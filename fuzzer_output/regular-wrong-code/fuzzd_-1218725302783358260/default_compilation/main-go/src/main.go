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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(856), _dafny.IntOfInt64(-444))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	return (false) && ((_dafny.IntOfInt64(-520)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-618))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality())) > 0)
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-158), _dafny.IntOfInt64(228))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-158)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(228)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_1_v0, _dafny.IntOfInt64(-841)), !(true) || (false))
			}
		}
		return _coll0.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf(true)
		}
		return _dafny.SeqOf(true)
	})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(true, true)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(968), true)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-115), true)
	})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-981), !(false)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("t")
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((Companion_D2_.Create_DC7_(_dafny.IntOfInt64(644))).Dtor_cf10()), _dafny.IntOfInt64(76))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(507), _dafny.IntOfInt64(23))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(507)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(23)) < 0) {
				_coll1.Add(Companion_Default___.SafeModuloInt(_2_v0, _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())))
			}
		}
		return _coll1.ToSet()
	}()).Union((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gj")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality())).Intersection(_dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('q'))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(383))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	})), true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dyck")).Cardinality()), (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("r"))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-66))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_4_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-114))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_5_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	})))).Cardinality())).Cardinality()))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 D2, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(572))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))).Cardinality()), (func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(535), _dafny.IntOfInt64(768))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _7_v0 _dafny.Int
			_7_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(535)).Cmp(_7_v0) <= 0) && ((_7_v0).Cmp(_dafny.IntOfInt64(768)) < 0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_7_v0, _dafny.IntOfInt64(455)))
			}
		}
		return _coll2.ToSet()
	}()).Cardinality())).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sqbhaywk")).Cardinality()), _dafny.IntOfInt64(160))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('e')
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 _dafny.Int, p2 D6, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SeqOf(false, true), _dafny.SeqOf(true, true, false, false), _dafny.SeqOf(false, true), _dafny.SeqOf(false))).Difference(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if !(true) {
			return _dafny.SeqOf((Companion_D8_.Create_DC24_(_dafny.SetOf(!(true), true), _dafny.SeqOf(!(false)))).Dtor_cf31(), _dafny.SeqOf(true), _dafny.SeqOf(true, true))
		}
		return _dafny.SeqOf(_dafny.SeqOf(true), _dafny.SeqOf(!(false)), _dafny.SeqOf(true), _dafny.SeqOf(false, true, true, false), _dafny.SeqOf(true, true))
	})()))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 bool, globalState *GlobalState) D2 {
	var _source0 D6 = Companion_D6_.Create_DC18_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hqs")).Cardinality()))).Cardinality()))
	_ = _source0
	if _source0.Is_DC16() {
		var _8___mcc_h0 _dafny.Array = _source0.Get_().(D6_DC16).Cf22
		_ = _8___mcc_h0
		var _9___mcc_h1 _dafny.Int = _source0.Get_().(D6_DC16).Cf23
		_ = _9___mcc_h1
		var _10_cf23 _dafny.Int = _9___mcc_h1
		_ = _10_cf23
		var _11_cf22 _dafny.Array = _8___mcc_h0
		_ = _11_cf22
		return Companion_D2_.Create_DC8_(false, ((Companion_D12_.Create_DC36_(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("rtpervwnl")))).Dtor_cf52()).Cardinality(), false)
	} else if _source0.Is_DC17() {
		var _12___mcc_h2 bool = _source0.Get_().(D6_DC17).Cf24
		_ = _12___mcc_h2
		var _13_cf24 bool = _12___mcc_h2
		_ = _13_cf24
		return Companion_D2_.Create_DC8_(_13_cf24, _dafny.IntOfInt64(877), _13_cf24)
	} else if _source0.Is_DC18() {
		var _14___mcc_h3 _dafny.Int = _source0.Get_().(D6_DC18).Cf25
		_ = _14___mcc_h3
		var _15_cf25 _dafny.Int = _14___mcc_h3
		_ = _15_cf25
		return Companion_D2_.Create_DC8_(false, _15_cf25, !(!(!(true))))
	} else if _source0.Is_DC15() {
		var _16___mcc_h4 _dafny.Sequence = _source0.Get_().(D6_DC15).Cf21
		_ = _16___mcc_h4
		var _17_cf21 _dafny.Sequence = _16___mcc_h4
		_ = _17_cf21
		return Companion_D2_.Create_DC8_(false, (func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.SetOf(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(647))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}))).Cardinality()))), _dafny.SetOf(_dafny.IntOfInt64(-146), _dafny.IntOfInt64(401), _dafny.IntOfInt64(-293), _dafny.IntOfUint32((_17_cf21).Cardinality()), (_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(471))).Cardinality()), _dafny.IntOfInt64(386))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality()))).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _19_v0 _dafny.Set
				_19_v0 = interface{}(_compr_3).(_dafny.Set)
				if (_dafny.SetOf(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(647))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				}))).Cardinality()))), _dafny.SetOf(_dafny.IntOfInt64(-146), _dafny.IntOfInt64(401), _dafny.IntOfInt64(-293), _dafny.IntOfUint32((_17_cf21).Cardinality()), (_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(471))).Cardinality()), _dafny.IntOfInt64(386))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality()))).Contains(_19_v0) {
					_coll3.Add(_19_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(708))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg7 _dafny.Int) interface{} {
							return coer7(arg7)
						}
					}(func(_20_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					}))).Cardinality()))
				}
			}
			return _coll3.ToMap()
		}()).Cardinality(), false)
	} else {
		var _21___mcc_h5 D6 = _source0.Get_().(D6_DC19).Cf26
		_ = _21___mcc_h5
		var _22_cf26 D6 = _21___mcc_h5
		_ = _22_cf26
		return Companion_D2_.Create_DC8_(!(true), _dafny.IntOfInt64(-990), false)
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((!(!(false))) || (false), false)
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(true))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 D5, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('v'), _dafny.CodePoint('u'), (func() _dafny.CodePoint {
		if true {
			return (Companion_D9_.Create_DC25_(_dafny.CodePoint('k'))).Dtor_cf32()
		}
		return _dafny.CodePoint('w')
	})())
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	var _source1 D8 = Companion_D8_.Create_DC23_(_dafny.IntOfInt64(663))
	_ = _source1
	if _source1.Is_DC23() {
		var _23___mcc_h0 _dafny.Int = _source1.Get_().(D8_DC23).Cf29
		_ = _23___mcc_h0
		var _24_cf29 _dafny.Int = _23___mcc_h0
		_ = _24_cf29
		return (_dafny.MultiSetOf(_24_cf29, _24_cf29, _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(240)))
	} else if _source1.Is_DC24() {
		var _25___mcc_h1 _dafny.Set = _source1.Get_().(D8_DC24).Cf30
		_ = _25___mcc_h1
		var _26___mcc_h2 _dafny.Sequence = _source1.Get_().(D8_DC24).Cf31
		_ = _26___mcc_h2
		var _27_cf31 _dafny.Sequence = _26___mcc_h2
		_ = _27_cf31
		var _28_cf30 _dafny.Set = _25___mcc_h1
		_ = _28_cf30
		return _dafny.MultiSetOf(_dafny.IntOfInt64(267))
	} else {
		var _29___mcc_h3 _dafny.MultiSet = _source1.Get_().(D8_DC22).Cf28
		_ = _29___mcc_h3
		var _30_cf28 _dafny.MultiSet = _29___mcc_h3
		_ = _30_cf28
		return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _dafny.IntOfInt64(992))).Cardinality())).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jgetcs")).Cardinality()))))
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC24_(_dafny.SetOf(false, true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if !(false) {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(459))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_31_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(220))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}(func(_32_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				}))).Cardinality())
			}))
		}
		return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-19), false), !(false))).Cardinality(), _dafny.IntOfInt64(49))
	})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-953))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_33_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-289)
	})))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D5_.Create_DC14_(true), Companion_D5_.Create_DC14_(true)), _dafny.SeqOf(Companion_D5_.Create_DC14_(true), Companion_D5_.Create_DC14_(false))), _dafny.SeqOf(Companion_D5_.Create_DC14_(true), Companion_D5_.Create_DC14_(true)))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC14_(!(false))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(739))).Uint32(), func(coer11 func(_dafny.Int) D5) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_34_i0 _dafny.Int) D5 {
		return Companion_D5_.Create_DC13_(true, false)
	}))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Map, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC26_()
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Map, _dafny.Sequence) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Map = _dafny.EmptyMap
	_ = r1
	var r2 _dafny.Sequence = _dafny.EmptySeq
	_ = r2
	if p3 {
		var _35_v0 _dafny.Sequence
		_ = _35_v0
		_35_v0 = _dafny.UnicodeSeqOfUtf8Bytes("gnxtuopx")
		var _36_v1 _dafny.MultiSet
		_ = _36_v1
		_36_v1 = _dafny.MultiSetOf(p3)
		var _37_v2 _dafny.Array
		_ = _37_v2
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) D2 = (func(_38_v1 _dafny.MultiSet) func(_dafny.Int) D2 {
				return func(_39_i2 _dafny.Int) D2 {
					return Companion_D2_.Create_DC7_((_dafny.MultiSetOf(_38_v1)).Cardinality())
				}
			})(_36_v1)
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
		_37_v2 = _nw0
		var _40_v3 T2
		_ = _40_v3
		var _nw1 *C3 = New_C3_()
		_ = _nw1
		_nw1.Ctor__(_37_v2, _dafny.CodePoint('v'), _dafny.IntOfInt64(419))
		_40_v3 = _nw1
		var _41_v4 _dafny.Map
		_ = _41_v4
		_41_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(668))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_42_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-194)
		})), _40_v3)
		var _43_v5 _dafny.Map
		_ = _43_v5
		_43_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-89))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_44_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})))
		var _45_v6 _dafny.Array
		_ = _45_v6
		var _nwElement0_0 _dafny.Sequence = _35_v0
		_ = _nwElement0_0
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(13))
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_0, 0)
		(_nw2).ArraySet1(Companion_Default___.Fm11((func() _dafny.Int {
			if (_36_v1).Contains(p3) {
				return (_36_v1).Multiplicity(p3)
			}
			return _dafny.IntOfUint32((_35_v0).Cardinality())
		})(), p2, p3, globalState), 1)
		(_nw2).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(120))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_46_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})), 2)
		(_nw2).ArraySet1(_35_v0, 3)
		(_nw2).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_35_v0, Companion_Default___.Fm11((_41_v4).Cardinality(), p0, p3, globalState)), 4)
		(_nw2).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("exf"), 5)
		(_nw2).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ivgkwls"), 6)
		(_nw2).ArraySet1((func() _dafny.Sequence {
			if (_43_v5).Contains(p3) {
				return (_43_v5).Get(p3).(_dafny.Sequence)
			}
			return _35_v0
		})(), 7)
		(_nw2).ArraySet1(_35_v0, 8)
		(_nw2).ArraySet1(_35_v0, 9)
		(_nw2).ArraySet1(_35_v0, 10)
		(_nw2).ArraySet1(_35_v0, 11)
		(_nw2).ArraySet1(_35_v0, 12)
		_45_v6 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))
		_ = _index0
		(_45_v6).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(473))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_47_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})), (_index0).Int())
		var _48_v7 _dafny.CodePoint
		_ = _48_v7
		_48_v7 = _dafny.CodePoint('c')
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))
		_ = _index1
		var _rhs0 _dafny.Int = (p1).Times(p0)
		_ = _rhs0
		var _rhs1 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_35_v0, (Companion_Default___.SafeIndex(_40_v3.F12(), _dafny.IntOfUint32((_35_v0).Cardinality()))).Uint32(), _48_v7)
		_ = _rhs1
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 _dafny.Array = _45_v6
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))
		_ = _lhs2
		_lhs0.F5 = _rhs0
		(_lhs1).ArraySet1(_rhs1, (_lhs2).Int())
		var _49_v8 _dafny.Sequence
		_ = _49_v8
		_49_v8 = _dafny.SeqOf(_40_v3.F12())
		var _50_v9 _dafny.Map
		_ = _50_v9
		_50_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_45_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))).Int()).(_dafny.Sequence)).Cardinality()), p0)
		var _51_v10 _dafny.Array
		_ = _51_v10
		var _nwElement0_1 _dafny.Sequence = _49_v8
		_ = _nwElement0_1
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(13))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_1, 0)
		(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_49_v8, _49_v8), 1)
		(_nw3).ArraySet1(_49_v8, 2)
		(_nw3).ArraySet1(_49_v8, 3)
		(_nw3).ArraySet1(_49_v8, 4)
		(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_49_v8, _dafny.SeqOf(p1)), 5)
		(_nw3).ArraySet1(_49_v8, 6)
		(_nw3).ArraySet1(_49_v8, 7)
		(_nw3).ArraySet1(Companion_Default___.Fm23(p3, _dafny.IntOfInt64(16), _dafny.IntOfInt64(-473), globalState), 8)
		(_nw3).ArraySet1(_49_v8, 9)
		(_nw3).ArraySet1(_dafny.SeqOf((_50_v9).Cardinality()), 10)
		(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_49_v8, _49_v8), 11)
		(_nw3).ArraySet1(_dafny.SeqOf(p2), 12)
		_51_v10 = _nw3
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_51_v10), 0))
		_ = _index2
		(_51_v10).ArraySet1(_49_v8, (_index2).Int())
		var _52_v11 _dafny.Map
		_ = _52_v11
		_52_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _49_v8)
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_51_v10), 0))
		_ = _index3
		(_51_v10).ArraySet1((func() _dafny.Sequence {
			if (_52_v11).Contains(true) {
				return (_52_v11).Get(true).(_dafny.Sequence)
			}
			return _49_v8
		})(), (_index3).Int())
		var _53_v12 D0
		_ = _53_v12
		_53_v12 = Companion_D0_.Create_DC0_(p3)
		var _source2 D0 = _53_v12
		_ = _source2
		if _source2.Is_DC1() {
			var _54___mcc_h0 _dafny.Array = _source2.Get_().(D0_DC1).Cf1
			_ = _54___mcc_h0
			var _55_cf1 _dafny.Array = _54___mcc_h0
			_ = _55_cf1
			var _56_v13 _dafny.Map
			_ = _56_v13
			_56_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v3.F12(), p3)
			_56_v13 = (_56_v13).Update(p2, p3)
			(_40_v3).F12_set_(p1)
			var _57_v14 D5
			_ = _57_v14
			_57_v14 = Companion_D5_.Create_DC14_(false)
			var _58_v15 _dafny.Sequence
			_ = _58_v15
			_58_v15 = _dafny.SeqOf(Companion_D5_.Create_DC14_(p3), _57_v14)
			var _59_v16 _dafny.Sequence
			_ = _59_v16
			_59_v16 = _dafny.SeqOf(_58_v15, _58_v15)
			var _60_v17 _dafny.Sequence
			_ = _60_v17
			_60_v17 = _dafny.SeqOf(_37_v2, _37_v2)
			var _61_v18 T1
			_ = _61_v18
			var _nw4 *C3 = New_C3_()
			_ = _nw4
			_nw4.Ctor__((_60_v17).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf((_dafny.Zero).Minus(p2))).Cardinality(), _dafny.IntOfUint32((_60_v17).Cardinality()))).Uint32()).(_dafny.Array), _dafny.CodePoint('s'), (_dafny.MultiSetOf(p3, p3)).Cardinality())
			_61_v18 = _nw4
			var _62_v19 _dafny.MultiSet
			_ = _62_v19
			_62_v19 = _dafny.MultiSetOf(_61_v18, _61_v18)
			var _63_v20 _dafny.Map
			_ = _63_v20
			_63_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_62_v19).Update(_61_v18, Companion_Default___.Abs(_40_v3.F12())), _48_v7)
			var _64_v21 _dafny.MultiSet
			_ = _64_v21
			_64_v21 = _dafny.MultiSetOf(p2, _40_v3.F12())
			var _65_v22 _dafny.Map
			_ = _65_v22
			_65_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_59_v16).Select((Companion_Default___.SafeIndex((_63_v20).Cardinality(), _dafny.IntOfUint32((_59_v16).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Fm24(p3, _40_v3.F12(), globalState)), (Companion_Default___.SafeIndex((_64_v21).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_59_v16).Select((Companion_Default___.SafeIndex((_63_v20).Cardinality(), _dafny.IntOfUint32((_59_v16).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Fm24(p3, _40_v3.F12(), globalState))).Cardinality()))).Uint32(), Companion_D5_.Create_DC14_(p3)), _40_v3)
			_65_v22 = (_65_v22).Update(_dafny.Companion_Sequence_.Concatenate(_58_v15, _dafny.Companion_Sequence_.Update(_58_v15, (Companion_Default___.SafeIndex(_61_v18.F12(), _dafny.IntOfUint32((_58_v15).Cardinality()))).Uint32(), _57_v14)), _40_v3)
			var _66_v23 bool
			_ = _66_v23
			_66_v23 = false
			_66_v23 = p3
		} else if _source2.Is_DC2() {
			var _67___mcc_h1 bool = _source2.Get_().(D0_DC2).Cf2
			_ = _67___mcc_h1
			var _68___mcc_h2 _dafny.Map = _source2.Get_().(D0_DC2).Cf3
			_ = _68___mcc_h2
			var _69___mcc_h3 _dafny.Map = _source2.Get_().(D0_DC2).Cf4
			_ = _69___mcc_h3
			var _70___mcc_h4 bool = _source2.Get_().(D0_DC2).Cf5
			_ = _70___mcc_h4
			var _71_cf5 bool = _70___mcc_h4
			_ = _71_cf5
			var _72_cf4 _dafny.Map = _69___mcc_h3
			_ = _72_cf4
			var _73_cf3 _dafny.Map = _68___mcc_h2
			_ = _73_cf3
			var _74_cf2 bool = _67___mcc_h1
			_ = _74_cf2
			_49_v8 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p2, (_dafny.Zero).Minus(p0), _dafny.IntOfInt64(973)), _dafny.SeqOf(_dafny.IntOfUint32((_35_v0).Cardinality())))
			var _75_v24 _dafny.Map
			_ = _75_v24
			_75_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v7, p3)
			var _76_v25 _dafny.Array
			_ = _76_v25
			var _nwElement0_2 _dafny.Int = (_40_v3.F12()).Minus(p2)
			_ = _nwElement0_2
			var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(18))
			_ = _nw5
			(_nw5).ArraySet1(_nwElement0_2, 0)
			(_nw5).ArraySet1(_dafny.IntOfUint32(((_45_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))).Int()).(_dafny.Sequence)).Cardinality()), 1)
			(_nw5).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_40_v3.F12(), p1)), 2)
			(_nw5).ArraySet1(p0, 3)
			(_nw5).ArraySet1(_40_v3.F12(), 4)
			(_nw5).ArraySet1((p0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("die")).Cardinality())), 5)
			(_nw5).ArraySet1((_dafny.Zero).Minus((Companion_Default___.Fm0(_dafny.IntOfInt64(493), (_75_v24).Cardinality(), globalState)).Times(_dafny.IntOfInt64(685))), 6)
			(_nw5).ArraySet1(p1, 7)
			(_nw5).ArraySet1(_dafny.IntOfUint32(((_45_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))).Int()).(_dafny.Sequence)).Cardinality()), 8)
			(_nw5).ArraySet1((_dafny.Zero).Minus(_40_v3.F12()), 9)
			(_nw5).ArraySet1(p2, 10)
			(_nw5).ArraySet1(p1, 11)
			(_nw5).ArraySet1(p1, 12)
			(_nw5).ArraySet1(p1, 13)
			(_nw5).ArraySet1(_dafny.IntOfUint32((_35_v0).Cardinality()), 14)
			(_nw5).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(540))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_77_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_78_i5 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_77_v8).Cardinality())
				}
			})(_49_v8)))).Cardinality())).Minus(p0), 15)
			(_nw5).ArraySet1(_40_v3.F12(), 16)
			(_nw5).ArraySet1(_40_v3.F12(), 17)
			_76_v25 = _nw5
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_76_v25), 0))
			_ = _index4
			(_76_v25).ArraySet1((_40_v3.F12()).Times(_40_v3.F12()), (_index4).Int())
			var _79_v26 _dafny.Set
			_ = _79_v26
			_79_v26 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("scckxsky"), (_45_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))).Int()).(_dafny.Sequence))
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_76_v25), 0))
			_ = _index5
			(_76_v25).ArraySet1((func() _dafny.Int {
				if _71_cf5 {
					return p2
				}
				return (_79_v26).Cardinality()
			})(), (_index5).Int())
			_48_v7 = _48_v7
			_71_cf5 = !((p2).Cmp(p0) < 0)
		} else if _source2.Is_DC0() {
			var _80___mcc_h5 bool = _source2.Get_().(D0_DC0).Cf0
			_ = _80___mcc_h5
			var _81_cf0 bool = _80___mcc_h5
			_ = _81_cf0
			var _82_v28 _dafny.Set
			_ = _82_v28
			_82_v28 = _dafny.SetOf(p1)
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__(func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((Companion_Default___.Fm23(_81_cf0, _40_v3.F12(), p2, globalState)).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _83_v27 _dafny.Int
					_83_v27 = interface{}(_compr_4).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm23(_81_cf0, _40_v3.F12(), p2, globalState), _83_v27) {
						_coll4.Add((_83_v27).Times(_40_v3.F12()), _81_cf0)
					}
				}
				return _coll4.ToMap()
			}(), (_82_v28).Cardinality())
			_40_v3 = _nw6
			var _84_v29 _dafny.Array
			_ = _84_v29
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw7
			_84_v29 = _nw7
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_84_v29), 0))
			_ = _index6
			(_84_v29).ArraySet1(_81_cf0, (_index6).Int())
			var _85_v30 D5
			_ = _85_v30
			_85_v30 = Companion_D5_.Create_DC14_(false)
			var _86_v31 _dafny.Sequence
			_ = _86_v31
			_86_v31 = _dafny.SeqOf((_45_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))).Int()).(_dafny.Sequence))
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_84_v29), 0))
			_ = _index7
			var _rhs2 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("y"), _35_v0), _86_v31)
			_ = _rhs2
			var _rhs3 D5 = Companion_Default___.Fm25(_81_cf0, _81_cf0, Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(p2, (_dafny.Zero).Minus(p2), globalState), _40_v3.F12()), globalState)
			_ = _rhs3
			var _rhs4 bool = _81_cf0
			_ = _rhs4
			var _lhs3 _dafny.Array = _84_v29
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_84_v29), 0))
			_ = _lhs4
			(_lhs3).ArraySet1(_rhs2, (_lhs4).Int())
			_85_v30 = _rhs3
			_81_cf0 = _rhs4
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_84_v29), 0))
			_ = _index8
			(_84_v29).ArraySet1(_81_cf0, (_index8).Int())
			var _rhs5 _dafny.Int = _40_v3.F12()
			_ = _rhs5
			var _rhs6 _dafny.Int = (func() _dafny.Int {
				if (_84_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_84_v29), 0))).Int()).(bool) {
					return _dafny.IntOfUint32(((_45_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))).Int()).(_dafny.Sequence)).Cardinality())
				}
				return ((_51_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_51_v10), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((_51_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_51_v10), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int)
			})()
			_ = _rhs6
			var _rhs7 _dafny.Int = (_dafny.Zero).Minus(_40_v3.F12())
			_ = _rhs7
			var _rhs8 _dafny.CodePoint = _48_v7
			_ = _rhs8
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			_lhs5.F5 = _rhs5
			_lhs6.F5 = _rhs6
			_lhs7.F5 = _rhs7
			_48_v7 = _rhs8
		} else {
			var _87___mcc_h6 D0 = _source2.Get_().(D0_DC3).Cf6
			_ = _87___mcc_h6
			var _88_cf6 D0 = _87___mcc_h6
			_ = _88_cf6
			var _89_v32 _dafny.Array
			_ = _89_v32
			var _nw8 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
			_ = _nw8
			_89_v32 = _nw8
			var _90_v33 *C3
			_ = _90_v33
			var _nw9 *C3 = New_C3_()
			_ = _nw9
			_nw9.Ctor__(_37_v2, _48_v7, p1)
			_90_v33 = _nw9
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_89_v32), 0))
			_ = _index9
			(_89_v32).ArraySet1(_90_v33, (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_89_v32), 0))
			_ = _index10
			var _nw10 *C3 = New_C3_()
			_ = _nw10
			_nw10.Ctor__(_37_v2, _48_v7, (_dafny.Zero).Minus(p1))
			(_89_v32).ArraySet1(_nw10, (_index10).Int())
			var _91_v34 _dafny.Array
			_ = _91_v34
			var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw11
			_91_v34 = _nw11
			var _92_v35 _dafny.MultiSet
			_ = _92_v35
			_92_v35 = _dafny.MultiSetOf(_91_v34, _91_v34)
			(_40_v3).F12_set_((func() _dafny.Int {
				if (_92_v35).Contains(_91_v34) {
					return (_92_v35).Multiplicity(_91_v34)
				}
				return _dafny.IntOfInt64(432)
			})())
			var _93_v36 _dafny.Map
			_ = _93_v36
			_93_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_40_v3.F12()).Times(_40_v3.F12()), p3)
			_93_v36 = (_93_v36).Update((_40_v3.F12()).Times((_36_v1).Cardinality()), p3)
			var _94_v37 _dafny.Set
			_ = _94_v37
			_94_v37 = _dafny.SetOf(_40_v3.F12())
			var _95_v38 *C4
			_ = _95_v38
			var _nw12 *C4 = New_C4_()
			_ = _nw12
			_nw12.Ctor__(((_94_v37).Union(_94_v37)).Cardinality(), (func() _dafny.Int {
				if Companion_Default___.Fm1(_dafny.IntOfInt64(-148), p3, p0, p3, globalState) {
					return _40_v3.F12()
				}
				return _40_v3.F12()
			})())
			_95_v38 = _nw12
		}
		var _96_v39 _dafny.Sequence
		_ = _96_v39
		_96_v39 = _dafny.SeqOf(p3)
		if !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(p3), _dafny.Companion_Sequence_.Concatenate(_96_v39, _96_v39)) {
			(_40_v3).F12_set_((func() _dafny.Int {
				if (_50_v9).Contains(p0) {
					return (_50_v9).Get(p0).(_dafny.Int)
				}
				return p0
			})())
			var _97_v40 _dafny.Map
			_ = _97_v40
			_97_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
			var _98_v41 _dafny.Array
			_ = _98_v41
			var _nwElement0_3 bool = Companion_Default___.Fm1(p0, false, (_49_v8).Select((Companion_Default___.SafeIndex((_36_v1).Cardinality(), _dafny.IntOfUint32((_49_v8).Cardinality()))).Uint32()).(_dafny.Int), p3, globalState)
			_ = _nwElement0_3
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(15))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_3, 0)
			(_nw13).ArraySet1(p3, 1)
			(_nw13).ArraySet1(p3, 2)
			(_nw13).ArraySet1(p3, 3)
			(_nw13).ArraySet1(p3, 4)
			(_nw13).ArraySet1(p3, 5)
			(_nw13).ArraySet1((p3) || (!(false)), 6)
			(_nw13).ArraySet1(false, 7)
			(_nw13).ArraySet1(p3, 8)
			(_nw13).ArraySet1(!(_dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm23(p3, Companion_Default___.Fm0((_97_v40).Cardinality(), p2, globalState), p1, globalState), (_51_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_51_v10), 0))).Int()).(_dafny.Sequence))), 9)
			(_nw13).ArraySet1(p3, 10)
			(_nw13).ArraySet1(p3, 11)
			(_nw13).ArraySet1(p3, 12)
			(_nw13).ArraySet1(p3, 13)
			(_nw13).ArraySet1((_96_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.IntOfUint32((_96_v39).Cardinality()))).Uint32()).(bool), 14)
			_98_v41 = _nw13
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_98_v41), 0))
			_ = _index11
			(_98_v41).ArraySet1(p3, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_98_v41), 0))
			_ = _index12
			(_98_v41).ArraySet1(p3, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_98_v41), 0))
			_ = _index13
			(_98_v41).ArraySet1((_98_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_98_v41), 0))).Int()).(bool), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_98_v41), 0))
			_ = _index14
			(_98_v41).ArraySet1(p3, (_index14).Int())
			(_40_v3).F12_set_((func() _dafny.Int {
				if (_97_v40).Contains((_98_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_98_v41), 0))).Int()).(bool)) {
					return (_97_v40).Get((_98_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_98_v41), 0))).Int()).(bool)).(_dafny.Int)
				}
				return p2
			})())
		} else {
			var _99_v42 _dafny.Array
			_ = _99_v42
			var _nw14 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
			_ = _nw14
			_99_v42 = _nw14
			var _100_v43 _dafny.Array
			_ = _100_v43
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_1
			var _nw15 _dafny.Array
			_ = _nw15
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw15 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_101_p3 bool) func(_dafny.Int) bool {
					return func(_102_i6 _dafny.Int) bool {
						return _101_p3
					}
				})(p3)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw15 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw15).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw15).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_100_v43 = _nw15
			var _103_v44 _dafny.Map
			_ = _103_v44
			_103_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v42, _100_v43)
			_103_v44 = _103_v44
			var _104_v45 D5
			_ = _104_v45
			_104_v45 = Companion_D5_.Create_DC13_(p3, p3)
			var _105_v46 _dafny.Sequence
			_ = _105_v46
			_105_v46 = _dafny.SeqOf(_104_v45, _104_v45, _104_v45, _104_v45)
			var _106_v47 _dafny.Map
			_ = _106_v47
			_106_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
			var _107_v48 _dafny.Set
			_ = _107_v48
			_107_v48 = _dafny.SetOf(_105_v46, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_104_v45, _104_v45), _105_v46), Companion_Default___.Fm26(p0, Companion_Default___.Fm15(_106_v47, _40_v3.F12(), globalState), p3, _48_v7, globalState))
			_107_v48 = _107_v48
			var _108_v49 bool
			_ = _108_v49
			_108_v49 = false
			var _109_v50 _dafny.Array
			_ = _109_v50
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
			_ = _nw16
			_109_v50 = _nw16
			var _110_v51 _dafny.Map
			_ = _110_v51
			_110_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _100_v43)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_109_v50), 0))
			_ = _index15
			(_109_v50).ArraySet1(_110_v51, (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_109_v50), 0))
			_ = _index16
			var _rhs9 bool = Companion_Default___.Fm1(p0, _108_v49, (p0).Plus(_dafny.IntOfInt64(-415)), p3, globalState)
			_ = _rhs9
			var _rhs10 _dafny.Map = _110_v51
			_ = _rhs10
			var _lhs8 _dafny.Array = _109_v50
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_109_v50), 0))
			_ = _lhs9
			_108_v49 = _rhs9
			(_lhs8).ArraySet1(_rhs10, (_lhs9).Int())
			var _111_v52 D10
			_ = _111_v52
			_111_v52 = Companion_D10_.Create_DC29_(_40_v3)
			var _112_v53 _dafny.Array
			_ = _112_v53
			var _nwElement0_4 T2 = _40_v3
			_ = _nwElement0_4
			var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(26))
			_ = _nw17
			(_nw17).ArraySet1(_nwElement0_4, 0)
			(_nw17).ArraySet1(_40_v3, 1)
			(_nw17).ArraySet1(_40_v3, 2)
			(_nw17).ArraySet1(_40_v3, 3)
			(_nw17).ArraySet1(_40_v3, 4)
			(_nw17).ArraySet1(_40_v3, 5)
			(_nw17).ArraySet1(_40_v3, 6)
			(_nw17).ArraySet1((func() T2 {
				if !(true) {
					return _40_v3
				}
				return _40_v3
			})(), 7)
			(_nw17).ArraySet1(_40_v3, 8)
			(_nw17).ArraySet1(_40_v3, 9)
			(_nw17).ArraySet1(_40_v3, 10)
			(_nw17).ArraySet1(_40_v3, 11)
			(_nw17).ArraySet1(_40_v3, 12)
			(_nw17).ArraySet1(_40_v3, 13)
			(_nw17).ArraySet1(_40_v3, 14)
			(_nw17).ArraySet1(_40_v3, 15)
			(_nw17).ArraySet1(_40_v3, 16)
			(_nw17).ArraySet1(_40_v3, 17)
			(_nw17).ArraySet1(_40_v3, 18)
			(_nw17).ArraySet1(_40_v3, 19)
			(_nw17).ArraySet1(_40_v3, 20)
			(_nw17).ArraySet1((func() T2 {
				if !(p3) {
					return _40_v3
				}
				return _40_v3
			})(), 21)
			(_nw17).ArraySet1(_40_v3, 22)
			(_nw17).ArraySet1(_40_v3, 23)
			(_nw17).ArraySet1((func() T2 {
				if _108_v49 {
					return (_111_v52).Dtor_cf39()
				}
				return _40_v3
			})(), 24)
			(_nw17).ArraySet1(_40_v3, 25)
			_112_v53 = _nw17
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_112_v53), 0))
			_ = _index17
			(_112_v53).ArraySet1(_40_v3, (_index17).Int())
			var _113_v54 D8
			_ = _113_v54
			_113_v54 = Companion_D8_.Create_DC23_(_40_v3.F12())
			var _114_v55 D9
			_ = _114_v55
			_114_v55 = Companion_D9_.Create_DC26_()
			var _115_v56 *C1
			_ = _115_v56
			var _nw18 *C1 = New_C1_()
			_ = _nw18
			_nw18.Ctor__(_dafny.CodePoint('g'), _48_v7, _dafny.IntOfInt64(696))
			_115_v56 = _nw18
			var _116_v57 _dafny.Set
			_ = _116_v57
			_116_v57 = _dafny.SetOf(_115_v56, _115_v56)
			var _117_v58 _dafny.Map
			_ = _117_v58
			_117_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v3.F12(), p3)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_112_v53), 0))
			_ = _index18
			var _rhs11 T2 = _40_v3
			_ = _rhs11
			var _rhs12 T2 = (func() T2 {
				if _108_v49 {
					return _40_v3
				}
				return _40_v3
			})()
			_ = _rhs12
			var _rhs13 _dafny.CodePoint = _48_v7
			_ = _rhs13
			var _rhs14 D8 = Companion_D8_.Create_DC23_(Companion_Default___.SafeModuloInt(p2, (_116_v57).Cardinality()))
			_ = _rhs14
			var _rhs15 D9 = Companion_Default___.Fm27(_117_v58, p3, _dafny.Companion_Sequence_.Concatenate(_35_v0, (_45_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_45_v6), 0))).Int()).(_dafny.Sequence)), globalState)
			_ = _rhs15
			var _lhs10 _dafny.Array = _112_v53
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_112_v53), 0))
			_ = _lhs11
			_40_v3 = _rhs11
			(_lhs10).ArraySet1(_rhs12, (_lhs11).Int())
			_48_v7 = _rhs13
			_113_v54 = _rhs14
			_114_v55 = _rhs15
			_35_v0 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(49))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_118_v56 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_119_i7 _dafny.Int) _dafny.CodePoint {
					return (_118_v56).F15()
				}
			})(_115_v56)))
		}
		var _120_v59 _dafny.MultiSet
		_ = _120_v59
		_120_v59 = _dafny.MultiSetOf(p2)
		var _121_v60 _dafny.Map
		_ = _121_v60
		_121_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
		r0 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_40_v3.F12(), _dafny.IntOfInt64(-35)), (func() _dafny.Int {
			if (_120_v59).Contains((_dafny.Zero).Minus(p1)) {
				return (_120_v59).Multiplicity((_dafny.Zero).Minus(p1))
			}
			return (_121_v60).Cardinality()
		})())
	} else {
		var _122_v61 bool
		_ = _122_v61
		_122_v61 = true
		var _123_v62 _dafny.MultiSet
		_ = _123_v62
		_123_v62 = _dafny.MultiSetOf(p3)
		var _124_v63 _dafny.MultiSet
		_ = _124_v63
		_124_v63 = _dafny.MultiSetOf(p1, (_123_v62).Cardinality())
		_122_v61 = (_dafny.MultiSetOf(Companion_Default___.Fm0(_dafny.IntOfInt64(864), p1, globalState))).Equals((_124_v63).Union(_124_v63))
		var _125_v64 D10
		_ = _125_v64
		_125_v64 = Companion_D10_.Create_DC30_(p0)
		(globalState).F5 = ((_125_v64).Dtor_cf40()).Minus((p0).Times(p1))
		var _126_v65 _dafny.Array
		_ = _126_v65
		var _nwElement0_5 bool = _122_v61
		_ = _nwElement0_5
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(24))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_5, 0)
		(_nw19).ArraySet1(_122_v61, 1)
		(_nw19).ArraySet1(p3, 2)
		(_nw19).ArraySet1(_122_v61, 3)
		(_nw19).ArraySet1(_122_v61, 4)
		(_nw19).ArraySet1(true, 5)
		(_nw19).ArraySet1(false, 6)
		(_nw19).ArraySet1(!(_122_v61), 7)
		(_nw19).ArraySet1(!(_122_v61), 8)
		(_nw19).ArraySet1(p3, 9)
		(_nw19).ArraySet1(_122_v61, 10)
		(_nw19).ArraySet1(_122_v61, 11)
		(_nw19).ArraySet1(true, 12)
		(_nw19).ArraySet1(p3, 13)
		(_nw19).ArraySet1(p3, 14)
		(_nw19).ArraySet1(p3, 15)
		(_nw19).ArraySet1(p3, 16)
		(_nw19).ArraySet1(_122_v61, 17)
		(_nw19).ArraySet1(p3, 18)
		(_nw19).ArraySet1(false, 19)
		(_nw19).ArraySet1(p3, 20)
		(_nw19).ArraySet1(_122_v61, 21)
		(_nw19).ArraySet1((Companion_D5_.Create_DC13_(_122_v61, p3)).Dtor_cf19(), 22)
		(_nw19).ArraySet1(p3, 23)
		_126_v65 = _nw19
		var _127_v66 _dafny.Map
		_ = _127_v66
		_127_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v61, _126_v65)
		var _128_v67 D11
		_ = _128_v67
		_128_v67 = Companion_D11_.Create_DC32_(_127_v66)
		_127_v66 = ((_128_v67).Dtor_cf42()).Merge((_127_v66).Merge(_127_v66))
		(globalState).F5 = (func() _dafny.Int {
			if p3 {
				return p2
			}
			return (p2).Times(p1)
		})()
		var _129_v68 *C4
		_ = _129_v68
		var _nw20 *C4 = New_C4_()
		_ = _nw20
		_nw20.Ctor__(p1, p0)
		_129_v68 = _nw20
	}
	var _130_v69 _dafny.Set
	_ = _130_v69
	_130_v69 = _dafny.SetOf(p3)
	var _131_v70 _dafny.Sequence
	_ = _131_v70
	_131_v70 = _dafny.SeqOf(p2, p0)
	var _132_v71 _dafny.Array
	_ = _132_v71
	var _nwElement0_6 _dafny.Int = (_dafny.IntOfInt64(646)).Minus(p1)
	_ = _nwElement0_6
	var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(17))
	_ = _nw21
	(_nw21).ArraySet1(_nwElement0_6, 0)
	(_nw21).ArraySet1(p2, 1)
	(_nw21).ArraySet1((p1).Plus((_130_v69).Cardinality()), 2)
	(_nw21).ArraySet1((func() _dafny.Int {
		if false {
			return p1
		}
		return _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality())
	})(), 3)
	(_nw21).ArraySet1(p0, 4)
	(_nw21).ArraySet1(p2, 5)
	(_nw21).ArraySet1(p1, 6)
	(_nw21).ArraySet1((p1).Plus(p1), 7)
	(_nw21).ArraySet1(p1, 8)
	(_nw21).ArraySet1(_dafny.IntOfUint32((_131_v70).Cardinality()), 9)
	(_nw21).ArraySet1(p1, 10)
	(_nw21).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p3, p3, p3)).Cardinality()), 11)
	(_nw21).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(652), p0), 12)
	(_nw21).ArraySet1(p0, 13)
	(_nw21).ArraySet1(p0, 14)
	(_nw21).ArraySet1(p2, 15)
	(_nw21).ArraySet1(Companion_Default___.SafeModuloInt(p2, p2), 16)
	_132_v71 = _nw21
	_132_v71 = _132_v71
	var _133_v72 bool
	_ = _133_v72
	_133_v72 = false
	var _134_v73 _dafny.CodePoint
	_ = _134_v73
	_134_v73 = _dafny.CodePoint('y')
	var _135_v74 *C1
	_ = _135_v74
	var _nw22 *C1 = New_C1_()
	_ = _nw22
	_nw22.Ctor__(_134_v73, _134_v73, p2)
	_135_v74 = _nw22
	var _136_v75 _dafny.Sequence
	_ = _136_v75
	_136_v75 = _dafny.UnicodeSeqOfUtf8Bytes("g")
	var _rhs16 bool = (func() bool {
		if _133_v72 {
			return _133_v72
		}
		return ((_dafny.Zero).Minus(_dafny.IntOfUint32((_136_v75).Cardinality()))).Cmp(p1) > 0
	})()
	_ = _rhs16
	var _rhs17 bool = (func() bool {
		if p3 {
			return _133_v72
		}
		return _133_v72
	})()
	_ = _rhs17
	var _rhs18 bool = _133_v72
	_ = _rhs18
	var _rhs19 *C1 = _135_v74
	_ = _rhs19
	var _rhs20 _dafny.Int = p2
	_ = _rhs20
	var _lhs12 *GlobalState = globalState
	_ = _lhs12
	_133_v72 = _rhs16
	_133_v72 = _rhs17
	_133_v72 = _rhs18
	_135_v74 = _rhs19
	_lhs12.F5 = _rhs20
	var _137_v76 _dafny.Array
	_ = _137_v76
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(21)
	_ = _len0_2
	var _nw23 _dafny.Array
	_ = _nw23
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw23 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) bool = (func(_138_v75 _dafny.Sequence, _139_p1 _dafny.Int, _140_v74 *C1) func(_dafny.Int) bool {
			return func(_141_i8 _dafny.Int) bool {
				return (_dafny.IntOfInt64(-476)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_138_v75, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_139_p1), _dafny.IntOfUint32((_138_v75).Cardinality()))).Uint32(), (_140_v74).F15())).Cardinality())) <= 0
			}
		})(_136_v75, p1, _135_v74)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw23 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw23).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw23).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_137_v76 = _nw23
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_137_v76), 0))
	_ = _index19
	(_137_v76).ArraySet1(p3, (_index19).Int())
	var _142_v77 _dafny.Map
	_ = _142_v77
	_142_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-649))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}((func(_143_v74 *C1) func(_dafny.Int) _dafny.CodePoint {
		return func(_144_i9 _dafny.Int) _dafny.CodePoint {
			return (_143_v74).F15()
		}
	})(_135_v74)))).Cardinality()), _136_v75)
	var _145_v78 D1
	_ = _145_v78
	_145_v78 = Companion_D1_.Create_DC4_(_132_v71)
	var _146_v79 _dafny.Set
	_ = _146_v79
	_146_v79 = _dafny.SetOf((_145_v78).Dtor_cf7(), _132_v71)
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_137_v76), 0))
	_ = _index20
	var _rhs21 bool = Companion_Default___.Fm1(p0, _133_v72, (_131_v70).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_131_v70).Cardinality()))).Uint32()).(_dafny.Int), p3, globalState)
	_ = _rhs21
	var _rhs22 _dafny.Int = Companion_Default___.Fm0(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
		if (_142_v77).Contains(_dafny.IntOfInt64(-120)) {
			return (_142_v77).Get(_dafny.IntOfInt64(-120)).(_dafny.Sequence)
		}
		return _136_v75
	})()).Cardinality()), p2), ((func() _dafny.Set {
		if p3 {
			return _146_v79
		}
		return _146_v79
	})()).Cardinality(), globalState)
	_ = _rhs22
	var _lhs13 _dafny.Array = _137_v76
	_ = _lhs13
	var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_137_v76), 0))
	_ = _lhs14
	(_lhs13).ArraySet1(_rhs21, (_lhs14).Int())
	r0 = _rhs22
	var _147_v80 *C4
	_ = _147_v80
	var _nw24 *C4 = New_C4_()
	_ = _nw24
	_nw24.Ctor__((_dafny.IntOfInt64(-263)).Plus(p0), p2)
	_147_v80 = _nw24
	_133_v72 = true
	r0 = Companion_Default___.SafeModuloInt(_147_v80.F11, (_147_v80).F10())
	var _148_v81 _dafny.Map
	_ = _148_v81
	_148_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfInt64(821))
	r1 = _148_v81
	var _149_v82 _dafny.Map
	_ = _149_v82
	_149_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_147_v80.F11, p0)
	var _150_v83 _dafny.Sequence
	_ = _150_v83
	_150_v83 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(347))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_151_i10 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	})), _136_v75, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_136_v75, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_136_v75).Cardinality()))).Uint32(), _134_v73), _136_v75), (Companion_Default___.SafeIndex((_149_v82).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_136_v75, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_136_v75).Cardinality()))).Uint32(), _134_v73), _136_v75)).Cardinality()))).Uint32(), _134_v73))
	r2 = (_150_v83).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-890), _dafny.IntOfUint32((_150_v83).Cardinality()))).Uint32()).(_dafny.Sequence)
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _152_v0 _dafny.Int
	_ = _152_v0
	_152_v0 = _dafny.IntOfInt64(172)
	var _153_v1 _dafny.Sequence
	_ = _153_v1
	_153_v1 = _dafny.SeqOf(_152_v0)
	var _154_v2 _dafny.Map
	_ = _154_v2
	_154_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v0, _153_v1)
	var _155_v3 _dafny.MultiSet
	_ = _155_v3
	_155_v3 = _dafny.MultiSetOf(_152_v0)
	var _156_v4 _dafny.CodePoint
	_ = _156_v4
	_156_v4 = _dafny.CodePoint('s')
	var _157_v5 _dafny.Sequence
	_ = _157_v5
	_157_v5 = _dafny.UnicodeSeqOfUtf8Bytes("o")
	var _158_v6 _dafny.Array
	_ = _158_v6
	var _nwElement0_7 _dafny.CodePoint = _156_v4
	_ = _nwElement0_7
	var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(11))
	_ = _nw25
	(_nw25).ArraySet1CodePoint(_nwElement0_7, 0)
	(_nw25).ArraySet1CodePoint(_156_v4, 1)
	(_nw25).ArraySet1CodePoint(_156_v4, 2)
	(_nw25).ArraySet1CodePoint(_156_v4, 3)
	(_nw25).ArraySet1CodePoint(_156_v4, 4)
	(_nw25).ArraySet1CodePoint(_156_v4, 5)
	(_nw25).ArraySet1CodePoint(_156_v4, 6)
	(_nw25).ArraySet1CodePoint(_156_v4, 7)
	(_nw25).ArraySet1CodePoint(_156_v4, 8)
	(_nw25).ArraySet1CodePoint((_157_v5).Select((Companion_Default___.SafeIndex(_152_v0, _dafny.IntOfUint32((_157_v5).Cardinality()))).Uint32()).(_dafny.CodePoint), 9)
	(_nw25).ArraySet1CodePoint(_156_v4, 10)
	_158_v6 = _nw25
	var _159_globalState *GlobalState
	_ = _159_globalState
	var _nw26 *GlobalState = New_GlobalState_()
	_ = _nw26
	_nw26.Ctor__((_dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if (_154_v2).Contains(_dafny.IntOfUint32((_153_v1).Cardinality())) {
			return (_154_v2).Get(_dafny.IntOfUint32((_153_v1).Cardinality())).(_dafny.Sequence)
		}
		return _dafny.SeqOf(_dafny.IntOfInt64(138))
	})())).Difference(_155_v3), true, false, true, _dafny.IntOfInt64(841), _dafny.IntOfInt64(338), true, false, _dafny.IntOfInt64(-192), _158_v6)
	_159_globalState = _nw26
	var _160_v7 _dafny.Map
	_ = _160_v7
	_160_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v0, _dafny.IntOfInt64(447))
	var _161_v8 bool
	_ = _161_v8
	_161_v8 = true
	var _162_v9 _dafny.Int
	_ = _162_v9
	var _163_v10 _dafny.Map
	_ = _163_v10
	var _164_v11 _dafny.Sequence
	_ = _164_v11
	var _out0 _dafny.Int
	_ = _out0
	var _out1 _dafny.Map
	_ = _out1
	var _out2 _dafny.Sequence
	_ = _out2
	_out0, _out1, _out2 = Companion_Default___.M0((func() _dafny.Int {
		if (_160_v7).Contains(_dafny.IntOfInt64(49)) {
			return (_160_v7).Get(_dafny.IntOfInt64(49)).(_dafny.Int)
		}
		return _152_v0
	})(), _152_v0, _152_v0, _161_v8, _159_globalState)
	_162_v9 = _out0
	_163_v10 = _out1
	_164_v11 = _out2
	_161_v8 = _161_v8
	var _hi0 _dafny.Int = _152_v0
	_ = _hi0
	for _165_i0 := _162_v9; _165_i0.Cmp(_hi0) < 0; _165_i0 = _165_i0.Plus(_dafny.One) {
		var _166_v12 _dafny.Map
		_ = _166_v12
		_166_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-269)).Minus(Companion_Default___.Fm0(_162_v9, _162_v9, _159_globalState)), _dafny.UnicodeSeqOfUtf8Bytes("ytxvk"))
		_157_v5 = (func() _dafny.Sequence {
			if (_166_v12).Contains((_dafny.Zero).Minus(_162_v9)) {
				return (_166_v12).Get((_dafny.Zero).Minus(_162_v9)).(_dafny.Sequence)
			}
			return _164_v11
		})()
		_161_v8 = _161_v8
		var _167_v13 _dafny.Sequence
		_ = _167_v13
		_167_v13 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xrlcvxj"), _157_v5, _dafny.UnicodeSeqOfUtf8Bytes("enoyjic"))
		var _168_v14 _dafny.Sequence
		_ = _168_v14
		_168_v14 = _dafny.SeqOf(_161_v8)
		var _169_v15 D0
		_ = _169_v15
		_169_v15 = Companion_D0_.Create_DC0_(_161_v8)
		var _170_v16 D0
		_ = _170_v16
		_170_v16 = Companion_D0_.Create_DC0_((_169_v15).Dtor_cf0())
		var _rhs23 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_167_v13, _dafny.Companion_Sequence_.Concatenate(_167_v13, _167_v13))
		_ = _rhs23
		var _rhs24 bool = !(_161_v8)
		_ = _rhs24
		var _rhs25 bool = !((_dafny.IntOfUint32((_168_v14).Cardinality())).Cmp(_165_i0) > 0) || (_161_v8)
		_ = _rhs25
		var _rhs26 _dafny.Int = _165_i0
		_ = _rhs26
		var _rhs27 bool = (_169_v15).Dtor_cf0()
		_ = _rhs27
		_167_v13 = _rhs23
		_161_v8 = _rhs24
		_161_v8 = _rhs25
		_152_v0 = _rhs26
		_161_v8 = _rhs27
		var _171_v17 _dafny.Array
		_ = _171_v17
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(29))
		_ = _nw27
		_171_v17 = _nw27
		_171_v17 = _171_v17
	}
	var _172_v18 _dafny.Array
	_ = _172_v18
	var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
	_ = _nw28
	_172_v18 = _nw28
	var _173_v19 _dafny.Array
	_ = _173_v19
	var _nwElement0_8 _dafny.Array = _172_v18
	_ = _nwElement0_8
	var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(14))
	_ = _nw29
	(_nw29).ArraySet1(_nwElement0_8, 0)
	(_nw29).ArraySet1(_172_v18, 1)
	(_nw29).ArraySet1((Companion_D1_.Create_DC4_(_172_v18)).Dtor_cf7(), 2)
	(_nw29).ArraySet1(_172_v18, 3)
	(_nw29).ArraySet1(_172_v18, 4)
	(_nw29).ArraySet1(_172_v18, 5)
	(_nw29).ArraySet1(_172_v18, 6)
	(_nw29).ArraySet1(_172_v18, 7)
	(_nw29).ArraySet1(_172_v18, 8)
	(_nw29).ArraySet1(_172_v18, 9)
	(_nw29).ArraySet1(_172_v18, 10)
	(_nw29).ArraySet1(_172_v18, 11)
	(_nw29).ArraySet1(_172_v18, 12)
	(_nw29).ArraySet1(_172_v18, 13)
	_173_v19 = _nw29
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_173_v19), 0))
	_ = _index21
	(_173_v19).ArraySet1(_172_v18, (_index21).Int())
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_173_v19), 0))
	_ = _index22
	(_173_v19).ArraySet1(_172_v18, (_index22).Int())
	var _174_v20 _dafny.Set
	_ = _174_v20
	_174_v20 = _dafny.SetOf(_dafny.IntOfInt64(-144))
	var _175_v21 _dafny.Int
	_ = _175_v21
	var _176_v22 _dafny.Map
	_ = _176_v22
	var _177_v23 _dafny.Sequence
	_ = _177_v23
	var _out3 _dafny.Int
	_ = _out3
	var _out4 _dafny.Map
	_ = _out4
	var _out5 _dafny.Sequence
	_ = _out5
	_out3, _out4, _out5 = Companion_Default___.M0(_152_v0, _162_v9, (_dafny.Zero).Minus((func() _dafny.Int {
		if true {
			return (_153_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_174_v20).Cardinality()), _dafny.IntOfUint32((_153_v1).Cardinality()))).Uint32()).(_dafny.Int)
		}
		return _dafny.IntOfUint32((_157_v5).Cardinality())
	})()), Companion_Default___.Fm1((_dafny.Zero).Minus(_152_v0), _161_v8, _dafny.IntOfInt64(-267), !(true), _159_globalState), _159_globalState)
	_175_v21 = _out3
	_176_v22 = _out4
	_177_v23 = _out5
	var _178_v24 _dafny.Array
	_ = _178_v24
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(5)
	_ = _len0_3
	var _nw30 _dafny.Array
	_ = _nw30
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw30 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) D2 = (func(_179_v0 _dafny.Int) func(_dafny.Int) D2 {
			return func(_180_i1 _dafny.Int) D2 {
				return Companion_D2_.Create_DC7_(_179_v0)
			}
		})(_152_v0)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw30 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw30).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw30).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_178_v24 = _nw30
	var _181_v25 *C3
	_ = _181_v25
	var _nw31 *C3 = New_C3_()
	_ = _nw31
	_nw31.Ctor__(_178_v24, _dafny.CodePoint('x'), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_175_v21, _175_v21)))
	_181_v25 = _nw31
	var _182_v26 _dafny.Array
	_ = _182_v26
	var _nw32 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
	_ = _nw32
	_182_v26 = _nw32
	var _183_v27 *C1
	_ = _183_v27
	var _nw33 *C1 = New_C1_()
	_ = _nw33
	_nw33.Ctor__(_156_v4, _156_v4, _162_v9)
	_183_v27 = _nw33
	var _184_v28 _dafny.Map
	_ = _184_v28
	_184_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v0, _161_v8)
	var _185_v29 _dafny.Map
	_ = _185_v29
	_185_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v28, (_183_v27).F15())
	var _186_v30 D9
	_ = _186_v30
	_186_v30 = Companion_D9_.Create_DC27_(_183_v27, (_155_v3).Cardinality(), _dafny.IntOfUint32((_157_v5).Cardinality()), _dafny.IntOfInt64(-208), _185_v29)
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_182_v26), 0))
	_ = _index23
	(_182_v26).ArraySet1((_186_v30).Dtor_cf33(), (_index23).Int())
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_182_v26), 0))
	_ = _index24
	var _nw34 *C1 = New_C1_()
	_ = _nw34
	_nw34.Ctor__(_156_v4, (_183_v27).F15(), (_dafny.IntOfUint32((_164_v11).Cardinality())).Minus(_152_v0))
	(_182_v26).ArraySet1(_nw34, (_index24).Int())
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))
	_ = _index25
	(_172_v18).ArraySet1(_152_v0, (_index25).Int())
	var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))
	_ = _index26
	(_172_v18).ArraySet1((_dafny.Zero).Minus((_181_v25).Fm4(_159_globalState)), (_index26).Int())
	var _187_v31 _dafny.Array
	_ = _187_v31
	var _nw35 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
	_ = _nw35
	_187_v31 = _nw35
	var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_187_v31), 0))
	_ = _index27
	(_187_v31).ArraySet1(_181_v25, (_index27).Int())
	var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_187_v31), 0))
	_ = _index28
	(_187_v31).ArraySet1((func() *C3 {
		if _161_v8 {
			return _181_v25
		}
		return _181_v25
	})(), (_index28).Int())
	var _188_v32 bool
	_ = _188_v32
	var _189_v33 _dafny.Array
	_ = _189_v33
	var _190_v34 _dafny.Map
	_ = _190_v34
	var _191_v35 _dafny.Int
	_ = _191_v35
	var _out6 bool
	_ = _out6
	var _out7 _dafny.Array
	_ = _out7
	var _out8 _dafny.Map
	_ = _out8
	var _out9 _dafny.Int
	_ = _out9
	_out6, _out7, _out8, _out9 = (_183_v27).M2(_159_globalState)
	_188_v32 = _out6
	_189_v33 = _out7
	_190_v34 = _out8
	_191_v35 = _out9
	var _192_v36 D5
	_ = _192_v36
	_192_v36 = Companion_D5_.Create_DC14_(_161_v8)
	if (_192_v36).Dtor_cf20() {
		var _193_v37 _dafny.Array
		_ = _193_v37
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(12))
		_ = _nw36
		_193_v37 = _nw36
		var _194_v38 T0
		_ = _194_v38
		var _nw37 *C2 = New_C2_()
		_ = _nw37
		_nw37.Ctor__(_188_v32, (_181_v25).Fm7(_159_globalState))
		_194_v38 = _nw37
		var _195_v39 _dafny.Map
		_ = _195_v39
		_195_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v37, _194_v38)
		_195_v39 = (_195_v39).Update(_193_v37, _194_v38)
		var _196_v40 _dafny.Array
		_ = _196_v40
		var _nwElement0_9 bool = !(_188_v32)
		_ = _nwElement0_9
		var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(5))
		_ = _nw38
		(_nw38).ArraySet1(_nwElement0_9, 0)
		(_nw38).ArraySet1(!(_161_v8), 1)
		(_nw38).ArraySet1(_188_v32, 2)
		(_nw38).ArraySet1(((_dafny.MultiSetOf(_191_v35)).Cardinality()).Cmp(_175_v21) < 0, 3)
		(_nw38).ArraySet1(_161_v8, 4)
		_196_v40 = _nw38
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))
		_ = _index29
		(_196_v40).ArraySet1(_161_v8, (_index29).Int())
		var _197_v41 D2
		_ = _197_v41
		_197_v41 = Companion_D2_.Create_DC8_(((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)).Cmp(_162_v9) >= 0, (func() _dafny.Int {
			if !(_188_v32) {
				return (_153_v1).Select((Companion_Default___.SafeIndex(_175_v21, _dafny.IntOfUint32((_153_v1).Cardinality()))).Uint32()).(_dafny.Int)
			}
			return (_155_v3).Cardinality()
		})(), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf((_183_v27).F15(), _dafny.CodePoint('n')), _177_v23))
		var _198_v42 _dafny.Map
		_ = _198_v42
		_198_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_dafny.ArrayCastTo((_173_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_173_v19), 0))).Int()))), _189_v33)
		var _199_v43 _dafny.Map
		_ = _199_v43
		_199_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_v32, _dafny.SetOf(_dafny.IntOfInt64(-880), _175_v21))
		var _200_v44 _dafny.Array
		_ = _200_v44
		var _nwElement0_10 _dafny.Set = (_dafny.SetOf(_191_v35)).Difference(_174_v20)
		_ = _nwElement0_10
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(12))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_10, 0)
		(_nw39).ArraySet1(Companion_Default___.Fm13(false, _161_v8, _159_globalState), 1)
		(_nw39).ArraySet1(_174_v20, 2)
		(_nw39).ArraySet1((_dafny.SetOf((_dafny.Zero).Minus(_175_v21))).Intersection(_174_v20), 3)
		(_nw39).ArraySet1(_dafny.SetOf(_191_v35), 4)
		(_nw39).ArraySet1(_174_v20, 5)
		(_nw39).ArraySet1((_dafny.SetOf((_198_v42).Cardinality())).Difference((func() _dafny.Set {
			if (_199_v43).Contains(_188_v32) {
				return (_199_v43).Get(_188_v32).(_dafny.Set)
			}
			return _174_v20
		})()), 6)
		(_nw39).ArraySet1(_dafny.SetOf(_194_v38.F12()), 7)
		(_nw39).ArraySet1((_174_v20).Union(_174_v20), 8)
		(_nw39).ArraySet1((_174_v20).Difference(_174_v20), 9)
		(_nw39).ArraySet1(_dafny.SetOf((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)), 10)
		(_nw39).ArraySet1(_174_v20, 11)
		_200_v44 = _nw39
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_200_v44), 0))
		_ = _index30
		(_200_v44).ArraySet1(_dafny.SetOf((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_152_v0)), (_index30).Int())
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))
		_ = _index31
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_200_v44), 0))
		_ = _index32
		var _rhs28 bool = _161_v8
		_ = _rhs28
		var _rhs29 _dafny.Int = (_181_v25).Fm7(_159_globalState)
		_ = _rhs29
		var _rhs30 D2 = func(_pat_let0_0 D2) D2 {
			return func(_201_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let1_0 bool) D2 {
					return func(_202_dt__update_hcf13_h0 bool) D2 {
						return func(_pat_let2_0 bool) D2 {
							return func(_203_dt__update_hcf11_h0 bool) D2 {
								return Companion_D2_.Create_DC8_(_203_dt__update_hcf11_h0, (_201_dt__update__tmp_h0).Dtor_cf12(), _202_dt__update_hcf13_h0)
							}(_pat_let2_0)
						}(false)
					}(_pat_let1_0)
				}(true)
			}(_pat_let0_0)
		}(_197_v41)
		_ = _rhs30
		var _rhs31 _dafny.Set = _174_v20
		_ = _rhs31
		var _lhs15 _dafny.Array = _196_v40
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))
		_ = _lhs16
		var _lhs17 _dafny.Array = _200_v44
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_200_v44), 0))
		_ = _lhs18
		(_lhs15).ArraySet1(_rhs28, (_lhs16).Int())
		_162_v9 = _rhs29
		_197_v41 = _rhs30
		(_lhs17).ArraySet1(_rhs31, (_lhs18).Int())
		var _204_v45 _dafny.Map
		_ = _204_v45
		_204_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v8, true)
		var _205_v46 _dafny.MultiSet
		_ = _205_v46
		_205_v46 = _dafny.MultiSetOf(_161_v8)
		var _206_v47 _dafny.Sequence
		_ = _206_v47
		_206_v47 = _dafny.SeqOf(_205_v46)
		var _207_v48 D5
		_ = _207_v48
		_207_v48 = Companion_D5_.Create_DC12_(_dafny.SeqOf(_dafny.MultiSetOf(_188_v32, !(_188_v32))))
		var _208_v49 _dafny.Sequence
		_ = _208_v49
		_208_v49 = _dafny.SeqOf(Companion_D5_.Create_DC12_(_206_v47), Companion_D5_.Create_DC12_(_206_v47), _207_v48, _207_v48, Companion_D5_.Create_DC12_(_206_v47))
		var _209_v50 D6
		_ = _209_v50
		_209_v50 = Companion_D6_.Create_DC15_((_183_v27).Fm5((func() bool {
			if (_204_v45).Contains(!(false)) {
				return (_204_v45).Get(!(false)).(bool)
			}
			return (_196_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))).Int()).(bool)
		})(), _161_v8, _188_v32, Companion_Default___.Fm1(_191_v35, true, (_153_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_208_v49).Cardinality()), _dafny.IntOfUint32((_153_v1).Cardinality()))).Uint32()).(_dafny.Int), _188_v32, _159_globalState), _159_globalState))
		var _pat_let_tv0 = _153_v1
		_ = _pat_let_tv0
		_209_v50 = func(_pat_let3_0 D6) D6 {
			return func(_210_dt__update__tmp_h1 D6) D6 {
				return func(_pat_let4_0 _dafny.Sequence) D6 {
					return func(_211_dt__update_hcf21_h0 _dafny.Sequence) D6 {
						return Companion_D6_.Create_DC15_(_211_dt__update_hcf21_h0)
					}(_pat_let4_0)
				}(_pat_let_tv0)
			}(_pat_let3_0)
		}(Companion_D6_.Create_DC15_(_153_v1))
		var _212_v51 _dafny.Array
		_ = _212_v51
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_4
		var _nw40 _dafny.Array
		_ = _nw40
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw40 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Sequence = (func(_213_v40 _dafny.Array, _214_v11 _dafny.Sequence, _215_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_216_i2 _dafny.Int) _dafny.Sequence {
					return (func() _dafny.Sequence {
						if (_213_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_213_v40), 0))).Int()).(bool) {
							return _214_v11
						}
						return _215_v23
					})()
				}
			})(_196_v40, _164_v11, _177_v23)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw40 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw40).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw40).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_212_v51 = _nw40
		_212_v51 = _212_v51
		var _217_v52 _dafny.Set
		_ = _217_v52
		_217_v52 = _dafny.SetOf(true, (_196_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))).Int()).(bool), _188_v32)
		if Companion_Default___.Fm1((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int), Companion_Default___.Fm1(_162_v9, (func() bool {
			if (_204_v45).Contains(_188_v32) {
				return (_204_v45).Get(_188_v32).(bool)
			}
			return _188_v32
		})(), _194_v38.F12(), Companion_Default___.Fm1(_dafny.IntOfInt64(125), (_196_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))).Int()).(bool), (_217_v52).Cardinality(), _188_v32, _159_globalState), _159_globalState), (_181_v25).Fm7(_159_globalState), _188_v32, _159_globalState) {
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))
			_ = _index33
			(_196_v40).ArraySet1(!(_161_v8) || ((_dafny.MultiSetFromSeq(_153_v1)).IsSubsetOf(Companion_Default___.Fm21(_152_v0, _175_v21, _156_v4, (_dafny.Zero).Minus(_191_v35), _159_globalState))), (_index33).Int())
			_161_v8 = _188_v32
			var _218_v53 bool
			_ = _218_v53
			var _219_v54 _dafny.Array
			_ = _219_v54
			var _220_v55 _dafny.Map
			_ = _220_v55
			var _221_v56 _dafny.Int
			_ = _221_v56
			var _out10 bool
			_ = _out10
			var _out11 _dafny.Array
			_ = _out11
			var _out12 _dafny.Map
			_ = _out12
			var _out13 _dafny.Int
			_ = _out13
			_out10, _out11, _out12, _out13 = (_181_v25).M2(_159_globalState)
			_218_v53 = _out10
			_219_v54 = _out11
			_220_v55 = _out12
			_221_v56 = _out13
			var _222_v57 _dafny.Map
			_ = _222_v57
			_222_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v38, _161_v8)
			var _223_v58 *C2
			_ = _223_v58
			var _nw41 *C2 = New_C2_()
			_ = _nw41
			_nw41.Ctor__(_161_v8, ((_222_v57).Merge(_222_v57)).Cardinality())
			_223_v58 = _nw41
			_196_v40 = _196_v40
		} else {
			(_159_globalState).F5 = (_dafny.Zero).Minus((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int))
			var _224_v59 _dafny.Map
			_ = _224_v59
			_224_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_181_v25).Fm7(_159_globalState), _157_v5)
			var _225_v60 _dafny.Set
			_ = _225_v60
			_225_v60 = _dafny.SetOf(_183_v27)
			_224_v59 = (_224_v59).Update(((_225_v60).Cardinality()).Times(_dafny.IntOfInt64(627)), _164_v11)
			_188_v32 = !((_196_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))).Int()).(bool)) || ((_196_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_196_v40), 0))).Int()).(bool))
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))
			_ = _index34
			(_172_v18).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_152_v0, _dafny.IntOfInt64(241))), Companion_Default___.Fm0((func() _dafny.Int {
				if (_160_v7).Contains((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)) {
					return (_160_v7).Get((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)).(_dafny.Int)
				}
				return _152_v0
			})(), _dafny.IntOfInt64(658), _159_globalState)), (_index34).Int())
			_177_v23 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-635))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_226_v11 _dafny.Sequence, _227_v9 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_228_i3 _dafny.Int) _dafny.CodePoint {
					return (_226_v11).Select((Companion_Default___.SafeIndex(_227_v9, _dafny.IntOfUint32((_226_v11).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_164_v11, _162_v9)))
		}
	} else {
		var _229_v61 *C3
		_ = _229_v61
		var _nw42 *C3 = New_C3_()
		_ = _nw42
		_nw42.Ctor__(_178_v24, _dafny.CodePoint('j'), _152_v0)
		_229_v61 = _nw42
		var _230_v62 _dafny.Sequence
		_ = _230_v62
		_230_v62 = _dafny.SeqOf(_188_v32)
		_230_v62 = _dafny.Companion_Sequence_.Update(_230_v62, (Companion_Default___.SafeIndex(_175_v21, _dafny.IntOfUint32((_230_v62).Cardinality()))).Uint32(), (func() bool {
			if true {
				return _188_v32
			}
			return Companion_Default___.Fm1(_152_v0, _161_v8, _152_v0, _188_v32, _159_globalState)
		})())
		var _231_v63 _dafny.Map
		_ = _231_v63
		_231_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v33, _175_v21)
		if !((_231_v63).Update(_172_v18, _175_v21)).Equals(_231_v63) {
			var _232_v64 _dafny.Map
			_ = _232_v64
			_232_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_v32, _182_v26)
			var _233_v65 _dafny.Array
			_ = _233_v65
			var _nwElement0_11 _dafny.Set = _174_v20
			_ = _nwElement0_11
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(3))
			_ = _nw43
			(_nw43).ArraySet1(_nwElement0_11, 0)
			(_nw43).ArraySet1(_174_v20, 1)
			(_nw43).ArraySet1(_dafny.SetOf((Companion_Default___.Fm19(_152_v0, _159_globalState)).Cardinality(), (_dafny.Zero).Minus((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)), _175_v21), 2)
			_233_v65 = _nw43
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_233_v65), 0))
			_ = _index35
			(_233_v65).ArraySet1(_174_v20, (_index35).Int())
			var _234_v66 D2
			_ = _234_v66
			_234_v66 = Companion_D2_.Create_DC9_(_152_v0)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_233_v65), 0))
			_ = _index36
			var _rhs32 _dafny.Map = (_232_v64).Update((_155_v3).IsSubsetOf(_155_v3), _182_v26)
			_ = _rhs32
			var _rhs33 _dafny.Set = _dafny.SetOf((_dafny.IntOfUint32(((_229_v61).Fm5(_188_v32, _188_v32, _188_v32, _161_v8, _159_globalState)).Cardinality())).Plus((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)), (_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int))
			_ = _rhs33
			var _rhs34 D2 = _234_v66
			_ = _rhs34
			var _rhs35 _dafny.Sequence = _177_v23
			_ = _rhs35
			var _lhs19 _dafny.Array = _233_v65
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_233_v65), 0))
			_ = _lhs20
			_232_v64 = _rhs32
			(_lhs19).ArraySet1(_rhs33, (_lhs20).Int())
			_234_v66 = _rhs34
			_177_v23 = _rhs35
			var _235_v67 *C2
			_ = _235_v67
			var _nw44 *C2 = New_C2_()
			_ = _nw44
			_nw44.Ctor__(_161_v8, _162_v9)
			_235_v67 = _nw44
			var _236_v68 _dafny.Array
			_ = _236_v68
			var _nw45 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
			_ = _nw45
			_236_v68 = _nw45
			var _237_v69 T2
			_ = _237_v69
			var _nw46 *C3 = New_C3_()
			_ = _nw46
			_nw46.Ctor__((_229_v61).F14(), _dafny.CodePoint('r'), _dafny.IntOfInt64(268))
			_237_v69 = _nw46
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_236_v68), 0))
			_ = _index37
			(_236_v68).ArraySet1(_237_v69, (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_236_v68), 0))
			_ = _index38
			var _rhs36 bool = ((_230_v62).Select((Companion_Default___.SafeIndex(_237_v69.F12(), _dafny.IntOfUint32((_230_v62).Cardinality()))).Uint32()).(bool)) == (_188_v32)
			_ = _rhs36
			var _rhs37 T2 = (func() T2 {
				if _161_v8 {
					return _237_v69
				}
				return _237_v69
			})()
			_ = _rhs37
			var _lhs21 _dafny.Array = _236_v68
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_236_v68), 0))
			_ = _lhs22
			_161_v8 = _rhs36
			(_lhs21).ArraySet1(_rhs37, (_lhs22).Int())
			var _238_v70 bool
			_ = _238_v70
			var _239_v71 _dafny.Array
			_ = _239_v71
			var _240_v72 _dafny.Map
			_ = _240_v72
			var _241_v73 _dafny.Int
			_ = _241_v73
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Array
			_ = _out15
			var _out16 _dafny.Map
			_ = _out16
			var _out17 _dafny.Int
			_ = _out17
			_out14, _out15, _out16, _out17 = (_181_v25).M2(_159_globalState)
			_238_v70 = _out14
			_239_v71 = _out15
			_240_v72 = _out16
			_241_v73 = _out17
			_177_v23 = _157_v5
		} else {
			var _242_v74 *C4
			_ = _242_v74
			var _nw47 *C4 = New_C4_()
			_ = _nw47
			_nw47.Ctor__((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int), _162_v9)
			_242_v74 = _nw47
			var _243_v75 _dafny.Sequence
			_ = _243_v75
			_243_v75 = _dafny.SeqOf(_242_v74, _242_v74, _242_v74)
			var _244_v76 _dafny.Map
			_ = _244_v76
			_244_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_176_v22).Cardinality(), _243_v75)
			_175_v21 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_244_v76).Contains((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)) {
					return (_244_v76).Get((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
				}
				return _243_v75
			})()).Cardinality()), _191_v35)
			var _245_v77 *C0
			_ = _245_v77
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__(_184_v28, _191_v35)
			_245_v77 = _nw48
			var _246_v78 _dafny.Map
			_ = _246_v78
			_246_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v77, _dafny.Companion_Sequence_.Update(_157_v5, (Companion_Default___.SafeIndex(_162_v9, _dafny.IntOfUint32((_157_v5).Cardinality()))).Uint32(), (_183_v27).F15()))
			_191_v35 = ((func() _dafny.Map {
				if false {
					return (_246_v78).Merge(_246_v78)
				}
				return (_246_v78).Merge(_246_v78)
			})()).Cardinality()
			_161_v8 = _188_v32
			_161_v8 = false
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_173_v19), 0))
			_ = _index39
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_5
			var _nw49 _dafny.Array
			_ = _nw49
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw49 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_247_v74 *C4) func(_dafny.Int) _dafny.Int {
					return func(_248_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_248_i4, (_247_v74).F10())
					}
				})(_242_v74)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw49 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw49).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw49).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			(_173_v19).ArraySet1(_nw49, (_index39).Int())
		}
		var _249_v79 D9
		_ = _249_v79
		_249_v79 = Companion_D9_.Create_DC26_()
		_249_v79 = (func() D9 {
			if !(_188_v32) {
				return _249_v79
			}
			return _249_v79
		})()
		var _250_v80 _dafny.Int
		_ = _250_v80
		var _251_v81 bool
		_ = _251_v81
		var _252_v82 _dafny.Int
		_ = _252_v82
		var _out18 _dafny.Int
		_ = _out18
		var _out19 bool
		_ = _out19
		var _out20 _dafny.Int
		_ = _out20
		_out18, _out19, _out20 = (_181_v25).M3(_dafny.MultiSetOf(false, !(_188_v32)), _159_globalState)
		_250_v80 = _out18
		_251_v81 = _out19
		_252_v82 = _out20
	}
	var _253_v83 _dafny.MultiSet
	_ = _253_v83
	_253_v83 = _dafny.MultiSetOf(_161_v8)
	if Companion_Default___.Fm1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_v32, _253_v83)).Cardinality(), _161_v8, _dafny.IntOfInt64(844), (_175_v21).Cmp((_253_v83).Cardinality()) < 0, _159_globalState) {
		_191_v35 = _dafny.IntOfInt64(214)
		_184_v28 = (_184_v28).Update((func() _dafny.Int {
			if _161_v8 {
				return (_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)
			}
			return (_dafny.Zero).Minus((_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int))
		})(), false)
		var _254_v84 _dafny.Map
		_ = _254_v84
		_254_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v8, _187_v31)
		var _255_v85 _dafny.Map
		_ = _255_v85
		_255_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_184_v28).Merge(Companion_Default___.Fm3(_188_v32, _161_v8, (_181_v25).Fm7(_159_globalState), _159_globalState)), (func() _dafny.Array {
			if (_254_v84).Contains(_161_v8) {
				return (_254_v84).Get(_161_v8).(_dafny.Array)
			}
			return _187_v31
		})())
		var _256_v87 _dafny.Map
		_ = _256_v87
		_256_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_dafny.IntOfInt64(46), _188_v32, _dafny.IntOfInt64(959), _161_v8, _159_globalState), (_183_v27).F15())
		var _257_v88 D0
		_ = _257_v88
		_257_v88 = Companion_D0_.Create_DC2_(_161_v8, func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(601), _dafny.IntOfInt64(587))); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _258_v86 _dafny.Int
				_258_v86 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(601)).Cmp(_258_v86) <= 0) && ((_258_v86).Cmp(_dafny.IntOfInt64(587)) < 0) {
					_coll5.Add(Companion_Default___.SafeDivisionInt(_258_v86, _dafny.IntOfUint32((_157_v5).Cardinality())), _188_v32)
				}
			}
			return _coll5.ToMap()
		}(), _256_v87, _161_v8)
		var _259_v89 D0
		_ = _259_v89
		_259_v89 = Companion_D0_.Create_DC2_(_161_v8, (_257_v88).Dtor_cf3(), (_256_v87).Update(_161_v8, (_183_v27).F15()), _161_v8)
		var _260_v90 _dafny.Map
		_ = _260_v90
		_260_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_191_v35, _187_v31)
		_255_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_257_v88).Dtor_cf3(), (func() _dafny.Array {
			if (_260_v90).Contains(_dafny.IntOfInt64(702)) {
				return (_260_v90).Get(_dafny.IntOfInt64(702)).(_dafny.Array)
			}
			return _187_v31
		})())
		_161_v8 = _161_v8
		var _261_v91 _dafny.Set
		_ = _261_v91
		_261_v91 = _dafny.SetOf(_188_v32)
		var _262_v92 _dafny.Sequence
		_ = _262_v92
		_262_v92 = _dafny.SeqOf(_188_v32)
		var _263_v93 _dafny.Sequence
		_ = _263_v93
		_263_v93 = _dafny.SeqOf(_262_v92, _dafny.SeqOf(_161_v8, _188_v32, _188_v32), _262_v92)
		var _264_v94 D8
		_ = _264_v94
		_264_v94 = Companion_D8_.Create_DC24_(_261_v91, (_263_v93).Select((Companion_Default___.SafeIndex(_175_v21, _dafny.IntOfUint32((_263_v93).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _265_v95 _dafny.Set
		_ = _265_v95
		_265_v95 = _dafny.SetOf(Companion_Default___.Fm22(_188_v32, _159_globalState), _264_v94)
		var _266_v96 _dafny.Map
		_ = _266_v96
		_266_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_265_v95).Intersection(_265_v95), (func() _dafny.Int {
			if _188_v32 {
				return _dafny.IntOfUint32((_153_v1).Cardinality())
			}
			return _191_v35
		})())
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_187_v31), 0))
		_ = _index40
		var _rhs38 _dafny.Map = (_266_v96).Merge(_266_v96)
		_ = _rhs38
		var _rhs39 *C3 = _181_v25
		_ = _rhs39
		var _rhs40 bool = true
		_ = _rhs40
		var _lhs23 _dafny.Array = _187_v31
		_ = _lhs23
		var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_187_v31), 0))
		_ = _lhs24
		_266_v96 = _rhs38
		(_lhs23).ArraySet1(_rhs39, (_lhs24).Int())
		_161_v8 = _rhs40
	} else {
		_188_v32 = _161_v8
		_156_v4 = _156_v4
		_156_v4 = (func() _dafny.CodePoint {
			if !(_188_v32) {
				return _dafny.CodePoint('m')
			}
			return (_164_v11).Select((Companion_Default___.SafeIndex(_175_v21, _dafny.IntOfUint32((_164_v11).Cardinality()))).Uint32()).(_dafny.CodePoint)
		})()
		(_159_globalState).F5 = _152_v0
		(_159_globalState).F5 = _dafny.IntOfInt64(578)
	}
	var _267_v97 bool
	_ = _267_v97
	var _268_v98 _dafny.Array
	_ = _268_v98
	var _269_v99 _dafny.Map
	_ = _269_v99
	var _270_v100 _dafny.Int
	_ = _270_v100
	var _out21 bool
	_ = _out21
	var _out22 _dafny.Array
	_ = _out22
	var _out23 _dafny.Map
	_ = _out23
	var _out24 _dafny.Int
	_ = _out24
	_out21, _out22, _out23, _out24 = (_183_v27).M2(_159_globalState)
	_267_v97 = _out21
	_268_v98 = _out22
	_269_v99 = _out23
	_270_v100 = _out24
	var _271_v101 *C3
	_ = _271_v101
	var _nw50 *C3 = New_C3_()
	_ = _nw50
	_nw50.Ctor__((_181_v25).F14(), _156_v4, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()))
	_271_v101 = _nw50
	var _272_i5 _dafny.Int
	_ = _272_i5
	_272_i5 = _dafny.Zero
	{
		for true {
			{
				if (_272_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_272_i5 = (_272_i5).Plus(_dafny.One)
				(_159_globalState).F5 = _175_v21
				if !(_188_v32) || (_267_v97) {
					var _273_v102 bool
					_ = _273_v102
					var _274_v103 _dafny.Array
					_ = _274_v103
					var _275_v104 _dafny.Map
					_ = _275_v104
					var _276_v105 _dafny.Int
					_ = _276_v105
					var _out25 bool
					_ = _out25
					var _out26 _dafny.Array
					_ = _out26
					var _out27 _dafny.Map
					_ = _out27
					var _out28 _dafny.Int
					_ = _out28
					_out25, _out26, _out27, _out28 = (_183_v27).M2(_159_globalState)
					_273_v102 = _out25
					_274_v103 = _out26
					_275_v104 = _out27
					_276_v105 = _out28
					_191_v35 = (func() _dafny.Int {
						if _267_v97 {
							return (_181_v25).Fm4(_159_globalState)
						}
						return (_162_v9).Minus(_162_v9)
					})()
					var _277_v106 _dafny.Array
					_ = _277_v106
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(4)
					_ = _len0_6
					var _nw51 _dafny.Array
					_ = _nw51
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw51 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) bool = (func(_278_v1 _dafny.Sequence) func(_dafny.Int) bool {
							return func(_279_i6 _dafny.Int) bool {
								return _dafny.Companion_Sequence_.IsProperPrefixOf(_278_v1, _278_v1)
							}
						})(_153_v1)
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
					_277_v106 = _nw51
					var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_277_v106), 0))
					_ = _index41
					(_277_v106).ArraySet1(_267_v97, (_index41).Int())
					var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_277_v106), 0))
					_ = _index42
					var _rhs41 _dafny.Array = _dafny.ArrayCastTo((_173_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_173_v19), 0))).Int()))
					_ = _rhs41
					var _rhs42 bool = !(_267_v97)
					_ = _rhs42
					var _lhs25 _dafny.Array = _277_v106
					_ = _lhs25
					var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_277_v106), 0))
					_ = _lhs26
					_274_v103 = _rhs41
					(_lhs25).ArraySet1(_rhs42, (_lhs26).Int())
					var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))
					_ = _index43
					var _rhs43 _dafny.Sequence = _157_v5
					_ = _rhs43
					var _rhs44 _dafny.Int = _191_v35
					_ = _rhs44
					var _rhs45 _dafny.Int = (_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)
					_ = _rhs45
					var _rhs46 bool = true
					_ = _rhs46
					var _lhs27 _dafny.Array = _172_v18
					_ = _lhs27
					var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))
					_ = _lhs28
					_177_v23 = _rhs43
					_175_v21 = _rhs44
					(_lhs27).ArraySet1(_rhs45, (_lhs28).Int())
					_267_v97 = _rhs46
					var _280_v107 *C3
					_ = _280_v107
					var _nw52 *C3 = New_C3_()
					_ = _nw52
					_nw52.Ctor__((_181_v25).F14(), (_183_v27).F15(), (_dafny.Zero).Minus((_276_v105).Minus(_191_v35)))
					_280_v107 = _nw52
				} else {
					_164_v11 = (func() _dafny.Sequence {
						if _161_v8 {
							return _dafny.Companion_Sequence_.Concatenate(_177_v23, _157_v5)
						}
						return _dafny.Companion_Sequence_.Concatenate(_157_v5, _177_v23)
					})()
					var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))
					_ = _index44
					(_172_v18).ArraySet1(_191_v35, (_index44).Int())
					var _281_v108 _dafny.Set
					_ = _281_v108
					_281_v108 = _dafny.SetOf(_156_v4)
					var _rhs47 _dafny.Set = ((_dafny.SetOf((_183_v27).F15(), _156_v4)).Union(_dafny.SetOf(_156_v4))).Union(_281_v108)
					_ = _rhs47
					var _rhs48 bool = ((_270_v100).Minus(_152_v0)).Cmp((_153_v1).Select((Companion_Default___.SafeIndex(_162_v9, _dafny.IntOfUint32((_153_v1).Cardinality()))).Uint32()).(_dafny.Int)) != 0
					_ = _rhs48
					var _rhs49 _dafny.Int = _162_v9
					_ = _rhs49
					var _rhs50 _dafny.Sequence = _153_v1
					_ = _rhs50
					var _lhs29 *GlobalState = _159_globalState
					_ = _lhs29
					_281_v108 = _rhs47
					_267_v97 = _rhs48
					_lhs29.F5 = _rhs49
					_153_v1 = _rhs50
					var _282_v109 _dafny.Sequence
					_ = _282_v109
					_282_v109 = _dafny.Companion_Sequence_.Concatenate(_157_v5, _157_v5)
					_282_v109 = _282_v109
					var _283_v110 bool
					_ = _283_v110
					var _284_v111 _dafny.Array
					_ = _284_v111
					var _285_v112 _dafny.Map
					_ = _285_v112
					var _286_v113 _dafny.Int
					_ = _286_v113
					var _out29 bool
					_ = _out29
					var _out30 _dafny.Array
					_ = _out30
					var _out31 _dafny.Map
					_ = _out31
					var _out32 _dafny.Int
					_ = _out32
					_out29, _out30, _out31, _out32 = (_183_v27).M2(_159_globalState)
					_283_v110 = _out29
					_284_v111 = _out30
					_285_v112 = _out31
					_286_v113 = _out32
				}
				if _161_v8 {
					var _287_v114 D6
					_ = _287_v114
					_287_v114 = Companion_D6_.Create_DC15_(_dafny.Companion_Sequence_.Update(_153_v1, (Companion_Default___.SafeIndex(_191_v35, _dafny.IntOfUint32((_153_v1).Cardinality()))).Uint32(), _175_v21))
					_287_v114 = (func() D6 {
						if !(_161_v8) {
							return _287_v114
						}
						return _287_v114
					})()
					_153_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_174_v20).Cardinality()), _153_v1)
					_153_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_270_v100, (_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int), _162_v9, _152_v0), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_152_v0))))
					var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))
					_ = _index45
					(_172_v18).ArraySet1(_dafny.IntOfInt64(399), (_index45).Int())
					_152_v0 = _175_v21
				} else {
					var _288_v115 *C4
					_ = _288_v115
					var _nw53 *C4 = New_C4_()
					_ = _nw53
					_nw53.Ctor__((_162_v9).Plus(_191_v35), Companion_Default___.SafeModuloInt((_176_v22).Cardinality(), _175_v21))
					_288_v115 = _nw53
					var _289_v116 bool
					_ = _289_v116
					var _290_v117 bool
					_ = _290_v117
					var _291_v118 _dafny.Int
					_ = _291_v118
					var _out33 bool
					_ = _out33
					var _out34 bool
					_ = _out34
					var _out35 _dafny.Int
					_ = _out35
					_out33, _out34, _out35 = (_271_v101).M4(_162_v9, _161_v8, _dafny.Companion_Sequence_.Concatenate(_157_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-669))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg21 _dafny.Int) interface{} {
							return coer21(arg21)
						}
					}((func(_292_v27 *C1) func(_dafny.Int) _dafny.CodePoint {
						return func(_293_i7 _dafny.Int) _dafny.CodePoint {
							return (_292_v27).F15()
						}
					})(_183_v27)))), _159_globalState)
					_289_v116 = _out33
					_290_v117 = _out34
					_291_v118 = _out35
					var _294_v119 _dafny.Array
					_ = _294_v119
					var _nw54 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
					_ = _nw54
					_294_v119 = _nw54
					var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_294_v119), 0))
					_ = _index46
					(_294_v119).ArraySet1(((_174_v20).Cardinality()).Cmp((_288_v115).F10()) != 0, (_index46).Int())
					var _295_v120 _dafny.Map
					_ = _295_v120
					_295_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v0, _dafny.SetOf(_175_v21))
					var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_294_v119), 0))
					_ = _index47
					(_294_v119).ArraySet1((_174_v20).IsSubsetOf((func() _dafny.Set {
						if (_295_v120).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_177_v23).Cardinality()))) {
							return (_295_v120).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_177_v23).Cardinality()))).(_dafny.Set)
						}
						return _174_v20
					})()), (_index47).Int())
					var _296_v121 _dafny.Int
					_ = _296_v121
					var _297_v122 bool
					_ = _297_v122
					var _298_v123 _dafny.Int
					_ = _298_v123
					var _out36 _dafny.Int
					_ = _out36
					var _out37 bool
					_ = _out37
					var _out38 _dafny.Int
					_ = _out38
					_out36, _out37, _out38 = (_271_v101).M3(_253_v83, _159_globalState)
					_296_v121 = _out36
					_297_v122 = _out37
					_298_v123 = _out38
					var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_294_v119), 0))
					_ = _index48
					(_294_v119).ArraySet1(false, (_index48).Int())
				}
				_152_v0 = (_172_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_172_v18), 0))).Int()).(_dafny.Int)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _rhs51 _dafny.Int = (_dafny.Zero).Minus(_162_v9)
	_ = _rhs51
	var _rhs52 bool = Companion_Default___.Fm1((_163_v10).Cardinality(), _267_v97, _270_v100, _267_v97, _159_globalState)
	_ = _rhs52
	_191_v35 = _rhs51
	_188_v32 = _rhs52
	_dafny.Print(_152_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_153_v1, _dafny.SeqOf(_dafny.IntOfInt64(6), _dafny.IntOfInt64(399), _dafny.IntOfInt64(583), _dafny.IntOfInt64(399), _dafny.IntOfInt64(399))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(172), _dafny.SeqOf(_dafny.IntOfInt64(172)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v3).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(172))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_156_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_157_v5, _dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v6).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F0()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(138))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_159_globalState).F9()).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v7).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(172), _dafny.IntOfInt64(447))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_162_v9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(821))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_164_v11, _dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v18).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v19).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v20).Equals(_dafny.SetOf(_dafny.IntOfInt64(-144))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_v21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_176_v22).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(821))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_177_v23, _dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v24).ArrayGet1((_dafny.Zero).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v24).ArrayGet1((_dafny.One).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v24).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v24).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v24).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_181_v25).F14()).ArrayGet1((_dafny.Zero).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_181_v25).F14()).ArrayGet1((_dafny.One).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_181_v25).F14()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_181_v25).F14()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_181_v25).F14()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_v26).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C1)).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_v26).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C1)).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v26).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C1).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_183_v27).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v28).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(171), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_v29).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(171), true), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_186_v30).Dtor_cf33()).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_186_v30).Dtor_cf33()).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v30).Dtor_cf33().F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v30).Dtor_cf34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v30).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v30).Dtor_cf36())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_186_v30).Dtor_cf37()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(171), true), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_187_v31).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C3)).F14()).ArrayGet1((_dafny.Zero).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_187_v31).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C3)).F14()).ArrayGet1((_dafny.One).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_187_v31).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C3)).F14()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_187_v31).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C3)).F14()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_187_v31).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C3)).F14()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_187_v31).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C3)).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v31).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C3).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_188_v32)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v33).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v34).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_191_v35)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v36).Dtor_cf20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v83).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_267_v97)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v98).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v99).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_270_v100)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_271_v101).F14()).ArrayGet1((_dafny.Zero).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_271_v101).F14()).ArrayGet1((_dafny.One).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_271_v101).F14()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_271_v101).F14()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_271_v101).F14()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D2)).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_272_i5)
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
	Cf1 _dafny.Array
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Array) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 bool
	Cf3 _dafny.Map
	Cf4 _dafny.Map
	Cf5 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 bool, Cf3 _dafny.Map, Cf4 _dafny.Map, Cf5 bool) D0 {
	return D0{D0_DC2{Cf2, Cf3, Cf4, Cf5}}
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
	Cf6 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf6 D0) D0 {
	return D0{D0_DC3{Cf6}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Map {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Map {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() bool {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf6() D0 {
	return _this.Get_().(D0_DC3).Cf6
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf6) + ")"
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
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3.Equals(data2.Cf3) && data1.Cf4.Equals(data2.Cf4) && data1.Cf5 == data2.Cf5
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	Cf8 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf8 bool) D1 {
	return D1{D1_DC5{Cf8}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf7 _dafny.Array
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 _dafny.Array) D1 {
	return D1{D1_DC4{Cf7}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC6 struct {
	Cf9 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf9 D1) D1 {
	return D1{D1_DC6{Cf9}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false)
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf9() D1 {
	return _this.Get_().(D1_DC6).Cf9
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf8 == data2.Cf8
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf7 == data2.Cf7
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf9.Equals(data2.Cf9)
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
	Cf11 bool
	Cf12 _dafny.Int
	Cf13 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf11 bool, Cf12 _dafny.Int, Cf13 bool) D2 {
	return D2{D2_DC8{Cf11, Cf12, Cf13}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC9 struct {
	Cf14 _dafny.Int
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf14 _dafny.Int) D2 {
	return D2{D2_DC9{Cf14}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC7 struct {
	Cf10 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf10 _dafny.Int) D2 {
	return D2{D2_DC7{Cf10}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(false, _dafny.Zero, false)
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC8).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf14
}

func (_this D2) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13 == data2.Cf13
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0
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
	Cf15 _dafny.Sequence
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf15 _dafny.Sequence) D3 {
	return D3{D3_DC10{Cf15}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D3) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf15
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + data.Cf15.VerbatimString(true) + ")"
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
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf16 _dafny.Set
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf16 _dafny.Set) D4 {
	return D4{D4_DC11{Cf16}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D4) Dtor_cf16() _dafny.Set {
	return _this.Get_().(D4_DC11).Cf16
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf16.Equals(data2.Cf16)
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

type D5_DC13 struct {
	Cf18 bool
	Cf19 bool
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf18 bool, Cf19 bool) D5 {
	return D5{D5_DC13{Cf18, Cf19}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC14 struct {
	Cf20 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf20 bool) D5 {
	return D5{D5_DC14{Cf20}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC12 struct {
	Cf17 _dafny.Sequence
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf17 _dafny.Sequence) D5 {
	return D5{D5_DC12{Cf17}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(false, false)
}

func (_this D5) Dtor_cf18() bool {
	return _this.Get_().(D5_DC13).Cf18
}

func (_this D5) Dtor_cf19() bool {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D5_DC12).Cf17
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf17) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf20 == data2.Cf20
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf17.Equals(data2.Cf17)
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
	Cf22 _dafny.Array
	Cf23 _dafny.Int
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf22 _dafny.Array, Cf23 _dafny.Int) D6 {
	return D6{D6_DC16{Cf22, Cf23}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC17 struct {
	Cf24 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf24 bool) D6 {
	return D6{D6_DC17{Cf24}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf25 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf25 _dafny.Int) D6 {
	return D6{D6_DC18{Cf25}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC15 struct {
	Cf21 _dafny.Sequence
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf21 _dafny.Sequence) D6 {
	return D6{D6_DC15{Cf21}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC19 struct {
	Cf26 D6
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf26 D6) D6 {
	return D6{D6_DC19{Cf26}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D6) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf23
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC17).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf25
}

func (_this D6) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D6_DC15).Cf21
}

func (_this D6) Dtor_cf26() D6 {
	return _this.Get_().(D6_DC19).Cf26
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf26) + ")"
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
			return ok && data1.Cf22 == data2.Cf22 && data1.Cf23.Cmp(data2.Cf23) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf21.Equals(data2.Cf21)
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D7_DC21 struct {
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_() D7 {
	return D7{D7_DC21{}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC20 struct {
	Cf27 _dafny.Map
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf27 _dafny.Map) D7 {
	return D7{D7_DC20{Cf27}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC21_()
}

func (_this D7) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D7_DC20).Cf27
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC21:
		{
			_, ok := other.Get_().(D7_DC21)
			return ok
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf27.Equals(data2.Cf27)
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

type D8_DC23 struct {
	Cf29 _dafny.Int
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf29 _dafny.Int) D8 {
	return D8{D8_DC23{Cf29}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf30 _dafny.Set
	Cf31 _dafny.Sequence
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf30 _dafny.Set, Cf31 _dafny.Sequence) D8 {
	return D8{D8_DC24{Cf30, Cf31}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC22 struct {
	Cf28 _dafny.MultiSet
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf28 _dafny.MultiSet) D8 {
	return D8{D8_DC22{Cf28}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(_dafny.Zero)
}

func (_this D8) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf29
}

func (_this D8) Dtor_cf30() _dafny.Set {
	return _this.Get_().(D8_DC24).Cf30
}

func (_this D8) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D8_DC24).Cf31
}

func (_this D8) Dtor_cf28() _dafny.MultiSet {
	return _this.Get_().(D8_DC22).Cf28
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf30.Equals(data2.Cf30) && data1.Cf31.Equals(data2.Cf31)
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf28.Equals(data2.Cf28)
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

type D9_DC26 struct {
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_() D9 {
	return D9{D9_DC26{}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC27 struct {
	Cf33 *C1
	Cf34 _dafny.Int
	Cf35 _dafny.Int
	Cf36 _dafny.Int
	Cf37 _dafny.Map
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf33 *C1, Cf34 _dafny.Int, Cf35 _dafny.Int, Cf36 _dafny.Int, Cf37 _dafny.Map) D9 {
	return D9{D9_DC27{Cf33, Cf34, Cf35, Cf36, Cf37}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC25 struct {
	Cf32 _dafny.CodePoint
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf32 _dafny.CodePoint) D9 {
	return D9{D9_DC25{Cf32}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC28 struct {
	Cf38 D9
}

func (D9_DC28) isD9() {}

func (CompanionStruct_D9_) Create_DC28_(Cf38 D9) D9 {
	return D9{D9_DC28{Cf38}}
}

func (_this D9) Is_DC28() bool {
	_, ok := _this.Get_().(D9_DC28)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC26_()
}

func (_this D9) Dtor_cf33() *C1 {
	return _this.Get_().(D9_DC27).Cf33
}

func (_this D9) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D9_DC27).Cf34
}

func (_this D9) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D9_DC27).Cf35
}

func (_this D9) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D9_DC27).Cf36
}

func (_this D9) Dtor_cf37() _dafny.Map {
	return _this.Get_().(D9_DC27).Cf37
}

func (_this D9) Dtor_cf32() _dafny.CodePoint {
	return _this.Get_().(D9_DC25).Cf32
}

func (_this D9) Dtor_cf38() D9 {
	return _this.Get_().(D9_DC28).Cf38
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D9_DC28:
		{
			return "D9.DC28" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC26:
		{
			_, ok := other.Get_().(D9_DC26)
			return ok
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf33 == data2.Cf33 && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37.Equals(data2.Cf37)
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf32 == data2.Cf32
		}
	case D9_DC28:
		{
			data2, ok := other.Get_().(D9_DC28)
			return ok && data1.Cf38.Equals(data2.Cf38)
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
	Cf40 _dafny.Int
}

func (D10_DC30) isD10() {}

func (CompanionStruct_D10_) Create_DC30_(Cf40 _dafny.Int) D10 {
	return D10{D10_DC30{Cf40}}
}

func (_this D10) Is_DC30() bool {
	_, ok := _this.Get_().(D10_DC30)
	return ok
}

type D10_DC29 struct {
	Cf39 T2
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf39 T2) D10 {
	return D10{D10_DC29{Cf39}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC31 struct {
	Cf41 D10
}

func (D10_DC31) isD10() {}

func (CompanionStruct_D10_) Create_DC31_(Cf41 D10) D10 {
	return D10{D10_DC31{Cf41}}
}

func (_this D10) Is_DC31() bool {
	_, ok := _this.Get_().(D10_DC31)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC30_(_dafny.Zero)
}

func (_this D10) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D10_DC30).Cf40
}

func (_this D10) Dtor_cf39() T2 {
	return _this.Get_().(D10_DC29).Cf39
}

func (_this D10) Dtor_cf41() D10 {
	return _this.Get_().(D10_DC31).Cf41
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC30:
		{
			return "D10.DC30" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D10_DC31:
		{
			return "D10.DC31" + "(" + _dafny.String(data.Cf41) + ")"
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
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0
		}
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && _dafny.AreEqual(data1.Cf39, data2.Cf39)
		}
	case D10_DC31:
		{
			data2, ok := other.Get_().(D10_DC31)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D11_DC33 struct {
	Cf43 _dafny.Map
	Cf44 _dafny.Int
	Cf45 bool
	Cf46 bool
	Cf47 _dafny.Int
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf43 _dafny.Map, Cf44 _dafny.Int, Cf45 bool, Cf46 bool, Cf47 _dafny.Int) D11 {
	return D11{D11_DC33{Cf43, Cf44, Cf45, Cf46, Cf47}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

type D11_DC34 struct {
	Cf48 D0
	Cf49 _dafny.Int
	Cf50 _dafny.Map
}

func (D11_DC34) isD11() {}

func (CompanionStruct_D11_) Create_DC34_(Cf48 D0, Cf49 _dafny.Int, Cf50 _dafny.Map) D11 {
	return D11{D11_DC34{Cf48, Cf49, Cf50}}
}

func (_this D11) Is_DC34() bool {
	_, ok := _this.Get_().(D11_DC34)
	return ok
}

type D11_DC32 struct {
	Cf42 _dafny.Map
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf42 _dafny.Map) D11 {
	return D11{D11_DC32{Cf42}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

type D11_DC35 struct {
	Cf51 D11
}

func (D11_DC35) isD11() {}

func (CompanionStruct_D11_) Create_DC35_(Cf51 D11) D11 {
	return D11{D11_DC35{Cf51}}
}

func (_this D11) Is_DC35() bool {
	_, ok := _this.Get_().(D11_DC35)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC33_(_dafny.EmptyMap, _dafny.Zero, false, false, _dafny.Zero)
}

func (_this D11) Dtor_cf43() _dafny.Map {
	return _this.Get_().(D11_DC33).Cf43
}

func (_this D11) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D11_DC33).Cf44
}

func (_this D11) Dtor_cf45() bool {
	return _this.Get_().(D11_DC33).Cf45
}

func (_this D11) Dtor_cf46() bool {
	return _this.Get_().(D11_DC33).Cf46
}

func (_this D11) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D11_DC33).Cf47
}

func (_this D11) Dtor_cf48() D0 {
	return _this.Get_().(D11_DC34).Cf48
}

func (_this D11) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D11_DC34).Cf49
}

func (_this D11) Dtor_cf50() _dafny.Map {
	return _this.Get_().(D11_DC34).Cf50
}

func (_this D11) Dtor_cf42() _dafny.Map {
	return _this.Get_().(D11_DC32).Cf42
}

func (_this D11) Dtor_cf51() D11 {
	return _this.Get_().(D11_DC35).Cf51
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D11_DC34:
		{
			return "D11.DC34" + "(" + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D11_DC35:
		{
			return "D11.DC35" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf43.Equals(data2.Cf43) && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45 == data2.Cf45 && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0
		}
	case D11_DC34:
		{
			data2, ok := other.Get_().(D11_DC34)
			return ok && data1.Cf48.Equals(data2.Cf48) && data1.Cf49.Cmp(data2.Cf49) == 0 && data1.Cf50.Equals(data2.Cf50)
		}
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D11_DC35:
		{
			data2, ok := other.Get_().(D11_DC35)
			return ok && data1.Cf51.Equals(data2.Cf51)
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

type D12_DC37 struct {
}

func (D12_DC37) isD12() {}

func (CompanionStruct_D12_) Create_DC37_() D12 {
	return D12{D12_DC37{}}
}

func (_this D12) Is_DC37() bool {
	_, ok := _this.Get_().(D12_DC37)
	return ok
}

type D12_DC38 struct {
	Cf53 D8
	Cf54 _dafny.Int
	Cf55 _dafny.Int
}

func (D12_DC38) isD12() {}

func (CompanionStruct_D12_) Create_DC38_(Cf53 D8, Cf54 _dafny.Int, Cf55 _dafny.Int) D12 {
	return D12{D12_DC38{Cf53, Cf54, Cf55}}
}

func (_this D12) Is_DC38() bool {
	_, ok := _this.Get_().(D12_DC38)
	return ok
}

type D12_DC39 struct {
	Cf56 bool
}

func (D12_DC39) isD12() {}

func (CompanionStruct_D12_) Create_DC39_(Cf56 bool) D12 {
	return D12{D12_DC39{Cf56}}
}

func (_this D12) Is_DC39() bool {
	_, ok := _this.Get_().(D12_DC39)
	return ok
}

type D12_DC36 struct {
	Cf52 _dafny.Set
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf52 _dafny.Set) D12 {
	return D12{D12_DC36{Cf52}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC37_()
}

func (_this D12) Dtor_cf53() D8 {
	return _this.Get_().(D12_DC38).Cf53
}

func (_this D12) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D12_DC38).Cf54
}

func (_this D12) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D12_DC38).Cf55
}

func (_this D12) Dtor_cf56() bool {
	return _this.Get_().(D12_DC39).Cf56
}

func (_this D12) Dtor_cf52() _dafny.Set {
	return _this.Get_().(D12_DC36).Cf52
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC37:
		{
			return "D12.DC37"
		}
	case D12_DC38:
		{
			return "D12.DC38" + "(" + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ")"
		}
	case D12_DC39:
		{
			return "D12.DC39" + "(" + _dafny.String(data.Cf56) + ")"
		}
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf52) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC37:
		{
			_, ok := other.Get_().(D12_DC37)
			return ok
		}
	case D12_DC38:
		{
			data2, ok := other.Get_().(D12_DC38)
			return ok && data1.Cf53.Equals(data2.Cf53) && data1.Cf54.Cmp(data2.Cf54) == 0 && data1.Cf55.Cmp(data2.Cf55) == 0
		}
	case D12_DC39:
		{
			data2, ok := other.Get_().(D12_DC39)
			return ok && data1.Cf56 == data2.Cf56
		}
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf52.Equals(data2.Cf52)
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

// Definition of trait T0
type T0 interface {
	String() string
	F12() _dafny.Int
	F12_set_(value _dafny.Int)
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
	F12() _dafny.Int
	F12_set_(value _dafny.Int)
	Fm4(globalState *GlobalState) _dafny.Int
	Fm5(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence
	M2(globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int)
	F13() _dafny.CodePoint
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
	F12() _dafny.Int
	F12_set_(value _dafny.Int)
	Fm6(globalState *GlobalState) D0
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
	F5  _dafny.Int
	_f0 _dafny.MultiSet
	_f1 bool
	_f2 bool
	_f3 bool
	_f4 _dafny.Int
	_f6 bool
	_f7 bool
	_f8 _dafny.Int
	_f9 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F5 = _dafny.Zero
	_this._f0 = _dafny.EmptyMultiSet
	_this._f1 = false
	_this._f2 = false
	_this._f3 = false
	_this._f4 = _dafny.Zero
	_this._f6 = false
	_this._f7 = false
	_this._f8 = _dafny.Zero
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.MultiSet, f1 bool, f2 bool, f3 bool, f4 _dafny.Int, f5 _dafny.Int, f6 bool, f7 bool, f8 _dafny.Int, f9 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *GlobalState) F0() _dafny.MultiSet {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
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
func (_this *GlobalState) F9() _dafny.Array {
	{
		return _this._f9
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f12 _dafny.Int
	F16  _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f12 = _dafny.Zero
	_this.F16 = _dafny.EmptyMap
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C0{}
var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F12() _dafny.Int {
	return _this._f12
}
func (_this *C0) F12_set_(value _dafny.Int) {
	_this._f12 = value
}
func (_this *C0) Ctor__(f16 _dafny.Map, f12 _dafny.Int) {
	{
		(_this).F16 = f16
		(_this)._f12 = f12
	}
}
func (_this *C0) Fm6(globalState *GlobalState) D0 {
	{
		return Companion_D0_.Create_DC0_(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("qxlwa"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(370))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_299_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		}))))
	}
}
func (_this *C0) Fm8(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lnadqpvyl"), (func() _dafny.Sequence {
			if false {
				return _dafny.UnicodeSeqOfUtf8Bytes("yutm")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("slatllus")
		})())
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f12 _dafny.Int
	_f13 _dafny.CodePoint
	_f15 _dafny.CodePoint
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.CodePoint('D')
	_this._f15 = _dafny.CodePoint('D')
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

func (_this *C1) F12() _dafny.Int {
	return _this._f12
}
func (_this *C1) F12_set_(value _dafny.Int) {
	_this._f12 = value
}
func (_this *C1) F13() _dafny.CodePoint {
	return _this._f13
}
func (_this *C1) Ctor__(f15 _dafny.CodePoint, f13 _dafny.CodePoint, f12 _dafny.Int) {
	{
		(_this)._f15 = f15
		(_this)._f13 = f13
		(_this)._f12 = f12
	}
}
func (_this *C1) Fm4(globalState *GlobalState) _dafny.Int {
	{
		return _this.F12()
	}
}
func (_this *C1) Fm5(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		if false {
			return _dafny.SeqOf(_this.F12())
		} else if true {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(155))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_300_i0 _dafny.Int) _dafny.Int {
				return _this.F12()
			}))
		} else {
			return _dafny.SeqOf(_dafny.IntOfInt64(234), _this.F12())
		}
	}
}
func (_this *C1) M2(globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _301_v0 _dafny.Array
		_ = _301_v0
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_7
		var _nw55 _dafny.Array
		_ = _nw55
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw55 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Int = func(_302_i0 _dafny.Int) _dafny.Int {
				return (_302_i0).Times(_this.F12())
			}
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw55 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw55).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw55).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_301_v0 = _nw55
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))
		_ = _index49
		(_301_v0).ArraySet1(_this.F12(), (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))
		_ = _index50
		(_301_v0).ArraySet1(_this.F12(), (_index50).Int())
		var _303_v1 bool
		_ = _303_v1
		_303_v1 = false
		var _304_v2 D1
		_ = _304_v2
		_304_v2 = Companion_D1_.Create_DC5_(_303_v1)
		_304_v2 = Companion_D1_.Create_DC5_(_303_v1)
		if (func() bool {
			if _303_v1 {
				return _303_v1
			}
			return _303_v1
		})() {
			var _305_v3 _dafny.Set
			_ = _305_v3
			_305_v3 = _dafny.SetOf(_303_v1, _303_v1)
			var _306_v4 _dafny.Map
			_ = _306_v4
			_306_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_305_v3).Cardinality(), false)
			var _307_v5 T2
			_ = _307_v5
			var _nw56 *C0 = New_C0_()
			_ = _nw56
			_nw56.Ctor__(_306_v4, (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
			_307_v5 = _nw56
			var _308_v6 _dafny.Sequence
			_ = _308_v6
			_308_v6 = _dafny.SeqOf(_303_v1, _303_v1, _303_v1)
			var _309_v7 _dafny.Sequence
			_ = _309_v7
			_309_v7 = _dafny.SeqOf(_301_v0)
			var _rhs53 bool = true
			_ = _rhs53
			var _rhs54 bool = !(_303_v1) || (_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_307_v5), _307_v5))
			_ = _rhs54
			var _rhs55 bool = !(_303_v1) || (!((_308_v6).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_307_v5.F12()), _dafny.IntOfUint32((_308_v6).Cardinality()))).Uint32()).(bool)))
			_ = _rhs55
			var _rhs56 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_309_v7, _309_v7), _309_v7), (Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_309_v7, _309_v7), _309_v7)).Cardinality()))).Uint32(), _301_v0)).Cardinality())
			_ = _rhs56
			var _lhs30 *GlobalState = globalState
			_ = _lhs30
			r0 = _rhs53
			r0 = _rhs54
			r0 = _rhs55
			_lhs30.F5 = _rhs56
			var _310_v8 _dafny.Sequence
			_ = _310_v8
			_310_v8 = _dafny.SeqOf(_this.F12())
			var _311_v9 _dafny.MultiSet
			_ = _311_v9
			_311_v9 = _dafny.MultiSetOf(_307_v5.F12())
			var _312_v10 _dafny.Map
			_ = _312_v10
			_312_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v1, _311_v9)
			var _313_v11 _dafny.Map
			_ = _313_v11
			_313_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_312_v10, _303_v1)
			var _314_v12 _dafny.Map
			_ = _314_v12
			_314_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v1, (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
			var _315_v13 D2
			_ = _315_v13
			_315_v13 = Companion_D2_.Create_DC7_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_311_v9).Contains(_this.F12()) {
					return (_311_v9).Multiplicity(_this.F12())
				}
				return (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)
			})(), (func() _dafny.Int {
				if (_314_v12).Contains(_303_v1) {
					return (_314_v12).Get(_303_v1).(_dafny.Int)
				}
				return _this.F12()
			})())).Cardinality())
			var _rhs57 bool = _303_v1
			_ = _rhs57
			var _rhs58 _dafny.Int = _307_v5.F12()
			_ = _rhs58
			var _rhs59 bool = (func() bool {
				if (_313_v11).Contains(_312_v10) {
					return (_313_v11).Get(_312_v10).(bool)
				}
				return !((_303_v1) || (_303_v1))
			})()
			_ = _rhs59
			var _rhs60 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_310_v8, (Companion_Default___.SafeIndex((func() _dafny.Int {
				if _303_v1 {
					return (_315_v13).Dtor_cf10()
				}
				return _dafny.IntOfInt64(-345)
			})(), _dafny.IntOfUint32((_310_v8).Cardinality()))).Uint32(), _307_v5.F12())
			_ = _rhs60
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			_303_v1 = _rhs57
			_lhs31.F5 = _rhs58
			r0 = _rhs59
			_310_v8 = _rhs60
			if false {
				var _rhs61 bool = (func() bool {
					if _303_v1 {
						return _303_v1
					}
					return !_dafny.Companion_Sequence_.Contains(_308_v6, _303_v1)
				})()
				_ = _rhs61
				var _rhs62 bool = _303_v1
				_ = _rhs62
				_303_v1 = _rhs61
				r0 = _rhs62
				var _316_v14 *C0
				_ = _316_v14
				var _nw57 *C0 = New_C0_()
				_ = _nw57
				_nw57.Ctor__(_306_v4, (Companion_D2_.Create_DC9_(_dafny.IntOfInt64(-834))).Dtor_cf14())
				_316_v14 = _nw57
				var _317_v15 _dafny.CodePoint
				_ = _317_v15
				_317_v15 = _dafny.CodePoint('n')
				_317_v15 = (_this).F13()
				var _318_v16 _dafny.Sequence
				_ = _318_v16
				_318_v16 = _dafny.UnicodeSeqOfUtf8Bytes("iiotodfiq")
				_318_v16 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_318_v16, _318_v16), _318_v16)
				_308_v6 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm9(!(_303_v1), globalState), _dafny.SeqOf(_303_v1, true, _303_v1))
			} else {
				var _319_v17 _dafny.Sequence
				_ = _319_v17
				_319_v17 = _dafny.UnicodeSeqOfUtf8Bytes("q")
				var _320_v18 _dafny.Sequence
				_ = _320_v18
				_320_v18 = _dafny.SeqOf(_319_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(718))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}(func(_321_i1 _dafny.Int) _dafny.CodePoint {
					return (_this).F15()
				})), _dafny.Companion_Sequence_.Concatenate(_319_v17, _319_v17))
				var _322_v19 _dafny.Map
				_ = _322_v19
				_322_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_307_v5.F12()), _dafny.IntOfInt64(210))
				_319_v17 = (_320_v18).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(440)).Plus((func() _dafny.Int {
					if (_322_v19).Contains(_dafny.IntOfUint32((_319_v17).Cardinality())) {
						return (_322_v19).Get(_dafny.IntOfUint32((_319_v17).Cardinality())).(_dafny.Int)
					}
					return (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)
				})()), _dafny.IntOfUint32((_320_v18).Cardinality()))).Uint32()).(_dafny.Sequence)
				var _323_v20 *C0
				_ = _323_v20
				var _nw58 *C0 = New_C0_()
				_ = _nw58
				_nw58.Ctor__(_306_v4, (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
				_323_v20 = _nw58
				_323_v20 = _323_v20
				var _324_v21 _dafny.Array
				_ = _324_v21
				var _nw59 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(22))
				_ = _nw59
				_324_v21 = _nw59
				_324_v21 = _324_v21
				var _325_v22 _dafny.MultiSet
				_ = _325_v22
				_325_v22 = _dafny.MultiSetOf(_308_v6)
				var _326_v23 _dafny.Sequence
				_ = _326_v23
				_326_v23 = _dafny.SeqOf(_308_v6)
				_325_v22 = _dafny.MultiSetFromSeq(_326_v23)
				(_323_v20).F16 = _323_v20.F16
			}
			_303_v1 = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("h"), _dafny.UnicodeSeqOfUtf8Bytes("lbyqqyufh"))
			var _327_v24 *C0
			_ = _327_v24
			var _nw60 *C0 = New_C0_()
			_ = _nw60
			_nw60.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _303_v1), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
			_327_v24 = _nw60
			var _328_v25 _dafny.MultiSet
			_ = _328_v25
			_328_v25 = _dafny.MultiSetOf(_327_v24, _327_v24, _327_v24)
			var _329_v26 _dafny.Map
			_ = _329_v26
			_329_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v1, (_this).F15())
			var _330_v27 D0
			_ = _330_v27
			_330_v27 = Companion_D0_.Create_DC2_(_303_v1, Companion_Default___.Fm10((_328_v25).Cardinality(), globalState), _329_v26, true)
			var _331_v28 _dafny.Map
			_ = _331_v28
			_331_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v27, _this.F12())
			(globalState).F5 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if true {
					return (func() _dafny.Int {
						if (_331_v28).Contains(_330_v27) {
							return (_331_v28).Get(_330_v27).(_dafny.Int)
						}
						return (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)
					})()
				}
				return _307_v5.F12()
			})(), (_307_v5.F12()).Minus((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)))
		} else {
			if (func() bool {
				if _303_v1 {
					return _303_v1
				}
				return false
			})() {
				var _332_v29 _dafny.Map
				_ = _332_v29
				_332_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), _303_v1)
				var _333_v30 _dafny.Sequence
				_ = _333_v30
				_333_v30 = _dafny.UnicodeSeqOfUtf8Bytes("ynbd")
				_303_v1 = !((func() bool {
					if (_332_v29).Contains(_dafny.CodePoint('w')) {
						return (_332_v29).Get(_dafny.CodePoint('w')).(bool)
					}
					return _303_v1
				})()) || (Companion_Default___.Fm1(_dafny.IntOfUint32((_333_v30).Cardinality()), Companion_Default___.Fm1(_this.F12(), !(_303_v1), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _303_v1, globalState), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _303_v1, globalState))
				var _334_v31 _dafny.Map
				_ = _334_v31
				_334_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-503), _303_v1)
				var _335_v32 *C0
				_ = _335_v32
				var _nw61 *C0 = New_C0_()
				_ = _nw61
				_nw61.Ctor__((_334_v31).Merge(_334_v31), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
				_335_v32 = _nw61
				_333_v30 = _333_v30
				r0 = _303_v1
				r0 = Companion_Default___.Fm1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_333_v30, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(198))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}(func(_336_i2 _dafny.Int) _dafny.CodePoint {
					return (_this).F13()
				})), (Companion_Default___.SafeIndex((_dafny.MultiSetOf(_this.F12())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(198))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}(func(_337_i2 _dafny.Int) _dafny.CodePoint {
					return (_this).F13()
				}))).Cardinality()))).Uint32(), _dafny.CodePoint('q')))).Cardinality()), (_this.F12()).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()))) < 0, (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), (_303_v1) && (_303_v1), globalState)
			} else {
				var _338_v33 _dafny.Sequence
				_ = _338_v33
				_338_v33 = _dafny.SeqOf(_303_v1, _303_v1, _303_v1)
				var _339_v34 _dafny.Sequence
				_ = _339_v34
				_339_v34 = _dafny.UnicodeSeqOfUtf8Bytes("cqtlrd")
				var _340_v35 _dafny.Sequence
				_ = _340_v35
				_340_v35 = _dafny.SeqOf(_339_v34)
				var _341_v36 D0
				_ = _341_v36
				_341_v36 = Companion_D0_.Create_DC0_(true)
				var _342_v37 D0
				_ = _342_v37
				_342_v37 = Companion_D0_.Create_DC3_(_341_v36)
				var _343_v38 _dafny.Array
				_ = _343_v38
				var _nwElement0_12 bool = _303_v1
				_ = _nwElement0_12
				var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(23))
				_ = _nw62
				(_nw62).ArraySet1(_nwElement0_12, 0)
				(_nw62).ArraySet1(_303_v1, 1)
				(_nw62).ArraySet1(false, 2)
				(_nw62).ArraySet1(_303_v1, 3)
				(_nw62).ArraySet1(Companion_Default___.Fm1(_this.F12(), _303_v1, _dafny.IntOfInt64(-989), _303_v1, globalState), 4)
				(_nw62).ArraySet1(_303_v1, 5)
				(_nw62).ArraySet1(_303_v1, 6)
				(_nw62).ArraySet1(_303_v1, 7)
				(_nw62).ArraySet1(_303_v1, 8)
				(_nw62).ArraySet1(_303_v1, 9)
				(_nw62).ArraySet1(_303_v1, 10)
				(_nw62).ArraySet1(_303_v1, 11)
				(_nw62).ArraySet1(false, 12)
				(_nw62).ArraySet1(_303_v1, 13)
				(_nw62).ArraySet1(_303_v1, 14)
				(_nw62).ArraySet1(_303_v1, 15)
				(_nw62).ArraySet1(false, 16)
				(_nw62).ArraySet1(_303_v1, 17)
				(_nw62).ArraySet1(_303_v1, 18)
				(_nw62).ArraySet1(_303_v1, 19)
				(_nw62).ArraySet1(_303_v1, 20)
				(_nw62).ArraySet1(_303_v1, 21)
				(_nw62).ArraySet1(_303_v1, 22)
				_343_v38 = _nw62
				var _344_v39 _dafny.Array
				_ = _344_v39
				var _nwElement0_13 bool = _303_v1
				_ = _nwElement0_13
				var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(19))
				_ = _nw63
				(_nw63).ArraySet1(_nwElement0_13, 0)
				(_nw63).ArraySet1(_303_v1, 1)
				(_nw63).ArraySet1(_303_v1, 2)
				(_nw63).ArraySet1((_338_v33).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_338_v33).Cardinality()))).Uint32()).(bool), 3)
				(_nw63).ArraySet1(_303_v1, 4)
				(_nw63).ArraySet1(Companion_Default___.Fm1(_this.F12(), true, _dafny.IntOfInt64(716), _303_v1, globalState), 5)
				(_nw63).ArraySet1(!(_303_v1) || (_303_v1), 6)
				(_nw63).ArraySet1(!(_303_v1) || (_303_v1), 7)
				(_nw63).ArraySet1(_303_v1, 8)
				(_nw63).ArraySet1(_303_v1, 9)
				(_nw63).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm11((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _this.F12(), _303_v1, globalState), _339_v34), 10)
				(_nw63).ArraySet1(!_dafny.Companion_Sequence_.Equal(_339_v34, (_340_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_338_v33, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))), _dafny.IntOfUint32((_338_v33).Cardinality()))).Uint32(), false)).Cardinality()), _dafny.IntOfUint32((_340_v35).Cardinality()))).Uint32()).(_dafny.Sequence)), 11)
				(_nw63).ArraySet1(_303_v1, 12)
				(_nw63).ArraySet1(_303_v1, 13)
				(_nw63).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v37, _343_v38)).Contains(_342_v37), 14)
				(_nw63).ArraySet1(_303_v1, 15)
				(_nw63).ArraySet1(_303_v1, 16)
				(_nw63).ArraySet1(!(_303_v1), 17)
				(_nw63).ArraySet1(_303_v1, 18)
				_344_v39 = _nw63
				var _rhs63 bool = _303_v1
				_ = _rhs63
				var _rhs64 _dafny.Int = _dafny.IntOfInt64(-754)
				_ = _rhs64
				var _rhs65 _dafny.Int = ((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)).Times((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
				_ = _rhs65
				var _rhs66 _dafny.Array = _344_v39
				_ = _rhs66
				var _rhs67 bool = !(!(!(_303_v1)) || (true))
				_ = _rhs67
				var _lhs32 *GlobalState = globalState
				_ = _lhs32
				var _lhs33 *GlobalState = globalState
				_ = _lhs33
				r0 = _rhs63
				_lhs32.F5 = _rhs64
				_lhs33.F5 = _rhs65
				_344_v39 = _rhs66
				r0 = _rhs67
				var _345_v40 _dafny.MultiSet
				_ = _345_v40
				_345_v40 = _dafny.MultiSetOf(_this.F12(), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
				var _346_v41 _dafny.Map
				_ = _346_v41
				_346_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v1, _303_v1)
				var _347_v42 _dafny.Map
				_ = _347_v42
				_347_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_345_v40).Contains((_346_v41).Cardinality()) {
						return (_345_v40).Multiplicity((_346_v41).Cardinality())
					}
					return (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)
				})(), false)
				var _348_v43 _dafny.Sequence
				_ = _348_v43
				_348_v43 = Companion_Default___.Fm11(Companion_Default___.Fm0(_this.F12(), _this.F12(), globalState), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _303_v1, globalState)
				var _349_v44 _dafny.Map
				_ = _349_v44
				_349_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v1, _dafny.IntOfUint32((_348_v43).Cardinality()))
				var _350_v45 *C0
				_ = _350_v45
				var _nw64 *C0 = New_C0_()
				_ = _nw64
				_nw64.Ctor__(_347_v42, (_349_v44).Cardinality())
				_350_v45 = _nw64
				var _351_v46 _dafny.Array
				_ = _351_v46
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_8
				var _nw65 _dafny.Array
				_ = _nw65
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw65 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Map = (func(_352_v41 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_353_i3 _dafny.Int) _dafny.Map {
							return _352_v41
						}
					})(_346_v41)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw65 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw65).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw65).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_351_v46 = _nw65
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_351_v46), 0))
				_ = _index51
				(_351_v46).ArraySet1((_346_v41).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v1, _303_v1)), (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.ArrayLen((_351_v46), 0))
				_ = _index52
				(_351_v46).ArraySet1(_346_v41, (_index52).Int())
				var _354_v47 _dafny.Map
				_ = _354_v47
				_354_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_338_v33, _this.F12())
				_354_v47 = (_354_v47).Update(_dafny.SeqOf(Companion_Default___.Fm1((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _303_v1, _this.F12(), _303_v1, globalState)), _this.F12())
				var _355_v48 _dafny.Set
				_ = _355_v48
				_355_v48 = _dafny.SetOf(true)
				var _356_v49 _dafny.Map
				_ = _356_v49
				_356_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_355_v48, (func() _dafny.Int {
					if (_345_v40).Contains((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)) {
						return (_345_v40).Multiplicity((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
					}
					return _dafny.IntOfUint32((_339_v34).Cardinality())
				})())
				var _357_v50 _dafny.CodePoint
				_ = _357_v50
				_357_v50 = _dafny.CodePoint('s')
				var _358_v51 _dafny.Array
				_ = _358_v51
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_9
				var _nw66 _dafny.Array
				_ = _nw66
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw66 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Map = (func(_359_v44 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_360_i4 _dafny.Int) _dafny.Map {
							return _359_v44
						}
					})(_349_v44)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw66 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw66).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw66).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_358_v51 = _nw66
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_358_v51), 0))
				_ = _index53
				(_358_v51).ArraySet1(_349_v44, (_index53).Int())
				var _361_v53 _dafny.Sequence
				_ = _361_v53
				_361_v53 = _dafny.SeqOf(_355_v48)
				var _362_v54 _dafny.Set
				_ = _362_v54
				_362_v54 = _355_v48
				var _363_v55 D2
				_ = _363_v55
				_363_v55 = Companion_D2_.Create_DC8_(_303_v1, (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _303_v1)
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_358_v51), 0))
				_ = _index54
				var _rhs68 _dafny.Map = ((func() _dafny.Map {
					var _coll6 = _dafny.NewMapBuilder()
					_ = _coll6
					for _iter6 := _dafny.Iterate((_361_v53).Elements()); ; {
						_compr_6, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _364_v52 _dafny.Set
						_364_v52 = interface{}(_compr_6).(_dafny.Set)
						if _dafny.Companion_Sequence_.Contains(_361_v53, _364_v52) {
							_coll6.Add(_364_v52, (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
						}
					}
					return _coll6.ToMap()
				}()).Update((_362_v54), (_363_v55).Dtor_cf12())).Update(_dafny.SetOf(_303_v1, _303_v1, _303_v1), Companion_Default___.Fm0(_this.F12(), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), globalState))
				_ = _rhs68
				var _rhs69 _dafny.CodePoint = _357_v50
				_ = _rhs69
				var _rhs70 _dafny.Map = _349_v44
				_ = _rhs70
				var _lhs34 _dafny.Array = _358_v51
				_ = _lhs34
				var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_358_v51), 0))
				_ = _lhs35
				_356_v49 = _rhs68
				_357_v50 = _rhs69
				(_lhs34).ArraySet1(_rhs70, (_lhs35).Int())
			}
			var _365_v56 _dafny.Sequence
			_ = _365_v56
			_365_v56 = _dafny.SeqOf(_303_v1)
			var _366_v57 D2
			_ = _366_v57
			_366_v57 = Companion_D2_.Create_DC7_(_dafny.IntOfInt64(706))
			var _367_v58 _dafny.MultiSet
			_ = _367_v58
			_367_v58 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_365_v56, _dafny.SeqOf(false)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(97))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_368_v56 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_369_i5 _dafny.Int) _dafny.Sequence {
					return _368_v56
				}
			})(_365_v56))))).Cardinality()), _this.F12(), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), (_366_v57).Dtor_cf10(), (_this.F12()).Minus(_this.F12()))
			_367_v58 = _367_v58
			_303_v1 = ((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)).Cmp(_this.F12()) <= 0
			var _370_v59 _dafny.Map
			_ = _370_v59
			_370_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _303_v1)
			var _371_v60 _dafny.MultiSet
			_ = _371_v60
			_371_v60 = _dafny.MultiSetOf(_303_v1, false, _303_v1, _303_v1, _303_v1)
			var _372_v63 _dafny.Array
			_ = _372_v63
			var _nwElement0_14 _dafny.Map = (_370_v59).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _303_v1))
			_ = _nwElement0_14
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(13))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_14, 0)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _303_v1), 1)
			(_nw67).ArraySet1((_370_v59).Merge(_370_v59), 2)
			(_nw67).ArraySet1(_370_v59, 3)
			(_nw67).ArraySet1(_370_v59, 4)
			(_nw67).ArraySet1(Companion_Default___.Fm10((_371_v60).Cardinality(), globalState), 5)
			(_nw67).ArraySet1(_370_v59, 6)
			(_nw67).ArraySet1((_370_v59).Merge(_370_v59), 7)
			(_nw67).ArraySet1(_370_v59, 8)
			(_nw67).ArraySet1(_370_v59, 9)
			(_nw67).ArraySet1(func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(14), _dafny.IntOfInt64(690))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _373_v61 _dafny.Int
					_373_v61 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(14)).Cmp(_373_v61) <= 0) && ((_373_v61).Cmp(_dafny.IntOfInt64(690)) < 0) {
						_coll7.Add(Companion_Default___.SafeModuloInt(_373_v61, _this.F12()), !(_303_v1))
					}
				}
				return _coll7.ToMap()
			}(), 10)
			(_nw67).ArraySet1(_370_v59, 11)
			(_nw67).ArraySet1(func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(382), _dafny.IntOfInt64(85))); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _374_v62 _dafny.Int
					_374_v62 = interface{}(_compr_8).(_dafny.Int)
					if ((_dafny.IntOfInt64(382)).Cmp(_374_v62) <= 0) && ((_374_v62).Cmp(_dafny.IntOfInt64(85)) < 0) {
						_coll8.Add((_374_v62).Plus(_this.F12()), _303_v1)
					}
				}
				return _coll8.ToMap()
			}(), 12)
			_372_v63 = _nw67
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_372_v63), 0))
			_ = _index55
			(_372_v63).ArraySet1((_370_v59).Update(_this.F12(), _303_v1), (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_372_v63), 0))
			_ = _index56
			(_372_v63).ArraySet1(_370_v59, (_index56).Int())
			var _375_v64 _dafny.Array
			_ = _375_v64
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_10
			var _nw68 _dafny.Array
			_ = _nw68
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw68 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) D2 = func(_376_i6 _dafny.Int) D2 {
					return Companion_D2_.Create_DC9_(_dafny.IntOfInt64(524))
				}
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw68 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw68).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw68).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_375_v64 = _nw68
			var _377_v65 _dafny.Map
			_ = _377_v65
			_377_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_v58, _375_v64)
			_377_v65 = (_377_v65).Update(_dafny.MultiSetOf(_this.F12()), _375_v64)
		}
		var _378_v66 _dafny.Map
		_ = _378_v66
		_378_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), !(!(true)))
		var _379_v67 _dafny.MultiSet
		_ = _379_v67
		_379_v67 = _dafny.MultiSetOf(_303_v1, _303_v1)
		var _380_v68 _dafny.Map
		_ = _380_v68
		_380_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(20), (_379_v67).Cardinality())
		var _381_v71 _dafny.Sequence
		_ = _381_v71
		_381_v71 = _dafny.SeqOf(_dafny.SetOf((_dafny.Zero).Minus((_dafny.SetOf(_378_v66)).Cardinality())), func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate((_380_v68).Keys().Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _382_v69 _dafny.Int
				_382_v69 = interface{}(_compr_9).(_dafny.Int)
				if (_380_v68).Contains(_382_v69) {
					_coll9.Add(Companion_Default___.SafeModuloInt(_382_v69, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC2_(!(true), func() _dafny.Map {
						var _coll10 = _dafny.NewMapBuilder()
						_ = _coll10
						for _iter10 := _dafny.Iterate((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(242), false)).Cardinality(), _dafny.IntOfInt64(192))).Elements()); ; {
							_compr_10, _ok10 := _iter10()
							if !_ok10 {
								break
							}
							var _383_v70 _dafny.Int
							_383_v70 = interface{}(_compr_10).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(242), false)).Cardinality(), _dafny.IntOfInt64(192)), _383_v70) {
								_coll10.Add((_383_v70).Plus(_dafny.IntOfInt64(851)), !(true))
							}
						}
						return _coll10.ToMap()
					}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('f')), true)).Dtor_cf2(), true)).Cardinality()))
				}
			}
			return _coll9.ToSet()
		}())
		var _384_v72 _dafny.Sequence
		_ = _384_v72
		_384_v72 = _dafny.UnicodeSeqOfUtf8Bytes("lvadps")
		var _385_v73 _dafny.Set
		_ = _385_v73
		_385_v73 = _dafny.SetOf(_this.F12())
		if ((_381_v71).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_384_v72).Cardinality()), _dafny.IntOfUint32((_381_v71).Cardinality()))).Uint32()).(_dafny.Set)).Equals((_385_v73).Difference(_385_v73)) {
			var _386_v74 _dafny.Sequence
			_ = _386_v74
			_386_v74 = _dafny.SeqOf((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))
			r0 = Companion_Default___.Fm1((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), false, (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), ((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_386_v74).Cardinality())) > 0, globalState)
			var _387_v75 _dafny.Array
			_ = _387_v75
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw69
			_387_v75 = _nw69
			var _388_v76 D0
			_ = _388_v76
			_388_v76 = Companion_D0_.Create_DC0_(_303_v1)
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_387_v75), 0))
			_ = _index57
			(_387_v75).ArraySet1((_388_v76).Dtor_cf0(), (_index57).Int())
			var _389_v77 _dafny.Array
			_ = _389_v77
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
			_ = _nw70
			_389_v77 = _nw70
			var _390_v78 _dafny.Set
			_ = _390_v78
			_390_v78 = _dafny.SetOf(_389_v77, _389_v77)
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))
			_ = _index58
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_387_v75), 0))
			_ = _index59
			var _rhs71 _dafny.Int = (_390_v78).Cardinality()
			_ = _rhs71
			var _rhs72 bool = ((_dafny.Zero).Minus((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))).Cmp(Companion_Default___.SafeModuloInt((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))) <= 0
			_ = _rhs72
			var _rhs73 _dafny.Int = _this.F12()
			_ = _rhs73
			var _rhs74 _dafny.Int = (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs74
			var _lhs36 _dafny.Array = _301_v0
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))
			_ = _lhs37
			var _lhs38 _dafny.Array = _387_v75
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_387_v75), 0))
			_ = _lhs39
			(_lhs36).ArraySet1(_rhs71, (_lhs37).Int())
			(_lhs38).ArraySet1(_rhs72, (_lhs39).Int())
			r3 = _rhs73
			r3 = _rhs74
			var _391_v79 *C0
			_ = _391_v79
			var _nw71 *C0 = New_C0_()
			_ = _nw71
			_nw71.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(550), true), _dafny.IntOfInt64(688))
			_391_v79 = _nw71
			var _392_v80 _dafny.Array
			_ = _392_v80
			var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw72
			_392_v80 = _nw72
			var _393_v81 _dafny.Sequence
			_ = _393_v81
			_393_v81 = _dafny.SeqOf(_dafny.SeqOf((_387_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_387_v75), 0))).Int()).(bool)), _dafny.SeqOf((_387_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_387_v75), 0))).Int()).(bool), _303_v1))
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_392_v80), 0))
			_ = _index60
			(_392_v80).ArraySet1(_393_v81, (_index60).Int())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_392_v80), 0))
			_ = _index61
			(_392_v80).ArraySet1(_393_v81, (_index61).Int())
			_387_v75 = _387_v75
		} else {
			var _394_v82 *C0
			_ = _394_v82
			var _nw73 *C0 = New_C0_()
			_ = _nw73
			_nw73.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_384_v72).Cardinality()), true), _this.F12())
			_394_v82 = _nw73
			var _395_v83 *C0
			_ = _395_v83
			var _nw74 *C0 = New_C0_()
			_ = _nw74
			_nw74.Ctor__(_378_v66, _this.F12())
			_395_v83 = _nw74
			_395_v83 = _395_v83
			var _396_v84 _dafny.Array
			_ = _396_v84
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_11
			var _nw75 _dafny.Array
			_ = _nw75
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw75 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Sequence = (func(_397_v72 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_398_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_397_v72, _397_v72)
					}
				})(_384_v72)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw75 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw75).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw75).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_396_v84 = _nw75
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_396_v84), 0))
			_ = _index62
			(_396_v84).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("r"), (_index62).Int())
			var _399_v85 _dafny.Sequence
			_ = _399_v85
			_399_v85 = _384_v72
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_396_v84), 0))
			_ = _index63
			(_396_v84).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if _303_v1 {
					return _384_v72
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("hu")
			})(), (_399_v85)), (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_396_v84), 0))
			_ = _index64
			(_396_v84).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_384_v72, _dafny.UnicodeSeqOfUtf8Bytes("y")), _dafny.UnicodeSeqOfUtf8Bytes("bsdwgeit")), (_index64).Int())
		}
		var _400_v86 _dafny.Sequence
		_ = _400_v86
		_400_v86 = _dafny.SeqOf(_303_v1)
		var _401_v87 D2
		_ = _401_v87
		_401_v87 = Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_400_v86).Cardinality()))
		var _402_v88 _dafny.Sequence
		_ = _402_v88
		_402_v88 = _dafny.SeqOf((_401_v87).Dtor_cf10())
		var _403_v89 D1
		_ = _403_v89
		_403_v89 = Companion_D1_.Create_DC4_(_301_v0)
		var _404_v90 _dafny.Map
		_ = _404_v90
		_404_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_402_v88).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_402_v88).Cardinality()))).Uint32()).(_dafny.Int), (_403_v89).Dtor_cf7())
		_404_v90 = (_404_v90).Update(((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)).Plus((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)), _301_v0)
		var _hi1 _dafny.Int = _this.F12()
		_ = _hi1
		for _405_i8 := (_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int); _405_i8.Cmp(_hi1) < 0; _405_i8 = _405_i8.Plus(_dafny.One) {
			var _406_v91 _dafny.Set
			_ = _406_v91
			_406_v91 = _dafny.SetOf(_400_v86)
			var _407_v92 _dafny.Map
			_ = _407_v92
			_407_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_303_v1, _303_v1), _dafny.MultiSetOf(_303_v1, _303_v1, false, Companion_Default___.Fm1((_385_v73).Cardinality(), _303_v1, _this.F12(), _303_v1, globalState)))
			if (func() bool {
				if true {
					return (func() _dafny.Set {
						var _coll11 = _dafny.NewBuilder()
						_ = _coll11
						for _iter11 := _dafny.Iterate((_407_v92).Keys().Elements()); ; {
							_compr_11, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _408_v93 _dafny.Sequence
							_408_v93 = interface{}(_compr_11).(_dafny.Sequence)
							if (_407_v92).Contains(_408_v93) {
								_coll11.Add(_408_v93)
							}
						}
						return _coll11.ToSet()
					}()).IsProperSubsetOf(_406_v91)
				}
				return !(!((func() bool {
					if _303_v1 {
						return _303_v1
					}
					return _303_v1
				})()))
			})() {
				_384_v72 = _dafny.Companion_Sequence_.Concatenate(_384_v72, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-898))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}(func(_409_i9 _dafny.Int) _dafny.CodePoint {
					return (_this).F15()
				})))
				_384_v72 = _384_v72
				r0 = _303_v1
				var _410_v94 _dafny.CodePoint
				_ = _410_v94
				_410_v94 = _dafny.CodePoint('p')
				_410_v94 = (_this).F13()
				var _411_v95 _dafny.Array
				_ = _411_v95
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(18))
				_ = _nw76
				_411_v95 = _nw76
				_411_v95 = _411_v95
			} else {
				var _412_v96 _dafny.Array
				_ = _412_v96
				var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
				_ = _nw77
				_412_v96 = _nw77
				var _413_v97 _dafny.Map
				_ = _413_v97
				_413_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_412_v96, _303_v1)
				_413_v97 = (_413_v97).Update(_412_v96, !(Companion_Default___.Fm1(_this.F12(), !(_303_v1), _405_i8, _303_v1, globalState)) || (_303_v1))
				var _414_v98 T2
				_ = _414_v98
				var _nw78 *C0 = New_C0_()
				_ = _nw78
				_nw78.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_402_v88).Cardinality()), _303_v1), _this.F12())
				_414_v98 = _nw78
				var _415_v99 _dafny.Sequence
				_ = _415_v99
				_415_v99 = _dafny.SeqOf(_414_v98, _414_v98, _414_v98, _414_v98, _414_v98)
				var _416_v100 _dafny.Map
				_ = _416_v100
				_416_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_415_v99, _414_v98.F12())
				_416_v100 = (_416_v100).Update(_dafny.SeqOf(_414_v98, _414_v98, _414_v98, _414_v98), (_this.F12()).Times((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)))
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))
				_ = _index65
				(_301_v0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_384_v72, (Companion_Default___.SafeIndex((_this).Fm4(globalState), _dafny.IntOfUint32((_384_v72).Cardinality()))).Uint32(), (_this).F15())).Cardinality()), (_index65).Int())
				_380_v68 = Companion_Default___.Fm12((_dafny.Zero).Minus((_this.F12()).Minus((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int))), _414_v98.F12(), _this.F12(), (Companion_Default___.Fm0(_dafny.IntOfUint32((_400_v86).Cardinality()), (_380_v68).Cardinality(), globalState)).Plus((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int)), globalState)
				var _417_v101 _dafny.Set
				_ = _417_v101
				_417_v101 = _dafny.SetOf(_380_v68, _380_v68, _380_v68)
				var _418_v102 _dafny.Array
				_ = _418_v102
				var _nwElement0_15 bool = _303_v1
				_ = _nwElement0_15
				var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(13))
				_ = _nw79
				(_nw79).ArraySet1(_nwElement0_15, 0)
				(_nw79).ArraySet1(_303_v1, 1)
				(_nw79).ArraySet1(_303_v1, 2)
				(_nw79).ArraySet1(_303_v1, 3)
				(_nw79).ArraySet1(_303_v1, 4)
				(_nw79).ArraySet1(_303_v1, 5)
				(_nw79).ArraySet1(true, 6)
				(_nw79).ArraySet1(!(_303_v1), 7)
				(_nw79).ArraySet1(_303_v1, 8)
				(_nw79).ArraySet1(_dafny.Companion_Sequence_.Contains(_400_v86, _303_v1), 9)
				(_nw79).ArraySet1(_303_v1, 10)
				(_nw79).ArraySet1(_303_v1, 11)
				(_nw79).ArraySet1((_417_v101).IsSubsetOf(_417_v101), 12)
				_418_v102 = _nw79
				_418_v102 = (func() _dafny.Array {
					if (func() bool {
						if (_378_v66).Contains((_380_v68).Cardinality()) {
							return (_378_v66).Get((_380_v68).Cardinality()).(bool)
						}
						return _303_v1
					})() {
						return _418_v102
					}
					return _418_v102
				})()
			}
			_402_v88 = _402_v88
			_303_v1 = !(true)
			var _419_v103 D0
			_ = _419_v103
			_419_v103 = Companion_D0_.Create_DC0_(_303_v1)
			_303_v1 = !((_419_v103).Dtor_cf0())
		}
		r0 = !(Companion_Default___.Fm1(_this.F12(), _303_v1, (_this.F12()).Times(_this.F12()), false, globalState))
		r1 = _301_v0
		var _420_v104 _dafny.Set
		_ = _420_v104
		_420_v104 = _dafny.SetOf(_402_v88)
		var _421_v105 _dafny.MultiSet
		_ = _421_v105
		_421_v105 = _dafny.MultiSetOf((_301_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_301_v0), 0))).Int()).(_dafny.Int), _this.F12())
		var _422_v106 _dafny.Map
		_ = _422_v106
		_422_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), !((func() bool {
			if (_378_v66).Contains((_420_v104).Cardinality()) {
				return (_378_v66).Get((_420_v104).Cardinality()).(bool)
			}
			return Companion_Default___.Fm1(_this.F12(), _303_v1, (_421_v105).Cardinality(), _303_v1, globalState)
		})()))
		r2 = (_422_v106).Merge(_422_v106)
		r3 = _dafny.IntOfUint32((_384_v72).Cardinality())
		return r0, r1, r2, r3
	}
}
func (_this *C1) M5(globalState *GlobalState) (_dafny.Int, bool, _dafny.Map) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var _423_v0 _dafny.Set
		_ = _423_v0
		_423_v0 = _dafny.SetOf(_this.F12(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lqyamhoax")).Cardinality()))
		r1 = (_423_v0).IsSubsetOf(_423_v0)
		(_this).F12_set_(func(_source3 D2) _dafny.Int {
			if _source3.Is_DC8() {
				var _424___mcc_h0 bool = _source3.Get_().(D2_DC8).Cf11
				_ = _424___mcc_h0
				var _425___mcc_h1 _dafny.Int = _source3.Get_().(D2_DC8).Cf12
				_ = _425___mcc_h1
				var _426___mcc_h2 bool = _source3.Get_().(D2_DC8).Cf13
				_ = _426___mcc_h2
				var _427_cf13 bool = _426___mcc_h2
				_ = _427_cf13
				var _428_cf12 _dafny.Int = _425___mcc_h1
				_ = _428_cf12
				var _429_cf11 bool = _424___mcc_h0
				_ = _429_cf11
				return _428_cf12
			} else if _source3.Is_DC9() {
				var _430___mcc_h3 _dafny.Int = _source3.Get_().(D2_DC9).Cf14
				_ = _430___mcc_h3
				var _431_cf14 _dafny.Int = _430___mcc_h3
				_ = _431_cf14
				return _431_cf14
			} else {
				var _432___mcc_h4 _dafny.Int = _source3.Get_().(D2_DC7).Cf10
				_ = _432___mcc_h4
				var _433_cf10 _dafny.Int = _432___mcc_h4
				_ = _433_cf10
				return _433_cf10
			}
		}(Companion_D2_.Create_DC9_(_this.F12())))
		var _434_v1 bool
		_ = _434_v1
		_434_v1 = true
		var _435_i0 _dafny.Int
		_ = _435_i0
		_435_i0 = _dafny.Zero
		{
			for _434_v1 {
				{
					if (_435_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_435_i0 = (_435_i0).Plus(_dafny.One)
					var _436_v2 _dafny.Sequence
					_ = _436_v2
					_436_v2 = _dafny.SeqOf(_434_v1, _434_v1)
					var _437_v3 _dafny.Sequence
					_ = _437_v3
					_437_v3 = _dafny.UnicodeSeqOfUtf8Bytes("iy")
					var _438_v4 _dafny.Map
					_ = _438_v4
					_438_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if _434_v1 {
							return _this.F12()
						}
						return _this.F12()
					})(), (Companion_Default___.Fm0(_dafny.IntOfUint32((_436_v2).Cardinality()), _dafny.IntOfUint32((_437_v3).Cardinality()), globalState)).Cmp(_this.F12()) <= 0)
					_438_v4 = (_438_v4).Update(_this.F12(), !(_434_v1))
					var _439_v5 _dafny.Array
					_ = _439_v5
					var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
					_ = _nw80
					_439_v5 = _nw80
					var _440_v6 D2
					_ = _440_v6
					_440_v6 = Companion_D2_.Create_DC8_(_434_v1, _this.F12(), _434_v1)
					var _441_v7 _dafny.Array
					_ = _441_v7
					var _nwElement0_16 bool = (_440_v6).Dtor_cf11()
					_ = _nwElement0_16
					var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(5))
					_ = _nw81
					(_nw81).ArraySet1(_nwElement0_16, 0)
					(_nw81).ArraySet1(_434_v1, 1)
					(_nw81).ArraySet1(_434_v1, 2)
					(_nw81).ArraySet1(_434_v1, 3)
					(_nw81).ArraySet1(_434_v1, 4)
					_441_v7 = _nw81
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_439_v5), 0))
					_ = _index66
					(_439_v5).ArraySet1(_441_v7, (_index66).Int())
					var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_439_v5), 0))
					_ = _index67
					var _rhs75 bool = !((Companion_Default___.Fm13(_434_v1, _434_v1, globalState)).Contains(_this.F12())) || (!(_434_v1))
					_ = _rhs75
					var _rhs76 _dafny.Array = _441_v7
					_ = _rhs76
					var _lhs40 _dafny.Array = _439_v5
					_ = _lhs40
					var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_439_v5), 0))
					_ = _lhs41
					_434_v1 = _rhs75
					(_lhs40).ArraySet1(_rhs76, (_lhs41).Int())
					var _442_v8 _dafny.Set
					_ = _442_v8
					_442_v8 = _dafny.SetOf(_434_v1)
					var _443_v9 D2
					_ = _443_v9
					_443_v9 = Companion_D2_.Create_DC9_(_this.F12())
					_442_v8 = Companion_Default___.Fm14(_this.F12(), _443_v9, globalState)
					var _444_v10 _dafny.Sequence
					_ = _444_v10
					_444_v10 = _dafny.SeqOf(_this.F12())
					var _445_v11 _dafny.MultiSet
					_ = _445_v11
					_445_v11 = _dafny.MultiSetOf(_dafny.CodePoint('j'), (_this).F15(), (_this).F15())
					(globalState).F5 = Companion_Default___.SafeModuloInt(_this.F12(), (_444_v10).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_445_v11).Contains((_this).F15()) {
							return (_445_v11).Multiplicity((_this).F15())
						}
						return _this.F12()
					})(), _dafny.IntOfUint32((_444_v10).Cardinality()))).Uint32()).(_dafny.Int))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		r1 = (func() bool {
			if _434_v1 {
				return (func() bool {
					if _434_v1 {
						return _434_v1
					}
					return _434_v1
				})()
			}
			return !(_434_v1)
		})()
		var _446_v12 _dafny.Sequence
		_ = _446_v12
		_446_v12 = _dafny.SeqOf(_this.F12(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(164), _this.F12()), _this.F12())
		var _447_v13 _dafny.Sequence
		_ = _447_v13
		_447_v13 = _dafny.SeqOf(true, _434_v1)
		var _448_v14 _dafny.Sequence
		_ = _448_v14
		_448_v14 = _dafny.UnicodeSeqOfUtf8Bytes("iiunb")
		_446_v12 = _dafny.Companion_Sequence_.Update((_this).Fm5(_dafny.Companion_Sequence_.IsPrefixOf(_447_v13, _447_v13), _434_v1, _434_v1, true, globalState), (Companion_Default___.SafeIndex((_dafny.IntOfUint32((_448_v14).Cardinality())).Minus(_this.F12()), _dafny.IntOfUint32(((_this).Fm5(_dafny.Companion_Sequence_.IsPrefixOf(_447_v13, _447_v13), _434_v1, _434_v1, true, globalState)).Cardinality()))).Uint32(), (_dafny.Zero).Minus((_this.F12()).Times(_this.F12())))
		var _449_v15 _dafny.Array
		_ = _449_v15
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
		_ = _nw82
		_449_v15 = _nw82
		var _450_v16 _dafny.Array
		_ = _450_v16
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_12
		var _nw83 _dafny.Array
		_ = _nw83
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw83 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Set = (func(_451_v0 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_452_i1 _dafny.Int) _dafny.Set {
					return _451_v0
				}
			})(_423_v0)
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
		_450_v16 = _nw83
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_449_v15), 0))
		_ = _index68
		(_449_v15).ArraySet1(_450_v16, (_index68).Int())
		var _453_v17 _dafny.Map
		_ = _453_v17
		_453_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v1, _450_v16)
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_449_v15), 0))
		_ = _index69
		(_449_v15).ArraySet1((func() _dafny.Array {
			if (_453_v17).Contains(Companion_Default___.Fm1((_dafny.Zero).Minus(_dafny.IntOfUint32((_448_v14).Cardinality())), _434_v1, _this.F12(), _434_v1, globalState)) {
				return (_453_v17).Get(Companion_Default___.Fm1((_dafny.Zero).Minus(_dafny.IntOfUint32((_448_v14).Cardinality())), _434_v1, _this.F12(), _434_v1, globalState)).(_dafny.Array)
			}
			return _450_v16
		})(), (_index69).Int())
		r0 = (_dafny.Zero).Minus(func(_source4 D0) _dafny.Int {
			if _source4.Is_DC1() {
				var _454___mcc_h5 _dafny.Array = _source4.Get_().(D0_DC1).Cf1
				_ = _454___mcc_h5
				var _455_cf1 _dafny.Array = _454___mcc_h5
				_ = _455_cf1
				return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(592), _this.F12())
			} else if _source4.Is_DC2() {
				var _456___mcc_h6 bool = _source4.Get_().(D0_DC2).Cf2
				_ = _456___mcc_h6
				var _457___mcc_h7 _dafny.Map = _source4.Get_().(D0_DC2).Cf3
				_ = _457___mcc_h7
				var _458___mcc_h8 _dafny.Map = _source4.Get_().(D0_DC2).Cf4
				_ = _458___mcc_h8
				var _459___mcc_h9 bool = _source4.Get_().(D0_DC2).Cf5
				_ = _459___mcc_h9
				var _460_cf5 bool = _459___mcc_h9
				_ = _460_cf5
				var _461_cf4 _dafny.Map = _458___mcc_h8
				_ = _461_cf4
				var _462_cf3 _dafny.Map = _457___mcc_h7
				_ = _462_cf3
				var _463_cf2 bool = _456___mcc_h6
				_ = _463_cf2
				return (_this.F12()).Plus(_this.F12())
			} else if _source4.Is_DC0() {
				var _464___mcc_h10 bool = _source4.Get_().(D0_DC0).Cf0
				_ = _464___mcc_h10
				var _465_cf0 bool = _464___mcc_h10
				_ = _465_cf0
				return _this.F12()
			} else {
				var _466___mcc_h11 D0 = _source4.Get_().(D0_DC3).Cf6
				_ = _466___mcc_h11
				var _467_cf6 D0 = _466___mcc_h11
				_ = _467_cf6
				return _dafny.IntOfInt64(103)
			}
		}(Companion_D0_.Create_DC0_(Companion_Default___.Fm1(_this.F12(), !(_434_v1), _this.F12(), _434_v1, globalState))))
		var _468_v18 D1
		_ = _468_v18
		_468_v18 = Companion_D1_.Create_DC5_(_434_v1)
		var _469_v19 _dafny.Map
		_ = _469_v19
		_469_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_468_v18).Dtor_cf8(), _dafny.IntOfInt64(623))
		r1 = (func() bool {
			if !(_434_v1) {
				return _434_v1
			}
			return Companion_Default___.Fm1(_this.F12(), !(_434_v1), (_dafny.Zero).Minus((_469_v19).Cardinality()), _434_v1, globalState)
		})()
		var _470_v20 _dafny.Set
		_ = _470_v20
		_470_v20 = _dafny.SetOf(_434_v1, false, _434_v1, _434_v1)
		var _471_v21 _dafny.Map
		_ = _471_v21
		_471_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _434_v1)
		var _472_v22 _dafny.Map
		_ = _472_v22
		_472_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() bool {
			if (_471_v21).Contains(_dafny.IntOfInt64(421)) {
				return (_471_v21).Get(_dafny.IntOfInt64(421)).(bool)
			}
			return !(_434_v1)
		})())
		var _473_v23 _dafny.MultiSet
		_ = _473_v23
		_473_v23 = _dafny.MultiSetOf(_this.F12(), (_472_v22).Cardinality())
		var _474_v24 _dafny.Map
		_ = _474_v24
		_474_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_434_v1), _473_v23)
		var _475_v26 _dafny.Sequence
		_ = _475_v26
		_475_v26 = _dafny.SeqOf(_470_v20)
		r2 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_470_v20, _473_v23)).Merge((_474_v24).Merge(func() _dafny.Map {
			var _coll12 = _dafny.NewMapBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate((_475_v26).Elements()); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _476_v25 _dafny.Set
				_476_v25 = interface{}(_compr_12).(_dafny.Set)
				if _dafny.Companion_Sequence_.Contains(_475_v26, _476_v25) {
					_coll12.Add(_476_v25, _473_v23)
				}
			}
			return _coll12.ToMap()
		}()))
		return r0, r1, r2
	}
}
func (_this *C1) F15() _dafny.CodePoint {
	{
		return _this._f15
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f12 _dafny.Int
	_f17 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f12 = _dafny.Zero
	_this._f17 = false
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

func (_this *C2) F12() _dafny.Int {
	return _this._f12
}
func (_this *C2) F12_set_(value _dafny.Int) {
	_this._f12 = value
}
func (_this *C2) Ctor__(f17 bool, f12 _dafny.Int) {
	{
		(_this)._f17 = f17
		(_this)._f12 = f12
	}
}
func (_this *C2) F17() bool {
	{
		return _this._f17
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f12 _dafny.Int
	_f13 _dafny.CodePoint
	_f14 _dafny.Array
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.CodePoint('D')
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C3) F12() _dafny.Int {
	return _this._f12
}
func (_this *C3) F12_set_(value _dafny.Int) {
	_this._f12 = value
}
func (_this *C3) F13() _dafny.CodePoint {
	return _this._f13
}
func (_this *C3) Ctor__(f14 _dafny.Array, f13 _dafny.CodePoint, f12 _dafny.Int) {
	{
		(_this)._f14 = f14
		(_this)._f13 = f13
		(_this)._f12 = f12
	}
}
func (_this *C3) Fm4(globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(_this.F12(), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(632), (_dafny.MultiSetOf(true)).Cardinality()))
	}
}
func (_this *C3) Fm5(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F12()), _dafny.SeqOf(_this.F12()))
	}
}
func (_this *C3) Fm6(globalState *GlobalState) D0 {
	{
		return Companion_D0_.Create_DC0_(false)
	}
}
func (_this *C3) Fm7(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if false {
				return _dafny.UnicodeSeqOfUtf8Bytes("kdgs")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("xerenpimg")
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sugwvjq"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(567))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_477_i0 _dafny.Int) _dafny.CodePoint {
			return (_this).F13()
		}))))).Cardinality())
	}
}
func (_this *C3) M2(globalState *GlobalState) (bool, _dafny.Array, _dafny.Map, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _478_v0 _dafny.Array
		_ = _478_v0
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_13
		var _nw84 _dafny.Array
		_ = _nw84
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw84 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.Int = func(_479_i0 _dafny.Int) _dafny.Int {
				return (_479_i0).Plus(_this.F12())
			}
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw84 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw84).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw84).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_478_v0 = _nw84
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))
		_ = _index70
		(_478_v0).ArraySet1(Companion_Default___.SafeModuloInt(_this.F12(), _this.F12()), (_index70).Int())
		var _480_v1 bool
		_ = _480_v1
		_480_v1 = true
		var _481_v2 D2
		_ = _481_v2
		_481_v2 = Companion_D2_.Create_DC8_(_480_v1, _this.F12(), _480_v1)
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))
		_ = _index71
		(_478_v0).ArraySet1(func(_source5 D2) _dafny.Int {
			if _source5.Is_DC8() {
				var _482___mcc_h0 bool = _source5.Get_().(D2_DC8).Cf11
				_ = _482___mcc_h0
				var _483___mcc_h1 _dafny.Int = _source5.Get_().(D2_DC8).Cf12
				_ = _483___mcc_h1
				var _484___mcc_h2 bool = _source5.Get_().(D2_DC8).Cf13
				_ = _484___mcc_h2
				var _485_cf13 bool = _484___mcc_h2
				_ = _485_cf13
				var _486_cf12 _dafny.Int = _483___mcc_h1
				_ = _486_cf12
				var _487_cf11 bool = _482___mcc_h0
				_ = _487_cf11
				return (_486_cf12).Plus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
			} else if _source5.Is_DC9() {
				var _488___mcc_h3 _dafny.Int = _source5.Get_().(D2_DC9).Cf14
				_ = _488___mcc_h3
				var _489_cf14 _dafny.Int = _488___mcc_h3
				_ = _489_cf14
				return _this.F12()
			} else {
				var _490___mcc_h4 _dafny.Int = _source5.Get_().(D2_DC7).Cf10
				_ = _490___mcc_h4
				var _491_cf10 _dafny.Int = _490___mcc_h4
				_ = _491_cf10
				return _this.F12()
			}
		}(_481_v2), (_index71).Int())
		var _492_i1 _dafny.Int
		_ = _492_i1
		_492_i1 = _dafny.Zero
		{
			for _480_v1 {
				{
					if (_492_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_492_i1 = (_492_i1).Plus(_dafny.One)
					var _493_v3 _dafny.Sequence
					_ = _493_v3
					_493_v3 = _dafny.UnicodeSeqOfUtf8Bytes("dii")
					(globalState).F5 = (Companion_Default___.Fm0(_dafny.IntOfUint32((_493_v3).Cardinality()), (_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int), globalState)).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_493_v3, (Companion_Default___.SafeIndex((_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_493_v3).Cardinality()))).Uint32(), (_this).F13())).Cardinality()))
					var _494_v4 _dafny.Map
					_ = _494_v4
					_494_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(858), true)
					var _495_v5 *C0
					_ = _495_v5
					var _nw85 *C0 = New_C0_()
					_ = _nw85
					_nw85.Ctor__(_494_v4, Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0((_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int), (_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int), globalState), (_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int)))
					_495_v5 = _nw85
					var _496_v6 _dafny.Sequence
					_ = _496_v6
					_496_v6 = _dafny.SeqOf(_480_v1)
					_496_v6 = (func() _dafny.Sequence {
						if _480_v1 {
							return _496_v6
						}
						return _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_480_v1, _480_v1, _480_v1, _480_v1, _480_v1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F12()), _dafny.IntOfUint32((_dafny.SeqOf(_480_v1, _480_v1, _480_v1, _480_v1, _480_v1)).Cardinality()))).Uint32(), _480_v1)
					})()
					var _497_v7 _dafny.Array
					_ = _497_v7
					var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
					_ = _nw86
					_497_v7 = _nw86
					var _498_v8 _dafny.Array
					_ = _498_v8
					var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
					_ = _nw87
					_498_v8 = _nw87
					var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_497_v7), 0))
					_ = _index72
					(_497_v7).ArraySet1(_498_v8, (_index72).Int())
					var _499_v9 _dafny.Sequence
					_ = _499_v9
					_499_v9 = _dafny.SeqOf(_498_v8)
					var _500_v10 _dafny.MultiSet
					_ = _500_v10
					_500_v10 = _dafny.MultiSetOf((_this).F13(), (_this).F13(), (_this).F13(), (_this).F13())
					var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_497_v7), 0))
					_ = _index73
					(_497_v7).ArraySet1((_499_v9).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_500_v10).Contains((_this).F13()) {
							return (_500_v10).Multiplicity((_this).F13())
						}
						return _this.F12()
					})(), _dafny.IntOfUint32((_499_v9).Cardinality()))).Uint32()).(_dafny.Array), (_index73).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _501_v11 _dafny.Map
		_ = _501_v11
		_501_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int), _478_v0)
		_501_v11 = (_501_v11).Update((_dafny.Zero).Minus((_dafny.IntOfInt64(497)).Minus((_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int))), _478_v0)
		r3 = (_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int)
		_480_v1 = ((_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int)).Cmp((_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int)) < 0
		r0 = _480_v1
		var _502_v12 _dafny.MultiSet
		_ = _502_v12
		_502_v12 = _dafny.MultiSetOf(_480_v1, _480_v1)
		var _503_v13 D5
		_ = _503_v13
		_503_v13 = Companion_D5_.Create_DC12_(_dafny.SeqOf(_502_v12, _502_v12))
		r0 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate((_503_v13).Dtor_cf17(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(605))).Uint32(), func(coer30 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}((func(_504_v12 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
			return func(_505_i2 _dafny.Int) _dafny.MultiSet {
				return _504_v12
			}
		})(_502_v12)))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(924))).Uint32(), func(coer31 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_506_v12 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
			return func(_507_i3 _dafny.Int) _dafny.MultiSet {
				return _506_v12
			}
		})(_502_v12))))
		r1 = _478_v0
		var _508_v14 _dafny.Map
		_ = _508_v14
		_508_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _480_v1)
		r2 = (_508_v14).Merge(_508_v14)
		var _509_v15 _dafny.Set
		_ = _509_v15
		_509_v15 = _dafny.SetOf(_480_v1, Companion_Default___.Fm1(_this.F12(), _480_v1, (_478_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_478_v0), 0))).Int()).(_dafny.Int), _480_v1, globalState))
		var _510_v16 _dafny.Set
		_ = _510_v16
		_510_v16 = _dafny.SetOf(_dafny.SetOf(_480_v1, _480_v1, _480_v1, false), _dafny.SetOf(_480_v1, false, _480_v1, _480_v1), _509_v15)
		r3 = (_510_v16).Cardinality()
		return r0, r1, r2, r3
	}
}
func (_this *C3) M3(p0 _dafny.MultiSet, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _511_v0 bool
		_ = _511_v0
		_511_v0 = true
		var _512_i0 _dafny.Int
		_ = _512_i0
		_512_i0 = _dafny.Zero
		{
			for _511_v0 {
				{
					if (_512_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_512_i0 = (_512_i0).Plus(_dafny.One)
					var _513_v1 _dafny.Map
					_ = _513_v1
					_513_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(902), _this.F12())
					var _514_v2 _dafny.Array
					_ = _514_v2
					var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
					_ = _nw88
					_514_v2 = _nw88
					var _515_v3 _dafny.Sequence
					_ = _515_v3
					_515_v3 = _dafny.SeqOf(_this.F12(), _dafny.IntOfInt64(842))
					var _516_v4 _dafny.Sequence
					_ = _516_v4
					_516_v4 = _dafny.SeqOf(_this.F12(), _dafny.IntOfUint32((_515_v3).Cardinality()))
					var _517_v5 _dafny.Array
					_ = _517_v5
					var _nwElement0_17 bool = _511_v0
					_ = _nwElement0_17
					var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(22))
					_ = _nw89
					(_nw89).ArraySet1(_nwElement0_17, 0)
					(_nw89).ArraySet1(true, 1)
					(_nw89).ArraySet1(_511_v0, 2)
					(_nw89).ArraySet1(_511_v0, 3)
					(_nw89).ArraySet1(_511_v0, 4)
					(_nw89).ArraySet1(Companion_Default___.Fm1(_this.F12(), _511_v0, _this.F12(), _511_v0, globalState), 5)
					(_nw89).ArraySet1(_511_v0, 6)
					(_nw89).ArraySet1(_511_v0, 7)
					(_nw89).ArraySet1(_511_v0, 8)
					(_nw89).ArraySet1(_511_v0, 9)
					(_nw89).ArraySet1(_511_v0, 10)
					(_nw89).ArraySet1(_511_v0, 11)
					(_nw89).ArraySet1(false, 12)
					(_nw89).ArraySet1(_511_v0, 13)
					(_nw89).ArraySet1(_511_v0, 14)
					(_nw89).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((_515_v3).Cardinality()), _511_v0, _dafny.IntOfInt64(129), _511_v0, globalState), 15)
					(_nw89).ArraySet1(Companion_Default___.Fm1(_this.F12(), _511_v0, _this.F12(), _511_v0, globalState), 16)
					(_nw89).ArraySet1(_511_v0, 17)
					(_nw89).ArraySet1(_511_v0, 18)
					(_nw89).ArraySet1(_511_v0, 19)
					(_nw89).ArraySet1(_511_v0, 20)
					(_nw89).ArraySet1(_511_v0, 21)
					_517_v5 = _nw89
					var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_514_v2), 0))
					_ = _index74
					(_514_v2).ArraySet1(_517_v5, (_index74).Int())
					var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_514_v2), 0))
					_ = _index75
					var _rhs77 _dafny.Map = _513_v1
					_ = _rhs77
					var _rhs78 _dafny.Array = (func() _dafny.Array {
						if _511_v0 {
							return _517_v5
						}
						return _517_v5
					})()
					_ = _rhs78
					var _lhs42 _dafny.Array = _514_v2
					_ = _lhs42
					var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_514_v2), 0))
					_ = _lhs43
					_513_v1 = _rhs77
					(_lhs42).ArraySet1(_rhs78, (_lhs43).Int())
					var _518_v6 _dafny.Sequence
					_ = _518_v6
					_518_v6 = _dafny.UnicodeSeqOfUtf8Bytes("doktuoubt")
					if (Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_518_v6).Cardinality())), _this.F12())).Cardinality(), _511_v0)).Cardinality(), _dafny.IntOfInt64(11))).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_518_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(94))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg32 _dafny.Int) interface{} {
							return coer32(arg32)
						}
					}((func(_519_v6 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
						return func(_520_i1 _dafny.Int) _dafny.CodePoint {
							return (_519_v6).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_519_v6).Cardinality()))).Uint32()).(_dafny.CodePoint)
						}
					})(_518_v6))))).Cardinality())) == 0 {
						var _521_v7 _dafny.Map
						_ = _521_v7
						_521_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_511_v0, _this.F12())
						(globalState).F5 = (Companion_Default___.SafeDivisionInt((_521_v7).Cardinality(), _this.F12())).Times(_this.F12())
						(_this).F12_set_(_this.F12())
						var _522_v8 _dafny.Array
						_ = _522_v8
						var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
						_ = _nw90
						_522_v8 = _nw90
						var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_522_v8), 0))
						_ = _index76
						(_522_v8).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_522_v8)).Cardinality()), _this.F12()), (_index76).Int())
						var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_522_v8), 0))
						_ = _index77
						(_522_v8).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F12())), (_index77).Int())
						var _523_v9 _dafny.Sequence
						_ = _523_v9
						_523_v9 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-655))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg33 _dafny.Int) interface{} {
								return coer33(arg33)
							}
						}(func(_524_i2 _dafny.Int) _dafny.CodePoint {
							return (_this).F13()
						})))
						var _525_v10 D2
						_ = _525_v10
						_525_v10 = Companion_D2_.Create_DC8_(_511_v0, (_522_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_522_v8), 0))).Int()).(_dafny.Int), _511_v0)
						var _526_v11 _dafny.Map
						_ = _526_v11
						_526_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_523_v9).Select((Companion_Default___.SafeIndex((_522_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_522_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_523_v9).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _525_v10)
						var _527_v12 _dafny.Sequence
						_ = _527_v12
						_527_v12 = _dafny.SeqOf(_511_v0)
						var _pat_let_tv1 = _511_v0
						_ = _pat_let_tv1
						var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_522_v8), 0))
						_ = _index78
						(_522_v8).ArraySet1(((((_526_v11).Update(_this.F12(), func(_pat_let5_0 D2) D2 {
							return func(_528_dt__update__tmp_h0 D2) D2 {
								return func(_pat_let6_0 bool) D2 {
									return func(_529_dt__update_hcf11_h0 bool) D2 {
										return Companion_D2_.Create_DC8_(_529_dt__update_hcf11_h0, (_528_dt__update__tmp_h0).Dtor_cf12(), (_528_dt__update__tmp_h0).Dtor_cf13())
									}(_pat_let6_0)
								}(_pat_let_tv1)
							}(_pat_let5_0)
						}(_525_v10))).Update(_dafny.IntOfUint32((_518_v6).Cardinality()), _525_v10)).Update((_522_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_522_v8), 0))).Int()).(_dafny.Int), Companion_D2_.Create_DC8_((_527_v12).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_527_v12).Cardinality()))).Uint32()).(bool), (_522_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_522_v8), 0))).Int()).(_dafny.Int), !(!(_511_v0))))).Cardinality(), (_index78).Int())
						var _530_v13 _dafny.Set
						_ = _530_v13
						_530_v13 = _dafny.SetOf((_this).F13(), (_this).F13(), (_518_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.IntOfUint32((_518_v6).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.CodePoint('x'))
						var _531_v14 _dafny.Set
						_ = _531_v14
						_531_v14 = _dafny.SetOf(_511_v0)
						var _532_v15 _dafny.Sequence
						_ = _532_v15
						_532_v15 = _dafny.SeqOf(_531_v14, _531_v14)
						var _533_v16 D6
						_ = _533_v16
						_533_v16 = Companion_D6_.Create_DC15_(_516_v4)
						var _534_v17 _dafny.Map
						_ = _534_v17
						_534_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_530_v13, (func() _dafny.Sequence {
							if _511_v0 {
								return _dafny.SeqOf(((_532_v15).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_532_v15).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), (_dafny.Zero).Minus(Companion_Default___.Fm0((_522_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_522_v8), 0))).Int()).(_dafny.Int), _this.F12(), globalState)))
							}
							return (_533_v16).Dtor_cf21()
						})())
						_534_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_530_v13).Union(_530_v13), _516_v4)
					} else {
						var _535_v18 _dafny.Set
						_ = _535_v18
						_535_v18 = _dafny.SetOf(_515_v3)
						var _536_v19 _dafny.Map
						_ = _536_v19
						_536_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_535_v18, _518_v6)
						_511_v0 = !(!(_511_v0) || (_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("hvvjigs"), (func() _dafny.Sequence {
							if (_536_v19).Contains(_535_v18) {
								return (_536_v19).Get(_535_v18).(_dafny.Sequence)
							}
							return _518_v6
						})())))
						var _537_v20 _dafny.Array
						_ = _537_v20
						var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
						_ = _nw91
						_537_v20 = _nw91
						var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_537_v20), 0))
						_ = _index79
						(_537_v20).ArraySet1(_this.F12(), (_index79).Int())
						var _538_v21 D0
						_ = _538_v21
						_538_v21 = Companion_D0_.Create_DC0_(!(_511_v0))
						var _539_v22 D0
						_ = _539_v22
						_539_v22 = Companion_D0_.Create_DC3_(_538_v21)
						var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_537_v20), 0))
						_ = _index80
						(_537_v20).ArraySet1(_dafny.IntOfInt64(894), (_index80).Int())
						var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_537_v20), 0))
						_ = _index81
						(_537_v20).ArraySet1(_this.F12(), (_index81).Int())
						var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_537_v20), 0))
						_ = _index82
						(_537_v20).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
							if (_513_v1).Contains(_this.F12()) {
								return (_513_v1).Get(_this.F12()).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_518_v6).Cardinality())
						})()), (_index82).Int())
						var _pat_let_tv2 = _538_v21
						_ = _pat_let_tv2
						var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_537_v20), 0))
						_ = _index83
						var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_537_v20), 0))
						_ = _index84
						var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_537_v20), 0))
						_ = _index85
						var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_537_v20), 0))
						_ = _index86
						var _rhs79 _dafny.Int = _this.F12()
						_ = _rhs79
						var _rhs80 D0 = func(_pat_let7_0 D0) D0 {
							return func(_540_dt__update__tmp_h1 D0) D0 {
								return func(_pat_let8_0 D0) D0 {
									return func(_541_dt__update_hcf6_h0 D0) D0 {
										return Companion_D0_.Create_DC3_(_541_dt__update_hcf6_h0)
									}(_pat_let8_0)
								}(Companion_D0_.Create_DC3_(_pat_let_tv2))
							}(_pat_let7_0)
						}(_539_v22)
						_ = _rhs80
						var _rhs81 _dafny.Int = _this.F12()
						_ = _rhs81
						var _rhs82 _dafny.Int = _this.F12()
						_ = _rhs82
						var _rhs83 _dafny.Int = ((func() _dafny.MultiSet {
							if (_511_v0) || (!(_511_v0)) {
								return p0
							}
							return p0
						})()).Cardinality()
						_ = _rhs83
						var _lhs44 _dafny.Array = _537_v20
						_ = _lhs44
						var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_537_v20), 0))
						_ = _lhs45
						var _lhs46 _dafny.Array = _537_v20
						_ = _lhs46
						var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_537_v20), 0))
						_ = _lhs47
						var _lhs48 _dafny.Array = _537_v20
						_ = _lhs48
						var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_537_v20), 0))
						_ = _lhs49
						var _lhs50 _dafny.Array = _537_v20
						_ = _lhs50
						var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_537_v20), 0))
						_ = _lhs51
						(_lhs44).ArraySet1(_rhs79, (_lhs45).Int())
						_539_v22 = _rhs80
						(_lhs46).ArraySet1(_rhs81, (_lhs47).Int())
						(_lhs48).ArraySet1(_rhs82, (_lhs49).Int())
						(_lhs50).ArraySet1(_rhs83, (_lhs51).Int())
						r1 = _511_v0
						var _542_v23 _dafny.Set
						_ = _542_v23
						_542_v23 = _dafny.SetOf(_511_v0, _511_v0)
						(_this).F12_set_((func() _dafny.Int {
							if _511_v0 {
								return (_537_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_537_v20), 0))).Int()).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_516_v4, (_this).Fm5(_511_v0, _511_v0, _511_v0, Companion_Default___.Fm1((func() _dafny.Int {
								if (_513_v1).Contains(_this.F12()) {
									return (_513_v1).Get(_this.F12()).(_dafny.Int)
								}
								return (_542_v23).Cardinality()
							})(), _511_v0, (_537_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_537_v20), 0))).Int()).(_dafny.Int), true, globalState), globalState))).Cardinality())
						})())
						r1 = !(_511_v0)
					}
					_511_v0 = false
					var _543_v24 D5
					_ = _543_v24
					_543_v24 = Companion_D5_.Create_DC14_(_511_v0)
					var _544_v25 _dafny.Set
					_ = _544_v25
					_544_v25 = _dafny.SetOf(_511_v0)
					var _545_v26 _dafny.MultiSet
					_ = _545_v26
					_545_v26 = _dafny.MultiSetOf(_this.F12(), (_544_v25).Cardinality())
					var _pat_let_tv3 = _545_v26
					_ = _pat_let_tv3
					var _pat_let_tv4 = _545_v26
					_ = _pat_let_tv4
					var _source6 D5 = func(_pat_let9_0 D5) D5 {
						return func(_546_dt__update__tmp_h2 D5) D5 {
							return func(_pat_let10_0 bool) D5 {
								return func(_547_dt__update_hcf20_h0 bool) D5 {
									return Companion_D5_.Create_DC14_(_547_dt__update_hcf20_h0)
								}(_pat_let10_0)
							}((_pat_let_tv3).IsSubsetOf(_pat_let_tv4))
						}(_pat_let9_0)
					}(_543_v24)
					_ = _source6
					if _source6.Is_DC13() {
						var _548___mcc_h0 bool = _source6.Get_().(D5_DC13).Cf18
						_ = _548___mcc_h0
						var _549___mcc_h1 bool = _source6.Get_().(D5_DC13).Cf19
						_ = _549___mcc_h1
						var _550_cf19 bool = _549___mcc_h1
						_ = _550_cf19
						var _551_cf18 bool = _548___mcc_h0
						_ = _551_cf18
						var _552_v27 _dafny.Array
						_ = _552_v27
						var _nwElement0_18 _dafny.Int = (_this.F12()).Times((_516_v4).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_516_v4).Cardinality()))).Uint32()).(_dafny.Int))
						_ = _nwElement0_18
						var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.One)
						_ = _nw92
						(_nw92).ArraySet1(_nwElement0_18, 0)
						_552_v27 = _nw92
						var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_552_v27), 0))
						_ = _index87
						(_552_v27).ArraySet1(_this.F12(), (_index87).Int())
						var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_552_v27), 0))
						_ = _index88
						(_552_v27).ArraySet1(_this.F12(), (_index88).Int())
						_518_v6 = _dafny.UnicodeSeqOfUtf8Bytes("nbfyp")
						(globalState).F5 = (_dafny.Zero).Minus(_this.F12())
						var _rhs84 bool = _551_cf18
						_ = _rhs84
						var _rhs85 _dafny.Int = (_dafny.Zero).Minus(_this.F12())
						_ = _rhs85
						r1 = _rhs84
						r2 = _rhs85
					} else if _source6.Is_DC14() {
						var _553___mcc_h2 bool = _source6.Get_().(D5_DC14).Cf20
						_ = _553___mcc_h2
						var _554_cf20 bool = _553___mcc_h2
						_ = _554_cf20
						var _555_v28 _dafny.Sequence
						_ = _555_v28
						_555_v28 = _dafny.SeqOf(_dafny.Companion_Sequence_.Contains(_518_v6, _dafny.CodePoint('n')))
						_511_v0 = (_555_v28).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_555_v28).Cardinality()))).Uint32()).(bool)
						var _556_v29 _dafny.MultiSet
						_ = _556_v29
						_556_v29 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("i"), _dafny.UnicodeSeqOfUtf8Bytes("tf"), _dafny.UnicodeSeqOfUtf8Bytes("tbjwfs"), _518_v6)
						r1 = Companion_Default___.Fm1(_this.F12(), !(Companion_Default___.Fm1(_this.F12(), _554_cf20, _this.F12(), _511_v0, globalState)), ((_556_v29).Intersection(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(653))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg34 _dafny.Int) interface{} {
								return coer34(arg34)
							}
						}(func(_557_i3 _dafny.Int) _dafny.CodePoint {
							return (_this).F13()
						})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-737))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg35 _dafny.Int) interface{} {
								return coer35(arg35)
							}
						}(func(_558_i4 _dafny.Int) _dafny.CodePoint {
							return (_this).F13()
						}))))).Cardinality(), _554_cf20, globalState)
						var _559_v30 _dafny.Map
						_ = _559_v30
						_559_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _511_v0)
						var _560_v31 _dafny.Sequence
						_ = _560_v31
						_560_v31 = _dafny.SeqOf(_544_v25, _544_v25, _544_v25, _544_v25, _dafny.SetOf(true, _554_cf20, _511_v0))
						var _561_v32 _dafny.Map
						_ = _561_v32
						_561_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_560_v31, (_this).F13())
						var _562_v33 _dafny.Array
						_ = _562_v33
						var _nwElement0_19 _dafny.CodePoint = (_this).F13()
						_ = _nwElement0_19
						var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(10))
						_ = _nw93
						(_nw93).ArraySet1CodePoint(_nwElement0_19, 0)
						(_nw93).ArraySet1CodePoint(Companion_Default___.Fm15(_559_v30, _dafny.IntOfUint32((_518_v6).Cardinality()), globalState), 1)
						(_nw93).ArraySet1CodePoint((func() _dafny.CodePoint {
							if (_561_v32).Contains(_560_v31) {
								return (_561_v32).Get(_560_v31).(_dafny.CodePoint)
							}
							return (_this).F13()
						})(), 2)
						(_nw93).ArraySet1CodePoint((_this).F13(), 3)
						(_nw93).ArraySet1CodePoint((_this).F13(), 4)
						(_nw93).ArraySet1CodePoint((_this).F13(), 5)
						(_nw93).ArraySet1CodePoint((_this).F13(), 6)
						(_nw93).ArraySet1CodePoint((_this).F13(), 7)
						(_nw93).ArraySet1CodePoint((_this).F13(), 8)
						(_nw93).ArraySet1CodePoint((_this).F13(), 9)
						_562_v33 = _nw93
						var _rhs86 _dafny.Array = _562_v33
						_ = _rhs86
						var _rhs87 _dafny.Int = ((_dafny.Zero).Minus(_this.F12())).Minus(_this.F12())
						_ = _rhs87
						var _lhs52 *GlobalState = globalState
						_ = _lhs52
						_562_v33 = _rhs86
						_lhs52.F5 = _rhs87
						_511_v0 = _554_cf20
					} else {
						var _563___mcc_h3 _dafny.Sequence = _source6.Get_().(D5_DC12).Cf17
						_ = _563___mcc_h3
						var _564_cf17 _dafny.Sequence = _563___mcc_h3
						_ = _564_cf17
						_518_v6 = _dafny.Companion_Sequence_.Update(_518_v6, (Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_518_v6).Cardinality()))).Uint32(), (_this).F13())
						var _565_v34 _dafny.Array
						_ = _565_v34
						var _len0_14 _dafny.Int = _dafny.One
						_ = _len0_14
						var _nw94 _dafny.Array
						_ = _nw94
						if _len0_14.Cmp(_dafny.Zero) == 0 {
							_nw94 = _dafny.NewArray(_len0_14)
						} else {
							var _init14 func(_dafny.Int) _dafny.Int = func(_566_i5 _dafny.Int) _dafny.Int {
								return (_566_i5).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(491))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg36 _dafny.Int) interface{} {
										return coer36(arg36)
									}
								}(func(_567_i6 _dafny.Int) _dafny.CodePoint {
									return (_this).F13()
								}))).Cardinality())))
							}
							_ = _init14
							var _element0_14 = _init14(_dafny.Zero)
							_ = _element0_14
							_nw94 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
							(_nw94).ArraySet1(_element0_14, 0)
							var _nativeLen0_14 = (_len0_14).Int()
							_ = _nativeLen0_14
							for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
								(_nw94).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
							}
						}
						_565_v34 = _nw94
						_565_v34 = (func() _dafny.Array {
							if (_511_v0) == (_511_v0) {
								return _565_v34
							}
							return (func() _dafny.Array {
								if _511_v0 {
									return _565_v34
								}
								return _565_v34
							})()
						})()
						r1 = !(!(_511_v0))
						var _568_v35 *C1
						_ = _568_v35
						var _nw95 *C1 = New_C1_()
						_ = _nw95
						_nw95.Ctor__((_this).F13(), (_this).F13(), _this.F12())
						_568_v35 = _nw95
					}
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _569_v36 _dafny.Map
		_ = _569_v36
		_569_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_511_v0, _511_v0)
		var _570_v37 _dafny.Sequence
		_ = _570_v37
		_570_v37 = _dafny.UnicodeSeqOfUtf8Bytes("gpstgyprm")
		var _571_v38 _dafny.Sequence
		_ = _571_v38
		_571_v38 = _dafny.SeqOf(Companion_Default___.Fm1((_569_v36).Cardinality(), _511_v0, _dafny.IntOfUint32((_570_v37).Cardinality()), _511_v0, globalState))
		var _572_v39 _dafny.Set
		_ = _572_v39
		_572_v39 = _dafny.SetOf(Companion_Default___.Fm1(_this.F12(), false, _dafny.IntOfUint32((_571_v38).Cardinality()), _511_v0, globalState), _511_v0, (func() bool {
			if (_569_v36).Contains(_511_v0) {
				return (_569_v36).Get(_511_v0).(bool)
			}
			return _511_v0
		})(), _511_v0)
		var _source7 _dafny.Set = _572_v39
		_ = _source7
		var _573___mcc_h4 _dafny.Set = _source7
		_ = _573___mcc_h4
		var _574_cf16 _dafny.Set = _573___mcc_h4
		_ = _574_cf16
		var _575_v40 _dafny.Array
		_ = _575_v40
		var _nwElement0_20 bool = !(Companion_Default___.Fm1(_dafny.IntOfInt64(783), _511_v0, _this.F12(), _511_v0, globalState))
		_ = _nwElement0_20
		var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(8))
		_ = _nw96
		(_nw96).ArraySet1(_nwElement0_20, 0)
		(_nw96).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_570_v37, _570_v37), 1)
		(_nw96).ArraySet1(true, 2)
		(_nw96).ArraySet1(_511_v0, 3)
		(_nw96).ArraySet1(_511_v0, 4)
		(_nw96).ArraySet1((_this.F12()).Cmp(_this.F12()) > 0, 5)
		(_nw96).ArraySet1(!((!(_511_v0)) || (_511_v0)), 6)
		(_nw96).ArraySet1(_511_v0, 7)
		_575_v40 = _nw96
		var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_575_v40), 0))
		_ = _index89
		(_575_v40).ArraySet1(_511_v0, (_index89).Int())
		var _576_v41 _dafny.Sequence
		_ = _576_v41
		_576_v41 = _dafny.SeqOf(_dafny.IntOfUint32((_570_v37).Cardinality()))
		var _577_v42 _dafny.Map
		_ = _577_v42
		_577_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_511_v0, (_this).F13())
		var _578_v44 _dafny.Set
		_ = _578_v44
		_578_v44 = _dafny.SetOf((_this).F13())
		var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_575_v40), 0))
		_ = _index90
		(_575_v40).ArraySet1((func() bool {
			if _511_v0 {
				return (_dafny.IntOfUint32((_576_v41).Cardinality())).Cmp(_this.F12()) == 0
			}
			return !(func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_578_v44).Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _579_v43 _dafny.CodePoint
					_579_v43 = interface{}(_compr_13).(_dafny.CodePoint)
					if (_578_v44).Contains(_579_v43) {
						_coll13.Add(_579_v43, _511_v0)
					}
				}
				return _coll13.ToMap()
			}()).Contains((func() _dafny.CodePoint {
				if (_577_v42).Contains(_511_v0) {
					return (_577_v42).Get(_511_v0).(_dafny.CodePoint)
				}
				return (_this).F13()
			})())
		})(), (_index90).Int())
		var _580_v46 _dafny.Map
		_ = _580_v46
		_580_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F12())
		var _581_v47 _dafny.Map
		_ = _581_v47
		_581_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), (_580_v46).Cardinality())
		var _582_v48 _dafny.Map
		_ = _582_v48
		_582_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_581_v47).Cardinality(), _511_v0)
		var _583_v49 *C0
		_ = _583_v49
		var _nw97 *C0 = New_C0_()
		_ = _nw97
		_nw97.Ctor__((func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate((_576_v41).Elements()); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _584_v45 _dafny.Int
				_584_v45 = interface{}(_compr_14).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_576_v41, _584_v45) {
					_coll14.Add((_584_v45).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-485))), !((_575_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_575_v40), 0))).Int()).(bool)))
				}
			}
			return _coll14.ToMap()
		}()).Merge(_582_v48), _this.F12())
		_583_v49 = _nw97
		var _585_v50 _dafny.Array
		_ = _585_v50
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_15
		var _nw98 _dafny.Array
		_ = _nw98
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw98 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) _dafny.Int = func(_586_i7 _dafny.Int) _dafny.Int {
				return (_586_i7).Minus(_this.F12())
			}
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw98 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw98).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw98).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_585_v50 = _nw98
		var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_585_v50), 0))
		_ = _index91
		(_585_v50).ArraySet1(_this.F12(), (_index91).Int())
		var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_585_v50), 0))
		_ = _index92
		(_585_v50).ArraySet1((_this.F12()).Plus((_574_cf16).Cardinality()), (_index92).Int())
		(_583_v49).F16 = (_583_v49.F16).Update(_dafny.IntOfInt64(437), !_dafny.Companion_Sequence_.Contains((_this).Fm5(!(_511_v0), _511_v0, _511_v0, _511_v0, globalState), _this.F12()))
		var _587_v51 _dafny.Array
		_ = _587_v51
		var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw99
		_587_v51 = _nw99
		var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))
		_ = _index93
		(_587_v51).ArraySet1(_511_v0, (_index93).Int())
		var _588_v52 _dafny.Sequence
		_ = _588_v52
		_588_v52 = _dafny.SeqOf(p0, p0)
		var _589_v53 D5
		_ = _589_v53
		_589_v53 = Companion_D5_.Create_DC12_(_588_v52)
		var _590_v54 _dafny.Sequence
		_ = _590_v54
		_590_v54 = _dafny.SeqOf(_589_v53)
		var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))
		_ = _index94
		(_587_v51).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_589_v53), _dafny.Companion_Sequence_.Update(_590_v54, (Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_590_v54).Cardinality()))).Uint32(), _589_v53)), _590_v54), (_index94).Int())
		var _591_v55 _dafny.MultiSet
		_ = _591_v55
		_591_v55 = _dafny.MultiSetOf(_571_v38, _dafny.SeqOf((_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool)), _571_v38)
		var _592_v56 _dafny.Set
		_ = _592_v56
		_592_v56 = _dafny.SetOf(_this.F12(), _this.F12())
		var _593_v57 _dafny.Sequence
		_ = _593_v57
		_593_v57 = _dafny.SeqOf(_this.F12(), _this.F12(), (_592_v56).Cardinality(), _this.F12())
		var _594_v58 D6
		_ = _594_v58
		_594_v58 = Companion_D6_.Create_DC15_(_593_v57)
		_511_v0 = ((_591_v55).Intersection(_591_v55)).IsSubsetOf((_591_v55).Union(Companion_Default___.Fm16(_511_v0, _this.F12(), _594_v58, _this.F12(), globalState)))
		if !_dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
			if true {
				return _570_v37
			}
			return _570_v37
		})(), _570_v37) {
			var _595_v59 D2
			_ = _595_v59
			_595_v59 = Companion_D2_.Create_DC8_((_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool), _this.F12(), _511_v0)
			var _pat_let_tv5 = _511_v0
			_ = _pat_let_tv5
			var _596_v60 _dafny.Array
			_ = _596_v60
			var _nwElement0_21 D2 = _595_v59
			_ = _nwElement0_21
			var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(29))
			_ = _nw100
			(_nw100).ArraySet1(_nwElement0_21, 0)
			(_nw100).ArraySet1(_595_v59, 1)
			(_nw100).ArraySet1(_595_v59, 2)
			(_nw100).ArraySet1(_595_v59, 3)
			(_nw100).ArraySet1(_595_v59, 4)
			(_nw100).ArraySet1(Companion_Default___.Fm17(_511_v0, Companion_Default___.Fm1(_this.F12(), _511_v0, _this.F12(), _511_v0, globalState), globalState), 5)
			(_nw100).ArraySet1(_595_v59, 6)
			(_nw100).ArraySet1(_595_v59, 7)
			(_nw100).ArraySet1(_595_v59, 8)
			(_nw100).ArraySet1(func(_pat_let11_0 D2) D2 {
				return func(_597_dt__update__tmp_h3 D2) D2 {
					return func(_pat_let12_0 bool) D2 {
						return func(_598_dt__update_hcf13_h0 bool) D2 {
							return Companion_D2_.Create_DC8_((_597_dt__update__tmp_h3).Dtor_cf11(), (_597_dt__update__tmp_h3).Dtor_cf12(), _598_dt__update_hcf13_h0)
						}(_pat_let12_0)
					}(_pat_let_tv5)
				}(_pat_let11_0)
			}(_595_v59), 9)
			(_nw100).ArraySet1(_595_v59, 10)
			(_nw100).ArraySet1(_595_v59, 11)
			(_nw100).ArraySet1(_595_v59, 12)
			(_nw100).ArraySet1(Companion_D2_.Create_DC8_(false, _dafny.IntOfUint32((_571_v38).Cardinality()), _511_v0), 13)
			(_nw100).ArraySet1(_595_v59, 14)
			(_nw100).ArraySet1(_595_v59, 15)
			(_nw100).ArraySet1(_595_v59, 16)
			(_nw100).ArraySet1(_595_v59, 17)
			(_nw100).ArraySet1(_595_v59, 18)
			(_nw100).ArraySet1(_595_v59, 19)
			(_nw100).ArraySet1(_595_v59, 20)
			(_nw100).ArraySet1(_595_v59, 21)
			(_nw100).ArraySet1(_595_v59, 22)
			(_nw100).ArraySet1(_595_v59, 23)
			(_nw100).ArraySet1(_595_v59, 24)
			(_nw100).ArraySet1(Companion_Default___.Fm17(_511_v0, _511_v0, globalState), 25)
			(_nw100).ArraySet1(_595_v59, 26)
			(_nw100).ArraySet1(_595_v59, 27)
			(_nw100).ArraySet1(_595_v59, 28)
			_596_v60 = _nw100
			var _599_v61 _dafny.Sequence
			_ = _599_v61
			_599_v61 = _dafny.SeqOf(_596_v60, _596_v60)
			_599_v61 = (func() _dafny.Sequence {
				if (_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool) {
					return _599_v61
				}
				return _599_v61
			})()
			var _600_v62 _dafny.Int
			_ = _600_v62
			var _601_v63 _dafny.Map
			_ = _601_v63
			var _602_v64 _dafny.Sequence
			_ = _602_v64
			var _out39 _dafny.Int
			_ = _out39
			var _out40 _dafny.Map
			_ = _out40
			var _out41 _dafny.Sequence
			_ = _out41
			_out39, _out40, _out41 = Companion_Default___.M0(_this.F12(), (_dafny.IntOfInt64(113)).Minus((_dafny.Zero).Minus(_this.F12())), _this.F12(), (_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool), globalState)
			_600_v62 = _out39
			_601_v63 = _out40
			_602_v64 = _out41
			_511_v0 = (_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool)
			var _603_v65 _dafny.Map
			_ = _603_v65
			_603_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(_511_v0)), _511_v0)
			_603_v65 = (_603_v65).Update(_571_v38, (_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool))
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))
			_ = _index95
			(_587_v51).ArraySet1(_511_v0, (_index95).Int())
		} else {
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))
			_ = _index96
			(_587_v51).ArraySet1((_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool), (_index96).Int())
			var _604_v66 _dafny.Array
			_ = _604_v66
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_16
			var _nw101 _dafny.Array
			_ = _nw101
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw101 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Map = func(_605_i8 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), (_dafny.Zero).Minus(_this.F12()))
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
			_604_v66 = _nw101
			var _606_v67 _dafny.Map
			_ = _606_v67
			_606_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), false)
			var _607_v68 _dafny.Sequence
			_ = _607_v68
			_607_v68 = _dafny.SeqOf(_571_v38, _dafny.SeqOf((_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool), false, false, _511_v0))
			var _608_v69 _dafny.Map
			_ = _608_v69
			_608_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_606_v67).Cardinality(), (Companion_D2_.Create_DC9_(Companion_Default___.Fm0(_this.F12(), (_dafny.MultiSetFromSeq((_607_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.IntOfUint32((_607_v68).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality(), globalState))).Dtor_cf14())
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_604_v66), 0))
			_ = _index97
			(_604_v66).ArraySet1((_608_v69).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F12())), (_index97).Int())
			var _609_v71 _dafny.MultiSet
			_ = _609_v71
			_609_v71 = _dafny.MultiSetOf((_this).F13())
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_604_v66), 0))
			_ = _index98
			var _rhs88 _dafny.Map = func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate((_593_v57).Elements()); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _610_v70 _dafny.Int
					_610_v70 = interface{}(_compr_15).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_593_v57, _610_v70) {
						_coll15.Add((_610_v70).Times(_dafny.IntOfInt64(901)), _this.F12())
					}
				}
				return _coll15.ToMap()
			}()
			_ = _rhs88
			var _rhs89 _dafny.Sequence = _570_v37
			_ = _rhs89
			var _rhs90 bool = (_dafny.MultiSetFromSeq(Companion_Default___.Fm11(_this.F12(), (_dafny.MultiSetFromSeq(_593_v57)).Cardinality(), (_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool), globalState))).IsSubsetOf((_dafny.MultiSetOf((_this).F13(), _dafny.CodePoint('f'), (_this).F13())).Difference(_609_v71))
			_ = _rhs90
			var _rhs91 _dafny.Sequence = _570_v37
			_ = _rhs91
			var _lhs53 _dafny.Array = _604_v66
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_604_v66), 0))
			_ = _lhs54
			(_lhs53).ArraySet1(_rhs88, (_lhs54).Int())
			_570_v37 = _rhs89
			_511_v0 = _rhs90
			_570_v37 = _rhs91
			var _611_v72 _dafny.Array
			_ = _611_v72
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw102
			_611_v72 = _nw102
			var _612_v73 _dafny.Map
			_ = _612_v73
			_612_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_this.F12(), false, _this.F12(), false, globalState), _611_v72)
			_612_v73 = (_612_v73).Update(_511_v0, (func() _dafny.Array {
				if (_612_v73).Contains(_511_v0) {
					return (_612_v73).Get(_511_v0).(_dafny.Array)
				}
				return _611_v72
			})())
			var _613_v74 _dafny.CodePoint
			_ = _613_v74
			_613_v74 = _dafny.CodePoint('o')
			_613_v74 = Companion_Default___.Fm15(_569_v36, (_593_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_570_v37).Cardinality()), _dafny.IntOfUint32((_593_v57).Cardinality()))).Uint32()).(_dafny.Int), globalState)
			_592_v56 = _592_v56
		}
		var _614_i9 _dafny.Int
		_ = _614_i9
		_614_i9 = _dafny.Zero
		{
			for (_571_v38).Select((Companion_Default___.SafeIndex(((_dafny.Zero).Minus(_this.F12())).Plus(_this.F12()), _dafny.IntOfUint32((_571_v38).Cardinality()))).Uint32()).(bool) {
				{
					if (_614_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_614_i9 = (_614_i9).Plus(_dafny.One)
					(globalState).F5 = _dafny.IntOfInt64(479)
					var _615_v75 _dafny.Map
					_ = _615_v75
					_615_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), (_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool))
					_615_v75 = (_615_v75).Update(_this.F12(), !((_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool)) || ((_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool)))
					var _616_v76 _dafny.CodePoint
					_ = _616_v76
					_616_v76 = _dafny.CodePoint('u')
					_616_v76 = _616_v76
					_511_v0 = (func() bool {
						if (_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool) {
							return !(_dafny.SetOf(_511_v0, _511_v0)).Contains((_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool))
						}
						return (_this.F12()).Cmp((_dafny.Zero).Minus(_this.F12())) == 0
					})()
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		r0 = (_dafny.IntOfInt64(832)).Times(_this.F12())
		r1 = (_587_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_587_v51), 0))).Int()).(bool)
		var _617_v77 _dafny.MultiSet
		_ = _617_v77
		_617_v77 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_570_v37, (Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_570_v37).Cardinality()))).Uint32(), (_this).F13()), _570_v37)
		r2 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_617_v77).Contains(_dafny.UnicodeSeqOfUtf8Bytes("ptyrk")) {
				return (_617_v77).Multiplicity(_dafny.UnicodeSeqOfUtf8Bytes("ptyrk"))
			}
			return _this.F12()
		})(), (_this.F12()).Times(_this.F12()))
		return r0, r1, r2
	}
}
func (_this *C3) M4(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _hi2 _dafny.Int = (_this.F12()).Minus(_this.F12())
		_ = _hi2
		for _618_i0 := p0; _618_i0.Cmp(_hi2) < 0; _618_i0 = _618_i0.Plus(_dafny.One) {
			var _619_v0 _dafny.Map
			_ = _619_v0
			_619_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F12())
			var _620_v1 _dafny.MultiSet
			_ = _620_v1
			_620_v1 = _dafny.MultiSetOf(_619_v0, _619_v0)
			_620_v1 = _dafny.MultiSetFromSeq(_dafny.SeqOf(_619_v0))
			var _621_v2 _dafny.Array
			_ = _621_v2
			var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw103
			_621_v2 = _nw103
			var _622_v3 D2
			_ = _622_v3
			_622_v3 = Companion_D2_.Create_DC8_(p1, _dafny.IntOfInt64(-238), p1)
			var _623_v4 _dafny.Set
			_ = _623_v4
			_623_v4 = _dafny.SetOf((_622_v3).Dtor_cf11())
			var _624_v5 _dafny.Map
			_ = _624_v5
			_624_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _625_v6 T0
			_ = _625_v6
			var _nw104 *C2 = New_C2_()
			_ = _nw104
			_nw104.Ctor__(true, (_dafny.Zero).Minus(_618_i0))
			_625_v6 = _nw104
			var _626_v7 _dafny.Sequence
			_ = _626_v7
			_626_v7 = _dafny.SeqOf(p0, p0, p0, p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _625_v6)).Cardinality())
			var _627_v8 _dafny.Map
			_ = _627_v8
			_627_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D2_.Create_DC8_(p1, p0, p1)).Dtor_cf12(), _this.F12())
			var _628_v9 _dafny.Array
			_ = _628_v9
			var _nwElement0_22 _dafny.Int = _this.F12()
			_ = _nwElement0_22
			var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(25))
			_ = _nw105
			(_nw105).ArraySet1(_nwElement0_22, 0)
			(_nw105).ArraySet1(p0, 1)
			(_nw105).ArraySet1((Companion_Default___.Fm18(_dafny.IntOfUint32((p2).Cardinality()), _dafny.IntOfInt64(421), p1, globalState)).Cardinality(), 2)
			(_nw105).ArraySet1(_this.F12(), 3)
			(_nw105).ArraySet1(_this.F12(), 4)
			(_nw105).ArraySet1(p0, 5)
			(_nw105).ArraySet1((_623_v4).Cardinality(), 6)
			(_nw105).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 7)
			(_nw105).ArraySet1(_this.F12(), 8)
			(_nw105).ArraySet1(_this.F12(), 9)
			(_nw105).ArraySet1(_618_i0, 10)
			(_nw105).ArraySet1(_dafny.IntOfInt64(203), 11)
			(_nw105).ArraySet1(p0, 12)
			(_nw105).ArraySet1(_618_i0, 13)
			(_nw105).ArraySet1(p0, 14)
			(_nw105).ArraySet1(_618_i0, 15)
			(_nw105).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((p2).Cardinality())), 16)
			(_nw105).ArraySet1(_dafny.IntOfInt64(968), 17)
			(_nw105).ArraySet1(_this.F12(), 18)
			(_nw105).ArraySet1(_618_i0, 19)
			(_nw105).ArraySet1(((_624_v5).Update(_dafny.IntOfUint32((_626_v7).Cardinality()), p1)).Cardinality(), 20)
			(_nw105).ArraySet1(p0, 21)
			(_nw105).ArraySet1(p0, 22)
			(_nw105).ArraySet1((_627_v8).Cardinality(), 23)
			(_nw105).ArraySet1(_618_i0, 24)
			_628_v9 = _nw105
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_621_v2), 0))
			_ = _index99
			(_621_v2).ArraySet1(_628_v9, (_index99).Int())
			var _629_v10 D6
			_ = _629_v10
			_629_v10 = Companion_D6_.Create_DC15_((func() _dafny.Sequence {
				if p1 {
					return _626_v7
				}
				return _626_v7
			})())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_621_v2), 0))
			_ = _index100
			var _rhs92 _dafny.Int = _625_v6.F12()
			_ = _rhs92
			var _rhs93 _dafny.Array = _628_v9
			_ = _rhs93
			var _rhs94 bool = p1
			_ = _rhs94
			var _rhs95 bool = p1
			_ = _rhs95
			var _rhs96 D6 = Companion_D6_.Create_DC15_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(684))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}((func(_630_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_631_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_630_p2).Cardinality())
				}
			})(p2))))
			_ = _rhs96
			var _lhs55 _dafny.Array = _621_v2
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_621_v2), 0))
			_ = _lhs56
			r2 = _rhs92
			(_lhs55).ArraySet1(_rhs93, (_lhs56).Int())
			r1 = _rhs94
			r1 = _rhs95
			_629_v10 = _rhs96
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_621_v2), 0))
			_ = _index101
			var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw106
			(_621_v2).ArraySet1(_nw106, (_index101).Int())
			var _632_v11 _dafny.Sequence
			_ = _632_v11
			_632_v11 = _dafny.SeqOf(p1, p1)
			var _633_v12 _dafny.Array
			_ = _633_v12
			var _nwElement0_23 bool = p1
			_ = _nwElement0_23
			var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(16))
			_ = _nw107
			(_nw107).ArraySet1(_nwElement0_23, 0)
			(_nw107).ArraySet1(Companion_Default___.Fm1(p0, p1, _dafny.IntOfInt64(-195), p1, globalState), 1)
			(_nw107).ArraySet1(p1, 2)
			(_nw107).ArraySet1(p1, 3)
			(_nw107).ArraySet1(p1, 4)
			(_nw107).ArraySet1(p1, 5)
			(_nw107).ArraySet1(p1, 6)
			(_nw107).ArraySet1(p1, 7)
			(_nw107).ArraySet1(false, 8)
			(_nw107).ArraySet1(p1, 9)
			(_nw107).ArraySet1((_632_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_632_v11).Cardinality()))).Uint32()).(bool), 10)
			(_nw107).ArraySet1(p1, 11)
			(_nw107).ArraySet1(p1, 12)
			(_nw107).ArraySet1(p1, 13)
			(_nw107).ArraySet1(p1, 14)
			(_nw107).ArraySet1(p1, 15)
			_633_v12 = _nw107
			var _634_v13 _dafny.Map
			_ = _634_v13
			_634_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_618_i0, _633_v12)
			var _635_v14 _dafny.Sequence
			_ = _635_v14
			_635_v14 = _dafny.SeqOf(_633_v12)
			_634_v13 = (_634_v13).Update(_this.F12(), (_635_v14).Select((Companion_Default___.SafeIndex((_626_v7).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_626_v7).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_635_v14).Cardinality()))).Uint32()).(_dafny.Array))
		}
		if p1 {
			var _636_v15 _dafny.Set
			_ = _636_v15
			_636_v15 = _dafny.SetOf(_this.F12())
			var _637_v16 _dafny.Map
			_ = _637_v16
			_637_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F12())
			var _638_v17 _dafny.Map
			_ = _638_v17
			_638_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), p1)
			var _639_v18 T2
			_ = _639_v18
			var _nw108 *C0 = New_C0_()
			_ = _nw108
			_nw108.Ctor__(_638_v17, _this.F12())
			_639_v18 = _nw108
			var _640_v19 _dafny.Sequence
			_ = _640_v19
			_640_v19 = _dafny.SeqOf(_639_v18)
			var _641_v20 _dafny.MultiSet
			_ = _641_v20
			_641_v20 = _dafny.MultiSetOf(false)
			var _642_v21 _dafny.Array
			_ = _642_v21
			var _nwElement0_24 _dafny.Int = p0
			_ = _nwElement0_24
			var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(29))
			_ = _nw109
			(_nw109).ArraySet1(_nwElement0_24, 0)
			(_nw109).ArraySet1(_this.F12(), 1)
			(_nw109).ArraySet1((_dafny.Zero).Minus(_this.F12()), 2)
			(_nw109).ArraySet1(_this.F12(), 3)
			(_nw109).ArraySet1((_636_v15).Cardinality(), 4)
			(_nw109).ArraySet1(p0, 5)
			(_nw109).ArraySet1(p0, 6)
			(_nw109).ArraySet1(p0, 7)
			(_nw109).ArraySet1(_this.F12(), 8)
			(_nw109).ArraySet1(_this.F12(), 9)
			(_nw109).ArraySet1(p0, 10)
			(_nw109).ArraySet1(_this.F12(), 11)
			(_nw109).ArraySet1((func() _dafny.Int {
				if (_637_v16).Contains(p1) {
					return (_637_v16).Get(p1).(_dafny.Int)
				}
				return p0
			})(), 12)
			(_nw109).ArraySet1(p0, 13)
			(_nw109).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_640_v19, (Companion_Default___.SafeIndex((_641_v20).Cardinality(), _dafny.IntOfUint32((_640_v19).Cardinality()))).Uint32(), _639_v18)).Cardinality()), 14)
			(_nw109).ArraySet1(p0, 15)
			(_nw109).ArraySet1(p0, 16)
			(_nw109).ArraySet1((_dafny.Zero).Minus(p0), 17)
			(_nw109).ArraySet1(_639_v18.F12(), 18)
			(_nw109).ArraySet1(_this.F12(), 19)
			(_nw109).ArraySet1(p0, 20)
			(_nw109).ArraySet1(_639_v18.F12(), 21)
			(_nw109).ArraySet1(_this.F12(), 22)
			(_nw109).ArraySet1((_dafny.Zero).Minus(p0), 23)
			(_nw109).ArraySet1(_this.F12(), 24)
			(_nw109).ArraySet1(p0, 25)
			(_nw109).ArraySet1(p0, 26)
			(_nw109).ArraySet1(_this.F12(), 27)
			(_nw109).ArraySet1(_639_v18.F12(), 28)
			_642_v21 = _nw109
			var _643_v22 D1
			_ = _643_v22
			_643_v22 = Companion_D1_.Create_DC4_(_642_v21)
			var _644_v23 _dafny.Map
			_ = _644_v23
			_644_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_643_v22).Dtor_cf7()), (func() bool {
				if p1 {
					return p1
				}
				return p1
			})())
			var _645_v24 _dafny.Map
			_ = _645_v24
			_645_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _642_v21)
			var _646_v25 D7
			_ = _646_v25
			_646_v25 = Companion_D7_.Create_DC20_(_645_v24)
			var _647_v26 _dafny.Sequence
			_ = _647_v26
			_647_v26 = _dafny.SeqOf(!(p1))
			_644_v23 = (_644_v23).Update((_646_v25).Dtor_cf27(), !(Companion_Default___.Fm1((_dafny.MultiSetFromSeq(_647_v26)).Cardinality(), false, p0, p1, globalState)))
			r2 = _this.F12()
			r0 = p1
			var _648_v27 _dafny.CodePoint
			_ = _648_v27
			_648_v27 = _dafny.CodePoint('b')
			_648_v27 = (_this).F13()
			var _649_v28 _dafny.Array
			_ = _649_v28
			var _nw110 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw110
			_649_v28 = _nw110
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_649_v28), 0))
			_ = _index102
			(_649_v28).ArraySet1(p1, (_index102).Int())
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_649_v28), 0))
			_ = _index103
			(_649_v28).ArraySet1(p1, (_index103).Int())
		} else {
			var _650_v29 _dafny.Sequence
			_ = _650_v29
			_650_v29 = _dafny.UnicodeSeqOfUtf8Bytes("cggpcap")
			_650_v29 = _650_v29
			var _651_v30 *C0
			_ = _651_v30
			var _nw111 *C0 = New_C0_()
			_ = _nw111
			_nw111.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), true)).Update(_dafny.IntOfInt64(25), p1), _this.F12())
			_651_v30 = _nw111
			_651_v30 = _651_v30
			if p1 {
				var _652_v31 _dafny.CodePoint
				_ = _652_v31
				_652_v31 = _dafny.CodePoint('e')
				_652_v31 = (_this).F13()
				var _653_v32 _dafny.Sequence
				_ = _653_v32
				_653_v32 = _dafny.SeqOf(!(p1))
				var _654_v33 _dafny.Map
				_ = _654_v33
				_654_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_653_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.IntOfUint32((_653_v32).Cardinality()))).Uint32()).(bool), p1)
				var _655_v34 _dafny.Set
				_ = _655_v34
				_655_v34 = _dafny.SetOf(p1, false, p1)
				var _656_v35 _dafny.Array
				_ = _656_v35
				var _nwElement0_25 bool = p1
				_ = _nwElement0_25
				var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(22))
				_ = _nw112
				(_nw112).ArraySet1(_nwElement0_25, 0)
				(_nw112).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Equals(_654_v33), 1)
				(_nw112).ArraySet1(Companion_Default___.Fm1(_this.F12(), p1, p0, p1, globalState), 2)
				(_nw112).ArraySet1(p1, 3)
				(_nw112).ArraySet1(p1, 4)
				(_nw112).ArraySet1((p0).Cmp(p0) >= 0, 5)
				(_nw112).ArraySet1(!(p1) || (p1), 6)
				(_nw112).ArraySet1(p1, 7)
				(_nw112).ArraySet1(p1, 8)
				(_nw112).ArraySet1(p1, 9)
				(_nw112).ArraySet1(p1, 10)
				(_nw112).ArraySet1(p1, 11)
				(_nw112).ArraySet1(p1, 12)
				(_nw112).ArraySet1(p1, 13)
				(_nw112).ArraySet1(p1, 14)
				(_nw112).ArraySet1((func() bool {
					if p1 {
						return p1
					}
					return p1
				})(), 15)
				(_nw112).ArraySet1(p1, 16)
				(_nw112).ArraySet1((_655_v34).IsDisjointFrom(_655_v34), 17)
				(_nw112).ArraySet1(true, 18)
				(_nw112).ArraySet1(p1, 19)
				(_nw112).ArraySet1(p1, 20)
				(_nw112).ArraySet1((func() bool {
					if p1 {
						return p1
					}
					return p1
				})(), 21)
				_656_v35 = _nw112
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_656_v35), 0))
				_ = _index104
				(_656_v35).ArraySet1(p1, (_index104).Int())
				var _657_v36 _dafny.Sequence
				_ = _657_v36
				_657_v36 = _dafny.SeqOf(_dafny.IntOfInt64(-803))
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_656_v35), 0))
				_ = _index105
				(_656_v35).ArraySet1(((_dafny.SetOf(_656_v35)).Cardinality()).Cmp(_dafny.IntOfUint32((_657_v36).Cardinality())) != 0, (_index105).Int())
				var _658_v37 _dafny.Array
				_ = _658_v37
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_17
				var _nw113 _dafny.Array
				_ = _nw113
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw113 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Sequence = (func(_659_v29 _dafny.Sequence, _660_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_661_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_659_v29, _660_p2)
						}
					})(_650_v29, p2)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw113 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw113).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw113).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_658_v37 = _nw113
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_658_v37), 0))
				_ = _index106
				(_658_v37).ArraySet1(p2, (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_658_v37), 0))
				_ = _index107
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_656_v35), 0))
				_ = _index108
				var _rhs97 bool = p1
				_ = _rhs97
				var _rhs98 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p2, _dafny.UnicodeSeqOfUtf8Bytes("b"))
				_ = _rhs98
				var _rhs99 _dafny.CodePoint = _dafny.CodePoint('d')
				_ = _rhs99
				var _rhs100 bool = p1
				_ = _rhs100
				var _lhs57 _dafny.Array = _658_v37
				_ = _lhs57
				var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_658_v37), 0))
				_ = _lhs58
				var _lhs59 _dafny.Array = _656_v35
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_656_v35), 0))
				_ = _lhs60
				r0 = _rhs97
				(_lhs57).ArraySet1(_rhs98, (_lhs58).Int())
				_652_v31 = _rhs99
				(_lhs59).ArraySet1(_rhs100, (_lhs60).Int())
				var _662_v40 *C0
				_ = _662_v40
				var _nw114 *C0 = New_C0_()
				_ = _nw114
				_nw114.Ctor__((_651_v30.F16).Merge(func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter16 := _dafny.Iterate((func() _dafny.Set {
						var _coll17 = _dafny.NewBuilder()
						_ = _coll17
						for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(307), _dafny.IntOfInt64(123))); ; {
							_compr_17, _ok17 := _iter17()
							if !_ok17 {
								break
							}
							var _663_v39 _dafny.Int
							_663_v39 = interface{}(_compr_17).(_dafny.Int)
							if ((_dafny.IntOfInt64(307)).Cmp(_663_v39) <= 0) && ((_663_v39).Cmp(_dafny.IntOfInt64(123)) < 0) {
								_coll17.Add(Companion_Default___.SafeDivisionInt(_663_v39, _this.F12()))
							}
						}
						return _coll17.ToSet()
					}()).Elements()); ; {
						_compr_16, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _664_v38 _dafny.Int
						_664_v38 = interface{}(_compr_16).(_dafny.Int)
						if (func() _dafny.Set {
							var _coll18 = _dafny.NewBuilder()
							_ = _coll18
							for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(307), _dafny.IntOfInt64(123))); ; {
								_compr_18, _ok18 := _iter18()
								if !_ok18 {
									break
								}
								var _665_v39 _dafny.Int
								_665_v39 = interface{}(_compr_18).(_dafny.Int)
								if ((_dafny.IntOfInt64(307)).Cmp(_665_v39) <= 0) && ((_665_v39).Cmp(_dafny.IntOfInt64(123)) < 0) {
									_coll18.Add(Companion_Default___.SafeDivisionInt(_665_v39, _this.F12()))
								}
							}
							return _coll18.ToSet()
						}()).Contains(_664_v38) {
							_coll16.Add((_664_v38).Minus(p0), (_656_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_656_v35), 0))).Int()).(bool))
						}
					}
					return _coll16.ToMap()
				}()), _this.F12())
				_662_v40 = _nw114
				_657_v36 = _dafny.SeqOf(_this.F12(), Companion_Default___.SafeModuloInt(_this.F12(), _dafny.IntOfInt64(476)), p0)
			} else {
				var _666_v41 _dafny.Map
				_ = _666_v41
				_666_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
				var _667_v42 _dafny.MultiSet
				_ = _667_v42
				_667_v42 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(653))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}(func(_668_i3 _dafny.Int) _dafny.CodePoint {
					return (_this).F13()
				}))).Cardinality()), (_666_v41).Cardinality())
				var _669_v43 _dafny.Sequence
				_ = _669_v43
				_669_v43 = _dafny.SeqOf((_dafny.Zero).Minus(((_dafny.MultiSetOf(_651_v30.F16)).Update(_651_v30.F16, Companion_Default___.Abs((_dafny.Zero).Minus(p0)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(p1, p1)).Cardinality()))
				var _670_v44 _dafny.Set
				_ = _670_v44
				_670_v44 = _dafny.SetOf(_669_v43)
				var _671_v45 _dafny.Map
				_ = _671_v45
				_671_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_667_v42, (_670_v44).Cardinality())
				_671_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_667_v42, _this.F12())
				(globalState).F5 = p0
				_650_v29 = _dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_this).F13())
				var _672_v46 _dafny.Array
				_ = _672_v46
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw115
				_672_v46 = _nw115
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_672_v46), 0))
				_ = _index109
				(_672_v46).ArraySet1(p0, (_index109).Int())
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_672_v46), 0))
				_ = _index110
				(_672_v46).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_650_v29).Cardinality())), (_index110).Int())
				var _673_v47 _dafny.Map
				_ = _673_v47
				_673_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F13())
				var _674_v48 _dafny.Array
				_ = _674_v48
				var _nwElement0_26 _dafny.Map = _673_v47
				_ = _nwElement0_26
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(29))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_26, 0)
				(_nw116).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.CodePoint('v')), 1)
				(_nw116).ArraySet1(_673_v47, 2)
				(_nw116).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F13()), 3)
				(_nw116).ArraySet1(((_673_v47).Update(!(p1), (_this).F13())).Merge(_673_v47), 4)
				(_nw116).ArraySet1((_673_v47).Merge(_673_v47), 5)
				(_nw116).ArraySet1(_673_v47, 6)
				(_nw116).ArraySet1(_673_v47, 7)
				(_nw116).ArraySet1((_673_v47).Merge(_673_v47), 8)
				(_nw116).ArraySet1(_673_v47, 9)
				(_nw116).ArraySet1(_673_v47, 10)
				(_nw116).ArraySet1((_673_v47).Merge(_673_v47), 11)
				(_nw116).ArraySet1((func() _dafny.Map {
					if p1 {
						return _673_v47
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F13())
				})(), 12)
				(_nw116).ArraySet1((_673_v47).Merge(_673_v47), 13)
				(_nw116).ArraySet1(_673_v47, 14)
				(_nw116).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F13()), 15)
				(_nw116).ArraySet1(_673_v47, 16)
				(_nw116).ArraySet1((((_673_v47).Update(p1, (_this).F13())).Update(p1, _dafny.CodePoint('t'))).Merge(_673_v47), 17)
				(_nw116).ArraySet1(_673_v47, 18)
				(_nw116).ArraySet1(_673_v47, 19)
				(_nw116).ArraySet1(_673_v47, 20)
				(_nw116).ArraySet1(_673_v47, 21)
				(_nw116).ArraySet1(_673_v47, 22)
				(_nw116).ArraySet1(_673_v47, 23)
				(_nw116).ArraySet1(_673_v47, 24)
				(_nw116).ArraySet1((_673_v47).Update(false, (_this).F13()), 25)
				(_nw116).ArraySet1(_673_v47, 26)
				(_nw116).ArraySet1((_673_v47).Merge(_673_v47), 27)
				(_nw116).ArraySet1(_673_v47, 28)
				_674_v48 = _nw116
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_674_v48), 0))
				_ = _index111
				(_674_v48).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F13())).Update(p1, (_this).F13()), (_index111).Int())
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_674_v48), 0))
				_ = _index112
				(_674_v48).ArraySet1(_673_v47, (_index112).Int())
			}
			var _675_v49 _dafny.Array
			_ = _675_v49
			var _nw117 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(6))
			_ = _nw117
			_675_v49 = _nw117
			var _676_v50 _dafny.Map
			_ = _676_v50
			_676_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _677_v51 _dafny.Array
			_ = _677_v51
			var _nwElement0_27 bool = p1
			_ = _nwElement0_27
			var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(28))
			_ = _nw118
			(_nw118).ArraySet1(_nwElement0_27, 0)
			(_nw118).ArraySet1(p1, 1)
			(_nw118).ArraySet1(p1, 2)
			(_nw118).ArraySet1(false, 3)
			(_nw118).ArraySet1(false, 4)
			(_nw118).ArraySet1(p1, 5)
			(_nw118).ArraySet1(false, 6)
			(_nw118).ArraySet1(p1, 7)
			(_nw118).ArraySet1(p1, 8)
			(_nw118).ArraySet1(p1, 9)
			(_nw118).ArraySet1(p1, 10)
			(_nw118).ArraySet1(p1, 11)
			(_nw118).ArraySet1(true, 12)
			(_nw118).ArraySet1(p1, 13)
			(_nw118).ArraySet1(p1, 14)
			(_nw118).ArraySet1(p1, 15)
			(_nw118).ArraySet1(p1, 16)
			(_nw118).ArraySet1(p1, 17)
			(_nw118).ArraySet1(p1, 18)
			(_nw118).ArraySet1(p1, 19)
			(_nw118).ArraySet1((func() bool {
				if (_676_v50).Contains(p1) {
					return (_676_v50).Get(p1).(bool)
				}
				return p1
			})(), 20)
			(_nw118).ArraySet1(p1, 21)
			(_nw118).ArraySet1(p1, 22)
			(_nw118).ArraySet1(p1, 23)
			(_nw118).ArraySet1(p1, 24)
			(_nw118).ArraySet1(false, 25)
			(_nw118).ArraySet1(!(false), 26)
			(_nw118).ArraySet1(p1, 27)
			_677_v51 = _nw118
			var _678_v52 D6
			_ = _678_v52
			_678_v52 = Companion_D6_.Create_DC16_(_677_v51, _this.F12())
			var _679_v53 D6
			_ = _679_v53
			_679_v53 = Companion_D6_.Create_DC19_(_678_v52)
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_675_v49), 0))
			_ = _index113
			(_675_v49).ArraySet1(_679_v53, (_index113).Int())
			var _680_v54 _dafny.Array
			_ = _680_v54
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_18
			var _nw119 _dafny.Array
			_ = _nw119
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw119 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) _dafny.CodePoint = func(_681_i4 _dafny.Int) _dafny.CodePoint {
					return (_this).F13()
				}
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw119 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw119).ArraySet1CodePoint(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw119).ArraySet1CodePoint(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_680_v54 = _nw119
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_680_v54), 0))
			_ = _index114
			(_680_v54).ArraySet1CodePoint(_dafny.CodePoint('l'), (_index114).Int())
			var _682_v55 _dafny.Map
			_ = _682_v55
			_682_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(864))
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_675_v49), 0))
			_ = _index115
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_680_v54), 0))
			_ = _index116
			var _rhs101 bool = p1
			_ = _rhs101
			var _rhs102 D6 = _679_v53
			_ = _rhs102
			var _rhs103 _dafny.CodePoint = _dafny.CodePoint('k')
			_ = _rhs103
			var _rhs104 _dafny.Map = (_682_v55).Merge(_682_v55)
			_ = _rhs104
			var _lhs61 _dafny.Array = _675_v49
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_675_v49), 0))
			_ = _lhs62
			var _lhs63 _dafny.Array = _680_v54
			_ = _lhs63
			var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_680_v54), 0))
			_ = _lhs64
			r1 = _rhs101
			(_lhs61).ArraySet1(_rhs102, (_lhs62).Int())
			(_lhs63).ArraySet1CodePoint(_rhs103, (_lhs64).Int())
			_682_v55 = _rhs104
			var _683_v56 _dafny.Sequence
			_ = _683_v56
			_683_v56 = _dafny.SeqOf(p1)
			var _684_v57 _dafny.Sequence
			_ = _684_v57
			_684_v57 = _dafny.SeqOf(_683_v56)
			var _685_v58 _dafny.Array
			_ = _685_v58
			var _nwElement0_28 _dafny.Sequence = _683_v56
			_ = _nwElement0_28
			var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(9))
			_ = _nw120
			(_nw120).ArraySet1(_nwElement0_28, 0)
			(_nw120).ArraySet1(_683_v56, 1)
			(_nw120).ArraySet1(_dafny.SeqOf(p1), 2)
			(_nw120).ArraySet1(_683_v56, 3)
			(_nw120).ArraySet1(_683_v56, 4)
			(_nw120).ArraySet1(_683_v56, 5)
			(_nw120).ArraySet1((_684_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_684_v57, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_684_v57).Cardinality()))).Uint32(), Companion_Default___.Fm9(p1, globalState))).Cardinality()), _dafny.IntOfUint32((_684_v57).Cardinality()))).Uint32()).(_dafny.Sequence), 6)
			(_nw120).ArraySet1(_683_v56, 7)
			(_nw120).ArraySet1(_683_v56, 8)
			_685_v58 = _nw120
			var _686_v59 _dafny.Sequence
			_ = _686_v59
			_686_v59 = _dafny.SeqOf(_685_v58)
			var _687_v60 _dafny.Array
			_ = _687_v60
			var _nwElement0_29 _dafny.Array = _685_v58
			_ = _nwElement0_29
			var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(6))
			_ = _nw121
			(_nw121).ArraySet1(_nwElement0_29, 0)
			(_nw121).ArraySet1(_685_v58, 1)
			(_nw121).ArraySet1(_685_v58, 2)
			(_nw121).ArraySet1(_685_v58, 3)
			(_nw121).ArraySet1(_685_v58, 4)
			(_nw121).ArraySet1((_686_v59).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F12()), _dafny.IntOfUint32((_686_v59).Cardinality()))).Uint32()).(_dafny.Array), 5)
			_687_v60 = _nw121
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_687_v60), 0))
			_ = _index117
			(_687_v60).ArraySet1(_685_v58, (_index117).Int())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_687_v60), 0))
			_ = _index118
			(_687_v60).ArraySet1(_685_v58, (_index118).Int())
		}
		var _688_v61 _dafny.MultiSet
		_ = _688_v61
		_688_v61 = _dafny.MultiSetOf(p0)
		var _689_i5 _dafny.Int
		_ = _689_i5
		_689_i5 = _dafny.Zero
		{
			for !(_688_v61).Contains(p0) {
				{
					if (_689_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_689_i5 = (_689_i5).Plus(_dafny.One)
					var _rhs105 _dafny.Int = (_dafny.Zero).Minus(_this.F12())
					_ = _rhs105
					var _rhs106 _dafny.Int = _this.F12()
					_ = _rhs106
					var _lhs65 *GlobalState = globalState
					_ = _lhs65
					_lhs65.F5 = _rhs105
					r2 = _rhs106
					var _690_v62 D5
					_ = _690_v62
					_690_v62 = Companion_D5_.Create_DC14_(p1)
					var _source8 D5 = _690_v62
					_ = _source8
					if _source8.Is_DC13() {
						var _691___mcc_h0 bool = _source8.Get_().(D5_DC13).Cf18
						_ = _691___mcc_h0
						var _692___mcc_h1 bool = _source8.Get_().(D5_DC13).Cf19
						_ = _692___mcc_h1
						var _693_cf19 bool = _692___mcc_h1
						_ = _693_cf19
						var _694_cf18 bool = _691___mcc_h0
						_ = _694_cf18
						var _695_v63 _dafny.Map
						_ = _695_v63
						_695_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), !(p1))
						_694_cf18 = (func() bool {
							if (_695_v63).Contains(_this.F12()) {
								return (_695_v63).Get(_this.F12()).(bool)
							}
							return !(_693_cf19)
						})()
						r1 = !(!(p1))
						(globalState).F5 = (func() _dafny.Int {
							if _693_cf19 {
								return _dafny.IntOfInt64(507)
							}
							return _this.F12()
						})()
						var _696_v64 _dafny.Array
						_ = _696_v64
						var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
						_ = _nw122
						_696_v64 = _nw122
						var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_696_v64), 0))
						_ = _index119
						(_696_v64).ArraySet1(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((p2).Cardinality())), (_index119).Int())
						var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_696_v64), 0))
						_ = _index120
						(_696_v64).ArraySet1((_dafny.Zero).Minus(p0), (_index120).Int())
					} else if _source8.Is_DC14() {
						var _697___mcc_h2 bool = _source8.Get_().(D5_DC14).Cf20
						_ = _697___mcc_h2
						var _698_cf20 bool = _697___mcc_h2
						_ = _698_cf20
						var _699_v65 _dafny.Sequence
						_ = _699_v65
						_699_v65 = _dafny.SeqOf(p1)
						var _700_v66 _dafny.Array
						_ = _700_v66
						var _nwElement0_30 bool = !(_698_cf20)
						_ = _nwElement0_30
						var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(21))
						_ = _nw123
						(_nw123).ArraySet1(_nwElement0_30, 0)
						(_nw123).ArraySet1(p1, 1)
						(_nw123).ArraySet1(_698_cf20, 2)
						(_nw123).ArraySet1(p1, 3)
						(_nw123).ArraySet1(!(p1), 4)
						(_nw123).ArraySet1(_698_cf20, 5)
						(_nw123).ArraySet1(p1, 6)
						(_nw123).ArraySet1(!((p0).Cmp(p0) < 0), 7)
						(_nw123).ArraySet1(p1, 8)
						(_nw123).ArraySet1(!(_698_cf20), 9)
						(_nw123).ArraySet1(p1, 10)
						(_nw123).ArraySet1(_698_cf20, 11)
						(_nw123).ArraySet1((_699_v65).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_699_v65).Cardinality()))).Uint32()).(bool), 12)
						(_nw123).ArraySet1(_698_cf20, 13)
						(_nw123).ArraySet1(!(p1) || (p1), 14)
						(_nw123).ArraySet1((func() bool {
							if p1 {
								return false
							}
							return _698_cf20
						})(), 15)
						(_nw123).ArraySet1(_698_cf20, 16)
						(_nw123).ArraySet1(!((func() bool {
							if _698_cf20 {
								return _698_cf20
							}
							return p1
						})()), 17)
						(_nw123).ArraySet1((_this.F12()).Cmp(_this.F12()) >= 0, 18)
						(_nw123).ArraySet1((p1) == (_698_cf20), 19)
						(_nw123).ArraySet1(_698_cf20, 20)
						_700_v66 = _nw123
						var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_700_v66), 0))
						_ = _index121
						(_700_v66).ArraySet1((p0).Cmp(p0) == 0, (_index121).Int())
						var _701_v67 _dafny.Sequence
						_ = _701_v67
						_701_v67 = _dafny.SeqOf(_dafny.IntOfInt64(850))
						var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_700_v66), 0))
						_ = _index122
						(_700_v66).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_this.F12()), (func() _dafny.Sequence {
							if p1 {
								return _701_v67
							}
							return (_this).Fm5(p1, _698_cf20, _698_cf20, !(p1), globalState)
						})()), (_index122).Int())
						(globalState).F5 = p0
						var _702_v68 D2
						_ = _702_v68
						_702_v68 = Companion_D2_.Create_DC8_(_698_cf20, p0, _698_cf20)
						var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_700_v66), 0))
						_ = _index123
						(_700_v66).ArraySet1(!(Companion_Default___.Fm1(_dafny.IntOfInt64(757), Companion_Default___.Fm1(_this.F12(), (_700_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_700_v66), 0))).Int()).(bool), _this.F12(), (_700_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_700_v66), 0))).Int()).(bool), globalState), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-788)), p1)).Cardinality(), (_702_v68).Dtor_cf13(), globalState)), (_index123).Int())
						var _703_v69 _dafny.Set
						_ = _703_v69
						_703_v69 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_698_cf20, !(Companion_Default___.Fm1(_dafny.IntOfInt64(-745), (_700_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_700_v66), 0))).Int()).(bool), p0, _698_cf20, globalState))))
						var _704_v71 _dafny.MultiSet
						_ = _704_v71
						_704_v71 = _dafny.MultiSetOf(_dafny.CodePoint('j'))
						var _705_v72 _dafny.Array
						_ = _705_v72
						var _nwElement0_31 _dafny.Int = _dafny.IntOfInt64(176)
						_ = _nwElement0_31
						var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(17))
						_ = _nw124
						(_nw124).ArraySet1(_nwElement0_31, 0)
						(_nw124).ArraySet1(_this.F12(), 1)
						(_nw124).ArraySet1(Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(_this.F12())), 2)
						(_nw124).ArraySet1((func() _dafny.Int {
							if false {
								return (_dafny.Zero).Minus((_703_v69).Cardinality())
							}
							return _this.F12()
						})(), 3)
						(_nw124).ArraySet1(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(440)), 4)
						(_nw124).ArraySet1(p0, 5)
						(_nw124).ArraySet1(_this.F12(), 6)
						(_nw124).ArraySet1(p0, 7)
						(_nw124).ArraySet1(_this.F12(), 8)
						(_nw124).ArraySet1((func() _dafny.Map {
							var _coll19 = _dafny.NewMapBuilder()
							_ = _coll19
							for _iter19 := _dafny.Iterate((_704_v71).Elements()); ; {
								_compr_19, _ok19 := _iter19()
								if !_ok19 {
									break
								}
								var _706_v70 _dafny.CodePoint
								_706_v70 = interface{}(_compr_19).(_dafny.CodePoint)
								if (_704_v71).Contains(_706_v70) {
									_coll19.Add(_706_v70, _698_cf20)
								}
							}
							return _coll19.ToMap()
						}()).Cardinality(), 9)
						(_nw124).ArraySet1(p0, 10)
						(_nw124).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-239)), 11)
						(_nw124).ArraySet1((_this.F12()).Plus((_dafny.MultiSetFromSeq(_699_v65)).Cardinality()), 12)
						(_nw124).ArraySet1((_this.F12()).Plus(p0), 13)
						(_nw124).ArraySet1(_this.F12(), 14)
						(_nw124).ArraySet1(_dafny.IntOfInt64(536), 15)
						(_nw124).ArraySet1(_this.F12(), 16)
						_705_v72 = _nw124
						_705_v72 = _705_v72
					} else {
						var _707___mcc_h3 _dafny.Sequence = _source8.Get_().(D5_DC12).Cf17
						_ = _707___mcc_h3
						var _708_cf17 _dafny.Sequence = _707___mcc_h3
						_ = _708_cf17
						var _709_v73 *C1
						_ = _709_v73
						var _nw125 *C1 = New_C1_()
						_ = _nw125
						_nw125.Ctor__((_this).F13(), (_this).F13(), _this.F12())
						_709_v73 = _nw125
						var _710_v74 _dafny.Array
						_ = _710_v74
						var _len0_19 _dafny.Int = _dafny.IntOfInt64(28)
						_ = _len0_19
						var _nw126 _dafny.Array
						_ = _nw126
						if _len0_19.Cmp(_dafny.Zero) == 0 {
							_nw126 = _dafny.NewArray(_len0_19)
						} else {
							var _init19 func(_dafny.Int) _dafny.Int = (func(_711_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
								return func(_712_i6 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeDivisionInt(_712_i6, (_dafny.SetOf(_dafny.IntOfUint32((_711_p2).Cardinality()))).Cardinality())
								}
							})(p2)
							_ = _init19
							var _element0_19 = _init19(_dafny.Zero)
							_ = _element0_19
							_nw126 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
							(_nw126).ArraySet1(_element0_19, 0)
							var _nativeLen0_19 = (_len0_19).Int()
							_ = _nativeLen0_19
							for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
								(_nw126).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
							}
						}
						_710_v74 = _nw126
						var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_710_v74), 0))
						_ = _index124
						(_710_v74).ArraySet1(_dafny.IntOfInt64(299), (_index124).Int())
						var _713_v75 _dafny.CodePoint
						_ = _713_v75
						_713_v75 = _dafny.CodePoint('m')
						var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_710_v74), 0))
						_ = _index125
						var _rhs107 *C1 = _709_v73
						_ = _rhs107
						var _rhs108 _dafny.Int = p0
						_ = _rhs108
						var _rhs109 _dafny.CodePoint = (_709_v73).F15()
						_ = _rhs109
						var _lhs66 _dafny.Array = _710_v74
						_ = _lhs66
						var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_710_v74), 0))
						_ = _lhs67
						_709_v73 = _rhs107
						(_lhs66).ArraySet1(_rhs108, (_lhs67).Int())
						_713_v75 = _rhs109
						_713_v75 = (_709_v73).F15()
						(globalState).F5 = (_709_v73).Fm4(globalState)
						var _714_v76 _dafny.Int
						_ = _714_v76
						var _715_v77 _dafny.Map
						_ = _715_v77
						var _716_v78 _dafny.Sequence
						_ = _716_v78
						var _out42 _dafny.Int
						_ = _out42
						var _out43 _dafny.Map
						_ = _out43
						var _out44 _dafny.Sequence
						_ = _out44
						_out42, _out43, _out44 = Companion_Default___.M0((_710_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_710_v74), 0))).Int()).(_dafny.Int), ((_710_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_710_v74), 0))).Int()).(_dafny.Int)).Times(_this.F12()), ((_710_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_710_v74), 0))).Int()).(_dafny.Int)).Plus((_710_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_710_v74), 0))).Int()).(_dafny.Int)), true, globalState)
						_714_v76 = _out42
						_715_v77 = _out43
						_716_v78 = _out44
					}
					var _717_v79 _dafny.Map
					_ = _717_v79
					_717_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F12()), p2)
					var _718_v80 _dafny.Sequence
					_ = _718_v80
					_718_v80 = _dafny.SeqOf(true)
					_717_v79 = (_717_v79).Update(_dafny.IntOfUint32((_718_v80).Cardinality()), _dafny.Companion_Sequence_.Concatenate(p2, p2))
					var _719_v81 _dafny.Array
					_ = _719_v81
					var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
					_ = _nw127
					_719_v81 = _nw127
					var _rhs110 _dafny.Array = _719_v81
					_ = _rhs110
					var _rhs111 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm0(p0, _this.F12(), globalState))
					_ = _rhs111
					var _lhs68 *C3 = _this
					_ = _lhs68
					_719_v81 = _rhs110
					_lhs68.F12_set_(_rhs111)
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _hi3 _dafny.Int = _this.F12()
		_ = _hi3
		for _720_i7 := (_dafny.IntOfInt64(-779)).Times(_dafny.IntOfInt64(357)); _720_i7.Cmp(_hi3) < 0; _720_i7 = _720_i7.Plus(_dafny.One) {
			r0 = (_dafny.MultiSetOf(p0)).IsSubsetOf(_688_v61)
			var _721_v82 *C0
			_ = _721_v82
			var _nw128 *C0 = New_C0_()
			_ = _nw128
			_nw128.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_i7, p1), _720_i7)
			_721_v82 = _nw128
			var _722_v83 _dafny.Map
			_ = _722_v83
			_722_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _723_v84 _dafny.Sequence
			_ = _723_v84
			_723_v84 = _dafny.SeqOf(p1, (func() bool {
				if (_722_v83).Contains(!(p1)) {
					return (_722_v83).Get(!(p1)).(bool)
				}
				return p1
			})(), p1, p1, p1)
			var _724_v85 _dafny.Sequence
			_ = _724_v85
			_724_v85 = _dafny.SeqOf(_723_v84)
			if (_dafny.IntOfUint32((_724_v85).Cardinality())).Cmp(_this.F12()) >= 0 {
				var _725_v86 _dafny.Array
				_ = _725_v86
				var _nw129 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw129
				_725_v86 = _nw129
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_725_v86), 0))
				_ = _index126
				(_725_v86).ArraySet1(p1, (_index126).Int())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_725_v86), 0))
				_ = _index127
				(_725_v86).ArraySet1(!_dafny.Companion_Sequence_.Equal(p2, p2), (_index127).Int())
				var _726_v87 _dafny.Array
				_ = _726_v87
				var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw130
				_726_v87 = _nw130
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_726_v87), 0))
				_ = _index128
				(_726_v87).ArraySet1(_720_i7, (_index128).Int())
				var _727_v88 _dafny.Set
				_ = _727_v88
				_727_v88 = _dafny.SetOf(false, (_725_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_725_v86), 0))).Int()).(bool))
				var _728_v89 _dafny.Map
				_ = _728_v89
				_728_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _727_v88)
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_726_v87), 0))
				_ = _index129
				(_726_v87).ArraySet1(((func() _dafny.Set {
					if (_728_v89).Contains((_this.F12()).Cmp(_720_i7) == 0) {
						return (_728_v89).Get((_this.F12()).Cmp(_720_i7) == 0).(_dafny.Set)
					}
					return _727_v88
				})()).Cardinality(), (_index129).Int())
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_725_v86), 0))
				_ = _index130
				(_725_v86).ArraySet1(((Companion_D0_.Create_DC0_(p1)).Dtor_cf0()) == ((_725_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_725_v86), 0))).Int()).(bool)), (_index130).Int())
				r1 = (_725_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_725_v86), 0))).Int()).(bool)
				var _729_v90 _dafny.Sequence
				_ = _729_v90
				_729_v90 = _dafny.SeqOf(_dafny.IntOfInt64(796))
				var _730_v91 _dafny.MultiSet
				_ = _730_v91
				_730_v91 = _dafny.MultiSetOf(_729_v90)
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_726_v87), 0))
				_ = _index131
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_726_v87), 0))
				_ = _index132
				var _rhs112 _dafny.Int = _720_i7
				_ = _rhs112
				var _rhs113 _dafny.Int = (func() _dafny.Int {
					if (_730_v91).Contains(_dafny.Companion_Sequence_.Concatenate(_729_v90, _729_v90)) {
						return (_730_v91).Multiplicity(_dafny.Companion_Sequence_.Concatenate(_729_v90, _729_v90))
					}
					return p0
				})()
				_ = _rhs113
				var _rhs114 _dafny.Int = _720_i7
				_ = _rhs114
				var _lhs69 _dafny.Array = _726_v87
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_726_v87), 0))
				_ = _lhs70
				var _lhs71 *GlobalState = globalState
				_ = _lhs71
				var _lhs72 _dafny.Array = _726_v87
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_726_v87), 0))
				_ = _lhs73
				(_lhs69).ArraySet1(_rhs112, (_lhs70).Int())
				_lhs71.F5 = _rhs113
				(_lhs72).ArraySet1(_rhs114, (_lhs73).Int())
			} else {
				(globalState).F5 = (_dafny.Zero).Minus(_720_i7)
				var _731_v92 _dafny.MultiSet
				_ = _731_v92
				_731_v92 = _dafny.MultiSetOf(p1, p1)
				_731_v92 = _dafny.MultiSetOf(p1)
				var _732_v93 _dafny.Array
				_ = _732_v93
				var _nw131 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw131
				_732_v93 = _nw131
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_732_v93), 0))
				_ = _index133
				(_732_v93).ArraySet1(Companion_Default___.Fm1(_this.F12(), p1, _720_i7, false, globalState), (_index133).Int())
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_732_v93), 0))
				_ = _index134
				(_732_v93).ArraySet1(p1, (_index134).Int())
				var _733_v94 _dafny.CodePoint
				_ = _733_v94
				_733_v94 = _dafny.CodePoint('j')
				_733_v94 = (_this).F13()
				var _734_v95 _dafny.Int
				_ = _734_v95
				var _735_v96 _dafny.Map
				_ = _735_v96
				var _736_v97 _dafny.Sequence
				_ = _736_v97
				var _out45 _dafny.Int
				_ = _out45
				var _out46 _dafny.Map
				_ = _out46
				var _out47 _dafny.Sequence
				_ = _out47
				_out45, _out46, _out47 = Companion_Default___.M0(_this.F12(), (_dafny.Zero).Minus(Companion_Default___.Fm0(p0, p0, globalState)), _dafny.IntOfInt64(406), (_732_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_732_v93), 0))).Int()).(bool), globalState)
				_734_v95 = _out45
				_735_v96 = _out46
				_736_v97 = _out47
			}
			var _737_v98 _dafny.Sequence
			_ = _737_v98
			_737_v98 = _dafny.UnicodeSeqOfUtf8Bytes("b")
			_737_v98 = _dafny.Companion_Sequence_.Concatenate(_737_v98, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(831))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg39 _dafny.Int) interface{} {
					return coer39(arg39)
				}
			}(func(_738_i8 _dafny.Int) _dafny.CodePoint {
				return (_this).F13()
			})), _dafny.UnicodeSeqOfUtf8Bytes("hxb")))
		}
		var _739_v99 _dafny.Array
		_ = _739_v99
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_20
		var _nw132 _dafny.Array
		_ = _nw132
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw132 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) bool = (func(_740_p1 bool) func(_dafny.Int) bool {
				return func(_741_i9 _dafny.Int) bool {
					return _740_p1
				}
			})(p1)
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw132 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw132).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw132).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_739_v99 = _nw132
		var _742_v100 D0
		_ = _742_v100
		_742_v100 = Companion_D0_.Create_DC1_(_739_v99)
		var _source9 D0 = _742_v100
		_ = _source9
		if _source9.Is_DC1() {
			var _743___mcc_h4 _dafny.Array = _source9.Get_().(D0_DC1).Cf1
			_ = _743___mcc_h4
			var _744_cf1 _dafny.Array = _743___mcc_h4
			_ = _744_cf1
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_739_v99), 0))
			_ = _index135
			(_739_v99).ArraySet1(!(!(p1)), (_index135).Int())
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_739_v99), 0))
			_ = _index136
			(_739_v99).ArraySet1(p1, (_index136).Int())
			var _745_v101 _dafny.Sequence
			_ = _745_v101
			_745_v101 = _dafny.UnicodeSeqOfUtf8Bytes("uadjarwl")
			_745_v101 = _745_v101
			var _746_v102 D2
			_ = _746_v102
			_746_v102 = Companion_D2_.Create_DC9_(p0)
			var _747_v103 _dafny.Set
			_ = _747_v103
			_747_v103 = Companion_Default___.Fm14(_dafny.IntOfInt64(-662), _746_v102, globalState)
			var _source10 _dafny.Set = _747_v103
			_ = _source10
			var _748___mcc_h11 _dafny.Set = _source10
			_ = _748___mcc_h11
			var _749_cf16 _dafny.Set = _748___mcc_h11
			_ = _749_cf16
			var _750_v104 _dafny.Map
			_ = _750_v104
			_750_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), true)
			_750_v104 = (_750_v104).Update(Companion_Default___.Fm1(p0, p1, p0, p1, globalState), (_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool))
			var _751_v105 _dafny.Sequence
			_ = _751_v105
			_751_v105 = _dafny.SeqOf(_this.F12(), p0, p0, (_this.F12()).Times(p0))
			_751_v105 = _751_v105
			_688_v61 = _688_v61
			var _752_v106 T1
			_ = _752_v106
			var _nw133 *C1 = New_C1_()
			_ = _nw133
			_nw133.Ctor__(_dafny.CodePoint('l'), (_this).F13(), _dafny.IntOfInt64(738))
			_752_v106 = _nw133
			_752_v106 = _752_v106
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_739_v99), 0))
			_ = _index137
			var _rhs115 _dafny.Int = (_dafny.IntOfUint32((_745_v101).Cardinality())).Times(p0)
			_ = _rhs115
			var _rhs116 bool = p1
			_ = _rhs116
			var _rhs117 _dafny.Int = Companion_Default___.SafeDivisionInt(_this.F12(), _dafny.IntOfInt64(979))
			_ = _rhs117
			var _lhs74 _dafny.Array = _739_v99
			_ = _lhs74
			var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_739_v99), 0))
			_ = _lhs75
			var _lhs76 *C3 = _this
			_ = _lhs76
			r2 = _rhs115
			(_lhs74).ArraySet1(_rhs116, (_lhs75).Int())
			_lhs76.F12_set_(_rhs117)
		} else if _source9.Is_DC2() {
			var _753___mcc_h5 bool = _source9.Get_().(D0_DC2).Cf2
			_ = _753___mcc_h5
			var _754___mcc_h6 _dafny.Map = _source9.Get_().(D0_DC2).Cf3
			_ = _754___mcc_h6
			var _755___mcc_h7 _dafny.Map = _source9.Get_().(D0_DC2).Cf4
			_ = _755___mcc_h7
			var _756___mcc_h8 bool = _source9.Get_().(D0_DC2).Cf5
			_ = _756___mcc_h8
			var _757_cf5 bool = _756___mcc_h8
			_ = _757_cf5
			var _758_cf4 _dafny.Map = _755___mcc_h7
			_ = _758_cf4
			var _759_cf3 _dafny.Map = _754___mcc_h6
			_ = _759_cf3
			var _760_cf2 bool = _753___mcc_h5
			_ = _760_cf2
			r0 = (_688_v61).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_this.F12(), p0, _this.F12())).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((p2).Cardinality())))))
			var _761_v107 T0
			_ = _761_v107
			var _nw134 *C2 = New_C2_()
			_ = _nw134
			_nw134.Ctor__(_760_cf2, _dafny.IntOfInt64(369))
			_761_v107 = _nw134
			var _rhs118 T0 = _761_v107
			_ = _rhs118
			var _rhs119 bool = _760_cf2
			_ = _rhs119
			_761_v107 = _rhs118
			_760_cf2 = _rhs119
			var _762_v108 *C1
			_ = _762_v108
			var _nw135 *C1 = New_C1_()
			_ = _nw135
			_nw135.Ctor__((_this).F13(), (_this).F13(), _this.F12())
			_762_v108 = _nw135
			if !(true) {
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))
				_ = _index138
				(_739_v99).ArraySet1(_760_cf2, (_index138).Int())
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))
				_ = _index139
				(_739_v99).ArraySet1(!(true), (_index139).Int())
				(globalState).F5 = _dafny.IntOfUint32((p2).Cardinality())
				var _763_v109 _dafny.Set
				_ = _763_v109
				_763_v109 = _dafny.SetOf(_760_cf2, Companion_Default___.Fm1(p0, _757_cf5, p0, (_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool), globalState))
				var _764_v110 _dafny.Set
				_ = _764_v110
				_764_v110 = _763_v109
				var _rhs120 _dafny.Int = p0
				_ = _rhs120
				var _rhs121 _dafny.Set = _764_v110
				_ = _rhs121
				var _lhs77 T0 = _761_v107
				_ = _lhs77
				_lhs77.F12_set_(_rhs120)
				_764_v110 = _rhs121
				var _765_v111 *C0
				_ = _765_v111
				var _nw136 *C0 = New_C0_()
				_ = _nw136
				_nw136.Ctor__((_759_cf3).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p2).Cardinality()), (_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool))), p0)
				_765_v111 = _nw136
				var _766_v112 _dafny.Array
				_ = _766_v112
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_21
				var _nw137 _dafny.Array
				_ = _nw137
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw137 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) bool = (func(_767_cf5 bool) func(_dafny.Int) bool {
						return func(_768_i10 _dafny.Int) bool {
							return _767_cf5
						}
					})(_757_cf5)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw137 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw137).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw137).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_766_v112 = _nw137
				var _769_v113 D6
				_ = _769_v113
				_769_v113 = Companion_D6_.Create_DC16_(_766_v112, _761_v107.F12())
				var _770_v115 _dafny.Sequence
				_ = _770_v115
				_770_v115 = _dafny.SeqOf(false)
				var _771_v116 _dafny.Map
				_ = _771_v116
				_771_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool), p0)
				var _772_v117 _dafny.Sequence
				_ = _772_v117
				_772_v117 = _dafny.SeqOf(_761_v107.F12(), _dafny.IntOfInt64(471), _this.F12())
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))
				_ = _index140
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))
				_ = _index141
				var _rhs122 _dafny.Int = (_769_v113).Dtor_cf23()
				_ = _rhs122
				var _rhs123 bool = ((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(945), _dafny.IntOfInt64(-119))); ; {
						_compr_20, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _773_v114 _dafny.Int
						_773_v114 = interface{}(_compr_20).(_dafny.Int)
						if ((_dafny.IntOfInt64(945)).Cmp(_773_v114) <= 0) && ((_773_v114).Cmp(_dafny.IntOfInt64(-119)) < 0) {
							_coll20.Add(Companion_Default___.SafeModuloInt(_773_v114, _761_v107.F12()), _761_v107.F12())
						}
					}
					return _coll20.ToMap()
				}()).Cardinality()))).Cmp(_761_v107.F12()) < 0
				_ = _rhs123
				var _rhs124 bool = (func() bool {
					if !((_770_v115).Select((Companion_Default___.SafeIndex((_771_v116).Cardinality(), _dafny.IntOfUint32((_770_v115).Cardinality()))).Uint32()).(bool)) || (false) {
						return _760_cf2
					}
					return Companion_Default___.Fm1(_dafny.IntOfUint32((_772_v117).Cardinality()), p1, _this.F12(), Companion_Default___.Fm1(_dafny.IntOfInt64(610), (_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool), _761_v107.F12(), _760_cf2, globalState), globalState)
				})()
				_ = _rhs124
				var _lhs78 T0 = _761_v107
				_ = _lhs78
				var _lhs79 _dafny.Array = _739_v99
				_ = _lhs79
				var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))
				_ = _lhs80
				var _lhs81 _dafny.Array = _739_v99
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_739_v99), 0))
				_ = _lhs82
				_lhs78.F12_set_(_rhs122)
				(_lhs79).ArraySet1(_rhs123, (_lhs80).Int())
				(_lhs81).ArraySet1(_rhs124, (_lhs82).Int())
			} else {
				var _774_v118 _dafny.Array
				_ = _774_v118
				var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw138
				_774_v118 = _nw138
				var _775_v119 _dafny.Set
				_ = _775_v119
				_775_v119 = _dafny.SetOf(_759_cf3)
				var _rhs125 _dafny.Array = _774_v118
				_ = _rhs125
				var _rhs126 _dafny.Int = Companion_Default___.SafeModuloInt((_775_v119).Cardinality(), _dafny.IntOfInt64(808))
				_ = _rhs126
				var _lhs83 T0 = _761_v107
				_ = _lhs83
				_774_v118 = _rhs125
				_lhs83.F12_set_(_rhs126)
				r2 = _761_v107.F12()
				r1 = !(!(_760_cf2))
				var _776_v120 _dafny.Sequence
				_ = _776_v120
				_776_v120 = _dafny.UnicodeSeqOfUtf8Bytes("nbjki")
				_776_v120 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("c"), p2)
				var _777_v121 _dafny.Sequence
				_ = _777_v121
				_777_v121 = _dafny.SeqOf(_761_v107.F12())
				var _778_v122 _dafny.Map
				_ = _778_v122
				_778_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_777_v121).Cardinality()), (_762_v108).F15())
				var _779_v123 _dafny.MultiSet
				_ = _779_v123
				_779_v123 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), (_this).F13()), _778_v122, (_778_v122).Update(_this.F12(), _dafny.CodePoint('h')), (_778_v122).Merge(_778_v122), (_778_v122).Merge(_778_v122))
				_779_v123 = _779_v123
			}
		} else if _source9.Is_DC0() {
			var _780___mcc_h9 bool = _source9.Get_().(D0_DC0).Cf0
			_ = _780___mcc_h9
			var _781_cf0 bool = _780___mcc_h9
			_ = _781_cf0
			var _782_v124 _dafny.MultiSet
			_ = _782_v124
			_782_v124 = _dafny.MultiSetOf(_781_cf0)
			var _783_v125 _dafny.Set
			_ = _783_v125
			_783_v125 = _dafny.SetOf((_782_v124).IsSubsetOf(_782_v124))
			_783_v125 = _783_v125
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))
			_ = _index142
			(_739_v99).ArraySet1(_781_cf0, (_index142).Int())
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))
			_ = _index143
			(_739_v99).ArraySet1(_781_cf0, (_index143).Int())
			var _784_v126 _dafny.Array
			_ = _784_v126
			var _nw139 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw139
			_784_v126 = _nw139
			var _785_v127 D6
			_ = _785_v127
			_785_v127 = Companion_D6_.Create_DC16_(_784_v126, p0)
			var _source11 D6 = _785_v127
			_ = _source11
			if _source11.Is_DC16() {
				var _786___mcc_h12 _dafny.Array = _source11.Get_().(D6_DC16).Cf22
				_ = _786___mcc_h12
				var _787___mcc_h13 _dafny.Int = _source11.Get_().(D6_DC16).Cf23
				_ = _787___mcc_h13
				var _788_cf23 _dafny.Int = _787___mcc_h13
				_ = _788_cf23
				var _789_cf22 _dafny.Array = _786___mcc_h12
				_ = _789_cf22
				var _790_v128 _dafny.Array
				_ = _790_v128
				var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw140
				_790_v128 = _nw140
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_790_v128), 0))
				_ = _index144
				(_790_v128).ArraySet1(_788_cf23, (_index144).Int())
				var _791_v129 _dafny.Sequence
				_ = _791_v129
				_791_v129 = _dafny.SeqOf(_789_cf22)
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_790_v128), 0))
				_ = _index145
				(_790_v128).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_791_v129, _791_v129)).Cardinality())).Times(p0), (_index145).Int())
				var _792_v130 _dafny.Map
				_ = _792_v130
				_792_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_790_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_790_v128), 0))).Int()).(_dafny.Int), Companion_Default___.SafeDivisionInt(p0, p0))
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_790_v128), 0))
				_ = _index146
				(_790_v128).ArraySet1((_792_v130).Cardinality(), (_index146).Int())
				var _793_v131 D2
				_ = _793_v131
				_793_v131 = Companion_D2_.Create_DC8_(_781_cf0, _dafny.IntOfInt64(288), _781_cf0)
				(_this).F12_set_(Companion_Default___.Fm0(Companion_Default___.SafeModuloInt((_790_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_790_v128), 0))).Int()).(_dafny.Int), (_793_v131).Dtor_cf12()), _dafny.IntOfInt64(37), globalState))
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))
				_ = _index147
				(_739_v99).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_789_cf22, _784_v126), _791_v129), (_index147).Int())
			} else if _source11.Is_DC17() {
				var _794___mcc_h14 bool = _source11.Get_().(D6_DC17).Cf24
				_ = _794___mcc_h14
				var _795_cf24 bool = _794___mcc_h14
				_ = _795_cf24
				r0 = (_this.F12()).Cmp(_this.F12()) > 0
				var _796_v132 T0
				_ = _796_v132
				var _nw141 *C2 = New_C2_()
				_ = _nw141
				_nw141.Ctor__(p1, _this.F12())
				_796_v132 = _nw141
				var _797_v133 _dafny.Map
				_ = _797_v133
				_797_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_796_v132, _dafny.IntOfInt64(919))
				var _798_v135 _dafny.MultiSet
				_ = _798_v135
				_798_v135 = _dafny.MultiSetOf((_this).F13(), (_this).F13())
				var _799_v136 *C2
				_ = _799_v136
				var _nw142 *C2 = New_C2_()
				_ = _nw142
				_nw142.Ctor__(false, (func() _dafny.Map {
					var _coll21 = _dafny.NewMapBuilder()
					_ = _coll21
					for _iter21 := _dafny.Iterate((_798_v135).Elements()); ; {
						_compr_21, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _800_v134 _dafny.CodePoint
						_800_v134 = interface{}(_compr_21).(_dafny.CodePoint)
						if (_798_v135).Contains(_800_v134) {
							_coll21.Add(_800_v134, p0)
						}
					}
					return _coll21.ToMap()
				}()).Cardinality())
				_799_v136 = _nw142
				_797_v133 = (_797_v133).Update(_796_v132, Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-507), _799_v136)).Cardinality(), p0))
				var _801_v137 _dafny.Sequence
				_ = _801_v137
				_801_v137 = _dafny.SeqOf(p1)
				var _802_v138 _dafny.Int
				_ = _802_v138
				var _803_v139 bool
				_ = _803_v139
				var _804_v140 _dafny.Int
				_ = _804_v140
				var _out48 _dafny.Int
				_ = _out48
				var _out49 bool
				_ = _out49
				var _out50 _dafny.Int
				_ = _out50
				_out48, _out49, _out50 = (_this).M3((func() _dafny.MultiSet {
					if _781_cf0 {
						return _dafny.MultiSetFromSeq(_801_v137)
					}
					return _782_v124
				})(), globalState)
				_802_v138 = _out48
				_803_v139 = _out49
				_804_v140 = _out50
				var _805_v141 _dafny.Sequence
				_ = _805_v141
				_805_v141 = _dafny.SeqOf(p0, _this.F12())
				var _806_v142 _dafny.MultiSet
				_ = _806_v142
				_806_v142 = _dafny.MultiSetOf(_805_v141, _805_v141, _805_v141)
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))
				_ = _index148
				(_739_v99).ArraySet1(((_806_v142).Intersection(_806_v142)).Equals(_806_v142), (_index148).Int())
			} else if _source11.Is_DC18() {
				var _807___mcc_h15 _dafny.Int = _source11.Get_().(D6_DC18).Cf25
				_ = _807___mcc_h15
				var _808_cf25 _dafny.Int = _807___mcc_h15
				_ = _808_cf25
				_808_cf25 = (_dafny.Zero).Minus(_808_cf25)
				var _809_v143 _dafny.Sequence
				_ = _809_v143
				_809_v143 = _dafny.UnicodeSeqOfUtf8Bytes("r")
				_809_v143 = _dafny.UnicodeSeqOfUtf8Bytes("kj")
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))
				_ = _index149
				(_739_v99).ArraySet1(_781_cf0, (_index149).Int())
				var _810_v144 _dafny.CodePoint
				_ = _810_v144
				_810_v144 = _dafny.CodePoint('l')
				var _811_v145 _dafny.Map
				_ = _811_v145
				_811_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _781_cf0)
				var _812_v146 _dafny.Map
				_ = _812_v146
				_812_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_781_cf0, _808_cf25)
				var _813_v147 *C0
				_ = _813_v147
				var _nw143 *C0 = New_C0_()
				_ = _nw143
				_nw143.Ctor__(_811_v145, (_dafny.IntOfInt64(95)).Times((_812_v146).Cardinality()))
				_813_v147 = _nw143
				var _814_v148 D6
				_ = _814_v148
				_814_v148 = Companion_D6_.Create_DC18_(Companion_Default___.SafeModuloInt(_this.F12(), (func() _dafny.Int {
					if (_782_v124).Contains((_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool)) {
						return (_782_v124).Multiplicity((_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool))
					}
					return _808_cf25
				})()))
				var _rhs127 _dafny.Int = (func() _dafny.Int {
					if (_688_v61).Contains(_dafny.IntOfInt64(900)) {
						return (_688_v61).Multiplicity(_dafny.IntOfInt64(900))
					}
					return _this.F12()
				})()
				_ = _rhs127
				var _rhs128 _dafny.Array = _784_v126
				_ = _rhs128
				var _rhs129 _dafny.CodePoint = _810_v144
				_ = _rhs129
				var _rhs130 *C0 = _813_v147
				_ = _rhs130
				var _rhs131 D6 = _814_v148
				_ = _rhs131
				var _lhs84 *C3 = _this
				_ = _lhs84
				_lhs84.F12_set_(_rhs127)
				_784_v126 = _rhs128
				_810_v144 = _rhs129
				_813_v147 = _rhs130
				_814_v148 = _rhs131
			} else if _source11.Is_DC15() {
				var _815___mcc_h16 _dafny.Sequence = _source11.Get_().(D6_DC15).Cf21
				_ = _815___mcc_h16
				var _816_cf21 _dafny.Sequence = _815___mcc_h16
				_ = _816_cf21
				var _817_v149 _dafny.MultiSet
				_ = _817_v149
				_817_v149 = _dafny.MultiSetOf(p2, p2)
				_781_cf0 = (_817_v149).IsDisjointFrom((_817_v149).Difference(_817_v149))
				var _818_v150 *C1
				_ = _818_v150
				var _nw144 *C1 = New_C1_()
				_ = _nw144
				_nw144.Ctor__((_this).F13(), (_this).F13(), _dafny.IntOfInt64(-857))
				_818_v150 = _nw144
				var _819_v151 D6
				_ = _819_v151
				_819_v151 = Companion_D6_.Create_DC15_(_dafny.Companion_Sequence_.Concatenate(_816_cf21, _816_cf21))
				_819_v151 = Companion_D6_.Create_DC15_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(872))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_820_cf21 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_821_i11 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_820_cf21).Cardinality())
					}
				})(_816_cf21))))
				var _822_v152 _dafny.Array
				_ = _822_v152
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_22
				var _nw145 _dafny.Array
				_ = _nw145
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw145 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Sequence = (func(_823_v99 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
						return func(_824_i12 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf((_823_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_823_v99), 0))).Int()).(bool))
						}
					})(_739_v99)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw145 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw145).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw145).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_822_v152 = _nw145
				var _825_v153 _dafny.Map
				_ = _825_v153
				_825_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F12()), _this.F12())
				var _826_v154 _dafny.Set
				_ = _826_v154
				_826_v154 = _dafny.SetOf((_dafny.Zero).Minus((_825_v153).Cardinality()))
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))
				_ = _index150
				var _rhs132 bool = (_826_v154).IsSubsetOf(_826_v154)
				_ = _rhs132
				var _rhs133 bool = (_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool)
				_ = _rhs133
				var _rhs134 _dafny.Int = (_dafny.Zero).Minus(p0)
				_ = _rhs134
				var _rhs135 _dafny.Array = _822_v152
				_ = _rhs135
				var _rhs136 bool = false
				_ = _rhs136
				var _lhs85 _dafny.Array = _739_v99
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))
				_ = _lhs86
				var _lhs87 *GlobalState = globalState
				_ = _lhs87
				r0 = _rhs132
				(_lhs85).ArraySet1(_rhs133, (_lhs86).Int())
				_lhs87.F5 = _rhs134
				_822_v152 = _rhs135
				_781_cf0 = _rhs136
			} else {
				var _827___mcc_h17 D6 = _source11.Get_().(D6_DC19).Cf26
				_ = _827___mcc_h17
				var _828_cf26 D6 = _827___mcc_h17
				_ = _828_cf26
				(globalState).F5 = p0
				var _829_v155 _dafny.Map
				_ = _829_v155
				_829_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_782_v124, _this.F12())
				var _830_v156 D2
				_ = _830_v156
				_830_v156 = Companion_D2_.Create_DC8_(_781_cf0, (_dafny.Zero).Minus(p0), p1)
				_829_v155 = (_829_v155).Update(Companion_Default___.Fm19((_830_v156).Dtor_cf12(), globalState), p0)
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))
				_ = _index151
				(_739_v99).ArraySet1(_781_cf0, (_index151).Int())
				var _831_v157 _dafny.Array
				_ = _831_v157
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw146
				_831_v157 = _nw146
				var _832_v158 _dafny.Map
				_ = _832_v158
				_832_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_739_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_739_v99), 0))).Int()).(bool), _831_v157)
				var _833_v159 _dafny.Array
				_ = _833_v159
				var _nwElement0_32 _dafny.Array = _831_v157
				_ = _nwElement0_32
				var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(5))
				_ = _nw147
				(_nw147).ArraySet1(_nwElement0_32, 0)
				(_nw147).ArraySet1(_831_v157, 1)
				(_nw147).ArraySet1(_831_v157, 2)
				(_nw147).ArraySet1((func() _dafny.Array {
					if (_832_v158).Contains(_781_cf0) {
						return (_832_v158).Get(_781_cf0).(_dafny.Array)
					}
					return _831_v157
				})(), 3)
				(_nw147).ArraySet1(_831_v157, 4)
				_833_v159 = _nw147
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_833_v159), 0))
				_ = _index152
				(_833_v159).ArraySet1(_831_v157, (_index152).Int())
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_833_v159), 0))
				_ = _index153
				(_833_v159).ArraySet1(_831_v157, (_index153).Int())
			}
			r1 = !(p1)
		} else {
			var _834___mcc_h10 D0 = _source9.Get_().(D0_DC3).Cf6
			_ = _834___mcc_h10
			var _835_cf6 D0 = _834___mcc_h10
			_ = _835_cf6
			var _836_v160 _dafny.Array
			_ = _836_v160
			var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw148
			_836_v160 = _nw148
			var _837_v161 _dafny.Set
			_ = _837_v161
			_837_v161 = _dafny.SetOf(_836_v160)
			var _838_v162 _dafny.Map
			_ = _838_v162
			_838_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_837_v161).Cardinality())
			var _839_v163 _dafny.Map
			_ = _839_v163
			_839_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _838_v162)
			_839_v163 = (_839_v163).Update(_dafny.IntOfUint32((p2).Cardinality()), _838_v162)
			if p1 {
				var _840_v164 _dafny.Set
				_ = _840_v164
				_840_v164 = _dafny.SetOf(p1)
				var _841_v165 D6
				_ = _841_v165
				_841_v165 = Companion_D6_.Create_DC17_(p1)
				var _842_v166 _dafny.Sequence
				_ = _842_v166
				_842_v166 = _dafny.SeqOf((_840_v164).IsSubsetOf(_840_v164), (func() bool {
					if p1 {
						return p1
					}
					return p1
				})(), (_dafny.IntOfInt64(188)).Cmp(p0) < 0, true, (_841_v165).Dtor_cf24())
				r1 = (_842_v166).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_842_v166).Cardinality()))).Uint32()).(bool)
				_836_v160 = _836_v160
				var _843_v167 _dafny.Array
				_ = _843_v167
				var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw149
				_843_v167 = _nw149
				var _844_v168 _dafny.Map
				_ = _844_v168
				_844_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_843_v167), 0))
				_ = _index154
				(_843_v167).ArraySet1(_844_v168, (_index154).Int())
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_843_v167), 0))
				_ = _index155
				(_843_v167).ArraySet1((_844_v168).Merge((Companion_Default___.Fm18(_this.F12(), _this.F12(), p1, globalState)).Merge(_844_v168)), (_index155).Int())
				r2 = ((_dafny.Zero).Minus(p0)).Times(_this.F12())
				var _845_v169 _dafny.MultiSet
				_ = _845_v169
				_845_v169 = _dafny.MultiSetOf(p1, true)
				var _846_v170 _dafny.Int
				_ = _846_v170
				var _847_v171 bool
				_ = _847_v171
				var _848_v172 _dafny.Int
				_ = _848_v172
				var _out51 _dafny.Int
				_ = _out51
				var _out52 bool
				_ = _out52
				var _out53 _dafny.Int
				_ = _out53
				_out51, _out52, _out53 = (_this).M3((_845_v169).Difference(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p1), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))).Uint32(), p1))), globalState)
				_846_v170 = _out51
				_847_v171 = _out52
				_848_v172 = _out53
			} else {
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_836_v160), 0))
				_ = _index156
				(_836_v160).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("va")).Cardinality()), (_index156).Int())
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_836_v160), 0))
				_ = _index157
				(_836_v160).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), (_index157).Int())
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_836_v160), 0))
				_ = _index158
				(_836_v160).ArraySet1(_this.F12(), (_index158).Int())
				var _849_v173 _dafny.Sequence
				_ = _849_v173
				_849_v173 = _dafny.SeqOf(p0)
				var _850_v174 _dafny.Map
				_ = _850_v174
				_850_v174 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_688_v61).Contains(_this.F12()) {
						return (_688_v61).Multiplicity(_this.F12())
					}
					return (_849_v173).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p2).Cardinality()), _dafny.IntOfUint32((_849_v173).Cardinality()))).Uint32()).(_dafny.Int)
				})(), _dafny.IntOfUint32((p2).Cardinality())), (_836_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_836_v160), 0))).Int()).(_dafny.Int))
				_850_v174 = (func() _dafny.Map {
					if p1 {
						return _850_v174
					}
					return _850_v174
				})()
				var _851_v175 _dafny.Array
				_ = _851_v175
				var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw150
				_851_v175 = _nw150
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_851_v175), 0))
				_ = _index159
				(_851_v175).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p2, p2), (_index159).Int())
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_851_v175), 0))
				_ = _index160
				(_851_v175).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p2, _dafny.UnicodeSeqOfUtf8Bytes("kemlgdi")), (_index160).Int())
				var _852_v176 _dafny.Set
				_ = _852_v176
				_852_v176 = _dafny.SetOf(p1)
				var _853_v177 _dafny.Sequence
				_ = _853_v177
				_853_v177 = _dafny.SeqOf(_dafny.SetOf(p1, !(true)), _852_v176, _852_v176)
				var _854_v178 _dafny.Set
				_ = _854_v178
				_854_v178 = _dafny.SetOf((_853_v177).Select((Companion_Default___.SafeIndex((_836_v160).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_836_v160), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_853_v177).Cardinality()))).Uint32()).(_dafny.Set), _852_v176, _852_v176)
				var _855_v179 D6
				_ = _855_v179
				_855_v179 = Companion_D6_.Create_DC17_(true)
				var _856_v180 _dafny.Sequence
				_ = _856_v180
				_856_v180 = _dafny.SeqOf(_739_v99, _739_v99, _739_v99)
				var _rhs137 _dafny.Set = _854_v178
				_ = _rhs137
				var _rhs138 D6 = _855_v179
				_ = _rhs138
				var _rhs139 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_856_v180, _856_v180), (Companion_Default___.SafeIndex((_849_v173).Select((Companion_Default___.SafeIndex(_this.F12(), _dafny.IntOfUint32((_849_v173).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_856_v180, _856_v180)).Cardinality()))).Uint32(), _739_v99)
				_ = _rhs139
				var _rhs140 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F12())), _this.F12())
				_ = _rhs140
				var _rhs141 _dafny.Int = _dafny.IntOfInt64(-338)
				_ = _rhs141
				var _lhs88 *GlobalState = globalState
				_ = _lhs88
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				_854_v178 = _rhs137
				_855_v179 = _rhs138
				_856_v180 = _rhs139
				_lhs88.F5 = _rhs140
				_lhs89.F5 = _rhs141
			}
			var _857_v181 _dafny.Array
			_ = _857_v181
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_23
			var _nw151 _dafny.Array
			_ = _nw151
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw151 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) _dafny.Sequence = (func(_858_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_859_i13 _dafny.Int) _dafny.Sequence {
						return _858_p2
					}
				})(p2)
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw151 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw151).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw151).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_857_v181 = _nw151
			_857_v181 = (func() _dafny.Array {
				if p1 {
					return _857_v181
				}
				return _857_v181
			})()
			var _860_v182 _dafny.Sequence
			_ = _860_v182
			_860_v182 = _dafny.SeqOf(p1, p1)
			var _861_v183 _dafny.Set
			_ = _861_v183
			_861_v183 = _dafny.SetOf(p1)
			var _862_v184 _dafny.Sequence
			_ = _862_v184
			_862_v184 = _dafny.SeqOf(p0, p0)
			_860_v182 = Companion_Default___.Fm9(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(((_838_v162).Update(_this.F12(), _dafny.IntOfInt64(105))).Cardinality(), p0, (_dafny.Zero).Minus((_861_v183).Cardinality())), _862_v184), globalState)
		}
		var _863_v185 _dafny.MultiSet
		_ = _863_v185
		_863_v185 = _dafny.MultiSetOf(p1)
		var _864_v186 _dafny.Map
		_ = _864_v186
		_864_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), _863_v185)
		var _865_v187 _dafny.Set
		_ = _865_v187
		_865_v187 = _dafny.SetOf(true, p1)
		r0 = !(Companion_Default___.Fm1(_this.F12(), (_863_v185).IsSubsetOf((func() _dafny.MultiSet {
			if (_864_v186).Contains(p1) {
				return (_864_v186).Get(p1).(_dafny.MultiSet)
			}
			return _863_v185
		})()), (_865_v187).Cardinality(), false, globalState))
		r0 = ((_688_v61).Cardinality()).Cmp(p0) > 0
		r1 = !((p1) == (p1)) || (!(p1) || (p1))
		r2 = Companion_Default___.SafeDivisionInt(p0, _this.F12())
		return r0, r1, r2
	}
}
func (_this *C3) F14() _dafny.Array {
	{
		return _this._f14
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	F11  _dafny.Int
	_f10 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this.F11 = _dafny.Zero
	_this._f10 = _dafny.Zero
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

func (_this *C4) Ctor__(f10 _dafny.Int, f11 _dafny.Int) {
	{
		(_this)._f10 = f10
		(_this).F11 = f11
	}
}
func (_this *C4) Fm2(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F10()
	}
}
func (_this *C4) M1(p0 _dafny.Sequence, p1 _dafny.CodePoint, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, _dafny.Set, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Set = _dafny.EmptySet
		_ = r2
		var r3 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r3
		var _866_v1 _dafny.Set
		_ = _866_v1
		_866_v1 = _dafny.SetOf(_dafny.IntOfInt64(999), (_this).F10(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("stclb")).Cardinality()), (_dafny.Zero).Minus((_this).F10()), _this.F11)
		var _867_v2 D0
		_ = _867_v2
		_867_v2 = Companion_D0_.Create_DC2_(p3, func() _dafny.Map {
			var _coll22 = _dafny.NewMapBuilder()
			_ = _coll22
			for _iter22 := _dafny.Iterate((_866_v1).Elements()); ; {
				_compr_22, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _868_v0 _dafny.Int
				_868_v0 = interface{}(_compr_22).(_dafny.Int)
				if (_866_v1).Contains(_868_v0) {
					_coll22.Add((_868_v0).Times((_this).F10()), true)
				}
			}
			return _coll22.ToMap()
		}(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)).Update(p3, p1), p2)
		var _source12 D0 = _867_v2
		_ = _source12
		if _source12.Is_DC1() {
			var _869___mcc_h0 _dafny.Array = _source12.Get_().(D0_DC1).Cf1
			_ = _869___mcc_h0
			var _870_cf1 _dafny.Array = _869___mcc_h0
			_ = _870_cf1
			r1 = _870_cf1
			if (func() bool {
				if p3 {
					return p3
				}
				return p2
			})() {
				var _871_v3 _dafny.MultiSet
				_ = _871_v3
				_871_v3 = _dafny.MultiSetOf(_this.F11)
				var _872_v4 D2
				_ = _872_v4
				_872_v4 = Companion_D2_.Create_DC8_(p3, (_871_v3).Cardinality(), p2)
				var _873_v5 _dafny.Map
				_ = _873_v5
				_873_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p3)
				var _874_v6 _dafny.Map
				_ = _874_v6
				_874_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F11)
				var _875_v7 _dafny.Sequence
				_ = _875_v7
				_875_v7 = _dafny.SeqOf(_this.F11)
				var _876_v8 _dafny.Map
				_ = _876_v8
				_876_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, p1)
				var _877_v9 _dafny.Array
				_ = _877_v9
				var _nwElement0_33 _dafny.Int = _this.F11
				_ = _nwElement0_33
				var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(27))
				_ = _nw152
				(_nw152).ArraySet1(_nwElement0_33, 0)
				(_nw152).ArraySet1(Companion_Default___.Fm0((_this).F10(), (_this).F10(), globalState), 1)
				(_nw152).ArraySet1((_this).F10(), 2)
				(_nw152).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((p0).Cardinality())), 3)
				(_nw152).ArraySet1(_dafny.IntOfInt64(528), 4)
				(_nw152).ArraySet1((_this).F10(), 5)
				(_nw152).ArraySet1((_dafny.Zero).Minus((_872_v4).Dtor_cf12()), 6)
				(_nw152).ArraySet1((_this).F10(), 7)
				(_nw152).ArraySet1(((func() _dafny.Map {
					if p2 {
						return _873_v5
					}
					return Companion_Default___.Fm3(p2, p3, (_874_v6).Cardinality(), globalState)
				})()).Cardinality(), 8)
				(_nw152).ArraySet1(_this.F11, 9)
				(_nw152).ArraySet1((_this).F10(), 10)
				(_nw152).ArraySet1(_dafny.IntOfInt64(171), 11)
				(_nw152).ArraySet1((_this).F10(), 12)
				(_nw152).ArraySet1(_dafny.IntOfInt64(-520), 13)
				(_nw152).ArraySet1((_this).F10(), 14)
				(_nw152).ArraySet1((_875_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.IntOfUint32((_875_v7).Cardinality()))).Uint32()).(_dafny.Int), 15)
				(_nw152).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, _dafny.UnicodeSeqOfUtf8Bytes("tup"))).Cardinality()), 16)
				(_nw152).ArraySet1((_876_v8).Cardinality(), 17)
				(_nw152).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(106), _this.F11), 18)
				(_nw152).ArraySet1(((_this).F10()).Minus(_this.F11), 19)
				(_nw152).ArraySet1(_this.F11, 20)
				(_nw152).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F11, _this.F11), 21)
				(_nw152).ArraySet1((_this).F10(), 22)
				(_nw152).ArraySet1((func() _dafny.Int {
					if p3 {
						return _this.F11
					}
					return _this.F11
				})(), 23)
				(_nw152).ArraySet1((_this).F10(), 24)
				(_nw152).ArraySet1(_this.F11, 25)
				(_nw152).ArraySet1(_this.F11, 26)
				_877_v9 = _nw152
				var _878_v10 D2
				_ = _878_v10
				_878_v10 = Companion_D2_.Create_DC9_(_dafny.IntOfInt64(779))
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_877_v9), 0))
				_ = _index161
				(_877_v9).ArraySet1((_878_v10).Dtor_cf14(), (_index161).Int())
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_877_v9), 0))
				_ = _index162
				(_877_v9).ArraySet1(_this.F11, (_index162).Int())
				var _879_v11 bool
				_ = _879_v11
				_879_v11 = true
				_879_v11 = false
				var _880_v12 _dafny.Set
				_ = _880_v12
				_880_v12 = _dafny.SetOf(_871_v3, _871_v3, _dafny.MultiSetOf((_this).F10()))
				var _rhs142 _dafny.Set = _880_v12
				_ = _rhs142
				var _rhs143 bool = false
				_ = _rhs143
				var _rhs144 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F10(), _this.F11)
				_ = _rhs144
				var _rhs145 bool = !(p2) || (_879_v11)
				_ = _rhs145
				_880_v12 = _rhs142
				_879_v11 = _rhs143
				r0 = _rhs144
				_879_v11 = _rhs145
				var _881_v13 _dafny.Set
				_ = _881_v13
				_881_v13 = _dafny.SetOf(p2)
				var _882_v14 *C2
				_ = _882_v14
				var _nw153 *C2 = New_C2_()
				_ = _nw153
				_nw153.Ctor__((_881_v13).IsSubsetOf(_881_v13), (_877_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_877_v9), 0))).Int()).(_dafny.Int))
				_882_v14 = _nw153
				(globalState).F5 = (_877_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_877_v9), 0))).Int()).(_dafny.Int)
			} else {
				var _883_v15 bool
				_ = _883_v15
				_883_v15 = false
				var _884_v16 _dafny.MultiSet
				_ = _884_v16
				_884_v16 = _dafny.MultiSetOf(_870_cf1, _870_cf1)
				_883_v15 = !((_884_v16).IsDisjointFrom(_dafny.MultiSetOf(_870_cf1, _870_cf1, _870_cf1))) || (p2)
				var _885_v17 _dafny.Array
				_ = _885_v17
				var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw154
				_885_v17 = _nw154
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_885_v17), 0))
				_ = _index163
				(_885_v17).ArraySet1(_this.F11, (_index163).Int())
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_885_v17), 0))
				_ = _index164
				(_885_v17).ArraySet1(_this.F11, (_index164).Int())
				_883_v15 = ((func() _dafny.Int {
					if _883_v15 {
						return _this.F11
					}
					return (_this).F10()
				})()).Cmp(_dafny.IntOfUint32((p0).Cardinality())) < 0
				(globalState).F5 = Companion_Default___.SafeModuloInt((_885_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_885_v17), 0))).Int()).(_dafny.Int), _this.F11)
				_883_v15 = (p2) || (!_dafny.Companion_Sequence_.Equal(p0, p0))
			}
			var _886_v18 D5
			_ = _886_v18
			_886_v18 = Companion_D5_.Create_DC13_(true, false)
			var _887_v19 _dafny.MultiSet
			_ = _887_v19
			_887_v19 = _dafny.MultiSetOf(p3, (_886_v18).Dtor_cf19())
			var _888_v20 _dafny.Sequence
			_ = _888_v20
			_888_v20 = _dafny.SeqOf(_dafny.IntOfInt64(681), _this.F11)
			var _889_v21 _dafny.Sequence
			_ = _889_v21
			_889_v21 = _dafny.SeqOf(_887_v19, _887_v19, (_887_v19).Update(p2, Companion_Default___.Abs(_dafny.IntOfUint32((_888_v20).Cardinality()))), _887_v19, _887_v19)
			var _source13 D5 = Companion_D5_.Create_DC12_(_889_v21)
			_ = _source13
			if _source13.Is_DC13() {
				var _890___mcc_h7 bool = _source13.Get_().(D5_DC13).Cf18
				_ = _890___mcc_h7
				var _891___mcc_h8 bool = _source13.Get_().(D5_DC13).Cf19
				_ = _891___mcc_h8
				var _892_cf19 bool = _891___mcc_h8
				_ = _892_cf19
				var _893_cf18 bool = _890___mcc_h7
				_ = _893_cf18
				var _894_v22 *C1
				_ = _894_v22
				var _nw155 *C1 = New_C1_()
				_ = _nw155
				_nw155.Ctor__(p1, p1, (_this).F10())
				_894_v22 = _nw155
				var _rhs146 *C1 = _894_v22
				_ = _rhs146
				var _rhs147 bool = _892_cf19
				_ = _rhs147
				_894_v22 = _rhs146
				_893_cf18 = _rhs147
				(globalState).F5 = _dafny.IntOfInt64(115)
				var _895_v23 _dafny.MultiSet
				_ = _895_v23
				_895_v23 = _dafny.MultiSetOf(_this.F11, _dafny.IntOfUint32((p0).Cardinality()))
				var _896_v24 D5
				_ = _896_v24
				_896_v24 = Companion_D5_.Create_DC12_(_dafny.SeqOf(_887_v19, _887_v19))
				(_this).F11 = (Companion_Default___.Fm20((_this).Fm2(_893_cf18, globalState), _this.F11, (_895_v23).Cardinality(), _896_v24, globalState)).Cardinality()
				(globalState).F5 = (_887_v19).Cardinality()
			} else if _source13.Is_DC14() {
				var _897___mcc_h9 bool = _source13.Get_().(D5_DC14).Cf20
				_ = _897___mcc_h9
				var _898_cf20 bool = _897___mcc_h9
				_ = _898_cf20
				var _899_v25 _dafny.Array
				_ = _899_v25
				var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw156
				_899_v25 = _nw156
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_899_v25), 0))
				_ = _index165
				(_899_v25).ArraySet1(p0, (_index165).Int())
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_899_v25), 0))
				_ = _index166
				(_899_v25).ArraySet1(p0, (_index166).Int())
				var _900_v26 *C1
				_ = _900_v26
				var _nw157 *C1 = New_C1_()
				_ = _nw157
				_nw157.Ctor__(p1, p1, _dafny.IntOfUint32(((_899_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_899_v25), 0))).Int()).(_dafny.Sequence)).Cardinality()))
				_900_v26 = _nw157
				_900_v26 = (func() *C1 {
					if false {
						return _900_v26
					}
					return _900_v26
				})()
				var _901_v27 _dafny.Array
				_ = _901_v27
				var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw158
				_901_v27 = _nw158
				r3 = _901_v27
				var _902_v28 _dafny.MultiSet
				_ = _902_v28
				_902_v28 = _dafny.MultiSetOf(p0, _dafny.UnicodeSeqOfUtf8Bytes("ed"))
				var _903_v29 _dafny.Map
				_ = _903_v29
				_903_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_870_cf1, !(_902_v28).Contains(p0))
				_903_v29 = (_903_v29).Update(_870_cf1, Companion_Default___.Fm1((_this).F10(), !(_898_cf20), (_this).F10(), p3, globalState))
			} else {
				var _904___mcc_h10 _dafny.Sequence = _source13.Get_().(D5_DC12).Cf17
				_ = _904___mcc_h10
				var _905_cf17 _dafny.Sequence = _904___mcc_h10
				_ = _905_cf17
				var _906_v30 *C1
				_ = _906_v30
				var _nw159 *C1 = New_C1_()
				_ = _nw159
				_nw159.Ctor__(p1, p1, _this.F11)
				_906_v30 = _nw159
				r0 = _this.F11
				var _907_v31 bool
				_ = _907_v31
				_907_v31 = true
				_907_v31 = p2
				var _908_v32 _dafny.Array
				_ = _908_v32
				var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
				_ = _nw160
				_908_v32 = _nw160
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_908_v32), 0))
				_ = _index167
				(_908_v32).ArraySet1(p0, (_index167).Int())
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_908_v32), 0))
				_ = _index168
				(_908_v32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), (_index168).Int())
			}
			var _909_v33 _dafny.Array
			_ = _909_v33
			var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw161
			_909_v33 = _nw161
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_909_v33), 0))
			_ = _index169
			(_909_v33).ArraySet1((_this).F10(), (_index169).Int())
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_909_v33), 0))
			_ = _index170
			(_909_v33).ArraySet1(_this.F11, (_index170).Int())
		} else if _source12.Is_DC2() {
			var _910___mcc_h1 bool = _source12.Get_().(D0_DC2).Cf2
			_ = _910___mcc_h1
			var _911___mcc_h2 _dafny.Map = _source12.Get_().(D0_DC2).Cf3
			_ = _911___mcc_h2
			var _912___mcc_h3 _dafny.Map = _source12.Get_().(D0_DC2).Cf4
			_ = _912___mcc_h3
			var _913___mcc_h4 bool = _source12.Get_().(D0_DC2).Cf5
			_ = _913___mcc_h4
			var _914_cf5 bool = _913___mcc_h4
			_ = _914_cf5
			var _915_cf4 _dafny.Map = _912___mcc_h3
			_ = _915_cf4
			var _916_cf3 _dafny.Map = _911___mcc_h2
			_ = _916_cf3
			var _917_cf2 bool = _910___mcc_h1
			_ = _917_cf2
			var _918_v34 D7
			_ = _918_v34
			_918_v34 = Companion_D7_.Create_DC21_()
			_918_v34 = Companion_D7_.Create_DC21_()
			if false {
				_914_cf5 = _914_cf5
				var _919_v35 _dafny.Sequence
				_ = _919_v35
				_919_v35 = _dafny.SeqOf(_this.F11, _this.F11, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (Companion_D2_.Create_DC7_((_this).F10())).Dtor_cf10())).Cardinality(), _dafny.IntOfInt64(533))
				var _920_v36 *C2
				_ = _920_v36
				var _nw162 *C2 = New_C2_()
				_ = _nw162
				_nw162.Ctor__(p3, (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_919_v35).Cardinality()), (_dafny.Zero).Minus(_this.F11))))
				_920_v36 = _nw162
				var _921_v37 _dafny.Map
				_ = _921_v37
				_921_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_920_v36).F17())
				var _922_v38 _dafny.MultiSet
				_ = _922_v38
				_922_v38 = _dafny.MultiSetOf((func() bool {
					if (_921_v37).Contains(_dafny.CodePoint('a')) {
						return (_921_v37).Get(_dafny.CodePoint('a')).(bool)
					}
					return !(p2)
				})(), _914_cf5)
				(globalState).F5 = (func() _dafny.Int {
					if (_922_v38).Contains(((_this).F10()).Cmp(_this.F11) < 0) {
						return (_922_v38).Multiplicity(((_this).F10()).Cmp(_this.F11) < 0)
					}
					return (_916_cf3).Cardinality()
				})()
				var _923_v39 _dafny.Array
				_ = _923_v39
				var _nwElement0_34 bool = _914_cf5
				_ = _nwElement0_34
				var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(28))
				_ = _nw163
				(_nw163).ArraySet1(_nwElement0_34, 0)
				(_nw163).ArraySet1(p3, 1)
				(_nw163).ArraySet1(_917_cf2, 2)
				(_nw163).ArraySet1((_920_v36).F17(), 3)
				(_nw163).ArraySet1(p3, 4)
				(_nw163).ArraySet1(_914_cf5, 5)
				(_nw163).ArraySet1((_920_v36).F17(), 6)
				(_nw163).ArraySet1(_914_cf5, 7)
				(_nw163).ArraySet1(p3, 8)
				(_nw163).ArraySet1((_920_v36).F17(), 9)
				(_nw163).ArraySet1((_920_v36).F17(), 10)
				(_nw163).ArraySet1(_917_cf2, 11)
				(_nw163).ArraySet1(_914_cf5, 12)
				(_nw163).ArraySet1(p2, 13)
				(_nw163).ArraySet1(true, 14)
				(_nw163).ArraySet1(p2, 15)
				(_nw163).ArraySet1(_917_cf2, 16)
				(_nw163).ArraySet1(_914_cf5, 17)
				(_nw163).ArraySet1(_914_cf5, 18)
				(_nw163).ArraySet1(p3, 19)
				(_nw163).ArraySet1(true, 20)
				(_nw163).ArraySet1(p2, 21)
				(_nw163).ArraySet1(p3, 22)
				(_nw163).ArraySet1(false, 23)
				(_nw163).ArraySet1(p3, 24)
				(_nw163).ArraySet1(p2, 25)
				(_nw163).ArraySet1(p2, 26)
				(_nw163).ArraySet1(p2, 27)
				_923_v39 = _nw163
				var _924_v40 T1
				_ = _924_v40
				var _nw164 *C1 = New_C1_()
				_ = _nw164
				_nw164.Ctor__(p1, _dafny.CodePoint('l'), (Companion_D6_.Create_DC16_(_923_v39, (_this).F10())).Dtor_cf23())
				_924_v40 = _nw164
				var _925_v41 _dafny.Array
				_ = _925_v41
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_24
				var _nw165 _dafny.Array
				_ = _nw165
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw165 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Sequence = (func(_926_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_927_i0 _dafny.Int) _dafny.Sequence {
							return _926_p0
						}
					})(p0)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw165 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw165).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw165).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_925_v41 = _nw165
				var _928_v42 _dafny.Sequence
				_ = _928_v42
				_928_v42 = _dafny.SeqOf(_924_v40, _924_v40)
				var _rhs148 T1 = (_928_v42).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-414), _dafny.IntOfUint32((_928_v42).Cardinality()))).Uint32()).(T1)
				_ = _rhs148
				var _rhs149 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(302))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_929_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_930_i2 _dafny.Int) _dafny.Sequence {
						return _929_p0
					}
				})(p0))), _dafny.SeqOf(p0, p0)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_931_v40 T1) func(_dafny.Int) _dafny.CodePoint {
					return func(_932_i1 _dafny.Int) _dafny.CodePoint {
						return (_931_v40).F13()
					}
				})(_924_v40))), p0))
				_ = _rhs149
				var _rhs150 _dafny.Array = _925_v41
				_ = _rhs150
				var _rhs151 _dafny.Int = (_dafny.Zero).Minus((_this).F10())
				_ = _rhs151
				var _lhs90 T1 = _924_v40
				_ = _lhs90
				_924_v40 = _rhs148
				_914_cf5 = _rhs149
				_925_v41 = _rhs150
				_lhs90.F12_set_(_rhs151)
				var _933_v43 _dafny.MultiSet
				_ = _933_v43
				_933_v43 = _dafny.MultiSetOf(Companion_Default___.Fm11(_924_v40.F12(), _924_v40.F12(), p2, globalState))
				var _934_v44 _dafny.Array
				_ = _934_v44
				var _nw166 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(23))
				_ = _nw166
				_934_v44 = _nw166
				var _935_v45 _dafny.Set
				_ = _935_v45
				_935_v45 = _dafny.SetOf((_924_v40).F13())
				var _936_v46 _dafny.Map
				_ = _936_v46
				_936_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_917_cf2, Companion_Default___.Fm0((_935_v45).Cardinality(), _this.F11, globalState))
				var _937_v47 *C3
				_ = _937_v47
				var _nw167 *C3 = New_C3_()
				_ = _nw167
				_nw167.Ctor__(_934_v44, p1, (func() _dafny.Int {
					if (_936_v46).Contains(p2) {
						return (_936_v46).Get(p2).(_dafny.Int)
					}
					return _this.F11
				})())
				_937_v47 = _nw167
				var _938_v48 _dafny.Map
				_ = _938_v48
				_938_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D8_.Create_DC22_(_933_v43)).Dtor_cf28(), _937_v47)
				_938_v48 = (_938_v48).Update(_dafny.MultiSetOf(p0, p0, p0, p0, p0), _937_v47)
			} else {
				var _939_v49 _dafny.Array
				_ = _939_v49
				var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(8))
				_ = _nw168
				_939_v49 = _nw168
				var _940_v50 _dafny.Set
				_ = _940_v50
				_940_v50 = _dafny.SetOf(p2, _914_cf5)
				var _941_v51 _dafny.MultiSet
				_ = _941_v51
				_941_v51 = _dafny.MultiSetOf((_940_v50).Cardinality())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_939_v49), 0))
				_ = _index171
				(_939_v49).ArraySet1(_941_v51, (_index171).Int())
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_939_v49), 0))
				_ = _index172
				(_939_v49).ArraySet1(_941_v51, (_index172).Int())
				var _942_v52 _dafny.Array
				_ = _942_v52
				var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw169
				_942_v52 = _nw169
				var _943_v53 _dafny.Map
				_ = _943_v53
				_943_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _this.F11)
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_942_v52), 0))
				_ = _index173
				(_942_v52).ArraySet1(((_943_v53).Update(p2, (_dafny.Zero).Minus(_this.F11))).Cardinality(), (_index173).Int())
				var _944_v54 _dafny.Map
				_ = _944_v54
				_944_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), (_this).F10())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_942_v52), 0))
				_ = _index174
				(_942_v52).ArraySet1(((_this.F11).Minus((_944_v54).Cardinality())).Times(_this.F11), (_index174).Int())
				_917_cf2 = _917_cf2
				_917_cf2 = Companion_Default___.Fm1(_this.F11, p3, (_this).F10(), _914_cf5, globalState)
				var _945_v55 _dafny.Set
				_ = _945_v55
				_945_v55 = _dafny.SetOf(_942_v52, _942_v52, _942_v52)
				var _946_v56 D5
				_ = _946_v56
				_946_v56 = Companion_D5_.Create_DC13_(true, ((_dafny.Zero).Minus(_this.F11)).Cmp((_945_v55).Cardinality()) != 0)
				_946_v56 = _946_v56
			}
			(globalState).F5 = (Companion_Default___.Fm10(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-926), _this.F11), globalState)).Cardinality()
			var _947_v57 _dafny.Sequence
			_ = _947_v57
			_947_v57 = _dafny.SeqOf(_917_cf2)
			var _948_v58 _dafny.Map
			_ = _948_v58
			_948_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_947_v57, _dafny.SetOf(p3))
			var _949_v59 _dafny.Set
			_ = _949_v59
			_949_v59 = _dafny.SetOf(p2, p3)
			_914_cf5 = (_949_v59).IsProperSubsetOf(((func() _dafny.Set {
				if (_948_v58).Contains(_947_v57) {
					return (_948_v58).Get(_947_v57).(_dafny.Set)
				}
				return _949_v59
			})()).Difference(_dafny.SetOf(p3, p3)))
		} else if _source12.Is_DC0() {
			var _950___mcc_h5 bool = _source12.Get_().(D0_DC0).Cf0
			_ = _950___mcc_h5
			var _951_cf0 bool = _950___mcc_h5
			_ = _951_cf0
			var _952_v60 _dafny.MultiSet
			_ = _952_v60
			_952_v60 = _dafny.MultiSetOf(_866_v1)
			_951_cf0 = (_952_v60).IsProperSubsetOf(_952_v60)
			var _953_v61 _dafny.Array
			_ = _953_v61
			var _nwElement0_35 _dafny.Int = _this.F11
			_ = _nwElement0_35
			var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(8))
			_ = _nw170
			(_nw170).ArraySet1(_nwElement0_35, 0)
			(_nw170).ArraySet1(_this.F11, 1)
			(_nw170).ArraySet1((_dafny.Zero).Minus(_this.F11), 2)
			(_nw170).ArraySet1((_866_v1).Cardinality(), 3)
			(_nw170).ArraySet1((_this).F10(), 4)
			(_nw170).ArraySet1(_this.F11, 5)
			(_nw170).ArraySet1((_this).F10(), 6)
			(_nw170).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _this.F11)).Cardinality(), 7)
			_953_v61 = _nw170
			var _954_v62 _dafny.Array
			_ = _954_v62
			var _nwElement0_36 bool = _951_cf0
			_ = _nwElement0_36
			var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(16))
			_ = _nw171
			(_nw171).ArraySet1(_nwElement0_36, 0)
			(_nw171).ArraySet1(false, 1)
			(_nw171).ArraySet1(p3, 2)
			(_nw171).ArraySet1(p3, 3)
			(_nw171).ArraySet1(p2, 4)
			(_nw171).ArraySet1(p2, 5)
			(_nw171).ArraySet1(p2, 6)
			(_nw171).ArraySet1(p3, 7)
			(_nw171).ArraySet1(p2, 8)
			(_nw171).ArraySet1(Companion_Default___.Fm1((_this).F10(), p3, (_this).F10(), p2, globalState), 9)
			(_nw171).ArraySet1(p3, 10)
			(_nw171).ArraySet1(p2, 11)
			(_nw171).ArraySet1(_951_cf0, 12)
			(_nw171).ArraySet1(_951_cf0, 13)
			(_nw171).ArraySet1(p3, 14)
			(_nw171).ArraySet1(_951_cf0, 15)
			_954_v62 = _nw171
			var _955_v63 _dafny.Map
			_ = _955_v63
			_955_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_953_v61, _954_v62)
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
			_ = _index175
			(_954_v62).ArraySet1(p2, (_index175).Int())
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
			_ = _index176
			var _rhs152 _dafny.Map = ((_955_v63).Update(_953_v61, _954_v62)).Merge(_955_v63)
			_ = _rhs152
			var _rhs153 bool = p2
			_ = _rhs153
			var _lhs91 _dafny.Array = _954_v62
			_ = _lhs91
			var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
			_ = _lhs92
			_955_v63 = _rhs152
			(_lhs91).ArraySet1(_rhs153, (_lhs92).Int())
			var _956_v64 _dafny.MultiSet
			_ = _956_v64
			_956_v64 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("hs"), _dafny.UnicodeSeqOfUtf8Bytes("yn"))
			var _957_v65 D8
			_ = _957_v65
			_957_v65 = Companion_D8_.Create_DC22_(_956_v64)
			var _source14 D8 = _957_v65
			_ = _source14
			if _source14.Is_DC23() {
				var _958___mcc_h11 _dafny.Int = _source14.Get_().(D8_DC23).Cf29
				_ = _958___mcc_h11
				var _959_cf29 _dafny.Int = _958___mcc_h11
				_ = _959_cf29
				var _960_v66 _dafny.Map
				_ = _960_v66
				_960_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _953_v61)
				var _961_v67 T1
				_ = _961_v67
				var _nw172 *C1 = New_C1_()
				_ = _nw172
				_nw172.Ctor__(p1, p1, _959_cf29)
				_961_v67 = _nw172
				var _pat_let_tv6 = _960_v66
				_ = _pat_let_tv6
				var _962_v68 _dafny.Map
				_ = _962_v68
				_962_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let13_0 D7) D7 {
					return func(_963_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let14_0 _dafny.Map) D7 {
							return func(_964_dt__update_hcf27_h0 _dafny.Map) D7 {
								return Companion_D7_.Create_DC20_(_964_dt__update_hcf27_h0)
							}(_pat_let14_0)
						}(_pat_let_tv6)
					}(_pat_let13_0)
				}(Companion_D7_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_954_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))).Int()).(bool), _953_v61))), _961_v67)
				var _965_v69 _dafny.Map
				_ = _965_v69
				_965_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_962_v68, !(p3))
				var _966_v70 _dafny.Map
				_ = _966_v70
				_966_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_959_cf29, _965_v69)
				var _967_v71 D2
				_ = _967_v71
				_967_v71 = Companion_D2_.Create_DC7_(_961_v67.F12())
				var _968_v72 _dafny.Map
				_ = _968_v72
				_968_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_954_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))).Int()).(bool), (_954_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))).Int()).(bool))
				var _pat_let_tv7 = _959_cf29
				_ = _pat_let_tv7
				var _969_v73 _dafny.Array
				_ = _969_v73
				var _nwElement0_37 D2 = _967_v71
				_ = _nwElement0_37
				var _nw173 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(15))
				_ = _nw173
				(_nw173).ArraySet1(_nwElement0_37, 0)
				(_nw173).ArraySet1(func(_pat_let15_0 D2) D2 {
					return func(_970_dt__update__tmp_h1 D2) D2 {
						return func(_pat_let16_0 _dafny.Int) D2 {
							return func(_971_dt__update_hcf10_h0 _dafny.Int) D2 {
								return Companion_D2_.Create_DC7_(_971_dt__update_hcf10_h0)
							}(_pat_let16_0)
						}(_pat_let_tv7)
					}(_pat_let15_0)
				}(_967_v71), 1)
				(_nw173).ArraySet1(_967_v71, 2)
				(_nw173).ArraySet1(Companion_D2_.Create_DC7_((_968_v72).Cardinality()), 3)
				(_nw173).ArraySet1(_967_v71, 4)
				(_nw173).ArraySet1(_967_v71, 5)
				(_nw173).ArraySet1(Companion_D2_.Create_DC7_((_this).F10()), 6)
				(_nw173).ArraySet1(_967_v71, 7)
				(_nw173).ArraySet1(_967_v71, 8)
				(_nw173).ArraySet1(_967_v71, 9)
				(_nw173).ArraySet1(_967_v71, 10)
				(_nw173).ArraySet1(_967_v71, 11)
				(_nw173).ArraySet1(Companion_D2_.Create_DC7_(_dafny.IntOfInt64(539)), 12)
				(_nw173).ArraySet1(_967_v71, 13)
				(_nw173).ArraySet1(_967_v71, 14)
				_969_v73 = _nw173
				var _972_v74 *C3
				_ = _972_v74
				var _nw174 *C3 = New_C3_()
				_ = _nw174
				_nw174.Ctor__(_969_v73, p1, _dafny.IntOfInt64(-596))
				_972_v74 = _nw174
				var _973_v75 _dafny.Sequence
				_ = _973_v75
				_973_v75 = _dafny.SeqOf(_972_v74)
				_966_v70 = (_966_v70).Update(Companion_Default___.SafeDivisionInt(_959_cf29, _dafny.IntOfUint32((_973_v75).Cardinality())), _965_v69)
				var _974_v76 D7
				_ = _974_v76
				_974_v76 = Companion_D7_.Create_DC20_(_960_v66)
				var _975_v77 _dafny.Sequence
				_ = _975_v77
				_975_v77 = _dafny.SeqOf(_974_v76)
				_975_v77 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_975_v77, _975_v77), _975_v77), (Companion_Default___.SafeIndex(_961_v67.F12(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_975_v77, _975_v77), _975_v77)).Cardinality()))).Uint32(), _974_v76)
				var _976_v78 _dafny.Array
				_ = _976_v78
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_25
				var _nw175 _dafny.Array
				_ = _nw175
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw175 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.Set = (func(_977_p3 bool) func(_dafny.Int) _dafny.Set {
						return func(_978_i3 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_977_p3)
						}
					})(p3)
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw175 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw175).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw175).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_976_v78 = _nw175
				var _979_v79 *C3
				_ = _979_v79
				var _nw176 *C3 = New_C3_()
				_ = _nw176
				_nw176.Ctor__((_972_v74).F14(), p1, (_dafny.MultiSetOf(_976_v78, _976_v78)).Cardinality())
				_979_v79 = _nw176
				var _980_v80 _dafny.CodePoint
				_ = _980_v80
				_980_v80 = _dafny.CodePoint('u')
				var _981_v81 _dafny.Map
				_ = _981_v81
				_981_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_961_v67).F13())
				_980_v80 = (func() _dafny.CodePoint {
					if (_981_v81).Contains((func() bool {
						if (_954_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))).Int()).(bool) {
							return p2
						}
						return p2
					})()) {
						return (_981_v81).Get((func() bool {
							if (_954_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))).Int()).(bool) {
								return p2
							}
							return p2
						})()).(_dafny.CodePoint)
					}
					return p1
				})()
			} else if _source14.Is_DC24() {
				var _982___mcc_h12 _dafny.Set = _source14.Get_().(D8_DC24).Cf30
				_ = _982___mcc_h12
				var _983___mcc_h13 _dafny.Sequence = _source14.Get_().(D8_DC24).Cf31
				_ = _983___mcc_h13
				var _984_cf31 _dafny.Sequence = _983___mcc_h13
				_ = _984_cf31
				var _985_cf30 _dafny.Set = _982___mcc_h12
				_ = _985_cf30
				var _986_v82 D2
				_ = _986_v82
				_986_v82 = Companion_D2_.Create_DC7_(_this.F11)
				var _987_v83 _dafny.Array
				_ = _987_v83
				var _nwElement0_38 D2 = _986_v82
				_ = _nwElement0_38
				var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(13))
				_ = _nw177
				(_nw177).ArraySet1(_nwElement0_38, 0)
				(_nw177).ArraySet1(_986_v82, 1)
				(_nw177).ArraySet1(_986_v82, 2)
				(_nw177).ArraySet1(_986_v82, 3)
				(_nw177).ArraySet1(_986_v82, 4)
				(_nw177).ArraySet1(_986_v82, 5)
				(_nw177).ArraySet1(_986_v82, 6)
				(_nw177).ArraySet1(_986_v82, 7)
				(_nw177).ArraySet1(_986_v82, 8)
				(_nw177).ArraySet1(_986_v82, 9)
				(_nw177).ArraySet1(_986_v82, 10)
				(_nw177).ArraySet1(_986_v82, 11)
				(_nw177).ArraySet1(_986_v82, 12)
				_987_v83 = _nw177
				var _988_v84 D9
				_ = _988_v84
				_988_v84 = Companion_D9_.Create_DC25_(p1)
				var _989_v85 *C3
				_ = _989_v85
				var _nw178 *C3 = New_C3_()
				_ = _nw178
				_nw178.Ctor__(_987_v83, (_988_v84).Dtor_cf32(), (_dafny.SetOf(_951_cf0)).Cardinality())
				_989_v85 = _nw178
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
				_ = _index177
				(_954_v62).ArraySet1(true, (_index177).Int())
				var _990_v86 _dafny.Map
				_ = _990_v86
				_990_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(!(p3)))
				var _991_v87 _dafny.Sequence
				_ = _991_v87
				_991_v87 = _dafny.SeqOf((func() _dafny.Int {
					if !((_954_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))).Int()).(bool)) {
						return (_990_v86).Cardinality()
					}
					return _this.F11
				})())
				(globalState).F5 = (_991_v87).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F10())), _dafny.IntOfUint32((_991_v87).Cardinality()))).Uint32()).(_dafny.Int)
				var _992_v88 D6
				_ = _992_v88
				_992_v88 = Companion_D6_.Create_DC18_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_951_cf0), p2)).Cardinality())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
				_ = _index178
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
				_ = _index179
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
				_ = _index180
				var _rhs154 bool = (_954_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))).Int()).(bool)
				_ = _rhs154
				var _rhs155 _dafny.Int = (_992_v88).Dtor_cf25()
				_ = _rhs155
				var _rhs156 bool = Companion_Default___.Fm1((_this).F10(), _951_cf0, (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_this.F11, _dafny.IntOfInt64(250))), p2, globalState)
				_ = _rhs156
				var _rhs157 bool = p2
				_ = _rhs157
				var _lhs93 _dafny.Array = _954_v62
				_ = _lhs93
				var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
				_ = _lhs94
				var _lhs95 *C4 = _this
				_ = _lhs95
				var _lhs96 _dafny.Array = _954_v62
				_ = _lhs96
				var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
				_ = _lhs97
				var _lhs98 _dafny.Array = _954_v62
				_ = _lhs98
				var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_954_v62), 0))
				_ = _lhs99
				(_lhs93).ArraySet1(_rhs154, (_lhs94).Int())
				_lhs95.F11 = _rhs155
				(_lhs96).ArraySet1(_rhs156, (_lhs97).Int())
				(_lhs98).ArraySet1(_rhs157, (_lhs99).Int())
			} else {
				var _993___mcc_h14 _dafny.MultiSet = _source14.Get_().(D8_DC22).Cf28
				_ = _993___mcc_h14
				var _994_cf28 _dafny.MultiSet = _993___mcc_h14
				_ = _994_cf28
				var _995_v89 _dafny.Map
				_ = _995_v89
				_995_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _this.F11)
				var _996_v90 _dafny.Sequence
				_ = _996_v90
				_996_v90 = _dafny.SeqOf(_dafny.IntOfInt64(858))
				_995_v89 = (_995_v89).Update(_this.F11, (_996_v90).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_996_v90).Cardinality()))).Uint32()).(_dafny.Int))
				var _997_v91 T1
				_ = _997_v91
				var _nw179 *C1 = New_C1_()
				_ = _nw179
				_nw179.Ctor__((p0).Select((Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint), p1, (_this).F10())
				_997_v91 = _nw179
				var _998_v92 _dafny.Array
				_ = _998_v92
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_26
				var _nw180 _dafny.Array
				_ = _nw180
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw180 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) D2 = (func(_999_v91 T1, _1000_p0 _dafny.Sequence) func(_dafny.Int) D2 {
						return func(_1001_i4 _dafny.Int) D2 {
							return Companion_D2_.Create_DC7_((_dafny.SetOf(_999_v91.F12(), (_this).F10(), _999_v91.F12(), _dafny.IntOfUint32((_1000_p0).Cardinality()))).Cardinality())
						}
					})(_997_v91, p0)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw180 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw180).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw180).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_998_v92 = _nw180
				var _nw181 *C3 = New_C3_()
				_ = _nw181
				_nw181.Ctor__(_998_v92, (_997_v91).F13(), (_996_v90).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_996_v90).Cardinality()))).Uint32()).(_dafny.Int))
				_997_v91 = _nw181
				var _1002_v93 _dafny.CodePoint
				_ = _1002_v93
				_1002_v93 = _dafny.CodePoint('n')
				_1002_v93 = _dafny.CodePoint('a')
				_866_v1 = _866_v1
			}
			_954_v62 = _954_v62
		} else {
			var _1003___mcc_h6 D0 = _source12.Get_().(D0_DC3).Cf6
			_ = _1003___mcc_h6
			var _1004_cf6 D0 = _1003___mcc_h6
			_ = _1004_cf6
			var _1005_v94 _dafny.Array
			_ = _1005_v94
			var _nw182 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw182
			_1005_v94 = _nw182
			var _1006_v95 D0
			_ = _1006_v95
			_1006_v95 = Companion_D0_.Create_DC1_(_1005_v94)
			var _1007_v97 _dafny.Sequence
			_ = _1007_v97
			_1007_v97 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter23 := _dafny.Iterate((p0).Elements()); ; {
					_compr_23, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _1008_v96 _dafny.CodePoint
					_1008_v96 = interface{}(_compr_23).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(p0, _1008_v96) {
						_coll23.Add(_1008_v96, p2)
					}
				}
				return _coll23.ToMap()
			}()).Cardinality())).Cardinality()), _dafny.IntOfInt64(257))
			var _1009_v98 _dafny.Map
			_ = _1009_v98
			_1009_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1006_v95).Dtor_cf1(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1007_v97, _dafny.SeqOf(_this.F11))).Cardinality()))
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_1005_v94), 0))
			_ = _index181
			(_1005_v94).ArraySet1(p3, (_index181).Int())
			var _1010_v99 _dafny.MultiSet
			_ = _1010_v99
			_1010_v99 = _dafny.MultiSetOf(p0, _dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(447))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_1011_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1012_i5 _dafny.Int) _dafny.CodePoint {
					return _1011_p1
				}
			})(p1)))))
			var _1013_v100 _dafny.MultiSet
			_ = _1013_v100
			_1013_v100 = _dafny.MultiSetOf(_this.F11)
			var _1014_v101 _dafny.Set
			_ = _1014_v101
			_1014_v101 = _dafny.SetOf(_dafny.MultiSetFromSeq(_1007_v97), _1013_v100, _1013_v100)
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_1005_v94), 0))
			_ = _index182
			var _rhs158 _dafny.Map = _1009_v98
			_ = _rhs158
			var _rhs159 _dafny.Int = _this.F11
			_ = _rhs159
			var _rhs160 bool = !(_1014_v101).Equals((_1014_v101).Intersection(_1014_v101))
			_ = _rhs160
			var _rhs161 _dafny.MultiSet = _1010_v99
			_ = _rhs161
			var _lhs100 *GlobalState = globalState
			_ = _lhs100
			var _lhs101 _dafny.Array = _1005_v94
			_ = _lhs101
			var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_1005_v94), 0))
			_ = _lhs102
			_1009_v98 = _rhs158
			_lhs100.F5 = _rhs159
			(_lhs101).ArraySet1(_rhs160, (_lhs102).Int())
			_1010_v99 = _rhs161
			var _1015_v102 _dafny.CodePoint
			_ = _1015_v102
			_1015_v102 = _dafny.CodePoint('u')
			_1015_v102 = _1015_v102
			var _1016_v103 _dafny.Map
			_ = _1016_v103
			_1016_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _1017_v104 _dafny.Map
			_ = _1017_v104
			_1017_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1016_v103, _this.F11)
			_1017_v104 = (_1017_v104).Update(_1016_v103, (_this).F10())
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_1005_v94), 0))
			_ = _index183
			(_1005_v94).ArraySet1(p2, (_index183).Int())
		}
		var _1018_v105 *C1
		_ = _1018_v105
		var _nw183 *C1 = New_C1_()
		_ = _nw183
		_nw183.Ctor__(p1, p1, _dafny.IntOfInt64(554))
		_1018_v105 = _nw183
		var _1019_v106 _dafny.MultiSet
		_ = _1019_v106
		_1019_v106 = _dafny.MultiSetOf(_dafny.IntOfInt64(-902))
		var _1020_v107 _dafny.Map
		_ = _1020_v107
		_1020_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(146), _1019_v106)
		if ((_1019_v106).Update(_this.F11, Companion_Default___.Abs((_this).F10()))).IsSubsetOf((func() _dafny.MultiSet {
			if p3 {
				return _1019_v106
			}
			return (func() _dafny.MultiSet {
				if (_1020_v107).Contains(_this.F11) {
					return (_1020_v107).Get(_this.F11).(_dafny.MultiSet)
				}
				return _1019_v106
			})()
		})()) {
			var _1021_v108 _dafny.Map
			_ = _1021_v108
			_1021_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F10())
			var _1022_v109 _dafny.Map
			_ = _1022_v109
			_1022_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1021_v108).Cardinality())
			var _1023_v110 T0
			_ = _1023_v110
			var _nw184 *C2 = New_C2_()
			_ = _nw184
			_nw184.Ctor__(p3, (func() _dafny.Int {
				if (_1022_v109).Contains(p2) {
					return (_1022_v109).Get(p2).(_dafny.Int)
				}
				return _this.F11
			})())
			_1023_v110 = _nw184
			var _1024_v111 _dafny.Array
			_ = _1024_v111
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_27
			var _nw185 _dafny.Array
			_ = _nw185
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw185 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Int = (func(_1025_v110 T0) func(_dafny.Int) _dafny.Int {
					return func(_1026_i6 _dafny.Int) _dafny.Int {
						return (_1026_i6).Minus(_1025_v110.F12())
					}
				})(_1023_v110)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw185 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw185).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw185).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1024_v111 = _nw185
			var _1027_v112 _dafny.Map
			_ = _1027_v112
			_1027_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1023_v110.F12(), _1024_v111)
			var _1028_v113 _dafny.Sequence
			_ = _1028_v113
			_1028_v113 = _dafny.SeqOf(_1024_v111)
			var _rhs162 _dafny.Int = (((_1019_v106).Intersection(_1019_v106)).Difference((_1019_v106).Update((_this).F10(), Companion_Default___.Abs((_this).F10())))).Cardinality()
			_ = _rhs162
			var _rhs163 _dafny.Int = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1024_v111)).Merge((_1027_v112).Update(_this.F11, _1024_v111))).Update(_this.F11, (_1028_v113).Select((Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((_1028_v113).Cardinality()))).Uint32()).(_dafny.Array))).Cardinality()
			_ = _rhs163
			var _rhs164 T0 = _1023_v110
			_ = _rhs164
			var _rhs165 _dafny.Int = _this.F11
			_ = _rhs165
			var _lhs103 *C4 = _this
			_ = _lhs103
			var _lhs104 *GlobalState = globalState
			_ = _lhs104
			var _lhs105 T0 = _1023_v110
			_ = _lhs105
			_lhs103.F11 = _rhs162
			_lhs104.F5 = _rhs163
			_1023_v110 = _rhs164
			_lhs105.F12_set_(_rhs165)
			var _1029_v114 _dafny.Array
			_ = _1029_v114
			var _nwElement0_39 _dafny.Array = _1024_v111
			_ = _nwElement0_39
			var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(25))
			_ = _nw186
			(_nw186).ArraySet1(_nwElement0_39, 0)
			(_nw186).ArraySet1(_1024_v111, 1)
			(_nw186).ArraySet1(_1024_v111, 2)
			(_nw186).ArraySet1(_1024_v111, 3)
			(_nw186).ArraySet1(_1024_v111, 4)
			(_nw186).ArraySet1(_1024_v111, 5)
			(_nw186).ArraySet1(_1024_v111, 6)
			(_nw186).ArraySet1(_1024_v111, 7)
			(_nw186).ArraySet1(_1024_v111, 8)
			(_nw186).ArraySet1(_1024_v111, 9)
			(_nw186).ArraySet1(_1024_v111, 10)
			(_nw186).ArraySet1(_1024_v111, 11)
			(_nw186).ArraySet1(_1024_v111, 12)
			(_nw186).ArraySet1(_1024_v111, 13)
			(_nw186).ArraySet1(_1024_v111, 14)
			(_nw186).ArraySet1(_1024_v111, 15)
			(_nw186).ArraySet1(_1024_v111, 16)
			(_nw186).ArraySet1(_1024_v111, 17)
			(_nw186).ArraySet1(_1024_v111, 18)
			(_nw186).ArraySet1(_1024_v111, 19)
			(_nw186).ArraySet1(_1024_v111, 20)
			(_nw186).ArraySet1(_1024_v111, 21)
			(_nw186).ArraySet1(_1024_v111, 22)
			(_nw186).ArraySet1(_1024_v111, 23)
			(_nw186).ArraySet1(_1024_v111, 24)
			_1029_v114 = _nw186
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_1029_v114), 0))
			_ = _index184
			(_1029_v114).ArraySet1(_1024_v111, (_index184).Int())
			var _1030_v115 _dafny.Sequence
			_ = _1030_v115
			_1030_v115 = _dafny.SeqOf(p3, !(true))
			var _1031_v116 _dafny.Array
			_ = _1031_v116
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_28
			var _nw187 _dafny.Array
			_ = _nw187
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw187 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) bool = (func(_1032_p3 bool) func(_dafny.Int) bool {
					return func(_1033_i7 _dafny.Int) bool {
						return _1032_p3
					}
				})(p3)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw187 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw187).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw187).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1031_v116 = _nw187
			var _1034_v117 _dafny.Map
			_ = _1034_v117
			_1034_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1023_v110.F12(), _1031_v116)
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(501), _dafny.ArrayLen((_1029_v114), 0))
			_ = _index185
			var _nwElement0_40 _dafny.Int = _dafny.IntOfUint32((_1030_v115).Cardinality())
			_ = _nwElement0_40
			var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(5))
			_ = _nw188
			(_nw188).ArraySet1(_nwElement0_40, 0)
			(_nw188).ArraySet1(_this.F11, 1)
			(_nw188).ArraySet1(((_1034_v117).Update(_this.F11, _1031_v116)).Cardinality(), 2)
			(_nw188).ArraySet1(_1023_v110.F12(), 3)
			(_nw188).ArraySet1((_this).F10(), 4)
			(_1029_v114).ArraySet1(_nw188, (_index185).Int())
			(_1023_v110).F12_set_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Concatenate(p0, p0))).Cardinality()))
			var _1035_v118 _dafny.Array
			_ = _1035_v118
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_29
			var _nw189 _dafny.Array
			_ = _nw189
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw189 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.CodePoint = (func(_1036_v105 *C1) func(_dafny.Int) _dafny.CodePoint {
					return func(_1037_i8 _dafny.Int) _dafny.CodePoint {
						return (_1036_v105).F15()
					}
				})(_1018_v105)
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw189 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw189).ArraySet1CodePoint(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw189).ArraySet1CodePoint(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_1035_v118 = _nw189
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_1035_v118), 0))
			_ = _index186
			(_1035_v118).ArraySet1CodePoint((_1018_v105).F15(), (_index186).Int())
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_1035_v118), 0))
			_ = _index187
			(_1035_v118).ArraySet1CodePoint(p1, (_index187).Int())
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_1031_v116), 0))
			_ = _index188
			(_1031_v116).ArraySet1(!(p2), (_index188).Int())
			var _1038_v119 _dafny.Map
			_ = _1038_v119
			_1038_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_1031_v116), 0))
			_ = _index189
			(_1031_v116).ArraySet1((func() bool {
				if p3 {
					return (func() bool {
						if (_1038_v119).Contains(Companion_Default___.Fm1((_1018_v105).Fm4(globalState), p3, _this.F11, !(false), globalState)) {
							return (_1038_v119).Get(Companion_Default___.Fm1((_1018_v105).Fm4(globalState), p3, _this.F11, !(false), globalState)).(bool)
						}
						return p2
					})()
				}
				return p3
			})(), (_index189).Int())
		} else {
			var _1039_v120 _dafny.Array
			_ = _1039_v120
			var _len0_30 _dafny.Int = _dafny.One
			_ = _len0_30
			var _nw190 _dafny.Array
			_ = _nw190
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw190 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) _dafny.Int = func(_1040_i9 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1040_i9, _this.F11)
				}
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw190 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw190).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw190).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1039_v120 = _nw190
			var _1041_v121 D1
			_ = _1041_v121
			_1041_v121 = Companion_D1_.Create_DC4_(_1039_v120)
			var _1042_v122 _dafny.Set
			_ = _1042_v122
			_1042_v122 = _dafny.SetOf(Companion_D1_.Create_DC4_(_1039_v120), _1041_v121, _1041_v121)
			var _1043_v123 _dafny.Map
			_ = _1043_v123
			_1043_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v105).F15(), _1042_v122)
			var _1044_v124 _dafny.Map
			_ = _1044_v124
			_1044_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1039_v120)), _this.F11)
			var _1045_v125 _dafny.Sequence
			_ = _1045_v125
			_1045_v125 = _dafny.SeqOf((_this).F10())
			var _1046_v126 _dafny.Map
			_ = _1046_v126
			_1046_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1045_v125)
			var _1047_v127 _dafny.Sequence
			_ = _1047_v127
			_1047_v127 = _dafny.SeqOf((_dafny.SetOf(_1041_v121)).Equals((func() _dafny.Set {
				if (_1043_v123).Contains((_1018_v105).F15()) {
					return (_1043_v123).Get((_1018_v105).F15()).(_dafny.Set)
				}
				return _1042_v122
			})()), ((_dafny.Zero).Minus((func() _dafny.Int {
				if (_1044_v124).Contains(Companion_D7_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1039_v120))) {
					return (_1044_v124).Get(Companion_D7_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1039_v120))).(_dafny.Int)
				}
				return (_1046_v126).Cardinality()
			})())).Cmp(_this.F11) == 0, p3, ((_this).F10()).Cmp(_this.F11) > 0)
			_1047_v127 = Companion_Default___.Fm9(p2, globalState)
			var _1048_v128 bool
			_ = _1048_v128
			_1048_v128 = false
			_1048_v128 = p2
			var _1049_v129 D1
			_ = _1049_v129
			_1049_v129 = Companion_D1_.Create_DC5_(p2)
			var _1050_v130 D1
			_ = _1050_v130
			_1050_v130 = Companion_D1_.Create_DC6_(_1049_v129)
			var _source15 D1 = _1050_v130
			_ = _source15
			if _source15.Is_DC5() {
				var _1051___mcc_h15 bool = _source15.Get_().(D1_DC5).Cf8
				_ = _1051___mcc_h15
				var _1052_cf8 bool = _1051___mcc_h15
				_ = _1052_cf8
				var _1053_v131 _dafny.Map
				_ = _1053_v131
				_1053_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1048_v128), p0)
				var _1054_v132 _dafny.Sequence
				_ = _1054_v132
				_1054_v132 = _dafny.SeqOf((func() _dafny.Sequence {
					if (_1053_v131).Contains(p3) {
						return (_1053_v131).Get(p3).(_dafny.Sequence)
					}
					return p0
				})())
				var _1055_v133 _dafny.Map
				_ = _1055_v133
				_1055_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _1054_v132)
				var _1056_v134 _dafny.Map
				_ = _1056_v134
				_1056_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, (func() _dafny.Sequence {
					if (_1055_v133).Contains(_this.F11) {
						return (_1055_v133).Get(_this.F11).(_dafny.Sequence)
					}
					return _1054_v132
				})())
				_1054_v132 = (func() _dafny.Sequence {
					if (_1056_v134).Contains(((_this).F10()).Plus(_dafny.IntOfUint32((_1047_v127).Cardinality()))) {
						return (_1056_v134).Get(((_this).F10()).Plus(_dafny.IntOfUint32((_1047_v127).Cardinality()))).(_dafny.Sequence)
					}
					return _1054_v132
				})()
				_1052_cf8 = p2
				var _1057_v136 *C0
				_ = _1057_v136
				var _nw191 *C0 = New_C0_()
				_ = _nw191
				_nw191.Ctor__(func() _dafny.Map {
					var _coll24 = _dafny.NewMapBuilder()
					_ = _coll24
					for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-138), _dafny.IntOfInt64(227))); ; {
						_compr_24, _ok24 := _iter24()
						if !_ok24 {
							break
						}
						var _1058_v135 _dafny.Int
						_1058_v135 = interface{}(_compr_24).(_dafny.Int)
						if ((_dafny.IntOfInt64(-138)).Cmp(_1058_v135) <= 0) && ((_1058_v135).Cmp(_dafny.IntOfInt64(227)) < 0) {
							_coll24.Add((_1058_v135).Plus(_this.F11), _1052_cf8)
						}
					}
					return _coll24.ToMap()
				}(), _this.F11)
				_1057_v136 = _nw191
				var _1059_v137 D1
				_ = _1059_v137
				_1059_v137 = Companion_D1_.Create_DC5_(_1048_v128)
				var _1060_v138 *C2
				_ = _1060_v138
				var _nw192 *C2 = New_C2_()
				_ = _nw192
				_nw192.Ctor__((_1059_v137).Dtor_cf8(), (_this).F10())
				_1060_v138 = _nw192
			} else if _source15.Is_DC4() {
				var _1061___mcc_h16 _dafny.Array = _source15.Get_().(D1_DC4).Cf7
				_ = _1061___mcc_h16
				var _1062_cf7 _dafny.Array = _1061___mcc_h16
				_ = _1062_cf7
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1062_cf7), 0))
				_ = _index190
				(_1062_cf7).ArraySet1(_this.F11, (_index190).Int())
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1062_cf7), 0))
				_ = _index191
				(_1062_cf7).ArraySet1(_this.F11, (_index191).Int())
				var _1063_v140 _dafny.Map
				_ = _1063_v140
				_1063_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1062_cf7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_1062_cf7), 0))).Int()).(_dafny.Int), (_1018_v105).F15())
				var _1064_v141 _dafny.Map
				_ = _1064_v141
				_1064_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), true)
				var _1065_v142 *C0
				_ = _1065_v142
				var _nw193 *C0 = New_C0_()
				_ = _nw193
				_nw193.Ctor__((func() _dafny.Map {
					var _coll25 = _dafny.NewMapBuilder()
					_ = _coll25
					for _iter25 := _dafny.Iterate((_1063_v140).Keys().Elements()); ; {
						_compr_25, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _1066_v139 _dafny.Int
						_1066_v139 = interface{}(_compr_25).(_dafny.Int)
						if (_1063_v140).Contains(_1066_v139) {
							_coll25.Add((_1066_v139).Plus(_this.F11), p2)
						}
					}
					return _coll25.ToMap()
				}()).Merge(_1064_v141), (_this).F10())
				_1065_v142 = _nw193
				var _1067_v143 _dafny.Sequence
				_ = _1067_v143
				_1067_v143 = _dafny.UnicodeSeqOfUtf8Bytes("puri")
				_1067_v143 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("s"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}((func(_1068_v105 *C1) func(_dafny.Int) _dafny.CodePoint {
					return func(_1069_i10 _dafny.Int) _dafny.CodePoint {
						return (_1068_v105).F15()
					}
				})(_1018_v105)))), _1067_v143)
				var _1070_v144 _dafny.Sequence
				_ = _1070_v144
				_1070_v144 = _1067_v143
				_1067_v143 = _dafny.Companion_Sequence_.Concatenate((_1070_v144), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-971))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_1071_v105 *C1) func(_dafny.Int) _dafny.CodePoint {
					return func(_1072_i11 _dafny.Int) _dafny.CodePoint {
						return (_1071_v105).F15()
					}
				})(_1018_v105))))
			} else {
				var _1073___mcc_h17 D1 = _source15.Get_().(D1_DC6).Cf9
				_ = _1073___mcc_h17
				var _1074_cf9 D1 = _1073___mcc_h17
				_ = _1074_cf9
				_1048_v128 = p3
				var _1075_v145 D5
				_ = _1075_v145
				_1075_v145 = Companion_D5_.Create_DC14_(p2)
				var _1076_v146 _dafny.Array
				_ = _1076_v146
				var _nwElement0_41 bool = p3
				_ = _nwElement0_41
				var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(11))
				_ = _nw194
				(_nw194).ArraySet1(_nwElement0_41, 0)
				(_nw194).ArraySet1(_1048_v128, 1)
				(_nw194).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_this).F10()), _1045_v125), 2)
				(_nw194).ArraySet1(_1048_v128, 3)
				(_nw194).ArraySet1(false, 4)
				(_nw194).ArraySet1(_1048_v128, 5)
				(_nw194).ArraySet1(_1048_v128, 6)
				(_nw194).ArraySet1(p2, 7)
				(_nw194).ArraySet1(p3, 8)
				(_nw194).ArraySet1(!(Companion_Default___.Fm1((_this).F10(), !(_1048_v128), (_this).F10(), _1048_v128, globalState)), 9)
				(_nw194).ArraySet1((_1075_v145).Dtor_cf20(), 10)
				_1076_v146 = _nw194
				var _1077_v147 _dafny.Set
				_ = _1077_v147
				_1077_v147 = _dafny.SetOf(p3, !(p3))
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_1076_v146), 0))
				_ = _index192
				(_1076_v146).ArraySet1(Companion_Default___.Fm1((_1077_v147).Cardinality(), p3, (_this).F10(), p2, globalState), (_index192).Int())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_1076_v146), 0))
				_ = _index193
				(_1076_v146).ArraySet1(true, (_index193).Int())
				var _1078_v148 _dafny.Array
				_ = _1078_v148
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_31
				var _nw195 _dafny.Array
				_ = _nw195
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw195 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Sequence = (func(_1079_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1080_i12 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_1079_p0).Cardinality())), _dafny.SeqOf(_this.F11))
						}
					})(p0)
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
				_1078_v148 = _nw195
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(38), _dafny.ArrayLen((_1078_v148), 0))
				_ = _index194
				(_1078_v148).ArraySet1(_1045_v125, (_index194).Int())
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(38), _dafny.ArrayLen((_1078_v148), 0))
				_ = _index195
				(_1078_v148).ArraySet1(_dafny.Companion_Sequence_.Update(_1045_v125, (Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((_1045_v125).Cardinality()))).Uint32(), (_this).F10()), (_index195).Int())
				_1048_v128 = _1048_v128
			}
			_1048_v128 = !(Companion_Default___.Fm1(_this.F11, _1048_v128, (_this).F10(), (p3) == (p3), globalState))
			var _1081_v149 _dafny.Map
			_ = _1081_v149
			_1081_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1048_v128)
			(globalState).F5 = (func() _dafny.Int {
				if Companion_Default___.Fm1(_this.F11, _1048_v128, (_this).F10(), _1048_v128, globalState) {
					return (_1081_v149).Cardinality()
				}
				return (_this).F10()
			})()
		}
		var _1082_i13 _dafny.Int
		_ = _1082_i13
		_1082_i13 = _dafny.Zero
		{
			for !(!(!(p3)) || (true)) {
				{
					if (_1082_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1082_i13 = (_1082_i13).Plus(_dafny.One)
					var _1083_v150 bool
					_ = _1083_v150
					_1083_v150 = false
					var _1084_v151 _dafny.Array
					_ = _1084_v151
					var _len0_32 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_32
					var _nw196 _dafny.Array
					_ = _nw196
					if _len0_32.Cmp(_dafny.Zero) == 0 {
						_nw196 = _dafny.NewArray(_len0_32)
					} else {
						var _init32 func(_dafny.Int) bool = (func(_1085_v150 bool) func(_dafny.Int) bool {
							return func(_1086_i14 _dafny.Int) bool {
								return _1085_v150
							}
						})(_1083_v150)
						_ = _init32
						var _element0_32 = _init32(_dafny.Zero)
						_ = _element0_32
						_nw196 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
						(_nw196).ArraySet1(_element0_32, 0)
						var _nativeLen0_32 = (_len0_32).Int()
						_ = _nativeLen0_32
						for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
							(_nw196).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
						}
					}
					_1084_v151 = _nw196
					var _1087_v152 _dafny.Set
					_ = _1087_v152
					_1087_v152 = _dafny.SetOf(_1084_v151)
					_1083_v150 = (_1087_v152).IsProperSubsetOf((_1087_v152).Difference(_1087_v152))
					var _1088_v154 _dafny.MultiSet
					_ = _1088_v154
					_1088_v154 = _dafny.MultiSetOf(false, false)
					var _1089_v155 _dafny.Map
					_ = _1089_v155
					_1089_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1088_v154).Cardinality(), true)
					var _1090_v156 D9
					_ = _1090_v156
					_1090_v156 = Companion_D9_.Create_DC27_(_1018_v105, _this.F11, _this.F11, (_this).F10(), func() _dafny.Map {
						var _coll26 = _dafny.NewMapBuilder()
						_ = _coll26
						for _iter26 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_1089_v155))).Elements()); ; {
							_compr_26, _ok26 := _iter26()
							if !_ok26 {
								break
							}
							var _1091_v153 _dafny.Map
							_1091_v153 = interface{}(_compr_26).(_dafny.Map)
							if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_1089_v155))).Contains(_1091_v153) {
								_coll26.Add(_1091_v153, (_1018_v105).F15())
							}
						}
						return _coll26.ToMap()
					}())
					_1083_v150 = Companion_Default___.Fm1((_this).F10(), p3, ((_dafny.SetOf((_this).F10(), _this.F11, (_1090_v156).Dtor_cf34(), (_this).F10())).Union(_866_v1)).Cardinality(), !(!((func() bool {
						if Companion_Default___.Fm1((_dafny.Zero).Minus(_this.F11), p3, (_this).F10(), false, globalState) {
							return !(p3)
						}
						return _1083_v150
					})())), globalState)
					var _1092_v157 _dafny.Sequence
					_ = _1092_v157
					_1092_v157 = _dafny.SeqOf(_this.F11)
					_1092_v157 = _1092_v157
					if _1083_v150 {
						var _1093_v158 _dafny.Array
						_ = _1093_v158
						var _nwElement0_42 _dafny.Sequence = _1092_v157
						_ = _nwElement0_42
						var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(5))
						_ = _nw197
						(_nw197).ArraySet1(_nwElement0_42, 0)
						(_nw197).ArraySet1(_1092_v157, 1)
						(_nw197).ArraySet1(_1092_v157, 2)
						(_nw197).ArraySet1(_1092_v157, 3)
						(_nw197).ArraySet1(_1092_v157, 4)
						_1093_v158 = _nw197
						var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_1093_v158), 0))
						_ = _index196
						(_1093_v158).ArraySet1(_dafny.SeqOf((_1092_v157).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_1092_v157).Cardinality()))).Uint32()).(_dafny.Int), (_this).F10()), (_index196).Int())
						var _1094_v159 _dafny.Array
						_ = _1094_v159
						var _len0_33 _dafny.Int = _dafny.IntOfInt64(12)
						_ = _len0_33
						var _nw198 _dafny.Array
						_ = _nw198
						if _len0_33.Cmp(_dafny.Zero) == 0 {
							_nw198 = _dafny.NewArray(_len0_33)
						} else {
							var _init33 func(_dafny.Int) D2 = func(_1095_i15 _dafny.Int) D2 {
								return Companion_D2_.Create_DC7_((_this).F10())
							}
							_ = _init33
							var _element0_33 = _init33(_dafny.Zero)
							_ = _element0_33
							_nw198 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
							(_nw198).ArraySet1(_element0_33, 0)
							var _nativeLen0_33 = (_len0_33).Int()
							_ = _nativeLen0_33
							for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
								(_nw198).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
							}
						}
						_1094_v159 = _nw198
						var _1096_v160 T1
						_ = _1096_v160
						var _nw199 *C3 = New_C3_()
						_ = _nw199
						_nw199.Ctor__(_1094_v159, (p0).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint), (_this).F10())
						_1096_v160 = _nw199
						var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_1093_v158), 0))
						_ = _index197
						var _rhs166 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-973), _1096_v160.F12(), _dafny.IntOfInt64(-185)), _1092_v157)
						_ = _rhs166
						var _rhs167 _dafny.Int = _dafny.IntOfInt64(-939)
						_ = _rhs167
						var _rhs168 _dafny.Int = ((_this).Fm2(false, globalState)).Times(_1096_v160.F12())
						_ = _rhs168
						var _rhs169 T1 = _1096_v160
						_ = _rhs169
						var _rhs170 bool = Companion_Default___.Fm1(_dafny.IntOfInt64(599), _1083_v150, _this.F11, _1083_v150, globalState)
						_ = _rhs170
						var _lhs106 _dafny.Array = _1093_v158
						_ = _lhs106
						var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_1093_v158), 0))
						_ = _lhs107
						var _lhs108 *GlobalState = globalState
						_ = _lhs108
						(_lhs106).ArraySet1(_rhs166, (_lhs107).Int())
						_lhs108.F5 = _rhs167
						r0 = _rhs168
						_1096_v160 = _rhs169
						_1083_v150 = _rhs170
						var _1097_v161 _dafny.Map
						_ = _1097_v161
						_1097_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1083_v150, (_this).F10())
						var _1098_v162 _dafny.Sequence
						_ = _1098_v162
						_1098_v162 = _dafny.SeqOf(true)
						var _rhs171 bool = (_1096_v160.F12()).Cmp(((_1097_v161).Merge(_1097_v161)).Cardinality()) <= 0
						_ = _rhs171
						var _rhs172 bool = Companion_Default___.Fm1((_1089_v155).Cardinality(), _dafny.Companion_Sequence_.Contains(_1098_v162, false), (func() _dafny.Int {
							if _1083_v150 {
								return (_this).F10()
							}
							return (_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality()))
						})(), p3, globalState)
						_ = _rhs172
						var _rhs173 bool = (_1096_v160.F12()).Cmp((_this).F10()) == 0
						_ = _rhs173
						_1083_v150 = _rhs171
						_1083_v150 = _rhs172
						_1083_v150 = _rhs173
						var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1084_v151), 0))
						_ = _index198
						(_1084_v151).ArraySet1(!(p2), (_index198).Int())
						var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1084_v151), 0))
						_ = _index199
						(_1084_v151).ArraySet1(p3, (_index199).Int())
						var _1099_v163 _dafny.Sequence
						_ = _1099_v163
						_1099_v163 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F10()))
						var _1100_v164 T0
						_ = _1100_v164
						var _nw200 *C2 = New_C2_()
						_ = _nw200
						_nw200.Ctor__(p3, (_dafny.Zero).Minus(Companion_Default___.Fm0(((_1099_v163).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1099_v163).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), (_1097_v161).Cardinality(), globalState)))
						_1100_v164 = _nw200
						_1100_v164 = _1100_v164
						var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_1084_v151), 0))
						_ = _index200
						(_1084_v151).ArraySet1(p3, (_index200).Int())
					} else {
						_1083_v150 = p3
						var _1101_v165 _dafny.Array
						_ = _1101_v165
						var _len0_34 _dafny.Int = _dafny.IntOfInt64(5)
						_ = _len0_34
						var _nw201 _dafny.Array
						_ = _nw201
						if _len0_34.Cmp(_dafny.Zero) == 0 {
							_nw201 = _dafny.NewArray(_len0_34)
						} else {
							var _init34 func(_dafny.Int) D5 = (func(_1102_v150 bool, _1103_p2 bool) func(_dafny.Int) D5 {
								return func(_1104_i16 _dafny.Int) D5 {
									return Companion_D5_.Create_DC13_(!(_1102_v150), !(_1103_p2))
								}
							})(_1083_v150, p2)
							_ = _init34
							var _element0_34 = _init34(_dafny.Zero)
							_ = _element0_34
							_nw201 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
							(_nw201).ArraySet1(_element0_34, 0)
							var _nativeLen0_34 = (_len0_34).Int()
							_ = _nativeLen0_34
							for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
								(_nw201).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
							}
						}
						_1101_v165 = _nw201
						var _1105_v166 D5
						_ = _1105_v166
						_1105_v166 = Companion_D5_.Create_DC13_(!(p3), p3)
						var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_1101_v165), 0))
						_ = _index201
						(_1101_v165).ArraySet1(_1105_v166, (_index201).Int())
						var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_1101_v165), 0))
						_ = _index202
						(_1101_v165).ArraySet1(_1105_v166, (_index202).Int())
						var _1106_v167 *C1
						_ = _1106_v167
						var _nw202 *C1 = New_C1_()
						_ = _nw202
						_nw202.Ctor__((_1018_v105).F15(), (_1018_v105).F15(), _this.F11)
						_1106_v167 = _nw202
						var _1107_v168 *C1
						_ = _1107_v168
						var _nw203 *C1 = New_C1_()
						_ = _nw203
						_nw203.Ctor__((_1106_v167).F15(), (_1106_v167).F15(), (_this).F10())
						_1107_v168 = _nw203
						var _1108_v169 _dafny.Array
						_ = _1108_v169
						var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
						_ = _nw204
						_1108_v169 = _nw204
						var _1109_v170 _dafny.Set
						_ = _1109_v170
						_1109_v170 = _dafny.SetOf(p2, true, _1083_v150)
						var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1108_v169), 0))
						_ = _index203
						(_1108_v169).ArraySet1((_1109_v170).Cardinality(), (_index203).Int())
						var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1108_v169), 0))
						_ = _index204
						(_1108_v169).ArraySet1((_this).F10(), (_index204).Int())
					}
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		if (Companion_D2_.Create_DC8_(p3, _this.F11, p2)).Dtor_cf11() {
			var _1110_v171 _dafny.Map
			_ = _1110_v171
			_1110_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p2)
			var _1111_v172 *C0
			_ = _1111_v172
			var _nw205 *C0 = New_C0_()
			_ = _nw205
			_nw205.Ctor__(_1110_v171, (_this).F10())
			_1111_v172 = _nw205
			var _1112_v173 _dafny.Map
			_ = _1112_v173
			_1112_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v105).F15(), _1111_v172)
			_1112_v173 = _1112_v173
			if (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eubmv")).Cardinality())).Cmp((_this).F10()) > 0 {
				var _1113_v174 _dafny.Int
				_ = _1113_v174
				var _1114_v175 bool
				_ = _1114_v175
				var _1115_v176 _dafny.Map
				_ = _1115_v176
				var _out54 _dafny.Int
				_ = _out54
				var _out55 bool
				_ = _out55
				var _out56 _dafny.Map
				_ = _out56
				_out54, _out55, _out56 = (_1018_v105).M5(globalState)
				_1113_v174 = _out54
				_1114_v175 = _out55
				_1115_v176 = _out56
				var _1116_v177 _dafny.Set
				_ = _1116_v177
				_1116_v177 = _dafny.SetOf(_1114_v175)
				var _1117_v178 D0
				_ = _1117_v178
				_1117_v178 = Companion_D0_.Create_DC0_(true)
				var _1118_v179 _dafny.Array
				_ = _1118_v179
				var _nwElement0_43 _dafny.Set = _1116_v177
				_ = _nwElement0_43
				var _nw206 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(28))
				_ = _nw206
				(_nw206).ArraySet1(_nwElement0_43, 0)
				(_nw206).ArraySet1(_1116_v177, 1)
				(_nw206).ArraySet1(_dafny.SetOf(p2), 2)
				(_nw206).ArraySet1((_1116_v177).Union(_1116_v177), 3)
				(_nw206).ArraySet1(_dafny.SetOf(true), 4)
				(_nw206).ArraySet1(_dafny.SetOf(_1114_v175), 5)
				(_nw206).ArraySet1(_1116_v177, 6)
				(_nw206).ArraySet1((_dafny.SetOf(p3, p3)).Difference(_1116_v177), 7)
				(_nw206).ArraySet1(_1116_v177, 8)
				(_nw206).ArraySet1(_dafny.SetOf((_1117_v178).Dtor_cf0(), p3, !(_1114_v175), Companion_Default___.Fm1(_dafny.IntOfInt64(818), p3, _1113_v174, p2, globalState)), 9)
				(_nw206).ArraySet1(_1116_v177, 10)
				(_nw206).ArraySet1(_1116_v177, 11)
				(_nw206).ArraySet1(_dafny.SetOf(p2, p2), 12)
				(_nw206).ArraySet1(_dafny.SetOf(p3), 13)
				(_nw206).ArraySet1(_1116_v177, 14)
				(_nw206).ArraySet1(_1116_v177, 15)
				(_nw206).ArraySet1(_1116_v177, 16)
				(_nw206).ArraySet1(_1116_v177, 17)
				(_nw206).ArraySet1(_dafny.SetOf(_1114_v175, _1114_v175, p3), 18)
				(_nw206).ArraySet1((_1116_v177).Difference(_1116_v177), 19)
				(_nw206).ArraySet1(_1116_v177, 20)
				(_nw206).ArraySet1((_1116_v177).Union(_dafny.SetOf(p3)), 21)
				(_nw206).ArraySet1(_1116_v177, 22)
				(_nw206).ArraySet1(_1116_v177, 23)
				(_nw206).ArraySet1((_1116_v177).Difference(_1116_v177), 24)
				(_nw206).ArraySet1(_1116_v177, 25)
				(_nw206).ArraySet1((_1116_v177).Intersection(_dafny.SetOf(p3)), 26)
				(_nw206).ArraySet1(_1116_v177, 27)
				_1118_v179 = _nw206
				var _1119_v180 _dafny.Set
				_ = _1119_v180
				_1119_v180 = _1116_v177
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_1118_v179), 0))
				_ = _index205
				(_1118_v179).ArraySet1((_1119_v180).Intersection((_1119_v180)), (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_1118_v179), 0))
				_ = _index206
				var _rhs174 bool = false
				_ = _rhs174
				var _rhs175 _dafny.Set = (_1116_v177).Difference((_1116_v177).Difference(_1116_v177))
				_ = _rhs175
				var _lhs109 _dafny.Array = _1118_v179
				_ = _lhs109
				var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_1118_v179), 0))
				_ = _lhs110
				_1114_v175 = _rhs174
				(_lhs109).ArraySet1(_rhs175, (_lhs110).Int())
				var _1120_v181 _dafny.Array
				_ = _1120_v181
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_35
				var _nw207 _dafny.Array
				_ = _nw207
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw207 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) _dafny.Sequence = (func(_1121_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1122_i17 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hhjkitey"), _1121_p0)
						}
					})(p0)
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw207 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw207).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw207).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1120_v181 = _nw207
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_1120_v181), 0))
				_ = _index207
				(_1120_v181).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), (_index207).Int())
				var _1123_v182 _dafny.Sequence
				_ = _1123_v182
				_1123_v182 = _dafny.SeqOf(Companion_Default___.Fm1(_1113_v174, false, (_this).F10(), p3, globalState))
				var _1124_v183 D8
				_ = _1124_v183
				_1124_v183 = Companion_D8_.Create_DC24_(_1116_v177, _1123_v182)
				var _pat_let_tv8 = _1118_v179
				_ = _pat_let_tv8
				var _pat_let_tv9 = _1118_v179
				_ = _pat_let_tv9
				var _1125_v184 _dafny.MultiSet
				_ = _1125_v184
				_1125_v184 = _dafny.MultiSetOf(func(_pat_let17_0 D8) D8 {
					return func(_1126_dt__update__tmp_h2 D8) D8 {
						return func(_pat_let18_0 _dafny.Set) D8 {
							return func(_1127_dt__update_hcf30_h0 _dafny.Set) D8 {
								return Companion_D8_.Create_DC24_(_1127_dt__update_hcf30_h0, (_1126_dt__update__tmp_h2).Dtor_cf31())
							}(_pat_let18_0)
						}((_pat_let_tv9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_pat_let_tv8), 0))).Int()).(_dafny.Set))
					}(_pat_let17_0)
				}(_1124_v183))
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_1120_v181), 0))
				_ = _index208
				(_1120_v181).ArraySet1((func() _dafny.Sequence {
					if (_1125_v184).IsSubsetOf(_1125_v184) {
						return p0
					}
					return p0
				})(), (_index208).Int())
				var _1128_v185 _dafny.Array
				_ = _1128_v185
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_36
				var _nw208 _dafny.Array
				_ = _nw208
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw208 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = (func(_1129_v106 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
						return func(_1130_i18 _dafny.Int) _dafny.Int {
							return (_1130_i18).Plus((_1129_v106).Cardinality())
						}
					})(_1019_v106)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw208 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw208).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw208).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1128_v185 = _nw208
				var _1131_v186 _dafny.Array
				_ = _1131_v186
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_37
				var _nw209 _dafny.Array
				_ = _nw209
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw209 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) D2 = (func(_1132_v1 _dafny.Set) func(_dafny.Int) D2 {
						return func(_1133_i19 _dafny.Int) D2 {
							return Companion_D2_.Create_DC7_((_1132_v1).Cardinality())
						}
					})(_866_v1)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw209 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw209).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw209).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1131_v186 = _nw209
				var _1134_v187 T1
				_ = _1134_v187
				var _nw210 *C3 = New_C3_()
				_ = _nw210
				_nw210.Ctor__(_1131_v186, (_1018_v105).F15(), (_this).F10())
				_1134_v187 = _nw210
				var _1135_v188 _dafny.Sequence
				_ = _1135_v188
				_1135_v188 = _dafny.SeqOf(_1134_v187)
				var _1136_v189 _dafny.Map
				_ = _1136_v189
				_1136_v189 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1128_v185, _1135_v188)
				_1136_v189 = (_1136_v189).Update((func() _dafny.Array {
					if false {
						return _1128_v185
					}
					return _1128_v185
				})(), _dafny.Companion_Sequence_.Concatenate(_1135_v188, _1135_v188))
				(globalState).F5 = Companion_Default___.SafeDivisionInt((_this).F10(), (_this).F10())
			} else {
				var _1137_v190 bool
				_ = _1137_v190
				_1137_v190 = false
				_1137_v190 = p3
				var _1138_v191 _dafny.Map
				_ = _1138_v191
				_1138_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v105).F15(), _dafny.MultiSetFromSeq((_1018_v105).Fm5(p2, !(p2), false, !(p3), globalState)))
				var _1139_v192 _dafny.Sequence
				_ = _1139_v192
				_1139_v192 = _dafny.SeqOf((_this).F10())
				var _1140_v193 D0
				_ = _1140_v193
				_1140_v193 = Companion_D0_.Create_DC0_(p2)
				var _1141_v194 _dafny.Sequence
				_ = _1141_v194
				_1141_v194 = _dafny.SeqOf(_1137_v190)
				var _1142_v196 _dafny.Array
				_ = _1142_v196
				var _nwElement0_44 bool = false
				_ = _nwElement0_44
				var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(20))
				_ = _nw211
				(_nw211).ArraySet1(_nwElement0_44, 0)
				(_nw211).ArraySet1((Companion_Default___.Fm21((_this).F10(), (_this).F10(), _dafny.CodePoint('h'), (_this).F10(), globalState)).IsSubsetOf(((func() _dafny.MultiSet {
					if (_1138_v191).Contains((_1018_v105).F15()) {
						return (_1138_v191).Get((_1018_v105).F15()).(_dafny.MultiSet)
					}
					return _1019_v106
				})()).Update((_this).F10(), Companion_Default___.Abs(_dafny.IntOfInt64(-600)))), 1)
				(_nw211).ArraySet1(p2, 2)
				(_nw211).ArraySet1(p3, 3)
				(_nw211).ArraySet1(p2, 4)
				(_nw211).ArraySet1(p2, 5)
				(_nw211).ArraySet1(!(p2), 6)
				(_nw211).ArraySet1((_dafny.IntOfUint32((_1139_v192).Cardinality())).Cmp((func() _dafny.Int {
					if (_1019_v106).Contains(_this.F11) {
						return (_1019_v106).Multiplicity(_this.F11)
					}
					return _this.F11
				})()) <= 0, 7)
				(_nw211).ArraySet1((_1140_v193).Dtor_cf0(), 8)
				(_nw211).ArraySet1(p3, 9)
				(_nw211).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((_1139_v192).Cardinality()), Companion_Default___.Fm1((_dafny.SetOf((_this).F10())).Cardinality(), p3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wma")).Cardinality()), (_1141_v194).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1141_v194).Cardinality()))).Uint32()).(bool), globalState), (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F10())), _1137_v190, globalState), 10)
				(_nw211).ArraySet1(true, 11)
				(_nw211).ArraySet1(true, 12)
				(_nw211).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(61))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}(func(_1143_i20 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), p0), 13)
				(_nw211).ArraySet1(p2, 14)
				(_nw211).ArraySet1(!((func() _dafny.Set {
					var _coll27 = _dafny.NewBuilder()
					_ = _coll27
					for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(612), _dafny.IntOfInt64(-261))); ; {
						_compr_27, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _1144_v195 _dafny.Int
						_1144_v195 = interface{}(_compr_27).(_dafny.Int)
						if ((_dafny.IntOfInt64(612)).Cmp(_1144_v195) <= 0) && ((_1144_v195).Cmp(_dafny.IntOfInt64(-261)) < 0) {
							_coll27.Add(Companion_Default___.SafeDivisionInt(_1144_v195, (_this).F10()))
						}
					}
					return _coll27.ToSet()
				}()).IsSubsetOf(_866_v1)), 15)
				(_nw211).ArraySet1(_1137_v190, 16)
				(_nw211).ArraySet1(_1137_v190, 17)
				(_nw211).ArraySet1(((_this).F10()).Cmp(_this.F11) >= 0, 18)
				(_nw211).ArraySet1(false, 19)
				_1142_v196 = _nw211
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_1142_v196), 0))
				_ = _index209
				(_1142_v196).ArraySet1(p2, (_index209).Int())
				var _1145_v197 _dafny.Set
				_ = _1145_v197
				_1145_v197 = _dafny.SetOf(false)
				var _1146_v198 _dafny.Set
				_ = _1146_v198
				_1146_v198 = _1145_v197
				var _1147_v199 _dafny.Sequence
				_ = _1147_v199
				_1147_v199 = _dafny.SeqOf(_1111_v172)
				var _1148_v200 _dafny.Array
				_ = _1148_v200
				var _nw212 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw212
				_1148_v200 = _nw212
				var _1149_v201 _dafny.Sequence
				_ = _1149_v201
				_1149_v201 = _dafny.SeqOf(_1148_v200, _1148_v200, _1148_v200)
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_1142_v196), 0))
				_ = _index210
				var _rhs176 _dafny.Int = Companion_Default___.SafeModuloInt((_866_v1).Cardinality(), _dafny.IntOfUint32((_1147_v199).Cardinality()))
				_ = _rhs176
				var _rhs177 _dafny.Int = ((_this).F10()).Times((_this).F10())
				_ = _rhs177
				var _rhs178 bool = ((_1019_v106).Intersection(_1019_v106)).IsProperSubsetOf(_1019_v106)
				_ = _rhs178
				var _rhs179 _dafny.Set = _1146_v198
				_ = _rhs179
				var _rhs180 _dafny.Array = (_1149_v201).Select((Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((_1149_v201).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs180
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				var _lhs113 _dafny.Array = _1142_v196
				_ = _lhs113
				var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_1142_v196), 0))
				_ = _lhs114
				_lhs111.F5 = _rhs176
				_lhs112.F5 = _rhs177
				(_lhs113).ArraySet1(_rhs178, (_lhs114).Int())
				_1146_v198 = _rhs179
				r3 = _rhs180
				var _1150_v202 D1
				_ = _1150_v202
				_1150_v202 = Companion_D1_.Create_DC5_(p3)
				_1137_v190 = (_1150_v202).Dtor_cf8()
				var _nw213 *C0 = New_C0_()
				_ = _nw213
				_nw213.Ctor__(func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(361), _dafny.IntOfInt64(634))); ; {
						_compr_28, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _1151_v203 _dafny.Int
						_1151_v203 = interface{}(_compr_28).(_dafny.Int)
						if ((_dafny.IntOfInt64(361)).Cmp(_1151_v203) <= 0) && ((_1151_v203).Cmp(_dafny.IntOfInt64(634)) < 0) {
							_coll28.Add(Companion_Default___.SafeDivisionInt(_1151_v203, (_this).F10()), true)
						}
					}
					return _coll28.ToMap()
				}(), _this.F11)
				_1111_v172 = _nw213
				var _1152_v204 _dafny.Map
				_ = _1152_v204
				_1152_v204 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1139_v192, _1139_v192), _1141_v194)
				_1141_v194 = (func() _dafny.Sequence {
					if (_1152_v204).Contains(_1139_v192) {
						return (_1152_v204).Get(_1139_v192).(_dafny.Sequence)
					}
					return _dafny.SeqOf(_1137_v190)
				})()
			}
			var _1153_v205 _dafny.Map
			_ = _1153_v205
			_1153_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), (p2) && (p3))
			_1153_v205 = (_1153_v205).Update(p3, p3)
			var _1154_v206 _dafny.Sequence
			_ = _1154_v206
			_1154_v206 = _dafny.SeqOf(_this.F11)
			(globalState).F5 = (_1154_v206).Select((Companion_Default___.SafeIndex((_1019_v106).Cardinality(), _dafny.IntOfUint32((_1154_v206).Cardinality()))).Uint32()).(_dafny.Int)
			var _1155_v207 _dafny.Array
			_ = _1155_v207
			var _nw214 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw214
			_1155_v207 = _nw214
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1155_v207), 0))
			_ = _index211
			(_1155_v207).ArraySet1((func() bool {
				if p3 {
					return p2
				}
				return p3
			})(), (_index211).Int())
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1155_v207), 0))
			_ = _index212
			(_1155_v207).ArraySet1(p3, (_index212).Int())
		} else {
			var _1156_v208 _dafny.Map
			_ = _1156_v208
			_1156_v208 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
			_1156_v208 = (_1156_v208).Update(p3, Companion_Default___.Fm1((_this).F10(), p2, _this.F11, !(p2), globalState))
			var _1157_v209 _dafny.Array
			_ = _1157_v209
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_38
			var _nw215 _dafny.Array
			_ = _nw215
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw215 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) _dafny.Set = (func(_1158_p3 bool, _1159_p2 bool) func(_dafny.Int) _dafny.Set {
					return func(_1160_i21 _dafny.Int) _dafny.Set {
						return _dafny.SetOf(_1158_p3, _1159_p2, _1159_p2, !(true), _1159_p2)
					}
				})(p3, p2)
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw215 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw215).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw215).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1157_v209 = _nw215
			_1157_v209 = _1157_v209
			var _1161_v210 D2
			_ = _1161_v210
			_1161_v210 = Companion_D2_.Create_DC7_(_this.F11)
			var _1162_v211 _dafny.Array
			_ = _1162_v211
			var _nwElement0_45 D2 = _1161_v210
			_ = _nwElement0_45
			var _nw216 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(3))
			_ = _nw216
			(_nw216).ArraySet1(_nwElement0_45, 0)
			(_nw216).ArraySet1(Companion_D2_.Create_DC7_(_this.F11), 1)
			(_nw216).ArraySet1(_1161_v210, 2)
			_1162_v211 = _nw216
			var _1163_v212 *C3
			_ = _1163_v212
			var _nw217 *C3 = New_C3_()
			_ = _nw217
			_nw217.Ctor__(_1162_v211, (_1018_v105).F15(), _this.F11)
			_1163_v212 = _nw217
			var _1164_v213 _dafny.Map
			_ = _1164_v213
			_1164_v213 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v212, _this.F11)
			var _1165_v214 _dafny.MultiSet
			_ = _1165_v214
			_1165_v214 = _dafny.MultiSetOf(_1164_v213, _1164_v213)
			var _1166_v215 _dafny.MultiSet
			_ = _1166_v215
			_1166_v215 = _dafny.MultiSetOf((_1165_v214).IsDisjointFrom(_1165_v214))
			_1166_v215 = _1166_v215
			var _1167_v216 _dafny.Array
			_ = _1167_v216
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_39
			var _nw218 _dafny.Array
			_ = _nw218
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw218 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Int = func(_1168_i22 _dafny.Int) _dafny.Int {
					return (_1168_i22).Minus((_this).F10())
				}
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw218 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw218).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw218).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1167_v216 = _nw218
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_1167_v216), 0))
			_ = _index213
			(_1167_v216).ArraySet1((_dafny.Zero).Minus((_this).F10()), (_index213).Int())
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_1167_v216), 0))
			_ = _index214
			(_1167_v216).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), (_index214).Int())
			var _1169_v217 _dafny.Array
			_ = _1169_v217
			var _nw219 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw219
			_1169_v217 = _nw219
			var _1170_v218 D0
			_ = _1170_v218
			_1170_v218 = Companion_D0_.Create_DC1_(_1169_v217)
			_1170_v218 = _1170_v218
		}
		_866_v1 = _866_v1
		r0 = _dafny.IntOfInt64(-907)
		var _1171_v219 _dafny.MultiSet
		_ = _1171_v219
		_1171_v219 = _dafny.MultiSetOf(p2)
		var _1172_v220 _dafny.Array
		_ = _1172_v220
		var _nwElement0_46 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(p0, p0)
		_ = _nwElement0_46
		var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(5))
		_ = _nw220
		(_nw220).ArraySet1(_nwElement0_46, 0)
		(_nw220).ArraySet1((_this.F11).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(302))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg47 _dafny.Int) interface{} {
				return coer47(arg47)
			}
		}((func(_1173_v105 *C1) func(_dafny.Int) _dafny.CodePoint {
			return func(_1174_i23 _dafny.Int) _dafny.CodePoint {
				return (_1173_v105).F15()
			}
		})(_1018_v105)))).Cardinality())) == 0, 1)
		(_nw220).ArraySet1(!(p3), 2)
		(_nw220).ArraySet1(!(Companion_Default___.Fm1((_this).F10(), p2, (_1171_v219).Cardinality(), p2, globalState)), 3)
		(_nw220).ArraySet1(p3, 4)
		_1172_v220 = _nw220
		r1 = _1172_v220
		var _1175_v221 _dafny.Array
		_ = _1175_v221
		var _len0_40 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_40
		var _nw221 _dafny.Array
		_ = _nw221
		if _len0_40.Cmp(_dafny.Zero) == 0 {
			_nw221 = _dafny.NewArray(_len0_40)
		} else {
			var _init40 func(_dafny.Int) D0 = (func(_1176_v2 D0) func(_dafny.Int) D0 {
				return func(_1177_i24 _dafny.Int) D0 {
					return _1176_v2
				}
			})(_867_v2)
			_ = _init40
			var _element0_40 = _init40(_dafny.Zero)
			_ = _element0_40
			_nw221 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
			(_nw221).ArraySet1(_element0_40, 0)
			var _nativeLen0_40 = (_len0_40).Int()
			_ = _nativeLen0_40
			for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
				(_nw221).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
			}
		}
		_1175_v221 = _nw221
		var _1178_v222 _dafny.Set
		_ = _1178_v222
		_1178_v222 = _dafny.SetOf(_1175_v221)
		r2 = ((_1178_v222).Intersection(_1178_v222)).Union(_1178_v222)
		var _1179_v223 _dafny.Array
		_ = _1179_v223
		var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
		_ = _nw222
		_1179_v223 = _nw222
		r3 = _1179_v223
		return r0, r1, r2, r3
	}
}
func (_this *C4) F10() _dafny.Int {
	{
		return _this._f10
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
