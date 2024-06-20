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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) _dafny.Int {
	if false {
		return _dafny.IntOfInt64(34)
	} else {
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(419), _dafny.IntOfUint32((_dafny.SeqOf(true, !(false), false, false, true)).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(447), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(881), _dafny.IntOfInt64(850))).Merge(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-920), _dafny.IntOfInt64(636))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-920)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(636)) < 0) {
				_coll0.Add((_0_v0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ff")).Cardinality())), _dafny.IntOfInt64(96))
			}
		}
		return _coll0.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D3 = Companion_D3_.Create_DC11_()
	_ = _source0
	if _source0.Is_DC11() {
		if true {
			return _dafny.CodePoint('h')
		} else {
			return _dafny.CodePoint('v')
		}
	} else {
		var _1___mcc_h0 _dafny.Sequence = _source0.Get_().(D3_DC10).Cf18
		_ = _1___mcc_h0
		var _2_cf18 _dafny.Sequence = _1___mcc_h0
		_ = _2_cf18
		return _dafny.CodePoint('t')
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	return (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-89))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	}))).Cardinality()), _dafny.IntOfInt64(-546))).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-556), _dafny.IntOfInt64(756))) < 0
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(733))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality()), (_dafny.MultiSetOf(!(false))).Cardinality()), false)
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(true, (_dafny.IntOfInt64(-765)).Cmp((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf((Companion_D1_.Create_DC4_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(false)).Cardinality())).Cardinality(), false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-523), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ylexru")).Cardinality()))))).Dtor_cf9(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality()))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v0 _dafny.Int
			_5_v0 = interface{}(_compr_1).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((Companion_D1_.Create_DC4_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(false)).Cardinality())).Cardinality(), false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-523), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ylexru")).Cardinality()))))).Dtor_cf9(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())), _5_v0) {
				_coll1.Add((_5_v0).Plus((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("djl")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, !(true), false)).Cardinality()))).Cardinality()), _dafny.IntOfInt64(529))
			}
		}
		return _coll1.ToMap()
	}()).Cardinality()) < 0, (true) || (!(true)))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!((_dafny.MultiSetOf(true)).IsProperSubsetOf(_dafny.MultiSetOf(false, true))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-409), _dafny.IntOfInt64(468))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _6_v0 _dafny.Int
			_6_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(-409)).Cmp(_6_v0) <= 0) && ((_6_v0).Cmp(_dafny.IntOfInt64(468)) < 0) {
				_coll2.Add((_6_v0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uicbdh")).Cardinality()))), (_dafny.MultiSetOf(_dafny.IntOfInt64(687), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(550))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				}))).Cardinality()))).Cardinality())
			}
		}
		return _coll2.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(-962))), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-196), _dafny.IntOfInt64(427)), _dafny.SeqOf(_dafny.IntOfInt64(223)))).Cardinality())), (func() D0 {
		if false {
			return Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tlkawrtbe")).Cardinality()), false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(124))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_8_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(675), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-276), false)))
		}
		return Companion_D0_.Create_DC0_(_dafny.IntOfInt64(205), false, _dafny.UnicodeSeqOfUtf8Bytes("xhtbrf"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bq"))).Cardinality()))).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-318), true)))
	})())
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(807), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wdqdlcew")).Cardinality()), _dafny.IntOfInt64(266), _dafny.IntOfInt64(135))).Cardinality(), Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("f")).Cardinality()), false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(839))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(717), func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(364), _dafny.IntOfInt64(851))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _10_v0 _dafny.Int
			_10_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(364)).Cmp(_10_v0) <= 0) && ((_10_v0).Cmp(_dafny.IntOfInt64(851)) < 0) {
				_coll3.Add((_10_v0).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-898), true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(569))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg5 _dafny.Int) interface{} {
						return coer5(arg5)
					}
				}(func(_11_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				})), func() _dafny.Map {
					var _coll4 = _dafny.NewMapBuilder()
					_ = _coll4
					for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(619), _dafny.IntOfInt64(917))); ; {
						_compr_4, _ok4 := _iter4()
						if !_ok4 {
							break
						}
						var _12_v1 _dafny.Int
						_12_v1 = interface{}(_compr_4).(_dafny.Int)
						if ((_dafny.IntOfInt64(619)).Cmp(_12_v1) <= 0) && ((_12_v1).Cmp(_dafny.IntOfInt64(917)) < 0) {
							_coll4.Add((_12_v1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("fcscl"))).Cardinality())), func() _dafny.Map {
								var _coll5 = _dafny.NewMapBuilder()
								_ = _coll5
								for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(712), _dafny.IntOfInt64(236))); ; {
									_compr_5, _ok5 := _iter5()
									if !_ok5 {
										break
									}
									var _13_v2 _dafny.Int
									_13_v2 = interface{}(_compr_5).(_dafny.Int)
									if ((_dafny.IntOfInt64(712)).Cmp(_13_v2) <= 0) && ((_13_v2).Cmp(_dafny.IntOfInt64(236)) < 0) {
										_coll5.Add(Companion_Default___.SafeDivisionInt(_13_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qokpwk")).Cardinality())), false)
									}
								}
								return _coll5.ToMap()
							}())
						}
					}
					return _coll4.ToMap()
				}()))).Cardinality()))), false)
			}
		}
		return _coll3.ToMap()
	}())))), _dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-962), ((Companion_D5_.Create_DC17_(func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-634), _dafny.IntOfInt64(9))); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _14_v3 _dafny.Int
			_14_v3 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(-634)).Cmp(_14_v3) <= 0) && ((_14_v3).Cmp(_dafny.IntOfInt64(9)) < 0) {
				_coll6.Add(Companion_Default___.SafeModuloInt(_14_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bk")).Cardinality())))
			}
		}
		return _coll6.ToSet()
	}())).Dtor_cf29()).Cardinality()), _dafny.IntOfInt64(662), Companion_D0_.Create_DC0_(_dafny.IntOfInt64(164), false, _dafny.UnicodeSeqOfUtf8Bytes("newt"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dxly")).Cardinality())), false)).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(447), true)))), Companion_D1_.Create_DC3_(func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(620), _dafny.IntOfInt64(477))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _15_v4 _dafny.Int
			_15_v4 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(620)).Cmp(_15_v4) <= 0) && ((_15_v4).Cmp(_dafny.IntOfInt64(477)) < 0) {
				_coll7.Add(Companion_Default___.SafeModuloInt(_15_v4, _dafny.IntOfInt64(77)), _dafny.IntOfInt64(486))
			}
		}
		return _coll7.ToMap()
	}(), _dafny.IntOfInt64(-215), Companion_D0_.Create_DC0_(_dafny.IntOfInt64(651), false, _dafny.UnicodeSeqOfUtf8Bytes("jhxxrjsnl"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(295), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-593), true)))), Companion_D1_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(414), (_dafny.MultiSetOf(false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(356))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_16_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))).Cardinality()), Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(867))).Cardinality(), true, _dafny.UnicodeSeqOfUtf8Bytes("ouh"), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(207), false))))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true, false, true, false), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(603))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), _dafny.IntOfInt64(-742))).Merge((Companion_D6_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true), true)), _dafny.IntOfInt64(916)))).Dtor_cf33()))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(758), func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), true)).Keys().Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _18_v0 _dafny.Int
			_18_v0 = interface{}(_compr_8).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), true)).Contains(_18_v0) {
				_coll8.Add(Companion_Default___.SafeDivisionInt(_18_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(623), false)).Cardinality()), true)
			}
		}
		return _coll8.ToMap()
	}())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bre")).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-798), true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-689), false))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 D0, p1 _dafny.Int, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mcyqkio"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("egwe"), _dafny.UnicodeSeqOfUtf8Bytes("sewnuy")))
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(_dafny.IntOfInt64(795), !(!(((func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(357), _dafny.IntOfInt64(568))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(357)).Cmp(_19_v0) <= 0) && ((_19_v0).Cmp(_dafny.IntOfInt64(568)) < 0) {
				_coll9.Add(Companion_Default___.SafeModuloInt(_19_v0, _dafny.IntOfInt64(793)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}))).Cardinality())))
			}
		}
		return _coll9.ToMap()
	}()).Cardinality()).Cmp(_dafny.IntOfInt64(907)) == 0)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-463))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_21_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	}))).Cardinality()), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(343))).Cardinality())).Cardinality(), _dafny.IntOfInt64(-382), (_dafny.Zero).Minus((func() _dafny.Set {
		var _coll10 = _dafny.NewBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(581))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_22_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		}))).Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _23_v1 _dafny.CodePoint
			_23_v1 = interface{}(_compr_10).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(581))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}(func(_22_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), _23_v1) {
				_coll10.Add(_23_v1)
			}
		}
		return _coll10.ToSet()
	}()).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(514), _dafny.MultiSetOf(_dafny.IntOfInt64(726)))))
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-785)), _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sxmprhotv")).Cardinality()))), _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.SetOf(false, true, false)).Cardinality(), _dafny.IntOfInt64(378), (_dafny.Zero).Minus(_dafny.IntOfInt64(-951)), (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality())))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-648))).Uint32(), func(coer12 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_24_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-950)))
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(242))).Uint32(), func(coer13 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_25_i1 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(_dafny.IntOfInt64(-132), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(914))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_26_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))).Cardinality())), _dafny.IntOfInt64(-96))
	}))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 _dafny.Set, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (!(false)) && (false))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 _dafny.Map, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-10), (_dafny.IntOfInt64(683)).Cmp((_dafny.SetOf(_dafny.IntOfInt64(316))).Cardinality()) > 0, (func() _dafny.Sequence {
		if true {
			return _dafny.UnicodeSeqOfUtf8Bytes("w")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("jpsbupx")
	})(), func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(671), _dafny.IntOfInt64(265))); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _27_v0 _dafny.Int
			_27_v0 = interface{}(_compr_11).(_dafny.Int)
			if ((_dafny.IntOfInt64(671)).Cmp(_27_v0) <= 0) && ((_27_v0).Cmp(_dafny.IntOfInt64(265)) < 0) {
				_coll11.Add(Companion_Default___.SafeModuloInt(_27_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-506))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}(func(_29_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}))).Cardinality()))), func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(858), _dafny.IntOfInt64(123))); ; {
						_compr_12, _ok12 := _iter12()
						if !_ok12 {
							break
						}
						var _28_v1 _dafny.Int
						_28_v1 = interface{}(_compr_12).(_dafny.Int)
						if ((_dafny.IntOfInt64(858)).Cmp(_28_v1) <= 0) && ((_28_v1).Cmp(_dafny.IntOfInt64(123)) < 0) {
							_coll12.Add(Companion_Default___.SafeDivisionInt(_28_v1, _dafny.IntOfInt64(410)), !(false))
						}
					}
					return _coll12.ToMap()
				}())
			}
		}
		return _coll11.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(642))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_30_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(822)
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(219), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("agbh")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xvgfxoaqc")).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(-211), _dafny.IntOfInt64(157), _dafny.IntOfInt64(330))).Cardinality(), _dafny.IntOfInt64(31)), _dafny.SeqOf(_dafny.IntOfInt64(801), _dafny.IntOfInt64(912))))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("vvxhidyxj"))).Difference(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(146))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_31_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	}))))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(878))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D6_.Create_DC21_(true, _dafny.CodePoint('g'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rjujka")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lxlnkch")).Cardinality()))).Dtor_cf34(), _dafny.IntOfInt64(317)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xpd")).Cardinality()))).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(931), _dafny.IntOfInt64(846)), _dafny.IntOfInt64(-596))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Array, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, _dafny.Set) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r1
	var r2 _dafny.Set = _dafny.EmptySet
	_ = r2
	var _32_v0 bool
	_ = _32_v0
	_32_v0 = true
	var _33_v1 _dafny.CodePoint
	_ = _33_v1
	_33_v1 = _dafny.CodePoint('c')
	var _34_v2 _dafny.Set
	_ = _34_v2
	_34_v2 = _dafny.SetOf(_33_v1)
	_32_v0 = (p1).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-677), (_34_v2).Cardinality())) > 0
	(globalState).F6 = p1
	r0 = ((_dafny.IntOfInt64(-453)).Plus(p1)).Plus((_dafny.Zero).Minus((p1).Times((_dafny.Zero).Minus(p1))))
	var _35_v3 _dafny.Sequence
	_ = _35_v3
	_35_v3 = _dafny.UnicodeSeqOfUtf8Bytes("mhpactow")
	var _36_v4 D0
	_ = _36_v4
	_36_v4 = Companion_D0_.Create_DC0_(p1, p3, _35_v3, Companion_Default___.Fm11(p1, _32_v0, globalState))
	var _pat_let_tv0 = _33_v1
	_ = _pat_let_tv0
	var _pat_let_tv1 = p3
	_ = _pat_let_tv1
	var _pat_let_tv2 = _36_v4
	_ = _pat_let_tv2
	var _pat_let_tv3 = p1
	_ = _pat_let_tv3
	var _pat_let_tv4 = _33_v1
	_ = _pat_let_tv4
	var _pat_let_tv5 = p1
	_ = _pat_let_tv5
	var _pat_let_tv6 = globalState
	_ = _pat_let_tv6
	var _source1 D1 = func(_source2 D0) D1 {
		var _37___mcc_h9 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
		_ = _37___mcc_h9
		var _38___mcc_h10 bool = _source2.Get_().(D0_DC0).Cf1
		_ = _38___mcc_h10
		var _39___mcc_h11 _dafny.Sequence = _source2.Get_().(D0_DC0).Cf2
		_ = _39___mcc_h11
		var _40___mcc_h12 _dafny.Map = _source2.Get_().(D0_DC0).Cf3
		_ = _40___mcc_h12
		var _41_cf3 _dafny.Map = _40___mcc_h12
		_ = _41_cf3
		var _42_cf2 _dafny.Sequence = _39___mcc_h11
		_ = _42_cf2
		var _43_cf1 bool = _38___mcc_h10
		_ = _43_cf1
		var _44_cf0 _dafny.Int = _37___mcc_h9
		_ = _44_cf0
		return Companion_D1_.Create_DC2_(_pat_let_tv0)
	}(func(_pat_let0_0 D0) D0 {
		return func(_45_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 bool) D0 {
				return func(_46_dt__update_hcf1_h0 bool) D0 {
					return func(_pat_let2_0 _dafny.Sequence) D0 {
						return func(_47_dt__update_hcf2_h0 _dafny.Sequence) D0 {
							return Companion_D0_.Create_DC0_((_45_dt__update__tmp_h0).Dtor_cf0(), _46_dt__update_hcf1_h0, _47_dt__update_hcf2_h0, (_45_dt__update__tmp_h0).Dtor_cf3())
						}(_pat_let2_0)
					}(Companion_Default___.Fm12(_pat_let_tv2, _pat_let_tv3, _pat_let_tv4, _pat_let_tv5, _pat_let_tv6))
				}(_pat_let1_0)
			}(_pat_let_tv1)
		}(_pat_let0_0)
	}(_36_v4))
	_ = _source1
	if _source1.Is_DC2() {
		var _48___mcc_h0 _dafny.CodePoint = _source1.Get_().(D1_DC2).Cf5
		_ = _48___mcc_h0
		var _49_cf5 _dafny.CodePoint = _48___mcc_h0
		_ = _49_cf5
		var _50_v5 _dafny.Array
		_ = _50_v5
		var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw0
		_50_v5 = _nw0
		var _rhs0 bool = p2
		_ = _rhs0
		var _rhs1 bool = _32_v0
		_ = _rhs1
		var _rhs2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_35_v3, _dafny.Companion_Sequence_.Concatenate(_35_v3, _dafny.UnicodeSeqOfUtf8Bytes("y")))
		_ = _rhs2
		var _rhs3 _dafny.Array = _50_v5
		_ = _rhs3
		var _rhs4 _dafny.Sequence = _35_v3
		_ = _rhs4
		_32_v0 = _rhs0
		_32_v0 = _rhs1
		_35_v3 = _rhs2
		_50_v5 = _rhs3
		_35_v3 = _rhs4
		var _51_v6 _dafny.Map
		_ = _51_v6
		_51_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_35_v3).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_35_v3, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_35_v3).Cardinality()))).Uint32(), _49_cf5)).Cardinality())), globalState), p1), (p1).Minus(p1))
		var _52_v7 _dafny.Set
		_ = _52_v7
		_52_v7 = _dafny.SetOf(_35_v3, _35_v3)
		_51_v6 = (_51_v6).Update(p1, (_52_v7).Cardinality())
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_50_v5), 0))
		_ = _index0
		(_50_v5).ArraySet1(true, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_50_v5), 0))
		_ = _index1
		var _rhs5 _dafny.Int = p1
		_ = _rhs5
		var _rhs6 bool = _32_v0
		_ = _rhs6
		var _rhs7 _dafny.Int = p1
		_ = _rhs7
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 _dafny.Array = _50_v5
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_50_v5), 0))
		_ = _lhs2
		var _lhs3 *GlobalState = globalState
		_ = _lhs3
		_lhs0.F11 = _rhs5
		(_lhs1).ArraySet1(_rhs6, (_lhs2).Int())
		_lhs3.F16 = _rhs7
		var _53_v8 _dafny.Array
		_ = _53_v8
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Int = (func(_54_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_55_i0 _dafny.Int) _dafny.Int {
					return (_55_i0).Plus(_54_p1)
				}
			})(p1)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw1 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw1).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw1).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_53_v8 = _nw1
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_53_v8), 0))
		_ = _index2
		(_53_v8).ArraySet1(p1, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_53_v8), 0))
		_ = _index3
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_50_v5), 0))
		_ = _index4
		var _rhs8 _dafny.CodePoint = _33_v1
		_ = _rhs8
		var _rhs9 _dafny.Int = p1
		_ = _rhs9
		var _rhs10 bool = _32_v0
		_ = _rhs10
		var _lhs4 _dafny.Array = _53_v8
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_53_v8), 0))
		_ = _lhs5
		var _lhs6 _dafny.Array = _50_v5
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_50_v5), 0))
		_ = _lhs7
		_33_v1 = _rhs8
		(_lhs4).ArraySet1(_rhs9, (_lhs5).Int())
		(_lhs6).ArraySet1(_rhs10, (_lhs7).Int())
	} else if _source1.Is_DC3() {
		var _56___mcc_h1 _dafny.Map = _source1.Get_().(D1_DC3).Cf6
		_ = _56___mcc_h1
		var _57___mcc_h2 _dafny.Int = _source1.Get_().(D1_DC3).Cf7
		_ = _57___mcc_h2
		var _58___mcc_h3 D0 = _source1.Get_().(D1_DC3).Cf8
		_ = _58___mcc_h3
		var _59_cf8 D0 = _58___mcc_h3
		_ = _59_cf8
		var _60_cf7 _dafny.Int = _57___mcc_h2
		_ = _60_cf7
		var _61_cf6 _dafny.Map = _56___mcc_h1
		_ = _61_cf6
		var _62_v9 *C0
		_ = _62_v9
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__(_32_v0)
		_62_v9 = _nw2
		var _63_v10 _dafny.MultiSet
		_ = _63_v10
		_63_v10 = _dafny.MultiSetOf(_62_v9)
		var _64_v11 _dafny.Map
		_ = _64_v11
		_64_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_62_v9, (func() _dafny.Int {
			if (_63_v10).Contains(_62_v9) {
				return (_63_v10).Multiplicity(_62_v9)
			}
			return p1
		})())
		var _65_v12 _dafny.Set
		_ = _65_v12
		_65_v12 = _dafny.SetOf(_64_v11, (_64_v11).Merge(_64_v11), _64_v11, (_64_v11).Merge(_64_v11))
		_65_v12 = _65_v12
		var _66_v14 _dafny.Sequence
		_ = _66_v14
		_66_v14 = _dafny.SeqOf(_60_cf7)
		var _67_v15 _dafny.Sequence
		_ = _67_v15
		_67_v15 = _dafny.SeqOf(_dafny.SeqOf((_62_v9).F17()))
		var _68_v16 _dafny.Map
		_ = _68_v16
		_68_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_60_cf7, p3)
		var _69_v17 _dafny.Set
		_ = _69_v17
		_69_v17 = _dafny.SetOf(_dafny.IntOfInt64(191))
		var _70_v18 _dafny.Map
		_ = _70_v18
		_70_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_68_v16).Contains(_dafny.IntOfInt64(898)) {
				return (_68_v16).Get(_dafny.IntOfInt64(898)).(bool)
			}
			return p2
		})(), _69_v17)
		var _71_v19 _dafny.Set
		_ = _71_v19
		_71_v19 = _dafny.SetOf(_60_cf7, (_dafny.Zero).Minus(((func() _dafny.Set {
			if (_70_v18).Contains(p2) {
				return (_70_v18).Get(p2).(_dafny.Set)
			}
			return _69_v17
		})()).Cardinality()), p1, _60_cf7)
		var _72_v20 _dafny.Array
		_ = _72_v20
		var _nwElement0_0 _dafny.Int = (_dafny.Zero).Minus(p1)
		_ = _nwElement0_0
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(28))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_0, 0)
		(_nw3).ArraySet1(_dafny.IntOfInt64(763), 1)
		(_nw3).ArraySet1((p1).Minus(Companion_Default___.Fm0(p1, _61_cf6, globalState)), 2)
		(_nw3).ArraySet1(p1, 3)
		(_nw3).ArraySet1(_dafny.IntOfInt64(947), 4)
		(_nw3).ArraySet1(p1, 5)
		(_nw3).ArraySet1((_60_cf7).Minus((func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_61_cf6).Keys().Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _73_v13 _dafny.Int
				_73_v13 = interface{}(_compr_13).(_dafny.Int)
				if (_61_cf6).Contains(_73_v13) {
					_coll13.Add((_73_v13).Times(_60_cf7), p1)
				}
			}
			return _coll13.ToMap()
		}()).Cardinality()), 6)
		(_nw3).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_35_v3, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_35_v3).Cardinality()))).Uint32(), _33_v1), _35_v3)).Cardinality()), 7)
		(_nw3).ArraySet1(p1, 8)
		(_nw3).ArraySet1((func() _dafny.Int {
			if (_63_v10).Contains(_62_v9) {
				return (_63_v10).Multiplicity(_62_v9)
			}
			return _60_cf7
		})(), 9)
		(_nw3).ArraySet1(p1, 10)
		(_nw3).ArraySet1(_60_cf7, 11)
		(_nw3).ArraySet1(_60_cf7, 12)
		(_nw3).ArraySet1(_60_cf7, 13)
		(_nw3).ArraySet1(_60_cf7, 14)
		(_nw3).ArraySet1(p1, 15)
		(_nw3).ArraySet1(_60_cf7, 16)
		(_nw3).ArraySet1((_66_v14).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_66_v14).Cardinality()))).Uint32()).(_dafny.Int), 17)
		(_nw3).ArraySet1((_66_v14).Select((Companion_Default___.SafeIndex(_60_cf7, _dafny.IntOfUint32((_66_v14).Cardinality()))).Uint32()).(_dafny.Int), 18)
		(_nw3).ArraySet1(_60_cf7, 19)
		(_nw3).ArraySet1(_dafny.IntOfUint32(((_67_v15).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_67_v15).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), 20)
		(_nw3).ArraySet1(Companion_Default___.SafeModuloInt(_60_cf7, _dafny.IntOfInt64(-884)), 21)
		(_nw3).ArraySet1((_71_v19).Cardinality(), 22)
		(_nw3).ArraySet1(p1, 23)
		(_nw3).ArraySet1(_60_cf7, 24)
		(_nw3).ArraySet1(_60_cf7, 25)
		(_nw3).ArraySet1(p1, 26)
		(_nw3).ArraySet1(_dafny.IntOfUint32((_35_v3).Cardinality()), 27)
		_72_v20 = _nw3
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_72_v20), 0))
		_ = _index5
		(_72_v20).ArraySet1((func() _dafny.Int {
			if !(p3) {
				return _dafny.IntOfInt64(183)
			}
			return p1
		})(), (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_72_v20), 0))
		_ = _index6
		var _rhs11 _dafny.Int = Companion_Default___.Fm0(p1, _61_cf6, globalState)
		_ = _rhs11
		var _rhs12 _dafny.Int = _dafny.IntOfInt64(909)
		_ = _rhs12
		var _rhs13 bool = true
		_ = _rhs13
		var _rhs14 bool = true
		_ = _rhs14
		var _rhs15 bool = _32_v0
		_ = _rhs15
		var _lhs8 _dafny.Array = _72_v20
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_72_v20), 0))
		_ = _lhs9
		var _lhs10 *GlobalState = globalState
		_ = _lhs10
		(_lhs8).ArraySet1(_rhs11, (_lhs9).Int())
		_lhs10.F16 = _rhs12
		_32_v0 = _rhs13
		_32_v0 = _rhs14
		_32_v0 = _rhs15
		var _74_v21 _dafny.MultiSet
		_ = _74_v21
		_74_v21 = _dafny.MultiSetOf(p1)
		var _75_v22 _dafny.Set
		_ = _75_v22
		_75_v22 = _dafny.SetOf((_62_v9).F17(), (_62_v9).F17())
		var _76_v23 _dafny.Map
		_ = _76_v23
		_76_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_74_v21).Difference(_74_v21), _75_v22)
		_76_v23 = _76_v23
		var _77_v24 _dafny.Sequence
		_ = _77_v24
		_77_v24 = _dafny.SeqOf(_72_v20, _72_v20, _72_v20, _72_v20)
		var _rhs16 _dafny.Sequence = _66_v14
		_ = _rhs16
		var _rhs17 *C0 = _62_v9
		_ = _rhs17
		var _rhs18 _dafny.Sequence = _77_v24
		_ = _rhs18
		_66_v14 = _rhs16
		_62_v9 = _rhs17
		_77_v24 = _rhs18
	} else if _source1.Is_DC4() {
		var _78___mcc_h4 _dafny.Int = _source1.Get_().(D1_DC4).Cf9
		_ = _78___mcc_h4
		var _79___mcc_h5 bool = _source1.Get_().(D1_DC4).Cf10
		_ = _79___mcc_h5
		var _80___mcc_h6 _dafny.Map = _source1.Get_().(D1_DC4).Cf11
		_ = _80___mcc_h6
		var _81_cf11 _dafny.Map = _80___mcc_h6
		_ = _81_cf11
		var _82_cf10 bool = _79___mcc_h5
		_ = _82_cf10
		var _83_cf9 _dafny.Int = _78___mcc_h4
		_ = _83_cf9
		var _84_v25 _dafny.Sequence
		_ = _84_v25
		_84_v25 = _dafny.SeqOf(_82_cf10, p2, p2, p3, p2)
		var _85_v26 D1
		_ = _85_v26
		_85_v26 = Companion_D1_.Create_DC1_(_84_v25)
		_85_v26 = Companion_D1_.Create_DC1_(_dafny.SeqOf(p2, p2))
		var _source3 D1 = Companion_D1_.Create_DC2_(_33_v1)
		_ = _source3
		if _source3.Is_DC2() {
			var _86___mcc_h13 _dafny.CodePoint = _source3.Get_().(D1_DC2).Cf5
			_ = _86___mcc_h13
			var _87_cf5 _dafny.CodePoint = _86___mcc_h13
			_ = _87_cf5
			var _88_v27 *C0
			_ = _88_v27
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__(p3)
			_88_v27 = _nw4
			var _89_v28 _dafny.Sequence
			_ = _89_v28
			_89_v28 = _dafny.SeqOf(_88_v27)
			var _90_v29 _dafny.Array
			_ = _90_v29
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw5
			_90_v29 = _nw5
			var _91_v30 _dafny.Map
			_ = _91_v30
			_91_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_cf9, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_v29, _87_cf5)).Cardinality())
			var _92_v31 _dafny.Map
			_ = _92_v31
			_92_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_cf10, Companion_D1_.Create_DC4_(_83_cf9, Companion_Default___.Fm4(_dafny.IntOfUint32((_89_v28).Cardinality()), (_88_v27).F17(), globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_cf9, _dafny.MultiSetOf(_83_cf9, Companion_Default___.Fm0(_dafny.IntOfUint32((_35_v3).Cardinality()), _91_v30, globalState)))))
			var _93_v33 _dafny.Map
			_ = _93_v33
			_93_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_89_v28).Cardinality()), _32_v0)
			var _94_v34 D1
			_ = _94_v34
			_94_v34 = Companion_D1_.Create_DC4_(_83_cf9, (Companion_Default___.Fm13(globalState)).Dtor_cf10(), func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_93_v33).Keys().Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _95_v32 _dafny.Int
					_95_v32 = interface{}(_compr_14).(_dafny.Int)
					if (_93_v33).Contains(_95_v32) {
						_coll14.Add((_95_v32).Plus(p1), _dafny.MultiSetOf(_83_cf9, p1))
					}
				}
				return _coll14.ToMap()
			}())
			_82_cf10 = !((_92_v31).Merge(_92_v31)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, _94_v34))
			var _96_v37 _dafny.Array
			_ = _96_v37
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_1
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Set = (func(_97_cf5 _dafny.CodePoint, _98_v3 _dafny.Sequence, _99_p1 _dafny.Int, _100_v1 _dafny.CodePoint, _101_cf9 _dafny.Int) func(_dafny.Int) _dafny.Set {
					return func(_102_i1 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(func() _dafny.Map {
							var _coll15 = _dafny.NewMapBuilder()
							_ = _coll15
							for _iter15 := _dafny.Iterate((_dafny.SeqOf((_98_v3).Select((Companion_Default___.SafeIndex(_99_p1, _dafny.IntOfUint32((_98_v3).Cardinality()))).Uint32()).(_dafny.CodePoint), _100_v1)).Elements()); ; {
								_compr_15, _ok15 := _iter15()
								if !_ok15 {
									break
								}
								var _103_v35 _dafny.CodePoint
								_103_v35 = interface{}(_compr_15).(_dafny.CodePoint)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_98_v3).Select((Companion_Default___.SafeIndex(_99_p1, _dafny.IntOfUint32((_98_v3).Cardinality()))).Uint32()).(_dafny.CodePoint), _100_v1), _103_v35) {
									_coll15.Add(_103_v35, _99_p1)
								}
							}
							return _coll15.ToMap()
						}(), func() _dafny.Map {
							var _coll16 = _dafny.NewMapBuilder()
							_ = _coll16
							for _iter16 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_100_v1))).Elements()); ; {
								_compr_16, _ok16 := _iter16()
								if !_ok16 {
									break
								}
								var _104_v36 _dafny.CodePoint
								_104_v36 = interface{}(_compr_16).(_dafny.CodePoint)
								if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_100_v1))).Contains(_104_v36) {
									_coll16.Add(_104_v36, (_dafny.Zero).Minus(_99_p1))
								}
							}
							return _coll16.ToMap()
						}())).Union(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_cf5, _101_cf9)))
					}
				})(_87_cf5, _35_v3, p1, _33_v1, _83_cf9)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw6 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw6).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw6).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_96_v37 = _nw6
			var _105_v39 _dafny.MultiSet
			_ = _105_v39
			_105_v39 = _dafny.MultiSetOf(_dafny.CodePoint('k'))
			var _106_v40 _dafny.Map
			_ = _106_v40
			_106_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_cf5, Companion_Default___.Fm0(p1, _91_v30, globalState))
			var _107_v41 _dafny.MultiSet
			_ = _107_v41
			_107_v41 = _dafny.MultiSetOf(func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate((_105_v39).Elements()); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _108_v38 _dafny.CodePoint
					_108_v38 = interface{}(_compr_17).(_dafny.CodePoint)
					if (_105_v39).Contains(_108_v38) {
						_coll17.Add(_108_v38, (_dafny.MultiSetOf(_82_cf10)).Cardinality())
					}
				}
				return _coll17.ToMap()
			}(), _106_v40)
			var _109_v43 _dafny.Set
			_ = _109_v43
			_109_v43 = _dafny.SetOf(_106_v40)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_96_v37), 0))
			_ = _index7
			(_96_v37).ArraySet1((func() _dafny.Set {
				var _coll18 = _dafny.NewBuilder()
				_ = _coll18
				for _iter18 := _dafny.Iterate((_107_v41).Elements()); ; {
					_compr_18, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _110_v42 _dafny.Map
					_110_v42 = interface{}(_compr_18).(_dafny.Map)
					if (_107_v41).Contains(_110_v42) {
						_coll18.Add(_110_v42)
					}
				}
				return _coll18.ToSet()
			}()).Difference(_109_v43), (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_96_v37), 0))
			_ = _index8
			(_96_v37).ArraySet1((_109_v43).Intersection(_109_v43), (_index8).Int())
			_91_v30 = (_91_v30).Merge(_91_v30)
		} else if _source3.Is_DC3() {
			var _111___mcc_h14 _dafny.Map = _source3.Get_().(D1_DC3).Cf6
			_ = _111___mcc_h14
			var _112___mcc_h15 _dafny.Int = _source3.Get_().(D1_DC3).Cf7
			_ = _112___mcc_h15
			var _113___mcc_h16 D0 = _source3.Get_().(D1_DC3).Cf8
			_ = _113___mcc_h16
			var _114_cf8 D0 = _113___mcc_h16
			_ = _114_cf8
			var _115_cf7 _dafny.Int = _112___mcc_h15
			_ = _115_cf7
			var _116_cf6 _dafny.Map = _111___mcc_h14
			_ = _116_cf6
			var _117_v44 _dafny.Set
			_ = _117_v44
			_117_v44 = _dafny.SetOf((_116_cf6).Cardinality())
			var _118_v45 _dafny.Array
			_ = _118_v45
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_2
			var _nw7 _dafny.Array
			_ = _nw7
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw7 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.CodePoint = (func(_119_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_120_i3 _dafny.Int) _dafny.CodePoint {
						return _119_v1
					}
				})(_33_v1)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw7 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw7).ArraySet1CodePoint(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw7).ArraySet1CodePoint(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_118_v45 = _nw7
			_82_cf10 = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(607))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_121_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_122_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_121_v3).Cardinality())
				}
			})(_35_v3)))).Cardinality())).Cmp(Companion_Default___.SafeModuloInt((_117_v44).Cardinality(), (_dafny.MultiSetOf(_118_v45)).Cardinality())) <= 0
			var _123_v46 _dafny.Array
			_ = _123_v46
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw8
			_123_v46 = _nw8
			var _124_v47 *C0
			_ = _124_v47
			var _nw9 *C0 = New_C0_()
			_ = _nw9
			_nw9.Ctor__(_82_cf10)
			_124_v47 = _nw9
			var _125_v48 _dafny.MultiSet
			_ = _125_v48
			_125_v48 = _dafny.MultiSetOf(p2)
			var _126_v49 _dafny.Map
			_ = _126_v49
			_126_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_124_v47).F17())
			var _127_v50 _dafny.Map
			_ = _127_v50
			_127_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_cf9, _126_v49)
			var _128_v51 _dafny.Map
			_ = _128_v51
			_128_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v47, Companion_Default___.Fm4(_83_cf9, (Companion_D0_.Create_DC0_((func() _dafny.Int {
				if (_125_v48).Contains(Companion_Default___.Fm4(_dafny.IntOfInt64(-968), p3, globalState)) {
					return (_125_v48).Multiplicity(Companion_Default___.Fm4(_dafny.IntOfInt64(-968), p3, globalState))
				}
				return _83_cf9
			})(), Companion_Default___.Fm4(_83_cf9, !(_32_v0), globalState), _dafny.UnicodeSeqOfUtf8Bytes("suaxnms"), _127_v50)).Dtor_cf1(), globalState))
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_123_v46), 0))
			_ = _index9
			(_123_v46).ArraySet1((func() bool {
				if (_128_v51).Contains(_124_v47) {
					return (_128_v51).Get(_124_v47).(bool)
				}
				return _32_v0
			})(), (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_123_v46), 0))
			_ = _index10
			(_123_v46).ArraySet1((_124_v47).F17(), (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_123_v46), 0))
			_ = _index11
			(_123_v46).ArraySet1((_117_v44).IsSubsetOf(func() _dafny.Set {
				var _coll19 = _dafny.NewBuilder()
				_ = _coll19
				for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(509), _dafny.IntOfInt64(976))); ; {
					_compr_19, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _129_v52 _dafny.Int
					_129_v52 = interface{}(_compr_19).(_dafny.Int)
					if ((_dafny.IntOfInt64(509)).Cmp(_129_v52) <= 0) && ((_129_v52).Cmp(_dafny.IntOfInt64(976)) < 0) {
						_coll19.Add((_129_v52).Plus(_83_cf9))
					}
				}
				return _coll19.ToSet()
			}()), (_index11).Int())
			var _130_v53 _dafny.Sequence
			_ = _130_v53
			_130_v53 = _dafny.SeqOf(_123_v46, _123_v46, _123_v46, _123_v46, _123_v46)
			_130_v53 = _130_v53
		} else if _source3.Is_DC4() {
			var _131___mcc_h17 _dafny.Int = _source3.Get_().(D1_DC4).Cf9
			_ = _131___mcc_h17
			var _132___mcc_h18 bool = _source3.Get_().(D1_DC4).Cf10
			_ = _132___mcc_h18
			var _133___mcc_h19 _dafny.Map = _source3.Get_().(D1_DC4).Cf11
			_ = _133___mcc_h19
			var _134_cf11 _dafny.Map = _133___mcc_h19
			_ = _134_cf11
			var _135_cf10 bool = _132___mcc_h18
			_ = _135_cf10
			var _136_cf9 _dafny.Int = _131___mcc_h17
			_ = _136_cf9
			_33_v1 = _33_v1
			var _137_v54 _dafny.MultiSet
			_ = _137_v54
			_137_v54 = _dafny.MultiSetOf(_32_v0)
			(globalState).F6 = (func() _dafny.Int {
				if (_137_v54).Contains(_32_v0) {
					return (_137_v54).Multiplicity(_32_v0)
				}
				return (_dafny.IntOfUint32((_dafny.SeqOf(!(p3))).Cardinality())).Plus(_136_cf9)
			})()
			var _138_v55 _dafny.Array
			_ = _138_v55
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw10
			_138_v55 = _nw10
			var _139_v56 _dafny.Map
			_ = _139_v56
			_139_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_135_cf10, _138_v55)
			var _140_v57 _dafny.Sequence
			_ = _140_v57
			_140_v57 = _dafny.SeqOf(_136_cf9)
			var _141_v58 _dafny.Array
			_ = _141_v58
			var _nwElement0_1 bool = p3
			_ = _nwElement0_1
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(3))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_1, 0)
			(_nw11).ArraySet1(p3, 1)
			(_nw11).ArraySet1(_135_cf10, 2)
			_141_v58 = _nw11
			var _142_v59 _dafny.Map
			_ = _142_v59
			_142_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Plus((_140_v57).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_140_v57).Cardinality()))).Uint32()).(_dafny.Int)), _141_v58)
			var _143_v60 _dafny.Sequence
			_ = _143_v60
			_143_v60 = _dafny.SeqOf(_139_v56, (_139_v56).Merge(_139_v56))
			var _144_v61 _dafny.Map
			_ = _144_v61
			_144_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(823))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_145_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_146_i4 _dafny.Int) _dafny.CodePoint {
					return _145_v1
				}
			})(_33_v1))), p1)
			var _147_v62 _dafny.Set
			_ = _147_v62
			_147_v62 = _dafny.SetOf(_136_cf9)
			var _148_v63 *C0
			_ = _148_v63
			var _nw12 *C0 = New_C0_()
			_ = _nw12
			_nw12.Ctor__(Companion_Default___.Fm4((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, p3)).Cardinality(), _136_cf9)).Cardinality(), p3, globalState))
			_148_v63 = _nw12
			var _149_v64 _dafny.MultiSet
			_ = _149_v64
			_149_v64 = _dafny.MultiSetOf(_148_v63, _148_v63, _148_v63)
			var _150_v65 _dafny.Map
			_ = _150_v65
			_150_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_147_v62).Cardinality(), (_dafny.Zero).Minus((_149_v64).Cardinality()))
			var _151_v66 _dafny.Map
			_ = _151_v66
			_151_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _150_v65)
			var _rhs19 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_136_cf9, _136_cf9)))
			_ = _rhs19
			var _rhs20 _dafny.Map = (_143_v60).Select((Companion_Default___.SafeIndex((_136_cf9).Minus(p1), _dafny.IntOfUint32((_143_v60).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs20
			var _rhs21 _dafny.Int = Companion_Default___.Fm0((p1).Times((_144_v61).Cardinality()), (func() _dafny.Map {
				if (_151_v66).Contains(true) {
					return (_151_v66).Get(true).(_dafny.Map)
				}
				return _150_v65
			})(), globalState)
			_ = _rhs21
			var _rhs22 _dafny.Map = _142_v59
			_ = _rhs22
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			_lhs11.F11 = _rhs19
			_139_v56 = _rhs20
			_lhs12.F11 = _rhs21
			_142_v59 = _rhs22
			var _152_v67 _dafny.Sequence
			_ = _152_v67
			_152_v67 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(982))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_153_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-300))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_154_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_155_i6 _dafny.Int) _dafny.CodePoint {
					return _154_v1
				}
			})(_33_v1))), _35_v3))
			_35_v3 = _dafny.Companion_Sequence_.Update((_152_v67).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p1, p1), _dafny.IntOfUint32((_152_v67).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_136_cf9, _dafny.IntOfUint32(((_152_v67).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p1, p1), _dafny.IntOfUint32((_152_v67).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _dafny.CodePoint('j'))
		} else if _source3.Is_DC1() {
			var _156___mcc_h20 _dafny.Sequence = _source3.Get_().(D1_DC1).Cf4
			_ = _156___mcc_h20
			var _157_cf4 _dafny.Sequence = _156___mcc_h20
			_ = _157_cf4
			(globalState).F6 = _83_cf9
			var _158_v69 *C0
			_ = _158_v69
			var _nw13 *C0 = New_C0_()
			_ = _nw13
			_nw13.Ctor__(!(_82_cf10))
			_158_v69 = _nw13
			var _159_v70 _dafny.Map
			_ = _159_v70
			_159_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v69, _83_cf9)
			(globalState).F6 = ((func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate((_35_v3).Elements()); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _160_v68 _dafny.CodePoint
					_160_v68 = interface{}(_compr_20).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_35_v3, _160_v68) {
						_coll20.Add(_160_v68, _dafny.IntOfInt64(860))
					}
				}
				return _coll20.ToMap()
			}()).Cardinality()).Minus(((_159_v70).Cardinality()).Plus(p1))
			_158_v69 = _158_v69
			var _161_v71 D1
			_ = _161_v71
			_161_v71 = Companion_D1_.Create_DC2_(Companion_Default___.Fm3(_82_cf10, globalState))
			_161_v71 = func(_pat_let3_0 D1) D1 {
				return func(_162_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let4_0 _dafny.CodePoint) D1 {
						return func(_163_dt__update_hcf5_h0 _dafny.CodePoint) D1 {
							return Companion_D1_.Create_DC2_(_163_dt__update_hcf5_h0)
						}(_pat_let4_0)
					}(_dafny.CodePoint('q'))
				}(_pat_let3_0)
			}(_161_v71)
		} else {
			var _164___mcc_h21 D1 = _source3.Get_().(D1_DC5).Cf12
			_ = _164___mcc_h21
			var _165_cf12 D1 = _164___mcc_h21
			_ = _165_cf12
			var _166_v72 D1
			_ = _166_v72
			_166_v72 = Companion_D1_.Create_DC4_(p1, true, _81_cf11)
			var _167_v73 _dafny.Map
			_ = _167_v73
			_167_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_166_v72).Dtor_cf10(), _dafny.SeqOf(_83_cf9))
			var _168_v74 _dafny.Array
			_ = _168_v74
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw14
			_168_v74 = _nw14
			var _169_v75 *C0
			_ = _169_v75
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(p2)
			_169_v75 = _nw15
			var _170_v76 _dafny.Array
			_ = _170_v76
			var _nwElement0_2 *C0 = _169_v75
			_ = _nwElement0_2
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(8))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_2, 0)
			(_nw16).ArraySet1(_169_v75, 1)
			(_nw16).ArraySet1(_169_v75, 2)
			(_nw16).ArraySet1(_169_v75, 3)
			(_nw16).ArraySet1(_169_v75, 4)
			(_nw16).ArraySet1(_169_v75, 5)
			(_nw16).ArraySet1(_169_v75, 6)
			(_nw16).ArraySet1(_169_v75, 7)
			_170_v76 = _nw16
			var _171_v77 _dafny.Map
			_ = _171_v77
			_171_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v74, _170_v76)
			var _172_v78 _dafny.Set
			_ = _172_v78
			_172_v78 = _dafny.SetOf((_dafny.MultiSetOf(_83_cf9)).Cardinality())
			var _173_v79 _dafny.Map
			_ = _173_v79
			_173_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_cf9, (_172_v78).Cardinality())
			var _rhs23 _dafny.Int = _83_cf9
			_ = _rhs23
			var _rhs24 _dafny.Int = ((_171_v77).Merge(_171_v77)).Cardinality()
			_ = _rhs24
			var _rhs25 _dafny.Map = (_167_v73).Merge(_167_v73)
			_ = _rhs25
			var _rhs26 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(p1, _83_cf9), (_173_v79).Cardinality())
			_ = _rhs26
			var _rhs27 _dafny.CodePoint = _33_v1
			_ = _rhs27
			var _lhs13 *GlobalState = globalState
			_ = _lhs13
			var _lhs14 *GlobalState = globalState
			_ = _lhs14
			r0 = _rhs23
			_lhs13.F16 = _rhs24
			_167_v73 = _rhs25
			_lhs14.F6 = _rhs26
			_33_v1 = _rhs27
			var _174_v80 _dafny.Array
			_ = _174_v80
			var _nw17 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw17
			_174_v80 = _nw17
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_168_v74), 0))
			_ = _index12
			(_168_v74).ArraySet1((func() _dafny.Int {
				if (_173_v79).Contains(p1) {
					return (_173_v79).Get(p1).(_dafny.Int)
				}
				return (_36_v4).Dtor_cf0()
			})(), (_index12).Int())
			var _175_v81 _dafny.Sequence
			_ = _175_v81
			_175_v81 = _dafny.SeqOf(_174_v80, _174_v80)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_168_v74), 0))
			_ = _index13
			var _rhs28 _dafny.Array = _174_v80
			_ = _rhs28
			var _rhs29 _dafny.Int = Companion_Default___.SafeDivisionInt(_83_cf9, p1)
			_ = _rhs29
			var _rhs30 _dafny.Int = _dafny.IntOfUint32((_175_v81).Cardinality())
			_ = _rhs30
			var _rhs31 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(258), p1)
			_ = _rhs31
			var _lhs15 _dafny.Array = _168_v74
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_168_v74), 0))
			_ = _lhs16
			var _lhs17 *GlobalState = globalState
			_ = _lhs17
			_174_v80 = _rhs28
			_83_cf9 = _rhs29
			(_lhs15).ArraySet1(_rhs30, (_lhs16).Int())
			_lhs17.F6 = _rhs31
			var _176_v82 _dafny.MultiSet
			_ = _176_v82
			_176_v82 = _dafny.MultiSetOf(p3, (func() bool {
				if (_169_v75).F17() {
					return _82_cf10
				}
				return false
			})(), _32_v0)
			_176_v82 = _176_v82
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_170_v76), 0))
			_ = _index14
			(_170_v76).ArraySet1(_169_v75, (_index14).Int())
			var _177_v83 _dafny.Array
			_ = _177_v83
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw18
			_177_v83 = _nw18
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_177_v83), 0))
			_ = _index15
			(_177_v83).ArraySet1((_175_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-373), _dafny.IntOfUint32((_175_v81).Cardinality()))).Uint32()).(_dafny.Array), (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_174_v80), 0))
			_ = _index16
			(_174_v80).ArraySet1(p3, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_170_v76), 0))
			_ = _index17
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_177_v83), 0))
			_ = _index18
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_174_v80), 0))
			_ = _index19
			var _rhs32 *C0 = _169_v75
			_ = _rhs32
			var _rhs33 bool = (Companion_Default___.SafeModuloInt(p1, Companion_Default___.Fm0((_168_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_168_v74), 0))).Int()).(_dafny.Int), _173_v79, globalState))).Cmp((_168_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_168_v74), 0))).Int()).(_dafny.Int)) > 0
			_ = _rhs33
			var _rhs34 _dafny.Array = _174_v80
			_ = _rhs34
			var _rhs35 bool = p2
			_ = _rhs35
			var _rhs36 bool = p2
			_ = _rhs36
			var _lhs18 _dafny.Array = _170_v76
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_170_v76), 0))
			_ = _lhs19
			var _lhs20 _dafny.Array = _177_v83
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_177_v83), 0))
			_ = _lhs21
			var _lhs22 _dafny.Array = _174_v80
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_174_v80), 0))
			_ = _lhs23
			(_lhs18).ArraySet1(_rhs32, (_lhs19).Int())
			_82_cf10 = _rhs33
			(_lhs20).ArraySet1(_rhs34, (_lhs21).Int())
			(_lhs22).ArraySet1(_rhs35, (_lhs23).Int())
			_32_v0 = _rhs36
		}
		var _178_v84 _dafny.Array
		_ = _178_v84
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
		_ = _nw19
		_178_v84 = _nw19
		_178_v84 = _178_v84
		var _179_v85 _dafny.Map
		_ = _179_v85
		_179_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _83_cf9)
		var _180_v86 _dafny.MultiSet
		_ = _180_v86
		_180_v86 = _dafny.MultiSetOf((_84_v25).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_84_v25).Cardinality()))).Uint32()).(bool))
		var _pat_let_tv7 = _180_v86
		_ = _pat_let_tv7
		var _source4 D1 = func(_pat_let5_0 D1) D1 {
			return func(_181_dt__update__tmp_h2 D1) D1 {
				return func(_pat_let6_0 _dafny.Int) D1 {
					return func(_182_dt__update_hcf7_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC3_((_181_dt__update__tmp_h2).Dtor_cf6(), _182_dt__update_hcf7_h0, (_181_dt__update__tmp_h2).Dtor_cf8())
					}(_pat_let6_0)
				}((_pat_let_tv7).Cardinality())
			}(_pat_let5_0)
		}(Companion_D1_.Create_DC3_(_179_v85, p1, _36_v4))
		_ = _source4
		if _source4.Is_DC2() {
			var _183___mcc_h22 _dafny.CodePoint = _source4.Get_().(D1_DC2).Cf5
			_ = _183___mcc_h22
			var _184_cf5 _dafny.CodePoint = _183___mcc_h22
			_ = _184_cf5
			var _rhs37 bool = (_82_cf10) || ((!(_32_v0)) || (_32_v0))
			_ = _rhs37
			var _rhs38 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("pufjq")
			_ = _rhs38
			_82_cf10 = _rhs37
			_35_v3 = _rhs38
			var _185_v87 _dafny.Array
			_ = _185_v87
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
			_ = _nw20
			_185_v87 = _nw20
			var _186_v88 _dafny.Array
			_ = _186_v88
			var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw21
			_186_v88 = _nw21
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_185_v87), 0))
			_ = _index20
			(_185_v87).ArraySet1(_dafny.SetOf(_186_v88, _186_v88, _186_v88), (_index20).Int())
			var _187_v89 _dafny.Set
			_ = _187_v89
			_187_v89 = _dafny.SetOf(_186_v88, _186_v88, _186_v88, _186_v88)
			var _188_v90 _dafny.Sequence
			_ = _188_v90
			_188_v90 = _dafny.SeqOf(_187_v89, _187_v89, _dafny.SetOf(_186_v88, _186_v88), _187_v89)
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_185_v87), 0))
			_ = _index21
			(_185_v87).ArraySet1(((_187_v89).Intersection((_188_v90).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_188_v90).Cardinality()))).Uint32()).(_dafny.Set))).Intersection(_187_v89), (_index21).Int())
			_82_cf10 = (_180_v86).IsProperSubsetOf((_dafny.MultiSetFromSeq(_84_v25)).Update(p3, Companion_Default___.Abs(_83_cf9)))
			var _189_v91 *C0
			_ = _189_v91
			var _nw22 *C0 = New_C0_()
			_ = _nw22
			_nw22.Ctor__(p3)
			_189_v91 = _nw22
		} else if _source4.Is_DC3() {
			var _190___mcc_h23 _dafny.Map = _source4.Get_().(D1_DC3).Cf6
			_ = _190___mcc_h23
			var _191___mcc_h24 _dafny.Int = _source4.Get_().(D1_DC3).Cf7
			_ = _191___mcc_h24
			var _192___mcc_h25 D0 = _source4.Get_().(D1_DC3).Cf8
			_ = _192___mcc_h25
			var _193_cf8 D0 = _192___mcc_h25
			_ = _193_cf8
			var _194_cf7 _dafny.Int = _191___mcc_h24
			_ = _194_cf7
			var _195_cf6 _dafny.Map = _190___mcc_h23
			_ = _195_cf6
			_33_v1 = _33_v1
			var _196_v92 _dafny.Array
			_ = _196_v92
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw23
			_196_v92 = _nw23
			var _197_v93 _dafny.Map
			_ = _197_v93
			_197_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.Companion_Sequence_.Update(_35_v3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.IntOfUint32((_35_v3).Cardinality()))).Uint32(), _dafny.CodePoint('t')))
			var _198_v94 _dafny.Map
			_ = _198_v94
			_198_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, p1)
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_196_v92), 0))
			_ = _index22
			(_196_v92).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_197_v93).Contains((func() _dafny.Int {
					if (_198_v94).Contains(_82_cf10) {
						return (_198_v94).Get(_82_cf10).(_dafny.Int)
					}
					return _194_cf7
				})()) {
					return (_197_v93).Get((func() _dafny.Int {
						if (_198_v94).Contains(_82_cf10) {
							return (_198_v94).Get(_82_cf10).(_dafny.Int)
						}
						return _194_cf7
					})()).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(587))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}((func(_199_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_200_i7 _dafny.Int) _dafny.CodePoint {
						return _199_v1
					}
				})(_33_v1)))
			})()).Cardinality()), (_index22).Int())
			var _201_v95 _dafny.Array
			_ = _201_v95
			var _nwElement0_3 bool = _82_cf10
			_ = _nwElement0_3
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(18))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_3, 0)
			(_nw24).ArraySet1(_82_cf10, 1)
			(_nw24).ArraySet1(p2, 2)
			(_nw24).ArraySet1(p2, 3)
			(_nw24).ArraySet1(p3, 4)
			(_nw24).ArraySet1(p2, 5)
			(_nw24).ArraySet1((_dafny.IntOfInt64(-500)).Cmp(p1) <= 0, 6)
			(_nw24).ArraySet1(_82_cf10, 7)
			(_nw24).ArraySet1(!((_84_v25).Select((Companion_Default___.SafeIndex(_194_cf7, _dafny.IntOfUint32((_84_v25).Cardinality()))).Uint32()).(bool)), 8)
			(_nw24).ArraySet1(_32_v0, 9)
			(_nw24).ArraySet1(((Companion_Default___.Fm6(p2, p3, p2, globalState)).Update(p2, Companion_Default___.Abs(_83_cf9))).IsDisjointFrom(_180_v86), 10)
			(_nw24).ArraySet1(p3, 11)
			(_nw24).ArraySet1(p3, 12)
			(_nw24).ArraySet1(false, 13)
			(_nw24).ArraySet1((p1).Cmp(_194_cf7) == 0, 14)
			(_nw24).ArraySet1(p2, 15)
			(_nw24).ArraySet1(p2, 16)
			(_nw24).ArraySet1(true, 17)
			_201_v95 = _nw24
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_201_v95), 0))
			_ = _index23
			(_201_v95).ArraySet1(_32_v0, (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_196_v92), 0))
			_ = _index24
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_201_v95), 0))
			_ = _index25
			var _rhs39 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("udu"), _35_v3), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("udu"), _35_v3)).Cardinality()))).Uint32(), _33_v1)).Cardinality())).Cmp(_dafny.IntOfInt64(345)) >= 0
			_ = _rhs39
			var _rhs40 bool = p3
			_ = _rhs40
			var _rhs41 _dafny.Int = Companion_Default___.SafeModuloInt(_194_cf7, _83_cf9)
			_ = _rhs41
			var _rhs42 bool = _82_cf10
			_ = _rhs42
			var _rhs43 _dafny.Int = (_194_cf7).Minus(p1)
			_ = _rhs43
			var _lhs24 _dafny.Array = _196_v92
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_196_v92), 0))
			_ = _lhs25
			var _lhs26 _dafny.Array = _201_v95
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_201_v95), 0))
			_ = _lhs27
			var _lhs28 *GlobalState = globalState
			_ = _lhs28
			_32_v0 = _rhs39
			_82_cf10 = _rhs40
			(_lhs24).ArraySet1(_rhs41, (_lhs25).Int())
			(_lhs26).ArraySet1(_rhs42, (_lhs27).Int())
			_lhs28.F16 = _rhs43
			var _202_v96 D2
			_ = _202_v96
			_202_v96 = Companion_D2_.Create_DC6_(_196_v92)
			var _203_v97 _dafny.Sequence
			_ = _203_v97
			_203_v97 = _dafny.SeqOf((_202_v96).Dtor_cf13(), _196_v92, _196_v92)
			var _204_v98 _dafny.Array
			_ = _204_v98
			var _nw25 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
			_ = _nw25
			_204_v98 = _nw25
			var _205_v99 *C0
			_ = _205_v99
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__(_82_cf10)
			_205_v99 = _nw26
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_204_v98), 0))
			_ = _index26
			(_204_v98).ArraySet1(_205_v99, (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_201_v95), 0))
			_ = _index27
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_204_v98), 0))
			_ = _index28
			var _rhs44 bool = (_84_v25).Select((Companion_Default___.SafeIndex(_83_cf9, _dafny.IntOfUint32((_84_v25).Cardinality()))).Uint32()).(bool)
			_ = _rhs44
			var _rhs45 _dafny.Sequence = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_83_cf9).Cmp(p1) > 0 {
					return _dafny.Companion_Sequence_.Concatenate(_203_v97, _203_v97)
				}
				return _203_v97
			})(), (Companion_Default___.SafeIndex((_196_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_196_v92), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_83_cf9).Cmp(p1) > 0 {
					return _dafny.Companion_Sequence_.Concatenate(_203_v97, _203_v97)
				}
				return _203_v97
			})()).Cardinality()))).Uint32(), _196_v92)
			_ = _rhs45
			var _rhs46 *C0 = _205_v99
			_ = _rhs46
			var _rhs47 _dafny.Array = _201_v95
			_ = _rhs47
			var _lhs29 _dafny.Array = _201_v95
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_201_v95), 0))
			_ = _lhs30
			var _lhs31 _dafny.Array = _204_v98
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_204_v98), 0))
			_ = _lhs32
			(_lhs29).ArraySet1(_rhs44, (_lhs30).Int())
			_203_v97 = _rhs45
			(_lhs31).ArraySet1(_rhs46, (_lhs32).Int())
			_201_v95 = _rhs47
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_201_v95), 0))
			_ = _index29
			(_201_v95).ArraySet1((_83_cf9).Cmp(p1) <= 0, (_index29).Int())
		} else if _source4.Is_DC4() {
			var _206___mcc_h26 _dafny.Int = _source4.Get_().(D1_DC4).Cf9
			_ = _206___mcc_h26
			var _207___mcc_h27 bool = _source4.Get_().(D1_DC4).Cf10
			_ = _207___mcc_h27
			var _208___mcc_h28 _dafny.Map = _source4.Get_().(D1_DC4).Cf11
			_ = _208___mcc_h28
			var _209_cf11 _dafny.Map = _208___mcc_h28
			_ = _209_cf11
			var _210_cf10 bool = _207___mcc_h27
			_ = _210_cf10
			var _211_cf9 _dafny.Int = _206___mcc_h26
			_ = _211_cf9
			var _212_v100 *C0
			_ = _212_v100
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__(_210_cf10)
			_212_v100 = _nw27
			var _213_v101 _dafny.Map
			_ = _213_v101
			_213_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, _210_cf10)
			(globalState).F6 = (_213_v101).Cardinality()
			var _214_v102 _dafny.Array
			_ = _214_v102
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_3
			var _nw28 _dafny.Array
			_ = _nw28
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw28 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = (func(_215_p3 bool) func(_dafny.Int) bool {
					return func(_216_i8 _dafny.Int) bool {
						return _215_p3
					}
				})(p3)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw28 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw28).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw28).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_214_v102 = _nw28
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_214_v102), 0))
			_ = _index30
			(_214_v102).ArraySet1(p2, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_214_v102), 0))
			_ = _index31
			(_214_v102).ArraySet1((func() bool {
				if (_213_v101).Contains(!(_dafny.Companion_Sequence_.IsPrefixOf(_35_v3, _35_v3))) {
					return (_213_v101).Get(!(_dafny.Companion_Sequence_.IsPrefixOf(_35_v3, _35_v3))).(bool)
				}
				return _210_cf10
			})(), (_index31).Int())
			(globalState).F16 = _dafny.IntOfUint32((_35_v3).Cardinality())
		} else if _source4.Is_DC1() {
			var _217___mcc_h29 _dafny.Sequence = _source4.Get_().(D1_DC1).Cf4
			_ = _217___mcc_h29
			var _218_cf4 _dafny.Sequence = _217___mcc_h29
			_ = _218_cf4
			var _219_v103 *C0
			_ = _219_v103
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__(!(p2) || (!(!(p2))))
			_219_v103 = _nw29
			var _220_v104 _dafny.Set
			_ = _220_v104
			_220_v104 = _dafny.SetOf((_219_v103).F17())
			var _221_v105 _dafny.Map
			_ = _221_v105
			_221_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_cf9, (_84_v25).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v103).F17(), (_219_v103).F17())).Cardinality(), _dafny.IntOfUint32((_84_v25).Cardinality()))).Uint32()).(bool))
			_32_v0 = Companion_Default___.Fm4(Companion_Default___.SafeDivisionInt(p1, (_220_v104).Cardinality()), (func() bool {
				if (_221_v105).Contains(_83_cf9) {
					return (_221_v105).Get(_83_cf9).(bool)
				}
				return p2
			})(), globalState)
			var _222_v106 _dafny.Array
			_ = _222_v106
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw30
			_222_v106 = _nw30
			_222_v106 = _222_v106
			var _223_v107 _dafny.Sequence
			_ = _223_v107
			_223_v107 = _dafny.SeqOf(_219_v103)
			var _224_v108 _dafny.Sequence
			_ = _224_v108
			_224_v108 = _dafny.SeqOf(p1, p1, _83_cf9, p1, p1)
			var _rhs48 *C0 = _219_v103
			_ = _rhs48
			var _rhs49 *C0 = _219_v103
			_ = _rhs49
			var _rhs50 bool = !(!_dafny.Companion_Sequence_.Equal(_223_v107, _223_v107))
			_ = _rhs50
			var _rhs51 *C0 = (_223_v107).Select((Companion_Default___.SafeIndex((p1).Times((_224_v108).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.IntOfUint32((_224_v108).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfUint32((_223_v107).Cardinality()))).Uint32()).(*C0)
			_ = _rhs51
			_219_v103 = _rhs48
			_219_v103 = _rhs49
			_32_v0 = _rhs50
			_219_v103 = _rhs51
		} else {
			var _225___mcc_h30 D1 = _source4.Get_().(D1_DC5).Cf12
			_ = _225___mcc_h30
			var _226_cf12 D1 = _225___mcc_h30
			_ = _226_cf12
			(globalState).F6 = p1
			var _227_v109 *C0
			_ = _227_v109
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__(p3)
			_227_v109 = _nw31
			var _228_v110 _dafny.Map
			_ = _228_v110
			_228_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_cf9, _227_v109)
			_228_v110 = (_228_v110).Update((_dafny.Zero).Minus(_83_cf9), _227_v109)
			var _229_v111 *C0
			_ = _229_v111
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__(p2)
			_229_v111 = _nw32
			var _rhs52 _dafny.Int = _dafny.IntOfInt64(-453)
			_ = _rhs52
			var _rhs53 _dafny.Int = (((_dafny.Zero).Minus(p1)).Minus(p1)).Times(_dafny.IntOfUint32((Companion_Default___.Fm14(globalState)).Cardinality()))
			_ = _rhs53
			var _lhs33 *GlobalState = globalState
			_ = _lhs33
			var _lhs34 *GlobalState = globalState
			_ = _lhs34
			_lhs33.F11 = _rhs52
			_lhs34.F6 = _rhs53
		}
	} else if _source1.Is_DC1() {
		var _230___mcc_h7 _dafny.Sequence = _source1.Get_().(D1_DC1).Cf4
		_ = _230___mcc_h7
		var _231_cf4 _dafny.Sequence = _230___mcc_h7
		_ = _231_cf4
		var _232_v112 *C0
		_ = _232_v112
		var _nw33 *C0 = New_C0_()
		_ = _nw33
		_nw33.Ctor__(p2)
		_232_v112 = _nw33
		_232_v112 = _232_v112
		var _233_v113 _dafny.Map
		_ = _233_v113
		_233_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v3, _232_v112)
		_232_v112 = (func() *C0 {
			if (_233_v113).Contains((func() _dafny.Sequence {
				if !(!(p2)) {
					return _35_v3
				}
				return _35_v3
			})()) {
				return (_233_v113).Get((func() _dafny.Sequence {
					if !(!(p2)) {
						return _35_v3
					}
					return _35_v3
				})()).(*C0)
			}
			return (func() *C0 {
				if false {
					return _232_v112
				}
				return _232_v112
			})()
		})()
		r0 = p1
		var _234_v114 _dafny.Map
		_ = _234_v114
		_234_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_232_v112).F17(), p1)
		var _235_v115 _dafny.Sequence
		_ = _235_v115
		_235_v115 = _dafny.SeqOf(_232_v112)
		var _236_v116 _dafny.Map
		_ = _236_v116
		_236_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _237_v117 _dafny.Map
		_ = _237_v117
		_237_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.IntOfUint32((_235_v115).Cardinality()), _236_v116, globalState), p1)
		var _238_v118 _dafny.Map
		_ = _238_v118
		_238_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _32_v0)
		var _239_v119 D1
		_ = _239_v119
		_239_v119 = Companion_D1_.Create_DC4_(Companion_Default___.SafeDivisionInt(p1, p1), (_231_cf4).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_231_cf4).Cardinality()))).Uint32()).(bool), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.MultiSetOf((func() _dafny.Int {
			if (_234_v114).Contains((_231_cf4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(446), (func() _dafny.Int {
				if (_236_v116).Contains(p1) {
					return (_236_v116).Get(p1).(_dafny.Int)
				}
				return p1
			})())).Cardinality()), _dafny.IntOfUint32((_231_cf4).Cardinality()))).Uint32()).(bool)) {
				return (_234_v114).Get((_231_cf4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(446), (func() _dafny.Int {
					if (_236_v116).Contains(p1) {
						return (_236_v116).Get(p1).(_dafny.Int)
					}
					return p1
				})())).Cardinality()), _dafny.IntOfUint32((_231_cf4).Cardinality()))).Uint32()).(bool)).(_dafny.Int)
			}
			return (_238_v118).Cardinality()
		})())))
		var _source5 D1 = _239_v119
		_ = _source5
		if _source5.Is_DC2() {
			var _240___mcc_h31 _dafny.CodePoint = _source5.Get_().(D1_DC2).Cf5
			_ = _240___mcc_h31
			var _241_cf5 _dafny.CodePoint = _240___mcc_h31
			_ = _241_cf5
			var _242_v120 _dafny.Set
			_ = _242_v120
			_242_v120 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(124))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_243_cf5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_244_i9 _dafny.Int) _dafny.CodePoint {
					return _243_cf5
				}
			})(_241_cf5))))
			var _245_v122 _dafny.Set
			_ = _245_v122
			_245_v122 = _dafny.SetOf(!((_232_v112).F17()), false, p3, (_232_v112).F17(), (Companion_Default___.Fm13(globalState)).Dtor_cf10())
			var _246_v124 _dafny.Sequence
			_ = _246_v124
			_246_v124 = _dafny.SeqOf(_238_v118, func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(92), _dafny.IntOfInt64(86))); ; {
					_compr_21, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _247_v123 _dafny.Int
					_247_v123 = interface{}(_compr_21).(_dafny.Int)
					if ((_dafny.IntOfInt64(92)).Cmp(_247_v123) <= 0) && ((_247_v123).Cmp(_dafny.IntOfInt64(86)) < 0) {
						_coll21.Add((_247_v123).Times(_dafny.IntOfInt64(2)), p2)
					}
				}
				return _coll21.ToMap()
			}())
			var _248_v125 _dafny.Array
			_ = _248_v125
			var _nwElement0_4 bool = p2
			_ = _nwElement0_4
			var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(15))
			_ = _nw34
			(_nw34).ArraySet1(_nwElement0_4, 0)
			(_nw34).ArraySet1(p3, 1)
			(_nw34).ArraySet1(!(p3), 2)
			(_nw34).ArraySet1((_232_v112).F17(), 3)
			(_nw34).ArraySet1(!(_32_v0), 4)
			(_nw34).ArraySet1((func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate((_242_v120).Elements()); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _249_v121 _dafny.Sequence
					_249_v121 = interface{}(_compr_22).(_dafny.Sequence)
					if (_242_v120).Contains(_249_v121) {
						_coll22.Add(_249_v121)
					}
				}
				return _coll22.ToSet()
			}()).IsProperSubsetOf(_242_v120), 5)
			(_nw34).ArraySet1((_232_v112).F17(), 6)
			(_nw34).ArraySet1(p3, 7)
			(_nw34).ArraySet1(((_245_v122).Cardinality()).Cmp(p1) > 0, 8)
			(_nw34).ArraySet1(p2, 9)
			(_nw34).ArraySet1(p2, 10)
			(_nw34).ArraySet1(p3, 11)
			(_nw34).ArraySet1(p3, 12)
			(_nw34).ArraySet1((_232_v112).F17(), 13)
			(_nw34).ArraySet1(_dafny.Companion_Sequence_.Contains(_246_v124, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)), 14)
			_248_v125 = _nw34
			var _250_v126 _dafny.MultiSet
			_ = _250_v126
			_250_v126 = _dafny.MultiSetOf(p1)
			var _251_v127 _dafny.Map
			_ = _251_v127
			_251_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _250_v126)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_248_v125), 0))
			_ = _index32
			(_248_v125).ArraySet1((func() bool {
				if (_232_v112).F17() {
					return (Companion_D1_.Create_DC4_(p1, p2, _251_v127)).Dtor_cf10()
				}
				return (_232_v112).F17()
			})(), (_index32).Int())
			var _252_v128 _dafny.Set
			_ = _252_v128
			_252_v128 = _dafny.SetOf(p1)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_248_v125), 0))
			_ = _index33
			(_248_v125).ArraySet1(((_252_v128).IsProperSubsetOf(_252_v128)) && (p3), (_index33).Int())
			var _253_v129 *C0
			_ = _253_v129
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__((p1).Cmp(p1) > 0)
			_253_v129 = _nw35
			var _254_v131 _dafny.Array
			_ = _254_v131
			var _nwElement0_5 _dafny.Int = p1
			_ = _nwElement0_5
			var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(27))
			_ = _nw36
			(_nw36).ArraySet1(_nwElement0_5, 0)
			(_nw36).ArraySet1(p1, 1)
			(_nw36).ArraySet1(p1, 2)
			(_nw36).ArraySet1(p1, 3)
			(_nw36).ArraySet1(p1, 4)
			(_nw36).ArraySet1(p1, 5)
			(_nw36).ArraySet1(p1, 6)
			(_nw36).ArraySet1(p1, 7)
			(_nw36).ArraySet1(p1, 8)
			(_nw36).ArraySet1(p1, 9)
			(_nw36).ArraySet1((_dafny.Zero).Minus(p1), 10)
			(_nw36).ArraySet1((func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter23 := _dafny.Iterate((_238_v118).Keys().Elements()); ; {
					_compr_23, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _255_v130 _dafny.Int
					_255_v130 = interface{}(_compr_23).(_dafny.Int)
					if (_238_v118).Contains(_255_v130) {
						_coll23.Add((_255_v130).Plus(p1), (_232_v112).F17())
					}
				}
				return _coll23.ToMap()
			}()).Cardinality(), 11)
			(_nw36).ArraySet1(p1, 12)
			(_nw36).ArraySet1(_dafny.IntOfUint32((_35_v3).Cardinality()), 13)
			(_nw36).ArraySet1(p1, 14)
			(_nw36).ArraySet1(p1, 15)
			(_nw36).ArraySet1(p1, 16)
			(_nw36).ArraySet1(p1, 17)
			(_nw36).ArraySet1((_250_v126).Cardinality(), 18)
			(_nw36).ArraySet1((_dafny.Zero).Minus(p1), 19)
			(_nw36).ArraySet1(_dafny.IntOfUint32((_35_v3).Cardinality()), 20)
			(_nw36).ArraySet1(p1, 21)
			(_nw36).ArraySet1(p1, 22)
			(_nw36).ArraySet1(p1, 23)
			(_nw36).ArraySet1(_dafny.IntOfInt64(930), 24)
			(_nw36).ArraySet1((_dafny.Zero).Minus(p1), 25)
			(_nw36).ArraySet1(p1, 26)
			_254_v131 = _nw36
			var _256_v132 _dafny.Sequence
			_ = _256_v132
			_256_v132 = _dafny.SeqOf(_254_v131, _254_v131)
			var _257_v133 _dafny.Array
			_ = _257_v133
			var _nwElement0_6 _dafny.Sequence = _256_v132
			_ = _nwElement0_6
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(12))
			_ = _nw37
			(_nw37).ArraySet1(_nwElement0_6, 0)
			(_nw37).ArraySet1(_256_v132, 1)
			(_nw37).ArraySet1(_dafny.SeqOf(_254_v131), 2)
			(_nw37).ArraySet1(_256_v132, 3)
			(_nw37).ArraySet1(_256_v132, 4)
			(_nw37).ArraySet1(_256_v132, 5)
			(_nw37).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_254_v131), _256_v132), 6)
			(_nw37).ArraySet1(_256_v132, 7)
			(_nw37).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_256_v132, _256_v132), 8)
			(_nw37).ArraySet1((func() _dafny.Sequence {
				if p2 {
					return _dafny.SeqOf(_254_v131)
				}
				return _256_v132
			})(), 9)
			(_nw37).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_256_v132, _256_v132), 10)
			(_nw37).ArraySet1(_256_v132, 11)
			_257_v133 = _nw37
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_257_v133), 0))
			_ = _index34
			(_257_v133).ArraySet1(_256_v132, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_257_v133), 0))
			_ = _index35
			var _rhs54 _dafny.Sequence = _256_v132
			_ = _rhs54
			var _rhs55 _dafny.Int = (((_245_v122).Cardinality()).Plus(p1)).Times(p1)
			_ = _rhs55
			var _rhs56 bool = !(!((p2) || ((_232_v112).F17())))
			_ = _rhs56
			var _rhs57 _dafny.Int = p1
			_ = _rhs57
			var _rhs58 _dafny.Array = _254_v131
			_ = _rhs58
			var _lhs35 _dafny.Array = _257_v133
			_ = _lhs35
			var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_257_v133), 0))
			_ = _lhs36
			var _lhs37 *GlobalState = globalState
			_ = _lhs37
			var _lhs38 *GlobalState = globalState
			_ = _lhs38
			(_lhs35).ArraySet1(_rhs54, (_lhs36).Int())
			_lhs37.F11 = _rhs55
			_32_v0 = _rhs56
			_lhs38.F11 = _rhs57
			_254_v131 = _rhs58
			_32_v0 = (_248_v125).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_248_v125), 0))).Int()).(bool)
		} else if _source5.Is_DC3() {
			var _258___mcc_h32 _dafny.Map = _source5.Get_().(D1_DC3).Cf6
			_ = _258___mcc_h32
			var _259___mcc_h33 _dafny.Int = _source5.Get_().(D1_DC3).Cf7
			_ = _259___mcc_h33
			var _260___mcc_h34 D0 = _source5.Get_().(D1_DC3).Cf8
			_ = _260___mcc_h34
			var _261_cf8 D0 = _260___mcc_h34
			_ = _261_cf8
			var _262_cf7 _dafny.Int = _259___mcc_h33
			_ = _262_cf7
			var _263_cf6 _dafny.Map = _258___mcc_h32
			_ = _263_cf6
			(globalState).F16 = (func() _dafny.Int {
				if _32_v0 {
					return p1
				}
				return _262_cf7
			})()
			var _264_v134 _dafny.Map
			_ = _264_v134
			_264_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_v1, !(p3))
			var _265_v135 _dafny.Set
			_ = _265_v135
			_265_v135 = _dafny.SetOf((_232_v112).F17(), true)
			_264_v134 = Companion_Default___.Fm15(_32_v0, _265_v135, !(p2), globalState)
			var _266_v136 D3
			_ = _266_v136
			_266_v136 = Companion_D3_.Create_DC10_(_235_v115)
			_235_v115 = (_266_v136).Dtor_cf18()
			_236_v116 = (_236_v116).Update(_262_cf7, (_262_cf7).Times(Companion_Default___.Fm0(p1, _263_cf6, globalState)))
		} else if _source5.Is_DC4() {
			var _267___mcc_h35 _dafny.Int = _source5.Get_().(D1_DC4).Cf9
			_ = _267___mcc_h35
			var _268___mcc_h36 bool = _source5.Get_().(D1_DC4).Cf10
			_ = _268___mcc_h36
			var _269___mcc_h37 _dafny.Map = _source5.Get_().(D1_DC4).Cf11
			_ = _269___mcc_h37
			var _270_cf11 _dafny.Map = _269___mcc_h37
			_ = _270_cf11
			var _271_cf10 bool = _268___mcc_h36
			_ = _271_cf10
			var _272_cf9 _dafny.Int = _267___mcc_h35
			_ = _272_cf9
			var _273_v137 _dafny.Array
			_ = _273_v137
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
			_ = _nw38
			_273_v137 = _nw38
			var _274_v138 _dafny.Set
			_ = _274_v138
			_274_v138 = _dafny.SetOf(_232_v112, _232_v112)
			var _275_v139 _dafny.Array
			_ = _275_v139
			var _nwElement0_7 _dafny.Int = _dafny.IntOfInt64(-142)
			_ = _nwElement0_7
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(4))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_7, 0)
			(_nw39).ArraySet1(_dafny.IntOfInt64(-626), 1)
			(_nw39).ArraySet1(p1, 2)
			(_nw39).ArraySet1((_274_v138).Cardinality(), 3)
			_275_v139 = _nw39
			var _276_v140 _dafny.MultiSet
			_ = _276_v140
			_276_v140 = _dafny.MultiSetOf(_275_v139)
			var _277_v141 _dafny.Sequence
			_ = _277_v141
			_277_v141 = _dafny.SeqOf(Companion_Default___.Fm0((_276_v140).Cardinality(), _236_v116, globalState))
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_273_v137), 0))
			_ = _index36
			(_273_v137).ArraySet1(_277_v141, (_index36).Int())
			var _278_v142 D1
			_ = _278_v142
			_278_v142 = Companion_D1_.Create_DC2_(_33_v1)
			var _279_v143 _dafny.Sequence
			_ = _279_v143
			_279_v143 = _dafny.SeqOf(_278_v142, _278_v142, _278_v142)
			var _280_v145 _dafny.MultiSet
			_ = _280_v145
			_280_v145 = _dafny.MultiSetOf(func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(998), _dafny.IntOfInt64(785))); ; {
					_compr_24, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _281_v144 _dafny.Int
					_281_v144 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(998)).Cmp(_281_v144) <= 0) && ((_281_v144).Cmp(_dafny.IntOfInt64(785)) < 0) {
						_coll24.Add((_281_v144).Minus(p1), _272_cf9)
					}
				}
				return _coll24.ToMap()
			}(), _236_v116, Companion_Default___.Fm1(_236_v116, _272_cf9, _272_cf9, globalState), _236_v116)
			var _282_v146 _dafny.Sequence
			_ = _282_v146
			_282_v146 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(541))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_283_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_284_i10 _dafny.Int) _dafny.Int {
					return _283_cf9
				}
			})(_272_cf9))), _277_v141, _dafny.SeqOf(p1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(454), _dafny.IntOfUint32((_279_v143).Cardinality()), p1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_35_v3).Cardinality())))).Cardinality()), p1), _277_v141, _dafny.Companion_Sequence_.Update(_dafny.SeqOf((func() _dafny.Int {
				if (_280_v145).Contains(_236_v116) {
					return (_280_v145).Multiplicity(_236_v116)
				}
				return _272_cf9
			})()), (Companion_Default___.SafeIndex(_272_cf9, _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Int {
				if (_280_v145).Contains(_236_v116) {
					return (_280_v145).Multiplicity(_236_v116)
				}
				return _272_cf9
			})())).Cardinality()))).Uint32(), p1))
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_273_v137), 0))
			_ = _index37
			(_273_v137).ArraySet1((_282_v146).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_282_v146).Cardinality()))).Uint32()).(_dafny.Sequence), (_index37).Int())
			var _285_v147 *C0
			_ = _285_v147
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__(_271_cf10)
			_285_v147 = _nw40
			var _286_v148 D4
			_ = _286_v148
			_286_v148 = Companion_D4_.Create_DC12_(_234_v114)
			_234_v114 = (_234_v114).Merge((_286_v148).Dtor_cf19())
			var _287_v149 _dafny.Array
			_ = _287_v149
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw41
			_287_v149 = _nw41
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_287_v149), 0))
			_ = _index38
			(_287_v149).ArraySet1(Companion_Default___.Fm4(Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ns")).Cardinality()), _236_v116, globalState), p2, globalState), (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_287_v149), 0))
			_ = _index39
			(_287_v149).ArraySet1(_271_cf10, (_index39).Int())
		} else if _source5.Is_DC1() {
			var _288___mcc_h38 _dafny.Sequence = _source5.Get_().(D1_DC1).Cf4
			_ = _288___mcc_h38
			var _289_cf4 _dafny.Sequence = _288___mcc_h38
			_ = _289_cf4
			_236_v116 = _236_v116
			var _290_v150 _dafny.MultiSet
			_ = _290_v150
			_290_v150 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hhlfxth")).Cardinality()))
			_290_v150 = _290_v150
			var _291_v151 _dafny.Sequence
			_ = _291_v151
			_291_v151 = _dafny.SeqOf(p1, p1)
			var _292_v152 _dafny.Set
			_ = _292_v152
			_292_v152 = _dafny.SetOf(p1, Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.SeqOf(p2, p3)).Cardinality()), _236_v116, globalState), p1, Companion_Default___.Fm0((_291_v151).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_291_v151).Cardinality()))).Uint32()).(_dafny.Int), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), globalState))
			var _293_v153 _dafny.Map
			_ = _293_v153
			_293_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _292_v152)
			var _294_v154 _dafny.Set
			_ = _294_v154
			_294_v154 = _dafny.SetOf((func() _dafny.Set {
				if (_293_v153).Contains(p1) {
					return (_293_v153).Get(p1).(_dafny.Set)
				}
				return _dafny.SetOf(p1, p1, p1)
			})(), _292_v152, _292_v152, _dafny.SetOf((_dafny.SetOf((_232_v112).F17())).Cardinality()))
			_32_v0 = (_294_v154).Contains(_292_v152)
			var _rhs59 bool = false
			_ = _rhs59
			var _rhs60 bool = p2
			_ = _rhs60
			_32_v0 = _rhs59
			_32_v0 = _rhs60
		} else {
			var _295___mcc_h39 D1 = _source5.Get_().(D1_DC5).Cf12
			_ = _295___mcc_h39
			var _296_cf12 D1 = _295___mcc_h39
			_ = _296_cf12
			var _297_v155 _dafny.Map
			_ = _297_v155
			_297_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _238_v118)
			var _pat_let_tv8 = _35_v3
			_ = _pat_let_tv8
			var _pat_let_tv9 = p1
			_ = _pat_let_tv9
			var _pat_let_tv10 = _297_v155
			_ = _pat_let_tv10
			var _298_v156 _dafny.Array
			_ = _298_v156
			var _nwElement0_8 D0 = func(_pat_let7_0 D0) D0 {
				return func(_299_dt__update__tmp_h3 D0) D0 {
					return func(_pat_let8_0 _dafny.Sequence) D0 {
						return func(_300_dt__update_hcf2_h1 _dafny.Sequence) D0 {
							return func(_pat_let9_0 _dafny.Int) D0 {
								return func(_301_dt__update_hcf0_h0 _dafny.Int) D0 {
									return func(_pat_let10_0 _dafny.Map) D0 {
										return func(_302_dt__update_hcf3_h0 _dafny.Map) D0 {
											return Companion_D0_.Create_DC0_(_301_dt__update_hcf0_h0, (_299_dt__update__tmp_h3).Dtor_cf1(), _300_dt__update_hcf2_h1, _302_dt__update_hcf3_h0)
										}(_pat_let10_0)
									}(_pat_let_tv10)
								}(_pat_let9_0)
							}(_pat_let_tv9)
						}(_pat_let8_0)
					}(_pat_let_tv8)
				}(_pat_let7_0)
			}(_36_v4)
			_ = _nwElement0_8
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(4))
			_ = _nw42
			(_nw42).ArraySet1(_nwElement0_8, 0)
			(_nw42).ArraySet1(_36_v4, 1)
			(_nw42).ArraySet1(_36_v4, 2)
			(_nw42).ArraySet1(_36_v4, 3)
			_298_v156 = _nw42
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_298_v156), 0))
			_ = _index40
			(_298_v156).ArraySet1(_36_v4, (_index40).Int())
			var _303_v158 _dafny.Map
			_ = _303_v158
			_303_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC11_(), (_232_v112).F17())
			var _pat_let_tv11 = _231_cf4
			_ = _pat_let_tv11
			var _pat_let_tv12 = p2
			_ = _pat_let_tv12
			var _pat_let_tv13 = globalState
			_ = _pat_let_tv13
			var _pat_let_tv14 = p1
			_ = _pat_let_tv14
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_298_v156), 0))
			_ = _index41
			(_298_v156).ArraySet1(func(_pat_let11_0 D0) D0 {
				return func(_305_dt__update__tmp_h4 D0) D0 {
					return func(_pat_let12_0 _dafny.Map) D0 {
						return func(_306_dt__update_hcf3_h1 _dafny.Map) D0 {
							return func(_pat_let13_0 _dafny.Int) D0 {
								return func(_307_dt__update_hcf0_h1 _dafny.Int) D0 {
									return Companion_D0_.Create_DC0_(_307_dt__update_hcf0_h1, (_305_dt__update__tmp_h4).Dtor_cf1(), (_305_dt__update__tmp_h4).Dtor_cf2(), _306_dt__update_hcf3_h1)
								}(_pat_let13_0)
							}(_pat_let_tv14)
						}(_pat_let12_0)
					}(Companion_Default___.Fm11(_dafny.IntOfUint32((_pat_let_tv11).Cardinality()), _pat_let_tv12, _pat_let_tv13))
				}(_pat_let11_0)
			}(Companion_Default___.Fm16(_33_v1, Companion_Default___.Fm17(_dafny.IntOfInt64(597), globalState), func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter25 := _dafny.Iterate((_303_v158).Keys().Elements()); ; {
					_compr_25, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _304_v157 D3
					_304_v157 = interface{}(_compr_25).(D3)
					if (_303_v158).Contains(_304_v157) {
						_coll25.Add(_304_v157, p1)
					}
				}
				return _coll25.ToMap()
			}(), globalState)), (_index41).Int())
			var _308_v159 *C0
			_ = _308_v159
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__((_232_v112).F17())
			_308_v159 = _nw43
			var _309_v160 _dafny.Array
			_ = _309_v160
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw44
			_309_v160 = _nw44
			var _310_v161 _dafny.Sequence
			_ = _310_v161
			_310_v161 = _dafny.SeqOf((func() _dafny.Array {
				if p3 {
					return _309_v160
				}
				return _309_v160
			})(), _309_v160, _309_v160, _309_v160)
			_310_v161 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_309_v160), _310_v161)
			_32_v0 = _32_v0
		}
	} else {
		var _311___mcc_h8 D1 = _source1.Get_().(D1_DC5).Cf12
		_ = _311___mcc_h8
		var _312_cf12 D1 = _311___mcc_h8
		_ = _312_cf12
		if _32_v0 {
			var _313_v162 _dafny.MultiSet
			_ = _313_v162
			_313_v162 = _dafny.MultiSetOf(p2, !(!(p3)))
			var _314_v163 *C0
			_ = _314_v163
			var _nw45 *C0 = New_C0_()
			_ = _nw45
			_nw45.Ctor__(p2)
			_314_v163 = _nw45
			var _315_v164 _dafny.Map
			_ = _315_v164
			_315_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_313_v162).Intersection((_313_v162).Update(_32_v0, Companion_Default___.Abs(p1))), _314_v163)
			_315_v164 = (_315_v164).Update(_313_v162, _314_v163)
			_32_v0 = p3
			var _316_v165 _dafny.Array
			_ = _316_v165
			var _nw46 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
			_ = _nw46
			_316_v165 = _nw46
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_316_v165), 0))
			_ = _index42
			(_316_v165).ArraySet1(_314_v163, (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_316_v165), 0))
			_ = _index43
			(_316_v165).ArraySet1(_314_v163, (_index43).Int())
			var _317_v166 _dafny.Map
			_ = _317_v166
			_317_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, true)
			_317_v166 = (func() _dafny.Map {
				if (_314_v163).F17() {
					return (_317_v166).Update((_314_v163).F17(), (_314_v163).F17())
				}
				return (_317_v166).Merge(_317_v166)
			})()
			(globalState).F11 = p1
		} else {
			var _rhs61 _dafny.Int = p1
			_ = _rhs61
			var _rhs62 bool = !(p2)
			_ = _rhs62
			var _lhs39 *GlobalState = globalState
			_ = _lhs39
			_lhs39.F6 = _rhs61
			_32_v0 = _rhs62
			var _318_v167 *C0
			_ = _318_v167
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__(p3)
			_318_v167 = _nw47
			var _319_v168 _dafny.Array
			_ = _319_v168
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_4
			var _nw48 _dafny.Array
			_ = _nw48
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw48 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = func(_320_i11 _dafny.Int) bool {
					return true
				}
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw48 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw48).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw48).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_319_v168 = _nw48
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_319_v168), 0))
			_ = _index44
			(_319_v168).ArraySet1((func() bool {
				if (_318_v167).F17() {
					return (_318_v167).F17()
				}
				return _32_v0
			})(), (_index44).Int())
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_319_v168), 0))
			_ = _index45
			(_319_v168).ArraySet1(p2, (_index45).Int())
			_32_v0 = Companion_Default___.Fm4(p1, (_319_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_319_v168), 0))).Int()).(bool), globalState)
			var _321_v169 _dafny.Map
			_ = _321_v169
			_321_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, p1)
			var _322_v170 _dafny.Set
			_ = _322_v170
			_322_v170 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_319_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_319_v168), 0))).Int()).(bool), (_319_v168).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_319_v168), 0))).Int()).(bool))).Cardinality()), (func() _dafny.Int {
				if (_321_v169).Contains(p2) {
					return (_321_v169).Get(p2).(_dafny.Int)
				}
				return p1
			})(), (_dafny.Zero).Minus(p1))
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_319_v168), 0))
			_ = _index46
			(_319_v168).ArraySet1(!((_dafny.SetOf(p1)).IsDisjointFrom(_322_v170)), (_index46).Int())
		}
		var _323_v171 _dafny.MultiSet
		_ = _323_v171
		_323_v171 = _dafny.MultiSetOf(_32_v0, true)
		if Companion_Default___.Fm4(p1, ((_323_v171).Cardinality()).Cmp(p1) != 0, globalState) {
			var _324_v172 _dafny.Array
			_ = _324_v172
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
			_ = _nw49
			_324_v172 = _nw49
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_324_v172), 0))
			_ = _index47
			(_324_v172).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kinwdfjx"), _35_v3), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_324_v172), 0))
			_ = _index48
			(_324_v172).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}(func(_325_i12 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})), (_index48).Int())
			var _326_v173 *C0
			_ = _326_v173
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__((_323_v171).IsProperSubsetOf(_323_v171))
			_326_v173 = _nw50
			var _327_v174 _dafny.Map
			_ = _327_v174
			_327_v174 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(522), p1)
			var _328_v175 _dafny.Map
			_ = _328_v175
			_328_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
			(globalState).F16 = (Companion_Default___.Fm0(p1, _327_v174, globalState)).Plus(((func() _dafny.Map {
				if p3 {
					return _328_v175
				}
				return Companion_Default___.Fm5(globalState)
			})()).Cardinality())
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_324_v172), 0))
			_ = _index49
			(_324_v172).ArraySet1((_324_v172).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_324_v172), 0))).Int()).(_dafny.Sequence), (_index49).Int())
			var _329_v176 _dafny.Array
			_ = _329_v176
			var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(12))
			_ = _nw51
			_329_v176 = _nw51
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_329_v176), 0))
			_ = _index50
			(_329_v176).ArraySet1((_323_v171).Update(_32_v0, Companion_Default___.Abs(p1)), (_index50).Int())
			var _330_v177 _dafny.Array
			_ = _330_v177
			var _nw52 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(20))
			_ = _nw52
			_330_v177 = _nw52
			var _331_v178 _dafny.Array
			_ = _331_v178
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_5
			var _nw53 _dafny.Array
			_ = _nw53
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw53 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.CodePoint = (func(_332_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_333_i13 _dafny.Int) _dafny.CodePoint {
						return _332_v1
					}
				})(_33_v1)
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
			_331_v178 = _nw53
			var _334_v179 _dafny.Set
			_ = _334_v179
			_334_v179 = _dafny.SetOf(_331_v178)
			var _335_v180 D4
			_ = _335_v180
			_335_v180 = Companion_D4_.Create_DC15_(_334_v179, _33_v1)
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_330_v177), 0))
			_ = _index51
			(_330_v177).ArraySet1(_335_v180, (_index51).Int())
			var _336_v181 _dafny.Array
			_ = _336_v181
			var _nw54 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw54
			_336_v181 = _nw54
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_329_v176), 0))
			_ = _index52
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_330_v177), 0))
			_ = _index53
			var _rhs63 _dafny.MultiSet = (func() _dafny.MultiSet {
				if p2 {
					return Companion_Default___.Fm6(p2, p3, p2, globalState)
				}
				return _323_v171
			})()
			_ = _rhs63
			var _rhs64 D4 = Companion_D4_.Create_DC15_(_334_v179, _33_v1)
			_ = _rhs64
			var _rhs65 bool = (Companion_D4_.Create_DC13_(Companion_Default___.Fm4(p1, p3, globalState), p1, _336_v181)).Dtor_cf20()
			_ = _rhs65
			var _lhs40 _dafny.Array = _329_v176
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_329_v176), 0))
			_ = _lhs41
			var _lhs42 _dafny.Array = _330_v177
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_330_v177), 0))
			_ = _lhs43
			(_lhs40).ArraySet1(_rhs63, (_lhs41).Int())
			(_lhs42).ArraySet1(_rhs64, (_lhs43).Int())
			_32_v0 = _rhs65
		} else {
			var _337_v182 _dafny.Sequence
			_ = _337_v182
			_337_v182 = _dafny.SeqOf(!(p2))
			_32_v0 = (_337_v182).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-856))), _dafny.IntOfUint32((_337_v182).Cardinality()))).Uint32()).(bool)
			var _338_v183 _dafny.Map
			_ = _338_v183
			_338_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _339_v184 *C0
			_ = _339_v184
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__(p2)
			_339_v184 = _nw55
			var _340_v185 _dafny.Map
			_ = _340_v185
			_340_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _339_v184)
			var _341_v186 _dafny.Map
			_ = _341_v186
			_341_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(p1, _338_v183, globalState), (_340_v185).Merge(_340_v185))
			var _342_v187 _dafny.Sequence
			_ = _342_v187
			_342_v187 = _dafny.SeqOf(_341_v186)
			_341_v186 = (_341_v186).Merge((_342_v187).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Cardinality(), _dafny.IntOfUint32((_342_v187).Cardinality()))).Uint32()).(_dafny.Map))
			var _343_v188 _dafny.Map
			_ = _343_v188
			_343_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_339_v184).F17(), _33_v1)
			var _344_v189 _dafny.Set
			_ = _344_v189
			_344_v189 = _dafny.SetOf(_35_v3)
			var _345_v190 D4
			_ = _345_v190
			_345_v190 = Companion_D4_.Create_DC14_((_343_v188).Cardinality(), ((_338_v183).Cardinality()).Minus(p1), _344_v189)
			_345_v190 = _345_v190
			var _346_v191 _dafny.Array
			_ = _346_v191
			var _nw56 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw56
			_346_v191 = _nw56
			_346_v191 = _346_v191
			_32_v0 = !(!(p3) || (_32_v0))
		}
		var _347_v192 _dafny.Array
		_ = _347_v192
		var _len0_6 _dafny.Int = _dafny.One
		_ = _len0_6
		var _nw57 _dafny.Array
		_ = _nw57
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw57 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) D4 = (func(_348_v171 _dafny.MultiSet, _349_v3 _dafny.Sequence) func(_dafny.Int) D4 {
				return func(_350_i14 _dafny.Int) D4 {
					return Companion_D4_.Create_DC14_((_348_v171).Cardinality(), _dafny.IntOfInt64(996), _dafny.SetOf(_349_v3, _349_v3))
				}
			})(_323_v171, _35_v3)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw57 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw57).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw57).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_347_v192 = _nw57
		var _351_v193 *C0
		_ = _351_v193
		var _nw58 *C0 = New_C0_()
		_ = _nw58
		_nw58.Ctor__(!(p3))
		_351_v193 = _nw58
		var _352_v194 D4
		_ = _352_v194
		_352_v194 = Companion_D4_.Create_DC14_(_dafny.IntOfUint32((_35_v3).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_351_v193)).Cardinality()), Companion_Default___.Fm18(p1, p1, p1, _33_v1, globalState))
		var _353_v195 _dafny.Set
		_ = _353_v195
		_353_v195 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("fpagmi"))
		var _pat_let_tv15 = _353_v195
		_ = _pat_let_tv15
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_347_v192), 0))
		_ = _index54
		(_347_v192).ArraySet1(func(_pat_let14_0 D4) D4 {
			return func(_354_dt__update__tmp_h5 D4) D4 {
				return func(_pat_let15_0 _dafny.Set) D4 {
					return func(_355_dt__update_hcf25_h0 _dafny.Set) D4 {
						return Companion_D4_.Create_DC14_((_354_dt__update__tmp_h5).Dtor_cf23(), (_354_dt__update__tmp_h5).Dtor_cf24(), _355_dt__update_hcf25_h0)
					}(_pat_let15_0)
				}(_pat_let_tv15)
			}(_pat_let14_0)
		}(_352_v194), (_index54).Int())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_347_v192), 0))
		_ = _index55
		(_347_v192).ArraySet1(Companion_D4_.Create_DC14_(p1, p1, _dafny.SetOf(_35_v3)), (_index55).Int())
		var _356_v196 _dafny.Sequence
		_ = _356_v196
		_356_v196 = _dafny.SeqOf(_32_v0)
		var _357_v197 _dafny.Map
		_ = _357_v197
		_357_v197 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfUint32((_356_v196).Cardinality()))
		var _358_v198 _dafny.Map
		_ = _358_v198
		_358_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, _357_v197)
		_358_v198 = Companion_Default___.Fm19((_dafny.MultiSetOf(_32_v0, p3, (_351_v193).F17(), p2)).Difference(_323_v171), globalState)
	}
	var _359_v199 *C0
	_ = _359_v199
	var _nw59 *C0 = New_C0_()
	_ = _nw59
	_nw59.Ctor__(p2)
	_359_v199 = _nw59
	var _360_i15 _dafny.Int
	_ = _360_i15
	_360_i15 = _dafny.Zero
	{
		for p3 {
			{
				if (_360_i15).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_360_i15 = (_360_i15).Plus(_dafny.One)
				var _361_v200 _dafny.Array
				_ = _361_v200
				var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw60
				_361_v200 = _nw60
				var _362_v201 _dafny.Map
				_ = _362_v201
				_362_v201 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_361_v200, _359_v199)
				_362_v201 = (_362_v201).Update(_361_v200, _359_v199)
				(globalState).F16 = p1
				if (_359_v199).F17() {
					_32_v0 = p3
					_33_v1 = _33_v1
					var _363_v202 _dafny.Array
					_ = _363_v202
					var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.One)
					_ = _nw61
					_363_v202 = _nw61
					var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_363_v202), 0))
					_ = _index56
					(_363_v202).ArraySet1(_dafny.SetOf(p3, Companion_Default___.Fm4(p1, (_359_v199).F17(), globalState)), (_index56).Int())
					var _364_v203 _dafny.Set
					_ = _364_v203
					_364_v203 = _dafny.SetOf(!((_359_v199).F17()))
					var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_363_v202), 0))
					_ = _index57
					(_363_v202).ArraySet1(_364_v203, (_index57).Int())
					var _nw62 *C0 = New_C0_()
					_ = _nw62
					_nw62.Ctor__((_359_v199).F17())
					_359_v199 = _nw62
					var _365_v204 _dafny.Array
					_ = _365_v204
					var _nw63 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
					_ = _nw63
					_365_v204 = _nw63
					var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_365_v204), 0))
					_ = _index58
					(_365_v204).ArraySet1(_359_v199, (_index58).Int())
					var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_365_v204), 0))
					_ = _index59
					var _rhs66 bool = _32_v0
					_ = _rhs66
					var _rhs67 _dafny.Sequence = _35_v3
					_ = _rhs67
					var _rhs68 _dafny.Int = _dafny.IntOfUint32((_35_v3).Cardinality())
					_ = _rhs68
					var _rhs69 *C0 = (func() *C0 {
						if (_359_v199).F17() {
							return _359_v199
						}
						return _359_v199
					})()
					_ = _rhs69
					var _lhs44 *GlobalState = globalState
					_ = _lhs44
					var _lhs45 _dafny.Array = _365_v204
					_ = _lhs45
					var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_365_v204), 0))
					_ = _lhs46
					_32_v0 = _rhs66
					_35_v3 = _rhs67
					_lhs44.F16 = _rhs68
					(_lhs45).ArraySet1(_rhs69, (_lhs46).Int())
				} else {
					var _366_v205 D2
					_ = _366_v205
					_366_v205 = Companion_D2_.Create_DC7_(p1, _33_v1)
					_366_v205 = _366_v205
					var _367_v206 _dafny.Array
					_ = _367_v206
					var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
					_ = _nw64
					_367_v206 = _nw64
					var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_361_v200), 0))
					_ = _index60
					(_361_v200).ArraySet1(_367_v206, (_index60).Int())
					var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_361_v200), 0))
					_ = _index61
					(_361_v200).ArraySet1(_367_v206, (_index61).Int())
					var _368_v207 _dafny.Map
					_ = _368_v207
					_368_v207 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_359_v199, _33_v1)
					_32_v0 = Companion_Default___.Fm4((_dafny.Zero).Minus(p1), !((_368_v207).Update(_359_v199, _33_v1)).Contains(_359_v199), globalState)
					_32_v0 = (func() bool {
						if (p1).Cmp((Companion_Default___.Fm20(p3, globalState)).Cardinality()) == 0 {
							return (_359_v199).F17()
						}
						return p2
					})()
					var _369_v208 _dafny.MultiSet
					_ = _369_v208
					_369_v208 = _dafny.MultiSetOf(p1)
					var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((p0), 0))
					_ = _index62
					(p0).ArraySet1(_369_v208, (_index62).Int())
					var _arr0 _dafny.Array = _dafny.ArrayCastTo((_361_v200).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_361_v200), 0))).Int()))
					_ = _arr0
					var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_dafny.ArrayCastTo((_361_v200).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_361_v200), 0))).Int()))), 0))
					_ = _index63
					(_arr0).ArraySet1(p1, (_index63).Int())
					var _370_v209 _dafny.Array
					_ = _370_v209
					var _nwElement0_9 bool = (_359_v199).F17()
					_ = _nwElement0_9
					var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.One)
					_ = _nw65
					(_nw65).ArraySet1(_nwElement0_9, 0)
					_370_v209 = _nw65
					var _371_v210 _dafny.Sequence
					_ = _371_v210
					_371_v210 = _dafny.SeqOf((_359_v199).F17(), false, (_359_v199).F17())
					var _372_v211 _dafny.Set
					_ = _372_v211
					_372_v211 = _dafny.SetOf(_dafny.IntOfUint32((_371_v210).Cardinality()), p1)
					var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_370_v209), 0))
					_ = _index64
					(_370_v209).ArraySet1((_dafny.SetOf(p1)).IsProperSubsetOf(_372_v211), (_index64).Int())
					var _373_v212 _dafny.Sequence
					_ = _373_v212
					_373_v212 = _dafny.SeqOf((_dafny.MultiSetOf(p1)).Difference(_369_v208), _369_v208, _369_v208, (_369_v208).Difference(_dafny.MultiSetOf(p1, p1)))
					var _374_v213 _dafny.Map
					_ = _374_v213
					_374_v213 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _359_v199)
					var _375_v214 _dafny.Map
					_ = _375_v214
					_375_v214 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _374_v213)
					var _376_v215 _dafny.Set
					_ = _376_v215
					_376_v215 = _dafny.SetOf(false)
					var _377_v216 _dafny.Map
					_ = _377_v216
					_377_v216 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
					var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((p0), 0))
					_ = _index65
					var _arr1 _dafny.Array = _dafny.ArrayCastTo((_361_v200).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_361_v200), 0))).Int()))
					_ = _arr1
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_dafny.ArrayCastTo((_361_v200).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_361_v200), 0))).Int()))), 0))
					_ = _index66
					var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_370_v209), 0))
					_ = _index67
					var _rhs70 bool = !(!((_32_v0) || ((_359_v199).F17())))
					_ = _rhs70
					var _rhs71 _dafny.MultiSet = (_373_v212).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_373_v212).Cardinality()))).Uint32()).(_dafny.MultiSet)
					_ = _rhs71
					var _rhs72 _dafny.Int = p1
					_ = _rhs72
					var _rhs73 _dafny.Int = (func() _dafny.Int {
						if (_369_v208).Contains((_375_v214).Cardinality()) {
							return (_369_v208).Multiplicity((_375_v214).Cardinality())
						}
						return Companion_Default___.Fm0((_376_v215).Cardinality(), _377_v216, globalState)
					})()
					_ = _rhs73
					var _rhs74 bool = false
					_ = _rhs74
					var _lhs47 _dafny.Array = p0
					_ = _lhs47
					var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((p0), 0))
					_ = _lhs48
					var _lhs49 _dafny.Array = _dafny.ArrayCastTo((_361_v200).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_361_v200), 0))).Int()))
					_ = _lhs49
					var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_dafny.ArrayCastTo((_361_v200).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_361_v200), 0))).Int()))), 0))
					_ = _lhs50
					var _lhs51 *GlobalState = globalState
					_ = _lhs51
					var _lhs52 _dafny.Array = _370_v209
					_ = _lhs52
					var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_370_v209), 0))
					_ = _lhs53
					_32_v0 = _rhs70
					(_lhs47).ArraySet1(_rhs71, (_lhs48).Int())
					(_lhs49).ArraySet1(_rhs72, (_lhs50).Int())
					_lhs51.F6 = _rhs73
					(_lhs52).ArraySet1(_rhs74, (_lhs53).Int())
				}
				var _378_v217 _dafny.Array
				_ = _378_v217
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw66
				_378_v217 = _nw66
				var _379_v218 _dafny.Sequence
				_ = _379_v218
				_379_v218 = _dafny.SeqOf(p1)
				var _380_v219 _dafny.Sequence
				_ = _380_v219
				_380_v219 = _dafny.SeqOf(_379_v218)
				var _381_v220 _dafny.Map
				_ = _381_v220
				_381_v220 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v0, p1)
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_378_v217), 0))
				_ = _index68
				(_378_v217).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_380_v219).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_381_v220).Contains((_359_v199).F17()) {
						return (_381_v220).Get((_359_v199).F17()).(_dafny.Int)
					}
					return p1
				})(), _dafny.IntOfUint32((_380_v219).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf(p1))).Cardinality()), (_index68).Int())
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_378_v217), 0))
				_ = _index69
				(_378_v217).ArraySet1((p1).Plus((p1).Plus(p1)), (_index69).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _382_v221 _dafny.Map
	_ = _382_v221
	_382_v221 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(415))
	var _383_v222 _dafny.Sequence
	_ = _383_v222
	_383_v222 = _dafny.SeqOf(p3)
	r0 = ((Companion_Default___.Fm0(p1, _382_v221, globalState)).Plus(_dafny.IntOfInt64(786))).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_383_v222, _dafny.SeqOf(p3))).Cardinality()))
	var _384_v223 D1
	_ = _384_v223
	_384_v223 = Companion_D1_.Create_DC2_(_33_v1)
	var _385_v224 _dafny.Map
	_ = _385_v224
	_385_v224 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v223, _32_v0)
	var _386_v225 _dafny.MultiSet
	_ = _386_v225
	_386_v225 = _dafny.MultiSetOf((_383_v222).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_383_v222).Cardinality()))).Uint32()).(bool))
	var _387_v226 _dafny.Map
	_ = _387_v226
	_387_v226 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
	var _388_v227 _dafny.Array
	_ = _388_v227
	var _nwElement0_10 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_32_v0, (_359_v199).F17(), (_359_v199).F17(), p2, _32_v0)).Cardinality())
	_ = _nwElement0_10
	var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(27))
	_ = _nw67
	(_nw67).ArraySet1(_nwElement0_10, 0)
	(_nw67).ArraySet1((p1).Plus((_dafny.Zero).Minus(p1)), 1)
	(_nw67).ArraySet1((func() _dafny.Int {
		if p3 {
			return (_dafny.Zero).Minus((_dafny.Zero).Minus(p1))
		}
		return p1
	})(), 2)
	(_nw67).ArraySet1(p1, 3)
	(_nw67).ArraySet1(p1, 4)
	(_nw67).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(290), p1), 5)
	(_nw67).ArraySet1(p1, 6)
	(_nw67).ArraySet1(_dafny.IntOfInt64(752), 7)
	(_nw67).ArraySet1((func() _dafny.Int {
		if p2 {
			return p1
		}
		return (_385_v224).Cardinality()
	})(), 8)
	(_nw67).ArraySet1(p1, 9)
	(_nw67).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_35_v3).Cardinality())), 10)
	(_nw67).ArraySet1(p1, 11)
	(_nw67).ArraySet1((_dafny.IntOfInt64(64)).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(552))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}((func(_389_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_390_i16 _dafny.Int) _dafny.CodePoint {
			return _389_v1
		}
	})(_33_v1))), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(552))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg27 _dafny.Int) interface{} {
			return coer27(arg27)
		}
	}((func(_391_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_392_i16 _dafny.Int) _dafny.CodePoint {
			return _391_v1
		}
	})(_33_v1)))).Cardinality()))).Uint32(), _33_v1)).Cardinality())), 12)
	(_nw67).ArraySet1(Companion_Default___.SafeDivisionInt((_386_v225).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if (_387_v226).Contains(p2) {
			return (_387_v226).Get(p2).(_dafny.Int)
		}
		return (_382_v221).Cardinality()
	})(), _359_v199)).Cardinality()), 13)
	(_nw67).ArraySet1(p1, 14)
	(_nw67).ArraySet1(p1, 15)
	(_nw67).ArraySet1(p1, 16)
	(_nw67).ArraySet1(p1, 17)
	(_nw67).ArraySet1((p1).Plus((_dafny.Zero).Minus(p1)), 18)
	(_nw67).ArraySet1(p1, 19)
	(_nw67).ArraySet1(p1, 20)
	(_nw67).ArraySet1(p1, 21)
	(_nw67).ArraySet1(p1, 22)
	(_nw67).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(470))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}((func(_393_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_394_i17 _dafny.Int) _dafny.Int {
			return _393_p1
		}
	})(p1)))).Cardinality()), 23)
	(_nw67).ArraySet1(p1, 24)
	(_nw67).ArraySet1(p1, 25)
	(_nw67).ArraySet1(p1, 26)
	_388_v227 = _nw67
	r1 = _388_v227
	var _395_v228 _dafny.Set
	_ = _395_v228
	_395_v228 = _dafny.SetOf((func() bool {
		if _32_v0 {
			return (_359_v199).F17()
		}
		return _32_v0
	})(), (_359_v199).F17(), p3)
	r2 = _395_v228
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _396_v1 _dafny.CodePoint
	_ = _396_v1
	_396_v1 = _dafny.CodePoint('d')
	var _397_v2 _dafny.Int
	_ = _397_v2
	_397_v2 = _dafny.IntOfInt64(514)
	var _398_v4 bool
	_ = _398_v4
	_398_v4 = false
	var _399_v5 _dafny.Sequence
	_ = _399_v5
	_399_v5 = _dafny.SeqOf(_398_v4, true)
	var _400_v6 _dafny.Map
	_ = _400_v6
	_400_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_399_v5).Cardinality()), true)
	var _401_v7 _dafny.Set
	_ = _401_v7
	_401_v7 = _dafny.SetOf((func() bool {
		if (_400_v6).Contains(_397_v2) {
			return (_400_v6).Get(_397_v2).(bool)
		}
		return !(_398_v4)
	})(), _398_v4, _398_v4, _398_v4)
	var _402_v8 _dafny.Sequence
	_ = _402_v8
	_402_v8 = _dafny.UnicodeSeqOfUtf8Bytes("frrp")
	var _403_v9 _dafny.Array
	_ = _403_v9
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(7)
	_ = _len0_7
	var _nw68 _dafny.Array
	_ = _nw68
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw68 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) _dafny.Int = func(_404_i0 _dafny.Int) _dafny.Int {
			return Companion_Default___.SafeDivisionInt(_404_i0, _dafny.IntOfInt64(535))
		}
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw68 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw68).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw68).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	_403_v9 = _nw68
	var _405_globalState *GlobalState
	_ = _405_globalState
	var _nw69 *GlobalState = New_GlobalState_()
	_ = _nw69
	_nw69.Ctor__(_dafny.IntOfInt64(455), _dafny.IntOfInt64(119), (func() _dafny.Map {
		var _coll26 = _dafny.NewMapBuilder()
		_ = _coll26
		for _iter26 := _dafny.Iterate((_dafny.SeqOf(_396_v1)).Elements()); ; {
			_compr_26, _ok26 := _iter26()
			if !_ok26 {
				break
			}
			var _406_v0 _dafny.CodePoint
			_406_v0 = interface{}(_compr_26).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_396_v1), _406_v0) {
				_coll26.Add(_406_v0, _dafny.IntOfInt64(68))
			}
		}
		return _coll26.ToMap()
	}()).Update(_396_v1, _397_v2), true, _dafny.IntOfInt64(936), _dafny.IntOfInt64(277), _dafny.IntOfInt64(911), _dafny.MultiSetOf(_397_v2, (_dafny.Zero).Minus((func() _dafny.Set {
		var _coll27 = _dafny.NewBuilder()
		_ = _coll27
		for _iter27 := _dafny.Iterate((_dafny.SeqOf(_397_v2)).Elements()); ; {
			_compr_27, _ok27 := _iter27()
			if !_ok27 {
				break
			}
			var _407_v3 _dafny.Int
			_407_v3 = interface{}(_compr_27).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_397_v2), _407_v3) {
				_coll27.Add((_407_v3).Times((_dafny.Zero).Minus(_dafny.IntOfInt64(-233))))
			}
		}
		return _coll27.ToSet()
	}()).Cardinality()), (_401_v7).Cardinality(), _dafny.IntOfUint32((_402_v8).Cardinality())), true, _dafny.IntOfInt64(828), _dafny.IntOfInt64(622), _dafny.IntOfInt64(-572), _dafny.IntOfInt64(262), false, _403_v9, _dafny.IntOfInt64(726), _dafny.IntOfInt64(237))
	_405_globalState = _nw69
	var _408_v10 _dafny.Map
	_ = _408_v10
	_408_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_397_v2, _397_v2)
	var _hi0 _dafny.Int = _397_v2
	_ = _hi0
	for _409_i1 := Companion_Default___.Fm0(_397_v2, (_408_v10).Update(_dafny.IntOfInt64(-923), _397_v2), _405_globalState); _409_i1.Cmp(_hi0) < 0; _409_i1 = _409_i1.Plus(_dafny.One) {
		(_405_globalState).F16 = _409_i1
		var _410_v11 _dafny.MultiSet
		_ = _410_v11
		_410_v11 = _dafny.MultiSetOf(!(_398_v4), _398_v4, _398_v4)
		var _411_v12 _dafny.MultiSet
		_ = _411_v12
		_411_v12 = _dafny.MultiSetOf(_410_v11)
		var _412_v13 _dafny.Sequence
		_ = _412_v13
		_412_v13 = _dafny.SeqOf((func() _dafny.Int {
			if (_411_v12).Contains(_410_v11) {
				return (_411_v12).Multiplicity(_410_v11)
			}
			return _409_i1
		})())
		_412_v13 = _412_v13
		var _413_v14 _dafny.Map
		_ = _413_v14
		_413_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_398_v4, (_400_v6).Merge(_400_v6))
		_413_v14 = _413_v14
		var _414_v16 D0
		_ = _414_v16
		_414_v16 = Companion_D0_.Create_DC0_(_409_i1, _398_v4, _dafny.UnicodeSeqOfUtf8Bytes("yicqicecd"), func() _dafny.Map {
			var _coll28 = _dafny.NewMapBuilder()
			_ = _coll28
			for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(396), _dafny.IntOfInt64(-156))); ; {
				_compr_28, _ok28 := _iter28()
				if !_ok28 {
					break
				}
				var _415_v15 _dafny.Int
				_415_v15 = interface{}(_compr_28).(_dafny.Int)
				if ((_dafny.IntOfInt64(396)).Cmp(_415_v15) <= 0) && ((_415_v15).Cmp(_dafny.IntOfInt64(-156)) < 0) {
					_coll28.Add((_415_v15).Minus(_dafny.IntOfUint32((_402_v8).Cardinality())), _400_v6)
				}
			}
			return _coll28.ToMap()
		}())
		var _source6 D0 = _414_v16
		_ = _source6
		var _416___mcc_h0 _dafny.Int = _source6.Get_().(D0_DC0).Cf0
		_ = _416___mcc_h0
		var _417___mcc_h1 bool = _source6.Get_().(D0_DC0).Cf1
		_ = _417___mcc_h1
		var _418___mcc_h2 _dafny.Sequence = _source6.Get_().(D0_DC0).Cf2
		_ = _418___mcc_h2
		var _419___mcc_h3 _dafny.Map = _source6.Get_().(D0_DC0).Cf3
		_ = _419___mcc_h3
		var _420_cf3 _dafny.Map = _419___mcc_h3
		_ = _420_cf3
		var _421_cf2 _dafny.Sequence = _418___mcc_h2
		_ = _421_cf2
		var _422_cf1 bool = _417___mcc_h1
		_ = _422_cf1
		var _423_cf0 _dafny.Int = _416___mcc_h0
		_ = _423_cf0
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_403_v9), 0))
		_ = _index70
		(_403_v9).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.Fm0(_423_cf0, _408_v10, _405_globalState)).Times(_409_i1)))), (_index70).Int())
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_403_v9), 0))
		_ = _index71
		(_403_v9).ArraySet1(_397_v2, (_index71).Int())
		_408_v10 = (_408_v10).Update((_dafny.Zero).Minus((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)), (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))
		var _424_v17 D1
		_ = _424_v17
		_424_v17 = Companion_D1_.Create_DC1_(_399_v5)
		var _425_v18 _dafny.MultiSet
		_ = _425_v18
		_425_v18 = _dafny.MultiSetOf(_423_cf0)
		var _426_v19 _dafny.Map
		_ = _426_v19
		_426_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_cf1, _398_v4)
		var _427_v20 _dafny.Map
		_ = _427_v20
		_427_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)), _399_v5)
		var _428_v21 _dafny.Array
		_ = _428_v21
		var _nwElement0_11 _dafny.Sequence = _399_v5
		_ = _nwElement0_11
		var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(25))
		_ = _nw70
		(_nw70).ArraySet1(_nwElement0_11, 0)
		(_nw70).ArraySet1(_399_v5, 1)
		(_nw70).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_424_v17).Dtor_cf4(), _399_v5), 2)
		(_nw70).ArraySet1(_399_v5, 3)
		(_nw70).ArraySet1(_399_v5, 4)
		(_nw70).ArraySet1(_399_v5, 5)
		(_nw70).ArraySet1(_399_v5, 6)
		(_nw70).ArraySet1(_dafny.Companion_Sequence_.Update(_399_v5, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_423_cf0), _dafny.IntOfUint32((_399_v5).Cardinality()))).Uint32(), _422_cf1), 7)
		(_nw70).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_399_v5, (Companion_Default___.SafeIndex((_425_v18).Cardinality(), _dafny.IntOfUint32((_399_v5).Cardinality()))).Uint32(), _398_v4), (Companion_Default___.SafeIndex(_409_i1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_399_v5, (Companion_Default___.SafeIndex((_425_v18).Cardinality(), _dafny.IntOfUint32((_399_v5).Cardinality()))).Uint32(), _398_v4)).Cardinality()))).Uint32(), _422_cf1), _dafny.SeqOf(false, (func() bool {
			if (_426_v19).Contains(_422_cf1) {
				return (_426_v19).Get(_422_cf1).(bool)
			}
			return _398_v4
		})())), 8)
		(_nw70).ArraySet1(_399_v5, 9)
		(_nw70).ArraySet1(_399_v5, 10)
		(_nw70).ArraySet1(_399_v5, 11)
		(_nw70).ArraySet1(_399_v5, 12)
		(_nw70).ArraySet1(_399_v5, 13)
		(_nw70).ArraySet1((func() _dafny.Sequence {
			if (_427_v20).Contains(_dafny.IntOfInt64(-371)) {
				return (_427_v20).Get(_dafny.IntOfInt64(-371)).(_dafny.Sequence)
			}
			return _dafny.SeqOf(_422_cf1)
		})(), 14)
		(_nw70).ArraySet1(_399_v5, 15)
		(_nw70).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_399_v5, _399_v5), 16)
		(_nw70).ArraySet1(_dafny.SeqOf(false), 17)
		(_nw70).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_399_v5, _399_v5), 18)
		(_nw70).ArraySet1(_399_v5, 19)
		(_nw70).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_399_v5, (Companion_D1_.Create_DC1_(_399_v5)).Dtor_cf4()), 20)
		(_nw70).ArraySet1(_399_v5, 21)
		(_nw70).ArraySet1(_399_v5, 22)
		(_nw70).ArraySet1(_399_v5, 23)
		(_nw70).ArraySet1(_399_v5, 24)
		_428_v21 = _nw70
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_428_v21), 0))
		_ = _index72
		(_428_v21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_398_v4, false), _399_v5), (_index72).Int())
		var _429_v22 _dafny.MultiSet
		_ = _429_v22
		_429_v22 = _dafny.MultiSetOf(_402_v8)
		var _430_v23 _dafny.Array
		_ = _430_v23
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_8
		var _nw71 _dafny.Array
		_ = _nw71
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw71 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) D0 = (func(_431_i1 _dafny.Int, _432_cf1 bool, _433_v8 _dafny.Sequence, _434_cf3 _dafny.Map) func(_dafny.Int) D0 {
				return func(_435_i2 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_(_431_i1, _432_cf1, _433_v8, _434_cf3)
				}
			})(_409_i1, _422_cf1, _402_v8, _420_cf3)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw71 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw71).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw71).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_430_v23 = _nw71
		var _436_v24 _dafny.Sequence
		_ = _436_v24
		_436_v24 = _dafny.SeqOf(_dafny.SetOf(_398_v4))
		var _437_v25 _dafny.Array
		_ = _437_v25
		var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw72
		_437_v25 = _nw72
		var _438_v26 _dafny.Set
		_ = _438_v26
		_438_v26 = _dafny.SetOf(_437_v25)
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_428_v21), 0))
		_ = _index73
		var _rhs75 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_399_v5, _399_v5)
		_ = _rhs75
		var _rhs76 _dafny.MultiSet = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jaeue"))
		_ = _rhs76
		var _rhs77 _dafny.Array = _430_v23
		_ = _rhs77
		var _rhs78 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(203))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}((func(_439_v7 _dafny.Set) func(_dafny.Int) _dafny.Set {
			return func(_440_i3 _dafny.Int) _dafny.Set {
				return _439_v7
			}
		})(_401_v7)))
		_ = _rhs78
		var _rhs79 _dafny.Int = Companion_Default___.Fm0(_397_v2, (Companion_Default___.Fm1(_408_v10, (_408_v10).Cardinality(), (_438_v26).Cardinality(), _405_globalState)).Merge(_408_v10), _405_globalState)
		_ = _rhs79
		var _lhs54 _dafny.Array = _428_v21
		_ = _lhs54
		var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_428_v21), 0))
		_ = _lhs55
		var _lhs56 *GlobalState = _405_globalState
		_ = _lhs56
		(_lhs54).ArraySet1(_rhs75, (_lhs55).Int())
		_429_v22 = _rhs76
		_430_v23 = _rhs77
		_436_v24 = _rhs78
		_lhs56.F6 = _rhs79
		var _441_v27 _dafny.MultiSet
		_ = _441_v27
		_441_v27 = _dafny.MultiSetOf(_437_v25)
		_421_cf2 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("flvuwul"), (Companion_Default___.SafeIndex(_423_cf0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("flvuwul")).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
			if _398_v4 {
				return _396_v1
			}
			return (_402_v8).Select((Companion_Default___.SafeIndex(_423_cf0, _dafny.IntOfUint32((_402_v8).Cardinality()))).Uint32()).(_dafny.CodePoint)
		})()), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(240), (_441_v27).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("flvuwul"), (Companion_Default___.SafeIndex(_423_cf0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("flvuwul")).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
			if _398_v4 {
				return _396_v1
			}
			return (_402_v8).Select((Companion_Default___.SafeIndex(_423_cf0, _dafny.IntOfUint32((_402_v8).Cardinality()))).Uint32()).(_dafny.CodePoint)
		})())).Cardinality()))).Uint32(), _396_v1)
	}
	var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))
	_ = _index74
	(_403_v9).ArraySet1(_397_v2, (_index74).Int())
	var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))
	_ = _index75
	(_403_v9).ArraySet1(Companion_Default___.SafeModuloInt(_397_v2, _397_v2), (_index75).Int())
	var _442_v28 *C0
	_ = _442_v28
	var _nw73 *C0 = New_C0_()
	_ = _nw73
	_nw73.Ctor__(_398_v4)
	_442_v28 = _nw73
	var _443_v29 _dafny.MultiSet
	_ = _443_v29
	_443_v29 = _dafny.MultiSetOf(_397_v2)
	var _444_v30 _dafny.Map
	_ = _444_v30
	_444_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _dafny.MultiSetOf(_dafny.IntOfInt64(76)))
	var _445_v31 D1
	_ = _445_v31
	_445_v31 = Companion_D1_.Create_DC4_(_397_v2, (_442_v28).F17(), _444_v30)
	var _446_v32 D1
	_ = _446_v32
	_446_v32 = Companion_D1_.Create_DC5_(_445_v31)
	var _pat_let_tv16 = _398_v4
	_ = _pat_let_tv16
	var _pat_let_tv17 = _442_v28
	_ = _pat_let_tv17
	var _pat_let_tv18 = _398_v4
	_ = _pat_let_tv18
	var _pat_let_tv19 = _442_v28
	_ = _pat_let_tv19
	var _pat_let_tv20 = _398_v4
	_ = _pat_let_tv20
	var _rhs80 _dafny.CodePoint = Companion_Default___.Fm3(_398_v4, _405_globalState)
	_ = _rhs80
	var _rhs81 bool = !(_398_v4) || ((_443_v29).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(187), (_408_v10).Cardinality())))
	_ = _rhs81
	var _rhs82 bool = func(_source7 D1) bool {
		if _source7.Is_DC2() {
			var _447___mcc_h4 _dafny.CodePoint = _source7.Get_().(D1_DC2).Cf5
			_ = _447___mcc_h4
			var _448_cf5 _dafny.CodePoint = _447___mcc_h4
			_ = _448_cf5
			return true
		} else if _source7.Is_DC3() {
			var _449___mcc_h5 _dafny.Map = _source7.Get_().(D1_DC3).Cf6
			_ = _449___mcc_h5
			var _450___mcc_h6 _dafny.Int = _source7.Get_().(D1_DC3).Cf7
			_ = _450___mcc_h6
			var _451___mcc_h7 D0 = _source7.Get_().(D1_DC3).Cf8
			_ = _451___mcc_h7
			var _452_cf8 D0 = _451___mcc_h7
			_ = _452_cf8
			var _453_cf7 _dafny.Int = _450___mcc_h6
			_ = _453_cf7
			var _454_cf6 _dafny.Map = _449___mcc_h5
			_ = _454_cf6
			return _pat_let_tv16
		} else if _source7.Is_DC4() {
			var _455___mcc_h8 _dafny.Int = _source7.Get_().(D1_DC4).Cf9
			_ = _455___mcc_h8
			var _456___mcc_h9 bool = _source7.Get_().(D1_DC4).Cf10
			_ = _456___mcc_h9
			var _457___mcc_h10 _dafny.Map = _source7.Get_().(D1_DC4).Cf11
			_ = _457___mcc_h10
			var _458_cf11 _dafny.Map = _457___mcc_h10
			_ = _458_cf11
			var _459_cf10 bool = _456___mcc_h9
			_ = _459_cf10
			var _460_cf9 _dafny.Int = _455___mcc_h8
			_ = _460_cf9
			return (_pat_let_tv17).F17()
		} else if _source7.Is_DC1() {
			var _461___mcc_h11 _dafny.Sequence = _source7.Get_().(D1_DC1).Cf4
			_ = _461___mcc_h11
			var _462_cf4 _dafny.Sequence = _461___mcc_h11
			_ = _462_cf4
			return _pat_let_tv18
		} else {
			var _463___mcc_h12 D1 = _source7.Get_().(D1_DC5).Cf12
			_ = _463___mcc_h12
			var _464_cf12 D1 = _463___mcc_h12
			_ = _464_cf12
			return !((_pat_let_tv19).F17()) || (_pat_let_tv20)
		}
	}(_446_v32)
	_ = _rhs82
	var _rhs83 _dafny.CodePoint = (func() _dafny.CodePoint {
		if _398_v4 {
			return _396_v1
		}
		return _396_v1
	})()
	_ = _rhs83
	var _rhs84 _dafny.CodePoint = _396_v1
	_ = _rhs84
	_396_v1 = _rhs80
	_398_v4 = _rhs81
	_398_v4 = _rhs82
	_396_v1 = _rhs83
	_396_v1 = _rhs84
	var _465_v33 _dafny.MultiSet
	_ = _465_v33
	_465_v33 = _dafny.MultiSetOf(false)
	var _466_v34 _dafny.Array
	_ = _466_v34
	var _nw74 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
	_ = _nw74
	_466_v34 = _nw74
	var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
	_ = _index76
	(_466_v34).ArraySet1(_398_v4, (_index76).Int())
	var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_466_v34), 0))
	_ = _index77
	(_466_v34).ArraySet1((_442_v28).F17(), (_index77).Int())
	var _467_v35 _dafny.Map
	_ = _467_v35
	_467_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_442_v28).F17(), (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))
	var _468_v36 _dafny.Sequence
	_ = _468_v36
	_468_v36 = _dafny.SeqOf(_408_v10, _408_v10)
	var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
	_ = _index78
	var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_466_v34), 0))
	_ = _index79
	var _rhs85 _dafny.MultiSet = _465_v33
	_ = _rhs85
	var _rhs86 bool = Companion_Default___.Fm4(((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Plus((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)), ((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Cmp((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)) >= 0, _405_globalState)
	_ = _rhs86
	var _rhs87 bool = Companion_Default___.Fm4(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_402_v8, (Companion_Default___.SafeIndex((func() _dafny.Int {
		if (_467_v35).Contains(Companion_Default___.Fm4((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_442_v28).F17(), _405_globalState)) {
			return (_467_v35).Get(Companion_Default___.Fm4((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_442_v28).F17(), _405_globalState)).(_dafny.Int)
		}
		return Companion_Default___.Fm0(_dafny.IntOfInt64(16), (_468_v36).Select((Companion_Default___.SafeIndex((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_468_v36).Cardinality()))).Uint32()).(_dafny.Map), _405_globalState)
	})(), _dafny.IntOfUint32((_402_v8).Cardinality()))).Uint32(), _396_v1)).Cardinality()), (_442_v28).F17(), _405_globalState)
	_ = _rhs87
	var _rhs88 bool = Companion_Default___.Fm4(_397_v2, (_442_v28).F17(), _405_globalState)
	_ = _rhs88
	var _lhs57 _dafny.Array = _466_v34
	_ = _lhs57
	var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
	_ = _lhs58
	var _lhs59 _dafny.Array = _466_v34
	_ = _lhs59
	var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_466_v34), 0))
	_ = _lhs60
	_465_v33 = _rhs85
	_398_v4 = _rhs86
	(_lhs57).ArraySet1(_rhs87, (_lhs58).Int())
	(_lhs59).ArraySet1(_rhs88, (_lhs60).Int())
	var _469_v37 *C0
	_ = _469_v37
	var _nw75 *C0 = New_C0_()
	_ = _nw75
	_nw75.Ctor__((_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool))
	_469_v37 = _nw75
	var _470_v38 _dafny.Map
	_ = _470_v38
	_470_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _402_v8)
	_470_v38 = (_470_v38).Update(_dafny.IntOfUint32((_399_v5).Cardinality()), _402_v8)
	var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
	_ = _index80
	(_466_v34).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
		if false {
			return _402_v8
		}
		return _dafny.Companion_Sequence_.Update(_402_v8, (Companion_Default___.SafeIndex(_397_v2, _dafny.IntOfUint32((_402_v8).Cardinality()))).Uint32(), _396_v1)
	})(), _402_v8), (_index80).Int())
	for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_403_v9), 0))); ; {
		_guard_loop_0, _ok29 := _iter29()
		if !_ok29 {
			break
		}
		var _471_i4 _dafny.Int
		_471_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_471_i4).Sign() != -1) && ((_471_i4).Cmp(_dafny.ArrayLen((_403_v9), 0)) < 0)) {
			(_403_v9).ArraySet1((_471_i4).Times(_dafny.IntOfInt64(671)), (_471_i4).Int())
		}
	}
	var _472_v39 _dafny.Array
	_ = _472_v39
	var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
	_ = _nw76
	_472_v39 = _nw76
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_472_v39), 0))); ; {
		_guard_loop_1, _ok30 := _iter30()
		if !_ok30 {
			break
		}
		var _473_i5 _dafny.Int
		_473_i5 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_473_i5).Sign() != -1) && ((_473_i5).Cmp(_dafny.ArrayLen((_472_v39), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_472_v39, (_473_i5).Int(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_397_v2), _dafny.SeqOf(_dafny.IntOfInt64(129))))))
		}
	}
	for _iter31 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok31 := _iter31()
		if !_ok31 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	(_405_globalState).F16 = _397_v2
	var _474_i6 _dafny.Int
	_ = _474_i6
	_474_i6 = _dafny.Zero
	{
		for (_469_v37).F17() {
			{
				if (_474_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_474_i6 = (_474_i6).Plus(_dafny.One)
				var _475_v40 _dafny.Map
				_ = _475_v40
				_475_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_dafny.IntOfInt64(861), (_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool), _444_v30), _403_v9)
				var _476_v41 D1
				_ = _476_v41
				_476_v41 = Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ksjsyh")).Cardinality()), (_442_v28).F17(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _443_v29))
				var _477_v42 _dafny.MultiSet
				_ = _477_v42
				_477_v42 = _dafny.MultiSetOf(_403_v9, _403_v9, (func() _dafny.Array {
					if (_475_v40).Contains(_476_v41) {
						return (_475_v40).Get(_476_v41).(_dafny.Array)
					}
					return _403_v9
				})(), _403_v9, _403_v9)
				_477_v42 = _477_v42
				var _478_v43 _dafny.Sequence
				_ = _478_v43
				_478_v43 = _dafny.SeqOf(_466_v34, _466_v34, _466_v34)
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
				_ = _index81
				(_466_v34).ArraySet1(!(_dafny.Companion_Sequence_.IsPrefixOf(_478_v43, _478_v43)), (_index81).Int())
				var _479_v44 _dafny.Array
				_ = _479_v44
				var _nwElement0_12 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("fqsutwajj")
				_ = _nwElement0_12
				var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(8))
				_ = _nw77
				(_nw77).ArraySet1(_nwElement0_12, 0)
				(_nw77).ArraySet1(_402_v8, 1)
				(_nw77).ArraySet1(_402_v8, 2)
				(_nw77).ArraySet1(_402_v8, 3)
				(_nw77).ArraySet1(_402_v8, 4)
				(_nw77).ArraySet1(_402_v8, 5)
				(_nw77).ArraySet1(_dafny.Companion_Sequence_.Update(_402_v8, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool), _397_v2)).Cardinality(), _dafny.IntOfUint32((_402_v8).Cardinality()))).Uint32(), _396_v1), 6)
				(_nw77).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ibgqhjni"), 7)
				_479_v44 = _nw77
				var _480_v45 _dafny.Map
				_ = _480_v45
				_480_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_442_v28).F17(), _479_v44)
				_480_v45 = (_480_v45).Update(true, _479_v44)
				var _481_v47 *C0
				_ = _481_v47
				var _nw78 *C0 = New_C0_()
				_ = _nw78
				_nw78.Ctor__(!(_400_v6).Equals(func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter32 := _dafny.Iterate((Companion_Default___.Fm5(_405_globalState)).Keys().Elements()); ; {
						_compr_29, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _482_v46 _dafny.Int
						_482_v46 = interface{}(_compr_29).(_dafny.Int)
						if (Companion_Default___.Fm5(_405_globalState)).Contains(_482_v46) {
							_coll29.Add((_482_v46).Minus(_397_v2), (_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool))
						}
					}
					return _coll29.ToMap()
				}()))
				_481_v47 = _nw78
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
	_ = _index82
	(_466_v34).ArraySet1((_469_v37).F17(), (_index82).Int())
	var _hi1 _dafny.Int = (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)
	_ = _hi1
	for _483_i7 := _dafny.IntOfInt64(85); _483_i7.Cmp(_hi1) < 0; _483_i7 = _483_i7.Plus(_dafny.One) {
		var _484_v48 *C0
		_ = _484_v48
		var _nw79 *C0 = New_C0_()
		_ = _nw79
		_nw79.Ctor__((_469_v37).F17())
		_484_v48 = _nw79
		var _485_v49 _dafny.Sequence
		_ = _485_v49
		_485_v49 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(874))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}((func(_486_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_487_i8 _dafny.Int) _dafny.Sequence {
				return _486_v8
			}
		})(_402_v8)))).Cardinality())))
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_472_v39), 0))
		_ = _index83
		(_472_v39).ArraySet1(_485_v49, (_index83).Int())
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(569), _dafny.ArrayLen((_472_v39), 0))
		_ = _index84
		(_472_v39).ArraySet1(_485_v49, (_index84).Int())
		var _488_v50 _dafny.Map
		_ = _488_v50
		_488_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_397_v2, _400_v6)
		var _489_v51 D0
		_ = _489_v51
		_489_v51 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(863), (_469_v37).F17(), _402_v8, _488_v50)
		var _490_v52 D1
		_ = _490_v52
		_490_v52 = Companion_D1_.Create_DC3_(_408_v10, _397_v2, _489_v51)
		var _491_v53 _dafny.Set
		_ = _491_v53
		_491_v53 = _dafny.SetOf(_490_v52, Companion_D1_.Create_DC3_(_408_v10, _483_i7, _489_v51), _490_v52)
		var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))
		_ = _index85
		(_403_v9).ArraySet1(Companion_Default___.SafeModuloInt(_483_i7, ((_491_v53).Intersection(_dafny.SetOf(_490_v52))).Cardinality()), (_index85).Int())
		_467_v35 = (_467_v35).Update((_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool), (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))
	}
	if _dafny.Companion_Sequence_.IsPrefixOf(_402_v8, _402_v8) {
		_397_v2 = (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)
		var _492_v54 D1
		_ = _492_v54
		_492_v54 = Companion_D1_.Create_DC2_(_396_v1)
		var _493_v55 D1
		_ = _493_v55
		_493_v55 = Companion_D1_.Create_DC4_(_397_v2, (_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool), _444_v30)
		var _494_v56 _dafny.Map
		_ = _494_v56
		_494_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_442_v28, _465_v33)
		var _495_v58 _dafny.Set
		_ = _495_v58
		_495_v58 = _dafny.SetOf((func() _dafny.Map {
			var _coll30 = _dafny.NewMapBuilder()
			_ = _coll30
			for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-794), _dafny.IntOfInt64(807))); ; {
				_compr_30, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _496_v57 _dafny.Int
				_496_v57 = interface{}(_compr_30).(_dafny.Int)
				if ((_dafny.IntOfInt64(-794)).Cmp(_496_v57) <= 0) && ((_496_v57).Cmp(_dafny.IntOfInt64(807)) < 0) {
					_coll30.Add(Companion_Default___.SafeDivisionInt(_496_v57, _397_v2), _396_v1)
				}
			}
			return _coll30.ToMap()
		}()).Cardinality())
		var _497_v59 _dafny.Array
		_ = _497_v59
		var _nwElement0_13 _dafny.MultiSet = Companion_Default___.Fm6(false, _398_v4, (_442_v28).F17(), _405_globalState)
		_ = _nwElement0_13
		var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(20))
		_ = _nw80
		(_nw80).ArraySet1(_nwElement0_13, 0)
		(_nw80).ArraySet1(_dafny.MultiSetOf((_493_v55).Dtor_cf10(), (_442_v28).F17(), false, (_442_v28).F17(), true), 1)
		(_nw80).ArraySet1(_465_v33, 2)
		(_nw80).ArraySet1(_465_v33, 3)
		(_nw80).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7((_492_v54).Dtor_cf5(), _405_globalState), _399_v5)), 4)
		(_nw80).ArraySet1(_465_v33, 5)
		(_nw80).ArraySet1(Companion_Default___.Fm6(true, (_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool), (_469_v37).F17(), _405_globalState), 6)
		(_nw80).ArraySet1(_465_v33, 7)
		(_nw80).ArraySet1(_465_v33, 8)
		(_nw80).ArraySet1(_465_v33, 9)
		(_nw80).ArraySet1(_465_v33, 10)
		(_nw80).ArraySet1(_465_v33, 11)
		(_nw80).ArraySet1(_465_v33, 12)
		(_nw80).ArraySet1((_465_v33).Update((_469_v37).F17(), Companion_Default___.Abs(_397_v2)), 13)
		(_nw80).ArraySet1((func() _dafny.MultiSet {
			if (_494_v56).Contains(_469_v37) {
				return (_494_v56).Get(_469_v37).(_dafny.MultiSet)
			}
			return (_465_v33).Update((_442_v28).F17(), Companion_Default___.Abs((_495_v58).Cardinality()))
		})(), 14)
		(_nw80).ArraySet1(_dafny.MultiSetOf(true, (_469_v37).F17()), 15)
		(_nw80).ArraySet1(_465_v33, 16)
		(_nw80).ArraySet1(_dafny.MultiSetOf((_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool), (_442_v28).F17()), 17)
		(_nw80).ArraySet1(_465_v33, 18)
		(_nw80).ArraySet1(_465_v33, 19)
		_497_v59 = _nw80
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_497_v59), 0))
		_ = _index86
		(_497_v59).ArraySet1(_465_v33, (_index86).Int())
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_497_v59), 0))
		_ = _index87
		var _rhs89 D1 = _492_v54
		_ = _rhs89
		var _rhs90 _dafny.MultiSet = (_dafny.MultiSetOf((_469_v37).F17())).Union(_465_v33)
		_ = _rhs90
		var _lhs61 _dafny.Array = _497_v59
		_ = _lhs61
		var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_497_v59), 0))
		_ = _lhs62
		_492_v54 = _rhs89
		(_lhs61).ArraySet1(_rhs90, (_lhs62).Int())
		var _498_v60 *C0
		_ = _498_v60
		var _nw81 *C0 = New_C0_()
		_ = _nw81
		_nw81.Ctor__((func() bool {
			if _398_v4 {
				return !((_469_v37).F17())
			}
			return (_469_v37).F17()
		})())
		_498_v60 = _nw81
		var _499_v61 _dafny.Set
		_ = _499_v61
		_499_v61 = _dafny.SetOf(_399_v5)
		if ((_499_v61).Intersection(_499_v61)).IsProperSubsetOf(_499_v61) {
			var _500_v62 _dafny.Array
			_ = _500_v62
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_9
			var _nw82 _dafny.Array
			_ = _nw82
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw82 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.CodePoint = func(_501_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw82 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw82).ArraySet1CodePoint(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw82).ArraySet1CodePoint(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_500_v62 = _nw82
			var _rhs91 _dafny.Array = _500_v62
			_ = _rhs91
			var _rhs92 bool = (_498_v60).F17()
			_ = _rhs92
			_500_v62 = _rhs91
			_398_v4 = _rhs92
			var _502_v63 _dafny.Array
			_ = _502_v63
			var _nw83 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
			_ = _nw83
			_502_v63 = _nw83
			var _503_v64 _dafny.Set
			_ = _503_v64
			_503_v64 = _dafny.SetOf(_502_v63)
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index88
			var _rhs93 _dafny.Array = _403_v9
			_ = _rhs93
			var _rhs94 bool = (_503_v64).IsProperSubsetOf(_dafny.SetOf(_502_v63))
			_ = _rhs94
			var _rhs95 bool = Companion_Default___.Fm4((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_469_v37).F17(), _405_globalState)
			_ = _rhs95
			var _rhs96 *C0 = _442_v28
			_ = _rhs96
			var _lhs63 _dafny.Array = _466_v34
			_ = _lhs63
			var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _lhs64
			_403_v9 = _rhs93
			(_lhs63).ArraySet1(_rhs94, (_lhs64).Int())
			_398_v4 = _rhs95
			_442_v28 = _rhs96
			var _504_v65 _dafny.Map
			_ = _504_v65
			_504_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_397_v2, _400_v6)
			var _505_v66 D0
			_ = _505_v66
			_505_v66 = Companion_D0_.Create_DC0_(_397_v2, (_498_v60).F17(), _402_v8, _504_v65)
			var _506_v67 D1
			_ = _506_v67
			_506_v67 = Companion_D1_.Create_DC3_(_408_v10, (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _505_v66)
			_506_v67 = Companion_Default___.Fm8((_442_v28).F17(), (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _405_globalState)
			var _507_v68 _dafny.Map
			_ = _507_v68
			_507_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_469_v37).F17(), _dafny.UnicodeSeqOfUtf8Bytes("fyk"))
			_507_v68 = (_507_v68).Update((_469_v37).F17(), _402_v8)
			var _508_v69 _dafny.Sequence
			_ = _508_v69
			_508_v69 = _dafny.SeqOf(Companion_D1_.Create_DC3_((_408_v10).Update(_397_v2, (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)), (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _505_v66), _506_v67, Companion_D1_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_495_v58, _495_v58)).Cardinality()), _dafny.IntOfInt64(214)), _397_v2, Companion_D0_.Create_DC0_((_401_v7).Cardinality(), _398_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_509_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_510_i10 _dafny.Int) _dafny.CodePoint {
					return _509_v1
				}
			})(_396_v1))), _504_v65)))
			var _511_v70 _dafny.MultiSet
			_ = _511_v70
			_511_v70 = _dafny.MultiSetOf(_dafny.SeqOf(_506_v67, _506_v67), _508_v69, _508_v69, _508_v69, Companion_Default___.Fm9(_405_globalState))
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index89
			(_466_v34).ArraySet1((_511_v70).IsProperSubsetOf(_511_v70), (_index89).Int())
		} else {
			var _512_v71 _dafny.Sequence
			_ = _512_v71
			_512_v71 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _397_v2), _397_v2)
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))
			_ = _index90
			(_403_v9).ArraySet1((_dafny.Zero).Minus((_512_v71).Select((Companion_Default___.SafeIndex(((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Plus((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_512_v71).Cardinality()))).Uint32()).(_dafny.Int)), (_index90).Int())
			var _513_v72 *C0
			_ = _513_v72
			var _nw84 *C0 = New_C0_()
			_ = _nw84
			_nw84.Ctor__((_401_v7).IsDisjointFrom(_401_v7))
			_513_v72 = _nw84
			var _514_v73 _dafny.Map
			_ = _514_v73
			_514_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-377), _400_v6)
			var _515_v74 D0
			_ = _515_v74
			_515_v74 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_402_v8).Cardinality()), (_469_v37).F17(), _402_v8, _514_v73)
			var _516_v75 D1
			_ = _516_v75
			_516_v75 = Companion_D1_.Create_DC3_(_408_v10, ((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Times(_397_v2), _515_v74)
			var _pat_let_tv21 = _512_v71
			_ = _pat_let_tv21
			var _pat_let_tv22 = _512_v71
			_ = _pat_let_tv22
			var _pat_let_tv23 = _408_v10
			_ = _pat_let_tv23
			var _pat_let_tv24 = _403_v9
			_ = _pat_let_tv24
			var _pat_let_tv25 = _403_v9
			_ = _pat_let_tv25
			_516_v75 = func(_pat_let16_0 D1) D1 {
				return func(_517_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let17_0 _dafny.Map) D1 {
						return func(_519_dt__update_hcf6_h0 _dafny.Map) D1 {
							return Companion_D1_.Create_DC3_(_519_dt__update_hcf6_h0, (_517_dt__update__tmp_h0).Dtor_cf7(), (_517_dt__update__tmp_h0).Dtor_cf8())
						}(_pat_let17_0)
					}(func() _dafny.Map {
						var _coll31 = _dafny.NewMapBuilder()
						_ = _coll31
						for _iter34 := _dafny.Iterate((_pat_let_tv21).Elements()); ; {
							_compr_31, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _518_v76 _dafny.Int
							_518_v76 = interface{}(_compr_31).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_pat_let_tv22, _518_v76) {
								_coll31.Add((_518_v76).Times((_pat_let_tv25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_pat_let_tv24), 0))).Int()).(_dafny.Int)), (_pat_let_tv23).Cardinality())
							}
						}
						return _coll31.ToMap()
					}())
				}(_pat_let16_0)
			}(_516_v75)
			var _520_v77 _dafny.Map
			_ = _520_v77
			_520_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_442_v28).F17(), (_469_v37).F17())
			var _521_v78 _dafny.Map
			_ = _521_v78
			_521_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("fh"), (func() bool {
				if (_520_v77).Contains((_469_v37).F17()) {
					return (_520_v77).Get((_469_v37).F17()).(bool)
				}
				return Companion_Default___.Fm4(_397_v2, Companion_Default___.Fm4(_397_v2, (_513_v72).F17(), _405_globalState), _405_globalState)
			})())
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index91
			(_466_v34).ArraySet1((func() bool {
				if (_521_v78).Contains(_dafny.Companion_Sequence_.Concatenate(_402_v8, _402_v8)) {
					return (_521_v78).Get(_dafny.Companion_Sequence_.Concatenate(_402_v8, _402_v8)).(bool)
				}
				return ((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_399_v5).Cardinality())) <= 0
			})(), (_index91).Int())
			(_405_globalState).F16 = (_dafny.Zero).Minus((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))
		}
		var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))
		_ = _index92
		(_403_v9).ArraySet1((func() _dafny.Int {
			if true {
				return (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)
			}
			return _397_v2
		})(), (_index92).Int())
	} else {
		_403_v9 = _403_v9
		if (_401_v7).IsProperSubsetOf(_401_v7) {
			(_405_globalState).F16 = (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)
			_398_v4 = (_442_v28).F17()
			var _522_v79 _dafny.Array
			_ = _522_v79
			var _nw85 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
			_ = _nw85
			_522_v79 = _nw85
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_522_v79), 0))
			_ = _index93
			(_522_v79).ArraySet1(_442_v28, (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_522_v79), 0))
			_ = _index94
			(_522_v79).ArraySet1(_442_v28, (_index94).Int())
			_397_v2 = Companion_Default___.SafeDivisionInt((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
				if (_467_v35).Contains(Companion_Default___.Fm4((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_442_v28).F17(), _405_globalState)) {
					return (_467_v35).Get(Companion_Default___.Fm4((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_442_v28).F17(), _405_globalState)).(_dafny.Int)
				}
				return Companion_Default___.Fm0((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _408_v10, _405_globalState)
			})())
			(_405_globalState).F6 = Companion_Default___.SafeModuloInt((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_dafny.IntOfUint32((_399_v5).Cardinality())))
		} else {
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index95
			(_466_v34).ArraySet1((_442_v28).F17(), (_index95).Int())
			_396_v1 = _396_v1
			var _523_v80 _dafny.Array
			_ = _523_v80
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(7))
			_ = _nw86
			_523_v80 = _nw86
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_523_v80), 0))
			_ = _index96
			(_523_v80).ArraySet1(_443_v29, (_index96).Int())
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_523_v80), 0))
			_ = _index97
			var _rhs97 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_402_v8, (Companion_Default___.SafeIndex((_dafny.IntOfInt64(-462)).Minus(_397_v2), _dafny.IntOfUint32((_402_v8).Cardinality()))).Uint32(), _396_v1)
			_ = _rhs97
			var _rhs98 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v9, _396_v1)).Cardinality(), Companion_Default___.SafeModuloInt(_397_v2, (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)))
			_ = _rhs98
			var _rhs99 _dafny.MultiSet = _dafny.MultiSetOf((_dafny.Zero).Minus(_397_v2))
			_ = _rhs99
			var _rhs100 *C0 = _442_v28
			_ = _rhs100
			var _rhs101 _dafny.Array = _403_v9
			_ = _rhs101
			var _lhs65 _dafny.Array = _523_v80
			_ = _lhs65
			var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_523_v80), 0))
			_ = _lhs66
			var _lhs67 *GlobalState = _405_globalState
			_ = _lhs67
			_402_v8 = _rhs97
			_397_v2 = _rhs98
			(_lhs65).ArraySet1(_rhs99, (_lhs66).Int())
			_442_v28 = _rhs100
			_lhs67.F14 = _rhs101
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index98
			(_466_v34).ArraySet1((_469_v37).F17(), (_index98).Int())
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))
			_ = _index99
			(_403_v9).ArraySet1(_397_v2, (_index99).Int())
		}
		if (_399_v5).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_408_v10).Contains((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)) {
				return (_408_v10).Get((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).(_dafny.Int)
			}
			return _397_v2
		})(), _dafny.IntOfUint32((_399_v5).Cardinality()))).Uint32()).(bool) {
			var _524_v81 _dafny.Map
			_ = _524_v81
			_524_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_402_v8, _443_v29)
			var _525_v82 _dafny.Array
			_ = _525_v82
			var _nwElement0_14 _dafny.MultiSet = _443_v29
			_ = _nwElement0_14
			var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(27))
			_ = _nw87
			(_nw87).ArraySet1(_nwElement0_14, 0)
			(_nw87).ArraySet1(_dafny.MultiSetOf(_397_v2), 1)
			(_nw87).ArraySet1(_dafny.MultiSetOf((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)), 2)
			(_nw87).ArraySet1(_443_v29, 3)
			(_nw87).ArraySet1(_443_v29, 4)
			(_nw87).ArraySet1(_443_v29, 5)
			(_nw87).ArraySet1(_443_v29, 6)
			(_nw87).ArraySet1(_443_v29, 7)
			(_nw87).ArraySet1(_443_v29, 8)
			(_nw87).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))), 9)
			(_nw87).ArraySet1(_443_v29, 10)
			(_nw87).ArraySet1(_443_v29, 11)
			(_nw87).ArraySet1(_443_v29, 12)
			(_nw87).ArraySet1(_443_v29, 13)
			(_nw87).ArraySet1(_443_v29, 14)
			(_nw87).ArraySet1((func() _dafny.MultiSet {
				if (_524_v81).Contains(_402_v8) {
					return (_524_v81).Get(_402_v8).(_dafny.MultiSet)
				}
				return _443_v29
			})(), 15)
			(_nw87).ArraySet1(_443_v29, 16)
			(_nw87).ArraySet1(_443_v29, 17)
			(_nw87).ArraySet1(_443_v29, 18)
			(_nw87).ArraySet1(_443_v29, 19)
			(_nw87).ArraySet1(_443_v29, 20)
			(_nw87).ArraySet1(_443_v29, 21)
			(_nw87).ArraySet1(_443_v29, 22)
			(_nw87).ArraySet1((func() _dafny.MultiSet {
				if (_444_v30).Contains((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)) {
					return (_444_v30).Get((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).(_dafny.MultiSet)
				}
				return _443_v29
			})(), 23)
			(_nw87).ArraySet1(_443_v29, 24)
			(_nw87).ArraySet1(_443_v29, 25)
			(_nw87).ArraySet1(_443_v29, 26)
			_525_v82 = _nw87
			var _526_v83 _dafny.Map
			_ = _526_v83
			_526_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_397_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_397_v2, _398_v4))
			var _527_v84 D0
			_ = _527_v84
			_527_v84 = Companion_D0_.Create_DC0_((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_442_v28).F17(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(105))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}((func(_528_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_529_i11 _dafny.Int) _dafny.CodePoint {
					return _528_v1
				}
			})(_396_v1))), _526_v83)
			var _530_v85 _dafny.Map
			_ = _530_v85
			_530_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(_408_v10, (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), _527_v84)), _469_v37)
			var _531_v86 _dafny.Sequence
			_ = _531_v86
			_531_v86 = _dafny.SeqOf((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_530_v85).Cardinality(), (_dafny.Zero).Minus((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)))
			var _532_v87 _dafny.Map
			_ = _532_v87
			_532_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v86, (_442_v28).F17())
			var _533_v88 _dafny.Int
			_ = _533_v88
			var _534_v89 _dafny.Array
			_ = _534_v89
			var _535_v90 _dafny.Set
			_ = _535_v90
			var _out0 _dafny.Int
			_ = _out0
			var _out1 _dafny.Array
			_ = _out1
			var _out2 _dafny.Set
			_ = _out2
			_out0, _out1, _out2 = Companion_Default___.M0(_525_v82, (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), ((_532_v87).Cardinality()).Cmp(_397_v2) > 0, ((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Cmp(Companion_Default___.Fm0(_397_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-732), _dafny.IntOfInt64(-173)), _405_globalState)) != 0, _405_globalState)
			_533_v88 = _out0
			_534_v89 = _out1
			_535_v90 = _out2
			_399_v5 = _399_v5
			var _536_v91 _dafny.Int
			_ = _536_v91
			var _537_v92 _dafny.Array
			_ = _537_v92
			var _538_v93 _dafny.Set
			_ = _538_v93
			var _out3 _dafny.Int
			_ = _out3
			var _out4 _dafny.Array
			_ = _out4
			var _out5 _dafny.Set
			_ = _out5
			_out3, _out4, _out5 = Companion_Default___.M0(_525_v82, _533_v88, !((_dafny.IntOfInt64(690)).Cmp((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)) <= 0), (_469_v37).F17(), _405_globalState)
			_536_v91 = _out3
			_537_v92 = _out4
			_538_v93 = _out5
			var _539_v94 _dafny.Set
			_ = _539_v94
			_539_v94 = _dafny.SetOf(_533_v88)
			var _540_v95 _dafny.Sequence
			_ = _540_v95
			_540_v95 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(611))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}((func(_541_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_542_i13 _dafny.Int) _dafny.CodePoint {
					return _541_v1
				}
			})(_396_v1))))
			var _543_v96 _dafny.Map
			_ = _543_v96
			_543_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_539_v94).Intersection(_539_v94), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(277))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_544_v84 D0) func(_dafny.Int) _dafny.Sequence {
				return func(_545_i12 _dafny.Int) _dafny.Sequence {
					return (_544_v84).Dtor_cf2()
				}
			})(_527_v84))), _540_v95))
			(_405_globalState).F16 = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_543_v96).Contains((_539_v94).Intersection(_539_v94)) {
					return (_543_v96).Get((_539_v94).Intersection(_539_v94)).(_dafny.Sequence)
				}
				return _540_v95
			})()).Cardinality())
			var _pat_let_tv26 = _445_v31
			_ = _pat_let_tv26
			_446_v32 = func(_pat_let18_0 D1) D1 {
				return func(_546_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let19_0 D1) D1 {
						return func(_547_dt__update_hcf12_h0 D1) D1 {
							return Companion_D1_.Create_DC5_(_547_dt__update_hcf12_h0)
						}(_pat_let19_0)
					}(_pat_let_tv26)
				}(_pat_let18_0)
			}(_446_v32)
		} else {
			_397_v2 = _dafny.IntOfInt64(781)
			var _548_v97 _dafny.Array
			_ = _548_v97
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
			_ = _nw88
			_548_v97 = _nw88
			var _549_v98 _dafny.Sequence
			_ = _549_v98
			_549_v98 = _dafny.SeqOf(_442_v28, _442_v28)
			var _550_v99 _dafny.Int
			_ = _550_v99
			var _551_v100 _dafny.Array
			_ = _551_v100
			var _552_v101 _dafny.Set
			_ = _552_v101
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Array
			_ = _out7
			var _out8 _dafny.Set
			_ = _out8
			_out6, _out7, _out8 = Companion_Default___.M0(_548_v97, (_dafny.Zero).Minus((func() _dafny.Int {
				if !((_442_v28).F17()) {
					return _dafny.IntOfUint32((_549_v98).Cardinality())
				}
				return _397_v2
			})()), (_442_v28).F17(), Companion_Default___.Fm4((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), (_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool), _405_globalState), _405_globalState)
			_550_v99 = _out6
			_551_v100 = _out7
			_552_v101 = _out8
			_398_v4 = (_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool)
			_442_v28 = (func() *C0 {
				if (_442_v28).F17() {
					return _469_v37
				}
				return _469_v37
			})()
			var _553_v102 _dafny.Set
			_ = _553_v102
			_553_v102 = _dafny.SetOf(_442_v28)
			var _554_v103 _dafny.Map
			_ = _554_v103
			_554_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v102, (_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool))
			var _555_v104 _dafny.Map
			_ = _555_v104
			_555_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4((_554_v103).Cardinality(), (_469_v37).F17(), _405_globalState), _466_v34)
			var _556_v105 _dafny.Map
			_ = _556_v105
			_556_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_549_v98, _466_v34)
			_555_v104 = (_555_v104).Update(Companion_Default___.Fm4(_397_v2, (_442_v28).F17(), _405_globalState), (func() _dafny.Array {
				if (_556_v105).Contains(_549_v98) {
					return (_556_v105).Get(_549_v98).(_dafny.Array)
				}
				return _466_v34
			})())
		}
		(_405_globalState).F11 = _397_v2
		if (true) || (_dafny.Companion_Sequence_.IsPrefixOf(_402_v8, _402_v8)) {
			_398_v4 = (_399_v5).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_401_v7).Cardinality())), _dafny.IntOfUint32((_399_v5).Cardinality()))).Uint32()).(bool)
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index100
			(_466_v34).ArraySet1(!((_469_v37).F17()), (_index100).Int())
			var _557_v106 _dafny.Map
			_ = _557_v106
			_557_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_467_v35).Contains((_442_v28).F17()) {
					return (_467_v35).Get((_442_v28).F17()).(_dafny.Int)
				}
				return _397_v2
			})(), _469_v37)
			var _558_v107 _dafny.Sequence
			_ = _558_v107
			_558_v107 = _dafny.SeqOf(_557_v106)
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index101
			(_466_v34).ArraySet1(!(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_558_v107, _558_v107), (_557_v106).Merge(_557_v106))), (_index101).Int())
			var _559_v108 _dafny.Array
			_ = _559_v108
			var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(8))
			_ = _nw89
			_559_v108 = _nw89
			var _560_v109 _dafny.Int
			_ = _560_v109
			var _561_v110 _dafny.Array
			_ = _561_v110
			var _562_v111 _dafny.Set
			_ = _562_v111
			var _out9 _dafny.Int
			_ = _out9
			var _out10 _dafny.Array
			_ = _out10
			var _out11 _dafny.Set
			_ = _out11
			_out9, _out10, _out11 = Companion_Default___.M0(_559_v108, (_dafny.SetOf((_469_v37).F17())).Cardinality(), ((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Int {
				if (_408_v10).Contains(_397_v2) {
					return (_408_v10).Get(_397_v2).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.SeqOf((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))).Cardinality())
			})()) < 0, (_469_v37).F17(), _405_globalState)
			_560_v109 = _out9
			_561_v110 = _out10
			_562_v111 = _out11
			var _563_v112 _dafny.Map
			_ = _563_v112
			_563_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_465_v33, (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))
			var _564_v115 _dafny.Sequence
			_ = _564_v115
			_564_v115 = _dafny.SeqOf(_465_v33)
			var _565_v116 _dafny.Array
			_ = _565_v116
			var _nwElement0_15 _dafny.Map = _563_v112
			_ = _nwElement0_15
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(8))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_15, 0)
			(_nw90).ArraySet1(_563_v112, 1)
			(_nw90).ArraySet1((_563_v112).Merge(_563_v112), 2)
			(_nw90).ArraySet1(func() _dafny.Map {
				var _coll32 = _dafny.NewMapBuilder()
				_ = _coll32
				for _iter35 := _dafny.Iterate((_563_v112).Keys().Elements()); ; {
					_compr_32, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _566_v113 _dafny.MultiSet
					_566_v113 = interface{}(_compr_32).(_dafny.MultiSet)
					if (_563_v112).Contains(_566_v113) {
						_coll32.Add(_566_v113, _dafny.IntOfInt64(560))
					}
				}
				return _coll32.ToMap()
			}(), 3)
			(_nw90).ArraySet1(Companion_Default___.Fm10((_469_v37).F17(), _405_globalState), 4)
			(_nw90).ArraySet1(func() _dafny.Map {
				var _coll33 = _dafny.NewMapBuilder()
				_ = _coll33
				for _iter36 := _dafny.Iterate((_dafny.MultiSetFromSeq(_564_v115)).Elements()); ; {
					_compr_33, _ok36 := _iter36()
					if !_ok36 {
						break
					}
					var _567_v114 _dafny.MultiSet
					_567_v114 = interface{}(_compr_33).(_dafny.MultiSet)
					if (_dafny.MultiSetFromSeq(_564_v115)).Contains(_567_v114) {
						_coll33.Add(_567_v114, (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))
					}
				}
				return _coll33.ToMap()
			}(), 5)
			(_nw90).ArraySet1((Companion_Default___.Fm10((_442_v28).F17(), _405_globalState)).Merge(_563_v112), 6)
			(_nw90).ArraySet1(_563_v112, 7)
			_565_v116 = _nw90
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_565_v116), 0))
			_ = _index102
			(_565_v116).ArraySet1(_563_v112, (_index102).Int())
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_565_v116), 0))
			_ = _index103
			(_565_v116).ArraySet1(_563_v112, (_index103).Int())
		} else {
			var _568_v117 _dafny.Sequence
			_ = _568_v117
			_568_v117 = _dafny.SeqOf(_466_v34)
			_398_v4 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_466_v34), _568_v117), _568_v117)
			_398_v4 = (_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool)
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index104
			var _rhs102 bool = Companion_Default___.Fm4((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), Companion_Default___.Fm4(_397_v2, (_442_v28).F17(), _405_globalState), _405_globalState)
			_ = _rhs102
			var _rhs103 *C0 = _469_v37
			_ = _rhs103
			var _rhs104 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_397_v2), (_dafny.Zero).Minus(((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Minus(_397_v2)))
			_ = _rhs104
			var _lhs68 _dafny.Array = _466_v34
			_ = _lhs68
			var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _lhs69
			var _lhs70 *GlobalState = _405_globalState
			_ = _lhs70
			(_lhs68).ArraySet1(_rhs102, (_lhs69).Int())
			_442_v28 = _rhs103
			_lhs70.F6 = _rhs104
			var _569_v118 _dafny.Array
			_ = _569_v118
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(8))
			_ = _nw91
			_569_v118 = _nw91
			var _570_v119 _dafny.Set
			_ = _570_v119
			_570_v119 = _dafny.SetOf((_401_v7).Cardinality(), (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int))
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_569_v118), 0))
			_ = _index105
			(_569_v118).ArraySet1(_570_v119, (_index105).Int())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_569_v118), 0))
			_ = _index106
			(_569_v118).ArraySet1((_570_v119).Union(_570_v119), (_index106).Int())
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))
			_ = _index107
			(_466_v34).ArraySet1(((_dafny.MultiSetFromSeq(_399_v5)).Difference(_465_v33)).IsSubsetOf(_465_v33), (_index107).Int())
		}
	}
	var _hi2 _dafny.Int = _397_v2
	_ = _hi2
	for _571_i14 := (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int); _571_i14.Cmp(_hi2) < 0; _571_i14 = _571_i14.Plus(_dafny.One) {
		_398_v4 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("to"), (func() _dafny.Sequence {
			if Companion_Default___.Fm4((_dafny.Zero).Minus(_571_i14), true, _405_globalState) {
				return _402_v8
			}
			return _402_v8
		})())
		var _572_v120 _dafny.Array
		_ = _572_v120
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_10
		var _nw92 _dafny.Array
		_ = _nw92
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw92 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.MultiSet = (func(_573_v9 _dafny.Array) func(_dafny.Int) _dafny.MultiSet {
				return func(_574_i15 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf((_573_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_573_v9), 0))).Int()).(_dafny.Int))
				}
			})(_403_v9)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw92 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw92).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw92).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_572_v120 = _nw92
		var _575_v121 _dafny.Int
		_ = _575_v121
		var _576_v122 _dafny.Array
		_ = _576_v122
		var _577_v123 _dafny.Set
		_ = _577_v123
		var _out12 _dafny.Int
		_ = _out12
		var _out13 _dafny.Array
		_ = _out13
		var _out14 _dafny.Set
		_ = _out14
		_out12, _out13, _out14 = Companion_Default___.M0(_572_v120, (_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int), ((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)).Cmp((_403_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_403_v9), 0))).Int()).(_dafny.Int)) != 0, _dafny.Companion_Sequence_.Contains(_402_v8, _dafny.CodePoint('t')), _405_globalState)
		_575_v121 = _out12
		_576_v122 = _out13
		_577_v123 = _out14
		var _578_v124 _dafny.Int
		_ = _578_v124
		var _579_v125 _dafny.Array
		_ = _579_v125
		var _580_v126 _dafny.Set
		_ = _580_v126
		var _out15 _dafny.Int
		_ = _out15
		var _out16 _dafny.Array
		_ = _out16
		var _out17 _dafny.Set
		_ = _out17
		_out15, _out16, _out17 = Companion_Default___.M0(_572_v120, _575_v121, (_401_v7).IsDisjointFrom(_dafny.SetOf((_466_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_466_v34), 0))).Int()).(bool))), (_442_v28).F17(), _405_globalState)
		_578_v124 = _out15
		_579_v125 = _out16
		_580_v126 = _out17
		var _581_v127 _dafny.Int
		_ = _581_v127
		var _582_v128 _dafny.Array
		_ = _582_v128
		var _583_v129 _dafny.Set
		_ = _583_v129
		var _out18 _dafny.Int
		_ = _out18
		var _out19 _dafny.Array
		_ = _out19
		var _out20 _dafny.Set
		_ = _out20
		_out18, _out19, _out20 = Companion_Default___.M0(_572_v120, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_402_v8).Cardinality()), _571_i14), (_442_v28).F17(), (_442_v28).F17(), _405_globalState)
		_581_v127 = _out18
		_582_v128 = _out19
		_583_v129 = _out20
	}
	_dafny.Print(_396_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_397_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_398_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_399_v5, _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_400_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_401_v7).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_402_v8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_403_v9).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_403_v9).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_403_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_403_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_403_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_403_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_403_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_405_globalState).F2()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _dafny.IntOfInt64(514))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_405_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_405_globalState).F7()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(514), _dafny.IntOfInt64(-1), _dafny.IntOfInt64(2), _dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_405_globalState.F11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState.F14).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState.F14).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState.F14).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState.F14).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState.F14).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState.F14).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState.F14).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_405_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_405_globalState.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_408_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(514), _dafny.IntOfInt64(514)).UpdateUnsafe(_dafny.IntOfInt64(-514), _dafny.IntOfInt64(514))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_442_v28).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_443_v29).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(514))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_444_v30).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.MultiSetOf(_dafny.IntOfInt64(76)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_445_v31).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_445_v31).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_445_v31).Dtor_cf11()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.MultiSetOf(_dafny.IntOfInt64(76)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_446_v32).Dtor_cf12()).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_446_v32).Dtor_cf12()).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_446_v32).Dtor_cf12()).Dtor_cf11()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.MultiSetOf(_dafny.IntOfInt64(76)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_465_v33).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_466_v34).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_466_v34).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_467_v35).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_468_v36, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(514), _dafny.IntOfInt64(514)).UpdateUnsafe(_dafny.IntOfInt64(-514), _dafny.IntOfInt64(514)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(514), _dafny.IntOfInt64(514)).UpdateUnsafe(_dafny.IntOfInt64(-514), _dafny.IntOfInt64(514)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_469_v37).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_470_v38).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.UnicodeSeqOfUtf8Bytes("frrp")).UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.UnicodeSeqOfUtf8Bytes("frrp"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(-874))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_472_v39).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(3355), _dafny.IntOfInt64(514), _dafny.IntOfInt64(129))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_474_i6)
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
	Cf1 bool
	Cf2 _dafny.Sequence
	Cf3 _dafny.Map
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int, Cf1 bool, Cf2 _dafny.Sequence, Cf3 _dafny.Map) D0 {
	return D0{D0_DC0{Cf0, Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero, false, _dafny.EmptySeq, _dafny.EmptyMap)
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ", " + data.Cf2.VerbatimString(true) + ", " + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1 == data2.Cf1 && data1.Cf2.Equals(data2.Cf2) && data1.Cf3.Equals(data2.Cf3)
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
	Cf5 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf5 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf5}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC3 struct {
	Cf6 _dafny.Map
	Cf7 _dafny.Int
	Cf8 D0
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 _dafny.Map, Cf7 _dafny.Int, Cf8 D0) D1 {
	return D1{D1_DC3{Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf9  _dafny.Int
	Cf10 bool
	Cf11 _dafny.Map
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf9 _dafny.Int, Cf10 bool, Cf11 _dafny.Map) D1 {
	return D1{D1_DC4{Cf9, Cf10, Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC1 struct {
	Cf4 _dafny.Sequence
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf4 _dafny.Sequence) D1 {
	return D1{D1_DC1{Cf4}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

type D1_DC5 struct {
	Cf12 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf12 D1) D1 {
	return D1{D1_DC5{Cf12}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.CodePoint('D'))
}

func (_this D1) Dtor_cf5() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Map {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf8() D0 {
	return _this.Get_().(D1_DC3).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC1).Cf4
}

func (_this D1) Dtor_cf12() D1 {
	return _this.Get_().(D1_DC5).Cf12
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf5 == data2.Cf5
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Equals(data2.Cf8)
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11.Equals(data2.Cf11)
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf14 _dafny.Int
	Cf15 _dafny.CodePoint
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf14 _dafny.Int, Cf15 _dafny.CodePoint) D2 {
	return D2{D2_DC7{Cf14, Cf15}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf16 _dafny.Array
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf16 _dafny.Array) D2 {
	return D2{D2_DC8{Cf16}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC9 struct {
	Cf17 _dafny.Int
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf17 _dafny.Int) D2 {
	return D2{D2_DC9{Cf17}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC6 struct {
	Cf13 _dafny.Array
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf13 _dafny.Array) D2 {
	return D2{D2_DC6{Cf13}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) Dtor_cf15() _dafny.CodePoint {
	return _this.Get_().(D2_DC7).Cf15
}

func (_this D2) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D2_DC8).Cf16
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf17
}

func (_this D2) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D2_DC6).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf16 == data2.Cf16
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf13 == data2.Cf13
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
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_() D3 {
	return D3{D3_DC11{}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC10 struct {
	Cf18 _dafny.Sequence
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf18 _dafny.Sequence) D3 {
	return D3{D3_DC10{Cf18}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC11_()
}

func (_this D3) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf18) + ")"
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
			_, ok := other.Get_().(D3_DC11)
			return ok
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
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

type D4_DC13 struct {
	Cf20 bool
	Cf21 _dafny.Int
	Cf22 _dafny.Array
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf20 bool, Cf21 _dafny.Int, Cf22 _dafny.Array) D4 {
	return D4{D4_DC13{Cf20, Cf21, Cf22}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf23 _dafny.Int
	Cf24 _dafny.Int
	Cf25 _dafny.Set
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf23 _dafny.Int, Cf24 _dafny.Int, Cf25 _dafny.Set) D4 {
	return D4{D4_DC14{Cf23, Cf24, Cf25}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC15 struct {
	Cf26 _dafny.Set
	Cf27 _dafny.CodePoint
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf26 _dafny.Set, Cf27 _dafny.CodePoint) D4 {
	return D4{D4_DC15{Cf26, Cf27}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

type D4_DC12 struct {
	Cf19 _dafny.Map
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf19 _dafny.Map) D4 {
	return D4{D4_DC12{Cf19}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC16 struct {
	Cf28 D4
}

func (D4_DC16) isD4() {}

func (CompanionStruct_D4_) Create_DC16_(Cf28 D4) D4 {
	return D4{D4_DC16{Cf28}}
}

func (_this D4) Is_DC16() bool {
	_, ok := _this.Get_().(D4_DC16)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC13).Cf20
}

func (_this D4) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf21
}

func (_this D4) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D4_DC14).Cf25
}

func (_this D4) Dtor_cf26() _dafny.Set {
	return _this.Get_().(D4_DC15).Cf26
}

func (_this D4) Dtor_cf27() _dafny.CodePoint {
	return _this.Get_().(D4_DC15).Cf27
}

func (_this D4) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) Dtor_cf28() D4 {
	return _this.Get_().(D4_DC16).Cf28
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC16:
		{
			return "D4.DC16" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22 == data2.Cf22
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Equals(data2.Cf25)
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf26.Equals(data2.Cf26) && data1.Cf27 == data2.Cf27
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf19.Equals(data2.Cf19)
		}
	case D4_DC16:
		{
			data2, ok := other.Get_().(D4_DC16)
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

type D5_DC18 struct {
	Cf30 _dafny.Int
	Cf31 _dafny.Int
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf30 _dafny.Int, Cf31 _dafny.Int) D5 {
	return D5{D5_DC18{Cf30, Cf31}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC17 struct {
	Cf29 _dafny.Set
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf29 _dafny.Set) D5 {
	return D5{D5_DC17{Cf29}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC19 struct {
	Cf32 D5
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf32 D5) D5 {
	return D5{D5_DC19{Cf32}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC18_(_dafny.Zero, _dafny.Zero)
}

func (_this D5) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf30
}

func (_this D5) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf31
}

func (_this D5) Dtor_cf29() _dafny.Set {
	return _this.Get_().(D5_DC17).Cf29
}

func (_this D5) Dtor_cf32() D5 {
	return _this.Get_().(D5_DC19).Cf32
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf32) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf32.Equals(data2.Cf32)
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

type D6_DC21 struct {
	Cf34 bool
	Cf35 _dafny.CodePoint
	Cf36 _dafny.Int
	Cf37 _dafny.Int
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_(Cf34 bool, Cf35 _dafny.CodePoint, Cf36 _dafny.Int, Cf37 _dafny.Int) D6 {
	return D6{D6_DC21{Cf34, Cf35, Cf36, Cf37}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

type D6_DC20 struct {
	Cf33 _dafny.Map
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf33 _dafny.Map) D6 {
	return D6{D6_DC20{Cf33}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC21_(false, _dafny.CodePoint('D'), _dafny.Zero, _dafny.Zero)
}

func (_this D6) Dtor_cf34() bool {
	return _this.Get_().(D6_DC21).Cf34
}

func (_this D6) Dtor_cf35() _dafny.CodePoint {
	return _this.Get_().(D6_DC21).Cf35
}

func (_this D6) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D6_DC21).Cf36
}

func (_this D6) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D6_DC21).Cf37
}

func (_this D6) Dtor_cf33() _dafny.Map {
	return _this.Get_().(D6_DC20).Cf33
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC21:
		{
			return "D6.DC21" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC21:
		{
			data2, ok := other.Get_().(D6_DC21)
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35 && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf33.Equals(data2.Cf33)
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

// Definition of class GlobalState
type GlobalState struct {
	F6   _dafny.Int
	F11  _dafny.Int
	F14  _dafny.Array
	F16  _dafny.Int
	_f0  _dafny.Int
	_f1  _dafny.Int
	_f2  _dafny.Map
	_f3  bool
	_f4  _dafny.Int
	_f5  _dafny.Int
	_f7  _dafny.MultiSet
	_f8  bool
	_f9  _dafny.Int
	_f10 _dafny.Int
	_f12 _dafny.Int
	_f13 bool
	_f15 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F6 = _dafny.Zero
	_this.F11 = _dafny.Zero
	_this.F14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F16 = _dafny.Zero
	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.EmptyMap
	_this._f3 = false
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f7 = _dafny.EmptyMultiSet
	_this._f8 = false
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.Zero
	_this._f12 = _dafny.Zero
	_this._f13 = false
	_this._f15 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 _dafny.Map, f3 bool, f4 _dafny.Int, f5 _dafny.Int, f6 _dafny.Int, f7 _dafny.MultiSet, f8 bool, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Int, f12 _dafny.Int, f13 bool, f14 _dafny.Array, f15 _dafny.Int, f16 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this).F16 = f16
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Map {
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
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F7() _dafny.MultiSet {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f17 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f17 = false
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

func (_this *C0) Ctor__(f17 bool) {
	{
		(_this)._f17 = f17
	}
}
func (_this *C0) Fm2(p0 bool, p1 bool, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(396))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_584_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jcohubbi"), _dafny.UnicodeSeqOfUtf8Bytes("bfaki")))
	}
}
func (_this *C0) F17() bool {
	{
		return _this._f17
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
