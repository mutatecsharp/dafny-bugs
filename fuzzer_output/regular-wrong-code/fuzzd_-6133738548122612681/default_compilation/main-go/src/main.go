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
	return (Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(984))).Cardinality()))).Cardinality(), (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(602), _dafny.IntOfInt64(705))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(602)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(705)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_0_v0, _dafny.IntOfInt64(803)), _dafny.SeqOf(_dafny.IntOfInt64(-786)))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())).Minus((_dafny.IntOfInt64(376)).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fsmqogqg")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(32))).Cardinality())).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(756)))).Intersection(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(43), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(874))).Cardinality())).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(268), _dafny.IntOfInt64(281)))))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) bool {
	return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(714))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality())).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hidhmds")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(427))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i1 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('b')))).Cardinality()
	}))).Cardinality()))) > 0
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('v')
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((Companion_D8_.Create_DC26_(Companion_D1_.Create_DC3_(_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('o'), _dafny.CodePoint('r'), _dafny.CodePoint('x'))), _dafny.UnicodeSeqOfUtf8Bytes("cu"))).Dtor_cf42(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ccsyvjc"), _dafny.UnicodeSeqOfUtf8Bytes("rdcttheo")))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(129), _dafny.IntOfInt64(577))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(711), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, ((Companion_D8_.Create_DC23_(_dafny.SetOf(true))).Dtor_cf36()).Cardinality())).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(202), _dafny.IntOfInt64(-405))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v0 _dafny.Int
			_3_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(202)).Cmp(_3_v0) <= 0) && ((_3_v0).Cmp(_dafny.IntOfInt64(-405)) < 0) {
				_coll1.Add(Companion_Default___.SafeModuloInt(_3_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(684), true)).Cardinality()), !(true))
			}
		}
		return _coll1.ToMap()
	}()).Cardinality()), ((Companion_D19_.Create_DC44_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false))).Dtor_cf67()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(-417))
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false)).Union((_dafny.MultiSetOf(true, false)).Union(_dafny.MultiSetOf(true, true, true, !(!(!(false))), true)))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(155))).Cardinality())).Times(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), (_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())).Minus((_dafny.MultiSetOf(_dafny.CodePoint('n'), _dafny.CodePoint('c'))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Map {
	if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vnbdljh")).Cardinality())))).Cardinality()).Cmp(_dafny.IntOfInt64(286)) < 0 {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-275), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(495), !(false)))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-59), !(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.CodePoint('p'))).Cardinality())), true))
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(true), _dafny.SeqOf(!(true), true, true)), (_dafny.IntOfInt64(-570)).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sx")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) _dafny.Set {
	return (Companion_D6_.Create_DC16_(func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(418), _dafny.IntOfInt64(-290))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(418)).Cmp(_4_v0) <= 0) && ((_4_v0).Cmp(_dafny.IntOfInt64(-290)) < 0) {
				_coll2.Add((_4_v0).Plus((_dafny.MultiSetOf(_dafny.IntOfInt64(891), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(661))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))).Cardinality()))).Cardinality()))
			}
		}
		return _coll2.ToSet()
	}())).Dtor_cf27()
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true), true), _dafny.SeqOf(false, true)), _dafny.SeqOf(true))
}
func (_static *CompanionStruct_Default___) Fm20(globalState *GlobalState) _dafny.Map {
	return ((func() D20 {
		if false {
			return Companion_D20_.Create_DC48_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(!(false)), false, !(false)), _dafny.IntOfInt64(812)))
		}
		return Companion_D20_.Create_DC48_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), _dafny.IntOfInt64(930)))
	})()).Dtor_cf75()
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, !(false))).Cardinality()))), ((_dafny.SetOf(_dafny.IntOfInt64(790))).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(494))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_6_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))).Cardinality())), (Companion_D10_.Create_DC29_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality(), false)).Dtor_cf46())
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(309))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(132))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_8_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))))
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(755))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	})), _dafny.UnicodeSeqOfUtf8Bytes("ytfqd")))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-428))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_10_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(814)
	}))
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC13_((_dafny.IntOfInt64(994)).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(38), _dafny.IntOfInt64(-926))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _11_v0 _dafny.Int
			_11_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(38)).Cmp(_11_v0) <= 0) && ((_11_v0).Cmp(_dafny.IntOfInt64(-926)) < 0) {
				_coll3.Add((_11_v0).Plus(_dafny.IntOfInt64(960)))
			}
		}
		return _coll3.ToSet()
	}()).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("folgy"), _dafny.UnicodeSeqOfUtf8Bytes("q")))
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC14_((Companion_D4_.Create_DC14_(Companion_D4_.Create_DC14_(Companion_D4_.Create_DC12_(_dafny.SeqOf((_dafny.MultiSetOf(false)).Cardinality()))))).Dtor_cf25())
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC6_()
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(!(false), true, true, !(true))).Difference(_dafny.SetOf(!((Companion_D1_.Create_DC4_(true, false, _dafny.CodePoint('e'), _dafny.IntOfInt64(554), _dafny.IntOfInt64(784))).Dtor_cf8()), true, false, false))).Difference(_dafny.SetOf(true, false, true, true, !(true)))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-424), (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(643))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_12_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(398)
	})))).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(58), _dafny.IntOfInt64(520))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xb")).Cardinality()))), _dafny.IntOfInt64(311))).Cardinality())), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(289))).Cardinality())))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(19), (_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(32), _dafny.IntOfInt64(765)), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(684))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_13_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-402))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_14_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		}))).Cardinality())
	}))).Cardinality()), (_dafny.SetOf(true, false)).Cardinality(), _dafny.IntOfInt64(-421))))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(232), _dafny.IntOfInt64(455))); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _15_v1 _dafny.Int
				_15_v1 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(232)).Cmp(_15_v1) <= 0) && ((_15_v1).Cmp(_dafny.IntOfInt64(455)) < 0) {
					_coll5.Add(Companion_Default___.SafeModuloInt(_15_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality())))
				}
			}
			return _coll5.ToSet()
		}()).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _16_v0 _dafny.Int
			_16_v0 = interface{}(_compr_4).(_dafny.Int)
			if (func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(232), _dafny.IntOfInt64(455))); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _17_v1 _dafny.Int
					_17_v1 = interface{}(_compr_6).(_dafny.Int)
					if ((_dafny.IntOfInt64(232)).Cmp(_17_v1) <= 0) && ((_17_v1).Cmp(_dafny.IntOfInt64(455)) < 0) {
						_coll6.Add(Companion_Default___.SafeModuloInt(_17_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality())))
					}
				}
				return _coll6.ToSet()
			}()).Contains(_16_v0) {
				_coll4.Add(Companion_Default___.SafeDivisionInt(_16_v0, _dafny.IntOfInt64(845)), _dafny.CodePoint('s'))
			}
		}
		return _coll4.ToMap()
	}()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(393))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))).Cardinality())), _dafny.CodePoint('e'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(695), _dafny.CodePoint('k'))))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC19_(Companion_D6_.Create_DC19_(Companion_D6_.Create_DC18_(false)))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(false, true, _dafny.CodePoint('r'), _dafny.IntOfInt64(-724), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-318))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_19_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))).Cardinality())), !(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(true, !(false), _dafny.CodePoint('l'), _dafny.IntOfInt64(-368), _dafny.IntOfInt64(-261)), false))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 _dafny.Int, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC10_()
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("aqy"))).Difference(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("bjuydaqli")))).Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _20_v0 _dafny.Sequence
			_20_v0 = interface{}(_compr_7).(_dafny.Sequence)
			if ((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("aqy"))).Difference(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("bjuydaqli")))).Contains(_20_v0) {
				_coll7.Add(_20_v0, Companion_D1_.Create_DC5_(!(false), !(true), !(true)))
			}
		}
		return _coll7.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_21_i0 _dafny.Int) _dafny.Set {
		return _dafny.SetOf((_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(51))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rvhgqeka")).Cardinality()))
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(348))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_22_i1 _dafny.Int) _dafny.Set {
		return _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.IntOfInt64(-183), (_dafny.Zero).Minus(_dafny.IntOfInt64(-324)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf(false, false), _dafny.MultiSetOf(!(true), true))).Cardinality()))).Cardinality()), _dafny.IntOfInt64(-687))
	})))
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-904))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_23_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D12_.Create_DC33_(_dafny.CodePoint('r'), !(false))).Dtor_cf55(), false)
	})))
}
func (_static *CompanionStruct_Default___) Fm37(globalState *GlobalState) D8 {
	if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ycqwr"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(737))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_24_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))) {
		return Companion_D8_.Create_DC26_(Companion_D1_.Create_DC3_((Companion_D8_.Create_DC26_(Companion_D1_.Create_DC3_(_dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('b'), _dafny.CodePoint('r'))), _dafny.UnicodeSeqOfUtf8Bytes("ndi"))).Dtor_cf42()), _dafny.UnicodeSeqOfUtf8Bytes("w"))
	} else {
		return Companion_D8_.Create_DC26_(Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(693))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_25_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(263))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}(func(_26_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm38(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_D3_.Create_DC11_(_dafny.IntOfInt64(442), false, true))
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(80), _dafny.IntOfInt64(406))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _27_v0 _dafny.Int
			_27_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(80)).Cmp(_27_v0) <= 0) && ((_27_v0).Cmp(_dafny.IntOfInt64(406)) < 0) {
				_coll8.Add((_27_v0).Minus(_dafny.IntOfInt64(-269)), _dafny.SetOf(_dafny.IntOfInt64(522), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gc")).Cardinality())))
			}
		}
		return _coll8.ToMap()
	}()).Merge((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(977), _dafny.SetOf(_dafny.IntOfInt64(826)))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.SetOf(_dafny.IntOfInt64(380)))
	})())
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(828))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_28_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("q")
	})), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("fil"))), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("canhopsk")))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nw")).Cardinality()))).Difference((Companion_D13_.Create_DC35_(_dafny.MultiSetOf(_dafny.IntOfInt64(98)))).Dtor_cf57()), _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('s'))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bcleahc")).Cardinality()))).Cardinality())), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-165)), (_dafny.SetOf(true, true, true)).Cardinality())).Cardinality())).Intersection(_dafny.MultiSetOf((func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(709), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jjwyqqi")).Cardinality()))).Cardinality()))).Keys().Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _29_v0 _dafny.Int
			_29_v0 = interface{}(_compr_9).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(709), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jjwyqqi")).Cardinality()))).Cardinality()))).Contains(_29_v0) {
				_coll9.Add(Companion_Default___.SafeModuloInt(_29_v0, _dafny.IntOfInt64(620)), false)
			}
		}
		return _coll9.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(-251), _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uqgceiacv")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _dafny.IntOfInt64(-853)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), (_dafny.SetOf(true)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(!(false) || (false), false, _dafny.CodePoint('i'), ((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(637))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg20 _dafny.Int) interface{} {
			return coer20(arg20)
		}
	}(func(_30_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(true, !(!(true)))).Cardinality())
	})))).Intersection(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(439))))).Cardinality(), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(124), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ahdj")).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var _31_v0 _dafny.Map
	_ = _31_v0
	_31_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(4))
	var _32_v1 _dafny.Sequence
	_ = _32_v1
	_32_v1 = _dafny.UnicodeSeqOfUtf8Bytes("nfp")
	_31_v0 = (_31_v0).Update(_dafny.IntOfUint32((_32_v1).Cardinality()), p1)
	var _33_v2 _dafny.MultiSet
	_ = _33_v2
	_33_v2 = _dafny.MultiSetOf(p1)
	var _34_v3 _dafny.Sequence
	_ = _34_v3
	_34_v3 = _dafny.SeqOf((_33_v2).Update((_31_v0).Cardinality(), Companion_Default___.Abs(Companion_Default___.Fm0(p0, p0, globalState))))
	var _35_v4 bool
	_ = _35_v4
	_35_v4 = false
	var _36_v5 _dafny.Set
	_ = _36_v5
	_36_v5 = _dafny.SetOf(_35_v4, _35_v4, _35_v4, true)
	var _37_v6 _dafny.CodePoint
	_ = _37_v6
	_37_v6 = _dafny.CodePoint('w')
	var _38_v7 _dafny.Map
	_ = _38_v7
	_38_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_37_v6, p1)
	var _39_v8 _dafny.Sequence
	_ = _39_v8
	_39_v8 = _dafny.SeqOf(_38_v7)
	var _40_v9 T0
	_ = _40_v9
	var _nw0 *C1 = New_C1_()
	_ = _nw0
	_nw0.Ctor__(_34_v3, p1, (_36_v5).Cardinality(), _39_v8)
	_40_v9 = _nw0
	var _41_v10 _dafny.Map
	_ = _41_v10
	_41_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _40_v9)
	var _42_v11 _dafny.Array
	_ = _42_v11
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
	_ = _nw1
	_42_v11 = _nw1
	var _43_v12 *C4
	_ = _43_v12
	var _nw2 *C4 = New_C4_()
	_ = _nw2
	_nw2.Ctor__(_42_v11, (_40_v9).F24(), (_40_v9).F25())
	_43_v12 = _nw2
	var _44_v13 _dafny.MultiSet
	_ = _44_v13
	_44_v13 = _dafny.MultiSetOf(_43_v12, _43_v12, _43_v12, _43_v12, _43_v12)
	var _45_v14 T0
	_ = _45_v14
	_45_v14 = _40_v9
	_41_v10 = (_41_v10).Update(Companion_Default___.Fm0((func() _dafny.Int {
		if (_44_v13).Contains(_43_v12) {
			return (_44_v13).Multiplicity(_43_v12)
		}
		return _dafny.IntOfInt64(235)
	})(), _dafny.IntOfInt64(986), globalState), (_40_v9))
	var _46_v15 _dafny.Map
	_ = _46_v15
	_46_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_43_v12).F28(), (_40_v9).F24())
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen(((_43_v12).F28()), 0))
	_ = _index0
	((_43_v12).F28()).ArraySet1((_46_v15).Cardinality(), (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen(((_43_v12).F28()), 0))
	_ = _index1
	((_43_v12).F28()).ArraySet1(p1, (_index1).Int())
	var _47_v16 _dafny.Map
	_ = _47_v16
	_47_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v4, (_43_v12).F28())
	_47_v16 = (_47_v16).Merge((_47_v16).Merge(_47_v16))
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-669))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}((func(_48_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_49_i1 _dafny.Int) _dafny.CodePoint {
			return _48_v6
		}
	})(_37_v6)))).Cardinality())
	_ = _hi0
	for _50_i0 := (_dafny.Zero).Minus(((_43_v12).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen(((_43_v12).F28()), 0))).Int()).(_dafny.Int)); _50_i0.Cmp(_hi0) < 0; _50_i0 = _50_i0.Plus(_dafny.One) {
		var _51_v17 _dafny.Array
		_ = _51_v17
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_0
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_52_v1 _dafny.Sequence) func(_dafny.Int) bool {
				return func(_53_i2 _dafny.Int) bool {
					return !_dafny.Companion_Sequence_.Contains(_52_v1, _dafny.CodePoint('w'))
				}
			})(_32_v1)
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
		_51_v17 = _nw3
		var _54_v18 _dafny.Map
		_ = _54_v18
		_54_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_40_v9).F24(), (func() _dafny.Array {
			if true {
				return _51_v17
			}
			return _51_v17
		})())
		_51_v17 = (func() _dafny.Array {
			if (_54_v18).Contains((_40_v9).F24()) {
				return (_54_v18).Get((_40_v9).F24()).(_dafny.Array)
			}
			return _51_v17
		})()
		var _55_v19 _dafny.MultiSet
		_ = _55_v19
		_55_v19 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(694))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}((func(_56_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_57_i3 _dafny.Int) _dafny.CodePoint {
				return _56_v6
			}
		})(_37_v6))))
		var _58_v20 _dafny.Sequence
		_ = _58_v20
		_58_v20 = _dafny.SeqOf(_50_i0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ipsloox")).Cardinality()))
		var _59_v21 _dafny.Set
		_ = _59_v21
		_59_v21 = _dafny.SetOf((_55_v19).Cardinality(), _dafny.IntOfUint32((_58_v20).Cardinality()))
		var _60_v22 D10
		_ = _60_v22
		_60_v22 = Companion_D10_.Create_DC29_((_40_v9).F24(), _35_v4)
		(globalState).F18 = (func() _dafny.Int {
			if _35_v4 {
				return (((_43_v12).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen(((_43_v12).F28()), 0))).Int()).(_dafny.Int)).Minus((_59_v21).Cardinality())
			}
			return Companion_Default___.SafeModuloInt((_60_v22).Dtor_cf45(), ((_43_v12).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen(((_43_v12).F28()), 0))).Int()).(_dafny.Int))
		})()
		var _61_v23 _dafny.Array
		_ = _61_v23
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
		_ = _nw4
		_61_v23 = _nw4
		var _62_v24 _dafny.Sequence
		_ = _62_v24
		_62_v24 = _dafny.SeqOf(_35_v4)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_61_v23), 0))
		_ = _index2
		(_61_v23).ArraySet1(_62_v24, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_61_v23), 0))
		_ = _index3
		var _rhs0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_62_v24, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_35_v4, _35_v4, _35_v4, _35_v4), (Companion_Default___.SafeIndex(_50_i0, _dafny.IntOfUint32((_dafny.SeqOf(_35_v4, _35_v4, _35_v4, _35_v4)).Cardinality()))).Uint32(), _35_v4))
		_ = _rhs0
		var _rhs1 _dafny.Int = (_dafny.Zero).Minus(((_40_v9).F24()).Plus(p0))
		_ = _rhs1
		var _rhs2 _dafny.Set = _36_v5
		_ = _rhs2
		var _lhs0 _dafny.Array = _61_v23
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_61_v23), 0))
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
		_lhs2.F1 = _rhs1
		_36_v5 = _rhs2
		var _63_v25 _dafny.MultiSet
		_ = _63_v25
		_63_v25 = _dafny.MultiSetOf(_35_v4)
		var _64_v26 _dafny.Map
		_ = _64_v26
		_64_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_63_v25).Difference(_63_v25), false)
		_64_v26 = (_64_v26).Update(_63_v25, (_35_v4) == (true))
	}
	for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_42_v11), 0))); ; {
		_guard_loop_0, _ok10 := _iter10()
		if !_ok10 {
			break
		}
		var _65_i4 _dafny.Int
		_65_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_65_i4).Sign() != -1) && ((_65_i4).Cmp(_dafny.ArrayLen((_42_v11), 0)) < 0)) {
			(_42_v11).ArraySet1((_65_i4).Times(p1), (_65_i4).Int())
		}
	}
	var _66_v27 _dafny.Map
	_ = _66_v27
	_66_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v4, p0)
	var _67_v28 D4
	_ = _67_v28
	_67_v28 = Companion_D4_.Create_DC13_(_dafny.IntOfInt64(-145), _66_v27, _32_v1)
	r0 = (_dafny.Zero).Minus(_dafny.IntOfUint32(((_67_v28).Dtor_cf24()).Cardinality()))
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _68_v0 _dafny.Int
	_ = _68_v0
	_68_v0 = _dafny.IntOfInt64(590)
	var _69_v1 _dafny.Map
	_ = _69_v1
	_69_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _68_v0)
	var _70_v2 bool
	_ = _70_v2
	_70_v2 = true
	var _71_v3 _dafny.Array
	_ = _71_v3
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(24)
	_ = _len0_1
	var _nw5 _dafny.Array
	_ = _nw5
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw5 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = func(_72_i0 _dafny.Int) _dafny.Int {
			return (_72_i0).Minus(_dafny.IntOfInt64(-719))
		}
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
	_71_v3 = _nw5
	var _73_v4 D0
	_ = _73_v4
	_73_v4 = Companion_D0_.Create_DC0_(_71_v3)
	var _74_v5 _dafny.Map
	_ = _74_v5
	_74_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v3, (_73_v4).Dtor_cf0())
	var _75_v6 _dafny.Sequence
	_ = _75_v6
	_75_v6 = _dafny.SeqOf(_70_v2)
	var _76_v7 _dafny.CodePoint
	_ = _76_v7
	_76_v7 = _dafny.CodePoint('c')
	var _77_v8 D1
	_ = _77_v8
	_77_v8 = Companion_D1_.Create_DC3_(_dafny.SeqOf(_76_v7, _dafny.CodePoint('h')))
	var _78_v9 _dafny.MultiSet
	_ = _78_v9
	_78_v9 = _dafny.MultiSetOf(_76_v7)
	var _79_v10 _dafny.Set
	_ = _79_v10
	_79_v10 = _dafny.SetOf(_dafny.MultiSetFromSeq((_77_v8).Dtor_cf7()), _78_v9)
	var _80_v11 _dafny.Array
	_ = _80_v11
	var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
	_ = _nw6
	_80_v11 = _nw6
	var _81_v12 _dafny.Map
	_ = _81_v12
	_81_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-991))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}((func(_82_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_83_i1 _dafny.Int) _dafny.Int {
			return _82_v0
		}
	})(_68_v0)))).Cardinality())), _80_v11)
	var _84_v13 _dafny.Set
	_ = _84_v13
	_84_v13 = _dafny.SetOf(_70_v2, _70_v2, _70_v2)
	var _85_v14 _dafny.Sequence
	_ = _85_v14
	_85_v14 = _dafny.UnicodeSeqOfUtf8Bytes("nbntyekh")
	var _86_v15 D1
	_ = _86_v15
	_86_v15 = Companion_D1_.Create_DC4_(_70_v2, _70_v2, (_85_v14).Select((Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_85_v14).Cardinality()))).Uint32()).(_dafny.CodePoint), _68_v0, _68_v0)
	var _87_v16 _dafny.Map
	_ = _87_v16
	_87_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(382), _68_v0))
	var _88_globalState *GlobalState
	_ = _88_globalState
	var _nw7 *GlobalState = New_GlobalState_()
	_ = _nw7
	_nw7.Ctor__(true, _dafny.IntOfInt64(554), true, (_69_v1).Update(_70_v2, _68_v0), false, _dafny.IntOfInt64(-61), _dafny.CodePoint('c'), _74_v5, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_70_v2, _70_v2, !(false)), _75_v6), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_68_v0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_70_v2, _70_v2, !(false)), _75_v6)).Cardinality()))).Uint32(), !(_70_v2)), _dafny.IntOfInt64(889), _dafny.IntOfInt64(204), _79_v10, _dafny.IntOfInt64(127), _81_v12, _dafny.IntOfInt64(996), _dafny.IntOfInt64(982), _84_v13, _85_v14, _dafny.IntOfInt64(-607), _dafny.MultiSetOf(_80_v11, _80_v11), _75_v6, _dafny.SeqOf(_70_v2, (_86_v15).Dtor_cf8()), (func() _dafny.Map {
		if (_87_v16).Contains(_68_v0) {
			return (_87_v16).Get(_68_v0).(_dafny.Map)
		}
		return func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(199), _dafny.IntOfInt64(972))); ; {
				_compr_10, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _89_v17 _dafny.Int
				_89_v17 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(199)).Cmp(_89_v17) <= 0) && ((_89_v17).Cmp(_dafny.IntOfInt64(972)) < 0) {
					_coll10.Add((_89_v17).Minus(_68_v0), _68_v0)
				}
			}
			return _coll10.ToMap()
		}()
	})(), _dafny.IntOfInt64(790))
	_88_globalState = _nw7
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
	_ = _index4
	(_80_v11).ArraySet1((_68_v0).Cmp(_68_v0) <= 0, (_index4).Int())
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
	_ = _index5
	(_80_v11).ArraySet1(_70_v2, (_index5).Int())
	_76_v7 = _dafny.CodePoint('u')
	var _90_v18 _dafny.MultiSet
	_ = _90_v18
	_90_v18 = _dafny.MultiSetOf(!((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)), (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool))
	var _91_v19 _dafny.Map
	_ = _91_v19
	_91_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v2, false)
	var _92_v20 _dafny.Sequence
	_ = _92_v20
	_92_v20 = _dafny.SeqOf(_68_v0, (_91_v19).Cardinality())
	if (Companion_Default___.SafeDivisionInt(_68_v0, (func() _dafny.Int {
		if (_90_v18).Contains(!((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool))) {
			return (_90_v18).Multiplicity(!((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)))
		}
		return _68_v0
	})())).Cmp(Companion_Default___.Fm0(_dafny.IntOfUint32((_92_v20).Cardinality()), _68_v0, _88_globalState)) >= 0 {
		var _93_v21 _dafny.MultiSet
		_ = _93_v21
		_93_v21 = _dafny.MultiSetOf(_68_v0)
		var _94_v22 _dafny.Map
		_ = _94_v22
		_94_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v7, _93_v21)
		var _95_v23 _dafny.Sequence
		_ = _95_v23
		_95_v23 = _dafny.SeqOf(_93_v21)
		(_88_globalState).F14 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_93_v21, Companion_Default___.Fm1(_68_v0, _68_v0, _88_globalState), (func() _dafny.MultiSet {
			if (_94_v22).Contains(_76_v7) {
				return (_94_v22).Get(_76_v7).(_dafny.MultiSet)
			}
			return _93_v21
		})()), _95_v23)).Cardinality())
		_73_v4 = Companion_D0_.Create_DC0_(_71_v3)
		var _96_v24 _dafny.Set
		_ = _96_v24
		_96_v24 = _dafny.SetOf(Companion_Default___.Fm0(_dafny.IntOfUint32((_85_v14).Cardinality()), _68_v0, _88_globalState))
		_96_v24 = _96_v24
		var _97_v25 _dafny.Map
		_ = _97_v25
		_97_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _84_v13)
		var _98_v26 _dafny.Map
		_ = _98_v26
		_98_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_84_v13).IsProperSubsetOf((func() _dafny.Set {
			if (_97_v25).Contains(true) {
				return (_97_v25).Get(true).(_dafny.Set)
			}
			return _dafny.SetOf(_70_v2, Companion_Default___.Fm2(_dafny.IntOfInt64(659), _dafny.IntOfUint32((_75_v6).Cardinality()), _76_v7, _88_globalState))
		})()), _71_v3)
		var _99_v27 _dafny.Array
		_ = _99_v27
		var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
		_ = _nw8
		_99_v27 = _nw8
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_99_v27), 0))
		_ = _index6
		(_99_v27).ArraySet1CodePoint((_85_v14).Select((Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_85_v14).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index6).Int())
		var _100_v28 _dafny.MultiSet
		_ = _100_v28
		_100_v28 = _dafny.MultiSetOf(_99_v27)
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_99_v27), 0))
		_ = _index7
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
		_ = _index8
		var _rhs3 _dafny.Map = _98_v26
		_ = _rhs3
		var _rhs4 _dafny.CodePoint = Companion_Default___.Fm3(_88_globalState)
		_ = _rhs4
		var _rhs5 _dafny.Int = _68_v0
		_ = _rhs5
		var _rhs6 bool = ((_100_v28).Union(_100_v28)).IsSubsetOf(_100_v28)
		_ = _rhs6
		var _rhs7 _dafny.Int = _68_v0
		_ = _rhs7
		var _lhs3 _dafny.Array = _99_v27
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(707), _dafny.ArrayLen((_99_v27), 0))
		_ = _lhs4
		var _lhs5 *GlobalState = _88_globalState
		_ = _lhs5
		var _lhs6 _dafny.Array = _80_v11
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
		_ = _lhs7
		var _lhs8 *GlobalState = _88_globalState
		_ = _lhs8
		_98_v26 = _rhs3
		(_lhs3).ArraySet1CodePoint(_rhs4, (_lhs4).Int())
		_lhs5.F9 = _rhs5
		(_lhs6).ArraySet1(_rhs6, (_lhs7).Int())
		_lhs8.F1 = _rhs7
		(_88_globalState).F9 = Companion_Default___.SafeModuloInt((_68_v0).Plus(_dafny.IntOfInt64(765)), ((_84_v13).Difference(_84_v13)).Cardinality())
	} else {
		var _101_v29 _dafny.Array
		_ = _101_v29
		var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
		_ = _nw9
		_101_v29 = _nw9
		var _102_v30 _dafny.Map
		_ = _102_v30
		_102_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _71_v3)
		var _103_v31 D0
		_ = _103_v31
		_103_v31 = Companion_D0_.Create_DC2_((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), (_dafny.Zero).Minus(_68_v0), _76_v7)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_101_v29), 0))
		_ = _index9
		(_101_v29).ArraySet1((func() _dafny.Array {
			if (_102_v30).Contains(!((_103_v31).Dtor_cf4())) {
				return (_102_v30).Get(!((_103_v31).Dtor_cf4())).(_dafny.Array)
			}
			return _71_v3
		})(), (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_101_v29), 0))
		_ = _index10
		(_101_v29).ArraySet1(_71_v3, (_index10).Int())
		var _104_v32 D0
		_ = _104_v32
		_104_v32 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(215), _68_v0, (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool))
		var _105_v33 _dafny.Map
		_ = _105_v33
		_105_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _85_v14)
		var _106_v34 _dafny.Array
		_ = _106_v34
		var _nwElement0_0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_85_v14, _85_v14)
		_ = _nwElement0_0
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(21))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_0, 0)
		(_nw10).ArraySet1(_dafny.Companion_Sequence_.Update(_85_v14, (Companion_Default___.SafeIndex(Companion_Default___.Fm0(_68_v0, _68_v0, _88_globalState), _dafny.IntOfUint32((_85_v14).Cardinality()))).Uint32(), _76_v7), 1)
		(_nw10).ArraySet1(_85_v14, 2)
		(_nw10).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("orkpvq"), (Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("orkpvq")).Cardinality()))).Uint32(), Companion_Default___.Fm3(_88_globalState)), 3)
		(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_85_v14, Companion_Default___.Fm4((_104_v32).Dtor_cf3(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_92_v20).Cardinality())), _76_v7, Companion_Default___.Fm2(_68_v0, _68_v0, _76_v7, _88_globalState), _88_globalState)), 4)
		(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_85_v14, _85_v14), 5)
		(_nw10).ArraySet1(_85_v14, 6)
		(_nw10).ArraySet1(_85_v14, 7)
		(_nw10).ArraySet1(_85_v14, 8)
		(_nw10).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-260))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}((func(_107_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_108_i2 _dafny.Int) _dafny.CodePoint {
				return _107_v7
			}
		})(_76_v7))), 9)
		(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_105_v33).Contains((_84_v13).Cardinality()) {
				return (_105_v33).Get((_84_v13).Cardinality()).(_dafny.Sequence)
			}
			return _85_v14
		})(), _dafny.Companion_Sequence_.Update(_85_v14, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_92_v20).Cardinality()), _dafny.IntOfUint32((_85_v14).Cardinality()))).Uint32(), _76_v7)), 10)
		(_nw10).ArraySet1(_85_v14, 11)
		(_nw10).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jmsoaylng"), 12)
		(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_85_v14, _85_v14), 13)
		(_nw10).ArraySet1(_85_v14, 14)
		(_nw10).ArraySet1(_85_v14, 15)
		(_nw10).ArraySet1(_85_v14, 16)
		(_nw10).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("arypuvscv"), 17)
		(_nw10).ArraySet1(_85_v14, 18)
		(_nw10).ArraySet1(_dafny.Companion_Sequence_.Update(_85_v14, (Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_85_v14).Cardinality()))).Uint32(), _76_v7), 19)
		(_nw10).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("bmic"), 20)
		_106_v34 = _nw10
		var _109_v35 _dafny.Sequence
		_ = _109_v35
		_109_v35 = _dafny.SeqOf(_106_v34, _106_v34)
		_106_v34 = (_109_v35).Select((Companion_Default___.SafeIndex((_86_v15).Dtor_cf11(), _dafny.IntOfUint32((_109_v35).Cardinality()))).Uint32()).(_dafny.Array)
		(_88_globalState).F2 = (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)
		var _110_v36 _dafny.Set
		_ = _110_v36
		_110_v36 = _dafny.SetOf(_85_v14, _85_v14, Companion_Default___.Fm4((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _70_v2)).Cardinality(), _76_v7, _70_v2, _88_globalState))
		var _111_v37 _dafny.Map
		_ = _111_v37
		_111_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_110_v36).Union(_110_v36), ((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)) || ((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)))
		_111_v37 = (_111_v37).Update(_110_v36, !((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)))
		var _112_v38 _dafny.Map
		_ = _112_v38
		_112_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _76_v7)
		_70_v2 = Companion_Default___.Fm2((_68_v0).Plus(_68_v0), Companion_Default___.SafeDivisionInt(_68_v0, _68_v0), (func() _dafny.CodePoint {
			if (_112_v38).Contains((_104_v32).Dtor_cf1()) {
				return (_112_v38).Get((_104_v32).Dtor_cf1()).(_dafny.CodePoint)
			}
			return _76_v7
		})(), _88_globalState)
	}
	var _113_v39 D1
	_ = _113_v39
	_113_v39 = Companion_D1_.Create_DC5_((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), (_68_v0).Cmp(_68_v0) >= 0, (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool))
	_113_v39 = _113_v39
	var _hi1 _dafny.Int = (_92_v20).Select((Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_92_v20).Cardinality()))).Uint32()).(_dafny.Int)
	_ = _hi1
	for _114_i3 := _68_v0; _114_i3.Cmp(_hi1) < 0; _114_i3 = _114_i3.Plus(_dafny.One) {
		_85_v14 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ug"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(922))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_115_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		}))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_68_v0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ug"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(922))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_116_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})))).Cardinality()))).Uint32(), _76_v7), _85_v14)
		var _117_v40 _dafny.Int
		_ = _117_v40
		var _out0 _dafny.Int
		_ = _out0
		_out0 = Companion_Default___.M0(_114_i3, _114_i3, _88_globalState)
		_117_v40 = _out0
		var _118_v41 _dafny.Map
		_ = _118_v41
		_118_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _117_v40)
		var _119_v43 _dafny.Sequence
		_ = _119_v43
		_119_v43 = _dafny.SeqOf(_118_v41, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _dafny.IntOfInt64(-223)), func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter12 := _dafny.Iterate((_92_v20).Elements()); ; {
				_compr_11, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _120_v42 _dafny.Int
				_120_v42 = interface{}(_compr_11).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_92_v20, _120_v42) {
					_coll11.Add(Companion_Default___.SafeDivisionInt(_120_v42, _dafny.IntOfInt64(135)), _dafny.IntOfInt64(597))
				}
			}
			return _coll11.ToMap()
		}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _114_i3), _118_v41)
		var _121_v44 _dafny.Map
		_ = _121_v44
		_121_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v2, _119_v43)
		_121_v44 = (_121_v44).Update(_70_v2, _119_v43)
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_71_v3), 0))
		_ = _index11
		(_71_v3).ArraySet1(_68_v0, (_index11).Int())
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_71_v3), 0))
		_ = _index12
		(_71_v3).ArraySet1(_117_v40, (_index12).Int())
	}
	(_88_globalState).F18 = _dafny.IntOfUint32((_85_v14).Cardinality())
	var _122_v45 _dafny.Array
	_ = _122_v45
	var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
	_ = _nw11
	_122_v45 = _nw11
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))
	_ = _index13
	(_122_v45).ArraySet1(_85_v14, (_index13).Int())
	var _123_v46 _dafny.Map
	_ = _123_v46
	_123_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_85_v14, _85_v14, _dafny.UnicodeSeqOfUtf8Bytes("ncd"), _85_v14, _85_v14), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("lbmfhj"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-133), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lbmfhj")).Cardinality()))).Uint32(), _76_v7))
	var _124_v47 _dafny.Sequence
	_ = _124_v47
	_124_v47 = _dafny.SeqOf(_85_v14)
	var _125_v48 _dafny.Sequence
	_ = _125_v48
	_125_v48 = _dafny.SeqOf(_85_v14, (_124_v47).Select((Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_124_v47).Cardinality()))).Uint32()).(_dafny.Sequence), _85_v14, _85_v14, _85_v14)
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))
	_ = _index14
	(_122_v45).ArraySet1((func() _dafny.Sequence {
		if (_123_v46).Contains(_124_v47) {
			return (_123_v46).Get(_124_v47).(_dafny.Sequence)
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("anrxvxycj")
	})(), (_index14).Int())
	var _rhs8 _dafny.Int = (func() _dafny.Int {
		if Companion_Default___.Fm2((_91_v19).Cardinality(), _68_v0, _76_v7, _88_globalState) {
			return (_dafny.Zero).Minus(_68_v0)
		}
		return (func() _dafny.Int {
			if !(_70_v2) {
				return _68_v0
			}
			return (_dafny.Zero).Minus(_68_v0)
		})()
	})()
	_ = _rhs8
	var _rhs9 bool = !(((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)) == (_70_v2))
	_ = _rhs9
	var _lhs9 *GlobalState = _88_globalState
	_ = _lhs9
	_lhs9.F14 = _rhs8
	_70_v2 = _rhs9
	var _126_v49 _dafny.Set
	_ = _126_v49
	_126_v49 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kyghqhrgx")).Cardinality()))
	var _127_v50 D0
	_ = _127_v50
	_127_v50 = Companion_D0_.Create_DC1_(_68_v0, _68_v0, !((_75_v6).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(760), _126_v49)).Cardinality(), _dafny.IntOfUint32((_75_v6).Cardinality()))).Uint32()).(bool)))
	var _source0 D0 = _127_v50
	_ = _source0
	if _source0.Is_DC1() {
		var _128___mcc_h0 _dafny.Int = _source0.Get_().(D0_DC1).Cf1
		_ = _128___mcc_h0
		var _129___mcc_h1 _dafny.Int = _source0.Get_().(D0_DC1).Cf2
		_ = _129___mcc_h1
		var _130___mcc_h2 bool = _source0.Get_().(D0_DC1).Cf3
		_ = _130___mcc_h2
		var _131_cf3 bool = _130___mcc_h2
		_ = _131_cf3
		var _132_cf2 _dafny.Int = _129___mcc_h1
		_ = _132_cf2
		var _133_cf1 _dafny.Int = _128___mcc_h0
		_ = _133_cf1
		var _134_v51 *C0
		_ = _134_v51
		var _nw12 *C0 = New_C0_()
		_ = _nw12
		_nw12.Ctor__(!(_131_cf3) || (_70_v2))
		_134_v51 = _nw12
		var _135_v52 _dafny.Int
		_ = _135_v52
		var _out1 _dafny.Int
		_ = _out1
		_out1 = Companion_Default___.M0(_133_cf1, Companion_Default___.SafeDivisionInt(_68_v0, _68_v0), _88_globalState)
		_135_v52 = _out1
		_131_cf3 = (_134_v51).F34()
		_131_cf3 = (_134_v51).F34()
	} else if _source0.Is_DC2() {
		var _136___mcc_h3 bool = _source0.Get_().(D0_DC2).Cf4
		_ = _136___mcc_h3
		var _137___mcc_h4 _dafny.Int = _source0.Get_().(D0_DC2).Cf5
		_ = _137___mcc_h4
		var _138___mcc_h5 _dafny.CodePoint = _source0.Get_().(D0_DC2).Cf6
		_ = _138___mcc_h5
		var _139_cf6 _dafny.CodePoint = _138___mcc_h5
		_ = _139_cf6
		var _140_cf5 _dafny.Int = _137___mcc_h4
		_ = _140_cf5
		var _141_cf4 bool = _136___mcc_h3
		_ = _141_cf4
		_70_v2 = !((func() bool {
			if Companion_Default___.Fm2(_68_v0, _dafny.IntOfUint32((_75_v6).Cardinality()), _dafny.CodePoint('t'), _88_globalState) {
				return _141_cf4
			}
			return _70_v2
		})()) || (!((_84_v13).IsSubsetOf(_84_v13)))
		var _142_v53 _dafny.Map
		_ = _142_v53
		_142_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _92_v20)
		_142_v53 = _142_v53
		(_88_globalState).F9 = (_140_cf5).Minus(_dafny.IntOfInt64(969))
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_71_v3), 0))
		_ = _index15
		(_71_v3).ArraySet1(Companion_Default___.SafeDivisionInt(_68_v0, _140_cf5), (_index15).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_71_v3), 0))
		_ = _index16
		(_71_v3).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(776), _140_cf5)), (_index16).Int())
	} else {
		var _143___mcc_h6 _dafny.Array = _source0.Get_().(D0_DC0).Cf0
		_ = _143___mcc_h6
		var _144_cf0 _dafny.Array = _143___mcc_h6
		_ = _144_cf0
		if (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool) {
			_68_v0 = _68_v0
			var _145_v54 _dafny.Array
			_ = _145_v54
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw13
			_145_v54 = _nw13
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_145_v54), 0))
			_ = _index17
			(_145_v54).ArraySet1(_92_v20, (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_145_v54), 0))
			_ = _index18
			(_145_v54).ArraySet1(_92_v20, (_index18).Int())
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_2
			var _nw14 _dafny.Array
			_ = _nw14
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw14 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = (func(_146_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_147_i5 _dafny.Int) _dafny.Int {
						return (_147_i5).Minus(_146_v0)
					}
				})(_68_v0)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw14 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw14).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw14).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_71_v3 = _nw14
			var _148_v55 _dafny.Map
			_ = _148_v55
			_148_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v7, _68_v0)
			var _149_v56 _dafny.Sequence
			_ = _149_v56
			_149_v56 = _dafny.SeqOf(_148_v55)
			var _150_v57 *C4
			_ = _150_v57
			var _nw15 *C4 = New_C4_()
			_ = _nw15
			_nw15.Ctor__(_71_v3, _68_v0, _149_v56)
			_150_v57 = _nw15
			var _151_v58 _dafny.Sequence
			_ = _151_v58
			_151_v58 = _dafny.SeqOf(_150_v57, _150_v57)
			var _152_v59 *C4
			_ = _152_v59
			var _nw16 *C4 = New_C4_()
			_ = _nw16
			_nw16.Ctor__(_144_cf0, _dafny.IntOfUint32((_151_v58).Cardinality()), _149_v56)
			_152_v59 = _nw16
			var _153_v60 _dafny.MultiSet
			_ = _153_v60
			_153_v60 = _dafny.MultiSetOf(_dafny.IntOfInt64(-453), _68_v0)
			var _154_v61 _dafny.Sequence
			_ = _154_v61
			_154_v61 = _dafny.SeqOf(_153_v60, _dafny.MultiSetOf(_68_v0))
			var _155_v62 *C1
			_ = _155_v62
			var _nw17 *C1 = New_C1_()
			_ = _nw17
			_nw17.Ctor__(_154_v61, _68_v0, (func() _dafny.Int {
				if (_69_v1).Contains(true) {
					return (_69_v1).Get(true).(_dafny.Int)
				}
				return (_153_v60).Cardinality()
			})(), _149_v56)
			_155_v62 = _nw17
			var _156_v63 *C1
			_ = _156_v63
			_156_v63 = _155_v62
			var _157_v64 _dafny.Sequence
			_ = _157_v64
			_157_v64 = _dafny.SeqOf(_155_v62)
			var _158_v65 _dafny.Array
			_ = _158_v65
			var _nwElement0_1 *C1 = _155_v62
			_ = _nwElement0_1
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(26))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_1, 0)
			(_nw18).ArraySet1(_155_v62, 1)
			(_nw18).ArraySet1(_155_v62, 2)
			(_nw18).ArraySet1(_155_v62, 3)
			(_nw18).ArraySet1(_155_v62, 4)
			(_nw18).ArraySet1(_155_v62, 5)
			(_nw18).ArraySet1(_155_v62, 6)
			(_nw18).ArraySet1(_155_v62, 7)
			(_nw18).ArraySet1(_155_v62, 8)
			(_nw18).ArraySet1((_156_v63), 9)
			(_nw18).ArraySet1(_155_v62, 10)
			(_nw18).ArraySet1(_155_v62, 11)
			(_nw18).ArraySet1(_155_v62, 12)
			(_nw18).ArraySet1(_155_v62, 13)
			(_nw18).ArraySet1((_157_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_157_v64).Cardinality()))).Uint32()).(*C1), 14)
			(_nw18).ArraySet1(_155_v62, 15)
			(_nw18).ArraySet1(_155_v62, 16)
			(_nw18).ArraySet1(_155_v62, 17)
			(_nw18).ArraySet1((_157_v64).Select((Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_157_v64).Cardinality()))).Uint32()).(*C1), 18)
			(_nw18).ArraySet1((_157_v64).Select((Companion_Default___.SafeIndex(_155_v62.F33, _dafny.IntOfUint32((_157_v64).Cardinality()))).Uint32()).(*C1), 19)
			(_nw18).ArraySet1(_155_v62, 20)
			(_nw18).ArraySet1(_155_v62, 21)
			(_nw18).ArraySet1(_155_v62, 22)
			(_nw18).ArraySet1(_155_v62, 23)
			(_nw18).ArraySet1(_155_v62, 24)
			(_nw18).ArraySet1(_155_v62, 25)
			_158_v65 = _nw18
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_158_v65), 0))
			_ = _index19
			(_158_v65).ArraySet1(_155_v62, (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_158_v65), 0))
			_ = _index20
			var _rhs10 _dafny.CodePoint = _76_v7
			_ = _rhs10
			var _rhs11 _dafny.Int = (_dafny.Zero).Minus(_68_v0)
			_ = _rhs11
			var _rhs12 *C1 = _155_v62
			_ = _rhs12
			var _rhs13 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_155_v62.F33, _dafny.IntOfUint32(((_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _76_v7), _85_v14), _85_v14)
			_ = _rhs13
			var _lhs10 *GlobalState = _88_globalState
			_ = _lhs10
			var _lhs11 _dafny.Array = _158_v65
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_158_v65), 0))
			_ = _lhs12
			_76_v7 = _rhs10
			_lhs10.F18 = _rhs11
			(_lhs11).ArraySet1(_rhs12, (_lhs12).Int())
			_70_v2 = _rhs13
		} else {
			var _159_v67 _dafny.Set
			_ = _159_v67
			_159_v67 = _dafny.SetOf(_76_v7, _76_v7)
			var _160_v68 D4
			_ = _160_v68
			_160_v68 = Companion_D4_.Create_DC12_(_92_v20)
			var _161_v69 _dafny.Map
			_ = _161_v69
			_161_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), _160_v68)
			var _162_v70 _dafny.Map
			_ = _162_v70
			_162_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v7, (_161_v69).Cardinality())
			var _163_v71 _dafny.Sequence
			_ = _163_v71
			_163_v71 = _dafny.SeqOf(func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter13 := _dafny.Iterate((_159_v67).Elements()); ; {
					_compr_12, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _164_v66 _dafny.CodePoint
					_164_v66 = interface{}(_compr_12).(_dafny.CodePoint)
					if (_159_v67).Contains(_164_v66) {
						_coll12.Add(_164_v66, _68_v0)
					}
				}
				return _coll12.ToMap()
			}(), _162_v70)
			var _165_v72 *C2
			_ = _165_v72
			var _nw19 *C2 = New_C2_()
			_ = _nw19
			_nw19.Ctor__((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), _dafny.IntOfUint32(((_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)).Cardinality()), _163_v71)
			_165_v72 = _nw19
			var _166_v73 _dafny.Sequence
			_ = _166_v73
			_166_v73 = _dafny.SeqOf(_165_v72)
			var _167_v74 _dafny.MultiSet
			_ = _167_v74
			_167_v74 = _dafny.MultiSetOf((_166_v73).Select((Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_166_v73).Cardinality()))).Uint32()).(*C2), _165_v72)
			var _168_v75 _dafny.Map
			_ = _168_v75
			_168_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _165_v72)
			var _169_v76 T0
			_ = _169_v76
			var _nw20 *C4 = New_C4_()
			_ = _nw20
			_nw20.Ctor__(_71_v3, (func() _dafny.Int {
				if (_167_v74).Contains((func() *C2 {
					if (_168_v75).Contains((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)) {
						return (_168_v75).Get((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)).(*C2)
					}
					return _165_v72
				})()) {
					return (_167_v74).Multiplicity((func() *C2 {
						if (_168_v75).Contains((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)) {
							return (_168_v75).Get((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)).(*C2)
						}
						return _165_v72
					})())
				}
				return _68_v0
			})(), _163_v71)
			_169_v76 = _nw20
			_169_v76 = _169_v76
			_69_v1 = _69_v1
			var _170_v77 *C2
			_ = _170_v77
			var _nw21 *C2 = New_C2_()
			_ = _nw21
			_nw21.Ctor__((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_92_v20).Select((Companion_Default___.SafeIndex((_169_v76).F24(), _dafny.IntOfUint32((_92_v20).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(687))), Companion_Default___.Fm42((_169_v76).F24(), _88_globalState))
			_170_v77 = _nw21
			(_88_globalState).F17 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yqdcd"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-23))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_171_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_172_i6 _dafny.Int) _dafny.CodePoint {
					return _171_v7
				}
			})(_76_v7))))
			var _173_v78 *C0
			_ = _173_v78
			var _nw22 *C0 = New_C0_()
			_ = _nw22
			_nw22.Ctor__((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool))
			_173_v78 = _nw22
			_173_v78 = _173_v78
		}
		(_88_globalState).F2 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_92_v20, _92_v20), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.Fm0(_68_v0, _68_v0, _88_globalState)), _68_v0))
		var _174_v79 _dafny.Array
		_ = _174_v79
		var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(25))
		_ = _nw23
		_174_v79 = _nw23
		var _175_v80 _dafny.Sequence
		_ = _175_v80
		_175_v80 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v7, _68_v0)).Update(_76_v7, _68_v0))
		var _176_v81 *C2
		_ = _176_v81
		var _nw24 *C2 = New_C2_()
		_ = _nw24
		_nw24.Ctor__(false, _68_v0, (_175_v80))
		_176_v81 = _nw24
		var _177_v82 _dafny.Map
		_ = _177_v82
		_177_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _dafny.IntOfInt64(83))
		var _178_v83 _dafny.Array
		_ = _178_v83
		var _nwElement0_2 _dafny.Map = _177_v82
		_ = _nwElement0_2
		var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(4))
		_ = _nw25
		(_nw25).ArraySet1(_nwElement0_2, 0)
		(_nw25).ArraySet1(_177_v82, 1)
		(_nw25).ArraySet1(_177_v82, 2)
		(_nw25).ArraySet1(_177_v82, 3)
		_178_v83 = _nw25
		var _179_v84 _dafny.Map
		_ = _179_v84
		_179_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v81, _178_v83)
		var _180_v85 *C3
		_ = _180_v85
		var _nw26 *C3 = New_C3_()
		_ = _nw26
		_nw26.Ctor__(_174_v79, (func() _dafny.Array {
			if (_179_v84).Contains(_176_v81) {
				return (_179_v84).Get(_176_v81).(_dafny.Array)
			}
			return _178_v83
		})(), _dafny.IntOfInt64(869), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(510))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_181_v7 _dafny.CodePoint, _182_v0 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_183_i7 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v7, _182_v0)
			}
		})(_76_v7, _68_v0))))
		_180_v85 = _nw26
		(_88_globalState).F18 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_85_v14).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_180_v85)).Cardinality()))
		var _184_v87 _dafny.Sequence
		_ = _184_v87
		_184_v87 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v7, _68_v0))
		var _185_v88 _dafny.Map
		_ = _185_v88
		_185_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v7, _68_v0)
		var _186_v89 *C4
		_ = _186_v89
		var _nw27 *C4 = New_C4_()
		_ = _nw27
		_nw27.Ctor__(_144_cf0, _68_v0, _dafny.Companion_Sequence_.Update(_184_v87, (Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_184_v87).Cardinality()))).Uint32(), _185_v88))
		_186_v89 = _nw27
		var _187_v90 _dafny.Map
		_ = _187_v90
		_187_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)), _76_v7)).Update(_176_v81.F31, _76_v7), _186_v89)
		_81_v12 = (_81_v12).Update(((func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(62), _dafny.IntOfInt64(277))); ; {
				_compr_13, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _188_v86 _dafny.Int
				_188_v86 = interface{}(_compr_13).(_dafny.Int)
				if ((_dafny.IntOfInt64(62)).Cmp(_188_v86) <= 0) && ((_188_v86).Cmp(_dafny.IntOfInt64(277)) < 0) {
					_coll13.Add(Companion_Default___.SafeModuloInt(_188_v86, _68_v0))
				}
			}
			return _coll13.ToSet()
		}()).Cardinality()).Times((_187_v90).Cardinality()), _80_v11)
	}
	var _hi2 _dafny.Int = (_dafny.Zero).Minus(_68_v0)
	_ = _hi2
	for _189_i8 := Companion_Default___.Fm0(_68_v0, _68_v0, _88_globalState); _189_i8.Cmp(_hi2) < 0; _189_i8 = _189_i8.Plus(_dafny.One) {
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
		_ = _index21
		(_80_v11).ArraySet1((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), (_index21).Int())
		(_88_globalState).F18 = _189_i8
		var _190_v91 _dafny.MultiSet
		_ = _190_v91
		_190_v91 = _dafny.MultiSetOf(_189_i8)
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
		_ = _index22
		(_80_v11).ArraySet1(((func() bool {
			if (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool) {
				return (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)
			}
			return _70_v2
		})()) || (!(_dafny.MultiSetOf(_189_i8)).Equals(_190_v91)), (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_71_v3), 0))
		_ = _index23
		(_71_v3).ArraySet1(_189_i8, (_index23).Int())
		var _191_v92 _dafny.Array
		_ = _191_v92
		var _nwElement0_3 _dafny.CodePoint = _dafny.CodePoint('y')
		_ = _nwElement0_3
		var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(27))
		_ = _nw28
		(_nw28).ArraySet1CodePoint(_nwElement0_3, 0)
		(_nw28).ArraySet1CodePoint(_76_v7, 1)
		(_nw28).ArraySet1CodePoint(_76_v7, 2)
		(_nw28).ArraySet1CodePoint(_76_v7, 3)
		(_nw28).ArraySet1CodePoint(_76_v7, 4)
		(_nw28).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool) {
				return _76_v7
			}
			return _76_v7
		})(), 5)
		(_nw28).ArraySet1CodePoint(_76_v7, 6)
		(_nw28).ArraySet1CodePoint(_76_v7, 7)
		(_nw28).ArraySet1CodePoint(_dafny.CodePoint('u'), 8)
		(_nw28).ArraySet1CodePoint(_76_v7, 9)
		(_nw28).ArraySet1CodePoint(((_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_189_i8, _dafny.IntOfUint32(((_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint), 10)
		(_nw28).ArraySet1CodePoint(_dafny.CodePoint('l'), 11)
		(_nw28).ArraySet1CodePoint(_76_v7, 12)
		(_nw28).ArraySet1CodePoint(_76_v7, 13)
		(_nw28).ArraySet1CodePoint(_76_v7, 14)
		(_nw28).ArraySet1CodePoint(_76_v7, 15)
		(_nw28).ArraySet1CodePoint(_76_v7, 16)
		(_nw28).ArraySet1CodePoint(_76_v7, 17)
		(_nw28).ArraySet1CodePoint(_76_v7, 18)
		(_nw28).ArraySet1CodePoint(Companion_Default___.Fm3(_88_globalState), 19)
		(_nw28).ArraySet1CodePoint(_76_v7, 20)
		(_nw28).ArraySet1CodePoint(_76_v7, 21)
		(_nw28).ArraySet1CodePoint(_76_v7, 22)
		(_nw28).ArraySet1CodePoint(_76_v7, 23)
		(_nw28).ArraySet1CodePoint(_76_v7, 24)
		(_nw28).ArraySet1CodePoint(_76_v7, 25)
		(_nw28).ArraySet1CodePoint(_76_v7, 26)
		_191_v92 = _nw28
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_191_v92), 0))
		_ = _index24
		(_191_v92).ArraySet1CodePoint(_76_v7, (_index24).Int())
		var _192_v93 _dafny.Map
		_ = _192_v93
		_192_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v2, _76_v7)
		var _193_v94 _dafny.Sequence
		_ = _193_v94
		_193_v94 = _dafny.SeqOf(_80_v11)
		var _194_v95 _dafny.Map
		_ = _194_v95
		_194_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v11, _80_v11)
		var _195_v96 *C0
		_ = _195_v96
		var _nw29 *C0 = New_C0_()
		_ = _nw29
		_nw29.Ctor__(_70_v2)
		_195_v96 = _nw29
		var _196_v97 _dafny.Sequence
		_ = _196_v97
		_196_v97 = _dafny.SeqOf(_195_v96)
		var _197_v98 _dafny.MultiSet
		_ = _197_v98
		_197_v98 = _dafny.MultiSetOf((_193_v94).Select((Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_193_v94).Cardinality()))).Uint32()).(_dafny.Array), _80_v11, (func() _dafny.Array {
			if (_194_v95).Contains(_80_v11) {
				return (_194_v95).Get(_80_v11).(_dafny.Array)
			}
			return (_193_v94).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_196_v97).Cardinality()), _dafny.IntOfUint32((_193_v94).Cardinality()))).Uint32()).(_dafny.Array)
		})())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_71_v3), 0))
		_ = _index25
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_191_v92), 0))
		_ = _index26
		var _rhs14 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_75_v6, _75_v6), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_192_v93).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm23(_88_globalState)).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_75_v6, _75_v6)).Cardinality()))).Uint32(), (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool))
		_ = _rhs14
		var _rhs15 _dafny.Int = Companion_Default___.Fm0(_68_v0, Companion_Default___.SafeDivisionInt((_192_v93).Cardinality(), _68_v0), _88_globalState)
		_ = _rhs15
		var _rhs16 _dafny.Int = _dafny.IntOfInt64(525)
		_ = _rhs16
		var _rhs17 _dafny.CodePoint = _76_v7
		_ = _rhs17
		var _rhs18 _dafny.Int = (func() _dafny.Int {
			if (_197_v98).Contains(_80_v11) {
				return (_197_v98).Multiplicity(_80_v11)
			}
			return _189_i8
		})()
		_ = _rhs18
		var _lhs13 *GlobalState = _88_globalState
		_ = _lhs13
		var _lhs14 _dafny.Array = _71_v3
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_71_v3), 0))
		_ = _lhs15
		var _lhs16 _dafny.Array = _191_v92
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_191_v92), 0))
		_ = _lhs17
		var _lhs18 *GlobalState = _88_globalState
		_ = _lhs18
		_lhs13.F20 = _rhs14
		(_lhs14).ArraySet1(_rhs15, (_lhs15).Int())
		_68_v0 = _rhs16
		(_lhs16).ArraySet1CodePoint(_rhs17, (_lhs17).Int())
		_lhs18.F9 = _rhs18
	}
	var _198_v99 _dafny.Array
	_ = _198_v99
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(6)
	_ = _len0_3
	var _nw30 _dafny.Array
	_ = _nw30
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw30 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Set = (func(_199_v0 _dafny.Int) func(_dafny.Int) _dafny.Set {
			return func(_200_i9 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_199_v0, _199_v0)
			}
		})(_68_v0)
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
	_198_v99 = _nw30
	var _201_v100 _dafny.Array
	_ = _201_v100
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(25)
	_ = _len0_4
	var _nw31 _dafny.Array
	_ = _nw31
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw31 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Map = (func(_202_v6 _dafny.Sequence, _203_v0 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_204_i10 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_202_v6).Cardinality()), _203_v0)
			}
		})(_75_v6, _68_v0)
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
	_201_v100 = _nw31
	var _205_v102 _dafny.Map
	_ = _205_v102
	_205_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v7, _68_v0)
	var _206_v103 _dafny.Sequence
	_ = _206_v103
	_206_v103 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v7, _dafny.IntOfInt64(-1)), _205_v102)
	var _207_v104 *C3
	_ = _207_v104
	var _nw32 *C3 = New_C3_()
	_ = _nw32
	_nw32.Ctor__(_198_v99, _201_v100, Companion_Default___.SafeModuloInt(_68_v0, _68_v0), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(775))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}((func(_208_v2 bool, _209_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
		return func(_210_i11 _dafny.Int) _dafny.Map {
			return func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter15 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _208_v2)).Keys().Elements()); ; {
					_compr_14, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _211_v101 _dafny.CodePoint
					_211_v101 = interface{}(_compr_14).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _208_v2)).Contains(_211_v101) {
						_coll14.Add(_211_v101, _dafny.IntOfUint32((_209_v6).Cardinality()))
					}
				}
				return _coll14.ToMap()
			}()
		}
	})(_70_v2, _75_v6))), _206_v103))
	_207_v104 = _nw32
	var _212_v105 _dafny.Map
	_ = _212_v105
	_212_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v0, _70_v2)
	var _213_v106 _dafny.Map
	_ = _213_v106
	_213_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
		if false {
			return (func() bool {
				if (_212_v105).Contains(_dafny.IntOfUint32((_75_v6).Cardinality())) {
					return (_212_v105).Get(_dafny.IntOfUint32((_75_v6).Cardinality())).(bool)
				}
				return true
			})()
		}
		return (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)
	})(), Companion_Default___.Fm43(_68_v0, !(Companion_Default___.Fm2(_68_v0, _68_v0, _76_v7, _88_globalState)), _70_v2, _88_globalState))
	_213_v106 = (_213_v106).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), _86_v15))
	var _214_v107 D3
	_ = _214_v107
	_214_v107 = Companion_D3_.Create_DC9_(_75_v6)
	var _source1 D3 = _214_v107
	_ = _source1
	if _source1.Is_DC10() {
		(_88_globalState).F2 = !(_70_v2)
		var _215_v108 bool
		_ = _215_v108
		var _216_v109 bool
		_ = _216_v109
		var _217_v110 _dafny.Int
		_ = _217_v110
		var _out2 bool
		_ = _out2
		var _out3 bool
		_ = _out3
		var _out4 _dafny.Int
		_ = _out4
		_out2, _out3, _out4 = (_207_v104).M5(_70_v2, _88_globalState)
		_215_v108 = _out2
		_216_v109 = _out3
		_217_v110 = _out4
		var _218_v111 _dafny.Array
		_ = _218_v111
		var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
		_ = _nw33
		_218_v111 = _nw33
		_218_v111 = _218_v111
		(_88_globalState).F1 = Companion_Default___.Fm0(_68_v0, (_68_v0).Times(_217_v110), _88_globalState)
	} else if _source1.Is_DC11() {
		var _219___mcc_h7 _dafny.Int = _source1.Get_().(D3_DC11).Cf18
		_ = _219___mcc_h7
		var _220___mcc_h8 bool = _source1.Get_().(D3_DC11).Cf19
		_ = _220___mcc_h8
		var _221___mcc_h9 bool = _source1.Get_().(D3_DC11).Cf20
		_ = _221___mcc_h9
		var _222_cf20 bool = _221___mcc_h9
		_ = _222_cf20
		var _223_cf19 bool = _220___mcc_h8
		_ = _223_cf19
		var _224_cf18 _dafny.Int = _219___mcc_h7
		_ = _224_cf18
		var _225_v112 _dafny.Set
		_ = _225_v112
		_225_v112 = _dafny.SetOf(_75_v6, _75_v6, _75_v6, _75_v6)
		var _226_v113 _dafny.Map
		_ = _226_v113
		_226_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_cf19, _225_v112)
		var _227_v114 _dafny.Set
		_ = _227_v114
		_227_v114 = (func() _dafny.Set {
			if (_226_v113).Contains(_70_v2) {
				return (_226_v113).Get(_70_v2).(_dafny.Set)
			}
			return _dafny.SetOf(_75_v6, _75_v6)
		})()
		var _source2 _dafny.Set = _227_v114
		_ = _source2
		var _228___mcc_h11 _dafny.Set = _source2
		_ = _228___mcc_h11
		var _229_cf65 _dafny.Set = _228___mcc_h11
		_ = _229_cf65
		(_88_globalState).F17 = (_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)
		(_207_v104).M6(_88_globalState)
		var _230_v115 *C4
		_ = _230_v115
		var _nw34 *C4 = New_C4_()
		_ = _nw34
		_nw34.Ctor__(_71_v3, _68_v0, _206_v103)
		_230_v115 = _nw34
		var _231_v116 _dafny.Map
		_ = _231_v116
		_231_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v20, _223_cf19)
		var _232_v117 _dafny.Sequence
		_ = _232_v117
		_232_v117 = _dafny.SeqOf(_69_v1)
		var _233_v118 _dafny.Int
		_ = _233_v118
		var _234_v119 _dafny.Int
		_ = _234_v119
		var _235_v120 _dafny.Sequence
		_ = _235_v120
		var _236_v121 bool
		_ = _236_v121
		var _out5 _dafny.Int
		_ = _out5
		var _out6 _dafny.Int
		_ = _out6
		var _out7 _dafny.Sequence
		_ = _out7
		var _out8 bool
		_ = _out8
		_out5, _out6, _out7, _out8 = (_230_v115).M4((func() bool {
			if !(!(_223_cf19)) {
				return (func() bool {
					if (_231_v116).Contains(_92_v20) {
						return (_231_v116).Get(_92_v20).(bool)
					}
					return !(true)
				})()
			}
			return (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)
		})(), !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_232_v117, (Companion_Default___.SafeIndex(_68_v0, _dafny.IntOfUint32((_232_v117).Cardinality()))).Uint32(), _69_v1), _232_v117)), _88_globalState)
		_233_v118 = _out5
		_234_v119 = _out6
		_235_v120 = _out7
		_236_v121 = _out8
		(_207_v104).M6(_88_globalState)
		if !(_223_cf19) {
			var _rhs19 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_85_v14, _85_v14)
			_ = _rhs19
			var _rhs20 _dafny.Int = _68_v0
			_ = _rhs20
			var _lhs19 *GlobalState = _88_globalState
			_ = _lhs19
			var _lhs20 *GlobalState = _88_globalState
			_ = _lhs20
			_lhs19.F17 = _rhs19
			_lhs20.F14 = _rhs20
			(_88_globalState).F14 = (_68_v0).Times(((_92_v20).Select((Companion_Default___.SafeIndex(_224_cf18, _dafny.IntOfUint32((_92_v20).Cardinality()))).Uint32()).(_dafny.Int)).Plus((_dafny.SetOf(_80_v11)).Cardinality()))
			var _237_v122 _dafny.Int
			_ = _237_v122
			var _out9 _dafny.Int
			_ = _out9
			_out9 = Companion_Default___.M0(_dafny.IntOfInt64(928), (_224_cf18).Times(_dafny.IntOfUint32((_92_v20).Cardinality())), _88_globalState)
			_237_v122 = _out9
			(_88_globalState).F21 = Companion_Default___.Fm23(_88_globalState)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
			_ = _index27
			(_80_v11).ArraySet1(_223_cf19, (_index27).Int())
		} else {
			var _238_v123 _dafny.Array
			_ = _238_v123
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(12))
			_ = _nw35
			_238_v123 = _nw35
			var _239_v124 _dafny.Map
			_ = _239_v124
			_239_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v2, _238_v123)
			_239_v124 = (_239_v124).Update(_70_v2, _238_v123)
			var _240_v125 _dafny.Array
			_ = _240_v125
			var _nw36 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
			_ = _nw36
			_240_v125 = _nw36
			_240_v125 = _240_v125
			(_88_globalState).F14 = (_dafny.IntOfInt64(369)).Minus(_224_cf18)
			(_88_globalState).F2 = _222_cf20
			(_88_globalState).F18 = _68_v0
		}
		(_88_globalState).F9 = _68_v0
	} else {
		var _241___mcc_h10 _dafny.Sequence = _source1.Get_().(D3_DC9).Cf17
		_ = _241___mcc_h10
		var _242_cf17 _dafny.Sequence = _241___mcc_h10
		_ = _242_cf17
		var _243_v126 *C4
		_ = _243_v126
		var _nw37 *C4 = New_C4_()
		_ = _nw37
		_nw37.Ctor__(_71_v3, _dafny.IntOfUint32(((_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)).Cardinality()), _206_v103)
		_243_v126 = _nw37
		var _244_v127 D3
		_ = _244_v127
		_244_v127 = Companion_D3_.Create_DC10_()
		var _245_v128 _dafny.Sequence
		_ = _245_v128
		_245_v128 = _dafny.SeqOf((_dafny.MultiSetOf(Companion_Default___.Fm0(_68_v0, _68_v0, _88_globalState), _dafny.IntOfUint32(((_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Update((_212_v105).Cardinality(), Companion_Default___.Abs(_68_v0)))
		var _246_v129 *C1
		_ = _246_v129
		var _nw38 *C1 = New_C1_()
		_ = _nw38
		_nw38.Ctor__(_dafny.Companion_Sequence_.Update(_245_v128, (Companion_Default___.SafeIndex((_207_v104).Fm8(_88_globalState), _dafny.IntOfUint32((_245_v128).Cardinality()))).Uint32(), _dafny.MultiSetOf(_dafny.IntOfInt64(-838), _68_v0, _68_v0)), _68_v0, (_69_v1).Cardinality(), _206_v103)
		_246_v129 = _nw38
		var _247_v130 _dafny.Map
		_ = _247_v130
		_247_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), _243_v126)
		var _248_v131 _dafny.MultiSet
		_ = _248_v131
		_248_v131 = _dafny.MultiSetOf(_dafny.SeqOf(_246_v129.F33, _dafny.IntOfUint32((_242_cf17).Cardinality()), (_dafny.Zero).Minus(_68_v0)), _92_v20, _dafny.SeqOf(_dafny.IntOfInt64(368)))
		var _rhs21 bool = (Companion_D3_.Create_DC11_(_246_v129.F33, _70_v2, _70_v2)).Dtor_cf20()
		_ = _rhs21
		var _rhs22 *C4 = (func() *C4 {
			if (_247_v130).Contains(_70_v2) {
				return (_247_v130).Get(_70_v2).(*C4)
			}
			return _243_v126
		})()
		_ = _rhs22
		var _rhs23 D3 = Companion_Default___.Fm33(!(((_248_v131).Cardinality()).Cmp(_68_v0) != 0), _246_v129.F33, _88_globalState)
		_ = _rhs23
		var _rhs24 _dafny.Int = _68_v0
		_ = _rhs24
		var _rhs25 *C1 = _246_v129
		_ = _rhs25
		var _lhs21 *GlobalState = _88_globalState
		_ = _lhs21
		var _lhs22 *GlobalState = _88_globalState
		_ = _lhs22
		_lhs21.F2 = _rhs21
		_243_v126 = _rhs22
		_244_v127 = _rhs23
		_lhs22.F18 = _rhs24
		_246_v129 = _rhs25
		var _249_v132 _dafny.Array
		_ = _249_v132
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
		_ = _nw39
		_249_v132 = _nw39
		_249_v132 = _249_v132
		if _70_v2 {
			(_88_globalState).F9 = _dafny.IntOfInt64(-974)
			(_88_globalState).F18 = (_68_v0).Times(_dafny.IntOfInt64(-481))
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
			_ = _index28
			(_80_v11).ArraySet1((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), (_index28).Int())
			var _250_v133 *C0
			_ = _250_v133
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__(!(true))
			_250_v133 = _nw40
			var _251_v134 _dafny.Array
			_ = _251_v134
			var _nwElement0_4 *C0 = _250_v133
			_ = _nwElement0_4
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.One)
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_4, 0)
			_251_v134 = _nw41
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_251_v134), 0))
			_ = _index29
			(_251_v134).ArraySet1(_250_v133, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_251_v134), 0))
			_ = _index30
			(_251_v134).ArraySet1(_250_v133, (_index30).Int())
			var _252_v135 T0
			_ = _252_v135
			var _nw42 *C4 = New_C4_()
			_ = _nw42
			_nw42.Ctor__(_71_v3, Companion_Default___.Fm0(_246_v129.F33, _246_v129.F33, _88_globalState), _dafny.Companion_Sequence_.Concatenate(_206_v103, _dafny.SeqOf(_205_v102)))
			_252_v135 = _nw42
			var _rhs26 T0 = _252_v135
			_ = _rhs26
			var _rhs27 _dafny.Sequence = (_122_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_122_v45), 0))).Int()).(_dafny.Sequence)
			_ = _rhs27
			var _rhs28 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfInt64(495)).Minus(((_246_v129).Fm11(_70_v2, (_250_v133).F34(), (_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), _68_v0, _88_globalState)).Plus((_252_v135).F24())))
			_ = _rhs28
			var _lhs23 *GlobalState = _88_globalState
			_ = _lhs23
			var _lhs24 *GlobalState = _88_globalState
			_ = _lhs24
			_252_v135 = _rhs26
			_lhs23.F17 = _rhs27
			_lhs24.F9 = _rhs28
		} else {
			var _253_v136 bool
			_ = _253_v136
			var _254_v137 bool
			_ = _254_v137
			var _255_v138 _dafny.Int
			_ = _255_v138
			var _out10 bool
			_ = _out10
			var _out11 bool
			_ = _out11
			var _out12 _dafny.Int
			_ = _out12
			_out10, _out11, _out12 = (_207_v104).M5((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool), _88_globalState)
			_253_v136 = _out10
			_254_v137 = _out11
			_255_v138 = _out12
			(_88_globalState).F18 = _68_v0
			(_88_globalState).F18 = _246_v129.F33
			(_88_globalState).F14 = (_dafny.Zero).Minus(_246_v129.F33)
			(_88_globalState).F18 = Companion_Default___.Fm0(_68_v0, (_dafny.Zero).Minus((_246_v129.F33).Plus(_246_v129.F33)), _88_globalState)
		}
		var _256_v139 _dafny.MultiSet
		_ = _256_v139
		_256_v139 = _dafny.MultiSetOf(_246_v129.F33, _dafny.IntOfUint32((_85_v14).Cardinality()))
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))
		_ = _index31
		(_80_v11).ArraySet1(!((_256_v139).IsProperSubsetOf(_256_v139)) || ((_80_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_80_v11), 0))).Int()).(bool)), (_index31).Int())
	}
	(_88_globalState).F18 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(((_dafny.MultiSetOf(_68_v0, _dafny.IntOfUint32((_92_v20).Cardinality()), _68_v0)).Cardinality()).Plus((_dafny.Zero).Minus(_68_v0))), _68_v0)
	(_88_globalState).F18 = _68_v0
	_212_v105 = (_212_v105).Update(_68_v0, !((_126_v49).IsSubsetOf(_126_v49)))
	_dafny.Print(_68_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v1).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(590))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_70_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v3).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_73_v4).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v5).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_75_v6, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_76_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_77_v8).Dtor_cf7(), _dafny.SeqOf(_dafny.CodePoint('c'), _dafny.CodePoint('h'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_78_v9).Equals(_dafny.MultiSetOf(_dafny.CodePoint('c'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_79_v10).Equals(_dafny.SetOf(_dafny.MultiSetOf(_dafny.CodePoint('c'), _dafny.CodePoint('h')), _dafny.MultiSetOf(_dafny.CodePoint('c')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v11).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v12).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_84_v13).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_v14.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v15).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v15).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v15).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v15).Dtor_cf11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v15).Dtor_cf12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_87_v16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(590), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(382), _dafny.IntOfInt64(590)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_88_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_88_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_88_globalState).F3()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(590)).UpdateUnsafe(true, _dafny.IntOfInt64(590))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_88_globalState).F7()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_88_globalState).F8(), _dafny.SeqOf(false, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_88_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_88_globalState).F11()).Equals(_dafny.SetOf(_dafny.MultiSetOf(_dafny.CodePoint('c'), _dafny.CodePoint('h')), _dafny.MultiSetOf(_dafny.CodePoint('c')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_88_globalState).F13()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_88_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_88_globalState).F16()).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_88_globalState.F17.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_88_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState.F19).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_88_globalState.F20, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_88_globalState.F21, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState.F22).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(382), _dafny.IntOfInt64(590))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_88_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v18).Equals(_dafny.MultiSetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_92_v20, _dafny.SeqOf(_dafny.IntOfInt64(590), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v39).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v39).Dtor_cf14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v39).Dtor_cf15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v45).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_v46).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"), _dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"), _dafny.UnicodeSeqOfUtf8Bytes("ncd"), _dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"), _dafny.UnicodeSeqOfUtf8Bytes("nbntyekh")), _dafny.UnicodeSeqOfUtf8Bytes("ubmfhj"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_124_v47, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_125_v48, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"), _dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"), _dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"), _dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"), _dafny.UnicodeSeqOfUtf8Bytes("nbntyekh"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v49).Equals(_dafny.SetOf(_dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v50).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v50).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v50).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v99).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v99).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v99).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v99).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v99).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v99).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_201_v100).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v102).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_206_v103, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), _dafny.IntOfInt64(-1)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), _dafny.IntOfInt64(525)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_207_v104).F29()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_207_v104).F29()).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_207_v104).F29()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_207_v104).F29()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_207_v104).F29()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_207_v104).F29()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_207_v104.F30).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(525))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v105).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(525), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_213_v106).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D1_.Create_DC4_(true, true, _dafny.CodePoint('k'), _dafny.IntOfInt64(590), _dafny.IntOfInt64(590)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_214_v107).Dtor_cf17(), _dafny.SeqOf(true)))
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
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 bool
	Cf5 _dafny.Int
	Cf6 _dafny.CodePoint
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 bool, Cf5 _dafny.Int, Cf6 _dafny.CodePoint) D0 {
	return D0{D0_DC2{Cf4, Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Array
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Array) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() _dafny.CodePoint {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf0() _dafny.Array {
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
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4 == data2.Cf4 && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6 == data2.Cf6
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

type D1_DC4 struct {
	Cf8  bool
	Cf9  bool
	Cf10 _dafny.CodePoint
	Cf11 _dafny.Int
	Cf12 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 bool, Cf9 bool, Cf10 _dafny.CodePoint, Cf11 _dafny.Int, Cf12 _dafny.Int) D1 {
	return D1{D1_DC4{Cf8, Cf9, Cf10, Cf11, Cf12}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf13 bool
	Cf14 bool
	Cf15 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf13 bool, Cf14 bool, Cf15 bool) D1 {
	return D1{D1_DC5{Cf13, Cf14, Cf15}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_() D1 {
	return D1{D1_DC6{}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC3 struct {
	Cf7 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false, false, _dafny.CodePoint('D'), _dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf12
}

func (_this D1) Dtor_cf13() bool {
	return _this.Get_().(D1_DC5).Cf13
}

func (_this D1) Dtor_cf14() bool {
	return _this.Get_().(D1_DC5).Cf14
}

func (_this D1) Dtor_cf15() bool {
	return _this.Get_().(D1_DC5).Cf15
}

func (_this D1) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + data.Cf7.VerbatimString(true) + ")"
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
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15
		}
	case D1_DC6:
		{
			_, ok := other.Get_().(D1_DC6)
			return ok
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_() D2 {
	return D2{D2_DC8{}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf16 _dafny.Array
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf16 _dafny.Array) D2 {
	return D2{D2_DC7{Cf16}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_()
}

func (_this D2) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf16
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf16) + ")"
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
			_, ok := other.Get_().(D2_DC8)
			return ok
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf16 == data2.Cf16
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
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_() D3 {
	return D3{D3_DC10{}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf18 _dafny.Int
	Cf19 bool
	Cf20 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf18 _dafny.Int, Cf19 bool, Cf20 bool) D3 {
	return D3{D3_DC11{Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC9 struct {
	Cf17 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf17 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf17}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_()
}

func (_this D3) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf18
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC11).Cf20
}

func (_this D3) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf17
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf17) + ")"
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
			_, ok := other.Get_().(D3_DC10)
			return ok
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf17.Equals(data2.Cf17)
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
	Cf22 _dafny.Int
	Cf23 _dafny.Map
	Cf24 _dafny.Sequence
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf22 _dafny.Int, Cf23 _dafny.Map, Cf24 _dafny.Sequence) D4 {
	return D4{D4_DC13{Cf22, Cf23, Cf24}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC12 struct {
	Cf21 _dafny.Sequence
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf21 _dafny.Sequence) D4 {
	return D4{D4_DC12{Cf21}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC14 struct {
	Cf25 D4
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf25 D4) D4 {
	return D4{D4_DC14{Cf25}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(_dafny.Zero, _dafny.EmptyMap, _dafny.EmptySeq)
}

func (_this D4) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Map {
	return _this.Get_().(D4_DC13).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D4_DC13).Cf24
}

func (_this D4) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D4_DC12).Cf21
}

func (_this D4) Dtor_cf25() D4 {
	return _this.Get_().(D4_DC14).Cf25
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + data.Cf24.VerbatimString(true) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf25) + ")"
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
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23.Equals(data2.Cf23) && data1.Cf24.Equals(data2.Cf24)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf21.Equals(data2.Cf21)
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
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

type D5_DC15 struct {
	Cf26 _dafny.MultiSet
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf26 _dafny.MultiSet) D5 {
	return D5{D5_DC15{Cf26}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D5) Dtor_cf26() _dafny.MultiSet {
	return _this.Get_().(D5_DC15).Cf26
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf26.Equals(data2.Cf26)
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
	Cf28 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf28 bool) D6 {
	return D6{D6_DC17{Cf28}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf29 bool
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf29 bool) D6 {
	return D6{D6_DC18{Cf29}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC16 struct {
	Cf27 _dafny.Set
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf27 _dafny.Set) D6 {
	return D6{D6_DC16{Cf27}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC19 struct {
	Cf30 D6
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf30 D6) D6 {
	return D6{D6_DC19{Cf30}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(false)
}

func (_this D6) Dtor_cf28() bool {
	return _this.Get_().(D6_DC17).Cf28
}

func (_this D6) Dtor_cf29() bool {
	return _this.Get_().(D6_DC18).Cf29
}

func (_this D6) Dtor_cf27() _dafny.Set {
	return _this.Get_().(D6_DC16).Cf27
}

func (_this D6) Dtor_cf30() D6 {
	return _this.Get_().(D6_DC19).Cf30
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf30) + ")"
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
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf29 == data2.Cf29
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf27.Equals(data2.Cf27)
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf30.Equals(data2.Cf30)
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
	Cf32 D1
	Cf33 D6
	Cf34 _dafny.Array
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf32 D1, Cf33 D6, Cf34 _dafny.Array) D7 {
	return D7{D7_DC21{Cf32, Cf33, Cf34}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC20 struct {
	Cf31 _dafny.Map
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf31 _dafny.Map) D7 {
	return D7{D7_DC20{Cf31}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC22 struct {
	Cf35 D7
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf35 D7) D7 {
	return D7{D7_DC22{Cf35}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC21_(Companion_D1_.Default(), Companion_D6_.Default(), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D7) Dtor_cf32() D1 {
	return _this.Get_().(D7_DC21).Cf32
}

func (_this D7) Dtor_cf33() D6 {
	return _this.Get_().(D7_DC21).Cf33
}

func (_this D7) Dtor_cf34() _dafny.Array {
	return _this.Get_().(D7_DC21).Cf34
}

func (_this D7) Dtor_cf31() _dafny.Map {
	return _this.Get_().(D7_DC20).Cf31
}

func (_this D7) Dtor_cf35() D7 {
	return _this.Get_().(D7_DC22).Cf35
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf31) + ")"
		}
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf35) + ")"
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
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf32.Equals(data2.Cf32) && data1.Cf33.Equals(data2.Cf33) && data1.Cf34 == data2.Cf34
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf31.Equals(data2.Cf31)
		}
	case D7_DC22:
		{
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf35.Equals(data2.Cf35)
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

type D8_DC24 struct {
	Cf37 _dafny.Int
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf37 _dafny.Int) D8 {
	return D8{D8_DC24{Cf37}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC25 struct {
	Cf38 _dafny.Int
	Cf39 bool
	Cf40 bool
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf38 _dafny.Int, Cf39 bool, Cf40 bool) D8 {
	return D8{D8_DC25{Cf38, Cf39, Cf40}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC26 struct {
	Cf41 D1
	Cf42 _dafny.Sequence
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf41 D1, Cf42 _dafny.Sequence) D8 {
	return D8{D8_DC26{Cf41, Cf42}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

type D8_DC23 struct {
	Cf36 _dafny.Set
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf36 _dafny.Set) D8 {
	return D8{D8_DC23{Cf36}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC24_(_dafny.Zero)
}

func (_this D8) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf37
}

func (_this D8) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D8_DC25).Cf38
}

func (_this D8) Dtor_cf39() bool {
	return _this.Get_().(D8_DC25).Cf39
}

func (_this D8) Dtor_cf40() bool {
	return _this.Get_().(D8_DC25).Cf40
}

func (_this D8) Dtor_cf41() D1 {
	return _this.Get_().(D8_DC26).Cf41
}

func (_this D8) Dtor_cf42() _dafny.Sequence {
	return _this.Get_().(D8_DC26).Cf42
}

func (_this D8) Dtor_cf36() _dafny.Set {
	return _this.Get_().(D8_DC23).Cf36
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf41) + ", " + data.Cf42.VerbatimString(true) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf36) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40
		}
	case D8_DC26:
		{
			data2, ok := other.Get_().(D8_DC26)
			return ok && data1.Cf41.Equals(data2.Cf41) && data1.Cf42.Equals(data2.Cf42)
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf36.Equals(data2.Cf36)
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

type D9_DC27 struct {
	Cf43 T0
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf43 T0) D9 {
	return D9{D9_DC27{Cf43}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

func (CompanionStruct_D9_) Default() T0 {
	return (T0)(nil)
}

func (_this D9) Dtor_cf43() T0 {
	return _this.Get_().(D9_DC27).Cf43
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && _dafny.AreEqual(data1.Cf43, data2.Cf43)
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

type D10_DC29 struct {
	Cf45 _dafny.Int
	Cf46 bool
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf45 _dafny.Int, Cf46 bool) D10 {
	return D10{D10_DC29{Cf45, Cf46}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC28 struct {
	Cf44 _dafny.Array
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf44 _dafny.Array) D10 {
	return D10{D10_DC28{Cf44}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_(_dafny.Zero, false)
}

func (_this D10) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf45
}

func (_this D10) Dtor_cf46() bool {
	return _this.Get_().(D10_DC29).Cf46
}

func (_this D10) Dtor_cf44() _dafny.Array {
	return _this.Get_().(D10_DC28).Cf44
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46 == data2.Cf46
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf44 == data2.Cf44
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

type D11_DC31 struct {
	Cf48 bool
	Cf49 _dafny.Int
	Cf50 _dafny.Array
	Cf51 bool
	Cf52 _dafny.Int
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf48 bool, Cf49 _dafny.Int, Cf50 _dafny.Array, Cf51 bool, Cf52 _dafny.Int) D11 {
	return D11{D11_DC31{Cf48, Cf49, Cf50, Cf51, Cf52}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC30 struct {
	Cf47 _dafny.Map
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf47 _dafny.Map) D11 {
	return D11{D11_DC30{Cf47}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC31_(false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.Zero)
}

func (_this D11) Dtor_cf48() bool {
	return _this.Get_().(D11_DC31).Cf48
}

func (_this D11) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf49
}

func (_this D11) Dtor_cf50() _dafny.Array {
	return _this.Get_().(D11_DC31).Cf50
}

func (_this D11) Dtor_cf51() bool {
	return _this.Get_().(D11_DC31).Cf51
}

func (_this D11) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf52
}

func (_this D11) Dtor_cf47() _dafny.Map {
	return _this.Get_().(D11_DC30).Cf47
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf47) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf48 == data2.Cf48 && data1.Cf49.Cmp(data2.Cf49) == 0 && data1.Cf50 == data2.Cf50 && data1.Cf51 == data2.Cf51 && data1.Cf52.Cmp(data2.Cf52) == 0
		}
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf47.Equals(data2.Cf47)
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

type D12_DC33 struct {
	Cf54 _dafny.CodePoint
	Cf55 bool
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf54 _dafny.CodePoint, Cf55 bool) D12 {
	return D12{D12_DC33{Cf54, Cf55}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC32 struct {
	Cf53 _dafny.Map
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf53 _dafny.Map) D12 {
	return D12{D12_DC32{Cf53}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC34 struct {
	Cf56 D12
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf56 D12) D12 {
	return D12{D12_DC34{Cf56}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC33_(_dafny.CodePoint('D'), false)
}

func (_this D12) Dtor_cf54() _dafny.CodePoint {
	return _this.Get_().(D12_DC33).Cf54
}

func (_this D12) Dtor_cf55() bool {
	return _this.Get_().(D12_DC33).Cf55
}

func (_this D12) Dtor_cf53() _dafny.Map {
	return _this.Get_().(D12_DC32).Cf53
}

func (_this D12) Dtor_cf56() D12 {
	return _this.Get_().(D12_DC34).Cf56
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf53) + ")"
		}
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf54 == data2.Cf54 && data1.Cf55 == data2.Cf55
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf53.Equals(data2.Cf53)
		}
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf56.Equals(data2.Cf56)
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

type D13_DC36 struct {
	Cf58 bool
	Cf59 bool
}

func (D13_DC36) isD13() {}

func (CompanionStruct_D13_) Create_DC36_(Cf58 bool, Cf59 bool) D13 {
	return D13{D13_DC36{Cf58, Cf59}}
}

func (_this D13) Is_DC36() bool {
	_, ok := _this.Get_().(D13_DC36)
	return ok
}

type D13_DC35 struct {
	Cf57 _dafny.MultiSet
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf57 _dafny.MultiSet) D13 {
	return D13{D13_DC35{Cf57}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

type D13_DC37 struct {
	Cf60 D13
}

func (D13_DC37) isD13() {}

func (CompanionStruct_D13_) Create_DC37_(Cf60 D13) D13 {
	return D13{D13_DC37{Cf60}}
}

func (_this D13) Is_DC37() bool {
	_, ok := _this.Get_().(D13_DC37)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC36_(false, false)
}

func (_this D13) Dtor_cf58() bool {
	return _this.Get_().(D13_DC36).Cf58
}

func (_this D13) Dtor_cf59() bool {
	return _this.Get_().(D13_DC36).Cf59
}

func (_this D13) Dtor_cf57() _dafny.MultiSet {
	return _this.Get_().(D13_DC35).Cf57
}

func (_this D13) Dtor_cf60() D13 {
	return _this.Get_().(D13_DC37).Cf60
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC36:
		{
			return "D13.DC36" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf57) + ")"
		}
	case D13_DC37:
		{
			return "D13.DC37" + "(" + _dafny.String(data.Cf60) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC36:
		{
			data2, ok := other.Get_().(D13_DC36)
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59 == data2.Cf59
		}
	case D13_DC35:
		{
			data2, ok := other.Get_().(D13_DC35)
			return ok && data1.Cf57.Equals(data2.Cf57)
		}
	case D13_DC37:
		{
			data2, ok := other.Get_().(D13_DC37)
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

type D14_DC38 struct {
	Cf61 _dafny.Sequence
}

func (D14_DC38) isD14() {}

func (CompanionStruct_D14_) Create_DC38_(Cf61 _dafny.Sequence) D14 {
	return D14{D14_DC38{Cf61}}
}

func (_this D14) Is_DC38() bool {
	_, ok := _this.Get_().(D14_DC38)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D14) Dtor_cf61() _dafny.Sequence {
	return _this.Get_().(D14_DC38).Cf61
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC38:
		{
			return "D14.DC38" + "(" + _dafny.String(data.Cf61) + ")"
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
			return ok && data1.Cf61.Equals(data2.Cf61)
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
	Cf62 _dafny.Map
}

func (D15_DC39) isD15() {}

func (CompanionStruct_D15_) Create_DC39_(Cf62 _dafny.Map) D15 {
	return D15{D15_DC39{Cf62}}
}

func (_this D15) Is_DC39() bool {
	_, ok := _this.Get_().(D15_DC39)
	return ok
}

func (CompanionStruct_D15_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D15) Dtor_cf62() _dafny.Map {
	return _this.Get_().(D15_DC39).Cf62
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC39:
		{
			return "D15.DC39" + "(" + _dafny.String(data.Cf62) + ")"
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
			return ok && data1.Cf62.Equals(data2.Cf62)
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
	Cf64 _dafny.Int
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf64 _dafny.Int) D16 {
	return D16{D16_DC41{Cf64}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

type D16_DC40 struct {
	Cf63 _dafny.Array
}

func (D16_DC40) isD16() {}

func (CompanionStruct_D16_) Create_DC40_(Cf63 _dafny.Array) D16 {
	return D16{D16_DC40{Cf63}}
}

func (_this D16) Is_DC40() bool {
	_, ok := _this.Get_().(D16_DC40)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC41_(_dafny.Zero)
}

func (_this D16) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D16_DC41).Cf64
}

func (_this D16) Dtor_cf63() _dafny.Array {
	return _this.Get_().(D16_DC40).Cf63
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf64) + ")"
		}
	case D16_DC40:
		{
			return "D16.DC40" + "(" + _dafny.String(data.Cf63) + ")"
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
			return ok && data1.Cf64.Cmp(data2.Cf64) == 0
		}
	case D16_DC40:
		{
			data2, ok := other.Get_().(D16_DC40)
			return ok && data1.Cf63 == data2.Cf63
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

type D17_DC42 struct {
	Cf65 _dafny.Set
}

func (D17_DC42) isD17() {}

func (CompanionStruct_D17_) Create_DC42_(Cf65 _dafny.Set) D17 {
	return D17{D17_DC42{Cf65}}
}

func (_this D17) Is_DC42() bool {
	_, ok := _this.Get_().(D17_DC42)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D17) Dtor_cf65() _dafny.Set {
	return _this.Get_().(D17_DC42).Cf65
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC42:
		{
			return "D17.DC42" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC42:
		{
			data2, ok := other.Get_().(D17_DC42)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

type D18_DC43 struct {
	Cf66 *C1
}

func (D18_DC43) isD18() {}

func (CompanionStruct_D18_) Create_DC43_(Cf66 *C1) D18 {
	return D18{D18_DC43{Cf66}}
}

func (_this D18) Is_DC43() bool {
	_, ok := _this.Get_().(D18_DC43)
	return ok
}

func (CompanionStruct_D18_) Default() *C1 {
	return (*C1)(nil)
}

func (_this D18) Dtor_cf66() *C1 {
	return _this.Get_().(D18_DC43).Cf66
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC43:
		{
			return "D18.DC43" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC43:
		{
			data2, ok := other.Get_().(D18_DC43)
			return ok && data1.Cf66 == data2.Cf66
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

type D19_DC45 struct {
	Cf68 _dafny.Int
	Cf69 bool
	Cf70 _dafny.Int
	Cf71 _dafny.Int
}

func (D19_DC45) isD19() {}

func (CompanionStruct_D19_) Create_DC45_(Cf68 _dafny.Int, Cf69 bool, Cf70 _dafny.Int, Cf71 _dafny.Int) D19 {
	return D19{D19_DC45{Cf68, Cf69, Cf70, Cf71}}
}

func (_this D19) Is_DC45() bool {
	_, ok := _this.Get_().(D19_DC45)
	return ok
}

type D19_DC46 struct {
	Cf72 _dafny.Int
	Cf73 bool
}

func (D19_DC46) isD19() {}

func (CompanionStruct_D19_) Create_DC46_(Cf72 _dafny.Int, Cf73 bool) D19 {
	return D19{D19_DC46{Cf72, Cf73}}
}

func (_this D19) Is_DC46() bool {
	_, ok := _this.Get_().(D19_DC46)
	return ok
}

type D19_DC47 struct {
	Cf74 _dafny.Int
}

func (D19_DC47) isD19() {}

func (CompanionStruct_D19_) Create_DC47_(Cf74 _dafny.Int) D19 {
	return D19{D19_DC47{Cf74}}
}

func (_this D19) Is_DC47() bool {
	_, ok := _this.Get_().(D19_DC47)
	return ok
}

type D19_DC44 struct {
	Cf67 _dafny.Map
}

func (D19_DC44) isD19() {}

func (CompanionStruct_D19_) Create_DC44_(Cf67 _dafny.Map) D19 {
	return D19{D19_DC44{Cf67}}
}

func (_this D19) Is_DC44() bool {
	_, ok := _this.Get_().(D19_DC44)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC45_(_dafny.Zero, false, _dafny.Zero, _dafny.Zero)
}

func (_this D19) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D19_DC45).Cf68
}

func (_this D19) Dtor_cf69() bool {
	return _this.Get_().(D19_DC45).Cf69
}

func (_this D19) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D19_DC45).Cf70
}

func (_this D19) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D19_DC45).Cf71
}

func (_this D19) Dtor_cf72() _dafny.Int {
	return _this.Get_().(D19_DC46).Cf72
}

func (_this D19) Dtor_cf73() bool {
	return _this.Get_().(D19_DC46).Cf73
}

func (_this D19) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D19_DC47).Cf74
}

func (_this D19) Dtor_cf67() _dafny.Map {
	return _this.Get_().(D19_DC44).Cf67
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC45:
		{
			return "D19.DC45" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D19_DC46:
		{
			return "D19.DC46" + "(" + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ")"
		}
	case D19_DC47:
		{
			return "D19.DC47" + "(" + _dafny.String(data.Cf74) + ")"
		}
	case D19_DC44:
		{
			return "D19.DC44" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC45:
		{
			data2, ok := other.Get_().(D19_DC45)
			return ok && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69 == data2.Cf69 && data1.Cf70.Cmp(data2.Cf70) == 0 && data1.Cf71.Cmp(data2.Cf71) == 0
		}
	case D19_DC46:
		{
			data2, ok := other.Get_().(D19_DC46)
			return ok && data1.Cf72.Cmp(data2.Cf72) == 0 && data1.Cf73 == data2.Cf73
		}
	case D19_DC47:
		{
			data2, ok := other.Get_().(D19_DC47)
			return ok && data1.Cf74.Cmp(data2.Cf74) == 0
		}
	case D19_DC44:
		{
			data2, ok := other.Get_().(D19_DC44)
			return ok && data1.Cf67.Equals(data2.Cf67)
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

type D20_DC49 struct {
	Cf76 _dafny.MultiSet
	Cf77 _dafny.Int
}

func (D20_DC49) isD20() {}

func (CompanionStruct_D20_) Create_DC49_(Cf76 _dafny.MultiSet, Cf77 _dafny.Int) D20 {
	return D20{D20_DC49{Cf76, Cf77}}
}

func (_this D20) Is_DC49() bool {
	_, ok := _this.Get_().(D20_DC49)
	return ok
}

type D20_DC48 struct {
	Cf75 _dafny.Map
}

func (D20_DC48) isD20() {}

func (CompanionStruct_D20_) Create_DC48_(Cf75 _dafny.Map) D20 {
	return D20{D20_DC48{Cf75}}
}

func (_this D20) Is_DC48() bool {
	_, ok := _this.Get_().(D20_DC48)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC49_(_dafny.EmptyMultiSet, _dafny.Zero)
}

func (_this D20) Dtor_cf76() _dafny.MultiSet {
	return _this.Get_().(D20_DC49).Cf76
}

func (_this D20) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D20_DC49).Cf77
}

func (_this D20) Dtor_cf75() _dafny.Map {
	return _this.Get_().(D20_DC48).Cf75
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC49:
		{
			return "D20.DC49" + "(" + _dafny.String(data.Cf76) + ", " + _dafny.String(data.Cf77) + ")"
		}
	case D20_DC48:
		{
			return "D20.DC48" + "(" + _dafny.String(data.Cf75) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC49:
		{
			data2, ok := other.Get_().(D20_DC49)
			return ok && data1.Cf76.Equals(data2.Cf76) && data1.Cf77.Cmp(data2.Cf77) == 0
		}
	case D20_DC48:
		{
			data2, ok := other.Get_().(D20_DC48)
			return ok && data1.Cf75.Equals(data2.Cf75)
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

// Definition of trait T0
type T0 interface {
	String() string
	M1(globalState *GlobalState) bool
	F24() _dafny.Int
	F25() _dafny.Sequence
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
	F1   _dafny.Int
	F2   bool
	F9   _dafny.Int
	F14  _dafny.Int
	F17  _dafny.Sequence
	F18  _dafny.Int
	F19  _dafny.MultiSet
	F20  _dafny.Sequence
	F21  _dafny.Sequence
	F22  _dafny.Map
	_f0  bool
	_f3  _dafny.Map
	_f4  bool
	_f5  _dafny.Int
	_f6  _dafny.CodePoint
	_f7  _dafny.Map
	_f8  _dafny.Sequence
	_f10 _dafny.Int
	_f11 _dafny.Set
	_f12 _dafny.Int
	_f13 _dafny.Map
	_f15 _dafny.Int
	_f16 _dafny.Set
	_f23 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this.F2 = false
	_this.F9 = _dafny.Zero
	_this.F14 = _dafny.Zero
	_this.F17 = _dafny.EmptySeq
	_this.F18 = _dafny.Zero
	_this.F19 = _dafny.EmptyMultiSet
	_this.F20 = _dafny.EmptySeq
	_this.F21 = _dafny.EmptySeq
	_this.F22 = _dafny.EmptyMap
	_this._f0 = false
	_this._f3 = _dafny.EmptyMap
	_this._f4 = false
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.CodePoint('D')
	_this._f7 = _dafny.EmptyMap
	_this._f8 = _dafny.EmptySeq
	_this._f10 = _dafny.Zero
	_this._f11 = _dafny.EmptySet
	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.EmptyMap
	_this._f15 = _dafny.Zero
	_this._f16 = _dafny.EmptySet
	_this._f23 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 bool, f3 _dafny.Map, f4 bool, f5 _dafny.Int, f6 _dafny.CodePoint, f7 _dafny.Map, f8 _dafny.Sequence, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Set, f12 _dafny.Int, f13 _dafny.Map, f14 _dafny.Int, f15 _dafny.Int, f16 _dafny.Set, f17 _dafny.Sequence, f18 _dafny.Int, f19 _dafny.MultiSet, f20 _dafny.Sequence, f21 _dafny.Sequence, f22 _dafny.Map, f23 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this).F17 = f17
		(_this).F18 = f18
		(_this).F19 = f19
		(_this).F20 = f20
		(_this).F21 = f21
		(_this).F22 = f22
		(_this)._f23 = f23
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F3() _dafny.Map {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.CodePoint {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Map {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Set {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Map {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Set {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F23() _dafny.Int {
	{
		return _this._f23
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f34 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f34 = false
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

func (_this *C0) Ctor__(f34 bool) {
	{
		(_this)._f34 = f34
	}
}
func (_this *C0) Fm12(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf(((_dafny.Zero).Minus(_dafny.IntOfInt64(-780))).Plus(_dafny.IntOfInt64(-448)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("a")).Cardinality()))
	}
}
func (_this *C0) F34() bool {
	{
		return _this._f34
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f24 _dafny.Int
	_f25 _dafny.Sequence
	F33  _dafny.Int
	_f32 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f24 = _dafny.Zero
	_this._f25 = _dafny.EmptySeq
	_this.F33 = _dafny.Zero
	_this._f32 = _dafny.EmptySeq
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

func (_this *C1) F24() _dafny.Int {
	return _this._f24
}
func (_this *C1) F25() _dafny.Sequence {
	return _this._f25
}
func (_this *C1) Ctor__(f32 _dafny.Sequence, f33 _dafny.Int, f24 _dafny.Int, f25 _dafny.Sequence) {
	{
		(_this)._f32 = f32
		(_this).F33 = f33
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C1) Fm10(globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C1) Fm11(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qgrpl")).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('l'), _dafny.CodePoint('b')), _dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('y'), _dafny.CodePoint('p')))).Cardinality()))
	}
}
func (_this *C1) M1(globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		if ((_this.F33).Minus(_this.F33)).Cmp(_this.F33) <= 0 {
			var _257_v0 bool
			_ = _257_v0
			_257_v0 = true
			(globalState).F2 = _257_v0
			if _257_v0 {
				(globalState).F1 = (_this).F24()
				var _258_v1 _dafny.CodePoint
				_ = _258_v1
				_258_v1 = _dafny.CodePoint('i')
				_258_v1 = _258_v1
				var _259_v2 _dafny.Array
				_ = _259_v2
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_5
				var _nw43 _dafny.Array
				_ = _nw43
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw43 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) bool = (func(_260_v0 bool) func(_dafny.Int) bool {
						return func(_261_i0 _dafny.Int) bool {
							return _260_v0
						}
					})(_257_v0)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw43 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw43).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw43).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_259_v2 = _nw43
				var _262_v3 _dafny.Sequence
				_ = _262_v3
				_262_v3 = _dafny.SeqOf(_257_v0)
				var _263_v4 _dafny.Set
				_ = _263_v4
				_263_v4 = _dafny.SetOf(_257_v0, (_262_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(730), _dafny.IntOfUint32((_262_v3).Cardinality()))).Uint32()).(bool), _257_v0)
				_259_v2 = (func() _dafny.Array {
					if (_dafny.SetOf(_257_v0)).IsSubsetOf(_263_v4) {
						return _259_v2
					}
					return _259_v2
				})()
				var _264_v5 _dafny.Map
				_ = _264_v5
				_264_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F33, _257_v0)
				var _265_v6 _dafny.Map
				_ = _265_v6
				_265_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_262_v3), (_this).Fm10(globalState))
				var _266_v7 _dafny.Map
				_ = _266_v7
				_266_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v0, _257_v0)
				var _267_v8 D3
				_ = _267_v8
				_267_v8 = Companion_D3_.Create_DC9_(_262_v3)
				var _268_v9 _dafny.Map
				_ = _268_v9
				_268_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_266_v7).Cardinality(), (_267_v8).Dtor_cf17())
				(globalState).F2 = (func() bool {
					if (_264_v5).Contains(_this.F33) {
						return (_264_v5).Get(_this.F33).(bool)
					}
					return (func() bool {
						if (_265_v6).Contains(_dafny.SeqOf((func() _dafny.Sequence {
							if (_268_v9).Contains((_dafny.Zero).Minus((_this).F24())) {
								return (_268_v9).Get((_dafny.Zero).Minus((_this).F24())).(_dafny.Sequence)
							}
							return _262_v3
						})())) {
							return (_265_v6).Get(_dafny.SeqOf((func() _dafny.Sequence {
								if (_268_v9).Contains((_dafny.Zero).Minus((_this).F24())) {
									return (_268_v9).Get((_dafny.Zero).Minus((_this).F24())).(_dafny.Sequence)
								}
								return _262_v3
							})())).(bool)
						}
						return _257_v0
					})()
				})()
				var _269_v10 _dafny.Sequence
				_ = _269_v10
				_269_v10 = _dafny.UnicodeSeqOfUtf8Bytes("clyji")
				var _270_v11 _dafny.Set
				_ = _270_v11
				_270_v11 = _dafny.SetOf(_269_v10)
				var _271_v12 _dafny.Sequence
				_ = _271_v12
				_271_v12 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_this).F24(), (_270_v11).Cardinality()), (_this).F24(), (_dafny.MultiSetOf(_257_v0)).Cardinality())
				var _rhs29 _dafny.Array = _259_v2
				_ = _rhs29
				var _rhs30 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("s"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(888))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_272_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_273_i1 _dafny.Int) _dafny.CodePoint {
						return _272_v1
					}
				})(_258_v1))))).Cardinality())
				_ = _rhs30
				var _rhs31 _dafny.Int = (_this).F24()
				_ = _rhs31
				var _rhs32 _dafny.Sequence = _271_v12
				_ = _rhs32
				var _lhs25 *GlobalState = globalState
				_ = _lhs25
				var _lhs26 *GlobalState = globalState
				_ = _lhs26
				_259_v2 = _rhs29
				_lhs25.F9 = _rhs30
				_lhs26.F9 = _rhs31
				_271_v12 = _rhs32
			} else {
				var _274_v13 _dafny.MultiSet
				_ = _274_v13
				_274_v13 = _dafny.MultiSetOf(_this.F33)
				var _275_v14 _dafny.MultiSet
				_ = _275_v14
				_275_v14 = _dafny.MultiSetOf(_257_v0, _257_v0, _257_v0)
				var _276_v15 *C0
				_ = _276_v15
				var _nw44 *C0 = New_C0_()
				_ = _nw44
				_nw44.Ctor__((((_274_v13).Update(_this.F33, Companion_Default___.Abs((_this).F24()))).Cardinality()).Cmp((func() _dafny.Int {
					if (_275_v14).Contains(_257_v0) {
						return (_275_v14).Multiplicity(_257_v0)
					}
					return (_dafny.SetOf(true, !(_257_v0), _257_v0)).Cardinality()
				})()) != 0)
				_276_v15 = _nw44
				var _277_v16 _dafny.Array
				_ = _277_v16
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_6
				var _nw45 _dafny.Array
				_ = _nw45
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw45 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Sequence = func(_278_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(648))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg31 _dafny.Int) interface{} {
								return coer31(arg31)
							}
						}(func(_279_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('t')
						})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg32 _dafny.Int) interface{} {
								return coer32(arg32)
							}
						}(func(_280_i4 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('y')
						})))
					}
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
				_277_v16 = _nw45
				var _281_v17 _dafny.Sequence
				_ = _281_v17
				_281_v17 = _dafny.UnicodeSeqOfUtf8Bytes("wtih")
				var _282_v18 _dafny.CodePoint
				_ = _282_v18
				_282_v18 = _dafny.CodePoint('q')
				var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_277_v16), 0))
				_ = _index32
				(_277_v16).ArraySet1(_dafny.Companion_Sequence_.Update(_281_v17, (Companion_Default___.SafeIndex((_this).Fm11((_276_v15).F34(), (_276_v15).F34(), true, _this.F33, globalState), _dafny.IntOfUint32((_281_v17).Cardinality()))).Uint32(), _282_v18), (_index32).Int())
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_277_v16), 0))
				_ = _index33
				(_277_v16).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vcjpd"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(558))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}((func(_283_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_284_i5 _dafny.Int) _dafny.CodePoint {
						return _283_v18
					}
				})(_282_v18)))), (_index33).Int())
				_276_v15 = _276_v15
				var _285_v19 _dafny.Map
				_ = _285_v19
				_285_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_276_v15).F34(), _282_v18)
				var _286_v20 _dafny.Map
				_ = _286_v20
				_286_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_276_v15).F34()), _257_v0)
				_285_v19 = (_285_v19).Update((func() bool {
					if (_286_v20).Contains((_276_v15).F34()) {
						return (_286_v20).Get((_276_v15).F34()).(bool)
					}
					return _257_v0
				})(), _282_v18)
				var _287_v21 D0
				_ = _287_v21
				_287_v21 = Companion_D0_.Create_DC1_((_this).F24(), _dafny.IntOfInt64(949), (func() bool {
					if Companion_Default___.Fm2(_this.F33, (_dafny.Zero).Minus((_this).F24()), _282_v18, globalState) {
						return (_276_v15).F34()
					}
					return (_276_v15).F34()
				})())
				var _288_v22 _dafny.Map
				_ = _288_v22
				_288_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_257_v0), _this.F33)
				var _289_v23 _dafny.Sequence
				_ = _289_v23
				_289_v23 = _dafny.SeqOf((_276_v15).F34(), (_276_v15).F34(), _257_v0)
				var _290_v24 _dafny.Sequence
				_ = _290_v24
				_290_v24 = _dafny.SeqOf((func() _dafny.Map {
					if _257_v0 {
						return _288_v22
					}
					return _288_v22
				})(), (_288_v22).Update(true, _dafny.IntOfUint32((_289_v23).Cardinality())), _288_v22, _288_v22, _288_v22)
				var _rhs33 bool = (_276_v15).F34()
				_ = _rhs33
				var _rhs34 D0 = _287_v21
				_ = _rhs34
				var _rhs35 _dafny.Int = (_this).F24()
				_ = _rhs35
				var _rhs36 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(655))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg34 _dafny.Int) interface{} {
						return coer34(arg34)
					}
				}((func(_291_v22 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_292_i6 _dafny.Int) _dafny.Map {
						return _291_v22
					}
				})(_288_v22)))
				_ = _rhs36
				var _lhs27 *GlobalState = globalState
				_ = _lhs27
				var _lhs28 *GlobalState = globalState
				_ = _lhs28
				_lhs27.F2 = _rhs33
				_287_v21 = _rhs34
				_lhs28.F1 = _rhs35
				_290_v24 = _rhs36
			}
			var _293_v25 _dafny.Array
			_ = _293_v25
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw46
			_293_v25 = _nw46
			var _294_v26 _dafny.Map
			_ = _294_v26
			_294_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v0, _257_v0)
			var _295_v27 _dafny.Sequence
			_ = _295_v27
			_295_v27 = _dafny.SeqOf(_294_v26)
			var _296_v28 _dafny.Array
			_ = _296_v28
			var _nwElement0_5 _dafny.Map = Companion_Default___.Fm13(globalState)
			_ = _nwElement0_5
			var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
			_ = _nw47
			(_nw47).ArraySet1(_nwElement0_5, 0)
			(_nw47).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v0, _257_v0), 1)
			(_nw47).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v0, _257_v0), 2)
			(_nw47).ArraySet1(_294_v26, 3)
			(_nw47).ArraySet1(_294_v26, 4)
			(_nw47).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v0, _257_v0), 5)
			(_nw47).ArraySet1(_294_v26, 6)
			(_nw47).ArraySet1(_294_v26, 7)
			(_nw47).ArraySet1(_294_v26, 8)
			(_nw47).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v0, !(_257_v0)), 9)
			(_nw47).ArraySet1(_294_v26, 10)
			(_nw47).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _257_v0), 11)
			(_nw47).ArraySet1(_294_v26, 12)
			(_nw47).ArraySet1(_294_v26, 13)
			(_nw47).ArraySet1(_294_v26, 14)
			(_nw47).ArraySet1((_295_v27).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_295_v27).Cardinality()))).Uint32()).(_dafny.Map), 15)
			(_nw47).ArraySet1(_294_v26, 16)
			(_nw47).ArraySet1(_294_v26, 17)
			(_nw47).ArraySet1(_294_v26, 18)
			(_nw47).ArraySet1(_294_v26, 19)
			(_nw47).ArraySet1(_294_v26, 20)
			(_nw47).ArraySet1((_294_v26).Update(_257_v0, _257_v0), 21)
			(_nw47).ArraySet1(_294_v26, 22)
			(_nw47).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v0, !(_257_v0)), 23)
			(_nw47).ArraySet1(_294_v26, 24)
			_296_v28 = _nw47
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_293_v25), 0))
			_ = _index34
			(_293_v25).ArraySet1(_296_v28, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_293_v25), 0))
			_ = _index35
			(_293_v25).ArraySet1(_296_v28, (_index35).Int())
			var _297_v29 _dafny.Set
			_ = _297_v29
			_297_v29 = _dafny.SetOf(_dafny.IntOfInt64(342))
			_297_v29 = _297_v29
			var _298_v30 _dafny.Array
			_ = _298_v30
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw48
			_298_v30 = _nw48
			var _299_v31 _dafny.Sequence
			_ = _299_v31
			_299_v31 = _dafny.SeqOf(_298_v30, _298_v30, _298_v30)
			var _300_v32 _dafny.Int
			_ = _300_v32
			var _301_v33 bool
			_ = _301_v33
			var _302_v34 _dafny.Int
			_ = _302_v34
			var _303_v35 _dafny.CodePoint
			_ = _303_v35
			var _out13 _dafny.Int
			_ = _out13
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Int
			_ = _out15
			var _out16 _dafny.CodePoint
			_ = _out16
			_out13, _out14, _out15, _out16 = (_this).M8((_299_v31).Select((Companion_Default___.SafeIndex(_this.F33, _dafny.IntOfUint32((_299_v31).Cardinality()))).Uint32()).(_dafny.Array), globalState)
			_300_v32 = _out13
			_301_v33 = _out14
			_302_v34 = _out15
			_303_v35 = _out16
		} else {
			var _304_v36 bool
			_ = _304_v36
			_304_v36 = true
			if _304_v36 {
				(globalState).F2 = _304_v36
				var _305_v37 _dafny.Map
				_ = _305_v37
				_305_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_304_v36, _this.F33)
				var _306_v38 D0
				_ = _306_v38
				_306_v38 = Companion_D0_.Create_DC1_(_this.F33, (_305_v37).Cardinality(), _304_v36)
				var _307_v39 _dafny.Map
				_ = _307_v39
				_307_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_306_v38).Dtor_cf3(), _this.F33)
				var _308_v40 _dafny.Map
				_ = _308_v40
				_308_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _304_v36)
				var _309_v41 _dafny.CodePoint
				_ = _309_v41
				_309_v41 = _dafny.CodePoint('c')
				var _310_v42 _dafny.MultiSet
				_ = _310_v42
				_310_v42 = _dafny.MultiSetOf(_304_v36)
				var _311_v43 D1
				_ = _311_v43
				_311_v43 = Companion_D1_.Create_DC4_((_this.F33).Cmp((_307_v39).Cardinality()) <= 0, (func() bool {
					if (_308_v40).Contains(_dafny.CodePoint('d')) {
						return (_308_v40).Get(_dafny.CodePoint('d')).(bool)
					}
					return _304_v36
				})(), _309_v41, (func() _dafny.Int {
					if (_310_v42).Contains(true) {
						return (_310_v42).Multiplicity(true)
					}
					return _this.F33
				})(), (_this).Fm11(_304_v36, !(_304_v36), _304_v36, _this.F33, globalState))
				_311_v43 = _311_v43
				var _312_v44 *C0
				_ = _312_v44
				var _nw49 *C0 = New_C0_()
				_ = _nw49
				_nw49.Ctor__(_304_v36)
				_312_v44 = _nw49
				_310_v42 = (_310_v42).Intersection(Companion_Default___.Fm14(Companion_Default___.Fm15(_dafny.IntOfInt64(224), _304_v36, (_this).F24(), (_312_v44).F34(), globalState), _this.F33, globalState))
				var _313_v45 _dafny.Sequence
				_ = _313_v45
				_313_v45 = _dafny.SeqOf(_this.F33, _this.F33)
				var _314_v46 D4
				_ = _314_v46
				_314_v46 = Companion_D4_.Create_DC12_(_313_v45)
				var _315_v47 _dafny.Map
				_ = _315_v47
				_315_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((((_310_v42).Update((_312_v44).F34(), Companion_Default___.Abs(_this.F33))).Cardinality()).Cmp(_this.F33) < 0, (_314_v46).Dtor_cf21())
				_315_v47 = (_315_v47).Update((_312_v44).F34(), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_this.F33, (_this).F24()), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F24()), _this.F33)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_this.F33, (_this).F24())).Cardinality()))).Uint32(), (_this).F24()))
			} else {
				var _316_v48 _dafny.Map
				_ = _316_v48
				_316_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (func() _dafny.Int {
					if _304_v36 {
						return (_dafny.Zero).Minus((_this).F24())
					}
					return (_this).F24()
				})())
				_316_v48 = (_316_v48).Update((_this).F24(), (_this).F24())
				var _317_v49 _dafny.Array
				_ = _317_v49
				var _nw50 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(24))
				_ = _nw50
				_317_v49 = _nw50
				_317_v49 = _317_v49
				var _318_v50 D0
				_ = _318_v50
				_318_v50 = Companion_D0_.Create_DC1_((_dafny.Zero).Minus((_this).F24()), _dafny.IntOfInt64(300), _304_v36)
				(globalState).F2 = !(((_this).F24()).Cmp((_this.F33).Plus((_318_v50).Dtor_cf2())) == 0)
				var _319_v51 _dafny.Sequence
				_ = _319_v51
				_319_v51 = _dafny.UnicodeSeqOfUtf8Bytes("i")
				(globalState).F17 = _319_v51
				var _320_v52 _dafny.Map
				_ = _320_v52
				_320_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(329))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}(func(_321_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})))
				_320_v52 = (_320_v52).Update((_this).F24(), _319_v51)
			}
			var _322_v53 _dafny.Array
			_ = _322_v53
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_7
			var _nw51 _dafny.Array
			_ = _nw51
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw51 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Int = func(_323_i8 _dafny.Int) _dafny.Int {
					return (_323_i8).Times(_this.F33)
				}
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw51 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw51).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw51).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_322_v53 = _nw51
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_322_v53), 0))
			_ = _index36
			(_322_v53).ArraySet1((_this).F24(), (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_322_v53), 0))
			_ = _index37
			(_322_v53).ArraySet1(((_this).F24()).Times((_this).F24()), (_index37).Int())
			(globalState).F14 = _this.F33
			var _324_v54 _dafny.Array
			_ = _324_v54
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_8
			var _nw52 _dafny.Array
			_ = _nw52
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw52 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = func(_325_i9 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("u")
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw52 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw52).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw52).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_324_v54 = _nw52
			var _326_v55 _dafny.Sequence
			_ = _326_v55
			_326_v55 = _dafny.UnicodeSeqOfUtf8Bytes("xwklb")
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_324_v54), 0))
			_ = _index38
			(_324_v54).ArraySet1(_326_v55, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_324_v54), 0))
			_ = _index39
			(_324_v54).ArraySet1(_326_v55, (_index39).Int())
			var _327_v56 _dafny.Map
			_ = _327_v56
			_327_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_304_v36, _304_v36)
			var _328_v57 _dafny.Map
			_ = _328_v57
			_328_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_304_v36, (_327_v56).Cardinality())
			var _rhs37 _dafny.Int = (Companion_D4_.Create_DC13_(_this.F33, _328_v57, _326_v55)).Dtor_cf22()
			_ = _rhs37
			var _rhs38 _dafny.Int = (_this).F24()
			_ = _rhs38
			var _rhs39 bool = _304_v36
			_ = _rhs39
			var _rhs40 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(133))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_329_i10 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			}))
			_ = _rhs40
			var _lhs29 *GlobalState = globalState
			_ = _lhs29
			var _lhs30 *GlobalState = globalState
			_ = _lhs30
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			_lhs29.F18 = _rhs37
			_lhs30.F9 = _rhs38
			_304_v36 = _rhs39
			_lhs31.F17 = _rhs40
		}
		(globalState).F9 = (_this).F24()
		var _330_v58 bool
		_ = _330_v58
		_330_v58 = true
		r0 = _330_v58
		_330_v58 = (_this).Fm10(globalState)
		var _331_v59 _dafny.Sequence
		_ = _331_v59
		_331_v59 = _dafny.UnicodeSeqOfUtf8Bytes("tvcjan")
		var _332_v60 _dafny.CodePoint
		_ = _332_v60
		_332_v60 = _dafny.CodePoint('t')
		var _333_v61 _dafny.Sequence
		_ = _333_v61
		_333_v61 = _dafny.SeqOf(_331_v59, _331_v59, Companion_Default___.Fm4(_330_v58, _this.F33, _332_v60, !(_330_v58), globalState))
		var _334_i11 _dafny.Int
		_ = _334_i11
		_334_i11 = _dafny.Zero
		{
			for !_dafny.Companion_Sequence_.Equal(_333_v61, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-168))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}((func(_376_v59 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_377_i12 _dafny.Int) _dafny.Sequence {
					return _376_v59
				}
			})(_331_v59)))) {
				{
					if (_334_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_334_i11 = (_334_i11).Plus(_dafny.One)
					var _335_v62 _dafny.Set
					_ = _335_v62
					_335_v62 = _dafny.SetOf(_this.F33)
					var _336_v63 _dafny.Map
					_ = _336_v63
					_336_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v58, (_335_v62).Cardinality())
					var _337_v64 _dafny.Map
					_ = _337_v64
					_337_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F33, _332_v60)
					var _338_v65 _dafny.Map
					_ = _338_v65
					_338_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_337_v64).Cardinality(), !(_330_v58))
					var _339_v66 _dafny.Array
					_ = _339_v66
					var _nwElement0_6 _dafny.Int = _this.F33
					_ = _nwElement0_6
					var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(19))
					_ = _nw53
					(_nw53).ArraySet1(_nwElement0_6, 0)
					(_nw53).ArraySet1(_this.F33, 1)
					(_nw53).ArraySet1((_336_v63).Cardinality(), 2)
					(_nw53).ArraySet1((_this).F24(), 3)
					(_nw53).ArraySet1(_dafny.IntOfInt64(676), 4)
					(_nw53).ArraySet1(_dafny.IntOfUint32((_331_v59).Cardinality()), 5)
					(_nw53).ArraySet1((Companion_Default___.Fm16(globalState)).Cardinality(), 6)
					(_nw53).ArraySet1((_dafny.Zero).Minus((_this).F24()), 7)
					(_nw53).ArraySet1((_this).F24(), 8)
					(_nw53).ArraySet1((_338_v65).Cardinality(), 9)
					(_nw53).ArraySet1((_this).F24(), 10)
					(_nw53).ArraySet1(((_this).F24()).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_331_v59, _331_v59)).Cardinality())), 11)
					(_nw53).ArraySet1(Companion_Default___.SafeModuloInt(_this.F33, _dafny.IntOfInt64(-6)), 12)
					(_nw53).ArraySet1((_this).F24(), 13)
					(_nw53).ArraySet1((_this).F24(), 14)
					(_nw53).ArraySet1((_338_v65).Cardinality(), 15)
					(_nw53).ArraySet1(_this.F33, 16)
					(_nw53).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfInt64(-523), _this.F33, globalState), 17)
					(_nw53).ArraySet1((_this).F24(), 18)
					_339_v66 = _nw53
					var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))
					_ = _index40
					(_339_v66).ArraySet1(_this.F33, (_index40).Int())
					var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))
					_ = _index41
					var _rhs41 _dafny.Int = (_this).F24()
					_ = _rhs41
					var _rhs42 _dafny.Array = (func() _dafny.Array {
						if _330_v58 {
							return _339_v66
						}
						return _339_v66
					})()
					_ = _rhs42
					var _lhs32 _dafny.Array = _339_v66
					_ = _lhs32
					var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))
					_ = _lhs33
					(_lhs32).ArraySet1(_rhs41, (_lhs33).Int())
					_339_v66 = _rhs42
					var _340_v67 D0
					_ = _340_v67
					_340_v67 = Companion_D0_.Create_DC2_(_330_v58, _dafny.IntOfInt64(-118), _332_v60)
					var _341_v68 _dafny.Map
					_ = _341_v68
					_341_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_340_v67, _330_v58)
					var _342_v69 _dafny.Map
					_ = _342_v69
					_342_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v58, (func() bool {
						if _330_v58 {
							return _330_v58
						}
						return (func() bool {
							if (_341_v68).Contains(_340_v67) {
								return (_341_v68).Get(_340_v67).(bool)
							}
							return _330_v58
						})()
					})())
					_342_v69 = (_342_v69).Update(((_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int)).Cmp((_this).F24()) != 0, _330_v58)
					var _343_v70 D2
					_ = _343_v70
					_343_v70 = Companion_D2_.Create_DC8_()
					var _source3 D2 = _343_v70
					_ = _source3
					if _source3.Is_DC8() {
						var _344_v71 _dafny.Sequence
						_ = _344_v71
						_344_v71 = _dafny.SeqOf(_dafny.IntOfUint32((_331_v59).Cardinality()))
						var _345_v72 _dafny.Sequence
						_ = _345_v72
						_345_v72 = _dafny.SeqOf(_330_v58, _330_v58)
						var _346_v73 _dafny.Array
						_ = _346_v73
						var _nwElement0_7 _dafny.Map = Companion_Default___.Fm17(_dafny.IntOfInt64(790), (_this).F24(), _dafny.IntOfInt64(580), false, globalState)
						_ = _nwElement0_7
						var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(27))
						_ = _nw54
						(_nw54).ArraySet1(_nwElement0_7, 0)
						(_nw54).ArraySet1(_336_v63, 1)
						(_nw54).ArraySet1(_336_v63, 2)
						(_nw54).ArraySet1((_336_v63).Update(_330_v58, (_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int)), 3)
						(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int), (_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int), _332_v60, globalState), (_this).F24()), 4)
						(_nw54).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v58, _dafny.IntOfInt64(641))).Update(_330_v58, _this.F33), 5)
						(_nw54).ArraySet1(Companion_Default___.Fm17(_dafny.IntOfUint32((Companion_Default___.Fm4(_330_v58, (_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int), _332_v60, _330_v58, globalState)).Cardinality()), _this.F33, _dafny.IntOfUint32((_344_v71).Cardinality()), _330_v58, globalState), 6)
						(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v58, _dafny.IntOfInt64(-683)), 7)
						(_nw54).ArraySet1(_336_v63, 8)
						(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v58, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v58, (_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_331_v59).Cardinality()))).Cardinality()), 9)
						(_nw54).ArraySet1(_336_v63, 10)
						(_nw54).ArraySet1(_336_v63, 11)
						(_nw54).ArraySet1(_336_v63, 12)
						(_nw54).ArraySet1(_336_v63, 13)
						(_nw54).ArraySet1(_336_v63, 14)
						(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v58, _this.F33), 15)
						(_nw54).ArraySet1(Companion_Default___.Fm17((Companion_Default___.Fm18(globalState)).Cardinality(), (_this).F24(), _dafny.IntOfUint32((_345_v72).Cardinality()), _330_v58, globalState), 16)
						(_nw54).ArraySet1(_336_v63, 17)
						(_nw54).ArraySet1(_336_v63, 18)
						(_nw54).ArraySet1(_336_v63, 19)
						(_nw54).ArraySet1(_336_v63, 20)
						(_nw54).ArraySet1(_336_v63, 21)
						(_nw54).ArraySet1(_336_v63, 22)
						(_nw54).ArraySet1(_336_v63, 23)
						(_nw54).ArraySet1((_336_v63).Update(_330_v58, (_this).F24()), 24)
						(_nw54).ArraySet1(_336_v63, 25)
						(_nw54).ArraySet1(_336_v63, 26)
						_346_v73 = _nw54
						var _347_v74 _dafny.Map
						_ = _347_v74
						_347_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg37 _dafny.Int) interface{} {
								return coer37(arg37)
							}
						}((func(_348_v60 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_349_i13 _dafny.Int) _dafny.CodePoint {
								return _348_v60
							}
						})(_332_v60)))).Cardinality()), _346_v73)
						var _350_v75 _dafny.Sequence
						_ = _350_v75
						_350_v75 = _dafny.SeqOf(_346_v73)
						_347_v74 = (_347_v74).Update((_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int), (_350_v75).Select((Companion_Default___.SafeIndex(_this.F33, _dafny.IntOfUint32((_350_v75).Cardinality()))).Uint32()).(_dafny.Array))
						var _351_v76 D3
						_ = _351_v76
						_351_v76 = Companion_D3_.Create_DC11_(_this.F33, _330_v58, _330_v58)
						var _352_v77 _dafny.MultiSet
						_ = _352_v77
						_352_v77 = _dafny.MultiSetOf(_this.F33, (_this).F24(), _this.F33, (_this).Fm11(_330_v58, !(_330_v58), _330_v58, _this.F33, globalState), (_351_v76).Dtor_cf18())
						var _353_v78 _dafny.Set
						_ = _353_v78
						_353_v78 = _dafny.SetOf(_339_v66)
						var _rhs43 bool = _330_v58
						_ = _rhs43
						var _rhs44 _dafny.MultiSet = _352_v77
						_ = _rhs44
						var _rhs45 _dafny.Set = (_353_v78).Intersection(_353_v78)
						_ = _rhs45
						var _lhs34 *GlobalState = globalState
						_ = _lhs34
						_lhs34.F2 = _rhs43
						_352_v77 = _rhs44
						_353_v78 = _rhs45
						var _354_v79 _dafny.MultiSet
						_ = _354_v79
						_354_v79 = _dafny.MultiSetOf(_330_v58, _330_v58)
						var _355_v80 _dafny.Map
						_ = _355_v80
						_355_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_354_v79, ((_this).F24()).Times((_this).F24()))
						_355_v80 = _355_v80
						var _356_v81 _dafny.Array
						_ = _356_v81
						var _len0_9 _dafny.Int = _dafny.IntOfInt64(24)
						_ = _len0_9
						var _nw55 _dafny.Array
						_ = _nw55
						if _len0_9.Cmp(_dafny.Zero) == 0 {
							_nw55 = _dafny.NewArray(_len0_9)
						} else {
							var _init9 func(_dafny.Int) bool = (func(_357_v58 bool) func(_dafny.Int) bool {
								return func(_358_i14 _dafny.Int) bool {
									return _357_v58
								}
							})(_330_v58)
							_ = _init9
							var _element0_9 = _init9(_dafny.Zero)
							_ = _element0_9
							_nw55 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
							(_nw55).ArraySet1(_element0_9, 0)
							var _nativeLen0_9 = (_len0_9).Int()
							_ = _nativeLen0_9
							for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
								(_nw55).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
							}
						}
						_356_v81 = _nw55
						var _359_v82 _dafny.Map
						_ = _359_v82
						_359_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19(_330_v58, (_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int), globalState), _356_v81)
						var _360_v83 _dafny.Sequence
						_ = _360_v83
						_360_v83 = _dafny.SeqOf(_dafny.SeqOf(_330_v58))
						_356_v81 = (func() _dafny.Array {
							if (_359_v82).Contains(_dafny.Companion_Sequence_.Concatenate(_360_v83, _dafny.Companion_Sequence_.Update(_360_v83, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_331_v59).Cardinality()), _dafny.IntOfUint32((_360_v83).Cardinality()))).Uint32(), _dafny.SeqOf(true)))) {
								return (_359_v82).Get(_dafny.Companion_Sequence_.Concatenate(_360_v83, _dafny.Companion_Sequence_.Update(_360_v83, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_331_v59).Cardinality()), _dafny.IntOfUint32((_360_v83).Cardinality()))).Uint32(), _dafny.SeqOf(true)))).(_dafny.Array)
							}
							return _356_v81
						})()
					} else {
						var _361___mcc_h0 _dafny.Array = _source3.Get_().(D2_DC7).Cf16
						_ = _361___mcc_h0
						var _362_cf16 _dafny.Array = _361___mcc_h0
						_ = _362_cf16
						(globalState).F2 = (Companion_Default___.SafeModuloInt((_338_v65).Cardinality(), (_dafny.Zero).Minus((_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int)))).Cmp((_this).F24()) != 0
						var _363_v84 _dafny.MultiSet
						_ = _363_v84
						_363_v84 = _dafny.MultiSetOf(_330_v58)
						var _364_v85 _dafny.Sequence
						_ = _364_v85
						_364_v85 = _dafny.SeqOf((_363_v84).Cardinality())
						var _365_v86 _dafny.Sequence
						_ = _365_v86
						_365_v86 = _dafny.SeqOf(_330_v58, _330_v58, _330_v58)
						var _366_v87 _dafny.Map
						_ = _366_v87
						_366_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_365_v86, _330_v58)
						var _367_v88 _dafny.Map
						_ = _367_v88
						_367_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_332_v60, _330_v58)
						var _368_v89 _dafny.MultiSet
						_ = _368_v89
						_368_v89 = _dafny.MultiSetOf(_330_v58)
						var _369_v90 _dafny.Sequence
						_ = _369_v90
						_369_v90 = _dafny.SeqOf(_363_v84, (_368_v89))
						var _370_v91 _dafny.Map
						_ = _370_v91
						_370_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F33, (_this).F24())
						var _371_v92 D1
						_ = _371_v92
						_371_v92 = Companion_D1_.Create_DC4_(_330_v58, _330_v58, _332_v60, (_370_v91).Cardinality(), (_this).F24())
						var _372_v94 _dafny.Array
						_ = _372_v94
						var _nwElement0_8 bool = ((_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int)).Cmp(_this.F33) != 0
						_ = _nwElement0_8
						var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(24))
						_ = _nw56
						(_nw56).ArraySet1(_nwElement0_8, 0)
						(_nw56).ArraySet1(true, 1)
						(_nw56).ArraySet1(Companion_Default___.Fm2((_364_v85).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_331_v59).Cardinality()), _dafny.IntOfUint32((_364_v85).Cardinality()))).Uint32()).(_dafny.Int), _this.F33, _332_v60, globalState), 2)
						(_nw56).ArraySet1(_dafny.Companion_Sequence_.Contains(_364_v85, (_this).F24()), 3)
						(_nw56).ArraySet1(_330_v58, 4)
						(_nw56).ArraySet1(Companion_Default___.Fm2(_this.F33, _this.F33, _dafny.CodePoint('c'), globalState), 5)
						(_nw56).ArraySet1(_330_v58, 6)
						(_nw56).ArraySet1(_330_v58, 7)
						(_nw56).ArraySet1(_330_v58, 8)
						(_nw56).ArraySet1(_330_v58, 9)
						(_nw56).ArraySet1(!(true), 10)
						(_nw56).ArraySet1(_330_v58, 11)
						(_nw56).ArraySet1((func() bool {
							if (_366_v87).Contains(_365_v86) {
								return (_366_v87).Get(_365_v86).(bool)
							}
							return _330_v58
						})(), 12)
						(_nw56).ArraySet1(_330_v58, 13)
						(_nw56).ArraySet1(!(Companion_Default___.Fm2((_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int), (_363_v84).Cardinality(), _332_v60, globalState)), 14)
						(_nw56).ArraySet1(_330_v58, 15)
						(_nw56).ArraySet1(_330_v58, 16)
						(_nw56).ArraySet1(_330_v58, 17)
						(_nw56).ArraySet1(((_367_v88).Cardinality()).Cmp(((_369_v90).Select((Companion_Default___.SafeIndex((_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_369_v90).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()) <= 0, 18)
						(_nw56).ArraySet1(Companion_Default___.Fm2((_370_v91).Cardinality(), (_371_v92).Dtor_cf12(), _332_v60, globalState), 19)
						(_nw56).ArraySet1(false, 20)
						(_nw56).ArraySet1(((_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int)).Cmp(((func() _dafny.Map {
							var _coll15 = _dafny.NewMapBuilder()
							_ = _coll15
							for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-804), _dafny.IntOfInt64(751))); ; {
								_compr_15, _ok16 := _iter16()
								if !_ok16 {
									break
								}
								var _373_v93 _dafny.Int
								_373_v93 = interface{}(_compr_15).(_dafny.Int)
								if ((_dafny.IntOfInt64(-804)).Cmp(_373_v93) <= 0) && ((_373_v93).Cmp(_dafny.IntOfInt64(751)) < 0) {
									_coll15.Add(Companion_Default___.SafeDivisionInt(_373_v93, (_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v58, _332_v60)).Cardinality())
								}
							}
							return _coll15.ToMap()
						}()).Update((_this).F24(), (_339_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_339_v66), 0))).Int()).(_dafny.Int))).Cardinality()) <= 0, 21)
						(_nw56).ArraySet1(_330_v58, 22)
						(_nw56).ArraySet1(_330_v58, 23)
						_372_v94 = _nw56
						_372_v94 = _372_v94
						var _374_v95 _dafny.Set
						_ = _374_v95
						_374_v95 = _dafny.SetOf(_365_v86, _365_v86)
						_374_v95 = ((func() _dafny.Set {
							var _coll16 = _dafny.NewBuilder()
							_ = _coll16
							for _iter17 := _dafny.Iterate((Companion_Default___.Fm20(globalState)).Keys().Elements()); ; {
								_compr_16, _ok17 := _iter17()
								if !_ok17 {
									break
								}
								var _375_v96 _dafny.Sequence
								_375_v96 = interface{}(_compr_16).(_dafny.Sequence)
								if (Companion_Default___.Fm20(globalState)).Contains(_375_v96) {
									_coll16.Add(_375_v96)
								}
							}
							return _coll16.ToSet()
						}()).Difference(_374_v95)).Difference(_374_v95)
						(globalState).F17 = _331_v59
					}
					(globalState).F1 = ((_335_v62).Cardinality()).Times(_this.F33)
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_331_v59, _331_v59), _dafny.UnicodeSeqOfUtf8Bytes("ndbanqk"))).Cardinality())
		r0 = !(!(_330_v58))
		return r0
	}
}
func (_this *C1) M7(p0 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		(globalState).F18 = _this.F33
		var _378_v0 bool
		_ = _378_v0
		_378_v0 = true
		var _379_v1 _dafny.Map
		_ = _379_v1
		_379_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), _378_v0)
		var _380_v2 _dafny.Sequence
		_ = _380_v2
		_380_v2 = _dafny.SeqOf((_379_v1).Cardinality())
		var _381_v3 _dafny.Map
		_ = _381_v3
		_381_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_380_v2, _380_v2), (_this).F24())
		(globalState).F1 = (func() _dafny.Int {
			if (_381_v3).Contains(_380_v2) {
				return (_381_v3).Get(_380_v2).(_dafny.Int)
			}
			return (_this).F24()
		})()
		if _378_v0 {
			r0 = _378_v0
			var _382_v4 _dafny.CodePoint
			_ = _382_v4
			_382_v4 = _dafny.CodePoint('k')
			var _383_v5 _dafny.Map
			_ = _383_v5
			_383_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(546), p0)
			var _384_v6 D1
			_ = _384_v6
			_384_v6 = Companion_D1_.Create_DC4_(_378_v0, _378_v0, _382_v4, (_this).F24(), (_383_v5).Cardinality())
			var _pat_let_tv0 = _382_v4
			_ = _pat_let_tv0
			var _source4 D1 = func(_pat_let0_0 D1) D1 {
				return func(_385_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let1_0 _dafny.CodePoint) D1 {
						return func(_386_dt__update_hcf10_h0 _dafny.CodePoint) D1 {
							return Companion_D1_.Create_DC4_((_385_dt__update__tmp_h0).Dtor_cf8(), (_385_dt__update__tmp_h0).Dtor_cf9(), _386_dt__update_hcf10_h0, (_385_dt__update__tmp_h0).Dtor_cf11(), (_385_dt__update__tmp_h0).Dtor_cf12())
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_384_v6)
			_ = _source4
			if _source4.Is_DC4() {
				var _387___mcc_h0 bool = _source4.Get_().(D1_DC4).Cf8
				_ = _387___mcc_h0
				var _388___mcc_h1 bool = _source4.Get_().(D1_DC4).Cf9
				_ = _388___mcc_h1
				var _389___mcc_h2 _dafny.CodePoint = _source4.Get_().(D1_DC4).Cf10
				_ = _389___mcc_h2
				var _390___mcc_h3 _dafny.Int = _source4.Get_().(D1_DC4).Cf11
				_ = _390___mcc_h3
				var _391___mcc_h4 _dafny.Int = _source4.Get_().(D1_DC4).Cf12
				_ = _391___mcc_h4
				var _392_cf12 _dafny.Int = _391___mcc_h4
				_ = _392_cf12
				var _393_cf11 _dafny.Int = _390___mcc_h3
				_ = _393_cf11
				var _394_cf10 _dafny.CodePoint = _389___mcc_h2
				_ = _394_cf10
				var _395_cf9 bool = _388___mcc_h1
				_ = _395_cf9
				var _396_cf8 bool = _387___mcc_h0
				_ = _396_cf8
				var _397_v7 _dafny.Array
				_ = _397_v7
				var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
				_ = _nw57
				_397_v7 = _nw57
				_397_v7 = _397_v7
				(globalState).F2 = _378_v0
				var _398_v8 D2
				_ = _398_v8
				_398_v8 = Companion_D2_.Create_DC8_()
				var _399_v9 _dafny.Set
				_ = _399_v9
				_399_v9 = _dafny.SetOf(_378_v0)
				var _400_v10 _dafny.Map
				_ = _400_v10
				_400_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-689))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_401_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_402_i0 _dafny.Int) _dafny.CodePoint {
						return _401_v4
					}
				})(_382_v4))), true)
				var _403_v11 _dafny.Sequence
				_ = _403_v11
				_403_v11 = _dafny.UnicodeSeqOfUtf8Bytes("qojo")
				var _404_v12 _dafny.Map
				_ = _404_v12
				_404_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_378_v0), _396_cf8)
				var _405_v13 _dafny.Sequence
				_ = _405_v13
				_405_v13 = _dafny.SeqOf(_404_v12)
				var _rhs46 D2 = _398_v8
				_ = _rhs46
				var _rhs47 bool = (func() bool {
					if (_399_v9).IsSubsetOf(_399_v9) {
						return _396_cf8
					}
					return (func() bool {
						if _396_cf8 {
							return _378_v0
						}
						return _395_cf9
					})()
				})()
				_ = _rhs47
				var _rhs48 _dafny.Int = (_dafny.Zero).Minus((((_400_v10).Update(_dafny.Companion_Sequence_.Update(_403_v11, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_380_v2), _this.F33)).Cardinality(), _dafny.IntOfUint32((_403_v11).Cardinality()))).Uint32(), _382_v4), _395_cf9)).Cardinality()).Times(((_this).F24()).Plus(_dafny.IntOfUint32((_405_v13).Cardinality()))))
				_ = _rhs48
				var _lhs35 *GlobalState = globalState
				_ = _lhs35
				_398_v8 = _rhs46
				_396_cf8 = _rhs47
				_lhs35.F1 = _rhs48
				var _406_v14 _dafny.Map
				_ = _406_v14
				_406_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_394_cf10, !(_395_cf9) || (_395_cf9))
				_406_v14 = (_406_v14).Update(_382_v4, _378_v0)
			} else if _source4.Is_DC5() {
				var _407___mcc_h5 bool = _source4.Get_().(D1_DC5).Cf13
				_ = _407___mcc_h5
				var _408___mcc_h6 bool = _source4.Get_().(D1_DC5).Cf14
				_ = _408___mcc_h6
				var _409___mcc_h7 bool = _source4.Get_().(D1_DC5).Cf15
				_ = _409___mcc_h7
				var _410_cf15 bool = _409___mcc_h7
				_ = _410_cf15
				var _411_cf14 bool = _408___mcc_h6
				_ = _411_cf14
				var _412_cf13 bool = _407___mcc_h5
				_ = _412_cf13
				var _413_v15 _dafny.Array
				_ = _413_v15
				var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(22))
				_ = _nw58
				_413_v15 = _nw58
				var _414_v16 _dafny.MultiSet
				_ = _414_v16
				_414_v16 = _dafny.MultiSetOf((p0).Select((Companion_Default___.SafeIndex(_this.F33, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(bool))
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_413_v15), 0))
				_ = _index42
				(_413_v15).ArraySet1((func() _dafny.MultiSet {
					if _378_v0 {
						return _414_v16
					}
					return _414_v16
				})(), (_index42).Int())
				var _415_v17 _dafny.Map
				_ = _415_v17
				_415_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, _414_v16)
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_413_v15), 0))
				_ = _index43
				(_413_v15).ArraySet1((func() _dafny.MultiSet {
					if (_415_v17).Contains(false) {
						return (_415_v17).Get(false).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(p0)
				})(), (_index43).Int())
				var _416_v18 _dafny.Map
				_ = _416_v18
				_416_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), (_this).F24())
				_416_v18 = (_416_v18).Update(_this.F33, _this.F33)
				var _417_v19 _dafny.Array
				_ = _417_v19
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_10
				var _nw59 _dafny.Array
				_ = _nw59
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw59 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) bool = (func(_418_v0 bool) func(_dafny.Int) bool {
						return func(_419_i1 _dafny.Int) bool {
							return _418_v0
						}
					})(_378_v0)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw59 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw59).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw59).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_417_v19 = _nw59
				var _420_v20 _dafny.MultiSet
				_ = _420_v20
				_420_v20 = _dafny.MultiSetOf(_417_v19)
				var _421_v21 _dafny.Set
				_ = _421_v21
				_421_v21 = _dafny.SetOf(_378_v0)
				var _422_v22 _dafny.Map
				_ = _422_v22
				_422_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, (_421_v21).Cardinality())
				var _423_v23 _dafny.Sequence
				_ = _423_v23
				_423_v23 = _dafny.UnicodeSeqOfUtf8Bytes("lxpkr")
				var _424_v24 _dafny.Int
				_ = _424_v24
				var _out17 _dafny.Int
				_ = _out17
				_out17 = Companion_Default___.M0(((func() _dafny.MultiSet {
					if _411_cf14 {
						return _dafny.MultiSetOf(_417_v19, _417_v19)
					}
					return _420_v20
				})()).Cardinality(), (Companion_D4_.Create_DC13_((_dafny.Zero).Minus((_this).F24()), _422_v22, _423_v23)).Dtor_cf22(), globalState)
				_424_v24 = _out17
				var _425_v25 D3
				_ = _425_v25
				_425_v25 = Companion_D3_.Create_DC11_((_this).F24(), _410_cf15, _412_cf13)
				(_this).F33 = ((_425_v25).Dtor_cf18()).Plus((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F33)))
			} else if _source4.Is_DC6() {
				var _426_v26 _dafny.Array
				_ = _426_v26
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_11
				var _nw60 _dafny.Array
				_ = _nw60
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw60 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) bool = (func(_427_v0 bool) func(_dafny.Int) bool {
						return func(_428_i2 _dafny.Int) bool {
							return _427_v0
						}
					})(_378_v0)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw60 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw60).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw60).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_426_v26 = _nw60
				var _429_v27 _dafny.Map
				_ = _429_v27
				_429_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F33, _426_v26)
				var _430_v28 _dafny.Sequence
				_ = _430_v28
				_430_v28 = _dafny.UnicodeSeqOfUtf8Bytes("lp")
				var _431_v29 _dafny.Map
				_ = _431_v29
				_431_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_429_v27).Merge(_429_v27), _430_v28)
				var _432_v30 _dafny.Sequence
				_ = _432_v30
				_432_v30 = _dafny.SeqOf(_430_v28)
				_431_v29 = (_431_v29).Update(_429_v27, _dafny.Companion_Sequence_.Concatenate((_432_v30).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_380_v2).Cardinality())), _dafny.IntOfUint32((_432_v30).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("ucibb")))
				var _433_v31 _dafny.Map
				_ = _433_v31
				_433_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, _378_v0)
				_433_v31 = (_433_v31).Update((_378_v0) && (false), _378_v0)
				var _434_v32 *C0
				_ = _434_v32
				var _nw61 *C0 = New_C0_()
				_ = _nw61
				_nw61.Ctor__(true)
				_434_v32 = _nw61
				var _435_v33 _dafny.Array
				_ = _435_v33
				var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(12))
				_ = _nw62
				_435_v33 = _nw62
				_435_v33 = _435_v33
			} else {
				var _436___mcc_h8 _dafny.Sequence = _source4.Get_().(D1_DC3).Cf7
				_ = _436___mcc_h8
				var _437_cf7 _dafny.Sequence = _436___mcc_h8
				_ = _437_cf7
				var _438_v34 _dafny.MultiSet
				_ = _438_v34
				_438_v34 = _dafny.MultiSetOf(_378_v0, _378_v0)
				r0 = (_438_v34).Equals((_438_v34).Difference(_438_v34))
				var _439_v35 _dafny.Map
				_ = _439_v35
				_439_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_378_v0) == (_378_v0), ((_this).F24()).Plus(_dafny.IntOfUint32((_437_cf7).Cardinality())))
				var _440_v36 _dafny.Array
				_ = _440_v36
				var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
				_ = _nw63
				_440_v36 = _nw63
				var _rhs49 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("pfkfe"), (Companion_Default___.SafeIndex(_this.F33, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pfkfe")).Cardinality()))).Uint32(), _382_v4), _437_cf7)).Cardinality()))
				_ = _rhs49
				var _rhs50 _dafny.MultiSet = _dafny.MultiSetOf(_440_v36)
				_ = _rhs50
				var _rhs51 _dafny.Int = (_this).F24()
				_ = _rhs51
				var _lhs36 *GlobalState = globalState
				_ = _lhs36
				var _lhs37 *GlobalState = globalState
				_ = _lhs37
				_439_v35 = _rhs49
				_lhs36.F19 = _rhs50
				_lhs37.F1 = _rhs51
				(globalState).F9 = _this.F33
				(globalState).F2 = false
			}
			var _441_v37 _dafny.Array
			_ = _441_v37
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw64
			_441_v37 = _nw64
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))
			_ = _index44
			(_441_v37).ArraySet1(_378_v0, (_index44).Int())
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))
			_ = _index45
			(_441_v37).ArraySet1(Companion_Default___.Fm2((_this).F24(), _this.F33, _dafny.CodePoint('w'), globalState), (_index45).Int())
			(globalState).F18 = (_this).F24()
			var _442_v38 D1
			_ = _442_v38
			_442_v38 = Companion_D1_.Create_DC5_((_441_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))).Int()).(bool), _378_v0, !(_378_v0))
			if (_442_v38).Dtor_cf13() {
				var _443_v39 _dafny.Map
				_ = _443_v39
				_443_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, _this.F33)
				var _444_v40 D4
				_ = _444_v40
				_444_v40 = Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pao")).Cardinality()), _443_v39, _dafny.UnicodeSeqOfUtf8Bytes("yvdcwh"))
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))
				_ = _index46
				(_441_v37).ArraySet1((Companion_Default___.Fm0(_this.F33, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_444_v40).Dtor_cf24(), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32(((_444_v40).Dtor_cf24()).Cardinality()))).Uint32(), _382_v4)).Cardinality()), globalState)).Cmp((func() _dafny.Int {
					if (_441_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))).Int()).(bool) {
						return _dafny.IntOfUint32((_380_v2).Cardinality())
					}
					return _this.F33
				})()) <= 0, (_index46).Int())
				_379_v1 = (_379_v1).Update(_this.F33, (!(_378_v0)) && (_378_v0))
				r0 = !(_378_v0)
				var _445_v41 _dafny.Array
				_ = _445_v41
				var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
				_ = _nw65
				_445_v41 = _nw65
				var _446_v42 _dafny.Map
				_ = _446_v42
				_446_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _378_v0)
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_445_v41), 0))
				_ = _index47
				(_445_v41).ArraySet1(_446_v42, (_index47).Int())
				var _447_v43 _dafny.Sequence
				_ = _447_v43
				_447_v43 = _dafny.UnicodeSeqOfUtf8Bytes("bxjwsbdy")
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_445_v41), 0))
				_ = _index48
				var _rhs52 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F24(), _this.F33)
				_ = _rhs52
				var _rhs53 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_447_v43, _dafny.UnicodeSeqOfUtf8Bytes("b")), _dafny.Companion_Sequence_.Concatenate(_447_v43, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-744))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_448_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_449_i3 _dafny.Int) _dafny.CodePoint {
						return _448_v4
					}
				})(_382_v4)))))
				_ = _rhs53
				var _rhs54 _dafny.Map = ((_446_v42).Update((_441_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))).Int()).(bool), (_441_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))).Int()).(bool))).Update((_441_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))).Int()).(bool), _378_v0)
				_ = _rhs54
				var _rhs55 _dafny.Int = (func(_pat_let2_0 D0) D0 {
					return func(_450_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let3_0 _dafny.Int) D0 {
							return func(_451_dt__update_hcf1_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC1_(_451_dt__update_hcf1_h0, (_450_dt__update__tmp_h1).Dtor_cf2(), (_450_dt__update__tmp_h1).Dtor_cf3())
							}(_pat_let3_0)
						}(_this.F33)
					}(_pat_let2_0)
				}(Companion_Default___.Fm21(globalState))).Dtor_cf2()
				_ = _rhs55
				var _lhs38 *GlobalState = globalState
				_ = _lhs38
				var _lhs39 _dafny.Array = _445_v41
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_445_v41), 0))
				_ = _lhs40
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				_lhs38.F18 = _rhs52
				_378_v0 = _rhs53
				(_lhs39).ArraySet1(_rhs54, (_lhs40).Int())
				_lhs41.F1 = _rhs55
				var _452_v44 _dafny.Int
				_ = _452_v44
				var _453_v45 bool
				_ = _453_v45
				var _454_v46 _dafny.Int
				_ = _454_v46
				var _455_v47 _dafny.CodePoint
				_ = _455_v47
				var _out18 _dafny.Int
				_ = _out18
				var _out19 bool
				_ = _out19
				var _out20 _dafny.Int
				_ = _out20
				var _out21 _dafny.CodePoint
				_ = _out21
				_out18, _out19, _out20, _out21 = (_this).M8(_441_v37, globalState)
				_452_v44 = _out18
				_453_v45 = _out19
				_454_v46 = _out20
				_455_v47 = _out21
			} else {
				var _456_v48 _dafny.Array
				_ = _456_v48
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw66
				_456_v48 = _nw66
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_456_v48), 0))
				_ = _index49
				(_456_v48).ArraySet1((_this).F24(), (_index49).Int())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_456_v48), 0))
				_ = _index50
				var _rhs56 _dafny.Int = _this.F33
				_ = _rhs56
				var _rhs57 bool = _378_v0
				_ = _rhs57
				var _lhs42 _dafny.Array = _456_v48
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_456_v48), 0))
				_ = _lhs43
				(_lhs42).ArraySet1(_rhs56, (_lhs43).Int())
				r2 = _rhs57
				r0 = _378_v0
				var _457_v49 _dafny.Map
				_ = _457_v49
				_457_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v6, (_441_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))).Int()).(bool))
				var _458_v50 _dafny.Array
				_ = _458_v50
				var _nwElement0_9 _dafny.Map = _457_v49
				_ = _nwElement0_9
				var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(7))
				_ = _nw67
				(_nw67).ArraySet1(_nwElement0_9, 0)
				(_nw67).ArraySet1((_457_v49).Merge(_457_v49), 1)
				(_nw67).ArraySet1(_457_v49, 2)
				(_nw67).ArraySet1((_457_v49).Update(_384_v6, (_441_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))).Int()).(bool)), 3)
				(_nw67).ArraySet1((_457_v49).Merge(_457_v49), 4)
				(_nw67).ArraySet1(_457_v49, 5)
				(_nw67).ArraySet1(_457_v49, 6)
				_458_v50 = _nw67
				_458_v50 = _458_v50
				var _459_v51 *C0
				_ = _459_v51
				var _nw68 *C0 = New_C0_()
				_ = _nw68
				_nw68.Ctor__(_378_v0)
				_459_v51 = _nw68
				var _460_v52 _dafny.Sequence
				_ = _460_v52
				_460_v52 = _dafny.SeqOf(_382_v4)
				var _461_v53 D1
				_ = _461_v53
				_461_v53 = Companion_D1_.Create_DC3_(_460_v52)
				var _462_v54 _dafny.Set
				_ = _462_v54
				_462_v54 = _dafny.SetOf(_461_v53, _461_v53, Companion_D1_.Create_DC3_(_460_v52), _461_v53, Companion_D1_.Create_DC3_(_460_v52))
				_459_v51 = (func() *C0 {
					if !((_462_v54).IsSubsetOf(_462_v54)) {
						return _459_v51
					}
					return _459_v51
				})()
				(globalState).F18 = ((func() _dafny.Int {
					if (_441_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_441_v37), 0))).Int()).(bool) {
						return _this.F33
					}
					return (_this).F24()
				})()).Plus(_dafny.IntOfInt64(214))
			}
		} else {
			var _463_v55 _dafny.CodePoint
			_ = _463_v55
			_463_v55 = _dafny.CodePoint('r')
			var _464_v56 _dafny.Sequence
			_ = _464_v56
			_464_v56 = _dafny.SeqOf(_463_v55)
			var _465_v57 _dafny.Set
			_ = _465_v57
			_465_v57 = _dafny.SetOf(_464_v56, _464_v56)
			_465_v57 = _465_v57
			var _rhs58 bool = _378_v0
			_ = _rhs58
			var _rhs59 _dafny.Sequence = p0
			_ = _rhs59
			var _rhs60 _dafny.Int = (_dafny.Zero).Minus(_this.F33)
			_ = _rhs60
			var _rhs61 _dafny.Set = (Companion_Default___.Fm22(_this.F33, _378_v0, _378_v0, Companion_Default___.Fm0(_this.F33, (_this).F24(), globalState), globalState)).Intersection((_465_v57).Union(_dafny.SetOf(_464_v56, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-411))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}((func(_466_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_467_i4 _dafny.Int) _dafny.CodePoint {
					return _466_v55
				}
			})(_463_v55))))))
			_ = _rhs61
			var _rhs62 _dafny.CodePoint = _463_v55
			_ = _rhs62
			var _lhs44 *GlobalState = globalState
			_ = _lhs44
			var _lhs45 *GlobalState = globalState
			_ = _lhs45
			r2 = _rhs58
			_lhs44.F21 = _rhs59
			_lhs45.F18 = _rhs60
			_465_v57 = _rhs61
			_463_v55 = _rhs62
			if _378_v0 {
				var _468_v58 _dafny.MultiSet
				_ = _468_v58
				_468_v58 = _dafny.MultiSetOf(_this.F33, _this.F33)
				var _469_v59 _dafny.Array
				_ = _469_v59
				var _nwElement0_10 _dafny.Sequence = p0
				_ = _nwElement0_10
				var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(13))
				_ = _nw69
				(_nw69).ArraySet1(_nwElement0_10, 0)
				(_nw69).ArraySet1(p0, 1)
				(_nw69).ArraySet1(p0, 2)
				(_nw69).ArraySet1(p0, 3)
				(_nw69).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), 4)
				(_nw69).ArraySet1(_dafny.SeqOf(_378_v0, false, _378_v0), 5)
				(_nw69).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqOf(_378_v0)), 6)
				(_nw69).ArraySet1(p0, 7)
				(_nw69).ArraySet1(p0, 8)
				(_nw69).ArraySet1(Companion_Default___.Fm23(globalState), 9)
				(_nw69).ArraySet1(Companion_Default___.Fm23(globalState), 10)
				(_nw69).ArraySet1(p0, 11)
				(_nw69).ArraySet1(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_468_v58).Contains(_dafny.IntOfUint32((p0).Cardinality())) {
						return (_468_v58).Multiplicity(_dafny.IntOfUint32((p0).Cardinality()))
					}
					return _this.F33
				})(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _378_v0), 12)
				_469_v59 = _nw69
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_469_v59), 0))
				_ = _index51
				(_469_v59).ArraySet1(p0, (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_469_v59), 0))
				_ = _index52
				(_469_v59).ArraySet1(p0, (_index52).Int())
				var _470_v60 _dafny.Sequence
				_ = _470_v60
				_470_v60 = _dafny.SeqOf(_464_v56)
				(globalState).F2 = Companion_Default___.Fm2(Companion_Default___.Fm0(_dafny.IntOfUint32(((_470_v60).Select((Companion_Default___.SafeIndex(_this.F33, _dafny.IntOfUint32((_470_v60).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (Companion_D3_.Create_DC11_(_this.F33, ((_469_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_469_v59), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32(((_469_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_469_v59), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(bool), _378_v0)).Dtor_cf18(), globalState), _this.F33, Companion_Default___.Fm3(globalState), globalState)
				var _471_v61 D4
				_ = _471_v61
				_471_v61 = Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-751))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_472_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_473_i5 _dafny.Int) _dafny.CodePoint {
						return _472_v55
					}
				})(_463_v55)))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F33), _464_v56)
				var _474_v62 _dafny.Map
				_ = _474_v62
				_474_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_471_v61).Dtor_cf22(), _dafny.IntOfInt64(-176))
				_474_v62 = (_474_v62).Update(_this.F33, (func() _dafny.Int {
					if _378_v0 {
						return _dafny.IntOfUint32((_464_v56).Cardinality())
					}
					return (_this).F24()
				})())
				var _475_v63 _dafny.Array
				_ = _475_v63
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw70
				_475_v63 = _nw70
				var _476_v64 _dafny.Map
				_ = _476_v64
				_476_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F33, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F33, _378_v0))
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_475_v63), 0))
				_ = _index53
				(_475_v63).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F33, (_476_v64).Cardinality()), (_index53).Int())
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_475_v63), 0))
				_ = _index54
				(_475_v63).ArraySet1(_this.F33, (_index54).Int())
				var _477_v65 *C0
				_ = _477_v65
				var _nw71 *C0 = New_C0_()
				_ = _nw71
				_nw71.Ctor__(_378_v0)
				_477_v65 = _nw71
			} else {
				var _478_v67 _dafny.Map
				_ = _478_v67
				_478_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, _this.F33)
				var _479_v68 _dafny.MultiSet
				_ = _479_v68
				_479_v68 = _dafny.MultiSetOf(func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter18 := _dafny.Iterate((_380_v2).Elements()); ; {
						_compr_17, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _480_v66 _dafny.Int
						_480_v66 = interface{}(_compr_17).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_380_v2, _480_v66) {
							_coll17.Add(Companion_Default___.SafeDivisionInt(_480_v66, _this.F33), _378_v0)
						}
					}
					return _coll17.ToMap()
				}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _378_v0), _379_v1, (_379_v1).Update((func() _dafny.Int {
					if (_478_v67).Contains(true) {
						return (_478_v67).Get(true).(_dafny.Int)
					}
					return _this.F33
				})(), _378_v0), _379_v1)
				var _481_v69 D4
				_ = _481_v69
				_481_v69 = Companion_D4_.Create_DC12_(_380_v2)
				var _482_v70 D3
				_ = _482_v70
				_482_v70 = Companion_D3_.Create_DC11_(_this.F33, _378_v0, _378_v0)
				var _483_v71 _dafny.MultiSet
				_ = _483_v71
				_483_v71 = _dafny.MultiSetOf(_378_v0, (_482_v70).Dtor_cf19(), _378_v0)
				var _484_v72 _dafny.Array
				_ = _484_v72
				var _nwElement0_11 _dafny.Int = _this.F33
				_ = _nwElement0_11
				var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(17))
				_ = _nw72
				(_nw72).ArraySet1(_nwElement0_11, 0)
				(_nw72).ArraySet1((_this).F24(), 1)
				(_nw72).ArraySet1(((_dafny.Zero).Minus(_this.F33)).Times(_this.F33), 2)
				(_nw72).ArraySet1(_this.F33, 3)
				(_nw72).ArraySet1((func() _dafny.Int {
					if Companion_Default___.Fm2(_this.F33, (func() _dafny.Int {
						if (_479_v68).Contains(_379_v1) {
							return (_479_v68).Multiplicity(_379_v1)
						}
						return _this.F33
					})(), _463_v55, globalState) {
						return _dafny.IntOfUint32(((_481_v69).Dtor_cf21()).Cardinality())
					}
					return _this.F33
				})(), 4)
				(_nw72).ArraySet1(((_this).F24()).Times(_this.F33), 5)
				(_nw72).ArraySet1((_483_v71).Cardinality(), 6)
				(_nw72).ArraySet1(_this.F33, 7)
				(_nw72).ArraySet1(_dafny.IntOfInt64(100), 8)
				(_nw72).ArraySet1(_this.F33, 9)
				(_nw72).ArraySet1(_this.F33, 10)
				(_nw72).ArraySet1((_this).F24(), 11)
				(_nw72).ArraySet1(_dafny.IntOfUint32((_464_v56).Cardinality()), 12)
				(_nw72).ArraySet1(_this.F33, 13)
				(_nw72).ArraySet1(_dafny.IntOfInt64(856), 14)
				(_nw72).ArraySet1((_dafny.IntOfUint32((_464_v56).Cardinality())).Plus((_this).F24()), 15)
				(_nw72).ArraySet1(_dafny.IntOfInt64(946), 16)
				_484_v72 = _nw72
				_484_v72 = _484_v72
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_484_v72), 0))
				_ = _index55
				(_484_v72).ArraySet1(_this.F33, (_index55).Int())
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_484_v72), 0))
				_ = _index56
				(_484_v72).ArraySet1((_this.F33).Plus(_dafny.IntOfInt64(802)), (_index56).Int())
				var _485_v73 _dafny.Array
				_ = _485_v73
				var _nw73 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw73
				_485_v73 = _nw73
				var _486_v74 _dafny.Set
				_ = _486_v74
				_486_v74 = _dafny.SetOf(_485_v73)
				var _487_v75 D1
				_ = _487_v75
				_487_v75 = Companion_D1_.Create_DC5_(_378_v0, (_486_v74).IsSubsetOf(_486_v74), _378_v0)
				_487_v75 = _487_v75
				var _488_v76 _dafny.Array
				_ = _488_v76
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(9))
				_ = _nw74
				_488_v76 = _nw74
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_488_v76), 0))
				_ = _index57
				(_488_v76).ArraySet1(_481_v69, (_index57).Int())
				var _489_v77 D1
				_ = _489_v77
				_489_v77 = Companion_D1_.Create_DC4_(_378_v0, _378_v0, _463_v55, _this.F33, (_this).F24())
				var _pat_let_tv1 = _380_v2
				_ = _pat_let_tv1
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_488_v76), 0))
				_ = _index58
				var _rhs63 D4 = func(_pat_let4_0 D4) D4 {
					return func(_490_dt__update__tmp_h2 D4) D4 {
						return func(_pat_let5_0 _dafny.Sequence) D4 {
							return func(_491_dt__update_hcf21_h0 _dafny.Sequence) D4 {
								return Companion_D4_.Create_DC12_(_491_dt__update_hcf21_h0)
							}(_pat_let5_0)
						}(_pat_let_tv1)
					}(_pat_let4_0)
				}(_481_v69)
				_ = _rhs63
				var _rhs64 _dafny.Map = (func() _dafny.Map {
					if (_489_v77).Dtor_cf9() {
						return (_478_v67).Merge(_478_v67)
					}
					return _478_v67
				})()
				_ = _rhs64
				var _lhs46 _dafny.Array = _488_v76
				_ = _lhs46
				var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_488_v76), 0))
				_ = _lhs47
				(_lhs46).ArraySet1(_rhs63, (_lhs47).Int())
				_478_v67 = _rhs64
				var _492_v78 _dafny.Map
				_ = _492_v78
				_492_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, Companion_Default___.Fm24((_380_v2).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_380_v2).Cardinality()))).Uint32()).(_dafny.Int), !(_378_v0), (_484_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_484_v72), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_this).F24()), globalState))
				_492_v78 = (_492_v78).Update(_378_v0, _380_v2)
			}
			var _493_v79 _dafny.Array
			_ = _493_v79
			var _nwElement0_12 _dafny.CodePoint = _463_v55
			_ = _nwElement0_12
			var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(7))
			_ = _nw75
			(_nw75).ArraySet1CodePoint(_nwElement0_12, 0)
			(_nw75).ArraySet1CodePoint(_463_v55, 1)
			(_nw75).ArraySet1CodePoint(_463_v55, 2)
			(_nw75).ArraySet1CodePoint(_dafny.CodePoint('p'), 3)
			(_nw75).ArraySet1CodePoint((_464_v56).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_464_v56).Cardinality()))).Uint32()).(_dafny.CodePoint), 4)
			(_nw75).ArraySet1CodePoint(_463_v55, 5)
			(_nw75).ArraySet1CodePoint(_463_v55, 6)
			_493_v79 = _nw75
			var _494_v80 _dafny.Sequence
			_ = _494_v80
			_494_v80 = _dafny.SeqOf(_493_v79)
			var _495_v81 D2
			_ = _495_v81
			_495_v81 = Companion_D2_.Create_DC7_((_494_v80).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.IntOfUint32((_494_v80).Cardinality()))).Uint32()).(_dafny.Array))
			var _source5 D2 = _495_v81
			_ = _source5
			if _source5.Is_DC8() {
				(globalState).F18 = (_this.F33).Plus((func() _dafny.Int {
					if _378_v0 {
						return (_this).F24()
					}
					return (_this).F24()
				})())
				var _496_v82 _dafny.Array
				_ = _496_v82
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw76
				_496_v82 = _nw76
				_496_v82 = _496_v82
				var _497_v83 *C0
				_ = _497_v83
				var _nw77 *C0 = New_C0_()
				_ = _nw77
				_nw77.Ctor__((func() bool {
					if _378_v0 {
						return _378_v0
					}
					return _378_v0
				})())
				_497_v83 = _nw77
				_378_v0 = ((_this).F24()).Cmp((_this).F24()) < 0
			} else {
				var _498___mcc_h9 _dafny.Array = _source5.Get_().(D2_DC7).Cf16
				_ = _498___mcc_h9
				var _499_cf16 _dafny.Array = _498___mcc_h9
				_ = _499_cf16
				var _500_v84 D1
				_ = _500_v84
				_500_v84 = Companion_D1_.Create_DC5_(_378_v0, _378_v0, true)
				var _pat_let_tv2 = _378_v0
				_ = _pat_let_tv2
				var _501_v85 _dafny.Map
				_ = _501_v85
				_501_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_380_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F33), _dafny.IntOfUint32((_380_v2).Cardinality()))).Uint32()).(_dafny.Int), func(_pat_let6_0 D1) D1 {
					return func(_502_dt__update__tmp_h3 D1) D1 {
						return func(_pat_let7_0 bool) D1 {
							return func(_503_dt__update_hcf13_h0 bool) D1 {
								return Companion_D1_.Create_DC5_(_503_dt__update_hcf13_h0, (_502_dt__update__tmp_h3).Dtor_cf14(), (_502_dt__update__tmp_h3).Dtor_cf15())
							}(_pat_let7_0)
						}(_pat_let_tv2)
					}(_pat_let6_0)
				}(_500_v84))
				var _504_v86 _dafny.Map
				_ = _504_v86
				_504_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F33, _501_v85)
				_504_v86 = func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter19 := _dafny.Iterate((Companion_Default___.Fm18(globalState)).Elements()); ; {
						_compr_18, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _505_v87 _dafny.Int
						_505_v87 = interface{}(_compr_18).(_dafny.Int)
						if (Companion_Default___.Fm18(globalState)).Contains(_505_v87) {
							_coll18.Add((_505_v87).Times(_this.F33), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F33)), _500_v84)).Merge(_501_v85))
						}
					}
					return _coll18.ToMap()
				}()
				var _506_v88 *C0
				_ = _506_v88
				var _nw78 *C0 = New_C0_()
				_ = _nw78
				_nw78.Ctor__(_378_v0)
				_506_v88 = _nw78
				(globalState).F17 = _464_v56
				var _507_v89 _dafny.Map
				_ = _507_v89
				_507_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_464_v56, !(_378_v0))
				_507_v89 = (_507_v89).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(901))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_508_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_509_i6 _dafny.Int) _dafny.CodePoint {
						return _508_v55
					}
				})(_463_v55))), ((_506_v88).F34()) == ((_506_v88).F34()))
			}
			(globalState).F14 = (_dafny.Zero).Minus((_this).F24())
		}
		var _510_v90 _dafny.Sequence
		_ = _510_v90
		_510_v90 = _dafny.UnicodeSeqOfUtf8Bytes("frfeqaonm")
		var _511_v91 D4
		_ = _511_v91
		_511_v91 = Companion_D4_.Create_DC13_(_this.F33, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, _this.F33), _510_v90)
		var _source6 D4 = _511_v91
		_ = _source6
		if _source6.Is_DC13() {
			var _512___mcc_h10 _dafny.Int = _source6.Get_().(D4_DC13).Cf22
			_ = _512___mcc_h10
			var _513___mcc_h11 _dafny.Map = _source6.Get_().(D4_DC13).Cf23
			_ = _513___mcc_h11
			var _514___mcc_h12 _dafny.Sequence = _source6.Get_().(D4_DC13).Cf24
			_ = _514___mcc_h12
			var _515_cf24 _dafny.Sequence = _514___mcc_h12
			_ = _515_cf24
			var _516_cf23 _dafny.Map = _513___mcc_h11
			_ = _516_cf23
			var _517_cf22 _dafny.Int = _512___mcc_h10
			_ = _517_cf22
			(globalState).F2 = _378_v0
			r1 = (_dafny.Zero).Minus(_this.F33)
			if _378_v0 {
				var _518_v92 _dafny.CodePoint
				_ = _518_v92
				_518_v92 = _dafny.CodePoint('x')
				var _rhs65 bool = Companion_Default___.Fm2(_this.F33, _this.F33, _518_v92, globalState)
				_ = _rhs65
				var _rhs66 _dafny.Int = (_this).Fm11(_378_v0, _378_v0, ((_this).F24()).Cmp(_this.F33) != 0, (Companion_Default___.Fm25(globalState)).Dtor_cf22(), globalState)
				_ = _rhs66
				var _rhs67 bool = _378_v0
				_ = _rhs67
				var _lhs48 *C1 = _this
				_ = _lhs48
				var _lhs49 *GlobalState = globalState
				_ = _lhs49
				_378_v0 = _rhs65
				_lhs48.F33 = _rhs66
				_lhs49.F2 = _rhs67
				_379_v1 = (_379_v1).Update(_517_cf22, true)
				(globalState).F2 = false
				r2 = (_378_v0) == (!(_378_v0))
				var _519_v93 _dafny.Array
				_ = _519_v93
				var _nw79 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw79
				_519_v93 = _nw79
				var _520_v94 _dafny.Set
				_ = _520_v94
				_520_v94 = _dafny.SetOf(_519_v93, _519_v93, _519_v93, _519_v93, _519_v93)
				_520_v94 = (_dafny.SetOf(_519_v93, _519_v93, _519_v93)).Intersection(_520_v94)
			} else {
				var _521_v95 _dafny.CodePoint
				_ = _521_v95
				_521_v95 = _dafny.CodePoint('l')
				var _522_v96 _dafny.MultiSet
				_ = _522_v96
				_522_v96 = _dafny.MultiSetOf((_379_v1).Merge(_379_v1))
				var _rhs68 _dafny.CodePoint = _521_v95
				_ = _rhs68
				var _rhs69 _dafny.MultiSet = _522_v96
				_ = _rhs69
				_521_v95 = _rhs68
				_522_v96 = _rhs69
				(globalState).F2 = ((_dafny.Zero).Minus((_this).F24())).Cmp(_this.F33) < 0
				var _523_v97 _dafny.Map
				_ = _523_v97
				_523_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_515_cf24), false)
				var _524_v98 D1
				_ = _524_v98
				_524_v98 = Companion_D1_.Create_DC3_(_510_v90)
				_523_v97 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_524_v98, _378_v0)).Merge(_523_v97)
				_378_v0 = (_this).Fm10(globalState)
				(_this).F33 = _this.F33
			}
			var _525_v99 _dafny.CodePoint
			_ = _525_v99
			_525_v99 = _dafny.CodePoint('n')
			_525_v99 = _525_v99
		} else if _source6.Is_DC12() {
			var _526___mcc_h13 _dafny.Sequence = _source6.Get_().(D4_DC12).Cf21
			_ = _526___mcc_h13
			var _527_cf21 _dafny.Sequence = _526___mcc_h13
			_ = _527_cf21
			(globalState).F2 = _378_v0
			(globalState).F2 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm23(globalState), p0), p0)
			(globalState).F18 = ((_this).F24()).Minus(_this.F33)
			var _528_v100 _dafny.Array
			_ = _528_v100
			var _nwElement0_13 bool = false
			_ = _nwElement0_13
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(7))
			_ = _nw80
			(_nw80).ArraySet1(_nwElement0_13, 0)
			(_nw80).ArraySet1(_378_v0, 1)
			(_nw80).ArraySet1((_378_v0) == (_378_v0), 2)
			(_nw80).ArraySet1(_378_v0, 3)
			(_nw80).ArraySet1(_378_v0, 4)
			(_nw80).ArraySet1(!(_dafny.Companion_Sequence_.Equal(_510_v90, _dafny.UnicodeSeqOfUtf8Bytes("lruyipny"))), 5)
			(_nw80).ArraySet1(true, 6)
			_528_v100 = _nw80
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_528_v100), 0))
			_ = _index59
			(_528_v100).ArraySet1(_378_v0, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_528_v100), 0))
			_ = _index60
			(_528_v100).ArraySet1(_378_v0, (_index60).Int())
		} else {
			var _529___mcc_h14 D4 = _source6.Get_().(D4_DC14).Cf25
			_ = _529___mcc_h14
			var _530_cf25 D4 = _529___mcc_h14
			_ = _530_cf25
			var _531_v101 _dafny.Array
			_ = _531_v101
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
			_ = _nw81
			_531_v101 = _nw81
			var _532_v102 _dafny.Map
			_ = _532_v102
			_532_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, _378_v0)
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_531_v101), 0))
			_ = _index61
			(_531_v101).ArraySet1(_532_v102, (_index61).Int())
			var _533_v103 _dafny.MultiSet
			_ = _533_v103
			_533_v103 = _dafny.MultiSetOf((_this).F24(), _this.F33)
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_531_v101), 0))
			_ = _index62
			var _rhs70 _dafny.Int = (_dafny.Zero).Minus((_this).F24())
			_ = _rhs70
			var _rhs71 _dafny.Map = _532_v102
			_ = _rhs71
			var _rhs72 _dafny.MultiSet = ((_533_v103).Intersection(_533_v103)).Difference(_533_v103)
			_ = _rhs72
			var _lhs50 *GlobalState = globalState
			_ = _lhs50
			var _lhs51 _dafny.Array = _531_v101
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen((_531_v101), 0))
			_ = _lhs52
			_lhs50.F9 = _rhs70
			(_lhs51).ArraySet1(_rhs71, (_lhs52).Int())
			_533_v103 = _rhs72
			var _534_v104 _dafny.MultiSet
			_ = _534_v104
			_534_v104 = _dafny.MultiSetOf(Companion_Default___.Fm26(globalState))
			var _535_v105 _dafny.Array
			_ = _535_v105
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_12
			var _nw82 _dafny.Array
			_ = _nw82
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw82 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) bool = (func(_536_v0 bool) func(_dafny.Int) bool {
					return func(_537_i7 _dafny.Int) bool {
						return _536_v0
					}
				})(_378_v0)
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
			_535_v105 = _nw82
			var _538_v106 _dafny.Map
			_ = _538_v106
			_538_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_535_v105, (_this).F24())
			var _539_v107 D4
			_ = _539_v107
			_539_v107 = Companion_D4_.Create_DC13_((func() _dafny.Int {
				if (_538_v106).Contains(_535_v105) {
					return (_538_v106).Get(_535_v105).(_dafny.Int)
				}
				return (_dafny.MultiSetOf(_this.F33)).Cardinality()
			})(), Companion_Default___.Fm17((_this).F24(), _dafny.IntOfUint32((_380_v2).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_378_v0, _378_v0)).Cardinality(), true, globalState), _510_v90)
			var _540_v108 D4
			_ = _540_v108
			_540_v108 = Companion_D4_.Create_DC14_(_539_v107)
			var _541_v109 D4
			_ = _541_v109
			_541_v109 = Companion_D4_.Create_DC14_(_539_v107)
			var _542_v110 D4
			_ = _542_v110
			_542_v110 = Companion_D4_.Create_DC14_(_540_v108)
			var _543_v111 _dafny.Map
			_ = _543_v111
			_543_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(97), _dafny.IntOfInt64(-620))
			var _544_v112 _dafny.CodePoint
			_ = _544_v112
			_544_v112 = _dafny.CodePoint('y')
			var _545_v113 _dafny.Array
			_ = _545_v113
			var _nwElement0_14 _dafny.Int = ((_this).F24()).Plus((_380_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_534_v104).Update(_542_v110, Companion_Default___.Abs(_this.F33))).Cardinality()), _dafny.IntOfUint32((_380_v2).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _nwElement0_14
			var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(12))
			_ = _nw83
			(_nw83).ArraySet1(_nwElement0_14, 0)
			(_nw83).ArraySet1(_this.F33, 1)
			(_nw83).ArraySet1((_this).F24(), 2)
			(_nw83).ArraySet1((_dafny.Zero).Minus((_this.F33).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bgboi")).Cardinality()))), 3)
			(_nw83).ArraySet1(((_533_v103).Cardinality()).Minus((_543_v111).Cardinality()), 4)
			(_nw83).ArraySet1((_this).F24(), 5)
			(_nw83).ArraySet1((_379_v1).Cardinality(), 6)
			(_nw83).ArraySet1((_this).F24(), 7)
			(_nw83).ArraySet1(_this.F33, 8)
			(_nw83).ArraySet1((_this).F24(), 9)
			(_nw83).ArraySet1((_this).F24(), 10)
			(_nw83).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_510_v90, (Companion_Default___.SafeIndex(_this.F33, _dafny.IntOfUint32((_510_v90).Cardinality()))).Uint32(), _544_v112)).Cardinality()), 11)
			_545_v113 = _nw83
			_545_v113 = _545_v113
			var _546_v114 D1
			_ = _546_v114
			_546_v114 = Companion_D1_.Create_DC6_()
			_546_v114 = Companion_Default___.Fm27(globalState)
			var _547_v115 _dafny.Map
			_ = _547_v115
			_547_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_544_v112, (_this).F24())
			var _548_v116 _dafny.Map
			_ = _548_v116
			_548_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_547_v115, (_this).F24())
			var _549_v117 _dafny.Set
			_ = _549_v117
			_549_v117 = _dafny.SetOf(_378_v0, _378_v0)
			var _550_v118 _dafny.Sequence
			_ = _550_v118
			_550_v118 = _dafny.SeqOf((_548_v116).Update(_547_v115, (_this).F24()), _548_v116, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_547_v115, (_549_v117).Cardinality()))
			_548_v116 = ((_550_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.IntOfUint32((_550_v118).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_548_v116)
		}
		(globalState).F14 = (_this).F24()
		(globalState).F14 = _this.F33
		r0 = _378_v0
		r1 = (_this).F24()
		var _551_v119 _dafny.Set
		_ = _551_v119
		_551_v119 = _dafny.SetOf(_this.F33)
		r2 = (func() _dafny.Set {
			var _coll19 = _dafny.NewBuilder()
			_ = _coll19
			for _iter20 := _dafny.Iterate((_380_v2).Elements()); ; {
				_compr_19, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _552_v120 _dafny.Int
				_552_v120 = interface{}(_compr_19).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_380_v2, _552_v120) {
					_coll19.Add(Companion_Default___.SafeDivisionInt(_552_v120, (_dafny.SetOf(_dafny.SetOf(true, true))).Cardinality()))
				}
			}
			return _coll19.ToSet()
		}()).IsSubsetOf(_551_v119)
		return r0, r1, r2
	}
}
func (_this *C1) M8(p0 _dafny.Array, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int, _dafny.CodePoint) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r3
		var _553_v0 bool
		_ = _553_v0
		_553_v0 = false
		r1 = (func() bool {
			if _553_v0 {
				return _553_v0
			}
			return (func() bool {
				if _553_v0 {
					return _553_v0
				}
				return _553_v0
			})()
		})()
		var _554_v1 _dafny.Map
		_ = _554_v1
		_554_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(319), (_dafny.Zero).Minus((_this).F24()))
		var _555_v2 _dafny.Sequence
		_ = _555_v2
		_555_v2 = _dafny.UnicodeSeqOfUtf8Bytes("nedubbppa")
		var _556_v3 _dafny.Array
		_ = _556_v3
		var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw84
		_556_v3 = _nw84
		var _557_v4 _dafny.Set
		_ = _557_v4
		_557_v4 = _dafny.SetOf(_556_v3, _556_v3, _556_v3)
		var _558_v5 _dafny.Set
		_ = _558_v5
		_558_v5 = _dafny.SetOf(_dafny.IntOfUint32((_555_v2).Cardinality()))
		var _559_v6 D6
		_ = _559_v6
		_559_v6 = Companion_D6_.Create_DC16_(_558_v5)
		var _560_v7 _dafny.Sequence
		_ = _560_v7
		_560_v7 = _dafny.SeqOf(((_559_v6).Dtor_cf27()).Cardinality())
		var _561_v8 _dafny.Sequence
		_ = _561_v8
		_561_v8 = _dafny.SeqOf(_553_v0, _553_v0)
		var _562_v9 _dafny.Array
		_ = _562_v9
		var _nwElement0_15 _dafny.Int = (_this).F24()
		_ = _nwElement0_15
		var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(26))
		_ = _nw85
		(_nw85).ArraySet1(_nwElement0_15, 0)
		(_nw85).ArraySet1((_this).F24(), 1)
		(_nw85).ArraySet1((_this).F24(), 2)
		(_nw85).ArraySet1((func() _dafny.Int {
			if (_554_v1).Contains(_this.F33) {
				return (_554_v1).Get(_this.F33).(_dafny.Int)
			}
			return _this.F33
		})(), 3)
		(_nw85).ArraySet1((_this).F24(), 4)
		(_nw85).ArraySet1((_this).F24(), 5)
		(_nw85).ArraySet1(_dafny.IntOfUint32((_555_v2).Cardinality()), 6)
		(_nw85).ArraySet1(_dafny.IntOfUint32((_555_v2).Cardinality()), 7)
		(_nw85).ArraySet1((_this).F24(), 8)
		(_nw85).ArraySet1((_557_v4).Cardinality(), 9)
		(_nw85).ArraySet1(Companion_Default___.Fm0((_this).F24(), _dafny.IntOfInt64(580), globalState), 10)
		(_nw85).ArraySet1(_this.F33, 11)
		(_nw85).ArraySet1((_this).F24(), 12)
		(_nw85).ArraySet1((_this).F24(), 13)
		(_nw85).ArraySet1(_this.F33, 14)
		(_nw85).ArraySet1(_dafny.IntOfUint32((_560_v7).Cardinality()), 15)
		(_nw85).ArraySet1(_this.F33, 16)
		(_nw85).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_561_v8).Cardinality())), 17)
		(_nw85).ArraySet1(_dafny.IntOfInt64(-528), 18)
		(_nw85).ArraySet1(_this.F33, 19)
		(_nw85).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_560_v7).Cardinality())), 20)
		(_nw85).ArraySet1((_this).F24(), 21)
		(_nw85).ArraySet1((_this).F24(), 22)
		(_nw85).ArraySet1(_this.F33, 23)
		(_nw85).ArraySet1(_dafny.IntOfInt64(108), 24)
		(_nw85).ArraySet1(_this.F33, 25)
		_562_v9 = _nw85
		var _563_v10 _dafny.MultiSet
		_ = _563_v10
		_563_v10 = _dafny.MultiSetOf(_562_v9)
		if !((_563_v10).Equals(_563_v10)) {
			var _564_v11 _dafny.Map
			_ = _564_v11
			_564_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_553_v0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_561_v8, _561_v8)).Cardinality()))
			_564_v11 = (_564_v11).Update(_553_v0, (_this).F24())
			var _565_v12 _dafny.Array
			_ = _565_v12
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
			_ = _nw86
			_565_v12 = _nw86
			var _566_v13 D4
			_ = _566_v13
			_566_v13 = Companion_D4_.Create_DC14_(Companion_D4_.Create_DC13_(_this.F33, _564_v11, _555_v2))
			var _567_v14 _dafny.Map
			_ = _567_v14
			_567_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_566_v13, _555_v2)
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_565_v12), 0))
			_ = _index63
			(_565_v12).ArraySet1(_567_v14, (_index63).Int())
			var _568_v15 D7
			_ = _568_v15
			_568_v15 = Companion_D7_.Create_DC20_(_567_v14)
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_565_v12), 0))
			_ = _index64
			(_565_v12).ArraySet1((_568_v15).Dtor_cf31(), (_index64).Int())
			(globalState).F18 = _this.F33
			(globalState).F2 = _553_v0
			var _569_v16 _dafny.CodePoint
			_ = _569_v16
			_569_v16 = _dafny.CodePoint('u')
			var _570_v17 _dafny.Map
			_ = _570_v17
			_570_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v0, _553_v0)
			var _571_v18 _dafny.MultiSet
			_ = _571_v18
			_571_v18 = _dafny.MultiSetOf((_570_v17).Cardinality(), (_this).F24())
			var _572_v19 _dafny.MultiSet
			_ = _572_v19
			_572_v19 = _dafny.MultiSetOf(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_569_v16, !(_553_v0))).Cardinality()).Cmp((_571_v18).Cardinality()) == 0)
			_572_v19 = Companion_Default___.Fm14(_554_v1, _this.F33, globalState)
		} else {
			(globalState).F2 = !_dafny.Companion_Sequence_.Equal(_561_v8, _561_v8)
			if ((_this).F24()).Cmp(Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState)) != 0 {
				var _573_v20 _dafny.Set
				_ = _573_v20
				_573_v20 = _dafny.SetOf(_553_v0, !(_553_v0), _553_v0, _553_v0, _553_v0)
				var _574_v21 _dafny.Set
				_ = _574_v21
				_574_v21 = _dafny.SetOf(_dafny.SeqOf(_dafny.IntOfUint32((_555_v2).Cardinality()), (_573_v20).Cardinality(), _this.F33))
				var _575_v22 _dafny.Array
				_ = _575_v22
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw87
				_575_v22 = _nw87
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_575_v22), 0))
				_ = _index65
				(_575_v22).ArraySet1(Companion_Default___.Fm24(_this.F33, _553_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dn")).Cardinality()), _this.F33, globalState), (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((p0), 0))
				_ = _index66
				(p0).ArraySet1(((_558_v5).Cardinality()).Cmp((_this).F24()) == 0, (_index66).Int())
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_575_v22), 0))
				_ = _index67
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((p0), 0))
				_ = _index68
				var _rhs73 _dafny.Sequence = Companion_Default___.Fm24((_dafny.Zero).Minus(_this.F33), _553_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_555_v2, _555_v2)).Cardinality()), _dafny.IntOfInt64(151), globalState)
				_ = _rhs73
				var _rhs74 _dafny.Set = _574_v21
				_ = _rhs74
				var _rhs75 _dafny.Int = _this.F33
				_ = _rhs75
				var _rhs76 _dafny.Sequence = _dafny.SeqOf((_this).F24(), _dafny.IntOfInt64(829), Companion_Default___.SafeModuloInt(_this.F33, _dafny.IntOfUint32((_555_v2).Cardinality())))
				_ = _rhs76
				var _rhs77 bool = _553_v0
				_ = _rhs77
				var _lhs53 *C1 = _this
				_ = _lhs53
				var _lhs54 _dafny.Array = _575_v22
				_ = _lhs54
				var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_575_v22), 0))
				_ = _lhs55
				var _lhs56 _dafny.Array = p0
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((p0), 0))
				_ = _lhs57
				_560_v7 = _rhs73
				_574_v21 = _rhs74
				_lhs53.F33 = _rhs75
				(_lhs54).ArraySet1(_rhs76, (_lhs55).Int())
				(_lhs56).ArraySet1(_rhs77, (_lhs57).Int())
				var _576_v23 *C0
				_ = _576_v23
				var _nw88 *C0 = New_C0_()
				_ = _nw88
				_nw88.Ctor__((_this.F33).Cmp(_this.F33) <= 0)
				_576_v23 = _nw88
				r1 = (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((p0), 0))).Int()).(bool)
				var _577_v24 _dafny.Map
				_ = _577_v24
				_577_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((_575_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_575_v22), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_this.F33, _dafny.IntOfUint32(((_575_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_575_v22), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _this.F33), Companion_D6_.Create_DC17_(_553_v0))
				var _578_v25 D6
				_ = _578_v25
				_578_v25 = Companion_D6_.Create_DC17_(!((_576_v23).F34()))
				_577_v24 = (_577_v24).Update(_560_v7, (func() D6 {
					if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
						return _578_v25
					}
					return _578_v25
				})())
				var _579_v26 *C0
				_ = _579_v26
				var _nw89 *C0 = New_C0_()
				_ = _nw89
				_nw89.Ctor__((_576_v23).F34())
				_579_v26 = _nw89
			} else {
				var _580_v27 _dafny.Set
				_ = _580_v27
				_580_v27 = _dafny.SetOf(_553_v0)
				var _rhs78 bool = _553_v0
				_ = _rhs78
				var _rhs79 _dafny.Set = (_dafny.SetOf(_553_v0)).Difference(_580_v27)
				_ = _rhs79
				var _rhs80 bool = false
				_ = _rhs80
				var _lhs58 *GlobalState = globalState
				_ = _lhs58
				_553_v0 = _rhs78
				_580_v27 = _rhs79
				_lhs58.F2 = _rhs80
				r1 = _553_v0
				var _581_v28 D1
				_ = _581_v28
				_581_v28 = Companion_D1_.Create_DC5_(_553_v0, _553_v0, _553_v0)
				var _582_v29 D6
				_ = _582_v29
				_582_v29 = Companion_D6_.Create_DC17_(_553_v0)
				var _pat_let_tv3 = _553_v0
				_ = _pat_let_tv3
				var _583_v30 D7
				_ = _583_v30
				_583_v30 = Companion_D7_.Create_DC21_(func(_pat_let8_0 D1) D1 {
					return func(_584_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let9_0 bool) D1 {
							return func(_585_dt__update_hcf15_h0 bool) D1 {
								return Companion_D1_.Create_DC5_((_584_dt__update__tmp_h0).Dtor_cf13(), (_584_dt__update__tmp_h0).Dtor_cf14(), _585_dt__update_hcf15_h0)
							}(_pat_let9_0)
						}(_pat_let_tv3)
					}(_pat_let8_0)
				}(_581_v28), Companion_D6_.Create_DC19_(_582_v29), p0)
				var _586_v31 _dafny.Map
				_ = _586_v31
				_586_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v0, _583_v30)
				var _587_v32 _dafny.MultiSet
				_ = _587_v32
				_587_v32 = _dafny.MultiSetOf(_553_v0)
				var _588_v33 _dafny.Map
				_ = _588_v33
				_588_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v0, true)
				var _rhs81 _dafny.Int = ((_dafny.IntOfUint32((_561_v8).Cardinality())).Minus(_this.F33)).Minus((func() _dafny.Int {
					if (_587_v32).Contains(_553_v0) {
						return (_587_v32).Multiplicity(_553_v0)
					}
					return _this.F33
				})())
				_ = _rhs81
				var _rhs82 _dafny.Int = _dafny.IntOfUint32((_561_v8).Cardinality())
				_ = _rhs82
				var _rhs83 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v0, _583_v30)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_553_v0), _583_v30))
				_ = _rhs83
				var _rhs84 bool = ((_this).F24()).Cmp((_dafny.Zero).Minus(((_588_v33).Cardinality()).Minus((_this).F24()))) >= 0
				_ = _rhs84
				var _lhs59 *GlobalState = globalState
				_ = _lhs59
				var _lhs60 *GlobalState = globalState
				_ = _lhs60
				var _lhs61 *GlobalState = globalState
				_ = _lhs61
				_lhs59.F9 = _rhs81
				_lhs60.F18 = _rhs82
				_586_v31 = _rhs83
				_lhs61.F2 = _rhs84
				var _589_v34 _dafny.Map
				_ = _589_v34
				_589_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_561_v8, (_this).F24())
				_589_v34 = (_589_v34).Update(_dafny.Companion_Sequence_.Concatenate(_561_v8, _561_v8), _dafny.IntOfInt64(997))
				(globalState).F1 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-374))
			}
			_562_v9 = _562_v9
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((p0), 0))
			_ = _index69
			(p0).ArraySet1((_553_v0) || (_553_v0), (_index69).Int())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((p0), 0))
			_ = _index70
			(p0).ArraySet1(!(!(_553_v0)), (_index70).Int())
			var _590_v35 _dafny.Set
			_ = _590_v35
			_590_v35 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(841))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_591_v2 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
				return func(_592_i0 _dafny.Int) _dafny.CodePoint {
					return (_591_v2).Select((Companion_Default___.SafeIndex(_this.F33, _dafny.IntOfUint32((_591_v2).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_555_v2))))
			var _593_v36 *C0
			_ = _593_v36
			var _nw90 *C0 = New_C0_()
			_ = _nw90
			_nw90.Ctor__(!((_590_v35).IsDisjointFrom(_590_v35)))
			_593_v36 = _nw90
		}
		(globalState).F2 = (_557_v4).IsDisjointFrom(_557_v4)
		var _594_v37 D1
		_ = _594_v37
		_594_v37 = Companion_D1_.Create_DC5_(_553_v0, true, _553_v0)
		var _595_v38 D6
		_ = _595_v38
		_595_v38 = Companion_D6_.Create_DC16_(_dafny.SetOf(_this.F33))
		var _596_v39 D7
		_ = _596_v39
		_596_v39 = Companion_D7_.Create_DC21_(_594_v37, Companion_D6_.Create_DC19_(_595_v38), p0)
		var _597_v40 D7
		_ = _597_v40
		_597_v40 = Companion_D7_.Create_DC22_(_596_v39)
		var _598_v41 _dafny.Map
		_ = _598_v41
		_598_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_597_v40, p0)
		_598_v41 = (_598_v41).Update(Companion_D7_.Create_DC22_(_596_v39), p0)
		var _599_v42 _dafny.MultiSet
		_ = _599_v42
		_599_v42 = _dafny.MultiSetOf((_this).F24(), (_dafny.Zero).Minus((_558_v5).Cardinality()))
		var _600_v43 _dafny.Set
		_ = _600_v43
		_600_v43 = _dafny.SetOf(((_599_v42).Cardinality()).Cmp(_dafny.IntOfInt64(-939)) < 0, _553_v0, (func() bool {
			if !(_553_v0) {
				return _553_v0
			}
			return _553_v0
		})(), _553_v0, _553_v0)
		var _601_v44 _dafny.MultiSet
		_ = _601_v44
		_601_v44 = _dafny.MultiSetOf((Companion_Default___.Fm28(_553_v0, _dafny.IntOfUint32((_555_v2).Cardinality()), _553_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v0, _this.F33)).Cardinality(), globalState)).IsProperSubsetOf(_600_v43), _553_v0, _553_v0)
		var _602_v45 D8
		_ = _602_v45
		_602_v45 = Companion_D8_.Create_DC23_(_dafny.SetOf(_553_v0, _553_v0, _553_v0, !(_553_v0)))
		var _rhs85 _dafny.Set = (_602_v45).Dtor_cf36()
		_ = _rhs85
		var _rhs86 _dafny.Int = ((_dafny.IntOfInt64(830)).Plus(_this.F33)).Plus(_dafny.IntOfInt64(689))
		_ = _rhs86
		var _rhs87 _dafny.MultiSet = (_601_v44).Update(_553_v0, Companion_Default___.Abs((_this).F24()))
		_ = _rhs87
		var _lhs62 *GlobalState = globalState
		_ = _lhs62
		_600_v43 = _rhs85
		_lhs62.F18 = _rhs86
		_601_v44 = _rhs87
		var _hi3 _dafny.Int = ((_this).F24()).Times(_dafny.IntOfUint32((_555_v2).Cardinality()))
		_ = _hi3
		for _603_i1 := (func() _dafny.Int {
			if _553_v0 {
				return (_this).F24()
			}
			return _this.F33
		})(); _603_i1.Cmp(_hi3) < 0; _603_i1 = _603_i1.Plus(_dafny.One) {
			_555_v2 = _555_v2
			r1 = ((_this).F24()).Cmp((_dafny.Zero).Minus((Companion_Default___.Fm0((_this).F24(), _dafny.IntOfInt64(179), globalState)).Minus(_this.F33))) < 0
			if _553_v0 {
				var _604_v46 _dafny.Map
				_ = _604_v46
				_604_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
				var _rhs88 _dafny.Int = Companion_Default___.SafeModuloInt((_604_v46).Cardinality(), (_603_i1).Plus(_this.F33))
				_ = _rhs88
				var _rhs89 bool = !((_559_v6).Dtor_cf27()).Equals(func() _dafny.Set {
					var _coll20 = _dafny.NewBuilder()
					_ = _coll20
					for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(-841))); ; {
						_compr_20, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _605_v47 _dafny.Int
						_605_v47 = interface{}(_compr_20).(_dafny.Int)
						if ((_dafny.IntOfInt64(280)).Cmp(_605_v47) <= 0) && ((_605_v47).Cmp(_dafny.IntOfInt64(-841)) < 0) {
							_coll20.Add(Companion_Default___.SafeModuloInt(_605_v47, (_601_v44).Cardinality()))
						}
					}
					return _coll20.ToSet()
				}())
				_ = _rhs89
				var _lhs63 *GlobalState = globalState
				_ = _lhs63
				_lhs63.F18 = _rhs88
				r1 = _rhs89
				var _606_v48 _dafny.Array
				_ = _606_v48
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_13
				var _nw91 _dafny.Array
				_ = _nw91
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw91 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Sequence = (func(_607_v2 _dafny.Sequence, _608_i1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_609_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Update(_607_v2, (Companion_Default___.SafeIndex(_608_i1, _dafny.IntOfUint32((_607_v2).Cardinality()))).Uint32(), _dafny.CodePoint('v'))
						}
					})(_555_v2, _603_i1)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw91 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw91).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw91).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_606_v48 = _nw91
				_606_v48 = _606_v48
				var _610_v49 _dafny.Int
				_ = _610_v49
				var _out22 _dafny.Int
				_ = _out22
				_out22 = Companion_Default___.M0(_this.F33, ((_this).F24()).Minus(_dafny.IntOfInt64(-447)), globalState)
				_610_v49 = _out22
				var _611_v50 D6
				_ = _611_v50
				_611_v50 = Companion_D6_.Create_DC17_(_553_v0)
				var _612_v51 _dafny.CodePoint
				_ = _612_v51
				_612_v51 = _dafny.CodePoint('d')
				var _613_v52 _dafny.Map
				_ = _613_v52
				_613_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_555_v2).Cardinality()), Companion_Default___.Fm4(_553_v0, _610_v49, _612_v51, _553_v0, globalState))
				var _614_v53 _dafny.Map
				_ = _614_v53
				_614_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_v50, _613_v52)
				var _615_v54 _dafny.Map
				_ = _615_v54
				_615_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_610_v49, _613_v52)
				_614_v53 = (_614_v53).Update(_611_v50, ((func() _dafny.Map {
					if (_615_v54).Contains(_610_v49) {
						return (_615_v54).Get(_610_v49).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_610_v49, _555_v2)
				})()).Merge(_613_v52))
				var _rhs90 bool = _553_v0
				_ = _rhs90
				var _rhs91 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(212))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_616_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_617_i3 _dafny.Int) _dafny.CodePoint {
						return _616_v51
					}
				})(_612_v51))), _555_v2)
				_ = _rhs91
				var _lhs64 *GlobalState = globalState
				_ = _lhs64
				var _lhs65 *GlobalState = globalState
				_ = _lhs65
				_lhs64.F2 = _rhs90
				_lhs65.F17 = _rhs91
			} else {
				(globalState).F21 = _561_v8
				var _618_v55 _dafny.Map
				_ = _618_v55
				_618_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_560_v7, (Companion_Default___.SafeIndex(_603_i1, _dafny.IntOfUint32((_560_v7).Cardinality()))).Uint32(), _603_i1)).Cardinality()))
				var _619_v56 _dafny.Map
				_ = _619_v56
				_619_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v0, (_618_v55).Cardinality())
				r0 = ((func() _dafny.Int {
					if (_619_v56).Contains(_553_v0) {
						return (_619_v56).Get(_553_v0).(_dafny.Int)
					}
					return _dafny.IntOfInt64(679)
				})()).Minus(Companion_Default___.SafeDivisionInt(_603_i1, _this.F33))
				var _620_v57 *C0
				_ = _620_v57
				var _nw92 *C0 = New_C0_()
				_ = _nw92
				_nw92.Ctor__(_553_v0)
				_620_v57 = _nw92
				var _621_v58 _dafny.Map
				_ = _621_v58
				_621_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}(func(_622_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))).Cardinality())), (_620_v57).F34())
				_621_v58 = _621_v58
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((p0), 0))
				_ = _index71
				(p0).ArraySet1(false, (_index71).Int())
				var _623_v59 _dafny.CodePoint
				_ = _623_v59
				_623_v59 = _dafny.CodePoint('s')
				var _624_v60 _dafny.Map
				_ = _624_v60
				_624_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_562_v9, _623_v59)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((p0), 0))
				_ = _index72
				var _rhs92 bool = !(((_dafny.Zero).Minus(_603_i1)).Cmp((func() _dafny.Int {
					if true {
						return _dafny.IntOfInt64(701)
					}
					return _603_i1
				})()) > 0)
				_ = _rhs92
				var _rhs93 _dafny.Int = (_560_v7).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(102)).Plus(_dafny.IntOfInt64(97)), _dafny.IntOfUint32((_560_v7).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs93
				var _rhs94 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_556_v3, _623_v59)).Merge((_624_v60).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_556_v3, _623_v59)))
				_ = _rhs94
				var _lhs66 _dafny.Array = p0
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((p0), 0))
				_ = _lhs67
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				(_lhs66).ArraySet1(_rhs92, (_lhs67).Int())
				_lhs68.F9 = _rhs93
				_624_v60 = _rhs94
			}
			r1 = (_dafny.IntOfInt64(134)).Cmp(_603_i1) == 0
		}
		r0 = (_this.F33).Times((Companion_Default___.Fm0((_this).F24(), _this.F33, globalState)).Times(_dafny.IntOfInt64(610)))
		r1 = _553_v0
		var _625_v61 _dafny.CodePoint
		_ = _625_v61
		_625_v61 = _dafny.CodePoint('f')
		var _626_v62 _dafny.Map
		_ = _626_v62
		_626_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_601_v44).Cardinality(), false)
		r2 = (_560_v7).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0((_dafny.Zero).Minus((Companion_D1_.Create_DC4_(_553_v0, _553_v0, _625_v61, _dafny.IntOfUint32((_560_v7).Cardinality()), Companion_Default___.Fm0(_this.F33, Companion_Default___.Fm0((_626_v62).Cardinality(), (_this).F24(), globalState), globalState))).Dtor_cf12()), _this.F33, globalState), _dafny.IntOfUint32((_560_v7).Cardinality()))).Uint32()).(_dafny.Int)
		r3 = (_555_v2).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(-930)).Minus(_dafny.IntOfInt64(447)), _dafny.IntOfUint32((_555_v2).Cardinality()))).Uint32()).(_dafny.CodePoint)
		return r0, r1, r2, r3
	}
}
func (_this *C1) F32() _dafny.Sequence {
	{
		return _this._f32
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f24 _dafny.Int
	_f25 _dafny.Sequence
	F31  bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f24 = _dafny.Zero
	_this._f25 = _dafny.EmptySeq
	_this.F31 = false
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

func (_this *C2) F24() _dafny.Int {
	return _this._f24
}
func (_this *C2) F25() _dafny.Sequence {
	return _this._f25
}
func (_this *C2) Ctor__(f31 bool, f24 _dafny.Int, f25 _dafny.Sequence) {
	{
		(_this).F31 = f31
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C2) Fm9(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf(Companion_D1_.Create_DC6_())).IsSubsetOf(_dafny.MultiSetOf(Companion_D1_.Create_DC6_(), Companion_D1_.Create_DC6_()))
	}
}
func (_this *C2) M1(globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _627_v0 D0
		_ = _627_v0
		_627_v0 = Companion_D0_.Create_DC1_((_this).F24(), (_this).F24(), _this.F31)
		if func(_source7 D0) bool {
			if _source7.Is_DC1() {
				var _628___mcc_h4 _dafny.Int = _source7.Get_().(D0_DC1).Cf1
				_ = _628___mcc_h4
				var _629___mcc_h5 _dafny.Int = _source7.Get_().(D0_DC1).Cf2
				_ = _629___mcc_h5
				var _630___mcc_h6 bool = _source7.Get_().(D0_DC1).Cf3
				_ = _630___mcc_h6
				var _631_cf3 bool = _630___mcc_h6
				_ = _631_cf3
				var _632_cf2 _dafny.Int = _629___mcc_h5
				_ = _632_cf2
				var _633_cf1 _dafny.Int = _628___mcc_h4
				_ = _633_cf1
				return _631_cf3
			} else if _source7.Is_DC2() {
				var _634___mcc_h7 bool = _source7.Get_().(D0_DC2).Cf4
				_ = _634___mcc_h7
				var _635___mcc_h8 _dafny.Int = _source7.Get_().(D0_DC2).Cf5
				_ = _635___mcc_h8
				var _636___mcc_h9 _dafny.CodePoint = _source7.Get_().(D0_DC2).Cf6
				_ = _636___mcc_h9
				var _637_cf6 _dafny.CodePoint = _636___mcc_h9
				_ = _637_cf6
				var _638_cf5 _dafny.Int = _635___mcc_h8
				_ = _638_cf5
				var _639_cf4 bool = _634___mcc_h7
				_ = _639_cf4
				return !(_this.F31)
			} else {
				var _640___mcc_h10 _dafny.Array = _source7.Get_().(D0_DC0).Cf0
				_ = _640___mcc_h10
				var _641_cf0 _dafny.Array = _640___mcc_h10
				_ = _641_cf0
				return _this.F31
			}
		}(_627_v0) {
			var _642_v1 _dafny.Array
			_ = _642_v1
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_14
			var _nw93 _dafny.Array
			_ = _nw93
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw93 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = func(_643_i0 _dafny.Int) bool {
					return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-591))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}(func(_644_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('x')
					}))).Cardinality())).Cmp((_this).F24()) != 0
				}
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw93 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw93).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw93).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_642_v1 = _nw93
			var _rhs95 _dafny.Array = _642_v1
			_ = _rhs95
			var _rhs96 bool = (_this.F31) == ((_this.F31) && (_this.F31))
			_ = _rhs96
			var _lhs69 *GlobalState = globalState
			_ = _lhs69
			_642_v1 = _rhs95
			_lhs69.F2 = _rhs96
			if _this.F31 {
				var _645_v2 _dafny.Sequence
				_ = _645_v2
				_645_v2 = _dafny.UnicodeSeqOfUtf8Bytes("dvbaaxfn")
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))
				_ = _index73
				(_642_v1).ArraySet1(((_this).F24()).Cmp(_dafny.IntOfUint32((_645_v2).Cardinality())) <= 0, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))
				_ = _index74
				(_642_v1).ArraySet1(_this.F31, (_index74).Int())
				var _646_v3 _dafny.Map
				_ = _646_v3
				_646_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(69))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}(func(_647_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				}))).Cardinality()))
				(globalState).F9 = ((_this).F24()).Minus(Companion_Default___.SafeDivisionInt((_646_v3).Cardinality(), (_this).F24()))
				(globalState).F18 = (_this).F24()
				var _648_v4 _dafny.MultiSet
				_ = _648_v4
				_648_v4 = _dafny.MultiSetOf((_642_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))).Int()).(bool), (_642_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))).Int()).(bool), _this.F31)
				var _649_v5 _dafny.CodePoint
				_ = _649_v5
				_649_v5 = _dafny.CodePoint('o')
				var _650_v6 _dafny.Array
				_ = _650_v6
				var _nwElement0_16 _dafny.CodePoint = _649_v5
				_ = _nwElement0_16
				var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.One)
				_ = _nw94
				(_nw94).ArraySet1CodePoint(_nwElement0_16, 0)
				_650_v6 = _nw94
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_650_v6), 0))
				_ = _index75
				(_650_v6).ArraySet1CodePoint((func() _dafny.CodePoint {
					if !(_this.F31) {
						return Companion_Default___.Fm3(globalState)
					}
					return _dafny.CodePoint('t')
				})(), (_index75).Int())
				var _651_v7 D0
				_ = _651_v7
				_651_v7 = Companion_D0_.Create_DC2_(_this.F31, (_this).F24(), _649_v5)
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))
				_ = _index76
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_650_v6), 0))
				_ = _index77
				var _rhs97 bool = !((_this).Fm9((_642_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))).Int()).(bool), Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState), (_this).F24(), globalState))
				_ = _rhs97
				var _rhs98 _dafny.MultiSet = _648_v4
				_ = _rhs98
				var _rhs99 _dafny.CodePoint = (_651_v7).Dtor_cf6()
				_ = _rhs99
				var _lhs70 _dafny.Array = _642_v1
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))
				_ = _lhs71
				var _lhs72 _dafny.Array = _650_v6
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_650_v6), 0))
				_ = _lhs73
				(_lhs70).ArraySet1(_rhs97, (_lhs71).Int())
				_648_v4 = _rhs98
				(_lhs72).ArraySet1CodePoint(_rhs99, (_lhs73).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))
				_ = _index78
				(_642_v1).ArraySet1(((_642_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_642_v1), 0))).Int()).(bool)) || (_this.F31), (_index78).Int())
			} else {
				r0 = _this.F31
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_642_v1), 0))
				_ = _index79
				(_642_v1).ArraySet1((_dafny.IntOfInt64(155)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-796))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}(func(_652_i3 _dafny.Int) _dafny.Int {
					return (_this).F24()
				}))).Cardinality())) <= 0, (_index79).Int())
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_642_v1), 0))
				_ = _index80
				(_642_v1).ArraySet1(_this.F31, (_index80).Int())
				(globalState).F2 = (_this).Fm9(_this.F31, (_this).F24(), (_this).F24(), globalState)
				var _653_v8 _dafny.Sequence
				_ = _653_v8
				_653_v8 = _dafny.UnicodeSeqOfUtf8Bytes("xuqmdjtcc")
				var _654_v9 _dafny.CodePoint
				_ = _654_v9
				_654_v9 = _dafny.CodePoint('e')
				var _655_v10 D0
				_ = _655_v10
				_655_v10 = Companion_D0_.Create_DC2_((_642_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_642_v1), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qoctswai")).Cardinality()), _654_v9)
				var _656_v11 _dafny.Map
				_ = _656_v11
				_656_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_655_v10).Dtor_cf6(), (_this).F24())
				var _657_v12 *C1
				_ = _657_v12
				var _nw95 *C1 = New_C1_()
				_ = _nw95
				_nw95.Ctor__(Companion_Default___.Fm29(_dafny.IntOfInt64(743), (_this).F24(), globalState), (_this).F24(), (_this).F24(), _dafny.Companion_Sequence_.Update((_this).F25(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_653_v8).Cardinality()), _dafny.IntOfUint32(((_this).F25()).Cardinality()))).Uint32(), _656_v11))
				_657_v12 = _nw95
				var _nw96 *C1 = New_C1_()
				_ = _nw96
				_nw96.Ctor__((_657_v12).F32(), (_this).F24(), (_this).F24(), (_this).F25())
				_657_v12 = _nw96
			}
			var _658_v13 _dafny.MultiSet
			_ = _658_v13
			_658_v13 = _dafny.MultiSetOf(_642_v1, _642_v1, _642_v1)
			var _659_v14 _dafny.Array
			_ = _659_v14
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_15
			var _nw97 _dafny.Array
			_ = _nw97
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw97 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Int = func(_660_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_660_i4, (_this).F24())
				}
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
			_659_v14 = _nw97
			var _661_v15 _dafny.Map
			_ = _661_v15
			_661_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_658_v13).Difference(_658_v13), _659_v14)
			_661_v15 = (_661_v15).Update(_658_v13, _659_v14)
			var _662_v16 D1
			_ = _662_v16
			_662_v16 = Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}(func(_663_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})))
			var _664_v17 D8
			_ = _664_v17
			_664_v17 = Companion_D8_.Create_DC26_(_662_v16, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(623))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}(func(_665_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})))
			var _666_v18 _dafny.Map
			_ = _666_v18
			_666_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F24()), _664_v17)).Cardinality())
			var _667_v19 _dafny.Sequence
			_ = _667_v19
			_667_v19 = _dafny.SeqOf(_666_v18, Companion_Default___.Fm17((_this).F24(), (_this).F24(), (_this).F24(), _this.F31, globalState), _666_v18, _666_v18, _666_v18)
			_667_v19 = _667_v19
			var _668_v20 D1
			_ = _668_v20
			_668_v20 = Companion_D1_.Create_DC5_(true, _this.F31, _this.F31)
			var _669_v21 D6
			_ = _669_v21
			_669_v21 = Companion_D6_.Create_DC19_(Companion_D6_.Create_DC18_(false))
			var _670_v22 D6
			_ = _670_v22
			_670_v22 = Companion_D6_.Create_DC19_(_669_v21)
			var _671_v23 D6
			_ = _671_v23
			_671_v23 = Companion_D6_.Create_DC19_((Companion_D7_.Create_DC21_(_668_v20, _670_v22, _642_v1)).Dtor_cf33())
			var _source8 D6 = _670_v22
			_ = _source8
			if _source8.Is_DC17() {
				var _672___mcc_h0 bool = _source8.Get_().(D6_DC17).Cf28
				_ = _672___mcc_h0
				var _673_cf28 bool = _672___mcc_h0
				_ = _673_cf28
				var _674_v24 _dafny.Sequence
				_ = _674_v24
				_674_v24 = _dafny.UnicodeSeqOfUtf8Bytes("ibwh")
				(globalState).F17 = _674_v24
				var _675_v25 _dafny.Map
				_ = _675_v25
				_675_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, _673_cf28)
				var _676_v26 D6
				_ = _676_v26
				_676_v26 = Companion_D6_.Create_DC18_((func() bool {
					if (_675_v25).Contains(_673_cf28) {
						return (_675_v25).Get(_673_cf28).(bool)
					}
					return _673_cf28
				})())
				_673_cf28 = (_676_v26).Dtor_cf29()
				var _677_v27 _dafny.Sequence
				_ = _677_v27
				_677_v27 = _dafny.SeqOf((_this).F24())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_659_v14), 0))
				_ = _index81
				(_659_v14).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_677_v27).Cardinality()), (_this).F24()), (_index81).Int())
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_659_v14), 0))
				_ = _index82
				(_659_v14).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F24(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_674_v24).Cardinality())))), (_index82).Int())
				var _678_v28 _dafny.Sequence
				_ = _678_v28
				_678_v28 = _dafny.SeqOf(_673_cf28, _this.F31, _673_cf28)
				var _679_v29 _dafny.MultiSet
				_ = _679_v29
				_679_v29 = _dafny.MultiSetOf((_659_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_659_v14), 0))).Int()).(_dafny.Int))
				var _680_v30 _dafny.Sequence
				_ = _680_v30
				_680_v30 = _dafny.SeqOf(_dafny.MultiSetOf((_this).F24()), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_678_v28).Cardinality()))), _dafny.MultiSetOf((_675_v25).Cardinality()), (_679_v29).Update(_dafny.IntOfUint32((_677_v27).Cardinality()), Companion_Default___.Abs((_659_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_659_v14), 0))).Int()).(_dafny.Int))))
				var _681_v31 *C1
				_ = _681_v31
				var _nw98 *C1 = New_C1_()
				_ = _nw98
				_nw98.Ctor__(_680_v30, _dafny.IntOfInt64(-808), (_this).F24(), (_this).F25())
				_681_v31 = _nw98
			} else if _source8.Is_DC18() {
				var _682___mcc_h1 bool = _source8.Get_().(D6_DC18).Cf29
				_ = _682___mcc_h1
				var _683_cf29 bool = _682___mcc_h1
				_ = _683_cf29
				var _684_v32 _dafny.CodePoint
				_ = _684_v32
				_684_v32 = _dafny.CodePoint('d')
				var _685_v33 _dafny.Sequence
				_ = _685_v33
				_685_v33 = _dafny.SeqOf(_this.F31, _683_cf29)
				var _686_v34 _dafny.Map
				_ = _686_v34
				_686_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_684_v32, _dafny.Companion_Sequence_.Concatenate(_685_v33, _685_v33))
				var _687_v35 _dafny.MultiSet
				_ = _687_v35
				_687_v35 = _dafny.MultiSetOf(_683_cf29)
				var _688_v36 _dafny.MultiSet
				_ = _688_v36
				_688_v36 = _687_v35
				var _rhs100 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _685_v33)
				_ = _rhs100
				var _rhs101 _dafny.MultiSet = _687_v35
				_ = _rhs101
				var _rhs102 bool = false
				_ = _rhs102
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				_686_v34 = _rhs100
				_688_v36 = _rhs101
				_lhs74.F2 = _rhs102
				_659_v14 = _659_v14
				_659_v14 = _659_v14
				_684_v32 = _684_v32
			} else if _source8.Is_DC16() {
				var _689___mcc_h2 _dafny.Set = _source8.Get_().(D6_DC16).Cf27
				_ = _689___mcc_h2
				var _690_cf27 _dafny.Set = _689___mcc_h2
				_ = _690_cf27
				var _691_v37 _dafny.Set
				_ = _691_v37
				_691_v37 = _dafny.SetOf(_this.F31)
				var _692_v38 _dafny.Map
				_ = _692_v38
				_692_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F24()), _691_v37)
				_692_v38 = (_692_v38).Update(_dafny.IntOfInt64(37), _691_v37)
				var _693_v39 _dafny.Sequence
				_ = _693_v39
				_693_v39 = _dafny.UnicodeSeqOfUtf8Bytes("h")
				var _694_v40 *C0
				_ = _694_v40
				var _nw99 *C0 = New_C0_()
				_ = _nw99
				_nw99.Ctor__((_dafny.IntOfUint32((_693_v39).Cardinality())).Cmp((_dafny.Zero).Minus((_this).F24())) > 0)
				_694_v40 = _nw99
				var _695_v41 _dafny.Sequence
				_ = _695_v41
				_695_v41 = _dafny.SeqOf((_694_v40).F34())
				var _696_v42 _dafny.CodePoint
				_ = _696_v42
				_696_v42 = _dafny.CodePoint('g')
				var _697_v43 D2
				_ = _697_v43
				_697_v43 = Companion_D2_.Create_DC8_()
				var _698_v44 _dafny.MultiSet
				_ = _698_v44
				_698_v44 = _dafny.MultiSetOf(_697_v43)
				var _699_v45 _dafny.MultiSet
				_ = _699_v45
				_699_v45 = _dafny.MultiSetOf((_this).F24(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_694_v40).F34(), _696_v42)).Cardinality(), (_dafny.Zero).Minus((_698_v44).Cardinality()))
				(globalState).F21 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if ((_694_v40).F34()) && (!((_694_v40).F34())) {
						return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm23(globalState), _695_v41)
					}
					return (func() _dafny.Sequence {
						if (_694_v40).F34() {
							return _695_v41
						}
						return _695_v41
					})()
				})(), (Companion_Default___.SafeIndex(((_this).F24()).Times((func() _dafny.Int {
					if (_699_v45).Contains((_this).F24()) {
						return (_699_v45).Multiplicity((_this).F24())
					}
					return (_this).F24()
				})()), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if ((_694_v40).F34()) && (!((_694_v40).F34())) {
						return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm23(globalState), _695_v41)
					}
					return (func() _dafny.Sequence {
						if (_694_v40).F34() {
							return _695_v41
						}
						return _695_v41
					})()
				})()).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.IsPrefixOf(_693_v39, _693_v39))
				_693_v39 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(663))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_700_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_701_i7 _dafny.Int) _dafny.CodePoint {
						return _700_v42
					}
				})(_696_v42)))
			} else {
				var _702___mcc_h3 D6 = _source8.Get_().(D6_DC19).Cf30
				_ = _702___mcc_h3
				var _703_cf30 D6 = _702___mcc_h3
				_ = _703_cf30
				(globalState).F2 = _this.F31
				(_this).F31 = _this.F31
				var _704_v46 _dafny.Sequence
				_ = _704_v46
				_704_v46 = _dafny.UnicodeSeqOfUtf8Bytes("ea")
				var _rhs103 _dafny.Sequence = _704_v46
				_ = _rhs103
				var _rhs104 _dafny.Sequence = (func() _dafny.Sequence {
					if _this.F31 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(180))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg53 _dafny.Int) interface{} {
								return coer53(arg53)
							}
						}(func(_705_i8 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('p')
						}))
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("hdlwah")
				})()
				_ = _rhs104
				var _rhs105 _dafny.Sequence = _704_v46
				_ = _rhs105
				var _rhs106 bool = (_dafny.IntOfInt64(-216)).Cmp((_this).F24()) == 0
				_ = _rhs106
				var _lhs75 *GlobalState = globalState
				_ = _lhs75
				var _lhs76 *GlobalState = globalState
				_ = _lhs76
				var _lhs77 *GlobalState = globalState
				_ = _lhs77
				_lhs75.F17 = _rhs103
				_lhs76.F17 = _rhs104
				_lhs77.F17 = _rhs105
				r0 = _rhs106
				var _706_v47 D8
				_ = _706_v47
				_706_v47 = Companion_D8_.Create_DC24_((_this).F24())
				var _707_v48 _dafny.CodePoint
				_ = _707_v48
				_707_v48 = _dafny.CodePoint('f')
				var _708_v49 _dafny.Map
				_ = _708_v49
				_708_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_707_v48, _707_v48)
				var _709_v51 _dafny.MultiSet
				_ = _709_v51
				_709_v51 = _dafny.MultiSetOf((func() _dafny.Set {
					var _coll21 = _dafny.NewBuilder()
					_ = _coll21
					for _iter22 := _dafny.Iterate((_708_v49).Keys().Elements()); ; {
						_compr_21, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _710_v50 _dafny.CodePoint
						_710_v50 = interface{}(_compr_21).(_dafny.CodePoint)
						if (_708_v49).Contains(_710_v50) {
							_coll21.Add(_710_v50)
						}
					}
					return _coll21.ToSet()
				}()).Cardinality(), (_this).F24())
				var _711_v52 _dafny.Set
				_ = _711_v52
				_711_v52 = _dafny.SetOf(_this.F31, true, _this.F31)
				_706_v47 = Companion_D8_.Create_DC24_((func() _dafny.Int {
					if (_709_v51).Contains((_this).F24()) {
						return (_709_v51).Multiplicity((_this).F24())
					}
					return (_711_v52).Cardinality()
				})())
			}
		} else {
			var _712_v53 _dafny.Map
			_ = _712_v53
			_712_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F24()).Minus((_this).F24()), (_this).F24())
			(globalState).F22 = _712_v53
			var _713_v54 _dafny.Sequence
			_ = _713_v54
			_713_v54 = _dafny.UnicodeSeqOfUtf8Bytes("o")
			var _714_v55 _dafny.Sequence
			_ = _714_v55
			_714_v55 = _dafny.SeqOf((_this).F24(), _dafny.IntOfUint32((_713_v54).Cardinality()), _dafny.IntOfInt64(920))
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_714_v55, _dafny.SeqOf((_this).F24(), (_this).F24(), (_dafny.Zero).Minus((_this).F24()), (_this).F24(), (_this).F24())) {
				var _715_v56 _dafny.Map
				_ = _715_v56
				_715_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, _dafny.IntOfInt64(170))
				var _716_v57 D4
				_ = _716_v57
				_716_v57 = Companion_D4_.Create_DC13_(((_this).F24()).Minus(_dafny.IntOfInt64(511)), (_715_v56).Merge(_715_v56), _713_v54)
				_716_v57 = _716_v57
				(globalState).F18 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F24()))
				_712_v53 = (func() _dafny.Map {
					if true {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
					}
					return _712_v53
				})()
				var _717_v58 _dafny.MultiSet
				_ = _717_v58
				_717_v58 = _dafny.MultiSetOf(true, _this.F31)
				var _718_v59 _dafny.CodePoint
				_ = _718_v59
				_718_v59 = _dafny.CodePoint('t')
				var _719_v60 _dafny.Map
				_ = _719_v60
				_719_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _713_v54)
				var _720_v61 *C0
				_ = _720_v61
				var _nw100 *C0 = New_C0_()
				_ = _nw100
				_nw100.Ctor__(_this.F31)
				_720_v61 = _nw100
				var _721_v62 _dafny.Sequence
				_ = _721_v62
				_721_v62 = _dafny.SeqOf(_720_v61, _720_v61)
				var _722_v63 _dafny.Set
				_ = _722_v63
				_722_v63 = _dafny.SetOf((_this).F24(), _dafny.IntOfUint32((_721_v62).Cardinality()))
				var _723_v64 _dafny.Array
				_ = _723_v64
				var _nwElement0_17 bool = _this.F31
				_ = _nwElement0_17
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(15))
				_ = _nw101
				(_nw101).ArraySet1(_nwElement0_17, 0)
				(_nw101).ArraySet1(!(false), 1)
				(_nw101).ArraySet1((Companion_D6_.Create_DC17_(_this.F31)).Dtor_cf28(), 2)
				(_nw101).ArraySet1(_this.F31, 3)
				(_nw101).ArraySet1((_this.F31) == (_this.F31), 4)
				(_nw101).ArraySet1((_717_v58).IsDisjointFrom(_717_v58), 5)
				(_nw101).ArraySet1(_this.F31, 6)
				(_nw101).ArraySet1(Companion_Default___.Fm2(_dafny.IntOfInt64(211), _dafny.IntOfUint32((_714_v55).Cardinality()), _718_v59, globalState), 7)
				(_nw101).ArraySet1(((_this).F24()).Cmp((_this).F24()) <= 0, 8)
				(_nw101).ArraySet1(!((_719_v60).Update((_this).F24(), _713_v54)).Contains((_this).F24()), 9)
				(_nw101).ArraySet1(!(Companion_Default___.Fm30(_718_v59, !(_this.F31), globalState)).Contains((_722_v63).Cardinality()), 10)
				(_nw101).ArraySet1((_dafny.IntOfInt64(755)).Cmp((_this).F24()) != 0, 11)
				(_nw101).ArraySet1(true, 12)
				(_nw101).ArraySet1(_this.F31, 13)
				(_nw101).ArraySet1((_720_v61).F34(), 14)
				_723_v64 = _nw101
				_723_v64 = _723_v64
				var _724_v65 _dafny.MultiSet
				_ = _724_v65
				_724_v65 = _dafny.MultiSetOf((_this).F24(), (_this).F24(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(684))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}(func(_725_i9 _dafny.Int) _dafny.Int {
					return (_this).F24()
				}))).Cardinality()), (_this).F24(), Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState))
				var _726_v66 _dafny.Set
				_ = _726_v66
				_726_v66 = _dafny.SetOf(_724_v65)
				_726_v66 = _726_v66
			} else {
				var _rhs107 bool = false
				_ = _rhs107
				var _rhs108 bool = _this.F31
				_ = _rhs108
				var _rhs109 _dafny.Int = _dafny.IntOfInt64(613)
				_ = _rhs109
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				_lhs78.F2 = _rhs107
				_lhs79.F2 = _rhs108
				_lhs80.F1 = _rhs109
				var _727_v67 _dafny.Array
				_ = _727_v67
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_16
				var _nw102 _dafny.Array
				_ = _nw102
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw102 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) bool = (func(_728_v55 _dafny.Sequence) func(_dafny.Int) bool {
						return func(_729_i10 _dafny.Int) bool {
							return _dafny.Companion_Sequence_.Equal(_728_v55, _dafny.SeqOf((_this).F24()))
						}
					})(_714_v55)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw102 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw102).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw102).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_727_v67 = _nw102
				_727_v67 = _727_v67
				var _rhs110 _dafny.Sequence = _713_v54
				_ = _rhs110
				var _rhs111 _dafny.Int = (((_this).F24()).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ljamq")).Cardinality()))).Minus(Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F24()))
				_ = _rhs111
				var _lhs81 *GlobalState = globalState
				_ = _lhs81
				var _lhs82 *GlobalState = globalState
				_ = _lhs82
				_lhs81.F17 = _rhs110
				_lhs82.F18 = _rhs111
				var _730_v68 *C0
				_ = _730_v68
				var _nw103 *C0 = New_C0_()
				_ = _nw103
				_nw103.Ctor__(_this.F31)
				_730_v68 = _nw103
				(_this).F31 = !(true) || ((_730_v68).F34())
			}
			if _this.F31 {
				var _731_v69 _dafny.Sequence
				_ = _731_v69
				_731_v69 = _dafny.SeqOf(_713_v54)
				(_this).F31 = !_dafny.Companion_Sequence_.Equal(_713_v54, _dafny.Companion_Sequence_.Concatenate((_731_v69).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_731_v69).Cardinality()))).Uint32()).(_dafny.Sequence), _713_v54))
				var _732_v70 D4
				_ = _732_v70
				_732_v70 = Companion_D4_.Create_DC12_(_714_v55)
				_714_v55 = (_732_v70).Dtor_cf21()
				var _733_v71 _dafny.Map
				_ = _733_v71
				_733_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, (_dafny.IntOfInt64(340)).Minus((_this).F24()))
				_733_v71 = (_733_v71).Update(false, (_this).F24())
				(globalState).F14 = (_714_v55).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_714_v55).Cardinality()))).Uint32()).(_dafny.Int)
				(globalState).F2 = !(_dafny.Companion_Sequence_.Equal(_714_v55, _dafny.SeqOf((_this).F24(), (_this).F24(), (_this).F24())))
			} else {
				var _734_v72 _dafny.Set
				_ = _734_v72
				_734_v72 = _dafny.SetOf((_this).F24())
				var _735_v73 _dafny.Set
				_ = _735_v73
				_735_v73 = _dafny.SetOf(_734_v72)
				(globalState).F18 = (_735_v73).Cardinality()
				var _736_v74 _dafny.MultiSet
				_ = _736_v74
				_736_v74 = _dafny.MultiSetOf((_this).F24())
				var _737_v75 T0
				_ = _737_v75
				var _nw104 *C1 = New_C1_()
				_ = _nw104
				_nw104.Ctor__(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_736_v74), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_714_v55).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_736_v74)).Cardinality()))).Uint32(), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(973))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}(func(_738_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				}))).Cardinality()))), (_this).F24(), (_this).F24(), (_this).F25())
				_737_v75 = _nw104
				var _739_v76 _dafny.Map
				_ = _739_v76
				_739_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _737_v75)
				var _740_v77 T0
				_ = _740_v77
				_740_v77 = _737_v75
				var _741_v78 _dafny.Array
				_ = _741_v78
				var _nwElement0_18 T0 = _737_v75
				_ = _nwElement0_18
				var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(19))
				_ = _nw105
				(_nw105).ArraySet1(_nwElement0_18, 0)
				(_nw105).ArraySet1(_737_v75, 1)
				(_nw105).ArraySet1(_737_v75, 2)
				(_nw105).ArraySet1(_737_v75, 3)
				(_nw105).ArraySet1((func() T0 {
					if (_739_v76).Contains(_this.F31) {
						return (_739_v76).Get(_this.F31).(T0)
					}
					return _737_v75
				})(), 4)
				(_nw105).ArraySet1(_737_v75, 5)
				(_nw105).ArraySet1(_737_v75, 6)
				(_nw105).ArraySet1(_737_v75, 7)
				(_nw105).ArraySet1(_737_v75, 8)
				(_nw105).ArraySet1(_737_v75, 9)
				(_nw105).ArraySet1(_737_v75, 10)
				(_nw105).ArraySet1(_737_v75, 11)
				(_nw105).ArraySet1(_737_v75, 12)
				(_nw105).ArraySet1(_737_v75, 13)
				(_nw105).ArraySet1(_737_v75, 14)
				(_nw105).ArraySet1(_737_v75, 15)
				(_nw105).ArraySet1(_737_v75, 16)
				(_nw105).ArraySet1(_737_v75, 17)
				(_nw105).ArraySet1((_740_v77), 18)
				_741_v78 = _nw105
				var _742_v79 _dafny.Sequence
				_ = _742_v79
				_742_v79 = _dafny.SeqOf(_737_v75)
				var _743_v80 _dafny.Map
				_ = _743_v80
				_743_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, _this.F31)
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_741_v78), 0))
				_ = _index83
				(_741_v78).ArraySet1((func() T0 {
					if _this.F31 {
						return (_742_v79).Select((Companion_Default___.SafeIndex((_743_v80).Cardinality(), _dafny.IntOfUint32((_742_v79).Cardinality()))).Uint32()).(T0)
					}
					return _737_v75
				})(), (_index83).Int())
				var _744_v81 _dafny.Map
				_ = _744_v81
				_744_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-79))), _737_v75)
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_741_v78), 0))
				_ = _index84
				(_741_v78).ArraySet1((func() T0 {
					if _this.F31 {
						return (func() T0 {
							if (_744_v81).Contains((_this).F24()) {
								return (_744_v81).Get((_this).F24()).(T0)
							}
							return _737_v75
						})()
					}
					return (func() T0 {
						if !(_this.F31) {
							return _737_v75
						}
						return _737_v75
					})()
				})(), (_index84).Int())
				(globalState).F2 = !(!(!(_this.F31)))
				(globalState).F9 = (_737_v75).F24()
				var _745_v82 _dafny.Set
				_ = _745_v82
				_745_v82 = _dafny.SetOf(_this.F31)
				var _746_v83 _dafny.Map
				_ = _746_v83
				_746_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_737_v75).F24(), _this.F31)
				var _747_v84 _dafny.CodePoint
				_ = _747_v84
				_747_v84 = _dafny.CodePoint('x')
				var _748_v85 _dafny.Set
				_ = _748_v85
				_748_v85 = _dafny.SetOf(_747_v84, _747_v84)
				var _749_v86 _dafny.Array
				_ = _749_v86
				var _nwElement0_19 _dafny.Set = _745_v82
				_ = _nwElement0_19
				var _nw106 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(12))
				_ = _nw106
				(_nw106).ArraySet1(_nwElement0_19, 0)
				(_nw106).ArraySet1(_745_v82, 1)
				(_nw106).ArraySet1(_745_v82, 2)
				(_nw106).ArraySet1((Companion_Default___.Fm28(_this.F31, (_this).F24(), _this.F31, (_dafny.Zero).Minus((_737_v75).F24()), globalState)).Union(Companion_Default___.Fm28((func() bool {
					if (_746_v83).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F24()))) {
						return (_746_v83).Get((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F24()))).(bool)
					}
					return _this.F31
				})(), (_748_v85).Cardinality(), _this.F31, (_this).F24(), globalState)), 3)
				(_nw106).ArraySet1((_745_v82).Difference(_745_v82), 4)
				(_nw106).ArraySet1(_745_v82, 5)
				(_nw106).ArraySet1(_745_v82, 6)
				(_nw106).ArraySet1(_745_v82, 7)
				(_nw106).ArraySet1(Companion_Default___.Fm28(_this.F31, (_this).F24(), _this.F31, _dafny.IntOfInt64(503), globalState), 8)
				(_nw106).ArraySet1(_dafny.SetOf(_this.F31, true, _this.F31), 9)
				(_nw106).ArraySet1((_745_v82).Difference(_745_v82), 10)
				(_nw106).ArraySet1(_745_v82, 11)
				_749_v86 = _nw106
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_749_v86), 0))
				_ = _index85
				(_749_v86).ArraySet1(_745_v82, (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_749_v86), 0))
				_ = _index86
				(_749_v86).ArraySet1(_745_v82, (_index86).Int())
			}
			(globalState).F2 = _this.F31
			(globalState).F17 = _dafny.UnicodeSeqOfUtf8Bytes("tlgqytw")
		}
		var _750_v87 _dafny.Array
		_ = _750_v87
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_17
		var _nw107 _dafny.Array
		_ = _nw107
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw107 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) _dafny.CodePoint = func(_751_i12 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw107 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw107).ArraySet1CodePoint(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw107).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_750_v87 = _nw107
		var _752_v88 _dafny.CodePoint
		_ = _752_v88
		_752_v88 = _dafny.CodePoint('y')
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))
		_ = _index87
		(_750_v87).ArraySet1CodePoint(_752_v88, (_index87).Int())
		var _753_v89 _dafny.Sequence
		_ = _753_v89
		_753_v89 = _dafny.UnicodeSeqOfUtf8Bytes("pe")
		var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))
		_ = _index88
		var _rhs112 _dafny.Int = (_this).F24()
		_ = _rhs112
		var _rhs113 _dafny.Int = (_this).F24()
		_ = _rhs113
		var _rhs114 _dafny.CodePoint = _752_v88
		_ = _rhs114
		var _rhs115 _dafny.Sequence = _753_v89
		_ = _rhs115
		var _rhs116 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F24()), Companion_Default___.SafeModuloInt((_this).F24(), _dafny.IntOfUint32((_753_v89).Cardinality())))
		_ = _rhs116
		var _lhs83 *GlobalState = globalState
		_ = _lhs83
		var _lhs84 *GlobalState = globalState
		_ = _lhs84
		var _lhs85 _dafny.Array = _750_v87
		_ = _lhs85
		var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))
		_ = _lhs86
		var _lhs87 *GlobalState = globalState
		_ = _lhs87
		var _lhs88 *GlobalState = globalState
		_ = _lhs88
		_lhs83.F9 = _rhs112
		_lhs84.F14 = _rhs113
		(_lhs85).ArraySet1CodePoint(_rhs114, (_lhs86).Int())
		_lhs87.F17 = _rhs115
		_lhs88.F9 = _rhs116
		var _hi4 _dafny.Int = (_this).F24()
		_ = _hi4
		for _754_i13 := (_this).F24(); _754_i13.Cmp(_hi4) < 0; _754_i13 = _754_i13.Plus(_dafny.One) {
			var _755_v90 _dafny.Sequence
			_ = _755_v90
			_755_v90 = _dafny.SeqOf(_this.F31, _this.F31, _this.F31, _this.F31)
			var _756_v91 _dafny.Map
			_ = _756_v91
			_756_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, (_this).Fm9(Companion_Default___.Fm2(_754_i13, _dafny.IntOfInt64(968), (_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()), globalState), _dafny.IntOfUint32((_755_v90).Cardinality()), _754_i13, globalState))
			_756_v91 = (_756_v91).Update(_this.F31, _this.F31)
			if (Companion_Default___.SafeDivisionInt((_this).F24(), _754_i13)).Cmp((Companion_Default___.Fm18(globalState)).Cardinality()) == 0 {
				var _757_v92 _dafny.Array
				_ = _757_v92
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw108
				_757_v92 = _nw108
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_757_v92), 0))
				_ = _index89
				(_757_v92).ArraySet1((_dafny.Zero).Minus(_754_i13), (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_757_v92), 0))
				_ = _index90
				(_757_v92).ArraySet1((_this).F24(), (_index90).Int())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_757_v92), 0))
				_ = _index91
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_757_v92), 0))
				_ = _index92
				var _rhs117 _dafny.Int = (_this).F24()
				_ = _rhs117
				var _rhs118 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F24(), _dafny.IntOfInt64(225))
				_ = _rhs118
				var _lhs89 _dafny.Array = _757_v92
				_ = _lhs89
				var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_757_v92), 0))
				_ = _lhs90
				var _lhs91 _dafny.Array = _757_v92
				_ = _lhs91
				var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_757_v92), 0))
				_ = _lhs92
				(_lhs89).ArraySet1(_rhs117, (_lhs90).Int())
				(_lhs91).ArraySet1(_rhs118, (_lhs92).Int())
				(globalState).F2 = (_this).Fm9(_this.F31, _754_i13, (_dafny.IntOfUint32((Companion_Default___.Fm4(_this.F31, _754_i13, _dafny.CodePoint('g'), _this.F31, globalState)).Cardinality())).Times((_this).F24()), globalState)
				var _758_v93 _dafny.Array
				_ = _758_v93
				var _nw109 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
				_ = _nw109
				_758_v93 = _nw109
				var _759_v94 _dafny.Array
				_ = _759_v94
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(13))
				_ = _nw110
				_759_v94 = _nw110
				var _760_v95 _dafny.Set
				_ = _760_v95
				_760_v95 = _dafny.SetOf((_757_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_757_v92), 0))).Int()).(_dafny.Int), _754_i13, _dafny.IntOfUint32((_753_v89).Cardinality()))
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_759_v94), 0))
				_ = _index93
				(_759_v94).ArraySet1(Companion_D6_.Create_DC16_(_760_v95), (_index93).Int())
				var _761_v96 D1
				_ = _761_v96
				_761_v96 = Companion_D1_.Create_DC5_(_this.F31, _this.F31, _this.F31)
				var _762_v97 D7
				_ = _762_v97
				_762_v97 = Companion_D7_.Create_DC21_(_761_v96, Companion_Default___.Fm31(_this.F31, _this.F31, globalState), _758_v93)
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_759_v94), 0))
				_ = _index94
				var _rhs119 _dafny.Array = (func() _dafny.Array {
					if _this.F31 {
						return _758_v93
					}
					return (_762_v97).Dtor_cf34()
				})()
				_ = _rhs119
				var _rhs120 D6 = Companion_D6_.Create_DC16_(_760_v95)
				_ = _rhs120
				var _lhs93 _dafny.Array = _759_v94
				_ = _lhs93
				var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_759_v94), 0))
				_ = _lhs94
				_758_v93 = _rhs119
				(_lhs93).ArraySet1(_rhs120, (_lhs94).Int())
				var _763_v98 _dafny.Array
				_ = _763_v98
				var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw111
				_763_v98 = _nw111
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_763_v98), 0))
				_ = _index95
				(_763_v98).ArraySet1((func() _dafny.Array {
					if _this.F31 {
						return _757_v92
					}
					return _757_v92
				})(), (_index95).Int())
				var _764_v99 D0
				_ = _764_v99
				_764_v99 = Companion_D0_.Create_DC0_(_757_v92)
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_763_v98), 0))
				_ = _index96
				(_763_v98).ArraySet1((_764_v99).Dtor_cf0(), (_index96).Int())
				(globalState).F9 = (_754_i13).Times(((_this).F24()).Minus((_757_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_757_v92), 0))).Int()).(_dafny.Int)))
			} else {
				var _765_v100 _dafny.MultiSet
				_ = _765_v100
				_765_v100 = _dafny.MultiSetOf(_754_i13, _754_i13)
				var _766_v101 _dafny.Sequence
				_ = _766_v101
				_766_v101 = _dafny.SeqOf(_765_v100)
				var _767_v102 _dafny.MultiSet
				_ = _767_v102
				_767_v102 = _dafny.MultiSetOf(_this.F31, _this.F31)
				var _768_v103 _dafny.Set
				_ = _768_v103
				_768_v103 = _dafny.SetOf(!(_this.F31))
				var _769_v104 _dafny.Map
				_ = _769_v104
				_769_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_752_v88, (_this).F24())
				var _770_v105 *C1
				_ = _770_v105
				var _nw112 *C1 = New_C1_()
				_ = _nw112
				_nw112.Ctor__(_dafny.Companion_Sequence_.Concatenate(_766_v101, _766_v101), _754_i13, Companion_Default___.SafeModuloInt(_754_i13, (func() _dafny.Int {
					if (_767_v102).Contains(_this.F31) {
						return (_767_v102).Multiplicity(_this.F31)
					}
					return (_768_v103).Cardinality()
				})()), (func() _dafny.Sequence {
					if true {
						return _dafny.SeqOf(_769_v104, _769_v104, _769_v104)
					}
					return _dafny.SeqOf(_769_v104)
				})())
				_770_v105 = _nw112
				(globalState).F2 = _this.F31
				_767_v102 = _767_v102
				var _771_v106 _dafny.Set
				_ = _771_v106
				_771_v106 = _dafny.SetOf(_754_i13)
				(globalState).F1 = Companion_Default___.SafeDivisionInt(_770_v105.F33, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_754_i13), (_771_v106).Cardinality()))
				(globalState).F18 = (_this).F24()
			}
			if ((_this).F24()).Cmp(_754_i13) != 0 {
				var _772_v107 *C0
				_ = _772_v107
				var _nw113 *C0 = New_C0_()
				_ = _nw113
				_nw113.Ctor__(_this.F31)
				_772_v107 = _nw113
				var _773_v108 _dafny.Map
				_ = _773_v108
				_773_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, _754_i13)
				(globalState).F17 = (func() _dafny.Sequence {
					if Companion_Default___.Fm2((_773_v108).Cardinality(), _dafny.IntOfInt64(392), Companion_Default___.Fm3(globalState), globalState) {
						return _753_v89
					}
					return _753_v89
				})()
				(globalState).F18 = _754_i13
				var _774_v109 _dafny.Array
				_ = _774_v109
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_18
				var _nw114 _dafny.Array
				_ = _nw114
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw114 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Sequence = (func(_775_v90 _dafny.Sequence, _776_v107 *C0) func(_dafny.Int) _dafny.Sequence {
						return func(_777_i14 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_775_v90), _dafny.SeqOf(_775_v90, _dafny.SeqOf((_776_v107).F34(), (_776_v107).F34(), false)))
						}
					})(_755_v90, _772_v107)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw114 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw114).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw114).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_774_v109 = _nw114
				var _778_v110 _dafny.Sequence
				_ = _778_v110
				_778_v110 = _dafny.SeqOf(_755_v90)
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_774_v109), 0))
				_ = _index97
				(_774_v109).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_778_v110, _778_v110), (_index97).Int())
				var _779_v111 _dafny.Map
				_ = _779_v111
				_779_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F24()), _755_v90)
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_774_v109), 0))
				_ = _index98
				(_774_v109).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_772_v107).F34() {
						return _dafny.Companion_Sequence_.Concatenate(_778_v110, _778_v110)
					}
					return (func() _dafny.Sequence {
						if !((_772_v107).F34()) {
							return _778_v110
						}
						return _dafny.SeqOf(_dafny.SeqOf((_772_v107).F34()), (func() _dafny.Sequence {
							if (_779_v111).Contains(_dafny.IntOfInt64(370)) {
								return (_779_v111).Get(_dafny.IntOfInt64(370)).(_dafny.Sequence)
							}
							return _755_v90
						})(), _755_v90, _755_v90)
					})()
				})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality())), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_772_v107).F34() {
						return _dafny.Companion_Sequence_.Concatenate(_778_v110, _778_v110)
					}
					return (func() _dafny.Sequence {
						if !((_772_v107).F34()) {
							return _778_v110
						}
						return _dafny.SeqOf(_dafny.SeqOf((_772_v107).F34()), (func() _dafny.Sequence {
							if (_779_v111).Contains(_dafny.IntOfInt64(370)) {
								return (_779_v111).Get(_dafny.IntOfInt64(370)).(_dafny.Sequence)
							}
							return _755_v90
						})(), _755_v90, _755_v90)
					})()
				})()).Cardinality()))).Uint32(), _755_v90), (_index98).Int())
				_755_v90 = _755_v90
			} else {
				var _780_v112 _dafny.MultiSet
				_ = _780_v112
				_780_v112 = _dafny.MultiSetOf((_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()), _752_v88)
				var _781_v113 _dafny.Sequence
				_ = _781_v113
				_781_v113 = _dafny.SeqOf((_this).F24(), (_this).F24())
				(globalState).F14 = Companion_Default___.Fm0(Companion_Default___.Fm0(_dafny.IntOfInt64(233), Companion_Default___.Fm0((_this).F24(), (func() _dafny.Int {
					if (_780_v112).Contains(_752_v88) {
						return (_780_v112).Multiplicity(_752_v88)
					}
					return _dafny.IntOfUint32((_781_v113).Cardinality())
				})(), globalState), globalState), (_this).F24(), globalState)
				var _782_v114 _dafny.Map
				_ = _782_v114
				_782_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_753_v89, _dafny.SeqOf(_this.F31, _this.F31))
				var _783_v115 _dafny.Set
				_ = _783_v115
				_783_v115 = _dafny.SetOf((func() _dafny.Sequence {
					if (_782_v114).Contains(_753_v89) {
						return (_782_v114).Get(_753_v89).(_dafny.Sequence)
					}
					return _755_v90
				})(), _755_v90, Companion_Default___.Fm23(globalState))
				var _784_v116 _dafny.Sequence
				_ = _784_v116
				_784_v116 = _dafny.SeqOf(_755_v90, _755_v90)
				_783_v115 = func() _dafny.Set {
					var _coll22 = _dafny.NewBuilder()
					_ = _coll22
					for _iter23 := _dafny.Iterate((_784_v116).Elements()); ; {
						_compr_22, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _785_v117 _dafny.Sequence
						_785_v117 = interface{}(_compr_22).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_784_v116, _785_v117) {
							_coll22.Add(_785_v117)
						}
					}
					return _coll22.ToSet()
				}()
				var _786_v118 _dafny.Map
				_ = _786_v118
				_786_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_783_v115).Cardinality())
				var _787_v119 _dafny.Int
				_ = _787_v119
				var _out23 _dafny.Int
				_ = _out23
				_out23 = Companion_Default___.M0((func() _dafny.Int {
					if (_786_v118).Contains((_this).F24()) {
						return (_786_v118).Get((_this).F24()).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-202)
				})(), (_dafny.Zero).Minus((_756_v91).Cardinality()), globalState)
				_787_v119 = _out23
				(globalState).F1 = (func() _dafny.Int {
					if _this.F31 {
						return _754_i13
					}
					return Companion_Default___.SafeDivisionInt(_787_v119, _dafny.IntOfInt64(958))
				})()
				(globalState).F14 = _787_v119
			}
			var _788_v120 _dafny.MultiSet
			_ = _788_v120
			_788_v120 = _dafny.MultiSetOf(_this.F31)
			var _789_v121 D6
			_ = _789_v121
			_789_v121 = Companion_D6_.Create_DC18_(false)
			var _790_v122 _dafny.Map
			_ = _790_v122
			_790_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, _dafny.IntOfInt64(513))
			var _791_v123 _dafny.Array
			_ = _791_v123
			var _nwElement0_20 bool = (_789_v121).Dtor_cf29()
			_ = _nwElement0_20
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(11))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_20, 0)
			(_nw115).ArraySet1(_this.F31, 1)
			(_nw115).ArraySet1((_this).Fm9(_this.F31, (_this).F24(), (_this).F24(), globalState), 2)
			(_nw115).ArraySet1((_this).Fm9(_this.F31, (_this).F24(), (_790_v122).Cardinality(), globalState), 3)
			(_nw115).ArraySet1(_this.F31, 4)
			(_nw115).ArraySet1(_this.F31, 5)
			(_nw115).ArraySet1(_this.F31, 6)
			(_nw115).ArraySet1(_this.F31, 7)
			(_nw115).ArraySet1(_this.F31, 8)
			(_nw115).ArraySet1(_this.F31, 9)
			(_nw115).ArraySet1(_this.F31, 10)
			_791_v123 = _nw115
			var _792_v124 _dafny.Map
			_ = _792_v124
			_792_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, _791_v123)
			var _793_v125 _dafny.Map
			_ = _793_v125
			_793_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()), ((_788_v120).Update(!(_this.F31), Companion_Default___.Abs((_792_v124).Cardinality()))).Union(_788_v120))
			_793_v125 = (_793_v125).Update(_752_v88, _788_v120)
		}
		var _794_v126 D6
		_ = _794_v126
		_794_v126 = Companion_D6_.Create_DC16_(Companion_Default___.Fm18(globalState))
		var _source9 D6 = Companion_D6_.Create_DC19_(_794_v126)
		_ = _source9
		if _source9.Is_DC17() {
			var _795___mcc_h11 bool = _source9.Get_().(D6_DC17).Cf28
			_ = _795___mcc_h11
			var _796_cf28 bool = _795___mcc_h11
			_ = _796_cf28
			var _797_v127 _dafny.Array
			_ = _797_v127
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_19
			var _nw116 _dafny.Array
			_ = _nw116
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw116 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Int = func(_798_i15 _dafny.Int) _dafny.Int {
					return (_798_i15).Plus((_this).F24())
				}
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw116 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw116).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw116).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_797_v127 = _nw116
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))
			_ = _index99
			(_797_v127).ArraySet1((_this).F24(), (_index99).Int())
			var _799_v128 _dafny.Set
			_ = _799_v128
			_799_v128 = _dafny.SetOf((_this).F24())
			var _800_v129 _dafny.MultiSet
			_ = _800_v129
			_800_v129 = _dafny.MultiSetOf(_796_cf28, _this.F31)
			var _801_v131 _dafny.MultiSet
			_ = _801_v131
			_801_v131 = _dafny.MultiSetOf(_799_v128, _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(51))).Cardinality()), (_this).F24(), (_800_v129).Cardinality()), func() _dafny.Set {
				var _coll23 = _dafny.NewBuilder()
				_ = _coll23
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(103), _dafny.IntOfInt64(967))); ; {
					_compr_23, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _802_v130 _dafny.Int
					_802_v130 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(103)).Cmp(_802_v130) <= 0) && ((_802_v130).Cmp(_dafny.IntOfInt64(967)) < 0) {
						_coll23.Add((_802_v130).Plus(_dafny.IntOfUint32((_753_v89).Cardinality())))
					}
				}
				return _coll23.ToSet()
			}())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))
			_ = _index100
			(_797_v127).ArraySet1((func() _dafny.Int {
				if _this.F31 {
					return (func() _dafny.Int {
						if (_801_v131).Contains(_799_v128) {
							return (_801_v131).Multiplicity(_799_v128)
						}
						return (_this).F24()
					})()
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _this.F31)).Cardinality()
			})(), (_index100).Int())
			if (_this).Fm9(true, _dafny.IntOfInt64(726), (_this).F24(), globalState) {
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))
				_ = _index101
				(_750_v87).ArraySet1CodePoint((_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()), (_index101).Int())
				var _803_v132 _dafny.MultiSet
				_ = _803_v132
				_803_v132 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(866))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_804_v87 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_805_i16 _dafny.Int) _dafny.CodePoint {
						return (_804_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_804_v87), 0))).Int())
					}
				})(_750_v87))), (func() _dafny.Sequence {
					if _this.F31 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(955))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg57 _dafny.Int) interface{} {
								return coer57(arg57)
							}
						}((func(_806_v88 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_807_i17 _dafny.Int) _dafny.CodePoint {
								return _806_v88
							}
						})(_752_v88)))
					}
					return _dafny.Companion_Sequence_.Update(_753_v89, (Companion_Default___.SafeIndex((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_753_v89).Cardinality()))).Uint32(), _dafny.CodePoint('o'))
				})(), _dafny.UnicodeSeqOfUtf8Bytes("hqys"))
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))
				_ = _index102
				var _rhs121 _dafny.Sequence = _753_v89
				_ = _rhs121
				var _rhs122 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), (_this).F24()))
				_ = _rhs122
				var _rhs123 _dafny.MultiSet = _803_v132
				_ = _rhs123
				var _lhs95 _dafny.Array = _797_v127
				_ = _lhs95
				var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))
				_ = _lhs96
				_753_v89 = _rhs121
				(_lhs95).ArraySet1(_rhs122, (_lhs96).Int())
				_803_v132 = _rhs123
				var _808_v133 _dafny.Sequence
				_ = _808_v133
				_808_v133 = _dafny.SeqOf(_796_cf28)
				var _rhs124 _dafny.Int = (_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int)
				_ = _rhs124
				var _rhs125 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_808_v133, _808_v133)
				_ = _rhs125
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				var _lhs98 *GlobalState = globalState
				_ = _lhs98
				_lhs97.F1 = _rhs124
				_lhs98.F20 = _rhs125
				(_this).F31 = ((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int)).Cmp((_this).F24()) <= 0
				var _809_v134 _dafny.Sequence
				_ = _809_v134
				_809_v134 = _dafny.SeqOf(_797_v127, _797_v127)
				var _810_v135 D1
				_ = _810_v135
				_810_v135 = Companion_D1_.Create_DC4_(_796_cf28, _this.F31, _752_v88, (_dafny.Zero).Minus((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int)), (_this).F24())
				var _811_v136 _dafny.Array
				_ = _811_v136
				var _nwElement0_21 _dafny.Array = _797_v127
				_ = _nwElement0_21
				var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(17))
				_ = _nw117
				(_nw117).ArraySet1(_nwElement0_21, 0)
				(_nw117).ArraySet1(_797_v127, 1)
				(_nw117).ArraySet1(_797_v127, 2)
				(_nw117).ArraySet1(_797_v127, 3)
				(_nw117).ArraySet1(_797_v127, 4)
				(_nw117).ArraySet1(_797_v127, 5)
				(_nw117).ArraySet1(_797_v127, 6)
				(_nw117).ArraySet1(_797_v127, 7)
				(_nw117).ArraySet1(_797_v127, 8)
				(_nw117).ArraySet1(_797_v127, 9)
				(_nw117).ArraySet1(_797_v127, 10)
				(_nw117).ArraySet1(_797_v127, 11)
				(_nw117).ArraySet1(_797_v127, 12)
				(_nw117).ArraySet1(_797_v127, 13)
				(_nw117).ArraySet1((_809_v134).Select((Companion_Default___.SafeIndex((_810_v135).Dtor_cf12(), _dafny.IntOfUint32((_809_v134).Cardinality()))).Uint32()).(_dafny.Array), 14)
				(_nw117).ArraySet1((_809_v134).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_809_v134).Cardinality()))).Uint32()).(_dafny.Array), 15)
				(_nw117).ArraySet1(_797_v127, 16)
				_811_v136 = _nw117
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_811_v136), 0))
				_ = _index103
				(_811_v136).ArraySet1(_797_v127, (_index103).Int())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_811_v136), 0))
				_ = _index104
				(_811_v136).ArraySet1(_797_v127, (_index104).Int())
			} else {
				var _812_v137 _dafny.MultiSet
				_ = _812_v137
				_812_v137 = _dafny.MultiSetOf((_this).F24(), (_this).F24())
				var _813_v138 _dafny.Map
				_ = _813_v138
				_813_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if _796_cf28 {
						return _796_cf28
					}
					return _this.F31
				})(), (_812_v137).Difference(_812_v137))
				(globalState).F9 = (_813_v138).Cardinality()
				var _814_v139 D6
				_ = _814_v139
				_814_v139 = Companion_D6_.Create_DC16_(_dafny.SetOf(_dafny.IntOfUint32((_753_v89).Cardinality())))
				var _815_v140 _dafny.Map
				_ = _815_v140
				_815_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, _814_v139)
				var _816_v141 _dafny.Map
				_ = _816_v141
				_816_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_815_v140, _796_cf28)
				_816_v141 = ((_816_v141).Merge(_816_v141)).Merge(_816_v141)
				var _817_v142 _dafny.Map
				_ = _817_v142
				_817_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_812_v137).Update((_this).F24(), Companion_Default___.Abs((_this).F24())), _752_v88)
				var _818_v144 _dafny.MultiSet
				_ = _818_v144
				_818_v144 = _dafny.MultiSetOf((_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()))
				_817_v142 = (_817_v142).Update(_dafny.MultiSetOf((func() _dafny.Map {
					var _coll24 = _dafny.NewMapBuilder()
					_ = _coll24
					for _iter25 := _dafny.Iterate((_818_v144).Elements()); ; {
						_compr_24, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _819_v143 _dafny.CodePoint
						_819_v143 = interface{}(_compr_24).(_dafny.CodePoint)
						if (_818_v144).Contains(_819_v143) {
							_coll24.Add(_819_v143, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, (_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int)))
						}
					}
					return _coll24.ToMap()
				}()).Cardinality()), (_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()))
				var _820_v145 *C0
				_ = _820_v145
				var _nw118 *C0 = New_C0_()
				_ = _nw118
				_nw118.Ctor__(_this.F31)
				_820_v145 = _nw118
				var _821_v146 *C1
				_ = _821_v146
				var _nw119 *C1 = New_C1_()
				_ = _nw119
				_nw119.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(286))).Uint32(), func(coer58 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}((func(_822_v137 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_823_i18 _dafny.Int) _dafny.MultiSet {
						return _822_v137
					}
				})(_812_v137))), (_dafny.MultiSetOf((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int))).Cardinality(), (_this).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(367))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}((func(_824_v88 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
					return func(_825_i19 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_824_v88, _dafny.IntOfInt64(-786))
					}
				})(_752_v88))))
				_821_v146 = _nw119
			}
			(globalState).F14 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-45), ((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfInt64(519)))
			var _826_v147 _dafny.Sequence
			_ = _826_v147
			_826_v147 = _dafny.SeqOf(false, _796_cf28, _796_cf28, !(_796_cf28), _796_cf28)
			if _dafny.Companion_Sequence_.Contains(_826_v147, _dafny.Companion_Sequence_.Contains(_826_v147, _796_cf28)) {
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))
				_ = _index105
				(_750_v87).ArraySet1CodePoint(_752_v88, (_index105).Int())
				var _827_v148 _dafny.Map
				_ = _827_v148
				_827_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_752_v88, _this.F31)
				_827_v148 = (_827_v148).Update((_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()), !_dafny.Companion_Sequence_.Equal(_826_v147, _826_v147))
				r0 = !((_this.F31) == ((func() bool {
					if _this.F31 {
						return true
					}
					return _796_cf28
				})()))
				(globalState).F2 = !_dafny.Companion_Sequence_.Equal(_753_v89, _dafny.Companion_Sequence_.Concatenate(_753_v89, _dafny.UnicodeSeqOfUtf8Bytes("bxrlmgp")))
				(globalState).F1 = _dafny.IntOfUint32((_753_v89).Cardinality())
			} else {
				var _828_v149 _dafny.Map
				_ = _828_v149
				_828_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, (_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()))
				var _829_v150 _dafny.Set
				_ = _829_v150
				_829_v150 = _dafny.SetOf((func() _dafny.CodePoint {
					if (_828_v149).Contains(_796_cf28) {
						return (_828_v149).Get(_796_cf28).(_dafny.CodePoint)
					}
					return _752_v88
				})())
				var _830_v151 D3
				_ = _830_v151
				_830_v151 = Companion_D3_.Create_DC11_((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), _796_cf28, Companion_Default___.Fm2((_829_v150).Cardinality(), (_this).F24(), Companion_Default___.Fm3(globalState), globalState))
				(globalState).F2 = (_830_v151).Dtor_cf19()
				(globalState).F18 = ((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int)).Minus((_this).F24())
				_796_cf28 = !(((Companion_Default___.Fm0((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), (_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), globalState)).Minus((_this).F24())).Cmp(Companion_Default___.SafeModuloInt((_this).F24(), (_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int))) >= 0)
				var _831_v152 _dafny.MultiSet
				_ = _831_v152
				_831_v152 = _dafny.MultiSetOf((_this).F24(), (_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), (_this).F24())
				var _832_v153 T0
				_ = _832_v153
				var _nw120 *C1 = New_C1_()
				_ = _nw120
				_nw120.Ctor__(_dafny.SeqOf(_831_v152, (_831_v152).Update((_831_v152).Cardinality(), Companion_Default___.Abs(_dafny.IntOfUint32((_826_v147).Cardinality()))), (_831_v152).Update((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), Companion_Default___.Abs((_this).F24())), _831_v152, _831_v152), _dafny.IntOfUint32((_753_v89).Cardinality()), (_this).F24(), (_this).F25())
				_832_v153 = _nw120
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))
				_ = _index106
				var _rhs126 T0 = _832_v153
				_ = _rhs126
				var _rhs127 bool = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_797_v127, (_832_v153).F24())).Cardinality()).Cmp((_831_v152).Cardinality()) >= 0
				_ = _rhs127
				var _rhs128 bool = _this.F31
				_ = _rhs128
				var _rhs129 _dafny.Sequence = _753_v89
				_ = _rhs129
				var _rhs130 _dafny.Int = Companion_Default___.SafeModuloInt((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_832_v153).F24()))
				_ = _rhs130
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				var _lhs100 *GlobalState = globalState
				_ = _lhs100
				var _lhs101 _dafny.Array = _797_v127
				_ = _lhs101
				var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))
				_ = _lhs102
				_832_v153 = _rhs126
				_lhs99.F2 = _rhs127
				r0 = _rhs128
				_lhs100.F17 = _rhs129
				(_lhs101).ArraySet1(_rhs130, (_lhs102).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))
				_ = _index107
				(_797_v127).ArraySet1((_797_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_797_v127), 0))).Int()).(_dafny.Int), (_index107).Int())
			}
		} else if _source9.Is_DC18() {
			var _833___mcc_h12 bool = _source9.Get_().(D6_DC18).Cf29
			_ = _833___mcc_h12
			var _834_cf29 bool = _833___mcc_h12
			_ = _834_cf29
			(globalState).F1 = (_this).F24()
			var _835_v154 _dafny.Map
			_ = _835_v154
			_835_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F24())
			_835_v154 = (_835_v154).Update(_834_cf29, (_dafny.IntOfInt64(188)).Minus((_this).F24()))
			var _836_v155 _dafny.Array
			_ = _836_v155
			var _nw121 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw121
			_836_v155 = _nw121
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_836_v155), 0))
			_ = _index108
			(_836_v155).ArraySet1(_this.F31, (_index108).Int())
			var _837_v156 _dafny.Set
			_ = _837_v156
			_837_v156 = _dafny.SetOf(_836_v155)
			var _838_v157 _dafny.Sequence
			_ = _838_v157
			_838_v157 = _dafny.SeqOf((_this).F24(), (_this).F24())
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_836_v155), 0))
			_ = _index109
			var _rhs131 _dafny.Int = (_this).F24()
			_ = _rhs131
			var _rhs132 _dafny.Int = (_838_v157).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_838_v157).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs132
			var _rhs133 bool = _834_cf29
			_ = _rhs133
			var _rhs134 _dafny.Set = (_837_v156).Union(_837_v156)
			_ = _rhs134
			var _lhs103 *GlobalState = globalState
			_ = _lhs103
			var _lhs104 *GlobalState = globalState
			_ = _lhs104
			var _lhs105 _dafny.Array = _836_v155
			_ = _lhs105
			var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_836_v155), 0))
			_ = _lhs106
			_lhs103.F18 = _rhs131
			_lhs104.F14 = _rhs132
			(_lhs105).ArraySet1(_rhs133, (_lhs106).Int())
			_837_v156 = _rhs134
			(globalState).F14 = (_this).F24()
		} else if _source9.Is_DC16() {
			var _839___mcc_h13 _dafny.Set = _source9.Get_().(D6_DC16).Cf27
			_ = _839___mcc_h13
			var _840_cf27 _dafny.Set = _839___mcc_h13
			_ = _840_cf27
			var _841_v158 _dafny.Map
			_ = _841_v158
			_841_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, (_this).F24())
			var _842_v159 _dafny.Int
			_ = _842_v159
			var _out24 _dafny.Int
			_ = _out24
			_out24 = Companion_Default___.M0(Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState), (func() _dafny.Int {
				if (_841_v158).Contains(_this.F31) {
					return (_841_v158).Get(_this.F31).(_dafny.Int)
				}
				return _dafny.IntOfInt64(464)
			})(), globalState)
			_842_v159 = _out24
			var _843_v160 D1
			_ = _843_v160
			_843_v160 = Companion_D1_.Create_DC4_(_this.F31, _this.F31, _dafny.CodePoint('i'), _dafny.IntOfInt64(983), (_this).F24())
			var _844_v161 _dafny.Map
			_ = _844_v161
			_844_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v160, _this.F31)
			var _845_v163 _dafny.Array
			_ = _845_v163
			var _nwElement0_22 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v160, _this.F31)
			_ = _nwElement0_22
			var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(12))
			_ = _nw122
			(_nw122).ArraySet1(_nwElement0_22, 0)
			(_nw122).ArraySet1(_844_v161, 1)
			(_nw122).ArraySet1((_844_v161).Merge(_844_v161), 2)
			(_nw122).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v160, _this.F31)).Merge(_844_v161), 3)
			(_nw122).ArraySet1((_844_v161).Merge(_844_v161), 4)
			(_nw122).ArraySet1(_844_v161, 5)
			(_nw122).ArraySet1(_844_v161, 6)
			(_nw122).ArraySet1(_844_v161, 7)
			(_nw122).ArraySet1(func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter26 := _dafny.Iterate((_844_v161).Keys().Elements()); ; {
					_compr_25, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _846_v162 D1
					_846_v162 = interface{}(_compr_25).(D1)
					if (_844_v161).Contains(_846_v162) {
						_coll25.Add(_846_v162, _this.F31)
					}
				}
				return _coll25.ToMap()
			}(), 8)
			(_nw122).ArraySet1((_844_v161).Merge(_844_v161), 9)
			(_nw122).ArraySet1((_844_v161).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v160, _this.F31)), 10)
			(_nw122).ArraySet1((_844_v161).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v160, Companion_Default___.Fm2(_842_v159, _842_v159, (_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()), globalState))), 11)
			_845_v163 = _nw122
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_845_v163), 0))
			_ = _index110
			(_845_v163).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v160, _this.F31)).Merge(_844_v161), (_index110).Int())
			var _847_v164 _dafny.MultiSet
			_ = _847_v164
			_847_v164 = _dafny.MultiSetOf(_this.F31, _this.F31)
			var _848_v165 _dafny.Sequence
			_ = _848_v165
			_848_v165 = _dafny.SeqOf(_this.F31)
			var _849_v166 D4
			_ = _849_v166
			_849_v166 = Companion_D4_.Create_DC13_(_842_v159, _841_v158, _dafny.UnicodeSeqOfUtf8Bytes("libuqyvxh"))
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_845_v163), 0))
			_ = _index111
			(_845_v163).ArraySet1(Companion_Default___.Fm32((_this).F24(), _dafny.SeqOf(_847_v164, _847_v164, _847_v164, _847_v164), _848_v165, _dafny.Companion_Sequence_.Update(_753_v89, (Companion_Default___.SafeIndex((_849_v166).Dtor_cf22(), _dafny.IntOfUint32((_753_v89).Cardinality()))).Uint32(), _752_v88), globalState), (_index111).Int())
			(globalState).F18 = Companion_Default___.Fm0((_this).F24(), _842_v159, globalState)
			var _850_v167 _dafny.Array
			_ = _850_v167
			var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
			_ = _nw123
			_850_v167 = _nw123
			var _851_v168 _dafny.Array
			_ = _851_v168
			var _nwElement0_23 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_753_v89, _753_v89)
			_ = _nwElement0_23
			var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(5))
			_ = _nw124
			(_nw124).ArraySet1(_nwElement0_23, 0)
			(_nw124).ArraySet1((func() _dafny.Sequence {
				if _this.F31 {
					return _753_v89
				}
				return Companion_Default___.Fm4(_this.F31, _842_v159, (_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()), _this.F31, globalState)
			})(), 1)
			(_nw124).ArraySet1(_753_v89, 2)
			(_nw124).ArraySet1(_753_v89, 3)
			(_nw124).ArraySet1(_753_v89, 4)
			_851_v168 = _nw124
			var _852_v169 _dafny.Map
			_ = _852_v169
			_852_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _this.F31)
			var _853_v170 _dafny.Sequence
			_ = _853_v170
			_853_v170 = _dafny.SeqOf(_842_v159)
			var _rhs135 _dafny.Int = (_dafny.Zero).Minus((((_this).F24()).Times((func() _dafny.Int {
				if (_847_v164).Contains(Companion_Default___.Fm2(Companion_Default___.Fm0((_this).F24(), _842_v159, globalState), (_this).F24(), _752_v88, globalState)) {
					return (_847_v164).Multiplicity(Companion_Default___.Fm2(Companion_Default___.Fm0((_this).F24(), _842_v159, globalState), (_this).F24(), _752_v88, globalState))
				}
				return _842_v159
			})())).Plus((_842_v159).Times(_842_v159)))
			_ = _rhs135
			var _rhs136 _dafny.Array = (func() _dafny.Array {
				if false {
					return _850_v167
				}
				return _850_v167
			})()
			_ = _rhs136
			var _rhs137 bool = (Companion_Default___.Fm16(globalState)).Equals((_852_v169).Merge(_852_v169))
			_ = _rhs137
			var _rhs138 _dafny.Int = ((_853_v170).Select((Companion_Default___.SafeIndex(_842_v159, _dafny.IntOfUint32((_853_v170).Cardinality()))).Uint32()).(_dafny.Int)).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_v159, _this.F31)).Cardinality())
			_ = _rhs138
			var _rhs139 _dafny.Array = _851_v168
			_ = _rhs139
			var _lhs107 *GlobalState = globalState
			_ = _lhs107
			var _lhs108 *GlobalState = globalState
			_ = _lhs108
			_lhs107.F18 = _rhs135
			_850_v167 = _rhs136
			r0 = _rhs137
			_lhs108.F14 = _rhs138
			_851_v168 = _rhs139
		} else {
			var _854___mcc_h14 D6 = _source9.Get_().(D6_DC19).Cf30
			_ = _854___mcc_h14
			var _855_cf30 D6 = _854___mcc_h14
			_ = _855_cf30
			var _856_v171 D1
			_ = _856_v171
			_856_v171 = Companion_D1_.Create_DC6_()
			_856_v171 = _856_v171
			var _857_v172 D6
			_ = _857_v172
			_857_v172 = Companion_D6_.Create_DC17_(_this.F31)
			(globalState).F2 = ((_857_v172).Dtor_cf28()) && (_this.F31)
			(globalState).F2 = _this.F31
			var _858_v173 _dafny.Sequence
			_ = _858_v173
			_858_v173 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_753_v89, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(849))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}((func(_859_v88 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_860_i20 _dafny.Int) _dafny.CodePoint {
					return _859_v88
				}
			})(_752_v88))))).Cardinality()), ((_this).F24()).Minus((_this).F24()))
			_858_v173 = _dafny.SeqOf((_this).F24(), (_this).F24())
		}
		var _861_v174 _dafny.Array
		_ = _861_v174
		var _nw125 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
		_ = _nw125
		_861_v174 = _nw125
		for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_861_v174), 0))); ; {
			_guard_loop_1, _ok27 := _iter27()
			if !_ok27 {
				break
			}
			var _862_i21 _dafny.Int
			_862_i21 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_862_i21).Sign() != -1) && ((_862_i21).Cmp(_dafny.ArrayLen((_861_v174), 0)) < 0)) {
				(_861_v174).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F24())), (_862_i21).Int())
			}
		}
		var _863_i22 _dafny.Int
		_ = _863_i22
		_863_i22 = _dafny.Zero
		{
			for !(_this.F31) {
				{
					if (_863_i22).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_863_i22 = (_863_i22).Plus(_dafny.One)
					if _this.F31 {
						var _864_v175 _dafny.Array
						_ = _864_v175
						var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
						_ = _nw126
						_864_v175 = _nw126
						var _865_v176 _dafny.Map
						_ = _865_v176
						_865_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_864_v175, _this.F31)
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_864_v175), 0))
						_ = _index112
						(_864_v175).ArraySet1(((_865_v176).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_864_v175, _this.F31)).Update(_864_v175, _this.F31))).Cardinality(), (_index112).Int())
						var _866_v177 _dafny.MultiSet
						_ = _866_v177
						_866_v177 = _dafny.MultiSetOf((_this).F24())
						var _867_v178 _dafny.Map
						_ = _867_v178
						_867_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rfeykmiww")).Cardinality()), (_this).F24(), globalState), _this.F31)
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_864_v175), 0))
						_ = _index113
						(_864_v175).ArraySet1((func() _dafny.Int {
							if _this.F31 {
								return (func() _dafny.Int {
									if (_866_v177).Contains((_this).F24()) {
										return (_866_v177).Multiplicity((_this).F24())
									}
									return (_this).F24()
								})()
							}
							return (_867_v178).Cardinality()
						})(), (_index113).Int())
						(globalState).F9 = (_864_v175).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_864_v175), 0))).Int()).(_dafny.Int)
						(globalState).F2 = ((_864_v175).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_864_v175), 0))).Int()).(_dafny.Int)).Cmp((_dafny.IntOfInt64(519)).Plus((_dafny.Zero).Minus((_864_v175).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_864_v175), 0))).Int()).(_dafny.Int)))) > 0
						var _868_v179 _dafny.Map
						_ = _868_v179
						_868_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_864_v175, (_this).F24())
						var _869_v180 _dafny.Sequence
						_ = _869_v180
						_869_v180 = _dafny.SeqOf(_864_v175)
						var _870_v181 _dafny.Set
						_ = _870_v181
						_870_v181 = _dafny.SetOf(_this.F31)
						_868_v179 = (_868_v179).Update((_869_v180).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(Companion_Default___.Fm0((_870_v181).Cardinality(), (_864_v175).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_864_v175), 0))).Int()).(_dafny.Int), globalState), (_this).F24(), globalState), _dafny.IntOfUint32((_869_v180).Cardinality()))).Uint32()).(_dafny.Array), (_this).F24())
						var _871_v182 _dafny.Set
						_ = _871_v182
						_871_v182 = _dafny.SetOf((_this).F24())
						var _872_v183 _dafny.Map
						_ = _872_v183
						_872_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, (_871_v182).Cardinality())
						(globalState).F2 = (_this).Fm9(_this.F31, Companion_Default___.Fm0((_872_v183).Cardinality(), (_864_v175).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_864_v175), 0))).Int()).(_dafny.Int), globalState), (_this).F24(), globalState)
					} else {
						var _873_v184 _dafny.Map
						_ = _873_v184
						_873_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm9(_this.F31, (_this).F24(), _dafny.IntOfInt64(593), globalState), _this.F31)
						_873_v184 = (_873_v184).Update(false, false)
						(globalState).F2 = _this.F31
						var _874_v185 _dafny.Map
						_ = _874_v185
						_874_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
						var _875_v186 _dafny.Sequence
						_ = _875_v186
						_875_v186 = _dafny.SeqOf(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_873_v184).Cardinality())).Update((_this).F24(), _dafny.IntOfInt64(-306))).Merge(_874_v185), _874_v185)
						_875_v186 = _dafny.Companion_Sequence_.Concatenate(_875_v186, _875_v186)
						(globalState).F2 = !(_this.F31) || ((_this.F31) == (_this.F31))
						var _876_v187 _dafny.Int
						_ = _876_v187
						var _out25 _dafny.Int
						_ = _out25
						_out25 = Companion_Default___.M0((_this).F24(), (_dafny.Zero).Minus((_this).F24()), globalState)
						_876_v187 = _out25
					}
					var _877_v188 _dafny.MultiSet
					_ = _877_v188
					_877_v188 = _dafny.MultiSetOf((_this).F24())
					var _878_v189 *C0
					_ = _878_v189
					var _nw127 *C0 = New_C0_()
					_ = _nw127
					_nw127.Ctor__((_dafny.MultiSetOf((_this).F24(), _dafny.IntOfUint32((_753_v89).Cardinality()))).IsSubsetOf(_877_v188))
					_878_v189 = _nw127
					if ((_this).F24()).Cmp((_this).F24()) > 0 {
						var _879_v190 _dafny.Set
						_ = _879_v190
						_879_v190 = _dafny.SetOf(_dafny.IntOfUint32((_753_v89).Cardinality()), (_dafny.Zero).Minus((_this).F24()))
						var _880_v191 D6
						_ = _880_v191
						_880_v191 = Companion_D6_.Create_DC16_(_879_v190)
						_880_v191 = _880_v191
						var _881_v192 _dafny.Map
						_ = _881_v192
						_881_v192 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), Companion_Default___.Fm18(globalState))
						var _882_v194 _dafny.Sequence
						_ = _882_v194
						_882_v194 = _dafny.SeqOf(_879_v190, func() _dafny.Set {
							var _coll26 = _dafny.NewBuilder()
							_ = _coll26
							for _iter28 := _dafny.Iterate((_877_v188).Elements()); ; {
								_compr_26, _ok28 := _iter28()
								if !_ok28 {
									break
								}
								var _883_v193 _dafny.Int
								_883_v193 = interface{}(_compr_26).(_dafny.Int)
								if (_877_v188).Contains(_883_v193) {
									_coll26.Add((_883_v193).Plus(_dafny.IntOfInt64(918)))
								}
							}
							return _coll26.ToSet()
						}(), _879_v190)
						var _884_v195 _dafny.Sequence
						_ = _884_v195
						_884_v195 = _dafny.SeqOf((_this).F24())
						_879_v190 = (func() _dafny.Set {
							if (_881_v192).Contains((_this).F24()) {
								return (_881_v192).Get((_this).F24()).(_dafny.Set)
							}
							return (func() _dafny.Set {
								if Companion_Default___.Fm2((_this).F24(), (_this).F24(), (_750_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_750_v87), 0))).Int()), globalState) {
									return (_882_v194).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_884_v195).Cardinality()), _dafny.IntOfUint32((_882_v194).Cardinality()))).Uint32()).(_dafny.Set)
								}
								return _879_v190
							})()
						})()
						var _885_v196 _dafny.Map
						_ = _885_v196
						_885_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm23(globalState), Companion_Default___.Fm2((_this).F24(), (_dafny.Zero).Minus((_this).F24()), _dafny.CodePoint('n'), globalState))
						var _886_v197 _dafny.Map
						_ = _886_v197
						_886_v197 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31, (_878_v189).F34())
						var _887_v198 _dafny.Sequence
						_ = _887_v198
						_887_v198 = _dafny.SeqOf((_878_v189).F34())
						_885_v196 = (_885_v196).Update((func() _dafny.Sequence {
							if (func() bool {
								if (_886_v197).Contains((_this).Fm9((_878_v189).F34(), (_this).F24(), (_this).F24(), globalState)) {
									return (_886_v197).Get((_this).Fm9((_878_v189).F34(), (_this).F24(), (_this).F24(), globalState)).(bool)
								}
								return (_878_v189).F34()
							})() {
								return _887_v198
							}
							return _887_v198
						})(), (_878_v189).F34())
						var _888_v199 D1
						_ = _888_v199
						_888_v199 = Companion_D1_.Create_DC3_(_753_v89)
						_888_v199 = Companion_D1_.Create_DC3_(_753_v89)
						(_this).F31 = _this.F31
					} else {
						var _889_v200 _dafny.Set
						_ = _889_v200
						_889_v200 = _dafny.SetOf(_dafny.IntOfInt64(104))
						var _890_v201 _dafny.Map
						_ = _890_v201
						_890_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_878_v189).F34(), _889_v200)
						var _891_v202 _dafny.Sequence
						_ = _891_v202
						_891_v202 = _dafny.SeqOf(!((_878_v189).F34()))
						var _892_v203 _dafny.Map
						_ = _892_v203
						_892_v203 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_890_v201, (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _891_v202))).Cardinality())
						_892_v203 = (_892_v203).Update(_890_v201, (_this).F24())
						var _893_v204 _dafny.Int
						_ = _893_v204
						var _out26 _dafny.Int
						_ = _out26
						_out26 = Companion_Default___.M0((_dafny.Zero).Minus((_this).F24()), (_this).F24(), globalState)
						_893_v204 = _out26
						var _894_v205 _dafny.Array
						_ = _894_v205
						var _len0_20 _dafny.Int = _dafny.IntOfInt64(26)
						_ = _len0_20
						var _nw128 _dafny.Array
						_ = _nw128
						if _len0_20.Cmp(_dafny.Zero) == 0 {
							_nw128 = _dafny.NewArray(_len0_20)
						} else {
							var _init20 func(_dafny.Int) bool = func(_895_i23 _dafny.Int) bool {
								return !(_this.F31)
							}
							_ = _init20
							var _element0_20 = _init20(_dafny.Zero)
							_ = _element0_20
							_nw128 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
							(_nw128).ArraySet1(_element0_20, 0)
							var _nativeLen0_20 = (_len0_20).Int()
							_ = _nativeLen0_20
							for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
								(_nw128).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
							}
						}
						_894_v205 = _nw128
						var _rhs140 _dafny.Int = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tlpy")).Cardinality())).Times((_dafny.Zero).Minus((_this).F24()))
						_ = _rhs140
						var _rhs141 _dafny.Array = _894_v205
						_ = _rhs141
						_893_v204 = _rhs140
						_894_v205 = _rhs141
						var _896_v206 _dafny.Sequence
						_ = _896_v206
						_896_v206 = _dafny.SeqOf(_893_v204)
						_896_v206 = _dafny.Companion_Sequence_.Concatenate(_896_v206, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(134))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg61 _dafny.Int) interface{} {
								return coer61(arg61)
							}
						}((func(_897_v204 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_898_i24 _dafny.Int) _dafny.Int {
								return _897_v204
							}
						})(_893_v204))))
						(globalState).F2 = (_878_v189).F34()
					}
					var _899_v207 _dafny.Array
					_ = _899_v207
					var _len0_21 _dafny.Int = _dafny.IntOfInt64(23)
					_ = _len0_21
					var _nw129 _dafny.Array
					_ = _nw129
					if _len0_21.Cmp(_dafny.Zero) == 0 {
						_nw129 = _dafny.NewArray(_len0_21)
					} else {
						var _init21 func(_dafny.Int) _dafny.MultiSet = (func(_900_v189 *C0) func(_dafny.Int) _dafny.MultiSet {
							return func(_901_i25 _dafny.Int) _dafny.MultiSet {
								return (_dafny.MultiSetOf((_900_v189).F34())).Union(_dafny.MultiSetOf(_this.F31))
							}
						})(_878_v189)
						_ = _init21
						var _element0_21 = _init21(_dafny.Zero)
						_ = _element0_21
						_nw129 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
						(_nw129).ArraySet1(_element0_21, 0)
						var _nativeLen0_21 = (_len0_21).Int()
						_ = _nativeLen0_21
						for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
							(_nw129).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
						}
					}
					_899_v207 = _nw129
					_899_v207 = _899_v207
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		r0 = _this.F31
		return r0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f24 _dafny.Int
	_f25 _dafny.Sequence
	F30  _dafny.Array
	_f29 _dafny.Array
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f24 = _dafny.Zero
	_this._f25 = _dafny.EmptySeq
	_this.F30 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f29 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C3) F24() _dafny.Int {
	return _this._f24
}
func (_this *C3) F25() _dafny.Sequence {
	return _this._f25
}
func (_this *C3) Ctor__(f29 _dafny.Array, f30 _dafny.Array, f24 _dafny.Int, f25 _dafny.Sequence) {
	{
		(_this)._f29 = f29
		(_this).F30 = f30
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C3) Fm8(globalState *GlobalState) _dafny.Int {
	{
		return (Companion_D0_.Create_DC2_(false, (_this).F24(), _dafny.CodePoint('h'))).Dtor_cf5()
	}
}
func (_this *C3) M1(globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _902_v0 _dafny.Sequence
		_ = _902_v0
		_902_v0 = _dafny.UnicodeSeqOfUtf8Bytes("olddh")
		var _903_v1 _dafny.CodePoint
		_ = _903_v1
		_903_v1 = _dafny.CodePoint('w')
		var _904_v2 _dafny.Sequence
		_ = _904_v2
		_904_v2 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_902_v0, (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_902_v0).Cardinality()))).Uint32(), _903_v1))
		var _hi5 _dafny.Int = _dafny.IntOfUint32((_904_v2).Cardinality())
		_ = _hi5
		for _905_i0 := (_this).F24(); _905_i0.Cmp(_hi5) < 0; _905_i0 = _905_i0.Plus(_dafny.One) {
			var _906_v3 bool
			_ = _906_v3
			_906_v3 = true
			(globalState).F17 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm4(false, (_this).F24(), _903_v1, _906_v3, globalState), _dafny.UnicodeSeqOfUtf8Bytes("pglhejcp"))
			var _907_v4 _dafny.Array
			_ = _907_v4
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_22
			var _nw130 _dafny.Array
			_ = _nw130
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw130 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) bool = (func(_908_v3 bool) func(_dafny.Int) bool {
					return func(_909_i1 _dafny.Int) bool {
						return (func() bool {
							if _908_v3 {
								return _908_v3
							}
							return _908_v3
						})()
					}
				})(_906_v3)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw130 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw130).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw130).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_907_v4 = _nw130
			var _nw131 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw131
			_907_v4 = _nw131
			var _910_v5 _dafny.Sequence
			_ = _910_v5
			_910_v5 = _dafny.SeqOf(_906_v3)
			var _911_v6 _dafny.MultiSet
			_ = _911_v6
			_911_v6 = _dafny.MultiSetOf(_dafny.SeqOf(_906_v3))
			var _912_v7 _dafny.Sequence
			_ = _912_v7
			_912_v7 = _dafny.SeqOf(_911_v6, _dafny.MultiSetOf(_910_v5, _910_v5), _911_v6)
			if ((_912_v7).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_912_v7).Cardinality()))).Uint32()).(_dafny.MultiSet)).Contains(_dafny.Companion_Sequence_.Concatenate(_910_v5, _910_v5)) {
				var _913_v8 _dafny.Array
				_ = _913_v8
				var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
				_ = _nw132
				_913_v8 = _nw132
				var _914_v9 D2
				_ = _914_v9
				_914_v9 = Companion_D2_.Create_DC7_(_913_v8)
				var _915_v10 _dafny.Array
				_ = _915_v10
				var _nwElement0_24 _dafny.Array = _913_v8
				_ = _nwElement0_24
				var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(6))
				_ = _nw133
				(_nw133).ArraySet1(_nwElement0_24, 0)
				(_nw133).ArraySet1(_913_v8, 1)
				(_nw133).ArraySet1(_913_v8, 2)
				(_nw133).ArraySet1(_913_v8, 3)
				(_nw133).ArraySet1(_913_v8, 4)
				(_nw133).ArraySet1((_914_v9).Dtor_cf16(), 5)
				_915_v10 = _nw133
				_915_v10 = _915_v10
				var _916_v11 *C0
				_ = _916_v11
				var _nw134 *C0 = New_C0_()
				_ = _nw134
				_nw134.Ctor__(!_dafny.Companion_Sequence_.Equal(_910_v5, _dafny.SeqOf(true, false, _906_v3, _906_v3, _906_v3)))
				_916_v11 = _nw134
				(globalState).F18 = _905_i0
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_907_v4), 0))
				_ = _index114
				(_907_v4).ArraySet1(true, (_index114).Int())
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_907_v4), 0))
				_ = _index115
				(_907_v4).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_902_v0, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_902_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("apcb")).Cardinality()), _dafny.IntOfUint32((_902_v0).Cardinality()))).Uint32(), _903_v1), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(426))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}(func(_917_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				})))), (_index115).Int())
				var _918_v12 _dafny.MultiSet
				_ = _918_v12
				_918_v12 = _dafny.MultiSetOf(!((_916_v11).F34()), (_907_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_907_v4), 0))).Int()).(bool), (_916_v11).F34())
				var _919_v13 _dafny.Set
				_ = _919_v13
				_919_v13 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(711))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_920_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_921_i3 _dafny.Int) _dafny.CodePoint {
						return _920_v1
					}
				})(_903_v1))))
				var _rhs142 _dafny.Int = (func() _dafny.Int {
					if (_918_v12).IsDisjointFrom(_dafny.MultiSetOf((_916_v11).F34())) {
						return (_this).F24()
					}
					return _dafny.IntOfUint32((_902_v0).Cardinality())
				})()
				_ = _rhs142
				var _rhs143 bool = ((_919_v13).IsSubsetOf(Companion_Default___.Fm22(_905_i0, (_916_v11).F34(), (_916_v11).F34(), _dafny.IntOfUint32((_902_v0).Cardinality()), globalState))) && ((_916_v11).F34())
				_ = _rhs143
				var _rhs144 _dafny.Int = (_this).F24()
				_ = _rhs144
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				_lhs109.F14 = _rhs142
				_lhs110.F2 = _rhs143
				_lhs111.F9 = _rhs144
			} else {
				var _922_v14 _dafny.Map
				_ = _922_v14
				_922_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_902_v0, _906_v3)
				_922_v14 = (_922_v14).Update(_902_v0, _906_v3)
				var _923_v15 _dafny.Set
				_ = _923_v15
				_923_v15 = _dafny.SetOf(_906_v3)
				var _924_v16 D0
				_ = _924_v16
				_924_v16 = Companion_D0_.Create_DC2_((_905_i0).Cmp((_923_v15).Cardinality()) > 0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-589))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_925_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_926_i4 _dafny.Int) _dafny.Int {
						return _925_i0
					}
				})(_905_i0)))).Cardinality()), _dafny.CodePoint('g'))
				_924_v16 = _924_v16
				var _927_v18 _dafny.Array
				_ = _927_v18
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_23
				var _nw135 _dafny.Array
				_ = _nw135
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw135 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Set = func(_928_i5 _dafny.Int) _dafny.Set {
						return func() _dafny.Set {
							var _coll27 = _dafny.NewBuilder()
							_ = _coll27
							for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(291), _dafny.IntOfInt64(-29))); ; {
								_compr_27, _ok29 := _iter29()
								if !_ok29 {
									break
								}
								var _929_v17 _dafny.Int
								_929_v17 = interface{}(_compr_27).(_dafny.Int)
								if ((_dafny.IntOfInt64(291)).Cmp(_929_v17) <= 0) && ((_929_v17).Cmp(_dafny.IntOfInt64(-29)) < 0) {
									_coll27.Add(Companion_Default___.SafeModuloInt(_929_v17, (_this).F24()))
								}
							}
							return _coll27.ToSet()
						}()
					}
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw135 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw135).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw135).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_927_v18 = _nw135
				var _rhs145 _dafny.Array = (_this).F29()
				_ = _rhs145
				var _rhs146 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_902_v0, _902_v0), _902_v0)).Cardinality())
				_ = _rhs146
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				_927_v18 = _rhs145
				_lhs112.F1 = _rhs146
				var _930_v19 D8
				_ = _930_v19
				_930_v19 = Companion_D8_.Create_DC23_(_923_v15)
				_930_v19 = Companion_D8_.Create_DC23_(_923_v15)
				var _931_v20 _dafny.Map
				_ = _931_v20
				_931_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.IntOfUint32((_910_v5).Cardinality())).Plus((_this).F24()))))
				_931_v20 = (_931_v20).Update(_905_i0, Companion_Default___.SafeModuloInt(_905_i0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_902_v0).Cardinality()))))
			}
			(globalState).F18 = Companion_Default___.Fm0(_905_i0, ((_this).F24()).Plus(_905_i0), globalState)
		}
		var _hi6 _dafny.Int = (_this).F24()
		_ = _hi6
		for _932_i6 := _dafny.IntOfUint32((_902_v0).Cardinality()); _932_i6.Cmp(_hi6) < 0; _932_i6 = _932_i6.Plus(_dafny.One) {
			(globalState).F2 = (_dafny.IntOfUint32((_902_v0).Cardinality())).Cmp((_932_i6).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nhg")).Cardinality()))) == 0
			var _933_v21 bool
			_ = _933_v21
			_933_v21 = true
			var _934_v22 _dafny.Map
			_ = _934_v22
			_934_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_933_v21, _933_v21)
			_934_v22 = (_934_v22).Update(((_this).F24()).Cmp(_932_i6) != 0, _933_v21)
			_902_v0 = _dafny.Companion_Sequence_.Concatenate(_902_v0, (func() _dafny.Sequence {
				if _933_v21 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(481))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg65 _dafny.Int) interface{} {
							return coer65(arg65)
						}
					}((func(_935_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_936_i7 _dafny.Int) _dafny.CodePoint {
							return _935_v1
						}
					})(_903_v1)))
				}
				return _902_v0
			})())
			(globalState).F2 = _933_v21
		}
		var _937_v23 _dafny.Int
		_ = _937_v23
		var _out27 _dafny.Int
		_ = _out27
		_out27 = Companion_Default___.M0((_this).F24(), (_this).F24(), globalState)
		_937_v23 = _out27
		var _938_v24 bool
		_ = _938_v24
		_938_v24 = true
		var _939_v25 _dafny.MultiSet
		_ = _939_v25
		_939_v25 = _dafny.MultiSetOf(_938_v24)
		var _940_v26 _dafny.Sequence
		_ = _940_v26
		_940_v26 = _dafny.SeqOf((_939_v25).Cardinality(), (_this).F24(), _937_v23)
		var _941_v27 _dafny.Map
		_ = _941_v27
		_941_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v24, _940_v26)
		var _942_v28 _dafny.Map
		_ = _942_v28
		_942_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_941_v27).Contains(_938_v24) {
				return (_941_v27).Get(_938_v24).(_dafny.Sequence)
			}
			return _940_v26
		})()).Cardinality()), _938_v24)
		var _943_v29 _dafny.MultiSet
		_ = _943_v29
		_943_v29 = _dafny.MultiSetOf(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_942_v28, _937_v23)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_942_v28, _937_v23))).Cardinality())
		_943_v29 = (_dafny.MultiSetOf(_937_v23)).Intersection(_943_v29)
		var _source10 D3 = Companion_Default___.Fm33(_938_v24, (_dafny.Zero).Minus((_this).F24()), globalState)
		_ = _source10
		if _source10.Is_DC10() {
			_938_v24 = _938_v24
			var _944_v30 _dafny.Array
			_ = _944_v30
			var _nw136 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw136
			_944_v30 = _nw136
			var _945_v31 _dafny.Set
			_ = _945_v31
			_945_v31 = _dafny.SetOf(_944_v30, _944_v30)
			var _946_v32 _dafny.Map
			_ = _946_v32
			_946_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_938_v24), _938_v24)
			var _947_v33 _dafny.Set
			_ = _947_v33
			_947_v33 = _dafny.SetOf(!(Companion_Default___.Fm2(_dafny.IntOfInt64(872), (_this).F24(), _dafny.CodePoint('x'), globalState)))
			var _948_v34 _dafny.Sequence
			_ = _948_v34
			_948_v34 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_940_v26))
			var _949_v35 _dafny.Map
			_ = _949_v35
			_949_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_903_v1, (_dafny.Zero).Minus((_this).F24()))
			var _950_v36 *C1
			_ = _950_v36
			var _nw137 *C1 = New_C1_()
			_ = _nw137
			_nw137.Ctor__(_948_v34, (_this).F24(), (_this).F24(), _dafny.SeqOf(_949_v35))
			_950_v36 = _nw137
			var _951_v37 _dafny.Map
			_ = _951_v37
			_951_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_946_v32).Cardinality(), _937_v23, _903_v1, globalState), (_dafny.Zero).Minus(_937_v23))
			var _952_v38 _dafny.Map
			_ = _952_v38
			_952_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_950_v36, _951_v37)
			var _953_v39 _dafny.Array
			_ = _953_v39
			var _nwElement0_25 _dafny.Int = (_945_v31).Cardinality()
			_ = _nwElement0_25
			var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(20))
			_ = _nw138
			(_nw138).ArraySet1(_nwElement0_25, 0)
			(_nw138).ArraySet1(((_this).F24()).Times(_937_v23), 1)
			(_nw138).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F24()), (_this).F24()), 2)
			(_nw138).ArraySet1((_dafny.Zero).Minus(((_this).F24()).Times(_937_v23)), 3)
			(_nw138).ArraySet1((_this).F24(), 4)
			(_nw138).ArraySet1(_937_v23, 5)
			(_nw138).ArraySet1(((_946_v32).Merge(_946_v32)).Cardinality(), 6)
			(_nw138).ArraySet1((_947_v33).Cardinality(), 7)
			(_nw138).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hwuajdbrj")).Cardinality()), (_this).F24()), 8)
			(_nw138).ArraySet1((_952_v38).Cardinality(), 9)
			(_nw138).ArraySet1((Companion_D1_.Create_DC4_(_938_v24, _938_v24, _903_v1, (_this).F24(), (_this).Fm8(globalState))).Dtor_cf12(), 10)
			(_nw138).ArraySet1(_950_v36.F33, 11)
			(_nw138).ArraySet1((_this).F24(), 12)
			(_nw138).ArraySet1((_this).F24(), 13)
			(_nw138).ArraySet1(((_this).F24()).Minus(_950_v36.F33), 14)
			(_nw138).ArraySet1((_947_v33).Cardinality(), 15)
			(_nw138).ArraySet1((_this).F24(), 16)
			(_nw138).ArraySet1((_940_v26).Select((Companion_Default___.SafeIndex(_937_v23, _dafny.IntOfUint32((_940_v26).Cardinality()))).Uint32()).(_dafny.Int), 17)
			(_nw138).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hhs")).Cardinality())).Plus((_dafny.Zero).Minus(_937_v23)), 18)
			(_nw138).ArraySet1((_this).F24(), 19)
			_953_v39 = _nw138
			var _954_v40 _dafny.Map
			_ = _954_v40
			_954_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_951_v37, (_dafny.Zero).Minus(_950_v36.F33))
			var _955_v41 _dafny.Map
			_ = _955_v41
			_955_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v24, (_951_v37).Update(_938_v24, _950_v36.F33))
			var _nwElement0_26 _dafny.Int = (_this).F24()
			_ = _nwElement0_26
			var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(5))
			_ = _nw139
			(_nw139).ArraySet1(_nwElement0_26, 0)
			(_nw139).ArraySet1((func() _dafny.Int {
				if (_954_v40).Contains((func() _dafny.Map {
					if (_955_v41).Contains(_938_v24) {
						return (_955_v41).Get(_938_v24).(_dafny.Map)
					}
					return _951_v37
				})()) {
					return (_954_v40).Get((func() _dafny.Map {
						if (_955_v41).Contains(_938_v24) {
							return (_955_v41).Get(_938_v24).(_dafny.Map)
						}
						return _951_v37
					})()).(_dafny.Int)
				}
				return (_this).F24()
			})(), 1)
			(_nw139).ArraySet1(_950_v36.F33, 2)
			(_nw139).ArraySet1((_this).F24(), 3)
			(_nw139).ArraySet1(_937_v23, 4)
			_953_v39 = _nw139
			var _956_v42 _dafny.Array
			_ = _956_v42
			var _nw140 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw140
			_956_v42 = _nw140
			var _957_v43 _dafny.Map
			_ = _957_v43
			_957_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_902_v0, _956_v42)
			_956_v42 = (func() _dafny.Array {
				if (_957_v43).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-401))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}((func(_958_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_959_i8 _dafny.Int) _dafny.CodePoint {
						return _958_v1
					}
				})(_903_v1)))) {
					return (_957_v43).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-401))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg67 _dafny.Int) interface{} {
							return coer67(arg67)
						}
					}((func(_960_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_961_i8 _dafny.Int) _dafny.CodePoint {
							return _960_v1
						}
					})(_903_v1)))).(_dafny.Array)
				}
				return _956_v42
			})()
			(globalState).F17 = _dafny.Companion_Sequence_.Concatenate(_902_v0, _902_v0)
		} else if _source10.Is_DC11() {
			var _962___mcc_h0 _dafny.Int = _source10.Get_().(D3_DC11).Cf18
			_ = _962___mcc_h0
			var _963___mcc_h1 bool = _source10.Get_().(D3_DC11).Cf19
			_ = _963___mcc_h1
			var _964___mcc_h2 bool = _source10.Get_().(D3_DC11).Cf20
			_ = _964___mcc_h2
			var _965_cf20 bool = _964___mcc_h2
			_ = _965_cf20
			var _966_cf19 bool = _963___mcc_h1
			_ = _966_cf19
			var _967_cf18 _dafny.Int = _962___mcc_h0
			_ = _967_cf18
			var _968_v44 _dafny.Map
			_ = _968_v44
			_968_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v24, !(!(true)) || (true))
			(globalState).F2 = (func() bool {
				if (_968_v44).Contains(Companion_Default___.Fm2(_937_v23, _937_v23, _903_v1, globalState)) {
					return (_968_v44).Get(Companion_Default___.Fm2(_937_v23, _937_v23, _903_v1, globalState)).(bool)
				}
				return _966_cf19
			})()
			var _969_v45 _dafny.Array
			_ = _969_v45
			var _nwElement0_27 bool = _966_cf19
			_ = _nwElement0_27
			var _nw141 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(20))
			_ = _nw141
			(_nw141).ArraySet1(_nwElement0_27, 0)
			(_nw141).ArraySet1(_938_v24, 1)
			(_nw141).ArraySet1(!(_965_cf20), 2)
			(_nw141).ArraySet1(_966_cf19, 3)
			(_nw141).ArraySet1(_966_cf19, 4)
			(_nw141).ArraySet1(_938_v24, 5)
			(_nw141).ArraySet1(_938_v24, 6)
			(_nw141).ArraySet1(_966_cf19, 7)
			(_nw141).ArraySet1(false, 8)
			(_nw141).ArraySet1(true, 9)
			(_nw141).ArraySet1(_938_v24, 10)
			(_nw141).ArraySet1(_965_cf20, 11)
			(_nw141).ArraySet1(true, 12)
			(_nw141).ArraySet1(_938_v24, 13)
			(_nw141).ArraySet1(_938_v24, 14)
			(_nw141).ArraySet1(!(false), 15)
			(_nw141).ArraySet1(_938_v24, 16)
			(_nw141).ArraySet1(_965_cf20, 17)
			(_nw141).ArraySet1(_965_cf20, 18)
			(_nw141).ArraySet1(_966_cf19, 19)
			_969_v45 = _nw141
			var _970_v46 _dafny.Sequence
			_ = _970_v46
			_970_v46 = _dafny.SeqOf(_969_v45)
			_970_v46 = _970_v46
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_969_v45), 0))
			_ = _index116
			(_969_v45).ArraySet1(_966_cf19, (_index116).Int())
			var _971_v47 _dafny.Sequence
			_ = _971_v47
			_971_v47 = _dafny.SeqOf(_966_cf19, _938_v24)
			var _972_v48 _dafny.MultiSet
			_ = _972_v48
			_972_v48 = _dafny.MultiSetOf(_971_v47)
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_969_v45), 0))
			_ = _index117
			(_969_v45).ArraySet1(((func() _dafny.MultiSet {
				if _966_cf19 {
					return _972_v48
				}
				return _972_v48
			})()).IsSubsetOf(_972_v48), (_index117).Int())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_969_v45), 0))
			_ = _index118
			(_969_v45).ArraySet1((_969_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_969_v45), 0))).Int()).(bool), (_index118).Int())
		} else {
			var _973___mcc_h3 _dafny.Sequence = _source10.Get_().(D3_DC9).Cf17
			_ = _973___mcc_h3
			var _974_cf17 _dafny.Sequence = _973___mcc_h3
			_ = _974_cf17
			var _975_v49 _dafny.Sequence
			_ = _975_v49
			_975_v49 = _dafny.SeqOf(_dafny.MultiSetOf(_937_v23), _dafny.MultiSetOf(_dafny.IntOfUint32((_902_v0).Cardinality())), _943_v29, _943_v29, _943_v29)
			var _976_v51 *C1
			_ = _976_v51
			var _nw142 *C1 = New_C1_()
			_ = _nw142
			_nw142.Ctor__(_975_v49, _dafny.IntOfUint32((_902_v0).Cardinality()), (_this).F24(), _dafny.SeqOf(func() _dafny.Map {
				var _coll28 = _dafny.NewMapBuilder()
				_ = _coll28
				for _iter30 := _dafny.Iterate((_902_v0).Elements()); ; {
					_compr_28, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _977_v50 _dafny.CodePoint
					_977_v50 = interface{}(_compr_28).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_902_v0, _977_v50) {
						_coll28.Add(_977_v50, _937_v23)
					}
				}
				return _coll28.ToMap()
			}()))
			_976_v51 = _nw142
			var _978_v52 _dafny.Set
			_ = _978_v52
			_978_v52 = _dafny.SetOf((_940_v26).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.IntOfUint32((_940_v26).Cardinality()))).Uint32()).(_dafny.Int))
			var _979_v54 _dafny.Map
			_ = _979_v54
			_979_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_978_v52).Union(func() _dafny.Set {
				var _coll29 = _dafny.NewBuilder()
				_ = _coll29
				for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(350), _dafny.IntOfInt64(677))); ; {
					_compr_29, _ok31 := _iter31()
					if !_ok31 {
						break
					}
					var _980_v53 _dafny.Int
					_980_v53 = interface{}(_compr_29).(_dafny.Int)
					if ((_dafny.IntOfInt64(350)).Cmp(_980_v53) <= 0) && ((_980_v53).Cmp(_dafny.IntOfInt64(677)) < 0) {
						_coll29.Add((_980_v53).Times(_937_v23))
					}
				}
				return _coll29.ToSet()
			}())).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("gsixn"))
			var _rhs147 _dafny.Sequence = (func() _dafny.Sequence {
				if (_979_v54).Contains(_976_v51.F33) {
					return (_979_v54).Get(_976_v51.F33).(_dafny.Sequence)
				}
				return _902_v0
			})()
			_ = _rhs147
			var _rhs148 _dafny.Int = (func() _dafny.Int {
				if Companion_Default___.Fm2(_dafny.IntOfUint32((_902_v0).Cardinality()), _dafny.IntOfInt64(189), _903_v1, globalState) {
					return _937_v23
				}
				return (_this).F24()
			})()
			_ = _rhs148
			var _rhs149 bool = false
			_ = _rhs149
			var _lhs113 *GlobalState = globalState
			_ = _lhs113
			var _lhs114 *GlobalState = globalState
			_ = _lhs114
			_lhs113.F17 = _rhs147
			_lhs114.F14 = _rhs148
			_938_v24 = _rhs149
			var _981_v55 _dafny.Array
			_ = _981_v55
			var _nwElement0_28 bool = _938_v24
			_ = _nwElement0_28
			var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(28))
			_ = _nw143
			(_nw143).ArraySet1(_nwElement0_28, 0)
			(_nw143).ArraySet1(_938_v24, 1)
			(_nw143).ArraySet1(_938_v24, 2)
			(_nw143).ArraySet1(_938_v24, 3)
			(_nw143).ArraySet1(_938_v24, 4)
			(_nw143).ArraySet1(_938_v24, 5)
			(_nw143).ArraySet1(_938_v24, 6)
			(_nw143).ArraySet1(_938_v24, 7)
			(_nw143).ArraySet1(((_this).F24()).Cmp((_this).F24()) == 0, 8)
			(_nw143).ArraySet1(_938_v24, 9)
			(_nw143).ArraySet1((_dafny.MultiSetOf(_938_v24, _938_v24)).IsDisjointFrom((_939_v25).Update(false, Companion_Default___.Abs(_976_v51.F33))), 10)
			(_nw143).ArraySet1(_938_v24, 11)
			(_nw143).ArraySet1(_938_v24, 12)
			(_nw143).ArraySet1(_938_v24, 13)
			(_nw143).ArraySet1(_938_v24, 14)
			(_nw143).ArraySet1((_976_v51.F33).Cmp(_976_v51.F33) != 0, 15)
			(_nw143).ArraySet1(_938_v24, 16)
			(_nw143).ArraySet1(!(_938_v24), 17)
			(_nw143).ArraySet1((_974_cf17).Select((Companion_Default___.SafeIndex(_937_v23, _dafny.IntOfUint32((_974_cf17).Cardinality()))).Uint32()).(bool), 18)
			(_nw143).ArraySet1(_938_v24, 19)
			(_nw143).ArraySet1(_938_v24, 20)
			(_nw143).ArraySet1((func() bool {
				if _938_v24 {
					return _938_v24
				}
				return _938_v24
			})(), 21)
			(_nw143).ArraySet1(!(_938_v24), 22)
			(_nw143).ArraySet1(_938_v24, 23)
			(_nw143).ArraySet1(_938_v24, 24)
			(_nw143).ArraySet1(_938_v24, 25)
			(_nw143).ArraySet1(_938_v24, 26)
			(_nw143).ArraySet1(_938_v24, 27)
			_981_v55 = _nw143
			_981_v55 = _981_v55
			var _982_v56 *C0
			_ = _982_v56
			var _nw144 *C0 = New_C0_()
			_ = _nw144
			_nw144.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_902_v0, _902_v0))
			_982_v56 = _nw144
		}
		var _983_v57 _dafny.Map
		_ = _983_v57
		_983_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_902_v0).Cardinality())), _937_v23)
		var _984_v58 _dafny.Map
		_ = _984_v58
		_984_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_937_v23, _983_v57)
		var _985_v59 _dafny.Map
		_ = _985_v59
		_985_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v24, _937_v23)
		var _986_v60 D4
		_ = _986_v60
		_986_v60 = Companion_D4_.Create_DC13_((_984_v58).Cardinality(), _985_v59, _902_v0)
		var _987_v61 _dafny.Map
		_ = _987_v61
		_987_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_986_v60, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(496))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg68 _dafny.Int) interface{} {
				return coer68(arg68)
			}
		}((func(_988_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_989_i9 _dafny.Int) _dafny.CodePoint {
				return _988_v1
			}
		})(_903_v1))))
		(globalState).F1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (func() bool {
				if _938_v24 {
					return _938_v24
				}
				return _938_v24
			})() {
				return (func() _dafny.Sequence {
					if (_987_v61).Contains(_986_v60) {
						return (_987_v61).Get(_986_v60).(_dafny.Sequence)
					}
					return _902_v0
				})()
			}
			return _902_v0
		})()).Cardinality())
		r0 = true
		return r0
	}
}
func (_this *C3) M5(p0 bool, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _990_v0 _dafny.Sequence
		_ = _990_v0
		_990_v0 = _dafny.SeqOf(p0)
		var _991_v1 D3
		_ = _991_v1
		_991_v1 = Companion_D3_.Create_DC9_(_dafny.Companion_Sequence_.Update(_990_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.IntOfUint32((_990_v0).Cardinality()))).Uint32(), p0))
		var _992_i0 _dafny.Int
		_ = _992_i0
		_992_i0 = _dafny.Zero
		{
			var _pat_let_tv6 = p0
			_ = _pat_let_tv6
			for func(_source12 D3) bool {
				if _source12.Is_DC10() {
					return _pat_let_tv6
				} else if _source12.Is_DC11() {
					var _1069___mcc_h7 _dafny.Int = _source12.Get_().(D3_DC11).Cf18
					_ = _1069___mcc_h7
					var _1070___mcc_h8 bool = _source12.Get_().(D3_DC11).Cf19
					_ = _1070___mcc_h8
					var _1071___mcc_h9 bool = _source12.Get_().(D3_DC11).Cf20
					_ = _1071___mcc_h9
					var _1072_cf20 bool = _1071___mcc_h9
					_ = _1072_cf20
					var _1073_cf19 bool = _1070___mcc_h8
					_ = _1073_cf19
					var _1074_cf18 _dafny.Int = _1069___mcc_h7
					_ = _1074_cf18
					return _1073_cf19
				} else {
					var _1075___mcc_h10 _dafny.Sequence = _source12.Get_().(D3_DC9).Cf17
					_ = _1075___mcc_h10
					var _1076_cf17 _dafny.Sequence = _1075___mcc_h10
					_ = _1076_cf17
					return true
				}
			}(_991_v1) {
				{
					if (_992_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_992_i0 = (_992_i0).Plus(_dafny.One)
					var _993_v2 _dafny.CodePoint
					_ = _993_v2
					_993_v2 = _dafny.CodePoint('o')
					var _994_v3 _dafny.Sequence
					_ = _994_v3
					_994_v3 = _dafny.SeqOf(_993_v2, _993_v2, _993_v2)
					var _995_v4 D1
					_ = _995_v4
					_995_v4 = Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Update(_994_v3, (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_994_v3).Cardinality()))).Uint32(), _993_v2))
					var _996_v5 D8
					_ = _996_v5
					_996_v5 = Companion_D8_.Create_DC26_(_995_v4, _994_v3)
					var _source11 D8 = _996_v5
					_ = _source11
					if _source11.Is_DC24() {
						var _997___mcc_h0 _dafny.Int = _source11.Get_().(D8_DC24).Cf37
						_ = _997___mcc_h0
						var _998_cf37 _dafny.Int = _997___mcc_h0
						_ = _998_cf37
						var _999_v6 D0
						_ = _999_v6
						_999_v6 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(980), (_this).F24(), p0)
						var _1000_v7 _dafny.Sequence
						_ = _1000_v7
						_1000_v7 = _dafny.SeqOf((_this).F24())
						var _1001_v8 D1
						_ = _1001_v8
						_1001_v8 = Companion_D1_.Create_DC5_(p0, p0, p0)
						var _1002_v9 D6
						_ = _1002_v9
						_1002_v9 = Companion_D6_.Create_DC18_(p0)
						var _1003_v10 _dafny.Array
						_ = _1003_v10
						var _len0_24 _dafny.Int = _dafny.IntOfInt64(17)
						_ = _len0_24
						var _nw145 _dafny.Array
						_ = _nw145
						if _len0_24.Cmp(_dafny.Zero) == 0 {
							_nw145 = _dafny.NewArray(_len0_24)
						} else {
							var _init24 func(_dafny.Int) bool = (func(_1004_p0 bool) func(_dafny.Int) bool {
								return func(_1005_i1 _dafny.Int) bool {
									return _1004_p0
								}
							})(p0)
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
						_1003_v10 = _nw145
						var _1006_v11 D7
						_ = _1006_v11
						_1006_v11 = Companion_D7_.Create_DC21_(_1001_v8, Companion_D6_.Create_DC19_(_1002_v9), _1003_v10)
						var _1007_v12 _dafny.Map
						_ = _1007_v12
						_1007_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1000_v7, (_1006_v11).Dtor_cf34())
						var _pat_let_tv4 = _990_v0
						_ = _pat_let_tv4
						var _pat_let_tv5 = _990_v0
						_ = _pat_let_tv5
						var _1008_v13 _dafny.Array
						_ = _1008_v13
						var _nwElement0_29 _dafny.Int = _998_cf37
						_ = _nwElement0_29
						var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(27))
						_ = _nw146
						(_nw146).ArraySet1(_nwElement0_29, 0)
						(_nw146).ArraySet1((_dafny.Zero).Minus((func(_pat_let10_0 D0) D0 {
							return func(_1009_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let11_0 bool) D0 {
									return func(_1010_dt__update_hcf3_h0 bool) D0 {
										return Companion_D0_.Create_DC1_((_1009_dt__update__tmp_h0).Dtor_cf1(), (_1009_dt__update__tmp_h0).Dtor_cf2(), _1010_dt__update_hcf3_h0)
									}(_pat_let11_0)
								}((_pat_let_tv4).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_pat_let_tv5).Cardinality()))).Uint32()).(bool))
							}(_pat_let10_0)
						}(_999_v6)).Dtor_cf1()), 1)
						(_nw146).ArraySet1(_998_cf37, 2)
						(_nw146).ArraySet1(_998_cf37, 3)
						(_nw146).ArraySet1(_998_cf37, 4)
						(_nw146).ArraySet1((_dafny.Zero).Minus((_this).F24()), 5)
						(_nw146).ArraySet1((_this).F24(), 6)
						(_nw146).ArraySet1((_this).F24(), 7)
						(_nw146).ArraySet1(_998_cf37, 8)
						(_nw146).ArraySet1((_1007_v12).Cardinality(), 9)
						(_nw146).ArraySet1(_998_cf37, 10)
						(_nw146).ArraySet1(_998_cf37, 11)
						(_nw146).ArraySet1(_998_cf37, 12)
						(_nw146).ArraySet1(_998_cf37, 13)
						(_nw146).ArraySet1(_dafny.IntOfUint32((_994_v3).Cardinality()), 14)
						(_nw146).ArraySet1((_this).F24(), 15)
						(_nw146).ArraySet1((_this).F24(), 16)
						(_nw146).ArraySet1(_dafny.IntOfInt64(597), 17)
						(_nw146).ArraySet1((_this).F24(), 18)
						(_nw146).ArraySet1((_this).F24(), 19)
						(_nw146).ArraySet1((_this).F24(), 20)
						(_nw146).ArraySet1(_998_cf37, 21)
						(_nw146).ArraySet1(_dafny.IntOfInt64(-506), 22)
						(_nw146).ArraySet1((_this).F24(), 23)
						(_nw146).ArraySet1((_this).F24(), 24)
						(_nw146).ArraySet1(_998_cf37, 25)
						(_nw146).ArraySet1((_this).F24(), 26)
						_1008_v13 = _nw146
						var _1011_v14 _dafny.Sequence
						_ = _1011_v14
						_1011_v14 = _dafny.SeqOf(_1008_v13, (func() _dafny.Array {
							if p0 {
								return _1008_v13
							}
							return _1008_v13
						})(), _1008_v13)
						var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1008_v13), 0))
						_ = _index119
						(_1008_v13).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_998_cf37)))).Cardinality()), (_index119).Int())
						var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_1008_v13), 0))
						_ = _index120
						(_1008_v13).ArraySet1(_998_cf37, (_index120).Int())
						var _1012_v15 _dafny.Map
						_ = _1012_v15
						_1012_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_993_v2, (_this).F24())
						var _1013_v16 T0
						_ = _1013_v16
						var _nw147 *C2 = New_C2_()
						_ = _nw147
						_nw147.Ctor__(p0, (_this).F24(), _dafny.SeqOf(_1012_v15))
						_1013_v16 = _nw147
						var _1014_v17 _dafny.Map
						_ = _1014_v17
						_1014_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1013_v16, _1008_v13)
						var _1015_v18 *C0
						_ = _1015_v18
						var _nw148 *C0 = New_C0_()
						_ = _nw148
						_nw148.Ctor__(p0)
						_1015_v18 = _nw148
						var _1016_v19 _dafny.Map
						_ = _1016_v19
						_1016_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1015_v18, (_1015_v18).F34())
						var _1017_v20 _dafny.Map
						_ = _1017_v20
						_1017_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1013_v16).F24(), true)
						var _1018_v21 _dafny.Set
						_ = _1018_v21
						_1018_v21 = _dafny.SetOf(_1000_v7, _1000_v7, _1000_v7)
						var _1019_v22 _dafny.Set
						_ = _1019_v22
						_1019_v22 = _dafny.SetOf(!((_1015_v18).F34()))
						var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1008_v13), 0))
						_ = _index121
						var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_1008_v13), 0))
						_ = _index122
						var _rhs150 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1011_v14, _dafny.SeqOf(_1008_v13)), _1011_v14)
						_ = _rhs150
						var _rhs151 _dafny.Int = ((_this).F24()).Minus((_1014_v17).Cardinality())
						_ = _rhs151
						var _rhs152 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm4(p0, ((_1016_v19).Merge(_1016_v19)).Cardinality(), _993_v2, p0, globalState), (Companion_Default___.SafeIndex((_1000_v7).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1000_v7).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm4(p0, ((_1016_v19).Merge(_1016_v19)).Cardinality(), _993_v2, p0, globalState)).Cardinality()))).Uint32(), _993_v2), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_994_v3).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm4(p0, ((_1016_v19).Merge(_1016_v19)).Cardinality(), _993_v2, p0, globalState), (Companion_Default___.SafeIndex((_1000_v7).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1000_v7).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm4(p0, ((_1016_v19).Merge(_1016_v19)).Cardinality(), _993_v2, p0, globalState)).Cardinality()))).Uint32(), _993_v2)).Cardinality()))).Uint32(), _993_v2)
						_ = _rhs152
						var _rhs153 bool = !((_1015_v18).F34()) || (p0)
						_ = _rhs153
						var _rhs154 _dafny.Int = Companion_Default___.SafeModuloInt((_1000_v7).Select((Companion_Default___.SafeIndex((_dafny.SetOf((_1017_v20).Cardinality(), (_this).F24(), (_this).F24(), (_1018_v21).Cardinality())).Cardinality(), _dafny.IntOfUint32((_1000_v7).Cardinality()))).Uint32()).(_dafny.Int), (_1019_v22).Cardinality())
						_ = _rhs154
						var _lhs115 _dafny.Array = _1008_v13
						_ = _lhs115
						var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1008_v13), 0))
						_ = _lhs116
						var _lhs117 *GlobalState = globalState
						_ = _lhs117
						var _lhs118 _dafny.Array = _1008_v13
						_ = _lhs118
						var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(583), _dafny.ArrayLen((_1008_v13), 0))
						_ = _lhs119
						_1011_v14 = _rhs150
						(_lhs115).ArraySet1(_rhs151, (_lhs116).Int())
						_lhs117.F17 = _rhs152
						r1 = _rhs153
						(_lhs118).ArraySet1(_rhs154, (_lhs119).Int())
						var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1008_v13), 0))
						_ = _index123
						(_1008_v13).ArraySet1((_1013_v16).F24(), (_index123).Int())
						(globalState).F1 = ((Companion_D0_.Create_DC2_(!((_1015_v18).F34()), (_this).F24(), _993_v2)).Dtor_cf5()).Times(_998_cf37)
						(globalState).F2 = !((_1015_v18).F34())
					} else if _source11.Is_DC25() {
						var _1020___mcc_h1 _dafny.Int = _source11.Get_().(D8_DC25).Cf38
						_ = _1020___mcc_h1
						var _1021___mcc_h2 bool = _source11.Get_().(D8_DC25).Cf39
						_ = _1021___mcc_h2
						var _1022___mcc_h3 bool = _source11.Get_().(D8_DC25).Cf40
						_ = _1022___mcc_h3
						var _1023_cf40 bool = _1022___mcc_h3
						_ = _1023_cf40
						var _1024_cf39 bool = _1021___mcc_h2
						_ = _1024_cf39
						var _1025_cf38 _dafny.Int = _1020___mcc_h1
						_ = _1025_cf38
						var _1026_v23 _dafny.Map
						_ = _1026_v23
						_1026_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_cf38, _1023_cf40)
						var _1027_v24 _dafny.Map
						_ = _1027_v24
						_1027_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1024_cf39, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1023_cf40, _1023_cf40, !(!(p0))), (Companion_Default___.SafeIndex(_1025_cf38, _dafny.IntOfUint32((_dafny.SeqOf(_1023_cf40, _1023_cf40, !(!(p0)))).Cardinality()))).Uint32(), (func() bool {
							if (_1026_v23).Contains((_this).F24()) {
								return (_1026_v23).Get((_this).F24()).(bool)
							}
							return _1023_cf40
						})())).Cardinality()))
						var _1028_v25 _dafny.Map
						_ = _1028_v25
						_1028_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_cf38, _dafny.IntOfInt64(679))
						_1027_v24 = (_1027_v24).Update((_1024_cf39) && (p0), Companion_Default___.SafeDivisionInt((_this).F24(), (_1028_v25).Cardinality()))
						(globalState).F1 = Companion_Default___.SafeDivisionInt(_1025_cf38, Companion_Default___.SafeDivisionInt(_1025_cf38, (_this).F24()))
						var _rhs155 bool = p0
						_ = _rhs155
						var _rhs156 bool = _1023_cf40
						_ = _rhs156
						var _rhs157 bool = !_dafny.Companion_Sequence_.Contains(_994_v3, _993_v2)
						_ = _rhs157
						var _rhs158 _dafny.CodePoint = (func() _dafny.CodePoint {
							if true {
								return _dafny.CodePoint('b')
							}
							return _993_v2
						})()
						_ = _rhs158
						var _rhs159 _dafny.Int = (_this).F24()
						_ = _rhs159
						var _lhs120 *GlobalState = globalState
						_ = _lhs120
						var _lhs121 *GlobalState = globalState
						_ = _lhs121
						var _lhs122 *GlobalState = globalState
						_ = _lhs122
						_lhs120.F2 = _rhs155
						_lhs121.F2 = _rhs156
						r0 = _rhs157
						_993_v2 = _rhs158
						_lhs122.F1 = _rhs159
						var _1029_v26 _dafny.Array
						_ = _1029_v26
						var _nwElement0_30 bool = true
						_ = _nwElement0_30
						var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(19))
						_ = _nw149
						(_nw149).ArraySet1(_nwElement0_30, 0)
						(_nw149).ArraySet1(true, 1)
						(_nw149).ArraySet1(p0, 2)
						(_nw149).ArraySet1(_1024_cf39, 3)
						(_nw149).ArraySet1(!(p0), 4)
						(_nw149).ArraySet1(p0, 5)
						(_nw149).ArraySet1(p0, 6)
						(_nw149).ArraySet1(_1024_cf39, 7)
						(_nw149).ArraySet1(false, 8)
						(_nw149).ArraySet1(_1023_cf40, 9)
						(_nw149).ArraySet1(!(p0), 10)
						(_nw149).ArraySet1(_1023_cf40, 11)
						(_nw149).ArraySet1(p0, 12)
						(_nw149).ArraySet1(_1023_cf40, 13)
						(_nw149).ArraySet1(_1023_cf40, 14)
						(_nw149).ArraySet1(_1024_cf39, 15)
						(_nw149).ArraySet1(!(!(true)), 16)
						(_nw149).ArraySet1(p0, 17)
						(_nw149).ArraySet1(true, 18)
						_1029_v26 = _nw149
						var _1030_v27 _dafny.Sequence
						_ = _1030_v27
						_1030_v27 = _dafny.SeqOf(_1029_v26, _1029_v26)
						var _1031_v28 _dafny.Map
						_ = _1031_v28
						_1031_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1030_v27, _1030_v27), (_1027_v24).Cardinality())
						_1031_v28 = (_1031_v28).Update(_1030_v27, (_dafny.Zero).Minus((_dafny.Zero).Minus((_1027_v24).Cardinality())))
					} else if _source11.Is_DC26() {
						var _1032___mcc_h4 D1 = _source11.Get_().(D8_DC26).Cf41
						_ = _1032___mcc_h4
						var _1033___mcc_h5 _dafny.Sequence = _source11.Get_().(D8_DC26).Cf42
						_ = _1033___mcc_h5
						var _1034_cf42 _dafny.Sequence = _1033___mcc_h5
						_ = _1034_cf42
						var _1035_cf41 D1 = _1032___mcc_h4
						_ = _1035_cf41
						var _1036_v29 _dafny.Map
						_ = _1036_v29
						_1036_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _993_v2)
						_1036_v29 = (_1036_v29).Update((_this).F24(), _993_v2)
						_1034_cf42 = _994_v3
						var _1037_v30 _dafny.Map
						_ = _1037_v30
						_1037_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
						var _1038_v31 _dafny.Map
						_ = _1038_v31
						_1038_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F24())
						var _1039_v32 _dafny.Map
						_ = _1039_v32
						_1039_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
						var _1040_v33 _dafny.Sequence
						_ = _1040_v33
						_1040_v33 = _dafny.SeqOf(_dafny.IntOfUint32((_994_v3).Cardinality()), (_this).F24(), (_this).F24(), (_dafny.SetOf(p0, p0)).Cardinality(), (_1039_v32).Cardinality())
						var _1041_v34 D8
						_ = _1041_v34
						_1041_v34 = Companion_D8_.Create_DC25_((_this).F24(), p0, p0)
						var _1042_v35 _dafny.Map
						_ = _1042_v35
						_1042_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm14(_1037_v30, (func() _dafny.Int {
							if (_1038_v31).Contains(p0) {
								return (_1038_v31).Get(p0).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_1040_v33).Cardinality())
						})(), globalState)).Cardinality(), (_1041_v34).Dtor_cf39())
						var _1043_v36 _dafny.Array
						_ = _1043_v36
						var _nw150 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
						_ = _nw150
						_1043_v36 = _nw150
						var _1044_v37 D1
						_ = _1044_v37
						_1044_v37 = Companion_D1_.Create_DC5_(p0, p0, !(p0))
						var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1043_v36), 0))
						_ = _index124
						(_1043_v36).ArraySet1((_1044_v37).Dtor_cf14(), (_index124).Int())
						var _1045_v38 _dafny.Sequence
						_ = _1045_v38
						_1045_v38 = _dafny.SeqOf(_1043_v36, _1043_v36, _1043_v36)
						var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1043_v36), 0))
						_ = _index125
						var _rhs160 _dafny.Map = _1042_v35
						_ = _rhs160
						var _rhs161 bool = !(true) || (p0)
						_ = _rhs161
						var _rhs162 _dafny.Sequence = _1045_v38
						_ = _rhs162
						var _lhs123 _dafny.Array = _1043_v36
						_ = _lhs123
						var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_1043_v36), 0))
						_ = _lhs124
						_1042_v35 = _rhs160
						(_lhs123).ArraySet1(_rhs161, (_lhs124).Int())
						_1045_v38 = _rhs162
						var _1046_v39 _dafny.Map
						_ = _1046_v39
						_1046_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_1040_v33).Cardinality()), _1038_v31, _dafny.UnicodeSeqOfUtf8Bytes("de")), _dafny.IntOfUint32((_1040_v33).Cardinality()))
						_1046_v39 = (_1046_v39).Update(Companion_Default___.Fm25(globalState), (_this).Fm8(globalState))
					} else {
						var _1047___mcc_h6 _dafny.Set = _source11.Get_().(D8_DC23).Cf36
						_ = _1047___mcc_h6
						var _1048_cf36 _dafny.Set = _1047___mcc_h6
						_ = _1048_cf36
						var _1049_v40 _dafny.Array
						_ = _1049_v40
						var _len0_25 _dafny.Int = _dafny.IntOfInt64(25)
						_ = _len0_25
						var _nw151 _dafny.Array
						_ = _nw151
						if _len0_25.Cmp(_dafny.Zero) == 0 {
							_nw151 = _dafny.NewArray(_len0_25)
						} else {
							var _init25 func(_dafny.Int) _dafny.MultiSet = func(_1050_i2 _dafny.Int) _dafny.MultiSet {
								return _dafny.MultiSetOf((_this).F24(), (_this).F24())
							}
							_ = _init25
							var _element0_25 = _init25(_dafny.Zero)
							_ = _element0_25
							_nw151 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
							(_nw151).ArraySet1(_element0_25, 0)
							var _nativeLen0_25 = (_len0_25).Int()
							_ = _nativeLen0_25
							for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
								(_nw151).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
							}
						}
						_1049_v40 = _nw151
						var _1051_v41 _dafny.Array
						_ = _1051_v41
						var _nwElement0_31 _dafny.Array = _1049_v40
						_ = _nwElement0_31
						var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(9))
						_ = _nw152
						(_nw152).ArraySet1(_nwElement0_31, 0)
						(_nw152).ArraySet1(_1049_v40, 1)
						(_nw152).ArraySet1(_1049_v40, 2)
						(_nw152).ArraySet1(_1049_v40, 3)
						(_nw152).ArraySet1(_1049_v40, 4)
						(_nw152).ArraySet1(_1049_v40, 5)
						(_nw152).ArraySet1(_1049_v40, 6)
						(_nw152).ArraySet1(_1049_v40, 7)
						(_nw152).ArraySet1(_1049_v40, 8)
						_1051_v41 = _nw152
						var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_1051_v41), 0))
						_ = _index126
						(_1051_v41).ArraySet1((func() _dafny.Array {
							if !(p0) {
								return _1049_v40
							}
							return _1049_v40
						})(), (_index126).Int())
						var _1052_v42 _dafny.Array
						_ = _1052_v42
						var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
						_ = _nw153
						_1052_v42 = _nw153
						var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))
						_ = _index127
						(_1052_v42).ArraySet1(p0, (_index127).Int())
						var _1053_v43 D1
						_ = _1053_v43
						_1053_v43 = Companion_D1_.Create_DC5_(p0, !(p0), !(p0))
						var _1054_v44 _dafny.Map
						_ = _1054_v44
						_1054_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_994_v3, _1053_v43)
						var _1055_v45 _dafny.Array
						_ = _1055_v45
						var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
						_ = _nw154
						_1055_v45 = _nw154
						var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1055_v45), 0))
						_ = _index128
						(_1055_v45).ArraySet1(_994_v3, (_index128).Int())
						var _1056_v46 _dafny.Sequence
						_ = _1056_v46
						_1056_v46 = _dafny.SeqOf(_1049_v40)
						var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_1051_v41), 0))
						_ = _index129
						var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))
						_ = _index130
						var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1055_v45), 0))
						_ = _index131
						var _rhs163 _dafny.Array = (_1056_v46).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1056_v46).Cardinality()))).Uint32()).(_dafny.Array)
						_ = _rhs163
						var _rhs164 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_994_v3, _994_v3)
						_ = _rhs164
						var _rhs165 bool = (((_dafny.MultiSetOf(p0)).Update(p0, Companion_Default___.Abs((_this).F24()))).Cardinality()).Cmp((_this).F24()) < 0
						_ = _rhs165
						var _rhs166 _dafny.Map = (Companion_Default___.Fm34(p0, globalState)).Merge(_1054_v44)
						_ = _rhs166
						var _rhs167 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_994_v3, _994_v3)
						_ = _rhs167
						var _lhs125 _dafny.Array = _1051_v41
						_ = _lhs125
						var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_1051_v41), 0))
						_ = _lhs126
						var _lhs127 *GlobalState = globalState
						_ = _lhs127
						var _lhs128 _dafny.Array = _1052_v42
						_ = _lhs128
						var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))
						_ = _lhs129
						var _lhs130 _dafny.Array = _1055_v45
						_ = _lhs130
						var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1055_v45), 0))
						_ = _lhs131
						(_lhs125).ArraySet1(_rhs163, (_lhs126).Int())
						_lhs127.F17 = _rhs164
						(_lhs128).ArraySet1(_rhs165, (_lhs129).Int())
						_1054_v44 = _rhs166
						(_lhs130).ArraySet1(_rhs167, (_lhs131).Int())
						var _1057_v47 _dafny.Set
						_ = _1057_v47
						_1057_v47 = _dafny.SetOf((_this).F24())
						var _1058_v48 _dafny.Set
						_ = _1058_v48
						_1058_v48 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1057_v47).Cardinality())).Cardinality())
						var _1059_v49 _dafny.Map
						_ = _1059_v49
						_1059_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
							if p0 {
								return _1057_v47
							}
							return _1057_v47
						})(), _dafny.IntOfInt64(821))
						var _rhs168 _dafny.Map = _1059_v49
						_ = _rhs168
						var _rhs169 _dafny.Int = (_this).F24()
						_ = _rhs169
						var _lhs132 *GlobalState = globalState
						_ = _lhs132
						_1059_v49 = _rhs168
						_lhs132.F9 = _rhs169
						var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))
						_ = _index132
						(_1052_v42).ArraySet1((_1052_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))).Int()).(bool), (_index132).Int())
						var _1060_v50 _dafny.Map
						_ = _1060_v50
						_1060_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))).Int()).(bool), (_1052_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))).Int()).(bool))
						var _1061_v51 _dafny.MultiSet
						_ = _1061_v51
						_1061_v51 = _dafny.MultiSetOf(_1060_v50)
						var _1062_v52 _dafny.MultiSet
						_ = _1062_v52
						_1062_v52 = _dafny.MultiSetOf(Companion_Default___.Fm0(Companion_Default___.Fm0((_1061_v51).Cardinality(), (_this).F24(), globalState), (Companion_Default___.Fm17(_dafny.IntOfInt64(157), _dafny.IntOfInt64(-613), (_this).F24(), (_1052_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))).Int()).(bool), globalState)).Cardinality(), globalState))
						var _1063_v53 _dafny.Map
						_ = _1063_v53
						_1063_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_993_v2, (_this).F24())
						var _1064_v54 _dafny.Map
						_ = _1064_v54
						_1064_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1052_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_1052_v42), 0))).Int()).(bool), (_this).F25())
						var _1065_v55 *C1
						_ = _1065_v55
						var _nw155 *C1 = New_C1_()
						_ = _nw155
						_nw155.Ctor__(_dafny.SeqOf(_1062_v52), (_this).F24(), Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F24()), (func() _dafny.Sequence {
							if p0 {
								return _dafny.SeqOf(_1063_v53)
							}
							return (func() _dafny.Sequence {
								if (_1064_v54).Contains(p0) {
									return (_1064_v54).Get(p0).(_dafny.Sequence)
								}
								return (_this).F25()
							})()
						})())
						_1065_v55 = _nw155
					}
					var _1066_v56 _dafny.Map
					_ = _1066_v56
					_1066_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F24())), p0)
					var _1067_v57 *C1
					_ = _1067_v57
					var _nw156 *C1 = New_C1_()
					_ = _nw156
					_nw156.Ctor__(_dafny.SeqOf(_dafny.MultiSetOf((_1066_v56).Cardinality())), (_this).F24(), _dafny.IntOfInt64(939), _dafny.Companion_Sequence_.Concatenate((_this).F25(), (_this).F25()))
					_1067_v57 = _nw156
					(globalState).F21 = _990_v0
					var _1068_v58 _dafny.Map
					_ = _1068_v58
					_1068_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_this).F24(), (_dafny.Zero).Minus((_this).F24()), _993_v2, globalState), (_1067_v57.F33).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_990_v0, p0)).Cardinality()) != 0)
					_1068_v58 = (_1068_v58).Update(p0, p0)
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _1077_v59 _dafny.Array
		_ = _1077_v59
		var _len0_26 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_26
		var _nw157 _dafny.Array
		_ = _nw157
		if _len0_26.Cmp(_dafny.Zero) == 0 {
			_nw157 = _dafny.NewArray(_len0_26)
		} else {
			var _init26 func(_dafny.Int) _dafny.Map = (func(_1078_p0 bool) func(_dafny.Int) _dafny.Map {
				return func(_1079_i3 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1078_p0, _1078_p0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1078_p0, _1078_p0))
				}
			})(p0)
			_ = _init26
			var _element0_26 = _init26(_dafny.Zero)
			_ = _element0_26
			_nw157 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
			(_nw157).ArraySet1(_element0_26, 0)
			var _nativeLen0_26 = (_len0_26).Int()
			_ = _nativeLen0_26
			for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
				(_nw157).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
			}
		}
		_1077_v59 = _nw157
		var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1077_v59), 0))
		_ = _index133
		(_1077_v59).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), (_index133).Int())
		var _1080_v60 _dafny.Set
		_ = _1080_v60
		_1080_v60 = _dafny.SetOf(_dafny.CodePoint('i'))
		var _1081_v61 _dafny.Array
		_ = _1081_v61
		var _nw158 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw158
		_1081_v61 = _nw158
		var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))
		_ = _index134
		(_1081_v61).ArraySet1(!(p0) || (p0), (_index134).Int())
		var _1082_v62 _dafny.CodePoint
		_ = _1082_v62
		_1082_v62 = _dafny.CodePoint('j')
		var _1083_v63 _dafny.Map
		_ = _1083_v63
		_1083_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(p0)), !(!(p0)))
		var _1084_v64 _dafny.MultiSet
		_ = _1084_v64
		_1084_v64 = _dafny.MultiSetOf((_this).F24(), (_this).F24())
		var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1077_v59), 0))
		_ = _index135
		var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))
		_ = _index136
		var _rhs170 _dafny.Map = _1083_v63
		_ = _rhs170
		var _rhs171 _dafny.Set = _1080_v60
		_ = _rhs171
		var _rhs172 bool = ((_dafny.Zero).Minus((_this).F24())).Cmp((_this).F24()) >= 0
		_ = _rhs172
		var _rhs173 _dafny.Int = ((_dafny.MultiSetOf((_this).F24())).Intersection((_1084_v64).Intersection(_1084_v64))).Cardinality()
		_ = _rhs173
		var _rhs174 _dafny.CodePoint = _1082_v62
		_ = _rhs174
		var _lhs133 _dafny.Array = _1077_v59
		_ = _lhs133
		var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1077_v59), 0))
		_ = _lhs134
		var _lhs135 _dafny.Array = _1081_v61
		_ = _lhs135
		var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))
		_ = _lhs136
		var _lhs137 *GlobalState = globalState
		_ = _lhs137
		(_lhs133).ArraySet1(_rhs170, (_lhs134).Int())
		_1080_v60 = _rhs171
		(_lhs135).ArraySet1(_rhs172, (_lhs136).Int())
		_lhs137.F9 = _rhs173
		_1082_v62 = _rhs174
		var _1085_v65 _dafny.Map
		_ = _1085_v65
		_1085_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), Companion_Default___.Fm15((_this).F24(), p0, (_this).F24(), (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool), globalState))
		_1085_v65 = _1085_v65
		if !(((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)) == ((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))) {
			var _1086_v66 D3
			_ = _1086_v66
			_1086_v66 = Companion_D3_.Create_DC11_(((_this).F24()).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality())), p0, true)
			var _1087_v67 _dafny.Sequence
			_ = _1087_v67
			_1087_v67 = _dafny.UnicodeSeqOfUtf8Bytes("fidp")
			var _rhs175 _dafny.Array = _1081_v61
			_ = _rhs175
			var _rhs176 D3 = _1086_v66
			_ = _rhs176
			var _rhs177 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(997))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}((func(_1088_v62 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1089_i4 _dafny.Int) _dafny.CodePoint {
					return _1088_v62
				}
			})(_1082_v62))), _1087_v67)
			_ = _rhs177
			var _lhs138 *GlobalState = globalState
			_ = _lhs138
			_1081_v61 = _rhs175
			_1086_v66 = _rhs176
			_lhs138.F17 = _rhs177
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))
			_ = _index137
			(_1081_v61).ArraySet1(!(!((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)) || (!(p0))) || (p0), (_index137).Int())
			(globalState).F2 = (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)
			var _1090_v68 _dafny.Sequence
			_ = _1090_v68
			_1090_v68 = _dafny.SeqOf((_this).F24())
			var _1091_v69 D4
			_ = _1091_v69
			_1091_v69 = Companion_D4_.Create_DC14_(Companion_D4_.Create_DC12_(_1090_v68))
			var _1092_v70 _dafny.Map
			_ = _1092_v70
			_1092_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1091_v69, _1087_v67)
			var _1093_v71 D7
			_ = _1093_v71
			_1093_v71 = Companion_D7_.Create_DC20_(_1092_v70)
			var _1094_v72 D7
			_ = _1094_v72
			_1094_v72 = Companion_D7_.Create_DC22_(_1093_v71)
			_1094_v72 = _1094_v72
			var _1095_v73 _dafny.Map
			_ = _1095_v73
			_1095_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
			var _1096_v74 _dafny.Map
			_ = _1096_v74
			_1096_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _dafny.IntOfUint32((_1087_v67).Cardinality()))
			var _1097_v75 D4
			_ = _1097_v75
			_1097_v75 = Companion_D4_.Create_DC13_((_this).F24(), _1096_v74, _1087_v67)
			var _1098_v76 _dafny.Set
			_ = _1098_v76
			_1098_v76 = _dafny.SetOf(_dafny.IntOfInt64(862))
			var _1099_v77 _dafny.MultiSet
			_ = _1099_v77
			_1099_v77 = _dafny.MultiSetOf(p0, p0, (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))
			_ = _index138
			var _rhs178 bool = ((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)) && (!(((_1095_v73).Cardinality()).Cmp((_this).F24()) <= 0))
			_ = _rhs178
			var _rhs179 _dafny.Int = (_1097_v75).Dtor_cf22()
			_ = _rhs179
			var _rhs180 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm0((_1098_v76).Cardinality(), (func() _dafny.Int {
				if (_1099_v77).Contains((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)) {
					return (_1099_v77).Multiplicity((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
				}
				return (_this).F24()
			})(), globalState))
			_ = _rhs180
			var _rhs181 bool = (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)
			_ = _rhs181
			var _rhs182 _dafny.Int = ((_this).F24()).Minus(((_this).F24()).Plus((_this).F24()))
			_ = _rhs182
			var _lhs139 _dafny.Array = _1081_v61
			_ = _lhs139
			var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))
			_ = _lhs140
			var _lhs141 *GlobalState = globalState
			_ = _lhs141
			var _lhs142 *GlobalState = globalState
			_ = _lhs142
			var _lhs143 *GlobalState = globalState
			_ = _lhs143
			(_lhs139).ArraySet1(_rhs178, (_lhs140).Int())
			_lhs141.F14 = _rhs179
			_lhs142.F9 = _rhs180
			r0 = _rhs181
			_lhs143.F1 = _rhs182
		} else {
			(globalState).F2 = ((func() bool {
				if p0 {
					return p0
				}
				return (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)
			})()) && ((func() bool {
				if p0 {
					return p0
				}
				return Companion_Default___.Fm2(_dafny.IntOfUint32((_990_v0).Cardinality()), (_this).F24(), _dafny.CodePoint('w'), globalState)
			})())
			var _1100_v78 _dafny.Sequence
			_ = _1100_v78
			_1100_v78 = _dafny.UnicodeSeqOfUtf8Bytes("omqcl")
			if (func() bool {
				if !((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)) {
					return p0
				}
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_1100_v78, _1100_v78)
			})() {
				var _rhs183 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1100_v78, _dafny.UnicodeSeqOfUtf8Bytes("sqal"))).Cardinality())).Plus((_this).F24())
				_ = _rhs183
				var _rhs184 bool = p0
				_ = _rhs184
				var _lhs144 *GlobalState = globalState
				_ = _lhs144
				var _lhs145 *GlobalState = globalState
				_ = _lhs145
				_lhs144.F1 = _rhs183
				_lhs145.F2 = _rhs184
				var _1101_v79 _dafny.Map
				_ = _1101_v79
				_1101_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(34), !(true))
				var _1102_v80 _dafny.Set
				_ = _1102_v80
				_1102_v80 = _dafny.SetOf(Companion_Default___.SafeDivisionInt((_1101_v79).Cardinality(), (_this).F24()))
				_1102_v80 = _1102_v80
				var _1103_v81 _dafny.Array
				_ = _1103_v81
				var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
				_ = _nw159
				_1103_v81 = _nw159
				var _1104_v82 _dafny.Map
				_ = _1104_v82
				_1104_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1083_v63, (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
				var _rhs185 _dafny.Array = _1103_v81
				_ = _rhs185
				var _rhs186 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if !(((_this).F24()).Cmp((_dafny.Zero).Minus((_this).F24())) != 0) {
						return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(294))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg70 _dafny.Int) interface{} {
								return coer70(arg70)
							}
						}((func(_1105_v61 _dafny.Array) func(_dafny.Int) _dafny.Map {
							return func(_1106_i5 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1105_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1105_v61), 0))).Int()).(bool), (_this).F24())
							}
						})(_1081_v61)))).Cardinality())
					}
					return ((_1104_v82).Merge(_1104_v82)).Cardinality()
				})())
				_ = _rhs186
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				_1103_v81 = _rhs185
				_lhs146.F1 = _rhs186
				var _1107_v83 _dafny.Map
				_ = _1107_v83
				_1107_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1081_v61)
				_1107_v83 = (_1107_v83).Update(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nnv")).Cardinality()), _1081_v61)
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))
				_ = _index139
				(_1081_v61).ArraySet1((_990_v0).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F24()), _dafny.IntOfUint32((_990_v0).Cardinality()))).Uint32()).(bool), (_index139).Int())
			} else {
				var _1108_v84 _dafny.Map
				_ = _1108_v84
				_1108_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool), (_this).F24())
				_1108_v84 = _1108_v84
				(globalState).F18 = (_this).F24()
				var _1109_v85 _dafny.Array
				_ = _1109_v85
				var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw160
				_1109_v85 = _nw160
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_1109_v85), 0))
				_ = _index140
				(_1109_v85).ArraySet1((func() _dafny.Int {
					if (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool) {
						return (_dafny.Zero).Minus((_this).F24())
					}
					return (_this).F24()
				})(), (_index140).Int())
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_1109_v85), 0))
				_ = _index141
				(_1109_v85).ArraySet1((_this).F24(), (_index141).Int())
				(globalState).F2 = p0
				(globalState).F14 = (_1109_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_1109_v85), 0))).Int()).(_dafny.Int)
			}
			(globalState).F9 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F24())), (_this).F24())
			if p0 {
				var _1110_v86 _dafny.Map
				_ = _1110_v86
				_1110_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1082_v62, p0)
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))
				_ = _index142
				(_1081_v61).ArraySet1(!((func() bool {
					if (_1110_v86).Contains(_1082_v62) {
						return (_1110_v86).Get(_1082_v62).(bool)
					}
					return p0
				})()), (_index142).Int())
				(globalState).F9 = (_this).F24()
				var _1111_v87 _dafny.Array
				_ = _1111_v87
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_27
				var _nw161 _dafny.Array
				_ = _nw161
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw161 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Sequence = (func(_1112_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1113_i6 _dafny.Int) _dafny.Sequence {
							return _1112_v0
						}
					})(_990_v0)
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
				_1111_v87 = _nw161
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_1111_v87), 0))
				_ = _index143
				(_1111_v87).ArraySet1(_990_v0, (_index143).Int())
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_1111_v87), 0))
				_ = _index144
				(_1111_v87).ArraySet1(_dafny.Companion_Sequence_.Update(_990_v0, (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_990_v0).Cardinality()))).Uint32(), false), (_index144).Int())
				var _1114_v88 _dafny.Map
				_ = _1114_v88
				_1114_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool), (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool)), _1081_v61)
				_1114_v88 = _1114_v88
				(globalState).F9 = (_this).F24()
			} else {
				var _1115_v89 _dafny.Map
				_ = _1115_v89
				_1115_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1082_v62)
				_1082_v62 = (func() _dafny.CodePoint {
					if (_1115_v89).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_990_v0, _990_v0)).Cardinality())) {
						return (_1115_v89).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_990_v0, _990_v0)).Cardinality())).(_dafny.CodePoint)
					}
					return _1082_v62
				})()
				r1 = p0
				var _1116_v90 _dafny.Map
				_ = _1116_v90
				_1116_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_this).F24(), (_this).Fm8(globalState)), (func() bool {
					if (_1083_v63).Contains(p0) {
						return (_1083_v63).Get(p0).(bool)
					}
					return p0
				})())
				_1116_v90 = (_1116_v90).Update((_this).F24(), (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
				var _1117_v91 *C2
				_ = _1117_v91
				var _nw162 *C2 = New_C2_()
				_ = _nw162
				_nw162.Ctor__(p0, Companion_Default___.SafeModuloInt((_this).F24(), (_this).F24()), (_this).F25())
				_1117_v91 = _nw162
				var _1118_v92 _dafny.Map
				_ = _1118_v92
				_1118_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1082_v62, !(true))
				r0 = (func() bool {
					if (_1118_v92).Contains(_1082_v62) {
						return (_1118_v92).Get(_1082_v62).(bool)
					}
					return ((_this).F24()).Cmp(_dafny.IntOfInt64(123)) > 0
				})()
			}
			var _1119_v93 _dafny.Set
			_ = _1119_v93
			_1119_v93 = _dafny.SetOf((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool), (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
			var _1120_v94 D0
			_ = _1120_v94
			_1120_v94 = Companion_D0_.Create_DC1_((_1119_v93).Cardinality(), (_this).F24(), (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
			var _1121_v95 *C0
			_ = _1121_v95
			var _nw163 *C0 = New_C0_()
			_ = _nw163
			_nw163.Ctor__((_1120_v94).Dtor_cf3())
			_1121_v95 = _nw163
		}
		_1081_v61 = _1081_v61
		var _1122_i7 _dafny.Int
		_ = _1122_i7
		_1122_i7 = _dafny.Zero
		{
			for ((_this).F24()).Cmp((_this).F24()) != 0 {
				{
					if (_1122_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_1122_i7 = (_1122_i7).Plus(_dafny.One)
					(globalState).F1 = (((_this).F24()).Plus((_this).F24())).Plus((_this).F24())
					var _1123_v96 _dafny.Set
					_ = _1123_v96
					_1123_v96 = _dafny.SetOf(true)
					_1123_v96 = _1123_v96
					var _1124_v97 _dafny.Map
					_ = _1124_v97
					_1124_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
					_1124_v97 = _1124_v97
					var _1125_v98 _dafny.MultiSet
					_ = _1125_v98
					_1125_v98 = _dafny.MultiSetOf((_1081_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1081_v61), 0))).Int()).(bool))
					r1 = (_dafny.MultiSetFromSeq(_990_v0)).IsProperSubsetOf(_1125_v98)
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _1126_v99 _dafny.Sequence
		_ = _1126_v99
		_1126_v99 = _dafny.UnicodeSeqOfUtf8Bytes("lq")
		var _1127_v100 _dafny.Set
		_ = _1127_v100
		_1127_v100 = _dafny.SetOf(_dafny.IntOfUint32((_1126_v99).Cardinality()))
		r0 = (((Companion_D6_.Create_DC16_(_1127_v100)).Dtor_cf27()).Union(_1127_v100)).IsDisjointFrom((_dafny.SetOf(_dafny.IntOfInt64(17))).Difference(_dafny.SetOf((_this).F24())))
		r1 = !((p0) == (p0))
		r2 = (_this).F24()
		return r0, r1, r2
	}
}
func (_this *C3) M6(globalState *GlobalState) {
	{
		var _1128_v0 _dafny.Map
		_ = _1128_v0
		_1128_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F24()).Plus((_this).F24()), true)
		var _1129_v1 bool
		_ = _1129_v1
		_1129_v1 = false
		_1128_v0 = (_1128_v0).Update((_this).F24(), _1129_v1)
		(globalState).F2 = !((((_this).F24()).Minus((_this).F24())).Cmp((_this).F24()) == 0)
		var _1130_v2 D1
		_ = _1130_v2
		_1130_v2 = Companion_D1_.Create_DC5_(false, _1129_v1, _1129_v1)
		var _1131_v3 D6
		_ = _1131_v3
		_1131_v3 = Companion_D6_.Create_DC17_(_1129_v1)
		var _1132_v4 D6
		_ = _1132_v4
		_1132_v4 = Companion_D6_.Create_DC19_(_1131_v3)
		var _1133_v5 _dafny.Sequence
		_ = _1133_v5
		_1133_v5 = _dafny.SeqOf(_1129_v1, _1129_v1)
		var _1134_v6 _dafny.Array
		_ = _1134_v6
		var _nwElement0_32 bool = _1129_v1
		_ = _nwElement0_32
		var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(12))
		_ = _nw164
		(_nw164).ArraySet1(_nwElement0_32, 0)
		(_nw164).ArraySet1(_1129_v1, 1)
		(_nw164).ArraySet1(_1129_v1, 2)
		(_nw164).ArraySet1(_1129_v1, 3)
		(_nw164).ArraySet1(!(_1129_v1), 4)
		(_nw164).ArraySet1((_1133_v5).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1133_v5).Cardinality()))).Uint32()).(bool), 5)
		(_nw164).ArraySet1(_1129_v1, 6)
		(_nw164).ArraySet1(_1129_v1, 7)
		(_nw164).ArraySet1(_1129_v1, 8)
		(_nw164).ArraySet1(_1129_v1, 9)
		(_nw164).ArraySet1(_1129_v1, 10)
		(_nw164).ArraySet1(_1129_v1, 11)
		_1134_v6 = _nw164
		var _1135_v7 D7
		_ = _1135_v7
		_1135_v7 = Companion_D7_.Create_DC21_(_1130_v2, _1132_v4, _1134_v6)
		var _1136_v8 _dafny.CodePoint
		_ = _1136_v8
		_1136_v8 = _dafny.CodePoint('k')
		var _pat_let_tv7 = _1136_v8
		_ = _pat_let_tv7
		var _pat_let_tv8 = globalState
		_ = _pat_let_tv8
		var _source13 D7 = func(_pat_let12_0 D7) D7 {
			return func(_1137_dt__update__tmp_h0 D7) D7 {
				return func(_pat_let13_0 D6) D7 {
					return func(_1138_dt__update_hcf33_h0 D6) D7 {
						return Companion_D7_.Create_DC21_((_1137_dt__update__tmp_h0).Dtor_cf32(), _1138_dt__update_hcf33_h0, (_1137_dt__update__tmp_h0).Dtor_cf34())
					}(_pat_let13_0)
				}(Companion_D6_.Create_DC19_(Companion_D6_.Create_DC18_(Companion_Default___.Fm2((_this).F24(), (_this).F24(), _pat_let_tv7, _pat_let_tv8))))
			}(_pat_let12_0)
		}(_1135_v7)
		_ = _source13
		if _source13.Is_DC21() {
			var _1139___mcc_h0 D1 = _source13.Get_().(D7_DC21).Cf32
			_ = _1139___mcc_h0
			var _1140___mcc_h1 D6 = _source13.Get_().(D7_DC21).Cf33
			_ = _1140___mcc_h1
			var _1141___mcc_h2 _dafny.Array = _source13.Get_().(D7_DC21).Cf34
			_ = _1141___mcc_h2
			var _1142_cf34 _dafny.Array = _1141___mcc_h2
			_ = _1142_cf34
			var _1143_cf33 D6 = _1140___mcc_h1
			_ = _1143_cf33
			var _1144_cf32 D1 = _1139___mcc_h0
			_ = _1144_cf32
			(globalState).F2 = (_1129_v1) && (_1129_v1)
			(globalState).F2 = _1129_v1
			var _1145_v9 _dafny.Sequence
			_ = _1145_v9
			_1145_v9 = _dafny.SeqOf(_1136_v8, _1136_v8)
			var _1146_v10 _dafny.Set
			_ = _1146_v10
			_1146_v10 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_1145_v9).Cardinality())), (_this).F24(), (_this).F24())
			var _1147_v11 D6
			_ = _1147_v11
			_1147_v11 = Companion_D6_.Create_DC18_(_1129_v1)
			var _1148_v12 _dafny.Sequence
			_ = _1148_v12
			_1148_v12 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(291))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg71 _dafny.Int) interface{} {
					return coer71(arg71)
				}
			}((func(_1149_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1150_i0 _dafny.Int) _dafny.CodePoint {
					return _1149_v8
				}
			})(_1136_v8))))
			var _1151_v13 _dafny.MultiSet
			_ = _1151_v13
			_1151_v13 = _dafny.MultiSetOf(_1129_v1)
			var _rhs187 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1148_v12, (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1148_v12).Cardinality()))).Uint32(), _1145_v9)).Cardinality())
			_ = _rhs187
			var _rhs188 _dafny.Set = _1146_v10
			_ = _rhs188
			var _rhs189 _dafny.Sequence = _1133_v5
			_ = _rhs189
			var _rhs190 D6 = func(_pat_let14_0 D6) D6 {
				return func(_1152_dt__update__tmp_h1 D6) D6 {
					return func(_pat_let15_0 bool) D6 {
						return func(_1153_dt__update_hcf29_h0 bool) D6 {
							return Companion_D6_.Create_DC18_(_1153_dt__update_hcf29_h0)
						}(_pat_let15_0)
					}(true)
				}(_pat_let14_0)
			}(_1147_v11)
			_ = _rhs190
			var _rhs191 _dafny.Int = Companion_Default___.Fm0(((_this).F24()).Minus(((_1151_v13).Update(_1129_v1, Companion_Default___.Abs((_this).F24()))).Cardinality()), _dafny.IntOfInt64(600), globalState)
			_ = _rhs191
			var _lhs147 *GlobalState = globalState
			_ = _lhs147
			var _lhs148 *GlobalState = globalState
			_ = _lhs148
			var _lhs149 *GlobalState = globalState
			_ = _lhs149
			_lhs147.F18 = _rhs187
			_1146_v10 = _rhs188
			_lhs148.F20 = _rhs189
			_1147_v11 = _rhs190
			_lhs149.F18 = _rhs191
			var _1154_v14 _dafny.Array
			_ = _1154_v14
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_28
			var _nw165 _dafny.Array
			_ = _nw165
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw165 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Int = func(_1155_i1 _dafny.Int) _dafny.Int {
					return (_1155_i1).Plus((_this).F24())
				}
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw165 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw165).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw165).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1154_v14 = _nw165
			var _1156_v15 _dafny.Map
			_ = _1156_v15
			_1156_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1154_v14, _1129_v1)
			var _1157_v16 _dafny.Sequence
			_ = _1157_v16
			_1157_v16 = _dafny.SeqOf(_1146_v10)
			var _1158_v17 _dafny.Map
			_ = _1158_v17
			_1158_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1145_v9, _1157_v16)
			var _1159_v18 _dafny.Sequence
			_ = _1159_v18
			_1159_v18 = _dafny.SeqOf(_dafny.IntOfInt64(57))
			var _rhs192 _dafny.Int = (((_this).F24()).Minus((_this).F24())).Plus((_this).Fm8(globalState))
			_ = _rhs192
			var _rhs193 _dafny.CodePoint = Companion_Default___.Fm3(globalState)
			_ = _rhs193
			var _rhs194 _dafny.Map = _1156_v15
			_ = _rhs194
			var _rhs195 _dafny.Sequence = (func() _dafny.Sequence {
				if (_1158_v17).Contains(_1145_v9) {
					return (_1158_v17).Get(_1145_v9).(_dafny.Sequence)
				}
				return Companion_Default___.Fm35(!(_1129_v1), _1129_v1, globalState)
			})()
			_ = _rhs195
			var _rhs196 bool = (_dafny.MultiSetFromSeq(_1159_v18)).IsSubsetOf(_dafny.MultiSetOf((_1128_v0).Cardinality(), (_this).F24(), (_this).F24(), _dafny.IntOfInt64(709)))
			_ = _rhs196
			var _lhs150 *GlobalState = globalState
			_ = _lhs150
			var _lhs151 *GlobalState = globalState
			_ = _lhs151
			_lhs150.F1 = _rhs192
			_1136_v8 = _rhs193
			_1156_v15 = _rhs194
			_1157_v16 = _rhs195
			_lhs151.F2 = _rhs196
		} else if _source13.Is_DC20() {
			var _1160___mcc_h3 _dafny.Map = _source13.Get_().(D7_DC20).Cf31
			_ = _1160___mcc_h3
			var _1161_cf31 _dafny.Map = _1160___mcc_h3
			_ = _1161_cf31
			var _1162_v19 _dafny.Array
			_ = _1162_v19
			var _nw166 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw166
			_1162_v19 = _nw166
			var _1163_v20 _dafny.MultiSet
			_ = _1163_v20
			_1163_v20 = _dafny.MultiSetOf(_1129_v1, _1129_v1, _1129_v1)
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1162_v19), 0))
			_ = _index145
			(_1162_v19).ArraySet1(((_this).F24()).Minus((_1163_v20).Cardinality()), (_index145).Int())
			var _1164_v21 _dafny.Sequence
			_ = _1164_v21
			_1164_v21 = _dafny.UnicodeSeqOfUtf8Bytes("gsj")
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1162_v19), 0))
			_ = _index146
			var _rhs197 _dafny.Sequence = _1164_v21
			_ = _rhs197
			var _rhs198 _dafny.Int = (_this).F24()
			_ = _rhs198
			var _lhs152 *GlobalState = globalState
			_ = _lhs152
			var _lhs153 _dafny.Array = _1162_v19
			_ = _lhs153
			var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1162_v19), 0))
			_ = _lhs154
			_lhs152.F17 = _rhs197
			(_lhs153).ArraySet1(_rhs198, (_lhs154).Int())
			_1129_v1 = _1129_v1
			(globalState).F2 = _1129_v1
			(globalState).F9 = ((_1162_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1162_v19), 0))).Int()).(_dafny.Int)).Plus((_1162_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1162_v19), 0))).Int()).(_dafny.Int))
		} else {
			var _1165___mcc_h4 D7 = _source13.Get_().(D7_DC22).Cf35
			_ = _1165___mcc_h4
			var _1166_cf35 D7 = _1165___mcc_h4
			_ = _1166_cf35
			var _1167_v22 _dafny.Map
			_ = _1167_v22
			_1167_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v8, (_this).F24())
			var _1168_v23 *C2
			_ = _1168_v23
			var _nw167 *C2 = New_C2_()
			_ = _nw167
			_nw167.Ctor__(_1129_v1, _dafny.IntOfInt64(-718), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1167_v22), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_dafny.SeqOf(_1167_v22)).Cardinality()))).Uint32(), _1167_v22))
			_1168_v23 = _nw167
			var _1169_v24 _dafny.Sequence
			_ = _1169_v24
			_1169_v24 = _dafny.SeqOf((_this).F24())
			var _1170_v25 _dafny.Sequence
			_ = _1170_v25
			_1170_v25 = _dafny.UnicodeSeqOfUtf8Bytes("qe")
			var _1171_v27 _dafny.Map
			_ = _1171_v27
			_1171_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F24())).Cardinality()))
			var _1172_v28 _dafny.MultiSet
			_ = _1172_v28
			_1172_v28 = _dafny.MultiSetOf((_this).F24())
			var _1173_v29 _dafny.Array
			_ = _1173_v29
			var _nwElement0_33 _dafny.Int = (_this).F24()
			_ = _nwElement0_33
			var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(13))
			_ = _nw168
			(_nw168).ArraySet1(_nwElement0_33, 0)
			(_nw168).ArraySet1((_this).F24(), 1)
			(_nw168).ArraySet1((_1169_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1170_v25).Cardinality()), _dafny.IntOfUint32((_1169_v24).Cardinality()))).Uint32()).(_dafny.Int), 2)
			(_nw168).ArraySet1((func() _dafny.Map {
				var _coll30 = _dafny.NewMapBuilder()
				_ = _coll30
				for _iter32 := _dafny.Iterate((_1171_v27).Keys().Elements()); ; {
					_compr_30, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _1174_v26 _dafny.Int
					_1174_v26 = interface{}(_compr_30).(_dafny.Int)
					if (_1171_v27).Contains(_1174_v26) {
						_coll30.Add((_1174_v26).Times(_dafny.IntOfUint32((_1133_v5).Cardinality())), _dafny.IntOfInt64(-366))
					}
				}
				return _coll30.ToMap()
			}()).Cardinality(), 3)
			(_nw168).ArraySet1((_this).F24(), 4)
			(_nw168).ArraySet1((_this).F24(), 5)
			(_nw168).ArraySet1((_this).F24(), 6)
			(_nw168).ArraySet1((_this).F24(), 7)
			(_nw168).ArraySet1(_dafny.IntOfInt64(495), 8)
			(_nw168).ArraySet1((_this).F24(), 9)
			(_nw168).ArraySet1(_dafny.IntOfInt64(488), 10)
			(_nw168).ArraySet1((_this).F24(), 11)
			(_nw168).ArraySet1((_1172_v28).Cardinality(), 12)
			_1173_v29 = _nw168
			var _1175_v30 _dafny.Map
			_ = _1175_v30
			_1175_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("i")).Cardinality()), _1173_v29)
			var _1176_v31 D0
			_ = _1176_v31
			_1176_v31 = Companion_D0_.Create_DC0_(_1173_v29)
			var _1177_v32 _dafny.Set
			_ = _1177_v32
			_1177_v32 = _dafny.SetOf(_1173_v29, (func() _dafny.Array {
				if (_1175_v30).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1176_v31, _1173_v29)).Cardinality()) {
					return (_1175_v30).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1176_v31, _1173_v29)).Cardinality()).(_dafny.Array)
				}
				return _1173_v29
			})(), _1173_v29, _1173_v29, _1173_v29)
			_1177_v32 = ((_1177_v32).Difference(_1177_v32)).Union(_1177_v32)
			(globalState).F21 = _1133_v5
			var _1178_v33 _dafny.Map
			_ = _1178_v33
			_1178_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1168_v23.F31, ((_dafny.Zero).Minus((_this).F24())).Cmp((_dafny.Zero).Minus((_this).F24())) > 0)
			_1178_v33 = (_1178_v33).Update(!(_1168_v23.F31), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("pfodpobt"), _1170_v25))
		}
		var _1179_v34 _dafny.Sequence
		_ = _1179_v34
		_1179_v34 = _dafny.UnicodeSeqOfUtf8Bytes("u")
		var _1180_v35 _dafny.Map
		_ = _1180_v35
		_1180_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1129_v1, _1179_v34)
		var _1181_v36 _dafny.Map
		_ = _1181_v36
		_1181_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if _1129_v1 {
				return (_this).F24()
			}
			return (_this).F24()
		})(), (_1180_v35).Merge(_1180_v35))
		_1181_v36 = (_1181_v36).Update((_this).F24(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1129_v1, _1179_v34))
		(globalState).F17 = _1179_v34
		var _1182_v38 _dafny.Map
		_ = _1182_v38
		_1182_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1129_v1, _dafny.IntOfUint32((_dafny.SeqOf(_1129_v1, _1129_v1, false, !(_1129_v1))).Cardinality()))
		var _1183_v39 D4
		_ = _1183_v39
		_1183_v39 = Companion_D4_.Create_DC14_(Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_1179_v34).Cardinality()), _1182_v38, _1179_v34))
		var _1184_v40 _dafny.Set
		_ = _1184_v40
		_1184_v40 = _dafny.SetOf(_1183_v39)
		var _1185_v41 D7
		_ = _1185_v41
		_1185_v41 = Companion_D7_.Create_DC20_(func() _dafny.Map {
			var _coll31 = _dafny.NewMapBuilder()
			_ = _coll31
			for _iter33 := _dafny.Iterate((_1184_v40).Elements()); ; {
				_compr_31, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _1186_v37 D4
				_1186_v37 = interface{}(_compr_31).(D4)
				if (_1184_v40).Contains(_1186_v37) {
					_coll31.Add(_1186_v37, _1179_v34)
				}
			}
			return _coll31.ToMap()
		}())
		var _1187_i2 _dafny.Int
		_ = _1187_i2
		_1187_i2 = _dafny.Zero
		{
			var _pat_let_tv9 = _1136_v8
			_ = _pat_let_tv9
			var _pat_let_tv10 = _1136_v8
			_ = _pat_let_tv10
			var _pat_let_tv11 = _1136_v8
			_ = _pat_let_tv11
			var _pat_let_tv12 = _1136_v8
			_ = _pat_let_tv12
			var _pat_let_tv13 = _1136_v8
			_ = _pat_let_tv13
			var _pat_let_tv14 = _1129_v1
			_ = _pat_let_tv14
			for func(_source14 D7) bool {
				if _source14.Is_DC21() {
					var _1221___mcc_h5 D1 = _source14.Get_().(D7_DC21).Cf32
					_ = _1221___mcc_h5
					var _1222___mcc_h6 D6 = _source14.Get_().(D7_DC21).Cf33
					_ = _1222___mcc_h6
					var _1223___mcc_h7 _dafny.Array = _source14.Get_().(D7_DC21).Cf34
					_ = _1223___mcc_h7
					var _1224_cf34 _dafny.Array = _1223___mcc_h7
					_ = _1224_cf34
					var _1225_cf33 D6 = _1222___mcc_h6
					_ = _1225_cf33
					var _1226_cf32 D1 = _1221___mcc_h5
					_ = _1226_cf32
					return (func() _dafny.Set {
						var _coll32 = _dafny.NewBuilder()
						_ = _coll32
						for _iter34 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv9, _pat_let_tv10)).Keys().Elements()); ; {
							_compr_32, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _1227_v42 _dafny.CodePoint
							_1227_v42 = interface{}(_compr_32).(_dafny.CodePoint)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv11, _pat_let_tv12)).Contains(_1227_v42) {
								_coll32.Add(_1227_v42)
							}
						}
						return _coll32.ToSet()
					}()).IsProperSubsetOf(_dafny.SetOf(_pat_let_tv13))
				} else if _source14.Is_DC20() {
					var _1228___mcc_h8 _dafny.Map = _source14.Get_().(D7_DC20).Cf31
					_ = _1228___mcc_h8
					var _1229_cf31 _dafny.Map = _1228___mcc_h8
					_ = _1229_cf31
					return !(!(!(true)))
				} else {
					var _1230___mcc_h9 D7 = _source14.Get_().(D7_DC22).Cf35
					_ = _1230___mcc_h9
					var _1231_cf35 D7 = _1230___mcc_h9
					_ = _1231_cf35
					return (_pat_let_tv14) == (true)
				}
			}(_1185_v41) {
				{
					if (_1187_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1187_i2 = (_1187_i2).Plus(_dafny.One)
					var _1188_v43 _dafny.MultiSet
					_ = _1188_v43
					_1188_v43 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1179_v34).Cardinality()))
					var _1189_v44 bool
					_ = _1189_v44
					var _1190_v45 bool
					_ = _1190_v45
					var _1191_v46 _dafny.Int
					_ = _1191_v46
					var _out28 bool
					_ = _out28
					var _out29 bool
					_ = _out29
					var _out30 _dafny.Int
					_ = _out30
					_out28, _out29, _out30 = (_this).M5(Companion_Default___.Fm2((_1188_v43).Cardinality(), (_this).F24(), _1136_v8, globalState), globalState)
					_1189_v44 = _out28
					_1190_v45 = _out29
					_1191_v46 = _out30
					var _1192_v47 _dafny.Map
					_ = _1192_v47
					_1192_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1190_v45, _1129_v1)
					var _1193_v48 *C2
					_ = _1193_v48
					var _nw169 *C2 = New_C2_()
					_ = _nw169
					_nw169.Ctor__(_1189_v44, (_1192_v47).Cardinality(), (_this).F25())
					_1193_v48 = _nw169
					var _1194_v49 _dafny.Sequence
					_ = _1194_v49
					_1194_v49 = _dafny.SeqOf(_1193_v48)
					var _1195_v50 _dafny.Sequence
					_ = _1195_v50
					_1195_v50 = _dafny.SeqOf((_this).F24(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1194_v49).Cardinality())))
					_1128_v0 = (_1128_v0).Update((_dafny.IntOfUint32((_1195_v50).Cardinality())).Minus(_1191_v46), _1193_v48.F31)
					if _1190_v45 {
						(globalState).F18 = (_1191_v46).Plus((_this).F24())
						var _1196_v51 _dafny.Sequence
						_ = _1196_v51
						_1196_v51 = _dafny.SeqOf(_1192_v47)
						var _1197_v52 _dafny.Sequence
						_ = _1197_v52
						_1197_v52 = _dafny.SeqOf(_1196_v51, _1196_v51)
						var _1198_v53 _dafny.Map
						_ = _1198_v53
						_1198_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1129_v1, _1196_v51)
						var _1199_v54 _dafny.Array
						_ = _1199_v54
						var _nwElement0_34 _dafny.Sequence = (func() _dafny.Sequence {
							if _1190_v45 {
								return _1196_v51
							}
							return _1196_v51
						})()
						_ = _nwElement0_34
						var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(29))
						_ = _nw170
						(_nw170).ArraySet1(_nwElement0_34, 0)
						(_nw170).ArraySet1(_1196_v51, 1)
						(_nw170).ArraySet1(_1196_v51, 2)
						(_nw170).ArraySet1(_1196_v51, 3)
						(_nw170).ArraySet1(_1196_v51, 4)
						(_nw170).ArraySet1(_1196_v51, 5)
						(_nw170).ArraySet1(_1196_v51, 6)
						(_nw170).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm36(_1133_v5, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg72 _dafny.Int) interface{} {
								return coer72(arg72)
							}
						}((func(_1200_v47 _dafny.Map) func(_dafny.Int) _dafny.Map {
							return func(_1201_i3 _dafny.Int) _dafny.Map {
								return _1200_v47
							}
						})(_1192_v47)))), 7)
						(_nw170).ArraySet1(_1196_v51, 8)
						(_nw170).ArraySet1(_1196_v51, 9)
						(_nw170).ArraySet1(_1196_v51, 10)
						(_nw170).ArraySet1((_1197_v52).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F24()), _dafny.IntOfUint32((_1197_v52).Cardinality()))).Uint32()).(_dafny.Sequence), 11)
						(_nw170).ArraySet1(_1196_v51, 12)
						(_nw170).ArraySet1(_1196_v51, 13)
						(_nw170).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1196_v51, _dafny.SeqOf(_1192_v47)), 14)
						(_nw170).ArraySet1((func() _dafny.Sequence {
							if (_1198_v53).Contains(_1193_v48.F31) {
								return (_1198_v53).Get(_1193_v48.F31).(_dafny.Sequence)
							}
							return _dafny.SeqOf(_1192_v47)
						})(), 15)
						(_nw170).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg73 _dafny.Int) interface{} {
								return coer73(arg73)
							}
						}((func(_1202_v47 _dafny.Map, _1203_v1 bool, _1204_v48 *C2) func(_dafny.Int) _dafny.Map {
							return func(_1205_i4 _dafny.Int) _dafny.Map {
								return (_1202_v47).Update(_1203_v1, _1204_v48.F31)
							}
						})(_1192_v47, _1129_v1, _1193_v48))), 16)
						(_nw170).ArraySet1(_1196_v51, 17)
						(_nw170).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1196_v51, _1196_v51), 18)
						(_nw170).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1196_v51, Companion_Default___.Fm36(_dafny.SeqOf(_1129_v1, _1129_v1), globalState)), 19)
						(_nw170).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1196_v51, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1196_v51, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.IntOfUint32((_1196_v51).Cardinality()))).Uint32(), _1192_v47), (Companion_Default___.SafeIndex(_1191_v46, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1196_v51, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.IntOfUint32((_1196_v51).Cardinality()))).Uint32(), _1192_v47)).Cardinality()))).Uint32(), _1192_v47)), 20)
						(_nw170).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(5))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg74 _dafny.Int) interface{} {
								return coer74(arg74)
							}
						}((func(_1206_v47 _dafny.Map) func(_dafny.Int) _dafny.Map {
							return func(_1207_i5 _dafny.Int) _dafny.Map {
								return _1206_v47
							}
						})(_1192_v47))), 21)
						(_nw170).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1196_v51, _1196_v51), 22)
						(_nw170).ArraySet1(_1196_v51, 23)
						(_nw170).ArraySet1(_1196_v51, 24)
						(_nw170).ArraySet1(_1196_v51, 25)
						(_nw170).ArraySet1(_dafny.SeqOf(_1192_v47, _1192_v47, _1192_v47), 26)
						(_nw170).ArraySet1(_dafny.SeqOf(_1192_v47), 27)
						(_nw170).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1196_v51, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(184))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg75 _dafny.Int) interface{} {
								return coer75(arg75)
							}
						}((func(_1208_v47 _dafny.Map, _1209_v48 *C2) func(_dafny.Int) _dafny.Map {
							return func(_1210_i6 _dafny.Int) _dafny.Map {
								return (_1208_v47).Update(_1209_v48.F31, _1209_v48.F31)
							}
						})(_1192_v47, _1193_v48)))), 28)
						_1199_v54 = _nw170
						var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1199_v54), 0))
						_ = _index147
						(_1199_v54).ArraySet1(_1196_v51, (_index147).Int())
						var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1199_v54), 0))
						_ = _index148
						(_1199_v54).ArraySet1(Companion_Default___.Fm36(Companion_Default___.Fm23(globalState), globalState), (_index148).Int())
						(globalState).F14 = (_this).F24()
						var _1211_v55 D1
						_ = _1211_v55
						_1211_v55 = Companion_D1_.Create_DC3_(_1179_v34)
						var _1212_v56 D8
						_ = _1212_v56
						_1212_v56 = Companion_D8_.Create_DC26_(_1211_v55, _1179_v34)
						var _1213_v57 _dafny.Map
						_ = _1213_v57
						_1213_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8(globalState), _1212_v56)
						_1213_v57 = (_1213_v57).Update(_1191_v46, Companion_Default___.Fm37(globalState))
						var _1214_v58 *C2
						_ = _1214_v58
						var _nw171 *C2 = New_C2_()
						_ = _nw171
						_nw171.Ctor__((_1191_v46).Cmp(_1191_v46) < 0, (_this).Fm8(globalState), (_this).F25())
						_1214_v58 = _nw171
					} else {
						var _1215_v59 _dafny.Sequence
						_ = _1215_v59
						_1215_v59 = _dafny.SeqOf(_1188_v43)
						var _1216_v60 *C1
						_ = _1216_v60
						var _nw172 *C1 = New_C1_()
						_ = _nw172
						_nw172.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(694))).Uint32(), func(coer76 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
							return func(arg76 _dafny.Int) interface{} {
								return coer76(arg76)
							}
						}((func(_1217_v43 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
							return func(_1218_i7 _dafny.Int) _dafny.MultiSet {
								return _1217_v43
							}
						})(_1188_v43))), _1215_v59), (_this).F24(), ((_this).F24()).Times((_this).F24()), (_this).F25())
						_1216_v60 = _nw172
						var _1219_v61 *C1
						_ = _1219_v61
						var _nw173 *C1 = New_C1_()
						_ = _nw173
						_nw173.Ctor__(_1215_v59, _1216_v60.F33, _1216_v60.F33, (_this).F25())
						_1219_v61 = _nw173
						_1129_v1 = (func() bool {
							if (_1128_v0).Contains(_1191_v46) {
								return (_1128_v0).Get(_1191_v46).(bool)
							}
							return _1193_v48.F31
						})()
						(_1216_v60).F33 = Companion_Default___.SafeModuloInt(_1219_v61.F33, (_this).F24())
						var _1220_v62 _dafny.MultiSet
						_ = _1220_v62
						_1220_v62 = _dafny.MultiSetOf(_1193_v48.F31, _1129_v1, (_1216_v60.F33).Cmp((_this).F24()) < 0, _1193_v48.F31, _1193_v48.F31)
						_1220_v62 = _1220_v62
					}
					(globalState).F18 = (_1195_v50).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1195_v50).Cardinality()))).Uint32()).(_dafny.Int)
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
	}
}
func (_this *C3) F29() _dafny.Array {
	{
		return _this._f29
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f24 _dafny.Int
	_f25 _dafny.Sequence
	_f28 _dafny.Array
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f24 = _dafny.Zero
	_this._f25 = _dafny.EmptySeq
	_this._f28 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C4) F24() _dafny.Int {
	return _this._f24
}
func (_this *C4) F25() _dafny.Sequence {
	return _this._f25
}
func (_this *C4) Ctor__(f28 _dafny.Array, f24 _dafny.Int, f25 _dafny.Sequence) {
	{
		(_this)._f28 = f28
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C4) Fm5(p0 _dafny.Int, globalState *GlobalState) D0 {
	{
		return Companion_D0_.Create_DC1_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), (_this).F24())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), (_this).F24()))).Cardinality(), (_this).F24(), true)
	}
}
func (_this *C4) M1(globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1232_v1 _dafny.MultiSet
		_ = _1232_v1
		_1232_v1 = _dafny.MultiSetOf((_this).F24())
		var _1233_v2 bool
		_ = _1233_v2
		_1233_v2 = false
		var _1234_v3 _dafny.Sequence
		_ = _1234_v3
		_1234_v3 = _dafny.SeqOf(_1233_v2, _1233_v2)
		var _1235_v4 _dafny.Map
		_ = _1235_v4
		_1235_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F24())
		var _1236_v5 _dafny.Sequence
		_ = _1236_v5
		_1236_v5 = _dafny.SeqOf(_1235_v4)
		var _1237_v6 _dafny.CodePoint
		_ = _1237_v6
		_1237_v6 = _dafny.CodePoint('s')
		var _1238_v7 D0
		_ = _1238_v7
		_1238_v7 = Companion_D0_.Create_DC2_(_1233_v2, ((_1236_v5).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1236_v5).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _1237_v6)
		var _1239_v8 _dafny.Sequence
		_ = _1239_v8
		_1239_v8 = _dafny.UnicodeSeqOfUtf8Bytes("dnox")
		var _1240_v9 _dafny.Map
		_ = _1240_v9
		_1240_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.IntOfUint32((_1239_v8).Cardinality()))
		var _1241_v10 _dafny.Map
		_ = _1241_v10
		_1241_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1240_v9)
		var _1242_v12 _dafny.Array
		_ = _1242_v12
		var _nw174 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw174
		_1242_v12 = _nw174
		var _1243_v13 _dafny.Map
		_ = _1243_v13
		_1243_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1233_v2, _1242_v12)
		var _1244_v14 _dafny.Array
		_ = _1244_v14
		var _nwElement0_35 _dafny.Map = func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter35 := _dafny.Iterate((_1232_v1).Elements()); ; {
				_compr_33, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _1245_v0 _dafny.Int
				_1245_v0 = interface{}(_compr_33).(_dafny.Int)
				if (_1232_v1).Contains(_1245_v0) {
					_coll33.Add(Companion_Default___.SafeModuloInt(_1245_v0, (_this).F24()), _dafny.IntOfInt64(98))
				}
			}
			return _coll33.ToMap()
		}()
		_ = _nwElement0_35
		var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(13))
		_ = _nw175
		(_nw175).ArraySet1(_nwElement0_35, 0)
		(_nw175).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1233_v2, _1233_v2)).Cardinality(), (_this).F24()), 1)
		(_nw175).ArraySet1((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_dafny.Zero).Minus((_this).F24()))).Update((_this).F24(), (_this).F24())).Update((_this).F24(), _dafny.IntOfInt64(698))).Update(_dafny.IntOfUint32((_1234_v3).Cardinality()), (_1238_v7).Dtor_cf5()), 2)
		(_nw175).ArraySet1(Companion_Default___.Fm6((_this).F24(), (_this).F24(), Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState), (_this).F24(), globalState), 3)
		(_nw175).ArraySet1(_1240_v9, 4)
		(_nw175).ArraySet1(_1240_v9, 5)
		(_nw175).ArraySet1(_1240_v9, 6)
		(_nw175).ArraySet1((func() _dafny.Map {
			if (_1241_v10).Contains((_this).F24()) {
				return (_1241_v10).Get((_this).F24()).(_dafny.Map)
			}
			return _1240_v9
		})(), 7)
		(_nw175).ArraySet1(_1240_v9, 8)
		(_nw175).ArraySet1(_1240_v9, 9)
		(_nw175).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())).Merge(func() _dafny.Map {
			var _coll34 = _dafny.NewMapBuilder()
			_ = _coll34
			for _iter36 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm7(globalState), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((Companion_Default___.Fm7(globalState)).Cardinality()))).Uint32(), (_this).F24())).Elements()); ; {
				_compr_34, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _1246_v11 _dafny.Int
				_1246_v11 = interface{}(_compr_34).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm7(globalState), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((Companion_Default___.Fm7(globalState)).Cardinality()))).Uint32(), (_this).F24()), _1246_v11) {
					_coll34.Add((_1246_v11).Minus((_this).F24()), (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1239_v8).Cardinality()), _1233_v2))).Cardinality())
				}
			}
			return _coll34.ToMap()
		}()), 10)
		(_nw175).ArraySet1((_1240_v9).Merge(_1240_v9), 11)
		(_nw175).ArraySet1((_1240_v9).Update((_1243_v13).Cardinality(), (_dafny.Zero).Minus((_this).F24())), 12)
		_1244_v14 = _nw175
		for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1244_v14), 0))); ; {
			_guard_loop_2, _ok37 := _iter37()
			if !_ok37 {
				break
			}
			var _1247_i0 _dafny.Int
			_1247_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1247_i0).Sign() != -1) && ((_1247_i0).Cmp(_dafny.ArrayLen((_1244_v14), 0)) < 0)) {
				(_1244_v14).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24()), (_1247_i0).Int())
			}
		}
		var _1248_v15 _dafny.Set
		_ = _1248_v15
		_1248_v15 = _dafny.SetOf(_dafny.IntOfInt64(575))
		var _hi7 _dafny.Int = ((_this).F24()).Plus((_1248_v15).Cardinality())
		_ = _hi7
		for _1249_i1 := (_this).F24(); _1249_i1.Cmp(_hi7) < 0; _1249_i1 = _1249_i1.Plus(_dafny.One) {
			var _1250_v16 _dafny.Array
			_ = _1250_v16
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_29
			var _nw176 _dafny.Array
			_ = _nw176
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw176 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.Sequence = (func(_1251_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1252_i2 _dafny.Int) _dafny.Sequence {
						return _1251_v8
					}
				})(_1239_v8)
				_ = _init29
				var _element0_29 = _init29(_dafny.Zero)
				_ = _element0_29
				_nw176 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
				(_nw176).ArraySet1(_element0_29, 0)
				var _nativeLen0_29 = (_len0_29).Int()
				_ = _nativeLen0_29
				for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
					(_nw176).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
				}
			}
			_1250_v16 = _nw176
			_1250_v16 = _1250_v16
			(globalState).F14 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1239_v8, (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1239_v8).Cardinality()))).Uint32(), _1237_v6)).Cardinality())).Minus(Companion_Default___.SafeModuloInt((_this).F24(), (_this).F24()))
			var _1253_v17 _dafny.Sequence
			_ = _1253_v17
			_1253_v17 = _dafny.SeqOf(_dafny.IntOfUint32((_1239_v8).Cardinality()), _1249_i1, _1249_i1)
			var _1254_v18 D1
			_ = _1254_v18
			_1254_v18 = Companion_D1_.Create_DC4_(((_this).F24()).Cmp(_1249_i1) <= 0, _1233_v2, _1237_v6, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm2(_1249_i1, _dafny.IntOfUint32((_1239_v8).Cardinality()), _1237_v6, globalState)), _1234_v3)).Cardinality()), (_1253_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1239_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1253_v17).Cardinality()), _dafny.IntOfUint32((_1239_v8).Cardinality()))).Uint32(), _1237_v6)).Cardinality()), _dafny.IntOfUint32((_1253_v17).Cardinality()))).Uint32()).(_dafny.Int))
			_1254_v18 = _1254_v18
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index149
			((_this).F28()).ArraySet1((_dafny.IntOfInt64(350)).Minus(_1249_i1), (_index149).Int())
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index150
			var _rhs199 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1239_v8, _1239_v8), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F24()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1239_v8, _1239_v8)).Cardinality()))).Uint32(), _1237_v6)).Cardinality()))
			_ = _rhs199
			var _rhs200 bool = _1233_v2
			_ = _rhs200
			var _lhs155 _dafny.Array = (_this).F28()
			_ = _lhs155
			var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _lhs156
			var _lhs157 *GlobalState = globalState
			_ = _lhs157
			(_lhs155).ArraySet1(_rhs199, (_lhs156).Int())
			_lhs157.F2 = _rhs200
		}
		var _1255_v19 _dafny.MultiSet
		_ = _1255_v19
		_1255_v19 = _dafny.MultiSetOf(_1237_v6, _1237_v6)
		var _1256_v21 _dafny.Set
		_ = _1256_v21
		_1256_v21 = _dafny.SetOf(_1237_v6, _1237_v6, _1237_v6)
		var _hi8 _dafny.Int = (_this).F24()
		_ = _hi8
		for _1257_i3 := ((func() _dafny.Set {
			var _coll35 = _dafny.NewBuilder()
			_ = _coll35
			for _iter38 := _dafny.Iterate((_1255_v19).Elements()); ; {
				_compr_35, _ok38 := _iter38()
				if !_ok38 {
					break
				}
				var _1280_v20 _dafny.CodePoint
				_1280_v20 = interface{}(_compr_35).(_dafny.CodePoint)
				if (_1255_v19).Contains(_1280_v20) {
					_coll35.Add(_1280_v20)
				}
			}
			return _coll35.ToSet()
		}()).Intersection(_1256_v21)).Cardinality(); _1257_i3.Cmp(_hi8) < 0; _1257_i3 = _1257_i3.Plus(_dafny.One) {
			if _1233_v2 {
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1242_v12), 0))
				_ = _index151
				(_1242_v12).ArraySet1((func(_pat_let16_0 D0) D0 {
					return func(_1258_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let17_0 _dafny.Int) D0 {
							return func(_1259_dt__update_hcf5_h0 _dafny.Int) D0 {
								return Companion_D0_.Create_DC2_((_1258_dt__update__tmp_h0).Dtor_cf4(), _1259_dt__update_hcf5_h0, (_1258_dt__update__tmp_h0).Dtor_cf6())
							}(_pat_let17_0)
						}(_1257_i3)
					}(_pat_let16_0)
				}(_1238_v7)).Dtor_cf4(), (_index151).Int())
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1242_v12), 0))
				_ = _index152
				(_1242_v12).ArraySet1(Companion_Default___.Fm2((_dafny.Zero).Minus((_1257_i3).Plus((_1240_v9).Cardinality())), _1257_i3, _1237_v6, globalState), (_index152).Int())
				var _1260_v22 _dafny.Sequence
				_ = _1260_v22
				_1260_v22 = _dafny.SeqOf(_dafny.IntOfUint32((_1239_v8).Cardinality()))
				var _1261_v23 _dafny.Sequence
				_ = _1261_v23
				_1261_v23 = _dafny.SeqOf(_1232_v1, _dafny.MultiSetOf(_1257_i3))
				var _1262_v24 _dafny.Map
				_ = _1262_v24
				_1262_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v6, _1257_i3)
				var _1263_v25 T0
				_ = _1263_v25
				var _nw177 *C1 = New_C1_()
				_ = _nw177
				_nw177.Ctor__(_1261_v23, (_this).F24(), (_this).F24(), _dafny.SeqOf(_1262_v24, _1262_v24, _1262_v24))
				_1263_v25 = _nw177
				var _1264_v26 _dafny.Map
				_ = _1264_v26
				_1264_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1263_v25, _1233_v2)
				(globalState).F1 = Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32((_1260_v22).Cardinality())).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_1260_v22).Cardinality()))), (_1264_v26).Cardinality())
				(globalState).F18 = _1257_i3
				var _1265_v27 _dafny.MultiSet
				_ = _1265_v27
				_1265_v27 = _dafny.MultiSetOf(_1240_v9)
				(globalState).F2 = (_1265_v27).IsProperSubsetOf(_1265_v27)
				var _rhs201 _dafny.Set = _dafny.SetOf((_dafny.Zero).Minus(_1257_i3))
				_ = _rhs201
				var _rhs202 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1239_v8, _dafny.UnicodeSeqOfUtf8Bytes("d")), _1239_v8)
				_ = _rhs202
				var _rhs203 bool = _1233_v2
				_ = _rhs203
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				_1248_v15 = _rhs201
				_lhs158.F17 = _rhs202
				r0 = _rhs203
			} else {
				var _1266_v28 _dafny.Sequence
				_ = _1266_v28
				_1266_v28 = _dafny.SeqOf(_1257_i3)
				var _1267_v29 D4
				_ = _1267_v29
				_1267_v29 = Companion_D4_.Create_DC12_(_1266_v28)
				var _1268_v30 D4
				_ = _1268_v30
				_1268_v30 = Companion_D4_.Create_DC14_(_1267_v29)
				var _1269_v31 _dafny.Map
				_ = _1269_v31
				_1269_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1268_v30, _1239_v8)
				var _1270_v32 D7
				_ = _1270_v32
				_1270_v32 = Companion_D7_.Create_DC20_(_1269_v31)
				var _1271_v33 _dafny.MultiSet
				_ = _1271_v33
				_1271_v33 = _dafny.MultiSetOf(_1270_v32, Companion_D7_.Create_DC20_(_1269_v31), _1270_v32)
				_1271_v33 = _1271_v33
				(globalState).F14 = ((_this).F24()).Times(_1257_i3)
				var _1272_v34 _dafny.Map
				_ = _1272_v34
				_1272_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1233_v2, _1233_v2)
				_1272_v34 = (_1272_v34).Update(!(_1233_v2), _1233_v2)
				(globalState).F2 = _1233_v2
				var _1273_v35 _dafny.MultiSet
				_ = _1273_v35
				_1273_v35 = _dafny.MultiSetOf(_1233_v2)
				var _1274_v36 D3
				_ = _1274_v36
				_1274_v36 = Companion_D3_.Create_DC10_()
				var _1275_v37 _dafny.MultiSet
				_ = _1275_v37
				_1275_v37 = _dafny.MultiSetOf(_1274_v36)
				_1273_v35 = _dafny.MultiSetOf((_1275_v37).IsProperSubsetOf(_1275_v37))
			}
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1242_v12), 0))
			_ = _index153
			(_1242_v12).ArraySet1(false, (_index153).Int())
			var _1276_v38 _dafny.Sequence
			_ = _1276_v38
			_1276_v38 = _dafny.SeqOf(_1239_v8)
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1242_v12), 0))
			_ = _index154
			var _rhs204 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xho"), _1239_v8), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg77 _dafny.Int) interface{} {
					return coer77(arg77)
				}
			}((func(_1277_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1278_i4 _dafny.Int) _dafny.CodePoint {
					return _1277_v6
				}
			})(_1237_v6))))
			_ = _rhs204
			var _rhs205 _dafny.MultiSet = _1255_v19
			_ = _rhs205
			var _rhs206 _dafny.Sequence = _dafny.SeqOf(_1239_v8, _1239_v8)
			_ = _rhs206
			var _rhs207 bool = _1233_v2
			_ = _rhs207
			var _lhs159 _dafny.Array = _1242_v12
			_ = _lhs159
			var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1242_v12), 0))
			_ = _lhs160
			var _lhs161 *GlobalState = globalState
			_ = _lhs161
			(_lhs159).ArraySet1(_rhs204, (_lhs160).Int())
			_1255_v19 = _rhs205
			_1276_v38 = _rhs206
			_lhs161.F2 = _rhs207
			_1235_v4 = (_1235_v4).Update(!((_1242_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1242_v12), 0))).Int()).(bool)), _1257_i3)
			var _1279_v39 _dafny.Sequence
			_ = _1279_v39
			_1279_v39 = _dafny.SeqOf((_this).F28(), (_this).F28(), (_this).F28())
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1242_v12), 0))
			_ = _index155
			var _rhs208 _dafny.Sequence = _1239_v8
			_ = _rhs208
			var _rhs209 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1279_v39, (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1279_v39).Cardinality()))).Uint32(), (_this).F28())
			_ = _rhs209
			var _rhs210 _dafny.Set = _1248_v15
			_ = _rhs210
			var _rhs211 bool = _1233_v2
			_ = _rhs211
			var _rhs212 bool = !((_1242_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1242_v12), 0))).Int()).(bool)) || (_1233_v2)
			_ = _rhs212
			var _lhs162 *GlobalState = globalState
			_ = _lhs162
			var _lhs163 _dafny.Array = _1242_v12
			_ = _lhs163
			var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.ArrayLen((_1242_v12), 0))
			_ = _lhs164
			var _lhs165 *GlobalState = globalState
			_ = _lhs165
			_lhs162.F17 = _rhs208
			_1279_v39 = _rhs209
			_1248_v15 = _rhs210
			(_lhs163).ArraySet1(_rhs211, (_lhs164).Int())
			_lhs165.F2 = _rhs212
		}
		var _1281_i5 _dafny.Int
		_ = _1281_i5
		_1281_i5 = _dafny.Zero
		{
			for ((_this).F24()).Cmp((_this).F24()) == 0 {
				{
					if (_1281_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1281_i5 = (_1281_i5).Plus(_dafny.One)
					var _pat_let_tv15 = _1233_v2
					_ = _pat_let_tv15
					var _1282_v40 _dafny.MultiSet
					_ = _1282_v40
					_1282_v40 = _dafny.MultiSetOf(func(_pat_let18_0 D3) D3 {
						return func(_1283_dt__update__tmp_h1 D3) D3 {
							return func(_pat_let19_0 bool) D3 {
								return func(_1284_dt__update_hcf20_h0 bool) D3 {
									return Companion_D3_.Create_DC11_((_1283_dt__update__tmp_h1).Dtor_cf18(), (_1283_dt__update__tmp_h1).Dtor_cf19(), _1284_dt__update_hcf20_h0)
								}(_pat_let19_0)
							}(_pat_let_tv15)
						}(_pat_let18_0)
					}(Companion_D3_.Create_DC11_(Companion_Default___.Fm0((_this).F24(), _dafny.IntOfInt64(-151), globalState), false, _1233_v2)), Companion_D3_.Create_DC11_((_this).F24(), _1233_v2, _1233_v2))
					if !((_1282_v40).IsSubsetOf(Companion_Default___.Fm38(globalState))) {
						var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index156
						((_this).F28()).ArraySet1((_1248_v15).Cardinality(), (_index156).Int())
						var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index157
						var _rhs213 _dafny.Int = (((_this).F24()).Plus((_this).F24())).Plus(Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState))
						_ = _rhs213
						var _rhs214 _dafny.Int = (_this).F24()
						_ = _rhs214
						var _lhs166 _dafny.Array = (_this).F28()
						_ = _lhs166
						var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _lhs167
						var _lhs168 *GlobalState = globalState
						_ = _lhs168
						(_lhs166).ArraySet1(_rhs213, (_lhs167).Int())
						_lhs168.F14 = _rhs214
						(globalState).F17 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(244))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg78 _dafny.Int) interface{} {
								return coer78(arg78)
							}
						}((func(_1285_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1286_i6 _dafny.Int) _dafny.CodePoint {
								return _1285_v6
							}
						})(_1237_v6))), _1239_v8)
						r0 = _1233_v2
						var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index158
						var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index159
						var _rhs215 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("f")
						_ = _rhs215
						var _rhs216 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(337), (_dafny.Zero).Minus((_this).F24())))
						_ = _rhs216
						var _rhs217 _dafny.Int = (_1255_v19).Cardinality()
						_ = _rhs217
						var _lhs169 *GlobalState = globalState
						_ = _lhs169
						var _lhs170 _dafny.Array = (_this).F28()
						_ = _lhs170
						var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _lhs171
						var _lhs172 _dafny.Array = (_this).F28()
						_ = _lhs172
						var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _lhs173
						_lhs169.F17 = _rhs215
						(_lhs170).ArraySet1(_rhs216, (_lhs171).Int())
						(_lhs172).ArraySet1(_rhs217, (_lhs173).Int())
						var _1287_v41 _dafny.Array
						_ = _1287_v41
						var _nw178 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
						_ = _nw178
						_1287_v41 = _nw178
						var _1288_v42 T0
						_ = _1288_v42
						var _nw179 *C2 = New_C2_()
						_ = _nw179
						_nw179.Ctor__(_1233_v2, ((_this).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))).Int()).(_dafny.Int), (_this).F25())
						_1288_v42 = _nw179
						var _1289_v43 _dafny.Map
						_ = _1289_v43
						_1289_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1234_v3).Select((Companion_Default___.SafeIndex(((_this).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen(((_this).F28()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1234_v3).Cardinality()))).Uint32()).(bool), _1288_v42)
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1287_v41), 0))
						_ = _index160
						(_1287_v41).ArraySet1((func() T0 {
							if (_1289_v43).Contains(_1233_v2) {
								return (_1289_v43).Get(_1233_v2).(T0)
							}
							return _1288_v42
						})(), (_index160).Int())
						var _1290_v44 _dafny.Map
						_ = _1290_v44
						_1290_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1240_v9, _1288_v42)
						var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1287_v41), 0))
						_ = _index161
						(_1287_v41).ArraySet1((func() T0 {
							if (_1290_v44).Contains((_1240_v9).Merge(_1240_v9)) {
								return (_1290_v44).Get((_1240_v9).Merge(_1240_v9)).(T0)
							}
							return _1288_v42
						})(), (_index161).Int())
					} else {
						var _1291_v45 _dafny.Array
						_ = _1291_v45
						var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
						_ = _nw180
						_1291_v45 = _nw180
						var _1292_v46 _dafny.Map
						_ = _1292_v46
						_1292_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1233_v2)
						var _1293_v47 D1
						_ = _1293_v47
						_1293_v47 = Companion_D1_.Create_DC3_(Companion_Default___.Fm4(false, (_this).F24(), _1237_v6, _1233_v2, globalState))
						var _1294_v48 D8
						_ = _1294_v48
						_1294_v48 = Companion_D8_.Create_DC26_(_1293_v47, _dafny.UnicodeSeqOfUtf8Bytes("b"))
						var _1295_v49 _dafny.Array
						_ = _1295_v49
						var _nwElement0_36 _dafny.Sequence = _1239_v8
						_ = _nwElement0_36
						var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(19))
						_ = _nw181
						(_nw181).ArraySet1(_nwElement0_36, 0)
						(_nw181).ArraySet1(_1239_v8, 1)
						(_nw181).ArraySet1(_1239_v8, 2)
						(_nw181).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("an"), 3)
						(_nw181).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg79 _dafny.Int) interface{} {
								return coer79(arg79)
							}
						}((func(_1296_v2 bool, _1297_v6 _dafny.CodePoint, _1298_v8 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
							return func(_1299_i7 _dafny.Int) _dafny.CodePoint {
								return (Companion_D1_.Create_DC4_(_1296_v2, _1296_v2, _1297_v6, (_this).F24(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1298_v8).Cardinality()))).Cardinality()))).Dtor_cf10()
							}
						})(_1233_v2, _1237_v6, _1239_v8))), 4)
						(_nw181).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(187))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg80 _dafny.Int) interface{} {
								return coer80(arg80)
							}
						}((func(_1300_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1301_i8 _dafny.Int) _dafny.CodePoint {
								return _1300_v6
							}
						})(_1237_v6))), 5)
						(_nw181).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-797))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg81 _dafny.Int) interface{} {
								return coer81(arg81)
							}
						}((func(_1302_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1303_i9 _dafny.Int) _dafny.CodePoint {
								return _1302_v6
							}
						})(_1237_v6))), (Companion_Default___.SafeIndex((_1292_v46).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-797))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg82 _dafny.Int) interface{} {
								return coer82(arg82)
							}
						}((func(_1304_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1305_i9 _dafny.Int) _dafny.CodePoint {
								return _1304_v6
							}
						})(_1237_v6)))).Cardinality()))).Uint32(), _1237_v6), 6)
						(_nw181).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kncwpm"), 7)
						(_nw181).ArraySet1(_1239_v8, 8)
						(_nw181).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mpjutg"), 9)
						(_nw181).ArraySet1(_1239_v8, 10)
						(_nw181).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("wx"), 11)
						(_nw181).ArraySet1(_1239_v8, 12)
						(_nw181).ArraySet1(_1239_v8, 13)
						(_nw181).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(911))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg83 _dafny.Int) interface{} {
								return coer83(arg83)
							}
						}((func(_1306_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1307_i10 _dafny.Int) _dafny.CodePoint {
								return _1306_v6
							}
						})(_1237_v6))), 14)
						(_nw181).ArraySet1(_1239_v8, 15)
						(_nw181).ArraySet1(_1239_v8, 16)
						(_nw181).ArraySet1(_1239_v8, 17)
						(_nw181).ArraySet1((_1294_v48).Dtor_cf42(), 18)
						_1295_v49 = _nw181
						var _1308_v50 D10
						_ = _1308_v50
						_1308_v50 = Companion_D10_.Create_DC28_(_1295_v49)
						var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1291_v45), 0))
						_ = _index162
						(_1291_v45).ArraySet1((_1308_v50).Dtor_cf44(), (_index162).Int())
						var _1309_v51 _dafny.Map
						_ = _1309_v51
						_1309_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1233_v2, _1295_v49)
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_1291_v45), 0))
						_ = _index163
						(_1291_v45).ArraySet1((func() _dafny.Array {
							if (_1309_v51).Contains(((_this).F24()).Cmp((_dafny.Zero).Minus((_this).F24())) < 0) {
								return (_1309_v51).Get(((_this).F24()).Cmp((_dafny.Zero).Minus((_this).F24())) < 0).(_dafny.Array)
							}
							return _1295_v49
						})(), (_index163).Int())
						_1235_v4 = _1235_v4
						var _1310_v52 _dafny.Array
						_ = _1310_v52
						var _nw182 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(26))
						_ = _nw182
						_1310_v52 = _nw182
						var _1311_v53 _dafny.Map
						_ = _1311_v53
						_1311_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v6, (_this).F24())
						var _1312_v54 *C3
						_ = _1312_v54
						var _nw183 *C3 = New_C3_()
						_ = _nw183
						_nw183.Ctor__(_1310_v52, (func() _dafny.Array {
							if _1233_v2 {
								return _1244_v14
							}
							return _1244_v14
						})(), Companion_Default___.SafeDivisionInt((Companion_D8_.Create_DC25_((_this).F24(), _1233_v2, _1233_v2)).Dtor_cf38(), (_this).F24()), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F25(), (_this).F25()), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F25(), (_this).F25())).Cardinality()))).Uint32(), _1311_v53))
						_1312_v54 = _nw183
						var _1313_v56 D1
						_ = _1313_v56
						_1313_v56 = Companion_D1_.Create_DC4_(_1233_v2, _1233_v2, _1237_v6, (_this).F24(), (func() _dafny.Map {
							var _coll36 = _dafny.NewMapBuilder()
							_ = _coll36
							for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(784), _dafny.IntOfInt64(805))); ; {
								_compr_36, _ok39 := _iter39()
								if !_ok39 {
									break
								}
								var _1314_v55 _dafny.Int
								_1314_v55 = interface{}(_compr_36).(_dafny.Int)
								if ((_dafny.IntOfInt64(784)).Cmp(_1314_v55) <= 0) && ((_1314_v55).Cmp(_dafny.IntOfInt64(805)) < 0) {
									_coll36.Add((_1314_v55).Plus((_this).F24()), _1237_v6)
								}
							}
							return _coll36.ToMap()
						}()).Cardinality())
						r0 = ((_1313_v56).Dtor_cf12()).Cmp((_this).F24()) >= 0
						(globalState).F2 = !(Companion_Default___.Fm2(((_this).F24()).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(125))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg84 _dafny.Int) interface{} {
								return coer84(arg84)
							}
						}((func(_1315_v2 bool) func(_dafny.Int) _dafny.Int {
							return func(_1316_i11 _dafny.Int) _dafny.Int {
								return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1315_v2)).Cardinality()))
							}
						})(_1233_v2))), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(125))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg85 _dafny.Int) interface{} {
								return coer85(arg85)
							}
						}((func(_1317_v2 bool) func(_dafny.Int) _dafny.Int {
							return func(_1318_i11 _dafny.Int) _dafny.Int {
								return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1317_v2)).Cardinality()))
							}
						})(_1233_v2)))).Cardinality()))).Uint32(), (_this).F24())).Cardinality())), _dafny.IntOfInt64(-970), _1237_v6, globalState))
					}
					(globalState).F2 = false
					var _1319_v57 _dafny.Array
					_ = _1319_v57
					var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
					_ = _nw184
					_1319_v57 = _nw184
					var _1320_v58 _dafny.Map
					_ = _1320_v58
					_1320_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v6, (_this).F24())
					var _1321_v59 *C3
					_ = _1321_v59
					var _nw185 *C3 = New_C3_()
					_ = _nw185
					_nw185.Ctor__(_1319_v57, _1244_v14, Companion_Default___.SafeModuloInt((_this).F24(), (_this).F24()), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v6, (_this).F24()), _1320_v58))
					_1321_v59 = _nw185
					var _1322_v60 _dafny.Map
					_ = _1322_v60
					_1322_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v12, false)
					var _1323_v61 *C2
					_ = _1323_v61
					var _nw186 *C2 = New_C2_()
					_ = _nw186
					_nw186.Ctor__((func() bool {
						if (_1322_v60).Contains(_1242_v12) {
							return (_1322_v60).Get(_1242_v12).(bool)
						}
						return !(!(_1233_v2))
					})(), (_this).F24(), _dafny.Companion_Sequence_.Update((_this).F25(), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32(((_this).F25()).Cardinality()))).Uint32(), _1320_v58))
					_1323_v61 = _nw186
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1324_i12 _dafny.Int
		_ = _1324_i12
		_1324_i12 = _dafny.Zero
		{
			for _1233_v2 {
				{
					if (_1324_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1324_i12 = (_1324_i12).Plus(_dafny.One)
					(globalState).F18 = (_this).F24()
					(globalState).F14 = (_this).F24()
					var _1325_v62 *C2
					_ = _1325_v62
					var _nw187 *C2 = New_C2_()
					_ = _nw187
					_nw187.Ctor__(!(_1233_v2), (_this).F24(), (_this).F25())
					_1325_v62 = _nw187
					if _1233_v2 {
						var _1326_v65 _dafny.Array
						_ = _1326_v65
						var _nwElement0_37 _dafny.Set = _dafny.SetOf((_this).F24())
						_ = _nwElement0_37
						var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(22))
						_ = _nw188
						(_nw188).ArraySet1(_nwElement0_37, 0)
						(_nw188).ArraySet1(_dafny.SetOf((_this).F24(), (_this).F24()), 1)
						(_nw188).ArraySet1(_1248_v15, 2)
						(_nw188).ArraySet1(_1248_v15, 3)
						(_nw188).ArraySet1(_dafny.SetOf((_this).F24(), (_this).F24()), 4)
						(_nw188).ArraySet1(_1248_v15, 5)
						(_nw188).ArraySet1(_1248_v15, 6)
						(_nw188).ArraySet1(_1248_v15, 7)
						(_nw188).ArraySet1(_1248_v15, 8)
						(_nw188).ArraySet1(_1248_v15, 9)
						(_nw188).ArraySet1(func() _dafny.Set {
							var _coll37 = _dafny.NewBuilder()
							_ = _coll37
							for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(678), _dafny.IntOfInt64(758))); ; {
								_compr_37, _ok40 := _iter40()
								if !_ok40 {
									break
								}
								var _1327_v63 _dafny.Int
								_1327_v63 = interface{}(_compr_37).(_dafny.Int)
								if ((_dafny.IntOfInt64(678)).Cmp(_1327_v63) <= 0) && ((_1327_v63).Cmp(_dafny.IntOfInt64(758)) < 0) {
									_coll37.Add((_1327_v63).Times((_this).F24()))
								}
							}
							return _coll37.ToSet()
						}(), 10)
						(_nw188).ArraySet1(func() _dafny.Set {
							var _coll38 = _dafny.NewBuilder()
							_ = _coll38
							for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(494), _dafny.IntOfInt64(-646))); ; {
								_compr_38, _ok41 := _iter41()
								if !_ok41 {
									break
								}
								var _1328_v64 _dafny.Int
								_1328_v64 = interface{}(_compr_38).(_dafny.Int)
								if ((_dafny.IntOfInt64(494)).Cmp(_1328_v64) <= 0) && ((_1328_v64).Cmp(_dafny.IntOfInt64(-646)) < 0) {
									_coll38.Add((_1328_v64).Plus((_this).F24()))
								}
							}
							return _coll38.ToSet()
						}(), 11)
						(_nw188).ArraySet1(_1248_v15, 12)
						(_nw188).ArraySet1(_1248_v15, 13)
						(_nw188).ArraySet1(_1248_v15, 14)
						(_nw188).ArraySet1(_1248_v15, 15)
						(_nw188).ArraySet1(_1248_v15, 16)
						(_nw188).ArraySet1(_1248_v15, 17)
						(_nw188).ArraySet1(_1248_v15, 18)
						(_nw188).ArraySet1(_1248_v15, 19)
						(_nw188).ArraySet1(_dafny.SetOf((_this).F24()), 20)
						(_nw188).ArraySet1(_1248_v15, 21)
						_1326_v65 = _nw188
						var _1329_v66 *C3
						_ = _1329_v66
						var _nw189 *C3 = New_C3_()
						_ = _nw189
						_nw189.Ctor__(_1326_v65, _1244_v14, _dafny.IntOfUint32((_1239_v8).Cardinality()), (_this).F25())
						_1329_v66 = _nw189
						var _1330_v67 _dafny.Sequence
						_ = _1330_v67
						_1330_v67 = _dafny.SeqOf((_this).F24())
						var _1331_v68 _dafny.Sequence
						_ = _1331_v68
						_1331_v68 = _dafny.SeqOf((_1330_v67).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1330_v67).Cardinality()))).Uint32()).(_dafny.Int))
						(globalState).F2 = (func() bool {
							if _1233_v2 {
								return !(_dafny.Companion_Sequence_.IsProperPrefixOf(_1331_v68, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(300))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg86 _dafny.Int) interface{} {
										return coer86(arg86)
									}
								}(func(_1332_i13 _dafny.Int) _dafny.Int {
									return _dafny.IntOfInt64(717)
								}))))
							}
							return _1325_v62.F31
						})()
						(globalState).F17 = _1239_v8
						var _1333_v69 _dafny.Map
						_ = _1333_v69
						_1333_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v12, _1233_v2)
						_1333_v69 = (_1333_v69).Update(_1242_v12, _1325_v62.F31)
						var _1334_v70 _dafny.Map
						_ = _1334_v70
						_1334_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1330_v67, _1325_v62.F31)
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index164
						((_this).F28()).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(100), (_dafny.Zero).Minus((_1334_v70).Cardinality())), (_index164).Int())
						var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(638), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index165
						((_this).F28()).ArraySet1((_this).F24(), (_index165).Int())
					} else {
						var _1335_v71 _dafny.Array
						_ = _1335_v71
						var _len0_30 _dafny.Int = _dafny.IntOfInt64(23)
						_ = _len0_30
						var _nw190 _dafny.Array
						_ = _nw190
						if _len0_30.Cmp(_dafny.Zero) == 0 {
							_nw190 = _dafny.NewArray(_len0_30)
						} else {
							var _init30 func(_dafny.Int) _dafny.Set = (func(_1336_v15 _dafny.Set) func(_dafny.Int) _dafny.Set {
								return func(_1337_i14 _dafny.Int) _dafny.Set {
									return _1336_v15
								}
							})(_1248_v15)
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
						_1335_v71 = _nw190
						var _1338_v73 _dafny.Map
						_ = _1338_v73
						_1338_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v6, _1325_v62.F31)
						var _1339_v74 _dafny.Map
						_ = _1339_v74
						_1339_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v6, (_this).F24())
						var _1340_v75 *C3
						_ = _1340_v75
						var _nw191 *C3 = New_C3_()
						_ = _nw191
						_nw191.Ctor__(_1335_v71, _1244_v14, (_this).F24(), _dafny.SeqOf(func() _dafny.Map {
							var _coll39 = _dafny.NewMapBuilder()
							_ = _coll39
							for _iter42 := _dafny.Iterate((_1338_v73).Keys().Elements()); ; {
								_compr_39, _ok42 := _iter42()
								if !_ok42 {
									break
								}
								var _1341_v72 _dafny.CodePoint
								_1341_v72 = interface{}(_compr_39).(_dafny.CodePoint)
								if (_1338_v73).Contains(_1341_v72) {
									_coll39.Add(_1341_v72, (_dafny.SetOf((Companion_D3_.Create_DC11_((_this).F24(), _1325_v62.F31, true)).Dtor_cf20())).Cardinality())
								}
							}
							return _coll39.ToMap()
						}(), _1339_v74, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v6, (_dafny.Zero).Minus((_this).F24())), _1339_v74))
						_1340_v75 = _nw191
						_1340_v75 = _1340_v75
						var _1342_v76 _dafny.Map
						_ = _1342_v76
						_1342_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1325_v62.F31, _1234_v3)
						var _1343_v77 *C2
						_ = _1343_v77
						var _nw192 *C2 = New_C2_()
						_ = _nw192
						_nw192.Ctor__(!(_1342_v76).Contains(_1233_v2), (_this).F24(), _dafny.SeqOf(_1339_v74))
						_1343_v77 = _nw192
						var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index166
						((_this).F28()).ArraySet1((_this).F24(), (_index166).Int())
						var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index167
						var _rhs218 bool = _1325_v62.F31
						_ = _rhs218
						var _rhs219 bool = false
						_ = _rhs219
						var _rhs220 bool = _1325_v62.F31
						_ = _rhs220
						var _rhs221 _dafny.Int = (_this).F24()
						_ = _rhs221
						var _rhs222 _dafny.Int = (_this).F24()
						_ = _rhs222
						var _lhs174 *GlobalState = globalState
						_ = _lhs174
						var _lhs175 *GlobalState = globalState
						_ = _lhs175
						var _lhs176 *GlobalState = globalState
						_ = _lhs176
						var _lhs177 _dafny.Array = (_this).F28()
						_ = _lhs177
						var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _lhs178
						var _lhs179 *GlobalState = globalState
						_ = _lhs179
						_lhs174.F2 = _rhs218
						_lhs175.F2 = _rhs219
						_lhs176.F2 = _rhs220
						(_lhs177).ArraySet1(_rhs221, (_lhs178).Int())
						_lhs179.F9 = _rhs222
						var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen(((_this).F28()), 0))
						_ = _index168
						((_this).F28()).ArraySet1(((_this).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen(((_this).F28()), 0))).Int()).(_dafny.Int), (_index168).Int())
						(globalState).F9 = ((_this).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen(((_this).F28()), 0))).Int()).(_dafny.Int)
					}
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1344_v78 _dafny.Map
		_ = _1344_v78
		_1344_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1233_v2, _1238_v7)
		var _hi9 _dafny.Int = (_1344_v78).Cardinality()
		_ = _hi9
		for _1345_i15 := (Companion_Default___.Fm39((_this).F24(), _1233_v2, globalState)).Cardinality(); _1345_i15.Cmp(_hi9) < 0; _1345_i15 = _1345_i15.Plus(_dafny.One) {
			(globalState).F9 = (_this).F24()
			var _1346_v79 _dafny.Array
			_ = _1346_v79
			var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(16))
			_ = _nw193
			_1346_v79 = _nw193
			var _1347_v80 _dafny.Set
			_ = _1347_v80
			_1347_v80 = _dafny.SetOf(true)
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_1346_v79), 0))
			_ = _index169
			(_1346_v79).ArraySet1((Companion_Default___.Fm28(_1233_v2, (_dafny.Zero).Minus(_1345_i15), true, (_dafny.Zero).Minus(_1345_i15), globalState)).Intersection(_1347_v80), (_index169).Int())
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_1346_v79), 0))
			_ = _index170
			(_1346_v79).ArraySet1((_1347_v80).Union((_1347_v80).Difference(_dafny.SetOf(_1233_v2))), (_index170).Int())
			if (_1345_i15).Cmp((_this).F24()) != 0 {
				var _1348_v81 _dafny.Array
				_ = _1348_v81
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_31
				var _nw194 _dafny.Array
				_ = _nw194
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw194 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Map = (func(_1349_i15 _dafny.Int, _1350_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
						return func(_1351_i16 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_1349_i15, _dafny.IntOfInt64(297), _1349_i15, _1349_i15, _1349_i15)).Cardinality()), _1350_v8)
						}
					})(_1345_i15, _1239_v8)
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
				_1348_v81 = _nw194
				var _1352_v82 _dafny.Map
				_ = _1352_v82
				_1352_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1239_v8)
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_1348_v81), 0))
				_ = _index171
				(_1348_v81).ArraySet1(_1352_v82, (_index171).Int())
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_1348_v81), 0))
				_ = _index172
				(_1348_v81).ArraySet1((_1352_v82).Update(_dafny.IntOfInt64(-230), _1239_v8), (_index172).Int())
				(globalState).F9 = _1345_i15
				var _1353_v84 _dafny.MultiSet
				_ = _1353_v84
				_1353_v84 = _dafny.MultiSetOf((func() _dafny.Set {
					var _coll40 = _dafny.NewBuilder()
					_ = _coll40
					for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(719), _dafny.IntOfInt64(-613))); ; {
						_compr_40, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _1354_v83 _dafny.Int
						_1354_v83 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(719)).Cmp(_1354_v83) <= 0) && ((_1354_v83).Cmp(_dafny.IntOfInt64(-613)) < 0) {
							_coll40.Add(Companion_Default___.SafeModuloInt(_1354_v83, (_this).F24()))
						}
					}
					return _coll40.ToSet()
				}()).Equals(_1248_v15), _1233_v2, (_1233_v2) == (_1233_v2), !(_1233_v2) || (_1233_v2))
				var _1355_v85 *C2
				_ = _1355_v85
				var _nw195 *C2 = New_C2_()
				_ = _nw195
				_nw195.Ctor__(!(_1233_v2), Companion_Default___.SafeDivisionInt(_1345_i15, _1345_i15), _dafny.Companion_Sequence_.Update((_this).F25(), (Companion_Default___.SafeIndex(_1345_i15, _dafny.IntOfUint32(((_this).F25()).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1237_v6, (_this).F24())))
				_1355_v85 = _nw195
				var _rhs223 _dafny.Int = _1345_i15
				_ = _rhs223
				var _rhs224 _dafny.MultiSet = _dafny.MultiSetOf(!(_1233_v2))
				_ = _rhs224
				var _rhs225 bool = (func() bool {
					if (func() bool {
						if _1233_v2 {
							return _1233_v2
						}
						return _1233_v2
					})() {
						return _1233_v2
					}
					return !(_1355_v85.F31) || (_1355_v85.F31)
				})()
				_ = _rhs225
				var _rhs226 _dafny.Int = (Companion_Default___.Fm0(_1345_i15, (_this).F24(), globalState)).Plus(_1345_i15)
				_ = _rhs226
				var _rhs227 *C2 = _1355_v85
				_ = _rhs227
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				var _lhs181 *GlobalState = globalState
				_ = _lhs181
				_lhs180.F9 = _rhs223
				_1353_v84 = _rhs224
				_1233_v2 = _rhs225
				_lhs181.F14 = _rhs226
				_1355_v85 = _rhs227
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen(((_this).F28()), 0))
				_ = _index173
				((_this).F28()).ArraySet1(_1345_i15, (_index173).Int())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen(((_this).F28()), 0))
				_ = _index174
				((_this).F28()).ArraySet1(_1345_i15, (_index174).Int())
				var _1356_v86 D11
				_ = _1356_v86
				_1356_v86 = Companion_D11_.Create_DC30_(_1240_v9)
				var _1357_v87 _dafny.Map
				_ = _1357_v87
				_1357_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1239_v8, _1239_v8), _1233_v2)
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen(((_this).F28()), 0))
				_ = _index175
				var _rhs228 _dafny.Int = ((_this).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen(((_this).F28()), 0))).Int()).(_dafny.Int)
				_ = _rhs228
				var _rhs229 _dafny.CodePoint = _1237_v6
				_ = _rhs229
				var _rhs230 bool = (_dafny.MultiSetOf(_1240_v9)).IsProperSubsetOf(_dafny.MultiSetOf(_1240_v9, (_1356_v86).Dtor_cf47(), _1240_v9))
				_ = _rhs230
				var _rhs231 bool = (func() bool {
					if (_1357_v87).Contains(_1239_v8) {
						return (_1357_v87).Get(_1239_v8).(bool)
					}
					return !(_1233_v2)
				})()
				_ = _rhs231
				var _lhs182 _dafny.Array = (_this).F28()
				_ = _lhs182
				var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen(((_this).F28()), 0))
				_ = _lhs183
				var _lhs184 *C2 = _1355_v85
				_ = _lhs184
				var _lhs185 *GlobalState = globalState
				_ = _lhs185
				(_lhs182).ArraySet1(_rhs228, (_lhs183).Int())
				_1237_v6 = _rhs229
				_lhs184.F31 = _rhs230
				_lhs185.F2 = _rhs231
			} else {
				var _1358_v88 _dafny.Map
				_ = _1358_v88
				_1358_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1345_i15, _1242_v12)
				var _rhs232 bool = _1233_v2
				_ = _rhs232
				var _rhs233 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1239_v8, _1239_v8)
				_ = _rhs233
				var _rhs234 _dafny.Map = _1358_v88
				_ = _rhs234
				var _lhs186 *GlobalState = globalState
				_ = _lhs186
				var _lhs187 *GlobalState = globalState
				_ = _lhs187
				_lhs186.F2 = _rhs232
				_lhs187.F17 = _rhs233
				_1358_v88 = _rhs234
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1242_v12), 0))
				_ = _index176
				(_1242_v12).ArraySet1((_1234_v3).Select((Companion_Default___.SafeIndex(_1345_i15, _dafny.IntOfUint32((_1234_v3).Cardinality()))).Uint32()).(bool), (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1242_v12), 0))
				_ = _index177
				(_1242_v12).ArraySet1((_1345_i15).Cmp(((_1235_v4).Update(_1233_v2, (_this).F24())).Cardinality()) == 0, (_index177).Int())
				var _1359_v89 _dafny.Array
				_ = _1359_v89
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_32
				var _nw196 _dafny.Array
				_ = _nw196
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw196 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) bool = (func(_1360_v12 _dafny.Array, _1361_v2 bool) func(_dafny.Int) bool {
						return func(_1362_i17 _dafny.Int) bool {
							return !((_1360_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1360_v12), 0))).Int()).(bool)) || (_1361_v2)
						}
					})(_1242_v12, _1233_v2)
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
				_1359_v89 = _nw196
				_1359_v89 = _1359_v89
				var _1363_v90 D8
				_ = _1363_v90
				_1363_v90 = Companion_D8_.Create_DC25_(_1345_i15, _1233_v2, (_1242_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1242_v12), 0))).Int()).(bool))
				var _1364_v91 _dafny.MultiSet
				_ = _1364_v91
				_1364_v91 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_1239_v8, (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1239_v8).Cardinality()))).Uint32(), _dafny.CodePoint('a')), Companion_Default___.Fm4((_1242_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1242_v12), 0))).Int()).(bool), _1345_i15, _1237_v6, (_1363_v90).Dtor_cf40(), globalState), _dafny.UnicodeSeqOfUtf8Bytes("vhh"), _dafny.Companion_Sequence_.Concatenate(_1239_v8, _1239_v8), _1239_v8)
				var _1365_v92 _dafny.Sequence
				_ = _1365_v92
				_1365_v92 = _dafny.SeqOf(_1364_v91, (func() _dafny.MultiSet {
					if (_1242_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_1242_v12), 0))).Int()).(bool) {
						return _dafny.MultiSetOf(_1239_v8, _1239_v8)
					}
					return _1364_v91
				})(), _dafny.MultiSetOf(_1239_v8, _1239_v8, _dafny.UnicodeSeqOfUtf8Bytes("lqkivho")), ((_1364_v91).Update(_1239_v8, Companion_Default___.Abs(_1345_i15))).Intersection(_1364_v91))
				_1364_v91 = (_1365_v92).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1365_v92).Cardinality()))).Uint32()).(_dafny.MultiSet)
				_1239_v8 = _dafny.Companion_Sequence_.Concatenate(_1239_v8, _1239_v8)
			}
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index178
			((_this).F28()).ArraySet1((_this).F24(), (_index178).Int())
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen(((_this).F28()), 0))
			_ = _index179
			((_this).F28()).ArraySet1((_this).F24(), (_index179).Int())
		}
		var _1366_v93 _dafny.Sequence
		_ = _1366_v93
		_1366_v93 = _dafny.SeqOf(_1232_v1, _1232_v1)
		var _1367_v94 *C1
		_ = _1367_v94
		var _nw197 *C1 = New_C1_()
		_ = _nw197
		_nw197.Ctor__(_1366_v93, (_this).F24(), (_this).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(526))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg87 _dafny.Int) interface{} {
				return coer87(arg87)
			}
		}((func(_1368_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
			return func(_1369_i18 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1368_v6, (_this).F24())
			}
		})(_1237_v6))))
		_1367_v94 = _nw197
		var _1370_v95 _dafny.Map
		_ = _1370_v95
		_1370_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1367_v94)
		var _1371_v96 _dafny.Sequence
		_ = _1371_v96
		_1371_v96 = _dafny.SeqOf(Companion_Default___.Fm0((_this).F24(), (_1370_v95).Cardinality(), globalState))
		r0 = (_dafny.IntOfInt64(140)).Cmp((_1371_v96).Select((Companion_Default___.SafeIndex((_1248_v15).Cardinality(), _dafny.IntOfUint32((_1371_v96).Cardinality()))).Uint32()).(_dafny.Int)) == 0
		return r0
	}
}
func (_this *C4) M4(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Sequence, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 bool = false
		_ = r3
		var _1372_v0 _dafny.Map
		_ = _1372_v0
		_1372_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F24())
		var _1373_v1 _dafny.Sequence
		_ = _1373_v1
		_1373_v1 = _dafny.SeqOf(_1372_v0)
		var _1374_v2 _dafny.MultiSet
		_ = _1374_v2
		_1374_v2 = _dafny.MultiSetOf(p1, p1)
		var _1375_v3 _dafny.CodePoint
		_ = _1375_v3
		_1375_v3 = _dafny.CodePoint('u')
		var _1376_v4 _dafny.Map
		_ = _1376_v4
		_1376_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1375_v3)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))
		var _1377_v5 _dafny.Map
		_ = _1377_v5
		_1377_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_D3_.Create_DC11_((_this).F24(), p1, p1))
		var _1378_v6 D12
		_ = _1378_v6
		_1378_v6 = Companion_D12_.Create_DC32_(_1377_v5)
		var _1379_v7 _dafny.Set
		_ = _1379_v7
		_1379_v7 = _dafny.SetOf(p0)
		var _1380_v8 _dafny.Array
		_ = _1380_v8
		var _nwElement0_38 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1373_v1, _1373_v1)).Cardinality())
		_ = _nwElement0_38
		var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(16))
		_ = _nw198
		(_nw198).ArraySet1(_nwElement0_38, 0)
		(_nw198).ArraySet1((_this).F24(), 1)
		(_nw198).ArraySet1((_this).F24(), 2)
		(_nw198).ArraySet1((_this).F24(), 3)
		(_nw198).ArraySet1((_this).F24(), 4)
		(_nw198).ArraySet1((_this).F24(), 5)
		(_nw198).ArraySet1((_this).F24(), 6)
		(_nw198).ArraySet1((func() _dafny.Int {
			if (_1374_v2).Contains(p0) {
				return (_1374_v2).Multiplicity(p0)
			}
			return (_1376_v4).Cardinality()
		})(), 7)
		(_nw198).ArraySet1((func() _dafny.Int {
			if p1 {
				return (_dafny.Zero).Minus((_this).F24())
			}
			return (_this).F24()
		})(), 8)
		(_nw198).ArraySet1((_this).F24(), 9)
		(_nw198).ArraySet1(((_this).F24()).Minus(_dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality())), 10)
		(_nw198).ArraySet1(Companion_Default___.SafeDivisionInt(((_1378_v6).Dtor_cf53()).Cardinality(), (_this).F24()), 11)
		(_nw198).ArraySet1(((_1379_v7).Intersection(_1379_v7)).Cardinality(), 12)
		(_nw198).ArraySet1((_this).F24(), 13)
		(_nw198).ArraySet1((_this).F24(), 14)
		(_nw198).ArraySet1((_this).F24(), 15)
		_1380_v8 = _nw198
		for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1380_v8), 0))); ; {
			_guard_loop_3, _ok44 := _iter44()
			if !_ok44 {
				break
			}
			var _1381_i0 _dafny.Int
			_1381_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1381_i0).Sign() != -1) && ((_1381_i0).Cmp(_dafny.ArrayLen((_1380_v8), 0)) < 0)) {
				(_1380_v8).ArraySet1((_1381_i0).Times(Companion_Default___.SafeModuloInt((_this).F24(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("spivrnx")).Cardinality())))), (_1381_i0).Int())
			}
		}
		var _1382_v9 *C0
		_ = _1382_v9
		var _nw199 *C0 = New_C0_()
		_ = _nw199
		_nw199.Ctor__(true)
		_1382_v9 = _nw199
		var _hi10 _dafny.Int = (_dafny.Zero).Minus((_this).F24())
		_ = _hi10
		for _1383_i1 := (_dafny.Zero).Minus((_dafny.SetOf(p1)).Cardinality()); _1383_i1.Cmp(_hi10) < 0; _1383_i1 = _1383_i1.Plus(_dafny.One) {
			if (_1382_v9).F34() {
				_1375_v3 = _1375_v3
				var _1384_v10 _dafny.Set
				_ = _1384_v10
				_1384_v10 = _dafny.SetOf((_this).F28())
				var _1385_v11 _dafny.Map
				_ = _1385_v11
				_1385_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1386_v12 _dafny.Sequence
				_ = _1386_v12
				_1386_v12 = _dafny.SeqOf((_1382_v9).F34())
				var _1387_v13 _dafny.Array
				_ = _1387_v13
				var _nwElement0_39 bool = p1
				_ = _nwElement0_39
				var _nw200 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(27))
				_ = _nw200
				(_nw200).ArraySet1(_nwElement0_39, 0)
				(_nw200).ArraySet1((_1382_v9).F34(), 1)
				(_nw200).ArraySet1((_1382_v9).F34(), 2)
				(_nw200).ArraySet1(p1, 3)
				(_nw200).ArraySet1((_1382_v9).F34(), 4)
				(_nw200).ArraySet1((_1382_v9).F34(), 5)
				(_nw200).ArraySet1(Companion_Default___.Fm2(_dafny.IntOfInt64(111), _1383_i1, _1375_v3, globalState), 6)
				(_nw200).ArraySet1((_1382_v9).F34(), 7)
				(_nw200).ArraySet1((func() bool {
					if (_1385_v11).Contains(p1) {
						return (_1385_v11).Get(p1).(bool)
					}
					return p0
				})(), 8)
				(_nw200).ArraySet1(p1, 9)
				(_nw200).ArraySet1((_1382_v9).F34(), 10)
				(_nw200).ArraySet1(p0, 11)
				(_nw200).ArraySet1(p0, 12)
				(_nw200).ArraySet1(false, 13)
				(_nw200).ArraySet1((_1382_v9).F34(), 14)
				(_nw200).ArraySet1(true, 15)
				(_nw200).ArraySet1(false, 16)
				(_nw200).ArraySet1(p1, 17)
				(_nw200).ArraySet1(p0, 18)
				(_nw200).ArraySet1(!(true), 19)
				(_nw200).ArraySet1((_1382_v9).F34(), 20)
				(_nw200).ArraySet1((_1382_v9).F34(), 21)
				(_nw200).ArraySet1((_1382_v9).F34(), 22)
				(_nw200).ArraySet1((_1386_v12).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1382_v9).F34(), (_1382_v9).F34())).Cardinality(), _dafny.IntOfUint32((_1386_v12).Cardinality()))).Uint32()).(bool), 23)
				(_nw200).ArraySet1((_1382_v9).F34(), 24)
				(_nw200).ArraySet1((_1382_v9).F34(), 25)
				(_nw200).ArraySet1((Companion_D6_.Create_DC17_(p0)).Dtor_cf28(), 26)
				_1387_v13 = _nw200
				var _1388_v14 _dafny.Map
				_ = _1388_v14
				_1388_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1384_v10).Cardinality(), _1387_v13)
				_1388_v14 = (_1388_v14).Update(Companion_Default___.Fm0(_dafny.IntOfInt64(951), _dafny.IntOfInt64(-204), globalState), _1387_v13)
				(globalState).F2 = true
				var _1389_v15 *C0
				_ = _1389_v15
				var _nw201 *C0 = New_C0_()
				_ = _nw201
				_nw201.Ctor__((_1382_v9).F34())
				_1389_v15 = _nw201
				(globalState).F1 = (_this).F24()
			} else {
				(globalState).F18 = (_this).F24()
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F28()), 0))
				_ = _index180
				((_this).F28()).ArraySet1(Companion_Default___.Fm0((_dafny.Zero).Minus(Companion_Default___.Fm0((_this).F24(), _dafny.IntOfInt64(572), globalState)), (_this).F24(), globalState), (_index180).Int())
				var _1390_v16 _dafny.Array
				_ = _1390_v16
				var _nw202 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(27))
				_ = _nw202
				_1390_v16 = _nw202
				var _1391_v17 _dafny.Map
				_ = _1391_v17
				_1391_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1390_v16, (_1382_v9).F34())
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F28()), 0))
				_ = _index181
				((_this).F28()).ArraySet1(((func() _dafny.Int {
					if p0 {
						return _1383_i1
					}
					return _dafny.IntOfInt64(-994)
				})()).Minus((_1391_v17).Cardinality()), (_index181).Int())
				var _1392_v18 _dafny.Array
				_ = _1392_v18
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_33
				var _nw203 _dafny.Array
				_ = _nw203
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw203 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1393_v9 *C0, _1394_p0 bool) func(_dafny.Int) bool {
						return func(_1395_i2 _dafny.Int) bool {
							return ((_1393_v9).F34()) || (_1394_p0)
						}
					})(_1382_v9, p0)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw203 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw203).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw203).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1392_v18 = _nw203
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_1392_v18), 0))
				_ = _index182
				(_1392_v18).ArraySet1((_1382_v9).F34(), (_index182).Int())
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_1392_v18), 0))
				_ = _index183
				(_1392_v18).ArraySet1((_1382_v9).F34(), (_index183).Int())
				_1392_v18 = _1392_v18
				var _1396_v19 _dafny.Array
				_ = _1396_v19
				var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
				_ = _nw204
				_1396_v19 = _nw204
				var _1397_v20 _dafny.Array
				_ = _1397_v20
				var _nw205 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
				_ = _nw205
				_1397_v20 = _nw205
				var _1398_v21 _dafny.Map
				_ = _1398_v21
				_1398_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1375_v3, _1383_i1)
				var _1399_v22 T0
				_ = _1399_v22
				var _nw206 *C3 = New_C3_()
				_ = _nw206
				_nw206.Ctor__(_1396_v19, _1397_v20, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-388)), ((_this).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F28()), 0))).Int()).(_dafny.Int)), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F25(), _dafny.Companion_Sequence_.Update((_this).F25(), (Companion_Default___.SafeIndex(_1383_i1, _dafny.IntOfUint32(((_this).F25()).Cardinality()))).Uint32(), _1398_v21)), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F25(), _dafny.Companion_Sequence_.Update((_this).F25(), (Companion_Default___.SafeIndex(_1383_i1, _dafny.IntOfUint32(((_this).F25()).Cardinality()))).Uint32(), _1398_v21))).Cardinality()))).Uint32(), (_1398_v21).Update(_1375_v3, ((_this).F28()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F28()), 0))).Int()).(_dafny.Int))))
				_1399_v22 = _nw206
				_1399_v22 = _1399_v22
			}
			var _1400_v23 _dafny.Sequence
			_ = _1400_v23
			_1400_v23 = _dafny.UnicodeSeqOfUtf8Bytes("epg")
			var _1401_v24 _dafny.Map
			_ = _1401_v24
			_1401_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_this).F24(), _dafny.IntOfUint32((_1400_v23).Cardinality())), true)
			var _1402_v25 _dafny.Map
			_ = _1402_v25
			_1402_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1382_v9).F34(), _1379_v7)
			_1401_v24 = (_1401_v24).Update((func() _dafny.Int {
				if (_1382_v9).F34() {
					return _1383_i1
				}
				return (_1402_v25).Cardinality()
			})(), !((_1382_v9).F34()) || (!(p1)))
			r1 = _dafny.IntOfInt64(602)
			var _1403_v26 _dafny.MultiSet
			_ = _1403_v26
			_1403_v26 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F24()))
			_1401_v24 = (_1401_v24).Update((_1403_v26).Cardinality(), false)
		}
		var _1404_v27 _dafny.Sequence
		_ = _1404_v27
		_1404_v27 = _dafny.UnicodeSeqOfUtf8Bytes("wlyo")
		var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1380_v8), 0))
		_ = _index184
		(_1380_v8).ArraySet1(Companion_Default___.SafeModuloInt((_this).F24(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1404_v27).Cardinality()))), (_index184).Int())
		var _1405_v28 _dafny.Sequence
		_ = _1405_v28
		_1405_v28 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_1404_v27).Cardinality())))
		var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1380_v8), 0))
		_ = _index185
		(_1380_v8).ArraySet1((Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState)).Times((_1405_v28).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1405_v28).Cardinality()))).Uint32()).(_dafny.Int)), (_index185).Int())
		var _1406_v29 _dafny.Map
		_ = _1406_v29
		_1406_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1375_v3, false)
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1380_v8), 0))
		_ = _index186
		(_1380_v8).ArraySet1(Companion_Default___.Fm0(((_1406_v29).Cardinality()).Times((_this).F24()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(148))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg88 _dafny.Int) interface{} {
				return coer88(arg88)
			}
		}((func(_1407_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1408_i3 _dafny.Int) _dafny.CodePoint {
				return _1407_v3
			}
		})(_1375_v3)))).Cardinality()), globalState), (_index186).Int())
		var _1409_v30 _dafny.Array
		_ = _1409_v30
		var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
		_ = _nw207
		_1409_v30 = _nw207
		var _1410_v31 _dafny.Sequence
		_ = _1410_v31
		_1410_v31 = _dafny.SeqOf(p0)
		var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1409_v30), 0))
		_ = _index187
		(_1409_v30).ArraySet1((func() _dafny.Sequence {
			if (_1382_v9).F34() {
				return _1410_v31
			}
			return _1410_v31
		})(), (_index187).Int())
		var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1409_v30), 0))
		_ = _index188
		(_1409_v30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1410_v31, _1410_v31), (_index188).Int())
		r0 = (_1380_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1380_v8), 0))).Int()).(_dafny.Int)
		r1 = (_this).F24()
		var _1411_v32 _dafny.Sequence
		_ = _1411_v32
		_1411_v32 = _dafny.SeqOf(_1405_v28, _1405_v28)
		var _1412_v33 D4
		_ = _1412_v33
		_1412_v33 = Companion_D4_.Create_DC12_(_dafny.SeqOf((_1380_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1380_v8), 0))).Int()).(_dafny.Int), (_1380_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1380_v8), 0))).Int()).(_dafny.Int), (_this).F24(), (_this).F24()))
		var _1413_v34 _dafny.Sequence
		_ = _1413_v34
		_1413_v34 = _dafny.SeqOf(_1405_v28, (_1411_v32).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1411_v32).Cardinality()))).Uint32()).(_dafny.Sequence), (_1412_v33).Dtor_cf21(), _dafny.Companion_Sequence_.Concatenate(_1405_v28, _1405_v28))
		r2 = (_1411_v32).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1411_v32).Cardinality()))).Uint32()).(_dafny.Sequence)
		r3 = p0
		return r0, r1, r2, r3
	}
}
func (_this *C4) F28() _dafny.Array {
	{
		return _this._f28
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f24 _dafny.Int
	_f25 _dafny.Sequence
	F27  _dafny.Int
	_f26 _dafny.Map
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f24 = _dafny.Zero
	_this._f25 = _dafny.EmptySeq
	_this.F27 = _dafny.Zero
	_this._f26 = _dafny.EmptyMap
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

func (_this *C5) F24() _dafny.Int {
	return _this._f24
}
func (_this *C5) F25() _dafny.Sequence {
	return _this._f25
}
func (_this *C5) Ctor__(f26 _dafny.Map, f27 _dafny.Int, f24 _dafny.Int, f25 _dafny.Sequence) {
	{
		(_this)._f26 = f26
		(_this).F27 = f27
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C5) M1(globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _hi11 _dafny.Int = _this.F27
		_ = _hi11
		for _1414_i0 := Companion_Default___.SafeModuloInt(_this.F27, _this.F27); _1414_i0.Cmp(_hi11) < 0; _1414_i0 = _1414_i0.Plus(_dafny.One) {
			(globalState).F14 = ((_this).F24()).Minus(_this.F27)
			if Companion_Default___.Fm2((_this.F27).Plus(_dafny.IntOfInt64(-676)), (_this).F24(), Companion_Default___.Fm3(globalState), globalState) {
				var _1415_v0 _dafny.Array
				_ = _1415_v0
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_34
				var _nw208 _dafny.Array
				_ = _nw208
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw208 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) bool = func(_1416_i1 _dafny.Int) bool {
						return !(!(true))
					}
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw208 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw208).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw208).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1415_v0 = _nw208
				var _1417_v1 bool
				_ = _1417_v1
				_1417_v1 = false
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1415_v0), 0))
				_ = _index189
				(_1415_v0).ArraySet1(_1417_v1, (_index189).Int())
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1415_v0), 0))
				_ = _index190
				(_1415_v0).ArraySet1(!(_1417_v1) || (_1417_v1), (_index190).Int())
				var _1418_v2 _dafny.Array
				_ = _1418_v2
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
				_ = _nw209
				_1418_v2 = _nw209
				var _1419_v3 _dafny.Array
				_ = _1419_v3
				var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
				_ = _nw210
				_1419_v3 = _nw210
				var _1420_v4 T0
				_ = _1420_v4
				var _nw211 *C4 = New_C4_()
				_ = _nw211
				_nw211.Ctor__(_1419_v3, _dafny.IntOfUint32((_dafny.SeqOf(_this.F27)).Cardinality()), (_this).F25())
				_1420_v4 = _nw211
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1418_v2), 0))
				_ = _index191
				(_1418_v2).ArraySet1(_dafny.SeqOf(_1420_v4), (_index191).Int())
				var _1421_v5 _dafny.MultiSet
				_ = _1421_v5
				_1421_v5 = _dafny.MultiSetOf(_1417_v1)
				var _1422_v6 _dafny.Sequence
				_ = _1422_v6
				_1422_v6 = _dafny.SeqOf(_1420_v4, _1420_v4)
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1418_v2), 0))
				_ = _index192
				(_1418_v2).ArraySet1((func() _dafny.Sequence {
					if (_1421_v5).IsDisjointFrom(_1421_v5) {
						return _dafny.Companion_Sequence_.Concatenate(_1422_v6, _1422_v6)
					}
					return _1422_v6
				})(), (_index192).Int())
				var _1423_v7 _dafny.Array
				_ = _1423_v7
				var _nw212 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
				_ = _nw212
				_1423_v7 = _nw212
				var _1424_v8 _dafny.CodePoint
				_ = _1424_v8
				_1424_v8 = _dafny.CodePoint('d')
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_1423_v7), 0))
				_ = _index193
				(_1423_v7).ArraySet1CodePoint(_1424_v8, (_index193).Int())
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_1423_v7), 0))
				_ = _index194
				(_1423_v7).ArraySet1CodePoint(_1424_v8, (_index194).Int())
				var _1425_v9 _dafny.Map
				_ = _1425_v9
				_1425_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1420_v4).F24(), (_1420_v4).F24())
				var _1426_v10 _dafny.Sequence
				_ = _1426_v10
				_1426_v10 = _dafny.UnicodeSeqOfUtf8Bytes("iccnie")
				var _1427_v11 _dafny.Sequence
				_ = _1427_v11
				_1427_v11 = _dafny.SeqOf(_dafny.IntOfUint32((_1426_v10).Cardinality()))
				var _1428_v12 _dafny.MultiSet
				_ = _1428_v12
				_1428_v12 = _dafny.MultiSetOf((func() _dafny.Int {
					if (_1425_v9).Contains((_1420_v4).F24()) {
						return (_1425_v9).Get((_1420_v4).F24()).(_dafny.Int)
					}
					return (_1420_v4).F24()
				})(), (_1427_v11).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1427_v11).Cardinality()))).Uint32()).(_dafny.Int))
				var _1429_v13 _dafny.Int
				_ = _1429_v13
				var _out31 _dafny.Int
				_ = _out31
				_out31 = Companion_Default___.M0((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_1414_i0, (_1428_v12).Cardinality())), Companion_Default___.SafeDivisionInt(_1414_i0, Companion_Default___.Fm0((_dafny.Zero).Minus(_dafny.IntOfUint32((_1426_v10).Cardinality())), _1414_i0, globalState)), globalState)
				_1429_v13 = _out31
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_1415_v0), 0))
				_ = _index195
				(_1415_v0).ArraySet1(_1417_v1, (_index195).Int())
			} else {
				(globalState).F9 = _this.F27
				var _1430_v14 bool
				_ = _1430_v14
				_1430_v14 = true
				var _1431_v15 *C0
				_ = _1431_v15
				var _nw213 *C0 = New_C0_()
				_ = _nw213
				_nw213.Ctor__(_1430_v14)
				_1431_v15 = _nw213
				var _1432_v16 D6
				_ = _1432_v16
				_1432_v16 = Companion_D6_.Create_DC18_(!((_1431_v15).F34()))
				var _1433_v17 D6
				_ = _1433_v17
				_1433_v17 = Companion_D6_.Create_DC19_(_1432_v16)
				var _1434_v18 _dafny.Sequence
				_ = _1434_v18
				_1434_v18 = _dafny.UnicodeSeqOfUtf8Bytes("rflync")
				var _1435_v19 _dafny.Map
				_ = _1435_v19
				_1435_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1434_v18, Companion_D6_.Create_DC18_((_1431_v15).F34()))
				var _pat_let_tv16 = _1432_v16
				_ = _pat_let_tv16
				var _pat_let_tv17 = _1435_v19
				_ = _pat_let_tv17
				var _pat_let_tv18 = _1434_v18
				_ = _pat_let_tv18
				var _pat_let_tv19 = _1435_v19
				_ = _pat_let_tv19
				var _pat_let_tv20 = _1434_v18
				_ = _pat_let_tv20
				var _pat_let_tv21 = _1432_v16
				_ = _pat_let_tv21
				var _pat_let_tv22 = _1432_v16
				_ = _pat_let_tv22
				var _1436_v20 _dafny.Array
				_ = _1436_v20
				var _nwElement0_40 D6 = _1433_v17
				_ = _nwElement0_40
				var _nw214 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(23))
				_ = _nw214
				(_nw214).ArraySet1(_nwElement0_40, 0)
				(_nw214).ArraySet1(_1433_v17, 1)
				(_nw214).ArraySet1(_1433_v17, 2)
				(_nw214).ArraySet1(Companion_D6_.Create_DC19_(_1432_v16), 3)
				(_nw214).ArraySet1(_1433_v17, 4)
				(_nw214).ArraySet1(_1433_v17, 5)
				(_nw214).ArraySet1(_1433_v17, 6)
				(_nw214).ArraySet1(Companion_D6_.Create_DC19_(_1432_v16), 7)
				(_nw214).ArraySet1(_1433_v17, 8)
				(_nw214).ArraySet1(_1433_v17, 9)
				(_nw214).ArraySet1(Companion_D6_.Create_DC19_(_1432_v16), 10)
				(_nw214).ArraySet1(_1433_v17, 11)
				(_nw214).ArraySet1(func(_pat_let20_0 D6) D6 {
					return func(_1437_dt__update__tmp_h0 D6) D6 {
						return func(_pat_let21_0 D6) D6 {
							return func(_1438_dt__update_hcf30_h0 D6) D6 {
								return Companion_D6_.Create_DC19_(_1438_dt__update_hcf30_h0)
							}(_pat_let21_0)
						}(_pat_let_tv16)
					}(_pat_let20_0)
				}(_1433_v17), 12)
				(_nw214).ArraySet1(_1433_v17, 13)
				(_nw214).ArraySet1(_1433_v17, 14)
				(_nw214).ArraySet1((func() D6 {
					if false {
						return _1433_v17
					}
					return _1433_v17
				})(), 15)
				(_nw214).ArraySet1(_1433_v17, 16)
				(_nw214).ArraySet1(_1433_v17, 17)
				(_nw214).ArraySet1(_1433_v17, 18)
				(_nw214).ArraySet1(func(_pat_let22_0 D6) D6 {
					return func(_1439_dt__update__tmp_h1 D6) D6 {
						return func(_pat_let23_0 D6) D6 {
							return func(_1440_dt__update_hcf30_h1 D6) D6 {
								return Companion_D6_.Create_DC19_(_1440_dt__update_hcf30_h1)
							}(_pat_let23_0)
						}((func() D6 {
							if (_pat_let_tv17).Contains(_pat_let_tv18) {
								return (_pat_let_tv19).Get(_pat_let_tv20).(D6)
							}
							return _pat_let_tv21
						})())
					}(_pat_let22_0)
				}(Companion_D6_.Create_DC19_(_1432_v16)), 19)
				(_nw214).ArraySet1(func(_pat_let24_0 D6) D6 {
					return func(_1441_dt__update__tmp_h2 D6) D6 {
						return func(_pat_let25_0 D6) D6 {
							return func(_1442_dt__update_hcf30_h2 D6) D6 {
								return Companion_D6_.Create_DC19_(_1442_dt__update_hcf30_h2)
							}(_pat_let25_0)
						}(_pat_let_tv22)
					}(_pat_let24_0)
				}(_1433_v17), 20)
				(_nw214).ArraySet1(_1433_v17, 21)
				(_nw214).ArraySet1(_1433_v17, 22)
				_1436_v20 = _nw214
				_1436_v20 = _1436_v20
				var _1443_v21 _dafny.Sequence
				_ = _1443_v21
				_1443_v21 = _dafny.SeqOf((_1431_v15).F34())
				(globalState).F21 = (func() _dafny.Sequence {
					if (_1431_v15).F34() {
						return _1443_v21
					}
					return _dafny.Companion_Sequence_.Concatenate(_1443_v21, _1443_v21)
				})()
				var _1444_v22 _dafny.Array
				_ = _1444_v22
				var _nw215 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw215
				_1444_v22 = _nw215
				var _1445_v23 _dafny.MultiSet
				_ = _1445_v23
				_1445_v23 = _dafny.MultiSetOf(_1444_v22, _1444_v22, _1444_v22, _1444_v22)
				_1445_v23 = (_1445_v23).Update(_1444_v22, Companion_Default___.Abs(((func() _dafny.Set {
					if (_1431_v15).F34() {
						return _dafny.SetOf(_1444_v22)
					}
					return _dafny.SetOf(_1444_v22, _1444_v22)
				})()).Cardinality()))
			}
			(globalState).F1 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(891), _this.F27)
			var _1446_v24 bool
			_ = _1446_v24
			_1446_v24 = false
			if _1446_v24 {
				_1446_v24 = _1446_v24
				var _1447_v25 bool
				_ = _1447_v25
				var _out32 bool
				_ = _out32
				_out32 = (_this).M2(globalState)
				_1447_v25 = _out32
				var _1448_v26 _dafny.Array
				_ = _1448_v26
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_35
				var _nw216 _dafny.Array
				_ = _nw216
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw216 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) _dafny.Int = (func(_1449_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1450_i2 _dafny.Int) _dafny.Int {
							return (_1450_i2).Plus(_1449_i0)
						}
					})(_1414_i0)
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
				_1448_v26 = _nw216
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_1448_v26), 0))
				_ = _index196
				(_1448_v26).ArraySet1(_this.F27, (_index196).Int())
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_1448_v26), 0))
				_ = _index197
				(_1448_v26).ArraySet1(_1414_i0, (_index197).Int())
				var _1451_v27 _dafny.CodePoint
				_ = _1451_v27
				_1451_v27 = _dafny.CodePoint('o')
				var _1452_v28 _dafny.Map
				_ = _1452_v28
				_1452_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1451_v27, _1414_i0)
				var _1453_v29 *C4
				_ = _1453_v29
				var _nw217 *C4 = New_C4_()
				_ = _nw217
				_nw217.Ctor__(_1448_v26, (_1448_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_1448_v26), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F25(), (_this).F25()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-423), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F25(), (_this).F25())).Cardinality()))).Uint32(), _1452_v28))
				_1453_v29 = _nw217
				(globalState).F9 = (_this).F24()
			} else {
				var _1454_v30 _dafny.Map
				_ = _1454_v30
				_1454_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1446_v24, _1414_i0)
				var _1455_v31 _dafny.Array
				_ = _1455_v31
				var _nwElement0_41 _dafny.Int = _this.F27
				_ = _nwElement0_41
				var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(2))
				_ = _nw218
				(_nw218).ArraySet1(_nwElement0_41, 0)
				(_nw218).ArraySet1((func() _dafny.Int {
					if (_1454_v30).Contains(_1446_v24) {
						return (_1454_v30).Get(_1446_v24).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(_this.F27)
				})(), 1)
				_1455_v31 = _nw218
				var _1456_v32 T0
				_ = _1456_v32
				var _nw219 *C4 = New_C4_()
				_ = _nw219
				_nw219.Ctor__(_1455_v31, (_this).F24(), (_this).F25())
				_1456_v32 = _nw219
				_1456_v32 = _1456_v32
				var _rhs235 bool = _1446_v24
				_ = _rhs235
				var _rhs236 _dafny.Int = Companion_Default___.SafeModuloInt(_this.F27, _1414_i0)
				_ = _rhs236
				var _rhs237 bool = (func() bool {
					if _1446_v24 {
						return false
					}
					return false
				})()
				_ = _rhs237
				var _lhs188 *GlobalState = globalState
				_ = _lhs188
				var _lhs189 *GlobalState = globalState
				_ = _lhs189
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				_lhs188.F2 = _rhs235
				_lhs189.F18 = _rhs236
				_lhs190.F2 = _rhs237
				var _1457_v33 _dafny.Sequence
				_ = _1457_v33
				_1457_v33 = _dafny.SeqOf(_1446_v24)
				var _1458_v34 _dafny.MultiSet
				_ = _1458_v34
				_1458_v34 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1457_v33).Cardinality()), (_this).F24(), _1414_i0, ((_1456_v32).F24()).Times((_this).F24()), (_dafny.IntOfInt64(564)).Plus(_this.F27))
				_1458_v34 = _1458_v34
				(globalState).F17 = _dafny.UnicodeSeqOfUtf8Bytes("tbquywj")
				var _1459_v35 *C4
				_ = _1459_v35
				var _nw220 *C4 = New_C4_()
				_ = _nw220
				_nw220.Ctor__(_1455_v31, (_1456_v32).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(462))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg89 _dafny.Int) interface{} {
						return coer89(arg89)
					}
				}(func(_1460_i3 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), (_this).F24())
				})))
				_1459_v35 = _nw220
				var _1461_v36 _dafny.Map
				_ = _1461_v36
				_1461_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(964), _1459_v35)
				_1461_v36 = _1461_v36
			}
		}
		var _1462_v37 bool
		_ = _1462_v37
		_1462_v37 = false
		var _1463_v38 _dafny.Sequence
		_ = _1463_v38
		_1463_v38 = _dafny.SeqOf((_this).F24())
		var _pat_let_tv23 = _1462_v37
		_ = _pat_let_tv23
		var _pat_let_tv24 = _1463_v38
		_ = _pat_let_tv24
		var _1464_v39 _dafny.Map
		_ = _1464_v39
		_1464_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1462_v37, func(_pat_let26_0 D3) D3 {
			return func(_1465_dt__update__tmp_h3 D3) D3 {
				return func(_pat_let27_0 bool) D3 {
					return func(_1466_dt__update_hcf20_h0 bool) D3 {
						return func(_pat_let28_0 _dafny.Int) D3 {
							return func(_1467_dt__update_hcf18_h0 _dafny.Int) D3 {
								return Companion_D3_.Create_DC11_(_1467_dt__update_hcf18_h0, (_1465_dt__update__tmp_h3).Dtor_cf19(), _1466_dt__update_hcf20_h0)
							}(_pat_let28_0)
						}(_dafny.IntOfUint32((_pat_let_tv24).Cardinality()))
					}(_pat_let27_0)
				}(_pat_let_tv23)
			}(_pat_let26_0)
		}(Companion_D3_.Create_DC11_(_this.F27, _1462_v37, _1462_v37)))
		var _1468_v40 D12
		_ = _1468_v40
		_1468_v40 = Companion_D12_.Create_DC32_(_1464_v39)
		var _source15 D12 = _1468_v40
		_ = _source15
		if _source15.Is_DC33() {
			var _1469___mcc_h0 _dafny.CodePoint = _source15.Get_().(D12_DC33).Cf54
			_ = _1469___mcc_h0
			var _1470___mcc_h1 bool = _source15.Get_().(D12_DC33).Cf55
			_ = _1470___mcc_h1
			var _1471_cf55 bool = _1470___mcc_h1
			_ = _1471_cf55
			var _1472_cf54 _dafny.CodePoint = _1469___mcc_h0
			_ = _1472_cf54
			var _1473_v41 _dafny.Set
			_ = _1473_v41
			_1473_v41 = _dafny.SetOf(_this.F27)
			_1471_cf55 = !(!(((_dafny.SetOf(_this.F27, (_this).F24())).Difference(_1473_v41)).IsSubsetOf(_1473_v41)))
			var _1474_v42 _dafny.Sequence
			_ = _1474_v42
			_1474_v42 = _dafny.UnicodeSeqOfUtf8Bytes("tdwmjxld")
			r0 = Companion_Default___.Fm2((_dafny.IntOfUint32((_1474_v42).Cardinality())).Times((_this).F24()), (_this.F27).Times(_dafny.IntOfInt64(891)), _1472_cf54, globalState)
			var _1475_v43 _dafny.Array
			_ = _1475_v43
			var _nw221 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw221
			_1475_v43 = _nw221
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1475_v43), 0))
			_ = _index198
			(_1475_v43).ArraySet1(_this.F27, (_index198).Int())
			var _1476_v44 _dafny.Map
			_ = _1476_v44
			_1476_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1471_cf55, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(114))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg90 _dafny.Int) interface{} {
					return coer90(arg90)
				}
			}(func(_1477_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality()))
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1475_v43), 0))
			_ = _index199
			(_1475_v43).ArraySet1(((_1476_v44).Cardinality()).Plus(Companion_Default___.SafeModuloInt((_this).F24(), _this.F27)), (_index199).Int())
			(_this).F27 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(613))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg91 _dafny.Int) interface{} {
					return coer91(arg91)
				}
			}(func(_1478_i5 _dafny.Int) _dafny.Int {
				return (_this).F24()
			}))).Cardinality()))
		} else if _source15.Is_DC32() {
			var _1479___mcc_h2 _dafny.Map = _source15.Get_().(D12_DC32).Cf53
			_ = _1479___mcc_h2
			var _1480_cf53 _dafny.Map = _1479___mcc_h2
			_ = _1480_cf53
			var _1481_v45 _dafny.Sequence
			_ = _1481_v45
			_1481_v45 = _dafny.SeqOf(true, _1462_v37)
			var _1482_v46 _dafny.Set
			_ = _1482_v46
			_1482_v46 = _dafny.SetOf(_1481_v45)
			r0 = (_1482_v46).IsProperSubsetOf(_1482_v46)
			var _1483_v47 _dafny.Sequence
			_ = _1483_v47
			_1483_v47 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_1463_v38))
			var _1484_v48 *C1
			_ = _1484_v48
			var _nw222 *C1 = New_C1_()
			_ = _nw222
			_nw222.Ctor__(_1483_v47, _this.F27, _this.F27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-502))).Uint32(), func(coer92 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg92 _dafny.Int) interface{} {
					return coer92(arg92)
				}
			}(func(_1485_i6 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), (_this).F24())
			})))
			_1484_v48 = _nw222
			var _1486_v49 _dafny.CodePoint
			_ = _1486_v49
			_1486_v49 = _dafny.CodePoint('w')
			var _1487_v51 _dafny.MultiSet
			_ = _1487_v51
			_1487_v51 = _dafny.MultiSetOf((_this).F24())
			var _1488_v52 D13
			_ = _1488_v52
			_1488_v52 = Companion_D13_.Create_DC35_(_1487_v51)
			var _1489_v53 _dafny.Sequence
			_ = _1489_v53
			_1489_v53 = _dafny.SeqOf(Companion_D3_.Create_DC11_(_dafny.IntOfInt64(950), _1462_v37, Companion_Default___.Fm2(_dafny.IntOfInt64(543), _this.F27, _1486_v49, globalState)))
			var _1490_v54 _dafny.MultiSet
			_ = _1490_v54
			_1490_v54 = _dafny.MultiSetOf(_1462_v37, _1462_v37)
			var _1491_v55 _dafny.Array
			_ = _1491_v55
			var _nwElement0_42 bool = !(true)
			_ = _nwElement0_42
			var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(16))
			_ = _nw223
			(_nw223).ArraySet1(_nwElement0_42, 0)
			(_nw223).ArraySet1(_1462_v37, 1)
			(_nw223).ArraySet1(_1462_v37, 2)
			(_nw223).ArraySet1(Companion_Default___.Fm2(_1484_v48.F33, _1484_v48.F33, _1486_v49, globalState), 3)
			(_nw223).ArraySet1(_1462_v37, 4)
			(_nw223).ArraySet1((_dafny.SetOf(_1484_v48.F33, _1484_v48.F33)).IsDisjointFrom(func() _dafny.Set {
				var _coll41 = _dafny.NewBuilder()
				_ = _coll41
				for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(794), _dafny.IntOfInt64(-319))); ; {
					_compr_41, _ok45 := _iter45()
					if !_ok45 {
						break
					}
					var _1492_v50 _dafny.Int
					_1492_v50 = interface{}(_compr_41).(_dafny.Int)
					if ((_dafny.IntOfInt64(794)).Cmp(_1492_v50) <= 0) && ((_1492_v50).Cmp(_dafny.IntOfInt64(-319)) < 0) {
						_coll41.Add(Companion_Default___.SafeDivisionInt(_1492_v50, (_this).F24()))
					}
				}
				return _coll41.ToSet()
			}()), 5)
			(_nw223).ArraySet1(true, 6)
			(_nw223).ArraySet1(true, 7)
			(_nw223).ArraySet1(((_1488_v52).Dtor_cf57()).IsDisjointFrom(_1487_v51), 8)
			(_nw223).ArraySet1(((_1487_v51).Update(_dafny.IntOfUint32((_1489_v53).Cardinality()), Companion_Default___.Abs(_1484_v48.F33))).IsSubsetOf(_1487_v51), 9)
			(_nw223).ArraySet1((_1484_v48.F33).Cmp(_this.F27) != 0, 10)
			(_nw223).ArraySet1((_1490_v54).IsDisjointFrom(_1490_v54), 11)
			(_nw223).ArraySet1(!(Companion_Default___.Fm2((_1490_v54).Cardinality(), _this.F27, _1486_v49, globalState)), 12)
			(_nw223).ArraySet1(_1462_v37, 13)
			(_nw223).ArraySet1((_1481_v45).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F27), _dafny.IntOfUint32((_1481_v45).Cardinality()))).Uint32()).(bool), 14)
			(_nw223).ArraySet1(_1462_v37, 15)
			_1491_v55 = _nw223
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1491_v55), 0))
			_ = _index200
			(_1491_v55).ArraySet1((_this.F27).Cmp((_this).F24()) <= 0, (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1491_v55), 0))
			_ = _index201
			(_1491_v55).ArraySet1(_1462_v37, (_index201).Int())
			var _1493_v56 _dafny.Map
			_ = _1493_v56
			_1493_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1486_v49, _1462_v37)
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1491_v55), 0))
			_ = _index202
			var _rhs238 _dafny.Int = Companion_Default___.SafeDivisionInt(_this.F27, Companion_Default___.SafeDivisionInt(_1484_v48.F33, _dafny.IntOfInt64(541)))
			_ = _rhs238
			var _rhs239 bool = (true) && ((func() bool {
				if (_1491_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1491_v55), 0))).Int()).(bool) {
					return _1462_v37
				}
				return _1462_v37
			})())
			_ = _rhs239
			var _rhs240 bool = !(_1493_v56).Equals(_1493_v56)
			_ = _rhs240
			var _rhs241 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm0(Companion_Default___.SafeDivisionInt((_this).F24(), _dafny.IntOfUint32((_1463_v38).Cardinality())), _this.F27, globalState))
			_ = _rhs241
			var _rhs242 bool = _1462_v37
			_ = _rhs242
			var _lhs191 *GlobalState = globalState
			_ = _lhs191
			var _lhs192 *GlobalState = globalState
			_ = _lhs192
			var _lhs193 *C1 = _1484_v48
			_ = _lhs193
			var _lhs194 _dafny.Array = _1491_v55
			_ = _lhs194
			var _lhs195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1491_v55), 0))
			_ = _lhs195
			_lhs191.F14 = _rhs238
			r0 = _rhs239
			_lhs192.F2 = _rhs240
			_lhs193.F33 = _rhs241
			(_lhs194).ArraySet1(_rhs242, (_lhs195).Int())
		} else {
			var _1494___mcc_h3 D12 = _source15.Get_().(D12_DC34).Cf56
			_ = _1494___mcc_h3
			var _1495_cf56 D12 = _1494___mcc_h3
			_ = _1495_cf56
			var _1496_v57 _dafny.Array
			_ = _1496_v57
			var _nw224 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw224
			_1496_v57 = _nw224
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1496_v57), 0))
			_ = _index203
			(_1496_v57).ArraySet1(false, (_index203).Int())
			var _1497_v58 _dafny.Sequence
			_ = _1497_v58
			_1497_v58 = _dafny.UnicodeSeqOfUtf8Bytes("kltkdmnl")
			var _1498_v59 _dafny.CodePoint
			_ = _1498_v59
			_1498_v59 = _dafny.CodePoint('h')
			var _1499_v60 D8
			_ = _1499_v60
			_1499_v60 = Companion_D8_.Create_DC25_(_this.F27, _1462_v37, _1462_v37)
			var _1500_v61 _dafny.Map
			_ = _1500_v61
			_1500_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1498_v59, (_1499_v60).Dtor_cf40())
			var _1501_v62 _dafny.Sequence
			_ = _1501_v62
			_1501_v62 = _dafny.SeqOf(_1462_v37)
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1496_v57), 0))
			_ = _index204
			var _rhs243 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg93 _dafny.Int) interface{} {
					return coer93(arg93)
				}
			}(func(_1502_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), _1497_v58), _1497_v58)
			_ = _rhs243
			var _rhs244 bool = (func() bool {
				if (_1500_v61).Contains(_1498_v59) {
					return (_1500_v61).Get(_1498_v59).(bool)
				}
				return _1462_v37
			})()
			_ = _rhs244
			var _rhs245 _dafny.Sequence = _1501_v62
			_ = _rhs245
			var _lhs196 _dafny.Array = _1496_v57
			_ = _lhs196
			var _lhs197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1496_v57), 0))
			_ = _lhs197
			var _lhs198 *GlobalState = globalState
			_ = _lhs198
			var _lhs199 *GlobalState = globalState
			_ = _lhs199
			(_lhs196).ArraySet1(_rhs243, (_lhs197).Int())
			_lhs198.F2 = _rhs244
			_lhs199.F20 = _rhs245
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1496_v57), 0))
			_ = _index205
			(_1496_v57).ArraySet1((_1496_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1496_v57), 0))).Int()).(bool), (_index205).Int())
			var _1503_v63 _dafny.MultiSet
			_ = _1503_v63
			_1503_v63 = _dafny.MultiSetOf(_1462_v37)
			var _1504_v64 _dafny.Map
			_ = _1504_v64
			_1504_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1496_v57, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ekhhngqe")).Cardinality()))
			var _1505_v65 _dafny.Sequence
			_ = _1505_v65
			_1505_v65 = _dafny.SeqOf(_1496_v57, _1496_v57, _1496_v57, _1496_v57, _1496_v57)
			var _1506_v66 _dafny.Array
			_ = _1506_v66
			var _nwElement0_43 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-453), _this.F27)
			_ = _nwElement0_43
			var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(24))
			_ = _nw225
			(_nw225).ArraySet1(_nwElement0_43, 0)
			(_nw225).ArraySet1((_dafny.Zero).Minus(_this.F27), 1)
			(_nw225).ArraySet1(Companion_Default___.Fm0((_dafny.Zero).Minus((_1503_v63).Cardinality()), _dafny.IntOfUint32((_1497_v58).Cardinality()), globalState), 2)
			(_nw225).ArraySet1(_this.F27, 3)
			(_nw225).ArraySet1(_dafny.IntOfInt64(685), 4)
			(_nw225).ArraySet1((_1504_v64).Cardinality(), 5)
			(_nw225).ArraySet1((_dafny.IntOfUint32((_1497_v58).Cardinality())).Plus(_this.F27), 6)
			(_nw225).ArraySet1(_this.F27, 7)
			(_nw225).ArraySet1((_dafny.IntOfUint32((_1463_v38).Cardinality())).Times((_this).F24()), 8)
			(_nw225).ArraySet1(_this.F27, 9)
			(_nw225).ArraySet1(_this.F27, 10)
			(_nw225).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F27, _dafny.IntOfInt64(-1)), 11)
			(_nw225).ArraySet1(_this.F27, 12)
			(_nw225).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1497_v58, _1497_v58)).Cardinality()), 13)
			(_nw225).ArraySet1(_this.F27, 14)
			(_nw225).ArraySet1(_this.F27, 15)
			(_nw225).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1505_v65, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1501_v62).Cardinality()), _dafny.IntOfUint32((_1505_v65).Cardinality()))).Uint32(), _1496_v57)).Cardinality()), 16)
			(_nw225).ArraySet1(_this.F27, 17)
			(_nw225).ArraySet1((_this).F24(), 18)
			(_nw225).ArraySet1(_this.F27, 19)
			(_nw225).ArraySet1(_dafny.IntOfUint32((_1497_v58).Cardinality()), 20)
			(_nw225).ArraySet1(_dafny.IntOfInt64(400), 21)
			(_nw225).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_1463_v38).Cardinality())).Minus((_this).F24())), 22)
			(_nw225).ArraySet1((func() _dafny.Int {
				if _1462_v37 {
					return _dafny.IntOfInt64(662)
				}
				return Companion_Default___.Fm0(_dafny.IntOfInt64(-584), (_this).F24(), globalState)
			})(), 23)
			_1506_v66 = _nw225
			var _1507_v67 _dafny.Sequence
			_ = _1507_v67
			_1507_v67 = _dafny.SeqOf(_1506_v66, _1506_v66, (func() _dafny.Array {
				if _1462_v37 {
					return _1506_v66
				}
				return _1506_v66
			})())
			_1506_v66 = (_1507_v67).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1507_v67).Cardinality()))).Uint32()).(_dafny.Array)
			(globalState).F2 = _1462_v37
		}
		var _1508_v68 _dafny.MultiSet
		_ = _1508_v68
		_1508_v68 = _dafny.MultiSetOf(!(_1462_v37))
		var _1509_v69 _dafny.Sequence
		_ = _1509_v69
		_1509_v69 = _dafny.SeqOf(_1462_v37, _1462_v37, _1462_v37)
		var _1510_v70 _dafny.MultiSet
		_ = _1510_v70
		_1510_v70 = _dafny.MultiSetOf((_this).F24(), _this.F27, (_this).F24(), _this.F27, _this.F27)
		var _1511_v71 *C1
		_ = _1511_v71
		var _nw226 *C1 = New_C1_()
		_ = _nw226
		_nw226.Ctor__(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_1509_v69).Cardinality())), _1510_v70), (_this).F24(), _this.F27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(382))).Uint32(), func(coer94 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg94 _dafny.Int) interface{} {
				return coer94(arg94)
			}
		}(func(_1512_i8 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), _this.F27)
		})))
		_1511_v71 = _nw226
		var _1513_v72 _dafny.Set
		_ = _1513_v72
		_1513_v72 = _dafny.SetOf(_1511_v71, _1511_v71)
		var _1514_v73 _dafny.Array
		_ = _1514_v73
		var _nw227 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw227
		_1514_v73 = _nw227
		var _1515_v75 _dafny.CodePoint
		_ = _1515_v75
		_1515_v75 = _dafny.CodePoint('k')
		var _1516_v76 _dafny.Map
		_ = _1516_v76
		_1516_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1515_v75, _dafny.UnicodeSeqOfUtf8Bytes("uoejhqk"))
		var _1517_v77 T0
		_ = _1517_v77
		var _nw228 *C4 = New_C4_()
		_ = _nw228
		_nw228.Ctor__(_1514_v73, _this.F27, _dafny.SeqOf(func() _dafny.Map {
			var _coll42 = _dafny.NewMapBuilder()
			_ = _coll42
			for _iter46 := _dafny.Iterate((_1516_v76).Keys().Elements()); ; {
				_compr_42, _ok46 := _iter46()
				if !_ok46 {
					break
				}
				var _1518_v74 _dafny.CodePoint
				_1518_v74 = interface{}(_compr_42).(_dafny.CodePoint)
				if (_1516_v76).Contains(_1518_v74) {
					_coll42.Add(_1518_v74, _1511_v71.F33)
				}
			}
			return _coll42.ToMap()
		}()))
		_1517_v77 = _nw228
		var _1519_v78 _dafny.Set
		_ = _1519_v78
		_1519_v78 = _dafny.SetOf(_1517_v77)
		var _1520_v79 _dafny.Map
		_ = _1520_v79
		_1520_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1462_v37)
		var _1521_v80 _dafny.Sequence
		_ = _1521_v80
		_1521_v80 = _dafny.UnicodeSeqOfUtf8Bytes("uuvexp")
		var _1522_v81 _dafny.Array
		_ = _1522_v81
		var _nwElement0_44 _dafny.Int = (_dafny.Zero).Minus((_this).F24())
		_ = _nwElement0_44
		var _nw229 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(23))
		_ = _nw229
		(_nw229).ArraySet1(_nwElement0_44, 0)
		(_nw229).ArraySet1((_this).F24(), 1)
		(_nw229).ArraySet1(_this.F27, 2)
		(_nw229).ArraySet1(_this.F27, 3)
		(_nw229).ArraySet1(_this.F27, 4)
		(_nw229).ArraySet1((_this).F24(), 5)
		(_nw229).ArraySet1((_this).F24(), 6)
		(_nw229).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bobuyv")).Cardinality())), 7)
		(_nw229).ArraySet1((_this).F24(), 8)
		(_nw229).ArraySet1((_dafny.Zero).Minus((_this).F24()), 9)
		(_nw229).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()), 10)
		(_nw229).ArraySet1(_dafny.IntOfInt64(-676), 11)
		(_nw229).ArraySet1(Companion_Default___.Fm0((_this).F24(), _this.F27, globalState), 12)
		(_nw229).ArraySet1(Companion_Default___.Fm0((_1513_v72).Cardinality(), _this.F27, globalState), 13)
		(_nw229).ArraySet1(_this.F27, 14)
		(_nw229).ArraySet1((_1519_v78).Cardinality(), 15)
		(_nw229).ArraySet1((_1520_v79).Cardinality(), 16)
		(_nw229).ArraySet1((_this).F24(), 17)
		(_nw229).ArraySet1((_this).F24(), 18)
		(_nw229).ArraySet1((_this).F24(), 19)
		(_nw229).ArraySet1(_1511_v71.F33, 20)
		(_nw229).ArraySet1((_this).F24(), 21)
		(_nw229).ArraySet1(_dafny.IntOfUint32((_1521_v80).Cardinality()), 22)
		_1522_v81 = _nw229
		var _1523_v82 _dafny.Map
		_ = _1523_v82
		_1523_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_1508_v68).Cardinality(), _dafny.IntOfInt64(461)), _1522_v81)
		var _1524_v83 _dafny.Array
		_ = _1524_v83
		var _nw230 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
		_ = _nw230
		_1524_v83 = _nw230
		var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_1524_v83), 0))
		_ = _index206
		(_1524_v83).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg95 _dafny.Int) interface{} {
				return coer95(arg95)
			}
		}(func(_1525_i9 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), (_index206).Int())
		var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_1524_v83), 0))
		_ = _index207
		var _rhs246 bool = false
		_ = _rhs246
		var _rhs247 bool = (func() bool {
			if ((_this).F26()).Contains(_1462_v37) {
				return ((_this).F26()).Get(_1462_v37).(bool)
			}
			return _dafny.Companion_Sequence_.IsProperPrefixOf(_1509_v69, _1509_v69)
		})()
		_ = _rhs247
		var _rhs248 _dafny.Map = (_1523_v82).Update(_dafny.IntOfUint32((_1509_v69).Cardinality()), _1514_v73)
		_ = _rhs248
		var _rhs249 bool = _1462_v37
		_ = _rhs249
		var _rhs250 _dafny.Sequence = _1521_v80
		_ = _rhs250
		var _lhs200 *GlobalState = globalState
		_ = _lhs200
		var _lhs201 _dafny.Array = _1524_v83
		_ = _lhs201
		var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_1524_v83), 0))
		_ = _lhs202
		_lhs200.F2 = _rhs246
		_1462_v37 = _rhs247
		_1523_v82 = _rhs248
		r0 = _rhs249
		(_lhs201).ArraySet1(_rhs250, (_lhs202).Int())
		_1514_v73 = _1522_v81
		var _1526_v84 _dafny.Map
		_ = _1526_v84
		_1526_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1511_v71.F33)
		var _1527_v85 D4
		_ = _1527_v85
		_1527_v85 = Companion_D4_.Create_DC13_(_this.F27, _1526_v84, _1521_v80)
		var _1528_v86 D4
		_ = _1528_v86
		_1528_v86 = Companion_D4_.Create_DC14_(_1527_v85)
		var _1529_v87 D4
		_ = _1529_v87
		_1529_v87 = Companion_D4_.Create_DC14_(_1528_v86)
		var _1530_v88 D4
		_ = _1530_v88
		_1530_v88 = Companion_D4_.Create_DC14_(_1528_v86)
		var _1531_v89 _dafny.Map
		_ = _1531_v89
		_1531_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_v88, _1521_v80)
		var _1532_v90 D7
		_ = _1532_v90
		_1532_v90 = Companion_D7_.Create_DC20_(_1531_v89)
		var _1533_v91 _dafny.Map
		_ = _1533_v91
		_1533_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1532_v90, (_1524_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_1524_v83), 0))).Int()).(_dafny.Sequence))
		_1533_v91 = (_1533_v91).Update(_1532_v90, _dafny.UnicodeSeqOfUtf8Bytes("aohh"))
		var _1534_v92 _dafny.Array
		_ = _1534_v92
		var _nw231 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
		_ = _nw231
		_1534_v92 = _nw231
		var _1535_v93 *C4
		_ = _1535_v93
		var _nw232 *C4 = New_C4_()
		_ = _nw232
		_nw232.Ctor__(_1522_v81, (_1517_v77).F24(), (_this).F25())
		_1535_v93 = _nw232
		var _1536_v94 _dafny.Sequence
		_ = _1536_v94
		_1536_v94 = _dafny.SeqOf(_1535_v93)
		var _1537_v95 _dafny.Map
		_ = _1537_v95
		_1537_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1462_v37, _1509_v69)
		var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1534_v92), 0))
		_ = _index208
		(_1534_v92).ArraySet1((_1536_v94).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1537_v95).Contains(_1462_v37) {
				return (_1537_v95).Get(_1462_v37).(_dafny.Sequence)
			}
			return _1509_v69
		})()).Cardinality()), _dafny.IntOfUint32((_1536_v94).Cardinality()))).Uint32()).(*C4), (_index208).Int())
		var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1534_v92), 0))
		_ = _index209
		(_1534_v92).ArraySet1((_1536_v94).Select((Companion_Default___.SafeIndex(((_this).F24()).Times(_dafny.IntOfInt64(492)), _dafny.IntOfUint32((_1536_v94).Cardinality()))).Uint32()).(*C4), (_index209).Int())
		r0 = ((_1511_v71).Fm10(globalState)) || (!(_1462_v37))
		return r0
	}
}
func (_this *C5) M2(globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1538_v0 bool
		_ = _1538_v0
		_1538_v0 = true
		if !(_1538_v0) {
			(globalState).F9 = (_this).F24()
			var _1539_v1 _dafny.Map
			_ = _1539_v1
			_1539_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1538_v0, ((_this).F26()).Cardinality())
			var _1540_v2 D4
			_ = _1540_v2
			_1540_v2 = Companion_D4_.Create_DC13_((_this).F24(), _1539_v1, _dafny.UnicodeSeqOfUtf8Bytes("qiphwxsf"))
			if ((_1540_v2).Dtor_cf22()).Cmp(_this.F27) == 0 {
				var _1541_v3 _dafny.Array
				_ = _1541_v3
				var _nw233 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw233
				_1541_v3 = _nw233
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1541_v3), 0))
				_ = _index210
				(_1541_v3).ArraySet1(_1538_v0, (_index210).Int())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1541_v3), 0))
				_ = _index211
				(_1541_v3).ArraySet1(_1538_v0, (_index211).Int())
				var _1542_v4 _dafny.CodePoint
				_ = _1542_v4
				_1542_v4 = _dafny.CodePoint('f')
				var _1543_v5 *C0
				_ = _1543_v5
				var _nw234 *C0 = New_C0_()
				_ = _nw234
				_nw234.Ctor__(_1538_v0)
				_1543_v5 = _nw234
				var _1544_v6 _dafny.Set
				_ = _1544_v6
				_1544_v6 = _dafny.SetOf(_this.F27)
				var _1545_v7 _dafny.MultiSet
				_ = _1545_v7
				_1545_v7 = _dafny.MultiSetOf((_1544_v6).Cardinality())
				var _1546_v8 _dafny.Map
				_ = _1546_v8
				_1546_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1542_v4, (_1545_v7).Cardinality())
				var _1547_v9 T0
				_ = _1547_v9
				var _nw235 *C2 = New_C2_()
				_ = _nw235
				_nw235.Ctor__((_1543_v5).F34(), (_this).F24(), _dafny.SeqOf(_1546_v8))
				_1547_v9 = _nw235
				var _1548_v10 _dafny.Array
				_ = _1548_v10
				var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw236
				_1548_v10 = _nw236
				var _1549_v11 _dafny.Map
				_ = _1549_v11
				_1549_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1547_v9, _1548_v10)
				var _rhs251 bool = (_1549_v11).Contains(_1547_v9)
				_ = _rhs251
				var _rhs252 _dafny.CodePoint = Companion_Default___.Fm3(globalState)
				_ = _rhs252
				var _rhs253 *C0 = _1543_v5
				_ = _rhs253
				var _lhs203 *GlobalState = globalState
				_ = _lhs203
				_lhs203.F2 = _rhs251
				_1542_v4 = _rhs252
				_1543_v5 = _rhs253
				(globalState).F14 = (_1547_v9).F24()
				var _1550_v12 _dafny.MultiSet
				_ = _1550_v12
				_1550_v12 = _dafny.MultiSetOf((_1541_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_1541_v3), 0))).Int()).(bool))
				var _1551_v13 _dafny.MultiSet
				_ = _1551_v13
				_1551_v13 = _dafny.MultiSetOf((_1550_v12).Difference(_1550_v12))
				_1551_v13 = (_1551_v13).Update(_1550_v12, Companion_Default___.Abs((_this).F24()))
				var _1552_v14 _dafny.Array
				_ = _1552_v14
				var _nwElement0_45 _dafny.Set = _1544_v6
				_ = _nwElement0_45
				var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(9))
				_ = _nw237
				(_nw237).ArraySet1(_nwElement0_45, 0)
				(_nw237).ArraySet1(_1544_v6, 1)
				(_nw237).ArraySet1(_1544_v6, 2)
				(_nw237).ArraySet1(_dafny.SetOf((_1547_v9).F24()), 3)
				(_nw237).ArraySet1(_1544_v6, 4)
				(_nw237).ArraySet1(_1544_v6, 5)
				(_nw237).ArraySet1(_1544_v6, 6)
				(_nw237).ArraySet1(_1544_v6, 7)
				(_nw237).ArraySet1(_dafny.SetOf((_this).F24(), (_this).F24()), 8)
				_1552_v14 = _nw237
				var _1553_v15 _dafny.Sequence
				_ = _1553_v15
				_1553_v15 = _dafny.SeqOf(_1552_v14)
				var _1554_v16 _dafny.Sequence
				_ = _1554_v16
				_1554_v16 = _dafny.SeqOf((_1543_v5).F34(), _1538_v0)
				var _1555_v17 _dafny.Array
				_ = _1555_v17
				var _nw238 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw238
				_1555_v17 = _nw238
				var _1556_v18 *C3
				_ = _1556_v18
				var _nw239 *C3 = New_C3_()
				_ = _nw239
				_nw239.Ctor__((_1553_v15).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1554_v16).Cardinality()), _dafny.IntOfUint32((_1553_v15).Cardinality()))).Uint32()).(_dafny.Array), _1555_v17, (_this).F24(), (_1547_v9).F25())
				_1556_v18 = _nw239
			} else {
				var _1557_v19 _dafny.Array
				_ = _1557_v19
				var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(2))
				_ = _nw240
				_1557_v19 = _nw240
				var _1558_v21 _dafny.Array
				_ = _1558_v21
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_36
				var _nw241 _dafny.Array
				_ = _nw241
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw241 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Map = func(_1559_i0 _dafny.Int) _dafny.Map {
						return (Companion_D11_.Create_DC30_(func() _dafny.Map {
							var _coll43 = _dafny.NewMapBuilder()
							_ = _coll43
							for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(429), _dafny.IntOfInt64(778))); ; {
								_compr_43, _ok47 := _iter47()
								if !_ok47 {
									break
								}
								var _1560_v20 _dafny.Int
								_1560_v20 = interface{}(_compr_43).(_dafny.Int)
								if ((_dafny.IntOfInt64(429)).Cmp(_1560_v20) <= 0) && ((_1560_v20).Cmp(_dafny.IntOfInt64(778)) < 0) {
									_coll43.Add((_1560_v20).Plus((_this).F24()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bgqm")).Cardinality())))
								}
							}
							return _coll43.ToMap()
						}())).Dtor_cf47()
					}
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw241 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw241).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw241).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1558_v21 = _nw241
				var _1561_v22 T0
				_ = _1561_v22
				var _nw242 *C3 = New_C3_()
				_ = _nw242
				_nw242.Ctor__(_1557_v19, _1558_v21, (_this).F24(), (_this).F25())
				_1561_v22 = _nw242
				var _1562_v23 T0
				_ = _1562_v23
				_1562_v23 = _1561_v22
				_1562_v23 = _1562_v23
				var _1563_v24 _dafny.MultiSet
				_ = _1563_v24
				_1563_v24 = _dafny.MultiSetOf(_1561_v22, _1561_v22)
				var _1564_v25 _dafny.Set
				_ = _1564_v25
				_1564_v25 = _dafny.SetOf((_1563_v24).IsDisjointFrom(_1563_v24))
				var _1565_v26 D0
				_ = _1565_v26
				_1565_v26 = Companion_D0_.Create_DC1_((_this).F24(), _dafny.IntOfInt64(279), _1538_v0)
				var _1566_v27 _dafny.CodePoint
				_ = _1566_v27
				_1566_v27 = _dafny.CodePoint('n')
				var _pat_let_tv25 = _1538_v0
				_ = _pat_let_tv25
				var _rhs254 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1561_v22).F24(), (_1561_v22).F24()))
				_ = _rhs254
				var _rhs255 _dafny.Set = _dafny.SetOf((func(_pat_let29_0 D0) D0 {
					return func(_1567_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let30_0 _dafny.Int) D0 {
							return func(_1568_dt__update_hcf1_h0 _dafny.Int) D0 {
								return func(_pat_let31_0 bool) D0 {
									return func(_1569_dt__update_hcf3_h0 bool) D0 {
										return Companion_D0_.Create_DC1_(_1568_dt__update_hcf1_h0, (_1567_dt__update__tmp_h0).Dtor_cf2(), _1569_dt__update_hcf3_h0)
									}(_pat_let31_0)
								}(_pat_let_tv25)
							}(_pat_let30_0)
						}(_this.F27)
					}(_pat_let29_0)
				}(_1565_v26)).Dtor_cf3(), Companion_Default___.Fm2(_this.F27, Companion_Default___.Fm0(_dafny.IntOfInt64(-397), (_dafny.Zero).Minus((_1561_v22).F24()), globalState), _1566_v27, globalState))
				_ = _rhs255
				var _rhs256 _dafny.Int = (_this).F24()
				_ = _rhs256
				var _lhs204 *GlobalState = globalState
				_ = _lhs204
				var _lhs205 *GlobalState = globalState
				_ = _lhs205
				_lhs204.F14 = _rhs254
				_1564_v25 = _rhs255
				_lhs205.F18 = _rhs256
				(globalState).F2 = (func() bool {
					if _1538_v0 {
						return _1538_v0
					}
					return !(_1538_v0)
				})()
				var _1570_v28 _dafny.Array
				_ = _1570_v28
				var _nw243 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw243
				_1570_v28 = _nw243
				_1570_v28 = _1570_v28
				var _1571_v29 _dafny.MultiSet
				_ = _1571_v29
				_1571_v29 = _dafny.MultiSetOf(_1538_v0, _1538_v0)
				var _1572_v30 _dafny.Sequence
				_ = _1572_v30
				_1572_v30 = _dafny.SeqOf(_this.F27)
				var _1573_v31 _dafny.Set
				_ = _1573_v31
				_1573_v31 = _dafny.SetOf((_this).F24(), _dafny.IntOfUint32((_1572_v30).Cardinality()), (_this).F24(), _this.F27, _this.F27)
				var _1574_v32 _dafny.Set
				_ = _1574_v32
				_1574_v32 = _dafny.SetOf(_1566_v27, _1566_v27)
				var _1575_v33 _dafny.Array
				_ = _1575_v33
				var _nwElement0_46 bool = !(_1538_v0) || (_1538_v0)
				_ = _nwElement0_46
				var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(29))
				_ = _nw244
				(_nw244).ArraySet1(_nwElement0_46, 0)
				(_nw244).ArraySet1(_1538_v0, 1)
				(_nw244).ArraySet1(_1538_v0, 2)
				(_nw244).ArraySet1(((_this).F24()).Cmp((_1561_v22).F24()) >= 0, 3)
				(_nw244).ArraySet1(_1538_v0, 4)
				(_nw244).ArraySet1(_1538_v0, 5)
				(_nw244).ArraySet1(false, 6)
				(_nw244).ArraySet1(_1538_v0, 7)
				(_nw244).ArraySet1(false, 8)
				(_nw244).ArraySet1(_1538_v0, 9)
				(_nw244).ArraySet1(_1538_v0, 10)
				(_nw244).ArraySet1(((_1561_v22).F24()).Cmp(_this.F27) >= 0, 11)
				(_nw244).ArraySet1((_1565_v26).Dtor_cf3(), 12)
				(_nw244).ArraySet1(true, 13)
				(_nw244).ArraySet1(_1538_v0, 14)
				(_nw244).ArraySet1((_this.F27).Cmp((_1571_v29).Cardinality()) == 0, 15)
				(_nw244).ArraySet1(((_1561_v22).F24()).Cmp((_1561_v22).F24()) < 0, 16)
				(_nw244).ArraySet1(_1538_v0, 17)
				(_nw244).ArraySet1(_1538_v0, 18)
				(_nw244).ArraySet1(_1538_v0, 19)
				(_nw244).ArraySet1((_1573_v31).IsSubsetOf(_1573_v31), 20)
				(_nw244).ArraySet1(_1538_v0, 21)
				(_nw244).ArraySet1(_1538_v0, 22)
				(_nw244).ArraySet1(!(((_this).F24()).Cmp((_1574_v32).Cardinality()) < 0), 23)
				(_nw244).ArraySet1(!(_1538_v0), 24)
				(_nw244).ArraySet1((_1573_v31).IsProperSubsetOf(_dafny.SetOf(_this.F27, _this.F27)), 25)
				(_nw244).ArraySet1((func() bool {
					if _1538_v0 {
						return _1538_v0
					}
					return _1538_v0
				})(), 26)
				(_nw244).ArraySet1(((_this).F24()).Cmp(_this.F27) == 0, 27)
				(_nw244).ArraySet1(_1538_v0, 28)
				_1575_v33 = _nw244
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_1575_v33), 0))
				_ = _index212
				(_1575_v33).ArraySet1(_1538_v0, (_index212).Int())
				var _1576_v34 _dafny.Sequence
				_ = _1576_v34
				_1576_v34 = _dafny.SeqOf(_1566_v27)
				var _1577_v35 _dafny.MultiSet
				_ = _1577_v35
				_1577_v35 = _dafny.MultiSetOf((_this).F24(), _dafny.IntOfUint32((_1576_v34).Cardinality()))
				var _1578_v36 *C0
				_ = _1578_v36
				var _nw245 *C0 = New_C0_()
				_ = _nw245
				_nw245.Ctor__(_1538_v0)
				_1578_v36 = _nw245
				var _1579_v37 _dafny.Map
				_ = _1579_v37
				_1579_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1578_v36, _dafny.MultiSetFromSeq(_1572_v30))
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_1575_v33), 0))
				_ = _index213
				(_1575_v33).ArraySet1(((_1577_v35).Intersection((func() _dafny.MultiSet {
					if (_1579_v37).Contains(_1578_v36) {
						return (_1579_v37).Get(_1578_v36).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F24()))
				})())).IsSubsetOf(_1577_v35), (_index213).Int())
			}
			(globalState).F18 = Companion_Default___.SafeDivisionInt((_this).F24(), _dafny.IntOfInt64(798))
			var _1580_v38 _dafny.CodePoint
			_ = _1580_v38
			_1580_v38 = _dafny.CodePoint('j')
			var _1581_v39 _dafny.Sequence
			_ = _1581_v39
			_1581_v39 = _dafny.UnicodeSeqOfUtf8Bytes("dhvqnlf")
			_1538_v0 = (func() bool {
				if _1538_v0 {
					return _dafny.Companion_Sequence_.Contains(_1581_v39, _1580_v38)
				}
				return !((func() bool {
					if _1538_v0 {
						return _1538_v0
					}
					return _1538_v0
				})())
			})()
			(globalState).F17 = _1581_v39
		} else {
			(globalState).F1 = (_dafny.Zero).Minus((_this).F24())
			var _1582_v40 _dafny.CodePoint
			_ = _1582_v40
			_1582_v40 = _dafny.CodePoint('t')
			if Companion_Default___.Fm2((_this).F24(), (_this).F24(), _1582_v40, globalState) {
				(globalState).F14 = (_this).F24()
				var _1583_v41 D12
				_ = _1583_v41
				_1583_v41 = Companion_D12_.Create_DC33_(_1582_v40, _1538_v0)
				var _1584_v42 D12
				_ = _1584_v42
				_1584_v42 = Companion_D12_.Create_DC34_((func() D12 {
					if !(_1538_v0) {
						return _1583_v41
					}
					return _1583_v41
				})())
				var _1585_v43 _dafny.Array
				_ = _1585_v43
				var _nw246 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw246
				_1585_v43 = _nw246
				var _1586_v44 _dafny.Sequence
				_ = _1586_v44
				_1586_v44 = _dafny.SeqOf(_1585_v43)
				var _1587_v45 T0
				_ = _1587_v45
				var _nw247 *C4 = New_C4_()
				_ = _nw247
				_nw247.Ctor__((_1586_v44).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1586_v44).Cardinality()))).Uint32()).(_dafny.Array), _dafny.IntOfInt64(191), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(181))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}(func(_1588_i1 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _this.F27)
				})))
				_1587_v45 = _nw247
				var _1589_v46 _dafny.Sequence
				_ = _1589_v46
				_1589_v46 = _dafny.SeqOf(_1538_v0, _1538_v0, _1538_v0)
				var _1590_v47 _dafny.Map
				_ = _1590_v47
				_1590_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(972))).Uint32(), func(coer97 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}((func(_1591_v45 T0, _1592_v0 bool) func(_dafny.Int) _dafny.Int {
					return func(_1593_i2 _dafny.Int) _dafny.Int {
						return (Companion_D8_.Create_DC25_((_1591_v45).F24(), _1592_v0, _1592_v0)).Dtor_cf38()
					}
				})(_1587_v45, _1538_v0))), _1587_v45)
				var _1594_v48 _dafny.Map
				_ = _1594_v48
				_1594_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_1587_v45).F24(), (_this).F24())).Cardinality(), (_1587_v45).F24())
				var _1595_v49 _dafny.Sequence
				_ = _1595_v49
				_1595_v49 = _dafny.SeqOf((_1594_v48).Cardinality(), Companion_Default___.Fm0((_1587_v45).F24(), _this.F27, globalState))
				var _rhs257 D12 = _1584_v42
				_ = _rhs257
				var _rhs258 bool = _dafny.Companion_Sequence_.IsPrefixOf(_1589_v46, _1589_v46)
				_ = _rhs258
				var _rhs259 T0 = (func() T0 {
					if (_1590_v47).Contains(_1595_v49) {
						return (_1590_v47).Get(_1595_v49).(T0)
					}
					return _1587_v45
				})()
				_ = _rhs259
				_1584_v42 = _rhs257
				_1538_v0 = _rhs258
				_1587_v45 = _rhs259
				var _1596_v50 *C4
				_ = _1596_v50
				var _nw248 *C4 = New_C4_()
				_ = _nw248
				_nw248.Ctor__(_1585_v43, (_1587_v45).F24(), (_1587_v45).F25())
				_1596_v50 = _nw248
				var _1597_v51 _dafny.Map
				_ = _1597_v51
				_1597_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1587_v45).F24(), _1596_v50)
				var _1598_v52 _dafny.Map
				_ = _1598_v52
				_1598_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_1597_v51).Cardinality())
				var _1599_v53 _dafny.MultiSet
				_ = _1599_v53
				_1599_v53 = _dafny.MultiSetOf(_1538_v0)
				var _1600_v54 _dafny.Sequence
				_ = _1600_v54
				_1600_v54 = _dafny.SeqOf((_1598_v52).Update(_1538_v0, (_1599_v53).Cardinality()), _1598_v52)
				_1600_v54 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_1595_v49).Cardinality()))).Update(_1538_v0, _this.F27), _1598_v52)
				r0 = _1538_v0
				(globalState).F2 = (_1538_v0) == (_1538_v0)
			} else {
				var _1601_v55 _dafny.Sequence
				_ = _1601_v55
				_1601_v55 = _dafny.UnicodeSeqOfUtf8Bytes("drbtc")
				var _1602_v56 _dafny.Set
				_ = _1602_v56
				_1602_v56 = _dafny.SetOf(_dafny.IntOfUint32((_1601_v55).Cardinality()))
				var _1603_v57 _dafny.Sequence
				_ = _1603_v57
				_1603_v57 = _dafny.SeqOf(_1538_v0, _1538_v0)
				var _1604_v58 _dafny.MultiSet
				_ = _1604_v58
				_1604_v58 = _dafny.MultiSetOf((_this).F24(), _dafny.IntOfUint32((_1601_v55).Cardinality()))
				var _1605_v60 D6
				_ = _1605_v60
				_1605_v60 = Companion_D6_.Create_DC16_(_1602_v56)
				var _1606_v61 _dafny.Array
				_ = _1606_v61
				var _nwElement0_47 _dafny.Set = _1602_v56
				_ = _nwElement0_47
				var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(28))
				_ = _nw249
				(_nw249).ArraySet1(_nwElement0_47, 0)
				(_nw249).ArraySet1(_1602_v56, 1)
				(_nw249).ArraySet1(_1602_v56, 2)
				(_nw249).ArraySet1(_dafny.SetOf(_this.F27, _this.F27, _dafny.IntOfUint32((_1603_v57).Cardinality())), 3)
				(_nw249).ArraySet1(_1602_v56, 4)
				(_nw249).ArraySet1(_1602_v56, 5)
				(_nw249).ArraySet1(_1602_v56, 6)
				(_nw249).ArraySet1(_1602_v56, 7)
				(_nw249).ArraySet1(func() _dafny.Set {
					var _coll44 = _dafny.NewBuilder()
					_ = _coll44
					for _iter48 := _dafny.Iterate((_1604_v58).Elements()); ; {
						_compr_44, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1607_v59 _dafny.Int
						_1607_v59 = interface{}(_compr_44).(_dafny.Int)
						if (_1604_v58).Contains(_1607_v59) {
							_coll44.Add((_1607_v59).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("omdiun")).Cardinality())))
						}
					}
					return _coll44.ToSet()
				}(), 8)
				(_nw249).ArraySet1(_1602_v56, 9)
				(_nw249).ArraySet1(_1602_v56, 10)
				(_nw249).ArraySet1((_1605_v60).Dtor_cf27(), 11)
				(_nw249).ArraySet1(_1602_v56, 12)
				(_nw249).ArraySet1(_1602_v56, 13)
				(_nw249).ArraySet1(_1602_v56, 14)
				(_nw249).ArraySet1(_1602_v56, 15)
				(_nw249).ArraySet1(_1602_v56, 16)
				(_nw249).ArraySet1(_1602_v56, 17)
				(_nw249).ArraySet1(_1602_v56, 18)
				(_nw249).ArraySet1(_1602_v56, 19)
				(_nw249).ArraySet1(_1602_v56, 20)
				(_nw249).ArraySet1(_dafny.SetOf(_dafny.IntOfInt64(510), (_1602_v56).Cardinality()), 21)
				(_nw249).ArraySet1(_1602_v56, 22)
				(_nw249).ArraySet1(_1602_v56, 23)
				(_nw249).ArraySet1(_1602_v56, 24)
				(_nw249).ArraySet1(_1602_v56, 25)
				(_nw249).ArraySet1(_1602_v56, 26)
				(_nw249).ArraySet1(_1602_v56, 27)
				_1606_v61 = _nw249
				var _1608_v62 _dafny.Array
				_ = _1608_v62
				var _nw250 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
				_ = _nw250
				_1608_v62 = _nw250
				var _1609_v63 _dafny.Sequence
				_ = _1609_v63
				_1609_v63 = (_this).F25()
				var _1610_v64 *C3
				_ = _1610_v64
				var _nw251 *C3 = New_C3_()
				_ = _nw251
				_nw251.Ctor__(_1606_v61, _1608_v62, (_dafny.Zero).Minus((_this).F24()), (_1609_v63))
				_1610_v64 = _nw251
				var _1611_v65 _dafny.Map
				_ = _1611_v65
				_1611_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1538_v0, _1610_v64)
				_1611_v65 = (_1611_v65).Update(false, _1610_v64)
				var _1612_v66 _dafny.Map
				_ = _1612_v66
				_1612_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1538_v0, _this.F27)
				var _1613_v67 _dafny.Map
				_ = _1613_v67
				_1613_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1538_v0), _1612_v66)
				_1613_v67 = ((_1613_v67).Merge(_1613_v67)).Merge(_1613_v67)
				_1612_v66 = Companion_Default___.Fm17(Companion_Default___.Fm0(_dafny.IntOfUint32((_1601_v55).Cardinality()), _dafny.IntOfInt64(-163), globalState), _this.F27, _this.F27, _1538_v0, globalState)
				var _1614_v68 _dafny.Array
				_ = _1614_v68
				var _nw252 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
				_ = _nw252
				_1614_v68 = _nw252
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1614_v68), 0))
				_ = _index214
				(_1614_v68).ArraySet1CodePoint(_1582_v40, (_index214).Int())
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_1614_v68), 0))
				_ = _index215
				(_1614_v68).ArraySet1CodePoint(_1582_v40, (_index215).Int())
				var _1615_v69 _dafny.Map
				_ = _1615_v69
				_1615_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_1603_v57).Cardinality())), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()), _dafny.IntOfUint32((_1601_v55).Cardinality()))))
				var _1616_v70 _dafny.Sequence
				_ = _1616_v70
				_1616_v70 = _dafny.SeqOf(_this.F27)
				var _1617_v71 _dafny.MultiSet
				_ = _1617_v71
				_1617_v71 = _dafny.MultiSetOf(_1538_v0)
				_1615_v69 = (_1615_v69).Update(_dafny.Companion_Sequence_.Concatenate(_1616_v70, _1616_v70), (_this.F27).Minus((_1617_v71).Cardinality()))
			}
			var _1618_v72 _dafny.Set
			_ = _1618_v72
			_1618_v72 = _dafny.SetOf(_this.F27)
			var _1619_v73 _dafny.Sequence
			_ = _1619_v73
			_1619_v73 = _dafny.SeqOf((_this).F24(), _dafny.IntOfInt64(-217))
			var _1620_v75 _dafny.Array
			_ = _1620_v75
			var _nwElement0_48 _dafny.Set = (_1618_v72).Difference(_1618_v72)
			_ = _nwElement0_48
			var _nw253 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(14))
			_ = _nw253
			(_nw253).ArraySet1(_nwElement0_48, 0)
			(_nw253).ArraySet1(_1618_v72, 1)
			(_nw253).ArraySet1(_dafny.SetOf(_this.F27, (_this).F24(), (_this).F24()), 2)
			(_nw253).ArraySet1((_dafny.SetOf(_this.F27, _this.F27)).Union(_1618_v72), 3)
			(_nw253).ArraySet1(_1618_v72, 4)
			(_nw253).ArraySet1(_1618_v72, 5)
			(_nw253).ArraySet1(_1618_v72, 6)
			(_nw253).ArraySet1(_1618_v72, 7)
			(_nw253).ArraySet1(_1618_v72, 8)
			(_nw253).ArraySet1(_1618_v72, 9)
			(_nw253).ArraySet1(_dafny.SetOf(_this.F27, _this.F27, (_this).F24(), _dafny.IntOfUint32((_1619_v73).Cardinality())), 10)
			(_nw253).ArraySet1(func() _dafny.Set {
				var _coll45 = _dafny.NewBuilder()
				_ = _coll45
				for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(234), _dafny.IntOfInt64(-416))); ; {
					_compr_45, _ok49 := _iter49()
					if !_ok49 {
						break
					}
					var _1621_v74 _dafny.Int
					_1621_v74 = interface{}(_compr_45).(_dafny.Int)
					if ((_dafny.IntOfInt64(234)).Cmp(_1621_v74) <= 0) && ((_1621_v74).Cmp(_dafny.IntOfInt64(-416)) < 0) {
						_coll45.Add(Companion_Default___.SafeModuloInt(_1621_v74, (_this).F24()))
					}
				}
				return _coll45.ToSet()
			}(), 11)
			(_nw253).ArraySet1(_1618_v72, 12)
			(_nw253).ArraySet1(_dafny.SetOf((_this).F24(), (_this).F24()), 13)
			_1620_v75 = _nw253
			_1620_v75 = _1620_v75
			var _1622_v76 _dafny.Array
			_ = _1622_v76
			var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
			_ = _nw254
			_1622_v76 = _nw254
			var _1623_v77 _dafny.Array
			_ = _1623_v77
			var _nw255 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw255
			_1623_v77 = _nw255
			var _1624_v78 _dafny.Map
			_ = _1624_v78
			_1624_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1623_v77)
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1622_v76), 0))
			_ = _index216
			(_1622_v76).ArraySet1((_1624_v78).Merge(_1624_v78), (_index216).Int())
			var _1625_v79 _dafny.Map
			_ = _1625_v79
			_1625_v79 = _1624_v78
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1622_v76), 0))
			_ = _index217
			(_1622_v76).ArraySet1((_1625_v79), (_index217).Int())
			(globalState).F18 = _dafny.IntOfInt64(772)
		}
		(globalState).F1 = _dafny.IntOfUint32((Companion_Default___.Fm40(globalState)).Cardinality())
		(_this).M3((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F24(), (_this).F24())), globalState)
		var _1626_v80 _dafny.MultiSet
		_ = _1626_v80
		_1626_v80 = _dafny.MultiSetOf(_this.F27)
		if (_1626_v80).IsProperSubsetOf(_1626_v80) {
			var _1627_v81 _dafny.Sequence
			_ = _1627_v81
			_1627_v81 = _dafny.SeqOf(_1538_v0)
			if (func() bool {
				if _1538_v0 {
					return !(_1538_v0) || (_1538_v0)
				}
				return (_1627_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-585), _dafny.IntOfUint32((_1627_v81).Cardinality()))).Uint32()).(bool)
			})() {
				(globalState).F18 = _this.F27
				r0 = !((func() bool {
					if (_this.F27).Cmp((_this).F24()) != 0 {
						return _1538_v0
					}
					return !(_1538_v0)
				})())
				var _1628_v82 _dafny.Sequence
				_ = _1628_v82
				_1628_v82 = _dafny.UnicodeSeqOfUtf8Bytes("lofb")
				_1538_v0 = !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("th"), _1628_v82)
				var _1629_v83 _dafny.Set
				_ = _1629_v83
				_1629_v83 = _dafny.SetOf(_1538_v0)
				_1629_v83 = (_1629_v83).Union((_1629_v83).Intersection(_1629_v83))
				(globalState).F17 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1628_v82, _1628_v82), _dafny.UnicodeSeqOfUtf8Bytes("fyxkjjg"))
			} else {
				(globalState).F2 = (func() bool {
					if _1538_v0 {
						return (func() bool {
							if ((_this).F26()).Contains(_1538_v0) {
								return ((_this).F26()).Get(_1538_v0).(bool)
							}
							return _1538_v0
						})()
					}
					return _1538_v0
				})()
				var _1630_v84 D0
				_ = _1630_v84
				_1630_v84 = Companion_D0_.Create_DC1_(_this.F27, _this.F27, _1538_v0)
				var _1631_v85 _dafny.Sequence
				_ = _1631_v85
				_1631_v85 = _dafny.UnicodeSeqOfUtf8Bytes("qljejp")
				var _1632_v86 _dafny.Set
				_ = _1632_v86
				_1632_v86 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_1631_v85).Cardinality())))
				_1630_v84 = Companion_D0_.Create_DC1_((_this).F24(), ((_this).F24()).Times(_dafny.IntOfInt64(632)), (_1632_v86).IsSubsetOf(_1632_v86))
				var _1633_v87 _dafny.Array
				_ = _1633_v87
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_37
				var _nw256 _dafny.Array
				_ = _nw256
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw256 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) bool = (func(_1634_v0 bool) func(_dafny.Int) bool {
						return func(_1635_i3 _dafny.Int) bool {
							return _1634_v0
						}
					})(_1538_v0)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw256 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw256).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw256).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1633_v87 = _nw256
				var _1636_v88 _dafny.Map
				_ = _1636_v88
				_1636_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, _1538_v0)
				var _1637_v89 _dafny.Map
				_ = _1637_v89
				_1637_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1636_v88).Contains(_dafny.IntOfInt64(918)) {
						return (_1636_v88).Get(_dafny.IntOfInt64(918)).(bool)
					}
					return _1538_v0
				})(), (_this).F24())
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1633_v87), 0))
				_ = _index218
				(_1633_v87).ArraySet1(!(!(_1637_v89).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1538_v0, _this.F27))), (_index218).Int())
				var _1638_v90 _dafny.Map
				_ = _1638_v90
				_1638_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1631_v85, _dafny.IntOfInt64(327))
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1633_v87), 0))
				_ = _index219
				var _rhs260 bool = _1538_v0
				_ = _rhs260
				var _rhs261 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1631_v85, Companion_Default___.Fm0(_this.F27, _this.F27, globalState))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1631_v85, (_this).F24())).Merge(func() _dafny.Map {
					var _coll46 = _dafny.NewMapBuilder()
					_ = _coll46
					for _iter50 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(422))).Uint32(), func(coer98 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg98 _dafny.Int) interface{} {
							return coer98(arg98)
						}
					}((func(_1639_v85 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1640_i4 _dafny.Int) _dafny.Sequence {
							return _1639_v85
						}
					})(_1631_v85)))).Elements()); ; {
						_compr_46, _ok50 := _iter50()
						if !_ok50 {
							break
						}
						var _1641_v91 _dafny.Sequence
						_1641_v91 = interface{}(_compr_46).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(422))).Uint32(), func(coer99 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg99 _dafny.Int) interface{} {
								return coer99(arg99)
							}
						}((func(_1642_v85 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_1640_i4 _dafny.Int) _dafny.Sequence {
								return _1642_v85
							}
						})(_1631_v85))), _1641_v91) {
							_coll46.Add(_1641_v91, _this.F27)
						}
					}
					return _coll46.ToMap()
				}()))
				_ = _rhs261
				var _lhs206 _dafny.Array = _1633_v87
				_ = _lhs206
				var _lhs207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1633_v87), 0))
				_ = _lhs207
				(_lhs206).ArraySet1(_rhs260, (_lhs207).Int())
				_1638_v90 = _rhs261
				(globalState).F9 = (_this).F24()
				var _1643_v92 _dafny.Array
				_ = _1643_v92
				var _nw257 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
				_ = _nw257
				_1643_v92 = _nw257
				var _1644_v93 _dafny.MultiSet
				_ = _1644_v93
				_1644_v93 = _dafny.MultiSetOf((_1633_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1633_v87), 0))).Int()).(bool))
				var _1645_v94 _dafny.Sequence
				_ = _1645_v94
				_1645_v94 = _dafny.SeqOf((_1637_v89).Cardinality())
				var _1646_v95 D3
				_ = _1646_v95
				_1646_v95 = Companion_D3_.Create_DC11_(_this.F27, (_1627_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1645_v94).Cardinality()), _dafny.IntOfUint32((_1627_v81).Cardinality()))).Uint32()).(bool), true)
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1633_v87), 0))
				_ = _index220
				var _rhs262 _dafny.MultiSet = (_1626_v80).Union(_1626_v80)
				_ = _rhs262
				var _rhs263 _dafny.Int = _this.F27
				_ = _rhs263
				var _rhs264 bool = (func() bool {
					if (_1633_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1633_v87), 0))).Int()).(bool) {
						return ((_1644_v93).Cardinality()).Cmp(_this.F27) < 0
					}
					return (_1646_v95).Dtor_cf20()
				})()
				_ = _rhs264
				var _rhs265 _dafny.Array = (func() _dafny.Array {
					if !((_1633_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1633_v87), 0))).Int()).(bool)) {
						return _1643_v92
					}
					return _1643_v92
				})()
				_ = _rhs265
				var _lhs208 *GlobalState = globalState
				_ = _lhs208
				var _lhs209 _dafny.Array = _1633_v87
				_ = _lhs209
				var _lhs210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1633_v87), 0))
				_ = _lhs210
				_1626_v80 = _rhs262
				_lhs208.F9 = _rhs263
				(_lhs209).ArraySet1(_rhs264, (_lhs210).Int())
				_1643_v92 = _rhs265
			}
			var _1647_v96 _dafny.Array
			_ = _1647_v96
			var _nw258 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw258
			_1647_v96 = _nw258
			var _1648_v97 D3
			_ = _1648_v97
			_1648_v97 = Companion_D3_.Create_DC11_(_this.F27, _1538_v0, _1538_v0)
			var _1649_v98 _dafny.Array
			_ = _1649_v98
			var _nwElement0_49 _dafny.Array = _1647_v96
			_ = _nwElement0_49
			var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(27))
			_ = _nw259
			(_nw259).ArraySet1(_nwElement0_49, 0)
			(_nw259).ArraySet1(_1647_v96, 1)
			(_nw259).ArraySet1((func() _dafny.Array {
				if _1538_v0 {
					return _1647_v96
				}
				return _1647_v96
			})(), 2)
			(_nw259).ArraySet1(_1647_v96, 3)
			(_nw259).ArraySet1(_1647_v96, 4)
			(_nw259).ArraySet1(_1647_v96, 5)
			(_nw259).ArraySet1(_1647_v96, 6)
			(_nw259).ArraySet1(_1647_v96, 7)
			(_nw259).ArraySet1(_1647_v96, 8)
			(_nw259).ArraySet1(_1647_v96, 9)
			(_nw259).ArraySet1(_1647_v96, 10)
			(_nw259).ArraySet1(_1647_v96, 11)
			(_nw259).ArraySet1(_1647_v96, 12)
			(_nw259).ArraySet1(_1647_v96, 13)
			(_nw259).ArraySet1(_1647_v96, 14)
			(_nw259).ArraySet1(_1647_v96, 15)
			(_nw259).ArraySet1(_1647_v96, 16)
			(_nw259).ArraySet1(_1647_v96, 17)
			(_nw259).ArraySet1(_1647_v96, 18)
			(_nw259).ArraySet1(_1647_v96, 19)
			(_nw259).ArraySet1(_1647_v96, 20)
			(_nw259).ArraySet1(_1647_v96, 21)
			(_nw259).ArraySet1((func() _dafny.Array {
				if (_1648_v97).Dtor_cf19() {
					return _1647_v96
				}
				return _1647_v96
			})(), 22)
			(_nw259).ArraySet1(_1647_v96, 23)
			(_nw259).ArraySet1(_1647_v96, 24)
			(_nw259).ArraySet1(_1647_v96, 25)
			(_nw259).ArraySet1(_1647_v96, 26)
			_1649_v98 = _nw259
			var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))
			_ = _index221
			(_1649_v98).ArraySet1(_1647_v96, (_index221).Int())
			var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))
			_ = _index222
			(_1649_v98).ArraySet1(_1647_v96, (_index222).Int())
			var _1650_v99 _dafny.MultiSet
			_ = _1650_v99
			_1650_v99 = _dafny.MultiSetOf(_1538_v0, _1538_v0, _1538_v0, _1538_v0, _1538_v0)
			var _1651_v100 _dafny.CodePoint
			_ = _1651_v100
			_1651_v100 = _dafny.CodePoint('n')
			var _1652_v101 _dafny.Map
			_ = _1652_v101
			_1652_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2((_this).F24(), (_this).F24(), _1651_v100, globalState), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _this.F27)).Cardinality())
			if (_dafny.MultiSetOf(!(_1538_v0), false)).IsSubsetOf((_1650_v99).Update(_1538_v0, Companion_Default___.Abs((_1652_v101).Cardinality()))) {
				r0 = (_1626_v80).IsDisjointFrom(_1626_v80)
				var _1653_v102 _dafny.Set
				_ = _1653_v102
				_1653_v102 = _dafny.SetOf((_this).F24())
				var _1654_v103 _dafny.Map
				_ = _1654_v103
				_1654_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1653_v102, (_this).F24())
				_1654_v103 = (_1654_v103).Update(Companion_Default___.Fm18(globalState), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1627_v81).Cardinality()), _this.F27))
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))
				_ = _arr0
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))), 0))
				_ = _index223
				(_arr0).ArraySet1(!(_1538_v0), (_index223).Int())
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))
				_ = _arr1
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))), 0))
				_ = _index224
				(_arr1).ArraySet1((Companion_Default___.Fm0((_this).F24(), (_this).F24(), globalState)).Cmp(_this.F27) <= 0, (_index224).Int())
				(globalState).F2 = _1538_v0
				var _1655_v104 _dafny.Sequence
				_ = _1655_v104
				_1655_v104 = _dafny.SeqOf(_1651_v100, _1651_v100, _1651_v100)
				var _1656_v105 D1
				_ = _1656_v105
				_1656_v105 = Companion_D1_.Create_DC3_(_1655_v104)
				var _1657_v106 _dafny.Set
				_ = _1657_v106
				_1657_v106 = _dafny.SetOf(_1626_v80)
				var _1658_v107 _dafny.Sequence
				_ = _1658_v107
				_1658_v107 = _dafny.SeqOf(_dafny.SetOf(_1626_v80, _1626_v80, _1626_v80))
				var _1659_v109 _dafny.Sequence
				_ = _1659_v109
				_1659_v109 = _dafny.SeqOf(_1626_v80, _1626_v80)
				var _1660_v110 *C1
				_ = _1660_v110
				var _nw260 *C1 = New_C1_()
				_ = _nw260
				_nw260.Ctor__(_1659_v109, (_this).F24(), (_this).F24(), (_this).F25())
				_1660_v110 = _nw260
				var _1661_v111 _dafny.Map
				_ = _1661_v111
				_1661_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1660_v110, (_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))), 0))).Int()).(bool))
				var _1662_v112 _dafny.Sequence
				_ = _1662_v112
				_1662_v112 = _dafny.SeqOf((func() _dafny.Int {
					if (_1626_v80).Contains((_this).F24()) {
						return (_1626_v80).Multiplicity((_this).F24())
					}
					return (_1661_v111).Cardinality()
				})())
				var _1663_v113 _dafny.Map
				_ = _1663_v113
				_1663_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1660_v110.F33, _dafny.SetOf(_dafny.MultiSetFromSeq(_1662_v112), _1626_v80))
				var _1664_v114 _dafny.Map
				_ = _1664_v114
				_1664_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1626_v80, _this.F27)
				var _1665_v116 _dafny.Array
				_ = _1665_v116
				var _nwElement0_50 _dafny.Set = _1657_v106
				_ = _nwElement0_50
				var _nw261 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(27))
				_ = _nw261
				(_nw261).ArraySet1(_nwElement0_50, 0)
				(_nw261).ArraySet1(_1657_v106, 1)
				(_nw261).ArraySet1((_1657_v106).Difference(_1657_v106), 2)
				(_nw261).ArraySet1(_1657_v106, 3)
				(_nw261).ArraySet1(_1657_v106, 4)
				(_nw261).ArraySet1((_1657_v106).Difference(_1657_v106), 5)
				(_nw261).ArraySet1(((_1658_v107).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1655_v104).Cardinality()), _dafny.IntOfUint32((_1658_v107).Cardinality()))).Uint32()).(_dafny.Set)).Difference(_1657_v106), 6)
				(_nw261).ArraySet1(_1657_v106, 7)
				(_nw261).ArraySet1(_1657_v106, 8)
				(_nw261).ArraySet1((func() _dafny.Set {
					if _1538_v0 {
						return _1657_v106
					}
					return _1657_v106
				})(), 9)
				(_nw261).ArraySet1(_1657_v106, 10)
				(_nw261).ArraySet1((_1657_v106).Intersection(Companion_Default___.Fm41((_this).F24(), Companion_Default___.Fm4(_1538_v0, (_this).F24(), _1651_v100, !(_1538_v0), globalState), !(!(false)), globalState)), 11)
				(_nw261).ArraySet1(func() _dafny.Set {
					var _coll47 = _dafny.NewBuilder()
					_ = _coll47
					for _iter51 := _dafny.Iterate((_1657_v106).Elements()); ; {
						_compr_47, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1666_v108 _dafny.MultiSet
						_1666_v108 = interface{}(_compr_47).(_dafny.MultiSet)
						if (_1657_v106).Contains(_1666_v108) {
							_coll47.Add(_1666_v108)
						}
					}
					return _coll47.ToSet()
				}(), 12)
				(_nw261).ArraySet1(((_1658_v107).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1662_v112).Cardinality()), _dafny.IntOfUint32((_1658_v107).Cardinality()))).Uint32()).(_dafny.Set)).Difference((func() _dafny.Set {
					if (_1663_v113).Contains((_dafny.Zero).Minus(_1660_v110.F33)) {
						return (_1663_v113).Get((_dafny.Zero).Minus(_1660_v110.F33)).(_dafny.Set)
					}
					return _1657_v106
				})()), 13)
				(_nw261).ArraySet1((func() _dafny.Set {
					if !(_1538_v0) {
						return _dafny.SetOf(_1626_v80)
					}
					return func() _dafny.Set {
						var _coll48 = _dafny.NewBuilder()
						_ = _coll48
						for _iter52 := _dafny.Iterate((_1664_v114).Keys().Elements()); ; {
							_compr_48, _ok52 := _iter52()
							if !_ok52 {
								break
							}
							var _1667_v115 _dafny.MultiSet
							_1667_v115 = interface{}(_compr_48).(_dafny.MultiSet)
							if (_1664_v114).Contains(_1667_v115) {
								_coll48.Add(_1667_v115)
							}
						}
						return _coll48.ToSet()
					}()
				})(), 14)
				(_nw261).ArraySet1(_1657_v106, 15)
				(_nw261).ArraySet1(_dafny.SetOf(_1626_v80, _dafny.MultiSetOf((_this).F24(), _1660_v110.F33)), 16)
				(_nw261).ArraySet1((_1657_v106).Difference(_1657_v106), 17)
				(_nw261).ArraySet1(_1657_v106, 18)
				(_nw261).ArraySet1(_1657_v106, 19)
				(_nw261).ArraySet1((_1658_v107).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_1658_v107).Cardinality()))).Uint32()).(_dafny.Set), 20)
				(_nw261).ArraySet1(_1657_v106, 21)
				(_nw261).ArraySet1(_1657_v106, 22)
				(_nw261).ArraySet1(_dafny.SetOf(_1626_v80), 23)
				(_nw261).ArraySet1(Companion_Default___.Fm41(_this.F27, _1655_v104, true, globalState), 24)
				(_nw261).ArraySet1(_1657_v106, 25)
				(_nw261).ArraySet1((_1657_v106).Difference(_1657_v106), 26)
				_1665_v116 = _nw261
				var _1668_v117 D8
				_ = _1668_v117
				_1668_v117 = Companion_D8_.Create_DC25_(_this.F27, (_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))), 0))).Int()).(bool), _1538_v0)
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1665_v116), 0))
				_ = _index225
				(_1665_v116).ArraySet1(Companion_Default___.Fm41((_this).F24(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1669_v100 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1670_i5 _dafny.Int) _dafny.CodePoint {
						return _1669_v100
					}
				})(_1651_v100))), (_1668_v117).Dtor_cf39(), globalState), (_index225).Int())
				var _1671_v118 _dafny.Array
				_ = _1671_v118
				var _nw262 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw262
				_1671_v118 = _nw262
				var _1672_v119 _dafny.Array
				_ = _1672_v119
				var _nwElement0_51 _dafny.Array = _1671_v118
				_ = _nwElement0_51
				var _nw263 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(4))
				_ = _nw263
				(_nw263).ArraySet1(_nwElement0_51, 0)
				(_nw263).ArraySet1(_1671_v118, 1)
				(_nw263).ArraySet1(_1671_v118, 2)
				(_nw263).ArraySet1(_1671_v118, 3)
				_1672_v119 = _nw263
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_1672_v119), 0))
				_ = _index226
				(_1672_v119).ArraySet1(_1671_v118, (_index226).Int())
				var _pat_let_tv26 = _1655_v104
				_ = _pat_let_tv26
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1665_v116), 0))
				_ = _index227
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_1672_v119), 0))
				_ = _index228
				var _rhs266 D1 = (func() D1 {
					if _1538_v0 {
						return _1656_v105
					}
					return (func() D1 {
						if (_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int()))), 0))).Int()).(bool) {
							return _1656_v105
						}
						return func(_pat_let32_0 D1) D1 {
							return func(_1673_dt__update__tmp_h1 D1) D1 {
								return func(_pat_let33_0 _dafny.Sequence) D1 {
									return func(_1674_dt__update_hcf7_h0 _dafny.Sequence) D1 {
										return Companion_D1_.Create_DC3_(_1674_dt__update_hcf7_h0)
									}(_pat_let33_0)
								}(_pat_let_tv26)
							}(_pat_let32_0)
						}(_1656_v105)
					})()
				})()
				_ = _rhs266
				var _rhs267 _dafny.Set = _1657_v106
				_ = _rhs267
				var _rhs268 _dafny.Array = _1671_v118
				_ = _rhs268
				var _lhs211 _dafny.Array = _1665_v116
				_ = _lhs211
				var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1665_v116), 0))
				_ = _lhs212
				var _lhs213 _dafny.Array = _1672_v119
				_ = _lhs213
				var _lhs214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_1672_v119), 0))
				_ = _lhs214
				_1656_v105 = _rhs266
				(_lhs211).ArraySet1(_rhs267, (_lhs212).Int())
				(_lhs213).ArraySet1(_rhs268, (_lhs214).Int())
			} else {
				(globalState).F14 = (func() _dafny.Int {
					if _1538_v0 {
						return (_this).F24()
					}
					return _this.F27
				})()
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))
				_ = _index229
				(_1649_v98).ArraySet1(_dafny.ArrayCastTo((_1649_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_1649_v98), 0))).Int())), (_index229).Int())
				var _1675_v120 *C2
				_ = _1675_v120
				var _nw264 *C2 = New_C2_()
				_ = _nw264
				_nw264.Ctor__((func() bool {
					if _1538_v0 {
						return _1538_v0
					}
					return _1538_v0
				})(), (_this).F24(), (_this).F25())
				_1675_v120 = _nw264
				(globalState).F14 = (_dafny.Zero).Minus(_this.F27)
				var _1676_v121 _dafny.Sequence
				_ = _1676_v121
				_1676_v121 = _dafny.UnicodeSeqOfUtf8Bytes("tlgkqke")
				var _1677_v122 *C2
				_ = _1677_v122
				var _nw265 *C2 = New_C2_()
				_ = _nw265
				_nw265.Ctor__(_1538_v0, (_this.F27).Minus(_dafny.IntOfUint32((_1676_v121).Cardinality())), (_this).F25())
				_1677_v122 = _nw265
			}
			var _1678_v123 _dafny.Array
			_ = _1678_v123
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_38
			var _nw266 _dafny.Array
			_ = _nw266
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw266 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) _dafny.Sequence = func(_1679_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sihbc"), _dafny.UnicodeSeqOfUtf8Bytes("egjfu"))
				}
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw266 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw266).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw266).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1678_v123 = _nw266
			var _1680_v124 _dafny.Sequence
			_ = _1680_v124
			_1680_v124 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(813))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg101 _dafny.Int) interface{} {
					return coer101(arg101)
				}
			}(func(_1681_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})))
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1678_v123), 0))
			_ = _index230
			(_1678_v123).ArraySet1((_1680_v124).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F27), _dafny.IntOfUint32((_1680_v124).Cardinality()))).Uint32()).(_dafny.Sequence), (_index230).Int())
			var _1682_v125 _dafny.Sequence
			_ = _1682_v125
			_1682_v125 = _dafny.UnicodeSeqOfUtf8Bytes("wugxlqdqp")
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1678_v123), 0))
			_ = _index231
			(_1678_v123).ArraySet1(_1682_v125, (_index231).Int())
			var _1683_v126 _dafny.Map
			_ = _1683_v126
			_1683_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this.F27).Times(_dafny.IntOfInt64(350)))
			var _1684_v127 _dafny.Sequence
			_ = _1684_v127
			_1684_v127 = _dafny.SeqOf((_1626_v80).Cardinality(), _this.F27)
			_1683_v126 = (_1683_v126).Update((_this).F24(), (_1684_v127).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_1684_v127).Cardinality()))).Uint32()).(_dafny.Int))
		} else {
			var _1685_v128 _dafny.Sequence
			_ = _1685_v128
			_1685_v128 = _dafny.SeqOf(_this.F27, (_this).F24())
			var _1686_v129 *C1
			_ = _1686_v129
			var _nw267 *C1 = New_C1_()
			_ = _nw267
			_nw267.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-395))).Uint32(), func(coer102 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}((func(_1687_v0 bool) func(_dafny.Int) _dafny.MultiSet {
				return func(_1688_i8 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_this.F27, (_dafny.MultiSetOf(_1687_v0, (Companion_D10_.Create_DC29_((_this).F24(), _1687_v0)).Dtor_cf46(), _1687_v0)).Cardinality())
				}
			})(_1538_v0))), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1685_v128, (Companion_Default___.SafeIndex((_1626_v80).Cardinality(), _dafny.IntOfUint32((_1685_v128).Cardinality()))).Uint32(), _this.F27)).Cardinality()), _this.F27, (_this).F25())
			_1686_v129 = _nw267
			var _1689_v130 _dafny.Map
			_ = _1689_v130
			_1689_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v129, _1538_v0)
			r0 = (func() bool {
				if (func() bool {
					if (_1689_v130).Contains(_1686_v129) {
						return (_1689_v130).Get(_1686_v129).(bool)
					}
					return _1538_v0
				})() {
					return true
				}
				return _1538_v0
			})()
			var _1690_v131 _dafny.Map
			_ = _1690_v131
			_1690_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1538_v0, _dafny.IntOfInt64(80))
			var _1691_v132 _dafny.Sequence
			_ = _1691_v132
			_1691_v132 = _dafny.SeqOf(_1538_v0, true)
			var _1692_v133 _dafny.Array
			_ = _1692_v133
			var _nwElement0_52 _dafny.Int = (_this).F24()
			_ = _nwElement0_52
			var _nw268 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(27))
			_ = _nw268
			(_nw268).ArraySet1(_nwElement0_52, 0)
			(_nw268).ArraySet1(Companion_Default___.SafeDivisionInt(_1686_v129.F33, (_1690_v131).Cardinality()), 1)
			(_nw268).ArraySet1(_1686_v129.F33, 2)
			(_nw268).ArraySet1((_this).F24(), 3)
			(_nw268).ArraySet1(_1686_v129.F33, 4)
			(_nw268).ArraySet1(_1686_v129.F33, 5)
			(_nw268).ArraySet1(_this.F27, 6)
			(_nw268).ArraySet1(_this.F27, 7)
			(_nw268).ArraySet1((_this).F24(), 8)
			(_nw268).ArraySet1((func() _dafny.Int {
				if (_1690_v131).Contains(_1538_v0) {
					return (_1690_v131).Get(_1538_v0).(_dafny.Int)
				}
				return (_this).F24()
			})(), 9)
			(_nw268).ArraySet1(_this.F27, 10)
			(_nw268).ArraySet1((_this).F24(), 11)
			(_nw268).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1691_v132).Cardinality()), _1686_v129.F33), 12)
			(_nw268).ArraySet1(Companion_Default___.SafeModuloInt((_this).F24(), _this.F27), 13)
			(_nw268).ArraySet1(_this.F27, 14)
			(_nw268).ArraySet1(_this.F27, 15)
			(_nw268).ArraySet1(Companion_Default___.SafeModuloInt(_this.F27, _this.F27), 16)
			(_nw268).ArraySet1(_this.F27, 17)
			(_nw268).ArraySet1(_1686_v129.F33, 18)
			(_nw268).ArraySet1(_dafny.IntOfInt64(676), 19)
			(_nw268).ArraySet1(_1686_v129.F33, 20)
			(_nw268).ArraySet1(_1686_v129.F33, 21)
			(_nw268).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfInt64(356)).Plus((_this).F24())), 22)
			(_nw268).ArraySet1(_this.F27, 23)
			(_nw268).ArraySet1((_this).F24(), 24)
			(_nw268).ArraySet1(_this.F27, 25)
			(_nw268).ArraySet1(_this.F27, 26)
			_1692_v133 = _nw268
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1692_v133), 0))
			_ = _index232
			(_1692_v133).ArraySet1(_1686_v129.F33, (_index232).Int())
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1692_v133), 0))
			_ = _index233
			(_1692_v133).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F24(), _dafny.IntOfInt64(574)), (_index233).Int())
			var _1693_v134 _dafny.CodePoint
			_ = _1693_v134
			_1693_v134 = _dafny.CodePoint('r')
			var _1694_v135 _dafny.Map
			_ = _1694_v135
			_1694_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1626_v80).Contains(_1686_v129.F33) {
					return (_1626_v80).Multiplicity(_1686_v129.F33)
				}
				return _1686_v129.F33
			})(), _1693_v134)
			var _1695_v137 _dafny.Sequence
			_ = _1695_v137
			_1695_v137 = _dafny.SeqOf(_1693_v134, _dafny.CodePoint('f'), _1693_v134, _1693_v134, _dafny.CodePoint('q'))
			var _1696_v138 _dafny.Sequence
			_ = _1696_v138
			_1696_v138 = _dafny.SeqOf(func() _dafny.Map {
				var _coll49 = _dafny.NewMapBuilder()
				_ = _coll49
				for _iter53 := _dafny.Iterate((_1695_v137).Elements()); ; {
					_compr_49, _ok53 := _iter53()
					if !_ok53 {
						break
					}
					var _1697_v136 _dafny.CodePoint
					_1697_v136 = interface{}(_compr_49).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_1695_v137, _1697_v136) {
						_coll49.Add(_1697_v136, true)
					}
				}
				return _coll49.ToMap()
			}())
			var _1698_v139 T0
			_ = _1698_v139
			var _nw269 *C1 = New_C1_()
			_ = _nw269
			_nw269.Ctor__((_1686_v129).F32(), _dafny.IntOfInt64(207), _dafny.IntOfUint32((_1696_v138).Cardinality()), (_this).F25())
			_1698_v139 = _nw269
			var _1699_v140 _dafny.Sequence
			_ = _1699_v140
			_1699_v140 = _dafny.SeqOf(_1698_v139)
			var _1700_v141 *C2
			_ = _1700_v141
			var _nw270 *C2 = New_C2_()
			_ = _nw270
			_nw270.Ctor__((_dafny.IntOfUint32((_dafny.SeqOf((_this).F24(), (_1694_v135).Cardinality(), _dafny.IntOfInt64(797), _this.F27, _dafny.IntOfUint32((_1699_v140).Cardinality()))).Cardinality())).Cmp(_1686_v129.F33) <= 0, (_this).F24(), (_this).F25())
			_1700_v141 = _nw270
			var _1701_v142 _dafny.Map
			_ = _1701_v142
			_1701_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(420), _1686_v129.F33)
			var _1702_v143 D8
			_ = _1702_v143
			_1702_v143 = Companion_D8_.Create_DC24_((func() _dafny.Int {
				if (_1701_v142).Contains(_this.F27) {
					return (_1701_v142).Get(_this.F27).(_dafny.Int)
				}
				return _1686_v129.F33
			})())
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1692_v133), 0))
			_ = _index234
			(_1692_v133).ArraySet1((_1702_v143).Dtor_cf37(), (_index234).Int())
			var _1703_v144 _dafny.Set
			_ = _1703_v144
			_1703_v144 = _dafny.SetOf(_1700_v141.F31)
			_1538_v0 = (_dafny.SetOf(_1538_v0)).IsProperSubsetOf(_1703_v144)
		}
		_1538_v0 = _1538_v0
		var _1704_v145 _dafny.Array
		_ = _1704_v145
		var _nw271 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(28))
		_ = _nw271
		_1704_v145 = _nw271
		for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1704_v145), 0))); ; {
			_guard_loop_4, _ok54 := _iter54()
			if !_ok54 {
				break
			}
			var _1705_i9 _dafny.Int
			_1705_i9 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1705_i9).Sign() != -1) && ((_1705_i9).Cmp(_dafny.ArrayLen((_1704_v145), 0)) < 0)) {
				(_1704_v145).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1538_v0), _dafny.SeqOf(_1538_v0))), (_1705_i9).Int())
			}
		}
		r0 = !(_1538_v0)
		return r0
	}
}
func (_this *C5) M3(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _1706_v0 _dafny.Sequence
		_ = _1706_v0
		_1706_v0 = _dafny.UnicodeSeqOfUtf8Bytes("jhyjv")
		var _1707_v1 _dafny.MultiSet
		_ = _1707_v1
		_1707_v1 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1706_v0).Cardinality()), Companion_Default___.Fm0(p0, _this.F27, globalState))
		var _1708_v2 _dafny.Map
		_ = _1708_v2
		_1708_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1707_v1, false)
		if (func() bool {
			if (_1708_v2).Contains(_1707_v1) {
				return (_1708_v2).Get(_1707_v1).(bool)
			}
			return (Companion_Default___.Fm1(_this.F27, _dafny.IntOfInt64(889), globalState)).IsSubsetOf(_1707_v1)
		})() {
			if (_dafny.IntOfInt64(866)).Cmp(p0) > 0 {
				var _1709_v3 bool
				_ = _1709_v3
				_1709_v3 = true
				var _1710_v4 *C2
				_ = _1710_v4
				var _nw272 *C2 = New_C2_()
				_ = _nw272
				_nw272.Ctor__(_1709_v3, p0, (_this).F25())
				_1710_v4 = _nw272
				(globalState).F17 = _1706_v0
				var _1711_v5 _dafny.Array
				_ = _1711_v5
				var _nw273 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw273
				_1711_v5 = _nw273
				var _1712_v6 _dafny.Sequence
				_ = _1712_v6
				_1712_v6 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(770))).Uint32(), func(coer103 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg103 _dafny.Int) interface{} {
						return coer103(arg103)
					}
				}((func(_1713_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1714_i0 _dafny.Int) _dafny.Int {
						return _1713_p0
					}
				})(p0)))).Cardinality()), _dafny.IntOfInt64(274), p0, _this.F27, p0)
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_1711_v5), 0))
				_ = _index235
				(_1711_v5).ArraySet1(_dafny.Companion_Sequence_.Update(_1712_v6, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1712_v6).Cardinality()))).Uint32(), _this.F27), (_index235).Int())
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_1711_v5), 0))
				_ = _index236
				(_1711_v5).ArraySet1(_1712_v6, (_index236).Int())
				var _1715_v7 _dafny.Array
				_ = _1715_v7
				var _nw274 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
				_ = _nw274
				_1715_v7 = _nw274
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1715_v7), 0))
				_ = _index237
				(_1715_v7).ArraySet1((_this.F27).Cmp(_this.F27) <= 0, (_index237).Int())
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1715_v7), 0))
				_ = _index238
				(_1715_v7).ArraySet1(_1709_v3, (_index238).Int())
				_1709_v3 = _1710_v4.F31
			} else {
				var _1716_v8 _dafny.Array
				_ = _1716_v8
				var _nw275 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
				_ = _nw275
				_1716_v8 = _nw275
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_1716_v8), 0))
				_ = _index239
				(_1716_v8).ArraySet1(_1706_v0, (_index239).Int())
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_1716_v8), 0))
				_ = _index240
				(_1716_v8).ArraySet1(_1706_v0, (_index240).Int())
				var _1717_v9 bool
				_ = _1717_v9
				_1717_v9 = true
				(globalState).F2 = !(_1717_v9)
				var _1718_v10 *C0
				_ = _1718_v10
				var _nw276 *C0 = New_C0_()
				_ = _nw276
				_nw276.Ctor__((true) == (_1717_v9))
				_1718_v10 = _nw276
				var _1719_v11 _dafny.Array
				_ = _1719_v11
				var _nwElement0_53 _dafny.Int = (_dafny.Zero).Minus((_this).F24())
				_ = _nwElement0_53
				var _nw277 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(15))
				_ = _nw277
				(_nw277).ArraySet1(_nwElement0_53, 0)
				(_nw277).ArraySet1(p0, 1)
				(_nw277).ArraySet1(p0, 2)
				(_nw277).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jk")).Cardinality()), 3)
				(_nw277).ArraySet1((_dafny.Zero).Minus((_this).F24()), 4)
				(_nw277).ArraySet1(p0, 5)
				(_nw277).ArraySet1(_this.F27, 6)
				(_nw277).ArraySet1((_this).F24(), 7)
				(_nw277).ArraySet1(_this.F27, 8)
				(_nw277).ArraySet1(_this.F27, 9)
				(_nw277).ArraySet1((_this).F24(), 10)
				(_nw277).ArraySet1(p0, 11)
				(_nw277).ArraySet1(_this.F27, 12)
				(_nw277).ArraySet1(p0, 13)
				(_nw277).ArraySet1(_this.F27, 14)
				_1719_v11 = _nw277
				var _1720_v12 _dafny.Sequence
				_ = _1720_v12
				_1720_v12 = _dafny.SeqOf(_1719_v11)
				var _1721_v13 _dafny.Map
				_ = _1721_v13
				_1721_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1719_v11)
				var _1722_v14 _dafny.Array
				_ = _1722_v14
				var _nwElement0_54 _dafny.Array = (_1720_v12).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1720_v12).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _nwElement0_54
				var _nw278 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(28))
				_ = _nw278
				(_nw278).ArraySet1(_nwElement0_54, 0)
				(_nw278).ArraySet1(_1719_v11, 1)
				(_nw278).ArraySet1(_1719_v11, 2)
				(_nw278).ArraySet1((func() _dafny.Array {
					if _1717_v9 {
						return _1719_v11
					}
					return _1719_v11
				})(), 3)
				(_nw278).ArraySet1(_1719_v11, 4)
				(_nw278).ArraySet1(_1719_v11, 5)
				(_nw278).ArraySet1(_1719_v11, 6)
				(_nw278).ArraySet1(_1719_v11, 7)
				(_nw278).ArraySet1(_1719_v11, 8)
				(_nw278).ArraySet1((_1720_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.IntOfUint32((_1720_v12).Cardinality()))).Uint32()).(_dafny.Array), 9)
				(_nw278).ArraySet1(_1719_v11, 10)
				(_nw278).ArraySet1(_1719_v11, 11)
				(_nw278).ArraySet1(_1719_v11, 12)
				(_nw278).ArraySet1(_1719_v11, 13)
				(_nw278).ArraySet1(_1719_v11, 14)
				(_nw278).ArraySet1((func() _dafny.Array {
					if _1717_v9 {
						return _1719_v11
					}
					return _1719_v11
				})(), 15)
				(_nw278).ArraySet1(_1719_v11, 16)
				(_nw278).ArraySet1(_1719_v11, 17)
				(_nw278).ArraySet1(_1719_v11, 18)
				(_nw278).ArraySet1(_1719_v11, 19)
				(_nw278).ArraySet1(_1719_v11, 20)
				(_nw278).ArraySet1(_1719_v11, 21)
				(_nw278).ArraySet1(_1719_v11, 22)
				(_nw278).ArraySet1(_1719_v11, 23)
				(_nw278).ArraySet1(_1719_v11, 24)
				(_nw278).ArraySet1(_1719_v11, 25)
				(_nw278).ArraySet1((func() _dafny.Array {
					if (_1721_v13).Contains((_this).F24()) {
						return (_1721_v13).Get((_this).F24()).(_dafny.Array)
					}
					return _1719_v11
				})(), 26)
				(_nw278).ArraySet1(_1719_v11, 27)
				_1722_v14 = _nw278
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_1722_v14), 0))
				_ = _index241
				(_1722_v14).ArraySet1(_1719_v11, (_index241).Int())
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_1722_v14), 0))
				_ = _index242
				(_1722_v14).ArraySet1(_1719_v11, (_index242).Int())
				var _1723_v15 _dafny.Array
				_ = _1723_v15
				var _nw279 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
				_ = _nw279
				_1723_v15 = _nw279
				var _1724_v16 _dafny.Array
				_ = _1724_v16
				var _nw280 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw280
				_1724_v16 = _nw280
				var _1725_v17 _dafny.Map
				_ = _1725_v17
				_1725_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1724_v16)
				var _1726_v18 D16
				_ = _1726_v18
				_1726_v18 = Companion_D16_.Create_DC40_(_1723_v15)
				var _rhs269 _dafny.Int = ((_dafny.IntOfInt64(691)).Times(_dafny.IntOfUint32((_1706_v0).Cardinality()))).Times((_1725_v17).Cardinality())
				_ = _rhs269
				var _rhs270 _dafny.Array = (_1726_v18).Dtor_cf63()
				_ = _rhs270
				var _lhs215 *C5 = _this
				_ = _lhs215
				_lhs215.F27 = _rhs269
				_1723_v15 = _rhs270
			}
			var _1727_v19 _dafny.Array
			_ = _1727_v19
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_39
			var _nw281 _dafny.Array
			_ = _nw281
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw281 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) bool = func(_1728_i1 _dafny.Int) bool {
					return true
				}
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw281 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw281).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw281).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1727_v19 = _nw281
			var _1729_v20 bool
			_ = _1729_v20
			_1729_v20 = true
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_1727_v19), 0))
			_ = _index243
			(_1727_v19).ArraySet1(_1729_v20, (_index243).Int())
			var _1730_v21 _dafny.Map
			_ = _1730_v21
			_1730_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1729_v20, _1729_v20)
			var _1731_v22 _dafny.Sequence
			_ = _1731_v22
			_1731_v22 = _dafny.SeqOf(_1706_v0, _1706_v0, _1706_v0)
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_1727_v19), 0))
			_ = _index244
			var _rhs271 bool = _1729_v20
			_ = _rhs271
			var _rhs272 _dafny.Map = (_1730_v21).Merge((_this).F26())
			_ = _rhs272
			var _rhs273 _dafny.Int = _dafny.IntOfUint32(((_1731_v22).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1731_v22).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
			_ = _rhs273
			var _lhs216 _dafny.Array = _1727_v19
			_ = _lhs216
			var _lhs217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_1727_v19), 0))
			_ = _lhs217
			var _lhs218 *GlobalState = globalState
			_ = _lhs218
			(_lhs216).ArraySet1(_rhs271, (_lhs217).Int())
			_1730_v21 = _rhs272
			_lhs218.F14 = _rhs273
			var _1732_v23 *C2
			_ = _1732_v23
			var _nw282 *C2 = New_C2_()
			_ = _nw282
			_nw282.Ctor__((_dafny.IntOfInt64(-986)).Cmp(_this.F27) <= 0, ((_dafny.Zero).Minus(_this.F27)).Minus(p0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(562))).Uint32(), func(coer104 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg104 _dafny.Int) interface{} {
					return coer104(arg104)
				}
			}((func(_1733_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_1734_i2 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(84))).Uint32(), func(coer105 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg105 _dafny.Int) interface{} {
							return coer105(arg105)
						}
					}((func(_1735_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1736_i3 _dafny.Int) _dafny.Int {
							return _1735_p0
						}
					})(_1733_p0)))).Cardinality()))
				}
			})(p0))))
			_1732_v23 = _nw282
			_1732_v23 = (func() *C2 {
				if !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("lylfv"), _1706_v0) {
					return _1732_v23
				}
				return _1732_v23
			})()
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_1727_v19), 0))
			_ = _index245
			(_1727_v19).ArraySet1(_1732_v23.F31, (_index245).Int())
			var _1737_v24 _dafny.CodePoint
			_ = _1737_v24
			_1737_v24 = _dafny.CodePoint('d')
			var _1738_v25 _dafny.Set
			_ = _1738_v25
			_1738_v25 = _dafny.SetOf(_1727_v19, _1727_v19)
			var _1739_v26 _dafny.Sequence
			_ = _1739_v26
			_1739_v26 = _dafny.SeqOf(_1738_v25, _1738_v25)
			var _1740_v27 _dafny.Map
			_ = _1740_v27
			_1740_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-379))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg106 _dafny.Int) interface{} {
					return coer106(arg106)
				}
			}(func(_1741_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), _dafny.UnicodeSeqOfUtf8Bytes("fg")), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-379))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg107 _dafny.Int) interface{} {
					return coer107(arg107)
				}
			}(func(_1742_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), _dafny.UnicodeSeqOfUtf8Bytes("fg"))).Cardinality()))).Uint32(), _1737_v24), (Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-379))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg108 _dafny.Int) interface{} {
					return coer108(arg108)
				}
			}(func(_1743_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), _dafny.UnicodeSeqOfUtf8Bytes("fg")), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-379))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg109 _dafny.Int) interface{} {
					return coer109(arg109)
				}
			}(func(_1744_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), _dafny.UnicodeSeqOfUtf8Bytes("fg"))).Cardinality()))).Uint32(), _1737_v24)).Cardinality()))).Uint32(), _1737_v24)).Cardinality()), (_1738_v25).Union((_1739_v26).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F27), _dafny.IntOfUint32((_1739_v26).Cardinality()))).Uint32()).(_dafny.Set)))
			var _1745_v28 _dafny.Map
			_ = _1745_v28
			_1745_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1737_v24, (_this).F24())
			var _1746_v29 _dafny.Array
			_ = _1746_v29
			var _nwElement0_55 _dafny.Int = _dafny.IntOfInt64(427)
			_ = _nwElement0_55
			var _nw283 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(23))
			_ = _nw283
			(_nw283).ArraySet1(_nwElement0_55, 0)
			(_nw283).ArraySet1(p0, 1)
			(_nw283).ArraySet1((_this).F24(), 2)
			(_nw283).ArraySet1(Companion_Default___.Fm0((_this).F24(), _this.F27, globalState), 3)
			(_nw283).ArraySet1(p0, 4)
			(_nw283).ArraySet1(p0, 5)
			(_nw283).ArraySet1(p0, 6)
			(_nw283).ArraySet1(p0, 7)
			(_nw283).ArraySet1(_this.F27, 8)
			(_nw283).ArraySet1((_this).F24(), 9)
			(_nw283).ArraySet1(_this.F27, 10)
			(_nw283).ArraySet1((_1745_v28).Cardinality(), 11)
			(_nw283).ArraySet1(p0, 12)
			(_nw283).ArraySet1(((_this).F26()).Cardinality(), 13)
			(_nw283).ArraySet1((_this).F24(), 14)
			(_nw283).ArraySet1(p0, 15)
			(_nw283).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1706_v0).Cardinality()))).Cardinality()), 16)
			(_nw283).ArraySet1(_dafny.IntOfInt64(-124), 17)
			(_nw283).ArraySet1(p0, 18)
			(_nw283).ArraySet1(_this.F27, 19)
			(_nw283).ArraySet1(_this.F27, 20)
			(_nw283).ArraySet1(_this.F27, 21)
			(_nw283).ArraySet1(_this.F27, 22)
			_1746_v29 = _nw283
			var _1747_v30 _dafny.Sequence
			_ = _1747_v30
			_1747_v30 = _dafny.SeqOf(_1746_v29, _1746_v29)
			var _rhs274 _dafny.Int = ((func() _dafny.Set {
				if (_1740_v27).Contains(_this.F27) {
					return (_1740_v27).Get(_this.F27).(_dafny.Set)
				}
				return _1738_v25
			})()).Cardinality()
			_ = _rhs274
			var _rhs275 bool = Companion_Default___.Fm2((p0).Plus(p0), p0, _1737_v24, globalState)
			_ = _rhs275
			var _rhs276 bool = (((_dafny.Zero).Minus(_this.F27)).Times(_dafny.IntOfUint32((_1747_v30).Cardinality()))).Cmp((_this.F27).Minus((_this).F24())) < 0
			_ = _rhs276
			var _rhs277 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0)
			_ = _rhs277
			var _rhs278 bool = true
			_ = _rhs278
			var _lhs219 *GlobalState = globalState
			_ = _lhs219
			var _lhs220 *GlobalState = globalState
			_ = _lhs220
			var _lhs221 *GlobalState = globalState
			_ = _lhs221
			var _lhs222 *GlobalState = globalState
			_ = _lhs222
			_lhs219.F14 = _rhs274
			_1729_v20 = _rhs275
			_lhs220.F2 = _rhs276
			_lhs221.F17 = _rhs277
			_lhs222.F2 = _rhs278
		} else {
			var _1748_v31 _dafny.CodePoint
			_ = _1748_v31
			_1748_v31 = _dafny.CodePoint('t')
			_1748_v31 = _1748_v31
			(globalState).F14 = p0
			var _1749_v32 bool
			_ = _1749_v32
			_1749_v32 = true
			var _1750_v33 _dafny.Array
			_ = _1750_v33
			var _nwElement0_56 bool = _1749_v32
			_ = _nwElement0_56
			var _nw284 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(2))
			_ = _nw284
			(_nw284).ArraySet1(_nwElement0_56, 0)
			(_nw284).ArraySet1(_1749_v32, 1)
			_1750_v33 = _nw284
			var _1751_v34 _dafny.Map
			_ = _1751_v34
			_1751_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, !(_1749_v32))
			var _1752_v35 _dafny.Map
			_ = _1752_v35
			_1752_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _1751_v34)
			var _1753_v36 _dafny.Array
			_ = _1753_v36
			var _nwElement0_57 _dafny.Map = (_1751_v34).Merge(_1751_v34)
			_ = _nwElement0_57
			var _nw285 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(22))
			_ = _nw285
			(_nw285).ArraySet1(_nwElement0_57, 0)
			(_nw285).ArraySet1(_1751_v34, 1)
			(_nw285).ArraySet1(_1751_v34, 2)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, _1749_v32), 3)
			(_nw285).ArraySet1(_1751_v34, 4)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, Companion_Default___.Fm2((_this).F24(), _this.F27, _1748_v31, globalState)), 5)
			(_nw285).ArraySet1(_1751_v34, 6)
			(_nw285).ArraySet1((_1751_v34).Merge((_1751_v34).Update(_1750_v33, !(_1749_v32))), 7)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, _1749_v32), 8)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, false), 9)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, _1749_v32), 10)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, _1749_v32), 11)
			(_nw285).ArraySet1((func() _dafny.Map {
				if _1749_v32 {
					return _1751_v34
				}
				return _1751_v34
			})(), 12)
			(_nw285).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, _1749_v32)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, _1749_v32)), 13)
			(_nw285).ArraySet1((func() _dafny.Map {
				if (_1752_v35).Contains((_dafny.Zero).Minus(p0)) {
					return (_1752_v35).Get((_dafny.Zero).Minus(p0)).(_dafny.Map)
				}
				return _1751_v34
			})(), 14)
			(_nw285).ArraySet1(_1751_v34, 15)
			(_nw285).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, !(_1749_v32)), 16)
			(_nw285).ArraySet1(_1751_v34, 17)
			(_nw285).ArraySet1(_1751_v34, 18)
			(_nw285).ArraySet1(_1751_v34, 19)
			(_nw285).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, _1749_v32)).Update(_1750_v33, _1749_v32)).Merge(_1751_v34), 20)
			(_nw285).ArraySet1(_1751_v34, 21)
			_1753_v36 = _nw285
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1753_v36), 0))
			_ = _index246
			(_1753_v36).ArraySet1(((_1751_v34).Update(_1750_v33, _1749_v32)).Update(_1750_v33, _1749_v32), (_index246).Int())
			var _1754_v37 _dafny.Sequence
			_ = _1754_v37
			_1754_v37 = _dafny.SeqOf(!(_1749_v32), _1749_v32)
			var _1755_v38 _dafny.Map
			_ = _1755_v38
			_1755_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1754_v37, _this.F27)
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_1753_v36), 0))
			_ = _index247
			(_1753_v36).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1750_v33, ((func() _dafny.Int {
				if (_1755_v38).Contains(Companion_Default___.Fm23(globalState)) {
					return (_1755_v38).Get(Companion_Default___.Fm23(globalState)).(_dafny.Int)
				}
				return (_this).F24()
			})()).Cmp(_this.F27) < 0), (_index247).Int())
			_1749_v32 = (_1754_v37).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1754_v37).Cardinality()))).Uint32()).(bool)
			(globalState).F9 = (_this.F27).Times((_this).F24())
		}
		var _1756_v39 bool
		_ = _1756_v39
		_1756_v39 = true
		var _1757_i5 _dafny.Int
		_ = _1757_i5
		_1757_i5 = _dafny.Zero
		{
			for _1756_v39 {
				{
					if (_1757_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1757_i5 = (_1757_i5).Plus(_dafny.One)
					var _1758_v40 _dafny.Map
					_ = _1758_v40
					_1758_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F24())
					_1758_v40 = (_1758_v40).Update((_this).F24(), ((_dafny.SetOf(_1756_v39, _1756_v39, _1756_v39)).Cardinality()).Minus(p0))
					var _1759_v41 *C0
					_ = _1759_v41
					var _nw286 *C0 = New_C0_()
					_ = _nw286
					_nw286.Ctor__(!(_1756_v39) || (_1756_v39))
					_1759_v41 = _nw286
					var _1760_v42 _dafny.CodePoint
					_ = _1760_v42
					_1760_v42 = _dafny.CodePoint('o')
					var _1761_v43 _dafny.Map
					_ = _1761_v43
					_1761_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("j"))
					var _1762_v44 _dafny.Sequence
					_ = _1762_v44
					_1762_v44 = _dafny.SeqOf((_1759_v41).F34(), true)
					var _1763_v45 _dafny.Map
					_ = _1763_v45
					_1763_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1762_v44).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1762_v44).Cardinality()))).Uint32()).(bool), Companion_Default___.Fm3(globalState))
					var _1764_v46 _dafny.Array
					_ = _1764_v46
					var _nwElement0_58 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("c")
					_ = _nwElement0_58
					var _nw287 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(29))
					_ = _nw287
					(_nw287).ArraySet1(_nwElement0_58, 0)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jgelqg"), _1706_v0), 1)
					(_nw287).ArraySet1(_1706_v0, 2)
					(_nw287).ArraySet1(Companion_Default___.Fm4(false, p0, _1760_v42, Companion_Default___.Fm2((_this).F24(), (_this).F24(), _1760_v42, globalState), globalState), 3)
					(_nw287).ArraySet1(_1706_v0, 4)
					(_nw287).ArraySet1((func() _dafny.Sequence {
						if (_1761_v43).Contains((_1759_v41).F34()) {
							return (_1761_v43).Get((_1759_v41).F34()).(_dafny.Sequence)
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(537))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg110 _dafny.Int) interface{} {
								return coer110(arg110)
							}
						}((func(_1765_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1766_i6 _dafny.Int) _dafny.CodePoint {
								return _1765_v42
							}
						})(_1760_v42)))
					})(), 5)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0), 6)
					(_nw287).ArraySet1(_1706_v0, 7)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1706_v0, (Companion_Default___.SafeIndex((_1763_v45).Cardinality(), _dafny.IntOfUint32((_1706_v0).Cardinality()))).Uint32(), _1760_v42), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1706_v0, (Companion_Default___.SafeIndex((_1763_v45).Cardinality(), _dafny.IntOfUint32((_1706_v0).Cardinality()))).Uint32(), _1760_v42)).Cardinality()))).Uint32(), _dafny.CodePoint('f')), 8)
					(_nw287).ArraySet1(_1706_v0, 9)
					(_nw287).ArraySet1(_1706_v0, 10)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0), 11)
					(_nw287).ArraySet1(_1706_v0, 12)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0), 13)
					(_nw287).ArraySet1(_1706_v0, 14)
					(_nw287).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg111 _dafny.Int) interface{} {
							return coer111(arg111)
						}
					}((func(_1767_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1768_i7 _dafny.Int) _dafny.CodePoint {
							return _1767_v42
						}
					})(_1760_v42))), 15)
					(_nw287).ArraySet1(Companion_Default___.Fm4(_1756_v39, _this.F27, _1760_v42, (_1759_v41).F34(), globalState), 16)
					(_nw287).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(947))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg112 _dafny.Int) interface{} {
							return coer112(arg112)
						}
					}((func(_1769_v42 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1770_i8 _dafny.Int) _dafny.CodePoint {
							return _1769_v42
						}
					})(_1760_v42))), 17)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0), 18)
					(_nw287).ArraySet1(_1706_v0, 19)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0), 20)
					(_nw287).ArraySet1(_1706_v0, 21)
					(_nw287).ArraySet1(_1706_v0, 22)
					(_nw287).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kdm"), 23)
					(_nw287).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("nlv"), 24)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Update(_1706_v0, (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1706_v0).Cardinality()))).Uint32(), _1760_v42), 25)
					(_nw287).ArraySet1(_1706_v0, 26)
					(_nw287).ArraySet1(_1706_v0, 27)
					(_nw287).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0), 28)
					_1764_v46 = _nw287
					var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_1764_v46), 0))
					_ = _index248
					(_1764_v46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0), (_index248).Int())
					var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_1764_v46), 0))
					_ = _index249
					(_1764_v46).ArraySet1(_1706_v0, (_index249).Int())
					var _1771_v47 _dafny.Array
					_ = _1771_v47
					var _nw288 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
					_ = _nw288
					_1771_v47 = _nw288
					var _1772_v49 _dafny.Sequence
					_ = _1772_v49
					_1772_v49 = _dafny.SeqOf(_1707_v1)
					var _1773_v50 _dafny.Map
					_ = _1773_v50
					_1773_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), (_this).F24())
					var _1774_v51 T0
					_ = _1774_v51
					var _nw289 *C1 = New_C1_()
					_ = _nw289
					_nw289.Ctor__(_1772_v49, (_1773_v50).Cardinality(), _dafny.IntOfUint32((_1706_v0).Cardinality()), (_this).F25())
					_1774_v51 = _nw289
					var _1775_v52 _dafny.Map
					_ = _1775_v52
					_1775_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-30), _1774_v51)
					var _1776_v53 _dafny.Sequence
					_ = _1776_v53
					_1776_v53 = _dafny.SeqOf(_this.F27, p0)
					var _1777_v54 _dafny.Sequence
					_ = _1777_v54
					_1777_v54 = _dafny.SeqOf((_1775_v52).Cardinality(), (_1776_v53).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-383), _dafny.IntOfUint32((_1776_v53).Cardinality()))).Uint32()).(_dafny.Int))
					var _1778_v55 D12
					_ = _1778_v55
					_1778_v55 = Companion_D12_.Create_DC33_(_1760_v42, (_1759_v41).F34())
					var _1779_v56 _dafny.Map
					_ = _1779_v56
					_1779_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1778_v55).Dtor_cf54(), (_1759_v41).F34())
					var _1780_v57 _dafny.Sequence
					_ = _1780_v57
					_1780_v57 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1777_v54, (Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_1777_v54).Cardinality()))).Uint32(), _dafny.IntOfInt64(132))).Cardinality()), (_this).F24(), (_1779_v56).Cardinality(), (_1707_v1).Cardinality(), _dafny.IntOfInt64(320))
					var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1771_v47), 0))
					_ = _index250
					(_1771_v47).ArraySet1(func() _dafny.Map {
						var _coll50 = _dafny.NewMapBuilder()
						_ = _coll50
						for _iter55 := _dafny.Iterate((_1780_v57).Elements()); ; {
							_compr_50, _ok55 := _iter55()
							if !_ok55 {
								break
							}
							var _1781_v48 _dafny.Int
							_1781_v48 = interface{}(_compr_50).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_1780_v57, _1781_v48) {
								_coll50.Add((_1781_v48).Plus(_dafny.IntOfInt64(9)), _this.F27)
							}
						}
						return _coll50.ToMap()
					}(), (_index250).Int())
					var _1782_v58 _dafny.Map
					_ = _1782_v58
					_1782_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1776_v53)
					var _1783_v59 _dafny.Map
					_ = _1783_v59
					_1783_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1759_v41).F34(), (func() _dafny.Sequence {
						if (_1782_v58).Contains(_1756_v39) {
							return (_1782_v58).Get(_1756_v39).(_dafny.Sequence)
						}
						return _1777_v54
					})())
					var _1784_v60 _dafny.MultiSet
					_ = _1784_v60
					_1784_v60 = _dafny.MultiSetOf(_1759_v41)
					var _1785_v61 D16
					_ = _1785_v61
					_1785_v61 = Companion_D16_.Create_DC41_((_1784_v60).Cardinality())
					var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1771_v47), 0))
					_ = _index251
					(_1771_v47).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, (func() _dafny.Int {
						if Companion_Default___.Fm2(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tvsoofsw")).Cardinality()), p0, _1760_v42, globalState) {
							return (_1783_v59).Cardinality()
						}
						return (_1785_v61).Dtor_cf64()
					})()), (_index251).Int())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		(globalState).F1 = p0
		var _1786_v62 *C0
		_ = _1786_v62
		var _nw290 *C0 = New_C0_()
		_ = _nw290
		_nw290.Ctor__(_1756_v39)
		_1786_v62 = _nw290
		var _1787_v63 _dafny.Array
		_ = _1787_v63
		var _nwElement0_59 *C0 = _1786_v62
		_ = _nwElement0_59
		var _nw291 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(25))
		_ = _nw291
		(_nw291).ArraySet1(_nwElement0_59, 0)
		(_nw291).ArraySet1(_1786_v62, 1)
		(_nw291).ArraySet1(_1786_v62, 2)
		(_nw291).ArraySet1(_1786_v62, 3)
		(_nw291).ArraySet1(_1786_v62, 4)
		(_nw291).ArraySet1(_1786_v62, 5)
		(_nw291).ArraySet1(_1786_v62, 6)
		(_nw291).ArraySet1(_1786_v62, 7)
		(_nw291).ArraySet1(_1786_v62, 8)
		(_nw291).ArraySet1(_1786_v62, 9)
		(_nw291).ArraySet1(_1786_v62, 10)
		(_nw291).ArraySet1(_1786_v62, 11)
		(_nw291).ArraySet1(_1786_v62, 12)
		(_nw291).ArraySet1(_1786_v62, 13)
		(_nw291).ArraySet1(_1786_v62, 14)
		(_nw291).ArraySet1(_1786_v62, 15)
		(_nw291).ArraySet1(_1786_v62, 16)
		(_nw291).ArraySet1(_1786_v62, 17)
		(_nw291).ArraySet1(_1786_v62, 18)
		(_nw291).ArraySet1(_1786_v62, 19)
		(_nw291).ArraySet1(_1786_v62, 20)
		(_nw291).ArraySet1(_1786_v62, 21)
		(_nw291).ArraySet1(_1786_v62, 22)
		(_nw291).ArraySet1(_1786_v62, 23)
		(_nw291).ArraySet1(_1786_v62, 24)
		_1787_v63 = _nw291
		_1787_v63 = _1787_v63
		var _1788_v64 _dafny.Array
		_ = _1788_v64
		var _len0_40 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_40
		var _nw292 _dafny.Array
		_ = _nw292
		if _len0_40.Cmp(_dafny.Zero) == 0 {
			_nw292 = _dafny.NewArray(_len0_40)
		} else {
			var _init40 func(_dafny.Int) _dafny.Int = (func(_1789_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1790_i9 _dafny.Int) _dafny.Int {
					return (_1790_i9).Plus(_1789_p0)
				}
			})(p0)
			_ = _init40
			var _element0_40 = _init40(_dafny.Zero)
			_ = _element0_40
			_nw292 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
			(_nw292).ArraySet1(_element0_40, 0)
			var _nativeLen0_40 = (_len0_40).Int()
			_ = _nativeLen0_40
			for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
				(_nw292).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
			}
		}
		_1788_v64 = _nw292
		var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))
		_ = _index252
		(_1788_v64).ArraySet1(_this.F27, (_index252).Int())
		var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))
		_ = _index253
		var _rhs279 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1706_v0, _1706_v0)).Cardinality())).Cmp((_this).F24()) <= 0
		_ = _rhs279
		var _rhs280 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1706_v0, _dafny.UnicodeSeqOfUtf8Bytes("tytwp"))).Cardinality()), Companion_Default___.Fm0((_this).F24(), p0, globalState)))
		_ = _rhs280
		var _lhs223 _dafny.Array = _1788_v64
		_ = _lhs223
		var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))
		_ = _lhs224
		_1756_v39 = _rhs279
		(_lhs223).ArraySet1(_rhs280, (_lhs224).Int())
		var _1791_v65 _dafny.CodePoint
		_ = _1791_v65
		_1791_v65 = _dafny.CodePoint('q')
		var _1792_i10 _dafny.Int
		_ = _1792_i10
		_1792_i10 = _dafny.Zero
		{
			for Companion_Default___.Fm2(p0, p0, _1791_v65, globalState) {
				{
					if (_1792_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1792_i10 = (_1792_i10).Plus(_dafny.One)
					_1707_v1 = _1707_v1
					if (func() bool {
						if !(!(_1756_v39)) || ((func() bool {
							if ((_this).F26()).Contains(_1756_v39) {
								return ((_this).F26()).Get(_1756_v39).(bool)
							}
							return (_1786_v62).F34()
						})()) {
							return false
						}
						return !((_1786_v62).F34())
					})() {
						var _1793_v66 _dafny.Sequence
						_ = _1793_v66
						_1793_v66 = _dafny.SeqOf((_1786_v62).F34())
						var _1794_v67 _dafny.Sequence
						_ = _1794_v67
						_1794_v67 = _dafny.SeqOf(_1793_v66)
						var _1795_v68 _dafny.Set
						_ = _1795_v68
						_1795_v68 = _dafny.SetOf((_1794_v67).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1706_v0).Cardinality()), _dafny.IntOfUint32((_1794_v67).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_1793_v66, _1793_v66), _1793_v66)
						var _1796_v69 D13
						_ = _1796_v69
						_1796_v69 = Companion_D13_.Create_DC36_((_1786_v62).F34(), _1756_v39)
						var _1797_v70 _dafny.Map
						_ = _1797_v70
						_1797_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1756_v39), (_1786_v62).F34())
						var _1798_v72 _dafny.Set
						_ = _1798_v72
						_1798_v72 = _1795_v68
						_1795_v68 = (func() _dafny.Set {
							if (_1796_v69).Dtor_cf58() {
								return _1795_v68
							}
							return (func() _dafny.Set {
								var _coll51 = _dafny.NewBuilder()
								_ = _coll51
								for _iter56 := _dafny.Iterate((_1797_v70).Keys().Elements()); ; {
									_compr_51, _ok56 := _iter56()
									if !_ok56 {
										break
									}
									var _1799_v71 _dafny.Sequence
									_1799_v71 = interface{}(_compr_51).(_dafny.Sequence)
									if (_1797_v70).Contains(_1799_v71) {
										_coll51.Add(_1799_v71)
									}
								}
								return _coll51.ToSet()
							}()).Intersection((_1798_v72))
						})()
						var _1800_v73 _dafny.Sequence
						_ = _1800_v73
						_1800_v73 = _dafny.SeqOf(_1707_v1, _1707_v1)
						var _1801_v74 _dafny.Map
						_ = _1801_v74
						_1801_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1756_v39, _1800_v73)
						var _1802_v75 *C1
						_ = _1802_v75
						var _nw293 *C1 = New_C1_()
						_ = _nw293
						_nw293.Ctor__((func() _dafny.Sequence {
							if (_1801_v74).Contains((_1786_v62).F34()) {
								return (_1801_v74).Get((_1786_v62).F34()).(_dafny.Sequence)
							}
							return _1800_v73
						})(), _this.F27, (_dafny.IntOfInt64(325)).Minus((_this).F24()), _dafny.Companion_Sequence_.Concatenate((_this).F25(), (_this).F25()))
						_1802_v75 = _nw293
						var _1803_v76 _dafny.Sequence
						_ = _1803_v76
						_1803_v76 = _dafny.SeqOf(_1802_v75)
						(globalState).F18 = (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1803_v76, _dafny.SeqOf(_1802_v75, _1802_v75))).Cardinality())).Times(p0))
						var _1804_v77 _dafny.Array
						_ = _1804_v77
						var _nw294 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
						_ = _nw294
						_1804_v77 = _nw294
						var _1805_v78 _dafny.Array
						_ = _1805_v78
						var _nwElement0_60 _dafny.Array = _1804_v77
						_ = _nwElement0_60
						var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(12))
						_ = _nw295
						(_nw295).ArraySet1(_nwElement0_60, 0)
						(_nw295).ArraySet1(_1804_v77, 1)
						(_nw295).ArraySet1(_1804_v77, 2)
						(_nw295).ArraySet1(_1804_v77, 3)
						(_nw295).ArraySet1(_1804_v77, 4)
						(_nw295).ArraySet1(_1804_v77, 5)
						(_nw295).ArraySet1((func() _dafny.Array {
							if false {
								return _1804_v77
							}
							return _1804_v77
						})(), 6)
						(_nw295).ArraySet1(_1804_v77, 7)
						(_nw295).ArraySet1(_1804_v77, 8)
						(_nw295).ArraySet1(_1804_v77, 9)
						(_nw295).ArraySet1(_1804_v77, 10)
						(_nw295).ArraySet1(_1804_v77, 11)
						_1805_v78 = _nw295
						var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_1805_v78), 0))
						_ = _index254
						(_1805_v78).ArraySet1(_1804_v77, (_index254).Int())
						var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_1805_v78), 0))
						_ = _index255
						(_1805_v78).ArraySet1(_1804_v77, (_index255).Int())
						(globalState).F1 = _1802_v75.F33
					} else {
						var _1806_v79 _dafny.Set
						_ = _1806_v79
						_1806_v79 = _dafny.SetOf((_this).F24())
						var _1807_v80 _dafny.Map
						_ = _1807_v80
						_1807_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int), _1806_v79)
						var _1808_v81 _dafny.Map
						_ = _1808_v81
						_1808_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), p0)
						var _1809_v82 _dafny.Map
						_ = _1809_v82
						_1809_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1791_v65, (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1706_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-891), _dafny.IntOfUint32((_1706_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), (_this).F24()), _1808_v81, _1808_v81, _1808_v81))
						var _1810_v83 *C2
						_ = _1810_v83
						var _nw296 *C2 = New_C2_()
						_ = _nw296
						_nw296.Ctor__(((func() _dafny.Set {
							if (_1807_v80).Contains(Companion_Default___.Fm0(p0, _this.F27, globalState)) {
								return (_1807_v80).Get(Companion_Default___.Fm0(p0, _this.F27, globalState)).(_dafny.Set)
							}
							return _1806_v79
						})()).IsDisjointFrom(_1806_v79), (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int), (func() _dafny.Sequence {
							if (_1809_v82).Contains((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)) {
								return (_1809_v82).Get((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
							}
							return (_this).F25()
						})())
						_1810_v83 = _nw296
						_1810_v83 = _1810_v83
						var _1811_v84 D13
						_ = _1811_v84
						_1811_v84 = Companion_D13_.Create_DC35_(_1707_v1)
						var _1812_v85 _dafny.Sequence
						_ = _1812_v85
						_1812_v85 = _dafny.SeqOf(_this.F27, (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)))
						(_1810_v83).F31 = ((_1811_v84).Dtor_cf57()).IsDisjointFrom(_dafny.MultiSetFromSeq(_1812_v85))
						var _1813_v86 _dafny.Set
						_ = _1813_v86
						_1813_v86 = _dafny.SetOf(_1756_v39, _1810_v83.F31)
						var _1814_v87 _dafny.Sequence
						_ = _1814_v87
						_1814_v87 = _dafny.SeqOf(_1813_v86)
						_1812_v85 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(796))).Uint32(), func(coer113 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg113 _dafny.Int) interface{} {
								return coer113(arg113)
							}
						}((func(_1815_v64 _dafny.Array) func(_dafny.Int) _dafny.Int {
							return func(_1816_i11 _dafny.Int) _dafny.Int {
								return (_1815_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1815_v64), 0))).Int()).(_dafny.Int)
							}
						})(_1788_v64))), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(796))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg114 _dafny.Int) interface{} {
								return coer114(arg114)
							}
						}((func(_1817_v64 _dafny.Array) func(_dafny.Int) _dafny.Int {
							return func(_1818_i11 _dafny.Int) _dafny.Int {
								return (_1817_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1817_v64), 0))).Int()).(_dafny.Int)
							}
						})(_1788_v64)))).Cardinality()))).Uint32(), _this.F27), (Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(796))).Uint32(), func(coer115 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg115 _dafny.Int) interface{} {
								return coer115(arg115)
							}
						}((func(_1819_v64 _dafny.Array) func(_dafny.Int) _dafny.Int {
							return func(_1820_i11 _dafny.Int) _dafny.Int {
								return (_1819_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1819_v64), 0))).Int()).(_dafny.Int)
							}
						})(_1788_v64))), (Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(796))).Uint32(), func(coer116 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg116 _dafny.Int) interface{} {
								return coer116(arg116)
							}
						}((func(_1821_v64 _dafny.Array) func(_dafny.Int) _dafny.Int {
							return func(_1822_i11 _dafny.Int) _dafny.Int {
								return (_1821_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1821_v64), 0))).Int()).(_dafny.Int)
							}
						})(_1788_v64)))).Cardinality()))).Uint32(), _this.F27)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1814_v87).Cardinality()))
						var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))
						_ = _index256
						(_1788_v64).ArraySet1((p0).Minus((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)), (_index256).Int())
						var _1823_v88 _dafny.Map
						_ = _1823_v88
						_1823_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xuayu"), (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int))).Cardinality(), _1810_v83.F31)
						var _1824_v89 _dafny.Array
						_ = _1824_v89
						var _nw297 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(25))
						_ = _nw297
						_1824_v89 = _nw297
						var _1825_v90 _dafny.Map
						_ = _1825_v90
						_1825_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, (func() _dafny.Array {
							if (func() bool {
								if (_1823_v88).Contains(_dafny.IntOfInt64(302)) {
									return (_1823_v88).Get(_dafny.IntOfInt64(302)).(bool)
								}
								return _1810_v83.F31
							})() {
								return _1824_v89
							}
							return _1824_v89
						})())
						var _1826_v91 _dafny.Map
						_ = _1826_v91
						_1826_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int), _1791_v65)
						var _rhs281 _dafny.Map = (_1825_v90).Merge(_1825_v90)
						_ = _rhs281
						var _rhs282 bool = (_1786_v62).F34()
						_ = _rhs282
						var _rhs283 _dafny.Map = (_1826_v91).Merge(_1826_v91)
						_ = _rhs283
						var _lhs225 *GlobalState = globalState
						_ = _lhs225
						_1825_v90 = _rhs281
						_lhs225.F2 = _rhs282
						_1826_v91 = _rhs283
					}
					if ((func() _dafny.Int {
						if _1756_v39 {
							return (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)
						}
						return (_this).F24()
					})()).Cmp((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)) == 0 {
						_1756_v39 = (_1786_v62).F34()
						var _1827_v92 _dafny.Sequence
						_ = _1827_v92
						_1827_v92 = _dafny.SeqOf(_1788_v64)
						var _1828_v93 _dafny.Array
						_ = _1828_v93
						var _nwElement0_61 _dafny.Array = _1788_v64
						_ = _nwElement0_61
						var _nw298 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(7))
						_ = _nw298
						(_nw298).ArraySet1(_nwElement0_61, 0)
						(_nw298).ArraySet1(_1788_v64, 1)
						(_nw298).ArraySet1(_1788_v64, 2)
						(_nw298).ArraySet1(_1788_v64, 3)
						(_nw298).ArraySet1(_1788_v64, 4)
						(_nw298).ArraySet1((_1827_v92).Select((Companion_Default___.SafeIndex(_this.F27, _dafny.IntOfUint32((_1827_v92).Cardinality()))).Uint32()).(_dafny.Array), 5)
						(_nw298).ArraySet1(_1788_v64, 6)
						_1828_v93 = _nw298
						var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_1828_v93), 0))
						_ = _index257
						(_1828_v93).ArraySet1(_1788_v64, (_index257).Int())
						var _1829_v94 _dafny.Sequence
						_ = _1829_v94
						_1829_v94 = _dafny.SeqOf(_dafny.IntOfInt64(-469))
						var _1830_v95 _dafny.Sequence
						_ = _1830_v95
						_1830_v95 = _dafny.SeqOf(true, true)
						var _1831_v96 _dafny.Sequence
						_ = _1831_v96
						_1831_v96 = _dafny.SeqOf(_1830_v95)
						var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))
						_ = _index258
						var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_1828_v93), 0))
						_ = _index259
						var _rhs284 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-763), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1829_v94, (Companion_Default___.SafeIndex((Companion_Default___.Fm6(_dafny.IntOfUint32((_1829_v94).Cardinality()), (_this).F24(), (_this).F24(), _this.F27, globalState)).Cardinality(), _dafny.IntOfUint32((_1829_v94).Cardinality()))).Uint32(), p0)).Cardinality())), (_dafny.IntOfInt64(423)).Times(_this.F27))
						_ = _rhs284
						var _rhs285 bool = (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32(((_1831_v96).Select((Companion_Default___.SafeIndex((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1831_v96).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Cardinality()))).IsDisjointFrom(_dafny.SetOf((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)))
						_ = _rhs285
						var _rhs286 _dafny.Array = _1788_v64
						_ = _rhs286
						var _rhs287 _dafny.Sequence = _1706_v0
						_ = _rhs287
						var _rhs288 _dafny.Sequence = _1706_v0
						_ = _rhs288
						var _lhs226 _dafny.Array = _1788_v64
						_ = _lhs226
						var _lhs227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))
						_ = _lhs227
						var _lhs228 *GlobalState = globalState
						_ = _lhs228
						var _lhs229 _dafny.Array = _1828_v93
						_ = _lhs229
						var _lhs230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_1828_v93), 0))
						_ = _lhs230
						var _lhs231 *GlobalState = globalState
						_ = _lhs231
						(_lhs226).ArraySet1(_rhs284, (_lhs227).Int())
						_lhs228.F2 = _rhs285
						(_lhs229).ArraySet1(_rhs286, (_lhs230).Int())
						_lhs231.F17 = _rhs287
						_1706_v0 = _rhs288
						var _1832_v97 _dafny.Array
						_ = _1832_v97
						var _nw299 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
						_ = _nw299
						_1832_v97 = _nw299
						var _1833_v98 _dafny.Sequence
						_ = _1833_v98
						_1833_v98 = _dafny.SeqOf(_1832_v97, _1832_v97)
						_1832_v97 = (_1833_v98).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F24()), _dafny.IntOfUint32((_1833_v98).Cardinality()))).Uint32()).(_dafny.Array)
						var _1834_v99 _dafny.Int
						_ = _1834_v99
						var _out33 _dafny.Int
						_ = _out33
						_out33 = Companion_Default___.M0(p0, (_this).F24(), globalState)
						_1834_v99 = _out33
						var _1835_v100 _dafny.Map
						_ = _1835_v100
						_1835_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_Default___.Fm23(globalState))
						var _1836_v101 _dafny.Map
						_ = _1836_v101
						_1836_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm1(((_1707_v1).Update(_this.F27, Companion_Default___.Abs((_dafny.Zero).Minus(_this.F27)))).Cardinality(), Companion_Default___.Fm0(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_1835_v100).Contains((_1786_v62).F34()) {
								return (_1835_v100).Get((_1786_v62).F34()).(_dafny.Sequence)
							}
							return _dafny.SeqOf((_1786_v62).F34())
						})()).Cardinality()), _this.F27, globalState), globalState)).Cardinality(), _1834_v99)
						var _1837_v103 _dafny.Sequence
						_ = _1837_v103
						_1837_v103 = _dafny.SeqOf(func() _dafny.Map {
							var _coll52 = _dafny.NewMapBuilder()
							_ = _coll52
							for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(593), _dafny.IntOfInt64(594))); ; {
								_compr_52, _ok57 := _iter57()
								if !_ok57 {
									break
								}
								var _1838_v102 _dafny.Int
								_1838_v102 = interface{}(_compr_52).(_dafny.Int)
								if ((_dafny.IntOfInt64(593)).Cmp(_1838_v102) <= 0) && ((_1838_v102).Cmp(_dafny.IntOfInt64(594)) < 0) {
									_coll52.Add((_1838_v102).Plus(_1834_v99), (_dafny.Zero).Minus((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)))
								}
							}
							return _coll52.ToMap()
						}(), _1836_v101)
						var _1839_v104 _dafny.Map
						_ = _1839_v104
						_1839_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1756_v39, p0)
						var _1840_v105 D3
						_ = _1840_v105
						_1840_v105 = Companion_D3_.Create_DC9_(_1830_v95)
						var _1841_v106 D1
						_ = _1841_v106
						_1841_v106 = Companion_D1_.Create_DC4_((_1786_v62).F34(), _1756_v39, _1791_v65, (_this).F24(), _dafny.IntOfUint32(((_1840_v105).Dtor_cf17()).Cardinality()))
						var _1842_v108 _dafny.Array
						_ = _1842_v108
						var _nwElement0_62 _dafny.Map = _1836_v101
						_ = _nwElement0_62
						var _nw300 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(22))
						_ = _nw300
						(_nw300).ArraySet1(_nwElement0_62, 0)
						(_nw300).ArraySet1((_1836_v101).Merge(_1836_v101), 1)
						(_nw300).ArraySet1((_1837_v103).Select((Companion_Default___.SafeIndex((_this).F24(), _dafny.IntOfUint32((_1837_v103).Cardinality()))).Uint32()).(_dafny.Map), 2)
						(_nw300).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F27), (_this).F24()), 3)
						(_nw300).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(871), _1834_v99)).Merge(_1836_v101), 4)
						(_nw300).ArraySet1((_1836_v101).Merge(_1836_v101), 5)
						(_nw300).ArraySet1(_1836_v101, 6)
						(_nw300).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, (_dafny.Zero).Minus(p0)), 7)
						(_nw300).ArraySet1((_1836_v101).Update(_dafny.IntOfInt64(187), (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)), 8)
						(_nw300).ArraySet1(_1836_v101, 9)
						(_nw300).ArraySet1((func() _dafny.Map {
							if true {
								return _1836_v101
							}
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
								if (_1839_v104).Contains(_1756_v39) {
									return (_1839_v104).Get(_1756_v39).(_dafny.Int)
								}
								return (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)
							})(), p0)
						})(), 10)
						(_nw300).ArraySet1(_1836_v101, 11)
						(_nw300).ArraySet1(_1836_v101, 12)
						(_nw300).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1830_v95).Cardinality()), _1834_v99), 13)
						(_nw300).ArraySet1((func() _dafny.Map {
							if _1756_v39 {
								return _1836_v101
							}
							return Companion_Default___.Fm15((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int), (_1841_v106).Dtor_cf8(), _dafny.IntOfInt64(-368), (_1786_v62).F34(), globalState)
						})(), 14)
						(_nw300).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, _dafny.IntOfUint32((_1829_v94).Cardinality()))).Merge(func() _dafny.Map {
							var _coll53 = _dafny.NewMapBuilder()
							_ = _coll53
							for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(794), _dafny.IntOfInt64(-694))); ; {
								_compr_53, _ok58 := _iter58()
								if !_ok58 {
									break
								}
								var _1843_v107 _dafny.Int
								_1843_v107 = interface{}(_compr_53).(_dafny.Int)
								if ((_dafny.IntOfInt64(794)).Cmp(_1843_v107) <= 0) && ((_1843_v107).Cmp(_dafny.IntOfInt64(-694)) < 0) {
									_coll53.Add((_1843_v107).Minus(p0), _this.F27)
								}
							}
							return _coll53.ToMap()
						}()), 15)
						(_nw300).ArraySet1(_1836_v101, 16)
						(_nw300).ArraySet1(_1836_v101, 17)
						(_nw300).ArraySet1(_1836_v101, 18)
						(_nw300).ArraySet1(_1836_v101, 19)
						(_nw300).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1834_v99, (_this).F24()), 20)
						(_nw300).ArraySet1(_1836_v101, 21)
						_1842_v108 = _nw300
						var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1842_v108), 0))
						_ = _index260
						(_1842_v108).ArraySet1(_1836_v101, (_index260).Int())
						var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1842_v108), 0))
						_ = _index261
						(_1842_v108).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F27, (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)), (_index261).Int())
					} else {
						_1791_v65 = _dafny.CodePoint('v')
						var _1844_v110 _dafny.MultiSet
						_ = _1844_v110
						_1844_v110 = _dafny.MultiSetOf(_1791_v65)
						var _1845_v112 *C2
						_ = _1845_v112
						var _nw301 *C2 = New_C2_()
						_ = _nw301
						_nw301.Ctor__(((_dafny.Zero).Minus((_this).F24())).Cmp(_dafny.IntOfInt64(318)) == 0, (_dafny.Zero).Minus(((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int)).Minus((_this).F24())), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(244))).Uint32(), func(coer117 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg117 _dafny.Int) interface{} {
								return coer117(arg117)
							}
						}((func(_1846_v65 _dafny.CodePoint, _1847_v64 _dafny.Array) func(_dafny.Int) _dafny.Map {
							return func(_1848_i12 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1846_v65, (_1847_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1847_v64), 0))).Int()).(_dafny.Int))
							}
						})(_1791_v65, _1788_v64))), (Companion_Default___.SafeIndex(Companion_Default___.Fm0((_dafny.Zero).Minus(p0), p0, globalState), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(244))).Uint32(), func(coer118 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg118 _dafny.Int) interface{} {
								return coer118(arg118)
							}
						}((func(_1849_v65 _dafny.CodePoint, _1850_v64 _dafny.Array) func(_dafny.Int) _dafny.Map {
							return func(_1851_i12 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1849_v65, (_1850_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1850_v64), 0))).Int()).(_dafny.Int))
							}
						})(_1791_v65, _1788_v64)))).Cardinality()))).Uint32(), func() _dafny.Map {
							var _coll54 = _dafny.NewMapBuilder()
							_ = _coll54
							for _iter59 := _dafny.Iterate((_1844_v110).Elements()); ; {
								_compr_54, _ok59 := _iter59()
								if !_ok59 {
									break
								}
								var _1852_v109 _dafny.CodePoint
								_1852_v109 = interface{}(_compr_54).(_dafny.CodePoint)
								if (_1844_v110).Contains(_1852_v109) {
									_coll54.Add(_1852_v109, (_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int))
								}
							}
							return _coll54.ToMap()
						}()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(586))).Uint32(), func(coer119 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg119 _dafny.Int) interface{} {
								return coer119(arg119)
							}
						}((func(_1853_v110 _dafny.MultiSet) func(_dafny.Int) _dafny.Map {
							return func(_1854_i13 _dafny.Int) _dafny.Map {
								return func() _dafny.Map {
									var _coll55 = _dafny.NewMapBuilder()
									_ = _coll55
									for _iter60 := _dafny.Iterate((_1853_v110).Elements()); ; {
										_compr_55, _ok60 := _iter60()
										if !_ok60 {
											break
										}
										var _1855_v111 _dafny.CodePoint
										_1855_v111 = interface{}(_compr_55).(_dafny.CodePoint)
										if (_1853_v110).Contains(_1855_v111) {
											_coll55.Add(_1855_v111, (_this).F24())
										}
									}
									return _coll55.ToMap()
								}()
							}
						})(_1844_v110)))))
						_1845_v112 = _nw301
						(globalState).F2 = (_1786_v62).F34()
						(_this).F27 = (p0).Times(Companion_Default___.SafeModuloInt((_this).F24(), (_1707_v1).Cardinality()))
						(globalState).F2 = _1845_v112.F31
					}
					var _1856_v114 _dafny.Sequence
					_ = _1856_v114
					_1856_v114 = _dafny.SeqOf((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int))
					var _1857_v115 _dafny.Map
					_ = _1857_v115
					_1857_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1856_v114).Cardinality()), p0)
					var _1858_v116 _dafny.Map
					_ = _1858_v116
					_1858_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1791_v65, _1857_v115)
					var _1859_v117 _dafny.Map
					_ = _1859_v117
					_1859_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1788_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1788_v64), 0))).Int()).(_dafny.Int), (_1786_v62).F34())
					var _1860_v118 _dafny.Set
					_ = _1860_v118
					_1860_v118 = _dafny.SetOf(_this.F27, (_dafny.Zero).Minus(((func() _dafny.Map {
						var _coll56 = _dafny.NewMapBuilder()
						_ = _coll56
						for _iter61 := _dafny.Iterate(((func() _dafny.Map {
							if (_1858_v116).Contains(_1791_v65) {
								return (_1858_v116).Get(_1791_v65).(_dafny.Map)
							}
							return _1857_v115
						})()).Keys().Elements()); ; {
							_compr_56, _ok61 := _iter61()
							if !_ok61 {
								break
							}
							var _1861_v113 _dafny.Int
							_1861_v113 = interface{}(_compr_56).(_dafny.Int)
							if ((func() _dafny.Map {
								if (_1858_v116).Contains(_1791_v65) {
									return (_1858_v116).Get(_1791_v65).(_dafny.Map)
								}
								return _1857_v115
							})()).Contains(_1861_v113) {
								_coll56.Add((_1861_v113).Plus(_this.F27), (_1786_v62).F34())
							}
						}
						return _coll56.ToMap()
					}()).Merge(_1859_v117)).Cardinality()))
					(globalState).F18 = (_1860_v118).Cardinality()
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
	}
}
func (_this *C5) F26() _dafny.Map {
	{
		return _this._f26
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
