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
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SetOf(false, true), _dafny.SetOf(true, true))).Union(_dafny.SetOf(_dafny.SetOf(false, true), _dafny.SetOf(false), _dafny.SetOf(true, false), _dafny.SetOf(false)))).Difference((_dafny.SetOf(_dafny.SetOf(false), _dafny.SetOf(false))))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("ed")
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("knr"), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(220))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))).Cardinality()), _dafny.IntOfInt64(771)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("cvrf"), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(817))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if (_dafny.MultiSetOf(_dafny.IntOfInt64(817))).Contains(_1_v0) {
				_coll0.Add((_1_v0).Times((_dafny.SetOf(true)).Cardinality()), (_dafny.MultiSetOf(false, false)).Cardinality())
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())).Cardinality()))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("dujbkapsg"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-864))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(733)
	}))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	return !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uh"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-759))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("snaam"), _dafny.UnicodeSeqOfUtf8Bytes("uupaiodb")))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, p3 D1, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(133))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	})), _dafny.UnicodeSeqOfUtf8Bytes("bndem"))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), (func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(false)
		}
		return _dafny.SeqOf(true)
	})())
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-535), _dafny.IntOfInt64(990))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v0 _dafny.Int
			_5_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(-535)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(990)) < 0) {
				_coll1.Add((_5_v0).Plus(_dafny.IntOfInt64(-90)), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-907), _dafny.IntOfInt64(-667))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(521))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg5 _dafny.Int) interface{} {
						return coer5(arg5)
					}
				}(func(_7_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}))).Cardinality()), (_dafny.MultiSetOf(func() _dafny.Map {
					var _coll2 = _dafny.NewMapBuilder()
					_ = _coll2
					for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(286), _dafny.IntOfInt64(-321))); ; {
						_compr_2, _ok2 := _iter2()
						if !_ok2 {
							break
						}
						var _8_v1 _dafny.Int
						_8_v1 = interface{}(_compr_2).(_dafny.Int)
						if ((_dafny.IntOfInt64(286)).Cmp(_8_v1) <= 0) && ((_8_v1).Cmp(_dafny.IntOfInt64(-321)) < 0) {
							_coll2.Add(Companion_Default___.SafeDivisionInt(_8_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("srsgjqmes")).Cardinality())))).Cardinality())).Cardinality()), _dafny.IntOfInt64(32))
						}
					}
					return _coll2.ToMap()
				}())).Cardinality())).Cardinality()))
			}
		}
		return _coll1.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-841), (_dafny.MultiSetOf(true)).Cardinality())).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(115))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((Companion_D15_.Create_DC33_(_dafny.UnicodeSeqOfUtf8Bytes("rcx"))).Dtor_cf54()).Cardinality()), _dafny.IntOfInt64(692))))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))
}
func (_static *CompanionStruct_Default___) Fm16(p0 D4, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, !(!(true)))).Difference(_dafny.SetOf(true, false))).Difference(_dafny.SetOf(true, false))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	var _source0 D9 = Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hvkyd")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(609), _dafny.IntOfInt64(-451))).Cardinality())), _dafny.SetOf(!(true))))
	_ = _source0
	if _source0.Is_DC19() {
		var _9___mcc_h0 _dafny.Int = _source0.Get_().(D9_DC19).Cf29
		_ = _9___mcc_h0
		var _10___mcc_h1 _dafny.Int = _source0.Get_().(D9_DC19).Cf30
		_ = _10___mcc_h1
		var _11_cf30 _dafny.Int = _10___mcc_h1
		_ = _11_cf30
		var _12_cf29 _dafny.Int = _9___mcc_h0
		_ = _12_cf29
		return (_dafny.Zero).Minus(_12_cf29)
	} else if _source0.Is_DC20() {
		var _13___mcc_h2 bool = _source0.Get_().(D9_DC20).Cf31
		_ = _13___mcc_h2
		var _14___mcc_h3 bool = _source0.Get_().(D9_DC20).Cf32
		_ = _14___mcc_h3
		var _15___mcc_h4 bool = _source0.Get_().(D9_DC20).Cf33
		_ = _15___mcc_h4
		var _16_cf33 bool = _15___mcc_h4
		_ = _16_cf33
		var _17_cf32 bool = _14___mcc_h3
		_ = _17_cf32
		var _18_cf31 bool = _13___mcc_h2
		_ = _18_cf31
		return ((_dafny.SetOf(_dafny.SeqOf(_dafny.Zero), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gynrfk")).Cardinality())), (_dafny.SetOf(!(_16_cf33))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wvp")).Cardinality())))).Intersection(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_19_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(432)
		})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_18_cf31)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gynpbladr")).Cardinality()))))).Cardinality()
	} else {
		var _20___mcc_h5 _dafny.Map = _source0.Get_().(D9_DC18).Cf28
		_ = _20___mcc_h5
		var _21_cf28 _dafny.Map = _20___mcc_h5
		_ = _21_cf28
		return ((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hyoldvcp")).Cardinality()), _dafny.IntOfInt64(978), _dafny.IntOfInt64(255), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iflkrw")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))).Cardinality()).Plus((_dafny.MultiSetOf(false)).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iqunqvo")).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-829))).Cardinality()))).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(495))).Uint32(), func(coer7 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_22_i0 _dafny.Int) D1 {
		return Companion_D1_.Create_DC4_(_dafny.IntOfInt64(82), _dafny.CodePoint('y'), false, _dafny.IntOfInt64(805))
	}))).Cardinality()))).Intersection(_dafny.SetOf(_dafny.IntOfInt64(314)))).Difference(func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-820), _dafny.IntOfInt64(382))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _23_v0 _dafny.Int
			_23_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(-820)).Cmp(_23_v0) <= 0) && ((_23_v0).Cmp(_dafny.IntOfInt64(382)) < 0) {
				_coll3.Add(Companion_Default___.SafeDivisionInt(_23_v0, _dafny.IntOfInt64(599)))
			}
		}
		return _coll3.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D15_.Create_DC33_(_dafny.UnicodeSeqOfUtf8Bytes("xebgvpgac"))).Dtor_cf54()
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(932), (func() _dafny.Int {
		if true {
			return (_dafny.SetOf((_dafny.SetOf(false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(563))).Cardinality())).Cardinality()
		}
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())))).Cardinality(), _dafny.IntOfInt64(265))).Cardinality()
	})(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality()), _dafny.IntOfInt64(-658), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC21_(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false))), false)).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(918), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vkmbe"), _dafny.UnicodeSeqOfUtf8Bytes("r"))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("vogq"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("q"), _dafny.UnicodeSeqOfUtf8Bytes("y")))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true, false, (Companion_D1_.Create_DC4_((_dafny.SetOf((Companion_D15_.Create_DC34_(_dafny.UnicodeSeqOfUtf8Bytes("ghf"), _dafny.IntOfInt64(164), false, Companion_D10_.Create_DC22_(), _dafny.IntOfInt64(693))).Dtor_cf57(), true)).Cardinality(), _dafny.CodePoint('c'), true, _dafny.IntOfInt64(228))).Dtor_cf8()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if !((Companion_D15_.Create_DC32_(false)).Dtor_cf53()) {
			return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jac"))
		}
		return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("hwf"))
	})()).Union((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("en"))).Difference(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sicipu"))))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if (_dafny.IntOfInt64(216)).Cmp(_dafny.IntOfInt64(-12)) != 0 {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (Companion_D15_.Create_DC33_(_dafny.UnicodeSeqOfUtf8Bytes("fbkybdo"))).Dtor_cf54())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(!(!(false))))), _dafny.UnicodeSeqOfUtf8Bytes("vudhbc")))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.UnicodeSeqOfUtf8Bytes("jpc"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("fgwlc")))
	}
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fykfcgc")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-242), _dafny.IntOfInt64(517)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iglpstssc")).Cardinality()), _dafny.IntOfInt64(-663)))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source1 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("bmoihfr"))).Cardinality())).Cardinality()), _dafny.IntOfInt64(715)), true)
	_ = _source1
	var _24___mcc_h0 _dafny.Map = _source1
	_ = _24___mcc_h0
	var _25_cf18 _dafny.Map = _24___mcc_h0
	_ = _25_cf18
	return _dafny.CodePoint('e')
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfInt64(-55), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(585), _dafny.IntOfInt64(832))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-24)), (_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality()), _dafny.IntOfInt64(249))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(697)))).Union((_dafny.MultiSetOf(_dafny.IntOfInt64(747), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(34), _dafny.IntOfInt64(873), _dafny.IntOfInt64(-285), (_dafny.SetOf(true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(138))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_26_i0 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(164), _dafny.IntOfInt64(-392))
		})), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ykxkeccls")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xrtrgqvhp")).Cardinality()))), func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-447), _dafny.IntOfInt64(500))); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _27_v1 _dafny.Int
				_27_v1 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(-447)).Cmp(_27_v1) <= 0) && ((_27_v1).Cmp(_dafny.IntOfInt64(500)) < 0) {
					_coll5.Add((_27_v1).Times(_dafny.IntOfInt64(226)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())
				}
			}
			return _coll5.ToMap()
		}()))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _28_v0 _dafny.Map
			_28_v0 = interface{}(_compr_4).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(138))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_26_i0 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(164), _dafny.IntOfInt64(-392))
			})), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ykxkeccls")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xrtrgqvhp")).Cardinality()))), func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-447), _dafny.IntOfInt64(500))); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _27_v1 _dafny.Int
					_27_v1 = interface{}(_compr_6).(_dafny.Int)
					if ((_dafny.IntOfInt64(-447)).Cmp(_27_v1) <= 0) && ((_27_v1).Cmp(_dafny.IntOfInt64(500)) < 0) {
						_coll6.Add((_27_v1).Times(_dafny.IntOfInt64(226)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())
					}
				}
				return _coll6.ToMap()
			}())), _28_v0) {
				_coll4.Add(_28_v0, _dafny.SetOf(!(false), false, !(!(false))))
			}
		}
		return _coll4.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Sequence, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D1 {
	var _source2 D1 = Companion_D1_.Create_DC2_(false)
	_ = _source2
	if _source2.Is_DC3() {
		var _29___mcc_h0 _dafny.Int = _source2.Get_().(D1_DC3).Cf5
		_ = _29___mcc_h0
		var _30_cf5 _dafny.Int = _29___mcc_h0
		_ = _30_cf5
		return Companion_D1_.Create_DC2_(false)
	} else if _source2.Is_DC4() {
		var _31___mcc_h1 _dafny.Int = _source2.Get_().(D1_DC4).Cf6
		_ = _31___mcc_h1
		var _32___mcc_h2 _dafny.CodePoint = _source2.Get_().(D1_DC4).Cf7
		_ = _32___mcc_h2
		var _33___mcc_h3 bool = _source2.Get_().(D1_DC4).Cf8
		_ = _33___mcc_h3
		var _34___mcc_h4 _dafny.Int = _source2.Get_().(D1_DC4).Cf9
		_ = _34___mcc_h4
		var _35_cf9 _dafny.Int = _34___mcc_h4
		_ = _35_cf9
		var _36_cf8 bool = _33___mcc_h3
		_ = _36_cf8
		var _37_cf7 _dafny.CodePoint = _32___mcc_h2
		_ = _37_cf7
		var _38_cf6 _dafny.Int = _31___mcc_h1
		_ = _38_cf6
		return Companion_D1_.Create_DC2_(_36_cf8)
	} else {
		var _39___mcc_h5 bool = _source2.Get_().(D1_DC2).Cf4
		_ = _39___mcc_h5
		var _40_cf4 bool = _39___mcc_h5
		_ = _40_cf4
		return Companion_D1_.Create_DC2_(_40_cf4)
	}
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _dafny.IntOfInt64(902))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), _dafny.IntOfInt64(712)))).Keys().Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _41_v0 _dafny.CodePoint
			_41_v0 = interface{}(_compr_7).(_dafny.CodePoint)
			if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _dafny.IntOfInt64(902))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), _dafny.IntOfInt64(712)))).Contains(_41_v0) {
				_coll7.Add(_41_v0, !(false))
			}
		}
		return _coll7.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_((_dafny.IntOfInt64(522)).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-290))).Cardinality())), (Companion_D1_.Create_DC4_(_dafny.IntOfInt64(721), _dafny.CodePoint('d'), false, _dafny.IntOfInt64(427))).Dtor_cf7(), !(!(false)), (func() _dafny.Int {
		if true {
			return _dafny.IntOfInt64(579)
		}
		return _dafny.IntOfInt64(602)
	})())
}
func (_static *CompanionStruct_Default___) M6(globalState *GlobalState) (_dafny.Sequence, T0) {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var r1 T0 = (T0)(nil)
	_ = r1
	var _42_v0 *C8
	_ = _42_v0
	var _nw0 *C8 = New_C8_()
	_ = _nw0
	_nw0.Ctor__()
	_42_v0 = _nw0
	var _nw1 *C8 = New_C8_()
	_ = _nw1
	_nw1.Ctor__()
	_42_v0 = _nw1
	var _43_v1 _dafny.Array
	_ = _43_v1
	var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
	_ = _nw2
	_43_v1 = _nw2
	var _44_v2 bool
	_ = _44_v2
	_44_v2 = false
	var _45_v3 _dafny.Map
	_ = _45_v3
	_45_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_44_v2, _dafny.CodePoint('i'))
	var _46_v4 _dafny.CodePoint
	_ = _46_v4
	_46_v4 = _dafny.CodePoint('v')
	var _47_v5 _dafny.MultiSet
	_ = _47_v5
	_47_v5 = _dafny.MultiSetOf(_44_v2)
	var _48_v6 _dafny.Map
	_ = _48_v6
	_48_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
		if (_45_v3).Contains(_44_v2) {
			return (_45_v3).Get(_44_v2).(_dafny.CodePoint)
		}
		return _46_v4
	})(), _47_v5)
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_43_v1), 0))
	_ = _index0
	(_43_v1).ArraySet1((func() _dafny.MultiSet {
		if (_48_v6).Contains(_46_v4) {
			return (_48_v6).Get(_46_v4).(_dafny.MultiSet)
		}
		return _47_v5
	})(), (_index0).Int())
	var _49_v7 _dafny.Map
	_ = _49_v7
	_49_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_44_v2, _44_v2)
	var _50_v8 _dafny.Int
	_ = _50_v8
	_50_v8 = _dafny.IntOfInt64(94)
	var _51_v9 _dafny.Map
	_ = _51_v9
	_51_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v8, _44_v2)
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_43_v1), 0))
	_ = _index1
	(_43_v1).ArraySet1((_42_v0).Fm0((_49_v7).Equals(_49_v7), (func() bool {
		if (func() bool {
			if (_51_v9).Contains(_50_v8) {
				return (_51_v9).Get(_50_v8).(bool)
			}
			return false
		})() {
			return !(_44_v2)
		}
		return _44_v2
	})(), _50_v8, _50_v8, globalState), (_index1).Int())
	var _52_v10 _dafny.Array
	_ = _52_v10
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(19)
	_ = _len0_0
	var _nw3 _dafny.Array
	_ = _nw3
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw3 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = func(_53_i0 _dafny.Int) _dafny.Int {
			return (_53_i0).Times(_dafny.IntOfInt64(800))
		}
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
	_52_v10 = _nw3
	var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
	_ = _nw4
	_52_v10 = _nw4
	_42_v0 = _42_v0
	var _54_v11 _dafny.Sequence
	_ = _54_v11
	_54_v11 = _dafny.UnicodeSeqOfUtf8Bytes("ny")
	_50_v8 = (func() _dafny.Int {
		if _44_v2 {
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_54_v11, (Companion_Default___.SafeIndex(_50_v8, _dafny.IntOfUint32((_54_v11).Cardinality()))).Uint32(), _dafny.CodePoint('y'))).Cardinality())
		}
		return _50_v8
	})()
	var _55_v12 D4
	_ = _55_v12
	_55_v12 = Companion_D4_.Create_DC11_(Companion_D4_.Create_DC9_(_52_v10))
	var _56_v13 D4
	_ = _56_v13
	_56_v13 = Companion_D4_.Create_DC11_(_55_v12)
	var _57_v14 _dafny.Map
	_ = _57_v14
	_57_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v13, Companion_Default___.Fm19(false, _dafny.IntOfInt64(961), globalState))
	var _58_v15 _dafny.Sequence
	_ = _58_v15
	_58_v15 = _dafny.SeqOf(_50_v8)
	_44_v2 = !(((func() _dafny.Int {
		if (_57_v14).Contains(_56_v13) {
			return (_57_v14).Get(_56_v13).(_dafny.Int)
		}
		return (_58_v15).Select((Companion_Default___.SafeIndex(_50_v8, _dafny.IntOfUint32((_58_v15).Cardinality()))).Uint32()).(_dafny.Int)
	})()).Cmp(_50_v8) <= 0)
	r0 = _dafny.Companion_Sequence_.Concatenate(_54_v11, _54_v11)
	var _59_v16 T0
	_ = _59_v16
	var _nw5 *C8 = New_C8_()
	_ = _nw5
	_nw5.Ctor__()
	_59_v16 = _nw5
	r1 = _59_v16
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _60_v0 _dafny.Array
	_ = _60_v0
	var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
	_ = _nw6
	_60_v0 = _nw6
	var _61_globalState *GlobalState
	_ = _61_globalState
	var _nw7 *GlobalState = New_GlobalState_()
	_ = _nw7
	_nw7.Ctor__(_60_v0, _dafny.IntOfInt64(-481))
	_61_globalState = _nw7
	var _62_v1 bool
	_ = _62_v1
	_62_v1 = false
	if _62_v1 {
		var _63_v2 _dafny.Int
		_ = _63_v2
		_63_v2 = _dafny.IntOfInt64(7)
		_63_v2 = _dafny.IntOfInt64(547)
		var _64_v3 _dafny.MultiSet
		_ = _64_v3
		_64_v3 = _dafny.MultiSetOf(_62_v1, false)
		var _65_v4 _dafny.Array
		_ = _65_v4
		var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw8
		_65_v4 = _nw8
		var _66_v5 *C6
		_ = _66_v5
		var _nw9 *C6 = New_C6_()
		_ = _nw9
		_nw9.Ctor__((_64_v3).Update(_62_v1, Companion_Default___.Abs(_63_v2)), _65_v4)
		_66_v5 = _nw9
		var _67_v6 _dafny.Sequence
		_ = _67_v6
		_67_v6 = _dafny.UnicodeSeqOfUtf8Bytes("yymfboxka")
		_67_v6 = _dafny.Companion_Sequence_.Concatenate(_67_v6, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm4(_62_v1, _62_v1, _62_v1, _61_globalState), _67_v6))
		var _68_v7 _dafny.Sequence
		_ = _68_v7
		_68_v7 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("kahofehgp"))
		var _69_v8 *C5
		_ = _69_v8
		var _nw10 *C5 = New_C5_()
		_ = _nw10
		_nw10.Ctor__((_68_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xiko")).Cardinality()), _dafny.IntOfUint32((_68_v7).Cardinality()))).Uint32()).(_dafny.Sequence))
		_69_v8 = _nw10
		var _70_v9 _dafny.Array
		_ = _70_v9
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
		_ = _nw11
		_70_v9 = _nw11
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_70_v9), 0))
		_ = _index2
		(_70_v9).ArraySet1(_60_v0, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_70_v9), 0))
		_ = _index3
		(_70_v9).ArraySet1(_60_v0, (_index3).Int())
	} else {
		var _71_v10 _dafny.Int
		_ = _71_v10
		_71_v10 = _dafny.IntOfInt64(849)
		var _72_v11 _dafny.Map
		_ = _72_v11
		_72_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v10, _71_v10)
		var _73_v12 _dafny.Sequence
		_ = _73_v12
		_73_v12 = _dafny.SeqOf(_62_v1, _62_v1)
		var _74_v13 _dafny.Sequence
		_ = _74_v13
		_74_v13 = _dafny.SeqOf(_71_v10, _71_v10)
		var _75_v14 _dafny.Map
		_ = _75_v14
		_75_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v10, _62_v1)
		_72_v11 = (_72_v11).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_73_v12, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_74_v13).Cardinality()), _dafny.IntOfUint32((_73_v12).Cardinality()))).Uint32(), (func() bool {
			if (_75_v14).Contains(_71_v10) {
				return (_75_v14).Get(_71_v10).(bool)
			}
			return _62_v1
		})())).Cardinality()), _71_v10)
		var _76_v15 _dafny.CodePoint
		_ = _76_v15
		_76_v15 = _dafny.CodePoint('f')
		var _77_v16 _dafny.Array
		_ = _77_v16
		var _nwElement0_0 _dafny.CodePoint = _76_v15
		_ = _nwElement0_0
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(8))
		_ = _nw12
		(_nw12).ArraySet1CodePoint(_nwElement0_0, 0)
		(_nw12).ArraySet1CodePoint(_76_v15, 1)
		(_nw12).ArraySet1CodePoint(_dafny.CodePoint('f'), 2)
		(_nw12).ArraySet1CodePoint(_76_v15, 3)
		(_nw12).ArraySet1CodePoint(_76_v15, 4)
		(_nw12).ArraySet1CodePoint(_76_v15, 5)
		(_nw12).ArraySet1CodePoint(_76_v15, 6)
		(_nw12).ArraySet1CodePoint(Companion_Default___.Fm30(Companion_Default___.Fm22(_62_v1, (_dafny.Zero).Minus(_71_v10), _61_globalState), _71_v10, _dafny.IntOfInt64(581), _61_globalState), 7)
		_77_v16 = _nw12
		var _78_v17 _dafny.Sequence
		_ = _78_v17
		_78_v17 = _dafny.UnicodeSeqOfUtf8Bytes("i")
		var _79_v18 _dafny.Set
		_ = _79_v18
		_79_v18 = _dafny.SetOf(_62_v1)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_77_v16), 0))
		_ = _index4
		(_77_v16).ArraySet1CodePoint((_78_v17).Select((Companion_Default___.SafeIndex((_79_v18).Cardinality(), _dafny.IntOfUint32((_78_v17).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((_77_v16), 0))
		_ = _index5
		(_77_v16).ArraySet1CodePoint(_76_v15, (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))
		_ = _index6
		(_60_v0).ArraySet1(_62_v1, (_index6).Int())
		var _80_v19 T1
		_ = _80_v19
		var _nw13 *C4 = New_C4_()
		_ = _nw13
		_nw13.Ctor__(_dafny.IntOfInt64(232))
		_80_v19 = _nw13
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))
		_ = _index7
		var _rhs0 bool = _62_v1
		_ = _rhs0
		var _rhs1 _dafny.Int = _71_v10
		_ = _rhs1
		var _rhs2 _dafny.Sequence = _78_v17
		_ = _rhs2
		var _rhs3 bool = _62_v1
		_ = _rhs3
		var _rhs4 T1 = _80_v19
		_ = _rhs4
		var _lhs0 _dafny.Array = _60_v0
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))
		_ = _lhs1
		_62_v1 = _rhs0
		_71_v10 = _rhs1
		_78_v17 = _rhs2
		(_lhs0).ArraySet1(_rhs3, (_lhs1).Int())
		_80_v19 = _rhs4
		var _81_v20 _dafny.Array
		_ = _81_v20
		var _nwElement0_1 bool = (_60_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))).Int()).(bool)
		_ = _nwElement0_1
		var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(14))
		_ = _nw14
		(_nw14).ArraySet1(_nwElement0_1, 0)
		(_nw14).ArraySet1((_60_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))).Int()).(bool), 1)
		(_nw14).ArraySet1((_60_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))).Int()).(bool), 2)
		(_nw14).ArraySet1(_62_v1, 3)
		(_nw14).ArraySet1((_60_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))).Int()).(bool), 4)
		(_nw14).ArraySet1((_60_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))).Int()).(bool), 5)
		(_nw14).ArraySet1(_62_v1, 6)
		(_nw14).ArraySet1((_60_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))).Int()).(bool), 7)
		(_nw14).ArraySet1((_60_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_60_v0), 0))).Int()).(bool), 8)
		(_nw14).ArraySet1(_62_v1, 9)
		(_nw14).ArraySet1(_62_v1, 10)
		(_nw14).ArraySet1(true, 11)
		(_nw14).ArraySet1(_62_v1, 12)
		(_nw14).ArraySet1(_62_v1, 13)
		_81_v20 = _nw14
		var _82_v21 _dafny.Sequence
		_ = _82_v21
		_82_v21 = _dafny.SeqOf(_81_v20, _81_v20)
		var _83_v22 _dafny.Array
		_ = _83_v22
		var _nwElement0_2 _dafny.Array = _81_v20
		_ = _nwElement0_2
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(26))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_2, 0)
		(_nw15).ArraySet1(_81_v20, 1)
		(_nw15).ArraySet1(_81_v20, 2)
		(_nw15).ArraySet1(_81_v20, 3)
		(_nw15).ArraySet1(_81_v20, 4)
		(_nw15).ArraySet1(_81_v20, 5)
		(_nw15).ArraySet1(_81_v20, 6)
		(_nw15).ArraySet1(_81_v20, 7)
		(_nw15).ArraySet1(_81_v20, 8)
		(_nw15).ArraySet1(_81_v20, 9)
		(_nw15).ArraySet1(_81_v20, 10)
		(_nw15).ArraySet1((_82_v21).Select((Companion_Default___.SafeIndex(_71_v10, _dafny.IntOfUint32((_82_v21).Cardinality()))).Uint32()).(_dafny.Array), 11)
		(_nw15).ArraySet1(_81_v20, 12)
		(_nw15).ArraySet1(_81_v20, 13)
		(_nw15).ArraySet1(_81_v20, 14)
		(_nw15).ArraySet1(_81_v20, 15)
		(_nw15).ArraySet1(_81_v20, 16)
		(_nw15).ArraySet1(_81_v20, 17)
		(_nw15).ArraySet1(_81_v20, 18)
		(_nw15).ArraySet1(_81_v20, 19)
		(_nw15).ArraySet1(_81_v20, 20)
		(_nw15).ArraySet1(_81_v20, 21)
		(_nw15).ArraySet1(_81_v20, 22)
		(_nw15).ArraySet1(_81_v20, 23)
		(_nw15).ArraySet1(_81_v20, 24)
		(_nw15).ArraySet1(_81_v20, 25)
		_83_v22 = _nw15
		var _84_v23 bool
		_ = _84_v23
		var _85_v24 _dafny.Map
		_ = _85_v24
		var _out0 bool
		_ = _out0
		var _out1 _dafny.Map
		_ = _out1
		_out0, _out1 = (_80_v19).M0(_83_v22, _61_globalState)
		_84_v23 = _out0
		_85_v24 = _out1
		var _86_v25 *C4
		_ = _86_v25
		var _nw16 *C4 = New_C4_()
		_ = _nw16
		_nw16.Ctor__(_71_v10)
		_86_v25 = _nw16
		_86_v25 = _86_v25
	}
	var _87_v26 _dafny.Int
	_ = _87_v26
	_87_v26 = _dafny.IntOfInt64(542)
	var _88_v27 _dafny.Map
	_ = _88_v27
	_88_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v26, _87_v26)
	var _89_v28 _dafny.Sequence
	_ = _89_v28
	_89_v28 = _dafny.SeqOf(_87_v26)
	var _90_v29 _dafny.Sequence
	_ = _90_v29
	_90_v29 = _dafny.SeqOf((_89_v28).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm24(_dafny.IntOfInt64(904), _62_v1, _62_v1, _87_v26, _61_globalState), _dafny.IntOfUint32((_89_v28).Cardinality()))).Uint32()).(_dafny.Int))
	var _91_v30 _dafny.Map
	_ = _91_v30
	_91_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v26, _89_v28)
	_87_v26 = (func() _dafny.Int {
		if (_88_v27).Contains((_91_v30).Cardinality()) {
			return (_88_v27).Get((_91_v30).Cardinality()).(_dafny.Int)
		}
		return (func() _dafny.Int {
			if true {
				return _87_v26
			}
			return _87_v26
		})()
	})()
	var _92_v31 _dafny.CodePoint
	_ = _92_v31
	_92_v31 = _dafny.CodePoint('a')
	var _source3 D1 = Companion_D1_.Create_DC4_(_87_v26, _92_v31, _62_v1, _87_v26)
	_ = _source3
	if _source3.Is_DC3() {
		var _93___mcc_h0 _dafny.Int = _source3.Get_().(D1_DC3).Cf5
		_ = _93___mcc_h0
		var _94_cf5 _dafny.Int = _93___mcc_h0
		_ = _94_cf5
		var _95_v32 _dafny.Sequence
		_ = _95_v32
		var _96_v33 T0
		_ = _96_v33
		var _out2 _dafny.Sequence
		_ = _out2
		var _out3 T0
		_ = _out3
		_out2, _out3 = Companion_Default___.M6(_61_globalState)
		_95_v32 = _out2
		_96_v33 = _out3
		var _97_v34 D9
		_ = _97_v34
		_97_v34 = Companion_D9_.Create_DC19_(((_89_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.IntOfUint32((_89_v28).Cardinality()))).Uint32()).(_dafny.Int)).Plus((_dafny.MultiSetOf(_62_v1)).Cardinality()), _94_cf5)
		var _98_v35 D2
		_ = _98_v35
		_98_v35 = Companion_D2_.Create_DC6_(_94_cf5, _62_v1)
		var _rhs5 D9 = _97_v34
		_ = _rhs5
		var _rhs6 _dafny.Int = Companion_Default___.SafeModuloInt(_87_v26, (_dafny.Zero).Minus(Companion_Default___.Fm19((_98_v35).Dtor_cf12(), _94_cf5, _61_globalState)))
		_ = _rhs6
		_97_v34 = _rhs5
		_94_cf5 = _rhs6
		var _99_v36 _dafny.Map
		_ = _99_v36
		_99_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v1, _62_v1)
		var _100_v37 D15
		_ = _100_v37
		_100_v37 = Companion_D15_.Create_DC31_(_99_v36)
		_99_v36 = (_100_v37).Dtor_cf52()
		var _101_v38 _dafny.Sequence
		_ = _101_v38
		_101_v38 = _dafny.SeqOf(_62_v1, _62_v1)
		_101_v38 = _dafny.Companion_Sequence_.Update(_101_v38, (Companion_Default___.SafeIndex(_87_v26, _dafny.IntOfUint32((_101_v38).Cardinality()))).Uint32(), _62_v1)
	} else if _source3.Is_DC4() {
		var _102___mcc_h1 _dafny.Int = _source3.Get_().(D1_DC4).Cf6
		_ = _102___mcc_h1
		var _103___mcc_h2 _dafny.CodePoint = _source3.Get_().(D1_DC4).Cf7
		_ = _103___mcc_h2
		var _104___mcc_h3 bool = _source3.Get_().(D1_DC4).Cf8
		_ = _104___mcc_h3
		var _105___mcc_h4 _dafny.Int = _source3.Get_().(D1_DC4).Cf9
		_ = _105___mcc_h4
		var _106_cf9 _dafny.Int = _105___mcc_h4
		_ = _106_cf9
		var _107_cf8 bool = _104___mcc_h3
		_ = _107_cf8
		var _108_cf7 _dafny.CodePoint = _103___mcc_h2
		_ = _108_cf7
		var _109_cf6 _dafny.Int = _102___mcc_h1
		_ = _109_cf6
		var _110_v39 *C3
		_ = _110_v39
		var _nw17 *C3 = New_C3_()
		_ = _nw17
		_nw17.Ctor__()
		_110_v39 = _nw17
		var _111_v40 _dafny.Map
		_ = _111_v40
		_111_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm14(_87_v26, _dafny.IntOfInt64(-263), _61_globalState), (func() *C3 {
			if _62_v1 {
				return _110_v39
			}
			return _110_v39
		})())
		var _112_v41 _dafny.Sequence
		_ = _112_v41
		_112_v41 = _dafny.SeqOf(_107_cf8, _107_cf8)
		_111_v40 = (_111_v40).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_112_v41, _112_v41)).Cardinality()), _110_v39)
		_87_v26 = _109_cf6
		_107_cf8 = ((_106_cf9).Plus(_109_cf6)).Cmp(_87_v26) != 0
	} else {
		var _113___mcc_h5 bool = _source3.Get_().(D1_DC2).Cf4
		_ = _113___mcc_h5
		var _114_cf4 bool = _113___mcc_h5
		_ = _114_cf4
		_62_v1 = (_87_v26).Cmp(_87_v26) != 0
		var _115_v42 _dafny.Sequence
		_ = _115_v42
		_115_v42 = _dafny.UnicodeSeqOfUtf8Bytes("aptekw")
		var _116_v43 _dafny.MultiSet
		_ = _116_v43
		_116_v43 = _dafny.MultiSetOf(_dafny.IntOfInt64(278), _87_v26, _dafny.IntOfUint32((_115_v42).Cardinality()))
		_114_cf4 = Companion_Default___.Fm8(_116_v43, _61_globalState)
		var _117_v44 _dafny.Map
		_ = _117_v44
		_117_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_cf4, Companion_Default___.Fm4(!(_114_cf4), _62_v1, _114_cf4, _61_globalState))
		var _118_v45 _dafny.MultiSet
		_ = _118_v45
		_118_v45 = _dafny.MultiSetOf(_114_cf4)
		var _119_v46 *C1
		_ = _119_v46
		var _nw18 *C1 = New_C1_()
		_ = _nw18
		_nw18.Ctor__(_118_v45, (Companion_Default___.Fm35(false, _62_v1, _62_v1, _87_v26, _61_globalState)).Cardinality())
		_119_v46 = _nw18
		var _120_v47 _dafny.Array
		_ = _120_v47
		var _nwElement0_3 *C1 = _119_v46
		_ = _nwElement0_3
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(2))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_3, 0)
		(_nw19).ArraySet1(_119_v46, 1)
		_120_v47 = _nw19
		var _121_v48 _dafny.Sequence
		_ = _121_v48
		_121_v48 = _dafny.SeqOf(_120_v47)
		var _rhs7 bool = !_dafny.Companion_Sequence_.Contains(_121_v48, _120_v47)
		_ = _rhs7
		var _rhs8 _dafny.Int = (_119_v46.F11).Minus((_dafny.Zero).Minus(_119_v46.F11))
		_ = _rhs8
		var _rhs9 _dafny.Map = (((_117_v44).Update(_62_v1, _115_v42)).Update(_62_v1, _115_v42)).Merge((_117_v44).Update(!(_62_v1), _115_v42))
		_ = _rhs9
		_114_cf4 = _rhs7
		_87_v26 = _rhs8
		_117_v44 = _rhs9
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_60_v0), 0))
		_ = _index8
		(_60_v0).ArraySet1(true, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_60_v0), 0))
		_ = _index9
		(_60_v0).ArraySet1(!(_62_v1) || (_114_cf4), (_index9).Int())
	}
	_62_v1 = _62_v1
	_87_v26 = Companion_Default___.SafeDivisionInt(_87_v26, _87_v26)
	var _122_v49 _dafny.MultiSet
	_ = _122_v49
	_122_v49 = _dafny.MultiSetOf(_62_v1, _62_v1, _62_v1, _62_v1, _62_v1)
	var _123_v50 T1
	_ = _123_v50
	var _nw20 *C1 = New_C1_()
	_ = _nw20
	_nw20.Ctor__(_122_v49, _87_v26)
	_123_v50 = _nw20
	var _124_v51 _dafny.Sequence
	_ = _124_v51
	_124_v51 = _dafny.UnicodeSeqOfUtf8Bytes("nmpkjd")
	var _125_v52 _dafny.Set
	_ = _125_v52
	_125_v52 = _dafny.SetOf(_87_v26)
	var _rhs10 bool = _62_v1
	_ = _rhs10
	var _rhs11 bool = _62_v1
	_ = _rhs11
	var _rhs12 T1 = _123_v50
	_ = _rhs12
	var _rhs13 _dafny.CodePoint = Companion_Default___.Fm30(_124_v51, _87_v26, (_87_v26).Plus((_125_v52).Cardinality()), _61_globalState)
	_ = _rhs13
	var _rhs14 bool = (func() bool {
		if _62_v1 {
			return _62_v1
		}
		return _62_v1
	})()
	_ = _rhs14
	_62_v1 = _rhs10
	_62_v1 = _rhs11
	_123_v50 = _rhs12
	_92_v31 = _rhs13
	_62_v1 = _rhs14
	var _126_v53 D1
	_ = _126_v53
	_126_v53 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_87_v26), _92_v31, _62_v1, Companion_Default___.Fm24(_87_v26, _62_v1, _62_v1, _87_v26, _61_globalState))
	var _127_v54 _dafny.Array
	_ = _127_v54
	var _nwElement0_4 D1 = _126_v53
	_ = _nwElement0_4
	var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(11))
	_ = _nw21
	(_nw21).ArraySet1(_nwElement0_4, 0)
	(_nw21).ArraySet1(_126_v53, 1)
	(_nw21).ArraySet1(_126_v53, 2)
	(_nw21).ArraySet1(_126_v53, 3)
	(_nw21).ArraySet1(_126_v53, 4)
	(_nw21).ArraySet1(_126_v53, 5)
	(_nw21).ArraySet1(Companion_D1_.Create_DC4_(_87_v26, _92_v31, _62_v1, _87_v26), 6)
	(_nw21).ArraySet1(Companion_Default___.Fm36(_61_globalState), 7)
	(_nw21).ArraySet1(_126_v53, 8)
	(_nw21).ArraySet1(_126_v53, 9)
	(_nw21).ArraySet1(Companion_Default___.Fm36(_61_globalState), 10)
	_127_v54 = _nw21
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_127_v54), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _128_i0 _dafny.Int
		_128_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_128_i0).Sign() != -1) && ((_128_i0).Cmp(_dafny.ArrayLen((_127_v54), 0)) < 0)) {
			(_127_v54).ArraySet1(_126_v53, (_128_i0).Int())
		}
	}
	var _129_v55 _dafny.Array
	_ = _129_v55
	var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
	_ = _nw22
	_129_v55 = _nw22
	var _130_v56 bool
	_ = _130_v56
	var _131_v57 _dafny.Map
	_ = _131_v57
	var _out4 bool
	_ = _out4
	var _out5 _dafny.Map
	_ = _out5
	_out4, _out5 = (_123_v50).M0(_129_v55, _61_globalState)
	_130_v56 = _out4
	_131_v57 = _out5
	_130_v56 = _130_v56
	var _hi0 _dafny.Int = _87_v26
	_ = _hi0
	for _132_i1 := (_87_v26).Minus(_87_v26); _132_i1.Cmp(_hi0) < 0; _132_i1 = _132_i1.Plus(_dafny.One) {
		var _133_v58 _dafny.Map
		_ = _133_v58
		_133_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v31, _130_v56)
		_62_v1 = (func() bool {
			if _62_v1 {
				return !((_133_v58).Update(_92_v31, false)).Equals(_133_v58)
			}
			return !_dafny.Companion_Sequence_.Equal(_89_v28, _dafny.SeqOf(_132_i1))
		})()
		var _134_v59 _dafny.Array
		_ = _134_v59
		var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw23
		_134_v59 = _nw23
		var _135_v60 _dafny.Array
		_ = _135_v60
		var _nwElement0_5 _dafny.Array = _134_v59
		_ = _nwElement0_5
		var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(12))
		_ = _nw24
		(_nw24).ArraySet1(_nwElement0_5, 0)
		(_nw24).ArraySet1(_134_v59, 1)
		(_nw24).ArraySet1(_134_v59, 2)
		(_nw24).ArraySet1(_134_v59, 3)
		(_nw24).ArraySet1(_134_v59, 4)
		(_nw24).ArraySet1(_134_v59, 5)
		(_nw24).ArraySet1(_134_v59, 6)
		(_nw24).ArraySet1(_134_v59, 7)
		(_nw24).ArraySet1(_134_v59, 8)
		(_nw24).ArraySet1(_134_v59, 9)
		(_nw24).ArraySet1(_134_v59, 10)
		(_nw24).ArraySet1(_134_v59, 11)
		_135_v60 = _nw24
		var _136_v61 _dafny.Map
		_ = _136_v61
		_136_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_62_v1), _132_i1)
		var _137_v62 _dafny.Map
		_ = _137_v62
		_137_v62 = _136_v61
		_134_v59 = (Companion_D14_.Create_DC30_(_dafny.IntOfInt64(762), _135_v60, _137_v62, _dafny.IntOfInt64(-318), _134_v59)).Dtor_cf51()
		var _138_v63 *C3
		_ = _138_v63
		var _nw25 *C3 = New_C3_()
		_ = _nw25
		_nw25.Ctor__()
		_138_v63 = _nw25
		var _139_v64 D11
		_ = _139_v64
		_139_v64 = Companion_D11_.Create_DC24_(_138_v63)
		_139_v64 = _139_v64
		var _140_v65 bool
		_ = _140_v65
		var _141_v66 _dafny.Map
		_ = _141_v66
		var _out6 bool
		_ = _out6
		var _out7 _dafny.Map
		_ = _out7
		_out6, _out7 = (_138_v63).M0(_129_v55, _61_globalState)
		_140_v65 = _out6
		_141_v66 = _out7
	}
	var _142_v67 _dafny.MultiSet
	_ = _142_v67
	_142_v67 = _dafny.MultiSetOf(Companion_D10_.Create_DC22_())
	var _143_i2 _dafny.Int
	_ = _143_i2
	_143_i2 = _dafny.Zero
	{
		for (func() bool {
			if (_142_v67).IsSubsetOf(_142_v67) {
				return !(false)
			}
			return _130_v56
		})() {
			{
				if (_143_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_143_i2 = (_143_i2).Plus(_dafny.One)
				var _144_v68 *C2
				_ = _144_v68
				var _nw26 *C2 = New_C2_()
				_ = _nw26
				_nw26.Ctor__(_dafny.Companion_Sequence_.Concatenate(_124_v51, _124_v51), _124_v51)
				_144_v68 = _nw26
				_130_v56 = !(!(_62_v1) || (_62_v1))
				var _145_v69 _dafny.Array
				_ = _145_v69
				var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(2))
				_ = _nw27
				_145_v69 = _nw27
				_145_v69 = (func() _dafny.Array {
					if (Companion_Default___.Fm20(_130_v56, (_144_v68).F9(), _62_v1, _61_globalState)).IsSubsetOf(_125_v52) {
						return (func() _dafny.Array {
							if _130_v56 {
								return _145_v69
							}
							return _145_v69
						})()
					}
					return _145_v69
				})()
				_87_v26 = Companion_Default___.SafeDivisionInt(_87_v26, _dafny.IntOfInt64(15))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _146_v70 _dafny.Array
	_ = _146_v70
	var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
	_ = _nw28
	_146_v70 = _nw28
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_146_v70), 0))
	_ = _index10
	(_146_v70).ArraySet1(_dafny.IntOfInt64(72), (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_146_v70), 0))
	_ = _index11
	(_146_v70).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-592), _87_v26), (_index11).Int())
	var _147_v71 bool
	_ = _147_v71
	var _148_v72 _dafny.Map
	_ = _148_v72
	var _out8 bool
	_ = _out8
	var _out9 _dafny.Map
	_ = _out9
	_out8, _out9 = (_123_v50).M0(_129_v55, _61_globalState)
	_147_v71 = _out8
	_148_v72 = _out9
	_62_v1 = !(_62_v1)
	_87_v26 = _dafny.IntOfUint32(((func() _dafny.Sequence {
		if _130_v56 {
			return _dafny.Companion_Sequence_.Update(_124_v51, (Companion_Default___.SafeIndex((_146_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_146_v70), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_124_v51).Cardinality()))).Uint32(), _92_v31)
		}
		return _124_v51
	})()).Cardinality())
	_124_v51 = _124_v51
	_dafny.Print((_60_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_61_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_61_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_62_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_87_v26)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_v27).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(542), _dafny.IntOfInt64(542))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_89_v28, _dafny.SeqOf(_dafny.IntOfInt64(542))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_90_v29, _dafny.SeqOf(_dafny.IntOfInt64(542))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v30).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(542), _dafny.SeqOf(_dafny.IntOfInt64(542)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_92_v31)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v49).Equals(_dafny.MultiSetOf(false, false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_124_v51.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v52).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v53).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v53).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v53).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v53).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D1)).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D1)).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D1)).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_v54).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D1)).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_130_v56)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v57).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.MultiSetOf(false, false, false, false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_142_v67).Equals(_dafny.MultiSetOf(Companion_D10_.Create_DC22_())))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_143_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_v70).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_v71)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_v72).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.MultiSetOf(false, false, false, false, false))))
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
	Cf0 _dafny.Int
	Cf1 _dafny.Sequence
	Cf2 _dafny.Array
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int, Cf1 _dafny.Sequence, Cf2 _dafny.Array) D0 {
	return D0{D0_DC0{Cf0, Cf1, Cf2}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf3 D0
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf3 D0) D0 {
	return D0{D0_DC1{Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero, _dafny.EmptySeq, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf2
}

func (_this D0) Dtor_cf3() D0 {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + data.Cf1.VerbatimString(true) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1.Equals(data2.Cf1) && data1.Cf2 == data2.Cf2
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
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

type D1_DC3 struct {
	Cf5 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.Int) D1 {
	return D1{D1_DC3{Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf6 _dafny.Int
	Cf7 _dafny.CodePoint
	Cf8 bool
	Cf9 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 _dafny.Int, Cf7 _dafny.CodePoint, Cf8 bool, Cf9 _dafny.Int) D1 {
	return D1{D1_DC4{Cf6, Cf7, Cf8, Cf9}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC2 struct {
	Cf4 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 bool) D1 {
	return D1{D1_DC2{Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.Zero)
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf4 == data2.Cf4
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
	Cf11 _dafny.Int
	Cf12 bool
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 _dafny.Int, Cf12 bool) D2 {
	return D2{D2_DC6{Cf11, Cf12}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf10 _dafny.Sequence
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf10 _dafny.Sequence) D2 {
	return D2{D2_DC5{Cf10}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.Zero, false)
}

func (_this D2) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) Dtor_cf12() bool {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12 == data2.Cf12
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
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

type D3_DC7 struct {
	Cf13 _dafny.MultiSet
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf13 _dafny.MultiSet) D3 {
	return D3{D3_DC7{Cf13}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(_dafny.Zero)
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) Dtor_cf13() _dafny.MultiSet {
	return _this.Get_().(D3_DC7).Cf13
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf13) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf16 bool
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf16 bool) D4 {
	return D4{D4_DC10{Cf16}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

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

type D4_DC11 struct {
	Cf17 D4
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 D4) D4 {
	return D4{D4_DC11{Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(false)
}

func (_this D4) Dtor_cf16() bool {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) Dtor_cf17() D4 {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf17) + ")"
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
			return ok && data1.Cf16 == data2.Cf16
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf17.Equals(data2.Cf17)
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
	Cf18 _dafny.Map
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf18 _dafny.Map) D5 {
	return D5{D5_DC12{Cf18}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D5) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D5_DC12).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf19 _dafny.Array
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf19 _dafny.Array) D6 {
	return D6{D6_DC13{Cf19}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D6) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D6_DC13).Cf19
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf19 == data2.Cf19
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
	Cf21 bool
	Cf22 _dafny.Int
	Cf23 _dafny.Int
	Cf24 _dafny.Int
	Cf25 *C0
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf21 bool, Cf22 _dafny.Int, Cf23 _dafny.Int, Cf24 _dafny.Int, Cf25 *C0) D7 {
	return D7{D7_DC15{Cf21, Cf22, Cf23, Cf24, Cf25}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC14 struct {
	Cf20 _dafny.Sequence
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf20 _dafny.Sequence) D7 {
	return D7{D7_DC14{Cf20}}
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
	return Companion_D7_.Create_DC15_(false, _dafny.Zero, _dafny.Zero, _dafny.Zero, (*C0)(nil))
}

func (_this D7) Dtor_cf21() bool {
	return _this.Get_().(D7_DC15).Cf21
}

func (_this D7) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf22
}

func (_this D7) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf23
}

func (_this D7) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf24
}

func (_this D7) Dtor_cf25() *C0 {
	return _this.Get_().(D7_DC15).Cf25
}

func (_this D7) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D7_DC14).Cf20
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
			return "D7.DC15" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf20) + ")"
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
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25 == data2.Cf25
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf20.Equals(data2.Cf20)
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

type D8_DC17 struct {
	Cf27 _dafny.Map
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf27 _dafny.Map) D8 {
	return D8{D8_DC17{Cf27}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D8) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D8_DC17).Cf27
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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

type D9_DC19 struct {
	Cf29 _dafny.Int
	Cf30 _dafny.Int
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf29 _dafny.Int, Cf30 _dafny.Int) D9 {
	return D9{D9_DC19{Cf29, Cf30}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

type D9_DC20 struct {
	Cf31 bool
	Cf32 bool
	Cf33 bool
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf31 bool, Cf32 bool, Cf33 bool) D9 {
	return D9{D9_DC20{Cf31, Cf32, Cf33}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC18 struct {
	Cf28 _dafny.Map
}

func (D9_DC18) isD9() {}

func (CompanionStruct_D9_) Create_DC18_(Cf28 _dafny.Map) D9 {
	return D9{D9_DC18{Cf28}}
}

func (_this D9) Is_DC18() bool {
	_, ok := _this.Get_().(D9_DC18)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC19_(_dafny.Zero, _dafny.Zero)
}

func (_this D9) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D9_DC19).Cf29
}

func (_this D9) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D9_DC19).Cf30
}

func (_this D9) Dtor_cf31() bool {
	return _this.Get_().(D9_DC20).Cf31
}

func (_this D9) Dtor_cf32() bool {
	return _this.Get_().(D9_DC20).Cf32
}

func (_this D9) Dtor_cf33() bool {
	return _this.Get_().(D9_DC20).Cf33
}

func (_this D9) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D9_DC18).Cf28
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D9_DC18:
		{
			return "D9.DC18" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33
		}
	case D9_DC18:
		{
			data2, ok := other.Get_().(D9_DC18)
			return ok && data1.Cf28.Equals(data2.Cf28)
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
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_() D10 {
	return D10{D10_DC22{}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

type D10_DC23 struct {
	Cf35 _dafny.MultiSet
	Cf36 _dafny.Int
	Cf37 bool
	Cf38 bool
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf35 _dafny.MultiSet, Cf36 _dafny.Int, Cf37 bool, Cf38 bool) D10 {
	return D10{D10_DC23{Cf35, Cf36, Cf37, Cf38}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

type D10_DC21 struct {
	Cf34 _dafny.MultiSet
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf34 _dafny.MultiSet) D10 {
	return D10{D10_DC21{Cf34}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC22_()
}

func (_this D10) Dtor_cf35() _dafny.MultiSet {
	return _this.Get_().(D10_DC23).Cf35
}

func (_this D10) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D10_DC23).Cf36
}

func (_this D10) Dtor_cf37() bool {
	return _this.Get_().(D10_DC23).Cf37
}

func (_this D10) Dtor_cf38() bool {
	return _this.Get_().(D10_DC23).Cf38
}

func (_this D10) Dtor_cf34() _dafny.MultiSet {
	return _this.Get_().(D10_DC21).Cf34
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC22:
		{
			return "D10.DC22"
		}
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf34) + ")"
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
			_, ok := other.Get_().(D10_DC22)
			return ok
		}
	case D10_DC23:
		{
			data2, ok := other.Get_().(D10_DC23)
			return ok && data1.Cf35.Equals(data2.Cf35) && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38
		}
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf34.Equals(data2.Cf34)
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
	Cf40 bool
	Cf41 D9
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf40 bool, Cf41 D9) D11 {
	return D11{D11_DC25{Cf40, Cf41}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC26 struct {
	Cf42 _dafny.Sequence
	Cf43 T1
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf42 _dafny.Sequence, Cf43 T1) D11 {
	return D11{D11_DC26{Cf42, Cf43}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

type D11_DC24 struct {
	Cf39 *C3
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf39 *C3) D11 {
	return D11{D11_DC24{Cf39}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC25_(false, Companion_D9_.Default())
}

func (_this D11) Dtor_cf40() bool {
	return _this.Get_().(D11_DC25).Cf40
}

func (_this D11) Dtor_cf41() D9 {
	return _this.Get_().(D11_DC25).Cf41
}

func (_this D11) Dtor_cf42() _dafny.Sequence {
	return _this.Get_().(D11_DC26).Cf42
}

func (_this D11) Dtor_cf43() T1 {
	return _this.Get_().(D11_DC26).Cf43
}

func (_this D11) Dtor_cf39() *C3 {
	return _this.Get_().(D11_DC24).Cf39
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ")"
		}
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf40 == data2.Cf40 && data1.Cf41.Equals(data2.Cf41)
		}
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && data1.Cf42.Equals(data2.Cf42) && _dafny.AreEqual(data1.Cf43, data2.Cf43)
		}
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf39 == data2.Cf39
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
	Cf44 _dafny.Map
}

func (D12_DC27) isD12() {}

func (CompanionStruct_D12_) Create_DC27_(Cf44 _dafny.Map) D12 {
	return D12{D12_DC27{Cf44}}
}

func (_this D12) Is_DC27() bool {
	_, ok := _this.Get_().(D12_DC27)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D12) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D12_DC27).Cf44
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC27:
		{
			return "D12.DC27" + "(" + _dafny.String(data.Cf44) + ")"
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

type D13_DC28 struct {
	Cf45 _dafny.Sequence
}

func (D13_DC28) isD13() {}

func (CompanionStruct_D13_) Create_DC28_(Cf45 _dafny.Sequence) D13 {
	return D13{D13_DC28{Cf45}}
}

func (_this D13) Is_DC28() bool {
	_, ok := _this.Get_().(D13_DC28)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D13) Dtor_cf45() _dafny.Sequence {
	return _this.Get_().(D13_DC28).Cf45
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC28:
		{
			return "D13.DC28" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC28:
		{
			data2, ok := other.Get_().(D13_DC28)
			return ok && data1.Cf45.Equals(data2.Cf45)
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

type D14_DC30 struct {
	Cf47 _dafny.Int
	Cf48 _dafny.Array
	Cf49 _dafny.Map
	Cf50 _dafny.Int
	Cf51 _dafny.Array
}

func (D14_DC30) isD14() {}

func (CompanionStruct_D14_) Create_DC30_(Cf47 _dafny.Int, Cf48 _dafny.Array, Cf49 _dafny.Map, Cf50 _dafny.Int, Cf51 _dafny.Array) D14 {
	return D14{D14_DC30{Cf47, Cf48, Cf49, Cf50, Cf51}}
}

func (_this D14) Is_DC30() bool {
	_, ok := _this.Get_().(D14_DC30)
	return ok
}

type D14_DC29 struct {
	Cf46 _dafny.Sequence
}

func (D14_DC29) isD14() {}

func (CompanionStruct_D14_) Create_DC29_(Cf46 _dafny.Sequence) D14 {
	return D14{D14_DC29{Cf46}}
}

func (_this D14) Is_DC29() bool {
	_, ok := _this.Get_().(D14_DC29)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC30_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMap, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D14) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D14_DC30).Cf47
}

func (_this D14) Dtor_cf48() _dafny.Array {
	return _this.Get_().(D14_DC30).Cf48
}

func (_this D14) Dtor_cf49() _dafny.Map {
	return _this.Get_().(D14_DC30).Cf49
}

func (_this D14) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D14_DC30).Cf50
}

func (_this D14) Dtor_cf51() _dafny.Array {
	return _this.Get_().(D14_DC30).Cf51
}

func (_this D14) Dtor_cf46() _dafny.Sequence {
	return _this.Get_().(D14_DC29).Cf46
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC30:
		{
			return "D14.DC30" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ")"
		}
	case D14_DC29:
		{
			return "D14.DC29" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC30:
		{
			data2, ok := other.Get_().(D14_DC30)
			return ok && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48 == data2.Cf48 && data1.Cf49.Equals(data2.Cf49) && data1.Cf50.Cmp(data2.Cf50) == 0 && data1.Cf51 == data2.Cf51
		}
	case D14_DC29:
		{
			data2, ok := other.Get_().(D14_DC29)
			return ok && data1.Cf46.Equals(data2.Cf46)
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

type D15_DC32 struct {
	Cf53 bool
}

func (D15_DC32) isD15() {}

func (CompanionStruct_D15_) Create_DC32_(Cf53 bool) D15 {
	return D15{D15_DC32{Cf53}}
}

func (_this D15) Is_DC32() bool {
	_, ok := _this.Get_().(D15_DC32)
	return ok
}

type D15_DC33 struct {
	Cf54 _dafny.Sequence
}

func (D15_DC33) isD15() {}

func (CompanionStruct_D15_) Create_DC33_(Cf54 _dafny.Sequence) D15 {
	return D15{D15_DC33{Cf54}}
}

func (_this D15) Is_DC33() bool {
	_, ok := _this.Get_().(D15_DC33)
	return ok
}

type D15_DC34 struct {
	Cf55 _dafny.Sequence
	Cf56 _dafny.Int
	Cf57 bool
	Cf58 D10
	Cf59 _dafny.Int
}

func (D15_DC34) isD15() {}

func (CompanionStruct_D15_) Create_DC34_(Cf55 _dafny.Sequence, Cf56 _dafny.Int, Cf57 bool, Cf58 D10, Cf59 _dafny.Int) D15 {
	return D15{D15_DC34{Cf55, Cf56, Cf57, Cf58, Cf59}}
}

func (_this D15) Is_DC34() bool {
	_, ok := _this.Get_().(D15_DC34)
	return ok
}

type D15_DC31 struct {
	Cf52 _dafny.Map
}

func (D15_DC31) isD15() {}

func (CompanionStruct_D15_) Create_DC31_(Cf52 _dafny.Map) D15 {
	return D15{D15_DC31{Cf52}}
}

func (_this D15) Is_DC31() bool {
	_, ok := _this.Get_().(D15_DC31)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC32_(false)
}

func (_this D15) Dtor_cf53() bool {
	return _this.Get_().(D15_DC32).Cf53
}

func (_this D15) Dtor_cf54() _dafny.Sequence {
	return _this.Get_().(D15_DC33).Cf54
}

func (_this D15) Dtor_cf55() _dafny.Sequence {
	return _this.Get_().(D15_DC34).Cf55
}

func (_this D15) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D15_DC34).Cf56
}

func (_this D15) Dtor_cf57() bool {
	return _this.Get_().(D15_DC34).Cf57
}

func (_this D15) Dtor_cf58() D10 {
	return _this.Get_().(D15_DC34).Cf58
}

func (_this D15) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D15_DC34).Cf59
}

func (_this D15) Dtor_cf52() _dafny.Map {
	return _this.Get_().(D15_DC31).Cf52
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC32:
		{
			return "D15.DC32" + "(" + _dafny.String(data.Cf53) + ")"
		}
	case D15_DC33:
		{
			return "D15.DC33" + "(" + data.Cf54.VerbatimString(true) + ")"
		}
	case D15_DC34:
		{
			return "D15.DC34" + "(" + data.Cf55.VerbatimString(true) + ", " + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D15_DC31:
		{
			return "D15.DC31" + "(" + _dafny.String(data.Cf52) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC32:
		{
			data2, ok := other.Get_().(D15_DC32)
			return ok && data1.Cf53 == data2.Cf53
		}
	case D15_DC33:
		{
			data2, ok := other.Get_().(D15_DC33)
			return ok && data1.Cf54.Equals(data2.Cf54)
		}
	case D15_DC34:
		{
			data2, ok := other.Get_().(D15_DC34)
			return ok && data1.Cf55.Equals(data2.Cf55) && data1.Cf56.Cmp(data2.Cf56) == 0 && data1.Cf57 == data2.Cf57 && data1.Cf58.Equals(data2.Cf58) && data1.Cf59.Cmp(data2.Cf59) == 0
		}
	case D15_DC31:
		{
			data2, ok := other.Get_().(D15_DC31)
			return ok && data1.Cf52.Equals(data2.Cf52)
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
	Cf60 _dafny.Set
}

func (D16_DC35) isD16() {}

func (CompanionStruct_D16_) Create_DC35_(Cf60 _dafny.Set) D16 {
	return D16{D16_DC35{Cf60}}
}

func (_this D16) Is_DC35() bool {
	_, ok := _this.Get_().(D16_DC35)
	return ok
}

func (CompanionStruct_D16_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D16) Dtor_cf60() _dafny.Set {
	return _this.Get_().(D16_DC35).Cf60
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC35:
		{
			return "D16.DC35" + "(" + _dafny.String(data.Cf60) + ")"
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
			return ok && data1.Cf60.Equals(data2.Cf60)
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
	Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet
	M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map)
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
	Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet
	M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map)
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
	_f0 _dafny.Array
	_f1 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
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

// End of class C0

// Definition of class C1
type C1 struct {
	F11  _dafny.Int
	_f10 _dafny.MultiSet
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F11 = _dafny.Zero
	_this._f10 = _dafny.EmptyMultiSet
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

func (_this *C1) Ctor__(f10 _dafny.MultiSet, f11 _dafny.Int) {
	{
		(_this)._f10 = f10
		(_this).F11 = f11
	}
}
func (_this *C1) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_this).F10()
	}
}
func (_this *C1) Fm21(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(!((func() bool {
			if !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("talxrgws")).Cardinality())), _dafny.SeqOf(_this.F11)) {
				return (_dafny.SetOf(!(true))).IsSubsetOf(_dafny.SetOf(false, true))
			}
			return (_this.F11).Cmp(_this.F11) == 0
		})()))
	}
}
func (_this *C1) M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _149_v0 bool
		_ = _149_v0
		_149_v0 = true
		var _150_v1 D1
		_ = _150_v1
		_150_v1 = Companion_D1_.Create_DC2_(!(_149_v0))
		var _151_v2 _dafny.Sequence
		_ = _151_v2
		_151_v2 = _dafny.SeqOf(false)
		var _152_v3 D4
		_ = _152_v3
		_152_v3 = Companion_D4_.Create_DC10_(!(false))
		var _153_v4 _dafny.CodePoint
		_ = _153_v4
		_153_v4 = _dafny.CodePoint('i')
		var _154_v5 _dafny.Array
		_ = _154_v5
		var _nwElement0_6 bool = _149_v0
		_ = _nwElement0_6
		var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(8))
		_ = _nw29
		(_nw29).ArraySet1(_nwElement0_6, 0)
		(_nw29).ArraySet1((_dafny.IntOfInt64(436)).Cmp(_this.F11) == 0, 1)
		(_nw29).ArraySet1(((_150_v1).Dtor_cf4()) || (_149_v0), 2)
		(_nw29).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_151_v2, (Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((_151_v2).Cardinality()))).Uint32(), _149_v0)).Cardinality())).Cmp(_this.F11) > 0, 3)
		(_nw29).ArraySet1(!((_152_v3).Dtor_cf16()), 4)
		(_nw29).ArraySet1(_149_v0, 5)
		(_nw29).ArraySet1(!(_149_v0) || (_149_v0), 6)
		(_nw29).ArraySet1((_this).Fm21(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v4, _dafny.IntOfInt64(179)), _this.F11, _this.F11, globalState), 7)
		_154_v5 = _nw29
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
		_ = _index12
		(_154_v5).ArraySet1(_149_v0, (_index12).Int())
		var _155_v6 _dafny.Map
		_ = _155_v6
		_155_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v4, _this.F11)
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
		_ = _index13
		(_154_v5).ArraySet1((_this).Fm21((_155_v6).Update(_153_v4, _this.F11), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(567), _this.F11)), _this.F11, globalState), (_index13).Int())
		var _156_v7 _dafny.Array
		_ = _156_v7
		var _nwElement0_7 _dafny.Int = _this.F11
		_ = _nwElement0_7
		var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(2))
		_ = _nw30
		(_nw30).ArraySet1(_nwElement0_7, 0)
		(_nw30).ArraySet1((_dafny.Zero).Minus(_this.F11), 1)
		_156_v7 = _nw30
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
		_ = _index14
		(_156_v7).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm22((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), _dafny.IntOfInt64(-599), globalState)).Cardinality()), (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
		_ = _index15
		(_156_v7).ArraySet1((_dafny.Zero).Minus(_this.F11), (_index15).Int())
		var _157_v8 _dafny.Map
		_ = _157_v8
		_157_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), (_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int))
		var _hi1 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf((_151_v2).Select((Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((_151_v2).Cardinality()))).Uint32()).(bool))).Cardinality()), _dafny.IntOfInt64(441))
		_ = _hi1
		for _158_i0 := (_157_v8).Cardinality(); _158_i0.Cmp(_hi1) < 0; _158_i0 = _158_i0.Plus(_dafny.One) {
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
			_ = _index16
			(_156_v7).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm23(_158_i0, globalState)).Cardinality()), (_index16).Int())
			var _159_v9 _dafny.Sequence
			_ = _159_v9
			_159_v9 = _dafny.SeqOf((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int))
			var _160_v10 _dafny.Sequence
			_ = _160_v10
			_160_v10 = _dafny.SeqOf(_this.F11, _dafny.IntOfUint32((_159_v9).Cardinality()), (_dafny.MultiSetOf(!(!(_149_v0)), _149_v0, (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool))).Cardinality())
			var _161_v11 D2
			_ = _161_v11
			_161_v11 = Companion_D2_.Create_DC5_(_159_v9)
			_161_v11 = Companion_D2_.Create_DC5_(_160_v10)
			var _162_v12 _dafny.Sequence
			_ = _162_v12
			_162_v12 = _dafny.UnicodeSeqOfUtf8Bytes("cltsjvhb")
			var _163_v13 _dafny.Array
			_ = _163_v13
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw31
			_163_v13 = _nw31
			var _164_v14 D0
			_ = _164_v14
			_164_v14 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(411), _162_v12, _163_v13)
			var _165_v15 _dafny.Map
			_ = _165_v15
			_165_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v5, _164_v14)
			var _166_v16 _dafny.Sequence
			_ = _166_v16
			_166_v16 = _dafny.SeqOf(_164_v14)
			var _167_v17 _dafny.Set
			_ = _167_v17
			_167_v17 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_151_v2)).Cardinality())))
			var _168_v18 D0
			_ = _168_v18
			_168_v18 = Companion_D0_.Create_DC1_((func() D0 {
				if (_165_v15).Contains(_154_v5) {
					return (_165_v15).Get(_154_v5).(D0)
				}
				return (_166_v16).Select((Companion_Default___.SafeIndex((_167_v17).Cardinality(), _dafny.IntOfUint32((_166_v16).Cardinality()))).Uint32()).(D0)
			})())
			var _source4 D0 = _168_v18
			_ = _source4
			if _source4.Is_DC0() {
				var _169___mcc_h0 _dafny.Int = _source4.Get_().(D0_DC0).Cf0
				_ = _169___mcc_h0
				var _170___mcc_h1 _dafny.Sequence = _source4.Get_().(D0_DC0).Cf1
				_ = _170___mcc_h1
				var _171___mcc_h2 _dafny.Array = _source4.Get_().(D0_DC0).Cf2
				_ = _171___mcc_h2
				var _172_cf2 _dafny.Array = _171___mcc_h2
				_ = _172_cf2
				var _173_cf1 _dafny.Sequence = _170___mcc_h1
				_ = _173_cf1
				var _174_cf0 _dafny.Int = _169___mcc_h0
				_ = _174_cf0
				var _175_v19 *C0
				_ = _175_v19
				var _nw32 *C0 = New_C0_()
				_ = _nw32
				_nw32.Ctor__()
				_175_v19 = _nw32
				var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
				_ = _index17
				(_154_v5).ArraySet1(!(_dafny.Companion_Sequence_.Equal(_160_v10, _159_v9)) || (false), (_index17).Int())
				var _176_v20 *C0
				_ = _176_v20
				var _nw33 *C0 = New_C0_()
				_ = _nw33
				_nw33.Ctor__()
				_176_v20 = _nw33
				var _177_v21 _dafny.MultiSet
				_ = _177_v21
				_177_v21 = _dafny.MultiSetOf(_this.F11)
				var _178_v22 _dafny.Map
				_ = _178_v22
				_178_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_149_v0) == (!((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool))), (_177_v21).Contains(_this.F11))
				_178_v22 = (_178_v22).Update((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), (_150_v1).Dtor_cf4())
			} else {
				var _179___mcc_h3 D0 = _source4.Get_().(D0_DC1).Cf3
				_ = _179___mcc_h3
				var _180_cf3 D0 = _179___mcc_h3
				_ = _180_cf3
				var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
				_ = _index18
				var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
				_ = _index19
				var _rhs15 _dafny.Int = _this.F11
				_ = _rhs15
				var _rhs16 _dafny.Int = _this.F11
				_ = _rhs16
				var _lhs2 _dafny.Array = _156_v7
				_ = _lhs2
				var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
				_ = _lhs3
				var _lhs4 _dafny.Array = _156_v7
				_ = _lhs4
				var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
				_ = _lhs5
				(_lhs2).ArraySet1(_rhs15, (_lhs3).Int())
				(_lhs4).ArraySet1(_rhs16, (_lhs5).Int())
				var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
				_ = _index20
				(_154_v5).ArraySet1((_this).Fm21(_155_v6, ((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)).Plus(_this.F11), Companion_Default___.Fm24(_this.F11, (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), _149_v0, _dafny.IntOfUint32((_159_v9).Cardinality()), globalState), globalState), (_index20).Int())
				_153_v4 = _153_v4
				_153_v4 = _153_v4
			}
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
			_ = _index21
			(_154_v5).ArraySet1(_149_v0, (_index21).Int())
		}
		var _181_v23 _dafny.Sequence
		_ = _181_v23
		_181_v23 = _dafny.UnicodeSeqOfUtf8Bytes("vtlum")
		var _hi2 _dafny.Int = _this.F11
		_ = _hi2
		for _182_i1 := Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_181_v23).Cardinality()), _this.F11); _182_i1.Cmp(_hi2) < 0; _182_i1 = _182_i1.Plus(_dafny.One) {
			var _183_v24 _dafny.Sequence
			_ = _183_v24
			_183_v24 = _dafny.SeqOf(_182_i1, _dafny.IntOfInt64(712))
			_149_v0 = (((_this).F10()).Update(_149_v0, Companion_Default___.Abs(_dafny.IntOfUint32((_183_v24).Cardinality())))).IsSubsetOf((_this).F10())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
			_ = _index22
			var _rhs17 bool = _149_v0
			_ = _rhs17
			var _rhs18 _dafny.Int = (_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)
			_ = _rhs18
			var _rhs19 _dafny.Int = (_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)
			_ = _rhs19
			var _lhs6 *C1 = _this
			_ = _lhs6
			var _lhs7 _dafny.Array = _156_v7
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
			_ = _lhs8
			r0 = _rhs17
			_lhs6.F11 = _rhs18
			(_lhs7).ArraySet1(_rhs19, (_lhs8).Int())
			var _184_v25 _dafny.Array
			_ = _184_v25
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(23))
			_ = _nw34
			_184_v25 = _nw34
			var _185_v26 D2
			_ = _185_v26
			_185_v26 = Companion_D2_.Create_DC6_((_dafny.MultiSetFromSeq(_183_v24)).Cardinality(), (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool))
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_184_v25), 0))
			_ = _index23
			(_184_v25).ArraySet1(_185_v26, (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_184_v25), 0))
			_ = _index24
			(_184_v25).ArraySet1(_185_v26, (_index24).Int())
			var _186_v27 _dafny.Array
			_ = _186_v27
			var _nwElement0_8 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-530))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_187_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_188_i2 _dafny.Int) _dafny.Sequence {
					return _187_v23
				}
			})(_181_v23)))
			_ = _nwElement0_8
			var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.One)
			_ = _nw35
			(_nw35).ArraySet1(_nwElement0_8, 0)
			_186_v27 = _nw35
			var _189_v28 _dafny.Sequence
			_ = _189_v28
			_189_v28 = _dafny.SeqOf(_181_v23)
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_186_v27), 0))
			_ = _index25
			(_186_v27).ArraySet1(_189_v28, (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_186_v27), 0))
			_ = _index26
			(_186_v27).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_189_v28, Companion_Default___.Fm25(_181_v23, Companion_Default___.Fm24((_dafny.Zero).Minus(_182_i1), (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), _149_v0, (_dafny.MultiSetFromSeq(_183_v24)).Cardinality(), globalState), _182_i1, globalState)), (_index26).Int())
		}
		if true {
			_181_v23 = _181_v23
			if (_149_v0) && (((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool)) && ((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool))) {
				_149_v0 = _149_v0
				var _190_v29 _dafny.Array
				_ = _190_v29
				var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw36
				_190_v29 = _nw36
				_190_v29 = p0
				var _191_v30 D1
				_ = _191_v30
				_191_v30 = Companion_D1_.Create_DC3_(_this.F11)
				var _192_v31 _dafny.Sequence
				_ = _192_v31
				_192_v31 = _dafny.SeqOf((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int))
				var _193_v32 _dafny.Map
				_ = _193_v32
				_193_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool) {
						return _dafny.SeqOf((_191_v30).Dtor_cf5(), _this.F11)
					}
					return _192_v31
				})(), (false) || ((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool)))
				_193_v32 = _193_v32
				var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
				_ = _index27
				(_156_v7).ArraySet1(_this.F11, (_index27).Int())
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
				_ = _index28
				(_154_v5).ArraySet1((false) || (_dafny.Companion_Sequence_.Equal(_192_v31, _dafny.SeqOf(_this.F11))), (_index28).Int())
			} else {
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
				_ = _index29
				(_156_v7).ArraySet1((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), (_index29).Int())
				_149_v0 = (_151_v2).Select((Companion_Default___.SafeIndex((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_151_v2).Cardinality()))).Uint32()).(bool)
				_150_v1 = _150_v1
				var _194_v33 _dafny.Sequence
				_ = _194_v33
				_194_v33 = _dafny.SeqOf(_this.F11, _dafny.IntOfUint32((_181_v23).Cardinality()), (_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int))
				var _195_v34 _dafny.Map
				_ = _195_v34
				_195_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v0, ((_dafny.MultiSetOf((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(274))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}((func(_196_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_197_i3 _dafny.Int) _dafny.CodePoint {
						return _196_v4
					}
				})(_153_v4)))).Cardinality()))).Update(_this.F11, Companion_Default___.Abs((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)))).Difference(_dafny.MultiSetFromSeq(_194_v33)))
				var _198_v36 _dafny.Map
				_ = _198_v36
				_198_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _153_v4)
				var _199_v37 _dafny.MultiSet
				_ = _199_v37
				_199_v37 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(304))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}((func(_200_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_201_i4 _dafny.Int) _dafny.CodePoint {
						return _200_v4
					}
				})(_153_v4))), (Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(304))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg13 _dafny.Int) interface{} {
						return coer13(arg13)
					}
				}((func(_202_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_203_i4 _dafny.Int) _dafny.CodePoint {
						return _202_v4
					}
				})(_153_v4)))).Cardinality()))).Uint32(), _153_v4)).Cardinality()))
				_195_v34 = (_195_v34).Update(!((_this).Fm21(func() _dafny.Map {
					var _coll8 = _dafny.NewMapBuilder()
					_ = _coll8
					for _iter9 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
						if (_198_v36).Contains(_this.F11) {
							return (_198_v36).Get(_this.F11).(_dafny.CodePoint)
						}
						return _153_v4
					})(), false)).Keys().Elements()); ; {
						_compr_8, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _204_v35 _dafny.CodePoint
						_204_v35 = interface{}(_compr_8).(_dafny.CodePoint)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
							if (_198_v36).Contains(_this.F11) {
								return (_198_v36).Get(_this.F11).(_dafny.CodePoint)
							}
							return _153_v4
						})(), false)).Contains(_204_v35) {
							_coll8.Add(_204_v35, _this.F11)
						}
					}
					return _coll8.ToMap()
				}(), (_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm26((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), globalState)).Cardinality()), globalState)) || (_149_v0), _199_v37)
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
				_ = _index30
				(_154_v5).ArraySet1((true) || ((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool)), (_index30).Int())
			}
			var _205_v38 _dafny.Sequence
			_ = _205_v38
			_205_v38 = _dafny.SeqOf(_this.F11)
			var _206_v39 D2
			_ = _206_v39
			_206_v39 = Companion_D2_.Create_DC5_(_205_v38)
			_206_v39 = _206_v39
			var _207_v40 *C0
			_ = _207_v40
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__()
			_207_v40 = _nw37
			var _208_v41 _dafny.Array
			_ = _208_v41
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(8))
			_ = _nw38
			_208_v41 = _nw38
			var _209_v42 _dafny.Array
			_ = _209_v42
			var _nw39 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
			_ = _nw39
			_209_v42 = _nw39
			var _210_v43 _dafny.Set
			_ = _210_v43
			_210_v43 = _dafny.SetOf(_209_v42)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_208_v41), 0))
			_ = _index31
			(_208_v41).ArraySet1(_210_v43, (_index31).Int())
			var _211_v44 _dafny.Sequence
			_ = _211_v44
			_211_v44 = _dafny.SeqOf(_209_v42, _209_v42, _209_v42, _209_v42)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_208_v41), 0))
			_ = _index32
			(_208_v41).ArraySet1((_210_v43).Difference(_dafny.SetOf((_211_v44).Select((Companion_Default___.SafeIndex((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_211_v44).Cardinality()))).Uint32()).(_dafny.Array), _209_v42)), (_index32).Int())
		} else {
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
			_ = _index33
			(_156_v7).ArraySet1((_dafny.Zero).Minus(_this.F11), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
			_ = _index34
			(_154_v5).ArraySet1(_149_v0, (_index34).Int())
			_149_v0 = !((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool))
			var _212_v45 *C0
			_ = _212_v45
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__()
			_212_v45 = _nw40
			(_this).F11 = ((_this).F10()).Cardinality()
		}
		if (_151_v2).Select((Companion_Default___.SafeIndex(((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)).Minus(_this.F11), _dafny.IntOfUint32((_151_v2).Cardinality()))).Uint32()).(bool) {
			var _213_v46 _dafny.MultiSet
			_ = _213_v46
			_213_v46 = _dafny.MultiSetOf(_153_v4, _153_v4, _153_v4)
			_213_v46 = _213_v46
			var _214_v47 _dafny.Sequence
			_ = _214_v47
			_214_v47 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("dqbpjl"))
			var _215_v48 _dafny.Set
			_ = _215_v48
			_215_v48 = _dafny.SetOf(_153_v4, _153_v4, _153_v4, _153_v4, _153_v4)
			var _216_v49 _dafny.MultiSet
			_ = _216_v49
			_216_v49 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("a"), _181_v23, _181_v23, _dafny.UnicodeSeqOfUtf8Bytes("rmkxsg"), _181_v23)
			var _217_v50 _dafny.Sequence
			_ = _217_v50
			_217_v50 = _dafny.SeqOf(_214_v47, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-244))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_218_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_219_i5 _dafny.Int) _dafny.Sequence {
					return _218_v23
				}
			})(_181_v23))))
			var _220_v51 _dafny.Array
			_ = _220_v51
			var _nwElement0_9 _dafny.Array = _156_v7
			_ = _nwElement0_9
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(28))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_9, 0)
			(_nw41).ArraySet1(_156_v7, 1)
			(_nw41).ArraySet1(_156_v7, 2)
			(_nw41).ArraySet1(_156_v7, 3)
			(_nw41).ArraySet1(_156_v7, 4)
			(_nw41).ArraySet1(_156_v7, 5)
			(_nw41).ArraySet1(_156_v7, 6)
			(_nw41).ArraySet1(_156_v7, 7)
			(_nw41).ArraySet1(_156_v7, 8)
			(_nw41).ArraySet1(_156_v7, 9)
			(_nw41).ArraySet1(_156_v7, 10)
			(_nw41).ArraySet1(_156_v7, 11)
			(_nw41).ArraySet1(_156_v7, 12)
			(_nw41).ArraySet1(_156_v7, 13)
			(_nw41).ArraySet1(_156_v7, 14)
			(_nw41).ArraySet1(_156_v7, 15)
			(_nw41).ArraySet1(_156_v7, 16)
			(_nw41).ArraySet1(_156_v7, 17)
			(_nw41).ArraySet1(_156_v7, 18)
			(_nw41).ArraySet1(_156_v7, 19)
			(_nw41).ArraySet1(_156_v7, 20)
			(_nw41).ArraySet1(_156_v7, 21)
			(_nw41).ArraySet1(_156_v7, 22)
			(_nw41).ArraySet1(_156_v7, 23)
			(_nw41).ArraySet1(_156_v7, 24)
			(_nw41).ArraySet1(_156_v7, 25)
			(_nw41).ArraySet1(_156_v7, 26)
			(_nw41).ArraySet1(_156_v7, 27)
			_220_v51 = _nw41
			var _221_v52 D0
			_ = _221_v52
			_221_v52 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(592), _dafny.UnicodeSeqOfUtf8Bytes("ajxyfq"), _220_v51)
			var _222_v53 _dafny.Array
			_ = _222_v53
			var _nwElement0_10 _dafny.MultiSet = _dafny.MultiSetFromSeq(_214_v47)
			_ = _nwElement0_10
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(28))
			_ = _nw42
			(_nw42).ArraySet1(_nwElement0_10, 0)
			(_nw42).ArraySet1((Companion_Default___.Fm27(globalState)).Update(_181_v23, Companion_Default___.Abs((_215_v48).Cardinality())), 1)
			(_nw42).ArraySet1(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool) {
					return _214_v47
				}
				return _214_v47
			})()), 2)
			(_nw42).ArraySet1(_216_v49, 3)
			(_nw42).ArraySet1(_216_v49, 4)
			(_nw42).ArraySet1(_216_v49, 5)
			(_nw42).ArraySet1(_216_v49, 6)
			(_nw42).ArraySet1(_216_v49, 7)
			(_nw42).ArraySet1(_dafny.MultiSetOf(_181_v23, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm22((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), _this.F11, globalState), (Companion_Default___.SafeIndex(((_this).F10()).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm22((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), _this.F11, globalState)).Cardinality()))).Uint32(), _153_v4), _181_v23, _181_v23, _181_v23), 8)
			(_nw42).ArraySet1(_dafny.MultiSetFromSeq((_217_v50).Select((Companion_Default___.SafeIndex((_221_v52).Dtor_cf0(), _dafny.IntOfUint32((_217_v50).Cardinality()))).Uint32()).(_dafny.Sequence)), 9)
			(_nw42).ArraySet1((_dafny.MultiSetOf(_181_v23, _181_v23)).Difference(_216_v49), 10)
			(_nw42).ArraySet1(_216_v49, 11)
			(_nw42).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_214_v47, _214_v47)), 12)
			(_nw42).ArraySet1(_216_v49, 13)
			(_nw42).ArraySet1((func() _dafny.MultiSet {
				if (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool) {
					return _216_v49
				}
				return _216_v49
			})(), 14)
			(_nw42).ArraySet1(_216_v49, 15)
			(_nw42).ArraySet1(_216_v49, 16)
			(_nw42).ArraySet1(_216_v49, 17)
			(_nw42).ArraySet1(Companion_Default___.Fm27(globalState), 18)
			(_nw42).ArraySet1((_216_v49).Union(_216_v49), 19)
			(_nw42).ArraySet1(_216_v49, 20)
			(_nw42).ArraySet1(_216_v49, 21)
			(_nw42).ArraySet1(_216_v49, 22)
			(_nw42).ArraySet1(_216_v49, 23)
			(_nw42).ArraySet1(_216_v49, 24)
			(_nw42).ArraySet1((_dafny.MultiSetFromSeq(_214_v47)).Difference(_dafny.MultiSetOf(_181_v23)), 25)
			(_nw42).ArraySet1(_dafny.MultiSetOf(_181_v23), 26)
			(_nw42).ArraySet1(_216_v49, 27)
			_222_v53 = _nw42
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_222_v53), 0))
			_ = _index35
			(_222_v53).ArraySet1(_216_v49, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_222_v53), 0))
			_ = _index36
			(_222_v53).ArraySet1(_dafny.MultiSetOf(_181_v23, _dafny.Companion_Sequence_.Concatenate(_181_v23, _dafny.UnicodeSeqOfUtf8Bytes("xce")), _181_v23), (_index36).Int())
			if (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool) {
				(_this).F11 = (_dafny.Zero).Minus((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int))
				_149_v0 = (func() bool {
					if _149_v0 {
						return _149_v0
					}
					return !((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool)) || (_149_v0)
				})()
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
				_ = _index37
				(_156_v7).ArraySet1(_this.F11, (_index37).Int())
				_149_v0 = !_dafny.Companion_Sequence_.Contains(_181_v23, _153_v4)
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
				_ = _index38
				(_154_v5).ArraySet1(_149_v0, (_index38).Int())
			} else {
				var _223_v54 _dafny.Array
				_ = _223_v54
				var _len0_1 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_1
				var _nw43 _dafny.Array
				_ = _nw43
				if _len0_1.Cmp(_dafny.Zero) == 0 {
					_nw43 = _dafny.NewArray(_len0_1)
				} else {
					var _init1 func(_dafny.Int) _dafny.Sequence = (func(_224_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_225_i6 _dafny.Int) _dafny.Sequence {
							return _224_v23
						}
					})(_181_v23)
					_ = _init1
					var _element0_1 = _init1(_dafny.Zero)
					_ = _element0_1
					_nw43 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
					(_nw43).ArraySet1(_element0_1, 0)
					var _nativeLen0_1 = (_len0_1).Int()
					_ = _nativeLen0_1
					for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
						(_nw43).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
					}
				}
				_223_v54 = _nw43
				_223_v54 = _223_v54
				var _226_v55 _dafny.Map
				_ = _226_v55
				_226_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, (func() _dafny.Int {
					if (_213_v46).Contains(_153_v4) {
						return (_213_v46).Multiplicity(_153_v4)
					}
					return (_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)
				})())
				(_this).F11 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-943), (func() _dafny.Int {
					if (_226_v55).Contains(_this.F11) {
						return (_226_v55).Get(_this.F11).(_dafny.Int)
					}
					return _this.F11
				})()))
				r0 = (_this).Fm21(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v4, (_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)), _this.F11, _this.F11, globalState)
				var _227_v56 _dafny.Set
				_ = _227_v56
				_227_v56 = _dafny.SetOf(false)
				var _228_v57 D1
				_ = _228_v57
				_228_v57 = Companion_D1_.Create_DC4_((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), _153_v4, _149_v0, _this.F11)
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))
				_ = _index39
				(_156_v7).ArraySet1((((_227_v56).Cardinality()).Minus((_228_v57).Dtor_cf9())).Minus(_this.F11), (_index39).Int())
				_214_v47 = _dafny.Companion_Sequence_.Concatenate(_214_v47, _214_v47)
			}
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))
			_ = _index40
			(_154_v5).ArraySet1((_dafny.One).Cmp(((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(-273))) != 0, (_index40).Int())
			var _229_v58 _dafny.Array
			_ = _229_v58
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw44
			_229_v58 = _nw44
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_229_v58), 0))
			_ = _index41
			(_229_v58).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("p"), _181_v23), (_index41).Int())
			var _230_v59 _dafny.Set
			_ = _230_v59
			_230_v59 = _dafny.SetOf((_dafny.IntOfInt64(-451)).Cmp(_this.F11) == 0)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_229_v58), 0))
			_ = _index42
			var _rhs20 _dafny.Sequence = _181_v23
			_ = _rhs20
			var _rhs21 _dafny.Set = _230_v59
			_ = _rhs21
			var _lhs9 _dafny.Array = _229_v58
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_229_v58), 0))
			_ = _lhs10
			(_lhs9).ArraySet1(_rhs20, (_lhs10).Int())
			_230_v59 = _rhs21
		} else {
			var _231_v60 _dafny.MultiSet
			_ = _231_v60
			_231_v60 = _dafny.MultiSetOf(_this.F11)
			var _232_v61 _dafny.Map
			_ = _232_v61
			_232_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), _dafny.MultiSetOf((func() _dafny.Int {
				if (_231_v60).Contains((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int)) {
					return (_231_v60).Multiplicity((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int))
				}
				return _dafny.IntOfUint32((_181_v23).Cardinality())
			})()))
			var _233_v62 D3
			_ = _233_v62
			_233_v62 = Companion_D3_.Create_DC8_((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int))
			var _234_v63 _dafny.Map
			_ = _234_v63
			_234_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v0, ((func() _dafny.MultiSet {
				if (_232_v61).Contains((_dafny.Zero).Minus((_233_v62).Dtor_cf14())) {
					return (_232_v61).Get((_dafny.Zero).Minus((_233_v62).Dtor_cf14())).(_dafny.MultiSet)
				}
				return _231_v60
			})()).Difference(_231_v60))
			_234_v63 = _234_v63
			_149_v0 = (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool)
			var _235_v64 _dafny.Array
			_ = _235_v64
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_2
			var _nw45 _dafny.Array
			_ = _nw45
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw45 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_236_v23 _dafny.Sequence, _237_v7 _dafny.Array, _238_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_239_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_236_v23, _dafny.UnicodeSeqOfUtf8Bytes("jtwjl")), (Companion_Default___.SafeIndex((_237_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_237_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_236_v23, _dafny.UnicodeSeqOfUtf8Bytes("jtwjl"))).Cardinality()))).Uint32(), _238_v4)
					}
				})(_181_v23, _156_v7, _153_v4)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw45 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw45).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw45).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_235_v64 = _nw45
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_235_v64), 0))
			_ = _index43
			(_235_v64).ArraySet1(_181_v23, (_index43).Int())
			var _240_v65 _dafny.Map
			_ = _240_v65
			_240_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v0, _dafny.UnicodeSeqOfUtf8Bytes("rotcipwa"))
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_235_v64), 0))
			_ = _index44
			(_235_v64).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm22(_149_v0, (_dafny.Zero).Minus(_this.F11), globalState), (func() _dafny.Sequence {
				if (_240_v65).Contains(!((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool))) {
					return (_240_v65).Get(!((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool))).(_dafny.Sequence)
				}
				return _181_v23
			})()), (_index44).Int())
			var _241_v66 _dafny.Set
			_ = _241_v66
			_241_v66 = _dafny.SetOf((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), (_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool))
			var _242_v67 _dafny.Map
			_ = _242_v67
			_242_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v5, _149_v0)
			var _243_v68 _dafny.Map
			_ = _243_v68
			_243_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_241_v66).Cardinality(), (_242_v67).Merge(_242_v67))
			_243_v68 = (_243_v68).Update((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), _242_v67)
			var _244_v69 *C0
			_ = _244_v69
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__()
			_244_v69 = _nw46
		}
		r0 = _dafny.Companion_Sequence_.Contains(_181_v23, _153_v4)
		var _245_v70 _dafny.Array
		_ = _245_v70
		var _nwElement0_11 _dafny.Array = _156_v7
		_ = _nwElement0_11
		var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(17))
		_ = _nw47
		(_nw47).ArraySet1(_nwElement0_11, 0)
		(_nw47).ArraySet1(_156_v7, 1)
		(_nw47).ArraySet1(_156_v7, 2)
		(_nw47).ArraySet1(_156_v7, 3)
		(_nw47).ArraySet1(_156_v7, 4)
		(_nw47).ArraySet1(_156_v7, 5)
		(_nw47).ArraySet1(_156_v7, 6)
		(_nw47).ArraySet1(_156_v7, 7)
		(_nw47).ArraySet1(_156_v7, 8)
		(_nw47).ArraySet1(_156_v7, 9)
		(_nw47).ArraySet1(_156_v7, 10)
		(_nw47).ArraySet1(_156_v7, 11)
		(_nw47).ArraySet1(_156_v7, 12)
		(_nw47).ArraySet1(_156_v7, 13)
		(_nw47).ArraySet1(_156_v7, 14)
		(_nw47).ArraySet1(_156_v7, 15)
		(_nw47).ArraySet1(_156_v7, 16)
		_245_v70 = _nw47
		var _246_v71 _dafny.Map
		_ = _246_v71
		_246_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_154_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_154_v5), 0))).Int()).(bool), Companion_D0_.Create_DC0_((_156_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_156_v7), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("lrqnahqbq"), _245_v70))
		var _247_v72 _dafny.Map
		_ = _247_v72
		_247_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_246_v71).Cardinality(), _this.F11), (_this).F10())
		r1 = _247_v72
		return r0, r1
	}
}
func (_this *C1) M5(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) {
	{
		var _248_v1 bool
		_ = _248_v1
		_248_v1 = false
		var _249_v2 _dafny.Sequence
		_ = _249_v2
		_249_v2 = _dafny.SeqOf(_248_v1)
		var _250_i0 _dafny.Int
		_ = _250_i0
		_250_i0 = _dafny.Zero
		{
			for !(_dafny.Companion_Sequence_.Contains(_249_v2, !(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(767))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_256_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_257_i1 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Set {
						var _coll9 = _dafny.NewBuilder()
						_ = _coll9
						for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(109), _dafny.IntOfInt64(909))); ; {
							_compr_9, _ok10 := _iter10()
							if !_ok10 {
								break
							}
							var _258_v0 _dafny.Int
							_258_v0 = interface{}(_compr_9).(_dafny.Int)
							if ((_dafny.IntOfInt64(109)).Cmp(_258_v0) <= 0) && ((_258_v0).Cmp(_dafny.IntOfInt64(909)) < 0) {
								_coll9.Add(Companion_Default___.SafeModuloInt(_258_v0, _256_p0))
							}
						}
						return _coll9.ToSet()
					}()).Cardinality())).Cardinality()
				}
			})(p0))), p0)))) {
				{
					if (_250_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_250_i0 = (_250_i0).Plus(_dafny.One)
					var _251_v3 _dafny.Array
					_ = _251_v3
					var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
					_ = _nw48
					_251_v3 = _nw48
					var _252_v4 _dafny.Sequence
					_ = _252_v4
					_252_v4 = _dafny.SeqOf(_251_v3)
					_252_v4 = _252_v4
					(_this).F11 = p0
					var _253_v5 _dafny.Map
					_ = _253_v5
					_253_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v1, p0)
					var _254_v6 _dafny.Map
					_ = _254_v6
					_254_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v5, _248_v1)
					_254_v6 = _254_v6
					var _255_v7 _dafny.Map
					_ = _255_v7
					_255_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint), p0)
					_248_v1 = !((func() bool {
						if _248_v1 {
							return (_this).Fm21(_255_v7, p0, p0, globalState)
						}
						return (_this).Fm21(_255_v7, p0, Companion_Default___.Fm24(p0, _248_v1, _248_v1, _dafny.IntOfUint32((_249_v2).Cardinality()), globalState), globalState)
					})())
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _259_v8 _dafny.Set
		_ = _259_v8
		_259_v8 = _dafny.SetOf(_this.F11)
		_248_v1 = (_259_v8).IsSubsetOf((_259_v8).Intersection(_259_v8))
		var _260_v9 _dafny.Sequence
		_ = _260_v9
		_260_v9 = _dafny.SeqOf(_this.F11, _dafny.IntOfUint32((_249_v2).Cardinality()))
		var _261_v10 D2
		_ = _261_v10
		_261_v10 = Companion_D2_.Create_DC5_(_260_v9)
		var _pat_let_tv0 = _248_v1
		_ = _pat_let_tv0
		_248_v1 = func(_source5 D2) bool {
			if _source5.Is_DC6() {
				var _262___mcc_h0 _dafny.Int = _source5.Get_().(D2_DC6).Cf11
				_ = _262___mcc_h0
				var _263___mcc_h1 bool = _source5.Get_().(D2_DC6).Cf12
				_ = _263___mcc_h1
				var _264_cf12 bool = _263___mcc_h1
				_ = _264_cf12
				var _265_cf11 _dafny.Int = _262___mcc_h0
				_ = _265_cf11
				return _pat_let_tv0
			} else {
				var _266___mcc_h2 _dafny.Sequence = _source5.Get_().(D2_DC5).Cf10
				_ = _266___mcc_h2
				var _267_cf10 _dafny.Sequence = _266___mcc_h2
				_ = _267_cf10
				return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("xiwja"), _dafny.CodePoint('o'))
			}
		}(_261_v10)
		var _268_v11 D1
		_ = _268_v11
		_268_v11 = Companion_D1_.Create_DC2_(_248_v1)
		var _pat_let_tv1 = _248_v1
		_ = _pat_let_tv1
		_268_v11 = func(_pat_let0_0 D1) D1 {
			return func(_269_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let1_0 bool) D1 {
					return func(_270_dt__update_hcf4_h0 bool) D1 {
						return Companion_D1_.Create_DC2_(_270_dt__update_hcf4_h0)
					}(_pat_let1_0)
				}(_pat_let_tv1)
			}(_pat_let0_0)
		}(Companion_D1_.Create_DC2_(_248_v1))
		var _271_v12 _dafny.Map
		_ = _271_v12
		_271_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wnc"), _dafny.IntOfInt64(791))
		var _272_v13 _dafny.Map
		_ = _272_v13
		_272_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), (_271_v12).Cardinality())
		var _273_v14 _dafny.Set
		_ = _273_v14
		_273_v14 = _dafny.SetOf(_248_v1)
		var _274_v15 _dafny.Map
		_ = _274_v15
		_274_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _248_v1)
		var _275_v16 D1
		_ = _275_v16
		_275_v16 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(-768), _dafny.CodePoint('o'), true, _dafny.IntOfUint32((p1).Cardinality()))
		var _276_v17 _dafny.Array
		_ = _276_v17
		var _nwElement0_12 bool = (_this).Fm21(_272_v13, (_dafny.Zero).Minus(p0), p0, globalState)
		_ = _nwElement0_12
		var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(22))
		_ = _nw49
		(_nw49).ArraySet1(_nwElement0_12, 0)
		(_nw49).ArraySet1(!_dafny.Companion_Sequence_.Equal(_260_v9, _dafny.SeqOf(_dafny.IntOfUint32((_249_v2).Cardinality()), _this.F11)), 1)
		(_nw49).ArraySet1(!(_273_v14).Equals(_273_v14), 2)
		(_nw49).ArraySet1(!(false), 3)
		(_nw49).ArraySet1(_248_v1, 4)
		(_nw49).ArraySet1(_248_v1, 5)
		(_nw49).ArraySet1(_248_v1, 6)
		(_nw49).ArraySet1(_248_v1, 7)
		(_nw49).ArraySet1(true, 8)
		(_nw49).ArraySet1(_248_v1, 9)
		(_nw49).ArraySet1(_248_v1, 10)
		(_nw49).ArraySet1(_248_v1, 11)
		(_nw49).ArraySet1((Companion_Default___.Fm24(_dafny.IntOfInt64(434), (_249_v2).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_249_v2).Cardinality()))).Uint32()).(bool), _248_v1, ((_this).F10()).Cardinality(), globalState)).Cmp((_274_v15).Cardinality()) >= 0, 12)
		(_nw49).ArraySet1(_248_v1, 13)
		(_nw49).ArraySet1(_248_v1, 14)
		(_nw49).ArraySet1((_this).Fm21(_272_v13, p0, _this.F11, globalState), 15)
		(_nw49).ArraySet1(_248_v1, 16)
		(_nw49).ArraySet1((_this).Fm21(_272_v13, Companion_Default___.Fm24(_dafny.IntOfInt64(899), _248_v1, false, _dafny.IntOfInt64(345), globalState), _dafny.IntOfUint32((_dafny.SeqOf(_248_v1, false, _248_v1, _248_v1, _248_v1)).Cardinality()), globalState), 17)
		(_nw49).ArraySet1(_248_v1, 18)
		(_nw49).ArraySet1(_248_v1, 19)
		(_nw49).ArraySet1((_275_v16).Dtor_cf8(), 20)
		(_nw49).ArraySet1(_248_v1, 21)
		_276_v17 = _nw49
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_276_v17), 0))); ; {
			_guard_loop_1, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _277_i2 _dafny.Int
			_277_i2 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_277_i2).Sign() != -1) && ((_277_i2).Cmp(_dafny.ArrayLen((_276_v17), 0)) < 0)) {
				(_276_v17).ArraySet1(((_dafny.MultiSetFromSeq(_260_v9)).Intersection((Companion_D3_.Create_DC7_(_dafny.MultiSetOf(_dafny.IntOfInt64(709), p0))).Dtor_cf13())).IsDisjointFrom((_dafny.MultiSetOf(p0, _this.F11, _dafny.IntOfUint32((p1).Cardinality()))).Union(_dafny.MultiSetFromSeq(_260_v9))), (_277_i2).Int())
			}
		}
		if _248_v1 {
			var _pat_let_tv2 = p0
			_ = _pat_let_tv2
			if ((_248_v1) && ((_this).Fm21(_272_v13, _this.F11, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("weqkh")).Cardinality()), globalState))) || ((func(_pat_let2_0 D2) D2 {
				return func(_278_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let3_0 _dafny.Int) D2 {
						return func(_279_dt__update_hcf11_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC6_(_279_dt__update_hcf11_h0, (_278_dt__update__tmp_h1).Dtor_cf12())
						}(_pat_let3_0)
					}(_pat_let_tv2)
				}(_pat_let2_0)
			}(Companion_D2_.Create_DC6_(p0, _248_v1))).Dtor_cf12()) {
				(_this).F11 = _this.F11
				_248_v1 = !(_248_v1)
				(_this).F11 = _this.F11
				var _280_v18 *C0
				_ = _280_v18
				var _nw50 *C0 = New_C0_()
				_ = _nw50
				_nw50.Ctor__()
				_280_v18 = _nw50
				var _281_v19 _dafny.Array
				_ = _281_v19
				var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw51
				_281_v19 = _nw51
				var _282_v20 _dafny.Array
				_ = _282_v20
				var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw52
				_282_v20 = _nw52
				var _283_v21 _dafny.Sequence
				_ = _283_v21
				_283_v21 = _dafny.SeqOf(_282_v20)
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_281_v19), 0))
				_ = _index45
				(_281_v19).ArraySet1((_283_v21).Select((Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((_283_v21).Cardinality()))).Uint32()).(_dafny.Array), (_index45).Int())
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_281_v19), 0))
				_ = _index46
				(_281_v19).ArraySet1(_282_v20, (_index46).Int())
			} else {
				_261_v10 = _261_v10
				_268_v11 = _268_v11
				var _284_v22 _dafny.Array
				_ = _284_v22
				var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw53
				_284_v22 = _nw53
				_284_v22 = _284_v22
				_248_v1 = (Companion_Default___.SafeDivisionInt(_this.F11, p0)).Cmp((_this.F11).Minus(_dafny.IntOfInt64(618))) <= 0
				var _285_v23 _dafny.Map
				_ = _285_v23
				_285_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_259_v8).Cardinality(), p0)
				var _286_v24 _dafny.CodePoint
				_ = _286_v24
				_286_v24 = _dafny.CodePoint('n')
				(_this).F11 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_285_v23).Cardinality(), _this.F11), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _286_v24))).Cardinality()))
			}
			var _287_v25 *C0
			_ = _287_v25
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__()
			_287_v25 = _nw54
			var _288_v26 _dafny.Array
			_ = _288_v26
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_3
			var _nw55 _dafny.Array
			_ = _nw55
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw55 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = (func(_289_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_290_i3 _dafny.Int) _dafny.Int {
						return (_290_i3).Plus(_289_p0)
					}
				})(p0)
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
			_288_v26 = _nw55
			var _291_v27 _dafny.Sequence
			_ = _291_v27
			_291_v27 = _dafny.SeqOf(_288_v26)
			(_this).F11 = (_dafny.IntOfUint32((_291_v27).Cardinality())).Plus(_this.F11)
			if false {
				(_this).F11 = p0
				var _292_v28 _dafny.Map
				_ = _292_v28
				_292_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _293_v29 _dafny.Map
				_ = _293_v29
				_293_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_292_v28, _248_v1)
				var _294_v30 _dafny.Map
				_ = _294_v30
				_294_v30 = _293_v29
				var _295_v31 _dafny.Array
				_ = _295_v31
				var _nwElement0_13 _dafny.Array = _288_v26
				_ = _nwElement0_13
				var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(6))
				_ = _nw56
				(_nw56).ArraySet1(_nwElement0_13, 0)
				(_nw56).ArraySet1(_288_v26, 1)
				(_nw56).ArraySet1(_288_v26, 2)
				(_nw56).ArraySet1(_288_v26, 3)
				(_nw56).ArraySet1(_288_v26, 4)
				(_nw56).ArraySet1(_288_v26, 5)
				_295_v31 = _nw56
				var _296_v32 D0
				_ = _296_v32
				_296_v32 = Companion_D0_.Create_DC0_(_this.F11, p1, _295_v31)
				var _297_v33 _dafny.Map
				_ = _297_v33
				_297_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_293_v29).Cardinality(), _296_v32)
				var _298_v34 _dafny.Array
				_ = _298_v34
				_298_v34 = _276_v17
				var _299_v35 _dafny.Array
				_ = _299_v35
				var _nwElement0_14 _dafny.Array = _276_v17
				_ = _nwElement0_14
				var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(28))
				_ = _nw57
				(_nw57).ArraySet1(_nwElement0_14, 0)
				(_nw57).ArraySet1(_276_v17, 1)
				(_nw57).ArraySet1(_276_v17, 2)
				(_nw57).ArraySet1((_298_v34), 3)
				(_nw57).ArraySet1(_276_v17, 4)
				(_nw57).ArraySet1(_276_v17, 5)
				(_nw57).ArraySet1(_276_v17, 6)
				(_nw57).ArraySet1((func() _dafny.Array {
					if (_268_v11).Dtor_cf4() {
						return _276_v17
					}
					return _276_v17
				})(), 7)
				(_nw57).ArraySet1(_276_v17, 8)
				(_nw57).ArraySet1(_276_v17, 9)
				(_nw57).ArraySet1(_276_v17, 10)
				(_nw57).ArraySet1(_276_v17, 11)
				(_nw57).ArraySet1(_276_v17, 12)
				(_nw57).ArraySet1(_276_v17, 13)
				(_nw57).ArraySet1(_276_v17, 14)
				(_nw57).ArraySet1(_276_v17, 15)
				(_nw57).ArraySet1(_276_v17, 16)
				(_nw57).ArraySet1(_276_v17, 17)
				(_nw57).ArraySet1(_276_v17, 18)
				(_nw57).ArraySet1(_276_v17, 19)
				(_nw57).ArraySet1(_276_v17, 20)
				(_nw57).ArraySet1(_276_v17, 21)
				(_nw57).ArraySet1(_276_v17, 22)
				(_nw57).ArraySet1((func() _dafny.Array {
					if _248_v1 {
						return _276_v17
					}
					return _276_v17
				})(), 23)
				(_nw57).ArraySet1(_276_v17, 24)
				(_nw57).ArraySet1(_276_v17, 25)
				(_nw57).ArraySet1(_276_v17, 26)
				(_nw57).ArraySet1((_298_v34), 27)
				_299_v35 = _nw57
				var _rhs22 _dafny.Map = _297_v33
				_ = _rhs22
				var _rhs23 _dafny.Array = _299_v35
				_ = _rhs23
				var _rhs24 _dafny.Int = ((_dafny.Zero).Minus(_this.F11)).Minus(_dafny.IntOfUint32((_260_v9).Cardinality()))
				_ = _rhs24
				var _lhs11 *C1 = _this
				_ = _lhs11
				_297_v33 = _rhs22
				_299_v35 = _rhs23
				_lhs11.F11 = _rhs24
				var _300_v36 _dafny.Map
				_ = _300_v36
				_300_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_249_v2, _249_v2), _249_v2)
				_300_v36 = (_300_v36).Update(_249_v2, _249_v2)
				var _301_v37 _dafny.Map
				_ = _301_v37
				_301_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _248_v1)
				var _302_v38 _dafny.Map
				_ = _302_v38
				_302_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_276_v17, _dafny.MultiSetOf((_301_v37).Cardinality()))
				var _303_v39 _dafny.MultiSet
				_ = _303_v39
				_303_v39 = _dafny.MultiSetOf(_dafny.IntOfUint32((p1).Cardinality()))
				var _304_v40 _dafny.CodePoint
				_ = _304_v40
				_304_v40 = _dafny.CodePoint('k')
				var _305_v41 _dafny.Map
				_ = _305_v41
				_305_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _304_v40)
				var _306_v42 _dafny.Map
				_ = _306_v42
				_306_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.MultiSet {
					if (_302_v38).Contains(_276_v17) {
						return (_302_v38).Get(_276_v17).(_dafny.MultiSet)
					}
					return _303_v39
				})()).Union(_dafny.MultiSetOf(Companion_Default___.Fm24(_this.F11, _248_v1, (func() bool {
					if (_301_v37).Contains(_dafny.IntOfUint32((_260_v9).Cardinality())) {
						return (_301_v37).Get(_dafny.IntOfUint32((_260_v9).Cardinality())).(bool)
					}
					return _248_v1
				})(), (_305_v41).Cardinality(), globalState), _this.F11, _dafny.IntOfInt64(-835), _this.F11)), Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_303_v39).Contains(p0) {
						return (_303_v39).Multiplicity(p0)
					}
					return p0
				})(), _this.F11))
				var _307_v43 D2
				_ = _307_v43
				_307_v43 = Companion_D2_.Create_DC6_(p0, _248_v1)
				var _308_v44 _dafny.Map
				_ = _308_v44
				_308_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v1, _dafny.IntOfInt64(727))
				var _rhs25 _dafny.Map = _306_v42
				_ = _rhs25
				var _rhs26 _dafny.Int = (func() _dafny.Int {
					if (_308_v44).Contains(_248_v1) {
						return _this.F11
					}
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_249_v2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_249_v2).Cardinality()))).Uint32(), _248_v1), _249_v2), (Companion_Default___.SafeIndex((_308_v44).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_249_v2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_249_v2).Cardinality()))).Uint32(), _248_v1), _249_v2)).Cardinality()))).Uint32(), true)).Cardinality())
				})()
				_ = _rhs26
				var _rhs27 D2 = (func() D2 {
					if (_259_v8).IsProperSubsetOf(_259_v8) {
						return _307_v43
					}
					return Companion_D2_.Create_DC6_(_this.F11, _248_v1)
				})()
				_ = _rhs27
				var _rhs28 bool = (_273_v14).IsSubsetOf(_dafny.SetOf(false))
				_ = _rhs28
				var _lhs12 *C1 = _this
				_ = _lhs12
				_306_v42 = _rhs25
				_lhs12.F11 = _rhs26
				_307_v43 = _rhs27
				_248_v1 = _rhs28
				var _309_v45 _dafny.Sequence
				_ = _309_v45
				_309_v45 = _dafny.UnicodeSeqOfUtf8Bytes("aqinu")
				_309_v45 = p1
			} else {
				(_this).F11 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((p0).Minus(p0)), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, (_dafny.MultiSetOf(_248_v1)).Cardinality())))
				var _310_v46 D3
				_ = _310_v46
				_310_v46 = Companion_D3_.Create_DC8_(_dafny.IntOfInt64(175))
				var _311_v47 _dafny.Map
				_ = _311_v47
				_311_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_310_v46, _248_v1)
				var _312_v48 _dafny.CodePoint
				_ = _312_v48
				_312_v48 = _dafny.CodePoint('x')
				_311_v47 = ((func() _dafny.Map {
					if _248_v1 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_310_v46, (_this).Fm21(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_312_v48, _this.F11), _this.F11, p0, globalState))
					}
					return _311_v47
				})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_310_v46, true))
				_276_v17 = _276_v17
				_248_v1 = _248_v1
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_276_v17), 0))
				_ = _index47
				(_276_v17).ArraySet1(_248_v1, (_index47).Int())
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_276_v17), 0))
				_ = _index48
				(_276_v17).ArraySet1(_248_v1, (_index48).Int())
			}
			_248_v1 = ((_dafny.Zero).Minus(p0)).Cmp(p0) >= 0
		} else {
			var _313_v49 _dafny.Map
			_ = _313_v49
			_313_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_248_v1)), p0)
			(_this).F11 = (func() _dafny.Int {
				if (_313_v49).Contains(_248_v1) {
					return (_313_v49).Get(_248_v1).(_dafny.Int)
				}
				return _this.F11
			})()
			var _314_v50 _dafny.CodePoint
			_ = _314_v50
			_314_v50 = _dafny.CodePoint('y')
			var _315_v51 _dafny.Map
			_ = _315_v51
			_315_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _314_v50)
			_315_v51 = (_315_v51).Update(_this.F11, (func() _dafny.CodePoint {
				if _248_v1 {
					return _314_v50
				}
				return _314_v50
			})())
			_249_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_248_v1), (Companion_D7_.Create_DC14_(_249_v2)).Dtor_cf20())
			var _316_v52 _dafny.Sequence
			_ = _316_v52
			_316_v52 = _dafny.SeqOf(_274_v15)
			var _317_v53 D4
			_ = _317_v53
			_317_v53 = Companion_D4_.Create_DC10_(_248_v1)
			var _318_v54 _dafny.Map
			_ = _318_v54
			_318_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_317_v53).Dtor_cf16(), _248_v1)))
			_248_v1 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_316_v52, _316_v52), (func() _dafny.Sequence {
				if (_318_v54).Contains(p0) {
					return (_318_v54).Get(p0).(_dafny.Sequence)
				}
				return _316_v52
			})())
			var _319_v55 _dafny.MultiSet
			_ = _319_v55
			_319_v55 = _dafny.MultiSetOf(Companion_Default___.Fm24(p0, _248_v1, _248_v1, p0, globalState), p0)
			var _320_v56 _dafny.Sequence
			_ = _320_v56
			_320_v56 = _dafny.SeqOf(_319_v55, _319_v55)
			(_this).F11 = (((_319_v55).Union((_320_v56).Select((Companion_Default___.SafeIndex(_this.F11, _dafny.IntOfUint32((_320_v56).Cardinality()))).Uint32()).(_dafny.MultiSet))).Cardinality()).Minus(p0)
		}
	}
}
func (_this *C1) F10() _dafny.MultiSet {
	{
		return _this._f10
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f8 _dafny.Sequence
	_f9 _dafny.Sequence
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f8 = _dafny.EmptySeq
	_this._f9 = _dafny.EmptySeq
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

func (_this *C2) Ctor__(f8 _dafny.Sequence, f9 _dafny.Sequence) {
	{
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *C2) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return ((_dafny.MultiSetOf(true)).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true), false, false)))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false))))
	}
}
func (_this *C2) Fm17(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_dafny.Zero).Minus(((_dafny.MultiSetOf(false)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D2_.Create_DC6_(_dafny.IntOfInt64(681), true)).Dtor_cf12())))).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(-645)))).Cardinality())) >= 0
	}
}
func (_this *C2) Fm18(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !((func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter12 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(216), !(true)))).Elements()); ; {
				_compr_10, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _321_v1 _dafny.Map
				_321_v1 = interface{}(_compr_10).(_dafny.Map)
				if (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(216), !(true)))).Contains(_321_v1) {
					_coll10.Add(_321_v1)
				}
			}
			return _coll10.ToSet()
		}()).IsSubsetOf((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(531), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-595), true))).Difference(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F8()).Cardinality()), true), func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(114), _dafny.IntOfInt64(409))); ; {
				_compr_11, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _322_v0 _dafny.Int
				_322_v0 = interface{}(_compr_11).(_dafny.Int)
				if ((_dafny.IntOfInt64(114)).Cmp(_322_v0) <= 0) && ((_322_v0).Cmp(_dafny.IntOfInt64(409)) < 0) {
					_coll11.Add(Companion_Default___.SafeDivisionInt(_322_v0, _dafny.IntOfInt64(736)), true)
				}
			}
			return _coll11.ToMap()
		}()))))
	}
}
func (_this *C2) M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _323_v0 _dafny.Int
		_ = _323_v0
		_323_v0 = _dafny.IntOfInt64(758)
		var _324_v1 bool
		_ = _324_v1
		_324_v1 = true
		_323_v0 = (_323_v0).Minus(Companion_Default___.Fm19(_324_v1, _323_v0, globalState))
		var _325_v2 _dafny.Array
		_ = _325_v2
		var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw58
		_325_v2 = _nw58
		for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_325_v2), 0))); ; {
			_guard_loop_2, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _326_i0 _dafny.Int
			_326_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_326_i0).Sign() != -1) && ((_326_i0).Cmp(_dafny.ArrayLen((_325_v2), 0)) < 0)) {
				(_325_v2).ArraySet1((_326_i0).Minus((func() _dafny.Int {
					if _324_v1 {
						return _323_v0
					}
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(570))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg16 _dafny.Int) interface{} {
							return coer16(arg16)
						}
					}(func(_327_i1 _dafny.Int) _dafny.Sequence {
						return (_this).F8()
					}))).Cardinality())
				})()), (_326_i0).Int())
			}
		}
		var _328_v3 _dafny.Array
		_ = _328_v3
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_4
		var _nw59 _dafny.Array
		_ = _nw59
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw59 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Set = (func(_329_v0 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_330_i2 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_329_v0)
				}
			})(_323_v0)
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
		_328_v3 = _nw59
		var _331_v4 _dafny.Set
		_ = _331_v4
		_331_v4 = _dafny.SetOf((_dafny.Zero).Minus(_323_v0), _323_v0, _323_v0)
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_328_v3), 0))
		_ = _index49
		(_328_v3).ArraySet1((_331_v4).Union(_331_v4), (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
		_ = _index50
		(_325_v2).ArraySet1((_dafny.Zero).Minus(_323_v0), (_index50).Int())
		var _332_v5 _dafny.CodePoint
		_ = _332_v5
		_332_v5 = _dafny.CodePoint('g')
		var _333_v6 _dafny.Set
		_ = _333_v6
		_333_v6 = _dafny.SetOf(_332_v5, _332_v5)
		var _334_v7 _dafny.Map
		_ = _334_v7
		_334_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_333_v6, _325_v2)
		var _335_v8 _dafny.Array
		_ = _335_v8
		var _nwElement0_15 _dafny.Array = (func() _dafny.Array {
			if _324_v1 {
				return _325_v2
			}
			return _325_v2
		})()
		_ = _nwElement0_15
		var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(24))
		_ = _nw60
		(_nw60).ArraySet1(_nwElement0_15, 0)
		(_nw60).ArraySet1(_325_v2, 1)
		(_nw60).ArraySet1(_325_v2, 2)
		(_nw60).ArraySet1(_325_v2, 3)
		(_nw60).ArraySet1((func() _dafny.Array {
			if (_334_v7).Contains(_333_v6) {
				return (_334_v7).Get(_333_v6).(_dafny.Array)
			}
			return _325_v2
		})(), 4)
		(_nw60).ArraySet1(_325_v2, 5)
		(_nw60).ArraySet1(_325_v2, 6)
		(_nw60).ArraySet1(_325_v2, 7)
		(_nw60).ArraySet1(_325_v2, 8)
		(_nw60).ArraySet1(_325_v2, 9)
		(_nw60).ArraySet1(_325_v2, 10)
		(_nw60).ArraySet1(_325_v2, 11)
		(_nw60).ArraySet1(_325_v2, 12)
		(_nw60).ArraySet1(_325_v2, 13)
		(_nw60).ArraySet1(_325_v2, 14)
		(_nw60).ArraySet1(_325_v2, 15)
		(_nw60).ArraySet1(_325_v2, 16)
		(_nw60).ArraySet1((func() _dafny.Array {
			if _324_v1 {
				return _325_v2
			}
			return _325_v2
		})(), 17)
		(_nw60).ArraySet1(_325_v2, 18)
		(_nw60).ArraySet1(_325_v2, 19)
		(_nw60).ArraySet1(_325_v2, 20)
		(_nw60).ArraySet1(_325_v2, 21)
		(_nw60).ArraySet1(_325_v2, 22)
		(_nw60).ArraySet1((func() _dafny.Array {
			if !(_324_v1) {
				return _325_v2
			}
			return _325_v2
		})(), 23)
		_335_v8 = _nw60
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_328_v3), 0))
		_ = _index51
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
		_ = _index52
		var _rhs29 _dafny.Set = Companion_Default___.Fm20(true, (_this).F9(), (_323_v0).Cmp(_323_v0) != 0, globalState)
		_ = _rhs29
		var _rhs30 _dafny.Int = _323_v0
		_ = _rhs30
		var _rhs31 _dafny.Array = _335_v8
		_ = _rhs31
		var _lhs13 _dafny.Array = _328_v3
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_328_v3), 0))
		_ = _lhs14
		var _lhs15 _dafny.Array = _325_v2
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
		_ = _lhs16
		(_lhs13).ArraySet1(_rhs29, (_lhs14).Int())
		(_lhs15).ArraySet1(_rhs30, (_lhs16).Int())
		_335_v8 = _rhs31
		var _336_v9 D1
		_ = _336_v9
		_336_v9 = Companion_D1_.Create_DC2_(!(_324_v1))
		var _source6 D1 = _336_v9
		_ = _source6
		if _source6.Is_DC3() {
			var _337___mcc_h0 _dafny.Int = _source6.Get_().(D1_DC3).Cf5
			_ = _337___mcc_h0
			var _338_cf5 _dafny.Int = _337___mcc_h0
			_ = _338_cf5
			var _339_v10 _dafny.Array
			_ = _339_v10
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw61
			_339_v10 = _nw61
			var _340_v11 _dafny.Map
			_ = _340_v11
			_340_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v10, _324_v1)
			var _341_v12 D4
			_ = _341_v12
			_341_v12 = Companion_D4_.Create_DC10_(true)
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
			_ = _index53
			(_325_v2).ArraySet1((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v10, true)).Merge(_340_v11)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v10, (_341_v12).Dtor_cf16()))).Cardinality(), (_index53).Int())
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
			_ = _index54
			(_325_v2).ArraySet1((_338_cf5).Plus(_dafny.IntOfInt64(908)), (_index54).Int())
			var _342_v13 _dafny.Sequence
			_ = _342_v13
			_342_v13 = _dafny.UnicodeSeqOfUtf8Bytes("ndj")
			var _343_v14 D0
			_ = _343_v14
			_343_v14 = Companion_D0_.Create_DC0_((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("flum"), _335_v8)
			var _344_v15 _dafny.Sequence
			_ = _344_v15
			_344_v15 = _dafny.SeqOf(_324_v1, _324_v1)
			var _345_v16 _dafny.MultiSet
			_ = _345_v16
			_345_v16 = _dafny.MultiSetOf(_324_v1)
			var _346_v17 _dafny.Sequence
			_ = _346_v17
			_346_v17 = _dafny.SeqOf((_dafny.MultiSetFromSeq(_344_v15)).Intersection(_345_v16), _345_v16, _dafny.MultiSetOf(_324_v1))
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
			_ = _index55
			var _rhs32 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_343_v14).Dtor_cf1(), _342_v13), (Companion_Default___.SafeIndex(_323_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_343_v14).Dtor_cf1(), _342_v13)).Cardinality()))).Uint32(), _332_v5)
			_ = _rhs32
			var _rhs33 _dafny.Int = _dafny.IntOfUint32((_346_v17).Cardinality())
			_ = _rhs33
			var _lhs17 _dafny.Array = _325_v2
			_ = _lhs17
			var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
			_ = _lhs18
			_342_v13 = _rhs32
			(_lhs17).ArraySet1(_rhs33, (_lhs18).Int())
			r0 = _324_v1
		} else if _source6.Is_DC4() {
			var _347___mcc_h1 _dafny.Int = _source6.Get_().(D1_DC4).Cf6
			_ = _347___mcc_h1
			var _348___mcc_h2 _dafny.CodePoint = _source6.Get_().(D1_DC4).Cf7
			_ = _348___mcc_h2
			var _349___mcc_h3 bool = _source6.Get_().(D1_DC4).Cf8
			_ = _349___mcc_h3
			var _350___mcc_h4 _dafny.Int = _source6.Get_().(D1_DC4).Cf9
			_ = _350___mcc_h4
			var _351_cf9 _dafny.Int = _350___mcc_h4
			_ = _351_cf9
			var _352_cf8 bool = _349___mcc_h3
			_ = _352_cf8
			var _353_cf7 _dafny.CodePoint = _348___mcc_h2
			_ = _353_cf7
			var _354_cf6 _dafny.Int = _347___mcc_h1
			_ = _354_cf6
			_354_cf6 = _323_v0
			var _355_v18 _dafny.Sequence
			_ = _355_v18
			_355_v18 = _dafny.SeqOf(_352_cf8)
			var _356_v19 _dafny.Map
			_ = _356_v19
			_356_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_353_cf7, _324_v1)
			var _357_v20 _dafny.MultiSet
			_ = _357_v20
			_357_v20 = _dafny.MultiSetOf(_dafny.IntOfInt64(121), (_356_v19).Cardinality(), _354_cf6)
			var _358_v21 _dafny.Map
			_ = _358_v21
			_358_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_354_cf6, !(_352_cf8))
			var _359_v22 *C1
			_ = _359_v22
			var _nw62 *C1 = New_C1_()
			_ = _nw62
			_nw62.Ctor__((_this).Fm0(_352_cf8, !((_this).Fm18(_355_v18, globalState)), (_357_v20).Cardinality(), (_358_v21).Cardinality(), globalState), _323_v0)
			_359_v22 = _nw62
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
			_ = _index56
			var _rhs34 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_355_v18, _dafny.SeqOf(_324_v1, _352_cf8, _324_v1))
			_ = _rhs34
			var _rhs35 _dafny.Int = _323_v0
			_ = _rhs35
			var _lhs19 _dafny.Array = _325_v2
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
			_ = _lhs20
			_352_cf8 = _rhs34
			(_lhs19).ArraySet1(_rhs35, (_lhs20).Int())
			var _360_v23 _dafny.Array
			_ = _360_v23
			var _nw63 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(28))
			_ = _nw63
			_360_v23 = _nw63
			var _361_v24 D2
			_ = _361_v24
			_361_v24 = Companion_D2_.Create_DC6_(_354_cf6, _324_v1)
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_360_v23), 0))
			_ = _index57
			(_360_v23).ArraySet1(_361_v24, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_360_v23), 0))
			_ = _index58
			(_360_v23).ArraySet1(_361_v24, (_index58).Int())
		} else {
			var _362___mcc_h5 bool = _source6.Get_().(D1_DC2).Cf4
			_ = _362___mcc_h5
			var _363_cf4 bool = _362___mcc_h5
			_ = _363_cf4
			if _324_v1 {
				var _364_v25 _dafny.Map
				_ = _364_v25
				_364_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v2, (Companion_Default___.Fm28((Companion_Default___.Fm29(globalState)).Cardinality(), (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState)).Cardinality())
				var _365_v26 _dafny.Sequence
				_ = _365_v26
				_365_v26 = _dafny.SeqOf(_364_v25)
				_364_v25 = ((_365_v26).Select((Companion_Default___.SafeIndex(_323_v0, _dafny.IntOfUint32((_365_v26).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_364_v25)
				_323_v0 = _323_v0
				var _366_v27 D0
				_ = _366_v27
				_366_v27 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(338), _dafny.Companion_Sequence_.Concatenate((_this).F8(), (_this).F9()), _335_v8)
				var _367_v28 _dafny.Map
				_ = _367_v28
				_367_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_332_v5, _323_v0)
				_366_v27 = Companion_D0_.Create_DC0_(((_367_v28).Merge(_367_v28)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(515))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}((func(_368_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_369_i3 _dafny.Int) _dafny.CodePoint {
						return _368_v5
					}
				})(_332_v5))), _335_v8)
				var _370_v29 _dafny.Sequence
				_ = _370_v29
				_370_v29 = _dafny.SeqOf(!(_324_v1), _324_v1)
				var _371_v30 _dafny.Sequence
				_ = _371_v30
				_371_v30 = _dafny.SeqOf(_370_v29)
				var _372_v31 _dafny.Map
				_ = _372_v31
				_372_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(51), _323_v0)
				var _373_v32 _dafny.Sequence
				_ = _373_v32
				_373_v32 = _dafny.SeqOf((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), _323_v0, (_372_v31).Cardinality(), (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int))
				var _374_v33 _dafny.Map
				_ = _374_v33
				_374_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_363_cf4, _dafny.IntOfInt64(345))
				var _375_v34 _dafny.MultiSet
				_ = _375_v34
				_375_v34 = _dafny.MultiSetOf((_374_v33).Cardinality())
				var _376_v35 *C0
				_ = _376_v35
				var _nw64 *C0 = New_C0_()
				_ = _nw64
				_nw64.Ctor__()
				_376_v35 = _nw64
				var _377_v36 D7
				_ = _377_v36
				_377_v36 = Companion_D7_.Create_DC15_(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_370_v29, _dafny.SeqOf(_363_cf4, _324_v1)), _371_v30), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-459))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_378_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
					return func(_379_i4 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-711))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg19 _dafny.Int) interface{} {
								return coer19(arg19)
							}
						}((func(_380_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_381_i5 _dafny.Int) _dafny.CodePoint {
								return _380_v5
							}
						})(_378_v5)))).Cardinality())
					}
				})(_332_v5))), _373_v32)).Cardinality()), (_375_v34).Cardinality(), (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), _376_v35)
				var _382_v37 _dafny.MultiSet
				_ = _382_v37
				_382_v37 = _dafny.MultiSetOf(true)
				var _383_v38 T1
				_ = _383_v38
				var _nw65 *C1 = New_C1_()
				_ = _nw65
				_nw65.Ctor__(_382_v37, (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int))
				_383_v38 = _nw65
				var _384_v39 _dafny.MultiSet
				_ = _384_v39
				_384_v39 = _dafny.MultiSetOf(_383_v38, _383_v38)
				var _pat_let_tv3 = _376_v35
				_ = _pat_let_tv3
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
				_ = _index59
				var _rhs36 _dafny.Int = _dafny.IntOfUint32(((_371_v30).Select((Companion_Default___.SafeIndex(_323_v0, _dafny.IntOfUint32((_371_v30).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
				_ = _rhs36
				var _rhs37 D7 = (func() D7 {
					if !((_dafny.MultiSetOf(_383_v38)).IsDisjointFrom(_384_v39)) {
						return func(_pat_let4_0 D7) D7 {
							return func(_385_dt__update__tmp_h0 D7) D7 {
								return func(_pat_let5_0 *C0) D7 {
									return func(_386_dt__update_hcf25_h0 *C0) D7 {
										return Companion_D7_.Create_DC15_((_385_dt__update__tmp_h0).Dtor_cf21(), (_385_dt__update__tmp_h0).Dtor_cf22(), (_385_dt__update__tmp_h0).Dtor_cf23(), (_385_dt__update__tmp_h0).Dtor_cf24(), _386_dt__update_hcf25_h0)
									}(_pat_let5_0)
								}(_pat_let_tv3)
							}(_pat_let4_0)
						}(_377_v36)
					}
					return _377_v36
				})()
				_ = _rhs37
				var _lhs21 _dafny.Array = _325_v2
				_ = _lhs21
				var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
				_ = _lhs22
				(_lhs21).ArraySet1(_rhs36, (_lhs22).Int())
				_377_v36 = _rhs37
				_323_v0 = Companion_Default___.SafeDivisionInt((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), Companion_Default___.Fm19(_363_cf4, (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState))
			} else {
				_324_v1 = !(_363_cf4)
				var _387_v40 _dafny.Array
				_ = _387_v40
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
				_ = _nw66
				_387_v40 = _nw66
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
				_ = _index60
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
				_ = _index61
				var _rhs38 _dafny.Int = (Companion_Default___.SafeDivisionInt((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int))).Times(((_dafny.Zero).Minus((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int))).Minus((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int)))
				_ = _rhs38
				var _rhs39 _dafny.Int = _dafny.IntOfUint32(((_this).F9()).Cardinality())
				_ = _rhs39
				var _rhs40 _dafny.Array = (func() _dafny.Array {
					if _363_cf4 {
						return _387_v40
					}
					return _387_v40
				})()
				_ = _rhs40
				var _rhs41 _dafny.Int = Companion_Default___.SafeDivisionInt(_323_v0, Companion_Default___.Fm24((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), (_this).Fm17(_332_v5, _363_cf4, (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), Companion_Default___.Fm24(_323_v0, _324_v1, !(_363_cf4), (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState), globalState), !(!(_363_cf4)), _dafny.IntOfUint32(((_this).F9()).Cardinality()), globalState))
				_ = _rhs41
				var _lhs23 _dafny.Array = _325_v2
				_ = _lhs23
				var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
				_ = _lhs24
				var _lhs25 _dafny.Array = _325_v2
				_ = _lhs25
				var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
				_ = _lhs26
				_323_v0 = _rhs38
				(_lhs23).ArraySet1(_rhs39, (_lhs24).Int())
				_387_v40 = _rhs40
				(_lhs25).ArraySet1(_rhs41, (_lhs26).Int())
				var _388_v41 _dafny.Sequence
				_ = _388_v41
				_388_v41 = _dafny.SeqOf(_dafny.IntOfInt64(819))
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_387_v40), 0))
				_ = _index62
				(_387_v40).ArraySet1(_388_v41, (_index62).Int())
				var _389_v42 _dafny.Set
				_ = _389_v42
				_389_v42 = _dafny.SetOf(_324_v1, _324_v1, _324_v1)
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_387_v40), 0))
				_ = _index63
				(_387_v40).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm23((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState), (Companion_Default___.SafeIndex(((_389_v42).Difference(_389_v42)).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm23((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState)).Cardinality()))).Uint32(), _dafny.IntOfInt64(778)), (Companion_Default___.SafeIndex((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm23((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState), (Companion_Default___.SafeIndex(((_389_v42).Difference(_389_v42)).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm23((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState)).Cardinality()))).Uint32(), _dafny.IntOfInt64(778))).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32(((_this).F8()).Cardinality())), _dafny.SeqOf(_323_v0, _323_v0, _dafny.IntOfInt64(766))), (Companion_Default___.SafeIndex(_323_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32(((_this).F8()).Cardinality())), _dafny.SeqOf(_323_v0, _323_v0, _dafny.IntOfInt64(766)))).Cardinality()))).Uint32(), (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int))).Cardinality())), (_index63).Int())
				var _390_v43 *C0
				_ = _390_v43
				var _nw67 *C0 = New_C0_()
				_ = _nw67
				_nw67.Ctor__()
				_390_v43 = _nw67
				var _391_v44 _dafny.Map
				_ = _391_v44
				_391_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_324_v1, _324_v1)
				var _392_v45 D3
				_ = _392_v45
				_392_v45 = Companion_D3_.Create_DC8_((_391_v44).Cardinality())
				var _393_v46 _dafny.Map
				_ = _393_v46
				_393_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D7_.Create_DC15_(_324_v1, ((_328_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_328_v3), 0))).Int()).(_dafny.Set)).Cardinality(), (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), _323_v0, _390_v43)).Dtor_cf24(), (_392_v45).Dtor_cf14())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
				_ = _index64
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_387_v40), 0))
				_ = _index65
				var _rhs42 _dafny.Int = (_dafny.Zero).Minus(_323_v0)
				_ = _rhs42
				var _rhs43 _dafny.Sequence = (func() _dafny.Sequence {
					if _363_cf4 {
						return _dafny.SeqOf((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int))
					}
					return (_387_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_387_v40), 0))).Int()).(_dafny.Sequence)
				})()
				_ = _rhs43
				var _rhs44 bool = true
				_ = _rhs44
				var _rhs45 bool = (func() bool {
					if ((_dafny.Zero).Minus(_323_v0)).Cmp((_393_v46).Cardinality()) <= 0 {
						return _324_v1
					}
					return true
				})()
				_ = _rhs45
				var _lhs27 _dafny.Array = _325_v2
				_ = _lhs27
				var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
				_ = _lhs28
				var _lhs29 _dafny.Array = _387_v40
				_ = _lhs29
				var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_387_v40), 0))
				_ = _lhs30
				(_lhs27).ArraySet1(_rhs42, (_lhs28).Int())
				(_lhs29).ArraySet1(_rhs43, (_lhs30).Int())
				_324_v1 = _rhs44
				_324_v1 = _rhs45
				_363_cf4 = !(_324_v1)
			}
			var _394_v47 _dafny.Array
			_ = _394_v47
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_5
			var _nw68 _dafny.Array
			_ = _nw68
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw68 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Sequence = (func(_395_cf4 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_396_i6 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_395_cf4, _395_cf4, _395_cf4)
					}
				})(_363_cf4)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw68 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw68).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw68).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_394_v47 = _nw68
			var _397_v48 _dafny.Sequence
			_ = _397_v48
			_397_v48 = _dafny.SeqOf(_324_v1, _363_cf4, _363_cf4, false, _324_v1)
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_394_v47), 0))
			_ = _index66
			(_394_v47).ArraySet1(_397_v48, (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_394_v47), 0))
			_ = _index67
			(_394_v47).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm26((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.IntOfUint32((Companion_Default___.Fm26((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), globalState)).Cardinality()))).Uint32(), _324_v1), (_index67).Int())
			var _398_v49 _dafny.Sequence
			_ = _398_v49
			_398_v49 = _dafny.SeqOf((Companion_Default___.Fm24(_323_v0, _363_cf4, false, _dafny.IntOfUint32(((_this).F9()).Cardinality()), globalState)).Minus(_323_v0))
			_323_v0 = _dafny.IntOfUint32((_398_v49).Cardinality())
			var _399_v50 _dafny.Sequence
			_ = _399_v50
			_399_v50 = _dafny.UnicodeSeqOfUtf8Bytes("tsqb")
			_399_v50 = (_this).F8()
		}
		var _400_v51 _dafny.Sequence
		_ = _400_v51
		_400_v51 = _dafny.SeqOf(_324_v1)
		var _401_v53 _dafny.Map
		_ = _401_v53
		_401_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_400_v51).Select((Companion_Default___.SafeIndex((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_400_v51).Cardinality()))).Uint32()).(bool), func() _dafny.Set {
			var _coll12 = _dafny.NewBuilder()
			_ = _coll12
			for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-873), _dafny.IntOfInt64(158))); ; {
				_compr_12, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _402_v52 _dafny.Int
				_402_v52 = interface{}(_compr_12).(_dafny.Int)
				if ((_dafny.IntOfInt64(-873)).Cmp(_402_v52) <= 0) && ((_402_v52).Cmp(_dafny.IntOfInt64(158)) < 0) {
					_coll12.Add((_402_v52).Plus(_dafny.IntOfUint32((_400_v51).Cardinality())))
				}
			}
			return _coll12.ToSet()
		}())
		var _403_v55 _dafny.Sequence
		_ = _403_v55
		_403_v55 = _dafny.SeqOf(_dafny.SetOf(_323_v0, (_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int)), (func() _dafny.Set {
			if (_401_v53).Contains(_324_v1) {
				return (_401_v53).Get(_324_v1).(_dafny.Set)
			}
			return _331_v4
		})(), (_328_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_328_v3), 0))).Int()).(_dafny.Set), func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(53), _dafny.IntOfInt64(673))); ; {
				_compr_13, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _404_v54 _dafny.Int
				_404_v54 = interface{}(_compr_13).(_dafny.Int)
				if ((_dafny.IntOfInt64(53)).Cmp(_404_v54) <= 0) && ((_404_v54).Cmp(_dafny.IntOfInt64(673)) < 0) {
					_coll13.Add((_404_v54).Minus(_323_v0))
				}
			}
			return _coll13.ToSet()
		}())
		var _hi3 _dafny.Int = _323_v0
		_ = _hi3
		for _405_i7 := ((_403_v55).Select((Companion_Default___.SafeIndex((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_403_v55).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(); _405_i7.Cmp(_hi3) < 0; _405_i7 = _405_i7.Plus(_dafny.One) {
			_324_v1 = _324_v1
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
			_ = _index68
			(_325_v2).ArraySet1((_323_v0).Minus(_405_i7), (_index68).Int())
			var _406_v57 _dafny.Array
			_ = _406_v57
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_6
			var _nw69 _dafny.Array
			_ = _nw69
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw69 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) bool = (func(_407_i7 _dafny.Int, _408_v1 bool, _409_v2 _dafny.Array, _410_v0 _dafny.Int) func(_dafny.Int) bool {
					return func(_411_i8 _dafny.Int) bool {
						return (func() _dafny.Set {
							var _coll14 = _dafny.NewBuilder()
							_ = _coll14
							for _iter17 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, (_409_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_409_v2), 0))).Int()).(_dafny.Int)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _410_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _dafny.IntOfUint32((_dafny.SeqOf((_this).F8(), (_this).F9())).Cardinality())))).Elements()); ; {
								_compr_14, _ok17 := _iter17()
								if !_ok17 {
									break
								}
								var _412_v56 _dafny.Map
								_412_v56 = interface{}(_compr_14).(_dafny.Map)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, (_409_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_409_v2), 0))).Int()).(_dafny.Int)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _410_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _dafny.IntOfUint32((_dafny.SeqOf((_this).F8(), (_this).F9())).Cardinality()))), _412_v56) {
									_coll14.Add(_412_v56)
								}
							}
							return _coll14.ToSet()
						}()).IsDisjointFrom(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v1, _407_i7))))
					}
				})(_405_i7, _324_v1, _325_v2, _323_v0)
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
			_406_v57 = _nw69
			_406_v57 = _406_v57
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))
			_ = _index69
			(_325_v2).ArraySet1(_323_v0, (_index69).Int())
		}
		_323_v0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_332_v5, _dafny.CodePoint('j')), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_413_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_414_i9 _dafny.Int) _dafny.CodePoint {
				return _413_v5
			}
		})(_332_v5))), (Companion_Default___.SafeIndex(_323_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_415_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_416_i9 _dafny.Int) _dafny.CodePoint {
				return _415_v5
			}
		})(_332_v5)))).Cardinality()))).Uint32(), _332_v5), (_this).F8()))).Cardinality())
		r0 = _324_v1
		r1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_325_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_325_v2), 0))).Int()).(_dafny.Int), _dafny.MultiSetOf(_324_v1, _324_v1, _324_v1, _324_v1, _324_v1))
		return r0, r1
	}
}
func (_this *C2) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *C2) F9() _dafny.Sequence {
	{
		return _this._f9
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) Ctor__() {
	{
	}
}
func (_this *C3) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		if (!(!(false))) && (!(!(false))) {
			return _dafny.MultiSetOf(!(true))
		} else {
			return _dafny.MultiSetOf(false)
		}
	}
}
func (_this *C3) M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _417_v0 _dafny.Int
		_ = _417_v0
		_417_v0 = _dafny.IntOfInt64(131)
		var _418_v1 _dafny.MultiSet
		_ = _418_v1
		_418_v1 = _dafny.MultiSetOf(_417_v0, _dafny.IntOfInt64(-663), _417_v0)
		var _419_v2 D3
		_ = _419_v2
		_419_v2 = Companion_D3_.Create_DC7_(_418_v1)
		_419_v2 = _419_v2
		var _420_v3 bool
		_ = _420_v3
		_420_v3 = true
		var _421_v4 D4
		_ = _421_v4
		_421_v4 = Companion_D4_.Create_DC10_(_420_v3)
		_421_v4 = _421_v4
		if _420_v3 {
			var _422_v5 _dafny.Set
			_ = _422_v5
			_422_v5 = _dafny.SetOf(_420_v3)
			var _423_v6 _dafny.Sequence
			_ = _423_v6
			_423_v6 = _dafny.SeqOf(_420_v3, _420_v3)
			var _424_v7 _dafny.Set
			_ = _424_v7
			_424_v7 = _dafny.SetOf(_417_v0, Companion_Default___.Fm14(Companion_Default___.Fm14(_417_v0, (_422_v5).Cardinality(), globalState), _417_v0, globalState), _dafny.IntOfUint32((_423_v6).Cardinality()), (_417_v0).Times(Companion_Default___.Fm14(_417_v0, _417_v0, globalState)), (_dafny.Zero).Minus(_417_v0))
			_424_v7 = _424_v7
			var _425_v8 _dafny.Array
			_ = _425_v8
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw70
			_425_v8 = _nw70
			var _426_v9 _dafny.Map
			_ = _426_v9
			_426_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_425_v8, _420_v3)
			_426_v9 = (_426_v9).Update(_425_v8, _420_v3)
			var _427_v10 _dafny.Sequence
			_ = _427_v10
			_427_v10 = _dafny.SeqOf(_423_v6)
			var _428_v11 _dafny.CodePoint
			_ = _428_v11
			_428_v11 = _dafny.CodePoint('q')
			var _429_v12 _dafny.Array
			_ = _429_v12
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_7
			var _nw71 _dafny.Array
			_ = _nw71
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw71 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Int = (func(_430_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_431_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_431_i1, _430_v0)
					}
				})(_417_v0)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw71 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw71).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw71).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_429_v12 = _nw71
			var _432_v13 _dafny.Array
			_ = _432_v13
			var _nwElement0_16 _dafny.Array = _429_v12
			_ = _nwElement0_16
			var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(15))
			_ = _nw72
			(_nw72).ArraySet1(_nwElement0_16, 0)
			(_nw72).ArraySet1(_429_v12, 1)
			(_nw72).ArraySet1(_429_v12, 2)
			(_nw72).ArraySet1(_429_v12, 3)
			(_nw72).ArraySet1(_429_v12, 4)
			(_nw72).ArraySet1(_429_v12, 5)
			(_nw72).ArraySet1(_429_v12, 6)
			(_nw72).ArraySet1(_429_v12, 7)
			(_nw72).ArraySet1(_429_v12, 8)
			(_nw72).ArraySet1(_429_v12, 9)
			(_nw72).ArraySet1(_429_v12, 10)
			(_nw72).ArraySet1(_429_v12, 11)
			(_nw72).ArraySet1(_429_v12, 12)
			(_nw72).ArraySet1(_429_v12, 13)
			(_nw72).ArraySet1(_429_v12, 14)
			_432_v13 = _nw72
			var _433_v14 D0
			_ = _433_v14
			_433_v14 = Companion_D0_.Create_DC0_(_417_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(642))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_434_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_435_i0 _dafny.Int) _dafny.CodePoint {
					return _434_v11
				}
			})(_428_v11))), _432_v13)
			var _436_v15 _dafny.Map
			_ = _436_v15
			_436_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_428_v11, _433_v14)
			var _437_v16 D0
			_ = _437_v16
			_437_v16 = Companion_D0_.Create_DC1_((func() D0 {
				if (_436_v15).Contains(_dafny.CodePoint('h')) {
					return (_436_v15).Get(_dafny.CodePoint('h')).(D0)
				}
				return _433_v14
			})())
			var _438_v17 _dafny.Map
			_ = _438_v17
			_438_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_427_v10, _437_v16)
			_438_v17 = (_438_v17).Update(_427_v10, _437_v16)
			_417_v0 = _417_v0
			var _439_v18 _dafny.Array
			_ = _439_v18
			var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
			_ = _nw73
			_439_v18 = _nw73
			var _440_v19 _dafny.MultiSet
			_ = _440_v19
			_440_v19 = _dafny.MultiSetOf(_420_v3)
			var _441_v20 _dafny.Map
			_ = _441_v20
			_441_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_440_v19).Cardinality()), _428_v11)
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_439_v18), 0))
			_ = _index70
			(_439_v18).ArraySet1((_441_v20).Merge(_441_v20), (_index70).Int())
			var _442_v22 _dafny.Map
			_ = _442_v22
			_442_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_425_v8, func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(11), _dafny.IntOfInt64(664))); ; {
					_compr_15, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _443_v21 _dafny.Int
					_443_v21 = interface{}(_compr_15).(_dafny.Int)
					if ((_dafny.IntOfInt64(11)).Cmp(_443_v21) <= 0) && ((_443_v21).Cmp(_dafny.IntOfInt64(664)) < 0) {
						_coll15.Add((_443_v21).Plus(_417_v0), _dafny.CodePoint('m'))
					}
				}
				return _coll15.ToMap()
			}())
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_439_v18), 0))
			_ = _index71
			(_439_v18).ArraySet1((func() _dafny.Map {
				if (_442_v22).Contains((func() _dafny.Array {
					if false {
						return _425_v8
					}
					return _425_v8
				})()) {
					return (_442_v22).Get((func() _dafny.Array {
						if false {
							return _425_v8
						}
						return _425_v8
					})()).(_dafny.Map)
				}
				return (_441_v20).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_417_v0), _428_v11))
			})(), (_index71).Int())
		} else {
			var _444_v23 _dafny.Array
			_ = _444_v23
			var _nw74 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw74
			_444_v23 = _nw74
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
			_ = _index72
			(_444_v23).ArraySet1(_420_v3, (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
			_ = _index73
			(_444_v23).ArraySet1(!((_418_v1).IsProperSubsetOf(_418_v1)), (_index73).Int())
			var _445_v24 _dafny.CodePoint
			_ = _445_v24
			_445_v24 = _dafny.CodePoint('a')
			var _446_v25 _dafny.Sequence
			_ = _446_v25
			_446_v25 = _dafny.SeqOf(_445_v24)
			var _447_v26 _dafny.Map
			_ = _447_v26
			_447_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v23, _446_v25)
			_446_v25 = (func() _dafny.Sequence {
				if (_447_v26).Contains(_444_v23) {
					return (_447_v26).Get(_444_v23).(_dafny.Sequence)
				}
				return _446_v25
			})()
			_417_v0 = Companion_Default___.Fm14(_417_v0, _417_v0, globalState)
			var _448_v27 _dafny.Map
			_ = _448_v27
			_448_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate(_446_v25, _dafny.Companion_Sequence_.Update(_446_v25, (Companion_Default___.SafeIndex(_417_v0, _dafny.IntOfUint32((_446_v25).Cardinality()))).Uint32(), _445_v24)))
			var _449_v28 _dafny.Sequence
			_ = _449_v28
			_449_v28 = _dafny.SeqOf(_420_v3)
			var _450_v29 D1
			_ = _450_v29
			_450_v29 = Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_449_v28).Cardinality()))
			var _451_v30 _dafny.MultiSet
			_ = _451_v30
			_451_v30 = _dafny.MultiSetOf((_421_v4).Dtor_cf16())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
			_ = _index74
			var _rhs46 _dafny.Map = _448_v27
			_ = _rhs46
			var _rhs47 bool = ((_dafny.Zero).Minus((_450_v29).Dtor_cf5())).Cmp(_dafny.IntOfInt64(431)) == 0
			_ = _rhs47
			var _rhs48 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("atwdm")
			_ = _rhs48
			var _rhs49 _dafny.Int = ((func() _dafny.Int {
				if (_451_v30).Contains(true) {
					return (_451_v30).Multiplicity(true)
				}
				return _417_v0
			})()).Times(_417_v0)
			_ = _rhs49
			var _lhs31 _dafny.Array = _444_v23
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
			_ = _lhs32
			_448_v27 = _rhs46
			(_lhs31).ArraySet1(_rhs47, (_lhs32).Int())
			_446_v25 = _rhs48
			_417_v0 = _rhs49
			var _452_v31 D1
			_ = _452_v31
			_452_v31 = Companion_D1_.Create_DC2_(true)
			var _source7 D1 = _452_v31
			_ = _source7
			if _source7.Is_DC3() {
				var _453___mcc_h0 _dafny.Int = _source7.Get_().(D1_DC3).Cf5
				_ = _453___mcc_h0
				var _454_cf5 _dafny.Int = _453___mcc_h0
				_ = _454_cf5
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _index75
				(_444_v23).ArraySet1((_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), (_index75).Int())
				var _455_v32 _dafny.Array
				_ = _455_v32
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_8
				var _nw75 _dafny.Array
				_ = _nw75
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw75 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Int = func(_456_i2 _dafny.Int) _dafny.Int {
						return (_456_i2).Times(_dafny.IntOfInt64(-103))
					}
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw75 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw75).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw75).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_455_v32 = _nw75
				_455_v32 = _455_v32
				var _457_v33 _dafny.Map
				_ = _457_v33
				_457_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(650), _417_v0)
				var _458_v34 _dafny.Map
				_ = _458_v34
				_458_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), false)
				var _459_v35 D1
				_ = _459_v35
				_459_v35 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_417_v0), _445_v24, (_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), _417_v0)
				var _rhs50 _dafny.Int = _417_v0
				_ = _rhs50
				var _rhs51 _dafny.Int = _dafny.IntOfInt64(754)
				_ = _rhs51
				var _rhs52 _dafny.Int = ((func() _dafny.Int {
					if (_457_v33).Contains((_458_v34).Cardinality()) {
						return (_457_v33).Get((_458_v34).Cardinality()).(_dafny.Int)
					}
					return _417_v0
				})()).Minus(_454_cf5)
				_ = _rhs52
				var _rhs53 _dafny.Int = (_459_v35).Dtor_cf9()
				_ = _rhs53
				_454_cf5 = _rhs50
				_454_cf5 = _rhs51
				_454_cf5 = _rhs52
				_417_v0 = _rhs53
				var _460_v36 _dafny.Sequence
				_ = _460_v36
				_460_v36 = _dafny.SeqOf(_444_v23, _444_v23)
				var _461_v37 _dafny.Set
				_ = _461_v37
				_461_v37 = _dafny.SetOf(_417_v0, _417_v0, _417_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_449_v28, (Companion_Default___.SafeIndex(_417_v0, _dafny.IntOfUint32((_449_v28).Cardinality()))).Uint32(), (_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool))).Cardinality()))
				var _462_v38 _dafny.Set
				_ = _462_v38
				_462_v38 = _dafny.SetOf((_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), (_421_v4).Dtor_cf16())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _index76
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _index77
				var _rhs54 bool = !((func() _dafny.Set {
					if _420_v3 {
						return _461_v37
					}
					return _461_v37
				})()).Contains((_dafny.IntOfInt64(423)).Plus((_457_v33).Cardinality()))
				_ = _rhs54
				var _rhs55 _dafny.Int = (_417_v0).Plus(_417_v0)
				_ = _rhs55
				var _rhs56 bool = ((_462_v38).Union(_462_v38)).IsSubsetOf(_dafny.SetOf(_420_v3))
				_ = _rhs56
				var _rhs57 _dafny.Sequence = _460_v36
				_ = _rhs57
				var _lhs33 _dafny.Array = _444_v23
				_ = _lhs33
				var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _lhs34
				var _lhs35 _dafny.Array = _444_v23
				_ = _lhs35
				var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _lhs36
				(_lhs33).ArraySet1(_rhs54, (_lhs34).Int())
				_417_v0 = _rhs55
				(_lhs35).ArraySet1(_rhs56, (_lhs36).Int())
				_460_v36 = _rhs57
			} else if _source7.Is_DC4() {
				var _463___mcc_h1 _dafny.Int = _source7.Get_().(D1_DC4).Cf6
				_ = _463___mcc_h1
				var _464___mcc_h2 _dafny.CodePoint = _source7.Get_().(D1_DC4).Cf7
				_ = _464___mcc_h2
				var _465___mcc_h3 bool = _source7.Get_().(D1_DC4).Cf8
				_ = _465___mcc_h3
				var _466___mcc_h4 _dafny.Int = _source7.Get_().(D1_DC4).Cf9
				_ = _466___mcc_h4
				var _467_cf9 _dafny.Int = _466___mcc_h4
				_ = _467_cf9
				var _468_cf8 bool = _465___mcc_h3
				_ = _468_cf8
				var _469_cf7 _dafny.CodePoint = _464___mcc_h2
				_ = _469_cf7
				var _470_cf6 _dafny.Int = _463___mcc_h1
				_ = _470_cf6
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _index78
				(_444_v23).ArraySet1(!(!(_468_cf8)), (_index78).Int())
				var _471_v39 _dafny.Array
				_ = _471_v39
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw76
				_471_v39 = _nw76
				var _472_v40 _dafny.Map
				_ = _472_v40
				_472_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), _420_v3)
				var _473_v41 _dafny.Array
				_ = _473_v41
				var _nwElement0_17 _dafny.Map = _472_v40
				_ = _nwElement0_17
				var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(21))
				_ = _nw77
				(_nw77).ArraySet1(_nwElement0_17, 0)
				(_nw77).ArraySet1(_472_v40, 1)
				(_nw77).ArraySet1(_472_v40, 2)
				(_nw77).ArraySet1(_472_v40, 3)
				(_nw77).ArraySet1(_472_v40, 4)
				(_nw77).ArraySet1(_472_v40, 5)
				(_nw77).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v3, (_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool)), 6)
				(_nw77).ArraySet1(_472_v40, 7)
				(_nw77).ArraySet1(_472_v40, 8)
				(_nw77).ArraySet1(_472_v40, 9)
				(_nw77).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), (_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool)), 10)
				(_nw77).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_420_v3), _420_v3), 11)
				(_nw77).ArraySet1((_472_v40).Update(_420_v3, _468_cf8), 12)
				(_nw77).ArraySet1(_472_v40, 13)
				(_nw77).ArraySet1(_472_v40, 14)
				(_nw77).ArraySet1((_472_v40).Update(_420_v3, false), 15)
				(_nw77).ArraySet1(Companion_Default___.Fm15(_470_cf6, _468_cf8, _469_cf7, _467_cf9, globalState), 16)
				(_nw77).ArraySet1(_472_v40, 17)
				(_nw77).ArraySet1(_472_v40, 18)
				(_nw77).ArraySet1(_472_v40, 19)
				(_nw77).ArraySet1(_472_v40, 20)
				_473_v41 = _nw77
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_471_v39), 0))
				_ = _index79
				(_471_v39).ArraySet1(_473_v41, (_index79).Int())
				var _474_v42 _dafny.Array
				_ = _474_v42
				var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw78
				_474_v42 = _nw78
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))
				_ = _index80
				(_474_v42).ArraySet1(_dafny.IntOfUint32((_446_v25).Cardinality()), (_index80).Int())
				var _475_v43 _dafny.Set
				_ = _475_v43
				_475_v43 = _dafny.SetOf((_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), _468_cf8)
				var _476_v44 _dafny.Sequence
				_ = _476_v44
				_476_v44 = _dafny.SeqOf(_dafny.SetOf(_468_cf8, (_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool)), _475_v43)
				var _477_v45 _dafny.Array
				_ = _477_v45
				var _nwElement0_18 _dafny.Set = (_475_v43).Intersection(_475_v43)
				_ = _nwElement0_18
				var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(12))
				_ = _nw79
				(_nw79).ArraySet1(_nwElement0_18, 0)
				(_nw79).ArraySet1((_476_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_449_v28).Cardinality()), _dafny.IntOfUint32((_476_v44).Cardinality()))).Uint32()).(_dafny.Set), 1)
				(_nw79).ArraySet1(_475_v43, 2)
				(_nw79).ArraySet1((_475_v43).Intersection(_475_v43), 3)
				(_nw79).ArraySet1((_475_v43).Difference(_475_v43), 4)
				(_nw79).ArraySet1(_475_v43, 5)
				(_nw79).ArraySet1((Companion_Default___.Fm16(_421_v4, globalState)).Union(_475_v43), 6)
				(_nw79).ArraySet1(_475_v43, 7)
				(_nw79).ArraySet1(_475_v43, 8)
				(_nw79).ArraySet1(_475_v43, 9)
				(_nw79).ArraySet1(_475_v43, 10)
				(_nw79).ArraySet1((_475_v43).Difference(_dafny.SetOf(_420_v3)), 11)
				_477_v45 = _nw79
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_477_v45), 0))
				_ = _index81
				(_477_v45).ArraySet1(_475_v43, (_index81).Int())
				var _478_v46 _dafny.Map
				_ = _478_v46
				_478_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_467_cf9), _473_v41)
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_471_v39), 0))
				_ = _index82
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))
				_ = _index83
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_477_v45), 0))
				_ = _index84
				var _rhs58 _dafny.Array = (func() _dafny.Array {
					if (_478_v46).Contains(_dafny.IntOfInt64(-336)) {
						return (_478_v46).Get(_dafny.IntOfInt64(-336)).(_dafny.Array)
					}
					return _473_v41
				})()
				_ = _rhs58
				var _rhs59 _dafny.Int = _417_v0
				_ = _rhs59
				var _rhs60 _dafny.Set = Companion_Default___.Fm16(_421_v4, globalState)
				_ = _rhs60
				var _lhs37 _dafny.Array = _471_v39
				_ = _lhs37
				var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_471_v39), 0))
				_ = _lhs38
				var _lhs39 _dafny.Array = _474_v42
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))
				_ = _lhs40
				var _lhs41 _dafny.Array = _477_v45
				_ = _lhs41
				var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_477_v45), 0))
				_ = _lhs42
				(_lhs37).ArraySet1(_rhs58, (_lhs38).Int())
				(_lhs39).ArraySet1(_rhs59, (_lhs40).Int())
				(_lhs41).ArraySet1(_rhs60, (_lhs42).Int())
				var _479_v47 _dafny.Map
				_ = _479_v47
				_479_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_469_cf7, false)
				_417_v0 = ((_dafny.IntOfUint32((_446_v25).Cardinality())).Times(_470_cf6)).Times((func() _dafny.Int {
					if (func() bool {
						if (_479_v47).Contains(_469_cf7) {
							return (_479_v47).Get(_469_cf7).(bool)
						}
						return (_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool)
					})() {
						return (_474_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))).Int()).(_dafny.Int)
					}
					return (_474_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))).Int()).(_dafny.Int)
				})())
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))
				_ = _index85
				var _rhs61 _dafny.Int = (_474_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))).Int()).(_dafny.Int)
				_ = _rhs61
				var _rhs62 _dafny.Int = (func() _dafny.Int {
					if false {
						return Companion_Default___.SafeModuloInt(_470_cf6, (_474_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))).Int()).(_dafny.Int))
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ppyxfavq")).Cardinality())
				})()
				_ = _rhs62
				var _lhs43 _dafny.Array = _474_v42
				_ = _lhs43
				var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_474_v42), 0))
				_ = _lhs44
				_470_cf6 = _rhs61
				(_lhs43).ArraySet1(_rhs62, (_lhs44).Int())
			} else {
				var _480___mcc_h5 bool = _source7.Get_().(D1_DC2).Cf4
				_ = _480___mcc_h5
				var _481_cf4 bool = _480___mcc_h5
				_ = _481_cf4
				var _482_v48 *C2
				_ = _482_v48
				var _nw80 *C2 = New_C2_()
				_ = _nw80
				_nw80.Ctor__((func() _dafny.Sequence {
					if _420_v3 {
						return _446_v25
					}
					return _446_v25
				})(), _446_v25)
				_482_v48 = _nw80
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _index86
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _index87
				var _rhs63 bool = _420_v3
				_ = _rhs63
				var _rhs64 bool = (_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool)
				_ = _rhs64
				var _rhs65 bool = ((_417_v0).Minus(_417_v0)).Cmp(_417_v0) != 0
				_ = _rhs65
				var _rhs66 bool = !(_420_v3) || (_dafny.Companion_Sequence_.IsProperPrefixOf((_482_v48).F9(), _446_v25))
				_ = _rhs66
				var _lhs45 _dafny.Array = _444_v23
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _lhs46
				var _lhs47 _dafny.Array = _444_v23
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _lhs48
				_481_cf4 = _rhs63
				_481_cf4 = _rhs64
				(_lhs45).ArraySet1(_rhs65, (_lhs46).Int())
				(_lhs47).ArraySet1(_rhs66, (_lhs48).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))
				_ = _index88
				(_444_v23).ArraySet1((_444_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_444_v23), 0))).Int()).(bool), (_index88).Int())
				var _483_v49 _dafny.Array
				_ = _483_v49
				var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw81
				_483_v49 = _nw81
				var _484_v50 _dafny.Array
				_ = _484_v50
				var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw82
				_484_v50 = _nw82
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_483_v49), 0))
				_ = _index89
				(_483_v49).ArraySet1(_484_v50, (_index89).Int())
				var _485_v51 _dafny.Sequence
				_ = _485_v51
				_485_v51 = _dafny.SeqOf(_484_v50, _484_v50, _484_v50)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_483_v49), 0))
				_ = _index90
				(_483_v49).ArraySet1((_485_v51).Select((Companion_Default___.SafeIndex(_417_v0, _dafny.IntOfUint32((_485_v51).Cardinality()))).Uint32()).(_dafny.Array), (_index90).Int())
			}
		}
		var _486_v52 _dafny.Set
		_ = _486_v52
		_486_v52 = _dafny.SetOf(_420_v3)
		var _487_i3 _dafny.Int
		_ = _487_i3
		_487_i3 = _dafny.Zero
		{
			for (_486_v52).IsProperSubsetOf(_486_v52) {
				{
					if (_487_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_487_i3 = (_487_i3).Plus(_dafny.One)
					var _488_v53 _dafny.Map
					_ = _488_v53
					_488_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v3, _dafny.IntOfInt64(-225))
					_488_v53 = _488_v53
					var _489_v54 _dafny.Sequence
					_ = _489_v54
					_489_v54 = _dafny.UnicodeSeqOfUtf8Bytes("pybwqy")
					var _490_v55 _dafny.Map
					_ = _490_v55
					_490_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_417_v0, _489_v54)
					var _491_v56 T0
					_ = _491_v56
					var _nw83 *C2 = New_C2_()
					_ = _nw83
					_nw83.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("poqcmavv"), (func() _dafny.Sequence {
						if (_490_v55).Contains(_417_v0) {
							return (_490_v55).Get(_417_v0).(_dafny.Sequence)
						}
						return _489_v54
					})()), _489_v54)
					_491_v56 = _nw83
					var _492_v57 _dafny.Map
					_ = _492_v57
					_492_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v3, _420_v3)
					_491_v56 = (func() T0 {
						if (func() bool {
							if (_492_v57).Contains(!(_420_v3)) {
								return (_492_v57).Get(!(_420_v3)).(bool)
							}
							return false
						})() {
							return _491_v56
						}
						return _491_v56
					})()
					var _493_v58 _dafny.Map
					_ = _493_v58
					_493_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v3, _421_v4)
					_417_v0 = Companion_Default___.SafeDivisionInt(_417_v0, (_493_v58).Cardinality())
					_492_v57 = (_492_v57).Update(_420_v3, _420_v3)
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _494_v59 _dafny.MultiSet
		_ = _494_v59
		_494_v59 = _dafny.MultiSetOf(_420_v3, _420_v3)
		var _495_v60 _dafny.Sequence
		_ = _495_v60
		_495_v60 = _dafny.SeqOf(_420_v3, _420_v3)
		var _496_v61 *C1
		_ = _496_v61
		var _nw84 *C1 = New_C1_()
		_ = _nw84
		_nw84.Ctor__(_494_v59, (_dafny.IntOfUint32((_495_v60).Cardinality())).Times((Companion_Default___.Fm16(_421_v4, globalState)).Cardinality()))
		_496_v61 = _nw84
		if _420_v3 {
			var _497_v62 _dafny.Array
			_ = _497_v62
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw85
			_497_v62 = _nw85
			var _498_v63 _dafny.MultiSet
			_ = _498_v63
			_498_v63 = _dafny.MultiSetOf(_497_v62, _497_v62)
			_498_v63 = _498_v63
			var _499_v64 _dafny.CodePoint
			_ = _499_v64
			_499_v64 = _dafny.CodePoint('v')
			var _500_v65 _dafny.Sequence
			_ = _500_v65
			_500_v65 = _dafny.UnicodeSeqOfUtf8Bytes("ned")
			var _rhs67 _dafny.CodePoint = Companion_Default___.Fm30(_500_v65, _496_v61.F11, _496_v61.F11, globalState)
			_ = _rhs67
			var _rhs68 _dafny.CodePoint = _499_v64
			_ = _rhs68
			var _rhs69 bool = !(_420_v3)
			_ = _rhs69
			var _rhs70 _dafny.Int = _496_v61.F11
			_ = _rhs70
			var _lhs49 *C1 = _496_v61
			_ = _lhs49
			_499_v64 = _rhs67
			_499_v64 = _rhs68
			_420_v3 = _rhs69
			_lhs49.F11 = _rhs70
			var _501_v66 _dafny.Sequence
			_ = _501_v66
			_501_v66 = _dafny.SeqOf(_496_v61.F11)
			_501_v66 = _dafny.Companion_Sequence_.Concatenate(_501_v66, _dafny.Companion_Sequence_.Concatenate(_501_v66, _501_v66))
			(_496_v61).F11 = _496_v61.F11
			(_496_v61).F11 = (_496_v61.F11).Times(Companion_Default___.SafeModuloInt(_496_v61.F11, _417_v0))
		} else {
			var _502_v67 _dafny.Sequence
			_ = _502_v67
			_502_v67 = _dafny.UnicodeSeqOfUtf8Bytes("bwgipadpu")
			(_496_v61).F11 = Companion_Default___.Fm14(_417_v0, _dafny.IntOfUint32((_502_v67).Cardinality()), globalState)
			var _503_v68 _dafny.CodePoint
			_ = _503_v68
			_503_v68 = _dafny.CodePoint('a')
			var _504_v69 _dafny.Map
			_ = _504_v69
			_504_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_502_v67, _502_v67), (Companion_Default___.SafeIndex(_417_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_502_v67, _502_v67)).Cardinality()))).Uint32(), _503_v68), _420_v3)
			_504_v69 = (_504_v69).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(962))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_505_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), _420_v3)
			if (func() bool {
				if !(_420_v3) {
					return !(_dafny.Companion_Sequence_.Contains(_502_v67, _503_v68))
				}
				return !(!((Companion_Default___.Fm31(_dafny.IntOfUint32((_502_v67).Cardinality()), globalState)).IsDisjointFrom(_418_v1)))
			})() {
				var _506_v70 *C0
				_ = _506_v70
				var _nw86 *C0 = New_C0_()
				_ = _nw86
				_nw86.Ctor__()
				_506_v70 = _nw86
				var _507_v71 *C1
				_ = _507_v71
				var _nw87 *C1 = New_C1_()
				_ = _nw87
				_nw87.Ctor__((_496_v61).Fm0(_420_v3, _420_v3, (_486_v52).Cardinality(), _dafny.IntOfInt64(-609), globalState), _496_v61.F11)
				_507_v71 = _nw87
				var _508_v72 _dafny.Map
				_ = _508_v72
				_508_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_503_v68, _496_v61.F11)
				var _509_v73 _dafny.Set
				_ = _509_v73
				_509_v73 = _dafny.SetOf(_502_v67)
				var _510_v74 D4
				_ = _510_v74
				_510_v74 = Companion_D4_.Create_DC10_((_496_v61).Fm21(_508_v72, (_509_v73).Cardinality(), _dafny.IntOfInt64(324), globalState))
				var _511_v75 D4
				_ = _511_v75
				_511_v75 = Companion_D4_.Create_DC11_(_510_v74)
				_511_v75 = _511_v75
				var _512_v76 _dafny.Array
				_ = _512_v76
				var _nw88 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
				_ = _nw88
				_512_v76 = _nw88
				var _513_v77 T0
				_ = _513_v77
				var _nw89 *C2 = New_C2_()
				_ = _nw89
				_nw89.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("gcchob"), _dafny.UnicodeSeqOfUtf8Bytes("f"))
				_513_v77 = _nw89
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_512_v76), 0))
				_ = _index91
				(_512_v76).ArraySet1(_513_v77, (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_512_v76), 0))
				_ = _index92
				(_512_v76).ArraySet1(_513_v77, (_index92).Int())
				var _514_v78 _dafny.Map
				_ = _514_v78
				_514_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v61.F11, _496_v61.F11)
				var _515_v79 _dafny.Map
				_ = _515_v79
				_515_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_514_v78, _486_v52)
				var _516_v80 _dafny.Set
				_ = _516_v80
				_516_v80 = _dafny.SetOf(_496_v61.F11, ((_496_v61).F10()).Cardinality())
				var _517_v81 D9
				_ = _517_v81
				_517_v81 = Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_495_v60).Cardinality()), (_516_v80).Cardinality()), _dafny.SetOf(_420_v3, _420_v3)))
				var _518_v82 _dafny.Map
				_ = _518_v82
				_518_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_507_v71.F11, _515_v79)
				var _519_v83 _dafny.Map
				_ = _519_v83
				_519_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v3, _516_v80)
				var _520_v84 _dafny.Map
				_ = _520_v84
				_520_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v3, (func() _dafny.Set {
					if (_519_v83).Contains(true) {
						return (_519_v83).Get(true).(_dafny.Set)
					}
					return _516_v80
				})())
				var _521_v86 _dafny.Map
				_ = _521_v86
				_521_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _486_v52)
				var _522_v87 _dafny.Sequence
				_ = _522_v87
				_522_v87 = _dafny.SeqOf(_515_v79)
				var _523_v89 _dafny.Set
				_ = _523_v89
				_523_v89 = _dafny.SetOf(_514_v78)
				var _pat_let_tv4 = _515_v79
				_ = _pat_let_tv4
				var _524_v90 _dafny.Array
				_ = _524_v90
				var _nwElement0_19 _dafny.Map = _515_v79
				_ = _nwElement0_19
				var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(25))
				_ = _nw90
				(_nw90).ArraySet1(_nwElement0_19, 0)
				(_nw90).ArraySet1(_515_v79, 1)
				(_nw90).ArraySet1((_517_v81).Dtor_cf28(), 2)
				(_nw90).ArraySet1(_515_v79, 3)
				(_nw90).ArraySet1(_515_v79, 4)
				(_nw90).ArraySet1((_515_v79).Merge(_515_v79), 5)
				(_nw90).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_514_v78, _486_v52), 6)
				(_nw90).ArraySet1((_515_v79).Merge(_515_v79), 7)
				(_nw90).ArraySet1((func() _dafny.Map {
					if _420_v3 {
						return _515_v79
					}
					return _515_v79
				})(), 8)
				(_nw90).ArraySet1((_515_v79).Merge(_515_v79), 9)
				(_nw90).ArraySet1((func() _dafny.Map {
					if (_518_v82).Contains(((func() _dafny.Set {
						if (_519_v83).Contains(_420_v3) {
							return (_519_v83).Get(_420_v3).(_dafny.Set)
						}
						return _dafny.SetOf(_496_v61.F11, Companion_Default___.Fm19(_420_v3, _496_v61.F11, globalState))
					})()).Cardinality()) {
						return (_518_v82).Get(((func() _dafny.Set {
							if (_519_v83).Contains(_420_v3) {
								return (_519_v83).Get(_420_v3).(_dafny.Set)
							}
							return _dafny.SetOf(_496_v61.F11, Companion_Default___.Fm19(_420_v3, _496_v61.F11, globalState))
						})()).Cardinality()).(_dafny.Map)
					}
					return _515_v79
				})(), 10)
				(_nw90).ArraySet1(_515_v79, 11)
				(_nw90).ArraySet1(_515_v79, 12)
				(_nw90).ArraySet1(_515_v79, 13)
				(_nw90).ArraySet1((Companion_Default___.Fm32(_420_v3, globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_417_v0, _496_v61.F11), _486_v52)), 14)
				(_nw90).ArraySet1(_515_v79, 15)
				(_nw90).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_514_v78, _486_v52), 16)
				(_nw90).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_514_v78, _486_v52)).Update((_514_v78).Update((_dafny.Zero).Minus(_417_v0), _496_v61.F11), _486_v52)).Merge(_515_v79), 17)
				(_nw90).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if ((_496_v61).F10()).Contains(_420_v3) {
						return ((_496_v61).F10()).Multiplicity(_420_v3)
					}
					return (func() _dafny.Map {
						var _coll16 = _dafny.NewMapBuilder()
						_ = _coll16
						for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(187), _dafny.IntOfInt64(166))); ; {
							_compr_16, _ok19 := _iter19()
							if !_ok19 {
								break
							}
							var _525_v85 _dafny.Int
							_525_v85 = interface{}(_compr_16).(_dafny.Int)
							if ((_dafny.IntOfInt64(187)).Cmp(_525_v85) <= 0) && ((_525_v85).Cmp(_dafny.IntOfInt64(166)) < 0) {
								_coll16.Add(Companion_Default___.SafeModuloInt(_525_v85, _496_v61.F11), _503_v68)
							}
						}
						return _coll16.ToMap()
					}()).Cardinality()
				})(), _417_v0), (func() _dafny.Set {
					if (_521_v86).Contains(_420_v3) {
						return (_521_v86).Get(_420_v3).(_dafny.Set)
					}
					return _486_v52
				})()), 18)
				(_nw90).ArraySet1(_515_v79, 19)
				(_nw90).ArraySet1((_515_v79).Merge((_522_v87).Select((Companion_Default___.SafeIndex(_496_v61.F11, _dafny.IntOfUint32((_522_v87).Cardinality()))).Uint32()).(_dafny.Map)), 20)
				(_nw90).ArraySet1(func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter20 := _dafny.Iterate((_523_v89).Elements()); ; {
						_compr_17, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _526_v88 _dafny.Map
						_526_v88 = interface{}(_compr_17).(_dafny.Map)
						if (_523_v89).Contains(_526_v88) {
							_coll17.Add(_526_v88, _486_v52)
						}
					}
					return _coll17.ToMap()
				}(), 21)
				(_nw90).ArraySet1(((func(_pat_let6_0 D9) D9 {
					return func(_527_dt__update__tmp_h0 D9) D9 {
						return func(_pat_let7_0 _dafny.Map) D9 {
							return func(_528_dt__update_hcf28_h0 _dafny.Map) D9 {
								return Companion_D9_.Create_DC18_(_528_dt__update_hcf28_h0)
							}(_pat_let7_0)
						}(_pat_let_tv4)
					}(_pat_let6_0)
				}(Companion_D9_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_514_v78, _486_v52)))).Dtor_cf28()).Merge(Companion_Default___.Fm32(_420_v3, globalState)), 22)
				(_nw90).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_514_v78, Companion_Default___.Fm16(_421_v4, globalState)), 23)
				(_nw90).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_514_v78).Update(_417_v0, _dafny.IntOfInt64(213)), _486_v52), 24)
				_524_v90 = _nw90
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_524_v90), 0))
				_ = _index93
				(_524_v90).ArraySet1(_515_v79, (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_524_v90), 0))
				_ = _index94
				var _rhs71 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_514_v78).Contains(_417_v0) {
						return (_514_v78).Get(_417_v0).(_dafny.Int)
					}
					return _496_v61.F11
				})(), Companion_Default___.SafeDivisionInt(_496_v61.F11, _507_v71.F11))
				_ = _rhs71
				var _rhs72 _dafny.Map = (_515_v79).Merge(Companion_Default___.Fm32(false, globalState))
				_ = _rhs72
				var _lhs50 *C1 = _496_v61
				_ = _lhs50
				var _lhs51 _dafny.Array = _524_v90
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_524_v90), 0))
				_ = _lhs52
				_lhs50.F11 = _rhs71
				(_lhs51).ArraySet1(_rhs72, (_lhs52).Int())
			} else {
				var _529_v91 _dafny.Array
				_ = _529_v91
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
				_ = _nw91
				_529_v91 = _nw91
				var _530_v92 _dafny.Map
				_ = _530_v92
				_530_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v61.F11, _529_v91)
				var _531_v93 _dafny.Sequence
				_ = _531_v93
				_531_v93 = _dafny.SeqOf(_529_v91, _529_v91)
				var _532_v94 _dafny.Map
				_ = _532_v94
				_532_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _496_v61.F11)
				var _533_v95 _dafny.Array
				_ = _533_v95
				var _nwElement0_20 _dafny.Array = _529_v91
				_ = _nwElement0_20
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(11))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_20, 0)
				(_nw92).ArraySet1(_529_v91, 1)
				(_nw92).ArraySet1(_529_v91, 2)
				(_nw92).ArraySet1(_529_v91, 3)
				(_nw92).ArraySet1(_529_v91, 4)
				(_nw92).ArraySet1(_529_v91, 5)
				(_nw92).ArraySet1(_529_v91, 6)
				(_nw92).ArraySet1(_529_v91, 7)
				(_nw92).ArraySet1(_529_v91, 8)
				(_nw92).ArraySet1((func() _dafny.Array {
					if (_530_v92).Contains(_496_v61.F11) {
						return (_530_v92).Get(_496_v61.F11).(_dafny.Array)
					}
					return (_531_v93).Select((Companion_Default___.SafeIndex((_532_v94).Cardinality(), _dafny.IntOfUint32((_531_v93).Cardinality()))).Uint32()).(_dafny.Array)
				})(), 9)
				(_nw92).ArraySet1(_529_v91, 10)
				_533_v95 = _nw92
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_533_v95), 0))
				_ = _index95
				(_533_v95).ArraySet1((_531_v93).Select((Companion_Default___.SafeIndex(_496_v61.F11, _dafny.IntOfUint32((_531_v93).Cardinality()))).Uint32()).(_dafny.Array), (_index95).Int())
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_533_v95), 0))
				_ = _index96
				(_533_v95).ArraySet1((_531_v93).Select((Companion_Default___.SafeIndex(_496_v61.F11, _dafny.IntOfUint32((_531_v93).Cardinality()))).Uint32()).(_dafny.Array), (_index96).Int())
				var _534_v96 _dafny.Map
				_ = _534_v96
				_534_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_496_v61).F10(), _417_v0)
				var _535_v97 _dafny.Map
				_ = _535_v97
				_535_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_534_v96).Update((_496_v61).F10(), _496_v61.F11)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_494_v59, _417_v0)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_502_v67).Cardinality())))
				(_496_v61).F11 = (func() _dafny.Int {
					if (_535_v97).Contains(_534_v96) {
						return (_535_v97).Get(_534_v96).(_dafny.Int)
					}
					return _496_v61.F11
				})()
				_502_v67 = _502_v67
				(_496_v61).F11 = (func() _dafny.Int {
					if !(_420_v3) || (_420_v3) {
						return _496_v61.F11
					}
					return _dafny.IntOfInt64(393)
				})()
				var _536_v98 _dafny.Array
				_ = _536_v98
				var _nwElement0_21 bool = _420_v3
				_ = _nwElement0_21
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(3))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_21, 0)
				(_nw93).ArraySet1(_420_v3, 1)
				(_nw93).ArraySet1(_420_v3, 2)
				_536_v98 = _nw93
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_536_v98), 0))
				_ = _index97
				(_536_v98).ArraySet1(_420_v3, (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_536_v98), 0))
				_ = _index98
				(_536_v98).ArraySet1((_486_v52).IsSubsetOf(_486_v52), (_index98).Int())
			}
			var _537_v99 *C0
			_ = _537_v99
			var _nw94 *C0 = New_C0_()
			_ = _nw94
			_nw94.Ctor__()
			_537_v99 = _nw94
			_417_v0 = (_417_v0).Plus(((_dafny.MultiSetOf(_420_v3)).Union(((_496_v61).F10()).Update(false, Companion_Default___.Abs(_dafny.IntOfInt64(506))))).Cardinality())
		}
		r0 = _420_v3
		var _538_v100 D2
		_ = _538_v100
		_538_v100 = Companion_D2_.Create_DC6_(_496_v61.F11, _420_v3)
		var _539_v101 _dafny.Map
		_ = _539_v101
		_539_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v61.F11, _dafny.MultiSetOf((_538_v100).Dtor_cf12(), false))
		r1 = (_539_v101).Update(Companion_Default___.Fm14(_496_v61.F11, (_dafny.Zero).Minus((_dafny.Zero).Minus(_417_v0)), globalState), (_496_v61).F10())
		return r0, r1
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	F7 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this.F7 = _dafny.Zero
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

func (_this *C4) Ctor__(f7 _dafny.Int) {
	{
		(_this).F7 = f7
	}
}
func (_this *C4) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return ((_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(false, false, false, true, true))).Difference(_dafny.MultiSetOf(!(!(!(true)))))
	}
}
func (_this *C4) M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _540_v0 *C3
		_ = _540_v0
		var _nw95 *C3 = New_C3_()
		_ = _nw95
		_nw95.Ctor__()
		_540_v0 = _nw95
		r0 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_this.F7), _dafny.IntOfInt64(471))
		var _541_v1 _dafny.Sequence
		_ = _541_v1
		_541_v1 = _dafny.UnicodeSeqOfUtf8Bytes("jmcwwhb")
		var _542_v2 *C2
		_ = _542_v2
		var _nw96 *C2 = New_C2_()
		_ = _nw96
		_nw96.Ctor__(_541_v1, _541_v1)
		_542_v2 = _nw96
		var _543_v3 bool
		_ = _543_v3
		_543_v3 = true
		var _544_v4 _dafny.Sequence
		_ = _544_v4
		_544_v4 = _dafny.SeqOf(true, true, _543_v3)
		r0 = (func() bool {
			if _543_v3 {
				return _543_v3
			}
			return (_544_v4).Select((Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_544_v4).Cardinality()))).Uint32()).(bool)
		})()
		var _545_v5 _dafny.Array
		_ = _545_v5
		var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
		_ = _nw97
		_545_v5 = _nw97
		var _546_v6 D4
		_ = _546_v6
		_546_v6 = Companion_D4_.Create_DC9_(_545_v5)
		_546_v6 = (func() D4 {
			if _543_v3 {
				return _546_v6
			}
			return _546_v6
		})()
		var _547_i0 _dafny.Int
		_ = _547_i0
		_547_i0 = _dafny.Zero
		{
			for true {
				{
					if (_547_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_547_i0 = (_547_i0).Plus(_dafny.One)
					_541_v1 = (_542_v2).F8()
					var _rhs73 _dafny.Int = (_dafny.Zero).Minus((_this.F7).Times(Companion_Default___.Fm14(_this.F7, _this.F7, globalState)))
					_ = _rhs73
					var _rhs74 _dafny.Int = _dafny.IntOfInt64(907)
					_ = _rhs74
					var _rhs75 bool = (false) || (_543_v3)
					_ = _rhs75
					var _lhs53 *C4 = _this
					_ = _lhs53
					var _lhs54 *C4 = _this
					_ = _lhs54
					_lhs53.F7 = _rhs73
					_lhs54.F7 = _rhs74
					r0 = _rhs75
					var _548_v7 _dafny.Array
					_ = _548_v7
					var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
					_ = _nw98
					_548_v7 = _nw98
					var _549_v8 _dafny.Array
					_ = _549_v8
					var _nwElement0_22 _dafny.Array = _545_v5
					_ = _nwElement0_22
					var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(26))
					_ = _nw99
					(_nw99).ArraySet1(_nwElement0_22, 0)
					(_nw99).ArraySet1(_545_v5, 1)
					(_nw99).ArraySet1(_545_v5, 2)
					(_nw99).ArraySet1(_545_v5, 3)
					(_nw99).ArraySet1(_545_v5, 4)
					(_nw99).ArraySet1(_545_v5, 5)
					(_nw99).ArraySet1(_545_v5, 6)
					(_nw99).ArraySet1(_545_v5, 7)
					(_nw99).ArraySet1(_545_v5, 8)
					(_nw99).ArraySet1(_545_v5, 9)
					(_nw99).ArraySet1(_545_v5, 10)
					(_nw99).ArraySet1(_545_v5, 11)
					(_nw99).ArraySet1(_545_v5, 12)
					(_nw99).ArraySet1(_545_v5, 13)
					(_nw99).ArraySet1(_545_v5, 14)
					(_nw99).ArraySet1(_545_v5, 15)
					(_nw99).ArraySet1(_545_v5, 16)
					(_nw99).ArraySet1(_545_v5, 17)
					(_nw99).ArraySet1(_545_v5, 18)
					(_nw99).ArraySet1(_545_v5, 19)
					(_nw99).ArraySet1(_545_v5, 20)
					(_nw99).ArraySet1(_545_v5, 21)
					(_nw99).ArraySet1(_545_v5, 22)
					(_nw99).ArraySet1(_545_v5, 23)
					(_nw99).ArraySet1(_545_v5, 24)
					(_nw99).ArraySet1(_545_v5, 25)
					_549_v8 = _nw99
					var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_548_v7), 0))
					_ = _index99
					(_548_v7).ArraySet1(_549_v8, (_index99).Int())
					var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.ArrayLen((_548_v7), 0))
					_ = _index100
					(_548_v7).ArraySet1(_549_v8, (_index100).Int())
					var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_545_v5), 0))
					_ = _index101
					(_545_v5).ArraySet1(_this.F7, (_index101).Int())
					var _550_v9 _dafny.MultiSet
					_ = _550_v9
					_550_v9 = _dafny.MultiSetOf(_543_v3, _543_v3)
					var _551_v10 _dafny.MultiSet
					_ = _551_v10
					_551_v10 = _dafny.MultiSetOf(_this.F7, (func() _dafny.Int {
						if (_550_v9).Contains(_543_v3) {
							return (_550_v9).Multiplicity(_543_v3)
						}
						return _this.F7
					})())
					var _552_v11 _dafny.Sequence
					_ = _552_v11
					_552_v11 = _dafny.SeqOf(_541_v1)
					var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_545_v5), 0))
					_ = _index102
					var _rhs76 _dafny.Int = _dafny.IntOfInt64(613)
					_ = _rhs76
					var _rhs77 _dafny.Int = _this.F7
					_ = _rhs77
					var _rhs78 bool = _543_v3
					_ = _rhs78
					var _rhs79 _dafny.Int = Companion_Default___.Fm19((_551_v10).IsDisjointFrom(_551_v10), _dafny.IntOfUint32((_552_v11).Cardinality()), globalState)
					_ = _rhs79
					var _lhs55 *C4 = _this
					_ = _lhs55
					var _lhs56 _dafny.Array = _545_v5
					_ = _lhs56
					var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_545_v5), 0))
					_ = _lhs57
					var _lhs58 *C4 = _this
					_ = _lhs58
					_lhs55.F7 = _rhs76
					(_lhs56).ArraySet1(_rhs77, (_lhs57).Int())
					_543_v3 = _rhs78
					_lhs58.F7 = _rhs79
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _553_v12 _dafny.CodePoint
		_ = _553_v12
		_553_v12 = _dafny.CodePoint('i')
		r0 = (_542_v2).Fm17(_553_v12, _543_v3, _this.F7, _this.F7, globalState)
		var _554_v13 _dafny.Map
		_ = _554_v13
		_554_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, _dafny.MultiSetFromSeq(_544_v4))
		var _555_v14 _dafny.MultiSet
		_ = _555_v14
		_555_v14 = _dafny.MultiSetOf(_543_v3)
		r1 = ((_554_v13).Merge(_554_v13)).Merge((_554_v13).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, _555_v14)))
		return r0, r1
	}
}
func (_this *C4) M4(p0 bool, p1 _dafny.Array, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		r0 = p0
		var _hi4 _dafny.Int = _this.F7
		_ = _hi4
		for _556_i0 := p2; _556_i0.Cmp(_hi4) < 0; _556_i0 = _556_i0.Plus(_dafny.One) {
			var _557_v0 *C0
			_ = _557_v0
			var _nw100 *C0 = New_C0_()
			_ = _nw100
			_nw100.Ctor__()
			_557_v0 = _nw100
			var _558_v1 D7
			_ = _558_v1
			_558_v1 = Companion_D7_.Create_DC15_(p0, _this.F7, _556_i0, _this.F7, _557_v0)
			var _559_v2 _dafny.MultiSet
			_ = _559_v2
			_559_v2 = _dafny.MultiSetOf(p3)
			var _560_v3 *C1
			_ = _560_v3
			var _nw101 *C1 = New_C1_()
			_ = _nw101
			_nw101.Ctor__(_559_v2, _556_i0)
			_560_v3 = _nw101
			var _561_v4 _dafny.Map
			_ = _561_v4
			_561_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_560_v3, p3)
			var _562_v5 _dafny.Map
			_ = _562_v5
			_562_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_561_v4).Cardinality())
			var _563_v6 _dafny.Map
			_ = _563_v6
			_563_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0) == (!(p3)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_558_v1, _562_v5)).Cardinality())
			_562_v5 = _563_v6
			var _564_v7 _dafny.Array
			_ = _564_v7
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw102
			_564_v7 = _nw102
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_564_v7), 0))
			_ = _index103
			(_564_v7).ArraySet1(p3, (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_564_v7), 0))
			_ = _index104
			(_564_v7).ArraySet1((func() bool {
				if (func() bool {
					if p3 {
						return p3
					}
					return p3
				})() {
					return ((_560_v3).F10()).Contains(p0)
				}
				return (_this.F7).Cmp(_dafny.IntOfInt64(-700)) <= 0
			})(), (_index104).Int())
			var _565_v8 _dafny.Array
			_ = _565_v8
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_9
			var _nw103 _dafny.Array
			_ = _nw103
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw103 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Sequence = func(_566_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg24 _dafny.Int) interface{} {
							return coer24(arg24)
						}
					}(func(_567_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('s')
					}))
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw103 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw103).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw103).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_565_v8 = _nw103
			var _568_v9 _dafny.Sequence
			_ = _568_v9
			_568_v9 = _dafny.UnicodeSeqOfUtf8Bytes("muqj")
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_565_v8), 0))
			_ = _index105
			(_565_v8).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_568_v9, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("yifcb"), (Companion_Default___.SafeIndex(_556_i0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yifcb")).Cardinality()))).Uint32(), _dafny.CodePoint('f'))), (_index105).Int())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_565_v8), 0))
			_ = _index106
			(_565_v8).ArraySet1((func() _dafny.Sequence {
				if p0 {
					return _568_v9
				}
				return _dafny.Companion_Sequence_.Concatenate(_568_v9, _568_v9)
			})(), (_index106).Int())
			(_this).F7 = (p2).Times(p2)
		}
		var _569_v10 _dafny.Array
		_ = _569_v10
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_10
		var _nw104 _dafny.Array
		_ = _nw104
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw104 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Int = func(_570_i4 _dafny.Int) _dafny.Int {
				return (_570_i4).Minus((_dafny.Zero).Minus(_this.F7))
			}
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw104 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw104).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw104).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_569_v10 = _nw104
		for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_569_v10), 0))); ; {
			_guard_loop_3, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _571_i3 _dafny.Int
			_571_i3 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_571_i3).Sign() != -1) && ((_571_i3).Cmp(_dafny.ArrayLen((_569_v10), 0)) < 0)) {
				(_569_v10).ArraySet1(Companion_Default___.SafeDivisionInt(_571_i3, _dafny.IntOfInt64(-97)), (_571_i3).Int())
			}
		}
		var _572_v11 _dafny.Array
		_ = _572_v11
		var _nw105 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
		_ = _nw105
		_572_v11 = _nw105
		var _573_v12 _dafny.Sequence
		_ = _573_v12
		_573_v12 = _dafny.UnicodeSeqOfUtf8Bytes("jjbqnymx")
		var _574_v13 T0
		_ = _574_v13
		var _nw106 *C2 = New_C2_()
		_ = _nw106
		_nw106.Ctor__(_573_v12, _573_v12)
		_574_v13 = _nw106
		var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_572_v11), 0))
		_ = _index107
		(_572_v11).ArraySet1(_574_v13, (_index107).Int())
		var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_572_v11), 0))
		_ = _index108
		(_572_v11).ArraySet1(_574_v13, (_index108).Int())
		var _575_i5 _dafny.Int
		_ = _575_i5
		_575_i5 = _dafny.Zero
		{
			for p0 {
				{
					if (_575_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_575_i5 = (_575_i5).Plus(_dafny.One)
					var _576_v14 _dafny.Sequence
					_ = _576_v14
					_576_v14 = _dafny.SeqOf(p3)
					var _577_v15 _dafny.Sequence
					_ = _577_v15
					_577_v15 = _dafny.SeqOf(_576_v14, _576_v14)
					var _578_v16 _dafny.Map
					_ = _578_v16
					_578_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_577_v15, _577_v15), !(p3))
					r0 = (func() bool {
						if (_578_v16).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(699))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg25 _dafny.Int) interface{} {
								return coer25(arg25)
							}
						}((func(_579_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_580_i6 _dafny.Int) _dafny.Sequence {
								return _579_v14
							}
						})(_576_v14)))) {
							return (_578_v16).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(699))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg26 _dafny.Int) interface{} {
									return coer26(arg26)
								}
							}((func(_581_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_582_i6 _dafny.Int) _dafny.Sequence {
									return _581_v14
								}
							})(_576_v14)))).(bool)
						}
						return Companion_Default___.Fm33(_573_v12, globalState)
					})()
					var _583_v17 _dafny.Array
					_ = _583_v17
					var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
					_ = _nw107
					_583_v17 = _nw107
					var _584_v18 bool
					_ = _584_v18
					var _585_v19 _dafny.Map
					_ = _585_v19
					var _out10 bool
					_ = _out10
					var _out11 _dafny.Map
					_ = _out11
					_out10, _out11 = (_574_v13).M0(_583_v17, globalState)
					_584_v18 = _out10
					_585_v19 = _out11
					var _586_v20 _dafny.CodePoint
					_ = _586_v20
					_586_v20 = _dafny.CodePoint('v')
					_586_v20 = _586_v20
					var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((p1), 0))
					_ = _index109
					(p1).ArraySet1(_dafny.IntOfInt64(25), (_index109).Int())
					var _587_v21 _dafny.Array
					_ = _587_v21
					var _nw108 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.One)
					_ = _nw108
					_587_v21 = _nw108
					var _588_v22 D4
					_ = _588_v22
					_588_v22 = Companion_D4_.Create_DC9_(_569_v10)
					var _589_v23 D4
					_ = _589_v23
					_589_v23 = Companion_D4_.Create_DC11_(_588_v22)
					var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_587_v21), 0))
					_ = _index110
					(_587_v21).ArraySet1(Companion_D4_.Create_DC11_(_588_v22), (_index110).Int())
					var _590_v24 _dafny.Sequence
					_ = _590_v24
					_590_v24 = _dafny.SeqOf(_this.F7)
					var _591_v25 _dafny.Sequence
					_ = _591_v25
					_591_v25 = _dafny.SeqOf(_dafny.IntOfUint32((_590_v24).Cardinality()))
					var _592_v26 _dafny.MultiSet
					_ = _592_v26
					_592_v26 = _dafny.MultiSetOf(_this.F7, _dafny.IntOfUint32((_590_v24).Cardinality()))
					var _593_v27 D1
					_ = _593_v27
					_593_v27 = Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_dafny.SeqOf(p0, p0, _584_v18)).Cardinality()), _586_v20, p0, p2)
					var _594_v28 _dafny.Array
					_ = _594_v28
					var _nwElement0_23 bool = p3
					_ = _nwElement0_23
					var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(12))
					_ = _nw109
					(_nw109).ArraySet1(_nwElement0_23, 0)
					(_nw109).ArraySet1(p3, 1)
					(_nw109).ArraySet1(true, 2)
					(_nw109).ArraySet1(_584_v18, 3)
					(_nw109).ArraySet1(p3, 4)
					(_nw109).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_576_v14, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_576_v14).Cardinality()))).Uint32(), false)).Cardinality()))).IsProperSubsetOf(_592_v26), 5)
					(_nw109).ArraySet1((func() bool {
						if (_593_v27).Dtor_cf8() {
							return _584_v18
						}
						return !(p0)
					})(), 6)
					(_nw109).ArraySet1(((_576_v14).Select((Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_576_v14).Cardinality()))).Uint32()).(bool)) && (p0), 7)
					(_nw109).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_573_v12, _573_v12), 8)
					(_nw109).ArraySet1(true, 9)
					(_nw109).ArraySet1(_584_v18, 10)
					(_nw109).ArraySet1(!(!(p3)) || (_584_v18), 11)
					_594_v28 = _nw109
					var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_594_v28), 0))
					_ = _index111
					(_594_v28).ArraySet1((func() bool {
						if !(p3) {
							return p0
						}
						return !(_584_v18)
					})(), (_index111).Int())
					var _595_v29 _dafny.MultiSet
					_ = _595_v29
					_595_v29 = _dafny.MultiSetOf(_584_v18)
					var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((p1), 0))
					_ = _index112
					var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_587_v21), 0))
					_ = _index113
					var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_594_v28), 0))
					_ = _index114
					var _rhs80 _dafny.Int = (func() _dafny.Int {
						if (_595_v29).Contains(p3) {
							return (_595_v29).Multiplicity(p3)
						}
						return p2
					})()
					_ = _rhs80
					var _rhs81 D4 = Companion_D4_.Create_DC11_(_589_v23)
					_ = _rhs81
					var _rhs82 bool = !_dafny.Companion_Sequence_.Equal(_576_v14, _576_v14)
					_ = _rhs82
					var _lhs59 _dafny.Array = p1
					_ = _lhs59
					var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((p1), 0))
					_ = _lhs60
					var _lhs61 _dafny.Array = _587_v21
					_ = _lhs61
					var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_587_v21), 0))
					_ = _lhs62
					var _lhs63 _dafny.Array = _594_v28
					_ = _lhs63
					var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_594_v28), 0))
					_ = _lhs64
					(_lhs59).ArraySet1(_rhs80, (_lhs60).Int())
					(_lhs61).ArraySet1(_rhs81, (_lhs62).Int())
					(_lhs63).ArraySet1(_rhs82, (_lhs64).Int())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _596_v30 _dafny.MultiSet
		_ = _596_v30
		_596_v30 = _dafny.MultiSetOf(p3)
		var _597_v31 _dafny.Sequence
		_ = _597_v31
		_597_v31 = _dafny.SeqOf(false)
		var _598_v32 _dafny.Sequence
		_ = _598_v32
		_598_v32 = _dafny.SeqOf(_596_v30, _596_v30, _dafny.MultiSetFromSeq(_597_v31))
		var _599_v33 D2
		_ = _599_v33
		_599_v33 = Companion_D2_.Create_DC6_(_this.F7, p0)
		var _600_v34 _dafny.Array
		_ = _600_v34
		var _nwElement0_24 _dafny.Array = p1
		_ = _nwElement0_24
		var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(7))
		_ = _nw110
		(_nw110).ArraySet1(_nwElement0_24, 0)
		(_nw110).ArraySet1(_569_v10, 1)
		(_nw110).ArraySet1(p1, 2)
		(_nw110).ArraySet1(_569_v10, 3)
		(_nw110).ArraySet1(_569_v10, 4)
		(_nw110).ArraySet1(_569_v10, 5)
		(_nw110).ArraySet1(p1, 6)
		_600_v34 = _nw110
		var _601_v35 D0
		_ = _601_v35
		_601_v35 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus((_599_v33).Dtor_cf11()), _573_v12, _600_v34)
		var _602_v36 *C1
		_ = _602_v36
		var _nw111 *C1 = New_C1_()
		_ = _nw111
		_nw111.Ctor__((Companion_D10_.Create_DC23_((_598_v32).Select((Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_598_v32).Cardinality()))).Uint32()).(_dafny.MultiSet), p2, true, p0)).Dtor_cf35(), (_601_v35).Dtor_cf0())
		_602_v36 = _nw111
		var _603_v37 _dafny.Set
		_ = _603_v37
		_603_v37 = _dafny.SetOf(_602_v36, _602_v36, (func() *C1 {
			if p3 {
				return _602_v36
			}
			return _602_v36
		})())
		_603_v37 = _603_v37
		r0 = p0
		return r0
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f6 _dafny.Sequence
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f6 = _dafny.EmptySeq
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

func (_this *C5) Ctor__(f6 _dafny.Sequence) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C5) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		if !(false) || (false) {
			return _dafny.MultiSetOf(true, false, true, !(!(true)), !(false))
		} else {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(true, false, true, true))
		}
	}
}
func (_this *C5) Fm13(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		if (_dafny.MultiSetOf(_dafny.IntOfInt64(565), _dafny.IntOfUint32(((_this).F6()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(825))).Cardinality()))).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(156), (_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(-408), _dafny.IntOfUint32((_dafny.SeqOf(true, true, false, false)).Cardinality()))) {
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F6(), (_this).F6())).Cardinality())
		} else {
			return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-330), (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(602))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_604_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), (_this).F6())).Cardinality())
		}
	}
}
func (_this *C5) M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _605_v0 _dafny.Int
		_ = _605_v0
		_605_v0 = _dafny.IntOfInt64(112)
		var _606_v1 bool
		_ = _606_v1
		_606_v1 = false
		var _607_v2 _dafny.MultiSet
		_ = _607_v2
		_607_v2 = _dafny.MultiSetOf(_605_v0)
		var _608_v3 _dafny.Sequence
		_ = _608_v3
		_608_v3 = _dafny.SeqOf(_605_v0)
		var _609_v4 _dafny.Sequence
		_ = _609_v4
		_609_v4 = _dafny.SeqOf(_606_v1)
		var _610_v5 _dafny.Array
		_ = _610_v5
		var _nwElement0_25 _dafny.Int = _605_v0
		_ = _nwElement0_25
		var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(20))
		_ = _nw112
		(_nw112).ArraySet1(_nwElement0_25, 0)
		(_nw112).ArraySet1((_605_v0).Minus(_605_v0), 1)
		(_nw112).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_605_v0, _605_v0)), 2)
		(_nw112).ArraySet1(_605_v0, 3)
		(_nw112).ArraySet1(((func() _dafny.MultiSet {
			if _606_v1 {
				return _607_v2
			}
			return _dafny.MultiSetFromSeq(_608_v3)
		})()).Cardinality(), 4)
		(_nw112).ArraySet1(_605_v0, 5)
		(_nw112).ArraySet1(_605_v0, 6)
		(_nw112).ArraySet1(_605_v0, 7)
		(_nw112).ArraySet1(_dafny.IntOfUint32((_608_v3).Cardinality()), 8)
		(_nw112).ArraySet1(_dafny.IntOfInt64(686), 9)
		(_nw112).ArraySet1(_605_v0, 10)
		(_nw112).ArraySet1(_dafny.IntOfInt64(967), 11)
		(_nw112).ArraySet1((_605_v0).Minus(_605_v0), 12)
		(_nw112).ArraySet1(_605_v0, 13)
		(_nw112).ArraySet1(_605_v0, 14)
		(_nw112).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_606_v1, _606_v1), _609_v4)).Cardinality()), 15)
		(_nw112).ArraySet1(Companion_Default___.SafeModuloInt(_605_v0, _dafny.IntOfInt64(753)), 16)
		(_nw112).ArraySet1(_605_v0, 17)
		(_nw112).ArraySet1(_605_v0, 18)
		(_nw112).ArraySet1((_this).Fm13(_606_v1, _605_v0, _606_v1, globalState), 19)
		_610_v5 = _nw112
		var _611_v6 _dafny.MultiSet
		_ = _611_v6
		_611_v6 = _dafny.MultiSetOf((Companion_D3_.Create_DC7_(_607_v2)).Dtor_cf13(), _607_v2, _607_v2)
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
		_ = _index115
		(_610_v5).ArraySet1((func() _dafny.Int {
			if (_611_v6).Contains(_607_v2) {
				return (_611_v6).Multiplicity(_607_v2)
			}
			return _605_v0
		})(), (_index115).Int())
		var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
		_ = _index116
		(_610_v5).ArraySet1(_605_v0, (_index116).Int())
		var _612_v7 _dafny.Sequence
		_ = _612_v7
		_612_v7 = _dafny.SeqOf(_608_v3)
		var _613_v8 D2
		_ = _613_v8
		_613_v8 = Companion_D2_.Create_DC5_((_612_v7).Select((Companion_Default___.SafeIndex((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_612_v7).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _614_v9 _dafny.Map
		_ = _614_v9
		_614_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_613_v8).Dtor_cf10(), _606_v1)
		_614_v9 = (_614_v9).Update(_608_v3, _606_v1)
		if _606_v1 {
			var _615_v10 _dafny.Sequence
			_ = _615_v10
			_615_v10 = _dafny.SeqOf(_607_v2)
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
			_ = _index117
			var _rhs83 bool = !(!((((_615_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F6()).Cardinality()), _dafny.IntOfUint32((_615_v10).Cardinality()))).Uint32()).(_dafny.MultiSet)).Difference(_607_v2)).IsDisjointFrom(_607_v2)))
			_ = _rhs83
			var _rhs84 _dafny.Int = _605_v0
			_ = _rhs84
			var _lhs65 _dafny.Array = _610_v5
			_ = _lhs65
			var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
			_ = _lhs66
			r0 = _rhs83
			(_lhs65).ArraySet1(_rhs84, (_lhs66).Int())
			var _616_v11 D4
			_ = _616_v11
			_616_v11 = Companion_D4_.Create_DC9_(_610_v5)
			var _617_v12 _dafny.Set
			_ = _617_v12
			_617_v12 = _dafny.SetOf((_616_v11).Dtor_cf15(), _610_v5, _610_v5, _610_v5)
			if ((_608_v3).Select((Companion_Default___.SafeIndex((_617_v12).Cardinality(), _dafny.IntOfUint32((_608_v3).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_605_v0) <= 0 {
				var _618_v13 _dafny.CodePoint
				_ = _618_v13
				_618_v13 = _dafny.CodePoint('a')
				var _619_v14 _dafny.Map
				_ = _619_v14
				_619_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v0, _618_v13)
				var _620_v15 *C2
				_ = _620_v15
				var _nw113 *C2 = New_C2_()
				_ = _nw113
				_nw113.Ctor__(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}(func(_621_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), (_this).F6()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_609_v4).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}(func(_622_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), (_this).F6())).Cardinality()))).Uint32(), _618_v13), _dafny.Companion_Sequence_.Update((_this).F6(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F6()).Cardinality()), _dafny.IntOfUint32(((_this).F6()).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if (_619_v14).Contains((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)) {
						return (_619_v14).Get((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)).(_dafny.CodePoint)
					}
					return _618_v13
				})()))
				_620_v15 = _nw113
				r0 = _606_v1
				var _623_v16 _dafny.Array
				_ = _623_v16
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_11
				var _nw114 _dafny.Array
				_ = _nw114
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw114 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) bool = (func(_624_v1 bool) func(_dafny.Int) bool {
						return func(_625_i1 _dafny.Int) bool {
							return _624_v1
						}
					})(_606_v1)
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
				_623_v16 = _nw114
				_623_v16 = _623_v16
				var _626_v17 _dafny.Sequence
				_ = _626_v17
				_626_v17 = _dafny.UnicodeSeqOfUtf8Bytes("nrp")
				var _627_v18 _dafny.Array
				_ = _627_v18
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw115
				_627_v18 = _nw115
				var _628_v19 _dafny.Set
				_ = _628_v19
				_628_v19 = _dafny.SetOf(_605_v0)
				var _629_v20 _dafny.Map
				_ = _629_v20
				_629_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v0, _606_v1)
				var _630_v21 _dafny.Sequence
				_ = _630_v21
				_630_v21 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("oexlyc"), (_620_v15).F8())
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_627_v18), 0))
				_ = _index118
				(_627_v18).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_628_v19).Cardinality(), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)), _dafny.SeqOf(Companion_Default___.Fm19(_606_v1, (_629_v20).Cardinality(), globalState), _dafny.IntOfInt64(129), _dafny.IntOfUint32(((_630_v21).Select((Companion_Default___.SafeIndex((_607_v2).Cardinality(), _dafny.IntOfUint32((_630_v21).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))), (_index118).Int())
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_627_v18), 0))
				_ = _index119
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
				_ = _index120
				var _rhs85 _dafny.CodePoint = _618_v13
				_ = _rhs85
				var _rhs86 bool = (_620_v15).Fm18(_609_v4, globalState)
				_ = _rhs86
				var _rhs87 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(981))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_631_v5 _dafny.Array, _632_v13 _dafny.CodePoint, _633_v1 bool) func(_dafny.Int) _dafny.CodePoint {
					return func(_634_i2 _dafny.Int) _dafny.CodePoint {
						return (Companion_D1_.Create_DC4_((_631_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_631_v5), 0))).Int()).(_dafny.Int), _632_v13, _633_v1, (_631_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_631_v5), 0))).Int()).(_dafny.Int))).Dtor_cf7()
					}
				})(_610_v5, _618_v13, _606_v1)))
				_ = _rhs87
				var _rhs88 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_608_v3, _608_v3)
				_ = _rhs88
				var _rhs89 _dafny.Int = (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)
				_ = _rhs89
				var _lhs67 _dafny.Array = _627_v18
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_627_v18), 0))
				_ = _lhs68
				var _lhs69 _dafny.Array = _610_v5
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
				_ = _lhs70
				_618_v13 = _rhs85
				_606_v1 = _rhs86
				_626_v17 = _rhs87
				(_lhs67).ArraySet1(_rhs88, (_lhs68).Int())
				(_lhs69).ArraySet1(_rhs89, (_lhs70).Int())
				var _635_v22 *C0
				_ = _635_v22
				var _nw116 *C0 = New_C0_()
				_ = _nw116
				_nw116.Ctor__()
				_635_v22 = _nw116
			} else {
				var _636_v23 _dafny.Array
				_ = _636_v23
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
				_ = _nw117
				_636_v23 = _nw117
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_636_v23), 0))
				_ = _index121
				(_636_v23).ArraySet1(_610_v5, (_index121).Int())
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_636_v23), 0))
				_ = _index122
				(_636_v23).ArraySet1(_610_v5, (_index122).Int())
				var _637_v24 _dafny.Sequence
				_ = _637_v24
				_637_v24 = _dafny.SeqOf((_this).F6(), (_this).F6())
				var _638_v25 _dafny.CodePoint
				_ = _638_v25
				_638_v25 = _dafny.CodePoint('c')
				var _639_v26 _dafny.MultiSet
				_ = _639_v26
				_639_v26 = _dafny.MultiSetOf(_638_v25)
				var _rhs90 _dafny.Sequence = _608_v3
				_ = _rhs90
				var _rhs91 _dafny.Int = ((_dafny.IntOfUint32((_637_v24).Cardinality())).Minus(((_639_v26).Update(_638_v25, Companion_Default___.Abs((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)))).Cardinality())).Plus((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))
				_ = _rhs91
				_608_v3 = _rhs90
				_605_v0 = _rhs91
				var _640_v27 _dafny.Map
				_ = _640_v27
				_640_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _605_v0)
				_640_v27 = (_640_v27).Update(_606_v1, _605_v0)
				var _641_v28 D4
				_ = _641_v28
				_641_v28 = Companion_D4_.Create_DC9_(_dafny.ArrayCastTo((_636_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_636_v23), 0))).Int())))
				var _642_v29 D4
				_ = _642_v29
				_642_v29 = Companion_D4_.Create_DC11_(_641_v28)
				var _643_v30 D4
				_ = _643_v30
				_643_v30 = Companion_D4_.Create_DC11_(_641_v28)
				var _644_v31 _dafny.Sequence
				_ = _644_v31
				_644_v31 = _dafny.SeqOf(_643_v30, _643_v30, _643_v30, _643_v30, _643_v30)
				var _645_v32 _dafny.Map
				_ = _645_v32
				_645_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _644_v31)
				var _646_v33 _dafny.Map
				_ = _646_v33
				_646_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), (_645_v32).Cardinality())
				var _647_v36 _dafny.Array
				_ = _647_v36
				var _nwElement0_26 _dafny.Map = Companion_Default___.Fm29(globalState)
				_ = _nwElement0_26
				var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(13))
				_ = _nw118
				(_nw118).ArraySet1(_nwElement0_26, 0)
				(_nw118).ArraySet1((_646_v33).Update((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), _605_v0), 1)
				(_nw118).ArraySet1(_646_v33, 2)
				(_nw118).ArraySet1(func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(744), _dafny.IntOfInt64(9))); ; {
						_compr_18, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _648_v34 _dafny.Int
						_648_v34 = interface{}(_compr_18).(_dafny.Int)
						if ((_dafny.IntOfInt64(744)).Cmp(_648_v34) <= 0) && ((_648_v34).Cmp(_dafny.IntOfInt64(9)) < 0) {
							_coll18.Add(Companion_Default___.SafeModuloInt(_648_v34, _605_v0), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))
						}
					}
					return _coll18.ToMap()
				}(), 3)
				(_nw118).ArraySet1(_646_v33, 4)
				(_nw118).ArraySet1(func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(446), _dafny.IntOfInt64(923))); ; {
						_compr_19, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _649_v35 _dafny.Int
						_649_v35 = interface{}(_compr_19).(_dafny.Int)
						if ((_dafny.IntOfInt64(446)).Cmp(_649_v35) <= 0) && ((_649_v35).Cmp(_dafny.IntOfInt64(923)) < 0) {
							_coll19.Add(Companion_Default___.SafeModuloInt(_649_v35, (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)), _605_v0)
						}
					}
					return _coll19.ToMap()
				}(), 5)
				(_nw118).ArraySet1(_646_v33, 6)
				(_nw118).ArraySet1(_646_v33, 7)
				(_nw118).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))).Update(_605_v0, (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)), 8)
				(_nw118).ArraySet1(_646_v33, 9)
				(_nw118).ArraySet1(_646_v33, 10)
				(_nw118).ArraySet1(_646_v33, 11)
				(_nw118).ArraySet1((_646_v33).Update((func() _dafny.Int {
					if (_640_v27).Contains(_606_v1) {
						return (_640_v27).Get(_606_v1).(_dafny.Int)
					}
					return (_dafny.MultiSetOf(_606_v1, _606_v1, _606_v1)).Cardinality()
				})(), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)), 12)
				_647_v36 = _nw118
				var _650_v37 _dafny.Sequence
				_ = _650_v37
				_650_v37 = _dafny.SeqOf(_647_v36)
				_647_v36 = (_650_v37).Select((Companion_Default___.SafeIndex((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_650_v37).Cardinality()))).Uint32()).(_dafny.Array)
				_605_v0 = (Companion_Default___.Fm19(_606_v1, _dafny.IntOfUint32(((_this).F6()).Cardinality()), globalState)).Times(_dafny.IntOfInt64(152))
			}
			_616_v11 = _616_v11
			var _651_v38 D9
			_ = _651_v38
			_651_v38 = Companion_D9_.Create_DC20_(_606_v1, _606_v1, _606_v1)
			var _source8 D9 = _651_v38
			_ = _source8
			if _source8.Is_DC19() {
				var _652___mcc_h0 _dafny.Int = _source8.Get_().(D9_DC19).Cf29
				_ = _652___mcc_h0
				var _653___mcc_h1 _dafny.Int = _source8.Get_().(D9_DC19).Cf30
				_ = _653___mcc_h1
				var _654_cf30 _dafny.Int = _653___mcc_h1
				_ = _654_cf30
				var _655_cf29 _dafny.Int = _652___mcc_h0
				_ = _655_cf29
				var _656_v39 *C0
				_ = _656_v39
				var _nw119 *C0 = New_C0_()
				_ = _nw119
				_nw119.Ctor__()
				_656_v39 = _nw119
				var _657_v40 _dafny.Array
				_ = _657_v40
				var _nwElement0_27 *C0 = _656_v39
				_ = _nwElement0_27
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(10))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_27, 0)
				(_nw120).ArraySet1(_656_v39, 1)
				(_nw120).ArraySet1(_656_v39, 2)
				(_nw120).ArraySet1(_656_v39, 3)
				(_nw120).ArraySet1(_656_v39, 4)
				(_nw120).ArraySet1(_656_v39, 5)
				(_nw120).ArraySet1(_656_v39, 6)
				(_nw120).ArraySet1(_656_v39, 7)
				(_nw120).ArraySet1(_656_v39, 8)
				(_nw120).ArraySet1(_656_v39, 9)
				_657_v40 = _nw120
				var _658_v41 _dafny.Sequence
				_ = _658_v41
				_658_v41 = _dafny.SeqOf(_657_v40, _657_v40, _657_v40, _657_v40)
				var _659_v42 _dafny.Array
				_ = _659_v42
				var _nwElement0_28 _dafny.Array = _657_v40
				_ = _nwElement0_28
				var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(20))
				_ = _nw121
				(_nw121).ArraySet1(_nwElement0_28, 0)
				(_nw121).ArraySet1(_657_v40, 1)
				(_nw121).ArraySet1(_657_v40, 2)
				(_nw121).ArraySet1(_657_v40, 3)
				(_nw121).ArraySet1(_657_v40, 4)
				(_nw121).ArraySet1(_657_v40, 5)
				(_nw121).ArraySet1(_657_v40, 6)
				(_nw121).ArraySet1((func() _dafny.Array {
					if _606_v1 {
						return _657_v40
					}
					return (_658_v41).Select((Companion_Default___.SafeIndex((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_658_v41).Cardinality()))).Uint32()).(_dafny.Array)
				})(), 7)
				(_nw121).ArraySet1(_657_v40, 8)
				(_nw121).ArraySet1(_657_v40, 9)
				(_nw121).ArraySet1(_657_v40, 10)
				(_nw121).ArraySet1(_657_v40, 11)
				(_nw121).ArraySet1(_657_v40, 12)
				(_nw121).ArraySet1(_657_v40, 13)
				(_nw121).ArraySet1(_657_v40, 14)
				(_nw121).ArraySet1(_657_v40, 15)
				(_nw121).ArraySet1(_657_v40, 16)
				(_nw121).ArraySet1((_658_v41).Select((Companion_Default___.SafeIndex(_655_cf29, _dafny.IntOfUint32((_658_v41).Cardinality()))).Uint32()).(_dafny.Array), 17)
				(_nw121).ArraySet1(_657_v40, 18)
				(_nw121).ArraySet1(_657_v40, 19)
				_659_v42 = _nw121
				var _660_v43 _dafny.CodePoint
				_ = _660_v43
				_660_v43 = _dafny.CodePoint('r')
				var _661_v44 _dafny.Map
				_ = _661_v44
				_661_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v0, _659_v42)
				var _662_v45 _dafny.Map
				_ = _662_v45
				_662_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_606_v1, _659_v42)
				var _rhs92 _dafny.Array = (func() _dafny.Array {
					if _606_v1 {
						return (func() _dafny.Array {
							if (_661_v44).Contains(_655_cf29) {
								return (_661_v44).Get(_655_cf29).(_dafny.Array)
							}
							return _659_v42
						})()
					}
					return (func() _dafny.Array {
						if _606_v1 {
							return (func() _dafny.Array {
								if (_662_v45).Contains(_606_v1) {
									return (_662_v45).Get(_606_v1).(_dafny.Array)
								}
								return _659_v42
							})()
						}
						return _659_v42
					})()
				})()
				_ = _rhs92
				var _rhs93 _dafny.CodePoint = Companion_Default___.Fm30(_dafny.Companion_Sequence_.Concatenate((_this).F6(), (_this).F6()), _dafny.IntOfUint32((_608_v3).Cardinality()), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), globalState)
				_ = _rhs93
				_659_v42 = _rhs92
				_660_v43 = _rhs93
				var _663_v46 *C2
				_ = _663_v46
				var _nw122 *C2 = New_C2_()
				_ = _nw122
				_nw122.Ctor__(_dafny.Companion_Sequence_.Concatenate((_this).F6(), (_this).F6()), _dafny.UnicodeSeqOfUtf8Bytes("bjvex"))
				_663_v46 = _nw122
				_663_v46 = _663_v46
				var _664_v47 _dafny.MultiSet
				_ = _664_v47
				_664_v47 = _dafny.MultiSetOf(_606_v1, _606_v1)
				var _665_v48 _dafny.Set
				_ = _665_v48
				_665_v48 = _dafny.SetOf((func() _dafny.Int {
					if (_664_v47).Contains(!(_606_v1)) {
						return (_664_v47).Multiplicity(!(_606_v1))
					}
					return _655_cf29
				})())
				_606_v1 = (_665_v48).IsDisjointFrom(_665_v48)
				var _666_v49 _dafny.Map
				_ = _666_v49
				_666_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((_663_v46).F9(), (Companion_Default___.SafeIndex((_this).Fm13(_606_v1, _655_cf29, _606_v1, globalState), _dafny.IntOfUint32(((_663_v46).F9()).Cardinality()))).Uint32(), _dafny.CodePoint('u')), _dafny.IntOfInt64(618))
				_666_v49 = (_666_v49).Update((_this).F6(), _654_cf30)
			} else if _source8.Is_DC20() {
				var _667___mcc_h2 bool = _source8.Get_().(D9_DC20).Cf31
				_ = _667___mcc_h2
				var _668___mcc_h3 bool = _source8.Get_().(D9_DC20).Cf32
				_ = _668___mcc_h3
				var _669___mcc_h4 bool = _source8.Get_().(D9_DC20).Cf33
				_ = _669___mcc_h4
				var _670_cf33 bool = _669___mcc_h4
				_ = _670_cf33
				var _671_cf32 bool = _668___mcc_h3
				_ = _671_cf32
				var _672_cf31 bool = _667___mcc_h2
				_ = _672_cf31
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
				_ = _index123
				(_610_v5).ArraySet1(_605_v0, (_index123).Int())
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
				_ = _index124
				(_610_v5).ArraySet1(_605_v0, (_index124).Int())
				var _673_v50 _dafny.MultiSet
				_ = _673_v50
				_673_v50 = _dafny.MultiSetOf(_606_v1, _672_cf31)
				var _674_v51 T0
				_ = _674_v51
				var _nw123 *C4 = New_C4_()
				_ = _nw123
				_nw123.Ctor__(_605_v0)
				_674_v51 = _nw123
				var _675_v52 _dafny.Int
				_ = _675_v52
				var _676_v53 _dafny.CodePoint
				_ = _676_v53
				var _out12 _dafny.Int
				_ = _out12
				var _out13 _dafny.CodePoint
				_ = _out13
				_out12, _out13 = (_this).M3(!(((_dafny.MultiSetOf(_670_cf33, _672_cf31)).Update(_671_cf32, Companion_Default___.Abs(_dafny.IntOfInt64(16)))).IsSubsetOf(_673_v50)), _674_v51, globalState)
				_675_v52 = _out12
				_676_v53 = _out13
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
				_ = _index125
				(_610_v5).ArraySet1(_675_v52, (_index125).Int())
			} else {
				var _677___mcc_h5 _dafny.Map = _source8.Get_().(D9_DC18).Cf28
				_ = _677___mcc_h5
				var _678_cf28 _dafny.Map = _677___mcc_h5
				_ = _678_cf28
				var _679_v54 _dafny.Array
				_ = _679_v54
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
				_ = _nw124
				_679_v54 = _nw124
				var _680_v55 _dafny.CodePoint
				_ = _680_v55
				_680_v55 = _dafny.CodePoint('x')
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_679_v54), 0))
				_ = _index126
				(_679_v54).ArraySet1CodePoint(_680_v55, (_index126).Int())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_679_v54), 0))
				_ = _index127
				(_679_v54).ArraySet1CodePoint(_680_v55, (_index127).Int())
				_605_v0 = Companion_Default___.Fm14((_dafny.SetOf(_606_v1, _606_v1)).Cardinality(), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), globalState)
				var _681_v56 _dafny.Set
				_ = _681_v56
				_681_v56 = _dafny.SetOf(!(_dafny.Companion_Sequence_.Contains((_this).F6(), _680_v55)))
				_681_v56 = _681_v56
				_605_v0 = (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)
			}
			r0 = _606_v1
		} else {
			var _682_v58 _dafny.Array
			_ = _682_v58
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_12
			var _nw125 _dafny.Array
			_ = _nw125
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw125 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Set = (func(_683_v5 _dafny.Array, _684_v1 bool) func(_dafny.Int) _dafny.Set {
					return func(_685_i3 _dafny.Int) _dafny.Set {
						return func() _dafny.Set {
							var _coll20 = _dafny.NewBuilder()
							_ = _coll20
							for _iter24 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_683_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_683_v5), 0))).Int()).(_dafny.Int), _684_v1)).Keys().Elements()); ; {
								_compr_20, _ok24 := _iter24()
								if !_ok24 {
									break
								}
								var _686_v57 _dafny.Int
								_686_v57 = interface{}(_compr_20).(_dafny.Int)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_683_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_683_v5), 0))).Int()).(_dafny.Int), _684_v1)).Contains(_686_v57) {
									_coll20.Add((_686_v57).Plus((_dafny.MultiSetOf(_dafny.SeqOf(false))).Cardinality()))
								}
							}
							return _coll20.ToSet()
						}()
					}
				})(_610_v5, _606_v1)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw125 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw125).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw125).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_682_v58 = _nw125
			var _687_v59 _dafny.Set
			_ = _687_v59
			_687_v59 = _dafny.SetOf((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), _605_v0)
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_682_v58), 0))
			_ = _index128
			(_682_v58).ArraySet1(_687_v59, (_index128).Int())
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_682_v58), 0))
			_ = _index129
			(_682_v58).ArraySet1(_687_v59, (_index129).Int())
			_606_v1 = !(_606_v1)
			var _688_v60 D4
			_ = _688_v60
			_688_v60 = Companion_D4_.Create_DC10_(_606_v1)
			var _689_v61 _dafny.Array
			_ = _689_v61
			var _nwElement0_29 bool = _606_v1
			_ = _nwElement0_29
			var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(13))
			_ = _nw126
			(_nw126).ArraySet1(_nwElement0_29, 0)
			(_nw126).ArraySet1(_606_v1, 1)
			(_nw126).ArraySet1(_606_v1, 2)
			(_nw126).ArraySet1(false, 3)
			(_nw126).ArraySet1(!(_606_v1), 4)
			(_nw126).ArraySet1(_606_v1, 5)
			(_nw126).ArraySet1((_688_v60).Dtor_cf16(), 6)
			(_nw126).ArraySet1(!(_606_v1), 7)
			(_nw126).ArraySet1(_606_v1, 8)
			(_nw126).ArraySet1(_606_v1, 9)
			(_nw126).ArraySet1(_606_v1, 10)
			(_nw126).ArraySet1(_606_v1, 11)
			(_nw126).ArraySet1(_606_v1, 12)
			_689_v61 = _nw126
			var _690_v62 _dafny.Array
			_ = _690_v62
			_690_v62 = _689_v61
			_689_v61 = (_690_v62)
			_605_v0 = Companion_Default___.SafeDivisionInt(_605_v0, (_dafny.Zero).Minus((_dafny.Zero).Minus((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))))
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
			_ = _index130
			(_610_v5).ArraySet1((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), (_index130).Int())
		}
		var _hi5 _dafny.Int = _605_v0
		_ = _hi5
		for _691_i4 := Companion_Default___.SafeDivisionInt((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rgtil")).Cardinality()))); _691_i4.Cmp(_hi5) < 0; _691_i4 = _691_i4.Plus(_dafny.One) {
			_608_v3 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_608_v3, _608_v3), _608_v3)
			r0 = _606_v1
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
			_ = _index131
			(_610_v5).ArraySet1((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), (_index131).Int())
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
			_ = _index132
			(_610_v5).ArraySet1(Companion_Default___.SafeDivisionInt(_691_i4, (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)), (_index132).Int())
		}
		var _692_v63 _dafny.Set
		_ = _692_v63
		_692_v63 = _dafny.SetOf((_this).F6(), (_this).F6(), (_this).F6())
		var _693_i5 _dafny.Int
		_ = _693_i5
		_693_i5 = _dafny.Zero
		{
			for (_dafny.SetOf((_this).F6())).IsDisjointFrom(_692_v63) {
				{
					if (_693_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_693_i5 = (_693_i5).Plus(_dafny.One)
					var _694_v64 *C3
					_ = _694_v64
					var _nw127 *C3 = New_C3_()
					_ = _nw127
					_nw127.Ctor__()
					_694_v64 = _nw127
					var _695_v65 D11
					_ = _695_v65
					_695_v65 = Companion_D11_.Create_DC24_(_694_v64)
					var _696_v66 _dafny.Sequence
					_ = _696_v66
					_696_v66 = _dafny.SeqOf(_694_v64, _694_v64, _694_v64)
					var _697_v67 _dafny.Array
					_ = _697_v67
					var _nwElement0_30 *C3 = _694_v64
					_ = _nwElement0_30
					var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(11))
					_ = _nw128
					(_nw128).ArraySet1(_nwElement0_30, 0)
					(_nw128).ArraySet1(_694_v64, 1)
					(_nw128).ArraySet1(_694_v64, 2)
					(_nw128).ArraySet1(_694_v64, 3)
					(_nw128).ArraySet1(_694_v64, 4)
					(_nw128).ArraySet1(_694_v64, 5)
					(_nw128).ArraySet1(_694_v64, 6)
					(_nw128).ArraySet1((_695_v65).Dtor_cf39(), 7)
					(_nw128).ArraySet1(_694_v64, 8)
					(_nw128).ArraySet1(_694_v64, 9)
					(_nw128).ArraySet1((_696_v66).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.IntOfUint32((_696_v66).Cardinality()))).Uint32()).(*C3), 10)
					_697_v67 = _nw128
					var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_697_v67), 0))
					_ = _index133
					(_697_v67).ArraySet1(_694_v64, (_index133).Int())
					var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_697_v67), 0))
					_ = _index134
					var _nw129 *C3 = New_C3_()
					_ = _nw129
					_nw129.Ctor__()
					(_697_v67).ArraySet1(_nw129, (_index134).Int())
					var _698_v68 _dafny.Map
					_ = _698_v68
					_698_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)), _dafny.UnicodeSeqOfUtf8Bytes("mvbygep"))
					var _699_v69 _dafny.Array
					_ = _699_v69
					var _nwElement0_31 bool = _606_v1
					_ = _nwElement0_31
					var _nw130 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(3))
					_ = _nw130
					(_nw130).ArraySet1(_nwElement0_31, 0)
					(_nw130).ArraySet1(_606_v1, 1)
					(_nw130).ArraySet1(_606_v1, 2)
					_699_v69 = _nw130
					var _700_v70 _dafny.Array
					_ = _700_v70
					_700_v70 = _699_v69
					var _701_v71 _dafny.Set
					_ = _701_v71
					_701_v71 = _dafny.SetOf((_700_v70), _699_v69)
					_698_v68 = (_698_v68).Update(Companion_Default___.SafeModuloInt(_605_v0, (_701_v71).Cardinality()), _dafny.Companion_Sequence_.Concatenate((_this).F6(), (_this).F6()))
					var _702_v72 _dafny.Array
					_ = _702_v72
					var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
					_ = _nw131
					_702_v72 = _nw131
					_702_v72 = p0
					var _703_v73 _dafny.Set
					_ = _703_v73
					_703_v73 = _dafny.SetOf(false, true)
					var _704_v74 _dafny.Map
					_ = _704_v74
					_704_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_605_v0).Minus(_605_v0), _703_v73)
					var _705_v75 _dafny.Map
					_ = _705_v75
					_705_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_607_v2, ((_dafny.Zero).Minus((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))).Minus((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)))
					var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
					_ = _index135
					var _rhs94 _dafny.Int = (func() _dafny.Int {
						if (_705_v75).Contains((_dafny.MultiSetOf((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), _605_v0)).Intersection(Companion_Default___.Fm31((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), globalState))) {
							return (_705_v75).Get((_dafny.MultiSetOf((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), _605_v0)).Intersection(Companion_Default___.Fm31((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), globalState))).(_dafny.Int)
						}
						return (_dafny.Zero).Minus((_this).Fm13(!(_606_v1), Companion_Default___.Fm24(_605_v0, false, _606_v1, _dafny.IntOfUint32(((_this).F6()).Cardinality()), globalState), _606_v1, globalState))
					})()
					_ = _rhs94
					var _rhs95 _dafny.Map = _704_v74
					_ = _rhs95
					var _lhs71 _dafny.Array = _610_v5
					_ = _lhs71
					var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
					_ = _lhs72
					(_lhs71).ArraySet1(_rhs94, (_lhs72).Int())
					_704_v74 = _rhs95
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		if _606_v1 {
			var _706_v76 _dafny.MultiSet
			_ = _706_v76
			_706_v76 = _dafny.MultiSetOf(_606_v1)
			var _707_v77 D10
			_ = _707_v77
			_707_v77 = Companion_D10_.Create_DC21_(_706_v76)
			var _708_v78 _dafny.Map
			_ = _708_v78
			_708_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_609_v4).Cardinality()), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))
			var _709_v80 _dafny.Map
			_ = _709_v80
			_709_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_707_v77, (func() _dafny.Set {
				var _coll21 = _dafny.NewBuilder()
				_ = _coll21
				for _iter25 := _dafny.Iterate((_708_v78).Keys().Elements()); ; {
					_compr_21, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _710_v79 _dafny.Int
					_710_v79 = interface{}(_compr_21).(_dafny.Int)
					if (_708_v78).Contains(_710_v79) {
						_coll21.Add(Companion_Default___.SafeModuloInt(_710_v79, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D9_.Create_DC19_(_dafny.IntOfInt64(-120), _dafny.IntOfInt64(390))).Dtor_cf29(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(768))).Cardinality())).Cardinality()))
					}
				}
				return _coll21.ToSet()
			}()).Cardinality())
			_709_v80 = (_709_v80).Update(_707_v77, _dafny.IntOfInt64(283))
			var _711_v81 *C1
			_ = _711_v81
			var _nw132 *C1 = New_C1_()
			_ = _nw132
			_nw132.Ctor__((_706_v76).Intersection(_706_v76), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))
			_711_v81 = _nw132
			if _606_v1 {
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
				_ = _index136
				(_610_v5).ArraySet1(_605_v0, (_index136).Int())
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
				_ = _index137
				(_610_v5).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-96)), (_index137).Int())
				var _712_v82 _dafny.CodePoint
				_ = _712_v82
				_712_v82 = _dafny.CodePoint('v')
				_712_v82 = ((_this).F6()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32(((_this).F6()).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_606_v1 = !(Companion_Default___.Fm33((_this).F6(), globalState))
				var _713_v83 _dafny.Array
				_ = _713_v83
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw133
				_713_v83 = _nw133
				_713_v83 = _713_v83
			} else {
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
				_ = _index138
				(_610_v5).ArraySet1(_711_v81.F11, (_index138).Int())
				_609_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_609_v4, _dafny.SeqOf(_606_v1)), _609_v4)
				var _714_v84 *C0
				_ = _714_v84
				var _nw134 *C0 = New_C0_()
				_ = _nw134
				_nw134.Ctor__()
				_714_v84 = _nw134
				var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw135
				_610_v5 = _nw135
				_605_v0 = (_dafny.Zero).Minus(_711_v81.F11)
			}
			var _715_v85 *C0
			_ = _715_v85
			var _nw136 *C0 = New_C0_()
			_ = _nw136
			_nw136.Ctor__()
			_715_v85 = _nw136
			var _716_v86 D7
			_ = _716_v86
			_716_v86 = Companion_D7_.Create_DC15_(_606_v1, _711_v81.F11, _605_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), _711_v81.F11)).Cardinality(), _715_v85)
			var _717_v87 _dafny.Map
			_ = _717_v87
			_717_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_716_v86).Dtor_cf24(), _606_v1)
			var _718_v88 _dafny.Map
			_ = _718_v88
			_718_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_717_v87, _605_v0)
			var _719_v90 _dafny.CodePoint
			_ = _719_v90
			_719_v90 = _dafny.CodePoint('m')
			var _720_v91 _dafny.Map
			_ = _720_v91
			_720_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_711_v81.F11, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_719_v90, _606_v1))
			_718_v88 = (_718_v88).Update(func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter26 := _dafny.Iterate((_720_v91).Keys().Elements()); ; {
					_compr_22, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _721_v89 _dafny.Int
					_721_v89 = interface{}(_compr_22).(_dafny.Int)
					if (_720_v91).Contains(_721_v89) {
						_coll22.Add((_721_v89).Minus(_711_v81.F11), !(false))
					}
				}
				return _coll22.ToMap()
			}(), (_this).Fm13(_606_v1, (func() _dafny.Int {
				if (_708_v78).Contains(_dafny.IntOfInt64(-45)) {
					return (_708_v78).Get(_dafny.IntOfInt64(-45)).(_dafny.Int)
				}
				return _605_v0
			})(), _606_v1, globalState))
			var _722_v92 _dafny.Array
			_ = _722_v92
			var _nw137 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
			_ = _nw137
			_722_v92 = _nw137
			_722_v92 = _722_v92
		} else {
			var _723_v94 _dafny.Array
			_ = _723_v94
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_13
			var _nw138 _dafny.Array
			_ = _nw138
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw138 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Set = (func(_724_v5 _dafny.Array, _725_v2 _dafny.MultiSet) func(_dafny.Int) _dafny.Set {
					return func(_726_i6 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf((_724_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_724_v5), 0))).Int()).(_dafny.Int), (_724_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_724_v5), 0))).Int()).(_dafny.Int))).Union(func() _dafny.Set {
							var _coll23 = _dafny.NewBuilder()
							_ = _coll23
							for _iter27 := _dafny.Iterate((_725_v2).Elements()); ; {
								_compr_23, _ok27 := _iter27()
								if !_ok27 {
									break
								}
								var _727_v93 _dafny.Int
								_727_v93 = interface{}(_compr_23).(_dafny.Int)
								if (_725_v2).Contains(_727_v93) {
									_coll23.Add((_727_v93).Times(_dafny.IntOfInt64(647)))
								}
							}
							return _coll23.ToSet()
						}())
					}
				})(_610_v5, _607_v2)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw138 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw138).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw138).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_723_v94 = _nw138
			var _728_v95 _dafny.Set
			_ = _728_v95
			_728_v95 = _dafny.SetOf((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_723_v94), 0))
			_ = _index139
			(_723_v94).ArraySet1((_728_v95).Difference(func() _dafny.Set {
				var _coll24 = _dafny.NewBuilder()
				_ = _coll24
				for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(734), _dafny.IntOfInt64(782))); ; {
					_compr_24, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _729_v96 _dafny.Int
					_729_v96 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(734)).Cmp(_729_v96) <= 0) && ((_729_v96).Cmp(_dafny.IntOfInt64(782)) < 0) {
						_coll24.Add((_729_v96).Minus((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)))
					}
				}
				return _coll24.ToSet()
			}()), (_index139).Int())
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_723_v94), 0))
			_ = _index140
			(_723_v94).ArraySet1((_728_v95).Difference(func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter29 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(9))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}((func(_730_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_731_i7 _dafny.Int) _dafny.Int {
						return _730_v0
					}
				})(_605_v0)))).Elements()); ; {
					_compr_25, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _732_v97 _dafny.Int
					_732_v97 = interface{}(_compr_25).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(9))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg32 _dafny.Int) interface{} {
							return coer32(arg32)
						}
					}((func(_733_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_731_i7 _dafny.Int) _dafny.Int {
							return _733_v0
						}
					})(_605_v0))), _732_v97) {
						_coll25.Add((_732_v97).Times(_dafny.IntOfInt64(907)))
					}
				}
				return _coll25.ToSet()
			}()), (_index140).Int())
			var _734_v98 _dafny.Set
			_ = _734_v98
			_734_v98 = _dafny.SetOf(_606_v1, _606_v1)
			var _735_v99 _dafny.MultiSet
			_ = _735_v99
			_735_v99 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_734_v98).Cardinality()), (_dafny.Zero).Minus((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm24(_605_v0, _606_v1, _606_v1, _605_v0, globalState), _605_v0))
			var _736_v100 _dafny.Array
			_ = _736_v100
			var _nwElement0_32 bool = true
			_ = _nwElement0_32
			var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
			_ = _nw139
			(_nw139).ArraySet1(_nwElement0_32, 0)
			(_nw139).ArraySet1(_606_v1, 1)
			(_nw139).ArraySet1(true, 2)
			(_nw139).ArraySet1(_606_v1, 3)
			(_nw139).ArraySet1(_606_v1, 4)
			(_nw139).ArraySet1((_606_v1) == (!(_606_v1)), 5)
			(_nw139).ArraySet1((_735_v99).IsSubsetOf(_735_v99), 6)
			(_nw139).ArraySet1(_606_v1, 7)
			(_nw139).ArraySet1(_606_v1, 8)
			(_nw139).ArraySet1(_606_v1, 9)
			(_nw139).ArraySet1(_606_v1, 10)
			(_nw139).ArraySet1(((_dafny.MultiSetFromSeq(_609_v4)).Update(_606_v1, Companion_Default___.Abs((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)))).Contains(_606_v1), 11)
			(_nw139).ArraySet1(!((_605_v0).Cmp(_dafny.IntOfInt64(317)) <= 0), 12)
			(_nw139).ArraySet1(_606_v1, 13)
			(_nw139).ArraySet1(!(_606_v1) || (_606_v1), 14)
			(_nw139).ArraySet1(_606_v1, 15)
			(_nw139).ArraySet1(!((func() bool {
				if _606_v1 {
					return _606_v1
				}
				return _606_v1
			})()), 16)
			(_nw139).ArraySet1(_606_v1, 17)
			(_nw139).ArraySet1(_606_v1, 18)
			(_nw139).ArraySet1(_606_v1, 19)
			(_nw139).ArraySet1(_606_v1, 20)
			(_nw139).ArraySet1(_606_v1, 21)
			(_nw139).ArraySet1(Companion_Default___.Fm33(_dafny.UnicodeSeqOfUtf8Bytes("lemmjobsq"), globalState), 22)
			(_nw139).ArraySet1(!(_606_v1), 23)
			(_nw139).ArraySet1(_606_v1, 24)
			_736_v100 = _nw139
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_736_v100), 0))
			_ = _index141
			(_736_v100).ArraySet1((_606_v1) == (_606_v1), (_index141).Int())
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_736_v100), 0))
			_ = _index142
			var _rhs96 _dafny.Sequence = _609_v4
			_ = _rhs96
			var _rhs97 _dafny.Array = _610_v5
			_ = _rhs97
			var _rhs98 bool = _606_v1
			_ = _rhs98
			var _rhs99 bool = ((_this).Fm13(_606_v1, (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), true, globalState)).Cmp(_605_v0) > 0
			_ = _rhs99
			var _lhs73 _dafny.Array = _736_v100
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_736_v100), 0))
			_ = _lhs74
			_609_v4 = _rhs96
			_610_v5 = _rhs97
			(_lhs73).ArraySet1(_rhs98, (_lhs74).Int())
			r0 = _rhs99
			var _rhs100 _dafny.Array = _610_v5
			_ = _rhs100
			var _rhs101 _dafny.Sequence = _608_v3
			_ = _rhs101
			_610_v5 = _rhs100
			_608_v3 = _rhs101
			var _737_v101 _dafny.Array
			_ = _737_v101
			var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
			_ = _nw140
			_737_v101 = _nw140
			var _738_v102 _dafny.Map
			_ = _738_v102
			_738_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_737_v101, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)), (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int)))
			_738_v102 = (_738_v102).Update(_737_v101, (_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int))
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))
			_ = _index143
			(_610_v5).ArraySet1((_610_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_610_v5), 0))).Int()).(_dafny.Int), (_index143).Int())
		}
		r0 = _606_v1
		var _739_v103 _dafny.MultiSet
		_ = _739_v103
		_739_v103 = _dafny.MultiSetOf(_606_v1)
		var _740_v104 _dafny.Map
		_ = _740_v104
		_740_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(706))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}(func(_741_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))).Cardinality()), _739_v103)
		var _742_v105 _dafny.Sequence
		_ = _742_v105
		_742_v105 = _dafny.SeqOf((func() _dafny.Map {
			if _606_v1 {
				return _740_v104
			}
			return _740_v104
		})(), _740_v104, _740_v104, _740_v104)
		r1 = (_742_v105).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_this).F6()).Cardinality()), _605_v0)), _dafny.IntOfUint32((_742_v105).Cardinality()))).Uint32()).(_dafny.Map)
		return r0, r1
	}
}
func (_this *C5) M3(p0 bool, p1 T0, globalState *GlobalState) (_dafny.Int, _dafny.CodePoint) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var _743_v0 _dafny.Int
		_ = _743_v0
		_743_v0 = _dafny.IntOfInt64(-822)
		var _744_v1 _dafny.Sequence
		_ = _744_v1
		_744_v1 = _dafny.SeqOf(_743_v0, _743_v0)
		var _745_v2 D1
		_ = _745_v2
		_745_v2 = Companion_D1_.Create_DC4_((_744_v1).Select((Companion_Default___.SafeIndex(_743_v0, _dafny.IntOfUint32((_744_v1).Cardinality()))).Uint32()).(_dafny.Int), _dafny.CodePoint('b'), false, _dafny.IntOfUint32(((_this).F6()).Cardinality()))
		var _746_i0 _dafny.Int
		_ = _746_i0
		_746_i0 = _dafny.Zero
		{
			var _pat_let_tv5 = p0
			_ = _pat_let_tv5
			var _pat_let_tv6 = p0
			_ = _pat_let_tv6
			for func(_source9 D1) bool {
				if _source9.Is_DC3() {
					var _757___mcc_h0 _dafny.Int = _source9.Get_().(D1_DC3).Cf5
					_ = _757___mcc_h0
					var _758_cf5 _dafny.Int = _757___mcc_h0
					_ = _758_cf5
					return !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.SeqOf(false)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(755))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg34 _dafny.Int) interface{} {
							return coer34(arg34)
						}
					}((func(_759_p0 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_760_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_759_p0, false)
						}
					})(_pat_let_tv5)))))
				} else if _source9.Is_DC4() {
					var _761___mcc_h1 _dafny.Int = _source9.Get_().(D1_DC4).Cf6
					_ = _761___mcc_h1
					var _762___mcc_h2 _dafny.CodePoint = _source9.Get_().(D1_DC4).Cf7
					_ = _762___mcc_h2
					var _763___mcc_h3 bool = _source9.Get_().(D1_DC4).Cf8
					_ = _763___mcc_h3
					var _764___mcc_h4 _dafny.Int = _source9.Get_().(D1_DC4).Cf9
					_ = _764___mcc_h4
					var _765_cf9 _dafny.Int = _764___mcc_h4
					_ = _765_cf9
					var _766_cf8 bool = _763___mcc_h3
					_ = _766_cf8
					var _767_cf7 _dafny.CodePoint = _762___mcc_h2
					_ = _767_cf7
					var _768_cf6 _dafny.Int = _761___mcc_h1
					_ = _768_cf6
					return _pat_let_tv6
				} else {
					var _769___mcc_h5 bool = _source9.Get_().(D1_DC2).Cf4
					_ = _769___mcc_h5
					var _770_cf4 bool = _769___mcc_h5
					_ = _770_cf4
					return !(!(_770_cf4))
				}
			}(_745_v2) {
				{
					if (_746_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_746_i0 = (_746_i0).Plus(_dafny.One)
					var _747_v3 bool
					_ = _747_v3
					_747_v3 = false
					_747_v3 = p0
					var _748_v4 _dafny.Array
					_ = _748_v4
					var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
					_ = _nw141
					_748_v4 = _nw141
					var _749_v5 D2
					_ = _749_v5
					_749_v5 = Companion_D2_.Create_DC6_(_743_v0, p0)
					var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))
					_ = _index144
					(_748_v4).ArraySet1((_749_v5).Dtor_cf11(), (_index144).Int())
					var _750_v7 _dafny.Map
					_ = _750_v7
					_750_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_747_v3, _747_v3)
					var _751_v8 _dafny.Set
					_ = _751_v8
					_751_v8 = _dafny.SetOf((_750_v7).Cardinality(), (_dafny.Zero).Minus(_743_v0))
					var _752_v9 _dafny.Sequence
					_ = _752_v9
					_752_v9 = _dafny.SeqOf(_747_v3, (_751_v8).IsProperSubsetOf(func() _dafny.Set {
						var _coll26 = _dafny.NewBuilder()
						_ = _coll26
						for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(553), _dafny.IntOfInt64(891))); ; {
							_compr_26, _ok30 := _iter30()
							if !_ok30 {
								break
							}
							var _753_v6 _dafny.Int
							_753_v6 = interface{}(_compr_26).(_dafny.Int)
							if ((_dafny.IntOfInt64(553)).Cmp(_753_v6) <= 0) && ((_753_v6).Cmp(_dafny.IntOfInt64(891)) < 0) {
								_coll26.Add(Companion_Default___.SafeModuloInt(_753_v6, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ba")).Cardinality())))
							}
						}
						return _coll26.ToSet()
					}()), p0, !(_747_v3))
					var _754_v10 _dafny.Sequence
					_ = _754_v10
					_754_v10 = _dafny.SeqOf(_752_v9)
					var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))
					_ = _index145
					var _rhs102 _dafny.Int = (_743_v0).Plus(_743_v0)
					_ = _rhs102
					var _rhs103 _dafny.Int = Companion_Default___.SafeDivisionInt(_743_v0, (func() _dafny.Int {
						if _747_v3 {
							return _dafny.IntOfInt64(355)
						}
						return _743_v0
					})())
					_ = _rhs103
					var _rhs104 _dafny.Int = _743_v0
					_ = _rhs104
					var _rhs105 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm26(_743_v0, globalState), _752_v9), (_754_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aaqjua")).Cardinality()), _dafny.IntOfUint32((_754_v10).Cardinality()))).Uint32()).(_dafny.Sequence))
					_ = _rhs105
					var _lhs75 _dafny.Array = _748_v4
					_ = _lhs75
					var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))
					_ = _lhs76
					r0 = _rhs102
					r0 = _rhs103
					(_lhs75).ArraySet1(_rhs104, (_lhs76).Int())
					_752_v9 = _rhs105
					var _755_v11 *C3
					_ = _755_v11
					var _nw142 *C3 = New_C3_()
					_ = _nw142
					_nw142.Ctor__()
					_755_v11 = _nw142
					var _756_v12 _dafny.Map
					_ = _756_v12
					_756_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_748_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))).Int()).(_dafny.Int), (_748_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))).Int()).(_dafny.Int)), (_this).F6())
					var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))
					_ = _index146
					var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))
					_ = _index147
					var _rhs106 _dafny.Int = (_756_v12).Cardinality()
					_ = _rhs106
					var _rhs107 _dafny.Int = _dafny.IntOfInt64(-470)
					_ = _rhs107
					var _rhs108 bool = _747_v3
					_ = _rhs108
					var _lhs77 _dafny.Array = _748_v4
					_ = _lhs77
					var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))
					_ = _lhs78
					var _lhs79 _dafny.Array = _748_v4
					_ = _lhs79
					var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_748_v4), 0))
					_ = _lhs80
					(_lhs77).ArraySet1(_rhs106, (_lhs78).Int())
					(_lhs79).ArraySet1(_rhs107, (_lhs80).Int())
					_747_v3 = _rhs108
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _771_v13 _dafny.Array
		_ = _771_v13
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_14
		var _nw143 _dafny.Array
		_ = _nw143
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw143 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Int = (func(_772_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_773_i2 _dafny.Int) _dafny.Int {
					return (_773_i2).Plus(_772_v0)
				}
			})(_743_v0)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw143 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw143).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw143).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_771_v13 = _nw143
		_771_v13 = _771_v13
		var _774_v14 _dafny.Array
		_ = _774_v14
		var _nw144 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw144
		_774_v14 = _nw144
		var _775_v15 _dafny.Map
		_ = _775_v15
		_775_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F6()).Cardinality()), _774_v14)
		if (_775_v15).Equals(_775_v15) {
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_774_v14), 0))
			_ = _index148
			(_774_v14).ArraySet1(p0, (_index148).Int())
			var _776_v16 _dafny.Array
			_ = _776_v16
			var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw145
			_776_v16 = _nw145
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_776_v16), 0))
			_ = _index149
			(_776_v16).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F6(), (_this).F6()), (_index149).Int())
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_774_v14), 0))
			_ = _index150
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_776_v16), 0))
			_ = _index151
			var _rhs109 bool = p0
			_ = _rhs109
			var _rhs110 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F6(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(948))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_777_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			})))).Cardinality())
			_ = _rhs110
			var _rhs111 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_this).F6(), (_this).F6()), _dafny.UnicodeSeqOfUtf8Bytes("npyc"))
			_ = _rhs111
			var _lhs81 _dafny.Array = _774_v14
			_ = _lhs81
			var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_774_v14), 0))
			_ = _lhs82
			var _lhs83 _dafny.Array = _776_v16
			_ = _lhs83
			var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_776_v16), 0))
			_ = _lhs84
			(_lhs81).ArraySet1(_rhs109, (_lhs82).Int())
			r0 = _rhs110
			(_lhs83).ArraySet1(_rhs111, (_lhs84).Int())
			_743_v0 = (_dafny.IntOfInt64(-516)).Plus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_v0, p0)).Cardinality()).Minus((_744_v1).Select((Companion_Default___.SafeIndex(_743_v0, _dafny.IntOfUint32((_744_v1).Cardinality()))).Uint32()).(_dafny.Int)))
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_776_v16), 0))
			_ = _index152
			(_776_v16).ArraySet1((_this).F6(), (_index152).Int())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_774_v14), 0))
			_ = _index153
			(_774_v14).ArraySet1((_774_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_774_v14), 0))).Int()).(bool), (_index153).Int())
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_774_v14), 0))
			_ = _index154
			(_774_v14).ArraySet1(p0, (_index154).Int())
		} else {
			r0 = _dafny.IntOfInt64(-976)
			var _778_v17 bool
			_ = _778_v17
			_778_v17 = false
			var _779_v18 _dafny.Sequence
			_ = _779_v18
			_779_v18 = _dafny.SeqOf(p0, _778_v17)
			var _780_v19 D1
			_ = _780_v19
			_780_v19 = Companion_D1_.Create_DC2_((_779_v18).Select((Companion_Default___.SafeIndex(_743_v0, _dafny.IntOfUint32((_779_v18).Cardinality()))).Uint32()).(bool))
			var _781_v20 _dafny.CodePoint
			_ = _781_v20
			_781_v20 = _dafny.CodePoint('g')
			var _pat_let_tv7 = _743_v0
			_ = _pat_let_tv7
			var _pat_let_tv8 = _781_v20
			_ = _pat_let_tv8
			var _pat_let_tv9 = _778_v17
			_ = _pat_let_tv9
			var _rhs112 bool = p0
			_ = _rhs112
			var _rhs113 bool = !(!(false) || (false))
			_ = _rhs113
			var _rhs114 bool = (_780_v19).Dtor_cf4()
			_ = _rhs114
			var _rhs115 _dafny.CodePoint = (func(_pat_let8_0 D1) D1 {
				return func(_785_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let12_0 bool) D1 {
						return func(_786_dt__update_hcf8_h0 bool) D1 {
							return Companion_D1_.Create_DC4_((_785_dt__update__tmp_h1).Dtor_cf6(), (_785_dt__update__tmp_h1).Dtor_cf7(), _786_dt__update_hcf8_h0, (_785_dt__update__tmp_h1).Dtor_cf9())
						}(_pat_let12_0)
					}(_pat_let_tv9)
				}(_pat_let8_0)
			}(func(_pat_let9_0 D1) D1 {
				return func(_782_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let10_0 _dafny.Int) D1 {
						return func(_783_dt__update_hcf9_h0 _dafny.Int) D1 {
							return func(_pat_let11_0 _dafny.CodePoint) D1 {
								return func(_784_dt__update_hcf7_h0 _dafny.CodePoint) D1 {
									return Companion_D1_.Create_DC4_((_782_dt__update__tmp_h0).Dtor_cf6(), _784_dt__update_hcf7_h0, (_782_dt__update__tmp_h0).Dtor_cf8(), _783_dt__update_hcf9_h0)
								}(_pat_let11_0)
							}(_pat_let_tv8)
						}(_pat_let10_0)
					}(_pat_let_tv7)
				}(_pat_let9_0)
			}(_745_v2))).Dtor_cf7()
			_ = _rhs115
			var _rhs116 bool = Companion_Default___.Fm33(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xsfeu"), _dafny.UnicodeSeqOfUtf8Bytes("mosbxp")), globalState)
			_ = _rhs116
			_778_v17 = _rhs112
			_778_v17 = _rhs113
			_778_v17 = _rhs114
			r1 = _rhs115
			_778_v17 = _rhs116
			_744_v1 = Companion_Default___.Fm23((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cho")).Cardinality())).Times(_743_v0), globalState)
			var _787_v21 _dafny.Map
			_ = _787_v21
			_787_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((false) == (p0), _771_v13)
			_787_v21 = (_787_v21).Update(_778_v17, _771_v13)
			var _788_v22 _dafny.Map
			_ = _788_v22
			_788_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
			_778_v17 = ((_788_v22).Merge(_788_v22)).Contains(_778_v17)
		}
		if p0 {
			var _789_v23 _dafny.Array
			_ = _789_v23
			var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw146
			_789_v23 = _nw146
			var _790_v24 _dafny.Map
			_ = _790_v24
			_790_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_v0, Companion_D0_.Create_DC0_(_743_v0, (_this).F6(), _789_v23))
			var _791_v25 _dafny.Sequence
			_ = _791_v25
			_791_v25 = _dafny.SeqOf(_790_v24)
			var _792_v26 T1
			_ = _792_v26
			var _nw147 *C3 = New_C3_()
			_ = _nw147
			_nw147.Ctor__()
			_792_v26 = _nw147
			var _793_v27 D11
			_ = _793_v27
			_793_v27 = Companion_D11_.Create_DC26_(_791_v25, _792_v26)
			var _pat_let_tv10 = _790_v24
			_ = _pat_let_tv10
			var _pat_let_tv11 = _790_v24
			_ = _pat_let_tv11
			var _pat_let_tv12 = _790_v24
			_ = _pat_let_tv12
			var _pat_let_tv13 = _790_v24
			_ = _pat_let_tv13
			var _794_v28 _dafny.Map
			_ = _794_v28
			_794_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf((_this).F6(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_795_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('c')
			}))), func(_pat_let13_0 D11) D11 {
				return func(_796_dt__update__tmp_h2 D11) D11 {
					return func(_pat_let14_0 _dafny.Sequence) D11 {
						return func(_797_dt__update_hcf42_h0 _dafny.Sequence) D11 {
							return Companion_D11_.Create_DC26_(_797_dt__update_hcf42_h0, (_796_dt__update__tmp_h2).Dtor_cf43())
						}(_pat_let14_0)
					}(_dafny.SeqOf(_pat_let_tv10, _pat_let_tv11, _pat_let_tv12, _pat_let_tv13))
				}(_pat_let13_0)
			}(_793_v27))
			_794_v28 = (_794_v28).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _793_v27)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _793_v27)))
			var _798_v29 _dafny.Sequence
			_ = _798_v29
			_798_v29 = _dafny.SeqOf(p0, p0, !(p0) || (p0), p0)
			_798_v29 = _dafny.Companion_Sequence_.Concatenate(_798_v29, _dafny.Companion_Sequence_.Concatenate(_798_v29, _798_v29))
			_743_v0 = _743_v0
			var _799_v30 bool
			_ = _799_v30
			_799_v30 = true
			var _800_v31 D0
			_ = _800_v31
			_800_v31 = Companion_D0_.Create_DC0_(_743_v0, _dafny.UnicodeSeqOfUtf8Bytes("mqbg"), _789_v23)
			_799_v30 = Companion_Default___.Fm33((_800_v31).Dtor_cf1(), globalState)
			var _801_v32 *C1
			_ = _801_v32
			var _nw148 *C1 = New_C1_()
			_ = _nw148
			_nw148.Ctor__(_dafny.MultiSetFromSeq(_798_v29), Companion_Default___.SafeDivisionInt(_743_v0, _743_v0))
			_801_v32 = _nw148
		} else {
			var _802_v33 *C0
			_ = _802_v33
			var _nw149 *C0 = New_C0_()
			_ = _nw149
			_nw149.Ctor__()
			_802_v33 = _nw149
			_743_v0 = _743_v0
			_743_v0 = _dafny.IntOfInt64(-933)
			var _803_v34 bool
			_ = _803_v34
			_803_v34 = false
			var _804_v35 _dafny.Sequence
			_ = _804_v35
			_804_v35 = _dafny.SeqOf(_803_v34)
			_803_v34 = !(_803_v34) || ((_804_v35).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_803_v34, _803_v34)).Cardinality(), _dafny.IntOfUint32((_804_v35).Cardinality()))).Uint32()).(bool))
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_771_v13), 0))
			_ = _index155
			(_771_v13).ArraySet1(_743_v0, (_index155).Int())
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_771_v13), 0))
			_ = _index156
			(_771_v13).ArraySet1(_743_v0, (_index156).Int())
		}
		var _805_v36 bool
		_ = _805_v36
		_805_v36 = false
		_805_v36 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ybptaos"), _dafny.Companion_Sequence_.Concatenate((_this).F6(), _dafny.UnicodeSeqOfUtf8Bytes("kgjwul")))
		var _806_v37 _dafny.Sequence
		_ = _806_v37
		_806_v37 = _dafny.SeqOf(_dafny.SeqOf(_743_v0))
		var _807_i5 _dafny.Int
		_ = _807_i5
		_807_i5 = _dafny.Zero
		{
			for ((_743_v0).Plus(_743_v0)).Cmp((_dafny.MultiSetFromSeq((_806_v37).Select((Companion_Default___.SafeIndex(_743_v0, _dafny.IntOfUint32((_806_v37).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()) == 0 {
				{
					if (_807_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_807_i5 = (_807_i5).Plus(_dafny.One)
					var _808_v38 _dafny.Sequence
					_ = _808_v38
					_808_v38 = _dafny.UnicodeSeqOfUtf8Bytes("g")
					var _809_v39 _dafny.Set
					_ = _809_v39
					_809_v39 = _dafny.SetOf(_743_v0, (_dafny.Zero).Minus(_743_v0))
					var _810_v40 _dafny.Map
					_ = _810_v40
					_810_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_v0, _809_v39)
					var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_774_v14), 0))
					_ = _index157
					(_774_v14).ArraySet1(((func() _dafny.Set {
						if (_810_v40).Contains(_743_v0) {
							return (_810_v40).Get(_743_v0).(_dafny.Set)
						}
						return _809_v39
					})()).IsSubsetOf(Companion_Default___.Fm20(_805_v36, _808_v38, Companion_Default___.Fm33(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(290))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg37 _dafny.Int) interface{} {
							return coer37(arg37)
						}
					}(func(_811_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('j')
					})), globalState), globalState)), (_index157).Int())
					var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_774_v14), 0))
					_ = _index158
					var _rhs117 bool = _dafny.Companion_Sequence_.IsProperPrefixOf((_this).F6(), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm22(_805_v36, _743_v0, globalState), (_this).F6()))
					_ = _rhs117
					var _rhs118 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_this).F6(), _dafny.UnicodeSeqOfUtf8Bytes("xq"))
					_ = _rhs118
					var _rhs119 bool = _805_v36
					_ = _rhs119
					var _rhs120 bool = _805_v36
					_ = _rhs120
					var _rhs121 _dafny.Int = _743_v0
					_ = _rhs121
					var _lhs85 _dafny.Array = _774_v14
					_ = _lhs85
					var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_774_v14), 0))
					_ = _lhs86
					_805_v36 = _rhs117
					_808_v38 = _rhs118
					(_lhs85).ArraySet1(_rhs119, (_lhs86).Int())
					_805_v36 = _rhs120
					r0 = _rhs121
					var _812_v41 _dafny.MultiSet
					_ = _812_v41
					_812_v41 = _dafny.MultiSetOf(_dafny.IntOfInt64(259), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-880), (_dafny.Zero).Minus(_743_v0)))
					_812_v41 = _812_v41
					var _813_v42 D4
					_ = _813_v42
					_813_v42 = Companion_D4_.Create_DC9_(_771_v13)
					var _814_v43 _dafny.Array
					_ = _814_v43
					var _nwElement0_33 _dafny.Array = _771_v13
					_ = _nwElement0_33
					var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(2))
					_ = _nw150
					(_nw150).ArraySet1(_nwElement0_33, 0)
					(_nw150).ArraySet1((_813_v42).Dtor_cf15(), 1)
					_814_v43 = _nw150
					var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_814_v43), 0))
					_ = _index159
					(_814_v43).ArraySet1(_771_v13, (_index159).Int())
					var _815_v44 _dafny.Map
					_ = _815_v44
					_815_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _743_v0)
					var _816_v45 _dafny.MultiSet
					_ = _816_v45
					_816_v45 = _dafny.MultiSetOf((_774_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_774_v14), 0))).Int()).(bool), p0, false, _805_v36)
					var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_814_v43), 0))
					_ = _index160
					var _nwElement0_34 _dafny.Int = _743_v0
					_ = _nwElement0_34
					var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(24))
					_ = _nw151
					(_nw151).ArraySet1(_nwElement0_34, 0)
					(_nw151).ArraySet1(_743_v0, 1)
					(_nw151).ArraySet1(_743_v0, 2)
					(_nw151).ArraySet1(_743_v0, 3)
					(_nw151).ArraySet1(_743_v0, 4)
					(_nw151).ArraySet1((_dafny.Zero).Minus((_743_v0).Minus(_743_v0)), 5)
					(_nw151).ArraySet1(_743_v0, 6)
					(_nw151).ArraySet1(_743_v0, 7)
					(_nw151).ArraySet1(Companion_Default___.SafeDivisionInt(_743_v0, (_dafny.Zero).Minus(_743_v0)), 8)
					(_nw151).ArraySet1(_743_v0, 9)
					(_nw151).ArraySet1(_743_v0, 10)
					(_nw151).ArraySet1((func() _dafny.Int {
						if _805_v36 {
							return (_815_v44).Cardinality()
						}
						return _743_v0
					})(), 11)
					(_nw151).ArraySet1(_743_v0, 12)
					(_nw151).ArraySet1(_743_v0, 13)
					(_nw151).ArraySet1(_743_v0, 14)
					(_nw151).ArraySet1(((_dafny.SetOf(!(Companion_Default___.Fm33((_this).F6(), globalState)))).Cardinality()).Minus(_dafny.IntOfInt64(902)), 15)
					(_nw151).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_805_v36)).Cardinality()), _dafny.IntOfInt64(351)), 16)
					(_nw151).ArraySet1(_743_v0, 17)
					(_nw151).ArraySet1((_dafny.Zero).Minus(_743_v0), 18)
					(_nw151).ArraySet1(_743_v0, 19)
					(_nw151).ArraySet1(Companion_Default___.Fm14((func() _dafny.Int {
						if (_816_v45).Contains(p0) {
							return (_816_v45).Multiplicity(p0)
						}
						return _743_v0
					})(), _743_v0, globalState), 20)
					(_nw151).ArraySet1(_dafny.IntOfUint32(((_this).F6()).Cardinality()), 21)
					(_nw151).ArraySet1((_dafny.MultiSetOf(!((_774_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_774_v14), 0))).Int()).(bool)))).Cardinality(), 22)
					(_nw151).ArraySet1(_743_v0, 23)
					(_814_v43).ArraySet1(_nw151, (_index160).Int())
					var _817_v46 _dafny.Sequence
					_ = _817_v46
					_817_v46 = _dafny.SeqOf(_805_v36, true)
					var _818_v47 _dafny.Sequence
					_ = _818_v47
					_818_v47 = _dafny.SeqOf((_dafny.IntOfUint32((_817_v46).Cardinality())).Cmp(_743_v0) >= 0)
					var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_774_v14), 0))
					_ = _index161
					(_774_v14).ArraySet1((_818_v47).Select((Companion_Default___.SafeIndex((_744_v1).Select((Companion_Default___.SafeIndex(_743_v0, _dafny.IntOfUint32((_744_v1).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_818_v47).Cardinality()))).Uint32()).(bool), (_index161).Int())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		r0 = _743_v0
		var _819_v48 _dafny.CodePoint
		_ = _819_v48
		_819_v48 = _dafny.CodePoint('j')
		r1 = _819_v48
		return r0, r1
	}
}
func (_this *C5) F6() _dafny.Sequence {
	{
		return _this._f6
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	F4  _dafny.MultiSet
	_f5 _dafny.Array
}

func New_C6_() *C6 {
	_this := C6{}

	_this.F4 = _dafny.EmptyMultiSet
	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C6) Ctor__(f4 _dafny.MultiSet, f5 _dafny.Array) {
	{
		(_this).F4 = f4
		(_this)._f5 = f5
	}
}
func (_this *C6) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return ((_this.F4).Union(_this.F4)).Difference(_this.F4)
	}
}
func (_this *C6) Fm9(globalState *GlobalState) _dafny.Int {
	{
		return (((func() _dafny.Map {
			if false {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(788))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}(func(_820_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality()), false)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), false)
		})()).Merge((func() _dafny.Map {
			if false {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), true)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(440), true)
		})())).Cardinality()
	}
}
func (_this *C6) Fm10(p0 D2, p1 _dafny.Int, p2 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return (Companion_D2_.Create_DC6_(_dafny.IntOfInt64(929), false)).Dtor_cf12()
	}
}
func (_this *C6) M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _821_v0 _dafny.Array
		_ = _821_v0
		var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
		_ = _nw152
		_821_v0 = _nw152
		var _822_v1 _dafny.Sequence
		_ = _822_v1
		_822_v1 = _dafny.UnicodeSeqOfUtf8Bytes("rxl")
		var _823_v2 _dafny.CodePoint
		_ = _823_v2
		_823_v2 = _dafny.CodePoint('v')
		var _824_v3 _dafny.Int
		_ = _824_v3
		_824_v3 = _dafny.IntOfInt64(403)
		var _825_v4 bool
		_ = _825_v4
		_825_v4 = false
		var _826_v5 D1
		_ = _826_v5
		_826_v5 = Companion_D1_.Create_DC2_(true)
		var _827_v6 _dafny.Array
		_ = _827_v6
		var _nwElement0_35 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(656))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_828_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))
		_ = _nwElement0_35
		var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(23))
		_ = _nw153
		(_nw153).ArraySet1(_nwElement0_35, 0)
		(_nw153).ArraySet1(_822_v1, 1)
		(_nw153).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pet"), 2)
		(_nw153).ArraySet1(_822_v1, 3)
		(_nw153).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tpfmyj"), 4)
		(_nw153).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ipd"), 5)
		(_nw153).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("voauqmys"), 6)
		(_nw153).ArraySet1(_822_v1, 7)
		(_nw153).ArraySet1(_822_v1, 8)
		(_nw153).ArraySet1(_822_v1, 9)
		(_nw153).ArraySet1(_822_v1, 10)
		(_nw153).ArraySet1(Companion_Default___.Fm11(_823_v2, _824_v3, _825_v4, _826_v5, globalState), 11)
		(_nw153).ArraySet1(_822_v1, 12)
		(_nw153).ArraySet1(_822_v1, 13)
		(_nw153).ArraySet1(_822_v1, 14)
		(_nw153).ArraySet1(_822_v1, 15)
		(_nw153).ArraySet1(_822_v1, 16)
		(_nw153).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("a"), 17)
		(_nw153).ArraySet1(_822_v1, 18)
		(_nw153).ArraySet1(_822_v1, 19)
		(_nw153).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("usguexh"), 20)
		(_nw153).ArraySet1(_822_v1, 21)
		(_nw153).ArraySet1(_822_v1, 22)
		_827_v6 = _nw153
		var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_821_v0), 0))
		_ = _index162
		(_821_v0).ArraySet1(_827_v6, (_index162).Int())
		var _829_v7 D2
		_ = _829_v7
		_829_v7 = Companion_D2_.Create_DC6_(_824_v3, _825_v4)
		var _830_v8 _dafny.Sequence
		_ = _830_v8
		_830_v8 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(166))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_831_v3 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
			return func(_832_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_831_v3, _dafny.IntOfInt64(294))
			}
		})(_824_v3)))).Cardinality()))
		var _833_v9 _dafny.Sequence
		_ = _833_v9
		_833_v9 = _dafny.SeqOf(_830_v8, _830_v8, _830_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(906))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}((func(_834_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_835_i2 _dafny.Int) _dafny.Int {
				return _834_v3
			}
		})(_824_v3))), _830_v8)
		var _836_v10 _dafny.MultiSet
		_ = _836_v10
		_836_v10 = _dafny.MultiSetOf(_824_v3)
		var _837_v11 _dafny.Map
		_ = _837_v11
		_837_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_825_v4, _836_v10)
		var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_821_v0), 0))
		_ = _index163
		var _rhs122 _dafny.Array = _827_v6
		_ = _rhs122
		var _rhs123 bool = (_this).Fm10(_829_v7, (_dafny.IntOfUint32(((_833_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm12(_824_v3, _824_v3, globalState)).Cardinality()), _dafny.IntOfUint32((_833_v9).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Plus(_dafny.IntOfInt64(933)), (func() _dafny.MultiSet {
			if (_837_v11).Contains(_825_v4) {
				return (_837_v11).Get(_825_v4).(_dafny.MultiSet)
			}
			return _836_v10
		})(), globalState)
		_ = _rhs123
		var _rhs124 bool = false
		_ = _rhs124
		var _lhs87 _dafny.Array = _821_v0
		_ = _lhs87
		var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_821_v0), 0))
		_ = _lhs88
		(_lhs87).ArraySet1(_rhs122, (_lhs88).Int())
		r0 = _rhs123
		_825_v4 = _rhs124
		var _838_i3 _dafny.Int
		_ = _838_i3
		_838_i3 = _dafny.Zero
		{
			for !(_dafny.Companion_Sequence_.IsProperPrefixOf(_822_v1, _dafny.UnicodeSeqOfUtf8Bytes("omtjm"))) {
				{
					if (_838_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_838_i3 = (_838_i3).Plus(_dafny.One)
					var _rhs125 D1 = _826_v5
					_ = _rhs125
					var _rhs126 _dafny.Int = (_dafny.Zero).Minus(((_824_v3).Plus(_824_v3)).Minus(_dafny.IntOfInt64(199)))
					_ = _rhs126
					_826_v5 = _rhs125
					_824_v3 = _rhs126
					var _839_v12 _dafny.Sequence
					_ = _839_v12
					_839_v12 = _dafny.SeqOf(false, _825_v4, true, _825_v4)
					_824_v3 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_839_v12).Cardinality()), _824_v3)
					var _840_v13 _dafny.Int
					_ = _840_v13
					var _841_v14 _dafny.Map
					_ = _841_v14
					var _842_v15 bool
					_ = _842_v15
					var _out14 _dafny.Int
					_ = _out14
					var _out15 _dafny.Map
					_ = _out15
					var _out16 bool
					_ = _out16
					_out14, _out15, _out16 = (_this).M2(globalState)
					_840_v13 = _out14
					_841_v14 = _out15
					_842_v15 = _out16
					if _842_v15 {
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen(((_this).F5()), 0))
						_ = _index164
						((_this).F5()).ArraySet1((_this).Fm9(globalState), (_index164).Int())
						var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen(((_this).F5()), 0))
						_ = _index165
						((_this).F5()).ArraySet1((_829_v7).Dtor_cf11(), (_index165).Int())
						_830_v8 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_830_v8, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg42 _dafny.Int) interface{} {
								return coer42(arg42)
							}
						}(func(_843_i4 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(174)
						})), (Companion_Default___.SafeIndex(_840_v13, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg43 _dafny.Int) interface{} {
								return coer43(arg43)
							}
						}(func(_844_i4 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(174)
						}))).Cardinality()))).Uint32(), _dafny.IntOfInt64(844))), (Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_836_v10).Contains(((_this).F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen(((_this).F5()), 0))).Int()).(_dafny.Int)) {
								return (_836_v10).Multiplicity(((_this).F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen(((_this).F5()), 0))).Int()).(_dafny.Int))
							}
							return _dafny.IntOfInt64(-687)
						})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_830_v8, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}(func(_845_i4 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(174)
						})), (Companion_Default___.SafeIndex(_840_v13, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}(func(_846_i4 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(174)
						}))).Cardinality()))).Uint32(), _dafny.IntOfInt64(844)))).Cardinality()))).Uint32(), _dafny.IntOfInt64(-367))
						var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen(((_this).F5()), 0))
						_ = _index166
						((_this).F5()).ArraySet1(_840_v13, (_index166).Int())
						var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_827_v6), 0))
						_ = _index167
						(_827_v6).ArraySet1(_822_v1, (_index167).Int())
						var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_827_v6), 0))
						_ = _index168
						(_827_v6).ArraySet1((func() _dafny.Sequence {
							if !(_842_v15) {
								return _dafny.UnicodeSeqOfUtf8Bytes("dgiljws")
							}
							return _822_v1
						})(), (_index168).Int())
						_840_v13 = (_830_v8).Select((Companion_Default___.SafeIndex(((_this).F5()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen(((_this).F5()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_830_v8).Cardinality()))).Uint32()).(_dafny.Int)
					} else {
						var _847_v16 _dafny.Array
						_ = _847_v16
						var _nw154 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
						_ = _nw154
						_847_v16 = _nw154
						var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_847_v16), 0))
						_ = _index169
						(_847_v16).ArraySet1(_842_v15, (_index169).Int())
						var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_847_v16), 0))
						_ = _index170
						(_847_v16).ArraySet1(!((_this).Fm10(_829_v7, _824_v3, _836_v10, globalState)), (_index170).Int())
						_830_v8 = _dafny.Companion_Sequence_.Update(_830_v8, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_822_v1).Cardinality())), _dafny.IntOfUint32((_830_v8).Cardinality()))).Uint32(), _824_v3)
						_824_v3 = _dafny.IntOfUint32((_830_v8).Cardinality())
						var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_827_v6), 0))
						_ = _index171
						(_827_v6).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rvrjpar"), (_index171).Int())
						var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_827_v6), 0))
						_ = _index172
						(_827_v6).ArraySet1(_822_v1, (_index172).Int())
						var _848_v17 *C1
						_ = _848_v17
						var _nw155 *C1 = New_C1_()
						_ = _nw155
						_nw155.Ctor__(_this.F4, _840_v13)
						_848_v17 = _nw155
					}
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _849_i5 _dafny.Int
		_ = _849_i5
		_849_i5 = _dafny.Zero
		{
			for _825_v4 {
				{
					if (_849_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_849_i5 = (_849_i5).Plus(_dafny.One)
					var _850_v18 _dafny.Array
					_ = _850_v18
					var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
					_ = _nw156
					_850_v18 = _nw156
					var _851_v19 _dafny.Map
					_ = _851_v19
					_851_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_824_v3, Companion_D0_.Create_DC0_(_824_v3, _822_v1, _850_v18))
					var _852_v20 _dafny.Sequence
					_ = _852_v20
					_852_v20 = _dafny.SeqOf(_851_v19, _851_v19, _851_v19)
					var _853_v21 T1
					_ = _853_v21
					var _nw157 *C1 = New_C1_()
					_ = _nw157
					_nw157.Ctor__(_this.F4, _824_v3)
					_853_v21 = _nw157
					var _854_v22 D11
					_ = _854_v22
					_854_v22 = Companion_D11_.Create_DC26_(_852_v20, _853_v21)
					var _855_v23 _dafny.MultiSet
					_ = _855_v23
					_855_v23 = _dafny.MultiSetOf(_854_v22)
					var _856_v24 _dafny.Sequence
					_ = _856_v24
					_856_v24 = _dafny.SeqOf(_825_v4)
					var _857_v25 _dafny.MultiSet
					_ = _857_v25
					_857_v25 = _dafny.MultiSetOf(_856_v24, Companion_Default___.Fm12(_824_v3, _824_v3, globalState))
					var _rhs127 _dafny.MultiSet = _855_v23
					_ = _rhs127
					var _rhs128 _dafny.Int = Companion_Default___.SafeDivisionInt((_857_v25).Cardinality(), (func() _dafny.Int {
						if _825_v4 {
							return _824_v3
						}
						return _dafny.IntOfInt64(579)
					})())
					_ = _rhs128
					_855_v23 = _rhs127
					_824_v3 = _rhs128
					var _858_v26 _dafny.Map
					_ = _858_v26
					_858_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm33(_822_v1, globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_856_v24, _824_v3))
					var _859_v27 _dafny.Map
					_ = _859_v27
					_859_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_856_v24, _824_v3)
					_858_v26 = (_858_v26).Update(_825_v4, _859_v27)
					var _860_v28 *C0
					_ = _860_v28
					var _nw158 *C0 = New_C0_()
					_ = _nw158
					_nw158.Ctor__()
					_860_v28 = _nw158
					var _861_v29 D7
					_ = _861_v29
					_861_v29 = Companion_D7_.Create_DC15_(_825_v4, _824_v3, _824_v3, _824_v3, _860_v28)
					var _862_v30 D10
					_ = _862_v30
					_862_v30 = Companion_D10_.Create_DC22_()
					var _863_v31 _dafny.Sequence
					_ = _863_v31
					_863_v31 = _dafny.SeqOf(_862_v30)
					var _864_v32 _dafny.Map
					_ = _864_v32
					_864_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_825_v4, (_dafny.Zero).Minus(_824_v3))
					var _865_v33 D9
					_ = _865_v33
					_865_v33 = Companion_D9_.Create_DC19_(_824_v3, (_864_v32).Cardinality())
					var _866_v34 _dafny.Array
					_ = _866_v34
					var _nwElement0_36 _dafny.Int = _824_v3
					_ = _nwElement0_36
					var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(18))
					_ = _nw159
					(_nw159).ArraySet1(_nwElement0_36, 0)
					(_nw159).ArraySet1((_861_v29).Dtor_cf22(), 1)
					(_nw159).ArraySet1((_dafny.IntOfInt64(697)).Times(_824_v3), 2)
					(_nw159).ArraySet1(_824_v3, 3)
					(_nw159).ArraySet1(_dafny.IntOfInt64(-819), 4)
					(_nw159).ArraySet1(_824_v3, 5)
					(_nw159).ArraySet1(_824_v3, 6)
					(_nw159).ArraySet1((_dafny.Zero).Minus(_824_v3), 7)
					(_nw159).ArraySet1(_824_v3, 8)
					(_nw159).ArraySet1(_dafny.IntOfInt64(-642), 9)
					(_nw159).ArraySet1(_824_v3, 10)
					(_nw159).ArraySet1(_824_v3, 11)
					(_nw159).ArraySet1(_824_v3, 12)
					(_nw159).ArraySet1(_824_v3, 13)
					(_nw159).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_862_v30, _862_v30), _863_v31)).Cardinality()), 14)
					(_nw159).ArraySet1(Companion_Default___.SafeModuloInt((_865_v33).Dtor_cf29(), _824_v3), 15)
					(_nw159).ArraySet1((_this).Fm9(globalState), 16)
					(_nw159).ArraySet1((_830_v8).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.Fm24(_824_v3, true, _825_v4, _824_v3, globalState)), _dafny.IntOfUint32((_830_v8).Cardinality()))).Uint32()).(_dafny.Int), 17)
					_866_v34 = _nw159
					_866_v34 = _866_v34
					_830_v8 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}((func(_867_v3 _dafny.Int, _868_v4 bool) func(_dafny.Int) _dafny.Int {
						return func(_869_i6 _dafny.Int) _dafny.Int {
							return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_867_v3, _868_v4)).Cardinality()).Minus(_dafny.IntOfInt64(668))
						}
					})(_824_v3, _825_v4)))
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _870_v35 _dafny.Array
		_ = _870_v35
		var _nw160 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
		_ = _nw160
		_870_v35 = _nw160
		var _rhs129 bool = (_829_v7).Dtor_cf12()
		_ = _rhs129
		var _rhs130 _dafny.Array = _870_v35
		_ = _rhs130
		r0 = _rhs129
		_870_v35 = _rhs130
		_824_v3 = Companion_Default___.SafeModuloInt(_824_v3, _824_v3)
		var _871_v36 *C1
		_ = _871_v36
		var _nw161 *C1 = New_C1_()
		_ = _nw161
		_nw161.Ctor__(_this.F4, _824_v3)
		_871_v36 = _nw161
		r0 = _825_v4
		var _872_v37 _dafny.Map
		_ = _872_v37
		_872_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-879)), _this.F4)
		r1 = (_872_v37)
		return r0, r1
	}
}
func (_this *C6) M2(globalState *GlobalState) (_dafny.Int, _dafny.Map, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 bool = false
		_ = r2
		var _873_v0 _dafny.Int
		_ = _873_v0
		_873_v0 = _dafny.IntOfInt64(392)
		r0 = _873_v0
		var _874_v1 bool
		_ = _874_v1
		_874_v1 = false
		var _875_v2 D10
		_ = _875_v2
		_875_v2 = Companion_D10_.Create_DC21_(_dafny.MultiSetOf(_874_v1))
		_875_v2 = _875_v2
		r2 = !(_874_v1)
		var _876_v3 _dafny.Array
		_ = _876_v3
		var _nw162 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
		_ = _nw162
		_876_v3 = _nw162
		var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))
		_ = _index173
		(_876_v3).ArraySet1(_874_v1, (_index173).Int())
		var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))
		_ = _index174
		(_876_v3).ArraySet1(_874_v1, (_index174).Int())
		var _hi6 _dafny.Int = _873_v0
		_ = _hi6
		for _877_i0 := _873_v0; _877_i0.Cmp(_hi6) < 0; _877_i0 = _877_i0.Plus(_dafny.One) {
			r0 = _877_i0
			_873_v0 = (_877_i0).Minus(_877_i0)
			var _878_v4 _dafny.Array
			_ = _878_v4
			var _nwElement0_37 _dafny.Array = (_this).F5()
			_ = _nwElement0_37
			var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(2))
			_ = _nw163
			(_nw163).ArraySet1(_nwElement0_37, 0)
			(_nw163).ArraySet1((_this).F5(), 1)
			_878_v4 = _nw163
			var _879_v5 D0
			_ = _879_v5
			_879_v5 = Companion_D0_.Create_DC0_(_873_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(79))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}(func(_880_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			})), _878_v4)
			r2 = Companion_Default___.Fm33((_879_v5).Dtor_cf1(), globalState)
			if (_876_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))).Int()).(bool) {
				var _881_v6 _dafny.MultiSet
				_ = _881_v6
				_881_v6 = _dafny.MultiSetOf(_dafny.SeqOf(_877_i0, _dafny.IntOfInt64(901)))
				_881_v6 = (_881_v6).Union((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(776))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_882_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_883_i2 _dafny.Int) _dafny.Int {
						return _882_v0
					}
				})(_873_v0))))).Union(_881_v6))
				var _884_v7 _dafny.CodePoint
				_ = _884_v7
				_884_v7 = _dafny.CodePoint('t')
				var _885_v8 _dafny.Sequence
				_ = _885_v8
				_885_v8 = _dafny.UnicodeSeqOfUtf8Bytes("cu")
				var _886_v9 _dafny.Set
				_ = _886_v9
				_886_v9 = _dafny.SetOf(_885_v8)
				var _887_v10 _dafny.Sequence
				_ = _887_v10
				_887_v10 = _dafny.SeqOf(((_dafny.Zero).Minus(_873_v0)).Cmp((_886_v9).Cardinality()) != 0)
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))
				_ = _index175
				var _rhs131 bool = (_887_v10).Select((Companion_Default___.SafeIndex((_877_i0).Minus(_877_i0), _dafny.IntOfUint32((_887_v10).Cardinality()))).Uint32()).(bool)
				_ = _rhs131
				var _rhs132 _dafny.CodePoint = _dafny.CodePoint('f')
				_ = _rhs132
				var _rhs133 _dafny.Int = _873_v0
				_ = _rhs133
				var _lhs89 _dafny.Array = _876_v3
				_ = _lhs89
				var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))
				_ = _lhs90
				(_lhs89).ArraySet1(_rhs131, (_lhs90).Int())
				_884_v7 = _rhs132
				r0 = _rhs133
				_884_v7 = _884_v7
				var _888_v11 _dafny.Array
				_ = _888_v11
				var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw164
				_888_v11 = _nw164
				var _889_v12 _dafny.Sequence
				_ = _889_v12
				_889_v12 = _dafny.SeqOf(_877_i0)
				var _890_v13 _dafny.Map
				_ = _890_v13
				_890_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_876_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))).Int()).(bool), _876_v3)
				var _891_v14 _dafny.Array
				_ = _891_v14
				_891_v14 = _876_v3
				var _892_v15 _dafny.Array
				_ = _892_v15
				var _nwElement0_38 _dafny.Array = _876_v3
				_ = _nwElement0_38
				var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(3))
				_ = _nw165
				(_nw165).ArraySet1(_nwElement0_38, 0)
				(_nw165).ArraySet1(_876_v3, 1)
				(_nw165).ArraySet1((func() _dafny.Array {
					if (_890_v13).Contains(_874_v1) {
						return (_890_v13).Get(_874_v1).(_dafny.Array)
					}
					return (_891_v14)
				})(), 2)
				_892_v15 = _nw165
				var _893_v16 _dafny.Map
				_ = _893_v16
				_893_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_889_v12, _892_v15)
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_888_v11), 0))
				_ = _index176
				(_888_v11).ArraySet1((func() _dafny.Array {
					if (_893_v16).Contains(_889_v12) {
						return (_893_v16).Get(_889_v12).(_dafny.Array)
					}
					return _892_v15
				})(), (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_888_v11), 0))
				_ = _index177
				(_888_v11).ArraySet1(_892_v15, (_index177).Int())
				_874_v1 = (_876_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))).Int()).(bool)
			} else {
				_873_v0 = _877_i0
				var _894_v17 D3
				_ = _894_v17
				_894_v17 = Companion_D3_.Create_DC8_((func() _dafny.Int {
					if _874_v1 {
						return (_dafny.Zero).Minus(_873_v0)
					}
					return _877_i0
				})())
				_894_v17 = _894_v17
				var _895_v18 _dafny.CodePoint
				_ = _895_v18
				_895_v18 = _dafny.CodePoint('w')
				_895_v18 = _895_v18
				var _896_v19 _dafny.Sequence
				_ = _896_v19
				_896_v19 = _dafny.UnicodeSeqOfUtf8Bytes("xjisbuee")
				var _897_v20 _dafny.Sequence
				_ = _897_v20
				_897_v20 = _dafny.SeqOf(_896_v19, _896_v19, _896_v19)
				var _898_v21 _dafny.Sequence
				_ = _898_v21
				_898_v21 = _897_v20
				var _899_v22 _dafny.Map
				_ = _899_v22
				_899_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_873_v0, _dafny.IntOfUint32((_898_v21).Cardinality()))
				_899_v22 = (_899_v22).Update(_873_v0, _873_v0)
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen(((_this).F5()), 0))
				_ = _index178
				((_this).F5()).ArraySet1(_877_i0, (_index178).Int())
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen(((_this).F5()), 0))
				_ = _index179
				((_this).F5()).ArraySet1(_877_i0, (_index179).Int())
			}
		}
		var _900_v23 D4
		_ = _900_v23
		_900_v23 = Companion_D4_.Create_DC10_(_874_v1)
		var _901_v24 _dafny.Sequence
		_ = _901_v24
		_901_v24 = _dafny.UnicodeSeqOfUtf8Bytes("e")
		_874_v1 = (func() bool {
			if (Companion_Default___.Fm16(_900_v23, globalState)).IsSubsetOf(_dafny.SetOf(_874_v1, (_876_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))).Int()).(bool))) {
				return Companion_Default___.Fm33(_901_v24, globalState)
			}
			return (_876_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))).Int()).(bool)
		})()
		var _902_v25 _dafny.Sequence
		_ = _902_v25
		_902_v25 = _dafny.SeqOf(_876_v3)
		var _903_v27 _dafny.CodePoint
		_ = _903_v27
		_903_v27 = _dafny.CodePoint('v')
		var _904_v28 _dafny.Map
		_ = _904_v28
		_904_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm33(_901_v24, globalState), (func() _dafny.Map {
			var _coll27 = _dafny.NewMapBuilder()
			_ = _coll27
			for _iter31 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('g'), _903_v27)).Elements()); ; {
				_compr_27, _ok31 := _iter31()
				if !_ok31 {
					break
				}
				var _905_v26 _dafny.CodePoint
				_905_v26 = interface{}(_compr_27).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('g'), _903_v27), _905_v26) {
					_coll27.Add(_905_v26, _873_v0)
				}
			}
			return _coll27.ToMap()
		}()).Cardinality())
		r0 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_902_v25).Cardinality()), _873_v0), (func() _dafny.Int {
			if (_904_v28).Contains(Companion_Default___.Fm33(_901_v24, globalState)) {
				return (_904_v28).Get(Companion_Default___.Fm33(_901_v24, globalState)).(_dafny.Int)
			}
			return _873_v0
		})())
		var _906_v29 _dafny.Sequence
		_ = _906_v29
		_906_v29 = _dafny.SeqOf((_dafny.Zero).Minus(_873_v0), _dafny.IntOfUint32((_901_v24).Cardinality()))
		var _907_v30 _dafny.Map
		_ = _907_v30
		_907_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC5_(_906_v29), _dafny.IntOfInt64(101))
		var _908_v31 _dafny.Map
		_ = _908_v31
		_908_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_907_v30).Merge(_907_v30), _874_v1)
		r1 = _908_v31
		r2 = (_876_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_876_v3), 0))).Int()).(bool)
		return r0, r1, r2
	}
}
func (_this *C6) F5() _dafny.Array {
	{
		return _this._f5
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f3 bool
	_f2 _dafny.Array
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f3 = false
	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C7) Ctor__(f2 _dafny.Array, f3 bool) {
	{
		(_this)._f2 = f2
		(_this)._f3 = f3
	}
}
func (_this *C7) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf((_this).F3(), ((_dafny.Zero).Minus((_dafny.MultiSetOf((_this).F3(), (_this).F3(), (_this).F3())).Cardinality())).Cmp(_dafny.IntOfInt64(-476)) > 0)
	}
}
func (_this *C7) Fm6(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
			var _coll28 = _dafny.NewMapBuilder()
			_ = _coll28
			for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(747), _dafny.IntOfInt64(-512))); ; {
				_compr_28, _ok32 := _iter32()
				if !_ok32 {
					break
				}
				var _909_v0 _dafny.Int
				_909_v0 = interface{}(_compr_28).(_dafny.Int)
				if ((_dafny.IntOfInt64(747)).Cmp(_909_v0) <= 0) && ((_909_v0).Cmp(_dafny.IntOfInt64(-512)) < 0) {
					_coll28.Add((_909_v0).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gtxa")).Cardinality()))), (_this).F3())
				}
			}
			return _coll28.ToMap()
		}()).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cwbgbfuq")).Cardinality()))).Cardinality())).Cardinality()))).Cardinality()), (_this).F3())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(512), (_this).F3()))).Cardinality()).Minus(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), !((_this).F3()))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-762))).Cardinality()))))
	}
}
func (_this *C7) Fm7(p0 D2, p1 _dafny.Int, p2 D2, p3 D1, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("xxbvgut")
	}
}
func (_this *C7) M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _910_v0 _dafny.Int
		_ = _910_v0
		_910_v0 = _dafny.IntOfInt64(-934)
		_910_v0 = _910_v0
		var _911_v1 _dafny.Array
		_ = _911_v1
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_15
		var _nw166 _dafny.Array
		_ = _nw166
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw166 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) bool = func(_912_i0 _dafny.Int) bool {
				return (_this).F3()
			}
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw166 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw166).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw166).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_911_v1 = _nw166
		var _913_v2 _dafny.CodePoint
		_ = _913_v2
		_913_v2 = _dafny.CodePoint('o')
		var _914_v3 _dafny.Map
		_ = _914_v3
		_914_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _913_v2)
		var _915_v4 _dafny.MultiSet
		_ = _915_v4
		_915_v4 = _dafny.MultiSetOf(_914_v3)
		var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
		_ = _index180
		(_911_v1).ArraySet1(((_915_v4).Update((_914_v3).Update((_this).F3(), _913_v2), Companion_Default___.Abs(_dafny.IntOfInt64(161)))).IsDisjointFrom(_915_v4), (_index180).Int())
		var _916_v5 _dafny.MultiSet
		_ = _916_v5
		_916_v5 = _dafny.MultiSetOf(_910_v0, _910_v0, _910_v0)
		var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
		_ = _index181
		(_911_v1).ArraySet1(Companion_Default___.Fm8(_916_v5, globalState), (_index181).Int())
		if (func() bool {
			if (_this).F3() {
				return (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)
			}
			return (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)
		})() {
			r0 = (_this).F3()
			var _917_v6 _dafny.Sequence
			_ = _917_v6
			_917_v6 = _dafny.UnicodeSeqOfUtf8Bytes("prubmuoud")
			var _918_v7 _dafny.Sequence
			_ = _918_v7
			_918_v7 = _dafny.SeqOf(_910_v0)
			var _919_v8 D2
			_ = _919_v8
			_919_v8 = Companion_D2_.Create_DC5_(_918_v7)
			var _920_v9 D1
			_ = _920_v9
			_920_v9 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_910_v0), _913_v2, false, _dafny.IntOfUint32((_917_v6).Cardinality()))
			var _921_v10 *C2
			_ = _921_v10
			var _nw167 *C2 = New_C2_()
			_ = _nw167
			_nw167.Ctor__(_917_v6, (_this).Fm7(_919_v8, _910_v0, _919_v8, _920_v9, globalState))
			_921_v10 = _nw167
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
			_ = _index182
			(_911_v1).ArraySet1((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), (_index182).Int())
			var _922_v11 _dafny.Map
			_ = _922_v11
			_922_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_913_v2, _910_v0)
			var _923_v12 _dafny.Map
			_ = _923_v12
			_923_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_913_v2, ((_922_v11).Merge(_922_v11)).Cardinality())
			_923_v12 = (_923_v12).Update(Companion_Default___.Fm30(_917_v6, _910_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(384))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}((func(_924_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_925_i1 _dafny.Int) _dafny.CodePoint {
					return _924_v2
				}
			})(_913_v2)))).Cardinality()), globalState), _910_v0)
			_917_v6 = _917_v6
		} else {
			var _926_v13 _dafny.Sequence
			_ = _926_v13
			_926_v13 = _dafny.UnicodeSeqOfUtf8Bytes("snuoik")
			var _927_v14 *C5
			_ = _927_v14
			var _nw168 *C5 = New_C5_()
			_ = _nw168
			_nw168.Ctor__(_926_v13)
			_927_v14 = _nw168
			_910_v0 = _910_v0
			var _928_v15 D1
			_ = _928_v15
			_928_v15 = Companion_D1_.Create_DC2_((_this).F3())
			var _929_v16 _dafny.MultiSet
			_ = _929_v16
			_929_v16 = _dafny.MultiSetOf((_this).F3(), (_928_v15).Dtor_cf4())
			var _930_v17 *C1
			_ = _930_v17
			var _nw169 *C1 = New_C1_()
			_ = _nw169
			_nw169.Ctor__(_929_v16, _910_v0)
			_930_v17 = _nw169
			if (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool) {
				var _931_v18 _dafny.Array
				_ = _931_v18
				var _nw170 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw170
				_931_v18 = _nw170
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_931_v18), 0))
				_ = _index183
				(_931_v18).ArraySet1(_930_v17.F11, (_index183).Int())
				var _932_v19 _dafny.Sequence
				_ = _932_v19
				_932_v19 = _dafny.SeqOf(_930_v17.F11)
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_931_v18), 0))
				_ = _index184
				(_931_v18).ArraySet1((func() _dafny.Int {
					if (_dafny.IntOfUint32((_932_v19).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-709))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg50 _dafny.Int) interface{} {
							return coer50(arg50)
						}
					}((func(_933_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_934_i2 _dafny.Int) _dafny.CodePoint {
							return _933_v2
						}
					})(_913_v2)))).Cardinality())) > 0 {
						return _930_v17.F11
					}
					return (_910_v0).Plus((_dafny.MultiSetOf(_930_v17.F11, _930_v17.F11, _910_v0)).Cardinality())
				})(), (_index184).Int())
				(_930_v17).F11 = (_930_v17.F11).Plus(Companion_Default___.SafeDivisionInt((_931_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_931_v18), 0))).Int()).(_dafny.Int), _930_v17.F11))
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen(((_this).F2()), 0))
				_ = _index185
				((_this).F2()).ArraySet1(_931_v18, (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen(((_this).F2()), 0))
				_ = _index186
				((_this).F2()).ArraySet1(_931_v18, (_index186).Int())
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
				_ = _index187
				(_911_v1).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate((_927_v14).F6(), (_927_v14).F6()), (_927_v14).F6()), (_index187).Int())
				r0 = (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)
			} else {
				var _935_v20 _dafny.Array
				_ = _935_v20
				var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
				_ = _nw171
				_935_v20 = _nw171
				var _936_v21 _dafny.Sequence
				_ = _936_v21
				_936_v21 = _dafny.SeqOf((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool))
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_935_v20), 0))
				_ = _index188
				(_935_v20).ArraySet1(_dafny.Companion_Sequence_.Update(_936_v21, (Companion_Default___.SafeIndex(_930_v17.F11, _dafny.IntOfUint32((_936_v21).Cardinality()))).Uint32(), (_936_v21).Select((Companion_Default___.SafeIndex(_910_v0, _dafny.IntOfUint32((_936_v21).Cardinality()))).Uint32()).(bool)), (_index188).Int())
				var _937_v22 _dafny.Array
				_ = _937_v22
				var _nwElement0_39 _dafny.Int = (_dafny.Zero).Minus(_930_v17.F11)
				_ = _nwElement0_39
				var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(8))
				_ = _nw172
				(_nw172).ArraySet1(_nwElement0_39, 0)
				(_nw172).ArraySet1((func() _dafny.Int {
					if (_916_v5).Contains(_930_v17.F11) {
						return (_916_v5).Multiplicity(_930_v17.F11)
					}
					return _dafny.IntOfUint32((_936_v21).Cardinality())
				})(), 1)
				(_nw172).ArraySet1(_930_v17.F11, 2)
				(_nw172).ArraySet1(_930_v17.F11, 3)
				(_nw172).ArraySet1(_910_v0, 4)
				(_nw172).ArraySet1((_927_v14).Fm13((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), _930_v17.F11, (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), globalState), 5)
				(_nw172).ArraySet1(_930_v17.F11, 6)
				(_nw172).ArraySet1(_930_v17.F11, 7)
				_937_v22 = _nw172
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen(((_this).F2()), 0))
				_ = _index189
				((_this).F2()).ArraySet1(_937_v22, (_index189).Int())
				var _938_v23 _dafny.Map
				_ = _938_v23
				_938_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_926_v13).Cardinality()), _929_v16)
				var _939_v24 _dafny.Map
				_ = _939_v24
				_939_v24 = _938_v23
				var _940_v25 _dafny.Map
				_ = _940_v25
				_940_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), _930_v17.F11)
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_935_v20), 0))
				_ = _index190
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen(((_this).F2()), 0))
				_ = _index191
				var _rhs134 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_936_v21, _dafny.Companion_Sequence_.Update(_936_v21, (Companion_Default___.SafeIndex((_940_v25).Cardinality(), _dafny.IntOfUint32((_936_v21).Cardinality()))).Uint32(), !((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool))))
				_ = _rhs134
				var _rhs135 bool = (_this).F3()
				_ = _rhs135
				var _rhs136 _dafny.CodePoint = _913_v2
				_ = _rhs136
				var _rhs137 _dafny.Array = _937_v22
				_ = _rhs137
				var _rhs138 _dafny.Map = _939_v24
				_ = _rhs138
				var _lhs91 _dafny.Array = _935_v20
				_ = _lhs91
				var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_935_v20), 0))
				_ = _lhs92
				var _lhs93 _dafny.Array = (_this).F2()
				_ = _lhs93
				var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen(((_this).F2()), 0))
				_ = _lhs94
				(_lhs91).ArraySet1(_rhs134, (_lhs92).Int())
				r0 = _rhs135
				_913_v2 = _rhs136
				(_lhs93).ArraySet1(_rhs137, (_lhs94).Int())
				_939_v24 = _rhs138
				var _941_v26 _dafny.Map
				_ = _941_v26
				_941_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), !(!(!((_this).F3()))))
				_941_v26 = (_941_v26).Update((_this).F3(), (_this).F3())
				_926_v13 = _dafny.Companion_Sequence_.Concatenate((_927_v14).F6(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(123))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_942_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_943_i3 _dafny.Int) _dafny.CodePoint {
						return _942_v2
					}
				})(_913_v2))))
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
				_ = _index192
				(_911_v1).ArraySet1((func() bool {
					if (_941_v26).Contains((func() bool {
						if (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool) {
							return (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)
						}
						return (_this).F3()
					})()) {
						return (_941_v26).Get((func() bool {
							if (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool) {
								return (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)
							}
							return (_this).F3()
						})()).(bool)
					}
					return true
				})(), (_index192).Int())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
				_ = _index193
				(_911_v1).ArraySet1(!(false), (_index193).Int())
			}
			var _944_v27 _dafny.Sequence
			_ = _944_v27
			_944_v27 = _dafny.SeqOf(true)
			r0 = (_944_v27).Select((Companion_Default___.SafeIndex(_930_v17.F11, _dafny.IntOfUint32((_944_v27).Cardinality()))).Uint32()).(bool)
		}
		_913_v2 = _913_v2
		var _945_v28 _dafny.Map
		_ = _945_v28
		_945_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v0, _916_v5)
		var _946_i4 _dafny.Int
		_ = _946_i4
		_946_i4 = _dafny.Zero
		{
			for (Companion_Default___.Fm31((func() _dafny.Map {
				var _coll31 = _dafny.NewMapBuilder()
				_ = _coll31
				for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(213), _dafny.IntOfInt64(546))); ; {
					_compr_31, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _973_v29 _dafny.Int
					_973_v29 = interface{}(_compr_31).(_dafny.Int)
					if ((_dafny.IntOfInt64(213)).Cmp(_973_v29) <= 0) && ((_973_v29).Cmp(_dafny.IntOfInt64(546)) < 0) {
						_coll31.Add(Companion_Default___.SafeDivisionInt(_973_v29, _910_v0), _910_v0)
					}
				}
				return _coll31.ToMap()
			}()).Cardinality(), globalState)).IsProperSubsetOf((func() _dafny.MultiSet {
				if (_945_v28).Contains(_910_v0) {
					return (_945_v28).Get(_910_v0).(_dafny.MultiSet)
				}
				return _916_v5
			})()) {
				{
					if (_946_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_946_i4 = (_946_i4).Plus(_dafny.One)
					r0 = (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)
					if (_this).F3() {
						_910_v0 = _910_v0
						var _947_v30 *C4
						_ = _947_v30
						var _nw173 *C4 = New_C4_()
						_ = _nw173
						_nw173.Ctor__((_dafny.Zero).Minus(_910_v0))
						_947_v30 = _nw173
						var _948_v31 _dafny.Map
						_ = _948_v31
						_948_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_947_v30.F7, _dafny.IntOfInt64(257)), (_this).F3())
						var _949_v32 _dafny.Sequence
						_ = _949_v32
						_949_v32 = _dafny.SeqOf(_dafny.IntOfInt64(-445), _910_v0)
						_948_v31 = ((_948_v31).Update(_dafny.IntOfUint32((_949_v32).Cardinality()), (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool))).Merge(_948_v31)
						var _950_v33 _dafny.Array
						_ = _950_v33
						var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
						_ = _nw174
						_950_v33 = _nw174
						var _951_v34 _dafny.Set
						_ = _951_v34
						_951_v34 = _dafny.SetOf(_910_v0)
						var _952_v35 D1
						_ = _952_v35
						_952_v35 = Companion_D1_.Create_DC4_((_951_v34).Cardinality(), _913_v2, !((_this).F3()), _dafny.IntOfInt64(-214))
						var _953_v36 _dafny.Sequence
						_ = _953_v36
						_953_v36 = _dafny.SeqOf(_913_v2, _913_v2, _913_v2, (_952_v35).Dtor_cf7(), _913_v2)
						var _954_v37 _dafny.Sequence
						_ = _954_v37
						_954_v37 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)))
						var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_950_v33), 0))
						_ = _index194
						(_950_v33).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_953_v36, _953_v36), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_954_v37).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_953_v36, _953_v36)).Cardinality()))).Uint32(), _913_v2)).Cardinality()), (_index194).Int())
						var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_950_v33), 0))
						_ = _index195
						(_950_v33).ArraySet1(_910_v0, (_index195).Int())
						var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_950_v33), 0))
						_ = _index196
						var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
						_ = _index197
						var _rhs139 _dafny.Int = Companion_Default___.SafeModuloInt(_947_v30.F7, (_910_v0).Minus(_947_v30.F7))
						_ = _rhs139
						var _rhs140 bool = true
						_ = _rhs140
						var _rhs141 _dafny.Sequence = _953_v36
						_ = _rhs141
						var _rhs142 _dafny.Int = _910_v0
						_ = _rhs142
						var _lhs95 _dafny.Array = _950_v33
						_ = _lhs95
						var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_950_v33), 0))
						_ = _lhs96
						var _lhs97 _dafny.Array = _911_v1
						_ = _lhs97
						var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
						_ = _lhs98
						var _lhs99 *C4 = _947_v30
						_ = _lhs99
						(_lhs95).ArraySet1(_rhs139, (_lhs96).Int())
						(_lhs97).ArraySet1(_rhs140, (_lhs98).Int())
						_953_v36 = _rhs141
						_lhs99.F7 = _rhs142
					} else {
						_913_v2 = _913_v2
						var _955_v38 _dafny.Map
						_ = _955_v38
						_955_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_910_v0), (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool))
						var _956_v39 _dafny.Sequence
						_ = _956_v39
						_956_v39 = _dafny.UnicodeSeqOfUtf8Bytes("ftwmf")
						var _pat_let_tv14 = _911_v1
						_ = _pat_let_tv14
						var _pat_let_tv15 = _911_v1
						_ = _pat_let_tv15
						var _pat_let_tv16 = _955_v38
						_ = _pat_let_tv16
						var _pat_let_tv17 = globalState
						_ = _pat_let_tv17
						var _957_v40 _dafny.Map
						_ = _957_v40
						_957_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D3 {
							if (_this).F3() {
								return func(_pat_let15_0 D3) D3 {
									return func(_958_dt__update__tmp_h0 D3) D3 {
										return func(_pat_let16_0 _dafny.Int) D3 {
											return func(_959_dt__update_hcf14_h0 _dafny.Int) D3 {
												return Companion_D3_.Create_DC8_(_959_dt__update_hcf14_h0)
											}(_pat_let16_0)
										}(Companion_Default___.Fm19((_pat_let_tv15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_pat_let_tv14), 0))).Int()).(bool), (_pat_let_tv16).Cardinality(), _pat_let_tv17))
									}(_pat_let15_0)
								}(Companion_D3_.Create_DC8_(_910_v0))
							}
							return Companion_D3_.Create_DC8_((_dafny.MultiSetOf((_this).F3())).Cardinality())
						})(), Companion_Default___.Fm33(_956_v39, globalState))
						var _960_v42 _dafny.Map
						_ = _960_v42
						_960_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F3(), false)
						var _961_v43 _dafny.Map
						_ = _961_v43
						_961_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), (_960_v42).Cardinality())
						var _962_v44 _dafny.Sequence
						_ = _962_v44
						_962_v44 = _dafny.SeqOf((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool))
						var _963_v45 _dafny.Sequence
						_ = _963_v45
						_963_v45 = _dafny.SeqOf(_dafny.IntOfUint32((_962_v44).Cardinality()), _910_v0)
						var _964_v46 _dafny.Set
						_ = _964_v46
						_964_v46 = _dafny.SetOf(_961_v43, _961_v43, (_961_v43).Update((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), _dafny.IntOfUint32((_963_v45).Cardinality())), _961_v43)
						var _965_v47 D1
						_ = _965_v47
						_965_v47 = Companion_D1_.Create_DC3_((_964_v46).Cardinality())
						var _966_v48 D3
						_ = _966_v48
						_966_v48 = Companion_D3_.Create_DC8_((_965_v47).Dtor_cf5())
						var _967_v49 _dafny.Sequence
						_ = _967_v49
						_967_v49 = _dafny.SeqOf(_966_v48)
						_957_v40 = func() _dafny.Map {
							var _coll29 = _dafny.NewMapBuilder()
							_ = _coll29
							for _iter33 := _dafny.Iterate((_967_v49).Elements()); ; {
								_compr_29, _ok33 := _iter33()
								if !_ok33 {
									break
								}
								var _968_v41 D3
								_968_v41 = interface{}(_compr_29).(D3)
								if _dafny.Companion_Sequence_.Contains(_967_v49, _968_v41) {
									_coll29.Add(_968_v41, (_this).F3())
								}
							}
							return _coll29.ToMap()
						}()
						var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
						_ = _index198
						(_911_v1).ArraySet1((((_916_v5).Update(_910_v0, Companion_Default___.Abs(_dafny.IntOfInt64(469)))).Difference(_916_v5)).IsProperSubsetOf(_916_v5), (_index198).Int())
						var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
						_ = _index199
						(_911_v1).ArraySet1((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), (_index199).Int())
						var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
						_ = _index200
						(_911_v1).ArraySet1(((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)) && ((_this).F3()), (_index200).Int())
					}
					var _969_v50 _dafny.Map
					_ = _969_v50
					_969_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool)) || ((_this).F3()), _911_v1)
					var _970_v51 _dafny.Sequence
					_ = _970_v51
					_970_v51 = _dafny.SeqOf((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), true)
					var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
					_ = _index201
					var _rhs143 bool = (func() bool {
						if (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool) {
							return !_dafny.Companion_Sequence_.Equal(_970_v51, _970_v51)
						}
						return !(!((_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
							var _coll30 = _dafny.NewMapBuilder()
							_ = _coll30
							for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(902), _dafny.IntOfInt64(637))); ; {
								_compr_30, _ok34 := _iter34()
								if !_ok34 {
									break
								}
								var _971_v52 _dafny.Int
								_971_v52 = interface{}(_compr_30).(_dafny.Int)
								if ((_dafny.IntOfInt64(902)).Cmp(_971_v52) <= 0) && ((_971_v52).Cmp(_dafny.IntOfInt64(637)) < 0) {
									_coll30.Add(Companion_Default___.SafeDivisionInt(_971_v52, _dafny.IntOfInt64(470)), false)
								}
							}
							return _coll30.ToMap()
						}()).Cardinality(), _910_v0, _dafny.IntOfInt64(119))).Cardinality())).Cmp(_dafny.IntOfInt64(314)) > 0))
					})()
					_ = _rhs143
					var _rhs144 _dafny.Map = (_969_v50).Update((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), _911_v1)
					_ = _rhs144
					var _rhs145 bool = !((_this).F3()) || ((_this).F3())
					_ = _rhs145
					var _rhs146 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality())
					_ = _rhs146
					var _lhs100 _dafny.Array = _911_v1
					_ = _lhs100
					var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
					_ = _lhs101
					r0 = _rhs143
					_969_v50 = _rhs144
					(_lhs100).ArraySet1(_rhs145, (_lhs101).Int())
					_910_v0 = _rhs146
					var _972_v53 _dafny.Sequence
					_ = _972_v53
					_972_v53 = _dafny.SeqOf(_910_v0)
					var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
					_ = _index202
					var _rhs147 bool = (_this).F3()
					_ = _rhs147
					var _rhs148 _dafny.Sequence = _972_v53
					_ = _rhs148
					var _lhs102 _dafny.Array = _911_v1
					_ = _lhs102
					var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
					_ = _lhs103
					(_lhs102).ArraySet1(_rhs147, (_lhs103).Int())
					_972_v53 = _rhs148
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _974_v54 _dafny.Map
		_ = _974_v54
		_974_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v0, (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool))
		var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))
		_ = _index203
		(_911_v1).ArraySet1((func() bool {
			if (_974_v54).Contains(_910_v0) {
				return (_974_v54).Get(_910_v0).(bool)
			}
			return (_this).F3()
		})(), (_index203).Int())
		r0 = (_this).F3()
		var _975_v55 _dafny.MultiSet
		_ = _975_v55
		_975_v55 = _dafny.MultiSetOf((_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool), (_911_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_911_v1), 0))).Int()).(bool))
		var _976_v56 _dafny.Map
		_ = _976_v56
		_976_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_974_v54).Cardinality(), _975_v55)
		r1 = (_976_v56).Merge(_976_v56)
		return r0, r1
	}
}
func (_this *C7) F3() bool {
	{
		return _this._f3
	}
}
func (_this *C7) F2() _dafny.Array {
	{
		return _this._f2
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	dummy byte
}

func New_C8_() *C8 {
	_this := C8{}

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

func (_this *C8) Ctor__() {
	{
	}
}
func (_this *C8) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf(false, !(!(!(false))), !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(78), true)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality())))).Cardinality(), false)))
	}
}
func (_this *C8) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality()))).Intersection((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetOf(_dafny.IntOfInt64(276))).Cardinality())).Cardinality(), _dafny.IntOfInt64(-163))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(136))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg52 _dafny.Int) interface{} {
				return coer52(arg52)
			}
		}(func(_977_i0 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-353), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nyftlaqe")).Cardinality()))).Cardinality()
		})))))
	}
}
func (_this *C8) Fm2(p0 _dafny.Int, p1 bool, p2 _dafny.Set, globalState *GlobalState) bool {
	{
		return !(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("dyqqsd"), _dafny.UnicodeSeqOfUtf8Bytes("mtkslws")))
	}
}
func (_this *C8) M0(p0 _dafny.Array, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _978_v0 _dafny.Int
		_ = _978_v0
		_978_v0 = _dafny.IntOfInt64(-859)
		var _hi7 _dafny.Int = _978_v0
		_ = _hi7
		for _979_i0 := _978_v0; _979_i0.Cmp(_hi7) < 0; _979_i0 = _979_i0.Plus(_dafny.One) {
			var _980_v1 bool
			_ = _980_v1
			_980_v1 = true
			var _981_v2 _dafny.Set
			_ = _981_v2
			_981_v2 = _dafny.SetOf(_980_v1, _980_v1)
			var _982_v3 _dafny.Set
			_ = _982_v3
			_982_v3 = _dafny.SetOf(_981_v2)
			_982_v3 = Companion_Default___.Fm3(globalState)
			var _983_v4 _dafny.Sequence
			_ = _983_v4
			_983_v4 = _dafny.UnicodeSeqOfUtf8Bytes("hwkemvrcr")
			var _984_v5 _dafny.Sequence
			_ = _984_v5
			_984_v5 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(269))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}(func(_985_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			})), _983_v4)
			_978_v0 = Companion_Default___.SafeModuloInt(_978_v0, _dafny.IntOfUint32(((_984_v5).Select((Companion_Default___.SafeIndex(_978_v0, _dafny.IntOfUint32((_984_v5).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
			var _986_v6 _dafny.Array
			_ = _986_v6
			var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw175
			_986_v6 = _nw175
			var _987_v7 D0
			_ = _987_v7
			_987_v7 = Companion_D0_.Create_DC0_(_978_v0, _983_v4, _986_v6)
			var _988_v8 D0
			_ = _988_v8
			_988_v8 = Companion_D0_.Create_DC1_(Companion_D0_.Create_DC1_(_987_v7))
			var _989_v9 D0
			_ = _989_v9
			_989_v9 = Companion_D0_.Create_DC1_(Companion_D0_.Create_DC1_(_988_v8))
			var _source10 D0 = _989_v9
			_ = _source10
			if _source10.Is_DC0() {
				var _990___mcc_h0 _dafny.Int = _source10.Get_().(D0_DC0).Cf0
				_ = _990___mcc_h0
				var _991___mcc_h1 _dafny.Sequence = _source10.Get_().(D0_DC0).Cf1
				_ = _991___mcc_h1
				var _992___mcc_h2 _dafny.Array = _source10.Get_().(D0_DC0).Cf2
				_ = _992___mcc_h2
				var _993_cf2 _dafny.Array = _992___mcc_h2
				_ = _993_cf2
				var _994_cf1 _dafny.Sequence = _991___mcc_h1
				_ = _994_cf1
				var _995_cf0 _dafny.Int = _990___mcc_h0
				_ = _995_cf0
				var _996_v10 _dafny.Map
				_ = _996_v10
				_996_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_980_v1, _983_v4)
				var _997_v11 _dafny.CodePoint
				_ = _997_v11
				_997_v11 = _dafny.CodePoint('s')
				var _998_v12 D1
				_ = _998_v12
				_998_v12 = Companion_D1_.Create_DC4_(_978_v0, _997_v11, _980_v1, _978_v0)
				var _999_v13 _dafny.Array
				_ = _999_v13
				var _nwElement0_40 _dafny.Sequence = (func() _dafny.Sequence {
					if (_996_v10).Contains(!(_980_v1)) {
						return (_996_v10).Get(!(_980_v1)).(_dafny.Sequence)
					}
					return _983_v4
				})()
				_ = _nwElement0_40
				var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(23))
				_ = _nw176
				(_nw176).ArraySet1(_nwElement0_40, 0)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Update(_983_v4, (Companion_Default___.SafeIndex(_978_v0, _dafny.IntOfUint32((_983_v4).Cardinality()))).Uint32(), _997_v11), 1)
				(_nw176).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("p"), 2)
				(_nw176).ArraySet1(_983_v4, 3)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1000_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1001_i2 _dafny.Int) _dafny.CodePoint {
						return _1000_v11
					}
				})(_997_v11))), _994_cf1), 4)
				(_nw176).ArraySet1(_983_v4, 5)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_994_cf1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_1002_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1003_i3 _dafny.Int) _dafny.CodePoint {
						return _1002_v11
					}
				})(_997_v11)))), (Companion_Default___.SafeIndex(_978_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_994_cf1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_1004_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1005_i3 _dafny.Int) _dafny.CodePoint {
						return _1004_v11
					}
				})(_997_v11))))).Cardinality()))).Uint32(), _dafny.CodePoint('c')), 6)
				(_nw176).ArraySet1(_983_v4, 7)
				(_nw176).ArraySet1(_994_cf1, 8)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pmh"), _dafny.UnicodeSeqOfUtf8Bytes("moycxky")), 9)
				(_nw176).ArraySet1(_994_cf1, 10)
				(_nw176).ArraySet1(_983_v4, 11)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(458))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_1006_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1007_i4 _dafny.Int) _dafny.CodePoint {
						return _1006_v11
					}
				})(_997_v11))), _983_v4), 12)
				(_nw176).ArraySet1(_983_v4, 13)
				(_nw176).ArraySet1(_994_cf1, 14)
				(_nw176).ArraySet1(_994_cf1, 15)
				(_nw176).ArraySet1(_994_cf1, 16)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Update(_994_cf1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(53))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}((func(_1008_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1009_i5 _dafny.Int) _dafny.Int {
						return _1008_v0
					}
				})(_978_v0)))).Cardinality()), _dafny.IntOfUint32((_994_cf1).Cardinality()))).Uint32(), _997_v11), 17)
				(_nw176).ArraySet1(_994_cf1, 18)
				(_nw176).ArraySet1(_994_cf1, 19)
				(_nw176).ArraySet1(_983_v4, 20)
				(_nw176).ArraySet1((func() _dafny.Sequence {
					if (_998_v12).Dtor_cf8() {
						return _994_cf1
					}
					return _994_cf1
				})(), 21)
				(_nw176).ArraySet1(Companion_Default___.Fm4(_980_v1, _980_v1, _980_v1, globalState), 22)
				_999_v13 = _nw176
				_999_v13 = _999_v13
				var _1010_v14 _dafny.MultiSet
				_ = _1010_v14
				_1010_v14 = _dafny.MultiSetOf(_978_v0, _995_cf0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(578))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}((func(_1011_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1012_i6 _dafny.Int) _dafny.CodePoint {
						return _1011_v11
					}
				})(_997_v11)))).Cardinality()))
				var _1013_v15 _dafny.Map
				_ = _1013_v15
				_1013_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_995_cf0, _1010_v14)
				_1013_v15 = (_1013_v15).Update(_995_cf0, _dafny.MultiSetOf(_978_v0))
				r0 = (_980_v1) == (_980_v1)
				var _1014_v16 _dafny.Sequence
				_ = _1014_v16
				_1014_v16 = _dafny.SeqOf(_980_v1)
				var _1015_v17 _dafny.Map
				_ = _1015_v17
				_1015_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_980_v1, _980_v1, _980_v1), _1014_v16), _1014_v16)
				_1015_v17 = (_1015_v17).Update(_1014_v16, _1014_v16)
			} else {
				var _1016___mcc_h3 D0 = _source10.Get_().(D0_DC1).Cf3
				_ = _1016___mcc_h3
				var _1017_cf3 D0 = _1016___mcc_h3
				_ = _1017_cf3
				var _1018_v18 _dafny.Array
				_ = _1018_v18
				var _nw177 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(15))
				_ = _nw177
				_1018_v18 = _nw177
				var _1019_v19 D1
				_ = _1019_v19
				_1019_v19 = Companion_D1_.Create_DC3_(_978_v0)
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1018_v18), 0))
				_ = _index204
				(_1018_v18).ArraySet1(_1019_v19, (_index204).Int())
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1018_v18), 0))
				_ = _index205
				(_1018_v18).ArraySet1(_1019_v19, (_index205).Int())
				var _1020_v20 D1
				_ = _1020_v20
				_1020_v20 = Companion_D1_.Create_DC2_(_980_v1)
				_1020_v20 = _1020_v20
				var _1021_v22 _dafny.Map
				_ = _1021_v22
				_1021_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter36 := _dafny.Iterate(((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("abvray"))).Update(_983_v4, Companion_Default___.Abs(_dafny.IntOfUint32((_983_v4).Cardinality())))).Elements()); ; {
						_compr_32, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _1022_v21 _dafny.Sequence
						_1022_v21 = interface{}(_compr_32).(_dafny.Sequence)
						if ((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("abvray"))).Update(_983_v4, Companion_Default___.Abs(_dafny.IntOfUint32((_983_v4).Cardinality())))).Contains(_1022_v21) {
							_coll32.Add(_1022_v21, (Companion_D2_.Create_DC5_(_dafny.SeqOf(_978_v0))).Dtor_cf10())
						}
					}
					return _coll32.ToMap()
				}(), _980_v1)
				_1021_v22 = (_1021_v22).Update(Companion_Default___.Fm5(globalState), !((_980_v1) || (_980_v1)))
				var _1023_v23 _dafny.MultiSet
				_ = _1023_v23
				_1023_v23 = _dafny.MultiSetOf(_980_v1)
				var _1024_v24 *C1
				_ = _1024_v24
				var _nw178 *C1 = New_C1_()
				_ = _nw178
				_nw178.Ctor__((_1023_v23).Update(_980_v1, Companion_Default___.Abs(_978_v0)), _978_v0)
				_1024_v24 = _nw178
			}
			var _1025_v25 T0
			_ = _1025_v25
			var _nw179 *C2 = New_C2_()
			_ = _nw179
			_nw179.Ctor__(_983_v4, _983_v4)
			_1025_v25 = _nw179
			var _1026_v26 _dafny.Map
			_ = _1026_v26
			_1026_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_v25, _dafny.IntOfInt64(219))
			var _1027_v27 _dafny.Map
			_ = _1027_v27
			_1027_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_980_v1, _1026_v26)
			_1026_v26 = (func() _dafny.Map {
				if _980_v1 {
					return _1026_v26
				}
				return (func() _dafny.Map {
					if (_1027_v27).Contains(_980_v1) {
						return (_1027_v27).Get(_980_v1).(_dafny.Map)
					}
					return _1026_v26
				})()
			})()
		}
		var _hi8 _dafny.Int = (_978_v0).Times(_978_v0)
		_ = _hi8
		for _1028_i7 := (_dafny.Zero).Minus(_978_v0); _1028_i7.Cmp(_hi8) < 0; _1028_i7 = _1028_i7.Plus(_dafny.One) {
			r0 = (_1028_i7).Cmp(_978_v0) < 0
			var _1029_v28 _dafny.Sequence
			_ = _1029_v28
			_1029_v28 = _dafny.UnicodeSeqOfUtf8Bytes("erb")
			var _1030_v29 _dafny.Array
			_ = _1030_v29
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_16
			var _nw180 _dafny.Array
			_ = _nw180
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw180 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Sequence = (func(_1031_v0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1032_i8 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_1031_v0, _1031_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Cardinality())
					}
				})(_978_v0)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw180 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw180).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw180).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_1030_v29 = _nw180
			var _1033_v30 *C0
			_ = _1033_v30
			var _nw181 *C0 = New_C0_()
			_ = _nw181
			_nw181.Ctor__()
			_1033_v30 = _nw181
			var _1034_v31 bool
			_ = _1034_v31
			_1034_v31 = false
			var _1035_v32 _dafny.Map
			_ = _1035_v32
			_1035_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1034_v31, _1028_i7)
			var _1036_v33 _dafny.Map
			_ = _1036_v33
			_1036_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1034_v31, _1034_v31)
			var _rhs149 _dafny.Sequence = _1029_v28
			_ = _rhs149
			var _rhs150 _dafny.Array = _1030_v29
			_ = _rhs150
			var _rhs151 _dafny.Int = (func() _dafny.Int {
				if (_1035_v32).Contains(!(_1034_v31) || (_1034_v31)) {
					return (_1035_v32).Get(!(_1034_v31) || (_1034_v31)).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if _1034_v31 {
						return _1028_i7
					}
					return _978_v0
				})()
			})()
			_ = _rhs151
			var _rhs152 *C0 = _1033_v30
			_ = _rhs152
			var _rhs153 _dafny.Int = ((_dafny.Zero).Minus(((_1036_v33).Merge(_1036_v33)).Cardinality())).Times(_dafny.IntOfInt64(380))
			_ = _rhs153
			_1029_v28 = _rhs149
			_1030_v29 = _rhs150
			_978_v0 = _rhs151
			_1033_v30 = _rhs152
			_978_v0 = _rhs153
			var _1037_v34 D9
			_ = _1037_v34
			_1037_v34 = Companion_D9_.Create_DC20_(_1034_v31, _1034_v31, _1034_v31)
			var _pat_let_tv18 = _1034_v31
			_ = _pat_let_tv18
			var _1038_v35 _dafny.Sequence
			_ = _1038_v35
			_1038_v35 = _dafny.SeqOf(func(_pat_let17_0 D9) D9 {
				return func(_1039_dt__update__tmp_h0 D9) D9 {
					return func(_pat_let18_0 bool) D9 {
						return func(_1040_dt__update_hcf31_h0 bool) D9 {
							return Companion_D9_.Create_DC20_(_1040_dt__update_hcf31_h0, (_1039_dt__update__tmp_h0).Dtor_cf32(), (_1039_dt__update__tmp_h0).Dtor_cf33())
						}(_pat_let18_0)
					}(_pat_let_tv18)
				}(_pat_let17_0)
			}(Companion_D9_.Create_DC20_(_1034_v31, _1034_v31, _1034_v31)), Companion_D9_.Create_DC20_(_1034_v31, true, _1034_v31), _1037_v34, _1037_v34, _1037_v34)
			var _1041_v36 D14
			_ = _1041_v36
			_1041_v36 = Companion_D14_.Create_DC29_(_1038_v35)
			_1038_v35 = (_1041_v36).Dtor_cf46()
			_1034_v31 = _1034_v31
		}
		var _1042_v37 bool
		_ = _1042_v37
		_1042_v37 = false
		var _1043_v38 _dafny.Sequence
		_ = _1043_v38
		_1043_v38 = _dafny.SeqOf(_1042_v37)
		var _1044_i9 _dafny.Int
		_ = _1044_i9
		_1044_i9 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1043_v38, _1043_v38), _1043_v38) {
				{
					if (_1044_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1044_i9 = (_1044_i9).Plus(_dafny.One)
					_1042_v37 = _1042_v37
					var _1045_v39 _dafny.Map
					_ = _1045_v39
					_1045_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_978_v0, _978_v0), (_978_v0).Cmp(_978_v0) <= 0)
					var _1046_v40 _dafny.Map
					_ = _1046_v40
					_1046_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_978_v0, _978_v0)
					_1045_v39 = (_1045_v39).Update(_1046_v40, _1042_v37)
					var _1047_v41 _dafny.Array
					_ = _1047_v41
					var _nwElement0_41 bool = _1042_v37
					_ = _nwElement0_41
					var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(16))
					_ = _nw182
					(_nw182).ArraySet1(_nwElement0_41, 0)
					(_nw182).ArraySet1(_1042_v37, 1)
					(_nw182).ArraySet1(_1042_v37, 2)
					(_nw182).ArraySet1(_1042_v37, 3)
					(_nw182).ArraySet1(_1042_v37, 4)
					(_nw182).ArraySet1(_1042_v37, 5)
					(_nw182).ArraySet1(_1042_v37, 6)
					(_nw182).ArraySet1(false, 7)
					(_nw182).ArraySet1(_1042_v37, 8)
					(_nw182).ArraySet1(_1042_v37, 9)
					(_nw182).ArraySet1(false, 10)
					(_nw182).ArraySet1(_1042_v37, 11)
					(_nw182).ArraySet1(_1042_v37, 12)
					(_nw182).ArraySet1(_1042_v37, 13)
					(_nw182).ArraySet1(_1042_v37, 14)
					(_nw182).ArraySet1(_1042_v37, 15)
					_1047_v41 = _nw182
					var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((p0), 0))
					_ = _index206
					(p0).ArraySet1(_1047_v41, (_index206).Int())
					var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((p0), 0))
					_ = _index207
					(p0).ArraySet1(_1047_v41, (_index207).Int())
					var _1048_v42 _dafny.Map
					_ = _1048_v42
					_1048_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1047_v41, !(!(!(_1042_v37) || (_1042_v37))))
					_1042_v37 = (func() bool {
						if (_1048_v42).Contains(_1047_v41) {
							return (_1048_v42).Get(_1047_v41).(bool)
						}
						return _1042_v37
					})()
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _1049_v43 _dafny.Array
		_ = _1049_v43
		var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
		_ = _nw183
		_1049_v43 = _nw183
		var _1050_v44 _dafny.Sequence
		_ = _1050_v44
		_1050_v44 = _dafny.UnicodeSeqOfUtf8Bytes("nhh")
		var _1051_v45 _dafny.Map
		_ = _1051_v45
		_1051_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1042_v37, _1050_v44)
		var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_1049_v43), 0))
		_ = _index208
		(_1049_v43).ArraySet1((func() _dafny.Sequence {
			if (_1051_v45).Contains(_1042_v37) {
				return (_1051_v45).Get(_1042_v37).(_dafny.Sequence)
			}
			return _1050_v44
		})(), (_index208).Int())
		var _1052_v47 _dafny.Map
		_ = _1052_v47
		_1052_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(445), _dafny.IntOfInt64(767))); ; {
				_compr_33, _ok37 := _iter37()
				if !_ok37 {
					break
				}
				var _1053_v46 _dafny.Int
				_1053_v46 = interface{}(_compr_33).(_dafny.Int)
				if ((_dafny.IntOfInt64(445)).Cmp(_1053_v46) <= 0) && ((_1053_v46).Cmp(_dafny.IntOfInt64(767)) < 0) {
					_coll33.Add((_1053_v46).Plus(_978_v0), _1042_v37)
				}
			}
			return _coll33.ToMap()
		}()).Cardinality(), _1042_v37)
		var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_1049_v43), 0))
		_ = _index209
		(_1049_v43).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1050_v44, Companion_Default___.Fm22(_1042_v37, (_1052_v47).Cardinality(), globalState)), (_index209).Int())
		_1042_v37 = _1042_v37
		_978_v0 = _978_v0
		r0 = _1042_v37
		var _1054_v48 _dafny.Array
		_ = _1054_v48
		var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw184
		_1054_v48 = _nw184
		var _1055_v49 _dafny.Sequence
		_ = _1055_v49
		_1055_v49 = _dafny.SeqOf(_1054_v48)
		var _1056_v50 _dafny.Sequence
		_ = _1056_v50
		_1056_v50 = _dafny.SeqOf((_1049_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_1049_v43), 0))).Int()).(_dafny.Sequence))
		var _1057_v51 _dafny.Sequence
		_ = _1057_v51
		_1057_v51 = _dafny.SeqOf((_1055_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1056_v50).Select((Companion_Default___.SafeIndex(_978_v0, _dafny.IntOfUint32((_1056_v50).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_1055_v49).Cardinality()))).Uint32()).(_dafny.Array))
		var _1058_v52 _dafny.MultiSet
		_ = _1058_v52
		_1058_v52 = _dafny.MultiSetOf(true)
		r1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1057_v51).Cardinality()), _1058_v52)
		return r0, r1
	}
}
func (_this *C8) M1(p0 bool, globalState *GlobalState) (bool, _dafny.Array, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		var _1059_v0 _dafny.Sequence
		_ = _1059_v0
		_1059_v0 = _dafny.UnicodeSeqOfUtf8Bytes("goxqvwwrh")
		_1059_v0 = Companion_Default___.Fm4(!(p0) || (p0), p0, p0, globalState)
		var _1060_v1 _dafny.Int
		_ = _1060_v1
		_1060_v1 = _dafny.IntOfInt64(669)
		var _1061_v2 _dafny.Sequence
		_ = _1061_v2
		_1061_v2 = _dafny.SeqOf(p0)
		_1060_v1 = _dafny.IntOfUint32((_1061_v2).Cardinality())
		var _1062_v3 _dafny.MultiSet
		_ = _1062_v3
		_1062_v3 = _dafny.MultiSetOf(_1060_v1, _dafny.IntOfInt64(-45), _1060_v1)
		_1062_v3 = _1062_v3
		var _1063_v4 D1
		_ = _1063_v4
		_1063_v4 = Companion_D1_.Create_DC3_(_1060_v1)
		var _1064_v5 D9
		_ = _1064_v5
		_1064_v5 = Companion_D9_.Create_DC19_(_1060_v1, _1060_v1)
		var _pat_let_tv19 = _1064_v5
		_ = _pat_let_tv19
		var _source11 D1 = func(_pat_let19_0 D1) D1 {
			return func(_1065_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let20_0 _dafny.Int) D1 {
					return func(_1066_dt__update_hcf5_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC3_(_1066_dt__update_hcf5_h0)
					}(_pat_let20_0)
				}((_pat_let_tv19).Dtor_cf29())
			}(_pat_let19_0)
		}(_1063_v4)
		_ = _source11
		if _source11.Is_DC3() {
			var _1067___mcc_h0 _dafny.Int = _source11.Get_().(D1_DC3).Cf5
			_ = _1067___mcc_h0
			var _1068_cf5 _dafny.Int = _1067___mcc_h0
			_ = _1068_cf5
			var _1069_v6 _dafny.Sequence
			_ = _1069_v6
			_1069_v6 = _dafny.SeqOf(_1060_v1, _1060_v1, _1068_cf5, _1068_cf5)
			var _1070_v7 _dafny.Map
			_ = _1070_v7
			_1070_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1060_v1, _1069_v6)
			var _1071_v8 _dafny.Map
			_ = _1071_v8
			_1071_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1060_v1, p0)
			var _1072_v9 D10
			_ = _1072_v9
			_1072_v9 = Companion_D10_.Create_DC21_(_dafny.MultiSetOf((func() bool {
				if (_1071_v8).Contains(_1060_v1) {
					return (_1071_v8).Get(_1060_v1).(bool)
				}
				return false
			})(), p0, p0))
			var _1073_v10 _dafny.Map
			_ = _1073_v10
			_1073_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1068_cf5)
			var _1074_v11 _dafny.CodePoint
			_ = _1074_v11
			_1074_v11 = _dafny.CodePoint('w')
			var _1075_v12 _dafny.Array
			_ = _1075_v12
			var _nwElement0_42 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1070_v7).Contains(_dafny.IntOfInt64(696)) {
					return (_1070_v7).Get(_dafny.IntOfInt64(696)).(_dafny.Sequence)
				}
				return _1069_v6
			})()).Cardinality()))
			_ = _nwElement0_42
			var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(12))
			_ = _nw185
			(_nw185).ArraySet1(_nwElement0_42, 0)
			(_nw185).ArraySet1(_1060_v1, 1)
			(_nw185).ArraySet1(((_1072_v9).Dtor_cf34()).Cardinality(), 2)
			(_nw185).ArraySet1(_1068_cf5, 3)
			(_nw185).ArraySet1((_1068_cf5).Minus(_1060_v1), 4)
			(_nw185).ArraySet1(_dafny.IntOfUint32((_1069_v6).Cardinality()), 5)
			(_nw185).ArraySet1((func() _dafny.Int {
				if (_1073_v10).Contains(p0) {
					return (_1073_v10).Get(p0).(_dafny.Int)
				}
				return _1068_cf5
			})(), 6)
			(_nw185).ArraySet1(_1060_v1, 7)
			(_nw185).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm11(_1074_v11, _1068_cf5, true, Companion_Default___.Fm34(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qbdinbg")).Cardinality()), _1068_cf5, globalState), globalState)).Cardinality())), 8)
			(_nw185).ArraySet1(Companion_Default___.Fm24(_dafny.IntOfInt64(857), true, p0, _1060_v1, globalState), 9)
			(_nw185).ArraySet1(_dafny.IntOfUint32((_1069_v6).Cardinality()), 10)
			(_nw185).ArraySet1(_dafny.IntOfInt64(321), 11)
			_1075_v12 = _nw185
			_1075_v12 = _1075_v12
			var _1076_v13 *C1
			_ = _1076_v13
			var _nw186 *C1 = New_C1_()
			_ = _nw186
			_nw186.Ctor__(_dafny.MultiSetFromSeq(_1061_v2), (_1071_v8).Cardinality())
			_1076_v13 = _nw186
			var _1077_v14 _dafny.Map
			_ = _1077_v14
			_1077_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1076_v13)
			var _1078_v15 _dafny.Map
			_ = _1078_v15
			_1078_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _1079_v16 _dafny.Sequence
			_ = _1079_v16
			_1079_v16 = _dafny.SeqOf(_1075_v12, _1075_v12)
			var _1080_v17 _dafny.Array
			_ = _1080_v17
			var _nwElement0_43 _dafny.Array = _1075_v12
			_ = _nwElement0_43
			var _nw187 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(19))
			_ = _nw187
			(_nw187).ArraySet1(_nwElement0_43, 0)
			(_nw187).ArraySet1(_1075_v12, 1)
			(_nw187).ArraySet1(_1075_v12, 2)
			(_nw187).ArraySet1(_1075_v12, 3)
			(_nw187).ArraySet1(_1075_v12, 4)
			(_nw187).ArraySet1((_1079_v16).Select((Companion_Default___.SafeIndex(_1068_cf5, _dafny.IntOfUint32((_1079_v16).Cardinality()))).Uint32()).(_dafny.Array), 5)
			(_nw187).ArraySet1((_1079_v16).Select((Companion_Default___.SafeIndex(_1060_v1, _dafny.IntOfUint32((_1079_v16).Cardinality()))).Uint32()).(_dafny.Array), 6)
			(_nw187).ArraySet1(_1075_v12, 7)
			(_nw187).ArraySet1((_1079_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.IntOfUint32((_1079_v16).Cardinality()))).Uint32()).(_dafny.Array), 8)
			(_nw187).ArraySet1(_1075_v12, 9)
			(_nw187).ArraySet1(_1075_v12, 10)
			(_nw187).ArraySet1(_1075_v12, 11)
			(_nw187).ArraySet1(_1075_v12, 12)
			(_nw187).ArraySet1(_1075_v12, 13)
			(_nw187).ArraySet1(_1075_v12, 14)
			(_nw187).ArraySet1(_1075_v12, 15)
			(_nw187).ArraySet1(_1075_v12, 16)
			(_nw187).ArraySet1(_1075_v12, 17)
			(_nw187).ArraySet1(_1075_v12, 18)
			_1080_v17 = _nw187
			var _1081_v18 _dafny.Map
			_ = _1081_v18
			_1081_v18 = _1073_v10
			var _1082_v19 _dafny.Set
			_ = _1082_v19
			_1082_v19 = _dafny.SetOf(p0)
			var _1083_v20 D14
			_ = _1083_v20
			_1083_v20 = Companion_D14_.Create_DC30_(_1060_v1, _1080_v17, _1081_v18, (_1082_v19).Cardinality(), _1075_v12)
			var _1084_v21 _dafny.Array
			_ = _1084_v21
			var _nwElement0_44 bool = (_1068_cf5).Cmp(((_1077_v14).Update(!(p0), _1076_v13)).Cardinality()) > 0
			_ = _nwElement0_44
			var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(11))
			_ = _nw188
			(_nw188).ArraySet1(_nwElement0_44, 0)
			(_nw188).ArraySet1((func() bool {
				if p0 {
					return !((func() bool {
						if (_1078_v15).Contains((func() bool {
							if (_1078_v15).Contains(p0) {
								return (_1078_v15).Get(p0).(bool)
							}
							return p0
						})()) {
							return (_1078_v15).Get((func() bool {
								if (_1078_v15).Contains(p0) {
									return (_1078_v15).Get(p0).(bool)
								}
								return p0
							})()).(bool)
						}
						return p0
					})())
				}
				return p0
			})(), 1)
			(_nw188).ArraySet1((_1060_v1).Cmp(_1060_v1) < 0, 2)
			(_nw188).ArraySet1(p0, 3)
			(_nw188).ArraySet1(p0, 4)
			(_nw188).ArraySet1(false, 5)
			(_nw188).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1069_v6, (_1083_v20).Dtor_cf47()), 6)
			(_nw188).ArraySet1(false, 7)
			(_nw188).ArraySet1(p0, 8)
			(_nw188).ArraySet1(p0, 9)
			(_nw188).ArraySet1(false, 10)
			_1084_v21 = _nw188
			r1 = _1084_v21
			_1069_v6 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(481))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}(func(_1085_i0 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-426))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}(func(_1086_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))).Cardinality()))
			})), _dafny.Companion_Sequence_.Concatenate(_1069_v6, _1069_v6))
			var _1087_v22 *C1
			_ = _1087_v22
			var _nw189 *C1 = New_C1_()
			_ = _nw189
			_nw189.Ctor__((_1076_v13).F10(), _1076_v13.F11)
			_1087_v22 = _nw189
		} else if _source11.Is_DC4() {
			var _1088___mcc_h1 _dafny.Int = _source11.Get_().(D1_DC4).Cf6
			_ = _1088___mcc_h1
			var _1089___mcc_h2 _dafny.CodePoint = _source11.Get_().(D1_DC4).Cf7
			_ = _1089___mcc_h2
			var _1090___mcc_h3 bool = _source11.Get_().(D1_DC4).Cf8
			_ = _1090___mcc_h3
			var _1091___mcc_h4 _dafny.Int = _source11.Get_().(D1_DC4).Cf9
			_ = _1091___mcc_h4
			var _1092_cf9 _dafny.Int = _1091___mcc_h4
			_ = _1092_cf9
			var _1093_cf8 bool = _1090___mcc_h3
			_ = _1093_cf8
			var _1094_cf7 _dafny.CodePoint = _1089___mcc_h2
			_ = _1094_cf7
			var _1095_cf6 _dafny.Int = _1088___mcc_h1
			_ = _1095_cf6
			r0 = (func() bool {
				if _1093_cf8 {
					return p0
				}
				return (p0) == (_1093_cf8)
			})()
			if !((_1060_v1).Cmp(_1092_cf9) < 0) {
				var _1096_v23 _dafny.Array
				_ = _1096_v23
				var _nw190 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw190
				_1096_v23 = _nw190
				var _1097_v24 *C7
				_ = _1097_v24
				var _nw191 *C7 = New_C7_()
				_ = _nw191
				_nw191.Ctor__(_1096_v23, p0)
				_1097_v24 = _nw191
				var _1098_v25 _dafny.Map
				_ = _1098_v25
				_1098_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
				var _1099_v26 _dafny.Set
				_ = _1099_v26
				_1099_v26 = _dafny.SetOf((_dafny.Zero).Minus(_1060_v1), _1060_v1, _dafny.IntOfUint32((_dafny.SeqOf((_1098_v25).Cardinality(), _1060_v1)).Cardinality()), _1060_v1)
				var _1100_v27 _dafny.Array
				_ = _1100_v27
				var _nwElement0_45 _dafny.Int = (_dafny.SetOf(_1097_v24)).Cardinality()
				_ = _nwElement0_45
				var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(14))
				_ = _nw192
				(_nw192).ArraySet1(_nwElement0_45, 0)
				(_nw192).ArraySet1(_1060_v1, 1)
				(_nw192).ArraySet1(_1092_cf9, 2)
				(_nw192).ArraySet1(_1095_cf6, 3)
				(_nw192).ArraySet1((_1099_v26).Cardinality(), 4)
				(_nw192).ArraySet1((_dafny.Zero).Minus((_1062_v3).Cardinality()), 5)
				(_nw192).ArraySet1(_1060_v1, 6)
				(_nw192).ArraySet1(_1095_cf6, 7)
				(_nw192).ArraySet1(_dafny.IntOfUint32((_1061_v2).Cardinality()), 8)
				(_nw192).ArraySet1((_1097_v24).Fm6((_1097_v24).F3(), _1095_cf6, _1060_v1, _1094_cf7, globalState), 9)
				(_nw192).ArraySet1(_dafny.IntOfUint32((_1059_v0).Cardinality()), 10)
				(_nw192).ArraySet1(_dafny.IntOfInt64(210), 11)
				(_nw192).ArraySet1(_1092_cf9, 12)
				(_nw192).ArraySet1(((_1098_v25).Update((_1097_v24).F3(), _1093_cf8)).Cardinality(), 13)
				_1100_v27 = _nw192
				var _1101_v28 _dafny.Array
				_ = _1101_v28
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_17
				var _nw193 _dafny.Array
				_ = _nw193
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw193 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Map = (func(_1102_v25 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_1103_i2 _dafny.Int) _dafny.Map {
							return _1102_v25
						}
					})(_1098_v25)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw193 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw193).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw193).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_1101_v28 = _nw193
				var _1104_v29 _dafny.Map
				_ = _1104_v29
				_1104_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1100_v27, _1101_v28)
				_1104_v29 = (_1104_v29).Update(_1100_v27, _1101_v28)
				var _1105_v30 _dafny.Array
				_ = _1105_v30
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_18
				var _nw194 _dafny.Array
				_ = _nw194
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw194 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Sequence = (func(_1106_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1107_i3 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_1106_v0, _1106_v0)
						}
					})(_1059_v0)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw194 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw194).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw194).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_1105_v30 = _nw194
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_1105_v30), 0))
				_ = _index210
				(_1105_v30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1059_v0, _1059_v0), (_index210).Int())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_1105_v30), 0))
				_ = _index211
				(_1105_v30).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(111))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_1108_cf7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1109_i4 _dafny.Int) _dafny.CodePoint {
						return _1108_cf7
					}
				})(_1094_cf7))), Companion_Default___.Fm22((_1097_v24).F3(), _1095_cf6, globalState)), (Companion_Default___.SafeIndex(_1092_cf9, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(111))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_1110_cf7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1111_i4 _dafny.Int) _dafny.CodePoint {
						return _1110_cf7
					}
				})(_1094_cf7))), Companion_Default___.Fm22((_1097_v24).F3(), _1095_cf6, globalState))).Cardinality()))).Uint32(), _dafny.CodePoint('w')), (_index211).Int())
				var _1112_v31 *C5
				_ = _1112_v31
				var _nw195 *C5 = New_C5_()
				_ = _nw195
				_nw195.Ctor__((_1105_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_1105_v30), 0))).Int()).(_dafny.Sequence))
				_1112_v31 = _nw195
				_1059_v0 = _dafny.Companion_Sequence_.Concatenate((_1112_v31).F6(), (_1105_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_1105_v30), 0))).Int()).(_dafny.Sequence))
				var _1113_v32 _dafny.Map
				_ = _1113_v32
				_1113_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1092_cf9).Minus(_1060_v1), false)
				_1113_v32 = (_1113_v32).Update(_1092_cf9, _1093_cf8)
			} else {
				var _1114_v33 *C4
				_ = _1114_v33
				var _nw196 *C4 = New_C4_()
				_ = _nw196
				_nw196.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1059_v0).Cardinality()), _1060_v1))
				_1114_v33 = _nw196
				_1114_v33 = _1114_v33
				var _1115_v34 *C0
				_ = _1115_v34
				var _nw197 *C0 = New_C0_()
				_ = _nw197
				_nw197.Ctor__()
				_1115_v34 = _nw197
				var _1116_v35 _dafny.Array
				_ = _1116_v35
				var _nw198 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw198
				_1116_v35 = _nw198
				var _1117_v36 _dafny.Set
				_ = _1117_v36
				_1117_v36 = _dafny.SetOf(_1116_v35, _1116_v35)
				var _rhs154 bool = _1093_cf8
				_ = _rhs154
				var _rhs155 _dafny.Set = _1117_v36
				_ = _rhs155
				var _rhs156 _dafny.Int = _1114_v33.F7
				_ = _rhs156
				var _rhs157 bool = p0
				_ = _rhs157
				var _rhs158 bool = p0
				_ = _rhs158
				r0 = _rhs154
				_1117_v36 = _rhs155
				_1095_cf6 = _rhs156
				_1093_cf8 = _rhs157
				r0 = _rhs158
				var _1118_v37 *C2
				_ = _1118_v37
				var _nw199 *C2 = New_C2_()
				_ = _nw199
				_nw199.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("sm"), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1059_v0, (Companion_Default___.SafeIndex(_1114_v33.F7, _dafny.IntOfUint32((_1059_v0).Cardinality()))).Uint32(), _1094_cf7), (Companion_Default___.SafeIndex(_1060_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1059_v0, (Companion_Default___.SafeIndex(_1114_v33.F7, _dafny.IntOfUint32((_1059_v0).Cardinality()))).Uint32(), _1094_cf7)).Cardinality()))).Uint32(), _1094_cf7))
				_1118_v37 = _nw199
				var _1119_v38 _dafny.Map
				_ = _1119_v38
				_1119_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1060_v1, _1118_v37)
				_1119_v38 = (_1119_v38).Update(_1060_v1, _1118_v37)
				var _1120_v39 _dafny.Sequence
				_ = _1120_v39
				_1120_v39 = _dafny.SeqOf(_1116_v35)
				var _1121_v40 _dafny.MultiSet
				_ = _1121_v40
				_1121_v40 = _dafny.MultiSetOf(_1116_v35, (func() _dafny.Array {
					if !(_1093_cf8) {
						return _1116_v35
					}
					return (_1120_v39).Select((Companion_Default___.SafeIndex(_1060_v1, _dafny.IntOfUint32((_1120_v39).Cardinality()))).Uint32()).(_dafny.Array)
				})(), _1116_v35, _1116_v35, _1116_v35)
				_1092_cf9 = (func() _dafny.Int {
					if (_1121_v40).Contains(_1116_v35) {
						return (_1121_v40).Multiplicity(_1116_v35)
					}
					return Companion_Default___.SafeModuloInt(_1060_v1, _dafny.IntOfInt64(382))
				})()
			}
			var _1122_v41 _dafny.Array
			_ = _1122_v41
			var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw200
			_1122_v41 = _nw200
			_1122_v41 = _1122_v41
			var _1123_v42 _dafny.Array
			_ = _1123_v42
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_19
			var _nw201 _dafny.Array
			_ = _nw201
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw201 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Set = (func(_1124_cf9 _dafny.Int, _1125_cf8 bool, _1126_cf6 _dafny.Int) func(_dafny.Int) _dafny.Set {
					return func(_1127_i5 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_dafny.SeqOf(_1124_cf9, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_1125_cf8))).Cardinality(), _1126_cf6))).Difference(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(174))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg64 _dafny.Int) interface{} {
								return coer64(arg64)
							}
						}(func(_1128_i6 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
						}))))
					}
				})(_1092_cf9, _1093_cf8, _1095_cf6)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw201 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw201).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw201).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_1123_v42 = _nw201
			var _1129_v43 _dafny.Sequence
			_ = _1129_v43
			_1129_v43 = _dafny.SeqOf(_1092_cf9)
			var _1130_v44 _dafny.Set
			_ = _1130_v44
			_1130_v44 = _dafny.SetOf(_1129_v43)
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_1123_v42), 0))
			_ = _index212
			(_1123_v42).ArraySet1(_1130_v44, (_index212).Int())
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_1123_v42), 0))
			_ = _index213
			(_1123_v42).ArraySet1(func() _dafny.Set {
				var _coll34 = _dafny.NewBuilder()
				_ = _coll34
				for _iter38 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1129_v43, _1060_v1)).Update(_dafny.SeqOf(_1095_cf6, _1092_cf9), _dafny.IntOfInt64(437))).Keys().Elements()); ; {
					_compr_34, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _1131_v45 _dafny.Sequence
					_1131_v45 = interface{}(_compr_34).(_dafny.Sequence)
					if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1129_v43, _1060_v1)).Update(_dafny.SeqOf(_1095_cf6, _1092_cf9), _dafny.IntOfInt64(437))).Contains(_1131_v45) {
						_coll34.Add(_1131_v45)
					}
				}
				return _coll34.ToSet()
			}(), (_index213).Int())
		} else {
			var _1132___mcc_h5 bool = _source11.Get_().(D1_DC2).Cf4
			_ = _1132___mcc_h5
			var _1133_cf4 bool = _1132___mcc_h5
			_ = _1133_cf4
			var _1134_v46 _dafny.Array
			_ = _1134_v46
			var _nw202 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw202
			_1134_v46 = _nw202
			var _1135_v47 _dafny.Map
			_ = _1135_v47
			_1135_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1134_v46)
			var _1136_v48 D4
			_ = _1136_v48
			_1136_v48 = Companion_D4_.Create_DC9_(_1134_v46)
			var _1137_v49 D4
			_ = _1137_v49
			_1137_v49 = Companion_D4_.Create_DC11_(_1136_v48)
			var _1138_v50 _dafny.Map
			_ = _1138_v50
			_1138_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1135_v47).Cardinality(), _1137_v49)
			_1138_v50 = (_1138_v50).Update((_dafny.Zero).Minus(_1060_v1), _1137_v49)
			var _1139_v51 _dafny.Map
			_ = _1139_v51
			_1139_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1134_v46, false)
			var _1140_v52 D9
			_ = _1140_v52
			_1140_v52 = Companion_D9_.Create_DC20_(p0, p0, _1133_cf4)
			_1139_v51 = (_1139_v51).Update(_1134_v46, (_1140_v52).Dtor_cf33())
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1134_v46), 0))
			_ = _index214
			(_1134_v46).ArraySet1(_1060_v1, (_index214).Int())
			var _1141_v53 _dafny.CodePoint
			_ = _1141_v53
			_1141_v53 = _dafny.CodePoint('x')
			var _1142_v54 _dafny.Set
			_ = _1142_v54
			_1142_v54 = _dafny.SetOf(_1141_v53)
			var _1143_v55 _dafny.Sequence
			_ = _1143_v55
			_1143_v55 = _dafny.SeqOf(_1142_v54, _dafny.SetOf(_1141_v53), _1142_v54)
			var _1144_v56 _dafny.Sequence
			_ = _1144_v56
			_1144_v56 = _dafny.SeqOf((_dafny.Zero).Minus(_1060_v1), _1060_v1, (_1063_v4).Dtor_cf5(), _1060_v1, _dafny.IntOfUint32((_1143_v55).Cardinality()))
			var _1145_v57 _dafny.Array
			_ = _1145_v57
			var _nw203 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw203
			_1145_v57 = _nw203
			var _1146_v58 _dafny.Sequence
			_ = _1146_v58
			_1146_v58 = _dafny.SeqOf(_1145_v57)
			var _1147_v59 _dafny.Map
			_ = _1147_v59
			_1147_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1060_v1, _1060_v1)
			var _1148_v60 _dafny.MultiSet
			_ = _1148_v60
			_1148_v60 = _dafny.MultiSetOf(Companion_Default___.Fm8(_1062_v3, globalState), false)
			var _1149_v61 *C6
			_ = _1149_v61
			var _nw204 *C6 = New_C6_()
			_ = _nw204
			_nw204.Ctor__(_1148_v60, _1134_v46)
			_1149_v61 = _nw204
			var _1150_v62 _dafny.Map
			_ = _1150_v62
			_1150_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1149_v61, _1145_v57)
			var _1151_v63 _dafny.Set
			_ = _1151_v63
			_1151_v63 = _dafny.SetOf(_1145_v57, _1145_v57)
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1134_v46), 0))
			_ = _index215
			var _rhs159 _dafny.Int = _1060_v1
			_ = _rhs159
			var _rhs160 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_1060_v1, (_1060_v1).Plus(_1060_v1)))
			_ = _rhs160
			var _rhs161 _dafny.Sequence = (func() _dafny.Sequence {
				if _1133_cf4 {
					return _1144_v56
				}
				return _1144_v56
			})()
			_ = _rhs161
			var _rhs162 _dafny.Int = _1060_v1
			_ = _rhs162
			var _rhs163 bool = (_1151_v63).IsProperSubsetOf(_dafny.SetOf(_1145_v57, _1145_v57, (_1146_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1060_v1), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1147_v59)).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf(_1060_v1)).Cardinality()))).Uint32(), _1060_v1)).Cardinality()), _dafny.IntOfUint32((_1146_v58).Cardinality()))).Uint32()).(_dafny.Array), (func() _dafny.Array {
				if (_1150_v62).Contains(_1149_v61) {
					return (_1150_v62).Get(_1149_v61).(_dafny.Array)
				}
				return _1145_v57
			})()))
			_ = _rhs163
			var _lhs104 _dafny.Array = _1134_v46
			_ = _lhs104
			var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1134_v46), 0))
			_ = _lhs105
			_1060_v1 = _rhs159
			(_lhs104).ArraySet1(_rhs160, (_lhs105).Int())
			_1144_v56 = _rhs161
			_1060_v1 = _rhs162
			r0 = _rhs163
			var _1152_v64 _dafny.Map
			_ = _1152_v64
			_1152_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1134_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1134_v46), 0))).Int()).(_dafny.Int), _1061_v2)
			var _1153_v65 *C0
			_ = _1153_v65
			var _nw205 *C0 = New_C0_()
			_ = _nw205
			_nw205.Ctor__()
			_1153_v65 = _nw205
			var _1154_v66 D7
			_ = _1154_v66
			_1154_v66 = Companion_D7_.Create_DC15_(true, _dafny.IntOfInt64(-529), _1060_v1, (_1134_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1134_v46), 0))).Int()).(_dafny.Int), _1153_v65)
			var _1155_v67 _dafny.MultiSet
			_ = _1155_v67
			_1155_v67 = _dafny.MultiSetOf(_1154_v66, _1154_v66)
			var _1156_v68 _dafny.Map
			_ = _1156_v68
			_1156_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1155_v67, _1059_v0)
			var _1157_v69 _dafny.Map
			_ = _1157_v69
			_1157_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1152_v64).Contains((_1134_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1134_v46), 0))).Int()).(_dafny.Int)) {
					return (_1152_v64).Get((_1134_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1134_v46), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1133_cf4), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1144_v56).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_1133_cf4)).Cardinality()))).Uint32(), p0)
			})()).Cardinality()), (_1060_v1).Cmp((_1156_v68).Cardinality()) >= 0)
			var _1158_v70 T1
			_ = _1158_v70
			var _nw206 *C1 = New_C1_()
			_ = _nw206
			_nw206.Ctor__(_1148_v60, (_1134_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1134_v46), 0))).Int()).(_dafny.Int))
			_1158_v70 = _nw206
			var _1159_v71 _dafny.MultiSet
			_ = _1159_v71
			_1159_v71 = _dafny.MultiSetOf(_1158_v70)
			var _1160_v72 D2
			_ = _1160_v72
			_1160_v72 = Companion_D2_.Create_DC6_(_1060_v1, _1133_cf4)
			var _1161_v73 _dafny.Map
			_ = _1161_v73
			_1161_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1149_v61, false)
			_1157_v69 = (_1157_v69).Update(((_1159_v71).Cardinality()).Plus((_1062_v3).Cardinality()), (_1149_v61).Fm10(_1160_v72, _dafny.IntOfUint32((_1059_v0).Cardinality()), (_this).Fm1(_1060_v1, _dafny.IntOfUint32((_dafny.SeqOf((_1161_v73).Cardinality())).Cardinality()), _1144_v56, globalState), globalState))
		}
		var _1162_v74 _dafny.Sequence
		_ = _1162_v74
		_1162_v74 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1059_v0, (Companion_Default___.SafeIndex((_1062_v3).Cardinality(), _dafny.IntOfUint32((_1059_v0).Cardinality()))).Uint32(), Companion_Default___.Fm30(_1059_v0, _1060_v1, _1060_v1, globalState))).Cardinality()), _1060_v1)
		var _hi9 _dafny.Int = _1060_v1
		_ = _hi9
		for _1163_i7 := _dafny.IntOfUint32((_1162_v74).Cardinality()); _1163_i7.Cmp(_hi9) < 0; _1163_i7 = _1163_i7.Plus(_dafny.One) {
			var _1164_v75 _dafny.Map
			_ = _1164_v75
			_1164_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_i7, _dafny.IntOfInt64(223))
			_1164_v75 = (_1164_v75).Update(_dafny.IntOfInt64(-742), _1060_v1)
			var _1165_v76 _dafny.Set
			_ = _1165_v76
			_1165_v76 = _dafny.SetOf(_1060_v1, _1163_i7, _dafny.IntOfUint32((_1059_v0).Cardinality()), _1163_i7, _dafny.IntOfInt64(924))
			var _1166_v77 _dafny.Map
			_ = _1166_v77
			_1166_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				if p0 {
					return _1165_v76
				}
				return _1165_v76
			})(), _1162_v74)
			var _1167_v79 _dafny.Map
			_ = _1167_v79
			_1167_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1162_v74)
			_1166_v77 = (_1166_v77).Update((func() _dafny.Set {
				if p0 {
					return _1165_v76
				}
				return func() _dafny.Set {
					var _coll35 = _dafny.NewBuilder()
					_ = _coll35
					for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(452), _dafny.IntOfInt64(-33))); ; {
						_compr_35, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _1168_v78 _dafny.Int
						_1168_v78 = interface{}(_compr_35).(_dafny.Int)
						if ((_dafny.IntOfInt64(452)).Cmp(_1168_v78) <= 0) && ((_1168_v78).Cmp(_dafny.IntOfInt64(-33)) < 0) {
							_coll35.Add((_1168_v78).Minus(_dafny.IntOfInt64(265)))
						}
					}
					return _coll35.ToSet()
				}()
			})(), (func() _dafny.Sequence {
				if (_1167_v79).Contains(p0) {
					return (_1167_v79).Get(p0).(_dafny.Sequence)
				}
				return _1162_v74
			})())
			_1060_v1 = _1060_v1
			var _1169_v80 _dafny.Array
			_ = _1169_v80
			var _nw207 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw207
			_1169_v80 = _nw207
			var _1170_v81 _dafny.Sequence
			_ = _1170_v81
			_1170_v81 = _dafny.SeqOf(_1169_v80)
			r2 = !_dafny.Companion_Sequence_.Equal(_1170_v81, _1170_v81)
		}
		r0 = Companion_Default___.Fm8(_1062_v3, globalState)
		var _1171_v82 D4
		_ = _1171_v82
		_1171_v82 = Companion_D4_.Create_DC10_(p0)
		r0 = (_1171_v82).Dtor_cf16()
		var _1172_v83 _dafny.Array
		_ = _1172_v83
		var _nw208 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw208
		_1172_v83 = _nw208
		r1 = _1172_v83
		var _1173_v84 _dafny.Map
		_ = _1173_v84
		_1173_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1060_v1, _1162_v74)
		r2 = Companion_Default___.Fm8(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
			if (_1173_v84).Contains(_1060_v1) {
				return (_1173_v84).Get(_1060_v1).(_dafny.Sequence)
			}
			return _1162_v74
		})()), globalState)
		return r0, r1, r2
	}
}

// End of class C8
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
