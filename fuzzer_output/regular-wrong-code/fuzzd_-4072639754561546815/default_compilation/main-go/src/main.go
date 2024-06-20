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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	if !(false) {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(863))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), _dafny.UnicodeSeqOfUtf8Bytes("qy"))
	} else {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	return ((_dafny.IntOfInt64(716)).Plus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (Companion_D9_.Create_DC25_((_dafny.Zero).Minus(_dafny.IntOfInt64(-341)), false, true, !(true))).Dtor_cf44())).Cardinality(), _dafny.CodePoint('b'))).Cardinality()))).Plus(_dafny.IntOfInt64(28))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(true, true), _dafny.SeqOf(false, false, false, false, !(true))), !(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(74)), _dafny.IntOfInt64(-680))), !(false))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("e"), _dafny.UnicodeSeqOfUtf8Bytes("hipjlan")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("csbbvc"), _dafny.UnicodeSeqOfUtf8Bytes("ejenrea")))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.MultiSetOf(true, false, true), _dafny.MultiSetOf(true, false, true, true))
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality())).Cmp((_dafny.SetOf(false, false)).Cardinality()) <= 0 {
		return _dafny.CodePoint('q')
	} else {
		return _dafny.CodePoint('v')
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(986)), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Cardinality()))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.IntOfInt64(-717)).Cmp(_dafny.IntOfInt64(45)) == 0
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) bool {
	return !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("g"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wueqr"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(631))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	})))))
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(640)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-931), _dafny.IntOfInt64(501)), _dafny.SeqOf(_dafny.IntOfInt64(97))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false, !(false), true, true), _dafny.SeqOf(true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, !(false)), _dafny.SeqOf(!(true), false)))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(((Companion_D17_.Create_DC46_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-338))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))).Cardinality()), (_dafny.MultiSetOf(true, false)).Cardinality()))).Dtor_cf80()).Cardinality())).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(861), true)).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _5_v0 _dafny.Int
			_5_v0 = interface{}(_compr_0).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(861), true)).Contains(_5_v0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_5_v0, _dafny.IntOfInt64(63)))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality(), !(true))).Cardinality(), _dafny.IntOfInt64(190), (func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-462), _dafny.CodePoint('g'))).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _6_v1 _dafny.Int
			_6_v1 = interface{}(_compr_1).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-462), _dafny.CodePoint('g'))).Contains(_6_v1) {
				_coll1.Add((_6_v1).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(103))).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll2 = _dafny.NewMapBuilder()
					_ = _coll2
					for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), false)).Keys().Elements()); ; {
						_compr_2, _ok2 := _iter2()
						if !_ok2 {
							break
						}
						var _7_v2 _dafny.CodePoint
						_7_v2 = interface{}(_compr_2).(_dafny.CodePoint)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), false)).Contains(_7_v2) {
							_coll2.Add(_7_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xgonjaa")).Cardinality()))
						}
					}
					return _coll2.ToMap()
				}()).Cardinality(), _dafny.IntOfInt64(106)))
			}
		}
		return _coll1.ToMap()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(867))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_8_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	}))).Cardinality())), _dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('p')))).Cardinality(), _dafny.IntOfInt64(62), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(705))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_9_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(369))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_10_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		}))).Cardinality())
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jbdseslwf")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	return (Companion_D17_.Create_DC47_(!(false), (_dafny.Zero).Minus(_dafny.IntOfInt64(-432)))).Dtor_cf82()
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (Companion_D18_.Create_DC48_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(-652)))).Dtor_cf83()
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('m')
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.MultiSet, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(707), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(248), true))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if true {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(414))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_11_i0 _dafny.Int) _dafny.Set {
			return _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-540), false)).Cardinality())
		})), _dafny.SeqOf(func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(680), _dafny.IntOfInt64(4))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _12_v0 _dafny.Int
				_12_v0 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(680)).Cmp(_12_v0) <= 0) && ((_12_v0).Cmp(_dafny.IntOfInt64(4)) < 0) {
					_coll3.Add((_12_v0).Plus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dqm")).Cardinality()))).Cardinality()))
				}
			}
			return _coll3.ToSet()
		}()))
	} else {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_13_i1 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(_dafny.IntOfInt64(322), (_dafny.Zero).Minus(_dafny.IntOfInt64(-322)), (_dafny.MultiSetOf(_dafny.IntOfInt64(-173))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-996)), _dafny.IntOfInt64(-776))
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Map, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-110), _dafny.IntOfInt64(-249)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-74))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_14_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(916))).Cardinality())
	}))), _dafny.SeqOf(_dafny.IntOfInt64(378), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(11), _dafny.UnicodeSeqOfUtf8Bytes("ucvrgkh"))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(178))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_15_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))).Cardinality()), _dafny.CodePoint('t'))).Cardinality(), _dafny.IntOfInt64(933)))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfInt64(32))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(Companion_D2_.Create_DC6_((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false)))).Cardinality())))).Intersection(_dafny.SetOf(Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("msuhavx")).Cardinality())), Companion_D2_.Create_DC6_(_dafny.IntOfInt64(-218)), Companion_D2_.Create_DC6_((func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(783), _dafny.IntOfInt64(-172))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _16_v0 _dafny.Int
			_16_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(783)).Cmp(_16_v0) <= 0) && ((_16_v0).Cmp(_dafny.IntOfInt64(-172)) < 0) {
				_coll4.Add(Companion_Default___.SafeDivisionInt(_16_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(208))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}(func(_17_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(682), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jlnaqhjk")).Cardinality())))).Cardinality())
				}))).Cardinality())), false)
			}
		}
		return _coll4.ToMap()
	}()).Cardinality()), Companion_D2_.Create_DC6_(_dafny.IntOfInt64(531))))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.IntOfInt64(855)).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bfehftfco")).Cardinality()), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.CodePoint('u'), _dafny.CodePoint('n'))).Cardinality()))) == 0
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) D9 {
	if (_dafny.MultiSetOf(false, true, true)).IsDisjointFrom(_dafny.MultiSetOf(true, false, false, !(!(false)))) {
		return Companion_D9_.Create_DC23_(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(984))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))))
	} else {
		return Companion_D9_.Create_DC23_(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("nttb")))
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-747))).Cardinality(), _dafny.IntOfInt64(-663), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-818), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC25_(_dafny.IntOfInt64(933), false, true, false), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-790), false)).Cardinality(), _dafny.IntOfInt64(80)))).Cardinality())).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, p1 bool, globalState *GlobalState) D0 {
	if false {
		return Companion_D0_.Create_DC1_(_dafny.CodePoint('c'), false, true)
	} else {
		return Companion_D0_.Create_DC1_(_dafny.CodePoint('j'), false, false)
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.CodePoint, p1 D0, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC3_(_dafny.CodePoint('u'))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-687), _dafny.IntOfInt64(640))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(-687)).Cmp(_19_v0) <= 0) && ((_19_v0).Cmp(_dafny.IntOfInt64(640)) < 0) {
				_coll5.Add((_19_v0).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(379))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}(func(_20_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(604))).Cardinality())
				}))).Cardinality())), _dafny.IntOfInt64(-983))
			}
		}
		return _coll5.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(-585))).Cardinality(), _dafny.IntOfInt64(-836)))
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 D3, p2 _dafny.Sequence, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC21_((func() bool {
		if true {
			return !(false)
		}
		return true
	})(), !(!((true) && (true))), ((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("td")).Cardinality()))))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return (Companion_D19_.Create_DC52_(_dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.SetOf(!(false), true, true)).Cardinality()), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('j'))).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _21_v0 _dafny.CodePoint
			_21_v0 = interface{}(_compr_6).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('j')), _21_v0) {
				_coll6.Add(_21_v0, false)
			}
		}
		return _coll6.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(965))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfInt64(719))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(461))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))).Cardinality()))), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vo")).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(69)), _dafny.MultiSetOf(_dafny.IntOfInt64(856), (_dafny.MultiSetOf(true)).Cardinality())))).Dtor_cf91()
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	var _source0 D17 = Companion_D17_.Create_DC47_(false, _dafny.IntOfInt64(897))
	_ = _source0
	if _source0.Is_DC47() {
		var _23___mcc_h0 bool = _source0.Get_().(D17_DC47).Cf81
		_ = _23___mcc_h0
		var _24___mcc_h1 _dafny.Int = _source0.Get_().(D17_DC47).Cf82
		_ = _24___mcc_h1
		var _25_cf82 _dafny.Int = _24___mcc_h1
		_ = _25_cf82
		var _26_cf81 bool = _23___mcc_h0
		_ = _26_cf81
		return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_cf81, _dafny.IntOfUint32((_dafny.SeqOf(_26_cf81)).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_cf81, _25_cf82))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_cf81, _25_cf82))))
	} else {
		var _27___mcc_h2 _dafny.Map = _source0.Get_().(D17_DC46).Cf80
		_ = _27___mcc_h2
		var _28_cf80 _dafny.Map = _27___mcc_h2
		_ = _28_cf80
		return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
				var _coll8 = _dafny.NewBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(869), _dafny.IntOfInt64(577))); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _29_v1 _dafny.Int
					_29_v1 = interface{}(_compr_8).(_dafny.Int)
					if ((_dafny.IntOfInt64(869)).Cmp(_29_v1) <= 0) && ((_29_v1).Cmp(_dafny.IntOfInt64(577)) < 0) {
						_coll8.Add(Companion_Default___.SafeModuloInt(_29_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(569))).Cardinality()))
					}
				}
				return _coll8.ToSet()
			}(), _dafny.IntOfInt64(89))).Keys().Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _30_v0 _dafny.Set
				_30_v0 = interface{}(_compr_7).(_dafny.Set)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
					var _coll9 = _dafny.NewBuilder()
					_ = _coll9
					for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(869), _dafny.IntOfInt64(577))); ; {
						_compr_9, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _31_v1 _dafny.Int
						_31_v1 = interface{}(_compr_9).(_dafny.Int)
						if ((_dafny.IntOfInt64(869)).Cmp(_31_v1) <= 0) && ((_31_v1).Cmp(_dafny.IntOfInt64(577)) < 0) {
							_coll9.Add(Companion_Default___.SafeModuloInt(_31_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(569))).Cardinality()))
						}
					}
					return _coll9.ToSet()
				}(), _dafny.IntOfInt64(89))).Contains(_30_v0) {
					_coll7.Add(_30_v0, !(false))
				}
			}
			return _coll7.ToMap()
		}()).Cardinality()))).Intersection(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())))
	}
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true)).Intersection((_dafny.SetOf(true)).Union(_dafny.SetOf(false, false, true)))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, globalState *GlobalState) {
	var _32_v0 D2
	_ = _32_v0
	_32_v0 = Companion_D2_.Create_DC6_(p0)
	var _33_v1 D2
	_ = _33_v1
	_33_v1 = Companion_D2_.Create_DC8_(_32_v0)
	var _34_v2 D2
	_ = _34_v2
	_34_v2 = Companion_D2_.Create_DC8_(_33_v1)
	var _35_v3 D2
	_ = _35_v3
	_35_v3 = Companion_D2_.Create_DC8_(_33_v1)
	(globalState).F16 = func(_source1 D2) bool {
		if _source1.Is_DC7() {
			var _36___mcc_h0 _dafny.CodePoint = _source1.Get_().(D2_DC7).Cf9
			_ = _36___mcc_h0
			var _37___mcc_h1 _dafny.Set = _source1.Get_().(D2_DC7).Cf10
			_ = _37___mcc_h1
			var _38_cf10 _dafny.Set = _37___mcc_h1
			_ = _38_cf10
			var _39_cf9 _dafny.CodePoint = _36___mcc_h0
			_ = _39_cf9
			return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(false)))).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.SeqOf(!(false))))
		} else if _source1.Is_DC6() {
			var _40___mcc_h2 _dafny.Int = _source1.Get_().(D2_DC6).Cf8
			_ = _40___mcc_h2
			var _41_cf8 _dafny.Int = _40___mcc_h2
			_ = _41_cf8
			return false
		} else {
			var _42___mcc_h3 D2 = _source1.Get_().(D2_DC8).Cf11
			_ = _42___mcc_h3
			var _43_cf11 D2 = _42___mcc_h3
			_ = _43_cf11
			return false
		}
	}(_35_v3)
	var _44_v4 _dafny.Sequence
	_ = _44_v4
	_44_v4 = _dafny.UnicodeSeqOfUtf8Bytes("m")
	var _45_v5 bool
	_ = _45_v5
	_45_v5 = true
	var _46_v6 D3
	_ = _46_v6
	_46_v6 = Companion_D3_.Create_DC10_(_44_v4, p0, _45_v5)
	var _47_v7 _dafny.Sequence
	_ = _47_v7
	_47_v7 = _dafny.SeqOf(Companion_D3_.Create_DC10_(_44_v4, p0, _45_v5), _46_v6, _46_v6, _46_v6, _46_v6)
	var _48_i0 _dafny.Int
	_ = _48_i0
	_48_i0 = _dafny.Zero
	{
		for !_dafny.Companion_Sequence_.Equal(_47_v7, _dafny.SeqOf(Companion_D3_.Create_DC10_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-571))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_112_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})), p0, _45_v5))) {
			{
				if (_48_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_48_i0 = (_48_i0).Plus(_dafny.One)
				var _49_v8 _dafny.CodePoint
				_ = _49_v8
				_49_v8 = _dafny.CodePoint('k')
				var _50_v9 _dafny.MultiSet
				_ = _50_v9
				_50_v9 = _dafny.MultiSetOf(_49_v8)
				var _51_v10 D14
				_ = _51_v10
				_51_v10 = Companion_D14_.Create_DC40_(_45_v5, _45_v5, (_50_v9).Cardinality())
				var _52_v11 D14
				_ = _52_v11
				_52_v11 = Companion_D14_.Create_DC42_(_51_v10)
				var _53_v12 D14
				_ = _53_v12
				_53_v12 = Companion_D14_.Create_DC42_(_52_v11)
				var _source2 D14 = _53_v12
				_ = _source2
				if _source2.Is_DC40() {
					var _54___mcc_h4 bool = _source2.Get_().(D14_DC40).Cf69
					_ = _54___mcc_h4
					var _55___mcc_h5 bool = _source2.Get_().(D14_DC40).Cf70
					_ = _55___mcc_h5
					var _56___mcc_h6 _dafny.Int = _source2.Get_().(D14_DC40).Cf71
					_ = _56___mcc_h6
					var _57_cf71 _dafny.Int = _56___mcc_h6
					_ = _57_cf71
					var _58_cf70 bool = _55___mcc_h5
					_ = _58_cf70
					var _59_cf69 bool = _54___mcc_h4
					_ = _59_cf69
					_49_v8 = _49_v8
					var _60_v13 *C0
					_ = _60_v13
					var _nw0 *C0 = New_C0_()
					_ = _nw0
					_nw0.Ctor__(true)
					_60_v13 = _nw0
					var _61_v14 _dafny.Array
					_ = _61_v14
					var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
					_ = _nw1
					_61_v14 = _nw1
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_61_v14), 0))
					_ = _index0
					(_61_v14).ArraySet1(p0, (_index0).Int())
					var _62_v15 _dafny.MultiSet
					_ = _62_v15
					_62_v15 = _dafny.MultiSetOf(p0)
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_61_v14), 0))
					_ = _index1
					(_61_v14).ArraySet1((((_dafny.Zero).Minus((_62_v15).Cardinality())).Minus(p0)).Plus((_dafny.Zero).Minus(p0)), (_index1).Int())
					var _63_v16 _dafny.Map
					_ = _63_v16
					_63_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_58_cf70), _45_v5)
					(globalState).F28 = (_57_cf71).Cmp((_dafny.Zero).Minus((_63_v16).Cardinality())) <= 0
				} else if _source2.Is_DC41() {
					var _64___mcc_h7 _dafny.Int = _source2.Get_().(D14_DC41).Cf72
					_ = _64___mcc_h7
					var _65___mcc_h8 _dafny.Int = _source2.Get_().(D14_DC41).Cf73
					_ = _65___mcc_h8
					var _66_cf73 _dafny.Int = _65___mcc_h8
					_ = _66_cf73
					var _67_cf72 _dafny.Int = _64___mcc_h7
					_ = _67_cf72
					var _68_v17 _dafny.Array
					_ = _68_v17
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_0
					var _nw2 _dafny.Array
					_ = _nw2
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw2 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) _dafny.Sequence = (func(_69_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_70_i2 _dafny.Int) _dafny.Sequence {
								return _69_v4
							}
						})(_44_v4)
						_ = _init0
						var _element0_0 = _init0(_dafny.Zero)
						_ = _element0_0
						_nw2 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
						(_nw2).ArraySet1(_element0_0, 0)
						var _nativeLen0_0 = (_len0_0).Int()
						_ = _nativeLen0_0
						for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
							(_nw2).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
						}
					}
					_68_v17 = _nw2
					var _71_v18 _dafny.MultiSet
					_ = _71_v18
					_71_v18 = _dafny.MultiSetOf(_44_v4)
					var _72_v19 _dafny.Map
					_ = _72_v19
					_72_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v17, _71_v18)
					_72_v19 = (_72_v19).Update(_68_v17, (_71_v18).Update(_44_v4, Companion_Default___.Abs(_dafny.IntOfInt64(-699))))
					_66_cf73 = _dafny.IntOfInt64(312)
					var _73_v20 _dafny.Array
					_ = _73_v20
					var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
					_ = _nw3
					_73_v20 = _nw3
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_73_v20), 0))
					_ = _index2
					(_73_v20).ArraySet1(_dafny.IntOfInt64(-670), (_index2).Int())
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_73_v20), 0))
					_ = _index3
					(_73_v20).ArraySet1(_66_cf73, (_index3).Int())
					var _74_v21 _dafny.Map
					_ = _74_v21
					_74_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v5, !(_45_v5) || (_45_v5))
					_74_v21 = (_74_v21).Update(_45_v5, _45_v5)
				} else if _source2.Is_DC39() {
					var _75___mcc_h9 _dafny.Array = _source2.Get_().(D14_DC39).Cf68
					_ = _75___mcc_h9
					var _76_cf68 _dafny.Array = _75___mcc_h9
					_ = _76_cf68
					_44_v4 = _dafny.Companion_Sequence_.Concatenate(_44_v4, _44_v4)
					var _77_v22 _dafny.Map
					_ = _77_v22
					_77_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _45_v5)
					var _78_v23 _dafny.Map
					_ = _78_v23
					_78_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _77_v22)
					var _79_v24 _dafny.Map
					_ = _79_v24
					_79_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
					var _80_v25 _dafny.Sequence
					_ = _80_v25
					_80_v25 = _dafny.SeqOf(_44_v4)
					var _81_v26 *C3
					_ = _81_v26
					var _nw4 *C3 = New_C3_()
					_ = _nw4
					_nw4.Ctor__(_76_cf68, _78_v23, Companion_Default___.Fm26(_45_v5, _45_v5, _45_v5, (func() _dafny.Int {
						if (_79_v24).Contains(p0) {
							return (_79_v24).Get(p0).(_dafny.Int)
						}
						return _dafny.IntOfUint32(((_80_v25).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_80_v25).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
					})(), globalState))
					_81_v26 = _nw4
					var _82_v27 _dafny.Set
					_ = _82_v27
					_82_v27 = _dafny.SetOf(p0)
					var _83_v28 _dafny.Map
					_ = _83_v28
					_83_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v5, p0)
					var _84_v29 _dafny.Map
					_ = _84_v29
					_84_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_v26, ((_82_v27).Cardinality()).Minus((_83_v28).Cardinality()))
					_84_v29 = (_84_v29).Update(_81_v26, p0)
					var _85_v30 D0
					_ = _85_v30
					_85_v30 = Companion_D0_.Create_DC1_(_49_v8, _45_v5, _45_v5)
					_85_v30 = _85_v30
					var _86_v31 _dafny.Array
					_ = _86_v31
					var _nwElement0_0 bool = _45_v5
					_ = _nwElement0_0
					var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.One)
					_ = _nw5
					(_nw5).ArraySet1(_nwElement0_0, 0)
					_86_v31 = _nw5
					var _87_v32 _dafny.Map
					_ = _87_v32
					_87_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v31, p0)
					var _88_v33 _dafny.Map
					_ = _88_v33
					_88_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v31, ((_87_v32).Cardinality()).Cmp(p0) > 0)
					_88_v33 = (_88_v33).Update(_86_v31, (Companion_Default___.Fm35(p0, _44_v4, (_dafny.Zero).Minus(p0), globalState)).Contains(_83_v28))
				} else {
					var _89___mcc_h10 D14 = _source2.Get_().(D14_DC42).Cf74
					_ = _89___mcc_h10
					var _90_cf74 D14 = _89___mcc_h10
					_ = _90_cf74
					var _91_v34 _dafny.Set
					_ = _91_v34
					_91_v34 = _dafny.SetOf(_45_v5, _45_v5, _45_v5, false, _45_v5)
					var _92_v35 _dafny.Map
					_ = _92_v35
					_92_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_91_v34).Cardinality(), _45_v5)
					var _93_v36 _dafny.Sequence
					_ = _93_v36
					_93_v36 = _dafny.SeqOf(p0, p0, p0, p0, p0)
					var _94_v37 D16
					_ = _94_v37
					_94_v37 = Companion_D16_.Create_DC44_(_93_v36)
					var _95_v38 _dafny.Array
					_ = _95_v38
					var _nwElement0_1 _dafny.Int = _dafny.IntOfInt64(-58)
					_ = _nwElement0_1
					var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(23))
					_ = _nw6
					(_nw6).ArraySet1(_nwElement0_1, 0)
					(_nw6).ArraySet1(p0, 1)
					(_nw6).ArraySet1((Companion_Default___.Fm36(globalState)).Cardinality(), 2)
					(_nw6).ArraySet1(p0, 3)
					(_nw6).ArraySet1(p0, 4)
					(_nw6).ArraySet1(p0, 5)
					(_nw6).ArraySet1((_92_v35).Cardinality(), 6)
					(_nw6).ArraySet1(p0, 7)
					(_nw6).ArraySet1(_dafny.IntOfUint32(((_94_v37).Dtor_cf76()).Cardinality()), 8)
					(_nw6).ArraySet1(p0, 9)
					(_nw6).ArraySet1(_dafny.IntOfUint32((_44_v4).Cardinality()), 10)
					(_nw6).ArraySet1(p0, 11)
					(_nw6).ArraySet1(p0, 12)
					(_nw6).ArraySet1(p0, 13)
					(_nw6).ArraySet1(_dafny.IntOfUint32((_44_v4).Cardinality()), 14)
					(_nw6).ArraySet1(p0, 15)
					(_nw6).ArraySet1(p0, 16)
					(_nw6).ArraySet1(_dafny.IntOfUint32((_93_v36).Cardinality()), 17)
					(_nw6).ArraySet1(p0, 18)
					(_nw6).ArraySet1(p0, 19)
					(_nw6).ArraySet1(p0, 20)
					(_nw6).ArraySet1(p0, 21)
					(_nw6).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(635))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg16 _dafny.Int) interface{} {
							return coer16(arg16)
						}
					}(func(_96_i3 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('g')
					}))).Cardinality())), 22)
					_95_v38 = _nw6
					var _97_v39 _dafny.Array
					_ = _97_v39
					_97_v39 = _95_v38
					(globalState).F24 = (_97_v39)
					var _98_v40 _dafny.Map
					_ = _98_v40
					_98_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_93_v36, _93_v36), false)
					_98_v40 = (_98_v40).Update(_93_v36, (_dafny.IntOfInt64(29)).Cmp(p0) > 0)
					var _99_v41 T0
					_ = _99_v41
					var _nw7 *C0 = New_C0_()
					_ = _nw7
					_nw7.Ctor__(!(_45_v5))
					_99_v41 = _nw7
					var _100_v42 _dafny.Map
					_ = _100_v42
					_100_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v5, _99_v41)
					var _101_v43 _dafny.Set
					_ = _101_v43
					_101_v43 = _dafny.SetOf(p0)
					var _102_v44 D12
					_ = _102_v44
					_102_v44 = Companion_D12_.Create_DC32_(true, (_99_v41).F31(), _99_v41, _101_v43, p0)
					var _103_v45 _dafny.Map
					_ = _103_v45
					_103_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _99_v41)
					var _104_v46 _dafny.Array
					_ = _104_v46
					var _nwElement0_2 T0 = _99_v41
					_ = _nwElement0_2
					var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(26))
					_ = _nw8
					(_nw8).ArraySet1(_nwElement0_2, 0)
					(_nw8).ArraySet1(_99_v41, 1)
					(_nw8).ArraySet1(_99_v41, 2)
					(_nw8).ArraySet1(_99_v41, 3)
					(_nw8).ArraySet1(_99_v41, 4)
					(_nw8).ArraySet1(_99_v41, 5)
					(_nw8).ArraySet1(_99_v41, 6)
					(_nw8).ArraySet1((func() T0 {
						if (_100_v42).Contains((_99_v41).F31()) {
							return (_100_v42).Get((_99_v41).F31()).(T0)
						}
						return _99_v41
					})(), 7)
					(_nw8).ArraySet1(_99_v41, 8)
					(_nw8).ArraySet1(_99_v41, 9)
					(_nw8).ArraySet1(_99_v41, 10)
					(_nw8).ArraySet1(_99_v41, 11)
					(_nw8).ArraySet1(_99_v41, 12)
					(_nw8).ArraySet1(_99_v41, 13)
					(_nw8).ArraySet1(_99_v41, 14)
					(_nw8).ArraySet1(_99_v41, 15)
					(_nw8).ArraySet1(_99_v41, 16)
					(_nw8).ArraySet1(_99_v41, 17)
					(_nw8).ArraySet1(_99_v41, 18)
					(_nw8).ArraySet1(_99_v41, 19)
					(_nw8).ArraySet1((_102_v44).Dtor_cf57(), 20)
					(_nw8).ArraySet1(_99_v41, 21)
					(_nw8).ArraySet1(_99_v41, 22)
					(_nw8).ArraySet1((func() T0 {
						if (_103_v45).Contains(p0) {
							return (_103_v45).Get(p0).(T0)
						}
						return _99_v41
					})(), 23)
					(_nw8).ArraySet1(_99_v41, 24)
					(_nw8).ArraySet1(_99_v41, 25)
					_104_v46 = _nw8
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_104_v46), 0))
					_ = _index4
					(_104_v46).ArraySet1(_99_v41, (_index4).Int())
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(10), _dafny.ArrayLen((_104_v46), 0))
					_ = _index5
					(_104_v46).ArraySet1((func() T0 {
						if !((_99_v41).F31()) {
							return _99_v41
						}
						return _99_v41
					})(), (_index5).Int())
					var _105_v47 _dafny.Int
					_ = _105_v47
					_105_v47 = _dafny.IntOfInt64(-327)
					_105_v47 = _dafny.IntOfInt64(774)
				}
				var _106_v48 _dafny.Int
				_ = _106_v48
				_106_v48 = _dafny.IntOfInt64(662)
				_106_v48 = _106_v48
				var _107_v49 *C5
				_ = _107_v49
				var _nw9 *C5 = New_C5_()
				_ = _nw9
				_nw9.Ctor__(!(!(_45_v5)), (_106_v48).Cmp(p0) > 0)
				_107_v49 = _nw9
				var _108_v50 _dafny.Map
				_ = _108_v50
				_108_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v5, _106_v48)
				var _109_v51 _dafny.Sequence
				_ = _109_v51
				_109_v51 = _dafny.SeqOf(_108_v50)
				var _110_v52 _dafny.MultiSet
				_ = _110_v52
				_110_v52 = _dafny.MultiSetOf(_108_v50)
				var _111_v53 *C5
				_ = _111_v53
				var _nw10 *C5 = New_C5_()
				_ = _nw10
				_nw10.Ctor__((_107_v49).F30(), (_110_v52).IsSubsetOf(_dafny.MultiSetFromSeq(_109_v51)))
				_111_v53 = _nw10
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _113_v54 _dafny.Array
	_ = _113_v54
	var _nw11 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
	_ = _nw11
	_113_v54 = _nw11
	for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_113_v54), 0))); ; {
		_guard_loop_0, _ok10 := _iter10()
		if !_ok10 {
			break
		}
		var _114_i4 _dafny.Int
		_114_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_114_i4).Sign() != -1) && ((_114_i4).Cmp(_dafny.ArrayLen((_113_v54), 0)) < 0)) {
			(_113_v54).ArraySet1((_114_i4).Minus(_dafny.IntOfInt64(57)), (_114_i4).Int())
		}
	}
	var _115_v55 _dafny.Array
	_ = _115_v55
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_1
	var _nw12 _dafny.Array
	_ = _nw12
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw12 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.CodePoint = func(_116_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw12 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw12).ArraySet1CodePoint(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw12).ArraySet1CodePoint(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_115_v55 = _nw12
	var _117_v56 _dafny.CodePoint
	_ = _117_v56
	_117_v56 = _dafny.CodePoint('m')
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_115_v55), 0))
	_ = _index6
	(_115_v55).ArraySet1CodePoint(_117_v56, (_index6).Int())
	var _118_v57 _dafny.MultiSet
	_ = _118_v57
	_118_v57 = _dafny.MultiSetOf(_45_v5)
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_115_v55), 0))
	_ = _index7
	(_115_v55).ArraySet1CodePoint(Companion_Default___.Fm5(_45_v5, ((_118_v57).Cardinality()).Times(_dafny.IntOfInt64(479)), globalState), (_index7).Int())
	(globalState).F28 = !_dafny.Companion_Sequence_.Equal(_44_v4, _44_v4)
	(globalState).F16 = ((func() bool {
		if _45_v5 {
			return _45_v5
		}
		return _45_v5
	})()) == (_45_v5)
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _119_v0 _dafny.CodePoint
	_ = _119_v0
	_119_v0 = _dafny.CodePoint('v')
	var _120_v1 _dafny.Array
	_ = _120_v1
	var _nwElement0_3 _dafny.CodePoint = _119_v0
	_ = _nwElement0_3
	var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(3))
	_ = _nw13
	(_nw13).ArraySet1CodePoint(_nwElement0_3, 0)
	(_nw13).ArraySet1CodePoint(_119_v0, 1)
	(_nw13).ArraySet1CodePoint(_119_v0, 2)
	_120_v1 = _nw13
	var _121_v2 _dafny.Sequence
	_ = _121_v2
	_121_v2 = _dafny.UnicodeSeqOfUtf8Bytes("mau")
	var _122_v3 bool
	_ = _122_v3
	_122_v3 = false
	var _123_v4 _dafny.Map
	_ = _123_v4
	_123_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v2, _122_v3)
	var _124_v5 _dafny.Array
	_ = _124_v5
	var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
	_ = _nw14
	_124_v5 = _nw14
	var _125_v6 _dafny.Map
	_ = _125_v6
	_125_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-850), _124_v5)
	var _126_v7 _dafny.Int
	_ = _126_v7
	_126_v7 = _dafny.IntOfInt64(159)
	var _127_v8 _dafny.Sequence
	_ = _127_v8
	_127_v8 = _dafny.SeqOf((func() _dafny.Array {
		if (_125_v6).Contains(_126_v7) {
			return (_125_v6).Get(_126_v7).(_dafny.Array)
		}
		return _124_v5
	})())
	var _128_v9 _dafny.Map
	_ = _128_v9
	_128_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _126_v7)
	var _129_v10 _dafny.Sequence
	_ = _129_v10
	_129_v10 = _dafny.SeqOf(!(_122_v3), _122_v3)
	var _130_v11 _dafny.Map
	_ = _130_v11
	_130_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_126_v7, (func() _dafny.Int {
		if (_128_v9).Contains(_122_v3) {
			return (_128_v9).Get(_122_v3).(_dafny.Int)
		}
		return (_128_v9).Cardinality()
	})()), (_dafny.MultiSetFromSeq(_129_v10)).Cardinality())
	var _131_v12 _dafny.Array
	_ = _131_v12
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(6)
	_ = _len0_2
	var _nw15 _dafny.Array
	_ = _nw15
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw15 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Sequence = (func(_132_v10 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_133_i0 _dafny.Int) _dafny.Sequence {
				return _132_v10
			}
		})(_129_v10)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw15 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw15).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw15).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_131_v12 = _nw15
	var _134_v13 D0
	_ = _134_v13
	_134_v13 = Companion_D0_.Create_DC0_(_131_v12)
	var _135_v14 _dafny.Array
	_ = _135_v14
	var _nwElement0_4 bool = _122_v3
	_ = _nwElement0_4
	var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(12))
	_ = _nw16
	(_nw16).ArraySet1(_nwElement0_4, 0)
	(_nw16).ArraySet1(!(_122_v3), 1)
	(_nw16).ArraySet1(false, 2)
	(_nw16).ArraySet1(_122_v3, 3)
	(_nw16).ArraySet1(_122_v3, 4)
	(_nw16).ArraySet1(_122_v3, 5)
	(_nw16).ArraySet1(false, 6)
	(_nw16).ArraySet1(false, 7)
	(_nw16).ArraySet1(_122_v3, 8)
	(_nw16).ArraySet1(_122_v3, 9)
	(_nw16).ArraySet1(_122_v3, 10)
	(_nw16).ArraySet1(_122_v3, 11)
	_135_v14 = _nw16
	var _136_v15 _dafny.Array
	_ = _136_v15
	_136_v15 = _124_v5
	var _137_v16 _dafny.Map
	_ = _137_v16
	_137_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _dafny.SetOf(_122_v3, _122_v3, _122_v3, false, _122_v3))
	var _138_globalState *GlobalState
	_ = _138_globalState
	var _nw17 *GlobalState = New_GlobalState_()
	_ = _nw17
	_nw17.Ctor__(_120_v1, false, false, _123_v4, false, true, _121_v2, _dafny.IntOfInt64(267), false, true, true, _127_v8, _130_v11, _dafny.IntOfInt64(175), true, _dafny.IntOfInt64(87), false, false, (_134_v13).Dtor_cf0(), true, _135_v14, (_136_v15), false, (_137_v16).Merge(_137_v16), _124_v5, _dafny.IntOfInt64(21), _dafny.IntOfInt64(-144), _135_v14, true)
	_138_globalState = _nw17
	var _139_v17 D0
	_ = _139_v17
	_139_v17 = Companion_D0_.Create_DC3_(_119_v0)
	var _140_v18 _dafny.Set
	_ = _140_v18
	_140_v18 = _dafny.SetOf(_139_v17)
	var _141_v19 _dafny.Sequence
	_ = _141_v19
	_141_v19 = _dafny.SeqOf(_126_v7, _126_v7, (_126_v7).Times((_140_v18).Cardinality()), _126_v7)
	_126_v7 = (_141_v19).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_141_v19).Cardinality()))).Uint32()).(_dafny.Int)
	var _hi0 _dafny.Int = (_126_v7).Times(_dafny.IntOfUint32((_121_v2).Cardinality()))
	_ = _hi0
	for _142_i1 := Companion_Default___.SafeModuloInt(_126_v7, _126_v7); _142_i1.Cmp(_hi0) < 0; _142_i1 = _142_i1.Plus(_dafny.One) {
		var _143_v20 _dafny.Set
		_ = _143_v20
		_143_v20 = _dafny.SetOf(_122_v3, true)
		var _144_v21 _dafny.Map
		_ = _144_v21
		_144_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, (_143_v20).Cardinality())
		Companion_Default___.M0((func() _dafny.Int {
			if (_144_v21).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_135_v14, _142_i1)).Cardinality()) {
				return (_144_v21).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_135_v14, _142_i1)).Cardinality()).(_dafny.Int)
			}
			return _126_v7
		})(), _138_globalState)
		var _145_v22 D0
		_ = _145_v22
		_145_v22 = Companion_D0_.Create_DC2_(_122_v3)
		var _pat_let_tv0 = _129_v10
		_ = _pat_let_tv0
		var _pat_let_tv1 = _129_v10
		_ = _pat_let_tv1
		var _rhs0 D0 = func(_pat_let0_0 D0) D0 {
			return func(_146_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 bool) D0 {
					return func(_147_dt__update_hcf4_h0 bool) D0 {
						return Companion_D0_.Create_DC2_(_147_dt__update_hcf4_h0)
					}(_pat_let1_0)
				}(_dafny.Companion_Sequence_.IsProperPrefixOf(_pat_let_tv0, _pat_let_tv1))
			}(_pat_let0_0)
		}(_145_v22)
		_ = _rhs0
		var _rhs1 bool = _122_v3
		_ = _rhs1
		var _rhs2 _dafny.Array = _135_v14
		_ = _rhs2
		var _lhs0 *GlobalState = _138_globalState
		_ = _lhs0
		var _lhs1 *GlobalState = _138_globalState
		_ = _lhs1
		_145_v22 = _rhs0
		_lhs0.F16 = _rhs1
		_lhs1.F20 = _rhs2
		var _source3 _dafny.Array = (func() _dafny.Array {
			if _122_v3 {
				return _136_v15
			}
			return _124_v5
		})()
		_ = _source3
		var _148___mcc_h0 _dafny.Array = _source3
		_ = _148___mcc_h0
		var _149_cf7 _dafny.Array = _148___mcc_h0
		_ = _149_cf7
		_122_v3 = (((_128_v9).Cardinality()).Times(_142_i1)).Cmp(_126_v7) == 0
		_134_v13 = _134_v13
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_149_cf7), 0))
		_ = _index8
		(_149_cf7).ArraySet1(_126_v7, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_149_cf7), 0))
		_ = _index9
		(_149_cf7).ArraySet1(_126_v7, (_index9).Int())
		_122_v3 = !(_122_v3) || (_122_v3)
		var _150_v23 D2
		_ = _150_v23
		_150_v23 = Companion_D2_.Create_DC6_(_126_v7)
		(_138_globalState).F28 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_127_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_121_v2, (Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_121_v2).Cardinality()))).Uint32(), _dafny.CodePoint('q')), (Companion_Default___.SafeIndex((_150_v23).Dtor_cf8(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_121_v2, (Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_121_v2).Cardinality()))).Uint32(), _dafny.CodePoint('q'))).Cardinality()))).Uint32(), _119_v0)).Cardinality()), _dafny.IntOfUint32((_127_v8).Cardinality()))).Uint32(), _124_v5), _dafny.Companion_Sequence_.Concatenate(_127_v8, _127_v8))
	}
	for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_135_v14), 0))); ; {
		_guard_loop_1, _ok11 := _iter11()
		if !_ok11 {
			break
		}
		var _151_i2 _dafny.Int
		_151_i2 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_151_i2).Sign() != -1) && ((_151_i2).Cmp(_dafny.ArrayLen((_135_v14), 0)) < 0)) {
			(_135_v14).ArraySet1(false, (_151_i2).Int())
		}
	}
	var _152_v24 _dafny.Map
	_ = _152_v24
	_152_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, _126_v7)
	var _153_v25 _dafny.Map
	_ = _153_v25
	_153_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, (func() _dafny.Int {
		if (_152_v24).Contains(_126_v7) {
			return (_152_v24).Get(_126_v7).(_dafny.Int)
		}
		return _126_v7
	})())
	(_138_globalState).F16 = !((_153_v25).Equals(_153_v25))
	if false {
		var _154_v26 _dafny.Map
		_ = _154_v26
		_154_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _122_v3)
		_126_v7 = (_154_v26).Cardinality()
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_124_v5), 0))
		_ = _index10
		(_124_v5).ArraySet1(_126_v7, (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_124_v5), 0))
		_ = _index11
		var _rhs3 bool = !(_122_v3) || (_122_v3)
		_ = _rhs3
		var _rhs4 _dafny.Int = (_dafny.Zero).Minus(_126_v7)
		_ = _rhs4
		var _rhs5 _dafny.Sequence = _129_v10
		_ = _rhs5
		var _rhs6 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_126_v7))
		_ = _rhs6
		var _lhs2 _dafny.Array = _124_v5
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_124_v5), 0))
		_ = _lhs3
		_122_v3 = _rhs3
		_126_v7 = _rhs4
		_129_v10 = _rhs5
		(_lhs2).ArraySet1(_rhs6, (_lhs3).Int())
		var _155_v27 _dafny.Map
		_ = _155_v27
		_155_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_126_v7).Cmp((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_124_v5), 0))).Int()).(_dafny.Int)) >= 0, _dafny.CodePoint('u'))
		_155_v27 = (_155_v27).Update(_122_v3, _119_v0)
		_126_v7 = (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((Companion_Default___.Fm0(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, !(_122_v3))), _138_globalState)).Cardinality()), (_dafny.MultiSetOf(true)).Cardinality())).Plus((func() _dafny.Int {
			if _122_v3 {
				return _dafny.IntOfInt64(-660)
			}
			return _126_v7
		})())
		var _156_v28 _dafny.Array
		_ = _156_v28
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
		_ = _nw18
		_156_v28 = _nw18
		var _157_v29 _dafny.Array
		_ = _157_v29
		var _nwElement0_5 _dafny.Sequence = _141_v19
		_ = _nwElement0_5
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(7))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_5, 0)
		(_nw19).ArraySet1(_141_v19, 1)
		(_nw19).ArraySet1(_141_v19, 2)
		(_nw19).ArraySet1(_141_v19, 3)
		(_nw19).ArraySet1(_141_v19, 4)
		(_nw19).ArraySet1(_141_v19, 5)
		(_nw19).ArraySet1(_dafny.SeqOf(_126_v7, (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_124_v5), 0))).Int()).(_dafny.Int)), 6)
		_157_v29 = _nw19
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_156_v28), 0))
		_ = _index12
		(_156_v28).ArraySet1(_157_v29, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_156_v28), 0))
		_ = _index13
		(_156_v28).ArraySet1(_157_v29, (_index13).Int())
	} else {
		_122_v3 = _122_v3
		var _158_v30 _dafny.Set
		_ = _158_v30
		_158_v30 = _dafny.SetOf(_141_v19)
		var _159_v31 _dafny.Set
		_ = _159_v31
		_159_v31 = _dafny.SetOf((_158_v30).Cardinality(), _126_v7)
		if !((_159_v31).IsProperSubsetOf(_dafny.SetOf(_126_v7))) {
			_121_v2 = _121_v2
			var _pat_let_tv2 = _119_v0
			_ = _pat_let_tv2
			var _160_v32 _dafny.Array
			_ = _160_v32
			var _nwElement0_6 D0 = Companion_D0_.Create_DC3_(_119_v0)
			_ = _nwElement0_6
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(26))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_6, 0)
			(_nw20).ArraySet1(_139_v17, 1)
			(_nw20).ArraySet1(func(_pat_let2_0 D0) D0 {
				return func(_161_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let3_0 _dafny.CodePoint) D0 {
						return func(_162_dt__update_hcf5_h0 _dafny.CodePoint) D0 {
							return Companion_D0_.Create_DC3_(_162_dt__update_hcf5_h0)
						}(_pat_let3_0)
					}(_pat_let_tv2)
				}(_pat_let2_0)
			}(_139_v17), 2)
			(_nw20).ArraySet1(_139_v17, 3)
			(_nw20).ArraySet1((func() D0 {
				if _122_v3 {
					return _139_v17
				}
				return Companion_D0_.Create_DC3_(_dafny.CodePoint('n'))
			})(), 4)
			(_nw20).ArraySet1(_139_v17, 5)
			(_nw20).ArraySet1(_139_v17, 6)
			(_nw20).ArraySet1(_139_v17, 7)
			(_nw20).ArraySet1(Companion_D0_.Create_DC3_(_dafny.CodePoint('p')), 8)
			(_nw20).ArraySet1(_139_v17, 9)
			(_nw20).ArraySet1(_139_v17, 10)
			(_nw20).ArraySet1(_139_v17, 11)
			(_nw20).ArraySet1(_139_v17, 12)
			(_nw20).ArraySet1(_139_v17, 13)
			(_nw20).ArraySet1(_139_v17, 14)
			(_nw20).ArraySet1(_139_v17, 15)
			(_nw20).ArraySet1(_139_v17, 16)
			(_nw20).ArraySet1(_139_v17, 17)
			(_nw20).ArraySet1(_139_v17, 18)
			(_nw20).ArraySet1(_139_v17, 19)
			(_nw20).ArraySet1(Companion_D0_.Create_DC3_(_119_v0), 20)
			(_nw20).ArraySet1(_139_v17, 21)
			(_nw20).ArraySet1(_139_v17, 22)
			(_nw20).ArraySet1(_139_v17, 23)
			(_nw20).ArraySet1(_139_v17, 24)
			(_nw20).ArraySet1(_139_v17, 25)
			_160_v32 = _nw20
			_160_v32 = _160_v32
			var _163_v33 _dafny.MultiSet
			_ = _163_v33
			_163_v33 = _dafny.MultiSetOf(_126_v7, (_126_v7).Minus(_126_v7))
			var _rhs7 bool = _122_v3
			_ = _rhs7
			var _rhs8 bool = !(_122_v3) || (_122_v3)
			_ = _rhs8
			var _rhs9 _dafny.MultiSet = _dafny.MultiSetOf(_dafny.IntOfInt64(-665))
			_ = _rhs9
			var _lhs4 *GlobalState = _138_globalState
			_ = _lhs4
			var _lhs5 *GlobalState = _138_globalState
			_ = _lhs5
			_lhs4.F16 = _rhs7
			_lhs5.F28 = _rhs8
			_163_v33 = _rhs9
			_126_v7 = _126_v7
			var _164_v34 _dafny.Map
			_ = _164_v34
			_164_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _dafny.CodePoint('g'))
			var _165_v35 _dafny.MultiSet
			_ = _165_v35
			_165_v35 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_((func() _dafny.CodePoint {
				if (_164_v34).Contains(_122_v3) {
					return (_164_v34).Get(_122_v3).(_dafny.CodePoint)
				}
				return _119_v0
			})(), _122_v3, false))
			var _166_v36 _dafny.Sequence
			_ = _166_v36
			_166_v36 = _dafny.SeqOf(_165_v35)
			var _167_v37 D0
			_ = _167_v37
			_167_v37 = Companion_D0_.Create_DC1_(_119_v0, _122_v3, _122_v3)
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_135_v14), 0))
			_ = _index14
			(_135_v14).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqOf(_167_v37))).IsSubsetOf((_166_v36).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_166_v36).Cardinality()))).Uint32()).(_dafny.MultiSet)), (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_135_v14), 0))
			_ = _index15
			(_135_v14).ArraySet1((_126_v7).Cmp(Companion_Default___.SafeDivisionInt(_126_v7, _126_v7)) >= 0, (_index15).Int())
		} else {
			var _168_v38 *C2
			_ = _168_v38
			var _nw21 *C2 = New_C2_()
			_ = _nw21
			_nw21.Ctor__(_129_v10, true)
			_168_v38 = _nw21
			var _169_v39 _dafny.Map
			_ = _169_v39
			_169_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _122_v3)
			var _170_v40 _dafny.Array
			_ = _170_v40
			var _nwElement0_7 _dafny.Map = _169_v39
			_ = _nwElement0_7
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(26))
			_ = _nw22
			(_nw22).ArraySet1(_nwElement0_7, 0)
			(_nw22).ArraySet1(_169_v39, 1)
			(_nw22).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_122_v3), _122_v3), 2)
			(_nw22).ArraySet1(Companion_Default___.Fm30(_119_v0, _139_v17, _138_globalState), 3)
			(_nw22).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, !(_122_v3)), 4)
			(_nw22).ArraySet1(_169_v39, 5)
			(_nw22).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _122_v3), 6)
			(_nw22).ArraySet1(_169_v39, 7)
			(_nw22).ArraySet1(_169_v39, 8)
			(_nw22).ArraySet1(_169_v39, 9)
			(_nw22).ArraySet1(_169_v39, 10)
			(_nw22).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _122_v3), 11)
			(_nw22).ArraySet1((_169_v39).Update(_122_v3, _122_v3), 12)
			(_nw22).ArraySet1(_169_v39, 13)
			(_nw22).ArraySet1(_169_v39, 14)
			(_nw22).ArraySet1(_169_v39, 15)
			(_nw22).ArraySet1(_169_v39, 16)
			(_nw22).ArraySet1(_169_v39, 17)
			(_nw22).ArraySet1(_169_v39, 18)
			(_nw22).ArraySet1(_169_v39, 19)
			(_nw22).ArraySet1(_169_v39, 20)
			(_nw22).ArraySet1(_169_v39, 21)
			(_nw22).ArraySet1(_169_v39, 22)
			(_nw22).ArraySet1(_169_v39, 23)
			(_nw22).ArraySet1(_169_v39, 24)
			(_nw22).ArraySet1(_169_v39, 25)
			_170_v40 = _nw22
			var _171_v41 _dafny.Map
			_ = _171_v41
			_171_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, _122_v3)
			var _172_v42 T0
			_ = _172_v42
			var _nw23 *C3 = New_C3_()
			_ = _nw23
			_nw23.Ctor__(_170_v40, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-190), _171_v41), _122_v3)
			_172_v42 = _nw23
			_172_v42 = _172_v42
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_124_v5), 0))
			_ = _index16
			(_124_v5).ArraySet1((func() _dafny.Int {
				if _122_v3 {
					return _126_v7
				}
				return _dafny.IntOfUint32((_121_v2).Cardinality())
			})(), (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_124_v5), 0))
			_ = _index17
			(_124_v5).ArraySet1((_168_v38).Fm13(_126_v7, _138_globalState), (_index17).Int())
			(_138_globalState).F28 = false
			_122_v3 = (_172_v42).F31()
		}
		var _173_v43 _dafny.Sequence
		_ = _173_v43
		_173_v43 = _dafny.SeqOf(_129_v10)
		var _174_v44 _dafny.Sequence
		_ = _174_v44
		_174_v44 = _dafny.SeqOf(_129_v10, (_173_v43).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_173_v43).Cardinality()))).Uint32()).(_dafny.Sequence), _129_v10, _dafny.SeqOf(_122_v3))
		_126_v7 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_174_v44).Cardinality()), _126_v7)
		if _122_v3 {
			_128_v9 = (_128_v9).Update((_122_v3) && (_122_v3), _dafny.IntOfInt64(62))
			Companion_Default___.M0(_126_v7, _138_globalState)
			_126_v7 = Companion_Default___.SafeDivisionInt(_126_v7, _126_v7)
			Companion_Default___.M0(_126_v7, _138_globalState)
			var _175_v45 _dafny.Array
			_ = _175_v45
			var _nwElement0_8 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_121_v2, (Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_121_v2).Cardinality()))).Uint32(), (_121_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_121_v2).Cardinality()), _dafny.IntOfUint32((_121_v2).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_ = _nwElement0_8
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.One)
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_8, 0)
			_175_v45 = _nw24
			_175_v45 = _175_v45
		} else {
			_126_v7 = _126_v7
			_139_v17 = Companion_Default___.Fm31(_126_v7, _138_globalState)
			var _176_v46 D5
			_ = _176_v46
			_176_v46 = Companion_D5_.Create_DC14_(_159_v31, _126_v7, true, _135_v14)
			var _177_v47 _dafny.Array
			_ = _177_v47
			var _nwElement0_9 _dafny.Array = _135_v14
			_ = _nwElement0_9
			var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(7))
			_ = _nw25
			(_nw25).ArraySet1(_nwElement0_9, 0)
			(_nw25).ArraySet1(_135_v14, 1)
			(_nw25).ArraySet1((_176_v46).Dtor_cf22(), 2)
			(_nw25).ArraySet1(_135_v14, 3)
			(_nw25).ArraySet1(_135_v14, 4)
			(_nw25).ArraySet1(_135_v14, 5)
			(_nw25).ArraySet1(_135_v14, 6)
			_177_v47 = _nw25
			var _178_v48 *C1
			_ = _178_v48
			var _nw26 *C1 = New_C1_()
			_ = _nw26
			_nw26.Ctor__(_126_v7, _177_v47, _122_v3)
			_178_v48 = _nw26
			var _179_v49 _dafny.Map
			_ = _179_v49
			_179_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, _122_v3)
			var _180_v50 D9
			_ = _180_v50
			_180_v50 = Companion_D9_.Create_DC26_(_122_v3, _121_v2, _126_v7, _179_v49, _122_v3)
			var _181_v51 _dafny.Map
			_ = _181_v51
			_181_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v50, _dafny.MultiSetOf(_dafny.IntOfInt64(-674), _dafny.IntOfUint32((_141_v19).Cardinality())))
			var _182_v52 _dafny.MultiSet
			_ = _182_v52
			_182_v52 = _dafny.MultiSetOf(_126_v7)
			var _183_v53 *C5
			_ = _183_v53
			var _nw27 *C5 = New_C5_()
			_ = _nw27
			_nw27.Ctor__(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("cf"), Companion_Default___.Fm20(_121_v2, (func() _dafny.MultiSet {
				if (_181_v51).Contains(_180_v50) {
					return (_181_v51).Get(_180_v50).(_dafny.MultiSet)
				}
				return _182_v52
			})(), !(!((_129_v10).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_129_v10).Cardinality()))).Uint32()).(bool))), (_182_v52).Cardinality(), _138_globalState)), _122_v3)
			_183_v53 = _nw27
			_126_v7 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_183_v53).F29(), _119_v0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _119_v0))).Cardinality()
		}
		var _184_v54 _dafny.MultiSet
		_ = _184_v54
		_184_v54 = _dafny.MultiSetOf(_126_v7)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_135_v14), 0))
		_ = _index18
		(_135_v14).ArraySet1((_184_v54).Contains(_126_v7), (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_135_v14), 0))
		_ = _index19
		(_135_v14).ArraySet1(_122_v3, (_index19).Int())
		var _185_v55 *C5
		_ = _185_v55
		var _nw28 *C5 = New_C5_()
		_ = _nw28
		_nw28.Ctor__(_122_v3, _122_v3)
		_185_v55 = _nw28
		var _186_v56 _dafny.Map
		_ = _186_v56
		_186_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _185_v55)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_135_v14), 0))
		_ = _index20
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_135_v14), 0))
		_ = _index21
		var _rhs10 bool = ((_126_v7).Minus(_dafny.IntOfUint32((_121_v2).Cardinality()))).Cmp(Companion_Default___.SafeModuloInt(_126_v7, _126_v7)) < 0
		_ = _rhs10
		var _rhs11 _dafny.Int = (((_186_v56).Merge((_186_v56).Update((_185_v55).F30(), _185_v55))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10(_119_v0, (_185_v55).F30(), _138_globalState), _185_v55)).Merge(_186_v56))).Cardinality()
		_ = _rhs11
		var _rhs12 bool = (_185_v55).F29()
		_ = _rhs12
		var _rhs13 _dafny.Array = _135_v14
		_ = _rhs13
		var _lhs6 _dafny.Array = _135_v14
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_135_v14), 0))
		_ = _lhs7
		var _lhs8 _dafny.Array = _135_v14
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_135_v14), 0))
		_ = _lhs9
		var _lhs10 *GlobalState = _138_globalState
		_ = _lhs10
		(_lhs6).ArraySet1(_rhs10, (_lhs7).Int())
		_126_v7 = _rhs11
		(_lhs8).ArraySet1(_rhs12, (_lhs9).Int())
		_lhs10.F20 = _rhs13
	}
	Companion_Default___.M0(Companion_Default___.SafeDivisionInt(_126_v7, _126_v7), _138_globalState)
	var _187_i3 _dafny.Int
	_ = _187_i3
	_187_i3 = _dafny.Zero
	{
		for _122_v3 {
			{
				if (_187_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_187_i3 = (_187_i3).Plus(_dafny.One)
				var _188_v57 *C0
				_ = _188_v57
				var _nw29 *C0 = New_C0_()
				_ = _nw29
				_nw29.Ctor__((_126_v7).Cmp(_126_v7) < 0)
				_188_v57 = _nw29
				_126_v7 = _126_v7
				_126_v7 = _dafny.IntOfInt64(796)
				Companion_Default___.M0(_126_v7, _138_globalState)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_124_v5), 0))); ; {
		_guard_loop_2, _ok12 := _iter12()
		if !_ok12 {
			break
		}
		var _189_i4 _dafny.Int
		_189_i4 = interface{}(_guard_loop_2).(_dafny.Int)
		if (true) && (((_189_i4).Sign() != -1) && ((_189_i4).Cmp(_dafny.ArrayLen((_124_v5), 0)) < 0)) {
			(_124_v5).ArraySet1((_189_i4).Times(_126_v7), (_189_i4).Int())
		}
	}
	var _190_i5 _dafny.Int
	_ = _190_i5
	_190_i5 = _dafny.Zero
	{
		for _122_v3 {
			{
				if (_190_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_190_i5 = (_190_i5).Plus(_dafny.One)
				var _191_v58 _dafny.MultiSet
				_ = _191_v58
				_191_v58 = _dafny.MultiSetOf(_dafny.IntOfInt64(49))
				var _192_v59 _dafny.Sequence
				_ = _192_v59
				_192_v59 = _dafny.SeqOf(_dafny.MultiSetOf(_126_v7), _191_v58, _191_v58)
				var _193_v60 D12
				_ = _193_v60
				_193_v60 = Companion_D12_.Create_DC31_(_129_v10)
				var _pat_let_tv3 = _122_v3
				_ = _pat_let_tv3
				var _rhs14 bool = (_dafny.MultiSetOf(_126_v7)).IsSubsetOf((_192_v59).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_192_v59).Cardinality()))).Uint32()).(_dafny.MultiSet))
				_ = _rhs14
				var _rhs15 _dafny.Array = _124_v5
				_ = _rhs15
				var _rhs16 bool = _122_v3
				_ = _rhs16
				var _rhs17 _dafny.Sequence = (func(_pat_let4_0 D12) D12 {
					return func(_194_dt__update__tmp_h2 D12) D12 {
						return func(_pat_let5_0 _dafny.Sequence) D12 {
							return func(_195_dt__update_hcf54_h0 _dafny.Sequence) D12 {
								return Companion_D12_.Create_DC31_(_195_dt__update_hcf54_h0)
							}(_pat_let5_0)
						}(_dafny.SeqOf(!(_pat_let_tv3)))
					}(_pat_let4_0)
				}(_193_v60)).Dtor_cf54()
				_ = _rhs17
				var _rhs18 bool = _122_v3
				_ = _rhs18
				var _lhs11 *GlobalState = _138_globalState
				_ = _lhs11
				var _lhs12 *GlobalState = _138_globalState
				_ = _lhs12
				var _lhs13 *GlobalState = _138_globalState
				_ = _lhs13
				_lhs11.F16 = _rhs14
				_lhs12.F24 = _rhs15
				_122_v3 = _rhs16
				_129_v10 = _rhs17
				_lhs13.F16 = _rhs18
				_126_v7 = _126_v7
				var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_135_v14), 0))
				_ = _index22
				(_135_v14).ArraySet1((_126_v7).Cmp(_126_v7) > 0, (_index22).Int())
				var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_135_v14), 0))
				_ = _index23
				var _rhs19 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_129_v10, _dafny.Companion_Sequence_.Concatenate(_129_v10, _dafny.SeqOf(_122_v3, true)))
				_ = _rhs19
				var _rhs20 bool = _122_v3
				_ = _rhs20
				var _rhs21 bool = _122_v3
				_ = _rhs21
				var _lhs14 _dafny.Array = _135_v14
				_ = _lhs14
				var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_135_v14), 0))
				_ = _lhs15
				var _lhs16 *GlobalState = _138_globalState
				_ = _lhs16
				_129_v10 = _rhs19
				(_lhs14).ArraySet1(_rhs20, (_lhs15).Int())
				_lhs16.F28 = _rhs21
				Companion_Default___.M0(_126_v7, _138_globalState)
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _196_v61 _dafny.MultiSet
	_ = _196_v61
	_196_v61 = _dafny.MultiSetOf(_122_v3)
	var _197_v62 _dafny.Sequence
	_ = _197_v62
	_197_v62 = _dafny.SeqOf(_dafny.MultiSetOf(_122_v3, _122_v3, _122_v3, _122_v3), (_196_v61).Update(_122_v3, Companion_Default___.Abs(_126_v7)))
	var _198_v63 D5
	_ = _198_v63
	_198_v63 = Companion_D5_.Create_DC13_((_197_v62).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_197_v62).Cardinality()))).Uint32()).(_dafny.MultiSet))
	var _source4 D5 = _198_v63
	_ = _source4
	if _source4.Is_DC14() {
		var _199___mcc_h1 _dafny.Set = _source4.Get_().(D5_DC14).Cf19
		_ = _199___mcc_h1
		var _200___mcc_h2 _dafny.Int = _source4.Get_().(D5_DC14).Cf20
		_ = _200___mcc_h2
		var _201___mcc_h3 bool = _source4.Get_().(D5_DC14).Cf21
		_ = _201___mcc_h3
		var _202___mcc_h4 _dafny.Array = _source4.Get_().(D5_DC14).Cf22
		_ = _202___mcc_h4
		var _203_cf22 _dafny.Array = _202___mcc_h4
		_ = _203_cf22
		var _204_cf21 bool = _201___mcc_h3
		_ = _204_cf21
		var _205_cf20 _dafny.Int = _200___mcc_h2
		_ = _205_cf20
		var _206_cf19 _dafny.Set = _199___mcc_h1
		_ = _206_cf19
		_205_cf20 = (_dafny.IntOfUint32((_141_v19).Cardinality())).Times(Companion_Default___.Fm7(_204_cf21, _122_v3, _138_globalState))
		var _207_v64 *C2
		_ = _207_v64
		var _nw30 *C2 = New_C2_()
		_ = _nw30
		_nw30.Ctor__(_dafny.Companion_Sequence_.Update(_129_v10, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.IntOfUint32((_129_v10).Cardinality()))).Uint32(), _122_v3), false)
		_207_v64 = _nw30
		_207_v64 = _207_v64
		var _208_v65 D0
		_ = _208_v65
		_208_v65 = Companion_D0_.Create_DC1_(_119_v0, !(_204_cf21), Companion_Default___.Fm26(_122_v3, _122_v3, _204_cf21, _205_cf20, _138_globalState))
		var _209_v66 _dafny.Sequence
		_ = _209_v66
		_209_v66 = _dafny.SeqOf(_208_v65)
		var _210_v67 _dafny.Map
		_ = _210_v67
		_210_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ks")).Cardinality()))).Cardinality())
		var _211_v68 _dafny.Map
		_ = _211_v68
		_211_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, _119_v0)
		var _212_v69 _dafny.Array
		_ = _212_v69
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw31
		_212_v69 = _nw31
		var _213_v70 *C4
		_ = _213_v70
		var _nw32 *C4 = New_C4_()
		_ = _nw32
		_nw32.Ctor__(_209_v66, _122_v3, (_dafny.Zero).Minus((func() _dafny.Int {
			if (_210_v67).Contains((func() _dafny.CodePoint {
				if (_211_v68).Contains(_dafny.IntOfUint32((_121_v2).Cardinality())) {
					return (_211_v68).Get(_dafny.IntOfUint32((_121_v2).Cardinality())).(_dafny.CodePoint)
				}
				return _119_v0
			})()) {
				return (_210_v67).Get((func() _dafny.CodePoint {
					if (_211_v68).Contains(_dafny.IntOfUint32((_121_v2).Cardinality())) {
						return (_211_v68).Get(_dafny.IntOfUint32((_121_v2).Cardinality())).(_dafny.CodePoint)
					}
					return _119_v0
				})()).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_121_v2).Cardinality())
		})()), _212_v69, _122_v3)
		_213_v70 = _nw32
		_213_v70 = _213_v70
		var _214_v71 _dafny.Set
		_ = _214_v71
		_214_v71 = _dafny.SetOf(_213_v70.F35)
		(_138_globalState).F16 = !((_214_v71).IsSubsetOf(_214_v71)) || (_204_cf21)
	} else if _source4.Is_DC13() {
		var _215___mcc_h5 _dafny.MultiSet = _source4.Get_().(D5_DC13).Cf18
		_ = _215___mcc_h5
		var _216_cf18 _dafny.MultiSet = _215___mcc_h5
		_ = _216_cf18
		_124_v5 = _124_v5
		var _217_v72 _dafny.Array
		_ = _217_v72
		var _nw33 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
		_ = _nw33
		_217_v72 = _nw33
		var _218_v73 _dafny.Array
		_ = _218_v73
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
		_ = _nw34
		_218_v73 = _nw34
		var _219_v74 T1
		_ = _219_v74
		var _nw35 *C1 = New_C1_()
		_ = _nw35
		_nw35.Ctor__(_126_v7, _218_v73, _122_v3)
		_219_v74 = _nw35
		var _220_v75 _dafny.Sequence
		_ = _220_v75
		_220_v75 = _dafny.SeqOf(_219_v74, _219_v74)
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_217_v72), 0))
		_ = _index24
		(_217_v72).ArraySet1((_220_v75).Select((Companion_Default___.SafeIndex((_141_v19).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_141_v19).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_220_v75).Cardinality()))).Uint32()).(T1), (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_217_v72), 0))
		_ = _index25
		(_217_v72).ArraySet1(_219_v74, (_index25).Int())
		_129_v10 = _129_v10
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_218_v73), 0))
		_ = _index26
		(_218_v73).ArraySet1(_135_v14, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_218_v73), 0))
		_ = _index27
		(_218_v73).ArraySet1(_135_v14, (_index27).Int())
	} else {
		var _221___mcc_h6 D5 = _source4.Get_().(D5_DC15).Cf23
		_ = _221___mcc_h6
		var _222_cf23 D5 = _221___mcc_h6
		_ = _222_cf23
		var _223_v76 D0
		_ = _223_v76
		_223_v76 = Companion_D0_.Create_DC2_(_122_v3)
		var _224_v77 D0
		_ = _224_v77
		_224_v77 = Companion_D0_.Create_DC1_(_119_v0, (_223_v76).Dtor_cf4(), _122_v3)
		var _225_v78 _dafny.Sequence
		_ = _225_v78
		_225_v78 = _dafny.SeqOf(_224_v77)
		var _226_v79 _dafny.Map
		_ = _226_v79
		_226_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_126_v7), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(106))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_227_v19 _dafny.Sequence, _228_v3 bool, _229_v7 _dafny.Int, _230_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
			return func(_231_i6 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf((_227_v19).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC2_(_228_v3), _229_v7)).Cardinality(), _dafny.IntOfUint32((_227_v19).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(971))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}((func(_232_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_233_i7 _dafny.Int) _dafny.CodePoint {
						return _232_v0
					}
				})(_230_v0)))).Cardinality()))).Cardinality())
			}
		})(_141_v19, _122_v3, _126_v7, _119_v0))))
		var _234_v80 _dafny.Sequence
		_ = _234_v80
		_234_v80 = _dafny.SeqOf(_141_v19)
		var _235_v81 _dafny.Array
		_ = _235_v81
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
		_ = _nw36
		_235_v81 = _nw36
		var _236_v82 *C4
		_ = _236_v82
		var _nw37 *C4 = New_C4_()
		_ = _nw37
		_nw37.Ctor__(_dafny.Companion_Sequence_.Concatenate(_225_v78, _225_v78), _122_v3, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_226_v79).Contains((_234_v80).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_234_v80).Cardinality()))).Uint32()).(_dafny.Sequence)) {
				return (_226_v79).Get((_234_v80).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_234_v80).Cardinality()))).Uint32()).(_dafny.Sequence)).(_dafny.Sequence)
			}
			return _141_v19
		})()).Cardinality()), _235_v81, !(false) || (_122_v3))
		_236_v82 = _nw37
		var _237_v83 _dafny.Set
		_ = _237_v83
		_237_v83 = _dafny.SetOf(Companion_Default___.Fm1(_119_v0, _138_globalState))
		var _238_v84 _dafny.Map
		_ = _238_v84
		_238_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _236_v82.F35)
		var _rhs22 _dafny.Int = Companion_Default___.Fm7(!(_236_v82.F35) || (_122_v3), (func() bool {
			if (_238_v84).Contains(!(_122_v3)) {
				return (_238_v84).Get(!(_122_v3)).(bool)
			}
			return _236_v82.F35
		})(), _138_globalState)
		_ = _rhs22
		var _rhs23 *C4 = (func() *C4 {
			if _122_v3 {
				return _236_v82
			}
			return _236_v82
		})()
		_ = _rhs23
		var _rhs24 bool = _236_v82.F35
		_ = _rhs24
		var _rhs25 _dafny.Set = _237_v83
		_ = _rhs25
		var _rhs26 _dafny.Sequence = (func() _dafny.Sequence {
			if _122_v3 {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}((func(_239_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_240_i8 _dafny.Int) _dafny.CodePoint {
						return _239_v0
					}
				})(_119_v0)))
			}
			return _121_v2
		})()
		_ = _rhs26
		var _lhs17 *GlobalState = _138_globalState
		_ = _lhs17
		_126_v7 = _rhs22
		_236_v82 = _rhs23
		_lhs17.F28 = _rhs24
		_237_v83 = _rhs25
		_121_v2 = _rhs26
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_135_v14), 0))
		_ = _index28
		(_135_v14).ArraySet1(!_dafny.Companion_Sequence_.Equal(_121_v2, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("w"), (Companion_Default___.SafeIndex((Companion_Default___.Fm32(_126_v7, _236_v82.F35, _236_v82.F35, _236_v82.F35, _138_globalState)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality()))).Uint32(), _119_v0)), (_index28).Int())
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_135_v14), 0))
		_ = _index29
		(_135_v14).ArraySet1(_236_v82.F35, (_index29).Int())
		var _241_v85 _dafny.Int
		_ = _241_v85
		var _242_v86 _dafny.Array
		_ = _242_v86
		var _243_v87 _dafny.Sequence
		_ = _243_v87
		var _244_v88 _dafny.Int
		_ = _244_v88
		var _out0 _dafny.Int
		_ = _out0
		var _out1 _dafny.Array
		_ = _out1
		var _out2 _dafny.Sequence
		_ = _out2
		var _out3 _dafny.Int
		_ = _out3
		_out0, _out1, _out2, _out3 = (_236_v82).M2((func() bool {
			if (_135_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_135_v14), 0))).Int()).(bool) {
				return _236_v82.F35
			}
			return _122_v3
		})(), (_196_v61).Union(_196_v61), _196_v61, _138_globalState)
		_241_v85 = _out0
		_242_v86 = _out1
		_243_v87 = _out2
		_244_v88 = _out3
		var _245_v89 D3
		_ = _245_v89
		_245_v89 = Companion_D3_.Create_DC10_(_dafny.UnicodeSeqOfUtf8Bytes("eavwb"), _241_v85, false)
		var _246_v90 D3
		_ = _246_v90
		_246_v90 = Companion_D3_.Create_DC11_(_245_v89)
		(_138_globalState).F16 = (Companion_Default___.Fm33(_244_v88, _246_v90, _121_v2, _138_globalState)).Dtor_cf32()
	}
	var _247_v91 _dafny.Map
	_ = _247_v91
	_247_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_121_v2).Cardinality()), _122_v3)
	var _248_v92 _dafny.Map
	_ = _248_v92
	_248_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_247_v91).Update(_126_v7, _122_v3), _126_v7)
	var _249_v93 D13
	_ = _249_v93
	_249_v93 = Companion_D13_.Create_DC34_(_248_v92)
	var _250_v94 _dafny.Map
	_ = _250_v94
	_250_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _135_v14)
	var _251_v95 _dafny.Sequence
	_ = _251_v95
	_251_v95 = _dafny.SeqOf((func() _dafny.Array {
		if (_250_v94).Contains(_122_v3) {
			return (_250_v94).Get(_122_v3).(_dafny.Array)
		}
		return _135_v14
	})())
	var _252_v96 _dafny.MultiSet
	_ = _252_v96
	_252_v96 = _dafny.MultiSetOf(_dafny.IntOfUint32((_251_v95).Cardinality()))
	var _253_v97 _dafny.Sequence
	_ = _253_v97
	_253_v97 = _dafny.SeqOf(_252_v96)
	var _254_v98 _dafny.MultiSet
	_ = _254_v98
	_254_v98 = _dafny.MultiSetOf(_252_v96, _252_v96, _252_v96)
	(_138_globalState).F16 = ((_dafny.MultiSetFromSeq(_253_v97)).Difference(_254_v98)).IsProperSubsetOf(Companion_Default___.Fm34((_249_v93).Dtor_cf61(), _138_globalState))
	var _255_v99 _dafny.Map
	_ = _255_v99
	_255_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _141_v19)
	Companion_Default___.M0((func() _dafny.Int {
		if (_153_v25).Contains(((_255_v99).Update(_122_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(608))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_256_v3 bool) func(_dafny.Int) _dafny.Int {
			return func(_257_i9 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(Companion_D13_.Create_DC36_(_256_v3, _256_v3))).Cardinality())
			}
		})(_122_v3))))).Cardinality()) {
			return (_153_v25).Get(((_255_v99).Update(_122_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(608))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_258_v3 bool) func(_dafny.Int) _dafny.Int {
				return func(_259_i9 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(Companion_D13_.Create_DC36_(_258_v3, _258_v3))).Cardinality())
				}
			})(_122_v3))))).Cardinality()).(_dafny.Int)
		}
		return _126_v7
	})(), _138_globalState)
	var _hi1 _dafny.Int = _126_v7
	_ = _hi1
	for _260_i10 := _126_v7; _260_i10.Cmp(_hi1) < 0; _260_i10 = _260_i10.Plus(_dafny.One) {
		var _261_v100 _dafny.Array
		_ = _261_v100
		var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
		_ = _nw38
		_261_v100 = _nw38
		var _262_v101 D14
		_ = _262_v101
		_262_v101 = Companion_D14_.Create_DC39_(_261_v100)
		var _263_v102 _dafny.Map
		_ = _263_v102
		_263_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_129_v10)).Cardinality(), _247_v91)
		var _264_v103 *C3
		_ = _264_v103
		var _nw39 *C3 = New_C3_()
		_ = _nw39
		_nw39.Ctor__((_262_v101).Dtor_cf68(), _263_v102, _122_v3)
		_264_v103 = _nw39
		var _265_v104 _dafny.Set
		_ = _265_v104
		_265_v104 = _dafny.SetOf(_264_v103)
		_265_v104 = ((_265_v104).Intersection(_265_v104)).Difference((_265_v104).Difference(_265_v104))
		var _266_v105 _dafny.Set
		_ = _266_v105
		_266_v105 = _dafny.SetOf(_124_v5)
		var _rhs27 _dafny.Int = (_dafny.Zero).Minus(_126_v7)
		_ = _rhs27
		var _rhs28 _dafny.Int = (_260_i10).Plus(_126_v7)
		_ = _rhs28
		var _rhs29 _dafny.Int = _126_v7
		_ = _rhs29
		var _rhs30 _dafny.Set = (_266_v105).Difference(_266_v105)
		_ = _rhs30
		_126_v7 = _rhs27
		_126_v7 = _rhs28
		_126_v7 = _rhs29
		_266_v105 = _rhs30
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_124_v5), 0))
		_ = _index30
		(_124_v5).ArraySet1(_126_v7, (_index30).Int())
		var _267_v106 _dafny.Array
		_ = _267_v106
		var _nwElement0_10 _dafny.Array = _135_v14
		_ = _nwElement0_10
		var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(6))
		_ = _nw40
		(_nw40).ArraySet1(_nwElement0_10, 0)
		(_nw40).ArraySet1(_135_v14, 1)
		(_nw40).ArraySet1(_135_v14, 2)
		(_nw40).ArraySet1(_135_v14, 3)
		(_nw40).ArraySet1(_135_v14, 4)
		(_nw40).ArraySet1(_135_v14, 5)
		_267_v106 = _nw40
		var _268_v107 T1
		_ = _268_v107
		var _nw41 *C1 = New_C1_()
		_ = _nw41
		_nw41.Ctor__(_126_v7, _267_v106, _122_v3)
		_268_v107 = _nw41
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_124_v5), 0))
		_ = _index31
		(_124_v5).ArraySet1((_dafny.SetOf(_268_v107)).Cardinality(), (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_124_v5), 0))
		_ = _index32
		(_124_v5).ArraySet1(_dafny.IntOfUint32((_121_v2).Cardinality()), (_index32).Int())
	}
	var _269_v108 *C5
	_ = _269_v108
	var _nw42 *C5 = New_C5_()
	_ = _nw42
	_nw42.Ctor__((_129_v10).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_129_v10).Cardinality()))).Uint32()).(bool), _122_v3)
	_269_v108 = _nw42
	var _270_v109 D6
	_ = _270_v109
	_270_v109 = Companion_D6_.Create_DC16_(_120_v1)
	var _source5 D6 = _270_v109
	_ = _source5
	if _source5.Is_DC17() {
		var _271___mcc_h7 bool = _source5.Get_().(D6_DC17).Cf25
		_ = _271___mcc_h7
		var _272___mcc_h8 _dafny.Array = _source5.Get_().(D6_DC17).Cf26
		_ = _272___mcc_h8
		var _273___mcc_h9 _dafny.Int = _source5.Get_().(D6_DC17).Cf27
		_ = _273___mcc_h9
		var _274_cf27 _dafny.Int = _273___mcc_h9
		_ = _274_cf27
		var _275_cf26 _dafny.Array = _272___mcc_h8
		_ = _275_cf26
		var _276_cf25 bool = _271___mcc_h7
		_ = _276_cf25
		var _277_v110 _dafny.Map
		_ = _277_v110
		_277_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10(Companion_Default___.Fm20(_121_v2, _252_v96, (_269_v108).F30(), (_196_v61).Cardinality(), _138_globalState), (_269_v108).F29(), _138_globalState), (_269_v108).F30())
		var _278_v111 _dafny.Map
		_ = _278_v111
		_278_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm18((func() bool {
			if (_277_v110).Contains(false) {
				return (_277_v110).Get(false).(bool)
			}
			return _276_cf25
		})(), _122_v3, _138_globalState), _277_v110)
		_278_v111 = (_278_v111).Update(_126_v7, _277_v110)
		_274_cf27 = Companion_Default___.SafeDivisionInt((_126_v7).Minus(_dafny.IntOfUint32((_129_v10).Cardinality())), _274_cf27)
		Companion_Default___.M0(_126_v7, _138_globalState)
		_252_v96 = (func() _dafny.MultiSet {
			if (_269_v108).F30() {
				return _dafny.MultiSetOf(_dafny.IntOfUint32((_121_v2).Cardinality()))
			}
			return _252_v96
		})()
	} else if _source5.Is_DC16() {
		var _279___mcc_h10 _dafny.Array = _source5.Get_().(D6_DC16).Cf24
		_ = _279___mcc_h10
		var _280_cf24 _dafny.Array = _279___mcc_h10
		_ = _280_cf24
		var _281_v112 D0
		_ = _281_v112
		_281_v112 = Companion_D0_.Create_DC1_(_119_v0, (_269_v108).F30(), (_269_v108).F30())
		var _282_v113 bool
		_ = _282_v113
		var _283_v114 bool
		_ = _283_v114
		var _out4 bool
		_ = _out4
		var _out5 bool
		_ = _out5
		_out4, _out5 = (_269_v108).M1(_141_v19, _dafny.IntOfInt64(-33), _281_v112, _138_globalState)
		_282_v113 = _out4
		_283_v114 = _out5
		if _283_v114 {
			_122_v3 = !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("ei"), Companion_Default___.Fm0(_196_v61, _138_globalState))
			_121_v2 = _121_v2
			Companion_Default___.M0(_126_v7, _138_globalState)
			var _284_v115 bool
			_ = _284_v115
			var _285_v116 bool
			_ = _285_v116
			var _out6 bool
			_ = _out6
			var _out7 bool
			_ = _out7
			_out6, _out7 = (_269_v108).M1(_141_v19, _126_v7, _281_v112, _138_globalState)
			_284_v115 = _out6
			_285_v116 = _out7
			_251_v95 = _dafny.Companion_Sequence_.Concatenate(_251_v95, _dafny.Companion_Sequence_.Concatenate(_251_v95, _251_v95))
		} else {
			(_138_globalState).F20 = _135_v14
			var _pat_let_tv4 = _119_v0
			_ = _pat_let_tv4
			var _286_v117 _dafny.Array
			_ = _286_v117
			var _nwElement0_11 D0 = func(_pat_let6_0 D0) D0 {
				return func(_287_dt__update__tmp_h3 D0) D0 {
					return func(_pat_let7_0 _dafny.CodePoint) D0 {
						return func(_288_dt__update_hcf5_h1 _dafny.CodePoint) D0 {
							return Companion_D0_.Create_DC3_(_288_dt__update_hcf5_h1)
						}(_pat_let7_0)
					}(_pat_let_tv4)
				}(_pat_let6_0)
			}(_139_v17)
			_ = _nwElement0_11
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(21))
			_ = _nw43
			(_nw43).ArraySet1(_nwElement0_11, 0)
			(_nw43).ArraySet1(_139_v17, 1)
			(_nw43).ArraySet1(_139_v17, 2)
			(_nw43).ArraySet1(Companion_D0_.Create_DC3_(_119_v0), 3)
			(_nw43).ArraySet1(_139_v17, 4)
			(_nw43).ArraySet1(_139_v17, 5)
			(_nw43).ArraySet1(_139_v17, 6)
			(_nw43).ArraySet1(_139_v17, 7)
			(_nw43).ArraySet1(_139_v17, 8)
			(_nw43).ArraySet1(_139_v17, 9)
			(_nw43).ArraySet1(_139_v17, 10)
			(_nw43).ArraySet1(Companion_D0_.Create_DC3_(_119_v0), 11)
			(_nw43).ArraySet1(_139_v17, 12)
			(_nw43).ArraySet1(_139_v17, 13)
			(_nw43).ArraySet1(_139_v17, 14)
			(_nw43).ArraySet1(_139_v17, 15)
			(_nw43).ArraySet1(_139_v17, 16)
			(_nw43).ArraySet1(_139_v17, 17)
			(_nw43).ArraySet1(_139_v17, 18)
			(_nw43).ArraySet1(_139_v17, 19)
			(_nw43).ArraySet1(_139_v17, 20)
			_286_v117 = _nw43
			var _289_v118 _dafny.Sequence
			_ = _289_v118
			_289_v118 = _dafny.SeqOf(_286_v117, _286_v117, _286_v117)
			var _290_v119 _dafny.Map
			_ = _290_v119
			_290_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7(_282_v113, (_269_v108).F30(), _138_globalState), _289_v118)
			_290_v119 = (_290_v119).Update(_126_v7, _289_v118)
			var _291_v120 D3
			_ = _291_v120
			_291_v120 = Companion_D3_.Create_DC10_(_dafny.UnicodeSeqOfUtf8Bytes("soglnufp"), _126_v7, (_269_v108).F29())
			var _292_v121 _dafny.Map
			_ = _292_v121
			_292_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v120, (_269_v108).F30())
			_292_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v120, (_122_v3) || (_283_v114))
			var _293_v122 D9
			_ = _293_v122
			_293_v122 = Companion_D9_.Create_DC25_(_126_v7, (_269_v108).F30(), Companion_Default___.Fm3(_138_globalState), _283_v114)
			var _294_v123 *C5
			_ = _294_v123
			var _nw44 *C5 = New_C5_()
			_ = _nw44
			_nw44.Ctor__(!(!(false)), !(false) || ((_293_v122).Dtor_cf44()))
			_294_v123 = _nw44
			var _295_v124 *C2
			_ = _295_v124
			var _nw45 *C2 = New_C2_()
			_ = _nw45
			_nw45.Ctor__(Companion_Default___.Fm2(_282_v113, _126_v7, (_269_v108).F29(), (func() bool {
				if (_247_v91).Contains(_dafny.IntOfUint32((_121_v2).Cardinality())) {
					return (_247_v91).Get(_dafny.IntOfUint32((_121_v2).Cardinality())).(bool)
				}
				return (_269_v108).F30()
			})(), _138_globalState), (_dafny.IntOfInt64(125)).Cmp(_126_v7) == 0)
			_295_v124 = _nw45
			_295_v124 = _295_v124
		}
		_283_v114 = (_269_v108).F29()
		var _296_v125 _dafny.Array
		_ = _296_v125
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_3
		var _nw46 _dafny.Array
		_ = _nw46
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw46 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Map = (func(_297_v3 bool, _298_v108 *C5) func(_dafny.Int) _dafny.Map {
				return func(_299_i11 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_297_v3, (_298_v108).F30())
				}
			})(_122_v3, _269_v108)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw46 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw46).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw46).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_296_v125 = _nw46
		var _300_v126 _dafny.Map
		_ = _300_v126
		_300_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, _247_v91)
		var _301_v127 *C3
		_ = _301_v127
		var _nw47 *C3 = New_C3_()
		_ = _nw47
		_nw47.Ctor__(_296_v125, _300_v126, (_269_v108).F30())
		_301_v127 = _nw47
		var _302_v128 _dafny.MultiSet
		_ = _302_v128
		_302_v128 = _dafny.MultiSetOf(_301_v127, _301_v127, _301_v127, _301_v127)
		_283_v114 = Companion_Default___.Fm8((_302_v128).IsSubsetOf(_dafny.MultiSetOf(_301_v127, _301_v127)), (_126_v7).Minus(_126_v7), _138_globalState)
	} else {
		var _303___mcc_h11 D6 = _source5.Get_().(D6_DC18).Cf28
		_ = _303___mcc_h11
		var _304_cf28 D6 = _303___mcc_h11
		_ = _304_cf28
		var _305_v129 *C0
		_ = _305_v129
		var _nw48 *C0 = New_C0_()
		_ = _nw48
		_nw48.Ctor__((_269_v108).F30())
		_305_v129 = _nw48
		var _306_v130 *C0
		_ = _306_v130
		_306_v130 = _305_v129
		var _307_v131 _dafny.Sequence
		_ = _307_v131
		_307_v131 = _dafny.SeqOf(_305_v129, _305_v129)
		var _308_v132 _dafny.Array
		_ = _308_v132
		var _nwElement0_12 *C0 = _305_v129
		_ = _nwElement0_12
		var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(29))
		_ = _nw49
		(_nw49).ArraySet1(_nwElement0_12, 0)
		(_nw49).ArraySet1(_305_v129, 1)
		(_nw49).ArraySet1(_305_v129, 2)
		(_nw49).ArraySet1(_305_v129, 3)
		(_nw49).ArraySet1(_305_v129, 4)
		(_nw49).ArraySet1(_305_v129, 5)
		(_nw49).ArraySet1(_305_v129, 6)
		(_nw49).ArraySet1(_305_v129, 7)
		(_nw49).ArraySet1(_305_v129, 8)
		(_nw49).ArraySet1(_305_v129, 9)
		(_nw49).ArraySet1(_305_v129, 10)
		(_nw49).ArraySet1(_305_v129, 11)
		(_nw49).ArraySet1((_306_v130), 12)
		(_nw49).ArraySet1(_305_v129, 13)
		(_nw49).ArraySet1(_305_v129, 14)
		(_nw49).ArraySet1(_305_v129, 15)
		(_nw49).ArraySet1(_305_v129, 16)
		(_nw49).ArraySet1(_305_v129, 17)
		(_nw49).ArraySet1(_305_v129, 18)
		(_nw49).ArraySet1((_307_v131).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_126_v7), _dafny.IntOfUint32((_307_v131).Cardinality()))).Uint32()).(*C0), 19)
		(_nw49).ArraySet1(_305_v129, 20)
		(_nw49).ArraySet1(_305_v129, 21)
		(_nw49).ArraySet1(_305_v129, 22)
		(_nw49).ArraySet1(_305_v129, 23)
		(_nw49).ArraySet1((_307_v131).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_307_v131).Cardinality()))).Uint32()).(*C0), 24)
		(_nw49).ArraySet1((_306_v130), 25)
		(_nw49).ArraySet1(_305_v129, 26)
		(_nw49).ArraySet1(_305_v129, 27)
		(_nw49).ArraySet1(_305_v129, 28)
		_308_v132 = _nw49
		_126_v7 = _dafny.IntOfUint32((_dafny.SeqOf(_308_v132, _308_v132, _308_v132, _308_v132, _308_v132)).Cardinality())
		var _rhs31 bool = (_129_v10).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_129_v10).Cardinality()))).Uint32()).(bool)
		_ = _rhs31
		var _rhs32 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_126_v7, _126_v7), (_126_v7).Plus(_dafny.IntOfUint32((_129_v10).Cardinality())))
		_ = _rhs32
		var _rhs33 _dafny.Int = Companion_Default___.Fm1(_119_v0, _138_globalState)
		_ = _rhs33
		var _lhs18 *GlobalState = _138_globalState
		_ = _lhs18
		_lhs18.F16 = _rhs31
		_126_v7 = _rhs32
		_126_v7 = _rhs33
		if _122_v3 {
			var _309_v133 _dafny.Array
			_ = _309_v133
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw50
			_309_v133 = _nw50
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_309_v133), 0))
			_ = _index33
			(_309_v133).ArraySet1(_dafny.SeqOf(_135_v14), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_124_v5), 0))
			_ = _index34
			(_124_v5).ArraySet1(Companion_Default___.SafeDivisionInt((_252_v96).Cardinality(), _dafny.IntOfInt64(-566)), (_index34).Int())
			var _310_v134 _dafny.Set
			_ = _310_v134
			_310_v134 = _dafny.SetOf((_269_v108).F30())
			var _311_v135 _dafny.Map
			_ = _311_v135
			_311_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, _310_v134)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_309_v133), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_124_v5), 0))
			_ = _index36
			var _rhs34 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_251_v95, _251_v95)
			_ = _rhs34
			var _rhs35 _dafny.Int = (Companion_Default___.SafeModuloInt(_126_v7, ((func() _dafny.Set {
				if (_311_v135).Contains(_126_v7) {
					return (_311_v135).Get(_126_v7).(_dafny.Set)
				}
				return _dafny.SetOf(true)
			})()).Cardinality())).Plus((_126_v7).Plus((func() _dafny.Int {
				if (_128_v9).Contains((_269_v108).F30()) {
					return (_128_v9).Get((_269_v108).F30()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(885)
			})()))
			_ = _rhs35
			var _rhs36 _dafny.Int = Companion_Default___.Fm7((_269_v108).F29(), _122_v3, _138_globalState)
			_ = _rhs36
			var _rhs37 _dafny.Int = _dafny.IntOfInt64(643)
			_ = _rhs37
			var _rhs38 bool = (_269_v108).F29()
			_ = _rhs38
			var _lhs19 _dafny.Array = _309_v133
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_309_v133), 0))
			_ = _lhs20
			var _lhs21 _dafny.Array = _124_v5
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_124_v5), 0))
			_ = _lhs22
			(_lhs19).ArraySet1(_rhs34, (_lhs20).Int())
			_126_v7 = _rhs35
			(_lhs21).ArraySet1(_rhs36, (_lhs22).Int())
			_126_v7 = _rhs37
			_122_v3 = _rhs38
			_126_v7 = (_dafny.Zero).Minus(Companion_Default___.Fm1(_dafny.CodePoint('r'), _138_globalState))
			var _312_v136 _dafny.Sequence
			_ = _312_v136
			_312_v136 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v7, (_269_v108).F30()))
			(_138_globalState).F16 = !_dafny.Companion_Sequence_.Contains(_312_v136, _247_v91)
			(_138_globalState).F16 = !(!((func() bool {
				if ((_269_v108).F30()) == ((_269_v108).F30()) {
					return ((_269_v108).F29()) && ((_269_v108).F30())
				}
				return !(!((_269_v108).F30()))
			})()))
			var _rhs39 _dafny.Int = _126_v7
			_ = _rhs39
			var _rhs40 _dafny.Array = ((_309_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_309_v133), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32(((_309_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_309_v133), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs40
			var _rhs41 _dafny.Sequence = _121_v2
			_ = _rhs41
			var _rhs42 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_121_v2, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_128_v9).Cardinality()), _dafny.IntOfUint32((_121_v2).Cardinality()))).Uint32(), _119_v0)
			_ = _rhs42
			var _lhs23 *GlobalState = _138_globalState
			_ = _lhs23
			_126_v7 = _rhs39
			_lhs23.F27 = _rhs40
			_121_v2 = _rhs41
			_121_v2 = _rhs42
		} else {
			var _313_v137 D3
			_ = _313_v137
			_313_v137 = Companion_D3_.Create_DC10_(_121_v2, _126_v7, _122_v3)
			_126_v7 = (_313_v137).Dtor_cf14()
			_126_v7 = _dafny.IntOfInt64(915)
			Companion_Default___.M0(_126_v7, _138_globalState)
			var _314_v138 _dafny.Sequence
			_ = _314_v138
			_314_v138 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-779))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_315_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_316_i12 _dafny.Int) _dafny.CodePoint {
					return _315_v0
				}
			})(_119_v0))))
			_126_v7 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_269_v108).F29() {
					return _314_v138
				}
				return _314_v138
			})(), _314_v138), (Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_269_v108).F29() {
					return _314_v138
				}
				return _314_v138
			})(), _314_v138)).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(_121_v2, _121_v2))).Cardinality())
			(_138_globalState).F16 = (_269_v108).F29()
		}
		_126_v7 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_126_v7))
	}
	var _317_i13 _dafny.Int
	_ = _317_i13
	_317_i13 = _dafny.Zero
	{
		for _122_v3 {
			{
				if (_317_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_317_i13 = (_317_i13).Plus(_dafny.One)
				if (_129_v10).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_129_v10).Cardinality()))).Uint32()).(bool) {
					var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_124_v5), 0))
					_ = _index37
					(_124_v5).ArraySet1(Companion_Default___.SafeDivisionInt(_126_v7, _126_v7), (_index37).Int())
					var _318_v139 _dafny.Array
					_ = _318_v139
					var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
					_ = _nw51
					_318_v139 = _nw51
					var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_318_v139), 0))
					_ = _index38
					(_318_v139).ArraySet1(_135_v14, (_index38).Int())
					var _319_v140 _dafny.Sequence
					_ = _319_v140
					_319_v140 = _dafny.SeqOf(Companion_Default___.Fm14(_126_v7, _126_v7, _126_v7, (_269_v108).F29(), _138_globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_269_v108).F29()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-53), _dafny.IntOfUint32((_dafny.SeqOf((_269_v108).F29())).Cardinality()))).Uint32(), (_269_v108).F30()), _129_v10))
					var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_124_v5), 0))
					_ = _index39
					var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_318_v139), 0))
					_ = _index40
					var _rhs43 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v19, (_269_v108).F30())).Cardinality()
					_ = _rhs43
					var _rhs44 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_121_v2, _121_v2)
					_ = _rhs44
					var _rhs45 _dafny.Array = _135_v14
					_ = _rhs45
					var _rhs46 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_319_v140, _319_v140)
					_ = _rhs46
					var _rhs47 _dafny.Int = _dafny.IntOfInt64(101)
					_ = _rhs47
					var _lhs24 _dafny.Array = _124_v5
					_ = _lhs24
					var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_124_v5), 0))
					_ = _lhs25
					var _lhs26 _dafny.Array = _318_v139
					_ = _lhs26
					var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_318_v139), 0))
					_ = _lhs27
					(_lhs24).ArraySet1(_rhs43, (_lhs25).Int())
					_121_v2 = _rhs44
					(_lhs26).ArraySet1(_rhs45, (_lhs27).Int())
					_319_v140 = _rhs46
					_126_v7 = _rhs47
					var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_124_v5), 0))
					_ = _index41
					(_124_v5).ArraySet1(_126_v7, (_index41).Int())
					(_138_globalState).F16 = (_269_v108).F30()
					var _320_v141 *C1
					_ = _320_v141
					var _nw52 *C1 = New_C1_()
					_ = _nw52
					_nw52.Ctor__((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_124_v5), 0))).Int()).(_dafny.Int), _318_v139, !((_269_v108).F30()))
					_320_v141 = _nw52
					_320_v141 = _320_v141
					var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_124_v5), 0))
					_ = _index42
					(_124_v5).ArraySet1((_126_v7).Times(((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.ArrayLen((_124_v5), 0))).Int()).(_dafny.Int)).Minus((_141_v19).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_141_v19).Cardinality()))).Uint32()).(_dafny.Int))), (_index42).Int())
				} else {
					var _321_v142 *C5
					_ = _321_v142
					var _nw53 *C5 = New_C5_()
					_ = _nw53
					_nw53.Ctor__((_269_v108).F29(), (_269_v108).F30())
					_321_v142 = _nw53
					var _322_v143 D0
					_ = _322_v143
					_322_v143 = Companion_D0_.Create_DC1_(_119_v0, (_269_v108).F30(), true)
					var _323_v144 bool
					_ = _323_v144
					var _324_v145 bool
					_ = _324_v145
					var _out8 bool
					_ = _out8
					var _out9 bool
					_ = _out9
					_out8, _out9 = (_321_v142).M1(_141_v19, Companion_Default___.SafeDivisionInt(_126_v7, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v10, _126_v7)).Cardinality()), _322_v143, _138_globalState)
					_323_v144 = _out8
					_324_v145 = _out9
					var _325_v146 bool
					_ = _325_v146
					var _326_v147 bool
					_ = _326_v147
					var _out10 bool
					_ = _out10
					var _out11 bool
					_ = _out11
					_out10, _out11 = (_321_v142).M1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(877))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg24 _dafny.Int) interface{} {
							return coer24(arg24)
						}
					}((func(_327_v7 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_328_i14 _dafny.Int) _dafny.Int {
							return _327_v7
						}
					})(_126_v7))), _141_v19), _126_v7, _322_v143, _138_globalState)
					_325_v146 = _out10
					_326_v147 = _out11
					_325_v146 = (_269_v108).F30()
					var _329_v148 _dafny.Array
					_ = _329_v148
					var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
					_ = _nw54
					_329_v148 = _nw54
					var _330_v149 _dafny.Array
					_ = _330_v149
					var _nwElement0_13 _dafny.Array = _329_v148
					_ = _nwElement0_13
					var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(17))
					_ = _nw55
					(_nw55).ArraySet1(_nwElement0_13, 0)
					(_nw55).ArraySet1(_329_v148, 1)
					(_nw55).ArraySet1(_329_v148, 2)
					(_nw55).ArraySet1(_329_v148, 3)
					(_nw55).ArraySet1(_329_v148, 4)
					(_nw55).ArraySet1(_329_v148, 5)
					(_nw55).ArraySet1(_329_v148, 6)
					(_nw55).ArraySet1(_329_v148, 7)
					(_nw55).ArraySet1(_329_v148, 8)
					(_nw55).ArraySet1(_329_v148, 9)
					(_nw55).ArraySet1(_329_v148, 10)
					(_nw55).ArraySet1(_329_v148, 11)
					(_nw55).ArraySet1(_329_v148, 12)
					(_nw55).ArraySet1(_329_v148, 13)
					(_nw55).ArraySet1(_329_v148, 14)
					(_nw55).ArraySet1(_329_v148, 15)
					(_nw55).ArraySet1(_329_v148, 16)
					_330_v149 = _nw55
					var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_330_v149), 0))
					_ = _index43
					(_330_v149).ArraySet1(_329_v148, (_index43).Int())
					var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_330_v149), 0))
					_ = _index44
					(_330_v149).ArraySet1(_329_v148, (_index44).Int())
				}
				if ((_129_v10).Select((Companion_Default___.SafeIndex((_248_v92).Cardinality(), _dafny.IntOfUint32((_129_v10).Cardinality()))).Uint32()).(bool)) && (!(Companion_Default___.Fm26((_269_v108).F29(), (_269_v108).F29(), (_129_v10).Select((Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((_129_v10).Cardinality()))).Uint32()).(bool), (func() _dafny.Int {
					if (_153_v25).Contains(Companion_Default___.Fm1(_119_v0, _138_globalState)) {
						return (_153_v25).Get(Companion_Default___.Fm1(_119_v0, _138_globalState)).(_dafny.Int)
					}
					return _126_v7
				})(), _138_globalState))) {
					var _331_v150 _dafny.Set
					_ = _331_v150
					_331_v150 = _dafny.SetOf(_119_v0)
					var _332_v151 _dafny.Map
					_ = _332_v151
					_332_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_331_v150, _121_v2)
					_332_v151 = (_332_v151).Update((_331_v150).Intersection(_331_v150), _121_v2)
					var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_135_v14), 0))
					_ = _index45
					(_135_v14).ArraySet1((_269_v108).F29(), (_index45).Int())
					var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(810), _dafny.ArrayLen((_135_v14), 0))
					_ = _index46
					(_135_v14).ArraySet1((_126_v7).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_129_v10).Cardinality()), _126_v7))) > 0, (_index46).Int())
					_126_v7 = _dafny.IntOfInt64(929)
					var _333_v152 _dafny.Map
					_ = _333_v152
					_333_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v5, _124_v5)
					_333_v152 = (_333_v152).Update(_124_v5, _124_v5)
					var _334_v153 D0
					_ = _334_v153
					_334_v153 = Companion_D0_.Create_DC1_(_119_v0, (_269_v108).F29(), (_269_v108).F30())
					var _335_v154 bool
					_ = _335_v154
					var _336_v155 bool
					_ = _336_v155
					var _out12 bool
					_ = _out12
					var _out13 bool
					_ = _out13
					_out12, _out13 = (_269_v108).M1(_141_v19, _126_v7, _334_v153, _138_globalState)
					_335_v154 = _out12
					_336_v155 = _out13
				} else {
					var _337_v156 _dafny.Set
					_ = _337_v156
					_337_v156 = _dafny.SetOf(_dafny.IntOfInt64(324))
					_126_v7 = ((_dafny.SetOf(_dafny.IntOfInt64(-408))).Intersection(_337_v156)).Cardinality()
					_247_v91 = (_247_v91).Update((_247_v91).Cardinality(), (_269_v108).F30())
					_126_v7 = Companion_Default___.SafeDivisionInt(_126_v7, (_dafny.IntOfUint32((_121_v2).Cardinality())).Times(_126_v7))
					_126_v7 = (_dafny.Zero).Minus(_126_v7)
					var _338_v157 _dafny.Array
					_ = _338_v157
					var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
					_ = _nw56
					_338_v157 = _nw56
					var _339_v158 _dafny.Map
					_ = _339_v158
					_339_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_119_v0, _138_globalState), _247_v91)
					var _340_v159 *C3
					_ = _340_v159
					var _nw57 *C3 = New_C3_()
					_ = _nw57
					_nw57.Ctor__(_338_v157, _339_v158, (_269_v108).F29())
					_340_v159 = _nw57
				}
				var _rhs48 bool = (_269_v108).F29()
				_ = _rhs48
				var _rhs49 _dafny.Sequence = (func() _dafny.Sequence {
					if _122_v3 {
						return _121_v2
					}
					return _dafny.Companion_Sequence_.Update(Companion_Default___.Fm0(_196_v61, _138_globalState), (Companion_Default___.SafeIndex(_126_v7, _dafny.IntOfUint32((Companion_Default___.Fm0(_196_v61, _138_globalState)).Cardinality()))).Uint32(), _119_v0)
				})()
				_ = _rhs49
				var _rhs50 bool = Companion_Default___.Fm10(_119_v0, (_126_v7).Cmp(_126_v7) == 0, _138_globalState)
				_ = _rhs50
				var _rhs51 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_269_v108).F30() {
						return _dafny.CodePoint('k')
					}
					return _119_v0
				})()
				_ = _rhs51
				var _lhs28 *GlobalState = _138_globalState
				_ = _lhs28
				var _lhs29 *GlobalState = _138_globalState
				_ = _lhs29
				_lhs28.F28 = _rhs48
				_121_v2 = _rhs49
				_lhs29.F28 = _rhs50
				_119_v0 = _rhs51
				var _341_v160 _dafny.Array
				_ = _341_v160
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_4
				var _nw58 _dafny.Array
				_ = _nw58
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw58 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) D14 = (func(_342_v108 *C5, _343_v7 _dafny.Int) func(_dafny.Int) D14 {
						return func(_344_i15 _dafny.Int) D14 {
							return Companion_D14_.Create_DC40_((_342_v108).F29(), (_342_v108).F29(), _343_v7)
						}
					})(_269_v108, _126_v7)
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw58 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw58).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw58).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_341_v160 = _nw58
				var _345_v161 D14
				_ = _345_v161
				_345_v161 = Companion_D14_.Create_DC40_(!((_269_v108).F30()), _122_v3, _126_v7)
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_341_v160), 0))
				_ = _index47
				(_341_v160).ArraySet1(_345_v161, (_index47).Int())
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_341_v160), 0))
				_ = _index48
				(_341_v160).ArraySet1((func() D14 {
					if (_269_v108).F29() {
						return _345_v161
					}
					return _345_v161
				})(), (_index48).Int())
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	_dafny.Print(_119_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v1).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v1).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v1).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_121_v2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_122_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mau"), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_126_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_127_v8).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(159)).UpdateUnsafe(true, _dafny.IntOfInt64(62))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_129_v10, _dafny.SeqOf(false, false, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(159), _dafny.IntOfInt64(159)), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_131_v12).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_131_v12).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_131_v12).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_131_v12).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_131_v12).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_131_v12).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_134_v13).Dtor_cf0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_134_v13).Dtor_cf0()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_134_v13).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_134_v13).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_134_v13).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_134_v13).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_v16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F0()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F0()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F0()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F3()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mau"), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F6().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_138_globalState).F11()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(159), _dafny.IntOfInt64(159)), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_138_globalState.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_138_globalState).F18()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_138_globalState).F18()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_138_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_138_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_138_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_138_globalState).F18()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F20).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_138_globalState).F23()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F24).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_globalState.F27).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_138_globalState.F28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v17).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_140_v18).Equals(_dafny.SetOf(Companion_D0_.Create_DC3_(_dafny.CodePoint('v')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_141_v19, _dafny.SeqOf(_dafny.IntOfInt64(159), _dafny.IntOfInt64(159), _dafny.IntOfInt64(159), _dafny.IntOfInt64(159))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v24).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(159), _dafny.IntOfInt64(159))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v25).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(159), _dafny.IntOfInt64(159))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_187_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_190_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_196_v61).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_197_v62, _dafny.SeqOf(_dafny.MultiSetOf(true, true, true, true), _dafny.MultiSetOf(true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v63).Dtor_cf18()).Equals(_dafny.MultiSetOf(true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_247_v91).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(3), true).UpdateUnsafe(_dafny.One, true).UpdateUnsafe(_dafny.IntOfInt64(2), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_248_v92).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(3), true).UpdateUnsafe(_dafny.IntOfInt64(796), true), _dafny.IntOfInt64(796))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_249_v93).Dtor_cf61()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(3), true).UpdateUnsafe(_dafny.IntOfInt64(796), true), _dafny.IntOfInt64(796))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v94).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_251_v95).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v96).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_253_v97, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v98).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.One), _dafny.MultiSetOf(_dafny.One), _dafny.MultiSetOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v99).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfInt64(159), _dafny.IntOfInt64(159), _dafny.IntOfInt64(159), _dafny.IntOfInt64(159)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v108).F29())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v108).F30())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v109).Dtor_cf24()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v109).Dtor_cf24()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_270_v109).Dtor_cf24()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_317_i13)
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
	Cf1 _dafny.CodePoint
	Cf2 bool
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.CodePoint, Cf2 bool, Cf3 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 bool) D0 {
	return D0{D0_DC2{Cf4}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf5 _dafny.CodePoint
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf5 _dafny.CodePoint) D0 {
	return D0{D0_DC3{Cf5}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
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

type D0_DC4 struct {
	Cf6 D0
}

func (D0_DC4) isD0() {}

func (CompanionStruct_D0_) Create_DC4_(Cf6 D0) D0 {
	return D0{D0_DC4{Cf6}}
}

func (_this D0) Is_DC4() bool {
	_, ok := _this.Get_().(D0_DC4)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.CodePoint('D'), false, false)
}

func (_this D0) Dtor_cf1() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.CodePoint {
	return _this.Get_().(D0_DC3).Cf5
}

func (_this D0) Dtor_cf0() _dafny.Array {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf6() D0 {
	return _this.Get_().(D0_DC4).Cf6
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
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC4:
		{
			return "D0.DC4" + "(" + _dafny.String(data.Cf6) + ")"
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
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4 == data2.Cf4
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf5 == data2.Cf5
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC4:
		{
			data2, ok := other.Get_().(D0_DC4)
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
	Cf7 _dafny.Array
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf7 _dafny.Array) D1 {
	return D1{D1_DC5{Cf7}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D1) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf7 == data2.Cf7
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
	Cf9  _dafny.CodePoint
	Cf10 _dafny.Set
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf9 _dafny.CodePoint, Cf10 _dafny.Set) D2 {
	return D2{D2_DC7{Cf9, Cf10}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf8 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf8 _dafny.Int) D2 {
	return D2{D2_DC6{Cf8}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC8 struct {
	Cf11 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf11 D2) D2 {
	return D2{D2_DC8{Cf11}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.CodePoint('D'), _dafny.EmptySet)
}

func (_this D2) Dtor_cf9() _dafny.CodePoint {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) Dtor_cf10() _dafny.Set {
	return _this.Get_().(D2_DC7).Cf10
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf8
}

func (_this D2) Dtor_cf11() D2 {
	return _this.Get_().(D2_DC8).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf9 == data2.Cf9 && data1.Cf10.Equals(data2.Cf10)
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf13 _dafny.Sequence
	Cf14 _dafny.Int
	Cf15 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf13 _dafny.Sequence, Cf14 _dafny.Int, Cf15 bool) D3 {
	return D3{D3_DC10{Cf13, Cf14, Cf15}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf12 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf12 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf12}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC11 struct {
	Cf16 D3
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf16 D3) D3 {
	return D3{D3_DC11{Cf16}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.EmptySeq, _dafny.Zero, false)
}

func (_this D3) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf14
}

func (_this D3) Dtor_cf15() bool {
	return _this.Get_().(D3_DC10).Cf15
}

func (_this D3) Dtor_cf12() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf12
}

func (_this D3) Dtor_cf16() D3 {
	return _this.Get_().(D3_DC11).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + data.Cf13.VerbatimString(true) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf13.Equals(data2.Cf13) && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf12.Equals(data2.Cf12)
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
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
	Cf17 _dafny.Set
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf17 _dafny.Set) D4 {
	return D4{D4_DC12{Cf17}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D4) Dtor_cf17() _dafny.Set {
	return _this.Get_().(D4_DC12).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf17) + ")"
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

type D5_DC14 struct {
	Cf19 _dafny.Set
	Cf20 _dafny.Int
	Cf21 bool
	Cf22 _dafny.Array
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf19 _dafny.Set, Cf20 _dafny.Int, Cf21 bool, Cf22 _dafny.Array) D5 {
	return D5{D5_DC14{Cf19, Cf20, Cf21, Cf22}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf18 _dafny.MultiSet
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf18 _dafny.MultiSet) D5 {
	return D5{D5_DC13{Cf18}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC15 struct {
	Cf23 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf23 D5) D5 {
	return D5{D5_DC15{Cf23}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(_dafny.EmptySet, _dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D5) Dtor_cf19() _dafny.Set {
	return _this.Get_().(D5_DC14).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC14).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf22
}

func (_this D5) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D5_DC13).Cf18
}

func (_this D5) Dtor_cf23() D5 {
	return _this.Get_().(D5_DC15).Cf23
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf19.Equals(data2.Cf19) && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf23.Equals(data2.Cf23)
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
	Cf25 bool
	Cf26 _dafny.Array
	Cf27 _dafny.Int
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf25 bool, Cf26 _dafny.Array, Cf27 _dafny.Int) D6 {
	return D6{D6_DC17{Cf25, Cf26, Cf27}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf24 _dafny.Array
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf24 _dafny.Array) D6 {
	return D6{D6_DC16{Cf24}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC18 struct {
	Cf28 D6
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf28 D6) D6 {
	return D6{D6_DC18{Cf28}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D6) Dtor_cf25() bool {
	return _this.Get_().(D6_DC17).Cf25
}

func (_this D6) Dtor_cf26() _dafny.Array {
	return _this.Get_().(D6_DC17).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf27
}

func (_this D6) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D6_DC16).Cf24
}

func (_this D6) Dtor_cf28() D6 {
	return _this.Get_().(D6_DC18).Cf28
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf28.Equals(data2.Cf28)
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
	Cf29 _dafny.Map
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf29 _dafny.Map) D7 {
	return D7{D7_DC19{Cf29}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D7) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D7_DC19).Cf29
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf29) + ")"
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

type D8_DC21 struct {
	Cf31 bool
	Cf32 bool
	Cf33 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf31 bool, Cf32 bool, Cf33 _dafny.Int) D8 {
	return D8{D8_DC21{Cf31, Cf32, Cf33}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC22 struct {
	Cf34 bool
	Cf35 bool
	Cf36 _dafny.MultiSet
	Cf37 _dafny.Array
	Cf38 bool
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf34 bool, Cf35 bool, Cf36 _dafny.MultiSet, Cf37 _dafny.Array, Cf38 bool) D8 {
	return D8{D8_DC22{Cf34, Cf35, Cf36, Cf37, Cf38}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC20 struct {
	Cf30 _dafny.Map
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf30 _dafny.Map) D8 {
	return D8{D8_DC20{Cf30}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(false, false, _dafny.Zero)
}

func (_this D8) Dtor_cf31() bool {
	return _this.Get_().(D8_DC21).Cf31
}

func (_this D8) Dtor_cf32() bool {
	return _this.Get_().(D8_DC21).Cf32
}

func (_this D8) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf33
}

func (_this D8) Dtor_cf34() bool {
	return _this.Get_().(D8_DC22).Cf34
}

func (_this D8) Dtor_cf35() bool {
	return _this.Get_().(D8_DC22).Cf35
}

func (_this D8) Dtor_cf36() _dafny.MultiSet {
	return _this.Get_().(D8_DC22).Cf36
}

func (_this D8) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D8_DC22).Cf37
}

func (_this D8) Dtor_cf38() bool {
	return _this.Get_().(D8_DC22).Cf38
}

func (_this D8) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D8_DC20).Cf30
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35 && data1.Cf36.Equals(data2.Cf36) && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf30.Equals(data2.Cf30)
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
	Cf40 _dafny.MultiSet
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf40 _dafny.MultiSet) D9 {
	return D9{D9_DC24{Cf40}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC25 struct {
	Cf41 _dafny.Int
	Cf42 bool
	Cf43 bool
	Cf44 bool
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf41 _dafny.Int, Cf42 bool, Cf43 bool, Cf44 bool) D9 {
	return D9{D9_DC25{Cf41, Cf42, Cf43, Cf44}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC26 struct {
	Cf45 bool
	Cf46 _dafny.Sequence
	Cf47 _dafny.Int
	Cf48 _dafny.Map
	Cf49 bool
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf45 bool, Cf46 _dafny.Sequence, Cf47 _dafny.Int, Cf48 _dafny.Map, Cf49 bool) D9 {
	return D9{D9_DC26{Cf45, Cf46, Cf47, Cf48, Cf49}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC23 struct {
	Cf39 _dafny.Sequence
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf39 _dafny.Sequence) D9 {
	return D9{D9_DC23{Cf39}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC24_(_dafny.EmptyMultiSet)
}

func (_this D9) Dtor_cf40() _dafny.MultiSet {
	return _this.Get_().(D9_DC24).Cf40
}

func (_this D9) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf41
}

func (_this D9) Dtor_cf42() bool {
	return _this.Get_().(D9_DC25).Cf42
}

func (_this D9) Dtor_cf43() bool {
	return _this.Get_().(D9_DC25).Cf43
}

func (_this D9) Dtor_cf44() bool {
	return _this.Get_().(D9_DC25).Cf44
}

func (_this D9) Dtor_cf45() bool {
	return _this.Get_().(D9_DC26).Cf45
}

func (_this D9) Dtor_cf46() _dafny.Sequence {
	return _this.Get_().(D9_DC26).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D9_DC26).Cf47
}

func (_this D9) Dtor_cf48() _dafny.Map {
	return _this.Get_().(D9_DC26).Cf48
}

func (_this D9) Dtor_cf49() bool {
	return _this.Get_().(D9_DC26).Cf49
}

func (_this D9) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D9_DC23).Cf39
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf45) + ", " + data.Cf46.VerbatimString(true) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf40.Equals(data2.Cf40)
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf41.Cmp(data2.Cf41) == 0 && data1.Cf42 == data2.Cf42 && data1.Cf43 == data2.Cf43 && data1.Cf44 == data2.Cf44
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf45 == data2.Cf45 && data1.Cf46.Equals(data2.Cf46) && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Equals(data2.Cf48) && data1.Cf49 == data2.Cf49
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	Cf51 _dafny.Int
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf51 _dafny.Int) D10 {
	return D10{D10_DC28{Cf51}}
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

type D10_DC29 struct {
	Cf52 D10
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf52 D10) D10 {
	return D10{D10_DC29{Cf52}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC28_(_dafny.Zero)
}

func (_this D10) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D10_DC28).Cf51
}

func (_this D10) Dtor_cf50() _dafny.Map {
	return _this.Get_().(D10_DC27).Cf50
}

func (_this D10) Dtor_cf52() D10 {
	return _this.Get_().(D10_DC29).Cf52
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf51) + ")"
		}
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf52) + ")"
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
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf51.Cmp(data2.Cf51) == 0
		}
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf50.Equals(data2.Cf50)
		}
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && data1.Cf52.Equals(data2.Cf52)
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
	Cf53 *C1
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf53 *C1) D11 {
	return D11{D11_DC30{Cf53}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

func (CompanionStruct_D11_) Default() *C1 {
	return (*C1)(nil)
}

func (_this D11) Dtor_cf53() *C1 {
	return _this.Get_().(D11_DC30).Cf53
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf53) + ")"
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
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf53 == data2.Cf53
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
	Cf55 bool
	Cf56 bool
	Cf57 T0
	Cf58 _dafny.Set
	Cf59 _dafny.Int
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf55 bool, Cf56 bool, Cf57 T0, Cf58 _dafny.Set, Cf59 _dafny.Int) D12 {
	return D12{D12_DC32{Cf55, Cf56, Cf57, Cf58, Cf59}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC33 struct {
	Cf60 _dafny.Sequence
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf60 _dafny.Sequence) D12 {
	return D12{D12_DC33{Cf60}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC31 struct {
	Cf54 _dafny.Sequence
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf54 _dafny.Sequence) D12 {
	return D12{D12_DC31{Cf54}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC32_(false, false, (T0)(nil), _dafny.EmptySet, _dafny.Zero)
}

func (_this D12) Dtor_cf55() bool {
	return _this.Get_().(D12_DC32).Cf55
}

func (_this D12) Dtor_cf56() bool {
	return _this.Get_().(D12_DC32).Cf56
}

func (_this D12) Dtor_cf57() T0 {
	return _this.Get_().(D12_DC32).Cf57
}

func (_this D12) Dtor_cf58() _dafny.Set {
	return _this.Get_().(D12_DC32).Cf58
}

func (_this D12) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf59
}

func (_this D12) Dtor_cf60() _dafny.Sequence {
	return _this.Get_().(D12_DC33).Cf60
}

func (_this D12) Dtor_cf54() _dafny.Sequence {
	return _this.Get_().(D12_DC31).Cf54
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D12_DC33:
		{
			return "D12.DC33" + "(" + data.Cf60.VerbatimString(true) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf54) + ")"
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
			return ok && data1.Cf55 == data2.Cf55 && data1.Cf56 == data2.Cf56 && _dafny.AreEqual(data1.Cf57, data2.Cf57) && data1.Cf58.Equals(data2.Cf58) && data1.Cf59.Cmp(data2.Cf59) == 0
		}
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf60.Equals(data2.Cf60)
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf54.Equals(data2.Cf54)
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
	Cf62 _dafny.Int
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf62 _dafny.Int) D13 {
	return D13{D13_DC35{Cf62}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

type D13_DC36 struct {
	Cf63 bool
	Cf64 bool
}

func (D13_DC36) isD13() {}

func (CompanionStruct_D13_) Create_DC36_(Cf63 bool, Cf64 bool) D13 {
	return D13{D13_DC36{Cf63, Cf64}}
}

func (_this D13) Is_DC36() bool {
	_, ok := _this.Get_().(D13_DC36)
	return ok
}

type D13_DC37 struct {
	Cf65 bool
	Cf66 _dafny.Int
}

func (D13_DC37) isD13() {}

func (CompanionStruct_D13_) Create_DC37_(Cf65 bool, Cf66 _dafny.Int) D13 {
	return D13{D13_DC37{Cf65, Cf66}}
}

func (_this D13) Is_DC37() bool {
	_, ok := _this.Get_().(D13_DC37)
	return ok
}

type D13_DC34 struct {
	Cf61 _dafny.Map
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf61 _dafny.Map) D13 {
	return D13{D13_DC34{Cf61}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

type D13_DC38 struct {
	Cf67 D13
}

func (D13_DC38) isD13() {}

func (CompanionStruct_D13_) Create_DC38_(Cf67 D13) D13 {
	return D13{D13_DC38{Cf67}}
}

func (_this D13) Is_DC38() bool {
	_, ok := _this.Get_().(D13_DC38)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC35_(_dafny.Zero)
}

func (_this D13) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D13_DC35).Cf62
}

func (_this D13) Dtor_cf63() bool {
	return _this.Get_().(D13_DC36).Cf63
}

func (_this D13) Dtor_cf64() bool {
	return _this.Get_().(D13_DC36).Cf64
}

func (_this D13) Dtor_cf65() bool {
	return _this.Get_().(D13_DC37).Cf65
}

func (_this D13) Dtor_cf66() _dafny.Int {
	return _this.Get_().(D13_DC37).Cf66
}

func (_this D13) Dtor_cf61() _dafny.Map {
	return _this.Get_().(D13_DC34).Cf61
}

func (_this D13) Dtor_cf67() D13 {
	return _this.Get_().(D13_DC38).Cf67
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf62) + ")"
		}
	case D13_DC36:
		{
			return "D13.DC36" + "(" + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D13_DC37:
		{
			return "D13.DC37" + "(" + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ")"
		}
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf61) + ")"
		}
	case D13_DC38:
		{
			return "D13.DC38" + "(" + _dafny.String(data.Cf67) + ")"
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
			return ok && data1.Cf62.Cmp(data2.Cf62) == 0
		}
	case D13_DC36:
		{
			data2, ok := other.Get_().(D13_DC36)
			return ok && data1.Cf63 == data2.Cf63 && data1.Cf64 == data2.Cf64
		}
	case D13_DC37:
		{
			data2, ok := other.Get_().(D13_DC37)
			return ok && data1.Cf65 == data2.Cf65 && data1.Cf66.Cmp(data2.Cf66) == 0
		}
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf61.Equals(data2.Cf61)
		}
	case D13_DC38:
		{
			data2, ok := other.Get_().(D13_DC38)
			return ok && data1.Cf67.Equals(data2.Cf67)
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

type D14_DC40 struct {
	Cf69 bool
	Cf70 bool
	Cf71 _dafny.Int
}

func (D14_DC40) isD14() {}

func (CompanionStruct_D14_) Create_DC40_(Cf69 bool, Cf70 bool, Cf71 _dafny.Int) D14 {
	return D14{D14_DC40{Cf69, Cf70, Cf71}}
}

func (_this D14) Is_DC40() bool {
	_, ok := _this.Get_().(D14_DC40)
	return ok
}

type D14_DC41 struct {
	Cf72 _dafny.Int
	Cf73 _dafny.Int
}

func (D14_DC41) isD14() {}

func (CompanionStruct_D14_) Create_DC41_(Cf72 _dafny.Int, Cf73 _dafny.Int) D14 {
	return D14{D14_DC41{Cf72, Cf73}}
}

func (_this D14) Is_DC41() bool {
	_, ok := _this.Get_().(D14_DC41)
	return ok
}

type D14_DC39 struct {
	Cf68 _dafny.Array
}

func (D14_DC39) isD14() {}

func (CompanionStruct_D14_) Create_DC39_(Cf68 _dafny.Array) D14 {
	return D14{D14_DC39{Cf68}}
}

func (_this D14) Is_DC39() bool {
	_, ok := _this.Get_().(D14_DC39)
	return ok
}

type D14_DC42 struct {
	Cf74 D14
}

func (D14_DC42) isD14() {}

func (CompanionStruct_D14_) Create_DC42_(Cf74 D14) D14 {
	return D14{D14_DC42{Cf74}}
}

func (_this D14) Is_DC42() bool {
	_, ok := _this.Get_().(D14_DC42)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC40_(false, false, _dafny.Zero)
}

func (_this D14) Dtor_cf69() bool {
	return _this.Get_().(D14_DC40).Cf69
}

func (_this D14) Dtor_cf70() bool {
	return _this.Get_().(D14_DC40).Cf70
}

func (_this D14) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D14_DC40).Cf71
}

func (_this D14) Dtor_cf72() _dafny.Int {
	return _this.Get_().(D14_DC41).Cf72
}

func (_this D14) Dtor_cf73() _dafny.Int {
	return _this.Get_().(D14_DC41).Cf73
}

func (_this D14) Dtor_cf68() _dafny.Array {
	return _this.Get_().(D14_DC39).Cf68
}

func (_this D14) Dtor_cf74() D14 {
	return _this.Get_().(D14_DC42).Cf74
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC40:
		{
			return "D14.DC40" + "(" + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D14_DC41:
		{
			return "D14.DC41" + "(" + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ")"
		}
	case D14_DC39:
		{
			return "D14.DC39" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D14_DC42:
		{
			return "D14.DC42" + "(" + _dafny.String(data.Cf74) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC40:
		{
			data2, ok := other.Get_().(D14_DC40)
			return ok && data1.Cf69 == data2.Cf69 && data1.Cf70 == data2.Cf70 && data1.Cf71.Cmp(data2.Cf71) == 0
		}
	case D14_DC41:
		{
			data2, ok := other.Get_().(D14_DC41)
			return ok && data1.Cf72.Cmp(data2.Cf72) == 0 && data1.Cf73.Cmp(data2.Cf73) == 0
		}
	case D14_DC39:
		{
			data2, ok := other.Get_().(D14_DC39)
			return ok && data1.Cf68 == data2.Cf68
		}
	case D14_DC42:
		{
			data2, ok := other.Get_().(D14_DC42)
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

type D15_DC43 struct {
	Cf75 *C0
}

func (D15_DC43) isD15() {}

func (CompanionStruct_D15_) Create_DC43_(Cf75 *C0) D15 {
	return D15{D15_DC43{Cf75}}
}

func (_this D15) Is_DC43() bool {
	_, ok := _this.Get_().(D15_DC43)
	return ok
}

func (CompanionStruct_D15_) Default() *C0 {
	return (*C0)(nil)
}

func (_this D15) Dtor_cf75() *C0 {
	return _this.Get_().(D15_DC43).Cf75
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC43:
		{
			return "D15.DC43" + "(" + _dafny.String(data.Cf75) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC43:
		{
			data2, ok := other.Get_().(D15_DC43)
			return ok && data1.Cf75 == data2.Cf75
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

type D16_DC45 struct {
	Cf77 _dafny.CodePoint
	Cf78 bool
	Cf79 _dafny.Int
}

func (D16_DC45) isD16() {}

func (CompanionStruct_D16_) Create_DC45_(Cf77 _dafny.CodePoint, Cf78 bool, Cf79 _dafny.Int) D16 {
	return D16{D16_DC45{Cf77, Cf78, Cf79}}
}

func (_this D16) Is_DC45() bool {
	_, ok := _this.Get_().(D16_DC45)
	return ok
}

type D16_DC44 struct {
	Cf76 _dafny.Sequence
}

func (D16_DC44) isD16() {}

func (CompanionStruct_D16_) Create_DC44_(Cf76 _dafny.Sequence) D16 {
	return D16{D16_DC44{Cf76}}
}

func (_this D16) Is_DC44() bool {
	_, ok := _this.Get_().(D16_DC44)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC45_(_dafny.CodePoint('D'), false, _dafny.Zero)
}

func (_this D16) Dtor_cf77() _dafny.CodePoint {
	return _this.Get_().(D16_DC45).Cf77
}

func (_this D16) Dtor_cf78() bool {
	return _this.Get_().(D16_DC45).Cf78
}

func (_this D16) Dtor_cf79() _dafny.Int {
	return _this.Get_().(D16_DC45).Cf79
}

func (_this D16) Dtor_cf76() _dafny.Sequence {
	return _this.Get_().(D16_DC44).Cf76
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC45:
		{
			return "D16.DC45" + "(" + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D16_DC44:
		{
			return "D16.DC44" + "(" + _dafny.String(data.Cf76) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC45:
		{
			data2, ok := other.Get_().(D16_DC45)
			return ok && data1.Cf77 == data2.Cf77 && data1.Cf78 == data2.Cf78 && data1.Cf79.Cmp(data2.Cf79) == 0
		}
	case D16_DC44:
		{
			data2, ok := other.Get_().(D16_DC44)
			return ok && data1.Cf76.Equals(data2.Cf76)
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

type D17_DC47 struct {
	Cf81 bool
	Cf82 _dafny.Int
}

func (D17_DC47) isD17() {}

func (CompanionStruct_D17_) Create_DC47_(Cf81 bool, Cf82 _dafny.Int) D17 {
	return D17{D17_DC47{Cf81, Cf82}}
}

func (_this D17) Is_DC47() bool {
	_, ok := _this.Get_().(D17_DC47)
	return ok
}

type D17_DC46 struct {
	Cf80 _dafny.Map
}

func (D17_DC46) isD17() {}

func (CompanionStruct_D17_) Create_DC46_(Cf80 _dafny.Map) D17 {
	return D17{D17_DC46{Cf80}}
}

func (_this D17) Is_DC46() bool {
	_, ok := _this.Get_().(D17_DC46)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC47_(false, _dafny.Zero)
}

func (_this D17) Dtor_cf81() bool {
	return _this.Get_().(D17_DC47).Cf81
}

func (_this D17) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D17_DC47).Cf82
}

func (_this D17) Dtor_cf80() _dafny.Map {
	return _this.Get_().(D17_DC46).Cf80
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC47:
		{
			return "D17.DC47" + "(" + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ")"
		}
	case D17_DC46:
		{
			return "D17.DC46" + "(" + _dafny.String(data.Cf80) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC47:
		{
			data2, ok := other.Get_().(D17_DC47)
			return ok && data1.Cf81 == data2.Cf81 && data1.Cf82.Cmp(data2.Cf82) == 0
		}
	case D17_DC46:
		{
			data2, ok := other.Get_().(D17_DC46)
			return ok && data1.Cf80.Equals(data2.Cf80)
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

type D18_DC49 struct {
}

func (D18_DC49) isD18() {}

func (CompanionStruct_D18_) Create_DC49_() D18 {
	return D18{D18_DC49{}}
}

func (_this D18) Is_DC49() bool {
	_, ok := _this.Get_().(D18_DC49)
	return ok
}

type D18_DC50 struct {
	Cf84 _dafny.Int
	Cf85 _dafny.Sequence
}

func (D18_DC50) isD18() {}

func (CompanionStruct_D18_) Create_DC50_(Cf84 _dafny.Int, Cf85 _dafny.Sequence) D18 {
	return D18{D18_DC50{Cf84, Cf85}}
}

func (_this D18) Is_DC50() bool {
	_, ok := _this.Get_().(D18_DC50)
	return ok
}

type D18_DC51 struct {
	Cf86 _dafny.Int
	Cf87 bool
	Cf88 bool
	Cf89 _dafny.Set
	Cf90 bool
}

func (D18_DC51) isD18() {}

func (CompanionStruct_D18_) Create_DC51_(Cf86 _dafny.Int, Cf87 bool, Cf88 bool, Cf89 _dafny.Set, Cf90 bool) D18 {
	return D18{D18_DC51{Cf86, Cf87, Cf88, Cf89, Cf90}}
}

func (_this D18) Is_DC51() bool {
	_, ok := _this.Get_().(D18_DC51)
	return ok
}

type D18_DC48 struct {
	Cf83 _dafny.Map
}

func (D18_DC48) isD18() {}

func (CompanionStruct_D18_) Create_DC48_(Cf83 _dafny.Map) D18 {
	return D18{D18_DC48{Cf83}}
}

func (_this D18) Is_DC48() bool {
	_, ok := _this.Get_().(D18_DC48)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC49_()
}

func (_this D18) Dtor_cf84() _dafny.Int {
	return _this.Get_().(D18_DC50).Cf84
}

func (_this D18) Dtor_cf85() _dafny.Sequence {
	return _this.Get_().(D18_DC50).Cf85
}

func (_this D18) Dtor_cf86() _dafny.Int {
	return _this.Get_().(D18_DC51).Cf86
}

func (_this D18) Dtor_cf87() bool {
	return _this.Get_().(D18_DC51).Cf87
}

func (_this D18) Dtor_cf88() bool {
	return _this.Get_().(D18_DC51).Cf88
}

func (_this D18) Dtor_cf89() _dafny.Set {
	return _this.Get_().(D18_DC51).Cf89
}

func (_this D18) Dtor_cf90() bool {
	return _this.Get_().(D18_DC51).Cf90
}

func (_this D18) Dtor_cf83() _dafny.Map {
	return _this.Get_().(D18_DC48).Cf83
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC49:
		{
			return "D18.DC49"
		}
	case D18_DC50:
		{
			return "D18.DC50" + "(" + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ")"
		}
	case D18_DC51:
		{
			return "D18.DC51" + "(" + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ", " + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ")"
		}
	case D18_DC48:
		{
			return "D18.DC48" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC49:
		{
			_, ok := other.Get_().(D18_DC49)
			return ok
		}
	case D18_DC50:
		{
			data2, ok := other.Get_().(D18_DC50)
			return ok && data1.Cf84.Cmp(data2.Cf84) == 0 && data1.Cf85.Equals(data2.Cf85)
		}
	case D18_DC51:
		{
			data2, ok := other.Get_().(D18_DC51)
			return ok && data1.Cf86.Cmp(data2.Cf86) == 0 && data1.Cf87 == data2.Cf87 && data1.Cf88 == data2.Cf88 && data1.Cf89.Equals(data2.Cf89) && data1.Cf90 == data2.Cf90
		}
	case D18_DC48:
		{
			data2, ok := other.Get_().(D18_DC48)
			return ok && data1.Cf83.Equals(data2.Cf83)
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

type D19_DC53 struct {
	Cf92 bool
}

func (D19_DC53) isD19() {}

func (CompanionStruct_D19_) Create_DC53_(Cf92 bool) D19 {
	return D19{D19_DC53{Cf92}}
}

func (_this D19) Is_DC53() bool {
	_, ok := _this.Get_().(D19_DC53)
	return ok
}

type D19_DC54 struct {
}

func (D19_DC54) isD19() {}

func (CompanionStruct_D19_) Create_DC54_() D19 {
	return D19{D19_DC54{}}
}

func (_this D19) Is_DC54() bool {
	_, ok := _this.Get_().(D19_DC54)
	return ok
}

type D19_DC52 struct {
	Cf91 _dafny.MultiSet
}

func (D19_DC52) isD19() {}

func (CompanionStruct_D19_) Create_DC52_(Cf91 _dafny.MultiSet) D19 {
	return D19{D19_DC52{Cf91}}
}

func (_this D19) Is_DC52() bool {
	_, ok := _this.Get_().(D19_DC52)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC53_(false)
}

func (_this D19) Dtor_cf92() bool {
	return _this.Get_().(D19_DC53).Cf92
}

func (_this D19) Dtor_cf91() _dafny.MultiSet {
	return _this.Get_().(D19_DC52).Cf91
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC53:
		{
			return "D19.DC53" + "(" + _dafny.String(data.Cf92) + ")"
		}
	case D19_DC54:
		{
			return "D19.DC54"
		}
	case D19_DC52:
		{
			return "D19.DC52" + "(" + _dafny.String(data.Cf91) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC53:
		{
			data2, ok := other.Get_().(D19_DC53)
			return ok && data1.Cf92 == data2.Cf92
		}
	case D19_DC54:
		{
			_, ok := other.Get_().(D19_DC54)
			return ok
		}
	case D19_DC52:
		{
			data2, ok := other.Get_().(D19_DC52)
			return ok && data1.Cf91.Equals(data2.Cf91)
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

// Definition of trait T0
type T0 interface {
	String() string
	F31() bool
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
	F31() bool
	Fm6(p0 bool, globalState *GlobalState) _dafny.Sequence
	M2(p0 bool, p1 _dafny.MultiSet, p2 _dafny.MultiSet, globalState *GlobalState) (_dafny.Int, _dafny.Array, _dafny.Sequence, _dafny.Int)
	F32() _dafny.Int
	F33() _dafny.Array
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
	F12  _dafny.Map
	F16  bool
	F20  _dafny.Array
	F24  _dafny.Array
	F27  _dafny.Array
	F28  bool
	_f0  _dafny.Array
	_f1  bool
	_f2  bool
	_f3  _dafny.Map
	_f4  bool
	_f5  bool
	_f6  _dafny.Sequence
	_f7  _dafny.Int
	_f8  bool
	_f9  bool
	_f10 bool
	_f11 _dafny.Sequence
	_f13 _dafny.Int
	_f14 bool
	_f15 _dafny.Int
	_f17 bool
	_f18 _dafny.Array
	_f19 bool
	_f21 _dafny.Array
	_f22 bool
	_f23 _dafny.Map
	_f25 _dafny.Int
	_f26 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F12 = _dafny.EmptyMap
	_this.F16 = false
	_this.F20 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F24 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F28 = false
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = false
	_this._f2 = false
	_this._f3 = _dafny.EmptyMap
	_this._f4 = false
	_this._f5 = false
	_this._f6 = _dafny.EmptySeq
	_this._f7 = _dafny.Zero
	_this._f8 = false
	_this._f9 = false
	_this._f10 = false
	_this._f11 = _dafny.EmptySeq
	_this._f13 = _dafny.Zero
	_this._f14 = false
	_this._f15 = _dafny.Zero
	_this._f17 = false
	_this._f18 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f19 = false
	_this._f21 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f22 = false
	_this._f23 = _dafny.EmptyMap
	_this._f25 = _dafny.Zero
	_this._f26 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 bool, f2 bool, f3 _dafny.Map, f4 bool, f5 bool, f6 _dafny.Sequence, f7 _dafny.Int, f8 bool, f9 bool, f10 bool, f11 _dafny.Sequence, f12 _dafny.Map, f13 _dafny.Int, f14 bool, f15 _dafny.Int, f16 bool, f17 bool, f18 _dafny.Array, f19 bool, f20 _dafny.Array, f21 _dafny.Array, f22 bool, f23 _dafny.Map, f24 _dafny.Array, f25 _dafny.Int, f26 _dafny.Int, f27 _dafny.Array, f28 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this).F12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this).F16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this).F20 = f20
		(_this)._f21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this).F24 = f24
		(_this)._f25 = f25
		(_this)._f26 = f26
		(_this).F27 = f27
		(_this).F28 = f28
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
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
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
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Sequence {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Sequence {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Array {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() bool {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F21() _dafny.Array {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F22() bool {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() _dafny.Map {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F25() _dafny.Int {
	{
		return _this._f25
	}
}
func (_this *GlobalState) F26() _dafny.Int {
	{
		return _this._f26
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f31 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f31 = false
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

func (_this *C0) F31() bool {
	return _this._f31
}
func (_this *C0) Ctor__(f31 bool) {
	{
		(_this)._f31 = f31
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f31 bool
	_f32 _dafny.Int
	_f33 _dafny.Array
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f31 = false
	_this._f32 = _dafny.Zero
	_this._f33 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C1) F31() bool {
	return _this._f31
}
func (_this *C1) F32() _dafny.Int {
	return _this._f32
}
func (_this *C1) F33() _dafny.Array {
	return _this._f33
}
func (_this *C1) Ctor__(f32 _dafny.Int, f33 _dafny.Array, f31 bool) {
	{
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f31 = f31
	}
}
func (_this *C1) Fm6(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf((_this).F31(), false, (_this).F31(), true, !(((_this).F32()).Cmp((_this).F32()) == 0))
	}
}
func (_this *C1) Fm16(p0 _dafny.Sequence, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(531))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_346_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_347_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})), _dafny.UnicodeSeqOfUtf8Bytes("lriggcwfc")))
	}
}
func (_this *C1) Fm17(p0 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F32()))).IsSubsetOf((func() _dafny.MultiSet {
			if (_this).F31() {
				return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ndqhhq")).Cardinality()), (_this).F32(), _dafny.IntOfInt64(679))
			}
			return _dafny.MultiSetOf(_dafny.IntOfInt64(-370), (_this).F32(), (_dafny.MultiSetOf((_this).F31(), !((_this).F31()), true, (_this).F31(), (_this).F31())).Cardinality())
		})())
	}
}
func (_this *C1) M2(p0 bool, p1 _dafny.MultiSet, p2 _dafny.MultiSet, globalState *GlobalState) (_dafny.Int, _dafny.Array, _dafny.Sequence, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _348_v0 _dafny.Sequence
		_ = _348_v0
		_348_v0 = _dafny.SeqOf(!(p0), (_this).F31(), !(false), (_this).F31())
		var _349_v1 _dafny.Sequence
		_ = _349_v1
		_349_v1 = _dafny.SeqOf((_this).F32(), (_this).F32(), _dafny.IntOfUint32((_348_v0).Cardinality()), Companion_Default___.Fm18((_this).F31(), p0, globalState), (_this).F32())
		var _350_v2 _dafny.Map
		_ = _350_v2
		_350_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(269), _349_v1)
		_350_v2 = (_350_v2).Update((_this).F32(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F32(), (_this).F32()), _349_v1))
		r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_348_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_349_v1).Cardinality()), _dafny.IntOfUint32((_348_v0).Cardinality()))).Uint32()).(bool)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F31(), (_this).F31(), (_this).F31()), _348_v0)), (Companion_Default___.SafeIndex(((_this).F32()).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-955))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_351_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_348_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_349_v1).Cardinality()), _dafny.IntOfUint32((_348_v0).Cardinality()))).Uint32()).(bool)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F31(), (_this).F31(), (_this).F31()), _348_v0))).Cardinality()))).Uint32(), p0)).Cardinality())
		var _352_v3 _dafny.Array
		_ = _352_v3
		var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
		_ = _nw59
		_352_v3 = _nw59
		_352_v3 = _352_v3
		var _353_i1 _dafny.Int
		_ = _353_i1
		_353_i1 = _dafny.Zero
		{
			for false {
				{
					if (_353_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_353_i1 = (_353_i1).Plus(_dafny.One)
					var _354_v5 _dafny.Map
					_ = _354_v5
					_354_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
					var _355_v6 _dafny.Sequence
					_ = _355_v6
					_355_v6 = _dafny.SeqOf(_354_v5)
					var _356_v7 *C0
					_ = _356_v7
					var _nw60 *C0 = New_C0_()
					_ = _nw60
					_nw60.Ctor__(p0)
					_356_v7 = _nw60
					var _357_v8 _dafny.MultiSet
					_ = _357_v8
					_357_v8 = _dafny.MultiSetOf(_356_v7, _356_v7)
					var _358_v9 _dafny.Map
					_ = _358_v9
					_358_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_354_v5, (_357_v8).Cardinality())
					r3 = ((func() _dafny.Map {
						var _coll10 = _dafny.NewMapBuilder()
						_ = _coll10
						for _iter13 := _dafny.Iterate((_355_v6).Elements()); ; {
							_compr_10, _ok13 := _iter13()
							if !_ok13 {
								break
							}
							var _359_v4 _dafny.Map
							_359_v4 = interface{}(_compr_10).(_dafny.Map)
							if _dafny.Companion_Sequence_.Contains(_355_v6, _359_v4) {
								_coll10.Add(_359_v4, (_this).F32())
							}
						}
						return _coll10.ToMap()
					}()).Merge(_358_v9)).Cardinality()
					r0 = (_this).F32()
					var _360_v10 _dafny.CodePoint
					_ = _360_v10
					_360_v10 = _dafny.CodePoint('m')
					if (_this).Fm17(_360_v10, globalState) {
						var _361_v11 _dafny.Set
						_ = _361_v11
						_361_v11 = _dafny.SetOf((_this).F32())
						(globalState).F16 = (func() bool {
							if false {
								return (_361_v11).IsProperSubsetOf(_361_v11)
							}
							return ((_this).F32()).Cmp((_this).F32()) > 0
						})()
						var _362_v12 _dafny.MultiSet
						_ = _362_v12
						_362_v12 = _dafny.MultiSetOf(((_this).F32()).Cmp((_this).F32()) != 0, ((_this).F32()).Cmp(_dafny.IntOfUint32((_348_v0).Cardinality())) < 0)
						_362_v12 = (func() _dafny.MultiSet {
							if p0 {
								return p1
							}
							return p2
						})()
						(globalState).F28 = p0
						var _363_v13 _dafny.Set
						_ = _363_v13
						_363_v13 = _dafny.SetOf((_this).F31(), p0)
						var _364_v14 D2
						_ = _364_v14
						_364_v14 = Companion_D2_.Create_DC7_(_360_v10, _363_v13)
						var _365_v15 _dafny.Sequence
						_ = _365_v15
						_365_v15 = _dafny.UnicodeSeqOfUtf8Bytes("hxuqqi")
						var _366_v16 _dafny.Map
						_ = _366_v16
						_366_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_364_v14).Dtor_cf10(), _dafny.Companion_Sequence_.Equal(_365_v15, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(726))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg28 _dafny.Int) interface{} {
								return coer28(arg28)
							}
						}((func(_367_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_368_i2 _dafny.Int) _dafny.CodePoint {
								return _367_v10
							}
						})(_360_v10)))))
						_366_v16 = _366_v16
						var _369_v17 *C0
						_ = _369_v17
						var _nw61 *C0 = New_C0_()
						_ = _nw61
						_nw61.Ctor__((_this).F31())
						_369_v17 = _nw61
					} else {
						var _370_v18 _dafny.Sequence
						_ = _370_v18
						_370_v18 = _dafny.SeqOf(p1)
						var _371_v19 _dafny.Array
						_ = _371_v19
						var _nwElement0_14 _dafny.CodePoint = _360_v10
						_ = _nwElement0_14
						var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(5))
						_ = _nw62
						(_nw62).ArraySet1CodePoint(_nwElement0_14, 0)
						(_nw62).ArraySet1CodePoint(_360_v10, 1)
						(_nw62).ArraySet1CodePoint(_360_v10, 2)
						(_nw62).ArraySet1CodePoint(_360_v10, 3)
						(_nw62).ArraySet1CodePoint(_dafny.CodePoint('v'), 4)
						_371_v19 = _nw62
						var _372_v20 D6
						_ = _372_v20
						_372_v20 = Companion_D6_.Create_DC16_(_371_v19)
						var _373_v21 _dafny.Set
						_ = _373_v21
						_373_v21 = _dafny.SetOf(_371_v19, (_372_v20).Dtor_cf24(), _371_v19, _371_v19, _371_v19)
						var _374_v22 _dafny.Map
						_ = _374_v22
						_374_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_370_v18).Select((Companion_Default___.SafeIndex((_373_v21).Cardinality(), _dafny.IntOfUint32((_370_v18).Cardinality()))).Uint32()).(_dafny.MultiSet), (_this).F31())
						var _375_v24 _dafny.Set
						_ = _375_v24
						_375_v24 = _dafny.SetOf(p1, p2)
						var _376_v25 _dafny.Array
						_ = _376_v25
						var _nwElement0_15 _dafny.Map = _374_v22
						_ = _nwElement0_15
						var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(10))
						_ = _nw63
						(_nw63).ArraySet1(_nwElement0_15, 0)
						(_nw63).ArraySet1((_374_v22).Merge(_374_v22), 1)
						(_nw63).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)).Merge(_374_v22), 2)
						(_nw63).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F31()), 3)
						(_nw63).ArraySet1(_374_v22, 4)
						(_nw63).ArraySet1(_374_v22, 5)
						(_nw63).ArraySet1(func() _dafny.Map {
							var _coll11 = _dafny.NewMapBuilder()
							_ = _coll11
							for _iter14 := _dafny.Iterate((_375_v24).Elements()); ; {
								_compr_11, _ok14 := _iter14()
								if !_ok14 {
									break
								}
								var _377_v23 _dafny.MultiSet
								_377_v23 = interface{}(_compr_11).(_dafny.MultiSet)
								if (_375_v24).Contains(_377_v23) {
									_coll11.Add(_377_v23, p0)
								}
							}
							return _coll11.ToMap()
						}(), 6)
						(_nw63).ArraySet1(_374_v22, 7)
						(_nw63).ArraySet1((_374_v22).Merge(_374_v22), 8)
						(_nw63).ArraySet1((_374_v22).Merge(_374_v22), 9)
						_376_v25 = _nw63
						var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_376_v25), 0))
						_ = _index49
						(_376_v25).ArraySet1(func() _dafny.Map {
							var _coll12 = _dafny.NewMapBuilder()
							_ = _coll12
							for _iter15 := _dafny.Iterate((_dafny.SeqOf(p2)).Elements()); ; {
								_compr_12, _ok15 := _iter15()
								if !_ok15 {
									break
								}
								var _378_v26 _dafny.MultiSet
								_378_v26 = interface{}(_compr_12).(_dafny.MultiSet)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p2), _378_v26) {
									_coll12.Add(_378_v26, !(p0))
								}
							}
							return _coll12.ToMap()
						}(), (_index49).Int())
						var _379_v27 _dafny.Map
						_ = _379_v27
						_379_v27 = _374_v22
						var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_376_v25), 0))
						_ = _index50
						(_376_v25).ArraySet1(((func() _dafny.Map {
							if p0 {
								return _379_v27
							}
							return _379_v27
						})()), (_index50).Int())
						var _380_v28 _dafny.Sequence
						_ = _380_v28
						_380_v28 = _dafny.UnicodeSeqOfUtf8Bytes("kyqocej")
						var _381_v29 D3
						_ = _381_v29
						_381_v29 = Companion_D3_.Create_DC10_(_380_v28, _dafny.IntOfInt64(398), false)
						var _382_v30 _dafny.Map
						_ = _382_v30
						_382_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_381_v29).Dtor_cf13(), p0)
						var _383_v31 _dafny.Array
						_ = _383_v31
						var _nwElement0_16 bool = (_this).F31()
						_ = _nwElement0_16
						var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(18))
						_ = _nw64
						(_nw64).ArraySet1(_nwElement0_16, 0)
						(_nw64).ArraySet1(false, 1)
						(_nw64).ArraySet1((_this).F31(), 2)
						(_nw64).ArraySet1(p0, 3)
						(_nw64).ArraySet1((_this).F31(), 4)
						(_nw64).ArraySet1((func() bool {
							if (_382_v30).Contains(_380_v28) {
								return (_382_v30).Get(_380_v28).(bool)
							}
							return !(p0)
						})(), 5)
						(_nw64).ArraySet1((_this).F31(), 6)
						(_nw64).ArraySet1((_this).F31(), 7)
						(_nw64).ArraySet1((_this).F31(), 8)
						(_nw64).ArraySet1((_this).F31(), 9)
						(_nw64).ArraySet1((_this).F31(), 10)
						(_nw64).ArraySet1(false, 11)
						(_nw64).ArraySet1((_dafny.MultiSetFromSeq(_349_v1)).IsDisjointFrom((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(360))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg29 _dafny.Int) interface{} {
								return coer29(arg29)
							}
						}((func(_384_p0 bool) func(_dafny.Int) _dafny.Int {
							return func(_385_i3 _dafny.Int) _dafny.Int {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _384_p0)).Cardinality()
							}
						})(p0))))).Update((_this).F32(), Companion_Default___.Abs((_this).F32()))), 12)
						(_nw64).ArraySet1(p0, 13)
						(_nw64).ArraySet1(((_this).F31()) == (p0), 14)
						(_nw64).ArraySet1(p0, 15)
						(_nw64).ArraySet1(p0, 16)
						(_nw64).ArraySet1(p0, 17)
						_383_v31 = _nw64
						var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_383_v31), 0))
						_ = _index51
						(_383_v31).ArraySet1(((_this).F31()) && ((_this).F31()), (_index51).Int())
						var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_383_v31), 0))
						_ = _index52
						(_383_v31).ArraySet1(false, (_index52).Int())
						(globalState).F16 = (_383_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_383_v31), 0))).Int()).(bool)
						var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_383_v31), 0))
						_ = _index53
						(_383_v31).ArraySet1((func() bool {
							if (_this).F31() {
								return ((_this).F32()).Cmp(_dafny.IntOfInt64(297)) <= 0
							}
							return (_dafny.IntOfUint32((_349_v1).Cardinality())).Cmp((p1).Cardinality()) == 0
						})(), (_index53).Int())
						r0 = (_this).F32()
					}
					var _386_v32 _dafny.Map
					_ = _386_v32
					_386_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), (_this).F32())
					_386_v32 = (_386_v32).Update(_dafny.IntOfInt64(982), _dafny.IntOfInt64(652))
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _387_v33 _dafny.Array
		_ = _387_v33
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_5
		var _nw65 _dafny.Array
		_ = _nw65
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw65 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Int = func(_388_i4 _dafny.Int) _dafny.Int {
				return (_388_i4).Times((_this).F32())
			}
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw65 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw65).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw65).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_387_v33 = _nw65
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_387_v33), 0))
		_ = _index54
		(_387_v33).ArraySet1((_this).F32(), (_index54).Int())
		var _389_v34 _dafny.Map
		_ = _389_v34
		_389_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), !((_this).F31()))
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_387_v33), 0))
		_ = _index55
		(_387_v33).ArraySet1(((_389_v34).Merge((_389_v34).Merge(_389_v34))).Cardinality(), (_index55).Int())
		var _390_v35 _dafny.MultiSet
		_ = _390_v35
		_390_v35 = _dafny.MultiSetOf((_this).F32())
		var _391_v36 _dafny.Array
		_ = _391_v36
		var _nwElement0_17 bool = (_this).F31()
		_ = _nwElement0_17
		var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(4))
		_ = _nw66
		(_nw66).ArraySet1(_nwElement0_17, 0)
		(_nw66).ArraySet1(((_387_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_387_v33), 0))).Int()).(_dafny.Int)).Cmp((_this).F32()) <= 0, 1)
		(_nw66).ArraySet1((_this).F31(), 2)
		(_nw66).ArraySet1(p0, 3)
		_391_v36 = _nw66
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_391_v36), 0))
		_ = _index56
		(_391_v36).ArraySet1(false, (_index56).Int())
		var _392_v37 D3
		_ = _392_v37
		_392_v37 = Companion_D3_.Create_DC10_(_dafny.UnicodeSeqOfUtf8Bytes("cjaamgms"), (_this).F32(), (_this).F31())
		var _393_v38 _dafny.Sequence
		_ = _393_v38
		_393_v38 = _dafny.SeqOf(_390_v35, _390_v35)
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_391_v36), 0))
		_ = _index57
		var _rhs52 _dafny.Sequence = (_392_v37).Dtor_cf13()
		_ = _rhs52
		var _rhs53 _dafny.MultiSet = ((_393_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_349_v1).Cardinality()), _dafny.IntOfUint32((_393_v38).Cardinality()))).Uint32()).(_dafny.MultiSet)).Union((_390_v35).Update((_this).F32(), Companion_Default___.Abs((_this).F32())))
		_ = _rhs53
		var _rhs54 bool = true
		_ = _rhs54
		var _lhs30 _dafny.Array = _391_v36
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_391_v36), 0))
		_ = _lhs31
		r2 = _rhs52
		_390_v35 = _rhs53
		(_lhs30).ArraySet1(_rhs54, (_lhs31).Int())
		r0 = (_this).F32()
		var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(14))
		_ = _nw67
		r1 = _nw67
		var _394_v39 _dafny.Sequence
		_ = _394_v39
		_394_v39 = _dafny.UnicodeSeqOfUtf8Bytes("kip")
		r2 = _394_v39
		r3 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F32(), ((func() _dafny.Int {
			if (p2).Contains((_this).F31()) {
				return (p2).Multiplicity((_this).F31())
			}
			return (_this).F32()
		})()).Plus((_387_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_387_v33), 0))).Int()).(_dafny.Int))))
		return r0, r1, r2, r3
	}
}
func (_this *C1) M5(globalState *GlobalState) (_dafny.Int, bool, _dafny.MultiSet, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r2
		var r3 bool = false
		_ = r3
		r0 = (_this).F32()
		var _395_v0 _dafny.MultiSet
		_ = _395_v0
		_395_v0 = _dafny.MultiSetOf((_this).F32())
		var _396_v1 _dafny.Sequence
		_ = _396_v1
		_396_v1 = _dafny.SeqOf((_395_v0).Cardinality(), (_this).F32())
		var _397_v2 _dafny.Sequence
		_ = _397_v2
		_397_v2 = _dafny.SeqOf(_396_v1, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F32()), _396_v1), _dafny.Companion_Sequence_.Concatenate(_396_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(595))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_398_i0 _dafny.Int) _dafny.Int {
			return (_this).F32()
		}))))
		var _399_v3 _dafny.Array
		_ = _399_v3
		var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
		_ = _nw68
		_399_v3 = _nw68
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
		_ = _index58
		(_399_v3).ArraySet1(_dafny.IntOfInt64(291), (_index58).Int())
		var _400_v4 _dafny.CodePoint
		_ = _400_v4
		_400_v4 = _dafny.CodePoint('o')
		var _401_v5 _dafny.Sequence
		_ = _401_v5
		_401_v5 = _dafny.SeqOf((_this).F31())
		var _402_v6 _dafny.Set
		_ = _402_v6
		_402_v6 = _dafny.SetOf(_401_v5, _dafny.SeqOf((_this).F31(), (_this).F31()), _401_v5, _401_v5, _dafny.SeqOf((_this).F31()))
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
		_ = _index59
		var _rhs55 _dafny.Int = (_this).F32()
		_ = _rhs55
		var _rhs56 _dafny.Sequence = _397_v2
		_ = _rhs56
		var _rhs57 _dafny.Int = (_this).F32()
		_ = _rhs57
		var _rhs58 _dafny.Int = Companion_Default___.Fm18((_this).Fm17(_400_v4, globalState), (_this).F31(), globalState)
		_ = _rhs58
		var _rhs59 _dafny.Int = (((_402_v6).Union(_402_v6)).Cardinality()).Times((_this).F32())
		_ = _rhs59
		var _lhs32 _dafny.Array = _399_v3
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
		_ = _lhs33
		r0 = _rhs55
		_397_v2 = _rhs56
		r0 = _rhs57
		r0 = _rhs58
		(_lhs32).ArraySet1(_rhs59, (_lhs33).Int())
		r1 = (_this).F31()
		var _403_i1 _dafny.Int
		_ = _403_i1
		_403_i1 = _dafny.Zero
		{
			for (_this).F31() {
				{
					if (_403_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_403_i1 = (_403_i1).Plus(_dafny.One)
					var _404_v7 _dafny.Map
					_ = _404_v7
					_404_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int), _401_v5)
					var _405_v10 _dafny.Array
					_ = _405_v10
					var _nwElement0_18 _dafny.Map = _404_v7
					_ = _nwElement0_18
					var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(14))
					_ = _nw69
					(_nw69).ArraySet1(_nwElement0_18, 0)
					(_nw69).ArraySet1((_404_v7).Merge(_404_v7), 1)
					(_nw69).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _dafny.SeqOf((_this).F31())), 2)
					(_nw69).ArraySet1(((_404_v7).Update((_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int), _401_v5)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _401_v5)).Update((_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int), _401_v5)), 3)
					(_nw69).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-777), _401_v5), 4)
					(_nw69).ArraySet1(_404_v7, 5)
					(_nw69).ArraySet1(func() _dafny.Map {
						var _coll13 = _dafny.NewMapBuilder()
						_ = _coll13
						for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(189), _dafny.IntOfInt64(37))); ; {
							_compr_13, _ok16 := _iter16()
							if !_ok16 {
								break
							}
							var _406_v8 _dafny.Int
							_406_v8 = interface{}(_compr_13).(_dafny.Int)
							if ((_dafny.IntOfInt64(189)).Cmp(_406_v8) <= 0) && ((_406_v8).Cmp(_dafny.IntOfInt64(37)) < 0) {
								_coll13.Add(Companion_Default___.SafeModuloInt(_406_v8, _dafny.IntOfInt64(58)), _401_v5)
							}
						}
						return _coll13.ToMap()
					}(), 6)
					(_nw69).ArraySet1(_404_v7, 7)
					(_nw69).ArraySet1(func() _dafny.Map {
						var _coll14 = _dafny.NewMapBuilder()
						_ = _coll14
						for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(719), _dafny.IntOfInt64(649))); ; {
							_compr_14, _ok17 := _iter17()
							if !_ok17 {
								break
							}
							var _407_v9 _dafny.Int
							_407_v9 = interface{}(_compr_14).(_dafny.Int)
							if ((_dafny.IntOfInt64(719)).Cmp(_407_v9) <= 0) && ((_407_v9).Cmp(_dafny.IntOfInt64(649)) < 0) {
								_coll14.Add((_407_v9).Plus((_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int)), _401_v5)
							}
						}
						return _coll14.ToMap()
					}(), 8)
					(_nw69).ArraySet1(_404_v7, 9)
					(_nw69).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int), _401_v5), 10)
					(_nw69).ArraySet1((func() _dafny.Map {
						if (_this).F31() {
							return _404_v7
						}
						return _404_v7
					})(), 11)
					(_nw69).ArraySet1(_404_v7, 12)
					(_nw69).ArraySet1(_404_v7, 13)
					_405_v10 = _nw69
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_405_v10), 0))
					_ = _index60
					(_405_v10).ArraySet1(_404_v7, (_index60).Int())
					var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_405_v10), 0))
					_ = _index61
					(_405_v10).ArraySet1(_404_v7, (_index61).Int())
					_396_v1 = _396_v1
					var _408_v11 _dafny.Sequence
					_ = _408_v11
					_408_v11 = _dafny.UnicodeSeqOfUtf8Bytes("smv")
					_408_v11 = _408_v11
					var _rhs60 _dafny.CodePoint = _400_v4
					_ = _rhs60
					var _rhs61 _dafny.Int = Companion_Default___.Fm18((func() bool {
						if (_this).F31() {
							return (_this).F31()
						}
						return (_this).F31()
					})(), (_this).F31(), globalState)
					_ = _rhs61
					var _rhs62 _dafny.Array = _399_v3
					_ = _rhs62
					_400_v4 = _rhs60
					r0 = _rhs61
					_399_v3 = _rhs62
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _409_v12 _dafny.Sequence
		_ = _409_v12
		_409_v12 = _dafny.UnicodeSeqOfUtf8Bytes("cxidk")
		var _hi2 _dafny.Int = _dafny.IntOfUint32((_409_v12).Cardinality())
		_ = _hi2
		for _410_i2 := (_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int); _410_i2.Cmp(_hi2) < 0; _410_i2 = _410_i2.Plus(_dafny.One) {
			var _411_v13 bool
			_ = _411_v13
			var _412_v14 bool
			_ = _412_v14
			var _413_v15 bool
			_ = _413_v15
			var _out14 bool
			_ = _out14
			var _out15 bool
			_ = _out15
			var _out16 bool
			_ = _out16
			_out14, _out15, _out16 = (_this).M6((_410_i2).Cmp((Companion_Default___.Fm19((_this).F32(), _dafny.IntOfUint32((_409_v12).Cardinality()), globalState)).Cardinality()) == 0, globalState)
			_411_v13 = _out14
			_412_v14 = _out15
			_413_v15 = _out16
			var _source6 D3 = func(_pat_let8_0 D3) D3 {
				return func(_414_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let9_0 bool) D3 {
						return func(_415_dt__update_hcf15_h0 bool) D3 {
							return func(_pat_let10_0 _dafny.Int) D3 {
								return func(_416_dt__update_hcf14_h0 _dafny.Int) D3 {
									return Companion_D3_.Create_DC10_((_414_dt__update__tmp_h0).Dtor_cf13(), _416_dt__update_hcf14_h0, _415_dt__update_hcf15_h0)
								}(_pat_let10_0)
							}((_this).F32())
						}(_pat_let9_0)
					}(true)
				}(_pat_let8_0)
			}(Companion_D3_.Create_DC10_(_409_v12, _dafny.IntOfInt64(377), (_this).F31()))
			_ = _source6
			if _source6.Is_DC10() {
				var _417___mcc_h0 _dafny.Sequence = _source6.Get_().(D3_DC10).Cf13
				_ = _417___mcc_h0
				var _418___mcc_h1 _dafny.Int = _source6.Get_().(D3_DC10).Cf14
				_ = _418___mcc_h1
				var _419___mcc_h2 bool = _source6.Get_().(D3_DC10).Cf15
				_ = _419___mcc_h2
				var _420_cf15 bool = _419___mcc_h2
				_ = _420_cf15
				var _421_cf14 _dafny.Int = _418___mcc_h1
				_ = _421_cf14
				var _422_cf13 _dafny.Sequence = _417___mcc_h0
				_ = _422_cf13
				r0 = _421_cf14
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _index62
				(_399_v3).ArraySet1((_this).F32(), (_index62).Int())
				var _423_v16 _dafny.Array
				_ = _423_v16
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw70
				_423_v16 = _nw70
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_423_v16), 0))
				_ = _index63
				(_423_v16).ArraySet1(_399_v3, (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_423_v16), 0))
				_ = _index64
				(_423_v16).ArraySet1(_399_v3, (_index64).Int())
				var _424_v18 _dafny.Sequence
				_ = _424_v18
				_424_v18 = _dafny.SeqOf(func() _dafny.Map {
					var _coll15 = _dafny.NewMapBuilder()
					_ = _coll15
					for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-16), _dafny.IntOfInt64(-893))); ; {
						_compr_15, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _425_v17 _dafny.Int
						_425_v17 = interface{}(_compr_15).(_dafny.Int)
						if ((_dafny.IntOfInt64(-16)).Cmp(_425_v17) <= 0) && ((_425_v17).Cmp(_dafny.IntOfInt64(-893)) < 0) {
							_coll15.Add((_425_v17).Times(_dafny.IntOfInt64(507)), (_this).F31())
						}
					}
					return _coll15.ToMap()
				}())
				var _426_v19 _dafny.MultiSet
				_ = _426_v19
				_426_v19 = _dafny.MultiSetOf(_411_v13)
				var _427_v20 _dafny.Set
				_ = _427_v20
				_427_v20 = _dafny.SetOf(_410_i2, (_426_v19).Cardinality(), (_this).F32())
				var _428_v21 *C0
				_ = _428_v21
				var _nw71 *C0 = New_C0_()
				_ = _nw71
				_nw71.Ctor__((_427_v20).Contains(_dafny.IntOfUint32((_424_v18).Cardinality())))
				_428_v21 = _nw71
			} else if _source6.Is_DC9() {
				var _429___mcc_h3 _dafny.Sequence = _source6.Get_().(D3_DC9).Cf12
				_ = _429___mcc_h3
				var _430_cf12 _dafny.Sequence = _429___mcc_h3
				_ = _430_cf12
				r0 = (func() _dafny.Int {
					if (_413_v15) && (_413_v15) {
						return (_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int)
					}
					return _410_i2
				})()
				var _431_v22 _dafny.Array
				_ = _431_v22
				var _nw72 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw72
				_431_v22 = _nw72
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_431_v22), 0))
				_ = _index65
				(_431_v22).ArraySet1(_413_v15, (_index65).Int())
				var _432_v23 _dafny.Map
				_ = _432_v23
				_432_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_413_v15, (_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int))
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _index66
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_431_v22), 0))
				_ = _index67
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _index68
				var _rhs63 _dafny.Int = (_this).F32()
				_ = _rhs63
				var _rhs64 bool = _411_v13
				_ = _rhs64
				var _rhs65 bool = (func() bool {
					if !(_432_v23).Contains(_412_v14) {
						return _411_v13
					}
					return (_this).F31()
				})()
				_ = _rhs65
				var _rhs66 _dafny.Int = (_this).F32()
				_ = _rhs66
				var _rhs67 bool = (_this).F31()
				_ = _rhs67
				var _lhs34 _dafny.Array = _399_v3
				_ = _lhs34
				var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _lhs35
				var _lhs36 *GlobalState = globalState
				_ = _lhs36
				var _lhs37 _dafny.Array = _431_v22
				_ = _lhs37
				var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_431_v22), 0))
				_ = _lhs38
				var _lhs39 _dafny.Array = _399_v3
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _lhs40
				(_lhs34).ArraySet1(_rhs63, (_lhs35).Int())
				_lhs36.F16 = _rhs64
				(_lhs37).ArraySet1(_rhs65, (_lhs38).Int())
				(_lhs39).ArraySet1(_rhs66, (_lhs40).Int())
				_411_v13 = _rhs67
				var _433_v24 _dafny.Array
				_ = _433_v24
				var _nwElement0_19 _dafny.Sequence = _401_v5
				_ = _nwElement0_19
				var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(6))
				_ = _nw73
				(_nw73).ArraySet1(_nwElement0_19, 0)
				(_nw73).ArraySet1(_401_v5, 1)
				(_nw73).ArraySet1(_401_v5, 2)
				(_nw73).ArraySet1(_401_v5, 3)
				(_nw73).ArraySet1(_dafny.SeqOf(_411_v13), 4)
				(_nw73).ArraySet1(_401_v5, 5)
				_433_v24 = _nw73
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_433_v24), 0))
				_ = _index69
				(_433_v24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_401_v5, _401_v5), (_index69).Int())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_433_v24), 0))
				_ = _index70
				var _rhs68 _dafny.Int = (_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int)
				_ = _rhs68
				var _rhs69 _dafny.Sequence = _401_v5
				_ = _rhs69
				var _rhs70 bool = _412_v14
				_ = _rhs70
				var _rhs71 bool = _413_v15
				_ = _rhs71
				var _rhs72 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F32(), _410_i2)
				_ = _rhs72
				var _lhs41 _dafny.Array = _433_v24
				_ = _lhs41
				var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_433_v24), 0))
				_ = _lhs42
				var _lhs43 *GlobalState = globalState
				_ = _lhs43
				r0 = _rhs68
				(_lhs41).ArraySet1(_rhs69, (_lhs42).Int())
				_lhs43.F16 = _rhs70
				_411_v13 = _rhs71
				r0 = _rhs72
				var _434_v25 bool
				_ = _434_v25
				var _435_v26 bool
				_ = _435_v26
				var _436_v27 bool
				_ = _436_v27
				var _out17 bool
				_ = _out17
				var _out18 bool
				_ = _out18
				var _out19 bool
				_ = _out19
				_out17, _out18, _out19 = (_this).M6(_412_v14, globalState)
				_434_v25 = _out17
				_435_v26 = _out18
				_436_v27 = _out19
			} else {
				var _437___mcc_h4 D3 = _source6.Get_().(D3_DC11).Cf16
				_ = _437___mcc_h4
				var _438_cf16 D3 = _437___mcc_h4
				_ = _438_cf16
				var _439_v28 _dafny.Set
				_ = _439_v28
				_439_v28 = _dafny.SetOf((_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int))
				var _440_v29 _dafny.Array
				_ = _440_v29
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw74
				_440_v29 = _nw74
				var _441_v30 D5
				_ = _441_v30
				_441_v30 = Companion_D5_.Create_DC14_(_439_v28, _410_i2, true, _440_v29)
				r0 = (_441_v30).Dtor_cf20()
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _index71
				(_399_v3).ArraySet1(_410_i2, (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _index72
				(_399_v3).ArraySet1((_dafny.Zero).Minus((_this).F32()), (_index72).Int())
				var _442_v31 *C0
				_ = _442_v31
				var _nw75 *C0 = New_C0_()
				_ = _nw75
				_nw75.Ctor__(!(_411_v13))
				_442_v31 = _nw75
			}
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
			_ = _index73
			(_399_v3).ArraySet1((_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int), (_index73).Int())
			if _411_v13 {
				_409_v12 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}((func(_443_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_444_i3 _dafny.Int) _dafny.CodePoint {
						return _443_v4
					}
				})(_400_v4)))
				r0 = (_this).F32()
				var _445_v32 _dafny.Set
				_ = _445_v32
				_445_v32 = _dafny.SetOf(_400_v4)
				_445_v32 = _445_v32
				var _446_v33 D0
				_ = _446_v33
				_446_v33 = Companion_D0_.Create_DC1_(_400_v4, (_this).F31(), false)
				var _pat_let_tv5 = _400_v4
				_ = _pat_let_tv5
				var _pat_let_tv6 = _409_v12
				_ = _pat_let_tv6
				var _pat_let_tv7 = _399_v3
				_ = _pat_let_tv7
				var _pat_let_tv8 = _399_v3
				_ = _pat_let_tv8
				_446_v33 = func(_pat_let11_0 D0) D0 {
					return func(_447_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let12_0 _dafny.CodePoint) D0 {
							return func(_448_dt__update_hcf1_h0 _dafny.CodePoint) D0 {
								return func(_pat_let13_0 bool) D0 {
									return func(_449_dt__update_hcf3_h0 bool) D0 {
										return Companion_D0_.Create_DC1_(_448_dt__update_hcf1_h0, (_447_dt__update__tmp_h1).Dtor_cf2(), _449_dt__update_hcf3_h0)
									}(_pat_let13_0)
								}((_dafny.IntOfUint32((_pat_let_tv6).Cardinality())).Cmp((_pat_let_tv8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_pat_let_tv7), 0))).Int()).(_dafny.Int)) > 0)
							}(_pat_let12_0)
						}(_pat_let_tv5)
					}(_pat_let11_0)
				}(_446_v33)
				var _450_v34 *C0
				_ = _450_v34
				var _nw76 *C0 = New_C0_()
				_ = _nw76
				_nw76.Ctor__(_413_v15)
				_450_v34 = _nw76
			} else {
				var _451_v35 _dafny.Map
				_ = _451_v35
				_451_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_395_v0).Cardinality(), (_this).F31())
				var _452_v36 _dafny.Array
				_ = _452_v36
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_6
				var _nw77 _dafny.Array
				_ = _nw77
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw77 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Set = (func(_453_i2 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_454_i4 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_453_i2, _453_i2)
						}
					})(_410_i2)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw77 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw77).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw77).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_452_v36 = _nw77
				var _455_v37 D6
				_ = _455_v37
				_455_v37 = Companion_D6_.Create_DC17_(_413_v15, _452_v36, _410_i2)
				var _456_v38 _dafny.Map
				_ = _456_v38
				_456_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_451_v35, _dafny.SeqOf((_401_v5).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_401_v5).Cardinality()))).Uint32()).(bool), (_455_v37).Dtor_cf25(), _411_v13))
				var _457_v39 _dafny.Array
				_ = _457_v39
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_7
				var _nw78 _dafny.Array
				_ = _nw78
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw78 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.CodePoint = func(_458_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('r')
					}
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw78 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw78).ArraySet1CodePoint(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw78).ArraySet1CodePoint(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_457_v39 = _nw78
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_457_v39), 0))
				_ = _index74
				(_457_v39).ArraySet1CodePoint(_400_v4, (_index74).Int())
				var _459_v40 D0
				_ = _459_v40
				_459_v40 = Companion_D0_.Create_DC2_((_this).F31())
				var _460_v41 _dafny.Map
				_ = _460_v41
				_460_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_459_v40, _400_v4)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_457_v39), 0))
				_ = _index75
				var _rhs73 _dafny.CodePoint = Companion_Default___.Fm20(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gswdyire"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_461_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_462_i6 _dafny.Int) _dafny.CodePoint {
						return _461_v4
					}
				})(_400_v4)))), _395_v0, (_this).F31(), (_this).F32(), globalState)
				_ = _rhs73
				var _rhs74 bool = (_412_v14) == (!(true) || (_411_v13))
				_ = _rhs74
				var _rhs75 _dafny.Map = (_456_v38).Update(_451_v35, _401_v5)
				_ = _rhs75
				var _rhs76 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_460_v41).Contains(_459_v40) {
						return (_460_v41).Get(_459_v40).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('o')
				})()
				_ = _rhs76
				var _lhs44 _dafny.Array = _457_v39
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_457_v39), 0))
				_ = _lhs45
				_400_v4 = _rhs73
				_411_v13 = _rhs74
				_456_v38 = _rhs75
				(_lhs44).ArraySet1CodePoint(_rhs76, (_lhs45).Int())
				var _rhs77 _dafny.CodePoint = _400_v4
				_ = _rhs77
				var _rhs78 bool = _411_v13
				_ = _rhs78
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				_400_v4 = _rhs77
				_lhs46.F16 = _rhs78
				var _pat_let_tv9 = _452_v36
				_ = _pat_let_tv9
				_452_v36 = ((func() D6 {
					if true {
						return func(_pat_let14_0 D6) D6 {
							return func(_463_dt__update__tmp_h2 D6) D6 {
								return func(_pat_let15_0 _dafny.Array) D6 {
									return func(_464_dt__update_hcf26_h0 _dafny.Array) D6 {
										return Companion_D6_.Create_DC17_((_463_dt__update__tmp_h2).Dtor_cf25(), _464_dt__update_hcf26_h0, (_463_dt__update__tmp_h2).Dtor_cf27())
									}(_pat_let15_0)
								}(_pat_let_tv9)
							}(_pat_let14_0)
						}(_455_v37)
					}
					return _455_v37
				})()).Dtor_cf26()
				var _465_v42 D3
				_ = _465_v42
				_465_v42 = Companion_D3_.Create_DC10_(_dafny.UnicodeSeqOfUtf8Bytes("rpccqjlr"), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_409_v12, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.IntOfUint32((_409_v12).Cardinality()))).Uint32(), (_457_v39).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_457_v39), 0))).Int()))).Cardinality()), _411_v13)
				_409_v12 = (_465_v42).Dtor_cf13()
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _index76
				var _rhs79 _dafny.Map = Companion_Default___.Fm21(_395_v0, _412_v14, !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_409_v12, (Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_409_v12).Cardinality()))).Uint32(), Companion_Default___.Fm20((_this).Fm16(_396_v1, _413_v15, _411_v13, _412_v14, globalState), _395_v0, (_this).F31(), (_395_v0).Cardinality(), globalState)), _409_v12), globalState)
				_ = _rhs79
				var _rhs80 _dafny.Int = _410_i2
				_ = _rhs80
				var _lhs47 _dafny.Array = _399_v3
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))
				_ = _lhs48
				_451_v35 = _rhs79
				(_lhs47).ArraySet1(_rhs80, (_lhs48).Int())
			}
		}
		(globalState).F28 = ((_399_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_399_v3), 0))).Int()).(_dafny.Int)).Cmp(((_this).F32()).Minus(_dafny.IntOfInt64(141))) == 0
		r0 = (_this).F32()
		r1 = ((_this).F32()).Cmp(Companion_Default___.Fm18((_this).F31(), !((_this).F31()), globalState)) == 0
		var _466_v43 _dafny.Set
		_ = _466_v43
		_466_v43 = _dafny.SetOf((_this).F32(), _dafny.IntOfUint32((_409_v12).Cardinality()))
		var _467_v44 _dafny.MultiSet
		_ = _467_v44
		_467_v44 = _dafny.MultiSetOf(_466_v43)
		var _468_v45 _dafny.Map
		_ = _468_v45
		_468_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F32()), (_this).F31())
		var _469_v46 _dafny.Map
		_ = _469_v46
		_469_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _468_v45)
		r2 = (func() _dafny.MultiSet {
			if (_this).F31() {
				return (_dafny.MultiSetFromSeq(Companion_Default___.Fm22((_this).F32(), globalState))).Difference(_467_v44)
			}
			return (_467_v44).Union((_467_v44).Update(_466_v43, Companion_Default___.Abs(_dafny.IntOfUint32((Companion_Default___.Fm23(_469_v46, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_409_v12).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("lqta"), globalState)).Cardinality()))))
		})()
		r3 = (_this).F31()
		return r0, r1, r2, r3
	}
}
func (_this *C1) M6(p0 bool, globalState *GlobalState) (bool, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _470_v0 _dafny.Array
		_ = _470_v0
		var _nwElement0_20 _dafny.Int = (_this).F32()
		_ = _nwElement0_20
		var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(6))
		_ = _nw79
		(_nw79).ArraySet1(_nwElement0_20, 0)
		(_nw79).ArraySet1((_this).F32(), 1)
		(_nw79).ArraySet1((_this).F32(), 2)
		(_nw79).ArraySet1((_this).F32(), 3)
		(_nw79).ArraySet1((_this).F32(), 4)
		(_nw79).ArraySet1((_this).F32(), 5)
		_470_v0 = _nw79
		(globalState).F24 = _470_v0
		var _471_v1 _dafny.Sequence
		_ = _471_v1
		_471_v1 = _dafny.UnicodeSeqOfUtf8Bytes("jymesegv")
		var _472_v2 _dafny.MultiSet
		_ = _472_v2
		_472_v2 = _dafny.MultiSetOf((_this).F32())
		var _473_v3 _dafny.Map
		_ = _473_v3
		_473_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F32()), p0)
		var _474_v4 _dafny.Sequence
		_ = _474_v4
		_474_v4 = _dafny.SeqOf(Companion_Default___.Fm21(_472_v2, (_this).F31(), (_this).F31(), globalState), _473_v3, _473_v3)
		var _475_v5 D3
		_ = _475_v5
		_475_v5 = Companion_D3_.Create_DC10_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(162))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}(func(_476_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		})), _471_v1), (_dafny.MultiSetFromSeq(_474_v4)).Cardinality(), (_this).F31())
		var _source7 D3 = _475_v5
		_ = _source7
		if _source7.Is_DC10() {
			var _477___mcc_h0 _dafny.Sequence = _source7.Get_().(D3_DC10).Cf13
			_ = _477___mcc_h0
			var _478___mcc_h1 _dafny.Int = _source7.Get_().(D3_DC10).Cf14
			_ = _478___mcc_h1
			var _479___mcc_h2 bool = _source7.Get_().(D3_DC10).Cf15
			_ = _479___mcc_h2
			var _480_cf15 bool = _479___mcc_h2
			_ = _480_cf15
			var _481_cf14 _dafny.Int = _478___mcc_h1
			_ = _481_cf14
			var _482_cf13 _dafny.Sequence = _477___mcc_h0
			_ = _482_cf13
			_481_cf14 = _481_cf14
			var _483_v6 _dafny.Sequence
			_ = _483_v6
			_483_v6 = _dafny.SeqOf(_480_cf15, (_this).F31())
			var _484_v7 _dafny.Array
			_ = _484_v7
			var _nwElement0_21 _dafny.Sequence = _483_v6
			_ = _nwElement0_21
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(17))
			_ = _nw80
			(_nw80).ArraySet1(_nwElement0_21, 0)
			(_nw80).ArraySet1(_483_v6, 1)
			(_nw80).ArraySet1(_483_v6, 2)
			(_nw80).ArraySet1(_483_v6, 3)
			(_nw80).ArraySet1(_483_v6, 4)
			(_nw80).ArraySet1(_483_v6, 5)
			(_nw80).ArraySet1(_483_v6, 6)
			(_nw80).ArraySet1(_483_v6, 7)
			(_nw80).ArraySet1(_483_v6, 8)
			(_nw80).ArraySet1(_dafny.SeqOf(_480_cf15, false), 9)
			(_nw80).ArraySet1(_483_v6, 10)
			(_nw80).ArraySet1(_483_v6, 11)
			(_nw80).ArraySet1(_483_v6, 12)
			(_nw80).ArraySet1(_483_v6, 13)
			(_nw80).ArraySet1(_483_v6, 14)
			(_nw80).ArraySet1(_483_v6, 15)
			(_nw80).ArraySet1(_dafny.SeqOf(p0), 16)
			_484_v7 = _nw80
			var _485_v8 D0
			_ = _485_v8
			_485_v8 = Companion_D0_.Create_DC0_(_484_v7)
			var _486_v9 D0
			_ = _486_v9
			_486_v9 = Companion_D0_.Create_DC4_(_485_v8)
			var _487_v10 _dafny.Array
			_ = _487_v10
			var _nwElement0_22 D0 = _486_v9
			_ = _nwElement0_22
			var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(5))
			_ = _nw81
			(_nw81).ArraySet1(_nwElement0_22, 0)
			(_nw81).ArraySet1(_486_v9, 1)
			(_nw81).ArraySet1(_486_v9, 2)
			(_nw81).ArraySet1(_486_v9, 3)
			(_nw81).ArraySet1(_486_v9, 4)
			_487_v10 = _nw81
			_487_v10 = _487_v10
			var _488_v11 D0
			_ = _488_v11
			_488_v11 = Companion_D0_.Create_DC2_((_this).F31())
			var _source8 D0 = _488_v11
			_ = _source8
			if _source8.Is_DC1() {
				var _489___mcc_h5 _dafny.CodePoint = _source8.Get_().(D0_DC1).Cf1
				_ = _489___mcc_h5
				var _490___mcc_h6 bool = _source8.Get_().(D0_DC1).Cf2
				_ = _490___mcc_h6
				var _491___mcc_h7 bool = _source8.Get_().(D0_DC1).Cf3
				_ = _491___mcc_h7
				var _492_cf3 bool = _491___mcc_h7
				_ = _492_cf3
				var _493_cf2 bool = _490___mcc_h6
				_ = _493_cf2
				var _494_cf1 _dafny.CodePoint = _489___mcc_h5
				_ = _494_cf1
				var _495_v12 _dafny.Map
				_ = _495_v12
				_495_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_494_cf1, (_this).Fm17(_494_cf1, globalState))
				_495_v12 = _495_v12
				var _496_v13 _dafny.Array
				_ = _496_v13
				var _nw82 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(11))
				_ = _nw82
				_496_v13 = _nw82
				var _497_v14 _dafny.Sequence
				_ = _497_v14
				_497_v14 = _dafny.SeqOf(_496_v13)
				var _498_v15 _dafny.Array
				_ = _498_v15
				var _nwElement0_23 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_496_v13), _dafny.SeqOf(_496_v13, _496_v13)), (Companion_Default___.SafeIndex(_481_cf14, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_496_v13), _dafny.SeqOf(_496_v13, _496_v13))).Cardinality()))).Uint32(), _496_v13)
				_ = _nwElement0_23
				var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(4))
				_ = _nw83
				(_nw83).ArraySet1(_nwElement0_23, 0)
				(_nw83).ArraySet1(_497_v14, 1)
				(_nw83).ArraySet1(_497_v14, 2)
				(_nw83).ArraySet1(_497_v14, 3)
				_498_v15 = _nw83
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_498_v15), 0))
				_ = _index77
				(_498_v15).ArraySet1(_497_v14, (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_498_v15), 0))
				_ = _index78
				(_498_v15).ArraySet1(_497_v14, (_index78).Int())
				var _499_v16 _dafny.Map
				_ = _499_v16
				_499_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(512), _dafny.IntOfUint32((_471_v1).Cardinality()))
				var _500_v17 _dafny.Set
				_ = _500_v17
				_500_v17 = _dafny.SetOf(_499_v16)
				var _501_v18 _dafny.Map
				_ = _501_v18
				_501_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_482_cf13).Cardinality())).Cmp((_this).F32()) == 0, (_500_v17).Union(_500_v17))
				var _502_v21 _dafny.Sequence
				_ = _502_v21
				_502_v21 = _dafny.SeqOf(func() _dafny.Set {
					var _coll16 = _dafny.NewBuilder()
					_ = _coll16
					for _iter19 := _dafny.Iterate((func() _dafny.Set {
						var _coll17 = _dafny.NewBuilder()
						_ = _coll17
						for _iter20 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg34 _dafny.Int) interface{} {
								return coer34(arg34)
							}
						}((func(_503_cf14 _dafny.Int, _504_cf13 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
							return func(_505_i1 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_503_cf14, _dafny.IntOfUint32((_504_cf13).Cardinality()))
							}
						})(_481_cf14, _482_cf13)))).Elements()); ; {
							_compr_17, _ok20 := _iter20()
							if !_ok20 {
								break
							}
							var _506_v19 _dafny.Map
							_506_v19 = interface{}(_compr_17).(_dafny.Map)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
								return func(arg35 _dafny.Int) interface{} {
									return coer35(arg35)
								}
							}((func(_507_cf14 _dafny.Int, _508_cf13 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
								return func(_505_i1 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_507_cf14, _dafny.IntOfUint32((_508_cf13).Cardinality()))
								}
							})(_481_cf14, _482_cf13))), _506_v19) {
								_coll17.Add(_506_v19)
							}
						}
						return _coll17.ToSet()
					}()).Elements()); ; {
						_compr_16, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _509_v20 _dafny.Map
						_509_v20 = interface{}(_compr_16).(_dafny.Map)
						if (func() _dafny.Set {
							var _coll18 = _dafny.NewBuilder()
							_ = _coll18
							for _iter21 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
								return func(arg36 _dafny.Int) interface{} {
									return coer36(arg36)
								}
							}((func(_510_cf14 _dafny.Int, _511_cf13 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
								return func(_505_i1 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_510_cf14, _dafny.IntOfUint32((_511_cf13).Cardinality()))
								}
							})(_481_cf14, _482_cf13)))).Elements()); ; {
								_compr_18, _ok21 := _iter21()
								if !_ok21 {
									break
								}
								var _512_v19 _dafny.Map
								_512_v19 = interface{}(_compr_18).(_dafny.Map)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
									return func(arg37 _dafny.Int) interface{} {
										return coer37(arg37)
									}
								}((func(_513_cf14 _dafny.Int, _514_cf13 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
									return func(_505_i1 _dafny.Int) _dafny.Map {
										return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_513_cf14, _dafny.IntOfUint32((_514_cf13).Cardinality()))
									}
								})(_481_cf14, _482_cf13))), _512_v19) {
									_coll18.Add(_512_v19)
								}
							}
							return _coll18.ToSet()
						}()).Contains(_509_v20) {
							_coll16.Add(_509_v20)
						}
					}
					return _coll16.ToSet()
				}(), (func() _dafny.Set {
					if (_501_v18).Contains(p0) {
						return (_501_v18).Get(p0).(_dafny.Set)
					}
					return _500_v17
				})())
				_501_v18 = (_501_v18).Update(p0, (_502_v21).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_502_v21).Cardinality()))).Uint32()).(_dafny.Set))
				_482_cf13 = _482_cf13
			} else if _source8.Is_DC2() {
				var _515___mcc_h8 bool = _source8.Get_().(D0_DC2).Cf4
				_ = _515___mcc_h8
				var _516_cf4 bool = _515___mcc_h8
				_ = _516_cf4
				var _517_v22 _dafny.CodePoint
				_ = _517_v22
				_517_v22 = _dafny.CodePoint('w')
				var _518_v23 _dafny.MultiSet
				_ = _518_v23
				_518_v23 = _dafny.MultiSetOf(_517_v22, Companion_Default___.Fm20(_471_v1, _dafny.MultiSetOf((_this).F32()), (_this).F31(), (_this).F32(), globalState))
				Companion_Default___.M0((func() _dafny.Int {
					if (_518_v23).Contains(_517_v22) {
						return (_518_v23).Multiplicity(_517_v22)
					}
					return (_this).F32()
				})(), globalState)
				var _519_v24 _dafny.Map
				_ = _519_v24
				_519_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _481_cf14)
				var _520_v25 _dafny.Map
				_ = _520_v25
				_520_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_519_v24).Contains(p0) {
						return (_519_v24).Get(p0).(_dafny.Int)
					}
					return (_this).F32()
				})(), Companion_Default___.Fm18(p0, (_this).F31(), globalState))
				_520_v25 = (_520_v25).Update((_dafny.IntOfUint32((_471_v1).Cardinality())).Plus(_481_cf14), _481_cf14)
				var _521_v26 _dafny.MultiSet
				_ = _521_v26
				_521_v26 = _dafny.MultiSetOf((Companion_D3_.Create_DC10_(_dafny.UnicodeSeqOfUtf8Bytes("hiecrh"), (_this).F32(), _516_cf4)).Dtor_cf13())
				var _522_v27 _dafny.Set
				_ = _522_v27
				_522_v27 = _dafny.SetOf(_516_cf4)
				var _523_v28 _dafny.Map
				_ = _523_v28
				_523_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_521_v26, (_522_v27).Intersection(_522_v27))
				_523_v28 = (_523_v28).Update(_dafny.MultiSetOf(_471_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(735))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_524_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_525_i2 _dafny.Int) _dafny.CodePoint {
						return _524_v22
					}
				})(_517_v22)))), _522_v27)
				var _526_v29 *C0
				_ = _526_v29
				var _nw84 *C0 = New_C0_()
				_ = _nw84
				_nw84.Ctor__(((_this).F31()) == ((_this).Fm17(_517_v22, globalState)))
				_526_v29 = _nw84
			} else if _source8.Is_DC3() {
				var _527___mcc_h9 _dafny.CodePoint = _source8.Get_().(D0_DC3).Cf5
				_ = _527___mcc_h9
				var _528_cf5 _dafny.CodePoint = _527___mcc_h9
				_ = _528_cf5
				var _529_v30 _dafny.Map
				_ = _529_v30
				_529_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_482_cf13).Cardinality())).Minus((_this).F32()), (_472_v2).Cardinality())
				_529_v30 = _529_v30
				var _530_v31 _dafny.Array
				_ = _530_v31
				var _nwElement0_24 bool = !(!((_this).F31())) || ((_this).F31())
				_ = _nwElement0_24
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(9))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_24, 0)
				(_nw85).ArraySet1(!(false) || ((_this).F31()), 1)
				(_nw85).ArraySet1(!(p0), 2)
				(_nw85).ArraySet1((_this).F31(), 3)
				(_nw85).ArraySet1((p0) == (_480_cf15), 4)
				(_nw85).ArraySet1(((_dafny.Zero).Minus((_this).F32())).Cmp(_481_cf14) > 0, 5)
				(_nw85).ArraySet1(p0, 6)
				(_nw85).ArraySet1((_this).F31(), 7)
				(_nw85).ArraySet1(!(_480_cf15), 8)
				_530_v31 = _nw85
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_530_v31), 0))
				_ = _index79
				(_530_v31).ArraySet1(_480_cf15, (_index79).Int())
				var _531_v32 _dafny.Map
				_ = _531_v32
				_531_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_486_v9, _481_cf14)
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_530_v31), 0))
				_ = _index80
				(_530_v31).ArraySet1((_483_v6).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_531_v32).Contains(_486_v9) {
						return (_531_v32).Get(_486_v9).(_dafny.Int)
					}
					return _dafny.IntOfInt64(759)
				})(), _dafny.IntOfUint32((_483_v6).Cardinality()))).Uint32()).(bool), (_index80).Int())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_470_v0), 0))
				_ = _index81
				(_470_v0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_482_cf13, _482_cf13)).Cardinality()), (_index81).Int())
				var _532_v33 _dafny.Sequence
				_ = _532_v33
				_532_v33 = _dafny.SeqOf(_482_cf13, _dafny.UnicodeSeqOfUtf8Bytes("ufldtem"), _482_cf13)
				var _533_v34 _dafny.Array
				_ = _533_v34
				var _nwElement0_25 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("eohrqs")
				_ = _nwElement0_25
				var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(21))
				_ = _nw86
				(_nw86).ArraySet1(_nwElement0_25, 0)
				(_nw86).ArraySet1(_471_v1, 1)
				(_nw86).ArraySet1(_471_v1, 2)
				(_nw86).ArraySet1(_471_v1, 3)
				(_nw86).ArraySet1(_482_cf13, 4)
				(_nw86).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_475_v5).Dtor_cf13(), _482_cf13), 5)
				(_nw86).ArraySet1(_471_v1, 6)
				(_nw86).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vet"), 7)
				(_nw86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_471_v1, _471_v1), 8)
				(_nw86).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("et"), 9)
				(_nw86).ArraySet1(_482_cf13, 10)
				(_nw86).ArraySet1((_532_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.IntOfUint32((_532_v33).Cardinality()))).Uint32()).(_dafny.Sequence), 11)
				(_nw86).ArraySet1(_482_cf13, 12)
				(_nw86).ArraySet1(_471_v1, 13)
				(_nw86).ArraySet1(_471_v1, 14)
				(_nw86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_471_v1, _471_v1), 15)
				(_nw86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_482_cf13, _dafny.UnicodeSeqOfUtf8Bytes("fvmafgib")), 16)
				(_nw86).ArraySet1(_482_cf13, 17)
				(_nw86).ArraySet1(_471_v1, 18)
				(_nw86).ArraySet1((_475_v5).Dtor_cf13(), 19)
				(_nw86).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(902))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_534_cf5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_535_i3 _dafny.Int) _dafny.CodePoint {
						return _534_cf5
					}
				})(_528_cf5))), 20)
				_533_v34 = _nw86
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_533_v34), 0))
				_ = _index82
				(_533_v34).ArraySet1(_471_v1, (_index82).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_470_v0), 0))
				_ = _index83
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_533_v34), 0))
				_ = _index84
				var _rhs81 _dafny.Int = (_this).F32()
				_ = _rhs81
				var _rhs82 _dafny.Int = (_this).F32()
				_ = _rhs82
				var _rhs83 _dafny.Int = (_481_cf14).Times(_dafny.IntOfInt64(-720))
				_ = _rhs83
				var _rhs84 _dafny.Sequence = _482_cf13
				_ = _rhs84
				var _lhs49 _dafny.Array = _470_v0
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_470_v0), 0))
				_ = _lhs50
				var _lhs51 _dafny.Array = _533_v34
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_533_v34), 0))
				_ = _lhs52
				_481_cf14 = _rhs81
				_481_cf14 = _rhs82
				(_lhs49).ArraySet1(_rhs83, (_lhs50).Int())
				(_lhs51).ArraySet1(_rhs84, (_lhs52).Int())
				var _536_v35 D2
				_ = _536_v35
				_536_v35 = Companion_D2_.Create_DC6_((_this).F32())
				var _537_v36 _dafny.Set
				_ = _537_v36
				_537_v36 = _dafny.SetOf(_536_v35)
				var _538_v37 _dafny.Map
				_ = _538_v37
				_538_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_481_cf14, _528_cf5)
				var _539_v38 _dafny.Array
				_ = _539_v38
				var _nwElement0_26 _dafny.CodePoint = _528_cf5
				_ = _nwElement0_26
				var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(13))
				_ = _nw87
				(_nw87).ArraySet1CodePoint(_nwElement0_26, 0)
				(_nw87).ArraySet1CodePoint(_dafny.CodePoint('u'), 1)
				(_nw87).ArraySet1CodePoint(_528_cf5, 2)
				(_nw87).ArraySet1CodePoint(_528_cf5, 3)
				(_nw87).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_538_v37).Contains((_470_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_470_v0), 0))).Int()).(_dafny.Int)) {
						return (_538_v37).Get((_470_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_470_v0), 0))).Int()).(_dafny.Int)).(_dafny.CodePoint)
					}
					return _528_cf5
				})(), 4)
				(_nw87).ArraySet1CodePoint(_528_cf5, 5)
				(_nw87).ArraySet1CodePoint(_528_cf5, 6)
				(_nw87).ArraySet1CodePoint(_528_cf5, 7)
				(_nw87).ArraySet1CodePoint(_528_cf5, 8)
				(_nw87).ArraySet1CodePoint(_528_cf5, 9)
				(_nw87).ArraySet1CodePoint(_dafny.CodePoint('v'), 10)
				(_nw87).ArraySet1CodePoint(_528_cf5, 11)
				(_nw87).ArraySet1CodePoint(_528_cf5, 12)
				_539_v38 = _nw87
				var _540_v39 _dafny.Map
				_ = _540_v39
				_540_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_537_v36, _539_v38)
				_540_v39 = (_540_v39).Update(_537_v36, _539_v38)
			} else if _source8.Is_DC0() {
				var _541___mcc_h10 _dafny.Array = _source8.Get_().(D0_DC0).Cf0
				_ = _541___mcc_h10
				var _542_cf0 _dafny.Array = _541___mcc_h10
				_ = _542_cf0
				_481_cf14 = (_this).F32()
				_481_cf14 = _481_cf14
				var _543_v40 _dafny.Map
				_ = _543_v40
				_543_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_480_cf15, p0)
				var _544_v41 _dafny.Map
				_ = _544_v41
				_544_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_543_v40).Update((_this).F31(), (func() bool {
					if (_543_v40).Contains((Companion_D0_.Create_DC2_((_this).F31())).Dtor_cf4()) {
						return (_543_v40).Get((Companion_D0_.Create_DC2_((_this).F31())).Dtor_cf4()).(bool)
					}
					return p0
				})())).Cardinality(), Companion_Default___.Fm24(_480_cf15, _480_cf15, _dafny.IntOfInt64(681), globalState))
				var _545_v42 _dafny.Map
				_ = _545_v42
				_545_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), ((func() _dafny.MultiSet {
					if (_544_v41).Contains(Companion_Default___.Fm18(_480_cf15, _480_cf15, globalState)) {
						return (_544_v41).Get(Companion_Default___.Fm18(_480_cf15, _480_cf15, globalState)).(_dafny.MultiSet)
					}
					return _472_v2
				})()).Cardinality())
				_545_v42 = (_545_v42).Update(p0, (_this).F32())
				var _546_v43 _dafny.Array
				_ = _546_v43
				var _nw88 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
				_ = _nw88
				_546_v43 = _nw88
				var _547_v44 *C0
				_ = _547_v44
				var _nw89 *C0 = New_C0_()
				_ = _nw89
				_nw89.Ctor__(p0)
				_547_v44 = _nw89
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_546_v43), 0))
				_ = _index85
				(_546_v43).ArraySet1(_547_v44, (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_546_v43), 0))
				_ = _index86
				(_546_v43).ArraySet1(_547_v44, (_index86).Int())
			} else {
				var _548___mcc_h11 D0 = _source8.Get_().(D0_DC4).Cf6
				_ = _548___mcc_h11
				var _549_cf6 D0 = _548___mcc_h11
				_ = _549_cf6
				var _550_v45 _dafny.MultiSet
				_ = _550_v45
				_550_v45 = _dafny.MultiSetOf(p0)
				var _551_v46 _dafny.Map
				_ = _551_v46
				_551_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_483_v6).Select((Companion_Default___.SafeIndex((_550_v45).Cardinality(), _dafny.IntOfUint32((_483_v6).Cardinality()))).Uint32()).(bool), (_this).F31())
				var _552_v47 _dafny.Array
				_ = _552_v47
				var _nwElement0_27 bool = _480_cf15
				_ = _nwElement0_27
				var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(4))
				_ = _nw90
				(_nw90).ArraySet1(_nwElement0_27, 0)
				(_nw90).ArraySet1(_480_cf15, 1)
				(_nw90).ArraySet1((func() bool {
					if (_551_v46).Contains(true) {
						return (_551_v46).Get(true).(bool)
					}
					return (_488_v11).Dtor_cf4()
				})(), 2)
				(_nw90).ArraySet1((_this).F31(), 3)
				_552_v47 = _nw90
				var _553_v48 _dafny.Map
				_ = _553_v48
				_553_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_552_v47, (_dafny.IntOfInt64(693)).Cmp(Companion_Default___.Fm18((_483_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.IntOfUint32((_483_v6).Cardinality()))).Uint32()).(bool), _480_cf15, globalState)) <= 0)
				_553_v48 = (_553_v48).Update(_552_v47, _480_cf15)
				var _554_v49 _dafny.Array
				_ = _554_v49
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
				_ = _nw91
				_554_v49 = _nw91
				var _555_v50 _dafny.Sequence
				_ = _555_v50
				_555_v50 = _dafny.SeqOf(_554_v49)
				var _556_v51 _dafny.Map
				_ = _556_v51
				_556_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm20(_471_v1, _472_v2, !(true), (_this).F32(), globalState), _dafny.IntOfInt64(-279))
				var _557_v52 _dafny.CodePoint
				_ = _557_v52
				_557_v52 = _dafny.CodePoint('u')
				var _rhs85 _dafny.Array = (_555_v50).Select((Companion_Default___.SafeIndex((_556_v51).Cardinality(), _dafny.IntOfUint32((_555_v50).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs85
				var _rhs86 _dafny.Map = _551_v46
				_ = _rhs86
				var _rhs87 bool = ((_550_v45).Difference(_550_v45)).IsProperSubsetOf((_550_v45).Difference(_550_v45))
				_ = _rhs87
				var _rhs88 bool = (_this).Fm17(_557_v52, globalState)
				_ = _rhs88
				var _lhs53 *GlobalState = globalState
				_ = _lhs53
				_554_v49 = _rhs85
				_551_v46 = _rhs86
				_lhs53.F28 = _rhs87
				r0 = _rhs88
				_481_cf14 = _dafny.IntOfInt64(-842)
				_480_cf15 = (_475_v5).Dtor_cf15()
			}
			var _558_v53 _dafny.Sequence
			_ = _558_v53
			_558_v53 = _dafny.SeqOf(_484_v7, _484_v7)
			var _559_v54 D0
			_ = _559_v54
			_559_v54 = Companion_D0_.Create_DC0_((_558_v53).Select((Companion_Default___.SafeIndex(_481_cf14, _dafny.IntOfUint32((_558_v53).Cardinality()))).Uint32()).(_dafny.Array))
			var _560_v55 _dafny.Map
			_ = _560_v55
			_560_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_559_v54, (_this).F31())
			_560_v55 = _560_v55
		} else if _source7.Is_DC9() {
			var _561___mcc_h3 _dafny.Sequence = _source7.Get_().(D3_DC9).Cf12
			_ = _561___mcc_h3
			var _562_cf12 _dafny.Sequence = _561___mcc_h3
			_ = _562_cf12
			Companion_Default___.M0(((_this).F32()).Plus((_dafny.Zero).Minus((_this).F32())), globalState)
			var _563_v56 _dafny.Set
			_ = _563_v56
			_563_v56 = Companion_Default___.Fm25(false, _dafny.IntOfInt64(-932), (_this).F32(), (_this).F31(), globalState)
			var _564_v57 _dafny.MultiSet
			_ = _564_v57
			_564_v57 = _dafny.MultiSetOf((_this).F31(), (_this).F31(), p0)
			var _565_v58 D2
			_ = _565_v58
			_565_v58 = Companion_D2_.Create_DC6_((_564_v57).Cardinality())
			var _566_v59 _dafny.Set
			_ = _566_v59
			_566_v59 = _dafny.SetOf(_565_v58, _565_v58, _565_v58, _565_v58)
			_563_v56 = _566_v59
			(globalState).F16 = p0
			var _567_v60 _dafny.Map
			_ = _567_v60
			_567_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F32())
			var _568_v61 _dafny.Sequence
			_ = _568_v61
			_568_v61 = _dafny.SeqOf((_this).F32(), (_567_v60).Cardinality(), (_this).F32(), (_this).F32(), (_this).F32())
			var _569_v62 _dafny.Set
			_ = _569_v62
			_569_v62 = _dafny.SetOf((_this).F32(), _dafny.IntOfUint32((_568_v61).Cardinality()), (_this).F32(), (_this).F32(), (_472_v2).Cardinality())
			_569_v62 = (_569_v62).Difference(_dafny.SetOf((_this).F32(), (_this).F32()))
		} else {
			var _570___mcc_h4 D3 = _source7.Get_().(D3_DC11).Cf16
			_ = _570___mcc_h4
			var _571_cf16 D3 = _570___mcc_h4
			_ = _571_cf16
			var _572_v63 _dafny.Int
			_ = _572_v63
			_572_v63 = _dafny.IntOfInt64(592)
			_572_v63 = _572_v63
			var _573_v64 _dafny.Sequence
			_ = _573_v64
			_573_v64 = _dafny.SeqOf(p0)
			var _574_v65 _dafny.Set
			_ = _574_v65
			_574_v65 = _dafny.SetOf(_dafny.IntOfUint32((_573_v64).Cardinality()))
			var _575_v66 _dafny.Map
			_ = _575_v66
			_575_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_470_v0, ((_574_v65).Intersection(_574_v65)).Cardinality())
			var _576_v67 _dafny.Array
			_ = _576_v67
			var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw92
			_576_v67 = _nw92
			var _577_v68 _dafny.CodePoint
			_ = _577_v68
			_577_v68 = _dafny.CodePoint('d')
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_576_v67), 0))
			_ = _index87
			(_576_v67).ArraySet1(_dafny.SeqOf(_577_v68), (_index87).Int())
			var _578_v69 _dafny.Sequence
			_ = _578_v69
			_578_v69 = _dafny.SeqOf((_this).F32(), (_this).F32())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_576_v67), 0))
			_ = _index88
			var _rhs89 _dafny.Map = ((_575_v66).Merge(_575_v66)).Merge(((_575_v66).Update(_470_v0, (_this).F32())).Merge(_575_v66))
			_ = _rhs89
			var _rhs90 bool = (_572_v63).Cmp((_578_v69).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(873), _dafny.IntOfUint32((_578_v69).Cardinality()))).Uint32()).(_dafny.Int)) <= 0
			_ = _rhs90
			var _rhs91 _dafny.Sequence = _471_v1
			_ = _rhs91
			var _rhs92 bool = p0
			_ = _rhs92
			var _lhs54 _dafny.Array = _576_v67
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_576_v67), 0))
			_ = _lhs55
			_575_v66 = _rhs89
			r2 = _rhs90
			(_lhs54).ArraySet1(_rhs91, (_lhs55).Int())
			r2 = _rhs92
			var _579_v70 _dafny.Map
			_ = _579_v70
			_579_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_dafny.Zero).Minus((_this).F32()))
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_470_v0), 0))
			_ = _index89
			(_470_v0).ArraySet1((func() _dafny.Int {
				if (_this).F31() {
					return _572_v63
				}
				return (_579_v70).Cardinality()
			})(), (_index89).Int())
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_470_v0), 0))
			_ = _index90
			(_470_v0).ArraySet1((_this).F32(), (_index90).Int())
			Companion_Default___.M0((_470_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_470_v0), 0))).Int()).(_dafny.Int), globalState)
		}
		var _580_v71 _dafny.Set
		_ = _580_v71
		_580_v71 = _dafny.SetOf(_470_v0, _470_v0)
		var _581_v72 _dafny.Map
		_ = _581_v72
		_581_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_580_v71, (_this).F32())
		_581_v72 = (_581_v72).Update(_580_v71, (_this).F32())
		var _582_v73 _dafny.Int
		_ = _582_v73
		_582_v73 = _dafny.IntOfInt64(355)
		_582_v73 = (_582_v73).Minus(_dafny.IntOfUint32((_dafny.SeqOf(p0, p0)).Cardinality()))
		var _583_v74 _dafny.CodePoint
		_ = _583_v74
		_583_v74 = _dafny.CodePoint('b')
		var _584_i4 _dafny.Int
		_ = _584_i4
		_584_i4 = _dafny.Zero
		{
			for !(!(p0) || ((_this).Fm17(_583_v74, globalState))) {
				{
					if (_584_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_584_i4 = (_584_i4).Plus(_dafny.One)
					_471_v1 = _471_v1
					var _585_v75 _dafny.Array
					_ = _585_v75
					var _len0_8 _dafny.Int = _dafny.One
					_ = _len0_8
					var _nw93 _dafny.Array
					_ = _nw93
					if _len0_8.Cmp(_dafny.Zero) == 0 {
						_nw93 = _dafny.NewArray(_len0_8)
					} else {
						var _init8 func(_dafny.Int) _dafny.Sequence = (func(_586_v74 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
							return func(_587_i5 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-543))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg40 _dafny.Int) interface{} {
										return coer40(arg40)
									}
								}((func(_588_v74 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_589_i6 _dafny.Int) _dafny.CodePoint {
										return _588_v74
									}
								})(_586_v74)))
							}
						})(_583_v74)
						_ = _init8
						var _element0_8 = _init8(_dafny.Zero)
						_ = _element0_8
						_nw93 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
						(_nw93).ArraySet1(_element0_8, 0)
						var _nativeLen0_8 = (_len0_8).Int()
						_ = _nativeLen0_8
						for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
							(_nw93).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
						}
					}
					_585_v75 = _nw93
					var _590_v76 _dafny.Array
					_ = _590_v76
					var _nwElement0_28 _dafny.Array = _585_v75
					_ = _nwElement0_28
					var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(6))
					_ = _nw94
					(_nw94).ArraySet1(_nwElement0_28, 0)
					(_nw94).ArraySet1(_585_v75, 1)
					(_nw94).ArraySet1(_585_v75, 2)
					(_nw94).ArraySet1(_585_v75, 3)
					(_nw94).ArraySet1(_585_v75, 4)
					(_nw94).ArraySet1(_585_v75, 5)
					_590_v76 = _nw94
					var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_590_v76), 0))
					_ = _index91
					(_590_v76).ArraySet1(_585_v75, (_index91).Int())
					var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_590_v76), 0))
					_ = _index92
					(_590_v76).ArraySet1(_585_v75, (_index92).Int())
					var _591_v77 *C0
					_ = _591_v77
					var _nw95 *C0 = New_C0_()
					_ = _nw95
					_nw95.Ctor__(_dafny.Companion_Sequence_.Equal(_471_v1, _471_v1))
					_591_v77 = _nw95
					var _rhs93 _dafny.Int = (_this).F32()
					_ = _rhs93
					var _rhs94 bool = !((_this).F31())
					_ = _rhs94
					var _lhs56 *GlobalState = globalState
					_ = _lhs56
					_582_v73 = _rhs93
					_lhs56.F16 = _rhs94
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _592_v78 _dafny.Sequence
		_ = _592_v78
		_592_v78 = _dafny.SeqOf(_582_v73)
		var _593_v79 _dafny.Sequence
		_ = _593_v79
		_593_v79 = _dafny.SeqOf(p0, _dafny.Companion_Sequence_.IsProperPrefixOf(_592_v78, _592_v78))
		_593_v79 = _593_v79
		var _594_v80 _dafny.Map
		_ = _594_v80
		_594_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _582_v73)
		var _595_v81 _dafny.Sequence
		_ = _595_v81
		_595_v81 = _dafny.SeqOf(_594_v80)
		r0 = !_dafny.Companion_Sequence_.Equal(_595_v81, _595_v81)
		r1 = true
		r2 = p0
		return r0, r1, r2
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f31 bool
	_f38 _dafny.Sequence
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f31 = false
	_this._f38 = _dafny.EmptySeq
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

func (_this *C2) F31() bool {
	return _this._f31
}
func (_this *C2) Ctor__(f38 _dafny.Sequence, f31 bool) {
	{
		(_this)._f38 = f38
		(_this)._f31 = f31
	}
}
func (_this *C2) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("my")).Cardinality()))).Cardinality()))).Dtor_cf8()
	}
}
func (_this *C2) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xwst")).Cardinality()))).Minus(_dafny.IntOfInt64(84))
	}
}
func (_this *C2) M4(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) (bool, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _596_v0 _dafny.Sequence
		_ = _596_v0
		_596_v0 = _dafny.UnicodeSeqOfUtf8Bytes("fnnicgsqf")
		var _597_v1 _dafny.CodePoint
		_ = _597_v1
		_597_v1 = _dafny.CodePoint('r')
		var _598_v2 D0
		_ = _598_v2
		_598_v2 = Companion_D0_.Create_DC1_(_597_v1, (_this).F31(), p2)
		var _599_v3 _dafny.Map
		_ = _599_v3
		_599_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_598_v2).Dtor_cf2(), p1)
		var _hi3 _dafny.Int = ((_599_v3).Merge(_599_v3)).Cardinality()
		_ = _hi3
		for _600_i0 := (_dafny.IntOfUint32((_596_v0).Cardinality())).Times(p1); _600_i0.Cmp(_hi3) < 0; _600_i0 = _600_i0.Plus(_dafny.One) {
			var _601_v4 _dafny.Array
			_ = _601_v4
			var _nw96 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw96
			_601_v4 = _nw96
			var _602_v5 _dafny.Map
			_ = _602_v5
			_602_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(910), _601_v4)
			var _603_v6 _dafny.Map
			_ = _603_v6
			_603_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm14(_600_i0, _600_i0, p1, true, globalState)).Cardinality()), p0)
			var _604_v7 _dafny.Set
			_ = _604_v7
			_604_v7 = _dafny.SetOf(p3, false)
			var _605_v8 _dafny.Set
			_ = _605_v8
			_605_v8 = _dafny.SetOf((_604_v7).Cardinality(), p0, _dafny.IntOfInt64(151), p0, p1)
			var _606_v9 D5
			_ = _606_v9
			_606_v9 = Companion_D5_.Create_DC14_(_605_v8, _600_i0, p2, _601_v4)
			var _607_v10 _dafny.Array
			_ = _607_v10
			var _nwElement0_29 _dafny.Array = (func() _dafny.Array {
				if (_602_v5).Contains(((_603_v6).Update(p1, _dafny.IntOfUint32((_596_v0).Cardinality()))).Cardinality()) {
					return (_602_v5).Get(((_603_v6).Update(p1, _dafny.IntOfUint32((_596_v0).Cardinality()))).Cardinality()).(_dafny.Array)
				}
				return _601_v4
			})()
			_ = _nwElement0_29
			var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(23))
			_ = _nw97
			(_nw97).ArraySet1(_nwElement0_29, 0)
			(_nw97).ArraySet1(_601_v4, 1)
			(_nw97).ArraySet1(_601_v4, 2)
			(_nw97).ArraySet1((func() _dafny.Array {
				if (_this).F31() {
					return _601_v4
				}
				return _601_v4
			})(), 3)
			(_nw97).ArraySet1(_601_v4, 4)
			(_nw97).ArraySet1(_601_v4, 5)
			(_nw97).ArraySet1(_601_v4, 6)
			(_nw97).ArraySet1(_601_v4, 7)
			(_nw97).ArraySet1(_601_v4, 8)
			(_nw97).ArraySet1((_606_v9).Dtor_cf22(), 9)
			(_nw97).ArraySet1(_601_v4, 10)
			(_nw97).ArraySet1(_601_v4, 11)
			(_nw97).ArraySet1(_601_v4, 12)
			(_nw97).ArraySet1(_601_v4, 13)
			(_nw97).ArraySet1(_601_v4, 14)
			(_nw97).ArraySet1(_601_v4, 15)
			(_nw97).ArraySet1(_601_v4, 16)
			(_nw97).ArraySet1(_601_v4, 17)
			(_nw97).ArraySet1(_601_v4, 18)
			(_nw97).ArraySet1(_601_v4, 19)
			(_nw97).ArraySet1(_601_v4, 20)
			(_nw97).ArraySet1(_601_v4, 21)
			(_nw97).ArraySet1(_601_v4, 22)
			_607_v10 = _nw97
			_607_v10 = _607_v10
			if p2 {
				var _608_v11 _dafny.Int
				_ = _608_v11
				_608_v11 = _dafny.IntOfInt64(-699)
				var _rhs95 bool = (func() bool {
					if p2 {
						return p2
					}
					return (_604_v7).IsDisjointFrom(_604_v7)
				})()
				_ = _rhs95
				var _rhs96 _dafny.Int = (Companion_D2_.Create_DC6_(p0)).Dtor_cf8()
				_ = _rhs96
				var _rhs97 _dafny.Int = (_603_v6).Cardinality()
				_ = _rhs97
				var _rhs98 _dafny.Int = (_dafny.IntOfInt64(-292)).Minus(_608_v11)
				_ = _rhs98
				var _lhs57 *GlobalState = globalState
				_ = _lhs57
				_lhs57.F28 = _rhs95
				_608_v11 = _rhs96
				_608_v11 = _rhs97
				_608_v11 = _rhs98
				var _609_v12 _dafny.Set
				_ = _609_v12
				_609_v12 = _dafny.SetOf(_597_v1, _597_v1, _597_v1, _597_v1, _597_v1)
				var _610_v13 *C0
				_ = _610_v13
				var _nw98 *C0 = New_C0_()
				_ = _nw98
				_nw98.Ctor__((_609_v12).Equals(_609_v12))
				_610_v13 = _nw98
				_597_v1 = _597_v1
				_608_v11 = (_608_v11).Minus(p1)
				var _611_v14 _dafny.MultiSet
				_ = _611_v14
				_611_v14 = _dafny.MultiSetOf(true, p3, (_606_v9).Dtor_cf21())
				(globalState).F28 = (_611_v14).IsProperSubsetOf(_dafny.MultiSetOf(p2, !(p2), p2))
			} else {
				var _612_v15 _dafny.Map
				_ = _612_v15
				_612_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _613_v16 _dafny.Map
				_ = _613_v16
				_613_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), p2)
				(globalState).F28 = !((func() bool {
					if (_612_v15).Contains(p3) {
						return (_612_v15).Get(p3).(bool)
					}
					return (func() bool {
						if (_613_v16).Contains(_597_v1) {
							return (_613_v16).Get(_597_v1).(bool)
						}
						return (_this).F31()
					})()
				})())
				var _614_v17 _dafny.Int
				_ = _614_v17
				_614_v17 = _dafny.IntOfInt64(100)
				_614_v17 = p1
				(globalState).F16 = p2
				var _615_v18 *C0
				_ = _615_v18
				var _nw99 *C0 = New_C0_()
				_ = _nw99
				_nw99.Ctor__(true)
				_615_v18 = _nw99
				var _616_v19 _dafny.Array
				_ = _616_v19
				var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw100
				_616_v19 = _nw100
				var _617_v20 _dafny.Map
				_ = _617_v20
				_617_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_616_v19, _614_v17)
				var _618_v21 _dafny.Map
				_ = _618_v21
				_618_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_600_i0, _616_v19)
				_617_v20 = (_617_v20).Update((func() _dafny.Array {
					if (_618_v21).Contains(_dafny.IntOfInt64(-299)) {
						return (_618_v21).Get(_dafny.IntOfInt64(-299)).(_dafny.Array)
					}
					return _616_v19
				})(), (_this).Fm13((_dafny.Zero).Minus(_614_v17), globalState))
			}
			(globalState).F28 = p3
			var _619_v22 _dafny.Sequence
			_ = _619_v22
			_619_v22 = _dafny.SeqOf(p0, p0, p1)
			var _620_v23 _dafny.Map
			_ = _620_v23
			_620_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _619_v22)
			_620_v23 = (_620_v23).Update(p2, _dafny.Companion_Sequence_.Concatenate(_619_v22, _dafny.SeqOf((_this).Fm12(p1, p0, _600_i0, _605_v8, globalState))))
		}
		var _hi4 _dafny.Int = p1
		_ = _hi4
		for _621_i1 := p1; _621_i1.Cmp(_hi4) < 0; _621_i1 = _621_i1.Plus(_dafny.One) {
			var _622_v24 _dafny.Map
			_ = _622_v24
			_622_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_621_i1, _621_i1)).Cardinality()), (_this).F31())
			_622_v24 = (_622_v24).Update(p1, (_this).F31())
			var _623_v25 _dafny.Array
			_ = _623_v25
			var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
			_ = _nw101
			_623_v25 = _nw101
			var _624_v26 _dafny.Map
			_ = _624_v26
			_624_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_623_v25, p1)
			_624_v26 = (_624_v26).Update(_623_v25, Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-690))))
			var _625_v27 *C0
			_ = _625_v27
			var _nw102 *C0 = New_C0_()
			_ = _nw102
			_nw102.Ctor__((_this).F31())
			_625_v27 = _nw102
			var _626_v28 _dafny.Array
			_ = _626_v28
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_9
			var _nw103 _dafny.Array
			_ = _nw103
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw103 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = (func(_627_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_628_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_628_i2, _627_p1)
					}
				})(p1)
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
			_626_v28 = _nw103
			var _629_v29 _dafny.Array
			_ = _629_v29
			_629_v29 = _626_v28
			var _source9 _dafny.Array = _629_v29
			_ = _source9
			var _630___mcc_h0 _dafny.Array = _source9
			_ = _630___mcc_h0
			var _631_cf7 _dafny.Array = _630___mcc_h0
			_ = _631_cf7
			var _632_v30 _dafny.Map
			_ = _632_v30
			_632_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), p3)
			var _633_v31 _dafny.Sequence
			_ = _633_v31
			_633_v31 = _dafny.SeqOf(_632_v30)
			var _634_v32 _dafny.MultiSet
			_ = _634_v32
			_634_v32 = _dafny.MultiSetOf(_621_i1, (_this).Fm13(p1, globalState), p1)
			var _635_v33 _dafny.Int
			_ = _635_v33
			_635_v33 = _dafny.IntOfInt64(465)
			var _636_v34 _dafny.Sequence
			_ = _636_v34
			_636_v34 = _dafny.SeqOf((_this).F38())
			var _637_v35 _dafny.Set
			_ = _637_v35
			_637_v35 = _dafny.SetOf(_dafny.IntOfUint32((_636_v34).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_this).F31())).Cardinality()), _621_i1)
			var _rhs99 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_633_v31, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm15(p2, globalState)).Cardinality()), _dafny.IntOfUint32((_633_v31).Cardinality()))).Uint32(), (_632_v30).Merge(_632_v30))
			_ = _rhs99
			var _rhs100 _dafny.MultiSet = _634_v32
			_ = _rhs100
			var _rhs101 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_596_v0, _596_v0), _596_v0)).Cardinality()))
			_ = _rhs101
			var _rhs102 _dafny.Int = (_this).Fm12(_dafny.IntOfUint32((_dafny.SeqOf(true, (_this).F31(), p2, (_this).F31(), p2)).Cardinality()), (_this).Fm13(p0, globalState), p0, (_dafny.SetOf(_621_i1)).Difference(_637_v35), globalState)
			_ = _rhs102
			var _rhs103 bool = p2
			_ = _rhs103
			var _lhs58 *GlobalState = globalState
			_ = _lhs58
			_633_v31 = _rhs99
			_634_v32 = _rhs100
			_635_v33 = _rhs101
			_635_v33 = _rhs102
			_lhs58.F28 = _rhs103
			var _638_v36 _dafny.MultiSet
			_ = _638_v36
			_638_v36 = _dafny.MultiSetOf((_this).F31())
			var _639_v37 _dafny.Map
			_ = _639_v37
			_639_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jhpbetysj"), Companion_Default___.Fm0(_638_v36, globalState)), (Companion_Default___.SafeIndex((_this).Fm13(p1, globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jhpbetysj"), Companion_Default___.Fm0(_638_v36, globalState))).Cardinality()))).Uint32(), _dafny.CodePoint('p')))
			_639_v37 = (_639_v37).Update(true, _596_v0)
			var _rhs104 _dafny.Int = (_this).Fm13(p1, globalState)
			_ = _rhs104
			var _rhs105 _dafny.Int = p1
			_ = _rhs105
			var _rhs106 _dafny.Array = (_629_v29)
			_ = _rhs106
			_635_v33 = _rhs104
			_635_v33 = _rhs105
			_631_cf7 = _rhs106
			var _640_v38 *C0
			_ = _640_v38
			var _nw104 *C0 = New_C0_()
			_ = _nw104
			_nw104.Ctor__((_this).F31())
			_640_v38 = _nw104
		}
		var _641_v39 _dafny.Set
		_ = _641_v39
		_641_v39 = _dafny.SetOf(p0)
		Companion_Default___.M0(((_641_v39).Difference(_641_v39)).Cardinality(), globalState)
		r0 = (func() bool {
			if p3 {
				return (_this).F31()
			}
			return p3
		})()
		_596_v0 = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(891))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}((func(_642_p3 bool, _643_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_644_i3 _dafny.Int) _dafny.CodePoint {
				return (func() _dafny.CodePoint {
					if _642_p3 {
						return _643_v1
					}
					return _643_v1
				})()
			}
		})(p3, _597_v1))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_596_v0).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(891))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg42 _dafny.Int) interface{} {
				return coer42(arg42)
			}
		}((func(_645_p3 bool, _646_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_647_i3 _dafny.Int) _dafny.CodePoint {
				return (func() _dafny.CodePoint {
					if _645_p3 {
						return _646_v1
					}
					return _646_v1
				})()
			}
		})(p3, _597_v1)))).Cardinality()))).Uint32(), _597_v1)
		var _648_i4 _dafny.Int
		_ = _648_i4
		_648_i4 = _dafny.Zero
		{
			for (_this).F31() {
				{
					if (_648_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_648_i4 = (_648_i4).Plus(_dafny.One)
					var _649_v40 _dafny.Array
					_ = _649_v40
					var _nw105 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
					_ = _nw105
					_649_v40 = _nw105
					var _650_v41 _dafny.Map
					_ = _650_v41
					_650_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_649_v40, (_this).Fm13(p0, globalState))
					var _651_v42 _dafny.Sequence
					_ = _651_v42
					_651_v42 = _dafny.SeqOf(p0, _dafny.IntOfUint32(((func() _dafny.Sequence {
						if p3 {
							return _596_v0
						}
						return _596_v0
					})()).Cardinality()), Companion_Default___.SafeDivisionInt((_650_v41).Cardinality(), p0), p0)
					_651_v42 = _651_v42
					var _652_v43 _dafny.Array
					_ = _652_v43
					var _len0_10 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_10
					var _nw106 _dafny.Array
					_ = _nw106
					if _len0_10.Cmp(_dafny.Zero) == 0 {
						_nw106 = _dafny.NewArray(_len0_10)
					} else {
						var _init10 func(_dafny.Int) bool = func(_653_i5 _dafny.Int) bool {
							return true
						}
						_ = _init10
						var _element0_10 = _init10(_dafny.Zero)
						_ = _element0_10
						_nw106 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
						(_nw106).ArraySet1(_element0_10, 0)
						var _nativeLen0_10 = (_len0_10).Int()
						_ = _nativeLen0_10
						for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
							(_nw106).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
						}
					}
					_652_v43 = _nw106
					(globalState).F20 = _652_v43
					r0 = !(p2) || (p2)
					if !_dafny.Companion_Sequence_.Contains((_this).F38(), !(_dafny.MultiSetFromSeq(_651_v42)).Contains(p1)) {
						var _654_v44 _dafny.Map
						_ = _654_v44
						_654_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(46), _597_v1)
						_654_v44 = (_654_v44).Update(p0, _597_v1)
						var _655_v45 _dafny.Int
						_ = _655_v45
						_655_v45 = _dafny.IntOfInt64(825)
						var _656_v46 _dafny.Map
						_ = _656_v46
						_656_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _597_v1)
						_655_v45 = (_this).Fm12(p1, p1, (_this).Fm12((_656_v46).Cardinality(), p0, _dafny.IntOfUint32((_596_v0).Cardinality()), _641_v39, globalState), _641_v39, globalState)
						var _657_v47 _dafny.MultiSet
						_ = _657_v47
						_657_v47 = _dafny.MultiSetOf(!(p2))
						_596_v0 = Companion_Default___.Fm0(_657_v47, globalState)
						var _658_v48 _dafny.Sequence
						_ = _658_v48
						_658_v48 = _dafny.SeqOf(_649_v40, _649_v40, _649_v40)
						var _659_v49 _dafny.Array
						_ = _659_v49
						var _nwElement0_30 _dafny.Array = _652_v43
						_ = _nwElement0_30
						var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(7))
						_ = _nw107
						(_nw107).ArraySet1(_nwElement0_30, 0)
						(_nw107).ArraySet1(_652_v43, 1)
						(_nw107).ArraySet1(_652_v43, 2)
						(_nw107).ArraySet1(_652_v43, 3)
						(_nw107).ArraySet1(_652_v43, 4)
						(_nw107).ArraySet1(_652_v43, 5)
						(_nw107).ArraySet1(_652_v43, 6)
						_659_v49 = _nw107
						var _660_v50 T1
						_ = _660_v50
						var _nw108 *C1 = New_C1_()
						_ = _nw108
						_nw108.Ctor__(p1, _659_v49, (_this).F31())
						_660_v50 = _nw108
						var _661_v51 _dafny.Map
						_ = _661_v51
						_661_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_660_v50, _655_v45)
						_655_v45 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_658_v48, (Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_661_v51).Contains(_660_v50) {
								return (_661_v51).Get(_660_v50).(_dafny.Int)
							}
							return p1
						})(), _dafny.IntOfUint32((_658_v48).Cardinality()))).Uint32(), _649_v40)).Cardinality())
						var _662_v52 _dafny.Map
						_ = _662_v52
						_662_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_660_v50).F32()), Companion_Default___.Fm26((_this).F31(), p2, p3, (_dafny.Zero).Minus((_660_v50).F32()), globalState))
						_662_v52 = (_662_v52).Update((_660_v50).F32(), p2)
					} else {
						var _663_v53 _dafny.Sequence
						_ = _663_v53
						_663_v53 = _dafny.SeqOf(_596_v0)
						_596_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xyaujnuy"), _596_v0), (_663_v53).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_663_v53).Cardinality()))).Uint32()).(_dafny.Sequence))
						var _664_v54 _dafny.Map
						_ = _664_v54
						_664_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_599_v3).Cardinality(), (_dafny.Zero).Minus(p0)), _652_v43)
						_664_v54 = (_664_v54).Merge((Companion_D8_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _652_v43))).Dtor_cf30())
						(globalState).F28 = p2
						var _665_v55 _dafny.Int
						_ = _665_v55
						_665_v55 = _dafny.IntOfInt64(590)
						var _666_v56 _dafny.Map
						_ = _666_v56
						_666_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _596_v0)
						_665_v55 = (_666_v56).Cardinality()
						var _667_v57 _dafny.Array
						_ = _667_v57
						var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
						_ = _nw109
						_667_v57 = _nw109
						_667_v57 = (func() _dafny.Array {
							if (_this).F31() {
								return _667_v57
							}
							return _667_v57
						})()
					}
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		r0 = p3
		var _668_v59 _dafny.Array
		_ = _668_v59
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_11
		var _nw110 _dafny.Array
		_ = _nw110
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw110 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Map = (func(_669_v1 _dafny.CodePoint, _670_p2 bool, _671_p3 bool) func(_dafny.Int) _dafny.Map {
				return func(_672_i6 _dafny.Int) _dafny.Map {
					return func() _dafny.Map {
						var _coll19 = _dafny.NewMapBuilder()
						_ = _coll19
						for _iter22 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC7_(_669_v1, _dafny.SetOf(_670_p2)), false)).Keys().Elements()); ; {
							_compr_19, _ok22 := _iter22()
							if !_ok22 {
								break
							}
							var _673_v58 D2
							_673_v58 = interface{}(_compr_19).(D2)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC7_(_669_v1, _dafny.SetOf(_670_p2)), false)).Contains(_673_v58) {
								_coll19.Add(_673_v58, _671_p3)
							}
						}
						return _coll19.ToMap()
					}()
				}
			})(_597_v1, p2, p3)
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
		_668_v59 = _nw110
		r1 = (func() _dafny.Array {
			if p2 {
				return (func() _dafny.Array {
					if false {
						return _668_v59
					}
					return _668_v59
				})()
			}
			return _668_v59
		})()
		return r0, r1
	}
}
func (_this *C2) F38() _dafny.Sequence {
	{
		return _this._f38
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f31 bool
	_f36 _dafny.Array
	_f37 _dafny.Map
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f31 = false
	_this._f36 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f37 = _dafny.EmptyMap
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

func (_this *C3) F31() bool {
	return _this._f31
}
func (_this *C3) Ctor__(f36 _dafny.Array, f37 _dafny.Map, f31 bool) {
	{
		(_this)._f36 = f36
		(_this)._f37 = f37
		(_this)._f31 = f31
	}
}
func (_this *C3) Fm9(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((func(_source10 D2) _dafny.Sequence {
			if _source10.Is_DC7() {
				var _674___mcc_h0 _dafny.CodePoint = _source10.Get_().(D2_DC7).Cf9
				_ = _674___mcc_h0
				var _675___mcc_h1 _dafny.Set = _source10.Get_().(D2_DC7).Cf10
				_ = _675___mcc_h1
				var _676_cf10 _dafny.Set = _675___mcc_h1
				_ = _676_cf10
				var _677_cf9 _dafny.CodePoint = _674___mcc_h0
				_ = _677_cf9
				if (_this).F31() {
					return _dafny.UnicodeSeqOfUtf8Bytes("tj")
				} else {
					return _dafny.UnicodeSeqOfUtf8Bytes("puwxcvuv")
				}
			} else if _source10.Is_DC6() {
				var _678___mcc_h2 _dafny.Int = _source10.Get_().(D2_DC6).Cf8
				_ = _678___mcc_h2
				var _679_cf8 _dafny.Int = _678___mcc_h2
				_ = _679_cf8
				return _dafny.UnicodeSeqOfUtf8Bytes("rasephr")
			} else {
				var _680___mcc_h3 D2 = _source10.Get_().(D2_DC8).Cf11
				_ = _680___mcc_h3
				var _681_cf11 D2 = _680___mcc_h3
				_ = _681_cf11
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wlyxgncg"), _dafny.UnicodeSeqOfUtf8Bytes("mit"))
			}
		}(Companion_D2_.Create_DC6_((func() _dafny.Map {
			var _coll20 = _dafny.NewMapBuilder()
			_ = _coll20
			for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(691), _dafny.IntOfInt64(135))); ; {
				_compr_20, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _682_v0 _dafny.Int
				_682_v0 = interface{}(_compr_20).(_dafny.Int)
				if ((_dafny.IntOfInt64(691)).Cmp(_682_v0) <= 0) && ((_682_v0).Cmp(_dafny.IntOfInt64(135)) < 0) {
					_coll20.Add((_682_v0).Minus((_dafny.SetOf((_this).F31(), (_this).F31(), (_this).F31(), (_this).F31(), (_this).F31())).Cardinality()), _dafny.IntOfInt64(640))
				}
			}
			return _coll20.ToMap()
		}()).Cardinality()))).Cardinality())
	}
}
func (_this *C3) M3(p0 D0, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) {
	{
		var _683_v0 _dafny.Int
		_ = _683_v0
		_683_v0 = _dafny.IntOfInt64(-523)
		_683_v0 = Companion_Default___.SafeModuloInt(p2, p2)
		if p3 {
			var _684_v1 _dafny.Sequence
			_ = _684_v1
			_684_v1 = _dafny.UnicodeSeqOfUtf8Bytes("uda")
			var _685_v2 _dafny.Array
			_ = _685_v2
			var _nwElement0_31 _dafny.Sequence = _684_v1
			_ = _nwElement0_31
			var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(10))
			_ = _nw111
			(_nw111).ArraySet1(_nwElement0_31, 0)
			(_nw111).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_684_v1, _684_v1), 1)
			(_nw111).ArraySet1(_684_v1, 2)
			(_nw111).ArraySet1(_684_v1, 3)
			(_nw111).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(60))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}(func(_686_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			})), _684_v1), 4)
			(_nw111).ArraySet1(_684_v1, 5)
			(_nw111).ArraySet1(_684_v1, 6)
			(_nw111).ArraySet1(_684_v1, 7)
			(_nw111).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_684_v1, _dafny.UnicodeSeqOfUtf8Bytes("j")), 8)
			(_nw111).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qviqoktq"), 9)
			_685_v2 = _nw111
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_685_v2), 0))
			_ = _index93
			(_685_v2).ArraySet1(_684_v1, (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_685_v2), 0))
			_ = _index94
			(_685_v2).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ldgnc"), (_index94).Int())
			var _687_v4 _dafny.Map
			_ = _687_v4
			_687_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_683_v0, _683_v0)
			var _688_v5 _dafny.MultiSet
			_ = _688_v5
			_688_v5 = _dafny.MultiSetOf((func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter24 := _dafny.Iterate((_687_v4).Keys().Elements()); ; {
					_compr_21, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _689_v3 _dafny.Int
					_689_v3 = interface{}(_compr_21).(_dafny.Int)
					if (_687_v4).Contains(_689_v3) {
						_coll21.Add((_689_v3).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("h"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()))).Uint32(), _dafny.CodePoint('o'))).Cardinality())), (_this).F31())
					}
				}
				return _coll21.ToMap()
			}()).Cardinality())
			var _690_v6 _dafny.Map
			_ = _690_v6
			_690_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_688_v5).IsDisjointFrom(_688_v5), (_this).F31())
			_690_v6 = (_690_v6).Update(p3, p3)
			var _691_v7 _dafny.MultiSet
			_ = _691_v7
			_691_v7 = _dafny.MultiSetOf((_this).F31(), p3)
			_684_v1 = Companion_Default___.Fm0(_691_v7, globalState)
			var _692_v8 _dafny.Array
			_ = _692_v8
			var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
			_ = _nw112
			_692_v8 = _nw112
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_692_v8), 0))
			_ = _index95
			(_692_v8).ArraySet1(_691_v7, (_index95).Int())
			var _693_v9 _dafny.Array
			_ = _693_v9
			var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw113
			_693_v9 = _nw113
			var _694_v10 _dafny.Set
			_ = _694_v10
			_694_v10 = _dafny.SetOf((_this).F31())
			var _695_v11 _dafny.CodePoint
			_ = _695_v11
			_695_v11 = _dafny.CodePoint('d')
			var _696_v12 D0
			_ = _696_v12
			_696_v12 = Companion_D0_.Create_DC1_(_695_v11, (_this).F31(), p3)
			var _697_v13 _dafny.Array
			_ = _697_v13
			var _nwElement0_32 bool = Companion_Default___.Fm10(_695_v11, (_this).F31(), globalState)
			_ = _nwElement0_32
			var _nw114 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
			_ = _nw114
			(_nw114).ArraySet1(_nwElement0_32, 0)
			(_nw114).ArraySet1(false, 1)
			(_nw114).ArraySet1((_this).F31(), 2)
			(_nw114).ArraySet1(p3, 3)
			(_nw114).ArraySet1((_this).F31(), 4)
			(_nw114).ArraySet1(true, 5)
			(_nw114).ArraySet1(p3, 6)
			(_nw114).ArraySet1((_this).F31(), 7)
			(_nw114).ArraySet1(p3, 8)
			(_nw114).ArraySet1((_this).F31(), 9)
			(_nw114).ArraySet1(true, 10)
			(_nw114).ArraySet1(p3, 11)
			(_nw114).ArraySet1((_this).F31(), 12)
			(_nw114).ArraySet1((_this).F31(), 13)
			(_nw114).ArraySet1((_696_v12).Dtor_cf3(), 14)
			(_nw114).ArraySet1(p3, 15)
			(_nw114).ArraySet1(!((_this).F31()), 16)
			(_nw114).ArraySet1((_this).F31(), 17)
			(_nw114).ArraySet1(Companion_Default___.Fm10(_695_v11, (_this).F31(), globalState), 18)
			(_nw114).ArraySet1(true, 19)
			(_nw114).ArraySet1(p3, 20)
			(_nw114).ArraySet1(true, 21)
			(_nw114).ArraySet1(p3, 22)
			(_nw114).ArraySet1(p3, 23)
			(_nw114).ArraySet1(!((_this).F31()), 24)
			_697_v13 = _nw114
			var _698_v14 _dafny.Map
			_ = _698_v14
			_698_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _697_v13)
			var _699_v15 _dafny.Array
			_ = _699_v15
			var _nwElement0_33 _dafny.Int = (_694_v10).Cardinality()
			_ = _nwElement0_33
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(5))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_33, 0)
			(_nw115).ArraySet1((_dafny.SetOf(_683_v0)).Cardinality(), 1)
			(_nw115).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (_691_v7).Contains((_this).F31()) {
					return (_691_v7).Multiplicity((_this).F31())
				}
				return (_698_v14).Cardinality()
			})()), 2)
			(_nw115).ArraySet1(_683_v0, 3)
			(_nw115).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 4)
			_699_v15 = _nw115
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))
			_ = _index96
			(_693_v9).ArraySet1(_699_v15, (_index96).Int())
			var _700_v16 D5
			_ = _700_v16
			_700_v16 = Companion_D5_.Create_DC13_(_691_v7)
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_692_v8), 0))
			_ = _index97
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))
			_ = _index98
			var _rhs107 _dafny.MultiSet = (_700_v16).Dtor_cf18()
			_ = _rhs107
			var _rhs108 _dafny.Array = _699_v15
			_ = _rhs108
			var _lhs59 _dafny.Array = _692_v8
			_ = _lhs59
			var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_692_v8), 0))
			_ = _lhs60
			var _lhs61 _dafny.Array = _693_v9
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))
			_ = _lhs62
			(_lhs59).ArraySet1(_rhs107, (_lhs60).Int())
			(_lhs61).ArraySet1(_rhs108, (_lhs62).Int())
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_693_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))).Int()))
			_ = _arr0
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_dafny.ArrayCastTo((_693_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))).Int()))), 0))
			_ = _index99
			(_arr0).ArraySet1((_dafny.Zero).Minus(p2), (_index99).Int())
			var _701_v17 _dafny.Sequence
			_ = _701_v17
			_701_v17 = _dafny.SeqOf(p2)
			var _702_v18 _dafny.Map
			_ = _702_v18
			_702_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_699_v15, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm11(globalState), _701_v17))
			var _703_v19 _dafny.Map
			_ = _703_v19
			_703_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), p2)
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_693_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))).Int()))
			_ = _arr1
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_dafny.ArrayCastTo((_693_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))).Int()))), 0))
			_ = _index100
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_685_v2), 0))
			_ = _index101
			var _rhs109 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_702_v18).Contains((func() _dafny.Array {
					if p3 {
						return _699_v15
					}
					return _699_v15
				})()) {
					return (_702_v18).Get((func() _dafny.Array {
						if p3 {
							return _699_v15
						}
						return _699_v15
					})()).(_dafny.Sequence)
				}
				return (func() _dafny.Sequence {
					if p3 {
						return _701_v17
					}
					return _701_v17
				})()
			})()).Cardinality())
			_ = _rhs109
			var _rhs110 _dafny.Int = p2
			_ = _rhs110
			var _rhs111 _dafny.Int = (_683_v0).Times(_683_v0)
			_ = _rhs111
			var _rhs112 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_685_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_685_v2), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(p2, _683_v0), _dafny.IntOfUint32(((_685_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_685_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), ((_685_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_685_v2), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_703_v19).Cardinality(), _dafny.IntOfUint32(((_685_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_685_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_ = _rhs112
			var _lhs63 _dafny.Array = _dafny.ArrayCastTo((_693_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))).Int()))
			_ = _lhs63
			var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_dafny.ArrayCastTo((_693_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(270), _dafny.ArrayLen((_693_v9), 0))).Int()))), 0))
			_ = _lhs64
			var _lhs65 _dafny.Array = _685_v2
			_ = _lhs65
			var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_685_v2), 0))
			_ = _lhs66
			_683_v0 = _rhs109
			_683_v0 = _rhs110
			(_lhs63).ArraySet1(_rhs111, (_lhs64).Int())
			(_lhs65).ArraySet1(_rhs112, (_lhs66).Int())
		} else {
			var _704_v20 _dafny.Set
			_ = _704_v20
			_704_v20 = _dafny.SetOf(_683_v0, p2)
			var _705_v21 _dafny.Map
			_ = _705_v21
			_705_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_704_v20).Cardinality())
			var _706_v22 _dafny.MultiSet
			_ = _706_v22
			_706_v22 = _dafny.MultiSetOf(p2, p2, (func() _dafny.Int {
				if (_705_v21).Contains(p3) {
					return (_705_v21).Get(p3).(_dafny.Int)
				}
				return _683_v0
			})())
			var _707_v23 _dafny.Sequence
			_ = _707_v23
			_707_v23 = _dafny.UnicodeSeqOfUtf8Bytes("eys")
			var _708_v24 _dafny.Sequence
			_ = _708_v24
			_708_v24 = _dafny.SeqOf(_683_v0)
			var _709_v25 _dafny.Array
			_ = _709_v25
			var _nwElement0_34 bool = (_this).F31()
			_ = _nwElement0_34
			var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(10))
			_ = _nw116
			(_nw116).ArraySet1(_nwElement0_34, 0)
			(_nw116).ArraySet1(!(_706_v22).Equals(_706_v22), 1)
			(_nw116).ArraySet1((p2).Cmp((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_710_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_711_i1 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_710_p2)
				}
			})(p2))))).Cardinality()) < 0, 2)
			(_nw116).ArraySet1(((_706_v22).Update(_683_v0, Companion_Default___.Abs(_683_v0))).IsSubsetOf(_706_v22), 3)
			(_nw116).ArraySet1(false, 4)
			(_nw116).ArraySet1(true, 5)
			(_nw116).ArraySet1((_683_v0).Cmp(_683_v0) < 0, 6)
			(_nw116).ArraySet1((_this).F31(), 7)
			(_nw116).ArraySet1(_dafny.Companion_Sequence_.Equal(_707_v23, _707_v23), 8)
			(_nw116).ArraySet1(_dafny.Companion_Sequence_.Equal(_708_v24, _708_v24), 9)
			_709_v25 = _nw116
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_709_v25), 0))
			_ = _index102
			(_709_v25).ArraySet1(!(p3), (_index102).Int())
			var _712_v26 _dafny.Array
			_ = _712_v26
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_12
			var _nw117 _dafny.Array
			_ = _nw117
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw117 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.CodePoint = func(_713_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw117 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw117).ArraySet1CodePoint(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw117).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_712_v26 = _nw117
			var _714_v27 _dafny.Map
			_ = _714_v27
			_714_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _712_v26)
			var _715_v28 _dafny.CodePoint
			_ = _715_v28
			_715_v28 = _dafny.CodePoint('h')
			var _716_v29 _dafny.Map
			_ = _716_v29
			_716_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _715_v28)
			var _717_v30 _dafny.Array
			_ = _717_v30
			var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw118
			_717_v30 = _nw118
			var _718_v31 _dafny.Map
			_ = _718_v31
			_718_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_717_v30, p2)
			var _719_v32 _dafny.Map
			_ = _719_v32
			_719_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _717_v30)
			var _720_v33 _dafny.Array
			_ = _720_v33
			var _nwElement0_35 _dafny.Int = (_714_v27).Cardinality()
			_ = _nwElement0_35
			var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(7))
			_ = _nw119
			(_nw119).ArraySet1(_nwElement0_35, 0)
			(_nw119).ArraySet1((_this).Fm9(_dafny.IntOfInt64(977), _708_v24, _683_v0, (_this).F31(), globalState), 1)
			(_nw119).ArraySet1((_this).Fm9((_this).Fm9(_dafny.IntOfInt64(557), _708_v24, _683_v0, p3, globalState), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-475))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg45 _dafny.Int) interface{} {
					return coer45(arg45)
				}
			}(func(_721_i3 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-745)
			})), (Companion_Default___.SafeIndex((_716_v29).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-475))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_722_i3 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-745)
			}))).Cardinality()))).Uint32(), (_718_v31).Cardinality()), p2, (_this).F31(), globalState), 2)
			(_nw119).ArraySet1(_683_v0, 3)
			(_nw119).ArraySet1(((_719_v32).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _717_v30))).Cardinality(), 4)
			(_nw119).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((p1).Cardinality())), 5)
			(_nw119).ArraySet1(_683_v0, 6)
			_720_v33 = _nw119
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))
			_ = _index103
			(_717_v30).ArraySet1(_683_v0, (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_709_v25), 0))
			_ = _index104
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))
			_ = _index105
			var _rhs113 _dafny.Array = _709_v25
			_ = _rhs113
			var _rhs114 _dafny.Int = p2
			_ = _rhs114
			var _rhs115 bool = (func() bool {
				if (_this).F31() {
					return (_this).F31()
				}
				return (_this).F31()
			})()
			_ = _rhs115
			var _rhs116 _dafny.Int = (_683_v0).Minus((_this).Fm9(_683_v0, _708_v24, _dafny.IntOfInt64(927), (_this).F31(), globalState))
			_ = _rhs116
			var _lhs67 *GlobalState = globalState
			_ = _lhs67
			var _lhs68 _dafny.Array = _709_v25
			_ = _lhs68
			var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_709_v25), 0))
			_ = _lhs69
			var _lhs70 _dafny.Array = _717_v30
			_ = _lhs70
			var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))
			_ = _lhs71
			_lhs67.F27 = _rhs113
			_683_v0 = _rhs114
			(_lhs68).ArraySet1(_rhs115, (_lhs69).Int())
			(_lhs70).ArraySet1(_rhs116, (_lhs71).Int())
			var _723_v34 *C0
			_ = _723_v34
			var _nw120 *C0 = New_C0_()
			_ = _nw120
			_nw120.Ctor__((_709_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_709_v25), 0))).Int()).(bool))
			_723_v34 = _nw120
			var _724_v35 _dafny.Array
			_ = _724_v35
			var _nwElement0_36 _dafny.Array = _720_v33
			_ = _nwElement0_36
			var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.One)
			_ = _nw121
			(_nw121).ArraySet1(_nwElement0_36, 0)
			_724_v35 = _nw121
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_724_v35), 0))
			_ = _index106
			(_724_v35).ArraySet1(_720_v33, (_index106).Int())
			var _725_v36 _dafny.Map
			_ = _725_v36
			_725_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_717_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))).Int()).(_dafny.Int))
			var _726_v37 _dafny.Sequence
			_ = _726_v37
			_726_v37 = _dafny.SeqOf((_706_v22).Intersection(_dafny.MultiSetOf(((_725_v36).Update(p2, _dafny.IntOfInt64(975))).Cardinality())), _dafny.MultiSetFromSeq(_708_v24), _706_v22)
			var _727_v38 _dafny.MultiSet
			_ = _727_v38
			_727_v38 = _dafny.MultiSetOf((_this).F31())
			var _728_v39 D8
			_ = _728_v39
			_728_v39 = Companion_D8_.Create_DC21_(p3, (_this).F31(), (_717_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))).Int()).(_dafny.Int))
			var _pat_let_tv10 = _709_v25
			_ = _pat_let_tv10
			var _pat_let_tv11 = _709_v25
			_ = _pat_let_tv11
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_724_v35), 0))
			_ = _index107
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))
			_ = _index108
			var _rhs117 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_726_v37).Cardinality()))
			_ = _rhs117
			var _rhs118 _dafny.Array = (func() _dafny.Array {
				if p3 {
					return _720_v33
				}
				return _720_v33
			})()
			_ = _rhs118
			var _rhs119 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_709_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_709_v25), 0))).Int()).(bool) {
					return _715_v28
				}
				return _715_v28
			})()
			_ = _rhs119
			var _rhs120 _dafny.Int = p2
			_ = _rhs120
			var _rhs121 _dafny.Sequence = Companion_Default___.Fm0((_727_v38).Update((_709_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_709_v25), 0))).Int()).(bool), Companion_Default___.Abs((func(_pat_let16_0 D8) D8 {
				return func(_729_dt__update__tmp_h0 D8) D8 {
					return func(_pat_let17_0 bool) D8 {
						return func(_730_dt__update_hcf31_h0 bool) D8 {
							return Companion_D8_.Create_DC21_(_730_dt__update_hcf31_h0, (_729_dt__update__tmp_h0).Dtor_cf32(), (_729_dt__update__tmp_h0).Dtor_cf33())
						}(_pat_let17_0)
					}((_pat_let_tv11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_pat_let_tv10), 0))).Int()).(bool))
				}(_pat_let16_0)
			}(_728_v39)).Dtor_cf33())), globalState)
			_ = _rhs121
			var _lhs72 _dafny.Array = _724_v35
			_ = _lhs72
			var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_724_v35), 0))
			_ = _lhs73
			var _lhs74 _dafny.Array = _717_v30
			_ = _lhs74
			var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))
			_ = _lhs75
			_683_v0 = _rhs117
			(_lhs72).ArraySet1(_rhs118, (_lhs73).Int())
			_715_v28 = _rhs119
			(_lhs74).ArraySet1(_rhs120, (_lhs75).Int())
			_707_v23 = _rhs121
			var _731_v40 _dafny.Array
			_ = _731_v40
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw122
			_731_v40 = _nw122
			var _732_v41 D5
			_ = _732_v41
			_732_v41 = Companion_D5_.Create_DC14_(_704_v20, _683_v0, p3, _709_v25)
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_731_v40), 0))
			_ = _index109
			(_731_v40).ArraySet1((func() _dafny.Array {
				if p3 {
					return (_732_v41).Dtor_cf22()
				}
				return _709_v25
			})(), (_index109).Int())
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_731_v40), 0))
			_ = _index110
			(_731_v40).ArraySet1(_709_v25, (_index110).Int())
			var _733_v42 _dafny.Sequence
			_ = _733_v42
			_733_v42 = _dafny.SeqOf(_707_v23)
			var _734_v43 _dafny.Sequence
			_ = _734_v43
			_734_v43 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_708_v24, (Companion_Default___.SafeIndex((_717_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_708_v24).Cardinality()))).Uint32(), (_717_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_717_v30), 0))).Int()).(_dafny.Int)), _708_v24)
			var _735_v44 T1
			_ = _735_v44
			var _nw123 *C1 = New_C1_()
			_ = _nw123
			_nw123.Ctor__(_dafny.IntOfUint32((_734_v43).Cardinality()), _731_v40, (_this).F31())
			_735_v44 = _nw123
			var _736_v45 _dafny.Map
			_ = _736_v45
			_736_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_683_v0, _735_v44)
			_733_v42 = (Companion_Default___.Fm27(p3, (_736_v45).Cardinality(), (_this).F31(), globalState)).Dtor_cf39()
		}
		var _737_v46 _dafny.MultiSet
		_ = _737_v46
		_737_v46 = _dafny.MultiSetOf(p3, (_this).F31(), false)
		var _738_v47 _dafny.Sequence
		_ = _738_v47
		_738_v47 = _dafny.SeqOf((_737_v46).IsSubsetOf(_737_v46))
		var _739_v48 _dafny.Map
		_ = _739_v48
		_739_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
		var _740_v49 _dafny.Sequence
		_ = _740_v49
		_740_v49 = _dafny.UnicodeSeqOfUtf8Bytes("ugtivxrgy")
		var _741_v50 D9
		_ = _741_v50
		_741_v50 = Companion_D9_.Create_DC26_(Companion_Default___.Fm26(p3, p3, (func() bool {
			if (_739_v48).Contains(_683_v0) {
				return (_739_v48).Get(_683_v0).(bool)
			}
			return p3
		})(), p2, globalState), _740_v49, _683_v0, _739_v48, p3)
		var _pat_let_tv12 = p1
		_ = _pat_let_tv12
		var _pat_let_tv13 = p1
		_ = _pat_let_tv13
		var _pat_let_tv14 = _738_v47
		_ = _pat_let_tv14
		var _pat_let_tv15 = _738_v47
		_ = _pat_let_tv15
		var _pat_let_tv16 = p1
		_ = _pat_let_tv16
		_738_v47 = func(_source11 D9) _dafny.Sequence {
			if _source11.Is_DC24() {
				var _742___mcc_h0 _dafny.MultiSet = _source11.Get_().(D9_DC24).Cf40
				_ = _742___mcc_h0
				var _743_cf40 _dafny.MultiSet = _742___mcc_h0
				_ = _743_cf40
				return _pat_let_tv12
			} else if _source11.Is_DC25() {
				var _744___mcc_h1 _dafny.Int = _source11.Get_().(D9_DC25).Cf41
				_ = _744___mcc_h1
				var _745___mcc_h2 bool = _source11.Get_().(D9_DC25).Cf42
				_ = _745___mcc_h2
				var _746___mcc_h3 bool = _source11.Get_().(D9_DC25).Cf43
				_ = _746___mcc_h3
				var _747___mcc_h4 bool = _source11.Get_().(D9_DC25).Cf44
				_ = _747___mcc_h4
				var _748_cf44 bool = _747___mcc_h4
				_ = _748_cf44
				var _749_cf43 bool = _746___mcc_h3
				_ = _749_cf43
				var _750_cf42 bool = _745___mcc_h2
				_ = _750_cf42
				var _751_cf41 _dafny.Int = _744___mcc_h1
				_ = _751_cf41
				return _pat_let_tv13
			} else if _source11.Is_DC26() {
				var _752___mcc_h5 bool = _source11.Get_().(D9_DC26).Cf45
				_ = _752___mcc_h5
				var _753___mcc_h6 _dafny.Sequence = _source11.Get_().(D9_DC26).Cf46
				_ = _753___mcc_h6
				var _754___mcc_h7 _dafny.Int = _source11.Get_().(D9_DC26).Cf47
				_ = _754___mcc_h7
				var _755___mcc_h8 _dafny.Map = _source11.Get_().(D9_DC26).Cf48
				_ = _755___mcc_h8
				var _756___mcc_h9 bool = _source11.Get_().(D9_DC26).Cf49
				_ = _756___mcc_h9
				var _757_cf49 bool = _756___mcc_h9
				_ = _757_cf49
				var _758_cf48 _dafny.Map = _755___mcc_h8
				_ = _758_cf48
				var _759_cf47 _dafny.Int = _754___mcc_h7
				_ = _759_cf47
				var _760_cf46 _dafny.Sequence = _753___mcc_h6
				_ = _760_cf46
				var _761_cf45 bool = _752___mcc_h5
				_ = _761_cf45
				return _pat_let_tv14
			} else {
				var _762___mcc_h10 _dafny.Sequence = _source11.Get_().(D9_DC23).Cf39
				_ = _762___mcc_h10
				var _763_cf39 _dafny.Sequence = _762___mcc_h10
				_ = _763_cf39
				return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv15, _pat_let_tv16)
			}
		}(_741_v50)
		_683_v0 = p2
		var _764_v51 _dafny.Array
		_ = _764_v51
		var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
		_ = _nw124
		_764_v51 = _nw124
		for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_764_v51), 0))); ; {
			_guard_loop_3, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _765_i4 _dafny.Int
			_765_i4 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_765_i4).Sign() != -1) && ((_765_i4).Cmp(_dafny.ArrayLen((_764_v51), 0)) < 0)) {
				(_764_v51).ArraySet1((_765_i4).Times((_dafny.IntOfUint32((p1).Cardinality())).Times(_dafny.IntOfInt64(47))), (_765_i4).Int())
			}
		}
		if (p2).Cmp((_683_v0).Times((_dafny.Zero).Minus(p2))) >= 0 {
			Companion_Default___.M0(p2, globalState)
			var _766_v52 _dafny.MultiSet
			_ = _766_v52
			_766_v52 = _dafny.MultiSetOf(p1, _738_v47, _738_v47)
			(globalState).F16 = (_766_v52).IsDisjointFrom(_766_v52)
			var _767_v53 _dafny.Map
			_ = _767_v53
			_767_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _683_v0)
			_767_v53 = _767_v53
			var _768_v54 D2
			_ = _768_v54
			_768_v54 = Companion_D2_.Create_DC6_(p2)
			var _769_v55 _dafny.Set
			_ = _769_v55
			_769_v55 = _dafny.SetOf(_768_v54)
			var _770_v56 _dafny.Set
			_ = _770_v56
			_770_v56 = _769_v55
			var _source12 _dafny.Set = _770_v56
			_ = _source12
			var _771___mcc_h11 _dafny.Set = _source12
			_ = _771___mcc_h11
			var _772_cf17 _dafny.Set = _771___mcc_h11
			_ = _772_cf17
			var _773_v57 _dafny.Sequence
			_ = _773_v57
			_773_v57 = _dafny.SeqOf(_741_v50)
			var _774_v59 _dafny.MultiSet
			_ = _774_v59
			_774_v59 = _dafny.MultiSetOf(func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter26 := _dafny.Iterate((_773_v57).Elements()); ; {
					_compr_22, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _775_v58 D9
					_775_v58 = interface{}(_compr_22).(D9)
					if _dafny.Companion_Sequence_.Contains(_773_v57, _775_v58) {
						_coll22.Add(_775_v58)
					}
				}
				return _coll22.ToSet()
			}())
			var _776_v60 _dafny.Map
			_ = _776_v60
			_776_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v59, (_dafny.Zero).Minus((p2).Times((_dafny.Zero).Minus(p2))))
			var _777_v61 _dafny.Sequence
			_ = _777_v61
			_777_v61 = _dafny.SeqOf(_774_v59)
			_683_v0 = (func() _dafny.Int {
				if (_776_v60).Contains(((_777_v61).Select((Companion_Default___.SafeIndex(_683_v0, _dafny.IntOfUint32((_777_v61).Cardinality()))).Uint32()).(_dafny.MultiSet)).Union(_774_v59)) {
					return (_776_v60).Get(((_777_v61).Select((Companion_Default___.SafeIndex(_683_v0, _dafny.IntOfUint32((_777_v61).Cardinality()))).Uint32()).(_dafny.MultiSet)).Union(_774_v59)).(_dafny.Int)
				}
				return _683_v0
			})()
			var _778_v62 _dafny.Set
			_ = _778_v62
			_778_v62 = _dafny.SetOf(true)
			var _779_v63 _dafny.Array
			_ = _779_v63
			var _nwElement0_37 _dafny.Set = (_dafny.SetOf(p3)).Intersection(_778_v62)
			_ = _nwElement0_37
			var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.One)
			_ = _nw125
			(_nw125).ArraySet1(_nwElement0_37, 0)
			_779_v63 = _nw125
			_779_v63 = _779_v63
			_683_v0 = p2
			var _780_v64 _dafny.Map
			_ = _780_v64
			_780_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_740_v49).Cardinality()), _683_v0)
			var _781_v65 _dafny.Map
			_ = _781_v65
			_781_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_780_v64).Cardinality(), _740_v49)
			var _782_v66 _dafny.Map
			_ = _782_v66
			_782_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _740_v49)
			var _783_v67 _dafny.CodePoint
			_ = _783_v67
			_783_v67 = _dafny.CodePoint('k')
			var _784_v68 _dafny.Array
			_ = _784_v68
			var _nwElement0_38 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_740_v49, (func() _dafny.Sequence {
				if (_781_v65).Contains(p2) {
					return (_781_v65).Get(p2).(_dafny.Sequence)
				}
				return _740_v49
			})())
			_ = _nwElement0_38
			var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(27))
			_ = _nw126
			(_nw126).ArraySet1(_nwElement0_38, 0)
			(_nw126).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lyh"), 1)
			(_nw126).ArraySet1(_740_v49, 2)
			(_nw126).ArraySet1(_740_v49, 3)
			(_nw126).ArraySet1((func() _dafny.Sequence {
				if (_782_v66).Contains((_738_v47).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_738_v47).Cardinality()))).Uint32()).(bool)) {
					return (_782_v66).Get((_738_v47).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_738_v47).Cardinality()))).Uint32()).(bool)).(_dafny.Sequence)
				}
				return _740_v49
			})(), 4)
			(_nw126).ArraySet1(_740_v49, 5)
			(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_740_v49, _740_v49), 6)
			(_nw126).ArraySet1(_740_v49, 7)
			(_nw126).ArraySet1((func() _dafny.Sequence {
				if p3 {
					return _740_v49
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("icxd")
			})(), 8)
			(_nw126).ArraySet1(_740_v49, 9)
			(_nw126).ArraySet1(_740_v49, 10)
			(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm0(_737_v46, globalState), _740_v49), 11)
			(_nw126).ArraySet1(_740_v49, 12)
			(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_740_v49, _740_v49), 13)
			(_nw126).ArraySet1(_740_v49, 14)
			(_nw126).ArraySet1(_740_v49, 15)
			(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_740_v49, _740_v49), 16)
			(_nw126).ArraySet1(_740_v49, 17)
			(_nw126).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("eghuqqi"), 18)
			(_nw126).ArraySet1(_740_v49, 19)
			(_nw126).ArraySet1(_dafny.Companion_Sequence_.Update(_740_v49, (Companion_Default___.SafeIndex(_683_v0, _dafny.IntOfUint32((_740_v49).Cardinality()))).Uint32(), _783_v67), 20)
			(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_740_v49, _740_v49), 21)
			(_nw126).ArraySet1(_740_v49, 22)
			(_nw126).ArraySet1(_740_v49, 23)
			(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_740_v49, _740_v49), 24)
			(_nw126).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("s"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()))).Uint32(), _783_v67), _740_v49), 25)
			(_nw126).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(493))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_785_v67 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_786_i5 _dafny.Int) _dafny.CodePoint {
					return _785_v67
				}
			})(_783_v67))), 26)
			_784_v68 = _nw126
			_784_v68 = _784_v68
			var _787_v69 _dafny.Array
			_ = _787_v69
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_13
			var _nw127 _dafny.Array
			_ = _nw127
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw127 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) bool = func(_788_i6 _dafny.Int) bool {
					return (_this).F31()
				}
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw127 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw127).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw127).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_787_v69 = _nw127
			var _789_v70 _dafny.Set
			_ = _789_v70
			_789_v70 = _dafny.SetOf(_dafny.IntOfUint32((_740_v49).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(p3)).Cardinality()))
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_787_v69), 0))
			_ = _index111
			(_787_v69).ArraySet1((_789_v70).IsSubsetOf(_789_v70), (_index111).Int())
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_787_v69), 0))
			_ = _index112
			(_787_v69).ArraySet1((_this).F31(), (_index112).Int())
		} else {
			var _790_v71 _dafny.Array
			_ = _790_v71
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_14
			var _nw128 _dafny.Array
			_ = _nw128
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw128 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.MultiSet = (func(_791_v0 _dafny.Int, _792_p2 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
					return func(_793_i7 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_791_v0, _792_p2))).Union(_dafny.MultiSetOf(_791_v0, _dafny.IntOfInt64(213), _792_p2))
					}
				})(_683_v0, p2)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw128 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw128).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw128).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_790_v71 = _nw128
			var _794_v72 _dafny.MultiSet
			_ = _794_v72
			_794_v72 = _dafny.MultiSetOf(p2, p2)
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_790_v71), 0))
			_ = _index113
			(_790_v71).ArraySet1((_794_v72).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_738_v47).Cardinality()), _683_v0)), (_index113).Int())
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((_790_v71), 0))
			_ = _index114
			(_790_v71).ArraySet1((_794_v72).Update(p2, Companion_Default___.Abs(_dafny.IntOfInt64(955))), (_index114).Int())
			_683_v0 = (p2).Minus(_dafny.IntOfInt64(-367))
			_740_v49 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_740_v49, _740_v49), _740_v49)
			var _795_v73 _dafny.Map
			_ = _795_v73
			_795_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_740_v49).Select((Companion_Default___.SafeIndex(_683_v0, _dafny.IntOfUint32((_740_v49).Cardinality()))).Uint32()).(_dafny.CodePoint), _683_v0)
			var _796_v74 _dafny.CodePoint
			_ = _796_v74
			_796_v74 = _dafny.CodePoint('n')
			_795_v73 = (_795_v73).Update(_796_v74, (_this).Fm9((_741_v50).Dtor_cf47(), _dafny.SeqOf(_683_v0, _683_v0), _dafny.IntOfUint32((_738_v47).Cardinality()), (_this).F31(), globalState))
			Companion_Default___.M0(p2, globalState)
		}
	}
}
func (_this *C3) F36() _dafny.Array {
	{
		return _this._f36
	}
}
func (_this *C3) F37() _dafny.Map {
	{
		return _this._f37
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f31 bool
	_f32 _dafny.Int
	_f33 _dafny.Array
	F35  bool
	_f34 _dafny.Sequence
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f31 = false
	_this._f32 = _dafny.Zero
	_this._f33 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F35 = false
	_this._f34 = _dafny.EmptySeq
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

func (_this *C4) F31() bool {
	return _this._f31
}
func (_this *C4) F32() _dafny.Int {
	return _this._f32
}
func (_this *C4) F33() _dafny.Array {
	return _this._f33
}
func (_this *C4) Ctor__(f34 _dafny.Sequence, f35 bool, f32 _dafny.Int, f33 _dafny.Array, f31 bool) {
	{
		(_this)._f34 = f34
		(_this).F35 = f35
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f31 = f31
	}
}
func (_this *C4) Fm6(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F35, (_this).F31()), _dafny.SeqOf(true)), _dafny.SeqOf(!(_this.F35)))
	}
}
func (_this *C4) M2(p0 bool, p1 _dafny.MultiSet, p2 _dafny.MultiSet, globalState *GlobalState) (_dafny.Int, _dafny.Array, _dafny.Sequence, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _797_i0 _dafny.Int
		_ = _797_i0
		_797_i0 = _dafny.Zero
		{
			for !((Companion_Default___.SafeModuloInt((_this).F32(), (_this).F32())).Cmp((_this).F32()) < 0) {
				{
					if (_797_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_797_i0 = (_797_i0).Plus(_dafny.One)
					(globalState).F28 = (_this).F31()
					r0 = _dafny.IntOfInt64(462)
					var _798_v1 _dafny.Sequence
					_ = _798_v1
					_798_v1 = _dafny.SeqOf((_this).F32(), (_dafny.Zero).Minus(((func() _dafny.Set {
						var _coll23 = _dafny.NewBuilder()
						_ = _coll23
						for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(180), _dafny.IntOfInt64(462))); ; {
							_compr_23, _ok27 := _iter27()
							if !_ok27 {
								break
							}
							var _799_v0 _dafny.Int
							_799_v0 = interface{}(_compr_23).(_dafny.Int)
							if ((_dafny.IntOfInt64(180)).Cmp(_799_v0) <= 0) && ((_799_v0).Cmp(_dafny.IntOfInt64(462)) < 0) {
								_coll23.Add((_799_v0).Minus((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F35))).Cardinality()))
							}
						}
						return _coll23.ToSet()
					}()).Cardinality()).Minus((_this).F32())), (_this).F32())
					_798_v1 = _798_v1
					r3 = Companion_Default___.Fm7((_this).F31(), (_this).F31(), globalState)
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _800_v2 _dafny.Sequence
		_ = _800_v2
		_800_v2 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		var _801_v3 _dafny.MultiSet
		_ = _801_v3
		_801_v3 = _dafny.MultiSetOf(_800_v2)
		var _802_v4 _dafny.Sequence
		_ = _802_v4
		_802_v4 = _dafny.SeqOf(p0)
		var _803_v5 _dafny.Sequence
		_ = _803_v5
		_803_v5 = _dafny.SeqOf((_this).F32(), (_this).F32())
		var _804_v6 _dafny.Sequence
		_ = _804_v6
		_804_v6 = _dafny.SeqOf((func() _dafny.Int {
			if (_801_v3).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(876))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}(func(_805_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('w')
			}))) {
				return (_801_v3).Multiplicity(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(876))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}(func(_806_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				})))
			}
			return _dafny.IntOfUint32((_802_v4).Cardinality())
		})(), _dafny.IntOfUint32((_803_v5).Cardinality()), (_this).F32(), (_this).F32(), _dafny.IntOfInt64(214))
		var _807_v7 _dafny.Set
		_ = _807_v7
		_807_v7 = _dafny.SetOf(_803_v5, _804_v6)
		if ((_807_v7).Difference(_807_v7)).IsSubsetOf((_807_v7).Difference(_807_v7)) {
			(globalState).F16 = _this.F35
			var _808_v8 _dafny.Array
			_ = _808_v8
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_15
			var _nw129 _dafny.Array
			_ = _nw129
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw129 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Sequence = (func(_809_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_810_i2 _dafny.Int) _dafny.Sequence {
						return _809_v4
					}
				})(_802_v4)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw129 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw129).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw129).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_808_v8 = _nw129
			var _811_v9 D0
			_ = _811_v9
			_811_v9 = Companion_D0_.Create_DC0_(_808_v8)
			var _812_v10 _dafny.Set
			_ = _812_v10
			_812_v10 = _dafny.SetOf(_this.F35)
			var _pat_let_tv17 = _808_v8
			_ = _pat_let_tv17
			var _pat_let_tv18 = _808_v8
			_ = _pat_let_tv18
			var _rhs122 _dafny.Int = (_812_v10).Cardinality()
			_ = _rhs122
			var _rhs123 D0 = func(_pat_let18_0 D0) D0 {
				return func(_815_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let21_0 _dafny.Array) D0 {
						return func(_816_dt__update_hcf0_h1 _dafny.Array) D0 {
							return Companion_D0_.Create_DC0_(_816_dt__update_hcf0_h1)
						}(_pat_let21_0)
					}(_pat_let_tv18)
				}(_pat_let18_0)
			}(func(_pat_let19_0 D0) D0 {
				return func(_813_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let20_0 _dafny.Array) D0 {
						return func(_814_dt__update_hcf0_h0 _dafny.Array) D0 {
							return Companion_D0_.Create_DC0_(_814_dt__update_hcf0_h0)
						}(_pat_let20_0)
					}(_pat_let_tv17)
				}(_pat_let19_0)
			}(_811_v9))
			_ = _rhs123
			r3 = _rhs122
			_811_v9 = _rhs123
			_802_v4 = _802_v4
			var _817_v11 _dafny.Array
			_ = _817_v11
			var _nw130 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw130
			_817_v11 = _nw130
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_817_v11), 0))
			_ = _index115
			(_817_v11).ArraySet1(Companion_Default___.Fm8(p0, _dafny.IntOfInt64(989), globalState), (_index115).Int())
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_817_v11), 0))
			_ = _index116
			(_817_v11).ArraySet1((_this).F31(), (_index116).Int())
			var _818_v12 _dafny.MultiSet
			_ = _818_v12
			_818_v12 = _dafny.MultiSetOf(_dafny.IntOfInt64(-823), (_this).F32(), (_this).F32())
			var _819_v13 _dafny.Array
			_ = _819_v13
			var _nwElement0_39 _dafny.Int = ((_818_v12).Intersection(_818_v12)).Cardinality()
			_ = _nwElement0_39
			var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(5))
			_ = _nw131
			(_nw131).ArraySet1(_nwElement0_39, 0)
			(_nw131).ArraySet1((_this).F32(), 1)
			(_nw131).ArraySet1((_dafny.Zero).Minus((_this).F32()), 2)
			(_nw131).ArraySet1(_dafny.IntOfInt64(592), 3)
			(_nw131).ArraySet1(_dafny.IntOfInt64(-12), 4)
			_819_v13 = _nw131
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_819_v13), 0))
			_ = _index117
			(_819_v13).ArraySet1(((_this).F32()).Minus(_dafny.IntOfInt64(439)), (_index117).Int())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_819_v13), 0))
			_ = _index118
			(_819_v13).ArraySet1((Companion_Default___.Fm7(p0, (_817_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_817_v11), 0))).Int()).(bool), globalState)).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F35, p0), _802_v4)).Cardinality())), (_index118).Int())
		} else {
			(globalState).F28 = (_this).F31()
			var _820_v14 D3
			_ = _820_v14
			_820_v14 = Companion_D3_.Create_DC10_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(455))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}(func(_821_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			})), (_this).F32(), p0)
			r3 = (func() _dafny.Int {
				if _this.F35 {
					return (_820_v14).Dtor_cf14()
				}
				return (_this).F32()
			})()
			var _822_v15 _dafny.CodePoint
			_ = _822_v15
			_822_v15 = _dafny.CodePoint('y')
			(globalState).F28 = (_dafny.MultiSetOf(_822_v15)).IsSubsetOf(_dafny.MultiSetOf(_822_v15, _822_v15))
			_802_v4 = _802_v4
			r3 = _dafny.IntOfUint32((_800_v2).Cardinality())
		}
		var _823_v16 _dafny.Array
		_ = _823_v16
		var _nwElement0_40 bool = !(p0)
		_ = _nwElement0_40
		var _nw132 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(22))
		_ = _nw132
		(_nw132).ArraySet1(_nwElement0_40, 0)
		(_nw132).ArraySet1((_this).F31(), 1)
		(_nw132).ArraySet1(p0, 2)
		(_nw132).ArraySet1(p0, 3)
		(_nw132).ArraySet1(_this.F35, 4)
		(_nw132).ArraySet1(!(p0), 5)
		(_nw132).ArraySet1(p0, 6)
		(_nw132).ArraySet1((_this).F31(), 7)
		(_nw132).ArraySet1(!(false), 8)
		(_nw132).ArraySet1((_this).F31(), 9)
		(_nw132).ArraySet1((_this).F31(), 10)
		(_nw132).ArraySet1(_this.F35, 11)
		(_nw132).ArraySet1(_this.F35, 12)
		(_nw132).ArraySet1(true, 13)
		(_nw132).ArraySet1(!((_this).F31()), 14)
		(_nw132).ArraySet1(Companion_Default___.Fm8(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(79))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}(func(_824_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		}))).Cardinality()), globalState), 15)
		(_nw132).ArraySet1(p0, 16)
		(_nw132).ArraySet1(p0, 17)
		(_nw132).ArraySet1(p0, 18)
		(_nw132).ArraySet1(_this.F35, 19)
		(_nw132).ArraySet1(p0, 20)
		(_nw132).ArraySet1((_this).F31(), 21)
		_823_v16 = _nw132
		var _825_v17 _dafny.Sequence
		_ = _825_v17
		_825_v17 = _dafny.SeqOf(_823_v16, _823_v16)
		(globalState).F20 = (_825_v17).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F32(), (_this).F32())), _dafny.IntOfUint32((_825_v17).Cardinality()))).Uint32()).(_dafny.Array)
		var _826_v18 _dafny.Map
		_ = _826_v18
		_826_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_this).F31())
		var _827_v19 _dafny.Array
		_ = _827_v19
		var _nwElement0_41 _dafny.Map = _826_v18
		_ = _nwElement0_41
		var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.One)
		_ = _nw133
		(_nw133).ArraySet1(_nwElement0_41, 0)
		_827_v19 = _nw133
		var _828_v20 D10
		_ = _828_v20
		_828_v20 = Companion_D10_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), (_this).F31())))
		var _829_v21 T0
		_ = _829_v21
		var _nw134 *C3 = New_C3_()
		_ = _nw134
		_nw134.Ctor__(_827_v19, (_828_v20).Dtor_cf50(), (_this).F31())
		_829_v21 = _nw134
		_829_v21 = _829_v21
		var _830_v22 _dafny.Map
		_ = _830_v22
		_830_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), _802_v4)
		_802_v4 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_830_v22).Contains((_this).F32()) {
				return (_830_v22).Get((_this).F32()).(_dafny.Sequence)
			}
			return _802_v4
		})(), _802_v4)
		(globalState).F16 = (func() bool {
			if (_826_v18).Contains(((_this).F32()).Cmp((_this).F32()) >= 0) {
				return (_826_v18).Get(((_this).F32()).Cmp((_this).F32()) >= 0).(bool)
			}
			return _this.F35
		})()
		r0 = (_this).F32()
		var _831_v23 _dafny.Array
		_ = _831_v23
		var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(8))
		_ = _nw135
		_831_v23 = _nw135
		r1 = _831_v23
		r2 = _dafny.UnicodeSeqOfUtf8Bytes("lsi")
		r3 = (_803_v5).Select((Companion_Default___.SafeIndex((_this).F32(), _dafny.IntOfUint32((_803_v5).Cardinality()))).Uint32()).(_dafny.Int)
		return r0, r1, r2, r3
	}
}
func (_this *C4) F34() _dafny.Sequence {
	{
		return _this._f34
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f29 bool
	_f30 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f29 = false
	_this._f30 = false
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__(f29 bool, f30 bool) {
	{
		(_this)._f29 = f29
		(_this)._f30 = f30
	}
}
func (_this *C5) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 D0, globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		if (_this).F29() {
			var _832_v0 _dafny.Int
			_ = _832_v0
			_832_v0 = _dafny.IntOfInt64(351)
			var _833_v1 _dafny.CodePoint
			_ = _833_v1
			_833_v1 = _dafny.CodePoint('r')
			_832_v0 = (_dafny.Zero).Minus(Companion_Default___.Fm1(_833_v1, globalState))
			var _834_v2 _dafny.Sequence
			_ = _834_v2
			_834_v2 = _dafny.UnicodeSeqOfUtf8Bytes("s")
			var _835_v3 _dafny.Sequence
			_ = _835_v3
			_835_v3 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_834_v2, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_834_v2).Cardinality()))).Uint32(), _833_v1), _834_v2), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(826))).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_834_v2, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_834_v2).Cardinality()))).Uint32(), _833_v1), _834_v2)).Cardinality()))).Uint32(), _833_v1))
			_835_v3 = _dafny.Companion_Sequence_.Concatenate(_835_v3, _835_v3)
			var _836_v4 _dafny.Array
			_ = _836_v4
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_16
			var _nw136 _dafny.Array
			_ = _nw136
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw136 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Int = (func(_837_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_838_i0 _dafny.Int) _dafny.Int {
						return (_838_i0).Plus(_837_v0)
					}
				})(_832_v0)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw136 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw136).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw136).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_836_v4 = _nw136
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_836_v4), 0))
			_ = _index119
			(_836_v4).ArraySet1(_832_v0, (_index119).Int())
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_836_v4), 0))
			_ = _index120
			(_836_v4).ArraySet1(p1, (_index120).Int())
			_832_v0 = (_836_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_836_v4), 0))).Int()).(_dafny.Int)
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_836_v4), 0))
			_ = _index121
			(_836_v4).ArraySet1(p1, (_index121).Int())
		} else {
			var _839_v5 _dafny.MultiSet
			_ = _839_v5
			_839_v5 = _dafny.MultiSetOf(p1)
			Companion_Default___.M0((func() _dafny.Int {
				if (_this).F30() {
					return (_839_v5).Cardinality()
				}
				return p1
			})(), globalState)
			var _840_v6 _dafny.Map
			_ = _840_v6
			_840_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F30())
			if !((func() bool {
				if (_840_v6).Contains(p1) {
					return (_840_v6).Get(p1).(bool)
				}
				return (_this).F29()
			})()) || ((_this).F30()) {
				var _841_v7 _dafny.MultiSet
				_ = _841_v7
				_841_v7 = _dafny.MultiSetOf((_this).F29())
				var _842_v8 _dafny.Sequence
				_ = _842_v8
				_842_v8 = _dafny.SeqOf(_841_v7, _841_v7)
				var _843_v9 _dafny.Map
				_ = _843_v9
				_843_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_842_v8).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm2((_this).F29(), p1, (_this).F30(), (_this).F29(), globalState)).Cardinality()))
				var _844_v10 _dafny.Sequence
				_ = _844_v10
				_844_v10 = _dafny.SeqOf(_843_v9)
				_844_v10 = _844_v10
				(globalState).F28 = (p1).Cmp(p1) == 0
				var _845_v11 _dafny.Int
				_ = _845_v11
				_845_v11 = _dafny.IntOfInt64(497)
				var _846_v12 _dafny.Set
				_ = _846_v12
				_846_v12 = _dafny.SetOf((_this).F30(), (_this).F29())
				var _847_v13 _dafny.Set
				_ = _847_v13
				_847_v13 = _dafny.SetOf(_dafny.IntOfInt64(410), (_846_v12).Cardinality(), _845_v11)
				var _848_v14 _dafny.Sequence
				_ = _848_v14
				_848_v14 = _dafny.SeqOf((_this).F29(), true)
				_845_v11 = ((_847_v13).Cardinality()).Plus(_dafny.IntOfUint32((_848_v14).Cardinality()))
				var _849_v15 _dafny.Map
				_ = _849_v15
				_849_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if Companion_Default___.Fm3(globalState) {
						return true
					}
					return Companion_Default___.Fm3(globalState)
				})(), p2)
				_849_v15 = (_849_v15).Update((_this).F30(), p2)
				var _850_v16 _dafny.Array
				_ = _850_v16
				var _nw137 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw137
				_850_v16 = _nw137
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_850_v16), 0))
				_ = _index122
				(_850_v16).ArraySet1(p1, (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_850_v16), 0))
				_ = _index123
				(_850_v16).ArraySet1((_dafny.Zero).Minus(p1), (_index123).Int())
			} else {
				(globalState).F28 = (p1).Cmp(p1) == 0
				var _851_v17 _dafny.MultiSet
				_ = _851_v17
				_851_v17 = _dafny.MultiSetOf(Companion_Default___.Fm3(globalState), (_this).F30(), (_this).F29(), (_this).F29(), (_this).F29())
				var _852_v18 _dafny.Sequence
				_ = _852_v18
				_852_v18 = _dafny.SeqOf(_851_v17)
				var _853_v19 D3
				_ = _853_v19
				_853_v19 = Companion_D3_.Create_DC9_(_852_v18)
				var _854_v20 D0
				_ = _854_v20
				_854_v20 = Companion_D0_.Create_DC2_(true)
				var _855_v21 _dafny.Array
				_ = _855_v21
				var _nwElement0_42 _dafny.Sequence = _852_v18
				_ = _nwElement0_42
				var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(12))
				_ = _nw138
				(_nw138).ArraySet1(_nwElement0_42, 0)
				(_nw138).ArraySet1(_852_v18, 1)
				(_nw138).ArraySet1(_852_v18, 2)
				(_nw138).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_852_v18, Companion_Default___.Fm4((_this).F30(), _dafny.IntOfInt64(700), globalState)), 3)
				(_nw138).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm4(!((_this).F30()), p1, globalState), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((Companion_Default___.Fm4(!((_this).F30()), p1, globalState)).Cardinality()))).Uint32(), _851_v17), 4)
				(_nw138).ArraySet1(Companion_Default___.Fm4((_this).F30(), p1, globalState), 5)
				(_nw138).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(654))).Uint32(), func(coer52 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_856_v17 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_857_i1 _dafny.Int) _dafny.MultiSet {
						return _856_v17
					}
				})(_851_v17))), _852_v18), 6)
				(_nw138).ArraySet1((_853_v19).Dtor_cf12(), 7)
				(_nw138).ArraySet1(Companion_Default___.Fm4((_854_v20).Dtor_cf4(), _dafny.IntOfInt64(-70), globalState), 8)
				(_nw138).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_851_v17, _851_v17), _852_v18), 9)
				(_nw138).ArraySet1(_852_v18, 10)
				(_nw138).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_852_v18, _852_v18), 11)
				_855_v21 = _nw138
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_855_v21), 0))
				_ = _index124
				(_855_v21).ArraySet1(_dafny.Companion_Sequence_.Update(_852_v18, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_852_v18).Cardinality()))).Uint32(), _851_v17), (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_855_v21), 0))
				_ = _index125
				(_855_v21).ArraySet1(_dafny.SeqOf(_851_v17, _851_v17, (_851_v17).Intersection(_851_v17), _851_v17, _851_v17), (_index125).Int())
				var _858_v22 _dafny.Sequence
				_ = _858_v22
				_858_v22 = _dafny.UnicodeSeqOfUtf8Bytes("fudavpasx")
				var _859_v23 _dafny.Set
				_ = _859_v23
				_859_v23 = _dafny.SetOf((_this).F30())
				var _860_v24 D2
				_ = _860_v24
				_860_v24 = Companion_D2_.Create_DC7_((_858_v22).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_858_v22).Cardinality()))).Uint32()).(_dafny.CodePoint), _859_v23)
				var _861_v25 _dafny.Array
				_ = _861_v25
				var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw139
				_861_v25 = _nw139
				var _rhs124 bool = ((_this).F30()) || ((_this).F29())
				_ = _rhs124
				var _rhs125 D2 = _860_v24
				_ = _rhs125
				var _rhs126 bool = (_dafny.IntOfInt64(-849)).Cmp((p1).Plus((_dafny.Zero).Minus(p1))) < 0
				_ = _rhs126
				var _rhs127 bool = (Companion_Default___.SafeModuloInt(p1, p1)).Cmp(p1) != 0
				_ = _rhs127
				var _rhs128 _dafny.Array = _861_v25
				_ = _rhs128
				var _lhs76 *GlobalState = globalState
				_ = _lhs76
				var _lhs77 *GlobalState = globalState
				_ = _lhs77
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				_lhs76.F28 = _rhs124
				_860_v24 = _rhs125
				_lhs77.F16 = _rhs126
				r0 = _rhs127
				_lhs78.F24 = _rhs128
				var _862_v26 D2
				_ = _862_v26
				_862_v26 = Companion_D2_.Create_DC6_((func() _dafny.Int {
					if (_this).F30() {
						return p1
					}
					return p1
				})())
				_862_v26 = _862_v26
				var _863_v27 _dafny.Int
				_ = _863_v27
				_863_v27 = _dafny.IntOfInt64(-488)
				_863_v27 = _863_v27
			}
			var _864_v28 _dafny.CodePoint
			_ = _864_v28
			_864_v28 = _dafny.CodePoint('d')
			var _865_v29 _dafny.Array
			_ = _865_v29
			var _nw140 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw140
			_865_v29 = _nw140
			var _866_v30 _dafny.Map
			_ = _866_v30
			_866_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_864_v28, _865_v29)
			var _867_v31 _dafny.Sequence
			_ = _867_v31
			_867_v31 = _dafny.SeqOf(_865_v29)
			_866_v30 = (_866_v30).Update(_864_v28, (_867_v31).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_867_v31).Cardinality()))).Uint32()).(_dafny.Array))
			(globalState).F16 = (_this).F29()
			var _868_v32 _dafny.Int
			_ = _868_v32
			_868_v32 = _dafny.IntOfInt64(277)
			_868_v32 = p1
		}
		var _869_v33 _dafny.Array
		_ = _869_v33
		var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
		_ = _nw141
		_869_v33 = _nw141
		for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_869_v33), 0))); ; {
			_guard_loop_4, _ok28 := _iter28()
			if !_ok28 {
				break
			}
			var _870_i2 _dafny.Int
			_870_i2 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_870_i2).Sign() != -1) && ((_870_i2).Cmp(_dafny.ArrayLen((_869_v33), 0)) < 0)) {
				(_869_v33).ArraySet1(_dafny.SeqOf((_this).F30(), (_this).F29(), false, (_dafny.MultiSetOf((_this).F29(), (_this).F29())).IsSubsetOf(_dafny.MultiSetOf((_this).F30())), (_dafny.SetOf(Companion_D2_.Create_DC6_(p1))).IsDisjointFrom(func() _dafny.Set {
					var _coll24 = _dafny.NewBuilder()
					_ = _coll24
					for _iter29 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ho")).Cardinality())), p1)).Keys().Elements()); ; {
						_compr_24, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _871_v34 D2
						_871_v34 = interface{}(_compr_24).(D2)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ho")).Cardinality())), p1)).Contains(_871_v34) {
							_coll24.Add(_871_v34)
						}
					}
					return _coll24.ToSet()
				}())), (_870_i2).Int())
			}
		}
		var _872_v35 _dafny.CodePoint
		_ = _872_v35
		_872_v35 = _dafny.CodePoint('u')
		var _873_v36 _dafny.Set
		_ = _873_v36
		_873_v36 = _dafny.SetOf((_this).F30(), true)
		var _874_v37 D2
		_ = _874_v37
		_874_v37 = Companion_D2_.Create_DC7_(_872_v35, _873_v36)
		var _pat_let_tv19 = _873_v36
		_ = _pat_let_tv19
		_872_v35 = (func(_pat_let22_0 D2) D2 {
			return func(_875_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let23_0 _dafny.Set) D2 {
					return func(_876_dt__update_hcf10_h0 _dafny.Set) D2 {
						return Companion_D2_.Create_DC7_((_875_dt__update__tmp_h0).Dtor_cf9(), _876_dt__update_hcf10_h0)
					}(_pat_let23_0)
				}(_pat_let_tv19)
			}(_pat_let22_0)
		}(_874_v37)).Dtor_cf9()
		var _877_v38 _dafny.Sequence
		_ = _877_v38
		_877_v38 = _dafny.UnicodeSeqOfUtf8Bytes("vjubcpq")
		var _878_v39 D3
		_ = _878_v39
		_878_v39 = Companion_D3_.Create_DC10_(_877_v38, p1, (_this).F30())
		var _879_v40 D3
		_ = _879_v40
		_879_v40 = Companion_D3_.Create_DC11_(_878_v39)
		var _880_v41 D3
		_ = _880_v41
		_880_v41 = Companion_D3_.Create_DC11_(_879_v40)
		var _source13 D3 = _880_v41
		_ = _source13
		if _source13.Is_DC10() {
			var _881___mcc_h0 _dafny.Sequence = _source13.Get_().(D3_DC10).Cf13
			_ = _881___mcc_h0
			var _882___mcc_h1 _dafny.Int = _source13.Get_().(D3_DC10).Cf14
			_ = _882___mcc_h1
			var _883___mcc_h2 bool = _source13.Get_().(D3_DC10).Cf15
			_ = _883___mcc_h2
			var _884_cf15 bool = _883___mcc_h2
			_ = _884_cf15
			var _885_cf14 _dafny.Int = _882___mcc_h1
			_ = _885_cf14
			var _886_cf13 _dafny.Sequence = _881___mcc_h0
			_ = _886_cf13
			_885_cf14 = _885_cf14
			var _887_v42 _dafny.Array
			_ = _887_v42
			var _nw142 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.One)
			_ = _nw142
			_887_v42 = _nw142
			var _888_v43 _dafny.Set
			_ = _888_v43
			_888_v43 = _dafny.SetOf(_887_v42)
			_884_cf15 = (_888_v43).IsDisjointFrom(_888_v43)
			var _889_v44 _dafny.Array
			_ = _889_v44
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_17
			var _nw143 _dafny.Array
			_ = _nw143
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw143 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.Int = (func(_890_cf14 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_891_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_891_i3, _890_cf14)
					}
				})(_885_cf14)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw143 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw143).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw143).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_889_v44 = _nw143
			var _892_v45 _dafny.Array
			_ = _892_v45
			_892_v45 = _889_v44
			var _source14 _dafny.Array = _892_v45
			_ = _source14
			var _893___mcc_h5 _dafny.Array = _source14
			_ = _893___mcc_h5
			var _894_cf7 _dafny.Array = _893___mcc_h5
			_ = _894_cf7
			var _895_v46 _dafny.Array
			_ = _895_v46
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_18
			var _nw144 _dafny.Array
			_ = _nw144
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw144 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = (func(_896_cf15 bool) func(_dafny.Int) bool {
					return func(_897_i4 _dafny.Int) bool {
						return !(!(_896_cf15))
					}
				})(_884_cf15)
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw144 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw144).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw144).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_895_v46 = _nw144
			(globalState).F27 = _895_v46
			var _898_v47 _dafny.Sequence
			_ = _898_v47
			_898_v47 = _dafny.SeqOf(_877_v38)
			var _899_v48 D3
			_ = _899_v48
			_899_v48 = Companion_D3_.Create_DC10_((_898_v47).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_898_v47).Cardinality()))).Uint32()).(_dafny.Sequence), p1, _884_cf15)
			var _900_v49 _dafny.Sequence
			_ = _900_v49
			_900_v49 = _dafny.SeqOf(true)
			(globalState).F16 = ((func() D3 {
				if _884_cf15 {
					return _899_v48
				}
				return Companion_D3_.Create_DC10_(_877_v38, _dafny.IntOfUint32((_900_v49).Cardinality()), _884_cf15)
			})()).Dtor_cf15()
			var _901_v50 _dafny.Map
			_ = _901_v50
			_901_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_877_v38, false)
			var _902_v51 _dafny.MultiSet
			_ = _902_v51
			_902_v51 = _dafny.MultiSetOf((_this).F29())
			var _903_v52 _dafny.Map
			_ = _903_v52
			_903_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_884_cf15, Companion_Default___.Fm0(_902_v51, globalState))
			_901_v50 = (_901_v50).Update((func() _dafny.Sequence {
				if (_903_v52).Contains((_this).F29()) {
					return (_903_v52).Get((_this).F29()).(_dafny.Sequence)
				}
				return _877_v38
			})(), (_this).F29())
			_884_cf15 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(46))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_904_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_905_i5 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus((_dafny.Zero).Minus(_904_p1))
				}
			})(p1))), p0)).Cardinality())).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_877_v38).Cardinality()))) != 0
			var _906_v53 _dafny.MultiSet
			_ = _906_v53
			_906_v53 = _dafny.MultiSetOf(p0)
			var _907_v54 _dafny.Map
			_ = _907_v54
			_907_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5((_this).F29(), _885_cf14, globalState), (_this).F29())
			var _908_v55 _dafny.MultiSet
			_ = _908_v55
			_908_v55 = _dafny.MultiSetOf((_this).F30())
			var _909_v57 _dafny.Set
			_ = _909_v57
			_909_v57 = _dafny.SetOf(_872_v35)
			var _rhs129 bool = (_this).F29()
			_ = _rhs129
			var _rhs130 _dafny.MultiSet = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(p0, p0))
			_ = _rhs130
			var _rhs131 _dafny.Int = (p1).Minus(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_908_v55, (_this).F29())).Cardinality(), (p0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(p1)), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)))
			_ = _rhs131
			var _rhs132 _dafny.Map = func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter30 := _dafny.Iterate((_909_v57).Elements()); ; {
					_compr_25, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _910_v56 _dafny.CodePoint
					_910_v56 = interface{}(_compr_25).(_dafny.CodePoint)
					if (_909_v57).Contains(_910_v56) {
						_coll25.Add(_910_v56, false)
					}
				}
				return _coll25.ToMap()
			}()
			_ = _rhs132
			var _rhs133 bool = (_dafny.IntOfInt64(387)).Cmp(_885_cf14) < 0
			_ = _rhs133
			var _lhs79 *GlobalState = globalState
			_ = _lhs79
			var _lhs80 *GlobalState = globalState
			_ = _lhs80
			_lhs79.F28 = _rhs129
			_906_v53 = _rhs130
			_885_cf14 = _rhs131
			_907_v54 = _rhs132
			_lhs80.F28 = _rhs133
		} else if _source13.Is_DC9() {
			var _911___mcc_h3 _dafny.Sequence = _source13.Get_().(D3_DC9).Cf12
			_ = _911___mcc_h3
			var _912_cf12 _dafny.Sequence = _911___mcc_h3
			_ = _912_cf12
			var _913_v58 _dafny.Map
			_ = _913_v58
			_913_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F30())
			var _914_v59 _dafny.MultiSet
			_ = _914_v59
			_914_v59 = _dafny.MultiSetOf((_this).F30(), false)
			var _915_v60 _dafny.MultiSet
			_ = _915_v60
			_915_v60 = _dafny.MultiSetOf(p1, (_913_v58).Cardinality(), p1, (_914_v59).Cardinality(), _dafny.IntOfInt64(364))
			var _916_v61 _dafny.Map
			_ = _916_v61
			_916_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_915_v60).Cardinality(), (_this).F29())
			var _917_v62 D2
			_ = _917_v62
			_917_v62 = Companion_D2_.Create_DC6_(p1)
			var _918_v63 _dafny.Set
			_ = _918_v63
			_918_v63 = _dafny.SetOf(_917_v62, _917_v62)
			var _919_v64 _dafny.Set
			_ = _919_v64
			_919_v64 = _918_v63
			var _920_v65 _dafny.Sequence
			_ = _920_v65
			_920_v65 = _dafny.SeqOf(_919_v64, _919_v64, _919_v64, _919_v64, _dafny.SetOf(_917_v62))
			var _921_v66 _dafny.Map
			_ = _921_v66
			_921_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_916_v61).Merge(_916_v61)).Cardinality(), _920_v65)
			var _922_v67 _dafny.Map
			_ = _922_v67
			_922_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), _920_v65)).Update(_dafny.IntOfUint32((p0).Cardinality()), _920_v65))
			_921_v66 = (func() _dafny.Map {
				if (_922_v67).Contains((_this).F30()) {
					return (_922_v67).Get((_this).F30()).(_dafny.Map)
				}
				return _921_v66
			})()
			(globalState).F28 = !((_this).F29()) || ((_this).F30())
			r0 = ((_dafny.IntOfInt64(453)).Plus(p1)).Cmp(p1) != 0
			var _rhs134 bool = (_this).F30()
			_ = _rhs134
			var _rhs135 bool = (_this).F29()
			_ = _rhs135
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			_lhs81.F16 = _rhs134
			_lhs82.F28 = _rhs135
		} else {
			var _923___mcc_h4 D3 = _source13.Get_().(D3_DC11).Cf16
			_ = _923___mcc_h4
			var _924_cf16 D3 = _923___mcc_h4
			_ = _924_cf16
			var _925_v68 _dafny.Int
			_ = _925_v68
			_925_v68 = _dafny.IntOfInt64(435)
			var _926_v69 _dafny.MultiSet
			_ = _926_v69
			_926_v69 = _dafny.MultiSetOf((_this).F30())
			var _927_v70 _dafny.Map
			_ = _927_v70
			_927_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_926_v69, (p0).Select((Companion_Default___.SafeIndex(_925_v68, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int))
			_925_v68 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F30() {
					return _925_v68
				}
				return (func() _dafny.Int {
					if (_927_v70).Contains(_926_v69) {
						return (_927_v70).Get(_926_v69).(_dafny.Int)
					}
					return p1
				})()
			})())
			var _928_v71 _dafny.Sequence
			_ = _928_v71
			_928_v71 = _dafny.SeqOf(p2, p2)
			var _929_v72 _dafny.Map
			_ = _929_v72
			_929_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
			var _930_v73 _dafny.Array
			_ = _930_v73
			var _nwElement0_43 bool = (_this).F29()
			_ = _nwElement0_43
			var _nw145 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(22))
			_ = _nw145
			(_nw145).ArraySet1(_nwElement0_43, 0)
			(_nw145).ArraySet1(false, 1)
			(_nw145).ArraySet1(false, 2)
			(_nw145).ArraySet1((_this).F30(), 3)
			(_nw145).ArraySet1((_this).F30(), 4)
			(_nw145).ArraySet1(!(true), 5)
			(_nw145).ArraySet1((_this).F29(), 6)
			(_nw145).ArraySet1((_this).F29(), 7)
			(_nw145).ArraySet1((_this).F30(), 8)
			(_nw145).ArraySet1((_this).F29(), 9)
			(_nw145).ArraySet1((_this).F29(), 10)
			(_nw145).ArraySet1(!((_this).F29()), 11)
			(_nw145).ArraySet1((_this).F30(), 12)
			(_nw145).ArraySet1((func() bool {
				if (_929_v72).Contains(_925_v68) {
					return (_929_v72).Get(_925_v68).(bool)
				}
				return (_this).F30()
			})(), 13)
			(_nw145).ArraySet1((_this).F29(), 14)
			(_nw145).ArraySet1((_this).F29(), 15)
			(_nw145).ArraySet1((func() bool {
				if (_929_v72).Contains(_925_v68) {
					return (_929_v72).Get(_925_v68).(bool)
				}
				return (_this).F29()
			})(), 16)
			(_nw145).ArraySet1((_this).F30(), 17)
			(_nw145).ArraySet1((_this).F29(), 18)
			(_nw145).ArraySet1((_this).F29(), 19)
			(_nw145).ArraySet1(true, 20)
			(_nw145).ArraySet1(true, 21)
			_930_v73 = _nw145
			var _931_v74 _dafny.Sequence
			_ = _931_v74
			_931_v74 = _dafny.SeqOf((_this).F30())
			var _932_v75 _dafny.Map
			_ = _932_v75
			_932_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v74, _930_v73)
			var _933_v76 _dafny.Set
			_ = _933_v76
			_933_v76 = _dafny.SetOf(_925_v68)
			var _934_v77 D5
			_ = _934_v77
			_934_v77 = Companion_D5_.Create_DC14_(_933_v76, p1, (_this).F30(), _930_v73)
			var _935_v78 _dafny.Array
			_ = _935_v78
			var _nwElement0_44 _dafny.Array = _930_v73
			_ = _nwElement0_44
			var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(29))
			_ = _nw146
			(_nw146).ArraySet1(_nwElement0_44, 0)
			(_nw146).ArraySet1(_930_v73, 1)
			(_nw146).ArraySet1(_930_v73, 2)
			(_nw146).ArraySet1(_930_v73, 3)
			(_nw146).ArraySet1(_930_v73, 4)
			(_nw146).ArraySet1(_930_v73, 5)
			(_nw146).ArraySet1(_930_v73, 6)
			(_nw146).ArraySet1(_930_v73, 7)
			(_nw146).ArraySet1(_930_v73, 8)
			(_nw146).ArraySet1(_930_v73, 9)
			(_nw146).ArraySet1(_930_v73, 10)
			(_nw146).ArraySet1(_930_v73, 11)
			(_nw146).ArraySet1(_930_v73, 12)
			(_nw146).ArraySet1(_930_v73, 13)
			(_nw146).ArraySet1(_930_v73, 14)
			(_nw146).ArraySet1((func() _dafny.Array {
				if (_932_v75).Contains(_931_v74) {
					return (_932_v75).Get(_931_v74).(_dafny.Array)
				}
				return _930_v73
			})(), 15)
			(_nw146).ArraySet1(_930_v73, 16)
			(_nw146).ArraySet1((_934_v77).Dtor_cf22(), 17)
			(_nw146).ArraySet1(_930_v73, 18)
			(_nw146).ArraySet1(_930_v73, 19)
			(_nw146).ArraySet1(_930_v73, 20)
			(_nw146).ArraySet1(_930_v73, 21)
			(_nw146).ArraySet1(_930_v73, 22)
			(_nw146).ArraySet1(_930_v73, 23)
			(_nw146).ArraySet1(_930_v73, 24)
			(_nw146).ArraySet1(_930_v73, 25)
			(_nw146).ArraySet1(_930_v73, 26)
			(_nw146).ArraySet1(_930_v73, 27)
			(_nw146).ArraySet1(_930_v73, 28)
			_935_v78 = _nw146
			var _936_v79 _dafny.Map
			_ = _936_v79
			_936_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _925_v68)
			var _937_v81 *C4
			_ = _937_v81
			var _nw147 *C4 = New_C4_()
			_ = _nw147
			_nw147.Ctor__(_928_v71, (_this).F30(), p1, _935_v78, !((func() _dafny.Set {
				var _coll26 = _dafny.NewBuilder()
				_ = _coll26
				for _iter31 := _dafny.Iterate((_dafny.MultiSetOf(_936_v79, _936_v79)).Elements()); ; {
					_compr_26, _ok31 := _iter31()
					if !_ok31 {
						break
					}
					var _938_v80 _dafny.Map
					_938_v80 = interface{}(_compr_26).(_dafny.Map)
					if (_dafny.MultiSetOf(_936_v79, _936_v79)).Contains(_938_v80) {
						_coll26.Add(_938_v80)
					}
				}
				return _coll26.ToSet()
			}()).IsDisjointFrom(_dafny.SetOf(_936_v79))))
			_937_v81 = _nw147
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_930_v73), 0))
			_ = _index126
			(_930_v73).ArraySet1(!((_this).F30()), (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_935_v78), 0))
			_ = _index127
			(_935_v78).ArraySet1(_930_v73, (_index127).Int())
			var _939_v82 _dafny.MultiSet
			_ = _939_v82
			_939_v82 = _dafny.MultiSetOf(_dafny.IntOfInt64(280), _925_v68, _925_v68)
			var _940_v83 _dafny.Array
			_ = _940_v83
			var _nwElement0_45 _dafny.Int = _925_v68
			_ = _nwElement0_45
			var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(5))
			_ = _nw148
			(_nw148).ArraySet1(_nwElement0_45, 0)
			(_nw148).ArraySet1(p1, 1)
			(_nw148).ArraySet1(_dafny.IntOfUint32((_877_v38).Cardinality()), 2)
			(_nw148).ArraySet1(p1, 3)
			(_nw148).ArraySet1(_dafny.IntOfUint32((_877_v38).Cardinality()), 4)
			_940_v83 = _nw148
			var _941_v84 _dafny.Map
			_ = _941_v84
			_941_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_940_v83, _930_v73)
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_930_v73), 0))
			_ = _index128
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_935_v78), 0))
			_ = _index129
			var _rhs136 bool = !_dafny.Companion_Sequence_.Contains(_877_v38, _872_v35)
			_ = _rhs136
			var _rhs137 bool = (_939_v82).IsDisjointFrom(_939_v82)
			_ = _rhs137
			var _rhs138 _dafny.Array = (func() _dafny.Array {
				if (_941_v84).Contains(_940_v83) {
					return (_941_v84).Get(_940_v83).(_dafny.Array)
				}
				return _930_v73
			})()
			_ = _rhs138
			var _rhs139 bool = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfInt64(6))).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_925_v68, _dafny.IntOfInt64(484)))) > 0
			_ = _rhs139
			var _rhs140 _dafny.Sequence = _877_v38
			_ = _rhs140
			var _lhs83 _dafny.Array = _930_v73
			_ = _lhs83
			var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_930_v73), 0))
			_ = _lhs84
			var _lhs85 *GlobalState = globalState
			_ = _lhs85
			var _lhs86 _dafny.Array = _935_v78
			_ = _lhs86
			var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_935_v78), 0))
			_ = _lhs87
			var _lhs88 *C4 = _937_v81
			_ = _lhs88
			(_lhs83).ArraySet1(_rhs136, (_lhs84).Int())
			_lhs85.F16 = _rhs137
			(_lhs86).ArraySet1(_rhs138, (_lhs87).Int())
			_lhs88.F35 = _rhs139
			_877_v38 = _rhs140
			_925_v68 = p1
		}
		var _942_v85 _dafny.Map
		_ = _942_v85
		_942_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.Fm3(globalState))
		var _943_v86 _dafny.Map
		_ = _943_v86
		_943_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F29())
		var _944_i6 _dafny.Int
		_ = _944_i6
		_944_i6 = _dafny.Zero
		{
			for (func() bool {
				if (_942_v85).Contains(p1) {
					return (_942_v85).Get(p1).(bool)
				}
				return (_dafny.IntOfUint32((_877_v38).Cardinality())).Cmp((_943_v86).Cardinality()) <= 0
			})() {
				{
					if (_944_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_944_i6 = (_944_i6).Plus(_dafny.One)
					var _945_v87 _dafny.Array
					_ = _945_v87
					var _nw149 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(7))
					_ = _nw149
					_945_v87 = _nw149
					var _946_v88 _dafny.Array
					_ = _946_v88
					var _len0_19 _dafny.Int = _dafny.IntOfInt64(18)
					_ = _len0_19
					var _nw150 _dafny.Array
					_ = _nw150
					if _len0_19.Cmp(_dafny.Zero) == 0 {
						_nw150 = _dafny.NewArray(_len0_19)
					} else {
						var _init19 func(_dafny.Int) _dafny.Set = (func(_947_p1 _dafny.Int) func(_dafny.Int) _dafny.Set {
							return func(_948_i7 _dafny.Int) _dafny.Set {
								return _dafny.SetOf(_947_p1, _947_p1)
							}
						})(p1)
						_ = _init19
						var _element0_19 = _init19(_dafny.Zero)
						_ = _element0_19
						_nw150 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
						(_nw150).ArraySet1(_element0_19, 0)
						var _nativeLen0_19 = (_len0_19).Int()
						_ = _nativeLen0_19
						for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
							(_nw150).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
						}
					}
					_946_v88 = _nw150
					var _949_v89 _dafny.MultiSet
					_ = _949_v89
					_949_v89 = _dafny.MultiSetOf(p1)
					var _950_v90 D6
					_ = _950_v90
					_950_v90 = Companion_D6_.Create_DC17_((_this).F30(), _946_v88, (_949_v89).Cardinality())
					var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_945_v87), 0))
					_ = _index130
					(_945_v87).ArraySet1(_950_v90, (_index130).Int())
					var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_945_v87), 0))
					_ = _index131
					(_945_v87).ArraySet1(_950_v90, (_index131).Int())
					var _951_v91 D0
					_ = _951_v91
					_951_v91 = Companion_D0_.Create_DC0_(_869_v33)
					var _952_v92 D0
					_ = _952_v92
					_952_v92 = Companion_D0_.Create_DC4_(_951_v91)
					var _953_v93 _dafny.Map
					_ = _953_v93
					_953_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_952_v92, Companion_Default___.Fm1(_872_v35, globalState))
					(globalState).F28 = !(_953_v93).Contains((func() D0 {
						if (_this).F29() {
							return _952_v92
						}
						return _952_v92
					})())
					var _954_v94 _dafny.MultiSet
					_ = _954_v94
					_954_v94 = _dafny.MultiSetOf((_this).F30())
					var _955_v95 _dafny.Sequence
					_ = _955_v95
					_955_v95 = _dafny.SeqOf((_this).F30(), (_this).F29(), (_this).F30(), (_this).F30())
					var _956_v96 _dafny.Sequence
					_ = _956_v96
					_956_v96 = _dafny.SeqOf(_942_v85, (_942_v85).Merge(_942_v85), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F30()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_954_v94).Cardinality(), (_955_v95).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.IntOfUint32((_955_v95).Cardinality()))).Uint32()).(bool)), _942_v85)
					_942_v85 = (_956_v96).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_956_v96).Cardinality()))).Uint32()).(_dafny.Map)
					var _957_v97 _dafny.Array
					_ = _957_v97
					var _nw151 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
					_ = _nw151
					_957_v97 = _nw151
					var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_957_v97), 0))
					_ = _index132
					(_957_v97).ArraySet1(Companion_Default___.Fm10(_872_v35, false, globalState), (_index132).Int())
					var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_957_v97), 0))
					_ = _index133
					(_957_v97).ArraySet1((_this).F29(), (_index133).Int())
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _958_v98 _dafny.Sequence
		_ = _958_v98
		_958_v98 = _dafny.SeqOf(_877_v38, _877_v38)
		var _959_v99 D9
		_ = _959_v99
		_959_v99 = Companion_D9_.Create_DC23_(_958_v98)
		var _source15 D9 = _959_v99
		_ = _source15
		if _source15.Is_DC24() {
			var _960___mcc_h6 _dafny.MultiSet = _source15.Get_().(D9_DC24).Cf40
			_ = _960___mcc_h6
			var _961_cf40 _dafny.MultiSet = _960___mcc_h6
			_ = _961_cf40
			var _962_v100 _dafny.Sequence
			_ = _962_v100
			_962_v100 = _dafny.SeqOf(p2)
			var _963_v101 _dafny.Array
			_ = _963_v101
			var _nwElement0_46 bool = (_this).F29()
			_ = _nwElement0_46
			var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(26))
			_ = _nw152
			(_nw152).ArraySet1(_nwElement0_46, 0)
			(_nw152).ArraySet1((_this).F30(), 1)
			(_nw152).ArraySet1((_this).F29(), 2)
			(_nw152).ArraySet1((_this).F29(), 3)
			(_nw152).ArraySet1((_this).F30(), 4)
			(_nw152).ArraySet1((_this).F30(), 5)
			(_nw152).ArraySet1((_this).F30(), 6)
			(_nw152).ArraySet1((_this).F29(), 7)
			(_nw152).ArraySet1(true, 8)
			(_nw152).ArraySet1((_this).F29(), 9)
			(_nw152).ArraySet1((_this).F29(), 10)
			(_nw152).ArraySet1((_this).F30(), 11)
			(_nw152).ArraySet1((_this).F30(), 12)
			(_nw152).ArraySet1((_this).F29(), 13)
			(_nw152).ArraySet1((_this).F29(), 14)
			(_nw152).ArraySet1((_this).F30(), 15)
			(_nw152).ArraySet1((_this).F30(), 16)
			(_nw152).ArraySet1((_this).F29(), 17)
			(_nw152).ArraySet1((_this).F29(), 18)
			(_nw152).ArraySet1((_this).F30(), 19)
			(_nw152).ArraySet1((_this).F29(), 20)
			(_nw152).ArraySet1((_this).F30(), 21)
			(_nw152).ArraySet1(!((_this).F29()), 22)
			(_nw152).ArraySet1(true, 23)
			(_nw152).ArraySet1(false, 24)
			(_nw152).ArraySet1((_this).F29(), 25)
			_963_v101 = _nw152
			var _964_v102 _dafny.Sequence
			_ = _964_v102
			_964_v102 = _dafny.SeqOf(_963_v101)
			var _965_v103 _dafny.Array
			_ = _965_v103
			var _nwElement0_47 _dafny.Array = _963_v101
			_ = _nwElement0_47
			var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(28))
			_ = _nw153
			(_nw153).ArraySet1(_nwElement0_47, 0)
			(_nw153).ArraySet1(_963_v101, 1)
			(_nw153).ArraySet1(_963_v101, 2)
			(_nw153).ArraySet1(_963_v101, 3)
			(_nw153).ArraySet1(_963_v101, 4)
			(_nw153).ArraySet1(_963_v101, 5)
			(_nw153).ArraySet1(_963_v101, 6)
			(_nw153).ArraySet1((_964_v102).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_964_v102).Cardinality()))).Uint32()).(_dafny.Array), 7)
			(_nw153).ArraySet1(_963_v101, 8)
			(_nw153).ArraySet1(_963_v101, 9)
			(_nw153).ArraySet1(_963_v101, 10)
			(_nw153).ArraySet1(_963_v101, 11)
			(_nw153).ArraySet1(_963_v101, 12)
			(_nw153).ArraySet1(_963_v101, 13)
			(_nw153).ArraySet1(_963_v101, 14)
			(_nw153).ArraySet1(_963_v101, 15)
			(_nw153).ArraySet1(_963_v101, 16)
			(_nw153).ArraySet1(_963_v101, 17)
			(_nw153).ArraySet1(_963_v101, 18)
			(_nw153).ArraySet1(_963_v101, 19)
			(_nw153).ArraySet1(_963_v101, 20)
			(_nw153).ArraySet1(_963_v101, 21)
			(_nw153).ArraySet1((_964_v102).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_964_v102).Cardinality()))).Uint32()).(_dafny.Array), 22)
			(_nw153).ArraySet1(_963_v101, 23)
			(_nw153).ArraySet1(_963_v101, 24)
			(_nw153).ArraySet1(_963_v101, 25)
			(_nw153).ArraySet1(_963_v101, 26)
			(_nw153).ArraySet1(_963_v101, 27)
			_965_v103 = _nw153
			var _966_v104 *C4
			_ = _966_v104
			var _nw154 *C4 = New_C4_()
			_ = _nw154
			_nw154.Ctor__(_962_v100, ((_961_cf40).Cardinality()).Cmp((_dafny.Zero).Minus(p1)) == 0, _dafny.IntOfUint32((_877_v38).Cardinality()), _965_v103, (_this).F30())
			_966_v104 = _nw154
			var _967_v105 _dafny.Int
			_ = _967_v105
			_967_v105 = _dafny.IntOfInt64(570)
			_967_v105 = _967_v105
			var _968_v106 _dafny.Array
			_ = _968_v106
			var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw155
			_968_v106 = _nw155
			var _969_v107 _dafny.Set
			_ = _969_v107
			_969_v107 = _dafny.SetOf(_968_v106)
			_969_v107 = _dafny.SetOf(_968_v106, _968_v106)
			var _970_v108 *C2
			_ = _970_v108
			var _nw156 *C2 = New_C2_()
			_ = _nw156
			_nw156.Ctor__(Companion_Default___.Fm14(p1, _967_v105, _dafny.IntOfInt64(372), (_this).F29(), globalState), (_this).F29())
			_970_v108 = _nw156
		} else if _source15.Is_DC25() {
			var _971___mcc_h7 _dafny.Int = _source15.Get_().(D9_DC25).Cf41
			_ = _971___mcc_h7
			var _972___mcc_h8 bool = _source15.Get_().(D9_DC25).Cf42
			_ = _972___mcc_h8
			var _973___mcc_h9 bool = _source15.Get_().(D9_DC25).Cf43
			_ = _973___mcc_h9
			var _974___mcc_h10 bool = _source15.Get_().(D9_DC25).Cf44
			_ = _974___mcc_h10
			var _975_cf44 bool = _974___mcc_h10
			_ = _975_cf44
			var _976_cf43 bool = _973___mcc_h9
			_ = _976_cf43
			var _977_cf42 bool = _972___mcc_h8
			_ = _977_cf42
			var _978_cf41 _dafny.Int = _971___mcc_h7
			_ = _978_cf41
			var _979_v109 D10
			_ = _979_v109
			_979_v109 = Companion_D10_.Create_DC28_(p1)
			var _source16 D10 = _979_v109
			_ = _source16
			if _source16.Is_DC28() {
				var _980___mcc_h17 _dafny.Int = _source16.Get_().(D10_DC28).Cf51
				_ = _980___mcc_h17
				var _981_cf51 _dafny.Int = _980___mcc_h17
				_ = _981_cf51
				var _982_v110 _dafny.Array
				_ = _982_v110
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw157
				_982_v110 = _nw157
				_982_v110 = _982_v110
				var _983_v111 _dafny.Array
				_ = _983_v111
				var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw158
				_983_v111 = _nw158
				var _984_v112 _dafny.Array
				_ = _984_v112
				_984_v112 = _983_v111
				_984_v112 = _984_v112
				var _985_v113 _dafny.Map
				_ = _985_v113
				_985_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), Companion_Default___.SafeDivisionInt(_981_cf51, _981_cf51))
				_985_v113 = (_985_v113).Update((_this).F29(), _981_cf51)
				_877_v38 = _877_v38
			} else if _source16.Is_DC27() {
				var _986___mcc_h18 _dafny.Map = _source16.Get_().(D10_DC27).Cf50
				_ = _986___mcc_h18
				var _987_cf50 _dafny.Map = _986___mcc_h18
				_ = _987_cf50
				var _988_v114 _dafny.Array
				_ = _988_v114
				var _nwElement0_48 _dafny.Int = _dafny.IntOfInt64(430)
				_ = _nwElement0_48
				var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(3))
				_ = _nw159
				(_nw159).ArraySet1(_nwElement0_48, 0)
				(_nw159).ArraySet1(_dafny.IntOfInt64(-353), 1)
				(_nw159).ArraySet1(_dafny.IntOfInt64(-342), 2)
				_988_v114 = _nw159
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_988_v114), 0))
				_ = _index134
				(_988_v114).ArraySet1((func() _dafny.Int {
					if (_this).F30() {
						return _978_cf41
					}
					return _978_cf41
				})(), (_index134).Int())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_988_v114), 0))
				_ = _index135
				(_988_v114).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_877_v38, _dafny.Companion_Sequence_.Concatenate((_958_v98).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_958_v98).Cardinality()))).Uint32()).(_dafny.Sequence), _877_v38))).Cardinality()), (_index135).Int())
				var _989_v115 _dafny.Sequence
				_ = _989_v115
				_989_v115 = _dafny.SeqOf(true)
				_989_v115 = _989_v115
				(globalState).F16 = _975_cf44
				var _990_v116 _dafny.Map
				_ = _990_v116
				_990_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), p1)
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_988_v114), 0))
				_ = _index136
				(_988_v114).ArraySet1(((func() _dafny.Int {
					if _976_cf43 {
						return (_990_v116).Cardinality()
					}
					return _dafny.IntOfUint32((_877_v38).Cardinality())
				})()).Times((_978_cf41).Times(Companion_Default___.Fm18((func() bool {
					if (_942_v85).Contains(p1) {
						return (_942_v85).Get(p1).(bool)
					}
					return (_this).F30()
				})(), (_this).F29(), globalState))), (_index136).Int())
			} else {
				var _991___mcc_h19 D10 = _source16.Get_().(D10_DC29).Cf52
				_ = _991___mcc_h19
				var _992_cf52 D10 = _991___mcc_h19
				_ = _992_cf52
				_975_cf44 = Companion_Default___.Fm8((_978_cf41).Cmp(_978_cf41) <= 0, _dafny.IntOfInt64(-349), globalState)
				_978_cf41 = _978_cf41
				var _993_v117 _dafny.Sequence
				_ = _993_v117
				_993_v117 = _dafny.SeqOf(_978_cf41)
				_993_v117 = _dafny.SeqOf(_978_cf41, Companion_Default___.Fm7(false, (_this).F30(), globalState), _dafny.IntOfInt64(661))
				var _994_v118 _dafny.Array
				_ = _994_v118
				var _nwElement0_49 bool = !(_977_cf42)
				_ = _nwElement0_49
				var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(2))
				_ = _nw160
				(_nw160).ArraySet1(_nwElement0_49, 0)
				(_nw160).ArraySet1((_this).F30(), 1)
				_994_v118 = _nw160
				var _995_v119 _dafny.Map
				_ = _995_v119
				_995_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _994_v118)
				var _996_v120 D8
				_ = _996_v120
				_996_v120 = Companion_D8_.Create_DC20_(_995_v119)
				var _rhs141 bool = ((_dafny.Zero).Minus(Companion_Default___.Fm1(_872_v35, globalState))).Cmp(_dafny.IntOfInt64(-520)) > 0
				_ = _rhs141
				var _rhs142 D8 = _996_v120
				_ = _rhs142
				var _rhs143 bool = (_this).F30()
				_ = _rhs143
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				_lhs89.F28 = _rhs141
				_996_v120 = _rhs142
				_976_cf43 = _rhs143
			}
			var _997_v122 _dafny.Array
			_ = _997_v122
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_20
			var _nw161 _dafny.Array
			_ = _nw161
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw161 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) bool = (func(_998_cf43 bool) func(_dafny.Int) bool {
					return func(_999_i8 _dafny.Int) bool {
						return _998_cf43
					}
				})(_976_cf43)
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw161 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw161).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw161).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_997_v122 = _nw161
			var _1000_v123 D5
			_ = _1000_v123
			_1000_v123 = Companion_D5_.Create_DC14_(func() _dafny.Set {
				var _coll27 = _dafny.NewBuilder()
				_ = _coll27
				for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-870), _dafny.IntOfInt64(990))); ; {
					_compr_27, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _1001_v121 _dafny.Int
					_1001_v121 = interface{}(_compr_27).(_dafny.Int)
					if ((_dafny.IntOfInt64(-870)).Cmp(_1001_v121) <= 0) && ((_1001_v121).Cmp(_dafny.IntOfInt64(990)) < 0) {
						_coll27.Add(Companion_Default___.SafeModuloInt(_1001_v121, p1))
					}
				}
				return _coll27.ToSet()
			}(), _978_cf41, !(_975_cf44), _997_v122)
			var _1002_v124 _dafny.Map
			_ = _1002_v124
			_1002_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1000_v123, _872_v35)
			var _1003_v125 _dafny.Set
			_ = _1003_v125
			_1003_v125 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ouwdevy")).Cardinality()))
			var _rhs144 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_1002_v124).Contains(Companion_D5_.Create_DC14_(_1003_v125, p1, !((_this).F29()), _997_v122)) {
					return (_1002_v124).Get(Companion_D5_.Create_DC14_(_1003_v125, p1, !((_this).F29()), _997_v122)).(_dafny.CodePoint)
				}
				return _872_v35
			})()
			_ = _rhs144
			var _rhs145 _dafny.Int = (_978_cf41).Minus(Companion_Default___.Fm1(_872_v35, globalState))
			_ = _rhs145
			var _rhs146 bool = Companion_Default___.Fm3(globalState)
			_ = _rhs146
			var _rhs147 bool = true
			_ = _rhs147
			_872_v35 = _rhs144
			_978_cf41 = _rhs145
			_976_cf43 = _rhs146
			_977_cf42 = _rhs147
			var _1004_v126 _dafny.Map
			_ = _1004_v126
			_1004_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_978_cf41, _dafny.SetOf(_978_cf41, _978_cf41))
			var _1005_v130 _dafny.Array
			_ = _1005_v130
			var _nwElement0_50 _dafny.Set = _1003_v125
			_ = _nwElement0_50
			var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(29))
			_ = _nw162
			(_nw162).ArraySet1(_nwElement0_50, 0)
			(_nw162).ArraySet1(_1003_v125, 1)
			(_nw162).ArraySet1(_1003_v125, 2)
			(_nw162).ArraySet1(_1003_v125, 3)
			(_nw162).ArraySet1(_1003_v125, 4)
			(_nw162).ArraySet1(_dafny.SetOf(p1, p1), 5)
			(_nw162).ArraySet1(_1003_v125, 6)
			(_nw162).ArraySet1(_dafny.SetOf(_978_cf41, p1, _978_cf41), 7)
			(_nw162).ArraySet1(_1003_v125, 8)
			(_nw162).ArraySet1(_1003_v125, 9)
			(_nw162).ArraySet1(_1003_v125, 10)
			(_nw162).ArraySet1(_1003_v125, 11)
			(_nw162).ArraySet1(_dafny.SetOf(_978_cf41, p1), 12)
			(_nw162).ArraySet1(_1003_v125, 13)
			(_nw162).ArraySet1((_1000_v123).Dtor_cf19(), 14)
			(_nw162).ArraySet1(_1003_v125, 15)
			(_nw162).ArraySet1(_1003_v125, 16)
			(_nw162).ArraySet1(_dafny.SetOf(p1, (_dafny.Zero).Minus(_978_cf41)), 17)
			(_nw162).ArraySet1(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-73))), 18)
			(_nw162).ArraySet1((func() _dafny.Set {
				if (_1004_v126).Contains(_978_cf41) {
					return (_1004_v126).Get(_978_cf41).(_dafny.Set)
				}
				return func() _dafny.Set {
					var _coll28 = _dafny.NewBuilder()
					_ = _coll28
					for _iter33 := _dafny.Iterate((func() _dafny.Set {
						var _coll29 = _dafny.NewBuilder()
						_ = _coll29
						for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(464), _dafny.IntOfInt64(420))); ; {
							_compr_29, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _1006_v127 _dafny.Int
							_1006_v127 = interface{}(_compr_29).(_dafny.Int)
							if ((_dafny.IntOfInt64(464)).Cmp(_1006_v127) <= 0) && ((_1006_v127).Cmp(_dafny.IntOfInt64(420)) < 0) {
								_coll29.Add(Companion_Default___.SafeDivisionInt(_1006_v127, p1))
							}
						}
						return _coll29.ToSet()
					}()).Elements()); ; {
						_compr_28, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _1007_v128 _dafny.Int
						_1007_v128 = interface{}(_compr_28).(_dafny.Int)
						if (func() _dafny.Set {
							var _coll30 = _dafny.NewBuilder()
							_ = _coll30
							for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(464), _dafny.IntOfInt64(420))); ; {
								_compr_30, _ok35 := _iter35()
								if !_ok35 {
									break
								}
								var _1008_v127 _dafny.Int
								_1008_v127 = interface{}(_compr_30).(_dafny.Int)
								if ((_dafny.IntOfInt64(464)).Cmp(_1008_v127) <= 0) && ((_1008_v127).Cmp(_dafny.IntOfInt64(420)) < 0) {
									_coll30.Add(Companion_Default___.SafeDivisionInt(_1008_v127, p1))
								}
							}
							return _coll30.ToSet()
						}()).Contains(_1007_v128) {
							_coll28.Add(Companion_Default___.SafeModuloInt(_1007_v128, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality())))
						}
					}
					return _coll28.ToSet()
				}()
			})(), 19)
			(_nw162).ArraySet1(_1003_v125, 20)
			(_nw162).ArraySet1(_dafny.SetOf(_978_cf41), 21)
			(_nw162).ArraySet1(_dafny.SetOf(p1), 22)
			(_nw162).ArraySet1(_1003_v125, 23)
			(_nw162).ArraySet1(_1003_v125, 24)
			(_nw162).ArraySet1(Companion_Default___.Fm28((_943_v86).Cardinality(), false, _975_cf44, globalState), 25)
			(_nw162).ArraySet1(_1003_v125, 26)
			(_nw162).ArraySet1(_1003_v125, 27)
			(_nw162).ArraySet1(func() _dafny.Set {
				var _coll31 = _dafny.NewBuilder()
				_ = _coll31
				for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(160), _dafny.IntOfInt64(-306))); ; {
					_compr_31, _ok36 := _iter36()
					if !_ok36 {
						break
					}
					var _1009_v129 _dafny.Int
					_1009_v129 = interface{}(_compr_31).(_dafny.Int)
					if ((_dafny.IntOfInt64(160)).Cmp(_1009_v129) <= 0) && ((_1009_v129).Cmp(_dafny.IntOfInt64(-306)) < 0) {
						_coll31.Add((_1009_v129).Times(_978_cf41))
					}
				}
				return _coll31.ToSet()
			}(), 28)
			_1005_v130 = _nw162
			var _1010_v131 D6
			_ = _1010_v131
			_1010_v131 = Companion_D6_.Create_DC17_(_975_cf44, _1005_v130, p1)
			_975_cf44 = (_1010_v131).Dtor_cf25()
			var _1011_v132 _dafny.Map
			_ = _1011_v132
			_1011_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_872_v35, (p0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int))
			var _1012_v133 _dafny.MultiSet
			_ = _1012_v133
			_1012_v133 = _dafny.MultiSetOf(p1)
			_1011_v132 = (_1011_v132).Update(Companion_Default___.Fm20(_877_v38, _1012_v133, false, _dafny.IntOfInt64(-38), globalState), Companion_Default___.SafeModuloInt(_978_cf41, _dafny.IntOfInt64(150)))
		} else if _source15.Is_DC26() {
			var _1013___mcc_h11 bool = _source15.Get_().(D9_DC26).Cf45
			_ = _1013___mcc_h11
			var _1014___mcc_h12 _dafny.Sequence = _source15.Get_().(D9_DC26).Cf46
			_ = _1014___mcc_h12
			var _1015___mcc_h13 _dafny.Int = _source15.Get_().(D9_DC26).Cf47
			_ = _1015___mcc_h13
			var _1016___mcc_h14 _dafny.Map = _source15.Get_().(D9_DC26).Cf48
			_ = _1016___mcc_h14
			var _1017___mcc_h15 bool = _source15.Get_().(D9_DC26).Cf49
			_ = _1017___mcc_h15
			var _1018_cf49 bool = _1017___mcc_h15
			_ = _1018_cf49
			var _1019_cf48 _dafny.Map = _1016___mcc_h14
			_ = _1019_cf48
			var _1020_cf47 _dafny.Int = _1015___mcc_h13
			_ = _1020_cf47
			var _1021_cf46 _dafny.Sequence = _1014___mcc_h12
			_ = _1021_cf46
			var _1022_cf45 bool = _1013___mcc_h11
			_ = _1022_cf45
			if ((_873_v36).Union(_873_v36)).Equals(_dafny.SetOf(_1018_cf49)) {
				var _1023_v134 _dafny.MultiSet
				_ = _1023_v134
				_1023_v134 = _dafny.MultiSetOf(!(_1018_cf49))
				var _1024_v135 _dafny.Array
				_ = _1024_v135
				var _nwElement0_51 bool = true
				_ = _nwElement0_51
				var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(17))
				_ = _nw163
				(_nw163).ArraySet1(_nwElement0_51, 0)
				(_nw163).ArraySet1(_1018_cf49, 1)
				(_nw163).ArraySet1(true, 2)
				(_nw163).ArraySet1((func() bool {
					if _1018_cf49 {
						return (_this).F29()
					}
					return (_this).F30()
				})(), 3)
				(_nw163).ArraySet1((_1022_cf45) == ((func() bool {
					if (_943_v86).Contains((_this).F30()) {
						return (_943_v86).Get((_this).F30()).(bool)
					}
					return _1022_cf45
				})()), 4)
				(_nw163).ArraySet1((_this).F30(), 5)
				(_nw163).ArraySet1(_1018_cf49, 6)
				(_nw163).ArraySet1((_this).F30(), 7)
				(_nw163).ArraySet1(_1022_cf45, 8)
				(_nw163).ArraySet1(_1022_cf45, 9)
				(_nw163).ArraySet1((_this).F29(), 10)
				(_nw163).ArraySet1((_dafny.MultiSetOf(_1020_cf47, p1)).IsSubsetOf(_dafny.MultiSetOf(_1020_cf47)), 11)
				(_nw163).ArraySet1((func() bool {
					if (_this).F29() {
						return !((_this).F29())
					}
					return _1022_cf45
				})(), 12)
				(_nw163).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1021_cf46, (Companion_Default___.SafeIndex((_1023_v134).Cardinality(), _dafny.IntOfUint32((_1021_cf46).Cardinality()))).Uint32(), _872_v35)).Cardinality())).Cmp(p1) >= 0, 13)
				(_nw163).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("vu"), _877_v38), 14)
				(_nw163).ArraySet1(_1018_cf49, 15)
				(_nw163).ArraySet1(false, 16)
				_1024_v135 = _nw163
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1024_v135), 0))
				_ = _index137
				(_1024_v135).ArraySet1(_1018_cf49, (_index137).Int())
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1024_v135), 0))
				_ = _index138
				(_1024_v135).ArraySet1((func() bool {
					if _1022_cf45 {
						return (_1020_cf47).Cmp(_1020_cf47) >= 0
					}
					return _1018_cf49
				})(), (_index138).Int())
				var _1025_v136 _dafny.Array
				_ = _1025_v136
				var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
				_ = _nw164
				_1025_v136 = _nw164
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1025_v136), 0))
				_ = _index139
				(_1025_v136).ArraySet1CodePoint(_872_v35, (_index139).Int())
				var _1026_v137 _dafny.Array
				_ = _1026_v137
				var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw165
				_1026_v137 = _nw165
				var _1027_v138 *C1
				_ = _1027_v138
				var _nw166 *C1 = New_C1_()
				_ = _nw166
				_nw166.Ctor__((Companion_Default___.Fm28(p1, (_this).F29(), _1018_cf49, globalState)).Cardinality(), _1026_v137, _1022_cf45)
				_1027_v138 = _nw166
				var _1028_v139 *C1
				_ = _1028_v139
				_1028_v139 = _1027_v138
				var _1029_v140 _dafny.Array
				_ = _1029_v140
				var _nwElement0_52 *C1 = _1027_v138
				_ = _nwElement0_52
				var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(13))
				_ = _nw167
				(_nw167).ArraySet1(_nwElement0_52, 0)
				(_nw167).ArraySet1(_1027_v138, 1)
				(_nw167).ArraySet1(_1027_v138, 2)
				(_nw167).ArraySet1((_1028_v139), 3)
				(_nw167).ArraySet1(_1027_v138, 4)
				(_nw167).ArraySet1(_1027_v138, 5)
				(_nw167).ArraySet1(_1027_v138, 6)
				(_nw167).ArraySet1(_1027_v138, 7)
				(_nw167).ArraySet1(_1027_v138, 8)
				(_nw167).ArraySet1(_1027_v138, 9)
				(_nw167).ArraySet1(_1027_v138, 10)
				(_nw167).ArraySet1((_1028_v139), 11)
				(_nw167).ArraySet1(_1027_v138, 12)
				_1029_v140 = _nw167
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_1029_v140), 0))
				_ = _index140
				(_1029_v140).ArraySet1(_1027_v138, (_index140).Int())
				var _1030_v141 _dafny.Array
				_ = _1030_v141
				var _nwElement0_53 _dafny.Map = (_943_v86).Update(_1022_cf45, _1018_cf49)
				_ = _nwElement0_53
				var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(13))
				_ = _nw168
				(_nw168).ArraySet1(_nwElement0_53, 0)
				(_nw168).ArraySet1(_943_v86, 1)
				(_nw168).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1018_cf49, _1022_cf45), 2)
				(_nw168).ArraySet1(_943_v86, 3)
				(_nw168).ArraySet1(_943_v86, 4)
				(_nw168).ArraySet1(_943_v86, 5)
				(_nw168).ArraySet1(_943_v86, 6)
				(_nw168).ArraySet1(_943_v86, 7)
				(_nw168).ArraySet1(_943_v86, 8)
				(_nw168).ArraySet1(_943_v86, 9)
				(_nw168).ArraySet1(_943_v86, 10)
				(_nw168).ArraySet1(_943_v86, 11)
				(_nw168).ArraySet1(_943_v86, 12)
				_1030_v141 = _nw168
				var _1031_v142 _dafny.Map
				_ = _1031_v142
				_1031_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1020_cf47), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1018_cf49))
				var _1032_v143 *C3
				_ = _1032_v143
				var _nw169 *C3 = New_C3_()
				_ = _nw169
				_nw169.Ctor__(_1030_v141, _1031_v142, _1022_cf45)
				_1032_v143 = _nw169
				var _1033_v144 _dafny.Set
				_ = _1033_v144
				_1033_v144 = _dafny.SetOf(_1032_v143, _1032_v143)
				var _1034_v145 _dafny.Map
				_ = _1034_v145
				_1034_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_943_v86, _1033_v144)
				var _1035_v146 _dafny.Array
				_ = _1035_v146
				var _nwElement0_54 _dafny.Set = _1033_v144
				_ = _nwElement0_54
				var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(4))
				_ = _nw170
				(_nw170).ArraySet1(_nwElement0_54, 0)
				(_nw170).ArraySet1(_1033_v144, 1)
				(_nw170).ArraySet1((func() _dafny.Set {
					if (_1034_v145).Contains(_943_v86) {
						return (_1034_v145).Get(_943_v86).(_dafny.Set)
					}
					return _dafny.SetOf(_1032_v143, _1032_v143, _1032_v143, _1032_v143, _1032_v143)
				})(), 2)
				(_nw170).ArraySet1(_dafny.SetOf(_1032_v143), 3)
				_1035_v146 = _nw170
				var _1036_v147 _dafny.Map
				_ = _1036_v147
				_1036_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1035_v146, (_this).F29())
				var _1037_v148 _dafny.Map
				_ = _1037_v148
				_1037_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm29(_1020_cf47, _1018_cf49, globalState), (p0).Select((Companion_Default___.SafeIndex(_1020_cf47, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int))
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1024_v135), 0))
				_ = _index141
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1025_v136), 0))
				_ = _index142
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_1029_v140), 0))
				_ = _index143
				var _rhs148 bool = (func() bool {
					if (_1036_v147).Contains(_1035_v146) {
						return (_1036_v147).Get(_1035_v146).(bool)
					}
					return (_dafny.IntOfUint32((p0).Cardinality())).Cmp(p1) != 0
				})()
				_ = _rhs148
				var _rhs149 _dafny.Int = (func() _dafny.Int {
					if (_1037_v148).Contains(p2) {
						return (_1037_v148).Get(p2).(_dafny.Int)
					}
					return _1020_cf47
				})()
				_ = _rhs149
				var _rhs150 _dafny.Sequence = _1021_cf46
				_ = _rhs150
				var _rhs151 _dafny.CodePoint = _872_v35
				_ = _rhs151
				var _rhs152 *C1 = (func() *C1 {
					if _1018_cf49 {
						return _1027_v138
					}
					return (func() *C1 {
						if (_this).F29() {
							return _1027_v138
						}
						return _1027_v138
					})()
				})()
				_ = _rhs152
				var _lhs90 _dafny.Array = _1024_v135
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1024_v135), 0))
				_ = _lhs91
				var _lhs92 _dafny.Array = _1025_v136
				_ = _lhs92
				var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1025_v136), 0))
				_ = _lhs93
				var _lhs94 _dafny.Array = _1029_v140
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_1029_v140), 0))
				_ = _lhs95
				(_lhs90).ArraySet1(_rhs148, (_lhs91).Int())
				_1020_cf47 = _rhs149
				_877_v38 = _rhs150
				(_lhs92).ArraySet1CodePoint(_rhs151, (_lhs93).Int())
				(_lhs94).ArraySet1(_rhs152, (_lhs95).Int())
				var _1038_v149 D3
				_ = _1038_v149
				_1038_v149 = Companion_D3_.Create_DC10_(_877_v38, _1020_cf47, _1018_cf49)
				var _1039_v150 _dafny.Array
				_ = _1039_v150
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_21
				var _nw171 _dafny.Array
				_ = _nw171
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw171 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Int = (func(_1040_cf47 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1041_i9 _dafny.Int) _dafny.Int {
							return (_1041_i9).Plus(_1040_cf47)
						}
					})(_1020_cf47)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw171 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw171).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw171).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_1039_v150 = _nw171
				var _rhs153 D3 = _1038_v149
				_ = _rhs153
				var _rhs154 _dafny.Array = _1039_v150
				_ = _rhs154
				var _rhs155 _dafny.Int = p1
				_ = _rhs155
				var _rhs156 bool = (_this).F30()
				_ = _rhs156
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				_1038_v149 = _rhs153
				_lhs96.F24 = _rhs154
				_1020_cf47 = _rhs155
				_1022_cf45 = _rhs156
				var _1042_v151 *C2
				_ = _1042_v151
				var _nw172 *C2 = New_C2_()
				_ = _nw172
				_nw172.Ctor__(_dafny.SeqOf(true), _1018_cf49)
				_1042_v151 = _nw172
				var _1043_v152 _dafny.Map
				_ = _1043_v152
				_1043_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1020_cf47, _1042_v151)
				var _rhs157 bool = (_this).F29()
				_ = _rhs157
				var _rhs158 *C2 = (func() *C2 {
					if (_1043_v152).Contains((_1032_v143).Fm9(_1020_cf47, p0, p1, (_this).F30(), globalState)) {
						return (_1043_v152).Get((_1032_v143).Fm9(_1020_cf47, p0, p1, (_this).F30(), globalState)).(*C2)
					}
					return _1042_v151
				})()
				_ = _rhs158
				_1022_cf45 = _rhs157
				_1042_v151 = _rhs158
				var _1044_v153 _dafny.Set
				_ = _1044_v153
				_1044_v153 = _dafny.SetOf(_1020_cf47)
				var _1045_v154 _dafny.Map
				_ = _1045_v154
				_1045_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(_dafny.IntOfUint32((p0).Cardinality())))
				var _1046_v155 _dafny.Map
				_ = _1046_v155
				_1046_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1044_v153).Cardinality(), (func() _dafny.Set {
					if (_1045_v154).Contains(_1022_cf45) {
						return (_1045_v154).Get(_1022_cf45).(_dafny.Set)
					}
					return _1044_v153
				})())
				var _1047_v156 _dafny.Map
				_ = _1047_v156
				_1047_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_877_v38, (_1023_v134).Cardinality())
				var _1048_v157 _dafny.Map
				_ = _1048_v157
				_1048_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(607), (_1047_v156).Cardinality()), _1046_v155)
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1024_v135), 0))
				_ = _index144
				var _rhs159 bool = !(!_dafny.Companion_Sequence_.Contains((_1042_v151).F38(), (_1018_cf49) || ((_1024_v135).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1024_v135), 0))).Int()).(bool))))
				_ = _rhs159
				var _rhs160 _dafny.Map = (func() _dafny.Map {
					if (_1048_v157).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1021_cf46, _1021_cf46)).Cardinality())) {
						return (_1048_v157).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1021_cf46, _1021_cf46)).Cardinality())).(_dafny.Map)
					}
					return _1046_v155
				})()
				_ = _rhs160
				var _rhs161 _dafny.MultiSet = (_1023_v134).Update(!(_1022_cf45), Companion_Default___.Abs(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus(p1)), p1)))
				_ = _rhs161
				var _rhs162 _dafny.Array = _1024_v135
				_ = _rhs162
				var _lhs97 _dafny.Array = _1024_v135
				_ = _lhs97
				var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1024_v135), 0))
				_ = _lhs98
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				(_lhs97).ArraySet1(_rhs159, (_lhs98).Int())
				_1046_v155 = _rhs160
				_1023_v134 = _rhs161
				_lhs99.F20 = _rhs162
			} else {
				var _1049_v158 _dafny.Array
				_ = _1049_v158
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_22
				var _nw173 _dafny.Array
				_ = _nw173
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw173 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) bool = (func(_1050_cf45 bool) func(_dafny.Int) bool {
						return func(_1051_i10 _dafny.Int) bool {
							return _1050_cf45
						}
					})(_1022_cf45)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw173 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw173).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw173).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_1049_v158 = _nw173
				(globalState).F20 = _1049_v158
				Companion_Default___.M0(Companion_Default___.Fm1(_872_v35, globalState), globalState)
				var _1052_v159 _dafny.Sequence
				_ = _1052_v159
				_1052_v159 = _dafny.SeqOf((_this).F29(), (_this).F29(), Companion_Default___.Fm26((_this).F29(), (_this).F29(), (_this).F30(), p1, globalState))
				var _1053_v160 _dafny.Map
				_ = _1053_v160
				_1053_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1052_v159, p1)
				_1053_v160 = _1053_v160
				var _1054_v161 _dafny.Sequence
				_ = _1054_v161
				_1054_v161 = _dafny.SeqOf(_1049_v158, _1049_v158, _1049_v158)
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1049_v158), 0))
				_ = _index145
				(_1049_v158).ArraySet1((_this).F30(), (_index145).Int())
				var _1055_v162 _dafny.Array
				_ = _1055_v162
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(23))
				_ = _nw174
				_1055_v162 = _nw174
				var _1056_v163 _dafny.Set
				_ = _1056_v163
				_1056_v163 = _dafny.SetOf(_872_v35)
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_1055_v162), 0))
				_ = _index146
				(_1055_v162).ArraySet1((_1056_v163).Union(_1056_v163), (_index146).Int())
				var _1057_v164 _dafny.Sequence
				_ = _1057_v164
				_1057_v164 = _dafny.SeqOf(_dafny.SetOf(_dafny.CodePoint('k'), _872_v35, _872_v35))
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1049_v158), 0))
				_ = _index147
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_1055_v162), 0))
				_ = _index148
				var _rhs163 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1052_v159, (func() _dafny.Sequence {
					if (_this).F29() {
						return _1052_v159
					}
					return _1052_v159
				})()), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1052_v159, (func() _dafny.Sequence {
					if (_this).F29() {
						return _1052_v159
					}
					return _1052_v159
				})())).Cardinality()))).Uint32(), !((_this).F30()) || (_1018_cf49))).Cardinality())
				_ = _rhs163
				var _rhs164 _dafny.Int = (_dafny.IntOfInt64(-756)).Plus(p1)
				_ = _rhs164
				var _rhs165 _dafny.Sequence = _1054_v161
				_ = _rhs165
				var _rhs166 bool = (_this).F29()
				_ = _rhs166
				var _rhs167 _dafny.Set = (_1057_v164).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_1057_v164).Cardinality()))).Uint32()).(_dafny.Set)
				_ = _rhs167
				var _lhs100 _dafny.Array = _1049_v158
				_ = _lhs100
				var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1049_v158), 0))
				_ = _lhs101
				var _lhs102 _dafny.Array = _1055_v162
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_1055_v162), 0))
				_ = _lhs103
				_1020_cf47 = _rhs163
				_1020_cf47 = _rhs164
				_1054_v161 = _rhs165
				(_lhs100).ArraySet1(_rhs166, (_lhs101).Int())
				(_lhs102).ArraySet1(_rhs167, (_lhs103).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1049_v158), 0))
				_ = _index149
				var _rhs168 bool = (func() bool {
					if (func() bool {
						if (_943_v86).Contains((_this).F29()) {
							return (_943_v86).Get((_this).F29()).(bool)
						}
						return (_1049_v158).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1049_v158), 0))).Int()).(bool)
					})() {
						return true
					}
					return (_this).F29()
				})()
				_ = _rhs168
				var _rhs169 _dafny.Int = (_1020_cf47).Plus((_dafny.Zero).Minus((_1020_cf47).Minus(_1020_cf47)))
				_ = _rhs169
				var _lhs104 _dafny.Array = _1049_v158
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1049_v158), 0))
				_ = _lhs105
				(_lhs104).ArraySet1(_rhs168, (_lhs105).Int())
				_1020_cf47 = _rhs169
			}
			var _1058_v165 _dafny.Array
			_ = _1058_v165
			var _nw175 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw175
			_1058_v165 = _nw175
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1058_v165), 0))
			_ = _index150
			(_1058_v165).ArraySet1(_1018_cf49, (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1058_v165), 0))
			_ = _index151
			(_1058_v165).ArraySet1((_this).F29(), (_index151).Int())
			var _1059_v166 _dafny.Sequence
			_ = _1059_v166
			_1059_v166 = _dafny.SeqOf((_this).F30())
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1058_v165), 0))
			_ = _index152
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1058_v165), 0))
			_ = _index153
			var _rhs170 bool = !(_873_v36).Contains((_1059_v166).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1059_v166).Cardinality()))).Uint32()).(bool))
			_ = _rhs170
			var _rhs171 bool = (_1058_v165).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1058_v165), 0))).Int()).(bool)
			_ = _rhs171
			var _rhs172 _dafny.Sequence = _1059_v166
			_ = _rhs172
			var _rhs173 bool = Companion_Default___.Fm3(globalState)
			_ = _rhs173
			var _lhs106 _dafny.Array = _1058_v165
			_ = _lhs106
			var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1058_v165), 0))
			_ = _lhs107
			var _lhs108 _dafny.Array = _1058_v165
			_ = _lhs108
			var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1058_v165), 0))
			_ = _lhs109
			var _lhs110 *GlobalState = globalState
			_ = _lhs110
			(_lhs106).ArraySet1(_rhs170, (_lhs107).Int())
			(_lhs108).ArraySet1(_rhs171, (_lhs109).Int())
			_1059_v166 = _rhs172
			_lhs110.F16 = _rhs173
			var _1060_v167 _dafny.Map
			_ = _1060_v167
			_1060_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1018_cf49, _872_v35)
			var _1061_v168 _dafny.Map
			_ = _1061_v168
			_1061_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), _872_v35)
			var _1062_v169 _dafny.Array
			_ = _1062_v169
			var _nwElement0_55 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_1060_v167).Contains(_1022_cf45) {
					return (_1060_v167).Get(_1022_cf45).(_dafny.CodePoint)
				}
				return _872_v35
			})()
			_ = _nwElement0_55
			var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(7))
			_ = _nw176
			(_nw176).ArraySet1CodePoint(_nwElement0_55, 0)
			(_nw176).ArraySet1CodePoint(_872_v35, 1)
			(_nw176).ArraySet1CodePoint(_dafny.CodePoint('k'), 2)
			(_nw176).ArraySet1CodePoint(_872_v35, 3)
			(_nw176).ArraySet1CodePoint(_872_v35, 4)
			(_nw176).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_1061_v168).Contains(_1020_cf47) {
					return (_1061_v168).Get(_1020_cf47).(_dafny.CodePoint)
				}
				return _872_v35
			})(), 5)
			(_nw176).ArraySet1CodePoint(_872_v35, 6)
			_1062_v169 = _nw176
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1062_v169), 0))
			_ = _index154
			(_1062_v169).ArraySet1CodePoint(_872_v35, (_index154).Int())
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1062_v169), 0))
			_ = _index155
			(_1062_v169).ArraySet1CodePoint(_872_v35, (_index155).Int())
		} else {
			var _1063___mcc_h16 _dafny.Sequence = _source15.Get_().(D9_DC23).Cf39
			_ = _1063___mcc_h16
			var _1064_cf39 _dafny.Sequence = _1063___mcc_h16
			_ = _1064_cf39
			var _1065_v170 _dafny.Int
			_ = _1065_v170
			_1065_v170 = _dafny.IntOfInt64(18)
			var _1066_v171 _dafny.Set
			_ = _1066_v171
			_1066_v171 = _dafny.SetOf(p1)
			_1065_v170 = (_1066_v171).Cardinality()
			_877_v38 = _877_v38
			var _1067_v172 _dafny.Array
			_ = _1067_v172
			var _nw177 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw177
			_1067_v172 = _nw177
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_1067_v172), 0))
			_ = _index156
			(_1067_v172).ArraySet1(!(_873_v36).Contains((_this).F29()), (_index156).Int())
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_1067_v172), 0))
			_ = _index157
			(_1067_v172).ArraySet1((_this).F30(), (_index157).Int())
			(globalState).F16 = (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-612), _1065_v170)).Cmp(_1065_v170) >= 0
		}
		r0 = (_this).F29()
		r1 = !((Companion_Default___.SafeDivisionInt(p1, p1)).Cmp(p1) != 0)
		return r0, r1
	}
}
func (_this *C5) F29() bool {
	{
		return _this._f29
	}
}
func (_this *C5) F30() bool {
	{
		return _this._f30
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
