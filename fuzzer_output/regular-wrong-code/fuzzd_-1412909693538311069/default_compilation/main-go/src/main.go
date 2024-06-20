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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.CodePoint('j')), _dafny.MultiSetOf((_dafny.SetOf(Companion_D9_.Create_DC24_(!(false), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality())))).Cardinality(), _dafny.IntOfInt64(596)))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hpvvkqoa")).Cardinality()))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(false)).Cardinality())), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(-524)))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(766)))))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) bool {
	return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false), !(false), true, true, false), _dafny.SeqOf(false))).Cardinality())).Cmp(((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false))).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.CodePoint
			_1_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false))).Contains(_1_v0) {
				_coll0.Add(_1_v0, false)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-803), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-538), _dafny.IntOfInt64(-54))).Cardinality()))).Cardinality())) <= 0
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('y')
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_D1_.Create_DC2_(false)).Dtor_cf2()) || (!(!(true))), _dafny.SeqOf(true))
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 bool, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(_dafny.CodePoint('c'))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, !(true), !(true)), _dafny.SeqOf(false)), _dafny.SeqOf(false, true, false))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("c")
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC7_((_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(409)), _dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(176))).Cardinality(), _dafny.IntOfInt64(766)), _dafny.MultiSetOf((_dafny.SetOf(true, false)).Cardinality()))).Intersection(_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-517), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))), _dafny.MultiSetOf(_dafny.IntOfInt64(264)))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), !(false)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), false)))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(Companion_D1_.Create_DC2_(true))
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-115)), _dafny.SeqOf(_dafny.IntOfInt64(879), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-823))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return (Companion_D7_.Create_DC19_(_dafny.CodePoint('n'), false)).Dtor_cf29()
	}))).Cardinality()))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i1 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)
	}))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(373))))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false, false)).Intersection((_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.CodePoint('t')))).Cardinality()), _dafny.CodePoint('f'))).Dtor_cf14())).Cardinality(), (_dafny.MultiSetFromSeq((Companion_D6_.Create_DC16_(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(916), _dafny.IntOfInt64(337))).Cardinality())))).Dtor_cf22())).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(44), _dafny.IntOfInt64(907)))
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) D3 {
	var _source0 D4 = (func() D4 {
		if false {
			return Companion_D4_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(516), _dafny.IntOfInt64(525)))
		}
		return Companion_D4_.Create_DC10_(func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(375), _dafny.IntOfInt64(105))); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _4_v1 _dafny.Int
					_4_v1 = interface{}(_compr_2).(_dafny.Int)
					if ((_dafny.IntOfInt64(375)).Cmp(_4_v1) <= 0) && ((_4_v1).Cmp(_dafny.IntOfInt64(105)) < 0) {
						_coll2.Add((_4_v1).Times((func() _dafny.Map {
							var _coll3 = _dafny.NewMapBuilder()
							_ = _coll3
							for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(916), _dafny.IntOfInt64(690))); ; {
								_compr_3, _ok3 := _iter3()
								if !_ok3 {
									break
								}
								var _6_v2 _dafny.Int
								_6_v2 = interface{}(_compr_3).(_dafny.Int)
								if ((_dafny.IntOfInt64(916)).Cmp(_6_v2) <= 0) && ((_6_v2).Cmp(_dafny.IntOfInt64(690)) < 0) {
									_coll3.Add((_6_v2).Minus(_dafny.IntOfInt64(893)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(90))).Cardinality(), _dafny.CodePoint('c'))).Cardinality())
								}
							}
							return _coll3.ToMap()
						}()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(434))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg3 _dafny.Int) interface{} {
								return coer3(arg3)
							}
						}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('h')
						}))).Cardinality()))
					}
				}
				return _coll2.ToMap()
			}()).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _7_v0 _dafny.Int
				_7_v0 = interface{}(_compr_1).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll4 = _dafny.NewMapBuilder()
					_ = _coll4
					for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(375), _dafny.IntOfInt64(105))); ; {
						_compr_4, _ok4 := _iter4()
						if !_ok4 {
							break
						}
						var _4_v1 _dafny.Int
						_4_v1 = interface{}(_compr_4).(_dafny.Int)
						if ((_dafny.IntOfInt64(375)).Cmp(_4_v1) <= 0) && ((_4_v1).Cmp(_dafny.IntOfInt64(105)) < 0) {
							_coll4.Add((_4_v1).Times((func() _dafny.Map {
								var _coll5 = _dafny.NewMapBuilder()
								_ = _coll5
								for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(916), _dafny.IntOfInt64(690))); ; {
									_compr_5, _ok5 := _iter5()
									if !_ok5 {
										break
									}
									var _6_v2 _dafny.Int
									_6_v2 = interface{}(_compr_5).(_dafny.Int)
									if ((_dafny.IntOfInt64(916)).Cmp(_6_v2) <= 0) && ((_6_v2).Cmp(_dafny.IntOfInt64(690)) < 0) {
										_coll5.Add((_6_v2).Minus(_dafny.IntOfInt64(893)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(90))).Cardinality(), _dafny.CodePoint('c'))).Cardinality())
									}
								}
								return _coll5.ToMap()
							}()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(434))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg4 _dafny.Int) interface{} {
									return coer4(arg4)
								}
							}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('h')
							}))).Cardinality()))
						}
					}
					return _coll4.ToMap()
				}()).Contains(_7_v0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_7_v0, _dafny.IntOfInt64(-34)), _dafny.IntOfInt64(360))
				}
			}
			return _coll1.ToMap()
		}())
	})()
	_ = _source0
	if _source0.Is_DC11() {
		var _8___mcc_h0 _dafny.Int = _source0.Get_().(D4_DC11).Cf13
		_ = _8___mcc_h0
		var _9___mcc_h1 _dafny.CodePoint = _source0.Get_().(D4_DC11).Cf14
		_ = _9___mcc_h1
		var _10_cf14 _dafny.CodePoint = _9___mcc_h1
		_ = _10_cf14
		var _11_cf13 _dafny.Int = _8___mcc_h0
		_ = _11_cf13
		return Companion_D3_.Create_DC9_(Companion_D3_.Create_DC9_(Companion_D3_.Create_DC7_(func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.SetOf(_dafny.MultiSetOf(_11_cf13, _11_cf13, _11_cf13))).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _12_v3 _dafny.MultiSet
				_12_v3 = interface{}(_compr_6).(_dafny.MultiSet)
				if (_dafny.SetOf(_dafny.MultiSetOf(_11_cf13, _11_cf13, _11_cf13))).Contains(_12_v3) {
					_coll6.Add(_12_v3)
				}
			}
			return _coll6.ToSet()
		}())))
	} else if _source0.Is_DC12() {
		var _13___mcc_h2 _dafny.Int = _source0.Get_().(D4_DC12).Cf15
		_ = _13___mcc_h2
		var _14___mcc_h3 _dafny.Sequence = _source0.Get_().(D4_DC12).Cf16
		_ = _14___mcc_h3
		var _15___mcc_h4 bool = _source0.Get_().(D4_DC12).Cf17
		_ = _15___mcc_h4
		var _16_cf17 bool = _15___mcc_h4
		_ = _16_cf17
		var _17_cf16 _dafny.Sequence = _14___mcc_h3
		_ = _17_cf16
		var _18_cf15 _dafny.Int = _13___mcc_h2
		_ = _18_cf15
		return Companion_D3_.Create_DC9_(Companion_D3_.Create_DC8_(_18_cf15))
	} else if _source0.Is_DC10() {
		var _19___mcc_h5 _dafny.Map = _source0.Get_().(D4_DC10).Cf12
		_ = _19___mcc_h5
		var _20_cf12 _dafny.Map = _19___mcc_h5
		_ = _20_cf12
		return Companion_D3_.Create_DC9_(Companion_D3_.Create_DC8_(_dafny.IntOfInt64(923)))
	} else {
		var _21___mcc_h6 D4 = _source0.Get_().(D4_DC13).Cf18
		_ = _21___mcc_h6
		var _22_cf18 D4 = _21___mcc_h6
		_ = _22_cf18
		return Companion_D3_.Create_DC9_(Companion_D3_.Create_DC7_(_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uh")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(573), true)).Cardinality())).Cardinality()), _dafny.MultiSetOf((_dafny.SetOf(!(!(true)))).Cardinality()))))
	}
}
func (_static *CompanionStruct_Default___) Fm20(p0 D2, p1 bool, globalState *GlobalState) _dafny.Set {
	var _source1 D7 = (func() D7 {
		if !(true) {
			return Companion_D7_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), false), func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(910), _dafny.IntOfInt64(290))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _23_v0 _dafny.Int
					_23_v0 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(910)).Cmp(_23_v0) <= 0) && ((_23_v0).Cmp(_dafny.IntOfInt64(290)) < 0) {
						_coll7.Add(Companion_Default___.SafeModuloInt(_23_v0, _dafny.IntOfInt64(50)), _dafny.CodePoint('q'))
					}
				}
				return _coll7.ToMap()
			}()))
		}
		return Companion_D7_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(37), _dafny.CodePoint('y'))))
	})()
	_ = _source1
	if _source1.Is_DC19() {
		var _24___mcc_h0 _dafny.CodePoint = _source1.Get_().(D7_DC19).Cf29
		_ = _24___mcc_h0
		var _25___mcc_h1 bool = _source1.Get_().(D7_DC19).Cf30
		_ = _25___mcc_h1
		var _26_cf30 bool = _25___mcc_h1
		_ = _26_cf30
		var _27_cf29 _dafny.CodePoint = _24___mcc_h0
		_ = _27_cf29
		return (_dafny.SetOf(_26_cf30, false, _26_cf30)).Intersection(_dafny.SetOf(_26_cf30))
	} else {
		var _28___mcc_h2 _dafny.Map = _source1.Get_().(D7_DC18).Cf28
		_ = _28___mcc_h2
		var _29_cf28 _dafny.Map = _28___mcc_h2
		_ = _29_cf28
		return (_dafny.SetOf(true)).Union(_dafny.SetOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pqgqi")).Cardinality())).Times(_dafny.IntOfInt64(-128)))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(719), _dafny.IntOfInt64(317))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _30_v0 _dafny.Int
			_30_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(719)).Cmp(_30_v0) <= 0) && ((_30_v0).Cmp(_dafny.IntOfInt64(317)) < 0) {
				_coll8.Add((_30_v0).Times((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('r'))).Cardinality())), ((Companion_D6_.Create_DC17_(_dafny.UnicodeSeqOfUtf8Bytes("uyc"), false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eda")).Cardinality()), true, false)).Dtor_cf25()).Minus(_dafny.IntOfInt64(47)))
			}
		}
		return _coll8.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC10_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(118))).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xbwtruw")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(304))))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) {
	var _31_v0 bool
	_ = _31_v0
	_31_v0 = true
	(globalState).F3 = _31_v0
	var _32_v1 _dafny.Int
	_ = _32_v1
	_32_v1 = _dafny.IntOfInt64(524)
	var _33_v2 _dafny.Map
	_ = _33_v2
	_33_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.CodePoint('x'), globalState), _32_v1)
	_33_v2 = (_33_v2).Update(_32_v1, ((_dafny.Zero).Minus(_32_v1)).Minus(_32_v1))
	var _34_v3 _dafny.Array
	_ = _34_v3
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(26)
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = (func(_35_v0 bool) func(_dafny.Int) bool {
			return func(_36_i0 _dafny.Int) bool {
				return (Companion_D7_.Create_DC19_(_dafny.CodePoint('o'), _35_v0)).Dtor_cf30()
			}
		})(_31_v0)
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
	_34_v3 = _nw0
	var _37_v4 D8
	_ = _37_v4
	_37_v4 = Companion_D8_.Create_DC21_(_32_v1)
	var _pat_let_tv0 = _31_v0
	_ = _pat_let_tv0
	var _pat_let_tv1 = _31_v0
	_ = _pat_let_tv1
	var _pat_let_tv2 = _31_v0
	_ = _pat_let_tv2
	var _pat_let_tv3 = _31_v0
	_ = _pat_let_tv3
	var _rhs0 bool = (_32_v1).Cmp(_32_v1) == 0
	_ = _rhs0
	var _rhs1 bool = func(_source2 D8) bool {
		if _source2.Is_DC21() {
			var _38___mcc_h0 _dafny.Int = _source2.Get_().(D8_DC21).Cf32
			_ = _38___mcc_h0
			var _39_cf32 _dafny.Int = _38___mcc_h0
			_ = _39_cf32
			return !((_dafny.MultiSetOf(_dafny.IntOfInt64(665), _39_cf32)).IsDisjointFrom(_dafny.MultiSetOf(_39_cf32, (_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_D6_.Create_DC17_(_dafny.UnicodeSeqOfUtf8Bytes("ekoegwtab"), _pat_let_tv0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mkjxih")).Cardinality()), _pat_let_tv1, _pat_let_tv2)).Dtor_cf25())))))
		} else {
			var _40___mcc_h1 _dafny.Sequence = _source2.Get_().(D8_DC20).Cf31
			_ = _40___mcc_h1
			var _41_cf31 _dafny.Sequence = _40___mcc_h1
			_ = _41_cf31
			return _pat_let_tv3
		}
	}(_37_v4)
	_ = _rhs1
	var _rhs2 _dafny.Array = (func() _dafny.Array {
		if _31_v0 {
			return _34_v3
		}
		return _34_v3
	})()
	_ = _rhs2
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	_lhs0.F3 = _rhs0
	_31_v0 = _rhs1
	_34_v3 = _rhs2
	var _42_v5 _dafny.Map
	_ = _42_v5
	_42_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jouwqcwf")).Cardinality()))
	var _43_v6 _dafny.CodePoint
	_ = _43_v6
	_43_v6 = _dafny.CodePoint('x')
	var _44_v7 D4
	_ = _44_v7
	_44_v7 = Companion_D4_.Create_DC11_(_32_v1, _43_v6)
	var _45_v8 _dafny.MultiSet
	_ = _45_v8
	_45_v8 = _dafny.MultiSetOf(_31_v0)
	var _46_v9 *C2
	_ = _46_v9
	var _nw1 *C2 = New_C2_()
	_ = _nw1
	_nw1.Ctor__((_44_v7).Dtor_cf13(), _32_v1, (func() _dafny.Int {
		if (_45_v8).Contains(!(!(_31_v0))) {
			return (_45_v8).Multiplicity(!(!(_31_v0)))
		}
		return _32_v1
	})(), (_33_v2).Cardinality())
	_46_v9 = _nw1
	var _47_v10 _dafny.Map
	_ = _47_v10
	_47_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_42_v5).Merge(_42_v5), _46_v9)
	var _48_v11 _dafny.Sequence
	_ = _48_v11
	_48_v11 = _dafny.UnicodeSeqOfUtf8Bytes("vyfp")
	_47_v10 = (_47_v10).Update(((_42_v5).Update(_31_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_48_v11).Cardinality())))).Merge(_42_v5), _46_v9)
	var _49_v12 _dafny.Set
	_ = _49_v12
	_49_v12 = _dafny.SetOf(_34_v3, _34_v3, _34_v3)
	var _hi0 _dafny.Int = (_46_v9).F9()
	_ = _hi0
	for _50_i1 := (_49_v12).Cardinality(); _50_i1.Cmp(_hi0) < 0; _50_i1 = _50_i1.Plus(_dafny.One) {
		if _31_v0 {
			var _51_v13 T0
			_ = _51_v13
			var _nw2 *C2 = New_C2_()
			_ = _nw2
			_nw2.Ctor__(_46_v9.F10, _32_v1, (_46_v9).F9(), (_46_v9).F9())
			_51_v13 = _nw2
			var _52_v14 _dafny.Sequence
			_ = _52_v14
			_52_v14 = _dafny.SeqOf(_51_v13)
			var _53_v15 _dafny.Map
			_ = _53_v15
			_53_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v0, _52_v14)
			(_46_v9).F10 = (_53_v15).Cardinality()
			(globalState).F3 = !((_31_v0) == (_31_v0)) || (_31_v0)
			(_46_v9).F10 = (_51_v13).F7()
			_34_v3 = (func() _dafny.Array {
				if _31_v0 {
					return _34_v3
				}
				return (func() _dafny.Array {
					if _31_v0 {
						return _34_v3
					}
					return _34_v3
				})()
			})()
			var _54_v16 _dafny.Sequence
			_ = _54_v16
			_54_v16 = _dafny.SeqOf(_32_v1, (_46_v9).F9())
			_48_v11 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm11(_31_v0, (_dafny.Zero).Minus(_50_i1), globalState), (Companion_Default___.SafeIndex((_54_v16).Select((Companion_Default___.SafeIndex(_50_i1, _dafny.IntOfUint32((_54_v16).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm11(_31_v0, (_dafny.Zero).Minus(_50_i1), globalState)).Cardinality()))).Uint32(), _43_v6)
		} else {
			var _55_v17 _dafny.Map
			_ = _55_v17
			_55_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _45_v8)
			(globalState).F1 = !(((func() _dafny.MultiSet {
				if (_55_v17).Contains(!(_31_v0)) {
					return (_55_v17).Get(!(_31_v0)).(_dafny.MultiSet)
				}
				return _45_v8
			})()).IsSubsetOf(Companion_Default___.Fm17(_31_v0, (_46_v9).F9(), globalState)))
			var _56_v18 _dafny.Map
			_ = _56_v18
			_56_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_43_v6, _31_v0)
			var _57_v20 _dafny.Map
			_ = _57_v20
			_57_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v18, func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(791), _dafny.IntOfInt64(145))); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _58_v19 _dafny.Int
					_58_v19 = interface{}(_compr_9).(_dafny.Int)
					if ((_dafny.IntOfInt64(791)).Cmp(_58_v19) <= 0) && ((_58_v19).Cmp(_dafny.IntOfInt64(145)) < 0) {
						_coll9.Add((_58_v19).Plus(_dafny.IntOfInt64(293)), _43_v6)
					}
				}
				return _coll9.ToMap()
			}())
			var _59_v21 *C4
			_ = _59_v21
			var _nw3 *C4 = New_C4_()
			_ = _nw3
			_nw3.Ctor__((_57_v20).Merge(_57_v20))
			_59_v21 = _nw3
			var _60_v22 _dafny.Sequence
			_ = _60_v22
			_60_v22 = _dafny.SeqOf(_31_v0)
			var _61_v23 D4
			_ = _61_v23
			_61_v23 = Companion_D4_.Create_DC12_((_46_v9).F9(), _60_v22, _31_v0)
			var _62_v24 *C1
			_ = _62_v24
			var _nw4 *C1 = New_C1_()
			_ = _nw4
			_nw4.Ctor__(_31_v0, (_dafny.Zero).Minus((_45_v8).Cardinality()), (_46_v9).F9(), _dafny.IntOfUint32(((_61_v23).Dtor_cf16()).Cardinality()))
			_62_v24 = _nw4
			var _63_v25 _dafny.Array
			_ = _63_v25
			var _nwElement0_0 *C1 = _62_v24
			_ = _nwElement0_0
			var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(3))
			_ = _nw5
			(_nw5).ArraySet1(_nwElement0_0, 0)
			(_nw5).ArraySet1(_62_v24, 1)
			(_nw5).ArraySet1(_62_v24, 2)
			_63_v25 = _nw5
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_63_v25), 0))
			_ = _index0
			(_63_v25).ArraySet1(_62_v24, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_63_v25), 0))
			_ = _index1
			(_63_v25).ArraySet1(_62_v24, (_index1).Int())
			_32_v1 = (_dafny.Zero).Minus(_50_i1)
			var _64_v26 _dafny.Array
			_ = _64_v26
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
			_ = _nw6
			_64_v26 = _nw6
			_64_v26 = (func() _dafny.Array {
				if _31_v0 {
					return _64_v26
				}
				return _64_v26
			})()
		}
		var _65_v27 _dafny.Array
		_ = _65_v27
		var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
		_ = _nw7
		_65_v27 = _nw7
		var _66_v28 D4
		_ = _66_v28
		_66_v28 = Companion_D4_.Create_DC13_(Companion_Default___.Fm23(globalState))
		var _67_v29 D4
		_ = _67_v29
		_67_v29 = Companion_D4_.Create_DC13_(_66_v28)
		var _68_v30 _dafny.Sequence
		_ = _68_v30
		_68_v30 = _dafny.SeqOf(_67_v29)
		var _69_v31 _dafny.Map
		_ = _69_v31
		_69_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v27, _68_v30)
		_69_v31 = (_69_v31).Update(_65_v27, _68_v30)
		var _70_v32 _dafny.Sequence
		_ = _70_v32
		_70_v32 = _dafny.SeqOf(_31_v0, true, _31_v0, false)
		(_46_v9).F10 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if _31_v0 {
				return _70_v32
			}
			return _70_v32
		})(), _70_v32)).Cardinality()))
		var _71_v34 _dafny.Map
		_ = _71_v34
		_71_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _33_v2)
		var _72_v35 _dafny.Sequence
		_ = _72_v35
		_72_v35 = _dafny.SeqOf((_46_v9).F9(), _50_i1, ((func() _dafny.Map {
			if (_71_v34).Contains(true) {
				return (_71_v34).Get(true).(_dafny.Map)
			}
			return _33_v2
		})()).Cardinality(), (_46_v9).Fm7(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm11(_31_v0, (_46_v9).F9(), globalState), (Companion_Default___.SafeIndex(_50_i1, _dafny.IntOfUint32((Companion_Default___.Fm11(_31_v0, (_46_v9).F9(), globalState)).Cardinality()))).Uint32(), _43_v6), (_46_v9).F9(), globalState))
		var _73_v36 *C1
		_ = _73_v36
		var _nw8 *C1 = New_C1_()
		_ = _nw8
		_nw8.Ctor__(((func() _dafny.Int {
			if (_42_v5).Contains(_31_v0) {
				return (_42_v5).Get(_31_v0).(_dafny.Int)
			}
			return (_46_v9).F9()
		})()).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_31_v0)).Cardinality()))) <= 0, (func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_72_v35).Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _74_v33 _dafny.Int
				_74_v33 = interface{}(_compr_10).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_72_v35, _74_v33) {
					_coll10.Add((_74_v33).Plus(_46_v9.F10), _46_v9.F10)
				}
			}
			return _coll10.ToMap()
		}()).Cardinality(), (_46_v9.F10).Plus((_46_v9).F9()), _46_v9.F10)
		_73_v36 = _nw8
		_73_v36 = _73_v36
	}
	var _75_v37 _dafny.Set
	_ = _75_v37
	_75_v37 = _dafny.SetOf(_31_v0, _31_v0, _31_v0)
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_34_v3), 0))
	_ = _index2
	(_34_v3).ArraySet1((_75_v37).IsProperSubsetOf(_75_v37), (_index2).Int())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_34_v3), 0))
	_ = _index3
	(_34_v3).ArraySet1(Companion_Default___.Fm2(_31_v0, globalState), (_index3).Int())
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _76_globalState *GlobalState
	_ = _76_globalState
	var _nw9 *GlobalState = New_GlobalState_()
	_ = _nw9
	_nw9.Ctor__(true, true, _dafny.IntOfInt64(135), true, false, _dafny.CodePoint('e'))
	_76_globalState = _nw9
	var _77_v0 _dafny.Sequence
	_ = _77_v0
	_77_v0 = _dafny.UnicodeSeqOfUtf8Bytes("jc")
	_77_v0 = _dafny.UnicodeSeqOfUtf8Bytes("pqitbtm")
	var _78_v1 _dafny.Int
	_ = _78_v1
	_78_v1 = _dafny.IntOfInt64(414)
	var _79_v2 bool
	_ = _79_v2
	_79_v2 = true
	var _80_v3 _dafny.Sequence
	_ = _80_v3
	_80_v3 = _dafny.SeqOf(_79_v2, true)
	var _81_v4 _dafny.CodePoint
	_ = _81_v4
	_81_v4 = _dafny.CodePoint('j')
	var _82_v5 _dafny.Map
	_ = _82_v5
	_82_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _dafny.IntOfUint32((_80_v3).Cardinality()))
	var _83_v6 _dafny.Map
	_ = _83_v6
	_83_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v3, (func() _dafny.Int {
		if (_82_v5).Contains(!(true)) {
			return (_82_v5).Get(!(true)).(_dafny.Int)
		}
		return _78_v1
	})())
	_78_v1 = Companion_Default___.SafeDivisionInt(_78_v1, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v3, Companion_Default___.Fm0(_81_v4, _76_globalState))).Merge(_83_v6)).Cardinality())
	if !(_79_v2) {
		var _84_v7 _dafny.Array
		_ = _84_v7
		var _nwElement0_1 _dafny.Int = _78_v1
		_ = _nwElement0_1
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(2))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_1, 0)
		(_nw10).ArraySet1(_78_v1, 1)
		_84_v7 = _nw10
		var _85_v8 _dafny.Map
		_ = _85_v8
		_85_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _81_v4)
		var _86_v9 _dafny.Map
		_ = _86_v9
		_86_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_85_v8).Update(_79_v2, _81_v4), Companion_Default___.Fm2(_79_v2, _76_globalState))
		var _87_v10 _dafny.Map
		_ = _87_v10
		_87_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1((func() bool {
			if (_86_v9).Contains(_85_v8) {
				return (_86_v9).Get(_85_v8).(bool)
			}
			return _79_v2
		})(), _76_globalState), _84_v7)
		var _88_v11 _dafny.MultiSet
		_ = _88_v11
		_88_v11 = _dafny.MultiSetOf(_82_v5)
		var _89_v12 D0
		_ = _89_v12
		_89_v12 = Companion_D0_.Create_DC0_(_84_v7)
		var _90_v13 _dafny.Array
		_ = _90_v13
		var _nwElement0_2 _dafny.Array = _84_v7
		_ = _nwElement0_2
		var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(28))
		_ = _nw11
		(_nw11).ArraySet1(_nwElement0_2, 0)
		(_nw11).ArraySet1(_84_v7, 1)
		(_nw11).ArraySet1(_84_v7, 2)
		(_nw11).ArraySet1(_84_v7, 3)
		(_nw11).ArraySet1(_84_v7, 4)
		(_nw11).ArraySet1(_84_v7, 5)
		(_nw11).ArraySet1(_84_v7, 6)
		(_nw11).ArraySet1(_84_v7, 7)
		(_nw11).ArraySet1(_84_v7, 8)
		(_nw11).ArraySet1(_84_v7, 9)
		(_nw11).ArraySet1(_84_v7, 10)
		(_nw11).ArraySet1(_84_v7, 11)
		(_nw11).ArraySet1(_84_v7, 12)
		(_nw11).ArraySet1(_84_v7, 13)
		(_nw11).ArraySet1((func() _dafny.Array {
			if (_87_v10).Contains(_88_v11) {
				return (_87_v10).Get(_88_v11).(_dafny.Array)
			}
			return _84_v7
		})(), 14)
		(_nw11).ArraySet1(_84_v7, 15)
		(_nw11).ArraySet1(_84_v7, 16)
		(_nw11).ArraySet1(_84_v7, 17)
		(_nw11).ArraySet1(_84_v7, 18)
		(_nw11).ArraySet1(_84_v7, 19)
		(_nw11).ArraySet1(_84_v7, 20)
		(_nw11).ArraySet1(_84_v7, 21)
		(_nw11).ArraySet1(_84_v7, 22)
		(_nw11).ArraySet1((func() _dafny.Array {
			if _79_v2 {
				return _84_v7
			}
			return (_89_v12).Dtor_cf0()
		})(), 23)
		(_nw11).ArraySet1(_84_v7, 24)
		(_nw11).ArraySet1(_84_v7, 25)
		(_nw11).ArraySet1(_84_v7, 26)
		(_nw11).ArraySet1(_84_v7, 27)
		_90_v13 = _nw11
		var _rhs3 _dafny.Array = _90_v13
		_ = _rhs3
		var _rhs4 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-42))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}((func(_91_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_92_i0 _dafny.Int) _dafny.CodePoint {
				return _91_v4
			}
		})(_81_v4))), _dafny.Companion_Sequence_.Concatenate(_77_v0, _77_v0))
		_ = _rhs4
		_90_v13 = _rhs3
		_77_v0 = _rhs4
		(_76_globalState).F3 = _79_v2
		_79_v2 = _79_v2
		var _93_v14 D0
		_ = _93_v14
		_93_v14 = Companion_D0_.Create_DC1_(_81_v4)
		_93_v14 = _93_v14
		_78_v1 = (_dafny.IntOfInt64(466)).Plus(_dafny.IntOfUint32((_77_v0).Cardinality()))
	} else {
		Companion_Default___.M0(_76_globalState)
		if (_80_v3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if _79_v2 {
				return _78_v1
			}
			return _78_v1
		})(), _dafny.IntOfUint32((_80_v3).Cardinality()))).Uint32()).(bool) {
			var _94_v15 _dafny.Set
			_ = _94_v15
			_94_v15 = _dafny.SetOf(_81_v4, _81_v4, _81_v4, Companion_Default___.Fm3(_79_v2, _76_globalState))
			var _95_v16 _dafny.Sequence
			_ = _95_v16
			_95_v16 = _dafny.SeqOf(_94_v15, _94_v15)
			_95_v16 = _dafny.Companion_Sequence_.Concatenate(_95_v16, _dafny.SeqOf(_94_v15, _94_v15, _94_v15))
			_77_v0 = _77_v0
			var _96_v17 _dafny.Array
			_ = _96_v17
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_1
			var _nw12 _dafny.Array
			_ = _nw12
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw12 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_97_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_98_i1 _dafny.Int) _dafny.Int {
						return (_98_i1).Minus(_97_v1)
					}
				})(_78_v1)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw12 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw12).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw12).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_96_v17 = _nw12
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_96_v17), 0))
			_ = _index4
			(_96_v17).ArraySet1((_78_v1).Times(_78_v1), (_index4).Int())
			var _99_v18 _dafny.Array
			_ = _99_v18
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw13
			_99_v18 = _nw13
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_99_v18), 0))
			_ = _index5
			(_99_v18).ArraySet1(_77_v0, (_index5).Int())
			var _100_v19 _dafny.Set
			_ = _100_v19
			_100_v19 = _dafny.SetOf(_79_v2)
			var _101_v20 _dafny.Array
			_ = _101_v20
			var _nwElement0_3 _dafny.CodePoint = _81_v4
			_ = _nwElement0_3
			var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(8))
			_ = _nw14
			(_nw14).ArraySet1CodePoint(_nwElement0_3, 0)
			(_nw14).ArraySet1CodePoint((_77_v0).Select((Companion_Default___.SafeIndex((_100_v19).Cardinality(), _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
			(_nw14).ArraySet1CodePoint(Companion_Default___.Fm3(_79_v2, _76_globalState), 2)
			(_nw14).ArraySet1CodePoint(_81_v4, 3)
			(_nw14).ArraySet1CodePoint(_81_v4, 4)
			(_nw14).ArraySet1CodePoint(_81_v4, 5)
			(_nw14).ArraySet1CodePoint(_81_v4, 6)
			(_nw14).ArraySet1CodePoint(_81_v4, 7)
			_101_v20 = _nw14
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_101_v20), 0))
			_ = _index6
			(_101_v20).ArraySet1CodePoint(_dafny.CodePoint('t'), (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_96_v17), 0))
			_ = _index7
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_99_v18), 0))
			_ = _index8
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_101_v20), 0))
			_ = _index9
			var _rhs5 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.Fm4((_dafny.Zero).Minus(_78_v1), _79_v2, _76_globalState)).Cardinality())
			_ = _rhs5
			var _rhs6 _dafny.Sequence = _77_v0
			_ = _rhs6
			var _rhs7 _dafny.CodePoint = _81_v4
			_ = _rhs7
			var _rhs8 _dafny.CodePoint = _81_v4
			_ = _rhs8
			var _rhs9 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_81_v4, _76_globalState), Companion_Default___.Fm0(_81_v4, _76_globalState))
			_ = _rhs9
			var _lhs1 _dafny.Array = _96_v17
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_96_v17), 0))
			_ = _lhs2
			var _lhs3 _dafny.Array = _99_v18
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_99_v18), 0))
			_ = _lhs4
			var _lhs5 _dafny.Array = _101_v20
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_101_v20), 0))
			_ = _lhs6
			(_lhs1).ArraySet1(_rhs5, (_lhs2).Int())
			(_lhs3).ArraySet1(_rhs6, (_lhs4).Int())
			_81_v4 = _rhs7
			(_lhs5).ArraySet1CodePoint(_rhs8, (_lhs6).Int())
			_78_v1 = _rhs9
			var _102_v21 D0
			_ = _102_v21
			_102_v21 = Companion_D0_.Create_DC1_(_81_v4)
			_102_v21 = Companion_D0_.Create_DC1_((_101_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_101_v20), 0))).Int()))
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_99_v18), 0))
			_ = _index10
			(_99_v18).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_99_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_99_v18), 0))).Int()).(_dafny.Sequence), (_99_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_99_v18), 0))).Int()).(_dafny.Sequence)), (_index10).Int())
		} else {
			Companion_Default___.M0(_76_globalState)
			var _103_v22 _dafny.Sequence
			_ = _103_v22
			_103_v22 = _dafny.SeqOf(Companion_Default___.Fm0(_81_v4, _76_globalState), _dafny.IntOfInt64(800))
			var _104_v23 *C1
			_ = _104_v23
			var _nw15 *C1 = New_C1_()
			_ = _nw15
			_nw15.Ctor__(_79_v2, _78_v1, _dafny.IntOfUint32((_103_v22).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_77_v0, _77_v0)).Cardinality()))
			_104_v23 = _nw15
			var _105_v24 _dafny.MultiSet
			_ = _105_v24
			_105_v24 = _dafny.MultiSetOf(_79_v2, (_104_v23).F12())
			var _106_v25 _dafny.Map
			_ = _106_v25
			_106_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_105_v24).Update((_104_v23).F12(), Companion_Default___.Abs(_78_v1))).Cardinality(), _dafny.Companion_Sequence_.Update(_77_v0, (Companion_Default___.SafeIndex(_104_v23.F13, _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32(), _dafny.CodePoint('q')))
			var _107_v26 T0
			_ = _107_v26
			var _nw16 *C2 = New_C2_()
			_ = _nw16
			_nw16.Ctor__((_106_v25).Cardinality(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(852), _78_v1), (func() _dafny.Int {
				if (_104_v23).F12() {
					return _104_v23.F13
				}
				return _104_v23.F13
			})(), (_103_v22).Select((Companion_Default___.SafeIndex(_78_v1, _dafny.IntOfUint32((_103_v22).Cardinality()))).Uint32()).(_dafny.Int))
			_107_v26 = _nw16
			_107_v26 = _107_v26
			_81_v4 = _81_v4
			(_76_globalState).F3 = (_dafny.IntOfInt64(105)).Cmp((_107_v26).F7()) != 0
		}
		var _108_v28 D1
		_ = _108_v28
		_108_v28 = Companion_D1_.Create_DC3_(_81_v4)
		var _109_v29 _dafny.Map
		_ = _109_v29
		_109_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm13(_78_v1, _76_globalState), _108_v28)
		var _110_v30 *C4
		_ = _110_v30
		var _nw17 *C4 = New_C4_()
		_ = _nw17
		_nw17.Ctor__(func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((_109_v29).Keys().Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _111_v27 _dafny.Map
				_111_v27 = interface{}(_compr_11).(_dafny.Map)
				if (_109_v29).Contains(_111_v27) {
					_coll11.Add(_111_v27, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _dafny.CodePoint('y')))
				}
			}
			return _coll11.ToMap()
		}())
		_110_v30 = _nw17
		_110_v30 = _110_v30
		var _112_v31 *C0
		_ = _112_v31
		var _nw18 *C0 = New_C0_()
		_ = _nw18
		_nw18.Ctor__(_79_v2)
		_112_v31 = _nw18
		var _113_v32 _dafny.Array
		_ = _113_v32
		var _nwElement0_4 *C0 = _112_v31
		_ = _nwElement0_4
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(12))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_4, 0)
		(_nw19).ArraySet1(_112_v31, 1)
		(_nw19).ArraySet1(_112_v31, 2)
		(_nw19).ArraySet1(_112_v31, 3)
		(_nw19).ArraySet1(_112_v31, 4)
		(_nw19).ArraySet1(_112_v31, 5)
		(_nw19).ArraySet1(_112_v31, 6)
		(_nw19).ArraySet1(_112_v31, 7)
		(_nw19).ArraySet1(_112_v31, 8)
		(_nw19).ArraySet1(_112_v31, 9)
		(_nw19).ArraySet1(_112_v31, 10)
		(_nw19).ArraySet1(_112_v31, 11)
		_113_v32 = _nw19
		_113_v32 = _113_v32
		_78_v1 = Companion_Default___.SafeModuloInt(_78_v1, (_78_v1).Times((_dafny.Zero).Minus((Companion_Default___.Fm21(_dafny.IntOfUint32((_80_v3).Cardinality()), _78_v1, _76_globalState)).Cardinality())))
	}
	_78_v1 = _dafny.IntOfInt64(-958)
	var _114_v34 _dafny.Map
	_ = _114_v34
	_114_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_77_v0).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _115_v33 _dafny.CodePoint
			_115_v33 = interface{}(_compr_12).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_77_v0, _115_v33) {
				_coll12.Add(_115_v33, _79_v2)
			}
		}
		return _coll12.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _dafny.CodePoint('q')))
	var _116_v35 D7
	_ = _116_v35
	_116_v35 = Companion_D7_.Create_DC18_(_114_v34)
	var _117_v36 *C4
	_ = _117_v36
	var _nw20 *C4 = New_C4_()
	_ = _nw20
	_nw20.Ctor__((Companion_D7_.Create_DC18_((_116_v35).Dtor_cf28())).Dtor_cf28())
	_117_v36 = _nw20
	if _79_v2 {
		if (_79_v2) || (_79_v2) {
			var _118_v37 *C0
			_ = _118_v37
			var _nw21 *C0 = New_C0_()
			_ = _nw21
			_nw21.Ctor__(_79_v2)
			_118_v37 = _nw21
			var _119_v38 _dafny.Sequence
			_ = _119_v38
			_119_v38 = _dafny.SeqOf(_118_v37)
			var _120_v39 _dafny.Map
			_ = _120_v39
			_120_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_118_v37).F11(), _119_v38)
			_119_v38 = _dafny.Companion_Sequence_.Concatenate(_119_v38, (func() _dafny.Sequence {
				if (_120_v39).Contains((_118_v37).F11()) {
					return (_120_v39).Get((_118_v37).F11()).(_dafny.Sequence)
				}
				return _119_v38
			})())
			var _121_v40 _dafny.Array
			_ = _121_v40
			var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
			_ = _nw22
			_121_v40 = _nw22
			var _122_v41 _dafny.Map
			_ = _122_v41
			_122_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_79_v2, _76_globalState), !(false))
			var _123_v42 _dafny.Sequence
			_ = _123_v42
			_123_v42 = _dafny.SeqOf(_dafny.IntOfInt64(897))
			var _124_v43 D0
			_ = _124_v43
			_124_v43 = Companion_D0_.Create_DC1_(_81_v4)
			var _125_v44 _dafny.Array
			_ = _125_v44
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw23
			_125_v44 = _nw23
			var _126_v45 _dafny.Sequence
			_ = _126_v45
			_126_v45 = _dafny.SeqOf(_125_v44, _125_v44, _125_v44)
			var _127_v46 D8
			_ = _127_v46
			_127_v46 = Companion_D8_.Create_DC20_(_126_v45)
			var _nwElement0_5 _dafny.Int = _dafny.IntOfInt64(935)
			_ = _nwElement0_5
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(17))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_5, 0)
			(_nw24).ArraySet1(_78_v1, 1)
			(_nw24).ArraySet1(((_122_v41).Update(_79_v2, (_118_v37).F11())).Cardinality(), 2)
			(_nw24).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_123_v42, _123_v42)).Cardinality()), 3)
			(_nw24).ArraySet1(_78_v1, 4)
			(_nw24).ArraySet1((_78_v1).Minus(_78_v1), 5)
			(_nw24).ArraySet1((_dafny.IntOfInt64(195)).Minus(_78_v1), 6)
			(_nw24).ArraySet1(_78_v1, 7)
			(_nw24).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm0((_124_v43).Dtor_cf1(), _76_globalState)), 8)
			(_nw24).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_78_v1), _78_v1), 9)
			(_nw24).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_77_v0, (Companion_Default___.SafeIndex(_78_v1, _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32(), _81_v4)).Cardinality())).Times(_78_v1), 10)
			(_nw24).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, true)).Cardinality(), 11)
			(_nw24).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if _79_v2 {
					return _123_v42
				}
				return _dafny.SeqOf(_78_v1, _78_v1, _78_v1)
			})()).Cardinality()), 12)
			(_nw24).ArraySet1((_dafny.Zero).Minus(_78_v1), 13)
			(_nw24).ArraySet1(_78_v1, 14)
			(_nw24).ArraySet1(_dafny.IntOfUint32(((_127_v46).Dtor_cf31()).Cardinality()), 15)
			(_nw24).ArraySet1(Companion_Default___.SafeModuloInt(_78_v1, _78_v1), 16)
			_121_v40 = _nw24
			var _128_v47 *C1
			_ = _128_v47
			var _nw25 *C1 = New_C1_()
			_ = _nw25
			_nw25.Ctor__((_118_v37).F11(), _78_v1, _78_v1, _78_v1)
			_128_v47 = _nw25
			var _129_v48 _dafny.Map
			_ = _129_v48
			_129_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v47, (_128_v47).F12())
			var _130_v49 _dafny.Map
			_ = _130_v49
			_130_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _128_v47)
			_129_v48 = (_129_v48).Update((func() *C1 {
				if (_130_v49).Contains(Companion_Default___.Fm0(_81_v4, _76_globalState)) {
					return (_130_v49).Get(Companion_Default___.Fm0(_81_v4, _76_globalState)).(*C1)
				}
				return _128_v47
			})(), _79_v2)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_125_v44), 0))
			_ = _index11
			(_125_v44).ArraySet1(_79_v2, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_125_v44), 0))
			_ = _index12
			(_125_v44).ArraySet1(false, (_index12).Int())
			var _131_v50 D6
			_ = _131_v50
			_131_v50 = Companion_D6_.Create_DC17_(_77_v0, _79_v2, _dafny.IntOfUint32((_80_v3).Cardinality()), Companion_Default___.Fm2((_125_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_125_v44), 0))).Int()).(bool), _76_globalState), (_125_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_125_v44), 0))).Int()).(bool))
			(_128_v47).F13 = (_131_v50).Dtor_cf25()
		} else {
			var _132_v51 _dafny.Array
			_ = _132_v51
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw26
			_132_v51 = _nw26
			var _133_v52 _dafny.Set
			_ = _133_v52
			_133_v52 = _dafny.SetOf(_78_v1)
			var _134_v53 _dafny.Sequence
			_ = _134_v53
			_134_v53 = _dafny.SeqOf((_82_v5).Cardinality(), _78_v1)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_132_v51), 0))
			_ = _index13
			(_132_v51).ArraySet1(Companion_Default___.SafeModuloInt((_133_v52).Cardinality(), (_134_v53).Select((Companion_Default___.SafeIndex(_78_v1, _dafny.IntOfUint32((_134_v53).Cardinality()))).Uint32()).(_dafny.Int)), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_132_v51), 0))
			_ = _index14
			var _rhs10 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfInt64(937)).Minus((_78_v1).Times(_78_v1)))
			_ = _rhs10
			var _rhs11 bool = _79_v2
			_ = _rhs11
			var _rhs12 _dafny.Int = _78_v1
			_ = _rhs12
			var _lhs7 *GlobalState = _76_globalState
			_ = _lhs7
			var _lhs8 _dafny.Array = _132_v51
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_132_v51), 0))
			_ = _lhs9
			_78_v1 = _rhs10
			_lhs7.F1 = _rhs11
			(_lhs8).ArraySet1(_rhs12, (_lhs9).Int())
			_132_v51 = _132_v51
			(_76_globalState).F3 = ((_132_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_132_v51), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(172)) <= 0
			var _135_v54 _dafny.Set
			_ = _135_v54
			_135_v54 = _dafny.SetOf(_79_v2)
			var _136_v55 _dafny.Sequence
			_ = _136_v55
			_136_v55 = _dafny.SeqOf(_135_v54)
			var _137_v56 _dafny.Array
			_ = _137_v56
			var _nwElement0_6 _dafny.Set = _135_v54
			_ = _nwElement0_6
			var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
			_ = _nw27
			(_nw27).ArraySet1(_nwElement0_6, 0)
			(_nw27).ArraySet1(_135_v54, 1)
			(_nw27).ArraySet1(_135_v54, 2)
			(_nw27).ArraySet1((_dafny.SetOf(_79_v2, _79_v2, (Companion_D6_.Create_DC17_(_77_v0, _79_v2, _78_v1, (_80_v3).Select((Companion_Default___.SafeIndex((_132_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_132_v51), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_80_v3).Cardinality()))).Uint32()).(bool), !(true))).Dtor_cf24())).Difference(_135_v54), 3)
			(_nw27).ArraySet1((_136_v55).Select((Companion_Default___.SafeIndex((_132_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_132_v51), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_136_v55).Cardinality()))).Uint32()).(_dafny.Set), 4)
			(_nw27).ArraySet1(_135_v54, 5)
			(_nw27).ArraySet1(_135_v54, 6)
			(_nw27).ArraySet1(_dafny.SetOf(_79_v2, _79_v2), 7)
			(_nw27).ArraySet1(_dafny.SetOf(_79_v2, _79_v2, _79_v2), 8)
			(_nw27).ArraySet1(_135_v54, 9)
			(_nw27).ArraySet1(_135_v54, 10)
			(_nw27).ArraySet1(_135_v54, 11)
			(_nw27).ArraySet1((_dafny.SetOf(_79_v2, _79_v2)).Intersection(_135_v54), 12)
			_137_v56 = _nw27
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_137_v56), 0))
			_ = _index15
			(_137_v56).ArraySet1(_135_v54, (_index15).Int())
			var _138_v57 _dafny.Array
			_ = _138_v57
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw28
			_138_v57 = _nw28
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_138_v57), 0))
			_ = _index16
			(_138_v57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vyygsh"), _77_v0), (_index16).Int())
			var _139_v58 D7
			_ = _139_v58
			_139_v58 = Companion_D7_.Create_DC19_((_77_v0).Select((Companion_Default___.SafeIndex(_78_v1, _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), _79_v2)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_137_v56), 0))
			_ = _index17
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_138_v57), 0))
			_ = _index18
			var _rhs13 bool = (_139_v58).Dtor_cf30()
			_ = _rhs13
			var _rhs14 _dafny.Set = _135_v54
			_ = _rhs14
			var _rhs15 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_77_v0, _77_v0)
			_ = _rhs15
			var _lhs10 *GlobalState = _76_globalState
			_ = _lhs10
			var _lhs11 _dafny.Array = _137_v56
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_137_v56), 0))
			_ = _lhs12
			var _lhs13 _dafny.Array = _138_v57
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_138_v57), 0))
			_ = _lhs14
			_lhs10.F1 = _rhs13
			(_lhs11).ArraySet1(_rhs14, (_lhs12).Int())
			(_lhs13).ArraySet1(_rhs15, (_lhs14).Int())
			var _140_v59 *C3
			_ = _140_v59
			var _nw29 *C3 = New_C3_()
			_ = _nw29
			_nw29.Ctor__(_78_v1, _78_v1)
			_140_v59 = _nw29
			_140_v59 = _140_v59
		}
		var _141_v60 *C0
		_ = _141_v60
		var _nw30 *C0 = New_C0_()
		_ = _nw30
		_nw30.Ctor__(_79_v2)
		_141_v60 = _nw30
		_141_v60 = _141_v60
		_77_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_77_v0, _77_v0), _dafny.Companion_Sequence_.Concatenate(_77_v0, _77_v0))
		_77_v0 = _77_v0
		var _142_v61 _dafny.Array
		_ = _142_v61
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
		_ = _nw31
		_142_v61 = _nw31
		var _143_v62 _dafny.Sequence
		_ = _143_v62
		_143_v62 = _dafny.SeqOf(_78_v1, _78_v1, _78_v1)
		var _144_v63 *C2
		_ = _144_v63
		var _nw32 *C2 = New_C2_()
		_ = _nw32
		_nw32.Ctor__(_78_v1, (_143_v62).Select((Companion_Default___.SafeIndex(_78_v1, _dafny.IntOfUint32((_143_v62).Cardinality()))).Uint32()).(_dafny.Int), _78_v1, _78_v1)
		_144_v63 = _nw32
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_142_v61), 0))
		_ = _index19
		(_142_v61).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v63, (_144_v63).Fm7(_77_v0, _144_v63.F10, _76_globalState)), (_index19).Int())
		var _145_v64 _dafny.Map
		_ = _145_v64
		_145_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v63, _144_v63.F10)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_142_v61), 0))
		_ = _index20
		(_142_v61).ArraySet1(_145_v64, (_index20).Int())
	} else {
		_78_v1 = _78_v1
		var _146_v65 _dafny.Map
		_ = _146_v65
		_146_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.Fm0(_81_v4, _76_globalState)), _78_v1))
		var _147_v66 D1
		_ = _147_v66
		_147_v66 = Companion_D1_.Create_DC2_(_79_v2)
		var _148_v67 _dafny.Map
		_ = _148_v67
		_148_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_82_v5).Contains(true) {
				return (_82_v5).Get(true).(_dafny.Int)
			}
			return _78_v1
		})(), _78_v1)
		_146_v65 = (_146_v65).Update((_147_v66).Dtor_cf2(), _148_v67)
		var _149_v68 _dafny.Array
		_ = _149_v68
		var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw33
		_149_v68 = _nw33
		var _150_v69 *C3
		_ = _150_v69
		var _nw34 *C3 = New_C3_()
		_ = _nw34
		_nw34.Ctor__(_dafny.IntOfInt64(777), _78_v1)
		_150_v69 = _nw34
		var _151_v70 _dafny.Array
		_ = _151_v70
		var _nwElement0_7 *C3 = _150_v69
		_ = _nwElement0_7
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(4))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_7, 0)
		(_nw35).ArraySet1(_150_v69, 1)
		(_nw35).ArraySet1(_150_v69, 2)
		(_nw35).ArraySet1(_150_v69, 3)
		_151_v70 = _nw35
		var _152_v71 D6
		_ = _152_v71
		_152_v71 = Companion_D6_.Create_DC17_(_77_v0, _79_v2, _78_v1, _79_v2, _79_v2)
		var _153_v72 _dafny.Map
		_ = _153_v72
		_153_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, true)
		var _154_v73 _dafny.Array
		_ = _154_v73
		var _nwElement0_8 _dafny.Int = (_dafny.Zero).Minus(_78_v1)
		_ = _nwElement0_8
		var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(14))
		_ = _nw36
		(_nw36).ArraySet1(_nwElement0_8, 0)
		(_nw36).ArraySet1(_78_v1, 1)
		(_nw36).ArraySet1(_78_v1, 2)
		(_nw36).ArraySet1(_78_v1, 3)
		(_nw36).ArraySet1(_78_v1, 4)
		(_nw36).ArraySet1((_dafny.Zero).Minus((_dafny.MultiSetOf(_151_v70, _151_v70, _151_v70)).Cardinality()), 5)
		(_nw36).ArraySet1((_150_v69).Fm7((_152_v71).Dtor_cf23(), _78_v1, _76_globalState), 6)
		(_nw36).ArraySet1(_78_v1, 7)
		(_nw36).ArraySet1((_150_v69).Fm7(_77_v0, _78_v1, _76_globalState), 8)
		(_nw36).ArraySet1(_dafny.IntOfInt64(608), 9)
		(_nw36).ArraySet1((_153_v72).Cardinality(), 10)
		(_nw36).ArraySet1((_150_v69).Fm7(_77_v0, _78_v1, _76_globalState), 11)
		(_nw36).ArraySet1((_dafny.Zero).Minus((_82_v5).Cardinality()), 12)
		(_nw36).ArraySet1(_78_v1, 13)
		_154_v73 = _nw36
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_149_v68), 0))
		_ = _index21
		(_149_v68).ArraySet1(_154_v73, (_index21).Int())
		var _155_v74 _dafny.MultiSet
		_ = _155_v74
		_155_v74 = _dafny.MultiSetOf(_79_v2)
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_149_v68), 0))
		_ = _index22
		var _nwElement0_9 _dafny.Int = _78_v1
		_ = _nwElement0_9
		var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(10))
		_ = _nw37
		(_nw37).ArraySet1(_nwElement0_9, 0)
		(_nw37).ArraySet1((_78_v1).Minus(_dafny.IntOfInt64(591)), 1)
		(_nw37).ArraySet1(_78_v1, 2)
		(_nw37).ArraySet1((_78_v1).Times((_155_v74).Cardinality()), 3)
		(_nw37).ArraySet1((_78_v1).Minus((_153_v72).Cardinality()), 4)
		(_nw37).ArraySet1(_78_v1, 5)
		(_nw37).ArraySet1((_78_v1).Times(_dafny.IntOfInt64(-336)), 6)
		(_nw37).ArraySet1(_78_v1, 7)
		(_nw37).ArraySet1(Companion_Default___.Fm0((_77_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), _76_globalState), 8)
		(_nw37).ArraySet1(_78_v1, 9)
		(_149_v68).ArraySet1(_nw37, (_index22).Int())
		var _156_v75 D4
		_ = _156_v75
		_156_v75 = Companion_D4_.Create_DC12_(_78_v1, _dafny.SeqOf(_79_v2), !(!(_79_v2)))
		var _157_v76 D7
		_ = _157_v76
		_157_v76 = Companion_D7_.Create_DC19_(_81_v4, _79_v2)
		var _158_v77 _dafny.Set
		_ = _158_v77
		_158_v77 = _dafny.SetOf(_78_v1, (_dafny.Zero).Minus((func() _dafny.Int {
			if (_155_v74).Contains(_79_v2) {
				return (_155_v74).Multiplicity(_79_v2)
			}
			return _dafny.IntOfInt64(179)
		})()), _78_v1)
		var _159_v78 _dafny.Array
		_ = _159_v78
		var _nwElement0_10 bool = _79_v2
		_ = _nwElement0_10
		var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(25))
		_ = _nw38
		(_nw38).ArraySet1(_nwElement0_10, 0)
		(_nw38).ArraySet1(_79_v2, 1)
		(_nw38).ArraySet1(_79_v2, 2)
		(_nw38).ArraySet1((_79_v2) || (false), 3)
		(_nw38).ArraySet1((_79_v2) || (_79_v2), 4)
		(_nw38).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_80_v3, _80_v3), 5)
		(_nw38).ArraySet1(!(_79_v2) || (_79_v2), 6)
		(_nw38).ArraySet1(_79_v2, 7)
		(_nw38).ArraySet1(_79_v2, 8)
		(_nw38).ArraySet1(!(_79_v2) || ((_156_v75).Dtor_cf17()), 9)
		(_nw38).ArraySet1(_79_v2, 10)
		(_nw38).ArraySet1(_79_v2, 11)
		(_nw38).ArraySet1(_79_v2, 12)
		(_nw38).ArraySet1(_79_v2, 13)
		(_nw38).ArraySet1(_79_v2, 14)
		(_nw38).ArraySet1(_79_v2, 15)
		(_nw38).ArraySet1(((_dafny.Zero).Minus((_156_v75).Dtor_cf15())).Cmp(_78_v1) > 0, 16)
		(_nw38).ArraySet1(_79_v2, 17)
		(_nw38).ArraySet1((_157_v76).Dtor_cf30(), 18)
		(_nw38).ArraySet1(!(!((_80_v3).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_158_v77).Cardinality()), _dafny.IntOfUint32((_80_v3).Cardinality()))).Uint32()).(bool)) || (_79_v2)), 19)
		(_nw38).ArraySet1(_79_v2, 20)
		(_nw38).ArraySet1(_79_v2, 21)
		(_nw38).ArraySet1(_79_v2, 22)
		(_nw38).ArraySet1(!(_79_v2) || (_79_v2), 23)
		(_nw38).ArraySet1(!(_79_v2), 24)
		_159_v78 = _nw38
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_159_v78), 0))
		_ = _index23
		(_159_v78).ArraySet1(_79_v2, (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_159_v78), 0))
		_ = _index24
		(_159_v78).ArraySet1(_79_v2, (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_154_v73), 0))
		_ = _index25
		(_154_v73).ArraySet1(_78_v1, (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_154_v73), 0))
		_ = _index26
		(_154_v73).ArraySet1(((_78_v1).Plus(_78_v1)).Minus(_78_v1), (_index26).Int())
	}
	(_76_globalState).F3 = _79_v2
	(_76_globalState).F1 = _79_v2
	var _160_v79 _dafny.Sequence
	_ = _160_v79
	_160_v79 = _dafny.SeqOf(_78_v1)
	_160_v79 = _160_v79
	var _161_v80 _dafny.Array
	_ = _161_v80
	var _nwElement0_11 _dafny.CodePoint = _81_v4
	_ = _nwElement0_11
	var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(16))
	_ = _nw39
	(_nw39).ArraySet1CodePoint(_nwElement0_11, 0)
	(_nw39).ArraySet1CodePoint((_77_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_80_v3).Cardinality()), _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
	(_nw39).ArraySet1CodePoint(_81_v4, 2)
	(_nw39).ArraySet1CodePoint(_dafny.CodePoint('a'), 3)
	(_nw39).ArraySet1CodePoint(_81_v4, 4)
	(_nw39).ArraySet1CodePoint(_81_v4, 5)
	(_nw39).ArraySet1CodePoint(_dafny.CodePoint('f'), 6)
	(_nw39).ArraySet1CodePoint(_81_v4, 7)
	(_nw39).ArraySet1CodePoint(_81_v4, 8)
	(_nw39).ArraySet1CodePoint(_81_v4, 9)
	(_nw39).ArraySet1CodePoint(_81_v4, 10)
	(_nw39).ArraySet1CodePoint(_81_v4, 11)
	(_nw39).ArraySet1CodePoint(_81_v4, 12)
	(_nw39).ArraySet1CodePoint(_81_v4, 13)
	(_nw39).ArraySet1CodePoint(_81_v4, 14)
	(_nw39).ArraySet1CodePoint(_dafny.CodePoint('j'), 15)
	_161_v80 = _nw39
	var _162_v81 D9
	_ = _162_v81
	_162_v81 = Companion_D9_.Create_DC22_(_161_v80)
	_161_v80 = (_162_v81).Dtor_cf33()
	var _hi1 _dafny.Int = (func() _dafny.Int {
		if _79_v2 {
			return _78_v1
		}
		return (_dafny.SetOf(_dafny.IntOfUint32((Companion_Default___.Fm16(_76_globalState)).Cardinality()), _78_v1)).Cardinality()
	})()
	_ = _hi1
	for _163_i2 := _78_v1; _163_i2.Cmp(_hi1) < 0; _163_i2 = _163_i2.Plus(_dafny.One) {
		if _79_v2 {
			_78_v1 = Companion_Default___.SafeModuloInt(_163_i2, _163_i2)
			var _164_v82 _dafny.Map
			_ = _164_v82
			_164_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_78_v1), _163_i2), _163_i2)
			_164_v82 = _164_v82
			var _165_v83 *C3
			_ = _165_v83
			var _nw40 *C3 = New_C3_()
			_ = _nw40
			_nw40.Ctor__(_163_i2, _163_i2)
			_165_v83 = _nw40
			_165_v83 = _165_v83
			var _166_v84 _dafny.Array
			_ = _166_v84
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw41
			_166_v84 = _nw41
			var _167_v85 _dafny.Array
			_ = _167_v85
			var _168_v86 bool
			_ = _168_v86
			var _169_v87 _dafny.Sequence
			_ = _169_v87
			var _out0 _dafny.Array
			_ = _out0
			var _out1 bool
			_ = _out1
			var _out2 _dafny.Sequence
			_ = _out2
			_out0, _out1, _out2 = (_165_v83).M4(_166_v84, _76_globalState)
			_167_v85 = _out0
			_168_v86 = _out1
			_169_v87 = _out2
			var _170_v88 _dafny.Array
			_ = _170_v88
			var _nwElement0_12 bool = _168_v86
			_ = _nwElement0_12
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(4))
			_ = _nw42
			(_nw42).ArraySet1(_nwElement0_12, 0)
			(_nw42).ArraySet1(_79_v2, 1)
			(_nw42).ArraySet1(_168_v86, 2)
			(_nw42).ArraySet1(_79_v2, 3)
			_170_v88 = _nw42
			var _171_v89 _dafny.Map
			_ = _171_v89
			_171_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_v4, _79_v2)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_170_v88), 0))
			_ = _index27
			(_170_v88).ArraySet1((func() bool {
				if (_171_v89).Contains(_81_v4) {
					return (_171_v89).Get(_81_v4).(bool)
				}
				return _79_v2
			})(), (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_170_v88), 0))
			_ = _index28
			var _rhs16 bool = _79_v2
			_ = _rhs16
			var _rhs17 _dafny.CodePoint = _81_v4
			_ = _rhs17
			var _rhs18 bool = !(_168_v86) || (_168_v86)
			_ = _rhs18
			var _lhs15 _dafny.Array = _170_v88
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_170_v88), 0))
			_ = _lhs16
			var _lhs17 *GlobalState = _76_globalState
			_ = _lhs17
			(_lhs15).ArraySet1(_rhs16, (_lhs16).Int())
			_81_v4 = _rhs17
			_lhs17.F3 = _rhs18
		} else {
			var _172_v90 _dafny.Array
			_ = _172_v90
			var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw43
			_172_v90 = _nw43
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_172_v90), 0))
			_ = _index29
			(_172_v90).ArraySet1(Companion_Default___.Fm0(_81_v4, _76_globalState), (_index29).Int())
			var _173_v91 _dafny.Sequence
			_ = _173_v91
			_173_v91 = _dafny.SeqOf(_77_v0)
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_172_v90), 0))
			_ = _index30
			(_172_v90).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_173_v91).Cardinality())).Times(_163_i2)), (_index30).Int())
			_78_v1 = (func() _dafny.Int {
				if (func() bool {
					if true {
						return _79_v2
					}
					return Companion_Default___.Fm2(_79_v2, _76_globalState)
				})() {
					return Companion_Default___.SafeDivisionInt(_78_v1, _163_i2)
				}
				return _78_v1
			})()
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_172_v90), 0))
			_ = _index31
			(_172_v90).ArraySet1(((_dafny.Zero).Minus(_163_i2)).Plus(Companion_Default___.Fm0(_81_v4, _76_globalState)), (_index31).Int())
			var _174_v92 _dafny.Map
			_ = _174_v92
			_174_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_172_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_172_v90), 0))).Int()).(_dafny.Int), _79_v2)
			var _175_v93 _dafny.Map
			_ = _175_v93
			_175_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _79_v2)
			var _176_v94 _dafny.Array
			_ = _176_v94
			var _nwElement0_13 bool = true
			_ = _nwElement0_13
			var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(7))
			_ = _nw44
			(_nw44).ArraySet1(_nwElement0_13, 0)
			(_nw44).ArraySet1(!(_79_v2), 1)
			(_nw44).ArraySet1(_79_v2, 2)
			(_nw44).ArraySet1(_79_v2, 3)
			(_nw44).ArraySet1((func() bool {
				if (_174_v92).Contains(_78_v1) {
					return (_174_v92).Get(_78_v1).(bool)
				}
				return _79_v2
			})(), 4)
			(_nw44).ArraySet1((func() bool {
				if (_175_v93).Contains(_79_v2) {
					return (_175_v93).Get(_79_v2).(bool)
				}
				return _79_v2
			})(), 5)
			(_nw44).ArraySet1(_79_v2, 6)
			_176_v94 = _nw44
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_176_v94), 0))
			_ = _index32
			(_176_v94).ArraySet1(_79_v2, (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_176_v94), 0))
			_ = _index33
			(_176_v94).ArraySet1((_80_v3).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_78_v1, _163_i2), _dafny.IntOfUint32((_80_v3).Cardinality()))).Uint32()).(bool), (_index33).Int())
			(_76_globalState).F1 = (func() bool {
				if _79_v2 {
					return _79_v2
				}
				return (_dafny.SetOf((_176_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_176_v94), 0))).Int()).(bool))).IsProperSubsetOf(_dafny.SetOf((_176_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_176_v94), 0))).Int()).(bool), true))
			})()
		}
		var _source3 D8 = Companion_D8_.Create_DC21_(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-625), _163_i2))
		_ = _source3
		if _source3.Is_DC21() {
			var _177___mcc_h0 _dafny.Int = _source3.Get_().(D8_DC21).Cf32
			_ = _177___mcc_h0
			var _178_cf32 _dafny.Int = _177___mcc_h0
			_ = _178_cf32
			var _179_v95 _dafny.Array
			_ = _179_v95
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_2
			var _nw45 _dafny.Array
			_ = _nw45
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw45 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_180_v79 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_181_i3 _dafny.Int) _dafny.Sequence {
						return _180_v79
					}
				})(_160_v79)
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
			_179_v95 = _nw45
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_179_v95), 0))
			_ = _index34
			(_179_v95).ArraySet1(_160_v79, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_179_v95), 0))
			_ = _index35
			(_179_v95).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_160_v79, _160_v79), (_index35).Int())
			(_117_v36).M1(_76_globalState)
			var _182_v97 _dafny.MultiSet
			_ = _182_v97
			_182_v97 = _dafny.MultiSetOf(_78_v1, _163_i2)
			var _183_v98 *C2
			_ = _183_v98
			var _nw46 *C2 = New_C2_()
			_ = _nw46
			_nw46.Ctor__((func() _dafny.Map {
				var _coll13 = _dafny.NewMapBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_182_v97).Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _184_v96 _dafny.Int
					_184_v96 = interface{}(_compr_13).(_dafny.Int)
					if (_182_v97).Contains(_184_v96) {
						_coll13.Add(Companion_Default___.SafeDivisionInt(_184_v96, (_dafny.Zero).Minus(_163_i2)), _79_v2)
					}
				}
				return _coll13.ToMap()
			}()).Cardinality(), (_dafny.Zero).Minus(_178_cf32), _178_cf32, _178_cf32)
			_183_v98 = _nw46
			var _185_v99 *C0
			_ = _185_v99
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__(_79_v2)
			_185_v99 = _nw47
			var _186_v100 _dafny.Sequence
			_ = _186_v100
			_186_v100 = _dafny.SeqOf(_185_v99, _185_v99)
			var _187_v101 D6
			_ = _187_v101
			_187_v101 = Companion_D6_.Create_DC17_(_77_v0, _79_v2, _dafny.IntOfInt64(-505), true, _79_v2)
			_82_v5 = (_82_v5).Update((_dafny.IntOfUint32((_186_v100).Cardinality())).Cmp((_183_v98).Fm7((_187_v101).Dtor_cf23(), _78_v1, _76_globalState)) <= 0, _163_i2)
		} else {
			var _188___mcc_h1 _dafny.Sequence = _source3.Get_().(D8_DC20).Cf31
			_ = _188___mcc_h1
			var _189_cf31 _dafny.Sequence = _188___mcc_h1
			_ = _189_cf31
			var _190_v102 _dafny.MultiSet
			_ = _190_v102
			_190_v102 = _dafny.MultiSetOf(_81_v4, _81_v4)
			var _191_v103 D4
			_ = _191_v103
			_191_v103 = Companion_D4_.Create_DC12_(_78_v1, _80_v3, _79_v2)
			var _192_v104 _dafny.Map
			_ = _192_v104
			_192_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, (_82_v5).Update(false, _163_i2))
			var _193_v105 _dafny.Array
			_ = _193_v105
			var _nwElement0_14 _dafny.Map = _82_v5
			_ = _nwElement0_14
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(27))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_14, 0)
			(_nw48).ArraySet1(_82_v5, 1)
			(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _dafny.IntOfInt64(797)), 2)
			(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _78_v1), 3)
			(_nw48).ArraySet1(_82_v5, 4)
			(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _163_i2), 5)
			(_nw48).ArraySet1(_82_v5, 6)
			(_nw48).ArraySet1(_82_v5, 7)
			(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _163_i2), 8)
			(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _78_v1), 9)
			(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _78_v1), 10)
			(_nw48).ArraySet1(Companion_Default___.Fm21((_82_v5).Cardinality(), (func() _dafny.Int {
				if (_190_v102).Contains(_81_v4) {
					return (_190_v102).Multiplicity(_81_v4)
				}
				return _dafny.IntOfUint32((_dafny.SeqOf((_191_v103).Dtor_cf15())).Cardinality())
			})(), _76_globalState), 11)
			(_nw48).ArraySet1(_82_v5, 12)
			(_nw48).ArraySet1(Companion_Default___.Fm21((_dafny.Zero).Minus(_163_i2), _78_v1, _76_globalState), 13)
			(_nw48).ArraySet1((func() _dafny.Map {
				if (_192_v104).Contains(_79_v2) {
					return (_192_v104).Get(_79_v2).(_dafny.Map)
				}
				return _82_v5
			})(), 14)
			(_nw48).ArraySet1(_82_v5, 15)
			(_nw48).ArraySet1(_82_v5, 16)
			(_nw48).ArraySet1(_82_v5, 17)
			(_nw48).ArraySet1(_82_v5, 18)
			(_nw48).ArraySet1(_82_v5, 19)
			(_nw48).ArraySet1(_82_v5, 20)
			(_nw48).ArraySet1(_82_v5, 21)
			(_nw48).ArraySet1(_82_v5, 22)
			(_nw48).ArraySet1(_82_v5, 23)
			(_nw48).ArraySet1(_82_v5, 24)
			(_nw48).ArraySet1(_82_v5, 25)
			(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _78_v1), 26)
			_193_v105 = _nw48
			var _194_v106 _dafny.Array
			_ = _194_v106
			var _nwElement0_15 _dafny.Array = _193_v105
			_ = _nwElement0_15
			var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(11))
			_ = _nw49
			(_nw49).ArraySet1(_nwElement0_15, 0)
			(_nw49).ArraySet1(_193_v105, 1)
			(_nw49).ArraySet1(_193_v105, 2)
			(_nw49).ArraySet1(_193_v105, 3)
			(_nw49).ArraySet1(_193_v105, 4)
			(_nw49).ArraySet1(_193_v105, 5)
			(_nw49).ArraySet1(_193_v105, 6)
			(_nw49).ArraySet1(_193_v105, 7)
			(_nw49).ArraySet1(_193_v105, 8)
			(_nw49).ArraySet1(_193_v105, 9)
			(_nw49).ArraySet1(_193_v105, 10)
			_194_v106 = _nw49
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_194_v106), 0))
			_ = _index36
			(_194_v106).ArraySet1(_193_v105, (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_194_v106), 0))
			_ = _index37
			(_194_v106).ArraySet1(_193_v105, (_index37).Int())
			_78_v1 = _dafny.IntOfUint32((_77_v0).Cardinality())
			var _195_v107 _dafny.Array
			_ = _195_v107
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_3
			var _nw50 _dafny.Array
			_ = _nw50
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw50 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = func(_196_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(117))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg6 _dafny.Int) interface{} {
							return coer6(arg6)
						}
					}(func(_197_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					}))
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
			_195_v107 = _nw50
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_195_v107), 0))
			_ = _index38
			(_195_v107).ArraySet1(_77_v0, (_index38).Int())
			var _198_v108 _dafny.Array
			_ = _198_v108
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_4
			var _nw51 _dafny.Array
			_ = _nw51
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw51 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Int = func(_199_i6 _dafny.Int) _dafny.Int {
					return (_199_i6).Times(_dafny.IntOfInt64(281))
				}
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw51 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw51).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw51).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_198_v108 = _nw51
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_198_v108), 0))
			_ = _index39
			(_198_v108).ArraySet1((func() _dafny.Int {
				if _79_v2 {
					return _78_v1
				}
				return _78_v1
			})(), (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_195_v107), 0))
			_ = _index40
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_198_v108), 0))
			_ = _index41
			var _rhs19 bool = _79_v2
			_ = _rhs19
			var _rhs20 _dafny.Sequence = _77_v0
			_ = _rhs20
			var _rhs21 _dafny.Int = _78_v1
			_ = _rhs21
			var _rhs22 bool = _79_v2
			_ = _rhs22
			var _lhs18 *GlobalState = _76_globalState
			_ = _lhs18
			var _lhs19 _dafny.Array = _195_v107
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_195_v107), 0))
			_ = _lhs20
			var _lhs21 _dafny.Array = _198_v108
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_198_v108), 0))
			_ = _lhs22
			var _lhs23 *GlobalState = _76_globalState
			_ = _lhs23
			_lhs18.F1 = _rhs19
			(_lhs19).ArraySet1(_rhs20, (_lhs20).Int())
			(_lhs21).ArraySet1(_rhs21, (_lhs22).Int())
			_lhs23.F3 = _rhs22
			var _200_v109 T0
			_ = _200_v109
			var _nw52 *C2 = New_C2_()
			_ = _nw52
			_nw52.Ctor__(_163_i2, _78_v1, _163_i2, (_198_v108).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_198_v108), 0))).Int()).(_dafny.Int))
			_200_v109 = _nw52
			var _201_v110 _dafny.Sequence
			_ = _201_v110
			_201_v110 = _dafny.SeqOf(_200_v109)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_198_v108), 0))
			_ = _index42
			(_198_v108).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if _79_v2 {
					return (_198_v108).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_198_v108), 0))).Int()).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if _79_v2 {
						return _dafny.IntOfUint32((_201_v110).Cardinality())
					}
					return (_200_v109).F8()
				})()
			})()), (_index42).Int())
		}
		(_76_globalState).F1 = (_80_v3).Select((Companion_Default___.SafeIndex(_78_v1, _dafny.IntOfUint32((_80_v3).Cardinality()))).Uint32()).(bool)
		if !((_78_v1).Cmp(_163_i2) > 0) {
			var _202_v111 _dafny.Map
			_ = _202_v111
			_202_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_i2, _163_i2)
			var _203_v112 _dafny.Set
			_ = _203_v112
			_203_v112 = _dafny.SetOf(true)
			var _204_v113 _dafny.Set
			_ = _204_v113
			_204_v113 = _dafny.SetOf((_203_v112).Cardinality(), _163_i2, _163_i2)
			var _205_v114 _dafny.Map
			_ = _205_v114
			_205_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_79_v2)).Update(_79_v2, Companion_Default___.Abs((_204_v113).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rxrjsdrs")).Cardinality()))
			_202_v111 = Companion_Default___.Fm22(_78_v1, (_163_i2).Minus((_205_v114).Cardinality()), true, _160_v79, _76_globalState)
			var _206_v115 _dafny.Array
			_ = _206_v115
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_5
			var _nw53 _dafny.Array
			_ = _nw53
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw53 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_207_v2 bool) func(_dafny.Int) bool {
					return func(_208_i7 _dafny.Int) bool {
						return _207_v2
					}
				})(_79_v2)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw53 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw53).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw53).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_206_v115 = _nw53
			var _209_v116 _dafny.MultiSet
			_ = _209_v116
			_209_v116 = _dafny.MultiSetOf(_206_v115, _206_v115, _206_v115)
			var _210_v117 _dafny.Array
			_ = _210_v117
			var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw54
			_210_v117 = _nw54
			var _211_v118 _dafny.Map
			_ = _211_v118
			_211_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_209_v116, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm11(_79_v2, _163_i2, _76_globalState), _210_v117))
			var _212_v119 _dafny.Map
			_ = _212_v119
			_212_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_77_v0, _210_v117)
			_211_v118 = (_211_v118).Update((_dafny.MultiSetOf(_206_v115, _206_v115)).Difference(_209_v116), _212_v119)
			var _213_v120 D0
			_ = _213_v120
			_213_v120 = Companion_D0_.Create_DC0_(_210_v117)
			var _214_v121 _dafny.Map
			_ = _214_v121
			_214_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(338), _79_v2))
			var _215_v122 _dafny.Map
			_ = _215_v122
			_215_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _79_v2)
			var _216_v123 *C0
			_ = _216_v123
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__(_79_v2)
			_216_v123 = _nw55
			var _217_v124 _dafny.Map
			_ = _217_v124
			_217_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if (_214_v121).Contains(_163_i2) {
					return (_214_v121).Get(_163_i2).(_dafny.Map)
				}
				return _215_v122
			})(), _216_v123)
			var _218_v125 T0
			_ = _218_v125
			var _nw56 *C3 = New_C3_()
			_ = _nw56
			_nw56.Ctor__((_217_v124).Cardinality(), _78_v1)
			_218_v125 = _nw56
			var _219_v126 _dafny.Map
			_ = _219_v126
			_219_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_i2, (func() T0 {
				if _79_v2 {
					return _218_v125
				}
				return _218_v125
			})())
			var _rhs23 D0 = _213_v120
			_ = _rhs23
			var _rhs24 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_82_v5).Contains((_216_v123).F11()) {
					return (_82_v5).Get((_216_v123).F11()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus((_218_v125).F8())
			})(), _218_v125)).Merge((_219_v126).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _218_v125)))
			_ = _rhs24
			var _rhs25 bool = _79_v2
			_ = _rhs25
			var _rhs26 _dafny.Int = (_218_v125).Fm7(_77_v0, (_dafny.IntOfUint32((_77_v0).Cardinality())).Times((_218_v125).F7()), _76_globalState)
			_ = _rhs26
			var _lhs24 *GlobalState = _76_globalState
			_ = _lhs24
			_213_v120 = _rhs23
			_219_v126 = _rhs24
			_lhs24.F3 = _rhs25
			_78_v1 = _rhs26
			_215_v122 = (_215_v122).Update(_dafny.IntOfUint32((_77_v0).Cardinality()), !(_79_v2))
			var _220_v127 _dafny.MultiSet
			_ = _220_v127
			_220_v127 = _dafny.MultiSetOf(false)
			var _221_v128 *C1
			_ = _221_v128
			var _nw57 *C1 = New_C1_()
			_ = _nw57
			_nw57.Ctor__(((_220_v127).Cardinality()).Cmp(_78_v1) != 0, _78_v1, _dafny.IntOfInt64(220), (_218_v125).F7())
			_221_v128 = _nw57
		} else {
			_79_v2 = _79_v2
			var _222_v129 _dafny.Array
			_ = _222_v129
			var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw58
			_222_v129 = _nw58
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_222_v129), 0))
			_ = _index43
			(_222_v129).ArraySet1(_78_v1, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_222_v129), 0))
			_ = _index44
			(_222_v129).ArraySet1((_dafny.Zero).Minus(_163_i2), (_index44).Int())
			var _223_v130 _dafny.Array
			_ = _223_v130
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw59
			_223_v130 = _nw59
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_223_v130), 0))
			_ = _index45
			(_223_v130).ArraySet1(false, (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_223_v130), 0))
			_ = _index46
			var _rhs27 bool = (_80_v3).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_160_v79).Cardinality()), _163_i2), _dafny.IntOfUint32((_80_v3).Cardinality()))).Uint32()).(bool)
			_ = _rhs27
			var _rhs28 bool = ((_222_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_222_v129), 0))).Int()).(_dafny.Int)).Cmp(((_222_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_222_v129), 0))).Int()).(_dafny.Int)).Minus(_78_v1)) >= 0
			_ = _rhs28
			var _lhs25 _dafny.Array = _223_v130
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_223_v130), 0))
			_ = _lhs26
			var _lhs27 *GlobalState = _76_globalState
			_ = _lhs27
			(_lhs25).ArraySet1(_rhs27, (_lhs26).Int())
			_lhs27.F3 = _rhs28
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_223_v130), 0))
			_ = _index47
			(_223_v130).ArraySet1((_223_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_223_v130), 0))).Int()).(bool), (_index47).Int())
			var _224_v131 _dafny.Map
			_ = _224_v131
			_224_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, (_223_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_223_v130), 0))).Int()).(bool))
			(_117_v36).M2(_224_v131, _dafny.Companion_Sequence_.Concatenate(_80_v3, _80_v3), !((_223_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_223_v130), 0))).Int()).(bool)), _76_globalState)
		}
	}
	_78_v1 = (_dafny.Zero).Minus((_78_v1).Times((_78_v1).Plus(_78_v1)))
	_79_v2 = _79_v2
	var _225_v132 _dafny.Set
	_ = _225_v132
	_225_v132 = _dafny.SetOf((_77_v0).Select((Companion_Default___.SafeIndex(_78_v1, _dafny.IntOfUint32((_77_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))
	var _226_v134 _dafny.MultiSet
	_ = _226_v134
	_226_v134 = _dafny.MultiSetOf(_225_v132, func() _dafny.Set {
		var _coll14 = _dafny.NewBuilder()
		_ = _coll14
		for _iter14 := _dafny.Iterate((_225_v132).Elements()); ; {
			_compr_14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _227_v133 _dafny.CodePoint
			_227_v133 = interface{}(_compr_14).(_dafny.CodePoint)
			if (_225_v132).Contains(_227_v133) {
				_coll14.Add(_227_v133)
			}
		}
		return _coll14.ToSet()
	}(), _225_v132)
	var _228_i8 _dafny.Int
	_ = _228_i8
	_228_i8 = _dafny.Zero
	{
		for (_226_v134).IsSubsetOf(_226_v134) {
			{
				if (_228_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_228_i8 = (_228_i8).Plus(_dafny.One)
				var _229_v135 _dafny.Array
				_ = _229_v135
				var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
				_ = _nw60
				_229_v135 = _nw60
				var _230_v136 T0
				_ = _230_v136
				var _nw61 *C1 = New_C1_()
				_ = _nw61
				_nw61.Ctor__(_79_v2, _78_v1, _78_v1, _78_v1)
				_230_v136 = _nw61
				var _231_v137 D2
				_ = _231_v137
				_231_v137 = Companion_D2_.Create_DC5_((_dafny.SetOf(_79_v2)).Cardinality(), _77_v0, _230_v136)
				var _232_v138 D2
				_ = _232_v138
				_232_v138 = Companion_D2_.Create_DC6_(_231_v137)
				var _233_v139 _dafny.Map
				_ = _233_v139
				_233_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_v4, _232_v138)
				var _234_v140 _dafny.Set
				_ = _234_v140
				_234_v140 = _dafny.SetOf(_233_v139)
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_229_v135), 0))
				_ = _index48
				(_229_v135).ArraySet1(_234_v140, (_index48).Int())
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_229_v135), 0))
				_ = _index49
				(_229_v135).ArraySet1(_234_v140, (_index49).Int())
				var _235_v141 D6
				_ = _235_v141
				_235_v141 = Companion_D6_.Create_DC17_(_77_v0, _79_v2, _78_v1, _79_v2, _79_v2)
				var _236_v142 *C1
				_ = _236_v142
				var _nw62 *C1 = New_C1_()
				_ = _nw62
				_nw62.Ctor__(!((_235_v141).Dtor_cf26()), (_230_v136).F7(), _78_v1, ((_230_v136).F7()).Plus((_230_v136).Fm7(_77_v0, (_230_v136).F7(), _76_globalState)))
				_236_v142 = _nw62
				_236_v142 = _236_v142
				var _237_v143 _dafny.Array
				_ = _237_v143
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_6
				var _nw63 _dafny.Array
				_ = _nw63
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw63 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Int = (func(_238_v136 T0) func(_dafny.Int) _dafny.Int {
						return func(_239_i9 _dafny.Int) _dafny.Int {
							return (_239_i9).Plus((_238_v136).F7())
						}
					})(_230_v136)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw63 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw63).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw63).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_237_v143 = _nw63
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_237_v143), 0))
				_ = _index50
				(_237_v143).ArraySet1((_230_v136).F7(), (_index50).Int())
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_237_v143), 0))
				_ = _index51
				(_237_v143).ArraySet1(_78_v1, (_index51).Int())
				var _240_v144 _dafny.Map
				_ = _240_v144
				_240_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v2, _79_v2)
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_237_v143), 0))
				_ = _index52
				(_237_v143).ArraySet1(((_240_v144).Cardinality()).Plus((_78_v1).Times((_236_v142).Fm7(_dafny.UnicodeSeqOfUtf8Bytes("qyb"), _dafny.IntOfUint32((_160_v79).Cardinality()), _76_globalState))), (_index52).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _241_v145 _dafny.Array
	_ = _241_v145
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(3)
	_ = _len0_7
	var _nw64 _dafny.Array
	_ = _nw64
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw64 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) D1 = func(_242_i10 _dafny.Int) D1 {
			return Companion_D1_.Create_DC2_(false)
		}
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw64 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw64).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw64).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	_241_v145 = _nw64
	var _243_v146 D1
	_ = _243_v146
	_243_v146 = Companion_D1_.Create_DC2_(false)
	var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_241_v145), 0))
	_ = _index53
	(_241_v145).ArraySet1(_243_v146, (_index53).Int())
	var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_241_v145), 0))
	_ = _index54
	(_241_v145).ArraySet1(Companion_D1_.Create_DC2_(_79_v2), (_index54).Int())
	if _79_v2 {
		(_76_globalState).F3 = !((_79_v2) || (_79_v2)) || (_79_v2)
		var _244_v147 _dafny.Map
		_ = _244_v147
		_244_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v1, _79_v2)
		var _245_v148 _dafny.Map
		_ = _245_v148
		_245_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_v4, _80_v3)
		var _246_v149 _dafny.Array
		_ = _246_v149
		var _nw65 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw65
		_246_v149 = _nw65
		var _247_v150 _dafny.Set
		_ = _247_v150
		_247_v150 = _dafny.SetOf(_246_v149, _246_v149, _246_v149)
		(_117_v36).M2(_244_v147, (func() _dafny.Sequence {
			if (_245_v148).Contains(_81_v4) {
				return (_245_v148).Get(_81_v4).(_dafny.Sequence)
			}
			return _80_v3
		})(), (_247_v150).IsProperSubsetOf(_dafny.SetOf(_246_v149, _246_v149)), _76_globalState)
		var _248_v151 D6
		_ = _248_v151
		_248_v151 = Companion_D6_.Create_DC17_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(833))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_249_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_250_i11 _dafny.Int) _dafny.CodePoint {
				return _249_v4
			}
		})(_81_v4))), _79_v2, _dafny.IntOfUint32((_77_v0).Cardinality()), _79_v2, _79_v2)
		_248_v151 = _248_v151
		(_117_v36).M1(_76_globalState)
		var _251_v152 *C4
		_ = _251_v152
		var _nw66 *C4 = New_C4_()
		_ = _nw66
		_nw66.Ctor__((_117_v36).F6())
		_251_v152 = _nw66
	} else {
		Companion_Default___.M0(_76_globalState)
		(_76_globalState).F1 = _79_v2
		var _252_v153 _dafny.Map
		_ = _252_v153
		_252_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v5, !(_79_v2))
		_78_v1 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_78_v1), (_252_v153).Cardinality())
		var _253_v154 _dafny.Array
		_ = _253_v154
		var _nw67 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
		_ = _nw67
		_253_v154 = _nw67
		_253_v154 = _253_v154
		(_117_v36).M1(_76_globalState)
	}
	_dafny.Print((_76_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_76_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_76_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_77_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_78_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_79_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_80_v3, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_81_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, true), _dafny.IntOfInt64(414))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v34).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), true).UpdateUnsafe(_dafny.CodePoint('q'), true).UpdateUnsafe(_dafny.CodePoint('i'), true).UpdateUnsafe(_dafny.CodePoint('t'), true).UpdateUnsafe(_dafny.CodePoint('b'), true).UpdateUnsafe(_dafny.CodePoint('m'), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-958), _dafny.CodePoint('q')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_116_v35).Dtor_cf28()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), true).UpdateUnsafe(_dafny.CodePoint('q'), true).UpdateUnsafe(_dafny.CodePoint('i'), true).UpdateUnsafe(_dafny.CodePoint('t'), true).UpdateUnsafe(_dafny.CodePoint('b'), true).UpdateUnsafe(_dafny.CodePoint('m'), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-958), _dafny.CodePoint('q')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_117_v36).F6()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), true).UpdateUnsafe(_dafny.CodePoint('q'), true).UpdateUnsafe(_dafny.CodePoint('i'), true).UpdateUnsafe(_dafny.CodePoint('t'), true).UpdateUnsafe(_dafny.CodePoint('b'), true).UpdateUnsafe(_dafny.CodePoint('m'), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-958), _dafny.CodePoint('q')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_160_v79, _dafny.SeqOf(_dafny.IntOfInt64(-958))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v81).Dtor_cf33()).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v132).Equals(_dafny.SetOf(_dafny.CodePoint('p'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_226_v134).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.CodePoint('p')), _dafny.SetOf(_dafny.CodePoint('p')), _dafny.SetOf(_dafny.CodePoint('p')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_i8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_241_v145).ArrayGet1((_dafny.Zero).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_241_v145).ArrayGet1((_dafny.One).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_241_v145).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D1)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_243_v146).Dtor_cf2())
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.CodePoint) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
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
	return Companion_D0_.Create_DC1_(_dafny.CodePoint('D'))
}

func (_this D0) Dtor_cf1() _dafny.CodePoint {
	return _this.Get_().(D0_DC1).Cf1
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf1 == data2.Cf1
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

type D1_DC3 struct {
	Cf3 _dafny.CodePoint
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.CodePoint) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf2 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 bool) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.CodePoint('D'))
}

func (_this D1) Dtor_cf3() _dafny.CodePoint {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf3 == data2.Cf3
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2 == data2.Cf2
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

type D2_DC5 struct {
	Cf5 _dafny.Int
	Cf6 _dafny.Sequence
	Cf7 T0
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf5 _dafny.Int, Cf6 _dafny.Sequence, Cf7 T0) D2 {
	return D2{D2_DC5{Cf5, Cf6, Cf7}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC4 struct {
	Cf4 _dafny.Int
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf4 _dafny.Int) D2 {
	return D2{D2_DC4{Cf4}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC6 struct {
	Cf8 D2
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf8 D2) D2 {
	return D2{D2_DC6{Cf8}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(_dafny.Zero, _dafny.EmptySeq, (T0)(nil))
}

func (_this D2) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf7() T0 {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf4
}

func (_this D2) Dtor_cf8() D2 {
	return _this.Get_().(D2_DC6).Cf8
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf5) + ", " + data.Cf6.VerbatimString(true) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf8) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6.Equals(data2.Cf6) && _dafny.AreEqual(data1.Cf7, data2.Cf7)
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf10 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf10 _dafny.Int) D3 {
	return D3{D3_DC8{Cf10}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC7 struct {
	Cf9 _dafny.Set
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf9 _dafny.Set) D3 {
	return D3{D3_DC7{Cf9}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC9 struct {
	Cf11 D3
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf11 D3) D3 {
	return D3{D3_DC9{Cf11}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(_dafny.Zero)
}

func (_this D3) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf10
}

func (_this D3) Dtor_cf9() _dafny.Set {
	return _this.Get_().(D3_DC7).Cf9
}

func (_this D3) Dtor_cf11() D3 {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf9.Equals(data2.Cf9)
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf13 _dafny.Int
	Cf14 _dafny.CodePoint
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf13 _dafny.Int, Cf14 _dafny.CodePoint) D4 {
	return D4{D4_DC11{Cf13, Cf14}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC12 struct {
	Cf15 _dafny.Int
	Cf16 _dafny.Sequence
	Cf17 bool
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf15 _dafny.Int, Cf16 _dafny.Sequence, Cf17 bool) D4 {
	return D4{D4_DC12{Cf15, Cf16, Cf17}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC10 struct {
	Cf12 _dafny.Map
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf12 _dafny.Map) D4 {
	return D4{D4_DC10{Cf12}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC13 struct {
	Cf18 D4
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf18 D4) D4 {
	return D4{D4_DC13{Cf18}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D4) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf13
}

func (_this D4) Dtor_cf14() _dafny.CodePoint {
	return _this.Get_().(D4_DC11).Cf14
}

func (_this D4) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D4_DC12).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC12).Cf17
}

func (_this D4) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D4_DC10).Cf12
}

func (_this D4) Dtor_cf18() D4 {
	return _this.Get_().(D4_DC13).Cf18
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Equals(data2.Cf16) && data1.Cf17 == data2.Cf17
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf12.Equals(data2.Cf12)
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf20 _dafny.Sequence
	Cf21 _dafny.Int
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf20 _dafny.Sequence, Cf21 _dafny.Int) D5 {
	return D5{D5_DC15{Cf20, Cf21}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC14 struct {
	Cf19 _dafny.Array
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf19 _dafny.Array) D5 {
	return D5{D5_DC14{Cf19}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this D5) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D5_DC15).Cf20
}

func (_this D5) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D5_DC15).Cf21
}

func (_this D5) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf19
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + data.Cf20.VerbatimString(true) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf20.Equals(data2.Cf20) && data1.Cf21.Cmp(data2.Cf21) == 0
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf19 == data2.Cf19
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
	Cf23 _dafny.Sequence
	Cf24 bool
	Cf25 _dafny.Int
	Cf26 bool
	Cf27 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf23 _dafny.Sequence, Cf24 bool, Cf25 _dafny.Int, Cf26 bool, Cf27 bool) D6 {
	return D6{D6_DC17{Cf23, Cf24, Cf25, Cf26, Cf27}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf22 _dafny.Sequence
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf22 _dafny.Sequence) D6 {
	return D6{D6_DC16{Cf22}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(_dafny.EmptySeq, false, _dafny.Zero, false, false)
}

func (_this D6) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf23
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC17).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf25
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC17).Cf26
}

func (_this D6) Dtor_cf27() bool {
	return _this.Get_().(D6_DC17).Cf27
}

func (_this D6) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + data.Cf23.VerbatimString(true) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf23.Equals(data2.Cf23) && data1.Cf24 == data2.Cf24 && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf22.Equals(data2.Cf22)
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
	Cf29 _dafny.CodePoint
	Cf30 bool
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf29 _dafny.CodePoint, Cf30 bool) D7 {
	return D7{D7_DC19{Cf29, Cf30}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC18 struct {
	Cf28 _dafny.Map
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf28 _dafny.Map) D7 {
	return D7{D7_DC18{Cf28}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(_dafny.CodePoint('D'), false)
}

func (_this D7) Dtor_cf29() _dafny.CodePoint {
	return _this.Get_().(D7_DC19).Cf29
}

func (_this D7) Dtor_cf30() bool {
	return _this.Get_().(D7_DC19).Cf30
}

func (_this D7) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D7_DC18).Cf28
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf28.Equals(data2.Cf28)
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
	Cf32 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf32 _dafny.Int) D8 {
	return D8{D8_DC21{Cf32}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC20 struct {
	Cf31 _dafny.Sequence
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf31 _dafny.Sequence) D8 {
	return D8{D8_DC20{Cf31}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(_dafny.Zero)
}

func (_this D8) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf32
}

func (_this D8) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D8_DC20).Cf31
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf32.Cmp(data2.Cf32) == 0
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf31.Equals(data2.Cf31)
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
	Cf34 _dafny.Int
	Cf35 _dafny.Int
	Cf36 _dafny.Sequence
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf34 _dafny.Int, Cf35 _dafny.Int, Cf36 _dafny.Sequence) D9 {
	return D9{D9_DC23{Cf34, Cf35, Cf36}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC24 struct {
	Cf37 bool
	Cf38 _dafny.Int
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf37 bool, Cf38 _dafny.Int) D9 {
	return D9{D9_DC24{Cf37, Cf38}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC22 struct {
	Cf33 _dafny.Array
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf33 _dafny.Array) D9 {
	return D9{D9_DC22{Cf33}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC23_(_dafny.Zero, _dafny.Zero, _dafny.EmptySeq)
}

func (_this D9) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf34
}

func (_this D9) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf35
}

func (_this D9) Dtor_cf36() _dafny.Sequence {
	return _this.Get_().(D9_DC23).Cf36
}

func (_this D9) Dtor_cf37() bool {
	return _this.Get_().(D9_DC24).Cf37
}

func (_this D9) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf38
}

func (_this D9) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D9_DC22).Cf33
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf33) + ")"
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
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36.Equals(data2.Cf36)
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf33 == data2.Cf33
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm6(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Set
	Fm7(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Int
	F7() _dafny.Int
	F8() _dafny.Int
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
	F1  bool
	F3  bool
	_f0 bool
	_f2 _dafny.Int
	_f4 bool
	_f5 _dafny.CodePoint
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = false
	_this.F3 = false
	_this._f0 = false
	_this._f2 = _dafny.Zero
	_this._f4 = false
	_this._f5 = _dafny.CodePoint('D')
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 _dafny.Int, f3 bool, f4 bool, f5 _dafny.CodePoint) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
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
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.CodePoint {
	{
		return _this._f5
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f11 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f11 = false
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

func (_this *C0) Ctor__(f11 bool) {
	{
		(_this)._f11 = f11
	}
}
func (_this *C0) F11() bool {
	{
		return _this._f11
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f7  _dafny.Int
	_f8  _dafny.Int
	F13  _dafny.Int
	_f12 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this.F13 = _dafny.Zero
	_this._f12 = false
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
func (_this *C1) F8() _dafny.Int {
	return _this._f8
}
func (_this *C1) Ctor__(f12 bool, f13 _dafny.Int, f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C1) Fm6(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	{
		return ((_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(175))).Uint32(), func(coer8 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_254_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf((_this).F12())
		}))).Cardinality())))).Difference((Companion_D3_.Create_DC7_(_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (_this).F12())).Cardinality(), (_this).F8()))))).Dtor_cf9())).Union((_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-398)), _dafny.MultiSetOf((_this).F8(), (_this).F8()), _dafny.MultiSetOf((_this).F7()), _dafny.MultiSetOf(_dafny.IntOfInt64(639)), _dafny.MultiSetOf((_this).F7()))).Intersection(_dafny.SetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _dafny.CodePoint('v'))).Cardinality(), (_this).F7(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(239))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_255_i1 _dafny.Int) _dafny.Int {
			return _this.F13
		}))).Cardinality())))))
	}
}
func (_this *C1) Fm7(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F7()
	}
}
func (_this *C1) M6(p0 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, D3) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 D3 = Companion_D3_.Default()
		_ = r1
		var _hi2 _dafny.Int = (_this).F7()
		_ = _hi2
		for _256_i0 := p0; _256_i0.Cmp(_hi2) < 0; _256_i0 = _256_i0.Plus(_dafny.One) {
			(globalState).F1 = ((_this).F7()).Cmp(p0) < 0
			var _257_v0 *C0
			_ = _257_v0
			var _nw68 *C0 = New_C0_()
			_ = _nw68
			_nw68.Ctor__((_this).F12())
			_257_v0 = _nw68
			var _258_v1 _dafny.Array
			_ = _258_v1
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw69
			_258_v1 = _nw69
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_258_v1), 0))
			_ = _index55
			(_258_v1).ArraySet1(Companion_Default___.SafeModuloInt((_this).F7(), (_this).F8()), (_index55).Int())
			var _259_v2 _dafny.Array
			_ = _259_v2
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(26))
			_ = _nw70
			_259_v2 = _nw70
			var _260_v3 _dafny.MultiSet
			_ = _260_v3
			_260_v3 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mpakoftbq")).Cardinality()), _this.F13, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_261_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()))
			var _262_v4 _dafny.Sequence
			_ = _262_v4
			_262_v4 = _dafny.SeqOf(_260_v3)
			var _263_v6 D3
			_ = _263_v6
			_263_v6 = Companion_D3_.Create_DC7_(func() _dafny.Set {
				var _coll15 = _dafny.NewBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate((_262_v4).Elements()); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _264_v5 _dafny.MultiSet
					_264_v5 = interface{}(_compr_15).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_262_v4, _264_v5) {
						_coll15.Add(_264_v5)
					}
				}
				return _coll15.ToSet()
			}())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_259_v2), 0))
			_ = _index56
			(_259_v2).ArraySet1(_263_v6, (_index56).Int())
			var _265_v7 _dafny.MultiSet
			_ = _265_v7
			_265_v7 = _dafny.MultiSetOf((_this).F12(), (_257_v0).F11(), (_this).F12())
			var _266_v8 _dafny.Map
			_ = _266_v8
			_266_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(440), _265_v7)
			var _267_v10 D4
			_ = _267_v10
			_267_v10 = Companion_D4_.Create_DC10_(func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(674), _dafny.IntOfInt64(811))); ; {
					_compr_16, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _268_v9 _dafny.Int
					_268_v9 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(674)).Cmp(_268_v9) <= 0) && ((_268_v9).Cmp(_dafny.IntOfInt64(811)) < 0) {
						_coll16.Add((_268_v9).Plus(_dafny.IntOfInt64(68)), _256_i0)
					}
				}
				return _coll16.ToMap()
			}())
			var _269_v11 _dafny.Set
			_ = _269_v11
			_269_v11 = _dafny.SetOf((_this).F12())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_258_v1), 0))
			_ = _index57
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_259_v2), 0))
			_ = _index58
			var _rhs29 _dafny.Int = ((func() _dafny.MultiSet {
				if (_266_v8).Contains(((_267_v10).Dtor_cf12()).Cardinality()) {
					return (_266_v8).Get(((_267_v10).Dtor_cf12()).Cardinality()).(_dafny.MultiSet)
				}
				return _265_v7
			})()).Cardinality()
			_ = _rhs29
			var _rhs30 D3 = Companion_Default___.Fm12((!((_257_v0).F11())) || (true), (_257_v0).F11(), Companion_Default___.Fm3((_this).F12(), globalState), globalState)
			_ = _rhs30
			var _rhs31 bool = (func() bool {
				if (_this).F12() {
					return ((_this).F7()).Cmp((_269_v11).Cardinality()) != 0
				}
				return (p0).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_257_v0).F11(), (_this).F8())).Cardinality()) == 0
			})()
			_ = _rhs31
			var _lhs28 _dafny.Array = _258_v1
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_258_v1), 0))
			_ = _lhs29
			var _lhs30 _dafny.Array = _259_v2
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_259_v2), 0))
			_ = _lhs31
			var _lhs32 *GlobalState = globalState
			_ = _lhs32
			(_lhs28).ArraySet1(_rhs29, (_lhs29).Int())
			(_lhs30).ArraySet1(_rhs30, (_lhs31).Int())
			_lhs32.F3 = _rhs31
			_258_v1 = _258_v1
		}
		var _270_v12 _dafny.Sequence
		_ = _270_v12
		_270_v12 = _dafny.UnicodeSeqOfUtf8Bytes("nhflilcl")
		_270_v12 = _270_v12
		var _271_v13 *C0
		_ = _271_v13
		var _nw71 *C0 = New_C0_()
		_ = _nw71
		_nw71.Ctor__((_this).F12())
		_271_v13 = _nw71
		var _272_v14 D1
		_ = _272_v14
		_272_v14 = Companion_D1_.Create_DC2_((_271_v13).F11())
		var _273_v15 *C0
		_ = _273_v15
		var _nw72 *C0 = New_C0_()
		_ = _nw72
		_nw72.Ctor__((func() bool {
			if (_272_v14).Dtor_cf2() {
				return (_this).F12()
			}
			return (_this).F12()
		})())
		_273_v15 = _nw72
		var _274_v16 _dafny.Set
		_ = _274_v16
		_274_v16 = _dafny.SetOf(Companion_Default___.SafeModuloInt(_this.F13, _this.F13), _this.F13, (_this).F7(), (_this).F7())
		_274_v16 = _274_v16
		(_this).F13 = Companion_Default___.SafeDivisionInt(p0, (_this).Fm7(_270_v12, (_this).F7(), globalState))
		var _275_v17 _dafny.Sequence
		_ = _275_v17
		_275_v17 = _dafny.SeqOf((_271_v13).F11())
		r0 = _dafny.Companion_Sequence_.Concatenate(_275_v17, _dafny.SeqOf((_this).F12()))
		var _276_v18 D3
		_ = _276_v18
		_276_v18 = Companion_D3_.Create_DC8_(((_this).F8()).Minus(p0))
		r1 = _276_v18
		return r0, r1
	}
}
func (_this *C1) F12() bool {
	{
		return _this._f12
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f7 _dafny.Int
	_f8 _dafny.Int
	F10 _dafny.Int
	_f9 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this.F10 = _dafny.Zero
	_this._f9 = _dafny.Zero
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
func (_this *C2) F8() _dafny.Int {
	return _this._f8
}
func (_this *C2) Ctor__(f9 _dafny.Int, f10 _dafny.Int, f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C2) Fm6(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	{
		return (_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false, !(false), true)).Cardinality())), _dafny.MultiSetOf((_this).F8()))).Difference(_dafny.SetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ql")).Cardinality()))).Cardinality(), (_this).F8())).Cardinality()), _dafny.MultiSetOf((_this).F8(), (_this).F7(), (_this).F9()), _dafny.MultiSetOf(_dafny.IntOfInt64(846), (_this).F8()), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(23))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_277_i0 _dafny.Int) _dafny.Set {
			return _dafny.SetOf((_this).F9())
		}))).Cardinality()))))
	}
}
func (_this *C2) Fm7(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F8()
	}
}
func (_this *C2) Fm9(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((Companion_D1_.Create_DC2_(!(false))).Dtor_cf2()) && (!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("see"), _dafny.UnicodeSeqOfUtf8Bytes("mu")))
	}
}
func (_this *C2) Fm10(p0 bool, p1 bool, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf(false)).IsSubsetOf((_dafny.MultiSetOf(true, false)).Intersection(_dafny.MultiSetOf(true)))
	}
}
func (_this *C2) M5(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		(globalState).F1 = ((_this.F10).Times((_this).F9())).Cmp(_this.F10) >= 0
		var _278_i0 _dafny.Int
		_ = _278_i0
		_278_i0 = _dafny.Zero
		{
			for false {
				{
					if (_278_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_278_i0 = (_278_i0).Plus(_dafny.One)
					var _279_v0 bool
					_ = _279_v0
					_279_v0 = true
					if !(_279_v0) || (true) {
						(_this).F10 = (_this).F7()
						var _280_v1 _dafny.Array
						_ = _280_v1
						var _len0_8 _dafny.Int = _dafny.IntOfInt64(8)
						_ = _len0_8
						var _nw73 _dafny.Array
						_ = _nw73
						if _len0_8.Cmp(_dafny.Zero) == 0 {
							_nw73 = _dafny.NewArray(_len0_8)
						} else {
							var _init8 func(_dafny.Int) bool = (func(_281_v0 bool) func(_dafny.Int) bool {
								return func(_282_i1 _dafny.Int) bool {
									return _281_v0
								}
							})(_279_v0)
							_ = _init8
							var _element0_8 = _init8(_dafny.Zero)
							_ = _element0_8
							_nw73 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
							(_nw73).ArraySet1(_element0_8, 0)
							var _nativeLen0_8 = (_len0_8).Int()
							_ = _nativeLen0_8
							for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
								(_nw73).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
							}
						}
						_280_v1 = _nw73
						_280_v1 = _280_v1
						var _283_v2 _dafny.CodePoint
						_ = _283_v2
						_283_v2 = _dafny.CodePoint('e')
						var _284_v3 *C0
						_ = _284_v3
						var _nw74 *C0 = New_C0_()
						_ = _nw74
						_nw74.Ctor__((_this).Fm10(!(_279_v0), _279_v0, _279_v0, _283_v2, globalState))
						_284_v3 = _nw74
						var _285_v4 _dafny.Sequence
						_ = _285_v4
						_285_v4 = _dafny.SeqOf(_279_v0)
						var _286_v5 D1
						_ = _286_v5
						_286_v5 = Companion_D1_.Create_DC3_(_283_v2)
						var _287_v6 _dafny.Map
						_ = _287_v6
						_287_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_this).F7(), (_dafny.MultiSetFromSeq(_285_v4)).Cardinality()), _286_v5)
						_287_v6 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(561), _286_v5)).Merge(_287_v6)
						var _288_v7 _dafny.Array
						_ = _288_v7
						var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
						_ = _nw75
						_288_v7 = _nw75
						_288_v7 = _288_v7
					} else {
						var _289_v8 _dafny.Set
						_ = _289_v8
						_289_v8 = _dafny.SetOf((_this).F8(), _this.F10)
						_289_v8 = _289_v8
						var _290_v9 D2
						_ = _290_v9
						_290_v9 = Companion_D2_.Create_DC4_((_this).F7())
						(_this).F10 = (_290_v9).Dtor_cf4()
						var _291_v10 _dafny.Set
						_ = _291_v10
						_291_v10 = _dafny.SetOf(_279_v0, Companion_Default___.Fm2(_279_v0, globalState), _279_v0, _279_v0)
						var _292_v11 _dafny.Map
						_ = _292_v11
						_292_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if !(_279_v0) {
								return _279_v0
							}
							return _279_v0
						})(), (_291_v10).Cardinality())
						var _293_v12 _dafny.CodePoint
						_ = _293_v12
						_293_v12 = _dafny.CodePoint('w')
						var _294_v13 _dafny.Array
						_ = _294_v13
						var _len0_9 _dafny.Int = _dafny.One
						_ = _len0_9
						var _nw76 _dafny.Array
						_ = _nw76
						if _len0_9.Cmp(_dafny.Zero) == 0 {
							_nw76 = _dafny.NewArray(_len0_9)
						} else {
							var _init9 func(_dafny.Int) _dafny.CodePoint = (func(_295_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_296_i2 _dafny.Int) _dafny.CodePoint {
									return _295_v12
								}
							})(_293_v12)
							_ = _init9
							var _element0_9 = _init9(_dafny.Zero)
							_ = _element0_9
							_nw76 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
							(_nw76).ArraySet1CodePoint(_element0_9, 0)
							var _nativeLen0_9 = (_len0_9).Int()
							_ = _nativeLen0_9
							for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
								(_nw76).ArraySet1CodePoint(_init9(_dafny.IntOf(_i0_9)), _i0_9)
							}
						}
						_294_v13 = _nw76
						var _297_v14 D1
						_ = _297_v14
						_297_v14 = Companion_D1_.Create_DC2_(_279_v0)
						var _rhs32 _dafny.Map = p1
						_ = _rhs32
						var _rhs33 bool = _279_v0
						_ = _rhs33
						var _rhs34 _dafny.CodePoint = (func() _dafny.CodePoint {
							if (_297_v14).Dtor_cf2() {
								return _dafny.CodePoint('x')
							}
							return _293_v12
						})()
						_ = _rhs34
						var _rhs35 _dafny.Array = _294_v13
						_ = _rhs35
						var _rhs36 bool = _279_v0
						_ = _rhs36
						var _lhs33 *GlobalState = globalState
						_ = _lhs33
						var _lhs34 *GlobalState = globalState
						_ = _lhs34
						_292_v11 = _rhs32
						_lhs33.F1 = _rhs33
						_293_v12 = _rhs34
						_294_v13 = _rhs35
						_lhs34.F3 = _rhs36
						(_this).F10 = (_this).F8()
						var _298_v15 *C0
						_ = _298_v15
						var _nw77 *C0 = New_C0_()
						_ = _nw77
						_nw77.Ctor__(_279_v0)
						_298_v15 = _nw77
					}
					(globalState).F3 = (_this.F10).Cmp(((_this).F8()).Plus(_dafny.IntOfInt64(-750))) < 0
					var _299_v16 D2
					_ = _299_v16
					_299_v16 = Companion_D2_.Create_DC4_((_this).F9())
					var _300_v17 _dafny.Map
					_ = _300_v17
					_300_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), func(_pat_let0_0 D2) D2 {
						return func(_301_dt__update__tmp_h0 D2) D2 {
							return func(_pat_let1_0 _dafny.Int) D2 {
								return func(_302_dt__update_hcf4_h0 _dafny.Int) D2 {
									return Companion_D2_.Create_DC4_(_302_dt__update_hcf4_h0)
								}(_pat_let1_0)
							}((_this).F9())
						}(_pat_let0_0)
					}(_299_v16))
					var _303_v18 _dafny.CodePoint
					_ = _303_v18
					_303_v18 = _dafny.CodePoint('h')
					var _304_v19 _dafny.Sequence
					_ = _304_v19
					_304_v19 = _dafny.SeqOf(_279_v0, _279_v0, _279_v0)
					var _305_v20 _dafny.MultiSet
					_ = _305_v20
					_305_v20 = _dafny.MultiSetOf(_dafny.SeqOf(_279_v0), _304_v19, _304_v19, _304_v19)
					var _306_v21 _dafny.Array
					_ = _306_v21
					var _nwElement0_16 _dafny.Int = (_this).F8()
					_ = _nwElement0_16
					var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(19))
					_ = _nw78
					(_nw78).ArraySet1(_nwElement0_16, 0)
					(_nw78).ArraySet1((_this).F9(), 1)
					(_nw78).ArraySet1((_this).F7(), 2)
					(_nw78).ArraySet1((_this).F7(), 3)
					(_nw78).ArraySet1((_this).F9(), 4)
					(_nw78).ArraySet1(_this.F10, 5)
					(_nw78).ArraySet1((_this).F7(), 6)
					(_nw78).ArraySet1((_this).F9(), 7)
					(_nw78).ArraySet1(_this.F10, 8)
					(_nw78).ArraySet1((_this).F7(), 9)
					(_nw78).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sdqprjgij")).Cardinality()), 10)
					(_nw78).ArraySet1((_this).F8(), 11)
					(_nw78).ArraySet1((_300_v17).Cardinality(), 12)
					(_nw78).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm0(_303_v18, globalState)), 13)
					(_nw78).ArraySet1(_dafny.IntOfInt64(-826), 14)
					(_nw78).ArraySet1((_this).F9(), 15)
					(_nw78).ArraySet1((_this).F7(), 16)
					(_nw78).ArraySet1((_this).F7(), 17)
					(_nw78).ArraySet1((_305_v20).Cardinality(), 18)
					_306_v21 = _nw78
					var _307_v22 D0
					_ = _307_v22
					_307_v22 = Companion_D0_.Create_DC0_(_306_v21)
					var _pat_let_tv4 = _306_v21
					_ = _pat_let_tv4
					var _308_v23 _dafny.Array
					_ = _308_v23
					var _nwElement0_17 D0 = func(_pat_let2_0 D0) D0 {
						return func(_309_dt__update__tmp_h1 D0) D0 {
							return func(_pat_let3_0 _dafny.Array) D0 {
								return func(_310_dt__update_hcf0_h0 _dafny.Array) D0 {
									return Companion_D0_.Create_DC0_(_310_dt__update_hcf0_h0)
								}(_pat_let3_0)
							}(_pat_let_tv4)
						}(_pat_let2_0)
					}(_307_v22)
					_ = _nwElement0_17
					var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(29))
					_ = _nw79
					(_nw79).ArraySet1(_nwElement0_17, 0)
					(_nw79).ArraySet1(_307_v22, 1)
					(_nw79).ArraySet1(_307_v22, 2)
					(_nw79).ArraySet1(_307_v22, 3)
					(_nw79).ArraySet1(_307_v22, 4)
					(_nw79).ArraySet1(Companion_D0_.Create_DC0_((_307_v22).Dtor_cf0()), 5)
					(_nw79).ArraySet1(_307_v22, 6)
					(_nw79).ArraySet1(_307_v22, 7)
					(_nw79).ArraySet1(_307_v22, 8)
					(_nw79).ArraySet1(_307_v22, 9)
					(_nw79).ArraySet1(_307_v22, 10)
					(_nw79).ArraySet1(_307_v22, 11)
					(_nw79).ArraySet1(Companion_D0_.Create_DC0_(_306_v21), 12)
					(_nw79).ArraySet1(_307_v22, 13)
					(_nw79).ArraySet1(_307_v22, 14)
					(_nw79).ArraySet1(_307_v22, 15)
					(_nw79).ArraySet1(_307_v22, 16)
					(_nw79).ArraySet1(_307_v22, 17)
					(_nw79).ArraySet1(_307_v22, 18)
					(_nw79).ArraySet1(Companion_D0_.Create_DC0_(_306_v21), 19)
					(_nw79).ArraySet1(Companion_D0_.Create_DC0_(_306_v21), 20)
					(_nw79).ArraySet1(_307_v22, 21)
					(_nw79).ArraySet1(_307_v22, 22)
					(_nw79).ArraySet1(_307_v22, 23)
					(_nw79).ArraySet1(_307_v22, 24)
					(_nw79).ArraySet1(_307_v22, 25)
					(_nw79).ArraySet1(_307_v22, 26)
					(_nw79).ArraySet1(_307_v22, 27)
					(_nw79).ArraySet1(_307_v22, 28)
					_308_v23 = _nw79
					_308_v23 = _308_v23
					var _311_v24 D0
					_ = _311_v24
					_311_v24 = Companion_D0_.Create_DC1_(_303_v18)
					var _312_v25 _dafny.Set
					_ = _312_v25
					_312_v25 = _dafny.SetOf(_311_v24, _311_v24)
					var _pat_let_tv5 = _303_v18
					_ = _pat_let_tv5
					var _313_v26 _dafny.Map
					_ = _313_v26
					_313_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v0, (_312_v25).IsProperSubsetOf(_dafny.SetOf(_311_v24, func(_pat_let4_0 D0) D0 {
						return func(_314_dt__update__tmp_h2 D0) D0 {
							return func(_pat_let5_0 _dafny.CodePoint) D0 {
								return func(_315_dt__update_hcf1_h0 _dafny.CodePoint) D0 {
									return Companion_D0_.Create_DC1_(_315_dt__update_hcf1_h0)
								}(_pat_let5_0)
							}(_pat_let_tv5)
						}(_pat_let4_0)
					}(_311_v24))))
					var _316_v27 T0
					_ = _316_v27
					var _nw80 *C1 = New_C1_()
					_ = _nw80
					_nw80.Ctor__(_279_v0, _this.F10, (_this).F7(), (_dafny.Zero).Minus(_this.F10))
					_316_v27 = _nw80
					var _317_v28 D2
					_ = _317_v28
					_317_v28 = Companion_D2_.Create_DC5_((_dafny.Zero).Minus((_this).F8()), Companion_Default___.Fm11(_279_v0, _this.F10, globalState), _316_v27)
					var _318_v29 _dafny.Map
					_ = _318_v29
					_318_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let6_0 D2) D2 {
						return func(_319_dt__update__tmp_h3 D2) D2 {
							return func(_pat_let7_0 _dafny.Int) D2 {
								return func(_320_dt__update_hcf5_h0 _dafny.Int) D2 {
									return Companion_D2_.Create_DC5_(_320_dt__update_hcf5_h0, (_319_dt__update__tmp_h3).Dtor_cf6(), (_319_dt__update__tmp_h3).Dtor_cf7())
								}(_pat_let7_0)
							}(_this.F10)
						}(_pat_let6_0)
					}(Companion_D2_.Create_DC5_((_this).F9(), p0, _316_v27)), _279_v0)
					if (func() bool {
						if (_313_v26).Contains(!(_318_v29).Contains(_317_v28)) {
							return (_313_v26).Get(!(_318_v29).Contains(_317_v28)).(bool)
						}
						return !(_305_v20).Equals(_305_v20)
					})() {
						var _321_v30 _dafny.MultiSet
						_ = _321_v30
						_321_v30 = _dafny.MultiSetOf((_this).F8())
						var _322_v31 _dafny.Set
						_ = _322_v31
						_322_v31 = _dafny.SetOf(_321_v30)
						var _323_v32 D3
						_ = _323_v32
						_323_v32 = Companion_D3_.Create_DC9_(Companion_D3_.Create_DC7_(_322_v31))
						var _324_v33 _dafny.Sequence
						_ = _324_v33
						_324_v33 = _dafny.SeqOf(_323_v32, _323_v32, _323_v32)
						(globalState).F1 = !_dafny.Companion_Sequence_.Contains(_324_v33, _323_v32)
						(globalState).F1 = Companion_Default___.Fm2((func() bool {
							if _279_v0 {
								return !(_279_v0)
							}
							return _279_v0
						})(), globalState)
						var _325_v34 _dafny.Array
						_ = _325_v34
						var _len0_10 _dafny.Int = _dafny.IntOfInt64(3)
						_ = _len0_10
						var _nw81 _dafny.Array
						_ = _nw81
						if _len0_10.Cmp(_dafny.Zero) == 0 {
							_nw81 = _dafny.NewArray(_len0_10)
						} else {
							var _init10 func(_dafny.Int) bool = (func(_326_v27 T0) func(_dafny.Int) bool {
								return func(_327_i3 _dafny.Int) bool {
									return ((_dafny.Zero).Minus((_326_v27).F7())).Cmp((_326_v27).F7()) <= 0
								}
							})(_316_v27)
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
						_325_v34 = _nw81
						var _328_v35 _dafny.Sequence
						_ = _328_v35
						_328_v35 = _dafny.SeqOf(_316_v27)
						var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_325_v34), 0))
						_ = _index59
						(_325_v34).ArraySet1((_dafny.IntOfUint32((_328_v35).Cardinality())).Cmp((_316_v27).F8()) != 0, (_index59).Int())
						var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_325_v34), 0))
						_ = _index60
						(_325_v34).ArraySet1(_279_v0, (_index60).Int())
						var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_306_v21), 0))
						_ = _index61
						(_306_v21).ArraySet1((_316_v27).F8(), (_index61).Int())
						var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_306_v21), 0))
						_ = _index62
						(_306_v21).ArraySet1(Companion_Default___.SafeModuloInt((_this).F9(), (_316_v27).F8()), (_index62).Int())
						(_this).F10 = ((func() _dafny.Int {
							if (_325_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_325_v34), 0))).Int()).(bool) {
								return (_this).F8()
							}
							return Companion_Default___.Fm0(_303_v18, globalState)
						})()).Times((_dafny.Zero).Minus((_316_v27).F8()))
					} else {
						var _329_v36 _dafny.Array
						_ = _329_v36
						var _len0_11 _dafny.Int = _dafny.IntOfInt64(22)
						_ = _len0_11
						var _nw82 _dafny.Array
						_ = _nw82
						if _len0_11.Cmp(_dafny.Zero) == 0 {
							_nw82 = _dafny.NewArray(_len0_11)
						} else {
							var _init11 func(_dafny.Int) bool = (func(_330_v0 bool) func(_dafny.Int) bool {
								return func(_331_i4 _dafny.Int) bool {
									return _330_v0
								}
							})(_279_v0)
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
						_329_v36 = _nw82
						var _332_v37 _dafny.Map
						_ = _332_v37
						_332_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v0, _329_v36)
						_332_v37 = _332_v37
						(globalState).F1 = ((_316_v27).F7()).Cmp((_this).F7()) == 0
						var _333_v38 _dafny.Set
						_ = _333_v38
						_333_v38 = _dafny.SetOf((_this).F9())
						var _334_v39 _dafny.Sequence
						_ = _334_v39
						_334_v39 = _dafny.SeqOf(_dafny.IntOfInt64(-828), (_333_v38).Cardinality(), (_316_v27).F8(), (_316_v27).F7(), _dafny.IntOfUint32((p0).Cardinality()))
						_334_v39 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_334_v39, _334_v39), _334_v39)
						var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_306_v21), 0))
						_ = _index63
						(_306_v21).ArraySet1((_this).F8(), (_index63).Int())
						var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_306_v21), 0))
						_ = _index64
						var _rhs37 _dafny.CodePoint = _303_v18
						_ = _rhs37
						var _rhs38 _dafny.Int = (_316_v27).F8()
						_ = _rhs38
						var _lhs35 _dafny.Array = _306_v21
						_ = _lhs35
						var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_306_v21), 0))
						_ = _lhs36
						_303_v18 = _rhs37
						(_lhs35).ArraySet1(_rhs38, (_lhs36).Int())
						(globalState).F1 = _279_v0
					}
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _335_v40 _dafny.Array
		_ = _335_v40
		var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
		_ = _nw83
		_335_v40 = _nw83
		var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
		_ = _nw84
		_335_v40 = _nw84
		var _336_v41 bool
		_ = _336_v41
		_336_v41 = true
		(globalState).F1 = (_336_v41) || (_336_v41)
		if (func() bool {
			if _336_v41 {
				return false
			}
			return _336_v41
		})() {
			(globalState).F1 = _336_v41
			if _336_v41 {
				var _337_v42 _dafny.Array
				_ = _337_v42
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_12
				var _nw85 _dafny.Array
				_ = _nw85
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw85 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) bool = (func(_338_v41 bool) func(_dafny.Int) bool {
						return func(_339_i5 _dafny.Int) bool {
							return (_338_v41) && (_338_v41)
						}
					})(_336_v41)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw85 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw85).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw85).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_337_v42 = _nw85
				_337_v42 = _337_v42
				var _340_v43 _dafny.Array
				_ = _340_v43
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(2))
				_ = _nw86
				_340_v43 = _nw86
				var _341_v44 _dafny.MultiSet
				_ = _341_v44
				_341_v44 = _dafny.MultiSetOf(_340_v43)
				_341_v44 = ((_dafny.MultiSetOf(_340_v43, _340_v43)).Intersection(_341_v44)).Update(_340_v43, Companion_Default___.Abs(_dafny.IntOfInt64(319)))
				r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cget"), p0)
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_337_v42), 0))
				_ = _index65
				(_337_v42).ArraySet1(_336_v41, (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_337_v42), 0))
				_ = _index66
				var _rhs39 bool = !(_336_v41)
				_ = _rhs39
				var _rhs40 _dafny.Int = _this.F10
				_ = _rhs40
				var _rhs41 _dafny.Int = (_this).F9()
				_ = _rhs41
				var _rhs42 bool = true
				_ = _rhs42
				var _rhs43 _dafny.Int = Companion_Default___.Fm0(_dafny.CodePoint('l'), globalState)
				_ = _rhs43
				var _lhs37 _dafny.Array = _337_v42
				_ = _lhs37
				var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_337_v42), 0))
				_ = _lhs38
				var _lhs39 *C2 = _this
				_ = _lhs39
				var _lhs40 *C2 = _this
				_ = _lhs40
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				var _lhs42 *C2 = _this
				_ = _lhs42
				(_lhs37).ArraySet1(_rhs39, (_lhs38).Int())
				_lhs39.F10 = _rhs40
				_lhs40.F10 = _rhs41
				_lhs41.F1 = _rhs42
				_lhs42.F10 = _rhs43
				var _342_v45 _dafny.Sequence
				_ = _342_v45
				_342_v45 = _dafny.SeqOf((_337_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_337_v42), 0))).Int()).(bool))
				var _343_v46 _dafny.Set
				_ = _343_v46
				_343_v46 = _dafny.SetOf(_342_v45, _342_v45, _342_v45)
				_336_v41 = !((((_343_v46).Intersection(_dafny.SetOf(_342_v45))).Cardinality()).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-139), _dafny.IntOfInt64(-421))) != 0)
			} else {
				(_this).F10 = _dafny.IntOfUint32((p0).Cardinality())
				var _rhs44 _dafny.Int = _this.F10
				_ = _rhs44
				var _rhs45 bool = _336_v41
				_ = _rhs45
				var _lhs43 *C2 = _this
				_ = _lhs43
				var _lhs44 *GlobalState = globalState
				_ = _lhs44
				_lhs43.F10 = _rhs44
				_lhs44.F1 = _rhs45
				(globalState).F1 = _336_v41
				var _344_v47 _dafny.Array
				_ = _344_v47
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw87
				_344_v47 = _nw87
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_344_v47), 0))
				_ = _index67
				(_344_v47).ArraySet1(_336_v41, (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_344_v47), 0))
				_ = _index68
				(_344_v47).ArraySet1(_336_v41, (_index68).Int())
				var _345_v48 *C0
				_ = _345_v48
				var _nw88 *C0 = New_C0_()
				_ = _nw88
				_nw88.Ctor__(_336_v41)
				_345_v48 = _nw88
			}
			var _346_v49 _dafny.Sequence
			_ = _346_v49
			_346_v49 = _dafny.SeqOf(_336_v41)
			var _347_v50 _dafny.Map
			_ = _347_v50
			_347_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_336_v41, globalState), _346_v49)
			var _348_v51 _dafny.Sequence
			_ = _348_v51
			_348_v51 = _dafny.SeqOf(_dafny.IntOfUint32((p0).Cardinality()), (_this).F7(), _this.F10, _this.F10, (_this).F8())
			var _349_v52 T0
			_ = _349_v52
			var _nw89 *C1 = New_C1_()
			_ = _nw89
			_nw89.Ctor__(_336_v41, (_this).F8(), (_this).F9(), _dafny.IntOfUint32((_346_v49).Cardinality()))
			_349_v52 = _nw89
			var _350_v53 _dafny.Map
			_ = _350_v53
			_350_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v41, _349_v52)
			var _351_v54 _dafny.MultiSet
			_ = _351_v54
			_351_v54 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq(_346_v49)).Cardinality(), _dafny.IntOfUint32((_346_v49).Cardinality()))
			var _352_v55 _dafny.CodePoint
			_ = _352_v55
			_352_v55 = _dafny.CodePoint('y')
			var _353_v56 _dafny.Array
			_ = _353_v56
			var _nwElement0_18 _dafny.Int = _this.F10
			_ = _nwElement0_18
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(29))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_18, 0)
			(_nw90).ArraySet1((_dafny.Zero).Minus(_this.F10), 1)
			(_nw90).ArraySet1((_dafny.IntOfUint32((_346_v49).Cardinality())).Plus((_this).F8()), 2)
			(_nw90).ArraySet1((_this).F7(), 3)
			(_nw90).ArraySet1((_347_v50).Cardinality(), 4)
			(_nw90).ArraySet1(_this.F10, 5)
			(_nw90).ArraySet1((_dafny.Zero).Minus((_this).F8()), 6)
			(_nw90).ArraySet1((_this).F7(), 7)
			(_nw90).ArraySet1((_this).F7(), 8)
			(_nw90).ArraySet1((func() _dafny.Int {
				if _336_v41 {
					return _dafny.IntOfUint32((_348_v51).Cardinality())
				}
				return (_this).F8()
			})(), 9)
			(_nw90).ArraySet1((_this).Fm7(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(389))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_354_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})), _dafny.IntOfInt64(454), globalState), 10)
			(_nw90).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F7()), (_this).F9()), 11)
			(_nw90).ArraySet1(_this.F10, 12)
			(_nw90).ArraySet1(((_dafny.Zero).Minus((_350_v53).Cardinality())).Minus((_this).F7()), 13)
			(_nw90).ArraySet1((_349_v52).F8(), 14)
			(_nw90).ArraySet1((_this).F9(), 15)
			(_nw90).ArraySet1((_349_v52).F8(), 16)
			(_nw90).ArraySet1((_this).F7(), 17)
			(_nw90).ArraySet1((_this).F8(), 18)
			(_nw90).ArraySet1(_this.F10, 19)
			(_nw90).ArraySet1((_dafny.Zero).Minus((_this).F7()), 20)
			(_nw90).ArraySet1((_this).F9(), 21)
			(_nw90).ArraySet1((_this).F7(), 22)
			(_nw90).ArraySet1((_dafny.Zero).Minus(((_this).F8()).Minus(_dafny.IntOfInt64(-46))), 23)
			(_nw90).ArraySet1((_349_v52).F8(), 24)
			(_nw90).ArraySet1(((func() _dafny.Int {
				if (_351_v54).Contains((_this).F9()) {
					return (_351_v54).Multiplicity((_this).F9())
				}
				return (_this).F8()
			})()).Plus((_this).F9()), 25)
			(_nw90).ArraySet1(_this.F10, 26)
			(_nw90).ArraySet1(Companion_Default___.Fm0(_352_v55, globalState), 27)
			(_nw90).ArraySet1(_this.F10, 28)
			_353_v56 = _nw90
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_353_v56), 0))
			_ = _index69
			(_353_v56).ArraySet1(_dafny.IntOfUint32((_348_v51).Cardinality()), (_index69).Int())
			var _355_v57 _dafny.Map
			_ = _355_v57
			_355_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_352_v55, _336_v41)
			var _356_v58 _dafny.Sequence
			_ = _356_v58
			_356_v58 = _dafny.SeqOf(_349_v52, _349_v52)
			var _357_v59 _dafny.Map
			_ = _357_v59
			_357_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(516))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_358_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_359_i7 _dafny.Int) _dafny.CodePoint {
					return _358_v55
				}
			})(_352_v55)))).Cardinality()), _356_v58)
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_353_v56), 0))
			_ = _index70
			(_353_v56).ArraySet1(((_355_v57).Merge(Companion_Default___.Fm13(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_357_v59).Contains(_this.F10) {
					return (_357_v59).Get(_this.F10).(_dafny.Sequence)
				}
				return _356_v58
			})()).Cardinality()), globalState))).Cardinality(), (_index70).Int())
			var _360_v60 _dafny.Map
			_ = _360_v60
			_360_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_336_v41), _346_v49)
			var _361_v61 *C1
			_ = _361_v61
			var _nw91 *C1 = New_C1_()
			_ = _nw91
			_nw91.Ctor__(_336_v41, (_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality())), _this.F10, (_360_v60).Cardinality())
			_361_v61 = _nw91
			var _362_v62 D1
			_ = _362_v62
			_362_v62 = Companion_D1_.Create_DC2_(_336_v41)
			var _pat_let_tv6 = _336_v41
			_ = _pat_let_tv6
			var _363_v63 _dafny.Set
			_ = _363_v63
			_363_v63 = _dafny.SetOf(func(_pat_let8_0 D1) D1 {
				return func(_364_dt__update__tmp_h4 D1) D1 {
					return func(_pat_let9_0 bool) D1 {
						return func(_365_dt__update_hcf2_h0 bool) D1 {
							return Companion_D1_.Create_DC2_(_365_dt__update_hcf2_h0)
						}(_pat_let9_0)
					}(_pat_let_tv6)
				}(_pat_let8_0)
			}(_362_v62))
			(globalState).F3 = (_363_v63).IsProperSubsetOf(Companion_Default___.Fm14(_361_v61.F13, !((_361_v61).F12()), globalState))
		} else {
			if _dafny.Companion_Sequence_.IsPrefixOf(p0, p0) {
				var _366_v64 _dafny.CodePoint
				_ = _366_v64
				_366_v64 = _dafny.CodePoint('s')
				var _367_v65 _dafny.Map
				_ = _367_v65
				_367_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v41, _366_v64)
				var _368_v66 _dafny.Map
				_ = _368_v66
				_368_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), (func() _dafny.CodePoint {
					if (_367_v65).Contains(_336_v41) {
						return (_367_v65).Get(_336_v41).(_dafny.CodePoint)
					}
					return _366_v64
				})())
				var _369_v67 _dafny.Map
				_ = _369_v67
				_369_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_v66, (_this).F9())
				var _370_v68 _dafny.Sequence
				_ = _370_v68
				_370_v68 = _dafny.SeqOf(_336_v41)
				var _371_v69 *C1
				_ = _371_v69
				var _nw92 *C1 = New_C1_()
				_ = _nw92
				_nw92.Ctor__(!(_336_v41), (func() _dafny.Int {
					if (_369_v67).Contains(_368_v66) {
						return (_369_v67).Get(_368_v66).(_dafny.Int)
					}
					return (Companion_Default___.Fm15(globalState)).Cardinality()
				})(), (_this).F9(), _dafny.IntOfUint32((_370_v68).Cardinality()))
				_371_v69 = _nw92
				var _372_v70 _dafny.Sequence
				_ = _372_v70
				var _373_v71 D3
				_ = _373_v71
				var _out3 _dafny.Sequence
				_ = _out3
				var _out4 D3
				_ = _out4
				_out3, _out4 = (_371_v69).M6(_371_v69.F13, globalState)
				_372_v70 = _out3
				_373_v71 = _out4
				var _374_v73 *C0
				_ = _374_v73
				var _nw93 *C0 = New_C0_()
				_ = _nw93
				_nw93.Ctor__(!(_336_v41))
				_374_v73 = _nw93
				var _375_v74 _dafny.MultiSet
				_ = _375_v74
				_375_v74 = _dafny.MultiSetOf(_374_v73, _374_v73)
				var _376_v75 _dafny.Map
				_ = _376_v75
				_376_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm7(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("glko"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("glko")).Cardinality()))).Uint32(), _dafny.CodePoint('o')), _dafny.IntOfInt64(-128), globalState), _375_v74)
				var _377_v76 _dafny.Set
				_ = _377_v76
				_377_v76 = _dafny.SetOf((_371_v69).F12())
				var _378_v77 _dafny.Array
				_ = _378_v77
				var _nwElement0_19 bool = (_371_v69).F12()
				_ = _nwElement0_19
				var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(19))
				_ = _nw94
				(_nw94).ArraySet1(_nwElement0_19, 0)
				(_nw94).ArraySet1(_336_v41, 1)
				(_nw94).ArraySet1((_371_v69).F12(), 2)
				(_nw94).ArraySet1(_336_v41, 3)
				(_nw94).ArraySet1((_371_v69).F12(), 4)
				(_nw94).ArraySet1((_374_v73).F11(), 5)
				(_nw94).ArraySet1(_336_v41, 6)
				(_nw94).ArraySet1(_336_v41, 7)
				(_nw94).ArraySet1(_336_v41, 8)
				(_nw94).ArraySet1((_371_v69).F12(), 9)
				(_nw94).ArraySet1(!((_371_v69).F12()), 10)
				(_nw94).ArraySet1((_371_v69).F12(), 11)
				(_nw94).ArraySet1(_336_v41, 12)
				(_nw94).ArraySet1(!((_374_v73).F11()), 13)
				(_nw94).ArraySet1((_371_v69).F12(), 14)
				(_nw94).ArraySet1((_374_v73).F11(), 15)
				(_nw94).ArraySet1(false, 16)
				(_nw94).ArraySet1((_371_v69).F12(), 17)
				(_nw94).ArraySet1(_336_v41, 18)
				_378_v77 = _nw94
				var _379_v78 _dafny.Sequence
				_ = _379_v78
				_379_v78 = _dafny.SeqOf(_378_v77)
				var _380_v79 _dafny.Map
				_ = _380_v79
				_380_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_374_v73).F11())
				var _381_v80 _dafny.Array
				_ = _381_v80
				var _nwElement0_20 _dafny.Int = (_dafny.IntOfUint32((p0).Cardinality())).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}((func(_382_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_383_i8 _dafny.Int) _dafny.CodePoint {
						return _382_v64
					}
				})(_366_v64)))).Cardinality()))
				_ = _nwElement0_20
				var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(19))
				_ = _nw95
				(_nw95).ArraySet1(_nwElement0_20, 0)
				(_nw95).ArraySet1((((func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter17 := _dafny.Iterate((_dafny.SeqOf((_this).F7())).Elements()); ; {
						_compr_17, _ok17 := _iter17()
						if !_ok17 {
							break
						}
						var _384_v72 _dafny.Int
						_384_v72 = interface{}(_compr_17).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_this).F7()), _384_v72) {
							_coll17.Add(Companion_Default___.SafeDivisionInt(_384_v72, _dafny.IntOfInt64(950)), _this.F10)
						}
					}
					return _coll17.ToMap()
				}()).Update(_371_v69.F13, _371_v69.F13)).Cardinality()).Times((_this).F7()), 1)
				(_nw95).ArraySet1((_this).F9(), 2)
				(_nw95).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
					if (_376_v75).Contains((_this).F8()) {
						return (_376_v75).Get((_this).F8()).(_dafny.MultiSet)
					}
					return _375_v74
				})(), (_this).F9())).Cardinality(), 3)
				(_nw95).ArraySet1((func() _dafny.Int {
					if _336_v41 {
						return (_this).F7()
					}
					return _this.F10
				})(), 4)
				(_nw95).ArraySet1((_this).F9(), 5)
				(_nw95).ArraySet1((_dafny.IntOfInt64(76)).Minus((_this).F9()), 6)
				(_nw95).ArraySet1(((_this).F7()).Minus(_this.F10), 7)
				(_nw95).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()), 8)
				(_nw95).ArraySet1((_dafny.Zero).Minus(_371_v69.F13), 9)
				(_nw95).ArraySet1(_dafny.IntOfInt64(-325), 10)
				(_nw95).ArraySet1((_this).F9(), 11)
				(_nw95).ArraySet1(((_this).F8()).Times((_377_v76).Cardinality()), 12)
				(_nw95).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F10)), _371_v69.F13), 13)
				(_nw95).ArraySet1(_371_v69.F13, 14)
				(_nw95).ArraySet1((_this).F7(), 15)
				(_nw95).ArraySet1((_this).F8(), 16)
				(_nw95).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_this.F10, (_this).F8())).Cardinality())).Times(_dafny.IntOfUint32((_379_v78).Cardinality())), 17)
				(_nw95).ArraySet1(((_380_v79).Merge(_380_v79)).Cardinality(), 18)
				_381_v80 = _nw95
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_381_v80), 0))
				_ = _index71
				(_381_v80).ArraySet1(_dafny.IntOfInt64(793), (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_381_v80), 0))
				_ = _index72
				(_381_v80).ArraySet1((_377_v76).Cardinality(), (_index72).Int())
				(globalState).F1 = _336_v41
				(_this).F10 = (_dafny.Zero).Minus(((_381_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_381_v80), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(-91)))
			} else {
				(globalState).F3 = (!(_336_v41) || (_336_v41)) || (_336_v41)
				var _385_v81 _dafny.CodePoint
				_ = _385_v81
				_385_v81 = _dafny.CodePoint('y')
				r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(495))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}(func(_386_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})), p0), (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(495))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}(func(_387_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})), p0)).Cardinality()))).Uint32(), _385_v81), p0)
				(globalState).F1 = false
				var _388_v82 _dafny.Array
				_ = _388_v82
				var _nwElement0_21 _dafny.Int = (_this).F8()
				_ = _nwElement0_21
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(8))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_21, 0)
				(_nw96).ArraySet1((_this).F9(), 1)
				(_nw96).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}(func(_389_i10 _dafny.Int) _dafny.Int {
					return (_this).F7()
				}))).Cardinality()), 2)
				(_nw96).ArraySet1((_this).F8(), 3)
				(_nw96).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf((_this).F7())).Cardinality())).Cardinality()), 4)
				(_nw96).ArraySet1(_dafny.IntOfInt64(404), 5)
				(_nw96).ArraySet1(_this.F10, 6)
				(_nw96).ArraySet1((_this).F8(), 7)
				_388_v82 = _nw96
				var _390_v83 _dafny.Sequence
				_ = _390_v83
				_390_v83 = _dafny.SeqOf(_388_v82, _388_v82)
				var _391_v84 _dafny.Array
				_ = _391_v84
				var _nwElement0_22 bool = _336_v41
				_ = _nwElement0_22
				var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(8))
				_ = _nw97
				(_nw97).ArraySet1(_nwElement0_22, 0)
				(_nw97).ArraySet1(_336_v41, 1)
				(_nw97).ArraySet1(false, 2)
				(_nw97).ArraySet1(_336_v41, 3)
				(_nw97).ArraySet1(_336_v41, 4)
				(_nw97).ArraySet1(_336_v41, 5)
				(_nw97).ArraySet1(_336_v41, 6)
				(_nw97).ArraySet1(_336_v41, 7)
				_391_v84 = _nw97
				var _392_v85 _dafny.Map
				_ = _392_v85
				_392_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_390_v83).Cardinality()), _391_v84)
				var _393_v86 _dafny.Sequence
				_ = _393_v86
				_393_v86 = _dafny.SeqOf((_392_v85).Cardinality(), (_this).F9(), (_this).F8())
				var _394_v87 _dafny.Sequence
				_ = _394_v87
				_394_v87 = _dafny.SeqOf((_393_v86).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_393_v86).Cardinality()))).Uint32()).(_dafny.Int), ((_this).F7()).Minus((_this).F7()))
				var _395_v88 _dafny.MultiSet
				_ = _395_v88
				_395_v88 = _dafny.MultiSetOf(Companion_Default___.Fm16(globalState), Companion_Default___.Fm16(globalState), _393_v86, _394_v87)
				var _396_v89 _dafny.Map
				_ = _396_v89
				_396_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_395_v88, (_this).F7())
				var _rhs46 bool = !(_336_v41)
				_ = _rhs46
				var _rhs47 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_this.F10, _this.F10, (func() _dafny.Int {
					if (_396_v89).Contains(_395_v88) {
						return (_396_v89).Get(_395_v88).(_dafny.Int)
					}
					return _this.F10
				})()), (Companion_Default___.SafeIndex((_dafny.IntOfInt64(498)).Plus((_this).F8()), _dafny.IntOfUint32((_dafny.SeqOf(_this.F10, _this.F10, (func() _dafny.Int {
					if (_396_v89).Contains(_395_v88) {
						return (_396_v89).Get(_395_v88).(_dafny.Int)
					}
					return _this.F10
				})())).Cardinality()))).Uint32(), (_this).F9())
				_ = _rhs47
				_336_v41 = _rhs46
				_394_v87 = _rhs47
				r0 = p0
			}
			if _336_v41 {
				var _397_v90 D2
				_ = _397_v90
				_397_v90 = Companion_D2_.Create_DC4_(_this.F10)
				var _398_v91 _dafny.Set
				_ = _398_v91
				_398_v91 = _dafny.SetOf(_336_v41)
				var _399_v92 _dafny.Array
				_ = _399_v92
				var _nwElement0_23 _dafny.Int = (_this).F8()
				_ = _nwElement0_23
				var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(4))
				_ = _nw98
				(_nw98).ArraySet1(_nwElement0_23, 0)
				(_nw98).ArraySet1(Companion_Default___.SafeModuloInt((Companion_Default___.Fm17(_336_v41, _this.F10, globalState)).Cardinality(), (_this).F8()), 1)
				(_nw98).ArraySet1((_397_v90).Dtor_cf4(), 2)
				(_nw98).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_this.F10, (_398_v91).Cardinality())), 3)
				_399_v92 = _nw98
				var _rhs48 _dafny.Array = _399_v92
				_ = _rhs48
				var _rhs49 _dafny.Int = _this.F10
				_ = _rhs49
				var _lhs45 *C2 = _this
				_ = _lhs45
				_399_v92 = _rhs48
				_lhs45.F10 = _rhs49
				(globalState).F1 = ((_dafny.SetOf(_336_v41)).Intersection(_398_v91)).IsSubsetOf(_398_v91)
				(globalState).F3 = _336_v41
				_397_v90 = _397_v90
				(globalState).F1 = _336_v41
			} else {
				var _400_v93 _dafny.Sequence
				_ = _400_v93
				_400_v93 = _dafny.SeqOf(_dafny.IntOfUint32((p0).Cardinality()))
				var _401_v94 _dafny.Set
				_ = _401_v94
				_401_v94 = _dafny.SetOf(p0)
				var _402_v95 _dafny.Sequence
				_ = _402_v95
				_402_v95 = _dafny.SeqOf(_336_v41)
				var _403_v96 _dafny.Array
				_ = _403_v96
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw99
				_403_v96 = _nw99
				var _404_v97 _dafny.Map
				_ = _404_v97
				_404_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v41, _403_v96)
				var _405_v98 _dafny.Sequence
				_ = _405_v98
				_405_v98 = _dafny.SeqOf(_403_v96)
				var _406_v99 _dafny.MultiSet
				_ = _406_v99
				_406_v99 = _dafny.MultiSetOf(_404_v97, (_404_v97).Update(_336_v41, (_405_v98).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-514), _dafny.IntOfUint32((_405_v98).Cardinality()))).Uint32()).(_dafny.Array)), _404_v97, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _403_v96), _404_v97)
				var _407_v100 _dafny.Array
				_ = _407_v100
				var _nwElement0_24 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm16(globalState), _400_v93)
				_ = _nwElement0_24
				var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(16))
				_ = _nw100
				(_nw100).ArraySet1(_nwElement0_24, 0)
				(_nw100).ArraySet1(_336_v41, 1)
				(_nw100).ArraySet1(_336_v41, 2)
				(_nw100).ArraySet1(_336_v41, 3)
				(_nw100).ArraySet1((_dafny.SetOf(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(93))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}(func(_408_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				})))).IsSubsetOf(_401_v94), 4)
				(_nw100).ArraySet1(((_this).F7()).Cmp((_this).F8()) != 0, 5)
				(_nw100).ArraySet1((func() bool {
					if _336_v41 {
						return _336_v41
					}
					return true
				})(), 6)
				(_nw100).ArraySet1((_336_v41) && (_336_v41), 7)
				(_nw100).ArraySet1(_336_v41, 8)
				(_nw100).ArraySet1((func() bool {
					if Companion_Default___.Fm2(_336_v41, globalState) {
						return _336_v41
					}
					return _336_v41
				})(), 9)
				(_nw100).ArraySet1(_336_v41, 10)
				(_nw100).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_402_v95, _402_v95), 11)
				(_nw100).ArraySet1(_336_v41, 12)
				(_nw100).ArraySet1(((func() _dafny.Int {
					if (_406_v99).Contains(_404_v97) {
						return (_406_v99).Multiplicity(_404_v97)
					}
					return (_this).F8()
				})()).Cmp((_this).F9()) > 0, 13)
				(_nw100).ArraySet1(_336_v41, 14)
				(_nw100).ArraySet1(_336_v41, 15)
				_407_v100 = _nw100
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_403_v96), 0))
				_ = _index73
				(_403_v96).ArraySet1(!(_336_v41), (_index73).Int())
				var _409_v101 _dafny.Map
				_ = _409_v101
				_409_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(p0, p0), (_this).F7())
				var _410_v102 _dafny.CodePoint
				_ = _410_v102
				_410_v102 = _dafny.CodePoint('l')
				var _411_v103 _dafny.Array
				_ = _411_v103
				var _nwElement0_25 _dafny.CodePoint = _410_v102
				_ = _nwElement0_25
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(10))
				_ = _nw101
				(_nw101).ArraySet1CodePoint(_nwElement0_25, 0)
				(_nw101).ArraySet1CodePoint(_410_v102, 1)
				(_nw101).ArraySet1CodePoint(_410_v102, 2)
				(_nw101).ArraySet1CodePoint(_410_v102, 3)
				(_nw101).ArraySet1CodePoint(_410_v102, 4)
				(_nw101).ArraySet1CodePoint(_410_v102, 5)
				(_nw101).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _336_v41 {
						return _410_v102
					}
					return _410_v102
				})(), 6)
				(_nw101).ArraySet1CodePoint(_410_v102, 7)
				(_nw101).ArraySet1CodePoint(Companion_Default___.Fm3(_336_v41, globalState), 8)
				(_nw101).ArraySet1CodePoint(_410_v102, 9)
				_411_v103 = _nw101
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_403_v96), 0))
				_ = _index74
				var _rhs50 bool = _336_v41
				_ = _rhs50
				var _rhs51 _dafny.Map = _409_v101
				_ = _rhs51
				var _rhs52 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F8())).Cardinality()
				_ = _rhs52
				var _rhs53 _dafny.Array = _411_v103
				_ = _rhs53
				var _rhs54 _dafny.Array = _403_v96
				_ = _rhs54
				var _lhs46 _dafny.Array = _403_v96
				_ = _lhs46
				var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_403_v96), 0))
				_ = _lhs47
				var _lhs48 *C2 = _this
				_ = _lhs48
				(_lhs46).ArraySet1(_rhs50, (_lhs47).Int())
				_409_v101 = _rhs51
				_lhs48.F10 = _rhs52
				_411_v103 = _rhs53
				_403_v96 = _rhs54
				var _412_v104 _dafny.Set
				_ = _412_v104
				_412_v104 = _dafny.SetOf(false, false)
				var _413_v105 *C1
				_ = _413_v105
				var _nw102 *C1 = New_C1_()
				_ = _nw102
				_nw102.Ctor__(true, (_412_v104).Cardinality(), ((_this).F7()).Minus(_this.F10), (func() _dafny.Int {
					if !(_336_v41) {
						return (_this).F9()
					}
					return (_dafny.Zero).Minus((_this).F8())
				})())
				_413_v105 = _nw102
				var _414_v106 _dafny.Array
				_ = _414_v106
				var _nwElement0_26 _dafny.Set = _412_v104
				_ = _nwElement0_26
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(15))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_26, 0)
				(_nw103).ArraySet1((_412_v104).Union(_412_v104), 1)
				(_nw103).ArraySet1(_412_v104, 2)
				(_nw103).ArraySet1(_412_v104, 3)
				(_nw103).ArraySet1(_412_v104, 4)
				(_nw103).ArraySet1((func() _dafny.Set {
					if (_403_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_403_v96), 0))).Int()).(bool) {
						return _412_v104
					}
					return _412_v104
				})(), 5)
				(_nw103).ArraySet1((_412_v104).Difference(_412_v104), 6)
				(_nw103).ArraySet1(_412_v104, 7)
				(_nw103).ArraySet1((func() _dafny.Set {
					if (_403_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_403_v96), 0))).Int()).(bool) {
						return _412_v104
					}
					return _412_v104
				})(), 8)
				(_nw103).ArraySet1(_412_v104, 9)
				(_nw103).ArraySet1(_dafny.SetOf((_413_v105).F12()), 10)
				(_nw103).ArraySet1(_dafny.SetOf((_413_v105).F12(), _336_v41, (_413_v105).F12(), false, _336_v41), 11)
				(_nw103).ArraySet1(_412_v104, 12)
				(_nw103).ArraySet1((_412_v104).Difference(_412_v104), 13)
				(_nw103).ArraySet1(_412_v104, 14)
				_414_v106 = _nw103
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_414_v106), 0))
				_ = _index75
				(_414_v106).ArraySet1(_412_v104, (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_414_v106), 0))
				_ = _index76
				var _rhs55 _dafny.Set = _412_v104
				_ = _rhs55
				var _rhs56 _dafny.Array = _407_v100
				_ = _rhs56
				var _rhs57 _dafny.CodePoint = (p0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}((func(_415_v105 *C1) func(_dafny.Int) _dafny.Int {
					return func(_416_i12 _dafny.Int) _dafny.Int {
						return _415_v105.F13
					}
				})(_413_v105)))).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs57
				var _lhs49 _dafny.Array = _414_v106
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_414_v106), 0))
				_ = _lhs50
				(_lhs49).ArraySet1(_rhs55, (_lhs50).Int())
				_407_v100 = _rhs56
				_410_v102 = _rhs57
				r0 = p0
				(globalState).F1 = false
			}
			(_this).F10 = _this.F10
			var _417_v107 _dafny.Array
			_ = _417_v107
			var _nw104 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw104
			_417_v107 = _nw104
			var _418_v108 _dafny.CodePoint
			_ = _418_v108
			_418_v108 = _dafny.CodePoint('i')
			var _419_v109 D4
			_ = _419_v109
			_419_v109 = Companion_D4_.Create_DC11_((_this).F7(), _418_v108)
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_417_v107), 0))
			_ = _index77
			(_417_v107).ArraySet1(((_419_v109).Dtor_cf13()).Cmp(_dafny.IntOfInt64(284)) != 0, (_index77).Int())
			var _420_v110 _dafny.MultiSet
			_ = _420_v110
			_420_v110 = _dafny.MultiSetOf((_this).Fm9(_336_v41, _this.F10, globalState), _336_v41)
			var _421_v111 _dafny.Sequence
			_ = _421_v111
			_421_v111 = _dafny.SeqOf(_336_v41, _336_v41)
			var _422_v112 _dafny.Sequence
			_ = _422_v112
			_422_v112 = _dafny.SeqOf(_420_v110, _dafny.MultiSetOf(_336_v41))
			var _423_v113 _dafny.Array
			_ = _423_v113
			var _nwElement0_27 _dafny.MultiSet = _420_v110
			_ = _nwElement0_27
			var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(12))
			_ = _nw105
			(_nw105).ArraySet1(_nwElement0_27, 0)
			(_nw105).ArraySet1((_420_v110).Intersection(_420_v110), 1)
			(_nw105).ArraySet1(_420_v110, 2)
			(_nw105).ArraySet1(_420_v110, 3)
			(_nw105).ArraySet1((_dafny.MultiSetFromSeq(_421_v111)).Update(_336_v41, Companion_Default___.Abs((_this).F9())), 4)
			(_nw105).ArraySet1((func() _dafny.MultiSet {
				if _336_v41 {
					return _420_v110
				}
				return _420_v110
			})(), 5)
			(_nw105).ArraySet1(_420_v110, 6)
			(_nw105).ArraySet1(_420_v110, 7)
			(_nw105).ArraySet1(_420_v110, 8)
			(_nw105).ArraySet1(_420_v110, 9)
			(_nw105).ArraySet1(_dafny.MultiSetOf(_336_v41, _336_v41, _336_v41), 10)
			(_nw105).ArraySet1((_422_v112).Select((Companion_Default___.SafeIndex((_this).F7(), _dafny.IntOfUint32((_422_v112).Cardinality()))).Uint32()).(_dafny.MultiSet), 11)
			_423_v113 = _nw105
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_423_v113), 0))
			_ = _index78
			(_423_v113).ArraySet1(_dafny.MultiSetOf(_336_v41), (_index78).Int())
			var _424_v114 _dafny.Set
			_ = _424_v114
			_424_v114 = _dafny.SetOf(_418_v108, _418_v108)
			var _425_v115 _dafny.Map
			_ = _425_v115
			_425_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_424_v114, _this.F10)
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_417_v107), 0))
			_ = _index79
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_423_v113), 0))
			_ = _index80
			var _rhs58 bool = (func() bool {
				if true {
					return _336_v41
				}
				return !((_425_v115).Equals(_425_v115))
			})()
			_ = _rhs58
			var _rhs59 _dafny.MultiSet = _420_v110
			_ = _rhs59
			var _rhs60 _dafny.Int = Companion_Default___.SafeDivisionInt(((_this).F8()).Times((_dafny.Zero).Minus(_this.F10)), (_this).F7())
			_ = _rhs60
			var _lhs51 _dafny.Array = _417_v107
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_417_v107), 0))
			_ = _lhs52
			var _lhs53 _dafny.Array = _423_v113
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_423_v113), 0))
			_ = _lhs54
			var _lhs55 *C2 = _this
			_ = _lhs55
			(_lhs51).ArraySet1(_rhs58, (_lhs52).Int())
			(_lhs53).ArraySet1(_rhs59, (_lhs54).Int())
			_lhs55.F10 = _rhs60
			var _426_v116 _dafny.Set
			_ = _426_v116
			_426_v116 = _dafny.SetOf((_this).F9(), _dafny.IntOfInt64(658), _dafny.IntOfInt64(163))
			var _427_v117 _dafny.Map
			_ = _427_v117
			_427_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_426_v116, true)
			var _428_v118 _dafny.Map
			_ = _428_v118
			_428_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_417_v107).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_417_v107), 0))).Int()).(bool), _427_v117)
			var _429_v119 _dafny.Map
			_ = _429_v119
			_429_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if (_428_v118).Contains(_336_v41) {
					return (_428_v118).Get(_336_v41).(_dafny.Map)
				}
				return _427_v117
			})(), _336_v41)
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_417_v107), 0))
			_ = _index81
			var _rhs61 bool = (_417_v107).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_417_v107), 0))).Int()).(bool)
			_ = _rhs61
			var _rhs62 _dafny.Array = (func() _dafny.Array {
				if _336_v41 {
					return _417_v107
				}
				return _417_v107
			})()
			_ = _rhs62
			var _rhs63 bool = false
			_ = _rhs63
			var _rhs64 bool = (func() bool {
				if (_429_v119).Contains(_427_v117) {
					return (_429_v119).Get(_427_v117).(bool)
				}
				return (_417_v107).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_417_v107), 0))).Int()).(bool)
			})()
			_ = _rhs64
			var _lhs56 _dafny.Array = _417_v107
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_417_v107), 0))
			_ = _lhs57
			_336_v41 = _rhs61
			_417_v107 = _rhs62
			_336_v41 = _rhs63
			(_lhs56).ArraySet1(_rhs64, (_lhs57).Int())
		}
		var _430_v120 _dafny.Sequence
		_ = _430_v120
		_430_v120 = _dafny.SeqOf((_this).F8())
		var _431_v121 _dafny.CodePoint
		_ = _431_v121
		_431_v121 = _dafny.CodePoint('q')
		r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_430_v120).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _431_v121), _dafny.UnicodeSeqOfUtf8Bytes("vhgenyj"))
		r0 = p0
		return r0
	}
}
func (_this *C2) F9() _dafny.Int {
	{
		return _this._f9
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f7 _dafny.Int
	_f8 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
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
func (_this *C3) F8() _dafny.Int {
	return _this._f8
}
func (_this *C3) Ctor__(f7 _dafny.Int, f8 _dafny.Int) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C3) Fm6(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	{
		return (_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_this).F7(), (_dafny.Zero).Minus((_this).F8())), _dafny.MultiSetOf((_this).F8(), (_this).F8()), _dafny.MultiSetOf((_this).F7(), (_this).F7(), (_this).F8(), (_this).F7()))).Difference(_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qt"))).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(670))))
	}
}
func (_this *C3) Fm7(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('e'))).Union((_dafny.MultiSetOf(_dafny.CodePoint('g'), _dafny.CodePoint('b'))).Union(_dafny.MultiSetOf(_dafny.CodePoint('q'), _dafny.CodePoint('w'))))).Cardinality()
	}
}
func (_this *C3) M3(p0 _dafny.Sequence, p1 bool, p2 D1, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		r2 = Companion_Default___.SafeDivisionInt(p3, (_this).F8())
		var _432_v0 _dafny.Array
		_ = _432_v0
		var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
		_ = _nw106
		_432_v0 = _nw106
		_432_v0 = _432_v0
		var _433_v1 _dafny.Sequence
		_ = _433_v1
		_433_v1 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F8(), _dafny.IntOfUint32((p0).Cardinality()))).Cardinality()), ((_this).F8()).Times(p3))
		var _rhs65 _dafny.Sequence = _433_v1
		_ = _rhs65
		var _rhs66 bool = p1
		_ = _rhs66
		_433_v1 = _rhs65
		r3 = _rhs66
		var _434_v2 _dafny.Map
		_ = _434_v2
		_434_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality())))
		_434_v2 = (_434_v2).Update((_dafny.Zero).Minus((_this).F8()), (_this).F7())
		var _435_v3 _dafny.Sequence
		_ = _435_v3
		_435_v3 = _dafny.SeqOf(!(p1), true)
		var _436_v4 _dafny.MultiSet
		_ = _436_v4
		_436_v4 = _dafny.MultiSetOf(Companion_D1_.Create_DC2_(p1), p2)
		_435_v3 = Companion_Default___.Fm8(_433_v1, _dafny.IntOfInt64(-163), p1, _436_v4, globalState)
		var _437_v5 _dafny.CodePoint
		_ = _437_v5
		_437_v5 = _dafny.CodePoint('c')
		var _438_v6 D1
		_ = _438_v6
		_438_v6 = Companion_D1_.Create_DC3_(_437_v5)
		var _source4 D1 = (func() D1 {
			if Companion_Default___.Fm2(p1, globalState) {
				return func(_pat_let10_0 D1) D1 {
					return func(_439_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let11_0 _dafny.CodePoint) D1 {
							return func(_440_dt__update_hcf3_h0 _dafny.CodePoint) D1 {
								return Companion_D1_.Create_DC3_(_440_dt__update_hcf3_h0)
							}(_pat_let11_0)
						}(_dafny.CodePoint('f'))
					}(_pat_let10_0)
				}(_438_v6)
			}
			return _438_v6
		})()
		_ = _source4
		if _source4.Is_DC3() {
			var _441___mcc_h0 _dafny.CodePoint = _source4.Get_().(D1_DC3).Cf3
			_ = _441___mcc_h0
			var _442_cf3 _dafny.CodePoint = _441___mcc_h0
			_ = _442_cf3
			var _443_v7 _dafny.MultiSet
			_ = _443_v7
			_443_v7 = _dafny.MultiSetOf(p1)
			var _444_v8 _dafny.Sequence
			_ = _444_v8
			_444_v8 = _dafny.SeqOf(_443_v7, _443_v7)
			_444_v8 = _444_v8
			if p1 {
				var _445_v9 *C2
				_ = _445_v9
				var _nw107 *C2 = New_C2_()
				_ = _nw107
				_nw107.Ctor__(p3, p3, Companion_Default___.Fm0(Companion_Default___.Fm3(false, globalState), globalState), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p3), (_this).F7()))
				_445_v9 = _nw107
				(_445_v9).F10 = ((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_435_v3, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_435_v3).Cardinality()))).Uint32(), p1)).Cardinality()))).Times((func() _dafny.Int {
					if p1 {
						return _dafny.IntOfInt64(761)
					}
					return (_445_v9).F9()
				})())
				var _446_v10 _dafny.Array
				_ = _446_v10
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw108
				_446_v10 = _nw108
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_446_v10), 0))
				_ = _index82
				(_446_v10).ArraySet1(_445_v9.F10, (_index82).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_446_v10), 0))
				_ = _index83
				(_446_v10).ArraySet1(Companion_Default___.Fm0(_442_cf3, globalState), (_index83).Int())
				var _447_v11 D2
				_ = _447_v11
				_447_v11 = Companion_D2_.Create_DC4_((_this).F8())
				var _448_v12 D2
				_ = _448_v12
				_448_v12 = Companion_D2_.Create_DC6_(_447_v11)
				var _449_v13 D2
				_ = _449_v13
				_449_v13 = Companion_D2_.Create_DC6_(_448_v12)
				var _450_v14 D2
				_ = _450_v14
				_450_v14 = Companion_D2_.Create_DC6_(_447_v11)
				var _451_v15 D2
				_ = _451_v15
				_451_v15 = Companion_D2_.Create_DC6_(_450_v14)
				var _452_v16 _dafny.Map
				_ = _452_v16
				_452_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_451_v15, (_this).F7())
				var _453_v17 _dafny.Map
				_ = _453_v17
				_453_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(_450_v14), _dafny.IntOfInt64(-176)))
				var _454_v18 _dafny.Map
				_ = _454_v18
				_454_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_442_cf3, p1)
				var _455_v19 _dafny.Map
				_ = _455_v19
				_455_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v16).Merge((func() _dafny.Map {
					if (_453_v17).Contains(p1) {
						return (_453_v17).Get(p1).(_dafny.Map)
					}
					return _452_v16
				})()), (func() bool {
					if (_454_v18).Contains(_442_cf3) {
						return (_454_v18).Get(_442_cf3).(bool)
					}
					return p1
				})())
				_455_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_452_v16, p1)
				var _456_v20 *C0
				_ = _456_v20
				var _nw109 *C0 = New_C0_()
				_ = _nw109
				_nw109.Ctor__(p1)
				_456_v20 = _nw109
			} else {
				var _457_v21 _dafny.Map
				_ = _457_v21
				_457_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((false) == (false), p1)
				_457_v21 = (_457_v21).Update(p1, p1)
				var _458_v22 T0
				_ = _458_v22
				var _nw110 *C1 = New_C1_()
				_ = _nw110
				_nw110.Ctor__(p1, Companion_Default___.Fm0(_442_cf3, globalState), (_this).F7(), (_this).F7())
				_458_v22 = _nw110
				var _459_v23 D2
				_ = _459_v23
				_459_v23 = Companion_D2_.Create_DC5_((_this).F7(), p0, _458_v22)
				var _460_v24 D2
				_ = _460_v24
				_460_v24 = Companion_D2_.Create_DC6_(_459_v23)
				var _461_v25 _dafny.Map
				_ = _461_v25
				_461_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _460_v24))
				_461_v25 = _461_v25
				var _462_v26 _dafny.Array
				_ = _462_v26
				var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
				_ = _nw111
				_462_v26 = _nw111
				_462_v26 = (Companion_D0_.Create_DC0_(_462_v26)).Dtor_cf0()
				var _463_v27 D0
				_ = _463_v27
				_463_v27 = Companion_D0_.Create_DC0_(_462_v26)
				var _464_v28 D0
				_ = _464_v28
				_464_v28 = Companion_D0_.Create_DC0_((_463_v27).Dtor_cf0())
				_463_v27 = _463_v27
				r2 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(486), (_458_v22).F7())
			}
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_432_v0), 0))
			_ = _index84
			(_432_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hql"), p0), (_index84).Int())
			var _465_v29 _dafny.Array
			_ = _465_v29
			var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw112
			_465_v29 = _nw112
			var _466_v30 _dafny.Set
			_ = _466_v30
			_466_v30 = _dafny.SetOf(_465_v29, _465_v29)
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_465_v29), 0))
			_ = _index85
			(_465_v29).ArraySet1(Companion_Default___.SafeDivisionInt((_466_v30).Cardinality(), p3), (_index85).Int())
			var _467_v32 _dafny.Set
			_ = _467_v32
			_467_v32 = _dafny.SetOf(_442_cf3)
			var _468_v33 _dafny.Set
			_ = _468_v33
			_468_v33 = _dafny.SetOf((_467_v32).Cardinality())
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_432_v0), 0))
			_ = _index86
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_465_v29), 0))
			_ = _index87
			var _rhs67 bool = Companion_Default___.Fm2(p1, globalState)
			_ = _rhs67
			var _rhs68 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p0, Companion_Default___.Fm11(p1, _dafny.IntOfUint32((p0).Cardinality()), globalState))
			_ = _rhs68
			var _rhs69 _dafny.Int = p3
			_ = _rhs69
			var _rhs70 _dafny.Int = (((_434_v2).Merge(_434_v2)).Merge(func() _dafny.Map {
				var _coll18 = _dafny.NewMapBuilder()
				_ = _coll18
				for _iter18 := _dafny.Iterate((_468_v33).Elements()); ; {
					_compr_18, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _469_v31 _dafny.Int
					_469_v31 = interface{}(_compr_18).(_dafny.Int)
					if (_468_v33).Contains(_469_v31) {
						_coll18.Add((_469_v31).Minus((_this).F7()), (Companion_D2_.Create_DC4_(p3)).Dtor_cf4())
					}
				}
				return _coll18.ToMap()
			}())).Cardinality()
			_ = _rhs70
			var _lhs58 *GlobalState = globalState
			_ = _lhs58
			var _lhs59 _dafny.Array = _432_v0
			_ = _lhs59
			var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.ArrayLen((_432_v0), 0))
			_ = _lhs60
			var _lhs61 _dafny.Array = _465_v29
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_465_v29), 0))
			_ = _lhs62
			_lhs58.F1 = _rhs67
			(_lhs59).ArraySet1(_rhs68, (_lhs60).Int())
			(_lhs61).ArraySet1(_rhs69, (_lhs62).Int())
			r0 = _rhs70
			Companion_Default___.M0(globalState)
		} else {
			var _470___mcc_h1 bool = _source4.Get_().(D1_DC2).Cf2
			_ = _470___mcc_h1
			var _471_cf2 bool = _470___mcc_h1
			_ = _471_cf2
			var _472_v34 _dafny.Array
			_ = _472_v34
			var _nw113 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw113
			_472_v34 = _nw113
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_472_v34), 0))
			_ = _index88
			(_472_v34).ArraySet1(p1, (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_472_v34), 0))
			_ = _index89
			(_472_v34).ArraySet1(_471_cf2, (_index89).Int())
			var _473_v35 _dafny.Array
			_ = _473_v35
			var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw114
			_473_v35 = _nw114
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_473_v35), 0))
			_ = _index90
			(_473_v35).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_dafny.CodePoint('d'), globalState), p3), (_index90).Int())
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_473_v35), 0))
			_ = _index91
			(_473_v35).ArraySet1((p3).Minus((_dafny.Zero).Minus((_this).F7())), (_index91).Int())
			(globalState).F3 = _471_cf2
			_471_cf2 = _dafny.Companion_Sequence_.IsPrefixOf(_433_v1, _dafny.Companion_Sequence_.Concatenate(_433_v1, _433_v1))
		}
		r0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_435_v3).Cardinality()), (_this).F7())
		var _474_v36 D4
		_ = _474_v36
		_474_v36 = Companion_D4_.Create_DC11_((_this).F7(), _437_v5)
		var _475_v37 D4
		_ = _475_v37
		_475_v37 = Companion_D4_.Create_DC13_(_474_v36)
		var _pat_let_tv7 = _437_v5
		_ = _pat_let_tv7
		var _pat_let_tv8 = p0
		_ = _pat_let_tv8
		var _pat_let_tv9 = _437_v5
		_ = _pat_let_tv9
		var _pat_let_tv10 = _437_v5
		_ = _pat_let_tv10
		r1 = func(_source5 D4) _dafny.Sequence {
			if _source5.Is_DC11() {
				var _476___mcc_h2 _dafny.Int = _source5.Get_().(D4_DC11).Cf13
				_ = _476___mcc_h2
				var _477___mcc_h3 _dafny.CodePoint = _source5.Get_().(D4_DC11).Cf14
				_ = _477___mcc_h3
				var _478_cf14 _dafny.CodePoint = _477___mcc_h3
				_ = _478_cf14
				var _479_cf13 _dafny.Int = _476___mcc_h2
				_ = _479_cf13
				return _dafny.SeqOf(_478_cf14, _pat_let_tv7)
			} else if _source5.Is_DC12() {
				var _480___mcc_h4 _dafny.Int = _source5.Get_().(D4_DC12).Cf15
				_ = _480___mcc_h4
				var _481___mcc_h5 _dafny.Sequence = _source5.Get_().(D4_DC12).Cf16
				_ = _481___mcc_h5
				var _482___mcc_h6 bool = _source5.Get_().(D4_DC12).Cf17
				_ = _482___mcc_h6
				var _483_cf17 bool = _482___mcc_h6
				_ = _483_cf17
				var _484_cf16 _dafny.Sequence = _481___mcc_h5
				_ = _484_cf16
				var _485_cf15 _dafny.Int = _480___mcc_h4
				_ = _485_cf15
				return _pat_let_tv8
			} else if _source5.Is_DC10() {
				var _486___mcc_h7 _dafny.Map = _source5.Get_().(D4_DC10).Cf12
				_ = _486___mcc_h7
				var _487_cf12 _dafny.Map = _486___mcc_h7
				_ = _487_cf12
				return _dafny.SeqOf(_pat_let_tv9)
			} else {
				var _488___mcc_h8 D4 = _source5.Get_().(D4_DC13).Cf18
				_ = _488___mcc_h8
				var _489_cf18 D4 = _488___mcc_h8
				_ = _489_cf18
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(873))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}((func(_490_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_491_i0 _dafny.Int) _dafny.CodePoint {
						return _490_v5
					}
				})(_pat_let_tv10)))
			}
		}(_475_v37)
		r2 = (_this).F8()
		r3 = p1
		return r0, r1, r2, r3
	}
}
func (_this *C3) M4(p0 _dafny.Array, globalState *GlobalState) (_dafny.Array, bool, _dafny.Sequence) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _492_v0 bool
		_ = _492_v0
		_492_v0 = false
		var _493_v1 _dafny.Sequence
		_ = _493_v1
		_493_v1 = _dafny.UnicodeSeqOfUtf8Bytes("xeoexah")
		var _494_v2 T0
		_ = _494_v2
		var _nw115 *C1 = New_C1_()
		_ = _nw115
		_nw115.Ctor__(_492_v0, (_this).F7(), _dafny.IntOfUint32((_493_v1).Cardinality()), (_this).F7())
		_494_v2 = _nw115
		var _495_v3 _dafny.Sequence
		_ = _495_v3
		_495_v3 = _dafny.SeqOf(_494_v2)
		_495_v3 = _495_v3
		var _496_v4 _dafny.Int
		_ = _496_v4
		_496_v4 = _dafny.IntOfInt64(96)
		_496_v4 = Companion_Default___.SafeDivisionInt(_496_v4, _dafny.IntOfInt64(-305))
		var _497_v5 _dafny.Array
		_ = _497_v5
		var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw116
		_497_v5 = _nw116
		var _498_v6 _dafny.Array
		_ = _498_v6
		var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
		_ = _nw117
		_498_v6 = _nw117
		var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_497_v5), 0))
		_ = _index92
		(_497_v5).ArraySet1(_498_v6, (_index92).Int())
		var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_497_v5), 0))
		_ = _index93
		(_497_v5).ArraySet1(_498_v6, (_index93).Int())
		var _hi3 _dafny.Int = _496_v4
		_ = _hi3
		for _499_i0 := _496_v4; _499_i0.Cmp(_hi3) < 0; _499_i0 = _499_i0.Plus(_dafny.One) {
			var _500_v7 _dafny.Map
			_ = _500_v7
			_500_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_492_v0, _492_v0)
			var _501_v8 _dafny.Array
			_ = _501_v8
			var _nwElement0_28 bool = false
			_ = _nwElement0_28
			var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(25))
			_ = _nw118
			(_nw118).ArraySet1(_nwElement0_28, 0)
			(_nw118).ArraySet1(_492_v0, 1)
			(_nw118).ArraySet1(false, 2)
			(_nw118).ArraySet1(_492_v0, 3)
			(_nw118).ArraySet1((func() bool {
				if (_500_v7).Contains((func() bool {
					if (_500_v7).Contains(_492_v0) {
						return (_500_v7).Get(_492_v0).(bool)
					}
					return _492_v0
				})()) {
					return (_500_v7).Get((func() bool {
						if (_500_v7).Contains(_492_v0) {
							return (_500_v7).Get(_492_v0).(bool)
						}
						return _492_v0
					})()).(bool)
				}
				return _492_v0
			})(), 4)
			(_nw118).ArraySet1(_492_v0, 5)
			(_nw118).ArraySet1(_492_v0, 6)
			(_nw118).ArraySet1(_492_v0, 7)
			(_nw118).ArraySet1(_492_v0, 8)
			(_nw118).ArraySet1(_492_v0, 9)
			(_nw118).ArraySet1(_492_v0, 10)
			(_nw118).ArraySet1(_492_v0, 11)
			(_nw118).ArraySet1(_492_v0, 12)
			(_nw118).ArraySet1(_492_v0, 13)
			(_nw118).ArraySet1(_492_v0, 14)
			(_nw118).ArraySet1(false, 15)
			(_nw118).ArraySet1(_492_v0, 16)
			(_nw118).ArraySet1(true, 17)
			(_nw118).ArraySet1(_492_v0, 18)
			(_nw118).ArraySet1(!(_492_v0), 19)
			(_nw118).ArraySet1(_492_v0, 20)
			(_nw118).ArraySet1(false, 21)
			(_nw118).ArraySet1(_492_v0, 22)
			(_nw118).ArraySet1(_492_v0, 23)
			(_nw118).ArraySet1(_492_v0, 24)
			_501_v8 = _nw118
			var _502_v9 _dafny.Map
			_ = _502_v9
			_502_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_492_v0, _501_v8)
			var _503_v10 D5
			_ = _503_v10
			_503_v10 = Companion_D5_.Create_DC14_(_501_v8)
			var _504_v11 _dafny.Array
			_ = _504_v11
			var _nwElement0_29 _dafny.Array = _501_v8
			_ = _nwElement0_29
			var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(24))
			_ = _nw119
			(_nw119).ArraySet1(_nwElement0_29, 0)
			(_nw119).ArraySet1(_501_v8, 1)
			(_nw119).ArraySet1(_501_v8, 2)
			(_nw119).ArraySet1(_501_v8, 3)
			(_nw119).ArraySet1(_501_v8, 4)
			(_nw119).ArraySet1(_501_v8, 5)
			(_nw119).ArraySet1(_501_v8, 6)
			(_nw119).ArraySet1((func() _dafny.Array {
				if (_502_v9).Contains(_492_v0) {
					return (_502_v9).Get(_492_v0).(_dafny.Array)
				}
				return _501_v8
			})(), 7)
			(_nw119).ArraySet1((_503_v10).Dtor_cf19(), 8)
			(_nw119).ArraySet1(_501_v8, 9)
			(_nw119).ArraySet1(_501_v8, 10)
			(_nw119).ArraySet1(_501_v8, 11)
			(_nw119).ArraySet1((_503_v10).Dtor_cf19(), 12)
			(_nw119).ArraySet1(_501_v8, 13)
			(_nw119).ArraySet1(_501_v8, 14)
			(_nw119).ArraySet1(_501_v8, 15)
			(_nw119).ArraySet1(_501_v8, 16)
			(_nw119).ArraySet1(_501_v8, 17)
			(_nw119).ArraySet1(_501_v8, 18)
			(_nw119).ArraySet1(_501_v8, 19)
			(_nw119).ArraySet1(_501_v8, 20)
			(_nw119).ArraySet1(_501_v8, 21)
			(_nw119).ArraySet1(_501_v8, 22)
			(_nw119).ArraySet1(_501_v8, 23)
			_504_v11 = _nw119
			var _505_v12 _dafny.CodePoint
			_ = _505_v12
			_505_v12 = _dafny.CodePoint('y')
			var _506_v13 _dafny.Map
			_ = _506_v13
			_506_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_493_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-796), _dafny.IntOfUint32((_493_v1).Cardinality()))).Uint32(), _505_v12)).Cardinality()))
			var _507_v14 _dafny.Set
			_ = _507_v14
			_507_v14 = _dafny.SetOf(_493_v1, _493_v1)
			var _rhs71 _dafny.Array = _504_v11
			_ = _rhs71
			var _rhs72 _dafny.Int = ((_dafny.Zero).Minus((_506_v13).Cardinality())).Minus(Companion_Default___.SafeModuloInt((_494_v2).F8(), (_507_v14).Cardinality()))
			_ = _rhs72
			var _rhs73 _dafny.Int = (_494_v2).F8()
			_ = _rhs73
			_504_v11 = _rhs71
			_496_v4 = _rhs72
			_496_v4 = _rhs73
			var _508_v15 _dafny.Map
			_ = _508_v15
			_508_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_494_v2).F8(), _dafny.SeqOf(p0))
			_508_v15 = (_508_v15).Update((_494_v2).F8(), _dafny.SeqOf(p0))
			(globalState).F3 = Companion_Default___.Fm2(_492_v0, globalState)
			_501_v8 = (func() _dafny.Array {
				if _492_v0 {
					return _501_v8
				}
				return _501_v8
			})()
		}
		var _509_v16 _dafny.CodePoint
		_ = _509_v16
		_509_v16 = _dafny.CodePoint('s')
		var _510_v17 *C2
		_ = _510_v17
		var _nw120 *C2 = New_C2_()
		_ = _nw120
		_nw120.Ctor__(Companion_Default___.Fm0(_509_v16, globalState), _496_v4, (_494_v2).F8(), _dafny.IntOfInt64(250))
		_510_v17 = _nw120
		var _511_v18 *C2
		_ = _511_v18
		var _nw121 *C2 = New_C2_()
		_ = _nw121
		_nw121.Ctor__((_494_v2).F8(), _496_v4, (_494_v2).F8(), (_this).F7())
		_511_v18 = _nw121
		r0 = (func() _dafny.Array {
			if _492_v0 {
				return (func() _dafny.Array {
					if _492_v0 {
						return _498_v6
					}
					return _dafny.ArrayCastTo((_497_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_497_v5), 0))).Int()))
				})()
			}
			return _dafny.ArrayCastTo((_497_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_497_v5), 0))).Int()))
		})()
		var _512_v19 _dafny.MultiSet
		_ = _512_v19
		_512_v19 = _dafny.MultiSetOf(!(_492_v0))
		r1 = (_512_v19).Contains((func() bool {
			if _492_v0 {
				return _492_v0
			}
			return _492_v0
		})())
		var _513_v20 _dafny.Sequence
		_ = _513_v20
		_513_v20 = _dafny.SeqOf(_dafny.IntOfInt64(996), _511_v18.F10)
		var _514_v21 D6
		_ = _514_v21
		_514_v21 = Companion_D6_.Create_DC16_(_513_v20)
		r2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_513_v20, (_514_v21).Dtor_cf22()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(786))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_515_v0 bool) func(_dafny.Int) _dafny.Int {
			return func(_516_i1 _dafny.Int) _dafny.Int {
				return (func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter19 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_dafny.CodePoint('n')), (_this).F8())).Keys().Elements()); ; {
						_compr_19, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _517_v22 D1
						_517_v22 = interface{}(_compr_19).(D1)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_dafny.CodePoint('n')), (_this).F8())).Contains(_517_v22) {
							_coll19.Add(_517_v22, _515_v0)
						}
					}
					return _coll19.ToMap()
				}()).Cardinality()
			}
		})(_492_v0))))
		return r0, r1, r2
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f6 _dafny.Map
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f6 = _dafny.EmptyMap
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

func (_this *C4) Ctor__(f6 _dafny.Map) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C4) M1(globalState *GlobalState) {
	{
		var _518_v0 D0
		_ = _518_v0
		_518_v0 = Companion_D0_.Create_DC1_(_dafny.CodePoint('e'))
		var _519_v1 bool
		_ = _519_v1
		_519_v1 = true
		var _520_v2 _dafny.CodePoint
		_ = _520_v2
		_520_v2 = _dafny.CodePoint('t')
		var _521_v3 D1
		_ = _521_v3
		_521_v3 = Companion_D1_.Create_DC2_(_519_v1)
		var _522_v4 _dafny.Array
		_ = _522_v4
		var _nwElement0_30 D0 = _518_v0
		_ = _nwElement0_30
		var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(23))
		_ = _nw122
		(_nw122).ArraySet1(_nwElement0_30, 0)
		(_nw122).ArraySet1(_518_v0, 1)
		(_nw122).ArraySet1(_518_v0, 2)
		(_nw122).ArraySet1(_518_v0, 3)
		(_nw122).ArraySet1(_518_v0, 4)
		(_nw122).ArraySet1(_518_v0, 5)
		(_nw122).ArraySet1(Companion_Default___.Fm5(_519_v1, _519_v1, _519_v1, globalState), 6)
		(_nw122).ArraySet1(Companion_D0_.Create_DC1_(_520_v2), 7)
		(_nw122).ArraySet1(Companion_D0_.Create_DC1_(_520_v2), 8)
		(_nw122).ArraySet1(Companion_D0_.Create_DC1_(_520_v2), 9)
		(_nw122).ArraySet1(_518_v0, 10)
		(_nw122).ArraySet1(_518_v0, 11)
		(_nw122).ArraySet1(Companion_Default___.Fm5(_519_v1, Companion_Default___.Fm2((_521_v3).Dtor_cf2(), globalState), _519_v1, globalState), 12)
		(_nw122).ArraySet1(Companion_D0_.Create_DC1_(_520_v2), 13)
		(_nw122).ArraySet1(_518_v0, 14)
		(_nw122).ArraySet1(Companion_Default___.Fm5(_519_v1, _519_v1, false, globalState), 15)
		(_nw122).ArraySet1(_518_v0, 16)
		(_nw122).ArraySet1(_518_v0, 17)
		(_nw122).ArraySet1(Companion_Default___.Fm5(_519_v1, _519_v1, _519_v1, globalState), 18)
		(_nw122).ArraySet1(_518_v0, 19)
		(_nw122).ArraySet1(_518_v0, 20)
		(_nw122).ArraySet1(_518_v0, 21)
		(_nw122).ArraySet1(Companion_D0_.Create_DC1_(_520_v2), 22)
		_522_v4 = _nw122
		var _523_v5 _dafny.Int
		_ = _523_v5
		_523_v5 = _dafny.IntOfInt64(-361)
		var _524_v6 _dafny.MultiSet
		_ = _524_v6
		_524_v6 = _dafny.MultiSetOf(_523_v5)
		var _525_v7 _dafny.Sequence
		_ = _525_v7
		_525_v7 = _dafny.SeqOf(_519_v1, _519_v1)
		var _526_v8 _dafny.MultiSet
		_ = _526_v8
		_526_v8 = _dafny.MultiSetOf((_524_v6).Cardinality(), _dafny.IntOfUint32((_525_v7).Cardinality()))
		var _527_v9 _dafny.Map
		_ = _527_v9
		_527_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_522_v4, (_524_v6).Cardinality())
		_527_v9 = (_527_v9).Update(_522_v4, _dafny.IntOfInt64(59))
		_519_v1 = _519_v1
		var _528_i0 _dafny.Int
		_ = _528_i0
		_528_i0 = _dafny.Zero
		{
			for _519_v1 {
				{
					if (_528_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_528_i0 = (_528_i0).Plus(_dafny.One)
					var _529_v10 _dafny.Array
					_ = _529_v10
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(18)
					_ = _len0_13
					var _nw123 _dafny.Array
					_ = _nw123
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw123 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) _dafny.Int = (func(_530_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_531_i1 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_531_i1, (_dafny.Zero).Minus(_530_v5))
							}
						})(_523_v5)
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw123 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw123).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw123).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_529_v10 = _nw123
					var _532_v11 D0
					_ = _532_v11
					_532_v11 = Companion_D0_.Create_DC0_(_529_v10)
					var _533_v12 _dafny.Sequence
					_ = _533_v12
					_533_v12 = _dafny.SeqOf(_529_v10)
					var _pat_let_tv11 = _533_v12
					_ = _pat_let_tv11
					var _pat_let_tv12 = _523_v5
					_ = _pat_let_tv12
					var _pat_let_tv13 = _533_v12
					_ = _pat_let_tv13
					var _source6 D0 = func(_pat_let12_0 D0) D0 {
						return func(_534_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let13_0 _dafny.Array) D0 {
								return func(_535_dt__update_hcf0_h0 _dafny.Array) D0 {
									return Companion_D0_.Create_DC0_(_535_dt__update_hcf0_h0)
								}(_pat_let13_0)
							}((_pat_let_tv11).Select((Companion_Default___.SafeIndex(_pat_let_tv12, _dafny.IntOfUint32((_pat_let_tv13).Cardinality()))).Uint32()).(_dafny.Array))
						}(_pat_let12_0)
					}(_532_v11)
					_ = _source6
					if _source6.Is_DC1() {
						var _536___mcc_h0 _dafny.CodePoint = _source6.Get_().(D0_DC1).Cf1
						_ = _536___mcc_h0
						var _537_cf1 _dafny.CodePoint = _536___mcc_h0
						_ = _537_cf1
						var _538_v13 *C3
						_ = _538_v13
						var _nw124 *C3 = New_C3_()
						_ = _nw124
						_nw124.Ctor__(_523_v5, (func() _dafny.Int {
							if _519_v1 {
								return _523_v5
							}
							return _523_v5
						})())
						_538_v13 = _nw124
						var _539_v14 _dafny.Map
						_ = _539_v14
						_539_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(524), _523_v5)
						var _540_v15 _dafny.Map
						_ = _540_v15
						_540_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_519_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_523_v5, _523_v5))
						var _541_v17 _dafny.Sequence
						_ = _541_v17
						_541_v17 = _dafny.SeqOf((func() _dafny.Map {
							if (_540_v15).Contains(_519_v1) {
								return (_540_v15).Get(_519_v1).(_dafny.Map)
							}
							return func() _dafny.Map {
								var _coll20 = _dafny.NewMapBuilder()
								_ = _coll20
								for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(856), _dafny.IntOfInt64(532))); ; {
									_compr_20, _ok20 := _iter20()
									if !_ok20 {
										break
									}
									var _542_v16 _dafny.Int
									_542_v16 = interface{}(_compr_20).(_dafny.Int)
									if ((_dafny.IntOfInt64(856)).Cmp(_542_v16) <= 0) && ((_542_v16).Cmp(_dafny.IntOfInt64(532)) < 0) {
										_coll20.Add(Companion_Default___.SafeModuloInt(_542_v16, _523_v5), _523_v5)
									}
								}
								return _coll20.ToMap()
							}()
						})())
						var _543_v18 _dafny.Map
						_ = _543_v18
						_543_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_520_v2, !(_dafny.Companion_Sequence_.Contains(_541_v17, _539_v14)))
						_543_v18 = (_543_v18).Update(_520_v2, true)
						var _544_v19 _dafny.Array
						_ = _544_v19
						var _len0_14 _dafny.Int = _dafny.IntOfInt64(26)
						_ = _len0_14
						var _nw125 _dafny.Array
						_ = _nw125
						if _len0_14.Cmp(_dafny.Zero) == 0 {
							_nw125 = _dafny.NewArray(_len0_14)
						} else {
							var _init14 func(_dafny.Int) bool = (func(_545_v1 bool) func(_dafny.Int) bool {
								return func(_546_i2 _dafny.Int) bool {
									return _545_v1
								}
							})(_519_v1)
							_ = _init14
							var _element0_14 = _init14(_dafny.Zero)
							_ = _element0_14
							_nw125 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
							(_nw125).ArraySet1(_element0_14, 0)
							var _nativeLen0_14 = (_len0_14).Int()
							_ = _nativeLen0_14
							for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
								(_nw125).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
							}
						}
						_544_v19 = _nw125
						var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_544_v19), 0))
						_ = _index94
						(_544_v19).ArraySet1((func() bool {
							if _519_v1 {
								return _519_v1
							}
							return _519_v1
						})(), (_index94).Int())
						var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_544_v19), 0))
						_ = _index95
						(_544_v19).ArraySet1(_519_v1, (_index95).Int())
						var _547_v20 _dafny.Sequence
						_ = _547_v20
						_547_v20 = _dafny.UnicodeSeqOfUtf8Bytes("xjhkseht")
						var _548_v21 D6
						_ = _548_v21
						_548_v21 = Companion_D6_.Create_DC17_(_547_v20, _519_v1, _523_v5, false, !((_544_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(679), _dafny.ArrayLen((_544_v19), 0))).Int()).(bool)))
						(globalState).F3 = !_dafny.Companion_Sequence_.Equal(_547_v20, (_548_v21).Dtor_cf23())
					} else {
						var _549___mcc_h1 _dafny.Array = _source6.Get_().(D0_DC0).Cf0
						_ = _549___mcc_h1
						var _550_cf0 _dafny.Array = _549___mcc_h1
						_ = _550_cf0
						var _551_v22 _dafny.Sequence
						_ = _551_v22
						_551_v22 = _dafny.UnicodeSeqOfUtf8Bytes("jsstabmc")
						_551_v22 = _dafny.Companion_Sequence_.Concatenate(_551_v22, (func() _dafny.Sequence {
							if _519_v1 {
								return _551_v22
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("f")
						})())
						var _552_v23 _dafny.Array
						_ = _552_v23
						var _len0_15 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_15
						var _nw126 _dafny.Array
						_ = _nw126
						if _len0_15.Cmp(_dafny.Zero) == 0 {
							_nw126 = _dafny.NewArray(_len0_15)
						} else {
							var _init15 func(_dafny.Int) bool = (func(_553_v1 bool) func(_dafny.Int) bool {
								return func(_554_i3 _dafny.Int) bool {
									return _553_v1
								}
							})(_519_v1)
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
						_552_v23 = _nw126
						_552_v23 = _552_v23
						_551_v22 = Companion_Default___.Fm11(_519_v1, _523_v5, globalState)
						_523_v5 = (_dafny.Zero).Minus(_523_v5)
					}
					_520_v2 = _520_v2
					var _555_v24 _dafny.Array
					_ = _555_v24
					var _nwElement0_31 bool = _519_v1
					_ = _nwElement0_31
					var _nw127 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(26))
					_ = _nw127
					(_nw127).ArraySet1(_nwElement0_31, 0)
					(_nw127).ArraySet1(_519_v1, 1)
					(_nw127).ArraySet1(!(_519_v1), 2)
					(_nw127).ArraySet1(_519_v1, 3)
					(_nw127).ArraySet1(_519_v1, 4)
					(_nw127).ArraySet1(_519_v1, 5)
					(_nw127).ArraySet1(Companion_Default___.Fm2(_519_v1, globalState), 6)
					(_nw127).ArraySet1(false, 7)
					(_nw127).ArraySet1(_519_v1, 8)
					(_nw127).ArraySet1(_519_v1, 9)
					(_nw127).ArraySet1(_519_v1, 10)
					(_nw127).ArraySet1(_519_v1, 11)
					(_nw127).ArraySet1(_519_v1, 12)
					(_nw127).ArraySet1(true, 13)
					(_nw127).ArraySet1(!(false), 14)
					(_nw127).ArraySet1(_519_v1, 15)
					(_nw127).ArraySet1(_519_v1, 16)
					(_nw127).ArraySet1(_519_v1, 17)
					(_nw127).ArraySet1(_519_v1, 18)
					(_nw127).ArraySet1(_519_v1, 19)
					(_nw127).ArraySet1(true, 20)
					(_nw127).ArraySet1(_519_v1, 21)
					(_nw127).ArraySet1(Companion_Default___.Fm2(_519_v1, globalState), 22)
					(_nw127).ArraySet1(_519_v1, 23)
					(_nw127).ArraySet1(_519_v1, 24)
					(_nw127).ArraySet1(!(!(_519_v1)), 25)
					_555_v24 = _nw127
					var _556_v25 _dafny.MultiSet
					_ = _556_v25
					_556_v25 = _dafny.MultiSetOf(_555_v24, _555_v24)
					var _557_v26 _dafny.Map
					_ = _557_v26
					_557_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_556_v25, _519_v1)
					_557_v26 = (_557_v26).Update(_556_v25, (_523_v5).Cmp(_523_v5) == 0)
					var _558_v27 _dafny.Sequence
					_ = _558_v27
					_558_v27 = _dafny.SeqOf(_523_v5)
					var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_529_v10), 0))
					_ = _index96
					(_529_v10).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-759))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}((func(_559_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_560_i4 _dafny.Int) _dafny.CodePoint {
							return _559_v2
						}
					})(_520_v2))), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("t"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_558_v27).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()))).Uint32(), _520_v2))).Cardinality()), (_index96).Int())
					var _561_v28 _dafny.MultiSet
					_ = _561_v28
					_561_v28 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_525_v7).Cardinality()), _523_v5))
					var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_529_v10), 0))
					_ = _index97
					(_529_v10).ArraySet1(((_561_v28).Intersection((_561_v28).Union(Companion_Default___.Fm18(globalState)))).Cardinality(), (_index97).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _562_v29 _dafny.Sequence
		_ = _562_v29
		_562_v29 = _dafny.UnicodeSeqOfUtf8Bytes("geidfsen")
		var _563_v30 _dafny.Map
		_ = _563_v30
		_563_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_523_v5, _dafny.IntOfInt64(188))
		var _pat_let_tv14 = _519_v1
		_ = _pat_let_tv14
		var _564_v32 D6
		_ = _564_v32
		_564_v32 = Companion_D6_.Create_DC17_(_562_v29, (func(_pat_let14_0 D4) D4 {
			return func(_566_dt__update__tmp_h1 D4) D4 {
				return func(_pat_let15_0 bool) D4 {
					return func(_567_dt__update_hcf17_h0 bool) D4 {
						return Companion_D4_.Create_DC12_((_566_dt__update__tmp_h1).Dtor_cf15(), (_566_dt__update__tmp_h1).Dtor_cf16(), _567_dt__update_hcf17_h0)
					}(_pat_let15_0)
				}(!(_pat_let_tv14))
			}(_pat_let14_0)
		}(Companion_D4_.Create_DC12_((func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate((_563_v30).Keys().Elements()); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _565_v31 _dafny.Int
				_565_v31 = interface{}(_compr_21).(_dafny.Int)
				if (_563_v30).Contains(_565_v31) {
					_coll21.Add((_565_v31).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(143))).Cardinality()))
				}
			}
			return _coll21.ToSet()
		}()).Cardinality(), _dafny.SeqOf(_519_v1), false))).Dtor_cf17(), (func() _dafny.Int {
			if (_526_v8).Contains(_523_v5) {
				return (_526_v8).Multiplicity(_523_v5)
			}
			return _523_v5
		})(), (_525_v7).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_523_v5), _dafny.IntOfUint32((_525_v7).Cardinality()))).Uint32()).(bool), false)
		_519_v1 = (_564_v32).Dtor_cf27()
		var _568_v33 _dafny.Array
		_ = _568_v33
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_16
		var _nw128 _dafny.Array
		_ = _nw128
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw128 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) bool = (func(_569_v1 bool) func(_dafny.Int) bool {
				return func(_570_i5 _dafny.Int) bool {
					return _569_v1
				}
			})(_519_v1)
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw128 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw128).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw128).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		_568_v33 = _nw128
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_568_v33), 0))
		_ = _index98
		(_568_v33).ArraySet1(_519_v1, (_index98).Int())
		var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_568_v33), 0))
		_ = _index99
		(_568_v33).ArraySet1(!((_525_v7).Select((Companion_Default___.SafeIndex(_523_v5, _dafny.IntOfUint32((_525_v7).Cardinality()))).Uint32()).(bool)), (_index99).Int())
		(globalState).F1 = (_568_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_568_v33), 0))).Int()).(bool)
	}
}
func (_this *C4) M2(p0 _dafny.Map, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) {
	{
		var _571_v0 _dafny.Array
		_ = _571_v0
		var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw129
		_571_v0 = _nw129
		var _572_v1 D0
		_ = _572_v1
		_572_v1 = Companion_D0_.Create_DC0_(_571_v0)
		var _573_v2 _dafny.Array
		_ = _573_v2
		var _nw130 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw130
		_573_v2 = _nw130
		var _574_v3 _dafny.Map
		_ = _574_v3
		_574_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_572_v1).Dtor_cf0(), _573_v2)
		_574_v3 = (_574_v3).Update(_571_v0, _573_v2)
		var _575_v4 _dafny.Array
		_ = _575_v4
		var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
		_ = _nw131
		_575_v4 = _nw131
		var _576_v5 _dafny.Int
		_ = _576_v5
		_576_v5 = _dafny.IntOfInt64(-872)
		var _577_v6 _dafny.Sequence
		_ = _577_v6
		_577_v6 = _dafny.SeqOf(_576_v5)
		var _578_v7 _dafny.Map
		_ = _578_v7
		_578_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_576_v5, _dafny.IntOfUint32(((Companion_D6_.Create_DC16_(_577_v6)).Dtor_cf22()).Cardinality()))
		var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_575_v4), 0))
		_ = _index100
		(_575_v4).ArraySet1(_578_v7, (_index100).Int())
		var _579_v8 _dafny.CodePoint
		_ = _579_v8
		_579_v8 = _dafny.CodePoint('a')
		var _580_v9 _dafny.MultiSet
		_ = _580_v9
		_580_v9 = _dafny.MultiSetOf(_571_v0, _571_v0, _571_v0, _571_v0, _571_v0)
		var _581_v10 _dafny.Sequence
		_ = _581_v10
		_581_v10 = _dafny.UnicodeSeqOfUtf8Bytes("jmjl")
		var _582_v11 _dafny.Map
		_ = _582_v11
		_582_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)
		var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_575_v4), 0))
		_ = _index101
		var _rhs74 _dafny.Map = _578_v7
		_ = _rhs74
		var _rhs75 _dafny.CodePoint = _579_v8
		_ = _rhs75
		var _rhs76 bool = (Companion_D4_.Create_DC12_(_dafny.IntOfUint32((_581_v10).Cardinality()), p1, p2)).Dtor_cf17()
		_ = _rhs76
		var _rhs77 _dafny.Int = (_582_v11).Cardinality()
		_ = _rhs77
		var _rhs78 _dafny.MultiSet = _dafny.MultiSetOf(_571_v0)
		_ = _rhs78
		var _lhs63 _dafny.Array = _575_v4
		_ = _lhs63
		var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_575_v4), 0))
		_ = _lhs64
		var _lhs65 *GlobalState = globalState
		_ = _lhs65
		(_lhs63).ArraySet1(_rhs74, (_lhs64).Int())
		_579_v8 = _rhs75
		_lhs65.F3 = _rhs76
		_576_v5 = _rhs77
		_580_v9 = _rhs78
		var _583_v12 _dafny.MultiSet
		_ = _583_v12
		_583_v12 = _dafny.MultiSetOf(_576_v5, _576_v5)
		var _584_v13 _dafny.MultiSet
		_ = _584_v13
		_584_v13 = _dafny.MultiSetOf(_579_v8, _579_v8)
		var _585_v14 _dafny.MultiSet
		_ = _585_v14
		_585_v14 = _dafny.MultiSetOf((func() _dafny.Int {
			if (_583_v12).Contains(_576_v5) {
				return (_583_v12).Multiplicity(_576_v5)
			}
			return (_584_v13).Cardinality()
		})())
		var _586_v15 _dafny.Set
		_ = _586_v15
		_586_v15 = _dafny.SetOf(_583_v12)
		var _587_v16 D3
		_ = _587_v16
		_587_v16 = Companion_D3_.Create_DC9_(Companion_D3_.Create_DC7_(_586_v15))
		_587_v16 = Companion_Default___.Fm19(globalState)
		(globalState).F1 = (!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_588_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_589_i0 _dafny.Int) _dafny.CodePoint {
				return _588_v8
			}
		})(_579_v8))), _581_v10)) || (!(p2) || (p2))
		if !((p1).Select((Companion_Default___.SafeIndex(_576_v5, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool)) {
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_573_v2), 0))
			_ = _index102
			(_573_v2).ArraySet1(p2, (_index102).Int())
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_573_v2), 0))
			_ = _index103
			(_573_v2).ArraySet1(!((p1).Select((Companion_Default___.SafeIndex(_576_v5, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool)), (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_573_v2), 0))
			_ = _index104
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_573_v2), 0))
			_ = _index105
			var _rhs79 bool = p2
			_ = _rhs79
			var _rhs80 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(311))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_590_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_591_i1 _dafny.Int) _dafny.CodePoint {
					return _590_v8
				}
			})(_579_v8))), _581_v10)).Cardinality())
			_ = _rhs80
			var _rhs81 _dafny.Int = Companion_Default___.Fm0(_579_v8, globalState)
			_ = _rhs81
			var _rhs82 _dafny.Int = _576_v5
			_ = _rhs82
			var _rhs83 bool = p2
			_ = _rhs83
			var _lhs66 _dafny.Array = _573_v2
			_ = _lhs66
			var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_573_v2), 0))
			_ = _lhs67
			var _lhs68 _dafny.Array = _573_v2
			_ = _lhs68
			var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_573_v2), 0))
			_ = _lhs69
			(_lhs66).ArraySet1(_rhs79, (_lhs67).Int())
			_576_v5 = _rhs80
			_576_v5 = _rhs81
			_576_v5 = _rhs82
			(_lhs68).ArraySet1(_rhs83, (_lhs69).Int())
			(globalState).F1 = (_576_v5).Cmp((_576_v5).Minus(_576_v5)) != 0
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_573_v2), 0))
			_ = _index106
			(_573_v2).ArraySet1(!(p2), (_index106).Int())
			(globalState).F1 = Companion_Default___.Fm2((_573_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_573_v2), 0))).Int()).(bool), globalState)
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_575_v4), 0))
			_ = _index107
			(_575_v4).ArraySet1(((_575_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(285), _dafny.ArrayLen((_575_v4), 0))).Int()).(_dafny.Map)).Update(_576_v5, Companion_Default___.SafeModuloInt(_576_v5, _576_v5)), (_index107).Int())
		} else {
			var _592_v17 *C2
			_ = _592_v17
			var _nw132 *C2 = New_C2_()
			_ = _nw132
			_nw132.Ctor__(_576_v5, (_dafny.MultiSetFromSeq(_577_v6)).Cardinality(), _576_v5, _576_v5)
			_592_v17 = _nw132
			(globalState).F3 = Companion_Default___.Fm2(false, globalState)
			(globalState).F3 = p2
			var _593_v18 D6
			_ = _593_v18
			_593_v18 = Companion_D6_.Create_DC17_(_581_v10, false, _592_v17.F10, p2, p2)
			_576_v5 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_592_v17).F9(), (_592_v17).F9()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_594_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_595_i2 _dafny.Int) _dafny.CodePoint {
					return _594_v8
				}
			})(_579_v8))), (_593_v18).Dtor_cf23())).Cardinality()))
			var _596_v19 D3
			_ = _596_v19
			_596_v19 = Companion_D3_.Create_DC8_(_592_v17.F10)
			var _pat_let_tv15 = _592_v17
			_ = _pat_let_tv15
			var _597_v20 _dafny.MultiSet
			_ = _597_v20
			_597_v20 = _dafny.MultiSetOf(_596_v19, _596_v19, func(_pat_let16_0 D3) D3 {
				return func(_598_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let17_0 _dafny.Int) D3 {
						return func(_599_dt__update_hcf10_h0 _dafny.Int) D3 {
							return Companion_D3_.Create_DC8_(_599_dt__update_hcf10_h0)
						}(_pat_let17_0)
					}((_pat_let_tv15).F9())
				}(_pat_let16_0)
			}(_596_v19))
			var _600_v21 D4
			_ = _600_v21
			_600_v21 = Companion_D4_.Create_DC11_((_dafny.MultiSetOf(_576_v5, (_592_v17).F9())).Cardinality(), _579_v8)
			var _601_v22 _dafny.Set
			_ = _601_v22
			_601_v22 = _dafny.SetOf(true)
			var _602_v23 D2
			_ = _602_v23
			_602_v23 = Companion_D2_.Create_DC4_(_dafny.IntOfInt64(776))
			var _603_v24 _dafny.Sequence
			_ = _603_v24
			_603_v24 = _dafny.SeqOf(_601_v22, Companion_Default___.Fm20(_602_v23, p2, globalState))
			var _604_v25 _dafny.Array
			_ = _604_v25
			var _nwElement0_32 _dafny.Sequence = _581_v10
			_ = _nwElement0_32
			var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(24))
			_ = _nw133
			(_nw133).ArraySet1(_nwElement0_32, 0)
			(_nw133).ArraySet1(_581_v10, 1)
			(_nw133).ArraySet1(_dafny.Companion_Sequence_.Update(_581_v10, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.IntOfUint32((_581_v10).Cardinality()))).Uint32(), _579_v8), 2)
			(_nw133).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_581_v10, (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_597_v20).Contains(_596_v19) {
					return (_597_v20).Multiplicity(_596_v19)
				}
				return (_dafny.Zero).Minus((_600_v21).Dtor_cf13())
			})(), _dafny.IntOfUint32((_581_v10).Cardinality()))).Uint32(), _579_v8), (Companion_Default___.SafeIndex(((_603_v24).Select((Companion_Default___.SafeIndex(_592_v17.F10, _dafny.IntOfUint32((_603_v24).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_581_v10, (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_597_v20).Contains(_596_v19) {
					return (_597_v20).Multiplicity(_596_v19)
				}
				return (_dafny.Zero).Minus((_600_v21).Dtor_cf13())
			})(), _dafny.IntOfUint32((_581_v10).Cardinality()))).Uint32(), _579_v8)).Cardinality()))).Uint32(), _579_v8), 3)
			(_nw133).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("t"), _581_v10), 4)
			(_nw133).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-519))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_605_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_606_i3 _dafny.Int) _dafny.CodePoint {
					return _605_v8
				}
			})(_579_v8))), 5)
			(_nw133).ArraySet1(_581_v10, 6)
			(_nw133).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("sval"), 7)
			(_nw133).ArraySet1(_581_v10, 8)
			(_nw133).ArraySet1(_581_v10, 9)
			(_nw133).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_581_v10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(743))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_607_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_608_i4 _dafny.Int) _dafny.CodePoint {
					return _607_v8
				}
			})(_579_v8)))), 10)
			(_nw133).ArraySet1(_581_v10, 11)
			(_nw133).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_581_v10, _581_v10), 12)
			(_nw133).ArraySet1(_581_v10, 13)
			(_nw133).ArraySet1(_581_v10, 14)
			(_nw133).ArraySet1(_581_v10, 15)
			(_nw133).ArraySet1(_581_v10, 16)
			(_nw133).ArraySet1(_581_v10, 17)
			(_nw133).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_581_v10, _dafny.Companion_Sequence_.Update(_581_v10, (Companion_Default___.SafeIndex(_576_v5, _dafny.IntOfUint32((_581_v10).Cardinality()))).Uint32(), _579_v8)), 18)
			(_nw133).ArraySet1(_581_v10, 19)
			(_nw133).ArraySet1(_581_v10, 20)
			(_nw133).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_581_v10, _dafny.UnicodeSeqOfUtf8Bytes("tswjrvwl")), 21)
			(_nw133).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_581_v10, _581_v10), 22)
			(_nw133).ArraySet1(_581_v10, 23)
			_604_v25 = _nw133
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_604_v25), 0))
			_ = _index108
			(_604_v25).ArraySet1(_581_v10, (_index108).Int())
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_604_v25), 0))
			_ = _index109
			(_604_v25).ArraySet1((func() _dafny.Sequence {
				if !(!(p2) || (!(!(p2)))) {
					return _581_v10
				}
				return _dafny.Companion_Sequence_.Concatenate(_581_v10, _dafny.UnicodeSeqOfUtf8Bytes("hupyfkie"))
			})(), (_index109).Int())
		}
		var _609_v26 _dafny.Array
		_ = _609_v26
		var _nw134 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
		_ = _nw134
		_609_v26 = _nw134
		var _610_v27 *C3
		_ = _610_v27
		var _nw135 *C3 = New_C3_()
		_ = _nw135
		_nw135.Ctor__((_dafny.Zero).Minus((_dafny.SetOf(p2, true)).Cardinality()), _dafny.IntOfInt64(383))
		_610_v27 = _nw135
		var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_609_v26), 0))
		_ = _index110
		(_609_v26).ArraySet1(_610_v27, (_index110).Int())
		var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_609_v26), 0))
		_ = _index111
		(_609_v26).ArraySet1(_610_v27, (_index111).Int())
	}
}
func (_this *C4) F6() _dafny.Map {
	{
		return _this._f6
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
