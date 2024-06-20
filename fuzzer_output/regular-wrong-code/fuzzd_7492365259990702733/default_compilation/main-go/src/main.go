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
	return (_dafny.IntOfInt64(529)).Plus((_dafny.IntOfInt64(64)).Minus((_dafny.SetOf((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(695))).Cardinality(), _dafny.IntOfInt64(768))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	return ((func() bool {
		if false {
			return true
		}
		return true
	})()) || (true)
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.MultiSetOf(true, true), _dafny.MultiSetFromSeq((_dafny.SeqOf(true))), _dafny.MultiSetOf(true))).Cardinality(), Companion_D0_.Create_DC2_(true)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(380))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		}))).Cardinality()))).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Map
			_1_v0 = interface{}(_compr_0).(_dafny.Map)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.MultiSetOf(true, true), _dafny.MultiSetFromSeq((_dafny.SeqOf(true))), _dafny.MultiSetOf(true))).Cardinality(), Companion_D0_.Create_DC2_(true)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(380))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()))).Contains(_1_v0) {
				_coll0.Add(_1_v0, _dafny.IntOfInt64(329))
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), Companion_D0_.Create_DC2_(false)), _dafny.IntOfUint32((_dafny.SeqOf(false, true, false, false, false)).Cardinality())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-222), Companion_D0_.Create_DC2_(true)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("swxgyamao")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(679))).Cardinality(), Companion_D0_.Create_DC2_(true)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(273), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(511))).Cardinality()))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(127), !(false))).Cardinality())).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_(!(false) || (!(true)))
}
func (_static *CompanionStruct_Default___) Fm4(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((_dafny.SeqOf(!(!(false)))), _dafny.SeqOf(true, true))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-864), _dafny.IntOfInt64(937))).Cardinality()), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()), _dafny.IntOfInt64(311))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hxltqisx"), _dafny.UnicodeSeqOfUtf8Bytes("tbiom"))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	if false {
		return _dafny.CodePoint('h')
	} else {
		return _dafny.CodePoint('v')
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.IntOfInt64(566)), Companion_D0_.Create_DC0_(_dafny.IntOfInt64(607))), _dafny.SeqOf(Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('g'))).Cardinality()), Companion_D0_.Create_DC0_((_dafny.SetOf(true, false, false, false, !(false))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bb"), _dafny.UnicodeSeqOfUtf8Bytes("mwmmnrp"), _dafny.UnicodeSeqOfUtf8Bytes("obby"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(247))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v0 _dafny.Sequence
			_3_v0 = interface{}(_compr_1).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bb"), _dafny.UnicodeSeqOfUtf8Bytes("mwmmnrp"), _dafny.UnicodeSeqOfUtf8Bytes("obby"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(247))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}))), _3_v0) {
				_coll1.Add(_3_v0, (_dafny.IntOfInt64(405)).Plus(_dafny.IntOfInt64(921)))
			}
		}
		return _coll1.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Intersection(_dafny.SetOf(true))).Union(_dafny.SetOf(!(true), false, false))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-131)).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bqfasi")).Cardinality())), !(true) || (!(true)))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 D1, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((true) == (false))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(771), false)).Contains(_dafny.IntOfInt64(531)) {
		return func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(942))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_4_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-874)
			}))).Elements()); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _5_v0 _dafny.Int
				_5_v0 = interface{}(_compr_2).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(942))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg5 _dafny.Int) interface{} {
						return coer5(arg5)
					}
				}(func(_4_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-874)
				})), _5_v0) {
					_coll2.Add((_5_v0).Times(_dafny.IntOfInt64(871)), (func() _dafny.Map {
						var _coll3 = _dafny.NewMapBuilder()
						_ = _coll3
						for _iter3 := _dafny.Iterate(((Companion_D8_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(760), func() _dafny.Map {
							var _coll4 = _dafny.NewMapBuilder()
							_ = _coll4
							for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(976), _dafny.IntOfInt64(745))); ; {
								_compr_4, _ok4 := _iter4()
								if !_ok4 {
									break
								}
								var _6_v2 _dafny.Int
								_6_v2 = interface{}(_compr_4).(_dafny.Int)
								if ((_dafny.IntOfInt64(976)).Cmp(_6_v2) <= 0) && ((_6_v2).Cmp(_dafny.IntOfInt64(745)) < 0) {
									_coll4.Add(Companion_Default___.SafeModuloInt(_6_v2, _dafny.IntOfInt64(710)), _dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()))
								}
							}
							return _coll4.ToMap()
						}()))).Dtor_cf32()).Keys().Elements()); ; {
							_compr_3, _ok3 := _iter3()
							if !_ok3 {
								break
							}
							var _7_v1 _dafny.Int
							_7_v1 = interface{}(_compr_3).(_dafny.Int)
							if ((Companion_D8_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(760), func() _dafny.Map {
								var _coll5 = _dafny.NewMapBuilder()
								_ = _coll5
								for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(976), _dafny.IntOfInt64(745))); ; {
									_compr_5, _ok5 := _iter5()
									if !_ok5 {
										break
									}
									var _6_v2 _dafny.Int
									_6_v2 = interface{}(_compr_5).(_dafny.Int)
									if ((_dafny.IntOfInt64(976)).Cmp(_6_v2) <= 0) && ((_6_v2).Cmp(_dafny.IntOfInt64(745)) < 0) {
										_coll5.Add(Companion_Default___.SafeModuloInt(_6_v2, _dafny.IntOfInt64(710)), _dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()))
									}
								}
								return _coll5.ToMap()
							}()))).Dtor_cf32()).Contains(_7_v1) {
								_coll3.Add((_7_v1).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
									var _coll6 = _dafny.NewBuilder()
									_ = _coll6
									for _iter6 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('y'), _dafny.CodePoint('x'))).Elements()); ; {
										_compr_6, _ok6 := _iter6()
										if !_ok6 {
											break
										}
										var _8_v3 _dafny.CodePoint
										_8_v3 = interface{}(_compr_6).(_dafny.CodePoint)
										if (_dafny.MultiSetOf(_dafny.CodePoint('y'), _dafny.CodePoint('x'))).Contains(_8_v3) {
											_coll6.Add(_8_v3)
										}
									}
									return _coll6.ToSet()
								}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(228))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg6 _dafny.Int) interface{} {
										return coer6(arg6)
									}
								}(func(_9_i1 _dafny.Int) _dafny.Int {
									return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mamhl")).Cardinality())
								}))).Cardinality()))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), true))
							}
						}
						return _coll3.ToMap()
					}()).Cardinality())
				}
			}
			return _coll2.ToMap()
		}()
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(95), _dafny.IntOfInt64(-232))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(967))).Cardinality(), _dafny.IntOfInt64(-697)))
	}
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(false), _dafny.SeqOf(!(false)))).Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _10_v0 _dafny.Sequence
			_10_v0 = interface{}(_compr_7).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.SeqOf(false), _dafny.SeqOf(!(false)))).Contains(_10_v0) {
				_coll7.Add(_10_v0, ((_dafny.SetOf(true, true)).Cardinality()).Cmp(_dafny.IntOfInt64(-608)) == 0)
			}
		}
		return _coll7.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).IsDisjointFrom(_dafny.MultiSetOf(true, false)), _dafny.CodePoint('d'))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, globalState *GlobalState) _dafny.Array {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var _hi0 _dafny.Int = p0
	_ = _hi0
	for _11_i0 := _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality()); _11_i0.Cmp(_hi0) < 0; _11_i0 = _11_i0.Plus(_dafny.One) {
		var _12_v0 _dafny.Set
		_ = _12_v0
		_12_v0 = _dafny.SetOf((_dafny.SetOf(_11_i0)).Cardinality(), p0, p0)
		var _13_v1 bool
		_ = _13_v1
		_13_v1 = false
		var _14_v2 *C0
		_ = _14_v2
		var _nw0 *C0 = New_C0_()
		_ = _nw0
		_nw0.Ctor__()
		_14_v2 = _nw0
		var _15_v3 _dafny.Map
		_ = _15_v3
		_15_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v2, _dafny.IntOfInt64(803))
		var _16_v4 D3
		_ = _16_v4
		_16_v4 = Companion_D3_.Create_DC7_(_15_v3)
		var _17_v5 _dafny.Map
		_ = _17_v5
		_17_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v1, _16_v4)
		var _rhs0 _dafny.Set = func() _dafny.Set {
			var _coll8 = _dafny.NewBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate(((_12_v0).Union(_12_v0)).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _18_v6 _dafny.Int
				_18_v6 = interface{}(_compr_8).(_dafny.Int)
				if ((_12_v0).Union(_12_v0)).Contains(_18_v6) {
					_coll8.Add((_18_v6).Plus((_dafny.IntOfInt64(201)).Plus((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-511))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg7 _dafny.Int) interface{} {
							return coer7(arg7)
						}
					}(func(_19_i1 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(-128)
					})))).Cardinality())))
				}
			}
			return _coll8.ToSet()
		}()
		_ = _rhs0
		var _rhs1 _dafny.Map = _17_v5
		_ = _rhs1
		_12_v0 = _rhs0
		_17_v5 = _rhs1
		(globalState).F11 = (_dafny.Zero).Minus(_11_i0)
		var _20_v7 _dafny.Array
		_ = _20_v7
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_21_v1 bool) func(_dafny.Int) bool {
				return func(_22_i2 _dafny.Int) bool {
					return _21_v1
				}
			})(_13_v1)
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
		_20_v7 = _nw1
		var _23_v8 _dafny.CodePoint
		_ = _23_v8
		_23_v8 = _dafny.CodePoint('t')
		var _24_v9 _dafny.Map
		_ = _24_v9
		_24_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v8, _20_v7)
		var _25_v10 D0
		_ = _25_v10
		_25_v10 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(443), _dafny.IntOfInt64(315), p0, p0, (func() _dafny.Array {
			if (_24_v9).Contains(_23_v8) {
				return (_24_v9).Get(_23_v8).(_dafny.Array)
			}
			return _20_v7
		})())
		var _26_v11 _dafny.Map
		_ = _26_v11
		_26_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v10, _20_v7)
		var _27_v12 _dafny.Array
		_ = _27_v12
		var _nwElement0_0 _dafny.Array = _20_v7
		_ = _nwElement0_0
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(17))
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_0, 0)
		(_nw2).ArraySet1(_20_v7, 1)
		(_nw2).ArraySet1(_20_v7, 2)
		(_nw2).ArraySet1(_20_v7, 3)
		(_nw2).ArraySet1(_20_v7, 4)
		(_nw2).ArraySet1((func() _dafny.Array {
			if (_26_v11).Contains(_25_v10) {
				return (_26_v11).Get(_25_v10).(_dafny.Array)
			}
			return _20_v7
		})(), 5)
		(_nw2).ArraySet1(_20_v7, 6)
		(_nw2).ArraySet1(_20_v7, 7)
		(_nw2).ArraySet1(_20_v7, 8)
		(_nw2).ArraySet1(_20_v7, 9)
		(_nw2).ArraySet1(_20_v7, 10)
		(_nw2).ArraySet1(_20_v7, 11)
		(_nw2).ArraySet1(_20_v7, 12)
		(_nw2).ArraySet1(_20_v7, 13)
		(_nw2).ArraySet1(_20_v7, 14)
		(_nw2).ArraySet1((_25_v10).Dtor_cf5(), 15)
		(_nw2).ArraySet1(_20_v7, 16)
		_27_v12 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_27_v12), 0))
		_ = _index0
		(_27_v12).ArraySet1((_25_v10).Dtor_cf5(), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_27_v12), 0))
		_ = _index1
		(_27_v12).ArraySet1(_20_v7, (_index1).Int())
		var _28_v13 _dafny.Array
		_ = _28_v13
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
		_ = _nw3
		_28_v13 = _nw3
		r0 = _28_v13
	}
	var _29_v14 _dafny.Map
	_ = _29_v14
	_29_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
	var _30_v15 bool
	_ = _30_v15
	_30_v15 = false
	var _31_v16 _dafny.Sequence
	_ = _31_v16
	_31_v16 = _dafny.SeqOf(false, !(_30_v15))
	var _32_v17 _dafny.Set
	_ = _32_v17
	_32_v17 = _dafny.SetOf(p0)
	var _33_v18 _dafny.Array
	_ = _33_v18
	var _nwElement0_1 bool = _30_v15
	_ = _nwElement0_1
	var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(16))
	_ = _nw4
	(_nw4).ArraySet1(_nwElement0_1, 0)
	(_nw4).ArraySet1(false, 1)
	(_nw4).ArraySet1(_30_v15, 2)
	(_nw4).ArraySet1(false, 3)
	(_nw4).ArraySet1(_30_v15, 4)
	(_nw4).ArraySet1(_30_v15, 5)
	(_nw4).ArraySet1(_30_v15, 6)
	(_nw4).ArraySet1((func() bool {
		if (_29_v14).Contains(p0) {
			return (_29_v14).Get(p0).(bool)
		}
		return _30_v15
	})(), 7)
	(_nw4).ArraySet1(!(_30_v15), 8)
	(_nw4).ArraySet1(_30_v15, 9)
	(_nw4).ArraySet1(_30_v15, 10)
	(_nw4).ArraySet1(!(_30_v15), 11)
	(_nw4).ArraySet1(_30_v15, 12)
	(_nw4).ArraySet1(!(_30_v15), 13)
	(_nw4).ArraySet1(!(_30_v15), 14)
	(_nw4).ArraySet1((_31_v16).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_32_v17).Cardinality()), _dafny.IntOfUint32((_31_v16).Cardinality()))).Uint32()).(bool), 15)
	_33_v18 = _nw4
	var _34_v19 D0
	_ = _34_v19
	_34_v19 = Companion_D0_.Create_DC1_(Companion_Default___.Fm0(globalState), ((func() _dafny.Map {
		if false {
			return _29_v14
		}
		return _29_v14
	})()).Cardinality(), p0, p0, _33_v18)
	var _source0 D0 = _34_v19
	_ = _source0
	if _source0.Is_DC1() {
		var _35___mcc_h0 _dafny.Int = _source0.Get_().(D0_DC1).Cf1
		_ = _35___mcc_h0
		var _36___mcc_h1 _dafny.Int = _source0.Get_().(D0_DC1).Cf2
		_ = _36___mcc_h1
		var _37___mcc_h2 _dafny.Int = _source0.Get_().(D0_DC1).Cf3
		_ = _37___mcc_h2
		var _38___mcc_h3 _dafny.Int = _source0.Get_().(D0_DC1).Cf4
		_ = _38___mcc_h3
		var _39___mcc_h4 _dafny.Array = _source0.Get_().(D0_DC1).Cf5
		_ = _39___mcc_h4
		var _40_cf5 _dafny.Array = _39___mcc_h4
		_ = _40_cf5
		var _41_cf4 _dafny.Int = _38___mcc_h3
		_ = _41_cf4
		var _42_cf3 _dafny.Int = _37___mcc_h2
		_ = _42_cf3
		var _43_cf2 _dafny.Int = _36___mcc_h1
		_ = _43_cf2
		var _44_cf1 _dafny.Int = _35___mcc_h0
		_ = _44_cf1
		var _45_v20 _dafny.Set
		_ = _45_v20
		_45_v20 = _dafny.SetOf(true)
		var _46_v21 *C1
		_ = _46_v21
		var _nw5 *C1 = New_C1_()
		_ = _nw5
		_nw5.Ctor__(((_45_v20).Cardinality()).Minus(_43_cf2))
		_46_v21 = _nw5
		var _47_v22 _dafny.Array
		_ = _47_v22
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_1
		var _nw6 _dafny.Array
		_ = _nw6
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw6 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Int = (func(_48_v15 bool, _49_v16 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_50_i3 _dafny.Int) _dafny.Int {
					return (_50_i3).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v15, _dafny.IntOfUint32((_49_v16).Cardinality()))).Cardinality())
				}
			})(_30_v15, _31_v16)
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
		_47_v22 = _nw6
		r0 = _47_v22
		_41_cf4 = _46_v21.F18
		(globalState).F12 = _30_v15
	} else if _source0.Is_DC2() {
		var _51___mcc_h5 bool = _source0.Get_().(D0_DC2).Cf6
		_ = _51___mcc_h5
		var _52_cf6 bool = _51___mcc_h5
		_ = _52_cf6
		var _53_v23 *C1
		_ = _53_v23
		var _nw7 *C1 = New_C1_()
		_ = _nw7
		_nw7.Ctor__(_dafny.IntOfInt64(-998))
		_53_v23 = _nw7
		_53_v23 = _53_v23
		_33_v18 = _33_v18
		_52_cf6 = (func() bool {
			if !(_30_v15) || (_30_v15) {
				return _52_cf6
			}
			return (p0).Cmp(p0) <= 0
		})()
		var _54_v24 *C0
		_ = _54_v24
		var _nw8 *C0 = New_C0_()
		_ = _nw8
		_nw8.Ctor__()
		_54_v24 = _nw8
	} else if _source0.Is_DC3() {
		var _55___mcc_h6 _dafny.Int = _source0.Get_().(D0_DC3).Cf7
		_ = _55___mcc_h6
		var _56___mcc_h7 bool = _source0.Get_().(D0_DC3).Cf8
		_ = _56___mcc_h7
		var _57___mcc_h8 _dafny.Int = _source0.Get_().(D0_DC3).Cf9
		_ = _57___mcc_h8
		var _58___mcc_h9 _dafny.Int = _source0.Get_().(D0_DC3).Cf10
		_ = _58___mcc_h9
		var _59_cf10 _dafny.Int = _58___mcc_h9
		_ = _59_cf10
		var _60_cf9 _dafny.Int = _57___mcc_h8
		_ = _60_cf9
		var _61_cf8 bool = _56___mcc_h7
		_ = _61_cf8
		var _62_cf7 _dafny.Int = _55___mcc_h6
		_ = _62_cf7
		var _63_v25 _dafny.Sequence
		_ = _63_v25
		_63_v25 = _dafny.UnicodeSeqOfUtf8Bytes("xbjalwtvy")
		var _64_v26 D0
		_ = _64_v26
		_64_v26 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_63_v25).Cardinality()))
		_64_v26 = _64_v26
		(globalState).F12 = _dafny.Companion_Sequence_.IsPrefixOf(_63_v25, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm8(globalState), (Companion_Default___.SafeIndex(_62_cf7, _dafny.IntOfUint32((Companion_Default___.Fm8(globalState)).Cardinality()))).Uint32(), _dafny.CodePoint('l')), _63_v25))
		var _65_v27 _dafny.Set
		_ = _65_v27
		_65_v27 = _dafny.SetOf(_30_v15, (_31_v16).Select((Companion_Default___.SafeIndex(_62_cf7, _dafny.IntOfUint32((_31_v16).Cardinality()))).Uint32()).(bool), _61_cf8, _30_v15, _30_v15)
		var _66_v28 _dafny.CodePoint
		_ = _66_v28
		_66_v28 = _dafny.CodePoint('s')
		_65_v27 = Companion_Default___.Fm12(_66_v28, p0, globalState)
		(globalState).F12 = _61_cf8
	} else {
		var _67___mcc_h10 _dafny.Int = _source0.Get_().(D0_DC0).Cf0
		_ = _67___mcc_h10
		var _68_cf0 _dafny.Int = _67___mcc_h10
		_ = _68_cf0
		_30_v15 = _30_v15
		var _69_v29 _dafny.Sequence
		_ = _69_v29
		_69_v29 = _dafny.UnicodeSeqOfUtf8Bytes("dw")
		var _70_v30 _dafny.Sequence
		_ = _70_v30
		_70_v30 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("igmkken"), _69_v29)
		if ((_31_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.IntOfUint32((_31_v16).Cardinality()))).Uint32()).(bool)) == (_dafny.Companion_Sequence_.Contains(_70_v30, _69_v29)) {
			var _71_v31 _dafny.Sequence
			_ = _71_v31
			_71_v31 = _dafny.SeqOf(p0, _68_cf0)
			var _72_v32 _dafny.Map
			_ = _72_v32
			_72_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v15, _71_v31)
			var _73_v33 D3
			_ = _73_v33
			_73_v33 = Companion_D3_.Create_DC8_(_71_v31, true, _30_v15, _69_v29, _30_v15)
			_72_v32 = (_72_v32).Update(Companion_Default___.Fm1(_30_v15, _30_v15, p0, globalState), (_73_v33).Dtor_cf16())
			var _74_v34 _dafny.MultiSet
			_ = _74_v34
			_74_v34 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_69_v29, _69_v29)).Cardinality()), (func() _dafny.Int {
				if _30_v15 {
					return _68_cf0
				}
				return p0
			})(), (p0).Times(p0), p0, p0)
			var _75_v35 _dafny.Map
			_ = _75_v35
			_75_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_30_v15, _30_v15, p0, globalState), (_68_cf0).Cmp(Companion_Default___.Fm0(globalState)) < 0)
			var _76_v36 _dafny.Set
			_ = _76_v36
			_76_v36 = _dafny.SetOf(_30_v15)
			var _77_v37 _dafny.MultiSet
			_ = _77_v37
			_77_v37 = _dafny.MultiSetOf(_31_v16)
			var _78_v38 _dafny.Map
			_ = _78_v38
			_78_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v15, _77_v37)
			var _79_v39 _dafny.Map
			_ = _79_v39
			_79_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_76_v36).Difference(_76_v36), ((func() _dafny.MultiSet {
				if (_78_v38).Contains(_30_v15) {
					return (_78_v38).Get(_30_v15).(_dafny.MultiSet)
				}
				return _77_v37
			})()).IsDisjointFrom(_77_v37))
			var _80_v40 _dafny.CodePoint
			_ = _80_v40
			_80_v40 = _dafny.CodePoint('b')
			var _81_v41 _dafny.Set
			_ = _81_v41
			_81_v41 = _dafny.SetOf(_80_v40)
			var _82_v42 D0
			_ = _82_v42
			_82_v42 = Companion_D0_.Create_DC2_(_30_v15)
			var _rhs2 bool = (func() bool {
				if (_79_v39).Contains(_76_v36) {
					return (_79_v39).Get(_76_v36).(bool)
				}
				return _30_v15
			})()
			_ = _rhs2
			var _rhs3 _dafny.MultiSet = _74_v34
			_ = _rhs3
			var _rhs4 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v15, !(_30_v15))
			_ = _rhs4
			var _rhs5 bool = (_81_v41).Equals((_dafny.SetOf(_80_v40, _80_v40, _80_v40, _80_v40, _80_v40)).Difference(_81_v41))
			_ = _rhs5
			var _rhs6 _dafny.Int = (func() _dafny.Int {
				if ((_82_v42).Dtor_cf6()) && (_30_v15) {
					return _68_cf0
				}
				return _68_cf0
			})()
			_ = _rhs6
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			var _lhs1 *GlobalState = globalState
			_ = _lhs1
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			_lhs0.F12 = _rhs2
			_74_v34 = _rhs3
			_75_v35 = _rhs4
			_lhs1.F12 = _rhs5
			_lhs2.F11 = _rhs6
			var _83_v43 *C0
			_ = _83_v43
			var _nw9 *C0 = New_C0_()
			_ = _nw9
			_nw9.Ctor__()
			_83_v43 = _nw9
			var _84_v44 D4
			_ = _84_v44
			_84_v44 = Companion_D4_.Create_DC9_(_83_v43)
			var _85_v45 D4
			_ = _85_v45
			_85_v45 = Companion_D4_.Create_DC11_(_84_v44)
			var _86_v46 D4
			_ = _86_v46
			_86_v46 = Companion_D4_.Create_DC11_(_84_v44)
			var _87_v47 D4
			_ = _87_v47
			_87_v47 = Companion_D4_.Create_DC11_(_86_v46)
			var _rhs7 _dafny.Int = _dafny.IntOfUint32((_69_v29).Cardinality())
			_ = _rhs7
			var _rhs8 D4 = _87_v47
			_ = _rhs8
			var _rhs9 _dafny.CodePoint = _80_v40
			_ = _rhs9
			var _lhs3 *GlobalState = globalState
			_ = _lhs3
			_lhs3.F11 = _rhs7
			_87_v47 = _rhs8
			_80_v40 = _rhs9
			var _88_v48 _dafny.Array
			_ = _88_v48
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw10
			_88_v48 = _nw10
			var _89_v49 _dafny.Sequence
			_ = _89_v49
			_89_v49 = _dafny.SeqOf(_81_v41)
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_88_v48), 0))
			_ = _index2
			(_88_v48).ArraySet1(_89_v49, (_index2).Int())
			var _90_v50 _dafny.MultiSet
			_ = _90_v50
			_90_v50 = _dafny.MultiSetOf(_30_v15)
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_88_v48), 0))
			_ = _index3
			var _rhs10 _dafny.Array = _33_v18
			_ = _rhs10
			var _rhs11 _dafny.Sequence = _89_v49
			_ = _rhs11
			var _rhs12 _dafny.MultiSet = (_90_v50).Intersection(_90_v50)
			_ = _rhs12
			var _lhs4 _dafny.Array = _88_v48
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen((_88_v48), 0))
			_ = _lhs5
			_33_v18 = _rhs10
			(_lhs4).ArraySet1(_rhs11, (_lhs5).Int())
			_90_v50 = _rhs12
			var _91_v51 _dafny.Array
			_ = _91_v51
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_2
			var _nw11 _dafny.Array
			_ = _nw11
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw11 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = (func(_92_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_93_i4 _dafny.Int) _dafny.Int {
						return (_93_i4).Plus(_92_p0)
					}
				})(p0)
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
			_91_v51 = _nw11
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_91_v51), 0))
			_ = _index4
			(_91_v51).ArraySet1(Companion_Default___.SafeModuloInt(_68_cf0, _68_cf0), (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_91_v51), 0))
			_ = _index5
			var _rhs13 _dafny.Int = p0
			_ = _rhs13
			var _rhs14 _dafny.Int = _dafny.IntOfUint32((_31_v16).Cardinality())
			_ = _rhs14
			var _rhs15 _dafny.Int = _68_cf0
			_ = _rhs15
			var _rhs16 _dafny.Int = ((func() _dafny.Int {
				if _30_v15 {
					return _dafny.IntOfInt64(-722)
				}
				return _68_cf0
			})()).Plus(Companion_Default___.Fm0(globalState))
			_ = _rhs16
			var _rhs17 _dafny.Int = Companion_Default___.SafeDivisionInt((_71_v31).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_71_v31).Cardinality()))).Uint32()).(_dafny.Int), (_71_v31).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_69_v29).Cardinality()), _dafny.IntOfUint32((_71_v31).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs17
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			var _lhs8 _dafny.Array = _91_v51
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_91_v51), 0))
			_ = _lhs9
			var _lhs10 *GlobalState = globalState
			_ = _lhs10
			_lhs6.F11 = _rhs13
			_68_cf0 = _rhs14
			_lhs7.F11 = _rhs15
			(_lhs8).ArraySet1(_rhs16, (_lhs9).Int())
			_lhs10.F10 = _rhs17
		} else {
			(globalState).F12 = (Companion_D3_.Create_DC8_(_dafny.SeqOf(p0), Companion_Default___.Fm1(_30_v15, !(_30_v15), _68_cf0, globalState), _30_v15, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_94_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			})), (Companion_Default___.SafeIndex(_68_cf0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_95_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality()))).Uint32(), (_69_v29).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_69_v29).Cardinality()))).Uint32()).(_dafny.CodePoint)), true)).Dtor_cf20()
			(globalState).F12 = _30_v15
			var _96_v52 _dafny.Map
			_ = _96_v52
			_96_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v15, (_68_cf0).Cmp(_dafny.IntOfInt64(491)) == 0)
			_96_v52 = (_96_v52).Update(!_dafny.Companion_Sequence_.Equal(_69_v29, _69_v29), _30_v15)
			_33_v18 = _33_v18
			var _97_v53 _dafny.Array
			_ = _97_v53
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_3
			var _nw12 _dafny.Array
			_ = _nw12
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw12 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = (func(_98_cf0 _dafny.Int, _99_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_100_i6 _dafny.Int) _dafny.Int {
						return (_100_i6).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_cf0, _99_p0)).Cardinality())
					}
				})(_68_cf0, p0)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw12 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw12).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw12).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_97_v53 = _nw12
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_97_v53), 0))
			_ = _index6
			(_97_v53).ArraySet1(p0, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_97_v53), 0))
			_ = _index7
			(_97_v53).ArraySet1(p0, (_index7).Int())
		}
		_68_cf0 = (p0).Minus(_68_cf0)
		var _101_v54 _dafny.MultiSet
		_ = _101_v54
		_101_v54 = _dafny.MultiSetOf(_30_v15, false)
		(globalState).F12 = ((_101_v54).Cardinality()).Cmp(_68_cf0) == 0
	}
	if (true) && (_30_v15) {
		(globalState).F11 = p0
		(globalState).F12 = !(_30_v15)
		var _102_v55 _dafny.CodePoint
		_ = _102_v55
		_102_v55 = _dafny.CodePoint('v')
		var _103_v56 _dafny.Map
		_ = _103_v56
		_103_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v15, _102_v55)
		var _104_v57 _dafny.Map
		_ = _104_v57
		_104_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v56, _30_v15)
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_33_v18), 0))
		_ = _index8
		(_33_v18).ArraySet1((func() bool {
			if (_104_v57).Contains(Companion_Default___.Fm17(_30_v15, globalState)) {
				return (_104_v57).Get(Companion_Default___.Fm17(_30_v15, globalState)).(bool)
			}
			return _30_v15
		})(), (_index8).Int())
		var _105_v58 _dafny.MultiSet
		_ = _105_v58
		_105_v58 = _dafny.MultiSetOf(p0)
		var _106_v59 _dafny.Sequence
		_ = _106_v59
		_106_v59 = _dafny.SeqOf(p0)
		var _107_v60 _dafny.Map
		_ = _107_v60
		_107_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v15, _30_v15)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_33_v18), 0))
		_ = _index9
		(_33_v18).ArraySet1(((func() _dafny.Int {
			if (_105_v58).Contains(p0) {
				return (_105_v58).Multiplicity(p0)
			}
			return p0
		})()).Cmp((_106_v59).Select((Companion_Default___.SafeIndex((_107_v60).Cardinality(), _dafny.IntOfUint32((_106_v59).Cardinality()))).Uint32()).(_dafny.Int)) <= 0, (_index9).Int())
		var _108_v61 _dafny.MultiSet
		_ = _108_v61
		_108_v61 = _dafny.MultiSetOf((_33_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_33_v18), 0))).Int()).(bool))
		var _109_v62 _dafny.Sequence
		_ = _109_v62
		_109_v62 = _dafny.UnicodeSeqOfUtf8Bytes("fjuprfrne")
		var _110_v63 _dafny.Set
		_ = _110_v63
		_110_v63 = _dafny.SetOf((_108_v61).Update(true, Companion_Default___.Abs(_dafny.IntOfUint32((_109_v62).Cardinality()))), _108_v61, _108_v61)
		var _111_v64 _dafny.Map
		_ = _111_v64
		_111_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_110_v63).Cardinality())
		var _112_v65 *C1
		_ = _112_v65
		var _nw13 *C1 = New_C1_()
		_ = _nw13
		_nw13.Ctor__((_111_v64).Cardinality())
		_112_v65 = _nw13
		var _113_v66 _dafny.Array
		_ = _113_v66
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_4
		var _nw14 _dafny.Array
		_ = _nw14
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw14 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.CodePoint = (func(_114_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_115_i7 _dafny.Int) _dafny.CodePoint {
					return _114_v55
				}
			})(_102_v55)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw14 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw14).ArraySet1CodePoint(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw14).ArraySet1CodePoint(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_113_v66 = _nw14
		var _116_v67 D6
		_ = _116_v67
		_116_v67 = Companion_D6_.Create_DC14_(_30_v15, _102_v55, _113_v66)
		var _pat_let_tv0 = _102_v55
		_ = _pat_let_tv0
		var _pat_let_tv1 = _30_v15
		_ = _pat_let_tv1
		var _pat_let_tv2 = _30_v15
		_ = _pat_let_tv2
		var _source1 D6 = func(_pat_let0_0 D6) D6 {
			return func(_117_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let1_0 _dafny.CodePoint) D6 {
					return func(_118_dt__update_hcf29_h0 _dafny.CodePoint) D6 {
						return func(_pat_let2_0 bool) D6 {
							return func(_119_dt__update_hcf28_h0 bool) D6 {
								return Companion_D6_.Create_DC14_(_119_dt__update_hcf28_h0, _118_dt__update_hcf29_h0, (_117_dt__update__tmp_h0).Dtor_cf30())
							}(_pat_let2_0)
						}((_pat_let_tv1) && (_pat_let_tv2))
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_116_v67)
		_ = _source1
		if _source1.Is_DC14() {
			var _120___mcc_h11 bool = _source1.Get_().(D6_DC14).Cf28
			_ = _120___mcc_h11
			var _121___mcc_h12 _dafny.CodePoint = _source1.Get_().(D6_DC14).Cf29
			_ = _121___mcc_h12
			var _122___mcc_h13 _dafny.Array = _source1.Get_().(D6_DC14).Cf30
			_ = _122___mcc_h13
			var _123_cf30 _dafny.Array = _122___mcc_h13
			_ = _123_cf30
			var _124_cf29 _dafny.CodePoint = _121___mcc_h12
			_ = _124_cf29
			var _125_cf28 bool = _120___mcc_h11
			_ = _125_cf28
			(globalState).F10 = _112_v65.F18
			_107_v60 = (_107_v60).Update(!(_30_v15), (_33_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_33_v18), 0))).Int()).(bool))
			var _126_v68 _dafny.Array
			_ = _126_v68
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
			_ = _nw15
			_126_v68 = _nw15
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_126_v68), 0))
			_ = _index10
			(_126_v68).ArraySet1(_109_v62, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_126_v68), 0))
			_ = _index11
			var _rhs18 _dafny.CodePoint = _102_v55
			_ = _rhs18
			var _rhs19 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jgxwsdkj"), (Companion_Default___.SafeIndex(_112_v65.F18, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jgxwsdkj")).Cardinality()))).Uint32(), _124_cf29)
			_ = _rhs19
			var _lhs11 _dafny.Array = _126_v68
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_126_v68), 0))
			_ = _lhs12
			_124_cf29 = _rhs18
			(_lhs11).ArraySet1(_rhs19, (_lhs12).Int())
			var _127_v69 _dafny.Array
			_ = _127_v69
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw16
			_127_v69 = _nw16
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_127_v69), 0))
			_ = _index12
			(_127_v69).ArraySet1(_112_v65.F18, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_127_v69), 0))
			_ = _index13
			(_127_v69).ArraySet1((_dafny.Zero).Minus(p0), (_index13).Int())
		} else {
			var _128___mcc_h14 _dafny.Map = _source1.Get_().(D6_DC13).Cf27
			_ = _128___mcc_h14
			var _129_cf27 _dafny.Map = _128___mcc_h14
			_ = _129_cf27
			(globalState).F15 = _102_v55
			var _130_v70 _dafny.Array
			_ = _130_v70
			var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
			_ = _nw17
			_130_v70 = _nw17
			var _131_v71 _dafny.Array
			_ = _131_v71
			var _nwElement0_2 _dafny.Array = _130_v70
			_ = _nwElement0_2
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(28))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_2, 0)
			(_nw18).ArraySet1(_130_v70, 1)
			(_nw18).ArraySet1(_130_v70, 2)
			(_nw18).ArraySet1(_130_v70, 3)
			(_nw18).ArraySet1(_130_v70, 4)
			(_nw18).ArraySet1(_130_v70, 5)
			(_nw18).ArraySet1(_130_v70, 6)
			(_nw18).ArraySet1(_130_v70, 7)
			(_nw18).ArraySet1((func() _dafny.Array {
				if _30_v15 {
					return _130_v70
				}
				return _130_v70
			})(), 8)
			(_nw18).ArraySet1(_130_v70, 9)
			(_nw18).ArraySet1(_130_v70, 10)
			(_nw18).ArraySet1(_130_v70, 11)
			(_nw18).ArraySet1(_130_v70, 12)
			(_nw18).ArraySet1(_130_v70, 13)
			(_nw18).ArraySet1(_130_v70, 14)
			(_nw18).ArraySet1(_130_v70, 15)
			(_nw18).ArraySet1(_130_v70, 16)
			(_nw18).ArraySet1(_130_v70, 17)
			(_nw18).ArraySet1(_130_v70, 18)
			(_nw18).ArraySet1(_130_v70, 19)
			(_nw18).ArraySet1(_130_v70, 20)
			(_nw18).ArraySet1(_130_v70, 21)
			(_nw18).ArraySet1(_130_v70, 22)
			(_nw18).ArraySet1(_130_v70, 23)
			(_nw18).ArraySet1(_130_v70, 24)
			(_nw18).ArraySet1(_130_v70, 25)
			(_nw18).ArraySet1(_130_v70, 26)
			(_nw18).ArraySet1(_130_v70, 27)
			_131_v71 = _nw18
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_131_v71), 0))
			_ = _index14
			(_131_v71).ArraySet1(_130_v70, (_index14).Int())
			var _132_v72 *C0
			_ = _132_v72
			var _nw19 *C0 = New_C0_()
			_ = _nw19
			_nw19.Ctor__()
			_132_v72 = _nw19
			var _133_v73 _dafny.Array
			_ = _133_v73
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw20
			_133_v73 = _nw20
			var _134_v74 _dafny.Array
			_ = _134_v74
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_5
			var _nw21 _dafny.Array
			_ = _nw21
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw21 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Map = (func(_135_v14 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_136_i8 _dafny.Int) _dafny.Map {
						return _135_v14
					}
				})(_29_v14)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw21 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw21).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw21).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_134_v74 = _nw21
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_133_v73), 0))
			_ = _index15
			(_133_v73).ArraySet1(_134_v74, (_index15).Int())
			var _137_v75 _dafny.Set
			_ = _137_v75
			_137_v75 = _dafny.SetOf(_130_v70, _130_v70)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_131_v71), 0))
			_ = _index16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_133_v73), 0))
			_ = _index17
			var _rhs20 _dafny.Array = _130_v70
			_ = _rhs20
			var _rhs21 *C0 = _132_v72
			_ = _rhs21
			var _rhs22 _dafny.Array = (func() _dafny.Array {
				if (_137_v75).IsProperSubsetOf(_137_v75) {
					return _134_v74
				}
				return _134_v74
			})()
			_ = _rhs22
			var _lhs13 _dafny.Array = _131_v71
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_131_v71), 0))
			_ = _lhs14
			var _lhs15 _dafny.Array = _133_v73
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_133_v73), 0))
			_ = _lhs16
			(_lhs13).ArraySet1(_rhs20, (_lhs14).Int())
			_132_v72 = _rhs21
			(_lhs15).ArraySet1(_rhs22, (_lhs16).Int())
			var _138_v76 _dafny.Array
			_ = _138_v76
			var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
			_ = _nw22
			_138_v76 = _nw22
			var _139_v77 _dafny.Map
			_ = _139_v77
			_139_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v62, _112_v65.F18)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_138_v76), 0))
			_ = _index18
			(_138_v76).ArraySet1(_139_v77, (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_138_v76), 0))
			_ = _index19
			(_138_v76).ArraySet1(_139_v77, (_index19).Int())
			var _140_v78 _dafny.MultiSet
			_ = _140_v78
			_140_v78 = _dafny.MultiSetOf(_109_v62)
			var _141_v79 _dafny.Map
			_ = _141_v79
			_141_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_33_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_33_v18), 0))).Int()).(bool), p0)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_130_v70), 0))
			_ = _index20
			(_130_v70).ArraySet1((func() _dafny.Int {
				if (_140_v78).Contains((_132_v72).Fm7(_30_v15, _dafny.IntOfUint32((_dafny.SeqOf(_102_v55, _102_v55)).Cardinality()), _141_v79, _109_v62, globalState)) {
					return (_140_v78).Multiplicity((_132_v72).Fm7(_30_v15, _dafny.IntOfUint32((_dafny.SeqOf(_102_v55, _102_v55)).Cardinality()), _141_v79, _109_v62, globalState))
				}
				return (_dafny.Zero).Minus((_108_v61).Cardinality())
			})(), (_index20).Int())
			var _142_v81 _dafny.Map
			_ = _142_v81
			_142_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
				var _coll9 = _dafny.NewBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-954), _dafny.IntOfInt64(990))); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _143_v80 _dafny.Int
					_143_v80 = interface{}(_compr_9).(_dafny.Int)
					if ((_dafny.IntOfInt64(-954)).Cmp(_143_v80) <= 0) && ((_143_v80).Cmp(_dafny.IntOfInt64(990)) < 0) {
						_coll9.Add((_143_v80).Times(_112_v65.F18))
					}
				}
				return _coll9.ToSet()
			}(), p0)
			var _144_v83 D6
			_ = _144_v83
			_144_v83 = Companion_D6_.Create_DC13_(_129_cf27)
			var _145_v84 _dafny.MultiSet
			_ = _145_v84
			_145_v84 = _dafny.MultiSetOf(_144_v83)
			var _146_v85 _dafny.Map
			_ = _146_v85
			_146_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v84, (_33_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_33_v18), 0))).Int()).(bool))
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_33_v18), 0))
			_ = _index21
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_130_v70), 0))
			_ = _index22
			var _rhs23 _dafny.Sequence = _106_v59
			_ = _rhs23
			var _rhs24 bool = !(_30_v15)
			_ = _rhs24
			var _rhs25 _dafny.Int = (func() _dafny.Int {
				if (_142_v81).Contains(func() _dafny.Set {
					var _coll10 = _dafny.NewBuilder()
					_ = _coll10
					for _iter10 := _dafny.Iterate((_29_v14).Keys().Elements()); ; {
						_compr_10, _ok10 := _iter10()
						if !_ok10 {
							break
						}
						var _147_v82 _dafny.Int
						_147_v82 = interface{}(_compr_10).(_dafny.Int)
						if (_29_v14).Contains(_147_v82) {
							_coll10.Add(Companion_Default___.SafeModuloInt(_147_v82, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(420))).Cardinality()))
						}
					}
					return _coll10.ToSet()
				}()) {
					return (_142_v81).Get(func() _dafny.Set {
						var _coll11 = _dafny.NewBuilder()
						_ = _coll11
						for _iter11 := _dafny.Iterate((_29_v14).Keys().Elements()); ; {
							_compr_11, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _148_v82 _dafny.Int
							_148_v82 = interface{}(_compr_11).(_dafny.Int)
							if (_29_v14).Contains(_148_v82) {
								_coll11.Add(Companion_Default___.SafeModuloInt(_148_v82, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(420))).Cardinality()))
							}
						}
						return _coll11.ToSet()
					}()).(_dafny.Int)
				}
				return p0
			})()
			_ = _rhs25
			var _rhs26 _dafny.Int = (_112_v65.F18).Minus(Companion_Default___.SafeModuloInt((_146_v85).Cardinality(), (_dafny.SetOf(_30_v15)).Cardinality()))
			_ = _rhs26
			var _rhs27 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_109_v62, _109_v62)
			_ = _rhs27
			var _lhs17 _dafny.Array = _33_v18
			_ = _lhs17
			var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_33_v18), 0))
			_ = _lhs18
			var _lhs19 *GlobalState = globalState
			_ = _lhs19
			var _lhs20 _dafny.Array = _130_v70
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_130_v70), 0))
			_ = _lhs21
			_106_v59 = _rhs23
			(_lhs17).ArraySet1(_rhs24, (_lhs18).Int())
			_lhs19.F10 = _rhs25
			(_lhs20).ArraySet1(_rhs26, (_lhs21).Int())
			_109_v62 = _rhs27
		}
	} else {
		var _149_v86 _dafny.MultiSet
		_ = _149_v86
		_149_v86 = _dafny.MultiSetOf(!(_30_v15))
		var _150_v87 _dafny.Map
		_ = _150_v87
		_150_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_v18, _33_v18)
		var _rhs28 bool = _30_v15
		_ = _rhs28
		var _rhs29 _dafny.MultiSet = _149_v86
		_ = _rhs29
		var _rhs30 _dafny.Int = ((_32_v17).Union(_32_v17)).Cardinality()
		_ = _rhs30
		var _rhs31 _dafny.Map = (func() _dafny.Map {
			if (p0).Cmp(p0) > 0 {
				return (_150_v87).Merge(_150_v87)
			}
			return _150_v87
		})()
		_ = _rhs31
		var _lhs22 *GlobalState = globalState
		_ = _lhs22
		var _lhs23 *GlobalState = globalState
		_ = _lhs23
		_lhs22.F12 = _rhs28
		_149_v86 = _rhs29
		_lhs23.F11 = _rhs30
		_150_v87 = _rhs31
		var _151_v88 _dafny.Set
		_ = _151_v88
		_151_v88 = _dafny.SetOf(true)
		(globalState).F10 = (p0).Plus((_151_v88).Cardinality())
		var _152_v89 _dafny.Map
		_ = _152_v89
		_152_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v15, _30_v15)
		var _153_v90 _dafny.MultiSet
		_ = _153_v90
		_153_v90 = _dafny.MultiSetOf(_152_v89, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_30_v15, _30_v15, p0, globalState), _30_v15))
		_153_v90 = (func() _dafny.MultiSet {
			if (_31_v16).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_31_v16).Cardinality()))).Uint32()).(bool) {
				return (_153_v90).Update(_152_v89, Companion_Default___.Abs(p0))
			}
			return (_153_v90).Union(_153_v90)
		})()
		var _154_v91 _dafny.Array
		_ = _154_v91
		var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
		_ = _nw23
		_154_v91 = _nw23
		var _155_v92 _dafny.Array
		_ = _155_v92
		var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
		_ = _nw24
		_155_v92 = _nw24
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_154_v91), 0))
		_ = _index23
		(_154_v91).ArraySet1(_155_v92, (_index23).Int())
		var _156_v93 _dafny.Sequence
		_ = _156_v93
		_156_v93 = _dafny.UnicodeSeqOfUtf8Bytes("hpyho")
		var _157_v94 _dafny.Sequence
		_ = _157_v94
		_157_v94 = _156_v93
		var _158_v95 _dafny.CodePoint
		_ = _158_v95
		_158_v95 = _dafny.CodePoint('c')
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_154_v91), 0))
		_ = _index24
		var _nwElement0_3 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lksj"), _156_v93), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lksj"), _156_v93)).Cardinality()))).Uint32(), (_156_v93).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_156_v93).Cardinality()))).Uint32()).(_dafny.CodePoint))
		_ = _nwElement0_3
		var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(5))
		_ = _nw25
		(_nw25).ArraySet1(_nwElement0_3, 0)
		(_nw25).ArraySet1((_157_v94), 1)
		(_nw25).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_156_v93, _156_v93), 2)
		(_nw25).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(116))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_159_i9 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		})), 3)
		(_nw25).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_156_v93, _156_v93), (Companion_Default___.SafeIndex(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_156_v93, _156_v93)).Cardinality()))).Uint32(), _158_v95), 4)
		(_154_v91).ArraySet1(_nw25, (_index24).Int())
		var _rhs32 bool = _30_v15
		_ = _rhs32
		var _rhs33 _dafny.CodePoint = _158_v95
		_ = _rhs33
		var _rhs34 _dafny.Int = p0
		_ = _rhs34
		var _rhs35 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8(globalState), _dafny.Companion_Sequence_.Update(_156_v93, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_156_v93).Cardinality()))).Uint32(), _158_v95))
		_ = _rhs35
		var _lhs24 *GlobalState = globalState
		_ = _lhs24
		var _lhs25 *GlobalState = globalState
		_ = _lhs25
		var _lhs26 *GlobalState = globalState
		_ = _lhs26
		_lhs24.F12 = _rhs32
		_158_v95 = _rhs33
		_lhs25.F10 = _rhs34
		_lhs26.F3 = _rhs35
	}
	var _160_i10 _dafny.Int
	_ = _160_i10
	_160_i10 = _dafny.Zero
	{
		for !(_30_v15) {
			{
				if (_160_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_160_i10 = (_160_i10).Plus(_dafny.One)
				var _161_v96 *C0
				_ = _161_v96
				var _nw26 *C0 = New_C0_()
				_ = _nw26
				_nw26.Ctor__()
				_161_v96 = _nw26
				var _162_v97 _dafny.Map
				_ = _162_v97
				_162_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v96, _30_v15)
				var _163_v98 _dafny.Map
				_ = _163_v98
				_163_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v15, _162_v97)
				_163_v98 = (_163_v98).Update(_30_v15, (_162_v97).Merge(_162_v97))
				(globalState).F11 = ((Companion_Default___.Fm15(p0, _30_v15, _dafny.IntOfInt64(794), globalState)).Cardinality()).Minus(p0)
				var _164_v99 _dafny.Sequence
				_ = _164_v99
				_164_v99 = _dafny.SeqOf(p0, p0)
				var _165_v100 *C1
				_ = _165_v100
				var _nw27 *C1 = New_C1_()
				_ = _nw27
				_nw27.Ctor__(_dafny.IntOfInt64(451))
				_165_v100 = _nw27
				var _166_v102 _dafny.CodePoint
				_ = _166_v102
				_166_v102 = _dafny.CodePoint('u')
				var _167_v103 _dafny.Sequence
				_ = _167_v103
				_167_v103 = _dafny.SeqOf(_166_v102, _166_v102)
				var _168_v104 _dafny.Map
				_ = _168_v104
				_168_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v100, func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter12 := _dafny.Iterate((_167_v103).Elements()); ; {
						_compr_12, _ok12 := _iter12()
						if !_ok12 {
							break
						}
						var _169_v101 _dafny.CodePoint
						_169_v101 = interface{}(_compr_12).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_167_v103, _169_v101) {
							_coll12.Add(_169_v101, _165_v100.F18)
						}
					}
					return _coll12.ToMap()
				}())
				var _170_v105 _dafny.Array
				_ = _170_v105
				var _nwElement0_4 _dafny.Int = (func() _dafny.Int {
					if _30_v15 {
						return p0
					}
					return _dafny.IntOfInt64(899)
				})()
				_ = _nwElement0_4
				var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(23))
				_ = _nw28
				(_nw28).ArraySet1(_nwElement0_4, 0)
				(_nw28).ArraySet1((p0).Times(p0), 1)
				(_nw28).ArraySet1(_dafny.IntOfInt64(586), 2)
				(_nw28).ArraySet1(_dafny.IntOfInt64(120), 3)
				(_nw28).ArraySet1(p0, 4)
				(_nw28).ArraySet1(p0, 5)
				(_nw28).ArraySet1(p0, 6)
				(_nw28).ArraySet1(p0, 7)
				(_nw28).ArraySet1(p0, 8)
				(_nw28).ArraySet1(p0, 9)
				(_nw28).ArraySet1(Companion_Default___.Fm0(globalState), 10)
				(_nw28).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_164_v99, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_164_v99).Cardinality()))).Uint32(), (_dafny.Zero).Minus(p0))).Cardinality()), 11)
				(_nw28).ArraySet1(p0, 12)
				(_nw28).ArraySet1(p0, 13)
				(_nw28).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_164_v99, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_164_v99).Cardinality()))).Uint32(), _dafny.IntOfInt64(472))).Cardinality())).Times((_168_v104).Cardinality()), 14)
				(_nw28).ArraySet1(_165_v100.F18, 15)
				(_nw28).ArraySet1(_165_v100.F18, 16)
				(_nw28).ArraySet1(p0, 17)
				(_nw28).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_167_v103, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_167_v103).Cardinality()))).Uint32(), (_167_v103).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_167_v103).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()), 18)
				(_nw28).ArraySet1(_165_v100.F18, 19)
				(_nw28).ArraySet1((func() _dafny.Int {
					if Companion_Default___.Fm1(_30_v15, _30_v15, _dafny.IntOfUint32((_167_v103).Cardinality()), globalState) {
						return Companion_Default___.Fm0(globalState)
					}
					return (_164_v99).Select((Companion_Default___.SafeIndex(_165_v100.F18, _dafny.IntOfUint32((_164_v99).Cardinality()))).Uint32()).(_dafny.Int)
				})(), 20)
				(_nw28).ArraySet1(_165_v100.F18, 21)
				(_nw28).ArraySet1(p0, 22)
				_170_v105 = _nw28
				r0 = _170_v105
				(globalState).F12 = !(_30_v15)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	_33_v18 = _33_v18
	var _171_v106 _dafny.Array
	_ = _171_v106
	var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
	_ = _nw29
	_171_v106 = _nw29
	_171_v106 = _171_v106
	var _172_v107 _dafny.Array
	_ = _172_v107
	var _nwElement0_5 _dafny.Int = p0
	_ = _nwElement0_5
	var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.One)
	_ = _nw30
	(_nw30).ArraySet1(_nwElement0_5, 0)
	_172_v107 = _nw30
	r0 = _172_v107
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _173_v0 _dafny.Sequence
	_ = _173_v0
	_173_v0 = _dafny.UnicodeSeqOfUtf8Bytes("yo")
	var _174_v2 _dafny.CodePoint
	_ = _174_v2
	_174_v2 = _dafny.CodePoint('k')
	var _175_v3 _dafny.Set
	_ = _175_v3
	_175_v3 = _dafny.SetOf(_174_v2)
	var _176_v4 _dafny.MultiSet
	_ = _176_v4
	_176_v4 = _dafny.MultiSetOf(func() _dafny.Set {
		var _coll13 = _dafny.NewBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('n'))).Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _177_v1 _dafny.CodePoint
			_177_v1 = interface{}(_compr_13).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('n')), _177_v1) {
				_coll13.Add(_177_v1)
			}
		}
		return _coll13.ToSet()
	}(), _175_v3)
	var _178_globalState *GlobalState
	_ = _178_globalState
	var _nw31 *GlobalState = New_GlobalState_()
	_ = _nw31
	_nw31.Ctor__(_dafny.IntOfInt64(134), true, _dafny.IntOfInt64(107), _173_v0, _dafny.UnicodeSeqOfUtf8Bytes("ci"), _dafny.IntOfInt64(380), _dafny.IntOfInt64(-147), false, false, _dafny.IntOfInt64(727), _dafny.IntOfInt64(540), _dafny.IntOfInt64(576), false, false, _176_v4, _dafny.CodePoint('c'), _dafny.IntOfInt64(403), _dafny.IntOfInt64(-68))
	_178_globalState = _nw31
	var _179_v5 _dafny.Int
	_ = _179_v5
	_179_v5 = _dafny.IntOfInt64(980)
	var _hi1 _dafny.Int = _179_v5
	_ = _hi1
	for _180_i0 := _dafny.IntOfInt64(-735); _180_i0.Cmp(_hi1) < 0; _180_i0 = _180_i0.Plus(_dafny.One) {
		var _181_v6 _dafny.Array
		_ = _181_v6
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw32
		_181_v6 = _nw32
		var _182_v7 _dafny.Map
		_ = _182_v7
		_182_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_181_v6, _181_v6, _181_v6), false)
		var _183_v8 _dafny.Sequence
		_ = _183_v8
		_183_v8 = _dafny.SeqOf(_181_v6, _181_v6, _181_v6, _181_v6, _181_v6)
		_182_v7 = (_182_v7).Update(_dafny.Companion_Sequence_.Concatenate(_183_v8, _183_v8), false)
		var _184_v9 _dafny.Array
		_ = _184_v9
		var _nwElement0_6 _dafny.Int = _dafny.IntOfInt64(-13)
		_ = _nwElement0_6
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(2))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_6, 0)
		(_nw33).ArraySet1((Companion_D0_.Create_DC0_(_180_i0)).Dtor_cf0(), 1)
		_184_v9 = _nw33
		var _185_v10 bool
		_ = _185_v10
		_185_v10 = true
		var _186_v11 _dafny.Sequence
		_ = _186_v11
		_186_v11 = _dafny.SeqOf(!(_185_v10), _185_v10)
		var _187_v12 _dafny.Map
		_ = _187_v12
		_187_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_179_v5).Times(_180_i0), _185_v10)
		var _rhs36 _dafny.Array = (func() _dafny.Array {
			if (_dafny.IntOfInt64(-438)).Cmp(_dafny.IntOfInt64(880)) > 0 {
				return _184_v9
			}
			return _184_v9
		})()
		_ = _rhs36
		var _rhs37 bool = !(!_dafny.Companion_Sequence_.Equal(_173_v0, _173_v0))
		_ = _rhs37
		var _rhs38 bool = (func() bool {
			if _185_v10 {
				return _185_v10
			}
			return (_186_v11).Select((Companion_Default___.SafeIndex(_179_v5, _dafny.IntOfUint32((_186_v11).Cardinality()))).Uint32()).(bool)
		})()
		_ = _rhs38
		var _rhs39 bool = (func() bool {
			if (_187_v12).Contains(_dafny.IntOfInt64(716)) {
				return (_187_v12).Get(_dafny.IntOfInt64(716)).(bool)
			}
			return _185_v10
		})()
		_ = _rhs39
		var _rhs40 _dafny.Int = ((_dafny.Zero).Minus(_180_i0)).Minus((_dafny.IntOfUint32((_173_v0).Cardinality())).Times((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(_178_globalState)))))
		_ = _rhs40
		var _lhs27 *GlobalState = _178_globalState
		_ = _lhs27
		var _lhs28 *GlobalState = _178_globalState
		_ = _lhs28
		var _lhs29 *GlobalState = _178_globalState
		_ = _lhs29
		_184_v9 = _rhs36
		_lhs27.F12 = _rhs37
		_lhs28.F12 = _rhs38
		_lhs29.F12 = _rhs39
		_179_v5 = _rhs40
		var _188_v13 _dafny.MultiSet
		_ = _188_v13
		_188_v13 = _dafny.MultiSetOf(Companion_Default___.Fm0(_178_globalState))
		_188_v13 = (_188_v13).Union(_188_v13)
		var _189_v14 _dafny.MultiSet
		_ = _189_v14
		_189_v14 = _dafny.MultiSetOf(_185_v10, true, true)
		var _190_v15 _dafny.Map
		_ = _190_v15
		_190_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v10, (_189_v14).Cardinality())
		var _191_v16 _dafny.Map
		_ = _191_v16
		_191_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v5, _190_v15)
		_191_v16 = (_191_v16).Update(_180_i0, _190_v15)
	}
	var _192_v17 bool
	_ = _192_v17
	_192_v17 = false
	var _193_v18 _dafny.Map
	_ = _193_v18
	_193_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v17, _179_v5)
	var _194_v19 _dafny.Map
	_ = _194_v19
	_194_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_179_v5, _179_v5, (_193_v18).Cardinality())).Cardinality(), !(_192_v17))
	if (func() bool {
		if (_194_v19).Contains(_179_v5) {
			return (_194_v19).Get(_179_v5).(bool)
		}
		return (_179_v5).Cmp(_179_v5) == 0
	})() {
		var _195_v20 _dafny.Array
		_ = _195_v20
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
		_ = _nw34
		_195_v20 = _nw34
		var _196_v21 D0
		_ = _196_v21
		_196_v21 = Companion_D0_.Create_DC3_(_179_v5, _192_v17, _179_v5, _dafny.IntOfUint32((_173_v0).Cardinality()))
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))
		_ = _index25
		(_195_v20).ArraySet1((func() bool {
			if (func() bool {
				if (_194_v19).Contains(_179_v5) {
					return (_194_v19).Get(_179_v5).(bool)
				}
				return (_196_v21).Dtor_cf8()
			})() {
				return Companion_Default___.Fm1(_192_v17, _192_v17, _179_v5, _178_globalState)
			}
			return _192_v17
		})(), (_index25).Int())
		var _197_v22 _dafny.Set
		_ = _197_v22
		_197_v22 = _dafny.SetOf(_192_v17, !(_192_v17))
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))
		_ = _index26
		(_195_v20).ArraySet1(Companion_Default___.Fm1(false, (_197_v22).IsDisjointFrom(_197_v22), Companion_Default___.SafeModuloInt(_179_v5, _179_v5), _178_globalState), (_index26).Int())
		var _198_v23 _dafny.Array
		_ = _198_v23
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
		_ = _nw35
		_198_v23 = _nw35
		var _199_v25 _dafny.Map
		_ = _199_v25
		_199_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v5, _179_v5)
		var _200_v26 _dafny.Map
		_ = _200_v26
		_200_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate((_199_v25).Keys().Elements()); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _201_v24 _dafny.Int
				_201_v24 = interface{}(_compr_14).(_dafny.Int)
				if (_199_v25).Contains(_201_v24) {
					_coll14.Add((_201_v24).Times(_179_v5), Companion_D0_.Create_DC2_((_195_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))).Int()).(bool)))
				}
			}
			return _coll14.ToMap()
		}(), _179_v5)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_198_v23), 0))
		_ = _index27
		(_198_v23).ArraySet1(_200_v26, (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_198_v23), 0))
		_ = _index28
		(_198_v23).ArraySet1(Companion_Default___.Fm2(_178_globalState), (_index28).Int())
		var _source2 D0 = Companion_Default___.Fm3(_dafny.IntOfUint32((_173_v0).Cardinality()), _178_globalState)
		_ = _source2
		if _source2.Is_DC1() {
			var _202___mcc_h0 _dafny.Int = _source2.Get_().(D0_DC1).Cf1
			_ = _202___mcc_h0
			var _203___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC1).Cf2
			_ = _203___mcc_h1
			var _204___mcc_h2 _dafny.Int = _source2.Get_().(D0_DC1).Cf3
			_ = _204___mcc_h2
			var _205___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC1).Cf4
			_ = _205___mcc_h3
			var _206___mcc_h4 _dafny.Array = _source2.Get_().(D0_DC1).Cf5
			_ = _206___mcc_h4
			var _207_cf5 _dafny.Array = _206___mcc_h4
			_ = _207_cf5
			var _208_cf4 _dafny.Int = _205___mcc_h3
			_ = _208_cf4
			var _209_cf3 _dafny.Int = _204___mcc_h2
			_ = _209_cf3
			var _210_cf2 _dafny.Int = _203___mcc_h1
			_ = _210_cf2
			var _211_cf1 _dafny.Int = _202___mcc_h0
			_ = _211_cf1
			var _212_v27 _dafny.Array
			_ = _212_v27
			var _out0 _dafny.Array
			_ = _out0
			_out0 = Companion_Default___.M0((func() _dafny.Int {
				if (_195_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))).Int()).(bool) {
					return _210_cf2
				}
				return _208_cf4
			})(), _178_globalState)
			_212_v27 = _out0
			(_178_globalState).F10 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_208_cf4))
			var _213_v28 _dafny.Array
			_ = _213_v28
			var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
			_ = _nw36
			_213_v28 = _nw36
			var _214_v29 _dafny.Sequence
			_ = _214_v29
			_214_v29 = _dafny.SeqOf(_210_cf2)
			var _215_v30 _dafny.Set
			_ = _215_v30
			_215_v30 = _dafny.SetOf(_214_v29)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_213_v28), 0))
			_ = _index29
			(_213_v28).ArraySet1(_215_v30, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_213_v28), 0))
			_ = _index30
			(_213_v28).ArraySet1(_215_v30, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))
			_ = _index31
			(_195_v20).ArraySet1(!(_dafny.Companion_Sequence_.Contains(_173_v0, _174_v2)), (_index31).Int())
		} else if _source2.Is_DC2() {
			var _216___mcc_h5 bool = _source2.Get_().(D0_DC2).Cf6
			_ = _216___mcc_h5
			var _217_cf6 bool = _216___mcc_h5
			_ = _217_cf6
			var _218_v31 _dafny.Sequence
			_ = _218_v31
			_218_v31 = _dafny.SeqOf((_195_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))).Int()).(bool))
			_218_v31 = _dafny.Companion_Sequence_.Concatenate(_218_v31, _218_v31)
			var _rhs41 _dafny.Int = Companion_Default___.SafeModuloInt(_179_v5, _179_v5)
			_ = _rhs41
			var _rhs42 _dafny.Int = ((_dafny.Zero).Minus(_179_v5)).Minus(_179_v5)
			_ = _rhs42
			var _lhs30 *GlobalState = _178_globalState
			_ = _lhs30
			var _lhs31 *GlobalState = _178_globalState
			_ = _lhs31
			_lhs30.F10 = _rhs41
			_lhs31.F10 = _rhs42
			(_178_globalState).F12 = false
			var _219_v32 _dafny.Array
			_ = _219_v32
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_6
			var _nw37 _dafny.Array
			_ = _nw37
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw37 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = (func(_220_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_221_i1 _dafny.Int) _dafny.Int {
						return (_221_i1).Plus(_220_v5)
					}
				})(_179_v5)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw37 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw37).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw37).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_219_v32 = _nw37
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_219_v32), 0))
			_ = _index32
			(_219_v32).ArraySet1(_179_v5, (_index32).Int())
			var _222_v33 _dafny.Map
			_ = _222_v33
			_222_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_217_cf6), (_195_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))).Int()).(bool))
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_219_v32), 0))
			_ = _index33
			var _rhs43 _dafny.Int = (func() _dafny.Int {
				if _217_cf6 {
					return _179_v5
				}
				return Companion_Default___.SafeModuloInt(_179_v5, (_222_v33).Cardinality())
			})()
			_ = _rhs43
			var _rhs44 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_173_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(401))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_223_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_224_i2 _dafny.Int) _dafny.CodePoint {
					return _223_v2
				}
			})(_174_v2))))
			_ = _rhs44
			var _lhs32 _dafny.Array = _219_v32
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_219_v32), 0))
			_ = _lhs33
			(_lhs32).ArraySet1(_rhs43, (_lhs33).Int())
			_173_v0 = _rhs44
		} else if _source2.Is_DC3() {
			var _225___mcc_h6 _dafny.Int = _source2.Get_().(D0_DC3).Cf7
			_ = _225___mcc_h6
			var _226___mcc_h7 bool = _source2.Get_().(D0_DC3).Cf8
			_ = _226___mcc_h7
			var _227___mcc_h8 _dafny.Int = _source2.Get_().(D0_DC3).Cf9
			_ = _227___mcc_h8
			var _228___mcc_h9 _dafny.Int = _source2.Get_().(D0_DC3).Cf10
			_ = _228___mcc_h9
			var _229_cf10 _dafny.Int = _228___mcc_h9
			_ = _229_cf10
			var _230_cf9 _dafny.Int = _227___mcc_h8
			_ = _230_cf9
			var _231_cf8 bool = _226___mcc_h7
			_ = _231_cf8
			var _232_cf7 _dafny.Int = _225___mcc_h6
			_ = _232_cf7
			var _233_v34 _dafny.Sequence
			_ = _233_v34
			_233_v34 = _dafny.SeqOf(_194_v19, (_194_v19).Update(_dafny.IntOfInt64(959), _231_cf8), _194_v19)
			var _234_v35 _dafny.MultiSet
			_ = _234_v35
			_234_v35 = _dafny.MultiSetOf(_174_v2)
			var _235_v36 _dafny.Sequence
			_ = _235_v36
			_235_v36 = _dafny.SeqOf(_231_cf8)
			var _236_v37 _dafny.MultiSet
			_ = _236_v37
			_236_v37 = _dafny.MultiSetOf(_179_v5, _229_cf10)
			var _237_v38 _dafny.Array
			_ = _237_v38
			var _nwElement0_7 _dafny.Int = Companion_Default___.SafeModuloInt(_179_v5, _232_cf7)
			_ = _nwElement0_7
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(29))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_7, 0)
			(_nw38).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if _231_cf8 {
					return _179_v5
				}
				return _229_cf10
			})()), 1)
			(_nw38).ArraySet1(_dafny.IntOfInt64(11), 2)
			(_nw38).ArraySet1((_232_cf7).Plus(_230_cf9), 3)
			(_nw38).ArraySet1(_230_cf9, 4)
			(_nw38).ArraySet1((Companion_Default___.Fm0(_178_globalState)).Times(_dafny.IntOfUint32((_233_v34).Cardinality())), 5)
			(_nw38).ArraySet1(_179_v5, 6)
			(_nw38).ArraySet1(_179_v5, 7)
			(_nw38).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(693))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_238_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_239_i3 _dafny.Int) _dafny.Int {
					return _238_cf9
				}
			})(_230_cf9)))).Cardinality()), 8)
			(_nw38).ArraySet1((_230_cf9).Times(_dafny.IntOfInt64(419)), 9)
			(_nw38).ArraySet1(_dafny.IntOfInt64(437), 10)
			(_nw38).ArraySet1((((_234_v35).Update(_174_v2, Companion_Default___.Abs(_179_v5))).Cardinality()).Minus(_232_cf7), 11)
			(_nw38).ArraySet1(_232_cf7, 12)
			(_nw38).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm4(_178_globalState), _235_v36)).Cardinality()), 13)
			(_nw38).ArraySet1(_232_cf7, 14)
			(_nw38).ArraySet1(_dafny.IntOfInt64(-540), 15)
			(_nw38).ArraySet1(_229_cf10, 16)
			(_nw38).ArraySet1((func() _dafny.Int {
				if (_236_v37).Contains(_229_cf10) {
					return (_236_v37).Multiplicity(_229_cf10)
				}
				return _229_cf10
			})(), 17)
			(_nw38).ArraySet1(_179_v5, 18)
			(_nw38).ArraySet1(_229_cf10, 19)
			(_nw38).ArraySet1(_179_v5, 20)
			(_nw38).ArraySet1(_179_v5, 21)
			(_nw38).ArraySet1((_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32((_173_v0).Cardinality()))).Minus(_dafny.IntOfUint32((Companion_Default___.Fm5(_229_cf10, _178_globalState)).Cardinality()))), 22)
			(_nw38).ArraySet1(_229_cf10, 23)
			(_nw38).ArraySet1((_179_v5).Times(_dafny.IntOfInt64(-530)), 24)
			(_nw38).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_179_v5, (_dafny.Zero).Minus(_dafny.IntOfUint32((_173_v0).Cardinality())))).Cardinality()), 25)
			(_nw38).ArraySet1(_229_cf10, 26)
			(_nw38).ArraySet1(_179_v5, 27)
			(_nw38).ArraySet1(_229_cf10, 28)
			_237_v38 = _nw38
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_237_v38), 0))
			_ = _index34
			(_237_v38).ArraySet1((_230_cf9).Plus(_232_cf7), (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_237_v38), 0))
			_ = _index35
			(_237_v38).ArraySet1(_dafny.IntOfUint32((_173_v0).Cardinality()), (_index35).Int())
			var _240_v39 D1
			_ = _240_v39
			_240_v39 = Companion_D1_.Create_DC4_(_174_v2)
			var _pat_let_tv3 = _174_v2
			_ = _pat_let_tv3
			(_178_globalState).F15 = (func(_pat_let3_0 D1) D1 {
				return func(_241_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let4_0 _dafny.CodePoint) D1 {
						return func(_242_dt__update_hcf11_h0 _dafny.CodePoint) D1 {
							return Companion_D1_.Create_DC4_(_242_dt__update_hcf11_h0)
						}(_pat_let4_0)
					}(_pat_let_tv3)
				}(_pat_let3_0)
			}(_240_v39)).Dtor_cf11()
			var _243_v40 _dafny.Array
			_ = _243_v40
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(25))
			_ = _nw39
			_243_v40 = _nw39
			var _244_v41 _dafny.Set
			_ = _244_v41
			_244_v41 = _dafny.SetOf(_243_v40)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_237_v38), 0))
			_ = _index36
			(_237_v38).ArraySet1(((_244_v41).Union(_244_v41)).Cardinality(), (_index36).Int())
			_231_cf8 = ((_195_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))).Int()).(bool)) && (_192_v17)
		} else {
			var _245___mcc_h10 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
			_ = _245___mcc_h10
			var _246_cf0 _dafny.Int = _245___mcc_h10
			_ = _246_cf0
			(_178_globalState).F12 = !(_192_v17) || ((_195_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))).Int()).(bool))
			(_178_globalState).F3 = _173_v0
			(_178_globalState).F15 = _174_v2
			var _247_v42 _dafny.Array
			_ = _247_v42
			var _out1 _dafny.Array
			_ = _out1
			_out1 = Companion_Default___.M0(Companion_Default___.Fm0(_178_globalState), _178_globalState)
			_247_v42 = _out1
		}
		var _248_v43 _dafny.MultiSet
		_ = _248_v43
		_248_v43 = _dafny.MultiSetOf(_dafny.IntOfUint32((_173_v0).Cardinality()))
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_195_v20), 0))
		_ = _index37
		(_195_v20).ArraySet1(!((_248_v43).IsSubsetOf((_248_v43).Union(_248_v43))), (_index37).Int())
		var _249_v44 _dafny.Map
		_ = _249_v44
		_249_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v17, !(_192_v17))
		var _250_v45 _dafny.Array
		_ = _250_v45
		var _out2 _dafny.Array
		_ = _out2
		_out2 = Companion_Default___.M0((func() _dafny.Int {
			if (_248_v43).Contains((_249_v44).Cardinality()) {
				return (_248_v43).Multiplicity((_249_v44).Cardinality())
			}
			return _179_v5
		})(), _178_globalState)
		_250_v45 = _out2
	} else {
		var _251_v46 _dafny.Array
		_ = _251_v46
		var _out3 _dafny.Array
		_ = _out3
		_out3 = Companion_Default___.M0(_179_v5, _178_globalState)
		_251_v46 = _out3
		var _252_v47 D0
		_ = _252_v47
		_252_v47 = Companion_D0_.Create_DC3_(_179_v5, _192_v17, _dafny.IntOfInt64(-428), _179_v5)
		if (_252_v47).Dtor_cf8() {
			var _253_v48 _dafny.Sequence
			_ = _253_v48
			_253_v48 = _173_v0
			(_178_globalState).F3 = (_253_v48)
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_251_v46), 0))
			_ = _index38
			(_251_v46).ArraySet1(_179_v5, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_251_v46), 0))
			_ = _index39
			var _rhs45 _dafny.Int = _179_v5
			_ = _rhs45
			var _rhs46 bool = _192_v17
			_ = _rhs46
			var _rhs47 _dafny.Int = (_dafny.IntOfUint32((_173_v0).Cardinality())).Minus(_179_v5)
			_ = _rhs47
			var _rhs48 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_173_v0, _173_v0), _173_v0)
			_ = _rhs48
			var _lhs34 _dafny.Array = _251_v46
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_251_v46), 0))
			_ = _lhs35
			var _lhs36 *GlobalState = _178_globalState
			_ = _lhs36
			_179_v5 = _rhs45
			_192_v17 = _rhs46
			(_lhs34).ArraySet1(_rhs47, (_lhs35).Int())
			_lhs36.F12 = _rhs48
			var _254_v49 _dafny.Array
			_ = _254_v49
			var _out4 _dafny.Array
			_ = _out4
			_out4 = Companion_Default___.M0((_251_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_251_v46), 0))).Int()).(_dafny.Int), _178_globalState)
			_254_v49 = _out4
			var _255_v50 D0
			_ = _255_v50
			_255_v50 = Companion_D0_.Create_DC2_(_192_v17)
			var _pat_let_tv4 = _192_v17
			_ = _pat_let_tv4
			var _256_v51 _dafny.Sequence
			_ = _256_v51
			_256_v51 = _dafny.SeqOf(func(_pat_let5_0 D0) D0 {
				return func(_257_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let6_0 bool) D0 {
						return func(_258_dt__update_hcf6_h0 bool) D0 {
							return Companion_D0_.Create_DC2_(_258_dt__update_hcf6_h0)
						}(_pat_let6_0)
					}(!(_pat_let_tv4))
				}(_pat_let5_0)
			}(_255_v50), _255_v50)
			_256_v51 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(344))).Uint32(), func(coer13 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_259_v50 D0) func(_dafny.Int) D0 {
				return func(_260_i4 _dafny.Int) D0 {
					return _259_v50
				}
			})(_255_v50)))
			var _261_v52 _dafny.Map
			_ = _261_v52
			_261_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v17, _192_v17)
			var _262_v53 _dafny.Set
			_ = _262_v53
			_262_v53 = _dafny.SetOf((func() bool {
				if (_261_v52).Contains(_192_v17) {
					return (_261_v52).Get(_192_v17).(bool)
				}
				return _192_v17
			})())
			var _263_v54 _dafny.Array
			_ = _263_v54
			var _out5 _dafny.Array
			_ = _out5
			_out5 = Companion_Default___.M0((func() _dafny.Int {
				if !(!(_192_v17)) {
					return _dafny.IntOfInt64(97)
				}
				return (_262_v53).Cardinality()
			})(), _178_globalState)
			_263_v54 = _out5
		} else {
			var _rhs49 _dafny.Int = (_179_v5).Plus(_179_v5)
			_ = _rhs49
			var _rhs50 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("hfnyb")
			_ = _rhs50
			var _lhs37 *GlobalState = _178_globalState
			_ = _lhs37
			var _lhs38 *GlobalState = _178_globalState
			_ = _lhs38
			_lhs37.F11 = _rhs49
			_lhs38.F3 = _rhs50
			var _264_v55 _dafny.Set
			_ = _264_v55
			_264_v55 = _dafny.SetOf(_192_v17, _192_v17)
			var _265_v56 _dafny.Map
			_ = _265_v56
			_265_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v17, _264_v55)
			(_178_globalState).F10 = (_179_v5).Minus((_265_v56).Cardinality())
			var _266_v57 _dafny.Array
			_ = _266_v57
			var _out6 _dafny.Array
			_ = _out6
			_out6 = Companion_Default___.M0(_179_v5, _178_globalState)
			_266_v57 = _out6
			var _267_v58 _dafny.Array
			_ = _267_v58
			var _out7 _dafny.Array
			_ = _out7
			_out7 = Companion_Default___.M0(_179_v5, _178_globalState)
			_267_v58 = _out7
			(_178_globalState).F12 = _192_v17
		}
		var _268_v59 _dafny.Sequence
		_ = _268_v59
		_268_v59 = _dafny.SeqOf(_192_v17, true, _192_v17)
		var _269_v60 _dafny.Array
		_ = _269_v60
		var _nwElement0_8 bool = _192_v17
		_ = _nwElement0_8
		var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(8))
		_ = _nw40
		(_nw40).ArraySet1(_nwElement0_8, 0)
		(_nw40).ArraySet1((_179_v5).Cmp(_179_v5) > 0, 1)
		(_nw40).ArraySet1(_192_v17, 2)
		(_nw40).ArraySet1(_192_v17, 3)
		(_nw40).ArraySet1(_192_v17, 4)
		(_nw40).ArraySet1(!((func() bool {
			if (_194_v19).Contains(_179_v5) {
				return (_194_v19).Get(_179_v5).(bool)
			}
			return _192_v17
		})()), 5)
		(_nw40).ArraySet1(!((_268_v59).Select((Companion_Default___.SafeIndex(_179_v5, _dafny.IntOfUint32((_268_v59).Cardinality()))).Uint32()).(bool)), 6)
		(_nw40).ArraySet1(_192_v17, 7)
		_269_v60 = _nw40
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_269_v60), 0))
		_ = _index40
		(_269_v60).ArraySet1(_192_v17, (_index40).Int())
		var _270_v61 _dafny.MultiSet
		_ = _270_v61
		_270_v61 = _dafny.MultiSetOf(_192_v17)
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_269_v60), 0))
		_ = _index41
		(_269_v60).ArraySet1(Companion_Default___.Fm1(_192_v17, !(_dafny.MultiSetOf((_270_v61).Cardinality(), _179_v5, _179_v5)).Equals((_dafny.MultiSetOf(_179_v5, _179_v5)).Update(_179_v5, Companion_Default___.Abs(_179_v5))), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wbxk")).Cardinality()), _179_v5)), _178_globalState), (_index41).Int())
		var _271_v62 _dafny.Array
		_ = _271_v62
		var _out8 _dafny.Array
		_ = _out8
		_out8 = Companion_Default___.M0(_dafny.IntOfInt64(5), _178_globalState)
		_271_v62 = _out8
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_269_v60), 0))
		_ = _index42
		(_269_v60).ArraySet1(!(!((_269_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_269_v60), 0))).Int()).(bool))) || (!((_269_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_269_v60), 0))).Int()).(bool)) || ((_269_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_269_v60), 0))).Int()).(bool))), (_index42).Int())
	}
	var _272_v63 _dafny.Array
	_ = _272_v63
	var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
	_ = _nw41
	_272_v63 = _nw41
	for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_272_v63), 0))); ; {
		_guard_loop_0, _ok15 := _iter15()
		if !_ok15 {
			break
		}
		var _273_i5 _dafny.Int
		_273_i5 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_273_i5).Sign() != -1) && ((_273_i5).Cmp(_dafny.ArrayLen((_272_v63), 0)) < 0)) {
			(_272_v63).ArraySet1(!(_dafny.MultiSetOf(_179_v5)).Contains(((_dafny.MultiSetOf(_173_v0, _173_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(873))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_274_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_275_i6 _dafny.Int) _dafny.CodePoint {
					return _274_v2
				}
			})(_174_v2))))).Intersection(_dafny.MultiSetOf(_173_v0))).Cardinality()), (_273_i5).Int())
		}
	}
	var _276_v64 _dafny.Set
	_ = _276_v64
	_276_v64 = _dafny.SetOf(_192_v17, Companion_Default___.Fm1(_192_v17, _192_v17, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yhaahdov")).Cardinality()), _178_globalState), _192_v17)
	_192_v17 = (_276_v64).IsProperSubsetOf(_276_v64)
	var _277_v65 _dafny.Map
	_ = _277_v65
	_277_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_272_v63, _179_v5)
	var _278_v66 _dafny.Sequence
	_ = _278_v66
	_278_v66 = _dafny.SeqOf(!(Companion_Default___.Fm1(!(!(_192_v17)), _192_v17, (_277_v65).Cardinality(), _178_globalState)), _192_v17, _192_v17, _192_v17, _192_v17)
	var _279_v67 _dafny.Set
	_ = _279_v67
	_279_v67 = _dafny.SetOf(_179_v5, _dafny.IntOfInt64(706), _179_v5)
	var _280_v68 D0
	_ = _280_v68
	_280_v68 = Companion_D0_.Create_DC3_(_179_v5, (_dafny.IntOfUint32((_278_v66).Cardinality())).Cmp(_179_v5) < 0, _179_v5, ((_279_v67).Cardinality()).Times(_179_v5))
	_280_v68 = _280_v68
	var _281_v69 _dafny.Map
	_ = _281_v69
	_281_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v18, _192_v17)
	var _hi2 _dafny.Int = _dafny.IntOfInt64(66)
	_ = _hi2
	for _282_i7 := (_dafny.Zero).Minus((_281_v69).Cardinality()); _282_i7.Cmp(_hi2) < 0; _282_i7 = _282_i7.Plus(_dafny.One) {
		var _283_v70 *C1
		_ = _283_v70
		var _nw42 *C1 = New_C1_()
		_ = _nw42
		_nw42.Ctor__(_282_i7)
		_283_v70 = _nw42
		(_178_globalState).F12 = (!(_192_v17) || (_192_v17)) == (_192_v17)
		(_283_v70).M2(_272_v63, _dafny.Companion_Sequence_.Concatenate(_173_v0, _173_v0), _dafny.IntOfUint32((_dafny.SeqOf(_283_v70.F18)).Cardinality()), _178_globalState)
		if !(_192_v17) || (_192_v17) {
			var _nw43 *C1 = New_C1_()
			_ = _nw43
			_nw43.Ctor__(_283_v70.F18)
			_283_v70 = _nw43
			var _284_v71 _dafny.Map
			_ = _284_v71
			_284_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(806), _283_v70)
			var _285_v72 _dafny.Sequence
			_ = _285_v72
			_285_v72 = _dafny.SeqOf(_284_v71)
			var _286_v73 _dafny.Sequence
			_ = _286_v73
			_286_v73 = _dafny.SeqOf((_285_v72).Select((Companion_Default___.SafeIndex(_283_v70.F18, _dafny.IntOfUint32((_285_v72).Cardinality()))).Uint32()).(_dafny.Map))
			(_178_globalState).F12 = ((_286_v73).Select((Companion_Default___.SafeIndex(_282_i7, _dafny.IntOfUint32((_286_v73).Cardinality()))).Uint32()).(_dafny.Map)).Equals(_284_v71)
			var _287_v74 _dafny.Array
			_ = _287_v74
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
			_ = _nw44
			_287_v74 = _nw44
			var _288_v75 _dafny.Map
			_ = _288_v75
			_288_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _192_v17)
			var _289_v76 _dafny.Map
			_ = _289_v76
			_289_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_282_i7, _dafny.IntOfInt64(556))
			var _290_v77 _dafny.Array
			_ = _290_v77
			var _nwElement0_9 _dafny.Int = _179_v5
			_ = _nwElement0_9
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(22))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_9, 0)
			(_nw45).ArraySet1((_288_v75).Cardinality(), 1)
			(_nw45).ArraySet1(_282_i7, 2)
			(_nw45).ArraySet1(_283_v70.F18, 3)
			(_nw45).ArraySet1(_282_i7, 4)
			(_nw45).ArraySet1(_179_v5, 5)
			(_nw45).ArraySet1(_282_i7, 6)
			(_nw45).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-670)), 7)
			(_nw45).ArraySet1(_179_v5, 8)
			(_nw45).ArraySet1(_dafny.IntOfInt64(297), 9)
			(_nw45).ArraySet1(_283_v70.F18, 10)
			(_nw45).ArraySet1(_283_v70.F18, 11)
			(_nw45).ArraySet1(_282_i7, 12)
			(_nw45).ArraySet1(_283_v70.F18, 13)
			(_nw45).ArraySet1(_283_v70.F18, 14)
			(_nw45).ArraySet1(_283_v70.F18, 15)
			(_nw45).ArraySet1(_283_v70.F18, 16)
			(_nw45).ArraySet1(_179_v5, 17)
			(_nw45).ArraySet1(Companion_Default___.Fm0(_178_globalState), 18)
			(_nw45).ArraySet1((_289_v76).Cardinality(), 19)
			(_nw45).ArraySet1(_282_i7, 20)
			(_nw45).ArraySet1(_179_v5, 21)
			_290_v77 = _nw45
			var _291_v78 _dafny.Set
			_ = _291_v78
			_291_v78 = _dafny.SetOf(_290_v77)
			var _292_v79 _dafny.Map
			_ = _292_v79
			_292_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_291_v78).Union(_291_v78), ((_193_v18).Cardinality()).Times(_179_v5))
			var _rhs51 _dafny.Int = (func() _dafny.Int {
				if (_292_v79).Contains(_291_v78) {
					return (_292_v79).Get(_291_v78).(_dafny.Int)
				}
				return _dafny.IntOfInt64(110)
			})()
			_ = _rhs51
			var _rhs52 _dafny.Array = _287_v74
			_ = _rhs52
			var _rhs53 _dafny.Int = Companion_Default___.SafeDivisionInt(_283_v70.F18, _179_v5)
			_ = _rhs53
			var _rhs54 bool = _192_v17
			_ = _rhs54
			var _lhs39 *GlobalState = _178_globalState
			_ = _lhs39
			var _lhs40 *GlobalState = _178_globalState
			_ = _lhs40
			_lhs39.F10 = _rhs51
			_287_v74 = _rhs52
			_179_v5 = _rhs53
			_lhs40.F12 = _rhs54
			var _293_v80 _dafny.Map
			_ = _293_v80
			_293_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v17, _283_v70)
			var _294_v81 *C1
			_ = _294_v81
			var _nw46 *C1 = New_C1_()
			_ = _nw46
			_nw46.Ctor__(((_293_v80).Merge(_293_v80)).Cardinality())
			_294_v81 = _nw46
			(_178_globalState).F11 = _283_v70.F18
		} else {
			var _295_v82 *C0
			_ = _295_v82
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__()
			_295_v82 = _nw47
			var _296_v83 _dafny.Map
			_ = _296_v83
			_296_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_295_v82, _283_v70.F18)
			var _297_v84 D3
			_ = _297_v84
			_297_v84 = Companion_D3_.Create_DC7_(_296_v83)
			var _298_v85 _dafny.Map
			_ = _298_v85
			_298_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_297_v84, _282_i7)
			(_283_v70).F18 = Companion_Default___.SafeDivisionInt(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_297_v84, _283_v70.F18)).Merge(_298_v85)).Cardinality(), _dafny.IntOfUint32((_173_v0).Cardinality()))
			var _299_v86 _dafny.MultiSet
			_ = _299_v86
			_299_v86 = _dafny.MultiSetOf(_295_v82, _295_v82)
			var _300_v87 *C1
			_ = _300_v87
			var _nw48 *C1 = New_C1_()
			_ = _nw48
			_nw48.Ctor__((_299_v86).Cardinality())
			_300_v87 = _nw48
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_272_v63), 0))
			_ = _index43
			(_272_v63).ArraySet1(_192_v17, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_272_v63), 0))
			_ = _index44
			(_272_v63).ArraySet1((_283_v70.F18).Cmp(_300_v87.F18) >= 0, (_index44).Int())
			var _301_v88 *C0
			_ = _301_v88
			var _nw49 *C0 = New_C0_()
			_ = _nw49
			_nw49.Ctor__()
			_301_v88 = _nw49
			(_178_globalState).F10 = (_dafny.Zero).Minus(_283_v70.F18)
		}
	}
	var _302_v89 _dafny.Sequence
	_ = _302_v89
	_302_v89 = _dafny.SeqOf(_179_v5, _dafny.IntOfInt64(-221), _179_v5, _179_v5, _dafny.IntOfUint32((_173_v0).Cardinality()))
	var _hi3 _dafny.Int = _179_v5
	_ = _hi3
	for _303_i8 := _dafny.IntOfUint32((_302_v89).Cardinality()); _303_i8.Cmp(_hi3) < 0; _303_i8 = _303_i8.Plus(_dafny.One) {
		var _304_v90 *C0
		_ = _304_v90
		var _nw50 *C0 = New_C0_()
		_ = _nw50
		_nw50.Ctor__()
		_304_v90 = _nw50
		_192_v17 = _192_v17
		var _305_v91 _dafny.Map
		_ = _305_v91
		_305_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _173_v0)
		var _306_v92 *C1
		_ = _306_v92
		var _nw51 *C1 = New_C1_()
		_ = _nw51
		_nw51.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(741), (_305_v91).Cardinality()))
		_306_v92 = _nw51
		var _307_v93 _dafny.Map
		_ = _307_v93
		_307_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v17, _272_v63)
		_307_v93 = _307_v93
	}
	var _308_v94 _dafny.Array
	_ = _308_v94
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(7)
	_ = _len0_7
	var _nw52 _dafny.Array
	_ = _nw52
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw52 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) _dafny.Sequence = (func(_309_v0 _dafny.Sequence, _310_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
			return func(_311_i9 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_309_v0, _dafny.Companion_Sequence_.Update(_309_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_309_v0).Cardinality()), _dafny.IntOfUint32((_309_v0).Cardinality()))).Uint32(), _310_v2))
			}
		})(_173_v0, _174_v2)
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw52 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw52).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw52).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	_308_v94 = _nw52
	_308_v94 = _308_v94
	(_178_globalState).F12 = false
	var _312_v95 _dafny.Array
	_ = _312_v95
	var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
	_ = _nw53
	_312_v95 = _nw53
	_312_v95 = (func() _dafny.Array {
		if _192_v17 {
			return (func() _dafny.Array {
				if _192_v17 {
					return _312_v95
				}
				return _312_v95
			})()
		}
		return _312_v95
	})()
	(_178_globalState).F10 = _179_v5
	var _313_v96 _dafny.Map
	_ = _313_v96
	_313_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v5, _179_v5)
	(_178_globalState).F12 = (_179_v5).Cmp((_dafny.IntOfInt64(-135)).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm4(_178_globalState), (Companion_Default___.SafeIndex((_313_v96).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm4(_178_globalState)).Cardinality()))).Uint32(), _192_v17)).Cardinality()))) != 0
	var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_308_v94), 0))
	_ = _index45
	(_308_v94).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(238))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}((func(_314_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_315_i10 _dafny.Int) _dafny.CodePoint {
			return _314_v2
		}
	})(_174_v2))), (_index45).Int())
	var _316_v97 _dafny.Sequence
	_ = _316_v97
	_316_v97 = _dafny.SeqOf(_173_v0)
	var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_308_v94), 0))
	_ = _index46
	(_308_v94).ArraySet1((_316_v97).Select((Companion_Default___.SafeIndex(_179_v5, _dafny.IntOfUint32((_316_v97).Cardinality()))).Uint32()).(_dafny.Sequence), (_index46).Int())
	if (func() bool {
		if _192_v17 {
			return _192_v17
		}
		return _192_v17
	})() {
		var _317_v98 _dafny.Array
		_ = _317_v98
		var _out9 _dafny.Array
		_ = _out9
		_out9 = Companion_Default___.M0(_179_v5, _178_globalState)
		_317_v98 = _out9
		(_178_globalState).F12 = (Companion_Default___.SafeDivisionInt(_179_v5, _179_v5)).Cmp(_179_v5) >= 0
		(_178_globalState).F10 = (func() _dafny.Int {
			if (_313_v96).Contains(_dafny.IntOfInt64(465)) {
				return (_313_v96).Get(_dafny.IntOfInt64(465)).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_302_v89).Cardinality())
		})()
		(_178_globalState).F10 = _179_v5
		var _318_v99 _dafny.Array
		_ = _318_v99
		var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
		_ = _nw54
		_318_v99 = _nw54
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_318_v99), 0))
		_ = _index47
		(_318_v99).ArraySet1(_dafny.SeqOf(_179_v5, Companion_Default___.Fm0(_178_globalState)), (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_318_v99), 0))
		_ = _index48
		(_318_v99).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-759)), _dafny.SeqOf(_179_v5, _179_v5, _dafny.IntOfInt64(877))), (_index48).Int())
	} else {
		_179_v5 = (_179_v5).Minus(_dafny.IntOfUint32((_278_v66).Cardinality()))
		var _319_v100 _dafny.Sequence
		_ = _319_v100
		_319_v100 = _dafny.SeqOf(_272_v63, _272_v63, _272_v63)
		var _320_v101 *C1
		_ = _320_v101
		var _nw55 *C1 = New_C1_()
		_ = _nw55
		_nw55.Ctor__(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if _192_v17 {
				return _319_v100
			}
			return _319_v100
		})()).Cardinality()))
		_320_v101 = _nw55
		_320_v101 = _320_v101
		_320_v101 = _320_v101
		if !((_179_v5).Cmp(_179_v5) >= 0) {
			var _321_v102 _dafny.Array
			_ = _321_v102
			var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(8))
			_ = _nw56
			_321_v102 = _nw56
			var _322_v103 _dafny.Array
			_ = _322_v103
			var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(20))
			_ = _nw57
			_322_v103 = _nw57
			var _323_v104 _dafny.Array
			_ = _323_v104
			_323_v104 = _322_v103
			var _324_v105 _dafny.Set
			_ = _324_v105
			_324_v105 = _dafny.SetOf(_323_v104, _323_v104, _322_v103, _323_v104)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_321_v102), 0))
			_ = _index49
			(_321_v102).ArraySet1((func() _dafny.Set {
				if _192_v17 {
					return _324_v105
				}
				return _324_v105
			})(), (_index49).Int())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_321_v102), 0))
			_ = _index50
			(_321_v102).ArraySet1(_324_v105, (_index50).Int())
			var _325_v106 _dafny.Array
			_ = _325_v106
			var _out10 _dafny.Array
			_ = _out10
			_out10 = Companion_Default___.M0(_179_v5, _178_globalState)
			_325_v106 = _out10
			_192_v17 = _192_v17
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_325_v106), 0))
			_ = _index51
			(_325_v106).ArraySet1(_320_v101.F18, (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_325_v106), 0))
			_ = _index52
			var _rhs55 _dafny.Int = _320_v101.F18
			_ = _rhs55
			var _rhs56 _dafny.Array = _272_v63
			_ = _rhs56
			var _lhs41 _dafny.Array = _325_v106
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_325_v106), 0))
			_ = _lhs42
			(_lhs41).ArraySet1(_rhs55, (_lhs42).Int())
			_272_v63 = _rhs56
			var _326_v107 *C0
			_ = _326_v107
			var _nw58 *C0 = New_C0_()
			_ = _nw58
			_nw58.Ctor__()
			_326_v107 = _nw58
			var _327_v108 _dafny.Map
			_ = _327_v108
			_327_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_192_v17), _326_v107)
			(_178_globalState).F12 = Companion_Default___.Fm1(!((((_327_v108).Update(_192_v17, _326_v107)).Cardinality()).Cmp((_279_v67).Cardinality()) == 0), _192_v17, _179_v5, _178_globalState)
		} else {
			(_320_v101).M2(_272_v63, _173_v0, Companion_Default___.SafeModuloInt((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(237))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_328_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_329_i11 _dafny.Int) _dafny.CodePoint {
					return _328_v2
				}
			})(_174_v2))))).Cardinality(), _dafny.IntOfInt64(243)), _178_globalState)
			_278_v66 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm1(_192_v17, _192_v17, _179_v5, _178_globalState)), _dafny.Companion_Sequence_.Concatenate(_278_v66, _278_v66))
			(_178_globalState).F11 = (_dafny.Zero).Minus(_dafny.IntOfUint32(((_308_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_308_v94), 0))).Int()).(_dafny.Sequence)).Cardinality()))
			(_178_globalState).F12 = ((_dafny.Zero).Minus(_179_v5)).Cmp(((_193_v18).Cardinality()).Minus(_179_v5)) == 0
			var _330_v109 *C1
			_ = _330_v109
			var _nw59 *C1 = New_C1_()
			_ = _nw59
			_nw59.Ctor__(_320_v101.F18)
			_330_v109 = _nw59
		}
		var _331_v110 _dafny.Sequence
		_ = _331_v110
		_331_v110 = _dafny.SeqOf(_320_v101, _320_v101, _320_v101, _320_v101)
		_320_v101 = (_331_v110).Select((Companion_Default___.SafeIndex(_179_v5, _dafny.IntOfUint32((_331_v110).Cardinality()))).Uint32()).(*C1)
	}
	var _332_v111 _dafny.MultiSet
	_ = _332_v111
	_332_v111 = _dafny.MultiSetOf(_192_v17, _192_v17, Companion_Default___.Fm1(_192_v17, _192_v17, _179_v5, _178_globalState), _192_v17)
	var _333_v112 _dafny.Map
	_ = _333_v112
	_333_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v17, _192_v17)
	var _334_v113 _dafny.Map
	_ = _334_v113
	_334_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_278_v66, !((func() bool {
		if (_333_v112).Contains(_192_v17) {
			return (_333_v112).Get(_192_v17).(bool)
		}
		return _192_v17
	})()))
	(_178_globalState).F12 = !(Companion_Default___.Fm16((_332_v111).Cardinality(), _178_globalState)).Equals(_334_v113)
	var _335_v114 _dafny.Array
	_ = _335_v114
	var _nw60 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
	_ = _nw60
	_335_v114 = _nw60
	var _336_v115 *C0
	_ = _336_v115
	var _nw61 *C0 = New_C0_()
	_ = _nw61
	_nw61.Ctor__()
	_336_v115 = _nw61
	var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_335_v114), 0))
	_ = _index53
	(_335_v114).ArraySet1(_336_v115, (_index53).Int())
	var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_335_v114), 0))
	_ = _index54
	(_335_v114).ArraySet1(_336_v115, (_index54).Int())
	_dafny.Print(_173_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_174_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_v3).Equals(_dafny.SetOf(_dafny.CodePoint('k'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_176_v4).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('n')), _dafny.SetOf(_dafny.CodePoint('k')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_178_globalState.F3.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F4().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_178_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_178_globalState.F11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_178_globalState.F12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_globalState).F14()).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('n')), _dafny.SetOf(_dafny.CodePoint('k')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_178_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_179_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_192_v17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_193_v18).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-2161))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_194_v19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v63).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v63).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v63).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v63).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v63).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v63).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v63).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v64).Equals(_dafny.SetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v65).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_278_v66, _dafny.SeqOf(true, false, false, false, false, false, false, false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v67).Equals(_dafny.SetOf(_dafny.IntOfInt64(-2161), _dafny.IntOfInt64(706))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_280_v68).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_280_v68).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_280_v68).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_280_v68).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v69).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-2161)), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_302_v89, _dafny.SeqOf(_dafny.One, _dafny.IntOfInt64(-221), _dafny.One, _dafny.One, _dafny.IntOfInt64(403))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v94).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v94).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v94).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v94).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v94).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v94).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v94).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_313_v96).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_316_v97, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("yokkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_332_v111).Equals(_dafny.MultiSetOf(false, false, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_333_v112).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_334_v113).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, false, false, false, false, false, false, false, false, false, false), true)))
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
	Cf3 _dafny.Int
	Cf4 _dafny.Int
	Cf5 _dafny.Array
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 _dafny.Int, Cf4 _dafny.Int, Cf5 _dafny.Array) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf6 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf6 bool) D0 {
	return D0{D0_DC2{Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf7  _dafny.Int
	Cf8  bool
	Cf9  _dafny.Int
	Cf10 _dafny.Int
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf7 _dafny.Int, Cf8 bool, Cf9 _dafny.Int, Cf10 _dafny.Int) D0 {
	return D0{D0_DC3{Cf7, Cf8, Cf9, Cf10}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Int
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D0_DC3).Cf7
}

func (_this D0) Dtor_cf8() bool {
	return _this.Get_().(D0_DC3).Cf8
}

func (_this D0) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D0_DC3).Cf9
}

func (_this D0) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D0_DC3).Cf10
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf6 == data2.Cf6
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Cmp(data2.Cf10) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
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
	Cf12 bool
	Cf13 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf12 bool, Cf13 _dafny.Int) D1 {
	return D1{D1_DC5{Cf12, Cf13}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf11 _dafny.CodePoint
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf11 _dafny.CodePoint) D1 {
	return D1{D1_DC4{Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false, _dafny.Zero)
}

func (_this D1) Dtor_cf12() bool {
	return _this.Get_().(D1_DC5).Cf12
}

func (_this D1) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf13
}

func (_this D1) Dtor_cf11() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf12 == data2.Cf12 && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf11 == data2.Cf11
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
	Cf14 _dafny.Sequence
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf14 _dafny.Sequence) D2 {
	return D2{D2_DC6{Cf14}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D2) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf14
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + data.Cf14.VerbatimString(true) + ")"
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
			return ok && data1.Cf14.Equals(data2.Cf14)
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
	Cf16 _dafny.Sequence
	Cf17 bool
	Cf18 bool
	Cf19 _dafny.Sequence
	Cf20 bool
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf16 _dafny.Sequence, Cf17 bool, Cf18 bool, Cf19 _dafny.Sequence, Cf20 bool) D3 {
	return D3{D3_DC8{Cf16, Cf17, Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC7 struct {
	Cf15 _dafny.Map
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf15 _dafny.Map) D3 {
	return D3{D3_DC7{Cf15}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(_dafny.EmptySeq, false, false, _dafny.EmptySeq, false)
}

func (_this D3) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf16
}

func (_this D3) Dtor_cf17() bool {
	return _this.Get_().(D3_DC8).Cf17
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC8).Cf18
}

func (_this D3) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf19
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC8).Cf20
}

func (_this D3) Dtor_cf15() _dafny.Map {
	return _this.Get_().(D3_DC7).Cf15
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + data.Cf19.VerbatimString(true) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf16.Equals(data2.Cf16) && data1.Cf17 == data2.Cf17 && data1.Cf18 == data2.Cf18 && data1.Cf19.Equals(data2.Cf19) && data1.Cf20 == data2.Cf20
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
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

type D4_DC10 struct {
	Cf22 _dafny.CodePoint
	Cf23 _dafny.Array
	Cf24 _dafny.Int
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf22 _dafny.CodePoint, Cf23 _dafny.Array, Cf24 _dafny.Int) D4 {
	return D4{D4_DC10{Cf22, Cf23, Cf24}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC9 struct {
	Cf21 *C0
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf21 *C0) D4 {
	return D4{D4_DC9{Cf21}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC11 struct {
	Cf25 D4
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf25 D4) D4 {
	return D4{D4_DC11{Cf25}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(_dafny.CodePoint('D'), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D4) Dtor_cf22() _dafny.CodePoint {
	return _this.Get_().(D4_DC10).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Array {
	return _this.Get_().(D4_DC10).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf24
}

func (_this D4) Dtor_cf21() *C0 {
	return _this.Get_().(D4_DC9).Cf21
}

func (_this D4) Dtor_cf25() D4 {
	return _this.Get_().(D4_DC11).Cf25
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf25) + ")"
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
			return ok && data1.Cf22 == data2.Cf22 && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf21 == data2.Cf21
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
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

type D5_DC12 struct {
	Cf26 _dafny.Array
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf26 _dafny.Array) D5 {
	return D5{D5_DC12{Cf26}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf26() _dafny.Array {
	return _this.Get_().(D5_DC12).Cf26
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf26) + ")"
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
			return ok && data1.Cf26 == data2.Cf26
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
	Cf28 bool
	Cf29 _dafny.CodePoint
	Cf30 _dafny.Array
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf28 bool, Cf29 _dafny.CodePoint, Cf30 _dafny.Array) D6 {
	return D6{D6_DC14{Cf28, Cf29, Cf30}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC13 struct {
	Cf27 _dafny.Map
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf27 _dafny.Map) D6 {
	return D6{D6_DC13{Cf27}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC14_(false, _dafny.CodePoint('D'), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D6) Dtor_cf28() bool {
	return _this.Get_().(D6_DC14).Cf28
}

func (_this D6) Dtor_cf29() _dafny.CodePoint {
	return _this.Get_().(D6_DC14).Cf29
}

func (_this D6) Dtor_cf30() _dafny.Array {
	return _this.Get_().(D6_DC14).Cf30
}

func (_this D6) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D6_DC13).Cf27
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf27) + ")"
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
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
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

type D7_DC15 struct {
	Cf31 _dafny.Sequence
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf31 _dafny.Sequence) D7 {
	return D7{D7_DC15{Cf31}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D7) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D7_DC15).Cf31
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf31.Equals(data2.Cf31)
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
	Cf33 bool
	Cf34 _dafny.Int
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf33 bool, Cf34 _dafny.Int) D8 {
	return D8{D8_DC17{Cf33, Cf34}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

type D8_DC18 struct {
	Cf35 _dafny.Sequence
	Cf36 _dafny.Sequence
	Cf37 bool
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf35 _dafny.Sequence, Cf36 _dafny.Sequence, Cf37 bool) D8 {
	return D8{D8_DC18{Cf35, Cf36, Cf37}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

type D8_DC16 struct {
	Cf32 _dafny.Map
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf32 _dafny.Map) D8 {
	return D8{D8_DC16{Cf32}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

type D8_DC19 struct {
	Cf38 D8
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf38 D8) D8 {
	return D8{D8_DC19{Cf38}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC17_(false, _dafny.Zero)
}

func (_this D8) Dtor_cf33() bool {
	return _this.Get_().(D8_DC17).Cf33
}

func (_this D8) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D8_DC17).Cf34
}

func (_this D8) Dtor_cf35() _dafny.Sequence {
	return _this.Get_().(D8_DC18).Cf35
}

func (_this D8) Dtor_cf36() _dafny.Sequence {
	return _this.Get_().(D8_DC18).Cf36
}

func (_this D8) Dtor_cf37() bool {
	return _this.Get_().(D8_DC18).Cf37
}

func (_this D8) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D8_DC16).Cf32
}

func (_this D8) Dtor_cf38() D8 {
	return _this.Get_().(D8_DC19).Cf38
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D8_DC18:
		{
			return "D8.DC18" + "(" + data.Cf35.VerbatimString(true) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D8_DC16:
		{
			return "D8.DC16" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf38) + ")"
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
			return ok && data1.Cf33 == data2.Cf33 && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D8_DC18:
		{
			data2, ok := other.Get_().(D8_DC18)
			return ok && data1.Cf35.Equals(data2.Cf35) && data1.Cf36.Equals(data2.Cf36) && data1.Cf37 == data2.Cf37
		}
	case D8_DC16:
		{
			data2, ok := other.Get_().(D8_DC16)
			return ok && data1.Cf32.Equals(data2.Cf32)
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf38.Equals(data2.Cf38)
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

// Definition of class GlobalState
type GlobalState struct {
	F3   _dafny.Sequence
	F10  _dafny.Int
	F11  _dafny.Int
	F12  bool
	F15  _dafny.CodePoint
	_f0  _dafny.Int
	_f1  bool
	_f2  _dafny.Int
	_f4  _dafny.Sequence
	_f5  _dafny.Int
	_f6  _dafny.Int
	_f7  bool
	_f8  bool
	_f9  _dafny.Int
	_f13 bool
	_f14 _dafny.MultiSet
	_f16 _dafny.Int
	_f17 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = _dafny.EmptySeq
	_this.F10 = _dafny.Zero
	_this.F11 = _dafny.Zero
	_this.F12 = false
	_this.F15 = _dafny.CodePoint('D')
	_this._f0 = _dafny.Zero
	_this._f1 = false
	_this._f2 = _dafny.Zero
	_this._f4 = _dafny.EmptySeq
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f7 = false
	_this._f8 = false
	_this._f9 = _dafny.Zero
	_this._f13 = false
	_this._f14 = _dafny.EmptyMultiSet
	_this._f16 = _dafny.Zero
	_this._f17 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 _dafny.Int, f3 _dafny.Sequence, f4 _dafny.Sequence, f5 _dafny.Int, f6 _dafny.Int, f7 bool, f8 bool, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Int, f12 bool, f13 bool, f14 _dafny.MultiSet, f15 _dafny.CodePoint, f16 _dafny.Int, f17 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this).F11 = f11
		(_this).F12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F4() _dafny.Sequence {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
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
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() _dafny.MultiSet {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Int {
	{
		return _this._f17
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
func (_this *C0) Fm6(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.IntOfInt64(-495)).Times((_dafny.SetOf(!(false), false)).Cardinality())), (func() _dafny.Int {
			if false {
				return _dafny.IntOfInt64(-643)
			}
			return _dafny.IntOfInt64(993)
		})())
	}
}
func (_this *C0) Fm7(p0 bool, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cdmixbs"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("blliiliu"), _dafny.UnicodeSeqOfUtf8Bytes("vthendu")))
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F18 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F18 = _dafny.Zero
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

func (_this *C1) Ctor__(f18 _dafny.Int) {
	{
		(_this).F18 = f18
	}
}
func (_this *C1) M1(p0 _dafny.Int, p1 bool, p2 _dafny.Array, p3 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _337_v0 _dafny.Array
		_ = _337_v0
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
		_ = _nw62
		_337_v0 = _nw62
		var _338_v1 _dafny.CodePoint
		_ = _338_v1
		_338_v1 = _dafny.CodePoint('a')
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_337_v0), 0))
		_ = _index55
		(_337_v0).ArraySet1CodePoint(_338_v1, (_index55).Int())
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_337_v0), 0))
		_ = _index56
		(_337_v0).ArraySet1CodePoint(_dafny.CodePoint('n'), (_index56).Int())
		var _339_v2 _dafny.Array
		_ = _339_v2
		var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw63
		_339_v2 = _nw63
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_339_v2), 0))); ; {
			_guard_loop_1, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _340_i0 _dafny.Int
			_340_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_340_i0).Sign() != -1) && ((_340_i0).Cmp(_dafny.ArrayLen((_339_v2), 0)) < 0)) {
				(_339_v2).ArraySet1(p1, (_340_i0).Int())
			}
		}
		(_this).F18 = p0
		var _341_v3 _dafny.Array
		_ = _341_v3
		var _out11 _dafny.Array
		_ = _out11
		_out11 = Companion_Default___.M0(p0, globalState)
		_341_v3 = _out11
		var _hi4 _dafny.Int = Companion_Default___.Fm0(globalState)
		_ = _hi4
		for _342_i1 := _this.F18; _342_i1.Cmp(_hi4) < 0; _342_i1 = _342_i1.Plus(_dafny.One) {
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_339_v2), 0))
			_ = _index57
			(_339_v2).ArraySet1(true, (_index57).Int())
			var _343_v4 _dafny.MultiSet
			_ = _343_v4
			_343_v4 = _dafny.MultiSetOf(_342_i1)
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_339_v2), 0))
			_ = _index58
			(_339_v2).ArraySet1((_343_v4).IsProperSubsetOf(_343_v4), (_index58).Int())
			var _344_v5 *C0
			_ = _344_v5
			var _nw64 *C0 = New_C0_()
			_ = _nw64
			_nw64.Ctor__()
			_344_v5 = _nw64
			var _345_v6 _dafny.Sequence
			_ = _345_v6
			_345_v6 = _dafny.UnicodeSeqOfUtf8Bytes("bfkwvkpem")
			var _346_v7 _dafny.Map
			_ = _346_v7
			_346_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_345_v6, _dafny.IntOfInt64(938))
			_346_v7 = _346_v7
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_341_v3), 0))
			_ = _index59
			(_341_v3).ArraySet1((_dafny.Zero).Minus(p3), (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_341_v3), 0))
			_ = _index60
			(_341_v3).ArraySet1((_dafny.Zero).Minus(((_dafny.Zero).Minus((_this.F18).Plus(_342_i1))).Plus(_this.F18)), (_index60).Int())
		}
		if Companion_Default___.Fm1(p1, _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("rrme"), Companion_Default___.Fm8(globalState)), p3, globalState) {
			var _347_v8 _dafny.Map
			_ = _347_v8
			_347_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v3, Companion_D0_.Create_DC0_(p3))
			_347_v8 = _347_v8
			var _348_v9 _dafny.MultiSet
			_ = _348_v9
			_348_v9 = _dafny.MultiSetOf(true, p1)
			var _349_v10 _dafny.Map
			_ = _349_v10
			_349_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v0, (_348_v9).IsProperSubsetOf(_dafny.MultiSetOf(p1, p1, p1, p1)))
			var _350_v11 _dafny.Map
			_ = _350_v11
			_350_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(globalState), _337_v0)
			_349_v10 = (_349_v10).Update((func() _dafny.Array {
				if (_350_v11).Contains(p0) {
					return (_350_v11).Get(p0).(_dafny.Array)
				}
				return _337_v0
			})(), p1)
			(globalState).F11 = p0
			(globalState).F12 = p1
			var _351_v12 D0
			_ = _351_v12
			_351_v12 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-927))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}(func(_352_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}))).Cardinality()))
			var _353_v13 _dafny.Set
			_ = _353_v13
			_353_v13 = _dafny.SetOf(Companion_D0_.Create_DC0_(p3), _351_v12, Companion_D0_.Create_DC0_(_dafny.IntOfInt64(537)))
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_339_v2), 0))
			_ = _index61
			(_339_v2).ArraySet1((_353_v13).IsProperSubsetOf(_353_v13), (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_339_v2), 0))
			_ = _index62
			var _rhs57 bool = false
			_ = _rhs57
			var _rhs58 bool = p1
			_ = _rhs58
			var _lhs43 *GlobalState = globalState
			_ = _lhs43
			var _lhs44 _dafny.Array = _339_v2
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_339_v2), 0))
			_ = _lhs45
			_lhs43.F12 = _rhs57
			(_lhs44).ArraySet1(_rhs58, (_lhs45).Int())
		} else {
			var _354_v14 _dafny.Sequence
			_ = _354_v14
			_354_v14 = _dafny.SeqOf((_337_v0).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_337_v0), 0))).Int()), (_337_v0).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_337_v0), 0))).Int()))
			var _355_v15 _dafny.Sequence
			_ = _355_v15
			_355_v15 = _354_v14
			var _356_v16 _dafny.Map
			_ = _356_v16
			_356_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_355_v15, p1)
			_356_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_354_v14, true)
			var _357_v17 _dafny.Sequence
			_ = _357_v17
			_357_v17 = _dafny.SeqOf(p1)
			var _358_v18 _dafny.Map
			_ = _358_v18
			_358_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_354_v14, _357_v17)
			_358_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_354_v14, _354_v14), _357_v17)
			var _359_v19 _dafny.Map
			_ = _359_v19
			_359_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v2, p1)
			var _360_v20 D1
			_ = _360_v20
			_360_v20 = Companion_D1_.Create_DC4_(_338_v1)
			var _rhs59 _dafny.Map = (_359_v19).Merge(_359_v19)
			_ = _rhs59
			var _rhs60 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs60
			var _rhs61 _dafny.Int = p3
			_ = _rhs61
			var _rhs62 D1 = _360_v20
			_ = _rhs62
			var _lhs46 *GlobalState = globalState
			_ = _lhs46
			var _lhs47 *GlobalState = globalState
			_ = _lhs47
			_359_v19 = _rhs59
			_lhs46.F11 = _rhs60
			_lhs47.F10 = _rhs61
			_360_v20 = _rhs62
			r0 = _339_v2
			var _361_v21 _dafny.Map
			_ = _361_v21
			_361_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _362_v22 _dafny.Set
			_ = _362_v22
			_362_v22 = _dafny.SetOf(_361_v21, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), _361_v21)
			(globalState).F12 = (_362_v22).IsDisjointFrom(_362_v22)
		}
		r0 = _339_v2
		r1 = p3
		return r0, r1
	}
}
func (_this *C1) M2(p0 _dafny.Array, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) {
	{
		var _363_v0 _dafny.Array
		_ = _363_v0
		var _nwElement0_10 _dafny.Int = (_dafny.Zero).Minus(p2)
		_ = _nwElement0_10
		var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(2))
		_ = _nw65
		(_nw65).ArraySet1(_nwElement0_10, 0)
		(_nw65).ArraySet1((_this.F18).Minus(_dafny.IntOfInt64(-962)), 1)
		_363_v0 = _nw65
		var _364_v1 bool
		_ = _364_v1
		_364_v1 = false
		_363_v0 = (func() _dafny.Array {
			if _364_v1 {
				return _363_v0
			}
			return _363_v0
		})()
		var _365_i0 _dafny.Int
		_ = _365_i0
		_365_i0 = _dafny.Zero
		{
			for (func() bool {
				if true {
					return (_this.F18).Cmp(_dafny.IntOfInt64(181)) >= 0
				}
				return (p2).Cmp(p2) >= 0
			})() {
				{
					if (_365_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_365_i0 = (_365_i0).Plus(_dafny.One)
					var _366_v2 _dafny.Array
					_ = _366_v2
					var _len0_8 _dafny.Int = _dafny.IntOfInt64(25)
					_ = _len0_8
					var _nw66 _dafny.Array
					_ = _nw66
					if _len0_8.Cmp(_dafny.Zero) == 0 {
						_nw66 = _dafny.NewArray(_len0_8)
					} else {
						var _init8 func(_dafny.Int) _dafny.Sequence = (func(_367_v1 bool, _368_p2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
							return func(_369_i1 _dafny.Int) _dafny.Sequence {
								return (func() _dafny.Sequence {
									if _367_v1 {
										return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(145))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
											return func(arg18 _dafny.Int) interface{} {
												return coer18(arg18)
											}
										}((func(_370_v1 bool) func(_dafny.Int) _dafny.Int {
											return func(_371_i2 _dafny.Int) _dafny.Int {
												return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_370_v1, _370_v1, _370_v1))).Cardinality()
											}
										})(_367_v1)))
									}
									return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(154))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
										return func(arg19 _dafny.Int) interface{} {
											return coer19(arg19)
										}
									}((func(_372_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
										return func(_373_i3 _dafny.Int) _dafny.Int {
											return _372_p2
										}
									})(_368_p2)))
								})()
							}
						})(_364_v1, p2)
						_ = _init8
						var _element0_8 = _init8(_dafny.Zero)
						_ = _element0_8
						_nw66 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
						(_nw66).ArraySet1(_element0_8, 0)
						var _nativeLen0_8 = (_len0_8).Int()
						_ = _nativeLen0_8
						for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
							(_nw66).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
						}
					}
					_366_v2 = _nw66
					var _374_v3 _dafny.Sequence
					_ = _374_v3
					_374_v3 = _dafny.SeqOf(p2)
					var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_366_v2), 0))
					_ = _index63
					(_366_v2).ArraySet1(_374_v3, (_index63).Int())
					var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_366_v2), 0))
					_ = _index64
					(_366_v2).ArraySet1(_374_v3, (_index64).Int())
					(globalState).F12 = (func() bool {
						if (_dafny.IntOfUint32((p1).Cardinality())).Cmp(_dafny.IntOfUint32((p1).Cardinality())) >= 0 {
							return !(_364_v1)
						}
						return _364_v1
					})()
					var _375_v4 *C0
					_ = _375_v4
					var _nw67 *C0 = New_C0_()
					_ = _nw67
					_nw67.Ctor__()
					_375_v4 = _nw67
					var _376_v5 _dafny.Map
					_ = _376_v5
					_376_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_375_v4, (_this.F18).Minus(p2))
					var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((p0), 0))
					_ = _index65
					(p0).ArraySet1(!(_364_v1), (_index65).Int())
					var _377_v6 D3
					_ = _377_v6
					_377_v6 = Companion_D3_.Create_DC7_(_376_v5)
					var _378_v7 D0
					_ = _378_v7
					_378_v7 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(908))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg20 _dafny.Int) interface{} {
							return coer20(arg20)
						}
					}(func(_379_i4 _dafny.Int) _dafny.Int {
						return _this.F18
					}))).Cardinality()), _this.F18, p2, _this.F18, p0)
					var _380_v8 _dafny.Sequence
					_ = _380_v8
					_380_v8 = _dafny.SeqOf(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(250), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mocct")).Cardinality()), (_dafny.Zero).Minus(_this.F18), _this.F18, p0))
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((p0), 0))
					_ = _index66
					var _rhs63 _dafny.Map = (_377_v6).Dtor_cf15()
					_ = _rhs63
					var _rhs64 bool = _364_v1
					_ = _rhs64
					var _rhs65 bool = !_dafny.Companion_Sequence_.Contains(_380_v8, _378_v7)
					_ = _rhs65
					var _rhs66 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(25))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg21 _dafny.Int) interface{} {
							return coer21(arg21)
						}
					}(func(_381_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('x')
					})))).Cardinality())
					_ = _rhs66
					var _lhs48 _dafny.Array = p0
					_ = _lhs48
					var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((p0), 0))
					_ = _lhs49
					var _lhs50 *GlobalState = globalState
					_ = _lhs50
					_376_v5 = _rhs63
					(_lhs48).ArraySet1(_rhs64, (_lhs49).Int())
					_364_v1 = _rhs65
					_lhs50.F10 = _rhs66
					var _382_v9 D0
					_ = _382_v9
					_382_v9 = Companion_D0_.Create_DC2_(_364_v1)
					var _383_v10 _dafny.Map
					_ = _383_v10
					_383_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((p0), 0))).Int()).(bool), _382_v9)
					var _384_v11 _dafny.Map
					_ = _384_v11
					_384_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v10, _this.F18)
					var _385_v12 _dafny.CodePoint
					_ = _385_v12
					_385_v12 = _dafny.CodePoint('c')
					var _386_v13 _dafny.MultiSet
					_ = _386_v13
					_386_v13 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _385_v12)).Cardinality()))
					var _387_v15 _dafny.Set
					_ = _387_v15
					_387_v15 = _dafny.SetOf(_383_v10, _383_v10)
					(globalState).F12 = (_384_v11).Equals((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v10, (_386_v13).Cardinality())).Merge(func() _dafny.Map {
						var _coll15 = _dafny.NewMapBuilder()
						_ = _coll15
						for _iter17 := _dafny.Iterate((_387_v15).Elements()); ; {
							_compr_15, _ok17 := _iter17()
							if !_ok17 {
								break
							}
							var _388_v14 _dafny.Map
							_388_v14 = interface{}(_compr_15).(_dafny.Map)
							if (_387_v15).Contains(_388_v14) {
								_coll15.Add(_388_v14, _dafny.IntOfInt64(2))
							}
						}
						return _coll15.ToMap()
					}()))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _389_v16 _dafny.Array
		_ = _389_v16
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_9
		var _nw68 _dafny.Array
		_ = _nw68
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw68 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) D3 = (func(_390_p1 _dafny.Sequence, _391_p2 _dafny.Int, _392_v1 bool) func(_dafny.Int) D3 {
				return func(_393_i7 _dafny.Int) D3 {
					return Companion_D3_.Create_DC8_(_dafny.SeqOf(_dafny.IntOfInt64(896), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_390_p1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.IntOfUint32((_390_p1).Cardinality()))).Uint32(), _dafny.CodePoint('v'))).Cardinality()), _391_p2))).Cardinality()), _392_v1, _392_v1, _390_p1, _392_v1)
				}
			})(p1, p2, _364_v1)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw68 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw68).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw68).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_389_v16 = _nw68
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_389_v16), 0))); ; {
			_guard_loop_2, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _394_i6 _dafny.Int
			_394_i6 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_394_i6).Sign() != -1) && ((_394_i6).Cmp(_dafny.ArrayLen((_389_v16), 0)) < 0)) {
				(_389_v16).ArraySet1((func() D3 {
					if false {
						return Companion_D3_.Create_DC8_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(998))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg22 _dafny.Int) interface{} {
								return coer22(arg22)
							}
						}((func(_395_v1 bool) func(_dafny.Int) _dafny.Int {
							return func(_396_i8 _dafny.Int) _dafny.Int {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_395_v1, _this.F18)).Cardinality()
							}
						})(_364_v1))), _364_v1, _364_v1, p1, _364_v1)
					}
					return Companion_D3_.Create_DC8_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(49))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg23 _dafny.Int) interface{} {
							return coer23(arg23)
						}
					}((func(_397_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_398_i9 _dafny.Int) _dafny.Int {
							return _397_p2
						}
					})(p2))), !(_364_v1), _364_v1, p1, _364_v1)
				})(), (_394_i6).Int())
			}
		}
		if false {
			if (p2).Cmp(p2) >= 0 {
				var _399_v17 *C0
				_ = _399_v17
				var _nw69 *C0 = New_C0_()
				_ = _nw69
				_nw69.Ctor__()
				_399_v17 = _nw69
				var _rhs67 bool = _364_v1
				_ = _rhs67
				var _rhs68 bool = true
				_ = _rhs68
				var _rhs69 _dafny.Int = p2
				_ = _rhs69
				var _rhs70 _dafny.Int = p2
				_ = _rhs70
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				var _lhs52 *GlobalState = globalState
				_ = _lhs52
				var _lhs53 *GlobalState = globalState
				_ = _lhs53
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				_lhs51.F12 = _rhs67
				_lhs52.F12 = _rhs68
				_lhs53.F10 = _rhs69
				_lhs54.F11 = _rhs70
				var _400_v18 _dafny.Map
				_ = _400_v18
				_400_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_364_v1), p1)
				var _401_v19 _dafny.Sequence
				_ = _401_v19
				_401_v19 = p1
				var _402_v20 _dafny.Sequence
				_ = _402_v20
				_402_v20 = _dafny.SeqOf(_this.F18)
				var _403_v21 D3
				_ = _403_v21
				_403_v21 = Companion_D3_.Create_DC8_(_402_v20, _364_v1, _364_v1, p1, _364_v1)
				var _404_v22 _dafny.Array
				_ = _404_v22
				var _nwElement0_11 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("lglnyxquc")
				_ = _nwElement0_11
				var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(29))
				_ = _nw70
				(_nw70).ArraySet1(_nwElement0_11, 0)
				(_nw70).ArraySet1(p1, 1)
				(_nw70).ArraySet1((func() _dafny.Sequence {
					if (_400_v18).Contains(_364_v1) {
						return (_400_v18).Get(_364_v1).(_dafny.Sequence)
					}
					return p1
				})(), 2)
				(_nw70).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-156))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}(func(_405_i10 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				})), 3)
				(_nw70).ArraySet1(p1, 4)
				(_nw70).ArraySet1(p1, 5)
				(_nw70).ArraySet1((p1), 6)
				(_nw70).ArraySet1(p1, 7)
				(_nw70).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(738))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}(func(_406_i11 _dafny.Int) _dafny.CodePoint {
					return (Companion_D1_.Create_DC4_(_dafny.CodePoint('l'))).Dtor_cf11()
				})), 8)
				(_nw70).ArraySet1(p1, 9)
				(_nw70).ArraySet1(p1, 10)
				(_nw70).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, (_403_v21).Dtor_cf19()), 11)
				(_nw70).ArraySet1(p1, 12)
				(_nw70).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hfqwixbp"), 13)
				(_nw70).ArraySet1(p1, 14)
				(_nw70).ArraySet1(p1, 15)
				(_nw70).ArraySet1(p1, 16)
				(_nw70).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, p1), 17)
				(_nw70).ArraySet1(p1, 18)
				(_nw70).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kyqwj"), 19)
				(_nw70).ArraySet1((_403_v21).Dtor_cf19(), 20)
				(_nw70).ArraySet1(p1, 21)
				(_nw70).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("aqevck"), 22)
				(_nw70).ArraySet1(p1, 23)
				(_nw70).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p1, p1), 24)
				(_nw70).ArraySet1(p1, 25)
				(_nw70).ArraySet1(p1, 26)
				(_nw70).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}(func(_407_i12 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				})), 27)
				(_nw70).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mgyrwllh"), 28)
				_404_v22 = _nw70
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_404_v22), 0))
				_ = _index67
				(_404_v22).ArraySet1(p1, (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_404_v22), 0))
				_ = _index68
				(_404_v22).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(360))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg27 _dafny.Int) interface{} {
						return coer27(arg27)
					}
				}(func(_408_i13 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), (_index68).Int())
				var _409_v23 _dafny.Sequence
				_ = _409_v23
				_409_v23 = _dafny.SeqOf(_399_v17)
				_399_v17 = (_409_v23).Select((Companion_Default___.SafeIndex((_this.F18).Minus((_dafny.Zero).Minus(_this.F18)), _dafny.IntOfUint32((_409_v23).Cardinality()))).Uint32()).(*C0)
				var _410_v24 _dafny.Array
				_ = _410_v24
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
				_ = _nw71
				_410_v24 = _nw71
				var _411_v25 _dafny.Sequence
				_ = _411_v25
				_411_v25 = _dafny.SeqOf(_363_v0, _363_v0, _363_v0, _363_v0)
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_410_v24), 0))
				_ = _index69
				(_410_v24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_363_v0, _363_v0, _363_v0), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqOf(_363_v0, _363_v0, _363_v0)).Cardinality()))).Uint32(), _363_v0), _411_v25), (_index69).Int())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_410_v24), 0))
				_ = _index70
				(_410_v24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_411_v25, _dafny.SeqOf(_363_v0, _363_v0)), (_index70).Int())
			} else {
				_363_v0 = _363_v0
				(globalState).F15 = Companion_Default___.Fm9(_364_v1, _364_v1, globalState)
				var _412_v26 _dafny.Array
				_ = _412_v26
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_10
				var _nw72 _dafny.Array
				_ = _nw72
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw72 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.CodePoint = func(_413_i14 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					}
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw72 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw72).ArraySet1CodePoint(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw72).ArraySet1CodePoint(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_412_v26 = _nw72
				_412_v26 = _412_v26
				_364_v1 = (func() bool {
					if _364_v1 {
						return _364_v1
					}
					return _364_v1
				})()
				var _414_v27 *C0
				_ = _414_v27
				var _nw73 *C0 = New_C0_()
				_ = _nw73
				_nw73.Ctor__()
				_414_v27 = _nw73
			}
			var _415_v28 _dafny.CodePoint
			_ = _415_v28
			_415_v28 = _dafny.CodePoint('h')
			var _416_v29 _dafny.Map
			_ = _416_v29
			_416_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(350))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_417_i15 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), _415_v28)
			_416_v29 = (_416_v29).Update(_dafny.Companion_Sequence_.Concatenate(p1, p1), _415_v28)
			var _418_v30 *C0
			_ = _418_v30
			var _nw74 *C0 = New_C0_()
			_ = _nw74
			_nw74.Ctor__()
			_418_v30 = _nw74
			var _419_v31 _dafny.Map
			_ = _419_v31
			_419_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_418_v30, (_dafny.Zero).Minus(p2))
			_419_v31 = (_419_v31).Update(_418_v30, _this.F18)
			var _420_v32 D0
			_ = _420_v32
			_420_v32 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((p1).Cardinality()), p2, p2, (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_this.F18), p2)), p0)
			var _source3 D0 = _420_v32
			_ = _source3
			if _source3.Is_DC1() {
				var _421___mcc_h0 _dafny.Int = _source3.Get_().(D0_DC1).Cf1
				_ = _421___mcc_h0
				var _422___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC1).Cf2
				_ = _422___mcc_h1
				var _423___mcc_h2 _dafny.Int = _source3.Get_().(D0_DC1).Cf3
				_ = _423___mcc_h2
				var _424___mcc_h3 _dafny.Int = _source3.Get_().(D0_DC1).Cf4
				_ = _424___mcc_h3
				var _425___mcc_h4 _dafny.Array = _source3.Get_().(D0_DC1).Cf5
				_ = _425___mcc_h4
				var _426_cf5 _dafny.Array = _425___mcc_h4
				_ = _426_cf5
				var _427_cf4 _dafny.Int = _424___mcc_h3
				_ = _427_cf4
				var _428_cf3 _dafny.Int = _423___mcc_h2
				_ = _428_cf3
				var _429_cf2 _dafny.Int = _422___mcc_h1
				_ = _429_cf2
				var _430_cf1 _dafny.Int = _421___mcc_h0
				_ = _430_cf1
				(globalState).F12 = (Companion_Default___.SafeModuloInt(_427_cf4, _this.F18)).Cmp(_428_cf3) <= 0
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_426_cf5), 0))
				_ = _index71
				(_426_cf5).ArraySet1(_364_v1, (_index71).Int())
				var _431_v33 _dafny.MultiSet
				_ = _431_v33
				_431_v33 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-308))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}((func(_432_cf1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_433_i16 _dafny.Int) _dafny.Int {
						return _432_cf1
					}
				})(_430_cf1)))))
				var _434_v34 _dafny.MultiSet
				_ = _434_v34
				_434_v34 = _dafny.MultiSetOf(_dafny.IntOfInt64(853), _this.F18)
				var _435_v35 _dafny.Sequence
				_ = _435_v35
				_435_v35 = _dafny.SeqOf(_364_v1, _364_v1)
				var _436_v36 _dafny.Map
				_ = _436_v36
				_436_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_415_v28, false)
				var _437_v37 _dafny.MultiSet
				_ = _437_v37
				_437_v37 = _dafny.MultiSetOf((func() bool {
					if (_436_v36).Contains(_415_v28) {
						return (_436_v36).Get(_415_v28).(bool)
					}
					return _364_v1
				})())
				var _438_v38 _dafny.Map
				_ = _438_v38
				_438_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_v1, _this.F18)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_426_cf5), 0))
				_ = _index72
				var _rhs71 bool = (_431_v33).Equals(((_431_v33).Update((_434_v34).Update(p2, Companion_Default___.Abs((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_364_v1), _364_v1)).Cardinality())), Companion_Default___.Abs(p2))).Union(_431_v33))
				_ = _rhs71
				var _rhs72 bool = (_437_v37).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_435_v35, _435_v35)))
				_ = _rhs72
				var _rhs73 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_418_v30).Fm7(_364_v1, _this.F18, _438_v38, p1, globalState), p1)
				_ = _rhs73
				var _rhs74 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_427_cf4, _430_cf1))
				_ = _rhs74
				var _lhs55 _dafny.Array = _426_cf5
				_ = _lhs55
				var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_426_cf5), 0))
				_ = _lhs56
				var _lhs57 *GlobalState = globalState
				_ = _lhs57
				var _lhs58 *GlobalState = globalState
				_ = _lhs58
				(_lhs55).ArraySet1(_rhs71, (_lhs56).Int())
				_364_v1 = _rhs72
				_lhs57.F3 = _rhs73
				_lhs58.F10 = _rhs74
				(globalState).F3 = p1
				var _439_v39 D0
				_ = _439_v39
				_439_v39 = Companion_D0_.Create_DC0_((_438_v38).Cardinality())
				var _440_v40 _dafny.Sequence
				_ = _440_v40
				_440_v40 = _dafny.SeqOf(_439_v39)
				_440_v40 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm10(_364_v1, _415_v28, globalState), _440_v40)
			} else if _source3.Is_DC2() {
				var _441___mcc_h5 bool = _source3.Get_().(D0_DC2).Cf6
				_ = _441___mcc_h5
				var _442_cf6 bool = _441___mcc_h5
				_ = _442_cf6
				var _rhs75 bool = false
				_ = _rhs75
				var _rhs76 _dafny.Int = (Companion_Default___.Fm0(globalState)).Minus(_this.F18)
				_ = _rhs76
				var _rhs77 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfUint32((p1).Cardinality())), _this.F18)
				_ = _rhs77
				var _lhs59 *GlobalState = globalState
				_ = _lhs59
				var _lhs60 *GlobalState = globalState
				_ = _lhs60
				_364_v1 = _rhs75
				_lhs59.F11 = _rhs76
				_lhs60.F10 = _rhs77
				var _443_v41 _dafny.Map
				_ = _443_v41
				_443_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_442_cf6, _442_cf6)
				_443_v41 = ((_443_v41).Update(_364_v1, _442_cf6)).Merge(_443_v41)
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((p0), 0))
				_ = _index73
				(p0).ArraySet1(_442_cf6, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((p0), 0))
				_ = _index74
				(p0).ArraySet1(false, (_index74).Int())
				var _444_v42 _dafny.Sequence
				_ = _444_v42
				_444_v42 = _dafny.SeqOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((p0), 0))).Int()).(bool), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((p0), 0))).Int()).(bool), _442_cf6)
				_444_v42 = _444_v42
			} else if _source3.Is_DC3() {
				var _445___mcc_h6 _dafny.Int = _source3.Get_().(D0_DC3).Cf7
				_ = _445___mcc_h6
				var _446___mcc_h7 bool = _source3.Get_().(D0_DC3).Cf8
				_ = _446___mcc_h7
				var _447___mcc_h8 _dafny.Int = _source3.Get_().(D0_DC3).Cf9
				_ = _447___mcc_h8
				var _448___mcc_h9 _dafny.Int = _source3.Get_().(D0_DC3).Cf10
				_ = _448___mcc_h9
				var _449_cf10 _dafny.Int = _448___mcc_h9
				_ = _449_cf10
				var _450_cf9 _dafny.Int = _447___mcc_h8
				_ = _450_cf9
				var _451_cf8 bool = _446___mcc_h7
				_ = _451_cf8
				var _452_cf7 _dafny.Int = _445___mcc_h6
				_ = _452_cf7
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((p0), 0))
				_ = _index75
				(p0).ArraySet1(_364_v1, (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((p0), 0))
				_ = _index76
				(p0).ArraySet1(!(_364_v1) || (_364_v1), (_index76).Int())
				_451_cf8 = (_dafny.MultiSetFromSeq(Companion_Default___.Fm4(globalState))).IsDisjointFrom(_dafny.MultiSetOf(_451_cf8))
				var _453_v43 D4
				_ = _453_v43
				_453_v43 = Companion_D4_.Create_DC9_(_418_v30)
				var _454_v44 _dafny.Map
				_ = _454_v44
				_454_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_451_cf8, _451_cf8, _this.F18, globalState), _418_v30)
				var _455_v45 _dafny.Array
				_ = _455_v45
				var _nwElement0_12 *C0 = _418_v30
				_ = _nwElement0_12
				var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(17))
				_ = _nw75
				(_nw75).ArraySet1(_nwElement0_12, 0)
				(_nw75).ArraySet1(_418_v30, 1)
				(_nw75).ArraySet1((_453_v43).Dtor_cf21(), 2)
				(_nw75).ArraySet1(_418_v30, 3)
				(_nw75).ArraySet1(_418_v30, 4)
				(_nw75).ArraySet1(_418_v30, 5)
				(_nw75).ArraySet1(_418_v30, 6)
				(_nw75).ArraySet1(_418_v30, 7)
				(_nw75).ArraySet1(_418_v30, 8)
				(_nw75).ArraySet1(_418_v30, 9)
				(_nw75).ArraySet1(_418_v30, 10)
				(_nw75).ArraySet1(_418_v30, 11)
				(_nw75).ArraySet1((func() *C0 {
					if _364_v1 {
						return _418_v30
					}
					return _418_v30
				})(), 12)
				(_nw75).ArraySet1(_418_v30, 13)
				(_nw75).ArraySet1(_418_v30, 14)
				(_nw75).ArraySet1(_418_v30, 15)
				(_nw75).ArraySet1((func() *C0 {
					if (_454_v44).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
						return (_454_v44).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((p0), 0))).Int()).(bool)).(*C0)
					}
					return _418_v30
				})(), 16)
				_455_v45 = _nw75
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_455_v45), 0))
				_ = _index77
				(_455_v45).ArraySet1(_418_v30, (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((p0), 0))
				_ = _index78
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_455_v45), 0))
				_ = _index79
				var _rhs78 bool = (_451_cf8) || (_451_cf8)
				_ = _rhs78
				var _rhs79 *C0 = _418_v30
				_ = _rhs79
				var _lhs61 _dafny.Array = p0
				_ = _lhs61
				var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((p0), 0))
				_ = _lhs62
				var _lhs63 _dafny.Array = _455_v45
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_455_v45), 0))
				_ = _lhs64
				(_lhs61).ArraySet1(_rhs78, (_lhs62).Int())
				(_lhs63).ArraySet1(_rhs79, (_lhs64).Int())
				(globalState).F10 = _452_cf7
			} else {
				var _456___mcc_h10 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
				_ = _456___mcc_h10
				var _457_cf0 _dafny.Int = _456___mcc_h10
				_ = _457_cf0
				var _458_v46 _dafny.Map
				_ = _458_v46
				_458_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18, _364_v1)
				var _459_v48 _dafny.Map
				_ = _459_v48
				_459_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(305), func() _dafny.Map {
					var _coll16 = _dafny.NewMapBuilder()
					_ = _coll16
					for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(306), _dafny.IntOfInt64(513))); ; {
						_compr_16, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _460_v47 _dafny.Int
						_460_v47 = interface{}(_compr_16).(_dafny.Int)
						if ((_dafny.IntOfInt64(306)).Cmp(_460_v47) <= 0) && ((_460_v47).Cmp(_dafny.IntOfInt64(513)) < 0) {
							_coll16.Add((_460_v47).Minus(_457_cf0), _364_v1)
						}
					}
					return _coll16.ToMap()
				}())
				var _461_v49 D0
				_ = _461_v49
				_461_v49 = Companion_D0_.Create_DC3_(p2, _364_v1, _this.F18, _this.F18)
				var _462_v50 _dafny.Map
				_ = _462_v50
				_462_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_458_v46).Merge((func() _dafny.Map {
					if (_459_v48).Contains(_457_cf0) {
						return (_459_v48).Get(_457_cf0).(_dafny.Map)
					}
					return _458_v46
				})()), _461_v49)
				_462_v50 = (_462_v50).Update(_458_v46, _461_v49)
				var _463_v51 _dafny.Map
				_ = _463_v51
				_463_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_v1, _this.F18)
				var _464_v52 _dafny.Sequence
				_ = _464_v52
				_464_v52 = p1
				(globalState).F3 = (_418_v30).Fm7(_364_v1, _this.F18, (func() _dafny.Map {
					if _364_v1 {
						return _463_v51
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_364_v1, true, _457_cf0, globalState), _dafny.IntOfInt64(175))
				})(), (_464_v52), globalState)
				(globalState).F12 = _364_v1
				(_this).F18 = (_dafny.IntOfInt64(252)).Minus(p2)
			}
			(globalState).F12 = false
		} else {
			var _465_v53 D1
			_ = _465_v53
			_465_v53 = Companion_D1_.Create_DC5_(_364_v1, _this.F18)
			var _source4 D1 = _465_v53
			_ = _source4
			if _source4.Is_DC5() {
				var _466___mcc_h11 bool = _source4.Get_().(D1_DC5).Cf12
				_ = _466___mcc_h11
				var _467___mcc_h12 _dafny.Int = _source4.Get_().(D1_DC5).Cf13
				_ = _467___mcc_h12
				var _468_cf13 _dafny.Int = _467___mcc_h12
				_ = _468_cf13
				var _469_cf12 bool = _466___mcc_h11
				_ = _469_cf12
				var _470_v54 _dafny.Map
				_ = _470_v54
				_470_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(528), _469_cf12)
				var _471_v55 _dafny.Sequence
				_ = _471_v55
				_471_v55 = _dafny.SeqOf(_469_cf12, _469_cf12)
				var _472_v56 _dafny.CodePoint
				_ = _472_v56
				_472_v56 = _dafny.CodePoint('a')
				var _473_v57 _dafny.Map
				_ = _473_v57
				_473_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains(_471_v55, (func() bool {
					if (_470_v54).Contains(_468_cf13) {
						return (_470_v54).Get(_468_cf13).(bool)
					}
					return _469_cf12
				})()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _472_v56))
				_473_v57 = (_473_v57).Merge(_473_v57)
				var _474_v58 _dafny.Array
				_ = _474_v58
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
				_ = _nw76
				_474_v58 = _nw76
				var _475_v59 _dafny.Array
				_ = _475_v59
				_475_v59 = _474_v58
				var _rhs80 _dafny.Array = (_475_v59)
				_ = _rhs80
				var _rhs81 _dafny.Int = p2
				_ = _rhs81
				var _lhs65 *C1 = _this
				_ = _lhs65
				_474_v58 = _rhs80
				_lhs65.F18 = _rhs81
				var _476_v60 _dafny.Map
				_ = _476_v60
				_476_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18, ((_dafny.Zero).Minus((Companion_Default___.Fm11(_364_v1, globalState)).Cardinality())).Plus(p2))
				_476_v60 = (_476_v60).Update(_dafny.IntOfInt64(-983), p2)
				(globalState).F10 = _dafny.IntOfInt64(-458)
			} else {
				var _477___mcc_h13 _dafny.CodePoint = _source4.Get_().(D1_DC4).Cf11
				_ = _477___mcc_h13
				var _478_cf11 _dafny.CodePoint = _477___mcc_h13
				_ = _478_cf11
				var _479_v61 *C0
				_ = _479_v61
				var _nw77 *C0 = New_C0_()
				_ = _nw77
				_nw77.Ctor__()
				_479_v61 = _nw77
				_479_v61 = _479_v61
				_364_v1 = _364_v1
				_364_v1 = _364_v1
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((p0), 0))
				_ = _index80
				(p0).ArraySet1((_this.F18).Cmp(_this.F18) >= 0, (_index80).Int())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((p0), 0))
				_ = _index81
				(p0).ArraySet1(_364_v1, (_index81).Int())
			}
			var _480_v63 _dafny.Set
			_ = _480_v63
			_480_v63 = _dafny.SetOf(_364_v1)
			var _481_v64 *C0
			_ = _481_v64
			var _nw78 *C0 = New_C0_()
			_ = _nw78
			_nw78.Ctor__()
			_481_v64 = _nw78
			var _482_v65 _dafny.Sequence
			_ = _482_v65
			_482_v65 = _dafny.SeqOf(_this.F18, _dafny.IntOfInt64(55))
			var _483_v66 _dafny.Sequence
			_ = _483_v66
			_483_v66 = _dafny.SeqOf((func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(905), _dafny.IntOfInt64(243))); ; {
					_compr_17, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _484_v62 _dafny.Int
					_484_v62 = interface{}(_compr_17).(_dafny.Int)
					if ((_dafny.IntOfInt64(905)).Cmp(_484_v62) <= 0) && ((_484_v62).Cmp(_dafny.IntOfInt64(243)) < 0) {
						_coll17.Add((_484_v62).Times(p2), _364_v1)
					}
				}
				return _coll17.ToMap()
			}()).Cardinality(), (_480_v63).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_481_v64)).Cardinality()), _dafny.IntOfUint32((_482_v65).Cardinality()))
			(globalState).F12 = Companion_Default___.Fm1(_364_v1, false, _dafny.IntOfUint32((_483_v66).Cardinality()), globalState)
			(globalState).F12 = false
			var _485_v67 _dafny.Array
			_ = _485_v67
			var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
			_ = _nw79
			_485_v67 = _nw79
			var _486_v68 _dafny.Array
			_ = _486_v68
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_11
			var _nw80 _dafny.Array
			_ = _nw80
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw80 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.MultiSet = func(_487_i17 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_this.F18)
				}
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw80 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw80).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw80).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_486_v68 = _nw80
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_485_v67), 0))
			_ = _index82
			(_485_v67).ArraySet1(_486_v68, (_index82).Int())
			var _488_v69 _dafny.MultiSet
			_ = _488_v69
			_488_v69 = _dafny.MultiSetOf(_364_v1)
			var _489_v70 _dafny.MultiSet
			_ = _489_v70
			_489_v70 = _dafny.MultiSetOf(_this.F18, p2, (_dafny.Zero).Minus(p2), (_488_v69).Cardinality())
			var _490_v71 _dafny.Sequence
			_ = _490_v71
			_490_v71 = _dafny.SeqOf(_489_v70)
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(589), _dafny.ArrayLen((_485_v67), 0))
			_ = _index83
			var _nwElement0_13 _dafny.MultiSet = (_490_v71).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_490_v71).Cardinality()))).Uint32()).(_dafny.MultiSet)
			_ = _nwElement0_13
			var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(15))
			_ = _nw81
			(_nw81).ArraySet1(_nwElement0_13, 0)
			(_nw81).ArraySet1(((_481_v64).Fm6(_dafny.IntOfUint32((_483_v66).Cardinality()), _364_v1, globalState)).Difference(_489_v70), 1)
			(_nw81).ArraySet1(_489_v70, 2)
			(_nw81).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_483_v66, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_483_v66).Cardinality()))).Uint32(), _dafny.IntOfInt64(706)), _483_v66)), 3)
			(_nw81).ArraySet1(_489_v70, 4)
			(_nw81).ArraySet1(_489_v70, 5)
			(_nw81).ArraySet1((_489_v70).Intersection(_489_v70), 6)
			(_nw81).ArraySet1(_489_v70, 7)
			(_nw81).ArraySet1((_489_v70).Update(p2, Companion_Default___.Abs(p2)), 8)
			(_nw81).ArraySet1(_489_v70, 9)
			(_nw81).ArraySet1(_489_v70, 10)
			(_nw81).ArraySet1((_489_v70).Difference(_489_v70), 11)
			(_nw81).ArraySet1(_489_v70, 12)
			(_nw81).ArraySet1(_dafny.MultiSetOf(p2, _this.F18), 13)
			(_nw81).ArraySet1((_489_v70).Union((_dafny.MultiSetOf(p2, _this.F18)).Update(p2, Companion_Default___.Abs(_this.F18))), 14)
			(_485_v67).ArraySet1(_nw81, (_index83).Int())
			var _491_v72 _dafny.Sequence
			_ = _491_v72
			_491_v72 = _dafny.SeqOf(p0, p0)
			_491_v72 = _491_v72
		}
		if (_364_v1) || (!(_364_v1) || (_364_v1)) {
			var _492_v73 _dafny.Set
			_ = _492_v73
			_492_v73 = _dafny.SetOf(true)
			var _493_v74 _dafny.CodePoint
			_ = _493_v74
			_493_v74 = _dafny.CodePoint('i')
			var _494_v75 _dafny.Sequence
			_ = _494_v75
			_494_v75 = _dafny.SeqOf(_492_v73, Companion_Default___.Fm12(_493_v74, _this.F18, globalState))
			var _495_v76 _dafny.Map
			_ = _495_v76
			_495_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_494_v75).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_494_v75).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _363_v0)
			var _496_v77 D6
			_ = _496_v77
			_496_v77 = Companion_D6_.Create_DC13_(_495_v76)
			var _source5 D0 = Companion_D0_.Create_DC3_(p2, Companion_Default___.Fm1(_364_v1, _364_v1, ((_496_v77).Dtor_cf27()).Cardinality(), globalState), p2, p2)
			_ = _source5
			if _source5.Is_DC1() {
				var _497___mcc_h14 _dafny.Int = _source5.Get_().(D0_DC1).Cf1
				_ = _497___mcc_h14
				var _498___mcc_h15 _dafny.Int = _source5.Get_().(D0_DC1).Cf2
				_ = _498___mcc_h15
				var _499___mcc_h16 _dafny.Int = _source5.Get_().(D0_DC1).Cf3
				_ = _499___mcc_h16
				var _500___mcc_h17 _dafny.Int = _source5.Get_().(D0_DC1).Cf4
				_ = _500___mcc_h17
				var _501___mcc_h18 _dafny.Array = _source5.Get_().(D0_DC1).Cf5
				_ = _501___mcc_h18
				var _502_cf5 _dafny.Array = _501___mcc_h18
				_ = _502_cf5
				var _503_cf4 _dafny.Int = _500___mcc_h17
				_ = _503_cf4
				var _504_cf3 _dafny.Int = _499___mcc_h16
				_ = _504_cf3
				var _505_cf2 _dafny.Int = _498___mcc_h15
				_ = _505_cf2
				var _506_cf1 _dafny.Int = _497___mcc_h14
				_ = _506_cf1
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_363_v0), 0))
				_ = _index84
				(_363_v0).ArraySet1(p2, (_index84).Int())
				var _507_v78 _dafny.MultiSet
				_ = _507_v78
				_507_v78 = _dafny.MultiSetOf(_364_v1, _364_v1)
				var _508_v79 _dafny.Sequence
				_ = _508_v79
				_508_v79 = _dafny.SeqOf(_364_v1, _364_v1, _364_v1, _364_v1, _364_v1)
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_363_v0), 0))
				_ = _index85
				(_363_v0).ArraySet1(((_505_cf2).Minus((func() _dafny.Int {
					if (_507_v78).Contains(_364_v1) {
						return (_507_v78).Multiplicity(_364_v1)
					}
					return p2
				})())).Times((Companion_Default___.Fm13(_364_v1, _508_v79, globalState)).Cardinality()), (_index85).Int())
				var _509_v80 *C0
				_ = _509_v80
				var _nw82 *C0 = New_C0_()
				_ = _nw82
				_nw82.Ctor__()
				_509_v80 = _nw82
				var _510_v81 *C0
				_ = _510_v81
				var _nw83 *C0 = New_C0_()
				_ = _nw83
				_nw83.Ctor__()
				_510_v81 = _nw83
				var _511_v82 _dafny.Map
				_ = _511_v82
				_511_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_509_v80, _505_cf2)
				var _512_v83 D3
				_ = _512_v83
				_512_v83 = Companion_D3_.Create_DC7_(_511_v82)
				var _pat_let_tv5 = _511_v82
				_ = _pat_let_tv5
				_512_v83 = func(_pat_let7_0 D3) D3 {
					return func(_513_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let8_0 _dafny.Map) D3 {
							return func(_514_dt__update_hcf15_h0 _dafny.Map) D3 {
								return Companion_D3_.Create_DC7_(_514_dt__update_hcf15_h0)
							}(_pat_let8_0)
						}(_pat_let_tv5)
					}(_pat_let7_0)
				}((func() D3 {
					if _364_v1 {
						return _512_v83
					}
					return _512_v83
				})())
			} else if _source5.Is_DC2() {
				var _515___mcc_h19 bool = _source5.Get_().(D0_DC2).Cf6
				_ = _515___mcc_h19
				var _516_cf6 bool = _515___mcc_h19
				_ = _516_cf6
				var _517_v84 *C0
				_ = _517_v84
				var _nw84 *C0 = New_C0_()
				_ = _nw84
				_nw84.Ctor__()
				_517_v84 = _nw84
				var _518_v85 _dafny.Map
				_ = _518_v85
				_518_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_516_cf6), _517_v84)
				_518_v85 = (_518_v85).Update(_364_v1, _517_v84)
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((p0), 0))
				_ = _index86
				(p0).ArraySet1(_516_cf6, (_index86).Int())
				var _519_v86 _dafny.Map
				_ = _519_v86
				_519_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_516_cf6, _364_v1)
				var _520_v87 _dafny.Sequence
				_ = _520_v87
				_520_v87 = _dafny.SeqOf(p2, (_519_v86).Cardinality())
				var _521_v88 _dafny.MultiSet
				_ = _521_v88
				_521_v88 = _dafny.MultiSetOf(p2)
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((p0), 0))
				_ = _index87
				(p0).ArraySet1(((_dafny.MultiSetFromSeq(_520_v87)).Intersection(_521_v88)).IsProperSubsetOf(_521_v88), (_index87).Int())
				var _522_v89 _dafny.Map
				_ = _522_v89
				_522_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(79))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_523_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_524_i18 _dafny.Int) _dafny.Int {
						return _523_p2
					}
				})(p2))), _this.F18)
				_522_v89 = (_522_v89).Update(_520_v87, (p2).Minus(p2))
				var _525_v90 _dafny.Array
				_ = _525_v90
				var _nwElement0_14 bool = (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((p0), 0))).Int()).(bool)
				_ = _nwElement0_14
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(18))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_14, 0)
				(_nw85).ArraySet1(!(_364_v1), 1)
				(_nw85).ArraySet1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((p0), 0))).Int()).(bool), 2)
				(_nw85).ArraySet1(false, 3)
				(_nw85).ArraySet1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((p0), 0))).Int()).(bool), 4)
				(_nw85).ArraySet1(false, 5)
				(_nw85).ArraySet1(false, 6)
				(_nw85).ArraySet1(_516_cf6, 7)
				(_nw85).ArraySet1(_516_cf6, 8)
				(_nw85).ArraySet1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((p0), 0))).Int()).(bool), 9)
				(_nw85).ArraySet1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((p0), 0))).Int()).(bool), 10)
				(_nw85).ArraySet1(_516_cf6, 11)
				(_nw85).ArraySet1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.ArrayLen((p0), 0))).Int()).(bool), 12)
				(_nw85).ArraySet1(_516_cf6, 13)
				(_nw85).ArraySet1(_364_v1, 14)
				(_nw85).ArraySet1(_516_cf6, 15)
				(_nw85).ArraySet1(_516_cf6, 16)
				(_nw85).ArraySet1(true, 17)
				_525_v90 = _nw85
				var _526_v91 _dafny.MultiSet
				_ = _526_v91
				_526_v91 = _dafny.MultiSetOf(_525_v90)
				(_this).F18 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_520_v87, _dafny.SeqOf(p2, (_526_v91).Cardinality()))).Cardinality())
			} else if _source5.Is_DC3() {
				var _527___mcc_h20 _dafny.Int = _source5.Get_().(D0_DC3).Cf7
				_ = _527___mcc_h20
				var _528___mcc_h21 bool = _source5.Get_().(D0_DC3).Cf8
				_ = _528___mcc_h21
				var _529___mcc_h22 _dafny.Int = _source5.Get_().(D0_DC3).Cf9
				_ = _529___mcc_h22
				var _530___mcc_h23 _dafny.Int = _source5.Get_().(D0_DC3).Cf10
				_ = _530___mcc_h23
				var _531_cf10 _dafny.Int = _530___mcc_h23
				_ = _531_cf10
				var _532_cf9 _dafny.Int = _529___mcc_h22
				_ = _532_cf9
				var _533_cf8 bool = _528___mcc_h21
				_ = _533_cf8
				var _534_cf7 _dafny.Int = _527___mcc_h20
				_ = _534_cf7
				var _535_v92 _dafny.Sequence
				_ = _535_v92
				_535_v92 = _dafny.SeqOf(_364_v1, _364_v1)
				var _536_v93 _dafny.MultiSet
				_ = _536_v93
				_536_v93 = _dafny.MultiSetOf(Companion_Default___.Fm1(_364_v1, _533_cf8, _this.F18, globalState), _364_v1, _364_v1, _364_v1, _364_v1)
				var _537_v94 _dafny.Sequence
				_ = _537_v94
				_537_v94 = _dafny.SeqOf((p2).Plus(_dafny.IntOfUint32((_535_v92).Cardinality())), (func() _dafny.Int {
					if (_536_v93).Contains(_533_cf8) {
						return (_536_v93).Multiplicity(_533_cf8)
					}
					return p2
				})(), _531_cf10)
				_537_v94 = (func() _dafny.Sequence {
					if _364_v1 {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(429))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg31 _dafny.Int) interface{} {
								return coer31(arg31)
							}
						}((func(_538_cf7 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_539_i19 _dafny.Int) _dafny.Int {
								return (_dafny.SetOf(_538_cf7)).Cardinality()
							}
						})(_534_cf7)))
					}
					return _537_v94
				})()
				var _540_v95 _dafny.MultiSet
				_ = _540_v95
				_540_v95 = _dafny.MultiSetOf(_493_v74)
				var _541_v96 _dafny.Map
				_ = _541_v96
				_541_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(globalState), p2), _540_v95)
				_541_v96 = (_541_v96).Update(_532_cf9, _540_v95)
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw86
				_363_v0 = _nw86
				(globalState).F10 = _this.F18
			} else {
				var _542___mcc_h24 _dafny.Int = _source5.Get_().(D0_DC0).Cf0
				_ = _542___mcc_h24
				var _543_cf0 _dafny.Int = _542___mcc_h24
				_ = _543_cf0
				var _544_v97 _dafny.Map
				_ = _544_v97
				_544_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_v1, _364_v1)
				var _rhs82 _dafny.Int = Companion_Default___.Fm0(globalState)
				_ = _rhs82
				var _rhs83 bool = (_543_cf0).Cmp(((_544_v97).Merge(_544_v97)).Cardinality()) >= 0
				_ = _rhs83
				var _rhs84 _dafny.Array = _363_v0
				_ = _rhs84
				var _lhs66 *GlobalState = globalState
				_ = _lhs66
				var _lhs67 *GlobalState = globalState
				_ = _lhs67
				_lhs66.F11 = _rhs82
				_lhs67.F12 = _rhs83
				_363_v0 = _rhs84
				var _545_v98 _dafny.MultiSet
				_ = _545_v98
				_545_v98 = _dafny.MultiSetOf(_364_v1)
				_545_v98 = ((_545_v98).Difference(_545_v98)).Union(Companion_Default___.Fm14(_this.F18, Companion_D1_.Create_DC5_(_364_v1, _543_cf0), globalState))
				(globalState).F12 = _364_v1
				var _546_v99 *C0
				_ = _546_v99
				var _nw87 *C0 = New_C0_()
				_ = _nw87
				_nw87.Ctor__()
				_546_v99 = _nw87
			}
			_363_v0 = _363_v0
			var _547_v100 _dafny.Set
			_ = _547_v100
			_547_v100 = _dafny.SetOf(p2)
			_547_v100 = _547_v100
			_364_v1 = _dafny.Companion_Sequence_.Equal(p1, p1)
			var _548_v101 _dafny.Map
			_ = _548_v101
			_548_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_v1, _this.F18)
			var _549_v102 D0
			_ = _549_v102
			_549_v102 = Companion_D0_.Create_DC3_(_dafny.IntOfInt64(429), _364_v1, (_548_v101).Cardinality(), _dafny.IntOfInt64(366))
			var _550_v103 _dafny.Map
			_ = _550_v103
			_550_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_549_v102).Dtor_cf8())
			_550_v103 = _550_v103
		} else {
			(globalState).F10 = _this.F18
			var _551_v104 _dafny.Array
			_ = _551_v104
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_12
			var _nw88 _dafny.Array
			_ = _nw88
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw88 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.CodePoint = func(_552_i20 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw88 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw88).ArraySet1CodePoint(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw88).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_551_v104 = _nw88
			_551_v104 = _551_v104
			var _553_v105 _dafny.CodePoint
			_ = _553_v105
			_553_v105 = _dafny.CodePoint('l')
			var _554_v106 _dafny.Set
			_ = _554_v106
			_554_v106 = _dafny.SetOf(_553_v105, Companion_Default___.Fm9(_364_v1, !(false), globalState), _553_v105, _553_v105)
			var _555_v107 D0
			_ = _555_v107
			_555_v107 = Companion_D0_.Create_DC2_((p2).Cmp((_554_v106).Cardinality()) >= 0)
			var _source6 D0 = _555_v107
			_ = _source6
			if _source6.Is_DC1() {
				var _556___mcc_h25 _dafny.Int = _source6.Get_().(D0_DC1).Cf1
				_ = _556___mcc_h25
				var _557___mcc_h26 _dafny.Int = _source6.Get_().(D0_DC1).Cf2
				_ = _557___mcc_h26
				var _558___mcc_h27 _dafny.Int = _source6.Get_().(D0_DC1).Cf3
				_ = _558___mcc_h27
				var _559___mcc_h28 _dafny.Int = _source6.Get_().(D0_DC1).Cf4
				_ = _559___mcc_h28
				var _560___mcc_h29 _dafny.Array = _source6.Get_().(D0_DC1).Cf5
				_ = _560___mcc_h29
				var _561_cf5 _dafny.Array = _560___mcc_h29
				_ = _561_cf5
				var _562_cf4 _dafny.Int = _559___mcc_h28
				_ = _562_cf4
				var _563_cf3 _dafny.Int = _558___mcc_h27
				_ = _563_cf3
				var _564_cf2 _dafny.Int = _557___mcc_h26
				_ = _564_cf2
				var _565_cf1 _dafny.Int = _556___mcc_h25
				_ = _565_cf1
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_561_cf5), 0))
				_ = _index88
				(_561_cf5).ArraySet1((Companion_Default___.Fm1(_364_v1, _364_v1, Companion_Default___.Fm0(globalState), globalState)) && (_364_v1), (_index88).Int())
				var _566_v108 _dafny.Sequence
				_ = _566_v108
				_566_v108 = _dafny.SeqOf(_565_cf1)
				var _567_v109 _dafny.Map
				_ = _567_v109
				_567_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_363_v0, _364_v1)
				var _568_v110 _dafny.MultiSet
				_ = _568_v110
				_568_v110 = _dafny.MultiSetOf(_dafny.IntOfInt64(786))
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_561_cf5), 0))
				_ = _index89
				(_561_cf5).ArraySet1(((_dafny.MultiSetFromSeq(_566_v108)).Update(_563_cf3, Companion_Default___.Abs((_567_v109).Cardinality()))).IsDisjointFrom(_568_v110), (_index89).Int())
				var _569_v111 _dafny.Sequence
				_ = _569_v111
				_569_v111 = _dafny.SeqOf((_561_cf5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_561_cf5), 0))).Int()).(bool))
				_569_v111 = _569_v111
				var _570_v112 *C0
				_ = _570_v112
				var _nw89 *C0 = New_C0_()
				_ = _nw89
				_nw89.Ctor__()
				_570_v112 = _nw89
				_364_v1 = (_569_v111).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(globalState), _dafny.IntOfUint32((_569_v111).Cardinality()))).Uint32()).(bool)
			} else if _source6.Is_DC2() {
				var _571___mcc_h30 bool = _source6.Get_().(D0_DC2).Cf6
				_ = _571___mcc_h30
				var _572_cf6 bool = _571___mcc_h30
				_ = _572_cf6
				var _573_v113 D6
				_ = _573_v113
				_573_v113 = Companion_D6_.Create_DC14_(false, _553_v105, _551_v104)
				_573_v113 = _573_v113
				var _574_v114 _dafny.Array
				_ = _574_v114
				var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw90
				_574_v114 = _nw90
				var _575_v115 D4
				_ = _575_v115
				_575_v115 = Companion_D4_.Create_DC10_(_553_v105, _574_v114, _this.F18)
				var _576_v116 _dafny.Set
				_ = _576_v116
				_576_v116 = _dafny.SetOf(_575_v115)
				var _577_v117 _dafny.Map
				_ = _577_v117
				_577_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_575_v115)).IsSubsetOf(_576_v116), _364_v1)
				var _578_v118 _dafny.Array
				_ = _578_v118
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_13
				var _nw91 _dafny.Array
				_ = _nw91
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw91 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) D0 = (func(_579_v107 D0) func(_dafny.Int) D0 {
						return func(_580_i21 _dafny.Int) D0 {
							return _579_v107
						}
					})(_555_v107)
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
				_578_v118 = _nw91
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_578_v118), 0))
				_ = _index90
				(_578_v118).ArraySet1(_555_v107, (_index90).Int())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_578_v118), 0))
				_ = _index91
				var _rhs85 _dafny.Map = _577_v117
				_ = _rhs85
				var _rhs86 _dafny.Sequence = _dafny.SeqOf(_553_v105)
				_ = _rhs86
				var _rhs87 D0 = _555_v107
				_ = _rhs87
				var _rhs88 bool = _572_cf6
				_ = _rhs88
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				var _lhs69 _dafny.Array = _578_v118
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_578_v118), 0))
				_ = _lhs70
				var _lhs71 *GlobalState = globalState
				_ = _lhs71
				_577_v117 = _rhs85
				_lhs68.F3 = _rhs86
				(_lhs69).ArraySet1(_rhs87, (_lhs70).Int())
				_lhs71.F12 = _rhs88
				var _581_v119 *C0
				_ = _581_v119
				var _nw92 *C0 = New_C0_()
				_ = _nw92
				_nw92.Ctor__()
				_581_v119 = _nw92
				var _582_v120 *C0
				_ = _582_v120
				var _nw93 *C0 = New_C0_()
				_ = _nw93
				_nw93.Ctor__()
				_582_v120 = _nw93
			} else if _source6.Is_DC3() {
				var _583___mcc_h31 _dafny.Int = _source6.Get_().(D0_DC3).Cf7
				_ = _583___mcc_h31
				var _584___mcc_h32 bool = _source6.Get_().(D0_DC3).Cf8
				_ = _584___mcc_h32
				var _585___mcc_h33 _dafny.Int = _source6.Get_().(D0_DC3).Cf9
				_ = _585___mcc_h33
				var _586___mcc_h34 _dafny.Int = _source6.Get_().(D0_DC3).Cf10
				_ = _586___mcc_h34
				var _587_cf10 _dafny.Int = _586___mcc_h34
				_ = _587_cf10
				var _588_cf9 _dafny.Int = _585___mcc_h33
				_ = _588_cf9
				var _589_cf8 bool = _584___mcc_h32
				_ = _589_cf8
				var _590_cf7 _dafny.Int = _583___mcc_h31
				_ = _590_cf7
				var _591_v121 *C0
				_ = _591_v121
				var _nw94 *C0 = New_C0_()
				_ = _nw94
				_nw94.Ctor__()
				_591_v121 = _nw94
				var _592_v122 _dafny.Map
				_ = _592_v122
				_592_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_553_v105, _590_cf7)
				(globalState).F11 = ((_592_v122).Merge(_592_v122)).Cardinality()
				var _593_v123 *C0
				_ = _593_v123
				var _nw95 *C0 = New_C0_()
				_ = _nw95
				_nw95.Ctor__()
				_593_v123 = _nw95
				var _594_v124 _dafny.Array
				_ = _594_v124
				var _nw96 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
				_ = _nw96
				_594_v124 = _nw96
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_594_v124), 0))
				_ = _index92
				(_594_v124).ArraySet1(_591_v121, (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_594_v124), 0))
				_ = _index93
				(_594_v124).ArraySet1(_593_v123, (_index93).Int())
			} else {
				var _595___mcc_h35 _dafny.Int = _source6.Get_().(D0_DC0).Cf0
				_ = _595___mcc_h35
				var _596_cf0 _dafny.Int = _595___mcc_h35
				_ = _596_cf0
				var _597_v125 _dafny.Set
				_ = _597_v125
				_597_v125 = _dafny.SetOf(_this.F18)
				(globalState).F12 = (_dafny.SetOf(_596_cf0, _596_cf0)).IsSubsetOf(_597_v125)
				var _598_v126 _dafny.Map
				_ = _598_v126
				_598_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18, p2)
				var _599_v127 _dafny.Sequence
				_ = _599_v127
				_599_v127 = _dafny.SeqOf(_364_v1)
				var _600_v131 *C0
				_ = _600_v131
				var _nw97 *C0 = New_C0_()
				_ = _nw97
				_nw97.Ctor__()
				_600_v131 = _nw97
				var _601_v132 _dafny.Set
				_ = _601_v132
				_601_v132 = _dafny.SetOf(_600_v131)
				var _602_v133 _dafny.Array
				_ = _602_v133
				var _nwElement0_15 _dafny.Map = (_598_v126).Merge(_598_v126)
				_ = _nwElement0_15
				var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(21))
				_ = _nw98
				(_nw98).ArraySet1(_nwElement0_15, 0)
				(_nw98).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18, _this.F18), 1)
				(_nw98).ArraySet1(_598_v126, 2)
				(_nw98).ArraySet1((func() _dafny.Map {
					if (_599_v127).Select((Companion_Default___.SafeIndex(_this.F18, _dafny.IntOfUint32((_599_v127).Cardinality()))).Uint32()).(bool) {
						return _598_v126
					}
					return _598_v126
				})(), 3)
				(_nw98).ArraySet1(_598_v126, 4)
				(_nw98).ArraySet1(_598_v126, 5)
				(_nw98).ArraySet1(func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(241), _dafny.IntOfInt64(378))); ; {
						_compr_18, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _603_v128 _dafny.Int
						_603_v128 = interface{}(_compr_18).(_dafny.Int)
						if ((_dafny.IntOfInt64(241)).Cmp(_603_v128) <= 0) && ((_603_v128).Cmp(_dafny.IntOfInt64(378)) < 0) {
							_coll18.Add(Companion_Default___.SafeModuloInt(_603_v128, (_dafny.Zero).Minus(_dafny.IntOfUint32((_599_v127).Cardinality()))), _this.F18)
						}
					}
					return _coll18.ToMap()
				}(), 6)
				(_nw98).ArraySet1(_598_v126, 7)
				(_nw98).ArraySet1((func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter22 := _dafny.Iterate((_598_v126).Keys().Elements()); ; {
						_compr_19, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _604_v129 _dafny.Int
						_604_v129 = interface{}(_compr_19).(_dafny.Int)
						if (_598_v126).Contains(_604_v129) {
							_coll19.Add((_604_v129).Plus((_dafny.Zero).Minus(p2)), _this.F18)
						}
					}
					return _coll19.ToMap()
				}()).Merge(func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(735), _dafny.IntOfInt64(507))); ; {
						_compr_20, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _605_v130 _dafny.Int
						_605_v130 = interface{}(_compr_20).(_dafny.Int)
						if ((_dafny.IntOfInt64(735)).Cmp(_605_v130) <= 0) && ((_605_v130).Cmp(_dafny.IntOfInt64(507)) < 0) {
							_coll20.Add(Companion_Default___.SafeDivisionInt(_605_v130, _596_cf0), (_dafny.MultiSetOf(_596_cf0)).Cardinality())
						}
					}
					return _coll20.ToMap()
				}()), 8)
				(_nw98).ArraySet1(Companion_Default___.Fm15((_601_v132).Cardinality(), _364_v1, _596_cf0, globalState), 9)
				(_nw98).ArraySet1(_598_v126, 10)
				(_nw98).ArraySet1(_598_v126, 11)
				(_nw98).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_596_cf0, _dafny.IntOfInt64(650)), 12)
				(_nw98).ArraySet1(_598_v126, 13)
				(_nw98).ArraySet1((_598_v126).Merge(_598_v126), 14)
				(_nw98).ArraySet1(_598_v126, 15)
				(_nw98).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F18, _596_cf0), 16)
				(_nw98).ArraySet1(_598_v126, 17)
				(_nw98).ArraySet1(_598_v126, 18)
				(_nw98).ArraySet1(_598_v126, 19)
				(_nw98).ArraySet1(_598_v126, 20)
				_602_v133 = _nw98
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
				_ = _nw99
				_602_v133 = _nw99
				var _606_v134 _dafny.Array
				_ = _606_v134
				var _nw100 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(14))
				_ = _nw100
				_606_v134 = _nw100
				var _607_v135 _dafny.Array
				_ = _607_v135
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
				_ = _nw101
				_607_v135 = _nw101
				var _608_v136 D4
				_ = _608_v136
				_608_v136 = Companion_D4_.Create_DC10_(_553_v105, _607_v135, (_dafny.Zero).Minus(_this.F18))
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_606_v134), 0))
				_ = _index94
				(_606_v134).ArraySet1(_608_v136, (_index94).Int())
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_606_v134), 0))
				_ = _index95
				(_606_v134).ArraySet1(_608_v136, (_index95).Int())
				var _609_v137 _dafny.MultiSet
				_ = _609_v137
				_609_v137 = _dafny.MultiSetOf(p1)
				_598_v126 = (_598_v126).Update((_609_v137).Cardinality(), p2)
			}
			var _610_v138 _dafny.Map
			_ = _610_v138
			_610_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_v1, _364_v1)
			var _611_v139 _dafny.Array
			_ = _611_v139
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_14
			var _nw102 _dafny.Array
			_ = _nw102
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw102 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Set = (func(_612_v1 bool) func(_dafny.Int) _dafny.Set {
					return func(_613_i22 _dafny.Int) _dafny.Set {
						return _dafny.SetOf(_612_v1)
					}
				})(_364_v1)
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
			_611_v139 = _nw102
			var _614_v140 _dafny.Array
			_ = _614_v140
			_614_v140 = _611_v139
			var _source7 _dafny.Array = (func() _dafny.Array {
				if (func() bool {
					if (_610_v138).Contains(_364_v1) {
						return (_610_v138).Get(_364_v1).(bool)
					}
					return _364_v1
				})() {
					return _611_v139
				}
				return _614_v140
			})()
			_ = _source7
			var _615___mcc_h36 _dafny.Array = _source7
			_ = _615___mcc_h36
			var _616_cf26 _dafny.Array = _615___mcc_h36
			_ = _616_cf26
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((p0), 0))
			_ = _index96
			(p0).ArraySet1((_this.F18).Cmp(p2) != 0, (_index96).Int())
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((p0), 0))
			_ = _index97
			(p0).ArraySet1(_364_v1, (_index97).Int())
			var _617_v141 *C0
			_ = _617_v141
			var _nw103 *C0 = New_C0_()
			_ = _nw103
			_nw103.Ctor__()
			_617_v141 = _nw103
			(globalState).F11 = _this.F18
			var _618_v142 _dafny.Sequence
			_ = _618_v142
			_618_v142 = _dafny.SeqOf(_364_v1, _364_v1, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((p0), 0))).Int()).(bool))
			var _619_v143 _dafny.Sequence
			_ = _619_v143
			_619_v143 = _dafny.SeqOf(_551_v104)
			var _620_v144 D6
			_ = _620_v144
			_620_v144 = Companion_D6_.Create_DC14_(false, _553_v105, (_619_v143).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_619_v143).Cardinality()))).Uint32()).(_dafny.Array))
			var _621_v145 _dafny.Map
			_ = _621_v145
			_621_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_618_v142, _618_v142), (func() bool {
				if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
					return true
				}
				return (_620_v144).Dtor_cf28()
			})())
			_621_v145 = (_621_v145).Update(_dafny.SeqOf(_364_v1), !(_364_v1))
			if !(_364_v1) || (!((func() bool {
				if _364_v1 {
					return _364_v1
				}
				return _364_v1
			})())) {
				var _622_v146 *C0
				_ = _622_v146
				var _nw104 *C0 = New_C0_()
				_ = _nw104
				_nw104.Ctor__()
				_622_v146 = _nw104
				var _623_v147 _dafny.Array
				_ = _623_v147
				var _out12 _dafny.Array
				_ = _out12
				_out12 = Companion_Default___.M0(p2, globalState)
				_623_v147 = _out12
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_363_v0), 0))
				_ = _index98
				(_363_v0).ArraySet1((func() _dafny.Int {
					if Companion_Default___.Fm1(!(_364_v1), true, _this.F18, globalState) {
						return _this.F18
					}
					return _this.F18
				})(), (_index98).Int())
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_363_v0), 0))
				_ = _index99
				(_363_v0).ArraySet1((((_dafny.Zero).Minus(_this.F18)).Minus(_this.F18)).Times((_dafny.Zero).Minus((p2).Times(_this.F18))), (_index99).Int())
				var _624_v148 _dafny.Map
				_ = _624_v148
				_624_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_364_v1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cudbn")).Cardinality())))
				(globalState).F3 = (_622_v146).Fm7(_364_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(554))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_625_v105 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_626_i23 _dafny.Int) _dafny.CodePoint {
						return _625_v105
					}
				})(_553_v105)))).Cardinality()), _624_v148, p1, globalState)
				var _rhs89 _dafny.CodePoint = _553_v105
				_ = _rhs89
				var _rhs90 _dafny.Int = _this.F18
				_ = _rhs90
				var _lhs72 *GlobalState = globalState
				_ = _lhs72
				_553_v105 = _rhs89
				_lhs72.F10 = _rhs90
			} else {
				var _627_v149 *C0
				_ = _627_v149
				var _nw105 *C0 = New_C0_()
				_ = _nw105
				_nw105.Ctor__()
				_627_v149 = _nw105
				var _628_v151 _dafny.Array
				_ = _628_v151
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_15
				var _nw106 _dafny.Array
				_ = _nw106
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw106 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.Sequence = (func(_629_p2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_630_i24 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_this.F18, (func() _dafny.Map {
								var _coll21 = _dafny.NewMapBuilder()
								_ = _coll21
								for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(640), _dafny.IntOfInt64(-60))); ; {
									_compr_21, _ok24 := _iter24()
									if !_ok24 {
										break
									}
									var _631_v150 _dafny.Int
									_631_v150 = interface{}(_compr_21).(_dafny.Int)
									if ((_dafny.IntOfInt64(640)).Cmp(_631_v150) <= 0) && ((_631_v150).Cmp(_dafny.IntOfInt64(-60)) < 0) {
										_coll21.Add((_631_v150).Times(_629_p2), _this.F18)
									}
								}
								return _coll21.ToMap()
							}()).Cardinality())
						}
					})(p2)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw106 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw106).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw106).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_628_v151 = _nw106
				var _632_v152 D0
				_ = _632_v152
				_632_v152 = Companion_D0_.Create_DC3_(_this.F18, false, _this.F18, (_554_v106).Cardinality())
				var _633_v153 D4
				_ = _633_v153
				_633_v153 = Companion_D4_.Create_DC10_(Companion_Default___.Fm9(_364_v1, _364_v1, globalState), _628_v151, (_dafny.SetOf(_632_v152)).Cardinality())
				_633_v153 = _633_v153
				var _634_v154 _dafny.Sequence
				_ = _634_v154
				_634_v154 = _dafny.SeqOf(_363_v0, _363_v0)
				_364_v1 = !(Companion_Default___.Fm1(_364_v1, _dafny.Companion_Sequence_.Contains(_634_v154, _363_v0), (p2).Plus(p2), globalState))
				(globalState).F10 = (func() _dafny.Int {
					if _364_v1 {
						return p2
					}
					return (_this.F18).Minus(p2)
				})()
				(globalState).F10 = _this.F18
			}
		}
		var _635_v155 _dafny.Map
		_ = _635_v155
		_635_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(globalState), !(_364_v1))
		var _636_v156 _dafny.Map
		_ = _636_v156
		_636_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Map {
			if _364_v1 {
				return _635_v155
			}
			return _635_v155
		})()).Cardinality(), p2)
		_636_v156 = (_636_v156).Update((_dafny.Zero).Minus(_this.F18), (_dafny.Zero).Minus(_dafny.IntOfUint32((p1).Cardinality())))
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
