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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqOf(Companion_D0_.Create_DC1_(true, false, _dafny.IntOfInt64(197)))).Cardinality())).Plus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("giqwu")).Cardinality()))), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(187), _dafny.IntOfInt64(69), _dafny.IntOfInt64(335))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.IntOfInt64(450)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(false), true)).Cardinality()), _dafny.IntOfInt64(-29)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(913))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ycbj")).Cardinality())
	}))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.CodePoint {
	if true {
		if true {
			return _dafny.CodePoint('g')
		} else {
			return _dafny.CodePoint('u')
		}
	} else {
		return _dafny.CodePoint('l')
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(_dafny.CodePoint('l'), (_dafny.One).Minus(_dafny.IntOfInt64(202)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, p2 D0, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), true)).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(165))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i0 _dafny.Int) _dafny.Int {
			return (_dafny.SetOf(true, true)).Cardinality()
		}))).Cardinality())) == 0 {
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("khvhcxm")).Cardinality())), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(321), true)).Cardinality())))))
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(550))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_2_i1 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))).Cardinality())
		}))
	})())
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("d")
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 D16 = Companion_D16_.Create_DC50_(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), true)))
	_ = _source0
	if _source0.Is_DC51() {
		var _3___mcc_h0 bool = _source0.Get_().(D16_DC51).Cf94
		_ = _3___mcc_h0
		var _4___mcc_h1 _dafny.Set = _source0.Get_().(D16_DC51).Cf95
		_ = _4___mcc_h1
		var _5___mcc_h2 bool = _source0.Get_().(D16_DC51).Cf96
		_ = _5___mcc_h2
		var _6___mcc_h3 _dafny.Int = _source0.Get_().(D16_DC51).Cf97
		_ = _6___mcc_h3
		var _7_cf97 _dafny.Int = _6___mcc_h3
		_ = _7_cf97
		var _8_cf96 bool = _5___mcc_h2
		_ = _8_cf96
		var _9_cf95 _dafny.Set = _4___mcc_h1
		_ = _9_cf95
		var _10_cf94 bool = _3___mcc_h0
		_ = _10_cf94
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(356))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_11_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})), _dafny.UnicodeSeqOfUtf8Bytes("clkt"))
	} else {
		var _12___mcc_h4 _dafny.MultiSet = _source0.Get_().(D16_DC50).Cf93
		_ = _12___mcc_h4
		var _13_cf93 _dafny.MultiSet = _12___mcc_h4
		_ = _13_cf93
		return _dafny.UnicodeSeqOfUtf8Bytes("may")
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(true)
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D0 {
	var _source1 D8 = Companion_D8_.Create_DC26_(!(!(false)), !(false), true, (_dafny.SetOf(_dafny.IntOfInt64(-666))).Cardinality())
	_ = _source1
	if _source1.Is_DC26() {
		var _14___mcc_h0 bool = _source1.Get_().(D8_DC26).Cf45
		_ = _14___mcc_h0
		var _15___mcc_h1 bool = _source1.Get_().(D8_DC26).Cf46
		_ = _15___mcc_h1
		var _16___mcc_h2 bool = _source1.Get_().(D8_DC26).Cf47
		_ = _16___mcc_h2
		var _17___mcc_h3 _dafny.Int = _source1.Get_().(D8_DC26).Cf48
		_ = _17___mcc_h3
		var _18_cf48 _dafny.Int = _17___mcc_h3
		_ = _18_cf48
		var _19_cf47 bool = _16___mcc_h2
		_ = _19_cf47
		var _20_cf46 bool = _15___mcc_h1
		_ = _20_cf46
		var _21_cf45 bool = _14___mcc_h0
		_ = _21_cf45
		if _19_cf47 {
			return Companion_D0_.Create_DC1_(_19_cf47, _19_cf47, _18_cf48)
		} else {
			return Companion_D0_.Create_DC1_(_20_cf46, _20_cf46, _18_cf48)
		}
	} else if _source1.Is_DC27() {
		var _22___mcc_h4 bool = _source1.Get_().(D8_DC27).Cf49
		_ = _22___mcc_h4
		var _23___mcc_h5 bool = _source1.Get_().(D8_DC27).Cf50
		_ = _23___mcc_h5
		var _24___mcc_h6 _dafny.Int = _source1.Get_().(D8_DC27).Cf51
		_ = _24___mcc_h6
		var _25___mcc_h7 _dafny.Int = _source1.Get_().(D8_DC27).Cf52
		_ = _25___mcc_h7
		var _26_cf52 _dafny.Int = _25___mcc_h7
		_ = _26_cf52
		var _27_cf51 _dafny.Int = _24___mcc_h6
		_ = _27_cf51
		var _28_cf50 bool = _23___mcc_h5
		_ = _28_cf50
		var _29_cf49 bool = _22___mcc_h4
		_ = _29_cf49
		return Companion_D0_.Create_DC1_(_29_cf49, _29_cf49, (func() _dafny.Set {
			var _coll0 = _dafny.NewBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((func() _dafny.Set {
				var _coll1 = _dafny.NewBuilder()
				_ = _coll1
				for _iter1 := _dafny.Iterate((func() _dafny.Set {
					var _coll2 = _dafny.NewBuilder()
					_ = _coll2
					for _iter2 := _dafny.Iterate((func() _dafny.Set {
						var _coll3 = _dafny.NewBuilder()
						_ = _coll3
						for _iter3 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr"))).Elements()); ; {
							_compr_3, _ok3 := _iter3()
							if !_ok3 {
								break
							}
							var _30_v0 _dafny.Sequence
							_30_v0 = interface{}(_compr_3).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr")), _30_v0) {
								_coll3.Add(_30_v0)
							}
						}
						return _coll3.ToSet()
					}()).Elements()); ; {
						_compr_2, _ok2 := _iter2()
						if !_ok2 {
							break
						}
						var _31_v1 _dafny.Sequence
						_31_v1 = interface{}(_compr_2).(_dafny.Sequence)
						if (func() _dafny.Set {
							var _coll4 = _dafny.NewBuilder()
							_ = _coll4
							for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr"))).Elements()); ; {
								_compr_4, _ok4 := _iter4()
								if !_ok4 {
									break
								}
								var _32_v0 _dafny.Sequence
								_32_v0 = interface{}(_compr_4).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr")), _32_v0) {
									_coll4.Add(_32_v0)
								}
							}
							return _coll4.ToSet()
						}()).Contains(_31_v1) {
							_coll2.Add(_31_v1)
						}
					}
					return _coll2.ToSet()
				}()).Elements()); ; {
					_compr_1, _ok1 := _iter1()
					if !_ok1 {
						break
					}
					var _33_v2 _dafny.Sequence
					_33_v2 = interface{}(_compr_1).(_dafny.Sequence)
					if (func() _dafny.Set {
						var _coll5 = _dafny.NewBuilder()
						_ = _coll5
						for _iter5 := _dafny.Iterate((func() _dafny.Set {
							var _coll6 = _dafny.NewBuilder()
							_ = _coll6
							for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr"))).Elements()); ; {
								_compr_6, _ok6 := _iter6()
								if !_ok6 {
									break
								}
								var _34_v0 _dafny.Sequence
								_34_v0 = interface{}(_compr_6).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr")), _34_v0) {
									_coll6.Add(_34_v0)
								}
							}
							return _coll6.ToSet()
						}()).Elements()); ; {
							_compr_5, _ok5 := _iter5()
							if !_ok5 {
								break
							}
							var _35_v1 _dafny.Sequence
							_35_v1 = interface{}(_compr_5).(_dafny.Sequence)
							if (func() _dafny.Set {
								var _coll7 = _dafny.NewBuilder()
								_ = _coll7
								for _iter7 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr"))).Elements()); ; {
									_compr_7, _ok7 := _iter7()
									if !_ok7 {
										break
									}
									var _36_v0 _dafny.Sequence
									_36_v0 = interface{}(_compr_7).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr")), _36_v0) {
										_coll7.Add(_36_v0)
									}
								}
								return _coll7.ToSet()
							}()).Contains(_35_v1) {
								_coll5.Add(_35_v1)
							}
						}
						return _coll5.ToSet()
					}()).Contains(_33_v2) {
						_coll1.Add(_33_v2)
					}
				}
				return _coll1.ToSet()
			}()).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _37_v3 _dafny.Sequence
				_37_v3 = interface{}(_compr_0).(_dafny.Sequence)
				if (func() _dafny.Set {
					var _coll8 = _dafny.NewBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate((func() _dafny.Set {
						var _coll9 = _dafny.NewBuilder()
						_ = _coll9
						for _iter9 := _dafny.Iterate((func() _dafny.Set {
							var _coll10 = _dafny.NewBuilder()
							_ = _coll10
							for _iter10 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr"))).Elements()); ; {
								_compr_10, _ok10 := _iter10()
								if !_ok10 {
									break
								}
								var _38_v0 _dafny.Sequence
								_38_v0 = interface{}(_compr_10).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr")), _38_v0) {
									_coll10.Add(_38_v0)
								}
							}
							return _coll10.ToSet()
						}()).Elements()); ; {
							_compr_9, _ok9 := _iter9()
							if !_ok9 {
								break
							}
							var _39_v1 _dafny.Sequence
							_39_v1 = interface{}(_compr_9).(_dafny.Sequence)
							if (func() _dafny.Set {
								var _coll11 = _dafny.NewBuilder()
								_ = _coll11
								for _iter11 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr"))).Elements()); ; {
									_compr_11, _ok11 := _iter11()
									if !_ok11 {
										break
									}
									var _40_v0 _dafny.Sequence
									_40_v0 = interface{}(_compr_11).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr")), _40_v0) {
										_coll11.Add(_40_v0)
									}
								}
								return _coll11.ToSet()
							}()).Contains(_39_v1) {
								_coll9.Add(_39_v1)
							}
						}
						return _coll9.ToSet()
					}()).Elements()); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _41_v2 _dafny.Sequence
						_41_v2 = interface{}(_compr_8).(_dafny.Sequence)
						if (func() _dafny.Set {
							var _coll12 = _dafny.NewBuilder()
							_ = _coll12
							for _iter12 := _dafny.Iterate((func() _dafny.Set {
								var _coll13 = _dafny.NewBuilder()
								_ = _coll13
								for _iter13 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr"))).Elements()); ; {
									_compr_13, _ok13 := _iter13()
									if !_ok13 {
										break
									}
									var _42_v0 _dafny.Sequence
									_42_v0 = interface{}(_compr_13).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr")), _42_v0) {
										_coll13.Add(_42_v0)
									}
								}
								return _coll13.ToSet()
							}()).Elements()); ; {
								_compr_12, _ok12 := _iter12()
								if !_ok12 {
									break
								}
								var _43_v1 _dafny.Sequence
								_43_v1 = interface{}(_compr_12).(_dafny.Sequence)
								if (func() _dafny.Set {
									var _coll14 = _dafny.NewBuilder()
									_ = _coll14
									for _iter14 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr"))).Elements()); ; {
										_compr_14, _ok14 := _iter14()
										if !_ok14 {
											break
										}
										var _44_v0 _dafny.Sequence
										_44_v0 = interface{}(_compr_14).(_dafny.Sequence)
										if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pisbhaerb"), _dafny.UnicodeSeqOfUtf8Bytes("ejfkrj"), _dafny.UnicodeSeqOfUtf8Bytes("qqubr")), _44_v0) {
											_coll14.Add(_44_v0)
										}
									}
									return _coll14.ToSet()
								}()).Contains(_43_v1) {
									_coll12.Add(_43_v1)
								}
							}
							return _coll12.ToSet()
						}()).Contains(_41_v2) {
							_coll8.Add(_41_v2)
						}
					}
					return _coll8.ToSet()
				}()).Contains(_37_v3) {
					_coll0.Add(_37_v3)
				}
			}
			return _coll0.ToSet()
		}()).Cardinality())
	} else if _source1.Is_DC28() {
		var _45___mcc_h8 bool = _source1.Get_().(D8_DC28).Cf53
		_ = _45___mcc_h8
		var _46___mcc_h9 _dafny.Int = _source1.Get_().(D8_DC28).Cf54
		_ = _46___mcc_h9
		var _47_cf54 _dafny.Int = _46___mcc_h9
		_ = _47_cf54
		var _48_cf53 bool = _45___mcc_h8
		_ = _48_cf53
		return Companion_D0_.Create_DC1_(_48_cf53, _48_cf53, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_cf54, _48_cf53)).Cardinality())
	} else if _source1.Is_DC25() {
		var _49___mcc_h10 _dafny.Map = _source1.Get_().(D8_DC25).Cf44
		_ = _49___mcc_h10
		var _50_cf44 _dafny.Map = _49___mcc_h10
		_ = _50_cf44
		if false {
			return Companion_D0_.Create_DC1_(false, false, _dafny.IntOfInt64(770))
		} else {
			return Companion_D0_.Create_DC1_(true, false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dbqyvcng")).Cardinality()))
		}
	} else {
		var _51___mcc_h11 D8 = _source1.Get_().(D8_DC29).Cf55
		_ = _51___mcc_h11
		var _52_cf55 D8 = _51___mcc_h11
		_ = _52_cf55
		return Companion_D0_.Create_DC1_(false, true, _dafny.IntOfInt64(859))
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Set, p1 _dafny.Sequence, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(true, (_dafny.MultiSetOf(_dafny.CodePoint('x'), _dafny.CodePoint('c'), _dafny.CodePoint('t'))).IsSubsetOf(_dafny.MultiSetOf(_dafny.CodePoint('i'))), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(125)), _dafny.SeqOf(_dafny.IntOfInt64(294)))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D18_.Create_DC55_((Companion_D18_.Create_DC55_(_dafny.UnicodeSeqOfUtf8Bytes("qjvd"), true)).Dtor_cf104(), true)).Dtor_cf104()
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, globalState *GlobalState) D4 {
	var _source2 D13 = Companion_D13_.Create_DC43_(Companion_D13_.Create_DC40_(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true, false))))
	_ = _source2
	if _source2.Is_DC41() {
		var _53___mcc_h0 _dafny.Int = _source2.Get_().(D13_DC41).Cf76
		_ = _53___mcc_h0
		var _54_cf76 _dafny.Int = _53___mcc_h0
		_ = _54_cf76
		return Companion_D4_.Create_DC11_(_dafny.SeqOf(false, true, false, false))
	} else if _source2.Is_DC42() {
		var _55___mcc_h1 _dafny.Sequence = _source2.Get_().(D13_DC42).Cf77
		_ = _55___mcc_h1
		var _56___mcc_h2 _dafny.MultiSet = _source2.Get_().(D13_DC42).Cf78
		_ = _56___mcc_h2
		var _57___mcc_h3 _dafny.Sequence = _source2.Get_().(D13_DC42).Cf79
		_ = _57___mcc_h3
		var _58___mcc_h4 bool = _source2.Get_().(D13_DC42).Cf80
		_ = _58___mcc_h4
		var _59___mcc_h5 _dafny.Sequence = _source2.Get_().(D13_DC42).Cf81
		_ = _59___mcc_h5
		var _60_cf81 _dafny.Sequence = _59___mcc_h5
		_ = _60_cf81
		var _61_cf80 bool = _58___mcc_h4
		_ = _61_cf80
		var _62_cf79 _dafny.Sequence = _57___mcc_h3
		_ = _62_cf79
		var _63_cf78 _dafny.MultiSet = _56___mcc_h2
		_ = _63_cf78
		var _64_cf77 _dafny.Sequence = _55___mcc_h1
		_ = _64_cf77
		return Companion_D4_.Create_DC11_(_64_cf77)
	} else if _source2.Is_DC40() {
		var _65___mcc_h6 _dafny.MultiSet = _source2.Get_().(D13_DC40).Cf75
		_ = _65___mcc_h6
		var _66_cf75 _dafny.MultiSet = _65___mcc_h6
		_ = _66_cf75
		return Companion_D4_.Create_DC11_(_dafny.SeqOf(true, true))
	} else {
		var _67___mcc_h7 D13 = _source2.Get_().(D13_DC43).Cf82
		_ = _67___mcc_h7
		var _68_cf82 D13 = _67___mcc_h7
		_ = _68_cf82
		return Companion_D4_.Create_DC11_(_dafny.SeqOf(false))
	}
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qcx"), _dafny.UnicodeSeqOfUtf8Bytes("yykt")), _dafny.UnicodeSeqOfUtf8Bytes("maxyot"))
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wcathxy"), _dafny.UnicodeSeqOfUtf8Bytes("ok"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(626))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_69_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})), _dafny.UnicodeSeqOfUtf8Bytes("epmmm"), _dafny.UnicodeSeqOfUtf8Bytes("lvjl")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-381))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_70_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(927))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_71_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))
	})))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('x')
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC12_()
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Sequence, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC7_(_dafny.IntOfInt64(46), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(940))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_72_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(415)
	}))).Cardinality()), (Companion_D2_.Create_DC7_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(69), true)).Cardinality(), _dafny.IntOfInt64(258), _dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.MultiSetOf(false)).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(true, false)).Cardinality(), _dafny.IntOfInt64(914))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_73_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))).Cardinality()))).Dtor_cf18(), _dafny.IntOfInt64(683), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("umst"), _dafny.UnicodeSeqOfUtf8Bytes("h"))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, globalState *GlobalState) _dafny.Map {
	var _source3 D1 = Companion_D1_.Create_DC4_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(122), !(true)), true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(88), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetOf(true, true, false)).Cardinality())).Cardinality())).Cardinality()), true)
	_ = _source3
	if _source3.Is_DC4() {
		var _74___mcc_h0 _dafny.Map = _source3.Get_().(D1_DC4).Cf9
		_ = _74___mcc_h0
		var _75___mcc_h1 bool = _source3.Get_().(D1_DC4).Cf10
		_ = _75___mcc_h1
		var _76___mcc_h2 _dafny.Int = _source3.Get_().(D1_DC4).Cf11
		_ = _76___mcc_h2
		var _77___mcc_h3 bool = _source3.Get_().(D1_DC4).Cf12
		_ = _77___mcc_h3
		var _78_cf12 bool = _77___mcc_h3
		_ = _78_cf12
		var _79_cf11 _dafny.Int = _76___mcc_h2
		_ = _79_cf11
		var _80_cf10 bool = _75___mcc_h1
		_ = _80_cf10
		var _81_cf9 _dafny.Map = _74___mcc_h0
		_ = _81_cf9
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_78_cf12, _78_cf12)).Cardinality()), _79_cf11)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(851), _79_cf11))
	} else {
		var _82___mcc_h4 _dafny.Sequence = _source3.Get_().(D1_DC3).Cf8
		_ = _82___mcc_h4
		var _83_cf8 _dafny.Sequence = _82___mcc_h4
		_ = _83_cf8
		return (func() _dafny.Map {
			var _coll15 = _dafny.NewMapBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(392), _dafny.IntOfInt64(89))); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _84_v0 _dafny.Int
				_84_v0 = interface{}(_compr_15).(_dafny.Int)
				if ((_dafny.IntOfInt64(392)).Cmp(_84_v0) <= 0) && ((_84_v0).Cmp(_dafny.IntOfInt64(89)) < 0) {
					_coll15.Add(Companion_Default___.SafeModuloInt(_84_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yy")).Cardinality())), _dafny.IntOfInt64(803))
				}
			}
			return _coll15.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()), (Companion_D5_.Create_DC17_(_dafny.IntOfInt64(258))).Dtor_cf35()))
	}
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 bool, p2 D5, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	var _source4 D6 = Companion_D6_.Create_DC20_(!(true), (Companion_D8_.Create_DC27_(true, true, (_dafny.SetOf(false)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("u")).Cardinality()))).Dtor_cf49())
	_ = _source4
	if _source4.Is_DC20() {
		var _85___mcc_h0 bool = _source4.Get_().(D6_DC20).Cf38
		_ = _85___mcc_h0
		var _86___mcc_h1 bool = _source4.Get_().(D6_DC20).Cf39
		_ = _86___mcc_h1
		var _87_cf39 bool = _86___mcc_h1
		_ = _87_cf39
		var _88_cf38 bool = _85___mcc_h0
		_ = _88_cf38
		return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-70), _dafny.IntOfInt64(693), (_dafny.Zero).Minus(_dafny.IntOfInt64(-986)))).Cardinality()))
	} else if _source4.Is_DC19() {
		var _89___mcc_h2 _dafny.Array = _source4.Get_().(D6_DC19).Cf37
		_ = _89___mcc_h2
		var _90_cf37 _dafny.Array = _89___mcc_h2
		_ = _90_cf37
		return _dafny.MultiSetOf(_dafny.IntOfInt64(-916), _dafny.IntOfInt64(212))
	} else {
		var _91___mcc_h3 D6 = _source4.Get_().(D6_DC21).Cf40
		_ = _91___mcc_h3
		var _92_cf40 D6 = _91___mcc_h3
		_ = _92_cf40
		return _dafny.MultiSetOf(_dafny.IntOfInt64(772))
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D5 {
	var _source5 D18 = Companion_D18_.Create_DC53_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("ujbhhlxdv")))
	_ = _source5
	if _source5.Is_DC54() {
		var _93___mcc_h0 bool = _source5.Get_().(D18_DC54).Cf100
		_ = _93___mcc_h0
		var _94___mcc_h1 bool = _source5.Get_().(D18_DC54).Cf101
		_ = _94___mcc_h1
		var _95___mcc_h2 _dafny.Int = _source5.Get_().(D18_DC54).Cf102
		_ = _95___mcc_h2
		var _96___mcc_h3 _dafny.Int = _source5.Get_().(D18_DC54).Cf103
		_ = _96___mcc_h3
		var _97_cf103 _dafny.Int = _96___mcc_h3
		_ = _97_cf103
		var _98_cf102 _dafny.Int = _95___mcc_h2
		_ = _98_cf102
		var _99_cf101 bool = _94___mcc_h1
		_ = _99_cf101
		var _100_cf100 bool = _93___mcc_h0
		_ = _100_cf100
		return Companion_D5_.Create_DC17_(_98_cf102)
	} else if _source5.Is_DC55() {
		var _101___mcc_h4 _dafny.Sequence = _source5.Get_().(D18_DC55).Cf104
		_ = _101___mcc_h4
		var _102___mcc_h5 bool = _source5.Get_().(D18_DC55).Cf105
		_ = _102___mcc_h5
		var _103_cf105 bool = _102___mcc_h5
		_ = _103_cf105
		var _104_cf104 _dafny.Sequence = _101___mcc_h4
		_ = _104_cf104
		return Companion_D5_.Create_DC17_(_dafny.IntOfUint32((_dafny.SeqOf(_103_cf105, _103_cf105, _103_cf105, _103_cf105, _103_cf105)).Cardinality()))
	} else {
		var _105___mcc_h6 _dafny.Map = _source5.Get_().(D18_DC53).Cf99
		_ = _105___mcc_h6
		var _106_cf99 _dafny.Map = _105___mcc_h6
		_ = _106_cf99
		return Companion_D5_.Create_DC17_((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()), _dafny.IntOfInt64(-18), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(75))).Cardinality(), !(false))).Cardinality(), (_dafny.Zero).Minus((_dafny.SetOf(!(true), false)).Cardinality()))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC26_(true, !((_dafny.MultiSetOf(_dafny.IntOfInt64(449))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-443))).Cardinality())))), (_dafny.SetOf(true)).IsProperSubsetOf(_dafny.SetOf(!(false), false)), _dafny.IntOfInt64(285))
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("whjhxs")).Cardinality()), _dafny.IntOfInt64(378))).Cardinality()), ((_dafny.SetOf(!(false))).Difference(_dafny.SetOf(false, false, true))).Cardinality(), ((_dafny.SetOf(false)).Cardinality()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bcax")).Cardinality()), func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-957), _dafny.IntOfInt64(399))); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _107_v0 _dafny.Int
			_107_v0 = interface{}(_compr_16).(_dafny.Int)
			if ((_dafny.IntOfInt64(-957)).Cmp(_107_v0) <= 0) && ((_107_v0).Cmp(_dafny.IntOfInt64(399)) < 0) {
				_coll16.Add(Companion_Default___.SafeDivisionInt(_107_v0, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), false)
			}
		}
		return _coll16.ToMap()
	}())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D4_.Create_DC13_(false, _dafny.MultiSetOf(_dafny.CodePoint('q')), false)).Dtor_cf31(), _dafny.IntOfInt64(-621))).Cardinality(), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xb")).Cardinality())).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC13_((_dafny.IntOfInt64(-795)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(853))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_108_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	}))).Cardinality())) >= 0, _dafny.MultiSetOf(_dafny.CodePoint('k')), _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("ljkdpw"), _dafny.UnicodeSeqOfUtf8Bytes("hrxq")))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(612)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(35))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_109_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(207)
	})))).Cardinality()))))), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Cardinality(), _dafny.IntOfInt64(-416))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(978), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cltxgkd")).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(648), _dafny.IntOfInt64(-694), _dafny.IntOfInt64(372), _dafny.IntOfInt64(280))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 bool, p2 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(true, (func() bool {
		if false {
			return false
		}
		return true
	})(), true, true)
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC29_(Companion_D8_.Create_DC28_((Companion_D19_.Create_DC58_(false)).Dtor_cf107(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf(false), _dafny.SetOf(!(false)))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	if !(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("s")).Cardinality())))).Contains(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xeh")).Cardinality())), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D10_.Create_DC33_(false), Companion_D10_.Create_DC33_(true), Companion_D10_.Create_DC33_(!(true)))).Cardinality()))).Cardinality())).Cardinality())) {
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf((_dafny.MultiSetOf((_dafny.SetOf(true, !(false), false)).Cardinality())).Cardinality())))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(987))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_110_i0 _dafny.Int) _dafny.Int {
			return (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(251), _dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()))).Cardinality(), (_dafny.SetOf(!(false))).Cardinality())).Cardinality()
		})), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(422))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(677))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_111_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(691)
		})), _dafny.SeqOf(_dafny.IntOfInt64(672)), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(54))).Uint32(), func(coer13 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_112_i2 _dafny.Int) D0 {
			return Companion_D0_.Create_DC1_(!(!(false)), true, (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D4_.Create_DC12_())).Cardinality()))).Cardinality())
		}))))).Cardinality(), _dafny.IntOfInt64(911))).Cardinality()))))
	} else {
		return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf((func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(389), _dafny.IntOfInt64(925))); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _113_v0 _dafny.Int
				_113_v0 = interface{}(_compr_17).(_dafny.Int)
				if ((_dafny.IntOfInt64(389)).Cmp(_113_v0) <= 0) && ((_113_v0).Cmp(_dafny.IntOfInt64(925)) < 0) {
					_coll17.Add(Companion_Default___.SafeModuloInt(_113_v0, (_dafny.SetOf(_dafny.IntOfInt64(-405), _dafny.IntOfInt64(630), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
						var _coll18 = _dafny.NewMapBuilder()
						_ = _coll18
						for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-370), _dafny.IntOfInt64(149))); ; {
							_compr_18, _ok18 := _iter18()
							if !_ok18 {
								break
							}
							var _114_v1 _dafny.Int
							_114_v1 = interface{}(_compr_18).(_dafny.Int)
							if ((_dafny.IntOfInt64(-370)).Cmp(_114_v1) <= 0) && ((_114_v1).Cmp(_dafny.IntOfInt64(149)) < 0) {
								_coll18.Add((_114_v1).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jegrswki")).Cardinality()))).Cardinality())), false)
							}
						}
						return _coll18.ToMap()
					}()).Cardinality(), false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false))).Cardinality()))).Cardinality()), _dafny.CodePoint('p'))
				}
			}
			return _coll17.ToMap()
		}()).Cardinality())))).Difference(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(178), _dafny.IntOfInt64(522), _dafny.IntOfInt64(647), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-523), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf((_dafny.SetOf(true)).Cardinality())).Cardinality())).Cardinality()))).Cardinality()), _dafny.SeqOf((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(426)), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.Zero).Minus(_dafny.IntOfInt64(-36)))).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_115_i3 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(321)
		})), _dafny.SeqOf(_dafny.IntOfInt64(-821))))
	}
}
func (_static *CompanionStruct_Default___) Fm35(p0 D13, p1 bool, globalState *GlobalState) D8 {
	if false {
		if true {
			return Companion_D8_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(554))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_116_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			})), (_dafny.MultiSetOf(false)).Cardinality()))
		} else {
			return Companion_D8_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("oth"), _dafny.IntOfInt64(247)))
		}
	} else {
		return Companion_D8_.Create_DC25_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("bg"), (_dafny.Zero).Minus(_dafny.IntOfInt64(-997))))
	}
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(false, (_dafny.IntOfInt64(-539)).Cmp(_dafny.IntOfInt64(923)) >= 0, true, !(false))
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.CodePoint, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC10_(true, _dafny.CodePoint('v'), false, true, false)
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(_dafny.IntOfInt64(587), (_dafny.SetOf(true, !(false))).Cardinality(), _dafny.IntOfInt64(103))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) _dafny.Sequence {
	if _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("udj"), _dafny.UnicodeSeqOfUtf8Bytes("jllaki")) {
		return _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-482), _dafny.IntOfInt64(297)))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func() _dafny.Map {
			var _coll19 = _dafny.NewMapBuilder()
			_ = _coll19
			for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(836), _dafny.IntOfInt64(538))); ; {
				_compr_19, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _117_v0 _dafny.Int
				_117_v0 = interface{}(_compr_19).(_dafny.Int)
				if ((_dafny.IntOfInt64(836)).Cmp(_117_v0) <= 0) && ((_117_v0).Cmp(_dafny.IntOfInt64(538)) < 0) {
					_coll19.Add((_117_v0).Plus((_dafny.SetOf(false, false, true)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()))
				}
			}
			return _coll19.ToMap()
		}()), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D16_.Create_DC51_(false, _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetOf(false, false)).Cardinality())), true, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Dtor_cf97(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-228))).Cardinality()))))
	}
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.CodePoint, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) D5 {
	var _source6 D2 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_dafny.SeqOf(!(true), false)).Cardinality()), _dafny.IntOfInt64(475), _dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(941), (_dafny.MultiSetOf(_dafny.IntOfInt64(197))).Cardinality())), _dafny.IntOfInt64(189), (_dafny.MultiSetOf(false, true)).Cardinality()))
	_ = _source6
	if _source6.Is_DC6() {
		var _118___mcc_h0 _dafny.CodePoint = _source6.Get_().(D2_DC6).Cf14
		_ = _118___mcc_h0
		var _119___mcc_h1 _dafny.Int = _source6.Get_().(D2_DC6).Cf15
		_ = _119___mcc_h1
		var _120_cf15 _dafny.Int = _119___mcc_h1
		_ = _120_cf15
		var _121_cf14 _dafny.CodePoint = _118___mcc_h0
		_ = _121_cf14
		return Companion_D5_.Create_DC15_(Companion_D4_.Create_DC11_(_dafny.SeqOf(true)), true)
	} else if _source6.Is_DC7() {
		var _122___mcc_h2 _dafny.Int = _source6.Get_().(D2_DC7).Cf16
		_ = _122___mcc_h2
		var _123___mcc_h3 _dafny.Int = _source6.Get_().(D2_DC7).Cf17
		_ = _123___mcc_h3
		var _124___mcc_h4 _dafny.MultiSet = _source6.Get_().(D2_DC7).Cf18
		_ = _124___mcc_h4
		var _125___mcc_h5 _dafny.Int = _source6.Get_().(D2_DC7).Cf19
		_ = _125___mcc_h5
		var _126___mcc_h6 _dafny.Int = _source6.Get_().(D2_DC7).Cf20
		_ = _126___mcc_h6
		var _127_cf20 _dafny.Int = _126___mcc_h6
		_ = _127_cf20
		var _128_cf19 _dafny.Int = _125___mcc_h5
		_ = _128_cf19
		var _129_cf18 _dafny.MultiSet = _124___mcc_h4
		_ = _129_cf18
		var _130_cf17 _dafny.Int = _123___mcc_h3
		_ = _130_cf17
		var _131_cf16 _dafny.Int = _122___mcc_h2
		_ = _131_cf16
		return Companion_D5_.Create_DC15_(Companion_D4_.Create_DC11_(_dafny.SeqOf(false, true)), !(false))
	} else if _source6.Is_DC5() {
		var _132___mcc_h7 _dafny.Sequence = _source6.Get_().(D2_DC5).Cf13
		_ = _132___mcc_h7
		var _133_cf13 _dafny.Sequence = _132___mcc_h7
		_ = _133_cf13
		return Companion_D5_.Create_DC15_(Companion_D4_.Create_DC11_(_dafny.SeqOf(false, true)), false)
	} else {
		var _134___mcc_h8 D2 = _source6.Get_().(D2_DC8).Cf21
		_ = _134___mcc_h8
		var _135_cf21 D2 = _134___mcc_h8
		_ = _135_cf21
		return Companion_D5_.Create_DC15_(Companion_D4_.Create_DC11_(_dafny.SeqOf(false)), true)
	}
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((Companion_D20_.Create_DC60_(_dafny.SetOf(_dafny.CodePoint('x'), _dafny.CodePoint('i')))).Dtor_cf109(), _dafny.SetOf(_dafny.CodePoint('s')), (_dafny.SetOf(_dafny.CodePoint('a'))).Union(_dafny.SetOf(_dafny.CodePoint('q'))), func() _dafny.Set {
		var _coll20 = _dafny.NewBuilder()
		_ = _coll20
		for _iter20 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('r'))).Elements()); ; {
			_compr_20, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _136_v0 _dafny.CodePoint
			_136_v0 = interface{}(_compr_20).(_dafny.CodePoint)
			if (_dafny.SetOf(_dafny.CodePoint('r'))).Contains(_136_v0) {
				_coll20.Add(_136_v0)
			}
		}
		return _coll20.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) M6(globalState *GlobalState) (bool, _dafny.Map) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.Map = _dafny.EmptyMap
	_ = r1
	var _137_v0 _dafny.Int
	_ = _137_v0
	_137_v0 = _dafny.IntOfInt64(387)
	var _138_v1 bool
	_ = _138_v1
	_138_v1 = true
	var _139_v2 _dafny.Map
	_ = _139_v2
	_139_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v1, _137_v0)
	var _140_v3 _dafny.Array
	_ = _140_v3
	var _nwElement0_0 _dafny.Map = _139_v2
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.One)
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	_140_v3 = _nw0
	var _141_v4 _dafny.Array
	_ = _141_v4
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(4)
	_ = _len0_0
	var _nw1 _dafny.Array
	_ = _nw1
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw1 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = (func(_142_v1 bool) func(_dafny.Int) bool {
			return func(_143_i1 _dafny.Int) bool {
				return _142_v1
			}
		})(_138_v1)
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
	_141_v4 = _nw1
	var _144_v5 T0
	_ = _144_v5
	var _nw2 *C2 = New_C2_()
	_ = _nw2
	_nw2.Ctor__(true, _137_v0, _140_v3, _138_v1, _141_v4)
	_144_v5 = _nw2
	var _145_v6 _dafny.MultiSet
	_ = _145_v6
	_145_v6 = _dafny.MultiSetOf(_144_v5, _144_v5, _144_v5, _144_v5)
	var _146_v7 D7
	_ = _146_v7
	_146_v7 = Companion_D7_.Create_DC23_(_137_v0)
	var _hi0 _dafny.Int = ((_146_v7).Dtor_cf42()).Minus(_137_v0)
	_ = _hi0
	for _147_i0 := (_145_v6).Cardinality(); _147_i0.Cmp(_hi0) < 0; _147_i0 = _147_i0.Plus(_dafny.One) {
		var _148_v8 _dafny.Array
		_ = _148_v8
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_1
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Int = func(_149_i2 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_149_i2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ussfu")).Cardinality()))
			}
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw3 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw3).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw3).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_148_v8 = _nw3
		_148_v8 = _148_v8
		_138_v1 = _138_v1
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))
		_ = _index0
		(_148_v8).ArraySet1(_dafny.IntOfInt64(-482), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))
		_ = _index1
		(_148_v8).ArraySet1(Companion_Default___.SafeModuloInt(_147_i0, _137_v0), (_index1).Int())
		var _150_v9 _dafny.MultiSet
		_ = _150_v9
		_150_v9 = _dafny.MultiSetOf((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(887))
		var _151_v10 _dafny.MultiSet
		_ = _151_v10
		_151_v10 = _dafny.MultiSetOf((_150_v9).Cardinality(), _147_i0)
		var _152_v11 _dafny.MultiSet
		_ = _152_v11
		_152_v11 = _dafny.MultiSetOf(_151_v10, _151_v10)
		var _153_v12 D2
		_ = _153_v12
		_153_v12 = Companion_D2_.Create_DC7_((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), (_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _152_v11, (_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-258))
		var _154_v13 D2
		_ = _154_v13
		_154_v13 = Companion_D2_.Create_DC8_(_153_v12)
		var _source7 D2 = (func() D2 {
			if _138_v1 {
				return (func() D2 {
					if _138_v1 {
						return _154_v13
					}
					return _154_v13
				})()
			}
			return _154_v13
		})()
		_ = _source7
		if _source7.Is_DC6() {
			var _155___mcc_h0 _dafny.CodePoint = _source7.Get_().(D2_DC6).Cf14
			_ = _155___mcc_h0
			var _156___mcc_h1 _dafny.Int = _source7.Get_().(D2_DC6).Cf15
			_ = _156___mcc_h1
			var _157_cf15 _dafny.Int = _156___mcc_h1
			_ = _157_cf15
			var _158_cf14 _dafny.CodePoint = _155___mcc_h0
			_ = _158_cf14
			_148_v8 = (func() _dafny.Array {
				if _138_v1 {
					return _148_v8
				}
				return _148_v8
			})()
			_157_cf15 = (_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int)
			_137_v0 = Companion_Default___.SafeDivisionInt((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _137_v0)
			var _159_v14 _dafny.Array
			_ = _159_v14
			var _nw4 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw4
			_159_v14 = _nw4
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_159_v14), 0))
			_ = _index2
			(_159_v14).ArraySet1(_144_v5, (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(766), _dafny.ArrayLen((_159_v14), 0))
			_ = _index3
			(_159_v14).ArraySet1(_144_v5, (_index3).Int())
		} else if _source7.Is_DC7() {
			var _160___mcc_h2 _dafny.Int = _source7.Get_().(D2_DC7).Cf16
			_ = _160___mcc_h2
			var _161___mcc_h3 _dafny.Int = _source7.Get_().(D2_DC7).Cf17
			_ = _161___mcc_h3
			var _162___mcc_h4 _dafny.MultiSet = _source7.Get_().(D2_DC7).Cf18
			_ = _162___mcc_h4
			var _163___mcc_h5 _dafny.Int = _source7.Get_().(D2_DC7).Cf19
			_ = _163___mcc_h5
			var _164___mcc_h6 _dafny.Int = _source7.Get_().(D2_DC7).Cf20
			_ = _164___mcc_h6
			var _165_cf20 _dafny.Int = _164___mcc_h6
			_ = _165_cf20
			var _166_cf19 _dafny.Int = _163___mcc_h5
			_ = _166_cf19
			var _167_cf18 _dafny.MultiSet = _162___mcc_h4
			_ = _167_cf18
			var _168_cf17 _dafny.Int = _161___mcc_h3
			_ = _168_cf17
			var _169_cf16 _dafny.Int = _160___mcc_h2
			_ = _169_cf16
			var _170_v15 _dafny.CodePoint
			_ = _170_v15
			_170_v15 = _dafny.CodePoint('i')
			_170_v15 = _170_v15
			var _171_v16 _dafny.Sequence
			_ = _171_v16
			_171_v16 = _dafny.SeqOf(_138_v1)
			var _172_v17 _dafny.Sequence
			_ = _172_v17
			_172_v17 = _dafny.SeqOf(_137_v0, _137_v0)
			_166_cf19 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(636), (func() _dafny.Int {
				if _138_v1 {
					return _137_v0
				}
				return (Companion_Default___.Fm26(_138_v1, _138_v1, Companion_Default___.Fm41(_170_v15, _171_v16, _172_v17, globalState), _166_cf19, globalState)).Cardinality()
			})())
			_168_cf17 = ((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int)).Times(_168_cf17)
			var _173_v18 *C0
			_ = _173_v18
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__()
			_173_v18 = _nw5
		} else if _source7.Is_DC5() {
			var _174___mcc_h7 _dafny.Sequence = _source7.Get_().(D2_DC5).Cf13
			_ = _174___mcc_h7
			var _175_cf13 _dafny.Sequence = _174___mcc_h7
			_ = _175_cf13
			var _176_v19 _dafny.Map
			_ = _176_v19
			_176_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _138_v1)
			var _177_v20 D1
			_ = _177_v20
			_177_v20 = Companion_D1_.Create_DC4_(_176_v19, !(_138_v1), (_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _138_v1)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))
			_ = _index4
			(_148_v8).ArraySet1((_177_v20).Dtor_cf11(), (_index4).Int())
			var _178_v21 D6
			_ = _178_v21
			_178_v21 = Companion_D6_.Create_DC20_(_138_v1, _138_v1)
			var _179_v22 _dafny.Set
			_ = _179_v22
			_179_v22 = _dafny.SetOf(_178_v21, _178_v21)
			var _180_v23 _dafny.Sequence
			_ = _180_v23
			_180_v23 = _dafny.SeqOf(_179_v22)
			var _181_v24 *C0
			_ = _181_v24
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__()
			_181_v24 = _nw6
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))
			_ = _index5
			var _rhs0 _dafny.Sequence = _180_v23
			_ = _rhs0
			var _rhs1 _dafny.Int = (((_139_v2).Merge(_139_v2)).Merge(_139_v2)).Cardinality()
			_ = _rhs1
			var _rhs2 _dafny.Array = _148_v8
			_ = _rhs2
			var _rhs3 _dafny.Int = (_dafny.Zero).Minus((_137_v0).Minus((_dafny.Zero).Minus((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int))))
			_ = _rhs3
			var _rhs4 _dafny.Int = Companion_Default___.SafeModuloInt(((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _181_v24)).Cardinality())).Plus(_147_i0), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), (_dafny.SetOf(_147_i0, _137_v0)).Cardinality())))
			_ = _rhs4
			var _lhs0 _dafny.Array = _148_v8
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))
			_ = _lhs1
			_180_v23 = _rhs0
			_137_v0 = _rhs1
			_148_v8 = _rhs2
			_137_v0 = _rhs3
			(_lhs0).ArraySet1(_rhs4, (_lhs1).Int())
			var _182_v25 _dafny.Set
			_ = _182_v25
			_182_v25 = _dafny.SetOf((_175_cf13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-185), _dafny.IntOfUint32((_175_cf13).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_138_v1 = (_dafny.IntOfUint32((_175_cf13).Cardinality())).Cmp(((_182_v25).Cardinality()).Minus(_137_v0)) < 0
			var _183_v26 _dafny.MultiSet
			_ = _183_v26
			_183_v26 = _dafny.MultiSetOf(_138_v1, _138_v1, _138_v1, _138_v1)
			var _184_v27 D8
			_ = _184_v27
			_184_v27 = Companion_D8_.Create_DC26_(_138_v1, _dafny.Companion_Sequence_.IsPrefixOf(_175_cf13, _175_cf13), ((_183_v26).Cardinality()).Cmp((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int)) < 0, (_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int))
			_184_v27 = _184_v27
		} else {
			var _185___mcc_h8 D2 = _source7.Get_().(D2_DC8).Cf21
			_ = _185___mcc_h8
			var _186_cf21 D2 = _185___mcc_h8
			_ = _186_cf21
			var _187_v28 _dafny.Set
			_ = _187_v28
			_187_v28 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int)), (Companion_Default___.SafeIndex((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int))).Cardinality()))).Uint32(), _147_i0), (Companion_Default___.SafeIndex(_147_i0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int)), (Companion_Default___.SafeIndex((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int))).Cardinality()))).Uint32(), _147_i0)).Cardinality()))).Uint32(), (_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int))).Cardinality()))
			_137_v0 = (func() _dafny.Int {
				if Companion_Default___.Fm4((_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int), _187_v28, globalState) {
					return (_148_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))).Int()).(_dafny.Int)
				}
				return (_144_v5).Fm2(globalState)
			})()
			var _188_v29 _dafny.Set
			_ = _188_v29
			_188_v29 = _dafny.SetOf(false)
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_148_v8), 0))
			_ = _index6
			(_148_v8).ArraySet1((_137_v0).Plus(Companion_Default___.SafeModuloInt((_188_v29).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_138_v1)).Cardinality()))), (_index6).Int())
			var _189_v30 D8
			_ = _189_v30
			_189_v30 = Companion_D8_.Create_DC28_(_138_v1, _dafny.IntOfInt64(-658))
			var _190_v31 D8
			_ = _190_v31
			_190_v31 = Companion_D8_.Create_DC29_(_189_v30)
			_190_v31 = _190_v31
			var _191_v32 _dafny.Sequence
			_ = _191_v32
			_191_v32 = _dafny.UnicodeSeqOfUtf8Bytes("yvdvmxg")
			var _192_v33 _dafny.Array
			_ = _192_v33
			var _nwElement0_1 _dafny.Sequence = _191_v32
			_ = _nwElement0_1
			var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(18))
			_ = _nw7
			(_nw7).ArraySet1(_nwElement0_1, 0)
			(_nw7).ArraySet1(_191_v32, 1)
			(_nw7).ArraySet1(_191_v32, 2)
			(_nw7).ArraySet1(_191_v32, 3)
			(_nw7).ArraySet1(_191_v32, 4)
			(_nw7).ArraySet1(_191_v32, 5)
			(_nw7).ArraySet1(_191_v32, 6)
			(_nw7).ArraySet1(_191_v32, 7)
			(_nw7).ArraySet1(_191_v32, 8)
			(_nw7).ArraySet1(_191_v32, 9)
			(_nw7).ArraySet1(_191_v32, 10)
			(_nw7).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("eiaukerpe"), 11)
			(_nw7).ArraySet1(_191_v32, 12)
			(_nw7).ArraySet1(_191_v32, 13)
			(_nw7).ArraySet1(_191_v32, 14)
			(_nw7).ArraySet1(_191_v32, 15)
			(_nw7).ArraySet1(_191_v32, 16)
			(_nw7).ArraySet1(_191_v32, 17)
			_192_v33 = _nw7
			var _193_v34 D19
			_ = _193_v34
			_193_v34 = Companion_D19_.Create_DC56_(_192_v33)
			var _pat_let_tv0 = _192_v33
			_ = _pat_let_tv0
			var _194_v35 _dafny.Array
			_ = _194_v35
			var _nwElement0_2 _dafny.Array = (func(_pat_let0_0 D19) D19 {
				return func(_195_dt__update__tmp_h0 D19) D19 {
					return func(_pat_let1_0 _dafny.Array) D19 {
						return func(_196_dt__update_hcf106_h0 _dafny.Array) D19 {
							return Companion_D19_.Create_DC56_(_196_dt__update_hcf106_h0)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_193_v34)).Dtor_cf106()
			_ = _nwElement0_2
			var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
			_ = _nw8
			(_nw8).ArraySet1(_nwElement0_2, 0)
			(_nw8).ArraySet1(_192_v33, 1)
			_194_v35 = _nw8
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_194_v35), 0))
			_ = _index7
			(_194_v35).ArraySet1(_192_v33, (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_194_v35), 0))
			_ = _index8
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_2
			var _nw9 _dafny.Array
			_ = _nw9
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw9 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_197_v32 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_198_i3 _dafny.Int) _dafny.Sequence {
						return _197_v32
					}
				})(_191_v32)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw9 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw9).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw9).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			(_194_v35).ArraySet1(_nw9, (_index8).Int())
		}
	}
	var _199_v36 _dafny.MultiSet
	_ = _199_v36
	_199_v36 = _dafny.MultiSetOf(!(_138_v1), _138_v1)
	var _200_v37 _dafny.Map
	_ = _200_v37
	_200_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_138_v1)).Equals(_199_v36), _138_v1)
	_200_v37 = (_200_v37).Update(true, false)
	var _hi1 _dafny.Int = _137_v0
	_ = _hi1
	for _201_i4 := _137_v0; _201_i4.Cmp(_hi1) < 0; _201_i4 = _201_i4.Plus(_dafny.One) {
		if _138_v1 {
			var _202_v38 _dafny.Array
			_ = _202_v38
			var _nwElement0_3 _dafny.Array = _144_v5.F2()
			_ = _nwElement0_3
			var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(5))
			_ = _nw10
			(_nw10).ArraySet1(_nwElement0_3, 0)
			(_nw10).ArraySet1(_144_v5.F2(), 1)
			(_nw10).ArraySet1(_141_v4, 2)
			(_nw10).ArraySet1(_144_v5.F2(), 3)
			(_nw10).ArraySet1(_141_v4, 4)
			_202_v38 = _nw10
			var _203_v39 D15
			_ = _203_v39
			_203_v39 = Companion_D15_.Create_DC47_(_202_v38)
			var _204_v40 D15
			_ = _204_v40
			_204_v40 = Companion_D15_.Create_DC49_(_203_v39)
			var _205_v41 D15
			_ = _205_v41
			_205_v41 = Companion_D15_.Create_DC49_(_203_v39)
			var _206_v42 _dafny.Sequence
			_ = _206_v42
			_206_v42 = _dafny.SeqOf(_205_v41)
			var _207_v43 *C4
			_ = _207_v43
			var _nw11 *C4 = New_C4_()
			_ = _nw11
			_nw11.Ctor__(!_dafny.Companion_Sequence_.Equal(_206_v42, _206_v42))
			_207_v43 = _nw11
			r0 = (_207_v43).F6()
			var _208_v44 _dafny.CodePoint
			_ = _208_v44
			_208_v44 = _dafny.CodePoint('t')
			var _209_v45 _dafny.Sequence
			_ = _209_v45
			_209_v45 = _dafny.SeqOf(_208_v44)
			var _210_v46 D13
			_ = _210_v46
			_210_v46 = Companion_D13_.Create_DC40_(_dafny.MultiSetOf((_207_v43).Fm6(_209_v45, (_207_v43).F6(), globalState), _138_v1))
			var _211_v47 _dafny.Set
			_ = _211_v47
			_211_v47 = _dafny.SetOf(_201_i4)
			r0 = !(Companion_Default___.Fm4(((_210_v46).Dtor_cf75()).Cardinality(), _211_v47, globalState))
			(globalState).F1 = _144_v5.F2()
			var _212_v48 *C2
			_ = _212_v48
			var _nw12 *C2 = New_C2_()
			_ = _nw12
			_nw12.Ctor__((_207_v43).F6(), _201_i4, _140_v3, _138_v1, _144_v5.F2())
			_212_v48 = _nw12
			var _213_v49 _dafny.Array
			_ = _213_v49
			var _nwElement0_4 *C2 = _212_v48
			_ = _nwElement0_4
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(17))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_4, 0)
			(_nw13).ArraySet1(_212_v48, 1)
			(_nw13).ArraySet1(_212_v48, 2)
			(_nw13).ArraySet1(_212_v48, 3)
			(_nw13).ArraySet1(_212_v48, 4)
			(_nw13).ArraySet1(_212_v48, 5)
			(_nw13).ArraySet1(_212_v48, 6)
			(_nw13).ArraySet1(_212_v48, 7)
			(_nw13).ArraySet1((func() *C2 {
				if (_207_v43).F6() {
					return _212_v48
				}
				return _212_v48
			})(), 8)
			(_nw13).ArraySet1(_212_v48, 9)
			(_nw13).ArraySet1(_212_v48, 10)
			(_nw13).ArraySet1(_212_v48, 11)
			(_nw13).ArraySet1(_212_v48, 12)
			(_nw13).ArraySet1(_212_v48, 13)
			(_nw13).ArraySet1(_212_v48, 14)
			(_nw13).ArraySet1(_212_v48, 15)
			(_nw13).ArraySet1(_212_v48, 16)
			_213_v49 = _nw13
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_213_v49), 0))
			_ = _index9
			(_213_v49).ArraySet1(_212_v48, (_index9).Int())
			var _214_v50 _dafny.Set
			_ = _214_v50
			_214_v50 = _dafny.SetOf(_138_v1)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_213_v49), 0))
			_ = _index10
			var _rhs5 _dafny.Int = _201_i4
			_ = _rhs5
			var _rhs6 *C2 = _212_v48
			_ = _rhs6
			var _rhs7 bool = (func() bool {
				if _212_v48.F7 {
					return (true) || (true)
				}
				return (_214_v50).IsDisjointFrom(_214_v50)
			})()
			_ = _rhs7
			var _rhs8 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_209_v45, _209_v45)
			_ = _rhs8
			var _lhs2 _dafny.Array = _213_v49
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_213_v49), 0))
			_ = _lhs3
			_137_v0 = _rhs5
			(_lhs2).ArraySet1(_rhs6, (_lhs3).Int())
			_138_v1 = _rhs7
			_209_v45 = _rhs8
		} else {
			var _215_v51 _dafny.Array
			_ = _215_v51
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
			_ = _nw14
			_215_v51 = _nw14
			var _216_v52 _dafny.CodePoint
			_ = _216_v52
			_216_v52 = _dafny.CodePoint('p')
			var _217_v53 _dafny.Set
			_ = _217_v53
			_217_v53 = _dafny.SetOf(_216_v52, _216_v52, _216_v52)
			var _218_v54 _dafny.Sequence
			_ = _218_v54
			_218_v54 = _dafny.SeqOf(_217_v53)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_215_v51), 0))
			_ = _index11
			(_215_v51).ArraySet1(_218_v54, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_215_v51), 0))
			_ = _index12
			(_215_v51).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_218_v54, Companion_Default___.Fm42(_137_v0, _138_v1, _138_v1, globalState)), (_index12).Int())
			_138_v1 = !(_138_v1)
			_216_v52 = _216_v52
			_137_v0 = _137_v0
			var _219_v55 _dafny.Sequence
			_ = _219_v55
			_219_v55 = _dafny.UnicodeSeqOfUtf8Bytes("fc")
			r0 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_219_v55).Cardinality()), _137_v0)).Cmp(_dafny.IntOfInt64(920)) >= 0
		}
		var _220_v56 _dafny.Array
		_ = _220_v56
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_3
		var _nw15 _dafny.Array
		_ = _nw15
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw15 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Sequence = (func(_221_i4 _dafny.Int, _222_v1 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_223_i5 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_221_i4), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_222_v1, _222_v1)).Cardinality())))
				}
			})(_201_i4, _138_v1)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw15 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw15).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw15).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_220_v56 = _nw15
		var _224_v57 _dafny.Sequence
		_ = _224_v57
		_224_v57 = _dafny.SeqOf(_201_i4)
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_220_v56), 0))
		_ = _index13
		(_220_v56).ArraySet1(_224_v57, (_index13).Int())
		var _225_v58 _dafny.Array
		_ = _225_v58
		var _nw16 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
		_ = _nw16
		_225_v58 = _nw16
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_220_v56), 0))
		_ = _index14
		var _rhs9 bool = _138_v1
		_ = _rhs9
		var _rhs10 _dafny.Sequence = _224_v57
		_ = _rhs10
		var _rhs11 _dafny.Array = _225_v58
		_ = _rhs11
		var _lhs4 _dafny.Array = _220_v56
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_220_v56), 0))
		_ = _lhs5
		r0 = _rhs9
		(_lhs4).ArraySet1(_rhs10, (_lhs5).Int())
		_225_v58 = _rhs11
		_138_v1 = _138_v1
		var _226_v59 D14
		_ = _226_v59
		_226_v59 = Companion_D14_.Create_DC45_(_138_v1)
		_226_v59 = Companion_D14_.Create_DC45_(_138_v1)
	}
	var _227_v60 _dafny.Array
	_ = _227_v60
	var _nw17 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
	_ = _nw17
	_227_v60 = _nw17
	var _228_v61 *C0
	_ = _228_v61
	var _nw18 *C0 = New_C0_()
	_ = _nw18
	_nw18.Ctor__()
	_228_v61 = _nw18
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_227_v60), 0))
	_ = _index15
	(_227_v60).ArraySet1(_228_v61, (_index15).Int())
	var _229_v62 D8
	_ = _229_v62
	_229_v62 = Companion_D8_.Create_DC27_(_138_v1, _138_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("begh")).Cardinality()), (_137_v0).Minus(_dafny.IntOfInt64(416)))
	var _230_v63 _dafny.Array
	_ = _230_v63
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(19)
	_ = _len0_4
	var _nw19 _dafny.Array
	_ = _nw19
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw19 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Int = func(_231_i6 _dafny.Int) _dafny.Int {
			return (_231_i6).Times(_dafny.IntOfInt64(430))
		}
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw19 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw19).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw19).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_230_v63 = _nw19
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_230_v63), 0))
	_ = _index16
	(_230_v63).ArraySet1(_137_v0, (_index16).Int())
	var _232_v64 _dafny.Map
	_ = _232_v64
	_232_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(194), _141_v4)
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_227_v60), 0))
	_ = _index17
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_230_v63), 0))
	_ = _index18
	var _rhs12 *C0 = _228_v61
	_ = _rhs12
	var _rhs13 D8 = _229_v62
	_ = _rhs13
	var _rhs14 _dafny.Int = Companion_Default___.SafeDivisionInt(((_232_v64).Update(_dafny.IntOfUint32((Companion_Default___.Fm18(_138_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, true), globalState)).Cardinality()), _144_v5.F2())).Cardinality(), (_dafny.Zero).Minus(_137_v0))
	_ = _rhs14
	var _lhs6 _dafny.Array = _227_v60
	_ = _lhs6
	var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_227_v60), 0))
	_ = _lhs7
	var _lhs8 _dafny.Array = _230_v63
	_ = _lhs8
	var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_230_v63), 0))
	_ = _lhs9
	(_lhs6).ArraySet1(_rhs12, (_lhs7).Int())
	_229_v62 = _rhs13
	(_lhs8).ArraySet1(_rhs14, (_lhs9).Int())
	var _233_v65 _dafny.CodePoint
	_ = _233_v65
	_233_v65 = _dafny.CodePoint('c')
	_233_v65 = _233_v65
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_230_v63), 0))
	_ = _index19
	(_230_v63).ArraySet1((_230_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_230_v63), 0))).Int()).(_dafny.Int), (_index19).Int())
	var _234_v66 D3
	_ = _234_v66
	_234_v66 = Companion_D3_.Create_DC10_(_138_v1, _233_v65, _138_v1, _138_v1, _138_v1)
	var _235_v67 _dafny.Map
	_ = _235_v67
	_235_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_234_v66, _dafny.IntOfInt64(324))
	var _pat_let_tv1 = _138_v1
	_ = _pat_let_tv1
	var _pat_let_tv2 = _138_v1
	_ = _pat_let_tv2
	r0 = (func() _dafny.Set {
		var _coll21 = _dafny.NewBuilder()
		_ = _coll21
		for _iter21 := _dafny.Iterate((_235_v67).Keys().Elements()); ; {
			_compr_21, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _236_v68 D3
			_236_v68 = interface{}(_compr_21).(D3)
			if (_235_v67).Contains(_236_v68) {
				_coll21.Add(_236_v68)
			}
		}
		return _coll21.ToSet()
	}()).Contains(func(_pat_let2_0 D3) D3 {
		return func(_237_dt__update__tmp_h1 D3) D3 {
			return func(_pat_let3_0 bool) D3 {
				return func(_238_dt__update_hcf26_h0 bool) D3 {
					return func(_pat_let4_0 bool) D3 {
						return func(_239_dt__update_hcf23_h0 bool) D3 {
							return Companion_D3_.Create_DC10_(_239_dt__update_hcf23_h0, (_237_dt__update__tmp_h1).Dtor_cf24(), (_237_dt__update__tmp_h1).Dtor_cf25(), _238_dt__update_hcf26_h0, (_237_dt__update__tmp_h1).Dtor_cf27())
						}(_pat_let4_0)
					}(_pat_let_tv2)
				}(_pat_let3_0)
			}(_pat_let_tv1)
		}(_pat_let2_0)
	}(_234_v66))
	var _240_v70 _dafny.Map
	_ = _240_v70
	_240_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), _139_v2)
	r1 = func() _dafny.Map {
		var _coll22 = _dafny.NewMapBuilder()
		_ = _coll22
		for _iter22 := _dafny.Iterate((_240_v70).Keys().Elements()); ; {
			_compr_22, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _241_v69 _dafny.MultiSet
			_241_v69 = interface{}(_compr_22).(_dafny.MultiSet)
			if (_240_v70).Contains(_241_v69) {
				_coll22.Add(_241_v69, (_138_v1) || (true))
			}
		}
		return _coll22.ToMap()
	}()
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _242_v0 bool
	_ = _242_v0
	_242_v0 = true
	var _243_v1 _dafny.Set
	_ = _243_v1
	_243_v1 = _dafny.SetOf(_242_v0, _242_v0, _242_v0, _242_v0)
	var _244_v2 _dafny.Array
	_ = _244_v2
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(14)
	_ = _len0_5
	var _nw20 _dafny.Array
	_ = _nw20
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw20 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) bool = func(_245_i0 _dafny.Int) bool {
			return false
		}
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw20 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw20).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw20).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_244_v2 = _nw20
	var _246_globalState *GlobalState
	_ = _246_globalState
	var _nw21 *GlobalState = New_GlobalState_()
	_ = _nw21
	_nw21.Ctor__(_243_v1, _244_v2)
	_246_globalState = _nw21
	var _247_v3 _dafny.Int
	_ = _247_v3
	_247_v3 = _dafny.IntOfInt64(923)
	var _hi2 _dafny.Int = _247_v3
	_ = _hi2
	for _248_i1 := _247_v3; _248_i1.Cmp(_hi2) < 0; _248_i1 = _248_i1.Plus(_dafny.One) {
		var _249_v4 _dafny.Map
		_ = _249_v4
		_249_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_247_v3), _248_i1)
		var _250_v5 _dafny.Sequence
		_ = _250_v5
		_250_v5 = _dafny.SeqOf(Companion_Default___.Fm0((_dafny.Zero).Minus(_247_v3), _246_globalState), _247_v3, _247_v3, (func() _dafny.Int {
			if (_249_v4).Contains(_247_v3) {
				return (_249_v4).Get(_247_v3).(_dafny.Int)
			}
			return _247_v3
		})(), _247_v3)
		_247_v3 = _dafny.IntOfUint32((_250_v5).Cardinality())
		var _251_v6 _dafny.Array
		_ = _251_v6
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw22
		_251_v6 = _nw22
		var _252_v7 _dafny.Map
		_ = _252_v7
		_252_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_251_v6, _dafny.IntOfUint32((_dafny.SeqOf(_248_i1)).Cardinality()))
		_242_v0 = ((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_249_v4).Contains((_dafny.Zero).Minus(_dafny.IntOfInt64(-586))) {
				return (_249_v4).Get((_dafny.Zero).Minus(_dafny.IntOfInt64(-586))).(_dafny.Int)
			}
			return _247_v3
		})(), _248_i1))).Cmp((_247_v3).Times(((_252_v7).Update(_251_v6, _dafny.IntOfInt64(464))).Cardinality())) < 0
		var _253_v8 _dafny.Sequence
		_ = _253_v8
		_253_v8 = _dafny.UnicodeSeqOfUtf8Bytes("rcrmu")
		if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("dbqp"), _253_v8) {
			var _254_v9 *C3
			_ = _254_v9
			var _nw23 *C3 = New_C3_()
			_ = _nw23
			_nw23.Ctor__(_244_v2)
			_254_v9 = _nw23
			var _255_v10 _dafny.Sequence
			_ = _255_v10
			_255_v10 = _dafny.SeqOf(!(_242_v0), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("vbki"), _253_v8), _242_v0)
			var _256_v11 _dafny.Set
			_ = _256_v11
			_256_v11 = _dafny.SetOf(_247_v3, _247_v3)
			var _257_v12 _dafny.MultiSet
			_ = _257_v12
			_257_v12 = _dafny.MultiSetOf(_242_v0, _242_v0, _242_v0, _242_v0, false)
			var _258_v13 _dafny.Map
			_ = _258_v13
			_258_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, Companion_Default___.Fm4(_248_i1, _256_v11, _246_globalState))
			var _259_v14 _dafny.Array
			_ = _259_v14
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
			_ = _nw24
			_259_v14 = _nw24
			var _260_v15 _dafny.Sequence
			_ = _260_v15
			_260_v15 = _dafny.SeqOf(_244_v2, _244_v2)
			var _261_v16 *C2
			_ = _261_v16
			var _nw25 *C2 = New_C2_()
			_ = _nw25
			_nw25.Ctor__(!(!((Companion_Default___.Fm29(_246_globalState)).IsProperSubsetOf(_256_v11))), (func() _dafny.Int {
				if (_257_v12).Contains(!(_242_v0)) {
					return (_257_v12).Multiplicity(!(_242_v0))
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D1_.Create_DC4_(_258_v13, _242_v0, (_256_v11).Cardinality(), true)).Dtor_cf11(), _247_v3)).Cardinality()
			})(), _259_v14, (_256_v11).IsSubsetOf(_256_v11), (_260_v15).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_260_v15).Cardinality()))).Uint32()).(_dafny.Array))
			_261_v16 = _nw25
			var _262_v17 _dafny.CodePoint
			_ = _262_v17
			_262_v17 = _dafny.CodePoint('c')
			var _rhs15 _dafny.Int = (_dafny.Zero).Minus(_247_v3)
			_ = _rhs15
			var _rhs16 _dafny.Sequence = _dafny.SeqOf(_dafny.Companion_Sequence_.Equal(_250_v5, _250_v5), !(_261_v16.F7), (Companion_Default___.Fm38(_dafny.IntOfInt64(301), _242_v0, _253_v8, _262_v17, _246_globalState)).Dtor_cf23())
			_ = _rhs16
			var _rhs17 _dafny.Int = (_250_v5).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_250_v5).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs17
			var _rhs18 *C2 = _261_v16
			_ = _rhs18
			_247_v3 = _rhs15
			_255_v10 = _rhs16
			_247_v3 = _rhs17
			_261_v16 = _rhs18
			var _263_v22 _dafny.Array
			_ = _263_v22
			var _nwElement0_5 _dafny.Map = _258_v13
			_ = _nwElement0_5
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(13))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_5, 0)
			(_nw26).ArraySet1(_258_v13, 1)
			(_nw26).ArraySet1(_258_v13, 2)
			(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_i1, _242_v0), 3)
			(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_253_v8).Cardinality()), (_255_v10).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_255_v10).Cardinality()))).Uint32()).(bool)), 4)
			(_nw26).ArraySet1(_258_v13, 5)
			(_nw26).ArraySet1(func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(54), _dafny.IntOfInt64(-989))); ; {
					_compr_23, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _264_v18 _dafny.Int
					_264_v18 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(54)).Cmp(_264_v18) <= 0) && ((_264_v18).Cmp(_dafny.IntOfInt64(-989)) < 0) {
						_coll23.Add((_264_v18).Plus(_247_v3), _261_v16.F7)
					}
				}
				return _coll23.ToMap()
			}(), 6)
			(_nw26).ArraySet1(_258_v13, 7)
			(_nw26).ArraySet1(func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-130), _dafny.IntOfInt64(973))); ; {
					_compr_24, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _265_v19 _dafny.Int
					_265_v19 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(-130)).Cmp(_265_v19) <= 0) && ((_265_v19).Cmp(_dafny.IntOfInt64(973)) < 0) {
						_coll24.Add(Companion_Default___.SafeDivisionInt(_265_v19, _248_i1), false)
					}
				}
				return _coll24.ToMap()
			}(), 8)
			(_nw26).ArraySet1((_258_v13).Merge(_258_v13), 9)
			(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_i1, _261_v16.F7), 10)
			(_nw26).ArraySet1((_258_v13).Merge(func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-922), _dafny.IntOfInt64(-131))); ; {
					_compr_25, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _266_v20 _dafny.Int
					_266_v20 = interface{}(_compr_25).(_dafny.Int)
					if ((_dafny.IntOfInt64(-922)).Cmp(_266_v20) <= 0) && ((_266_v20).Cmp(_dafny.IntOfInt64(-131)) < 0) {
						_coll25.Add((_266_v20).Minus(_dafny.IntOfInt64(28)), _261_v16.F7)
					}
				}
				return _coll25.ToMap()
			}()), 11)
			(_nw26).ArraySet1(func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(558), _dafny.IntOfInt64(765))); ; {
					_compr_26, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _267_v21 _dafny.Int
					_267_v21 = interface{}(_compr_26).(_dafny.Int)
					if ((_dafny.IntOfInt64(558)).Cmp(_267_v21) <= 0) && ((_267_v21).Cmp(_dafny.IntOfInt64(765)) < 0) {
						_coll26.Add((_267_v21).Plus(_248_i1), (_255_v10).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_250_v5)).Cardinality(), _dafny.IntOfUint32((_255_v10).Cardinality()))).Uint32()).(bool))
					}
				}
				return _coll26.ToMap()
			}(), 12)
			_263_v22 = _nw26
			var _268_v23 D1
			_ = _268_v23
			_268_v23 = Companion_D1_.Create_DC4_(_258_v13, _242_v0, _248_i1, _242_v0)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_263_v22), 0))
			_ = _index20
			(_263_v22).ArraySet1((_268_v23).Dtor_cf9(), (_index20).Int())
			var _269_v24 _dafny.Array
			_ = _269_v24
			var _nwElement0_6 _dafny.Sequence = _253_v8
			_ = _nwElement0_6
			var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(9))
			_ = _nw27
			(_nw27).ArraySet1(_nwElement0_6, 0)
			(_nw27).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("oqt"), 1)
			(_nw27).ArraySet1(_253_v8, 2)
			(_nw27).ArraySet1(_253_v8, 3)
			(_nw27).ArraySet1(_253_v8, 4)
			(_nw27).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cpvuli"), 5)
			(_nw27).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_253_v8, _dafny.UnicodeSeqOfUtf8Bytes("v")), 6)
			(_nw27).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(521))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_270_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_271_i2 _dafny.Int) _dafny.CodePoint {
					return _270_v17
				}
			})(_262_v17))), 7)
			(_nw27).ArraySet1(_253_v8, 8)
			_269_v24 = _nw27
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_269_v24), 0))
			_ = _index21
			(_269_v24).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fipalxnh"), (_index21).Int())
			var _272_v26 D2
			_ = _272_v26
			_272_v26 = Companion_D2_.Create_DC5_(_dafny.SeqOf(_262_v17, _262_v17, _262_v17))
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_263_v22), 0))
			_ = _index22
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_269_v24), 0))
			_ = _index23
			var _rhs19 _dafny.Map = func() _dafny.Map {
				var _coll27 = _dafny.NewMapBuilder()
				_ = _coll27
				for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(65), _dafny.IntOfInt64(555))); ; {
					_compr_27, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _273_v25 _dafny.Int
					_273_v25 = interface{}(_compr_27).(_dafny.Int)
					if ((_dafny.IntOfInt64(65)).Cmp(_273_v25) <= 0) && ((_273_v25).Cmp(_dafny.IntOfInt64(555)) < 0) {
						_coll27.Add((_273_v25).Plus((_261_v16).F8()), (func() bool {
							if (_258_v13).Contains(_247_v3) {
								return (_258_v13).Get(_247_v3).(bool)
							}
							return _261_v16.F7
						})())
					}
				}
				return _coll27.ToMap()
			}()
			_ = _rhs19
			var _rhs20 *C3 = _254_v9
			_ = _rhs20
			var _rhs21 _dafny.Sequence = (_272_v26).Dtor_cf13()
			_ = _rhs21
			var _rhs22 _dafny.Int = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(616), _dafny.IntOfUint32((_250_v5).Cardinality()))).Plus((_261_v16).F8())
			_ = _rhs22
			var _lhs10 _dafny.Array = _263_v22
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_263_v22), 0))
			_ = _lhs11
			var _lhs12 _dafny.Array = _269_v24
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_269_v24), 0))
			_ = _lhs13
			(_lhs10).ArraySet1(_rhs19, (_lhs11).Int())
			_254_v9 = _rhs20
			(_lhs12).ArraySet1(_rhs21, (_lhs13).Int())
			_247_v3 = _rhs22
			var _274_v27 _dafny.Array
			_ = _274_v27
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
			_ = _nw28
			_274_v27 = _nw28
			var _275_v28 _dafny.MultiSet
			_ = _275_v28
			var _276_v29 bool
			_ = _276_v29
			var _out0 _dafny.MultiSet
			_ = _out0
			var _out1 bool
			_ = _out1
			_out0, _out1 = (_261_v16).M1(_247_v3, _262_v17, _269_v24, _274_v27, _246_globalState)
			_275_v28 = _out0
			_276_v29 = _out1
			_276_v29 = !(!(!(_276_v29) || (Companion_Default___.Fm20(_242_v0, _246_globalState))))
		} else {
			_247_v3 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-2), _247_v3)
			var _277_v30 bool
			_ = _277_v30
			var _278_v31 _dafny.Map
			_ = _278_v31
			var _out2 bool
			_ = _out2
			var _out3 _dafny.Map
			_ = _out3
			_out2, _out3 = Companion_Default___.M6(_246_globalState)
			_277_v30 = _out2
			_278_v31 = _out3
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_244_v2), 0))
			_ = _index24
			(_244_v2).ArraySet1(_277_v30, (_index24).Int())
			var _279_v32 _dafny.Array
			_ = _279_v32
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_6
			var _nw29 _dafny.Array
			_ = _nw29
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw29 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Map = (func(_280_v30 bool, _281_i1 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_282_i3 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v30, _281_i1)
					}
				})(_277_v30, _248_i1)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw29 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw29).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw29).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_279_v32 = _nw29
			var _283_v33 *C5
			_ = _283_v33
			var _nw30 *C5 = New_C5_()
			_ = _nw30
			_nw30.Ctor__(_277_v30, _244_v2, _279_v32, _242_v0)
			_283_v33 = _nw30
			var _284_v34 _dafny.Set
			_ = _284_v34
			_284_v34 = _dafny.SetOf(_283_v33, _283_v33)
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(454), _dafny.ArrayLen((_244_v2), 0))
			_ = _index25
			(_244_v2).ArraySet1(!(((_284_v34).Union(_284_v34)).IsSubsetOf((_284_v34).Difference(_284_v34))), (_index25).Int())
			var _285_v35 _dafny.Sequence
			_ = _285_v35
			_285_v35 = _dafny.SeqOf((_283_v33).F5())
			_247_v3 = (_247_v3).Minus(_dafny.IntOfUint32((_285_v35).Cardinality()))
			var _286_v36 _dafny.MultiSet
			_ = _286_v36
			_286_v36 = _dafny.MultiSetOf(_251_v6)
			var _287_v37 _dafny.Array
			_ = _287_v37
			var _nw31 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
			_ = _nw31
			_287_v37 = _nw31
			var _288_v38 _dafny.Map
			_ = _288_v38
			_288_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_287_v37, _247_v3)
			_247_v3 = ((func() _dafny.Map {
				if (_286_v36).IsDisjointFrom(_286_v36) {
					return (_288_v38).Update(_287_v37, _248_i1)
				}
				return (func() _dafny.Map {
					if _242_v0 {
						return _288_v38
					}
					return _288_v38
				})()
			})()).Cardinality()
		}
		var _289_v39 _dafny.Map
		_ = _289_v39
		_289_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_244_v2, _249_v4)
		var _290_v40 _dafny.Array
		_ = _290_v40
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_7
		var _nw32 _dafny.Array
		_ = _nw32
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw32 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Map = (func(_291_v0 bool, _292_v3 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_293_i4 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v0, _292_v3)
				}
			})(_242_v0, _247_v3)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw32 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw32).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw32).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_290_v40 = _nw32
		var _294_v41 *C1
		_ = _294_v41
		var _nw33 *C1 = New_C1_()
		_ = _nw33
		_nw33.Ctor__(_290_v40, _242_v0, _244_v2)
		_294_v41 = _nw33
		var _295_v42 D10
		_ = _295_v42
		_295_v42 = Companion_D10_.Create_DC31_(_294_v41)
		var _rhs23 _dafny.Map = _289_v39
		_ = _rhs23
		var _rhs24 bool = (_dafny.IntOfInt64(669)).Cmp(_248_i1) <= 0
		_ = _rhs24
		var _rhs25 *C1 = (_295_v42).Dtor_cf57()
		_ = _rhs25
		_289_v39 = _rhs23
		_242_v0 = _rhs24
		_294_v41 = _rhs25
	}
	var _296_v43 bool
	_ = _296_v43
	var _297_v44 _dafny.Map
	_ = _297_v44
	var _out4 bool
	_ = _out4
	var _out5 _dafny.Map
	_ = _out5
	_out4, _out5 = Companion_Default___.M6(_246_globalState)
	_296_v43 = _out4
	_297_v44 = _out5
	var _298_v45 _dafny.Array
	_ = _298_v45
	var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
	_ = _nw34
	_298_v45 = _nw34
	var _299_v46 D5
	_ = _299_v46
	_299_v46 = Companion_D5_.Create_DC14_(_298_v45)
	var _source8 D5 = _299_v46
	_ = _source8
	if _source8.Is_DC15() {
		var _300___mcc_h0 D4 = _source8.Get_().(D5_DC15).Cf33
		_ = _300___mcc_h0
		var _301___mcc_h1 bool = _source8.Get_().(D5_DC15).Cf34
		_ = _301___mcc_h1
		var _302_cf34 bool = _301___mcc_h1
		_ = _302_cf34
		var _303_cf33 D4 = _300___mcc_h0
		_ = _303_cf33
		_243_v1 = _dafny.SetOf(_302_cf34, _302_cf34, _242_v0)
		_247_v3 = Companion_Default___.Fm0(_247_v3, _246_globalState)
		var _304_v47 _dafny.Array
		_ = _304_v47
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
		_ = _nw35
		_304_v47 = _nw35
		var _305_v48 _dafny.Sequence
		_ = _305_v48
		_305_v48 = _dafny.UnicodeSeqOfUtf8Bytes("u")
		var _306_v49 _dafny.Sequence
		_ = _306_v49
		_306_v49 = _dafny.SeqOf(_305_v48)
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_304_v47), 0))
		_ = _index26
		(_304_v47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_306_v49, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}((func(_307_v48 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_308_i5 _dafny.Int) _dafny.Sequence {
				return _307_v48
			}
		})(_305_v48)))), (_index26).Int())
		var _309_v50 _dafny.CodePoint
		_ = _309_v50
		_309_v50 = _dafny.CodePoint('e')
		var _310_v51 *C3
		_ = _310_v51
		var _nw36 *C3 = New_C3_()
		_ = _nw36
		_nw36.Ctor__(_244_v2)
		_310_v51 = _nw36
		var _311_v52 _dafny.Map
		_ = _311_v52
		_311_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_v48, _310_v51)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_304_v47), 0))
		_ = _index27
		var _rhs26 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_312_v49 _dafny.Sequence, _313_v3 _dafny.Int, _314_v48 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_315_i6 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_312_v49).Select((Companion_Default___.SafeIndex(_313_v3, _dafny.IntOfUint32((_312_v49).Cardinality()))).Uint32()).(_dafny.Sequence), _314_v48), (Companion_Default___.SafeIndex(_313_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_312_v49).Select((Companion_Default___.SafeIndex(_313_v3, _dafny.IntOfUint32((_312_v49).Cardinality()))).Uint32()).(_dafny.Sequence), _314_v48)).Cardinality()))).Uint32(), _dafny.CodePoint('q'))
			}
		})(_306_v49, _247_v3, _305_v48)))
		_ = _rhs26
		var _rhs27 _dafny.CodePoint = _309_v50
		_ = _rhs27
		var _rhs28 bool = _296_v43
		_ = _rhs28
		var _rhs29 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_v48, _310_v51)
		_ = _rhs29
		var _lhs14 _dafny.Array = _304_v47
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_304_v47), 0))
		_ = _lhs15
		(_lhs14).ArraySet1(_rhs26, (_lhs15).Int())
		_309_v50 = _rhs27
		_296_v43 = _rhs28
		_311_v52 = _rhs29
		_247_v3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_304_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_304_v47), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm21(_246_globalState))).Cardinality())
	} else if _source8.Is_DC16() {
		var _316_v53 _dafny.MultiSet
		_ = _316_v53
		_316_v53 = _dafny.MultiSetOf(_dafny.IntOfInt64(729), _247_v3, _247_v3)
		var _317_v54 _dafny.Sequence
		_ = _317_v54
		_317_v54 = _dafny.SeqOf(true, _296_v43, _296_v43, _242_v0)
		var _318_v55 D4
		_ = _318_v55
		_318_v55 = Companion_D4_.Create_DC11_(_317_v54)
		var _319_v56 _dafny.Set
		_ = _319_v56
		_319_v56 = _dafny.SetOf((_dafny.MultiSetOf(_247_v3)).Cardinality(), (_dafny.Zero).Minus(_247_v3))
		var _320_v57 _dafny.Sequence
		_ = _320_v57
		_320_v57 = _dafny.UnicodeSeqOfUtf8Bytes("yfwppgog")
		var _321_v58 _dafny.Map
		_ = _321_v58
		_321_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_dafny.IntOfInt64(243), _319_v56, _246_globalState), _dafny.IntOfUint32((_320_v57).Cardinality()))
		var _322_v59 _dafny.Map
		_ = _322_v59
		_322_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-66), _242_v0)
		var _323_v60 _dafny.Sequence
		_ = _323_v60
		_323_v60 = _dafny.SeqOf(_244_v2)
		var _324_v61 D1
		_ = _324_v61
		_324_v61 = Companion_D1_.Create_DC4_(_322_v59, _296_v43, _dafny.IntOfUint32((_323_v60).Cardinality()), _296_v43)
		var _325_v62 _dafny.Array
		_ = _325_v62
		var _nwElement0_7 _dafny.Int = (func() _dafny.Int {
			if (_316_v53).Contains(_dafny.IntOfInt64(-547)) {
				return (_316_v53).Multiplicity(_dafny.IntOfInt64(-547))
			}
			return _247_v3
		})()
		_ = _nwElement0_7
		var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(22))
		_ = _nw37
		(_nw37).ArraySet1(_nwElement0_7, 0)
		(_nw37).ArraySet1(Companion_Default___.SafeDivisionInt(_247_v3, _dafny.IntOfInt64(165)), 1)
		(_nw37).ArraySet1((func() _dafny.Int {
			if (_316_v53).Contains(_247_v3) {
				return (_316_v53).Multiplicity(_247_v3)
			}
			return _247_v3
		})(), 2)
		(_nw37).ArraySet1(_247_v3, 3)
		(_nw37).ArraySet1(_247_v3, 4)
		(_nw37).ArraySet1(_247_v3, 5)
		(_nw37).ArraySet1(Companion_Default___.SafeModuloInt(_247_v3, _247_v3), 6)
		(_nw37).ArraySet1(_247_v3, 7)
		(_nw37).ArraySet1(_247_v3, 8)
		(_nw37).ArraySet1(_247_v3, 9)
		(_nw37).ArraySet1(_dafny.IntOfInt64(909), 10)
		(_nw37).ArraySet1(_247_v3, 11)
		(_nw37).ArraySet1(_dafny.IntOfUint32(((_318_v55).Dtor_cf28()).Cardinality()), 12)
		(_nw37).ArraySet1((Companion_Default___.Fm0(_247_v3, _246_globalState)).Times(_247_v3), 13)
		(_nw37).ArraySet1(_247_v3, 14)
		(_nw37).ArraySet1(_247_v3, 15)
		(_nw37).ArraySet1(_247_v3, 16)
		(_nw37).ArraySet1((_dafny.SetOf(_247_v3)).Cardinality(), 17)
		(_nw37).ArraySet1(_247_v3, 18)
		(_nw37).ArraySet1(((_321_v58).Merge(_321_v58)).Cardinality(), 19)
		(_nw37).ArraySet1(_247_v3, 20)
		(_nw37).ArraySet1((_324_v61).Dtor_cf11(), 21)
		_325_v62 = _nw37
		_325_v62 = _325_v62
		var _326_v63 *C3
		_ = _326_v63
		var _nw38 *C3 = New_C3_()
		_ = _nw38
		_nw38.Ctor__(_244_v2)
		_326_v63 = _nw38
		var _327_v64 _dafny.Map
		_ = _327_v64
		_327_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_326_v63, _dafny.IntOfInt64(885))
		_327_v64 = (_327_v64).Update(_326_v63, _247_v3)
		if _242_v0 {
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_244_v2), 0))
			_ = _index28
			(_244_v2).ArraySet1(true, (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_244_v2), 0))
			_ = _index29
			(_244_v2).ArraySet1(_296_v43, (_index29).Int())
			var _328_v65 _dafny.Array
			_ = _328_v65
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_8
			var _nw39 _dafny.Array
			_ = _nw39
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw39 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = func(_329_i7 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("uyjkwf")
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw39 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw39).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw39).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_328_v65 = _nw39
			_328_v65 = _328_v65
			var _330_v66 _dafny.Array
			_ = _330_v66
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_9
			var _nw40 _dafny.Array
			_ = _nw40
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw40 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) bool = (func(_331_v0 bool) func(_dafny.Int) bool {
					return func(_332_i8 _dafny.Int) bool {
						return _331_v0
					}
				})(_242_v0)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw40 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw40).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw40).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_330_v66 = _nw40
			var _333_v67 _dafny.Array
			_ = _333_v67
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
			_ = _nw41
			_333_v67 = _nw41
			var _334_v68 *C5
			_ = _334_v68
			var _nw42 *C5 = New_C5_()
			_ = _nw42
			_nw42.Ctor__((_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool), _330_v66, _333_v67, (_243_v1).IsProperSubsetOf(_243_v1))
			_334_v68 = _nw42
			var _335_v69 D8
			_ = _335_v69
			_335_v69 = Companion_D8_.Create_DC26_(false, (_334_v68).F5(), _296_v43, _247_v3)
			var _336_v70 D8
			_ = _336_v70
			_336_v70 = Companion_D8_.Create_DC29_(_335_v69)
			_336_v70 = _336_v70
			var _337_v71 _dafny.CodePoint
			_ = _337_v71
			_337_v71 = _dafny.CodePoint('x')
			var _338_v72 _dafny.Map
			_ = _338_v72
			_338_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v71, _dafny.CodePoint('h'))
			_296_v43 = _dafny.Companion_Sequence_.Contains(_320_v57, (func() _dafny.CodePoint {
				if (_338_v72).Contains(_337_v71) {
					return (_338_v72).Get(_337_v71).(_dafny.CodePoint)
				}
				return _337_v71
			})())
		} else {
			_247_v3 = (_247_v3).Minus(_dafny.IntOfUint32((_320_v57).Cardinality()))
			var _339_v73 _dafny.CodePoint
			_ = _339_v73
			_339_v73 = _dafny.CodePoint('y')
			_320_v57 = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("rasbc"), (Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rasbc")).Cardinality()))).Uint32(), _339_v73)
			_296_v43 = _296_v43
			var _340_v74 *C4
			_ = _340_v74
			var _nw43 *C4 = New_C4_()
			_ = _nw43
			_nw43.Ctor__((_243_v1).IsProperSubsetOf(_dafny.SetOf(_296_v43)))
			_340_v74 = _nw43
			var _341_v75 D6
			_ = _341_v75
			_341_v75 = Companion_D6_.Create_DC20_(Companion_Default___.Fm20(_296_v43, _246_globalState), _242_v0)
			var _342_v76 bool
			_ = _342_v76
			var _343_v77 _dafny.Int
			_ = _343_v77
			var _344_v78 bool
			_ = _344_v78
			var _out6 bool
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			var _out8 bool
			_ = _out8
			_out6, _out7, _out8 = (_340_v74).M3(_325_v62, (_341_v75).Dtor_cf38(), _246_globalState)
			_342_v76 = _out6
			_343_v77 = _out7
			_344_v78 = _out8
		}
		_247_v3 = (_326_v63).Fm2(_246_globalState)
	} else if _source8.Is_DC17() {
		var _345___mcc_h2 _dafny.Int = _source8.Get_().(D5_DC17).Cf35
		_ = _345___mcc_h2
		var _346_cf35 _dafny.Int = _345___mcc_h2
		_ = _346_cf35
		var _347_v79 _dafny.CodePoint
		_ = _347_v79
		_347_v79 = _dafny.CodePoint('d')
		var _348_v80 _dafny.Sequence
		_ = _348_v80
		_348_v80 = _dafny.SeqOf(_347_v79, _347_v79, _347_v79, _347_v79)
		var _source9 D2 = Companion_D2_.Create_DC5_(_348_v80)
		_ = _source9
		if _source9.Is_DC6() {
			var _349___mcc_h5 _dafny.CodePoint = _source9.Get_().(D2_DC6).Cf14
			_ = _349___mcc_h5
			var _350___mcc_h6 _dafny.Int = _source9.Get_().(D2_DC6).Cf15
			_ = _350___mcc_h6
			var _351_cf15 _dafny.Int = _350___mcc_h6
			_ = _351_cf15
			var _352_cf14 _dafny.CodePoint = _349___mcc_h5
			_ = _352_cf14
			var _353_v81 _dafny.Array
			_ = _353_v81
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw44
			_353_v81 = _nw44
			_353_v81 = (func() _dafny.Array {
				if _296_v43 {
					return _353_v81
				}
				return _353_v81
			})()
			var _354_v82 _dafny.Array
			_ = _354_v82
			var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw45
			_354_v82 = _nw45
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_354_v82), 0))
			_ = _index30
			(_354_v82).ArraySet1(_247_v3, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_354_v82), 0))
			_ = _index31
			(_354_v82).ArraySet1(_247_v3, (_index31).Int())
			var _355_v83 bool
			_ = _355_v83
			var _356_v84 _dafny.Map
			_ = _356_v84
			var _out9 bool
			_ = _out9
			var _out10 _dafny.Map
			_ = _out10
			_out9, _out10 = Companion_Default___.M6(_246_globalState)
			_355_v83 = _out9
			_356_v84 = _out10
			var _357_v85 _dafny.Array
			_ = _357_v85
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
			_ = _nw46
			_357_v85 = _nw46
			var _358_v86 _dafny.Array
			_ = _358_v86
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(27))
			_ = _nw47
			_358_v86 = _nw47
			var _359_v87 _dafny.Map
			_ = _359_v87
			_359_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_358_v86, _346_cf35)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_357_v85), 0))
			_ = _index32
			(_357_v85).ArraySet1(_359_v87, (_index32).Int())
			var _360_v88 _dafny.Map
			_ = _360_v88
			_360_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_354_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_354_v82), 0))).Int()).(_dafny.Int), _242_v0)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_357_v85), 0))
			_ = _index33
			var _rhs30 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_358_v86, _346_cf35)
			_ = _rhs30
			var _rhs31 _dafny.Int = Companion_Default___.SafeDivisionInt(_247_v3, ((_360_v88).Merge(_360_v88)).Cardinality())
			_ = _rhs31
			var _rhs32 bool = _242_v0
			_ = _rhs32
			var _rhs33 _dafny.Int = _dafny.IntOfInt64(982)
			_ = _rhs33
			var _lhs16 _dafny.Array = _357_v85
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_357_v85), 0))
			_ = _lhs17
			(_lhs16).ArraySet1(_rhs30, (_lhs17).Int())
			_247_v3 = _rhs31
			_242_v0 = _rhs32
			_247_v3 = _rhs33
		} else if _source9.Is_DC7() {
			var _361___mcc_h7 _dafny.Int = _source9.Get_().(D2_DC7).Cf16
			_ = _361___mcc_h7
			var _362___mcc_h8 _dafny.Int = _source9.Get_().(D2_DC7).Cf17
			_ = _362___mcc_h8
			var _363___mcc_h9 _dafny.MultiSet = _source9.Get_().(D2_DC7).Cf18
			_ = _363___mcc_h9
			var _364___mcc_h10 _dafny.Int = _source9.Get_().(D2_DC7).Cf19
			_ = _364___mcc_h10
			var _365___mcc_h11 _dafny.Int = _source9.Get_().(D2_DC7).Cf20
			_ = _365___mcc_h11
			var _366_cf20 _dafny.Int = _365___mcc_h11
			_ = _366_cf20
			var _367_cf19 _dafny.Int = _364___mcc_h10
			_ = _367_cf19
			var _368_cf18 _dafny.MultiSet = _363___mcc_h9
			_ = _368_cf18
			var _369_cf17 _dafny.Int = _362___mcc_h8
			_ = _369_cf17
			var _370_cf16 _dafny.Int = _361___mcc_h7
			_ = _370_cf16
			var _371_v89 _dafny.Map
			_ = _371_v89
			_371_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_cf19, _242_v0)
			_366_cf20 = ((_dafny.IntOfUint32((_348_v80).Cardinality())).Plus(_346_cf35)).Plus((_346_cf35).Minus((_371_v89).Cardinality()))
			var _372_v90 _dafny.Sequence
			_ = _372_v90
			_372_v90 = _dafny.SeqOf(_366_cf20)
			var _373_v91 D8
			_ = _373_v91
			_373_v91 = Companion_D8_.Create_DC27_(_242_v0, _242_v0, (_dafny.SetOf(_372_v90)).Cardinality(), _247_v3)
			var _374_v92 _dafny.Array
			_ = _374_v92
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_10
			var _nw48 _dafny.Array
			_ = _nw48
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw48 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Map = (func(_375_cf17 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_376_i9 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _375_cf17)
					}
				})(_369_cf17)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw48 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw48).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw48).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_374_v92 = _nw48
			var _377_v93 *C2
			_ = _377_v93
			var _nw49 *C2 = New_C2_()
			_ = _nw49
			_nw49.Ctor__((_373_v91).Dtor_cf50(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(272))).Cardinality()), _374_v92, _242_v0, _244_v2)
			_377_v93 = _nw49
			var _378_v94 _dafny.Sequence
			_ = _378_v94
			_378_v94 = _dafny.SeqOf(_377_v93)
			var _379_v96 _dafny.Map
			_ = _379_v96
			_379_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _367_cf19)
			var _380_v97 D12
			_ = _380_v97
			_380_v97 = Companion_D12_.Create_DC38_(func() _dafny.Set {
				var _coll28 = _dafny.NewBuilder()
				_ = _coll28
				for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(180), _dafny.IntOfInt64(360))); ; {
					_compr_28, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _381_v95 _dafny.Int
					_381_v95 = interface{}(_compr_28).(_dafny.Int)
					if ((_dafny.IntOfInt64(180)).Cmp(_381_v95) <= 0) && ((_381_v95).Cmp(_dafny.IntOfInt64(360)) < 0) {
						_coll28.Add((_381_v95).Minus(_367_cf19))
					}
				}
				return _coll28.ToSet()
			}(), _366_cf20, _377_v93, (_379_v96).Cardinality())
			var _382_v98 _dafny.Array
			_ = _382_v98
			var _nwElement0_8 _dafny.Sequence = _378_v94
			_ = _nwElement0_8
			var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(17))
			_ = _nw50
			(_nw50).ArraySet1(_nwElement0_8, 0)
			(_nw50).ArraySet1(_dafny.Companion_Sequence_.Update(_378_v94, (Companion_Default___.SafeIndex((_377_v93).F8(), _dafny.IntOfUint32((_378_v94).Cardinality()))).Uint32(), (_380_v97).Dtor_cf72()), 1)
			(_nw50).ArraySet1(_378_v94, 2)
			(_nw50).ArraySet1(_378_v94, 3)
			(_nw50).ArraySet1(_378_v94, 4)
			(_nw50).ArraySet1(_378_v94, 5)
			(_nw50).ArraySet1(_dafny.SeqOf(_377_v93), 6)
			(_nw50).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_378_v94, _378_v94), 7)
			(_nw50).ArraySet1(_378_v94, 8)
			(_nw50).ArraySet1(_dafny.Companion_Sequence_.Update(_378_v94, (Companion_Default___.SafeIndex(_369_cf17, _dafny.IntOfUint32((_378_v94).Cardinality()))).Uint32(), _377_v93), 9)
			(_nw50).ArraySet1(_378_v94, 10)
			(_nw50).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_378_v94, _378_v94), 11)
			(_nw50).ArraySet1(_378_v94, 12)
			(_nw50).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_378_v94, _378_v94), 13)
			(_nw50).ArraySet1(_378_v94, 14)
			(_nw50).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_378_v94, _dafny.SeqOf(_377_v93)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_348_v80).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_378_v94, _dafny.SeqOf(_377_v93))).Cardinality()))).Uint32(), _377_v93), 15)
			(_nw50).ArraySet1(_378_v94, 16)
			_382_v98 = _nw50
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_382_v98), 0))
			_ = _index34
			(_382_v98).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_378_v94, _dafny.SeqOf(_377_v93)), (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_382_v98), 0))
			_ = _index35
			var _rhs34 _dafny.Sequence = _dafny.SeqOf(_377_v93, _377_v93, _377_v93, _377_v93, (func() *C2 {
				if _377_v93.F7 {
					return _377_v93
				}
				return _377_v93
			})())
			_ = _rhs34
			var _rhs35 _dafny.Int = _366_cf20
			_ = _rhs35
			var _lhs18 _dafny.Array = _382_v98
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_382_v98), 0))
			_ = _lhs19
			(_lhs18).ArraySet1(_rhs34, (_lhs19).Int())
			_247_v3 = _rhs35
			var _383_v99 _dafny.MultiSet
			_ = _383_v99
			_383_v99 = _dafny.MultiSetOf(_dafny.IntOfInt64(132))
			var _384_v100 _dafny.Set
			_ = _384_v100
			_384_v100 = _dafny.SetOf(_370_cf16)
			var _385_v101 _dafny.Set
			_ = _385_v101
			_385_v101 = _243_v1
			var _386_v102 _dafny.Set
			_ = _386_v102
			_386_v102 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_372_v90, (Companion_Default___.SafeIndex(_367_cf19, _dafny.IntOfUint32((_372_v90).Cardinality()))).Uint32(), _346_cf35), _372_v90, _372_v90, _372_v90)
			var _387_v103 _dafny.Array
			_ = _387_v103
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_11
			var _nw51 _dafny.Array
			_ = _nw51
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw51 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) D4 = (func(_388_v43 bool, _389_v80 _dafny.Sequence, _390_v93 *C2) func(_dafny.Int) D4 {
					return func(_391_i10 _dafny.Int) D4 {
						return Companion_D4_.Create_DC13_(_388_v43, _dafny.MultiSetFromSeq(_389_v80), _390_v93.F7)
					}
				})(_296_v43, _348_v80, _377_v93)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw51 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw51).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw51).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_387_v103 = _nw51
			var _392_v104 _dafny.Set
			_ = _392_v104
			_392_v104 = _dafny.SetOf(_387_v103)
			var _393_v105 _dafny.Array
			_ = _393_v105
			var _nwElement0_9 _dafny.Int = _367_cf19
			_ = _nwElement0_9
			var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(26))
			_ = _nw52
			(_nw52).ArraySet1(_nwElement0_9, 0)
			(_nw52).ArraySet1(_dafny.IntOfUint32((_348_v80).Cardinality()), 1)
			(_nw52).ArraySet1(_dafny.IntOfUint32((_372_v90).Cardinality()), 2)
			(_nw52).ArraySet1(_dafny.IntOfInt64(565), 3)
			(_nw52).ArraySet1((_377_v93).F8(), 4)
			(_nw52).ArraySet1(_370_cf16, 5)
			(_nw52).ArraySet1(_dafny.IntOfUint32((_372_v90).Cardinality()), 6)
			(_nw52).ArraySet1((_377_v93).F8(), 7)
			(_nw52).ArraySet1(((_379_v96).Update(_296_v43, _247_v3)).Cardinality(), 8)
			(_nw52).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_383_v99)).Cardinality()), 9)
			(_nw52).ArraySet1(_366_cf20, 10)
			(_nw52).ArraySet1(_367_cf19, 11)
			(_nw52).ArraySet1((_384_v100).Cardinality(), 12)
			(_nw52).ArraySet1(_346_cf35, 13)
			(_nw52).ArraySet1((_dafny.Zero).Minus(_370_cf16), 14)
			(_nw52).ArraySet1((_377_v93).F8(), 15)
			(_nw52).ArraySet1((_385_v101).Cardinality(), 16)
			(_nw52).ArraySet1(_dafny.IntOfUint32((_348_v80).Cardinality()), 17)
			(_nw52).ArraySet1(_370_cf16, 18)
			(_nw52).ArraySet1(_366_cf20, 19)
			(_nw52).ArraySet1((_dafny.Zero).Minus((_386_v102).Cardinality()), 20)
			(_nw52).ArraySet1((_392_v104).Cardinality(), 21)
			(_nw52).ArraySet1(_247_v3, 22)
			(_nw52).ArraySet1((_377_v93).F8(), 23)
			(_nw52).ArraySet1(_247_v3, 24)
			(_nw52).ArraySet1(_dafny.IntOfUint32((_372_v90).Cardinality()), 25)
			_393_v105 = _nw52
			var _394_v106 _dafny.MultiSet
			_ = _394_v106
			_394_v106 = _dafny.MultiSetOf(_393_v105)
			_394_v106 = _394_v106
			var _395_v107 *C5
			_ = _395_v107
			var _nw53 *C5 = New_C5_()
			_ = _nw53
			_nw53.Ctor__(_296_v43, _244_v2, _374_v92, _296_v43)
			_395_v107 = _nw53
		} else if _source9.Is_DC5() {
			var _396___mcc_h12 _dafny.Sequence = _source9.Get_().(D2_DC5).Cf13
			_ = _396___mcc_h12
			var _397_cf13 _dafny.Sequence = _396___mcc_h12
			_ = _397_cf13
			var _398_v108 bool
			_ = _398_v108
			var _399_v109 _dafny.Map
			_ = _399_v109
			var _out11 bool
			_ = _out11
			var _out12 _dafny.Map
			_ = _out12
			_out11, _out12 = Companion_Default___.M6(_246_globalState)
			_398_v108 = _out11
			_399_v109 = _out12
			_242_v0 = (func() bool {
				if (Companion_D0_.Create_DC2_(_247_v3, _398_v108, _296_v43, _247_v3)).Dtor_cf6() {
					return Companion_Default___.Fm20(true, _246_globalState)
				}
				return _242_v0
			})()
			var _400_v110 _dafny.Sequence
			_ = _400_v110
			_400_v110 = _dafny.SeqOf(_398_v108, _296_v43, !(_242_v0), _242_v0)
			var _401_v111 _dafny.Map
			_ = _401_v111
			_401_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_400_v110, (_dafny.IntOfUint32((_348_v80).Cardinality())).Cmp(_247_v3) >= 0)
			_401_v111 = (_401_v111).Update(_dafny.SeqOf(_296_v43, _242_v0), !((_346_cf35).Cmp(_346_cf35) != 0))
			var _402_v112 _dafny.Sequence
			_ = _402_v112
			_402_v112 = _dafny.SeqOf((_243_v1).Cardinality())
			var _403_v113 D8
			_ = _403_v113
			_403_v113 = Companion_D8_.Create_DC28_(_296_v43, _dafny.IntOfUint32((_397_cf13).Cardinality()))
			var _404_v114 _dafny.Map
			_ = _404_v114
			_404_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_403_v113).Dtor_cf53(), _dafny.IntOfInt64(24))
			var _405_v115 _dafny.Array
			_ = _405_v115
			var _nwElement0_10 _dafny.Map = _404_v114
			_ = _nwElement0_10
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(29))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_10, 0)
			(_nw54).ArraySet1(_404_v114, 1)
			(_nw54).ArraySet1(_404_v114, 2)
			(_nw54).ArraySet1(_404_v114, 3)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(716)), 4)
			(_nw54).ArraySet1(_404_v114, 5)
			(_nw54).ArraySet1(_404_v114, 6)
			(_nw54).ArraySet1(_404_v114, 7)
			(_nw54).ArraySet1(_404_v114, 8)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _346_cf35), 9)
			(_nw54).ArraySet1(_404_v114, 10)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_398_v108, _346_cf35), 11)
			(_nw54).ArraySet1(_404_v114, 12)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_296_v43, _dafny.IntOfUint32((_400_v110).Cardinality())), 13)
			(_nw54).ArraySet1(_404_v114, 14)
			(_nw54).ArraySet1(Companion_Default___.Fm39(Companion_Default___.Fm0(_346_cf35, _246_globalState), _398_v108, _246_globalState), 15)
			(_nw54).ArraySet1(_404_v114, 16)
			(_nw54).ArraySet1((_404_v114).Update(_296_v43, _346_cf35), 17)
			(_nw54).ArraySet1(_404_v114, 18)
			(_nw54).ArraySet1(_404_v114, 19)
			(_nw54).ArraySet1(_404_v114, 20)
			(_nw54).ArraySet1((Companion_D11_.Create_DC34_(_404_v114)).Dtor_cf62(), 21)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_296_v43), _dafny.IntOfInt64(107)), 22)
			(_nw54).ArraySet1(_404_v114, 23)
			(_nw54).ArraySet1(_404_v114, 24)
			(_nw54).ArraySet1(_404_v114, 25)
			(_nw54).ArraySet1(Companion_Default___.Fm39(_dafny.IntOfUint32((_397_cf13).Cardinality()), _242_v0, _246_globalState), 26)
			(_nw54).ArraySet1(_404_v114, 27)
			(_nw54).ArraySet1(_404_v114, 28)
			_405_v115 = _nw54
			var _406_v116 _dafny.Map
			_ = _406_v116
			_406_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_402_v112).Cardinality()), _405_v115)
			var _407_v117 *C5
			_ = _407_v117
			var _nw55 *C5 = New_C5_()
			_ = _nw55
			_nw55.Ctor__(((_dafny.Zero).Minus(_247_v3)).Cmp(_247_v3) == 0, _244_v2, (func() _dafny.Array {
				if (_406_v116).Contains(_247_v3) {
					return (_406_v116).Get(_247_v3).(_dafny.Array)
				}
				return _405_v115
			})(), !(_398_v108))
			_407_v117 = _nw55
		} else {
			var _408___mcc_h13 D2 = _source9.Get_().(D2_DC8).Cf21
			_ = _408___mcc_h13
			var _409_cf21 D2 = _408___mcc_h13
			_ = _409_cf21
			var _410_v118 _dafny.Set
			_ = _410_v118
			_410_v118 = _dafny.SetOf(_247_v3, _dafny.IntOfInt64(486))
			_296_v43 = (func() bool {
				if !(_410_v118).Equals(_410_v118) {
					return _242_v0
				}
				return _242_v0
			})()
			var _411_v119 _dafny.Sequence
			_ = _411_v119
			_411_v119 = _dafny.SeqOf(_dafny.IntOfInt64(857), _247_v3)
			var _412_v120 _dafny.MultiSet
			_ = _412_v120
			_412_v120 = _dafny.MultiSetOf(_244_v2)
			var _413_v121 D10
			_ = _413_v121
			_413_v121 = Companion_D10_.Create_DC32_(_412_v120, _346_cf35, _247_v3)
			var _414_v122 _dafny.Map
			_ = _414_v122
			_414_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_247_v3).Cmp((_411_v119).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_411_v119).Cardinality()))).Uint32()).(_dafny.Int)) >= 0, _413_v121)
			_414_v122 = (_414_v122).Update(_296_v43, _413_v121)
			var _415_v123 _dafny.Array
			_ = _415_v123
			var _nw56 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
			_ = _nw56
			_415_v123 = _nw56
			var _416_v124 _dafny.Sequence
			_ = _416_v124
			_416_v124 = _dafny.SeqOf(_415_v123)
			var _417_v125 _dafny.Map
			_ = _417_v125
			_417_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_415_v123), _416_v124), _346_cf35)
			_417_v125 = (_417_v125).Update(_416_v124, _247_v3)
			var _418_v126 _dafny.Array
			_ = _418_v126
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_12
			var _nw57 _dafny.Array
			_ = _nw57
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw57 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Map = (func(_419_v43 bool, _420_v3 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_421_i11 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v43, _420_v3)
					}
				})(_296_v43, _247_v3)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw57 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw57).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw57).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_418_v126 = _nw57
			var _422_v127 _dafny.Map
			_ = _422_v127
			_422_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_242_v0), _418_v126)
			var _423_v128 *C2
			_ = _423_v128
			var _nw58 *C2 = New_C2_()
			_ = _nw58
			_nw58.Ctor__(!(_242_v0), _247_v3, (func() _dafny.Array {
				if (_422_v127).Contains(_296_v43) {
					return (_422_v127).Get(_296_v43).(_dafny.Array)
				}
				return _418_v126
			})(), _242_v0, _244_v2)
			_423_v128 = _nw58
		}
		_346_cf35 = _346_cf35
		var _424_v129 _dafny.Array
		_ = _424_v129
		var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
		_ = _nw59
		_424_v129 = _nw59
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_424_v129), 0))
		_ = _index36
		(_424_v129).ArraySet1(_348_v80, (_index36).Int())
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_424_v129), 0))
		_ = _index37
		(_424_v129).ArraySet1(_348_v80, (_index37).Int())
		if _242_v0 {
			var _425_v131 _dafny.Array
			_ = _425_v131
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_13
			var _nw60 _dafny.Array
			_ = _nw60
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw60 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Map = (func(_426_cf35 _dafny.Int, _427_v3 _dafny.Int, _428_v79 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
					return func(_429_i12 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_426_cf35, _427_v3)).Merge(func() _dafny.Map {
							var _coll29 = _dafny.NewMapBuilder()
							_ = _coll29
							for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-584), _dafny.IntOfInt64(501))); ; {
								_compr_29, _ok29 := _iter29()
								if !_ok29 {
									break
								}
								var _430_v130 _dafny.Int
								_430_v130 = interface{}(_compr_29).(_dafny.Int)
								if ((_dafny.IntOfInt64(-584)).Cmp(_430_v130) <= 0) && ((_430_v130).Cmp(_dafny.IntOfInt64(501)) < 0) {
									_coll29.Add((_430_v130).Times(_dafny.IntOfInt64(-372)), _dafny.IntOfUint32((_dafny.SeqOf(_428_v79)).Cardinality()))
								}
							}
							return _coll29.ToMap()
						}())
					}
				})(_346_cf35, _247_v3, _347_v79)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw60 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw60).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw60).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_425_v131 = _nw60
			_425_v131 = _425_v131
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_244_v2), 0))
			_ = _index38
			(_244_v2).ArraySet1(_242_v0, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_244_v2), 0))
			_ = _index39
			(_244_v2).ArraySet1(_296_v43, (_index39).Int())
			var _431_v132 _dafny.Map
			_ = _431_v132
			_431_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(463), _247_v3)
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_425_v131), 0))
			_ = _index40
			(_425_v131).ArraySet1((func() _dafny.Map {
				if _242_v0 {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _346_cf35)
				}
				return _431_v132
			})(), (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_425_v131), 0))
			_ = _index41
			(_425_v131).ArraySet1(_431_v132, (_index41).Int())
			var _432_v133 bool
			_ = _432_v133
			var _433_v134 _dafny.Map
			_ = _433_v134
			var _out13 bool
			_ = _out13
			var _out14 _dafny.Map
			_ = _out14
			_out13, _out14 = Companion_Default___.M6(_246_globalState)
			_432_v133 = _out13
			_433_v134 = _out14
			var _434_v135 _dafny.Array
			_ = _434_v135
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_14
			var _nw61 _dafny.Array
			_ = _nw61
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw61 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = (func(_435_v1 _dafny.Set) func(_dafny.Int) bool {
					return func(_436_i13 _dafny.Int) bool {
						return (_435_v1).IsDisjointFrom(_435_v1)
					}
				})(_243_v1)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw61 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw61).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw61).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_434_v135 = _nw61
			(_246_globalState).F1 = _434_v135
		} else {
			var _437_v136 _dafny.Map
			_ = _437_v136
			_437_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _247_v3)
			var _438_v137 _dafny.Sequence
			_ = _438_v137
			_438_v137 = _dafny.SeqOf(_437_v136, _437_v136)
			_438_v137 = _dafny.Companion_Sequence_.Concatenate(_438_v137, _dafny.Companion_Sequence_.Concatenate(_438_v137, Companion_Default___.Fm40(_246_globalState)))
			var _439_v138 *C0
			_ = _439_v138
			var _nw62 *C0 = New_C0_()
			_ = _nw62
			_nw62.Ctor__()
			_439_v138 = _nw62
			var _440_v139 _dafny.Sequence
			_ = _440_v139
			_440_v139 = _dafny.SeqOf((Companion_Default___.Fm25(true, _246_globalState)).Cardinality())
			var _441_v140 _dafny.Sequence
			_ = _441_v140
			_441_v140 = _dafny.SeqOf(_247_v3, (_440_v139).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_424_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_424_v129), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_440_v139).Cardinality()))).Uint32()).(_dafny.Int))
			var _442_v141 _dafny.MultiSet
			_ = _442_v141
			_442_v141 = _dafny.MultiSetOf(_244_v2, _244_v2)
			var _443_v142 D10
			_ = _443_v142
			_443_v142 = Companion_D10_.Create_DC32_(_442_v141, _dafny.IntOfUint32((Companion_Default___.Fm11(_346_cf35, _242_v0, _346_cf35, _246_globalState)).Cardinality()), _247_v3)
			var _444_v143 _dafny.Set
			_ = _444_v143
			_444_v143 = _dafny.SetOf(_247_v3, (_443_v142).Dtor_cf60(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _dafny.IntOfInt64(-861))).Cardinality())
			var _445_v144 _dafny.Map
			_ = _445_v144
			_445_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_cf35, _244_v2)
			var _446_v145 _dafny.Array
			_ = _446_v145
			var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
			_ = _nw63
			_446_v145 = _nw63
			var _447_v146 D10
			_ = _447_v146
			_447_v146 = Companion_D10_.Create_DC33_(_296_v43)
			var _448_v147 T1
			_ = _448_v147
			var _nw64 *C5 = New_C5_()
			_ = _nw64
			_nw64.Ctor__(_296_v43, (func() _dafny.Array {
				if (_445_v144).Contains(_247_v3) {
					return (_445_v144).Get(_247_v3).(_dafny.Array)
				}
				return _244_v2
			})(), _446_v145, (_447_v146).Dtor_cf61())
			_448_v147 = _nw64
			var _449_v148 _dafny.Map
			_ = _449_v148
			_449_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_448_v147, (_dafny.Zero).Minus((_448_v147).Fm3(_242_v0, true, _346_cf35, _247_v3, _246_globalState)))
			var _450_v149 _dafny.Map
			_ = _450_v149
			_450_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_449_v148, _347_v79)
			var _451_v150 _dafny.Sequence
			_ = _451_v150
			_451_v150 = _dafny.SeqOf(!(false), _242_v0, (_448_v147).F4(), _296_v43, _296_v43)
			var _452_v151 _dafny.MultiSet
			_ = _452_v151
			_452_v151 = _dafny.MultiSetOf((_448_v147).F4())
			var _453_v152 _dafny.Sequence
			_ = _453_v152
			_453_v152 = _dafny.SeqOf((_424_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_424_v129), 0))).Int()).(_dafny.Sequence), _348_v80, _dafny.UnicodeSeqOfUtf8Bytes("rn"), _348_v80)
			var _454_v153 *C5
			_ = _454_v153
			var _nw65 *C5 = New_C5_()
			_ = _nw65
			_nw65.Ctor__(false, _244_v2, (_448_v147).F3(), _296_v43)
			_454_v153 = _nw65
			var _455_v154 D13
			_ = _455_v154
			_455_v154 = Companion_D13_.Create_DC42_(_dafny.Companion_Sequence_.Update(_451_v150, (Companion_Default___.SafeIndex(_346_cf35, _dafny.IntOfUint32((_451_v150).Cardinality()))).Uint32(), false), (_452_v151).Update(_242_v0, Companion_Default___.Abs(_247_v3)), _dafny.Companion_Sequence_.Update(_453_v152, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_454_v153, (_448_v147).F4())).Cardinality(), _dafny.IntOfUint32((_453_v152).Cardinality()))).Uint32(), _348_v80), _296_v43, _dafny.SeqOf(_346_cf35))
			var _rhs36 bool = Companion_Default___.Fm4(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(551), _346_cf35), _444_v143, _246_globalState)
			_ = _rhs36
			var _rhs37 bool = ((_dafny.Zero).Minus(_346_cf35)).Cmp((_450_v149).Cardinality()) == 0
			_ = _rhs37
			var _rhs38 _dafny.Sequence = (_455_v154).Dtor_cf81()
			_ = _rhs38
			var _rhs39 _dafny.CodePoint = _347_v79
			_ = _rhs39
			var _rhs40 *C0 = _439_v138
			_ = _rhs40
			_242_v0 = _rhs36
			_242_v0 = _rhs37
			_441_v140 = _rhs38
			_347_v79 = _rhs39
			_439_v138 = _rhs40
			var _456_v155 _dafny.MultiSet
			_ = _456_v155
			_456_v155 = _dafny.MultiSetOf(_346_cf35)
			_346_cf35 = (_440_v139).Select((Companion_Default___.SafeIndex(((_448_v147).Fm3(!((_448_v147).F4()), false, _247_v3, (_456_v155).Cardinality(), _246_globalState)).Plus((_454_v153).Fm3((_448_v147).F4(), (_454_v153).F5(), _346_cf35, _346_cf35, _246_globalState)), _dafny.IntOfUint32((_440_v139).Cardinality()))).Uint32()).(_dafny.Int)
			var _457_v156 _dafny.Sequence
			_ = _457_v156
			_457_v156 = _dafny.SeqOf(_439_v138, _439_v138)
			var _458_v157 _dafny.Sequence
			_ = _458_v157
			_458_v157 = _dafny.SeqOf(_457_v156)
			var _459_v158 _dafny.Map
			_ = _459_v158
			_459_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_454_v153).F5(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(326))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_460_v79 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_461_i14 _dafny.Int) _dafny.CodePoint {
					return _460_v79
				}
			})(_347_v79))))
			var _462_v159 D18
			_ = _462_v159
			_462_v159 = Companion_D18_.Create_DC53_(_459_v158)
			var _463_v160 _dafny.Sequence
			_ = _463_v160
			_463_v160 = _dafny.SeqOf((_462_v159).Dtor_cf99(), _459_v158)
			var _464_v161 _dafny.Array
			_ = _464_v161
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_15
			var _nw66 _dafny.Array
			_ = _nw66
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw66 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Int = (func(_465_v129 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_466_i15 _dafny.Int) _dafny.Int {
						return (_466_i15).Minus(_dafny.IntOfUint32(((_465_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_465_v129), 0))).Int()).(_dafny.Sequence)).Cardinality()))
					}
				})(_424_v129)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw66 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw66).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw66).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_464_v161 = _nw66
			var _467_v162 _dafny.Set
			_ = _467_v162
			_467_v162 = _dafny.SetOf(_464_v161)
			var _468_v163 _dafny.Array
			_ = _468_v163
			var _nwElement0_11 _dafny.Int = (_dafny.MultiSetFromSeq(_458_v157)).Cardinality()
			_ = _nwElement0_11
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(16))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_11, 0)
			(_nw67).ArraySet1(_dafny.IntOfInt64(371), 1)
			(_nw67).ArraySet1(_346_cf35, 2)
			(_nw67).ArraySet1(_346_cf35, 3)
			(_nw67).ArraySet1(((_463_v160).Select((Companion_Default___.SafeIndex(_346_cf35, _dafny.IntOfUint32((_463_v160).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 4)
			(_nw67).ArraySet1((_247_v3).Minus(_247_v3), 5)
			(_nw67).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_247_v3), _346_cf35)), 6)
			(_nw67).ArraySet1(_247_v3, 7)
			(_nw67).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if _296_v43 {
					return (_467_v162).Cardinality()
				}
				return _346_cf35
			})())), 8)
			(_nw67).ArraySet1(_346_cf35, 9)
			(_nw67).ArraySet1(_dafny.IntOfInt64(747), 10)
			(_nw67).ArraySet1(_346_cf35, 11)
			(_nw67).ArraySet1(_346_cf35, 12)
			(_nw67).ArraySet1(_247_v3, 13)
			(_nw67).ArraySet1(_247_v3, 14)
			(_nw67).ArraySet1(Companion_Default___.SafeModuloInt(_346_cf35, (_441_v140).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_441_v140).Cardinality()))).Uint32()).(_dafny.Int)), 15)
			_468_v163 = _nw67
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_464_v161), 0))
			_ = _index42
			(_464_v161).ArraySet1(_346_cf35, (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_464_v161), 0))
			_ = _index43
			(_464_v161).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_451_v150).Cardinality()), _346_cf35), (_437_v136).Cardinality()), (_index43).Int())
		}
	} else if _source8.Is_DC14() {
		var _469___mcc_h3 _dafny.Array = _source8.Get_().(D5_DC14).Cf32
		_ = _469___mcc_h3
		var _470_cf32 _dafny.Array = _469___mcc_h3
		_ = _470_cf32
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_244_v2), 0))
		_ = _index44
		(_244_v2).ArraySet1(_296_v43, (_index44).Int())
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_244_v2), 0))
		_ = _index45
		(_244_v2).ArraySet1(_242_v0, (_index45).Int())
		var _471_v164 _dafny.Sequence
		_ = _471_v164
		_471_v164 = _dafny.UnicodeSeqOfUtf8Bytes("fryvwxa")
		_471_v164 = (func() _dafny.Sequence {
			if (_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool) {
				return _dafny.UnicodeSeqOfUtf8Bytes("h")
			}
			return _471_v164
		})()
		var _472_v165 _dafny.Array
		_ = _472_v165
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_16
		var _nw68 _dafny.Array
		_ = _nw68
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw68 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) _dafny.Int = (func(_473_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_474_i16 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_474_i16, _473_v3)
				}
			})(_247_v3)
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw68 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw68).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw68).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		_472_v165 = _nw68
		var _475_v166 _dafny.Map
		_ = _475_v166
		_475_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(214), _472_v165)
		_475_v166 = (_475_v166).Update(_247_v3, _472_v165)
		var _476_v167 _dafny.CodePoint
		_ = _476_v167
		_476_v167 = _dafny.CodePoint('v')
		var _477_v168 _dafny.Sequence
		_ = _477_v168
		_477_v168 = _dafny.SeqOf(_471_v164)
		var _478_v169 _dafny.Array
		_ = _478_v169
		var _nwElement0_12 _dafny.Sequence = _471_v164
		_ = _nwElement0_12
		var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(9))
		_ = _nw69
		(_nw69).ArraySet1(_nwElement0_12, 0)
		(_nw69).ArraySet1(_471_v164, 1)
		(_nw69).ArraySet1(_471_v164, 2)
		(_nw69).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_471_v164, _471_v164), 3)
		(_nw69).ArraySet1(_471_v164, 4)
		(_nw69).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}(func(_479_i17 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})), 5)
		(_nw69).ArraySet1(_471_v164, 6)
		(_nw69).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_471_v164, _dafny.SeqOf(_476_v167)), 7)
		(_nw69).ArraySet1((_477_v168).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_477_v168).Cardinality()))).Uint32()).(_dafny.Sequence), 8)
		_478_v169 = _nw69
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_17
		var _nw70 _dafny.Array
		_ = _nw70
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw70 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) _dafny.Sequence = (func(_480_v164 _dafny.Sequence, _481_v3 _dafny.Int, _482_v167 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_483_i18 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_480_v164, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(686))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg21 _dafny.Int) interface{} {
							return coer21(arg21)
						}
					}(func(_484_i19 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('j')
					})), (Companion_Default___.SafeIndex(_481_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(686))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}(func(_485_i19 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('j')
					}))).Cardinality()))).Uint32(), _482_v167))
				}
			})(_471_v164, _247_v3, _476_v167)
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw70 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw70).ArraySet1(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw70).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_478_v169 = _nw70
	} else {
		var _486___mcc_h4 D5 = _source8.Get_().(D5_DC18).Cf36
		_ = _486___mcc_h4
		var _487_cf36 D5 = _486___mcc_h4
		_ = _487_cf36
		_247_v3 = _247_v3
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))
		_ = _index46
		(_244_v2).ArraySet1(false, (_index46).Int())
		var _488_v170 _dafny.Sequence
		_ = _488_v170
		_488_v170 = _dafny.UnicodeSeqOfUtf8Bytes("rh")
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))
		_ = _index47
		var _rhs41 _dafny.Int = _dafny.IntOfUint32((_488_v170).Cardinality())
		_ = _rhs41
		var _rhs42 bool = _296_v43
		_ = _rhs42
		var _lhs20 _dafny.Array = _244_v2
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))
		_ = _lhs21
		_247_v3 = _rhs41
		(_lhs20).ArraySet1(_rhs42, (_lhs21).Int())
		var _489_v171 _dafny.Map
		_ = _489_v171
		_489_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool), _247_v3)
		var _490_v172 _dafny.Set
		_ = _490_v172
		_490_v172 = _dafny.SetOf(_247_v3)
		var _491_v173 _dafny.Set
		_ = _491_v173
		_491_v173 = _dafny.SetOf(_490_v172, _490_v172)
		var _492_v174 _dafny.MultiSet
		_ = _492_v174
		_492_v174 = _dafny.MultiSetOf((_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool))
		var _493_v175 D11
		_ = _493_v175
		_493_v175 = Companion_D11_.Create_DC34_(_489_v171)
		var _494_v176 _dafny.Array
		_ = _494_v176
		var _nwElement0_13 _dafny.Map = _489_v171
		_ = _nwElement0_13
		var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(22))
		_ = _nw71
		(_nw71).ArraySet1(_nwElement0_13, 0)
		(_nw71).ArraySet1(_489_v171, 1)
		(_nw71).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v0, _247_v3), 2)
		(_nw71).ArraySet1((_489_v171).Update((_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool), (_491_v173).Cardinality()), 3)
		(_nw71).ArraySet1((_489_v171).Update(_296_v43, (_dafny.Zero).Minus((_492_v174).Cardinality())), 4)
		(_nw71).ArraySet1(_489_v171, 5)
		(_nw71).ArraySet1((_493_v175).Dtor_cf62(), 6)
		(_nw71).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool), _247_v3), 7)
		(_nw71).ArraySet1(_489_v171, 8)
		(_nw71).ArraySet1(_489_v171, 9)
		(_nw71).ArraySet1(_489_v171, 10)
		(_nw71).ArraySet1(_489_v171, 11)
		(_nw71).ArraySet1(_489_v171, 12)
		(_nw71).ArraySet1((_489_v171).Update((_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool), _247_v3), 13)
		(_nw71).ArraySet1(_489_v171, 14)
		(_nw71).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v0, _247_v3), 15)
		(_nw71).ArraySet1(_489_v171, 16)
		(_nw71).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v0, _247_v3), 17)
		(_nw71).ArraySet1(Companion_Default___.Fm39(_247_v3, !((_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool)), _246_globalState), 18)
		(_nw71).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_296_v43, _247_v3), 19)
		(_nw71).ArraySet1(_489_v171, 20)
		(_nw71).ArraySet1(_489_v171, 21)
		_494_v176 = _nw71
		var _495_v177 _dafny.Array
		_ = _495_v177
		var _nwElement0_14 bool = _296_v43
		_ = _nwElement0_14
		var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(4))
		_ = _nw72
		(_nw72).ArraySet1(_nwElement0_14, 0)
		(_nw72).ArraySet1(_296_v43, 1)
		(_nw72).ArraySet1(_242_v0, 2)
		(_nw72).ArraySet1(_296_v43, 3)
		_495_v177 = _nw72
		var _496_v178 *C1
		_ = _496_v178
		var _nw73 *C1 = New_C1_()
		_ = _nw73
		_nw73.Ctor__(_494_v176, !(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_488_v170, _488_v170))), _495_v177)
		_496_v178 = _nw73
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_298_v45), 0))
		_ = _index48
		(_298_v45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_247_v3, (_dafny.Zero).Minus(_247_v3)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gaxsxcrjr")).Cardinality()))), (_index48).Int())
		var _497_v179 _dafny.Sequence
		_ = _497_v179
		_497_v179 = _dafny.SeqOf(_247_v3)
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_298_v45), 0))
		_ = _index49
		(_298_v45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_497_v179, _dafny.Companion_Sequence_.Concatenate(_497_v179, _497_v179)), (_index49).Int())
	}
	var _498_v180 _dafny.CodePoint
	_ = _498_v180
	_498_v180 = _dafny.CodePoint('b')
	_498_v180 = Companion_Default___.Fm8(_498_v180, _246_globalState)
	var _499_v181 _dafny.Array
	_ = _499_v181
	var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
	_ = _nw74
	_499_v181 = _nw74
	var _500_v183 *C1
	_ = _500_v183
	var _nw75 *C1 = New_C1_()
	_ = _nw75
	_nw75.Ctor__(_499_v181, Companion_Default___.Fm4(_247_v3, func() _dafny.Set {
		var _coll30 = _dafny.NewBuilder()
		_ = _coll30
		for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(604), _dafny.IntOfInt64(782))); ; {
			_compr_30, _ok30 := _iter30()
			if !_ok30 {
				break
			}
			var _501_v182 _dafny.Int
			_501_v182 = interface{}(_compr_30).(_dafny.Int)
			if ((_dafny.IntOfInt64(604)).Cmp(_501_v182) <= 0) && ((_501_v182).Cmp(_dafny.IntOfInt64(782)) < 0) {
				_coll30.Add((_501_v182).Minus(_247_v3))
			}
		}
		return _coll30.ToSet()
	}(), _246_globalState), (func() _dafny.Array {
		if _242_v0 {
			return _244_v2
		}
		return _244_v2
	})())
	_500_v183 = _nw75
	var _hi3 _dafny.Int = _247_v3
	_ = _hi3
	for _502_i20 := ((_243_v1).Intersection(_243_v1)).Cardinality(); _502_i20.Cmp(_hi3) < 0; _502_i20 = _502_i20.Plus(_dafny.One) {
		var _503_v184 _dafny.Map
		_ = _503_v184
		_503_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _296_v43)
		var _504_v185 _dafny.MultiSet
		_ = _504_v185
		_504_v185 = _dafny.MultiSetOf(_503_v184)
		var _505_v186 _dafny.Map
		_ = _505_v186
		_505_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_502_i20)), (func() _dafny.Int {
			if (_504_v185).Contains(_503_v184) {
				return (_504_v185).Multiplicity(_503_v184)
			}
			return _502_i20
		})())
		var _506_v187 _dafny.Map
		_ = _506_v187
		_506_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_505_v186).Cardinality(), (_dafny.Zero).Minus((_502_i20).Times(_502_i20)))
		var _507_v188 *C5
		_ = _507_v188
		var _nw76 *C5 = New_C5_()
		_ = _nw76
		_nw76.Ctor__(_296_v43, _244_v2, _499_v181, _242_v0)
		_507_v188 = _nw76
		var _508_v189 _dafny.Sequence
		_ = _508_v189
		_508_v189 = _dafny.SeqOf(_507_v188)
		_506_v187 = (_506_v187).Update(_dafny.IntOfUint32((_508_v189).Cardinality()), _247_v3)
		var _509_v190 _dafny.MultiSet
		_ = _509_v190
		_509_v190 = _dafny.MultiSetOf(_244_v2, _244_v2)
		var _510_v191 _dafny.Map
		_ = _510_v191
		_510_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D13_.Create_DC43_(Companion_D13_.Create_DC40_(_dafny.MultiSetOf((_507_v188).F5(), true))), _dafny.IntOfUint32((Companion_Default___.Fm18(!(_296_v43), _503_v184, _246_globalState)).Cardinality()))
		_247_v3 = (func() _dafny.Int {
			if true {
				return (func() _dafny.Int {
					if (_509_v190).Contains(_244_v2) {
						return (_509_v190).Multiplicity(_244_v2)
					}
					return (_510_v191).Cardinality()
				})()
			}
			return (_502_i20).Minus(_247_v3)
		})()
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_244_v2), 0))
		_ = _index50
		(_244_v2).ArraySet1(true, (_index50).Int())
		var _511_v192 _dafny.Set
		_ = _511_v192
		_511_v192 = _243_v1
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_244_v2), 0))
		_ = _index51
		(_244_v2).ArraySet1((_243_v1).IsProperSubsetOf((_511_v192)), (_index51).Int())
		var _512_v193 _dafny.Array
		_ = _512_v193
		var _len0_18 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_18
		var _nw77 _dafny.Array
		_ = _nw77
		if _len0_18.Cmp(_dafny.Zero) == 0 {
			_nw77 = _dafny.NewArray(_len0_18)
		} else {
			var _init18 func(_dafny.Int) _dafny.Int = (func(_513_i20 _dafny.Int, _514_v187 _dafny.Map, _515_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_516_i21 _dafny.Int) _dafny.Int {
					return (_516_i21).Plus((func() _dafny.Int {
						if (_514_v187).Contains(_513_i20) {
							return (_514_v187).Get(_513_i20).(_dafny.Int)
						}
						return _515_v3
					})())
				}
			})(_502_i20, _506_v187, _247_v3)
			_ = _init18
			var _element0_18 = _init18(_dafny.Zero)
			_ = _element0_18
			_nw77 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
			(_nw77).ArraySet1(_element0_18, 0)
			var _nativeLen0_18 = (_len0_18).Int()
			_ = _nativeLen0_18
			for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
				(_nw77).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
			}
		}
		_512_v193 = _nw77
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_512_v193), 0))
		_ = _index52
		(_512_v193).ArraySet1(_502_i20, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_512_v193), 0))
		_ = _index53
		(_512_v193).ArraySet1(_247_v3, (_index53).Int())
	}
	var _517_v194 _dafny.Sequence
	_ = _517_v194
	_517_v194 = _dafny.UnicodeSeqOfUtf8Bytes("epaistma")
	var _518_v195 D15
	_ = _518_v195
	_518_v195 = Companion_D15_.Create_DC48_(_296_v43, _244_v2, _247_v3, _296_v43, _517_v194)
	var _519_v196 _dafny.Array
	_ = _519_v196
	var _nwElement0_15 _dafny.Array = _244_v2
	_ = _nwElement0_15
	var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(12))
	_ = _nw78
	(_nw78).ArraySet1(_nwElement0_15, 0)
	(_nw78).ArraySet1(_244_v2, 1)
	(_nw78).ArraySet1(_244_v2, 2)
	(_nw78).ArraySet1(_244_v2, 3)
	(_nw78).ArraySet1(_244_v2, 4)
	(_nw78).ArraySet1(_244_v2, 5)
	(_nw78).ArraySet1(_244_v2, 6)
	(_nw78).ArraySet1(_244_v2, 7)
	(_nw78).ArraySet1(_244_v2, 8)
	(_nw78).ArraySet1((_518_v195).Dtor_cf88(), 9)
	(_nw78).ArraySet1(_244_v2, 10)
	(_nw78).ArraySet1(_244_v2, 11)
	_519_v196 = _nw78
	var _520_v197 D15
	_ = _520_v197
	_520_v197 = Companion_D15_.Create_DC47_(_519_v196)
	var _source10 D15 = _520_v197
	_ = _source10
	if _source10.Is_DC48() {
		var _521___mcc_h14 bool = _source10.Get_().(D15_DC48).Cf87
		_ = _521___mcc_h14
		var _522___mcc_h15 _dafny.Array = _source10.Get_().(D15_DC48).Cf88
		_ = _522___mcc_h15
		var _523___mcc_h16 _dafny.Int = _source10.Get_().(D15_DC48).Cf89
		_ = _523___mcc_h16
		var _524___mcc_h17 bool = _source10.Get_().(D15_DC48).Cf90
		_ = _524___mcc_h17
		var _525___mcc_h18 _dafny.Sequence = _source10.Get_().(D15_DC48).Cf91
		_ = _525___mcc_h18
		var _526_cf91 _dafny.Sequence = _525___mcc_h18
		_ = _526_cf91
		var _527_cf90 bool = _524___mcc_h17
		_ = _527_cf90
		var _528_cf89 _dafny.Int = _523___mcc_h16
		_ = _528_cf89
		var _529_cf88 _dafny.Array = _522___mcc_h15
		_ = _529_cf88
		var _530_cf87 bool = _521___mcc_h14
		_ = _530_cf87
		var _531_v198 _dafny.Map
		_ = _531_v198
		_531_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_528_cf89, _529_cf88)
		_531_v198 = (_531_v198).Update(_dafny.IntOfInt64(262), _244_v2)
		var _532_v199 _dafny.Array
		_ = _532_v199
		var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
		_ = _nw79
		_532_v199 = _nw79
		var _533_v200 _dafny.MultiSet
		_ = _533_v200
		_533_v200 = _dafny.MultiSetOf(_528_cf89)
		var _534_v201 *C2
		_ = _534_v201
		var _nw80 *C2 = New_C2_()
		_ = _nw80
		_nw80.Ctor__(_242_v0, _528_cf89, _499_v181, _242_v0, _529_cf88)
		_534_v201 = _nw80
		var _535_v202 _dafny.MultiSet
		_ = _535_v202
		_535_v202 = _dafny.MultiSetOf(_534_v201, _534_v201)
		var _536_v203 D12
		_ = _536_v203
		_536_v203 = Companion_D12_.Create_DC38_(_dafny.SetOf(_528_cf89, (_533_v200).Cardinality(), _dafny.IntOfInt64(358), _dafny.IntOfUint32((_526_cf91).Cardinality())), (_535_v202).Cardinality(), _534_v201, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(943))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}((func(_537_v180 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_538_i22 _dafny.Int) _dafny.CodePoint {
				return _537_v180
			}
		})(_498_v180)))).Cardinality()))
		var _539_v204 D12
		_ = _539_v204
		_539_v204 = Companion_D12_.Create_DC39_(_536_v203)
		var _540_v205 _dafny.Sequence
		_ = _540_v205
		_540_v205 = _dafny.SeqOf(_539_v204)
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_532_v199), 0))
		_ = _index54
		(_532_v199).ArraySet1(_540_v205, (_index54).Int())
		var _541_v206 _dafny.Sequence
		_ = _541_v206
		_541_v206 = _dafny.SeqOf(_247_v3, _528_cf89, _247_v3, (_534_v201).F8())
		var _542_v207 D14
		_ = _542_v207
		_542_v207 = Companion_D14_.Create_DC45_(_dafny.Companion_Sequence_.IsProperPrefixOf(_541_v206, Companion_Default___.Fm31(_247_v3, !(_527_cf90), (_534_v201).F8(), _296_v43, _246_globalState)))
		var _543_v208 _dafny.Map
		_ = _543_v208
		_543_v208 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _530_cf87)
		var _544_v209 D0
		_ = _544_v209
		_544_v209 = Companion_D0_.Create_DC1_(_242_v0, _534_v201.F7, _247_v3)
		var _pat_let_tv3 = _536_v203
		_ = _pat_let_tv3
		var _pat_let_tv4 = _536_v203
		_ = _pat_let_tv4
		var _pat_let_tv5 = _242_v0
		_ = _pat_let_tv5
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_532_v199), 0))
		_ = _index55
		var _rhs43 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func(_pat_let5_0 D12) D12 {
			return func(_545_dt__update__tmp_h0 D12) D12 {
				return func(_pat_let6_0 D12) D12 {
					return func(_546_dt__update_hcf74_h0 D12) D12 {
						return Companion_D12_.Create_DC39_(_546_dt__update_hcf74_h0)
					}(_pat_let6_0)
				}(_pat_let_tv3)
			}(_pat_let5_0)
		}(_539_v204)), _dafny.SeqOf(func(_pat_let7_0 D12) D12 {
			return func(_547_dt__update__tmp_h1 D12) D12 {
				return func(_pat_let8_0 D12) D12 {
					return func(_548_dt__update_hcf74_h1 D12) D12 {
						return Companion_D12_.Create_DC39_(_548_dt__update_hcf74_h1)
					}(_pat_let8_0)
				}(_pat_let_tv4)
			}(_pat_let7_0)
		}(_539_v204), _539_v204)), _dafny.Companion_Sequence_.Concatenate(_540_v205, _dafny.Companion_Sequence_.Update(_540_v205, (Companion_Default___.SafeIndex((_543_v208).Cardinality(), _dafny.IntOfUint32((_540_v205).Cardinality()))).Uint32(), Companion_D12_.Create_DC39_(_536_v203))))
		_ = _rhs43
		var _rhs44 D14 = func(_pat_let9_0 D14) D14 {
			return func(_549_dt__update__tmp_h2 D14) D14 {
				return func(_pat_let10_0 bool) D14 {
					return func(_550_dt__update_hcf84_h0 bool) D14 {
						return Companion_D14_.Create_DC45_(_550_dt__update_hcf84_h0)
					}(_pat_let10_0)
				}(_pat_let_tv5)
			}(_pat_let9_0)
		}(_542_v207)
		_ = _rhs44
		var _rhs45 _dafny.Array = _244_v2
		_ = _rhs45
		var _rhs46 bool = _530_cf87
		_ = _rhs46
		var _rhs47 bool = Companion_Default___.Fm4(((_541_v206).Select((Companion_Default___.SafeIndex((_544_v209).Dtor_cf3(), _dafny.IntOfUint32((_541_v206).Cardinality()))).Uint32()).(_dafny.Int)).Times((_534_v201).F8()), Companion_Default___.Fm29(_246_globalState), _246_globalState)
		_ = _rhs47
		var _lhs22 _dafny.Array = _532_v199
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_532_v199), 0))
		_ = _lhs23
		(_lhs22).ArraySet1(_rhs43, (_lhs23).Int())
		_542_v207 = _rhs44
		_244_v2 = _rhs45
		_527_cf90 = _rhs46
		_296_v43 = _rhs47
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_244_v2), 0))
		_ = _index56
		(_244_v2).ArraySet1((func() bool {
			if _530_cf87 {
				return !(true)
			}
			return _242_v0
		})(), (_index56).Int())
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_244_v2), 0))
		_ = _index57
		(_244_v2).ArraySet1(false, (_index57).Int())
		_498_v180 = _498_v180
	} else if _source10.Is_DC47() {
		var _551___mcc_h19 _dafny.Array = _source10.Get_().(D15_DC47).Cf86
		_ = _551___mcc_h19
		var _552_cf86 _dafny.Array = _551___mcc_h19
		_ = _552_cf86
		var _553_v210 *C4
		_ = _553_v210
		var _nw81 *C4 = New_C4_()
		_ = _nw81
		_nw81.Ctor__(true)
		_553_v210 = _nw81
		var _554_v211 _dafny.Sequence
		_ = _554_v211
		_554_v211 = _dafny.SeqOf(_553_v210, _553_v210)
		var _rhs48 bool = (_553_v210).F6()
		_ = _rhs48
		var _rhs49 *C4 = (func() *C4 {
			if _296_v43 {
				return (_554_v211).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(true), (Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Uint32(), (_553_v210).F6())).Cardinality()), _dafny.IntOfUint32((_554_v211).Cardinality()))).Uint32()).(*C4)
			}
			return _553_v210
		})()
		_ = _rhs49
		var _rhs50 _dafny.CodePoint = (func() _dafny.CodePoint {
			if _242_v0 {
				return _498_v180
			}
			return _498_v180
		})()
		_ = _rhs50
		var _rhs51 _dafny.Int = (_500_v183).Fm3(_242_v0, false, _247_v3, _247_v3, _246_globalState)
		_ = _rhs51
		_296_v43 = _rhs48
		_553_v210 = _rhs49
		_498_v180 = _rhs50
		_247_v3 = _rhs51
		var _555_v212 _dafny.Sequence
		_ = _555_v212
		_555_v212 = _dafny.SeqOf(_dafny.IntOfUint32((_517_v194).Cardinality()), _247_v3, _247_v3, _247_v3)
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_298_v45), 0))
		_ = _index58
		(_298_v45).ArraySet1(_555_v212, (_index58).Int())
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_298_v45), 0))
		_ = _index59
		(_298_v45).ArraySet1(_555_v212, (_index59).Int())
		var _556_v213 _dafny.Sequence
		_ = _556_v213
		_556_v213 = _dafny.SeqOf(_296_v43)
		var _557_v214 _dafny.Map
		_ = _557_v214
		_557_v214 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_556_v213).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_556_v213).Cardinality()))).Uint32()).(bool), _296_v43)
		var _558_v215 _dafny.MultiSet
		_ = _558_v215
		_558_v215 = _dafny.MultiSetOf(_557_v214, Companion_Default___.Fm36(_242_v0, _246_globalState))
		var _559_v216 _dafny.Sequence
		_ = _559_v216
		_559_v216 = _dafny.SeqOf((_558_v215).IsDisjointFrom(_558_v215), Companion_Default___.Fm4(_247_v3, _dafny.SetOf(_247_v3), _246_globalState), (_553_v210).F6(), _242_v0)
		_242_v0 = (_559_v216).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_559_v216).Cardinality()))).Uint32()).(bool)
		var _560_v217 _dafny.Map
		_ = _560_v217
		_560_v217 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v0, _517_v194)
		var _561_v218 *C4
		_ = _561_v218
		var _nw82 *C4 = New_C4_()
		_ = _nw82
		_nw82.Ctor__((_560_v217).Contains((_553_v210).F6()))
		_561_v218 = _nw82
	} else {
		var _562___mcc_h20 D15 = _source10.Get_().(D15_DC49).Cf92
		_ = _562___mcc_h20
		var _563_cf92 D15 = _562___mcc_h20
		_ = _563_cf92
		var _564_v219 _dafny.Array
		_ = _564_v219
		var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
		_ = _nw83
		_564_v219 = _nw83
		_564_v219 = _564_v219
		_247_v3 = _dafny.IntOfInt64(294)
		_242_v0 = _296_v43
		var _565_v220 _dafny.Sequence
		_ = _565_v220
		_565_v220 = _dafny.SeqOf(_296_v43, _242_v0, _296_v43)
		if !((_565_v220).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_565_v220).Cardinality()))).Uint32()).(bool)) {
			var _566_v221 _dafny.Array
			_ = _566_v221
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_19
			var _nw84 _dafny.Array
			_ = _nw84
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw84 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Int = (func(_567_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_568_i23 _dafny.Int) _dafny.Int {
						return (_568_i23).Plus(_567_v3)
					}
				})(_247_v3)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw84 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw84).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw84).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_566_v221 = _nw84
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))
			_ = _index60
			(_566_v221).ArraySet1(_247_v3, (_index60).Int())
			var _569_v222 _dafny.Map
			_ = _569_v222
			_569_v222 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_566_v221, _296_v43)
			var _570_v223 D18
			_ = _570_v223
			_570_v223 = Companion_D18_.Create_DC54_(true, _296_v43, _247_v3, _dafny.IntOfUint32((_517_v194).Cardinality()))
			var _571_v224 _dafny.Map
			_ = _571_v224
			_571_v224 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _296_v43)
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))
			_ = _index61
			var _rhs52 _dafny.Int = Companion_Default___.SafeModuloInt((_569_v222).Cardinality(), (_247_v3).Plus(_247_v3))
			_ = _rhs52
			var _rhs53 _dafny.Int = (func() _dafny.Int {
				if (_570_v223).Dtor_cf100() {
					return (_247_v3).Times((_571_v224).Cardinality())
				}
				return _247_v3
			})()
			_ = _rhs53
			var _lhs24 _dafny.Array = _566_v221
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))
			_ = _lhs25
			(_lhs24).ArraySet1(_rhs52, (_lhs25).Int())
			_247_v3 = _rhs53
			var _572_v225 _dafny.Array
			_ = _572_v225
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw85
			_572_v225 = _nw85
			var _573_v227 _dafny.Array
			_ = _573_v227
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_20
			var _nw86 _dafny.Array
			_ = _nw86
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw86 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Map = (func(_574_v221 _dafny.Array, _575_v3 _dafny.Int, _576_v194 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
					return func(_577_i24 _dafny.Int) _dafny.Map {
						return func() _dafny.Map {
							var _coll31 = _dafny.NewMapBuilder()
							_ = _coll31
							for _iter31 := _dafny.Iterate((_dafny.SeqOf((_574_v221).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_574_v221), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(461))).Elements()); ; {
								_compr_31, _ok31 := _iter31()
								if !_ok31 {
									break
								}
								var _578_v226 _dafny.Int
								_578_v226 = interface{}(_compr_31).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_574_v221).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_574_v221), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(461)), _578_v226) {
									_coll31.Add((_578_v226).Plus(_575_v3), _dafny.IntOfUint32((_576_v194).Cardinality()))
								}
							}
							return _coll31.ToMap()
						}()
					}
				})(_566_v221, _247_v3, _517_v194)
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw86 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw86).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw86).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_573_v227 = _nw86
			var _579_v228 _dafny.MultiSet
			_ = _579_v228
			var _580_v229 bool
			_ = _580_v229
			var _out15 _dafny.MultiSet
			_ = _out15
			var _out16 bool
			_ = _out16
			_out15, _out16 = (_500_v183).M1((_566_v221).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))).Int()).(_dafny.Int), _498_v180, _572_v225, _573_v227, _246_globalState)
			_579_v228 = _out15
			_580_v229 = _out16
			var _581_v230 _dafny.MultiSet
			_ = _581_v230
			_581_v230 = _dafny.MultiSetOf(_580_v229)
			var _582_v231 _dafny.Map
			_ = _582_v231
			_582_v231 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_572_v225, (_247_v3).Minus((_dafny.Zero).Minus((_581_v230).Cardinality())))
			_582_v231 = (_582_v231).Update(_572_v225, _dafny.IntOfInt64(379))
			var _583_v232 _dafny.MultiSet
			_ = _583_v232
			_583_v232 = _dafny.MultiSetOf(_dafny.IntOfInt64(929), (_566_v221).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))).Int()).(_dafny.Int), (_566_v221).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _296_v43 {
					return _dafny.UnicodeSeqOfUtf8Bytes("yuc")
				}
				return _517_v194
			})()).Cardinality()))
			_583_v232 = (_dafny.MultiSetOf((_566_v221).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))).Int()).(_dafny.Int))).Intersection(_583_v232)
			var _584_v233 _dafny.Map
			_ = _584_v233
			_584_v233 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm8(_498_v180, _246_globalState), _498_v180)
			_247_v3 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_584_v233).Cardinality()), (_566_v221).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))).Int()).(_dafny.Int)), (_dafny.Zero).Minus((_566_v221).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_566_v221), 0))).Int()).(_dafny.Int)))
		} else {
			var _585_v234 *C3
			_ = _585_v234
			var _nw87 *C3 = New_C3_()
			_ = _nw87
			_nw87.Ctor__(_244_v2)
			_585_v234 = _nw87
			var _586_v235 _dafny.Sequence
			_ = _586_v235
			_586_v235 = _dafny.SeqOf(_244_v2, _244_v2, _244_v2)
			var _587_v236 _dafny.Map
			_ = _587_v236
			_587_v236 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v0, _244_v2)
			var _588_v237 bool
			_ = _588_v237
			var _out17 bool
			_ = _out17
			_out17 = (_500_v183).M0(_247_v3, _586_v235, _587_v236, _246_globalState)
			_588_v237 = _out17
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_244_v2), 0))
			_ = _index62
			(_244_v2).ArraySet1(_296_v43, (_index62).Int())
			var _589_v238 _dafny.Map
			_ = _589_v238
			_589_v238 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _588_v237)
			var _590_v239 _dafny.Map
			_ = _590_v239
			_590_v239 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if _242_v0 {
					return _dafny.IntOfUint32((_517_v194).Cardinality())
				}
				return _247_v3
			})(), ((_589_v238).Cardinality()).Times(_247_v3))
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_244_v2), 0))
			_ = _index63
			var _rhs54 bool = !_dafny.Companion_Sequence_.Equal(_517_v194, _517_v194)
			_ = _rhs54
			var _rhs55 _dafny.Map = Companion_Default___.Fm25(_296_v43, _246_globalState)
			_ = _rhs55
			var _rhs56 bool = Companion_Default___.Fm20(_296_v43, _246_globalState)
			_ = _rhs56
			var _lhs26 _dafny.Array = _244_v2
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_244_v2), 0))
			_ = _lhs27
			(_lhs26).ArraySet1(_rhs54, (_lhs27).Int())
			_590_v239 = _rhs55
			_242_v0 = _rhs56
			_498_v180 = _498_v180
			var _591_v240 _dafny.Sequence
			_ = _591_v240
			_591_v240 = _dafny.SeqOf(_500_v183)
			var _592_v241 _dafny.Map
			_ = _592_v241
			_592_v241 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_591_v240), true)
			var _593_v242 _dafny.MultiSet
			_ = _593_v242
			_593_v242 = _dafny.MultiSetOf(_500_v183)
			var _594_v243 _dafny.Map
			_ = _594_v243
			_594_v243 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_588_v237, true)
			var _595_v244 _dafny.Array
			_ = _595_v244
			var _nwElement0_16 _dafny.Int = _247_v3
			_ = _nwElement0_16
			var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(11))
			_ = _nw88
			(_nw88).ArraySet1(_nwElement0_16, 0)
			(_nw88).ArraySet1((_dafny.Zero).Minus(_247_v3), 1)
			(_nw88).ArraySet1((_247_v3).Plus(_dafny.IntOfUint32((Companion_Default___.Fm11(_dafny.IntOfInt64(615), (func() bool {
				if (_592_v241).Contains(_593_v242) {
					return (_592_v241).Get(_593_v242).(bool)
				}
				return (_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool)
			})(), _247_v3, _246_globalState)).Cardinality())), 2)
			(_nw88).ArraySet1(_247_v3, 3)
			(_nw88).ArraySet1(_dafny.IntOfUint32((_517_v194).Cardinality()), 4)
			(_nw88).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_594_v243).Cardinality()), _247_v3), 5)
			(_nw88).ArraySet1((_500_v183).Fm16((_244_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_244_v2), 0))).Int()).(bool), _246_globalState), 6)
			(_nw88).ArraySet1(_247_v3, 7)
			(_nw88).ArraySet1(_247_v3, 8)
			(_nw88).ArraySet1(_247_v3, 9)
			(_nw88).ArraySet1(_247_v3, 10)
			_595_v244 = _nw88
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_595_v244), 0))
			_ = _index64
			(_595_v244).ArraySet1(_247_v3, (_index64).Int())
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_595_v244), 0))
			_ = _index65
			(_595_v244).ArraySet1(_247_v3, (_index65).Int())
		}
	}
	var _hi4 _dafny.Int = _247_v3
	_ = _hi4
	for _596_i25 := _247_v3; _596_i25.Cmp(_hi4) < 0; _596_i25 = _596_i25.Plus(_dafny.One) {
		var _597_v245 bool
		_ = _597_v245
		var _598_v246 _dafny.Map
		_ = _598_v246
		var _out18 bool
		_ = _out18
		var _out19 _dafny.Map
		_ = _out19
		_out18, _out19 = Companion_Default___.M6(_246_globalState)
		_597_v245 = _out18
		_598_v246 = _out19
		_498_v180 = _498_v180
		var _599_v247 _dafny.Map
		_ = _599_v247
		_599_v247 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_596_i25, _242_v0)
		var _600_v248 _dafny.Sequence
		_ = _600_v248
		_600_v248 = _dafny.SeqOf(_242_v0)
		var _601_v249 _dafny.Sequence
		_ = _601_v249
		_601_v249 = _dafny.SeqOf((_599_v247).Cardinality(), _dafny.IntOfUint32((_600_v248).Cardinality()), (_243_v1).Cardinality())
		var _602_v250 *C5
		_ = _602_v250
		var _nw89 *C5 = New_C5_()
		_ = _nw89
		_nw89.Ctor__(_597_v245, _244_v2, _499_v181, _296_v43)
		_602_v250 = _nw89
		var _603_v251 _dafny.MultiSet
		_ = _603_v251
		_603_v251 = _dafny.MultiSetOf(_602_v250, _602_v250, _602_v250)
		var _604_v252 *C2
		_ = _604_v252
		var _nw90 *C2 = New_C2_()
		_ = _nw90
		_nw90.Ctor__(_597_v245, Companion_Default___.SafeModuloInt((_601_v249).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.IntOfUint32((_601_v249).Cardinality()))).Uint32()).(_dafny.Int), _596_i25), _499_v181, (_603_v251).IsSubsetOf(_603_v251), _244_v2)
		_604_v252 = _nw90
		_604_v252 = _604_v252
		(_604_v252).F7 = (_600_v248).Select((Companion_Default___.SafeIndex(_596_i25, _dafny.IntOfUint32((_600_v248).Cardinality()))).Uint32()).(bool)
	}
	var _605_v253 _dafny.Array
	_ = _605_v253
	var _len0_21 _dafny.Int = _dafny.IntOfInt64(23)
	_ = _len0_21
	var _nw91 _dafny.Array
	_ = _nw91
	if _len0_21.Cmp(_dafny.Zero) == 0 {
		_nw91 = _dafny.NewArray(_len0_21)
	} else {
		var _init21 func(_dafny.Int) D0 = (func(_606_v0 bool) func(_dafny.Int) D0 {
			return func(_607_i27 _dafny.Int) D0 {
				return Companion_D0_.Create_DC0_(_606_v0)
			}
		})(_242_v0)
		_ = _init21
		var _element0_21 = _init21(_dafny.Zero)
		_ = _element0_21
		_nw91 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
		(_nw91).ArraySet1(_element0_21, 0)
		var _nativeLen0_21 = (_len0_21).Int()
		_ = _nativeLen0_21
		for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
			(_nw91).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
		}
	}
	_605_v253 = _nw91
	for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_605_v253), 0))); ; {
		_guard_loop_0, _ok32 := _iter32()
		if !_ok32 {
			break
		}
		var _608_i26 _dafny.Int
		_608_i26 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_608_i26).Sign() != -1) && ((_608_i26).Cmp(_dafny.ArrayLen((_605_v253), 0)) < 0)) {
			(_605_v253).ArraySet1(Companion_D0_.Create_DC0_((_dafny.MultiSetOf(_247_v3)).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(666))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_609_v0 bool, _610_v3 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_611_i28 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_609_v0, _610_v3)
				}
			})(_242_v0, _247_v3)))).Cardinality())))), (_608_i26).Int())
		}
	}
	var _612_v254 _dafny.Map
	_ = _612_v254
	_612_v254 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v0, _247_v3)).Cardinality(), _247_v3)
	var _613_v255 _dafny.Sequence
	_ = _613_v255
	_613_v255 = _dafny.SeqOf(_247_v3)
	var _614_v256 _dafny.Map
	_ = _614_v256
	_614_v256 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_296_v43, _dafny.IntOfUint32((_613_v255).Cardinality()))
	var _615_v257 _dafny.Map
	_ = _615_v257
	_615_v257 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_v256, _612_v254)
	var _616_v259 D0
	_ = _616_v259
	_616_v259 = Companion_D0_.Create_DC2_(_247_v3, _296_v43, _296_v43, _247_v3)
	var _617_v260 _dafny.Set
	_ = _617_v260
	_617_v260 = _dafny.SetOf((_616_v259).Dtor_cf4(), _247_v3)
	var _618_v261 _dafny.Sequence
	_ = _618_v261
	_618_v261 = _dafny.SeqOf(_242_v0, _296_v43)
	var _619_v262 _dafny.Array
	_ = _619_v262
	var _nwElement0_17 _dafny.Map = Companion_Default___.Fm25(_296_v43, _246_globalState)
	_ = _nwElement0_17
	var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(17))
	_ = _nw92
	(_nw92).ArraySet1(_nwElement0_17, 0)
	(_nw92).ArraySet1((_612_v254).Merge(((_612_v254).Update(_247_v3, _247_v3)).Update(_247_v3, _247_v3)), 1)
	(_nw92).ArraySet1(_612_v254, 2)
	(_nw92).ArraySet1((_612_v254).Merge(_612_v254), 3)
	(_nw92).ArraySet1(_612_v254, 4)
	(_nw92).ArraySet1((func() _dafny.Map {
		if (_615_v257).Contains(_614_v256) {
			return (_615_v257).Get(_614_v256).(_dafny.Map)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _247_v3)
	})(), 5)
	(_nw92).ArraySet1(_612_v254, 6)
	(_nw92).ArraySet1(func() _dafny.Map {
		var _coll32 = _dafny.NewMapBuilder()
		_ = _coll32
		for _iter33 := _dafny.Iterate((_617_v260).Elements()); ; {
			_compr_32, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _620_v258 _dafny.Int
			_620_v258 = interface{}(_compr_32).(_dafny.Int)
			if (_617_v260).Contains(_620_v258) {
				_coll32.Add((_620_v258).Times(_247_v3), _247_v3)
			}
		}
		return _coll32.ToMap()
	}(), 7)
	(_nw92).ArraySet1((_612_v254).Merge(_612_v254), 8)
	(_nw92).ArraySet1((_612_v254).Merge(_612_v254), 9)
	(_nw92).ArraySet1((_612_v254).Update(_247_v3, (_dafny.Zero).Minus((_500_v183).Fm3((_618_v261).Select((Companion_Default___.SafeIndex(_247_v3, _dafny.IntOfUint32((_618_v261).Cardinality()))).Uint32()).(bool), _242_v0, _dafny.IntOfUint32((_517_v194).Cardinality()), _247_v3, _246_globalState))), 10)
	(_nw92).ArraySet1(_612_v254, 11)
	(_nw92).ArraySet1(_612_v254, 12)
	(_nw92).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(191))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}((func(_621_v180 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_622_i30 _dafny.Int) _dafny.CodePoint {
			return _621_v180
		}
	})(_498_v180)))).Cardinality()), _dafny.IntOfUint32((_517_v194).Cardinality())), 13)
	(_nw92).ArraySet1(_612_v254, 14)
	(_nw92).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _247_v3), 15)
	(_nw92).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("q")).Cardinality())), 16)
	_619_v262 = _nw92
	for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_619_v262), 0))); ; {
		_guard_loop_1, _ok34 := _iter34()
		if !_ok34 {
			break
		}
		var _623_i29 _dafny.Int
		_623_i29 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_623_i29).Sign() != -1) && ((_623_i29).Cmp(_dafny.ArrayLen((_619_v262), 0)) < 0)) {
			(_619_v262).ArraySet1(((_612_v254).Merge(_612_v254)).Merge(_612_v254), (_623_i29).Int())
		}
	}
	var _624_v263 _dafny.Sequence
	_ = _624_v263
	_624_v263 = _dafny.SeqOf(_244_v2, _244_v2)
	var _625_v264 bool
	_ = _625_v264
	var _out20 bool
	_ = _out20
	_out20 = (_500_v183).M0(_247_v3, _dafny.Companion_Sequence_.Concatenate(_624_v263, _624_v263), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_242_v0, (_518_v195).Dtor_cf88()), _246_globalState)
	_625_v264 = _out20
	var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_244_v2), 0))
	_ = _index66
	(_244_v2).ArraySet1(_296_v43, (_index66).Int())
	var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_244_v2), 0))
	_ = _index67
	(_244_v2).ArraySet1(!(false) || ((func() bool {
		if _625_v264 {
			return _242_v0
		}
		return _296_v43
	})()), (_index67).Int())
	var _626_i31 _dafny.Int
	_ = _626_i31
	_626_i31 = _dafny.Zero
	{
		for (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_247_v3), _dafny.IntOfInt64(428))).Cmp(_247_v3) <= 0 {
			{
				if (_626_i31).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_626_i31 = (_626_i31).Plus(_dafny.One)
				var _627_v265 _dafny.Array
				_ = _627_v265
				var _nwElement0_18 _dafny.Int = _247_v3
				_ = _nwElement0_18
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(25))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_18, 0)
				(_nw93).ArraySet1((func() _dafny.Int {
					if (_612_v254).Contains(_dafny.IntOfUint32((_517_v194).Cardinality())) {
						return (_612_v254).Get(_dafny.IntOfUint32((_517_v194).Cardinality())).(_dafny.Int)
					}
					return _247_v3
				})(), 1)
				(_nw93).ArraySet1((_247_v3).Plus(_247_v3), 2)
				(_nw93).ArraySet1(_247_v3, 3)
				(_nw93).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-363)), 4)
				(_nw93).ArraySet1((_247_v3).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hkxt")).Cardinality())), 5)
				(_nw93).ArraySet1(_247_v3, 6)
				(_nw93).ArraySet1((_247_v3).Times(_247_v3), 7)
				(_nw93).ArraySet1(_247_v3, 8)
				(_nw93).ArraySet1(_247_v3, 9)
				(_nw93).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(228), _247_v3), 10)
				(_nw93).ArraySet1(_247_v3, 11)
				(_nw93).ArraySet1(Companion_Default___.SafeDivisionInt(_247_v3, _247_v3), 12)
				(_nw93).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(234))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}((func(_628_v180 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_629_i32 _dafny.Int) _dafny.CodePoint {
						return _628_v180
					}
				})(_498_v180)))).Cardinality())).Minus(_247_v3), 13)
				(_nw93).ArraySet1(_247_v3, 14)
				(_nw93).ArraySet1((_247_v3).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pmdl")).Cardinality()))), 15)
				(_nw93).ArraySet1(_247_v3, 16)
				(_nw93).ArraySet1(_247_v3, 17)
				(_nw93).ArraySet1(Companion_Default___.SafeModuloInt(_247_v3, _247_v3), 18)
				(_nw93).ArraySet1((_247_v3).Times(Companion_Default___.Fm0(_dafny.IntOfInt64(-471), _246_globalState)), 19)
				(_nw93).ArraySet1(_dafny.IntOfInt64(-39), 20)
				(_nw93).ArraySet1(_247_v3, 21)
				(_nw93).ArraySet1((_dafny.Zero).Minus(_247_v3), 22)
				(_nw93).ArraySet1(_247_v3, 23)
				(_nw93).ArraySet1((_247_v3).Times(_247_v3), 24)
				_627_v265 = _nw93
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_627_v265), 0))
				_ = _index68
				(_627_v265).ArraySet1(_247_v3, (_index68).Int())
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_627_v265), 0))
				_ = _index69
				(_627_v265).ArraySet1(_247_v3, (_index69).Int())
				var _630_v266 _dafny.Map
				_ = _630_v266
				_630_v266 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_625_v264), _244_v2)
				var _631_v267 bool
				_ = _631_v267
				var _out21 bool
				_ = _out21
				_out21 = (_500_v183).M0(_dafny.IntOfInt64(290), _dafny.SeqOf(_244_v2, _244_v2, _244_v2), _630_v266, _246_globalState)
				_631_v267 = _out21
				var _632_v268 _dafny.MultiSet
				_ = _632_v268
				_632_v268 = _dafny.MultiSetOf(_498_v180)
				var _633_v269 T1
				_ = _633_v269
				var _nw94 *C2 = New_C2_()
				_ = _nw94
				_nw94.Ctor__(_242_v0, _dafny.IntOfUint32((_618_v261).Cardinality()), _499_v181, _625_v264, _244_v2)
				_633_v269 = _nw94
				var _634_v270 D4
				_ = _634_v270
				_634_v270 = Companion_D4_.Create_DC13_(_242_v0, _632_v268, (_618_v261).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_633_v269)).Cardinality()), _dafny.IntOfUint32((_618_v261).Cardinality()))).Uint32()).(bool))
				_242_v0 = (_634_v270).Dtor_cf29()
				var _635_v271 _dafny.Array
				_ = _635_v271
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_22
				var _nw95 _dafny.Array
				_ = _nw95
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw95 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Sequence = (func(_636_v194 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_637_i33 _dafny.Int) _dafny.Sequence {
							return _636_v194
						}
					})(_517_v194)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw95 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw95).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw95).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_635_v271 = _nw95
				var _638_v272 _dafny.MultiSet
				_ = _638_v272
				var _639_v273 bool
				_ = _639_v273
				var _out22 _dafny.MultiSet
				_ = _out22
				var _out23 bool
				_ = _out23
				_out22, _out23 = (_500_v183).M1((_627_v265).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_627_v265), 0))).Int()).(_dafny.Int), _dafny.CodePoint('t'), _635_v271, _619_v262, _246_globalState)
				_638_v272 = _out22
				_639_v273 = _out23
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _640_v274 _dafny.Array
	_ = _640_v274
	var _len0_23 _dafny.Int = _dafny.IntOfInt64(14)
	_ = _len0_23
	var _nw96 _dafny.Array
	_ = _nw96
	if _len0_23.Cmp(_dafny.Zero) == 0 {
		_nw96 = _dafny.NewArray(_len0_23)
	} else {
		var _init23 func(_dafny.Int) _dafny.Sequence = (func(_641_v194 _dafny.Sequence, _642_v3 _dafny.Int, _643_v180 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
			return func(_644_i34 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_641_v194, (Companion_Default___.SafeIndex(_642_v3, _dafny.IntOfUint32((_641_v194).Cardinality()))).Uint32(), _643_v180), _641_v194)
			}
		})(_517_v194, _247_v3, _498_v180)
		_ = _init23
		var _element0_23 = _init23(_dafny.Zero)
		_ = _element0_23
		_nw96 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
		(_nw96).ArraySet1(_element0_23, 0)
		var _nativeLen0_23 = (_len0_23).Int()
		_ = _nativeLen0_23
		for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
			(_nw96).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
		}
	}
	_640_v274 = _nw96
	var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_640_v274), 0))
	_ = _index70
	(_640_v274).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(737))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg27 _dafny.Int) interface{} {
			return coer27(arg27)
		}
	}(func(_645_i35 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	})), (_index70).Int())
	var _646_v275 _dafny.Array
	_ = _646_v275
	var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
	_ = _nw97
	_646_v275 = _nw97
	var _647_v276 _dafny.Sequence
	_ = _647_v276
	_647_v276 = _dafny.SeqOf(_646_v275)
	var _648_v277 _dafny.Map
	_ = _648_v277
	_648_v277 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_647_v276, _517_v194)
	var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_640_v274), 0))
	_ = _index71
	(_640_v274).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
		if (_648_v277).Contains(_dafny.SeqOf(_646_v275)) {
			return (_648_v277).Get(_dafny.SeqOf(_646_v275)).(_dafny.Sequence)
		}
		return (_518_v195).Dtor_cf91()
	})(), (Companion_Default___.SafeIndex(((_500_v183).Fm2(_246_globalState)).Plus(_247_v3), _dafny.IntOfUint32(((func() _dafny.Sequence {
		if (_648_v277).Contains(_dafny.SeqOf(_646_v275)) {
			return (_648_v277).Get(_dafny.SeqOf(_646_v275)).(_dafny.Sequence)
		}
		return (_518_v195).Dtor_cf91()
	})()).Cardinality()))).Uint32(), _498_v180), (_index71).Int())
	var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_244_v2), 0))
	_ = _index72
	(_244_v2).ArraySet1((_dafny.SetOf(_247_v3, _247_v3)).IsProperSubsetOf(_dafny.SetOf(_247_v3)), (_index72).Int())
	var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_640_v274), 0))
	_ = _index73
	(_640_v274).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(313))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}((func(_649_v180 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_650_i36 _dafny.Int) _dafny.CodePoint {
			return _649_v180
		}
	})(_498_v180))), (func() _dafny.Sequence {
		if _242_v0 {
			return _dafny.UnicodeSeqOfUtf8Bytes("ndi")
		}
		return (_640_v274).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_640_v274), 0))).Int()).(_dafny.Sequence)
	})()), (_index73).Int())
	_dafny.Print(_242_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_243_v1).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v2).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_246_globalState).F0()).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_246_globalState.F1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_247_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_296_v43)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_297_v44).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_298_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(8), _dafny.Zero, _dafny.Zero, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_299_v46).Dtor_cf32()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(8), _dafny.Zero, _dafny.Zero, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_498_v180)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_517_v194.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_518_v195).Dtor_cf87())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_518_v195).Dtor_cf88()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_518_v195).Dtor_cf89())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_518_v195).Dtor_cf90())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_518_v195).Dtor_cf91().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_519_v196).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo(((_520_v197).Dtor_cf86()).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.Zero).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.One).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_605_v253).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(D0)).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_612_v254).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_613_v255, _dafny.SeqOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_614_v256).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_615_v257).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.One), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_616_v259).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_616_v259).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_616_v259).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_616_v259).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_617_v260).Equals(_dafny.SetOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_618_v261, _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_619_v262).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_624_v263).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_625_v264)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_626_i31)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_640_v274).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_647_v276).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_648_v277).Cardinality())
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
	Cf2 bool
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 bool, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 _dafny.Int
	Cf5 bool
	Cf6 bool
	Cf7 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 _dafny.Int, Cf5 bool, Cf6 bool, Cf7 _dafny.Int) D0 {
	return D0{D0_DC2{Cf4, Cf5, Cf6, Cf7}}
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

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() bool {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf7
}

func (_this D0) Dtor_cf0() bool {
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
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5 && data1.Cf6 == data2.Cf6 && data1.Cf7.Cmp(data2.Cf7) == 0
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
	Cf9  _dafny.Map
	Cf10 bool
	Cf11 _dafny.Int
	Cf12 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf9 _dafny.Map, Cf10 bool, Cf11 _dafny.Int, Cf12 bool) D1 {
	return D1{D1_DC4{Cf9, Cf10, Cf11, Cf12}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf8 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf8 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf8}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.EmptyMap, false, _dafny.Zero, false)
}

func (_this D1) Dtor_cf9() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf12() bool {
	return _this.Get_().(D1_DC4).Cf12
}

func (_this D1) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf9.Equals(data2.Cf9) && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12 == data2.Cf12
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf14 _dafny.CodePoint
	Cf15 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf14 _dafny.CodePoint, Cf15 _dafny.Int) D2 {
	return D2{D2_DC6{Cf14, Cf15}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf16 _dafny.Int
	Cf17 _dafny.Int
	Cf18 _dafny.MultiSet
	Cf19 _dafny.Int
	Cf20 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf16 _dafny.Int, Cf17 _dafny.Int, Cf18 _dafny.MultiSet, Cf19 _dafny.Int, Cf20 _dafny.Int) D2 {
	return D2{D2_DC7{Cf16, Cf17, Cf18, Cf19, Cf20}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC5 struct {
	Cf13 _dafny.Sequence
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf13 _dafny.Sequence) D2 {
	return D2{D2_DC5{Cf13}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC8 struct {
	Cf21 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf21 D2) D2 {
	return D2{D2_DC8{Cf21}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D2) Dtor_cf14() _dafny.CodePoint {
	return _this.Get_().(D2_DC6).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf15
}

func (_this D2) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf16
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf17
}

func (_this D2) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D2_DC7).Cf18
}

func (_this D2) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf19
}

func (_this D2) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf20
}

func (_this D2) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf13
}

func (_this D2) Dtor_cf21() D2 {
	return _this.Get_().(D2_DC8).Cf21
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + data.Cf13.VerbatimString(true) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18.Equals(data2.Cf18) && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf13.Equals(data2.Cf13)
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf23 bool
	Cf24 _dafny.CodePoint
	Cf25 bool
	Cf26 bool
	Cf27 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf23 bool, Cf24 _dafny.CodePoint, Cf25 bool, Cf26 bool, Cf27 bool) D3 {
	return D3{D3_DC10{Cf23, Cf24, Cf25, Cf26, Cf27}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf22 _dafny.Map
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf22 _dafny.Map) D3 {
	return D3{D3_DC9{Cf22}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, _dafny.CodePoint('D'), false, false, false)
}

func (_this D3) Dtor_cf23() bool {
	return _this.Get_().(D3_DC10).Cf23
}

func (_this D3) Dtor_cf24() _dafny.CodePoint {
	return _this.Get_().(D3_DC10).Cf24
}

func (_this D3) Dtor_cf25() bool {
	return _this.Get_().(D3_DC10).Cf25
}

func (_this D3) Dtor_cf26() bool {
	return _this.Get_().(D3_DC10).Cf26
}

func (_this D3) Dtor_cf27() bool {
	return _this.Get_().(D3_DC10).Cf27
}

func (_this D3) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D3_DC9).Cf22
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24 == data2.Cf24 && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf22.Equals(data2.Cf22)
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
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_() D4 {
	return D4{D4_DC12{}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf29 bool
	Cf30 _dafny.MultiSet
	Cf31 bool
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf29 bool, Cf30 _dafny.MultiSet, Cf31 bool) D4 {
	return D4{D4_DC13{Cf29, Cf30, Cf31}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC11 struct {
	Cf28 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf28 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf28}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_()
}

func (_this D4) Dtor_cf29() bool {
	return _this.Get_().(D4_DC13).Cf29
}

func (_this D4) Dtor_cf30() _dafny.MultiSet {
	return _this.Get_().(D4_DC13).Cf30
}

func (_this D4) Dtor_cf31() bool {
	return _this.Get_().(D4_DC13).Cf31
}

func (_this D4) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf28
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf28) + ")"
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
			_, ok := other.Get_().(D4_DC12)
			return ok
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf29 == data2.Cf29 && data1.Cf30.Equals(data2.Cf30) && data1.Cf31 == data2.Cf31
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
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

type D5_DC15 struct {
	Cf33 D4
	Cf34 bool
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf33 D4, Cf34 bool) D5 {
	return D5{D5_DC15{Cf33, Cf34}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC16 struct {
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_() D5 {
	return D5{D5_DC16{}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC17 struct {
	Cf35 _dafny.Int
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf35 _dafny.Int) D5 {
	return D5{D5_DC17{Cf35}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC14 struct {
	Cf32 _dafny.Array
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf32 _dafny.Array) D5 {
	return D5{D5_DC14{Cf32}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC18 struct {
	Cf36 D5
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf36 D5) D5 {
	return D5{D5_DC18{Cf36}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(Companion_D4_.Default(), false)
}

func (_this D5) Dtor_cf33() D4 {
	return _this.Get_().(D5_DC15).Cf33
}

func (_this D5) Dtor_cf34() bool {
	return _this.Get_().(D5_DC15).Cf34
}

func (_this D5) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D5_DC17).Cf35
}

func (_this D5) Dtor_cf32() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf32
}

func (_this D5) Dtor_cf36() D5 {
	return _this.Get_().(D5_DC18).Cf36
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16"
		}
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf36) + ")"
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
			return ok && data1.Cf33.Equals(data2.Cf33) && data1.Cf34 == data2.Cf34
		}
	case D5_DC16:
		{
			_, ok := other.Get_().(D5_DC16)
			return ok
		}
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf32 == data2.Cf32
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf36.Equals(data2.Cf36)
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

type D6_DC20 struct {
	Cf38 bool
	Cf39 bool
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf38 bool, Cf39 bool) D6 {
	return D6{D6_DC20{Cf38, Cf39}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

type D6_DC19 struct {
	Cf37 _dafny.Array
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf37 _dafny.Array) D6 {
	return D6{D6_DC19{Cf37}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC21 struct {
	Cf40 D6
}

func (D6_DC21) isD6() {}

func (CompanionStruct_D6_) Create_DC21_(Cf40 D6) D6 {
	return D6{D6_DC21{Cf40}}
}

func (_this D6) Is_DC21() bool {
	_, ok := _this.Get_().(D6_DC21)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC20_(false, false)
}

func (_this D6) Dtor_cf38() bool {
	return _this.Get_().(D6_DC20).Cf38
}

func (_this D6) Dtor_cf39() bool {
	return _this.Get_().(D6_DC20).Cf39
}

func (_this D6) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D6_DC19).Cf37
}

func (_this D6) Dtor_cf40() D6 {
	return _this.Get_().(D6_DC21).Cf40
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D6_DC21:
		{
			return "D6.DC21" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf38 == data2.Cf38 && data1.Cf39 == data2.Cf39
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf37 == data2.Cf37
		}
	case D6_DC21:
		{
			data2, ok := other.Get_().(D6_DC21)
			return ok && data1.Cf40.Equals(data2.Cf40)
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

type D7_DC23 struct {
	Cf42 _dafny.Int
}

func (D7_DC23) isD7() {}

func (CompanionStruct_D7_) Create_DC23_(Cf42 _dafny.Int) D7 {
	return D7{D7_DC23{Cf42}}
}

func (_this D7) Is_DC23() bool {
	_, ok := _this.Get_().(D7_DC23)
	return ok
}

type D7_DC24 struct {
	Cf43 bool
}

func (D7_DC24) isD7() {}

func (CompanionStruct_D7_) Create_DC24_(Cf43 bool) D7 {
	return D7{D7_DC24{Cf43}}
}

func (_this D7) Is_DC24() bool {
	_, ok := _this.Get_().(D7_DC24)
	return ok
}

type D7_DC22 struct {
	Cf41 _dafny.MultiSet
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf41 _dafny.MultiSet) D7 {
	return D7{D7_DC22{Cf41}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC23_(_dafny.Zero)
}

func (_this D7) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D7_DC23).Cf42
}

func (_this D7) Dtor_cf43() bool {
	return _this.Get_().(D7_DC24).Cf43
}

func (_this D7) Dtor_cf41() _dafny.MultiSet {
	return _this.Get_().(D7_DC22).Cf41
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC23:
		{
			return "D7.DC23" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D7_DC24:
		{
			return "D7.DC24" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC23:
		{
			data2, ok := other.Get_().(D7_DC23)
			return ok && data1.Cf42.Cmp(data2.Cf42) == 0
		}
	case D7_DC24:
		{
			data2, ok := other.Get_().(D7_DC24)
			return ok && data1.Cf43 == data2.Cf43
		}
	case D7_DC22:
		{
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D8_DC26 struct {
	Cf45 bool
	Cf46 bool
	Cf47 bool
	Cf48 _dafny.Int
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf45 bool, Cf46 bool, Cf47 bool, Cf48 _dafny.Int) D8 {
	return D8{D8_DC26{Cf45, Cf46, Cf47, Cf48}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

type D8_DC27 struct {
	Cf49 bool
	Cf50 bool
	Cf51 _dafny.Int
	Cf52 _dafny.Int
}

func (D8_DC27) isD8() {}

func (CompanionStruct_D8_) Create_DC27_(Cf49 bool, Cf50 bool, Cf51 _dafny.Int, Cf52 _dafny.Int) D8 {
	return D8{D8_DC27{Cf49, Cf50, Cf51, Cf52}}
}

func (_this D8) Is_DC27() bool {
	_, ok := _this.Get_().(D8_DC27)
	return ok
}

type D8_DC28 struct {
	Cf53 bool
	Cf54 _dafny.Int
}

func (D8_DC28) isD8() {}

func (CompanionStruct_D8_) Create_DC28_(Cf53 bool, Cf54 _dafny.Int) D8 {
	return D8{D8_DC28{Cf53, Cf54}}
}

func (_this D8) Is_DC28() bool {
	_, ok := _this.Get_().(D8_DC28)
	return ok
}

type D8_DC25 struct {
	Cf44 _dafny.Map
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf44 _dafny.Map) D8 {
	return D8{D8_DC25{Cf44}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC29 struct {
	Cf55 D8
}

func (D8_DC29) isD8() {}

func (CompanionStruct_D8_) Create_DC29_(Cf55 D8) D8 {
	return D8{D8_DC29{Cf55}}
}

func (_this D8) Is_DC29() bool {
	_, ok := _this.Get_().(D8_DC29)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC26_(false, false, false, _dafny.Zero)
}

func (_this D8) Dtor_cf45() bool {
	return _this.Get_().(D8_DC26).Cf45
}

func (_this D8) Dtor_cf46() bool {
	return _this.Get_().(D8_DC26).Cf46
}

func (_this D8) Dtor_cf47() bool {
	return _this.Get_().(D8_DC26).Cf47
}

func (_this D8) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D8_DC26).Cf48
}

func (_this D8) Dtor_cf49() bool {
	return _this.Get_().(D8_DC27).Cf49
}

func (_this D8) Dtor_cf50() bool {
	return _this.Get_().(D8_DC27).Cf50
}

func (_this D8) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D8_DC27).Cf51
}

func (_this D8) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D8_DC27).Cf52
}

func (_this D8) Dtor_cf53() bool {
	return _this.Get_().(D8_DC28).Cf53
}

func (_this D8) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D8_DC28).Cf54
}

func (_this D8) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D8_DC25).Cf44
}

func (_this D8) Dtor_cf55() D8 {
	return _this.Get_().(D8_DC29).Cf55
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D8_DC27:
		{
			return "D8.DC27" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D8_DC28:
		{
			return "D8.DC28" + "(" + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf44) + ")"
		}
	case D8_DC29:
		{
			return "D8.DC29" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC26:
		{
			data2, ok := other.Get_().(D8_DC26)
			return ok && data1.Cf45 == data2.Cf45 && data1.Cf46 == data2.Cf46 && data1.Cf47 == data2.Cf47 && data1.Cf48.Cmp(data2.Cf48) == 0
		}
	case D8_DC27:
		{
			data2, ok := other.Get_().(D8_DC27)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50 == data2.Cf50 && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52.Cmp(data2.Cf52) == 0
		}
	case D8_DC28:
		{
			data2, ok := other.Get_().(D8_DC28)
			return ok && data1.Cf53 == data2.Cf53 && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf44.Equals(data2.Cf44)
		}
	case D8_DC29:
		{
			data2, ok := other.Get_().(D8_DC29)
			return ok && data1.Cf55.Equals(data2.Cf55)
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

type D9_DC30 struct {
	Cf56 _dafny.Array
}

func (D9_DC30) isD9() {}

func (CompanionStruct_D9_) Create_DC30_(Cf56 _dafny.Array) D9 {
	return D9{D9_DC30{Cf56}}
}

func (_this D9) Is_DC30() bool {
	_, ok := _this.Get_().(D9_DC30)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D9) Dtor_cf56() _dafny.Array {
	return _this.Get_().(D9_DC30).Cf56
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC30:
		{
			return "D9.DC30" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC30:
		{
			data2, ok := other.Get_().(D9_DC30)
			return ok && data1.Cf56 == data2.Cf56
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

type D10_DC32 struct {
	Cf58 _dafny.MultiSet
	Cf59 _dafny.Int
	Cf60 _dafny.Int
}

func (D10_DC32) isD10() {}

func (CompanionStruct_D10_) Create_DC32_(Cf58 _dafny.MultiSet, Cf59 _dafny.Int, Cf60 _dafny.Int) D10 {
	return D10{D10_DC32{Cf58, Cf59, Cf60}}
}

func (_this D10) Is_DC32() bool {
	_, ok := _this.Get_().(D10_DC32)
	return ok
}

type D10_DC33 struct {
	Cf61 bool
}

func (D10_DC33) isD10() {}

func (CompanionStruct_D10_) Create_DC33_(Cf61 bool) D10 {
	return D10{D10_DC33{Cf61}}
}

func (_this D10) Is_DC33() bool {
	_, ok := _this.Get_().(D10_DC33)
	return ok
}

type D10_DC31 struct {
	Cf57 *C1
}

func (D10_DC31) isD10() {}

func (CompanionStruct_D10_) Create_DC31_(Cf57 *C1) D10 {
	return D10{D10_DC31{Cf57}}
}

func (_this D10) Is_DC31() bool {
	_, ok := _this.Get_().(D10_DC31)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC32_(_dafny.EmptyMultiSet, _dafny.Zero, _dafny.Zero)
}

func (_this D10) Dtor_cf58() _dafny.MultiSet {
	return _this.Get_().(D10_DC32).Cf58
}

func (_this D10) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D10_DC32).Cf59
}

func (_this D10) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D10_DC32).Cf60
}

func (_this D10) Dtor_cf61() bool {
	return _this.Get_().(D10_DC33).Cf61
}

func (_this D10) Dtor_cf57() *C1 {
	return _this.Get_().(D10_DC31).Cf57
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC32:
		{
			return "D10.DC32" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D10_DC33:
		{
			return "D10.DC33" + "(" + _dafny.String(data.Cf61) + ")"
		}
	case D10_DC31:
		{
			return "D10.DC31" + "(" + _dafny.String(data.Cf57) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC32:
		{
			data2, ok := other.Get_().(D10_DC32)
			return ok && data1.Cf58.Equals(data2.Cf58) && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60.Cmp(data2.Cf60) == 0
		}
	case D10_DC33:
		{
			data2, ok := other.Get_().(D10_DC33)
			return ok && data1.Cf61 == data2.Cf61
		}
	case D10_DC31:
		{
			data2, ok := other.Get_().(D10_DC31)
			return ok && data1.Cf57 == data2.Cf57
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

type D11_DC35 struct {
	Cf63 _dafny.Int
	Cf64 bool
	Cf65 bool
}

func (D11_DC35) isD11() {}

func (CompanionStruct_D11_) Create_DC35_(Cf63 _dafny.Int, Cf64 bool, Cf65 bool) D11 {
	return D11{D11_DC35{Cf63, Cf64, Cf65}}
}

func (_this D11) Is_DC35() bool {
	_, ok := _this.Get_().(D11_DC35)
	return ok
}

type D11_DC34 struct {
	Cf62 _dafny.Map
}

func (D11_DC34) isD11() {}

func (CompanionStruct_D11_) Create_DC34_(Cf62 _dafny.Map) D11 {
	return D11{D11_DC34{Cf62}}
}

func (_this D11) Is_DC34() bool {
	_, ok := _this.Get_().(D11_DC34)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC35_(_dafny.Zero, false, false)
}

func (_this D11) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D11_DC35).Cf63
}

func (_this D11) Dtor_cf64() bool {
	return _this.Get_().(D11_DC35).Cf64
}

func (_this D11) Dtor_cf65() bool {
	return _this.Get_().(D11_DC35).Cf65
}

func (_this D11) Dtor_cf62() _dafny.Map {
	return _this.Get_().(D11_DC34).Cf62
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC35:
		{
			return "D11.DC35" + "(" + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D11_DC34:
		{
			return "D11.DC34" + "(" + _dafny.String(data.Cf62) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC35:
		{
			data2, ok := other.Get_().(D11_DC35)
			return ok && data1.Cf63.Cmp(data2.Cf63) == 0 && data1.Cf64 == data2.Cf64 && data1.Cf65 == data2.Cf65
		}
	case D11_DC34:
		{
			data2, ok := other.Get_().(D11_DC34)
			return ok && data1.Cf62.Equals(data2.Cf62)
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
	Cf67 bool
	Cf68 _dafny.Sequence
	Cf69 _dafny.Int
}

func (D12_DC37) isD12() {}

func (CompanionStruct_D12_) Create_DC37_(Cf67 bool, Cf68 _dafny.Sequence, Cf69 _dafny.Int) D12 {
	return D12{D12_DC37{Cf67, Cf68, Cf69}}
}

func (_this D12) Is_DC37() bool {
	_, ok := _this.Get_().(D12_DC37)
	return ok
}

type D12_DC38 struct {
	Cf70 _dafny.Set
	Cf71 _dafny.Int
	Cf72 *C2
	Cf73 _dafny.Int
}

func (D12_DC38) isD12() {}

func (CompanionStruct_D12_) Create_DC38_(Cf70 _dafny.Set, Cf71 _dafny.Int, Cf72 *C2, Cf73 _dafny.Int) D12 {
	return D12{D12_DC38{Cf70, Cf71, Cf72, Cf73}}
}

func (_this D12) Is_DC38() bool {
	_, ok := _this.Get_().(D12_DC38)
	return ok
}

type D12_DC36 struct {
	Cf66 _dafny.Map
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf66 _dafny.Map) D12 {
	return D12{D12_DC36{Cf66}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

type D12_DC39 struct {
	Cf74 D12
}

func (D12_DC39) isD12() {}

func (CompanionStruct_D12_) Create_DC39_(Cf74 D12) D12 {
	return D12{D12_DC39{Cf74}}
}

func (_this D12) Is_DC39() bool {
	_, ok := _this.Get_().(D12_DC39)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC37_(false, _dafny.EmptySeq, _dafny.Zero)
}

func (_this D12) Dtor_cf67() bool {
	return _this.Get_().(D12_DC37).Cf67
}

func (_this D12) Dtor_cf68() _dafny.Sequence {
	return _this.Get_().(D12_DC37).Cf68
}

func (_this D12) Dtor_cf69() _dafny.Int {
	return _this.Get_().(D12_DC37).Cf69
}

func (_this D12) Dtor_cf70() _dafny.Set {
	return _this.Get_().(D12_DC38).Cf70
}

func (_this D12) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D12_DC38).Cf71
}

func (_this D12) Dtor_cf72() *C2 {
	return _this.Get_().(D12_DC38).Cf72
}

func (_this D12) Dtor_cf73() _dafny.Int {
	return _this.Get_().(D12_DC38).Cf73
}

func (_this D12) Dtor_cf66() _dafny.Map {
	return _this.Get_().(D12_DC36).Cf66
}

func (_this D12) Dtor_cf74() D12 {
	return _this.Get_().(D12_DC39).Cf74
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC37:
		{
			return "D12.DC37" + "(" + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D12_DC38:
		{
			return "D12.DC38" + "(" + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ")"
		}
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf66) + ")"
		}
	case D12_DC39:
		{
			return "D12.DC39" + "(" + _dafny.String(data.Cf74) + ")"
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
			data2, ok := other.Get_().(D12_DC37)
			return ok && data1.Cf67 == data2.Cf67 && data1.Cf68.Equals(data2.Cf68) && data1.Cf69.Cmp(data2.Cf69) == 0
		}
	case D12_DC38:
		{
			data2, ok := other.Get_().(D12_DC38)
			return ok && data1.Cf70.Equals(data2.Cf70) && data1.Cf71.Cmp(data2.Cf71) == 0 && data1.Cf72 == data2.Cf72 && data1.Cf73.Cmp(data2.Cf73) == 0
		}
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf66.Equals(data2.Cf66)
		}
	case D12_DC39:
		{
			data2, ok := other.Get_().(D12_DC39)
			return ok && data1.Cf74.Equals(data2.Cf74)
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

type D13_DC41 struct {
	Cf76 _dafny.Int
}

func (D13_DC41) isD13() {}

func (CompanionStruct_D13_) Create_DC41_(Cf76 _dafny.Int) D13 {
	return D13{D13_DC41{Cf76}}
}

func (_this D13) Is_DC41() bool {
	_, ok := _this.Get_().(D13_DC41)
	return ok
}

type D13_DC42 struct {
	Cf77 _dafny.Sequence
	Cf78 _dafny.MultiSet
	Cf79 _dafny.Sequence
	Cf80 bool
	Cf81 _dafny.Sequence
}

func (D13_DC42) isD13() {}

func (CompanionStruct_D13_) Create_DC42_(Cf77 _dafny.Sequence, Cf78 _dafny.MultiSet, Cf79 _dafny.Sequence, Cf80 bool, Cf81 _dafny.Sequence) D13 {
	return D13{D13_DC42{Cf77, Cf78, Cf79, Cf80, Cf81}}
}

func (_this D13) Is_DC42() bool {
	_, ok := _this.Get_().(D13_DC42)
	return ok
}

type D13_DC40 struct {
	Cf75 _dafny.MultiSet
}

func (D13_DC40) isD13() {}

func (CompanionStruct_D13_) Create_DC40_(Cf75 _dafny.MultiSet) D13 {
	return D13{D13_DC40{Cf75}}
}

func (_this D13) Is_DC40() bool {
	_, ok := _this.Get_().(D13_DC40)
	return ok
}

type D13_DC43 struct {
	Cf82 D13
}

func (D13_DC43) isD13() {}

func (CompanionStruct_D13_) Create_DC43_(Cf82 D13) D13 {
	return D13{D13_DC43{Cf82}}
}

func (_this D13) Is_DC43() bool {
	_, ok := _this.Get_().(D13_DC43)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC41_(_dafny.Zero)
}

func (_this D13) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D13_DC41).Cf76
}

func (_this D13) Dtor_cf77() _dafny.Sequence {
	return _this.Get_().(D13_DC42).Cf77
}

func (_this D13) Dtor_cf78() _dafny.MultiSet {
	return _this.Get_().(D13_DC42).Cf78
}

func (_this D13) Dtor_cf79() _dafny.Sequence {
	return _this.Get_().(D13_DC42).Cf79
}

func (_this D13) Dtor_cf80() bool {
	return _this.Get_().(D13_DC42).Cf80
}

func (_this D13) Dtor_cf81() _dafny.Sequence {
	return _this.Get_().(D13_DC42).Cf81
}

func (_this D13) Dtor_cf75() _dafny.MultiSet {
	return _this.Get_().(D13_DC40).Cf75
}

func (_this D13) Dtor_cf82() D13 {
	return _this.Get_().(D13_DC43).Cf82
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC41:
		{
			return "D13.DC41" + "(" + _dafny.String(data.Cf76) + ")"
		}
	case D13_DC42:
		{
			return "D13.DC42" + "(" + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ")"
		}
	case D13_DC40:
		{
			return "D13.DC40" + "(" + _dafny.String(data.Cf75) + ")"
		}
	case D13_DC43:
		{
			return "D13.DC43" + "(" + _dafny.String(data.Cf82) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC41:
		{
			data2, ok := other.Get_().(D13_DC41)
			return ok && data1.Cf76.Cmp(data2.Cf76) == 0
		}
	case D13_DC42:
		{
			data2, ok := other.Get_().(D13_DC42)
			return ok && data1.Cf77.Equals(data2.Cf77) && data1.Cf78.Equals(data2.Cf78) && data1.Cf79.Equals(data2.Cf79) && data1.Cf80 == data2.Cf80 && data1.Cf81.Equals(data2.Cf81)
		}
	case D13_DC40:
		{
			data2, ok := other.Get_().(D13_DC40)
			return ok && data1.Cf75.Equals(data2.Cf75)
		}
	case D13_DC43:
		{
			data2, ok := other.Get_().(D13_DC43)
			return ok && data1.Cf82.Equals(data2.Cf82)
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

type D14_DC45 struct {
	Cf84 bool
}

func (D14_DC45) isD14() {}

func (CompanionStruct_D14_) Create_DC45_(Cf84 bool) D14 {
	return D14{D14_DC45{Cf84}}
}

func (_this D14) Is_DC45() bool {
	_, ok := _this.Get_().(D14_DC45)
	return ok
}

type D14_DC44 struct {
	Cf83 _dafny.MultiSet
}

func (D14_DC44) isD14() {}

func (CompanionStruct_D14_) Create_DC44_(Cf83 _dafny.MultiSet) D14 {
	return D14{D14_DC44{Cf83}}
}

func (_this D14) Is_DC44() bool {
	_, ok := _this.Get_().(D14_DC44)
	return ok
}

type D14_DC46 struct {
	Cf85 D14
}

func (D14_DC46) isD14() {}

func (CompanionStruct_D14_) Create_DC46_(Cf85 D14) D14 {
	return D14{D14_DC46{Cf85}}
}

func (_this D14) Is_DC46() bool {
	_, ok := _this.Get_().(D14_DC46)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC45_(false)
}

func (_this D14) Dtor_cf84() bool {
	return _this.Get_().(D14_DC45).Cf84
}

func (_this D14) Dtor_cf83() _dafny.MultiSet {
	return _this.Get_().(D14_DC44).Cf83
}

func (_this D14) Dtor_cf85() D14 {
	return _this.Get_().(D14_DC46).Cf85
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC45:
		{
			return "D14.DC45" + "(" + _dafny.String(data.Cf84) + ")"
		}
	case D14_DC44:
		{
			return "D14.DC44" + "(" + _dafny.String(data.Cf83) + ")"
		}
	case D14_DC46:
		{
			return "D14.DC46" + "(" + _dafny.String(data.Cf85) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC45:
		{
			data2, ok := other.Get_().(D14_DC45)
			return ok && data1.Cf84 == data2.Cf84
		}
	case D14_DC44:
		{
			data2, ok := other.Get_().(D14_DC44)
			return ok && data1.Cf83.Equals(data2.Cf83)
		}
	case D14_DC46:
		{
			data2, ok := other.Get_().(D14_DC46)
			return ok && data1.Cf85.Equals(data2.Cf85)
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

type D15_DC48 struct {
	Cf87 bool
	Cf88 _dafny.Array
	Cf89 _dafny.Int
	Cf90 bool
	Cf91 _dafny.Sequence
}

func (D15_DC48) isD15() {}

func (CompanionStruct_D15_) Create_DC48_(Cf87 bool, Cf88 _dafny.Array, Cf89 _dafny.Int, Cf90 bool, Cf91 _dafny.Sequence) D15 {
	return D15{D15_DC48{Cf87, Cf88, Cf89, Cf90, Cf91}}
}

func (_this D15) Is_DC48() bool {
	_, ok := _this.Get_().(D15_DC48)
	return ok
}

type D15_DC47 struct {
	Cf86 _dafny.Array
}

func (D15_DC47) isD15() {}

func (CompanionStruct_D15_) Create_DC47_(Cf86 _dafny.Array) D15 {
	return D15{D15_DC47{Cf86}}
}

func (_this D15) Is_DC47() bool {
	_, ok := _this.Get_().(D15_DC47)
	return ok
}

type D15_DC49 struct {
	Cf92 D15
}

func (D15_DC49) isD15() {}

func (CompanionStruct_D15_) Create_DC49_(Cf92 D15) D15 {
	return D15{D15_DC49{Cf92}}
}

func (_this D15) Is_DC49() bool {
	_, ok := _this.Get_().(D15_DC49)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC48_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, false, _dafny.EmptySeq)
}

func (_this D15) Dtor_cf87() bool {
	return _this.Get_().(D15_DC48).Cf87
}

func (_this D15) Dtor_cf88() _dafny.Array {
	return _this.Get_().(D15_DC48).Cf88
}

func (_this D15) Dtor_cf89() _dafny.Int {
	return _this.Get_().(D15_DC48).Cf89
}

func (_this D15) Dtor_cf90() bool {
	return _this.Get_().(D15_DC48).Cf90
}

func (_this D15) Dtor_cf91() _dafny.Sequence {
	return _this.Get_().(D15_DC48).Cf91
}

func (_this D15) Dtor_cf86() _dafny.Array {
	return _this.Get_().(D15_DC47).Cf86
}

func (_this D15) Dtor_cf92() D15 {
	return _this.Get_().(D15_DC49).Cf92
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC48:
		{
			return "D15.DC48" + "(" + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ", " + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ", " + data.Cf91.VerbatimString(true) + ")"
		}
	case D15_DC47:
		{
			return "D15.DC47" + "(" + _dafny.String(data.Cf86) + ")"
		}
	case D15_DC49:
		{
			return "D15.DC49" + "(" + _dafny.String(data.Cf92) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC48:
		{
			data2, ok := other.Get_().(D15_DC48)
			return ok && data1.Cf87 == data2.Cf87 && data1.Cf88 == data2.Cf88 && data1.Cf89.Cmp(data2.Cf89) == 0 && data1.Cf90 == data2.Cf90 && data1.Cf91.Equals(data2.Cf91)
		}
	case D15_DC47:
		{
			data2, ok := other.Get_().(D15_DC47)
			return ok && data1.Cf86 == data2.Cf86
		}
	case D15_DC49:
		{
			data2, ok := other.Get_().(D15_DC49)
			return ok && data1.Cf92.Equals(data2.Cf92)
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

type D16_DC51 struct {
	Cf94 bool
	Cf95 _dafny.Set
	Cf96 bool
	Cf97 _dafny.Int
}

func (D16_DC51) isD16() {}

func (CompanionStruct_D16_) Create_DC51_(Cf94 bool, Cf95 _dafny.Set, Cf96 bool, Cf97 _dafny.Int) D16 {
	return D16{D16_DC51{Cf94, Cf95, Cf96, Cf97}}
}

func (_this D16) Is_DC51() bool {
	_, ok := _this.Get_().(D16_DC51)
	return ok
}

type D16_DC50 struct {
	Cf93 _dafny.MultiSet
}

func (D16_DC50) isD16() {}

func (CompanionStruct_D16_) Create_DC50_(Cf93 _dafny.MultiSet) D16 {
	return D16{D16_DC50{Cf93}}
}

func (_this D16) Is_DC50() bool {
	_, ok := _this.Get_().(D16_DC50)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC51_(false, _dafny.EmptySet, false, _dafny.Zero)
}

func (_this D16) Dtor_cf94() bool {
	return _this.Get_().(D16_DC51).Cf94
}

func (_this D16) Dtor_cf95() _dafny.Set {
	return _this.Get_().(D16_DC51).Cf95
}

func (_this D16) Dtor_cf96() bool {
	return _this.Get_().(D16_DC51).Cf96
}

func (_this D16) Dtor_cf97() _dafny.Int {
	return _this.Get_().(D16_DC51).Cf97
}

func (_this D16) Dtor_cf93() _dafny.MultiSet {
	return _this.Get_().(D16_DC50).Cf93
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC51:
		{
			return "D16.DC51" + "(" + _dafny.String(data.Cf94) + ", " + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ", " + _dafny.String(data.Cf97) + ")"
		}
	case D16_DC50:
		{
			return "D16.DC50" + "(" + _dafny.String(data.Cf93) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC51:
		{
			data2, ok := other.Get_().(D16_DC51)
			return ok && data1.Cf94 == data2.Cf94 && data1.Cf95.Equals(data2.Cf95) && data1.Cf96 == data2.Cf96 && data1.Cf97.Cmp(data2.Cf97) == 0
		}
	case D16_DC50:
		{
			data2, ok := other.Get_().(D16_DC50)
			return ok && data1.Cf93.Equals(data2.Cf93)
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

type D17_DC52 struct {
	Cf98 _dafny.Set
}

func (D17_DC52) isD17() {}

func (CompanionStruct_D17_) Create_DC52_(Cf98 _dafny.Set) D17 {
	return D17{D17_DC52{Cf98}}
}

func (_this D17) Is_DC52() bool {
	_, ok := _this.Get_().(D17_DC52)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D17) Dtor_cf98() _dafny.Set {
	return _this.Get_().(D17_DC52).Cf98
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC52:
		{
			return "D17.DC52" + "(" + _dafny.String(data.Cf98) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC52:
		{
			data2, ok := other.Get_().(D17_DC52)
			return ok && data1.Cf98.Equals(data2.Cf98)
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

type D18_DC54 struct {
	Cf100 bool
	Cf101 bool
	Cf102 _dafny.Int
	Cf103 _dafny.Int
}

func (D18_DC54) isD18() {}

func (CompanionStruct_D18_) Create_DC54_(Cf100 bool, Cf101 bool, Cf102 _dafny.Int, Cf103 _dafny.Int) D18 {
	return D18{D18_DC54{Cf100, Cf101, Cf102, Cf103}}
}

func (_this D18) Is_DC54() bool {
	_, ok := _this.Get_().(D18_DC54)
	return ok
}

type D18_DC55 struct {
	Cf104 _dafny.Sequence
	Cf105 bool
}

func (D18_DC55) isD18() {}

func (CompanionStruct_D18_) Create_DC55_(Cf104 _dafny.Sequence, Cf105 bool) D18 {
	return D18{D18_DC55{Cf104, Cf105}}
}

func (_this D18) Is_DC55() bool {
	_, ok := _this.Get_().(D18_DC55)
	return ok
}

type D18_DC53 struct {
	Cf99 _dafny.Map
}

func (D18_DC53) isD18() {}

func (CompanionStruct_D18_) Create_DC53_(Cf99 _dafny.Map) D18 {
	return D18{D18_DC53{Cf99}}
}

func (_this D18) Is_DC53() bool {
	_, ok := _this.Get_().(D18_DC53)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC54_(false, false, _dafny.Zero, _dafny.Zero)
}

func (_this D18) Dtor_cf100() bool {
	return _this.Get_().(D18_DC54).Cf100
}

func (_this D18) Dtor_cf101() bool {
	return _this.Get_().(D18_DC54).Cf101
}

func (_this D18) Dtor_cf102() _dafny.Int {
	return _this.Get_().(D18_DC54).Cf102
}

func (_this D18) Dtor_cf103() _dafny.Int {
	return _this.Get_().(D18_DC54).Cf103
}

func (_this D18) Dtor_cf104() _dafny.Sequence {
	return _this.Get_().(D18_DC55).Cf104
}

func (_this D18) Dtor_cf105() bool {
	return _this.Get_().(D18_DC55).Cf105
}

func (_this D18) Dtor_cf99() _dafny.Map {
	return _this.Get_().(D18_DC53).Cf99
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC54:
		{
			return "D18.DC54" + "(" + _dafny.String(data.Cf100) + ", " + _dafny.String(data.Cf101) + ", " + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ")"
		}
	case D18_DC55:
		{
			return "D18.DC55" + "(" + data.Cf104.VerbatimString(true) + ", " + _dafny.String(data.Cf105) + ")"
		}
	case D18_DC53:
		{
			return "D18.DC53" + "(" + _dafny.String(data.Cf99) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC54:
		{
			data2, ok := other.Get_().(D18_DC54)
			return ok && data1.Cf100 == data2.Cf100 && data1.Cf101 == data2.Cf101 && data1.Cf102.Cmp(data2.Cf102) == 0 && data1.Cf103.Cmp(data2.Cf103) == 0
		}
	case D18_DC55:
		{
			data2, ok := other.Get_().(D18_DC55)
			return ok && data1.Cf104.Equals(data2.Cf104) && data1.Cf105 == data2.Cf105
		}
	case D18_DC53:
		{
			data2, ok := other.Get_().(D18_DC53)
			return ok && data1.Cf99.Equals(data2.Cf99)
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

type D19_DC57 struct {
}

func (D19_DC57) isD19() {}

func (CompanionStruct_D19_) Create_DC57_() D19 {
	return D19{D19_DC57{}}
}

func (_this D19) Is_DC57() bool {
	_, ok := _this.Get_().(D19_DC57)
	return ok
}

type D19_DC58 struct {
	Cf107 bool
}

func (D19_DC58) isD19() {}

func (CompanionStruct_D19_) Create_DC58_(Cf107 bool) D19 {
	return D19{D19_DC58{Cf107}}
}

func (_this D19) Is_DC58() bool {
	_, ok := _this.Get_().(D19_DC58)
	return ok
}

type D19_DC56 struct {
	Cf106 _dafny.Array
}

func (D19_DC56) isD19() {}

func (CompanionStruct_D19_) Create_DC56_(Cf106 _dafny.Array) D19 {
	return D19{D19_DC56{Cf106}}
}

func (_this D19) Is_DC56() bool {
	_, ok := _this.Get_().(D19_DC56)
	return ok
}

type D19_DC59 struct {
	Cf108 D19
}

func (D19_DC59) isD19() {}

func (CompanionStruct_D19_) Create_DC59_(Cf108 D19) D19 {
	return D19{D19_DC59{Cf108}}
}

func (_this D19) Is_DC59() bool {
	_, ok := _this.Get_().(D19_DC59)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC57_()
}

func (_this D19) Dtor_cf107() bool {
	return _this.Get_().(D19_DC58).Cf107
}

func (_this D19) Dtor_cf106() _dafny.Array {
	return _this.Get_().(D19_DC56).Cf106
}

func (_this D19) Dtor_cf108() D19 {
	return _this.Get_().(D19_DC59).Cf108
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC57:
		{
			return "D19.DC57"
		}
	case D19_DC58:
		{
			return "D19.DC58" + "(" + _dafny.String(data.Cf107) + ")"
		}
	case D19_DC56:
		{
			return "D19.DC56" + "(" + _dafny.String(data.Cf106) + ")"
		}
	case D19_DC59:
		{
			return "D19.DC59" + "(" + _dafny.String(data.Cf108) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC57:
		{
			_, ok := other.Get_().(D19_DC57)
			return ok
		}
	case D19_DC58:
		{
			data2, ok := other.Get_().(D19_DC58)
			return ok && data1.Cf107 == data2.Cf107
		}
	case D19_DC56:
		{
			data2, ok := other.Get_().(D19_DC56)
			return ok && data1.Cf106 == data2.Cf106
		}
	case D19_DC59:
		{
			data2, ok := other.Get_().(D19_DC59)
			return ok && data1.Cf108.Equals(data2.Cf108)
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

type D20_DC61 struct {
	Cf110 bool
	Cf111 bool
}

func (D20_DC61) isD20() {}

func (CompanionStruct_D20_) Create_DC61_(Cf110 bool, Cf111 bool) D20 {
	return D20{D20_DC61{Cf110, Cf111}}
}

func (_this D20) Is_DC61() bool {
	_, ok := _this.Get_().(D20_DC61)
	return ok
}

type D20_DC62 struct {
	Cf112 _dafny.Int
}

func (D20_DC62) isD20() {}

func (CompanionStruct_D20_) Create_DC62_(Cf112 _dafny.Int) D20 {
	return D20{D20_DC62{Cf112}}
}

func (_this D20) Is_DC62() bool {
	_, ok := _this.Get_().(D20_DC62)
	return ok
}

type D20_DC60 struct {
	Cf109 _dafny.Set
}

func (D20_DC60) isD20() {}

func (CompanionStruct_D20_) Create_DC60_(Cf109 _dafny.Set) D20 {
	return D20{D20_DC60{Cf109}}
}

func (_this D20) Is_DC60() bool {
	_, ok := _this.Get_().(D20_DC60)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC61_(false, false)
}

func (_this D20) Dtor_cf110() bool {
	return _this.Get_().(D20_DC61).Cf110
}

func (_this D20) Dtor_cf111() bool {
	return _this.Get_().(D20_DC61).Cf111
}

func (_this D20) Dtor_cf112() _dafny.Int {
	return _this.Get_().(D20_DC62).Cf112
}

func (_this D20) Dtor_cf109() _dafny.Set {
	return _this.Get_().(D20_DC60).Cf109
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC61:
		{
			return "D20.DC61" + "(" + _dafny.String(data.Cf110) + ", " + _dafny.String(data.Cf111) + ")"
		}
	case D20_DC62:
		{
			return "D20.DC62" + "(" + _dafny.String(data.Cf112) + ")"
		}
	case D20_DC60:
		{
			return "D20.DC60" + "(" + _dafny.String(data.Cf109) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC61:
		{
			data2, ok := other.Get_().(D20_DC61)
			return ok && data1.Cf110 == data2.Cf110 && data1.Cf111 == data2.Cf111
		}
	case D20_DC62:
		{
			data2, ok := other.Get_().(D20_DC62)
			return ok && data1.Cf112.Cmp(data2.Cf112) == 0
		}
	case D20_DC60:
		{
			data2, ok := other.Get_().(D20_DC60)
			return ok && data1.Cf109.Equals(data2.Cf109)
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
	F2() _dafny.Array
	F2_set_(value _dafny.Array)
	Fm1(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map
	Fm2(globalState *GlobalState) _dafny.Int
	M0(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Map, globalState *GlobalState) bool
	M1(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Array, p3 _dafny.Array, globalState *GlobalState) (_dafny.MultiSet, bool)
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
	Fm3(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int
	F3() _dafny.Array
	F4() bool
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
	F1  _dafny.Array
	_f0 _dafny.Set
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = _dafny.EmptySet
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

func (_this *GlobalState) Ctor__(f0 _dafny.Set, f1 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
	}
}
func (_this *GlobalState) F0() _dafny.Set {
	{
		return _this._f0
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
func (_this *C0) Fm7(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(202), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.CodePoint('k'))).Cardinality())).Cardinality(), false)).Cardinality())).Cardinality()))).IsProperSubsetOf(_dafny.MultiSetOf(func() _dafny.Map {
			var _coll33 = _dafny.NewMapBuilder()
			_ = _coll33
			for _iter35 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.MultiSetOf(!(!(true)))).Cardinality()), _dafny.IntOfInt64(107))).Elements()); ; {
				_compr_33, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _651_v0 _dafny.Int
				_651_v0 = interface{}(_compr_33).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.MultiSetOf(!(!(true)))).Cardinality()), _dafny.IntOfInt64(107)), _651_v0) {
					_coll33.Add((_651_v0).Minus(_dafny.IntOfInt64(429)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.MultiSetOf(false)).Cardinality())).Cardinality())
				}
			}
			return _coll33.ToMap()
		}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(130), (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-100))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_652_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(602)
		})))).Cardinality()))))
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f2 _dafny.Array
	_f3 _dafny.Array
	_f4 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = false
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

func (_this *C1) F2() _dafny.Array {
	return _this._f2
}
func (_this *C1) F2_set_(value _dafny.Array) {
	_this._f2 = value
}
func (_this *C1) F3() _dafny.Array {
	return _this._f3
}
func (_this *C1) F4() bool {
	return _this._f4
}
func (_this *C1) Ctor__(f3 _dafny.Array, f4 bool, f2 _dafny.Array) {
	{
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f2 = f2
	}
}
func (_this *C1) Fm3(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(163), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())).Cardinality()))).Cardinality()))
	}
}
func (_this *C1) Fm1(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			if (_this).F4() {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(373), (_this).F4())
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(928), (_this).F4())
		})()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(956), (_this).F4())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(379), (_this).F4())))
	}
}
func (_this *C1) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (((_dafny.SetOf(false, (_this).F4())).Cardinality()).Times(_dafny.IntOfInt64(-718))).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(234))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_653_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		})), _dafny.UnicodeSeqOfUtf8Bytes("vfbddobcp"))).Cardinality()))
	}
}
func (_this *C1) Fm16(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		var _source11 D2 = (func() D2 {
			if (_this).F4() {
				return Companion_D2_.Create_DC7_(_dafny.IntOfInt64(921), _dafny.IntOfInt64(443), _dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(32), false)).Cardinality(), _dafny.IntOfInt64(429)), false)).Cardinality(), (func() _dafny.Map {
					var _coll34 = _dafny.NewMapBuilder()
					_ = _coll34
					for _iter36 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('p'))).Elements()); ; {
						_compr_34, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _654_v0 _dafny.CodePoint
						_654_v0 = interface{}(_compr_34).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('p')), _654_v0) {
							_coll34.Add(_654_v0, (_this).F4())
						}
					}
					return _coll34.ToMap()
				}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())).Cardinality(), _dafny.IntOfInt64(363))).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(586), _dafny.IntOfInt64(972), _dafny.IntOfInt64(-986))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-871))).Cardinality())), _dafny.IntOfInt64(450), _dafny.IntOfInt64(181))
			}
			return Companion_D2_.Create_DC7_(_dafny.IntOfInt64(268), _dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter37 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _dafny.CodePoint('q'))).Cardinality())).Elements()); ; {
					_compr_35, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _655_v1 _dafny.Int
					_655_v1 = interface{}(_compr_35).(_dafny.Int)
					if (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _dafny.CodePoint('q'))).Cardinality())).Contains(_655_v1) {
						_coll35.Add(Companion_Default___.SafeDivisionInt(_655_v1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(397))).Cardinality())), (_this).F4())
					}
				}
				return _coll35.ToMap()
			}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true)).Cardinality(), (_this).F4()))).Cardinality()), _dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-452))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}(func(_656_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(309)
			}))).Cardinality()))).Cardinality(), (Companion_D0_.Create_DC1_((_this).F4(), true, _dafny.IntOfInt64(247))).Dtor_cf3(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F4(), (_this).F4())).Cardinality()))), _dafny.IntOfInt64(417), _dafny.IntOfUint32((_dafny.SeqOf((_this).F4(), (_this).F4(), (_this).F4(), (_this).F4(), (_this).F4())).Cardinality()))
		})()
		_ = _source11
		if _source11.Is_DC6() {
			var _657___mcc_h0 _dafny.CodePoint = _source11.Get_().(D2_DC6).Cf14
			_ = _657___mcc_h0
			var _658___mcc_h1 _dafny.Int = _source11.Get_().(D2_DC6).Cf15
			_ = _658___mcc_h1
			var _659_cf15 _dafny.Int = _658___mcc_h1
			_ = _659_cf15
			var _660_cf14 _dafny.CodePoint = _657___mcc_h0
			_ = _660_cf14
			return _659_cf15
		} else if _source11.Is_DC7() {
			var _661___mcc_h2 _dafny.Int = _source11.Get_().(D2_DC7).Cf16
			_ = _661___mcc_h2
			var _662___mcc_h3 _dafny.Int = _source11.Get_().(D2_DC7).Cf17
			_ = _662___mcc_h3
			var _663___mcc_h4 _dafny.MultiSet = _source11.Get_().(D2_DC7).Cf18
			_ = _663___mcc_h4
			var _664___mcc_h5 _dafny.Int = _source11.Get_().(D2_DC7).Cf19
			_ = _664___mcc_h5
			var _665___mcc_h6 _dafny.Int = _source11.Get_().(D2_DC7).Cf20
			_ = _665___mcc_h6
			var _666_cf20 _dafny.Int = _665___mcc_h6
			_ = _666_cf20
			var _667_cf19 _dafny.Int = _664___mcc_h5
			_ = _667_cf19
			var _668_cf18 _dafny.MultiSet = _663___mcc_h4
			_ = _668_cf18
			var _669_cf17 _dafny.Int = _662___mcc_h3
			_ = _669_cf17
			var _670_cf16 _dafny.Int = _661___mcc_h2
			_ = _670_cf16
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F4(), (_this).F4()), _dafny.SeqOf((_this).F4(), (_this).F4()))).Cardinality())
		} else if _source11.Is_DC5() {
			var _671___mcc_h7 _dafny.Sequence = _source11.Get_().(D2_DC5).Cf13
			_ = _671___mcc_h7
			var _672_cf13 _dafny.Sequence = _671___mcc_h7
			_ = _672_cf13
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_672_cf13, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(903))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}(func(_673_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})))).Cardinality())
		} else {
			var _674___mcc_h8 D2 = _source11.Get_().(D2_DC8).Cf21
			_ = _674___mcc_h8
			var _675_cf21 D2 = _674___mcc_h8
			_ = _675_cf21
			return (_dafny.IntOfInt64(465)).Minus(_dafny.IntOfInt64(536))
		}
	}
}
func (_this *C1) M0(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Map, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _676_v0 _dafny.Array
		_ = _676_v0
		var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
		_ = _nw98
		_676_v0 = _nw98
		var _677_v1 _dafny.Map
		_ = _677_v1
		_677_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_676_v0, (_this).F4())
		var _678_v2 _dafny.Sequence
		_ = _678_v2
		_678_v2 = _dafny.SeqOf(_676_v0, _676_v0)
		_677_v1 = (_677_v1).Update((_678_v2).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_678_v2).Cardinality()))).Uint32()).(_dafny.Array), (_this).F4())
		var _679_v3 _dafny.Array
		_ = _679_v3
		var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw99
		_679_v3 = _nw99
		_679_v3 = _679_v3
		var _680_v4 *C0
		_ = _680_v4
		var _nw100 *C0 = New_C0_()
		_ = _nw100
		_nw100.Ctor__()
		_680_v4 = _nw100
		_680_v4 = _680_v4
		var _681_v5 _dafny.Int
		_ = _681_v5
		_681_v5 = _dafny.IntOfInt64(146)
		_681_v5 = Companion_Default___.Fm0(_681_v5, globalState)
		var _hi5 _dafny.Int = p0
		_ = _hi5
		for _682_i0 := Companion_Default___.Fm0(_dafny.IntOfInt64(581), globalState); _682_i0.Cmp(_hi5) < 0; _682_i0 = _682_i0.Plus(_dafny.One) {
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_679_v3), 0))
			_ = _index74
			(_679_v3).ArraySet1(p0, (_index74).Int())
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_679_v3), 0))
			_ = _index75
			(_679_v3).ArraySet1((_dafny.Zero).Minus(_681_v5), (_index75).Int())
			var _683_v6 _dafny.Map
			_ = _683_v6
			_683_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
			var _684_v7 _dafny.Map
			_ = _684_v7
			_684_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _683_v6)
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_679_v3), 0))
			_ = _index76
			(_679_v3).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if (_684_v7).Contains((_this).F4()) {
					return (_684_v7).Get((_this).F4()).(_dafny.Map)
				}
				return _683_v6
			})(), (_this).F4())).Cardinality()).Times(Companion_Default___.Fm0(_dafny.IntOfInt64(891), globalState)), (_index76).Int())
			var _685_v8 _dafny.Map
			_ = _685_v8
			_685_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_682_i0, (_this).F4())
			var _686_v9 _dafny.Sequence
			_ = _686_v9
			_686_v9 = _dafny.SeqOf(Companion_D1_.Create_DC4_(_685_v8, (_680_v4).Fm7(true, _682_i0, (_this).F4(), _dafny.IntOfInt64(843), globalState), _682_i0, (_this).F4()))
			_686_v9 = _686_v9
			var _687_v10 D1
			_ = _687_v10
			_687_v10 = Companion_D1_.Create_DC4_((_this).Fm1((_this).F4(), (_this).F4(), globalState), !((_this).F4()), p0, (_this).F4())
			var _688_v11 _dafny.Set
			_ = _688_v11
			_688_v11 = _dafny.SetOf(p0, _681_v5)
			var _689_v12 _dafny.Set
			_ = _689_v12
			_689_v12 = _dafny.SetOf(_688_v11)
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_679_v3), 0))
			_ = _index77
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_679_v3), 0))
			_ = _index78
			var _rhs57 _dafny.Int = _681_v5
			_ = _rhs57
			var _rhs58 bool = (_this).F4()
			_ = _rhs58
			var _rhs59 _dafny.Int = (_687_v10).Dtor_cf11()
			_ = _rhs59
			var _rhs60 _dafny.Int = (p0).Plus((_dafny.Zero).Minus((_dafny.Zero).Minus((_679_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_679_v3), 0))).Int()).(_dafny.Int))))
			_ = _rhs60
			var _rhs61 _dafny.Int = (_689_v12).Cardinality()
			_ = _rhs61
			var _lhs28 _dafny.Array = _679_v3
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_679_v3), 0))
			_ = _lhs29
			var _lhs30 _dafny.Array = _679_v3
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_679_v3), 0))
			_ = _lhs31
			_681_v5 = _rhs57
			r0 = _rhs58
			_681_v5 = _rhs59
			(_lhs28).ArraySet1(_rhs60, (_lhs29).Int())
			(_lhs30).ArraySet1(_rhs61, (_lhs31).Int())
		}
		for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_this.F2()), 0))); ; {
			_guard_loop_2, _ok38 := _iter38()
			if !_ok38 {
				break
			}
			var _690_i1 _dafny.Int
			_690_i1 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_690_i1).Sign() != -1) && ((_690_i1).Cmp(_dafny.ArrayLen((_this.F2()), 0)) < 0)) {
				var _arr0 _dafny.Array = _this.F2()
				_ = _arr0
				(_arr0).ArraySet1((_this).F4(), (_690_i1).Int())
			}
		}
		r0 = (_this).F4()
		return r0
	}
}
func (_this *C1) M1(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Array, p3 _dafny.Array, globalState *GlobalState) (_dafny.MultiSet, bool) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 bool = false
		_ = r1
		for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_this.F2()), 0))); ; {
			_guard_loop_3, _ok39 := _iter39()
			if !_ok39 {
				break
			}
			var _691_i0 _dafny.Int
			_691_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_691_i0).Sign() != -1) && ((_691_i0).Cmp(_dafny.ArrayLen((_this.F2()), 0)) < 0)) {
				var _arr1 _dafny.Array = _this.F2()
				_ = _arr1
				(_arr1).ArraySet1(!(!(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("meicjlcv"), _dafny.UnicodeSeqOfUtf8Bytes("yewbcbs")), _dafny.UnicodeSeqOfUtf8Bytes("guyoe")))), (_691_i0).Int())
			}
		}
		var _692_v0 _dafny.Sequence
		_ = _692_v0
		_692_v0 = _dafny.SeqOf((_this).F4(), false, (_this).F4())
		var _693_v1 _dafny.Sequence
		_ = _693_v1
		_693_v1 = _dafny.SeqOf(_dafny.SeqOf((_this).F4()), _692_v0)
		var _694_v2 D0
		_ = _694_v2
		_694_v2 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(154), (_this).F4(), !_dafny.Companion_Sequence_.Contains(_693_v1, _692_v0), (_this).Fm2(globalState))
		var _695_v3 _dafny.Set
		_ = _695_v3
		_695_v3 = _dafny.SetOf(p0)
		var _696_v4 _dafny.Sequence
		_ = _696_v4
		_696_v4 = _dafny.UnicodeSeqOfUtf8Bytes("b")
		var _pat_let_tv6 = p0
		_ = _pat_let_tv6
		var _pat_let_tv7 = _694_v2
		_ = _pat_let_tv7
		var _pat_let_tv8 = _694_v2
		_ = _pat_let_tv8
		_694_v2 = func(_source12 D0) D0 {
			if _source12.Is_DC1() {
				var _697___mcc_h0 bool = _source12.Get_().(D0_DC1).Cf1
				_ = _697___mcc_h0
				var _698___mcc_h1 bool = _source12.Get_().(D0_DC1).Cf2
				_ = _698___mcc_h1
				var _699___mcc_h2 _dafny.Int = _source12.Get_().(D0_DC1).Cf3
				_ = _699___mcc_h2
				var _700_cf3 _dafny.Int = _699___mcc_h2
				_ = _700_cf3
				var _701_cf2 bool = _698___mcc_h1
				_ = _701_cf2
				var _702_cf1 bool = _697___mcc_h0
				_ = _702_cf1
				return Companion_D0_.Create_DC2_(_pat_let_tv6, !((_this).F4()), _702_cf1, _700_cf3)
			} else if _source12.Is_DC2() {
				var _703___mcc_h3 _dafny.Int = _source12.Get_().(D0_DC2).Cf4
				_ = _703___mcc_h3
				var _704___mcc_h4 bool = _source12.Get_().(D0_DC2).Cf5
				_ = _704___mcc_h4
				var _705___mcc_h5 bool = _source12.Get_().(D0_DC2).Cf6
				_ = _705___mcc_h5
				var _706___mcc_h6 _dafny.Int = _source12.Get_().(D0_DC2).Cf7
				_ = _706___mcc_h6
				var _707_cf7 _dafny.Int = _706___mcc_h6
				_ = _707_cf7
				var _708_cf6 bool = _705___mcc_h5
				_ = _708_cf6
				var _709_cf5 bool = _704___mcc_h4
				_ = _709_cf5
				var _710_cf4 _dafny.Int = _703___mcc_h3
				_ = _710_cf4
				return _pat_let_tv7
			} else {
				var _711___mcc_h7 bool = _source12.Get_().(D0_DC0).Cf0
				_ = _711___mcc_h7
				var _712_cf0 bool = _711___mcc_h7
				_ = _712_cf0
				return _pat_let_tv8
			}
		}(Companion_Default___.Fm17(_695_v3, _696_v4, globalState))
		var _713_v5 _dafny.Array
		_ = _713_v5
		var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw101
		_713_v5 = _nw101
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
		_ = _index79
		(_713_v5).ArraySet1(p0, (_index79).Int())
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
		_ = _index80
		(_713_v5).ArraySet1(p0, (_index80).Int())
		if ((_this).F4()) || ((_this).F4()) {
			var _714_v6 *C0
			_ = _714_v6
			var _nw102 *C0 = New_C0_()
			_ = _nw102
			_nw102.Ctor__()
			_714_v6 = _nw102
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
			_ = _index81
			(_713_v5).ArraySet1(p0, (_index81).Int())
			var _715_v7 D0
			_ = _715_v7
			_715_v7 = Companion_D0_.Create_DC1_((_this).F4(), (_692_v0).Select((Companion_Default___.SafeIndex((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_692_v0).Cardinality()))).Uint32()).(bool), (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int))
			if (p0).Cmp((_715_v7).Dtor_cf3()) == 0 {
				r1 = (p0).Cmp(_dafny.IntOfUint32((_692_v0).Cardinality())) != 0
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index82
				(_713_v5).ArraySet1(p0, (_index82).Int())
				var _716_v8 _dafny.Set
				_ = _716_v8
				_716_v8 = _dafny.SetOf((_this).F4(), (_this).F4(), (_this).F4())
				_716_v8 = _716_v8
				var _arr2 _dafny.Array = _this.F2()
				_ = _arr2
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index83
				(_arr2).ArraySet1((true) == ((_this).F4()), (_index83).Int())
				var _717_v9 _dafny.Array
				_ = _717_v9
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw103
				_717_v9 = _nw103
				var _arr3 _dafny.Array = _this.F2()
				_ = _arr3
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index84
				var _rhs62 bool = (func() bool {
					if (_dafny.IntOfInt64(-366)).Cmp(p0) >= 0 {
						return (_this).F4()
					}
					return (_this).F4()
				})()
				_ = _rhs62
				var _rhs63 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_696_v4, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_696_v4, _696_v4), (Companion_Default___.SafeIndex((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_696_v4, _696_v4)).Cardinality()))).Uint32(), p1))
				_ = _rhs63
				var _rhs64 _dafny.Array = _717_v9
				_ = _rhs64
				var _rhs65 bool = (_this).F4()
				_ = _rhs65
				var _lhs32 _dafny.Array = _this.F2()
				_ = _lhs32
				var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs33
				(_lhs32).ArraySet1(_rhs62, (_lhs33).Int())
				_696_v4 = _rhs63
				_717_v9 = _rhs64
				r1 = _rhs65
				var _718_v10 _dafny.Map
				_ = _718_v10
				_718_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int))
				var _719_v11 _dafny.Map
				_ = _719_v11
				_719_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F4())
				var _720_v12 _dafny.Array
				_ = _720_v12
				var _nwElement0_19 bool = (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)
				_ = _nwElement0_19
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(24))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_19, 0)
				(_nw104).ArraySet1(!((_718_v10).Equals(_718_v10)), 1)
				(_nw104).ArraySet1((_this).F4(), 2)
				(_nw104).ArraySet1(_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm18((func() bool {
					if (_719_v11).Contains((_716_v8).Cardinality()) {
						return (_719_v11).Get((_716_v8).Cardinality()).(bool)
					}
					return (_this).F4()
				})(), _719_v11, globalState), _dafny.UnicodeSeqOfUtf8Bytes("wsxhj")), 3)
				(_nw104).ArraySet1((_this).F4(), 4)
				(_nw104).ArraySet1((_this).F4(), 5)
				(_nw104).ArraySet1((_this).F4(), 6)
				(_nw104).ArraySet1((_this).F4(), 7)
				(_nw104).ArraySet1((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), 8)
				(_nw104).ArraySet1((_this).F4(), 9)
				(_nw104).ArraySet1((_this).F4(), 10)
				(_nw104).ArraySet1((_714_v6).Fm7(false, _dafny.IntOfInt64(224), (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), _dafny.IntOfInt64(-510), globalState), 11)
				(_nw104).ArraySet1(!((_this).F4()), 12)
				(_nw104).ArraySet1(((_this).Fm3((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), !((_this).F4()), (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), p0, globalState)).Cmp((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)) >= 0, 13)
				(_nw104).ArraySet1(true, 14)
				(_nw104).ArraySet1((_this).F4(), 15)
				(_nw104).ArraySet1((_this).F4(), 16)
				(_nw104).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_696_v4, _dafny.SeqOf(p1, p1, p1)), 17)
				(_nw104).ArraySet1((_this).F4(), 18)
				(_nw104).ArraySet1(false, 19)
				(_nw104).ArraySet1(!(true), 20)
				(_nw104).ArraySet1((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), 21)
				(_nw104).ArraySet1((_this).F4(), 22)
				(_nw104).ArraySet1((_this).F4(), 23)
				_720_v12 = _nw104
				(globalState).F1 = _720_v12
			} else {
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index85
				(_713_v5).ArraySet1((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index86
				(_713_v5).ArraySet1((_dafny.Zero).Minus((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)), (_index86).Int())
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index87
				(_713_v5).ArraySet1(p0, (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index88
				(_713_v5).ArraySet1((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), (_index88).Int())
				var _721_v13 *C0
				_ = _721_v13
				var _nw105 *C0 = New_C0_()
				_ = _nw105
				_nw105.Ctor__()
				_721_v13 = _nw105
			}
			var _arr4 _dafny.Array = _this.F2()
			_ = _arr4
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index89
			(_arr4).ArraySet1(((_dafny.Zero).Minus(p0)).Cmp((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)) < 0, (_index89).Int())
			var _722_v14 _dafny.CodePoint
			_ = _722_v14
			_722_v14 = _dafny.CodePoint('k')
			var _723_v15 _dafny.Sequence
			_ = _723_v15
			_723_v15 = _dafny.SeqOf((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int))
			var _724_v16 _dafny.MultiSet
			_ = _724_v16
			_724_v16 = _dafny.MultiSetOf(p0, _dafny.IntOfUint32((_692_v0).Cardinality()), _dafny.IntOfInt64(899))
			var _725_v17 _dafny.Map
			_ = _725_v17
			_725_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), _dafny.SeqOf((_this).F4()))
			var _726_v19 _dafny.MultiSet
			_ = _726_v19
			_726_v19 = _dafny.MultiSetOf(_713_v5)
			var _727_v20 _dafny.Map
			_ = _727_v20
			_727_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Update(_723_v15, (Companion_Default___.SafeIndex((_726_v19).Cardinality(), _dafny.IntOfUint32((_723_v15).Cardinality()))).Uint32(), (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)))
			var _arr5 _dafny.Array = _this.F2()
			_ = _arr5
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index90
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
			_ = _index91
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
			_ = _index92
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
			_ = _index93
			var _rhs66 bool = !(_dafny.MultiSetFromSeq(_723_v15)).Equals(_724_v16)
			_ = _rhs66
			var _rhs67 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_696_v4, _696_v4), _696_v4)).Cardinality())
			_ = _rhs67
			var _rhs68 _dafny.Int = (((_725_v17).Merge(func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter40 := _dafny.Iterate(((func() _dafny.Sequence {
					if (_727_v20).Contains(p0) {
						return (_727_v20).Get(p0).(_dafny.Sequence)
					}
					return _dafny.SeqOf(_dafny.IntOfInt64(713))
				})()).Elements()); ; {
					_compr_36, _ok40 := _iter40()
					if !_ok40 {
						break
					}
					var _728_v18 _dafny.Int
					_728_v18 = interface{}(_compr_36).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
						if (_727_v20).Contains(p0) {
							return (_727_v20).Get(p0).(_dafny.Sequence)
						}
						return _dafny.SeqOf(_dafny.IntOfInt64(713))
					})(), _728_v18) {
						_coll36.Add(Companion_Default___.SafeDivisionInt(_728_v18, _dafny.IntOfInt64(72)), _692_v0)
					}
				}
				return _coll36.ToMap()
			}())).Cardinality()).Minus(Companion_Default___.SafeModuloInt((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)))
			_ = _rhs68
			var _rhs69 _dafny.Int = Companion_Default___.SafeDivisionInt((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), (_dafny.IntOfInt64(532)).Times((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)))
			_ = _rhs69
			var _rhs70 _dafny.CodePoint = p1
			_ = _rhs70
			var _lhs34 _dafny.Array = _this.F2()
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_this.F2()), 0))
			_ = _lhs35
			var _lhs36 _dafny.Array = _713_v5
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
			_ = _lhs37
			var _lhs38 _dafny.Array = _713_v5
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
			_ = _lhs39
			var _lhs40 _dafny.Array = _713_v5
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
			_ = _lhs41
			(_lhs34).ArraySet1(_rhs66, (_lhs35).Int())
			(_lhs36).ArraySet1(_rhs67, (_lhs37).Int())
			(_lhs38).ArraySet1(_rhs68, (_lhs39).Int())
			(_lhs40).ArraySet1(_rhs69, (_lhs41).Int())
			_722_v14 = _rhs70
			_692_v0 = (func() _dafny.Sequence {
				if (_this).F4() {
					return _692_v0
				}
				return _dafny.Companion_Sequence_.Concatenate((Companion_Default___.Fm19((_this).F4(), p0, globalState)).Dtor_cf28(), _692_v0)
			})()
		} else {
			var _729_v21 _dafny.Map
			_ = _729_v21
			_729_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _this.F2())
			(_this).F2_set_((func() _dafny.Array {
				if (_729_v21).Contains((_this).F4()) {
					return (_729_v21).Get((_this).F4()).(_dafny.Array)
				}
				return _this.F2()
			})())
			var _730_v22 _dafny.Sequence
			_ = _730_v22
			_730_v22 = _dafny.SeqOf(_696_v4, _696_v4, _696_v4)
			var _731_v23 _dafny.MultiSet
			_ = _731_v23
			_731_v23 = _dafny.MultiSetOf(p0, _dafny.IntOfUint32((_696_v4).Cardinality()), _dafny.IntOfInt64(896), p0)
			var _732_v24 _dafny.Map
			_ = _732_v24
			_732_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), false)
			var _733_v25 _dafny.Map
			_ = _733_v25
			_733_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_732_v24).Cardinality(), true)
			var _734_v26 _dafny.Map
			_ = _734_v26
			_734_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_730_v22).Cardinality()), Companion_D2_.Create_DC7_(p0, (_dafny.Zero).Minus((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)), _dafny.MultiSetOf(_731_v23, _731_v23), ((_733_v25).Update(_dafny.IntOfUint32((_696_v4).Cardinality()), false)).Cardinality(), p0))
			var _735_v27 D2
			_ = _735_v27
			_735_v27 = Companion_D2_.Create_DC8_((func() D2 {
				if (_734_v26).Contains((_dafny.Zero).Minus(p0)) {
					return (_734_v26).Get((_dafny.Zero).Minus(p0)).(D2)
				}
				return Companion_D2_.Create_DC5_(_dafny.SeqOf(p1))
			})())
			var _source13 D2 = _735_v27
			_ = _source13
			if _source13.Is_DC6() {
				var _736___mcc_h8 _dafny.CodePoint = _source13.Get_().(D2_DC6).Cf14
				_ = _736___mcc_h8
				var _737___mcc_h9 _dafny.Int = _source13.Get_().(D2_DC6).Cf15
				_ = _737___mcc_h9
				var _738_cf15 _dafny.Int = _737___mcc_h9
				_ = _738_cf15
				var _739_cf14 _dafny.CodePoint = _736___mcc_h8
				_ = _739_cf14
				var _740_v28 _dafny.Map
				_ = _740_v28
				_740_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), _738_cf15)
				var _741_v29 _dafny.Set
				_ = _741_v29
				_741_v29 = _dafny.SetOf(_dafny.CodePoint('t'), p1)
				var _742_v32 D1
				_ = _742_v32
				_742_v32 = Companion_D1_.Create_DC4_(func() _dafny.Map {
					var _coll37 = _dafny.NewMapBuilder()
					_ = _coll37
					for _iter41 := _dafny.Iterate((func() _dafny.Map {
						var _coll38 = _dafny.NewMapBuilder()
						_ = _coll38
						for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-366), _dafny.IntOfInt64(537))); ; {
							_compr_38, _ok42 := _iter42()
							if !_ok42 {
								break
							}
							var _743_v31 _dafny.Int
							_743_v31 = interface{}(_compr_38).(_dafny.Int)
							if ((_dafny.IntOfInt64(-366)).Cmp(_743_v31) <= 0) && ((_743_v31).Cmp(_dafny.IntOfInt64(537)) < 0) {
								_coll38.Add((_743_v31).Plus(_dafny.IntOfInt64(814)), _696_v4)
							}
						}
						return _coll38.ToMap()
					}()).Keys().Elements()); ; {
						_compr_37, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _744_v30 _dafny.Int
						_744_v30 = interface{}(_compr_37).(_dafny.Int)
						if (func() _dafny.Map {
							var _coll39 = _dafny.NewMapBuilder()
							_ = _coll39
							for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-366), _dafny.IntOfInt64(537))); ; {
								_compr_39, _ok43 := _iter43()
								if !_ok43 {
									break
								}
								var _743_v31 _dafny.Int
								_743_v31 = interface{}(_compr_39).(_dafny.Int)
								if ((_dafny.IntOfInt64(-366)).Cmp(_743_v31) <= 0) && ((_743_v31).Cmp(_dafny.IntOfInt64(537)) < 0) {
									_coll39.Add((_743_v31).Plus(_dafny.IntOfInt64(814)), _696_v4)
								}
							}
							return _coll39.ToMap()
						}()).Contains(_744_v30) {
							_coll37.Add(Companion_Default___.SafeModuloInt(_744_v30, (_dafny.SetOf((_this).F4(), (_this).F4())).Cardinality()), (_this).F4())
						}
					}
					return _coll37.ToMap()
				}(), (_this).F4(), p0, (_this).F4())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index94
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index95
				var _rhs71 _dafny.Int = (Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_740_v28).Cardinality()), p0)).Minus((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int))
				_ = _rhs71
				var _rhs72 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _741_v29)).Cardinality()).Plus((_742_v32).Dtor_cf11())
				_ = _rhs72
				var _rhs73 bool = (_this).F4()
				_ = _rhs73
				var _rhs74 bool = (_this).F4()
				_ = _rhs74
				var _lhs42 _dafny.Array = _713_v5
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _lhs43
				var _lhs44 _dafny.Array = _713_v5
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _lhs45
				(_lhs42).ArraySet1(_rhs71, (_lhs43).Int())
				(_lhs44).ArraySet1(_rhs72, (_lhs45).Int())
				r1 = _rhs73
				r1 = _rhs74
				_738_cf15 = (func() _dafny.Int {
					if (_731_v23).Contains(_dafny.IntOfInt64(71)) {
						return (_731_v23).Multiplicity(_dafny.IntOfInt64(71))
					}
					return (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)
				})()
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index96
				(_713_v5).ArraySet1(_738_cf15, (_index96).Int())
				_739_cf14 = (_696_v4).Select((Companion_Default___.SafeIndex((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_696_v4).Cardinality()))).Uint32()).(_dafny.CodePoint)
			} else if _source13.Is_DC7() {
				var _745___mcc_h10 _dafny.Int = _source13.Get_().(D2_DC7).Cf16
				_ = _745___mcc_h10
				var _746___mcc_h11 _dafny.Int = _source13.Get_().(D2_DC7).Cf17
				_ = _746___mcc_h11
				var _747___mcc_h12 _dafny.MultiSet = _source13.Get_().(D2_DC7).Cf18
				_ = _747___mcc_h12
				var _748___mcc_h13 _dafny.Int = _source13.Get_().(D2_DC7).Cf19
				_ = _748___mcc_h13
				var _749___mcc_h14 _dafny.Int = _source13.Get_().(D2_DC7).Cf20
				_ = _749___mcc_h14
				var _750_cf20 _dafny.Int = _749___mcc_h14
				_ = _750_cf20
				var _751_cf19 _dafny.Int = _748___mcc_h13
				_ = _751_cf19
				var _752_cf18 _dafny.MultiSet = _747___mcc_h12
				_ = _752_cf18
				var _753_cf17 _dafny.Int = _746___mcc_h11
				_ = _753_cf17
				var _754_cf16 _dafny.Int = _745___mcc_h10
				_ = _754_cf16
				var _arr6 _dafny.Array = _this.F2()
				_ = _arr6
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index97
				(_arr6).ArraySet1((_this).F4(), (_index97).Int())
				var _755_v33 _dafny.Map
				_ = _755_v33
				_755_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm20(Companion_Default___.Fm20((_this).F4(), globalState), globalState), (_this).F4())
				var _756_v34 _dafny.Sequence
				_ = _756_v34
				_756_v34 = _dafny.SeqOf(_dafny.IntOfInt64(407), _751_cf19)
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index98
				var _arr7 _dafny.Array = _this.F2()
				_ = _arr7
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index99
				var _rhs75 bool = (func() bool {
					if (_755_v33).Contains((func() bool {
						if (_this).F4() {
							return (_this).F4()
						}
						return (_this).F4()
					})()) {
						return (_755_v33).Get((func() bool {
							if (_this).F4() {
								return (_this).F4()
							}
							return (_this).F4()
						})()).(bool)
					}
					return (_this).F4()
				})()
				_ = _rhs75
				var _rhs76 _dafny.Int = (func() _dafny.Int {
					if (_this).F4() {
						return (_756_v34).Select((Companion_Default___.SafeIndex((_dafny.SetOf((_this).F4())).Cardinality(), _dafny.IntOfUint32((_756_v34).Cardinality()))).Uint32()).(_dafny.Int)
					}
					return ((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)).Minus((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int))
				})()
				_ = _rhs76
				var _rhs77 _dafny.Int = Companion_Default___.Fm0(Companion_Default___.SafeDivisionInt(_750_cf20, _dafny.IntOfUint32((_692_v0).Cardinality())), globalState)
				_ = _rhs77
				var _rhs78 bool = (_this).F4()
				_ = _rhs78
				var _lhs46 _dafny.Array = _713_v5
				_ = _lhs46
				var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _lhs47
				var _lhs48 _dafny.Array = _this.F2()
				_ = _lhs48
				var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs49
				r1 = _rhs75
				_754_cf16 = _rhs76
				(_lhs46).ArraySet1(_rhs77, (_lhs47).Int())
				(_lhs48).ArraySet1(_rhs78, (_lhs49).Int())
				var _arr8 _dafny.Array = _this.F2()
				_ = _arr8
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index100
				(_arr8).ArraySet1(!((_731_v23).Difference(_731_v23)).Contains((_750_cf20).Times(_754_cf16)), (_index100).Int())
				_694_v2 = _694_v2
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index101
				(_713_v5).ArraySet1((func() _dafny.Int {
					if (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool) {
						return _754_cf16
					}
					return _750_cf20
				})(), (_index101).Int())
			} else if _source13.Is_DC5() {
				var _757___mcc_h15 _dafny.Sequence = _source13.Get_().(D2_DC5).Cf13
				_ = _757___mcc_h15
				var _758_cf13 _dafny.Sequence = _757___mcc_h15
				_ = _758_cf13
				var _759_v35 *C0
				_ = _759_v35
				var _nw106 *C0 = New_C0_()
				_ = _nw106
				_nw106.Ctor__()
				_759_v35 = _nw106
				_758_cf13 = _696_v4
				_758_cf13 = _696_v4
				var _760_v36 _dafny.Map
				_ = _760_v36
				_760_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _730_v22)
				var _761_v37 _dafny.Array
				_ = _761_v37
				var _nwElement0_20 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_730_v22, (Companion_Default___.SafeIndex((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_730_v22).Cardinality()))).Uint32(), _758_cf13)
				_ = _nwElement0_20
				var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(13))
				_ = _nw107
				(_nw107).ArraySet1(_nwElement0_20, 0)
				(_nw107).ArraySet1(_730_v22, 1)
				(_nw107).ArraySet1(_730_v22, 2)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_730_v22, Companion_Default___.Fm21(globalState)), 3)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_696_v4), _730_v22), 4)
				(_nw107).ArraySet1(_730_v22, 5)
				(_nw107).ArraySet1(_730_v22, 6)
				(_nw107).ArraySet1(_dafny.SeqOf(_758_cf13, _758_cf13), 7)
				(_nw107).ArraySet1((func() _dafny.Sequence {
					if (_760_v36).Contains(Companion_Default___.Fm22(_758_cf13, globalState)) {
						return (_760_v36).Get(Companion_Default___.Fm22(_758_cf13, globalState)).(_dafny.Sequence)
					}
					return Companion_Default___.Fm21(globalState)
				})(), 8)
				(_nw107).ArraySet1(_730_v22, 9)
				(_nw107).ArraySet1(_730_v22, 10)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_730_v22, _730_v22), 11)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_730_v22, _730_v22), 12)
				_761_v37 = _nw107
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_761_v37), 0))
				_ = _index102
				(_761_v37).ArraySet1(_730_v22, (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_761_v37), 0))
				_ = _index103
				(_761_v37).ArraySet1(Companion_Default___.Fm21(globalState), (_index103).Int())
			} else {
				var _762___mcc_h16 D2 = _source13.Get_().(D2_DC8).Cf21
				_ = _762___mcc_h16
				var _763_cf21 D2 = _762___mcc_h16
				_ = _763_cf21
				var _764_v38 _dafny.Array
				_ = _764_v38
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_24
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) D4 = func(_765_i1 _dafny.Int) D4 {
						return Companion_D4_.Create_DC12_()
					}
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw108 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw108).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw108).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_764_v38 = _nw108
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_764_v38), 0))
				_ = _index104
				(_764_v38).ArraySet1(Companion_Default___.Fm23(globalState), (_index104).Int())
				var _766_v39 D4
				_ = _766_v39
				_766_v39 = Companion_D4_.Create_DC12_()
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_764_v38), 0))
				_ = _index105
				(_764_v38).ArraySet1(_766_v39, (_index105).Int())
				var _767_v40 _dafny.Map
				_ = _767_v40
				_767_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int))
				var _768_v41 _dafny.Map
				_ = _768_v41
				_768_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				_767_v40 = (_767_v40).Update(_this.F2(), ((_768_v41).Cardinality()).Minus((_this).Fm16(false, globalState)))
				var _nw109 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw109
				(_this).F2_set_(_nw109)
				var _arr9 _dafny.Array = _this.F2()
				_ = _arr9
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index106
				(_arr9).ArraySet1((_this).F4(), (_index106).Int())
				var _769_v42 _dafny.MultiSet
				_ = _769_v42
				_769_v42 = _dafny.MultiSetOf(p1)
				var _arr10 _dafny.Array = _this.F2()
				_ = _arr10
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index107
				(_arr10).ArraySet1((_769_v42).IsProperSubsetOf(_769_v42), (_index107).Int())
			}
			var _770_v43 _dafny.Map
			_ = _770_v43
			_770_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), true)
			var _771_v44 _dafny.Sequence
			_ = _771_v44
			_771_v44 = _dafny.SeqOf(p0, (_770_v43).Cardinality(), p0)
			_771_v44 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F4())).Cardinality()))
			var _772_v45 _dafny.Sequence
			_ = _772_v45
			_772_v45 = _dafny.SeqOf(_731_v23)
			var _773_v46 _dafny.MultiSet
			_ = _773_v46
			_773_v46 = _dafny.MultiSetOf((_772_v45).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_772_v45).Cardinality()))).Uint32()).(_dafny.MultiSet))
			var _774_v47 D2
			_ = _774_v47
			_774_v47 = Companion_D2_.Create_DC7_(p0, _dafny.IntOfUint32((_696_v4).Cardinality()), _773_v46, (_dafny.Zero).Minus(p0), (_this).Fm2(globalState))
			_774_v47 = Companion_Default___.Fm24(_730_v22, globalState)
			var _775_v48 _dafny.Map
			_ = _775_v48
			_775_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int))
			var _776_v49 _dafny.Map
			_ = _776_v49
			_776_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_696_v4, _775_v48)
			var _777_v51 _dafny.Set
			_ = _777_v51
			_777_v51 = _dafny.SetOf(_696_v4)
			if ((_777_v51).Difference(_777_v51)).IsSubsetOf(func() _dafny.Set {
				var _coll40 = _dafny.NewBuilder()
				_ = _coll40
				for _iter44 := _dafny.Iterate((_776_v49).Keys().Elements()); ; {
					_compr_40, _ok44 := _iter44()
					if !_ok44 {
						break
					}
					var _778_v50 _dafny.Sequence
					_778_v50 = interface{}(_compr_40).(_dafny.Sequence)
					if (_776_v49).Contains(_778_v50) {
						_coll40.Add(_778_v50)
					}
				}
				return _coll40.ToSet()
			}()) {
				var _779_v52 *C0
				_ = _779_v52
				var _nw110 *C0 = New_C0_()
				_ = _nw110
				_nw110.Ctor__()
				_779_v52 = _nw110
				var _780_v53 _dafny.Array
				_ = _780_v53
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_25
				var _nw111 _dafny.Array
				_ = _nw111
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw111 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.Sequence = (func(_781_v44 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_782_i2 _dafny.Int) _dafny.Sequence {
							return _781_v44
						}
					})(_771_v44)
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw111 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw111).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw111).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_780_v53 = _nw111
				var _783_v54 D5
				_ = _783_v54
				_783_v54 = Companion_D5_.Create_DC14_(_780_v53)
				var _784_v55 _dafny.Array
				_ = _784_v55
				var _nwElement0_21 _dafny.Array = _780_v53
				_ = _nwElement0_21
				var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(24))
				_ = _nw112
				(_nw112).ArraySet1(_nwElement0_21, 0)
				(_nw112).ArraySet1((_783_v54).Dtor_cf32(), 1)
				(_nw112).ArraySet1((func() _dafny.Array {
					if (func() bool {
						if (_770_v43).Contains((_this).F4()) {
							return (_770_v43).Get((_this).F4()).(bool)
						}
						return true
					})() {
						return _780_v53
					}
					return _780_v53
				})(), 2)
				(_nw112).ArraySet1(_780_v53, 3)
				(_nw112).ArraySet1(_780_v53, 4)
				(_nw112).ArraySet1(_780_v53, 5)
				(_nw112).ArraySet1(_780_v53, 6)
				(_nw112).ArraySet1(_780_v53, 7)
				(_nw112).ArraySet1(_780_v53, 8)
				(_nw112).ArraySet1(_780_v53, 9)
				(_nw112).ArraySet1(_780_v53, 10)
				(_nw112).ArraySet1(_780_v53, 11)
				(_nw112).ArraySet1(_780_v53, 12)
				(_nw112).ArraySet1(_780_v53, 13)
				(_nw112).ArraySet1(_780_v53, 14)
				(_nw112).ArraySet1(_780_v53, 15)
				(_nw112).ArraySet1(_780_v53, 16)
				(_nw112).ArraySet1(_780_v53, 17)
				(_nw112).ArraySet1(_780_v53, 18)
				(_nw112).ArraySet1(_780_v53, 19)
				(_nw112).ArraySet1(_780_v53, 20)
				(_nw112).ArraySet1(_780_v53, 21)
				(_nw112).ArraySet1(_780_v53, 22)
				(_nw112).ArraySet1(_780_v53, 23)
				_784_v55 = _nw112
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_784_v55), 0))
				_ = _index108
				(_784_v55).ArraySet1(_780_v53, (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_784_v55), 0))
				_ = _index109
				(_784_v55).ArraySet1(_780_v53, (_index109).Int())
				var _785_v56 _dafny.Array
				_ = _785_v56
				var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
				_ = _nw113
				_785_v56 = _nw113
				var _786_v58 _dafny.Set
				_ = _786_v58
				_786_v58 = _dafny.SetOf(_775_v48)
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_785_v56), 0))
				_ = _index110
				(_785_v56).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_775_v48, p1)).Merge((func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter45 := _dafny.Iterate((_786_v58).Elements()); ; {
						_compr_41, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _787_v57 _dafny.Map
						_787_v57 = interface{}(_compr_41).(_dafny.Map)
						if (_786_v58).Contains(_787_v57) {
							_coll41.Add(_787_v57, p1)
						}
					}
					return _coll41.ToMap()
				}()).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), p0), _dafny.CodePoint('q'))), (_index110).Int())
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_785_v56), 0))
				_ = _index111
				(_785_v56).ArraySet1(func() _dafny.Map {
					var _coll42 = _dafny.NewMapBuilder()
					_ = _coll42
					for _iter46 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(585))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg33 _dafny.Int) interface{} {
							return coer33(arg33)
						}
					}((func(_788_v5 _dafny.Array) func(_dafny.Int) _dafny.Map {
						return func(_789_i3 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_788_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_788_v5), 0))).Int()).(_dafny.Int))
						}
					})(_713_v5)))).Elements()); ; {
						_compr_42, _ok46 := _iter46()
						if !_ok46 {
							break
						}
						var _790_v59 _dafny.Map
						_790_v59 = interface{}(_compr_42).(_dafny.Map)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(585))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg34 _dafny.Int) interface{} {
								return coer34(arg34)
							}
						}((func(_791_v5 _dafny.Array) func(_dafny.Int) _dafny.Map {
							return func(_789_i3 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_791_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_791_v5), 0))).Int()).(_dafny.Int))
							}
						})(_713_v5))), _790_v59) {
							_coll42.Add(_790_v59, _dafny.CodePoint('e'))
						}
					}
					return _coll42.ToMap()
				}(), (_index111).Int())
				var _792_v60 *C0
				_ = _792_v60
				var _nw114 *C0 = New_C0_()
				_ = _nw114
				_nw114.Ctor__()
				_792_v60 = _nw114
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index112
				(_713_v5).ArraySet1((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int), (_index112).Int())
			} else {
				var _arr11 _dafny.Array = _this.F2()
				_ = _arr11
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index113
				(_arr11).ArraySet1((_this).F4(), (_index113).Int())
				var _arr12 _dafny.Array = _this.F2()
				_ = _arr12
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index115
				var _rhs79 bool = !(_dafny.Companion_Sequence_.IsPrefixOf(_692_v0, _692_v0))
				_ = _rhs79
				var _rhs80 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, p0)
				_ = _rhs80
				var _lhs50 _dafny.Array = _this.F2()
				_ = _lhs50
				var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs51
				var _lhs52 _dafny.Array = _713_v5
				_ = _lhs52
				var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _lhs53
				(_lhs50).ArraySet1(_rhs79, (_lhs51).Int())
				(_lhs52).ArraySet1(_rhs80, (_lhs53).Int())
				_775_v48 = (_775_v48).Update((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), p0)
				var _793_v61 _dafny.Array
				_ = _793_v61
				var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
				_ = _nw115
				_793_v61 = _nw115
				var _794_v62 _dafny.Array
				_ = _794_v62
				var _nwElement0_22 _dafny.Map = _770_v43
				_ = _nwElement0_22
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(12))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_22, 0)
				(_nw116).ArraySet1(_770_v43, 1)
				(_nw116).ArraySet1(_770_v43, 2)
				(_nw116).ArraySet1(_770_v43, 3)
				(_nw116).ArraySet1(_770_v43, 4)
				(_nw116).ArraySet1(_770_v43, 5)
				(_nw116).ArraySet1((_770_v43).Update((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)), 6)
				(_nw116).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)), 7)
				(_nw116).ArraySet1(_770_v43, 8)
				(_nw116).ArraySet1((_770_v43).Update((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), !((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))), 9)
				(_nw116).ArraySet1(_770_v43, 10)
				(_nw116).ArraySet1(_770_v43, 11)
				_794_v62 = _nw116
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_793_v61), 0))
				_ = _index116
				(_793_v61).ArraySet1(_794_v62, (_index116).Int())
				var _795_v63 D6
				_ = _795_v63
				_795_v63 = Companion_D6_.Create_DC19_(_794_v62)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_793_v61), 0))
				_ = _index117
				(_793_v61).ArraySet1((_795_v63).Dtor_cf37(), (_index117).Int())
				var _796_v64 D2
				_ = _796_v64
				_796_v64 = Companion_D2_.Create_DC6_(p1, p0)
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
				_ = _index118
				(_713_v5).ArraySet1((_796_v64).Dtor_cf15(), (_index118).Int())
				var _797_v65 _dafny.Set
				_ = _797_v65
				_797_v65 = _dafny.SetOf(_713_v5, _713_v5, _713_v5)
				var _798_v66 _dafny.Sequence
				_ = _798_v66
				_798_v66 = _dafny.SeqOf(_797_v65)
				r1 = (_797_v65).IsProperSubsetOf((_798_v66).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_696_v4).Cardinality()), _dafny.IntOfUint32((_798_v66).Cardinality()))).Uint32()).(_dafny.Set))
			}
		}
		var _799_v67 _dafny.Map
		_ = _799_v67
		_799_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_692_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.IntOfUint32((_692_v0).Cardinality()))).Uint32()).(bool), (_this).F4())
		var _800_v68 _dafny.Sequence
		_ = _800_v68
		_800_v68 = _dafny.SeqOf(_799_v67, _799_v67, (_799_v67).Merge(_799_v67))
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
		_ = _index119
		var _rhs81 _dafny.Sequence = _800_v68
		_ = _rhs81
		var _rhs82 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int)))
		_ = _rhs82
		var _rhs83 bool = !(!((_this).F4()))
		_ = _rhs83
		var _rhs84 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nfnnwvpf"), _696_v4)
		_ = _rhs84
		var _rhs85 bool = (_this).F4()
		_ = _rhs85
		var _lhs54 _dafny.Array = _713_v5
		_ = _lhs54
		var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))
		_ = _lhs55
		_800_v68 = _rhs81
		(_lhs54).ArraySet1(_rhs82, (_lhs55).Int())
		r1 = _rhs83
		_696_v4 = _rhs84
		r1 = _rhs85
		(_this).F2_set_(_this.F2())
		var _801_v69 _dafny.MultiSet
		_ = _801_v69
		_801_v69 = _dafny.MultiSetOf(_this.F2())
		var _802_v71 _dafny.Map
		_ = _802_v71
		_802_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_713_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(400), _dafny.ArrayLen((_713_v5), 0))).Int()).(_dafny.Int))
		var _803_v72 _dafny.Sequence
		_ = _803_v72
		_803_v72 = _dafny.SeqOf((func() _dafny.Map {
			var _coll43 = _dafny.NewMapBuilder()
			_ = _coll43
			for _iter47 := _dafny.Iterate((_802_v71).Keys().Elements()); ; {
				_compr_43, _ok47 := _iter47()
				if !_ok47 {
					break
				}
				var _804_v70 _dafny.CodePoint
				_804_v70 = interface{}(_compr_43).(_dafny.CodePoint)
				if (_802_v71).Contains(_804_v70) {
					_coll43.Add(_804_v70, (_this).F4())
				}
			}
			return _coll43.ToMap()
		}()).Cardinality(), p0)
		r0 = (_801_v69).Update((func() _dafny.Array {
			if (_this).F4() {
				return _this.F2()
			}
			return _this.F2()
		})(), Companion_Default___.Abs((_803_v72).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_803_v72).Cardinality()))).Uint32()).(_dafny.Int)))
		r1 = (_this).F4()
		return r0, r1
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f2 _dafny.Array
	_f3 _dafny.Array
	_f4 bool
	F7  bool
	_f8 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = false
	_this.F7 = false
	_this._f8 = _dafny.Zero
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

func (_this *C2) F2() _dafny.Array {
	return _this._f2
}
func (_this *C2) F2_set_(value _dafny.Array) {
	_this._f2 = value
}
func (_this *C2) F3() _dafny.Array {
	return _this._f3
}
func (_this *C2) F4() bool {
	return _this._f4
}
func (_this *C2) Ctor__(f7 bool, f8 _dafny.Int, f3 _dafny.Array, f4 bool, f2 _dafny.Array) {
	{
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f2 = f2
	}
}
func (_this *C2) Fm3(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F8()
	}
}
func (_this *C2) Fm1(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F8()).Plus((_this).F8()), (_this).F4())
	}
}
func (_this *C2) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F8()).Minus(_dafny.IntOfInt64(396))
	}
}
func (_this *C2) Fm12(globalState *GlobalState) bool {
	{
		return _this.F7
	}
}
func (_this *C2) M0(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Map, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		r0 = (_this).Fm12(globalState)
		var _805_i0 _dafny.Int
		_ = _805_i0
		_805_i0 = _dafny.Zero
		{
			for !((_dafny.IntOfInt64(578)).Cmp(_dafny.IntOfInt64(298)) <= 0) || ((_this).F4()) {
				{
					if (_805_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_805_i0 = (_805_i0).Plus(_dafny.One)
					(_this).F7 = (_this).F4()
					var _806_v0 _dafny.Map
					_ = _806_v0
					_806_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), p0)
					var _807_v1 _dafny.Set
					_ = _807_v1
					_807_v1 = _dafny.SetOf(_806_v0)
					var _808_v2 _dafny.Map
					_ = _808_v2
					_808_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_807_v1).Cardinality())
					_808_v2 = (_808_v2).Update(!((_this).Fm12(globalState)), (p0).Plus((_this).F8()))
					var _809_v3 _dafny.Int
					_ = _809_v3
					_809_v3 = _dafny.IntOfInt64(615)
					_809_v3 = p0
					var _810_v4 _dafny.CodePoint
					_ = _810_v4
					_810_v4 = _dafny.CodePoint('c')
					var _811_v5 D2
					_ = _811_v5
					_811_v5 = Companion_D2_.Create_DC6_(_810_v4, (_this).F8())
					var _812_v6 D2
					_ = _812_v6
					_812_v6 = Companion_D2_.Create_DC8_(_811_v5)
					var _813_v7 _dafny.Set
					_ = _813_v7
					_813_v7 = _dafny.SetOf((func() D2 {
						if _this.F7 {
							return _812_v6
						}
						return _812_v6
					})(), _812_v6, _812_v6, (func() D2 {
						if false {
							return _812_v6
						}
						return _812_v6
					})(), _812_v6)
					var _814_v8 _dafny.Sequence
					_ = _814_v8
					_814_v8 = _dafny.SeqOf((_813_v7).Union(_dafny.SetOf(_812_v6, _812_v6, _812_v6, _812_v6, _812_v6)), _813_v7, (_813_v7).Intersection(_813_v7), (_813_v7).Difference(_813_v7))
					_813_v7 = (_814_v8).Select((Companion_Default___.SafeIndex(_809_v3, _dafny.IntOfUint32((_814_v8).Cardinality()))).Uint32()).(_dafny.Set)
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _815_v9 _dafny.Sequence
		_ = _815_v9
		_815_v9 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F8()))
		var _816_i1 _dafny.Int
		_ = _816_i1
		_816_i1 = _dafny.Zero
		{
			for !(((_815_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_815_v9).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_this).F8()) == 0) {
				{
					if (_816_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_816_i1 = (_816_i1).Plus(_dafny.One)
					var _817_v10 _dafny.Int
					_ = _817_v10
					_817_v10 = _dafny.IntOfInt64(-909)
					_817_v10 = _dafny.IntOfInt64(-880)
					var _818_v11 _dafny.Sequence
					_ = _818_v11
					_818_v11 = _dafny.UnicodeSeqOfUtf8Bytes("dsko")
					(_this).F7 = !(!_dafny.Companion_Sequence_.Equal(_818_v11, _818_v11))
					_817_v10 = p0
					var _819_v12 _dafny.Set
					_ = _819_v12
					_819_v12 = _dafny.SetOf(_this.F7, _this.F7, !(true), (_this).F4())
					var _820_v13 _dafny.Map
					_ = _820_v13
					_820_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, ((_819_v12).Cardinality()).Plus(_817_v10))
					_820_v13 = (_820_v13).Update(((_this).F8()).Cmp(_dafny.IntOfInt64(573)) >= 0, (_this).F8())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _821_v14 _dafny.CodePoint
		_ = _821_v14
		_821_v14 = _dafny.CodePoint('j')
		var _822_v15 D3
		_ = _822_v15
		_822_v15 = Companion_D3_.Create_DC10_((_this).F4(), _821_v14, _this.F7, (_this).F4(), _this.F7)
		var _source14 D3 = _822_v15
		_ = _source14
		if _source14.Is_DC10() {
			var _823___mcc_h0 bool = _source14.Get_().(D3_DC10).Cf23
			_ = _823___mcc_h0
			var _824___mcc_h1 _dafny.CodePoint = _source14.Get_().(D3_DC10).Cf24
			_ = _824___mcc_h1
			var _825___mcc_h2 bool = _source14.Get_().(D3_DC10).Cf25
			_ = _825___mcc_h2
			var _826___mcc_h3 bool = _source14.Get_().(D3_DC10).Cf26
			_ = _826___mcc_h3
			var _827___mcc_h4 bool = _source14.Get_().(D3_DC10).Cf27
			_ = _827___mcc_h4
			var _828_cf27 bool = _827___mcc_h4
			_ = _828_cf27
			var _829_cf26 bool = _826___mcc_h3
			_ = _829_cf26
			var _830_cf25 bool = _825___mcc_h2
			_ = _830_cf25
			var _831_cf24 _dafny.CodePoint = _824___mcc_h1
			_ = _831_cf24
			var _832_cf23 bool = _823___mcc_h0
			_ = _832_cf23
			var _833_v16 _dafny.Map
			_ = _833_v16
			_833_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-883))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_834_i2 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_this).F8())
			})), _this.F2())
			_833_v16 = _833_v16
			var _835_v17 _dafny.Sequence
			_ = _835_v17
			_835_v17 = _dafny.UnicodeSeqOfUtf8Bytes("ngnnvi")
			var _836_v18 _dafny.MultiSet
			_ = _836_v18
			_836_v18 = _dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_835_v17).Cardinality())))
			var _837_v19 _dafny.MultiSet
			_ = _837_v19
			_837_v19 = _dafny.MultiSetOf(_836_v18)
			var _838_v20 _dafny.Map
			_ = _838_v20
			_838_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_830_cf25, p0)
			var _839_v21 _dafny.Sequence
			_ = _839_v21
			_839_v21 = _dafny.SeqOf(_832_cf23, _828_cf27)
			var _840_v22 _dafny.Array
			_ = _840_v22
			var _nwElement0_23 _dafny.Int = p0
			_ = _nwElement0_23
			var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(13))
			_ = _nw117
			(_nw117).ArraySet1(_nwElement0_23, 0)
			(_nw117).ArraySet1(p0, 1)
			(_nw117).ArraySet1((_this).Fm2(globalState), 2)
			(_nw117).ArraySet1(p0, 3)
			(_nw117).ArraySet1((_this).F8(), 4)
			(_nw117).ArraySet1(p0, 5)
			(_nw117).ArraySet1((_this).F8(), 6)
			(_nw117).ArraySet1((func() _dafny.Int {
				if (_837_v19).Contains(_836_v18) {
					return (_837_v19).Multiplicity(_836_v18)
				}
				return (_this).F8()
			})(), 7)
			(_nw117).ArraySet1((func() _dafny.Int {
				if (_838_v20).Contains(_832_cf23) {
					return (_838_v20).Get(_832_cf23).(_dafny.Int)
				}
				return p0
			})(), 8)
			(_nw117).ArraySet1(_dafny.IntOfUint32((_835_v17).Cardinality()), 9)
			(_nw117).ArraySet1((_815_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_815_v9).Cardinality()))).Uint32()).(_dafny.Int), 10)
			(_nw117).ArraySet1(_dafny.IntOfUint32((_839_v21).Cardinality()), 11)
			(_nw117).ArraySet1((_this).Fm2(globalState), 12)
			_840_v22 = _nw117
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_840_v22), 0))
			_ = _index120
			(_840_v22).ArraySet1((_this).F8(), (_index120).Int())
			var _841_v23 _dafny.Int
			_ = _841_v23
			_841_v23 = _dafny.IntOfInt64(371)
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_840_v22), 0))
			_ = _index121
			var _rhs86 bool = _832_cf23
			_ = _rhs86
			var _rhs87 _dafny.Int = (_this).F8()
			_ = _rhs87
			var _rhs88 _dafny.Int = (_this).F8()
			_ = _rhs88
			var _rhs89 bool = _828_cf27
			_ = _rhs89
			var _rhs90 _dafny.Int = (_this).Fm2(globalState)
			_ = _rhs90
			var _lhs56 *C2 = _this
			_ = _lhs56
			var _lhs57 _dafny.Array = _840_v22
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_840_v22), 0))
			_ = _lhs58
			_lhs56.F7 = _rhs86
			(_lhs57).ArraySet1(_rhs87, (_lhs58).Int())
			_841_v23 = _rhs88
			_832_cf23 = _rhs89
			_841_v23 = _rhs90
			var _842_v24 _dafny.Map
			_ = _842_v24
			_842_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_821_v14, _829_cf26)
			var _843_v25 _dafny.Map
			_ = _843_v25
			_843_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC9_(_842_v24), _832_cf23)
			var _844_v26 D3
			_ = _844_v26
			_844_v26 = Companion_D3_.Create_DC9_(_842_v24)
			_843_v25 = (_843_v25).Update((func() D3 {
				if (_this).F4() {
					return _844_v26
				}
				return Companion_D3_.Create_DC9_(_842_v24)
			})(), _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm13(_830_cf25, _830_cf25, _832_cf23, globalState), _835_v17))
			_828_cf27 = (_this).Fm12(globalState)
		} else {
			var _845___mcc_h5 _dafny.Map = _source14.Get_().(D3_DC9).Cf22
			_ = _845___mcc_h5
			var _846_cf22 _dafny.Map = _845___mcc_h5
			_ = _846_cf22
			var _847_v27 _dafny.Sequence
			_ = _847_v27
			_847_v27 = _dafny.SeqOf(_this.F7, _this.F7)
			(_this).F7 = _dafny.Companion_Sequence_.Equal(_847_v27, _847_v27)
			var _848_v28 *C0
			_ = _848_v28
			var _nw118 *C0 = New_C0_()
			_ = _nw118
			_nw118.Ctor__()
			_848_v28 = _nw118
			if _this.F7 {
				(_this).F7 = (_847_v27).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-704), _dafny.IntOfUint32((_847_v27).Cardinality()))).Uint32()).(bool)
				var _849_v29 _dafny.Array
				_ = _849_v29
				var _nw119 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(2))
				_ = _nw119
				_849_v29 = _nw119
				var _850_v30 D0
				_ = _850_v30
				_850_v30 = Companion_D0_.Create_DC0_((_this).F4())
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_849_v29), 0))
				_ = _index122
				(_849_v29).ArraySet1(_850_v30, (_index122).Int())
				var _851_v31 _dafny.Map
				_ = _851_v31
				_851_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F8()), (_this).F4())
				var _852_v32 D1
				_ = _852_v32
				_852_v32 = Companion_D1_.Create_DC4_(_851_v31, _this.F7, (_this).F8(), (_this).F4())
				var _853_v33 _dafny.MultiSet
				_ = _853_v33
				_853_v33 = _dafny.MultiSetOf((_this).F4())
				var _854_v34 _dafny.Map
				_ = _854_v34
				_854_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_849_v29), 0))
				_ = _index123
				(_849_v29).ArraySet1(Companion_Default___.Fm14((_852_v32).Dtor_cf11(), p0, (_853_v33).IsProperSubsetOf(_853_v33), (func() bool {
					if (_854_v34).Contains(false) {
						return (_854_v34).Get(false).(bool)
					}
					return (_this).F4()
				})(), globalState), (_index123).Int())
				var _855_v35 _dafny.Set
				_ = _855_v35
				_855_v35 = _dafny.SetOf(_821_v14)
				var _856_v36 _dafny.Map
				_ = _856_v36
				_856_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_855_v35, _852_v32)
				_856_v36 = (_856_v36).Update(_855_v35, _852_v32)
				var _857_v37 _dafny.Int
				_ = _857_v37
				_857_v37 = _dafny.IntOfInt64(697)
				var _858_v38 _dafny.Map
				_ = _858_v38
				_858_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), ((_853_v33).Cardinality()).Times(p0))
				_857_v37 = (func() _dafny.Int {
					if (_858_v38).Contains(_this.F7) {
						return (_858_v38).Get(_this.F7).(_dafny.Int)
					}
					return _857_v37
				})()
				var _859_v39 _dafny.Set
				_ = _859_v39
				var _860_v40 bool
				_ = _860_v40
				var _out24 _dafny.Set
				_ = _out24
				var _out25 bool
				_ = _out25
				_out24, _out25 = (_this).M4((_this).F4(), globalState)
				_859_v39 = _out24
				_860_v40 = _out25
			} else {
				(_this).M5((_this).F4(), (_this.F7) && (_this.F7), globalState)
				var _arr13 _dafny.Array = _this.F2()
				_ = _arr13
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index124
				(_arr13).ArraySet1(!(_dafny.Companion_Sequence_.IsPrefixOf(_847_v27, _dafny.SeqOf(_this.F7, (_this).F4()))), (_index124).Int())
				var _arr14 _dafny.Array = _this.F2()
				_ = _arr14
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index125
				(_arr14).ArraySet1((_this).Fm12(globalState), (_index125).Int())
				var _861_v41 _dafny.Set
				_ = _861_v41
				_861_v41 = _dafny.SetOf(_821_v14)
				_861_v41 = _861_v41
				var _862_v42 _dafny.Map
				_ = _862_v42
				_862_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-992), (_this).F4())
				var _863_v43 _dafny.Map
				_ = _863_v43
				_863_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_862_v42, !((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)))
				_863_v43 = (_863_v43).Update((_this).Fm1((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), _this.F7, globalState), (_this).F4())
				var _864_v44 _dafny.Array
				_ = _864_v44
				var _nwElement0_24 bool = (_this).F4()
				_ = _nwElement0_24
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.One)
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_24, 0)
				_864_v44 = _nw120
				var _865_v45 _dafny.Map
				_ = _865_v45
				_865_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _864_v44), _847_v27)
				var _866_v46 _dafny.Map
				_ = _866_v46
				_866_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _864_v44)
				_865_v45 = (_865_v45).Update((_866_v46).Merge(_866_v46), _dafny.Companion_Sequence_.Concatenate(_847_v27, _847_v27))
			}
			var _867_v47 _dafny.Sequence
			_ = _867_v47
			_867_v47 = _dafny.UnicodeSeqOfUtf8Bytes("mseo")
			_821_v14 = (_867_v47).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_867_v47).Cardinality()))).Uint32()).(_dafny.CodePoint)
		}
		if _this.F7 {
			var _868_v48 _dafny.Array
			_ = _868_v48
			var _nw121 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw121
			_868_v48 = _nw121
			var _869_v49 _dafny.Sequence
			_ = _869_v49
			_869_v49 = _dafny.SeqOf(_868_v48)
			var _870_v50 _dafny.Map
			_ = _870_v50
			_870_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _868_v48)
			var _871_v51 _dafny.Array
			_ = _871_v51
			var _nwElement0_25 _dafny.Array = _868_v48
			_ = _nwElement0_25
			var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(28))
			_ = _nw122
			(_nw122).ArraySet1(_nwElement0_25, 0)
			(_nw122).ArraySet1(_868_v48, 1)
			(_nw122).ArraySet1(_868_v48, 2)
			(_nw122).ArraySet1(_868_v48, 3)
			(_nw122).ArraySet1(_868_v48, 4)
			(_nw122).ArraySet1(_868_v48, 5)
			(_nw122).ArraySet1(_868_v48, 6)
			(_nw122).ArraySet1(_868_v48, 7)
			(_nw122).ArraySet1(_868_v48, 8)
			(_nw122).ArraySet1(_868_v48, 9)
			(_nw122).ArraySet1(_868_v48, 10)
			(_nw122).ArraySet1(_868_v48, 11)
			(_nw122).ArraySet1(_868_v48, 12)
			(_nw122).ArraySet1(_868_v48, 13)
			(_nw122).ArraySet1((_869_v49).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F8()), _dafny.IntOfUint32((_869_v49).Cardinality()))).Uint32()).(_dafny.Array), 14)
			(_nw122).ArraySet1(_868_v48, 15)
			(_nw122).ArraySet1(_868_v48, 16)
			(_nw122).ArraySet1((func() _dafny.Array {
				if (_870_v50).Contains(p0) {
					return (_870_v50).Get(p0).(_dafny.Array)
				}
				return _868_v48
			})(), 17)
			(_nw122).ArraySet1((_869_v49).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_869_v49).Cardinality()))).Uint32()).(_dafny.Array), 18)
			(_nw122).ArraySet1(_868_v48, 19)
			(_nw122).ArraySet1(_868_v48, 20)
			(_nw122).ArraySet1(_868_v48, 21)
			(_nw122).ArraySet1(_868_v48, 22)
			(_nw122).ArraySet1(_868_v48, 23)
			(_nw122).ArraySet1(_868_v48, 24)
			(_nw122).ArraySet1(_868_v48, 25)
			(_nw122).ArraySet1(_868_v48, 26)
			(_nw122).ArraySet1(_868_v48, 27)
			_871_v51 = _nw122
			_871_v51 = _871_v51
			(_this).F7 = _this.F7
			var _872_v52 _dafny.Int
			_ = _872_v52
			_872_v52 = _dafny.IntOfInt64(555)
			_872_v52 = p0
			var _873_v53 _dafny.Map
			_ = _873_v53
			_873_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_this).F4())
			var _874_v54 D0
			_ = _874_v54
			_874_v54 = Companion_D0_.Create_DC0_((func() bool {
				if (_873_v53).Contains(_this.F7) {
					return (_873_v53).Get(_this.F7).(bool)
				}
				return !((_this).F4())
			})())
			var _source15 D0 = func(_pat_let11_0 D0) D0 {
				return func(_875_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let12_0 bool) D0 {
						return func(_876_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_876_dt__update_hcf0_h0)
						}(_pat_let12_0)
					}(true)
				}(_pat_let11_0)
			}(_874_v54)
			_ = _source15
			if _source15.Is_DC1() {
				var _877___mcc_h6 bool = _source15.Get_().(D0_DC1).Cf1
				_ = _877___mcc_h6
				var _878___mcc_h7 bool = _source15.Get_().(D0_DC1).Cf2
				_ = _878___mcc_h7
				var _879___mcc_h8 _dafny.Int = _source15.Get_().(D0_DC1).Cf3
				_ = _879___mcc_h8
				var _880_cf3 _dafny.Int = _879___mcc_h8
				_ = _880_cf3
				var _881_cf2 bool = _878___mcc_h7
				_ = _881_cf2
				var _882_cf1 bool = _877___mcc_h6
				_ = _882_cf1
				var _883_v55 _dafny.Sequence
				_ = _883_v55
				_883_v55 = _dafny.SeqOf(_dafny.CodePoint('p'), _821_v14)
				var _884_v56 _dafny.Sequence
				_ = _884_v56
				_884_v56 = _dafny.SeqOf(_883_v55, _883_v55)
				var _885_v57 _dafny.Map
				_ = _885_v57
				_885_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_884_v56, _872_v52)
				_885_v57 = (_885_v57).Update(_884_v56, p0)
				var _arr15 _dafny.Array = _this.F2()
				_ = _arr15
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index126
				(_arr15).ArraySet1(_882_cf1, (_index126).Int())
				var _arr16 _dafny.Array = _this.F2()
				_ = _arr16
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index127
				(_arr16).ArraySet1(!(_this.F7) || ((_this).Fm12(globalState)), (_index127).Int())
				var _886_v58 _dafny.Array
				_ = _886_v58
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_26
				var _nw123 _dafny.Array
				_ = _nw123
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw123 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Sequence = (func(_887_v55 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_888_i3 _dafny.Int) _dafny.Sequence {
							return _887_v55
						}
					})(_883_v55)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw123 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw123).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw123).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_886_v58 = _nw123
				var _889_v59 _dafny.MultiSet
				_ = _889_v59
				_889_v59 = _dafny.MultiSetOf(_dafny.IntOfInt64(367), (_this).F8())
				var _890_v60 _dafny.Map
				_ = _890_v60
				_890_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_886_v58, (_889_v59).IsDisjointFrom(_889_v59))
				_881_cf2 = (func() bool {
					if (_890_v60).Contains(_886_v58) {
						return (_890_v60).Get(_886_v58).(bool)
					}
					return _881_cf2
				})()
				_872_v52 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
					if (_this).Fm12(globalState) {
						return _872_v52
					}
					return _880_cf3
				})(), _872_v52)
			} else if _source15.Is_DC2() {
				var _891___mcc_h9 _dafny.Int = _source15.Get_().(D0_DC2).Cf4
				_ = _891___mcc_h9
				var _892___mcc_h10 bool = _source15.Get_().(D0_DC2).Cf5
				_ = _892___mcc_h10
				var _893___mcc_h11 bool = _source15.Get_().(D0_DC2).Cf6
				_ = _893___mcc_h11
				var _894___mcc_h12 _dafny.Int = _source15.Get_().(D0_DC2).Cf7
				_ = _894___mcc_h12
				var _895_cf7 _dafny.Int = _894___mcc_h12
				_ = _895_cf7
				var _896_cf6 bool = _893___mcc_h11
				_ = _896_cf6
				var _897_cf5 bool = _892___mcc_h10
				_ = _897_cf5
				var _898_cf4 _dafny.Int = _891___mcc_h9
				_ = _898_cf4
				var _899_v61 _dafny.Set
				_ = _899_v61
				_899_v61 = _dafny.SetOf(_dafny.IntOfInt64(892), _dafny.IntOfInt64(516))
				var _900_v62 _dafny.Set
				_ = _900_v62
				_900_v62 = _dafny.SetOf(_896_cf6)
				_897_cf5 = !((_899_v61).IsProperSubsetOf(_899_v61)) || ((_900_v62).IsSubsetOf(_dafny.SetOf(_896_cf6, _897_cf5)))
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_868_v48), 0))
				_ = _index128
				(_868_v48).ArraySet1((_895_cf7).Plus(_895_cf7), (_index128).Int())
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_868_v48), 0))
				_ = _index129
				(_868_v48).ArraySet1(_872_v52, (_index129).Int())
				r0 = _897_cf5
				var _901_v63 D0
				_ = _901_v63
				_901_v63 = Companion_D0_.Create_DC1_(false, !(_897_cf5) || ((_this).F4()), _dafny.IntOfInt64(450))
				var _902_v64 _dafny.Sequence
				_ = _902_v64
				_902_v64 = _dafny.SeqOf(false)
				var _903_v65 _dafny.Map
				_ = _903_v65
				_903_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_872_v52, _896_cf6)
				var _904_v66 _dafny.Map
				_ = _904_v66
				_904_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_898_cf4, _903_v65)
				var _rhs91 _dafny.CodePoint = (func() _dafny.CodePoint {
					if _896_cf6 {
						return _dafny.CodePoint('k')
					}
					return _dafny.CodePoint('o')
				})()
				_ = _rhs91
				var _rhs92 _dafny.Array = _this.F2()
				_ = _rhs92
				var _rhs93 bool = (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_902_v64).Cardinality()), (_904_v66).Cardinality())).Cmp((_dafny.IntOfUint32((_902_v64).Cardinality())).Minus(p0)) == 0
				_ = _rhs93
				var _rhs94 bool = _897_cf5
				_ = _rhs94
				var _rhs95 D0 = _901_v63
				_ = _rhs95
				var _lhs59 *GlobalState = globalState
				_ = _lhs59
				_821_v14 = _rhs91
				_lhs59.F1 = _rhs92
				_896_cf6 = _rhs93
				r0 = _rhs94
				_901_v63 = _rhs95
			} else {
				var _905___mcc_h13 bool = _source15.Get_().(D0_DC0).Cf0
				_ = _905___mcc_h13
				var _906_cf0 bool = _905___mcc_h13
				_ = _906_cf0
				var _907_v67 _dafny.MultiSet
				_ = _907_v67
				_907_v67 = _dafny.MultiSetOf(_this.F7, _906_cf0, (_this).F4())
				var _908_v68 _dafny.Map
				_ = _908_v68
				_908_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm15(_dafny.IntOfInt64(-368), _872_v52, globalState)).Dtor_cf1(), _dafny.MultiSetOf(false, _this.F7, _906_cf0, (_this).F4()))
				_907_v67 = (func() _dafny.MultiSet {
					if (_908_v68).Contains((_872_v52).Cmp(_dafny.IntOfInt64(755)) < 0) {
						return (_908_v68).Get((_872_v52).Cmp(_dafny.IntOfInt64(755)) < 0).(_dafny.MultiSet)
					}
					return _907_v67
				})()
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw124
				_868_v48 = _nw124
				_872_v52 = (_this).F8()
				var _909_v69 _dafny.Array
				_ = _909_v69
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_27
				var _nw125 _dafny.Array
				_ = _nw125
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw125 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Sequence = (func(_910_cf0 bool, _911_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_912_i4 _dafny.Int) _dafny.Sequence {
							return (func() _dafny.Sequence {
								if _910_cf0 {
									return _dafny.UnicodeSeqOfUtf8Bytes("pi")
								}
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-39))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg36 _dafny.Int) interface{} {
										return coer36(arg36)
									}
								}((func(_913_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_914_i5 _dafny.Int) _dafny.CodePoint {
										return _913_v14
									}
								})(_911_v14)))
							})()
						}
					})(_906_cf0, _821_v14)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw125 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw125).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw125).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_909_v69 = _nw125
				var _915_v70 _dafny.Sequence
				_ = _915_v70
				_915_v70 = _dafny.UnicodeSeqOfUtf8Bytes("jemmqpy")
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_909_v69), 0))
				_ = _index130
				(_909_v69).ArraySet1(_915_v70, (_index130).Int())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_909_v69), 0))
				_ = _index131
				var _rhs96 _dafny.Int = ((p0).Plus(_872_v52)).Plus((_this).Fm2(globalState))
				_ = _rhs96
				var _rhs97 _dafny.Int = (_dafny.IntOfInt64(981)).Plus((func() _dafny.Int {
					if (_822_v15).Dtor_cf25() {
						return _872_v52
					}
					return _dafny.IntOfUint32((_dafny.SeqOf(_906_cf0, _this.F7)).Cardinality())
				})())
				_ = _rhs97
				var _rhs98 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wrwcm"), _915_v70)
				_ = _rhs98
				var _lhs60 _dafny.Array = _909_v69
				_ = _lhs60
				var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_909_v69), 0))
				_ = _lhs61
				_872_v52 = _rhs96
				_872_v52 = _rhs97
				(_lhs60).ArraySet1(_rhs98, (_lhs61).Int())
			}
			(_this).F7 = (_dafny.Companion_Sequence_.Contains(_815_v9, _dafny.IntOfInt64(819))) || ((_this).F4())
		} else {
			var _916_v71 _dafny.Int
			_ = _916_v71
			_916_v71 = _dafny.IntOfInt64(-829)
			_916_v71 = p0
			var _917_v72 *C0
			_ = _917_v72
			var _nw126 *C0 = New_C0_()
			_ = _nw126
			_nw126.Ctor__()
			_917_v72 = _nw126
			var _918_v73 _dafny.Array
			_ = _918_v73
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_28
			var _nw127 _dafny.Array
			_ = _nw127
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw127 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.CodePoint = (func(_919_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_920_i6 _dafny.Int) _dafny.CodePoint {
						return _919_v14
					}
				})(_821_v14)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw127 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw127).ArraySet1CodePoint(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw127).ArraySet1CodePoint(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_918_v73 = _nw127
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_918_v73), 0))
			_ = _index132
			(_918_v73).ArraySet1CodePoint(_821_v14, (_index132).Int())
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_918_v73), 0))
			_ = _index133
			(_918_v73).ArraySet1CodePoint(_821_v14, (_index133).Int())
			_916_v71 = _dafny.IntOfInt64(423)
			var _921_v74 _dafny.Set
			_ = _921_v74
			_921_v74 = _dafny.SetOf(_917_v72)
			var _922_v75 _dafny.Map
			_ = _922_v75
			_922_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), (_dafny.SetOf(_917_v72)).Difference(_921_v74))
			_922_v75 = (_922_v75).Update(_dafny.CodePoint('l'), (_921_v74).Union(_921_v74))
		}
		var _923_v76 _dafny.Int
		_ = _923_v76
		_923_v76 = _dafny.IntOfInt64(100)
		var _924_v77 T0
		_ = _924_v77
		var _nw128 *C1 = New_C1_()
		_ = _nw128
		_nw128.Ctor__((_this).F3(), (_this).F4(), _this.F2())
		_924_v77 = _nw128
		var _925_v78 _dafny.Map
		_ = _925_v78
		_925_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_924_v77, _924_v77)
		_923_v76 = (_dafny.Zero).Minus((_923_v76).Minus(((_dafny.SetOf(p0, ((_925_v78).Update(_924_v77, _924_v77)).Cardinality())).Cardinality()).Minus(_923_v76)))
		r0 = (func() bool {
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_815_v9, _dafny.SeqOf(p0, _923_v76)) {
				return (_this).F4()
			}
			return (_this).F4()
		})()
		return r0
	}
}
func (_this *C2) M1(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Array, p3 _dafny.Array, globalState *GlobalState) (_dafny.MultiSet, bool) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 bool = false
		_ = r1
		for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_this.F2()), 0))); ; {
			_guard_loop_4, _ok48 := _iter48()
			if !_ok48 {
				break
			}
			var _926_i0 _dafny.Int
			_926_i0 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_926_i0).Sign() != -1) && ((_926_i0).Cmp(_dafny.ArrayLen((_this.F2()), 0)) < 0)) {
				var _arr17 _dafny.Array = _this.F2()
				_ = _arr17
				(_arr17).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(834))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}(func(_927_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-224))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg38 _dafny.Int) interface{} {
							return coer38(arg38)
						}
					}(func(_928_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('q')
					}))
				})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(388))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}(func(_929_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("yv")
				}))), (_926_i0).Int())
			}
		}
		var _930_v0 _dafny.Sequence
		_ = _930_v0
		_930_v0 = _dafny.SeqOf(false)
		var _931_v1 D3
		_ = _931_v1
		_931_v1 = Companion_D3_.Create_DC10_((_this).F4(), p1, false, (_930_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_930_v0).Cardinality()), _dafny.IntOfUint32((_930_v0).Cardinality()))).Uint32()).(bool), (_this).F4())
		var _932_v2 _dafny.Sequence
		_ = _932_v2
		_932_v2 = _dafny.SeqOf(Companion_D0_.Create_DC1_((_this).F4(), true, p0))
		var _933_v3 D1
		_ = _933_v3
		_933_v3 = Companion_D1_.Create_DC3_(_932_v2)
		var _934_v4 _dafny.Sequence
		_ = _934_v4
		_934_v4 = _dafny.UnicodeSeqOfUtf8Bytes("fre")
		var _935_v5 _dafny.Map
		_ = _935_v5
		_935_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_933_v3, _dafny.IntOfUint32((_934_v4).Cardinality()))
		var _936_v6 _dafny.Map
		_ = _936_v6
		_936_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_931_v1).Dtor_cf24())).Update(_this.F7, p1)).Cardinality(), ((_935_v5).Update(Companion_D1_.Create_DC3_(_932_v2), (_this).F8())).Cardinality())
		(_this).F7 = (_936_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F8()))
		var _937_v7 _dafny.Sequence
		_ = _937_v7
		_937_v7 = _dafny.SeqOf(p0, _dafny.IntOfInt64(773), (_this).F8(), (_dafny.Zero).Minus(p0), p0)
		var _938_v8 D4
		_ = _938_v8
		_938_v8 = Companion_D4_.Create_DC11_(_dafny.Companion_Sequence_.Update(_930_v0, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_930_v0).Cardinality()))).Uint32(), _this.F7))
		var _939_v9 _dafny.Map
		_ = _939_v9
		_939_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _938_v8)
		var _rhs99 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(845))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_940_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_941_i4 _dafny.Int) _dafny.Int {
				return _940_p0
			}
		})(p0)))
		_ = _rhs99
		var _rhs100 _dafny.Map = (_939_v9).Merge(_939_v9)
		_ = _rhs100
		_937_v7 = _rhs99
		_939_v9 = _rhs100
		var _hi6 _dafny.Int = p0
		_ = _hi6
		for _942_i5 := ((_this).F8()).Times(p0); _942_i5.Cmp(_hi6) < 0; _942_i5 = _942_i5.Plus(_dafny.One) {
			var _943_v10 _dafny.Array
			_ = _943_v10
			var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(6))
			_ = _nw129
			_943_v10 = _nw129
			var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.One)
			_ = _nw130
			_943_v10 = _nw130
			if false {
				var _944_v11 _dafny.Int
				_ = _944_v11
				_944_v11 = _dafny.IntOfInt64(124)
				_944_v11 = _942_i5
				var _945_v12 _dafny.Set
				_ = _945_v12
				_945_v12 = _dafny.SetOf(_944_v11, p0)
				var _946_v13 _dafny.Array
				_ = _946_v13
				var _nwElement0_26 _dafny.Sequence = _dafny.SeqOf(_942_i5)
				_ = _nwElement0_26
				var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(19))
				_ = _nw131
				(_nw131).ArraySet1(_nwElement0_26, 0)
				(_nw131).ArraySet1(_937_v7, 1)
				(_nw131).ArraySet1(_dafny.SeqOf((_this).F8()), 2)
				(_nw131).ArraySet1(_937_v7, 3)
				(_nw131).ArraySet1(_dafny.SeqOf((_this).F8(), _942_i5), 4)
				(_nw131).ArraySet1(_937_v7, 5)
				(_nw131).ArraySet1(_937_v7, 6)
				(_nw131).ArraySet1(_dafny.Companion_Sequence_.Update(_937_v7, (Companion_Default___.SafeIndex(_942_i5, _dafny.IntOfUint32((_937_v7).Cardinality()))).Uint32(), _944_v11), 7)
				(_nw131).ArraySet1(_937_v7, 8)
				(_nw131).ArraySet1(_937_v7, 9)
				(_nw131).ArraySet1(_937_v7, 10)
				(_nw131).ArraySet1(_dafny.Companion_Sequence_.Update(_937_v7, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_937_v7).Cardinality()))).Uint32(), (_this).F8()), 11)
				(_nw131).ArraySet1(_937_v7, 12)
				(_nw131).ArraySet1(_937_v7, 13)
				(_nw131).ArraySet1(_dafny.SeqOf(p0, _944_v11, (_945_v12).Cardinality()), 14)
				(_nw131).ArraySet1(_dafny.SeqOf((_this).F8()), 15)
				(_nw131).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(740))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_947_i6 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf((_this).F4(), (_this).F4())).Cardinality())
				})), 16)
				(_nw131).ArraySet1(_937_v7, 17)
				(_nw131).ArraySet1(_937_v7, 18)
				_946_v13 = _nw131
				var _948_v14 _dafny.Map
				_ = _948_v14
				_948_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC14_(_946_v13), _944_v11)
				_948_v14 = _948_v14
				r1 = _this.F7
				r1 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_934_v4, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_934_v4).Cardinality()))).Uint32(), p1), p1)
				r1 = (_this).F4()
			} else {
				var _arr18 _dafny.Array = _this.F2()
				_ = _arr18
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index134
				(_arr18).ArraySet1((_this).F4(), (_index134).Int())
				var _949_v15 D0
				_ = _949_v15
				_949_v15 = Companion_D0_.Create_DC2_((_this).F8(), (_this).F4(), (_this).F4(), _942_i5)
				var _950_v16 _dafny.Sequence
				_ = _950_v16
				_950_v16 = _dafny.SeqOf(func(_pat_let13_0 D0) D0 {
					return func(_951_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let14_0 bool) D0 {
							return func(_952_dt__update_hcf5_h0 bool) D0 {
								return Companion_D0_.Create_DC2_((_951_dt__update__tmp_h0).Dtor_cf4(), _952_dt__update_hcf5_h0, (_951_dt__update__tmp_h0).Dtor_cf6(), (_951_dt__update__tmp_h0).Dtor_cf7())
							}(_pat_let14_0)
						}(false)
					}(_pat_let13_0)
				}(_949_v15))
				var _953_v17 _dafny.Sequence
				_ = _953_v17
				_953_v17 = _dafny.SeqOf(_950_v16)
				var _arr19 _dafny.Array = _this.F2()
				_ = _arr19
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index135
				(_arr19).ArraySet1((Companion_Default___.Fm25((_this).F4(), globalState)).Contains((func() _dafny.Int {
					if _this.F7 {
						return _dafny.IntOfUint32((_953_v17).Cardinality())
					}
					return p0
				})()), (_index135).Int())
				var _954_v18 _dafny.Set
				_ = _954_v18
				var _955_v19 bool
				_ = _955_v19
				var _out26 _dafny.Set
				_ = _out26
				var _out27 bool
				_ = _out27
				_out26, _out27 = (_this).M4(_this.F7, globalState)
				_954_v18 = _out26
				_955_v19 = _out27
				_936_v6 = (_936_v6).Update((_942_i5).Times(p0), _942_i5)
				var _956_v20 _dafny.Array
				_ = _956_v20
				var _len0_29 _dafny.Int = _dafny.One
				_ = _len0_29
				var _nw132 _dafny.Array
				_ = _nw132
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw132 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Int = (func(_957_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_958_i7 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_958_i7, _957_p0)
						}
					})(p0)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw132 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw132).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw132).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_956_v20 = _nw132
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_956_v20), 0))
				_ = _index136
				(_956_v20).ArraySet1((_this).F8(), (_index136).Int())
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_956_v20), 0))
				_ = _index137
				(_956_v20).ArraySet1((_dafny.Zero).Minus(_942_i5), (_index137).Int())
				_955_v19 = (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)
			}
			var _959_v21 _dafny.Int
			_ = _959_v21
			_959_v21 = _dafny.IntOfInt64(337)
			_959_v21 = ((_this).F8()).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_934_v4, (Companion_Default___.SafeIndex(Companion_Default___.Fm0((_this).F8(), globalState), _dafny.IntOfUint32((_934_v4).Cardinality()))).Uint32(), p1)).Cardinality()))
			r1 = _this.F7
		}
		var _960_v22 _dafny.Int
		_ = _960_v22
		_960_v22 = _dafny.IntOfInt64(722)
		var _961_v23 _dafny.Map
		_ = _961_v23
		_961_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_934_v4).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_934_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_((_this).F4()), _930_v0))
		var _962_v24 D0
		_ = _962_v24
		_962_v24 = Companion_D0_.Create_DC0_(false)
		var _963_v25 _dafny.Map
		_ = _963_v25
		_963_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_934_v4, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rhdr")).Cardinality()))
		var _964_v26 _dafny.Map
		_ = _964_v26
		_964_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			if (_961_v23).Contains(p1) {
				return (_961_v23).Get(p1).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_962_v24, _930_v0)
		})(), ((_963_v25).Cardinality()).Plus((_this).F8()))
		var _965_v27 _dafny.Map
		_ = _965_v27
		_965_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_962_v24, _930_v0)
		_960_v22 = (func() _dafny.Int {
			if (_964_v26).Contains(_965_v27) {
				return (_964_v26).Get(_965_v27).(_dafny.Int)
			}
			return (_960_v22).Times((_this).F8())
		})()
		var _966_v28 _dafny.Sequence
		_ = _966_v28
		_966_v28 = _dafny.SeqOf(_this.F2(), _this.F2())
		var _967_i8 _dafny.Int
		_ = _967_i8
		_967_i8 = _dafny.Zero
		{
			var _pat_let_tv9 = _930_v0
			_ = _pat_let_tv9
			for ((_this).Fm3(_this.F7, _this.F7, _dafny.IntOfUint32((_966_v28).Cardinality()), p0, globalState)).Cmp(((_dafny.Zero).Minus((Companion_Default___.Fm26(_this.F7, _this.F7, func(_pat_let15_0 D5) D5 {
				return func(_994_dt__update__tmp_h1 D5) D5 {
					return func(_pat_let16_0 D4) D5 {
						return func(_995_dt__update_hcf33_h0 D4) D5 {
							return Companion_D5_.Create_DC15_(_995_dt__update_hcf33_h0, (_994_dt__update__tmp_h1).Dtor_cf34())
						}(_pat_let16_0)
					}(Companion_D4_.Create_DC11_(_pat_let_tv9))
				}(_pat_let15_0)
			}(Companion_D5_.Create_DC15_(_938_v8, (_this).F4())), _dafny.IntOfInt64(546), globalState)).Cardinality())).Minus(Companion_Default___.Fm0((_this).F8(), globalState))) == 0 {
				{
					if (_967_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_967_i8 = (_967_i8).Plus(_dafny.One)
					var _968_v29 _dafny.Map
					_ = _968_v29
					_968_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _this.F7)
					var _969_v30 *C1
					_ = _969_v30
					var _nw133 *C1 = New_C1_()
					_ = _nw133
					_nw133.Ctor__((_this).F3(), (func() bool {
						if (_968_v29).Contains(_this.F7) {
							return (_968_v29).Get(_this.F7).(bool)
						}
						return false
					})(), _this.F2())
					_969_v30 = _nw133
					if _this.F7 {
						var _970_v31 _dafny.Map
						_ = _970_v31
						_970_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(737))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg42 _dafny.Int) interface{} {
								return coer42(arg42)
							}
						}((func(_971_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_972_i9 _dafny.Int) _dafny.CodePoint {
								return _971_p1
							}
						})(p1)))).Cardinality()))
						var _973_v32 _dafny.Map
						_ = _973_v32
						_973_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-765), _970_v31)
						var _974_v33 _dafny.Map
						_ = _974_v33
						_974_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_934_v4, _973_v32)
						var _975_v34 _dafny.Map
						_ = _975_v34
						_975_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_937_v7, _this.F7)
						_974_v33 = (_974_v33).Update(_934_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if (_975_v34).Contains(_dafny.SeqOf(_dafny.IntOfInt64(688))) {
								return (_975_v34).Get(_dafny.SeqOf(_dafny.IntOfInt64(688))).(bool)
							}
							return false
						})(), _968_v29)).Cardinality(), _970_v31))
						(_this).F7 = _this.F7
						var _976_v35 _dafny.Set
						_ = _976_v35
						_976_v35 = _dafny.SetOf(_this.F7)
						var _977_v36 _dafny.MultiSet
						_ = _977_v36
						_977_v36 = _dafny.MultiSetOf((_this).F8(), (_this).F8())
						var _978_v37 _dafny.Array
						_ = _978_v37
						var _nwElement0_27 _dafny.Int = (_977_v36).Cardinality()
						_ = _nwElement0_27
						var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(3))
						_ = _nw134
						(_nw134).ArraySet1(_nwElement0_27, 0)
						(_nw134).ArraySet1(_dafny.IntOfUint32((_930_v0).Cardinality()), 1)
						(_nw134).ArraySet1((_this).F8(), 2)
						_978_v37 = _nw134
						var _979_v38 _dafny.Map
						_ = _979_v38
						_979_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v22, _this.F7)
						var _980_v39 _dafny.Map
						_ = _980_v39
						_980_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_978_v37, _979_v38)
						var _981_v40 _dafny.Map
						_ = _981_v40
						_981_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, _980_v39)
						var _982_v41 _dafny.Map
						_ = _982_v41
						_982_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), p0)
						var _rhs101 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_934_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(987))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg43 _dafny.Int) interface{} {
								return coer43(arg43)
							}
						}((func(_983_p1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_984_i10 _dafny.Int) _dafny.CodePoint {
								return _983_p1
							}
						})(p1))))
						_ = _rhs101
						var _rhs102 bool = !((func() _dafny.Map {
							if (_981_v40).Contains(_this.F7) {
								return (_981_v40).Get(_this.F7).(_dafny.Map)
							}
							return (_980_v39).Update(_978_v37, _979_v38)
						})()).Equals((_980_v39).Merge(_980_v39))
						_ = _rhs102
						var _rhs103 _dafny.Int = (func() _dafny.Int {
							if (_982_v41).Contains(_this.F2()) {
								return (_982_v41).Get(_this.F2()).(_dafny.Int)
							}
							return _dafny.IntOfInt64(197)
						})()
						_ = _rhs103
						var _rhs104 _dafny.Set = _976_v35
						_ = _rhs104
						_934_v4 = _rhs101
						r1 = _rhs102
						_960_v22 = _rhs103
						_976_v35 = _rhs104
						_979_v38 = ((_979_v38).Merge(_979_v38)).Merge(_979_v38)
						(_this).F2_set_(_this.F2())
					} else {
						(_this).F7 = _this.F7
						_960_v22 = _960_v22
						var _985_v42 *C1
						_ = _985_v42
						var _nw135 *C1 = New_C1_()
						_ = _nw135
						_nw135.Ctor__((_this).F3(), (_this).F4(), _this.F2())
						_985_v42 = _nw135
						var _986_v43 _dafny.Map
						_ = _986_v43
						_986_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm13(_this.F7, _this.F7, _this.F7, globalState), _934_v4)
						r1 = !(_986_v43).Contains(_934_v4)
						var _987_v44 _dafny.Map
						_ = _987_v44
						_987_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v22, false)
						var _988_v45 D1
						_ = _988_v45
						_988_v45 = Companion_D1_.Create_DC4_(_987_v44, false, _dafny.IntOfInt64(135), (_this).F4())
						var _989_v46 _dafny.Set
						_ = _989_v46
						_989_v46 = _dafny.SetOf(_988_v45)
						_989_v46 = _dafny.SetOf(_988_v45)
					}
					var _990_v47 _dafny.Sequence
					_ = _990_v47
					_990_v47 = _dafny.SeqOf(_969_v30, _969_v30, _969_v30)
					r1 = _dafny.Companion_Sequence_.IsPrefixOf(_990_v47, _990_v47)
					var _991_v48 _dafny.Set
					_ = _991_v48
					_991_v48 = _dafny.SetOf((_this).F4())
					var _992_v49 _dafny.Map
					_ = _992_v49
					_992_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(713), _991_v48)
					var _993_v50 _dafny.Sequence
					_ = _993_v50
					_993_v50 = _dafny.SeqOf(_991_v48, (func() _dafny.Set {
						if (_992_v49).Contains(_960_v22) {
							return (_992_v49).Get(_960_v22).(_dafny.Set)
						}
						return _991_v48
					})())
					r1 = (func() bool {
						if (_991_v48).IsProperSubsetOf((_993_v50).Select((Companion_Default___.SafeIndex(_960_v22, _dafny.IntOfUint32((_993_v50).Cardinality()))).Uint32()).(_dafny.Set)) {
							return (!((_this).F4())) && ((_this).F4())
						}
						return (true) || (_this.F7)
					})()
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _996_v51 _dafny.MultiSet
		_ = _996_v51
		_996_v51 = _dafny.MultiSetOf((_966_v28).Select((Companion_Default___.SafeIndex(_960_v22, _dafny.IntOfUint32((_966_v28).Cardinality()))).Uint32()).(_dafny.Array), _this.F2())
		var _997_v52 D7
		_ = _997_v52
		_997_v52 = Companion_D7_.Create_DC22_(_996_v51)
		r0 = ((_996_v51).Difference((_997_v52).Dtor_cf41())).Difference(_996_v51)
		r1 = _this.F7
		return r0, r1
	}
}
func (_this *C2) M4(p0 bool, globalState *GlobalState) (_dafny.Set, bool) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 bool = false
		_ = r1
		var _998_v0 _dafny.Array
		_ = _998_v0
		var _len0_30 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_30
		var _nw136 _dafny.Array
		_ = _nw136
		if _len0_30.Cmp(_dafny.Zero) == 0 {
			_nw136 = _dafny.NewArray(_len0_30)
		} else {
			var _init30 func(_dafny.Int) _dafny.Sequence = func(_999_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eporvvi"), _dafny.UnicodeSeqOfUtf8Bytes("mrfgh"))
			}
			_ = _init30
			var _element0_30 = _init30(_dafny.Zero)
			_ = _element0_30
			_nw136 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
			(_nw136).ArraySet1(_element0_30, 0)
			var _nativeLen0_30 = (_len0_30).Int()
			_ = _nativeLen0_30
			for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
				(_nw136).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
			}
		}
		_998_v0 = _nw136
		for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_998_v0), 0))); ; {
			_guard_loop_5, _ok49 := _iter49()
			if !_ok49 {
				break
			}
			var _1000_i0 _dafny.Int
			_1000_i0 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1000_i0).Sign() != -1) && ((_1000_i0).Cmp(_dafny.ArrayLen((_998_v0), 0)) < 0)) {
				(_998_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_this).F4() {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-422))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}(func(_1001_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('x')
						}))
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(540))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}(func(_1002_i3 _dafny.Int) _dafny.CodePoint {
						return (Companion_D2_.Create_DC6_(_dafny.CodePoint('n'), (_this).F8())).Dtor_cf14()
					}))
				})(), _dafny.UnicodeSeqOfUtf8Bytes("side")), (_1000_i0).Int())
			}
		}
		var _1003_i4 _dafny.Int
		_ = _1003_i4
		_1003_i4 = _dafny.Zero
		{
			for ((_this).F8()).Cmp((_this).F8()) == 0 {
				{
					if (_1003_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1003_i4 = (_1003_i4).Plus(_dafny.One)
					var _1004_v1 _dafny.Array
					_ = _1004_v1
					var _len0_31 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_31
					var _nw137 _dafny.Array
					_ = _nw137
					if _len0_31.Cmp(_dafny.Zero) == 0 {
						_nw137 = _dafny.NewArray(_len0_31)
					} else {
						var _init31 func(_dafny.Int) _dafny.Int = func(_1005_i5 _dafny.Int) _dafny.Int {
							return (_1005_i5).Times((_this).F8())
						}
						_ = _init31
						var _element0_31 = _init31(_dafny.Zero)
						_ = _element0_31
						_nw137 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
						(_nw137).ArraySet1(_element0_31, 0)
						var _nativeLen0_31 = (_len0_31).Int()
						_ = _nativeLen0_31
						for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
							(_nw137).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
						}
					}
					_1004_v1 = _nw137
					var _1006_v2 _dafny.Sequence
					_ = _1006_v2
					_1006_v2 = _dafny.UnicodeSeqOfUtf8Bytes("niyp")
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))
					_ = _index138
					(_1004_v1).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F2(), _dafny.IntOfUint32((_1006_v2).Cardinality()))).Cardinality()).Times((_this).F8()), (_index138).Int())
					var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_1004_v1), 0))
					_ = _index139
					(_1004_v1).ArraySet1((_this).F8(), (_index139).Int())
					var _1007_v3 _dafny.CodePoint
					_ = _1007_v3
					_1007_v3 = _dafny.CodePoint('t')
					var _1008_v4 D2
					_ = _1008_v4
					_1008_v4 = Companion_D2_.Create_DC6_(_1007_v3, (_this).F8())
					var _1009_v5 _dafny.Sequence
					_ = _1009_v5
					_1009_v5 = _dafny.SeqOf(_1008_v4, _1008_v4, Companion_D2_.Create_DC6_(_1007_v3, (_this).Fm3(_this.F7, p0, (_this).F8(), _dafny.IntOfUint32((_1006_v2).Cardinality()), globalState)))
					var _1010_v6 _dafny.Sequence
					_ = _1010_v6
					_1010_v6 = _dafny.SeqOf(_1009_v5)
					var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))
					_ = _index140
					var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_1004_v1), 0))
					_ = _index141
					var _rhs105 _dafny.Int = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1007_v3, _1004_v1)).Cardinality()).Plus((_this).F8())).Times((_this).F8())
					_ = _rhs105
					var _rhs106 _dafny.Int = (_this).F8()
					_ = _rhs106
					var _rhs107 _dafny.Sequence = _1010_v6
					_ = _rhs107
					var _lhs62 _dafny.Array = _1004_v1
					_ = _lhs62
					var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))
					_ = _lhs63
					var _lhs64 _dafny.Array = _1004_v1
					_ = _lhs64
					var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_1004_v1), 0))
					_ = _lhs65
					(_lhs62).ArraySet1(_rhs105, (_lhs63).Int())
					(_lhs64).ArraySet1(_rhs106, (_lhs65).Int())
					_1010_v6 = _rhs107
					var _1011_v7 _dafny.MultiSet
					_ = _1011_v7
					_1011_v7 = _dafny.MultiSetOf(_dafny.IntOfInt64(818))
					var _1012_v8 _dafny.Map
					_ = _1012_v8
					_1012_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1011_v7).Cardinality(), !((_this).F4()))
					var _1013_v9 _dafny.Set
					_ = _1013_v9
					_1013_v9 = _dafny.SetOf(_this.F7)
					var _1014_v10 _dafny.Map
					_ = _1014_v10
					_1014_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1004_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))).Int()).(_dafny.Int), _1013_v9)
					var _rhs108 bool = !(((_1012_v8).Cardinality()).Cmp(_dafny.IntOfUint32((_1006_v2).Cardinality())) >= 0) || (!(_this.F7) || (p0))
					_ = _rhs108
					var _rhs109 bool = (_this.F7) && (((_1004_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((_1014_v10).Cardinality())) == 0)
					_ = _rhs109
					var _lhs66 *C2 = _this
					_ = _lhs66
					var _lhs67 *C2 = _this
					_ = _lhs67
					_lhs66.F7 = _rhs108
					_lhs67.F7 = _rhs109
					var _1015_v11 _dafny.Set
					_ = _1015_v11
					_1015_v11 = _dafny.SetOf(_1006_v2)
					var _arr20 _dafny.Array = _this.F2()
					_ = _arr20
					var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_this.F2()), 0))
					_ = _index142
					(_arr20).ArraySet1(!(_1015_v11).Contains(_1006_v2), (_index142).Int())
					var _arr21 _dafny.Array = _this.F2()
					_ = _arr21
					var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_this.F2()), 0))
					_ = _index143
					(_arr21).ArraySet1((_this).F4(), (_index143).Int())
					var _1016_v12 D7
					_ = _1016_v12
					_1016_v12 = Companion_D7_.Create_DC24_(p0)
					var _rhs110 D7 = _1016_v12
					_ = _rhs110
					var _rhs111 bool = (func() bool {
						if (_1012_v8).Contains(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_1004_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))).Int()).(_dafny.Int)), (_1004_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))).Int()).(_dafny.Int))) {
							return (_1012_v8).Get(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_1004_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))).Int()).(_dafny.Int)), (_1004_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1004_v1), 0))).Int()).(_dafny.Int))).(bool)
						}
						return (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)
					})()
					_ = _rhs111
					var _rhs112 _dafny.Sequence = _1006_v2
					_ = _rhs112
					_1016_v12 = _rhs110
					r1 = _rhs111
					_1006_v2 = _rhs112
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _1017_v13 *C1
		_ = _1017_v13
		var _nw138 *C1 = New_C1_()
		_ = _nw138
		_nw138.Ctor__((_this).F3(), (_this.F7) == (_this.F7), _this.F2())
		_1017_v13 = _nw138
		for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_998_v0), 0))); ; {
			_guard_loop_6, _ok50 := _iter50()
			if !_ok50 {
				break
			}
			var _1018_i6 _dafny.Int
			_1018_i6 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_1018_i6).Sign() != -1) && ((_1018_i6).Cmp(_dafny.ArrayLen((_998_v0), 0)) < 0)) {
				(_998_v0).ArraySet1((func() _dafny.Sequence {
					if false {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("koiruaywh"), _dafny.UnicodeSeqOfUtf8Bytes("m"))
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(533))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}(func(_1019_i7 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('e')
					}))
				})(), (_1018_i6).Int())
			}
		}
		var _hi7 _dafny.Int = (_this).F8()
		_ = _hi7
		for _1020_i8 := (_dafny.IntOfInt64(556)).Plus((_this).F8()); _1020_i8.Cmp(_hi7) < 0; _1020_i8 = _1020_i8.Plus(_dafny.One) {
			var _1021_v14 _dafny.Map
			_ = _1021_v14
			_1021_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			_1021_v14 = (_1021_v14).Update((_this).F4(), _this.F7)
			var _1022_v15 _dafny.Int
			_ = _1022_v15
			_1022_v15 = _dafny.IntOfInt64(727)
			var _1023_v16 _dafny.Array
			_ = _1023_v16
			var _nw139 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(13))
			_ = _nw139
			_1023_v16 = _nw139
			var _1024_v17 _dafny.MultiSet
			_ = _1024_v17
			_1024_v17 = _dafny.MultiSetOf(_1023_v16, _1023_v16, _1023_v16)
			_1022_v15 = (_1024_v17).Cardinality()
			var _1025_v18 _dafny.Array
			_ = _1025_v18
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_32
			var _nw140 _dafny.Array
			_ = _nw140
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw140 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Map = func(_1026_i9 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("i"), _dafny.IntOfInt64(668))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(442))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}(func(_1027_i10 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('y')
					})), _dafny.IntOfInt64(906)))
				}
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw140 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw140).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw140).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1025_v18 = _nw140
			var _1028_v19 _dafny.Sequence
			_ = _1028_v19
			_1028_v19 = _dafny.UnicodeSeqOfUtf8Bytes("acwigaa")
			var _1029_v20 _dafny.Map
			_ = _1029_v20
			_1029_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1028_v19, _dafny.IntOfInt64(48))
			var _1030_v21 D8
			_ = _1030_v21
			_1030_v21 = Companion_D8_.Create_DC25_(_1029_v20)
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1025_v18), 0))
			_ = _index144
			(_1025_v18).ArraySet1(((_1030_v21).Dtor_cf44()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1028_v19, _1022_v15)), (_index144).Int())
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1025_v18), 0))
			_ = _index145
			var _rhs113 _dafny.Map = (_1029_v20).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(252))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}(func(_1031_i11 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), (_dafny.Zero).Minus((_this).Fm2(globalState)))
			_ = _rhs113
			var _rhs114 bool = p0
			_ = _rhs114
			var _lhs68 _dafny.Array = _1025_v18
			_ = _lhs68
			var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1025_v18), 0))
			_ = _lhs69
			(_lhs68).ArraySet1(_rhs113, (_lhs69).Int())
			r1 = _rhs114
			var _1032_v22 *C1
			_ = _1032_v22
			var _nw141 *C1 = New_C1_()
			_ = _nw141
			_nw141.Ctor__((_this).F3(), (_this).F4(), _this.F2())
			_1032_v22 = _nw141
		}
		var _1033_v23 _dafny.Sequence
		_ = _1033_v23
		_1033_v23 = _dafny.UnicodeSeqOfUtf8Bytes("apes")
		var _1034_v24 D0
		_ = _1034_v24
		_1034_v24 = Companion_D0_.Create_DC2_((_this).F8(), (_this).F4(), !((_this).F4()), (_1017_v13).Fm3(p0, false, _dafny.IntOfUint32((_1033_v23).Cardinality()), (_dafny.Zero).Minus((_this).F8()), globalState))
		(_this).M5((_1034_v24).Dtor_cf6(), p0, globalState)
		var _1035_v25 _dafny.Map
		_ = _1035_v25
		_1035_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1033_v23, (_this).F8())
		var _1036_v27 _dafny.CodePoint
		_ = _1036_v27
		_1036_v27 = _dafny.CodePoint('t')
		var _1037_v28 _dafny.Set
		_ = _1037_v28
		_1037_v28 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1033_v23, _1033_v23), (Companion_Default___.SafeIndex((func() _dafny.Set {
			var _coll44 = _dafny.NewBuilder()
			_ = _coll44
			for _iter51 := _dafny.Iterate((_1035_v25).Keys().Elements()); ; {
				_compr_44, _ok51 := _iter51()
				if !_ok51 {
					break
				}
				var _1038_v26 _dafny.Sequence
				_1038_v26 = interface{}(_compr_44).(_dafny.Sequence)
				if (_1035_v25).Contains(_1038_v26) {
					_coll44.Add(_1038_v26)
				}
			}
			return _coll44.ToSet()
		}()).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1033_v23, _1033_v23)).Cardinality()))).Uint32(), _1036_v27), _1033_v23, _1033_v23, _dafny.UnicodeSeqOfUtf8Bytes("uwwdmx"), _dafny.Companion_Sequence_.Concatenate(_1033_v23, _1033_v23))
		r0 = _1037_v28
		r1 = p0
		return r0, r1
	}
}
func (_this *C2) M5(p0 bool, p1 bool, globalState *GlobalState) {
	{
		if p1 {
			var _1039_v0 _dafny.Sequence
			_ = _1039_v0
			_1039_v0 = _dafny.UnicodeSeqOfUtf8Bytes("almht")
			_1039_v0 = _1039_v0
			var _1040_v1 D5
			_ = _1040_v1
			_1040_v1 = Companion_D5_.Create_DC17_((_this).F8())
			var _1041_v2 _dafny.Map
			_ = _1041_v2
			_1041_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_1040_v1).Dtor_cf35())
			var _1042_v3 _dafny.Int
			_ = _1042_v3
			_1042_v3 = _dafny.IntOfInt64(-992)
			var _rhs115 bool = (_1042_v3).Cmp((_this).F8()) >= 0
			_ = _rhs115
			var _rhs116 _dafny.Map = _1041_v2
			_ = _rhs116
			var _rhs117 bool = ((_this).F8()).Cmp(_1042_v3) != 0
			_ = _rhs117
			var _rhs118 _dafny.Int = _dafny.IntOfInt64(984)
			_ = _rhs118
			var _lhs70 *C2 = _this
			_ = _lhs70
			var _lhs71 *C2 = _this
			_ = _lhs71
			_lhs70.F7 = _rhs115
			_1041_v2 = _rhs116
			_lhs71.F7 = _rhs117
			_1042_v3 = _rhs118
			var _arr22 _dafny.Array = _this.F2()
			_ = _arr22
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index146
			(_arr22).ArraySet1((_this).Fm12(globalState), (_index146).Int())
			var _1043_v4 _dafny.Map
			_ = _1043_v4
			_1043_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), p1)
			var _1044_v5 _dafny.Set
			_ = _1044_v5
			_1044_v5 = _dafny.SetOf(p0, p0, (_this).F4(), (func() bool {
				if (_1043_v4).Contains((_this).F4()) {
					return (_1043_v4).Get((_this).F4()).(bool)
				}
				return (_this).F4()
			})())
			var _1045_v6 _dafny.MultiSet
			_ = _1045_v6
			_1045_v6 = _dafny.MultiSetOf((_this).F8(), (_1044_v5).Cardinality())
			var _1046_v7 _dafny.MultiSet
			_ = _1046_v7
			_1046_v7 = _dafny.MultiSetOf(p0, _this.F7)
			var _1047_v8 _dafny.Sequence
			_ = _1047_v8
			_1047_v8 = _dafny.SeqOf(_1046_v7)
			var _arr23 _dafny.Array = _this.F2()
			_ = _arr23
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index147
			(_arr23).ArraySet1(!(!((_1042_v3).Cmp((func() _dafny.Int {
				if (_1045_v6).Contains(_1042_v3) {
					return (_1045_v6).Multiplicity(_1042_v3)
				}
				return _dafny.IntOfUint32((_1047_v8).Cardinality())
			})()) > 0)) || (!(_this.F7) || (p0)), (_index147).Int())
			if _this.F7 {
				var _1048_v9 _dafny.Array
				_ = _1048_v9
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_33
				var _nw142 _dafny.Array
				_ = _nw142
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw142 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.Map = func(_1049_i0 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('s'))
					}
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw142 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw142).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw142).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1048_v9 = _nw142
				var _1050_v10 _dafny.Map
				_ = _1050_v10
				_1050_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), _dafny.CodePoint('x'))
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1048_v9), 0))
				_ = _index148
				(_1048_v9).ArraySet1(_1050_v10, (_index148).Int())
				var _1051_v11 _dafny.Array
				_ = _1051_v11
				var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw143
				_1051_v11 = _nw143
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))
				_ = _index149
				(_1051_v11).ArraySet1((_this).F8(), (_index149).Int())
				var _1052_v12 _dafny.Set
				_ = _1052_v12
				_1052_v12 = _dafny.SetOf(Companion_Default___.Fm27(_dafny.IntOfInt64(60), _1042_v3, globalState))
				var _1053_v13 _dafny.CodePoint
				_ = _1053_v13
				_1053_v13 = _dafny.CodePoint('o')
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1048_v9), 0))
				_ = _index150
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))
				_ = _index151
				var _arr24 _dafny.Array = _this.F2()
				_ = _arr24
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index152
				var _rhs119 _dafny.Map = ((_1050_v10).Merge(_1050_v10)).Update((_1052_v12).IsDisjointFrom(_1052_v12), _1053_v13)
				_ = _rhs119
				var _rhs120 bool = _this.F7
				_ = _rhs120
				var _rhs121 _dafny.Int = _1042_v3
				_ = _rhs121
				var _rhs122 bool = (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)
				_ = _rhs122
				var _lhs72 _dafny.Array = _1048_v9
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1048_v9), 0))
				_ = _lhs73
				var _lhs74 *C2 = _this
				_ = _lhs74
				var _lhs75 _dafny.Array = _1051_v11
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))
				_ = _lhs76
				var _lhs77 _dafny.Array = _this.F2()
				_ = _lhs77
				var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs78
				(_lhs72).ArraySet1(_rhs119, (_lhs73).Int())
				_lhs74.F7 = _rhs120
				(_lhs75).ArraySet1(_rhs121, (_lhs76).Int())
				(_lhs77).ArraySet1(_rhs122, (_lhs78).Int())
				var _1054_v14 _dafny.Array
				_ = _1054_v14
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_34
				var _nw144 _dafny.Array
				_ = _nw144
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw144 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Sequence = (func(_1055_v3 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_1056_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1055_v3, (_this).F8()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(122))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg49 _dafny.Int) interface{} {
									return coer49(arg49)
								}
							}((func(_1057_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1058_i2 _dafny.Int) _dafny.Int {
									return _1057_v3
								}
							})(_1055_v3))))
						}
					})(_1042_v3)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw144 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw144).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw144).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1054_v14 = _nw144
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_1054_v14), 0))
				_ = _index153
				(_1054_v14).ArraySet1(_dafny.SeqOf((_this).F8(), (_1051_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))).Int()).(_dafny.Int)), (_index153).Int())
				var _1059_v15 _dafny.Array
				_ = _1059_v15
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_35
				var _nw145 _dafny.Array
				_ = _nw145
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw145 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) D4 = func(_1060_i3 _dafny.Int) D4 {
						return Companion_D4_.Create_DC11_(_dafny.SeqOf(!((_this).F4())))
					}
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw145 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw145).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw145).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1059_v15 = _nw145
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_1059_v15), 0))
				_ = _index154
				(_1059_v15).ArraySet1(Companion_Default___.Fm19(p0, (_1051_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))).Int()).(_dafny.Int), globalState), (_index154).Int())
				var _1061_v16 _dafny.Sequence
				_ = _1061_v16
				_1061_v16 = _dafny.SeqOf(p1, !(!(p0)), _this.F7, (_this).F4(), true)
				var _1062_v17 D4
				_ = _1062_v17
				_1062_v17 = Companion_D4_.Create_DC11_(_1061_v16)
				var _arr25 _dafny.Array = _this.F2()
				_ = _arr25
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index155
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))
				_ = _index156
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_1054_v14), 0))
				_ = _index157
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_1059_v15), 0))
				_ = _index158
				var _rhs123 bool = _this.F7
				_ = _rhs123
				var _rhs124 _dafny.Int = _1042_v3
				_ = _rhs124
				var _rhs125 _dafny.Sequence = _dafny.SeqOf((_1045_v6).Cardinality(), _dafny.IntOfUint32((_1061_v16).Cardinality()), _1042_v3, (_1051_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))).Int()).(_dafny.Int))
				_ = _rhs125
				var _rhs126 D4 = _1062_v17
				_ = _rhs126
				var _lhs79 _dafny.Array = _this.F2()
				_ = _lhs79
				var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs80
				var _lhs81 _dafny.Array = _1051_v11
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))
				_ = _lhs82
				var _lhs83 _dafny.Array = _1054_v14
				_ = _lhs83
				var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(493), _dafny.ArrayLen((_1054_v14), 0))
				_ = _lhs84
				var _lhs85 _dafny.Array = _1059_v15
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_1059_v15), 0))
				_ = _lhs86
				(_lhs79).ArraySet1(_rhs123, (_lhs80).Int())
				(_lhs81).ArraySet1(_rhs124, (_lhs82).Int())
				(_lhs83).ArraySet1(_rhs125, (_lhs84).Int())
				(_lhs85).ArraySet1(_rhs126, (_lhs86).Int())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))
				_ = _index159
				(_1051_v11).ArraySet1(_1042_v3, (_index159).Int())
				var _1063_v18 D2
				_ = _1063_v18
				_1063_v18 = Companion_D2_.Create_DC5_(_1039_v0)
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))
				_ = _index160
				(_1051_v11).ArraySet1(_dafny.IntOfUint32(((_1063_v18).Dtor_cf13()).Cardinality()), (_index160).Int())
				var _1064_v19 _dafny.Map
				_ = _1064_v19
				_1064_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_1051_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))).Int()).(_dafny.Int))
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))
				_ = _index161
				(_1051_v11).ArraySet1((func() _dafny.Int {
					if (_1064_v19).Contains((_this).F4()) {
						return (_1064_v19).Get((_this).F4()).(_dafny.Int)
					}
					return (_1051_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_1051_v11), 0))).Int()).(_dafny.Int)
				})(), (_index161).Int())
			} else {
				(_this).F7 = p0
				var _1065_v20 _dafny.CodePoint
				_ = _1065_v20
				_1065_v20 = _dafny.CodePoint('i')
				(_this).F7 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-936))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_1066_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1067_i4 _dafny.Int) _dafny.CodePoint {
						return _1066_v20
					}
				})(_1065_v20))), (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-936))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_1068_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1069_i4 _dafny.Int) _dafny.CodePoint {
						return _1068_v20
					}
				})(_1065_v20)))).Cardinality()))).Uint32(), _1065_v20), _1065_v20)
				var _1070_v21 _dafny.Map
				_ = _1070_v21
				_1070_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1065_v20, _1042_v3)
				_1070_v21 = _1070_v21
				var _1071_v22 _dafny.Map
				_ = _1071_v22
				_1071_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1042_v3, (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
				var _1072_v23 _dafny.Map
				_ = _1072_v23
				_1072_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), Companion_Default___.Fm18(p0, (_1071_v22).Update((_this).F8(), p1), globalState))
				_1039_v0 = (func() _dafny.Sequence {
					if (_1072_v23).Contains((_this).F4()) {
						return (_1072_v23).Get((_this).F4()).(_dafny.Sequence)
					}
					return _1039_v0
				})()
				var _1073_v24 _dafny.Array
				_ = _1073_v24
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw146
				_1073_v24 = _nw146
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1073_v24), 0))
				_ = _index162
				(_1073_v24).ArraySet1(_1042_v3, (_index162).Int())
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1073_v24), 0))
				_ = _index163
				(_1073_v24).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(688), (_this).F8()), (_index163).Int())
			}
			var _1074_v25 _dafny.Array
			_ = _1074_v25
			var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
			_ = _nw147
			_1074_v25 = _nw147
			var _1075_v26 _dafny.Sequence
			_ = _1075_v26
			_1075_v26 = _dafny.SeqOf(_1042_v3)
			var _1076_v27 _dafny.Array
			_ = _1076_v27
			var _nwElement0_28 _dafny.Int = (_dafny.Zero).Minus((_this).F8())
			_ = _nwElement0_28
			var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(24))
			_ = _nw148
			(_nw148).ArraySet1(_nwElement0_28, 0)
			(_nw148).ArraySet1((_this).F8(), 1)
			(_nw148).ArraySet1(_dafny.IntOfInt64(387), 2)
			(_nw148).ArraySet1(_dafny.IntOfUint32((_1039_v0).Cardinality()), 3)
			(_nw148).ArraySet1((_this).F8(), 4)
			(_nw148).ArraySet1(_dafny.IntOfUint32((_1039_v0).Cardinality()), 5)
			(_nw148).ArraySet1((_this).F8(), 6)
			(_nw148).ArraySet1((_this).F8(), 7)
			(_nw148).ArraySet1(_dafny.IntOfInt64(-572), 8)
			(_nw148).ArraySet1(_1042_v3, 9)
			(_nw148).ArraySet1((_this).F8(), 10)
			(_nw148).ArraySet1(_1042_v3, 11)
			(_nw148).ArraySet1((_this).F8(), 12)
			(_nw148).ArraySet1(_1042_v3, 13)
			(_nw148).ArraySet1(_1042_v3, 14)
			(_nw148).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("soi")).Cardinality()), 15)
			(_nw148).ArraySet1(_1042_v3, 16)
			(_nw148).ArraySet1(_dafny.IntOfUint32((_1075_v26).Cardinality()), 17)
			(_nw148).ArraySet1((_this).F8(), 18)
			(_nw148).ArraySet1((_this).F8(), 19)
			(_nw148).ArraySet1((_this).F8(), 20)
			(_nw148).ArraySet1((_this).Fm3(p1, p0, _dafny.IntOfInt64(486), (_this).F8(), globalState), 21)
			(_nw148).ArraySet1(_1042_v3, 22)
			(_nw148).ArraySet1(_1042_v3, 23)
			_1076_v27 = _nw148
			var _1077_v28 _dafny.Array
			_ = _1077_v28
			_1077_v28 = _1076_v27
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1074_v25), 0))
			_ = _index164
			(_1074_v25).ArraySet1((_1077_v28), (_index164).Int())
			var _1078_v29 _dafny.Sequence
			_ = _1078_v29
			_1078_v29 = _dafny.SeqOf(_1076_v27, _1076_v27)
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1074_v25), 0))
			_ = _index165
			(_1074_v25).ArraySet1((func() _dafny.Array {
				if ((_this).F8()).Cmp((_dafny.Zero).Minus((_1046_v7).Cardinality())) != 0 {
					return (_1078_v29).Select((Companion_Default___.SafeIndex(_1042_v3, _dafny.IntOfUint32((_1078_v29).Cardinality()))).Uint32()).(_dafny.Array)
				}
				return _1076_v27
			})(), (_index165).Int())
		} else {
			var _1079_v30 D8
			_ = _1079_v30
			_1079_v30 = Companion_D8_.Create_DC27_(!(p1), !(p1), (_dafny.Zero).Minus((_this).F8()), (_this).F8())
			var _1080_v31 _dafny.Map
			_ = _1080_v31
			_1080_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1079_v30, (_this).F8())
			var _1081_v32 _dafny.Set
			_ = _1081_v32
			_1081_v32 = _dafny.SetOf((_this).F8())
			_1080_v31 = (_1080_v31).Update(_1079_v30, (_1081_v32).Cardinality())
			var _1082_v33 _dafny.Int
			_ = _1082_v33
			_1082_v33 = _dafny.IntOfInt64(544)
			var _1083_v34 _dafny.Map
			_ = _1083_v34
			_1083_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), !(_this.F7))
			var _1084_v35 _dafny.Map
			_ = _1084_v35
			_1084_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F8())))
			var _arr26 _dafny.Array = _this.F2()
			_ = _arr26
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index166
			(_arr26).ArraySet1((_1084_v35).Contains((func() bool {
				if (_1083_v34).Contains(false) {
					return (_1083_v34).Get(false).(bool)
				}
				return (_this).F4()
			})()), (_index166).Int())
			var _1085_v36 _dafny.MultiSet
			_ = _1085_v36
			_1085_v36 = _dafny.MultiSetOf((_this).F8())
			var _arr27 _dafny.Array = _this.F2()
			_ = _arr27
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index167
			var _rhs127 _dafny.Int = _1082_v33
			_ = _rhs127
			var _rhs128 bool = (_this).F4()
			_ = _rhs128
			var _rhs129 bool = (_1085_v36).IsSubsetOf(_1085_v36)
			_ = _rhs129
			var _rhs130 bool = (_this).Fm12(globalState)
			_ = _rhs130
			var _lhs87 *C2 = _this
			_ = _lhs87
			var _lhs88 *C2 = _this
			_ = _lhs88
			var _lhs89 _dafny.Array = _this.F2()
			_ = _lhs89
			var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))
			_ = _lhs90
			_1082_v33 = _rhs127
			_lhs87.F7 = _rhs128
			_lhs88.F7 = _rhs129
			(_lhs89).ArraySet1(_rhs130, (_lhs90).Int())
			var _1086_v37 _dafny.Sequence
			_ = _1086_v37
			_1086_v37 = _dafny.SeqOf(_1082_v33, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-202))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}(func(_1087_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))).Cardinality()), (_this).F8())
			var _1088_v38 _dafny.Sequence
			_ = _1088_v38
			_1088_v38 = _dafny.UnicodeSeqOfUtf8Bytes("cqil")
			var _1089_v39 _dafny.Array
			_ = _1089_v39
			var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw149
			_1089_v39 = _nw149
			var _1090_v40 _dafny.CodePoint
			_ = _1090_v40
			_1090_v40 = _dafny.CodePoint('j')
			var _1091_v41 _dafny.Sequence
			_ = _1091_v41
			_1091_v41 = _dafny.SeqOf(_1089_v39)
			var _1092_v42 _dafny.MultiSet
			_ = _1092_v42
			_1092_v42 = _dafny.MultiSetOf((_this).F4(), (_this).F4(), (_this).F4())
			var _1093_v43 _dafny.Sequence
			_ = _1093_v43
			_1093_v43 = _dafny.SeqOf(p0, true)
			var _1094_v44 _dafny.Map
			_ = _1094_v44
			_1094_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_1093_v43)).Cardinality()), (_this).F8())
			var _rhs131 _dafny.Sequence = _dafny.SeqOf((_this).F8(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1088_v38).Cardinality())), Companion_Default___.SafeDivisionInt((_this).F8(), (_this).F8()))
			_ = _rhs131
			var _rhs132 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1088_v38, (Companion_Default___.SafeIndex(_1082_v33, _dafny.IntOfUint32((_1088_v38).Cardinality()))).Uint32(), _1090_v40), _dafny.Companion_Sequence_.Concatenate(_1088_v38, _1088_v38))
			_ = _rhs132
			var _rhs133 _dafny.Array = (_1091_v41).Select((Companion_Default___.SafeIndex(((_this).F8()).Plus(_1082_v33), _dafny.IntOfUint32((_1091_v41).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs133
			var _rhs134 _dafny.Int = (((_1092_v42).Union(_1092_v42)).Cardinality()).Plus((func() _dafny.Int {
				if (_1094_v44).Contains((_1084_v35).Cardinality()) {
					return (_1094_v44).Get((_1084_v35).Cardinality()).(_dafny.Int)
				}
				return _1082_v33
			})())
			_ = _rhs134
			var _rhs135 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1088_v38, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("b"), _1088_v38))).Cardinality())
			_ = _rhs135
			_1086_v37 = _rhs131
			_1088_v38 = _rhs132
			_1089_v39 = _rhs133
			_1082_v33 = _rhs134
			_1082_v33 = _rhs135
			var _arr28 _dafny.Array = _this.F2()
			_ = _arr28
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index168
			(_arr28).ArraySet1(((_1093_v43).Select((Companion_Default___.SafeIndex(_1082_v33, _dafny.IntOfUint32((_1093_v43).Cardinality()))).Uint32()).(bool)) || (p0), (_index168).Int())
			if !(p1) {
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1089_v39), 0))
				_ = _index169
				(_1089_v39).ArraySet1(_1082_v33, (_index169).Int())
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1089_v39), 0))
				_ = _index170
				var _arr29 _dafny.Array = _this.F2()
				_ = _arr29
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index171
				var _rhs136 _dafny.Int = Companion_Default___.Fm0(_1082_v33, globalState)
				_ = _rhs136
				var _rhs137 _dafny.Map = _1094_v44
				_ = _rhs137
				var _rhs138 bool = (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)
				_ = _rhs138
				var _lhs91 _dafny.Array = _1089_v39
				_ = _lhs91
				var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1089_v39), 0))
				_ = _lhs92
				var _lhs93 _dafny.Array = _this.F2()
				_ = _lhs93
				var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs94
				(_lhs91).ArraySet1(_rhs136, (_lhs92).Int())
				_1094_v44 = _rhs137
				(_lhs93).ArraySet1(_rhs138, (_lhs94).Int())
				_1084_v35 = (_1084_v35).Update(!(false) || ((_this).F4()), Companion_Default___.SafeDivisionInt((_1086_v37).Select((Companion_Default___.SafeIndex((_1089_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1089_v39), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1086_v37).Cardinality()))).Uint32()).(_dafny.Int), (_this).F8()))
				_1083_v34 = (_1083_v34).Update((_this).F4(), (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
				var _1095_v45 D2
				_ = _1095_v45
				_1095_v45 = Companion_D2_.Create_DC5_(_1088_v38)
				_1088_v38 = (_1095_v45).Dtor_cf13()
				var _1096_v46 _dafny.Array
				_ = _1096_v46
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_36
				var _nw150 _dafny.Array
				_ = _nw150
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw150 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.MultiSet = (func(_1097_v39 _dafny.Array) func(_dafny.Int) _dafny.MultiSet {
						return func(_1098_i6 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf((_1097_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1097_v39), 0))).Int()).(_dafny.Int))
						}
					})(_1089_v39)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw150 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw150).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw150).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1096_v46 = _nw150
				var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(20))
				_ = _nw151
				_1096_v46 = _nw151
			} else {
				(_this).F7 = !(!(p0) || (_this.F7))
				_1082_v33 = _1082_v33
				_1088_v38 = _dafny.Companion_Sequence_.Concatenate(_1088_v38, _1088_v38)
				_1090_v40 = _1090_v40
				var _1099_v47 D8
				_ = _1099_v47
				_1099_v47 = Companion_D8_.Create_DC26_(p0, true, false, (_this).F8())
				var _1100_v48 _dafny.Map
				_ = _1100_v48
				_1100_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), p0)
				var _pat_let_tv10 = _1100_v48
				_ = _pat_let_tv10
				var _1101_v49 _dafny.Array
				_ = _1101_v49
				var _nwElement0_29 D8 = _1099_v47
				_ = _nwElement0_29
				var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(11))
				_ = _nw152
				(_nw152).ArraySet1(_nwElement0_29, 0)
				(_nw152).ArraySet1(func(_pat_let17_0 D8) D8 {
					return func(_1102_dt__update__tmp_h0 D8) D8 {
						return func(_pat_let18_0 _dafny.Int) D8 {
							return func(_1103_dt__update_hcf48_h0 _dafny.Int) D8 {
								return func(_pat_let19_0 bool) D8 {
									return func(_1104_dt__update_hcf46_h0 bool) D8 {
										return Companion_D8_.Create_DC26_((_1102_dt__update__tmp_h0).Dtor_cf45(), _1104_dt__update_hcf46_h0, (_1102_dt__update__tmp_h0).Dtor_cf47(), _1103_dt__update_hcf48_h0)
									}(_pat_let19_0)
								}(_this.F7)
							}(_pat_let18_0)
						}((_pat_let_tv10).Cardinality())
					}(_pat_let17_0)
				}(Companion_D8_.Create_DC26_(p0, (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), _this.F7, _1082_v33)), 1)
				(_nw152).ArraySet1(_1099_v47, 2)
				(_nw152).ArraySet1(Companion_Default___.Fm28((_this).F8(), (_1086_v37).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1086_v37).Cardinality()))).Uint32()).(_dafny.Int), globalState), 3)
				(_nw152).ArraySet1(Companion_D8_.Create_DC26_(p1, _this.F7, false, _1082_v33), 4)
				(_nw152).ArraySet1(_1099_v47, 5)
				(_nw152).ArraySet1(func(_pat_let20_0 D8) D8 {
					return func(_1105_dt__update__tmp_h1 D8) D8 {
						return func(_pat_let21_0 bool) D8 {
							return func(_1106_dt__update_hcf45_h0 bool) D8 {
								return Companion_D8_.Create_DC26_(_1106_dt__update_hcf45_h0, (_1105_dt__update__tmp_h1).Dtor_cf46(), (_1105_dt__update__tmp_h1).Dtor_cf47(), (_1105_dt__update__tmp_h1).Dtor_cf48())
							}(_pat_let21_0)
						}((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
					}(_pat_let20_0)
				}(_1099_v47), 6)
				(_nw152).ArraySet1(_1099_v47, 7)
				(_nw152).ArraySet1(Companion_D8_.Create_DC26_(p1, true, true, (_this).F8()), 8)
				(_nw152).ArraySet1(_1099_v47, 9)
				(_nw152).ArraySet1(Companion_D8_.Create_DC26_(!((_this).F4()), (_this).F4(), (_this).Fm12(globalState), _1082_v33), 10)
				_1101_v49 = _nw152
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1101_v49), 0))
				_ = _index172
				(_1101_v49).ArraySet1(_1099_v47, (_index172).Int())
				var _pat_let_tv11 = _1082_v33
				_ = _pat_let_tv11
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_1101_v49), 0))
				_ = _index173
				(_1101_v49).ArraySet1(func(_pat_let22_0 D8) D8 {
					return func(_1107_dt__update__tmp_h2 D8) D8 {
						return func(_pat_let23_0 _dafny.Int) D8 {
							return func(_1108_dt__update_hcf48_h1 _dafny.Int) D8 {
								return Companion_D8_.Create_DC26_((_1107_dt__update__tmp_h2).Dtor_cf45(), (_1107_dt__update__tmp_h2).Dtor_cf46(), (_1107_dt__update__tmp_h2).Dtor_cf47(), _1108_dt__update_hcf48_h1)
							}(_pat_let23_0)
						}(_pat_let_tv11)
					}(_pat_let22_0)
				}(Companion_D8_.Create_DC26_((_this).Fm12(globalState), true, p0, _1082_v33)), (_index173).Int())
			}
		}
		var _rhs139 bool = (((_this).F8()).Times((_this).F8())).Cmp((_this).F8()) >= 0
		_ = _rhs139
		var _rhs140 bool = p0
		_ = _rhs140
		var _lhs95 *C2 = _this
		_ = _lhs95
		var _lhs96 *C2 = _this
		_ = _lhs96
		_lhs95.F7 = _rhs139
		_lhs96.F7 = _rhs140
		var _1109_v50 _dafny.Sequence
		_ = _1109_v50
		_1109_v50 = _dafny.UnicodeSeqOfUtf8Bytes("jwlocp")
		var _1110_i7 _dafny.Int
		_ = _1110_i7
		_1110_i7 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_1109_v50, _1109_v50) {
				{
					if (_1110_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1110_i7 = (_1110_i7).Plus(_dafny.One)
					_1109_v50 = _dafny.Companion_Sequence_.Concatenate(_1109_v50, _1109_v50)
					var _1111_v51 _dafny.Set
					_ = _1111_v51
					_1111_v51 = _dafny.SetOf((_this).F8())
					if (_1111_v51).IsDisjointFrom(func() _dafny.Set {
						var _coll45 = _dafny.NewBuilder()
						_ = _coll45
						for _iter52 := _dafny.Iterate((Companion_Default___.Fm29(globalState)).Elements()); ; {
							_compr_45, _ok52 := _iter52()
							if !_ok52 {
								break
							}
							var _1112_v52 _dafny.Int
							_1112_v52 = interface{}(_compr_45).(_dafny.Int)
							if (Companion_Default___.Fm29(globalState)).Contains(_1112_v52) {
								_coll45.Add((_1112_v52).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), (Companion_D0_.Create_DC1_(true, true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Dtor_cf1())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), false)).Cardinality())).Cardinality()))
							}
						}
						return _coll45.ToSet()
					}()) {
						(_this).F7 = p0
						var _1113_v53 _dafny.Array
						_ = _1113_v53
						var _len0_37 _dafny.Int = _dafny.IntOfInt64(4)
						_ = _len0_37
						var _nw153 _dafny.Array
						_ = _nw153
						if _len0_37.Cmp(_dafny.Zero) == 0 {
							_nw153 = _dafny.NewArray(_len0_37)
						} else {
							var _init37 func(_dafny.Int) _dafny.Sequence = func(_1114_i8 _dafny.Int) _dafny.Sequence {
								return (Companion_D4_.Create_DC11_(_dafny.SeqOf(_this.F7))).Dtor_cf28()
							}
							_ = _init37
							var _element0_37 = _init37(_dafny.Zero)
							_ = _element0_37
							_nw153 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
							(_nw153).ArraySet1(_element0_37, 0)
							var _nativeLen0_37 = (_len0_37).Int()
							_ = _nativeLen0_37
							for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
								(_nw153).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
							}
						}
						_1113_v53 = _nw153
						_1113_v53 = _1113_v53
						var _1115_v54 _dafny.CodePoint
						_ = _1115_v54
						_1115_v54 = _dafny.CodePoint('c')
						var _1116_v55 _dafny.Map
						_ = _1116_v55
						_1116_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v54, (_this).F8())
						var _1117_v56 _dafny.Sequence
						_ = _1117_v56
						_1117_v56 = _dafny.SeqOf((_this).F8())
						var _1118_v57 _dafny.MultiSet
						_ = _1118_v57
						_1118_v57 = _dafny.MultiSetOf(_dafny.IntOfInt64(-168))
						var _1119_v58 _dafny.Map
						_ = _1119_v58
						_1119_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _this.F7)
						_1116_v55 = (_1116_v55).Update(_1115_v54, ((func() _dafny.MultiSet {
							if p1 {
								return _dafny.MultiSetFromSeq(_1117_v56)
							}
							return (_1118_v57).Update((_this).F8(), Companion_Default___.Abs((_1119_v58).Cardinality()))
						})()).Cardinality())
						(_this).F7 = p0
						(_this).F7 = (_this).F4()
					} else {
						var _1120_v59 _dafny.CodePoint
						_ = _1120_v59
						_1120_v59 = _dafny.CodePoint('v')
						_1120_v59 = _1120_v59
						var _1121_v60 T1
						_ = _1121_v60
						var _nw154 *C1 = New_C1_()
						_ = _nw154
						_nw154.Ctor__((_this).F3(), p0, _this.F2())
						_1121_v60 = _nw154
						var _1122_v61 _dafny.MultiSet
						_ = _1122_v61
						_1122_v61 = _dafny.MultiSetOf(_1121_v60)
						var _1123_v62 _dafny.Set
						_ = _1123_v62
						_1123_v62 = _dafny.SetOf(_1122_v61)
						var _1124_v63 _dafny.MultiSet
						_ = _1124_v63
						_1124_v63 = _dafny.MultiSetOf((_1123_v62).Cardinality())
						_1111_v51 = (_1111_v51).Intersection(func() _dafny.Set {
							var _coll46 = _dafny.NewBuilder()
							_ = _coll46
							for _iter53 := _dafny.Iterate((_1124_v63).Elements()); ; {
								_compr_46, _ok53 := _iter53()
								if !_ok53 {
									break
								}
								var _1125_v64 _dafny.Int
								_1125_v64 = interface{}(_compr_46).(_dafny.Int)
								if (_1124_v63).Contains(_1125_v64) {
									_coll46.Add((_1125_v64).Plus((_dafny.Zero).Minus(_dafny.IntOfInt64(-584))))
								}
							}
							return _coll46.ToSet()
						}())
						(_this).F7 = _this.F7
						var _1126_v65 _dafny.Array
						_ = _1126_v65
						var _len0_38 _dafny.Int = _dafny.IntOfInt64(10)
						_ = _len0_38
						var _nw155 _dafny.Array
						_ = _nw155
						if _len0_38.Cmp(_dafny.Zero) == 0 {
							_nw155 = _dafny.NewArray(_len0_38)
						} else {
							var _init38 func(_dafny.Int) _dafny.Int = func(_1127_i9 _dafny.Int) _dafny.Int {
								return (_1127_i9).Plus((_this).F8())
							}
							_ = _init38
							var _element0_38 = _init38(_dafny.Zero)
							_ = _element0_38
							_nw155 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
							(_nw155).ArraySet1(_element0_38, 0)
							var _nativeLen0_38 = (_len0_38).Int()
							_ = _nativeLen0_38
							for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
								(_nw155).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
							}
						}
						_1126_v65 = _nw155
						var _1128_v66 _dafny.Map
						_ = _1128_v66
						_1128_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wbtfe"), _1109_v50), _1126_v65)
						_1128_v66 = (_1128_v66).Merge(((_1128_v66).Update(_1109_v50, _1126_v65)).Update(_1109_v50, _1126_v65))
						(_this).F7 = !(((_this).F8()).Cmp(Companion_Default___.SafeDivisionInt((_this).F8(), (_this).F8())) < 0)
					}
					if !(p0) {
						var _1129_v67 _dafny.Int
						_ = _1129_v67
						_1129_v67 = _dafny.IntOfInt64(483)
						var _1130_v68 _dafny.Sequence
						_ = _1130_v68
						_1130_v68 = _dafny.SeqOf(_1129_v67)
						_1129_v67 = _dafny.IntOfUint32((_1130_v68).Cardinality())
						_1129_v67 = Companion_Default___.SafeModuloInt((Companion_Default___.Fm29(globalState)).Cardinality(), _dafny.IntOfInt64(471))
						var _1131_v69 _dafny.Map
						_ = _1131_v69
						_1131_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), Companion_D0_.Create_DC1_((_this).Fm12(globalState), p1, _1129_v67))
						var _1132_v70 D0
						_ = _1132_v70
						_1132_v70 = Companion_D0_.Create_DC1_(!(_this.F7), !(false), (_dafny.Zero).Minus(_1129_v67))
						_1131_v69 = (_1131_v69).Update(Companion_Default___.Fm0((_this).F8(), globalState), _1132_v70)
						_1129_v67 = _dafny.IntOfUint32((_1109_v50).Cardinality())
						var _rhs141 _dafny.Int = _1129_v67
						_ = _rhs141
						var _rhs142 _dafny.Int = (_this).F8()
						_ = _rhs142
						var _rhs143 bool = (_dafny.IntOfInt64(967)).Cmp((_this).F8()) > 0
						_ = _rhs143
						var _lhs97 *C2 = _this
						_ = _lhs97
						_1129_v67 = _rhs141
						_1129_v67 = _rhs142
						_lhs97.F7 = _rhs143
					} else {
						var _1133_v71 _dafny.CodePoint
						_ = _1133_v71
						_1133_v71 = _dafny.CodePoint('q')
						var _1134_v72 _dafny.MultiSet
						_ = _1134_v72
						_1134_v72 = _dafny.MultiSetOf(_1133_v71)
						var _1135_v73 D4
						_ = _1135_v73
						_1135_v73 = Companion_D4_.Create_DC13_(p1, _1134_v72, !(true))
						var _pat_let_tv12 = _1133_v71
						_ = _pat_let_tv12
						var _pat_let_tv13 = p0
						_ = _pat_let_tv13
						var _1136_v74 _dafny.Array
						_ = _1136_v74
						var _nwElement0_30 D4 = Companion_D4_.Create_DC13_(false, (_dafny.MultiSetOf(_1133_v71, _1133_v71, _1133_v71)).Update(_1133_v71, Companion_Default___.Abs(_dafny.IntOfInt64(505))), (_this).F4())
						_ = _nwElement0_30
						var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(18))
						_ = _nw156
						(_nw156).ArraySet1(_nwElement0_30, 0)
						(_nw156).ArraySet1(_1135_v73, 1)
						(_nw156).ArraySet1(_1135_v73, 2)
						(_nw156).ArraySet1(Companion_D4_.Create_DC13_(p1, _1134_v72, (_this).F4()), 3)
						(_nw156).ArraySet1(Companion_Default___.Fm30((_this).F8(), p1, false, p1, globalState), 4)
						(_nw156).ArraySet1(Companion_Default___.Fm30((_this).F8(), _this.F7, (_this).F4(), _this.F7, globalState), 5)
						(_nw156).ArraySet1(_1135_v73, 6)
						(_nw156).ArraySet1(_1135_v73, 7)
						(_nw156).ArraySet1(_1135_v73, 8)
						(_nw156).ArraySet1(_1135_v73, 9)
						(_nw156).ArraySet1(func(_pat_let24_0 D4) D4 {
							return func(_1137_dt__update__tmp_h3 D4) D4 {
								return func(_pat_let25_0 _dafny.MultiSet) D4 {
									return func(_1138_dt__update_hcf30_h0 _dafny.MultiSet) D4 {
										return func(_pat_let26_0 bool) D4 {
											return func(_1139_dt__update_hcf29_h0 bool) D4 {
												return Companion_D4_.Create_DC13_(_1139_dt__update_hcf29_h0, _1138_dt__update_hcf30_h0, (_1137_dt__update__tmp_h3).Dtor_cf31())
											}(_pat_let26_0)
										}(_pat_let_tv13)
									}(_pat_let25_0)
								}(_dafny.MultiSetOf(_pat_let_tv12))
							}(_pat_let24_0)
						}(Companion_D4_.Create_DC13_(p1, _dafny.MultiSetOf(_1133_v71), p1)), 10)
						(_nw156).ArraySet1(_1135_v73, 11)
						(_nw156).ArraySet1(_1135_v73, 12)
						(_nw156).ArraySet1(_1135_v73, 13)
						(_nw156).ArraySet1(_1135_v73, 14)
						(_nw156).ArraySet1(_1135_v73, 15)
						(_nw156).ArraySet1(func(_pat_let27_0 D4) D4 {
							return func(_1140_dt__update__tmp_h4 D4) D4 {
								return func(_pat_let28_0 bool) D4 {
									return func(_1141_dt__update_hcf29_h1 bool) D4 {
										return Companion_D4_.Create_DC13_(_1141_dt__update_hcf29_h1, (_1140_dt__update__tmp_h4).Dtor_cf30(), (_1140_dt__update__tmp_h4).Dtor_cf31())
									}(_pat_let28_0)
								}((_this).F4())
							}(_pat_let27_0)
						}(_1135_v73), 16)
						(_nw156).ArraySet1(_1135_v73, 17)
						_1136_v74 = _nw156
						var _1142_v75 _dafny.Set
						_ = _1142_v75
						_1142_v75 = _dafny.SetOf(_1136_v74)
						_1142_v75 = _1142_v75
						var _arr30 _dafny.Array = _this.F2()
						_ = _arr30
						var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_this.F2()), 0))
						_ = _index174
						(_arr30).ArraySet1(Companion_Default___.Fm20((_this).F4(), globalState), (_index174).Int())
						var _arr31 _dafny.Array = _this.F2()
						_ = _arr31
						var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_this.F2()), 0))
						_ = _index175
						(_arr31).ArraySet1(!(_this.F7) || (_this.F7), (_index175).Int())
						var _1143_v76 _dafny.Int
						_ = _1143_v76
						_1143_v76 = _dafny.IntOfInt64(482)
						var _1144_v77 _dafny.Array
						_ = _1144_v77
						var _len0_39 _dafny.Int = _dafny.IntOfInt64(13)
						_ = _len0_39
						var _nw157 _dafny.Array
						_ = _nw157
						if _len0_39.Cmp(_dafny.Zero) == 0 {
							_nw157 = _dafny.NewArray(_len0_39)
						} else {
							var _init39 func(_dafny.Int) _dafny.Int = (func(_1145_v76 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1146_i10 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeDivisionInt(_1146_i10, _1145_v76)
								}
							})(_1143_v76)
							_ = _init39
							var _element0_39 = _init39(_dafny.Zero)
							_ = _element0_39
							_nw157 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
							(_nw157).ArraySet1(_element0_39, 0)
							var _nativeLen0_39 = (_len0_39).Int()
							_ = _nativeLen0_39
							for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
								(_nw157).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
							}
						}
						_1144_v77 = _nw157
						var _1147_v78 _dafny.Sequence
						_ = _1147_v78
						_1147_v78 = _dafny.SeqOf((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
						var _1148_v79 _dafny.Map
						_ = _1148_v79
						_1148_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1144_v77, _dafny.SeqOf(_dafny.SeqOf((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)), _1147_v78))
						_1143_v76 = (_dafny.IntOfInt64(566)).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_1148_v79).Contains(_1144_v77) {
								return (_1148_v79).Get(_1144_v77).(_dafny.Sequence)
							}
							return _dafny.SeqOf(_1147_v78, _dafny.SeqOf(p0), _1147_v78)
						})()).Cardinality()))
						var _1149_v80 _dafny.MultiSet
						_ = _1149_v80
						_1149_v80 = _dafny.MultiSetOf((_dafny.Zero).Minus(_1143_v76))
						var _arr32 _dafny.Array = _this.F2()
						_ = _arr32
						var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_this.F2()), 0))
						_ = _index176
						(_arr32).ArraySet1(((_1149_v80).Intersection(_1149_v80)).IsSubsetOf(_1149_v80), (_index176).Int())
						(_this).F7 = _this.F7
					}
					var _1150_v81 _dafny.MultiSet
					_ = _1150_v81
					_1150_v81 = _dafny.MultiSetOf((_this).F8(), (_this).F8())
					var _1151_v82 _dafny.MultiSet
					_ = _1151_v82
					_1151_v82 = _dafny.MultiSetOf(_this.F7)
					var _1152_v83 _dafny.Sequence
					_ = _1152_v83
					_1152_v83 = _dafny.SeqOf((func() _dafny.Int {
						if (_1151_v82).Contains(p0) {
							return (_1151_v82).Multiplicity(p0)
						}
						return (_this).F8()
					})(), (_this).F8())
					var _1153_v84 D0
					_ = _1153_v84
					_1153_v84 = Companion_D0_.Create_DC2_((_dafny.MultiSetOf(true)).Cardinality(), !(_this.F7), p0, (_dafny.Zero).Minus((_this).F8()))
					var _1154_v85 _dafny.Map
					_ = _1154_v85
					_1154_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_this).F8())
					var _1155_v86 D8
					_ = _1155_v86
					_1155_v86 = Companion_D8_.Create_DC26_((_this).F4(), true, (_1153_v84).Dtor_cf5(), (func() _dafny.Int {
						if (_1154_v85).Contains((_this).F8()) {
							return (_1154_v85).Get((_this).F8()).(_dafny.Int)
						}
						return (_this).F8()
					})())
					var _1156_v87 D2
					_ = _1156_v87
					_1156_v87 = Companion_D2_.Create_DC7_(Companion_Default___.Fm0((_this).F8(), globalState), (_this).F8(), _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_1152_v83)), (_this).F8(), (_this).F8())
					var _1157_v88 _dafny.Map
					_ = _1157_v88
					_1157_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1156_v87, (_this).F8())
					var _1158_v89 _dafny.Array
					_ = _1158_v89
					var _nwElement0_31 _dafny.Int = (_this).F8()
					_ = _nwElement0_31
					var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(26))
					_ = _nw158
					(_nw158).ArraySet1(_nwElement0_31, 0)
					(_nw158).ArraySet1(((_this).F8()).Plus((_this).F8()), 1)
					(_nw158).ArraySet1(_dafny.IntOfInt64(-705), 2)
					(_nw158).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1150_v81).Cardinality(), (_this).F8())), 3)
					(_nw158).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}(func(_1159_i11 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('r')
					})), _1109_v50)).Cardinality()), 4)
					(_nw158).ArraySet1((_this).F8(), 5)
					(_nw158).ArraySet1((func() _dafny.Int {
						if _this.F7 {
							return (_this).F8()
						}
						return _dafny.IntOfInt64(471)
					})(), 6)
					(_nw158).ArraySet1((_this).F8(), 7)
					(_nw158).ArraySet1((_this).F8(), 8)
					(_nw158).ArraySet1(((_this).F8()).Times((_this).F8()), 9)
					(_nw158).ArraySet1((_1111_v51).Cardinality(), 10)
					(_nw158).ArraySet1((_1152_v83).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1152_v83).Cardinality())), _dafny.IntOfUint32((_1152_v83).Cardinality()))).Uint32()).(_dafny.Int), 11)
					(_nw158).ArraySet1((_this).F8(), 12)
					(_nw158).ArraySet1((_this).F8(), 13)
					(_nw158).ArraySet1((_this).F8(), 14)
					(_nw158).ArraySet1((_this).F8(), 15)
					(_nw158).ArraySet1((_this).F8(), 16)
					(_nw158).ArraySet1((func() _dafny.Int {
						if (_1151_v82).Contains(!(p1)) {
							return (_1151_v82).Multiplicity(!(p1))
						}
						return (_1155_v86).Dtor_cf48()
					})(), 17)
					(_nw158).ArraySet1((_this).F8(), 18)
					(_nw158).ArraySet1(((_this).F8()).Times((_this).F8()), 19)
					(_nw158).ArraySet1(Companion_Default___.SafeModuloInt((_this).F8(), (_this).F8()), 20)
					(_nw158).ArraySet1((_this).F8(), 21)
					(_nw158).ArraySet1((_this).F8(), 22)
					(_nw158).ArraySet1((_this).F8(), 23)
					(_nw158).ArraySet1((_dafny.Zero).Minus((_1157_v88).Cardinality()), 24)
					(_nw158).ArraySet1(_dafny.IntOfInt64(75), 25)
					_1158_v89 = _nw158
					var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_1158_v89), 0))
					_ = _index177
					(_1158_v89).ArraySet1((_this).F8(), (_index177).Int())
					var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_1158_v89), 0))
					_ = _index178
					(_1158_v89).ArraySet1((_this).F8(), (_index178).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		if !(true) {
			var _1160_v90 _dafny.Array
			_ = _1160_v90
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_40
			var _nw159 _dafny.Array
			_ = _nw159
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw159 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) D3 = (func(_1161_p1 bool, _1162_p0 bool) func(_dafny.Int) D3 {
					return func(_1163_i12 _dafny.Int) D3 {
						return Companion_D3_.Create_DC10_(_1161_p1, _dafny.CodePoint('i'), false, true, _1162_p0)
					}
				})(p1, p0)
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw159 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw159).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw159).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1160_v90 = _nw159
			var _1164_v91 _dafny.CodePoint
			_ = _1164_v91
			_1164_v91 = _dafny.CodePoint('q')
			var _1165_v92 D3
			_ = _1165_v92
			_1165_v92 = Companion_D3_.Create_DC10_(p1, _1164_v91, (_this).F4(), p1, p1)
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1160_v90), 0))
			_ = _index179
			(_1160_v90).ArraySet1(_1165_v92, (_index179).Int())
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1160_v90), 0))
			_ = _index180
			(_1160_v90).ArraySet1(_1165_v92, (_index180).Int())
			_1164_v91 = _1164_v91
			var _1166_v93 *C1
			_ = _1166_v93
			var _nw160 *C1 = New_C1_()
			_ = _nw160
			_nw160.Ctor__((_this).F3(), p1, _this.F2())
			_1166_v93 = _nw160
			_1109_v50 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(448))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}((func(_1167_v91 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1168_i13 _dafny.Int) _dafny.CodePoint {
					return _1167_v91
				}
			})(_1164_v91))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(194))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}((func(_1169_v91 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1170_i14 _dafny.Int) _dafny.CodePoint {
					return _1169_v91
				}
			})(_1164_v91))))
			var _1171_v94 _dafny.Int
			_ = _1171_v94
			_1171_v94 = _dafny.IntOfInt64(800)
			_1171_v94 = (_1171_v94).Plus((_this).F8())
		} else {
			var _1172_v95 _dafny.Map
			_ = _1172_v95
			_1172_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if true {
					return _this.F7
				}
				return _this.F7
			})(), p0)
			_1172_v95 = (_1172_v95).Update((_this).F4(), p1)
			var _1173_v96 _dafny.Array
			_ = _1173_v96
			var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw161
			_1173_v96 = _nw161
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))
			_ = _index181
			(_1173_v96).ArraySet1((_this).F8(), (_index181).Int())
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))
			_ = _index182
			(_1173_v96).ArraySet1(((_dafny.Zero).Minus((func() _dafny.Int {
				if true {
					return (_this).F8()
				}
				return (_this).F8()
			})())).Times((_this).F8()), (_index182).Int())
			var _1174_v97 _dafny.Sequence
			_ = _1174_v97
			_1174_v97 = _dafny.SeqOf((_1173_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))).Int()).(_dafny.Int))
			if !(!(_dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
				if (_this).F4() {
					return _1174_v97
				}
				return _1174_v97
			})(), _dafny.SeqOf(_dafny.IntOfUint32((_1109_v50).Cardinality()))))) {
				var _1175_v98 _dafny.Map
				_ = _1175_v98
				_1175_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _this.F7)
				var _1176_v99 _dafny.Set
				_ = _1176_v99
				_1176_v99 = _dafny.SetOf(Companion_Default___.Fm18(_this.F7, _1175_v98, globalState), (func() _dafny.Sequence {
					if p0 {
						return _1109_v50
					}
					return _1109_v50
				})(), _dafny.Companion_Sequence_.Concatenate(_1109_v50, _dafny.UnicodeSeqOfUtf8Bytes("cspctqgt")))
				_1176_v99 = _1176_v99
				var _arr33 _dafny.Array = _this.F2()
				_ = _arr33
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index183
				(_arr33).ArraySet1((func() bool {
					if (_1172_v95).Contains(_this.F7) {
						return (_1172_v95).Get(_this.F7).(bool)
					}
					return _this.F7
				})(), (_index183).Int())
				var _1177_v100 _dafny.CodePoint
				_ = _1177_v100
				_1177_v100 = _dafny.CodePoint('y')
				var _1178_v101 _dafny.MultiSet
				_ = _1178_v101
				_1178_v101 = _dafny.MultiSetOf(_1177_v100, _1177_v100, _1177_v100, _1177_v100, _1177_v100)
				var _1179_v102 D4
				_ = _1179_v102
				_1179_v102 = Companion_D4_.Create_DC13_(p1, _1178_v101, p1)
				var _arr34 _dafny.Array = _this.F2()
				_ = _arr34
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index184
				(_arr34).ArraySet1((func(_pat_let29_0 D4) D4 {
					return func(_1180_dt__update__tmp_h5 D4) D4 {
						return func(_pat_let30_0 bool) D4 {
							return func(_1181_dt__update_hcf29_h2 bool) D4 {
								return Companion_D4_.Create_DC13_(_1181_dt__update_hcf29_h2, (_1180_dt__update__tmp_h5).Dtor_cf30(), (_1180_dt__update__tmp_h5).Dtor_cf31())
							}(_pat_let30_0)
						}(true)
					}(_pat_let29_0)
				}(_1179_v102)).Dtor_cf29(), (_index184).Int())
				var _1182_v103 D8
				_ = _1182_v103
				_1182_v103 = Companion_D8_.Create_DC26_(p1, _this.F7, (_this).Fm12(globalState), (_1173_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))).Int()).(_dafny.Int))
				var _1183_v104 _dafny.MultiSet
				_ = _1183_v104
				_1183_v104 = _dafny.MultiSetOf(_1182_v103, _1182_v103)
				var _1184_v105 _dafny.Map
				_ = _1184_v105
				_1184_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), (_1173_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))).Int()).(_dafny.Int))
				var _arr35 _dafny.Array = _this.F2()
				_ = _arr35
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index185
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))
				_ = _index186
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))
				_ = _index187
				var _rhs144 bool = (_this).F4()
				_ = _rhs144
				var _rhs145 _dafny.Int = (_1173_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))).Int()).(_dafny.Int)
				_ = _rhs145
				var _rhs146 _dafny.Int = (func() _dafny.Int {
					if (_1183_v104).Contains(_1182_v103) {
						return (_1183_v104).Multiplicity(_1182_v103)
					}
					return ((_1173_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))).Int()).(_dafny.Int)).Minus((func() _dafny.Int {
						if (_1184_v105).Contains(_1177_v100) {
							return (_1184_v105).Get(_1177_v100).(_dafny.Int)
						}
						return (_1172_v95).Cardinality()
					})())
				})()
				_ = _rhs146
				var _lhs98 _dafny.Array = _this.F2()
				_ = _lhs98
				var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs99
				var _lhs100 _dafny.Array = _1173_v96
				_ = _lhs100
				var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))
				_ = _lhs101
				var _lhs102 _dafny.Array = _1173_v96
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))
				_ = _lhs103
				(_lhs98).ArraySet1(_rhs144, (_lhs99).Int())
				(_lhs100).ArraySet1(_rhs145, (_lhs101).Int())
				(_lhs102).ArraySet1(_rhs146, (_lhs103).Int())
				var _1185_v106 *C0
				_ = _1185_v106
				var _nw162 *C0 = New_C0_()
				_ = _nw162
				_nw162.Ctor__()
				_1185_v106 = _nw162
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))
				_ = _index188
				(_1173_v96).ArraySet1((_this).F8(), (_index188).Int())
			} else {
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))
				_ = _index189
				(_1173_v96).ArraySet1(((func() _dafny.Int {
					if p1 {
						return (_1173_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1173_v96), 0))).Int()).(_dafny.Int)
					}
					return _dafny.IntOfInt64(487)
				})()).Times((_this).F8()), (_index189).Int())
				(_this).F7 = (_this).Fm12(globalState)
				_1109_v50 = _1109_v50
				var _1186_v107 _dafny.Array
				_ = _1186_v107
				var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw163
				_1186_v107 = _nw163
				_1186_v107 = _1186_v107
				var _arr36 _dafny.Array = _this.F2()
				_ = _arr36
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index190
				(_arr36).ArraySet1((_this).F4(), (_index190).Int())
				var _arr37 _dafny.Array = _this.F2()
				_ = _arr37
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index191
				(_arr37).ArraySet1(!(((_this).F8()).Cmp((_this).F8()) >= 0), (_index191).Int())
			}
			(_this).F7 = (_this).Fm12(globalState)
			var _1187_v108 *C1
			_ = _1187_v108
			var _nw164 *C1 = New_C1_()
			_ = _nw164
			_nw164.Ctor__((_this).F3(), p1, _this.F2())
			_1187_v108 = _nw164
		}
		var _1188_v109 _dafny.Array
		_ = _1188_v109
		var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
		_ = _nw165
		_1188_v109 = _nw165
		var _1189_v110 _dafny.Array
		_ = _1189_v110
		_1189_v110 = _1188_v109
		_1188_v109 = (_1189_v110)
		var _1190_i15 _dafny.Int
		_ = _1190_i15
		_1190_i15 = _dafny.Zero
		{
			for !(p1) || (p1) {
				{
					if (_1190_i15).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1190_i15 = (_1190_i15).Plus(_dafny.One)
					var _arr38 _dafny.Array = _this.F2()
					_ = _arr38
					var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_this.F2()), 0))
					_ = _index192
					(_arr38).ArraySet1((func() bool {
						if p1 {
							return !(_this.F7)
						}
						return true
					})(), (_index192).Int())
					var _arr39 _dafny.Array = _this.F2()
					_ = _arr39
					var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_this.F2()), 0))
					_ = _index193
					(_arr39).ArraySet1(_this.F7, (_index193).Int())
					var _1191_v111 D0
					_ = _1191_v111
					_1191_v111 = Companion_D0_.Create_DC0_(p0)
					var _1192_v112 _dafny.Sequence
					_ = _1192_v112
					_1192_v112 = _dafny.SeqOf(_1191_v111, _1191_v111)
					_1192_v112 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1192_v112, _1192_v112), _1192_v112)
					var _1193_v113 _dafny.Map
					_ = _1193_v113
					_1193_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
					var _1194_v114 _dafny.Array
					_ = _1194_v114
					var _nw166 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
					_ = _nw166
					_1194_v114 = _nw166
					var _1195_v115 T1
					_ = _1195_v115
					var _nw167 *C1 = New_C1_()
					_ = _nw167
					_nw167.Ctor__((_this).F3(), (_1193_v113).Contains(_this.F7), _1194_v114)
					_1195_v115 = _nw167
					_1195_v115 = _1195_v115
					var _1196_v116 _dafny.Set
					_ = _1196_v116
					_1196_v116 = _dafny.SetOf((_1195_v115).F4())
					var _1197_v117 _dafny.Map
					_ = _1197_v117
					_1197_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _dafny.SeqOf((_1196_v116).Cardinality()))
					var _1198_v118 _dafny.Map
					_ = _1198_v118
					_1198_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F8()), (_this).F8())
					var _1199_v119 _dafny.Sequence
					_ = _1199_v119
					_1199_v119 = _dafny.SeqOf(p0)
					var _arr40 _dafny.Array = _this.F2()
					_ = _arr40
					var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_this.F2()), 0))
					_ = _index194
					(_arr40).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
						if (_1197_v117).Contains((_1198_v118).Cardinality()) {
							return (_1197_v117).Get((_1198_v118).Cardinality()).(_dafny.Sequence)
						}
						return _dafny.SeqOf((_this).F8())
					})(), Companion_Default___.Fm31((_this).F8(), !((_1195_v115).F4()), (_this).F8(), (_1199_v119).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1199_v119).Cardinality()))).Uint32()).(bool), globalState)), (_index194).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
	}
}
func (_this *C2) F8() _dafny.Int {
	{
		return _this._f8
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f2 _dafny.Array
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C3) F2() _dafny.Array {
	return _this._f2
}
func (_this *C3) F2_set_(value _dafny.Array) {
	_this._f2 = value
}
func (_this *C3) Ctor__(f2 _dafny.Array) {
	{
		(_this)._f2 = f2
	}
}
func (_this *C3) Fm1(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			if false {
				return func() _dafny.Map {
					var _coll47 = _dafny.NewMapBuilder()
					_ = _coll47
					for _iter54 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-814)), _dafny.IntOfInt64(241))).Keys().Elements()); ; {
						_compr_47, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _1200_v0 _dafny.Int
						_1200_v0 = interface{}(_compr_47).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-814)), _dafny.IntOfInt64(241))).Contains(_1200_v0) {
							_coll47.Add(Companion_Default___.SafeDivisionInt(_1200_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(195))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg56 _dafny.Int) interface{} {
									return coer56(arg56)
								}
							}(func(_1201_i0 _dafny.Int) _dafny.Int {
								return _dafny.IntOfInt64(92)
							}))).Cardinality()))).Cardinality())), false)
						}
					}
					return _coll47.ToMap()
				}()
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(929), (_dafny.Zero).Minus(_dafny.IntOfInt64(-512)), _dafny.IntOfInt64(856))).Cardinality()), true)
		})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hyrhotgv")).Cardinality()), false))
	}
}
func (_this *C3) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(859)
	}
}
func (_this *C3) M0(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Map, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1202_v0 _dafny.Array
		_ = _1202_v0
		var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw168
		_1202_v0 = _nw168
		var _1203_v1 _dafny.Map
		_ = _1203_v1
		_1203_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1202_v0, p0)
		var _1204_v2 _dafny.Array
		_ = _1204_v2
		var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
		_ = _nw169
		_1204_v2 = _nw169
		var _1205_v3 bool
		_ = _1205_v3
		_1205_v3 = true
		var _1206_v4 T0
		_ = _1206_v4
		var _nw170 *C2 = New_C2_()
		_ = _nw170
		_nw170.Ctor__(true, (func() _dafny.Int {
			if (_1203_v1).Contains(_1202_v0) {
				return (_1203_v1).Get(_1202_v0).(_dafny.Int)
			}
			return p0
		})(), _1204_v2, _1205_v3, _this.F2())
		_1206_v4 = _nw170
		_1206_v4 = _1206_v4
		r0 = _1205_v3
		var _1207_v5 _dafny.Sequence
		_ = _1207_v5
		_1207_v5 = _dafny.UnicodeSeqOfUtf8Bytes("akmdxvj")
		var _1208_v6 _dafny.Map
		_ = _1208_v6
		_1208_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1207_v5)
		var _1209_v7 _dafny.CodePoint
		_ = _1209_v7
		_1209_v7 = _dafny.CodePoint('c')
		if !(_dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
			if false {
				return _1207_v5
			}
			return (func() _dafny.Sequence {
				if (_1208_v6).Contains(p0) {
					return (_1208_v6).Get(p0).(_dafny.Sequence)
				}
				return _1207_v5
			})()
		})(), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ebwdtouqi"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ebwdtouqi")).Cardinality()))).Uint32(), _1209_v7))) {
			var _1210_v8 _dafny.Array
			_ = _1210_v8
			var _nwElement0_32 _dafny.Sequence = _1207_v5
			_ = _nwElement0_32
			var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(13))
			_ = _nw171
			(_nw171).ArraySet1(_nwElement0_32, 0)
			(_nw171).ArraySet1(_1207_v5, 1)
			(_nw171).ArraySet1(_1207_v5, 2)
			(_nw171).ArraySet1(_1207_v5, 3)
			(_nw171).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-257))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}((func(_1211_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1212_i0 _dafny.Int) _dafny.CodePoint {
					return _1211_v7
				}
			})(_1209_v7))), 4)
			(_nw171).ArraySet1(_1207_v5, 5)
			(_nw171).ArraySet1(_dafny.SeqOf(_1209_v7, Companion_Default___.Fm22(_1207_v5, globalState)), 6)
			(_nw171).ArraySet1(_1207_v5, 7)
			(_nw171).ArraySet1(_1207_v5, 8)
			(_nw171).ArraySet1(_1207_v5, 9)
			(_nw171).ArraySet1(_dafny.SeqOf(_1209_v7, _1209_v7), 10)
			(_nw171).ArraySet1(_dafny.SeqOf(_1209_v7), 11)
			(_nw171).ArraySet1(_1207_v5, 12)
			_1210_v8 = _nw171
			var _1213_v9 _dafny.Sequence
			_ = _1213_v9
			_1213_v9 = _dafny.SeqOf(p0)
			var _1214_v10 _dafny.Map
			_ = _1214_v10
			_1214_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1213_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1213_v9).Cardinality()))).Uint32()).(_dafny.Int), p0)
			var _1215_v12 _dafny.Array
			_ = _1215_v12
			var _nwElement0_33 _dafny.Map = _1214_v10
			_ = _nwElement0_33
			var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(7))
			_ = _nw172
			(_nw172).ArraySet1(_nwElement0_33, 0)
			(_nw172).ArraySet1(_1214_v10, 1)
			(_nw172).ArraySet1(_1214_v10, 2)
			(_nw172).ArraySet1(_1214_v10, 3)
			(_nw172).ArraySet1(_1214_v10, 4)
			(_nw172).ArraySet1((func() _dafny.Map {
				var _coll48 = _dafny.NewMapBuilder()
				_ = _coll48
				for _iter55 := _dafny.Iterate((_1213_v9).Elements()); ; {
					_compr_48, _ok55 := _iter55()
					if !_ok55 {
						break
					}
					var _1216_v11 _dafny.Int
					_1216_v11 = interface{}(_compr_48).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_1213_v9, _1216_v11) {
						_coll48.Add((_1216_v11).Plus(p0), p0)
					}
				}
				return _coll48.ToMap()
			}()).Update(p0, _dafny.IntOfUint32((_1213_v9).Cardinality())), 5)
			(_nw172).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), 6)
			_1215_v12 = _nw172
			var _1217_v13 _dafny.MultiSet
			_ = _1217_v13
			var _1218_v14 bool
			_ = _1218_v14
			var _out28 _dafny.MultiSet
			_ = _out28
			var _out29 bool
			_ = _out29
			_out28, _out29 = (_1206_v4).M1(p0, _dafny.CodePoint('b'), _1210_v8, _1215_v12, globalState)
			_1217_v13 = _out28
			_1218_v14 = _out29
			var _1219_v15 _dafny.Sequence
			_ = _1219_v15
			_1219_v15 = _dafny.SeqOf(_1218_v14)
			_1205_v3 = (func() bool {
				if _dafny.Companion_Sequence_.IsProperPrefixOf(_1219_v15, _1219_v15) {
					return false
				}
				return (p0).Cmp((_dafny.SetOf(_1209_v7)).Cardinality()) >= 0
			})()
			r0 = _1205_v3
			if _1205_v3 {
				var _1220_v16 _dafny.Map
				_ = _1220_v16
				_1220_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1218_v14, _1210_v8)
				_1220_v16 = (_1220_v16).Update(_1218_v14, _1210_v8)
				var _1221_v17 _dafny.Array
				_ = _1221_v17
				var _nw173 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw173
				_1221_v17 = _nw173
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_1221_v17), 0))
				_ = _index195
				(_1221_v17).ArraySet1(Companion_Default___.Fm31((_1213_v9).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1207_v5).Cardinality())), _dafny.IntOfUint32((_1213_v9).Cardinality()))).Uint32()).(_dafny.Int), _1205_v3, (_dafny.Zero).Minus(p0), Companion_Default___.Fm20(!(_1205_v3), globalState), globalState), (_index195).Int())
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_1221_v17), 0))
				_ = _index196
				(_1221_v17).ArraySet1(_1213_v9, (_index196).Int())
				var _1222_v18 _dafny.Map
				_ = _1222_v18
				_1222_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1205_v3), _1205_v3)
				var _1223_v19 D8
				_ = _1223_v19
				_1223_v19 = Companion_D8_.Create_DC26_(_1218_v14, _1218_v14, _1205_v3, (_1222_v18).Cardinality())
				_1205_v3 = (_1223_v19).Dtor_cf45()
				_1218_v14 = _1218_v14
				var _1224_v20 _dafny.Array
				_ = _1224_v20
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw174
				_1224_v20 = _nw174
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_1224_v20), 0))
				_ = _index197
				(_1224_v20).ArraySet1(_1202_v0, (_index197).Int())
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_1224_v20), 0))
				_ = _index198
				(_1224_v20).ArraySet1(_1202_v0, (_index198).Int())
			} else {
				var _1225_v21 _dafny.Int
				_ = _1225_v21
				_1225_v21 = _dafny.IntOfInt64(-152)
				_1225_v21 = _1225_v21
				var _1226_v22 *C1
				_ = _1226_v22
				var _nw175 *C1 = New_C1_()
				_ = _nw175
				_nw175.Ctor__(_1204_v2, false, _this.F2())
				_1226_v22 = _nw175
				var _1227_v23 _dafny.Map
				_ = _1227_v23
				_1227_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1226_v22, (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(382))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}((func(_1228_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1229_i1 _dafny.Int) _dafny.CodePoint {
						return _1228_v7
					}
				})(_1209_v7)))).Cardinality())).Cmp(_1225_v21) >= 0)
				var _1230_v24 D10
				_ = _1230_v24
				_1230_v24 = Companion_D10_.Create_DC31_(_1226_v22)
				_1227_v23 = (_1227_v23).Update((_1230_v24).Dtor_cf57(), (Companion_Default___.Fm20(_1205_v3, globalState)) && (_1218_v14))
				r0 = false
				var _1231_v25 _dafny.MultiSet
				_ = _1231_v25
				var _1232_v26 bool
				_ = _1232_v26
				var _out30 _dafny.MultiSet
				_ = _out30
				var _out31 bool
				_ = _out31
				_out30, _out31 = (_1206_v4).M1(Companion_Default___.SafeDivisionInt(p0, p0), _1209_v7, _1210_v8, _1215_v12, globalState)
				_1231_v25 = _out30
				_1232_v26 = _out31
				var _rhs147 bool = _1205_v3
				_ = _rhs147
				var _rhs148 _dafny.Int = (p0).Times(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1225_v21), p0))
				_ = _rhs148
				_1232_v26 = _rhs147
				_1225_v21 = _rhs148
			}
			var _1233_v27 _dafny.Set
			_ = _1233_v27
			_1233_v27 = _dafny.SetOf(p0)
			var _1234_v28 _dafny.Sequence
			_ = _1234_v28
			_1234_v28 = _dafny.SeqOf(_1233_v27)
			var _1235_v29 _dafny.Map
			_ = _1235_v29
			_1235_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1207_v5).Cardinality()), _1234_v28)
			_1235_v29 = (_1235_v29).Update(p0, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1233_v27), _dafny.SeqOf(_1233_v27, _1233_v27)))
		} else {
			var _1236_v30 _dafny.Array
			_ = _1236_v30
			var _nwElement0_34 _dafny.CodePoint = _1209_v7
			_ = _nwElement0_34
			var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(28))
			_ = _nw176
			(_nw176).ArraySet1CodePoint(_nwElement0_34, 0)
			(_nw176).ArraySet1CodePoint(_1209_v7, 1)
			(_nw176).ArraySet1CodePoint((_1207_v5).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1207_v5).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
			(_nw176).ArraySet1CodePoint((func() _dafny.CodePoint {
				if true {
					return _1209_v7
				}
				return _1209_v7
			})(), 3)
			(_nw176).ArraySet1CodePoint((_1207_v5).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1207_v5).Cardinality()))).Uint32()).(_dafny.CodePoint), 4)
			(_nw176).ArraySet1CodePoint(_1209_v7, 5)
			(_nw176).ArraySet1CodePoint(_1209_v7, 6)
			(_nw176).ArraySet1CodePoint(_1209_v7, 7)
			(_nw176).ArraySet1CodePoint(_1209_v7, 8)
			(_nw176).ArraySet1CodePoint(_1209_v7, 9)
			(_nw176).ArraySet1CodePoint(_1209_v7, 10)
			(_nw176).ArraySet1CodePoint(_1209_v7, 11)
			(_nw176).ArraySet1CodePoint(_1209_v7, 12)
			(_nw176).ArraySet1CodePoint(_1209_v7, 13)
			(_nw176).ArraySet1CodePoint(_1209_v7, 14)
			(_nw176).ArraySet1CodePoint(_1209_v7, 15)
			(_nw176).ArraySet1CodePoint(_1209_v7, 16)
			(_nw176).ArraySet1CodePoint((func() _dafny.CodePoint {
				if !(_1205_v3) {
					return _1209_v7
				}
				return _1209_v7
			})(), 17)
			(_nw176).ArraySet1CodePoint(_1209_v7, 18)
			(_nw176).ArraySet1CodePoint(_1209_v7, 19)
			(_nw176).ArraySet1CodePoint(_dafny.CodePoint('d'), 20)
			(_nw176).ArraySet1CodePoint(_1209_v7, 21)
			(_nw176).ArraySet1CodePoint(Companion_Default___.Fm22(_1207_v5, globalState), 22)
			(_nw176).ArraySet1CodePoint(_1209_v7, 23)
			(_nw176).ArraySet1CodePoint(_1209_v7, 24)
			(_nw176).ArraySet1CodePoint(_1209_v7, 25)
			(_nw176).ArraySet1CodePoint(_dafny.CodePoint('u'), 26)
			(_nw176).ArraySet1CodePoint(_1209_v7, 27)
			_1236_v30 = _nw176
			var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(27))
			_ = _nw177
			_1236_v30 = _nw177
			var _1237_v31 _dafny.Int
			_ = _1237_v31
			_1237_v31 = _dafny.IntOfInt64(-732)
			_1237_v31 = _1237_v31
			var _arr41 _dafny.Array = _this.F2()
			_ = _arr41
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index199
			(_arr41).ArraySet1(Companion_Default___.Fm20(_1205_v3, globalState), (_index199).Int())
			var _arr42 _dafny.Array = _this.F2()
			_ = _arr42
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index200
			(_arr42).ArraySet1(_1205_v3, (_index200).Int())
			_1207_v5 = _1207_v5
			var _1238_v32 _dafny.Sequence
			_ = _1238_v32
			_1238_v32 = _dafny.SeqOf(_1237_v31)
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_1202_v0), 0))
			_ = _index201
			(_1202_v0).ArraySet1((Companion_D8_.Create_DC28_(_1205_v3, _dafny.IntOfUint32((_1238_v32).Cardinality()))).Dtor_cf54(), (_index201).Int())
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_1202_v0), 0))
			_ = _index202
			(_1202_v0).ArraySet1((_1238_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.IntOfUint32((_1238_v32).Cardinality()))).Uint32()).(_dafny.Int), (_index202).Int())
		}
		var _1239_v33 _dafny.Sequence
		_ = _1239_v33
		_1239_v33 = _dafny.SeqOf((true) && (_1205_v3), _1205_v3, !(_1205_v3))
		var _1240_v34 _dafny.Map
		_ = _1240_v34
		_1240_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1205_v3, _1205_v3)
		_1239_v33 = Companion_Default___.Fm32(_dafny.IntOfUint32((_dafny.SeqOf(_1202_v0)).Cardinality()), !(_1205_v3), _1240_v34, globalState)
		var _1241_v35 _dafny.Map
		_ = _1241_v35
		_1241_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-184))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}((func(_1242_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1243_i2 _dafny.Int) _dafny.CodePoint {
				return _1242_v7
			}
		})(_1209_v7))), p0)
		var _1244_v36 D8
		_ = _1244_v36
		_1244_v36 = Companion_D8_.Create_DC25_(_1241_v35)
		var _source16 D8 = _1244_v36
		_ = _source16
		if _source16.Is_DC26() {
			var _1245___mcc_h0 bool = _source16.Get_().(D8_DC26).Cf45
			_ = _1245___mcc_h0
			var _1246___mcc_h1 bool = _source16.Get_().(D8_DC26).Cf46
			_ = _1246___mcc_h1
			var _1247___mcc_h2 bool = _source16.Get_().(D8_DC26).Cf47
			_ = _1247___mcc_h2
			var _1248___mcc_h3 _dafny.Int = _source16.Get_().(D8_DC26).Cf48
			_ = _1248___mcc_h3
			var _1249_cf48 _dafny.Int = _1248___mcc_h3
			_ = _1249_cf48
			var _1250_cf47 bool = _1247___mcc_h2
			_ = _1250_cf47
			var _1251_cf46 bool = _1246___mcc_h1
			_ = _1251_cf46
			var _1252_cf45 bool = _1245___mcc_h0
			_ = _1252_cf45
			var _arr43 _dafny.Array = _this.F2()
			_ = _arr43
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index203
			(_arr43).ArraySet1(_1251_cf46, (_index203).Int())
			var _arr44 _dafny.Array = _this.F2()
			_ = _arr44
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_this.F2()), 0))
			_ = _index204
			(_arr44).ArraySet1(!(((_dafny.IntOfUint32((_1207_v5).Cardinality())).Cmp(_dafny.IntOfUint32((_1207_v5).Cardinality())) < 0) || (_1250_cf47)), (_index204).Int())
			var _1253_v37 _dafny.Sequence
			_ = _1253_v37
			_1253_v37 = _dafny.SeqOf(_dafny.IntOfUint32((_1207_v5).Cardinality()), _dafny.IntOfUint32((_1207_v5).Cardinality()), p0)
			var _1254_v38 _dafny.MultiSet
			_ = _1254_v38
			_1254_v38 = _dafny.MultiSetOf(p0, (_this).Fm2(globalState))
			var _rhs149 bool = ((func() bool {
				if _1205_v3 {
					return false
				}
				return false
			})()) && ((_1254_v38).IsProperSubsetOf(_dafny.MultiSetFromSeq(_1253_v37)))
			_ = _rhs149
			var _rhs150 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1207_v5, _1207_v5), _1207_v5), (Companion_Default___.SafeIndex(_1249_cf48, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1207_v5, _1207_v5), _1207_v5)).Cardinality()))).Uint32(), _1209_v7)
			_ = _rhs150
			var _rhs151 bool = _1252_cf45
			_ = _rhs151
			_1205_v3 = _rhs149
			_1207_v5 = _rhs150
			_1251_cf46 = _rhs151
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1202_v0), 0))
			_ = _index205
			(_1202_v0).ArraySet1(_1249_cf48, (_index205).Int())
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1202_v0), 0))
			_ = _index206
			(_1202_v0).ArraySet1(p0, (_index206).Int())
			var _1255_v39 D6
			_ = _1255_v39
			_1255_v39 = Companion_D6_.Create_DC20_(_1251_cf46, _1251_cf46)
			var _1256_v40 D6
			_ = _1256_v40
			_1256_v40 = Companion_D6_.Create_DC21_(_1255_v39)
			var _1257_v41 D6
			_ = _1257_v41
			_1257_v41 = Companion_D6_.Create_DC21_(_1255_v39)
			var _pat_let_tv14 = _1205_v3
			_ = _pat_let_tv14
			var _pat_let_tv15 = _1252_cf45
			_ = _pat_let_tv15
			var _1258_v42 _dafny.Map
			_ = _1258_v42
			_1258_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm14((_1202_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_1202_v0), 0))).Int()).(_dafny.Int), _1249_cf48, _1251_cf46, (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), globalState), func(_pat_let31_0 D6) D6 {
				return func(_1259_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let32_0 D6) D6 {
						return func(_1260_dt__update_hcf40_h0 D6) D6 {
							return Companion_D6_.Create_DC21_(_1260_dt__update_hcf40_h0)
						}(_pat_let32_0)
					}(Companion_D6_.Create_DC20_(_pat_let_tv14, _pat_let_tv15))
				}(_pat_let31_0)
			}(_1257_v41))
			var _1261_v43 D0
			_ = _1261_v43
			_1261_v43 = Companion_D0_.Create_DC0_((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
			_1258_v42 = (_1258_v42).Update(_1261_v43, Companion_D6_.Create_DC21_(_1256_v40))
		} else if _source16.Is_DC27() {
			var _1262___mcc_h4 bool = _source16.Get_().(D8_DC27).Cf49
			_ = _1262___mcc_h4
			var _1263___mcc_h5 bool = _source16.Get_().(D8_DC27).Cf50
			_ = _1263___mcc_h5
			var _1264___mcc_h6 _dafny.Int = _source16.Get_().(D8_DC27).Cf51
			_ = _1264___mcc_h6
			var _1265___mcc_h7 _dafny.Int = _source16.Get_().(D8_DC27).Cf52
			_ = _1265___mcc_h7
			var _1266_cf52 _dafny.Int = _1265___mcc_h7
			_ = _1266_cf52
			var _1267_cf51 _dafny.Int = _1264___mcc_h6
			_ = _1267_cf51
			var _1268_cf50 bool = _1263___mcc_h5
			_ = _1268_cf50
			var _1269_cf49 bool = _1262___mcc_h4
			_ = _1269_cf49
			var _1270_v44 _dafny.Array
			_ = _1270_v44
			var _nwElement0_35 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-36))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}((func(_1271_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1272_i3 _dafny.Int) _dafny.CodePoint {
					return _1271_v7
				}
			})(_1209_v7)))
			_ = _nwElement0_35
			var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(28))
			_ = _nw178
			(_nw178).ArraySet1(_nwElement0_35, 0)
			(_nw178).ArraySet1(_1207_v5, 1)
			(_nw178).ArraySet1(_1207_v5, 2)
			(_nw178).ArraySet1(_dafny.Companion_Sequence_.Update(_1207_v5, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1207_v5).Cardinality()))).Uint32(), _1209_v7), 3)
			(_nw178).ArraySet1(_1207_v5, 4)
			(_nw178).ArraySet1(_1207_v5, 5)
			(_nw178).ArraySet1(_1207_v5, 6)
			(_nw178).ArraySet1(_1207_v5, 7)
			(_nw178).ArraySet1(_dafny.Companion_Sequence_.Update(_1207_v5, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1207_v5).Cardinality()))).Uint32(), _1209_v7), 8)
			(_nw178).ArraySet1(_1207_v5, 9)
			(_nw178).ArraySet1(_1207_v5, 10)
			(_nw178).ArraySet1(_1207_v5, 11)
			(_nw178).ArraySet1(_1207_v5, 12)
			(_nw178).ArraySet1(_1207_v5, 13)
			(_nw178).ArraySet1(_1207_v5, 14)
			(_nw178).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(152))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}((func(_1273_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1274_i4 _dafny.Int) _dafny.CodePoint {
					return _1273_v7
				}
			})(_1209_v7))), 15)
			(_nw178).ArraySet1(_1207_v5, 16)
			(_nw178).ArraySet1(_1207_v5, 17)
			(_nw178).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(481))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}(func(_1275_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), 18)
			(_nw178).ArraySet1(_1207_v5, 19)
			(_nw178).ArraySet1(_1207_v5, 20)
			(_nw178).ArraySet1(Companion_Default___.Fm13(false, _1205_v3, true, globalState), 21)
			(_nw178).ArraySet1(_1207_v5, 22)
			(_nw178).ArraySet1(_1207_v5, 23)
			(_nw178).ArraySet1(_1207_v5, 24)
			(_nw178).ArraySet1(_1207_v5, 25)
			(_nw178).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(890))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}((func(_1276_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1277_i6 _dafny.Int) _dafny.CodePoint {
					return _1276_v7
				}
			})(_1209_v7))), 26)
			(_nw178).ArraySet1(_dafny.SeqOf(_1209_v7), 27)
			_1270_v44 = _nw178
			var _1278_v45 _dafny.Array
			_ = _1278_v45
			var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
			_ = _nw179
			_1278_v45 = _nw179
			var _1279_v46 _dafny.MultiSet
			_ = _1279_v46
			var _1280_v47 bool
			_ = _1280_v47
			var _out32 _dafny.MultiSet
			_ = _out32
			var _out33 bool
			_ = _out33
			_out32, _out33 = (_1206_v4).M1(_1266_cf52, _1209_v7, _1270_v44, _1278_v45, globalState)
			_1279_v46 = _out32
			_1280_v47 = _out33
			var _1281_v48 _dafny.Map
			_ = _1281_v48
			_1281_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gran")).Cardinality()), _1205_v3)
			_1207_v5 = Companion_Default___.Fm18(_1205_v3, (_1281_v48).Merge(_1281_v48), globalState)
			var _1282_v49 _dafny.Array
			_ = _1282_v49
			var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(11))
			_ = _nw180
			_1282_v49 = _nw180
			var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_1282_v49), 0))
			_ = _index207
			(_1282_v49).ArraySet1CodePoint(_dafny.CodePoint('x'), (_index207).Int())
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_1282_v49), 0))
			_ = _index208
			(_1282_v49).ArraySet1CodePoint((func() _dafny.CodePoint {
				if ((_dafny.Zero).Minus(p0)).Cmp(_dafny.IntOfInt64(539)) > 0 {
					return _1209_v7
				}
				return (_1207_v5).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1207_v5).Cardinality()))).Uint32()).(_dafny.CodePoint)
			})(), (_index208).Int())
			_1266_cf52 = Companion_Default___.SafeModuloInt(_1266_cf52, _1267_cf51)
		} else if _source16.Is_DC28() {
			var _1283___mcc_h8 bool = _source16.Get_().(D8_DC28).Cf53
			_ = _1283___mcc_h8
			var _1284___mcc_h9 _dafny.Int = _source16.Get_().(D8_DC28).Cf54
			_ = _1284___mcc_h9
			var _1285_cf54 _dafny.Int = _1284___mcc_h9
			_ = _1285_cf54
			var _1286_cf53 bool = _1283___mcc_h8
			_ = _1286_cf53
			r0 = _1205_v3
			var _1287_v50 *C0
			_ = _1287_v50
			var _nw181 *C0 = New_C0_()
			_ = _nw181
			_nw181.Ctor__()
			_1287_v50 = _nw181
			_1285_cf54 = Companion_Default___.SafeDivisionInt(_1285_cf54, (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(415))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg64 _dafny.Int) interface{} {
					return coer64(arg64)
				}
			}(func(_1288_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))).Cardinality())).Times(p0)))
			var _1289_v51 _dafny.MultiSet
			_ = _1289_v51
			_1289_v51 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(31))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}((func(_1290_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1291_i8 _dafny.Int) _dafny.Int {
					return _1290_p0
				}
			})(p0))))
			var _1292_v52 _dafny.Sequence
			_ = _1292_v52
			_1292_v52 = _dafny.SeqOf(_1285_cf54, _1285_cf54)
			var _1293_v53 _dafny.Map
			_ = _1293_v53
			_1293_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1286_cf53, (_1292_v52).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.IntOfUint32((_1292_v52).Cardinality()))).Uint32()).(_dafny.Int))
			var _1294_v54 D11
			_ = _1294_v54
			_1294_v54 = Companion_D11_.Create_DC34_(_1293_v53)
			var _1295_v55 _dafny.Sequence
			_ = _1295_v55
			_1295_v55 = _dafny.SeqOf(((_1294_v54).Dtor_cf62()).Cardinality(), _dafny.IntOfUint32((_1292_v52).Cardinality()))
			var _1296_v56 _dafny.Sequence
			_ = _1296_v56
			_1296_v56 = _dafny.SeqOf(_1295_v55)
			var _1297_v57 _dafny.MultiSet
			_ = _1297_v57
			_1297_v57 = _dafny.MultiSetOf(_1285_cf54)
			var _1298_v58 D10
			_ = _1298_v58
			_1298_v58 = Companion_D10_.Create_DC33_(((_dafny.MultiSetFromSeq(_1296_v56)).Update(_1295_v55, Companion_Default___.Abs((_1297_v57).Cardinality()))).IsProperSubsetOf(_1289_v51))
			var _source17 D10 = _1298_v58
			_ = _source17
			if _source17.Is_DC32() {
				var _1299___mcc_h12 _dafny.MultiSet = _source17.Get_().(D10_DC32).Cf58
				_ = _1299___mcc_h12
				var _1300___mcc_h13 _dafny.Int = _source17.Get_().(D10_DC32).Cf59
				_ = _1300___mcc_h13
				var _1301___mcc_h14 _dafny.Int = _source17.Get_().(D10_DC32).Cf60
				_ = _1301___mcc_h14
				var _1302_cf60 _dafny.Int = _1301___mcc_h14
				_ = _1302_cf60
				var _1303_cf59 _dafny.Int = _1300___mcc_h13
				_ = _1303_cf59
				var _1304_cf58 _dafny.MultiSet = _1299___mcc_h12
				_ = _1304_cf58
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index209
				(_1202_v0).ArraySet1(_1302_cf60, (_index209).Int())
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index210
				(_1202_v0).ArraySet1(_1302_cf60, (_index210).Int())
				_1302_cf60 = _1302_cf60
				var _1305_v59 D4
				_ = _1305_v59
				_1305_v59 = Companion_D4_.Create_DC13_(_1205_v3, _dafny.MultiSetOf(_dafny.CodePoint('j')), _1205_v3)
				var _1306_v60 D4
				_ = _1306_v60
				_1306_v60 = Companion_D4_.Create_DC13_(_1205_v3, (_1305_v59).Dtor_cf30(), _1205_v3)
				var _1307_v61 *C1
				_ = _1307_v61
				var _nw182 *C1 = New_C1_()
				_ = _nw182
				_nw182.Ctor__(_1204_v2, ((_1305_v59).Dtor_cf29()) || (_1286_cf53), _1206_v4.F2())
				_1307_v61 = _nw182
				var _1308_v62 *C0
				_ = _1308_v62
				var _nw183 *C0 = New_C0_()
				_ = _nw183
				_nw183.Ctor__()
				_1308_v62 = _nw183
			} else if _source17.Is_DC33() {
				var _1309___mcc_h15 bool = _source17.Get_().(D10_DC33).Cf61
				_ = _1309___mcc_h15
				var _1310_cf61 bool = _1309___mcc_h15
				_ = _1310_cf61
				var _1311_v63 _dafny.Map
				_ = _1311_v63
				_1311_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-460), _1205_v3)
				var _1312_v64 _dafny.Map
				_ = _1312_v64
				_1312_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1311_v63, _1298_v58)
				_1312_v64 = (_1312_v64).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1239_v33).Cardinality()), _1286_cf53)).Merge(_1311_v63), (func() D10 {
					if _1310_cf61 {
						return _1298_v58
					}
					return _1298_v58
				})())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index211
				(_1202_v0).ArraySet1(p0, (_index211).Int())
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index212
				(_1202_v0).ArraySet1((p0).Minus(p0), (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index213
				(_1202_v0).ArraySet1(((_1202_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1202_v0), 0))).Int()).(_dafny.Int)).Plus(p0), (_index213).Int())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index214
				(_1202_v0).ArraySet1((_1202_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1202_v0), 0))).Int()).(_dafny.Int), (_index214).Int())
			} else {
				var _1313___mcc_h16 *C1 = _source17.Get_().(D10_DC31).Cf57
				_ = _1313___mcc_h16
				var _1314_cf57 *C1 = _1313___mcc_h16
				_ = _1314_cf57
				_1285_cf54 = (_dafny.Zero).Minus((p0).Minus(_dafny.IntOfInt64(768)))
				var _1315_v65 D10
				_ = _1315_v65
				_1315_v65 = Companion_D10_.Create_DC31_((func() *C1 {
					if false {
						return _1314_cf57
					}
					return _1314_cf57
				})())
				var _rhs152 _dafny.Int = (((_1293_v53).Cardinality()).Minus(_dafny.IntOfUint32((_1207_v5).Cardinality()))).Plus(p0)
				_ = _rhs152
				var _rhs153 D10 = _1315_v65
				_ = _rhs153
				var _rhs154 _dafny.Int = p0
				_ = _rhs154
				_1285_cf54 = _rhs152
				_1315_v65 = _rhs153
				_1285_cf54 = _rhs154
				var _1316_v66 D6
				_ = _1316_v66
				_1316_v66 = Companion_D6_.Create_DC20_(_1286_cf53, !(_1286_cf53))
				var _pat_let_tv16 = _1205_v3
				_ = _pat_let_tv16
				var _1317_v67 _dafny.Array
				_ = _1317_v67
				var _nwElement0_36 D6 = _1316_v66
				_ = _nwElement0_36
				var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(8))
				_ = _nw184
				(_nw184).ArraySet1(_nwElement0_36, 0)
				(_nw184).ArraySet1(_1316_v66, 1)
				(_nw184).ArraySet1(_1316_v66, 2)
				(_nw184).ArraySet1(func(_pat_let33_0 D6) D6 {
					return func(_1318_dt__update__tmp_h1 D6) D6 {
						return func(_pat_let34_0 bool) D6 {
							return func(_1319_dt__update_hcf39_h0 bool) D6 {
								return Companion_D6_.Create_DC20_((_1318_dt__update__tmp_h1).Dtor_cf38(), _1319_dt__update_hcf39_h0)
							}(_pat_let34_0)
						}(_pat_let_tv16)
					}(_pat_let33_0)
				}(_1316_v66), 3)
				(_nw184).ArraySet1(_1316_v66, 4)
				(_nw184).ArraySet1(_1316_v66, 5)
				(_nw184).ArraySet1(_1316_v66, 6)
				(_nw184).ArraySet1(Companion_D6_.Create_DC20_(_1286_cf53, _1205_v3), 7)
				_1317_v67 = _nw184
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_1317_v67), 0))
				_ = _index215
				(_1317_v67).ArraySet1(_1316_v66, (_index215).Int())
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_1317_v67), 0))
				_ = _index216
				(_1317_v67).ArraySet1(_1316_v66, (_index216).Int())
				var _1320_v68 D5
				_ = _1320_v68
				_1320_v68 = Companion_D5_.Create_DC16_()
				var _1321_v69 _dafny.Sequence
				_ = _1321_v69
				_1321_v69 = _dafny.SeqOf(_1320_v68, _1320_v68)
				var _1322_v70 _dafny.Sequence
				_ = _1322_v70
				_1322_v70 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1207_v5, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1321_v69).Cardinality()), _dafny.IntOfUint32((_1207_v5).Cardinality()))).Uint32(), _1209_v7), _1207_v5, _1207_v5)
				var _1323_v71 _dafny.Array
				_ = _1323_v71
				var _nw185 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
				_ = _nw185
				_1323_v71 = _nw185
				var _1324_v72 _dafny.Map
				_ = _1324_v72
				_1324_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1206_v4.F2(), p0)
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_1323_v71), 0))
				_ = _index217
				(_1323_v71).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1206_v4.F2(), _1285_cf54)).Merge(_1324_v72), (_index217).Int())
				var _1325_v73 _dafny.MultiSet
				_ = _1325_v73
				_1325_v73 = _dafny.MultiSetOf(_1205_v3, !(_1205_v3), (_1239_v33).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1239_v33).Cardinality()))).Uint32()).(bool))
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_1323_v71), 0))
				_ = _index218
				var _rhs155 _dafny.Int = (p0).Plus((_dafny.IntOfUint32((_1292_v52).Cardinality())).Minus((_dafny.Zero).Minus(p0)))
				_ = _rhs155
				var _rhs156 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1322_v70, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(854))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}((func(_1326_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1327_i9 _dafny.Int) _dafny.Sequence {
						return _1326_v5
					}
				})(_1207_v5))))
				_ = _rhs156
				var _rhs157 bool = ((_dafny.MultiSetOf(!(false), (_1287_v50).Fm7(_1205_v3, _1285_cf54, _1205_v3, _dafny.IntOfUint32((_1239_v33).Cardinality()), globalState))).Update(_1286_cf53, Companion_Default___.Abs(_1285_cf54))).IsSubsetOf(_1325_v73)
				_ = _rhs157
				var _rhs158 _dafny.Sequence = _dafny.SeqOf((_dafny.Zero).Minus(_1285_cf54), p0, _dafny.IntOfInt64(276), (_dafny.Zero).Minus(_1285_cf54))
				_ = _rhs158
				var _rhs159 _dafny.Map = _1324_v72
				_ = _rhs159
				var _lhs104 _dafny.Array = _1323_v71
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(979), _dafny.ArrayLen((_1323_v71), 0))
				_ = _lhs105
				_1285_cf54 = _rhs155
				_1322_v70 = _rhs156
				r0 = _rhs157
				_1295_v55 = _rhs158
				(_lhs104).ArraySet1(_rhs159, (_lhs105).Int())
			}
		} else if _source16.Is_DC25() {
			var _1328___mcc_h10 _dafny.Map = _source16.Get_().(D8_DC25).Cf44
			_ = _1328___mcc_h10
			var _1329_cf44 _dafny.Map = _1328___mcc_h10
			_ = _1329_cf44
			var _1330_v74 *C0
			_ = _1330_v74
			var _nw186 *C0 = New_C0_()
			_ = _nw186
			_nw186.Ctor__()
			_1330_v74 = _nw186
			var _1331_v75 _dafny.Map
			_ = _1331_v75
			_1331_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1330_v74, _1207_v5)
			_1331_v75 = (_1331_v75).Update(_1330_v74, _dafny.UnicodeSeqOfUtf8Bytes("nervdqu"))
			if Companion_Default___.Fm20(_1205_v3, globalState) {
				_1209_v7 = _1209_v7
				var _1332_v76 _dafny.Int
				_ = _1332_v76
				_1332_v76 = _dafny.IntOfInt64(-353)
				_1332_v76 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1207_v5, _1207_v5)).Cardinality())
				var _arr45 _dafny.Array = _1206_v4.F2()
				_ = _arr45
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_1206_v4.F2()), 0))
				_ = _index219
				(_arr45).ArraySet1(((func() _dafny.Set {
					var _coll49 = _dafny.NewBuilder()
					_ = _coll49
					for _iter56 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1205_v3)).Keys().Elements()); ; {
						_compr_49, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1333_v77 _dafny.Int
						_1333_v77 = interface{}(_compr_49).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1205_v3)).Contains(_1333_v77) {
							_coll49.Add(Companion_Default___.SafeDivisionInt(_1333_v77, _dafny.IntOfInt64(893)))
						}
					}
					return _coll49.ToSet()
				}()).Cardinality()).Cmp(p0) >= 0, (_index219).Int())
				var _arr46 _dafny.Array = _1206_v4.F2()
				_ = _arr46
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_1206_v4.F2()), 0))
				_ = _index220
				(_arr46).ArraySet1(_1205_v3, (_index220).Int())
				var _1334_v78 _dafny.Map
				_ = _1334_v78
				_1334_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1207_v5, _dafny.UnicodeSeqOfUtf8Bytes("ockkobn"))
				var _1335_v79 D10
				_ = _1335_v79
				_1335_v79 = Companion_D10_.Create_DC33_(_1205_v3)
				var _1336_v80 _dafny.Sequence
				_ = _1336_v80
				_1336_v80 = _dafny.SeqOf(_1335_v79, _1335_v79, _1335_v79, Companion_D10_.Create_DC33_((_1206_v4.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_1206_v4.F2()), 0))).Int()).(bool)))
				var _1337_v81 _dafny.Map
				_ = _1337_v81
				_1337_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if (_1334_v78).Contains(_1207_v5) {
						return (_1334_v78).Get(_1207_v5).(_dafny.Sequence)
					}
					return _1207_v5
				})(), _1336_v80)
				var _1338_v82 D12
				_ = _1338_v82
				_1338_v82 = Companion_D12_.Create_DC36_(_1337_v81)
				_1337_v81 = (_1338_v82).Dtor_cf66()
				var _1339_v83 _dafny.Array
				_ = _1339_v83
				var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw187
				_1339_v83 = _nw187
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(308), _dafny.ArrayLen((_1339_v83), 0))
				_ = _index221
				(_1339_v83).ArraySet1(_this.F2(), (_index221).Int())
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index222
				(_1202_v0).ArraySet1(_1332_v76, (_index222).Int())
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index223
				(_1202_v0).ArraySet1(_1332_v76, (_index223).Int())
				var _1340_v84 D4
				_ = _1340_v84
				_1340_v84 = Companion_D4_.Create_DC11_(_dafny.SeqOf((_1206_v4.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_1206_v4.F2()), 0))).Int()).(bool)))
				var _pat_let_tv17 = _1239_v33
				_ = _pat_let_tv17
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(308), _dafny.ArrayLen((_1339_v83), 0))
				_ = _index224
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index225
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1202_v0), 0))
				_ = _index226
				var _rhs160 _dafny.Array = _1206_v4.F2()
				_ = _rhs160
				var _rhs161 _dafny.Int = _1332_v76
				_ = _rhs161
				var _rhs162 _dafny.Int = _dafny.IntOfUint32(((func(_pat_let35_0 D4) D4 {
					return func(_1341_dt__update__tmp_h2 D4) D4 {
						return func(_pat_let36_0 _dafny.Sequence) D4 {
							return func(_1342_dt__update_hcf28_h0 _dafny.Sequence) D4 {
								return Companion_D4_.Create_DC11_(_1342_dt__update_hcf28_h0)
							}(_pat_let36_0)
						}(_pat_let_tv17)
					}(_pat_let35_0)
				}(_1340_v84)).Dtor_cf28()).Cardinality())
				_ = _rhs162
				var _rhs163 _dafny.Sequence = _1207_v5
				_ = _rhs163
				var _rhs164 _dafny.Int = p0
				_ = _rhs164
				var _lhs106 _dafny.Array = _1339_v83
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(308), _dafny.ArrayLen((_1339_v83), 0))
				_ = _lhs107
				var _lhs108 _dafny.Array = _1202_v0
				_ = _lhs108
				var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1202_v0), 0))
				_ = _lhs109
				var _lhs110 _dafny.Array = _1202_v0
				_ = _lhs110
				var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1202_v0), 0))
				_ = _lhs111
				(_lhs106).ArraySet1(_rhs160, (_lhs107).Int())
				(_lhs108).ArraySet1(_rhs161, (_lhs109).Int())
				_1332_v76 = _rhs162
				_1207_v5 = _rhs163
				(_lhs110).ArraySet1(_rhs164, (_lhs111).Int())
			} else {
				(_1206_v4).F2_set_(_1206_v4.F2())
				var _1343_v85 _dafny.Map
				_ = _1343_v85
				_1343_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1205_v3)
				var _1344_v86 _dafny.Map
				_ = _1344_v86
				_1344_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				_1343_v85 = (_1343_v85).Update((func() _dafny.Int {
					if (_1344_v86).Contains(p0) {
						return (_1344_v86).Get(p0).(_dafny.Int)
					}
					return p0
				})(), _1205_v3)
				var _arr47 _dafny.Array = _this.F2()
				_ = _arr47
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index227
				(_arr47).ArraySet1(_1205_v3, (_index227).Int())
				var _arr48 _dafny.Array = _this.F2()
				_ = _arr48
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index228
				(_arr48).ArraySet1(_1205_v3, (_index228).Int())
				var _1345_v87 _dafny.Int
				_ = _1345_v87
				_1345_v87 = _dafny.IntOfInt64(427)
				var _1346_v88 _dafny.Map
				_ = _1346_v88
				_1346_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), _1345_v87)
				var _1347_v89 _dafny.Sequence
				_ = _1347_v89
				_1347_v89 = _dafny.SeqOf(p0, _1345_v87, _1345_v87, _dafny.IntOfInt64(-288), _1345_v87)
				_1345_v87 = (func() _dafny.Int {
					if (_1346_v88).Contains(_1205_v3) {
						return (_1346_v88).Get(_1205_v3).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_1347_v89).Cardinality())
				})()
				var _arr49 _dafny.Array = _this.F2()
				_ = _arr49
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index229
				(_arr49).ArraySet1(_1205_v3, (_index229).Int())
			}
			var _1348_v90 _dafny.Array
			_ = _1348_v90
			var _len0_41 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_41
			var _nw188 _dafny.Array
			_ = _nw188
			if _len0_41.Cmp(_dafny.Zero) == 0 {
				_nw188 = _dafny.NewArray(_len0_41)
			} else {
				var _init41 func(_dafny.Int) _dafny.Sequence = (func(_1349_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1350_i10 _dafny.Int) _dafny.Sequence {
						return _1349_v5
					}
				})(_1207_v5)
				_ = _init41
				var _element0_41 = _init41(_dafny.Zero)
				_ = _element0_41
				_nw188 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
				(_nw188).ArraySet1(_element0_41, 0)
				var _nativeLen0_41 = (_len0_41).Int()
				_ = _nativeLen0_41
				for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
					(_nw188).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
				}
			}
			_1348_v90 = _nw188
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_1348_v90), 0))
			_ = _index230
			(_1348_v90).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("nrklejhh"), (_index230).Int())
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_1348_v90), 0))
			_ = _index231
			(_1348_v90).ArraySet1(_1207_v5, (_index231).Int())
			r0 = (_1330_v74).Fm7(true, p0, !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(691))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}((func(_1351_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1352_i11 _dafny.Int) _dafny.CodePoint {
					return _1351_v7
				}
			})(_1209_v7))), (_1348_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_1348_v90), 0))).Int()).(_dafny.Sequence)), Companion_Default___.SafeDivisionInt(p0, p0), globalState)
		} else {
			var _1353___mcc_h11 D8 = _source16.Get_().(D8_DC29).Cf55
			_ = _1353___mcc_h11
			var _1354_cf55 D8 = _1353___mcc_h11
			_ = _1354_cf55
			var _1355_v91 _dafny.Int
			_ = _1355_v91
			_1355_v91 = _dafny.IntOfInt64(730)
			var _1356_v92 D5
			_ = _1356_v92
			_1356_v92 = Companion_D5_.Create_DC17_(p0)
			_1355_v91 = (_1356_v92).Dtor_cf35()
			var _1357_v93 _dafny.Map
			_ = _1357_v93
			_1357_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1205_v3), _1209_v7)
			var _1358_v94 D3
			_ = _1358_v94
			_1358_v94 = Companion_D3_.Create_DC10_(_1205_v3, (func() _dafny.CodePoint {
				if (_1357_v93).Contains(_1205_v3) {
					return (_1357_v93).Get(_1205_v3).(_dafny.CodePoint)
				}
				return _1209_v7
			})(), _1205_v3, _dafny.Companion_Sequence_.IsProperPrefixOf(_1207_v5, _dafny.UnicodeSeqOfUtf8Bytes("ligdce")), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("drun"), _1207_v5))
			var _1359_v95 _dafny.MultiSet
			_ = _1359_v95
			_1359_v95 = _dafny.MultiSetOf(p0, _1355_v91, _1355_v91)
			_1358_v94 = Companion_D3_.Create_DC10_(!((_1355_v91).Cmp(Companion_Default___.Fm0((func() _dafny.Int {
				if (_1359_v95).Contains(p0) {
					return (_1359_v95).Multiplicity(p0)
				}
				return _1355_v91
			})(), globalState)) > 0), _1209_v7, Companion_Default___.Fm20(_1205_v3, globalState), _1205_v3, Companion_Default___.Fm20(_1205_v3, globalState))
			var _arr50 _dafny.Array = _1206_v4.F2()
			_ = _arr50
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_1206_v4.F2()), 0))
			_ = _index232
			(_arr50).ArraySet1(Companion_Default___.Fm20(false, globalState), (_index232).Int())
			var _1360_v96 _dafny.Sequence
			_ = _1360_v96
			_1360_v96 = _dafny.SeqOf(_1202_v0)
			var _1361_v97 _dafny.Array
			_ = _1361_v97
			var _nwElement0_37 _dafny.Sequence = (func() _dafny.Sequence {
				if _1205_v3 {
					return _dafny.SeqOf(_1202_v0, _1202_v0)
				}
				return _1360_v96
			})()
			_ = _nwElement0_37
			var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(2))
			_ = _nw189
			(_nw189).ArraySet1(_nwElement0_37, 0)
			(_nw189).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1360_v96, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1360_v96).Cardinality()))).Uint32(), _1202_v0), _dafny.SeqOf(_1202_v0, _1202_v0, _1202_v0)), 1)
			_1361_v97 = _nw189
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_1361_v97), 0))
			_ = _index233
			(_1361_v97).ArraySet1(_1360_v96, (_index233).Int())
			var _1362_v98 _dafny.Map
			_ = _1362_v98
			_1362_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1202_v0)
			var _1363_v99 _dafny.Sequence
			_ = _1363_v99
			_1363_v99 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1360_v96, _dafny.SeqOf((func() _dafny.Array {
				if (_1362_v98).Contains(_1355_v91) {
					return (_1362_v98).Get(_1355_v91).(_dafny.Array)
				}
				return _1202_v0
			})())), _1360_v96)
			var _arr51 _dafny.Array = _1206_v4.F2()
			_ = _arr51
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_1206_v4.F2()), 0))
			_ = _index234
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_1361_v97), 0))
			_ = _index235
			var _rhs165 bool = _1205_v3
			_ = _rhs165
			var _rhs166 _dafny.Sequence = (_1363_v99).Select((Companion_Default___.SafeIndex(_1355_v91, _dafny.IntOfUint32((_1363_v99).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs166
			var _rhs167 _dafny.Map = _1240_v34
			_ = _rhs167
			var _lhs112 _dafny.Array = _1206_v4.F2()
			_ = _lhs112
			var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_1206_v4.F2()), 0))
			_ = _lhs113
			var _lhs114 _dafny.Array = _1361_v97
			_ = _lhs114
			var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_1361_v97), 0))
			_ = _lhs115
			(_lhs112).ArraySet1(_rhs165, (_lhs113).Int())
			(_lhs114).ArraySet1(_rhs166, (_lhs115).Int())
			_1240_v34 = _rhs167
			var _arr52 _dafny.Array = _1206_v4.F2()
			_ = _arr52
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(253), _dafny.ArrayLen((_1206_v4.F2()), 0))
			_ = _index236
			(_arr52).ArraySet1(_1205_v3, (_index236).Int())
		}
		var _1364_v100 _dafny.MultiSet
		_ = _1364_v100
		_1364_v100 = _dafny.MultiSetOf((func() _dafny.Array {
			if (p2).Contains(_1205_v3) {
				return (p2).Get(_1205_v3).(_dafny.Array)
			}
			return _this.F2()
		})())
		var _1365_v101 D7
		_ = _1365_v101
		_1365_v101 = Companion_D7_.Create_DC22_(_1364_v100)
		var _source18 D7 = _1365_v101
		_ = _source18
		if _source18.Is_DC23() {
			var _1366___mcc_h17 _dafny.Int = _source18.Get_().(D7_DC23).Cf42
			_ = _1366___mcc_h17
			var _1367_cf42 _dafny.Int = _1366___mcc_h17
			_ = _1367_cf42
			var _1368_v102 _dafny.Sequence
			_ = _1368_v102
			_1368_v102 = _dafny.SeqOf(_1239_v33)
			var _1369_v103 _dafny.MultiSet
			_ = _1369_v103
			_1369_v103 = _dafny.MultiSetOf(false)
			var _1370_v104 _dafny.Sequence
			_ = _1370_v104
			_1370_v104 = _dafny.SeqOf(_1207_v5)
			var _1371_v105 _dafny.Sequence
			_ = _1371_v105
			_1371_v105 = _dafny.SeqOf(_1367_cf42, _1367_cf42)
			var _1372_v106 D13
			_ = _1372_v106
			_1372_v106 = Companion_D13_.Create_DC42_((_1368_v102).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_1367_cf42)).Cardinality()), _dafny.IntOfUint32((_1368_v102).Cardinality()))).Uint32()).(_dafny.Sequence), _1369_v103, _1370_v104, _1205_v3, _1371_v105)
			_1367_cf42 = ((_1372_v106).Dtor_cf78()).Cardinality()
			_1367_cf42 = _1367_cf42
			var _source19 D8 = Companion_Default___.Fm33(_1371_v105, p0, globalState)
			_ = _source19
			if _source19.Is_DC26() {
				var _1373___mcc_h20 bool = _source19.Get_().(D8_DC26).Cf45
				_ = _1373___mcc_h20
				var _1374___mcc_h21 bool = _source19.Get_().(D8_DC26).Cf46
				_ = _1374___mcc_h21
				var _1375___mcc_h22 bool = _source19.Get_().(D8_DC26).Cf47
				_ = _1375___mcc_h22
				var _1376___mcc_h23 _dafny.Int = _source19.Get_().(D8_DC26).Cf48
				_ = _1376___mcc_h23
				var _1377_cf48 _dafny.Int = _1376___mcc_h23
				_ = _1377_cf48
				var _1378_cf47 bool = _1375___mcc_h22
				_ = _1378_cf47
				var _1379_cf46 bool = _1374___mcc_h21
				_ = _1379_cf46
				var _1380_cf45 bool = _1373___mcc_h20
				_ = _1380_cf45
				var _arr53 _dafny.Array = _1206_v4.F2()
				_ = _arr53
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1206_v4.F2()), 0))
				_ = _index237
				(_arr53).ArraySet1((func() bool {
					if _1205_v3 {
						return _1380_cf45
					}
					return _1380_cf45
				})(), (_index237).Int())
				var _arr54 _dafny.Array = _1206_v4.F2()
				_ = _arr54
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1206_v4.F2()), 0))
				_ = _index238
				var _rhs168 bool = (_1239_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1207_v5).Cardinality()), _dafny.IntOfUint32((_1239_v33).Cardinality()))).Uint32()).(bool)
				_ = _rhs168
				var _rhs169 bool = !(_1380_cf45) || (_1205_v3)
				_ = _rhs169
				var _rhs170 T0 = _1206_v4
				_ = _rhs170
				var _lhs116 _dafny.Array = _1206_v4.F2()
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1206_v4.F2()), 0))
				_ = _lhs117
				_1379_cf46 = _rhs168
				(_lhs116).ArraySet1(_rhs169, (_lhs117).Int())
				_1206_v4 = _rhs170
				_1379_cf46 = !(_1205_v3)
				_1380_cf45 = !(false)
				r0 = _1380_cf45
			} else if _source19.Is_DC27() {
				var _1381___mcc_h24 bool = _source19.Get_().(D8_DC27).Cf49
				_ = _1381___mcc_h24
				var _1382___mcc_h25 bool = _source19.Get_().(D8_DC27).Cf50
				_ = _1382___mcc_h25
				var _1383___mcc_h26 _dafny.Int = _source19.Get_().(D8_DC27).Cf51
				_ = _1383___mcc_h26
				var _1384___mcc_h27 _dafny.Int = _source19.Get_().(D8_DC27).Cf52
				_ = _1384___mcc_h27
				var _1385_cf52 _dafny.Int = _1384___mcc_h27
				_ = _1385_cf52
				var _1386_cf51 _dafny.Int = _1383___mcc_h26
				_ = _1386_cf51
				var _1387_cf50 bool = _1382___mcc_h25
				_ = _1387_cf50
				var _1388_cf49 bool = _1381___mcc_h24
				_ = _1388_cf49
				_1388_cf49 = _1388_cf49
				_1205_v3 = _1205_v3
				var _rhs171 _dafny.Int = _dafny.IntOfUint32((_1371_v105).Cardinality())
				_ = _rhs171
				var _rhs172 bool = (_1369_v103).IsProperSubsetOf(_dafny.MultiSetOf(false))
				_ = _rhs172
				var _rhs173 _dafny.Int = (_dafny.Zero).Minus(_1367_cf42)
				_ = _rhs173
				_1385_cf52 = _rhs171
				_1388_cf49 = _rhs172
				_1385_cf52 = _rhs173
				var _1389_v107 _dafny.MultiSet
				_ = _1389_v107
				_1389_v107 = _dafny.MultiSetOf(_1369_v103, _1369_v103)
				var _1390_v108 _dafny.Map
				_ = _1390_v108
				_1390_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_v7, _dafny.MultiSetOf(_dafny.MultiSetOf(_1387_cf50, _1387_cf50, _1388_cf49, _1205_v3)))
				_1389_v107 = ((func() _dafny.MultiSet {
					if (_1390_v108).Contains(_1209_v7) {
						return (_1390_v108).Get(_1209_v7).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(_dafny.SeqOf(_1369_v103, _1369_v103))
				})()).Union((_1389_v107).Update(_dafny.MultiSetOf(_1387_cf50), Companion_Default___.Abs(_1367_cf42)))
			} else if _source19.Is_DC28() {
				var _1391___mcc_h28 bool = _source19.Get_().(D8_DC28).Cf53
				_ = _1391___mcc_h28
				var _1392___mcc_h29 _dafny.Int = _source19.Get_().(D8_DC28).Cf54
				_ = _1392___mcc_h29
				var _1393_cf54 _dafny.Int = _1392___mcc_h29
				_ = _1393_cf54
				var _1394_cf53 bool = _1391___mcc_h28
				_ = _1394_cf53
				_1205_v3 = _dafny.Companion_Sequence_.IsPrefixOf(_1239_v33, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, _1205_v3, _1205_v3), _1239_v33))
				var _1395_v109 _dafny.Set
				_ = _1395_v109
				_1395_v109 = _dafny.SetOf(p0, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1202_v0, _dafny.IntOfUint32((_1207_v5).Cardinality()))).Merge(_1203_v1)).Cardinality(), (_dafny.Zero).Minus(_1367_cf42))
				var _rhs174 bool = _1205_v3
				_ = _rhs174
				var _rhs175 bool = _1394_cf53
				_ = _rhs175
				var _rhs176 _dafny.Set = _1395_v109
				_ = _rhs176
				var _rhs177 bool = _1394_cf53
				_ = _rhs177
				_1205_v3 = _rhs174
				_1394_cf53 = _rhs175
				_1395_v109 = _rhs176
				_1205_v3 = _rhs177
				_1367_cf42 = _1367_cf42
				var _1396_v110 _dafny.Map
				_ = _1396_v110
				_1396_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1205_v3, _1207_v5)
				var _rhs178 bool = _1394_cf53
				_ = _rhs178
				var _rhs179 bool = ((_1393_cf54).Times(_dafny.IntOfUint32((_1371_v105).Cardinality()))).Cmp(((_1240_v34).Update((_1239_v33).Select((Companion_Default___.SafeIndex((_1396_v110).Cardinality(), _dafny.IntOfUint32((_1239_v33).Cardinality()))).Uint32()).(bool), !(_1394_cf53))).Cardinality()) == 0
				_ = _rhs179
				r0 = _rhs178
				_1205_v3 = _rhs179
			} else if _source19.Is_DC25() {
				var _1397___mcc_h30 _dafny.Map = _source19.Get_().(D8_DC25).Cf44
				_ = _1397___mcc_h30
				var _1398_cf44 _dafny.Map = _1397___mcc_h30
				_ = _1398_cf44
				_1202_v0 = _1202_v0
				_1205_v3 = _1205_v3
				var _1399_v111 D3
				_ = _1399_v111
				_1399_v111 = Companion_D3_.Create_DC10_(!(_1205_v3), _1209_v7, _1205_v3, _1205_v3, true)
				var _1400_v112 _dafny.MultiSet
				_ = _1400_v112
				_1400_v112 = _dafny.MultiSetOf(_1399_v111)
				var _1401_v113 _dafny.Sequence
				_ = _1401_v113
				_1401_v113 = _dafny.SeqOf(_1399_v111, _1399_v111, _1399_v111, _1399_v111, Companion_D3_.Create_DC10_(_1205_v3, _1209_v7, _1205_v3, Companion_Default___.Fm20(_1205_v3, globalState), _1205_v3))
				var _1402_v114 _dafny.Array
				_ = _1402_v114
				var _nwElement0_38 _dafny.MultiSet = _1400_v112
				_ = _nwElement0_38
				var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(3))
				_ = _nw190
				(_nw190).ArraySet1(_nwElement0_38, 0)
				(_nw190).ArraySet1((_1400_v112).Difference(_1400_v112), 1)
				(_nw190).ArraySet1((_dafny.MultiSetFromSeq(_1401_v113)).Union(_1400_v112), 2)
				_1402_v114 = _nw190
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_1402_v114), 0))
				_ = _index239
				(_1402_v114).ArraySet1(_1400_v112, (_index239).Int())
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_1402_v114), 0))
				_ = _index240
				(_1402_v114).ArraySet1((_1400_v112).Union(_1400_v112), (_index240).Int())
				var _1403_v115 D4
				_ = _1403_v115
				_1403_v115 = Companion_D4_.Create_DC11_(_1239_v33)
				var _1404_v116 D5
				_ = _1404_v116
				_1404_v116 = Companion_D5_.Create_DC15_(_1403_v115, _1205_v3)
				var _1405_v117 _dafny.Sequence
				_ = _1405_v117
				_1405_v117 = _dafny.SeqOf(_1404_v116)
				var _rhs180 bool = (_1367_cf42).Cmp(Companion_Default___.SafeDivisionInt(p0, p0)) != 0
				_ = _rhs180
				var _rhs181 bool = !(_1205_v3)
				_ = _rhs181
				var _rhs182 bool = Companion_Default___.Fm20(_dafny.Companion_Sequence_.Contains(_1405_v117, _1404_v116), globalState)
				_ = _rhs182
				_1205_v3 = _rhs180
				_1205_v3 = _rhs181
				r0 = _rhs182
			} else {
				var _1406___mcc_h31 D8 = _source19.Get_().(D8_DC29).Cf55
				_ = _1406___mcc_h31
				var _1407_cf55 D8 = _1406___mcc_h31
				_ = _1407_cf55
				r0 = _1205_v3
				_1367_cf42 = Companion_Default___.SafeModuloInt(p0, p0)
				var _1408_v118 _dafny.Map
				_ = _1408_v118
				_1408_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1205_v3, p0)
				_1408_v118 = (_1408_v118).Update(_1205_v3, Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mgnbqbnka")).Cardinality())))
				var _1409_v119 _dafny.MultiSet
				_ = _1409_v119
				_1409_v119 = _dafny.MultiSetOf(_1206_v4, _1206_v4)
				var _1410_v120 _dafny.Set
				_ = _1410_v120
				_1410_v120 = _dafny.SetOf(_1409_v119)
				var _1411_v121 _dafny.Map
				_ = _1411_v121
				_1411_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1410_v120, _1369_v103)
				_1411_v121 = (_1411_v121).Update(_1410_v120, _dafny.MultiSetFromSeq((func() _dafny.Sequence {
					if _1205_v3 {
						return _dafny.SeqOf(_1205_v3, _1205_v3)
					}
					return _1239_v33
				})()))
			}
			var _rhs183 _dafny.Array = _this.F2()
			_ = _rhs183
			var _rhs184 bool = (true) && (_1205_v3)
			_ = _rhs184
			var _rhs185 _dafny.Array = _1202_v0
			_ = _rhs185
			var _lhs118 T0 = _1206_v4
			_ = _lhs118
			_lhs118.F2_set_(_rhs183)
			_1205_v3 = _rhs184
			_1202_v0 = _rhs185
		} else if _source18.Is_DC24() {
			var _1412___mcc_h18 bool = _source18.Get_().(D7_DC24).Cf43
			_ = _1412___mcc_h18
			var _1413_cf43 bool = _1412___mcc_h18
			_ = _1413_cf43
			var _1414_v122 *C0
			_ = _1414_v122
			var _nw191 *C0 = New_C0_()
			_ = _nw191
			_nw191.Ctor__()
			_1414_v122 = _nw191
			var _1415_v123 *C2
			_ = _1415_v123
			var _nw192 *C2 = New_C2_()
			_ = _nw192
			_nw192.Ctor__(_1205_v3, p0, _1204_v2, _1413_cf43, _1206_v4.F2())
			_1415_v123 = _nw192
			_1415_v123 = _1415_v123
			var _1416_v124 D2
			_ = _1416_v124
			_1416_v124 = Companion_D2_.Create_DC6_(_1209_v7, (_1415_v123).F8())
			var _1417_v125 _dafny.Set
			_ = _1417_v125
			_1417_v125 = _dafny.SetOf(_1209_v7, (_1416_v124).Dtor_cf14(), _1209_v7, _dafny.CodePoint('b'))
			(_1415_v123).F7 = (_dafny.SetOf((_1207_v5).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1207_v5).Cardinality()))).Uint32()).(_dafny.CodePoint))).IsSubsetOf(_1417_v125)
			(_1415_v123).F7 = !((_1413_cf43) == (_1413_cf43))
		} else {
			var _1418___mcc_h19 _dafny.MultiSet = _source18.Get_().(D7_DC22).Cf41
			_ = _1418___mcc_h19
			var _1419_cf41 _dafny.MultiSet = _1418___mcc_h19
			_ = _1419_cf41
			_1205_v3 = _1205_v3
			var _1420_v126 T1
			_ = _1420_v126
			var _nw193 *C2 = New_C2_()
			_ = _nw193
			_nw193.Ctor__(_1205_v3, (_this).Fm2(globalState), _1204_v2, _1205_v3, _1206_v4.F2())
			_1420_v126 = _nw193
			var _1421_v127 _dafny.Array
			_ = _1421_v127
			var _nwElement0_39 T1 = _1420_v126
			_ = _nwElement0_39
			var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(13))
			_ = _nw194
			(_nw194).ArraySet1(_nwElement0_39, 0)
			(_nw194).ArraySet1(_1420_v126, 1)
			(_nw194).ArraySet1(_1420_v126, 2)
			(_nw194).ArraySet1(_1420_v126, 3)
			(_nw194).ArraySet1(_1420_v126, 4)
			(_nw194).ArraySet1(_1420_v126, 5)
			(_nw194).ArraySet1(_1420_v126, 6)
			(_nw194).ArraySet1(_1420_v126, 7)
			(_nw194).ArraySet1(_1420_v126, 8)
			(_nw194).ArraySet1(_1420_v126, 9)
			(_nw194).ArraySet1(_1420_v126, 10)
			(_nw194).ArraySet1(_1420_v126, 11)
			(_nw194).ArraySet1(_1420_v126, 12)
			_1421_v127 = _nw194
			var _1422_v128 _dafny.Sequence
			_ = _1422_v128
			_1422_v128 = _dafny.SeqOf(Companion_Default___.Fm34(_1205_v3, p0, globalState))
			var _1423_v129 _dafny.Map
			_ = _1423_v129
			_1423_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1421_v127, (_1422_v128).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1422_v128).Cardinality()))).Uint32()).(_dafny.MultiSet))
			var _1424_v130 _dafny.Map
			_ = _1424_v130
			_1424_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1421_v127)
			var _1425_v131 _dafny.Map
			_ = _1425_v131
			_1425_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_v7, p0)
			var _1426_v132 _dafny.Map
			_ = _1426_v132
			_1426_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.CodePoint('p'))
			var _1427_v133 _dafny.Sequence
			_ = _1427_v133
			_1427_v133 = _dafny.SeqOf(((_1425_v131).Update(_1209_v7, ((_1426_v132).Update(p0, _1209_v7)).Cardinality())).Cardinality())
			var _1428_v134 _dafny.MultiSet
			_ = _1428_v134
			_1428_v134 = _dafny.MultiSetOf(_1427_v133, _1427_v133, _dafny.SeqOf(p0), _1427_v133, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(968))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}((func(_1429_v33 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1430_i12 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_1429_v33).Cardinality())
				}
			})(_1239_v33))))
			_1205_v3 = ((func() _dafny.MultiSet {
				if (_1423_v129).Contains((func() _dafny.Array {
					if (_1424_v130).Contains(p0) {
						return (_1424_v130).Get(p0).(_dafny.Array)
					}
					return _1421_v127
				})()) {
					return (_1423_v129).Get((func() _dafny.Array {
						if (_1424_v130).Contains(p0) {
							return (_1424_v130).Get(p0).(_dafny.Array)
						}
						return _1421_v127
					})()).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(_dafny.SeqOf(p0))
			})()).Equals(_1428_v134)
			_1420_v126 = _1420_v126
			var _1431_v135 *C0
			_ = _1431_v135
			var _nw195 *C0 = New_C0_()
			_ = _nw195
			_nw195.Ctor__()
			_1431_v135 = _nw195
			var _1432_v136 _dafny.Map
			_ = _1432_v136
			_1432_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1205_v3, _1431_v135)
			var _1433_v137 _dafny.MultiSet
			_ = _1433_v137
			_1433_v137 = _dafny.MultiSetOf((_1432_v136).Cardinality())
			var _1434_v138 _dafny.MultiSet
			_ = _1434_v138
			_1434_v138 = _dafny.MultiSetOf(_1433_v137)
			var _1435_v139 D2
			_ = _1435_v139
			_1435_v139 = Companion_D2_.Create_DC7_(p0, p0, _1434_v138, p0, p0)
			var _1436_v140 D2
			_ = _1436_v140
			_1436_v140 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC8_(_1435_v139))
			var _pat_let_tv18 = _1209_v7
			_ = _pat_let_tv18
			var _pat_let_tv19 = p0
			_ = _pat_let_tv19
			var _1437_v141 _dafny.Map
			_ = _1437_v141
			_1437_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, func(_pat_let37_0 D2) D2 {
				return func(_1438_dt__update__tmp_h3 D2) D2 {
					return func(_pat_let38_0 D2) D2 {
						return func(_1439_dt__update_hcf21_h0 D2) D2 {
							return Companion_D2_.Create_DC8_(_1439_dt__update_hcf21_h0)
						}(_pat_let38_0)
					}(Companion_D2_.Create_DC8_(Companion_D2_.Create_DC6_(_pat_let_tv18, _pat_let_tv19)))
				}(_pat_let37_0)
			}(_1436_v140))
			_1437_v141 = (_1437_v141).Update(p0, _1436_v140)
		}
		var _1440_v142 D4
		_ = _1440_v142
		_1440_v142 = Companion_D4_.Create_DC12_()
		var _pat_let_tv20 = _1240_v34
		_ = _pat_let_tv20
		var _pat_let_tv21 = _1205_v3
		_ = _pat_let_tv21
		var _pat_let_tv22 = _1240_v34
		_ = _pat_let_tv22
		var _pat_let_tv23 = _1205_v3
		_ = _pat_let_tv23
		var _pat_let_tv24 = _1205_v3
		_ = _pat_let_tv24
		var _pat_let_tv25 = _1205_v3
		_ = _pat_let_tv25
		r0 = func(_source20 D4) bool {
			if _source20.Is_DC12() {
				if (_pat_let_tv20).Contains(_pat_let_tv21) {
					return (_pat_let_tv22).Get(_pat_let_tv23).(bool)
				} else {
					return !(_pat_let_tv24)
				}
			} else if _source20.Is_DC13() {
				var _1441___mcc_h32 bool = _source20.Get_().(D4_DC13).Cf29
				_ = _1441___mcc_h32
				var _1442___mcc_h33 _dafny.MultiSet = _source20.Get_().(D4_DC13).Cf30
				_ = _1442___mcc_h33
				var _1443___mcc_h34 bool = _source20.Get_().(D4_DC13).Cf31
				_ = _1443___mcc_h34
				var _1444_cf31 bool = _1443___mcc_h34
				_ = _1444_cf31
				var _1445_cf30 _dafny.MultiSet = _1442___mcc_h33
				_ = _1445_cf30
				var _1446_cf29 bool = _1441___mcc_h32
				_ = _1446_cf29
				return false
			} else {
				var _1447___mcc_h35 _dafny.Sequence = _source20.Get_().(D4_DC11).Cf28
				_ = _1447___mcc_h35
				var _1448_cf28 _dafny.Sequence = _1447___mcc_h35
				_ = _1448_cf28
				return _pat_let_tv25
			}
		}(_1440_v142)
		return r0
	}
}
func (_this *C3) M1(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Array, p3 _dafny.Array, globalState *GlobalState) (_dafny.MultiSet, bool) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 bool = false
		_ = r1
		var _1449_v0 _dafny.Int
		_ = _1449_v0
		_1449_v0 = _dafny.IntOfInt64(980)
		var _1450_v1 bool
		_ = _1450_v1
		_1450_v1 = false
		var _1451_v2 _dafny.Sequence
		_ = _1451_v2
		_1451_v2 = _dafny.SeqOf(_1450_v1)
		_1449_v0 = (_1449_v0).Times((func() _dafny.Int {
			if false {
				return _1449_v0
			}
			return _dafny.IntOfUint32((_1451_v2).Cardinality())
		})())
		var _1452_v3 D4
		_ = _1452_v3
		_1452_v3 = Companion_D4_.Create_DC11_(_1451_v2)
		var _1453_v4 D5
		_ = _1453_v4
		_1453_v4 = Companion_D5_.Create_DC15_(_1452_v3, _1450_v1)
		var _pat_let_tv26 = _1451_v2
		_ = _pat_let_tv26
		var _pat_let_tv27 = _1450_v1
		_ = _pat_let_tv27
		var _pat_let_tv28 = p1
		_ = _pat_let_tv28
		var _pat_let_tv29 = _1451_v2
		_ = _pat_let_tv29
		var _pat_let_tv30 = p1
		_ = _pat_let_tv30
		var _pat_let_tv31 = _1450_v1
		_ = _pat_let_tv31
		var _pat_let_tv32 = _1450_v1
		_ = _pat_let_tv32
		var _pat_let_tv33 = _1450_v1
		_ = _pat_let_tv33
		var _pat_let_tv34 = _1451_v2
		_ = _pat_let_tv34
		var _pat_let_tv35 = _1450_v1
		_ = _pat_let_tv35
		var _pat_let_tv36 = _1451_v2
		_ = _pat_let_tv36
		var _pat_let_tv37 = _1450_v1
		_ = _pat_let_tv37
		var _pat_let_tv38 = _1450_v1
		_ = _pat_let_tv38
		var _pat_let_tv39 = _1450_v1
		_ = _pat_let_tv39
		var _pat_let_tv40 = _1450_v1
		_ = _pat_let_tv40
		var _pat_let_tv41 = p1
		_ = _pat_let_tv41
		var _pat_let_tv42 = _1450_v1
		_ = _pat_let_tv42
		var _pat_let_tv43 = _1450_v1
		_ = _pat_let_tv43
		var _pat_let_tv44 = _1451_v2
		_ = _pat_let_tv44
		var _pat_let_tv45 = _1450_v1
		_ = _pat_let_tv45
		var _pat_let_tv46 = p1
		_ = _pat_let_tv46
		var _pat_let_tv47 = _1450_v1
		_ = _pat_let_tv47
		var _pat_let_tv48 = _1450_v1
		_ = _pat_let_tv48
		var _pat_let_tv49 = _1450_v1
		_ = _pat_let_tv49
		var _pat_let_tv50 = _1450_v1
		_ = _pat_let_tv50
		var _pat_let_tv51 = _1450_v1
		_ = _pat_let_tv51
		var _pat_let_tv52 = p1
		_ = _pat_let_tv52
		var _pat_let_tv53 = _1450_v1
		_ = _pat_let_tv53
		var _pat_let_tv54 = _1450_v1
		_ = _pat_let_tv54
		var _pat_let_tv55 = _1450_v1
		_ = _pat_let_tv55
		var _pat_let_tv56 = _1450_v1
		_ = _pat_let_tv56
		var _pat_let_tv57 = _1450_v1
		_ = _pat_let_tv57
		var _pat_let_tv58 = _1451_v2
		_ = _pat_let_tv58
		var _pat_let_tv59 = _1449_v0
		_ = _pat_let_tv59
		var _pat_let_tv60 = _1451_v2
		_ = _pat_let_tv60
		var _pat_let_tv61 = p1
		_ = _pat_let_tv61
		var _pat_let_tv62 = _1450_v1
		_ = _pat_let_tv62
		var _pat_let_tv63 = _1450_v1
		_ = _pat_let_tv63
		var _pat_let_tv64 = _1450_v1
		_ = _pat_let_tv64
		var _pat_let_tv65 = _1450_v1
		_ = _pat_let_tv65
		var _pat_let_tv66 = _1450_v1
		_ = _pat_let_tv66
		var _pat_let_tv67 = _1450_v1
		_ = _pat_let_tv67
		var _pat_let_tv68 = _1450_v1
		_ = _pat_let_tv68
		var _pat_let_tv69 = _1450_v1
		_ = _pat_let_tv69
		var _pat_let_tv70 = p1
		_ = _pat_let_tv70
		var _pat_let_tv71 = _1450_v1
		_ = _pat_let_tv71
		var _pat_let_tv72 = _1450_v1
		_ = _pat_let_tv72
		var _pat_let_tv73 = _1450_v1
		_ = _pat_let_tv73
		var _pat_let_tv74 = _1451_v2
		_ = _pat_let_tv74
		var _pat_let_tv75 = _1450_v1
		_ = _pat_let_tv75
		var _pat_let_tv76 = p1
		_ = _pat_let_tv76
		var _pat_let_tv77 = _1450_v1
		_ = _pat_let_tv77
		var _pat_let_tv78 = _1450_v1
		_ = _pat_let_tv78
		var _pat_let_tv79 = _1450_v1
		_ = _pat_let_tv79
		var _pat_let_tv80 = _1450_v1
		_ = _pat_let_tv80
		var _pat_let_tv81 = _1450_v1
		_ = _pat_let_tv81
		var _pat_let_tv82 = _1450_v1
		_ = _pat_let_tv82
		var _pat_let_tv83 = _1451_v2
		_ = _pat_let_tv83
		var _pat_let_tv84 = _1450_v1
		_ = _pat_let_tv84
		var _pat_let_tv85 = p1
		_ = _pat_let_tv85
		var _pat_let_tv86 = _1450_v1
		_ = _pat_let_tv86
		var _pat_let_tv87 = _1450_v1
		_ = _pat_let_tv87
		var _pat_let_tv88 = _1450_v1
		_ = _pat_let_tv88
		var _pat_let_tv89 = _1450_v1
		_ = _pat_let_tv89
		var _pat_let_tv90 = _1450_v1
		_ = _pat_let_tv90
		var _pat_let_tv91 = _1451_v2
		_ = _pat_let_tv91
		var _pat_let_tv92 = _1449_v0
		_ = _pat_let_tv92
		var _pat_let_tv93 = _1451_v2
		_ = _pat_let_tv93
		var _pat_let_tv94 = _1449_v0
		_ = _pat_let_tv94
		var _pat_let_tv95 = _1450_v1
		_ = _pat_let_tv95
		var _pat_let_tv96 = p1
		_ = _pat_let_tv96
		var _pat_let_tv97 = _1450_v1
		_ = _pat_let_tv97
		var _pat_let_tv98 = _1450_v1
		_ = _pat_let_tv98
		_1449_v0 = (func(_source21 D5) _dafny.Map {
			if _source21.Is_DC15() {
				var _1454___mcc_h0 D4 = _source21.Get_().(D5_DC15).Cf33
				_ = _1454___mcc_h0
				var _1455___mcc_h1 bool = _source21.Get_().(D5_DC15).Cf34
				_ = _1455___mcc_h1
				var _1456_cf34 bool = _1455___mcc_h1
				_ = _1456_cf34
				var _1457_cf33 D4 = _1454___mcc_h0
				_ = _1457_cf33
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv26, _dafny.MultiSetOf(Companion_D3_.Create_DC10_(_pat_let_tv27, _pat_let_tv28, true, _1456_cf34, _1456_cf34)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv29, (Companion_D14_.Create_DC44_(_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D3_.Create_DC10_(true, _pat_let_tv30, _pat_let_tv31, _pat_let_tv32, _pat_let_tv33))))).Dtor_cf83()))
			} else if _source21.Is_DC16() {
				return func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter57 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv34, _pat_let_tv35)).Keys().Elements()); ; {
						_compr_50, _ok57 := _iter57()
						if !_ok57 {
							break
						}
						var _1458_v5 _dafny.Sequence
						_1458_v5 = interface{}(_compr_50).(_dafny.Sequence)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv36, _pat_let_tv37)).Contains(_1458_v5) {
							_coll50.Add(_1458_v5, _dafny.MultiSetOf(Companion_D3_.Create_DC10_(_pat_let_tv38, _dafny.CodePoint('p'), false, true, _pat_let_tv39), Companion_D3_.Create_DC10_(_pat_let_tv40, _pat_let_tv41, _pat_let_tv42, false, _pat_let_tv43)))
						}
					}
					return _coll50.ToMap()
				}()
			} else if _source21.Is_DC17() {
				var _1459___mcc_h2 _dafny.Int = _source21.Get_().(D5_DC17).Cf35
				_ = _1459___mcc_h2
				var _1460_cf35 _dafny.Int = _1459___mcc_h2
				_ = _1460_cf35
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv44, _dafny.MultiSetOf(Companion_D3_.Create_DC10_(_pat_let_tv45, _pat_let_tv46, _pat_let_tv47, _pat_let_tv48, _pat_let_tv49), Companion_D3_.Create_DC10_(_pat_let_tv50, (Companion_D3_.Create_DC10_(_pat_let_tv51, _pat_let_tv52, _pat_let_tv53, _pat_let_tv54, !(true))).Dtor_cf24(), _pat_let_tv55, _pat_let_tv56, _pat_let_tv57), Companion_D3_.Create_DC10_((_pat_let_tv58).Select((Companion_Default___.SafeIndex(_pat_let_tv59, _dafny.IntOfUint32((_pat_let_tv60).Cardinality()))).Uint32()).(bool), _pat_let_tv61, _pat_let_tv62, _pat_let_tv63, _pat_let_tv64), Companion_D3_.Create_DC10_(_pat_let_tv65, _dafny.CodePoint('y'), !(_pat_let_tv66), _pat_let_tv67, _pat_let_tv68), Companion_D3_.Create_DC10_(_pat_let_tv69, _pat_let_tv70, _pat_let_tv71, _pat_let_tv72, _pat_let_tv73)))
			} else if _source21.Is_DC14() {
				var _1461___mcc_h3 _dafny.Array = _source21.Get_().(D5_DC14).Cf32
				_ = _1461___mcc_h3
				var _1462_cf32 _dafny.Array = _1461___mcc_h3
				_ = _1462_cf32
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv74, _dafny.MultiSetOf(Companion_D3_.Create_DC10_(_pat_let_tv75, _pat_let_tv76, _pat_let_tv77, _pat_let_tv78, _pat_let_tv79), Companion_D3_.Create_DC10_(true, _dafny.CodePoint('o'), _pat_let_tv80, _pat_let_tv81, _pat_let_tv82)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv83, _dafny.MultiSetOf(Companion_D3_.Create_DC10_(_pat_let_tv84, _pat_let_tv85, false, _pat_let_tv86, _pat_let_tv87), Companion_D3_.Create_DC10_(_pat_let_tv88, _dafny.CodePoint('x'), true, _pat_let_tv89, _pat_let_tv90))))
			} else {
				var _1463___mcc_h4 D5 = _source21.Get_().(D5_DC18).Cf36
				_ = _1463___mcc_h4
				var _1464_cf36 D5 = _1463___mcc_h4
				_ = _1464_cf36
				return func() _dafny.Map {
					var _coll51 = _dafny.NewMapBuilder()
					_ = _coll51
					for _iter58 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv91, _pat_let_tv92)).Keys().Elements()); ; {
						_compr_51, _ok58 := _iter58()
						if !_ok58 {
							break
						}
						var _1465_v6 _dafny.Sequence
						_1465_v6 = interface{}(_compr_51).(_dafny.Sequence)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv93, _pat_let_tv94)).Contains(_1465_v6) {
							_coll51.Add(_1465_v6, _dafny.MultiSetOf(Companion_D3_.Create_DC10_(_pat_let_tv95, _pat_let_tv96, _pat_let_tv97, _pat_let_tv98, false)))
						}
					}
					return _coll51.ToMap()
				}()
			}
		}(_1453_v4)).Cardinality()
		var _1466_i0 _dafny.Int
		_ = _1466_i0
		_1466_i0 = _dafny.Zero
		{
			var _pat_let_tv99 = p0
			_ = _pat_let_tv99
			var _pat_let_tv100 = _1450_v1
			_ = _pat_let_tv100
			var _pat_let_tv101 = _1450_v1
			_ = _pat_let_tv101
			for func(_source22 D8) bool {
				if _source22.Is_DC26() {
					var _1485___mcc_h5 bool = _source22.Get_().(D8_DC26).Cf45
					_ = _1485___mcc_h5
					var _1486___mcc_h6 bool = _source22.Get_().(D8_DC26).Cf46
					_ = _1486___mcc_h6
					var _1487___mcc_h7 bool = _source22.Get_().(D8_DC26).Cf47
					_ = _1487___mcc_h7
					var _1488___mcc_h8 _dafny.Int = _source22.Get_().(D8_DC26).Cf48
					_ = _1488___mcc_h8
					var _1489_cf48 _dafny.Int = _1488___mcc_h8
					_ = _1489_cf48
					var _1490_cf47 bool = _1487___mcc_h7
					_ = _1490_cf47
					var _1491_cf46 bool = _1486___mcc_h6
					_ = _1491_cf46
					var _1492_cf45 bool = _1485___mcc_h5
					_ = _1492_cf45
					return _1492_cf45
				} else if _source22.Is_DC27() {
					var _1493___mcc_h9 bool = _source22.Get_().(D8_DC27).Cf49
					_ = _1493___mcc_h9
					var _1494___mcc_h10 bool = _source22.Get_().(D8_DC27).Cf50
					_ = _1494___mcc_h10
					var _1495___mcc_h11 _dafny.Int = _source22.Get_().(D8_DC27).Cf51
					_ = _1495___mcc_h11
					var _1496___mcc_h12 _dafny.Int = _source22.Get_().(D8_DC27).Cf52
					_ = _1496___mcc_h12
					var _1497_cf52 _dafny.Int = _1496___mcc_h12
					_ = _1497_cf52
					var _1498_cf51 _dafny.Int = _1495___mcc_h11
					_ = _1498_cf51
					var _1499_cf50 bool = _1494___mcc_h10
					_ = _1499_cf50
					var _1500_cf49 bool = _1493___mcc_h9
					_ = _1500_cf49
					return _1500_cf49
				} else if _source22.Is_DC28() {
					var _1501___mcc_h13 bool = _source22.Get_().(D8_DC28).Cf53
					_ = _1501___mcc_h13
					var _1502___mcc_h14 _dafny.Int = _source22.Get_().(D8_DC28).Cf54
					_ = _1502___mcc_h14
					var _1503_cf54 _dafny.Int = _1502___mcc_h14
					_ = _1503_cf54
					var _1504_cf53 bool = _1501___mcc_h13
					_ = _1504_cf53
					return (_dafny.IntOfInt64(208)).Cmp(_pat_let_tv99) != 0
				} else if _source22.Is_DC25() {
					var _1505___mcc_h15 _dafny.Map = _source22.Get_().(D8_DC25).Cf44
					_ = _1505___mcc_h15
					var _1506_cf44 _dafny.Map = _1505___mcc_h15
					_ = _1506_cf44
					return _pat_let_tv100
				} else {
					var _1507___mcc_h16 D8 = _source22.Get_().(D8_DC29).Cf55
					_ = _1507___mcc_h16
					var _1508_cf55 D8 = _1507___mcc_h16
					_ = _1508_cf55
					return _pat_let_tv101
				}
			}(Companion_D8_.Create_DC28_(_1450_v1, Companion_Default___.Fm0(p0, globalState))) {
				{
					if (_1466_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1466_i0 = (_1466_i0).Plus(_dafny.One)
					r1 = (_1451_v2).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1449_v0), _1449_v0), _dafny.IntOfUint32((_1451_v2).Cardinality()))).Uint32()).(bool)
					_1450_v1 = _1450_v1
					_1450_v1 = !(_1450_v1)
					if true {
						var _1467_v7 _dafny.CodePoint
						_ = _1467_v7
						_1467_v7 = _dafny.CodePoint('y')
						var _1468_v8 *C0
						_ = _1468_v8
						var _nw196 *C0 = New_C0_()
						_ = _nw196
						_nw196.Ctor__()
						_1468_v8 = _nw196
						var _1469_v9 _dafny.Map
						_ = _1469_v9
						_1469_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1468_v8, _1467_v7)
						_1467_v7 = (func() _dafny.CodePoint {
							if (_1469_v9).Contains(_1468_v8) {
								return (_1469_v9).Get(_1468_v8).(_dafny.CodePoint)
							}
							return _dafny.CodePoint('l')
						})()
						r1 = _1450_v1
						_1451_v2 = _1451_v2
						var _1470_v10 _dafny.Array
						_ = _1470_v10
						var _len0_42 _dafny.Int = _dafny.IntOfInt64(27)
						_ = _len0_42
						var _nw197 _dafny.Array
						_ = _nw197
						if _len0_42.Cmp(_dafny.Zero) == 0 {
							_nw197 = _dafny.NewArray(_len0_42)
						} else {
							var _init42 func(_dafny.Int) _dafny.Map = (func(_1471_v2 _dafny.Sequence, _1472_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
								return func(_1473_i1 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1471_v2, _1472_p0)
								}
							})(_1451_v2, p0)
							_ = _init42
							var _element0_42 = _init42(_dafny.Zero)
							_ = _element0_42
							_nw197 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
							(_nw197).ArraySet1(_element0_42, 0)
							var _nativeLen0_42 = (_len0_42).Int()
							_ = _nativeLen0_42
							for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
								(_nw197).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
							}
						}
						_1470_v10 = _nw197
						var _1474_v11 _dafny.Map
						_ = _1474_v11
						_1474_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1451_v2, p0)
						var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_1470_v10), 0))
						_ = _index241
						(_1470_v10).ArraySet1(_1474_v11, (_index241).Int())
						var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_1470_v10), 0))
						_ = _index242
						(_1470_v10).ArraySet1(_1474_v11, (_index242).Int())
						var _1475_v12 _dafny.Map
						_ = _1475_v12
						_1475_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1450_v1, _1450_v1)
						_1475_v12 = (_1475_v12).Update(_1450_v1, _1450_v1)
					} else {
						var _1476_v13 _dafny.MultiSet
						_ = _1476_v13
						_1476_v13 = _dafny.MultiSetOf(_1450_v1, _1450_v1)
						var _1477_v14 _dafny.Sequence
						_ = _1477_v14
						_1477_v14 = _dafny.SeqOf(_1449_v0, (func() _dafny.Int {
							if (_1476_v13).Contains(_1450_v1) {
								return (_1476_v13).Multiplicity(_1450_v1)
							}
							return (_this).Fm2(globalState)
						})(), p0, _dafny.IntOfUint32((_1451_v2).Cardinality()))
						var _1478_v15 _dafny.Sequence
						_ = _1478_v15
						_1478_v15 = _dafny.SeqOf(_1477_v14)
						var _1479_v16 _dafny.Set
						_ = _1479_v16
						_1479_v16 = _dafny.SetOf((_1478_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1478_v15).Cardinality()))).Uint32()).(_dafny.Sequence))
						_1479_v16 = _1479_v16
						_1449_v0 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1451_v2, _1451_v2)).Cardinality())).Minus(p0)
						_1449_v0 = _1449_v0
						_1449_v0 = (p0).Plus(p0)
						var _1480_v17 _dafny.Array
						_ = _1480_v17
						var _len0_43 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_43
						var _nw198 _dafny.Array
						_ = _nw198
						if _len0_43.Cmp(_dafny.Zero) == 0 {
							_nw198 = _dafny.NewArray(_len0_43)
						} else {
							var _init43 func(_dafny.Int) _dafny.Map = (func(_1481_v1 bool, _1482_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
								return func(_1483_i2 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1481_v1, _dafny.IntOfUint32((_1482_v14).Cardinality()))
								}
							})(_1450_v1, _1477_v14)
							_ = _init43
							var _element0_43 = _init43(_dafny.Zero)
							_ = _element0_43
							_nw198 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
							(_nw198).ArraySet1(_element0_43, 0)
							var _nativeLen0_43 = (_len0_43).Int()
							_ = _nativeLen0_43
							for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
								(_nw198).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
							}
						}
						_1480_v17 = _nw198
						var _1484_v18 *C1
						_ = _1484_v18
						var _nw199 *C1 = New_C1_()
						_ = _nw199
						_nw199.Ctor__(_1480_v17, !(_1450_v1), _this.F2())
						_1484_v18 = _nw199
					}
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1509_v19 _dafny.Array
		_ = _1509_v19
		var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw200
		_1509_v19 = _nw200
		_1509_v19 = _1509_v19
		var _1510_v20 _dafny.Sequence
		_ = _1510_v20
		_1510_v20 = _dafny.UnicodeSeqOfUtf8Bytes("qbtir")
		_1510_v20 = _dafny.Companion_Sequence_.Concatenate(_1510_v20, _dafny.Companion_Sequence_.Concatenate(_1510_v20, _dafny.UnicodeSeqOfUtf8Bytes("nxsjpdcw")))
		_1449_v0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1449_v0))
		var _1511_v21 _dafny.MultiSet
		_ = _1511_v21
		_1511_v21 = _dafny.MultiSetOf(_this.F2())
		r0 = _1511_v21
		r1 = (func() bool {
			if _1450_v1 {
				return true
			}
			return !(false)
		})()
		return r0, r1
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f6 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f6 = false
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

func (_this *C4) Ctor__(f6 bool) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C4) Fm5(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6()))).Cardinality()).Times(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-550), (_this).F6())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("su")).Cardinality()), (Companion_D0_.Create_DC1_(false, (_this).F6(), _dafny.IntOfInt64(972))).Dtor_cf2()))).Cardinality())
	}
}
func (_this *C4) Fm6(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(_dafny.MultiSetOf(!((_this).F6())))).Intersection(_dafny.SetOf(_dafny.MultiSetOf((_this).F6(), (_this).F6())))).IsDisjointFrom((func() _dafny.Set {
			var _coll52 = _dafny.NewBuilder()
			_ = _coll52
			for _iter59 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_this).F6(), (_this).F6(), (_this).F6(), (_this).F6(), false), (_this).F6())).Keys().Elements()); ; {
				_compr_52, _ok59 := _iter59()
				if !_ok59 {
					break
				}
				var _1512_v0 _dafny.MultiSet
				_1512_v0 = interface{}(_compr_52).(_dafny.MultiSet)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_this).F6(), (_this).F6(), (_this).F6(), (_this).F6(), false), (_this).F6())).Contains(_1512_v0) {
					_coll52.Add(_1512_v0)
				}
			}
			return _coll52.ToSet()
		}()).Intersection(_dafny.SetOf(_dafny.MultiSetOf((_this).F6(), (_this).F6()), _dafny.MultiSetOf((_this).F6(), (_this).F6()))))
	}
}
func (_this *C4) M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (bool, _dafny.Array, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _1513_v0 _dafny.Array
		_ = _1513_v0
		var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
		_ = _nw201
		_1513_v0 = _nw201
		_1513_v0 = _1513_v0
		var _1514_v1 _dafny.Array
		_ = _1514_v1
		var _len0_44 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_44
		var _nw202 _dafny.Array
		_ = _nw202
		if _len0_44.Cmp(_dafny.Zero) == 0 {
			_nw202 = _dafny.NewArray(_len0_44)
		} else {
			var _init44 func(_dafny.Int) _dafny.Int = (func(_1515_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1516_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1516_i1, (_dafny.Zero).Minus(_1515_p0))
				}
			})(p0)
			_ = _init44
			var _element0_44 = _init44(_dafny.Zero)
			_ = _element0_44
			_nw202 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
			(_nw202).ArraySet1(_element0_44, 0)
			var _nativeLen0_44 = (_len0_44).Int()
			_ = _nativeLen0_44
			for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
				(_nw202).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
			}
		}
		_1514_v1 = _nw202
		for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1514_v1), 0))); ; {
			_guard_loop_7, _ok60 := _iter60()
			if !_ok60 {
				break
			}
			var _1517_i0 _dafny.Int
			_1517_i0 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_1517_i0).Sign() != -1) && ((_1517_i0).Cmp(_dafny.ArrayLen((_1514_v1), 0)) < 0)) {
				(_1514_v1).ArraySet1((_1517_i0).Times(p0), (_1517_i0).Int())
			}
		}
		var _hi8 _dafny.Int = p2
		_ = _hi8
		for _1518_i2 := (_dafny.Zero).Minus(p2); _1518_i2.Cmp(_hi8) < 0; _1518_i2 = _1518_i2.Plus(_dafny.One) {
			var _1519_v2 _dafny.Int
			_ = _1519_v2
			_1519_v2 = _dafny.IntOfInt64(578)
			var _1520_v3 _dafny.Sequence
			_ = _1520_v3
			_1520_v3 = _dafny.UnicodeSeqOfUtf8Bytes("s")
			var _rhs186 _dafny.Int = p2
			_ = _rhs186
			var _rhs187 _dafny.Int = (func() _dafny.Int {
				if (_this).F6() {
					return _1518_i2
				}
				return (_dafny.Zero).Minus(_1518_i2)
			})()
			_ = _rhs187
			var _rhs188 _dafny.Sequence = _1520_v3
			_ = _rhs188
			var _rhs189 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs189
			var _rhs190 bool = (_this).F6()
			_ = _rhs190
			_1519_v2 = _rhs186
			_1519_v2 = _rhs187
			r2 = _rhs188
			_1519_v2 = _rhs189
			r0 = _rhs190
			var _1521_v4 *C0
			_ = _1521_v4
			var _nw203 *C0 = New_C0_()
			_ = _nw203
			_nw203.Ctor__()
			_1521_v4 = _nw203
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_1514_v1), 0))
			_ = _index243
			(_1514_v1).ArraySet1(p0, (_index243).Int())
			var _1522_v5 D0
			_ = _1522_v5
			_1522_v5 = Companion_D0_.Create_DC1_((_this).F6(), (_this).F6(), _dafny.IntOfInt64(-38))
			var _1523_v6 _dafny.Sequence
			_ = _1523_v6
			_1523_v6 = _dafny.SeqOf(_1522_v5, _1522_v5, Companion_D0_.Create_DC1_((_this).F6(), !(true), _1519_v2), Companion_D0_.Create_DC1_((_this).F6(), (_this).F6(), p0), Companion_D0_.Create_DC1_((_this).F6(), (_this).F6(), _1519_v2))
			var _1524_v7 D1
			_ = _1524_v7
			_1524_v7 = Companion_D1_.Create_DC3_(_dafny.SeqOf(_1522_v5))
			var _1525_v8 _dafny.MultiSet
			_ = _1525_v8
			_1525_v8 = _dafny.MultiSetOf(_1523_v6, (_1524_v7).Dtor_cf8())
			var _1526_v9 _dafny.Map
			_ = _1526_v9
			_1526_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()), (_this).F6())
			var _1527_v10 _dafny.Map
			_ = _1527_v10
			_1527_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1525_v8, (_1526_v9).Update((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(885))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}(func(_1528_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})))).Cardinality(), (_this).F6()))
			var _1529_v11 _dafny.Map
			_ = _1529_v11
			_1529_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()), false)
			var _1530_v12 _dafny.Sequence
			_ = _1530_v12
			_1530_v12 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1529_v11).Cardinality(), (_this).F6()), _1526_v9)
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_1514_v1), 0))
			_ = _index244
			(_1514_v1).ArraySet1(((func() _dafny.Map {
				if (_1527_v10).Contains(_1525_v8) {
					return (_1527_v10).Get(_1525_v8).(_dafny.Map)
				}
				return (_1530_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-8), _dafny.IntOfUint32((_1530_v12).Cardinality()))).Uint32()).(_dafny.Map)
			})()).Cardinality(), (_index244).Int())
			var _1531_v13 _dafny.CodePoint
			_ = _1531_v13
			_1531_v13 = _dafny.CodePoint('n')
			_1531_v13 = _1531_v13
		}
		var _1532_v14 _dafny.Map
		_ = _1532_v14
		_1532_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F6())
		if (p0).Cmp((_1532_v14).Cardinality()) > 0 {
			var _1533_v15 _dafny.CodePoint
			_ = _1533_v15
			_1533_v15 = _dafny.CodePoint('t')
			var _1534_v16 _dafny.Sequence
			_ = _1534_v16
			_1534_v16 = _dafny.SeqOf(p0)
			var _1535_v17 _dafny.Set
			_ = _1535_v17
			_1535_v17 = _dafny.SetOf(p2)
			var _1536_v18 _dafny.Map
			_ = _1536_v18
			_1536_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1535_v17).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))
			var _1537_v19 _dafny.Sequence
			_ = _1537_v19
			_1537_v19 = _dafny.SeqOf((_1534_v16).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1536_v18).Contains((_1534_v16).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1534_v16).Cardinality()))).Uint32()).(_dafny.Int)) {
					return (_1536_v18).Get((_1534_v16).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1534_v16).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Int)
				}
				return p0
			})(), _dafny.IntOfUint32((_1534_v16).Cardinality()))).Uint32()).(_dafny.Int), p2)
			var _1538_v20 _dafny.Sequence
			_ = _1538_v20
			_1538_v20 = _dafny.SeqOf(_1537_v19)
			var _1539_v21 _dafny.Array
			_ = _1539_v21
			var _nwElement0_40 bool = (_this).F6()
			_ = _nwElement0_40
			var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(29))
			_ = _nw204
			(_nw204).ArraySet1(_nwElement0_40, 0)
			(_nw204).ArraySet1(((_this).F6()) || (!((_this).F6())), 1)
			(_nw204).ArraySet1(true, 2)
			(_nw204).ArraySet1((_this).F6(), 3)
			(_nw204).ArraySet1((_this).F6(), 4)
			(_nw204).ArraySet1(!((_this).F6()), 5)
			(_nw204).ArraySet1((p0).Cmp(p0) > 0, 6)
			(_nw204).ArraySet1((_this).F6(), 7)
			(_nw204).ArraySet1((_this).Fm6(_dafny.SeqOf(_1533_v15, _1533_v15, _1533_v15, _1533_v15, _1533_v15), (_this).F6(), globalState), 8)
			(_nw204).ArraySet1((_this).F6(), 9)
			(_nw204).ArraySet1(!((_dafny.IntOfUint32((_1538_v20).Cardinality())).Cmp(p2) >= 0), 10)
			(_nw204).ArraySet1((func() bool {
				if (_this).F6() {
					return (_this).F6()
				}
				return (_this).F6()
			})(), 11)
			(_nw204).ArraySet1((_this).F6(), 12)
			(_nw204).ArraySet1(false, 13)
			(_nw204).ArraySet1((_this).F6(), 14)
			(_nw204).ArraySet1((!(!((_this).F6()))) && ((_this).F6()), 15)
			(_nw204).ArraySet1((_this).F6(), 16)
			(_nw204).ArraySet1((_this).F6(), 17)
			(_nw204).ArraySet1((_this).F6(), 18)
			(_nw204).ArraySet1((_this).F6(), 19)
			(_nw204).ArraySet1((_this).F6(), 20)
			(_nw204).ArraySet1(!((_this).F6()) || ((_this).F6()), 21)
			(_nw204).ArraySet1((_this).F6(), 22)
			(_nw204).ArraySet1((_this).F6(), 23)
			(_nw204).ArraySet1((_this).F6(), 24)
			(_nw204).ArraySet1((_this).F6(), 25)
			(_nw204).ArraySet1(false, 26)
			(_nw204).ArraySet1((_this).F6(), 27)
			(_nw204).ArraySet1(!((_this).F6()) || ((_this).F6()), 28)
			_1539_v21 = _nw204
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1539_v21), 0))
			_ = _index245
			(_1539_v21).ArraySet1((_this).F6(), (_index245).Int())
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1539_v21), 0))
			_ = _index246
			(_1539_v21).ArraySet1((_this).F6(), (_index246).Int())
			var _1540_v22 _dafny.Sequence
			_ = _1540_v22
			_1540_v22 = _dafny.SeqOf(_dafny.CodePoint('r'), _1533_v15)
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1539_v21), 0))
			_ = _index247
			var _rhs191 _dafny.Sequence = _1540_v22
			_ = _rhs191
			var _rhs192 bool = ((func() _dafny.Int {
				if (_this).F6() {
					return p2
				}
				return (_dafny.Zero).Minus(p2)
			})()).Cmp(p0) != 0
			_ = _rhs192
			var _lhs119 _dafny.Array = _1539_v21
			_ = _lhs119
			var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1539_v21), 0))
			_ = _lhs120
			r2 = _rhs191
			(_lhs119).ArraySet1(_rhs192, (_lhs120).Int())
			var _1541_v23 _dafny.Int
			_ = _1541_v23
			_1541_v23 = _dafny.IntOfInt64(284)
			var _1542_v24 _dafny.Set
			_ = _1542_v24
			_1542_v24 = _dafny.SetOf(Companion_Default___.Fm8(_1533_v15, globalState))
			_1541_v23 = ((_1542_v24).Cardinality()).Minus(Companion_Default___.Fm0(_1541_v23, globalState))
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_1514_v1), 0))
			_ = _index248
			(_1514_v1).ArraySet1(_dafny.IntOfUint32((_1540_v22).Cardinality()), (_index248).Int())
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_1514_v1), 0))
			_ = _index249
			(_1514_v1).ArraySet1((_dafny.Zero).Minus((p2).Minus((p2).Minus(p2))), (_index249).Int())
			var _1543_v25 _dafny.MultiSet
			_ = _1543_v25
			_1543_v25 = _dafny.MultiSetOf((_1539_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1539_v21), 0))).Int()).(bool), (_1539_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1539_v21), 0))).Int()).(bool))
			_1541_v23 = ((_1514_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_1514_v1), 0))).Int()).(_dafny.Int)).Minus((p2).Times(((_1543_v25).Update((_1539_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1539_v21), 0))).Int()).(bool), Companion_Default___.Abs(p2))).Cardinality()))
		} else {
			var _1544_v26 _dafny.Sequence
			_ = _1544_v26
			_1544_v26 = _dafny.UnicodeSeqOfUtf8Bytes("jtgco")
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))
			_ = _index250
			(_1513_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1544_v26, _1544_v26), (_index250).Int())
			var _1545_v27 _dafny.Array
			_ = _1545_v27
			var _nw205 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw205
			_1545_v27 = _nw205
			var _1546_v28 _dafny.Int
			_ = _1546_v28
			_1546_v28 = _dafny.IntOfInt64(260)
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))
			_ = _index251
			var _rhs193 bool = false
			_ = _rhs193
			var _rhs194 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1544_v26, _1544_v26)
			_ = _rhs194
			var _rhs195 _dafny.Array = _1545_v27
			_ = _rhs195
			var _rhs196 _dafny.Int = (_dafny.Zero).Minus(p2)
			_ = _rhs196
			var _lhs121 _dafny.Array = _1513_v0
			_ = _lhs121
			var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))
			_ = _lhs122
			r0 = _rhs193
			(_lhs121).ArraySet1(_rhs194, (_lhs122).Int())
			_1545_v27 = _rhs195
			_1546_v28 = _rhs196
			var _1547_v29 D2
			_ = _1547_v29
			_1547_v29 = Companion_D2_.Create_DC5_(_1544_v26)
			var _1548_v30 _dafny.CodePoint
			_ = _1548_v30
			_1548_v30 = _dafny.CodePoint('x')
			if (_this).Fm6((_1547_v29).Dtor_cf13(), _dafny.Companion_Sequence_.Contains(_1544_v26, _1548_v30), globalState) {
				var _1549_v31 D2
				_ = _1549_v31
				_1549_v31 = Companion_D2_.Create_DC6_(_1548_v30, _1546_v28)
				var _1550_v32 D2
				_ = _1550_v32
				_1550_v32 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC5_(_1544_v26))
				var _1551_v33 _dafny.Map
				_ = _1551_v33
				_1551_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1513_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))).Int()).(_dafny.Sequence))
				var _1552_v34 _dafny.Sequence
				_ = _1552_v34
				_1552_v34 = _dafny.SeqOf(_1550_v32, _1550_v32)
				_1549_v31 = Companion_Default___.Fm9(p0, ((_this).F6()) && ((_this).F6()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(_1550_v32), _dafny.SeqOf(Companion_D2_.Create_DC8_(Companion_D2_.Create_DC6_(_1548_v30, (_1551_v33).Cardinality()))), _1552_v34)).Cardinality()), globalState)
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))
				_ = _index252
				(_1513_v0).ArraySet1((_1513_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))).Int()).(_dafny.Sequence), (_index252).Int())
				var _1553_v35 _dafny.Array
				_ = _1553_v35
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
				_ = _nw206
				_1553_v35 = _nw206
				_1553_v35 = _1553_v35
				_1546_v28 = (p0).Plus(p2)
				_1546_v28 = p2
			} else {
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_1514_v1), 0))
				_ = _index253
				(_1514_v1).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(854), _1546_v28), (_index253).Int())
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_1514_v1), 0))
				_ = _index254
				(_1514_v1).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(604))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}((func(_1554_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1555_i4 _dafny.Int) _dafny.CodePoint {
						return _1554_v30
					}
				})(_1548_v30)))).Cardinality())).Plus((_dafny.Zero).Minus(_1546_v28)), (_index254).Int())
				r0 = (_this).F6()
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_1514_v1), 0))
				_ = _index255
				(_1514_v1).ArraySet1((_dafny.IntOfInt64(390)).Plus((_1514_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_1514_v1), 0))).Int()).(_dafny.Int)), (_index255).Int())
				r0 = (p2).Cmp(_1546_v28) != 0
				var _1556_v36 _dafny.Array
				_ = _1556_v36
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
				_ = _nw207
				_1556_v36 = _nw207
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1556_v36), 0))
				_ = _index256
				(_1556_v36).ArraySet1((p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1544_v26).Cardinality()), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), (_index256).Int())
				var _1557_v37 _dafny.Map
				_ = _1557_v37
				_1557_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, ((_this).F6()) || (!((_this).F6())))
				var _1558_v38 D2
				_ = _1558_v38
				_1558_v38 = Companion_D2_.Create_DC6_(_1548_v30, _1546_v28)
				var _1559_v39 D2
				_ = _1559_v39
				_1559_v39 = Companion_D2_.Create_DC8_(_1558_v38)
				var _pat_let_tv102 = _1558_v38
				_ = _pat_let_tv102
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1556_v36), 0))
				_ = _index257
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_1514_v1), 0))
				_ = _index258
				var _rhs197 bool = ((p2).Plus(p2)).Cmp((_dafny.Zero).Minus(p2)) <= 0
				_ = _rhs197
				var _rhs198 bool = (func() bool {
					if (_1532_v14).Contains(_dafny.IntOfInt64(915)) {
						return (_1532_v14).Get(_dafny.IntOfInt64(915)).(bool)
					}
					return (_this).F6()
				})()
				_ = _rhs198
				var _rhs199 _dafny.Int = (p0).Minus((_dafny.IntOfUint32((_dafny.SeqOf(func(_pat_let39_0 D2) D2 {
					return func(_1560_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let40_0 D2) D2 {
							return func(_1561_dt__update_hcf21_h0 D2) D2 {
								return Companion_D2_.Create_DC8_(_1561_dt__update_hcf21_h0)
							}(_pat_let40_0)
						}(_pat_let_tv102)
					}(_pat_let39_0)
				}(_1559_v39), Companion_D2_.Create_DC8_(_1558_v38), _1559_v39, _1559_v39, _1559_v39)).Cardinality())).Minus(p2))
				_ = _rhs199
				var _rhs200 _dafny.Map = (_1557_v37).Update(_dafny.Companion_Sequence_.Concatenate(p1, p1), ((_this).F6()) || ((_this).F6()))
				_ = _rhs200
				var _lhs123 _dafny.Array = _1556_v36
				_ = _lhs123
				var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1556_v36), 0))
				_ = _lhs124
				var _lhs125 _dafny.Array = _1514_v1
				_ = _lhs125
				var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_1514_v1), 0))
				_ = _lhs126
				r0 = _rhs197
				(_lhs123).ArraySet1(_rhs198, (_lhs124).Int())
				(_lhs125).ArraySet1(_rhs199, (_lhs126).Int())
				_1557_v37 = _rhs200
			}
			_1548_v30 = Companion_Default___.Fm8(_1548_v30, globalState)
			if true {
				var _1562_v40 _dafny.Array
				_ = _1562_v40
				var _nw208 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw208
				_1562_v40 = _nw208
				(globalState).F1 = (func() _dafny.Array {
					if false {
						return _1562_v40
					}
					return _1562_v40
				})()
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1562_v40), 0))
				_ = _index259
				(_1562_v40).ArraySet1((_this).F6(), (_index259).Int())
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1562_v40), 0))
				_ = _index260
				(_1562_v40).ArraySet1((_this).F6(), (_index260).Int())
				var _1563_v41 _dafny.Map
				_ = _1563_v41
				_1563_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1546_v28, _1546_v28)
				_1563_v41 = (_1563_v41).Update((_dafny.IntOfInt64(-499)).Times(_dafny.IntOfUint32(((_1513_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))).Int()).(_dafny.Sequence)).Cardinality())), _1546_v28)
				var _1564_v42 _dafny.MultiSet
				_ = _1564_v42
				_1564_v42 = _dafny.MultiSetOf((_this).F6(), (_this).F6())
				var _1565_v43 D0
				_ = _1565_v43
				_1565_v43 = Companion_D0_.Create_DC2_(p2, (_this).F6(), (_this).F6(), (_1564_v42).Cardinality())
				var _1566_v44 _dafny.Map
				_ = _1566_v44
				_1566_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
				var _1567_v45 _dafny.MultiSet
				_ = _1567_v45
				_1567_v45 = _dafny.MultiSetOf(_dafny.IntOfInt64(-423), _1546_v28, p2)
				var _1568_v46 _dafny.Set
				_ = _1568_v46
				_1568_v46 = _dafny.SetOf(Companion_Default___.Fm10((_1562_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1562_v40), 0))).Int()).(bool), (_dafny.Zero).Minus(_1546_v28), _1565_v43, globalState), (_dafny.MultiSetOf((_1566_v44).Cardinality())).Intersection(_1567_v45), (_1567_v45).Difference(_1567_v45))
				_1568_v46 = (_dafny.SetOf(_1567_v45)).Union(_1568_v46)
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))
				_ = _index261
				(_1513_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1513_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate((_1513_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm11(p0, (_1562_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((_1562_v40), 0))).Int()).(bool), _1546_v28, globalState))), (_index261).Int())
			} else {
				_1544_v26 = _1544_v26
				var _1569_v47 _dafny.Map
				_ = _1569_v47
				_1569_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1513_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))).Int()).(_dafny.Sequence), _1532_v14)
				var _1570_v48 _dafny.Map
				_ = _1570_v48
				_1570_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1569_v47)
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))
				_ = _index262
				(_1514_v1).ArraySet1((((func() _dafny.Map {
					if (_1570_v48).Contains(false) {
						return (_1570_v48).Get(false).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1513_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))).Int()).(_dafny.Sequence), _1532_v14)
				})()).Merge(func() _dafny.Map {
					var _coll53 = _dafny.NewMapBuilder()
					_ = _coll53
					for _iter61 := _dafny.Iterate((_dafny.SetOf(_1544_v26)).Elements()); ; {
						_compr_53, _ok61 := _iter61()
						if !_ok61 {
							break
						}
						var _1571_v49 _dafny.Sequence
						_1571_v49 = interface{}(_compr_53).(_dafny.Sequence)
						if (_dafny.SetOf(_1544_v26)).Contains(_1571_v49) {
							_coll53.Add(_1571_v49, _1532_v14)
						}
					}
					return _coll53.ToMap()
				}())).Cardinality(), (_index262).Int())
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))
				_ = _index263
				(_1514_v1).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.IntOfInt64(308), p2, p2), (Companion_Default___.SafeIndex(_1546_v28, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(308), p2, p2)).Cardinality()))).Uint32(), p0)).Cardinality()), _1546_v28), (_index263).Int())
				var _1572_v50 _dafny.Set
				_ = _1572_v50
				_1572_v50 = _dafny.SetOf((_1513_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_1513_v0), 0))).Int()).(_dafny.Sequence), _1544_v26)
				var _1573_v51 _dafny.Array
				_ = _1573_v51
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw209
				_1573_v51 = _nw209
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_1573_v51), 0))
				_ = _index264
				(_1573_v51).ArraySet1(!((_this).F6()) || ((_this).F6()), (_index264).Int())
				var _1574_v52 _dafny.Set
				_ = _1574_v52
				_1574_v52 = _dafny.SetOf((_this).F6(), (_this).F6())
				var _1575_v53 _dafny.Set
				_ = _1575_v53
				_1575_v53 = _dafny.SetOf(_1574_v52, (_1574_v52).Intersection(_1574_v52))
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))
				_ = _index265
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_1573_v51), 0))
				_ = _index266
				var _rhs201 _dafny.Int = (_dafny.Zero).Minus((_1575_v53).Cardinality())
				_ = _rhs201
				var _rhs202 _dafny.Set = _1572_v50
				_ = _rhs202
				var _rhs203 bool = (_this).F6()
				_ = _rhs203
				var _rhs204 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1544_v26, _1544_v26)
				_ = _rhs204
				var _lhs127 _dafny.Array = _1514_v1
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))
				_ = _lhs128
				var _lhs129 _dafny.Array = _1573_v51
				_ = _lhs129
				var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_1573_v51), 0))
				_ = _lhs130
				(_lhs127).ArraySet1(_rhs201, (_lhs128).Int())
				_1572_v50 = _rhs202
				(_lhs129).ArraySet1(_rhs203, (_lhs130).Int())
				_1544_v26 = _rhs204
				var _1576_v54 _dafny.Sequence
				_ = _1576_v54
				_1576_v54 = _dafny.SeqOf((_1514_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))).Int()).(_dafny.Int))
				var _1577_v55 _dafny.Map
				_ = _1577_v55
				_1577_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfInt64(848))
				var _1578_v56 _dafny.Set
				_ = _1578_v56
				_1578_v56 = _dafny.SetOf(_1577_v55, _1577_v55)
				var _1579_v57 _dafny.MultiSet
				_ = _1579_v57
				_1579_v57 = _dafny.MultiSetOf((func() bool {
					if (_this).F6() {
						return (_this).F6()
					}
					return !((_1573_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_1573_v51), 0))).Int()).(bool))
				})())
				var _1580_v59 _dafny.Map
				_ = _1580_v59
				_1580_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll54 = _dafny.NewMapBuilder()
					_ = _coll54
					for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(393), _dafny.IntOfInt64(951))); ; {
						_compr_54, _ok62 := _iter62()
						if !_ok62 {
							break
						}
						var _1581_v58 _dafny.Int
						_1581_v58 = interface{}(_compr_54).(_dafny.Int)
						if ((_dafny.IntOfInt64(393)).Cmp(_1581_v58) <= 0) && ((_1581_v58).Cmp(_dafny.IntOfInt64(951)) < 0) {
							_coll54.Add((_1581_v58).Minus((_1514_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))).Int()).(_dafny.Int)), (_1514_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))).Int()).(_dafny.Int))
						}
					}
					return _coll54.ToMap()
				}(), (_1573_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_1573_v51), 0))).Int()).(bool))
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_1573_v51), 0))
				_ = _index267
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))
				_ = _index268
				var _rhs205 _dafny.Sequence = _1576_v54
				_ = _rhs205
				var _rhs206 _dafny.Set = (_1578_v56).Difference(func() _dafny.Set {
					var _coll55 = _dafny.NewBuilder()
					_ = _coll55
					for _iter63 := _dafny.Iterate((_1580_v59).Keys().Elements()); ; {
						_compr_55, _ok63 := _iter63()
						if !_ok63 {
							break
						}
						var _1582_v60 _dafny.Map
						_1582_v60 = interface{}(_compr_55).(_dafny.Map)
						if (_1580_v59).Contains(_1582_v60) {
							_coll55.Add(_1582_v60)
						}
					}
					return _coll55.ToSet()
				}())
				_ = _rhs206
				var _rhs207 _dafny.MultiSet = (_1579_v57).Update((_this).F6(), Companion_Default___.Abs(p0))
				_ = _rhs207
				var _rhs208 bool = (_1573_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_1573_v51), 0))).Int()).(bool)
				_ = _rhs208
				var _rhs209 _dafny.Int = p0
				_ = _rhs209
				var _lhs131 _dafny.Array = _1573_v51
				_ = _lhs131
				var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_1573_v51), 0))
				_ = _lhs132
				var _lhs133 _dafny.Array = _1514_v1
				_ = _lhs133
				var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1514_v1), 0))
				_ = _lhs134
				_1576_v54 = _rhs205
				_1578_v56 = _rhs206
				_1579_v57 = _rhs207
				(_lhs131).ArraySet1(_rhs208, (_lhs132).Int())
				(_lhs133).ArraySet1(_rhs209, (_lhs134).Int())
				var _1583_v61 *C0
				_ = _1583_v61
				var _nw210 *C0 = New_C0_()
				_ = _nw210
				_nw210.Ctor__()
				_1583_v61 = _nw210
			}
			r0 = (_this).F6()
		}
		var _1584_v62 _dafny.CodePoint
		_ = _1584_v62
		_1584_v62 = _dafny.CodePoint('c')
		var _1585_v63 _dafny.Sequence
		_ = _1585_v63
		_1585_v63 = _dafny.SeqOf(_1584_v62, _1584_v62, _1584_v62, _1584_v62, _1584_v62)
		var _1586_v64 D2
		_ = _1586_v64
		_1586_v64 = Companion_D2_.Create_DC5_(_1585_v63)
		var _1587_v65 _dafny.Sequence
		_ = _1587_v65
		_1587_v65 = _dafny.SeqOf(_1585_v63, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(925))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg71 _dafny.Int) interface{} {
				return coer71(arg71)
			}
		}((func(_1588_v62 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1589_i5 _dafny.Int) _dafny.CodePoint {
				return _1588_v62
			}
		})(_1584_v62))), _1585_v63)
		var _1590_v66 _dafny.Array
		_ = _1590_v66
		var _nwElement0_41 D2 = _1586_v64
		_ = _nwElement0_41
		var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(10))
		_ = _nw211
		(_nw211).ArraySet1(_nwElement0_41, 0)
		(_nw211).ArraySet1(_1586_v64, 1)
		(_nw211).ArraySet1(_1586_v64, 2)
		(_nw211).ArraySet1(_1586_v64, 3)
		(_nw211).ArraySet1(_1586_v64, 4)
		(_nw211).ArraySet1(_1586_v64, 5)
		(_nw211).ArraySet1(_1586_v64, 6)
		(_nw211).ArraySet1(_1586_v64, 7)
		(_nw211).ArraySet1(Companion_D2_.Create_DC5_((_1587_v65).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1587_v65).Cardinality()))).Uint32()).(_dafny.Sequence)), 8)
		(_nw211).ArraySet1(_1586_v64, 9)
		_1590_v66 = _nw211
		var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1590_v66), 0))
		_ = _index269
		(_1590_v66).ArraySet1(_1586_v64, (_index269).Int())
		var _1591_v67 _dafny.Int
		_ = _1591_v67
		_1591_v67 = _dafny.IntOfInt64(702)
		var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1590_v66), 0))
		_ = _index270
		var _rhs210 D2 = _1586_v64
		_ = _rhs210
		var _rhs211 _dafny.Int = p2
		_ = _rhs211
		var _lhs135 _dafny.Array = _1590_v66
		_ = _lhs135
		var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1590_v66), 0))
		_ = _lhs136
		(_lhs135).ArraySet1(_rhs210, (_lhs136).Int())
		_1591_v67 = _rhs211
		var _1592_v68 _dafny.MultiSet
		_ = _1592_v68
		_1592_v68 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.Zero)).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg72 _dafny.Int) interface{} {
				return coer72(arg72)
			}
		}((func(_1593_v62 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1594_i6 _dafny.Int) _dafny.CodePoint {
				return _1593_v62
			}
		})(_1584_v62))))
		r0 = (_1592_v68).IsSubsetOf(_1592_v68)
		r0 = (_this).F6()
		var _1595_v69 _dafny.Array
		_ = _1595_v69
		var _len0_45 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_45
		var _nw212 _dafny.Array
		_ = _nw212
		if _len0_45.Cmp(_dafny.Zero) == 0 {
			_nw212 = _dafny.NewArray(_len0_45)
		} else {
			var _init45 func(_dafny.Int) bool = func(_1596_i7 _dafny.Int) bool {
				return (_this).F6()
			}
			_ = _init45
			var _element0_45 = _init45(_dafny.Zero)
			_ = _element0_45
			_nw212 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
			(_nw212).ArraySet1(_element0_45, 0)
			var _nativeLen0_45 = (_len0_45).Int()
			_ = _nativeLen0_45
			for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
				(_nw212).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
			}
		}
		_1595_v69 = _nw212
		r1 = _1595_v69
		r2 = (func() _dafny.Sequence {
			if (_this).F6() {
				return _dafny.Companion_Sequence_.Concatenate(_1585_v63, _dafny.UnicodeSeqOfUtf8Bytes("flcs"))
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("tudymb")
		})()
		return r0, r1, r2
	}
}
func (_this *C4) M3(p0 _dafny.Array, p1 bool, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1597_v0 _dafny.Sequence
		_ = _1597_v0
		_1597_v0 = _dafny.SeqOf((_this).F6())
		var _1598_v1 _dafny.Set
		_ = _1598_v1
		_1598_v1 = _dafny.SetOf(_1597_v0, _1597_v0, _dafny.Companion_Sequence_.Concatenate(_1597_v0, _1597_v0), _1597_v0)
		_1598_v1 = _1598_v1
		var _1599_v2 _dafny.Int
		_ = _1599_v2
		_1599_v2 = _dafny.IntOfInt64(482)
		var _hi9 _dafny.Int = _1599_v2
		_ = _hi9
		for _1600_i0 := Companion_Default___.Fm0(_1599_v2, globalState); _1600_i0.Cmp(_hi9) < 0; _1600_i0 = _1600_i0.Plus(_dafny.One) {
			var _1601_v3 _dafny.Array
			_ = _1601_v3
			var _len0_46 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_46
			var _nw213 _dafny.Array
			_ = _nw213
			if _len0_46.Cmp(_dafny.Zero) == 0 {
				_nw213 = _dafny.NewArray(_len0_46)
			} else {
				var _init46 func(_dafny.Int) bool = (func(_1602_p1 bool) func(_dafny.Int) bool {
					return func(_1603_i1 _dafny.Int) bool {
						return !(_1602_p1)
					}
				})(p1)
				_ = _init46
				var _element0_46 = _init46(_dafny.Zero)
				_ = _element0_46
				_nw213 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
				(_nw213).ArraySet1(_element0_46, 0)
				var _nativeLen0_46 = (_len0_46).Int()
				_ = _nativeLen0_46
				for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
					(_nw213).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
				}
			}
			_1601_v3 = _nw213
			(globalState).F1 = _1601_v3
			var _1604_v4 _dafny.Sequence
			_ = _1604_v4
			_1604_v4 = _dafny.UnicodeSeqOfUtf8Bytes("bhimfjb")
			if _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1604_v4, _dafny.UnicodeSeqOfUtf8Bytes("uu")), _1604_v4) {
				var _1605_v5 _dafny.CodePoint
				_ = _1605_v5
				_1605_v5 = _dafny.CodePoint('k')
				var _1606_v6 _dafny.Map
				_ = _1606_v6
				_1606_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1597_v0)
				var _1607_v7 _dafny.Map
				_ = _1607_v7
				_1607_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1605_v5, (_this).F6())
				var _1608_v8 D3
				_ = _1608_v8
				_1608_v8 = Companion_D3_.Create_DC9_(_1607_v7)
				var _1609_v9 _dafny.Map
				_ = _1609_v9
				_1609_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1608_v8).Dtor_cf22(), (_this).F6())
				var _1610_v10 D2
				_ = _1610_v10
				_1610_v10 = Companion_D2_.Create_DC6_(_1605_v5, (_dafny.Zero).Minus((_1609_v9).Cardinality()))
				var _1611_v11 _dafny.Map
				_ = _1611_v11
				_1611_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1599_v2, (func() D2 {
					if p1 {
						return Companion_D2_.Create_DC6_(_1605_v5, (_1606_v6).Cardinality())
					}
					return _1610_v10
				})())
				_1611_v11 = (_1611_v11).Update((func() _dafny.Int {
					if !(p1) {
						return _1600_i0
					}
					return _1599_v2
				})(), _1610_v10)
				var _1612_v12 _dafny.MultiSet
				_ = _1612_v12
				_1612_v12 = _dafny.MultiSetOf(true)
				var _1613_v13 D0
				_ = _1613_v13
				_1613_v13 = Companion_D0_.Create_DC0_((_this).F6())
				var _1614_v14 _dafny.Map
				_ = _1614_v14
				_1614_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1613_v13, _1599_v2)
				var _1615_v15 _dafny.MultiSet
				_ = _1615_v15
				_1615_v15 = _dafny.MultiSetOf((_1612_v12).Cardinality(), _dafny.IntOfUint32((_1604_v4).Cardinality()), _1600_i0, (_1614_v14).Cardinality(), _dafny.IntOfInt64(956))
				r1 = (_1600_i0).Minus((func() _dafny.Int {
					if (_1615_v15).Contains(_1599_v2) {
						return (_1615_v15).Multiplicity(_1599_v2)
					}
					return (_dafny.Zero).Minus(_1600_i0)
				})())
				var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_1601_v3), 0))
				_ = _index271
				(_1601_v3).ArraySet1(p1, (_index271).Int())
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_1601_v3), 0))
				_ = _index272
				(_1601_v3).ArraySet1(((_1599_v2).Times(_1599_v2)).Cmp(_1600_i0) <= 0, (_index272).Int())
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((p0), 0))
				_ = _index273
				(p0).ArraySet1(_dafny.IntOfInt64(-736), (_index273).Int())
				var _1616_v16 _dafny.Sequence
				_ = _1616_v16
				_1616_v16 = _dafny.SeqOf(_1599_v2, _1600_i0)
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((p0), 0))
				_ = _index274
				(p0).ArraySet1((_1616_v16).Select((Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_1616_v16).Cardinality()))).Uint32()).(_dafny.Int), (_index274).Int())
				r0 = p1
			} else {
				var _1617_v17 _dafny.MultiSet
				_ = _1617_v17
				_1617_v17 = _dafny.MultiSetOf((_this).F6())
				_1617_v17 = ((_1617_v17).Update(!((_this).F6()), Companion_Default___.Abs(_1600_i0))).Intersection((_1617_v17).Update(p1, Companion_Default___.Abs(_dafny.IntOfInt64(491))))
				r1 = (_dafny.Zero).Minus(_1599_v2)
				_1599_v2 = (_1599_v2).Minus(_1600_i0)
				_1599_v2 = ((_1617_v17).Cardinality()).Minus(_1600_i0)
				var _1618_v18 _dafny.CodePoint
				_ = _1618_v18
				_1618_v18 = _dafny.CodePoint('e')
				var _1619_v19 _dafny.Sequence
				_ = _1619_v19
				_1619_v19 = _dafny.SeqOf((_1600_i0).Minus(_1600_i0), (_dafny.Zero).Minus(_1599_v2), (_dafny.Zero).Minus(_1600_i0))
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_1601_v3), 0))
				_ = _index275
				(_1601_v3).ArraySet1(p1, (_index275).Int())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_1601_v3), 0))
				_ = _index276
				var _rhs212 _dafny.CodePoint = _1618_v18
				_ = _rhs212
				var _rhs213 _dafny.Sequence = _1619_v19
				_ = _rhs213
				var _rhs214 bool = p1
				_ = _rhs214
				var _rhs215 bool = (func() bool {
					if p1 {
						return (_this).F6()
					}
					return (_this).F6()
				})()
				_ = _rhs215
				var _lhs137 _dafny.Array = _1601_v3
				_ = _lhs137
				var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_1601_v3), 0))
				_ = _lhs138
				_1618_v18 = _rhs212
				_1619_v19 = _rhs213
				r0 = _rhs214
				(_lhs137).ArraySet1(_rhs215, (_lhs138).Int())
			}
			_1604_v4 = _dafny.UnicodeSeqOfUtf8Bytes("uca")
			var _1620_v20 *C0
			_ = _1620_v20
			var _nw214 *C0 = New_C0_()
			_ = _nw214
			_nw214.Ctor__()
			_1620_v20 = _nw214
		}
		var _1621_v21 _dafny.MultiSet
		_ = _1621_v21
		_1621_v21 = _dafny.MultiSetOf(_1599_v2, _dafny.IntOfInt64(19), _1599_v2, _1599_v2, _1599_v2)
		var _1622_v22 _dafny.Sequence
		_ = _1622_v22
		_1622_v22 = _dafny.SeqOf(_1621_v21, _1621_v21)
		if !((_1622_v22).Select((Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_1622_v22).Cardinality()))).Uint32()).(_dafny.MultiSet)).Equals(_1621_v21) {
			_1599_v2 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_1599_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ppwwf")).Cardinality())), _1599_v2)
			var _1623_v23 _dafny.Sequence
			_ = _1623_v23
			_1623_v23 = _dafny.SeqOf(_1599_v2, _dafny.IntOfInt64(792))
			var _1624_v24 _dafny.Map
			_ = _1624_v24
			_1624_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1623_v23)).Cardinality(), (_this).F6())
			_1624_v24 = (_1624_v24).Update(_1599_v2, !(p1))
			_1599_v2 = _1599_v2
			var _1625_v25 _dafny.Array
			_ = _1625_v25
			var _nw215 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw215
			_1625_v25 = _nw215
			var _1626_v26 _dafny.Sequence
			_ = _1626_v26
			_1626_v26 = _dafny.SeqOf(_1625_v25)
			var _1627_v27 _dafny.Map
			_ = _1627_v27
			_1627_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1599_v2, (_1626_v26).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1623_v23).Cardinality()), _dafny.IntOfUint32((_1626_v26).Cardinality()))).Uint32()).(_dafny.Array))
			var _rhs216 _dafny.Array = (func() _dafny.Array {
				if (_1627_v27).Contains(_1599_v2) {
					return (_1627_v27).Get(_1599_v2).(_dafny.Array)
				}
				return _1625_v25
			})()
			_ = _rhs216
			var _rhs217 _dafny.Int = (_1599_v2).Plus(_dafny.IntOfUint32((_1623_v23).Cardinality()))
			_ = _rhs217
			var _rhs218 bool = (_this).F6()
			_ = _rhs218
			var _rhs219 _dafny.Array = _1625_v25
			_ = _rhs219
			var _lhs139 *GlobalState = globalState
			_ = _lhs139
			var _lhs140 *GlobalState = globalState
			_ = _lhs140
			_lhs139.F1 = _rhs216
			r1 = _rhs217
			r0 = _rhs218
			_lhs140.F1 = _rhs219
			var _1628_v28 _dafny.Array
			_ = _1628_v28
			var _len0_47 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_47
			var _nw216 _dafny.Array
			_ = _nw216
			if _len0_47.Cmp(_dafny.Zero) == 0 {
				_nw216 = _dafny.NewArray(_len0_47)
			} else {
				var _init47 func(_dafny.Int) _dafny.Map = (func(_1629_v2 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_1630_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1629_v2)
					}
				})(_1599_v2)
				_ = _init47
				var _element0_47 = _init47(_dafny.Zero)
				_ = _element0_47
				_nw216 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
				(_nw216).ArraySet1(_element0_47, 0)
				var _nativeLen0_47 = (_len0_47).Int()
				_ = _nativeLen0_47
				for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
					(_nw216).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
				}
			}
			_1628_v28 = _nw216
			var _1631_v29 T0
			_ = _1631_v29
			var _nw217 *C1 = New_C1_()
			_ = _nw217
			_nw217.Ctor__(_1628_v28, (_this).F6(), _1625_v25)
			_1631_v29 = _nw217
			_1631_v29 = _1631_v29
		} else {
			var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
			_ = _index277
			(p0).ArraySet1(_1599_v2, (_index277).Int())
			var _1632_v30 _dafny.Map
			_ = _1632_v30
			_1632_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1597_v0, _1597_v0)).Cardinality()), _1599_v2)
			var _1633_v31 _dafny.Set
			_ = _1633_v31
			_1633_v31 = _dafny.SetOf(p0, p0)
			var _1634_v32 D5
			_ = _1634_v32
			_1634_v32 = Companion_D5_.Create_DC17_((_1633_v31).Cardinality())
			var _pat_let_tv103 = _1599_v2
			_ = _pat_let_tv103
			var _pat_let_tv104 = _1599_v2
			_ = _pat_let_tv104
			var _1635_v33 _dafny.Array
			_ = _1635_v33
			var _nwElement0_42 D5 = _1634_v32
			_ = _nwElement0_42
			var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(13))
			_ = _nw218
			(_nw218).ArraySet1(_nwElement0_42, 0)
			(_nw218).ArraySet1(_1634_v32, 1)
			(_nw218).ArraySet1(_1634_v32, 2)
			(_nw218).ArraySet1(_1634_v32, 3)
			(_nw218).ArraySet1(Companion_D5_.Create_DC17_(_1599_v2), 4)
			(_nw218).ArraySet1(Companion_D5_.Create_DC17_((_dafny.Zero).Minus(_dafny.IntOfUint32((_1597_v0).Cardinality()))), 5)
			(_nw218).ArraySet1(_1634_v32, 6)
			(_nw218).ArraySet1(_1634_v32, 7)
			(_nw218).ArraySet1(Companion_Default___.Fm27(_1599_v2, _1599_v2, globalState), 8)
			(_nw218).ArraySet1(Companion_Default___.Fm27(_1599_v2, _1599_v2, globalState), 9)
			(_nw218).ArraySet1(_1634_v32, 10)
			(_nw218).ArraySet1(func(_pat_let41_0 D5) D5 {
				return func(_1636_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let42_0 _dafny.Int) D5 {
						return func(_1637_dt__update_hcf35_h0 _dafny.Int) D5 {
							return Companion_D5_.Create_DC17_(_1637_dt__update_hcf35_h0)
						}(_pat_let42_0)
					}(_pat_let_tv103)
				}(_pat_let41_0)
			}(_1634_v32), 11)
			(_nw218).ArraySet1(func(_pat_let43_0 D5) D5 {
				return func(_1638_dt__update__tmp_h1 D5) D5 {
					return func(_pat_let44_0 _dafny.Int) D5 {
						return func(_1639_dt__update_hcf35_h1 _dafny.Int) D5 {
							return Companion_D5_.Create_DC17_(_1639_dt__update_hcf35_h1)
						}(_pat_let44_0)
					}(_pat_let_tv104)
				}(_pat_let43_0)
			}(_1634_v32), 12)
			_1635_v33 = _nw218
			var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1635_v33), 0))
			_ = _index278
			(_1635_v33).ArraySet1(_1634_v32, (_index278).Int())
			var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
			_ = _index279
			var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1635_v33), 0))
			_ = _index280
			var _rhs220 _dafny.Int = _1599_v2
			_ = _rhs220
			var _rhs221 _dafny.Map = _1632_v30
			_ = _rhs221
			var _rhs222 D5 = _1634_v32
			_ = _rhs222
			var _lhs141 _dafny.Array = p0
			_ = _lhs141
			var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
			_ = _lhs142
			var _lhs143 _dafny.Array = _1635_v33
			_ = _lhs143
			var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_1635_v33), 0))
			_ = _lhs144
			(_lhs141).ArraySet1(_rhs220, (_lhs142).Int())
			_1632_v30 = _rhs221
			(_lhs143).ArraySet1(_rhs222, (_lhs144).Int())
			if (_this).F6() {
				var _1640_v34 _dafny.Array
				_ = _1640_v34
				var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
				_ = _nw219
				_1640_v34 = _nw219
				_1640_v34 = _1640_v34
				var _1641_v35 _dafny.CodePoint
				_ = _1641_v35
				_1641_v35 = _dafny.CodePoint('k')
				_1641_v35 = _1641_v35
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _index281
				(p0).ArraySet1(_1599_v2, (_index281).Int())
				var _1642_v36 _dafny.Sequence
				_ = _1642_v36
				_1642_v36 = _dafny.SeqOf(_1599_v2)
				var _1643_v37 D0
				_ = _1643_v37
				_1643_v37 = Companion_D0_.Create_DC0_(p1)
				var _1644_v38 _dafny.Array
				_ = _1644_v38
				var _nwElement0_43 bool = !(!(p1))
				_ = _nwElement0_43
				var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(15))
				_ = _nw220
				(_nw220).ArraySet1(_nwElement0_43, 0)
				(_nw220).ArraySet1(!((_this).F6()), 1)
				(_nw220).ArraySet1((_this).F6(), 2)
				(_nw220).ArraySet1((_this).F6(), 3)
				(_nw220).ArraySet1(p1, 4)
				(_nw220).ArraySet1((_this).F6(), 5)
				(_nw220).ArraySet1(!(p1) || ((_this).F6()), 6)
				(_nw220).ArraySet1(((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_1642_v36).Cardinality())) > 0, 7)
				(_nw220).ArraySet1((_this).F6(), 8)
				(_nw220).ArraySet1((_1643_v37).Dtor_cf0(), 9)
				(_nw220).ArraySet1(p1, 10)
				(_nw220).ArraySet1(!((_this).F6()), 11)
				(_nw220).ArraySet1(!(p1) || ((_this).F6()), 12)
				(_nw220).ArraySet1((_1621_v21).IsDisjointFrom(_1621_v21), 13)
				(_nw220).ArraySet1((_this).F6(), 14)
				_1644_v38 = _nw220
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_1644_v38), 0))
				_ = _index282
				(_1644_v38).ArraySet1((_this).F6(), (_index282).Int())
				var _1645_v39 _dafny.Array
				_ = _1645_v39
				var _nw221 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
				_ = _nw221
				_1645_v39 = _nw221
				var _1646_v40 D10
				_ = _1646_v40
				_1646_v40 = Companion_D10_.Create_DC32_(_dafny.MultiSetOf(_1644_v38), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
				var _1647_v41 _dafny.MultiSet
				_ = _1647_v41
				_1647_v41 = _dafny.MultiSetOf(_1644_v38, _1644_v38)
				var _pat_let_tv105 = _1647_v41
				_ = _pat_let_tv105
				var _1648_v42 _dafny.MultiSet
				_ = _1648_v42
				_1648_v42 = _dafny.MultiSetOf(func(_pat_let45_0 D10) D10 {
					return func(_1649_dt__update__tmp_h2 D10) D10 {
						return func(_pat_let46_0 _dafny.MultiSet) D10 {
							return func(_1650_dt__update_hcf58_h0 _dafny.MultiSet) D10 {
								return Companion_D10_.Create_DC32_(_1650_dt__update_hcf58_h0, (_1649_dt__update__tmp_h2).Dtor_cf59(), (_1649_dt__update__tmp_h2).Dtor_cf60())
							}(_pat_let46_0)
						}(_pat_let_tv105)
					}(_pat_let45_0)
				}(_1646_v40))
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_1645_v39), 0))
				_ = _index283
				(_1645_v39).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1597_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1648_v42).Cardinality()), _dafny.IntOfUint32((_1597_v0).Cardinality()))).Uint32(), p1), (Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1597_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1648_v42).Cardinality()), _dafny.IntOfUint32((_1597_v0).Cardinality()))).Uint32(), p1)).Cardinality()))).Uint32(), p1), (_index283).Int())
				var _1651_v43 _dafny.Sequence
				_ = _1651_v43
				_1651_v43 = _dafny.UnicodeSeqOfUtf8Bytes("hpkbj")
				var _1652_v44 _dafny.Array
				_ = _1652_v44
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_48
				var _nw222 _dafny.Array
				_ = _nw222
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw222 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) _dafny.Int = (func(_1653_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1654_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1654_i3, _1653_v2)
						}
					})(_1599_v2)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw222 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw222).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw222).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1652_v44 = _nw222
				var _1655_v45 _dafny.Map
				_ = _1655_v45
				_1655_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1652_v44)
				var _1656_v46 _dafny.Array
				_ = _1656_v46
				var _nwElement0_44 _dafny.Array = _1652_v44
				_ = _nwElement0_44
				var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(18))
				_ = _nw223
				(_nw223).ArraySet1(_nwElement0_44, 0)
				(_nw223).ArraySet1(_1652_v44, 1)
				(_nw223).ArraySet1(_1652_v44, 2)
				(_nw223).ArraySet1(_1652_v44, 3)
				(_nw223).ArraySet1(_1652_v44, 4)
				(_nw223).ArraySet1(_1652_v44, 5)
				(_nw223).ArraySet1(_1652_v44, 6)
				(_nw223).ArraySet1(_1652_v44, 7)
				(_nw223).ArraySet1((func() _dafny.Array {
					if false {
						return _1652_v44
					}
					return _1652_v44
				})(), 8)
				(_nw223).ArraySet1(_1652_v44, 9)
				(_nw223).ArraySet1((func() _dafny.Array {
					if (_1655_v45).Contains(!(p1)) {
						return (_1655_v45).Get(!(p1)).(_dafny.Array)
					}
					return _1652_v44
				})(), 10)
				(_nw223).ArraySet1(_1652_v44, 11)
				(_nw223).ArraySet1((func() _dafny.Array {
					if p1 {
						return _1652_v44
					}
					return _1652_v44
				})(), 12)
				(_nw223).ArraySet1(_1652_v44, 13)
				(_nw223).ArraySet1(_1652_v44, 14)
				(_nw223).ArraySet1(_1652_v44, 15)
				(_nw223).ArraySet1(_1652_v44, 16)
				(_nw223).ArraySet1(_1652_v44, 17)
				_1656_v46 = _nw223
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1656_v46), 0))
				_ = _index284
				(_1656_v46).ArraySet1(_1652_v44, (_index284).Int())
				var _1657_v47 _dafny.Array
				_ = _1657_v47
				_1657_v47 = _1652_v44
				var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_1644_v38), 0))
				_ = _index285
				var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _index286
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_1645_v39), 0))
				_ = _index287
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1656_v46), 0))
				_ = _index288
				var _rhs223 bool = (func() bool {
					if _dafny.Companion_Sequence_.IsProperPrefixOf(_1642_v36, _1642_v36) {
						return false
					}
					return (_this).F6()
				})()
				_ = _rhs223
				var _rhs224 _dafny.Int = (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)
				_ = _rhs224
				var _rhs225 _dafny.Sequence = _1597_v0
				_ = _rhs225
				var _rhs226 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("ghblvvi")
				_ = _rhs226
				var _rhs227 _dafny.Array = (_1657_v47)
				_ = _rhs227
				var _lhs145 _dafny.Array = _1644_v38
				_ = _lhs145
				var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_1644_v38), 0))
				_ = _lhs146
				var _lhs147 _dafny.Array = p0
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _lhs148
				var _lhs149 _dafny.Array = _1645_v39
				_ = _lhs149
				var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_1645_v39), 0))
				_ = _lhs150
				var _lhs151 _dafny.Array = _1656_v46
				_ = _lhs151
				var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1656_v46), 0))
				_ = _lhs152
				(_lhs145).ArraySet1(_rhs223, (_lhs146).Int())
				(_lhs147).ArraySet1(_rhs224, (_lhs148).Int())
				(_lhs149).ArraySet1(_rhs225, (_lhs150).Int())
				_1651_v43 = _rhs226
				(_lhs151).ArraySet1(_rhs227, (_lhs152).Int())
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_1644_v38), 0))
				_ = _index289
				(_1644_v38).ArraySet1((_1597_v0).Select((Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1597_v0).Cardinality()))).Uint32()).(bool), (_index289).Int())
			} else {
				r2 = p1
				var _1658_v48 _dafny.Sequence
				_ = _1658_v48
				_1658_v48 = _dafny.UnicodeSeqOfUtf8Bytes("ask")
				var _1659_v49 _dafny.Map
				_ = _1659_v49
				_1659_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1658_v48, _1599_v2)
				_1659_v49 = _1659_v49
				var _1660_v50 _dafny.Array
				_ = _1660_v50
				var _nw224 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw224
				_1660_v50 = _nw224
				var _1661_v51 *C3
				_ = _1661_v51
				var _nw225 *C3 = New_C3_()
				_ = _nw225
				_nw225.Ctor__(_1660_v50)
				_1661_v51 = _nw225
				var _1662_v52 _dafny.CodePoint
				_ = _1662_v52
				_1662_v52 = _dafny.CodePoint('h')
				_1662_v52 = _1662_v52
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _index290
				(p0).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_this).F6() {
						return _1658_v48
					}
					return _1658_v48
				})()).Cardinality()), (_index290).Int())
			}
			var _1663_v53 _dafny.Array
			_ = _1663_v53
			var _len0_49 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_49
			var _nw226 _dafny.Array
			_ = _nw226
			if _len0_49.Cmp(_dafny.Zero) == 0 {
				_nw226 = _dafny.NewArray(_len0_49)
			} else {
				var _init49 func(_dafny.Int) bool = func(_1664_i4 _dafny.Int) bool {
					return (_this).F6()
				}
				_ = _init49
				var _element0_49 = _init49(_dafny.Zero)
				_ = _element0_49
				_nw226 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
				(_nw226).ArraySet1(_element0_49, 0)
				var _nativeLen0_49 = (_len0_49).Int()
				_ = _nativeLen0_49
				for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
					(_nw226).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
				}
			}
			_1663_v53 = _nw226
			var _1665_v54 *C3
			_ = _1665_v54
			var _nw227 *C3 = New_C3_()
			_ = _nw227
			_nw227.Ctor__(_1663_v53)
			_1665_v54 = _nw227
			var _1666_v55 _dafny.Array
			_ = _1666_v55
			var _len0_50 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_50
			var _nw228 _dafny.Array
			_ = _nw228
			if _len0_50.Cmp(_dafny.Zero) == 0 {
				_nw228 = _dafny.NewArray(_len0_50)
			} else {
				var _init50 func(_dafny.Int) _dafny.Map = (func(_1667_p1 bool, _1668_p0 _dafny.Array) func(_dafny.Int) _dafny.Map {
					return func(_1669_i5 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1667_p1, (_1668_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1668_p0), 0))).Int()).(_dafny.Int))
					}
				})(p1, p0)
				_ = _init50
				var _element0_50 = _init50(_dafny.Zero)
				_ = _element0_50
				_nw228 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
				(_nw228).ArraySet1(_element0_50, 0)
				var _nativeLen0_50 = (_len0_50).Int()
				_ = _nativeLen0_50
				for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
					(_nw228).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
				}
			}
			_1666_v55 = _nw228
			var _1670_v56 *C1
			_ = _1670_v56
			var _nw229 *C1 = New_C1_()
			_ = _nw229
			_nw229.Ctor__(_1666_v55, p1, _1663_v53)
			_1670_v56 = _nw229
			var _1671_v57 D10
			_ = _1671_v57
			_1671_v57 = Companion_D10_.Create_DC31_(_1670_v56)
			var _1672_v58 _dafny.Sequence
			_ = _1672_v58
			_1672_v58 = _dafny.SeqOf((_1671_v57).Dtor_cf57())
			var _1673_v59 _dafny.CodePoint
			_ = _1673_v59
			_1673_v59 = _dafny.CodePoint('w')
			var _1674_v60 _dafny.Set
			_ = _1674_v60
			_1674_v60 = _dafny.SetOf(p1)
			var _rhs228 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1672_v58, (Companion_Default___.SafeIndex((_dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("yilayj"), (Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yilayj")).Cardinality()))).Uint32(), _1673_v59))).Cardinality(), _dafny.IntOfUint32((_1672_v58).Cardinality()))).Uint32(), _1670_v56), _1672_v58)).Cardinality())
			_ = _rhs228
			var _rhs229 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1599_v2))
			_ = _rhs229
			var _rhs230 _dafny.Int = (func() _dafny.Int {
				if p1 {
					return ((_1674_v60).Union(_dafny.SetOf(p1, p1, (_this).F6(), false))).Cardinality()
				}
				return ((_1621_v21).Difference(_dafny.MultiSetOf(_1599_v2, _1599_v2, _1599_v2, _dafny.IntOfInt64(-508)))).Cardinality()
			})()
			_ = _rhs230
			var _rhs231 *C3 = _1665_v54
			_ = _rhs231
			_1599_v2 = _rhs228
			r1 = _rhs229
			_1599_v2 = _rhs230
			_1665_v54 = _rhs231
			var _1675_v62 _dafny.MultiSet
			_ = _1675_v62
			_1675_v62 = _dafny.MultiSetOf(func() _dafny.Set {
				var _coll56 = _dafny.NewBuilder()
				_ = _coll56
				for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-757), _dafny.IntOfInt64(981))); ; {
					_compr_56, _ok64 := _iter64()
					if !_ok64 {
						break
					}
					var _1676_v61 _dafny.Int
					_1676_v61 = interface{}(_compr_56).(_dafny.Int)
					if ((_dafny.IntOfInt64(-757)).Cmp(_1676_v61) <= 0) && ((_1676_v61).Cmp(_dafny.IntOfInt64(981)) < 0) {
						_coll56.Add((_1676_v61).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(79))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg73 _dafny.Int) interface{} {
								return coer73(arg73)
							}
						}((func(_1677_v59 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1678_i6 _dafny.Int) _dafny.CodePoint {
								return _1677_v59
							}
						})(_1673_v59)))).Cardinality()))))
					}
				}
				return _coll56.ToSet()
			}())
			_1675_v62 = _1675_v62
			var _1679_v63 _dafny.MultiSet
			_ = _1679_v63
			_1679_v63 = _dafny.MultiSetOf(_1673_v59, _1673_v59)
			var _1680_v64 D4
			_ = _1680_v64
			_1680_v64 = Companion_D4_.Create_DC13_(((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Cmp(_1599_v2) < 0, _1679_v63, (p1) && ((_this).F6()))
			var _source23 D4 = _1680_v64
			_ = _source23
			if _source23.Is_DC12() {
				var _1681_v65 _dafny.Array
				_ = _1681_v65
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_51
				var _nw230 _dafny.Array
				_ = _nw230
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw230 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) _dafny.Int = (func(_1682_p0 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_1683_i7 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1683_i7, (_1682_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1682_p0), 0))).Int()).(_dafny.Int))
						}
					})(p0)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw230 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw230).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw230).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_1681_v65 = _nw230
				var _1684_v66 _dafny.Map
				_ = _1684_v66
				_1684_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1681_v65)
				var _1685_v67 _dafny.Map
				_ = _1685_v67
				_1685_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1684_v66).Cardinality()), p1)
				var _1686_v68 _dafny.Sequence
				_ = _1686_v68
				_1686_v68 = _dafny.SeqOf(_1685_v67)
				_1686_v68 = _1686_v68
				var _1687_v69 _dafny.Map
				_ = _1687_v69
				_1687_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
				_1687_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !(false))
				(globalState).F1 = _1663_v53
				var _1688_v70 _dafny.Map
				_ = _1688_v70
				_1688_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(true)), (_this).F6())
				_1688_v70 = _1688_v70
			} else if _source23.Is_DC13() {
				var _1689___mcc_h0 bool = _source23.Get_().(D4_DC13).Cf29
				_ = _1689___mcc_h0
				var _1690___mcc_h1 _dafny.MultiSet = _source23.Get_().(D4_DC13).Cf30
				_ = _1690___mcc_h1
				var _1691___mcc_h2 bool = _source23.Get_().(D4_DC13).Cf31
				_ = _1691___mcc_h2
				var _1692_cf31 bool = _1691___mcc_h2
				_ = _1692_cf31
				var _1693_cf30 _dafny.MultiSet = _1690___mcc_h1
				_ = _1693_cf30
				var _1694_cf29 bool = _1689___mcc_h0
				_ = _1694_cf29
				var _1695_v71 *C2
				_ = _1695_v71
				var _nw231 *C2 = New_C2_()
				_ = _nw231
				_nw231.Ctor__(_1692_cf31, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _1666_v55, (_1599_v2).Cmp(_1599_v2) != 0, _1663_v53)
				_1695_v71 = _nw231
				var _1696_v72 _dafny.Sequence
				_ = _1696_v72
				_1696_v72 = _dafny.SeqOf(_dafny.CodePoint('m'), _1673_v59, _dafny.CodePoint('u'))
				var _1697_v73 _dafny.Set
				_ = _1697_v73
				_1697_v73 = _dafny.SetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _1599_v2, _1599_v2, (_1695_v71).F8())
				var _1698_v74 D12
				_ = _1698_v74
				_1698_v74 = Companion_D12_.Create_DC37_(_1695_v71.F7, _1597_v0, (_1697_v73).Cardinality())
				var _1699_v75 _dafny.Map
				_ = _1699_v75
				_1699_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_1698_v74).Dtor_cf69(), globalState), _1673_v59)
				var _1700_v76 _dafny.Array
				_ = _1700_v76
				var _nwElement0_45 _dafny.CodePoint = _1673_v59
				_ = _nwElement0_45
				var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(16))
				_ = _nw232
				(_nw232).ArraySet1CodePoint(_nwElement0_45, 0)
				(_nw232).ArraySet1CodePoint((_1696_v72).Select((Companion_Default___.SafeIndex((_1695_v71).F8(), _dafny.IntOfUint32((_1696_v72).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
				(_nw232).ArraySet1CodePoint(_1673_v59, 2)
				(_nw232).ArraySet1CodePoint(Companion_Default___.Fm22(_1696_v72, globalState), 3)
				(_nw232).ArraySet1CodePoint(_1673_v59, 4)
				(_nw232).ArraySet1CodePoint(_1673_v59, 5)
				(_nw232).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_1699_v75).Contains((_1695_v71).F8()) {
						return (_1699_v75).Get((_1695_v71).F8()).(_dafny.CodePoint)
					}
					return _1673_v59
				})(), 6)
				(_nw232).ArraySet1CodePoint(_1673_v59, 7)
				(_nw232).ArraySet1CodePoint(_1673_v59, 8)
				(_nw232).ArraySet1CodePoint(_1673_v59, 9)
				(_nw232).ArraySet1CodePoint(_1673_v59, 10)
				(_nw232).ArraySet1CodePoint(_1673_v59, 11)
				(_nw232).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _1694_cf29 {
						return _1673_v59
					}
					return _1673_v59
				})(), 12)
				(_nw232).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_1699_v75).Contains(_dafny.IntOfInt64(299)) {
						return (_1699_v75).Get(_dafny.IntOfInt64(299)).(_dafny.CodePoint)
					}
					return _1673_v59
				})(), 13)
				(_nw232).ArraySet1CodePoint(_1673_v59, 14)
				(_nw232).ArraySet1CodePoint(_1673_v59, 15)
				_1700_v76 = _nw232
				var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1700_v76), 0))
				_ = _index291
				(_1700_v76).ArraySet1CodePoint(_1673_v59, (_index291).Int())
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1700_v76), 0))
				_ = _index292
				var _rhs232 *C2 = _1695_v71
				_ = _rhs232
				var _rhs233 _dafny.CodePoint = _1673_v59
				_ = _rhs233
				var _lhs153 _dafny.Array = _1700_v76
				_ = _lhs153
				var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1700_v76), 0))
				_ = _lhs154
				_1695_v71 = _rhs232
				(_lhs153).ArraySet1CodePoint(_rhs233, (_lhs154).Int())
				var _1701_v77 _dafny.Array
				_ = _1701_v77
				var _nw233 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
				_ = _nw233
				_1701_v77 = _nw233
				var _1702_v78 D6
				_ = _1702_v78
				_1702_v78 = Companion_D6_.Create_DC19_(_1701_v77)
				var _1703_v79 D6
				_ = _1703_v79
				_1703_v79 = Companion_D6_.Create_DC21_(_1702_v78)
				var _1704_v80 _dafny.Array
				_ = _1704_v80
				var _nwElement0_46 D6 = _1703_v79
				_ = _nwElement0_46
				var _nw234 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(5))
				_ = _nw234
				(_nw234).ArraySet1(_nwElement0_46, 0)
				(_nw234).ArraySet1(_1703_v79, 1)
				(_nw234).ArraySet1(_1703_v79, 2)
				(_nw234).ArraySet1(_1703_v79, 3)
				(_nw234).ArraySet1(_1703_v79, 4)
				_1704_v80 = _nw234
				var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_1704_v80), 0))
				_ = _index293
				(_1704_v80).ArraySet1(Companion_D6_.Create_DC21_(_1702_v78), (_index293).Int())
				var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_1704_v80), 0))
				_ = _index294
				(_1704_v80).ArraySet1(_1703_v79, (_index294).Int())
				var _1705_v81 _dafny.Array
				_ = _1705_v81
				var _nw235 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw235
				_1705_v81 = _nw235
				_1705_v81 = _1705_v81
				var _1706_v82 _dafny.MultiSet
				_ = _1706_v82
				_1706_v82 = _dafny.MultiSetOf(_1663_v53)
				var _1707_v83 _dafny.Sequence
				_ = _1707_v83
				_1707_v83 = _dafny.SeqOf(Companion_D7_.Create_DC22_(_1706_v82))
				var _1708_v84 _dafny.Map
				_ = _1708_v84
				_1708_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1670_v56, Companion_Default___.Fm0(_1599_v2, globalState))
				var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _index295
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1700_v76), 0))
				_ = _index296
				var _rhs234 _dafny.Int = (_1695_v71).F8()
				_ = _rhs234
				var _rhs235 _dafny.Int = (_1708_v84).Cardinality()
				_ = _rhs235
				var _rhs236 bool = p1
				_ = _rhs236
				var _rhs237 _dafny.CodePoint = (_1700_v76).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1700_v76), 0))).Int())
				_ = _rhs237
				var _rhs238 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1707_v83, _1707_v83)
				_ = _rhs238
				var _lhs155 _dafny.Array = p0
				_ = _lhs155
				var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _lhs156
				var _lhs157 _dafny.Array = _1700_v76
				_ = _lhs157
				var _lhs158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_1700_v76), 0))
				_ = _lhs158
				(_lhs155).ArraySet1(_rhs234, (_lhs156).Int())
				_1599_v2 = _rhs235
				r2 = _rhs236
				(_lhs157).ArraySet1CodePoint(_rhs237, (_lhs158).Int())
				_1707_v83 = _rhs238
			} else {
				var _1709___mcc_h3 _dafny.Sequence = _source23.Get_().(D4_DC11).Cf28
				_ = _1709___mcc_h3
				var _1710_cf28 _dafny.Sequence = _1709___mcc_h3
				_ = _1710_cf28
				var _1711_v85 *C0
				_ = _1711_v85
				var _nw236 *C0 = New_C0_()
				_ = _nw236
				_nw236.Ctor__()
				_1711_v85 = _nw236
				var _1712_v86 _dafny.Array
				_ = _1712_v86
				var _nwElement0_47 *C0 = _1711_v85
				_ = _nwElement0_47
				var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(5))
				_ = _nw237
				(_nw237).ArraySet1(_nwElement0_47, 0)
				(_nw237).ArraySet1(_1711_v85, 1)
				(_nw237).ArraySet1(_1711_v85, 2)
				(_nw237).ArraySet1(_1711_v85, 3)
				(_nw237).ArraySet1(_1711_v85, 4)
				_1712_v86 = _nw237
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_1712_v86), 0))
				_ = _index297
				(_1712_v86).ArraySet1(_1711_v85, (_index297).Int())
				var _1713_v87 _dafny.Sequence
				_ = _1713_v87
				_1713_v87 = _dafny.UnicodeSeqOfUtf8Bytes("jkney")
				var _1714_v88 _dafny.Map
				_ = _1714_v88
				_1714_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1713_v87, (_this).F6())
				var _1715_v89 _dafny.Map
				_ = _1715_v89
				_1715_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1714_v88).Contains(_1713_v87) {
						return (_1714_v88).Get(_1713_v87).(bool)
					}
					return p1
				})(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
				var _1716_v90 _dafny.Map
				_ = _1716_v90
				_1716_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F6())
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _index298
				var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _index299
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_1712_v86), 0))
				_ = _index300
				var _rhs239 _dafny.Int = _1599_v2
				_ = _rhs239
				var _rhs240 _dafny.Int = _1599_v2
				_ = _rhs240
				var _rhs241 *C0 = _1711_v85
				_ = _rhs241
				var _rhs242 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(184))).Merge(_1715_v89)
				_ = _rhs242
				var _rhs243 _dafny.Int = (((func() _dafny.Map {
					if p1 {
						return _1716_v90
					}
					return _1716_v90
				})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1))).Cardinality()
				_ = _rhs243
				var _lhs159 _dafny.Array = p0
				_ = _lhs159
				var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _lhs160
				var _lhs161 _dafny.Array = p0
				_ = _lhs161
				var _lhs162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))
				_ = _lhs162
				var _lhs163 _dafny.Array = _1712_v86
				_ = _lhs163
				var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_1712_v86), 0))
				_ = _lhs164
				(_lhs159).ArraySet1(_rhs239, (_lhs160).Int())
				(_lhs161).ArraySet1(_rhs240, (_lhs162).Int())
				(_lhs163).ArraySet1(_rhs241, (_lhs164).Int())
				_1715_v89 = _rhs242
				_1599_v2 = _rhs243
				var _1717_v91 _dafny.MultiSet
				_ = _1717_v91
				_1717_v91 = _dafny.MultiSetOf((_this).F6(), p1, p1)
				var _1718_v92 _dafny.Map
				_ = _1718_v92
				_1718_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1673_v59, _1717_v91)
				var _1719_v93 D6
				_ = _1719_v93
				_1719_v93 = Companion_D6_.Create_DC20_(p1, (_this).F6())
				var _1720_v94 _dafny.Map
				_ = _1720_v94
				_1720_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm5((func() _dafny.MultiSet {
					if (_1718_v92).Contains(_1673_v59) {
						return (_1718_v92).Get(_1673_v59).(_dafny.MultiSet)
					}
					return _1717_v91
				})(), globalState), _1719_v93)
				_1720_v94 = (_1720_v94).Update((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _1719_v93)
				var _1721_v95 _dafny.Array
				_ = _1721_v95
				var _nw238 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(23))
				_ = _nw238
				_1721_v95 = _nw238
				var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1721_v95), 0))
				_ = _index301
				(_1721_v95).ArraySet1CodePoint(_1673_v59, (_index301).Int())
				var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1721_v95), 0))
				_ = _index302
				(_1721_v95).ArraySet1CodePoint(_1673_v59, (_index302).Int())
				var _1722_v96 _dafny.Map
				_ = _1722_v96
				_1722_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1713_v87).Cardinality()), (func() bool {
					if p1 {
						return !((_this).F6())
					}
					return (_this).F6()
				})())
				_1722_v96 = (_1722_v96).Update((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_this).F6())
			}
		}
		var _1723_v97 _dafny.Array
		_ = _1723_v97
		var _nw239 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
		_ = _nw239
		_1723_v97 = _nw239
		_1723_v97 = p0
		if (_this).F6() {
			r0 = (_this).F6()
			var _1724_v98 _dafny.Sequence
			_ = _1724_v98
			_1724_v98 = _dafny.UnicodeSeqOfUtf8Bytes("upfnqt")
			var _1725_v99 _dafny.Map
			_ = _1725_v99
			_1725_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm11(_1599_v2, (_this).F6(), _dafny.IntOfUint32((_1724_v98).Cardinality()), globalState), _1599_v2)
			var _1726_v100 D8
			_ = _1726_v100
			_1726_v100 = Companion_D8_.Create_DC25_((_1725_v99).Merge(_1725_v99))
			var _1727_v101 D13
			_ = _1727_v101
			_1727_v101 = Companion_D13_.Create_DC41_((_dafny.Zero).Minus(_1599_v2))
			var _1728_v102 _dafny.Map
			_ = _1728_v102
			_1728_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1599_v2)
			_1726_v100 = Companion_Default___.Fm35(_1727_v101, !(_1728_v102).Contains((_this).F6()), globalState)
			r2 = ((func() _dafny.Int {
				if (_1728_v102).Contains(p1) {
					return (_1728_v102).Get(p1).(_dafny.Int)
				}
				return _1599_v2
			})()).Cmp(_dafny.IntOfUint32((_1724_v98).Cardinality())) >= 0
			var _rhs244 _dafny.Int = _dafny.IntOfInt64(368)
			_ = _rhs244
			var _rhs245 _dafny.Sequence = _1597_v0
			_ = _rhs245
			r1 = _rhs244
			_1597_v0 = _rhs245
			r0 = (_this).F6()
		} else {
			var _1729_v103 _dafny.Map
			_ = _1729_v103
			_1729_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), p0)
			_1729_v103 = (_1729_v103).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1723_v97))
			var _1730_v104 _dafny.Sequence
			_ = _1730_v104
			_1730_v104 = _dafny.SeqOf(_dafny.IntOfUint32((_1597_v0).Cardinality()))
			var _1731_v105 _dafny.Array
			_ = _1731_v105
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_52
			var _nw240 _dafny.Array
			_ = _nw240
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw240 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) _dafny.Map = (func(_1732_p1 bool, _1733_v2 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_1734_i8 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1732_p1, _1733_v2)
					}
				})(p1, _1599_v2)
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw240 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw240).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw240).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_1731_v105 = _nw240
			var _1735_v106 _dafny.Array
			_ = _1735_v106
			var _nwElement0_48 bool = p1
			_ = _nwElement0_48
			var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(27))
			_ = _nw241
			(_nw241).ArraySet1(_nwElement0_48, 0)
			(_nw241).ArraySet1((_this).F6(), 1)
			(_nw241).ArraySet1(p1, 2)
			(_nw241).ArraySet1((_this).F6(), 3)
			(_nw241).ArraySet1(p1, 4)
			(_nw241).ArraySet1(p1, 5)
			(_nw241).ArraySet1((_this).F6(), 6)
			(_nw241).ArraySet1((_this).F6(), 7)
			(_nw241).ArraySet1((_this).F6(), 8)
			(_nw241).ArraySet1((_this).F6(), 9)
			(_nw241).ArraySet1((_this).F6(), 10)
			(_nw241).ArraySet1((_this).F6(), 11)
			(_nw241).ArraySet1(p1, 12)
			(_nw241).ArraySet1(p1, 13)
			(_nw241).ArraySet1(p1, 14)
			(_nw241).ArraySet1((_this).F6(), 15)
			(_nw241).ArraySet1(Companion_Default___.Fm20(p1, globalState), 16)
			(_nw241).ArraySet1(p1, 17)
			(_nw241).ArraySet1((_this).F6(), 18)
			(_nw241).ArraySet1((_this).F6(), 19)
			(_nw241).ArraySet1((_this).F6(), 20)
			(_nw241).ArraySet1(p1, 21)
			(_nw241).ArraySet1(p1, 22)
			(_nw241).ArraySet1((_this).F6(), 23)
			(_nw241).ArraySet1((_this).F6(), 24)
			(_nw241).ArraySet1((_this).F6(), 25)
			(_nw241).ArraySet1(p1, 26)
			_1735_v106 = _nw241
			var _1736_v107 *C2
			_ = _1736_v107
			var _nw242 *C2 = New_C2_()
			_ = _nw242
			_nw242.Ctor__(p1, _1599_v2, _1731_v105, (_this).F6(), _1735_v106)
			_1736_v107 = _nw242
			var _1737_v108 _dafny.Map
			_ = _1737_v108
			_1737_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if p1 {
					return _1599_v2
				}
				return (_1730_v104).Select((Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_1730_v104).Cardinality()))).Uint32()).(_dafny.Int)
			})(), _1736_v107)
			_1737_v108 = (_1737_v108).Update(_1599_v2, _1736_v107)
			var _1738_v109 bool
			_ = _1738_v109
			var _1739_v110 _dafny.Array
			_ = _1739_v110
			var _1740_v111 _dafny.Sequence
			_ = _1740_v111
			var _out34 bool
			_ = _out34
			var _out35 _dafny.Array
			_ = _out35
			var _out36 _dafny.Sequence
			_ = _out36
			_out34, _out35, _out36 = (_this).M2(_1599_v2, _dafny.SeqOf((_this).F6(), p1, p1), _dafny.IntOfUint32((_1730_v104).Cardinality()), globalState)
			_1738_v109 = _out34
			_1739_v110 = _out35
			_1740_v111 = _out36
			var _1741_v112 _dafny.Array
			_ = _1741_v112
			var _len0_53 _dafny.Int = _dafny.One
			_ = _len0_53
			var _nw243 _dafny.Array
			_ = _nw243
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw243 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) _dafny.Sequence = (func(_1742_v111 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1743_i9 _dafny.Int) _dafny.Sequence {
						return _1742_v111
					}
				})(_1740_v111)
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw243 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw243).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw243).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_1741_v112 = _nw243
			var _1744_v113 _dafny.Array
			_ = _1744_v113
			var _nwElement0_49 _dafny.Map = Companion_Default___.Fm25(p1, globalState)
			_ = _nwElement0_49
			var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.One)
			_ = _nw244
			(_nw244).ArraySet1(_nwElement0_49, 0)
			_1744_v113 = _nw244
			var _1745_v114 _dafny.MultiSet
			_ = _1745_v114
			var _1746_v115 bool
			_ = _1746_v115
			var _out37 _dafny.MultiSet
			_ = _out37
			var _out38 bool
			_ = _out38
			_out37, _out38 = (_1736_v107).M1(_dafny.IntOfInt64(515), Companion_Default___.Fm22(_1740_v111, globalState), _1741_v112, _1744_v113, globalState)
			_1745_v114 = _out37
			_1746_v115 = _out38
			if (_this).F6() {
				r0 = !(p1)
				var _1747_v116 _dafny.Array
				_ = _1747_v116
				var _len0_54 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_54
				var _nw245 _dafny.Array
				_ = _nw245
				if _len0_54.Cmp(_dafny.Zero) == 0 {
					_nw245 = _dafny.NewArray(_len0_54)
				} else {
					var _init54 func(_dafny.Int) _dafny.Map = (func(_1748_v115 bool, _1749_p1 bool) func(_dafny.Int) _dafny.Map {
						return func(_1750_i10 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1748_v115, _1749_p1)
						}
					})(_1746_v115, p1)
					_ = _init54
					var _element0_54 = _init54(_dafny.Zero)
					_ = _element0_54
					_nw245 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
					(_nw245).ArraySet1(_element0_54, 0)
					var _nativeLen0_54 = (_len0_54).Int()
					_ = _nativeLen0_54
					for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
						(_nw245).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
					}
				}
				_1747_v116 = _nw245
				var _1751_v117 D6
				_ = _1751_v117
				_1751_v117 = Companion_D6_.Create_DC19_(_1747_v116)
				var _1752_v118 _dafny.CodePoint
				_ = _1752_v118
				_1752_v118 = _dafny.CodePoint('y')
				var _1753_v119 _dafny.Map
				_ = _1753_v119
				_1753_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1752_v118, (_this).F6())
				var _1754_v120 _dafny.Map
				_ = _1754_v120
				_1754_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1738_v109, _1739_v110)
				var _1755_v121 _dafny.Sequence
				_ = _1755_v121
				_1755_v121 = _dafny.SeqOf((func() _dafny.Array {
					if (_1754_v120).Contains(_1736_v107.F7) {
						return (_1754_v120).Get(_1736_v107.F7).(_dafny.Array)
					}
					return _1739_v110
				})())
				var _rhs246 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1740_v111, _1740_v111)
				_ = _rhs246
				var _rhs247 D6 = _1751_v117
				_ = _rhs247
				var _rhs248 _dafny.Int = (func() _dafny.Int {
					if _1736_v107.F7 {
						return (_dafny.Zero).Minus((_1736_v107).F8())
					}
					return _dafny.IntOfUint32((_1740_v111).Cardinality())
				})()
				_ = _rhs248
				var _rhs249 _dafny.Int = (_1599_v2).Plus(((_1753_v119).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1752_v118, p1))).Cardinality())
				_ = _rhs249
				var _rhs250 _dafny.Array = (_1755_v121).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1599_v2, _1746_v115)).Cardinality(), (_dafny.Zero).Minus((_1621_v21).Cardinality())), _dafny.IntOfUint32((_1755_v121).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs250
				var _lhs165 *GlobalState = globalState
				_ = _lhs165
				_1740_v111 = _rhs246
				_1751_v117 = _rhs247
				_1599_v2 = _rhs248
				r1 = _rhs249
				_lhs165.F1 = _rhs250
				var _1756_v122 _dafny.Map
				_ = _1756_v122
				_1756_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1730_v104, (_1736_v107).F8())
				_1756_v122 = (_1756_v122).Update(_1730_v104, ((_1736_v107).F8()).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-420))))
				var _1757_v123 _dafny.Set
				_ = _1757_v123
				_1757_v123 = _dafny.SetOf((_this).F6())
				var _1758_v124 _dafny.Map
				_ = _1758_v124
				_1758_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1736_v107.F7)
				var _1759_v125 _dafny.Map
				_ = _1759_v125
				_1759_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1758_v124, (_1736_v107).F8())
				var _rhs251 bool = ((_this).F6()) && (p1)
				_ = _rhs251
				var _rhs252 bool = ((_1621_v21).Union(_dafny.MultiSetFromSeq(_1730_v104))).IsSubsetOf(_1621_v21)
				_ = _rhs252
				var _rhs253 bool = !(((_dafny.IntOfUint32((_1740_v111).Cardinality())).Times((_dafny.Zero).Minus((_dafny.Zero).Minus((_1757_v123).Cardinality())))).Cmp((_1759_v125).Cardinality()) == 0)
				_ = _rhs253
				var _rhs254 bool = _1746_v115
				_ = _rhs254
				var _lhs166 *C2 = _1736_v107
				_ = _lhs166
				r2 = _rhs251
				_1738_v109 = _rhs252
				_lhs166.F7 = _rhs253
				r0 = _rhs254
				(_1736_v107).F7 = (_1736_v107.F7) && (_1746_v115)
			} else {
				var _1760_v126 _dafny.Map
				_ = _1760_v126
				_1760_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D8_.Create_DC27_(_1736_v107.F7, (_this).F6(), _1599_v2, (_1736_v107).F8())).Dtor_cf51(), (_dafny.Zero).Minus((_1736_v107).F8()))
				_1760_v126 = (_1760_v126).Update((_1736_v107).F8(), _1599_v2)
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1735_v106), 0))
				_ = _index303
				(_1735_v106).ArraySet1(((_1736_v107).F8()).Cmp(_1599_v2) != 0, (_index303).Int())
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1735_v106), 0))
				_ = _index304
				(_1735_v106).ArraySet1(!(_1736_v107.F7), (_index304).Int())
				r1 = (func() _dafny.Int {
					if (_this).F6() {
						return (_1736_v107).F8()
					}
					return _1599_v2
				})()
				_1740_v111 = _dafny.Companion_Sequence_.Concatenate(_1740_v111, _dafny.UnicodeSeqOfUtf8Bytes("ngbnk"))
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1735_v106), 0))
				_ = _index305
				(_1735_v106).ArraySet1(!(!(!((func() bool {
					if false {
						return !(p1) || (_1746_v115)
					}
					return (func() bool {
						if _1736_v107.F7 {
							return p1
						}
						return p1
					})()
				})()))), (_index305).Int())
			}
		}
		if (_this).F6() {
			if p1 {
				r0 = p1
				_1599_v2 = _1599_v2
				r0 = Companion_Default___.Fm20(p1, globalState)
				var _1761_v127 _dafny.Sequence
				_ = _1761_v127
				_1761_v127 = _dafny.SeqOf(_dafny.IntOfInt64(-302))
				var _1762_v128 _dafny.Sequence
				_ = _1762_v128
				_1762_v128 = _dafny.SeqOf(_1761_v127)
				var _1763_v129 _dafny.Sequence
				_ = _1763_v129
				_1763_v129 = _dafny.SeqOf((_1762_v128).Select((Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_1762_v128).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _1764_v130 _dafny.Sequence
				_ = _1764_v130
				_1764_v130 = _dafny.SeqOf(_1761_v127, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-658))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}((func(_1765_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1766_i11 _dafny.Int) _dafny.Int {
						return _1765_v2
					}
				})(_1599_v2))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(246))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1767_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1768_i12 _dafny.Int) _dafny.Int {
						return _1767_v2
					}
				})(_1599_v2))), _1761_v127, (_1762_v128).Select((Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_1762_v128).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _1769_v131 _dafny.CodePoint
				_ = _1769_v131
				_1769_v131 = _dafny.CodePoint('a')
				var _1770_v132 _dafny.Map
				_ = _1770_v132
				_1770_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_1769_v131)).Cardinality(), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-785))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1771_v127 _dafny.Sequence, _1772_v2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1773_i13 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_dafny.IntOfUint32((_1771_v127).Cardinality()), _1772_v2, _dafny.IntOfInt64(151))
					}
				})(_1761_v127, _1599_v2))), (Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-785))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}((func(_1774_v127 _dafny.Sequence, _1775_v2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1776_i13 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_dafny.IntOfUint32((_1774_v127).Cardinality()), _1775_v2, _dafny.IntOfInt64(151))
					}
				})(_1761_v127, _1599_v2)))).Cardinality()))).Uint32(), _1761_v127), (Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-785))).Uint32(), func(coer78 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}((func(_1777_v127 _dafny.Sequence, _1778_v2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1779_i13 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_dafny.IntOfUint32((_1777_v127).Cardinality()), _1778_v2, _dafny.IntOfInt64(151))
					}
				})(_1761_v127, _1599_v2))), (Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-785))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}((func(_1780_v127 _dafny.Sequence, _1781_v2 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1782_i13 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_dafny.IntOfUint32((_1780_v127).Cardinality()), _1781_v2, _dafny.IntOfInt64(151))
					}
				})(_1761_v127, _1599_v2)))).Cardinality()))).Uint32(), _1761_v127)).Cardinality()))).Uint32(), _dafny.SeqOf(_1599_v2)))
				var _1783_v133 _dafny.MultiSet
				_ = _1783_v133
				_1783_v133 = _dafny.MultiSetOf((_this).F6(), p1)
				var _1784_v134 _dafny.Sequence
				_ = _1784_v134
				_1784_v134 = _dafny.UnicodeSeqOfUtf8Bytes("orstbaa")
				var _1785_v135 _dafny.MultiSet
				_ = _1785_v135
				_1785_v135 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_1763_v129, _1764_v130), _dafny.SeqOf(_1761_v127, _1761_v127, _1761_v127, _1761_v127), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1770_v132).Contains((_this).Fm5(_1783_v133, globalState)) {
						return (_1770_v132).Get((_this).Fm5(_1783_v133, globalState)).(_dafny.Sequence)
					}
					return _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1761_v127), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1784_v134).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_1761_v127)).Cardinality()))).Uint32(), _1761_v127)
				})(), _1762_v128), _1764_v130, _1763_v129)
				_1599_v2 = (func() _dafny.Int {
					if (_1785_v135).Contains(_1762_v128) {
						return (_1785_v135).Multiplicity(_1762_v128)
					}
					return _1599_v2
				})()
				var _1786_v136 _dafny.Map
				_ = _1786_v136
				_1786_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1599_v2, _dafny.SetOf(_1769_v131))
				var _1787_v137 _dafny.Set
				_ = _1787_v137
				_1787_v137 = _dafny.SetOf(_dafny.CodePoint('j'), _dafny.CodePoint('g'), _1769_v131, (Companion_D3_.Create_DC10_(p1, _1769_v131, (_this).Fm6(_dafny.SeqOf(_1769_v131, _1769_v131), p1, globalState), (_this).F6(), p1)).Dtor_cf24())
				_1786_v136 = (_1786_v136).Update((_1761_v127).Select((Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_1761_v127).Cardinality()))).Uint32()).(_dafny.Int), _1787_v137)
			} else {
				var _1788_v138 _dafny.CodePoint
				_ = _1788_v138
				_1788_v138 = _dafny.CodePoint('q')
				var _1789_v139 _dafny.Sequence
				_ = _1789_v139
				_1789_v139 = _dafny.SeqOf(_1788_v138, _1788_v138, _1788_v138, _1788_v138, _1788_v138)
				var _1790_v140 D2
				_ = _1790_v140
				_1790_v140 = Companion_D2_.Create_DC5_(_1789_v139)
				var _1791_v141 _dafny.Set
				_ = _1791_v141
				_1791_v141 = _dafny.SetOf((_1790_v140).Dtor_cf13(), _dafny.UnicodeSeqOfUtf8Bytes("ev"), _dafny.UnicodeSeqOfUtf8Bytes("jfvgu"), _dafny.UnicodeSeqOfUtf8Bytes("al"), _1789_v139)
				var _1792_v142 _dafny.Sequence
				_ = _1792_v142
				_1792_v142 = _dafny.SeqOf(_1791_v141)
				var _rhs255 bool = p1
				_ = _rhs255
				var _rhs256 _dafny.Set = (func() _dafny.Set {
					if (_this).Fm6(_1789_v139, (_this).F6(), globalState) {
						return (_1792_v142).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_1599_v2, globalState), _dafny.IntOfUint32((_1792_v142).Cardinality()))).Uint32()).(_dafny.Set)
					}
					return _dafny.SetOf(_1789_v139)
				})()
				_ = _rhs256
				var _rhs257 bool = (_1597_v0).Select((Companion_Default___.SafeIndex(_1599_v2, _dafny.IntOfUint32((_1597_v0).Cardinality()))).Uint32()).(bool)
				_ = _rhs257
				var _rhs258 _dafny.Int = Companion_Default___.Fm0(_1599_v2, globalState)
				_ = _rhs258
				r0 = _rhs255
				_1791_v141 = _rhs256
				r0 = _rhs257
				r1 = _rhs258
				var _1793_v143 _dafny.Array
				_ = _1793_v143
				var _len0_55 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_55
				var _nw246 _dafny.Array
				_ = _nw246
				if _len0_55.Cmp(_dafny.Zero) == 0 {
					_nw246 = _dafny.NewArray(_len0_55)
				} else {
					var _init55 func(_dafny.Int) bool = func(_1794_i14 _dafny.Int) bool {
						return (_this).F6()
					}
					_ = _init55
					var _element0_55 = _init55(_dafny.Zero)
					_ = _element0_55
					_nw246 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
					(_nw246).ArraySet1(_element0_55, 0)
					var _nativeLen0_55 = (_len0_55).Int()
					_ = _nativeLen0_55
					for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
						(_nw246).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
					}
				}
				_1793_v143 = _nw246
				var _1795_v144 *C3
				_ = _1795_v144
				var _nw247 *C3 = New_C3_()
				_ = _nw247
				_nw247.Ctor__(_1793_v143)
				_1795_v144 = _nw247
				_1788_v138 = _1788_v138
				var _1796_v145 _dafny.Array
				_ = _1796_v145
				var _nw248 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
				_ = _nw248
				_1796_v145 = _nw248
				var _1797_v146 _dafny.Set
				_ = _1797_v146
				_1797_v146 = _dafny.SetOf((_this).F6(), p1, p1)
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1796_v145), 0))
				_ = _index306
				(_1796_v145).ArraySet1(_1797_v146, (_index306).Int())
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1796_v145), 0))
				_ = _index307
				(_1796_v145).ArraySet1(_1797_v146, (_index307).Int())
				var _nw249 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw249
				_1723_v97 = _nw249
			}
			var _1798_v147 D1
			_ = _1798_v147
			_1798_v147 = Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-27))).Uint32(), func(coer80 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg80 _dafny.Int) interface{} {
					return coer80(arg80)
				}
			}((func(_1799_p1 bool, _1800_v2 _dafny.Int) func(_dafny.Int) D0 {
				return func(_1801_i15 _dafny.Int) D0 {
					return Companion_D0_.Create_DC1_(_1799_p1, _1799_p1, _1800_v2)
				}
			})(p1, _1599_v2))))
			var _1802_v148 _dafny.Sequence
			_ = _1802_v148
			_1802_v148 = _dafny.SeqOf(_1798_v147, _1798_v147, _1798_v147)
			var _1803_v149 _dafny.Sequence
			_ = _1803_v149
			_1803_v149 = _dafny.UnicodeSeqOfUtf8Bytes("epdgsg")
			var _rhs259 _dafny.Array = p0
			_ = _rhs259
			var _rhs260 _dafny.Int = (_dafny.MultiSetFromSeq(_1802_v148)).Cardinality()
			_ = _rhs260
			var _rhs261 _dafny.Int = _1599_v2
			_ = _rhs261
			var _rhs262 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_1599_v2, (_dafny.Zero).Minus(_1599_v2)), (func() _dafny.Int {
				if p1 {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1599_v2, _1803_v149)).Cardinality()
				}
				return _1599_v2
			})()))
			_ = _rhs262
			_1723_v97 = _rhs259
			_1599_v2 = _rhs260
			r1 = _rhs261
			r1 = _rhs262
			var _1804_v150 _dafny.Array
			_ = _1804_v150
			var _nw250 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw250
			_1804_v150 = _nw250
			(globalState).F1 = _1804_v150
			var _1805_v151 _dafny.Set
			_ = _1805_v151
			_1805_v151 = _dafny.SetOf((_this).F6())
			_1805_v151 = _1805_v151
			_1597_v0 = _dafny.Companion_Sequence_.Concatenate(_1597_v0, _1597_v0)
		} else {
			var _1806_v152 _dafny.CodePoint
			_ = _1806_v152
			_1806_v152 = _dafny.CodePoint('j')
			var _1807_v153 _dafny.Sequence
			_ = _1807_v153
			_1807_v153 = _dafny.SeqOf(_1806_v152)
			var _rhs263 _dafny.Int = (Companion_Default___.SafeDivisionInt(_1599_v2, _1599_v2)).Minus(_1599_v2)
			_ = _rhs263
			var _rhs264 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1807_v153, _dafny.SeqOf(Companion_Default___.Fm22(_1807_v153, globalState)))).Cardinality())
			_ = _rhs264
			_1599_v2 = _rhs263
			_1599_v2 = _rhs264
			r0 = ((_1621_v21).Union(_1621_v21)).Equals(_1621_v21)
			var _rhs265 _dafny.Array = p0
			_ = _rhs265
			var _rhs266 bool = (_this).F6()
			_ = _rhs266
			var _rhs267 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1807_v153, _1807_v153), _dafny.Companion_Sequence_.Concatenate(_1807_v153, _dafny.UnicodeSeqOfUtf8Bytes("eyatql")))
			_ = _rhs267
			_1723_v97 = _rhs265
			r0 = _rhs266
			_1807_v153 = _rhs267
			r1 = (_1599_v2).Plus((_1599_v2).Minus(_1599_v2))
			var _1808_v154 _dafny.Sequence
			_ = _1808_v154
			_1808_v154 = _dafny.SeqOf(_1599_v2, _1599_v2, _1599_v2, _1599_v2)
			_1621_v21 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1808_v154, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1807_v153).Cardinality()), _dafny.IntOfUint32((_1808_v154).Cardinality()))).Uint32(), _1599_v2), _1808_v154))).Union(_1621_v21)
		}
		r0 = ((_1621_v21).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_1599_v2, _1599_v2)))).IsDisjointFrom(_1621_v21)
		r1 = _1599_v2
		var _1809_v155 D0
		_ = _1809_v155
		_1809_v155 = Companion_D0_.Create_DC0_(true)
		var _pat_let_tv106 = p1
		_ = _pat_let_tv106
		var _pat_let_tv107 = p1
		_ = _pat_let_tv107
		var _pat_let_tv108 = p1
		_ = _pat_let_tv108
		var _pat_let_tv109 = p1
		_ = _pat_let_tv109
		var _pat_let_tv110 = p1
		_ = _pat_let_tv110
		var _pat_let_tv111 = _1599_v2
		_ = _pat_let_tv111
		var _pat_let_tv112 = _1599_v2
		_ = _pat_let_tv112
		var _pat_let_tv113 = _1599_v2
		_ = _pat_let_tv113
		r2 = func(_source24 D0) bool {
			if _source24.Is_DC1() {
				var _1810___mcc_h4 bool = _source24.Get_().(D0_DC1).Cf1
				_ = _1810___mcc_h4
				var _1811___mcc_h5 bool = _source24.Get_().(D0_DC1).Cf2
				_ = _1811___mcc_h5
				var _1812___mcc_h6 _dafny.Int = _source24.Get_().(D0_DC1).Cf3
				_ = _1812___mcc_h6
				var _1813_cf3 _dafny.Int = _1812___mcc_h6
				_ = _1813_cf3
				var _1814_cf2 bool = _1811___mcc_h5
				_ = _1814_cf2
				var _1815_cf1 bool = _1810___mcc_h4
				_ = _1815_cf1
				return (_this).F6()
			} else if _source24.Is_DC2() {
				var _1816___mcc_h7 _dafny.Int = _source24.Get_().(D0_DC2).Cf4
				_ = _1816___mcc_h7
				var _1817___mcc_h8 bool = _source24.Get_().(D0_DC2).Cf5
				_ = _1817___mcc_h8
				var _1818___mcc_h9 bool = _source24.Get_().(D0_DC2).Cf6
				_ = _1818___mcc_h9
				var _1819___mcc_h10 _dafny.Int = _source24.Get_().(D0_DC2).Cf7
				_ = _1819___mcc_h10
				var _1820_cf7 _dafny.Int = _1819___mcc_h10
				_ = _1820_cf7
				var _1821_cf6 bool = _1818___mcc_h9
				_ = _1821_cf6
				var _1822_cf5 bool = _1817___mcc_h8
				_ = _1822_cf5
				var _1823_cf4 _dafny.Int = _1816___mcc_h7
				_ = _1823_cf4
				return (_dafny.MultiSetOf(Companion_D13_.Create_DC40_(_dafny.MultiSetOf((_this).F6(), _pat_let_tv106, _1821_cf6, false)))).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D13_.Create_DC40_(_dafny.MultiSetOf(_1822_cf5)), Companion_D13_.Create_DC40_(_dafny.MultiSetOf((_this).F6(), (_this).F6())), Companion_D13_.Create_DC40_(_dafny.MultiSetOf(_pat_let_tv107)), Companion_D13_.Create_DC40_(_dafny.MultiSetOf(_pat_let_tv108)), Companion_D13_.Create_DC40_(_dafny.MultiSetOf(_1821_cf6, _pat_let_tv109, true, _pat_let_tv110)))))
			} else {
				var _1824___mcc_h11 bool = _source24.Get_().(D0_DC0).Cf0
				_ = _1824___mcc_h11
				var _1825_cf0 bool = _1824___mcc_h11
				_ = _1825_cf0
				return !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_pat_let_tv111, _pat_let_tv112), _dafny.SeqOf(_pat_let_tv113)))
			}
		}(_1809_v155)
		return r0, r1, r2
	}
}
func (_this *C4) F6() bool {
	{
		return _this._f6
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f2 _dafny.Array
	_f3 _dafny.Array
	_f4 bool
	_f5 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = false
	_this._f5 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C5{}
var _ T1 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F2() _dafny.Array {
	return _this._f2
}
func (_this *C5) F2_set_(value _dafny.Array) {
	_this._f2 = value
}
func (_this *C5) F3() _dafny.Array {
	return _this._f3
}
func (_this *C5) F4() bool {
	return _this._f4
}
func (_this *C5) Ctor__(f5 bool, f2 _dafny.Array, f3 _dafny.Array, f4 bool) {
	{
		(_this)._f5 = f5
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *C5) Fm1(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		if (_this).F4() {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(672), (_this).F5())
		} else {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(780), (_this).F4())
		}
	}
}
func (_this *C5) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(661)
	}
}
func (_this *C5) Fm3(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ns")).Cardinality()), (func() _dafny.Int {
			if (_this).F5() {
				return (_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()))))).Cardinality()
			}
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf((_this).F5(), (_this).F5())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _dafny.IntOfInt64(431))).Cardinality())).Cardinality()))).Cardinality()))
		})())
	}
}
func (_this *C5) M0(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Map, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		r0 = (_this).F5()
		var _1826_v0 D0
		_ = _1826_v0
		_1826_v0 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(397), (_this).F4(), false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}(func(_1827_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		}))).Cardinality()))
		var _1828_i0 _dafny.Int
		_ = _1828_i0
		_1828_i0 = _dafny.Zero
		{
			for ((func() D0 {
				if (_this).F4() {
					return _1826_v0
				}
				return _1826_v0
			})()).Dtor_cf6() {
				{
					if (_1828_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1828_i0 = (_1828_i0).Plus(_dafny.One)
					if Companion_Default___.Fm4(_dafny.IntOfInt64(-585), (func() _dafny.Set {
						if false {
							return func() _dafny.Set {
								var _coll57 = _dafny.NewBuilder()
								_ = _coll57
								for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(361), _dafny.IntOfInt64(82))); ; {
									_compr_57, _ok65 := _iter65()
									if !_ok65 {
										break
									}
									var _1829_v1 _dafny.Int
									_1829_v1 = interface{}(_compr_57).(_dafny.Int)
									if ((_dafny.IntOfInt64(361)).Cmp(_1829_v1) <= 0) && ((_1829_v1).Cmp(_dafny.IntOfInt64(82)) < 0) {
										_coll57.Add(Companion_Default___.SafeDivisionInt(_1829_v1, p0))
									}
								}
								return _coll57.ToSet()
							}()
						}
						return _dafny.SetOf(p0, p0, p0)
					})(), globalState) {
						var _1830_v2 D0
						_ = _1830_v2
						_1830_v2 = Companion_D0_.Create_DC0_((_this).F5())
						var _1831_v3 _dafny.Map
						_ = _1831_v3
						_1831_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _1830_v2)
						_1831_v3 = _1831_v3
						var _1832_v4 _dafny.Sequence
						_ = _1832_v4
						_1832_v4 = _dafny.UnicodeSeqOfUtf8Bytes("mw")
						var _1833_v5 _dafny.Set
						_ = _1833_v5
						_1833_v5 = _dafny.SetOf(_dafny.IntOfUint32((_1832_v4).Cardinality()))
						_1833_v5 = _1833_v5
						var _1834_v6 _dafny.Sequence
						_ = _1834_v6
						_1834_v6 = _dafny.SeqOf((_this).F4())
						var _1835_v7 *C2
						_ = _1835_v7
						var _nw251 *C2 = New_C2_()
						_ = _nw251
						_nw251.Ctor__((_this).F4(), _dafny.IntOfUint32((_1834_v6).Cardinality()), (_this).F3(), (_this).F5(), _this.F2())
						_1835_v7 = _nw251
						var _1836_v8 _dafny.Array
						_ = _1836_v8
						var _nw252 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
						_ = _nw252
						_1836_v8 = _nw252
						var _1837_v9 _dafny.Array
						_ = _1837_v9
						var _nw253 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
						_ = _nw253
						_1837_v9 = _nw253
						var _1838_v10 _dafny.Sequence
						_ = _1838_v10
						_1838_v10 = _dafny.SeqOf(_1837_v9)
						var _1839_v11 _dafny.Array
						_ = _1839_v11
						_1839_v11 = (_1838_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1838_v10).Cardinality()))).Uint32()).(_dafny.Array)
						var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_1836_v8), 0))
						_ = _index308
						(_1836_v8).ArraySet1(_1839_v11, (_index308).Int())
						var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_1836_v8), 0))
						_ = _index309
						(_1836_v8).ArraySet1(_1837_v9, (_index309).Int())
						var _1840_v12 *C3
						_ = _1840_v12
						var _nw254 *C3 = New_C3_()
						_ = _nw254
						_nw254.Ctor__(_this.F2())
						_1840_v12 = _nw254
					} else {
						var _1841_v13 _dafny.Array
						_ = _1841_v13
						var _len0_56 _dafny.Int = _dafny.IntOfInt64(24)
						_ = _len0_56
						var _nw255 _dafny.Array
						_ = _nw255
						if _len0_56.Cmp(_dafny.Zero) == 0 {
							_nw255 = _dafny.NewArray(_len0_56)
						} else {
							var _init56 func(_dafny.Int) _dafny.Int = func(_1842_i2 _dafny.Int) _dafny.Int {
								return (_1842_i2).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gwl")).Cardinality()))
							}
							_ = _init56
							var _element0_56 = _init56(_dafny.Zero)
							_ = _element0_56
							_nw255 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
							(_nw255).ArraySet1(_element0_56, 0)
							var _nativeLen0_56 = (_len0_56).Int()
							_ = _nativeLen0_56
							for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
								(_nw255).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
							}
						}
						_1841_v13 = _nw255
						var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1841_v13), 0))
						_ = _index310
						(_1841_v13).ArraySet1(p0, (_index310).Int())
						var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1841_v13), 0))
						_ = _index311
						(_1841_v13).ArraySet1(p0, (_index311).Int())
						var _1843_v14 *C1
						_ = _1843_v14
						var _nw256 *C1 = New_C1_()
						_ = _nw256
						_nw256.Ctor__((_this).F3(), (Companion_Default___.Fm29(globalState)).Contains((_1841_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1841_v13), 0))).Int()).(_dafny.Int)), _this.F2())
						_1843_v14 = _nw256
						_1843_v14 = _1843_v14
						var _1844_v15 _dafny.Sequence
						_ = _1844_v15
						_1844_v15 = _dafny.UnicodeSeqOfUtf8Bytes("f")
						var _1845_v16 _dafny.MultiSet
						_ = _1845_v16
						_1845_v16 = _dafny.MultiSetOf(!((_this).F4()), (_this).F4(), (_this).F5())
						var _1846_v17 _dafny.CodePoint
						_ = _1846_v17
						_1846_v17 = _dafny.CodePoint('s')
						var _1847_v18 *C2
						_ = _1847_v18
						var _nw257 *C2 = New_C2_()
						_ = _nw257
						_nw257.Ctor__((_this).F4(), _dafny.IntOfInt64(403), (_this).F3(), _dafny.Companion_Sequence_.IsProperPrefixOf(_1844_v15, _dafny.Companion_Sequence_.Update(_1844_v15, (Companion_Default___.SafeIndex((_1845_v16).Cardinality(), _dafny.IntOfUint32((_1844_v15).Cardinality()))).Uint32(), _1846_v17)), _this.F2())
						_1847_v18 = _nw257
						var _1848_v19 _dafny.Map
						_ = _1848_v19
						_1848_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_1843_v14).Fm2(globalState), globalState), p0)
						var _1849_v20 _dafny.MultiSet
						_ = _1849_v20
						_1849_v20 = _dafny.MultiSetOf((_1847_v18).F8())
						_1848_v19 = (_1848_v19).Update(((_1849_v20).Intersection(_1849_v20)).Cardinality(), p0)
						var _arr55 _dafny.Array = _this.F2()
						_ = _arr55
						var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_this.F2()), 0))
						_ = _index312
						(_arr55).ArraySet1((_this).F5(), (_index312).Int())
						var _1850_v21 _dafny.Sequence
						_ = _1850_v21
						_1850_v21 = _dafny.SeqOf(Companion_Default___.Fm11((_1847_v18).F8(), !((_this).F5()), (_dafny.Zero).Minus((_1847_v18).F8()), globalState), _1844_v15, _dafny.UnicodeSeqOfUtf8Bytes("ylns"), _1844_v15)
						var _arr56 _dafny.Array = _this.F2()
						_ = _arr56
						var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_this.F2()), 0))
						_ = _index313
						var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1841_v13), 0))
						_ = _index314
						var _rhs268 _dafny.Array = (func() _dafny.Array {
							if (_this).F4() {
								return _1841_v13
							}
							return _1841_v13
						})()
						_ = _rhs268
						var _rhs269 bool = !(_1847_v18.F7)
						_ = _rhs269
						var _rhs270 _dafny.Int = (_1841_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1841_v13), 0))).Int()).(_dafny.Int)
						_ = _rhs270
						var _rhs271 _dafny.Sequence = _dafny.SeqOf(Companion_Default___.Fm13((_this).F5(), true, (_this).F4(), globalState), _dafny.Companion_Sequence_.Concatenate(_1844_v15, _dafny.UnicodeSeqOfUtf8Bytes("lpisqxnf")), _1844_v15, Companion_Default___.Fm11(p0, (_this).F4(), (_1847_v18).F8(), globalState))
						_ = _rhs271
						var _lhs167 _dafny.Array = _this.F2()
						_ = _lhs167
						var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_this.F2()), 0))
						_ = _lhs168
						var _lhs169 _dafny.Array = _1841_v13
						_ = _lhs169
						var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1841_v13), 0))
						_ = _lhs170
						_1841_v13 = _rhs268
						(_lhs167).ArraySet1(_rhs269, (_lhs168).Int())
						(_lhs169).ArraySet1(_rhs270, (_lhs170).Int())
						_1850_v21 = _rhs271
					}
					var _1851_v22 _dafny.Sequence
					_ = _1851_v22
					_1851_v22 = _dafny.UnicodeSeqOfUtf8Bytes("acgo")
					var _1852_v23 _dafny.Map
					_ = _1852_v23
					_1852_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1851_v22, p2)
					var _1853_v24 _dafny.Set
					_ = _1853_v24
					_1853_v24 = _dafny.SetOf((_this).F4())
					var _1854_v25 _dafny.Set
					_ = _1854_v25
					_1854_v25 = _dafny.SetOf((_1853_v24).Cardinality(), p0)
					var _1855_v26 _dafny.Map
					_ = _1855_v26
					_1855_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1854_v25).Cardinality(), Companion_Default___.Fm20((_this).F5(), globalState))
					var _1856_v27 _dafny.Sequence
					_ = _1856_v27
					_1856_v27 = _dafny.SeqOf(p2)
					var _1857_v28 _dafny.Map
					_ = _1857_v28
					_1857_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_1856_v27).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1856_v27).Cardinality()))).Uint32()).(_dafny.Map))
					_1852_v23 = (_1852_v23).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(482))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg82 _dafny.Int) interface{} {
							return coer82(arg82)
						}
					}(func(_1858_i3 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('a')
					})), Companion_Default___.Fm18((_this).F4(), _1855_v26, globalState)), (func() _dafny.Map {
						if (_1857_v28).Contains((_this).F5()) {
							return (_1857_v28).Get((_this).F5()).(_dafny.Map)
						}
						return p2
					})())
					r0 = (func() bool {
						if (_1855_v26).Contains(p0) {
							return (_1855_v26).Get(p0).(bool)
						}
						return (_this).F4()
					})()
					r0 = (_this).F5()
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1859_v29 _dafny.CodePoint
		_ = _1859_v29
		_1859_v29 = _dafny.CodePoint('q')
		_1859_v29 = _1859_v29
		var _1860_v30 _dafny.Array
		_ = _1860_v30
		var _len0_57 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_57
		var _nw258 _dafny.Array
		_ = _nw258
		if _len0_57.Cmp(_dafny.Zero) == 0 {
			_nw258 = _dafny.NewArray(_len0_57)
		} else {
			var _init57 func(_dafny.Int) _dafny.Int = (func(_1861_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1862_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1862_i4, _1861_p0)
				}
			})(p0)
			_ = _init57
			var _element0_57 = _init57(_dafny.Zero)
			_ = _element0_57
			_nw258 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
			(_nw258).ArraySet1(_element0_57, 0)
			var _nativeLen0_57 = (_len0_57).Int()
			_ = _nativeLen0_57
			for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
				(_nw258).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
			}
		}
		_1860_v30 = _nw258
		var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
		_ = _index315
		(_1860_v30).ArraySet1(p0, (_index315).Int())
		var _1863_v31 _dafny.MultiSet
		_ = _1863_v31
		_1863_v31 = _dafny.MultiSetOf(p0)
		var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
		_ = _index316
		(_1860_v30).ArraySet1((Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(326), p0)).Plus(((_1863_v31).Intersection(_1863_v31)).Cardinality()), (_index316).Int())
		var _1864_v32 _dafny.Map
		_ = _1864_v32
		_1864_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F5())
		var _arr57 _dafny.Array = _this.F2()
		_ = _arr57
		var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
		_ = _index317
		(_arr57).ArraySet1((p0).Cmp(_dafny.IntOfInt64(65)) < 0, (_index317).Int())
		var _1865_v33 _dafny.Array
		_ = _1865_v33
		var _nw259 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
		_ = _nw259
		_1865_v33 = _nw259
		var _1866_v34 T1
		_ = _1866_v34
		var _nw260 *C1 = New_C1_()
		_ = _nw260
		_nw260.Ctor__((_this).F3(), !((_this).F4()), _this.F2())
		_1866_v34 = _nw260
		var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_1865_v33), 0))
		_ = _index318
		(_1865_v33).ArraySet1(_1866_v34, (_index318).Int())
		var _arr58 _dafny.Array = _this.F2()
		_ = _arr58
		var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
		_ = _index319
		var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_1865_v33), 0))
		_ = _index320
		var _rhs272 bool = (func() bool {
			if (_this).F5() {
				return (_1866_v34).F4()
			}
			return !((_1866_v34).F4()) || (true)
		})()
		_ = _rhs272
		var _rhs273 _dafny.Map = _1864_v32
		_ = _rhs273
		var _rhs274 bool = ((_dafny.IntOfInt64(440)).Plus((_dafny.Zero).Minus(p0))).Cmp((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int)) < 0
		_ = _rhs274
		var _rhs275 T1 = _1866_v34
		_ = _rhs275
		var _rhs276 _dafny.Array = _1860_v30
		_ = _rhs276
		var _lhs171 _dafny.Array = _this.F2()
		_ = _lhs171
		var _lhs172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
		_ = _lhs172
		var _lhs173 _dafny.Array = _1865_v33
		_ = _lhs173
		var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_1865_v33), 0))
		_ = _lhs174
		r0 = _rhs272
		_1864_v32 = _rhs273
		(_lhs171).ArraySet1(_rhs274, (_lhs172).Int())
		(_lhs173).ArraySet1(_rhs275, (_lhs174).Int())
		_1860_v30 = _rhs276
		if (_1866_v34).F4() {
			var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
			_ = _index321
			(_1860_v30).ArraySet1(p0, (_index321).Int())
			if (func() bool {
				if (_1866_v34).F4() {
					return (_this).F5()
				}
				return (_1866_v34).F4()
			})() {
				var _1867_v35 _dafny.Sequence
				_ = _1867_v35
				_1867_v35 = _dafny.SeqOf(_dafny.IntOfInt64(229), _dafny.IntOfInt64(-848), p0)
				var _1868_v36 _dafny.Map
				_ = _1868_v36
				_1868_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_1867_v35, _dafny.SeqOf(p0)), (func() bool {
					if (_1864_v32).Contains((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int)) {
						return (_1864_v32).Get((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int)).(bool)
					}
					return (_1866_v34).F4()
				})())
				_1868_v36 = (_1868_v36).Update((p0).Cmp((_dafny.Zero).Minus((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int))) == 0, (_1866_v34).F4())
				var _1869_v37 *C0
				_ = _1869_v37
				var _nw261 *C0 = New_C0_()
				_ = _nw261
				_nw261.Ctor__()
				_1869_v37 = _nw261
				var _1870_v38 _dafny.Sequence
				_ = _1870_v38
				_1870_v38 = _dafny.SeqOf(_1869_v37)
				var _1871_v39 _dafny.Array
				_ = _1871_v39
				var _nwElement0_50 *C0 = _1869_v37
				_ = _nwElement0_50
				var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(6))
				_ = _nw262
				(_nw262).ArraySet1(_nwElement0_50, 0)
				(_nw262).ArraySet1(_1869_v37, 1)
				(_nw262).ArraySet1(_1869_v37, 2)
				(_nw262).ArraySet1(_1869_v37, 3)
				(_nw262).ArraySet1(_1869_v37, 4)
				(_nw262).ArraySet1((_1870_v38).Select((Companion_Default___.SafeIndex((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1870_v38).Cardinality()))).Uint32()).(*C0), 5)
				_1871_v39 = _nw262
				_1871_v39 = _1871_v39
				var _1872_v40 _dafny.Sequence
				_ = _1872_v40
				_1872_v40 = _dafny.UnicodeSeqOfUtf8Bytes("egcrltj")
				var _1873_v41 _dafny.Sequence
				_ = _1873_v41
				_1873_v41 = _dafny.SeqOf(!((_1866_v34).F4()))
				var _1874_v42 _dafny.MultiSet
				_ = _1874_v42
				_1874_v42 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm32(_dafny.IntOfUint32((_1872_v40).Cardinality()), (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), _1868_v36, globalState), _1873_v41))
				var _1875_v43 _dafny.Array
				_ = _1875_v43
				var _len0_58 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_58
				var _nw263 _dafny.Array
				_ = _nw263
				if _len0_58.Cmp(_dafny.Zero) == 0 {
					_nw263 = _dafny.NewArray(_len0_58)
				} else {
					var _init58 func(_dafny.Int) _dafny.MultiSet = func(_1876_i5 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf((_this).F4())
					}
					_ = _init58
					var _element0_58 = _init58(_dafny.Zero)
					_ = _element0_58
					_nw263 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
					(_nw263).ArraySet1(_element0_58, 0)
					var _nativeLen0_58 = (_len0_58).Int()
					_ = _nativeLen0_58
					for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
						(_nw263).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
					}
				}
				_1875_v43 = _nw263
				var _1877_v44 _dafny.MultiSet
				_ = _1877_v44
				_1877_v44 = _dafny.MultiSetOf(!((_this).F4()))
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1875_v43), 0))
				_ = _index322
				(_1875_v43).ArraySet1(_1877_v44, (_index322).Int())
				var _1878_v45 _dafny.Array
				_ = _1878_v45
				var _len0_59 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_59
				var _nw264 _dafny.Array
				_ = _nw264
				if _len0_59.Cmp(_dafny.Zero) == 0 {
					_nw264 = _dafny.NewArray(_len0_59)
				} else {
					var _init59 func(_dafny.Int) bool = func(_1879_i6 _dafny.Int) bool {
						return true
					}
					_ = _init59
					var _element0_59 = _init59(_dafny.Zero)
					_ = _element0_59
					_nw264 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
					(_nw264).ArraySet1(_element0_59, 0)
					var _nativeLen0_59 = (_len0_59).Int()
					_ = _nativeLen0_59
					for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
						(_nw264).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
					}
				}
				_1878_v45 = _nw264
				var _1880_v46 *C2
				_ = _1880_v46
				var _nw265 *C2 = New_C2_()
				_ = _nw265
				_nw265.Ctor__(false, (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), (_1866_v34).F3(), (_1866_v34).F4(), _1878_v45)
				_1880_v46 = _nw265
				var _1881_v47 _dafny.Set
				_ = _1881_v47
				_1881_v47 = _dafny.SetOf(_1880_v46)
				var _1882_v48 _dafny.Map
				_ = _1882_v48
				_1882_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), _1881_v47)
				var _1883_v49 *C4
				_ = _1883_v49
				var _nw266 *C4 = New_C4_()
				_ = _nw266
				_nw266.Ctor__(_1880_v46.F7)
				_1883_v49 = _nw266
				var _1884_v50 _dafny.Sequence
				_ = _1884_v50
				_1884_v50 = _dafny.SeqOf(_1883_v49, _1883_v49)
				var _1885_v51 _dafny.Map
				_ = _1885_v51
				_1885_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1880_v46, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1884_v50).Cardinality()), _1881_v47)).Update(p0, _1881_v47))
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1875_v43), 0))
				_ = _index323
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
				_ = _index324
				var _rhs277 _dafny.MultiSet = _1874_v42
				_ = _rhs277
				var _rhs278 _dafny.MultiSet = (_1877_v44).Difference(_dafny.MultiSetOf(_1880_v46.F7))
				_ = _rhs278
				var _rhs279 _dafny.Int = (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int)
				_ = _rhs279
				var _rhs280 _dafny.Map = (_1882_v48).Merge((func() _dafny.Map {
					if (_1885_v51).Contains(_1880_v46) {
						return (_1885_v51).Get(_1880_v46).(_dafny.Map)
					}
					return _1882_v48
				})())
				_ = _rhs280
				var _lhs175 _dafny.Array = _1875_v43
				_ = _lhs175
				var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1875_v43), 0))
				_ = _lhs176
				var _lhs177 _dafny.Array = _1860_v30
				_ = _lhs177
				var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
				_ = _lhs178
				_1874_v42 = _rhs277
				(_lhs175).ArraySet1(_rhs278, (_lhs176).Int())
				(_lhs177).ArraySet1(_rhs279, (_lhs178).Int())
				_1882_v48 = _rhs280
				var _1886_v52 _dafny.Array
				_ = _1886_v52
				var _nw267 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
				_ = _nw267
				_1886_v52 = _nw267
				var _1887_v53 _dafny.Sequence
				_ = _1887_v53
				_1887_v53 = _dafny.SeqOf(_dafny.SetOf((_this).F4(), (_1866_v34).F4()))
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1886_v52), 0))
				_ = _index325
				(_1886_v52).ArraySet1(_1887_v53, (_index325).Int())
				var _1888_v54 _dafny.Sequence
				_ = _1888_v54
				_1888_v54 = _dafny.SeqOf(_1887_v53)
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1886_v52), 0))
				_ = _index326
				(_1886_v52).ArraySet1((_1888_v54).Select((Companion_Default___.SafeIndex((_1880_v46).F8(), _dafny.IntOfUint32((_1888_v54).Cardinality()))).Uint32()).(_dafny.Sequence), (_index326).Int())
				var _1889_v55 _dafny.Set
				_ = _1889_v55
				_1889_v55 = _dafny.SetOf((_1880_v46).F8())
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
				_ = _index327
				(_1860_v30).ArraySet1(((_1889_v55).Intersection(_1889_v55)).Cardinality(), (_index327).Int())
			} else {
				var _1890_v56 *C4
				_ = _1890_v56
				var _nw268 *C4 = New_C4_()
				_ = _nw268
				_nw268.Ctor__(false)
				_1890_v56 = _nw268
				_1890_v56 = _1890_v56
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
				_ = _index328
				(_1860_v30).ArraySet1((_this).Fm2(globalState), (_index328).Int())
				r0 = ((_dafny.Zero).Minus((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int))).Cmp((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int)) < 0
				var _1891_v57 _dafny.Array
				_ = _1891_v57
				var _len0_60 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_60
				var _nw269 _dafny.Array
				_ = _nw269
				if _len0_60.Cmp(_dafny.Zero) == 0 {
					_nw269 = _dafny.NewArray(_len0_60)
				} else {
					var _init60 func(_dafny.Int) bool = (func(_1892_v34 T1) func(_dafny.Int) bool {
						return func(_1893_i7 _dafny.Int) bool {
							return (_1892_v34).F4()
						}
					})(_1866_v34)
					_ = _init60
					var _element0_60 = _init60(_dafny.Zero)
					_ = _element0_60
					_nw269 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
					(_nw269).ArraySet1(_element0_60, 0)
					var _nativeLen0_60 = (_len0_60).Int()
					_ = _nativeLen0_60
					for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
						(_nw269).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
					}
				}
				_1891_v57 = _nw269
				var _1894_v58 *C1
				_ = _1894_v58
				var _nw270 *C1 = New_C1_()
				_ = _nw270
				_nw270.Ctor__((_1866_v34).F3(), (_1866_v34).F4(), _1891_v57)
				_1894_v58 = _nw270
				var _1895_v59 _dafny.Set
				_ = _1895_v59
				_1895_v59 = _dafny.SetOf(_dafny.IntOfInt64(-573))
				r0 = Companion_Default___.Fm4(p0, _1895_v59, globalState)
			}
			var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
			_ = _index329
			(_1860_v30).ArraySet1((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), (_index329).Int())
			var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
			_ = _index330
			(_1860_v30).ArraySet1((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), (_index330).Int())
			var _1896_v60 _dafny.Sequence
			_ = _1896_v60
			_1896_v60 = _dafny.UnicodeSeqOfUtf8Bytes("aoltl")
			var _1897_v61 _dafny.Set
			_ = _1897_v61
			_1897_v61 = _dafny.SetOf(p0, (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int))
			var _1898_v62 _dafny.Map
			_ = _1898_v62
			_1898_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(960))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg83 _dafny.Int) interface{} {
					return coer83(arg83)
				}
			}((func(_1899_v29 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1900_i8 _dafny.Int) _dafny.CodePoint {
					return _1899_v29
				}
			})(_1859_v29)))).Cardinality()), _dafny.IntOfUint32((_1896_v60).Cardinality())), _1897_v61)
			var _1901_v63 _dafny.Sequence
			_ = _1901_v63
			_1901_v63 = _dafny.SeqOf((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), false)
			var _1902_v64 D12
			_ = _1902_v64
			_1902_v64 = Companion_D12_.Create_DC37_((_this).F5(), _1901_v63, p0)
			var _1903_v65 D3
			_ = _1903_v65
			_1903_v65 = Companion_D3_.Create_DC10_((_1866_v34).F4(), _1859_v29, (_1866_v34).F4(), (_this).F5(), (_this).F4())
			_1898_v62 = (_1898_v62).Update((_1902_v64).Dtor_cf69(), (_dafny.SetOf((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int))).Intersection(_dafny.SetOf((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_1903_v65).Dtor_cf27())).Cardinality()))))
		} else {
			var _1904_v66 _dafny.Sequence
			_ = _1904_v66
			_1904_v66 = _dafny.UnicodeSeqOfUtf8Bytes("dvygmlw")
			var _1905_v67 D10
			_ = _1905_v67
			_1905_v67 = Companion_D10_.Create_DC33_((_1866_v34).F4())
			var _1906_v68 _dafny.Sequence
			_ = _1906_v68
			_1906_v68 = _dafny.SeqOf(_1905_v67)
			var _1907_v69 D12
			_ = _1907_v69
			_1907_v69 = Companion_D12_.Create_DC36_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1904_v66, _1906_v68))
			var _1908_v70 D12
			_ = _1908_v70
			_1908_v70 = Companion_D12_.Create_DC39_(_1907_v69)
			var _pat_let_tv114 = _1907_v69
			_ = _pat_let_tv114
			var _source25 D12 = func(_pat_let47_0 D12) D12 {
				return func(_1909_dt__update__tmp_h0 D12) D12 {
					return func(_pat_let48_0 D12) D12 {
						return func(_1910_dt__update_hcf74_h0 D12) D12 {
							return Companion_D12_.Create_DC39_(_1910_dt__update_hcf74_h0)
						}(_pat_let48_0)
					}(_pat_let_tv114)
				}(_pat_let47_0)
			}(_1908_v70)
			_ = _source25
			if _source25.Is_DC37() {
				var _1911___mcc_h0 bool = _source25.Get_().(D12_DC37).Cf67
				_ = _1911___mcc_h0
				var _1912___mcc_h1 _dafny.Sequence = _source25.Get_().(D12_DC37).Cf68
				_ = _1912___mcc_h1
				var _1913___mcc_h2 _dafny.Int = _source25.Get_().(D12_DC37).Cf69
				_ = _1913___mcc_h2
				var _1914_cf69 _dafny.Int = _1913___mcc_h2
				_ = _1914_cf69
				var _1915_cf68 _dafny.Sequence = _1912___mcc_h1
				_ = _1915_cf68
				var _1916_cf67 bool = _1911___mcc_h0
				_ = _1916_cf67
				_1916_cf67 = (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)
				var _1917_v71 _dafny.Map
				_ = _1917_v71
				_1917_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _dafny.UnicodeSeqOfUtf8Bytes("dj"))
				var _1918_v72 _dafny.Set
				_ = _1918_v72
				_1918_v72 = _dafny.SetOf(_1914_cf69, (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), ((_1917_v71).Cardinality()).Minus(_dafny.IntOfInt64(-994)), (func() _dafny.Int {
					if (_1866_v34).F4() {
						return _1914_cf69
					}
					return (_dafny.Zero).Minus(_1914_cf69)
				})(), p0)
				_1918_v72 = ((_dafny.SetOf(p0, p0)).Difference(func() _dafny.Set {
					var _coll58 = _dafny.NewBuilder()
					_ = _coll58
					for _iter66 := _dafny.Iterate((_1863_v31).Elements()); ; {
						_compr_58, _ok66 := _iter66()
						if !_ok66 {
							break
						}
						var _1919_v73 _dafny.Int
						_1919_v73 = interface{}(_compr_58).(_dafny.Int)
						if (_1863_v31).Contains(_1919_v73) {
							_coll58.Add((_1919_v73).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(93))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg84 _dafny.Int) interface{} {
									return coer84(arg84)
								}
							}(func(_1920_i9 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('k')
							}))).Cardinality())))
						}
					}
					return _coll58.ToSet()
				}())).Union(_1918_v72)
				var _1921_v74 D2
				_ = _1921_v74
				_1921_v74 = Companion_D2_.Create_DC6_(_dafny.CodePoint('a'), _1914_cf69)
				var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
				_ = _index331
				(_1860_v30).ArraySet1((_1921_v74).Dtor_cf15(), (_index331).Int())
				var _1922_v75 *C0
				_ = _1922_v75
				var _nw271 *C0 = New_C0_()
				_ = _nw271
				_nw271.Ctor__()
				_1922_v75 = _nw271
				var _1923_v76 _dafny.MultiSet
				_ = _1923_v76
				_1923_v76 = _dafny.MultiSetOf((_this).F5())
				var _1924_v77 _dafny.Sequence
				_ = _1924_v77
				_1924_v77 = _dafny.SeqOf(_1904_v66, _1904_v66)
				var _1925_v78 D13
				_ = _1925_v78
				_1925_v78 = Companion_D13_.Create_DC42_(_1915_cf68, _1923_v76, _1924_v77, true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(693))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_1926_cf69 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1927_i10 _dafny.Int) _dafny.Int {
						return _1926_cf69
					}
				})(_1914_cf69))))
				var _1928_v79 _dafny.Map
				_ = _1928_v79
				_1928_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1859_v29, (_1866_v34).F4())
				var _1929_v80 _dafny.Map
				_ = _1929_v80
				_1929_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), (_1928_v79).Cardinality())
				var _pat_let_tv115 = _1866_v34
				_ = _pat_let_tv115
				var _arr59 _dafny.Array = _this.F2()
				_ = _arr59
				var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index332
				var _arr60 _dafny.Array = _this.F2()
				_ = _arr60
				var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index333
				var _rhs281 bool = Companion_Default___.Fm20(_1916_cf67, globalState)
				_ = _rhs281
				var _rhs282 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1925_v78).Dtor_cf77(), _1915_cf68)
				_ = _rhs282
				var _rhs283 _dafny.Int = ((_1929_v80).Merge(_1929_v80)).Cardinality()
				_ = _rhs283
				var _rhs284 *C0 = _1922_v75
				_ = _rhs284
				var _rhs285 bool = !((func(_pat_let49_0 D10) D10 {
					return func(_1930_dt__update__tmp_h1 D10) D10 {
						return func(_pat_let50_0 bool) D10 {
							return func(_1931_dt__update_hcf61_h0 bool) D10 {
								return Companion_D10_.Create_DC33_(_1931_dt__update_hcf61_h0)
							}(_pat_let50_0)
						}((_pat_let_tv115).F4())
					}(_pat_let49_0)
				}(_1905_v67)).Dtor_cf61()) || (_1916_cf67)
				_ = _rhs285
				var _lhs179 _dafny.Array = _this.F2()
				_ = _lhs179
				var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs180
				var _lhs181 _dafny.Array = _this.F2()
				_ = _lhs181
				var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs182
				(_lhs179).ArraySet1(_rhs281, (_lhs180).Int())
				_1915_cf68 = _rhs282
				_1914_cf69 = _rhs283
				_1922_v75 = _rhs284
				(_lhs181).ArraySet1(_rhs285, (_lhs182).Int())
			} else if _source25.Is_DC38() {
				var _1932___mcc_h3 _dafny.Set = _source25.Get_().(D12_DC38).Cf70
				_ = _1932___mcc_h3
				var _1933___mcc_h4 _dafny.Int = _source25.Get_().(D12_DC38).Cf71
				_ = _1933___mcc_h4
				var _1934___mcc_h5 *C2 = _source25.Get_().(D12_DC38).Cf72
				_ = _1934___mcc_h5
				var _1935___mcc_h6 _dafny.Int = _source25.Get_().(D12_DC38).Cf73
				_ = _1935___mcc_h6
				var _1936_cf73 _dafny.Int = _1935___mcc_h6
				_ = _1936_cf73
				var _1937_cf72 *C2 = _1934___mcc_h5
				_ = _1937_cf72
				var _1938_cf71 _dafny.Int = _1933___mcc_h4
				_ = _1938_cf71
				var _1939_cf70 _dafny.Set = _1932___mcc_h3
				_ = _1939_cf70
				var _1940_v81 _dafny.Array
				_ = _1940_v81
				var _nw272 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw272
				_1940_v81 = _nw272
				var _1941_v82 _dafny.Map
				_ = _1941_v82
				_1941_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _1940_v81)
				var _1942_v83 D15
				_ = _1942_v83
				_1942_v83 = Companion_D15_.Create_DC47_(_1940_v81)
				var _pat_let_tv116 = _1940_v81
				_ = _pat_let_tv116
				var _1943_v84 _dafny.Array
				_ = _1943_v84
				var _nwElement0_51 _dafny.Array = _1940_v81
				_ = _nwElement0_51
				var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(29))
				_ = _nw273
				(_nw273).ArraySet1(_nwElement0_51, 0)
				(_nw273).ArraySet1(_1940_v81, 1)
				(_nw273).ArraySet1(_1940_v81, 2)
				(_nw273).ArraySet1(_1940_v81, 3)
				(_nw273).ArraySet1(_1940_v81, 4)
				(_nw273).ArraySet1(_1940_v81, 5)
				(_nw273).ArraySet1(_1940_v81, 6)
				(_nw273).ArraySet1(_1940_v81, 7)
				(_nw273).ArraySet1((func() _dafny.Array {
					if (_1941_v82).Contains((_1866_v34).F4()) {
						return (_1941_v82).Get((_1866_v34).F4()).(_dafny.Array)
					}
					return _1940_v81
				})(), 8)
				(_nw273).ArraySet1(_1940_v81, 9)
				(_nw273).ArraySet1(_1940_v81, 10)
				(_nw273).ArraySet1(_1940_v81, 11)
				(_nw273).ArraySet1(_1940_v81, 12)
				(_nw273).ArraySet1((func() _dafny.Array {
					if _1937_cf72.F7 {
						return _1940_v81
					}
					return _1940_v81
				})(), 13)
				(_nw273).ArraySet1(_1940_v81, 14)
				(_nw273).ArraySet1(_1940_v81, 15)
				(_nw273).ArraySet1(_1940_v81, 16)
				(_nw273).ArraySet1(_1940_v81, 17)
				(_nw273).ArraySet1(_1940_v81, 18)
				(_nw273).ArraySet1(_1940_v81, 19)
				(_nw273).ArraySet1(_1940_v81, 20)
				(_nw273).ArraySet1(_1940_v81, 21)
				(_nw273).ArraySet1(_1940_v81, 22)
				(_nw273).ArraySet1(_1940_v81, 23)
				(_nw273).ArraySet1(_1940_v81, 24)
				(_nw273).ArraySet1((func(_pat_let51_0 D15) D15 {
					return func(_1944_dt__update__tmp_h2 D15) D15 {
						return func(_pat_let52_0 _dafny.Array) D15 {
							return func(_1945_dt__update_hcf86_h0 _dafny.Array) D15 {
								return Companion_D15_.Create_DC47_(_1945_dt__update_hcf86_h0)
							}(_pat_let52_0)
						}(_pat_let_tv116)
					}(_pat_let51_0)
				}(_1942_v83)).Dtor_cf86(), 25)
				(_nw273).ArraySet1(_1940_v81, 26)
				(_nw273).ArraySet1(_1940_v81, 27)
				(_nw273).ArraySet1(_1940_v81, 28)
				_1943_v84 = _nw273
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1943_v84), 0))
				_ = _index334
				(_1943_v84).ArraySet1(_1940_v81, (_index334).Int())
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1943_v84), 0))
				_ = _index335
				(_1943_v84).ArraySet1(_1940_v81, (_index335).Int())
				var _1946_v85 _dafny.Map
				_ = _1946_v85
				_1946_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
				var _1947_v86 D3
				_ = _1947_v86
				_1947_v86 = Companion_D3_.Create_DC9_(_1946_v85)
				var _1948_v88 _dafny.MultiSet
				_ = _1948_v88
				_1948_v88 = _dafny.MultiSetOf(((_1947_v86).Dtor_cf22()).Merge(_1946_v85), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1859_v29, true), _1946_v85, (func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter67 := _dafny.Iterate((_1946_v85).Keys().Elements()); ; {
						_compr_59, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _1949_v87 _dafny.CodePoint
						_1949_v87 = interface{}(_compr_59).(_dafny.CodePoint)
						if (_1946_v85).Contains(_1949_v87) {
							_coll59.Add(_1949_v87, (_1866_v34).F4())
						}
					}
					return _coll59.ToMap()
				}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1859_v29, (_1866_v34).F4())))
				var _1950_v89 D16
				_ = _1950_v89
				_1950_v89 = Companion_D16_.Create_DC50_(_1948_v88)
				var _arr61 _dafny.Array = _this.F2()
				_ = _arr61
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index336
				var _rhs286 bool = _1937_cf72.F7
				_ = _rhs286
				var _rhs287 _dafny.MultiSet = (_1950_v89).Dtor_cf93()
				_ = _rhs287
				var _rhs288 bool = (_1866_v34).F4()
				_ = _rhs288
				var _rhs289 bool = !(Companion_Default___.Fm20((_1866_v34).F4(), globalState))
				_ = _rhs289
				var _rhs290 _dafny.Array = _1860_v30
				_ = _rhs290
				var _lhs183 _dafny.Array = _this.F2()
				_ = _lhs183
				var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
				_ = _lhs184
				var _lhs185 *C2 = _1937_cf72
				_ = _lhs185
				var _lhs186 *C2 = _1937_cf72
				_ = _lhs186
				(_lhs183).ArraySet1(_rhs286, (_lhs184).Int())
				_1948_v88 = _rhs287
				_lhs185.F7 = _rhs288
				_lhs186.F7 = _rhs289
				_1860_v30 = _rhs290
				_1939_cf70 = (_1939_cf70).Intersection(_1939_cf70)
				var _1951_v90 D7
				_ = _1951_v90
				_1951_v90 = Companion_D7_.Create_DC23_(_1938_cf71)
				_1936_cf73 = (_1951_v90).Dtor_cf42()
			} else if _source25.Is_DC36() {
				var _1952___mcc_h7 _dafny.Map = _source25.Get_().(D12_DC36).Cf66
				_ = _1952___mcc_h7
				var _1953_cf66 _dafny.Map = _1952___mcc_h7
				_ = _1953_cf66
				r0 = !((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
				var _1954_v91 _dafny.MultiSet
				_ = _1954_v91
				_1954_v91 = _dafny.MultiSetOf(!((_this).F5()), !((_1866_v34).F4()), (_this).F5())
				var _1955_v92 _dafny.Map
				_ = _1955_v92
				_1955_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1866_v34).F4(), _dafny.IntOfUint32((_1904_v66).Cardinality()))
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index337
				((_this).F3()).ArraySet1((_1955_v92).Merge(_1955_v92), (_index337).Int())
				var _1956_v93 _dafny.Sequence
				_ = _1956_v93
				_1956_v93 = _dafny.SeqOf(false, (_1866_v34).F4(), (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), (_1866_v34).F4(), (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool))
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _index338
				var _rhs291 _dafny.MultiSet = ((_1954_v91).Intersection(_dafny.MultiSetFromSeq(_1956_v93))).Union(_dafny.MultiSetOf((_1866_v34).F4()))
				_ = _rhs291
				var _rhs292 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), p0)).Update((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool), (_1863_v31).Cardinality())
				_ = _rhs292
				var _lhs187 _dafny.Array = (_this).F3()
				_ = _lhs187
				var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen(((_this).F3()), 0))
				_ = _lhs188
				_1954_v91 = _rhs291
				(_lhs187).ArraySet1(_rhs292, (_lhs188).Int())
				var _1957_v94 _dafny.Map
				_ = _1957_v94
				_1957_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), (p0).Plus((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int)))
				_1957_v94 = _1957_v94
				var _1958_v95 _dafny.Map
				_ = _1958_v95
				_1958_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F4(), _1956_v93)
				_1958_v95 = (_1958_v95).Update(true, _1956_v93)
			} else {
				var _1959___mcc_h8 D12 = _source25.Get_().(D12_DC39).Cf74
				_ = _1959___mcc_h8
				var _1960_cf74 D12 = _1959___mcc_h8
				_ = _1960_cf74
				var _1961_v96 _dafny.Array
				_ = _1961_v96
				var _len0_61 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_61
				var _nw274 _dafny.Array
				_ = _nw274
				if _len0_61.Cmp(_dafny.Zero) == 0 {
					_nw274 = _dafny.NewArray(_len0_61)
				} else {
					var _init61 func(_dafny.Int) bool = func(_1962_i11 _dafny.Int) bool {
						return (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)
					}
					_ = _init61
					var _element0_61 = _init61(_dafny.Zero)
					_ = _element0_61
					_nw274 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
					(_nw274).ArraySet1(_element0_61, 0)
					var _nativeLen0_61 = (_len0_61).Int()
					_ = _nativeLen0_61
					for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
						(_nw274).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
					}
				}
				_1961_v96 = _nw274
				var _1963_v97 _dafny.Map
				_ = _1963_v97
				_1963_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1961_v96, _1961_v96)
				var _1964_v98 *C2
				_ = _1964_v98
				var _nw275 *C2 = New_C2_()
				_ = _nw275
				_nw275.Ctor__(!((_1866_v34).F4()), (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), (_1866_v34).F3(), (_this).F5(), (func() _dafny.Array {
					if (_1963_v97).Contains(_1961_v96) {
						return (_1963_v97).Get(_1961_v96).(_dafny.Array)
					}
					return _1961_v96
				})())
				_1964_v98 = _nw275
				_1859_v29 = (func() _dafny.CodePoint {
					if (_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool) {
						return _1859_v29
					}
					return _1859_v29
				})()
				var _1965_v99 _dafny.Set
				_ = _1965_v99
				_1965_v99 = _dafny.SetOf((_1964_v98).F8(), (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int))
				var _1966_v100 D15
				_ = _1966_v100
				_1966_v100 = Companion_D15_.Create_DC48_(_1964_v98.F7, _1961_v96, (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), Companion_Default___.Fm4(p0, _1965_v99, globalState), _1904_v66)
				var _1967_v101 *C3
				_ = _1967_v101
				var _nw276 *C3 = New_C3_()
				_ = _nw276
				_nw276.Ctor__((_1966_v100).Dtor_cf88())
				_1967_v101 = _nw276
				var _arr62 _dafny.Array = _this.F2()
				_ = _arr62
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))
				_ = _index339
				(_arr62).ArraySet1((_this).F4(), (_index339).Int())
			}
			var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
			_ = _index340
			(_1860_v30).ArraySet1(p0, (_index340).Int())
			var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
			_ = _index341
			var _rhs293 _dafny.CodePoint = _1859_v29
			_ = _rhs293
			var _rhs294 _dafny.Int = (_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int)
			_ = _rhs294
			var _lhs189 _dafny.Array = _1860_v30
			_ = _lhs189
			var _lhs190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
			_ = _lhs190
			_1859_v29 = _rhs293
			(_lhs189).ArraySet1(_rhs294, (_lhs190).Int())
			var _1968_v102 _dafny.Map
			_ = _1968_v102
			_1968_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_this.F2()), 0))).Int()).(bool)), (_this).F4())
			var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))
			_ = _index342
			(_1860_v30).ArraySet1((_1968_v102).Cardinality(), (_index342).Int())
			var _1969_v103 _dafny.Sequence
			_ = _1969_v103
			_1969_v103 = _dafny.SeqOf(_1968_v102, _1968_v102, _1968_v102, Companion_Default___.Fm36((_1866_v34).F4(), globalState))
			var _1970_v104 D7
			_ = _1970_v104
			_1970_v104 = Companion_D7_.Create_DC23_((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int))
			var _pat_let_tv117 = _1860_v30
			_ = _pat_let_tv117
			var _pat_let_tv118 = _1860_v30
			_ = _pat_let_tv118
			_1968_v102 = (_1969_v103).Select((Companion_Default___.SafeIndex((func(_pat_let53_0 D7) D7 {
				return func(_1971_dt__update__tmp_h3 D7) D7 {
					return func(_pat_let54_0 _dafny.Int) D7 {
						return func(_1972_dt__update_hcf42_h0 _dafny.Int) D7 {
							return Companion_D7_.Create_DC23_(_1972_dt__update_hcf42_h0)
						}(_pat_let54_0)
					}((_pat_let_tv118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_pat_let_tv117), 0))).Int()).(_dafny.Int))
				}(_pat_let53_0)
			}(_1970_v104)).Dtor_cf42(), _dafny.IntOfUint32((_1969_v103).Cardinality()))).Uint32()).(_dafny.Map)
		}
		var _1973_v105 _dafny.Array
		_ = _1973_v105
		var _nw277 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
		_ = _nw277
		_1973_v105 = _nw277
		var _1974_v106 _dafny.Sequence
		_ = _1974_v106
		_1974_v106 = _dafny.SeqOf(_1973_v105)
		r0 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1974_v106, _dafny.SeqOf(_1973_v105, _1973_v105)), (Companion_Default___.SafeIndex((_1860_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1860_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1974_v106, _dafny.SeqOf(_1973_v105, _1973_v105))).Cardinality()))).Uint32(), _1973_v105), _1973_v105)
		return r0
	}
}
func (_this *C5) M1(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Array, p3 _dafny.Array, globalState *GlobalState) (_dafny.MultiSet, bool) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 bool = false
		_ = r1
		var _1975_v0 T0
		_ = _1975_v0
		var _nw278 *C2 = New_C2_()
		_ = _nw278
		_nw278.Ctor__(true, (_dafny.Zero).Minus(p0), (_this).F3(), (_this).F4(), _this.F2())
		_1975_v0 = _nw278
		_1975_v0 = _1975_v0
		r1 = (_this).F5()
		var _1976_v1 _dafny.Map
		_ = _1976_v1
		_1976_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if !((_this).F4()) {
				return _1975_v0.F2()
			}
			return _1975_v0.F2()
		})(), ((_this).F5()) == (true))
		_1976_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1975_v0.F2(), Companion_Default___.Fm20((_this).F5(), globalState))
		if (func() bool {
			if (p0).Cmp(p0) == 0 {
				return (_this).F5()
			}
			return (_this).F4()
		})() {
			r1 = (_dafny.IntOfInt64(407)).Cmp(p0) >= 0
			var _1977_v2 _dafny.Map
			_ = _1977_v2
			_1977_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), (_this).F4())
			var _1978_v3 *C1
			_ = _1978_v3
			var _nw279 *C1 = New_C1_()
			_ = _nw279
			_nw279.Ctor__((_this).F3(), Companion_Default___.Fm4(p0, _dafny.SetOf((_dafny.Zero).Minus(p0), p0, (_1977_v2).Cardinality()), globalState), _1975_v0.F2())
			_1978_v3 = _nw279
			var _1979_v4 _dafny.Int
			_ = _1979_v4
			_1979_v4 = _dafny.IntOfInt64(819)
			_1979_v4 = _1979_v4
			_1979_v4 = p0
			var _1980_v5 *C4
			_ = _1980_v5
			var _nw280 *C4 = New_C4_()
			_ = _nw280
			_nw280.Ctor__((_this).F5())
			_1980_v5 = _nw280
			_1980_v5 = _1980_v5
		} else {
			var _1981_v6 _dafny.Array
			_ = _1981_v6
			var _nw281 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw281
			_1981_v6 = _nw281
			var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))
			_ = _index343
			(_1981_v6).ArraySet1(p0, (_index343).Int())
			var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))
			_ = _index344
			(_1981_v6).ArraySet1(Companion_Default___.SafeModuloInt(p0, Companion_Default___.SafeModuloInt((Companion_Default___.Fm37(p0, true, p0, globalState)).Cardinality(), p0)), (_index344).Int())
			var _arr63 _dafny.Array = _1975_v0.F2()
			_ = _arr63
			var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
			_ = _index345
			(_arr63).ArraySet1(false, (_index345).Int())
			var _1982_v7 _dafny.MultiSet
			_ = _1982_v7
			_1982_v7 = _dafny.MultiSetOf((_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int))
			var _arr64 _dafny.Array = _1975_v0.F2()
			_ = _arr64
			var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
			_ = _index346
			(_arr64).ArraySet1((_dafny.MultiSetOf((_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int), p0, p0, (_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int))).IsProperSubsetOf(_1982_v7), (_index346).Int())
			var _1983_v8 _dafny.Array
			_ = _1983_v8
			_1983_v8 = _1981_v6
			_1981_v6 = (_1983_v8)
			var _1984_v9 _dafny.Map
			_ = _1984_v9
			_1984_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool))
			var _1985_v10 _dafny.MultiSet
			_ = _1985_v10
			_1985_v10 = _dafny.MultiSetOf((_this).F5())
			var _1986_v11 D1
			_ = _1986_v11
			_1986_v11 = Companion_D1_.Create_DC4_((_1984_v9).Update((_dafny.Zero).Minus((_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int)), (_this).F4()), (_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool), (_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int), (_dafny.MultiSetOf((_this).F5(), (_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool))).IsDisjointFrom(_1985_v10))
			var _source26 D1 = _1986_v11
			_ = _source26
			if _source26.Is_DC4() {
				var _1987___mcc_h0 _dafny.Map = _source26.Get_().(D1_DC4).Cf9
				_ = _1987___mcc_h0
				var _1988___mcc_h1 bool = _source26.Get_().(D1_DC4).Cf10
				_ = _1988___mcc_h1
				var _1989___mcc_h2 _dafny.Int = _source26.Get_().(D1_DC4).Cf11
				_ = _1989___mcc_h2
				var _1990___mcc_h3 bool = _source26.Get_().(D1_DC4).Cf12
				_ = _1990___mcc_h3
				var _1991_cf12 bool = _1990___mcc_h3
				_ = _1991_cf12
				var _1992_cf11 _dafny.Int = _1989___mcc_h2
				_ = _1992_cf11
				var _1993_cf10 bool = _1988___mcc_h1
				_ = _1993_cf10
				var _1994_cf9 _dafny.Map = _1987___mcc_h0
				_ = _1994_cf9
				var _arr65 _dafny.Array = _1975_v0.F2()
				_ = _arr65
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
				_ = _index347
				(_arr65).ArraySet1((_1984_v9).Equals(_1994_cf9), (_index347).Int())
				var _1995_v12 _dafny.Sequence
				_ = _1995_v12
				_1995_v12 = _dafny.UnicodeSeqOfUtf8Bytes("bns")
				_1992_cf11 = ((func() _dafny.Int {
					if false {
						return (_1975_v0).Fm2(globalState)
					}
					return _dafny.IntOfUint32((_1995_v12).Cardinality())
				})()).Times(((_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int)).Times(((_1982_v7).Update(p0, Companion_Default___.Abs(_1992_cf11))).Cardinality()))
				var _1996_v13 _dafny.Array
				_ = _1996_v13
				var _nw282 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
				_ = _nw282
				_1996_v13 = _nw282
				var _1997_v14 _dafny.Sequence
				_ = _1997_v14
				_1997_v14 = _dafny.SeqOf(_1991_cf12)
				var _1998_v15 _dafny.Sequence
				_ = _1998_v15
				_1998_v15 = _dafny.SeqOf((_1997_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1995_v12).Cardinality()), _dafny.IntOfUint32((_1997_v14).Cardinality()))).Uint32()).(bool))
				var _1999_v16 D4
				_ = _1999_v16
				_1999_v16 = Companion_D4_.Create_DC11_(_1997_v14)
				var _2000_v17 D5
				_ = _2000_v17
				_2000_v17 = Companion_D5_.Create_DC18_(Companion_D5_.Create_DC15_(_1999_v16, (_this).F5()))
				var _2001_v18 _dafny.Map
				_ = _2001_v18
				_2001_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _2000_v17)
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_1996_v13), 0))
				_ = _index348
				(_1996_v13).ArraySet1(_2001_v18, (_index348).Int())
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_1996_v13), 0))
				_ = _index349
				(_1996_v13).ArraySet1((_2001_v18).Merge(_2001_v18), (_index349).Int())
				var _2002_v20 _dafny.Set
				_ = _2002_v20
				_2002_v20 = _dafny.SetOf((_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int))
				var _2003_v21 _dafny.Sequence
				_ = _2003_v21
				_2003_v21 = _dafny.SeqOf(func() _dafny.Set {
					var _coll60 = _dafny.NewBuilder()
					_ = _coll60
					for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(569), _dafny.IntOfInt64(153))); ; {
						_compr_60, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _2004_v19 _dafny.Int
						_2004_v19 = interface{}(_compr_60).(_dafny.Int)
						if ((_dafny.IntOfInt64(569)).Cmp(_2004_v19) <= 0) && ((_2004_v19).Cmp(_dafny.IntOfInt64(153)) < 0) {
							_coll60.Add((_2004_v19).Times(p0))
						}
					}
					return _coll60.ToSet()
				}(), _2002_v20)
				var _2005_v22 D0
				_ = _2005_v22
				_2005_v22 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_1997_v14).Cardinality()), _1991_cf12, (_this).F5(), ((_2003_v21).Select((Companion_Default___.SafeIndex((_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2003_v21).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality())
				var _2006_v23 *C2
				_ = _2006_v23
				var _nw283 *C2 = New_C2_()
				_ = _nw283
				_nw283.Ctor__((_2005_v22).Dtor_cf6(), ((_dafny.Zero).Minus(p0)).Plus((_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int)), (_this).F3(), !((_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool)), _this.F2())
				_2006_v23 = _nw283
			} else {
				var _2007___mcc_h4 _dafny.Sequence = _source26.Get_().(D1_DC3).Cf8
				_ = _2007___mcc_h4
				var _2008_cf8 _dafny.Sequence = _2007___mcc_h4
				_ = _2008_cf8
				var _2009_v24 *C0
				_ = _2009_v24
				var _nw284 *C0 = New_C0_()
				_ = _nw284
				_nw284.Ctor__()
				_2009_v24 = _nw284
				var _2010_v25 _dafny.Map
				_ = _2010_v25
				_2010_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F5()), !((_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool)))
				var _2011_v26 _dafny.Sequence
				_ = _2011_v26
				_2011_v26 = _dafny.UnicodeSeqOfUtf8Bytes("qahpyyt")
				var _2012_v27 _dafny.Map
				_ = _2012_v27
				_2012_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2010_v25).Cardinality(), _dafny.IntOfUint32((_2011_v26).Cardinality()))
				var _2013_v28 _dafny.Set
				_ = _2013_v28
				_2013_v28 = _dafny.SetOf((_1981_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))).Int()).(_dafny.Int))
				var _arr66 _dafny.Array = _1975_v0.F2()
				_ = _arr66
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
				_ = _index350
				var _arr67 _dafny.Array = _1975_v0.F2()
				_ = _arr67
				var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
				_ = _index351
				var _arr68 _dafny.Array = _1975_v0.F2()
				_ = _arr68
				var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
				_ = _index352
				var _rhs295 bool = Companion_Default___.Fm4(p0, (_2013_v28).Intersection(_2013_v28), globalState)
				_ = _rhs295
				var _rhs296 bool = false
				_ = _rhs296
				var _rhs297 bool = (_2009_v24).Fm7((_this).F4(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool), (_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool))).Cardinality(), (false) == ((_this).F5()), p0, globalState)
				_ = _rhs297
				var _rhs298 _dafny.Map = (_2012_v27).Merge(_2012_v27)
				_ = _rhs298
				var _lhs191 _dafny.Array = _1975_v0.F2()
				_ = _lhs191
				var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
				_ = _lhs192
				var _lhs193 _dafny.Array = _1975_v0.F2()
				_ = _lhs193
				var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
				_ = _lhs194
				var _lhs195 _dafny.Array = _1975_v0.F2()
				_ = _lhs195
				var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
				_ = _lhs196
				(_lhs191).ArraySet1(_rhs295, (_lhs192).Int())
				(_lhs193).ArraySet1(_rhs296, (_lhs194).Int())
				(_lhs195).ArraySet1(_rhs297, (_lhs196).Int())
				_2012_v27 = _rhs298
				var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1981_v6), 0))
				_ = _index353
				(_1981_v6).ArraySet1((_dafny.Zero).Minus(p0), (_index353).Int())
				var _arr69 _dafny.Array = _1975_v0.F2()
				_ = _arr69
				var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))
				_ = _index354
				(_arr69).ArraySet1((_this).F5(), (_index354).Int())
			}
			var _2014_v29 _dafny.Sequence
			_ = _2014_v29
			_2014_v29 = _dafny.UnicodeSeqOfUtf8Bytes("cemk")
			var _2015_v30 _dafny.Set
			_ = _2015_v30
			_2015_v30 = _dafny.SetOf(p1)
			var _2016_v31 _dafny.Map
			_ = _2016_v31
			_2016_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(670))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg86 _dafny.Int) interface{} {
					return coer86(arg86)
				}
			}((func(_2017_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2018_i0 _dafny.Int) _dafny.Int {
					return _2017_p0
				}
			})(p0))), _2015_v30)
			var _2019_v32 _dafny.Sequence
			_ = _2019_v32
			_2019_v32 = _dafny.SeqOf((_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool), (_1975_v0.F2()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_1975_v0.F2()), 0))).Int()).(bool))
			r1 = !_dafny.Companion_Sequence_.Contains(_2019_v32, ((func() _dafny.Set {
				if (_2016_v31).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(136))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_2020_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2021_i1 _dafny.Int) _dafny.Int {
						return _2020_p0
					}
				})(p0)))) {
					return (_2016_v31).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(136))).Uint32(), func(coer88 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg88 _dafny.Int) interface{} {
							return coer88(arg88)
						}
					}((func(_2022_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2023_i1 _dafny.Int) _dafny.Int {
							return _2022_p0
						}
					})(p0)))).(_dafny.Set)
				}
				return _2015_v30
			})()).IsSubsetOf(_dafny.SetOf(p1, Companion_Default___.Fm22(_2014_v29, globalState))))
		}
		var _hi10 _dafny.Int = (_this).Fm2(globalState)
		_ = _hi10
		for _2024_i2 := p0; _2024_i2.Cmp(_hi10) < 0; _2024_i2 = _2024_i2.Plus(_dafny.One) {
			r1 = (_2024_i2).Cmp(_2024_i2) < 0
			var _2025_v33 _dafny.Array
			_ = _2025_v33
			var _nw285 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw285
			_2025_v33 = _nw285
			_2025_v33 = _2025_v33
			var _2026_v34 *C4
			_ = _2026_v34
			var _nw286 *C4 = New_C4_()
			_ = _nw286
			_nw286.Ctor__((_this).F5())
			_2026_v34 = _nw286
			var _2027_v35 _dafny.MultiSet
			_ = _2027_v35
			_2027_v35 = _dafny.MultiSetOf((_this).F5())
			var _2028_v36 D13
			_ = _2028_v36
			_2028_v36 = Companion_D13_.Create_DC40_((_2027_v35).Union(_dafny.MultiSetOf((_this).F5())))
			_2028_v36 = _2028_v36
		}
		var _2029_v37 _dafny.Int
		_ = _2029_v37
		_2029_v37 = _dafny.IntOfInt64(586)
		var _rhs299 bool = (func() bool {
			if !(Companion_Default___.Fm20((_this).F4(), globalState)) {
				return !(true) || (true)
			}
			return Companion_Default___.Fm20(!((_this).F5()), globalState)
		})()
		_ = _rhs299
		var _rhs300 _dafny.Int = p0
		_ = _rhs300
		r1 = _rhs299
		_2029_v37 = _rhs300
		var _2030_v38 _dafny.MultiSet
		_ = _2030_v38
		_2030_v38 = _dafny.MultiSetOf(_1975_v0.F2(), _this.F2(), _this.F2(), _this.F2())
		r0 = _2030_v38
		r1 = (_2029_v37).Cmp(_2029_v37) >= 0
		return r0, r1
	}
}
func (_this *C5) F5() bool {
	{
		return _this._f5
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
