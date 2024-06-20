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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false))).Cardinality()), (func() _dafny.Int {
		if !(!(false)) {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rakhhgk")).Cardinality())
		}
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ry")).Cardinality())
	})())
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("nkmmphev"), _dafny.CodePoint('d'))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("olpkpd"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(775))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 bool, p2 D1, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(981))).Uint32(), func(coer2 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) D1 {
		return Companion_D1_.Create_DC2_(_dafny.IntOfInt64(61), true, _dafny.IntOfInt64(-978), _dafny.IntOfInt64(52))
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D1_.Create_DC2_((_dafny.SetOf(false)).Cardinality(), true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ropqb")).Cardinality())), Companion_D1_.Create_DC2_(_dafny.IntOfInt64(422), false, _dafny.IntOfInt64(-728), _dafny.IntOfInt64(-268)), Companion_D1_.Create_DC2_(_dafny.IntOfInt64(157), false, _dafny.IntOfInt64(267), _dafny.IntOfInt64(-998))), _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(179), true, _dafny.IntOfInt64(207), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-541))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xv")).Cardinality())
	}))).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	if (func() bool {
		if !(true) {
			return false
		}
		return false
	})() {
		return _dafny.SeqOf(_dafny.IntOfInt64(-925))
	} else {
		return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(420), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-368))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-659)), _dafny.IntOfInt64(-693), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(548))).Cardinality())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf(false, true, true)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), !(true))).Cardinality())), _dafny.IntOfInt64(916)))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true)).Union(_dafny.SetOf(!(false), !(true)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D16 = (func() D16 {
		if true {
			return Companion_D16_.Create_DC45_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(142)))
		}
		return Companion_D16_.Create_DC45_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(87), false)).Cardinality(), _dafny.IntOfInt64(220)))
	})()
	_ = _source0
	if _source0.Is_DC46() {
		var _4___mcc_h0 _dafny.Int = _source0.Get_().(D16_DC46).Cf79
		_ = _4___mcc_h0
		var _5_cf79 _dafny.Int = _4___mcc_h0
		_ = _5_cf79
		return _dafny.CodePoint('d')
	} else if _source0.Is_DC45() {
		var _6___mcc_h1 _dafny.Map = _source0.Get_().(D16_DC45).Cf78
		_ = _6___mcc_h1
		var _7_cf78 _dafny.Map = _6___mcc_h1
		_ = _7_cf78
		return _dafny.CodePoint('g')
	} else {
		var _8___mcc_h2 D16 = _source0.Get_().(D16_DC47).Cf80
		_ = _8___mcc_h2
		var _9_cf80 D16 = _8___mcc_h2
		_ = _9_cf80
		return _dafny.CodePoint('a')
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gighne")).Cardinality()), _dafny.IntOfInt64(946))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(756), _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(289))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_10_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality()))).Cardinality())))).Merge(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(855), _dafny.IntOfInt64(-260), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(310), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dnvebqbov")).Cardinality()))).Cardinality())).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(509), _dafny.IntOfInt64(851))).Cardinality(), _dafny.IntOfInt64(965))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _11_v0 _dafny.Int
			_11_v0 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(855), _dafny.IntOfInt64(-260), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(310), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dnvebqbov")).Cardinality()))).Cardinality())).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(509), _dafny.IntOfInt64(851))).Cardinality(), _dafny.IntOfInt64(965)), _11_v0) {
				_coll0.Add((_11_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(!(false)))).Cardinality()))).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("n")).Cardinality())))
			}
		}
		return _coll0.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ahpb")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(202))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_12_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))).Cardinality())), (_dafny.Zero).Minus((_dafny.Zero).Minus(((_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false), false))).Cardinality()))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(904)))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('u'), _dafny.CodePoint('f'))).Cardinality()), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_13_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(627)
	}))).Cardinality()), false))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Map, p1 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(814), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-505))), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(704), _dafny.IntOfInt64(455))).Cardinality()), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("gf"), _dafny.UnicodeSeqOfUtf8Bytes("s"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-308))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false)).Cardinality()))).Intersection(_dafny.SetOf(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(893))).Cardinality(), _dafny.IntOfInt64(646))))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(Companion_D3_.Create_DC8_())).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _15_v0 D3
			_15_v0 = interface{}(_compr_1).(D3)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D3_.Create_DC8_()), _15_v0) {
				_coll1.Add(_15_v0, ((_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality())).Plus(_dafny.IntOfInt64(558)))
			}
		}
		return _coll1.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC1_(_dafny.IntOfInt64(230))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("gxerujrc"))
		}
		return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("nwiu"), _dafny.UnicodeSeqOfUtf8Bytes("fghlrjlu"))
	})()).Difference(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("efts")))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true)).Intersection(_dafny.MultiSetOf(true))).Union((_dafny.MultiSetOf(true)).Difference(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(558), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(417))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _17_v0 _dafny.MultiSet
			_17_v0 = interface{}(_compr_2).(_dafny.MultiSet)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(558), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(417))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Contains(_17_v0) {
				_coll2.Add(_17_v0, !(true))
			}
		}
		return _coll2.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(57), true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()), _dafny.IntOfInt64(656))).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _18_v1 _dafny.MultiSet
			_18_v1 = interface{}(_compr_3).(_dafny.MultiSet)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(57), true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()), _dafny.IntOfInt64(656))).Contains(_18_v1) {
				_coll3.Add(_18_v1, true)
			}
		}
		return _coll3.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(911))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(559))))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(!(true))).Intersection(_dafny.SetOf(true, false, false, false))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.MultiSet, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(957))).Cardinality(), (!(false)) && (false), (_dafny.IntOfInt64(568)).Minus(_dafny.IntOfInt64(800)), _dafny.IntOfInt64(154))
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('j'))).Cardinality(), false, (_dafny.SetOf(true, true, !(false), false, true)).Cardinality(), _dafny.IntOfInt64(15)), Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xc"), (_dafny.SetOf(true)).Cardinality())).Cardinality(), !(true), _dafny.IntOfInt64(679), _dafny.IntOfInt64(163))), _dafny.SeqOf(Companion_D1_.Create_DC2_((_dafny.SetOf(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(837))).Cardinality())).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_4).(_dafny.Int)
			if (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(837))).Cardinality())).Contains(_19_v0) {
				_coll4.Add(Companion_Default___.SafeModuloInt(_19_v0, _dafny.IntOfInt64(866)))
			}
		}
		return _coll4.ToSet()
	}(), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xydlmxbg")).Cardinality())))).Cardinality(), false, _dafny.IntOfInt64(683), (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("mxrqqpsh"))).Cardinality()), Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(31), _dafny.IntOfInt64(106))).Cardinality()), true, _dafny.IntOfInt64(860), _dafny.IntOfInt64(-392)), Companion_D1_.Create_DC2_(_dafny.IntOfInt64(116), true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mn")).Cardinality()), _dafny.IntOfInt64(-325)))), _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(627), !(false), _dafny.IntOfInt64(837), _dafny.IntOfInt64(429))))
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC10_()
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("wfu")
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(897), _dafny.IntOfInt64(-821)))
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.Map {
	if true {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))
	}
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.CodePoint {
	var _source1 D12 = Companion_D12_.Create_DC32_(_dafny.IntOfInt64(-862), true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("knta")).Cardinality()))
	_ = _source1
	if _source1.Is_DC32() {
		var _20___mcc_h0 _dafny.Int = _source1.Get_().(D12_DC32).Cf53
		_ = _20___mcc_h0
		var _21___mcc_h1 bool = _source1.Get_().(D12_DC32).Cf54
		_ = _21___mcc_h1
		var _22___mcc_h2 _dafny.Int = _source1.Get_().(D12_DC32).Cf55
		_ = _22___mcc_h2
		var _23_cf55 _dafny.Int = _22___mcc_h2
		_ = _23_cf55
		var _24_cf54 bool = _21___mcc_h1
		_ = _24_cf54
		var _25_cf53 _dafny.Int = _20___mcc_h0
		_ = _25_cf53
		return _dafny.CodePoint('h')
	} else if _source1.Is_DC31() {
		var _26___mcc_h3 _dafny.Sequence = _source1.Get_().(D12_DC31).Cf52
		_ = _26___mcc_h3
		var _27_cf52 _dafny.Sequence = _26___mcc_h3
		_ = _27_cf52
		return _dafny.CodePoint('p')
	} else {
		var _28___mcc_h4 D12 = _source1.Get_().(D12_DC33).Cf56
		_ = _28___mcc_h4
		var _29_cf56 D12 = _28___mcc_h4
		_ = _29_cf56
		return _dafny.CodePoint('c')
	}
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("opsm")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("tmbalwqgj"), _dafny.UnicodeSeqOfUtf8Bytes("uetnpshyt")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("s"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-573))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_30_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	})), _dafny.UnicodeSeqOfUtf8Bytes("s"))))
}
func (_static *CompanionStruct_Default___) Fm35(globalState *GlobalState) _dafny.CodePoint {
	if (func() bool {
		if true {
			return true
		}
		return !(!(true))
	})() {
		return _dafny.CodePoint('c')
	} else {
		return _dafny.CodePoint('b')
	}
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(901), _dafny.IntOfInt64(-213))).Cardinality()), _dafny.IntOfInt64(-297))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-323)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("nxungteu"))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('l'), _dafny.CodePoint('t'))
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_((_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-290))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(!(false)))), _dafny.IntOfInt64(519)))).Cardinality()), true, (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("pntnad"), _dafny.UnicodeSeqOfUtf8Bytes("vo"))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qvbi")).Cardinality()), _dafny.IntOfInt64(448))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm39(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D12_.Create_DC32_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Cardinality(), false, (func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(37), _dafny.IntOfInt64(25))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _31_v0 _dafny.Int
			_31_v0 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(37)).Cmp(_31_v0) <= 0) && ((_31_v0).Cmp(_dafny.IntOfInt64(25)) < 0) {
				_coll5.Add(Companion_Default___.SafeDivisionInt(_31_v0, (_dafny.MultiSetOf(_dafny.IntOfInt64(207))).Cardinality()))
			}
		}
		return _coll5.ToSet()
	}()).Cardinality())).Dtor_cf54(), false)))).Intersection(_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(true))))
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(213))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_32_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	})), _dafny.UnicodeSeqOfUtf8Bytes("rp")))
}
func (_static *CompanionStruct_Default___) Fm41(p0 D10, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(false)
}
func (_static *CompanionStruct_Default___) Fm42(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false, true, false))
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if !(false) {
			return _dafny.MultiSetOf(_dafny.SetOf(true, false, !(false)))
		}
		return _dafny.MultiSetOf(_dafny.SetOf(false, true))
	})()).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf(true, true, false, true), _dafny.SetOf(!(false)), _dafny.SetOf(false, false, false, !(false)), _dafny.SetOf(true), _dafny.SetOf(!(!(false))))))
}
func (_static *CompanionStruct_Default___) Fm44(globalState *GlobalState) D7 {
	var _source2 D14 = Companion_D14_.Create_DC38_(_dafny.CodePoint('p'), !(true), _dafny.SeqOf((Companion_D7_.Create_DC20_(true, _dafny.CodePoint('k'))).Dtor_cf36()))
	_ = _source2
	if _source2.Is_DC38() {
		var _33___mcc_h0 _dafny.CodePoint = _source2.Get_().(D14_DC38).Cf67
		_ = _33___mcc_h0
		var _34___mcc_h1 bool = _source2.Get_().(D14_DC38).Cf68
		_ = _34___mcc_h1
		var _35___mcc_h2 _dafny.Sequence = _source2.Get_().(D14_DC38).Cf69
		_ = _35___mcc_h2
		var _36_cf69 _dafny.Sequence = _35___mcc_h2
		_ = _36_cf69
		var _37_cf68 bool = _34___mcc_h1
		_ = _37_cf68
		var _38_cf67 _dafny.CodePoint = _33___mcc_h0
		_ = _38_cf67
		return Companion_D7_.Create_DC20_(false, _38_cf67)
	} else if _source2.Is_DC39() {
		var _39___mcc_h3 _dafny.Int = _source2.Get_().(D14_DC39).Cf70
		_ = _39___mcc_h3
		var _40___mcc_h4 _dafny.Map = _source2.Get_().(D14_DC39).Cf71
		_ = _40___mcc_h4
		var _41_cf71 _dafny.Map = _40___mcc_h4
		_ = _41_cf71
		var _42_cf70 _dafny.Int = _39___mcc_h3
		_ = _42_cf70
		if false {
			return Companion_D7_.Create_DC20_(true, _dafny.CodePoint('h'))
		} else {
			return Companion_D7_.Create_DC20_(false, _dafny.CodePoint('t'))
		}
	} else {
		var _43___mcc_h5 _dafny.Sequence = _source2.Get_().(D14_DC37).Cf66
		_ = _43___mcc_h5
		var _44_cf66 _dafny.Sequence = _43___mcc_h5
		_ = _44_cf66
		return Companion_D7_.Create_DC20_(false, _dafny.CodePoint('g'))
	}
}
func (_static *CompanionStruct_Default___) Fm45(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D1_.Create_DC3_(true, _dafny.UnicodeSeqOfUtf8Bytes("dwhk"), _dafny.IntOfInt64(-798))), _dafny.SeqOf(Companion_D1_.Create_DC3_(!(false), _dafny.UnicodeSeqOfUtf8Bytes("cc"), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nsv")).Cardinality())))))), _dafny.SeqOf(Companion_D1_.Create_DC3_(!(true), (Companion_D13_.Create_DC36_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.Zero)).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_45_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})), _dafny.IntOfInt64(369), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.MultiSetOf(true, false)).Cardinality())).Dtor_cf62(), _dafny.IntOfInt64(246)), Companion_D1_.Create_DC3_(true, _dafny.UnicodeSeqOfUtf8Bytes("e"), _dafny.IntOfInt64(89))))
}
func (_static *CompanionStruct_Default___) Fm46(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(703))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_46_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(945)
		}))).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _47_v0 _dafny.Int
			_47_v0 = interface{}(_compr_6).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(703))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}(func(_46_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(945)
			})), _47_v0) {
				_coll6.Add(Companion_Default___.SafeModuloInt(_47_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(655), _dafny.IntOfInt64(235))).Cardinality())), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bwex")).Cardinality()), _dafny.IntOfInt64(493))).Cardinality()))
			}
		}
		return _coll6.ToMap()
	}()).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(673))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_48_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	})))).Merge(func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(706), _dafny.IntOfInt64(867))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _49_v1 _dafny.Int
			_49_v1 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(706)).Cmp(_49_v1) <= 0) && ((_49_v1).Cmp(_dafny.IntOfInt64(867)) < 0) {
				_coll7.Add((_49_v1).Times((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(941)), _dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true)))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(420))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}(func(_50_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				}))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(284))))).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("pkrufmho"))
			}
		}
		return _coll7.ToMap()
	}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true, false, false)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(492))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_51_i3 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))))
}
func (_static *CompanionStruct_Default___) Fm47(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(936)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ti")).Cardinality()))).Merge(func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))).Cardinality()))).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _52_v0 _dafny.Sequence
			_52_v0 = interface{}(_compr_8).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))).Cardinality())), _52_v0) {
				_coll8.Add(_52_v0, _dafny.IntOfInt64(180))
			}
		}
		return _coll8.ToMap()
	}())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(347), (_dafny.SetOf(_dafny.IntOfInt64(-530))).Cardinality()), (_dafny.SetOf(_dafny.MultiSetOf(true))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(100))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_53_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(531), true)).Cardinality())).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _54_v1 _dafny.Int
			_54_v1 = interface{}(_compr_9).(_dafny.Int)
			if (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(531), true)).Cardinality())).Contains(_54_v1) {
				_coll9.Add((_54_v1).Times(_dafny.IntOfInt64(801)))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality())).Cardinality(), !(false))).Cardinality()), _dafny.IntOfInt64(660))))
}
func (_static *CompanionStruct_Default___) Fm48(p0 bool, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ks")).Cardinality()), (_dafny.Zero).Minus(((_dafny.SetOf(_dafny.IntOfInt64(-433), _dafny.IntOfInt64(647), _dafny.IntOfInt64(919))).Cardinality()).Minus(_dafny.IntOfInt64(852))))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, globalState *GlobalState) {
	var _55_v1 _dafny.Int
	_ = _55_v1
	_55_v1 = _dafny.IntOfInt64(-50)
	var _56_v2 _dafny.Map
	_ = _56_v2
	_56_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v1, _55_v1)
	var _57_v3 _dafny.Map
	_ = _57_v3
	_57_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_dafny.SeqOf(_55_v1)).Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _58_v0 _dafny.Int
			_58_v0 = interface{}(_compr_10).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_55_v1), _58_v0) {
				_coll10.Add((_58_v0).Times(_55_v1), p0)
			}
		}
		return _coll10.ToMap()
	}(), (_56_v2).Merge((_56_v2).Update(_55_v1, _55_v1)))
	var _59_v4 _dafny.MultiSet
	_ = _59_v4
	_59_v4 = _dafny.MultiSetOf(_55_v1)
	var _60_v5 _dafny.Map
	_ = _60_v5
	_60_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v1, Companion_Default___.Fm1(_59_v4, _55_v1, globalState))
	var _61_v6 D18
	_ = _61_v6
	_61_v6 = Companion_D18_.Create_DC52_(_57_v3)
	_57_v3 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_60_v5, _56_v2)).Merge((_61_v6).Dtor_cf88())).Merge(((_57_v3).Update(_60_v5, _56_v2)).Merge(_57_v3))
	_55_v1 = _55_v1
	if p0 {
		var _62_v7 _dafny.Sequence
		_ = _62_v7
		_62_v7 = _dafny.UnicodeSeqOfUtf8Bytes("dlxanf")
		_55_v1 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if (_56_v2).Contains((_dafny.Zero).Minus((_dafny.MultiSetOf(p0, p0)).Cardinality())) {
				return (_56_v2).Get((_dafny.Zero).Minus((_dafny.MultiSetOf(p0, p0)).Cardinality())).(_dafny.Int)
			}
			return _55_v1
		})(), _dafny.IntOfUint32((_62_v7).Cardinality()))
		var _63_v8 _dafny.Sequence
		_ = _63_v8
		_63_v8 = _dafny.SeqOf(_55_v1)
		_55_v1 = ((_dafny.MultiSetFromSeq(_63_v8)).Cardinality()).Times(_55_v1)
		var _64_v9 _dafny.Map
		_ = _64_v9
		_64_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_55_v1).Cmp(_55_v1) != 0)
		(globalState).F2 = (func() bool {
			if (_64_v9).Contains(p0) {
				return (_64_v9).Get(p0).(bool)
			}
			return Companion_Default___.Fm1(_59_v4, _55_v1, globalState)
		})()
		var _65_v10 _dafny.Sequence
		_ = _65_v10
		_65_v10 = _dafny.SeqOf(p0)
		_65_v10 = _65_v10
		var _66_v11 _dafny.CodePoint
		_ = _66_v11
		_66_v11 = _dafny.CodePoint('f')
		var _67_v12 T0
		_ = _67_v12
		var _nw0 *C1 = New_C1_()
		_ = _nw0
		_nw0.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer19 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_68_p0 bool, _69_v4 _dafny.MultiSet, _70_v1 _dafny.Int) func(_dafny.Int) D1 {
			return func(_71_i0 _dafny.Int) D1 {
				return Companion_D1_.Create_DC2_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_p0, _68_p0)).Cardinality()), _68_p0, (_dafny.SetOf((_69_v4).Cardinality())).Cardinality(), _70_v1)
			}
		})(p0, _59_v4, _55_v1))))
		_67_v12 = _nw0
		var _72_v13 D17
		_ = _72_v13
		_72_v13 = Companion_D17_.Create_DC50_(_66_v11, p0, _62_v7, _55_v1, _67_v12)
		_62_v7 = _dafny.Companion_Sequence_.Update((_72_v13).Dtor_cf84(), (Companion_Default___.SafeIndex(_55_v1, _dafny.IntOfUint32(((_72_v13).Dtor_cf84()).Cardinality()))).Uint32(), _66_v11)
	} else {
		var _73_v14 D1
		_ = _73_v14
		_73_v14 = Companion_D1_.Create_DC2_(_55_v1, (_55_v1).Cmp(_55_v1) <= 0, _55_v1, _55_v1)
		var _74_v15 _dafny.MultiSet
		_ = _74_v15
		_74_v15 = _dafny.MultiSetOf(p0, p0)
		var _75_v16 *C4
		_ = _75_v16
		var _nw1 *C4 = New_C4_()
		_ = _nw1
		_nw1.Ctor__(p0)
		_75_v16 = _nw1
		var _76_v17 _dafny.Map
		_ = _76_v17
		_76_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v16, p0)
		_73_v14 = Companion_D1_.Create_DC2_((_55_v1).Times(_55_v1), p0, _55_v1, ((_74_v15).Cardinality()).Plus((_76_v17).Cardinality()))
		(globalState).F2 = Companion_Default___.Fm1(_59_v4, Companion_Default___.SafeModuloInt(_55_v1, _55_v1), globalState)
		var _77_v18 _dafny.Sequence
		_ = _77_v18
		_77_v18 = _dafny.SeqOf(_73_v14)
		var _78_v19 *C1
		_ = _78_v19
		var _nw2 *C1 = New_C1_()
		_ = _nw2
		_nw2.Ctor__(_77_v18)
		_78_v19 = _nw2
		var _79_v20 _dafny.Sequence
		_ = _79_v20
		_79_v20 = _dafny.SeqOf((_dafny.MultiSetOf(p0, p0, p0, p0)).Cardinality(), _55_v1, _dafny.IntOfInt64(419), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0), (Companion_Default___.SafeIndex(_55_v1, _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))).Uint32(), true)).Cardinality()))
		var _80_v21 _dafny.Map
		_ = _80_v21
		_80_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_79_v20, _79_v20), _55_v1)
		var _81_v22 _dafny.Sequence
		_ = _81_v22
		_81_v22 = _dafny.UnicodeSeqOfUtf8Bytes("kqwwfaqw")
		var _82_v23 *C5
		_ = _82_v23
		var _nw3 *C5 = New_C5_()
		_ = _nw3
		_nw3.Ctor__(_55_v1, p0)
		_82_v23 = _nw3
		var _83_v24 _dafny.Sequence
		_ = _83_v24
		_83_v24 = _dafny.SeqOf(_82_v23, _82_v23, _82_v23)
		var _rhs0 bool = p0
		_ = _rhs0
		var _rhs1 _dafny.Map = _80_v21
		_ = _rhs1
		var _rhs2 bool = p0
		_ = _rhs2
		var _rhs3 _dafny.Sequence = _81_v22
		_ = _rhs3
		var _rhs4 _dafny.Int = _dafny.IntOfUint32((_83_v24).Cardinality())
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		_lhs0.F2 = _rhs0
		_80_v21 = _rhs1
		_lhs1.F2 = _rhs2
		_81_v22 = _rhs3
		_55_v1 = _rhs4
		var _84_v25 _dafny.Set
		_ = _84_v25
		_84_v25 = _dafny.SetOf(_79_v20)
		var _85_v26 _dafny.Sequence
		_ = _85_v26
		_85_v26 = _dafny.SeqOf(p0, p0, false, p0, p0)
		var _86_v27 _dafny.Sequence
		_ = _86_v27
		_86_v27 = _dafny.SeqOf(!(p0), (_85_v26).Select((Companion_Default___.SafeIndex((_82_v23).F10(), _dafny.IntOfUint32((_85_v26).Cardinality()))).Uint32()).(bool), true)
		var _87_v28 D11
		_ = _87_v28
		_87_v28 = Companion_D11_.Create_DC30_()
		if !((_82_v23).Fm5((_84_v25).Difference(_84_v25), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_82_v23).Fm4(p0, p0, p0, p0, globalState), _85_v26), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if p0 {
				return _dafny.SeqOf(_87_v28)
			}
			return _dafny.SeqOf(_87_v28)
		})()).Cardinality()), globalState)) {
			var _88_v29 _dafny.CodePoint
			_ = _88_v29
			_88_v29 = _dafny.CodePoint('x')
			var _89_v30 _dafny.Map
			_ = _89_v30
			_89_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _88_v29)
			var _90_v31 _dafny.Array
			_ = _90_v31
			var _nwElement0_0 _dafny.CodePoint = _88_v29
			_ = _nwElement0_0
			var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(23))
			_ = _nw4
			(_nw4).ArraySet1CodePoint(_nwElement0_0, 0)
			(_nw4).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_89_v30).Contains(p0) {
					return (_89_v30).Get(p0).(_dafny.CodePoint)
				}
				return _88_v29
			})(), 1)
			(_nw4).ArraySet1CodePoint(_88_v29, 2)
			(_nw4).ArraySet1CodePoint(_88_v29, 3)
			(_nw4).ArraySet1CodePoint(_88_v29, 4)
			(_nw4).ArraySet1CodePoint((func() _dafny.CodePoint {
				if !(!(!(p0))) {
					return _88_v29
				}
				return _88_v29
			})(), 5)
			(_nw4).ArraySet1CodePoint(_88_v29, 6)
			(_nw4).ArraySet1CodePoint(_88_v29, 7)
			(_nw4).ArraySet1CodePoint(_88_v29, 8)
			(_nw4).ArraySet1CodePoint(_88_v29, 9)
			(_nw4).ArraySet1CodePoint(_88_v29, 10)
			(_nw4).ArraySet1CodePoint(_88_v29, 11)
			(_nw4).ArraySet1CodePoint(_88_v29, 12)
			(_nw4).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return _88_v29
				}
				return _88_v29
			})(), 13)
			(_nw4).ArraySet1CodePoint(_88_v29, 14)
			(_nw4).ArraySet1CodePoint(_dafny.CodePoint('s'), 15)
			(_nw4).ArraySet1CodePoint(_dafny.CodePoint('d'), 16)
			(_nw4).ArraySet1CodePoint(_88_v29, 17)
			(_nw4).ArraySet1CodePoint(_dafny.CodePoint('i'), 18)
			(_nw4).ArraySet1CodePoint(_dafny.CodePoint('r'), 19)
			(_nw4).ArraySet1CodePoint(_88_v29, 20)
			(_nw4).ArraySet1CodePoint(_88_v29, 21)
			(_nw4).ArraySet1CodePoint((_81_v22).Select((Companion_Default___.SafeIndex((_82_v23).F10(), _dafny.IntOfUint32((_81_v22).Cardinality()))).Uint32()).(_dafny.CodePoint), 22)
			_90_v31 = _nw4
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_90_v31), 0))
			_ = _index0
			(_90_v31).ArraySet1CodePoint(_88_v29, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_90_v31), 0))
			_ = _index1
			(_90_v31).ArraySet1CodePoint(_88_v29, (_index1).Int())
			_79_v20 = _79_v20
			var _91_v32 *C2
			_ = _91_v32
			var _nw5 *C2 = New_C2_()
			_ = _nw5
			_nw5.Ctor__((_82_v23).F10(), !(p0) || (p0))
			_91_v32 = _nw5
			var _92_v33 _dafny.Array
			_ = _92_v33
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw6
			_92_v33 = _nw6
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_92_v33), 0))
			_ = _index2
			(_92_v33).ArraySet1(((_79_v20).Select((Companion_Default___.SafeIndex((_82_v23).F10(), _dafny.IntOfUint32((_79_v20).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_82_v23).F10()) >= 0, (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_92_v33), 0))
			_ = _index3
			(_92_v33).ArraySet1(false, (_index3).Int())
			(globalState).F2 = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("jb"), Companion_Default___.Fm6(_86_v27, globalState))
		} else {
			var _93_v34 _dafny.Array
			_ = _93_v34
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw7
			_93_v34 = _nw7
			var _94_v35 _dafny.Set
			_ = _94_v35
			_94_v35 = _dafny.SetOf(_93_v34, _93_v34)
			var _95_v36 _dafny.CodePoint
			_ = _95_v36
			_95_v36 = _dafny.CodePoint('m')
			var _96_v37 _dafny.Map
			_ = _96_v37
			_96_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_93_v34)).Union(_94_v35), _dafny.Companion_Sequence_.Update(_81_v22, (Companion_Default___.SafeIndex(_55_v1, _dafny.IntOfUint32((_81_v22).Cardinality()))).Uint32(), _95_v36))
			_96_v37 = (_96_v37).Update((_94_v35).Intersection(_94_v35), _81_v22)
			var _97_v38 *C0
			_ = _97_v38
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__(_84_v25, ((_82_v23).F10()).Minus((_82_v23).F10()))
			_97_v38 = _nw8
			var _98_v39 _dafny.Sequence
			_ = _98_v39
			_98_v39 = _dafny.SeqOf((_93_v34))
			var _rhs5 bool = p0
			_ = _rhs5
			var _rhs6 _dafny.Array = (_98_v39).Select((Companion_Default___.SafeIndex((_82_v23).F10(), _dafny.IntOfUint32((_98_v39).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs6
			var _rhs7 _dafny.Sequence = _79_v20
			_ = _rhs7
			var _rhs8 *C0 = _97_v38
			_ = _rhs8
			var _rhs9 _dafny.Int = _dafny.IntOfInt64(-96)
			_ = _rhs9
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			var _lhs3 *C0 = _97_v38
			_ = _lhs3
			_lhs2.F2 = _rhs5
			_93_v34 = _rhs6
			_79_v20 = _rhs7
			_97_v38 = _rhs8
			_lhs3.F12 = _rhs9
			(globalState).F2 = p0
			var _rhs10 _dafny.Int = ((_82_v23).F10()).Minus((_dafny.IntOfInt64(603)).Plus(_97_v38.F12))
			_ = _rhs10
			var _rhs11 _dafny.Int = _55_v1
			_ = _rhs11
			var _lhs4 *C0 = _97_v38
			_ = _lhs4
			_55_v1 = _rhs10
			_lhs4.F12 = _rhs11
			_55_v1 = (_dafny.Zero).Minus((_82_v23).F10())
		}
	}
	var _99_v40 _dafny.Set
	_ = _99_v40
	_99_v40 = _dafny.SetOf(_dafny.IntOfInt64(-54), _55_v1, _55_v1, _55_v1, _55_v1)
	var _100_i1 _dafny.Int
	_ = _100_i1
	_100_i1 = _dafny.Zero
	{
		for Companion_Default___.Fm1(_59_v4, Companion_Default___.SafeModuloInt((_99_v40).Cardinality(), _55_v1), globalState) {
			{
				if (_100_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_100_i1 = (_100_i1).Plus(_dafny.One)
				var _101_v41 _dafny.Array
				_ = _101_v41
				var _len0_0 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_0
				var _nw9 _dafny.Array
				_ = _nw9
				if _len0_0.Cmp(_dafny.Zero) == 0 {
					_nw9 = _dafny.NewArray(_len0_0)
				} else {
					var _init0 func(_dafny.Int) bool = (func(_102_p0 bool) func(_dafny.Int) bool {
						return func(_103_i2 _dafny.Int) bool {
							return _102_p0
						}
					})(p0)
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
				_101_v41 = _nw9
				var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_101_v41), 0))
				_ = _index4
				(_101_v41).ArraySet1(p0, (_index4).Int())
				var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_101_v41), 0))
				_ = _index5
				(_101_v41).ArraySet1(!(p0), (_index5).Int())
				(globalState).F2 = p0
				_55_v1 = (_55_v1).Plus(_55_v1)
				var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_101_v41), 0))
				_ = _index6
				(_101_v41).ArraySet1(p0, (_index6).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	if !(p0) || (p0) {
		var _104_v42 _dafny.Set
		_ = _104_v42
		_104_v42 = _99_v40
		var _105_v44 _dafny.Array
		_ = _105_v44
		var _nwElement0_1 _dafny.Set = _99_v40
		_ = _nwElement0_1
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(4))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_1, 0)
		(_nw10).ArraySet1((_104_v42), 1)
		(_nw10).ArraySet1(_99_v40, 2)
		(_nw10).ArraySet1(func() _dafny.Set {
			var _coll11 = _dafny.NewBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((_56_v2).Keys().Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _106_v43 _dafny.Int
				_106_v43 = interface{}(_compr_11).(_dafny.Int)
				if (_56_v2).Contains(_106_v43) {
					_coll11.Add((_106_v43).Times((_dafny.SetOf((_dafny.SetOf(_dafny.IntOfInt64(92), (_dafny.MultiSetOf(false, false)).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oaq")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(false, false)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(125))).Cardinality())).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(290))).Cardinality())).Cardinality(), _dafny.IntOfInt64(623), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()))
				}
			}
			return _coll11.ToSet()
		}(), 3)
		_105_v44 = _nw10
		var _107_v45 *C3
		_ = _107_v45
		var _nw11 *C3 = New_C3_()
		_ = _nw11
		_nw11.Ctor__(p0, _105_v44, !(p0))
		_107_v45 = _nw11
		_55_v1 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(globalState)))
		_55_v1 = (_55_v1).Times(_55_v1)
		(globalState).F2 = (func() bool {
			if p0 {
				return !(p0) || (p0)
			}
			return p0
		})()
		(globalState).F2 = !((_107_v45).F15())
	} else {
		var _108_v46 _dafny.Map
		_ = _108_v46
		_108_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v40, _55_v1)
		_108_v46 = (_108_v46).Merge(_108_v46)
		var _109_v47 _dafny.Sequence
		_ = _109_v47
		_109_v47 = _dafny.SeqOf(p0)
		var _110_v48 _dafny.Map
		_ = _110_v48
		_110_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _109_v47)
		var _111_v49 _dafny.MultiSet
		_ = _111_v49
		_111_v49 = _dafny.MultiSetOf(p0)
		var _112_v50 _dafny.Array
		_ = _112_v50
		var _nwElement0_2 _dafny.MultiSet = _dafny.MultiSetFromSeq((func() _dafny.Sequence {
			if (_110_v48).Contains(p0) {
				return (_110_v48).Get(p0).(_dafny.Sequence)
			}
			return _dafny.SeqOf(p0)
		})())
		_ = _nwElement0_2
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(10))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_2, 0)
		(_nw12).ArraySet1((_111_v49).Difference(_dafny.MultiSetOf(p0)), 1)
		(_nw12).ArraySet1((_111_v49).Update(p0, Companion_Default___.Abs(_55_v1)), 2)
		(_nw12).ArraySet1((_111_v49).Intersection(_111_v49), 3)
		(_nw12).ArraySet1(_111_v49, 4)
		(_nw12).ArraySet1(_111_v49, 5)
		(_nw12).ArraySet1((_111_v49).Union(_111_v49), 6)
		(_nw12).ArraySet1(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
			if p0 {
				return _109_v47
			}
			return _109_v47
		})()), 7)
		(_nw12).ArraySet1(_111_v49, 8)
		(_nw12).ArraySet1(_111_v49, 9)
		_112_v50 = _nw12
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_112_v50), 0))
		_ = _index7
		(_112_v50).ArraySet1(_111_v49, (_index7).Int())
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_112_v50), 0))
		_ = _index8
		(_112_v50).ArraySet1(_111_v49, (_index8).Int())
		var _113_v51 _dafny.Set
		_ = _113_v51
		_113_v51 = _dafny.SetOf(p0, p0)
		var _114_v52 D1
		_ = _114_v52
		_114_v52 = Companion_D1_.Create_DC2_(_55_v1, p0, (_113_v51).Cardinality(), _55_v1)
		var _115_v53 _dafny.Sequence
		_ = _115_v53
		_115_v53 = _dafny.SeqOf(_114_v52)
		var _116_v54 T0
		_ = _116_v54
		var _nw13 *C1 = New_C1_()
		_ = _nw13
		_nw13.Ctor__(_115_v53)
		_116_v54 = _nw13
		var _117_v55 _dafny.Set
		_ = _117_v55
		_117_v55 = _dafny.SetOf(_116_v54)
		_55_v1 = (_117_v55).Cardinality()
		var _118_v56 D10
		_ = _118_v56
		_118_v56 = Companion_D10_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _55_v1))
		var _119_v57 _dafny.Map
		_ = _119_v57
		_119_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _55_v1)
		var _pat_let_tv0 = _119_v57
		_ = _pat_let_tv0
		var _pat_let_tv1 = _119_v57
		_ = _pat_let_tv1
		var _source3 D10 = func(_pat_let0_0 D10) D10 {
			return func(_120_dt__update__tmp_h0 D10) D10 {
				return func(_pat_let1_0 _dafny.Map) D10 {
					return func(_121_dt__update_hcf50_h0 _dafny.Map) D10 {
						return Companion_D10_.Create_DC27_(_121_dt__update_hcf50_h0)
					}(_pat_let1_0)
				}((_pat_let_tv0).Merge(_pat_let_tv1))
			}(_pat_let0_0)
		}(_118_v56)
		_ = _source3
		if _source3.Is_DC28() {
			var _122_v58 _dafny.Array
			_ = _122_v58
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_1
			var _nw14 _dafny.Array
			_ = _nw14
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw14 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Sequence = func(_123_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("fvtkev")
				}
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw14 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw14).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw14).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_122_v58 = _nw14
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_122_v58), 0))
			_ = _index9
			(_122_v58).ArraySet1((Companion_D13_.Create_DC36_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(579))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_124_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			})), _dafny.IntOfInt64(327), _55_v1, _55_v1)).Dtor_cf62(), (_index9).Int())
			var _125_v59 _dafny.CodePoint
			_ = _125_v59
			_125_v59 = _dafny.CodePoint('m')
			var _126_v60 _dafny.Map
			_ = _126_v60
			_126_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v1, _116_v54)
			var _127_v61 _dafny.Map
			_ = _127_v61
			_127_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_126_v60).Cardinality(), _125_v59)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_122_v58), 0))
			_ = _index10
			var _rhs12 _dafny.Int = (func() _dafny.Int {
				if false {
					return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v1, _125_v59)).Merge(_127_v61)).Cardinality()
				}
				return _55_v1
			})()
			_ = _rhs12
			var _rhs13 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_128_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			}))
			_ = _rhs13
			var _rhs14 bool = Companion_Default___.Fm1(_59_v4, (func() _dafny.Int {
				if (func() bool {
					if (_60_v5).Contains(_55_v1) {
						return (_60_v5).Get(_55_v1).(bool)
					}
					return p0
				})() {
					return (_59_v4).Cardinality()
				}
				return _55_v1
			})(), globalState)
			_ = _rhs14
			var _lhs5 _dafny.Array = _122_v58
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_122_v58), 0))
			_ = _lhs6
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			_55_v1 = _rhs12
			(_lhs5).ArraySet1(_rhs13, (_lhs6).Int())
			_lhs7.F2 = _rhs14
			_55_v1 = _55_v1
			_113_v51 = _113_v51
			(globalState).F2 = (func() bool {
				if p0 {
					return !(p0)
				}
				return p0
			})()
		} else {
			var _129___mcc_h0 _dafny.Map = _source3.Get_().(D10_DC27).Cf50
			_ = _129___mcc_h0
			var _130_cf50 _dafny.Map = _129___mcc_h0
			_ = _130_cf50
			var _131_v62 *C4
			_ = _131_v62
			var _nw15 *C4 = New_C4_()
			_ = _nw15
			_nw15.Ctor__(true)
			_131_v62 = _nw15
			var _132_v63 _dafny.Int
			_ = _132_v63
			var _133_v64 D4
			_ = _133_v64
			var _134_v65 _dafny.MultiSet
			_ = _134_v65
			var _135_v66 bool
			_ = _135_v66
			var _out0 _dafny.Int
			_ = _out0
			var _out1 D4
			_ = _out1
			var _out2 _dafny.MultiSet
			_ = _out2
			var _out3 bool
			_ = _out3
			_out0, _out1, _out2, _out3 = (_131_v62).M7(_55_v1, _55_v1, false, globalState)
			_132_v63 = _out0
			_133_v64 = _out1
			_134_v65 = _out2
			_135_v66 = _out3
			(globalState).F2 = (_dafny.IntOfInt64(-737)).Cmp((_111_v49).Cardinality()) != 0
			var _136_v67 _dafny.Int
			_ = _136_v67
			var _137_v68 D4
			_ = _137_v68
			var _138_v69 _dafny.MultiSet
			_ = _138_v69
			var _139_v70 bool
			_ = _139_v70
			var _out4 _dafny.Int
			_ = _out4
			var _out5 D4
			_ = _out5
			var _out6 _dafny.MultiSet
			_ = _out6
			var _out7 bool
			_ = _out7
			_out4, _out5, _out6, _out7 = (_131_v62).M7(_132_v63, _55_v1, p0, globalState)
			_136_v67 = _out4
			_137_v68 = _out5
			_138_v69 = _out6
			_139_v70 = _out7
		}
		var _140_v71 D16
		_ = _140_v71
		_140_v71 = Companion_D16_.Create_DC46_(_55_v1)
		var _141_v72 _dafny.Map
		_ = _141_v72
		_141_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v71, p0)
		var _rhs15 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_55_v1))
		_ = _rhs15
		var _rhs16 _dafny.Int = (_55_v1).Times((func() _dafny.Int {
			if p0 {
				return Companion_Default___.Fm0(globalState)
			}
			return _55_v1
		})())
		_ = _rhs16
		var _rhs17 bool = Companion_Default___.Fm1(_59_v4, _55_v1, globalState)
		_ = _rhs17
		var _rhs18 _dafny.Map = (func() _dafny.Map {
			if (Companion_Default___.Fm0(globalState)).Cmp(_55_v1) <= 0 {
				return _141_v72
			}
			return _141_v72
		})()
		_ = _rhs18
		var _lhs8 *GlobalState = globalState
		_ = _lhs8
		_55_v1 = _rhs15
		_55_v1 = _rhs16
		_lhs8.F2 = _rhs17
		_141_v72 = _rhs18
	}
	var _142_i6 _dafny.Int
	_ = _142_i6
	_142_i6 = _dafny.Zero
	{
		for p0 {
			{
				if (_142_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_142_i6 = (_142_i6).Plus(_dafny.One)
				var _143_v73 _dafny.Sequence
				_ = _143_v73
				_143_v73 = _dafny.SeqOf(false, p0)
				var _144_v74 T2
				_ = _144_v74
				var _nw16 *C5 = New_C5_()
				_ = _nw16
				_nw16.Ctor__(_55_v1, p0)
				_144_v74 = _nw16
				var _145_v75 _dafny.Map
				_ = _145_v75
				_145_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_143_v73).Cardinality()), _144_v74)
				var _146_v76 _dafny.Sequence
				_ = _146_v76
				_146_v76 = _dafny.UnicodeSeqOfUtf8Bytes("mivytavqf")
				var _147_v77 *C4
				_ = _147_v77
				var _nw17 *C4 = New_C4_()
				_ = _nw17
				_nw17.Ctor__(p0)
				_147_v77 = _nw17
				var _148_v78 _dafny.Sequence
				_ = _148_v78
				_148_v78 = _dafny.SeqOf(_147_v77, _147_v77)
				var _149_v79 _dafny.Sequence
				_ = _149_v79
				_149_v79 = _dafny.SeqOf(_dafny.SeqOf(_147_v77), _148_v78, _dafny.SeqOf(_147_v77, _147_v77, _147_v77))
				var _150_v80 _dafny.Map
				_ = _150_v80
				_150_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v76, _55_v1)
				var _151_v81 _dafny.Sequence
				_ = _151_v81
				_151_v81 = _dafny.SeqOf(_146_v76)
				var _152_v82 _dafny.Array
				_ = _152_v82
				var _nwElement0_3 _dafny.Int = (_dafny.SetOf(_143_v73, _143_v73)).Cardinality()
				_ = _nwElement0_3
				var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(29))
				_ = _nw18
				(_nw18).ArraySet1(_nwElement0_3, 0)
				(_nw18).ArraySet1(_dafny.IntOfInt64(616), 1)
				(_nw18).ArraySet1((_55_v1).Times(_55_v1), 2)
				(_nw18).ArraySet1((((_145_v75).Update((_59_v4).Cardinality(), _144_v74)).Merge(_145_v75)).Cardinality(), 3)
				(_nw18).ArraySet1(_55_v1, 4)
				(_nw18).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_55_v1).Times(_55_v1))), 5)
				(_nw18).ArraySet1((_55_v1).Plus(_55_v1), 6)
				(_nw18).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(569))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_153_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				})), _146_v76)).Cardinality()), 7)
				(_nw18).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((Companion_Default___.Fm6(_143_v73, globalState)).Cardinality()), _dafny.IntOfUint32((_149_v79).Cardinality())), 8)
				(_nw18).ArraySet1(_55_v1, 9)
				(_nw18).ArraySet1(_55_v1, 10)
				(_nw18).ArraySet1(_55_v1, 11)
				(_nw18).ArraySet1(_dafny.IntOfUint32((_146_v76).Cardinality()), 12)
				(_nw18).ArraySet1((_55_v1).Minus(_55_v1), 13)
				(_nw18).ArraySet1(_dafny.IntOfInt64(989), 14)
				(_nw18).ArraySet1(_dafny.IntOfInt64(912), 15)
				(_nw18).ArraySet1((_55_v1).Minus((func() _dafny.Int {
					if (_150_v80).Contains((_151_v81).Select((Companion_Default___.SafeIndex(_55_v1, _dafny.IntOfUint32((_151_v81).Cardinality()))).Uint32()).(_dafny.Sequence)) {
						return (_150_v80).Get((_151_v81).Select((Companion_Default___.SafeIndex(_55_v1, _dafny.IntOfUint32((_151_v81).Cardinality()))).Uint32()).(_dafny.Sequence)).(_dafny.Int)
					}
					return _dafny.IntOfInt64(967)
				})()), 16)
				(_nw18).ArraySet1(_55_v1, 17)
				(_nw18).ArraySet1(_dafny.IntOfInt64(530), 18)
				(_nw18).ArraySet1(_55_v1, 19)
				(_nw18).ArraySet1(_dafny.IntOfInt64(322), 20)
				(_nw18).ArraySet1(_55_v1, 21)
				(_nw18).ArraySet1((_dafny.IntOfUint32((Companion_Default___.Fm20(p0, !((_144_v74).F9()), _dafny.IntOfInt64(-541), _146_v76, globalState)).Cardinality())).Times(_55_v1), 22)
				(_nw18).ArraySet1(Companion_Default___.Fm0(globalState), 23)
				(_nw18).ArraySet1((_dafny.Zero).Minus(_55_v1), 24)
				(_nw18).ArraySet1((_55_v1).Times(_55_v1), 25)
				(_nw18).ArraySet1(_55_v1, 26)
				(_nw18).ArraySet1(_55_v1, 27)
				(_nw18).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()), 28)
				_152_v82 = _nw18
				var _154_v83 _dafny.CodePoint
				_ = _154_v83
				_154_v83 = _dafny.CodePoint('a')
				var _155_v84 _dafny.Set
				_ = _155_v84
				_155_v84 = _dafny.SetOf(_154_v83, _154_v83)
				var _156_v85 _dafny.Map
				_ = _156_v85
				_156_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _155_v84)
				var _rhs19 _dafny.Array = _152_v82
				_ = _rhs19
				var _rhs20 bool = (_156_v85).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_144_v74).F9()), _155_v84))
				_ = _rhs20
				var _lhs9 *GlobalState = globalState
				_ = _lhs9
				_152_v82 = _rhs19
				_lhs9.F2 = _rhs20
				var _157_v86 _dafny.Map
				_ = _157_v86
				_157_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_144_v74).F9(), _144_v74)
				_55_v1 = (((_157_v86).Merge(_157_v86)).Cardinality()).Minus(_55_v1)
				var _158_v87 _dafny.MultiSet
				_ = _158_v87
				_158_v87 = _dafny.MultiSetOf(_56_v2, _56_v2)
				_55_v1 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_55_v1, _dafny.IntOfUint32((_146_v76).Cardinality())), (func() _dafny.Int {
					if (_158_v87).Contains(_56_v2) {
						return (_158_v87).Multiplicity(_56_v2)
					}
					return _dafny.IntOfInt64(540)
				})())
				var _159_v88 _dafny.Array
				_ = _159_v88
				var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
				_ = _nw19
				_159_v88 = _nw19
				var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_159_v88), 0))
				_ = _index11
				(_159_v88).ArraySet1((_60_v5).Merge(Companion_Default___.Fm15((_144_v74).F9(), globalState)), (_index11).Int())
				var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_159_v88), 0))
				_ = _index12
				(_159_v88).ArraySet1((func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter12 := _dafny.Iterate((_60_v5).Keys().Elements()); ; {
						_compr_12, _ok12 := _iter12()
						if !_ok12 {
							break
						}
						var _160_v89 _dafny.Int
						_160_v89 = interface{}(_compr_12).(_dafny.Int)
						if (_60_v5).Contains(_160_v89) {
							_coll12.Add((_160_v89).Plus(_55_v1), p0)
						}
					}
					return _coll12.ToMap()
				}()).Merge(_60_v5), (_index12).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _161_v0 bool
	_ = _161_v0
	_161_v0 = true
	var _162_v1 _dafny.Array
	_ = _162_v1
	var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
	_ = _nw20
	_162_v1 = _nw20
	var _163_v2 _dafny.Sequence
	_ = _163_v2
	_163_v2 = _dafny.UnicodeSeqOfUtf8Bytes("pa")
	var _164_globalState *GlobalState
	_ = _164_globalState
	var _nw21 *GlobalState = New_GlobalState_()
	_ = _nw21
	_nw21.Ctor__(true, _dafny.IntOfInt64(162), true, false, (func() _dafny.Array {
		if _161_v0 {
			return _162_v1
		}
		return _162_v1
	})(), true, _dafny.IntOfInt64(253), _163_v2, false)
	_164_globalState = _nw21
	var _165_v3 _dafny.Array
	_ = _165_v3
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(7)
	_ = _len0_2
	var _nw22 _dafny.Array
	_ = _nw22
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw22 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Int = (func(_166_v0 bool) func(_dafny.Int) _dafny.Int {
			return func(_167_i1 _dafny.Int) _dafny.Int {
				return (_167_i1).Plus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v0, _dafny.IntOfInt64(882))).Cardinality()))
			}
		})(_161_v0)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw22 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw22).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw22).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_165_v3 = _nw22
	for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_165_v3), 0))); ; {
		_guard_loop_0, _ok13 := _iter13()
		if !_ok13 {
			break
		}
		var _168_i0 _dafny.Int
		_168_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_168_i0).Sign() != -1) && ((_168_i0).Cmp(_dafny.ArrayLen((_165_v3), 0)) < 0)) {
			(_165_v3).ArraySet1(Companion_Default___.SafeDivisionInt(_168_i0, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _161_v0 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(105))).Uint32(), func(coer23 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
						return func(arg23 _dafny.Int) interface{} {
							return coer23(arg23)
						}
					}((func(_169_v0 bool) func(_dafny.Int) _dafny.MultiSet {
						return func(_170_i2 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf(_169_v0, _169_v0)
						}
					})(_161_v0)))
				}
				return _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_161_v0, _161_v0)), _dafny.MultiSetOf(_161_v0, _161_v0, _161_v0))
			})()).Cardinality())), (_168_i0).Int())
		}
	}
	var _171_v4 _dafny.Int
	_ = _171_v4
	_171_v4 = _dafny.IntOfInt64(490)
	var _hi0 _dafny.Int = Companion_Default___.Fm0(_164_globalState)
	_ = _hi0
	for _172_i3 := _171_v4; _172_i3.Cmp(_hi0) < 0; _172_i3 = _172_i3.Plus(_dafny.One) {
		Companion_Default___.M0(_161_v0, _164_globalState)
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_165_v3), 0))
		_ = _index13
		(_165_v3).ArraySet1((_dafny.Zero).Minus((_172_i3).Times(_172_i3)), (_index13).Int())
		var _173_v5 _dafny.Array
		_ = _173_v5
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_3
		var _nw23 _dafny.Array
		_ = _nw23
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw23 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = func(_174_i4 _dafny.Int) bool {
				return false
			}
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw23 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw23).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw23).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_173_v5 = _nw23
		var _175_v6 _dafny.CodePoint
		_ = _175_v6
		_175_v6 = _dafny.CodePoint('g')
		var _176_v7 _dafny.Map
		_ = _176_v7
		_176_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_173_v5, _175_v6)
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_165_v3), 0))
		_ = _index14
		var _rhs21 _dafny.Int = (_176_v7).Cardinality()
		_ = _rhs21
		var _rhs22 _dafny.Int = _172_i3
		_ = _rhs22
		var _lhs10 _dafny.Array = _165_v3
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_165_v3), 0))
		_ = _lhs11
		_171_v4 = _rhs21
		(_lhs10).ArraySet1(_rhs22, (_lhs11).Int())
		var _177_v8 _dafny.Map
		_ = _177_v8
		_177_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v4, (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int))
		var _178_v9 _dafny.Sequence
		_ = _178_v9
		_178_v9 = _dafny.SeqOf(_173_v5, _173_v5)
		_171_v4 = Companion_Default___.SafeDivisionInt((_177_v8).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_178_v9, _178_v9)).Cardinality()))
		var _179_v10 _dafny.MultiSet
		_ = _179_v10
		_179_v10 = _dafny.MultiSetOf((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int), (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int))
		_161_v0 = Companion_Default___.Fm1((func() _dafny.MultiSet {
			if _161_v0 {
				return _179_v10
			}
			return _179_v10
		})(), (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int), _164_globalState)
	}
	var _180_i5 _dafny.Int
	_ = _180_i5
	_180_i5 = _dafny.Zero
	{
		for _161_v0 {
			{
				if (_180_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_180_i5 = (_180_i5).Plus(_dafny.One)
				var _181_v11 _dafny.CodePoint
				_ = _181_v11
				_181_v11 = _dafny.CodePoint('j')
				var _182_v12 _dafny.Map
				_ = _182_v12
				_182_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v11, _171_v4), _161_v0)
				var _183_v13 _dafny.Sequence
				_ = _183_v13
				_183_v13 = _dafny.SeqOf(_163_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(655))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}((func(_184_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_185_i6 _dafny.Int) _dafny.CodePoint {
						return _184_v11
					}
				})(_181_v11))), _163_v2, _163_v2, _163_v2)
				var _186_v15 _dafny.Map
				_ = _186_v15
				_186_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v11, (func() _dafny.Set {
					var _coll13 = _dafny.NewBuilder()
					_ = _coll13
					for _iter14 := _dafny.Iterate((_183_v13).Elements()); ; {
						_compr_13, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _187_v14 _dafny.Sequence
						_187_v14 = interface{}(_compr_13).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_183_v13, _187_v14) {
							_coll13.Add(_187_v14)
						}
					}
					return _coll13.ToSet()
				}()).Cardinality())
				_182_v12 = (_182_v12).Update((_186_v15).Merge(_186_v15), _161_v0)
				var _188_v16 *C4
				_ = _188_v16
				var _nw24 *C4 = New_C4_()
				_ = _nw24
				_nw24.Ctor__(_161_v0)
				_188_v16 = _nw24
				(_164_globalState).F2 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm28(_171_v4, !(_161_v0), _164_globalState), _163_v2), (Companion_Default___.SafeIndex(_171_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm28(_171_v4, !(_161_v0), _164_globalState), _163_v2)).Cardinality()))).Uint32(), _181_v11), _dafny.SeqOf(_181_v11))
				if !(_161_v0) || (_161_v0) {
					var _189_v17 _dafny.Sequence
					_ = _189_v17
					_189_v17 = _dafny.SeqOf(_161_v0)
					var _190_v18 _dafny.MultiSet
					_ = _190_v18
					_190_v18 = _dafny.MultiSetOf((_dafny.MultiSetOf(_171_v4)).Cardinality(), _dafny.IntOfUint32((_189_v17).Cardinality()))
					var _191_v19 _dafny.Set
					_ = _191_v19
					_191_v19 = _dafny.SetOf(_dafny.SeqOf((_190_v18).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-165))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}((func(_192_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_193_i7 _dafny.Int) _dafny.Int {
							return _192_v4
						}
					})(_171_v4))))
					var _194_v20 *C0
					_ = _194_v20
					var _nw25 *C0 = New_C0_()
					_ = _nw25
					_nw25.Ctor__(_191_v19, (_dafny.Zero).Minus(_171_v4))
					_194_v20 = _nw25
					var _195_v21 D1
					_ = _195_v21
					_195_v21 = Companion_D1_.Create_DC2_(_194_v20.F12, (_194_v20.F12).Cmp(_171_v4) == 0, Companion_Default___.SafeModuloInt(_171_v4, _194_v20.F12), _171_v4)
					_195_v21 = _195_v21
					var _196_v22 _dafny.Map
					_ = _196_v22
					_196_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v4, _165_v3)
					var _rhs23 _dafny.Map = (_196_v22).Merge((_196_v22).Merge(_196_v22))
					_ = _rhs23
					var _rhs24 _dafny.Array = _162_v1
					_ = _rhs24
					_196_v22 = _rhs23
					_162_v1 = _rhs24
					var _197_v23 *C2
					_ = _197_v23
					var _nw26 *C2 = New_C2_()
					_ = _nw26
					_nw26.Ctor__(_171_v4, _161_v0)
					_197_v23 = _nw26
					_197_v23 = _197_v23
					_161_v0 = (_dafny.IntOfUint32((_163_v2).Cardinality())).Cmp(_171_v4) >= 0
				} else {
					var _198_v24 _dafny.Array
					_ = _198_v24
					var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(28))
					_ = _nw27
					_198_v24 = _nw27
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_198_v24), 0))
					_ = _index15
					(_198_v24).ArraySet1CodePoint((func() _dafny.CodePoint {
						if _161_v0 {
							return _181_v11
						}
						return _dafny.CodePoint('o')
					})(), (_index15).Int())
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_198_v24), 0))
					_ = _index16
					(_198_v24).ArraySet1CodePoint(_181_v11, (_index16).Int())
					var _rhs25 _dafny.Int = _171_v4
					_ = _rhs25
					var _rhs26 _dafny.Sequence = _163_v2
					_ = _rhs26
					_171_v4 = _rhs25
					_163_v2 = _rhs26
					var _199_v25 _dafny.Sequence
					_ = _199_v25
					_199_v25 = _dafny.SeqOf(_161_v0)
					var _200_v26 *C5
					_ = _200_v26
					var _nw28 *C5 = New_C5_()
					_ = _nw28
					_nw28.Ctor__(_dafny.IntOfUint32((_199_v25).Cardinality()), (_161_v0) || (!(_161_v0)))
					_200_v26 = _nw28
					var _201_v27 D1
					_ = _201_v27
					_201_v27 = Companion_D1_.Create_DC2_((_200_v26).F10(), true, _171_v4, (_dafny.Zero).Minus((_200_v26).F10()))
					var _202_v28 _dafny.Sequence
					_ = _202_v28
					_202_v28 = _dafny.SeqOf(_201_v27)
					var _203_v29 T0
					_ = _203_v29
					var _nw29 *C1 = New_C1_()
					_ = _nw29
					_nw29.Ctor__(_202_v28)
					_203_v29 = _nw29
					var _204_v30 _dafny.Map
					_ = _204_v30
					_204_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_203_v29, _171_v4)
					_204_v30 = (_204_v30).Update(_203_v29, _dafny.IntOfInt64(324))
					var _205_v31 _dafny.Sequence
					_ = _205_v31
					_205_v31 = _dafny.SeqOf((_200_v26).F10())
					var _206_v32 _dafny.Set
					_ = _206_v32
					_206_v32 = _dafny.SetOf(_205_v31)
					var _207_v33 _dafny.Map
					_ = _207_v33
					_207_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v26).F10(), _171_v4)
					var _208_v34 _dafny.MultiSet
					_ = _208_v34
					_208_v34 = _dafny.MultiSetOf(_165_v3, _165_v3)
					var _209_v35 *C0
					_ = _209_v35
					var _nw30 *C0 = New_C0_()
					_ = _nw30
					_nw30.Ctor__(_206_v32, (func() _dafny.Int {
						if (_207_v33).Contains(_dafny.IntOfInt64(901)) {
							return (_207_v33).Get(_dafny.IntOfInt64(901)).(_dafny.Int)
						}
						return (func() _dafny.Int {
							if (_208_v34).Contains(_165_v3) {
								return (_208_v34).Multiplicity(_165_v3)
							}
							return (_dafny.Zero).Minus(_171_v4)
						})()
					})())
					_209_v35 = _nw30
				}
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _hi1 _dafny.Int = _171_v4
	_ = _hi1
	for _210_i8 := _dafny.IntOfUint32((_163_v2).Cardinality()); _210_i8.Cmp(_hi1) < 0; _210_i8 = _210_i8.Plus(_dafny.One) {
		var _211_v36 _dafny.CodePoint
		_ = _211_v36
		_211_v36 = _dafny.CodePoint('n')
		var _212_v37 _dafny.MultiSet
		_ = _212_v37
		_212_v37 = _dafny.MultiSetOf(_211_v36, _211_v36, _211_v36)
		var _213_v38 _dafny.Sequence
		_ = _213_v38
		_213_v38 = _dafny.SeqOf(Companion_D1_.Create_DC2_(_210_i8, _161_v0, _171_v4, (_212_v37).Cardinality()))
		var _214_v39 T0
		_ = _214_v39
		var _nw31 *C1 = New_C1_()
		_ = _nw31
		_nw31.Ctor__(_dafny.Companion_Sequence_.Concatenate(_213_v38, _213_v38))
		_214_v39 = _nw31
		var _215_v41 D17
		_ = _215_v41
		_215_v41 = Companion_D17_.Create_DC48_(_214_v39)
		var _rhs27 bool = (Companion_Default___.Fm46(_161_v0, _161_v0, !(!(_161_v0)), _164_globalState)).Equals(func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-582), _dafny.IntOfInt64(489))); ; {
				_compr_14, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _216_v40 _dafny.Int
				_216_v40 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(-582)).Cmp(_216_v40) <= 0) && ((_216_v40).Cmp(_dafny.IntOfInt64(489)) < 0) {
					_coll14.Add(Companion_Default___.SafeModuloInt(_216_v40, _171_v4), _163_v2)
				}
			}
			return _coll14.ToMap()
		}())
		_ = _rhs27
		var _rhs28 _dafny.Int = _171_v4
		_ = _rhs28
		var _rhs29 T0 = (_215_v41).Dtor_cf81()
		_ = _rhs29
		var _lhs12 *GlobalState = _164_globalState
		_ = _lhs12
		_lhs12.F2 = _rhs27
		_171_v4 = _rhs28
		_214_v39 = _rhs29
		Companion_Default___.M0(_161_v0, _164_globalState)
		(_164_globalState).F2 = (_210_i8).Cmp(_171_v4) != 0
		var _217_v42 _dafny.Sequence
		_ = _217_v42
		_217_v42 = _dafny.SeqOf(_161_v0)
		var _rhs30 _dafny.Int = _210_i8
		_ = _rhs30
		var _rhs31 bool = _161_v0
		_ = _rhs31
		var _rhs32 bool = !(_161_v0) || ((func() bool {
			if _161_v0 {
				return (_217_v42).Select((Companion_Default___.SafeIndex(_171_v4, _dafny.IntOfUint32((_217_v42).Cardinality()))).Uint32()).(bool)
			}
			return _161_v0
		})())
		_ = _rhs32
		var _lhs13 *GlobalState = _164_globalState
		_ = _lhs13
		_171_v4 = _rhs30
		_lhs13.F2 = _rhs31
		_161_v0 = _rhs32
	}
	Companion_Default___.M0(_161_v0, _164_globalState)
	var _218_i9 _dafny.Int
	_ = _218_i9
	_218_i9 = _dafny.Zero
	{
		for _161_v0 {
			{
				if (_218_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_218_i9 = (_218_i9).Plus(_dafny.One)
				var _219_v43 *C2
				_ = _219_v43
				var _nw32 *C2 = New_C2_()
				_ = _nw32
				_nw32.Ctor__(_dafny.IntOfInt64(98), !(_161_v0))
				_219_v43 = _nw32
				var _220_v44 _dafny.Sequence
				_ = _220_v44
				_220_v44 = _dafny.SeqOf(_219_v43, _219_v43, _219_v43, _219_v43, _219_v43)
				_220_v44 = _220_v44
				var _221_v45 _dafny.MultiSet
				_ = _221_v45
				_221_v45 = _dafny.MultiSetOf(_171_v4)
				var _222_v47 _dafny.Array
				_ = _222_v47
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_4
				var _nw33 _dafny.Array
				_ = _nw33
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw33 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) _dafny.Set = (func(_223_v43 *C2) func(_dafny.Int) _dafny.Set {
						return func(_224_i10 _dafny.Int) _dafny.Set {
							return func() _dafny.Set {
								var _coll15 = _dafny.NewBuilder()
								_ = _coll15
								for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(809), _dafny.IntOfInt64(-805))); ; {
									_compr_15, _ok16 := _iter16()
									if !_ok16 {
										break
									}
									var _225_v46 _dafny.Int
									_225_v46 = interface{}(_compr_15).(_dafny.Int)
									if ((_dafny.IntOfInt64(809)).Cmp(_225_v46) <= 0) && ((_225_v46).Cmp(_dafny.IntOfInt64(-805)) < 0) {
										_coll15.Add((_225_v46).Minus((_223_v43).F14()))
									}
								}
								return _coll15.ToSet()
							}()
						}
					})(_219_v43)
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw33 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw33).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw33).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_222_v47 = _nw33
				var _226_v48 *C3
				_ = _226_v48
				var _nw34 *C3 = New_C3_()
				_ = _nw34
				_nw34.Ctor__(false, _222_v47, _161_v0)
				_226_v48 = _nw34
				var _227_v49 _dafny.Map
				_ = _227_v49
				_227_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v43).F14(), _226_v48)
				var _228_v50 _dafny.Map
				_ = _228_v50
				_228_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v0, _163_v2)
				var _229_v51 _dafny.Sequence
				_ = _229_v51
				_229_v51 = _dafny.SeqOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_228_v50).Contains((_226_v48).F15()) {
						return (_228_v50).Get((_226_v48).F15()).(_dafny.Sequence)
					}
					return _163_v2
				})()).Cardinality()), _171_v4)
				var _230_v52 _dafny.Sequence
				_ = _230_v52
				_230_v52 = _dafny.SeqOf(_dafny.IntOfUint32((_229_v51).Cardinality()))
				var _231_v53 D9
				_ = _231_v53
				_231_v53 = Companion_D9_.Create_DC24_(Companion_Default___.Fm1(_221_v45, (_227_v49).Cardinality(), _164_globalState), _dafny.IntOfUint32((_229_v51).Cardinality()), _163_v2)
				Companion_Default___.M0((_231_v53).Dtor_cf41(), _164_globalState)
				_163_v2 = _163_v2
				var _232_v54 *C3
				_ = _232_v54
				var _nw35 *C3 = New_C3_()
				_ = _nw35
				_nw35.Ctor__((_226_v48).F15(), _226_v48.F16, _161_v0)
				_232_v54 = _nw35
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_165_v3), 0))); ; {
		_guard_loop_1, _ok17 := _iter17()
		if !_ok17 {
			break
		}
		var _233_i11 _dafny.Int
		_233_i11 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_233_i11).Sign() != -1) && ((_233_i11).Cmp(_dafny.ArrayLen((_165_v3), 0)) < 0)) {
			(_165_v3).ArraySet1(Companion_Default___.SafeDivisionInt(_233_i11, _171_v4), (_233_i11).Int())
		}
	}
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))
	_ = _index17
	(_165_v3).ArraySet1(_dafny.IntOfInt64(77), (_index17).Int())
	var _234_v55 _dafny.CodePoint
	_ = _234_v55
	_234_v55 = _dafny.CodePoint('m')
	var _235_v56 D1
	_ = _235_v56
	_235_v56 = Companion_D1_.Create_DC2_(_171_v4, _161_v0, _171_v4, _dafny.IntOfUint32((_163_v2).Cardinality()))
	var _236_v57 _dafny.Sequence
	_ = _236_v57
	_236_v57 = _dafny.SeqOf(_235_v56)
	var _237_v58 T0
	_ = _237_v58
	var _nw36 *C1 = New_C1_()
	_ = _nw36
	_nw36.Ctor__(_236_v57)
	_237_v58 = _nw36
	var _238_v59 D17
	_ = _238_v59
	_238_v59 = Companion_D17_.Create_DC50_(_234_v55, _161_v0, _dafny.UnicodeSeqOfUtf8Bytes("awbtnvw"), _171_v4, _237_v58)
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))
	_ = _index18
	var _rhs33 _dafny.Int = _171_v4
	_ = _rhs33
	var _rhs34 D17 = _238_v59
	_ = _rhs34
	var _lhs14 _dafny.Array = _165_v3
	_ = _lhs14
	var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))
	_ = _lhs15
	(_lhs14).ArraySet1(_rhs33, (_lhs15).Int())
	_238_v59 = _rhs34
	var _239_v60 *C1
	_ = _239_v60
	var _nw37 *C1 = New_C1_()
	_ = _nw37
	_nw37.Ctor__(_236_v57)
	_239_v60 = _nw37
	var _240_v61 D13
	_ = _240_v61
	_240_v61 = Companion_D13_.Create_DC35_(_161_v0, _234_v55, _239_v60, _161_v0)
	var _241_v62 _dafny.Array
	_ = _241_v62
	var _nwElement0_4 *C1 = _239_v60
	_ = _nwElement0_4
	var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(25))
	_ = _nw38
	(_nw38).ArraySet1(_nwElement0_4, 0)
	(_nw38).ArraySet1(_239_v60, 1)
	(_nw38).ArraySet1(_239_v60, 2)
	(_nw38).ArraySet1(_239_v60, 3)
	(_nw38).ArraySet1(_239_v60, 4)
	(_nw38).ArraySet1(_239_v60, 5)
	(_nw38).ArraySet1(_239_v60, 6)
	(_nw38).ArraySet1(_239_v60, 7)
	(_nw38).ArraySet1(_239_v60, 8)
	(_nw38).ArraySet1(_239_v60, 9)
	(_nw38).ArraySet1(_239_v60, 10)
	(_nw38).ArraySet1(_239_v60, 11)
	(_nw38).ArraySet1(_239_v60, 12)
	(_nw38).ArraySet1(_239_v60, 13)
	(_nw38).ArraySet1(_239_v60, 14)
	(_nw38).ArraySet1((func() *C1 {
		if false {
			return _239_v60
		}
		return _239_v60
	})(), 15)
	(_nw38).ArraySet1(_239_v60, 16)
	(_nw38).ArraySet1(_239_v60, 17)
	(_nw38).ArraySet1((_240_v61).Dtor_cf60(), 18)
	(_nw38).ArraySet1(_239_v60, 19)
	(_nw38).ArraySet1(_239_v60, 20)
	(_nw38).ArraySet1(_239_v60, 21)
	(_nw38).ArraySet1(_239_v60, 22)
	(_nw38).ArraySet1(_239_v60, 23)
	(_nw38).ArraySet1(_239_v60, 24)
	_241_v62 = _nw38
	_241_v62 = _241_v62
	var _242_v63 _dafny.Sequence
	_ = _242_v63
	_242_v63 = _dafny.SeqOf(_165_v3)
	var _243_i12 _dafny.Int
	_ = _243_i12
	_243_i12 = _dafny.Zero
	{
		for (func() bool {
			if _161_v0 {
				return false
			}
			return (_dafny.IntOfUint32((_242_v63).Cardinality())).Cmp(Companion_Default___.Fm0(_164_globalState)) <= 0
		})() {
			{
				if (_243_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L4
				}
				_243_i12 = (_243_i12).Plus(_dafny.One)
				var _244_v64 _dafny.Array
				_ = _244_v64
				var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
				_ = _nw39
				_244_v64 = _nw39
				_244_v64 = _244_v64
				var _245_v65 D16
				_ = _245_v65
				_245_v65 = Companion_D16_.Create_DC46_(_171_v4)
				var _246_v66 D16
				_ = _246_v66
				_246_v66 = Companion_D16_.Create_DC47_(_245_v65)
				var _247_v67 D16
				_ = _247_v67
				_247_v67 = Companion_D16_.Create_DC47_(_246_v66)
				var _248_v68 _dafny.Map
				_ = _248_v68
				_248_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v4, _247_v67)
				_248_v68 = (_248_v68).Update(_171_v4, _247_v67)
				if !(_161_v0) {
					var _249_v69 _dafny.Map
					_ = _249_v69
					_249_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v0, _161_v0)
					_249_v69 = (_249_v69).Update(_161_v0, _161_v0)
					var _250_v70 _dafny.Sequence
					_ = _250_v70
					_250_v70 = _dafny.SeqOf((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int))
					var _251_v71 _dafny.Set
					_ = _251_v71
					_251_v71 = _dafny.SetOf(_250_v70, _dafny.SeqOf(_dafny.IntOfInt64(948), (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int)), _250_v70)
					var _252_v72 *C0
					_ = _252_v72
					var _nw40 *C0 = New_C0_()
					_ = _nw40
					_nw40.Ctor__(_251_v71, ((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int)).Plus((_249_v69).Cardinality()))
					_252_v72 = _nw40
					var _253_v73 _dafny.Map
					_ = _253_v73
					_253_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(Companion_Default___.Fm0(_164_globalState)), _161_v0)
					var _254_v74 D7
					_ = _254_v74
					_254_v74 = Companion_D7_.Create_DC18_(_253_v73)
					var _255_v75 _dafny.Map
					_ = _255_v75
					_255_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v4, _254_v74)
					_255_v75 = (_255_v75).Update((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int), _254_v74)
					var _pat_let_tv2 = _164_globalState
					_ = _pat_let_tv2
					_254_v74 = func(_pat_let2_0 D7) D7 {
						return func(_256_dt__update__tmp_h0 D7) D7 {
							return func(_pat_let3_0 _dafny.Map) D7 {
								return func(_257_dt__update_hcf32_h0 _dafny.Map) D7 {
									return Companion_D7_.Create_DC18_(_257_dt__update_hcf32_h0)
								}(_pat_let3_0)
							}(Companion_Default___.Fm22(_pat_let_tv2))
						}(_pat_let2_0)
					}(_254_v74)
					(_252_v72).F12 = _171_v4
				} else {
					_161_v0 = _161_v0
					var _258_v76 _dafny.Array
					_ = _258_v76
					var _len0_5 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_5
					var _nw41 _dafny.Array
					_ = _nw41
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw41 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) _dafny.Map = (func(_259_v3 _dafny.Array, _260_v4 _dafny.Int) func(_dafny.Int) _dafny.Map {
							return func(_261_i13 _dafny.Int) _dafny.Map {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_259_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_259_v3), 0))).Int()).(_dafny.Int)), (_259_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_259_v3), 0))).Int()).(_dafny.Int))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(694))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg26 _dafny.Int) interface{} {
										return coer26(arg26)
									}
								}((func(_262_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
									return func(_263_i14 _dafny.Int) _dafny.Int {
										return _262_v4
									}
								})(_260_v4))), (_dafny.SetOf((_259_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_259_v3), 0))).Int()).(_dafny.Int))).Cardinality()))
							}
						})(_165_v3, _171_v4)
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw41 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw41).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw41).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_258_v76 = _nw41
					var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_258_v76), 0))
					_ = _index19
					(_258_v76).ArraySet1(Companion_Default___.Fm47(_164_globalState), (_index19).Int())
					var _264_v77 _dafny.Sequence
					_ = _264_v77
					_264_v77 = _dafny.SeqOf(_171_v4, (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int))
					var _265_v78 _dafny.Sequence
					_ = _265_v78
					_265_v78 = _264_v77
					var _266_v79 _dafny.Map
					_ = _266_v79
					_266_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_v78, _171_v4)
					var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_258_v76), 0))
					_ = _index20
					(_258_v76).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_v78, _171_v4)).Merge((_266_v79).Update(_264_v77, _dafny.IntOfUint32((_264_v77).Cardinality())))).Update(_265_v78, _171_v4), (_index20).Int())
					var _267_v80 _dafny.MultiSet
					_ = _267_v80
					_267_v80 = _dafny.MultiSetOf((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int))
					var _268_v81 _dafny.Array
					_ = _268_v81
					_268_v81 = _165_v3
					var _rhs35 bool = _161_v0
					_ = _rhs35
					var _rhs36 bool = _161_v0
					_ = _rhs36
					var _rhs37 bool = ((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int)).Cmp((_267_v80).Cardinality()) <= 0
					_ = _rhs37
					var _rhs38 bool = _161_v0
					_ = _rhs38
					var _rhs39 _dafny.Array = (_268_v81)
					_ = _rhs39
					var _lhs16 *GlobalState = _164_globalState
					_ = _lhs16
					var _lhs17 *GlobalState = _164_globalState
					_ = _lhs17
					var _lhs18 *GlobalState = _164_globalState
					_ = _lhs18
					_lhs16.F2 = _rhs35
					_lhs17.F2 = _rhs36
					_161_v0 = _rhs37
					_lhs18.F2 = _rhs38
					_165_v3 = _rhs39
					var _269_v82 D15
					_ = _269_v82
					_269_v82 = Companion_D15_.Create_DC43_()
					var _270_v83 D15
					_ = _270_v83
					_270_v83 = Companion_D15_.Create_DC44_(_269_v82)
					var _271_v84 T1
					_ = _271_v84
					var _nw42 *C2 = New_C2_()
					_ = _nw42
					_nw42.Ctor__(_dafny.IntOfUint32((_163_v2).Cardinality()), !(true))
					_271_v84 = _nw42
					var _272_v85 _dafny.Sequence
					_ = _272_v85
					_272_v85 = _dafny.SeqOf(_271_v84, _271_v84)
					var _273_v86 *C2
					_ = _273_v86
					var _nw43 *C2 = New_C2_()
					_ = _nw43
					_nw43.Ctor__(Companion_Default___.Fm0(_164_globalState), true)
					_273_v86 = _nw43
					var _274_v87 _dafny.Map
					_ = _274_v87
					_274_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_272_v85, _272_v85)).Cardinality()), _273_v86)
					var _275_v88 _dafny.Array
					_ = _275_v88
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(3)
					_ = _len0_6
					var _nw44 _dafny.Array
					_ = _nw44
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw44 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.Set = (func(_276_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
							return func(_277_i15 _dafny.Int) _dafny.Set {
								return _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_276_v2).Cardinality())))
							}
						})(_163_v2)
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw44 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw44).ArraySet1(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw44).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_275_v88 = _nw44
					var _278_v89 T2
					_ = _278_v89
					var _nw45 *C3 = New_C3_()
					_ = _nw45
					_nw45.Ctor__(false, _275_v88, (_271_v84).F9())
					_278_v89 = _nw45
					var _279_v90 D15
					_ = _279_v90
					_279_v90 = Companion_D15_.Create_DC42_(_161_v0, (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int), _278_v89, _171_v4)
					var _280_v91 _dafny.Map
					_ = _280_v91
					_280_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v90, ((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus(_171_v4)) < 0)
					var _pat_let_tv3 = _269_v82
					_ = _pat_let_tv3
					var _rhs40 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_171_v4, ((_273_v86).F14()).Times((_273_v86).F14()))))
					_ = _rhs40
					var _rhs41 D15 = func(_pat_let4_0 D15) D15 {
						return func(_281_dt__update__tmp_h2 D15) D15 {
							return func(_pat_let5_0 D15) D15 {
								return func(_282_dt__update_hcf77_h0 D15) D15 {
									return Companion_D15_.Create_DC44_(_282_dt__update_hcf77_h0)
								}(_pat_let5_0)
							}(_pat_let_tv3)
						}(_pat_let4_0)
					}(_270_v83)
					_ = _rhs41
					var _rhs42 bool = (func() bool {
						if (_280_v91).Contains(_279_v90) {
							return (_280_v91).Get(_279_v90).(bool)
						}
						return (_278_v89).F9()
					})()
					_ = _rhs42
					var _rhs43 _dafny.Map = (_274_v87).Merge(_274_v87)
					_ = _rhs43
					var _lhs19 *GlobalState = _164_globalState
					_ = _lhs19
					_171_v4 = _rhs40
					_270_v83 = _rhs41
					_lhs19.F2 = _rhs42
					_274_v87 = _rhs43
					var _283_v92 _dafny.Array
					_ = _283_v92
					var _nwElement0_5 bool = (_278_v89).F9()
					_ = _nwElement0_5
					var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(2))
					_ = _nw46
					(_nw46).ArraySet1(_nwElement0_5, 0)
					(_nw46).ArraySet1((_278_v89).F9(), 1)
					_283_v92 = _nw46
					var _284_v93 bool
					_ = _284_v93
					var _285_v94 bool
					_ = _285_v94
					var _out8 bool
					_ = _out8
					var _out9 bool
					_ = _out9
					_out8, _out9 = (_239_v60).M6((func() T0 {
						if _161_v0 {
							return _237_v58
						}
						return _237_v58
					})(), _283_v92, (_271_v84).F9(), _164_globalState)
					_284_v93 = _out8
					_285_v94 = _out9
				}
				_171_v4 = (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int)
				goto C4
			}
		C4:
		}
		goto L4
	}
L4:
	var _286_v95 _dafny.Map
	_ = _286_v95
	_286_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v4, ((Companion_Default___.Fm23(_161_v0, _161_v0, _171_v4, _164_globalState)).Cardinality()).Plus(_171_v4))
	var _287_v96 _dafny.Map
	_ = _287_v96
	_287_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v4, _161_v0)
	var _288_v97 D4
	_ = _288_v97
	_288_v97 = Companion_D4_.Create_DC13_(_171_v4)
	var _289_v98 *C2
	_ = _289_v98
	var _nw47 *C2 = New_C2_()
	_ = _nw47
	_nw47.Ctor__((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int), _161_v0)
	_289_v98 = _nw47
	var _290_v99 _dafny.MultiSet
	_ = _290_v99
	_290_v99 = _dafny.MultiSetOf(_289_v98)
	var _291_v100 _dafny.Sequence
	_ = _291_v100
	_291_v100 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v0, (_289_v98).F14()))
	var _292_v101 _dafny.Sequence
	_ = _292_v101
	_292_v101 = _dafny.SeqOf((_291_v100).Select((Companion_Default___.SafeIndex((_289_v98).F14(), _dafny.IntOfUint32((_291_v100).Cardinality()))).Uint32()).(_dafny.Map))
	var _293_v102 _dafny.Map
	_ = _293_v102
	_293_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v0, (_dafny.SetOf(_161_v0, _161_v0)).Cardinality())
	var _294_v103 _dafny.Map
	_ = _294_v103
	_294_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v0, _161_v0)
	var _295_v104 D9
	_ = _295_v104
	_295_v104 = Companion_D9_.Create_DC24_(_161_v0, (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("fqe"))
	var _pat_let_tv4 = _165_v3
	_ = _pat_let_tv4
	var _pat_let_tv5 = _165_v3
	_ = _pat_let_tv5
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))
	_ = _index21
	var _rhs44 _dafny.Map = (_286_v95).Merge((Companion_Default___.Fm48(_161_v0, Companion_Default___.Fm11(false, _164_globalState), _161_v0, (func() _dafny.Int {
		if (_286_v95).Contains((_dafny.Zero).Minus((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int))) {
			return (_286_v95).Get((_dafny.Zero).Minus((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int))).(_dafny.Int)
		}
		return (_287_v96).Cardinality()
	})(), _164_globalState)).Update(_dafny.IntOfInt64(-614), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let6_0 D4) D4 {
		return func(_296_dt__update__tmp_h3 D4) D4 {
			return func(_pat_let7_0 _dafny.Int) D4 {
				return func(_297_dt__update_hcf25_h0 _dafny.Int) D4 {
					return Companion_D4_.Create_DC13_(_297_dt__update_hcf25_h0)
				}(_pat_let7_0)
			}((_pat_let_tv5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_pat_let_tv4), 0))).Int()).(_dafny.Int))
		}(_pat_let6_0)
	}(_288_v97), _237_v58)).Cardinality()))
	_ = _rhs44
	var _rhs45 _dafny.Int = (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int)
	_ = _rhs45
	var _rhs46 _dafny.CodePoint = _234_v55
	_ = _rhs46
	var _rhs47 _dafny.Int = (func() _dafny.Int {
		if (_290_v99).Contains(_289_v98) {
			return (_290_v99).Multiplicity(_289_v98)
		}
		return ((func() _dafny.Map {
			if false {
				return (_291_v100).Select((Companion_Default___.SafeIndex(_171_v4, _dafny.IntOfUint32((_291_v100).Cardinality()))).Uint32()).(_dafny.Map)
			}
			return _293_v102
		})()).Cardinality()
	})()
	_ = _rhs47
	var _rhs48 bool = !((func() bool {
		if _161_v0 {
			return false
		}
		return !((_295_v104).Dtor_cf41())
	})()) || (((_294_v103).Cardinality()).Cmp((_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int)) > 0)
	_ = _rhs48
	var _lhs20 _dafny.Array = _165_v3
	_ = _lhs20
	var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))
	_ = _lhs21
	var _lhs22 *GlobalState = _164_globalState
	_ = _lhs22
	_286_v95 = _rhs44
	_171_v4 = _rhs45
	_234_v55 = _rhs46
	(_lhs20).ArraySet1(_rhs47, (_lhs21).Int())
	_lhs22.F2 = _rhs48
	var _298_v105 _dafny.Array
	_ = _298_v105
	var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
	_ = _nw48
	_298_v105 = _nw48
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))
	_ = _index22
	var _rhs49 _dafny.Array = _298_v105
	_ = _rhs49
	var _rhs50 _dafny.Int = (_165_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))).Int()).(_dafny.Int)
	_ = _rhs50
	var _rhs51 D4 = Companion_D4_.Create_DC13_(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_163_v2).Cardinality()), (_289_v98).F14()))
	_ = _rhs51
	var _rhs52 bool = !(_161_v0)
	_ = _rhs52
	var _lhs23 _dafny.Array = _165_v3
	_ = _lhs23
	var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))
	_ = _lhs24
	_298_v105 = _rhs49
	(_lhs23).ArraySet1(_rhs50, (_lhs24).Int())
	_288_v97 = _rhs51
	_161_v0 = _rhs52
	var _299_v106 _dafny.Array
	_ = _299_v106
	var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
	_ = _nw49
	_299_v106 = _nw49
	_299_v106 = _299_v106
	var _300_v107 D15
	_ = _300_v107
	_300_v107 = Companion_D15_.Create_DC41_()
	var _301_v108 D15
	_ = _301_v108
	_301_v108 = Companion_D15_.Create_DC44_(_300_v107)
	var _302_v109 _dafny.Set
	_ = _302_v109
	_302_v109 = _dafny.SetOf(_161_v0, _161_v0, _161_v0, _161_v0)
	var _303_v110 _dafny.Map
	_ = _303_v110
	_303_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_v108, (_302_v109).Difference(_302_v109))
	_303_v110 = (_303_v110).Merge(_303_v110)
	var _304_v111 D17
	_ = _304_v111
	_304_v111 = Companion_D17_.Create_DC48_((func() T0 {
		if _161_v0 {
			return _237_v58
		}
		return _237_v58
	})())
	_304_v111 = _304_v111
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_165_v3), 0))
	_ = _index23
	(_165_v3).ArraySet1(_171_v4, (_index23).Int())
	_dafny.Print(_161_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_163_v2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_164_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_globalState).F7().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v3).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_171_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_180_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_218_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_234_v55)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v56).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v56).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v56).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v56).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_236_v57, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v59).Dtor_cf82())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v59).Dtor_cf83())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v59).Dtor_cf84().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v59).Dtor_cf85())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_239_v60.F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_v61).Dtor_cf58())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_v61).Dtor_cf59())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_240_v61).Dtor_cf60().F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_v61).Dtor_cf61())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.Zero).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.One).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_241_v62).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(*C1).F13, _dafny.SeqOf(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(489), true, _dafny.IntOfInt64(489), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_242_v63).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_243_i12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v95).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(489), _dafny.IntOfInt64(491)).UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.IntOfInt64(849)).UpdateUnsafe(_dafny.IntOfInt64(-614), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_287_v96).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(489), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_v97).Dtor_cf25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_v98).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v99).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_291_v100, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(489)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_292_v101, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(489)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_293_v102).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_294_v103).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_295_v104).Dtor_cf41())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_295_v104).Dtor_cf42())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_295_v104).Dtor_cf43().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_302_v109).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_303_v110).Cardinality())
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

func (CompanionStruct_D0_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
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
	Cf3 bool
	Cf4 _dafny.Int
	Cf5 _dafny.Int
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int, Cf3 bool, Cf4 _dafny.Int, Cf5 _dafny.Int) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4, Cf5}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC3 struct {
	Cf6 bool
	Cf7 _dafny.Sequence
	Cf8 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 bool, Cf7 _dafny.Sequence, Cf8 _dafny.Int) D1 {
	return D1{D1_DC3{Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf9  _dafny.Int
	Cf10 bool
	Cf11 bool
	Cf12 bool
	Cf13 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf9 _dafny.Int, Cf10 bool, Cf11 bool, Cf12 bool, Cf13 bool) D1 {
	return D1{D1_DC4{Cf9, Cf10, Cf11, Cf12, Cf13}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

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

type D1_DC5 struct {
	Cf14 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf14 D1) D1 {
	return D1{D1_DC5{Cf14}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, false, _dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() bool {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf12() bool {
	return _this.Get_().(D1_DC4).Cf12
}

func (_this D1) Dtor_cf13() bool {
	return _this.Get_().(D1_DC4).Cf13
}

func (_this D1) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) Dtor_cf14() D1 {
	return _this.Get_().(D1_DC5).Cf14
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf6) + ", " + data.Cf7.VerbatimString(true) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7.Equals(data2.Cf7) && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf14.Equals(data2.Cf14)
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
	Cf15 _dafny.Set
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf15 _dafny.Set) D2 {
	return D2{D2_DC6{Cf15}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D2) Dtor_cf15() _dafny.Set {
	return _this.Get_().(D2_DC6).Cf15
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf15.Equals(data2.Cf15)
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
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_() D3 {
	return D3{D3_DC8{}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf17 bool
	Cf18 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf17 bool, Cf18 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf17, Cf18}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

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

type D3_DC7 struct {
	Cf16 _dafny.MultiSet
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf16 _dafny.MultiSet) D3 {
	return D3{D3_DC7{Cf16}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_()
}

func (_this D3) Dtor_cf17() bool {
	return _this.Get_().(D3_DC9).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf18
}

func (_this D3) Dtor_cf16() _dafny.MultiSet {
	return _this.Get_().(D3_DC7).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf16) + ")"
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
			_, ok := other.Get_().(D3_DC8)
			return ok
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Equals(data2.Cf18)
		}
	case D3_DC10:
		{
			_, ok := other.Get_().(D3_DC10)
			return ok
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
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

type D4_DC12 struct {
	Cf20 bool
	Cf21 bool
	Cf22 _dafny.Int
	Cf23 _dafny.Int
	Cf24 bool
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf20 bool, Cf21 bool, Cf22 _dafny.Int, Cf23 _dafny.Int, Cf24 bool) D4 {
	return D4{D4_DC12{Cf20, Cf21, Cf22, Cf23, Cf24}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf25 _dafny.Int
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf25 _dafny.Int) D4 {
	return D4{D4_DC13{Cf25}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf26 bool
	Cf27 bool
	Cf28 _dafny.Int
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf26 bool, Cf27 bool, Cf28 _dafny.Int) D4 {
	return D4{D4_DC14{Cf26, Cf27, Cf28}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC11 struct {
	Cf19 _dafny.Array
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf19 _dafny.Array) D4 {
	return D4{D4_DC11{Cf19}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC15 struct {
	Cf29 D4
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf29 D4) D4 {
	return D4{D4_DC15{Cf29}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(false, false, _dafny.Zero, _dafny.Zero, false)
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC12).Cf20
}

func (_this D4) Dtor_cf21() bool {
	return _this.Get_().(D4_DC12).Cf21
}

func (_this D4) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf23
}

func (_this D4) Dtor_cf24() bool {
	return _this.Get_().(D4_DC12).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf25
}

func (_this D4) Dtor_cf26() bool {
	return _this.Get_().(D4_DC14).Cf26
}

func (_this D4) Dtor_cf27() bool {
	return _this.Get_().(D4_DC14).Cf27
}

func (_this D4) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf28
}

func (_this D4) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D4_DC11).Cf19
}

func (_this D4) Dtor_cf29() D4 {
	return _this.Get_().(D4_DC15).Cf29
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf29) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21 && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24 == data2.Cf24
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28.Cmp(data2.Cf28) == 0
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf19 == data2.Cf19
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf29.Equals(data2.Cf29)
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
	Cf30 _dafny.Array
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf30 _dafny.Array) D5 {
	return D5{D5_DC16{Cf30}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf30() _dafny.Array {
	return _this.Get_().(D5_DC16).Cf30
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf30) + ")"
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
			return ok && data1.Cf30 == data2.Cf30
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
	Cf31 _dafny.MultiSet
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf31 _dafny.MultiSet) D6 {
	return D6{D6_DC17{Cf31}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D6) Dtor_cf31() _dafny.MultiSet {
	return _this.Get_().(D6_DC17).Cf31
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf31.Equals(data2.Cf31)
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
	Cf33 _dafny.Array
	Cf34 bool
	Cf35 _dafny.Sequence
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf33 _dafny.Array, Cf34 bool, Cf35 _dafny.Sequence) D7 {
	return D7{D7_DC19{Cf33, Cf34, Cf35}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC20 struct {
	Cf36 bool
	Cf37 _dafny.CodePoint
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf36 bool, Cf37 _dafny.CodePoint) D7 {
	return D7{D7_DC20{Cf36, Cf37}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC18 struct {
	Cf32 _dafny.Map
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf32 _dafny.Map) D7 {
	return D7{D7_DC18{Cf32}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC21 struct {
	Cf38 D7
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf38 D7) D7 {
	return D7{D7_DC21{Cf38}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.EmptySeq)
}

func (_this D7) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D7_DC19).Cf33
}

func (_this D7) Dtor_cf34() bool {
	return _this.Get_().(D7_DC19).Cf34
}

func (_this D7) Dtor_cf35() _dafny.Sequence {
	return _this.Get_().(D7_DC19).Cf35
}

func (_this D7) Dtor_cf36() bool {
	return _this.Get_().(D7_DC20).Cf36
}

func (_this D7) Dtor_cf37() _dafny.CodePoint {
	return _this.Get_().(D7_DC20).Cf37
}

func (_this D7) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D7_DC18).Cf32
}

func (_this D7) Dtor_cf38() D7 {
	return _this.Get_().(D7_DC21).Cf38
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + data.Cf35.VerbatimString(true) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf38) + ")"
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
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf33 == data2.Cf33 && data1.Cf34 == data2.Cf34 && data1.Cf35.Equals(data2.Cf35)
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf32.Equals(data2.Cf32)
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf38.Equals(data2.Cf38)
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
	Cf39 _dafny.Array
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf39 _dafny.Array) D8 {
	return D8{D8_DC22{Cf39}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D8) Dtor_cf39() _dafny.Array {
	return _this.Get_().(D8_DC22).Cf39
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf39 == data2.Cf39
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

type D9_DC24 struct {
	Cf41 bool
	Cf42 _dafny.Int
	Cf43 _dafny.Sequence
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf41 bool, Cf42 _dafny.Int, Cf43 _dafny.Sequence) D9 {
	return D9{D9_DC24{Cf41, Cf42, Cf43}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC25 struct {
	Cf44 _dafny.Int
	Cf45 _dafny.Int
	Cf46 _dafny.MultiSet
	Cf47 _dafny.Map
	Cf48 _dafny.Int
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf44 _dafny.Int, Cf45 _dafny.Int, Cf46 _dafny.MultiSet, Cf47 _dafny.Map, Cf48 _dafny.Int) D9 {
	return D9{D9_DC25{Cf44, Cf45, Cf46, Cf47, Cf48}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC23 struct {
	Cf40 _dafny.Map
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf40 _dafny.Map) D9 {
	return D9{D9_DC23{Cf40}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC26 struct {
	Cf49 D9
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf49 D9) D9 {
	return D9{D9_DC26{Cf49}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC24_(false, _dafny.Zero, _dafny.EmptySeq)
}

func (_this D9) Dtor_cf41() bool {
	return _this.Get_().(D9_DC24).Cf41
}

func (_this D9) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf42
}

func (_this D9) Dtor_cf43() _dafny.Sequence {
	return _this.Get_().(D9_DC24).Cf43
}

func (_this D9) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf44
}

func (_this D9) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf45
}

func (_this D9) Dtor_cf46() _dafny.MultiSet {
	return _this.Get_().(D9_DC25).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Map {
	return _this.Get_().(D9_DC25).Cf47
}

func (_this D9) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf48
}

func (_this D9) Dtor_cf40() _dafny.Map {
	return _this.Get_().(D9_DC23).Cf40
}

func (_this D9) Dtor_cf49() D9 {
	return _this.Get_().(D9_DC26).Cf49
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ", " + data.Cf43.VerbatimString(true) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf49) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf41 == data2.Cf41 && data1.Cf42.Cmp(data2.Cf42) == 0 && data1.Cf43.Equals(data2.Cf43)
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46.Equals(data2.Cf46) && data1.Cf47.Equals(data2.Cf47) && data1.Cf48.Cmp(data2.Cf48) == 0
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf40.Equals(data2.Cf40)
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D10_DC28 struct {
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_() D10 {
	return D10{D10_DC28{}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

type D10_DC27 struct {
	Cf50 _dafny.Map
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf50 _dafny.Map) D10 {
	return D10{D10_DC27{Cf50}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC28_()
}

func (_this D10) Dtor_cf50() _dafny.Map {
	return _this.Get_().(D10_DC27).Cf50
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC28:
		{
			return "D10.DC28"
		}
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC28:
		{
			_, ok := other.Get_().(D10_DC28)
			return ok
		}
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf50.Equals(data2.Cf50)
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

type D11_DC30 struct {
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_() D11 {
	return D11{D11_DC30{}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC29 struct {
	Cf51 _dafny.Map
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf51 _dafny.Map) D11 {
	return D11{D11_DC29{Cf51}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC30_()
}

func (_this D11) Dtor_cf51() _dafny.Map {
	return _this.Get_().(D11_DC29).Cf51
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC30:
		{
			_, ok := other.Get_().(D11_DC30)
			return ok
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
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

type D12_DC32 struct {
	Cf53 _dafny.Int
	Cf54 bool
	Cf55 _dafny.Int
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf53 _dafny.Int, Cf54 bool, Cf55 _dafny.Int) D12 {
	return D12{D12_DC32{Cf53, Cf54, Cf55}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC31 struct {
	Cf52 _dafny.Sequence
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf52 _dafny.Sequence) D12 {
	return D12{D12_DC31{Cf52}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC33 struct {
	Cf56 D12
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf56 D12) D12 {
	return D12{D12_DC33{Cf56}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC32_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D12) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf53
}

func (_this D12) Dtor_cf54() bool {
	return _this.Get_().(D12_DC32).Cf54
}

func (_this D12) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf55
}

func (_this D12) Dtor_cf52() _dafny.Sequence {
	return _this.Get_().(D12_DC31).Cf52
}

func (_this D12) Dtor_cf56() D12 {
	return _this.Get_().(D12_DC33).Cf56
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf52) + ")"
		}
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf53.Cmp(data2.Cf53) == 0 && data1.Cf54 == data2.Cf54 && data1.Cf55.Cmp(data2.Cf55) == 0
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf52.Equals(data2.Cf52)
		}
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
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

type D13_DC35 struct {
	Cf58 bool
	Cf59 _dafny.CodePoint
	Cf60 *C1
	Cf61 bool
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf58 bool, Cf59 _dafny.CodePoint, Cf60 *C1, Cf61 bool) D13 {
	return D13{D13_DC35{Cf58, Cf59, Cf60, Cf61}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

type D13_DC36 struct {
	Cf62 _dafny.Sequence
	Cf63 _dafny.Int
	Cf64 _dafny.Int
	Cf65 _dafny.Int
}

func (D13_DC36) isD13() {}

func (CompanionStruct_D13_) Create_DC36_(Cf62 _dafny.Sequence, Cf63 _dafny.Int, Cf64 _dafny.Int, Cf65 _dafny.Int) D13 {
	return D13{D13_DC36{Cf62, Cf63, Cf64, Cf65}}
}

func (_this D13) Is_DC36() bool {
	_, ok := _this.Get_().(D13_DC36)
	return ok
}

type D13_DC34 struct {
	Cf57 _dafny.Array
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf57 _dafny.Array) D13 {
	return D13{D13_DC34{Cf57}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC35_(false, _dafny.CodePoint('D'), (*C1)(nil), false)
}

func (_this D13) Dtor_cf58() bool {
	return _this.Get_().(D13_DC35).Cf58
}

func (_this D13) Dtor_cf59() _dafny.CodePoint {
	return _this.Get_().(D13_DC35).Cf59
}

func (_this D13) Dtor_cf60() *C1 {
	return _this.Get_().(D13_DC35).Cf60
}

func (_this D13) Dtor_cf61() bool {
	return _this.Get_().(D13_DC35).Cf61
}

func (_this D13) Dtor_cf62() _dafny.Sequence {
	return _this.Get_().(D13_DC36).Cf62
}

func (_this D13) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D13_DC36).Cf63
}

func (_this D13) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D13_DC36).Cf64
}

func (_this D13) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D13_DC36).Cf65
}

func (_this D13) Dtor_cf57() _dafny.Array {
	return _this.Get_().(D13_DC34).Cf57
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D13_DC36:
		{
			return "D13.DC36" + "(" + data.Cf62.VerbatimString(true) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf57) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC35:
		{
			data2, ok := other.Get_().(D13_DC35)
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59 == data2.Cf59 && data1.Cf60 == data2.Cf60 && data1.Cf61 == data2.Cf61
		}
	case D13_DC36:
		{
			data2, ok := other.Get_().(D13_DC36)
			return ok && data1.Cf62.Equals(data2.Cf62) && data1.Cf63.Cmp(data2.Cf63) == 0 && data1.Cf64.Cmp(data2.Cf64) == 0 && data1.Cf65.Cmp(data2.Cf65) == 0
		}
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf57 == data2.Cf57
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
	Cf67 _dafny.CodePoint
	Cf68 bool
	Cf69 _dafny.Sequence
}

func (D14_DC38) isD14() {}

func (CompanionStruct_D14_) Create_DC38_(Cf67 _dafny.CodePoint, Cf68 bool, Cf69 _dafny.Sequence) D14 {
	return D14{D14_DC38{Cf67, Cf68, Cf69}}
}

func (_this D14) Is_DC38() bool {
	_, ok := _this.Get_().(D14_DC38)
	return ok
}

type D14_DC39 struct {
	Cf70 _dafny.Int
	Cf71 _dafny.Map
}

func (D14_DC39) isD14() {}

func (CompanionStruct_D14_) Create_DC39_(Cf70 _dafny.Int, Cf71 _dafny.Map) D14 {
	return D14{D14_DC39{Cf70, Cf71}}
}

func (_this D14) Is_DC39() bool {
	_, ok := _this.Get_().(D14_DC39)
	return ok
}

type D14_DC37 struct {
	Cf66 _dafny.Sequence
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf66 _dafny.Sequence) D14 {
	return D14{D14_DC37{Cf66}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC38_(_dafny.CodePoint('D'), false, _dafny.EmptySeq)
}

func (_this D14) Dtor_cf67() _dafny.CodePoint {
	return _this.Get_().(D14_DC38).Cf67
}

func (_this D14) Dtor_cf68() bool {
	return _this.Get_().(D14_DC38).Cf68
}

func (_this D14) Dtor_cf69() _dafny.Sequence {
	return _this.Get_().(D14_DC38).Cf69
}

func (_this D14) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D14_DC39).Cf70
}

func (_this D14) Dtor_cf71() _dafny.Map {
	return _this.Get_().(D14_DC39).Cf71
}

func (_this D14) Dtor_cf66() _dafny.Sequence {
	return _this.Get_().(D14_DC37).Cf66
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC38:
		{
			return "D14.DC38" + "(" + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D14_DC39:
		{
			return "D14.DC39" + "(" + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D14_DC37:
		{
			return "D14.DC37" + "(" + _dafny.String(data.Cf66) + ")"
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
			return ok && data1.Cf67 == data2.Cf67 && data1.Cf68 == data2.Cf68 && data1.Cf69.Equals(data2.Cf69)
		}
	case D14_DC39:
		{
			data2, ok := other.Get_().(D14_DC39)
			return ok && data1.Cf70.Cmp(data2.Cf70) == 0 && data1.Cf71.Equals(data2.Cf71)
		}
	case D14_DC37:
		{
			data2, ok := other.Get_().(D14_DC37)
			return ok && data1.Cf66.Equals(data2.Cf66)
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

type D15_DC41 struct {
}

func (D15_DC41) isD15() {}

func (CompanionStruct_D15_) Create_DC41_() D15 {
	return D15{D15_DC41{}}
}

func (_this D15) Is_DC41() bool {
	_, ok := _this.Get_().(D15_DC41)
	return ok
}

type D15_DC42 struct {
	Cf73 bool
	Cf74 _dafny.Int
	Cf75 T2
	Cf76 _dafny.Int
}

func (D15_DC42) isD15() {}

func (CompanionStruct_D15_) Create_DC42_(Cf73 bool, Cf74 _dafny.Int, Cf75 T2, Cf76 _dafny.Int) D15 {
	return D15{D15_DC42{Cf73, Cf74, Cf75, Cf76}}
}

func (_this D15) Is_DC42() bool {
	_, ok := _this.Get_().(D15_DC42)
	return ok
}

type D15_DC43 struct {
}

func (D15_DC43) isD15() {}

func (CompanionStruct_D15_) Create_DC43_() D15 {
	return D15{D15_DC43{}}
}

func (_this D15) Is_DC43() bool {
	_, ok := _this.Get_().(D15_DC43)
	return ok
}

type D15_DC40 struct {
	Cf72 _dafny.Set
}

func (D15_DC40) isD15() {}

func (CompanionStruct_D15_) Create_DC40_(Cf72 _dafny.Set) D15 {
	return D15{D15_DC40{Cf72}}
}

func (_this D15) Is_DC40() bool {
	_, ok := _this.Get_().(D15_DC40)
	return ok
}

type D15_DC44 struct {
	Cf77 D15
}

func (D15_DC44) isD15() {}

func (CompanionStruct_D15_) Create_DC44_(Cf77 D15) D15 {
	return D15{D15_DC44{Cf77}}
}

func (_this D15) Is_DC44() bool {
	_, ok := _this.Get_().(D15_DC44)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC41_()
}

func (_this D15) Dtor_cf73() bool {
	return _this.Get_().(D15_DC42).Cf73
}

func (_this D15) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D15_DC42).Cf74
}

func (_this D15) Dtor_cf75() T2 {
	return _this.Get_().(D15_DC42).Cf75
}

func (_this D15) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D15_DC42).Cf76
}

func (_this D15) Dtor_cf72() _dafny.Set {
	return _this.Get_().(D15_DC40).Cf72
}

func (_this D15) Dtor_cf77() D15 {
	return _this.Get_().(D15_DC44).Cf77
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC41:
		{
			return "D15.DC41"
		}
	case D15_DC42:
		{
			return "D15.DC42" + "(" + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D15_DC43:
		{
			return "D15.DC43"
		}
	case D15_DC40:
		{
			return "D15.DC40" + "(" + _dafny.String(data.Cf72) + ")"
		}
	case D15_DC44:
		{
			return "D15.DC44" + "(" + _dafny.String(data.Cf77) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC41:
		{
			_, ok := other.Get_().(D15_DC41)
			return ok
		}
	case D15_DC42:
		{
			data2, ok := other.Get_().(D15_DC42)
			return ok && data1.Cf73 == data2.Cf73 && data1.Cf74.Cmp(data2.Cf74) == 0 && _dafny.AreEqual(data1.Cf75, data2.Cf75) && data1.Cf76.Cmp(data2.Cf76) == 0
		}
	case D15_DC43:
		{
			_, ok := other.Get_().(D15_DC43)
			return ok
		}
	case D15_DC40:
		{
			data2, ok := other.Get_().(D15_DC40)
			return ok && data1.Cf72.Equals(data2.Cf72)
		}
	case D15_DC44:
		{
			data2, ok := other.Get_().(D15_DC44)
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

type D16_DC46 struct {
	Cf79 _dafny.Int
}

func (D16_DC46) isD16() {}

func (CompanionStruct_D16_) Create_DC46_(Cf79 _dafny.Int) D16 {
	return D16{D16_DC46{Cf79}}
}

func (_this D16) Is_DC46() bool {
	_, ok := _this.Get_().(D16_DC46)
	return ok
}

type D16_DC45 struct {
	Cf78 _dafny.Map
}

func (D16_DC45) isD16() {}

func (CompanionStruct_D16_) Create_DC45_(Cf78 _dafny.Map) D16 {
	return D16{D16_DC45{Cf78}}
}

func (_this D16) Is_DC45() bool {
	_, ok := _this.Get_().(D16_DC45)
	return ok
}

type D16_DC47 struct {
	Cf80 D16
}

func (D16_DC47) isD16() {}

func (CompanionStruct_D16_) Create_DC47_(Cf80 D16) D16 {
	return D16{D16_DC47{Cf80}}
}

func (_this D16) Is_DC47() bool {
	_, ok := _this.Get_().(D16_DC47)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC46_(_dafny.Zero)
}

func (_this D16) Dtor_cf79() _dafny.Int {
	return _this.Get_().(D16_DC46).Cf79
}

func (_this D16) Dtor_cf78() _dafny.Map {
	return _this.Get_().(D16_DC45).Cf78
}

func (_this D16) Dtor_cf80() D16 {
	return _this.Get_().(D16_DC47).Cf80
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC46:
		{
			return "D16.DC46" + "(" + _dafny.String(data.Cf79) + ")"
		}
	case D16_DC45:
		{
			return "D16.DC45" + "(" + _dafny.String(data.Cf78) + ")"
		}
	case D16_DC47:
		{
			return "D16.DC47" + "(" + _dafny.String(data.Cf80) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC46:
		{
			data2, ok := other.Get_().(D16_DC46)
			return ok && data1.Cf79.Cmp(data2.Cf79) == 0
		}
	case D16_DC45:
		{
			data2, ok := other.Get_().(D16_DC45)
			return ok && data1.Cf78.Equals(data2.Cf78)
		}
	case D16_DC47:
		{
			data2, ok := other.Get_().(D16_DC47)
			return ok && data1.Cf80.Equals(data2.Cf80)
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

type D17_DC49 struct {
}

func (D17_DC49) isD17() {}

func (CompanionStruct_D17_) Create_DC49_() D17 {
	return D17{D17_DC49{}}
}

func (_this D17) Is_DC49() bool {
	_, ok := _this.Get_().(D17_DC49)
	return ok
}

type D17_DC50 struct {
	Cf82 _dafny.CodePoint
	Cf83 bool
	Cf84 _dafny.Sequence
	Cf85 _dafny.Int
	Cf86 T0
}

func (D17_DC50) isD17() {}

func (CompanionStruct_D17_) Create_DC50_(Cf82 _dafny.CodePoint, Cf83 bool, Cf84 _dafny.Sequence, Cf85 _dafny.Int, Cf86 T0) D17 {
	return D17{D17_DC50{Cf82, Cf83, Cf84, Cf85, Cf86}}
}

func (_this D17) Is_DC50() bool {
	_, ok := _this.Get_().(D17_DC50)
	return ok
}

type D17_DC48 struct {
	Cf81 T0
}

func (D17_DC48) isD17() {}

func (CompanionStruct_D17_) Create_DC48_(Cf81 T0) D17 {
	return D17{D17_DC48{Cf81}}
}

func (_this D17) Is_DC48() bool {
	_, ok := _this.Get_().(D17_DC48)
	return ok
}

type D17_DC51 struct {
	Cf87 D17
}

func (D17_DC51) isD17() {}

func (CompanionStruct_D17_) Create_DC51_(Cf87 D17) D17 {
	return D17{D17_DC51{Cf87}}
}

func (_this D17) Is_DC51() bool {
	_, ok := _this.Get_().(D17_DC51)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC49_()
}

func (_this D17) Dtor_cf82() _dafny.CodePoint {
	return _this.Get_().(D17_DC50).Cf82
}

func (_this D17) Dtor_cf83() bool {
	return _this.Get_().(D17_DC50).Cf83
}

func (_this D17) Dtor_cf84() _dafny.Sequence {
	return _this.Get_().(D17_DC50).Cf84
}

func (_this D17) Dtor_cf85() _dafny.Int {
	return _this.Get_().(D17_DC50).Cf85
}

func (_this D17) Dtor_cf86() T0 {
	return _this.Get_().(D17_DC50).Cf86
}

func (_this D17) Dtor_cf81() T0 {
	return _this.Get_().(D17_DC48).Cf81
}

func (_this D17) Dtor_cf87() D17 {
	return _this.Get_().(D17_DC51).Cf87
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC49:
		{
			return "D17.DC49"
		}
	case D17_DC50:
		{
			return "D17.DC50" + "(" + _dafny.String(data.Cf82) + ", " + _dafny.String(data.Cf83) + ", " + data.Cf84.VerbatimString(true) + ", " + _dafny.String(data.Cf85) + ", " + _dafny.String(data.Cf86) + ")"
		}
	case D17_DC48:
		{
			return "D17.DC48" + "(" + _dafny.String(data.Cf81) + ")"
		}
	case D17_DC51:
		{
			return "D17.DC51" + "(" + _dafny.String(data.Cf87) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC49:
		{
			_, ok := other.Get_().(D17_DC49)
			return ok
		}
	case D17_DC50:
		{
			data2, ok := other.Get_().(D17_DC50)
			return ok && data1.Cf82 == data2.Cf82 && data1.Cf83 == data2.Cf83 && data1.Cf84.Equals(data2.Cf84) && data1.Cf85.Cmp(data2.Cf85) == 0 && _dafny.AreEqual(data1.Cf86, data2.Cf86)
		}
	case D17_DC48:
		{
			data2, ok := other.Get_().(D17_DC48)
			return ok && _dafny.AreEqual(data1.Cf81, data2.Cf81)
		}
	case D17_DC51:
		{
			data2, ok := other.Get_().(D17_DC51)
			return ok && data1.Cf87.Equals(data2.Cf87)
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

type D18_DC53 struct {
	Cf89 _dafny.Int
	Cf90 *C0
	Cf91 _dafny.Sequence
	Cf92 bool
}

func (D18_DC53) isD18() {}

func (CompanionStruct_D18_) Create_DC53_(Cf89 _dafny.Int, Cf90 *C0, Cf91 _dafny.Sequence, Cf92 bool) D18 {
	return D18{D18_DC53{Cf89, Cf90, Cf91, Cf92}}
}

func (_this D18) Is_DC53() bool {
	_, ok := _this.Get_().(D18_DC53)
	return ok
}

type D18_DC52 struct {
	Cf88 _dafny.Map
}

func (D18_DC52) isD18() {}

func (CompanionStruct_D18_) Create_DC52_(Cf88 _dafny.Map) D18 {
	return D18{D18_DC52{Cf88}}
}

func (_this D18) Is_DC52() bool {
	_, ok := _this.Get_().(D18_DC52)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC53_(_dafny.Zero, (*C0)(nil), _dafny.EmptySeq, false)
}

func (_this D18) Dtor_cf89() _dafny.Int {
	return _this.Get_().(D18_DC53).Cf89
}

func (_this D18) Dtor_cf90() *C0 {
	return _this.Get_().(D18_DC53).Cf90
}

func (_this D18) Dtor_cf91() _dafny.Sequence {
	return _this.Get_().(D18_DC53).Cf91
}

func (_this D18) Dtor_cf92() bool {
	return _this.Get_().(D18_DC53).Cf92
}

func (_this D18) Dtor_cf88() _dafny.Map {
	return _this.Get_().(D18_DC52).Cf88
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC53:
		{
			return "D18.DC53" + "(" + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ")"
		}
	case D18_DC52:
		{
			return "D18.DC52" + "(" + _dafny.String(data.Cf88) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC53:
		{
			data2, ok := other.Get_().(D18_DC53)
			return ok && data1.Cf89.Cmp(data2.Cf89) == 0 && data1.Cf90 == data2.Cf90 && data1.Cf91.Equals(data2.Cf91) && data1.Cf92 == data2.Cf92
		}
	case D18_DC52:
		{
			data2, ok := other.Get_().(D18_DC52)
			return ok && data1.Cf88.Equals(data2.Cf88)
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

type D19_DC54 struct {
	Cf93 _dafny.Map
}

func (D19_DC54) isD19() {}

func (CompanionStruct_D19_) Create_DC54_(Cf93 _dafny.Map) D19 {
	return D19{D19_DC54{Cf93}}
}

func (_this D19) Is_DC54() bool {
	_, ok := _this.Get_().(D19_DC54)
	return ok
}

func (CompanionStruct_D19_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D19) Dtor_cf93() _dafny.Map {
	return _this.Get_().(D19_DC54).Cf93
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC54:
		{
			return "D19.DC54" + "(" + _dafny.String(data.Cf93) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC54:
		{
			data2, ok := other.Get_().(D19_DC54)
			return ok && data1.Cf93.Equals(data2.Cf93)
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

type D20_DC55 struct {
	Cf94 _dafny.MultiSet
}

func (D20_DC55) isD20() {}

func (CompanionStruct_D20_) Create_DC55_(Cf94 _dafny.MultiSet) D20 {
	return D20{D20_DC55{Cf94}}
}

func (_this D20) Is_DC55() bool {
	_, ok := _this.Get_().(D20_DC55)
	return ok
}

func (CompanionStruct_D20_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D20) Dtor_cf94() _dafny.MultiSet {
	return _this.Get_().(D20_DC55).Cf94
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC55:
		{
			return "D20.DC55" + "(" + _dafny.String(data.Cf94) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC55:
		{
			data2, ok := other.Get_().(D20_DC55)
			return ok && data1.Cf94.Equals(data2.Cf94)
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
	Fm2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
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
	Fm2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	M1(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Array, globalState *GlobalState)
	F9() bool
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
	Fm2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	M1(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Array, globalState *GlobalState)
	F9() bool
	Fm3(globalState *GlobalState) bool
	M2(p0 _dafny.Array, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) (bool, bool, _dafny.Int, _dafny.Set)
	M3(p0 bool, globalState *GlobalState) bool
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
	F2  bool
	_f0 bool
	_f1 _dafny.Int
	_f3 bool
	_f4 _dafny.Array
	_f5 bool
	_f6 _dafny.Int
	_f7 _dafny.Sequence
	_f8 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = false
	_this._f0 = false
	_this._f1 = _dafny.Zero
	_this._f3 = false
	_this._f4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f5 = false
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.EmptySeq
	_this._f8 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 bool, f3 bool, f4 _dafny.Array, f5 bool, f6 _dafny.Int, f7 _dafny.Sequence, f8 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Array {
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
func (_this *GlobalState) F7() _dafny.Sequence {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F11 _dafny.Set
	F12 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F11 = _dafny.EmptySet
	_this.F12 = _dafny.Zero
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

func (_this *C0) Ctor__(f11 _dafny.Set, f12 _dafny.Int) {
	{
		(_this).F11 = f11
		(_this).F12 = f12
	}
}
func (_this *C0) Fm8(p0 bool, p1 bool, p2 _dafny.Map, globalState *GlobalState) bool {
	{
		var _source4 D1 = Companion_D1_.Create_DC2_(_this.F12, !(!(true)), _this.F12, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(67))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_305_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))).Cardinality()))
		_ = _source4
		if _source4.Is_DC2() {
			var _306___mcc_h0 _dafny.Int = _source4.Get_().(D1_DC2).Cf2
			_ = _306___mcc_h0
			var _307___mcc_h1 bool = _source4.Get_().(D1_DC2).Cf3
			_ = _307___mcc_h1
			var _308___mcc_h2 _dafny.Int = _source4.Get_().(D1_DC2).Cf4
			_ = _308___mcc_h2
			var _309___mcc_h3 _dafny.Int = _source4.Get_().(D1_DC2).Cf5
			_ = _309___mcc_h3
			var _310_cf5 _dafny.Int = _309___mcc_h3
			_ = _310_cf5
			var _311_cf4 _dafny.Int = _308___mcc_h2
			_ = _311_cf4
			var _312_cf3 bool = _307___mcc_h1
			_ = _312_cf3
			var _313_cf2 _dafny.Int = _306___mcc_h0
			_ = _313_cf2
			return _312_cf3
		} else if _source4.Is_DC3() {
			var _314___mcc_h4 bool = _source4.Get_().(D1_DC3).Cf6
			_ = _314___mcc_h4
			var _315___mcc_h5 _dafny.Sequence = _source4.Get_().(D1_DC3).Cf7
			_ = _315___mcc_h5
			var _316___mcc_h6 _dafny.Int = _source4.Get_().(D1_DC3).Cf8
			_ = _316___mcc_h6
			var _317_cf8 _dafny.Int = _316___mcc_h6
			_ = _317_cf8
			var _318_cf7 _dafny.Sequence = _315___mcc_h5
			_ = _318_cf7
			var _319_cf6 bool = _314___mcc_h4
			_ = _319_cf6
			if _319_cf6 {
				return _319_cf6
			} else {
				return _319_cf6
			}
		} else if _source4.Is_DC4() {
			var _320___mcc_h7 _dafny.Int = _source4.Get_().(D1_DC4).Cf9
			_ = _320___mcc_h7
			var _321___mcc_h8 bool = _source4.Get_().(D1_DC4).Cf10
			_ = _321___mcc_h8
			var _322___mcc_h9 bool = _source4.Get_().(D1_DC4).Cf11
			_ = _322___mcc_h9
			var _323___mcc_h10 bool = _source4.Get_().(D1_DC4).Cf12
			_ = _323___mcc_h10
			var _324___mcc_h11 bool = _source4.Get_().(D1_DC4).Cf13
			_ = _324___mcc_h11
			var _325_cf13 bool = _324___mcc_h11
			_ = _325_cf13
			var _326_cf12 bool = _323___mcc_h10
			_ = _326_cf12
			var _327_cf11 bool = _322___mcc_h9
			_ = _327_cf11
			var _328_cf10 bool = _321___mcc_h8
			_ = _328_cf10
			var _329_cf9 _dafny.Int = _320___mcc_h7
			_ = _329_cf9
			return true
		} else if _source4.Is_DC1() {
			var _330___mcc_h12 _dafny.Int = _source4.Get_().(D1_DC1).Cf1
			_ = _330___mcc_h12
			var _331_cf1 _dafny.Int = _330___mcc_h12
			_ = _331_cf1
			return (_this.F12).Cmp(_331_cf1) >= 0
		} else {
			var _332___mcc_h13 D1 = _source4.Get_().(D1_DC5).Cf14
			_ = _332___mcc_h13
			var _333_cf14 D1 = _332___mcc_h13
			_ = _333_cf14
			return (Companion_D1_.Create_DC3_((Companion_D1_.Create_DC4_(_this.F12, false, true, false, false)).Dtor_cf11(), _dafny.UnicodeSeqOfUtf8Bytes("wfyipkfqa"), _this.F12)).Dtor_cf6()
		}
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F13 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F13 = _dafny.EmptySeq
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

func (_this *C1) Ctor__(f13 _dafny.Sequence) {
	{
		(_this).F13 = f13
	}
}
func (_this *C1) Fm2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eiomh")).Cardinality())), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(628))))).Cardinality(), false)).Cardinality()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(649))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_334_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		}))).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(56))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_335_i1 _dafny.Int) _dafny.Int {
			return (_dafny.SetOf(false, !(!(!(true))), !(false), !(!((Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(976))).Cardinality()), false, false, false, false)).Dtor_cf12())))).Cardinality()
		})))
	}
}
func (_this *C1) M6(p0 T0, p1 _dafny.Array, p2 bool, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _336_v0 _dafny.Int
		_ = _336_v0
		_336_v0 = _dafny.IntOfInt64(226)
		var _337_v1 _dafny.Sequence
		_ = _337_v1
		_337_v1 = _dafny.SeqOf(_336_v0)
		var _338_v2 _dafny.Sequence
		_ = _338_v2
		_338_v2 = _dafny.SeqOf(_337_v1)
		var _339_v3 _dafny.MultiSet
		_ = _339_v3
		_339_v3 = _dafny.MultiSetOf(_336_v0)
		var _340_v4 _dafny.Sequence
		_ = _340_v4
		_340_v4 = _dafny.SeqOf(!(p2), p2, p2)
		var _341_v5 _dafny.Array
		_ = _341_v5
		var _nwElement0_6 _dafny.Int = _336_v0
		_ = _nwElement0_6
		var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(23))
		_ = _nw50
		(_nw50).ArraySet1(_nwElement0_6, 0)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_338_v2).Cardinality()), 1)
		(_nw50).ArraySet1(_336_v0, 2)
		(_nw50).ArraySet1(_336_v0, 3)
		(_nw50).ArraySet1(_336_v0, 4)
		(_nw50).ArraySet1(_336_v0, 5)
		(_nw50).ArraySet1(_336_v0, 6)
		(_nw50).ArraySet1(_336_v0, 7)
		(_nw50).ArraySet1(_336_v0, 8)
		(_nw50).ArraySet1((_dafny.Zero).Minus((_339_v3).Cardinality()), 9)
		(_nw50).ArraySet1(_336_v0, 10)
		(_nw50).ArraySet1(_336_v0, 11)
		(_nw50).ArraySet1(_dafny.IntOfInt64(893), 12)
		(_nw50).ArraySet1(_336_v0, 13)
		(_nw50).ArraySet1(_336_v0, 14)
		(_nw50).ArraySet1((_dafny.Zero).Minus(_336_v0), 15)
		(_nw50).ArraySet1(_dafny.IntOfUint32((_340_v4).Cardinality()), 16)
		(_nw50).ArraySet1(_336_v0, 17)
		(_nw50).ArraySet1(_336_v0, 18)
		(_nw50).ArraySet1(_336_v0, 19)
		(_nw50).ArraySet1(_336_v0, 20)
		(_nw50).ArraySet1(_336_v0, 21)
		(_nw50).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm0(globalState)), 22)
		_341_v5 = _nw50
		var _342_v6 _dafny.Map
		_ = _342_v6
		_342_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v5, _341_v5)
		_342_v6 = (_342_v6).Update(_341_v5, _341_v5)
		var _343_v7 _dafny.Array
		_ = _343_v7
		var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(11))
		_ = _nw51
		_343_v7 = _nw51
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_343_v7), 0))); ; {
			_guard_loop_2, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _344_i0 _dafny.Int
			_344_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_344_i0).Sign() != -1) && ((_344_i0).Cmp(_dafny.ArrayLen((_343_v7), 0)) < 0)) {
				(_343_v7).ArraySet1(((_dafny.SetOf(p2)).Union(_dafny.SetOf(p2, p2, p2, p2, p2))).Difference(_dafny.SetOf(p2, p2)), (_344_i0).Int())
			}
		}
		_336_v0 = (func() _dafny.Int {
			if p2 {
				return Companion_Default___.Fm0(globalState)
			}
			return Companion_Default___.Fm0(globalState)
		})()
		if ((_dafny.Zero).Minus(_336_v0)).Cmp((_336_v0).Times(_336_v0)) >= 0 {
			var _345_v8 _dafny.MultiSet
			_ = _345_v8
			_345_v8 = _dafny.MultiSetOf(p1, p1, p1)
			_345_v8 = ((Companion_D3_.Create_DC7_(_dafny.MultiSetOf(p1))).Dtor_cf16()).Difference(_345_v8)
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_341_v5), 0))
			_ = _index24
			(_341_v5).ArraySet1(((_dafny.Zero).Minus(_336_v0)).Plus(_336_v0), (_index24).Int())
			var _346_v9 _dafny.Map
			_ = _346_v9
			_346_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v0, p2)
			var _347_v10 _dafny.Map
			_ = _347_v10
			_347_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v0, (_346_v9).Cardinality())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_341_v5), 0))
			_ = _index25
			(_341_v5).ArraySet1((_336_v0).Minus((func() _dafny.Int {
				if (_347_v10).Contains(_dafny.IntOfInt64(180)) {
					return (_347_v10).Get(_dafny.IntOfInt64(180)).(_dafny.Int)
				}
				return (_337_v1).Select((Companion_Default___.SafeIndex(_336_v0, _dafny.IntOfUint32((_337_v1).Cardinality()))).Uint32()).(_dafny.Int)
			})()), (_index25).Int())
			var _348_v11 _dafny.Set
			_ = _348_v11
			_348_v11 = _dafny.SetOf(_337_v1)
			var _349_v12 *C0
			_ = _349_v12
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__((_348_v11).Union(_dafny.SetOf(_337_v1, _337_v1)), Companion_Default___.Fm0(globalState))
			_349_v12 = _nw52
			var _350_v13 _dafny.CodePoint
			_ = _350_v13
			_350_v13 = _dafny.CodePoint('j')
			var _351_v14 D1
			_ = _351_v14
			_351_v14 = Companion_D1_.Create_DC3_(false, _dafny.UnicodeSeqOfUtf8Bytes("hgnqr"), (_341_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_341_v5), 0))).Int()).(_dafny.Int))
			var _352_v15 _dafny.Map
			_ = _352_v15
			_352_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_350_v13, ((_dafny.Zero).Minus(_349_v12.F12)).Cmp((Companion_D1_.Create_DC4_(_336_v0, (_351_v14).Dtor_cf6(), p2, p2, p2)).Dtor_cf9()) != 0)
			var _353_v16 _dafny.Set
			_ = _353_v16
			_353_v16 = _dafny.SetOf(_349_v12, _349_v12, _349_v12)
			var _354_v17 _dafny.Map
			_ = _354_v17
			_354_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_349_v12)).IsDisjointFrom(_353_v16), (func() *C0 {
				if p2 {
					return _349_v12
				}
				return _349_v12
			})())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_341_v5), 0))
			_ = _index26
			var _rhs53 *C0 = (func() *C0 {
				if (_354_v17).Contains(p2) {
					return (_354_v17).Get(p2).(*C0)
				}
				return _349_v12
			})()
			_ = _rhs53
			var _rhs54 _dafny.Int = _349_v12.F12
			_ = _rhs54
			var _rhs55 bool = p2
			_ = _rhs55
			var _rhs56 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_350_v13, !(p2))
			_ = _rhs56
			var _lhs25 _dafny.Array = _341_v5
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_341_v5), 0))
			_ = _lhs26
			var _lhs27 *GlobalState = globalState
			_ = _lhs27
			_349_v12 = _rhs53
			(_lhs25).ArraySet1(_rhs54, (_lhs26).Int())
			_lhs27.F2 = _rhs55
			_352_v15 = _rhs56
			r1 = p2
		} else {
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((p1), 0))
			_ = _index27
			(p1).ArraySet1(p2, (_index27).Int())
			var _355_v18 _dafny.MultiSet
			_ = _355_v18
			_355_v18 = _dafny.MultiSetOf(p2, !(p2))
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((p1), 0))
			_ = _index28
			var _rhs57 bool = p2
			_ = _rhs57
			var _rhs58 bool = p2
			_ = _rhs58
			var _rhs59 _dafny.Int = (((_355_v18).Union(_dafny.MultiSetFromSeq(_340_v4))).Intersection(_355_v18)).Cardinality()
			_ = _rhs59
			var _rhs60 _dafny.Int = ((_339_v3).Cardinality()).Times((_336_v0).Times(_336_v0))
			_ = _rhs60
			var _rhs61 _dafny.Int = _336_v0
			_ = _rhs61
			var _lhs28 *GlobalState = globalState
			_ = _lhs28
			var _lhs29 _dafny.Array = p1
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((p1), 0))
			_ = _lhs30
			_lhs28.F2 = _rhs57
			(_lhs29).ArraySet1(_rhs58, (_lhs30).Int())
			_336_v0 = _rhs59
			_336_v0 = _rhs60
			_336_v0 = _rhs61
			var _rhs62 bool = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((p1), 0))).Int()).(bool)
			_ = _rhs62
			var _rhs63 _dafny.Int = _336_v0
			_ = _rhs63
			var _rhs64 _dafny.MultiSet = _339_v3
			_ = _rhs64
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			_lhs31.F2 = _rhs62
			_336_v0 = _rhs63
			_339_v3 = _rhs64
			var _356_v20 _dafny.Set
			_ = _356_v20
			_356_v20 = _dafny.SetOf(p2)
			var _357_v21 _dafny.MultiSet
			_ = _357_v21
			_357_v21 = _dafny.MultiSetOf(_356_v20, _356_v20, _356_v20, _356_v20)
			_336_v0 = ((func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter19 := _dafny.Iterate((_357_v21).Elements()); ; {
					_compr_16, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _358_v19 _dafny.Set
					_358_v19 = interface{}(_compr_16).(_dafny.Set)
					if (_357_v21).Contains(_358_v19) {
						_coll16.Add(_358_v19, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((p1), 0))).Int()).(bool))
					}
				}
				return _coll16.ToMap()
			}()).Cardinality()).Times(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(47), _336_v0))
			var _359_v22 _dafny.Set
			_ = _359_v22
			_359_v22 = _dafny.SetOf(_dafny.SeqOf(_336_v0))
			var _360_v23 *C0
			_ = _360_v23
			var _nw53 *C0 = New_C0_()
			_ = _nw53
			_nw53.Ctor__(_359_v22, _336_v0)
			_360_v23 = _nw53
			var _361_v24 _dafny.Map
			_ = _361_v24
			_361_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_360_v23, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((p1), 0))).Int()).(bool))
			_361_v24 = (_361_v24).Update(_360_v23, !((_360_v23.F12).Cmp((_dafny.Zero).Minus(_336_v0)) >= 0))
			(_360_v23).F12 = _360_v23.F12
		}
		var _362_v25 _dafny.Sequence
		_ = _362_v25
		_362_v25 = _337_v1
		var _source5 _dafny.Sequence = _362_v25
		_ = _source5
		var _363___mcc_h0 _dafny.Sequence = _source5
		_ = _363___mcc_h0
		var _364_cf0 _dafny.Sequence = _363___mcc_h0
		_ = _364_cf0
		var _365_v26 _dafny.CodePoint
		_ = _365_v26
		_365_v26 = _dafny.CodePoint('l')
		var _366_v27 _dafny.Map
		_ = _366_v27
		_366_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v0, _365_v26)
		_336_v0 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_366_v27).Cardinality()).Times((_dafny.Zero).Minus(_336_v0)), _336_v0))
		var _rhs65 bool = (false) || (true)
		_ = _rhs65
		var _rhs66 _dafny.Int = (func() _dafny.Int {
			if !(p2) {
				return _336_v0
			}
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).Fm2(_336_v0, p2, p2, _dafny.IntOfUint32((_340_v4).Cardinality()), globalState), _364_cf0)).Cardinality()))
		})()
		_ = _rhs66
		var _rhs67 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(globalState), (_dafny.Zero).Minus(_336_v0))
		_ = _rhs67
		var _rhs68 bool = (!(p2)) && (p2)
		_ = _rhs68
		var _lhs32 *GlobalState = globalState
		_ = _lhs32
		var _lhs33 *GlobalState = globalState
		_ = _lhs33
		_lhs32.F2 = _rhs65
		_336_v0 = _rhs66
		_336_v0 = _rhs67
		_lhs33.F2 = _rhs68
		r0 = p2
		var _367_v28 D3
		_ = _367_v28
		_367_v28 = Companion_D3_.Create_DC10_()
		var _368_v29 _dafny.Map
		_ = _368_v29
		_368_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _367_v28)
		var _369_v30 _dafny.Sequence
		_ = _369_v30
		_369_v30 = _dafny.SeqOf(_368_v29)
		var _370_v31 _dafny.Map
		_ = _370_v31
		_370_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v0, _336_v0)
		r1 = !((_369_v30).Select((Companion_Default___.SafeIndex((_370_v31).Cardinality(), _dafny.IntOfUint32((_369_v30).Cardinality()))).Uint32()).(_dafny.Map)).Contains(p2)
		if true {
			var _371_v32 _dafny.CodePoint
			_ = _371_v32
			_371_v32 = _dafny.CodePoint('t')
			var _372_v33 _dafny.Map
			_ = _372_v33
			_372_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v32, p2)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))
			_ = _index29
			(p1).ArraySet1((func() bool {
				if (_372_v33).Contains(_371_v32) {
					return (_372_v33).Get(_371_v32).(bool)
				}
				return p2
			})(), (_index29).Int())
			var _373_v34 _dafny.Set
			_ = _373_v34
			_373_v34 = _dafny.SetOf(_dafny.IntOfUint32(((p0).Fm2((_dafny.Zero).Minus(_336_v0), true, p2, _336_v0, globalState)).Cardinality()))
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))
			_ = _index30
			(p1).ArraySet1(((_dafny.SetOf(_336_v0, (_dafny.Zero).Minus(_336_v0))).Difference(_373_v34)).IsSubsetOf(_373_v34), (_index30).Int())
			var _374_v35 _dafny.Sequence
			_ = _374_v35
			_374_v35 = _dafny.UnicodeSeqOfUtf8Bytes("dcyxdyb")
			var _375_v36 _dafny.Map
			_ = _375_v36
			_375_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v0, true)
			var _376_v37 _dafny.Array
			_ = _376_v37
			var _nwElement0_7 bool = p2
			_ = _nwElement0_7
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(2))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_7, 0)
			(_nw54).ArraySet1((_dafny.SetOf(_336_v0, _dafny.IntOfUint32((_337_v1).Cardinality()))).IsDisjointFrom(Companion_Default___.Fm14(_336_v0, _336_v0, (_375_v36).Cardinality(), globalState)), 1)
			_376_v37 = _nw54
			var _rhs69 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_374_v35, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(404))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_377_v32 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_378_i1 _dafny.Int) _dafny.CodePoint {
					return _377_v32
				}
			})(_371_v32)))), _374_v35)
			_ = _rhs69
			var _rhs70 _dafny.Array = _376_v37
			_ = _rhs70
			_374_v35 = _rhs69
			_376_v37 = _rhs70
			if (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))).Int()).(bool) {
				var _379_v38 _dafny.Set
				_ = _379_v38
				_379_v38 = _dafny.SetOf(_337_v1, _dafny.SeqOf(_336_v0), _337_v1, _337_v1)
				var _380_v39 *C0
				_ = _380_v39
				var _nw55 *C0 = New_C0_()
				_ = _nw55
				_nw55.Ctor__((_379_v38).Intersection(_379_v38), (_336_v0).Plus(_dafny.IntOfInt64(73)))
				_380_v39 = _nw55
				var _381_v40 _dafny.Map
				_ = _381_v40
				_381_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))).Int()).(bool), p2)
				(_380_v39).F12 = (func() _dafny.Int {
					if p2 {
						return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(144), _dafny.IntOfUint32((_dafny.SeqOf(_380_v39.F12)).Cardinality()))
					}
					return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("aajbnx"), (Companion_Default___.SafeIndex((_381_v40).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aajbnx")).Cardinality()))).Uint32(), _371_v32)).Cardinality())).Minus(_336_v0)
				})()
				var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))
				_ = _index31
				(p1).ArraySet1(false, (_index31).Int())
				var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))
				_ = _index32
				(p1).ArraySet1(Companion_Default___.Fm1(_339_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fgnkdwmjc")).Cardinality()), globalState), (_index32).Int())
				var _382_v41 _dafny.Map
				_ = _382_v41
				_382_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_380_v39.F12, _374_v35)
				_382_v41 = (_382_v41).Update(_380_v39.F12, _374_v35)
			} else {
				var _383_v42 _dafny.Array
				_ = _383_v42
				var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
				_ = _nw56
				_383_v42 = _nw56
				var _384_v43 _dafny.Array
				_ = _384_v43
				var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw57
				_384_v43 = _nw57
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_383_v42), 0))
				_ = _index33
				(_383_v42).ArraySet1(_384_v43, (_index33).Int())
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_383_v42), 0))
				_ = _index34
				(_383_v42).ArraySet1(_384_v43, (_index34).Int())
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))
				_ = _index35
				(_341_v5).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p2, p2)).Cardinality()), (_index35).Int())
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))
				_ = _index36
				(_341_v5).ArraySet1((func() _dafny.Int {
					if Companion_Default___.Fm1(_339_v3, (_373_v34).Cardinality(), globalState) {
						return _dafny.IntOfUint32((_337_v1).Cardinality())
					}
					return _336_v0
				})(), (_index36).Int())
				var _385_v44 _dafny.Array
				_ = _385_v44
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_7
				var _nw58 _dafny.Array
				_ = _nw58
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw58 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Int = (func(_386_v35 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_387_i2 _dafny.Int) _dafny.Int {
							return (_387_i2).Minus((Companion_D1_.Create_DC1_(_dafny.IntOfUint32((_386_v35).Cardinality()))).Dtor_cf1())
						}
					})(_374_v35)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw58 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw58).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw58).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_385_v44 = _nw58
				var _388_v45 _dafny.Map
				_ = _388_v45
				_388_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))).Int()).(bool), _385_v44)
				_388_v45 = (_388_v45).Update(Companion_Default___.Fm1(_339_v3, (_341_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))).Int()).(_dafny.Int), globalState), (func() _dafny.Array {
					if (_388_v45).Contains(!(p2)) {
						return (_388_v45).Get(!(p2)).(_dafny.Array)
					}
					return _385_v44
				})())
				var _389_v47 _dafny.Set
				_ = _389_v47
				_389_v47 = _dafny.SetOf(_337_v1)
				var _390_v48 *C0
				_ = _390_v48
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__(_389_v47, (_341_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))).Int()).(_dafny.Int))
				_390_v48 = _nw59
				var _391_v49 _dafny.Map
				_ = _391_v49
				_391_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_390_v48, p2)
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))
				_ = _index37
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))
				_ = _index38
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))
				_ = _index39
				var _rhs71 _dafny.Map = (func() _dafny.Map {
					if true {
						return Companion_Default___.Fm15(p2, globalState)
					}
					return (_375_v36).Merge(func() _dafny.Map {
						var _coll17 = _dafny.NewMapBuilder()
						_ = _coll17
						for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-724), _dafny.IntOfInt64(374))); ; {
							_compr_17, _ok20 := _iter20()
							if !_ok20 {
								break
							}
							var _392_v46 _dafny.Int
							_392_v46 = interface{}(_compr_17).(_dafny.Int)
							if ((_dafny.IntOfInt64(-724)).Cmp(_392_v46) <= 0) && ((_392_v46).Cmp(_dafny.IntOfInt64(374)) < 0) {
								_coll17.Add(Companion_Default___.SafeDivisionInt(_392_v46, _336_v0), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))).Int()).(bool))
							}
						}
						return _coll17.ToMap()
					}())
				})()
				_ = _rhs71
				var _rhs72 bool = (_340_v4).Select((Companion_Default___.SafeIndex(((_391_v49).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_390_v48, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))).Int()).(bool)))).Cardinality(), _dafny.IntOfUint32((_340_v4).Cardinality()))).Uint32()).(bool)
				_ = _rhs72
				var _rhs73 _dafny.Int = (_390_v48.F12).Times((_390_v48.F12).Plus((_341_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))).Int()).(_dafny.Int)))
				_ = _rhs73
				var _rhs74 _dafny.Int = (_341_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))).Int()).(_dafny.Int)
				_ = _rhs74
				var _lhs34 _dafny.Array = p1
				_ = _lhs34
				var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))
				_ = _lhs35
				var _lhs36 _dafny.Array = _341_v5
				_ = _lhs36
				var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))
				_ = _lhs37
				var _lhs38 _dafny.Array = _341_v5
				_ = _lhs38
				var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_341_v5), 0))
				_ = _lhs39
				_375_v36 = _rhs71
				(_lhs34).ArraySet1(_rhs72, (_lhs35).Int())
				(_lhs36).ArraySet1(_rhs73, (_lhs37).Int())
				(_lhs38).ArraySet1(_rhs74, (_lhs39).Int())
				var _393_v50 _dafny.Map
				_ = _393_v50
				_393_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))).Int()).(bool), _dafny.IntOfInt64(521))
				r1 = (func() bool {
					if !(_393_v50).Equals(_393_v50) {
						return true
					}
					return (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((p1), 0))).Int()).(bool)
				})()
			}
			_336_v0 = Companion_Default___.SafeModuloInt(_336_v0, (_dafny.Zero).Minus(_336_v0))
			_336_v0 = Companion_Default___.SafeModuloInt(_336_v0, _336_v0)
		} else {
			var _394_v51 _dafny.Array
			_ = _394_v51
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_8
			var _nw60 _dafny.Array
			_ = _nw60
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw60 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) D3 = func(_395_i3 _dafny.Int) D3 {
					return Companion_D3_.Create_DC8_()
				}
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
			_394_v51 = _nw60
			_394_v51 = _394_v51
			var _396_v52 _dafny.Map
			_ = _396_v52
			_396_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _336_v0)
			_396_v52 = _396_v52
			var _397_v53 _dafny.Array
			_ = _397_v53
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_9
			var _nw61 _dafny.Array
			_ = _nw61
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw61 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.CodePoint = func(_398_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw61 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw61).ArraySet1CodePoint(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw61).ArraySet1CodePoint(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_397_v53 = _nw61
			var _399_v54 _dafny.CodePoint
			_ = _399_v54
			_399_v54 = _dafny.CodePoint('v')
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_397_v53), 0))
			_ = _index40
			(_397_v53).ArraySet1CodePoint(_399_v54, (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_397_v53), 0))
			_ = _index41
			(_397_v53).ArraySet1CodePoint(_399_v54, (_index41).Int())
			var _400_v55 _dafny.MultiSet
			_ = _400_v55
			_400_v55 = _dafny.MultiSetOf(p2, p2, p2, true, p2)
			var _401_v56 D1
			_ = _401_v56
			_401_v56 = Companion_D1_.Create_DC4_((_400_v55).Cardinality(), p2, p2, p2, p2)
			var _402_v57 D1
			_ = _402_v57
			_402_v57 = Companion_D1_.Create_DC5_(_401_v56)
			var _403_v58 _dafny.Map
			_ = _403_v58
			_403_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_399_v54, p2)
			var _404_v59 *C0
			_ = _404_v59
			var _nw62 *C0 = New_C0_()
			_ = _nw62
			_nw62.Ctor__(Companion_Default___.Fm16(_403_v58, p2, globalState), _336_v0)
			_404_v59 = _nw62
			var _405_v60 _dafny.Map
			_ = _405_v60
			_405_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D1 {
				if Companion_Default___.Fm1(_339_v3, _336_v0, globalState) {
					return _402_v57
				}
				return _402_v57
			})(), _404_v59)
			_405_v60 = _405_v60
			if !((!(p2)) == (p2)) {
				var _406_v61 _dafny.Map
				_ = _406_v61
				_406_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _404_v59.F12)
				var _407_v62 _dafny.Sequence
				_ = _407_v62
				_407_v62 = _dafny.SeqOf(p1, p1)
				var _408_v63 _dafny.Sequence
				_ = _408_v63
				_408_v63 = _dafny.UnicodeSeqOfUtf8Bytes("xuorgwewj")
				var _409_v64 _dafny.Map
				_ = _409_v64
				_409_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_404_v59.F12, p2)
				var _410_v65 _dafny.Map
				_ = _410_v65
				_410_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(p2, (_404_v59).Fm8(p2, p2, _409_v64, globalState), p2, p2, p2)).Cardinality(), p2)
				var _411_v66 _dafny.Sequence
				_ = _411_v66
				_411_v66 = _dafny.SeqOf(_409_v64)
				var _412_v67 _dafny.Map
				_ = _412_v67
				_412_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
				var _413_v68 D1
				_ = _413_v68
				_413_v68 = Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_337_v1).Cardinality()), p2, (func() bool {
					if (_412_v67).Contains(p1) {
						return (_412_v67).Get(p1).(bool)
					}
					return p2
				})(), p2, (_404_v59).Fm8(p2, p2, _410_v65, globalState))
				var _rhs75 _dafny.Int = ((_406_v61).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_407_v62).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_408_v63).Cardinality()), _dafny.IntOfUint32((_407_v62).Cardinality()))).Uint32()).(_dafny.Array), _404_v59.F12))).Cardinality()
				_ = _rhs75
				var _rhs76 bool = (func() bool {
					if p2 {
						return p2
					}
					return (_dafny.IntOfUint32((_411_v66).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_413_v68, _413_v68)).Cardinality())) != 0
				})()
				_ = _rhs76
				var _rhs77 _dafny.Array = _341_v5
				_ = _rhs77
				var _lhs40 *GlobalState = globalState
				_ = _lhs40
				_336_v0 = _rhs75
				_lhs40.F2 = _rhs76
				_341_v5 = _rhs77
				var _414_v69 _dafny.Map
				_ = _414_v69
				_414_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_340_v4, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_339_v3).Contains(_336_v0) {
						return (_339_v3).Multiplicity(_336_v0)
					}
					return _404_v59.F12
				})(), _dafny.IntOfUint32((_340_v4).Cardinality()))).Uint32(), p2), !(p2))
				_414_v69 = (_414_v69).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p2), _340_v4), (p2) && (p2))
				var _415_v70 _dafny.Set
				_ = _415_v70
				_415_v70 = _dafny.SetOf(p1, p1, p1)
				(globalState).F2 = !(p2) || ((_404_v59.F12).Cmp((_dafny.Zero).Minus((_415_v70).Cardinality())) >= 0)
				(_404_v59).F12 = (_dafny.Zero).Minus(_336_v0)
				var _416_v71 _dafny.Map
				_ = _416_v71
				_416_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v1, _404_v59.F12)
				var _417_v72 _dafny.Map
				_ = _417_v72
				_417_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v0, _404_v59.F12)
				_416_v71 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v1, _404_v59.F12)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_404_v59.F12, (func() _dafny.Int {
					if (_417_v72).Contains((Companion_Default___.Fm15(p2, globalState)).Cardinality()) {
						return (_417_v72).Get((Companion_Default___.Fm15(p2, globalState)).Cardinality()).(_dafny.Int)
					}
					return _336_v0
				})()), _336_v0)).Merge(_416_v71))
			} else {
				var _418_v73 _dafny.Set
				_ = _418_v73
				_418_v73 = _dafny.SetOf(_404_v59.F12, _336_v0)
				var _419_v74 _dafny.Set
				_ = _419_v74
				_419_v74 = _418_v73
				var _420_v75 _dafny.Map
				_ = _420_v75
				_420_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(), _336_v0)
				var _421_v76 _dafny.Array
				_ = _421_v76
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_10
				var _nw63 _dafny.Array
				_ = _nw63
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw63 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Set = (func(_422_v73 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_423_i5 _dafny.Int) _dafny.Set {
							return _422_v73
						}
					})(_418_v73)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw63 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw63).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw63).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_421_v76 = _nw63
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_421_v76), 0))
				_ = _index42
				(_421_v76).ArraySet1(_418_v73, (_index42).Int())
				var _424_v77 _dafny.Sequence
				_ = _424_v77
				_424_v77 = _dafny.UnicodeSeqOfUtf8Bytes("rrhtmhjuf")
				var _425_v78 D1
				_ = _425_v78
				_425_v78 = Companion_D1_.Create_DC3_(!(Companion_Default___.Fm1(_339_v3, _404_v59.F12, globalState)), _424_v77, _dafny.IntOfInt64(404))
				var _426_v79 _dafny.Map
				_ = _426_v79
				_426_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_424_v77, p2)
				var _427_v80 _dafny.Map
				_ = _427_v80
				_427_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_404_v59.F12, _424_v77)
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_421_v76), 0))
				_ = _index43
				var _rhs78 _dafny.Set = _419_v74
				_ = _rhs78
				var _rhs79 _dafny.Map = Companion_Default___.Fm17(_404_v59.F12, _426_v79, p2, (func() _dafny.Int {
					if (_400_v55).Contains(p2) {
						return (_400_v55).Multiplicity(p2)
					}
					return (_427_v80).Cardinality()
				})(), globalState)
				_ = _rhs79
				var _rhs80 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfUint32(((_338_v2).Select((Companion_Default___.SafeIndex(_404_v59.F12, _dafny.IntOfUint32((_338_v2).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Minus((func() _dafny.Int {
					if p2 {
						return _404_v59.F12
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality())
				})()))
				_ = _rhs80
				var _rhs81 _dafny.Set = (_418_v73).Intersection(_418_v73)
				_ = _rhs81
				var _rhs82 D1 = _425_v78
				_ = _rhs82
				var _lhs41 *C0 = _404_v59
				_ = _lhs41
				var _lhs42 _dafny.Array = _421_v76
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_421_v76), 0))
				_ = _lhs43
				_419_v74 = _rhs78
				_420_v75 = _rhs79
				_lhs41.F12 = _rhs80
				(_lhs42).ArraySet1(_rhs81, (_lhs43).Int())
				_425_v78 = _rhs82
				_337_v1 = _dafny.SeqOf(_404_v59.F12)
				var _428_v81 _dafny.Map
				_ = _428_v81
				_428_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), false)).Cardinality(), p2)
				Companion_Default___.M0((func() bool {
					if (_428_v81).Contains(_dafny.IntOfInt64(701)) {
						return (_428_v81).Get(_dafny.IntOfInt64(701)).(bool)
					}
					return p2
				})(), globalState)
				var _rhs83 _dafny.Int = (_dafny.IntOfInt64(217)).Times(Companion_Default___.Fm0(globalState))
				_ = _rhs83
				var _rhs84 bool = p2
				_ = _rhs84
				var _rhs85 _dafny.Int = _dafny.IntOfInt64(-297)
				_ = _rhs85
				var _lhs44 *C0 = _404_v59
				_ = _lhs44
				_lhs44.F12 = _rhs83
				r0 = _rhs84
				_336_v0 = _rhs85
				var _429_v82 _dafny.Map
				_ = _429_v82
				_429_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _404_v59)
				_429_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _404_v59)
			}
		}
		var _430_v83 _dafny.CodePoint
		_ = _430_v83
		_430_v83 = _dafny.CodePoint('x')
		var _431_v84 _dafny.MultiSet
		_ = _431_v84
		_431_v84 = _dafny.MultiSetOf(_430_v83, _430_v83)
		r0 = (func() bool {
			if !((_dafny.IntOfInt64(565)).Cmp(_336_v0) >= 0) {
				return p2
			}
			return ((func() _dafny.Set {
				var _coll18 = _dafny.NewBuilder()
				_ = _coll18
				for _iter21 := _dafny.Iterate((_431_v84).Elements()); ; {
					_compr_18, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _432_v85 _dafny.CodePoint
					_432_v85 = interface{}(_compr_18).(_dafny.CodePoint)
					if (_431_v84).Contains(_432_v85) {
						_coll18.Add(_432_v85)
					}
				}
				return _coll18.ToSet()
			}()).Cardinality()).Cmp(_336_v0) != 0
		})()
		r1 = p2
		return r0, r1
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f9  bool
	_f14 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f9 = false
	_this._f14 = _dafny.Zero
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

func (_this *C2) F9() bool {
	return _this._f9
}
func (_this *C2) Ctor__(f14 _dafny.Int, f9 bool) {
	{
		(_this)._f14 = f14
		(_this)._f9 = f9
	}
}
func (_this *C2) Fm2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F14()), _dafny.SeqOf((_this).F14(), (_this).F14(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(460))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}(func(_433_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality())))
	}
}
func (_this *C2) M1(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Array, globalState *GlobalState) {
	{
		var _434_v0 _dafny.Int
		_ = _434_v0
		_434_v0 = _dafny.IntOfInt64(884)
		var _435_v1 _dafny.Sequence
		_ = _435_v1
		_435_v1 = _dafny.UnicodeSeqOfUtf8Bytes("vqsey")
		var _436_v2 _dafny.Sequence
		_ = _436_v2
		_436_v2 = _dafny.SeqOf(p0)
		var _437_v3 _dafny.Set
		_ = _437_v3
		_437_v3 = _dafny.SetOf(_436_v2)
		var _438_v4 *C0
		_ = _438_v4
		var _nw64 *C0 = New_C0_()
		_ = _nw64
		_nw64.Ctor__(_437_v3, Companion_Default___.Fm0(globalState))
		_438_v4 = _nw64
		var _439_v5 _dafny.Sequence
		_ = _439_v5
		_439_v5 = _dafny.SeqOf(_438_v4, _438_v4)
		var _440_v6 _dafny.CodePoint
		_ = _440_v6
		_440_v6 = _dafny.CodePoint('o')
		var _441_v7 _dafny.MultiSet
		_ = _441_v7
		_441_v7 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_435_v1, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_439_v5, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_439_v5).Cardinality()))).Uint32(), _438_v4), (Companion_Default___.SafeIndex(_438_v4.F12, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_439_v5, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_439_v5).Cardinality()))).Uint32(), _438_v4)).Cardinality()))).Uint32(), _438_v4)).Cardinality())), _dafny.IntOfUint32((_435_v1).Cardinality()))).Uint32(), _440_v6), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), p1)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_435_v1, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_439_v5, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_439_v5).Cardinality()))).Uint32(), _438_v4), (Companion_Default___.SafeIndex(_438_v4.F12, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_439_v5, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_439_v5).Cardinality()))).Uint32(), _438_v4)).Cardinality()))).Uint32(), _438_v4)).Cardinality())), _dafny.IntOfUint32((_435_v1).Cardinality()))).Uint32(), _440_v6)).Cardinality()))).Uint32(), _440_v6)).Cardinality()))
		_434_v0 = (func() _dafny.Int {
			if p2 {
				return (func() _dafny.Int {
					if (_441_v7).Contains((_this).F14()) {
						return (_441_v7).Multiplicity((_this).F14())
					}
					return (_this).F14()
				})()
			}
			return _dafny.IntOfUint32((_435_v1).Cardinality())
		})()
		(globalState).F2 = !(p2)
		if p1 {
			var _442_v8 _dafny.Array
			_ = _442_v8
			var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw65
			_442_v8 = _nw65
			var _443_v9 _dafny.Array
			_ = _443_v9
			_443_v9 = _442_v8
			var _444_v10 _dafny.Array
			_ = _444_v10
			var _nw66 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw66
			_444_v10 = _nw66
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))
			_ = _index44
			(_444_v10).ArraySet1((_this).F9(), (_index44).Int())
			var _445_v11 _dafny.Map
			_ = _445_v11
			_445_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _440_v6)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))
			_ = _index45
			var _rhs86 _dafny.Array = _443_v9
			_ = _rhs86
			var _rhs87 bool = !(_445_v11).Equals(_445_v11)
			_ = _rhs87
			var _lhs45 _dafny.Array = _444_v10
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))
			_ = _lhs46
			_443_v9 = _rhs86
			(_lhs45).ArraySet1(_rhs87, (_lhs46).Int())
			var _446_v12 _dafny.Map
			_ = _446_v12
			_446_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (func() bool {
				if Companion_Default___.Fm1(_441_v7, _dafny.IntOfInt64(620), globalState) {
					return (_this).F9()
				}
				return (_444_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))).Int()).(bool)
			})())
			_446_v12 = _446_v12
			var _447_v13 _dafny.Set
			_ = _447_v13
			_447_v13 = _dafny.SetOf(_434_v0)
			(globalState).F2 = (_447_v13).Equals(_447_v13)
			if true {
				(_438_v4).F12 = _438_v4.F12
				var _448_v14 _dafny.Set
				_ = _448_v14
				_448_v14 = _dafny.SetOf(_435_v1)
				var _449_v15 _dafny.Array
				_ = _449_v15
				var _nwElement0_8 _dafny.Int = (_448_v14).Cardinality()
				_ = _nwElement0_8
				var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(2))
				_ = _nw67
				(_nw67).ArraySet1(_nwElement0_8, 0)
				(_nw67).ArraySet1(Companion_Default___.SafeModuloInt(_438_v4.F12, _dafny.IntOfInt64(679)), 1)
				_449_v15 = _nw67
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_449_v15), 0))
				_ = _index46
				(_449_v15).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_438_v4.F12), _438_v4.F12, _dafny.IntOfUint32((_436_v2).Cardinality()), _438_v4.F12, _438_v4.F12)).Cardinality()), (_index46).Int())
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_449_v15), 0))
				_ = _index47
				(_449_v15).ArraySet1(_438_v4.F12, (_index47).Int())
				var _450_v16 D1
				_ = _450_v16
				_450_v16 = Companion_D1_.Create_DC2_((_this).F14(), p2, _438_v4.F12, _438_v4.F12)
				_450_v16 = Companion_D1_.Create_DC2_(_dafny.IntOfInt64(718), (_444_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))).Int()).(bool), _438_v4.F12, (_438_v4.F12).Times((_dafny.Zero).Minus(_434_v0)))
				var _451_v17 _dafny.Array
				_ = _451_v17
				var _nw68 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
				_ = _nw68
				_451_v17 = _nw68
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_451_v17), 0))
				_ = _index48
				(_451_v17).ArraySet1(_438_v4, (_index48).Int())
				var _452_v18 _dafny.MultiSet
				_ = _452_v18
				_452_v18 = _dafny.MultiSetOf(p1, _dafny.Companion_Sequence_.Equal(_435_v1, _435_v1))
				var _453_v19 _dafny.Map
				_ = _453_v19
				_453_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_v7, (_444_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))).Int()).(bool))
				var _454_v20 D7
				_ = _454_v20
				_454_v20 = Companion_D7_.Create_DC18_(_453_v19)
				var _455_v21 _dafny.Sequence
				_ = _455_v21
				_455_v21 = _dafny.SeqOf((_444_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))).Int()).(bool), p1, p1, (_this).F9(), false)
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))
				_ = _index49
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_449_v15), 0))
				_ = _index50
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_451_v17), 0))
				_ = _index51
				var _rhs88 bool = (_444_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))).Int()).(bool)
				_ = _rhs88
				var _rhs89 _dafny.Int = _438_v4.F12
				_ = _rhs89
				var _rhs90 *C0 = _438_v4
				_ = _rhs90
				var _rhs91 _dafny.MultiSet = ((_452_v18).Intersection(_452_v18)).Union((_dafny.MultiSetFromSeq(_455_v21)).Difference(_dafny.MultiSetOf((_this).F9())))
				_ = _rhs91
				var _rhs92 D7 = _454_v20
				_ = _rhs92
				var _lhs47 _dafny.Array = _444_v10
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_444_v10), 0))
				_ = _lhs48
				var _lhs49 _dafny.Array = _449_v15
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_449_v15), 0))
				_ = _lhs50
				var _lhs51 _dafny.Array = _451_v17
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_451_v17), 0))
				_ = _lhs52
				(_lhs47).ArraySet1(_rhs88, (_lhs48).Int())
				(_lhs49).ArraySet1(_rhs89, (_lhs50).Int())
				(_lhs51).ArraySet1(_rhs90, (_lhs52).Int())
				_452_v18 = _rhs91
				_454_v20 = _rhs92
				_455_v21 = _455_v21
			} else {
				var _456_v22 _dafny.Map
				_ = _456_v22
				_456_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_435_v1, _444_v10)
				_456_v22 = _456_v22
				var _457_v23 _dafny.Array
				_ = _457_v23
				var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
				_ = _nw69
				_457_v23 = _nw69
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
				_ = _nw70
				_457_v23 = _nw70
				var _458_v24 _dafny.Set
				_ = _458_v24
				_458_v24 = _dafny.SetOf(!(true), (_dafny.MultiSetOf((_dafny.Zero).Minus(_438_v4.F12))).IsDisjointFrom(_441_v7), p1, p1, p1)
				_458_v24 = _458_v24
				(globalState).F2 = p1
				var _459_v25 _dafny.Map
				_ = _459_v25
				_459_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
				var _460_v26 _dafny.Array
				_ = _460_v26
				var _nwElement0_9 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_436_v2, (Companion_Default___.SafeIndex(_438_v4.F12, _dafny.IntOfUint32((_436_v2).Cardinality()))).Uint32(), _dafny.IntOfUint32((_436_v2).Cardinality()))).Cardinality())
				_ = _nwElement0_9
				var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(21))
				_ = _nw71
				(_nw71).ArraySet1(_nwElement0_9, 0)
				(_nw71).ArraySet1(_434_v0, 1)
				(_nw71).ArraySet1(_dafny.IntOfInt64(-78), 2)
				(_nw71).ArraySet1(_438_v4.F12, 3)
				(_nw71).ArraySet1(_434_v0, 4)
				(_nw71).ArraySet1(p0, 5)
				(_nw71).ArraySet1(p0, 6)
				(_nw71).ArraySet1(_438_v4.F12, 7)
				(_nw71).ArraySet1(Companion_Default___.Fm0(globalState), 8)
				(_nw71).ArraySet1(_438_v4.F12, 9)
				(_nw71).ArraySet1(_434_v0, 10)
				(_nw71).ArraySet1(_438_v4.F12, 11)
				(_nw71).ArraySet1((_this).F14(), 12)
				(_nw71).ArraySet1((_459_v25).Cardinality(), 13)
				(_nw71).ArraySet1(_438_v4.F12, 14)
				(_nw71).ArraySet1((_this).F14(), 15)
				(_nw71).ArraySet1(p0, 16)
				(_nw71).ArraySet1(p0, 17)
				(_nw71).ArraySet1(_434_v0, 18)
				(_nw71).ArraySet1(_438_v4.F12, 19)
				(_nw71).ArraySet1(_dafny.IntOfInt64(976), 20)
				_460_v26 = _nw71
				var _461_v27 _dafny.MultiSet
				_ = _461_v27
				_461_v27 = _dafny.MultiSetOf(_460_v26, _460_v26)
				(_this).M8(!((_this).F9()), _461_v27, globalState)
			}
			(_438_v4).F12 = _438_v4.F12
		} else {
			var _462_v28 _dafny.Sequence
			_ = _462_v28
			_462_v28 = _dafny.SeqOf((_this).F9())
			var _463_v29 _dafny.Array
			_ = _463_v29
			var _nwElement0_10 _dafny.Int = (_this).F14()
			_ = _nwElement0_10
			var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(21))
			_ = _nw72
			(_nw72).ArraySet1(_nwElement0_10, 0)
			(_nw72).ArraySet1(p0, 1)
			(_nw72).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_462_v28).Cardinality())), 2)
			(_nw72).ArraySet1(_438_v4.F12, 3)
			(_nw72).ArraySet1(_438_v4.F12, 4)
			(_nw72).ArraySet1((_this).F14(), 5)
			(_nw72).ArraySet1(p0, 6)
			(_nw72).ArraySet1(_438_v4.F12, 7)
			(_nw72).ArraySet1(_dafny.IntOfUint32((_435_v1).Cardinality()), 8)
			(_nw72).ArraySet1(_438_v4.F12, 9)
			(_nw72).ArraySet1((_this).F14(), 10)
			(_nw72).ArraySet1(_dafny.IntOfUint32((_436_v2).Cardinality()), 11)
			(_nw72).ArraySet1(_434_v0, 12)
			(_nw72).ArraySet1(_438_v4.F12, 13)
			(_nw72).ArraySet1(Companion_Default___.Fm0(globalState), 14)
			(_nw72).ArraySet1(_dafny.IntOfInt64(958), 15)
			(_nw72).ArraySet1(_438_v4.F12, 16)
			(_nw72).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), 17)
			(_nw72).ArraySet1(_438_v4.F12, 18)
			(_nw72).ArraySet1(_dafny.IntOfUint32((_435_v1).Cardinality()), 19)
			(_nw72).ArraySet1(_dafny.IntOfInt64(55), 20)
			_463_v29 = _nw72
			var _464_v30 _dafny.Sequence
			_ = _464_v30
			_464_v30 = _dafny.SeqOf(_463_v29)
			(_438_v4).F12 = (func() _dafny.Int {
				if (_this).F9() {
					return (Companion_Default___.Fm24(_438_v4.F12, !(true), (_this).F9(), p1, globalState)).Cardinality()
				}
				return (_dafny.IntOfUint32((_464_v30).Cardinality())).Times(_438_v4.F12)
			})()
			(globalState).F2 = (_this).F9()
			var _465_v31 _dafny.MultiSet
			_ = _465_v31
			_465_v31 = _dafny.MultiSetOf(p2, p2)
			var _466_v32 D1
			_ = _466_v32
			_466_v32 = Companion_D1_.Create_DC3_((_this).F9(), _dafny.UnicodeSeqOfUtf8Bytes("doecrfd"), p0)
			var _rhs93 _dafny.Sequence = _435_v1
			_ = _rhs93
			var _rhs94 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_436_v2, (_this).Fm2(_438_v4.F12, p1, (_this).F9(), (_465_v31).Cardinality(), globalState)), (Companion_Default___.SafeIndex((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("e"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("e")).Cardinality()))).Uint32(), _440_v6), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-52), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("e"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("e")).Cardinality()))).Uint32(), _440_v6)).Cardinality()))).Uint32(), _440_v6)).Cardinality())).Minus(_438_v4.F12), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_436_v2, (_this).Fm2(_438_v4.F12, p1, (_this).F9(), (_465_v31).Cardinality(), globalState))).Cardinality()))).Uint32(), (_this).F14())
			_ = _rhs94
			var _rhs95 bool = !(p1)
			_ = _rhs95
			var _rhs96 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_466_v32).Dtor_cf7(), _435_v1)
			_ = _rhs96
			var _lhs53 *GlobalState = globalState
			_ = _lhs53
			_435_v1 = _rhs93
			_436_v2 = _rhs94
			_lhs53.F2 = _rhs95
			_435_v1 = _rhs96
			var _467_v33 D4
			_ = _467_v33
			_467_v33 = Companion_D4_.Create_DC12_((_this).F9(), p1, (_this).F14(), _dafny.IntOfUint32((_436_v2).Cardinality()), (_this).F9())
			var _468_v34 D1
			_ = _468_v34
			_468_v34 = Companion_D1_.Create_DC2_(Companion_Default___.SafeModuloInt((_this).F14(), _438_v4.F12), (_462_v28).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_462_v28).Cardinality()))).Uint32()).(bool), (func() _dafny.Int {
				if p1 {
					return (_dafny.Zero).Minus(_438_v4.F12)
				}
				return _434_v0
			})(), ((_this).F14()).Minus((_467_v33).Dtor_cf23()))
			var _469_v35 _dafny.Map
			_ = _469_v35
			_469_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _470_v36 _dafny.Set
			_ = _470_v36
			_470_v36 = _dafny.SetOf((_469_v35).Cardinality())
			var _471_v37 _dafny.MultiSet
			_ = _471_v37
			_471_v37 = _dafny.MultiSetOf(_463_v29, _463_v29)
			var _pat_let_tv6 = _471_v37
			_ = _pat_let_tv6
			var _pat_let_tv7 = _471_v37
			_ = _pat_let_tv7
			var _pat_let_tv8 = p1
			_ = _pat_let_tv8
			_468_v34 = func(_pat_let8_0 D1) D1 {
				return func(_472_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let9_0 bool) D1 {
						return func(_473_dt__update_hcf3_h0 bool) D1 {
							return func(_pat_let10_0 _dafny.Int) D1 {
								return func(_474_dt__update_hcf4_h0 _dafny.Int) D1 {
									return Companion_D1_.Create_DC2_((_472_dt__update__tmp_h0).Dtor_cf2(), _473_dt__update_hcf3_h0, _474_dt__update_hcf4_h0, (_472_dt__update__tmp_h0).Dtor_cf5())
								}(_pat_let10_0)
							}(_dafny.IntOfUint32((_dafny.SeqOf((_this).F9(), true, _pat_let_tv8)).Cardinality()))
						}(_pat_let9_0)
					}((_pat_let_tv6).IsProperSubsetOf(_pat_let_tv7))
				}(_pat_let8_0)
			}(Companion_Default___.Fm25((_470_v36).Cardinality(), p1, p1, _441_v7, globalState))
			var _475_v38 _dafny.Map
			_ = _475_v38
			_475_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2) == (!(p1)), (_this).F14())
			_475_v38 = (_475_v38).Update(p1, Companion_Default___.Fm0(globalState))
		}
		var _476_v39 _dafny.Array
		_ = _476_v39
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_11
		var _nw73 _dafny.Array
		_ = _nw73
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw73 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) bool = (func(_477_p2 bool) func(_dafny.Int) bool {
				return func(_478_i0 _dafny.Int) bool {
					return _477_p2
				}
			})(p2)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw73 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw73).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw73).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_476_v39 = _nw73
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_476_v39), 0))
		_ = _index52
		(_476_v39).ArraySet1(p1, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_476_v39), 0))
		_ = _index53
		(_476_v39).ArraySet1(!((_this).F9()), (_index53).Int())
		var _479_v40 _dafny.Array
		_ = _479_v40
		var _nw74 _dafny.Array = _dafny.NewArray(_dafny.One)
		_ = _nw74
		_479_v40 = _nw74
		var _480_v41 D4
		_ = _480_v41
		_480_v41 = Companion_D4_.Create_DC11_(_479_v40)
		var _481_v42 _dafny.Array
		_ = _481_v42
		var _nwElement0_11 _dafny.Array = _479_v40
		_ = _nwElement0_11
		var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(7))
		_ = _nw75
		(_nw75).ArraySet1(_nwElement0_11, 0)
		(_nw75).ArraySet1((_480_v41).Dtor_cf19(), 1)
		(_nw75).ArraySet1(_479_v40, 2)
		(_nw75).ArraySet1((_480_v41).Dtor_cf19(), 3)
		(_nw75).ArraySet1(_479_v40, 4)
		(_nw75).ArraySet1(_479_v40, 5)
		(_nw75).ArraySet1(_479_v40, 6)
		_481_v42 = _nw75
		var _source6 _dafny.Array = _481_v42
		_ = _source6
		var _482___mcc_h0 _dafny.Array = _source6
		_ = _482___mcc_h0
		var _483_cf30 _dafny.Array = _482___mcc_h0
		_ = _483_cf30
		(_438_v4).F12 = _434_v0
		var _484_v43 D4
		_ = _484_v43
		_484_v43 = Companion_D4_.Create_DC12_(!((_this).F9()), p2, (_dafny.Zero).Minus(_438_v4.F12), (_this).F14(), false)
		(_438_v4).F12 = ((_484_v43).Dtor_cf23()).Plus(Companion_Default___.SafeModuloInt(_438_v4.F12, _dafny.IntOfUint32((_435_v1).Cardinality())))
		var _485_v44 _dafny.Sequence
		_ = _485_v44
		_485_v44 = _dafny.SeqOf(_438_v4.F11, _438_v4.F11)
		var _486_v46 *C0
		_ = _486_v46
		var _nw76 *C0 = New_C0_()
		_ = _nw76
		_nw76.Ctor__(func() _dafny.Set {
			var _coll19 = _dafny.NewBuilder()
			_ = _coll19
			for _iter22 := _dafny.Iterate(((_485_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_435_v1).Cardinality()), _dafny.IntOfUint32((_485_v44).Cardinality()))).Uint32()).(_dafny.Set)).Elements()); ; {
				_compr_19, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _487_v45 _dafny.Sequence
				_487_v45 = interface{}(_compr_19).(_dafny.Sequence)
				if ((_485_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_435_v1).Cardinality()), _dafny.IntOfUint32((_485_v44).Cardinality()))).Uint32()).(_dafny.Set)).Contains(_487_v45) {
					_coll19.Add(_487_v45)
				}
			}
			return _coll19.ToSet()
		}(), _438_v4.F12)
		_486_v46 = _nw76
		var _488_v47 _dafny.Map
		_ = _488_v47
		_488_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_440_v6, p2)
		var _489_v48 _dafny.Set
		_ = _489_v48
		_489_v48 = _dafny.SetOf(_488_v47, _488_v47)
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_476_v39), 0))
		_ = _index54
		(_476_v39).ArraySet1(!(((_489_v48).Difference(_489_v48)).IsSubsetOf(_489_v48)), (_index54).Int())
		var _490_v49 _dafny.Set
		_ = _490_v49
		_490_v49 = _dafny.SetOf((_this).F14(), (_this).F14())
		var _491_v50 _dafny.Sequence
		_ = _491_v50
		_491_v50 = _dafny.SeqOf((_this).F9())
		var _492_v51 _dafny.MultiSet
		_ = _492_v51
		_492_v51 = _dafny.MultiSetOf((_476_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_476_v39), 0))).Int()).(bool))
		var _493_v52 _dafny.Array
		_ = _493_v52
		var _nwElement0_12 _dafny.Int = (_434_v0).Plus(_dafny.IntOfUint32((_491_v50).Cardinality()))
		_ = _nwElement0_12
		var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(17))
		_ = _nw77
		(_nw77).ArraySet1(_nwElement0_12, 0)
		(_nw77).ArraySet1(p0, 1)
		(_nw77).ArraySet1(_dafny.IntOfUint32((_435_v1).Cardinality()), 2)
		(_nw77).ArraySet1(Companion_Default___.SafeModuloInt((_492_v51).Cardinality(), _dafny.IntOfUint32((_436_v2).Cardinality())), 3)
		(_nw77).ArraySet1(p0, 4)
		(_nw77).ArraySet1(Companion_Default___.SafeModuloInt((_this).F14(), p0), 5)
		(_nw77).ArraySet1(_434_v0, 6)
		(_nw77).ArraySet1(_434_v0, 7)
		(_nw77).ArraySet1((_434_v0).Plus((_this).F14()), 8)
		(_nw77).ArraySet1(_438_v4.F12, 9)
		(_nw77).ArraySet1((_this).F14(), 10)
		(_nw77).ArraySet1(_434_v0, 11)
		(_nw77).ArraySet1(Companion_Default___.Fm0(globalState), 12)
		(_nw77).ArraySet1(_434_v0, 13)
		(_nw77).ArraySet1(Companion_Default___.SafeModuloInt(_438_v4.F12, _438_v4.F12), 14)
		(_nw77).ArraySet1(_434_v0, 15)
		(_nw77).ArraySet1((func() _dafny.Int {
			if (_476_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_476_v39), 0))).Int()).(bool) {
				return (_this).F14()
			}
			return (_dafny.Zero).Minus(Companion_Default___.Fm0(globalState))
		})(), 16)
		_493_v52 = _nw77
		var _494_v53 *C1
		_ = _494_v53
		var _nw78 *C1 = New_C1_()
		_ = _nw78
		_nw78.Ctor__(Companion_Default___.Fm26(globalState))
		_494_v53 = _nw78
		var _495_v54 _dafny.Map
		_ = _495_v54
		_495_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(414))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}((func(_496_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_497_i1 _dafny.Int) _dafny.CodePoint {
				return _496_v6
			}
		})(_440_v6))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_494_v53, _438_v4.F12))
		var _498_v55 _dafny.Map
		_ = _498_v55
		_498_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_494_v53, p0)
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_493_v52), 0))
		_ = _index55
		(_493_v52).ArraySet1(((func() _dafny.Map {
			if (_495_v54).Contains(_dafny.UnicodeSeqOfUtf8Bytes("nukgadch")) {
				return (_495_v54).Get(_dafny.UnicodeSeqOfUtf8Bytes("nukgadch")).(_dafny.Map)
			}
			return _498_v55
		})()).Cardinality(), (_index55).Int())
		var _499_v56 _dafny.Sequence
		_ = _499_v56
		_499_v56 = _dafny.SeqOf(_dafny.SetOf(_434_v0), _490_v49, _490_v49)
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_493_v52), 0))
		_ = _index56
		var _rhs97 _dafny.Set = (_490_v49).Difference((_499_v56).Select((Companion_Default___.SafeIndex(_434_v0, _dafny.IntOfUint32((_499_v56).Cardinality()))).Uint32()).(_dafny.Set))
		_ = _rhs97
		var _rhs98 _dafny.Int = p0
		_ = _rhs98
		var _lhs54 _dafny.Array = _493_v52
		_ = _lhs54
		var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_493_v52), 0))
		_ = _lhs55
		_490_v49 = _rhs97
		(_lhs54).ArraySet1(_rhs98, (_lhs55).Int())
	}
}
func (_this *C2) M8(p0 bool, p1 _dafny.MultiSet, globalState *GlobalState) {
	{
		var _500_v0 _dafny.MultiSet
		_ = _500_v0
		_500_v0 = _dafny.MultiSetOf((_this).F14(), (_this).F14())
		var _501_v1 _dafny.Sequence
		_ = _501_v1
		_501_v1 = _dafny.SeqOf(p0)
		var _502_v2 _dafny.Set
		_ = _502_v2
		_502_v2 = _dafny.SetOf((_this).F14())
		var _503_v3 _dafny.Map
		_ = _503_v3
		_503_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F9())
		var _504_v4 _dafny.MultiSet
		_ = _504_v4
		_504_v4 = _dafny.MultiSetOf((func() bool {
			if (_503_v3).Contains((_this).F9()) {
				return (_503_v3).Get((_this).F9()).(bool)
			}
			return (_this).F9()
		})(), (_this).F9())
		var _505_v5 _dafny.Array
		_ = _505_v5
		var _nwElement0_13 bool = (_500_v0).IsProperSubsetOf(_500_v0)
		_ = _nwElement0_13
		var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(3))
		_ = _nw79
		(_nw79).ArraySet1(_nwElement0_13, 0)
		(_nw79).ArraySet1(((_this).F14()).Cmp((_this).F14()) >= 0, 1)
		(_nw79).ArraySet1(((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_501_v1, (Companion_Default___.SafeIndex((_this).F14(), _dafny.IntOfUint32((_501_v1).Cardinality()))).Uint32(), (_this).F9()))).Update((_this).F9(), Companion_Default___.Abs((_502_v2).Cardinality()))).IsDisjointFrom(_504_v4), 2)
		_505_v5 = _nw79
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))
		_ = _index57
		(_505_v5).ArraySet1((_this).F9(), (_index57).Int())
		var _506_v6 _dafny.Array
		_ = _506_v6
		var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
		_ = _nw80
		_506_v6 = _nw80
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
		_ = _index58
		(_506_v6).ArraySet1(_dafny.IntOfInt64(-89), (_index58).Int())
		var _507_v7 _dafny.Array
		_ = _507_v7
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(12))
		_ = _nw81
		_507_v7 = _nw81
		var _508_v8 _dafny.Sequence
		_ = _508_v8
		_508_v8 = _dafny.SeqOf(_507_v7)
		var _509_v9 _dafny.Array
		_ = _509_v9
		var _nwElement0_14 _dafny.Array = _507_v7
		_ = _nwElement0_14
		var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(16))
		_ = _nw82
		(_nw82).ArraySet1(_nwElement0_14, 0)
		(_nw82).ArraySet1(_507_v7, 1)
		(_nw82).ArraySet1(_507_v7, 2)
		(_nw82).ArraySet1(_507_v7, 3)
		(_nw82).ArraySet1(_507_v7, 4)
		(_nw82).ArraySet1(_507_v7, 5)
		(_nw82).ArraySet1(_507_v7, 6)
		(_nw82).ArraySet1(_507_v7, 7)
		(_nw82).ArraySet1((_508_v8).Select((Companion_Default___.SafeIndex((_this).F14(), _dafny.IntOfUint32((_508_v8).Cardinality()))).Uint32()).(_dafny.Array), 8)
		(_nw82).ArraySet1(_507_v7, 9)
		(_nw82).ArraySet1(_507_v7, 10)
		(_nw82).ArraySet1(_507_v7, 11)
		(_nw82).ArraySet1(_507_v7, 12)
		(_nw82).ArraySet1(_507_v7, 13)
		(_nw82).ArraySet1(_507_v7, 14)
		(_nw82).ArraySet1(_507_v7, 15)
		_509_v9 = _nw82
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))
		_ = _index59
		(_509_v9).ArraySet1(_507_v7, (_index59).Int())
		var _510_v10 _dafny.Sequence
		_ = _510_v10
		_510_v10 = _dafny.SeqOf((_this).F14())
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))
		_ = _index60
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
		_ = _index61
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))
		_ = _index62
		var _rhs99 bool = (func() bool {
			if (_this).F9() {
				return (_this).F9()
			}
			return p0
		})()
		_ = _rhs99
		var _rhs100 bool = Companion_Default___.Fm1(_500_v0, (_dafny.Zero).Minus((_this).F14()), globalState)
		_ = _rhs100
		var _rhs101 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F14(), (_dafny.Zero).Minus((_dafny.IntOfUint32((_510_v10).Cardinality())).Minus((_this).F14())))
		_ = _rhs101
		var _rhs102 _dafny.Array = _507_v7
		_ = _rhs102
		var _lhs56 *GlobalState = globalState
		_ = _lhs56
		var _lhs57 _dafny.Array = _505_v5
		_ = _lhs57
		var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))
		_ = _lhs58
		var _lhs59 _dafny.Array = _506_v6
		_ = _lhs59
		var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
		_ = _lhs60
		var _lhs61 _dafny.Array = _509_v9
		_ = _lhs61
		var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))
		_ = _lhs62
		_lhs56.F2 = _rhs99
		(_lhs57).ArraySet1(_rhs100, (_lhs58).Int())
		(_lhs59).ArraySet1(_rhs101, (_lhs60).Int())
		(_lhs61).ArraySet1(_rhs102, (_lhs62).Int())
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
		_ = _index63
		(_506_v6).ArraySet1((_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int), (_index63).Int())
		var _511_v11 D1
		_ = _511_v11
		_511_v11 = Companion_D1_.Create_DC5_(Companion_D1_.Create_DC4_((_this).F14(), (_this).F9(), true, (Companion_D3_.Create_DC9_((_this).F9(), _dafny.SeqOf(p0, p0))).Dtor_cf17(), (_this).F9()))
		var _512_v12 D1
		_ = _512_v12
		_512_v12 = Companion_D1_.Create_DC5_(_511_v11)
		var _pat_let_tv9 = _511_v11
		_ = _pat_let_tv9
		var _pat_let_tv10 = _511_v11
		_ = _pat_let_tv10
		var _513_v13 _dafny.Map
		_ = _513_v13
		_513_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(func(_pat_let11_0 D1) D1 {
			return func(_514_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let12_0 D1) D1 {
					return func(_515_dt__update_hcf14_h0 D1) D1 {
						return Companion_D1_.Create_DC5_(_515_dt__update_hcf14_h0)
					}(_pat_let12_0)
				}(_pat_let_tv9)
			}(_pat_let11_0)
		}(_512_v12), _512_v12, func(_pat_let13_0 D1) D1 {
			return func(_516_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let14_0 D1) D1 {
					return func(_517_dt__update_hcf14_h1 D1) D1 {
						return Companion_D1_.Create_DC5_(_517_dt__update_hcf14_h1)
					}(_pat_let14_0)
				}(_pat_let_tv10)
			}(_pat_let13_0)
		}(_512_v12)), _500_v0)
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))).Int()))
		_ = _arr0
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))).Int()))), 0))
		_ = _index64
		(_arr0).ArraySet1((func() _dafny.MultiSet {
			if (_513_v13).Contains(_dafny.MultiSetOf(_512_v12, _512_v12, _512_v12)) {
				return (_513_v13).Get(_dafny.MultiSetOf(_512_v12, _512_v12, _512_v12)).(_dafny.MultiSet)
			}
			return _dafny.MultiSetOf((_this).F14(), (_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int))
		})(), (_index64).Int())
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))).Int()))
		_ = _arr1
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))).Int()))), 0))
		_ = _index65
		(_arr1).ArraySet1((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_510_v10, _dafny.SeqOf((_this).F14())))).Difference(_500_v0), (_index65).Int())
		var _518_i0 _dafny.Int
		_ = _518_i0
		_518_i0 = _dafny.Zero
		{
			for (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool) {
				{
					if (_518_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_518_i0 = (_518_i0).Plus(_dafny.One)
					var _519_v14 _dafny.Sequence
					_ = _519_v14
					_519_v14 = _dafny.UnicodeSeqOfUtf8Bytes("a")
					var _520_v15 D1
					_ = _520_v15
					_520_v15 = Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_519_v14).Cardinality()), !(p0), p0, (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool), (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool))
					if Companion_Default___.Fm1(((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))).Int()))), 0))).Int()).(_dafny.MultiSet)).Union(_500_v0), Companion_Default___.SafeDivisionInt((_520_v15).Dtor_cf9(), (_this).F14()), globalState) {
						var _521_v16 _dafny.Map
						_ = _521_v16
						_521_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_500_v0, (_this).F9())
						var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
						_ = _index66
						(_506_v6).ArraySet1(((_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int)).Plus((_dafny.Zero).Minus((_dafny.IntOfUint32((_519_v14).Cardinality())).Minus((_521_v16).Cardinality()))), (_index66).Int())
						var _522_v17 D3
						_ = _522_v17
						_522_v17 = Companion_D3_.Create_DC9_((_this).F9(), _501_v1)
						var _523_v18 _dafny.MultiSet
						_ = _523_v18
						_523_v18 = _dafny.MultiSetOf(_505_v5)
						var _524_v19 D3
						_ = _524_v19
						_524_v19 = Companion_D3_.Create_DC7_(_523_v18)
						var _525_v20 _dafny.Map
						_ = _525_v20
						_525_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_519_v14, _519_v14), p0)
						var _526_v21 D9
						_ = _526_v21
						_526_v21 = Companion_D9_.Create_DC23_(_525_v20)
						var _pat_let_tv11 = _523_v18
						_ = _pat_let_tv11
						var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
						_ = _index67
						var _rhs103 _dafny.Int = ((_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int)).Times((_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int))
						_ = _rhs103
						var _rhs104 D3 = _522_v17
						_ = _rhs104
						var _rhs105 D3 = func(_pat_let15_0 D3) D3 {
							return func(_527_dt__update__tmp_h2 D3) D3 {
								return func(_pat_let16_0 _dafny.MultiSet) D3 {
									return func(_528_dt__update_hcf16_h0 _dafny.MultiSet) D3 {
										return Companion_D3_.Create_DC7_(_528_dt__update_hcf16_h0)
									}(_pat_let16_0)
								}(_pat_let_tv11)
							}(_pat_let15_0)
						}(_524_v19)
						_ = _rhs105
						var _rhs106 _dafny.Map = (_526_v21).Dtor_cf40()
						_ = _rhs106
						var _lhs63 _dafny.Array = _506_v6
						_ = _lhs63
						var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
						_ = _lhs64
						(_lhs63).ArraySet1(_rhs103, (_lhs64).Int())
						_522_v17 = _rhs104
						_524_v19 = _rhs105
						_525_v20 = _rhs106
						var _529_v22 _dafny.CodePoint
						_ = _529_v22
						_529_v22 = _dafny.CodePoint('e')
						var _530_v23 _dafny.Sequence
						_ = _530_v23
						_530_v23 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_519_v14, (Companion_Default___.SafeIndex((_this).F14(), _dafny.IntOfUint32((_519_v14).Cardinality()))).Uint32(), _529_v22))
						var _531_v24 _dafny.Map
						_ = _531_v24
						_531_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_500_v0).Update((_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int), Companion_Default___.Abs(_dafny.IntOfUint32((_530_v23).Cardinality()))), (_dafny.Zero).Minus((_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int)))
						var _532_v25 _dafny.Map
						_ = _532_v25
						_532_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v24, (_this).F9())
						_532_v25 = (_532_v25).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(218), _dafny.ArrayLen((_509_v9), 0))).Int()))), 0))).Int()).(_dafny.MultiSet), (_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int)), (_this).F9())
						var _533_v26 _dafny.Array
						_ = _533_v26
						var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(13))
						_ = _nw83
						_533_v26 = _nw83
						var _534_v27 _dafny.Map
						_ = _534_v27
						_534_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int), _529_v22)
						var _535_v28 _dafny.Set
						_ = _535_v28
						_535_v28 = _dafny.SetOf(_534_v27, _534_v27, _534_v27)
						var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_533_v26), 0))
						_ = _index68
						(_533_v26).ArraySet1(_535_v28, (_index68).Int())
						var _536_v29 _dafny.Array
						_ = _536_v29
						var _len0_12 _dafny.Int = _dafny.IntOfInt64(15)
						_ = _len0_12
						var _nw84 _dafny.Array
						_ = _nw84
						if _len0_12.Cmp(_dafny.Zero) == 0 {
							_nw84 = _dafny.NewArray(_len0_12)
						} else {
							var _init12 func(_dafny.Int) _dafny.Sequence = (func(_537_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_538_i1 _dafny.Int) _dafny.Sequence {
									return _537_v14
								}
							})(_519_v14)
							_ = _init12
							var _element0_12 = _init12(_dafny.Zero)
							_ = _element0_12
							_nw84 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
							(_nw84).ArraySet1(_element0_12, 0)
							var _nativeLen0_12 = (_len0_12).Int()
							_ = _nativeLen0_12
							for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
								(_nw84).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
							}
						}
						_536_v29 = _nw84
						var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_536_v29), 0))
						_ = _index69
						(_536_v29).ArraySet1(_519_v14, (_index69).Int())
						var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_533_v26), 0))
						_ = _index70
						var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_536_v29), 0))
						_ = _index71
						var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))
						_ = _index72
						var _rhs107 bool = (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool)
						_ = _rhs107
						var _rhs108 _dafny.Set = (_535_v28).Intersection((_535_v28).Intersection(_535_v28))
						_ = _rhs108
						var _rhs109 _dafny.Sequence = _519_v14
						_ = _rhs109
						var _rhs110 _dafny.CodePoint = (func() _dafny.CodePoint {
							if (_this).F9() {
								return _dafny.CodePoint('r')
							}
							return _529_v22
						})()
						_ = _rhs110
						var _rhs111 bool = !((func() bool {
							if (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool) {
								return p0
							}
							return (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool)
						})())
						_ = _rhs111
						var _lhs65 *GlobalState = globalState
						_ = _lhs65
						var _lhs66 _dafny.Array = _533_v26
						_ = _lhs66
						var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_533_v26), 0))
						_ = _lhs67
						var _lhs68 _dafny.Array = _536_v29
						_ = _lhs68
						var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_536_v29), 0))
						_ = _lhs69
						var _lhs70 _dafny.Array = _505_v5
						_ = _lhs70
						var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))
						_ = _lhs71
						_lhs65.F2 = _rhs107
						(_lhs66).ArraySet1(_rhs108, (_lhs67).Int())
						(_lhs68).ArraySet1(_rhs109, (_lhs69).Int())
						_529_v22 = _rhs110
						(_lhs70).ArraySet1(_rhs111, (_lhs71).Int())
						(globalState).F2 = (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool)
					} else {
						_519_v14 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(337))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg33 _dafny.Int) interface{} {
								return coer33(arg33)
							}
						}(func(_539_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('k')
						})), _519_v14)
						var _540_v30 _dafny.CodePoint
						_ = _540_v30
						_540_v30 = _dafny.CodePoint('w')
						(globalState).F2 = !(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_519_v14, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_519_v14).Cardinality()), _dafny.IntOfUint32((_519_v14).Cardinality()))).Uint32(), _540_v30), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg34 _dafny.Int) interface{} {
								return coer34(arg34)
							}
						}((func(_541_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_542_i3 _dafny.Int) _dafny.CodePoint {
								return _541_v30
							}
						})(_540_v30)))), (Companion_Default___.SafeIndex((_this).F14(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_519_v14, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_519_v14).Cardinality()), _dafny.IntOfUint32((_519_v14).Cardinality()))).Uint32(), _540_v30), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(766))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg35 _dafny.Int) interface{} {
								return coer35(arg35)
							}
						}((func(_543_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_544_i3 _dafny.Int) _dafny.CodePoint {
								return _543_v30
							}
						})(_540_v30))))).Cardinality()))).Uint32(), _540_v30), _dafny.Companion_Sequence_.Concatenate(_519_v14, _519_v14)))
						var _545_v31 _dafny.Array
						_ = _545_v31
						_545_v31 = _506_v6
						var _546_v32 _dafny.Map
						_ = _546_v32
						_546_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool), (_this).F14())
						var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
						_ = _index73
						var _rhs112 _dafny.Int = (func() _dafny.Int {
							if (_546_v32).Contains(p0) {
								return (_546_v32).Get(p0).(_dafny.Int)
							}
							return _dafny.IntOfInt64(738)
						})()
						_ = _rhs112
						var _rhs113 _dafny.Array = _545_v31
						_ = _rhs113
						var _lhs72 _dafny.Array = _506_v6
						_ = _lhs72
						var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
						_ = _lhs73
						(_lhs72).ArraySet1(_rhs112, (_lhs73).Int())
						_545_v31 = _rhs113
						var _547_v33 T0
						_ = _547_v33
						var _nw85 *C1 = New_C1_()
						_ = _nw85
						_nw85.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(324))).Uint32(), func(coer36 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
							return func(arg36 _dafny.Int) interface{} {
								return coer36(arg36)
							}
						}((func(_548_p0 bool, _549_v6 _dafny.Array) func(_dafny.Int) D1 {
							return func(_550_i4 _dafny.Int) D1 {
								return Companion_D1_.Create_DC2_((_dafny.SetOf(_548_p0)).Cardinality(), _548_p0, (_549_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_549_v6), 0))).Int()).(_dafny.Int), (_this).F14())
							}
						})(p0, _506_v6))))
						_547_v33 = _nw85
						var _551_v34 _dafny.MultiSet
						_ = _551_v34
						_551_v34 = _dafny.MultiSetOf(_547_v33)
						var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
						_ = _index74
						(_506_v6).ArraySet1((func() _dafny.Int {
							if (_551_v34).IsSubsetOf(_551_v34) {
								return (_this).F14()
							}
							return (_this).F14()
						})(), (_index74).Int())
						var _552_v35 _dafny.Set
						_ = _552_v35
						_552_v35 = _dafny.SetOf(_506_v6)
						var _553_v36 _dafny.Map
						_ = _553_v36
						_553_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_552_v35).Intersection(_552_v35), (_this).F14())
						var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))
						_ = _index75
						(_506_v6).ArraySet1((func() _dafny.Int {
							if (_553_v36).Contains(_552_v35) {
								return (_553_v36).Get(_552_v35).(_dafny.Int)
							}
							return (_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int)
						})(), (_index75).Int())
					}
					_506_v6 = _506_v6
					var _554_v37 D3
					_ = _554_v37
					_554_v37 = Companion_D3_.Create_DC10_()
					_554_v37 = Companion_Default___.Fm27(globalState)
					(globalState).F2 = (_this).F9()
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _555_v38 _dafny.CodePoint
		_ = _555_v38
		_555_v38 = _dafny.CodePoint('k')
		var _556_v39 _dafny.Set
		_ = _556_v39
		_556_v39 = _dafny.SetOf(_555_v38, _dafny.CodePoint('o'), _555_v38, _555_v38, _555_v38)
		_556_v39 = _556_v39
		var _557_v40 D4
		_ = _557_v40
		_557_v40 = Companion_D4_.Create_DC14_((_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool), false, (_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int))
		var _558_v41 _dafny.MultiSet
		_ = _558_v41
		_558_v41 = _dafny.MultiSetOf(_555_v38, _555_v38)
		var _559_v42 D1
		_ = _559_v42
		_559_v42 = Companion_D1_.Create_DC3_((_this).F9(), Companion_Default___.Fm28((_558_v41).Cardinality(), true, globalState), (_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int))
		var _560_v43 _dafny.MultiSet
		_ = _560_v43
		_560_v43 = _dafny.MultiSetOf(_559_v42)
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))
		_ = _index76
		var _rhs114 D4 = Companion_D4_.Create_DC14_((func() bool {
			if (_this).F9() {
				return (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool)
			}
			return !(p0)
		})(), (_505_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))).Int()).(bool), (_506_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_506_v6), 0))).Int()).(_dafny.Int))
		_ = _rhs114
		var _rhs115 _dafny.MultiSet = (_dafny.MultiSetOf(_559_v42)).Update(_559_v42, Companion_Default___.Abs((_this).F14()))
		_ = _rhs115
		var _rhs116 bool = p0
		_ = _rhs116
		var _lhs74 _dafny.Array = _505_v5
		_ = _lhs74
		var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_505_v5), 0))
		_ = _lhs75
		_557_v40 = _rhs114
		_560_v43 = _rhs115
		(_lhs74).ArraySet1(_rhs116, (_lhs75).Int())
	}
}
func (_this *C2) F14() _dafny.Int {
	{
		return _this._f14
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f9  bool
	F16  _dafny.Array
	_f15 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f9 = false
	_this.F16 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = false
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C3{}
var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F9() bool {
	return _this._f9
}
func (_this *C3) Ctor__(f15 bool, f16 _dafny.Array, f9 bool) {
	{
		(_this)._f15 = f15
		(_this).F16 = f16
		(_this)._f9 = f9
	}
}
func (_this *C3) Fm3(globalState *GlobalState) bool {
	{
		return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kthtybr")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (func() _dafny.Set {
			var _coll20 = _dafny.NewBuilder()
			_ = _coll20
			for _iter23 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SetOf(true, (_this).F9()), _dafny.SetOf(false))).Elements()); ; {
				_compr_20, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _561_v0 _dafny.Set
				_561_v0 = interface{}(_compr_20).(_dafny.Set)
				if (_dafny.MultiSetOf(_dafny.SetOf(true, (_this).F9()), _dafny.SetOf(false))).Contains(_561_v0) {
					_coll20.Add(_561_v0)
				}
			}
			return _coll20.ToSet()
		}()).Cardinality()))).Cardinality()).Cmp(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-746))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg37 _dafny.Int) interface{} {
				return coer37(arg37)
			}
		}(func(_562_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		}))).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC28_(), (_this).F9())).Cardinality())) >= 0
	}
}
func (_this *C3) Fm2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(286))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}(func(_563_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-817)
		})), _dafny.SeqOf(_dafny.IntOfInt64(611))))
	}
}
func (_this *C3) Fm32(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return !(!((_this).F9()))
	}
}
func (_this *C3) Fm33(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("chkqh")
	}
}
func (_this *C3) M2(p0 _dafny.Array, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) (bool, bool, _dafny.Int, _dafny.Set) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Set = _dafny.EmptySet
		_ = r3
		var _564_v0 _dafny.Int
		_ = _564_v0
		_564_v0 = _dafny.IntOfInt64(418)
		var _565_v1 _dafny.MultiSet
		_ = _565_v1
		_565_v1 = _dafny.MultiSetOf(_564_v0)
		var _566_v2 _dafny.Map
		_ = _566_v2
		_566_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_565_v1, (_this).F9())
		var _567_v3 D7
		_ = _567_v3
		_567_v3 = Companion_D7_.Create_DC18_(_566_v2)
		_567_v3 = _567_v3
		var _568_v4 _dafny.Array
		_ = _568_v4
		var _nwElement0_15 bool = (_this).F15()
		_ = _nwElement0_15
		var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(11))
		_ = _nw86
		(_nw86).ArraySet1(_nwElement0_15, 0)
		(_nw86).ArraySet1((_this).F15(), 1)
		(_nw86).ArraySet1((_this).F15(), 2)
		(_nw86).ArraySet1((_this).F9(), 3)
		(_nw86).ArraySet1((_this).F15(), 4)
		(_nw86).ArraySet1((_this).F15(), 5)
		(_nw86).ArraySet1((_this).F9(), 6)
		(_nw86).ArraySet1((_this).F15(), 7)
		(_nw86).ArraySet1((_this).F9(), 8)
		(_nw86).ArraySet1((_this).F15(), 9)
		(_nw86).ArraySet1((_this).F9(), 10)
		_568_v4 = _nw86
		var _569_v5 _dafny.MultiSet
		_ = _569_v5
		_569_v5 = _dafny.MultiSetOf(p1, p1)
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))
		_ = _index77
		(_568_v4).ArraySet1(!(!(!(_569_v5).Contains(p1))), (_index77).Int())
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))
		_ = _index78
		(_568_v4).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("jrugtgog"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(882))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}((func(_570_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_571_i0 _dafny.Int) _dafny.CodePoint {
				return _570_p2
			}
		})(p2)))), (_index78).Int())
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))
		_ = _index79
		(_568_v4).ArraySet1((_this).F15(), (_index79).Int())
		var _572_v6 _dafny.Set
		_ = _572_v6
		_572_v6 = _dafny.SetOf(_564_v0, _564_v0)
		var _573_i1 _dafny.Int
		_ = _573_i1
		_573_i1 = _dafny.Zero
		{
			for (_572_v6).IsProperSubsetOf(_572_v6) {
				{
					if (_573_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_573_i1 = (_573_i1).Plus(_dafny.One)
					var _574_v7 _dafny.Array
					_ = _574_v7
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(15)
					_ = _len0_13
					var _nw87 _dafny.Array
					_ = _nw87
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw87 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) _dafny.Int = (func(_575_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_576_i2 _dafny.Int) _dafny.Int {
								return (_576_i2).Minus(_575_v0)
							}
						})(_564_v0)
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw87 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw87).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw87).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_574_v7 = _nw87
					var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_574_v7), 0))
					_ = _index80
					(_574_v7).ArraySet1(_564_v0, (_index80).Int())
					var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_574_v7), 0))
					_ = _index81
					(_574_v7).ArraySet1(_564_v0, (_index81).Int())
					var _577_v8 D9
					_ = _577_v8
					_577_v8 = Companion_D9_.Create_DC25_(_564_v0, (_574_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_574_v7), 0))).Int()).(_dafny.Int), _565_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_574_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_574_v7), 0))).Int()).(_dafny.Int), p2), _dafny.IntOfInt64(-777))
					var _578_v9 _dafny.Map
					_ = _578_v9
					_578_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_574_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_574_v7), 0))).Int()).(_dafny.Int), _577_v8)
					_578_v9 = _578_v9
					var _579_v10 D9
					_ = _579_v10
					_579_v10 = Companion_D9_.Create_DC24_(false, (_574_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_574_v7), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("prdj"))
					var _580_v11 _dafny.Set
					_ = _580_v11
					_580_v11 = _dafny.SetOf(p1, p1, (_579_v10).Dtor_cf43())
					var _581_v12 _dafny.Sequence
					_ = _581_v12
					_581_v12 = _dafny.SeqOf((_580_v11).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(p2, p2)).Cardinality()))
					var _582_v13 _dafny.Map
					_ = _582_v13
					_582_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _564_v0)
					var _583_v14 _dafny.Sequence
					_ = _583_v14
					_583_v14 = _dafny.SeqOf(_dafny.SeqOf((_582_v13).Cardinality()), p3, p3, p3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(775))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg40 _dafny.Int) interface{} {
							return coer40(arg40)
						}
					}((func(_584_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_585_i3 _dafny.Int) _dafny.Int {
							return _584_v0
						}
					})(_564_v0))))
					_581_v12 = _dafny.Companion_Sequence_.Concatenate((_583_v14).Select((Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((_583_v14).Cardinality()))).Uint32()).(_dafny.Sequence), p3)
					var _586_v15 _dafny.Sequence
					_ = _586_v15
					_586_v15 = _dafny.SeqOf(p1)
					var _587_v16 _dafny.Map
					_ = _587_v16
					_587_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F15()) || ((_this).F15()), _586_v15)
					var _588_v17 _dafny.Sequence
					_ = _588_v17
					_588_v17 = _dafny.SeqOf((_this).F15())
					var _589_v18 _dafny.Map
					_ = _589_v18
					_589_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_564_v0, _dafny.IntOfInt64(859))
					_587_v16 = (_587_v16).Update((_588_v17).Select((Companion_Default___.SafeIndex((_589_v18).Cardinality(), _dafny.IntOfUint32((_588_v17).Cardinality()))).Uint32()).(bool), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm34((_dafny.Zero).Minus(_564_v0), (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool), (_this).F15(), globalState), (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((Companion_Default___.Fm34((_dafny.Zero).Minus(_564_v0), (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool), (_this).F15(), globalState)).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("sooucaqvp"), (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sooucaqvp")).Cardinality()))).Uint32(), p2), (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("sooucaqvp"), (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sooucaqvp")).Cardinality()))).Uint32(), p2)).Cardinality()))).Uint32(), p2)))
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		if (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool) {
			var _590_v19 _dafny.Sequence
			_ = _590_v19
			_590_v19 = _dafny.SeqOf(Companion_Default___.Fm0(globalState))
			var _591_v20 _dafny.CodePoint
			_ = _591_v20
			_591_v20 = _dafny.CodePoint('s')
			var _592_v21 _dafny.Sequence
			_ = _592_v21
			_592_v21 = _dafny.SeqOf(true)
			var _rhs117 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_564_v0), p3)
			_ = _rhs117
			var _rhs118 _dafny.CodePoint = p2
			_ = _rhs118
			var _rhs119 _dafny.Sequence = _592_v21
			_ = _rhs119
			_590_v19 = _rhs117
			_591_v20 = _rhs118
			_592_v21 = _rhs119
			if (_this).F9() {
				var _593_v22 _dafny.Map
				_ = _593_v22
				_593_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_564_v0).Minus(_564_v0)), (_this).F15())
				_593_v22 = (_593_v22).Update(_564_v0, (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool))
				_568_v4 = _568_v4
				var _594_v23 _dafny.Array
				_ = _594_v23
				var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw88
				_594_v23 = _nw88
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_594_v23), 0))
				_ = _index82
				(_594_v23).ArraySet1(_564_v0, (_index82).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_594_v23), 0))
				_ = _index83
				(_594_v23).ArraySet1(_564_v0, (_index83).Int())
				(globalState).F2 = (_this).F9()
				r2 = _564_v0
			} else {
				var _595_v24 _dafny.Map
				_ = _595_v24
				_595_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-632))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}((func(_596_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_597_i4 _dafny.Int) _dafny.Int {
						return _596_v0
					}
				})(_564_v0)))), _dafny.IntOfInt64(308), globalState), (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool))
				r2 = (_595_v24).Cardinality()
				var _598_v25 _dafny.Array
				_ = _598_v25
				var _nwElement0_16 _dafny.Sequence = p1
				_ = _nwElement0_16
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(22))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_16, 0)
				(_nw89).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qvqafp"), 1)
				(_nw89).ArraySet1(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _591_v20), 2)
				(_nw89).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("n"), 3)
				(_nw89).ArraySet1(p1, 4)
				(_nw89).ArraySet1(p1, 5)
				(_nw89).ArraySet1(p1, 6)
				(_nw89).ArraySet1(p1, 7)
				(_nw89).ArraySet1(p1, 8)
				(_nw89).ArraySet1(p1, 9)
				(_nw89).ArraySet1(p1, 10)
				(_nw89).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_599_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_600_i5 _dafny.Int) _dafny.CodePoint {
						return _599_p2
					}
				})(p2))), 11)
				(_nw89).ArraySet1((_this).Fm33(globalState), 12)
				(_nw89).ArraySet1(p1, 13)
				(_nw89).ArraySet1((_this).Fm33(globalState), 14)
				(_nw89).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("osyhtuwn"), p1), 15)
				(_nw89).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p1, p1), (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()))).Uint32(), p2), 16)
				(_nw89).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("iqiqgyo"), 17)
				(_nw89).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("amy"), 18)
				(_nw89).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, (_this).Fm33(globalState)), 19)
				(_nw89).ArraySet1(p1, 20)
				(_nw89).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, p1), 21)
				_598_v25 = _nw89
				_598_v25 = _598_v25
				var _601_v26 _dafny.Set
				_ = _601_v26
				_601_v26 = _dafny.SetOf(false, true)
				(globalState).F2 = (_601_v26).IsSubsetOf((_601_v26).Intersection(_601_v26))
				var _602_v27 _dafny.Array
				_ = _602_v27
				var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw90
				_602_v27 = _nw90
				_602_v27 = _602_v27
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))
				_ = _index84
				(_568_v4).ArraySet1(false, (_index84).Int())
			}
			var _603_v28 _dafny.Map
			_ = _603_v28
			_603_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_564_v0, _dafny.IntOfInt64(986))
			var _604_v29 _dafny.Map
			_ = _604_v29
			_604_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), _564_v0)
			var _605_v30 _dafny.Array
			_ = _605_v30
			var _nwElement0_17 _dafny.Int = (func() _dafny.Int {
				if (_603_v28).Contains(_564_v0) {
					return (_603_v28).Get(_564_v0).(_dafny.Int)
				}
				return _564_v0
			})()
			_ = _nwElement0_17
			var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(9))
			_ = _nw91
			(_nw91).ArraySet1(_nwElement0_17, 0)
			(_nw91).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F15() {
					return _dafny.IntOfUint32((_dafny.SeqOf((_604_v29).Cardinality())).Cardinality())
				}
				return _564_v0
			})()))), 1)
			(_nw91).ArraySet1(_564_v0, 2)
			(_nw91).ArraySet1(_564_v0, 3)
			(_nw91).ArraySet1(_564_v0, 4)
			(_nw91).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(899))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_606_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_607_i6 _dafny.Int) _dafny.Sequence {
					return (Companion_D1_.Create_DC3_((_this).F9(), _606_p1, _dafny.IntOfInt64(973))).Dtor_cf7()
				}
			})(p1)))).Cardinality()), 5)
			(_nw91).ArraySet1(_564_v0, 6)
			(_nw91).ArraySet1(_564_v0, 7)
			(_nw91).ArraySet1(Companion_Default___.Fm0(globalState), 8)
			_605_v30 = _nw91
			var _nwElement0_18 _dafny.Int = Companion_Default___.SafeDivisionInt(_564_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_564_v0, (_this).F15())).Cardinality())
			_ = _nwElement0_18
			var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.One)
			_ = _nw92
			(_nw92).ArraySet1(_nwElement0_18, 0)
			_605_v30 = _nw92
			r0 = !(!((_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool)) || (!((_this).F9())))
			(globalState).F2 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(p1, (_this).Fm33(globalState)), (func() _dafny.CodePoint {
				if (_this).F9() {
					return (p1).Select((Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
				return p2
			})())
		} else {
			if !(false) {
				r1 = ((_564_v0).Times(_564_v0)).Cmp(_564_v0) <= 0
				var _608_v31 _dafny.Map
				_ = _608_v31
				_608_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_568_v4, _564_v0)
				var _609_v32 _dafny.Array
				_ = _609_v32
				var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw93
				_609_v32 = _nw93
				var _610_v33 _dafny.Array
				_ = _610_v33
				_610_v33 = _609_v32
				var _611_v34 _dafny.Set
				_ = _611_v34
				_611_v34 = _dafny.SetOf(_610_v33)
				var _612_v35 _dafny.Array
				_ = _612_v35
				var _nwElement0_19 _dafny.Int = _564_v0
				_ = _nwElement0_19
				var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(9))
				_ = _nw94
				(_nw94).ArraySet1(_nwElement0_19, 0)
				(_nw94).ArraySet1((_611_v34).Cardinality(), 1)
				(_nw94).ArraySet1(_564_v0, 2)
				(_nw94).ArraySet1(_564_v0, 3)
				(_nw94).ArraySet1(_564_v0, 4)
				(_nw94).ArraySet1(_564_v0, 5)
				(_nw94).ArraySet1(_564_v0, 6)
				(_nw94).ArraySet1(_dafny.IntOfInt64(323), 7)
				(_nw94).ArraySet1(_564_v0, 8)
				_612_v35 = _nw94
				var _613_v36 *C2
				_ = _613_v36
				var _nw95 *C2 = New_C2_()
				_ = _nw95
				_nw95.Ctor__(_564_v0, (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool))
				_613_v36 = _nw95
				var _614_v37 _dafny.Sequence
				_ = _614_v37
				_614_v37 = _dafny.SeqOf(_613_v36, _613_v36, _613_v36)
				var _615_v38 _dafny.Map
				_ = _615_v38
				_615_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_612_v35, _dafny.IntOfUint32((_614_v37).Cardinality()))
				var _616_v39 D11
				_ = _616_v39
				_616_v39 = Companion_D11_.Create_DC29_(_615_v38)
				var _rhs120 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_568_v4, _564_v0)).Merge(_608_v31)).Merge((func() _dafny.Map {
					if (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_568_v4, (_613_v36).F14())
					}
					return _608_v31
				})())
				_ = _rhs120
				var _rhs121 _dafny.Map = (_616_v39).Dtor_cf51()
				_ = _rhs121
				var _rhs122 _dafny.Int = (func() _dafny.Int {
					if ((_613_v36).F14()).Cmp((_613_v36).F14()) >= 0 {
						return (_613_v36).F14()
					}
					return (_613_v36).F14()
				})()
				_ = _rhs122
				_608_v31 = _rhs120
				_615_v38 = _rhs121
				_564_v0 = _rhs122
				var _617_v40 _dafny.Map
				_ = _617_v40
				_617_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_564_v0), _564_v0)
				var _618_v41 _dafny.Sequence
				_ = _618_v41
				_618_v41 = _dafny.SeqOf(_617_v40, _617_v40, _617_v40)
				var _619_v42 _dafny.Map
				_ = _619_v42
				_619_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_618_v41).Cardinality())), (_this).F15())
				(globalState).F2 = (func() bool {
					if ((_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool)) == ((_this).F9()) {
						return (func() bool {
							if (_619_v42).Contains(_dafny.IntOfInt64(22)) {
								return (_619_v42).Get(_dafny.IntOfInt64(22)).(bool)
							}
							return (_this).F9()
						})()
					}
					return !((_this).F15())
				})()
				var _620_v43 _dafny.Array
				_ = _620_v43
				var _nwElement0_20 _dafny.Sequence = p1
				_ = _nwElement0_20
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(13))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_20, 0)
				(_nw96).ArraySet1(p1, 1)
				(_nw96).ArraySet1(p1, 2)
				(_nw96).ArraySet1(p1, 3)
				(_nw96).ArraySet1(p1, 4)
				(_nw96).ArraySet1(p1, 5)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex((_613_v36).F14(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), p2), p1), (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex((_613_v36).F14(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), p2), p1)).Cardinality()))).Uint32(), p2), 6)
				(_nw96).ArraySet1(p1, 7)
				(_nw96).ArraySet1(p1, 8)
				(_nw96).ArraySet1(p1, 9)
				(_nw96).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ofiikthw"), 10)
				(_nw96).ArraySet1(p1, 11)
				(_nw96).ArraySet1(_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), p2), 12)
				_620_v43 = _nw96
				_620_v43 = _620_v43
				_565_v1 = _565_v1
			} else {
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))
				_ = _index85
				(_568_v4).ArraySet1((_565_v1).IsSubsetOf(_565_v1), (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))
				_ = _index86
				(_568_v4).ArraySet1((_this).F15(), (_index86).Int())
				var _621_v44 _dafny.Map
				_ = _621_v44
				_621_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool), _568_v4)
				r2 = Companion_Default___.SafeDivisionInt(((_621_v44).Merge(_621_v44)).Cardinality(), (_564_v0).Times((_dafny.Zero).Minus(_564_v0)))
				var _622_v45 _dafny.Sequence
				_ = _622_v45
				_622_v45 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(p1, p1), p1, p1, p1, p1)
				r2 = _dafny.IntOfUint32(((_622_v45).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_565_v1).Contains(_dafny.IntOfUint32((p3).Cardinality())) {
						return (_565_v1).Multiplicity(_dafny.IntOfUint32((p3).Cardinality()))
					}
					return _dafny.IntOfInt64(275)
				})(), _dafny.IntOfUint32((_622_v45).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
				var _623_v46 D1
				_ = _623_v46
				_623_v46 = Companion_D1_.Create_DC2_(_564_v0, (_this).F9(), _dafny.IntOfUint32((p1).Cardinality()), _564_v0)
				var _624_v47 _dafny.Sequence
				_ = _624_v47
				_624_v47 = _dafny.SeqOf(_623_v46, _623_v46, _623_v46, _623_v46)
				var _625_v48 D12
				_ = _625_v48
				_625_v48 = Companion_D12_.Create_DC31_(_624_v47)
				var _626_v49 *C1
				_ = _626_v49
				var _nw97 *C1 = New_C1_()
				_ = _nw97
				_nw97.Ctor__(_dafny.Companion_Sequence_.Update((_625_v48).Dtor_cf52(), (Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32(((_625_v48).Dtor_cf52()).Cardinality()))).Uint32(), _623_v46))
				_626_v49 = _nw97
			}
			var _627_v50 _dafny.CodePoint
			_ = _627_v50
			_627_v50 = _dafny.CodePoint('n')
			var _rhs123 _dafny.Int = _564_v0
			_ = _rhs123
			var _rhs124 bool = (_this).F9()
			_ = _rhs124
			var _rhs125 bool = (_this).F9()
			_ = _rhs125
			var _rhs126 _dafny.CodePoint = (p1).Select((Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs126
			_564_v0 = _rhs123
			r0 = _rhs124
			r1 = _rhs125
			_627_v50 = _rhs126
			r2 = (_564_v0).Times(_564_v0)
			Companion_Default___.M0((_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool), globalState)
			r0 = (func() bool {
				if (func() bool {
					if (_this).F15() {
						return (_this).F15()
					}
					return (_this).F9()
				})() {
					return (_this).F15()
				}
				return (_this).F15()
			})()
		}
		var _628_v51 _dafny.Sequence
		_ = _628_v51
		_628_v51 = _dafny.SeqOf((_this).F9(), (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool))
		var _629_v52 _dafny.Set
		_ = _629_v52
		_629_v52 = _dafny.SetOf((_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool), (_628_v51).Select((Companion_Default___.SafeIndex(_564_v0, _dafny.IntOfUint32((_628_v51).Cardinality()))).Uint32()).(bool), (_this).F15(), (func() bool {
			if (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool) {
				return true
			}
			return (_this).F15()
		})(), (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool))
		_629_v52 = (_629_v52).Union(_629_v52)
		r0 = (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool)
		r1 = (_568_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_568_v4), 0))).Int()).(bool)
		r2 = Companion_Default___.SafeDivisionInt(_564_v0, _dafny.IntOfInt64(186))
		r3 = (_572_v6).Difference(_572_v6)
		return r0, r1, r2, r3
	}
}
func (_this *C3) M3(p0 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _630_v0 *C2
		_ = _630_v0
		var _nw98 *C2 = New_C2_()
		_ = _nw98
		_nw98.Ctor__(_dafny.IntOfInt64(660), (_this).F9())
		_630_v0 = _nw98
		(globalState).F2 = (_this).F15()
		var _631_v1 _dafny.Set
		_ = _631_v1
		_631_v1 = _dafny.SetOf(_dafny.IntOfInt64(-456), (_630_v0).F14())
		var _632_v2 _dafny.Sequence
		_ = _632_v2
		_632_v2 = _dafny.UnicodeSeqOfUtf8Bytes("amua")
		if (_this).Fm32(_dafny.SeqOf(_631_v1), true, (_630_v0).F14(), (_dafny.IntOfUint32((_632_v2).Cardinality())).Cmp((_630_v0).F14()) <= 0, globalState) {
			var _633_v3 _dafny.Int
			_ = _633_v3
			_633_v3 = _dafny.IntOfInt64(-903)
			_633_v3 = (_630_v0).F14()
			var _634_v4 _dafny.CodePoint
			_ = _634_v4
			_634_v4 = _dafny.CodePoint('d')
			var _635_v5 _dafny.Map
			_ = _635_v5
			_635_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(369), _634_v4)
			var _636_v6 _dafny.Map
			_ = _636_v6
			_636_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (func() _dafny.CodePoint {
				if (_635_v5).Contains(_633_v3) {
					return (_635_v5).Get(_633_v3).(_dafny.CodePoint)
				}
				return Companion_Default___.Fm35(globalState)
			})())
			var _637_v7 _dafny.Sequence
			_ = _637_v7
			_637_v7 = _dafny.SeqOf(_636_v6)
			_637_v7 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(369))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_638_v6 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_639_i0 _dafny.Int) _dafny.Map {
					return _638_v6
				}
			})(_636_v6)))
			var _640_v8 _dafny.Array
			_ = _640_v8
			var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw99
			_640_v8 = _nw99
			_640_v8 = _640_v8
			var _641_v9 _dafny.Array
			_ = _641_v9
			var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
			_ = _nw100
			_641_v9 = _nw100
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_641_v9), 0))
			_ = _index87
			(_641_v9).ArraySet1(_632_v2, (_index87).Int())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_641_v9), 0))
			_ = _index88
			(_641_v9).ArraySet1(_632_v2, (_index88).Int())
			var _642_v10 _dafny.Map
			_ = _642_v10
			_642_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_630_v0).F14())
			var _643_v11 _dafny.MultiSet
			_ = _643_v11
			_643_v11 = _dafny.MultiSetOf(_dafny.IntOfInt64(992))
			var _644_v12 _dafny.Map
			_ = _644_v12
			_644_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_642_v10, (Companion_D9_.Create_DC25_(_633_v3, (_630_v0).F14(), (_643_v11).Update((_643_v11).Cardinality(), Companion_Default___.Abs(_633_v3)), _635_v5, (_630_v0).F14())).Dtor_cf45())
			_633_v3 = (((Companion_Default___.Fm36((_this).F15(), (_this).F15(), p0, globalState)).Merge(_644_v12)).Cardinality()).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_633_v3), _633_v3))
		} else {
			(globalState).F2 = ((_this).F15()) || (((_630_v0).F14()).Cmp((_630_v0).F14()) < 0)
			var _645_v13 *C2
			_ = _645_v13
			var _nw101 *C2 = New_C2_()
			_ = _nw101
			_nw101.Ctor__(((_dafny.Zero).Minus((_630_v0).F14())).Times((_630_v0).F14()), !(((_dafny.Zero).Minus((_630_v0).F14())).Cmp((_630_v0).F14()) == 0))
			_645_v13 = _nw101
			var _646_v14 _dafny.Int
			_ = _646_v14
			_646_v14 = _dafny.IntOfInt64(678)
			_646_v14 = (_630_v0).F14()
			var _647_v15 _dafny.Map
			_ = _647_v15
			_647_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this).F9())
			var _648_v16 _dafny.Map
			_ = _648_v16
			_648_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_630_v0).F14(), !((_this).F15()) || ((func() bool {
				if (_647_v15).Contains(!(true)) {
					return (_647_v15).Get(!(true)).(bool)
				}
				return (_this).F15()
			})()))
			var _649_v17 _dafny.Map
			_ = _649_v17
			_649_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_630_v0).F14(), (_648_v16).Cardinality())
			var _650_v18 _dafny.Map
			_ = _650_v18
			_650_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _649_v17)
			if (func() bool {
				if (_648_v16).Contains(_646_v14) {
					return (_648_v16).Get(_646_v14).(bool)
				}
				return !(_650_v18).Contains((_this).F9())
			})() {
				(globalState).F2 = false
				(globalState).F2 = ((_this).F9()) == (!((_this).F9()) || (p0))
				var _651_v19 _dafny.MultiSet
				_ = _651_v19
				_651_v19 = _dafny.MultiSetOf((_this).F9())
				var _652_v20 _dafny.Array
				_ = _652_v20
				var _nwElement0_21 _dafny.Int = _646_v14
				_ = _nwElement0_21
				var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(13))
				_ = _nw102
				(_nw102).ArraySet1(_nwElement0_21, 0)
				(_nw102).ArraySet1(_dafny.IntOfInt64(-248), 1)
				(_nw102).ArraySet1(_dafny.IntOfInt64(-829), 2)
				(_nw102).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_651_v19).Cardinality())).Cardinality()), 3)
				(_nw102).ArraySet1(_646_v14, 4)
				(_nw102).ArraySet1((_645_v13).F14(), 5)
				(_nw102).ArraySet1(_646_v14, 6)
				(_nw102).ArraySet1((_645_v13).F14(), 7)
				(_nw102).ArraySet1((_645_v13).F14(), 8)
				(_nw102).ArraySet1((_630_v0).F14(), 9)
				(_nw102).ArraySet1((_630_v0).F14(), 10)
				(_nw102).ArraySet1(_646_v14, 11)
				(_nw102).ArraySet1((_645_v13).F14(), 12)
				_652_v20 = _nw102
				var _653_v21 _dafny.MultiSet
				_ = _653_v21
				_653_v21 = _dafny.MultiSetOf(_652_v20, _652_v20, _652_v20)
				(_630_v0).M8((_this).F9(), (_653_v21).Union(_653_v21), globalState)
				var _654_v22 _dafny.Array
				_ = _654_v22
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
				_ = _nw103
				_654_v22 = _nw103
				var _655_v23 _dafny.Array
				_ = _655_v23
				var _nwElement0_22 _dafny.Array = _654_v22
				_ = _nwElement0_22
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(14))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_22, 0)
				(_nw104).ArraySet1(_654_v22, 1)
				(_nw104).ArraySet1((func() _dafny.Array {
					if (_this).F9() {
						return _654_v22
					}
					return _654_v22
				})(), 2)
				(_nw104).ArraySet1(_654_v22, 3)
				(_nw104).ArraySet1(_654_v22, 4)
				(_nw104).ArraySet1(_654_v22, 5)
				(_nw104).ArraySet1(_654_v22, 6)
				(_nw104).ArraySet1(_654_v22, 7)
				(_nw104).ArraySet1(_654_v22, 8)
				(_nw104).ArraySet1(_654_v22, 9)
				(_nw104).ArraySet1(_654_v22, 10)
				(_nw104).ArraySet1(_654_v22, 11)
				(_nw104).ArraySet1(_654_v22, 12)
				(_nw104).ArraySet1(_654_v22, 13)
				_655_v23 = _nw104
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_655_v23), 0))
				_ = _index89
				(_655_v23).ArraySet1(_654_v22, (_index89).Int())
				var _656_v24 D13
				_ = _656_v24
				_656_v24 = Companion_D13_.Create_DC34_(_654_v22)
				var _657_v25 _dafny.Sequence
				_ = _657_v25
				_657_v25 = _dafny.SeqOf(_654_v22, _654_v22, (_656_v24).Dtor_cf57(), _654_v22, _654_v22)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_655_v23), 0))
				_ = _index90
				(_655_v23).ArraySet1((_657_v25).Select((Companion_Default___.SafeIndex((_630_v0).F14(), _dafny.IntOfUint32((_657_v25).Cardinality()))).Uint32()).(_dafny.Array), (_index90).Int())
				_646_v14 = ((_645_v13).F14()).Minus(_646_v14)
			} else {
				var _658_v26 _dafny.CodePoint
				_ = _658_v26
				_658_v26 = _dafny.CodePoint('q')
				_658_v26 = _658_v26
				var _659_v27 _dafny.Set
				_ = _659_v27
				_659_v27 = _dafny.SetOf(_dafny.SeqOf((_645_v13).F14(), (_630_v0).F14()), _dafny.SeqOf(_646_v14))
				var _660_v28 *C0
				_ = _660_v28
				var _nw105 *C0 = New_C0_()
				_ = _nw105
				_nw105.Ctor__(_659_v27, (_645_v13).F14())
				_660_v28 = _nw105
				var _661_v29 _dafny.MultiSet
				_ = _661_v29
				_661_v29 = _dafny.MultiSetOf(_660_v28)
				var _662_v30 _dafny.Map
				_ = _662_v30
				_662_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), Companion_Default___.Fm37(p0, (_this).F9(), globalState))
				(globalState).F2 = !(_662_v30).Contains((_661_v29).IsDisjointFrom(_dafny.MultiSetOf(_660_v28)))
				var _663_v31 _dafny.Array
				_ = _663_v31
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw106
				_663_v31 = _nw106
				var _664_v32 _dafny.Sequence
				_ = _664_v32
				_664_v32 = _dafny.SeqOf(_660_v28, _660_v28, _660_v28)
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_663_v31), 0))
				_ = _index91
				(_663_v31).ArraySet1(_664_v32, (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_663_v31), 0))
				_ = _index92
				(_663_v31).ArraySet1(_664_v32, (_index92).Int())
				var _665_v33 _dafny.Array
				_ = _665_v33
				var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw107
				_665_v33 = _nw107
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_665_v33), 0))
				_ = _index93
				(_665_v33).ArraySet1((_630_v0).F14(), (_index93).Int())
				var _666_v34 _dafny.Sequence
				_ = _666_v34
				_666_v34 = _dafny.SeqOf((_this).F15(), p0)
				var _667_v35 D1
				_ = _667_v35
				_667_v35 = Companion_D1_.Create_DC3_((_this).F15(), _632_v2, (_645_v13).F14())
				var _668_v36 D4
				_ = _668_v36
				_668_v36 = Companion_D4_.Create_DC14_(!((_this).F9()), (_this).F15(), (_630_v0).F14())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_665_v33), 0))
				_ = _index94
				var _rhs127 _dafny.Int = (_dafny.IntOfUint32((_666_v34).Cardinality())).Minus((_645_v13).F14())
				_ = _rhs127
				var _rhs128 _dafny.CodePoint = Companion_Default___.Fm35(globalState)
				_ = _rhs128
				var _rhs129 _dafny.Int = Companion_Default___.SafeDivisionInt((_630_v0).F14(), _dafny.IntOfInt64(928))
				_ = _rhs129
				var _rhs130 bool = !((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("oyablh"), (_667_v35).Dtor_cf7())).Cardinality())).Cmp(((_630_v0).F14()).Plus((_668_v36).Dtor_cf28())) >= 0)
				_ = _rhs130
				var _rhs131 _dafny.CodePoint = _658_v26
				_ = _rhs131
				var _lhs76 *C0 = _660_v28
				_ = _lhs76
				var _lhs77 _dafny.Array = _665_v33
				_ = _lhs77
				var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(861), _dafny.ArrayLen((_665_v33), 0))
				_ = _lhs78
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				_lhs76.F12 = _rhs127
				_658_v26 = _rhs128
				(_lhs77).ArraySet1(_rhs129, (_lhs78).Int())
				_lhs79.F2 = _rhs130
				_658_v26 = _rhs131
				(globalState).F2 = !((_this).F9())
			}
			_646_v14 = _dafny.IntOfInt64(335)
		}
		r0 = !((p0) == ((_this).F15()))
		var _rhs132 bool = p0
		_ = _rhs132
		var _rhs133 bool = !((_this).F9())
		_ = _rhs133
		var _rhs134 bool = (_this).F9()
		_ = _rhs134
		var _rhs135 _dafny.Sequence = _632_v2
		_ = _rhs135
		var _lhs80 *GlobalState = globalState
		_ = _lhs80
		var _lhs81 *GlobalState = globalState
		_ = _lhs81
		var _lhs82 *GlobalState = globalState
		_ = _lhs82
		_lhs80.F2 = _rhs132
		_lhs81.F2 = _rhs133
		_lhs82.F2 = _rhs134
		_632_v2 = _rhs135
		var _hi2 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(globalState), (_630_v0).F14())
		_ = _hi2
		for _669_i1 := (_630_v0).F14(); _669_i1.Cmp(_hi2) < 0; _669_i1 = _669_i1.Plus(_dafny.One) {
			var _670_v37 _dafny.Map
			_ = _670_v37
			_670_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), !((_this).F9()))
			var _671_v38 _dafny.MultiSet
			_ = _671_v38
			_671_v38 = _dafny.MultiSetOf((_630_v0).F14())
			var _672_v39 _dafny.Map
			_ = _672_v39
			_672_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_630_v0).F14())
			_670_v37 = (_670_v37).Update((_671_v38).IsDisjointFrom(_671_v38), (_672_v39).Equals(_672_v39))
			var _673_v40 _dafny.Int
			_ = _673_v40
			_673_v40 = _dafny.IntOfInt64(-700)
			_673_v40 = _673_v40
			var _674_v41 _dafny.MultiSet
			_ = _674_v41
			_674_v41 = _dafny.MultiSetOf(p0, (_this).F9(), (_this).F9())
			var _675_v42 _dafny.Sequence
			_ = _675_v42
			_675_v42 = _dafny.SeqOf((_dafny.Zero).Minus((_dafny.IntOfInt64(530)).Times(_669_i1)), Companion_Default___.SafeDivisionInt((_670_v37).Cardinality(), (_630_v0).F14()), (_674_v41).Cardinality())
			_673_v40 = _dafny.IntOfUint32((_675_v42).Cardinality())
			var _676_v43 _dafny.Map
			_ = _676_v43
			_676_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(378), _669_i1), !((_this).F9()))
			var _rhs136 bool = (_this).F9()
			_ = _rhs136
			var _rhs137 _dafny.Map = ((func() _dafny.Map {
				if true {
					return _676_v43
				}
				return _676_v43
			})()).Update(_669_i1, (_673_v40).Cmp(_669_i1) > 0)
			_ = _rhs137
			var _lhs83 *GlobalState = globalState
			_ = _lhs83
			_lhs83.F2 = _rhs136
			_676_v43 = _rhs137
		}
		r0 = (_this).F15()
		return r0
	}
}
func (_this *C3) M1(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Array, globalState *GlobalState) {
	{
		var _677_v0 _dafny.Int
		_ = _677_v0
		_677_v0 = _dafny.IntOfInt64(605)
		var _678_v1 _dafny.Map
		_ = _678_v1
		_678_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _677_v0)
		var _679_v2 _dafny.Sequence
		_ = _679_v2
		_679_v2 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0), _678_v1, _678_v1, _678_v1)
		var _680_v3 D14
		_ = _680_v3
		_680_v3 = Companion_D14_.Create_DC37_(_679_v2)
		var _681_v4 _dafny.Set
		_ = _681_v4
		_681_v4 = _dafny.SetOf(_677_v0)
		var _pat_let_tv12 = _679_v2
		_ = _pat_let_tv12
		var _pat_let_tv13 = _677_v0
		_ = _pat_let_tv13
		var _pat_let_tv14 = _679_v2
		_ = _pat_let_tv14
		var _pat_let_tv15 = _678_v1
		_ = _pat_let_tv15
		var _pat_let_tv16 = _681_v4
		_ = _pat_let_tv16
		var _pat_let_tv17 = _678_v1
		_ = _pat_let_tv17
		_677_v0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_679_v2, (func(_pat_let17_0 D14) D14 {
			return func(_682_dt__update__tmp_h0 D14) D14 {
				return func(_pat_let18_0 _dafny.Sequence) D14 {
					return func(_683_dt__update_hcf66_h0 _dafny.Sequence) D14 {
						return Companion_D14_.Create_DC37_(_683_dt__update_hcf66_h0)
					}(_pat_let18_0)
				}(_dafny.SeqOf((_pat_let_tv12).Select((Companion_Default___.SafeIndex(_pat_let_tv13, _dafny.IntOfUint32((_pat_let_tv14).Cardinality()))).Uint32()).(_dafny.Map), (_pat_let_tv15).Update((_this).F9(), (_pat_let_tv16).Cardinality()), _pat_let_tv17))
			}(_pat_let17_0)
		}(_680_v3)).Dtor_cf66())).Cardinality())
		var _684_v5 _dafny.Sequence
		_ = _684_v5
		_684_v5 = _dafny.UnicodeSeqOfUtf8Bytes("thqlb")
		var _685_v6 _dafny.MultiSet
		_ = _685_v6
		_685_v6 = _dafny.MultiSetOf(_677_v0)
		var _686_v7 D1
		_ = _686_v7
		_686_v7 = Companion_D1_.Create_DC2_(_dafny.IntOfInt64(937), p1, (_685_v6).Cardinality(), _677_v0)
		var _687_v8 *C1
		_ = _687_v8
		var _nw108 *C1 = New_C1_()
		_ = _nw108
		_nw108.Ctor__(_dafny.SeqOf(Companion_Default___.Fm38((_this).F15(), _dafny.IntOfUint32((_684_v5).Cardinality()), globalState), _686_v7, _686_v7))
		_687_v8 = _nw108
		var _688_v9 _dafny.Sequence
		_ = _688_v9
		_688_v9 = _dafny.SeqOf(_687_v8)
		var _689_v10 _dafny.Sequence
		_ = _689_v10
		_689_v10 = _dafny.SeqOf(p1)
		var _690_v11 _dafny.Sequence
		_ = _690_v11
		_690_v11 = _dafny.SeqOf(p0)
		var _691_v12 _dafny.Array
		_ = _691_v12
		var _nwElement0_23 bool = p2
		_ = _nwElement0_23
		var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(15))
		_ = _nw109
		(_nw109).ArraySet1(_nwElement0_23, 0)
		(_nw109).ArraySet1(true, 1)
		(_nw109).ArraySet1((_this).F15(), 2)
		(_nw109).ArraySet1(!((_this).F15()), 3)
		(_nw109).ArraySet1((_this).F9(), 4)
		(_nw109).ArraySet1((_this).F15(), 5)
		(_nw109).ArraySet1((_this).F15(), 6)
		(_nw109).ArraySet1(p2, 7)
		(_nw109).ArraySet1((_this).F9(), 8)
		(_nw109).ArraySet1((_this).F15(), 9)
		(_nw109).ArraySet1((_this).F9(), 10)
		(_nw109).ArraySet1((_this).F9(), 11)
		(_nw109).ArraySet1((_this).F15(), 12)
		(_nw109).ArraySet1((_this).F9(), 13)
		(_nw109).ArraySet1(p1, 14)
		_691_v12 = _nw109
		var _692_v13 _dafny.MultiSet
		_ = _692_v13
		_692_v13 = _dafny.MultiSetOf(_691_v12, _691_v12)
		var _693_v14 _dafny.Map
		_ = _693_v14
		_693_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_677_v0, p1)
		var _694_v15 _dafny.Array
		_ = _694_v15
		var _nwElement0_24 _dafny.Int = (func() _dafny.Int {
			if (_this).F9() {
				return _677_v0
			}
			return _677_v0
		})()
		_ = _nwElement0_24
		var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(25))
		_ = _nw110
		(_nw110).ArraySet1(_nwElement0_24, 0)
		(_nw110).ArraySet1(_dafny.IntOfUint32((_688_v9).Cardinality()), 1)
		(_nw110).ArraySet1(_dafny.IntOfUint32((_689_v10).Cardinality()), 2)
		(_nw110).ArraySet1(_dafny.IntOfUint32((_690_v11).Cardinality()), 3)
		(_nw110).ArraySet1(_677_v0, 4)
		(_nw110).ArraySet1((p0).Plus(_677_v0), 5)
		(_nw110).ArraySet1(p0, 6)
		(_nw110).ArraySet1(p0, 7)
		(_nw110).ArraySet1((_692_v13).Cardinality(), 8)
		(_nw110).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_684_v5, _684_v5)).Cardinality()), 9)
		(_nw110).ArraySet1(_677_v0, 10)
		(_nw110).ArraySet1(_677_v0, 11)
		(_nw110).ArraySet1(_677_v0, 12)
		(_nw110).ArraySet1(p0, 13)
		(_nw110).ArraySet1(_677_v0, 14)
		(_nw110).ArraySet1(_677_v0, 15)
		(_nw110).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32(((_687_v8).Fm2(_677_v0, true, p1, p0, globalState)).Cardinality()))).Cardinality()), 16)
		(_nw110).ArraySet1((_dafny.Zero).Minus(_677_v0), 17)
		(_nw110).ArraySet1(((_dafny.Zero).Minus((_693_v14).Cardinality())).Times(_dafny.IntOfInt64(-316)), 18)
		(_nw110).ArraySet1(p0, 19)
		(_nw110).ArraySet1(_dafny.IntOfInt64(44), 20)
		(_nw110).ArraySet1(_dafny.IntOfInt64(660), 21)
		(_nw110).ArraySet1((Companion_Default___.Fm39(globalState)).Cardinality(), 22)
		(_nw110).ArraySet1(p0, 23)
		(_nw110).ArraySet1(_677_v0, 24)
		_694_v15 = _nw110
		var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
		_ = _index95
		(_694_v15).ArraySet1(p0, (_index95).Int())
		var _695_v16 _dafny.MultiSet
		_ = _695_v16
		_695_v16 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("afneutgkf"), _dafny.UnicodeSeqOfUtf8Bytes("eaxwwu"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-209))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg45 _dafny.Int) interface{} {
				return coer45(arg45)
			}
		}(func(_696_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		})))
		var _697_v17 _dafny.Array
		_ = _697_v17
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_14
		var _nw111 _dafny.Array
		_ = _nw111
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw111 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Map = (func(_698_p2 bool) func(_dafny.Int) _dafny.Map {
				return func(_699_i1 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_698_p2, true)
				}
			})(p2)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw111 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw111).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw111).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_697_v17 = _nw111
		var _700_v18 D7
		_ = _700_v18
		_700_v18 = Companion_D7_.Create_DC19_(_697_v17, p1, _684_v5)
		var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
		_ = _index96
		var _rhs138 _dafny.Int = Companion_Default___.SafeDivisionInt((_690_v11).Select((Companion_Default___.SafeIndex(_677_v0, _dafny.IntOfUint32((_690_v11).Cardinality()))).Uint32()).(_dafny.Int), p0)
		_ = _rhs138
		var _rhs139 _dafny.MultiSet = (_dafny.MultiSetOf(_684_v5, (_700_v18).Dtor_cf35(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}(func(_701_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), _684_v5)).Union(Companion_Default___.Fm40(globalState))
		_ = _rhs139
		var _rhs140 bool = (_this).F15()
		_ = _rhs140
		var _rhs141 _dafny.Int = (_677_v0).Minus((_690_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_690_v11).Cardinality()))).Uint32()).(_dafny.Int))
		_ = _rhs141
		var _lhs84 _dafny.Array = _694_v15
		_ = _lhs84
		var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
		_ = _lhs85
		var _lhs86 *GlobalState = globalState
		_ = _lhs86
		(_lhs84).ArraySet1(_rhs138, (_lhs85).Int())
		_695_v16 = _rhs139
		_lhs86.F2 = _rhs140
		_677_v0 = _rhs141
		var _702_v19 _dafny.Array
		_ = _702_v19
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_15
		var _nw112 _dafny.Array
		_ = _nw112
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw112 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) _dafny.CodePoint = func(_703_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw112 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw112).ArraySet1CodePoint(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw112).ArraySet1CodePoint(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_702_v19 = _nw112
		var _704_v20 _dafny.CodePoint
		_ = _704_v20
		_704_v20 = _dafny.CodePoint('m')
		var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_702_v19), 0))
		_ = _index97
		(_702_v19).ArraySet1CodePoint(_704_v20, (_index97).Int())
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_702_v19), 0))
		_ = _index98
		(_702_v19).ArraySet1CodePoint(_704_v20, (_index98).Int())
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_16
		var _nw113 _dafny.Array
		_ = _nw113
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw113 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) _dafny.Int = (func(_705_v15 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_706_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_706_i4, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_705_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_705_v15), 0))).Int()).(_dafny.Int), (_705_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_705_v15), 0))).Int()).(_dafny.Int))).Cardinality())
				}
			})(_694_v15)
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw113 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw113).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw113).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		_694_v15 = _nw113
		var _707_v21 D10
		_ = _707_v21
		_707_v21 = Companion_D10_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32(((_this).Fm2((_dafny.Zero).Minus(_dafny.IntOfUint32((_684_v5).Cardinality())), (_this).F9(), (_this).F9(), _677_v0, globalState)).Cardinality())))
		var _708_v22 D14
		_ = _708_v22
		_708_v22 = Companion_D14_.Create_DC38_((_702_v19).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_702_v19), 0))).Int()), (_this).F15(), Companion_Default___.Fm41(_707_v21, _677_v0, _690_v11, globalState))
		var _pat_let_tv18 = p1
		_ = _pat_let_tv18
		var _pat_let_tv19 = _689_v10
		_ = _pat_let_tv19
		var _709_v23 _dafny.Set
		_ = _709_v23
		_709_v23 = _dafny.SetOf(_708_v22, _708_v22, _708_v22, _708_v22, func(_pat_let19_0 D14) D14 {
			return func(_710_dt__update__tmp_h1 D14) D14 {
				return func(_pat_let20_0 bool) D14 {
					return func(_711_dt__update_hcf68_h0 bool) D14 {
						return func(_pat_let21_0 _dafny.Sequence) D14 {
							return func(_712_dt__update_hcf69_h0 _dafny.Sequence) D14 {
								return Companion_D14_.Create_DC38_((_710_dt__update__tmp_h1).Dtor_cf67(), _711_dt__update_hcf68_h0, _712_dt__update_hcf69_h0)
							}(_pat_let21_0)
						}(_pat_let_tv19)
					}(_pat_let20_0)
				}(_pat_let_tv18)
			}(_pat_let19_0)
		}(_708_v22))
		_709_v23 = (_709_v23).Intersection((_709_v23).Difference(_709_v23))
		var _713_v24 _dafny.Sequence
		_ = _713_v24
		_713_v24 = _dafny.SeqOf(_685_v6, _dafny.MultiSetFromSeq(_690_v11), _685_v6)
		var _714_v25 _dafny.MultiSet
		_ = _714_v25
		_714_v25 = (_713_v24).Select((Companion_Default___.SafeIndex((_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_713_v24).Cardinality()))).Uint32()).(_dafny.MultiSet)
		if func(_source7 _dafny.MultiSet) bool {
			var _715___mcc_h0 _dafny.MultiSet = _source7
			_ = _715___mcc_h0
			var _716_cf31 _dafny.MultiSet = _715___mcc_h0
			_ = _716_cf31
			return false
		}(_714_v25) {
			(globalState).F2 = ((_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int)).Cmp(_677_v0) >= 0
			if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-978))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_717_v22 D14) func(_dafny.Int) _dafny.CodePoint {
				return func(_718_i5 _dafny.Int) _dafny.CodePoint {
					return (_717_v22).Dtor_cf67()
				}
			})(_708_v22))), _684_v5) {
				var _719_v26 _dafny.Map
				_ = _719_v26
				_719_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), p1)
				var _720_v27 _dafny.Set
				_ = _720_v27
				_720_v27 = _dafny.SetOf(!(p1))
				var _721_v28 D15
				_ = _721_v28
				_721_v28 = Companion_D15_.Create_DC40_(_dafny.SetOf((_this).F9()))
				var _722_v29 _dafny.Sequence
				_ = _722_v29
				_722_v29 = _dafny.SeqOf(_720_v27, _720_v27, (_721_v28).Dtor_cf72(), _720_v27)
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _index99
				(_694_v15).ArraySet1(Companion_Default___.SafeDivisionInt((_719_v26).Cardinality(), ((_722_v29).Select((Companion_Default___.SafeIndex(_677_v0, _dafny.IntOfUint32((_722_v29).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()), (_index99).Int())
				var _723_v30 _dafny.Map
				_ = _723_v30
				_723_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _704_v20)
				(globalState).F2 = !((p2) || (_dafny.Companion_Sequence_.Contains((_this).Fm33(globalState), (func() _dafny.CodePoint {
					if (_723_v30).Contains(p1) {
						return (_723_v30).Get(p1).(_dafny.CodePoint)
					}
					return (_702_v19).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_702_v19), 0))).Int())
				})())))
				_677_v0 = (_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int)
				var _724_v31 _dafny.Map
				_ = _724_v31
				_724_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_684_v5, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(511))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_725_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_726_i6 _dafny.Int) _dafny.Int {
						return _725_p0
					}
				})(p0)))).Cardinality()))
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _index100
				(_694_v15).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_724_v31).Contains(_dafny.Companion_Sequence_.Concatenate(_684_v5, (_this).Fm33(globalState))) {
						return (_724_v31).Get(_dafny.Companion_Sequence_.Concatenate(_684_v5, (_this).Fm33(globalState))).(_dafny.Int)
					}
					return (Companion_Default___.Fm42(p2, p1, _dafny.IntOfUint32((_689_v10).Cardinality()), globalState)).Cardinality()
				})()), (_index100).Int())
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw114
				_694_v15 = _nw114
			} else {
				var _727_v32 _dafny.Map
				_ = _727_v32
				_727_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), p1)
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _index101
				(_694_v15).ArraySet1((_727_v32).Cardinality(), (_index101).Int())
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _index102
				(_694_v15).ArraySet1(p0, (_index102).Int())
				var _728_v33 _dafny.Set
				_ = _728_v33
				_728_v33 = _dafny.SetOf((_677_v0).Cmp((_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int)) > 0)
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _index103
				var _rhs142 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(p1, p1)).Cardinality()), _677_v0)
				_ = _rhs142
				var _rhs143 _dafny.Set = _728_v33
				_ = _rhs143
				var _rhs144 _dafny.Array = _694_v15
				_ = _rhs144
				var _rhs145 bool = (_this).F9()
				_ = _rhs145
				var _rhs146 bool = (p0).Cmp(_677_v0) <= 0
				_ = _rhs146
				var _lhs87 _dafny.Array = _694_v15
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _lhs88
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				(_lhs87).ArraySet1(_rhs142, (_lhs88).Int())
				_728_v33 = _rhs143
				_694_v15 = _rhs144
				_lhs89.F2 = _rhs145
				_lhs90.F2 = _rhs146
				var _729_v34 _dafny.Array
				_ = _729_v34
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(29))
				_ = _nw115
				_729_v34 = _nw115
				var _730_v35 D3
				_ = _730_v35
				_730_v35 = Companion_D3_.Create_DC8_()
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_729_v34), 0))
				_ = _index104
				(_729_v34).ArraySet1(_730_v35, (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _index105
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_729_v34), 0))
				_ = _index106
				var _rhs147 bool = (Companion_Default___.SafeModuloInt((_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int), p0)).Cmp(p0) >= 0
				_ = _rhs147
				var _rhs148 _dafny.Map = _727_v32
				_ = _rhs148
				var _rhs149 _dafny.Int = (func() _dafny.Int {
					if p1 {
						return p0
					}
					return (_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int)
				})()
				_ = _rhs149
				var _rhs150 D3 = (func() D3 {
					if p1 {
						return _730_v35
					}
					return _730_v35
				})()
				_ = _rhs150
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				var _lhs92 _dafny.Array = _694_v15
				_ = _lhs92
				var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _lhs93
				var _lhs94 _dafny.Array = _729_v34
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_729_v34), 0))
				_ = _lhs95
				_lhs91.F2 = _rhs147
				_727_v32 = _rhs148
				(_lhs92).ArraySet1(_rhs149, (_lhs93).Int())
				(_lhs94).ArraySet1(_rhs150, (_lhs95).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
				_ = _index107
				(_694_v15).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm0(globalState)), (_index107).Int())
			}
			var _731_v36 D4
			_ = _731_v36
			_731_v36 = Companion_D4_.Create_DC13_(_677_v0)
			_677_v0 = (_dafny.Zero).Minus((_731_v36).Dtor_cf25())
			var _732_v37 _dafny.Array
			_ = _732_v37
			var _nwElement0_25 _dafny.Sequence = _684_v5
			_ = _nwElement0_25
			var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(8))
			_ = _nw116
			(_nw116).ArraySet1(_nwElement0_25, 0)
			(_nw116).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rmco"), 1)
			(_nw116).ArraySet1((_this).Fm33(globalState), 2)
			(_nw116).ArraySet1(_684_v5, 3)
			(_nw116).ArraySet1(_684_v5, 4)
			(_nw116).ArraySet1(_684_v5, 5)
			(_nw116).ArraySet1(_684_v5, 6)
			(_nw116).ArraySet1(_684_v5, 7)
			_732_v37 = _nw116
			var _733_v38 _dafny.Sequence
			_ = _733_v38
			_733_v38 = _dafny.SeqOf(_732_v37, _732_v37, _732_v37)
			var _734_v39 _dafny.Map
			_ = _734_v39
			_734_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_687_v8, _dafny.IntOfInt64(590)), (_733_v38).Select((Companion_Default___.SafeIndex((_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_733_v38).Cardinality()))).Uint32()).(_dafny.Array))
			var _735_v40 _dafny.Map
			_ = _735_v40
			_735_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_687_v8, p0)
			_734_v39 = (_734_v39).Update(_735_v40, _732_v37)
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
			_ = _index108
			(_694_v15).ArraySet1(_677_v0, (_index108).Int())
		} else {
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))
			_ = _index109
			(_694_v15).ArraySet1(p0, (_index109).Int())
			var _736_v41 _dafny.Map
			_ = _736_v41
			_736_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int), (_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int))
			var _737_v42 _dafny.Map
			_ = _737_v42
			_737_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((Companion_D16_.Create_DC45_(_736_v41)).Dtor_cf78()).Merge(_736_v41), ((_this).F9()) || ((_this).F15()))
			(globalState).F2 = (func() bool {
				if (_737_v42).Contains(_736_v41) {
					return (_737_v42).Get(_736_v41).(bool)
				}
				return p1
			})()
			_677_v0 = (_694_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_694_v15), 0))).Int()).(_dafny.Int)
			_677_v0 = Companion_Default___.Fm0(globalState)
			_690_v11 = _690_v11
		}
	}
}
func (_this *C3) F15() bool {
	{
		return _this._f15
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f9 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f9 = false
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

func (_this *C4) F9() bool {
	return _this._f9
}
func (_this *C4) Ctor__(f9 bool) {
	{
		(_this)._f9 = f9
	}
}
func (_this *C4) Fm2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(222))).Cardinality(), (_dafny.MultiSetOf(!((_this).F9()))).Cardinality()), _dafny.SeqOf((_dafny.MultiSetOf(_dafny.CodePoint('y'), _dafny.CodePoint('v'))).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(413))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}(func(_738_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(725)
		})), _dafny.SeqOf(_dafny.IntOfInt64(381))))
	}
}
func (_this *C4) M1(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Array, globalState *GlobalState) {
	{
		var _739_v0 _dafny.MultiSet
		_ = _739_v0
		_739_v0 = _dafny.MultiSetOf(p0)
		var _740_i0 _dafny.Int
		_ = _740_i0
		_740_i0 = _dafny.Zero
		{
			for (((_739_v0).Intersection(_739_v0)).Cardinality()).Cmp(p0) <= 0 {
				{
					if (_740_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_740_i0 = (_740_i0).Plus(_dafny.One)
					var _741_v1 _dafny.Array
					_ = _741_v1
					var _len0_17 _dafny.Int = _dafny.IntOfInt64(14)
					_ = _len0_17
					var _nw117 _dafny.Array
					_ = _nw117
					if _len0_17.Cmp(_dafny.Zero) == 0 {
						_nw117 = _dafny.NewArray(_len0_17)
					} else {
						var _init17 func(_dafny.Int) _dafny.CodePoint = func(_742_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('x')
						}
						_ = _init17
						var _element0_17 = _init17(_dafny.Zero)
						_ = _element0_17
						_nw117 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
						(_nw117).ArraySet1CodePoint(_element0_17, 0)
						var _nativeLen0_17 = (_len0_17).Int()
						_ = _nativeLen0_17
						for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
							(_nw117).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
						}
					}
					_741_v1 = _nw117
					_741_v1 = _741_v1
					var _743_v2 _dafny.Array
					_ = _743_v2
					var _len0_18 _dafny.Int = _dafny.IntOfInt64(26)
					_ = _len0_18
					var _nw118 _dafny.Array
					_ = _nw118
					if _len0_18.Cmp(_dafny.Zero) == 0 {
						_nw118 = _dafny.NewArray(_len0_18)
					} else {
						var _init18 func(_dafny.Int) bool = (func(_744_p2 bool) func(_dafny.Int) bool {
							return func(_745_i2 _dafny.Int) bool {
								return _744_p2
							}
						})(p2)
						_ = _init18
						var _element0_18 = _init18(_dafny.Zero)
						_ = _element0_18
						_nw118 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
						(_nw118).ArraySet1(_element0_18, 0)
						var _nativeLen0_18 = (_len0_18).Int()
						_ = _nativeLen0_18
						for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
							(_nw118).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
						}
					}
					_743_v2 = _nw118
					var _746_v3 _dafny.Map
					_ = _746_v3
					_746_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
					var _747_v4 _dafny.Sequence
					_ = _747_v4
					_747_v4 = _dafny.SeqOf((_746_v3).Cardinality())
					var _748_v5 _dafny.Sequence
					_ = _748_v5
					_748_v5 = _dafny.SeqOf((p0).Plus(p0), p0, (p0).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _743_v2)).Cardinality()), (_747_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_747_v4).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F9())).Cardinality())
					_747_v4 = _dafny.Companion_Sequence_.Concatenate(_747_v4, _748_v5)
					var _749_v6 _dafny.Int
					_ = _749_v6
					_749_v6 = _dafny.IntOfInt64(442)
					_749_v6 = (_dafny.IntOfUint32((_747_v4).Cardinality())).Times(p0)
					var _750_v7 _dafny.Array
					_ = _750_v7
					var _len0_19 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_19
					var _nw119 _dafny.Array
					_ = _nw119
					if _len0_19.Cmp(_dafny.Zero) == 0 {
						_nw119 = _dafny.NewArray(_len0_19)
					} else {
						var _init19 func(_dafny.Int) _dafny.Sequence = func(_751_i3 _dafny.Int) _dafny.Sequence {
							return _dafny.UnicodeSeqOfUtf8Bytes("kwxxvgrr")
						}
						_ = _init19
						var _element0_19 = _init19(_dafny.Zero)
						_ = _element0_19
						_nw119 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
						(_nw119).ArraySet1(_element0_19, 0)
						var _nativeLen0_19 = (_len0_19).Int()
						_ = _nativeLen0_19
						for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
							(_nw119).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
						}
					}
					_750_v7 = _nw119
					var _752_v8 D1
					_ = _752_v8
					_752_v8 = Companion_D1_.Create_DC3_(p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(388))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg50 _dafny.Int) interface{} {
							return coer50(arg50)
						}
					}(func(_753_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('e')
					})), _749_v6)
					var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_750_v7), 0))
					_ = _index110
					(_750_v7).ArraySet1((_752_v8).Dtor_cf7(), (_index110).Int())
					var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_750_v7), 0))
					_ = _index111
					(_750_v7).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(997))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg51 _dafny.Int) interface{} {
							return coer51(arg51)
						}
					}(func(_754_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('d')
					})), (_index111).Int())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		(globalState).F2 = !(!(true))
		var _755_v9 _dafny.Sequence
		_ = _755_v9
		_755_v9 = _dafny.SeqOf(p0)
		var _756_v10 _dafny.MultiSet
		_ = _756_v10
		_756_v10 = _dafny.MultiSetOf(_755_v9)
		var _757_v11 D1
		_ = _757_v11
		_757_v11 = Companion_D1_.Create_DC2_(p0, (_this).F9(), p0, (_756_v10).Cardinality())
		var _758_v12 _dafny.Sequence
		_ = _758_v12
		_758_v12 = _dafny.SeqOf(_757_v11)
		var _759_v13 *C1
		_ = _759_v13
		var _nw120 *C1 = New_C1_()
		_ = _nw120
		_nw120.Ctor__(_dafny.Companion_Sequence_.Concatenate(_758_v12, _758_v12))
		_759_v13 = _nw120
		var _760_v14 _dafny.Map
		_ = _760_v14
		_760_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
		var _761_v15 _dafny.CodePoint
		_ = _761_v15
		_761_v15 = _dafny.CodePoint('o')
		var _762_v16 _dafny.Sequence
		_ = _762_v16
		_762_v16 = _dafny.SeqOf(p2, false)
		var _763_v17 _dafny.Map
		_ = _763_v17
		_763_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _762_v16)
		var _764_v18 _dafny.Array
		_ = _764_v18
		var _nwElement0_26 bool = p2
		_ = _nwElement0_26
		var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(21))
		_ = _nw121
		(_nw121).ArraySet1(_nwElement0_26, 0)
		(_nw121).ArraySet1((_this).F9(), 1)
		(_nw121).ArraySet1(p1, 2)
		(_nw121).ArraySet1((func() bool {
			if p2 {
				return p2
			}
			return p2
		})(), 3)
		(_nw121).ArraySet1(p2, 4)
		(_nw121).ArraySet1(p1, 5)
		(_nw121).ArraySet1(!((_this).F9()), 6)
		(_nw121).ArraySet1((p0).Cmp(_dafny.IntOfInt64(868)) != 0, 7)
		(_nw121).ArraySet1(!((_739_v0).Update(p0, Companion_Default___.Abs((_dafny.Zero).Minus((_760_v14).Cardinality())))).Contains(p0), 8)
		(_nw121).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_761_v15, _dafny.MultiSetFromSeq((func() _dafny.Sequence {
			if (_763_v17).Contains(_761_v15) {
				return (_763_v17).Get(_761_v15).(_dafny.Sequence)
			}
			return _762_v16
		})()))).Contains(_761_v15), 9)
		(_nw121).ArraySet1((p1) || (p1), 10)
		(_nw121).ArraySet1(false, 11)
		(_nw121).ArraySet1((p0).Cmp(p0) >= 0, 12)
		(_nw121).ArraySet1((true) == (!(p1)), 13)
		(_nw121).ArraySet1(Companion_Default___.Fm1(_739_v0, p0, globalState), 14)
		(_nw121).ArraySet1((_this).F9(), 15)
		(_nw121).ArraySet1((_this).F9(), 16)
		(_nw121).ArraySet1((_this).F9(), 17)
		(_nw121).ArraySet1(p2, 18)
		(_nw121).ArraySet1((_762_v16).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_762_v16).Cardinality()))).Uint32()).(bool), 19)
		(_nw121).ArraySet1(p1, 20)
		_764_v18 = _nw121
		for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_764_v18), 0))); ; {
			_guard_loop_3, _ok24 := _iter24()
			if !_ok24 {
				break
			}
			var _765_i6 _dafny.Int
			_765_i6 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_765_i6).Sign() != -1) && ((_765_i6).Cmp(_dafny.ArrayLen((_764_v18), 0)) < 0)) {
				(_764_v18).ArraySet1(false, (_765_i6).Int())
			}
		}
		var _pat_let_tv20 = p0
		_ = _pat_let_tv20
		var _pat_let_tv21 = p0
		_ = _pat_let_tv21
		var _pat_let_tv22 = p0
		_ = _pat_let_tv22
		var _pat_let_tv23 = p0
		_ = _pat_let_tv23
		var _pat_let_tv24 = p0
		_ = _pat_let_tv24
		var _source8 D1 = func(_source9 D3) D1 {
			if _source9.Is_DC8() {
				return Companion_D1_.Create_DC1_(_pat_let_tv20)
			} else if _source9.Is_DC9() {
				var _766___mcc_h14 bool = _source9.Get_().(D3_DC9).Cf17
				_ = _766___mcc_h14
				var _767___mcc_h15 _dafny.Sequence = _source9.Get_().(D3_DC9).Cf18
				_ = _767___mcc_h15
				var _768_cf18 _dafny.Sequence = _767___mcc_h15
				_ = _768_cf18
				var _769_cf17 bool = _766___mcc_h14
				_ = _769_cf17
				if _769_cf17 {
					return Companion_D1_.Create_DC1_((_dafny.Zero).Minus(_pat_let_tv21))
				} else {
					return Companion_D1_.Create_DC1_(_pat_let_tv22)
				}
			} else if _source9.Is_DC10() {
				return Companion_D1_.Create_DC1_(_pat_let_tv23)
			} else {
				var _770___mcc_h16 _dafny.MultiSet = _source9.Get_().(D3_DC7).Cf16
				_ = _770___mcc_h16
				var _771_cf16 _dafny.MultiSet = _770___mcc_h16
				_ = _771_cf16
				return Companion_D1_.Create_DC1_(_pat_let_tv24)
			}
		}(Companion_D3_.Create_DC8_())
		_ = _source8
		if _source8.Is_DC2() {
			var _772___mcc_h0 _dafny.Int = _source8.Get_().(D1_DC2).Cf2
			_ = _772___mcc_h0
			var _773___mcc_h1 bool = _source8.Get_().(D1_DC2).Cf3
			_ = _773___mcc_h1
			var _774___mcc_h2 _dafny.Int = _source8.Get_().(D1_DC2).Cf4
			_ = _774___mcc_h2
			var _775___mcc_h3 _dafny.Int = _source8.Get_().(D1_DC2).Cf5
			_ = _775___mcc_h3
			var _776_cf5 _dafny.Int = _775___mcc_h3
			_ = _776_cf5
			var _777_cf4 _dafny.Int = _774___mcc_h2
			_ = _777_cf4
			var _778_cf3 bool = _773___mcc_h1
			_ = _778_cf3
			var _779_cf2 _dafny.Int = _772___mcc_h0
			_ = _779_cf2
			var _780_v19 _dafny.Map
			_ = _780_v19
			_780_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			var _781_v20 D1
			_ = _781_v20
			_781_v20 = Companion_D1_.Create_DC4_(_776_cf5, p2, (func() bool {
				if (_780_v19).Contains(_777_cf4) {
					return (_780_v19).Get(_777_cf4).(bool)
				}
				return p2
			})(), true, _778_cf3)
			var _782_v21 D1
			_ = _782_v21
			_782_v21 = Companion_D1_.Create_DC5_(_781_v20)
			var _source10 D1 = _782_v21
			_ = _source10
			if _source10.Is_DC2() {
				var _783___mcc_h17 _dafny.Int = _source10.Get_().(D1_DC2).Cf2
				_ = _783___mcc_h17
				var _784___mcc_h18 bool = _source10.Get_().(D1_DC2).Cf3
				_ = _784___mcc_h18
				var _785___mcc_h19 _dafny.Int = _source10.Get_().(D1_DC2).Cf4
				_ = _785___mcc_h19
				var _786___mcc_h20 _dafny.Int = _source10.Get_().(D1_DC2).Cf5
				_ = _786___mcc_h20
				var _787_cf5 _dafny.Int = _786___mcc_h20
				_ = _787_cf5
				var _788_cf4 _dafny.Int = _785___mcc_h19
				_ = _788_cf4
				var _789_cf3 bool = _784___mcc_h18
				_ = _789_cf3
				var _790_cf2 _dafny.Int = _783___mcc_h17
				_ = _790_cf2
				var _791_v22 _dafny.Array
				_ = _791_v22
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_20
				var _nw122 _dafny.Array
				_ = _nw122
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw122 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Int = (func(_792_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_793_i7 _dafny.Int) _dafny.Int {
							return (_793_i7).Minus((_dafny.Zero).Minus(_792_cf4))
						}
					})(_788_cf4)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw122 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw122).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw122).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_791_v22 = _nw122
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_791_v22), 0))
				_ = _index112
				(_791_v22).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_787_cf5)).Cardinality()), (_index112).Int())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_791_v22), 0))
				_ = _index113
				(_791_v22).ArraySet1(_790_cf2, (_index113).Int())
				var _794_v23 T0
				_ = _794_v23
				var _nw123 *C1 = New_C1_()
				_ = _nw123
				_nw123.Ctor__(_759_v13.F13)
				_794_v23 = _nw123
				var _795_v24 bool
				_ = _795_v24
				var _796_v25 bool
				_ = _796_v25
				var _out10 bool
				_ = _out10
				var _out11 bool
				_ = _out11
				_out10, _out11 = (_759_v13).M6(_794_v23, _764_v18, true, globalState)
				_795_v24 = _out10
				_796_v25 = _out11
				var _797_v26 _dafny.Array
				_ = _797_v26
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
				_ = _nw124
				_797_v26 = _nw124
				var _798_v27 _dafny.Array
				_ = _798_v27
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_21
				var _nw125 _dafny.Array
				_ = _nw125
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw125 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Map = (func(_799_v0 _dafny.MultiSet) func(_dafny.Int) _dafny.Map {
						return func(_800_i8 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_799_v0, true)
						}
					})(_739_v0)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw125 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw125).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw125).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_798_v27 = _nw125
				var _801_v28 _dafny.Map
				_ = _801_v28
				_801_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_739_v0, _796_v25)
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_798_v27), 0))
				_ = _index114
				(_798_v27).ArraySet1(_801_v28, (_index114).Int())
				var _802_v29 _dafny.Sequence
				_ = _802_v29
				_802_v29 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(499))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_803_v24 bool, _804_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_805_i9 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_803_v24, _804_p0)).Cardinality()
					}
				})(_795_v24, p0))))
				var _806_v31 D4
				_ = _806_v31
				_806_v31 = Companion_D4_.Create_DC12_(_795_v24, _796_v25, p0, (_791_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_791_v22), 0))).Int()).(_dafny.Int), (_this).F9())
				var _807_v32 D7
				_ = _807_v32
				_807_v32 = Companion_D7_.Create_DC18_(func() _dafny.Map {
					var _coll21 = _dafny.NewMapBuilder()
					_ = _coll21
					for _iter25 := _dafny.Iterate((_dafny.SeqOf(_739_v0, _739_v0, _dafny.MultiSetOf((_dafny.Zero).Minus((_806_v31).Dtor_cf22()), _788_cf4))).Elements()); ; {
						_compr_21, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _808_v30 _dafny.MultiSet
						_808_v30 = interface{}(_compr_21).(_dafny.MultiSet)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_739_v0, _739_v0, _dafny.MultiSetOf((_dafny.Zero).Minus((_806_v31).Dtor_cf22()), _788_cf4)), _808_v30) {
							_coll21.Add(_808_v30, true)
						}
					}
					return _coll21.ToMap()
				}())
				var _pat_let_tv25 = _781_v20
				_ = _pat_let_tv25
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_798_v27), 0))
				_ = _index115
				var _rhs151 D1 = func(_pat_let22_0 D1) D1 {
					return func(_809_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let23_0 D1) D1 {
							return func(_810_dt__update_hcf14_h0 D1) D1 {
								return Companion_D1_.Create_DC5_(_810_dt__update_hcf14_h0)
							}(_pat_let23_0)
						}(_pat_let_tv25)
					}(_pat_let22_0)
				}(Companion_D1_.Create_DC5_(_781_v20))
				_ = _rhs151
				var _rhs152 _dafny.Array = _797_v26
				_ = _rhs152
				var _rhs153 bool = (func() bool {
					if !((_this).F9()) {
						return !_dafny.Companion_Sequence_.Equal(_802_v29, _802_v29)
					}
					return !(_795_v24)
				})()
				_ = _rhs153
				var _rhs154 _dafny.Map = (_807_v32).Dtor_cf32()
				_ = _rhs154
				var _lhs96 _dafny.Array = _798_v27
				_ = _lhs96
				var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_798_v27), 0))
				_ = _lhs97
				_782_v21 = _rhs151
				_797_v26 = _rhs152
				_789_cf3 = _rhs153
				(_lhs96).ArraySet1(_rhs154, (_lhs97).Int())
				_787_cf5 = (_790_cf2).Times((_739_v0).Cardinality())
			} else if _source10.Is_DC3() {
				var _811___mcc_h21 bool = _source10.Get_().(D1_DC3).Cf6
				_ = _811___mcc_h21
				var _812___mcc_h22 _dafny.Sequence = _source10.Get_().(D1_DC3).Cf7
				_ = _812___mcc_h22
				var _813___mcc_h23 _dafny.Int = _source10.Get_().(D1_DC3).Cf8
				_ = _813___mcc_h23
				var _814_cf8 _dafny.Int = _813___mcc_h23
				_ = _814_cf8
				var _815_cf7 _dafny.Sequence = _812___mcc_h22
				_ = _815_cf7
				var _816_cf6 bool = _811___mcc_h21
				_ = _816_cf6
				Companion_Default___.M0(p1, globalState)
				_761_v15 = _761_v15
				var _817_v33 _dafny.Set
				_ = _817_v33
				_817_v33 = _dafny.SetOf(_755_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(881))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}(func(_818_i10 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(561)
				})))
				var _819_v34 *C0
				_ = _819_v34
				var _nw126 *C0 = New_C0_()
				_ = _nw126
				_nw126.Ctor__(_817_v33, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _778_cf3)).Cardinality())
				_819_v34 = _nw126
				var _820_v35 _dafny.Map
				_ = _820_v35
				_820_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_816_cf6, _819_v34)
				var _821_v36 _dafny.Sequence
				_ = _821_v36
				_821_v36 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _819_v34), (_820_v35).Merge(_820_v35), (_820_v35).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_816_cf6, _819_v34)))
				var _822_v37 _dafny.MultiSet
				_ = _822_v37
				_822_v37 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(720))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_823_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_824_i11 _dafny.Int) _dafny.CodePoint {
						return _823_v15
					}
				})(_761_v15))), _dafny.UnicodeSeqOfUtf8Bytes("unmdd"))
				var _825_v38 _dafny.Sequence
				_ = _825_v38
				_825_v38 = _dafny.SeqOf(_822_v37)
				var _826_v39 _dafny.Sequence
				_ = _826_v39
				_826_v39 = _dafny.SeqOf(_815_cf7, _815_cf7)
				var _827_v40 _dafny.Array
				_ = _827_v40
				var _nwElement0_27 _dafny.MultiSet = (_822_v37).Union(_822_v37)
				_ = _nwElement0_27
				var _nw127 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(21))
				_ = _nw127
				(_nw127).ArraySet1(_nwElement0_27, 0)
				(_nw127).ArraySet1((_822_v37).Union((_825_v38).Select((Companion_Default___.SafeIndex(_779_cf2, _dafny.IntOfUint32((_825_v38).Cardinality()))).Uint32()).(_dafny.MultiSet)), 1)
				(_nw127).ArraySet1(_822_v37, 2)
				(_nw127).ArraySet1(_822_v37, 3)
				(_nw127).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(954))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_828_cf7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_829_i12 _dafny.Int) _dafny.Sequence {
						return _828_cf7
					}
				})(_815_cf7))))).Union(_822_v37), 4)
				(_nw127).ArraySet1((_dafny.MultiSetFromSeq(_826_v39)).Difference(_822_v37), 5)
				(_nw127).ArraySet1(Companion_Default___.Fm19(p0, globalState), 6)
				(_nw127).ArraySet1(_822_v37, 7)
				(_nw127).ArraySet1((_dafny.MultiSetFromSeq(_826_v39)).Intersection(_822_v37), 8)
				(_nw127).ArraySet1(_822_v37, 9)
				(_nw127).ArraySet1(_822_v37, 10)
				(_nw127).ArraySet1(_822_v37, 11)
				(_nw127).ArraySet1(_822_v37, 12)
				(_nw127).ArraySet1(Companion_Default___.Fm19(_777_cf4, globalState), 13)
				(_nw127).ArraySet1(_822_v37, 14)
				(_nw127).ArraySet1(_822_v37, 15)
				(_nw127).ArraySet1((_822_v37).Update(_815_cf7, Companion_Default___.Abs(_819_v34.F12)), 16)
				(_nw127).ArraySet1(_822_v37, 17)
				(_nw127).ArraySet1(_822_v37, 18)
				(_nw127).ArraySet1(_822_v37, 19)
				(_nw127).ArraySet1(_822_v37, 20)
				_827_v40 = _nw127
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_827_v40), 0))
				_ = _index116
				(_827_v40).ArraySet1(_822_v37, (_index116).Int())
				var _830_v41 _dafny.Map
				_ = _830_v41
				_830_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_819_v34.F12, _819_v34)
				var _831_v42 _dafny.Array
				_ = _831_v42
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_22
				var _nw128 _dafny.Array
				_ = _nw128
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw128 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Map = (func(_832_p1 bool, _833_cf3 bool) func(_dafny.Int) _dafny.Map {
						return func(_834_i13 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_832_p1, _833_cf3)
						}
					})(p1, _778_cf3)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw128 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw128).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw128).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_831_v42 = _nw128
				var _835_v43 D7
				_ = _835_v43
				_835_v43 = Companion_D7_.Create_DC19_(_831_v42, p1, _815_cf7)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_827_v40), 0))
				_ = _index117
				var _rhs155 _dafny.Int = _dafny.IntOfInt64(-627)
				_ = _rhs155
				var _rhs156 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_821_v36, _821_v36), (Companion_Default___.SafeIndex(_776_cf5, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_821_v36, _821_v36)).Cardinality()))).Uint32(), _820_v35)
				_ = _rhs156
				var _rhs157 _dafny.MultiSet = _822_v37
				_ = _rhs157
				var _rhs158 _dafny.Map = _830_v41
				_ = _rhs158
				var _rhs159 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_835_v43).Dtor_cf35(), _815_cf7)).Cardinality())
				_ = _rhs159
				var _lhs98 _dafny.Array = _827_v40
				_ = _lhs98
				var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_827_v40), 0))
				_ = _lhs99
				_777_cf4 = _rhs155
				_821_v36 = _rhs156
				(_lhs98).ArraySet1(_rhs157, (_lhs99).Int())
				_830_v41 = _rhs158
				_777_cf4 = _rhs159
				var _836_v44 D1
				_ = _836_v44
				_836_v44 = Companion_D1_.Create_DC4_(_777_cf4, true, !(!(_778_cf3)), (_this).F9(), false)
				var _837_v45 *C0
				_ = _837_v45
				var _nw129 *C0 = New_C0_()
				_ = _nw129
				_nw129.Ctor__(_817_v33, (_836_v44).Dtor_cf9())
				_837_v45 = _nw129
			} else if _source10.Is_DC4() {
				var _838___mcc_h24 _dafny.Int = _source10.Get_().(D1_DC4).Cf9
				_ = _838___mcc_h24
				var _839___mcc_h25 bool = _source10.Get_().(D1_DC4).Cf10
				_ = _839___mcc_h25
				var _840___mcc_h26 bool = _source10.Get_().(D1_DC4).Cf11
				_ = _840___mcc_h26
				var _841___mcc_h27 bool = _source10.Get_().(D1_DC4).Cf12
				_ = _841___mcc_h27
				var _842___mcc_h28 bool = _source10.Get_().(D1_DC4).Cf13
				_ = _842___mcc_h28
				var _843_cf13 bool = _842___mcc_h28
				_ = _843_cf13
				var _844_cf12 bool = _841___mcc_h27
				_ = _844_cf12
				var _845_cf11 bool = _840___mcc_h26
				_ = _845_cf11
				var _846_cf10 bool = _839___mcc_h25
				_ = _846_cf10
				var _847_cf9 _dafny.Int = _838___mcc_h24
				_ = _847_cf9
				var _848_v46 _dafny.Map
				_ = _848_v46
				_848_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_739_v0).Difference(_739_v0), !((p1) && (p2)))
				_848_v46 = (_848_v46).Update((_739_v0).Intersection(_dafny.MultiSetOf(_776_cf5)), !((_762_v16).Select((Companion_Default___.SafeIndex(_847_cf9, _dafny.IntOfUint32((_762_v16).Cardinality()))).Uint32()).(bool)))
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_764_v18), 0))
				_ = _index118
				(_764_v18).ArraySet1((_778_cf3) == (_778_cf3), (_index118).Int())
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_764_v18), 0))
				_ = _index119
				(_764_v18).ArraySet1(_844_cf12, (_index119).Int())
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_764_v18), 0))
				_ = _index120
				(_764_v18).ArraySet1((_764_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_764_v18), 0))).Int()).(bool), (_index120).Int())
				(globalState).F2 = p2
			} else if _source10.Is_DC1() {
				var _849___mcc_h29 _dafny.Int = _source10.Get_().(D1_DC1).Cf1
				_ = _849___mcc_h29
				var _850_cf1 _dafny.Int = _849___mcc_h29
				_ = _850_cf1
				var _851_v47 _dafny.Set
				_ = _851_v47
				_851_v47 = _dafny.SetOf(_755_v9)
				var _852_v48 *C0
				_ = _852_v48
				var _nw130 *C0 = New_C0_()
				_ = _nw130
				_nw130.Ctor__((_851_v47).Intersection(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(534))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_853_cf5 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_854_i14 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_853_cf5)
					}
				})(_776_cf5))))), _dafny.IntOfInt64(197))
				_852_v48 = _nw130
				_852_v48 = _852_v48
				var _855_v49 D1
				_ = _855_v49
				_855_v49 = Companion_D1_.Create_DC3_((_778_cf3) == (p2), _dafny.UnicodeSeqOfUtf8Bytes("srewo"), _777_cf4)
				_855_v49 = Companion_D1_.Create_DC3_(true, _dafny.UnicodeSeqOfUtf8Bytes("lrcaagrw"), _850_cf1)
				var _856_v50 _dafny.Array
				_ = _856_v50
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_23
				var _nw131 _dafny.Array
				_ = _nw131
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw131 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = (func(_857_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_858_i15 _dafny.Int) _dafny.Int {
							return (_858_i15).Times(_857_p0)
						}
					})(p0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw131 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw131).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw131).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_856_v50 = _nw131
				_856_v50 = _856_v50
			} else {
				var _859___mcc_h30 D1 = _source10.Get_().(D1_DC5).Cf14
				_ = _859___mcc_h30
				var _860_cf14 D1 = _859___mcc_h30
				_ = _860_cf14
				var _861_v51 T0
				_ = _861_v51
				var _nw132 *C1 = New_C1_()
				_ = _nw132
				_nw132.Ctor__(_758_v12)
				_861_v51 = _nw132
				var _862_v52 bool
				_ = _862_v52
				var _863_v53 bool
				_ = _863_v53
				var _out12 bool
				_ = _out12
				var _out13 bool
				_ = _out13
				_out12, _out13 = (_759_v13).M6(_861_v51, _764_v18, p1, globalState)
				_862_v52 = _out12
				_863_v53 = _out13
				var _864_v54 _dafny.Sequence
				_ = _864_v54
				_864_v54 = _dafny.UnicodeSeqOfUtf8Bytes("wslaqc")
				_864_v54 = _dafny.Companion_Sequence_.Update(_864_v54, (Companion_Default___.SafeIndex(_779_cf2, _dafny.IntOfUint32((_864_v54).Cardinality()))).Uint32(), _761_v15)
				_761_v15 = _dafny.CodePoint('r')
				_864_v54 = _864_v54
			}
			var _rhs160 bool = !(true)
			_ = _rhs160
			var _rhs161 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_777_cf4), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_755_v9).Cardinality()), p0)).Cardinality())).Times(_776_cf5))
			_ = _rhs161
			var _lhs100 *GlobalState = globalState
			_ = _lhs100
			_lhs100.F2 = _rhs160
			_777_cf4 = _rhs161
			_778_cf3 = p1
			_777_cf4 = (p0).Plus(_779_cf2)
		} else if _source8.Is_DC3() {
			var _865___mcc_h4 bool = _source8.Get_().(D1_DC3).Cf6
			_ = _865___mcc_h4
			var _866___mcc_h5 _dafny.Sequence = _source8.Get_().(D1_DC3).Cf7
			_ = _866___mcc_h5
			var _867___mcc_h6 _dafny.Int = _source8.Get_().(D1_DC3).Cf8
			_ = _867___mcc_h6
			var _868_cf8 _dafny.Int = _867___mcc_h6
			_ = _868_cf8
			var _869_cf7 _dafny.Sequence = _866___mcc_h5
			_ = _869_cf7
			var _870_cf6 bool = _865___mcc_h4
			_ = _870_cf6
			var _871_v55 _dafny.Array
			_ = _871_v55
			var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
			_ = _nw133
			_871_v55 = _nw133
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_871_v55), 0))
			_ = _index121
			(_871_v55).ArraySet1(_869_cf7, (_index121).Int())
			var _872_v56 _dafny.Array
			_ = _872_v56
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_24
			var _nw134 _dafny.Array
			_ = _nw134
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw134 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Int = (func(_873_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_874_i16 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_874_i16, _873_p0)
					}
				})(p0)
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw134 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw134).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw134).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_872_v56 = _nw134
			var _875_v57 _dafny.MultiSet
			_ = _875_v57
			_875_v57 = _dafny.MultiSetOf(p2)
			var _876_v58 _dafny.Map
			_ = _876_v58
			_876_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.SeqOf(p0, (_875_v57).Cardinality()))
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))
			_ = _index122
			(_872_v56).ArraySet1(Companion_Default___.SafeModuloInt((_760_v14).Cardinality(), (_876_v58).Cardinality()), (_index122).Int())
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_871_v55), 0))
			_ = _index123
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))
			_ = _index124
			var _rhs162 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(491))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}((func(_877_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_878_i17 _dafny.Int) _dafny.CodePoint {
					return _877_v15
				}
			})(_761_v15))), _869_cf7)
			_ = _rhs162
			var _rhs163 _dafny.Int = p0
			_ = _rhs163
			var _rhs164 _dafny.Sequence = Companion_Default___.Fm20(!(false), p1, Companion_Default___.Fm0(globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}((func(_879_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_880_i18 _dafny.Int) _dafny.CodePoint {
					return _879_v15
				}
			})(_761_v15))), globalState)
			_ = _rhs164
			var _rhs165 _dafny.Int = _dafny.IntOfInt64(812)
			_ = _rhs165
			var _lhs101 _dafny.Array = _871_v55
			_ = _lhs101
			var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_871_v55), 0))
			_ = _lhs102
			var _lhs103 _dafny.Array = _872_v56
			_ = _lhs103
			var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))
			_ = _lhs104
			(_lhs101).ArraySet1(_rhs162, (_lhs102).Int())
			(_lhs103).ArraySet1(_rhs163, (_lhs104).Int())
			_762_v16 = _rhs164
			_868_cf8 = _rhs165
			var _881_v59 _dafny.Set
			_ = _881_v59
			_881_v59 = _dafny.SetOf(_dafny.SeqOf(_870_cf6), _762_v16)
			if !(_881_v59).Contains(_762_v16) {
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))
				_ = _index125
				var _rhs166 _dafny.Int = _868_cf8
				_ = _rhs166
				var _rhs167 _dafny.Array = _872_v56
				_ = _rhs167
				var _lhs105 _dafny.Array = _872_v56
				_ = _lhs105
				var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))
				_ = _lhs106
				(_lhs105).ArraySet1(_rhs166, (_lhs106).Int())
				_872_v56 = _rhs167
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_764_v18), 0))
				_ = _index126
				(_764_v18).ArraySet1(false, (_index126).Int())
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_764_v18), 0))
				_ = _index127
				(_764_v18).ArraySet1(p2, (_index127).Int())
				_868_cf8 = (_872_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))).Int()).(_dafny.Int)
				var _882_v60 D1
				_ = _882_v60
				_882_v60 = Companion_D1_.Create_DC4_(p0, false, p1, !(_870_cf6), (_764_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_764_v18), 0))).Int()).(bool))
				_882_v60 = _882_v60
				_870_cf6 = ((func() _dafny.Int {
					if (_760_v14).Contains((_764_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_764_v18), 0))).Int()).(bool)) {
						return (_760_v14).Get((_764_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_764_v18), 0))).Int()).(bool)).(_dafny.Int)
					}
					return (_872_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))).Int()).(_dafny.Int)
				})()).Cmp(Companion_Default___.SafeDivisionInt((_872_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))).Int()).(_dafny.Int), _868_cf8)) != 0
			} else {
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))
				_ = _index128
				(_872_v56).ArraySet1((Companion_Default___.SafeModuloInt(p0, _868_cf8)).Times((_872_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))).Int()).(_dafny.Int)), (_index128).Int())
				Companion_Default___.M0(_870_cf6, globalState)
				var _883_v61 _dafny.Set
				_ = _883_v61
				_883_v61 = _dafny.SetOf(_755_v9)
				var _884_v62 *C0
				_ = _884_v62
				var _nw135 *C0 = New_C0_()
				_ = _nw135
				_nw135.Ctor__((func() _dafny.Set {
					if _870_cf6 {
						return _883_v61
					}
					return _dafny.SetOf((_this).Fm2(_868_cf8, (_this).F9(), _870_cf6, _868_cf8, globalState), _755_v9, _755_v9)
				})(), (func() _dafny.Int {
					if (_760_v14).Contains(p1) {
						return (_760_v14).Get(p1).(_dafny.Int)
					}
					return _dafny.IntOfUint32(((_871_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_871_v55), 0))).Int()).(_dafny.Sequence)).Cardinality())
				})())
				_884_v62 = _nw135
				var _885_v63 _dafny.Map
				_ = _885_v63
				_885_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_870_cf6, _872_v56)
				var _886_v64 _dafny.Array
				_ = _886_v64
				_886_v64 = (func() _dafny.Array {
					if (_885_v63).Contains(false) {
						return (_885_v63).Get(false).(_dafny.Array)
					}
					return _872_v56
				})()
				var _rhs168 _dafny.Array = (_886_v64)
				_ = _rhs168
				var _rhs169 bool = (_this).F9()
				_ = _rhs169
				var _lhs107 *GlobalState = globalState
				_ = _lhs107
				_872_v56 = _rhs168
				_lhs107.F2 = _rhs169
				var _887_v65 _dafny.Map
				_ = _887_v65
				_887_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_868_cf8, _dafny.UnicodeSeqOfUtf8Bytes("sfasfg"))
				_887_v65 = (_887_v65).Update(_dafny.IntOfInt64(568), _dafny.Companion_Sequence_.Concatenate(_869_cf7, _869_cf7))
			}
			_868_cf8 = (_872_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))).Int()).(_dafny.Int)
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_872_v56), 0))
			_ = _index129
			(_872_v56).ArraySet1(((_875_v57).Union((_875_v57).Difference(Companion_Default___.Fm21(_dafny.IntOfUint32((_755_v9).Cardinality()), _870_cf6, globalState)))).Cardinality(), (_index129).Int())
		} else if _source8.Is_DC4() {
			var _888___mcc_h7 _dafny.Int = _source8.Get_().(D1_DC4).Cf9
			_ = _888___mcc_h7
			var _889___mcc_h8 bool = _source8.Get_().(D1_DC4).Cf10
			_ = _889___mcc_h8
			var _890___mcc_h9 bool = _source8.Get_().(D1_DC4).Cf11
			_ = _890___mcc_h9
			var _891___mcc_h10 bool = _source8.Get_().(D1_DC4).Cf12
			_ = _891___mcc_h10
			var _892___mcc_h11 bool = _source8.Get_().(D1_DC4).Cf13
			_ = _892___mcc_h11
			var _893_cf13 bool = _892___mcc_h11
			_ = _893_cf13
			var _894_cf12 bool = _891___mcc_h10
			_ = _894_cf12
			var _895_cf11 bool = _890___mcc_h9
			_ = _895_cf11
			var _896_cf10 bool = _889___mcc_h8
			_ = _896_cf10
			var _897_cf9 _dafny.Int = _888___mcc_h7
			_ = _897_cf9
			var _898_v66 _dafny.Map
			_ = _898_v66
			_898_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_764_v18, _897_cf9)
			if ((_898_v66).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_755_v9, _755_v9)).Cardinality())) != 0 {
				var _899_v67 _dafny.Array
				_ = _899_v67
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_25
				var _nw136 _dafny.Array
				_ = _nw136
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw136 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.Sequence = (func(_900_v9 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_901_i19 _dafny.Int) _dafny.Sequence {
							return _900_v9
						}
					})(_755_v9)
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw136 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw136).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw136).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_899_v67 = _nw136
				var _902_v68 _dafny.Sequence
				_ = _902_v68
				_902_v68 = _755_v9
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_899_v67), 0))
				_ = _index130
				(_899_v67).ArraySet1(_902_v68, (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_899_v67), 0))
				_ = _index131
				(_899_v67).ArraySet1(_902_v68, (_index131).Int())
				_760_v14 = (_760_v14).Update(_895_cf11, p0)
				var _903_v69 _dafny.Set
				_ = _903_v69
				_903_v69 = _dafny.SetOf(_dafny.SeqOf(_897_cf9))
				var _904_v70 *C0
				_ = _904_v70
				var _nw137 *C0 = New_C0_()
				_ = _nw137
				_nw137.Ctor__(_903_v69, _dafny.IntOfInt64(663))
				_904_v70 = _nw137
				Companion_Default___.M0(p1, globalState)
				(_904_v70).F12 = p0
			} else {
				var _905_v71 _dafny.Sequence
				_ = _905_v71
				_905_v71 = _dafny.UnicodeSeqOfUtf8Bytes("j")
				_905_v71 = _dafny.UnicodeSeqOfUtf8Bytes("hick")
				var _906_v72 _dafny.Array
				_ = _906_v72
				var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw138
				_906_v72 = _nw138
				var _907_v73 _dafny.Array
				_ = _907_v73
				var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
				_ = _nw139
				_907_v73 = _nw139
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_906_v72), 0))
				_ = _index132
				(_906_v72).ArraySet1(_907_v73, (_index132).Int())
				var _908_v74 _dafny.Sequence
				_ = _908_v74
				_908_v74 = _dafny.SeqOf(_907_v73)
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_906_v72), 0))
				_ = _index133
				var _rhs170 _dafny.Array = (func() _dafny.Array {
					if Companion_Default___.Fm1(_739_v0, _897_cf9, globalState) {
						return _907_v73
					}
					return (_908_v74).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.IntOfUint32((_908_v74).Cardinality()))).Uint32()).(_dafny.Array)
				})()
				_ = _rhs170
				var _rhs171 _dafny.Array = _764_v18
				_ = _rhs171
				var _rhs172 bool = (_this).F9()
				_ = _rhs172
				var _lhs108 _dafny.Array = _906_v72
				_ = _lhs108
				var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_906_v72), 0))
				_ = _lhs109
				(_lhs108).ArraySet1(_rhs170, (_lhs109).Int())
				_764_v18 = _rhs171
				_893_cf13 = _rhs172
				var _909_v75 D7
				_ = _909_v75
				_909_v75 = Companion_D7_.Create_DC18_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_739_v0, p1)).Merge(Companion_Default___.Fm22(globalState)))
				var _910_v76 _dafny.Map
				_ = _910_v76
				_910_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_739_v0, _893_cf13)
				_909_v75 = Companion_D7_.Create_DC18_(_910_v76)
				var _911_v77 _dafny.Map
				_ = _911_v77
				_911_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _897_cf9)
				var _912_v78 _dafny.MultiSet
				_ = _912_v78
				_912_v78 = _dafny.MultiSetOf(_911_v77)
				var _913_v79 _dafny.Map
				_ = _913_v79
				_913_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_905_v71, _897_cf9)
				_912_v78 = ((_912_v78).Difference(_912_v78)).Update(_911_v77, Companion_Default___.Abs(((_755_v9).Select((Companion_Default___.SafeIndex(_897_cf9, _dafny.IntOfUint32((_755_v9).Cardinality()))).Uint32()).(_dafny.Int)).Times((_913_v79).Cardinality())))
				_896_cf10 = (_896_cf10) == (p1)
			}
			var _914_v80 _dafny.Set
			_ = _914_v80
			_914_v80 = _dafny.SetOf(_895_cf11, (_this).F9(), _893_cf13, p2, _896_cf10)
			var _915_v81 _dafny.Map
			_ = _915_v81
			_915_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v80, _896_cf10)
			_897_cf9 = (p0).Minus((_dafny.Zero).Minus((_915_v81).Cardinality()))
			var _916_v82 _dafny.Sequence
			_ = _916_v82
			_916_v82 = _dafny.UnicodeSeqOfUtf8Bytes("l")
			var _917_v83 _dafny.Sequence
			_ = _917_v83
			_917_v83 = _dafny.SeqOf(_916_v82, _916_v82)
			_897_cf9 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_916_v82), _917_v83)).Cardinality())
			var _918_v84 _dafny.Array
			_ = _918_v84
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_26
			var _nw140 _dafny.Array
			_ = _nw140
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw140 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) _dafny.Sequence = (func(_919_cf9 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_920_i20 _dafny.Int) _dafny.Sequence {
						return (_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-154))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg59 _dafny.Int) interface{} {
								return coer59(arg59)
							}
						}((func(_921_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_922_i21 _dafny.Int) _dafny.Int {
								return _921_cf9
							}
						})(_919_cf9))))
					}
				})(_897_cf9)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw140 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw140).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw140).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_918_v84 = _nw140
			_918_v84 = _918_v84
		} else if _source8.Is_DC1() {
			var _923___mcc_h12 _dafny.Int = _source8.Get_().(D1_DC1).Cf1
			_ = _923___mcc_h12
			var _924_cf1 _dafny.Int = _923___mcc_h12
			_ = _924_cf1
			var _925_v85 _dafny.Array
			_ = _925_v85
			var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw141
			_925_v85 = _nw141
			var _926_v86 _dafny.Map
			_ = _926_v86
			_926_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !(p2))
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_925_v85), 0))
			_ = _index134
			(_925_v85).ArraySet1((_926_v86).Cardinality(), (_index134).Int())
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_925_v85), 0))
			_ = _index135
			(_925_v85).ArraySet1(_924_cf1, (_index135).Int())
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_925_v85), 0))
			_ = _index136
			(_925_v85).ArraySet1(_924_cf1, (_index136).Int())
			var _rhs173 bool = (_this).F9()
			_ = _rhs173
			var _rhs174 bool = !(!(Companion_Default___.Fm1(_739_v0, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-397), (_925_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_925_v85), 0))).Int()).(_dafny.Int)), globalState)))
			_ = _rhs174
			var _lhs110 *GlobalState = globalState
			_ = _lhs110
			var _lhs111 *GlobalState = globalState
			_ = _lhs111
			_lhs110.F2 = _rhs173
			_lhs111.F2 = _rhs174
			(globalState).F2 = p2
		} else {
			var _927___mcc_h13 D1 = _source8.Get_().(D1_DC5).Cf14
			_ = _927___mcc_h13
			var _928_cf14 D1 = _927___mcc_h13
			_ = _928_cf14
			var _929_v87 _dafny.Array
			_ = _929_v87
			var _nw142 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw142
			_929_v87 = _nw142
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_929_v87), 0))
			_ = _index137
			(_929_v87).ArraySet1(p0, (_index137).Int())
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_929_v87), 0))
			_ = _index138
			(_929_v87).ArraySet1(p0, (_index138).Int())
			var _930_v88 _dafny.Sequence
			_ = _930_v88
			_930_v88 = _dafny.UnicodeSeqOfUtf8Bytes("j")
			var _931_v90 _dafny.Set
			_ = _931_v90
			_931_v90 = _dafny.SetOf(_dafny.IntOfInt64(-638))
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_929_v87), 0))
			_ = _index139
			var _rhs175 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.Fm23(!((func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(319), _dafny.IntOfInt64(877))); ; {
					_compr_22, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _932_v89 _dafny.Int
					_932_v89 = interface{}(_compr_22).(_dafny.Int)
					if ((_dafny.IntOfInt64(319)).Cmp(_932_v89) <= 0) && ((_932_v89).Cmp(_dafny.IntOfInt64(877)) < 0) {
						_coll22.Add(Companion_Default___.SafeModuloInt(_932_v89, (_929_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_929_v87), 0))).Int()).(_dafny.Int)))
					}
				}
				return _coll22.ToSet()
			}()).IsSubsetOf(_931_v90)), !_dafny.Companion_Sequence_.Equal(_930_v88, _dafny.UnicodeSeqOfUtf8Bytes("tnvrovbj")), p0, globalState)).Cardinality())
			_ = _rhs175
			var _rhs176 bool = Companion_Default___.Fm1(_dafny.MultiSetOf((_929_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_929_v87), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(844)), Companion_Default___.SafeDivisionInt(p0, p0), globalState)
			_ = _rhs176
			var _rhs177 bool = ((_this).F9()) && ((_this).F9())
			_ = _rhs177
			var _rhs178 _dafny.Sequence = _930_v88
			_ = _rhs178
			var _rhs179 _dafny.Sequence = _930_v88
			_ = _rhs179
			var _lhs112 _dafny.Array = _929_v87
			_ = _lhs112
			var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_929_v87), 0))
			_ = _lhs113
			var _lhs114 *GlobalState = globalState
			_ = _lhs114
			var _lhs115 *GlobalState = globalState
			_ = _lhs115
			(_lhs112).ArraySet1(_rhs175, (_lhs113).Int())
			_lhs114.F2 = _rhs176
			_lhs115.F2 = _rhs177
			_930_v88 = _rhs178
			_930_v88 = _rhs179
			var _933_v91 _dafny.Array
			_ = _933_v91
			var _nw143 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
			_ = _nw143
			_933_v91 = _nw143
			var _934_v92 _dafny.Map
			_ = _934_v92
			_934_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_755_v9, _739_v0)
			var _935_v93 T1
			_ = _935_v93
			var _nw144 *C2 = New_C2_()
			_ = _nw144
			_nw144.Ctor__(((func() _dafny.MultiSet {
				if (_934_v92).Contains(_755_v9) {
					return (_934_v92).Get(_755_v9).(_dafny.MultiSet)
				}
				return _739_v0
			})()).Cardinality(), p1)
			_935_v93 = _nw144
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_933_v91), 0))
			_ = _index140
			(_933_v91).ArraySet1(_935_v93, (_index140).Int())
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_933_v91), 0))
			_ = _index141
			(_933_v91).ArraySet1(_935_v93, (_index141).Int())
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_929_v87), 0))
			_ = _index142
			(_929_v87).ArraySet1((_929_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_929_v87), 0))).Int()).(_dafny.Int), (_index142).Int())
		}
		_755_v9 = _755_v9
	}
}
func (_this *C4) M7(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) (_dafny.Int, D4, _dafny.MultiSet, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 D4 = Companion_D4_.Default()
		_ = r1
		var r2 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r2
		var r3 bool = false
		_ = r3
		if !((_this).F9()) || (!((_this).F9())) {
			var _936_v0 _dafny.MultiSet
			_ = _936_v0
			_936_v0 = _dafny.MultiSetOf(_dafny.IntOfInt64(230), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_this).F9())).Cardinality(), p0)).Cardinality())
			var _937_v1 _dafny.Sequence
			_ = _937_v1
			_937_v1 = _dafny.SeqOf(p2, p2)
			var _938_v2 _dafny.Map
			_ = _938_v2
			_938_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
			var _939_v3 _dafny.MultiSet
			_ = _939_v3
			_939_v3 = _dafny.MultiSetOf((_this).F9())
			var _940_v4 _dafny.Map
			_ = _940_v4
			_940_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), p0)
			var _941_v5 _dafny.Map
			_ = _941_v5
			_941_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _942_v6 _dafny.Set
			_ = _942_v6
			_942_v6 = _dafny.SetOf(_941_v5)
			var _943_v7 _dafny.Array
			_ = _943_v7
			var _nwElement0_28 _dafny.Int = (_dafny.IntOfInt64(261)).Times(p0)
			_ = _nwElement0_28
			var _nw145 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(28))
			_ = _nw145
			(_nw145).ArraySet1(_nwElement0_28, 0)
			(_nw145).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(295), p1), 1)
			(_nw145).ArraySet1(p1, 2)
			(_nw145).ArraySet1(p1, 3)
			(_nw145).ArraySet1(p0, 4)
			(_nw145).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_936_v0).Contains(p0) {
					return (_936_v0).Multiplicity(p0)
				}
				return p0
			})(), Companion_Default___.Fm0(globalState))), 5)
			(_nw145).ArraySet1(p1, 6)
			(_nw145).ArraySet1(p0, 7)
			(_nw145).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_937_v1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_937_v1).Cardinality()))).Uint32(), p2)).Cardinality()), (_dafny.Zero).Minus(p0)), 8)
			(_nw145).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), 9)
			(_nw145).ArraySet1(p0, 10)
			(_nw145).ArraySet1(p0, 11)
			(_nw145).ArraySet1(p1, 12)
			(_nw145).ArraySet1(p0, 13)
			(_nw145).ArraySet1(_dafny.IntOfInt64(898), 14)
			(_nw145).ArraySet1(p1, 15)
			(_nw145).ArraySet1(p1, 16)
			(_nw145).ArraySet1(Companion_Default___.SafeModuloInt((_938_v2).Cardinality(), (_939_v3).Cardinality()), 17)
			(_nw145).ArraySet1(p0, 18)
			(_nw145).ArraySet1((_dafny.Zero).Minus(p1), 19)
			(_nw145).ArraySet1(p1, 20)
			(_nw145).ArraySet1(((_940_v4).Merge(Companion_Default___.Fm23((_this).F9(), p2, p0, globalState))).Cardinality(), 21)
			(_nw145).ArraySet1(_dafny.IntOfInt64(899), 22)
			(_nw145).ArraySet1((p1).Minus(_dafny.IntOfInt64(944)), 23)
			(_nw145).ArraySet1((func() _dafny.Int {
				if (_this).F9() {
					return _dafny.IntOfInt64(229)
				}
				return p0
			})(), 24)
			(_nw145).ArraySet1(p0, 25)
			(_nw145).ArraySet1(p1, 26)
			(_nw145).ArraySet1((p1).Minus((_942_v6).Cardinality()), 27)
			_943_v7 = _nw145
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))
			_ = _index143
			(_943_v7).ArraySet1(p1, (_index143).Int())
			var _944_v8 _dafny.Map
			_ = _944_v8
			_944_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(p0, p1, p0), p2)
			var _945_v9 D7
			_ = _945_v9
			_945_v9 = Companion_D7_.Create_DC18_((_944_v8).Update(_936_v0, p2))
			var _946_v10 _dafny.Array
			_ = _946_v10
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_27
			var _nw146 _dafny.Array
			_ = _nw146
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw146 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) D7 = func(_947_i0 _dafny.Int) D7 {
					return Companion_D7_.Create_DC20_((_this).F9(), _dafny.CodePoint('q'))
				}
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw146 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw146).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw146).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_946_v10 = _nw146
			var _948_v11 _dafny.Sequence
			_ = _948_v11
			_948_v11 = _dafny.SeqOf((_941_v5).Merge(_941_v5))
			var _949_v12 _dafny.Sequence
			_ = _949_v12
			_949_v12 = _dafny.SeqOf(p1)
			var _950_v13 _dafny.Sequence
			_ = _950_v13
			_950_v13 = _dafny.UnicodeSeqOfUtf8Bytes("kcofbj")
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))
			_ = _index144
			var _rhs180 _dafny.Int = _dafny.IntOfUint32((_948_v11).Cardinality())
			_ = _rhs180
			var _rhs181 _dafny.Sequence = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm20((p1).Cmp(_dafny.IntOfUint32((_937_v1).Cardinality())) != 0, Companion_Default___.Fm1(Companion_Default___.Fm29(p2, (_this).F9(), globalState), (_949_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.IntOfUint32((_949_v12).Cardinality()))).Uint32()).(_dafny.Int), globalState), (p0).Plus(_dafny.IntOfUint32((_937_v1).Cardinality())), _950_v13, globalState), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((Companion_Default___.Fm20((p1).Cmp(_dafny.IntOfUint32((_937_v1).Cardinality())) != 0, Companion_Default___.Fm1(Companion_Default___.Fm29(p2, (_this).F9(), globalState), (_949_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.IntOfUint32((_949_v12).Cardinality()))).Uint32()).(_dafny.Int), globalState), (p0).Plus(_dafny.IntOfUint32((_937_v1).Cardinality())), _950_v13, globalState)).Cardinality()))).Uint32(), p2)
			_ = _rhs181
			var _rhs182 D7 = Companion_D7_.Create_DC18_(_944_v8)
			_ = _rhs182
			var _rhs183 _dafny.Int = p0
			_ = _rhs183
			var _rhs184 _dafny.Array = _946_v10
			_ = _rhs184
			var _lhs116 _dafny.Array = _943_v7
			_ = _lhs116
			var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))
			_ = _lhs117
			(_lhs116).ArraySet1(_rhs180, (_lhs117).Int())
			_937_v1 = _rhs181
			_945_v9 = _rhs182
			r0 = _rhs183
			_946_v10 = _rhs184
			var _951_v14 _dafny.Sequence
			_ = _951_v14
			_951_v14 = _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('c')))
			var _952_v15 D1
			_ = _952_v15
			_952_v15 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_950_v13).Cardinality()), (_this).F9(), _dafny.IntOfInt64(503), p0)
			var _953_v16 _dafny.Sequence
			_ = _953_v16
			_953_v16 = _dafny.SeqOf(func(_pat_let24_0 D1) D1 {
				return func(_954_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let25_0 bool) D1 {
						return func(_955_dt__update_hcf3_h0 bool) D1 {
							return Companion_D1_.Create_DC2_((_954_dt__update__tmp_h0).Dtor_cf2(), _955_dt__update_hcf3_h0, (_954_dt__update__tmp_h0).Dtor_cf4(), (_954_dt__update__tmp_h0).Dtor_cf5())
						}(_pat_let25_0)
					}((_this).F9())
				}(_pat_let24_0)
			}(_952_v15), _952_v15)
			var _956_v17 _dafny.CodePoint
			_ = _956_v17
			_956_v17 = _dafny.CodePoint('e')
			var _pat_let_tv26 = _943_v7
			_ = _pat_let_tv26
			var _pat_let_tv27 = _943_v7
			_ = _pat_let_tv27
			var _pat_let_tv28 = p0
			_ = _pat_let_tv28
			var _pat_let_tv29 = _943_v7
			_ = _pat_let_tv29
			var _pat_let_tv30 = _943_v7
			_ = _pat_let_tv30
			var _pat_let_tv31 = p0
			_ = _pat_let_tv31
			var _pat_let_tv32 = p1
			_ = _pat_let_tv32
			var _pat_let_tv33 = _943_v7
			_ = _pat_let_tv33
			var _pat_let_tv34 = _943_v7
			_ = _pat_let_tv34
			var _pat_let_tv35 = p0
			_ = _pat_let_tv35
			var _pat_let_tv36 = p0
			_ = _pat_let_tv36
			var _pat_let_tv37 = _943_v7
			_ = _pat_let_tv37
			var _pat_let_tv38 = _943_v7
			_ = _pat_let_tv38
			var _pat_let_tv39 = p0
			_ = _pat_let_tv39
			var _pat_let_tv40 = p1
			_ = _pat_let_tv40
			var _pat_let_tv41 = _943_v7
			_ = _pat_let_tv41
			var _pat_let_tv42 = _943_v7
			_ = _pat_let_tv42
			var _pat_let_tv43 = p1
			_ = _pat_let_tv43
			var _pat_let_tv44 = _956_v17
			_ = _pat_let_tv44
			var _pat_let_tv45 = _936_v0
			_ = _pat_let_tv45
			var _pat_let_tv46 = globalState
			_ = _pat_let_tv46
			var _957_v18 *C1
			_ = _957_v18
			var _nw147 *C1 = New_C1_()
			_ = _nw147
			_nw147.Ctor__((func() _dafny.Sequence {
				if Companion_Default___.Fm1(_dafny.MultiSetOf(_dafny.IntOfInt64(637), ((_951_v14).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_951_v14).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), (_dafny.Zero).Minus((_943_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))).Int()).(_dafny.Int))), p0, globalState) {
					return _953_v16
				}
				return _dafny.SeqOf(func(_pat_let26_0 D1) D1 {
					return func(_958_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let27_0 _dafny.Int) D1 {
							return func(_959_dt__update_hcf5_h0 _dafny.Int) D1 {
								return func(_pat_let28_0 _dafny.Int) D1 {
									return func(_962_dt__update_hcf4_h0 _dafny.Int) D1 {
										return func(_pat_let29_0 bool) D1 {
											return func(_964_dt__update_hcf3_h1 bool) D1 {
												return Companion_D1_.Create_DC2_((_958_dt__update__tmp_h1).Dtor_cf2(), _964_dt__update_hcf3_h1, _962_dt__update_hcf4_h0, _959_dt__update_hcf5_h0)
											}(_pat_let29_0)
										}(Companion_Default___.Fm1(_pat_let_tv45, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-302))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
											return func(arg62 _dafny.Int) interface{} {
												return coer62(arg62)
											}
										}(func(_963_i1 _dafny.Int) _dafny.CodePoint {
											return _dafny.CodePoint('b')
										}))).Cardinality()), _pat_let_tv46))
									}(_pat_let28_0)
								}(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(949))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg60 _dafny.Int) interface{} {
										return coer60(arg60)
									}
								}(func(_960_i2 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('c')
								})), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_pat_let_tv28, (_pat_let_tv30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_pat_let_tv29), 0))).Int()).(_dafny.Int), _pat_let_tv31, _pat_let_tv32, (_pat_let_tv34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_pat_let_tv33), 0))).Int()).(_dafny.Int)), (Companion_Default___.SafeIndex(_pat_let_tv35, _dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv36, (_pat_let_tv38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_pat_let_tv37), 0))).Int()).(_dafny.Int), _pat_let_tv39, _pat_let_tv40, (_pat_let_tv42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_pat_let_tv41), 0))).Int()).(_dafny.Int))).Cardinality()))).Uint32(), _pat_let_tv43)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(949))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg61 _dafny.Int) interface{} {
										return coer61(arg61)
									}
								}(func(_961_i2 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('c')
								}))).Cardinality()))).Uint32(), _pat_let_tv44)).Cardinality()))
							}(_pat_let27_0)
						}((_pat_let_tv27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_pat_let_tv26), 0))).Int()).(_dafny.Int))
					}(_pat_let26_0)
				}(_952_v15), _952_v15, _952_v15, _952_v15, _952_v15)
			})())
			_957_v18 = _nw147
			var _965_v19 _dafny.Sequence
			_ = _965_v19
			_965_v19 = _dafny.SeqOf(Companion_Default___.Fm30(globalState))
			_938_v2 = (_965_v19).Select((Companion_Default___.SafeIndex((_943_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_965_v19).Cardinality()))).Uint32()).(_dafny.Map)
			var _966_v20 _dafny.Map
			_ = _966_v20
			_966_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_943_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))).Int()).(_dafny.Int)), _956_v17)
			var _967_v21 _dafny.Sequence
			_ = _967_v21
			_967_v21 = _dafny.SeqOf(_940_v4, _940_v4, _940_v4)
			var _968_v22 _dafny.Set
			_ = _968_v22
			_968_v22 = _dafny.SetOf(p2, (_this).F9(), p2)
			var _969_v23 _dafny.Sequence
			_ = _969_v23
			_969_v23 = _949_v12
			var _970_v24 _dafny.Array
			_ = _970_v24
			var _nwElement0_29 bool = (_939_v3).IsDisjointFrom(_dafny.MultiSetOf((_this).F9()))
			_ = _nwElement0_29
			var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(19))
			_ = _nw148
			(_nw148).ArraySet1(_nwElement0_29, 0)
			(_nw148).ArraySet1(p2, 1)
			(_nw148).ArraySet1(!((_this).F9()), 2)
			(_nw148).ArraySet1(p2, 3)
			(_nw148).ArraySet1((p1).Cmp((_966_v20).Cardinality()) == 0, 4)
			(_nw148).ArraySet1((_this).F9(), 5)
			(_nw148).ArraySet1((_this).F9(), 6)
			(_nw148).ArraySet1((_this).F9(), 7)
			(_nw148).ArraySet1(p2, 8)
			(_nw148).ArraySet1(Companion_Default___.Fm1((_dafny.MultiSetOf(p1)).Update(p0, Companion_Default___.Abs(((_967_v21).Select((Companion_Default___.SafeIndex((_943_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_967_v21).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())), (_968_v22).Cardinality(), globalState), 9)
			(_nw148).ArraySet1((_this).F9(), 10)
			(_nw148).ArraySet1(p2, 11)
			(_nw148).ArraySet1(p2, 12)
			(_nw148).ArraySet1((_this).F9(), 13)
			(_nw148).ArraySet1((_dafny.IntOfInt64(559)).Cmp((_dafny.Zero).Minus(p1)) == 0, 14)
			(_nw148).ArraySet1(false, 15)
			(_nw148).ArraySet1((p2) == (!(p2)), 16)
			(_nw148).ArraySet1(false, 17)
			(_nw148).ArraySet1(_dafny.Companion_Sequence_.Contains((_dafny.SeqOf((_943_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))).Int()).(_dafny.Int), (_943_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))).Int()).(_dafny.Int))), p0), 18)
			_970_v24 = _nw148
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_970_v24), 0))
			_ = _index145
			(_970_v24).ArraySet1((_this).F9(), (_index145).Int())
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_970_v24), 0))
			_ = _index146
			(_970_v24).ArraySet1(!(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-568))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}((func(_971_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_972_i3 _dafny.Int) _dafny.CodePoint {
					return _971_v17
				}
			})(_956_v17))), _956_v17)), (_index146).Int())
			var _973_v25 _dafny.Map
			_ = _973_v25
			_973_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _956_v17)
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_970_v24), 0))
			_ = _index147
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_970_v24), 0))
			_ = _index148
			var _rhs185 _dafny.Int = p1
			_ = _rhs185
			var _rhs186 bool = !(_939_v3).Contains(p2)
			_ = _rhs186
			var _rhs187 bool = (p0).Cmp(((_943_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_943_v7), 0))).Int()).(_dafny.Int)).Times((_973_v25).Cardinality())) == 0
			_ = _rhs187
			var _rhs188 bool = !(p2) || (false)
			_ = _rhs188
			var _rhs189 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_950_v13, Companion_Default___.Fm28(p1, !((_this).F9()), globalState)), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_950_v13, Companion_Default___.Fm28(p1, !((_this).F9()), globalState))).Cardinality()))).Uint32(), _956_v17), _950_v13)
			_ = _rhs189
			var _lhs118 *GlobalState = globalState
			_ = _lhs118
			var _lhs119 _dafny.Array = _970_v24
			_ = _lhs119
			var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_970_v24), 0))
			_ = _lhs120
			var _lhs121 _dafny.Array = _970_v24
			_ = _lhs121
			var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_970_v24), 0))
			_ = _lhs122
			r0 = _rhs185
			_lhs118.F2 = _rhs186
			(_lhs119).ArraySet1(_rhs187, (_lhs120).Int())
			(_lhs121).ArraySet1(_rhs188, (_lhs122).Int())
			_950_v13 = _rhs189
			var _974_v26 D10
			_ = _974_v26
			_974_v26 = Companion_D10_.Create_DC27_(_940_v4)
			_940_v4 = (((_974_v26).Dtor_cf50()).Merge(_940_v4)).Merge(_940_v4)
		} else {
			r0 = (p1).Times(p1)
			var _975_v27 _dafny.Sequence
			_ = _975_v27
			_975_v27 = _dafny.SeqOf(p2, true)
			var _976_v28 _dafny.Set
			_ = _976_v28
			_976_v28 = _dafny.SetOf(p2)
			var _977_v29 _dafny.Array
			_ = _977_v29
			var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
			_ = _nw149
			_977_v29 = _nw149
			var _978_v30 _dafny.Sequence
			_ = _978_v30
			_978_v30 = _dafny.UnicodeSeqOfUtf8Bytes("n")
			var _979_v31 D7
			_ = _979_v31
			_979_v31 = Companion_D7_.Create_DC19_(_977_v29, (_this).F9(), _978_v30)
			var _980_v32 D7
			_ = _980_v32
			_980_v32 = Companion_D7_.Create_DC21_(_979_v31)
			var _981_v33 _dafny.Set
			_ = _981_v33
			_981_v33 = _dafny.SetOf(_980_v32, _980_v32)
			var _982_v34 _dafny.MultiSet
			_ = _982_v34
			_982_v34 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0))
			var _983_v35 _dafny.Array
			_ = _983_v35
			var _nwElement0_30 bool = (_975_v27).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_975_v27).Cardinality()))).Uint32()).(bool)
			_ = _nwElement0_30
			var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(19))
			_ = _nw150
			(_nw150).ArraySet1(_nwElement0_30, 0)
			(_nw150).ArraySet1((_this).F9(), 1)
			(_nw150).ArraySet1(p2, 2)
			(_nw150).ArraySet1(p2, 3)
			(_nw150).ArraySet1((_976_v28).IsSubsetOf(_976_v28), 4)
			(_nw150).ArraySet1((_981_v33).IsProperSubsetOf(_981_v33), 5)
			(_nw150).ArraySet1(true, 6)
			(_nw150).ArraySet1(Companion_Default___.Fm1(_982_v34, (_dafny.Zero).Minus(p0), globalState), 7)
			(_nw150).ArraySet1((_this).F9(), 8)
			(_nw150).ArraySet1((_this).F9(), 9)
			(_nw150).ArraySet1((_this).F9(), 10)
			(_nw150).ArraySet1(p2, 11)
			(_nw150).ArraySet1(!((_this).F9()) || (true), 12)
			(_nw150).ArraySet1(!(false), 13)
			(_nw150).ArraySet1(!(false), 14)
			(_nw150).ArraySet1(!((_976_v28).IsProperSubsetOf(_976_v28)), 15)
			(_nw150).ArraySet1((_this).F9(), 16)
			(_nw150).ArraySet1(p2, 17)
			(_nw150).ArraySet1(p2, 18)
			_983_v35 = _nw150
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))
			_ = _index149
			(_983_v35).ArraySet1((_this).F9(), (_index149).Int())
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))
			_ = _index150
			(_983_v35).ArraySet1((_this).F9(), (_index150).Int())
			var _984_v36 _dafny.Set
			_ = _984_v36
			_984_v36 = _dafny.SetOf(_dafny.IntOfUint32((_978_v30).Cardinality()), p1)
			var _985_v37 _dafny.Map
			_ = _985_v37
			_985_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_984_v36, p2)
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))
			_ = _index151
			(_983_v35).ArraySet1(!((func() bool {
				if (_985_v37).Contains(_984_v36) {
					return (_985_v37).Get(_984_v36).(bool)
				}
				return !(!(!((_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool)))) || (true)
			})()), (_index151).Int())
			var _986_v38 _dafny.CodePoint
			_ = _986_v38
			_986_v38 = _dafny.CodePoint('y')
			var _987_v39 D1
			_ = _987_v39
			_987_v39 = Companion_D1_.Create_DC3_((_this).F9(), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("vwjn"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vwjn")).Cardinality()))).Uint32(), _986_v38), p0)
			var _988_v40 _dafny.Set
			_ = _988_v40
			_988_v40 = _dafny.SetOf(_986_v38, _dafny.CodePoint('j'), _986_v38, _986_v38)
			var _989_v41 _dafny.Sequence
			_ = _989_v41
			_989_v41 = _dafny.SeqOf(p1)
			var _990_v42 _dafny.Map
			_ = _990_v42
			_990_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _978_v30)
			var _pat_let_tv47 = _978_v30
			_ = _pat_let_tv47
			var _pat_let_tv48 = p1
			_ = _pat_let_tv48
			var _pat_let_tv49 = p1
			_ = _pat_let_tv49
			var _pat_let_tv50 = p1
			_ = _pat_let_tv50
			var _pat_let_tv51 = _988_v40
			_ = _pat_let_tv51
			var _pat_let_tv52 = p1
			_ = _pat_let_tv52
			var _pat_let_tv53 = p2
			_ = _pat_let_tv53
			var _pat_let_tv54 = _978_v30
			_ = _pat_let_tv54
			var _pat_let_tv55 = p1
			_ = _pat_let_tv55
			var _pat_let_tv56 = p1
			_ = _pat_let_tv56
			var _pat_let_tv57 = p1
			_ = _pat_let_tv57
			var _pat_let_tv58 = _988_v40
			_ = _pat_let_tv58
			var _pat_let_tv59 = p1
			_ = _pat_let_tv59
			var _pat_let_tv60 = p2
			_ = _pat_let_tv60
			var _991_v43 _dafny.Array
			_ = _991_v43
			var _nwElement0_31 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_978_v30, _978_v30)
			_ = _nwElement0_31
			var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(18))
			_ = _nw151
			(_nw151).ArraySet1(_nwElement0_31, 0)
			(_nw151).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func(_pat_let30_0 D1) D1 {
				return func(_995_dt__update__tmp_h4 D1) D1 {
					return func(_pat_let34_0 _dafny.Int) D1 {
						return func(_996_dt__update_hcf8_h0 _dafny.Int) D1 {
							return func(_pat_let35_0 bool) D1 {
								return func(_997_dt__update_hcf6_h1 bool) D1 {
									return Companion_D1_.Create_DC3_(_997_dt__update_hcf6_h1, (_995_dt__update__tmp_h4).Dtor_cf7(), _996_dt__update_hcf8_h0)
								}(_pat_let35_0)
							}(_pat_let_tv53)
						}(_pat_let34_0)
					}(_dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv48, _pat_let_tv49, _pat_let_tv50, (_pat_let_tv51).Cardinality(), _pat_let_tv52)).Cardinality()))
				}(_pat_let30_0)
			}(func(_pat_let31_0 D1) D1 {
				return func(_992_dt__update__tmp_h3 D1) D1 {
					return func(_pat_let32_0 _dafny.Sequence) D1 {
						return func(_993_dt__update_hcf7_h0 _dafny.Sequence) D1 {
							return func(_pat_let33_0 bool) D1 {
								return func(_994_dt__update_hcf6_h0 bool) D1 {
									return Companion_D1_.Create_DC3_(_994_dt__update_hcf6_h0, _993_dt__update_hcf7_h0, (_992_dt__update__tmp_h3).Dtor_cf8())
								}(_pat_let33_0)
							}(true)
						}(_pat_let32_0)
					}(_pat_let_tv47)
				}(_pat_let31_0)
			}(_987_v39))).Dtor_cf7(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_989_v41).Cardinality()), _dafny.IntOfUint32(((func(_pat_let36_0 D1) D1 {
				return func(_1001_dt__update__tmp_h6 D1) D1 {
					return func(_pat_let40_0 _dafny.Int) D1 {
						return func(_1002_dt__update_hcf8_h1 _dafny.Int) D1 {
							return func(_pat_let41_0 bool) D1 {
								return func(_1003_dt__update_hcf6_h3 bool) D1 {
									return Companion_D1_.Create_DC3_(_1003_dt__update_hcf6_h3, (_1001_dt__update__tmp_h6).Dtor_cf7(), _1002_dt__update_hcf8_h1)
								}(_pat_let41_0)
							}(_pat_let_tv60)
						}(_pat_let40_0)
					}(_dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv55, _pat_let_tv56, _pat_let_tv57, (_pat_let_tv58).Cardinality(), _pat_let_tv59)).Cardinality()))
				}(_pat_let36_0)
			}(func(_pat_let37_0 D1) D1 {
				return func(_998_dt__update__tmp_h5 D1) D1 {
					return func(_pat_let38_0 _dafny.Sequence) D1 {
						return func(_999_dt__update_hcf7_h1 _dafny.Sequence) D1 {
							return func(_pat_let39_0 bool) D1 {
								return func(_1000_dt__update_hcf6_h2 bool) D1 {
									return Companion_D1_.Create_DC3_(_1000_dt__update_hcf6_h2, _999_dt__update_hcf7_h1, (_998_dt__update__tmp_h5).Dtor_cf8())
								}(_pat_let39_0)
							}(true)
						}(_pat_let38_0)
					}(_pat_let_tv54)
				}(_pat_let37_0)
			}(_987_v39))).Dtor_cf7()).Cardinality()))).Uint32(), _986_v38), _dafny.UnicodeSeqOfUtf8Bytes("ltwca")), 1)
			(_nw151).ArraySet1(_978_v30, 2)
			(_nw151).ArraySet1(_978_v30, 3)
			(_nw151).ArraySet1(_978_v30, 4)
			(_nw151).ArraySet1((func() _dafny.Sequence {
				if (_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool) {
					return _978_v30
				}
				return _dafny.Companion_Sequence_.Update(_978_v30, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_978_v30).Cardinality()))).Uint32(), _986_v38)
			})(), 5)
			(_nw151).ArraySet1(_978_v30, 6)
			(_nw151).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_978_v30, _978_v30), 7)
			(_nw151).ArraySet1(Companion_Default___.Fm28(p1, (_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool), globalState), 8)
			(_nw151).ArraySet1(_978_v30, 9)
			(_nw151).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hrwvirwn"), 10)
			(_nw151).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_978_v30, _978_v30), 11)
			(_nw151).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("silymvy"), 12)
			(_nw151).ArraySet1(_978_v30, 13)
			(_nw151).ArraySet1((func() _dafny.Sequence {
				if (_990_v42).Contains(!((_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool))) {
					return (_990_v42).Get(!((_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool))).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("jmvllgg")
			})(), 14)
			(_nw151).ArraySet1(_978_v30, 15)
			(_nw151).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-884))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg64 _dafny.Int) interface{} {
					return coer64(arg64)
				}
			}((func(_1004_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1005_i4 _dafny.Int) _dafny.CodePoint {
					return _1004_v38
				}
			})(_986_v38))), _978_v30), 16)
			(_nw151).ArraySet1(_978_v30, 17)
			_991_v43 = _nw151
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_991_v43), 0))
			_ = _index152
			(_991_v43).ArraySet1(_978_v30, (_index152).Int())
			var _1006_v44 _dafny.Map
			_ = _1006_v44
			_1006_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			var _1007_v45 D9
			_ = _1007_v45
			_1007_v45 = Companion_D9_.Create_DC24_((func() bool {
				if (_1006_v44).Contains(p0) {
					return (_1006_v44).Get(p0).(bool)
				}
				return p2
			})(), p0, _dafny.UnicodeSeqOfUtf8Bytes("i"))
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_991_v43), 0))
			_ = _index153
			(_991_v43).ArraySet1((_1007_v45).Dtor_cf43(), (_index153).Int())
			if (_975_v27).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_975_v27).Cardinality()))).Uint32()).(bool) {
				_986_v38 = Companion_Default___.Fm31(globalState)
				var _1008_v46 _dafny.Map
				_ = _1008_v46
				_1008_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(470), p1, p1, p0), (_this).F9())
				var _1009_v47 D7
				_ = _1009_v47
				_1009_v47 = Companion_D7_.Create_DC18_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_989_v41), p2)).Merge(_1008_v46))
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))
				_ = _index154
				var _rhs190 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(p1))
				_ = _rhs190
				var _rhs191 bool = p2
				_ = _rhs191
				var _rhs192 D7 = _1009_v47
				_ = _rhs192
				var _lhs123 _dafny.Array = _983_v35
				_ = _lhs123
				var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))
				_ = _lhs124
				r0 = _rhs190
				(_lhs123).ArraySet1(_rhs191, (_lhs124).Int())
				_1009_v47 = _rhs192
				r2 = Companion_Default___.Fm21(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_975_v27, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_975_v27).Cardinality()))).Uint32(), (_this).F9())).Cardinality()), Companion_Default___.Fm1(_982_v34, p1, globalState), globalState)
				var _1010_v48 D4
				_ = _1010_v48
				_1010_v48 = Companion_D4_.Create_DC14_(p2, p2, p1)
				r0 = (_1010_v48).Dtor_cf28()
				(globalState).F2 = false
			} else {
				var _1011_v49 _dafny.Array
				_ = _1011_v49
				var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw152
				_1011_v49 = _nw152
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_1011_v49), 0))
				_ = _index155
				(_1011_v49).ArraySet1(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_dafny.SeqOf((_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool))).Cardinality())), (_index155).Int())
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_1011_v49), 0))
				_ = _index156
				(_1011_v49).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-642), (_dafny.Zero).Minus(p0)), (_index156).Int())
				_978_v30 = _978_v30
				var _1012_v50 *C2
				_ = _1012_v50
				var _nw153 *C2 = New_C2_()
				_ = _nw153
				_nw153.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_978_v30, _978_v30)).Cardinality()), (_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool))
				_1012_v50 = _nw153
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_991_v43), 0))
				_ = _index157
				(_991_v43).ArraySet1(_978_v30, (_index157).Int())
				var _1013_v51 _dafny.MultiSet
				_ = _1013_v51
				_1013_v51 = _dafny.MultiSetOf(p2)
				var _1014_v52 D1
				_ = _1014_v52
				_1014_v52 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(385), (_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool), p2, (_983_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_983_v35), 0))).Int()).(bool), p2)
				r2 = (_1013_v51).Update((_1014_v52).Dtor_cf12(), Companion_Default___.Abs(_dafny.IntOfUint32((_975_v27).Cardinality())))
			}
		}
		var _source11 D3 = Companion_D3_.Create_DC10_()
		_ = _source11
		if _source11.Is_DC8() {
			var _1015_v53 _dafny.Sequence
			_ = _1015_v53
			_1015_v53 = _dafny.UnicodeSeqOfUtf8Bytes("ujap")
			r0 = _dafny.IntOfUint32((_1015_v53).Cardinality())
			var _1016_v54 _dafny.Array
			_ = _1016_v54
			var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw154
			_1016_v54 = _nw154
			var _1017_v55 _dafny.Array
			_ = _1017_v55
			_1017_v55 = _1016_v54
			_1016_v54 = (_1017_v55)
			var _source12 D1 = Companion_D1_.Create_DC3_(p2, _1015_v53, p1)
			_ = _source12
			if _source12.Is_DC2() {
				var _1018___mcc_h3 _dafny.Int = _source12.Get_().(D1_DC2).Cf2
				_ = _1018___mcc_h3
				var _1019___mcc_h4 bool = _source12.Get_().(D1_DC2).Cf3
				_ = _1019___mcc_h4
				var _1020___mcc_h5 _dafny.Int = _source12.Get_().(D1_DC2).Cf4
				_ = _1020___mcc_h5
				var _1021___mcc_h6 _dafny.Int = _source12.Get_().(D1_DC2).Cf5
				_ = _1021___mcc_h6
				var _1022_cf5 _dafny.Int = _1021___mcc_h6
				_ = _1022_cf5
				var _1023_cf4 _dafny.Int = _1020___mcc_h5
				_ = _1023_cf4
				var _1024_cf3 bool = _1019___mcc_h4
				_ = _1024_cf3
				var _1025_cf2 _dafny.Int = _1018___mcc_h3
				_ = _1025_cf2
				_1023_cf4 = (_1022_cf5).Times(_1022_cf5)
				var _1026_v56 _dafny.Map
				_ = _1026_v56
				_1026_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1024_cf3)
				var _1027_v57 _dafny.MultiSet
				_ = _1027_v57
				_1027_v57 = _dafny.MultiSetOf(p0, (_1026_v56).Cardinality())
				var _1028_v58 _dafny.CodePoint
				_ = _1028_v58
				_1028_v58 = _dafny.CodePoint('h')
				var _1029_v59 D7
				_ = _1029_v59
				_1029_v59 = Companion_D7_.Create_DC20_(false, _1028_v58)
				var _1030_v60 D4
				_ = _1030_v60
				_1030_v60 = Companion_D4_.Create_DC15_(Companion_D4_.Create_DC12_(Companion_Default___.Fm1(_1027_v57, _1023_cf4, globalState), false, _1025_cf2, _1025_cf2, (_1029_v59).Dtor_cf36()))
				_1030_v60 = _1030_v60
				var _1031_v61 _dafny.Array
				_ = _1031_v61
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_28
				var _nw155 _dafny.Array
				_ = _nw155
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw155 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) bool = (func(_1032_v53 _dafny.Sequence, _1033_p0 _dafny.Int, _1034_v58 _dafny.CodePoint, _1035_p1 _dafny.Int) func(_dafny.Int) bool {
						return func(_1036_i5 _dafny.Int) bool {
							return !((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1032_v53, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1033_p0), _dafny.IntOfUint32((_1032_v53).Cardinality()))).Uint32(), _1034_v58)).Cardinality())).Cmp(_1035_p1) < 0)
						}
					})(_1015_v53, p0, _1028_v58, p1)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw155 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw155).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw155).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1031_v61 = _nw155
				_1031_v61 = _1031_v61
				var _1037_v62 *C2
				_ = _1037_v62
				var _nw156 *C2 = New_C2_()
				_ = _nw156
				_nw156.Ctor__(_1025_cf2, p2)
				_1037_v62 = _nw156
				var _1038_v63 _dafny.Map
				_ = _1038_v63
				_1038_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1037_v62, _1024_cf3)
				var _1039_v64 _dafny.Array
				_ = _1039_v64
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
				_ = _nw157
				_1039_v64 = _nw157
				var _1040_v65 T2
				_ = _1040_v65
				var _nw158 *C3 = New_C3_()
				_ = _nw158
				_nw158.Ctor__((func() bool {
					if (_1038_v63).Contains(_1037_v62) {
						return (_1038_v63).Get(_1037_v62).(bool)
					}
					return !(false)
				})(), _1039_v64, _1024_cf3)
				_1040_v65 = _nw158
				var _1041_v66 _dafny.Sequence
				_ = _1041_v66
				_1041_v66 = _dafny.SeqOf(_1040_v65, _1040_v65)
				r3 = !(Companion_Default___.Fm1(_1027_v57, _dafny.IntOfUint32((_1041_v66).Cardinality()), globalState))
			} else if _source12.Is_DC3() {
				var _1042___mcc_h7 bool = _source12.Get_().(D1_DC3).Cf6
				_ = _1042___mcc_h7
				var _1043___mcc_h8 _dafny.Sequence = _source12.Get_().(D1_DC3).Cf7
				_ = _1043___mcc_h8
				var _1044___mcc_h9 _dafny.Int = _source12.Get_().(D1_DC3).Cf8
				_ = _1044___mcc_h9
				var _1045_cf8 _dafny.Int = _1044___mcc_h9
				_ = _1045_cf8
				var _1046_cf7 _dafny.Sequence = _1043___mcc_h8
				_ = _1046_cf7
				var _1047_cf6 bool = _1042___mcc_h7
				_ = _1047_cf6
				(globalState).F2 = _1047_cf6
				var _1048_v67 _dafny.Array
				_ = _1048_v67
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_29
				var _nw159 _dafny.Array
				_ = _nw159
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw159 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) bool = (func(_1049_cf6 bool, _1050_p2 bool) func(_dafny.Int) bool {
						return func(_1051_i6 _dafny.Int) bool {
							return (func() bool {
								if !(_1049_cf6) {
									return _1049_cf6
								}
								return _1050_p2
							})()
						}
					})(_1047_cf6, p2)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw159 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw159).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw159).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1048_v67 = _nw159
				var _1052_v68 _dafny.Sequence
				_ = _1052_v68
				_1052_v68 = _dafny.SeqOf(p1, _1045_cf8)
				var _1053_v69 _dafny.Map
				_ = _1053_v69
				_1053_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1052_v68, p0)
				var _1054_v71 *C0
				_ = _1054_v71
				var _nw160 *C0 = New_C0_()
				_ = _nw160
				_nw160.Ctor__(func() _dafny.Set {
					var _coll23 = _dafny.NewBuilder()
					_ = _coll23
					for _iter27 := _dafny.Iterate((_1053_v69).Keys().Elements()); ; {
						_compr_23, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _1055_v70 _dafny.Sequence
						_1055_v70 = interface{}(_compr_23).(_dafny.Sequence)
						if (_1053_v69).Contains(_1055_v70) {
							_coll23.Add(_1055_v70)
						}
					}
					return _coll23.ToSet()
				}(), _1045_cf8)
				_1054_v71 = _nw160
				var _1056_v72 _dafny.MultiSet
				_ = _1056_v72
				_1056_v72 = _dafny.MultiSetOf(_1054_v71, _1054_v71, _1054_v71)
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1048_v67), 0))
				_ = _index158
				(_1048_v67).ArraySet1(Companion_Default___.Fm1(_dafny.MultiSetOf(p1, p0), (_1056_v72).Cardinality(), globalState), (_index158).Int())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1048_v67), 0))
				_ = _index159
				(_1048_v67).ArraySet1(_1047_cf6, (_index159).Int())
				var _1057_v73 _dafny.Map
				_ = _1057_v73
				_1057_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1054_v71.F12, (_1054_v71.F12).Plus((_dafny.Zero).Minus((_dafny.Zero).Minus(p0))))
				_1057_v73 = (_1057_v73).Update((_dafny.Zero).Minus(p0), Companion_Default___.Fm0(globalState))
				var _1058_v74 *C2
				_ = _1058_v74
				var _nw161 *C2 = New_C2_()
				_ = _nw161
				_nw161.Ctor__(_1045_cf8, _1047_cf6)
				_1058_v74 = _nw161
				var _1059_v75 _dafny.Sequence
				_ = _1059_v75
				_1059_v75 = _dafny.SeqOf(_1058_v74, _1058_v74)
				var _1060_v76 _dafny.Map
				_ = _1060_v76
				_1060_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1059_v75, _1048_v67)
				var _1061_v77 _dafny.Array
				_ = _1061_v77
				var _nwElement0_32 _dafny.Array = (func() _dafny.Array {
					if (_1060_v76).Contains(_1059_v75) {
						return (_1060_v76).Get(_1059_v75).(_dafny.Array)
					}
					return _1048_v67
				})()
				_ = _nwElement0_32
				var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(4))
				_ = _nw162
				(_nw162).ArraySet1(_nwElement0_32, 0)
				(_nw162).ArraySet1(_1048_v67, 1)
				(_nw162).ArraySet1(_1048_v67, 2)
				(_nw162).ArraySet1(_1048_v67, 3)
				_1061_v77 = _nw162
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1061_v77), 0))
				_ = _index160
				(_1061_v77).ArraySet1(_1048_v67, (_index160).Int())
				var _1062_v78 _dafny.Map
				_ = _1062_v78
				_1062_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1015_v53).Cardinality()), _1048_v67)
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1061_v77), 0))
				_ = _index161
				(_1061_v77).ArraySet1((func() _dafny.Array {
					if (_1062_v78).Contains(_1054_v71.F12) {
						return (_1062_v78).Get(_1054_v71.F12).(_dafny.Array)
					}
					return _1048_v67
				})(), (_index161).Int())
			} else if _source12.Is_DC4() {
				var _1063___mcc_h10 _dafny.Int = _source12.Get_().(D1_DC4).Cf9
				_ = _1063___mcc_h10
				var _1064___mcc_h11 bool = _source12.Get_().(D1_DC4).Cf10
				_ = _1064___mcc_h11
				var _1065___mcc_h12 bool = _source12.Get_().(D1_DC4).Cf11
				_ = _1065___mcc_h12
				var _1066___mcc_h13 bool = _source12.Get_().(D1_DC4).Cf12
				_ = _1066___mcc_h13
				var _1067___mcc_h14 bool = _source12.Get_().(D1_DC4).Cf13
				_ = _1067___mcc_h14
				var _1068_cf13 bool = _1067___mcc_h14
				_ = _1068_cf13
				var _1069_cf12 bool = _1066___mcc_h13
				_ = _1069_cf12
				var _1070_cf11 bool = _1065___mcc_h12
				_ = _1070_cf11
				var _1071_cf10 bool = _1064___mcc_h11
				_ = _1071_cf10
				var _1072_cf9 _dafny.Int = _1063___mcc_h10
				_ = _1072_cf9
				var _1073_v79 _dafny.CodePoint
				_ = _1073_v79
				_1073_v79 = _dafny.CodePoint('d')
				_1073_v79 = _1073_v79
				var _1074_v80 _dafny.Array
				_ = _1074_v80
				var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
				_ = _nw163
				_1074_v80 = _nw163
				_1074_v80 = _1074_v80
				(globalState).F2 = true
				(globalState).F2 = (true) && ((_this).F9())
			} else if _source12.Is_DC1() {
				var _1075___mcc_h15 _dafny.Int = _source12.Get_().(D1_DC1).Cf1
				_ = _1075___mcc_h15
				var _1076_cf1 _dafny.Int = _1075___mcc_h15
				_ = _1076_cf1
				_1076_cf1 = Companion_Default___.SafeModuloInt(p1, p0)
				var _1077_v81 _dafny.CodePoint
				_ = _1077_v81
				_1077_v81 = _dafny.CodePoint('x')
				_1015_v53 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1015_v53, _1015_v53), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jor"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1015_v53).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jor")).Cardinality()))).Uint32(), _1077_v81), (Companion_Default___.SafeIndex(_1076_cf1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jor"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1015_v53).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jor")).Cardinality()))).Uint32(), _1077_v81)).Cardinality()))).Uint32(), Companion_Default___.Fm31(globalState)))
				r3 = true
				r3 = p2
			} else {
				var _1078___mcc_h16 D1 = _source12.Get_().(D1_DC5).Cf14
				_ = _1078___mcc_h16
				var _1079_cf14 D1 = _1078___mcc_h16
				_ = _1079_cf14
				var _1080_v82 _dafny.Map
				_ = _1080_v82
				_1080_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), p2)
				r0 = (p0).Times((p1).Minus((_1080_v82).Cardinality()))
				var _1081_v83 _dafny.Array
				_ = _1081_v83
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_30
				var _nw164 _dafny.Array
				_ = _nw164
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw164 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Sequence = func(_1082_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("jmokwm")
					}
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw164 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw164).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw164).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1081_v83 = _nw164
				_1081_v83 = _1081_v83
				r0 = Companion_Default___.SafeDivisionInt(p1, ((_1080_v82).Merge(_1080_v82)).Cardinality())
				var _1083_v84 D1
				_ = _1083_v84
				_1083_v84 = Companion_D1_.Create_DC2_(p1, true, p1, Companion_Default___.Fm0(globalState))
				var _1084_v85 _dafny.Sequence
				_ = _1084_v85
				_1084_v85 = _dafny.SeqOf((_this).F9())
				var _1085_v86 *C1
				_ = _1085_v86
				var _nw165 *C1 = New_C1_()
				_ = _nw165
				_nw165.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1083_v84, _1083_v84, _1083_v84, _1083_v84, Companion_D1_.Create_DC2_(_dafny.IntOfInt64(-578), (_this).F9(), _dafny.IntOfUint32((_1084_v85).Cardinality()), p0)), _dafny.SeqOf(_1083_v84, _1083_v84, _1083_v84, _1083_v84, Companion_D1_.Create_DC2_(Companion_Default___.Fm0(globalState), p2, p0, p0))))
				_1085_v86 = _nw165
			}
			var _1086_v87 _dafny.Map
			_ = _1086_v87
			_1086_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1016_v54, (_this).F9())
			_1086_v87 = (_1086_v87).Update(_1016_v54, !((_this).F9()))
		} else if _source11.Is_DC9() {
			var _1087___mcc_h0 bool = _source11.Get_().(D3_DC9).Cf17
			_ = _1087___mcc_h0
			var _1088___mcc_h1 _dafny.Sequence = _source11.Get_().(D3_DC9).Cf18
			_ = _1088___mcc_h1
			var _1089_cf18 _dafny.Sequence = _1088___mcc_h1
			_ = _1089_cf18
			var _1090_cf17 bool = _1087___mcc_h0
			_ = _1090_cf17
			var _1091_v88 _dafny.Array
			_ = _1091_v88
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_31
			var _nw166 _dafny.Array
			_ = _nw166
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw166 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.Int = (func(_1092_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1093_i8 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1093_i8, _1092_p1)
					}
				})(p1)
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw166 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw166).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw166).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1091_v88 = _nw166
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1091_v88), 0))
			_ = _index162
			(_1091_v88).ArraySet1(_dafny.IntOfUint32((_1089_cf18).Cardinality()), (_index162).Int())
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1091_v88), 0))
			_ = _index163
			(_1091_v88).ArraySet1(p1, (_index163).Int())
			var _1094_v89 _dafny.CodePoint
			_ = _1094_v89
			_1094_v89 = _dafny.CodePoint('v')
			var _1095_v90 _dafny.Map
			_ = _1095_v90
			_1095_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1094_v89, _1094_v89)
			var _1096_v91 _dafny.Sequence
			_ = _1096_v91
			_1096_v91 = _dafny.SeqOf(_1094_v89, _1094_v89)
			_1095_v90 = ((_1095_v90).Merge(_1095_v90)).Merge((_1095_v90).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), (_1096_v91).Select((Companion_Default___.SafeIndex((_1091_v88).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1091_v88), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1096_v91).Cardinality()))).Uint32()).(_dafny.CodePoint))))
			var _1097_v92 D10
			_ = _1097_v92
			_1097_v92 = Companion_D10_.Create_DC28_()
			var _1098_v93 _dafny.Map
			_ = _1098_v93
			_1098_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
			var _1099_v94 _dafny.Map
			_ = _1099_v94
			_1099_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1097_v92, _1098_v93)
			var _1100_v95 _dafny.MultiSet
			_ = _1100_v95
			_1100_v95 = _dafny.MultiSetOf((_1099_v94).Cardinality())
			var _rhs193 _dafny.Int = (_dafny.Zero).Minus(((_1091_v88).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1091_v88), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus(p0)))
			_ = _rhs193
			var _rhs194 bool = Companion_Default___.Fm1(_1100_v95, (_1091_v88).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1091_v88), 0))).Int()).(_dafny.Int), globalState)
			_ = _rhs194
			var _rhs195 _dafny.Int = p1
			_ = _rhs195
			r0 = _rhs193
			_1090_cf17 = _rhs194
			r0 = _rhs195
			var _1101_v96 D1
			_ = _1101_v96
			_1101_v96 = Companion_D1_.Create_DC2_(p0, false, p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("chnn")).Cardinality()))
			var _1102_v97 _dafny.Sequence
			_ = _1102_v97
			_1102_v97 = _dafny.SeqOf(_1101_v96)
			var _1103_v98 *C1
			_ = _1103_v98
			var _nw167 *C1 = New_C1_()
			_ = _nw167
			_nw167.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1102_v97, _1102_v97))
			_1103_v98 = _nw167
			var _1104_v99 *C2
			_ = _1104_v99
			var _nw168 *C2 = New_C2_()
			_ = _nw168
			_nw168.Ctor__((_1091_v88).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1091_v88), 0))).Int()).(_dafny.Int), !(p2))
			_1104_v99 = _nw168
			var _1105_v100 D4
			_ = _1105_v100
			_1105_v100 = Companion_D4_.Create_DC13_((_1104_v99).F14())
			var _1106_v101 _dafny.Set
			_ = _1106_v101
			_1106_v101 = _dafny.SetOf(_1090_cf17, p2)
			var _pat_let_tv61 = _1100_v95
			_ = _pat_let_tv61
			var _pat_let_tv62 = _1091_v88
			_ = _pat_let_tv62
			var _pat_let_tv63 = _1091_v88
			_ = _pat_let_tv63
			var _pat_let_tv64 = p1
			_ = _pat_let_tv64
			var _1107_v102 _dafny.Array
			_ = _1107_v102
			var _nwElement0_33 D4 = _1105_v100
			_ = _nwElement0_33
			var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(25))
			_ = _nw169
			(_nw169).ArraySet1(_nwElement0_33, 0)
			(_nw169).ArraySet1(_1105_v100, 1)
			(_nw169).ArraySet1(_1105_v100, 2)
			(_nw169).ArraySet1(_1105_v100, 3)
			(_nw169).ArraySet1(_1105_v100, 4)
			(_nw169).ArraySet1(_1105_v100, 5)
			(_nw169).ArraySet1(_1105_v100, 6)
			(_nw169).ArraySet1(func(_pat_let42_0 D4) D4 {
				return func(_1108_dt__update__tmp_h7 D4) D4 {
					return func(_pat_let43_0 _dafny.Int) D4 {
						return func(_1109_dt__update_hcf25_h0 _dafny.Int) D4 {
							return Companion_D4_.Create_DC13_(_1109_dt__update_hcf25_h0)
						}(_pat_let43_0)
					}((_pat_let_tv61).Cardinality())
				}(_pat_let42_0)
			}(_1105_v100), 7)
			(_nw169).ArraySet1(_1105_v100, 8)
			(_nw169).ArraySet1(_1105_v100, 9)
			(_nw169).ArraySet1(func(_pat_let44_0 D4) D4 {
				return func(_1110_dt__update__tmp_h8 D4) D4 {
					return func(_pat_let45_0 _dafny.Int) D4 {
						return func(_1111_dt__update_hcf25_h1 _dafny.Int) D4 {
							return Companion_D4_.Create_DC13_(_1111_dt__update_hcf25_h1)
						}(_pat_let45_0)
					}((_pat_let_tv63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_pat_let_tv62), 0))).Int()).(_dafny.Int))
				}(_pat_let44_0)
			}(_1105_v100), 10)
			(_nw169).ArraySet1(_1105_v100, 11)
			(_nw169).ArraySet1(_1105_v100, 12)
			(_nw169).ArraySet1(_1105_v100, 13)
			(_nw169).ArraySet1(_1105_v100, 14)
			(_nw169).ArraySet1(func(_pat_let46_0 D4) D4 {
				return func(_1112_dt__update__tmp_h9 D4) D4 {
					return func(_pat_let47_0 _dafny.Int) D4 {
						return func(_1113_dt__update_hcf25_h2 _dafny.Int) D4 {
							return Companion_D4_.Create_DC13_(_1113_dt__update_hcf25_h2)
						}(_pat_let47_0)
					}(_pat_let_tv64)
				}(_pat_let46_0)
			}(_1105_v100), 15)
			(_nw169).ArraySet1(_1105_v100, 16)
			(_nw169).ArraySet1(_1105_v100, 17)
			(_nw169).ArraySet1(_1105_v100, 18)
			(_nw169).ArraySet1(_1105_v100, 19)
			(_nw169).ArraySet1(Companion_D4_.Create_DC13_((_dafny.Zero).Minus((_1106_v101).Cardinality())), 20)
			(_nw169).ArraySet1(_1105_v100, 21)
			(_nw169).ArraySet1(Companion_D4_.Create_DC13_(p0), 22)
			(_nw169).ArraySet1(_1105_v100, 23)
			(_nw169).ArraySet1(_1105_v100, 24)
			_1107_v102 = _nw169
			var _1114_v103 _dafny.Array
			_ = _1114_v103
			var _nwElement0_34 _dafny.Array = _1107_v102
			_ = _nwElement0_34
			var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(13))
			_ = _nw170
			(_nw170).ArraySet1(_nwElement0_34, 0)
			(_nw170).ArraySet1(_1107_v102, 1)
			(_nw170).ArraySet1(_1107_v102, 2)
			(_nw170).ArraySet1(_1107_v102, 3)
			(_nw170).ArraySet1(_1107_v102, 4)
			(_nw170).ArraySet1(_1107_v102, 5)
			(_nw170).ArraySet1(_1107_v102, 6)
			(_nw170).ArraySet1(_1107_v102, 7)
			(_nw170).ArraySet1(_1107_v102, 8)
			(_nw170).ArraySet1(_1107_v102, 9)
			(_nw170).ArraySet1(_1107_v102, 10)
			(_nw170).ArraySet1(_1107_v102, 11)
			(_nw170).ArraySet1(_1107_v102, 12)
			_1114_v103 = _nw170
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1091_v88), 0))
			_ = _index164
			var _rhs196 *C1 = _1103_v98
			_ = _rhs196
			var _rhs197 _dafny.Int = _dafny.IntOfInt64(416)
			_ = _rhs197
			var _rhs198 *C2 = _1104_v99
			_ = _rhs198
			var _rhs199 _dafny.Array = _1114_v103
			_ = _rhs199
			var _lhs125 _dafny.Array = _1091_v88
			_ = _lhs125
			var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1091_v88), 0))
			_ = _lhs126
			_1103_v98 = _rhs196
			(_lhs125).ArraySet1(_rhs197, (_lhs126).Int())
			_1104_v99 = _rhs198
			_1114_v103 = _rhs199
		} else if _source11.Is_DC10() {
			r3 = (_this).F9()
			var _1115_v104 _dafny.CodePoint
			_ = _1115_v104
			_1115_v104 = _dafny.CodePoint('k')
			var _1116_v105 _dafny.Set
			_ = _1116_v105
			_1116_v105 = _dafny.SetOf(_1115_v104, _1115_v104)
			var _1117_v106 _dafny.Sequence
			_ = _1117_v106
			_1117_v106 = _dafny.SeqOf(true)
			var _1118_v107 _dafny.Set
			_ = _1118_v107
			_1118_v107 = _dafny.SetOf(p0)
			var _1119_v108 _dafny.MultiSet
			_ = _1119_v108
			_1119_v108 = _dafny.MultiSetOf(p1)
			var _1120_v109 _dafny.Array
			_ = _1120_v109
			var _nwElement0_35 bool = (_this).F9()
			_ = _nwElement0_35
			var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(29))
			_ = _nw171
			(_nw171).ArraySet1(_nwElement0_35, 0)
			(_nw171).ArraySet1(p2, 1)
			(_nw171).ArraySet1((_this).F9(), 2)
			(_nw171).ArraySet1(p2, 3)
			(_nw171).ArraySet1((_1117_v106).Select((Companion_Default___.SafeIndex((_1118_v107).Cardinality(), _dafny.IntOfUint32((_1117_v106).Cardinality()))).Uint32()).(bool), 4)
			(_nw171).ArraySet1(p2, 5)
			(_nw171).ArraySet1((_this).F9(), 6)
			(_nw171).ArraySet1(false, 7)
			(_nw171).ArraySet1(p2, 8)
			(_nw171).ArraySet1((_this).F9(), 9)
			(_nw171).ArraySet1(true, 10)
			(_nw171).ArraySet1((_this).F9(), 11)
			(_nw171).ArraySet1((_this).F9(), 12)
			(_nw171).ArraySet1((_this).F9(), 13)
			(_nw171).ArraySet1(!(p2), 14)
			(_nw171).ArraySet1(Companion_Default___.Fm1(_1119_v108, p1, globalState), 15)
			(_nw171).ArraySet1((_this).F9(), 16)
			(_nw171).ArraySet1(p2, 17)
			(_nw171).ArraySet1((_this).F9(), 18)
			(_nw171).ArraySet1(true, 19)
			(_nw171).ArraySet1((_this).F9(), 20)
			(_nw171).ArraySet1(Companion_Default___.Fm1(_1119_v108, (_dafny.Zero).Minus(p1), globalState), 21)
			(_nw171).ArraySet1((_this).F9(), 22)
			(_nw171).ArraySet1((_this).F9(), 23)
			(_nw171).ArraySet1((_this).F9(), 24)
			(_nw171).ArraySet1((_this).F9(), 25)
			(_nw171).ArraySet1((_this).F9(), 26)
			(_nw171).ArraySet1(!(p2), 27)
			(_nw171).ArraySet1(!((_this).F9()), 28)
			_1120_v109 = _nw171
			var _1121_v110 _dafny.Map
			_ = _1121_v110
			_1121_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1116_v105).Equals(_1116_v105), _1120_v109)
			var _1122_v111 _dafny.Sequence
			_ = _1122_v111
			_1122_v111 = _dafny.SeqOf(_1121_v110)
			var _1123_v112 *C2
			_ = _1123_v112
			var _nw172 *C2 = New_C2_()
			_ = _nw172
			_nw172.Ctor__(_dafny.IntOfInt64(-876), p2)
			_1123_v112 = _nw172
			var _1124_v113 _dafny.Sequence
			_ = _1124_v113
			_1124_v113 = _dafny.SeqOf(_1123_v112)
			_1121_v110 = (((_1122_v111).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1124_v113).Cardinality()), _dafny.IntOfUint32((_1122_v111).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1121_v110)).Update(Companion_Default___.Fm1(_dafny.MultiSetOf(p1, p0), (_dafny.SetOf((_this).F9())).Cardinality(), globalState), _1120_v109)
			r0 = (_1123_v112).F14()
			r0 = (_dafny.Zero).Minus((p1).Plus(Companion_Default___.Fm0(globalState)))
		} else {
			var _1125___mcc_h2 _dafny.MultiSet = _source11.Get_().(D3_DC7).Cf16
			_ = _1125___mcc_h2
			var _1126_cf16 _dafny.MultiSet = _1125___mcc_h2
			_ = _1126_cf16
			var _1127_v114 D4
			_ = _1127_v114
			_1127_v114 = Companion_D4_.Create_DC13_((_dafny.Zero).Minus(p0))
			var _1128_v115 _dafny.CodePoint
			_ = _1128_v115
			_1128_v115 = _dafny.CodePoint('y')
			var _1129_v116 *C1
			_ = _1129_v116
			var _nw173 *C1 = New_C1_()
			_ = _nw173
			_nw173.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(670))).Uint32(), func(coer65 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}((func(_1130_p0 _dafny.Int, _1131_p2 bool) func(_dafny.Int) D1 {
				return func(_1132_i9 _dafny.Int) D1 {
					return Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1130_p0)).Cardinality(), _1131_p2, _1130_p0, _1130_p0)
				}
			})(p0, p2))))
			_1129_v116 = _nw173
			var _1133_v117 _dafny.MultiSet
			_ = _1133_v117
			_1133_v117 = _dafny.MultiSetOf((Companion_D13_.Create_DC35_(p2, _1128_v115, _1129_v116, (_this).F9())).Dtor_cf61(), !(false))
			var _1134_v118 _dafny.Map
			_ = _1134_v118
			_1134_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1128_v115, (func() _dafny.Int {
				if (_1133_v117).Contains((_this).F9()) {
					return (_1133_v117).Multiplicity((_this).F9())
				}
				return p0
			})())
			var _1135_v119 _dafny.Map
			_ = _1135_v119
			_1135_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1127_v114, (func() _dafny.Int {
				if (_1134_v118).Contains(_1128_v115) {
					return (_1134_v118).Get(_1128_v115).(_dafny.Int)
				}
				return p1
			})())
			_1135_v119 = (_1135_v119).Update(_1127_v114, p1)
			var _1136_v120 _dafny.Array
			_ = _1136_v120
			var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
			_ = _nw174
			_1136_v120 = _nw174
			var _1137_v121 _dafny.Map
			_ = _1137_v121
			_1137_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F9())
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_1136_v120), 0))
			_ = _index165
			(_1136_v120).ArraySet1((_1137_v121).Merge(_1137_v121), (_index165).Int())
			var _1138_v123 _dafny.MultiSet
			_ = _1138_v123
			_1138_v123 = _dafny.MultiSetOf(p1)
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_1136_v120), 0))
			_ = _index166
			(_1136_v120).ArraySet1((_1137_v121).Merge(func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter28 := _dafny.Iterate((_1138_v123).Elements()); ; {
					_compr_24, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _1139_v122 _dafny.Int
					_1139_v122 = interface{}(_compr_24).(_dafny.Int)
					if (_1138_v123).Contains(_1139_v122) {
						_coll24.Add(Companion_Default___.SafeModuloInt(_1139_v122, p0), false)
					}
				}
				return _coll24.ToMap()
			}()), (_index166).Int())
			var _1140_v124 _dafny.Sequence
			_ = _1140_v124
			_1140_v124 = _dafny.SeqOf((_this).F9())
			var _1141_v125 D4
			_ = _1141_v125
			_1141_v125 = Companion_D4_.Create_DC12_(p2, p2, _dafny.IntOfUint32((_1140_v124).Cardinality()), p1, (_this).F9())
			var _1142_v126 D4
			_ = _1142_v126
			_1142_v126 = Companion_D4_.Create_DC15_((func() D4 {
				if p2 {
					return _1141_v125
				}
				return _1141_v125
			})())
			var _source13 D4 = _1142_v126
			_ = _source13
			if _source13.Is_DC12() {
				var _1143___mcc_h17 bool = _source13.Get_().(D4_DC12).Cf20
				_ = _1143___mcc_h17
				var _1144___mcc_h18 bool = _source13.Get_().(D4_DC12).Cf21
				_ = _1144___mcc_h18
				var _1145___mcc_h19 _dafny.Int = _source13.Get_().(D4_DC12).Cf22
				_ = _1145___mcc_h19
				var _1146___mcc_h20 _dafny.Int = _source13.Get_().(D4_DC12).Cf23
				_ = _1146___mcc_h20
				var _1147___mcc_h21 bool = _source13.Get_().(D4_DC12).Cf24
				_ = _1147___mcc_h21
				var _1148_cf24 bool = _1147___mcc_h21
				_ = _1148_cf24
				var _1149_cf23 _dafny.Int = _1146___mcc_h20
				_ = _1149_cf23
				var _1150_cf22 _dafny.Int = _1145___mcc_h19
				_ = _1150_cf22
				var _1151_cf21 bool = _1144___mcc_h18
				_ = _1151_cf21
				var _1152_cf20 bool = _1143___mcc_h17
				_ = _1152_cf20
				var _1153_v127 _dafny.Array
				_ = _1153_v127
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_32
				var _nw175 _dafny.Array
				_ = _nw175
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw175 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.Int = (func(_1154_cf23 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1155_i10 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1155_i10, _1154_cf23)
						}
					})(_1149_cf23)
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw175 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw175).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw175).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1153_v127 = _nw175
				r0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1153_v127, _1153_v127)).Cardinality()
				_1153_v127 = _1153_v127
				var _1156_v128 *C1
				_ = _1156_v128
				var _nw176 *C1 = New_C1_()
				_ = _nw176
				_nw176.Ctor__(_1129_v116.F13)
				_1156_v128 = _nw176
				var _1157_v129 _dafny.Sequence
				_ = _1157_v129
				_1157_v129 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_cf22, p0)).Cardinality())
				var _1158_v130 _dafny.Set
				_ = _1158_v130
				_1158_v130 = _dafny.SetOf(p0)
				var _1159_v131 _dafny.Map
				_ = _1159_v131
				_1159_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1151_cf21, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("acs")).Cardinality()))
				var _1160_v132 _dafny.Array
				_ = _1160_v132
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_33
				var _nw177 _dafny.Array
				_ = _nw177
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw177 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1161_cf20 bool) func(_dafny.Int) bool {
						return func(_1162_i11 _dafny.Int) bool {
							return _1161_cf20
						}
					})(_1152_cf20)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw177 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw177).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw177).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1160_v132 = _nw177
				var _1163_v133 _dafny.Map
				_ = _1163_v133
				_1163_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1160_v132, _1150_cf22)
				var _1164_v134 _dafny.Array
				_ = _1164_v134
				var _nwElement0_36 _dafny.Set = _1158_v130
				_ = _nwElement0_36
				var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(19))
				_ = _nw178
				(_nw178).ArraySet1(_nwElement0_36, 0)
				(_nw178).ArraySet1(_1158_v130, 1)
				(_nw178).ArraySet1(_1158_v130, 2)
				(_nw178).ArraySet1(_1158_v130, 3)
				(_nw178).ArraySet1(_1158_v130, 4)
				(_nw178).ArraySet1(_1158_v130, 5)
				(_nw178).ArraySet1(_dafny.SetOf(p0), 6)
				(_nw178).ArraySet1(_1158_v130, 7)
				(_nw178).ArraySet1(_dafny.SetOf((_1159_v131).Cardinality()), 8)
				(_nw178).ArraySet1(_1158_v130, 9)
				(_nw178).ArraySet1(_1158_v130, 10)
				(_nw178).ArraySet1(_1158_v130, 11)
				(_nw178).ArraySet1(_1158_v130, 12)
				(_nw178).ArraySet1(_1158_v130, 13)
				(_nw178).ArraySet1(_dafny.SetOf((_1163_v133).Cardinality(), _1149_cf23, _dafny.IntOfInt64(748)), 14)
				(_nw178).ArraySet1(_1158_v130, 15)
				(_nw178).ArraySet1(_1158_v130, 16)
				(_nw178).ArraySet1(_1158_v130, 17)
				(_nw178).ArraySet1(_1158_v130, 18)
				_1164_v134 = _nw178
				var _1165_v135 T2
				_ = _1165_v135
				var _nw179 *C3 = New_C3_()
				_ = _nw179
				_nw179.Ctor__(_1151_cf21, _1164_v134, _1152_cf20)
				_1165_v135 = _nw179
				var _1166_v136 _dafny.Sequence
				_ = _1166_v136
				_1166_v136 = _dafny.SeqOf(_1165_v135, _1165_v135)
				var _1167_v137 _dafny.MultiSet
				_ = _1167_v137
				_1167_v137 = _dafny.MultiSetOf(_1128_v115)
				var _1168_v138 _dafny.Array
				_ = _1168_v138
				var _nwElement0_37 bool = (_dafny.IntOfInt64(62)).Cmp((_dafny.Zero).Minus(_1150_cf22)) != 0
				_ = _nwElement0_37
				var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(24))
				_ = _nw180
				(_nw180).ArraySet1(_nwElement0_37, 0)
				(_nw180).ArraySet1(!(Companion_Default___.Fm1(_1138_v123, _dafny.IntOfInt64(186), globalState)), 1)
				(_nw180).ArraySet1(_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm28((_1157_v129).Select((Companion_Default___.SafeIndex((_1138_v123).Cardinality(), _dafny.IntOfUint32((_1157_v129).Cardinality()))).Uint32()).(_dafny.Int), p2, globalState), _1128_v115), 2)
				(_nw180).ArraySet1((_dafny.IntOfUint32((_1166_v136).Cardinality())).Cmp(_dafny.IntOfInt64(-521)) > 0, 3)
				(_nw180).ArraySet1(!(!(_1152_cf20)), 4)
				(_nw180).ArraySet1(true, 5)
				(_nw180).ArraySet1((_1165_v135).F9(), 6)
				(_nw180).ArraySet1((_this).F9(), 7)
				(_nw180).ArraySet1((_1167_v137).IsSubsetOf(_1167_v137), 8)
				(_nw180).ArraySet1((_1165_v135).F9(), 9)
				(_nw180).ArraySet1(true, 10)
				(_nw180).ArraySet1(!(p2) || (!(_1152_cf20)), 11)
				(_nw180).ArraySet1((_this).F9(), 12)
				(_nw180).ArraySet1(_1148_cf24, 13)
				(_nw180).ArraySet1((_this).F9(), 14)
				(_nw180).ArraySet1(p2, 15)
				(_nw180).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf((_this).Fm2(_1149_cf23, (_1165_v135).F9(), !((func() bool {
					if (_1137_v121).Contains(p0) {
						return (_1137_v121).Get(p0).(bool)
					}
					return (_1165_v135).F9()
				})()), Companion_Default___.Fm0(globalState), globalState), _1157_v129), 16)
				(_nw180).ArraySet1(false, 17)
				(_nw180).ArraySet1(_1151_cf21, 18)
				(_nw180).ArraySet1(p2, 19)
				(_nw180).ArraySet1(_1151_cf21, 20)
				(_nw180).ArraySet1(_1152_cf20, 21)
				(_nw180).ArraySet1(_1152_cf20, 22)
				(_nw180).ArraySet1(_1152_cf20, 23)
				_1168_v138 = _nw180
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_1160_v132), 0))
				_ = _index167
				(_1160_v132).ArraySet1(true, (_index167).Int())
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1153_v127), 0))
				_ = _index168
				(_1153_v127).ArraySet1(_1150_cf22, (_index168).Int())
				var _1169_v139 _dafny.Sequence
				_ = _1169_v139
				_1169_v139 = _dafny.SeqOf(_1159_v131, _1159_v131, _1159_v131, _1159_v131)
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1153_v127), 0))
				_ = _index169
				(_1153_v127).ArraySet1(_dafny.IntOfUint32((_1169_v139).Cardinality()), (_index169).Int())
				var _1170_v140 _dafny.MultiSet
				_ = _1170_v140
				_1170_v140 = _dafny.MultiSetOf(_1153_v127, _1153_v127, _1153_v127)
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_1160_v132), 0))
				_ = _index170
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1153_v127), 0))
				_ = _index171
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1153_v127), 0))
				_ = _index172
				var _rhs200 bool = (_1170_v140).IsDisjointFrom(_1170_v140)
				_ = _rhs200
				var _rhs201 _dafny.Int = p1
				_ = _rhs201
				var _rhs202 _dafny.Int = (func() _dafny.Int {
					if (_1159_v131).Contains((_1133_v117).IsProperSubsetOf(_1133_v117)) {
						return (_1159_v131).Get((_1133_v117).IsProperSubsetOf(_1133_v117)).(_dafny.Int)
					}
					return _1150_cf22
				})()
				_ = _rhs202
				var _rhs203 _dafny.Int = (_dafny.Zero).Minus(_1149_cf23)
				_ = _rhs203
				var _rhs204 bool = true
				_ = _rhs204
				var _lhs127 _dafny.Array = _1160_v132
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_1160_v132), 0))
				_ = _lhs128
				var _lhs129 _dafny.Array = _1153_v127
				_ = _lhs129
				var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1153_v127), 0))
				_ = _lhs130
				var _lhs131 _dafny.Array = _1153_v127
				_ = _lhs131
				var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1153_v127), 0))
				_ = _lhs132
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				(_lhs127).ArraySet1(_rhs200, (_lhs128).Int())
				(_lhs129).ArraySet1(_rhs201, (_lhs130).Int())
				(_lhs131).ArraySet1(_rhs202, (_lhs132).Int())
				_1150_cf22 = _rhs203
				_lhs133.F2 = _rhs204
			} else if _source13.Is_DC13() {
				var _1171___mcc_h22 _dafny.Int = _source13.Get_().(D4_DC13).Cf25
				_ = _1171___mcc_h22
				var _1172_cf25 _dafny.Int = _1171___mcc_h22
				_ = _1172_cf25
				var _1173_v141 _dafny.Array
				_ = _1173_v141
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_34
				var _nw181 _dafny.Array
				_ = _nw181
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw181 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Int = (func(_1174_cf25 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1175_i12 _dafny.Int) _dafny.Int {
							return (_1175_i12).Plus(_1174_cf25)
						}
					})(_1172_cf25)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw181 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw181).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw181).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1173_v141 = _nw181
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_1173_v141), 0))
				_ = _index173
				(_1173_v141).ArraySet1((func() _dafny.Int {
					if (_this).F9() {
						return p1
					}
					return _1172_cf25
				})(), (_index173).Int())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_1173_v141), 0))
				_ = _index174
				(_1173_v141).ArraySet1(Companion_Default___.SafeModuloInt(_1172_cf25, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(34))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}((func(_1176_v115 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1177_i13 _dafny.Int) _dafny.CodePoint {
						return _1176_v115
					}
				})(_1128_v115)))).Cardinality())), (_index174).Int())
				var _1178_v142 D3
				_ = _1178_v142
				_1178_v142 = Companion_D3_.Create_DC9_(Companion_Default___.Fm1(_dafny.MultiSetOf(_dafny.IntOfInt64(527), _1172_cf25, p1, _1172_cf25), p0, globalState), _dafny.SeqOf(!(true)))
				var _1179_v143 _dafny.Map
				_ = _1179_v143
				_1179_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(p1)).Cmp((_dafny.Zero).Minus(p1)) < 0, (_1178_v142).Dtor_cf17())
				_1179_v143 = (_1179_v143).Update(p2, !(false) || (p2))
				var _1180_v144 _dafny.Array
				_ = _1180_v144
				var _nw182 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw182
				_1180_v144 = _nw182
				var _1181_v145 _dafny.Map
				_ = _1181_v145
				_1181_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1180_v144, (_this).F9())
				var _1182_v146 _dafny.Sequence
				_ = _1182_v146
				_1182_v146 = _dafny.SeqOf(_1180_v144, _1180_v144)
				var _1183_v147 _dafny.Sequence
				_ = _1183_v147
				_1183_v147 = _dafny.UnicodeSeqOfUtf8Bytes("n")
				var _1184_v148 _dafny.MultiSet
				_ = _1184_v148
				_1184_v148 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1172_cf25), (Companion_Default___.SafeIndex((_1173_v141).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_1173_v141), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(_1172_cf25)).Cardinality()))).Uint32(), _1172_cf25)), _1138_v123)
				_1181_v145 = (_1181_v145).Update((_1182_v146).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1183_v147).Cardinality()), _dafny.IntOfUint32((_1182_v146).Cardinality()))).Uint32()).(_dafny.Array), !(_1184_v148).Contains(_1138_v123))
				r0 = Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfInt64(84))
			} else if _source13.Is_DC14() {
				var _1185___mcc_h23 bool = _source13.Get_().(D4_DC14).Cf26
				_ = _1185___mcc_h23
				var _1186___mcc_h24 bool = _source13.Get_().(D4_DC14).Cf27
				_ = _1186___mcc_h24
				var _1187___mcc_h25 _dafny.Int = _source13.Get_().(D4_DC14).Cf28
				_ = _1187___mcc_h25
				var _1188_cf28 _dafny.Int = _1187___mcc_h25
				_ = _1188_cf28
				var _1189_cf27 bool = _1186___mcc_h24
				_ = _1189_cf27
				var _1190_cf26 bool = _1185___mcc_h23
				_ = _1190_cf26
				var _1191_v149 _dafny.Array
				_ = _1191_v149
				var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw183
				_1191_v149 = _nw183
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_1191_v149), 0))
				_ = _index175
				(_1191_v149).ArraySet1(p1, (_index175).Int())
				var _1192_v150 _dafny.Sequence
				_ = _1192_v150
				_1192_v150 = _dafny.SeqOf(_1138_v123)
				var _1193_v151 _dafny.Map
				_ = _1193_v151
				_1193_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1189_cf27, (_1192_v150).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1192_v150).Cardinality()))).Uint32()).(_dafny.MultiSet))
				var _1194_v152 _dafny.Sequence
				_ = _1194_v152
				_1194_v152 = _dafny.SeqOf(p1, _dafny.IntOfUint32((_1140_v124).Cardinality()))
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_1191_v149), 0))
				_ = _index176
				var _rhs205 _dafny.MultiSet = ((func() _dafny.MultiSet {
					if (_1193_v151).Contains(!((_this).F9())) {
						return (_1193_v151).Get(!((_this).F9())).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(_1194_v152)
				})()).Union(_1138_v123)
				_ = _rhs205
				var _rhs206 _dafny.Int = p1
				_ = _rhs206
				var _lhs134 _dafny.Array = _1191_v149
				_ = _lhs134
				var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_1191_v149), 0))
				_ = _lhs135
				_1138_v123 = _rhs205
				(_lhs134).ArraySet1(_rhs206, (_lhs135).Int())
				var _1195_v153 _dafny.MultiSet
				_ = _1195_v153
				_1195_v153 = _dafny.MultiSetFromSeq(_1194_v152)
				var _1196_v154 _dafny.MultiSet
				_ = _1196_v154
				_1196_v154 = _dafny.MultiSetOf(_1195_v153)
				_1188_cf28 = (func() _dafny.Int {
					if (_1196_v154).Contains(_1195_v153) {
						return (_1196_v154).Multiplicity(_1195_v153)
					}
					return p0
				})()
				var _1197_v155 _dafny.Sequence
				_ = _1197_v155
				_1197_v155 = _dafny.SeqOf(_1140_v124, _1140_v124, _1140_v124)
				var _1198_v156 _dafny.Sequence
				_ = _1198_v156
				_1198_v156 = _dafny.UnicodeSeqOfUtf8Bytes("ifap")
				var _1199_v157 _dafny.Array
				_ = _1199_v157
				var _nwElement0_38 _dafny.Sequence = _1140_v124
				_ = _nwElement0_38
				var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(12))
				_ = _nw184
				(_nw184).ArraySet1(_nwElement0_38, 0)
				(_nw184).ArraySet1((_1197_v155).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1197_v155).Cardinality()))).Uint32()).(_dafny.Sequence), 1)
				(_nw184).ArraySet1(_1140_v124, 2)
				(_nw184).ArraySet1(Companion_Default___.Fm20((_this).F9(), _1189_cf27, (_1191_v149).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_1191_v149), 0))).Int()).(_dafny.Int), _1198_v156, globalState), 3)
				(_nw184).ArraySet1(_1140_v124, 4)
				(_nw184).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1190_cf26), _1140_v124), 5)
				(_nw184).ArraySet1(_1140_v124, 6)
				(_nw184).ArraySet1(_1140_v124, 7)
				(_nw184).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1140_v124, _1140_v124), 8)
				(_nw184).ArraySet1(_dafny.Companion_Sequence_.Update(_1140_v124, (Companion_Default___.SafeIndex(_1188_cf28, _dafny.IntOfUint32((_1140_v124).Cardinality()))).Uint32(), true), 9)
				(_nw184).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1140_v124, _1140_v124), 10)
				(_nw184).ArraySet1(_dafny.Companion_Sequence_.Update(_1140_v124, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1140_v124).Cardinality()))).Uint32(), _1189_cf27), 11)
				_1199_v157 = _nw184
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_1199_v157), 0))
				_ = _index177
				(_1199_v157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1140_v124, _1140_v124), (_index177).Int())
				var _1200_v158 _dafny.Map
				_ = _1200_v158
				_1200_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1190_cf26, _dafny.IntOfUint32((Companion_Default___.Fm28(_dafny.IntOfInt64(284), Companion_Default___.Fm1(_1138_v123, p0, globalState), globalState)).Cardinality()))
				var _1201_v159 D10
				_ = _1201_v159
				_1201_v159 = Companion_D10_.Create_DC27_(_1200_v158)
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_1199_v157), 0))
				_ = _index178
				(_1199_v157).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1140_v124, (Companion_Default___.SafeIndex(_1188_cf28, _dafny.IntOfUint32((_1140_v124).Cardinality()))).Uint32(), p2), Companion_Default___.Fm41(_1201_v159, (_1191_v149).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_1191_v149), 0))).Int()).(_dafny.Int), _1194_v152, globalState)), (_index178).Int())
				_1190_cf26 = _1189_cf27
			} else if _source13.Is_DC11() {
				var _1202___mcc_h26 _dafny.Array = _source13.Get_().(D4_DC11).Cf19
				_ = _1202___mcc_h26
				var _1203_cf19 _dafny.Array = _1202___mcc_h26
				_ = _1203_cf19
				var _1204_v160 _dafny.Array
				_ = _1204_v160
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_35
				var _nw185 _dafny.Array
				_ = _nw185
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw185 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) bool = (func(_1205_p2 bool) func(_dafny.Int) bool {
						return func(_1206_i14 _dafny.Int) bool {
							return _1205_p2
						}
					})(p2)
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw185 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw185).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw185).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1204_v160 = _nw185
				var _1207_v161 _dafny.Map
				_ = _1207_v161
				_1207_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1204_v160)
				var _1208_v162 _dafny.Map
				_ = _1208_v162
				_1208_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1207_v161).Merge(_1207_v161), (_this).F9())
				var _1209_v163 D9
				_ = _1209_v163
				_1209_v163 = Companion_D9_.Create_DC24_(!(p2), p0, Companion_Default___.Fm28(p1, p2, globalState))
				_1208_v162 = (_1208_v162).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1204_v160)).Merge(_1207_v161), Companion_Default___.Fm1(_1138_v123, (_1209_v163).Dtor_cf42(), globalState))
				var _1210_v164 T1
				_ = _1210_v164
				var _nw186 *C2 = New_C2_()
				_ = _nw186
				_nw186.Ctor__(p1, !(p2))
				_1210_v164 = _nw186
				_1210_v164 = (func() T1 {
					if p2 {
						return _1210_v164
					}
					return _1210_v164
				})()
				var _1211_v165 _dafny.Sequence
				_ = _1211_v165
				_1211_v165 = _dafny.SeqOf(_1204_v160)
				(globalState).F2 = (func() bool {
					if !(_1133_v117).Equals(_1133_v117) {
						return (_this).F9()
					}
					return _dafny.Companion_Sequence_.IsProperPrefixOf(_1211_v165, _1211_v165)
				})()
				r0 = _dafny.IntOfUint32((_1140_v124).Cardinality())
			} else {
				var _1212___mcc_h27 D4 = _source13.Get_().(D4_DC15).Cf29
				_ = _1212___mcc_h27
				var _1213_cf29 D4 = _1212___mcc_h27
				_ = _1213_cf29
				r0 = p1
				var _1214_v166 _dafny.Set
				_ = _1214_v166
				_1214_v166 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("mgsxcsdpe"))
				_1214_v166 = _1214_v166
				var _1215_v167 _dafny.Array
				_ = _1215_v167
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_36
				var _nw187 _dafny.Array
				_ = _nw187
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw187 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = (func(_1216_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1217_i15 _dafny.Int) _dafny.Int {
							return (_1217_i15).Minus(_1216_p0)
						}
					})(p0)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw187 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw187).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw187).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1215_v167 = _nw187
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_1215_v167), 0))
				_ = _index179
				(_1215_v167).ArraySet1(_dafny.IntOfInt64(688), (_index179).Int())
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_1215_v167), 0))
				_ = _index180
				(_1215_v167).ArraySet1(p1, (_index180).Int())
				var _1218_v168 _dafny.Sequence
				_ = _1218_v168
				_1218_v168 = _dafny.UnicodeSeqOfUtf8Bytes("t")
				r0 = _dafny.IntOfUint32((_1218_v168).Cardinality())
			}
			if (_this).F9() {
				_1138_v123 = _1138_v123
				var _rhs207 _dafny.MultiSet = _1133_v117
				_ = _rhs207
				var _rhs208 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(p1))
				_ = _rhs208
				var _rhs209 _dafny.CodePoint = (Companion_D14_.Create_DC38_(_1128_v115, p2, _1140_v124)).Dtor_cf67()
				_ = _rhs209
				_1133_v117 = _rhs207
				r0 = _rhs208
				_1128_v115 = _rhs209
				var _1219_v169 _dafny.Sequence
				_ = _1219_v169
				_1219_v169 = _dafny.UnicodeSeqOfUtf8Bytes("xqbsdfpq")
				r3 = !_dafny.Companion_Sequence_.Contains(_1219_v169, _1128_v115)
				var _1220_v170 _dafny.Set
				_ = _1220_v170
				_1220_v170 = _dafny.SetOf(_1129_v116)
				(globalState).F2 = (func() bool {
					if !((_this).F9()) {
						return true
					}
					return (_1220_v170).Equals(_1220_v170)
				})()
				var _1221_v171 _dafny.Array
				_ = _1221_v171
				var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
				_ = _nw188
				_1221_v171 = _nw188
				var _1222_v172 _dafny.Array
				_ = _1222_v172
				var _nwElement0_39 bool = p2
				_ = _nwElement0_39
				var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(11))
				_ = _nw189
				(_nw189).ArraySet1(_nwElement0_39, 0)
				(_nw189).ArraySet1(p2, 1)
				(_nw189).ArraySet1(true, 2)
				(_nw189).ArraySet1(Companion_Default___.Fm1(_dafny.MultiSetOf(p1), p0, globalState), 3)
				(_nw189).ArraySet1(p2, 4)
				(_nw189).ArraySet1(true, 5)
				(_nw189).ArraySet1(p2, 6)
				(_nw189).ArraySet1(Companion_Default___.Fm1(_1138_v123, p0, globalState), 7)
				(_nw189).ArraySet1((_this).F9(), 8)
				(_nw189).ArraySet1(false, 9)
				(_nw189).ArraySet1(Companion_Default___.Fm1(_1138_v123, (_dafny.MultiSetOf((_this).F9())).Cardinality(), globalState), 10)
				_1222_v172 = _nw189
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_1221_v171), 0))
				_ = _index181
				(_1221_v171).ArraySet1(_dafny.SetOf(_1222_v172), (_index181).Int())
				var _1223_v173 _dafny.Set
				_ = _1223_v173
				_1223_v173 = _dafny.SetOf(_1222_v172, _1222_v172, _1222_v172)
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_1221_v171), 0))
				_ = _index182
				(_1221_v171).ArraySet1(_1223_v173, (_index182).Int())
			} else {
				(globalState).F2 = (_this).F9()
				var _1224_v174 _dafny.Sequence
				_ = _1224_v174
				_1224_v174 = _dafny.UnicodeSeqOfUtf8Bytes("ie")
				var _1225_v175 _dafny.Array
				_ = _1225_v175
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_37
				var _nw190 _dafny.Array
				_ = _nw190
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw190 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) _dafny.Map = func(_1226_i17 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
					}
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw190 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw190).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw190).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1225_v175 = _nw190
				var _1227_v176 D7
				_ = _1227_v176
				_1227_v176 = Companion_D7_.Create_DC19_(_1225_v175, (_this).F9(), _1224_v174)
				var _1228_v177 _dafny.Sequence
				_ = _1228_v177
				_1228_v177 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(153))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}(func(_1229_i16 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				})), _1224_v174, (_1227_v176).Dtor_cf35())
				var _rhs210 _dafny.Sequence = _1228_v177
				_ = _rhs210
				var _rhs211 _dafny.Int = p0
				_ = _rhs211
				var _rhs212 bool = _dafny.Companion_Sequence_.Contains(_1140_v124, (_this).F9())
				_ = _rhs212
				var _rhs213 _dafny.CodePoint = _1128_v115
				_ = _rhs213
				_1228_v177 = _rhs210
				r0 = _rhs211
				r3 = _rhs212
				_1128_v115 = _rhs213
				var _1230_v178 _dafny.Array
				_ = _1230_v178
				var _nw191 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(24))
				_ = _nw191
				_1230_v178 = _nw191
				var _pat_let_tv65 = _1140_v124
				_ = _pat_let_tv65
				var _pat_let_tv66 = p1
				_ = _pat_let_tv66
				var _pat_let_tv67 = _1140_v124
				_ = _pat_let_tv67
				var _pat_let_tv68 = p2
				_ = _pat_let_tv68
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_1230_v178), 0))
				_ = _index183
				(_1230_v178).ArraySet1(func(_pat_let48_0 D3) D3 {
					return func(_1231_dt__update__tmp_h10 D3) D3 {
						return func(_pat_let49_0 _dafny.Sequence) D3 {
							return func(_1232_dt__update_hcf18_h0 _dafny.Sequence) D3 {
								return Companion_D3_.Create_DC9_((_1231_dt__update__tmp_h10).Dtor_cf17(), _1232_dt__update_hcf18_h0)
							}(_pat_let49_0)
						}(_dafny.Companion_Sequence_.Update(_pat_let_tv65, (Companion_Default___.SafeIndex(_pat_let_tv66, _dafny.IntOfUint32((_pat_let_tv67).Cardinality()))).Uint32(), _pat_let_tv68))
					}(_pat_let48_0)
				}(Companion_D3_.Create_DC9_((_this).F9(), _1140_v124)), (_index183).Int())
				var _1233_v179 D3
				_ = _1233_v179
				_1233_v179 = Companion_D3_.Create_DC9_((_dafny.IntOfInt64(657)).Cmp(Companion_Default___.Fm0(globalState)) > 0, _dafny.SeqOf((_this).F9(), p2))
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_1230_v178), 0))
				_ = _index184
				(_1230_v178).ArraySet1(_1233_v179, (_index184).Int())
				var _1234_v180 _dafny.MultiSet
				_ = _1234_v180
				_1234_v180 = _dafny.MultiSetOf(_1128_v115)
				r0 = (func() _dafny.Set {
					var _coll25 = _dafny.NewBuilder()
					_ = _coll25
					for _iter29 := _dafny.Iterate(((_1234_v180).Union(_1234_v180)).Elements()); ; {
						_compr_25, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _1235_v181 _dafny.CodePoint
						_1235_v181 = interface{}(_compr_25).(_dafny.CodePoint)
						if ((_1234_v180).Union(_1234_v180)).Contains(_1235_v181) {
							_coll25.Add(_1235_v181)
						}
					}
					return _coll25.ToSet()
				}()).Cardinality()
				var _1236_v182 _dafny.Array
				_ = _1236_v182
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_38
				var _nw192 _dafny.Array
				_ = _nw192
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw192 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Set = (func(_1237_v115 _dafny.CodePoint, _1238_v117 _dafny.MultiSet) func(_dafny.Int) _dafny.Set {
						return func(_1239_i18 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(846))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg68 _dafny.Int) interface{} {
									return coer68(arg68)
								}
							}((func(_1240_v115 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1241_i19 _dafny.Int) _dafny.CodePoint {
									return _1240_v115
								}
							})(_1237_v115)))).Cardinality()), (_1238_v117).Cardinality())
						}
					})(_1128_v115, _1133_v117)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw192 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw192).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw192).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1236_v182 = _nw192
				var _1242_v183 _dafny.Sequence
				_ = _1242_v183
				_1242_v183 = _dafny.SeqOf(_1236_v182, _1236_v182)
				var _1243_v184 _dafny.Sequence
				_ = _1243_v184
				_1243_v184 = _dafny.SeqOf((_1242_v183).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_1242_v183).Cardinality()))).Uint32()).(_dafny.Array))
				var _1244_v185 *C3
				_ = _1244_v185
				var _nw193 *C3 = New_C3_()
				_ = _nw193
				_nw193.Ctor__(true, (_1243_v184).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1243_v184).Cardinality()))).Uint32()).(_dafny.Array), p2)
				_1244_v185 = _nw193
				var _1245_v186 _dafny.Map
				_ = _1245_v186
				_1245_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1244_v185)
				var _1246_v187 _dafny.Sequence
				_ = _1246_v187
				_1246_v187 = _dafny.SeqOf(p0, p1, (_dafny.Zero).Minus((_1245_v186).Cardinality()))
				_1246_v187 = _1246_v187
			}
		}
		var _1247_v188 _dafny.Sequence
		_ = _1247_v188
		_1247_v188 = _dafny.UnicodeSeqOfUtf8Bytes("canqxbdox")
		var _1248_v189 D1
		_ = _1248_v189
		_1248_v189 = Companion_D1_.Create_DC2_(p1, true, _dafny.IntOfUint32((_1247_v188).Cardinality()), p0)
		var _1249_v190 *C1
		_ = _1249_v190
		var _nw194 *C1 = New_C1_()
		_ = _nw194
		_nw194.Ctor__(_dafny.SeqOf(_1248_v189))
		_1249_v190 = _nw194
		var _1250_v191 _dafny.Set
		_ = _1250_v191
		_1250_v191 = _dafny.SetOf(_1249_v190, _1249_v190)
		var _hi3 _dafny.Int = (Companion_Default___.Fm0(globalState)).Minus(p0)
		_ = _hi3
		for _1251_i20 := (_1250_v191).Cardinality(); _1251_i20.Cmp(_hi3) < 0; _1251_i20 = _1251_i20.Plus(_dafny.One) {
			var _1252_v192 *C1
			_ = _1252_v192
			var _nw195 *C1 = New_C1_()
			_ = _nw195
			_nw195.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1249_v190.F13, _1249_v190.F13))
			_1252_v192 = _nw195
			var _1253_v193 _dafny.Sequence
			_ = _1253_v193
			_1253_v193 = _dafny.SeqOf((_this).F9(), !((_this).F9()) || (!((_this).F9())), true)
			r0 = _dafny.IntOfUint32((_1253_v193).Cardinality())
			var _1254_v194 _dafny.Set
			_ = _1254_v194
			_1254_v194 = _dafny.SetOf(_dafny.SeqOf(p0))
			var _1255_v195 _dafny.MultiSet
			_ = _1255_v195
			_1255_v195 = _dafny.MultiSetOf(p0, (_dafny.Zero).Minus((_dafny.Zero).Minus((_1251_i20).Times((_1254_v194).Cardinality()))))
			_1255_v195 = _1255_v195
			r0 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}((func(_1256_v188 _dafny.Sequence, _1257_p0 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_1258_i21 _dafny.Int) _dafny.CodePoint {
					return (_1256_v188).Select((Companion_Default___.SafeIndex(_1257_p0, _dafny.IntOfUint32((_1256_v188).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_1247_v188, p0)))).Cardinality())
		}
		var _1259_v196 _dafny.CodePoint
		_ = _1259_v196
		_1259_v196 = _dafny.CodePoint('b')
		var _1260_v197 D7
		_ = _1260_v197
		_1260_v197 = Companion_D7_.Create_DC20_((_this).F9(), _1259_v196)
		var _1261_v198 D7
		_ = _1261_v198
		_1261_v198 = Companion_D7_.Create_DC21_(_1260_v197)
		var _1262_v199 D7
		_ = _1262_v199
		_1262_v199 = Companion_D7_.Create_DC21_(_1260_v197)
		var _1263_v200 D7
		_ = _1263_v200
		_1263_v200 = Companion_D7_.Create_DC21_(_1261_v198)
		var _1264_v201 _dafny.Sequence
		_ = _1264_v201
		_1264_v201 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lo")).Cardinality()))
		var _rhs214 bool = ((_this).F9()) == ((_this).F9())
		_ = _rhs214
		var _rhs215 bool = p2
		_ = _rhs215
		var _rhs216 D7 = _1263_v200
		_ = _rhs216
		var _rhs217 bool = (p0).Cmp((func() _dafny.Int {
			if (_this).F9() {
				return p0
			}
			return _dafny.IntOfUint32((_1264_v201).Cardinality())
		})()) == 0
		_ = _rhs217
		var _rhs218 _dafny.Int = Companion_Default___.SafeModuloInt(p0, p0)
		_ = _rhs218
		var _lhs136 *GlobalState = globalState
		_ = _lhs136
		var _lhs137 *GlobalState = globalState
		_ = _lhs137
		var _lhs138 *GlobalState = globalState
		_ = _lhs138
		_lhs136.F2 = _rhs214
		_lhs137.F2 = _rhs215
		_1263_v200 = _rhs216
		_lhs138.F2 = _rhs217
		r0 = _rhs218
		var _1265_v202 _dafny.Array
		_ = _1265_v202
		var _nw196 _dafny.Array = _dafny.NewArrayWithValue(Companion_D15_.Default(), _dafny.IntOfInt64(19))
		_ = _nw196
		_1265_v202 = _nw196
		var _1266_v203 _dafny.Set
		_ = _1266_v203
		_1266_v203 = _dafny.SetOf(p0, p1)
		var _1267_v204 _dafny.Sequence
		_ = _1267_v204
		_1267_v204 = _dafny.SeqOf(_dafny.SetOf(p1))
		var _1268_v206 _dafny.Array
		_ = _1268_v206
		var _nwElement0_40 _dafny.Set = _1266_v203
		_ = _nwElement0_40
		var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(15))
		_ = _nw197
		(_nw197).ArraySet1(_nwElement0_40, 0)
		(_nw197).ArraySet1(_1266_v203, 1)
		(_nw197).ArraySet1(_1266_v203, 2)
		(_nw197).ArraySet1(_1266_v203, 3)
		(_nw197).ArraySet1(_1266_v203, 4)
		(_nw197).ArraySet1(_1266_v203, 5)
		(_nw197).ArraySet1(_1266_v203, 6)
		(_nw197).ArraySet1((_1267_v204).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_1267_v204).Cardinality()))).Uint32()).(_dafny.Set), 7)
		(_nw197).ArraySet1(_dafny.SetOf(p0, _dafny.IntOfInt64(-854)), 8)
		(_nw197).ArraySet1(_1266_v203, 9)
		(_nw197).ArraySet1(_1266_v203, 10)
		(_nw197).ArraySet1(_1266_v203, 11)
		(_nw197).ArraySet1(_1266_v203, 12)
		(_nw197).ArraySet1(_1266_v203, 13)
		(_nw197).ArraySet1(func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(96), _dafny.IntOfInt64(923))); ; {
				_compr_26, _ok30 := _iter30()
				if !_ok30 {
					break
				}
				var _1269_v205 _dafny.Int
				_1269_v205 = interface{}(_compr_26).(_dafny.Int)
				if ((_dafny.IntOfInt64(96)).Cmp(_1269_v205) <= 0) && ((_1269_v205).Cmp(_dafny.IntOfInt64(923)) < 0) {
					_coll26.Add(Companion_Default___.SafeDivisionInt(_1269_v205, p1))
				}
			}
			return _coll26.ToSet()
		}(), 14)
		_1268_v206 = _nw197
		var _1270_v207 T2
		_ = _1270_v207
		var _nw198 *C3 = New_C3_()
		_ = _nw198
		_nw198.Ctor__((_this).F9(), _1268_v206, Companion_Default___.Fm1(Companion_Default___.Fm29(!(p2), false, globalState), _dafny.IntOfUint32((_1247_v188).Cardinality()), globalState))
		_1270_v207 = _nw198
		var _1271_v208 D15
		_ = _1271_v208
		_1271_v208 = Companion_D15_.Create_DC42_((_this).F9(), p1, _1270_v207, (_1266_v203).Cardinality())
		var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1265_v202), 0))
		_ = _index185
		(_1265_v202).ArraySet1(_1271_v208, (_index185).Int())
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1265_v202), 0))
		_ = _index186
		(_1265_v202).ArraySet1(_1271_v208, (_index186).Int())
		(globalState).F2 = p2
		var _1272_v209 _dafny.MultiSet
		_ = _1272_v209
		_1272_v209 = _dafny.MultiSetOf(p1)
		var _1273_v210 _dafny.Map
		_ = _1273_v210
		_1273_v210 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_1270_v207).F9())
		var _1274_v211 D1
		_ = _1274_v211
		_1274_v211 = Companion_D1_.Create_DC4_(p1, (_1270_v207).F9(), (_this).F9(), p2, Companion_Default___.Fm1(_1272_v209, (_1273_v210).Cardinality(), globalState))
		r0 = (_dafny.Zero).Minus(((_1274_v211).Dtor_cf9()).Minus(p1))
		var _1275_v212 D4
		_ = _1275_v212
		_1275_v212 = Companion_D4_.Create_DC13_(p0)
		r1 = _1275_v212
		var _1276_v213 _dafny.MultiSet
		_ = _1276_v213
		_1276_v213 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_1247_v188, _dafny.Companion_Sequence_.Update(_1247_v188, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf((_this).F9())).Cardinality())).Cardinality()), _dafny.IntOfUint32((_1247_v188).Cardinality()))).Uint32(), _1259_v196)))
		r2 = _1276_v213
		r3 = p2
		return r0, r1, r2, r3
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f9  bool
	_f10 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f9 = false
	_this._f10 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C5{}
var _ T2 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F9() bool {
	return _this._f9
}
func (_this *C5) Ctor__(f10 _dafny.Int, f9 bool) {
	{
		(_this)._f10 = f10
		(_this)._f9 = f9
	}
}
func (_this *C5) Fm2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return (_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F9(), (_this).F9())).Cardinality())))
	}
}
func (_this *C5) Fm3(globalState *GlobalState) bool {
	{
		return (_this).F9()
	}
}
func (_this *C5) Fm4(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F10()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).F10(), (_this).F10()), (_dafny.Zero).Minus((_this).F10()))).Cardinality())
	}
}
func (_this *C5) Fm5(p0 _dafny.Set, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf((_dafny.Zero).Minus((_this).F10()), (_this).F10())).IsSubsetOf(_dafny.MultiSetOf((_this).F10(), _dafny.IntOfInt64(884)))
	}
}
func (_this *C5) M1(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Array, globalState *GlobalState) {
	{
		var _1277_v0 _dafny.Sequence
		_ = _1277_v0
		_1277_v0 = _dafny.UnicodeSeqOfUtf8Bytes("m")
		var _1278_v1 _dafny.Sequence
		_ = _1278_v1
		_1278_v1 = _dafny.SeqOf(_1277_v0, _1277_v0)
		var _hi4 _dafny.Int = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1278_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(379))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg70 _dafny.Int) interface{} {
				return coer70(arg70)
			}
		}((func(_1279_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_1280_i1 _dafny.Int) _dafny.Sequence {
				return _1279_v0
			}
		})(_1277_v0)))))).Cardinality()
		_ = _hi4
		for _1281_i0 := (_this).F10(); _1281_i0.Cmp(_hi4) < 0; _1281_i0 = _1281_i0.Plus(_dafny.One) {
			var _1282_v2 _dafny.CodePoint
			_ = _1282_v2
			_1282_v2 = _dafny.CodePoint('x')
			var _rhs219 bool = p2
			_ = _rhs219
			var _rhs220 _dafny.CodePoint = _dafny.CodePoint('c')
			_ = _rhs220
			var _lhs139 *GlobalState = globalState
			_ = _lhs139
			_lhs139.F2 = _rhs219
			_1282_v2 = _rhs220
			var _1283_v3 _dafny.Int
			_ = _1283_v3
			_1283_v3 = _dafny.IntOfInt64(504)
			var _1284_v4 _dafny.Array
			_ = _1284_v4
			var _nw199 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw199
			_1284_v4 = _nw199
			var _1285_v5 _dafny.Array
			_ = _1285_v5
			var _nw200 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw200
			_1285_v5 = _nw200
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))
			_ = _index187
			(_1284_v4).ArraySet1(_1285_v5, (_index187).Int())
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))
			_ = _index188
			var _rhs221 bool = p2
			_ = _rhs221
			var _rhs222 bool = p1
			_ = _rhs222
			var _rhs223 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if p2 {
					return _1281_i0
				}
				return _1281_i0
			})(), _dafny.IntOfUint32((_1277_v0).Cardinality()))
			_ = _rhs223
			var _rhs224 _dafny.Array = _1285_v5
			_ = _rhs224
			var _rhs225 bool = (_this).F9()
			_ = _rhs225
			var _lhs140 *GlobalState = globalState
			_ = _lhs140
			var _lhs141 *GlobalState = globalState
			_ = _lhs141
			var _lhs142 _dafny.Array = _1284_v4
			_ = _lhs142
			var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))
			_ = _lhs143
			var _lhs144 *GlobalState = globalState
			_ = _lhs144
			_lhs140.F2 = _rhs221
			_lhs141.F2 = _rhs222
			_1283_v3 = _rhs223
			(_lhs142).ArraySet1(_rhs224, (_lhs143).Int())
			_lhs144.F2 = _rhs225
			var _1286_v6 _dafny.Sequence
			_ = _1286_v6
			_1286_v6 = _dafny.SeqOf(p1, p1)
			var _1287_v7 _dafny.Map
			_ = _1287_v7
			_1287_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int())))
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))
			_ = _index189
			var _rhs226 _dafny.Array = (func() _dafny.Array {
				if (_1287_v7).Contains(p2) {
					return (_1287_v7).Get(p2).(_dafny.Array)
				}
				return _dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))
			})()
			_ = _rhs226
			var _rhs227 _dafny.Int = p0
			_ = _rhs227
			var _rhs228 _dafny.Sequence = _1286_v6
			_ = _rhs228
			var _lhs145 _dafny.Array = _1284_v4
			_ = _lhs145
			var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))
			_ = _lhs146
			(_lhs145).ArraySet1(_rhs226, (_lhs146).Int())
			_1283_v3 = _rhs227
			_1286_v6 = _rhs228
			var _1288_v8 _dafny.MultiSet
			_ = _1288_v8
			_1288_v8 = _dafny.MultiSetOf(_1281_i0)
			var _1289_v9 _dafny.MultiSet
			_ = _1289_v9
			_1289_v9 = _dafny.MultiSetOf(p1, p1, p1)
			var _1290_v10 _dafny.Map
			_ = _1290_v10
			_1290_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), Companion_Default___.Fm1(_1288_v8, (_1289_v9).Cardinality(), globalState))
			var _1291_v11 _dafny.MultiSet
			_ = _1291_v11
			_1291_v11 = _dafny.MultiSetOf(_1290_v10)
			if (_dafny.MultiSetOf(_1290_v10, _1290_v10)).IsProperSubsetOf(_1291_v11) {
				var _1292_v12 _dafny.Array
				_ = _1292_v12
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_39
				var _nw201 _dafny.Array
				_ = _nw201
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw201 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) _dafny.Int = (func(_1293_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1294_i2 _dafny.Int) _dafny.Int {
							return (_1294_i2).Minus(_1293_i0)
						}
					})(_1281_i0)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw201 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw201).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw201).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1292_v12 = _nw201
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v12), 0))
				_ = _index190
				(_1292_v12).ArraySet1(p0, (_index190).Int())
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(852), _dafny.ArrayLen((_1292_v12), 0))
				_ = _index191
				(_1292_v12).ArraySet1(_1281_i0, (_index191).Int())
				_1283_v3 = p0
				(globalState).F2 = !((true) == ((_this).F9()))
				_1290_v10 = (_1290_v10).Update(p1, (_this).F9())
				_1290_v10 = (_1290_v10).Update(!(p2), p2)
			} else {
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))
				_ = _arr2
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))), 0))
				_ = _index192
				(_arr2).ArraySet1((_this).F9(), (_index192).Int())
				var _1295_v13 _dafny.Array
				_ = _1295_v13
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_40
				var _nw202 _dafny.Array
				_ = _nw202
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw202 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Map = (func(_1296_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_1297_i3 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1296_p0)
						}
					})(p0)
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw202 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw202).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw202).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1295_v13 = _nw202
				var _1298_v14 _dafny.Map
				_ = _1298_v14
				_1298_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm6(_dafny.SeqOf(p2), globalState), _1277_v0), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm6(_dafny.SeqOf(p2), globalState), _1277_v0)).Cardinality()))).Uint32(), _1282_v2), _1295_v13)
				var _arr3 _dafny.Array = _dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))
				_ = _arr3
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))), 0))
				_ = _index193
				var _rhs229 _dafny.CodePoint = _1282_v2
				_ = _rhs229
				var _rhs230 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("srhyeyrp"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(574))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1299_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1300_i4 _dafny.Int) _dafny.CodePoint {
						return _1299_v2
					}
				})(_1282_v2)))), _1277_v0)
				_ = _rhs230
				var _rhs231 _dafny.Array = (func() _dafny.Array {
					if (_1298_v14).Contains(_dafny.Companion_Sequence_.Concatenate(_1277_v0, _1277_v0)) {
						return (_1298_v14).Get(_dafny.Companion_Sequence_.Concatenate(_1277_v0, _1277_v0)).(_dafny.Array)
					}
					return _1295_v13
				})()
				_ = _rhs231
				var _lhs147 _dafny.Array = _dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))), 0))
				_ = _lhs148
				_1282_v2 = _rhs229
				(_lhs147).ArraySet1(_rhs230, (_lhs148).Int())
				_1295_v13 = _rhs231
				var _1301_v15 _dafny.Map
				_ = _1301_v15
				_1301_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1286_v6).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1286_v6).Cardinality()))).Uint32()).(bool), _1282_v2)
				_1301_v15 = (_1301_v15).Update((_1288_v8).IsProperSubsetOf(_1288_v8), _1282_v2)
				var _1302_v16 _dafny.Map
				_ = _1302_v16
				_1302_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1290_v10, _1281_i0)
				var _1303_v17 _dafny.Sequence
				_ = _1303_v17
				_1303_v17 = _dafny.SeqOf((_this).F10())
				_1283_v3 = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_this).F9() {
						return (func() _dafny.Int {
							if (_1302_v16).Contains(_1290_v10) {
								return (_1302_v16).Get(_1290_v10).(_dafny.Int)
							}
							return (_this).F10()
						})()
					}
					return (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1303_v17, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1303_v17).Cardinality()))).Uint32(), p0))).Cardinality()
				})())
				var _1304_v18 D1
				_ = _1304_v18
				_1304_v18 = Companion_D1_.Create_DC1_(_1283_v3)
				var _1305_v19 _dafny.Map
				_ = _1305_v19
				_1305_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1277_v0, p2)
				var _1306_v20 _dafny.Map
				_ = _1306_v20
				_1306_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1303_v17, (_this).Fm2(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-758), _dafny.IntOfInt64(-196))).Cardinality()), p2, p1, (_1304_v18).Dtor_cf1(), globalState))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1305_v19).Cardinality(), Companion_Default___.Fm0(globalState)))
				var _1307_v21 _dafny.Set
				_ = _1307_v21
				_1307_v21 = _dafny.SetOf((_1286_v6).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))), 0))).Int()).(bool), p2)).Cardinality(), _dafny.IntOfUint32((_1286_v6).Cardinality()))).Uint32()).(bool))
				var _1308_v22 _dafny.Array
				_ = _1308_v22
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw203
				_1308_v22 = _nw203
				var _1309_v23 _dafny.Map
				_ = _1309_v23
				_1309_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1308_v22, _1281_i0)
				var _1310_v24 _dafny.Map
				_ = _1310_v24
				_1310_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1307_v21, p1)
				var _1311_v25 _dafny.MultiSet
				_ = _1311_v25
				_1311_v25 = _dafny.MultiSetOf(_1310_v24)
				var _1312_v26 _dafny.Array
				_ = _1312_v26
				var _nwElement0_41 _dafny.Int = (func() _dafny.Int {
					if (_1309_v23).Contains(_1308_v22) {
						return (_1309_v23).Get(_1308_v22).(_dafny.Int)
					}
					return _dafny.IntOfInt64(245)
				})()
				_ = _nwElement0_41
				var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(27))
				_ = _nw204
				(_nw204).ArraySet1(_nwElement0_41, 0)
				(_nw204).ArraySet1((_1288_v8).Cardinality(), 1)
				(_nw204).ArraySet1(p0, 2)
				(_nw204).ArraySet1((_1289_v9).Cardinality(), 3)
				(_nw204).ArraySet1(Companion_Default___.SafeDivisionInt(_1283_v3, _dafny.IntOfInt64(457)), 4)
				(_nw204).ArraySet1(Companion_Default___.Fm0(globalState), 5)
				(_nw204).ArraySet1(p0, 6)
				(_nw204).ArraySet1(((_this).F10()).Minus(_1281_i0), 7)
				(_nw204).ArraySet1(_1281_i0, 8)
				(_nw204).ArraySet1((_1281_i0).Plus(_dafny.IntOfUint32((_1277_v0).Cardinality())), 9)
				(_nw204).ArraySet1((_this).F10(), 10)
				(_nw204).ArraySet1((func() _dafny.Int {
					if (_1302_v16).Contains(_1290_v10) {
						return (_1302_v16).Get(_1290_v10).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_1277_v0).Cardinality())
				})(), 11)
				(_nw204).ArraySet1((_1281_i0).Minus((_1288_v8).Cardinality()), 12)
				(_nw204).ArraySet1(_dafny.IntOfUint32((_1277_v0).Cardinality()), 13)
				(_nw204).ArraySet1((_this).F10(), 14)
				(_nw204).ArraySet1((_1289_v9).Cardinality(), 15)
				(_nw204).ArraySet1(p0, 16)
				(_nw204).ArraySet1((_1283_v3).Times((_this).F10()), 17)
				(_nw204).ArraySet1((_1283_v3).Minus(_1283_v3), 18)
				(_nw204).ArraySet1(_dafny.IntOfInt64(845), 19)
				(_nw204).ArraySet1(Companion_Default___.SafeModuloInt(_1283_v3, _1281_i0), 20)
				(_nw204).ArraySet1((_dafny.Zero).Minus(p0), 21)
				(_nw204).ArraySet1(_1283_v3, 22)
				(_nw204).ArraySet1((_this).F10(), 23)
				(_nw204).ArraySet1((_1281_i0).Plus(p0), 24)
				(_nw204).ArraySet1((_dafny.IntOfUint32((_1277_v0).Cardinality())).Times((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1311_v25).Contains(_1310_v24) {
						return (_1311_v25).Multiplicity(_1310_v24)
					}
					return p0
				})())), 25)
				(_nw204).ArraySet1((func() _dafny.Int {
					if p1 {
						return _dafny.IntOfUint32((_1277_v0).Cardinality())
					}
					return _dafny.IntOfInt64(678)
				})(), 26)
				_1312_v26 = _nw204
				var _1313_v27 _dafny.Map
				_ = _1313_v27
				_1313_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-168))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}((func(_1314_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1315_i5 _dafny.Int) _dafny.CodePoint {
						return _1314_v2
					}
				})(_1282_v2)))).Cardinality()), (_dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))), 0))).Int()).(bool))
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1312_v26), 0))
				_ = _index194
				(_1312_v26).ArraySet1(((func() _dafny.Map {
					if p1 {
						return _1313_v27
					}
					return _1313_v27
				})()).Cardinality(), (_index194).Int())
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1312_v26), 0))
				_ = _index195
				var _rhs232 _dafny.Int = (_this).F10()
				_ = _rhs232
				var _rhs233 _dafny.Int = _dafny.IntOfInt64(854)
				_ = _rhs233
				var _rhs234 _dafny.Map = ((_1306_v20).Merge(_1306_v20)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1277_v0).Cardinality()), func() _dafny.Map {
					var _coll27 = _dafny.NewMapBuilder()
					_ = _coll27
					for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(644), _dafny.IntOfInt64(494))); ; {
						_compr_27, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _1316_v28 _dafny.Int
						_1316_v28 = interface{}(_compr_27).(_dafny.Int)
						if ((_dafny.IntOfInt64(644)).Cmp(_1316_v28) <= 0) && ((_1316_v28).Cmp(_dafny.IntOfInt64(494)) < 0) {
							_coll27.Add(Companion_Default___.SafeModuloInt(_1316_v28, p0), _1283_v3)
						}
					}
					return _coll27.ToMap()
				}()))
				_ = _rhs234
				var _rhs235 _dafny.Set = _1307_v21
				_ = _rhs235
				var _rhs236 _dafny.Int = _1281_i0
				_ = _rhs236
				var _lhs149 _dafny.Array = _1312_v26
				_ = _lhs149
				var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1312_v26), 0))
				_ = _lhs150
				_1283_v3 = _rhs232
				_1283_v3 = _rhs233
				_1306_v20 = _rhs234
				_1307_v21 = _rhs235
				(_lhs149).ArraySet1(_rhs236, (_lhs150).Int())
				var _arr4 _dafny.Array = _dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))
				_ = _arr4
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_dafny.ArrayCastTo((_1284_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1284_v4), 0))).Int()))), 0))
				_ = _index196
				(_arr4).ArraySet1((_this).Fm3(globalState), (_index196).Int())
			}
		}
		var _1317_v29 _dafny.MultiSet
		_ = _1317_v29
		_1317_v29 = _dafny.MultiSetOf(p0)
		var _rhs237 bool = ((func() _dafny.MultiSet {
			if (_this).F9() {
				return _1317_v29
			}
			return _1317_v29
		})()).IsDisjointFrom(_1317_v29)
		_ = _rhs237
		var _rhs238 _dafny.Sequence = _1277_v0
		_ = _rhs238
		var _lhs151 *GlobalState = globalState
		_ = _lhs151
		_lhs151.F2 = _rhs237
		_1277_v0 = _rhs238
		var _1318_v30 D1
		_ = _1318_v30
		_1318_v30 = Companion_D1_.Create_DC2_((_this).F10(), (_this).F9(), p0, (Companion_D1_.Create_DC4_(_dafny.IntOfInt64(347), (_this).F9(), p1, false, p2)).Dtor_cf9())
		var _1319_v31 _dafny.Sequence
		_ = _1319_v31
		_1319_v31 = _dafny.SeqOf(p0, (_1318_v30).Dtor_cf2())
		(globalState).F2 = ((_dafny.IntOfUint32((_1319_v31).Cardinality())).Times(p0)).Cmp(p0) > 0
		var _hi5 _dafny.Int = (_this).F10()
		_ = _hi5
		for _1320_i6 := p0; _1320_i6.Cmp(_hi5) < 0; _1320_i6 = _1320_i6.Plus(_dafny.One) {
			var _1321_v32 bool
			_ = _1321_v32
			var _1322_v33 _dafny.Int
			_ = _1322_v33
			var _1323_v34 _dafny.Int
			_ = _1323_v34
			var _out14 bool
			_ = _out14
			var _out15 _dafny.Int
			_ = _out15
			var _out16 _dafny.Int
			_ = _out16
			_out14, _out15, _out16 = (_this).M5(globalState)
			_1321_v32 = _out14
			_1322_v33 = _out15
			_1323_v34 = _out16
			var _1324_v35 _dafny.Map
			_ = _1324_v35
			_1324_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1321_v32, p1)
			var _1325_v36 _dafny.Sequence
			_ = _1325_v36
			_1325_v36 = _dafny.SeqOf((_this).F9(), !(_1321_v32))
			var _1326_v37 _dafny.Map
			_ = _1326_v37
			_1326_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1324_v35, _1325_v36)
			var _1327_v38 _dafny.Set
			_ = _1327_v38
			_1327_v38 = _dafny.SetOf(_1319_v31)
			var _1328_v39 _dafny.Map
			_ = _1328_v39
			_1328_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1323_v34, _1325_v36)
			_1322_v33 = _dafny.IntOfUint32((Companion_Default___.Fm7((_1326_v37).Cardinality(), p2, Companion_D1_.Create_DC4_(_1322_v33, p1, !((_this).Fm5(_1327_v38, _1328_v39, _dafny.IntOfUint32((_1277_v0).Cardinality()), globalState)), !(_1321_v32), false), (_dafny.IntOfUint32((_1325_v36).Cardinality())).Cmp(_1320_i6) > 0, globalState)).Cardinality())
			var _1329_v40 _dafny.Map
			_ = _1329_v40
			_1329_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).Fm2(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1319_v31, (Companion_Default___.SafeIndex(_1323_v34, _dafny.IntOfUint32((_1319_v31).Cardinality()))).Uint32(), _dafny.IntOfInt64(263))).Cardinality()), (_this).F9(), false, (_this).F10(), globalState)).Cardinality()), Companion_D1_.Create_DC1_(_dafny.IntOfInt64(794)))
			var _1330_v41 D1
			_ = _1330_v41
			_1330_v41 = Companion_D1_.Create_DC1_(_1323_v34)
			_1329_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_1320_i6, _1323_v34), _1330_v41)
			var _1331_v42 *C0
			_ = _1331_v42
			var _nw205 *C0 = New_C0_()
			_ = _nw205
			_nw205.Ctor__(_1327_v38, _1322_v33)
			_1331_v42 = _nw205
		}
		var _1332_v43 D1
		_ = _1332_v43
		_1332_v43 = Companion_D1_.Create_DC4_(p0, (_this).F9(), true, (_this).F9(), (_this).F9())
		var _1333_v44 _dafny.Sequence
		_ = _1333_v44
		_1333_v44 = _dafny.SeqOf(_1332_v43, _1332_v43, Companion_D1_.Create_DC4_((_this).F10(), !(p1), !(p1), p2, p2), _1332_v43, _1332_v43)
		var _1334_v45 _dafny.Sequence
		_ = _1334_v45
		_1334_v45 = _dafny.SeqOf(_1333_v44, _1333_v44, _1333_v44)
		_1333_v44 = _dafny.Companion_Sequence_.Concatenate((_1334_v45).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1334_v45).Cardinality()))).Uint32()).(_dafny.Sequence), (_1334_v45).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1334_v45).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _1335_v46 _dafny.Int
		_ = _1335_v46
		_1335_v46 = _dafny.IntOfInt64(38)
		var _1336_v47 _dafny.Sequence
		_ = _1336_v47
		_1336_v47 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(884))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg73 _dafny.Int) interface{} {
				return coer73(arg73)
			}
		}((func(_1337_v46 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1338_i7 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1337_v46, (_this).F9())).Cardinality()
			}
		})(_1335_v46)))
		var _1339_v48 _dafny.Set
		_ = _1339_v48
		_1339_v48 = _dafny.SetOf(_1319_v31, _1319_v31)
		var _1340_v49 *C0
		_ = _1340_v49
		var _nw206 *C0 = New_C0_()
		_ = _nw206
		_nw206.Ctor__(_1339_v48, _1335_v46)
		_1340_v49 = _nw206
		var _1341_v50 _dafny.Map
		_ = _1341_v50
		_1341_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1340_v49)
		var _rhs239 bool = (p0).Cmp((_dafny.Zero).Minus(_1335_v46)) != 0
		_ = _rhs239
		var _rhs240 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1319_v31).Cardinality()), ((_this).F10()).Times((_1341_v50).Cardinality()))
		_ = _rhs240
		var _rhs241 _dafny.Sequence = Companion_Default___.Fm9(Companion_Default___.SafeDivisionInt((_this).F10(), p0), _1335_v46, false, globalState)
		_ = _rhs241
		var _rhs242 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("yuoebdhit")
		_ = _rhs242
		var _lhs152 *GlobalState = globalState
		_ = _lhs152
		_lhs152.F2 = _rhs239
		_1335_v46 = _rhs240
		_1336_v47 = _rhs241
		_1277_v0 = _rhs242
	}
}
func (_this *C5) M2(p0 _dafny.Array, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) (bool, bool, _dafny.Int, _dafny.Set) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Set = _dafny.EmptySet
		_ = r3
		var _1342_i0 _dafny.Int
		_ = _1342_i0
		_1342_i0 = _dafny.Zero
		{
			for (_this).F9() {
				{
					if (_1342_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1342_i0 = (_1342_i0).Plus(_dafny.One)
					var _1343_v0 _dafny.Set
					_ = _1343_v0
					_1343_v0 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(p3, (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((p3).Cardinality()))).Uint32(), (_this).F10()))
					var _1344_v1 *C0
					_ = _1344_v1
					var _nw207 *C0 = New_C0_()
					_ = _nw207
					_nw207.Ctor__(_1343_v0, ((p3).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((p3).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_this).F10()))
					_1344_v1 = _nw207
					var _1345_v2 _dafny.Array
					_ = _1345_v2
					var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
					_ = _nw208
					_1345_v2 = _nw208
					var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_1345_v2), 0))
					_ = _index197
					(_1345_v2).ArraySet1(Companion_Default___.SafeDivisionInt(_1344_v1.F12, _1344_v1.F12), (_index197).Int())
					var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_1345_v2), 0))
					_ = _index198
					(_1345_v2).ArraySet1(_1344_v1.F12, (_index198).Int())
					r1 = (_this).F9()
					var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_1345_v2), 0))
					_ = _index199
					(_1345_v2).ArraySet1(Companion_Default___.SafeDivisionInt((_1345_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_1345_v2), 0))).Int()).(_dafny.Int), (_1345_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_1345_v2), 0))).Int()).(_dafny.Int)), (_index199).Int())
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1346_v3 _dafny.Sequence
		_ = _1346_v3
		_1346_v3 = _dafny.UnicodeSeqOfUtf8Bytes("qgopc")
		var _1347_v4 _dafny.Set
		_ = _1347_v4
		_1347_v4 = _dafny.SetOf((_this).F10())
		var _1348_v5 _dafny.Map
		_ = _1348_v5
		_1348_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_1347_v4).IsDisjointFrom(_1347_v4))
		var _1349_v6 _dafny.Array
		_ = _1349_v6
		var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
		_ = _nw209
		_1349_v6 = _nw209
		var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))
		_ = _index200
		(_1349_v6).ArraySet1(((_this).F10()).Minus((_this).F10()), (_index200).Int())
		var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))
		_ = _index201
		var _rhs243 _dafny.Sequence = p1
		_ = _rhs243
		var _rhs244 _dafny.Map = _1348_v5
		_ = _rhs244
		var _rhs245 bool = (_dafny.SetOf((_this).F10(), _dafny.IntOfInt64(280), (_this).F10(), (_this).F10())).Contains(Companion_Default___.Fm0(globalState))
		_ = _rhs245
		var _rhs246 _dafny.Int = (_this).F10()
		_ = _rhs246
		var _rhs247 bool = !((_this).F9()) || (((_this).F9()) && (false))
		_ = _rhs247
		var _lhs153 _dafny.Array = _1349_v6
		_ = _lhs153
		var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))
		_ = _lhs154
		_1346_v3 = _rhs243
		_1348_v5 = _rhs244
		r1 = _rhs245
		(_lhs153).ArraySet1(_rhs246, (_lhs154).Int())
		r1 = _rhs247
		if (_this).F9() {
			var _1350_v7 _dafny.Map
			_ = _1350_v7
			_1350_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), true)
			var _1351_v8 D1
			_ = _1351_v8
			_1351_v8 = Companion_D1_.Create_DC3_(!(_1350_v7).Contains(Companion_Default___.Fm0(globalState)), p1, ((_this).F10()).Times((_this).F10()))
			var _1352_v9 _dafny.Set
			_ = _1352_v9
			_1352_v9 = _dafny.SetOf(p2, p2, p2, (func() _dafny.CodePoint {
				if (_this).F9() {
					return p2
				}
				return p2
			})())
			var _1353_v10 _dafny.MultiSet
			_ = _1353_v10
			_1353_v10 = _dafny.MultiSetOf((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int))
			var _1354_v11 _dafny.Array
			_ = _1354_v11
			var _nwElement0_42 bool = ((Companion_Default___.Fm10((_this).F9(), globalState)).Cardinality()).Cmp((_this).F10()) <= 0
			_ = _nwElement0_42
			var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(21))
			_ = _nw210
			(_nw210).ArraySet1(_nwElement0_42, 0)
			(_nw210).ArraySet1(!((_this).F9()), 1)
			(_nw210).ArraySet1((_this).F9(), 2)
			(_nw210).ArraySet1((_this).F9(), 3)
			(_nw210).ArraySet1((_this).F9(), 4)
			(_nw210).ArraySet1((_1351_v8).Dtor_cf6(), 5)
			(_nw210).ArraySet1((_this).F9(), 6)
			(_nw210).ArraySet1((_this).F9(), 7)
			(_nw210).ArraySet1((_this).F9(), 8)
			(_nw210).ArraySet1((_this).F9(), 9)
			(_nw210).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_1346_v3, (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1346_v3).Cardinality()))).Uint32(), Companion_Default___.Fm11((_this).F9(), globalState)), p2), 10)
			(_nw210).ArraySet1((_this).F9(), 11)
			(_nw210).ArraySet1(!((_this).F9()), 12)
			(_nw210).ArraySet1(((Companion_D1_.Create_DC1_((_this).F10())).Dtor_cf1()).Cmp((_this).F10()) >= 0, 13)
			(_nw210).ArraySet1((_this).F9(), 14)
			(_nw210).ArraySet1((_this).F9(), 15)
			(_nw210).ArraySet1((_this).F9(), 16)
			(_nw210).ArraySet1(Companion_Default___.Fm1(_dafny.MultiSetFromSeq(p3), (_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), globalState), 17)
			(_nw210).ArraySet1((_1353_v10).IsProperSubsetOf(_1353_v10), 18)
			(_nw210).ArraySet1((_this).F9(), 19)
			(_nw210).ArraySet1((_this).F9(), 20)
			_1354_v11 = _nw210
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1354_v11), 0))
			_ = _index202
			(_1354_v11).ArraySet1((Companion_Default___.Fm12((_1351_v8).Dtor_cf6(), _dafny.IntOfUint32((_1346_v3).Cardinality()), (_this).F9(), _dafny.IntOfUint32((_1346_v3).Cardinality()), globalState)).IsDisjointFrom(Companion_Default___.Fm12(true, (_this).F10(), (_this).F9(), (_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), globalState)), (_index202).Int())
			var _1355_v12 _dafny.Sequence
			_ = _1355_v12
			_1355_v12 = _dafny.SeqOf((_this).F9())
			var _1356_v13 D1
			_ = _1356_v13
			_1356_v13 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(144), (_this).F9(), (_this).F9(), (_this).F9(), (_1355_v12).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1355_v12).Cardinality()))).Uint32()).(bool))
			var _pat_let_tv69 = _1349_v6
			_ = _pat_let_tv69
			var _pat_let_tv70 = _1346_v3
			_ = _pat_let_tv70
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1354_v11), 0))
			_ = _index203
			var _rhs248 D1 = func(_pat_let50_0 D1) D1 {
				return func(_1357_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let51_0 bool) D1 {
						return func(_1360_dt__update_hcf6_h0 bool) D1 {
							return Companion_D1_.Create_DC3_(_1360_dt__update_hcf6_h0, (_1357_dt__update__tmp_h0).Dtor_cf7(), (_1357_dt__update__tmp_h0).Dtor_cf8())
						}(_pat_let51_0)
					}(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(68))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg74 _dafny.Int) interface{} {
							return coer74(arg74)
						}
					}((func(_1358_v6 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_1359_i1 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus((_1358_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1358_v6), 0))).Int()).(_dafny.Int))
						}
					})(_pat_let_tv69))), _dafny.SeqOf(_dafny.IntOfUint32((_pat_let_tv70).Cardinality()))))
				}(_pat_let50_0)
			}(_1351_v8)
			_ = _rhs248
			var _rhs249 _dafny.Set = _1352_v9
			_ = _rhs249
			var _rhs250 bool = (_this).F9()
			_ = _rhs250
			var _rhs251 _dafny.Int = (_dafny.Zero).Minus((_1356_v13).Dtor_cf9())
			_ = _rhs251
			var _rhs252 _dafny.Int = (_this).F10()
			_ = _rhs252
			var _lhs155 _dafny.Array = _1354_v11
			_ = _lhs155
			var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1354_v11), 0))
			_ = _lhs156
			_1351_v8 = _rhs248
			_1352_v9 = _rhs249
			(_lhs155).ArraySet1(_rhs250, (_lhs156).Int())
			r2 = _rhs251
			r2 = _rhs252
			if ((_1354_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1354_v11), 0))).Int()).(bool)) || ((_this).F9()) {
				var _1361_v14 _dafny.Set
				_ = _1361_v14
				_1361_v14 = _dafny.SetOf(p3)
				var _1362_v15 _dafny.Map
				_ = _1362_v15
				_1362_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), _1355_v12)
				(globalState).F2 = (_this).Fm5(_1361_v14, _1362_v15, ((_this).F10()).Times((_this).F10()), globalState)
				r0 = (_1354_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1354_v11), 0))).Int()).(bool)
				var _1363_v16 _dafny.Map
				_ = _1363_v16
				_1363_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.UnicodeSeqOfUtf8Bytes("gjurujc"))
				_1346_v3 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1363_v16).Contains((_dafny.Zero).Minus((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int))) {
						return (_1363_v16).Get((_dafny.Zero).Minus((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int))).(_dafny.Sequence)
					}
					return p1
				})(), p1)
				var _1364_v17 _dafny.MultiSet
				_ = _1364_v17
				_1364_v17 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.IsPrefixOf(p1, _dafny.UnicodeSeqOfUtf8Bytes("vfyqtatu")))
				_1364_v17 = _1364_v17
				_1353_v10 = _1353_v10
			} else {
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))
				_ = _index204
				var _rhs253 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F10(), (_this).F10())
				_ = _rhs253
				var _rhs254 _dafny.Int = (_this).F10()
				_ = _rhs254
				var _rhs255 bool = !((_this).F9())
				_ = _rhs255
				var _rhs256 _dafny.Int = (_this).F10()
				_ = _rhs256
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				var _lhs158 _dafny.Array = _1349_v6
				_ = _lhs158
				var _lhs159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))
				_ = _lhs159
				r2 = _rhs253
				r2 = _rhs254
				_lhs157.F2 = _rhs255
				(_lhs158).ArraySet1(_rhs256, (_lhs159).Int())
				var _1365_v18 _dafny.Map
				_ = _1365_v18
				_1365_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p2)
				var _1366_v20 _dafny.Sequence
				_ = _1366_v20
				_1366_v20 = _dafny.SeqOf(func() _dafny.Set {
					var _coll28 = _dafny.NewBuilder()
					_ = _coll28
					for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(545), _dafny.IntOfInt64(560))); ; {
						_compr_28, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _1367_v19 _dafny.Int
						_1367_v19 = interface{}(_compr_28).(_dafny.Int)
						if ((_dafny.IntOfInt64(545)).Cmp(_1367_v19) <= 0) && ((_1367_v19).Cmp(_dafny.IntOfInt64(560)) < 0) {
							_coll28.Add((_1367_v19).Plus((_this).F10()))
						}
					}
					return _coll28.ToSet()
				}(), (_1347_v4), _1347_v4)
				var _1368_v21 _dafny.Map
				_ = _1368_v21
				_1368_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _1369_v22 _dafny.Array
				_ = _1369_v22
				var _nwElement0_43 _dafny.CodePoint = p2
				_ = _nwElement0_43
				var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(8))
				_ = _nw211
				(_nw211).ArraySet1CodePoint(_nwElement0_43, 0)
				(_nw211).ArraySet1CodePoint(p2, 1)
				(_nw211).ArraySet1CodePoint(p2, 2)
				(_nw211).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_1365_v18).Contains(_dafny.IntOfUint32((_1366_v20).Cardinality())) {
						return (_1365_v18).Get(_dafny.IntOfUint32((_1366_v20).Cardinality())).(_dafny.CodePoint)
					}
					return p2
				})(), 3)
				(_nw211).ArraySet1CodePoint(p2, 4)
				(_nw211).ArraySet1CodePoint(p2, 5)
				(_nw211).ArraySet1CodePoint(p2, 6)
				(_nw211).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_1368_v21).Contains(p2) {
						return (_1368_v21).Get(p2).(_dafny.CodePoint)
					}
					return p2
				})(), 7)
				_1369_v22 = _nw211
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1369_v22), 0))
				_ = _index205
				(_1369_v22).ArraySet1CodePoint(p2, (_index205).Int())
				var _1370_v23 _dafny.Sequence
				_ = _1370_v23
				_1370_v23 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()), ((_this).F10()).Times((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int)))
				var _1371_v24 *C0
				_ = _1371_v24
				var _nw212 *C0 = New_C0_()
				_ = _nw212
				_nw212.Ctor__(_dafny.SetOf(_1370_v23), (_this).F10())
				_1371_v24 = _nw212
				var _1372_v25 _dafny.Sequence
				_ = _1372_v25
				_1372_v25 = _dafny.SeqOf(_1371_v24)
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1369_v22), 0))
				_ = _index206
				var _rhs257 _dafny.CodePoint = p2
				_ = _rhs257
				var _rhs258 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("lg")
				_ = _rhs258
				var _rhs259 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("lf")
				_ = _rhs259
				var _rhs260 _dafny.Sequence = _1370_v23
				_ = _rhs260
				var _rhs261 _dafny.Sequence = _1372_v25
				_ = _rhs261
				var _lhs160 _dafny.Array = _1369_v22
				_ = _lhs160
				var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_1369_v22), 0))
				_ = _lhs161
				(_lhs160).ArraySet1CodePoint(_rhs257, (_lhs161).Int())
				_1346_v3 = _rhs258
				_1346_v3 = _rhs259
				_1370_v23 = _rhs260
				_1372_v25 = _rhs261
				(globalState).F2 = ((_this).F10()).Cmp(_dafny.IntOfUint32((p3).Cardinality())) >= 0
				var _1373_v26 _dafny.Array
				_ = _1373_v26
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_41
				var _nw213 _dafny.Array
				_ = _nw213
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw213 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) _dafny.Map = (func(_1374_v6 _dafny.Array, _1375_v10 _dafny.MultiSet) func(_dafny.Int) _dafny.Map {
						return func(_1376_i2 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1374_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1374_v6), 0))).Int()).(_dafny.Int), _1375_v10)
						}
					})(_1349_v6, _1353_v10)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw213 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw213).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw213).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1373_v26 = _nw213
				var _1377_v27 _dafny.Map
				_ = _1377_v27
				_1377_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1347_v4).Cardinality(), _dafny.MultiSetFromSeq(p3))
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_1373_v26), 0))
				_ = _index207
				(_1373_v26).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), _1353_v10)).Merge(_1377_v27), (_index207).Int())
				var _1378_v28 D1
				_ = _1378_v28
				_1378_v28 = Companion_D1_.Create_DC2_((_this).F10(), false, (_this).F10(), (_dafny.Zero).Minus((_this).F10()))
				var _1379_v29 _dafny.Sequence
				_ = _1379_v29
				_1379_v29 = _dafny.SeqOf(_1378_v28, _1378_v28)
				var _1380_v30 T0
				_ = _1380_v30
				var _nw214 *C1 = New_C1_()
				_ = _nw214
				_nw214.Ctor__(_1379_v29)
				_1380_v30 = _nw214
				var _1381_v31 _dafny.Map
				_ = _1381_v31
				_1381_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1380_v30)
				var _1382_v32 D1
				_ = _1382_v32
				_1382_v32 = Companion_D1_.Create_DC2_((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), (_this).F9(), (_1381_v31).Cardinality(), (_this).F10())
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_1373_v26), 0))
				_ = _index208
				(_1373_v26).ArraySet1((_1377_v27).Merge((Companion_Default___.Fm13((_this).Fm4((_1354_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1354_v11), 0))).Int()).(bool), (_this).F9(), (_this).F9(), (_1382_v32).Dtor_cf3(), globalState), (_1354_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1354_v11), 0))).Int()).(bool), (_this).F9(), globalState)).Merge(_1377_v27)), (_index208).Int())
				var _1383_v33 *C0
				_ = _1383_v33
				var _nw215 *C0 = New_C0_()
				_ = _nw215
				_nw215.Ctor__(_1371_v24.F11, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(839))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1384_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1385_i3 _dafny.Int) _dafny.CodePoint {
						return _1384_p2
					}
				})(p2)))).Cardinality()))
				_1383_v33 = _nw215
			}
			var _1386_v34 _dafny.Array
			_ = _1386_v34
			var _nw216 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
			_ = _nw216
			_1386_v34 = _nw216
			var _1387_v35 _dafny.Sequence
			_ = _1387_v35
			_1387_v35 = _dafny.SeqOf(_1386_v34)
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1354_v11), 0))
			_ = _index209
			(_1354_v11).ArraySet1(_dafny.Companion_Sequence_.Contains(_1387_v35, _1386_v34), (_index209).Int())
			Companion_Default___.M0((_this).F9(), globalState)
			var _1388_v36 D1
			_ = _1388_v36
			_1388_v36 = Companion_D1_.Create_DC2_((_this).F10(), (_this).F9(), (_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), (_this).F10())
			var _1389_v37 _dafny.Sequence
			_ = _1389_v37
			_1389_v37 = _dafny.SeqOf(_1388_v36, _1388_v36)
			var _1390_v38 *C1
			_ = _1390_v38
			var _nw217 *C1 = New_C1_()
			_ = _nw217
			_nw217.Ctor__(_1389_v37)
			_1390_v38 = _nw217
		} else {
			var _1391_v39 _dafny.Array
			_ = _1391_v39
			var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw218
			_1391_v39 = _nw218
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1391_v39), 0))
			_ = _index210
			(_1391_v39).ArraySet1(_1346_v3, (_index210).Int())
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1391_v39), 0))
			_ = _index211
			(_1391_v39).ArraySet1(_1346_v3, (_index211).Int())
			var _1392_v40 _dafny.Array
			_ = _1392_v40
			var _nw219 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
			_ = _nw219
			_1392_v40 = _nw219
			var _1393_v41 D4
			_ = _1393_v41
			_1393_v41 = Companion_D4_.Create_DC11_(_1392_v40)
			var _1394_v42 _dafny.Array
			_ = _1394_v42
			var _nwElement0_44 _dafny.Array = _1392_v40
			_ = _nwElement0_44
			var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(19))
			_ = _nw220
			(_nw220).ArraySet1(_nwElement0_44, 0)
			(_nw220).ArraySet1(_1392_v40, 1)
			(_nw220).ArraySet1(_1392_v40, 2)
			(_nw220).ArraySet1(_1392_v40, 3)
			(_nw220).ArraySet1(_1392_v40, 4)
			(_nw220).ArraySet1(_1392_v40, 5)
			(_nw220).ArraySet1(_1392_v40, 6)
			(_nw220).ArraySet1(_1392_v40, 7)
			(_nw220).ArraySet1(_1392_v40, 8)
			(_nw220).ArraySet1(_1392_v40, 9)
			(_nw220).ArraySet1(_1392_v40, 10)
			(_nw220).ArraySet1(_1392_v40, 11)
			(_nw220).ArraySet1(_1392_v40, 12)
			(_nw220).ArraySet1(_1392_v40, 13)
			(_nw220).ArraySet1(_1392_v40, 14)
			(_nw220).ArraySet1(_1392_v40, 15)
			(_nw220).ArraySet1(_1392_v40, 16)
			(_nw220).ArraySet1((_1393_v41).Dtor_cf19(), 17)
			(_nw220).ArraySet1(_1392_v40, 18)
			_1394_v42 = _nw220
			var _1395_v43 _dafny.Array
			_ = _1395_v43
			_1395_v43 = _1394_v42
			_1394_v42 = (_1395_v43)
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))
			_ = _index212
			(_1349_v6).ArraySet1((_this).F10(), (_index212).Int())
			var _1396_v44 D1
			_ = _1396_v44
			_1396_v44 = Companion_D1_.Create_DC2_((_this).F10(), (_this).F9(), (_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(247))
			var _1397_v45 *C1
			_ = _1397_v45
			var _nw221 *C1 = New_C1_()
			_ = _nw221
			_nw221.Ctor__(_dafny.SeqOf(_1396_v44, _1396_v44))
			_1397_v45 = _nw221
			var _1398_v46 _dafny.Map
			_ = _1398_v46
			_1398_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_dafny.Zero).Minus((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int)))
			_1398_v46 = (_1398_v46).Merge(_1398_v46)
		}
		r2 = ((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int)).Minus((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int))
		var _1399_v47 _dafny.Array
		_ = _1399_v47
		var _nw222 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw222
		_1399_v47 = _nw222
		var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1399_v47), 0))
		_ = _index213
		(_1399_v47).ArraySet1((_this).F9(), (_index213).Int())
		var _1400_v49 _dafny.Sequence
		_ = _1400_v49
		_1400_v49 = _dafny.SeqOf(_1399_v47)
		var _1401_v50 _dafny.MultiSet
		_ = _1401_v50
		_1401_v50 = _dafny.MultiSetOf(_dafny.IntOfUint32((p1).Cardinality()))
		var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1399_v47), 0))
		_ = _index214
		var _rhs262 bool = !((Companion_Default___.Fm12((_this).F9(), (_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), (_this).F9(), (_this).F10(), globalState)).IsSubsetOf(func() _dafny.Set {
			var _coll29 = _dafny.NewBuilder()
			_ = _coll29
			for _iter33 := _dafny.Iterate((_1347_v4).Elements()); ; {
				_compr_29, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _1402_v48 _dafny.Int
				_1402_v48 = interface{}(_compr_29).(_dafny.Int)
				if (_1347_v4).Contains(_1402_v48) {
					_coll29.Add((_1402_v48).Times(_dafny.IntOfInt64(78)))
				}
			}
			return _coll29.ToSet()
		}()))
		_ = _rhs262
		var _rhs263 _dafny.Array = _1399_v47
		_ = _rhs263
		var _rhs264 _dafny.Array = (_1400_v49).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1400_v49).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs264
		var _rhs265 bool = Companion_Default___.Fm1(_1401_v50, (_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), globalState)
		_ = _rhs265
		var _lhs162 _dafny.Array = _1399_v47
		_ = _lhs162
		var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1399_v47), 0))
		_ = _lhs163
		(_lhs162).ArraySet1(_rhs262, (_lhs163).Int())
		_1399_v47 = _rhs263
		_1399_v47 = _rhs264
		r1 = _rhs265
		var _1403_v51 D3
		_ = _1403_v51
		_1403_v51 = Companion_D3_.Create_DC10_()
		_1403_v51 = (func() D3 {
			if (_this).F9() {
				return _1403_v51
			}
			return _1403_v51
		})()
		r0 = (_1399_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1399_v47), 0))).Int()).(bool)
		var _1404_v52 _dafny.Sequence
		_ = _1404_v52
		_1404_v52 = _dafny.SeqOf(p3)
		var _1405_v54 _dafny.Sequence
		_ = _1405_v54
		_1405_v54 = _dafny.SeqOf((_this).F9())
		var _1406_v55 D3
		_ = _1406_v55
		_1406_v55 = Companion_D3_.Create_DC9_((_this).F9(), _1405_v54)
		var _1407_v56 _dafny.Map
		_ = _1407_v56
		_1407_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), _dafny.SeqOf(_dafny.IntOfUint32(((_1406_v55).Dtor_cf18()).Cardinality()), (_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), (_this).F10()))
		var _1408_v57 _dafny.Set
		_ = _1408_v57
		_1408_v57 = _dafny.SetOf((func() _dafny.Sequence {
			if (_1407_v56).Contains((_this).F10()) {
				return (_1407_v56).Get((_this).F10()).(_dafny.Sequence)
			}
			return p3
		})())
		var _1409_v58 _dafny.Map
		_ = _1409_v58
		_1409_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1405_v54)
		r1 = (_this).Fm5((func() _dafny.Set {
			var _coll30 = _dafny.NewBuilder()
			_ = _coll30
			for _iter34 := _dafny.Iterate((_1404_v52).Elements()); ; {
				_compr_30, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _1410_v53 _dafny.Sequence
				_1410_v53 = interface{}(_compr_30).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_1404_v52, _1410_v53) {
					_coll30.Add(_1410_v53)
				}
			}
			return _coll30.ToSet()
		}()).Union(_1408_v57), (_1409_v58).Merge(_1409_v58), (_1349_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_1349_v6), 0))).Int()).(_dafny.Int), globalState)
		r2 = (_this).F10()
		r3 = (_1347_v4).Intersection(_1347_v4)
		return r0, r1, r2, r3
	}
}
func (_this *C5) M3(p0 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _hi6 _dafny.Int = (_this).F10()
		_ = _hi6
		for _1411_i0 := (_this).F10(); _1411_i0.Cmp(_hi6) < 0; _1411_i0 = _1411_i0.Plus(_dafny.One) {
			if p0 {
				var _1412_v0 _dafny.Set
				_ = _1412_v0
				_1412_v0 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-810))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1413_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1414_i1 _dafny.Int) _dafny.Int {
						return _1413_i0
					}
				})(_1411_i0))))
				var _1415_v1 *C0
				_ = _1415_v1
				var _nw223 *C0 = New_C0_()
				_ = _nw223
				_nw223.Ctor__(_1412_v0, ((_this).F10()).Plus(_1411_i0))
				_1415_v1 = _nw223
				var _1416_v2 _dafny.MultiSet
				_ = _1416_v2
				_1416_v2 = _dafny.MultiSetOf((_this).F10(), (_this).F10())
				var _1417_v3 D1
				_ = _1417_v3
				_1417_v3 = Companion_D1_.Create_DC4_((_dafny.Zero).Minus((_dafny.Zero).Minus(((_this).F10()).Times(_1415_v1.F12))), !((_1416_v2).IsProperSubsetOf(_1416_v2)), (_this).F9(), (_this).F9(), !(p0))
				_1417_v3 = Companion_D1_.Create_DC4_(_1411_i0, (_this).F9(), p0, (_this).F9(), p0)
				var _1418_v4 _dafny.Sequence
				_ = _1418_v4
				_1418_v4 = _dafny.SeqOf((_this).F10(), _dafny.IntOfInt64(267))
				var _1419_v5 _dafny.Sequence
				_ = _1419_v5
				_1419_v5 = _dafny.UnicodeSeqOfUtf8Bytes("fdwey")
				var _1420_v6 D1
				_ = _1420_v6
				_1420_v6 = Companion_D1_.Create_DC1_(_dafny.IntOfUint32((_1419_v5).Cardinality()))
				var _pat_let_tv71 = globalState
				_ = _pat_let_tv71
				var _1421_v7 _dafny.Array
				_ = _1421_v7
				var _nwElement0_45 D1 = Companion_Default___.Fm18(_dafny.IntOfUint32((_1418_v4).Cardinality()), p0, p0, globalState)
				_ = _nwElement0_45
				var _nw224 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(27))
				_ = _nw224
				(_nw224).ArraySet1(_nwElement0_45, 0)
				(_nw224).ArraySet1(_1420_v6, 1)
				(_nw224).ArraySet1(_1420_v6, 2)
				(_nw224).ArraySet1(Companion_Default___.Fm18(_1411_i0, p0, p0, globalState), 3)
				(_nw224).ArraySet1((func() D1 {
					if p0 {
						return _1420_v6
					}
					return _1420_v6
				})(), 4)
				(_nw224).ArraySet1(Companion_D1_.Create_DC1_((_dafny.Zero).Minus(_1411_i0)), 5)
				(_nw224).ArraySet1(_1420_v6, 6)
				(_nw224).ArraySet1(_1420_v6, 7)
				(_nw224).ArraySet1(_1420_v6, 8)
				(_nw224).ArraySet1(_1420_v6, 9)
				(_nw224).ArraySet1(_1420_v6, 10)
				(_nw224).ArraySet1(func(_pat_let52_0 D1) D1 {
					return func(_1422_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let53_0 _dafny.Int) D1 {
							return func(_1423_dt__update_hcf1_h0 _dafny.Int) D1 {
								return Companion_D1_.Create_DC1_(_1423_dt__update_hcf1_h0)
							}(_pat_let53_0)
						}(_1411_i0)
					}(_pat_let52_0)
				}(_1420_v6), 11)
				(_nw224).ArraySet1(_1420_v6, 12)
				(_nw224).ArraySet1(Companion_D1_.Create_DC1_(_1411_i0), 13)
				(_nw224).ArraySet1((func() D1 {
					if p0 {
						return _1420_v6
					}
					return _1420_v6
				})(), 14)
				(_nw224).ArraySet1(_1420_v6, 15)
				(_nw224).ArraySet1(_1420_v6, 16)
				(_nw224).ArraySet1(_1420_v6, 17)
				(_nw224).ArraySet1(_1420_v6, 18)
				(_nw224).ArraySet1(_1420_v6, 19)
				(_nw224).ArraySet1(_1420_v6, 20)
				(_nw224).ArraySet1((func() D1 {
					if (_this).F9() {
						return _1420_v6
					}
					return _1420_v6
				})(), 21)
				(_nw224).ArraySet1(_1420_v6, 22)
				(_nw224).ArraySet1(_1420_v6, 23)
				(_nw224).ArraySet1(func(_pat_let54_0 D1) D1 {
					return func(_1424_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let55_0 _dafny.Int) D1 {
							return func(_1425_dt__update_hcf1_h1 _dafny.Int) D1 {
								return Companion_D1_.Create_DC1_(_1425_dt__update_hcf1_h1)
							}(_pat_let55_0)
						}(Companion_Default___.Fm0(_pat_let_tv71))
					}(_pat_let54_0)
				}(_1420_v6), 24)
				(_nw224).ArraySet1(_1420_v6, 25)
				(_nw224).ArraySet1(_1420_v6, 26)
				_1421_v7 = _nw224
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_1421_v7), 0))
				_ = _index215
				(_1421_v7).ArraySet1(_1420_v6, (_index215).Int())
				var _pat_let_tv72 = _1415_v1
				_ = _pat_let_tv72
				var _pat_let_tv73 = _1415_v1
				_ = _pat_let_tv73
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_1421_v7), 0))
				_ = _index216
				(_1421_v7).ArraySet1(func(_pat_let56_0 D1) D1 {
					return func(_1426_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let57_0 _dafny.Int) D1 {
							return func(_1427_dt__update_hcf1_h2 _dafny.Int) D1 {
								return Companion_D1_.Create_DC1_(_1427_dt__update_hcf1_h2)
							}(_pat_let57_0)
						}(Companion_Default___.SafeDivisionInt(_pat_let_tv72.F12, _pat_let_tv73.F12))
					}(_pat_let56_0)
				}(_1420_v6), (_index216).Int())
				var _1428_v8 _dafny.Sequence
				_ = _1428_v8
				_1428_v8 = _dafny.SeqOf((_this).F9())
				(_1415_v1).F12 = (_dafny.Zero).Minus(((_1411_i0).Times(_dafny.IntOfUint32((_1428_v8).Cardinality()))).Plus(_1411_i0))
				var _1429_v9 *C1
				_ = _1429_v9
				var _nw225 *C1 = New_C1_()
				_ = _nw225
				_nw225.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer77 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}((func(_1430_p0 bool, _1431_i0 _dafny.Int, _1432_v4 _dafny.Sequence) func(_dafny.Int) D1 {
					return func(_1433_i2 _dafny.Int) D1 {
						return Companion_D1_.Create_DC2_((_this).F10(), _1430_p0, _1431_i0, _dafny.IntOfUint32((_1432_v4).Cardinality()))
					}
				})(p0, _1411_i0, _1418_v4))))
				_1429_v9 = _nw225
				var _1434_v10 _dafny.Map
				_ = _1434_v10
				_1434_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1415_v1.F12, _1429_v9)
				_1429_v9 = (func() *C1 {
					if (_1434_v10).Contains((_this).F10()) {
						return (_1434_v10).Get((_this).F10()).(*C1)
					}
					return _1429_v9
				})()
			} else {
				var _1435_v11 _dafny.Sequence
				_ = _1435_v11
				_1435_v11 = _dafny.UnicodeSeqOfUtf8Bytes("l")
				var _1436_v12 _dafny.Set
				_ = _1436_v12
				_1436_v12 = Companion_Default___.Fm14(_1411_i0, _dafny.IntOfUint32((_1435_v11).Cardinality()), (_this).F10(), globalState)
				_1436_v12 = (func() _dafny.Set {
					if (_this).F9() {
						return _1436_v12
					}
					return _1436_v12
				})()
				var _1437_v13 _dafny.MultiSet
				_ = _1437_v13
				_1437_v13 = _dafny.MultiSetOf((_this).F9())
				var _1438_v14 _dafny.Sequence
				_ = _1438_v14
				_1438_v14 = _dafny.SeqOf(_1435_v11)
				var _1439_v15 _dafny.Sequence
				_ = _1439_v15
				_1439_v15 = _dafny.SeqOf((_this).F9(), (_this).F9())
				var _1440_v16 _dafny.Sequence
				_ = _1440_v16
				_1440_v16 = _dafny.SeqOf(_1411_i0, _1411_i0, _dafny.IntOfUint32((_1439_v15).Cardinality()), (_this).F10())
				var _1441_v17 _dafny.Sequence
				_ = _1441_v17
				_1441_v17 = _dafny.SeqOf(Companion_Default___.Fm0(globalState), _1411_i0, _dafny.IntOfUint32((_1440_v16).Cardinality()))
				var _1442_v18 _dafny.Map
				_ = _1442_v18
				_1442_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1435_v11).Cardinality()), _1441_v17)
				var _1443_v19 _dafny.MultiSet
				_ = _1443_v19
				_1443_v19 = _dafny.MultiSetOf(_1411_i0)
				var _1444_v20 _dafny.Map
				_ = _1444_v20
				_1444_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1435_v11).Cardinality()), _1411_i0)
				var _1445_v21 _dafny.Map
				_ = _1445_v21
				_1445_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.MultiSetOf((_this).F10(), _dafny.IntOfInt64(756)))
				var _1446_v22 _dafny.Set
				_ = _1446_v22
				_1446_v22 = _dafny.SetOf(_1443_v19, _dafny.MultiSetOf((_this).F10(), (_1444_v20).Cardinality()), (func() _dafny.MultiSet {
					if (_1445_v21).Contains(true) {
						return (_1445_v21).Get(true).(_dafny.MultiSet)
					}
					return _1443_v19
				})())
				var _1447_v23 _dafny.Array
				_ = _1447_v23
				var _nwElement0_46 _dafny.Int = (_this).Fm4(p0, p0, (_this).F9(), p0, globalState)
				_ = _nwElement0_46
				var _nw226 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(21))
				_ = _nw226
				(_nw226).ArraySet1(_nwElement0_46, 0)
				(_nw226).ArraySet1(((_1437_v13).Update(!((_this).F9()), Companion_Default___.Abs((_this).F10()))).Cardinality(), 1)
				(_nw226).ArraySet1((_this).F10(), 2)
				(_nw226).ArraySet1(Companion_Default___.SafeModuloInt(_1411_i0, (_this).F10()), 3)
				(_nw226).ArraySet1((_dafny.IntOfUint32((_1438_v14).Cardinality())).Plus((_dafny.Zero).Minus((_this).F10())), 4)
				(_nw226).ArraySet1((_this).F10(), 5)
				(_nw226).ArraySet1(_1411_i0, 6)
				(_nw226).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.MultiSetFromSeq(_dafny.SeqOf(p0))).Cardinality(), (_1442_v18).Cardinality()), 7)
				(_nw226).ArraySet1((_this).F10(), 8)
				(_nw226).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_1441_v17).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.IntOfUint32((_1441_v17).Cardinality()))).Uint32()).(_dafny.Int), (_this).F10())), 9)
				(_nw226).ArraySet1((_this).F10(), 10)
				(_nw226).ArraySet1(_1411_i0, 11)
				(_nw226).ArraySet1(_dafny.IntOfUint32((_1439_v15).Cardinality()), 12)
				(_nw226).ArraySet1((_this).Fm4((_this).F9(), p0, !(false), !(p0), globalState), 13)
				(_nw226).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(807), _dafny.IntOfInt64(9))), 14)
				(_nw226).ArraySet1((_this).F10(), 15)
				(_nw226).ArraySet1((_this).F10(), 16)
				(_nw226).ArraySet1((_this).F10(), 17)
				(_nw226).ArraySet1((_1446_v22).Cardinality(), 18)
				(_nw226).ArraySet1(_1411_i0, 19)
				(_nw226).ArraySet1(_1411_i0, 20)
				_1447_v23 = _nw226
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_1447_v23), 0))
				_ = _index217
				(_1447_v23).ArraySet1(_1411_i0, (_index217).Int())
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_1447_v23), 0))
				_ = _index218
				(_1447_v23).ArraySet1((_dafny.Zero).Minus((_this).F10()), (_index218).Int())
				var _1448_v24 _dafny.Array
				_ = _1448_v24
				var _nw227 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw227
				_1448_v24 = _nw227
				var _1449_v25 _dafny.Sequence
				_ = _1449_v25
				_1449_v25 = _dafny.SeqOf(_1448_v24)
				_1448_v24 = (_1449_v25).Select((Companion_Default___.SafeIndex((_1447_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_1447_v23), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1449_v25).Cardinality()))).Uint32()).(_dafny.Array)
				_1435_v11 = _1435_v11
				var _1450_v27 _dafny.MultiSet
				_ = _1450_v27
				_1450_v27 = _1443_v19
				_1444_v20 = func() _dafny.Map {
					var _coll31 = _dafny.NewMapBuilder()
					_ = _coll31
					for _iter35 := _dafny.Iterate((_1450_v27).Elements()); ; {
						_compr_31, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _1451_v26 _dafny.Int
						_1451_v26 = interface{}(_compr_31).(_dafny.Int)
						if (_1450_v27).Contains(_1451_v26) {
							_coll31.Add((_1451_v26).Minus(_1411_i0), (_1447_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_1447_v23), 0))).Int()).(_dafny.Int))
						}
					}
					return _coll31.ToMap()
				}()
			}
			var _1452_v28 _dafny.Array
			_ = _1452_v28
			var _nwElement0_47 _dafny.Int = (_1411_i0).Times(_1411_i0)
			_ = _nwElement0_47
			var _nw228 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(7))
			_ = _nw228
			(_nw228).ArraySet1(_nwElement0_47, 0)
			(_nw228).ArraySet1(_dafny.IntOfInt64(584), 1)
			(_nw228).ArraySet1((_this).F10(), 2)
			(_nw228).ArraySet1(_1411_i0, 3)
			(_nw228).ArraySet1((_this).F10(), 4)
			(_nw228).ArraySet1((_dafny.Zero).Minus((_this).F10()), 5)
			(_nw228).ArraySet1(((_this).F10()).Times(_1411_i0), 6)
			_1452_v28 = _nw228
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
			_ = _index219
			(_1452_v28).ArraySet1(_1411_i0, (_index219).Int())
			var _1453_v29 T1
			_ = _1453_v29
			var _nw229 *C2 = New_C2_()
			_ = _nw229
			_nw229.Ctor__(_1411_i0, p0)
			_1453_v29 = _nw229
			var _1454_v30 _dafny.Set
			_ = _1454_v30
			_1454_v30 = _dafny.SetOf((func() T1 {
				if p0 {
					return _1453_v29
				}
				return _1453_v29
			})(), _1453_v29, _1453_v29, _1453_v29, _1453_v29)
			var _1455_v31 _dafny.Int
			_ = _1455_v31
			_1455_v31 = _dafny.IntOfInt64(485)
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
			_ = _index220
			var _rhs266 _dafny.Int = ((_1455_v31).Plus((_this).F10())).Plus((_1455_v31).Plus(_dafny.IntOfInt64(114)))
			_ = _rhs266
			var _rhs267 _dafny.Set = (func() _dafny.Set {
				if true {
					return (_1454_v30).Union(_dafny.SetOf(_1453_v29))
				}
				return _1454_v30
			})()
			_ = _rhs267
			var _rhs268 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F10(), (_this).F10())
			_ = _rhs268
			var _lhs164 _dafny.Array = _1452_v28
			_ = _lhs164
			var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
			_ = _lhs165
			(_lhs164).ArraySet1(_rhs266, (_lhs165).Int())
			_1454_v30 = _rhs267
			_1455_v31 = _rhs268
			var _1456_v32 _dafny.Sequence
			_ = _1456_v32
			_1456_v32 = _dafny.UnicodeSeqOfUtf8Bytes("jk")
			var _1457_v33 _dafny.CodePoint
			_ = _1457_v33
			_1457_v33 = _dafny.CodePoint('q')
			var _1458_v34 _dafny.MultiSet
			_ = _1458_v34
			_1458_v34 = _dafny.MultiSetOf(_1457_v33, _1457_v33)
			var _1459_v35 _dafny.Sequence
			_ = _1459_v35
			_1459_v35 = _dafny.SeqOf((_1458_v34).Cardinality(), _1455_v31)
			var _1460_v36 _dafny.Set
			_ = _1460_v36
			_1460_v36 = _dafny.SetOf(_1459_v35)
			var _1461_v38 _dafny.Array
			_ = _1461_v38
			var _nw230 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
			_ = _nw230
			_1461_v38 = _nw230
			var _1462_v39 T2
			_ = _1462_v39
			var _nw231 *C3 = New_C3_()
			_ = _nw231
			_nw231.Ctor__(p0, _1461_v38, (_this).F9())
			_1462_v39 = _nw231
			var _1463_v40 _dafny.Sequence
			_ = _1463_v40
			_1463_v40 = _dafny.SeqOf(_1462_v39)
			var _1464_v41 _dafny.Sequence
			_ = _1464_v41
			_1464_v41 = _dafny.SeqOf(_1463_v40)
			var _1465_v42 _dafny.MultiSet
			_ = _1465_v42
			_1465_v42 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1464_v41).Cardinality()), (_this).F10())
			var _1466_v43 _dafny.Sequence
			_ = _1466_v43
			_1466_v43 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(96))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg78 _dafny.Int) interface{} {
					return coer78(arg78)
				}
			}(func(_1467_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("x"), (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("x")).Cardinality()))).Uint32(), _1457_v33), (func() _dafny.Sequence {
				if (_this).Fm5(_1460_v36, func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter36 := _dafny.Iterate((_1465_v42).Elements()); ; {
						_compr_32, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _1468_v37 _dafny.Int
						_1468_v37 = interface{}(_compr_32).(_dafny.Int)
						if (_1465_v42).Contains(_1468_v37) {
							_coll32.Add((_1468_v37).Times(_dafny.IntOfInt64(173)), _dafny.SeqOf((_1462_v39).F9()))
						}
					}
					return _coll32.ToMap()
				}(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1456_v32).Cardinality())), globalState) {
					return _1456_v32
				}
				return _1456_v32
			})())
			var _1469_v44 _dafny.Set
			_ = _1469_v44
			_1469_v44 = _dafny.SetOf(!((_this).F9()), (_1462_v39).F9(), (_1453_v29).F9())
			var _1470_v45 _dafny.Map
			_ = _1470_v45
			_1470_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1462_v39).F9(), _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("goklh"), _1457_v33))
			var _1471_v46 _dafny.Sequence
			_ = _1471_v46
			_1471_v46 = _dafny.SeqOf((_1453_v29).F9())
			var _rhs269 _dafny.Sequence = (_1466_v43).Select((Companion_Default___.SafeIndex((_1459_v35).Select((Companion_Default___.SafeIndex((_1469_v44).Cardinality(), _dafny.IntOfUint32((_1459_v35).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1466_v43).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs269
			var _rhs270 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-945))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg79 _dafny.Int) interface{} {
					return coer79(arg79)
				}
			}((func(_1472_v33 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1473_i4 _dafny.Int) _dafny.CodePoint {
					return _1472_v33
				}
			})(_1457_v33)))
			_ = _rhs270
			var _rhs271 bool = !((func() bool {
				if (_1470_v45).Contains((_1471_v46).Select((Companion_Default___.SafeIndex((_1452_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1471_v46).Cardinality()))).Uint32()).(bool)) {
					return (_1470_v45).Get((_1471_v46).Select((Companion_Default___.SafeIndex((_1452_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1471_v46).Cardinality()))).Uint32()).(bool)).(bool)
				}
				return (_this).F9()
			})())
			_ = _rhs271
			var _rhs272 bool = (_1462_v39).F9()
			_ = _rhs272
			var _lhs166 *GlobalState = globalState
			_ = _lhs166
			var _lhs167 *GlobalState = globalState
			_ = _lhs167
			_1456_v32 = _rhs269
			_1456_v32 = _rhs270
			_lhs166.F2 = _rhs271
			_lhs167.F2 = _rhs272
			var _1474_v47 _dafny.Map
			_ = _1474_v47
			_1474_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F10())
			var _1475_v48 D4
			_ = _1475_v48
			_1475_v48 = Companion_D4_.Create_DC12_(!((_this).F9()), (_1453_v29).F9(), (_1474_v47).Cardinality(), (_this).F10(), !((_1453_v29).F9()))
			var _1476_v49 D4
			_ = _1476_v49
			_1476_v49 = Companion_D4_.Create_DC15_(_1475_v48)
			var _source14 D4 = _1476_v49
			_ = _source14
			if _source14.Is_DC12() {
				var _1477___mcc_h0 bool = _source14.Get_().(D4_DC12).Cf20
				_ = _1477___mcc_h0
				var _1478___mcc_h1 bool = _source14.Get_().(D4_DC12).Cf21
				_ = _1478___mcc_h1
				var _1479___mcc_h2 _dafny.Int = _source14.Get_().(D4_DC12).Cf22
				_ = _1479___mcc_h2
				var _1480___mcc_h3 _dafny.Int = _source14.Get_().(D4_DC12).Cf23
				_ = _1480___mcc_h3
				var _1481___mcc_h4 bool = _source14.Get_().(D4_DC12).Cf24
				_ = _1481___mcc_h4
				var _1482_cf24 bool = _1481___mcc_h4
				_ = _1482_cf24
				var _1483_cf23 _dafny.Int = _1480___mcc_h3
				_ = _1483_cf23
				var _1484_cf22 _dafny.Int = _1479___mcc_h2
				_ = _1484_cf22
				var _1485_cf21 bool = _1478___mcc_h1
				_ = _1485_cf21
				var _1486_cf20 bool = _1477___mcc_h0
				_ = _1486_cf20
				var _1487_v50 _dafny.Array
				_ = _1487_v50
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_42
				var _nw232 _dafny.Array
				_ = _nw232
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw232 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) bool = func(_1488_i5 _dafny.Int) bool {
						return false
					}
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw232 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw232).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw232).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1487_v50 = _nw232
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1487_v50), 0))
				_ = _index221
				(_1487_v50).ArraySet1((_1453_v29).F9(), (_index221).Int())
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1487_v50), 0))
				_ = _index222
				(_1487_v50).ArraySet1((_1462_v39).F9(), (_index222).Int())
				_1483_cf23 = (_1459_v35).Select((Companion_Default___.SafeIndex((_1452_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1459_v35).Cardinality()))).Uint32()).(_dafny.Int)
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
				_ = _index223
				(_1452_v28).ArraySet1((_1459_v35).Select((Companion_Default___.SafeIndex(_1411_i0, _dafny.IntOfUint32((_1459_v35).Cardinality()))).Uint32()).(_dafny.Int), (_index223).Int())
				var _1489_v51 *C2
				_ = _1489_v51
				var _nw233 *C2 = New_C2_()
				_ = _nw233
				_nw233.Ctor__((_dafny.Zero).Minus((_1469_v44).Cardinality()), (_1487_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1487_v50), 0))).Int()).(bool))
				_1489_v51 = _nw233
			} else if _source14.Is_DC13() {
				var _1490___mcc_h5 _dafny.Int = _source14.Get_().(D4_DC13).Cf25
				_ = _1490___mcc_h5
				var _1491_cf25 _dafny.Int = _1490___mcc_h5
				_ = _1491_cf25
				var _1492_v52 _dafny.MultiSet
				_ = _1492_v52
				_1492_v52 = _dafny.MultiSetOf((_1471_v46).Select((Companion_Default___.SafeIndex((_1452_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1471_v46).Cardinality()))).Uint32()).(bool))
				_1492_v52 = (func() _dafny.MultiSet {
					if !((_1462_v39).F9()) {
						return _dafny.MultiSetFromSeq(_1471_v46)
					}
					return (_1492_v52).Update(p0, Companion_Default___.Abs((_1452_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))).Int()).(_dafny.Int)))
				})()
				var _1493_v53 _dafny.Array
				_ = _1493_v53
				var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw234
				_1493_v53 = _nw234
				var _1494_v54 D14
				_ = _1494_v54
				_1494_v54 = Companion_D14_.Create_DC38_(_1457_v33, (_1453_v29).F9(), _1471_v46)
				var _1495_v55 _dafny.Array
				_ = _1495_v55
				var _nwElement0_48 bool = !(p0)
				_ = _nwElement0_48
				var _nw235 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(15))
				_ = _nw235
				(_nw235).ArraySet1(_nwElement0_48, 0)
				(_nw235).ArraySet1((_1453_v29).F9(), 1)
				(_nw235).ArraySet1(p0, 2)
				(_nw235).ArraySet1((_this).F9(), 3)
				(_nw235).ArraySet1((_1453_v29).F9(), 4)
				(_nw235).ArraySet1((_this).F9(), 5)
				(_nw235).ArraySet1((_1494_v54).Dtor_cf68(), 6)
				(_nw235).ArraySet1((_1453_v29).F9(), 7)
				(_nw235).ArraySet1((_1462_v39).F9(), 8)
				(_nw235).ArraySet1(!((_1453_v29).F9()), 9)
				(_nw235).ArraySet1((_1462_v39).F9(), 10)
				(_nw235).ArraySet1((_1462_v39).F9(), 11)
				(_nw235).ArraySet1((_1453_v29).F9(), 12)
				(_nw235).ArraySet1((_1462_v39).F9(), 13)
				(_nw235).ArraySet1((_this).F9(), 14)
				_1495_v55 = _nw235
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_1493_v53), 0))
				_ = _index224
				(_1493_v53).ArraySet1(_1495_v55, (_index224).Int())
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_1493_v53), 0))
				_ = _index225
				(_1493_v53).ArraySet1(_1495_v55, (_index225).Int())
				var _1496_v56 *C2
				_ = _1496_v56
				var _nw236 *C2 = New_C2_()
				_ = _nw236
				_nw236.Ctor__((_this).F10(), (func() bool {
					if (_1453_v29).F9() {
						return (_1462_v39).F9()
					}
					return p0
				})())
				_1496_v56 = _nw236
				var _1497_v57 _dafny.Sequence
				_ = _1497_v57
				_1497_v57 = _dafny.SeqOf(_1496_v56, _1496_v56, _1496_v56, _1496_v56)
				_1496_v56 = (_1497_v57).Select((Companion_Default___.SafeIndex(_1455_v31, _dafny.IntOfUint32((_1497_v57).Cardinality()))).Uint32()).(*C2)
				var _1498_v58 _dafny.MultiSet
				_ = _1498_v58
				_1498_v58 = _dafny.MultiSetOf(_1469_v44)
				var _1499_v59 D3
				_ = _1499_v59
				_1499_v59 = Companion_D3_.Create_DC9_(true, _1471_v46)
				_1498_v58 = (Companion_Default___.Fm43(_dafny.SetOf(_dafny.IntOfUint32(((_1499_v59).Dtor_cf18()).Cardinality()), _1411_i0), (_1474_v47).Cardinality(), Companion_Default___.Fm0(globalState), globalState)).Update(_1469_v44, Companion_Default___.Abs(_1491_cf25))
			} else if _source14.Is_DC14() {
				var _1500___mcc_h6 bool = _source14.Get_().(D4_DC14).Cf26
				_ = _1500___mcc_h6
				var _1501___mcc_h7 bool = _source14.Get_().(D4_DC14).Cf27
				_ = _1501___mcc_h7
				var _1502___mcc_h8 _dafny.Int = _source14.Get_().(D4_DC14).Cf28
				_ = _1502___mcc_h8
				var _1503_cf28 _dafny.Int = _1502___mcc_h8
				_ = _1503_cf28
				var _1504_cf27 bool = _1501___mcc_h7
				_ = _1504_cf27
				var _1505_cf26 bool = _1500___mcc_h6
				_ = _1505_cf26
				var _1506_v60 _dafny.Array
				_ = _1506_v60
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_43
				var _nw237 _dafny.Array
				_ = _nw237
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw237 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.Sequence = (func(_1507_cf28 _dafny.Int, _1508_v33 _dafny.CodePoint, _1509_v32 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1510_i6 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("hksmc"), (Companion_Default___.SafeIndex(_1507_cf28, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hksmc")).Cardinality()))).Uint32(), _1508_v33), _1509_v32)
						}
					})(_1503_cf28, _1457_v33, _1456_v32)
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw237 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw237).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw237).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1506_v60 = _nw237
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
				_ = _index226
				var _rhs273 _dafny.Sequence = _1459_v35
				_ = _rhs273
				var _rhs274 _dafny.Int = Companion_Default___.Fm0(globalState)
				_ = _rhs274
				var _rhs275 _dafny.Array = _1506_v60
				_ = _rhs275
				var _lhs168 _dafny.Array = _1452_v28
				_ = _lhs168
				var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
				_ = _lhs169
				_1459_v35 = _rhs273
				(_lhs168).ArraySet1(_rhs274, (_lhs169).Int())
				_1506_v60 = _rhs275
				var _1511_v61 _dafny.Map
				_ = _1511_v61
				_1511_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1455_v31), (_1453_v29).F9())
				var _1512_v62 _dafny.Set
				_ = _1512_v62
				_1512_v62 = _dafny.SetOf((_this).F10(), (_this).F10())
				var _1513_v63 _dafny.Set
				_ = _1513_v63
				_1513_v63 = _1512_v62
				var _1514_v64 _dafny.Map
				_ = _1514_v64
				_1514_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1513_v63, _1503_cf28)
				var _1515_v65 _dafny.Set
				_ = _1515_v65
				_1515_v65 = _dafny.SetOf(_1514_v64)
				var _rhs276 bool = (func() bool {
					if false {
						return p0
					}
					return (func() bool {
						if (_1511_v61).Contains((_1452_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))).Int()).(_dafny.Int)) {
							return (_1511_v61).Get((_1452_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))).Int()).(_dafny.Int)).(bool)
						}
						return Companion_Default___.Fm1(_dafny.MultiSetOf((_dafny.Zero).Minus(_1503_cf28), (_this).F10()), _1503_cf28, globalState)
					})()
				})()
				_ = _rhs276
				var _rhs277 bool = !(((func() _dafny.Set {
					if !((_1462_v39).F9()) {
						return _1515_v65
					}
					return _1515_v65
				})()).IsProperSubsetOf(_1515_v65))
				_ = _rhs277
				var _lhs170 *GlobalState = globalState
				_ = _lhs170
				r0 = _rhs276
				_lhs170.F2 = _rhs277
				_1456_v32 = _1456_v32
				var _1516_v66 _dafny.Map
				_ = _1516_v66
				_1516_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_1455_v31), (_1453_v29).F9())
				var _1517_v67 D7
				_ = _1517_v67
				_1517_v67 = Companion_D7_.Create_DC18_(_1516_v66)
				var _pat_let_tv74 = _1516_v66
				_ = _pat_let_tv74
				var _1518_v68 _dafny.Array
				_ = _1518_v68
				var _nwElement0_49 D7 = _1517_v67
				_ = _nwElement0_49
				var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(9))
				_ = _nw238
				(_nw238).ArraySet1(_nwElement0_49, 0)
				(_nw238).ArraySet1(_1517_v67, 1)
				(_nw238).ArraySet1(_1517_v67, 2)
				(_nw238).ArraySet1(_1517_v67, 3)
				(_nw238).ArraySet1(Companion_D7_.Create_DC18_(_1516_v66), 4)
				(_nw238).ArraySet1(_1517_v67, 5)
				(_nw238).ArraySet1(_1517_v67, 6)
				(_nw238).ArraySet1(func(_pat_let58_0 D7) D7 {
					return func(_1519_dt__update__tmp_h3 D7) D7 {
						return func(_pat_let59_0 _dafny.Map) D7 {
							return func(_1520_dt__update_hcf32_h0 _dafny.Map) D7 {
								return Companion_D7_.Create_DC18_(_1520_dt__update_hcf32_h0)
							}(_pat_let59_0)
						}(_pat_let_tv74)
					}(_pat_let58_0)
				}(_1517_v67), 7)
				(_nw238).ArraySet1(Companion_D7_.Create_DC18_(_1516_v66), 8)
				_1518_v68 = _nw238
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1518_v68), 0))
				_ = _index227
				(_1518_v68).ArraySet1(_1517_v67, (_index227).Int())
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1518_v68), 0))
				_ = _index228
				(_1518_v68).ArraySet1(_1517_v67, (_index228).Int())
			} else if _source14.Is_DC11() {
				var _1521___mcc_h9 _dafny.Array = _source14.Get_().(D4_DC11).Cf19
				_ = _1521___mcc_h9
				var _1522_cf19 _dafny.Array = _1521___mcc_h9
				_ = _1522_cf19
				var _1523_v69 *C4
				_ = _1523_v69
				var _nw239 *C4 = New_C4_()
				_ = _nw239
				_nw239.Ctor__((_1453_v29).F9())
				_1523_v69 = _nw239
				var _1524_v70 _dafny.MultiSet
				_ = _1524_v70
				_1524_v70 = _dafny.MultiSetOf(Companion_Default___.Fm1(Companion_Default___.Fm29((_this).F9(), (_1462_v39).F9(), globalState), _1411_i0, globalState), (_1462_v39).F9(), (_1453_v29).F9())
				_1455_v31 = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_1524_v70).Contains((_1469_v44).IsDisjointFrom(_1469_v44)) {
						return (_1524_v70).Multiplicity((_1469_v44).IsDisjointFrom(_1469_v44))
					}
					return (_dafny.Zero).Minus(_1411_i0)
				})())
				_1457_v33 = _1457_v33
				_1455_v31 = Companion_Default___.Fm0(globalState)
			} else {
				var _1525___mcc_h10 D4 = _source14.Get_().(D4_DC15).Cf29
				_ = _1525___mcc_h10
				var _1526_cf29 D4 = _1525___mcc_h10
				_ = _1526_cf29
				var _1527_v71 *C2
				_ = _1527_v71
				var _nw240 *C2 = New_C2_()
				_ = _nw240
				_nw240.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xf")).Cardinality()), (_this).F9())
				_1527_v71 = _nw240
				var _1528_v72 _dafny.Sequence
				_ = _1528_v72
				_1528_v72 = _dafny.SeqOf(_dafny.SeqOf((_1462_v39).F9(), true))
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
				_ = _index229
				(_1452_v28).ArraySet1(_dafny.IntOfUint32((_1528_v72).Cardinality()), (_index229).Int())
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
				_ = _index230
				(_1452_v28).ArraySet1(((_1474_v47).Cardinality()).Plus((_dafny.Zero).Minus((_1527_v71).F14())), (_index230).Int())
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1452_v28), 0))
				_ = _index231
				(_1452_v28).ArraySet1(_1411_i0, (_index231).Int())
			}
		}
		var _1529_v73 D1
		_ = _1529_v73
		_1529_v73 = Companion_D1_.Create_DC4_((_this).F10(), !(p0), p0, (_this).F9(), false)
		var _1530_v74 *C1
		_ = _1530_v74
		var _nw241 *C1 = New_C1_()
		_ = _nw241
		_nw241.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(207))).Uint32(), func(coer80 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
			return func(arg80 _dafny.Int) interface{} {
				return coer80(arg80)
			}
		}(func(_1531_i7 _dafny.Int) D1 {
			return Companion_D1_.Create_DC2_(_dafny.IntOfInt64(828), (_this).F9(), (_this).F10(), (_this).F10())
		})), Companion_Default___.Fm7((_this).F10(), (_this).F9(), _1529_v73, true, globalState)))
		_1530_v74 = _nw241
		_1530_v74 = _1530_v74
		var _1532_v75 _dafny.Array
		_ = _1532_v75
		var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(18))
		_ = _nw242
		_1532_v75 = _nw242
		var _1533_v76 *C3
		_ = _1533_v76
		var _nw243 *C3 = New_C3_()
		_ = _nw243
		_nw243.Ctor__(p0, _1532_v75, true)
		_1533_v76 = _nw243
		var _1534_v78 _dafny.Map
		_ = _1534_v78
		_1534_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p0)
		var _1535_v79 _dafny.Sequence
		_ = _1535_v79
		_1535_v79 = _dafny.UnicodeSeqOfUtf8Bytes("urpoc")
		var _1536_v81 _dafny.Sequence
		_ = _1536_v81
		_1536_v81 = _dafny.SeqOf(_1534_v78, _1534_v78)
		var _1537_v82 _dafny.Array
		_ = _1537_v82
		var _nwElement0_50 _dafny.Map = (func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(294), _dafny.IntOfInt64(229))); ; {
				_compr_33, _ok37 := _iter37()
				if !_ok37 {
					break
				}
				var _1538_v77 _dafny.Int
				_1538_v77 = interface{}(_compr_33).(_dafny.Int)
				if ((_dafny.IntOfInt64(294)).Cmp(_1538_v77) <= 0) && ((_1538_v77).Cmp(_dafny.IntOfInt64(229)) < 0) {
					_coll33.Add(Companion_Default___.SafeModuloInt(_1538_v77, (_this).F10()), p0)
				}
			}
			return _coll33.ToMap()
		}()).Update((_dafny.Zero).Minus((_this).F10()), (_1533_v76).F15())
		_ = _nwElement0_50
		var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(27))
		_ = _nw244
		(_nw244).ArraySet1(_nwElement0_50, 0)
		(_nw244).ArraySet1((Companion_Default___.Fm15((_1533_v76).F15(), globalState)).Merge(_1534_v78), 1)
		(_nw244).ArraySet1(_1534_v78, 2)
		(_nw244).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_1533_v76).F15()), 3)
		(_nw244).ArraySet1(_1534_v78, 4)
		(_nw244).ArraySet1(_1534_v78, 5)
		(_nw244).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p0), 6)
		(_nw244).ArraySet1(_1534_v78, 7)
		(_nw244).ArraySet1(_1534_v78, 8)
		(_nw244).ArraySet1(((_1534_v78).Update((_this).F10(), !((_1533_v76).F15()))).Merge(_1534_v78), 9)
		(_nw244).ArraySet1(_1534_v78, 10)
		(_nw244).ArraySet1(_1534_v78, 11)
		(_nw244).ArraySet1(_1534_v78, 12)
		(_nw244).ArraySet1(_1534_v78, 13)
		(_nw244).ArraySet1(_1534_v78, 14)
		(_nw244).ArraySet1(_1534_v78, 15)
		(_nw244).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1535_v79).Cardinality()), p0), 16)
		(_nw244).ArraySet1(_1534_v78, 17)
		(_nw244).ArraySet1(_1534_v78, 18)
		(_nw244).ArraySet1((_1534_v78).Merge(_1534_v78), 19)
		(_nw244).ArraySet1(Companion_Default___.Fm15((_this).F9(), globalState), 20)
		(_nw244).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p0), 21)
		(_nw244).ArraySet1(_1534_v78, 22)
		(_nw244).ArraySet1(func() _dafny.Map {
			var _coll34 = _dafny.NewMapBuilder()
			_ = _coll34
			for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(670), _dafny.IntOfInt64(-668))); ; {
				_compr_34, _ok38 := _iter38()
				if !_ok38 {
					break
				}
				var _1539_v80 _dafny.Int
				_1539_v80 = interface{}(_compr_34).(_dafny.Int)
				if ((_dafny.IntOfInt64(670)).Cmp(_1539_v80) <= 0) && ((_1539_v80).Cmp(_dafny.IntOfInt64(-668)) < 0) {
					_coll34.Add((_1539_v80).Times((_this).F10()), (_this).F9())
				}
			}
			return _coll34.ToMap()
		}(), 23)
		(_nw244).ArraySet1((Companion_Default___.Fm15(p0, globalState)).Merge((_1536_v81).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1536_v81).Cardinality()))).Uint32()).(_dafny.Map)), 24)
		(_nw244).ArraySet1(_1534_v78, 25)
		(_nw244).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p0)).Merge(_1534_v78), 26)
		_1537_v82 = _nw244
		var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_1537_v82), 0))
		_ = _index232
		(_1537_v82).ArraySet1(_1534_v78, (_index232).Int())
		var _1540_v83 D1
		_ = _1540_v83
		_1540_v83 = Companion_D1_.Create_DC2_((_this).F10(), true, (_this).F10(), (_this).F10())
		var _pat_let_tv75 = _1534_v78
		_ = _pat_let_tv75
		var _pat_let_tv76 = _1534_v78
		_ = _pat_let_tv76
		var _pat_let_tv77 = _1536_v81
		_ = _pat_let_tv77
		var _pat_let_tv78 = _1536_v81
		_ = _pat_let_tv78
		var _pat_let_tv79 = _1534_v78
		_ = _pat_let_tv79
		var _pat_let_tv80 = _1534_v78
		_ = _pat_let_tv80
		var _pat_let_tv81 = _1534_v78
		_ = _pat_let_tv81
		var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_1537_v82), 0))
		_ = _index233
		(_1537_v82).ArraySet1(func(_source15 D1) _dafny.Map {
			if _source15.Is_DC2() {
				var _1541___mcc_h11 _dafny.Int = _source15.Get_().(D1_DC2).Cf2
				_ = _1541___mcc_h11
				var _1542___mcc_h12 bool = _source15.Get_().(D1_DC2).Cf3
				_ = _1542___mcc_h12
				var _1543___mcc_h13 _dafny.Int = _source15.Get_().(D1_DC2).Cf4
				_ = _1543___mcc_h13
				var _1544___mcc_h14 _dafny.Int = _source15.Get_().(D1_DC2).Cf5
				_ = _1544___mcc_h14
				var _1545_cf5 _dafny.Int = _1544___mcc_h14
				_ = _1545_cf5
				var _1546_cf4 _dafny.Int = _1543___mcc_h13
				_ = _1546_cf4
				var _1547_cf3 bool = _1542___mcc_h12
				_ = _1547_cf3
				var _1548_cf2 _dafny.Int = _1541___mcc_h11
				_ = _1548_cf2
				return (_pat_let_tv75).Merge(_pat_let_tv76)
			} else if _source15.Is_DC3() {
				var _1549___mcc_h15 bool = _source15.Get_().(D1_DC3).Cf6
				_ = _1549___mcc_h15
				var _1550___mcc_h16 _dafny.Sequence = _source15.Get_().(D1_DC3).Cf7
				_ = _1550___mcc_h16
				var _1551___mcc_h17 _dafny.Int = _source15.Get_().(D1_DC3).Cf8
				_ = _1551___mcc_h17
				var _1552_cf8 _dafny.Int = _1551___mcc_h17
				_ = _1552_cf8
				var _1553_cf7 _dafny.Sequence = _1550___mcc_h16
				_ = _1553_cf7
				var _1554_cf6 bool = _1549___mcc_h15
				_ = _1554_cf6
				return (_pat_let_tv77).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_pat_let_tv78).Cardinality()))).Uint32()).(_dafny.Map)
			} else if _source15.Is_DC4() {
				var _1555___mcc_h18 _dafny.Int = _source15.Get_().(D1_DC4).Cf9
				_ = _1555___mcc_h18
				var _1556___mcc_h19 bool = _source15.Get_().(D1_DC4).Cf10
				_ = _1556___mcc_h19
				var _1557___mcc_h20 bool = _source15.Get_().(D1_DC4).Cf11
				_ = _1557___mcc_h20
				var _1558___mcc_h21 bool = _source15.Get_().(D1_DC4).Cf12
				_ = _1558___mcc_h21
				var _1559___mcc_h22 bool = _source15.Get_().(D1_DC4).Cf13
				_ = _1559___mcc_h22
				var _1560_cf13 bool = _1559___mcc_h22
				_ = _1560_cf13
				var _1561_cf12 bool = _1558___mcc_h21
				_ = _1561_cf12
				var _1562_cf11 bool = _1557___mcc_h20
				_ = _1562_cf11
				var _1563_cf10 bool = _1556___mcc_h19
				_ = _1563_cf10
				var _1564_cf9 _dafny.Int = _1555___mcc_h18
				_ = _1564_cf9
				return _pat_let_tv79
			} else if _source15.Is_DC1() {
				var _1565___mcc_h23 _dafny.Int = _source15.Get_().(D1_DC1).Cf1
				_ = _1565___mcc_h23
				var _1566_cf1 _dafny.Int = _1565___mcc_h23
				_ = _1566_cf1
				return _pat_let_tv80
			} else {
				var _1567___mcc_h24 D1 = _source15.Get_().(D1_DC5).Cf14
				_ = _1567___mcc_h24
				var _1568_cf14 D1 = _1567___mcc_h24
				_ = _1568_cf14
				return _pat_let_tv81
			}
		}(_1540_v83), (_index233).Int())
		var _1569_v84 _dafny.Array
		_ = _1569_v84
		var _len0_44 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_44
		var _nw245 _dafny.Array
		_ = _nw245
		if _len0_44.Cmp(_dafny.Zero) == 0 {
			_nw245 = _dafny.NewArray(_len0_44)
		} else {
			var _init44 func(_dafny.Int) bool = func(_1570_i8 _dafny.Int) bool {
				return (_this).F9()
			}
			_ = _init44
			var _element0_44 = _init44(_dafny.Zero)
			_ = _element0_44
			_nw245 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
			(_nw245).ArraySet1(_element0_44, 0)
			var _nativeLen0_44 = (_len0_44).Int()
			_ = _nativeLen0_44
			for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
				(_nw245).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
			}
		}
		_1569_v84 = _nw245
		var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1569_v84), 0))
		_ = _index234
		(_1569_v84).ArraySet1(true, (_index234).Int())
		var _1571_v85 _dafny.Sequence
		_ = _1571_v85
		_1571_v85 = _dafny.SeqOf(p0)
		var _1572_v86 _dafny.MultiSet
		_ = _1572_v86
		_1572_v86 = _dafny.MultiSetOf(p0, (_1533_v76).F15())
		var _1573_v87 _dafny.MultiSet
		_ = _1573_v87
		_1573_v87 = _dafny.MultiSetOf((_this).F10(), (_1572_v86).Cardinality(), _dafny.IntOfInt64(771), (_this).F10(), (_this).F10())
		var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1569_v84), 0))
		_ = _index235
		(_1569_v84).ArraySet1((func() bool {
			if true {
				return (_1571_v85).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1573_v87).Contains((_this).F10()) {
						return (_1573_v87).Multiplicity((_this).F10())
					}
					return (_this).F10()
				})(), _dafny.IntOfUint32((_1571_v85).Cardinality()))).Uint32()).(bool)
			}
			return false
		})(), (_index235).Int())
		if ((_this).F10()).Cmp((_this).F10()) >= 0 {
			var _1574_v88 _dafny.Sequence
			_ = _1574_v88
			_1574_v88 = _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_1535_v79).Cardinality()), (_this).F10()))
			var _1575_v89 _dafny.Map
			_ = _1575_v89
			_1575_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1533_v76).Fm32(_1574_v88, !((_this).F9()), (_this).F10(), false, globalState), _1535_v79)
			var _1576_v90 _dafny.Sequence
			_ = _1576_v90
			_1576_v90 = _dafny.SeqOf((_this).F10(), _dafny.IntOfInt64(668))
			var _1577_v91 _dafny.CodePoint
			_ = _1577_v91
			_1577_v91 = _dafny.CodePoint('e')
			_1575_v89 = (_1575_v89).Update(_dafny.Companion_Sequence_.IsPrefixOf(_1535_v79, _1535_v79), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("cgiaxc"), (Companion_Default___.SafeIndex((_1576_v90).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1576_v90).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cgiaxc")).Cardinality()))).Uint32(), _1577_v91), _dafny.UnicodeSeqOfUtf8Bytes("ehohxsxbk")))
			var _1578_v92 _dafny.Int
			_ = _1578_v92
			_1578_v92 = _dafny.IntOfInt64(256)
			_1578_v92 = (_this).F10()
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1569_v84), 0))
			_ = _index236
			(_1569_v84).ArraySet1((_this).F9(), (_index236).Int())
			_1578_v92 = _1578_v92
			var _1579_v93 _dafny.Array
			_ = _1579_v93
			var _len0_45 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_45
			var _nw246 _dafny.Array
			_ = _nw246
			if _len0_45.Cmp(_dafny.Zero) == 0 {
				_nw246 = _dafny.NewArray(_len0_45)
			} else {
				var _init45 func(_dafny.Int) _dafny.Sequence = (func(_1580_v79 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1581_i9 _dafny.Int) _dafny.Sequence {
						return _1580_v79
					}
				})(_1535_v79)
				_ = _init45
				var _element0_45 = _init45(_dafny.Zero)
				_ = _element0_45
				_nw246 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
				(_nw246).ArraySet1(_element0_45, 0)
				var _nativeLen0_45 = (_len0_45).Int()
				_ = _nativeLen0_45
				for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
					(_nw246).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
				}
			}
			_1579_v93 = _nw246
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_1579_v93), 0))
			_ = _index237
			(_1579_v93).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1535_v79, (Companion_Default___.SafeIndex(_1578_v92, _dafny.IntOfUint32((_1535_v79).Cardinality()))).Uint32(), (Companion_D13_.Create_DC35_((_1533_v76).F15(), _1577_v91, _1530_v74, (_1533_v76).F15())).Dtor_cf59()), (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1535_v79, (Companion_Default___.SafeIndex(_1578_v92, _dafny.IntOfUint32((_1535_v79).Cardinality()))).Uint32(), (Companion_D13_.Create_DC35_((_1533_v76).F15(), _1577_v91, _1530_v74, (_1533_v76).F15())).Dtor_cf59())).Cardinality()))).Uint32(), _1577_v91), (_index237).Int())
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_1579_v93), 0))
			_ = _index238
			(_1579_v93).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1535_v79, _dafny.Companion_Sequence_.Update(_1535_v79, (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1535_v79).Cardinality()))).Uint32(), _1577_v91)), (_index238).Int())
		} else {
			var _1582_v94 *C4
			_ = _1582_v94
			var _nw247 *C4 = New_C4_()
			_ = _nw247
			_nw247.Ctor__(true)
			_1582_v94 = _nw247
			var _1583_v95 _dafny.Array
			_ = _1583_v95
			var _nw248 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw248
			_1583_v95 = _nw248
			(_1533_v76).M1((_dafny.Zero).Minus(_dafny.IntOfInt64(-225)), p0, p0, _1583_v95, globalState)
			var _1584_v96 _dafny.Sequence
			_ = _1584_v96
			_1584_v96 = _dafny.SeqOf(Companion_D4_.Create_DC13_(_dafny.IntOfInt64(923)))
			var _1585_v97 D4
			_ = _1585_v97
			_1585_v97 = Companion_D4_.Create_DC15_((_1584_v96).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1584_v96).Cardinality()))).Uint32()).(D4))
			var _1586_v98 _dafny.Set
			_ = _1586_v98
			_1586_v98 = _dafny.SetOf(Companion_D4_.Create_DC15_(Companion_D4_.Create_DC13_(_dafny.IntOfInt64(863))), _1585_v97, _1585_v97)
			_1586_v98 = _1586_v98
			var _1587_v99 _dafny.Sequence
			_ = _1587_v99
			_1587_v99 = _dafny.SeqOf((_this).F10())
			var _1588_v100 _dafny.Set
			_ = _1588_v100
			_1588_v100 = _dafny.SetOf(_1587_v99)
			var _1589_v101 *C0
			_ = _1589_v101
			var _nw249 *C0 = New_C0_()
			_ = _nw249
			_nw249.Ctor__((_1588_v100).Intersection(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(905))).Uint32(), func(coer81 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg81 _dafny.Int) interface{} {
					return coer81(arg81)
				}
			}(func(_1590_i10 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-421)
			})), _dafny.SeqOf((_this).F10()))), (_this).F10())
			_1589_v101 = _nw249
			var _source16 D7 = Companion_Default___.Fm44(globalState)
			_ = _source16
			if _source16.Is_DC19() {
				var _1591___mcc_h25 _dafny.Array = _source16.Get_().(D7_DC19).Cf33
				_ = _1591___mcc_h25
				var _1592___mcc_h26 bool = _source16.Get_().(D7_DC19).Cf34
				_ = _1592___mcc_h26
				var _1593___mcc_h27 _dafny.Sequence = _source16.Get_().(D7_DC19).Cf35
				_ = _1593___mcc_h27
				var _1594_cf35 _dafny.Sequence = _1593___mcc_h27
				_ = _1594_cf35
				var _1595_cf34 bool = _1592___mcc_h26
				_ = _1595_cf34
				var _1596_cf33 _dafny.Array = _1591___mcc_h25
				_ = _1596_cf33
				(_1589_v101).F12 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if _1595_cf34 {
						return (_this).F10()
					}
					return _dafny.IntOfUint32((_1571_v85).Cardinality())
				})(), _1589_v101.F12)
				var _1597_v102 *C3
				_ = _1597_v102
				var _nw250 *C3 = New_C3_()
				_ = _nw250
				_nw250.Ctor__(true, _1532_v75, (_1533_v76).F15())
				_1597_v102 = _nw250
				var _1598_v103 _dafny.Sequence
				_ = _1598_v103
				_1598_v103 = _dafny.SeqOf(_1569_v84, _1569_v84, _1569_v84, _1569_v84, _1569_v84)
				_1569_v84 = (_1598_v103).Select((Companion_Default___.SafeIndex(_1589_v101.F12, _dafny.IntOfUint32((_1598_v103).Cardinality()))).Uint32()).(_dafny.Array)
				var _1599_v104 _dafny.Map
				_ = _1599_v104
				_1599_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality())), (_1572_v86).Update(false, Companion_Default___.Abs((_dafny.Zero).Minus(_1589_v101.F12))))
				var _pat_let_tv82 = _1589_v101
				_ = _pat_let_tv82
				_1599_v104 = (_1599_v104).Update(func(_pat_let60_0 D1) D1 {
					return func(_1600_dt__update__tmp_h4 D1) D1 {
						return func(_pat_let61_0 _dafny.Int) D1 {
							return func(_1601_dt__update_hcf1_h3 _dafny.Int) D1 {
								return Companion_D1_.Create_DC1_(_1601_dt__update_hcf1_h3)
							}(_pat_let61_0)
						}(_pat_let_tv82.F12)
					}(_pat_let60_0)
				}(Companion_D1_.Create_DC1_((func() _dafny.Int {
					if (_1572_v86).Contains((_1533_v76).F15()) {
						return (_1572_v86).Multiplicity((_1533_v76).F15())
					}
					return _1589_v101.F12
				})())), _1572_v86)
			} else if _source16.Is_DC20() {
				var _1602___mcc_h28 bool = _source16.Get_().(D7_DC20).Cf36
				_ = _1602___mcc_h28
				var _1603___mcc_h29 _dafny.CodePoint = _source16.Get_().(D7_DC20).Cf37
				_ = _1603___mcc_h29
				var _1604_cf37 _dafny.CodePoint = _1603___mcc_h29
				_ = _1604_cf37
				var _1605_cf36 bool = _1602___mcc_h28
				_ = _1605_cf36
				var _1606_v105 _dafny.Map
				_ = _1606_v105
				_1606_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(globalState), _1589_v101.F12)
				_1605_cf36 = !(!(_1606_v105).Contains((_this).F10()))
				var _1607_v106 _dafny.Map
				_ = _1607_v106
				_1607_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus((_dafny.IntOfUint32((_1587_v99).Cardinality())).Minus(_1589_v101.F12)))
				_1607_v106 = (_1607_v106).Update(!(!((_1569_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1569_v84), 0))).Int()).(bool))), _dafny.IntOfInt64(414))
				var _1608_v107 _dafny.Map
				_ = _1608_v107
				_1608_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1533_v76).F15(), _1605_cf36)
				var _1609_v108 _dafny.Set
				_ = _1609_v108
				_1609_v108 = _dafny.SetOf(p0, p0, p0)
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1569_v84), 0))
				_ = _index239
				var _rhs278 _dafny.Int = (_1589_v101.F12).Minus((_this).F10())
				_ = _rhs278
				var _rhs279 bool = (_1572_v86).IsSubsetOf(_dafny.MultiSetOf(_1605_cf36))
				_ = _rhs279
				var _rhs280 _dafny.MultiSet = _1572_v86
				_ = _rhs280
				var _rhs281 bool = (func() bool {
					if (_1608_v107).Contains((_dafny.SetOf((_1533_v76).F15())).IsDisjointFrom(_1609_v108)) {
						return (_1608_v107).Get((_dafny.SetOf((_1533_v76).F15())).IsDisjointFrom(_1609_v108)).(bool)
					}
					return false
				})()
				_ = _rhs281
				var _lhs171 *C0 = _1589_v101
				_ = _lhs171
				var _lhs172 *GlobalState = globalState
				_ = _lhs172
				var _lhs173 _dafny.Array = _1569_v84
				_ = _lhs173
				var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1569_v84), 0))
				_ = _lhs174
				_lhs171.F12 = _rhs278
				_lhs172.F2 = _rhs279
				_1572_v86 = _rhs280
				(_lhs173).ArraySet1(_rhs281, (_lhs174).Int())
				(_1589_v101).F12 = Companion_Default___.Fm0(globalState)
			} else if _source16.Is_DC18() {
				var _1610___mcc_h30 _dafny.Map = _source16.Get_().(D7_DC18).Cf32
				_ = _1610___mcc_h30
				var _1611_cf32 _dafny.Map = _1610___mcc_h30
				_ = _1611_cf32
				var _1612_v109 D15
				_ = _1612_v109
				_1612_v109 = Companion_D15_.Create_DC40_(_dafny.SetOf((_1533_v76).F15()))
				var _1613_v110 _dafny.Sequence
				_ = _1613_v110
				_1613_v110 = _dafny.SeqOf((_1612_v109).Dtor_cf72(), (_1612_v109).Dtor_cf72())
				var _1614_v111 _dafny.CodePoint
				_ = _1614_v111
				_1614_v111 = _dafny.CodePoint('s')
				var _1615_v112 _dafny.Map
				_ = _1615_v112
				_1615_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1589_v101.F12, _1614_v111)
				var _1616_v113 _dafny.Map
				_ = _1616_v113
				_1616_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_1613_v110).Cardinality())).Plus(_dafny.IntOfUint32((_1535_v79).Cardinality())), _1615_v112)
				_1616_v113 = (_1616_v113).Update(_1589_v101.F12, _1615_v112)
				var _1617_v114 *C1
				_ = _1617_v114
				var _nw251 *C1 = New_C1_()
				_ = _nw251
				_nw251.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1530_v74.F13, _1530_v74.F13))
				_1617_v114 = _nw251
				var _1618_v116 D14
				_ = _1618_v116
				_1618_v116 = Companion_D14_.Create_DC38_(_1614_v111, (_1571_v85).Select((Companion_Default___.SafeIndex((func() _dafny.Map {
					var _coll35 = _dafny.NewMapBuilder()
					_ = _coll35
					for _iter39 := _dafny.Iterate((Companion_Default___.Fm45(globalState)).Elements()); ; {
						_compr_35, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _1619_v115 D1
						_1619_v115 = interface{}(_compr_35).(D1)
						if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm45(globalState), _1619_v115) {
							_coll35.Add(_1619_v115, (_this).F9())
						}
					}
					return _coll35.ToMap()
				}()).Cardinality(), _dafny.IntOfUint32((_1571_v85).Cardinality()))).Uint32()).(bool), _1571_v85)
				(globalState).F2 = (_1618_v116).Dtor_cf68()
				var _1620_v117 _dafny.Array
				_ = _1620_v117
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_46
				var _nw252 _dafny.Array
				_ = _nw252
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw252 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Int = (func(_1621_v101 *C0) func(_dafny.Int) _dafny.Int {
						return func(_1622_i11 _dafny.Int) _dafny.Int {
							return (_1622_i11).Plus(_1621_v101.F12)
						}
					})(_1589_v101)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw252 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw252).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw252).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1620_v117 = _nw252
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1620_v117), 0))
				_ = _index240
				(_1620_v117).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1535_v79, _1535_v79)).Cardinality()), (_index240).Int())
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1620_v117), 0))
				_ = _index241
				(_1620_v117).ArraySet1((_this).F10(), (_index241).Int())
			} else {
				var _1623___mcc_h31 D7 = _source16.Get_().(D7_DC21).Cf38
				_ = _1623___mcc_h31
				var _1624_cf38 D7 = _1623___mcc_h31
				_ = _1624_cf38
				(globalState).F2 = (_1569_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1569_v84), 0))).Int()).(bool)
				(_1589_v101).F12 = (_this).F10()
				(globalState).F2 = (_1533_v76).F15()
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1583_v95), 0))
				_ = _index242
				(_1583_v95).ArraySet1(_1571_v85, (_index242).Int())
				var _1625_v118 _dafny.Map
				_ = _1625_v118
				_1625_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.IntOfInt64(653))
				var _1626_v119 _dafny.Map
				_ = _1626_v119
				_1626_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1533_v76, _1535_v79)
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_1583_v95), 0))
				_ = _index243
				(_1583_v95).ArraySet1(Companion_Default___.Fm41(Companion_D10_.Create_DC27_(_1625_v118), (_this).F10(), _dafny.Companion_Sequence_.Update(_1587_v99, (Companion_Default___.SafeIndex(_1589_v101.F12, _dafny.IntOfUint32((_1587_v99).Cardinality()))).Uint32(), (_1626_v119).Cardinality()), globalState), (_index243).Int())
			}
		}
		r0 = (_1569_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_1569_v84), 0))).Int()).(bool)
		return r0
	}
}
func (_this *C5) M4(p0 _dafny.Array, p1 _dafny.Int, p2 bool, globalState *GlobalState) {
	{
		var _1627_i0 _dafny.Int
		_ = _1627_i0
		_1627_i0 = _dafny.Zero
		{
			for !((_this).F9()) {
				{
					if (_1627_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1627_i0 = (_1627_i0).Plus(_dafny.One)
					var _1628_v0 _dafny.Int
					_ = _1628_v0
					_1628_v0 = _dafny.IntOfInt64(-100)
					var _1629_v1 _dafny.CodePoint
					_ = _1629_v1
					_1629_v1 = _dafny.CodePoint('r')
					var _1630_v2 _dafny.Sequence
					_ = _1630_v2
					_1630_v2 = _dafny.SeqOf(_1629_v1)
					var _1631_v3 _dafny.MultiSet
					_ = _1631_v3
					_1631_v3 = _dafny.MultiSetOf(_1629_v1, _1629_v1)
					var _1632_v4 _dafny.Map
					_ = _1632_v4
					_1632_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_dafny.MultiSetFromSeq(_1630_v2)).IsDisjointFrom(_1631_v3))
					_1628_v0 = (_1632_v4).Cardinality()
					var _1633_v5 _dafny.Array
					_ = _1633_v5
					var _nw253 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
					_ = _nw253
					_1633_v5 = _nw253
					var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_1633_v5), 0))
					_ = _index244
					(_1633_v5).ArraySet1(p2, (_index244).Int())
					var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_1633_v5), 0))
					_ = _index245
					(_1633_v5).ArraySet1((_this).F9(), (_index245).Int())
					_1628_v0 = Companion_Default___.SafeDivisionInt((_this).F10(), _1628_v0)
					var _1634_v6 _dafny.Sequence
					_ = _1634_v6
					_1634_v6 = _dafny.SeqOf((_1633_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_1633_v5), 0))).Int()).(bool))
					var _1635_v7 _dafny.Map
					_ = _1635_v7
					_1635_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.IntOfUint32((_1634_v6).Cardinality()))
					var _1636_v8 _dafny.Set
					_ = _1636_v8
					_1636_v8 = _dafny.SetOf((_dafny.Zero).Minus((_1635_v7).Cardinality()), (_this).F10(), _1628_v0, _dafny.IntOfInt64(-171))
					var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_1633_v5), 0))
					_ = _index246
					var _rhs282 bool = true
					_ = _rhs282
					var _rhs283 _dafny.Int = (Companion_Default___.SafeModuloInt(_1628_v0, _1628_v0)).Plus(((_dafny.SetOf((_this).F10(), _dafny.IntOfInt64(-206))).Difference(_1636_v8)).Cardinality())
					_ = _rhs283
					var _rhs284 bool = false
					_ = _rhs284
					var _rhs285 bool = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1628_v0, (_this).F9())).Contains(_1628_v0)
					_ = _rhs285
					var _rhs286 _dafny.Int = (_this).F10()
					_ = _rhs286
					var _lhs175 _dafny.Array = _1633_v5
					_ = _lhs175
					var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_1633_v5), 0))
					_ = _lhs176
					var _lhs177 *GlobalState = globalState
					_ = _lhs177
					var _lhs178 *GlobalState = globalState
					_ = _lhs178
					(_lhs175).ArraySet1(_rhs282, (_lhs176).Int())
					_1628_v0 = _rhs283
					_lhs177.F2 = _rhs284
					_lhs178.F2 = _rhs285
					_1628_v0 = _rhs286
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1637_v9 _dafny.Sequence
		_ = _1637_v9
		_1637_v9 = _dafny.UnicodeSeqOfUtf8Bytes("ua")
		_1637_v9 = _dafny.UnicodeSeqOfUtf8Bytes("iatft")
		var _1638_v10 _dafny.Int
		_ = _1638_v10
		_1638_v10 = _dafny.IntOfInt64(415)
		var _1639_v11 _dafny.Array
		_ = _1639_v11
		var _len0_47 _dafny.Int = _dafny.One
		_ = _len0_47
		var _nw254 _dafny.Array
		_ = _nw254
		if _len0_47.Cmp(_dafny.Zero) == 0 {
			_nw254 = _dafny.NewArray(_len0_47)
		} else {
			var _init47 func(_dafny.Int) bool = func(_1640_i1 _dafny.Int) bool {
				return (_this).F9()
			}
			_ = _init47
			var _element0_47 = _init47(_dafny.Zero)
			_ = _element0_47
			_nw254 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
			(_nw254).ArraySet1(_element0_47, 0)
			var _nativeLen0_47 = (_len0_47).Int()
			_ = _nativeLen0_47
			for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
				(_nw254).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
			}
		}
		_1639_v11 = _nw254
		var _1641_v12 _dafny.Map
		_ = _1641_v12
		_1641_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1639_v11, (p1).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality())) < 0)
		var _1642_v13 D11
		_ = _1642_v13
		_1642_v13 = Companion_D11_.Create_DC30_()
		var _1643_v14 _dafny.Sequence
		_ = _1643_v14
		_1643_v14 = _dafny.SeqOf(Companion_D11_.Create_DC30_(), _1642_v13, _1642_v13, _1642_v13, _1642_v13)
		var _1644_v15 _dafny.Map
		_ = _1644_v15
		_1644_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1643_v14)
		var _rhs287 _dafny.Int = (_dafny.Zero).Minus((p1).Plus((_dafny.Zero).Minus(p1)))
		_ = _rhs287
		var _rhs288 _dafny.Int = ((_1644_v15).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.Companion_Sequence_.Update(_1643_v14, (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1643_v14).Cardinality()))).Uint32(), _1642_v13))).Merge(_1644_v15))).Cardinality()
		_ = _rhs288
		var _rhs289 _dafny.Map = _1641_v12
		_ = _rhs289
		_1638_v10 = _rhs287
		_1638_v10 = _rhs288
		_1641_v12 = _rhs289
		var _1645_v16 _dafny.Map
		_ = _1645_v16
		_1645_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1637_v9, (_this).F9())
		var _1646_v17 D9
		_ = _1646_v17
		_1646_v17 = Companion_D9_.Create_DC23_(_1645_v16)
		var _pat_let_tv83 = _1645_v16
		_ = _pat_let_tv83
		var _1647_v18 _dafny.Map
		_ = _1647_v18
		_1647_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D9 {
			if (_this).F9() {
				return func(_pat_let62_0 D9) D9 {
					return func(_1648_dt__update__tmp_h0 D9) D9 {
						return func(_pat_let63_0 _dafny.Map) D9 {
							return func(_1649_dt__update_hcf40_h0 _dafny.Map) D9 {
								return Companion_D9_.Create_DC23_(_1649_dt__update_hcf40_h0)
							}(_pat_let63_0)
						}(_pat_let_tv83)
					}(_pat_let62_0)
				}(_1646_v17)
			}
			return _1646_v17
		})(), _1637_v9)
		_1647_v18 = (_1647_v18).Update(_1646_v17, _1637_v9)
		var _1650_v19 _dafny.CodePoint
		_ = _1650_v19
		_1650_v19 = _dafny.CodePoint('j')
		_1650_v19 = _1650_v19
		var _1651_v20 _dafny.Array
		_ = _1651_v20
		var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
		_ = _nw255
		_1651_v20 = _nw255
		var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1651_v20), 0))
		_ = _index247
		(_1651_v20).ArraySet1(p0, (_index247).Int())
		var _1652_v21 _dafny.Array
		_ = _1652_v21
		var _nw256 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
		_ = _nw256
		_1652_v21 = _nw256
		var _1653_v22 _dafny.Sequence
		_ = _1653_v22
		_1653_v22 = _dafny.SeqOf(_1652_v21)
		var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1651_v20), 0))
		_ = _index248
		var _rhs290 _dafny.Array = p0
		_ = _rhs290
		var _rhs291 _dafny.Array = (_1653_v22).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1653_v22).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs291
		var _lhs179 _dafny.Array = _1651_v20
		_ = _lhs179
		var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1651_v20), 0))
		_ = _lhs180
		(_lhs179).ArraySet1(_rhs290, (_lhs180).Int())
		_1652_v21 = _rhs291
	}
}
func (_this *C5) M5(globalState *GlobalState) (bool, _dafny.Int, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1654_i0 _dafny.Int
		_ = _1654_i0
		_1654_i0 = _dafny.Zero
		{
			for (func() bool {
				if (_this).F9() {
					return (_this).F9()
				}
				return (_this).F9()
			})() {
				{
					if (_1654_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1654_i0 = (_1654_i0).Plus(_dafny.One)
					var _1655_v0 _dafny.Array
					_ = _1655_v0
					var _len0_48 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_48
					var _nw257 _dafny.Array
					_ = _nw257
					if _len0_48.Cmp(_dafny.Zero) == 0 {
						_nw257 = _dafny.NewArray(_len0_48)
					} else {
						var _init48 func(_dafny.Int) _dafny.CodePoint = func(_1656_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('x')
						}
						_ = _init48
						var _element0_48 = _init48(_dafny.Zero)
						_ = _element0_48
						_nw257 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
						(_nw257).ArraySet1CodePoint(_element0_48, 0)
						var _nativeLen0_48 = (_len0_48).Int()
						_ = _nativeLen0_48
						for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
							(_nw257).ArraySet1CodePoint(_init48(_dafny.IntOf(_i0_48)), _i0_48)
						}
					}
					_1655_v0 = _nw257
					var _1657_v1 D13
					_ = _1657_v1
					_1657_v1 = Companion_D13_.Create_DC34_(_1655_v0)
					var _1658_v2 _dafny.Map
					_ = _1658_v2
					_1658_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1657_v1)
					var _1659_v3 _dafny.Map
					_ = _1659_v3
					_1659_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_1658_v2).Cardinality())
					_1659_v3 = (_1659_v3).Merge((func() _dafny.Map {
						if false {
							return _1659_v3
						}
						return (_1659_v3).Update((_this).F9(), (_this).F10())
					})())
					var _1660_v4 _dafny.Sequence
					_ = _1660_v4
					_1660_v4 = _dafny.SeqOf((_this).F9(), (_this).F9())
					var _1661_v5 _dafny.Set
					_ = _1661_v5
					_1661_v5 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1660_v4, _1660_v4)).Cardinality()), (_this).F10(), (func() _dafny.Int {
						if (_this).F9() {
							return _dafny.IntOfUint32((_dafny.SeqOf((_this).F9(), (_this).F9())).Cardinality())
						}
						return (_this).F10()
					})(), _dafny.IntOfInt64(827))
					_1661_v5 = _1661_v5
					var _1662_v6 _dafny.Array
					_ = _1662_v6
					var _len0_49 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_49
					var _nw258 _dafny.Array
					_ = _nw258
					if _len0_49.Cmp(_dafny.Zero) == 0 {
						_nw258 = _dafny.NewArray(_len0_49)
					} else {
						var _init49 func(_dafny.Int) _dafny.MultiSet = func(_1663_i2 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf(_dafny.IntOfInt64(430), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality()), (_this).F10(), (_this).F10(), (_dafny.Zero).Minus((_this).F10()))
						}
						_ = _init49
						var _element0_49 = _init49(_dafny.Zero)
						_ = _element0_49
						_nw258 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
						(_nw258).ArraySet1(_element0_49, 0)
						var _nativeLen0_49 = (_len0_49).Int()
						_ = _nativeLen0_49
						for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
							(_nw258).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
						}
					}
					_1662_v6 = _nw258
					_1662_v6 = _1662_v6
					var _1664_v7 _dafny.Array
					_ = _1664_v7
					var _nw259 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(13))
					_ = _nw259
					_1664_v7 = _nw259
					var _1665_v8 _dafny.Sequence
					_ = _1665_v8
					_1665_v8 = _dafny.SeqOf(_1664_v7, _1664_v7)
					r0 = ((_this).F9()) || ((_dafny.IntOfUint32((_1665_v8).Cardinality())).Cmp((_this).F10()) >= 0)
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1666_v9 _dafny.Array
		_ = _1666_v9
		var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
		_ = _nw260
		_1666_v9 = _nw260
		var _1667_v10 *C3
		_ = _1667_v10
		var _nw261 *C3 = New_C3_()
		_ = _nw261
		_nw261.Ctor__((_this).F9(), _1666_v9, (_this).F9())
		_1667_v10 = _nw261
		var _1668_v11 _dafny.Array
		_ = _1668_v11
		var _nwElement0_51 *C3 = _1667_v10
		_ = _nwElement0_51
		var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(12))
		_ = _nw262
		(_nw262).ArraySet1(_nwElement0_51, 0)
		(_nw262).ArraySet1(_1667_v10, 1)
		(_nw262).ArraySet1(_1667_v10, 2)
		(_nw262).ArraySet1(_1667_v10, 3)
		(_nw262).ArraySet1(_1667_v10, 4)
		(_nw262).ArraySet1(_1667_v10, 5)
		(_nw262).ArraySet1(_1667_v10, 6)
		(_nw262).ArraySet1(_1667_v10, 7)
		(_nw262).ArraySet1(_1667_v10, 8)
		(_nw262).ArraySet1(_1667_v10, 9)
		(_nw262).ArraySet1(_1667_v10, 10)
		(_nw262).ArraySet1(_1667_v10, 11)
		_1668_v11 = _nw262
		var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_1668_v11), 0))
		_ = _index249
		(_1668_v11).ArraySet1(_1667_v10, (_index249).Int())
		var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_1668_v11), 0))
		_ = _index250
		(_1668_v11).ArraySet1(_1667_v10, (_index250).Int())
		var _1669_v12 _dafny.Sequence
		_ = _1669_v12
		_1669_v12 = _dafny.SeqOf((_1667_v10).F15(), (_this).Fm3(globalState))
		var _hi7 _dafny.Int = _dafny.IntOfUint32((_1669_v12).Cardinality())
		_ = _hi7
		for _1670_i3 := (_dafny.Zero).Minus((_this).F10()); _1670_i3.Cmp(_hi7) < 0; _1670_i3 = _1670_i3.Plus(_dafny.One) {
			var _1671_v13 D4
			_ = _1671_v13
			_1671_v13 = Companion_D4_.Create_DC13_((_this).F10())
			var _1672_v14 _dafny.Sequence
			_ = _1672_v14
			_1672_v14 = _dafny.SeqOf(_dafny.SetOf((_1671_v13).Dtor_cf25()))
			var _1673_v15 _dafny.Sequence
			_ = _1673_v15
			_1673_v15 = _dafny.SeqOf((_this).F10())
			var _1674_v16 _dafny.Map
			_ = _1674_v16
			_1674_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_dafny.MultiSetFromSeq(_1673_v15), _1670_i3, globalState), (_1667_v10).F15())
			(globalState).F2 = ((_1667_v10).F15()) || ((_1667_v10).Fm32(_1672_v14, (_1667_v10).F15(), (_1674_v16).Cardinality(), (_this).F9(), globalState))
			var _1675_v17 T1
			_ = _1675_v17
			var _nw263 *C2 = New_C2_()
			_ = _nw263
			_nw263.Ctor__(_1670_i3, (_this).F9())
			_1675_v17 = _nw263
			var _1676_v18 _dafny.Map
			_ = _1676_v18
			_1676_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1667_v10).F15(), _1675_v17)
			var _1677_v19 _dafny.Set
			_ = _1677_v19
			_1677_v19 = _dafny.SetOf(_1670_i3, (_1676_v18).Cardinality())
			var _1678_v20 _dafny.Set
			_ = _1678_v20
			_1678_v20 = _1677_v19
			var _source17 _dafny.Set = _1678_v20
			_ = _source17
			var _1679___mcc_h0 _dafny.Set = _source17
			_ = _1679___mcc_h0
			var _1680_cf15 _dafny.Set = _1679___mcc_h0
			_ = _1680_cf15
			var _1681_v21 D13
			_ = _1681_v21
			_1681_v21 = Companion_D13_.Create_DC36_(_dafny.UnicodeSeqOfUtf8Bytes("hpwdus"), _1670_i3, (_this).F10(), (_this).F10())
			r2 = ((_1681_v21).Dtor_cf64()).Times(_1670_i3)
			var _1682_v22 _dafny.Array
			_ = _1682_v22
			var _nw264 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw264
			_1682_v22 = _nw264
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1682_v22), 0))
			_ = _index251
			(_1682_v22).ArraySet1((_1667_v10).F15(), (_index251).Int())
			var _1683_v23 _dafny.Map
			_ = _1683_v23
			_1683_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1675_v17).F9(), _1670_i3)
			var _1684_v25 _dafny.Set
			_ = _1684_v25
			_1684_v25 = _dafny.SetOf(_1673_v15)
			var _1685_v26 *C0
			_ = _1685_v26
			var _nw265 *C0 = New_C0_()
			_ = _nw265
			_nw265.Ctor__(_1684_v25, _1670_i3)
			_1685_v26 = _nw265
			var _1686_v27 _dafny.Map
			_ = _1686_v27
			_1686_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1685_v26, (_this).F10())
			var _1687_v28 _dafny.Map
			_ = _1687_v28
			_1687_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-4), _1686_v27)
			var _1688_v29 _dafny.MultiSet
			_ = _1688_v29
			_1688_v29 = _dafny.MultiSetOf((func() _dafny.Int {
				if (_1683_v23).Contains((_1667_v10).F15()) {
					return (_1683_v23).Get((_1667_v10).F15()).(_dafny.Int)
				}
				return (func() _dafny.Set {
					var _coll36 = _dafny.NewBuilder()
					_ = _coll36
					for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-117), _dafny.IntOfInt64(410))); ; {
						_compr_36, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _1689_v24 _dafny.Int
						_1689_v24 = interface{}(_compr_36).(_dafny.Int)
						if ((_dafny.IntOfInt64(-117)).Cmp(_1689_v24) <= 0) && ((_1689_v24).Cmp(_dafny.IntOfInt64(410)) < 0) {
							_coll36.Add(Companion_Default___.SafeDivisionInt(_1689_v24, _dafny.IntOfInt64(-521)))
						}
					}
					return _coll36.ToSet()
				}()).Cardinality()
			})(), _1670_i3, ((func() _dafny.Map {
				if (_1687_v28).Contains((_this).F10()) {
					return (_1687_v28).Get((_this).F10()).(_dafny.Map)
				}
				return _1686_v27
			})()).Cardinality())
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1682_v22), 0))
			_ = _index252
			(_1682_v22).ArraySet1((func() bool {
				if (_1667_v10).F15() {
					return (_1667_v10).F15()
				}
				return !((Companion_Default___.Fm29((_this).F9(), (_this).F9(), globalState)).IsSubsetOf(_1688_v29))
			})(), (_index252).Int())
			r1 = _dafny.IntOfInt64(342)
			(_1685_v26).F12 = (_dafny.Zero).Minus(_1670_i3)
			var _1690_v30 _dafny.Set
			_ = _1690_v30
			_1690_v30 = _dafny.SetOf((_1675_v17).F9(), (_1667_v10).F15(), (_1667_v10).F15())
			var _rhs292 bool = ((_this).F10()).Cmp((_this).F10()) != 0
			_ = _rhs292
			var _rhs293 _dafny.Set = _1690_v30
			_ = _rhs293
			var _rhs294 _dafny.Int = Companion_Default___.SafeModuloInt(_1670_i3, Companion_Default___.SafeModuloInt((_1673_v15).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F10()), _dafny.IntOfUint32((_1673_v15).Cardinality()))).Uint32()).(_dafny.Int), _1670_i3))
			_ = _rhs294
			r0 = _rhs292
			_1690_v30 = _rhs293
			r2 = _rhs294
			var _1691_v31 *C2
			_ = _1691_v31
			var _nw266 *C2 = New_C2_()
			_ = _nw266
			_nw266.Ctor__(_1670_i3, !((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uyy")).Cardinality())).Cmp(Companion_Default___.Fm0(globalState)) > 0))
			_1691_v31 = _nw266
		}
		var _1692_v32 _dafny.Set
		_ = _1692_v32
		_1692_v32 = _dafny.SetOf((_1667_v10).F15())
		var _1693_v33 D15
		_ = _1693_v33
		_1693_v33 = Companion_D15_.Create_DC40_(_1692_v32)
		var _source18 D15 = _1693_v33
		_ = _source18
		if _source18.Is_DC41() {
			r2 = (_this).F10()
			var _1694_v34 _dafny.Array
			_ = _1694_v34
			var _len0_50 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_50
			var _nw267 _dafny.Array
			_ = _nw267
			if _len0_50.Cmp(_dafny.Zero) == 0 {
				_nw267 = _dafny.NewArray(_len0_50)
			} else {
				var _init50 func(_dafny.Int) bool = (func(_1695_v10 *C3) func(_dafny.Int) bool {
					return func(_1696_i4 _dafny.Int) bool {
						return (_1695_v10).F15()
					}
				})(_1667_v10)
				_ = _init50
				var _element0_50 = _init50(_dafny.Zero)
				_ = _element0_50
				_nw267 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
				(_nw267).ArraySet1(_element0_50, 0)
				var _nativeLen0_50 = (_len0_50).Int()
				_ = _nativeLen0_50
				for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
					(_nw267).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
				}
			}
			_1694_v34 = _nw267
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1694_v34), 0))
			_ = _index253
			(_1694_v34).ArraySet1((_this).F9(), (_index253).Int())
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1694_v34), 0))
			_ = _index254
			(_1694_v34).ArraySet1((_1667_v10).F15(), (_index254).Int())
			var _1697_v35 _dafny.Sequence
			_ = _1697_v35
			_1697_v35 = _dafny.SeqOf((_1692_v32).Cardinality(), (_this).F10())
			var _1698_v36 _dafny.Set
			_ = _1698_v36
			_1698_v36 = _dafny.SetOf(_1697_v35)
			var _1699_v37 _dafny.Array
			_ = _1699_v37
			var _len0_51 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_51
			var _nw268 _dafny.Array
			_ = _nw268
			if _len0_51.Cmp(_dafny.Zero) == 0 {
				_nw268 = _dafny.NewArray(_len0_51)
			} else {
				var _init51 func(_dafny.Int) _dafny.Int = func(_1700_i5 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1700_i5, (_this).F10())
				}
				_ = _init51
				var _element0_51 = _init51(_dafny.Zero)
				_ = _element0_51
				_nw268 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
				(_nw268).ArraySet1(_element0_51, 0)
				var _nativeLen0_51 = (_len0_51).Int()
				_ = _nativeLen0_51
				for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
					(_nw268).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
				}
			}
			_1699_v37 = _nw268
			var _1701_v38 _dafny.Set
			_ = _1701_v38
			_1701_v38 = _dafny.SetOf(_1699_v37, _1699_v37, _1699_v37)
			var _1702_v39 _dafny.Map
			_ = _1702_v39
			_1702_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm5(_1698_v36, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1701_v38).Cardinality(), _1669_v12), _dafny.IntOfUint32((_dafny.SeqOf((_this).F9())).Cardinality()), globalState), (_this).F10())
			var _1703_v40 _dafny.MultiSet
			_ = _1703_v40
			_1703_v40 = _dafny.MultiSetOf((_1702_v39).Cardinality(), (_this).F10())
			var _1704_v41 _dafny.Map
			_ = _1704_v41
			_1704_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1667_v10).F15(), Companion_Default___.Fm1(_1703_v40, (_1697_v35).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1697_v35).Cardinality()))).Uint32()).(_dafny.Int), globalState))
			_1704_v41 = (_1704_v41).Update((_1667_v10).F15(), !((_1694_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1694_v34), 0))).Int()).(bool)))
			var _1705_v42 _dafny.MultiSet
			_ = _1705_v42
			_1705_v42 = _dafny.MultiSetOf(_dafny.CodePoint('x'))
			var _1706_v43 *C0
			_ = _1706_v43
			var _nw269 *C0 = New_C0_()
			_ = _nw269
			_nw269.Ctor__(_1698_v36, (_1705_v42).Cardinality())
			_1706_v43 = _nw269
			_1706_v43 = _1706_v43
		} else if _source18.Is_DC42() {
			var _1707___mcc_h1 bool = _source18.Get_().(D15_DC42).Cf73
			_ = _1707___mcc_h1
			var _1708___mcc_h2 _dafny.Int = _source18.Get_().(D15_DC42).Cf74
			_ = _1708___mcc_h2
			var _1709___mcc_h3 T2 = _source18.Get_().(D15_DC42).Cf75
			_ = _1709___mcc_h3
			var _1710___mcc_h4 _dafny.Int = _source18.Get_().(D15_DC42).Cf76
			_ = _1710___mcc_h4
			var _1711_cf76 _dafny.Int = _1710___mcc_h4
			_ = _1711_cf76
			var _1712_cf75 T2 = _1709___mcc_h3
			_ = _1712_cf75
			var _1713_cf74 _dafny.Int = _1708___mcc_h2
			_ = _1713_cf74
			var _1714_cf73 bool = _1707___mcc_h1
			_ = _1714_cf73
			var _1715_v45 _dafny.Sequence
			_ = _1715_v45
			_1715_v45 = _dafny.SeqOf((_dafny.MultiSetOf(_1711_cf76, (_this).F10())).Cardinality())
			var _1716_v46 _dafny.MultiSet
			_ = _1716_v46
			_1716_v46 = _dafny.MultiSetOf(_dafny.IntOfInt64(48), _1711_cf76)
			var _1717_v47 _dafny.MultiSet
			_ = _1717_v47
			_1717_v47 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_1715_v45), _1716_v46)
			var _1718_v48 _dafny.CodePoint
			_ = _1718_v48
			_1718_v48 = _dafny.CodePoint('h')
			var _1719_v49 _dafny.Map
			_ = _1719_v49
			_1719_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_this).F10(), _1713_cf74), _1718_v48)
			var _1720_v50 _dafny.Map
			_ = _1720_v50
			_1720_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if !(!(_1714_cf73)) {
					return func() _dafny.Map {
						var _coll37 = _dafny.NewMapBuilder()
						_ = _coll37
						for _iter41 := _dafny.Iterate((_1717_v47).Elements()); ; {
							_compr_37, _ok41 := _iter41()
							if !_ok41 {
								break
							}
							var _1721_v44 _dafny.MultiSet
							_1721_v44 = interface{}(_compr_37).(_dafny.MultiSet)
							if (_1717_v47).Contains(_1721_v44) {
								_coll37.Add(_1721_v44, _dafny.CodePoint('r'))
							}
						}
						return _coll37.ToMap()
					}()
				}
				return _1719_v49
			})(), Companion_Default___.SafeModuloInt((_this).F10(), _dafny.IntOfUint32((_1715_v45).Cardinality())))
			_1720_v50 = (_1720_v50).Update(_1719_v49, (_this).F10())
			var _1722_v51 _dafny.Array
			_ = _1722_v51
			var _nw270 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(9))
			_ = _nw270
			_1722_v51 = _nw270
			var _1723_v52 _dafny.Set
			_ = _1723_v52
			_1723_v52 = _dafny.SetOf((func() _dafny.Array {
				if (_this).F9() {
					return _1722_v51
				}
				return _1722_v51
			})())
			_1723_v52 = _1723_v52
			_1713_cf74 = _1713_cf74
			var _1724_v53 _dafny.Set
			_ = _1724_v53
			_1724_v53 = _dafny.SetOf((_this).F10())
			var _1725_v54 _dafny.Sequence
			_ = _1725_v54
			_1725_v54 = _dafny.SeqOf(_1724_v53, _dafny.SetOf(_1711_cf76))
			var _1726_v55 _dafny.MultiSet
			_ = _1726_v55
			_1726_v55 = _dafny.MultiSetOf((_1667_v10).F15())
			var _1727_v56 _dafny.Sequence
			_ = _1727_v56
			_1727_v56 = _dafny.UnicodeSeqOfUtf8Bytes("bxndbo")
			var _1728_v57 D1
			_ = _1728_v57
			_1728_v57 = Companion_D1_.Create_DC2_(((_1726_v55).Update(true, Companion_Default___.Abs(_dafny.IntOfUint32((_1727_v56).Cardinality())))).Cardinality(), _1714_cf73, _dafny.IntOfUint32((_1727_v56).Cardinality()), _1711_cf76)
			var _1729_v58 _dafny.Map
			_ = _1729_v58
			_1729_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1718_v48, _1716_v46)
			var _1730_v59 _dafny.Array
			_ = _1730_v59
			var _nwElement0_52 bool = (_1667_v10).Fm32(_1725_v54, (_this).Fm3(globalState), _dafny.IntOfInt64(328), (_1728_v57).Dtor_cf3(), globalState)
			_ = _nwElement0_52
			var _nw271 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(25))
			_ = _nw271
			(_nw271).ArraySet1(_nwElement0_52, 0)
			(_nw271).ArraySet1((func() bool {
				if (_1667_v10).F15() {
					return true
				}
				return (_1712_cf75).F9()
			})(), 1)
			(_nw271).ArraySet1((_this).F9(), 2)
			(_nw271).ArraySet1(!((func() bool {
				if (_1712_cf75).F9() {
					return (_1712_cf75).F9()
				}
				return (_1712_cf75).F9()
			})()), 3)
			(_nw271).ArraySet1(_1714_cf73, 4)
			(_nw271).ArraySet1(false, 5)
			(_nw271).ArraySet1((_1712_cf75).F9(), 6)
			(_nw271).ArraySet1((_1667_v10).F15(), 7)
			(_nw271).ArraySet1((_this).F9(), 8)
			(_nw271).ArraySet1((_this).F9(), 9)
			(_nw271).ArraySet1(_1714_cf73, 10)
			(_nw271).ArraySet1(!((_1712_cf75).Fm3(globalState)), 11)
			(_nw271).ArraySet1((_1713_cf74).Cmp((_this).F10()) <= 0, 12)
			(_nw271).ArraySet1((_1712_cf75).F9(), 13)
			(_nw271).ArraySet1((_this).F9(), 14)
			(_nw271).ArraySet1((_1667_v10).F15(), 15)
			(_nw271).ArraySet1(!((_1667_v10).F15()) || (false), 16)
			(_nw271).ArraySet1((_1712_cf75).Fm3(globalState), 17)
			(_nw271).ArraySet1(_1714_cf73, 18)
			(_nw271).ArraySet1(((_this).F10()).Cmp(_1711_cf76) > 0, 19)
			(_nw271).ArraySet1((_1712_cf75).F9(), 20)
			(_nw271).ArraySet1((_1667_v10).F15(), 21)
			(_nw271).ArraySet1(((_this).F10()).Cmp(_1713_cf74) < 0, 22)
			(_nw271).ArraySet1(!(Companion_Default___.Fm1((func() _dafny.MultiSet {
				if (_1729_v58).Contains((_1727_v56).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.IntOfUint32((_1727_v56).Cardinality()))).Uint32()).(_dafny.CodePoint)) {
					return (_1729_v58).Get((_1727_v56).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.IntOfUint32((_1727_v56).Cardinality()))).Uint32()).(_dafny.CodePoint)).(_dafny.MultiSet)
				}
				return _1716_v46
			})(), (_this).F10(), globalState)), 23)
			(_nw271).ArraySet1((_1667_v10).F15(), 24)
			_1730_v59 = _nw271
			var _1731_v60 *C1
			_ = _1731_v60
			var _nw272 *C1 = New_C1_()
			_ = _nw272
			_nw272.Ctor__(_dafny.SeqOf(_1728_v57, _1728_v57))
			_1731_v60 = _nw272
			var _1732_v61 D13
			_ = _1732_v61
			_1732_v61 = Companion_D13_.Create_DC35_(_1714_cf73, _1718_v48, _1731_v60, _1714_cf73)
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_1730_v59), 0))
			_ = _index255
			(_1730_v59).ArraySet1((_1732_v61).Dtor_cf58(), (_index255).Int())
			var _1733_v62 _dafny.Map
			_ = _1733_v62
			_1733_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1713_cf74, _1713_cf74)
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_1730_v59), 0))
			_ = _index256
			var _rhs295 bool = (_1669_v12).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_1733_v62).Cardinality(), _1713_cf74), _dafny.IntOfUint32((_1669_v12).Cardinality()))).Uint32()).(bool)
			_ = _rhs295
			var _rhs296 _dafny.Sequence = _1669_v12
			_ = _rhs296
			var _rhs297 bool = _1714_cf73
			_ = _rhs297
			var _lhs181 _dafny.Array = _1730_v59
			_ = _lhs181
			var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_1730_v59), 0))
			_ = _lhs182
			var _lhs183 *GlobalState = globalState
			_ = _lhs183
			(_lhs181).ArraySet1(_rhs295, (_lhs182).Int())
			_1669_v12 = _rhs296
			_lhs183.F2 = _rhs297
		} else if _source18.Is_DC43() {
			var _1734_v63 _dafny.Sequence
			_ = _1734_v63
			_1734_v63 = _dafny.UnicodeSeqOfUtf8Bytes("nf")
			var _1735_v64 D1
			_ = _1735_v64
			_1735_v64 = Companion_D1_.Create_DC2_((_this).F10(), (_1667_v10).F15(), (_this).F10(), _dafny.IntOfUint32((_1734_v63).Cardinality()))
			var _1736_v65 _dafny.Sequence
			_ = _1736_v65
			_1736_v65 = _dafny.SeqOf(_1735_v64)
			var _1737_v66 *C1
			_ = _1737_v66
			var _nw273 *C1 = New_C1_()
			_ = _nw273
			_nw273.Ctor__(_1736_v65)
			_1737_v66 = _nw273
			var _1738_v67 _dafny.MultiSet
			_ = _1738_v67
			_1738_v67 = _dafny.MultiSetOf(_1737_v66, _1737_v66)
			var _1739_v68 _dafny.Map
			_ = _1739_v68
			_1739_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1738_v67, (_1667_v10).F15())
			_1739_v68 = (_1739_v68).Update((_1738_v67).Update(_1737_v66, Companion_Default___.Abs(_dafny.IntOfInt64(677))), (_1667_v10).F15())
			var _1740_v69 _dafny.Set
			_ = _1740_v69
			_1740_v69 = _dafny.SetOf(_dafny.IntOfUint32((_1734_v63).Cardinality()))
			var _1741_v70 _dafny.Sequence
			_ = _1741_v70
			_1741_v70 = _dafny.SeqOf(_dafny.SetOf((_this).F10(), (_this).F10()), _1740_v69)
			var _1742_v71 D3
			_ = _1742_v71
			_1742_v71 = Companion_D3_.Create_DC9_((_this).F9(), _1669_v12)
			var _1743_v72 _dafny.Sequence
			_ = _1743_v72
			_1743_v72 = _dafny.SeqOf(_dafny.IntOfInt64(-45), Companion_Default___.SafeModuloInt((_this).Fm4(false, (_1667_v10).Fm32(_1741_v70, (_1742_v71).Dtor_cf17(), (_this).F10(), (_this).Fm3(globalState), globalState), (_1667_v10).F15(), (_1667_v10).F15(), globalState), (_this).F10()))
			_1743_v72 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-392))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg82 _dafny.Int) interface{} {
					return coer82(arg82)
				}
			}(func(_1744_i6 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_this).F10())
			}))
			r0 = (_1669_v12).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1669_v12).Cardinality()))).Uint32()).(bool)
			var _1745_v73 _dafny.Map
			_ = _1745_v73
			_1745_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1743_v72, _1734_v63)
			var _1746_v75 _dafny.Map
			_ = _1746_v75
			_1746_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.SeqOf((_1667_v10).F15()))
			var _1747_v76 _dafny.Array
			_ = _1747_v76
			var _nwElement0_53 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(false), _1669_v12)
			_ = _nwElement0_53
			var _nw274 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(11))
			_ = _nw274
			(_nw274).ArraySet1(_nwElement0_53, 0)
			(_nw274).ArraySet1((_this).F9(), 1)
			(_nw274).ArraySet1((_this).F9(), 2)
			(_nw274).ArraySet1((_this).Fm5(func() _dafny.Set {
				var _coll38 = _dafny.NewBuilder()
				_ = _coll38
				for _iter42 := _dafny.Iterate((_1745_v73).Keys().Elements()); ; {
					_compr_38, _ok42 := _iter42()
					if !_ok42 {
						break
					}
					var _1748_v74 _dafny.Sequence
					_1748_v74 = interface{}(_compr_38).(_dafny.Sequence)
					if (_1745_v73).Contains(_1748_v74) {
						_coll38.Add(_1748_v74)
					}
				}
				return _coll38.ToSet()
			}(), _1746_v75, (_this).F10(), globalState), 3)
			(_nw274).ArraySet1((_1667_v10).F15(), 4)
			(_nw274).ArraySet1((_1667_v10).F15(), 5)
			(_nw274).ArraySet1((_this).F9(), 6)
			(_nw274).ArraySet1((_1667_v10).F15(), 7)
			(_nw274).ArraySet1(false, 8)
			(_nw274).ArraySet1((_1667_v10).Fm3(globalState), 9)
			(_nw274).ArraySet1(true, 10)
			_1747_v76 = _nw274
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_1747_v76), 0))
			_ = _index257
			(_1747_v76).ArraySet1((_1667_v10).F15(), (_index257).Int())
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_1747_v76), 0))
			_ = _index258
			(_1747_v76).ArraySet1(!((_1667_v10).F15()), (_index258).Int())
		} else if _source18.Is_DC40() {
			var _1749___mcc_h5 _dafny.Set = _source18.Get_().(D15_DC40).Cf72
			_ = _1749___mcc_h5
			var _1750_cf72 _dafny.Set = _1749___mcc_h5
			_ = _1750_cf72
			var _1751_v77 _dafny.Sequence
			_ = _1751_v77
			_1751_v77 = _dafny.UnicodeSeqOfUtf8Bytes("lf")
			var _1752_v78 D1
			_ = _1752_v78
			_1752_v78 = Companion_D1_.Create_DC2_((_this).F10(), (_1667_v10).F15(), _dafny.IntOfUint32((_1751_v77).Cardinality()), (_this).F10())
			var _1753_v79 _dafny.Sequence
			_ = _1753_v79
			_1753_v79 = _dafny.SeqOf(_1752_v78, _1752_v78)
			var _1754_v80 *C1
			_ = _1754_v80
			var _nw275 *C1 = New_C1_()
			_ = _nw275
			_nw275.Ctor__(_1753_v79)
			_1754_v80 = _nw275
			var _1755_v81 _dafny.Sequence
			_ = _1755_v81
			_1755_v81 = _dafny.SeqOf((_this).F10())
			_1755_v81 = _1755_v81
			r1 = (_this).F10()
			_1755_v81 = _1755_v81
		} else {
			var _1756___mcc_h6 D15 = _source18.Get_().(D15_DC44).Cf77
			_ = _1756___mcc_h6
			var _1757_cf77 D15 = _1756___mcc_h6
			_ = _1757_cf77
			var _1758_v82 _dafny.Sequence
			_ = _1758_v82
			_1758_v82 = _dafny.UnicodeSeqOfUtf8Bytes("dxjy")
			var _1759_v83 _dafny.Map
			_ = _1759_v83
			_1759_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1758_v82, (_dafny.Zero).Minus((_this).F10()))
			var _1760_v84 _dafny.Set
			_ = _1760_v84
			_1760_v84 = _dafny.SetOf(((_dafny.Zero).Minus((_1759_v83).Cardinality())).Minus((_this).F10()), (_this).F10(), (_this).F10())
			var _1761_v85 _dafny.CodePoint
			_ = _1761_v85
			_1761_v85 = _dafny.CodePoint('l')
			var _1762_v86 _dafny.Map
			_ = _1762_v86
			_1762_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D7_.Create_DC20_(true, _1761_v85)).Dtor_cf36(), !(true))
			_1760_v84 = ((_1760_v84).Difference(_dafny.SetOf((_this).F10(), _dafny.IntOfUint32((_1669_v12).Cardinality()), (_1762_v86).Cardinality(), (_this).F10()))).Union(_1760_v84)
			var _1763_v87 T2
			_ = _1763_v87
			var _nw276 *C3 = New_C3_()
			_ = _nw276
			_nw276.Ctor__((_1667_v10).F15(), _1666_v9, (_1667_v10).F15())
			_1763_v87 = _nw276
			var _1764_v88 _dafny.Array
			_ = _1764_v88
			var _nwElement0_54 T2 = _1763_v87
			_ = _nwElement0_54
			var _nw277 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(28))
			_ = _nw277
			(_nw277).ArraySet1(_nwElement0_54, 0)
			(_nw277).ArraySet1(_1763_v87, 1)
			(_nw277).ArraySet1(_1763_v87, 2)
			(_nw277).ArraySet1(_1763_v87, 3)
			(_nw277).ArraySet1(_1763_v87, 4)
			(_nw277).ArraySet1(_1763_v87, 5)
			(_nw277).ArraySet1(_1763_v87, 6)
			(_nw277).ArraySet1(_1763_v87, 7)
			(_nw277).ArraySet1(_1763_v87, 8)
			(_nw277).ArraySet1(_1763_v87, 9)
			(_nw277).ArraySet1(_1763_v87, 10)
			(_nw277).ArraySet1(_1763_v87, 11)
			(_nw277).ArraySet1(_1763_v87, 12)
			(_nw277).ArraySet1(_1763_v87, 13)
			(_nw277).ArraySet1(_1763_v87, 14)
			(_nw277).ArraySet1(_1763_v87, 15)
			(_nw277).ArraySet1(_1763_v87, 16)
			(_nw277).ArraySet1(_1763_v87, 17)
			(_nw277).ArraySet1(_1763_v87, 18)
			(_nw277).ArraySet1(_1763_v87, 19)
			(_nw277).ArraySet1(_1763_v87, 20)
			(_nw277).ArraySet1(_1763_v87, 21)
			(_nw277).ArraySet1(_1763_v87, 22)
			(_nw277).ArraySet1(_1763_v87, 23)
			(_nw277).ArraySet1(_1763_v87, 24)
			(_nw277).ArraySet1(_1763_v87, 25)
			(_nw277).ArraySet1(_1763_v87, 26)
			(_nw277).ArraySet1(_1763_v87, 27)
			_1764_v88 = _nw277
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_1764_v88), 0))
			_ = _index259
			(_1764_v88).ArraySet1(_1763_v87, (_index259).Int())
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_1764_v88), 0))
			_ = _index260
			(_1764_v88).ArraySet1(_1763_v87, (_index260).Int())
			_1758_v82 = _1758_v82
			if !((func() bool {
				if ((_this).F10()).Cmp((_this).F10()) <= 0 {
					return (func() bool {
						if (_1667_v10).F15() {
							return (_1667_v10).F15()
						}
						return (_1763_v87).F9()
					})()
				}
				return (_this).F9()
			})()) {
				var _1765_v89 _dafny.Sequence
				_ = _1765_v89
				_1765_v89 = _dafny.SeqOf(((_this).F10()).Minus((_this).F10()))
				var _1766_v90 _dafny.Array
				_ = _1766_v90
				var _nw278 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw278
				_1766_v90 = _nw278
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1766_v90), 0))
				_ = _index261
				(_1766_v90).ArraySet1(true, (_index261).Int())
				var _1767_v91 _dafny.MultiSet
				_ = _1767_v91
				_1767_v91 = _dafny.MultiSetOf((_this).F10())
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1766_v90), 0))
				_ = _index262
				var _rhs298 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1765_v89, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F10()), _1765_v89))
				_ = _rhs298
				var _rhs299 bool = Companion_Default___.Fm1(_1767_v91, (_this).F10(), globalState)
				_ = _rhs299
				var _lhs184 _dafny.Array = _1766_v90
				_ = _lhs184
				var _lhs185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1766_v90), 0))
				_ = _lhs185
				_1765_v89 = _rhs298
				(_lhs184).ArraySet1(_rhs299, (_lhs185).Int())
				var _1768_v92 _dafny.Map
				_ = _1768_v92
				_1768_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(242), _1758_v82)
				var _1769_v93 _dafny.Map
				_ = _1769_v93
				_1769_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1763_v87).F9(), (func() _dafny.Sequence {
					if (_1768_v92).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lm")).Cardinality())) {
						return (_1768_v92).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lm")).Cardinality())).(_dafny.Sequence)
					}
					return _1758_v82
				})())
				var _1770_v94 _dafny.Sequence
				_ = _1770_v94
				_1770_v94 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-116))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1771_v85 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1772_i9 _dafny.Int) _dafny.CodePoint {
						return _1771_v85
					}
				})(_1761_v85))))
				var _1773_v95 _dafny.Array
				_ = _1773_v95
				var _nwElement0_55 _dafny.Sequence = _1758_v82
				_ = _nwElement0_55
				var _nw279 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(25))
				_ = _nw279
				(_nw279).ArraySet1(_nwElement0_55, 0)
				(_nw279).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1758_v82, _dafny.UnicodeSeqOfUtf8Bytes("ouulicrh")), 1)
				(_nw279).ArraySet1(_1758_v82, 2)
				(_nw279).ArraySet1(_1758_v82, 3)
				(_nw279).ArraySet1((func() _dafny.Sequence {
					if (_1667_v10).F15() {
						return _dafny.UnicodeSeqOfUtf8Bytes("cc")
					}
					return _1758_v82
				})(), 4)
				(_nw279).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(652))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}((func(_1774_v85 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1775_i7 _dafny.Int) _dafny.CodePoint {
						return _1774_v85
					}
				})(_1761_v85))), 5)
				(_nw279).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("a"), 6)
				(_nw279).ArraySet1(_1758_v82, 7)
				(_nw279).ArraySet1(_1758_v82, 8)
				(_nw279).ArraySet1(_1758_v82, 9)
				(_nw279).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("anxityp"), _dafny.UnicodeSeqOfUtf8Bytes("j")), 10)
				(_nw279).ArraySet1(_1758_v82, 11)
				(_nw279).ArraySet1((func() _dafny.Sequence {
					if (_1769_v93).Contains((_1763_v87).F9()) {
						return (_1769_v93).Get((_1763_v87).F9()).(_dafny.Sequence)
					}
					return _1758_v82
				})(), 12)
				(_nw279).ArraySet1(_1758_v82, 13)
				(_nw279).ArraySet1(Companion_Default___.Fm28((_this).F10(), !((_this).F9()), globalState), 14)
				(_nw279).ArraySet1(_1758_v82, 15)
				(_nw279).ArraySet1(_1758_v82, 16)
				(_nw279).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(839))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_1776_v85 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1777_i8 _dafny.Int) _dafny.CodePoint {
						return _1776_v85
					}
				})(_1761_v85))), 17)
				(_nw279).ArraySet1((_1770_v94).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1770_v94).Cardinality()))).Uint32()).(_dafny.Sequence), 18)
				(_nw279).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(859))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}((func(_1778_v85 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1779_i10 _dafny.Int) _dafny.CodePoint {
						return _1778_v85
					}
				})(_1761_v85))), 19)
				(_nw279).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_1780_v85 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1781_i11 _dafny.Int) _dafny.CodePoint {
						return _1780_v85
					}
				})(_1761_v85))), 20)
				(_nw279).ArraySet1(_1758_v82, 21)
				(_nw279).ArraySet1(_1758_v82, 22)
				(_nw279).ArraySet1(_1758_v82, 23)
				(_nw279).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fc"), 24)
				_1773_v95 = _nw279
				var _1782_v96 _dafny.Map
				_ = _1782_v96
				_1782_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1768_v92, _1773_v95)
				_1773_v95 = (func() _dafny.Array {
					if (_1782_v96).Contains((_1768_v92).Merge(_1768_v92)) {
						return (_1782_v96).Get((_1768_v92).Merge(_1768_v92)).(_dafny.Array)
					}
					return _1773_v95
				})()
				var _1783_v97 _dafny.MultiSet
				_ = _1783_v97
				_1783_v97 = _dafny.MultiSetOf((_1766_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1766_v90), 0))).Int()).(bool))
				_1783_v97 = _1783_v97
				var _1784_v98 *C2
				_ = _1784_v98
				var _nw280 *C2 = New_C2_()
				_ = _nw280
				_nw280.Ctor__((_this).F10(), (_1763_v87).F9())
				_1784_v98 = _nw280
				var _1785_v99 _dafny.Map
				_ = _1785_v99
				_1785_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), ((_1784_v98).F14()).Plus((_this).F10()))
				_1785_v99 = (_1785_v99).Update(false, (_1784_v98).F14())
			} else {
				r1 = (_this).F10()
				var _1786_v100 _dafny.Sequence
				_ = _1786_v100
				_1786_v100 = _dafny.SeqOf((_this).F10(), (func() _dafny.Int {
					if (_1667_v10).F15() {
						return (_this).F10()
					}
					return (_this).F10()
				})(), (_this).F10(), (_this).F10(), _dafny.IntOfUint32((_1758_v82).Cardinality()))
				_1786_v100 = _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F10()), (Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F10())).Cardinality()))).Uint32(), (_this).F10())
				var _1787_v101 *C2
				_ = _1787_v101
				var _nw281 *C2 = New_C2_()
				_ = _nw281
				_nw281.Ctor__((_dafny.Zero).Minus((_this).F10()), (_1667_v10).F15())
				_1787_v101 = _nw281
				var _1788_v102 _dafny.MultiSet
				_ = _1788_v102
				_1788_v102 = _dafny.MultiSetOf((_1763_v87).F9(), (_1763_v87).F9())
				var _1789_v103 bool
				_ = _1789_v103
				var _out17 bool
				_ = _out17
				_out17 = (_1667_v10).M3(((_dafny.MultiSetOf((_1667_v10).F15())).Update((_1669_v12).Select((Companion_Default___.SafeIndex((_this).F10(), _dafny.IntOfUint32((_1669_v12).Cardinality()))).Uint32()).(bool), Companion_Default___.Abs((_this).F10()))).IsDisjointFrom(_1788_v102), globalState)
				_1789_v103 = _out17
				r0 = (_dafny.IntOfUint32((_1669_v12).Cardinality())).Cmp((_dafny.Zero).Minus((_1787_v101).F14())) > 0
			}
		}
		var _1790_v104 _dafny.Array
		_ = _1790_v104
		var _nw282 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw282
		_1790_v104 = _nw282
		for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1790_v104), 0))); ; {
			_guard_loop_4, _ok43 := _iter43()
			if !_ok43 {
				break
			}
			var _1791_i12 _dafny.Int
			_1791_i12 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1791_i12).Sign() != -1) && ((_1791_i12).Cmp(_dafny.ArrayLen((_1790_v104), 0)) < 0)) {
				(_1790_v104).ArraySet1((_1791_i12).Plus((_this).F10()), (_1791_i12).Int())
			}
		}
		var _1792_v105 D16
		_ = _1792_v105
		_1792_v105 = Companion_D16_.Create_DC46_(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-44))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg88 _dafny.Int) interface{} {
				return coer88(arg88)
			}
		}(func(_1793_i13 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))).Cardinality()))).Minus((_this).F10()))
		_1792_v105 = _1792_v105
		r0 = (_1667_v10).F15()
		r1 = (_this).F10()
		r2 = (_this).F10()
		return r0, r1, r2
	}
}
func (_this *C5) F10() _dafny.Int {
	{
		return _this._f10
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
