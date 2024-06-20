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
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(643))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	})), _dafny.UnicodeSeqOfUtf8Bytes("svlade")), _dafny.UnicodeSeqOfUtf8Bytes("ls"))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-106))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(767)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eoqmvc")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Set, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC9_((func() _dafny.CodePoint {
		if true {
			return _dafny.CodePoint('v')
		}
		return _dafny.CodePoint('v')
	})())
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false)), _dafny.SeqOf(false)))).Cardinality()).Plus(((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(492), _dafny.IntOfInt64(-555))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(492)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(-555)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_1_v0, (func() _dafny.Map {
					var _coll1 = _dafny.NewMapBuilder()
					_ = _coll1
					for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(581), _dafny.IntOfInt64(375))); ; {
						_compr_1, _ok1 := _iter1()
						if !_ok1 {
							break
						}
						var _2_v1 _dafny.Int
						_2_v1 = interface{}(_compr_1).(_dafny.Int)
						if ((_dafny.IntOfInt64(581)).Cmp(_2_v1) <= 0) && ((_2_v1).Cmp(_dafny.IntOfInt64(375)) < 0) {
							_coll1.Add(Companion_Default___.SafeModuloInt(_2_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vfftcgcr")).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(540))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg1 _dafny.Int) interface{} {
									return coer1(arg1)
								}
							}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('s')
							}))).Cardinality()))
						}
					}
					return _coll1.ToMap()
				}()).Cardinality()))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality()).Times(_dafny.IntOfInt64(431)))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-246))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_4_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(316)
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-167))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_5_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-88)
		})))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _6_v0 _dafny.Int
			_6_v0 = interface{}(_compr_2).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-246))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_4_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(316)
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-167))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_5_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-88)
			}))), _6_v0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_6_v0, (_dafny.SetOf(_dafny.IntOfInt64(453), _dafny.IntOfInt64(243))).Cardinality()), !(true))
			}
		}
		return _coll2.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("hafht")
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.CodePoint, p1 _dafny.CodePoint, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("scdrbxig"), false)
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
		if !(!(!(false))) {
			return true
		}
		return !(true)
	})()), !(!(false)))
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Sequence {
	if true {
		return _dafny.SeqOf(false)
	} else if false {
		return _dafny.SeqOf(false)
	} else {
		return (Companion_D14_.Create_DC36_((Companion_D14_.Create_DC36_(_dafny.SeqOf(true, false))).Dtor_cf53())).Dtor_cf53()
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 bool, p2 _dafny.Set, p3 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 D14 = Companion_D14_.Create_DC35_(_dafny.SeqOf(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-786), _dafny.IntOfInt64(405))); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _7_v1 _dafny.Int
				_7_v1 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(-786)).Cmp(_7_v1) <= 0) && ((_7_v1).Cmp(_dafny.IntOfInt64(405)) < 0) {
					_coll4.Add((_7_v1).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(628))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg6 _dafny.Int) interface{} {
							return coer6(arg6)
						}
					}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('o')
					}))).Cardinality())), false)
				}
			}
			return _coll4.ToMap()
		}()).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _9_v0 _dafny.Int
			_9_v0 = interface{}(_compr_3).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-786), _dafny.IntOfInt64(405))); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _7_v1 _dafny.Int
					_7_v1 = interface{}(_compr_5).(_dafny.Int)
					if ((_dafny.IntOfInt64(-786)).Cmp(_7_v1) <= 0) && ((_7_v1).Cmp(_dafny.IntOfInt64(405)) < 0) {
						_coll5.Add((_7_v1).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(628))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg7 _dafny.Int) interface{} {
								return coer7(arg7)
							}
						}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('o')
						}))).Cardinality())), false)
					}
				}
				return _coll5.ToMap()
			}()).Contains(_9_v0) {
				_coll3.Add((_9_v0).Minus(_dafny.IntOfInt64(696)), false)
			}
		}
		return _coll3.ToMap()
	}(), func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(227), _dafny.IntOfInt64(175))); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _10_v2 _dafny.Int
			_10_v2 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(227)).Cmp(_10_v2) <= 0) && ((_10_v2).Cmp(_dafny.IntOfInt64(175)) < 0) {
				_coll6.Add(Companion_Default___.SafeModuloInt(_10_v2, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-493))).Cardinality())), false)
			}
		}
		return _coll6.ToMap()
	}()))
	_ = _source0
	if _source0.Is_DC36() {
		var _11___mcc_h0 _dafny.Sequence = _source0.Get_().(D14_DC36).Cf53
		_ = _11___mcc_h0
		var _12_cf53 _dafny.Sequence = _11___mcc_h0
		_ = _12_cf53
		return _dafny.UnicodeSeqOfUtf8Bytes("fsqaxwqte")
	} else if _source0.Is_DC35() {
		var _13___mcc_h1 _dafny.Sequence = _source0.Get_().(D14_DC35).Cf52
		_ = _13___mcc_h1
		var _14_cf52 _dafny.Sequence = _13___mcc_h1
		_ = _14_cf52
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dxhjglaic"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-301))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_15_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})))
	} else {
		var _16___mcc_h2 D14 = _source0.Get_().(D14_DC37).Cf54
		_ = _16___mcc_h2
		var _17_cf54 D14 = _16___mcc_h2
		_ = _17_cf54
		return _dafny.UnicodeSeqOfUtf8Bytes("mplnkdwif")
	}
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfInt64(72))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(77), (_dafny.MultiSetOf(!((Companion_D20_.Create_DC55_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(186), (_dafny.SetOf(false, false)).Cardinality())).Cardinality()), ((Companion_D15_.Create_DC38_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(116))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality()), true))).Dtor_cf55()).Cardinality(), false, true)).Dtor_cf87()), true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xmniuvfes")).Cardinality()), _dafny.IntOfInt64(728), _dafny.IntOfInt64(310)))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Set, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll7 = _dafny.NewBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(23), _dafny.IntOfInt64(808))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(23)).Cmp(_19_v0) <= 0) && ((_19_v0).Cmp(_dafny.IntOfInt64(808)) < 0) {
				_coll7.Add((_19_v0).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kwgcuel")).Cardinality())))
			}
		}
		return _coll7.ToSet()
	}()).Union(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-804), _dafny.IntOfInt64(56))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _20_v1 _dafny.Int
			_20_v1 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(-804)).Cmp(_20_v1) <= 0) && ((_20_v1).Cmp(_dafny.IntOfInt64(56)) < 0) {
				_coll8.Add((_20_v1).Times(_dafny.IntOfInt64(867)), false)
			}
		}
		return _coll8.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(-849), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tm")).Cardinality()), _dafny.IntOfInt64(84), _dafny.IntOfInt64(249))).Cardinality()), _dafny.IntOfInt64(843), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('a'))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((false) || (!(!(false))), false)
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true, false)).Intersection(_dafny.MultiSetOf(false, true))).Intersection(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(true))))
}
func (_static *CompanionStruct_Default___) Fm35(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(113)).Plus(_dafny.IntOfInt64(-13)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-78))).Contains(false))
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(false)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(true, false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(false, true))))
}
func (_static *CompanionStruct_Default___) Fm37(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, false)).Difference(_dafny.SetOf(false))).Union((_dafny.SetOf(true)).Union(_dafny.SetOf(false, true)))
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source1 D12 = Companion_D12_.Create_DC33_(true)
	_ = _source1
	if _source1.Is_DC33() {
		var _21___mcc_h0 bool = _source1.Get_().(D12_DC33).Cf50
		_ = _21___mcc_h0
		var _22_cf50 bool = _21___mcc_h0
		_ = _22_cf50
		return _dafny.CodePoint('s')
	} else {
		var _23___mcc_h1 _dafny.Set = _source1.Get_().(D12_DC32).Cf49
		_ = _23___mcc_h1
		var _24_cf49 _dafny.Set = _23___mcc_h1
		_ = _24_cf49
		if true {
			return _dafny.CodePoint('w')
		} else {
			return _dafny.CodePoint('t')
		}
	}
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(669), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(930)), _dafny.MultiSetOf(_dafny.IntOfInt64(897), _dafny.IntOfInt64(342), (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(503))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_25_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(162)
	}))).Cardinality())))).Cardinality())), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(732))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fnqkq")).Cardinality()))).Cardinality())), (func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf(_dafny.IntOfInt64(157), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rysmqpp")).Cardinality()))
		}
		return _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-412)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), _dafny.IntOfInt64(193), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ng")).Cardinality()), false)).Cardinality(), _dafny.CodePoint('x'))).Cardinality())).Cardinality()), _dafny.IntOfInt64(-685))).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 bool, p2 bool, p3 D8, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_((func() _dafny.Int {
		if true {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("clskkt")).Cardinality())
		}
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_26_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality())
	})())
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hjghubcnr")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, !((Companion_D19_.Create_DC52_(_dafny.IntOfInt64(272), true, true)).Dtor_cf82()))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-738))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_27_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))).Cardinality()), (func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cxpjsdcn")).Cardinality())))).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _28_v0 _dafny.Int
			_28_v0 = interface{}(_compr_9).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cxpjsdcn")).Cardinality()))), _28_v0) {
				_coll9.Add(Companion_Default___.SafeDivisionInt(_28_v0, (_dafny.MultiSetOf(_dafny.IntOfInt64(854))).Cardinality()))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(448), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-717), _dafny.IntOfInt64(806))))
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D1 {
	if (!(true)) == (!(true)) {
		return Companion_D1_.Create_DC4_(_dafny.MultiSetOf(_dafny.IntOfInt64(-88)))
	} else {
		return Companion_D1_.Create_DC4_(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(819))).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(func() _dafny.Set {
		var _coll10 = _dafny.NewBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(858), _dafny.IntOfInt64(165))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _29_v0 _dafny.Int
			_29_v0 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(858)).Cmp(_29_v0) <= 0) && ((_29_v0).Cmp(_dafny.IntOfInt64(165)) < 0) {
				_coll10.Add((_29_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality())).Cardinality())))
			}
		}
		return _coll10.ToSet()
	}(), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ade")).Cardinality()))).Cardinality())), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("euwroheny")).Cardinality())))).Difference((_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("icwgdatuh")).Cardinality()), _dafny.IntOfInt64(632)))).Intersection(_dafny.MultiSetOf(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-518), false)).Cardinality()), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pfnisofy")).Cardinality())), _dafny.SetOf(_dafny.IntOfInt64(-276)))))
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf((Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("frlkbglny"), !(false))).Dtor_cf3(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-612))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_30_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))), _dafny.IntOfInt64(698))
}
func (_static *CompanionStruct_Default___) Fm45(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("klqhbvu")).Cardinality()), false)), _dafny.SeqOf(func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-567), _dafny.IntOfInt64(334))); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _31_v0 _dafny.Int
			_31_v0 = interface{}(_compr_11).(_dafny.Int)
			if ((_dafny.IntOfInt64(-567)).Cmp(_31_v0) <= 0) && ((_31_v0).Cmp(_dafny.IntOfInt64(334)) < 0) {
				_coll11.Add((_31_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), (Companion_D12_.Create_DC33_(true)).Dtor_cf50())
			}
		}
		return _coll11.ToMap()
	}(), func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(393))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_32_i0 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(true), _dafny.SeqOf(false))).Cardinality()))
		}))).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _33_v1 _dafny.Int
			_33_v1 = interface{}(_compr_12).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(393))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_32_i0 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(true), _dafny.SeqOf(false))).Cardinality()))
			})), _33_v1) {
				_coll12.Add(Companion_Default___.SafeModuloInt(_33_v1, (_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality())), !(!(true)))
			}
		}
		return _coll12.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm46(p0 bool, globalState *GlobalState) D16 {
	return Companion_D16_.Create_DC42_(!(false), true, (true) == (true))
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(29))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_34_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(168))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_35_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})), _dafny.UnicodeSeqOfUtf8Bytes("pabhldsjv"), _dafny.UnicodeSeqOfUtf8Bytes("cksrrod"), _dafny.UnicodeSeqOfUtf8Bytes("mwrut"))).Union(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ddnqnu"), _dafny.UnicodeSeqOfUtf8Bytes("nb")))).Intersection((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("itm"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_36_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})), _dafny.UnicodeSeqOfUtf8Bytes("juvss")))).Intersection(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(669))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_37_i3 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})))))
}
func (_static *CompanionStruct_Default___) Fm48(p0 bool, p1 bool, p2 bool, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC35_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(582), _dafny.IntOfUint32((_dafny.SeqOf(true, true, !(false))).Cardinality()))).Cardinality(), !(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(343))).Cardinality()), false)))
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC36_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false), _dafny.SeqOf(false, false)))
}
func (_static *CompanionStruct_Default___) Fm50(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.MultiSetOf(Companion_D2_.Create_DC6_(_dafny.IntOfInt64(-174))))).Union(_dafny.MultiSetOf(_dafny.MultiSetOf(Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("udq")).Cardinality())), Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ivkmw")).Cardinality())))))).Intersection(_dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D2_.Create_DC6_(_dafny.IntOfInt64(214)), Companion_D2_.Create_DC6_(_dafny.IntOfInt64(533)))), _dafny.MultiSetOf(Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(654))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg20 _dafny.Int) interface{} {
			return coer20(arg20)
		}
	}(func(_38_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dxvtbcyv")).Cardinality()), true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())), Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(286))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_39_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))).Cardinality()))), _dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D2_.Create_DC6_((_dafny.MultiSetOf(false)).Cardinality()))), _dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D2_.Create_DC6_((_dafny.MultiSetOf(Companion_D16_.Create_DC42_(false, false, true), Companion_D16_.Create_DC42_(!(true), !(false), false), Companion_D16_.Create_DC42_(!(false), false, false))).Cardinality()), Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("resvq")).Cardinality())), Companion_D2_.Create_DC6_((_dafny.MultiSetOf(true)).Cardinality()), Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lhcaqip")).Cardinality()))).Cardinality())), Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("og")).Cardinality())))), _dafny.MultiSetOf(Companion_D2_.Create_DC6_(_dafny.IntOfInt64(473)), Companion_D2_.Create_DC6_((func() _dafny.Map {
		var _coll13 = _dafny.NewMapBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-474), _dafny.IntOfInt64(-262))); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _40_v0 _dafny.Int
			_40_v0 = interface{}(_compr_13).(_dafny.Int)
			if ((_dafny.IntOfInt64(-474)).Cmp(_40_v0) <= 0) && ((_40_v0).Cmp(_dafny.IntOfInt64(-262)) < 0) {
				_coll13.Add((_40_v0).Times(_dafny.IntOfInt64(357)), true)
			}
		}
		return _coll13.ToMap()
	}()).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC31_(Companion_D11_.Create_DC31_(Companion_D11_.Create_DC30_(true, _dafny.IntOfInt64(684))))
}
func (_static *CompanionStruct_Default___) Fm52(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(Companion_D11_.Create_DC31_(Companion_D11_.Create_DC28_(func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_dafny.SeqOf((func() _dafny.Map {
					var _coll15 = _dafny.NewMapBuilder()
					_ = _coll15
					for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(334), _dafny.IntOfInt64(762))); ; {
						_compr_15, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _41_v0 _dafny.Int
						_41_v0 = interface{}(_compr_15).(_dafny.Int)
						if ((_dafny.IntOfInt64(334)).Cmp(_41_v0) <= 0) && ((_41_v0).Cmp(_dafny.IntOfInt64(762)) < 0) {
							_coll15.Add(Companion_Default___.SafeDivisionInt(_41_v0, _dafny.IntOfInt64(610)), (_dafny.SetOf(!(true))).Cardinality())
						}
					}
					return _coll15.ToMap()
				}()).Cardinality())).Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _42_v1 _dafny.Int
					_42_v1 = interface{}(_compr_14).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((func() _dafny.Map {
						var _coll16 = _dafny.NewMapBuilder()
						_ = _coll16
						for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(334), _dafny.IntOfInt64(762))); ; {
							_compr_16, _ok16 := _iter16()
							if !_ok16 {
								break
							}
							var _41_v0 _dafny.Int
							_41_v0 = interface{}(_compr_16).(_dafny.Int)
							if ((_dafny.IntOfInt64(334)).Cmp(_41_v0) <= 0) && ((_41_v0).Cmp(_dafny.IntOfInt64(762)) < 0) {
								_coll16.Add(Companion_Default___.SafeDivisionInt(_41_v0, _dafny.IntOfInt64(610)), (_dafny.SetOf(!(true))).Cardinality())
							}
						}
						return _coll16.ToMap()
					}()).Cardinality()), _42_v1) {
						_coll14.Add((_42_v1).Times(_dafny.IntOfUint32((_dafny.SeqOf(true, false, false)).Cardinality())))
					}
				}
				return _coll14.ToSet()
			}())), Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_()))
		}
		return _dafny.SeqOf(Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_()), Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_()), Companion_D11_.Create_DC31_(Companion_D11_.Create_DC28_(_dafny.SetOf(_dafny.IntOfInt64(-644)))))
	})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer22 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_43_i0 _dafny.Int) D11 {
		return Companion_D11_.Create_DC31_(Companion_D11_.Create_DC28_(_dafny.SetOf(_dafny.IntOfInt64(695))))
	})), _dafny.SeqOf(Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_()), Companion_D11_.Create_DC31_(Companion_D11_.Create_DC31_(Companion_D11_.Create_DC30_(true, (func() _dafny.Map {
		var _coll17 = _dafny.NewMapBuilder()
		_ = _coll17
		for _iter17 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fviyvd")).Cardinality()), _dafny.CodePoint('h'))).Cardinality(), _dafny.IntOfInt64(-349))).Cardinality(), false)).Keys().Elements()); ; {
			_compr_17, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _44_v2 _dafny.Int
			_44_v2 = interface{}(_compr_17).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fviyvd")).Cardinality()), _dafny.CodePoint('h'))).Cardinality(), _dafny.IntOfInt64(-349))).Cardinality(), false)).Contains(_44_v2) {
				_coll17.Add((_44_v2).Plus(_dafny.IntOfInt64(461)), _dafny.MultiSetOf(true))
			}
		}
		return _coll17.ToMap()
	}()).Cardinality()))))))
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Int, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC10_(_dafny.IntOfInt64(-282), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-116), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm54(p0 bool, p1 _dafny.Map, p2 _dafny.MultiSet, globalState *GlobalState) D19 {
	var _source2 D12 = Companion_D12_.Create_DC32_(_dafny.SetOf(false, false, false))
	_ = _source2
	if _source2.Is_DC33() {
		var _45___mcc_h0 bool = _source2.Get_().(D12_DC33).Cf50
		_ = _45___mcc_h0
		var _46_cf50 bool = _45___mcc_h0
		_ = _46_cf50
		return Companion_D19_.Create_DC53_()
	} else {
		var _47___mcc_h1 _dafny.Set = _source2.Get_().(D12_DC32).Cf49
		_ = _47___mcc_h1
		var _48_cf49 _dafny.Set = _47___mcc_h1
		_ = _48_cf49
		return Companion_D19_.Create_DC53_()
	}
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, p1 _dafny.Int, p2 bool, p3 D14, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), true)), (func() _dafny.Map {
		var _coll18 = _dafny.NewMapBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _dafny.CodePoint('t'))).Keys().Elements()); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _49_v0 _dafny.CodePoint
			_49_v0 = interface{}(_compr_18).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _dafny.CodePoint('t'))).Contains(_49_v0) {
				_coll18.Add(_49_v0, false)
			}
		}
		return _coll18.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), true)))
}
func (_static *CompanionStruct_Default___) Fm56(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC22_((_dafny.SetOf(!(false), false)).IsSubsetOf(_dafny.SetOf(!(false))), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cwvgtdwty"), _dafny.UnicodeSeqOfUtf8Bytes("q"))).Cardinality()), (_dafny.MultiSetOf(false, true, true, true, true)).IsProperSubsetOf(_dafny.MultiSetOf(false, true, true)), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(498))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_50_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Cardinality()
	}))).Cardinality())).Times(_dafny.IntOfInt64(-797)))
}
func (_static *CompanionStruct_Default___) Fm57(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(Companion_D7_.Create_DC22_(false, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality(), false, _dafny.IntOfInt64(483)), Companion_D7_.Create_DC22_(true, _dafny.IntOfInt64(-171), false, _dafny.IntOfInt64(747)))).Difference(_dafny.MultiSetOf(Companion_D7_.Create_DC22_(false, _dafny.IntOfInt64(640), false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(219), _dafny.IntOfInt64(-399))).Cardinality()), Companion_D7_.Create_DC22_(!(!(true)), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.CodePoint('t'))).Cardinality())), false, (_dafny.SetOf(false, true, true)).Cardinality())))).Intersection(_dafny.MultiSetOf(Companion_D7_.Create_DC22_(!(true), _dafny.IntOfInt64(898), false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-133))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_51_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm58(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-157), Companion_D21_.Create_DC57_(_dafny.IntOfInt64(428), _dafny.IntOfInt64(475), Companion_D5_.Create_DC14_(_dafny.CodePoint('p'), true, (_dafny.SetOf(true)).Cardinality(), (Companion_D18_.Create_DC49_(true, !(false), _dafny.CodePoint('o'), false)).Dtor_cf76())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("irgs")).Cardinality()), Companion_D21_.Create_DC57_((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(315))).Cardinality(), _dafny.IntOfInt64(639), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(202))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}(func(_52_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	}))).Cardinality())), _dafny.IntOfInt64(680), _dafny.IntOfInt64(30))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yptolnu")).Cardinality()), Companion_D5_.Create_DC14_(_dafny.CodePoint('o'), false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tusd")).Cardinality()), false)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(490), Companion_D21_.Create_DC57_(_dafny.IntOfInt64(814), _dafny.IntOfInt64(-888), Companion_D5_.Create_DC14_(_dafny.CodePoint('i'), !(true), _dafny.IntOfInt64(374), true)))))
}
func (_static *CompanionStruct_Default___) Fm59(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-493))), _dafny.IntOfInt64(469))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xamnu")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm60(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) D19 {
	return Companion_D19_.Create_DC52_(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("i")).Cardinality()), _dafny.IntOfInt64(994)), (_dafny.MultiSetOf(false)).IsSubsetOf(_dafny.MultiSetOf(false)), !(!(false)))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _53_v0 bool
	_ = _53_v0
	_53_v0 = false
	var _54_v1 _dafny.Int
	_ = _54_v1
	_54_v1 = _dafny.IntOfInt64(215)
	var _55_v2 _dafny.Map
	_ = _55_v2
	_55_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_53_v0), _54_v1)
	var _56_v3 _dafny.Set
	_ = _56_v3
	_56_v3 = _dafny.SetOf(_54_v1)
	var _57_v4 _dafny.Map
	_ = _57_v4
	_57_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _53_v0)
	var _58_v5 _dafny.Sequence
	_ = _58_v5
	_58_v5 = _dafny.SeqOf(_53_v0)
	var _59_v6 _dafny.MultiSet
	_ = _59_v6
	_59_v6 = _dafny.MultiSetOf(_58_v5)
	var _60_globalState *GlobalState
	_ = _60_globalState
	var _nw0 *GlobalState = New_GlobalState_()
	_ = _nw0
	_nw0.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_53_v0, _53_v0), _dafny.SeqOf(_53_v0)), true, _dafny.CodePoint('t'), _dafny.IntOfInt64(51), _dafny.IntOfInt64(503), true, _55_v2, _dafny.IntOfInt64(231), _56_v3, true, true, (_57_v4).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_53_v0), _53_v0)), _dafny.IntOfInt64(38), _dafny.IntOfInt64(113), true, _dafny.IntOfInt64(964), _59_v6, true, _dafny.IntOfInt64(454), true)
	_60_globalState = _nw0
	var _61_v7 *C4
	_ = _61_v7
	var _nw1 *C4 = New_C4_()
	_ = _nw1
	_nw1.Ctor__()
	_61_v7 = _nw1
	var _hi0 _dafny.Int = _54_v1
	_ = _hi0
	for _62_i0 := _54_v1; _62_i0.Cmp(_hi0) < 0; _62_i0 = _62_i0.Plus(_dafny.One) {
		var _63_v8 _dafny.Array
		_ = _63_v8
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_0
		var _nw2 _dafny.Array
		_ = _nw2
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw2 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_64_v0 bool) func(_dafny.Int) bool {
				return func(_65_i1 _dafny.Int) bool {
					return _64_v0
				}
			})(_53_v0)
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
		_63_v8 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))
		_ = _index0
		(_63_v8).ArraySet1(!((func() bool {
			if _53_v0 {
				return true
			}
			return _53_v0
		})()), (_index0).Int())
		var _66_v9 _dafny.Sequence
		_ = _66_v9
		_66_v9 = _dafny.UnicodeSeqOfUtf8Bytes("eabrtymy")
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))
		_ = _index1
		var _rhs0 bool = !(true) || (_53_v0)
		_ = _rhs0
		var _rhs1 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}(func(_67_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))
		_ = _rhs1
		var _rhs2 _dafny.Int = _54_v1
		_ = _rhs2
		var _rhs3 bool = (func() bool {
			if _53_v0 {
				return (((_dafny.MultiSetFromSeq(_58_v5)).Update(_53_v0, Companion_Default___.Abs(_62_i0))).Cardinality()).Cmp(Companion_Default___.Fm17(_54_v1, _60_globalState)) == 0
			}
			return _53_v0
		})()
		_ = _rhs3
		var _rhs4 _dafny.Int = _62_i0
		_ = _rhs4
		var _lhs0 _dafny.Array = _63_v8
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))
		_ = _lhs1
		(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
		_66_v9 = _rhs1
		_54_v1 = _rhs2
		_53_v0 = _rhs3
		_54_v1 = _rhs4
		_58_v5 = _58_v5
		var _68_v10 _dafny.Sequence
		_ = _68_v10
		_68_v10 = _dafny.SeqOf(_66_v9, _66_v9, _66_v9, _66_v9)
		_68_v10 = _68_v10
		if true {
			var _69_v11 _dafny.Map
			_ = _69_v11
			_69_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v0, Companion_D4_.Create_DC11_(_68_v10))
			var _70_v12 _dafny.Map
			_ = _70_v12
			_70_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_v4).Cardinality(), _63_v8)
			var _71_v13 _dafny.Set
			_ = _71_v13
			_71_v13 = _dafny.SetOf((func() _dafny.Array {
				if (_70_v12).Contains(_dafny.IntOfInt64(898)) {
					return (_70_v12).Get(_dafny.IntOfInt64(898)).(_dafny.Array)
				}
				return _63_v8
			})())
			var _72_v14 D1
			_ = _72_v14
			_72_v14 = Companion_D1_.Create_DC2_(_62_i0)
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))
			_ = _index2
			var _rhs5 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(918), (_dafny.SetOf(_54_v1, (_dafny.Zero).Minus(_62_i0), _62_i0)).Cardinality()))).Plus(_54_v1)
			_ = _rhs5
			var _rhs6 bool = (((_61_v7).Fm23(_69_v11, _dafny.IntOfUint32((_66_v9).Cardinality()), (_71_v13).Cardinality(), _60_globalState)).Times(_dafny.IntOfUint32((_66_v9).Cardinality()))).Cmp(((_56_v3).Intersection(_dafny.SetOf((_72_v14).Dtor_cf2()))).Cardinality()) <= 0
			_ = _rhs6
			var _lhs2 *GlobalState = _60_globalState
			_ = _lhs2
			var _lhs3 _dafny.Array = _63_v8
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))
			_ = _lhs4
			_lhs2.F15 = _rhs5
			(_lhs3).ArraySet1(_rhs6, (_lhs4).Int())
			(_60_globalState).F4 = _dafny.IntOfInt64(642)
			var _73_v15 D0
			_ = _73_v15
			_73_v15 = Companion_D0_.Create_DC1_(_63_v8)
			var _74_v16 _dafny.Map
			_ = _74_v16
			_74_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_63_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))).Int()).(bool), _63_v8)
			_73_v15 = Companion_D0_.Create_DC1_((func() _dafny.Array {
				if (_74_v16).Contains((_63_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))).Int()).(bool)) {
					return (_74_v16).Get((_63_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _63_v8
			})())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))
			_ = _index3
			var _rhs7 bool = (_63_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))).Int()).(bool)
			_ = _rhs7
			var _rhs8 _dafny.Int = (func() _dafny.Int {
				if _53_v0 {
					return (_dafny.Zero).Minus(_54_v1)
				}
				return _54_v1
			})()
			_ = _rhs8
			var _lhs5 _dafny.Array = _63_v8
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))
			_ = _lhs6
			(_lhs5).ArraySet1(_rhs7, (_lhs6).Int())
			_54_v1 = _rhs8
			var _75_v17 _dafny.MultiSet
			_ = _75_v17
			_75_v17 = _dafny.MultiSetOf(_53_v0, (_63_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))).Int()).(bool))
			var _76_v18 _dafny.Array
			_ = _76_v18
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw3
			_76_v18 = _nw3
			var _77_v19 D2
			_ = _77_v19
			_77_v19 = Companion_D2_.Create_DC7_(false, _75_v17, _76_v18)
			var _78_v20 D2
			_ = _78_v20
			_78_v20 = Companion_D2_.Create_DC8_(_77_v19)
			var _79_v21 *C3
			_ = _79_v21
			var _nw4 *C3 = New_C3_()
			_ = _nw4
			_nw4.Ctor__(_78_v20, (_63_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))).Int()).(bool))
			_79_v21 = _nw4
			var _80_v22 _dafny.Map
			_ = _80_v22
			_80_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v21, (_62_i0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ojorbrkgu")).Cardinality())))
			_80_v22 = (func() _dafny.Map {
				if _79_v21.F25 {
					return (_80_v22).Update(_79_v21, _62_i0)
				}
				return (_80_v22).Merge(_80_v22)
			})()
		} else {
			var _81_v23 _dafny.MultiSet
			_ = _81_v23
			_81_v23 = _dafny.MultiSetOf(_62_i0, _62_i0, (_dafny.IntOfInt64(86)).Times(_54_v1))
			_81_v23 = _81_v23
			var _82_v24 _dafny.CodePoint
			_ = _82_v24
			_82_v24 = _dafny.CodePoint('w')
			var _83_v25 _dafny.Map
			_ = _83_v25
			_83_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-880), _82_v24)
			_83_v25 = (_83_v25).Update(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(953))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_84_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_85_i3 _dafny.Int) _dafny.Int {
					return _84_v1
				}
			})(_54_v1)))).Cardinality()), _62_i0), (func() _dafny.CodePoint {
				if !((_63_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_63_v8), 0))).Int()).(bool)) {
					return _82_v24
				}
				return _82_v24
			})())
			(_60_globalState).F15 = _54_v1
			var _86_v26 *C8
			_ = _86_v26
			var _nw5 *C8 = New_C8_()
			_ = _nw5
			_nw5.Ctor__()
			_86_v26 = _nw5
			var _87_v27 _dafny.Array
			_ = _87_v27
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw6
			_87_v27 = _nw6
			_87_v27 = _87_v27
		}
	}
	var _88_v28 *C6
	_ = _88_v28
	var _nw7 *C6 = New_C6_()
	_ = _nw7
	_nw7.Ctor__()
	_88_v28 = _nw7
	var _89_v29 _dafny.MultiSet
	_ = _89_v29
	_89_v29 = _dafny.MultiSetOf(_88_v28)
	var _90_v30 D20
	_ = _90_v30
	_90_v30 = Companion_D20_.Create_DC55_(_54_v1, _54_v1, _53_v0, !(_dafny.MultiSetOf(_88_v28)).Equals(_89_v29))
	var _source3 D20 = _90_v30
	_ = _source3
	if _source3.Is_DC55() {
		var _91___mcc_h0 _dafny.Int = _source3.Get_().(D20_DC55).Cf85
		_ = _91___mcc_h0
		var _92___mcc_h1 _dafny.Int = _source3.Get_().(D20_DC55).Cf86
		_ = _92___mcc_h1
		var _93___mcc_h2 bool = _source3.Get_().(D20_DC55).Cf87
		_ = _93___mcc_h2
		var _94___mcc_h3 bool = _source3.Get_().(D20_DC55).Cf88
		_ = _94___mcc_h3
		var _95_cf88 bool = _94___mcc_h3
		_ = _95_cf88
		var _96_cf87 bool = _93___mcc_h2
		_ = _96_cf87
		var _97_cf86 _dafny.Int = _92___mcc_h1
		_ = _97_cf86
		var _98_cf85 _dafny.Int = _91___mcc_h0
		_ = _98_cf85
		var _99_v31 _dafny.Sequence
		_ = _99_v31
		_99_v31 = _dafny.SeqOf(_97_cf86)
		var _100_v32 _dafny.Sequence
		_ = _100_v32
		_100_v32 = _dafny.UnicodeSeqOfUtf8Bytes("dnextlf")
		(_60_globalState).F15 = (_54_v1).Minus((_99_v31).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_100_v32).Cardinality()), _dafny.IntOfUint32((_99_v31).Cardinality()))).Uint32()).(_dafny.Int))
		var _101_v33 D16
		_ = _101_v33
		_101_v33 = Companion_D16_.Create_DC42_(false, _96_cf87, _95_cf88)
		var _102_v34 _dafny.Sequence
		_ = _102_v34
		_102_v34 = _dafny.SeqOf(_101_v33, _101_v33, _101_v33)
		var _rhs9 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(740))).Uint32(), func(coer28 func(_dafny.Int) D16) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_103_v33 D16) func(_dafny.Int) D16 {
			return func(_104_i4 _dafny.Int) D16 {
				return _103_v33
			}
		})(_101_v33))), _102_v34)
		_ = _rhs9
		var _rhs10 _dafny.Int = (_99_v31).Select((Companion_Default___.SafeIndex(_97_cf86, _dafny.IntOfUint32((_99_v31).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs10
		var _rhs11 _dafny.Int = (_98_cf85).Plus(_97_cf86)
		_ = _rhs11
		var _rhs12 bool = !(_95_cf88) || (_95_cf88)
		_ = _rhs12
		_102_v34 = _rhs9
		_98_cf85 = _rhs10
		_97_cf86 = _rhs11
		_53_v0 = _rhs12
		var _105_v35 *C8
		_ = _105_v35
		var _nw8 *C8 = New_C8_()
		_ = _nw8
		_nw8.Ctor__()
		_105_v35 = _nw8
		(_60_globalState).F7 = _98_cf85
	} else {
		var _106___mcc_h4 _dafny.Array = _source3.Get_().(D20_DC54).Cf84
		_ = _106___mcc_h4
		var _107_cf84 _dafny.Array = _106___mcc_h4
		_ = _107_cf84
		var _108_v36 _dafny.Int
		_ = _108_v36
		var _109_v37 _dafny.MultiSet
		_ = _109_v37
		var _110_v38 bool
		_ = _110_v38
		var _out0 _dafny.Int
		_ = _out0
		var _out1 _dafny.MultiSet
		_ = _out1
		var _out2 bool
		_ = _out2
		_out0, _out1, _out2 = (_88_v28).M9(_54_v1, _53_v0, !(!(_53_v0)), _60_globalState)
		_108_v36 = _out0
		_109_v37 = _out1
		_110_v38 = _out2
		_110_v38 = _53_v0
		_53_v0 = _53_v0
		_54_v1 = (_dafny.Zero).Minus((_dafny.IntOfUint32((_58_v5).Cardinality())).Minus((func() _dafny.Int {
			if false {
				return _54_v1
			}
			return _108_v36
		})()))
	}
	var _111_v39 D11
	_ = _111_v39
	_111_v39 = Companion_D11_.Create_DC29_()
	var _112_v40 D11
	_ = _112_v40
	_112_v40 = Companion_D11_.Create_DC31_(_111_v39)
	var _113_v41 D11
	_ = _113_v41
	_113_v41 = Companion_D11_.Create_DC31_(_111_v39)
	var _114_v42 _dafny.Sequence
	_ = _114_v42
	_114_v42 = _dafny.SeqOf(_113_v41, _113_v41)
	var _115_v43 D19
	_ = _115_v43
	_115_v43 = Companion_D19_.Create_DC51_(_114_v42)
	_115_v43 = _115_v43
	var _116_v44 _dafny.Sequence
	_ = _116_v44
	_116_v44 = _dafny.SeqOf((func() _dafny.Int {
		if (_55_v2).Contains(_53_v0) {
			return (_55_v2).Get(_53_v0).(_dafny.Int)
		}
		return _54_v1
	})())
	(_60_globalState).F15 = _dafny.IntOfUint32((_116_v44).Cardinality())
	(_60_globalState).F15 = (_dafny.IntOfInt64(253)).Plus(_54_v1)
	var _117_v45 _dafny.Array
	_ = _117_v45
	var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
	_ = _nw9
	_117_v45 = _nw9
	_117_v45 = _117_v45
	_53_v0 = !(_53_v0)
	var _118_v46 _dafny.MultiSet
	_ = _118_v46
	_118_v46 = _dafny.MultiSetOf(_56_v3, _56_v3)
	var _119_v47 D5
	_ = _119_v47
	_119_v47 = Companion_D5_.Create_DC15_(Companion_D5_.Create_DC13_(_118_v46))
	var _120_v48 _dafny.Sequence
	_ = _120_v48
	_120_v48 = _dafny.SeqOf(_119_v47, _119_v47, _119_v47)
	_120_v48 = _120_v48
	_53_v0 = _53_v0
	var _121_v49 _dafny.Array
	_ = _121_v49
	var _nwElement0_0 _dafny.Sequence = _58_v5
	_ = _nwElement0_0
	var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(10))
	_ = _nw10
	(_nw10).ArraySet1(_nwElement0_0, 0)
	(_nw10).ArraySet1(_dafny.SeqOf(_53_v0), 1)
	(_nw10).ArraySet1(_58_v5, 2)
	(_nw10).ArraySet1(_dafny.SeqOf(_53_v0), 3)
	(_nw10).ArraySet1(_58_v5, 4)
	(_nw10).ArraySet1(_58_v5, 5)
	(_nw10).ArraySet1(_58_v5, 6)
	(_nw10).ArraySet1(_58_v5, 7)
	(_nw10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_58_v5, _58_v5), 8)
	(_nw10).ArraySet1(_dafny.Companion_Sequence_.Update(_58_v5, (Companion_Default___.SafeIndex(_54_v1, _dafny.IntOfUint32((_58_v5).Cardinality()))).Uint32(), _53_v0), 9)
	_121_v49 = _nw10
	var _122_v50 _dafny.Sequence
	_ = _122_v50
	_122_v50 = _dafny.UnicodeSeqOfUtf8Bytes("ipyktu")
	var _123_v51 *C12
	_ = _123_v51
	var _nw11 *C12 = New_C12_()
	_ = _nw11
	_nw11.Ctor__(_dafny.IntOfUint32((_122_v50).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("xofwgsc"))
	_123_v51 = _nw11
	var _124_v52 _dafny.Array
	_ = _124_v52
	var _nwElement0_1 *C12 = _123_v51
	_ = _nwElement0_1
	var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(4))
	_ = _nw12
	(_nw12).ArraySet1(_nwElement0_1, 0)
	(_nw12).ArraySet1(_123_v51, 1)
	(_nw12).ArraySet1(_123_v51, 2)
	(_nw12).ArraySet1(_123_v51, 3)
	_124_v52 = _nw12
	var _125_v53 _dafny.Sequence
	_ = _125_v53
	_125_v53 = _dafny.SeqOf(_124_v52, _124_v52, _124_v52, _124_v52)
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_121_v49), 0))
	_ = _index4
	(_121_v49).ArraySet1(_dafny.SeqOf(false, (_61_v7).Fm8(_dafny.IntOfUint32((_125_v53).Cardinality()), _53_v0, (_123_v51).F22(), _60_globalState)), (_index4).Int())
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_121_v49), 0))
	_ = _index5
	(_121_v49).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_58_v5, Companion_Default___.Fm27(_60_globalState)), (_index5).Int())
	var _126_v54 T0
	_ = _126_v54
	var _nw13 *C11 = New_C11_()
	_ = _nw13
	_nw13.Ctor__()
	_126_v54 = _nw13
	var _127_v55 _dafny.Sequence
	_ = _127_v55
	_127_v55 = _dafny.SeqOf(_126_v54)
	var _128_v56 _dafny.Map
	_ = _128_v56
	_128_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v0, _127_v55)
	if (_58_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
		if (_128_v56).Contains(_53_v0) {
			return (_128_v56).Get(_53_v0).(_dafny.Sequence)
		}
		return _127_v55
	})()).Cardinality()), _dafny.IntOfUint32((_58_v5).Cardinality()))).Uint32()).(bool) {
		(_60_globalState).F15 = (_123_v51).F22()
		var _129_v57 D1
		_ = _129_v57
		_129_v57 = Companion_D1_.Create_DC3_(_122_v50, _53_v0)
		var _source4 D1 = _129_v57
		_ = _source4
		if _source4.Is_DC3() {
			var _130___mcc_h5 _dafny.Sequence = _source4.Get_().(D1_DC3).Cf3
			_ = _130___mcc_h5
			var _131___mcc_h6 bool = _source4.Get_().(D1_DC3).Cf4
			_ = _131___mcc_h6
			var _132_cf4 bool = _131___mcc_h6
			_ = _132_cf4
			var _133_cf3 _dafny.Sequence = _130___mcc_h5
			_ = _133_cf3
			(_60_globalState).F13 = ((_123_v51).F22()).Plus(_dafny.IntOfInt64(604))
			var _134_v58 _dafny.Array
			_ = _134_v58
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_1
			var _nw14 _dafny.Array
			_ = _nw14
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw14 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_135_v51 *C12) func(_dafny.Int) _dafny.Int {
					return func(_136_i5 _dafny.Int) _dafny.Int {
						return (_136_i5).Minus((_135_v51).F22())
					}
				})(_123_v51)
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
			_134_v58 = _nw14
			var _137_v59 _dafny.Map
			_ = _137_v59
			_137_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v0, _134_v58)
			(_126_v54).M1(Companion_Default___.Fm19(_132_cf4, _53_v0, _60_globalState), (_137_v59).Merge(_137_v59), _60_globalState)
			(_60_globalState).F13 = (_123_v51).F22()
			var _138_v60 _dafny.CodePoint
			_ = _138_v60
			_138_v60 = _dafny.CodePoint('w')
			var _139_v61 _dafny.MultiSet
			_ = _139_v61
			_139_v61 = _dafny.MultiSetOf(_138_v60)
			_139_v61 = _139_v61
		} else if _source4.Is_DC4() {
			var _140___mcc_h7 _dafny.MultiSet = _source4.Get_().(D1_DC4).Cf5
			_ = _140___mcc_h7
			var _141_cf5 _dafny.MultiSet = _140___mcc_h7
			_ = _141_cf5
			var _142_v62 _dafny.Map
			_ = _142_v62
			_142_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v54, (!(true)) && (_53_v0))
			_53_v0 = (func() bool {
				if (_142_v62).Contains((func() T0 {
					if _53_v0 {
						return _126_v54
					}
					return _126_v54
				})()) {
					return (_142_v62).Get((func() T0 {
						if _53_v0 {
							return _126_v54
						}
						return _126_v54
					})()).(bool)
				}
				return _53_v0
			})()
			_54_v1 = Companion_Default___.SafeModuloInt(_54_v1, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _53_v0 {
					return _dafny.UnicodeSeqOfUtf8Bytes("cbcyxwik")
				}
				return _122_v50
			})()).Cardinality()))
			var _143_v63 _dafny.CodePoint
			_ = _143_v63
			_143_v63 = _dafny.CodePoint('r')
			var _144_v64 _dafny.Array
			_ = _144_v64
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_2
			var _nw15 _dafny.Array
			_ = _nw15
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw15 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) bool = (func(_145_v0 bool) func(_dafny.Int) bool {
					return func(_146_i6 _dafny.Int) bool {
						return _145_v0
					}
				})(_53_v0)
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
			_144_v64 = _nw15
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_144_v64), 0))
			_ = _index6
			(_144_v64).ArraySet1(_53_v0, (_index6).Int())
			var _147_v65 _dafny.Array
			_ = _147_v65
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_3
			var _nw16 _dafny.Array
			_ = _nw16
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw16 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = (func(_148_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_149_i7 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_149_i7, _148_v1)
					}
				})(_54_v1)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw16 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw16).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw16).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_147_v65 = _nw16
			var _150_v66 D2
			_ = _150_v66
			_150_v66 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC5_(_147_v65))
			var _151_v67 D2
			_ = _151_v67
			_151_v67 = Companion_D2_.Create_DC5_(_147_v65)
			var _152_v68 D2
			_ = _152_v68
			_152_v68 = Companion_D2_.Create_DC8_(_151_v67)
			var _pat_let_tv0 = _152_v68
			_ = _pat_let_tv0
			var _153_v69 *C3
			_ = _153_v69
			var _nw17 *C3 = New_C3_()
			_ = _nw17
			_nw17.Ctor__(func(_pat_let0_0 D2) D2 {
				return func(_154_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let1_0 D2) D2 {
						return func(_155_dt__update_hcf11_h0 D2) D2 {
							return Companion_D2_.Create_DC8_(_155_dt__update_hcf11_h0)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_150_v66), _53_v0)
			_153_v69 = _nw17
			var _156_v70 _dafny.Map
			_ = _156_v70
			_156_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_153_v69, (_123_v51).F22())
			var _157_v71 _dafny.Set
			_ = _157_v71
			_157_v71 = _dafny.SetOf(_53_v0, _53_v0)
			var _158_v72 _dafny.Sequence
			_ = _158_v72
			_158_v72 = _dafny.SeqOf(_157_v71)
			var _159_v73 _dafny.Array
			_ = _159_v73
			var _nwElement0_2 _dafny.Int = ((_156_v70).Cardinality()).Plus((_dafny.Zero).Minus((_153_v69).Fm2((_123_v51).Fm3(_153_v69.F25, _60_globalState), _60_globalState)))
			_ = _nwElement0_2
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(22))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_2, 0)
			(_nw18).ArraySet1((_123_v51).F22(), 1)
			(_nw18).ArraySet1(_dafny.IntOfInt64(801), 2)
			(_nw18).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jqh"), _dafny.UnicodeSeqOfUtf8Bytes("usr"))).Cardinality()), 3)
			(_nw18).ArraySet1((_123_v51).F22(), 4)
			(_nw18).ArraySet1((_123_v51).F22(), 5)
			(_nw18).ArraySet1(_dafny.IntOfInt64(152), 6)
			(_nw18).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_158_v72).Cardinality())).Times(_54_v1)), 7)
			(_nw18).ArraySet1((_123_v51).F22(), 8)
			(_nw18).ArraySet1((_dafny.IntOfUint32((_122_v50).Cardinality())).Times(_dafny.IntOfInt64(-562)), 9)
			(_nw18).ArraySet1((_dafny.IntOfUint32(((_123_v51).F23()).Cardinality())).Times(_dafny.IntOfInt64(-945)), 10)
			(_nw18).ArraySet1((func() _dafny.Int {
				if _53_v0 {
					return (_123_v51).F22()
				}
				return _54_v1
			})(), 11)
			(_nw18).ArraySet1((_116_v44).Select((Companion_Default___.SafeIndex((_55_v2).Cardinality(), _dafny.IntOfUint32((_116_v44).Cardinality()))).Uint32()).(_dafny.Int), 12)
			(_nw18).ArraySet1((_123_v51).F22(), 13)
			(_nw18).ArraySet1(_dafny.IntOfInt64(798), 14)
			(_nw18).ArraySet1(_dafny.IntOfInt64(-850), 15)
			(_nw18).ArraySet1((_123_v51).F22(), 16)
			(_nw18).ArraySet1(((_123_v51).F22()).Times((_123_v51).F22()), 17)
			(_nw18).ArraySet1(_dafny.IntOfInt64(495), 18)
			(_nw18).ArraySet1((_116_v44).Select((Companion_Default___.SafeIndex((_123_v51).F22(), _dafny.IntOfUint32((_116_v44).Cardinality()))).Uint32()).(_dafny.Int), 19)
			(_nw18).ArraySet1(_54_v1, 20)
			(_nw18).ArraySet1(_dafny.IntOfInt64(-503), 21)
			_159_v73 = _nw18
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_147_v65), 0))
			_ = _index7
			(_147_v65).ArraySet1((_123_v51).F22(), (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_144_v64), 0))
			_ = _index8
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_147_v65), 0))
			_ = _index9
			var _rhs13 _dafny.Sequence = _122_v50
			_ = _rhs13
			var _rhs14 _dafny.CodePoint = _143_v63
			_ = _rhs14
			var _rhs15 bool = true
			_ = _rhs15
			var _rhs16 bool = _153_v69.F25
			_ = _rhs16
			var _rhs17 _dafny.Int = _54_v1
			_ = _rhs17
			var _lhs7 _dafny.Array = _144_v64
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_144_v64), 0))
			_ = _lhs8
			var _lhs9 _dafny.Array = _147_v65
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_147_v65), 0))
			_ = _lhs10
			_122_v50 = _rhs13
			_143_v63 = _rhs14
			_53_v0 = _rhs15
			(_lhs7).ArraySet1(_rhs16, (_lhs8).Int())
			(_lhs9).ArraySet1(_rhs17, (_lhs10).Int())
			_57_v4 = (_57_v4).Update(true, (_144_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_144_v64), 0))).Int()).(bool))
		} else {
			var _160___mcc_h8 _dafny.Int = _source4.Get_().(D1_DC2).Cf2
			_ = _160___mcc_h8
			var _161_cf2 _dafny.Int = _160___mcc_h8
			_ = _161_cf2
			(_60_globalState).F15 = _161_cf2
			var _162_v74 *C13
			_ = _162_v74
			var _nw19 *C13 = New_C13_()
			_ = _nw19
			_nw19.Ctor__((_123_v51).F22(), !((_88_v28).Fm16(_dafny.IntOfInt64(366), _60_globalState)))
			_162_v74 = _nw19
			var _163_v75 _dafny.MultiSet
			_ = _163_v75
			_163_v75 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_123_v51).F23()).Cardinality()), _54_v1)
			var _rhs18 *C13 = _162_v74
			_ = _rhs18
			var _rhs19 bool = false
			_ = _rhs19
			var _rhs20 _dafny.Int = Companion_Default___.SafeModuloInt((_163_v75).Cardinality(), (_162_v74).F20())
			_ = _rhs20
			var _lhs11 *GlobalState = _60_globalState
			_ = _lhs11
			_162_v74 = _rhs18
			_53_v0 = _rhs19
			_lhs11.F13 = _rhs20
			(_123_v51).M2((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_116_v44, _116_v44)).Cardinality())), _53_v0, _dafny.IntOfInt64(-853), _60_globalState)
			(_162_v74).F21 = false
		}
		if _53_v0 {
			var _164_v76 _dafny.Array
			_ = _164_v76
			var _nwElement0_3 _dafny.Int = _dafny.IntOfInt64(259)
			_ = _nwElement0_3
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(3))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_3, 0)
			(_nw20).ArraySet1(_dafny.IntOfUint32(((_123_v51).F23()).Cardinality()), 1)
			(_nw20).ArraySet1((_123_v51).F22(), 2)
			_164_v76 = _nw20
			_164_v76 = _164_v76
			_55_v2 = (_55_v2).Merge(_55_v2)
			var _165_v77 _dafny.Map
			_ = _165_v77
			_165_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v0, _164_v76)
			var _166_v78 D23
			_ = _166_v78
			_166_v78 = Companion_D23_.Create_DC63_(_165_v77)
			(_126_v54).M1((_123_v51).F23(), (_165_v77).Merge((_166_v78).Dtor_cf106()), _60_globalState)
			var _167_v79 D15
			_ = _167_v79
			_167_v79 = Companion_D15_.Create_DC39_(_57_v4, !(_53_v0), (_dafny.MultiSetOf(_54_v1)).Cardinality())
			var _168_v80 _dafny.Map
			_ = _168_v80
			_168_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(985), _167_v79)
			_168_v80 = _168_v80
			_164_v76 = _164_v76
		} else {
			var _169_v81 _dafny.CodePoint
			_ = _169_v81
			_169_v81 = _dafny.CodePoint('p')
			_169_v81 = _169_v81
			var _170_v82 _dafny.Set
			_ = _170_v82
			_170_v82 = _dafny.SetOf(_53_v0)
			var _171_v83 D12
			_ = _171_v83
			_171_v83 = Companion_D12_.Create_DC32_(_170_v82)
			_171_v83 = Companion_D12_.Create_DC32_((_170_v82).Union(_170_v82))
			var _172_v84 _dafny.Array
			_ = _172_v84
			var _nw21 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw21
			_172_v84 = _nw21
			var _173_v85 D0
			_ = _173_v85
			_173_v85 = Companion_D0_.Create_DC0_(_172_v84)
			var _174_v86 _dafny.Sequence
			_ = _174_v86
			_174_v86 = _dafny.SeqOf(_173_v85)
			(_60_globalState).F4 = (_dafny.IntOfUint32(((func() _dafny.Sequence {
				if _53_v0 {
					return _174_v86
				}
				return _dafny.SeqOf(_173_v85)
			})()).Cardinality())).Times(_54_v1)
			(_60_globalState).F4 = _dafny.IntOfInt64(857)
			var _175_v87 T1
			_ = _175_v87
			var _nw22 *C2 = New_C2_()
			_ = _nw22
			_nw22.Ctor__()
			_175_v87 = _nw22
			var _176_v88 _dafny.Sequence
			_ = _176_v88
			_176_v88 = _dafny.SeqOf(_175_v87, _175_v87)
			(_123_v51).M2(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_176_v88, _176_v88)).Cardinality()), _53_v0, (_123_v51).F22(), _60_globalState)
		}
		(_60_globalState).F13 = (_dafny.Zero).Minus(_54_v1)
		var _177_v89 _dafny.Array
		_ = _177_v89
		var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw23
		_177_v89 = _nw23
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_177_v89), 0))
		_ = _index10
		(_177_v89).ArraySet1((func() _dafny.Int {
			if (_55_v2).Contains(!(_53_v0)) {
				return (_55_v2).Get(!(_53_v0)).(_dafny.Int)
			}
			return (_123_v51).F22()
		})(), (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_177_v89), 0))
		_ = _index11
		(_177_v89).ArraySet1(_dafny.IntOfUint32(((_123_v51).F23()).Cardinality()), (_index11).Int())
	} else {
		var _178_v90 _dafny.Set
		_ = _178_v90
		_178_v90 = _dafny.SetOf(_53_v0)
		if (Companion_Default___.Fm37(_60_globalState)).IsSubsetOf(_178_v90) {
			var _179_v91 T1
			_ = _179_v91
			var _nw24 *C9 = New_C9_()
			_ = _nw24
			_nw24.Ctor__()
			_179_v91 = _nw24
			var _180_v93 _dafny.Map
			_ = _180_v93
			_180_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v91, _dafny.MultiSetOf(_56_v3, func() _dafny.Set {
				var _coll19 = _dafny.NewBuilder()
				_ = _coll19
				for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-884), _dafny.IntOfInt64(510))); ; {
					_compr_19, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _181_v92 _dafny.Int
					_181_v92 = interface{}(_compr_19).(_dafny.Int)
					if ((_dafny.IntOfInt64(-884)).Cmp(_181_v92) <= 0) && ((_181_v92).Cmp(_dafny.IntOfInt64(510)) < 0) {
						_coll19.Add((_181_v92).Plus(_dafny.IntOfInt64(-656)))
					}
				}
				return _coll19.ToSet()
			}()))
			_180_v93 = (_180_v93).Update(_179_v91, _dafny.MultiSetOf(_56_v3))
			var _182_v94 *C7
			_ = _182_v94
			var _nw25 *C7 = New_C7_()
			_ = _nw25
			_nw25.Ctor__()
			_182_v94 = _nw25
			var _183_v95 _dafny.Map
			_ = _183_v95
			_183_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v94, _54_v1)
			(_60_globalState).F15 = (func() _dafny.Int {
				if (_183_v95).Contains(_182_v94) {
					return (_183_v95).Get(_182_v94).(_dafny.Int)
				}
				return (_123_v51).F22()
			})()
			var _184_v96 _dafny.CodePoint
			_ = _184_v96
			_184_v96 = _dafny.CodePoint('c')
			var _185_v97 _dafny.Map
			_ = _185_v97
			_185_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_123_v51).F22(), _184_v96)
			var _186_v98 D19
			_ = _186_v98
			_186_v98 = Companion_D19_.Create_DC52_(((_185_v97).Merge(_185_v97)).Cardinality(), _53_v0, _53_v0)
			var _187_v99 _dafny.Map
			_ = _187_v99
			_187_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v1, _53_v0)
			_186_v98 = Companion_Default___.Fm60(_187_v99, (_dafny.Zero).Minus(_dafny.IntOfUint32(((_123_v51).F23()).Cardinality())), _60_globalState)
			(_60_globalState).F15 = ((_123_v51).F22()).Plus((_123_v51).F22())
			var _188_v100 _dafny.Sequence
			_ = _188_v100
			_188_v100 = _dafny.SeqOf((_57_v4).Merge(_57_v4), _57_v4, _57_v4)
			var _189_v101 _dafny.Map
			_ = _189_v101
			_189_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(896), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v0, _53_v0))
			var _rhs21 _dafny.Int = (_123_v51).F22()
			_ = _rhs21
			var _rhs22 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_188_v100, _dafny.Companion_Sequence_.Concatenate(_188_v100, _dafny.SeqOf((func() _dafny.Map {
				if (_189_v101).Contains((_123_v51).F22()) {
					return (_189_v101).Get((_123_v51).F22()).(_dafny.Map)
				}
				return _57_v4
			})(), _57_v4)))
			_ = _rhs22
			var _rhs23 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(466), _54_v1), _dafny.IntOfInt64(-194))
			_ = _rhs23
			var _lhs12 *GlobalState = _60_globalState
			_ = _lhs12
			var _lhs13 *GlobalState = _60_globalState
			_ = _lhs13
			_lhs12.F4 = _rhs21
			_188_v100 = _rhs22
			_lhs13.F7 = _rhs23
		} else {
			var _190_v102 _dafny.MultiSet
			_ = _190_v102
			_190_v102 = _dafny.MultiSetOf(_53_v0, _53_v0, _53_v0)
			var _191_v103 _dafny.Map
			_ = _191_v103
			_191_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _190_v102)
			var _192_v104 _dafny.Array
			_ = _192_v104
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw26
			_192_v104 = _nw26
			var _193_v105 D2
			_ = _193_v105
			_193_v105 = Companion_D2_.Create_DC5_(_192_v104)
			(_61_v7).M6(_54_v1, ((func() _dafny.MultiSet {
				if (_191_v103).Contains(true) {
					return (_191_v103).Get(true).(_dafny.MultiSet)
				}
				return _190_v102
			})()).Update(!(_53_v0), Companion_Default___.Abs((Companion_Default___.Fm32((_123_v51).F22(), (_123_v51).F22(), _dafny.SetOf(_53_v0, (func() bool {
				if (_57_v4).Contains(true) {
					return (_57_v4).Get(true).(bool)
				}
				return _53_v0
			})(), _53_v0), _60_globalState)).Cardinality())), _53_v0, _193_v105, _60_globalState)
			var _194_v106 _dafny.MultiSet
			_ = _194_v106
			_194_v106 = _dafny.MultiSetOf((_123_v51).F22())
			var _195_v107 _dafny.Map
			_ = _195_v107
			_195_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_116_v44).Cardinality()), _194_v106)
			(_123_v51).M2((_195_v107).Cardinality(), _53_v0, _54_v1, _60_globalState)
			var _196_v108 D7
			_ = _196_v108
			_196_v108 = Companion_D7_.Create_DC21_(_dafny.IntOfUint32(((_123_v51).F23()).Cardinality()), _dafny.IntOfInt64(-681), _53_v0, _53_v0, _53_v0)
			(_123_v51).M2((_123_v51).F22(), _53_v0, (_123_v51).Fm2((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_196_v108).Dtor_cf32()), (_123_v51).F22())).Cardinality(), _60_globalState), _60_globalState)
			var _197_v109 _dafny.Map
			_ = _197_v109
			_197_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v0, _57_v4)
			var _198_v110 D14
			_ = _198_v110
			_198_v110 = Companion_D14_.Create_DC36_(_dafny.SeqOf(_53_v0, _53_v0, _53_v0, (_123_v51).Fm1(_197_v109, _dafny.IntOfInt64(79), _60_globalState)))
			var _199_v111 _dafny.Set
			_ = _199_v111
			_199_v111 = _dafny.SetOf(_116_v44, _116_v44, _116_v44, _116_v44)
			var _rhs24 _dafny.Int = Companion_Default___.SafeModuloInt(_54_v1, (_123_v51).F22())
			_ = _rhs24
			var _rhs25 D14 = _198_v110
			_ = _rhs25
			var _rhs26 bool = (_199_v111).IsProperSubsetOf(_199_v111)
			_ = _rhs26
			var _rhs27 bool = ((_55_v2).Equals(_55_v2)) == ((_59_v6).IsProperSubsetOf(_59_v6))
			_ = _rhs27
			var _lhs14 *GlobalState = _60_globalState
			_ = _lhs14
			_lhs14.F7 = _rhs24
			_198_v110 = _rhs25
			_53_v0 = _rhs26
			_53_v0 = _rhs27
			(_60_globalState).F4 = (_dafny.Zero).Minus((_123_v51).F22())
		}
		(_60_globalState).F15 = ((_123_v51).F22()).Plus((_178_v90).Cardinality())
		_53_v0 = (_178_v90).IsProperSubsetOf(_178_v90)
		var _200_v112 _dafny.Int
		_ = _200_v112
		var _201_v113 _dafny.MultiSet
		_ = _201_v113
		var _202_v114 bool
		_ = _202_v114
		var _out3 _dafny.Int
		_ = _out3
		var _out4 _dafny.MultiSet
		_ = _out4
		var _out5 bool
		_ = _out5
		_out3, _out4, _out5 = (_88_v28).M9(Companion_Default___.SafeDivisionInt((_123_v51).F22(), (_123_v51).F22()), false, _53_v0, _60_globalState)
		_200_v112 = _out3
		_201_v113 = _out4
		_202_v114 = _out5
		_202_v114 = _dafny.Companion_Sequence_.IsProperPrefixOf(_116_v44, _dafny.Companion_Sequence_.Concatenate(_116_v44, _116_v44))
	}
	var _203_v115 *C5
	_ = _203_v115
	var _nw27 *C5 = New_C5_()
	_ = _nw27
	_nw27.Ctor__()
	_203_v115 = _nw27
	var _204_v116 _dafny.Set
	_ = _204_v116
	_204_v116 = _dafny.SetOf(_203_v115, _203_v115)
	var _205_v117 *C1
	_ = _205_v117
	var _nw28 *C1 = New_C1_()
	_ = _nw28
	_nw28.Ctor__()
	_205_v117 = _nw28
	var _206_v118 *C8
	_ = _206_v118
	var _nw29 *C8 = New_C8_()
	_ = _nw29
	_nw29.Ctor__()
	_206_v118 = _nw29
	var _207_v119 _dafny.Map
	_ = _207_v119
	_207_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_205_v117, _206_v118)
	var _208_v120 _dafny.Array
	_ = _208_v120
	var _nwElement0_4 _dafny.Int = (_dafny.Zero).Minus((_123_v51).F22())
	_ = _nwElement0_4
	var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(13))
	_ = _nw30
	(_nw30).ArraySet1(_nwElement0_4, 0)
	(_nw30).ArraySet1((_123_v51).F22(), 1)
	(_nw30).ArraySet1((_54_v1).Minus(_dafny.IntOfUint32((Companion_Default___.Fm5(_60_globalState)).Cardinality())), 2)
	(_nw30).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_121_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_121_v49), 0))).Int()).(_dafny.Sequence), (_121_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(986), _dafny.ArrayLen((_121_v49), 0))).Int()).(_dafny.Sequence))).Cardinality()), 3)
	(_nw30).ArraySet1(_54_v1, 4)
	(_nw30).ArraySet1(_54_v1, 5)
	(_nw30).ArraySet1(((_123_v51).F22()).Minus((_123_v51).F22()), 6)
	(_nw30).ArraySet1(_54_v1, 7)
	(_nw30).ArraySet1(Companion_Default___.SafeDivisionInt(_54_v1, _54_v1), 8)
	(_nw30).ArraySet1(_54_v1, 9)
	(_nw30).ArraySet1((_204_v116).Cardinality(), 10)
	(_nw30).ArraySet1(_54_v1, 11)
	(_nw30).ArraySet1(Companion_Default___.SafeDivisionInt(((_207_v119).Update(_205_v117, _206_v118)).Cardinality(), _dafny.IntOfInt64(382)), 12)
	_208_v120 = _nw30
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_208_v120), 0))
	_ = _index12
	(_208_v120).ArraySet1(_54_v1, (_index12).Int())
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(52), _dafny.ArrayLen((_208_v120), 0))
	_ = _index13
	(_208_v120).ArraySet1(Companion_Default___.SafeModuloInt((_123_v51).F22(), (_123_v51).F22()), (_index13).Int())
	var _209_v121 _dafny.CodePoint
	_ = _209_v121
	_209_v121 = _dafny.CodePoint('i')
	var _210_v122 D6
	_ = _210_v122
	_210_v122 = Companion_D6_.Create_DC16_(_116_v44)
	var _211_v123 _dafny.Map
	_ = _211_v123
	_211_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm39((_116_v44).Select((Companion_Default___.SafeIndex(_54_v1, _dafny.IntOfUint32((_116_v44).Cardinality()))).Uint32()).(_dafny.Int), _209_v121, (_210_v122).Dtor_cf23(), _53_v0, _60_globalState), _dafny.Companion_Sequence_.Update((_123_v51).F23(), (Companion_Default___.SafeIndex(_54_v1, _dafny.IntOfUint32(((_123_v51).F23()).Cardinality()))).Uint32(), _209_v121))
	_211_v123 = (_211_v123).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(278))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}((func(_212_v51 *C12) func(_dafny.Int) _dafny.Int {
		return func(_213_i8 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32(((_212_v51).F23()).Cardinality())
		}
	})(_123_v51))), _dafny.Companion_Sequence_.Concatenate((_123_v51).F23(), (_123_v51).F23()))
	var _214_v124 _dafny.Map
	_ = _214_v124
	_214_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v0, _122_v50)
	var _rhs28 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if (_214_v124).Contains(_53_v0) {
			return (_214_v124).Get(_53_v0).(_dafny.Sequence)
		}
		return _122_v50
	})(), (_123_v51).F23())
	_ = _rhs28
	var _rhs29 _dafny.Int = (_123_v51).F22()
	_ = _rhs29
	var _lhs15 *GlobalState = _60_globalState
	_ = _lhs15
	_122_v50 = _rhs28
	_lhs15.F15 = _rhs29
	_55_v2 = ((_55_v2).Merge(_55_v2)).Merge(_55_v2)
	_dafny.Print(_53_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_54_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_55_v2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(215))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v3).Equals(_dafny.SetOf(_dafny.IntOfInt64(215))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_57_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_58_v5, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_v6).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_60_globalState).F0(), _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_60_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_globalState).F6()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(215))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_60_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_globalState).F8()).Equals(_dafny.SetOf(_dafny.IntOfInt64(215))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_globalState).F11()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false).UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_60_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_60_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_globalState).F16()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_89_v29).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v30).Dtor_cf85())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v30).Dtor_cf86())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v30).Dtor_cf87())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_90_v30).Dtor_cf88())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_114_v42, _dafny.SeqOf(Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_()), Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_()))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_115_v43).Dtor_cf80(), _dafny.SeqOf(Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_()), Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_()))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_116_v44, _dafny.SeqOf(_dafny.IntOfInt64(215))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_v46).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(215)), _dafny.SetOf(_dafny.IntOfInt64(215)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_119_v47).Dtor_cf22()).Dtor_cf17()).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(215)), _dafny.SetOf(_dafny.IntOfInt64(215)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_120_v48, _dafny.SeqOf(Companion_D5_.Create_DC15_(Companion_D5_.Create_DC13_(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(215)), _dafny.SetOf(_dafny.IntOfInt64(215))))), Companion_D5_.Create_DC15_(Companion_D5_.Create_DC13_(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(215)), _dafny.SetOf(_dafny.IntOfInt64(215))))), Companion_D5_.Create_DC15_(Companion_D5_.Create_DC13_(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(215)), _dafny.SetOf(_dafny.IntOfInt64(215))))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v49).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_122_v50.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_v51).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_v51).F23().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v52).ArrayGet1((_dafny.Zero).Int()).(*C12)).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v52).ArrayGet1((_dafny.Zero).Int()).(*C12)).F23().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v52).ArrayGet1((_dafny.One).Int()).(*C12)).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v52).ArrayGet1((_dafny.One).Int()).(*C12)).F23().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v52).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C12)).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v52).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C12)).F23().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v52).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C12)).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v52).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C12)).F23().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_125_v53).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_127_v55).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v56).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_204_v116).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_207_v119).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v120).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_209_v121)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_210_v122).Dtor_cf23(), _dafny.SeqOf(_dafny.IntOfInt64(215))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_211_v123).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(669), _dafny.IntOfInt64(2), _dafny.One, _dafny.IntOfInt64(412), _dafny.One, _dafny.IntOfInt64(193), _dafny.One, _dafny.IntOfInt64(2)), _dafny.UnicodeSeqOfUtf8Bytes("xofwgic")).UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7), _dafny.IntOfInt64(7)), _dafny.UnicodeSeqOfUtf8Bytes("xofwgscxofwgsc"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v124).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("ipyktu"))))
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
	return Companion_D0_.Create_DC1_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf1() _dafny.Array {
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
	Cf3 _dafny.Sequence
	Cf4 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Sequence, Cf4 bool) D1 {
	return D1{D1_DC3{Cf3, Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf5 _dafny.MultiSet
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 _dafny.MultiSet) D1 {
	return D1{D1_DC4{Cf5}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC2 struct {
	Cf2 _dafny.Int
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.EmptySeq, false)
}

func (_this D1) Dtor_cf3() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() _dafny.MultiSet {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + data.Cf3.VerbatimString(true) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf3.Equals(data2.Cf3) && data1.Cf4 == data2.Cf4
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf5.Equals(data2.Cf5)
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0
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
	Cf7 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf7 _dafny.Int) D2 {
	return D2{D2_DC6{Cf7}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf8  bool
	Cf9  _dafny.MultiSet
	Cf10 _dafny.Array
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf8 bool, Cf9 _dafny.MultiSet, Cf10 _dafny.Array) D2 {
	return D2{D2_DC7{Cf8, Cf9, Cf10}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC5 struct {
	Cf6 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf6 _dafny.Array) D2 {
	return D2{D2_DC5{Cf6}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
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
	return Companion_D2_.Create_DC6_(_dafny.Zero)
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf7
}

func (_this D2) Dtor_cf8() bool {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) Dtor_cf9() _dafny.MultiSet {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) Dtor_cf10() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf10
}

func (_this D2) Dtor_cf6() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf11() D2 {
	return _this.Get_().(D2_DC8).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf6) + ")"
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
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9.Equals(data2.Cf9) && data1.Cf10 == data2.Cf10
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf6 == data2.Cf6
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
	Cf13 _dafny.Int
	Cf14 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf13 _dafny.Int, Cf14 _dafny.Int) D3 {
	return D3{D3_DC10{Cf13, Cf14}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf12 _dafny.CodePoint
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf12 _dafny.CodePoint) D3 {
	return D3{D3_DC9{Cf12}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.Zero, _dafny.Zero)
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf14
}

func (_this D3) Dtor_cf12() _dafny.CodePoint {
	return _this.Get_().(D3_DC9).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf12 == data2.Cf12
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
	Cf16 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf16 _dafny.Int) D4 {
	return D4{D4_DC12{Cf16}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC11 struct {
	Cf15 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf15 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf15}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(_dafny.Zero)
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf16
}

func (_this D4) Dtor_cf15() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf18 _dafny.CodePoint
	Cf19 bool
	Cf20 _dafny.Int
	Cf21 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf18 _dafny.CodePoint, Cf19 bool, Cf20 _dafny.Int, Cf21 bool) D5 {
	return D5{D5_DC14{Cf18, Cf19, Cf20, Cf21}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf17 _dafny.MultiSet
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf17 _dafny.MultiSet) D5 {
	return D5{D5_DC13{Cf17}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC15 struct {
	Cf22 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf22 D5) D5 {
	return D5{D5_DC15{Cf22}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(_dafny.CodePoint('D'), false, _dafny.Zero, false)
}

func (_this D5) Dtor_cf18() _dafny.CodePoint {
	return _this.Get_().(D5_DC14).Cf18
}

func (_this D5) Dtor_cf19() bool {
	return _this.Get_().(D5_DC14).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC14).Cf21
}

func (_this D5) Dtor_cf17() _dafny.MultiSet {
	return _this.Get_().(D5_DC13).Cf17
}

func (_this D5) Dtor_cf22() D5 {
	return _this.Get_().(D5_DC15).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf17.Equals(data2.Cf17)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf22.Equals(data2.Cf22)
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
	Cf24 bool
	Cf25 bool
	Cf26 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf24 bool, Cf25 bool, Cf26 bool) D6 {
	return D6{D6_DC17{Cf24, Cf25, Cf26}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf23 _dafny.Sequence
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf23 _dafny.Sequence) D6 {
	return D6{D6_DC16{Cf23}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC18 struct {
	Cf27 D6
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf27 D6) D6 {
	return D6{D6_DC18{Cf27}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(false, false, false)
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC17).Cf24
}

func (_this D6) Dtor_cf25() bool {
	return _this.Get_().(D6_DC17).Cf25
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC17).Cf26
}

func (_this D6) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf23
}

func (_this D6) Dtor_cf27() D6 {
	return _this.Get_().(D6_DC18).Cf27
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf27) + ")"
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
			return ok && data1.Cf24 == data2.Cf24 && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf23.Equals(data2.Cf23)
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
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

type D7_DC20 struct {
	Cf29 _dafny.Int
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf29 _dafny.Int) D7 {
	return D7{D7_DC20{Cf29}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC21 struct {
	Cf30 _dafny.Int
	Cf31 _dafny.Int
	Cf32 bool
	Cf33 bool
	Cf34 bool
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf30 _dafny.Int, Cf31 _dafny.Int, Cf32 bool, Cf33 bool, Cf34 bool) D7 {
	return D7{D7_DC21{Cf30, Cf31, Cf32, Cf33, Cf34}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC22 struct {
	Cf35 bool
	Cf36 _dafny.Int
	Cf37 bool
	Cf38 _dafny.Int
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf35 bool, Cf36 _dafny.Int, Cf37 bool, Cf38 _dafny.Int) D7 {
	return D7{D7_DC22{Cf35, Cf36, Cf37, Cf38}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC19 struct {
	Cf28 _dafny.Sequence
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf28 _dafny.Sequence) D7 {
	return D7{D7_DC19{Cf28}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC20_(_dafny.Zero)
}

func (_this D7) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D7_DC20).Cf29
}

func (_this D7) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D7_DC21).Cf30
}

func (_this D7) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D7_DC21).Cf31
}

func (_this D7) Dtor_cf32() bool {
	return _this.Get_().(D7_DC21).Cf32
}

func (_this D7) Dtor_cf33() bool {
	return _this.Get_().(D7_DC21).Cf33
}

func (_this D7) Dtor_cf34() bool {
	return _this.Get_().(D7_DC21).Cf34
}

func (_this D7) Dtor_cf35() bool {
	return _this.Get_().(D7_DC22).Cf35
}

func (_this D7) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D7_DC22).Cf36
}

func (_this D7) Dtor_cf37() bool {
	return _this.Get_().(D7_DC22).Cf37
}

func (_this D7) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D7_DC22).Cf38
}

func (_this D7) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D7_DC19).Cf28
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33 && data1.Cf34 == data2.Cf34
		}
	case D7_DC22:
		{
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
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

type D8_DC24 struct {
	Cf40 _dafny.Int
	Cf41 bool
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf40 _dafny.Int, Cf41 bool) D8 {
	return D8{D8_DC24{Cf40, Cf41}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC25 struct {
	Cf42 _dafny.Int
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf42 _dafny.Int) D8 {
	return D8{D8_DC25{Cf42}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC23 struct {
	Cf39 _dafny.Sequence
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf39 _dafny.Sequence) D8 {
	return D8{D8_DC23{Cf39}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC24_(_dafny.Zero, false)
}

func (_this D8) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf40
}

func (_this D8) Dtor_cf41() bool {
	return _this.Get_().(D8_DC24).Cf41
}

func (_this D8) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D8_DC25).Cf42
}

func (_this D8) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D8_DC23).Cf39
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41 == data2.Cf41
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf42.Cmp(data2.Cf42) == 0
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	Cf43 _dafny.Array
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf43 _dafny.Array) D9 {
	return D9{D9_DC26{Cf43}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D9) Dtor_cf43() _dafny.Array {
	return _this.Get_().(D9_DC26).Cf43
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf43) + ")"
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
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf43 == data2.Cf43
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

type D10_DC27 struct {
	Cf44 _dafny.Map
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf44 _dafny.Map) D10 {
	return D10{D10_DC27{Cf44}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D10) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D10_DC27).Cf44
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf44.Equals(data2.Cf44)
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

type D11_DC29 struct {
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_() D11 {
	return D11{D11_DC29{}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

type D11_DC30 struct {
	Cf46 bool
	Cf47 _dafny.Int
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf46 bool, Cf47 _dafny.Int) D11 {
	return D11{D11_DC30{Cf46, Cf47}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC28 struct {
	Cf45 _dafny.Set
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf45 _dafny.Set) D11 {
	return D11{D11_DC28{Cf45}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC31 struct {
	Cf48 D11
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf48 D11) D11 {
	return D11{D11_DC31{Cf48}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC29_()
}

func (_this D11) Dtor_cf46() bool {
	return _this.Get_().(D11_DC30).Cf46
}

func (_this D11) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D11_DC30).Cf47
}

func (_this D11) Dtor_cf45() _dafny.Set {
	return _this.Get_().(D11_DC28).Cf45
}

func (_this D11) Dtor_cf48() D11 {
	return _this.Get_().(D11_DC31).Cf48
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC29:
		{
			return "D11.DC29"
		}
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC29:
		{
			_, ok := other.Get_().(D11_DC29)
			return ok
		}
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0
		}
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf48.Equals(data2.Cf48)
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
	Cf50 bool
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf50 bool) D12 {
	return D12{D12_DC33{Cf50}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC32 struct {
	Cf49 _dafny.Set
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf49 _dafny.Set) D12 {
	return D12{D12_DC32{Cf49}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC33_(false)
}

func (_this D12) Dtor_cf50() bool {
	return _this.Get_().(D12_DC33).Cf50
}

func (_this D12) Dtor_cf49() _dafny.Set {
	return _this.Get_().(D12_DC32).Cf49
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf49) + ")"
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
			return ok && data1.Cf50 == data2.Cf50
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D13_DC34 struct {
	Cf51 _dafny.Map
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf51 _dafny.Map) D13 {
	return D13{D13_DC34{Cf51}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D13) Dtor_cf51() _dafny.Map {
	return _this.Get_().(D13_DC34).Cf51
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf51.Equals(data2.Cf51)
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

type D14_DC36 struct {
	Cf53 _dafny.Sequence
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf53 _dafny.Sequence) D14 {
	return D14{D14_DC36{Cf53}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

type D14_DC35 struct {
	Cf52 _dafny.Sequence
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf52 _dafny.Sequence) D14 {
	return D14{D14_DC35{Cf52}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

type D14_DC37 struct {
	Cf54 D14
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf54 D14) D14 {
	return D14{D14_DC37{Cf54}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC36_(_dafny.EmptySeq)
}

func (_this D14) Dtor_cf53() _dafny.Sequence {
	return _this.Get_().(D14_DC36).Cf53
}

func (_this D14) Dtor_cf52() _dafny.Sequence {
	return _this.Get_().(D14_DC35).Cf52
}

func (_this D14) Dtor_cf54() D14 {
	return _this.Get_().(D14_DC37).Cf54
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf53) + ")"
		}
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf52) + ")"
		}
	case D14_DC37:
		{
			return "D14.DC37" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf53.Equals(data2.Cf53)
		}
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf52.Equals(data2.Cf52)
		}
	case D14_DC37:
		{
			data2, ok := other.Get_().(D14_DC37)
			return ok && data1.Cf54.Equals(data2.Cf54)
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
	Cf56 _dafny.Map
	Cf57 bool
	Cf58 _dafny.Int
}

func (D15_DC39) isD15() {}

func (CompanionStruct_D15_) Create_DC39_(Cf56 _dafny.Map, Cf57 bool, Cf58 _dafny.Int) D15 {
	return D15{D15_DC39{Cf56, Cf57, Cf58}}
}

func (_this D15) Is_DC39() bool {
	_, ok := _this.Get_().(D15_DC39)
	return ok
}

type D15_DC38 struct {
	Cf55 _dafny.Map
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf55 _dafny.Map) D15 {
	return D15{D15_DC38{Cf55}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

type D15_DC40 struct {
	Cf59 D15
}

func (D15_DC40) isD15() {}

func (CompanionStruct_D15_) Create_DC40_(Cf59 D15) D15 {
	return D15{D15_DC40{Cf59}}
}

func (_this D15) Is_DC40() bool {
	_, ok := _this.Get_().(D15_DC40)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC39_(_dafny.EmptyMap, false, _dafny.Zero)
}

func (_this D15) Dtor_cf56() _dafny.Map {
	return _this.Get_().(D15_DC39).Cf56
}

func (_this D15) Dtor_cf57() bool {
	return _this.Get_().(D15_DC39).Cf57
}

func (_this D15) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D15_DC39).Cf58
}

func (_this D15) Dtor_cf55() _dafny.Map {
	return _this.Get_().(D15_DC38).Cf55
}

func (_this D15) Dtor_cf59() D15 {
	return _this.Get_().(D15_DC40).Cf59
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC39:
		{
			return "D15.DC39" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ")"
		}
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D15_DC40:
		{
			return "D15.DC40" + "(" + _dafny.String(data.Cf59) + ")"
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
			return ok && data1.Cf56.Equals(data2.Cf56) && data1.Cf57 == data2.Cf57 && data1.Cf58.Cmp(data2.Cf58) == 0
		}
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && data1.Cf55.Equals(data2.Cf55)
		}
	case D15_DC40:
		{
			data2, ok := other.Get_().(D15_DC40)
			return ok && data1.Cf59.Equals(data2.Cf59)
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

type D16_DC42 struct {
	Cf61 bool
	Cf62 bool
	Cf63 bool
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf61 bool, Cf62 bool, Cf63 bool) D16 {
	return D16{D16_DC42{Cf61, Cf62, Cf63}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

type D16_DC43 struct {
	Cf64 _dafny.Int
	Cf65 _dafny.CodePoint
}

func (D16_DC43) isD16() {}

func (CompanionStruct_D16_) Create_DC43_(Cf64 _dafny.Int, Cf65 _dafny.CodePoint) D16 {
	return D16{D16_DC43{Cf64, Cf65}}
}

func (_this D16) Is_DC43() bool {
	_, ok := _this.Get_().(D16_DC43)
	return ok
}

type D16_DC41 struct {
	Cf60 _dafny.Sequence
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf60 _dafny.Sequence) D16 {
	return D16{D16_DC41{Cf60}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

type D16_DC44 struct {
	Cf66 D16
}

func (D16_DC44) isD16() {}

func (CompanionStruct_D16_) Create_DC44_(Cf66 D16) D16 {
	return D16{D16_DC44{Cf66}}
}

func (_this D16) Is_DC44() bool {
	_, ok := _this.Get_().(D16_DC44)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC42_(false, false, false)
}

func (_this D16) Dtor_cf61() bool {
	return _this.Get_().(D16_DC42).Cf61
}

func (_this D16) Dtor_cf62() bool {
	return _this.Get_().(D16_DC42).Cf62
}

func (_this D16) Dtor_cf63() bool {
	return _this.Get_().(D16_DC42).Cf63
}

func (_this D16) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D16_DC43).Cf64
}

func (_this D16) Dtor_cf65() _dafny.CodePoint {
	return _this.Get_().(D16_DC43).Cf65
}

func (_this D16) Dtor_cf60() _dafny.Sequence {
	return _this.Get_().(D16_DC41).Cf60
}

func (_this D16) Dtor_cf66() D16 {
	return _this.Get_().(D16_DC44).Cf66
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ")"
		}
	case D16_DC43:
		{
			return "D16.DC43" + "(" + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf60) + ")"
		}
	case D16_DC44:
		{
			return "D16.DC44" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC42:
		{
			data2, ok := other.Get_().(D16_DC42)
			return ok && data1.Cf61 == data2.Cf61 && data1.Cf62 == data2.Cf62 && data1.Cf63 == data2.Cf63
		}
	case D16_DC43:
		{
			data2, ok := other.Get_().(D16_DC43)
			return ok && data1.Cf64.Cmp(data2.Cf64) == 0 && data1.Cf65 == data2.Cf65
		}
	case D16_DC41:
		{
			data2, ok := other.Get_().(D16_DC41)
			return ok && data1.Cf60.Equals(data2.Cf60)
		}
	case D16_DC44:
		{
			data2, ok := other.Get_().(D16_DC44)
			return ok && data1.Cf66.Equals(data2.Cf66)
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

type D17_DC46 struct {
	Cf68 _dafny.Int
}

func (D17_DC46) isD17() {}

func (CompanionStruct_D17_) Create_DC46_(Cf68 _dafny.Int) D17 {
	return D17{D17_DC46{Cf68}}
}

func (_this D17) Is_DC46() bool {
	_, ok := _this.Get_().(D17_DC46)
	return ok
}

type D17_DC47 struct {
	Cf69 T0
	Cf70 bool
	Cf71 _dafny.Int
	Cf72 _dafny.Int
	Cf73 _dafny.Int
}

func (D17_DC47) isD17() {}

func (CompanionStruct_D17_) Create_DC47_(Cf69 T0, Cf70 bool, Cf71 _dafny.Int, Cf72 _dafny.Int, Cf73 _dafny.Int) D17 {
	return D17{D17_DC47{Cf69, Cf70, Cf71, Cf72, Cf73}}
}

func (_this D17) Is_DC47() bool {
	_, ok := _this.Get_().(D17_DC47)
	return ok
}

type D17_DC45 struct {
	Cf67 _dafny.Map
}

func (D17_DC45) isD17() {}

func (CompanionStruct_D17_) Create_DC45_(Cf67 _dafny.Map) D17 {
	return D17{D17_DC45{Cf67}}
}

func (_this D17) Is_DC45() bool {
	_, ok := _this.Get_().(D17_DC45)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC46_(_dafny.Zero)
}

func (_this D17) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D17_DC46).Cf68
}

func (_this D17) Dtor_cf69() T0 {
	return _this.Get_().(D17_DC47).Cf69
}

func (_this D17) Dtor_cf70() bool {
	return _this.Get_().(D17_DC47).Cf70
}

func (_this D17) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D17_DC47).Cf71
}

func (_this D17) Dtor_cf72() _dafny.Int {
	return _this.Get_().(D17_DC47).Cf72
}

func (_this D17) Dtor_cf73() _dafny.Int {
	return _this.Get_().(D17_DC47).Cf73
}

func (_this D17) Dtor_cf67() _dafny.Map {
	return _this.Get_().(D17_DC45).Cf67
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC46:
		{
			return "D17.DC46" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D17_DC47:
		{
			return "D17.DC47" + "(" + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ")"
		}
	case D17_DC45:
		{
			return "D17.DC45" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC46:
		{
			data2, ok := other.Get_().(D17_DC46)
			return ok && data1.Cf68.Cmp(data2.Cf68) == 0
		}
	case D17_DC47:
		{
			data2, ok := other.Get_().(D17_DC47)
			return ok && _dafny.AreEqual(data1.Cf69, data2.Cf69) && data1.Cf70 == data2.Cf70 && data1.Cf71.Cmp(data2.Cf71) == 0 && data1.Cf72.Cmp(data2.Cf72) == 0 && data1.Cf73.Cmp(data2.Cf73) == 0
		}
	case D17_DC45:
		{
			data2, ok := other.Get_().(D17_DC45)
			return ok && data1.Cf67.Equals(data2.Cf67)
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
	Cf75 bool
	Cf76 bool
	Cf77 _dafny.CodePoint
	Cf78 bool
}

func (D18_DC49) isD18() {}

func (CompanionStruct_D18_) Create_DC49_(Cf75 bool, Cf76 bool, Cf77 _dafny.CodePoint, Cf78 bool) D18 {
	return D18{D18_DC49{Cf75, Cf76, Cf77, Cf78}}
}

func (_this D18) Is_DC49() bool {
	_, ok := _this.Get_().(D18_DC49)
	return ok
}

type D18_DC48 struct {
	Cf74 _dafny.Map
}

func (D18_DC48) isD18() {}

func (CompanionStruct_D18_) Create_DC48_(Cf74 _dafny.Map) D18 {
	return D18{D18_DC48{Cf74}}
}

func (_this D18) Is_DC48() bool {
	_, ok := _this.Get_().(D18_DC48)
	return ok
}

type D18_DC50 struct {
	Cf79 D18
}

func (D18_DC50) isD18() {}

func (CompanionStruct_D18_) Create_DC50_(Cf79 D18) D18 {
	return D18{D18_DC50{Cf79}}
}

func (_this D18) Is_DC50() bool {
	_, ok := _this.Get_().(D18_DC50)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC49_(false, false, _dafny.CodePoint('D'), false)
}

func (_this D18) Dtor_cf75() bool {
	return _this.Get_().(D18_DC49).Cf75
}

func (_this D18) Dtor_cf76() bool {
	return _this.Get_().(D18_DC49).Cf76
}

func (_this D18) Dtor_cf77() _dafny.CodePoint {
	return _this.Get_().(D18_DC49).Cf77
}

func (_this D18) Dtor_cf78() bool {
	return _this.Get_().(D18_DC49).Cf78
}

func (_this D18) Dtor_cf74() _dafny.Map {
	return _this.Get_().(D18_DC48).Cf74
}

func (_this D18) Dtor_cf79() D18 {
	return _this.Get_().(D18_DC50).Cf79
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC49:
		{
			return "D18.DC49" + "(" + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ", " + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ")"
		}
	case D18_DC48:
		{
			return "D18.DC48" + "(" + _dafny.String(data.Cf74) + ")"
		}
	case D18_DC50:
		{
			return "D18.DC50" + "(" + _dafny.String(data.Cf79) + ")"
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
			data2, ok := other.Get_().(D18_DC49)
			return ok && data1.Cf75 == data2.Cf75 && data1.Cf76 == data2.Cf76 && data1.Cf77 == data2.Cf77 && data1.Cf78 == data2.Cf78
		}
	case D18_DC48:
		{
			data2, ok := other.Get_().(D18_DC48)
			return ok && data1.Cf74.Equals(data2.Cf74)
		}
	case D18_DC50:
		{
			data2, ok := other.Get_().(D18_DC50)
			return ok && data1.Cf79.Equals(data2.Cf79)
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

type D19_DC52 struct {
	Cf81 _dafny.Int
	Cf82 bool
	Cf83 bool
}

func (D19_DC52) isD19() {}

func (CompanionStruct_D19_) Create_DC52_(Cf81 _dafny.Int, Cf82 bool, Cf83 bool) D19 {
	return D19{D19_DC52{Cf81, Cf82, Cf83}}
}

func (_this D19) Is_DC52() bool {
	_, ok := _this.Get_().(D19_DC52)
	return ok
}

type D19_DC53 struct {
}

func (D19_DC53) isD19() {}

func (CompanionStruct_D19_) Create_DC53_() D19 {
	return D19{D19_DC53{}}
}

func (_this D19) Is_DC53() bool {
	_, ok := _this.Get_().(D19_DC53)
	return ok
}

type D19_DC51 struct {
	Cf80 _dafny.Sequence
}

func (D19_DC51) isD19() {}

func (CompanionStruct_D19_) Create_DC51_(Cf80 _dafny.Sequence) D19 {
	return D19{D19_DC51{Cf80}}
}

func (_this D19) Is_DC51() bool {
	_, ok := _this.Get_().(D19_DC51)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC52_(_dafny.Zero, false, false)
}

func (_this D19) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D19_DC52).Cf81
}

func (_this D19) Dtor_cf82() bool {
	return _this.Get_().(D19_DC52).Cf82
}

func (_this D19) Dtor_cf83() bool {
	return _this.Get_().(D19_DC52).Cf83
}

func (_this D19) Dtor_cf80() _dafny.Sequence {
	return _this.Get_().(D19_DC51).Cf80
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC52:
		{
			return "D19.DC52" + "(" + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ", " + _dafny.String(data.Cf83) + ")"
		}
	case D19_DC53:
		{
			return "D19.DC53"
		}
	case D19_DC51:
		{
			return "D19.DC51" + "(" + _dafny.String(data.Cf80) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC52:
		{
			data2, ok := other.Get_().(D19_DC52)
			return ok && data1.Cf81.Cmp(data2.Cf81) == 0 && data1.Cf82 == data2.Cf82 && data1.Cf83 == data2.Cf83
		}
	case D19_DC53:
		{
			_, ok := other.Get_().(D19_DC53)
			return ok
		}
	case D19_DC51:
		{
			data2, ok := other.Get_().(D19_DC51)
			return ok && data1.Cf80.Equals(data2.Cf80)
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
	Cf85 _dafny.Int
	Cf86 _dafny.Int
	Cf87 bool
	Cf88 bool
}

func (D20_DC55) isD20() {}

func (CompanionStruct_D20_) Create_DC55_(Cf85 _dafny.Int, Cf86 _dafny.Int, Cf87 bool, Cf88 bool) D20 {
	return D20{D20_DC55{Cf85, Cf86, Cf87, Cf88}}
}

func (_this D20) Is_DC55() bool {
	_, ok := _this.Get_().(D20_DC55)
	return ok
}

type D20_DC54 struct {
	Cf84 _dafny.Array
}

func (D20_DC54) isD20() {}

func (CompanionStruct_D20_) Create_DC54_(Cf84 _dafny.Array) D20 {
	return D20{D20_DC54{Cf84}}
}

func (_this D20) Is_DC54() bool {
	_, ok := _this.Get_().(D20_DC54)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC55_(_dafny.Zero, _dafny.Zero, false, false)
}

func (_this D20) Dtor_cf85() _dafny.Int {
	return _this.Get_().(D20_DC55).Cf85
}

func (_this D20) Dtor_cf86() _dafny.Int {
	return _this.Get_().(D20_DC55).Cf86
}

func (_this D20) Dtor_cf87() bool {
	return _this.Get_().(D20_DC55).Cf87
}

func (_this D20) Dtor_cf88() bool {
	return _this.Get_().(D20_DC55).Cf88
}

func (_this D20) Dtor_cf84() _dafny.Array {
	return _this.Get_().(D20_DC54).Cf84
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC55:
		{
			return "D20.DC55" + "(" + _dafny.String(data.Cf85) + ", " + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ")"
		}
	case D20_DC54:
		{
			return "D20.DC54" + "(" + _dafny.String(data.Cf84) + ")"
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
			return ok && data1.Cf85.Cmp(data2.Cf85) == 0 && data1.Cf86.Cmp(data2.Cf86) == 0 && data1.Cf87 == data2.Cf87 && data1.Cf88 == data2.Cf88
		}
	case D20_DC54:
		{
			data2, ok := other.Get_().(D20_DC54)
			return ok && data1.Cf84 == data2.Cf84
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

// Definition of datatype D21
type D21 struct {
	Data_D21_
}

func (_this D21) Get_() Data_D21_ {
	return _this.Data_D21_
}

type Data_D21_ interface {
	isD21()
}

type CompanionStruct_D21_ struct {
}

var Companion_D21_ = CompanionStruct_D21_{}

type D21_DC57 struct {
	Cf90 _dafny.Int
	Cf91 _dafny.Int
	Cf92 D5
}

func (D21_DC57) isD21() {}

func (CompanionStruct_D21_) Create_DC57_(Cf90 _dafny.Int, Cf91 _dafny.Int, Cf92 D5) D21 {
	return D21{D21_DC57{Cf90, Cf91, Cf92}}
}

func (_this D21) Is_DC57() bool {
	_, ok := _this.Get_().(D21_DC57)
	return ok
}

type D21_DC56 struct {
	Cf89 _dafny.Array
}

func (D21_DC56) isD21() {}

func (CompanionStruct_D21_) Create_DC56_(Cf89 _dafny.Array) D21 {
	return D21{D21_DC56{Cf89}}
}

func (_this D21) Is_DC56() bool {
	_, ok := _this.Get_().(D21_DC56)
	return ok
}

type D21_DC58 struct {
	Cf93 D21
}

func (D21_DC58) isD21() {}

func (CompanionStruct_D21_) Create_DC58_(Cf93 D21) D21 {
	return D21{D21_DC58{Cf93}}
}

func (_this D21) Is_DC58() bool {
	_, ok := _this.Get_().(D21_DC58)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC57_(_dafny.Zero, _dafny.Zero, Companion_D5_.Default())
}

func (_this D21) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D21_DC57).Cf90
}

func (_this D21) Dtor_cf91() _dafny.Int {
	return _this.Get_().(D21_DC57).Cf91
}

func (_this D21) Dtor_cf92() D5 {
	return _this.Get_().(D21_DC57).Cf92
}

func (_this D21) Dtor_cf89() _dafny.Array {
	return _this.Get_().(D21_DC56).Cf89
}

func (_this D21) Dtor_cf93() D21 {
	return _this.Get_().(D21_DC58).Cf93
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC57:
		{
			return "D21.DC57" + "(" + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ")"
		}
	case D21_DC56:
		{
			return "D21.DC56" + "(" + _dafny.String(data.Cf89) + ")"
		}
	case D21_DC58:
		{
			return "D21.DC58" + "(" + _dafny.String(data.Cf93) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC57:
		{
			data2, ok := other.Get_().(D21_DC57)
			return ok && data1.Cf90.Cmp(data2.Cf90) == 0 && data1.Cf91.Cmp(data2.Cf91) == 0 && data1.Cf92.Equals(data2.Cf92)
		}
	case D21_DC56:
		{
			data2, ok := other.Get_().(D21_DC56)
			return ok && data1.Cf89 == data2.Cf89
		}
	case D21_DC58:
		{
			data2, ok := other.Get_().(D21_DC58)
			return ok && data1.Cf93.Equals(data2.Cf93)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D21) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D21)
	return ok && _this.Equals(typed)
}

func Type_D21_() _dafny.TypeDescriptor {
	return type_D21_{}
}

type type_D21_ struct {
}

func (_this type_D21_) Default() interface{} {
	return Companion_D21_.Default()
}

func (_this type_D21_) String() string {
	return "main.D21"
}
func (_this D21) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D21{}

// End of datatype D21

// Definition of datatype D22
type D22 struct {
	Data_D22_
}

func (_this D22) Get_() Data_D22_ {
	return _this.Data_D22_
}

type Data_D22_ interface {
	isD22()
}

type CompanionStruct_D22_ struct {
}

var Companion_D22_ = CompanionStruct_D22_{}

type D22_DC60 struct {
	Cf95 _dafny.Int
	Cf96 _dafny.Int
}

func (D22_DC60) isD22() {}

func (CompanionStruct_D22_) Create_DC60_(Cf95 _dafny.Int, Cf96 _dafny.Int) D22 {
	return D22{D22_DC60{Cf95, Cf96}}
}

func (_this D22) Is_DC60() bool {
	_, ok := _this.Get_().(D22_DC60)
	return ok
}

type D22_DC61 struct {
	Cf97  _dafny.Int
	Cf98  _dafny.Int
	Cf99  _dafny.Int
	Cf100 _dafny.Sequence
}

func (D22_DC61) isD22() {}

func (CompanionStruct_D22_) Create_DC61_(Cf97 _dafny.Int, Cf98 _dafny.Int, Cf99 _dafny.Int, Cf100 _dafny.Sequence) D22 {
	return D22{D22_DC61{Cf97, Cf98, Cf99, Cf100}}
}

func (_this D22) Is_DC61() bool {
	_, ok := _this.Get_().(D22_DC61)
	return ok
}

type D22_DC62 struct {
	Cf101 _dafny.Sequence
	Cf102 _dafny.Int
	Cf103 _dafny.Int
	Cf104 _dafny.Int
	Cf105 bool
}

func (D22_DC62) isD22() {}

func (CompanionStruct_D22_) Create_DC62_(Cf101 _dafny.Sequence, Cf102 _dafny.Int, Cf103 _dafny.Int, Cf104 _dafny.Int, Cf105 bool) D22 {
	return D22{D22_DC62{Cf101, Cf102, Cf103, Cf104, Cf105}}
}

func (_this D22) Is_DC62() bool {
	_, ok := _this.Get_().(D22_DC62)
	return ok
}

type D22_DC59 struct {
	Cf94 _dafny.MultiSet
}

func (D22_DC59) isD22() {}

func (CompanionStruct_D22_) Create_DC59_(Cf94 _dafny.MultiSet) D22 {
	return D22{D22_DC59{Cf94}}
}

func (_this D22) Is_DC59() bool {
	_, ok := _this.Get_().(D22_DC59)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC60_(_dafny.Zero, _dafny.Zero)
}

func (_this D22) Dtor_cf95() _dafny.Int {
	return _this.Get_().(D22_DC60).Cf95
}

func (_this D22) Dtor_cf96() _dafny.Int {
	return _this.Get_().(D22_DC60).Cf96
}

func (_this D22) Dtor_cf97() _dafny.Int {
	return _this.Get_().(D22_DC61).Cf97
}

func (_this D22) Dtor_cf98() _dafny.Int {
	return _this.Get_().(D22_DC61).Cf98
}

func (_this D22) Dtor_cf99() _dafny.Int {
	return _this.Get_().(D22_DC61).Cf99
}

func (_this D22) Dtor_cf100() _dafny.Sequence {
	return _this.Get_().(D22_DC61).Cf100
}

func (_this D22) Dtor_cf101() _dafny.Sequence {
	return _this.Get_().(D22_DC62).Cf101
}

func (_this D22) Dtor_cf102() _dafny.Int {
	return _this.Get_().(D22_DC62).Cf102
}

func (_this D22) Dtor_cf103() _dafny.Int {
	return _this.Get_().(D22_DC62).Cf103
}

func (_this D22) Dtor_cf104() _dafny.Int {
	return _this.Get_().(D22_DC62).Cf104
}

func (_this D22) Dtor_cf105() bool {
	return _this.Get_().(D22_DC62).Cf105
}

func (_this D22) Dtor_cf94() _dafny.MultiSet {
	return _this.Get_().(D22_DC59).Cf94
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC60:
		{
			return "D22.DC60" + "(" + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ")"
		}
	case D22_DC61:
		{
			return "D22.DC61" + "(" + _dafny.String(data.Cf97) + ", " + _dafny.String(data.Cf98) + ", " + _dafny.String(data.Cf99) + ", " + data.Cf100.VerbatimString(true) + ")"
		}
	case D22_DC62:
		{
			return "D22.DC62" + "(" + data.Cf101.VerbatimString(true) + ", " + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ", " + _dafny.String(data.Cf104) + ", " + _dafny.String(data.Cf105) + ")"
		}
	case D22_DC59:
		{
			return "D22.DC59" + "(" + _dafny.String(data.Cf94) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC60:
		{
			data2, ok := other.Get_().(D22_DC60)
			return ok && data1.Cf95.Cmp(data2.Cf95) == 0 && data1.Cf96.Cmp(data2.Cf96) == 0
		}
	case D22_DC61:
		{
			data2, ok := other.Get_().(D22_DC61)
			return ok && data1.Cf97.Cmp(data2.Cf97) == 0 && data1.Cf98.Cmp(data2.Cf98) == 0 && data1.Cf99.Cmp(data2.Cf99) == 0 && data1.Cf100.Equals(data2.Cf100)
		}
	case D22_DC62:
		{
			data2, ok := other.Get_().(D22_DC62)
			return ok && data1.Cf101.Equals(data2.Cf101) && data1.Cf102.Cmp(data2.Cf102) == 0 && data1.Cf103.Cmp(data2.Cf103) == 0 && data1.Cf104.Cmp(data2.Cf104) == 0 && data1.Cf105 == data2.Cf105
		}
	case D22_DC59:
		{
			data2, ok := other.Get_().(D22_DC59)
			return ok && data1.Cf94.Equals(data2.Cf94)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D22) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D22)
	return ok && _this.Equals(typed)
}

func Type_D22_() _dafny.TypeDescriptor {
	return type_D22_{}
}

type type_D22_ struct {
}

func (_this type_D22_) Default() interface{} {
	return Companion_D22_.Default()
}

func (_this type_D22_) String() string {
	return "main.D22"
}
func (_this D22) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D22{}

// End of datatype D22

// Definition of datatype D23
type D23 struct {
	Data_D23_
}

func (_this D23) Get_() Data_D23_ {
	return _this.Data_D23_
}

type Data_D23_ interface {
	isD23()
}

type CompanionStruct_D23_ struct {
}

var Companion_D23_ = CompanionStruct_D23_{}

type D23_DC64 struct {
	Cf107 _dafny.Int
	Cf108 _dafny.Int
}

func (D23_DC64) isD23() {}

func (CompanionStruct_D23_) Create_DC64_(Cf107 _dafny.Int, Cf108 _dafny.Int) D23 {
	return D23{D23_DC64{Cf107, Cf108}}
}

func (_this D23) Is_DC64() bool {
	_, ok := _this.Get_().(D23_DC64)
	return ok
}

type D23_DC63 struct {
	Cf106 _dafny.Map
}

func (D23_DC63) isD23() {}

func (CompanionStruct_D23_) Create_DC63_(Cf106 _dafny.Map) D23 {
	return D23{D23_DC63{Cf106}}
}

func (_this D23) Is_DC63() bool {
	_, ok := _this.Get_().(D23_DC63)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC64_(_dafny.Zero, _dafny.Zero)
}

func (_this D23) Dtor_cf107() _dafny.Int {
	return _this.Get_().(D23_DC64).Cf107
}

func (_this D23) Dtor_cf108() _dafny.Int {
	return _this.Get_().(D23_DC64).Cf108
}

func (_this D23) Dtor_cf106() _dafny.Map {
	return _this.Get_().(D23_DC63).Cf106
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC64:
		{
			return "D23.DC64" + "(" + _dafny.String(data.Cf107) + ", " + _dafny.String(data.Cf108) + ")"
		}
	case D23_DC63:
		{
			return "D23.DC63" + "(" + _dafny.String(data.Cf106) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC64:
		{
			data2, ok := other.Get_().(D23_DC64)
			return ok && data1.Cf107.Cmp(data2.Cf107) == 0 && data1.Cf108.Cmp(data2.Cf108) == 0
		}
	case D23_DC63:
		{
			data2, ok := other.Get_().(D23_DC63)
			return ok && data1.Cf106.Equals(data2.Cf106)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D23) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D23)
	return ok && _this.Equals(typed)
}

func Type_D23_() _dafny.TypeDescriptor {
	return type_D23_{}
}

type type_D23_ struct {
}

func (_this type_D23_) Default() interface{} {
	return Companion_D23_.Default()
}

func (_this type_D23_) String() string {
	return "main.D23"
}
func (_this D23) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D23{}

// End of datatype D23

// Definition of trait T0
type T0 interface {
	String() string
	Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool
	Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState)
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
	Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet
	Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool
	M6(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 D2, globalState *GlobalState)
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
	F4   _dafny.Int
	F7   _dafny.Int
	F13  _dafny.Int
	F15  _dafny.Int
	_f0  _dafny.Sequence
	_f1  bool
	_f2  _dafny.CodePoint
	_f3  _dafny.Int
	_f5  bool
	_f6  _dafny.Map
	_f8  _dafny.Set
	_f9  bool
	_f10 bool
	_f11 _dafny.Map
	_f12 _dafny.Int
	_f14 bool
	_f16 _dafny.MultiSet
	_f17 bool
	_f18 _dafny.Int
	_f19 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F4 = _dafny.Zero
	_this.F7 = _dafny.Zero
	_this.F13 = _dafny.Zero
	_this.F15 = _dafny.Zero
	_this._f0 = _dafny.EmptySeq
	_this._f1 = false
	_this._f2 = _dafny.CodePoint('D')
	_this._f3 = _dafny.Zero
	_this._f5 = false
	_this._f6 = _dafny.EmptyMap
	_this._f8 = _dafny.EmptySet
	_this._f9 = false
	_this._f10 = false
	_this._f11 = _dafny.EmptyMap
	_this._f12 = _dafny.Zero
	_this._f14 = false
	_this._f16 = _dafny.EmptyMultiSet
	_this._f17 = false
	_this._f18 = _dafny.Zero
	_this._f19 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Sequence, f1 bool, f2 _dafny.CodePoint, f3 _dafny.Int, f4 _dafny.Int, f5 bool, f6 _dafny.Map, f7 _dafny.Int, f8 _dafny.Set, f9 bool, f10 bool, f11 _dafny.Map, f12 _dafny.Int, f13 _dafny.Int, f14 bool, f15 _dafny.Int, f16 _dafny.MultiSet, f17 bool, f18 _dafny.Int, f19 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
	}
}
func (_this *GlobalState) F0() _dafny.Sequence {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.CodePoint {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Map {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F8() _dafny.Set {
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
func (_this *GlobalState) F11() _dafny.Map {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F16() _dafny.MultiSet {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Int {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() bool {
	{
		return _this._f19
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
func (_this *C0) Fm28(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(852))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_215_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), _dafny.UnicodeSeqOfUtf8Bytes("o")), _dafny.UnicodeSeqOfUtf8Bytes("fdqgbyeq"))
	}
}
func (_this *C0) Fm29(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (Companion_D5_.Create_DC14_(_dafny.CodePoint('n'), false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D7_.Create_DC21_(_dafny.IntOfInt64(739), (func() _dafny.Map {
			var _coll20 = _dafny.NewMapBuilder()
			_ = _coll20
			for _iter20 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sosvg")).Cardinality()), _dafny.IntOfInt64(968))).Keys().Elements()); ; {
				_compr_20, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _216_v0 _dafny.Int
				_216_v0 = interface{}(_compr_20).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sosvg")).Cardinality()), _dafny.IntOfInt64(968))).Contains(_216_v0) {
					_coll20.Add(Companion_Default___.SafeModuloInt(_216_v0, (_dafny.MultiSetOf(false, false)).Cardinality()), _dafny.IntOfInt64(-322))
				}
			}
			return _coll20.ToMap()
		}()).Cardinality(), true, !(true), false))).Cardinality(), _dafny.IntOfInt64(206))).Cardinality(), !(true))).Dtor_cf21()
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	dummy byte
}

func New_C1_() *C1 {
	_this := C1{}

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

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) Fm25(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(3), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-692), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(997))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}(func(_217_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))).Cardinality())))
	}
}
func (_this *C1) M13(p0 D7, p1 _dafny.Int, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, D7, D1) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 D7 = Companion_D7_.Default()
		_ = r2
		var r3 D1 = Companion_D1_.Default()
		_ = r3
		var _218_v0 bool
		_ = _218_v0
		_218_v0 = true
		var _219_v1 _dafny.Array
		_ = _219_v1
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_4
		var _nw31 _dafny.Array
		_ = _nw31
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw31 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Int = (func(_220_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_221_i0 _dafny.Int) _dafny.Int {
					return (_221_i0).Times(_220_p1)
				}
			})(p1)
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
		_219_v1 = _nw31
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
		_ = _index14
		(_219_v1).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), (_index14).Int())
		var _222_v2 _dafny.Set
		_ = _222_v2
		_222_v2 = _dafny.SetOf(p1)
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
		_ = _index15
		var _rhs30 bool = (_222_v2).IsSubsetOf(_222_v2)
		_ = _rhs30
		var _rhs31 _dafny.Int = _dafny.IntOfInt64(38)
		_ = _rhs31
		var _lhs16 _dafny.Array = _219_v1
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
		_ = _lhs17
		_218_v0 = _rhs30
		(_lhs16).ArraySet1(_rhs31, (_lhs17).Int())
		var _223_v3 _dafny.Sequence
		_ = _223_v3
		_223_v3 = _dafny.SeqOf((_dafny.Zero).Minus((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)))
		var _224_v4 D6
		_ = _224_v4
		_224_v4 = Companion_D6_.Create_DC16_(_223_v3)
		var _225_v5 _dafny.Map
		_ = _225_v5
		_225_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_218_v0, _218_v0)
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
		_ = _index16
		(_219_v1).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.MultiSetFromSeq((_224_v4).Dtor_cf23())).Cardinality()), ((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)).Times((_225_v5).Cardinality())), (_index16).Int())
		var _226_i1 _dafny.Int
		_ = _226_i1
		_226_i1 = _dafny.Zero
		{
			for Companion_Default___.Fm26(globalState) {
				{
					if (_226_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_226_i1 = (_226_i1).Plus(_dafny.One)
					var _227_v6 _dafny.Array
					_ = _227_v6
					var _len0_5 _dafny.Int = _dafny.IntOfInt64(23)
					_ = _len0_5
					var _nw32 _dafny.Array
					_ = _nw32
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw32 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) bool = (func(_228_v0 bool) func(_dafny.Int) bool {
							return func(_229_i2 _dafny.Int) bool {
								return _228_v0
							}
						})(_218_v0)
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw32 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw32).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw32).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_227_v6 = _nw32
					var _230_v7 _dafny.Sequence
					_ = _230_v7
					_230_v7 = _dafny.SeqOf(_227_v6, _227_v6)
					var _source5 D7 = Companion_D7_.Create_DC19_(_dafny.Companion_Sequence_.Concatenate(_230_v7, _dafny.Companion_Sequence_.Update(_230_v7, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_230_v7).Cardinality()))).Uint32(), _227_v6)))
					_ = _source5
					if _source5.Is_DC20() {
						var _231___mcc_h0 _dafny.Int = _source5.Get_().(D7_DC20).Cf29
						_ = _231___mcc_h0
						var _232_cf29 _dafny.Int = _231___mcc_h0
						_ = _232_cf29
						var _233_v8 _dafny.Sequence
						_ = _233_v8
						_233_v8 = _dafny.UnicodeSeqOfUtf8Bytes("slrsv")
						var _234_v9 _dafny.CodePoint
						_ = _234_v9
						_234_v9 = _dafny.CodePoint('c')
						var _235_v10 _dafny.Map
						_ = _235_v10
						_235_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!_dafny.Companion_Sequence_.Equal(_233_v8, _dafny.Companion_Sequence_.Update(_233_v8, (Companion_Default___.SafeIndex((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_233_v8).Cardinality()))).Uint32(), _234_v9)))), p1)
						_235_v10 = (_235_v10).Update(_218_v0, p1)
						var _236_v11 _dafny.MultiSet
						_ = _236_v11
						_236_v11 = _dafny.MultiSetOf(_218_v0)
						var _237_v12 _dafny.Sequence
						_ = _237_v12
						_237_v12 = _dafny.SeqOf(_218_v0)
						_218_v0 = !((((_236_v11).Difference(_dafny.MultiSetFromSeq(_237_v12))).Cardinality()).Cmp(_232_cf29) > 0)
						_218_v0 = Companion_Default___.Fm26(globalState)
						var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_227_v6), 0))
						_ = _index17
						(_227_v6).ArraySet1(_218_v0, (_index17).Int())
						var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_227_v6), 0))
						_ = _index18
						(_227_v6).ArraySet1(_218_v0, (_index18).Int())
					} else if _source5.Is_DC21() {
						var _238___mcc_h1 _dafny.Int = _source5.Get_().(D7_DC21).Cf30
						_ = _238___mcc_h1
						var _239___mcc_h2 _dafny.Int = _source5.Get_().(D7_DC21).Cf31
						_ = _239___mcc_h2
						var _240___mcc_h3 bool = _source5.Get_().(D7_DC21).Cf32
						_ = _240___mcc_h3
						var _241___mcc_h4 bool = _source5.Get_().(D7_DC21).Cf33
						_ = _241___mcc_h4
						var _242___mcc_h5 bool = _source5.Get_().(D7_DC21).Cf34
						_ = _242___mcc_h5
						var _243_cf34 bool = _242___mcc_h5
						_ = _243_cf34
						var _244_cf33 bool = _241___mcc_h4
						_ = _244_cf33
						var _245_cf32 bool = _240___mcc_h3
						_ = _245_cf32
						var _246_cf31 _dafny.Int = _239___mcc_h2
						_ = _246_cf31
						var _247_cf30 _dafny.Int = _238___mcc_h1
						_ = _247_cf30
						_246_cf31 = _246_cf31
						var _248_v13 _dafny.Sequence
						_ = _248_v13
						_248_v13 = _dafny.SeqOf(_245_cf32)
						var _249_v14 _dafny.Array
						_ = _249_v14
						var _nwElement0_5 _dafny.Sequence = _248_v13
						_ = _nwElement0_5
						var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(28))
						_ = _nw33
						(_nw33).ArraySet1(_nwElement0_5, 0)
						(_nw33).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_248_v13, _dafny.Companion_Sequence_.Update(_248_v13, (Companion_Default___.SafeIndex(_246_cf31, _dafny.IntOfUint32((_248_v13).Cardinality()))).Uint32(), _243_cf34)), 1)
						(_nw33).ArraySet1(_248_v13, 2)
						(_nw33).ArraySet1(_dafny.Companion_Sequence_.Update(_248_v13, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_248_v13).Cardinality()))).Uint32(), !(_218_v0)), 3)
						(_nw33).ArraySet1(_dafny.SeqOf(_218_v0, !(!(_245_cf32))), 4)
						(_nw33).ArraySet1(_248_v13, 5)
						(_nw33).ArraySet1(_248_v13, 6)
						(_nw33).ArraySet1(Companion_Default___.Fm27(globalState), 7)
						(_nw33).ArraySet1(_248_v13, 8)
						(_nw33).ArraySet1(_248_v13, 9)
						(_nw33).ArraySet1(Companion_Default___.Fm27(globalState), 10)
						(_nw33).ArraySet1(_dafny.Companion_Sequence_.Update(_248_v13, (Companion_Default___.SafeIndex(_246_cf31, _dafny.IntOfUint32((_248_v13).Cardinality()))).Uint32(), Companion_Default___.Fm26(globalState)), 11)
						(_nw33).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, _245_cf32, _218_v0, _243_cf34), _248_v13), 12)
						(_nw33).ArraySet1(_dafny.SeqOf(_245_cf32), 13)
						(_nw33).ArraySet1(_248_v13, 14)
						(_nw33).ArraySet1(_248_v13, 15)
						(_nw33).ArraySet1(_248_v13, 16)
						(_nw33).ArraySet1(_248_v13, 17)
						(_nw33).ArraySet1(_248_v13, 18)
						(_nw33).ArraySet1(_248_v13, 19)
						(_nw33).ArraySet1(_dafny.SeqOf(_243_cf34, _218_v0), 20)
						(_nw33).ArraySet1(_248_v13, 21)
						(_nw33).ArraySet1(_dafny.SeqOf(_244_cf33, _218_v0), 22)
						(_nw33).ArraySet1(_dafny.SeqOf(_244_cf33, _243_cf34, false, true, _245_cf32), 23)
						(_nw33).ArraySet1(_248_v13, 24)
						(_nw33).ArraySet1(_248_v13, 25)
						(_nw33).ArraySet1(_248_v13, 26)
						(_nw33).ArraySet1(_248_v13, 27)
						_249_v14 = _nw33
						var _250_v15 _dafny.Array
						_ = _250_v15
						_250_v15 = _249_v14
						var _251_v16 _dafny.Sequence
						_ = _251_v16
						_251_v16 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_223_v3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.IntOfUint32((_223_v3).Cardinality()))).Uint32(), p1))
						var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
						_ = _index19
						var _rhs32 bool = (p1).Cmp(_247_cf30) <= 0
						_ = _rhs32
						var _rhs33 _dafny.Array = (_250_v15)
						_ = _rhs33
						var _rhs34 _dafny.Int = _dafny.IntOfUint32((_251_v16).Cardinality())
						_ = _rhs34
						var _rhs35 _dafny.Int = (_247_cf30).Plus(p1)
						_ = _rhs35
						var _lhs18 *GlobalState = globalState
						_ = _lhs18
						var _lhs19 _dafny.Array = _219_v1
						_ = _lhs19
						var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
						_ = _lhs20
						_244_cf33 = _rhs32
						_249_v14 = _rhs33
						_lhs18.F7 = _rhs34
						(_lhs19).ArraySet1(_rhs35, (_lhs20).Int())
						var _252_v17 _dafny.MultiSet
						_ = _252_v17
						_252_v17 = _dafny.MultiSetOf(_247_cf30, (_dafny.Zero).Minus((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)))
						var _253_v18 _dafny.CodePoint
						_ = _253_v18
						_253_v18 = _dafny.CodePoint('r')
						var _254_v19 _dafny.Sequence
						_ = _254_v19
						_254_v19 = _dafny.UnicodeSeqOfUtf8Bytes("fk")
						var _255_v20 _dafny.Sequence
						_ = _255_v20
						_255_v20 = _dafny.SeqOf((_254_v19).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_254_v19).Cardinality()))).Uint32()).(_dafny.CodePoint))
						var _256_v21 _dafny.Map
						_ = _256_v21
						_256_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_252_v17, Companion_D5_.Create_DC14_(_253_v18, (_248_v13).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_248_v13).Cardinality()))).Uint32()).(bool), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_254_v19).Cardinality()))), _218_v0))
						_256_v21 = _256_v21
						_250_v15 = _249_v14
					} else if _source5.Is_DC22() {
						var _257___mcc_h6 bool = _source5.Get_().(D7_DC22).Cf35
						_ = _257___mcc_h6
						var _258___mcc_h7 _dafny.Int = _source5.Get_().(D7_DC22).Cf36
						_ = _258___mcc_h7
						var _259___mcc_h8 bool = _source5.Get_().(D7_DC22).Cf37
						_ = _259___mcc_h8
						var _260___mcc_h9 _dafny.Int = _source5.Get_().(D7_DC22).Cf38
						_ = _260___mcc_h9
						var _261_cf38 _dafny.Int = _260___mcc_h9
						_ = _261_cf38
						var _262_cf37 bool = _259___mcc_h8
						_ = _262_cf37
						var _263_cf36 _dafny.Int = _258___mcc_h7
						_ = _263_cf36
						var _264_cf35 bool = _257___mcc_h6
						_ = _264_cf35
						_264_cf35 = !(Companion_Default___.Fm26(globalState))
						var _265_v22 _dafny.Set
						_ = _265_v22
						_265_v22 = _dafny.SetOf(_219_v1, _219_v1)
						var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
						_ = _index20
						(_219_v1).ArraySet1(((_265_v22).Intersection((_265_v22).Difference(_265_v22))).Cardinality(), (_index20).Int())
						_264_cf35 = _264_cf35
						_261_cf38 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-589))
					} else {
						var _266___mcc_h10 _dafny.Sequence = _source5.Get_().(D7_DC19).Cf28
						_ = _266___mcc_h10
						var _267_cf28 _dafny.Sequence = _266___mcc_h10
						_ = _267_cf28
						_218_v0 = _218_v0
						_218_v0 = _218_v0
						var _268_v23 *C0
						_ = _268_v23
						var _nw34 *C0 = New_C0_()
						_ = _nw34
						_nw34.Ctor__()
						_268_v23 = _nw34
						var _269_v24 _dafny.Array
						_ = _269_v24
						var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
						_ = _nw35
						_269_v24 = _nw35
						var _270_v25 _dafny.Sequence
						_ = _270_v25
						_270_v25 = _dafny.UnicodeSeqOfUtf8Bytes("yaw")
						var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_269_v24), 0))
						_ = _index21
						(_269_v24).ArraySet1(_270_v25, (_index21).Int())
						var _271_v26 _dafny.MultiSet
						_ = _271_v26
						_271_v26 = _dafny.MultiSetOf(_219_v1)
						var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_269_v24), 0))
						_ = _index22
						var _rhs36 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_270_v25, _270_v25)
						_ = _rhs36
						var _rhs37 _dafny.Int = (p1).Minus((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int))
						_ = _rhs37
						var _rhs38 bool = !((_271_v26).IsSubsetOf(_271_v26))
						_ = _rhs38
						var _rhs39 _dafny.Int = p1
						_ = _rhs39
						var _rhs40 bool = !(_218_v0)
						_ = _rhs40
						var _lhs21 _dafny.Array = _269_v24
						_ = _lhs21
						var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_269_v24), 0))
						_ = _lhs22
						var _lhs23 *GlobalState = globalState
						_ = _lhs23
						var _lhs24 *GlobalState = globalState
						_ = _lhs24
						(_lhs21).ArraySet1(_rhs36, (_lhs22).Int())
						_lhs23.F7 = _rhs37
						_218_v0 = _rhs38
						_lhs24.F4 = _rhs39
						_218_v0 = _rhs40
					}
					_227_v6 = _227_v6
					var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_227_v6), 0))
					_ = _index23
					(_227_v6).ArraySet1(_218_v0, (_index23).Int())
					var _272_v27 _dafny.Array
					_ = _272_v27
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_6
					var _nw36 _dafny.Array
					_ = _nw36
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw36 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.Sequence = func(_273_i3 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(false)
						}
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw36 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw36).ArraySet1(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw36).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_272_v27 = _nw36
					var _274_v28 _dafny.Sequence
					_ = _274_v28
					_274_v28 = _dafny.SeqOf(_272_v27)
					var _275_v29 _dafny.Array
					_ = _275_v29
					_275_v29 = (_274_v28).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_274_v28).Cardinality()))).Uint32()).(_dafny.Array)
					var _276_v30 _dafny.Set
					_ = _276_v30
					_276_v30 = _dafny.SetOf(_275_v29)
					var _277_v31 _dafny.Map
					_ = _277_v31
					_277_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), _276_v30)
					var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_227_v6), 0))
					_ = _index24
					(_227_v6).ArraySet1((_277_v31).Contains(p1), (_index24).Int())
					var _278_v32 D8
					_ = _278_v32
					_278_v32 = Companion_D8_.Create_DC25_((_dafny.Zero).Minus(_dafny.IntOfUint32((_223_v3).Cardinality())))
					var _279_v33 _dafny.Sequence
					_ = _279_v33
					_279_v33 = _dafny.UnicodeSeqOfUtf8Bytes("cd")
					var _280_v34 _dafny.Sequence
					_ = _280_v34
					_280_v34 = _dafny.SeqOf(_218_v0, _218_v0, _218_v0)
					var _rhs41 D8 = _278_v32
					_ = _rhs41
					var _rhs42 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-416))
					_ = _rhs42
					var _rhs43 _dafny.Array = _219_v1
					_ = _rhs43
					var _rhs44 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_279_v33).Cardinality()))
					_ = _rhs44
					var _rhs45 _dafny.Int = (_dafny.IntOfUint32((_280_v34).Cardinality())).Times((p1).Times((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)))
					_ = _rhs45
					var _lhs25 *GlobalState = globalState
					_ = _lhs25
					var _lhs26 *GlobalState = globalState
					_ = _lhs26
					var _lhs27 *GlobalState = globalState
					_ = _lhs27
					_278_v32 = _rhs41
					_lhs25.F4 = _rhs42
					_219_v1 = _rhs43
					_lhs26.F13 = _rhs44
					_lhs27.F4 = _rhs45
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _281_v35 D2
		_ = _281_v35
		_281_v35 = Companion_D2_.Create_DC5_(_219_v1)
		var _source6 D2 = _281_v35
		_ = _source6
		if _source6.Is_DC6() {
			var _282___mcc_h11 _dafny.Int = _source6.Get_().(D2_DC6).Cf7
			_ = _282___mcc_h11
			var _283_cf7 _dafny.Int = _282___mcc_h11
			_ = _283_cf7
			var _284_v36 _dafny.Map
			_ = _284_v36
			_284_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int), _218_v0)
			var _285_v37 _dafny.Sequence
			_ = _285_v37
			_285_v37 = _dafny.SeqOf(!(_218_v0))
			_218_v0 = !(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _218_v0)).Equals(_284_v36)) == ((_285_v37).Select((Companion_Default___.SafeIndex(_283_cf7, _dafny.IntOfUint32((_285_v37).Cardinality()))).Uint32()).(bool)))
			var _286_v38 _dafny.Map
			_ = _286_v38
			_286_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm25(_218_v0, _218_v0, globalState), _dafny.CodePoint('b'))
			var _287_v39 _dafny.Array
			_ = _287_v39
			var _nwElement0_6 _dafny.Sequence = _285_v37
			_ = _nwElement0_6
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(20))
			_ = _nw37
			(_nw37).ArraySet1(_nwElement0_6, 0)
			(_nw37).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_218_v0), (Companion_Default___.SafeIndex((_286_v38).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_218_v0)).Cardinality()))).Uint32(), false), 1)
			(_nw37).ArraySet1(_285_v37, 2)
			(_nw37).ArraySet1(_285_v37, 3)
			(_nw37).ArraySet1(_dafny.SeqOf(_218_v0), 4)
			(_nw37).ArraySet1(_285_v37, 5)
			(_nw37).ArraySet1(_285_v37, 6)
			(_nw37).ArraySet1(_285_v37, 7)
			(_nw37).ArraySet1(_285_v37, 8)
			(_nw37).ArraySet1(Companion_Default___.Fm27(globalState), 9)
			(_nw37).ArraySet1(_dafny.SeqOf(_218_v0, _218_v0, _218_v0), 10)
			(_nw37).ArraySet1(_285_v37, 11)
			(_nw37).ArraySet1(_285_v37, 12)
			(_nw37).ArraySet1(_285_v37, 13)
			(_nw37).ArraySet1(_dafny.SeqOf(_218_v0, Companion_Default___.Fm26(globalState)), 14)
			(_nw37).ArraySet1(_285_v37, 15)
			(_nw37).ArraySet1(_dafny.SeqOf(!(false), true), 16)
			(_nw37).ArraySet1(_285_v37, 17)
			(_nw37).ArraySet1(_285_v37, 18)
			(_nw37).ArraySet1(_dafny.Companion_Sequence_.Update(_285_v37, (Companion_Default___.SafeIndex((_this).Fm25(_218_v0, !(_218_v0), globalState), _dafny.IntOfUint32((_285_v37).Cardinality()))).Uint32(), _218_v0), 19)
			_287_v39 = _nw37
			var _288_v40 _dafny.Array
			_ = _288_v40
			_288_v40 = _287_v39
			var _source7 _dafny.Array = _288_v40
			_ = _source7
			var _289___mcc_h17 _dafny.Array = _source7
			_ = _289___mcc_h17
			var _290_cf43 _dafny.Array = _289___mcc_h17
			_ = _290_cf43
			var _291_v41 _dafny.Sequence
			_ = _291_v41
			_291_v41 = _dafny.UnicodeSeqOfUtf8Bytes("owcdswdik")
			var _292_v42 _dafny.Sequence
			_ = _292_v42
			_292_v42 = _dafny.SeqOf((_291_v41).Select((Companion_Default___.SafeIndex((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_291_v41).Cardinality()))).Uint32()).(_dafny.CodePoint))
			(globalState).F15 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_291_v41).Cardinality()))
			var _293_v43 D8
			_ = _293_v43
			_293_v43 = Companion_D8_.Create_DC23_(_285_v37)
			var _294_v44 _dafny.Map
			_ = _294_v44
			_294_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_218_v0, _291_v41)
			var _295_v45 _dafny.Map
			_ = _295_v45
			_295_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_294_v44).Update(_218_v0, _291_v41), _218_v0)
			var _296_v46 _dafny.MultiSet
			_ = _296_v46
			_296_v46 = _dafny.MultiSetOf(false, !(_218_v0), _218_v0, (func() bool {
				if (_295_v45).Contains(_294_v44) {
					return (_295_v45).Get(_294_v44).(bool)
				}
				return _218_v0
			})())
			var _rhs46 D8 = _293_v43
			_ = _rhs46
			var _rhs47 _dafny.MultiSet = _296_v46
			_ = _rhs47
			var _rhs48 bool = _218_v0
			_ = _rhs48
			_293_v43 = _rhs46
			_296_v46 = _rhs47
			_218_v0 = _rhs48
			(globalState).F7 = (_283_cf7).Times((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int))
			var _297_v47 _dafny.Map
			_ = _297_v47
			_297_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_283_cf7, _dafny.IntOfUint32((_291_v41).Cardinality())), _283_cf7)
			var _298_v48 _dafny.MultiSet
			_ = _298_v48
			_298_v48 = _dafny.MultiSetOf((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int))
			var _299_v49 _dafny.Set
			_ = _299_v49
			_299_v49 = _dafny.SetOf(_218_v0)
			var _300_v50 _dafny.CodePoint
			_ = _300_v50
			_300_v50 = _dafny.CodePoint('t')
			var _301_v51 _dafny.Array
			_ = _301_v51
			var _nwElement0_7 bool = (_299_v49).IsDisjointFrom(_299_v49)
			_ = _nwElement0_7
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(11))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_7, 0)
			(_nw38).ArraySet1(_218_v0, 1)
			(_nw38).ArraySet1((_dafny.SetOf(_dafny.IntOfUint32((_223_v3).Cardinality()))).IsDisjointFrom(_222_v2), 2)
			(_nw38).ArraySet1(true, 3)
			(_nw38).ArraySet1(((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)).Cmp((_this).Fm25(false, _218_v0, globalState)) == 0, 4)
			(_nw38).ArraySet1(_218_v0, 5)
			(_nw38).ArraySet1(_218_v0, 6)
			(_nw38).ArraySet1(_218_v0, 7)
			(_nw38).ArraySet1(_dafny.Companion_Sequence_.Contains(_292_v42, _300_v50), 8)
			(_nw38).ArraySet1((!(_218_v0)) && (_218_v0), 9)
			(_nw38).ArraySet1((_dafny.MultiSetOf(_218_v0)).IsProperSubsetOf(_dafny.MultiSetOf(_218_v0, _218_v0)), 10)
			_301_v51 = _nw38
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_301_v51), 0))
			_ = _index25
			(_301_v51).ArraySet1((func() bool {
				if (_284_v36).Contains((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)) {
					return (_284_v36).Get((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)).(bool)
				}
				return true
			})(), (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_301_v51), 0))
			_ = _index26
			var _rhs49 _dafny.Map = _297_v47
			_ = _rhs49
			var _rhs50 _dafny.MultiSet = _298_v48
			_ = _rhs50
			var _rhs51 bool = !(_218_v0) || (_218_v0)
			_ = _rhs51
			var _lhs28 _dafny.Array = _301_v51
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_301_v51), 0))
			_ = _lhs29
			_297_v47 = _rhs49
			_298_v48 = _rhs50
			(_lhs28).ArraySet1(_rhs51, (_lhs29).Int())
			var _302_v52 _dafny.CodePoint
			_ = _302_v52
			_302_v52 = _dafny.CodePoint('a')
			var _303_v53 _dafny.Sequence
			_ = _303_v53
			_303_v53 = _dafny.UnicodeSeqOfUtf8Bytes("uljkewecy")
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_303_v53, _303_v53), _302_v52) {
				var _304_v54 _dafny.Array
				_ = _304_v54
				var _nwElement0_8 bool = (func() bool {
					if (_284_v36).Contains(p1) {
						return (_284_v36).Get(p1).(bool)
					}
					return _218_v0
				})()
				_ = _nwElement0_8
				var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(5))
				_ = _nw39
				(_nw39).ArraySet1(_nwElement0_8, 0)
				(_nw39).ArraySet1(_218_v0, 1)
				(_nw39).ArraySet1(_218_v0, 2)
				(_nw39).ArraySet1(_218_v0, 3)
				(_nw39).ArraySet1(_218_v0, 4)
				_304_v54 = _nw39
				var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _index27
				(_304_v54).ArraySet1(false, (_index27).Int())
				var _305_v55 _dafny.Array
				_ = _305_v55
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_7
				var _nw40 _dafny.Array
				_ = _nw40
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw40 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Map = (func(_306_v0 bool, _307_p1 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_308_i4 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_306_v0), _307_p1)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_306_v0, _307_p1))
						}
					})(_218_v0, p1)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw40 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw40).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw40).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_305_v55 = _nw40
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _index28
				var _rhs52 bool = _218_v0
				_ = _rhs52
				var _rhs53 _dafny.Array = _305_v55
				_ = _rhs53
				var _lhs30 _dafny.Array = _304_v54
				_ = _lhs30
				var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _lhs31
				(_lhs30).ArraySet1(_rhs52, (_lhs31).Int())
				_305_v55 = _rhs53
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _index29
				var _rhs54 _dafny.CodePoint = _302_v52
				_ = _rhs54
				var _rhs55 bool = !(!(_218_v0))
				_ = _rhs55
				var _lhs32 _dafny.Array = _304_v54
				_ = _lhs32
				var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _lhs33
				_302_v52 = _rhs54
				(_lhs32).ArraySet1(_rhs55, (_lhs33).Int())
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _index30
				var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _index31
				var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _index32
				var _rhs56 bool = (_304_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))).Int()).(bool)
				_ = _rhs56
				var _rhs57 _dafny.Int = _283_cf7
				_ = _rhs57
				var _rhs58 bool = !(!(true) || (_218_v0)) || (_218_v0)
				_ = _rhs58
				var _rhs59 bool = Companion_Default___.Fm26(globalState)
				_ = _rhs59
				var _rhs60 bool = (_304_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))).Int()).(bool)
				_ = _rhs60
				var _lhs34 _dafny.Array = _304_v54
				_ = _lhs34
				var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _lhs35
				var _lhs36 *GlobalState = globalState
				_ = _lhs36
				var _lhs37 _dafny.Array = _304_v54
				_ = _lhs37
				var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _lhs38
				var _lhs39 _dafny.Array = _304_v54
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))
				_ = _lhs40
				(_lhs34).ArraySet1(_rhs56, (_lhs35).Int())
				_lhs36.F7 = _rhs57
				(_lhs37).ArraySet1(_rhs58, (_lhs38).Int())
				_218_v0 = _rhs59
				(_lhs39).ArraySet1(_rhs60, (_lhs40).Int())
				var _309_v56 _dafny.Map
				_ = _309_v56
				_309_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SeqOf((_304_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))).Int()).(bool)))
				r0 = (func() _dafny.CodePoint {
					if (_286_v38).Contains((_dafny.IntOfInt64(879)).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_309_v56).Contains((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)) {
							return (_309_v56).Get((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
						}
						return _285_v37
					})()).Cardinality()))) {
						return (_286_v38).Get((_dafny.IntOfInt64(879)).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_309_v56).Contains((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)) {
								return (_309_v56).Get((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
							}
							return _285_v37
						})()).Cardinality()))).(_dafny.CodePoint)
					}
					return _302_v52
				})()
				var _310_v57 _dafny.Set
				_ = _310_v57
				_310_v57 = _dafny.SetOf(_218_v0)
				var _311_v58 D7
				_ = _311_v58
				_311_v58 = Companion_D7_.Create_DC22_(_218_v0, _283_cf7, (_285_v37).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_285_v37).Cardinality()))).Uint32()).(bool), (_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int))
				var _312_v59 _dafny.Sequence
				_ = _312_v59
				_312_v59 = _dafny.SeqOf(_311_v58, _311_v58)
				_303_v53 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_303_v53, (Companion_Default___.SafeIndex((_310_v57).Cardinality(), _dafny.IntOfUint32((_303_v53).Cardinality()))).Uint32(), _302_v52), Companion_Default___.Fm30(_dafny.IntOfUint32((_312_v59).Cardinality()), Companion_Default___.Fm26(globalState), func() _dafny.Set {
					var _coll21 = _dafny.NewBuilder()
					_ = _coll21
					for _iter21 := _dafny.Iterate((_223_v3).Elements()); ; {
						_compr_21, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _313_v60 _dafny.Int
						_313_v60 = interface{}(_compr_21).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_223_v3, _313_v60) {
							_coll21.Add(Companion_Default___.SafeDivisionInt(_313_v60, _dafny.IntOfInt64(22)))
						}
					}
					return _coll21.ToSet()
				}(), (_304_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_304_v54), 0))).Int()).(bool), globalState)), _303_v53)
			} else {
				var _314_v61 _dafny.Array
				_ = _314_v61
				var _nwElement0_9 bool = _218_v0
				_ = _nwElement0_9
				var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(12))
				_ = _nw41
				(_nw41).ArraySet1(_nwElement0_9, 0)
				(_nw41).ArraySet1(true, 1)
				(_nw41).ArraySet1(_218_v0, 2)
				(_nw41).ArraySet1(_218_v0, 3)
				(_nw41).ArraySet1(_218_v0, 4)
				(_nw41).ArraySet1(_218_v0, 5)
				(_nw41).ArraySet1(_218_v0, 6)
				(_nw41).ArraySet1(_218_v0, 7)
				(_nw41).ArraySet1(_218_v0, 8)
				(_nw41).ArraySet1(_218_v0, 9)
				(_nw41).ArraySet1(true, 10)
				(_nw41).ArraySet1(_218_v0, 11)
				_314_v61 = _nw41
				_218_v0 = (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_314_v61, _dafny.IntOfInt64(292))).Cardinality()), _283_cf7)).Cmp((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)) <= 0
				(globalState).F7 = (_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)
				var _315_v62 _dafny.Array
				_ = _315_v62
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_8
				var _nw42 _dafny.Array
				_ = _nw42
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw42 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) D4 = (func(_316_cf7 _dafny.Int) func(_dafny.Int) D4 {
						return func(_317_i5 _dafny.Int) D4 {
							return Companion_D4_.Create_DC12_(_316_cf7)
						}
					})(_283_cf7)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw42 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw42).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw42).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_315_v62 = _nw42
				var _318_v63 D4
				_ = _318_v63
				_318_v63 = Companion_D4_.Create_DC12_(_283_cf7)
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_315_v62), 0))
				_ = _index33
				(_315_v62).ArraySet1(_318_v63, (_index33).Int())
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_315_v62), 0))
				_ = _index34
				(_315_v62).ArraySet1(_318_v63, (_index34).Int())
				var _319_v64 _dafny.Map
				_ = _319_v64
				_319_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_225_v5).Contains(_218_v0) {
						return (_225_v5).Get(_218_v0).(bool)
					}
					return _218_v0
				})(), p1)
				var _320_v65 _dafny.MultiSet
				_ = _320_v65
				_320_v65 = _dafny.MultiSetOf(_218_v0, _218_v0)
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_314_v61), 0))
				_ = _index35
				(_314_v61).ArraySet1(((_320_v65).Cardinality()).Cmp(_dafny.IntOfInt64(-438)) == 0, (_index35).Int())
				var _321_v66 D7
				_ = _321_v66
				_321_v66 = Companion_D7_.Create_DC21_(_dafny.IntOfInt64(376), _283_cf7, _218_v0, !(_218_v0), _218_v0)
				var _322_v67 D6
				_ = _322_v67
				_322_v67 = Companion_D6_.Create_DC17_((_321_v66).Dtor_cf33(), _218_v0, _218_v0)
				var _323_v68 _dafny.MultiSet
				_ = _323_v68
				_323_v68 = _dafny.MultiSetOf(_322_v67)
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_314_v61), 0))
				_ = _index36
				var _rhs61 _dafny.Map = (_319_v64).Merge(_319_v64)
				_ = _rhs61
				var _rhs62 bool = _218_v0
				_ = _rhs62
				var _rhs63 bool = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-919))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_324_v0 bool) func(_dafny.Int) _dafny.Int {
					return func(_325_i6 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_324_v0, _324_v0)).Cardinality()
					}
				})(_218_v0))), (_dafny.IntOfInt64(105)).Minus((_323_v68).Cardinality()))
				_ = _rhs63
				var _rhs64 _dafny.Int = _dafny.IntOfInt64(-271)
				_ = _rhs64
				var _rhs65 bool = !_dafny.Companion_Sequence_.Equal(_223_v3, _223_v3)
				_ = _rhs65
				var _lhs41 _dafny.Array = _314_v61
				_ = _lhs41
				var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_314_v61), 0))
				_ = _lhs42
				var _lhs43 *GlobalState = globalState
				_ = _lhs43
				_319_v64 = _rhs61
				(_lhs41).ArraySet1(_rhs62, (_lhs42).Int())
				_218_v0 = _rhs63
				_lhs43.F15 = _rhs64
				_218_v0 = _rhs65
				var _326_v69 D6
				_ = _326_v69
				_326_v69 = Companion_D6_.Create_DC16_(_223_v3)
				var _327_v70 D6
				_ = _327_v70
				_327_v70 = Companion_D6_.Create_DC18_(_326_v69)
				var _328_v71 _dafny.MultiSet
				_ = _328_v71
				_328_v71 = _dafny.MultiSetOf(_327_v70)
				var _329_v72 _dafny.Sequence
				_ = _329_v72
				_329_v72 = _dafny.SeqOf(Companion_D6_.Create_DC18_(_326_v69))
				_328_v71 = ((_328_v71).Difference(_328_v71)).Intersection(_dafny.MultiSetFromSeq(_329_v72))
			}
			r1 = _283_cf7
		} else if _source6.Is_DC7() {
			var _330___mcc_h12 bool = _source6.Get_().(D2_DC7).Cf8
			_ = _330___mcc_h12
			var _331___mcc_h13 _dafny.MultiSet = _source6.Get_().(D2_DC7).Cf9
			_ = _331___mcc_h13
			var _332___mcc_h14 _dafny.Array = _source6.Get_().(D2_DC7).Cf10
			_ = _332___mcc_h14
			var _333_cf10 _dafny.Array = _332___mcc_h14
			_ = _333_cf10
			var _334_cf9 _dafny.MultiSet = _331___mcc_h13
			_ = _334_cf9
			var _335_cf8 bool = _330___mcc_h12
			_ = _335_cf8
			var _336_v75 _dafny.Array
			_ = _336_v75
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_9
			var _nw43 _dafny.Array
			_ = _nw43
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw43 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Set = (func(_337_p1 _dafny.Int, _338_cf8 bool) func(_dafny.Int) _dafny.Set {
					return func(_339_i7 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_dafny.MultiSetOf(_337_p1, _337_p1), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_338_cf8, _338_cf8)).Cardinality())))).Intersection(func() _dafny.Set {
							var _coll22 = _dafny.NewBuilder()
							_ = _coll22
							for _iter22 := _dafny.Iterate((func() _dafny.Set {
								var _coll23 = _dafny.NewBuilder()
								_ = _coll23
								for _iter23 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(130)))).Elements()); ; {
									_compr_23, _ok23 := _iter23()
									if !_ok23 {
										break
									}
									var _340_v73 _dafny.MultiSet
									_340_v73 = interface{}(_compr_23).(_dafny.MultiSet)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(130))), _340_v73) {
										_coll23.Add(_340_v73)
									}
								}
								return _coll23.ToSet()
							}()).Elements()); ; {
								_compr_22, _ok22 := _iter22()
								if !_ok22 {
									break
								}
								var _341_v74 _dafny.MultiSet
								_341_v74 = interface{}(_compr_22).(_dafny.MultiSet)
								if (func() _dafny.Set {
									var _coll24 = _dafny.NewBuilder()
									_ = _coll24
									for _iter24 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(130)))).Elements()); ; {
										_compr_24, _ok24 := _iter24()
										if !_ok24 {
											break
										}
										var _342_v73 _dafny.MultiSet
										_342_v73 = interface{}(_compr_24).(_dafny.MultiSet)
										if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(130))), _342_v73) {
											_coll24.Add(_342_v73)
										}
									}
									return _coll24.ToSet()
								}()).Contains(_341_v74) {
									_coll22.Add(_341_v74)
								}
							}
							return _coll22.ToSet()
						}())
					}
				})(p1, _335_cf8)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw43 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw43).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw43).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_336_v75 = _nw43
			var _343_v76 _dafny.MultiSet
			_ = _343_v76
			_343_v76 = _dafny.MultiSetOf(p1, (_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int))
			var _344_v77 _dafny.Set
			_ = _344_v77
			_344_v77 = _dafny.SetOf(_343_v76, _343_v76, Companion_Default___.Fm31((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int), (_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int), globalState))
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_336_v75), 0))
			_ = _index37
			(_336_v75).ArraySet1(_344_v77, (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_336_v75), 0))
			_ = _index38
			(_336_v75).ArraySet1((_344_v77).Intersection(_344_v77), (_index38).Int())
			(globalState).F13 = p1
			var _345_v78 _dafny.Sequence
			_ = _345_v78
			_345_v78 = _dafny.UnicodeSeqOfUtf8Bytes("vs")
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
			_ = _index39
			(_219_v1).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_345_v78, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pq"), _345_v78))).Cardinality()), (_index39).Int())
			if _218_v0 {
				(globalState).F7 = (_dafny.Zero).Minus(p1)
				var _346_v79 *C0
				_ = _346_v79
				var _nw44 *C0 = New_C0_()
				_ = _nw44
				_nw44.Ctor__()
				_346_v79 = _nw44
				(globalState).F15 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_345_v78, _dafny.Companion_Sequence_.Concatenate(_345_v78, _345_v78))).Cardinality())
				(globalState).F13 = (_dafny.Zero).Minus(((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)).Plus((_this).Fm25(_335_cf8, _335_cf8, globalState)))
				var _347_v80 _dafny.Map
				_ = _347_v80
				_347_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v35, _218_v0)
				var _348_v81 _dafny.Map
				_ = _348_v81
				_348_v81 = _347_v80
				_347_v80 = (_348_v81)
			} else {
				_219_v1 = _333_cf10
				var _349_v82 *C0
				_ = _349_v82
				var _nw45 *C0 = New_C0_()
				_ = _nw45
				_nw45.Ctor__()
				_349_v82 = _nw45
				var _350_v83 _dafny.CodePoint
				_ = _350_v83
				_350_v83 = _dafny.CodePoint('w')
				r1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_345_v78, _dafny.Companion_Sequence_.Update(_345_v78, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.IntOfUint32((_345_v78).Cardinality()))).Uint32(), _350_v83))).Cardinality())
				_335_cf8 = _335_cf8
				var _351_v84 *C0
				_ = _351_v84
				var _nw46 *C0 = New_C0_()
				_ = _nw46
				_nw46.Ctor__()
				_351_v84 = _nw46
			}
		} else if _source6.Is_DC5() {
			var _352___mcc_h15 _dafny.Array = _source6.Get_().(D2_DC5).Cf6
			_ = _352___mcc_h15
			var _353_cf6 _dafny.Array = _352___mcc_h15
			_ = _353_cf6
			var _354_v85 _dafny.Sequence
			_ = _354_v85
			_354_v85 = _dafny.SeqOf((_218_v0) == (_218_v0), Companion_Default___.Fm26(globalState), !(_218_v0) || (_218_v0), true, _218_v0)
			_354_v85 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_354_v85, _354_v85), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_354_v85, _354_v85)).Cardinality()))).Uint32(), _218_v0)
			_218_v0 = (p1).Cmp((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)) > 0
			_281_v35 = _281_v35
			_218_v0 = (Companion_Default___.SafeDivisionInt(p1, p1)).Cmp((_dafny.IntOfInt64(114)).Times(p1)) != 0
		} else {
			var _355___mcc_h16 D2 = _source6.Get_().(D2_DC8).Cf11
			_ = _355___mcc_h16
			var _356_cf11 D2 = _355___mcc_h16
			_ = _356_cf11
			_218_v0 = Companion_Default___.Fm26(globalState)
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))
			_ = _index40
			(_219_v1).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)), (_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int))), (_index40).Int())
			var _357_v86 _dafny.Sequence
			_ = _357_v86
			_357_v86 = _dafny.UnicodeSeqOfUtf8Bytes("wxgofwv")
			var _358_v87 _dafny.CodePoint
			_ = _358_v87
			_358_v87 = _dafny.CodePoint('e')
			var _359_v88 _dafny.Map
			_ = _359_v88
			_359_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_218_v0, _358_v87)
			var _360_v89 _dafny.Sequence
			_ = _360_v89
			_360_v89 = _dafny.SeqOf(_218_v0)
			var _361_v90 _dafny.Map
			_ = _361_v90
			_361_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_223_v3).Cardinality()), _357_v86)
			var _362_v91 _dafny.Array
			_ = _362_v91
			var _nwElement0_10 _dafny.Sequence = _357_v86
			_ = _nwElement0_10
			var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(12))
			_ = _nw47
			(_nw47).ArraySet1(_nwElement0_10, 0)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_357_v86, _357_v86), 1)
			(_nw47).ArraySet1(_357_v86, 2)
			(_nw47).ArraySet1(_357_v86, 3)
			(_nw47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vsd"), Companion_Default___.Fm30((_222_v2).Cardinality(), _218_v0, _dafny.SetOf((_359_v88).Cardinality(), (_dafny.MultiSetFromSeq(_360_v89)).Cardinality()), _218_v0, globalState)), 4)
			(_nw47).ArraySet1(_357_v86, 5)
			(_nw47).ArraySet1((func() _dafny.Sequence {
				if (_361_v90).Contains((_222_v2).Cardinality()) {
					return (_361_v90).Get((_222_v2).Cardinality()).(_dafny.Sequence)
				}
				return _357_v86
			})(), 6)
			(_nw47).ArraySet1(_357_v86, 7)
			(_nw47).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("otlcts"), 8)
			(_nw47).ArraySet1(_357_v86, 9)
			(_nw47).ArraySet1(_357_v86, 10)
			(_nw47).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ke"), 11)
			_362_v91 = _nw47
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_362_v91), 0))
			_ = _index41
			(_362_v91).ArraySet1(_357_v86, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_362_v91), 0))
			_ = _index42
			var _rhs66 _dafny.Sequence = _357_v86
			_ = _rhs66
			var _rhs67 _dafny.Int = p1
			_ = _rhs67
			var _lhs44 _dafny.Array = _362_v91
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_362_v91), 0))
			_ = _lhs45
			var _lhs46 *GlobalState = globalState
			_ = _lhs46
			(_lhs44).ArraySet1(_rhs66, (_lhs45).Int())
			_lhs46.F15 = _rhs67
			var _363_v92 D7
			_ = _363_v92
			_363_v92 = Companion_D7_.Create_DC20_(p1)
			_363_v92 = _363_v92
		}
		var _364_v93 _dafny.Array
		_ = _364_v93
		var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw48
		_364_v93 = _nw48
		var _365_v94 _dafny.Sequence
		_ = _365_v94
		_365_v94 = _dafny.SeqOf(_218_v0)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))
		_ = _index43
		(_364_v93).ArraySet1((_365_v94).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_365_v94).Cardinality()))).Uint32()).(bool), (_index43).Int())
		var _366_v95 _dafny.MultiSet
		_ = _366_v95
		_366_v95 = _dafny.MultiSetOf(_218_v0)
		var _367_v96 _dafny.Set
		_ = _367_v96
		_367_v96 = _dafny.SetOf(_366_v95)
		var _368_v98 _dafny.Map
		_ = _368_v98
		_368_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() _dafny.Int {
			if (_366_v95).Contains(_218_v0) {
				return (_366_v95).Multiplicity(_218_v0)
			}
			return _dafny.IntOfInt64(909)
		})())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))
		_ = _index44
		var _rhs68 bool = _218_v0
		_ = _rhs68
		var _rhs69 bool = (_367_v96).IsDisjointFrom((func() _dafny.Set {
			var _coll25 = _dafny.NewBuilder()
			_ = _coll25
			for _iter25 := _dafny.Iterate((_dafny.SeqOf(_366_v95)).Elements()); ; {
				_compr_25, _ok25 := _iter25()
				if !_ok25 {
					break
				}
				var _369_v97 _dafny.MultiSet
				_369_v97 = interface{}(_compr_25).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_366_v95), _369_v97) {
					_coll25.Add(_369_v97)
				}
			}
			return _coll25.ToSet()
		}()).Difference(_367_v96))
		_ = _rhs69
		var _rhs70 _dafny.Int = (_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)
		_ = _rhs70
		var _rhs71 _dafny.Int = (p1).Minus(((_368_v98).Cardinality()).Plus(_dafny.IntOfInt64(188)))
		_ = _rhs71
		var _rhs72 bool = _218_v0
		_ = _rhs72
		var _lhs47 *GlobalState = globalState
		_ = _lhs47
		var _lhs48 *GlobalState = globalState
		_ = _lhs48
		var _lhs49 _dafny.Array = _364_v93
		_ = _lhs49
		var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))
		_ = _lhs50
		_218_v0 = _rhs68
		_218_v0 = _rhs69
		_lhs47.F15 = _rhs70
		_lhs48.F15 = _rhs71
		(_lhs49).ArraySet1(_rhs72, (_lhs50).Int())
		var _hi1 _dafny.Int = (_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int)
		_ = _hi1
		for _370_i8 := (_this).Fm25(_218_v0, _218_v0, globalState); _370_i8.Cmp(_hi1) < 0; _370_i8 = _370_i8.Plus(_dafny.One) {
			_223_v3 = _dafny.Companion_Sequence_.Concatenate(_223_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(612))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_371_i9 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(574)
			})))
			(globalState).F15 = (func() _dafny.Int {
				if (_368_v98).Contains(((_222_v2).Intersection(_222_v2)).Cardinality()) {
					return (_368_v98).Get(((_222_v2).Intersection(_222_v2)).Cardinality()).(_dafny.Int)
				}
				return ((_223_v3).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_223_v3).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_370_i8)
			})()
			var _372_v99 _dafny.Set
			_ = _372_v99
			_372_v99 = _dafny.SetOf(_218_v0, _218_v0, _218_v0, !((_364_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))).Int()).(bool)), true)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))
			_ = _index45
			var _rhs73 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_365_v94, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_365_v94).Cardinality()))).Uint32(), (_364_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))).Int()).(bool)), _dafny.Companion_Sequence_.Concatenate(_365_v94, _365_v94))
			_ = _rhs73
			var _rhs74 _dafny.Int = (_this).Fm25(_218_v0, (_364_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))).Int()).(bool), globalState)
			_ = _rhs74
			var _rhs75 bool = (Companion_Default___.Fm32(_dafny.IntOfUint32((_365_v94).Cardinality()), (_219_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_219_v1), 0))).Int()).(_dafny.Int), _372_v99, globalState)).Equals((Companion_Default___.Fm32(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(188))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_373_v1 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_374_i10 _dafny.Int) _dafny.Int {
					return (_373_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_373_v1), 0))).Int()).(_dafny.Int)
				}
			})(_219_v1)))).Cardinality()), _dafny.IntOfInt64(466), _372_v99, globalState)).Difference(_222_v2))
			_ = _rhs75
			var _lhs51 *GlobalState = globalState
			_ = _lhs51
			var _lhs52 _dafny.Array = _364_v93
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))
			_ = _lhs53
			_365_v94 = _rhs73
			_lhs51.F7 = _rhs74
			(_lhs52).ArraySet1(_rhs75, (_lhs53).Int())
			(globalState).F15 = Companion_Default___.SafeDivisionInt((_this).Fm25(!(!(true)), (_364_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))).Int()).(bool), globalState), _370_i8)
		}
		var _375_v100 D6
		_ = _375_v100
		_375_v100 = Companion_D6_.Create_DC17_(true, (_364_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))).Int()).(bool), !((_364_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))).Int()).(bool)))
		var _pat_let_tv1 = _364_v93
		_ = _pat_let_tv1
		var _pat_let_tv2 = _364_v93
		_ = _pat_let_tv2
		r0 = func(_source8 D6) _dafny.CodePoint {
			if _source8.Is_DC17() {
				var _376___mcc_h18 bool = _source8.Get_().(D6_DC17).Cf24
				_ = _376___mcc_h18
				var _377___mcc_h19 bool = _source8.Get_().(D6_DC17).Cf25
				_ = _377___mcc_h19
				var _378___mcc_h20 bool = _source8.Get_().(D6_DC17).Cf26
				_ = _378___mcc_h20
				var _379_cf26 bool = _378___mcc_h20
				_ = _379_cf26
				var _380_cf25 bool = _377___mcc_h19
				_ = _380_cf25
				var _381_cf24 bool = _376___mcc_h18
				_ = _381_cf24
				return _dafny.CodePoint('c')
			} else if _source8.Is_DC16() {
				var _382___mcc_h21 _dafny.Sequence = _source8.Get_().(D6_DC16).Cf23
				_ = _382___mcc_h21
				var _383_cf23 _dafny.Sequence = _382___mcc_h21
				_ = _383_cf23
				if (_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_pat_let_tv1), 0))).Int()).(bool) {
					return _dafny.CodePoint('b')
				} else {
					return _dafny.CodePoint('y')
				}
			} else {
				var _384___mcc_h22 D6 = _source8.Get_().(D6_DC18).Cf27
				_ = _384___mcc_h22
				var _385_cf27 D6 = _384___mcc_h22
				_ = _385_cf27
				return _dafny.CodePoint('x')
			}
		}(_375_v100)
		r1 = p1
		var _386_v101 _dafny.Sequence
		_ = _386_v101
		_386_v101 = _dafny.SeqOf(_364_v93)
		r2 = (func() D7 {
			if !((_364_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_364_v93), 0))).Int()).(bool)) {
				return Companion_D7_.Create_DC19_(_386_v101)
			}
			return p0
		})()
		var _387_v102 _dafny.Sequence
		_ = _387_v102
		_387_v102 = _dafny.UnicodeSeqOfUtf8Bytes("lpsnk")
		var _388_v103 D1
		_ = _388_v103
		_388_v103 = Companion_D1_.Create_DC3_(_387_v102, _218_v0)
		r3 = (func() D1 {
			if (_218_v0) && (false) {
				return Companion_D1_.Create_DC3_((_388_v103).Dtor_cf3(), _218_v0)
			}
			return _388_v103
		})()
		return r0, r1, r2, r3
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	dummy byte
}

func New_C2_() *C2 {
	_this := C2{}

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

func (_this *C2) Ctor__() {
	{
	}
}
func (_this *C2) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(117))).Difference(_dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality(), _dafny.IntOfInt64(679))).Cardinality()))), _dafny.SetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("yjblw"), (Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("ywhli"), false)).Dtor_cf4())).Cardinality()), (_dafny.Zero).Minus((_dafny.SetOf((Companion_D7_.Create_DC20_(_dafny.IntOfInt64(898))).Dtor_cf29(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-596)), _dafny.IntOfInt64(483))).Cardinality())).Cardinality()), _dafny.IntOfInt64(46))).Cardinality())), func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter26 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(258), (_dafny.SetOf(false, false, true, true)).Cardinality())).Elements()); ; {
				_compr_26, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _389_v0 _dafny.Int
				_389_v0 = interface{}(_compr_26).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(258), (_dafny.SetOf(false, false, true, true)).Cardinality())).Contains(_389_v0) {
					_coll26.Add(Companion_Default___.SafeModuloInt(_389_v0, (func() _dafny.Set {
						var _coll27 = _dafny.NewBuilder()
						_ = _coll27
						for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(568), _dafny.IntOfInt64(301))); ; {
							_compr_27, _ok27 := _iter27()
							if !_ok27 {
								break
							}
							var _390_v1 _dafny.Int
							_390_v1 = interface{}(_compr_27).(_dafny.Int)
							if ((_dafny.IntOfInt64(568)).Cmp(_390_v1) <= 0) && ((_390_v1).Cmp(_dafny.IntOfInt64(301)) < 0) {
								_coll27.Add((_390_v1).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), !(true))).Cardinality()))
							}
						}
						return _coll27.ToSet()
					}()).Cardinality()))
				}
			}
			return _coll26.ToSet()
		}(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, !(false))).Cardinality()))).Difference(_dafny.SetOf(_dafny.IntOfInt64(-746))))
	}
}
func (_this *C2) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		if !(true) || (!(false)) {
			return !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC10_(_dafny.IntOfInt64(818), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fve")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aclnkmm")).Cardinality())))).Contains(Companion_D3_.Create_DC10_((func() _dafny.Map {
				var _coll28 = _dafny.NewMapBuilder()
				_ = _coll28
				for _iter28 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(813))).Elements()); ; {
					_compr_28, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _391_v0 _dafny.Int
					_391_v0 = interface{}(_compr_28).(_dafny.Int)
					if (_dafny.MultiSetOf(_dafny.IntOfInt64(813))).Contains(_391_v0) {
						_coll28.Add((_391_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(595), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('a'))).Cardinality(), true)).Cardinality())).Cardinality())).Cardinality()))).Cardinality())), true)
					}
				}
				return _coll28.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))))
		} else {
			return (_dafny.IntOfInt64(653)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(488))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_392_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(782)
			}))).Cardinality())) < 0
		}
	}
}
func (_this *C2) Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C2) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(690))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg36 _dafny.Int) interface{} {
				return coer36(arg36)
			}
		}(func(_393_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality()), _dafny.IntOfInt64(-61))
	}
}
func (_this *C2) Fm24(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(true)
	}
}
func (_this *C2) M6(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 D2, globalState *GlobalState) {
	{
		var _hi2 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p0), _dafny.IntOfInt64(693))
		_ = _hi2
		for _394_i0 := p0; _394_i0.Cmp(_hi2) < 0; _394_i0 = _394_i0.Plus(_dafny.One) {
			var _395_v0 bool
			_ = _395_v0
			_395_v0 = false
			var _396_v1 _dafny.Array
			_ = _396_v1
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw49
			_396_v1 = _nw49
			var _397_v2 _dafny.MultiSet
			_ = _397_v2
			_397_v2 = _dafny.MultiSetOf(_396_v1, _396_v1)
			_395_v0 = !((_397_v2).Intersection(_397_v2)).Contains(_396_v1)
			(globalState).F4 = _394_i0
			var _398_v3 _dafny.Array
			_ = _398_v3
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw50
			_398_v3 = _nw50
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_398_v3), 0))
			_ = _index46
			(_398_v3).ArraySet1(p0, (_index46).Int())
			var _399_v4 _dafny.Map
			_ = _399_v4
			_399_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _395_v0)
			var _400_v5 _dafny.Map
			_ = _400_v5
			_400_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_395_v0, _399_v4)
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_398_v3), 0))
			_ = _index47
			var _rhs76 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(p0, _394_i0), (_dafny.SetOf(_395_v0)).Cardinality())
			_ = _rhs76
			var _rhs77 bool = (_this).Fm1((_400_v5).Merge((_400_v5).Update(p2, _399_v4)), _394_i0, globalState)
			_ = _rhs77
			var _lhs54 _dafny.Array = _398_v3
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_398_v3), 0))
			_ = _lhs55
			(_lhs54).ArraySet1(_rhs76, (_lhs55).Int())
			_395_v0 = _rhs77
			_395_v0 = !(_395_v0)
		}
		var _401_v6 bool
		_ = _401_v6
		_401_v6 = false
		var _402_v7 _dafny.Sequence
		_ = _402_v7
		_402_v7 = _dafny.UnicodeSeqOfUtf8Bytes("qtrupg")
		_401_v6 = (Companion_Default___.SafeModuloInt(p0, p0)).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_402_v7).Cardinality()), _dafny.IntOfInt64(-925))) >= 0
		var _403_i1 _dafny.Int
		_ = _403_i1
		_403_i1 = _dafny.Zero
		{
			for _401_v6 {
				{
					if (_403_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_403_i1 = (_403_i1).Plus(_dafny.One)
					_401_v6 = _401_v6
					var _404_v8 _dafny.Set
					_ = _404_v8
					_404_v8 = _dafny.SetOf(_401_v6)
					var _405_v9 _dafny.Map
					_ = _405_v9
					_405_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.MultiSetOf(_404_v8, _404_v8))
					var _406_v10 _dafny.MultiSet
					_ = _406_v10
					_406_v10 = _dafny.MultiSetOf(_404_v8)
					_405_v9 = (_405_v9).Update(p0, _406_v10)
					var _407_v11 *C0
					_ = _407_v11
					var _nw51 *C0 = New_C0_()
					_ = _nw51
					_nw51.Ctor__()
					_407_v11 = _nw51
					var _408_v12 _dafny.Array
					_ = _408_v12
					var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(9))
					_ = _nw52
					_408_v12 = _nw52
					var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_408_v12), 0))
					_ = _index48
					(_408_v12).ArraySet1(_404_v8, (_index48).Int())
					var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_408_v12), 0))
					_ = _index49
					(_408_v12).ArraySet1((_404_v8).Intersection((_404_v8).Intersection(_404_v8)), (_index49).Int())
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _409_v13 _dafny.Array
		_ = _409_v13
		var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(11))
		_ = _nw53
		_409_v13 = _nw53
		var _410_v14 _dafny.Set
		_ = _410_v14
		_410_v14 = _dafny.SetOf(true)
		var _411_v15 _dafny.Set
		_ = _411_v15
		_411_v15 = _dafny.SetOf(_410_v14, _410_v14, _410_v14)
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_409_v13), 0))
		_ = _index50
		(_409_v13).ArraySet1((func() _dafny.Set {
			if _401_v6 {
				return _411_v15
			}
			return _411_v15
		})(), (_index50).Int())
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_409_v13), 0))
		_ = _index51
		(_409_v13).ArraySet1(_411_v15, (_index51).Int())
		var _412_v16 _dafny.Array
		_ = _412_v16
		var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
		_ = _nw54
		_412_v16 = _nw54
		var _413_v17 _dafny.Set
		_ = _413_v17
		_413_v17 = _dafny.SetOf(p0, (_this).Fm2((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(p2))).Cardinality(), globalState))
		var _414_v18 _dafny.Map
		_ = _414_v18
		_414_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
		var _415_v19 _dafny.Map
		_ = _415_v19
		_415_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
		var _416_v20 _dafny.Map
		_ = _416_v20
		_416_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(756), _413_v17)
		var _417_v21 _dafny.Map
		_ = _417_v21
		_417_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_402_v7, Companion_Default___.Fm30(p0, _401_v6, _413_v17, _401_v6, globalState)), (_414_v18).Merge(Companion_Default___.Fm33(!(p2), (func() bool {
			if (_415_v19).Contains(p0) {
				return (_415_v19).Get(p0).(bool)
			}
			return _401_v6
		})(), ((func() _dafny.Set {
			if (_416_v20).Contains(p0) {
				return (_416_v20).Get(p0).(_dafny.Set)
			}
			return _413_v17
		})()).Cardinality(), globalState)))
		var _418_v22 _dafny.Array
		_ = _418_v22
		var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(2))
		_ = _nw55
		_418_v22 = _nw55
		var _419_v23 *C0
		_ = _419_v23
		var _nw56 *C0 = New_C0_()
		_ = _nw56
		_nw56.Ctor__()
		_419_v23 = _nw56
		var _420_v24 _dafny.Map
		_ = _420_v24
		_420_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v23, p0)
		var _421_v25 _dafny.MultiSet
		_ = _421_v25
		_421_v25 = _dafny.MultiSetOf(_420_v24)
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_418_v22), 0))
		_ = _index52
		(_418_v22).ArraySet1((_421_v25).Union(_421_v25), (_index52).Int())
		var _422_v27 D11
		_ = _422_v27
		_422_v27 = Companion_D11_.Create_DC28_(_413_v17)
		var _423_v28 _dafny.MultiSet
		_ = _423_v28
		_423_v28 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(954))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg37 _dafny.Int) interface{} {
				return coer37(arg37)
			}
		}(func(_424_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		})), Companion_Default___.Fm30(p0, p2, (_422_v27).Dtor_cf45(), p2, globalState), _402_v7)
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_418_v22), 0))
		_ = _index53
		var _rhs78 _dafny.Array = _412_v16
		_ = _rhs78
		var _rhs79 _dafny.Map = (func() _dafny.Map {
			var _coll29 = _dafny.NewMapBuilder()
			_ = _coll29
			for _iter29 := _dafny.Iterate((_423_v28).Elements()); ; {
				_compr_29, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _425_v26 _dafny.Sequence
				_425_v26 = interface{}(_compr_29).(_dafny.Sequence)
				if (_423_v28).Contains(_425_v26) {
					_coll29.Add(_425_v26, _414_v18)
				}
			}
			return _coll29.ToMap()
		}()).Update(_402_v7, (_414_v18).Merge(_414_v18))
		_ = _rhs79
		var _rhs80 _dafny.MultiSet = _421_v25
		_ = _rhs80
		var _rhs81 _dafny.Int = p0
		_ = _rhs81
		var _rhs82 bool = _401_v6
		_ = _rhs82
		var _lhs56 _dafny.Array = _418_v22
		_ = _lhs56
		var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_418_v22), 0))
		_ = _lhs57
		var _lhs58 *GlobalState = globalState
		_ = _lhs58
		_412_v16 = _rhs78
		_417_v21 = _rhs79
		(_lhs56).ArraySet1(_rhs80, (_lhs57).Int())
		_lhs58.F4 = _rhs81
		_401_v6 = _rhs82
		var _426_v29 _dafny.Array
		_ = _426_v29
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_10
		var _nw57 _dafny.Array
		_ = _nw57
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw57 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.CodePoint = (func(_427_v6 bool) func(_dafny.Int) _dafny.CodePoint {
				return func(_428_i3 _dafny.Int) _dafny.CodePoint {
					return (func() _dafny.CodePoint {
						if _427_v6 {
							return _dafny.CodePoint('g')
						}
						return _dafny.CodePoint('e')
					})()
				}
			})(_401_v6)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw57 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw57).ArraySet1CodePoint(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw57).ArraySet1CodePoint(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_426_v29 = _nw57
		var _429_v30 _dafny.CodePoint
		_ = _429_v30
		_429_v30 = _dafny.CodePoint('l')
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_426_v29), 0))
		_ = _index54
		(_426_v29).ArraySet1CodePoint(_429_v30, (_index54).Int())
		var _430_v31 _dafny.Array
		_ = _430_v31
		var _nwElement0_11 bool = p2
		_ = _nwElement0_11
		var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.One)
		_ = _nw58
		(_nw58).ArraySet1(_nwElement0_11, 0)
		_430_v31 = _nw58
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_430_v31), 0))
		_ = _index55
		(_430_v31).ArraySet1(_401_v6, (_index55).Int())
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_426_v29), 0))
		_ = _index56
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_430_v31), 0))
		_ = _index57
		var _rhs83 _dafny.CodePoint = _429_v30
		_ = _rhs83
		var _rhs84 D11 = Companion_D11_.Create_DC28_(_413_v17)
		_ = _rhs84
		var _rhs85 bool = p2
		_ = _rhs85
		var _rhs86 _dafny.Int = p0
		_ = _rhs86
		var _rhs87 _dafny.Int = p0
		_ = _rhs87
		var _lhs59 _dafny.Array = _426_v29
		_ = _lhs59
		var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_426_v29), 0))
		_ = _lhs60
		var _lhs61 _dafny.Array = _430_v31
		_ = _lhs61
		var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_430_v31), 0))
		_ = _lhs62
		var _lhs63 *GlobalState = globalState
		_ = _lhs63
		var _lhs64 *GlobalState = globalState
		_ = _lhs64
		(_lhs59).ArraySet1CodePoint(_rhs83, (_lhs60).Int())
		_422_v27 = _rhs84
		(_lhs61).ArraySet1(_rhs85, (_lhs62).Int())
		_lhs63.F7 = _rhs86
		_lhs64.F13 = _rhs87
	}
}
func (_this *C2) M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _431_v0 _dafny.Sequence
		_ = _431_v0
		_431_v0 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-77))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}(func(_432_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})))
		var _433_v1 D4
		_ = _433_v1
		_433_v1 = Companion_D4_.Create_DC11_(_dafny.Companion_Sequence_.Concatenate(_431_v0, _dafny.SeqOf(p0)))
		var _source9 D4 = _433_v1
		_ = _source9
		if _source9.Is_DC12() {
			var _434___mcc_h0 _dafny.Int = _source9.Get_().(D4_DC12).Cf16
			_ = _434___mcc_h0
			var _435_cf16 _dafny.Int = _434___mcc_h0
			_ = _435_cf16
			var _436_v2 _dafny.Array
			_ = _436_v2
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_11
			var _nw59 _dafny.Array
			_ = _nw59
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw59 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Int = func(_437_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_437_i1, _dafny.IntOfInt64(344))
				}
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw59 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw59).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw59).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_436_v2 = _nw59
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))
			_ = _index58
			(_436_v2).ArraySet1(_435_cf16, (_index58).Int())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))
			_ = _index59
			(_436_v2).ArraySet1(_435_cf16, (_index59).Int())
			var _438_v3 bool
			_ = _438_v3
			_438_v3 = false
			if _438_v3 {
				(globalState).F4 = (_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int)
				var _439_v4 _dafny.Map
				_ = _439_v4
				_439_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int), (_435_cf16).Minus(_435_cf16))
				_439_v4 = _439_v4
				var _440_v5 _dafny.Array
				_ = _440_v5
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_12
				var _nw60 _dafny.Array
				_ = _nw60
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw60 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) D3 = func(_441_i2 _dafny.Int) D3 {
						return Companion_D3_.Create_DC9_(_dafny.CodePoint('v'))
					}
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw60 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw60).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw60).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_440_v5 = _nw60
				var _442_v6 _dafny.CodePoint
				_ = _442_v6
				_442_v6 = _dafny.CodePoint('m')
				var _443_v7 D3
				_ = _443_v7
				_443_v7 = Companion_D3_.Create_DC9_(_442_v6)
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_440_v5), 0))
				_ = _index60
				(_440_v5).ArraySet1((func() D3 {
					if _438_v3 {
						return _443_v7
					}
					return _443_v7
				})(), (_index60).Int())
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_440_v5), 0))
				_ = _index61
				(_440_v5).ArraySet1(_443_v7, (_index61).Int())
				var _rhs88 _dafny.Array = _436_v2
				_ = _rhs88
				var _rhs89 _dafny.Int = Companion_Default___.SafeDivisionInt((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int), (_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int))
				_ = _rhs89
				var _lhs65 *GlobalState = globalState
				_ = _lhs65
				_436_v2 = _rhs88
				_lhs65.F7 = _rhs89
				_438_v3 = !(!(!(_438_v3)))
			} else {
				(globalState).F13 = (_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int)
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))
				_ = _index62
				(_436_v2).ArraySet1((_dafny.Zero).Minus((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int)), (_index62).Int())
				(globalState).F15 = (_dafny.Zero).Minus((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int))
				var _444_v8 D3
				_ = _444_v8
				_444_v8 = Companion_D3_.Create_DC10_((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_435_cf16))
				var _445_v9 _dafny.Map
				_ = _445_v9
				_445_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v8, _438_v3)
				var _446_v10 _dafny.MultiSet
				_ = _446_v10
				_446_v10 = _dafny.MultiSetOf((_445_v9).Cardinality())
				var _447_v11 _dafny.Sequence
				_ = _447_v11
				_447_v11 = _dafny.SeqOf(_dafny.IntOfInt64(720))
				var _448_v12 D1
				_ = _448_v12
				_448_v12 = Companion_D1_.Create_DC4_(_446_v10)
				var _449_v13 _dafny.Sequence
				_ = _449_v13
				_449_v13 = _dafny.SeqOf(_dafny.MultiSetOf((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int), _435_cf16))
				var _450_v14 _dafny.Map
				_ = _450_v14
				_450_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_446_v10).Cardinality(), _438_v3)
				var _451_v15 _dafny.Array
				_ = _451_v15
				var _nwElement0_12 _dafny.MultiSet = (_dafny.MultiSetOf((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int))).Difference(_446_v10)
				_ = _nwElement0_12
				var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(23))
				_ = _nw61
				(_nw61).ArraySet1(_nwElement0_12, 0)
				(_nw61).ArraySet1(_446_v10, 1)
				(_nw61).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_447_v11, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(222))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_452_cf16 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_453_i3 _dafny.Int) _dafny.Int {
						return (_dafny.SetOf(_452_cf16)).Cardinality()
					}
				})(_435_cf16))))), 2)
				(_nw61).ArraySet1(_446_v10, 3)
				(_nw61).ArraySet1(_446_v10, 4)
				(_nw61).ArraySet1(_446_v10, 5)
				(_nw61).ArraySet1(_446_v10, 6)
				(_nw61).ArraySet1((_448_v12).Dtor_cf5(), 7)
				(_nw61).ArraySet1(_446_v10, 8)
				(_nw61).ArraySet1(Companion_Default___.Fm31(_435_cf16, (_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int), globalState), 9)
				(_nw61).ArraySet1(_446_v10, 10)
				(_nw61).ArraySet1((_449_v13).Select((Companion_Default___.SafeIndex((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_449_v13).Cardinality()))).Uint32()).(_dafny.MultiSet), 11)
				(_nw61).ArraySet1(_446_v10, 12)
				(_nw61).ArraySet1(_446_v10, 13)
				(_nw61).ArraySet1(_dafny.MultiSetOf((_450_v14).Cardinality(), (_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int)), 14)
				(_nw61).ArraySet1((_446_v10).Difference(_446_v10), 15)
				(_nw61).ArraySet1(_446_v10, 16)
				(_nw61).ArraySet1(_dafny.MultiSetOf((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int)), 17)
				(_nw61).ArraySet1(_dafny.MultiSetFromSeq(_447_v11), 18)
				(_nw61).ArraySet1(_446_v10, 19)
				(_nw61).ArraySet1(_446_v10, 20)
				(_nw61).ArraySet1(_446_v10, 21)
				(_nw61).ArraySet1(_dafny.MultiSetOf((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality())), _dafny.IntOfUint32((_447_v11).Cardinality()), _dafny.IntOfInt64(526)), 22)
				_451_v15 = _nw61
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_451_v15), 0))
				_ = _index63
				(_451_v15).ArraySet1((_446_v10).Difference(_446_v10), (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_451_v15), 0))
				_ = _index64
				var _rhs90 D3 = _444_v8
				_ = _rhs90
				var _rhs91 bool = _438_v3
				_ = _rhs91
				var _rhs92 _dafny.MultiSet = _446_v10
				_ = _rhs92
				var _lhs66 _dafny.Array = _451_v15
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_451_v15), 0))
				_ = _lhs67
				_444_v8 = _rhs90
				_438_v3 = _rhs91
				(_lhs66).ArraySet1(_rhs92, (_lhs67).Int())
				(_this).M12(_438_v3, _438_v3, _436_v2, (Companion_D7_.Create_DC22_(_438_v3, (_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int), _438_v3, (_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int))).Dtor_cf35(), globalState)
			}
			var _454_v16 D8
			_ = _454_v16
			_454_v16 = Companion_D8_.Create_DC25_((_436_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_436_v2), 0))).Int()).(_dafny.Int))
			_454_v16 = _454_v16
			var _455_v17 _dafny.Sequence
			_ = _455_v17
			_455_v17 = _dafny.SeqOf((_dafny.IntOfUint32((p0).Cardinality())).Cmp(_dafny.IntOfUint32((p0).Cardinality())) <= 0)
			_438_v3 = (_455_v17).Select((Companion_Default___.SafeIndex(_435_cf16, _dafny.IntOfUint32((_455_v17).Cardinality()))).Uint32()).(bool)
		} else {
			var _456___mcc_h1 _dafny.Sequence = _source9.Get_().(D4_DC11).Cf15
			_ = _456___mcc_h1
			var _457_cf15 _dafny.Sequence = _456___mcc_h1
			_ = _457_cf15
			var _458_v18 _dafny.Int
			_ = _458_v18
			_458_v18 = _dafny.IntOfInt64(757)
			(globalState).F7 = (_this).Fm2(_458_v18, globalState)
			var _459_v19 bool
			_ = _459_v19
			_459_v19 = false
			_459_v19 = (_dafny.IntOfInt64(936)).Cmp((_dafny.Zero).Minus((_this).Fm2(_458_v18, globalState))) == 0
			var _460_v20 _dafny.Set
			_ = _460_v20
			_460_v20 = _dafny.SetOf(_459_v19)
			var _461_v21 _dafny.Sequence
			_ = _461_v21
			_461_v21 = _dafny.SeqOf(_459_v19)
			var _462_v22 D1
			_ = _462_v22
			_462_v22 = Companion_D1_.Create_DC3_(p0, !(false))
			var _463_v23 _dafny.Array
			_ = _463_v23
			var _nwElement0_13 bool = _459_v19
			_ = _nwElement0_13
			var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(10))
			_ = _nw62
			(_nw62).ArraySet1(_nwElement0_13, 0)
			(_nw62).ArraySet1(_459_v19, 1)
			(_nw62).ArraySet1(_459_v19, 2)
			(_nw62).ArraySet1((_459_v19) == (_459_v19), 3)
			(_nw62).ArraySet1((func() bool {
				if _459_v19 {
					return _459_v19
				}
				return _459_v19
			})(), 4)
			(_nw62).ArraySet1(_459_v19, 5)
			(_nw62).ArraySet1((_dafny.SetOf(_459_v19)).IsSubsetOf(_460_v20), 6)
			(_nw62).ArraySet1(_459_v19, 7)
			(_nw62).ArraySet1(!(((_461_v21).Select((Companion_Default___.SafeIndex(_458_v18, _dafny.IntOfUint32((_461_v21).Cardinality()))).Uint32()).(bool)) == (_459_v19)), 8)
			(_nw62).ArraySet1((_462_v22).Dtor_cf4(), 9)
			_463_v23 = _nw62
			_463_v23 = _463_v23
			if _459_v19 {
				var _464_v24 _dafny.Array
				_ = _464_v24
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_13
				var _nw63 _dafny.Array
				_ = _nw63
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw63 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Sequence = (func(_465_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_466_i4 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_465_p0, _465_p0)
						}
					})(p0)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw63 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw63).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw63).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_464_v24 = _nw63
				_464_v24 = _464_v24
				_464_v24 = _464_v24
				var _467_v25 *C0
				_ = _467_v25
				var _nw64 *C0 = New_C0_()
				_ = _nw64
				_nw64.Ctor__()
				_467_v25 = _nw64
				var _468_v26 _dafny.Map
				_ = _468_v26
				_468_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_458_v18, _458_v18), _458_v18)
				_468_v26 = (_468_v26).Update((_this).Fm2(_458_v18, globalState), _458_v18)
				_459_v19 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_461_v21, _461_v21), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm27(globalState), _461_v21))
			} else {
				(globalState).F13 = (_dafny.IntOfInt64(972)).Times(_458_v18)
				var _469_v27 _dafny.Sequence
				_ = _469_v27
				_469_v27 = _dafny.SeqOf(_458_v18, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_459_v19, _dafny.IntOfInt64(377))).Cardinality(), _dafny.IntOfInt64(735))
				var _470_v28 _dafny.Map
				_ = _470_v28
				_470_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_469_v27, _463_v23)
				var _rhs93 _dafny.Int = (_458_v18).Plus(_458_v18)
				_ = _rhs93
				var _rhs94 _dafny.Array = _463_v23
				_ = _rhs94
				var _rhs95 bool = (_470_v28).Contains(_469_v27)
				_ = _rhs95
				var _rhs96 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tmislcqxm"), _dafny.UnicodeSeqOfUtf8Bytes("jtsjc")))).Cardinality())
				_ = _rhs96
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				_lhs68.F15 = _rhs93
				_463_v23 = _rhs94
				_459_v19 = _rhs95
				_lhs69.F4 = _rhs96
				(globalState).F13 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(815))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}(func(_471_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}))).Cardinality())
				var _472_v29 *C0
				_ = _472_v29
				var _nw65 *C0 = New_C0_()
				_ = _nw65
				_nw65.Ctor__()
				_472_v29 = _nw65
				var _473_v30 *C1
				_ = _473_v30
				var _nw66 *C1 = New_C1_()
				_ = _nw66
				_nw66.Ctor__()
				_473_v30 = _nw66
			}
		}
		var _474_v31 _dafny.Array
		_ = _474_v31
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_14
		var _nw67 _dafny.Array
		_ = _nw67
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw67 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) bool = func(_475_i6 _dafny.Int) bool {
				return (Companion_D8_.Create_DC24_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(959), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('d')), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('d'))))).Cardinality())).Cardinality(), (_dafny.MultiSetOf(true, true)).Cardinality())).Cardinality(), true)).Dtor_cf41()
			}
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw67 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw67).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw67).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_474_v31 = _nw67
		var _476_v32 bool
		_ = _476_v32
		_476_v32 = false
		var _477_v33 D6
		_ = _477_v33
		_477_v33 = Companion_D6_.Create_DC17_(_476_v32, false, true)
		var _478_v34 _dafny.Sequence
		_ = _478_v34
		_478_v34 = _dafny.SeqOf(_476_v32, _476_v32, (_477_v33).Dtor_cf26())
		var _479_v35 _dafny.Int
		_ = _479_v35
		_479_v35 = _dafny.IntOfInt64(348)
		var _480_v36 _dafny.Map
		_ = _480_v36
		_480_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _479_v35)
		var _481_v37 _dafny.Map
		_ = _481_v37
		_481_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_478_v34).Select((Companion_Default___.SafeIndex((_480_v36).Cardinality(), _dafny.IntOfUint32((_478_v34).Cardinality()))).Uint32()).(bool), _474_v31)
		var _482_v38 _dafny.Map
		_ = _482_v38
		_482_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_474_v31, (func() _dafny.Array {
			if (_481_v37).Contains(_476_v32) {
				return (_481_v37).Get(_476_v32).(_dafny.Array)
			}
			return _474_v31
		})()), _479_v35)
		var _483_v39 _dafny.MultiSet
		_ = _483_v39
		_483_v39 = _dafny.MultiSetOf((_this).Fm8((_dafny.MultiSetFromSeq(Companion_Default___.Fm27(globalState))).Cardinality(), _476_v32, _479_v35, globalState), !(false))
		(globalState).F15 = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if (_482_v38).Contains(_dafny.SeqOf(_474_v31)) {
				return (_482_v38).Get(_dafny.SeqOf(_474_v31)).(_dafny.Int)
			}
			return Companion_Default___.SafeDivisionInt((_483_v39).Cardinality(), _479_v35)
		})()))
		(globalState).F7 = ((_dafny.MultiSetOf(!(_476_v32))).Cardinality()).Plus((_dafny.MultiSetFromSeq(p0)).Cardinality())
		var _484_v40 _dafny.Array
		_ = _484_v40
		var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
		_ = _nw68
		_484_v40 = _nw68
		var _485_v41 _dafny.Map
		_ = _485_v41
		_485_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_479_v35, _474_v31)
		var _486_v42 _dafny.Array
		_ = _486_v42
		var _nwElement0_14 _dafny.Array = _474_v31
		_ = _nwElement0_14
		var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(29))
		_ = _nw69
		(_nw69).ArraySet1(_nwElement0_14, 0)
		(_nw69).ArraySet1(_474_v31, 1)
		(_nw69).ArraySet1(_474_v31, 2)
		(_nw69).ArraySet1(_474_v31, 3)
		(_nw69).ArraySet1(_474_v31, 4)
		(_nw69).ArraySet1(_474_v31, 5)
		(_nw69).ArraySet1(_474_v31, 6)
		(_nw69).ArraySet1(_474_v31, 7)
		(_nw69).ArraySet1(_474_v31, 8)
		(_nw69).ArraySet1(_474_v31, 9)
		(_nw69).ArraySet1(_474_v31, 10)
		(_nw69).ArraySet1(_474_v31, 11)
		(_nw69).ArraySet1(_474_v31, 12)
		(_nw69).ArraySet1(_474_v31, 13)
		(_nw69).ArraySet1(_474_v31, 14)
		(_nw69).ArraySet1(_474_v31, 15)
		(_nw69).ArraySet1(_474_v31, 16)
		(_nw69).ArraySet1(_474_v31, 17)
		(_nw69).ArraySet1(_474_v31, 18)
		(_nw69).ArraySet1(_474_v31, 19)
		(_nw69).ArraySet1(_474_v31, 20)
		(_nw69).ArraySet1(_474_v31, 21)
		(_nw69).ArraySet1(_474_v31, 22)
		(_nw69).ArraySet1(_474_v31, 23)
		(_nw69).ArraySet1(_474_v31, 24)
		(_nw69).ArraySet1(_474_v31, 25)
		(_nw69).ArraySet1(_474_v31, 26)
		(_nw69).ArraySet1(_474_v31, 27)
		(_nw69).ArraySet1((func() _dafny.Array {
			if (_485_v41).Contains(_dafny.IntOfInt64(-38)) {
				return (_485_v41).Get(_dafny.IntOfInt64(-38)).(_dafny.Array)
			}
			return _474_v31
		})(), 28)
		_486_v42 = _nw69
		var _487_v43 D0
		_ = _487_v43
		_487_v43 = Companion_D0_.Create_DC1_(_474_v31)
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_486_v42), 0))
		_ = _index65
		(_486_v42).ArraySet1((_487_v43).Dtor_cf1(), (_index65).Int())
		var _488_v44 _dafny.Array
		_ = _488_v44
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_15
		var _nw70 _dafny.Array
		_ = _nw70
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw70 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) _dafny.MultiSet = (func(_489_v39 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_490_i7 _dafny.Int) _dafny.MultiSet {
					return _489_v39
				}
			})(_483_v39)
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw70 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw70).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw70).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_488_v44 = _nw70
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_488_v44), 0))
		_ = _index66
		(_488_v44).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_476_v32), _478_v34)), (_index66).Int())
		var _491_v45 _dafny.Set
		_ = _491_v45
		_491_v45 = _dafny.SetOf(_476_v32, _476_v32, false, _476_v32, _476_v32)
		var _492_v46 _dafny.Sequence
		_ = _492_v46
		_492_v46 = _dafny.SeqOf(_484_v40, _484_v40, _484_v40, _484_v40, _484_v40)
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_486_v42), 0))
		_ = _index67
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_488_v44), 0))
		_ = _index68
		var _rhs97 _dafny.Array = (_492_v46).Select((Companion_Default___.SafeIndex(_479_v35, _dafny.IntOfUint32((_492_v46).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs97
		var _rhs98 _dafny.Array = _474_v31
		_ = _rhs98
		var _rhs99 _dafny.MultiSet = _dafny.MultiSetOf(_476_v32, _476_v32)
		_ = _rhs99
		var _rhs100 _dafny.Set = _491_v45
		_ = _rhs100
		var _lhs70 _dafny.Array = _486_v42
		_ = _lhs70
		var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_486_v42), 0))
		_ = _lhs71
		var _lhs72 _dafny.Array = _488_v44
		_ = _lhs72
		var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_488_v44), 0))
		_ = _lhs73
		_484_v40 = _rhs97
		(_lhs70).ArraySet1(_rhs98, (_lhs71).Int())
		(_lhs72).ArraySet1(_rhs99, (_lhs73).Int())
		_491_v45 = _rhs100
		var _pat_let_tv3 = p0
		_ = _pat_let_tv3
		var _pat_let_tv4 = p0
		_ = _pat_let_tv4
		var _pat_let_tv5 = _478_v34
		_ = _pat_let_tv5
		var _pat_let_tv6 = _479_v35
		_ = _pat_let_tv6
		(globalState).F15 = func(_source10 D4) _dafny.Int {
			if _source10.Is_DC12() {
				var _493___mcc_h2 _dafny.Int = _source10.Get_().(D4_DC12).Cf16
				_ = _493___mcc_h2
				var _494_cf16 _dafny.Int = _493___mcc_h2
				_ = _494_cf16
				return (_dafny.MultiSetOf((_dafny.SetOf(_dafny.MultiSetOf(_dafny.CodePoint('y'), _dafny.CodePoint('f')), _dafny.MultiSetFromSeq(_pat_let_tv3), _dafny.MultiSetFromSeq(_pat_let_tv4))).Cardinality(), _dafny.IntOfUint32((_pat_let_tv5).Cardinality()), _494_cf16, _dafny.IntOfInt64(561), (_dafny.Zero).Minus(_dafny.IntOfInt64(-846)))).Cardinality()
			} else {
				var _495___mcc_h3 _dafny.Sequence = _source10.Get_().(D4_DC11).Cf15
				_ = _495___mcc_h3
				var _496_cf15 _dafny.Sequence = _495___mcc_h3
				_ = _496_cf15
				return _pat_let_tv6
			}
		}(_433_v1)
		var _497_v47 _dafny.Array
		_ = _497_v47
		var _nwElement0_15 _dafny.Int = _dafny.IntOfUint32((_478_v34).Cardinality())
		_ = _nwElement0_15
		var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(20))
		_ = _nw71
		(_nw71).ArraySet1(_nwElement0_15, 0)
		(_nw71).ArraySet1(_479_v35, 1)
		(_nw71).ArraySet1(_479_v35, 2)
		(_nw71).ArraySet1((_this).Fm2(_479_v35, globalState), 3)
		(_nw71).ArraySet1(_479_v35, 4)
		(_nw71).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 5)
		(_nw71).ArraySet1(_479_v35, 6)
		(_nw71).ArraySet1(_479_v35, 7)
		(_nw71).ArraySet1(_479_v35, 8)
		(_nw71).ArraySet1(_479_v35, 9)
		(_nw71).ArraySet1(_479_v35, 10)
		(_nw71).ArraySet1(_dafny.IntOfUint32((_478_v34).Cardinality()), 11)
		(_nw71).ArraySet1(_479_v35, 12)
		(_nw71).ArraySet1(_479_v35, 13)
		(_nw71).ArraySet1(_479_v35, 14)
		(_nw71).ArraySet1(_479_v35, 15)
		(_nw71).ArraySet1(_dafny.Zero, 16)
		(_nw71).ArraySet1(_479_v35, 17)
		(_nw71).ArraySet1(_dafny.IntOfInt64(-971), 18)
		(_nw71).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
			if (_483_v39).Contains(_476_v32) {
				return (_483_v39).Multiplicity(_476_v32)
			}
			return _dafny.IntOfInt64(-672)
		})()), 19)
		_497_v47 = _nw71
		var _498_v48 _dafny.Map
		_ = _498_v48
		_498_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_497_v47, true)
		_498_v48 = (_498_v48).Update(_497_v47, _476_v32)
	}
}
func (_this *C2) M11(p0 _dafny.Array, p1 _dafny.Sequence, p2 _dafny.Set, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _499_v0 _dafny.Int
		_ = _499_v0
		_499_v0 = _dafny.IntOfInt64(978)
		var _500_v1 _dafny.Set
		_ = _500_v1
		_500_v1 = _dafny.SetOf(_499_v0, _499_v0)
		if (_500_v1).Equals(_dafny.SetOf(_499_v0)) {
			var _501_v2 _dafny.Array
			_ = _501_v2
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_16
			var _nw72 _dafny.Array
			_ = _nw72
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw72 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Int = (func(_502_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_503_i0 _dafny.Int) _dafny.Int {
						return (_503_i0).Times(_502_v0)
					}
				})(_499_v0)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw72 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw72).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw72).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_501_v2 = _nw72
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))
			_ = _index69
			(_501_v2).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_501_v2, _501_v2, _501_v2)).Cardinality()), (_index69).Int())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))
			_ = _index70
			(_501_v2).ArraySet1((_499_v0).Times(_499_v0), (_index70).Int())
			(globalState).F7 = (_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int)
			var _504_v3 D2
			_ = _504_v3
			_504_v3 = Companion_D2_.Create_DC7_(p3, Companion_Default___.Fm34(_499_v0, _499_v0, p3, globalState), _501_v2)
			var _source11 D2 = _504_v3
			_ = _source11
			if _source11.Is_DC6() {
				var _505___mcc_h0 _dafny.Int = _source11.Get_().(D2_DC6).Cf7
				_ = _505___mcc_h0
				var _506_cf7 _dafny.Int = _505___mcc_h0
				_ = _506_cf7
				var _507_v4 bool
				_ = _507_v4
				_507_v4 = false
				_507_v4 = !((_this).Fm8(_506_cf7, p3, (_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int), globalState))
				var _508_v5 *C1
				_ = _508_v5
				var _nw73 *C1 = New_C1_()
				_ = _nw73
				_nw73.Ctor__()
				_508_v5 = _nw73
				var _509_v6 _dafny.Map
				_ = _509_v6
				_509_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_500_v1, _499_v0)
				_509_v6 = _509_v6
				_507_v4 = (func() bool {
					if _507_v4 {
						return _507_v4
					}
					return _507_v4
				})()
			} else if _source11.Is_DC7() {
				var _510___mcc_h1 bool = _source11.Get_().(D2_DC7).Cf8
				_ = _510___mcc_h1
				var _511___mcc_h2 _dafny.MultiSet = _source11.Get_().(D2_DC7).Cf9
				_ = _511___mcc_h2
				var _512___mcc_h3 _dafny.Array = _source11.Get_().(D2_DC7).Cf10
				_ = _512___mcc_h3
				var _513_cf10 _dafny.Array = _512___mcc_h3
				_ = _513_cf10
				var _514_cf9 _dafny.MultiSet = _511___mcc_h2
				_ = _514_cf9
				var _515_cf8 bool = _510___mcc_h1
				_ = _515_cf8
				var _516_v7 _dafny.CodePoint
				_ = _516_v7
				_516_v7 = _dafny.CodePoint('q')
				var _517_v8 _dafny.MultiSet
				_ = _517_v8
				_517_v8 = _dafny.MultiSetOf(_499_v0)
				var _518_v9 _dafny.Map
				_ = _518_v9
				_518_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_499_v0, (_dafny.SetOf((func() _dafny.Int {
					if (_517_v8).Contains(_499_v0) {
						return (_517_v8).Multiplicity(_499_v0)
					}
					return _dafny.IntOfInt64(286)
				})(), ((Companion_Default___.Fm35(globalState)).Update((_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int), _515_cf8)).Cardinality())).Union(_500_v1))
				var _519_v10 _dafny.Array
				_ = _519_v10
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw74
				_519_v10 = _nw74
				var _520_v11 _dafny.Map
				_ = _520_v11
				_520_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_499_v0, p0)
				var _521_v12 _dafny.Sequence
				_ = _521_v12
				_521_v12 = _dafny.UnicodeSeqOfUtf8Bytes("jf")
				var _522_v13 _dafny.Map
				_ = _522_v13
				_522_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_517_v8).Update(_dafny.IntOfUint32((_521_v12).Cardinality()), Companion_Default___.Abs((_dafny.Zero).Minus((_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int)))), _516_v7)
				var _523_v14 _dafny.Map
				_ = _523_v14
				_523_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_521_v12).Cardinality()), p3)
				var _rhs101 _dafny.CodePoint = (func() _dafny.CodePoint {
					if p3 {
						return (func() _dafny.CodePoint {
							if (_522_v13).Contains(_517_v8) {
								return (_522_v13).Get(_517_v8).(_dafny.CodePoint)
							}
							return _dafny.CodePoint('j')
						})()
					}
					return _516_v7
				})()
				_ = _rhs101
				var _rhs102 _dafny.Map = _518_v9
				_ = _rhs102
				var _rhs103 _dafny.Array = p0
				_ = _rhs103
				var _rhs104 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_523_v14).Cardinality(), p0)).Merge(_520_v11)).Merge(_520_v11)
				_ = _rhs104
				var _rhs105 bool = (func() bool {
					if _515_cf8 {
						return p3
					}
					return !(p3)
				})()
				_ = _rhs105
				_516_v7 = _rhs101
				_518_v9 = _rhs102
				_519_v10 = _rhs103
				_520_v11 = _rhs104
				_515_cf8 = _rhs105
				_523_v14 = (_523_v14).Update(_499_v0, (p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool))
				_515_cf8 = true
				var _524_v15 _dafny.Set
				_ = _524_v15
				_524_v15 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("jfqn"))
				var _rhs106 _dafny.Int = (_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int)
				_ = _rhs106
				var _rhs107 _dafny.Set = _524_v15
				_ = _rhs107
				var _lhs74 *GlobalState = globalState
				_ = _lhs74
				_lhs74.F4 = _rhs106
				_524_v15 = _rhs107
			} else if _source11.Is_DC5() {
				var _525___mcc_h4 _dafny.Array = _source11.Get_().(D2_DC5).Cf6
				_ = _525___mcc_h4
				var _526_cf6 _dafny.Array = _525___mcc_h4
				_ = _526_cf6
				var _527_v16 _dafny.Array
				_ = _527_v16
				var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
				_ = _nw75
				_527_v16 = _nw75
				var _528_v17 _dafny.Map
				_ = _528_v17
				_528_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_527_v16), 0))
				_ = _index71
				(_527_v16).ArraySet1((_528_v17).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)), (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_527_v16), 0))
				_ = _index72
				(_527_v16).ArraySet1((_528_v17).Merge((_528_v17).Merge(Companion_Default___.Fm36((_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int), globalState))), (_index72).Int())
				_526_cf6 = (func() _dafny.Array {
					if p3 {
						return (_504_v3).Dtor_cf10()
					}
					return (func() _dafny.Array {
						if p3 {
							return _526_cf6
						}
						return _526_cf6
					})()
				})()
				var _529_v18 _dafny.Set
				_ = _529_v18
				_529_v18 = _dafny.SetOf(_501_v2, _526_cf6, _526_cf6, _526_cf6)
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((p0), 0))
				_ = _index73
				(p0).ArraySet1((_529_v18).IsProperSubsetOf(_529_v18), (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((p0), 0))
				_ = _index74
				var _rhs108 _dafny.Int = ((_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int)).Minus(Companion_Default___.SafeModuloInt((_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-246)))
				_ = _rhs108
				var _rhs109 bool = p3
				_ = _rhs109
				var _rhs110 _dafny.Int = (_dafny.Zero).Minus(_499_v0)
				_ = _rhs110
				var _lhs75 _dafny.Array = p0
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((p0), 0))
				_ = _lhs76
				_499_v0 = _rhs108
				(_lhs75).ArraySet1(_rhs109, (_lhs76).Int())
				_499_v0 = _rhs110
				var _530_v19 D2
				_ = _530_v19
				_530_v19 = Companion_D2_.Create_DC5_(_501_v2)
				var _531_v20 _dafny.Map
				_ = _531_v20
				_531_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_530_v19, p3)
				var _532_v21 _dafny.Array
				_ = _532_v21
				var _nwElement0_16 _dafny.Map = _531_v20
				_ = _nwElement0_16
				var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(2))
				_ = _nw76
				(_nw76).ArraySet1(_nwElement0_16, 0)
				(_nw76).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_530_v19, p3), 1)
				_532_v21 = _nw76
				var _533_v22 _dafny.Map
				_ = _533_v22
				_533_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_530_v19, true)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_532_v21), 0))
				_ = _index75
				(_532_v21).ArraySet1(_533_v22, (_index75).Int())
				var _534_v23 _dafny.Sequence
				_ = _534_v23
				_534_v23 = _dafny.UnicodeSeqOfUtf8Bytes("racjbb")
				var _535_v24 D1
				_ = _535_v24
				_535_v24 = Companion_D1_.Create_DC3_(_534_v23, p3)
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_532_v21), 0))
				_ = _index76
				var _rhs111 _dafny.Map = _533_v22
				_ = _rhs111
				var _rhs112 D1 = (func() D1 {
					if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
						return _535_v24
					}
					return _535_v24
				})()
				_ = _rhs112
				var _rhs113 _dafny.Int = (_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int)
				_ = _rhs113
				var _lhs77 _dafny.Array = _532_v21
				_ = _lhs77
				var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_532_v21), 0))
				_ = _lhs78
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				(_lhs77).ArraySet1(_rhs111, (_lhs78).Int())
				_535_v24 = _rhs112
				_lhs79.F13 = _rhs113
			} else {
				var _536___mcc_h5 D2 = _source11.Get_().(D2_DC8).Cf11
				_ = _536___mcc_h5
				var _537_cf11 D2 = _536___mcc_h5
				_ = _537_cf11
				var _538_v25 _dafny.Map
				_ = _538_v25
				_538_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_dafny.Zero).Minus(_499_v0))
				var _539_v26 _dafny.Sequence
				_ = _539_v26
				_539_v26 = _dafny.SeqOf(_dafny.IntOfInt64(331), (_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int))
				var _540_v27 _dafny.MultiSet
				_ = _540_v27
				_540_v27 = _dafny.MultiSetOf((_539_v26).Select((Companion_Default___.SafeIndex((_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_539_v26).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(86))
				var _541_v28 _dafny.Sequence
				_ = _541_v28
				_541_v28 = _dafny.SeqOf(_dafny.MultiSetOf((func() _dafny.Int {
					if (_538_v25).Contains(p3) {
						return (_538_v25).Get(p3).(_dafny.Int)
					}
					return (_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int)
				})()), _540_v27, _540_v27)
				_541_v28 = _541_v28
				var _542_v29 _dafny.Map
				_ = _542_v29
				_542_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_499_v0, p3)
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))
				_ = _index77
				(_501_v2).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf((_542_v29).Cardinality())).Cardinality()), _499_v0)).Times(((_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int)).Times(_499_v0)))), (_index77).Int())
				(globalState).F13 = Companion_Default___.SafeDivisionInt(_499_v0, (_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int))
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((p0), 0))
				_ = _index78
				(p0).ArraySet1(Companion_Default___.Fm26(globalState), (_index78).Int())
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((p0), 0))
				_ = _index79
				(p0).ArraySet1(false, (_index79).Int())
			}
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))
			_ = _index80
			(_501_v2).ArraySet1(Companion_Default___.SafeDivisionInt(_499_v0, _499_v0), (_index80).Int())
			var _543_v30 _dafny.Sequence
			_ = _543_v30
			_543_v30 = _dafny.SeqOf(_501_v2)
			var _544_v31 _dafny.Sequence
			_ = _544_v31
			_544_v31 = _dafny.SeqOf((_543_v30).Select((Companion_Default___.SafeIndex((_501_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_501_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_543_v30).Cardinality()))).Uint32()).(_dafny.Array), _501_v2, _501_v2)
			_544_v31 = _544_v31
		} else {
			var _545_v32 _dafny.Array
			_ = _545_v32
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_17
			var _nw77 _dafny.Array
			_ = _nw77
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw77 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) bool = (func(_546_p1 _dafny.Sequence, _547_v0 _dafny.Int) func(_dafny.Int) bool {
					return func(_548_i1 _dafny.Int) bool {
						return (_546_p1).Select((Companion_Default___.SafeIndex(_547_v0, _dafny.IntOfUint32((_546_p1).Cardinality()))).Uint32()).(bool)
					}
				})(p1, _499_v0)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw77 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw77).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw77).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_545_v32 = _nw77
			var _549_v33 _dafny.Array
			_ = _549_v33
			var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw78
			_549_v33 = _nw78
			var _rhs114 _dafny.Int = _dafny.IntOfInt64(-932)
			_ = _rhs114
			var _rhs115 _dafny.Int = _499_v0
			_ = _rhs115
			var _rhs116 _dafny.Array = _545_v32
			_ = _rhs116
			var _rhs117 _dafny.Array = _549_v33
			_ = _rhs117
			var _lhs80 *GlobalState = globalState
			_ = _lhs80
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			_lhs80.F15 = _rhs114
			_lhs81.F7 = _rhs115
			_545_v32 = _rhs116
			_549_v33 = _rhs117
			var _550_v34 bool
			_ = _550_v34
			_550_v34 = false
			_550_v34 = p3
			var _551_v35 *C1
			_ = _551_v35
			var _nw79 *C1 = New_C1_()
			_ = _nw79
			_nw79.Ctor__()
			_551_v35 = _nw79
			var _552_v36 _dafny.Sequence
			_ = _552_v36
			_552_v36 = _dafny.SeqOf(_499_v0, _499_v0)
			_550_v34 = (_499_v0).Cmp((func() _dafny.Int {
				if p3 {
					return (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_552_v36, (Companion_Default___.SafeIndex(_499_v0, _dafny.IntOfUint32((_552_v36).Cardinality()))).Uint32(), _499_v0))).Cardinality()
				}
				return _499_v0
			})()) >= 0
			_550_v34 = p3
		}
		var _553_v37 _dafny.Array
		_ = _553_v37
		var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw80
		_553_v37 = _nw80
		var _554_v38 bool
		_ = _554_v38
		_554_v38 = false
		var _rhs118 _dafny.Int = (_dafny.IntOfInt64(784)).Plus(_dafny.IntOfInt64(46))
		_ = _rhs118
		var _rhs119 _dafny.Array = _553_v37
		_ = _rhs119
		var _rhs120 bool = !(p3)
		_ = _rhs120
		var _rhs121 bool = _554_v38
		_ = _rhs121
		var _rhs122 bool = _554_v38
		_ = _rhs122
		var _lhs82 *GlobalState = globalState
		_ = _lhs82
		_lhs82.F13 = _rhs118
		_553_v37 = _rhs119
		_554_v38 = _rhs120
		_554_v38 = _rhs121
		_554_v38 = _rhs122
		var _hi3 _dafny.Int = _dafny.IntOfInt64(-680)
		_ = _hi3
		for _555_i2 := _499_v0; _555_i2.Cmp(_hi3) < 0; _555_i2 = _555_i2.Plus(_dafny.One) {
			var _556_v39 D0
			_ = _556_v39
			_556_v39 = Companion_D0_.Create_DC1_(p0)
			var _557_v40 _dafny.MultiSet
			_ = _557_v40
			_557_v40 = _dafny.MultiSetOf(_556_v39, _556_v39, _556_v39, _556_v39)
			var _558_v41 _dafny.Map
			_ = _558_v41
			_558_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_555_i2, _557_v40)
			var _559_v42 _dafny.Sequence
			_ = _559_v42
			_559_v42 = _dafny.SeqOf(_556_v39)
			_558_v41 = (_558_v41).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ayhpbe"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(613))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}(func(_560_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})))).Cardinality()), _dafny.MultiSetFromSeq(_559_v42))
			var _561_v43 _dafny.Array
			_ = _561_v43
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw81
			_561_v43 = _nw81
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_561_v43), 0))
			_ = _index81
			(_561_v43).ArraySet1(_555_i2, (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_561_v43), 0))
			_ = _index82
			(_561_v43).ArraySet1(_555_i2, (_index82).Int())
			_554_v38 = (_555_i2).Cmp(_555_i2) == 0
			if p3 {
				var _562_v44 _dafny.Sequence
				_ = _562_v44
				_562_v44 = _dafny.UnicodeSeqOfUtf8Bytes("bypilo")
				var _rhs123 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_562_v44, _562_v44)
				_ = _rhs123
				var _rhs124 _dafny.Int = (_dafny.Zero).Minus((_561_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_561_v43), 0))).Int()).(_dafny.Int))
				_ = _rhs124
				var _lhs83 *GlobalState = globalState
				_ = _lhs83
				_562_v44 = _rhs123
				_lhs83.F15 = _rhs124
				var _563_v45 *C0
				_ = _563_v45
				var _nw82 *C0 = New_C0_()
				_ = _nw82
				_nw82.Ctor__()
				_563_v45 = _nw82
				_554_v38 = ((_561_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_561_v43), 0))).Int()).(_dafny.Int)).Cmp(_499_v0) != 0
				var _564_v46 _dafny.Map
				_ = _564_v46
				_564_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), p3)
				var _565_v47 _dafny.MultiSet
				_ = _565_v47
				_565_v47 = _dafny.MultiSetOf(p3, !(_554_v38))
				var _566_v48 _dafny.MultiSet
				_ = _566_v48
				_566_v48 = _dafny.MultiSetOf((_565_v47).Cardinality())
				_564_v46 = (_564_v46).Update(_554_v38, (_566_v48).IsDisjointFrom(_566_v48))
				var _567_v49 _dafny.Array
				_ = _567_v49
				var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
				_ = _nw83
				_567_v49 = _nw83
				_567_v49 = _567_v49
			} else {
				var _568_v50 _dafny.CodePoint
				_ = _568_v50
				_568_v50 = _dafny.CodePoint('s')
				var _569_v51 _dafny.Map
				_ = _569_v51
				_569_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_568_v50, p0)
				(globalState).F4 = (_569_v51).Cardinality()
				var _570_v52 _dafny.Sequence
				_ = _570_v52
				_570_v52 = _dafny.SeqOf(_554_v38, _554_v38, !(_554_v38))
				_570_v52 = Companion_Default___.Fm27(globalState)
				var _571_v53 _dafny.Array
				_ = _571_v53
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_18
				var _nw84 _dafny.Array
				_ = _nw84
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw84 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Sequence = (func(_572_p3 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_573_i4 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_572_p3, _572_p3, !(true))
						}
					})(p3)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw84 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw84).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw84).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_571_v53 = _nw84
				var _574_v54 _dafny.Map
				_ = _574_v54
				_574_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_571_v53, _554_v38)
				var _575_v55 _dafny.MultiSet
				_ = _575_v55
				_575_v55 = _dafny.MultiSetOf(_574_v54, _574_v54, _574_v54)
				_554_v38 = !((_575_v55).IsProperSubsetOf(_575_v55))
				var _576_v56 _dafny.Sequence
				_ = _576_v56
				_576_v56 = _dafny.UnicodeSeqOfUtf8Bytes("xcbscyt")
				var _577_v57 _dafny.Array
				_ = _577_v57
				var _nwElement0_17 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("xfaydljs")
				_ = _nwElement0_17
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(3))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_17, 0)
				(_nw85).ArraySet1(_576_v56, 1)
				(_nw85).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_576_v56, _576_v56), 2)
				_577_v57 = _nw85
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_577_v57), 0))
				_ = _index83
				(_577_v57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_576_v56, _576_v56), (_index83).Int())
				var _578_v58 _dafny.Sequence
				_ = _578_v58
				_578_v58 = _dafny.SeqOf(_576_v56)
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_577_v57), 0))
				_ = _index84
				var _rhs125 bool = _554_v38
				_ = _rhs125
				var _rhs126 _dafny.Sequence = (func() _dafny.Sequence {
					if ((_dafny.Zero).Minus((_561_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_561_v43), 0))).Int()).(_dafny.Int))).Cmp((_561_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_561_v43), 0))).Int()).(_dafny.Int)) >= 0 {
						return _576_v56
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(352))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}((func(_579_v50 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_580_i5 _dafny.Int) _dafny.CodePoint {
							return _579_v50
						}
					})(_568_v50)))
				})()
				_ = _rhs126
				var _rhs127 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_578_v58, (Companion_Default___.SafeIndex((_561_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen((_561_v43), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_578_v58).Cardinality()))).Uint32(), _576_v56), _578_v58), _578_v58)
				_ = _rhs127
				var _rhs128 bool = !(p3) || (!(_554_v38))
				_ = _rhs128
				var _lhs84 _dafny.Array = _577_v57
				_ = _lhs84
				var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_577_v57), 0))
				_ = _lhs85
				_554_v38 = _rhs125
				(_lhs84).ArraySet1(_rhs126, (_lhs85).Int())
				_578_v58 = _rhs127
				_554_v38 = _rhs128
				var _581_v59 *C1
				_ = _581_v59
				var _nw86 *C1 = New_C1_()
				_ = _nw86
				_nw86.Ctor__()
				_581_v59 = _nw86
			}
		}
		_554_v38 = !(!(_554_v38)) || (p3)
		var _582_v60 D6
		_ = _582_v60
		_582_v60 = Companion_D6_.Create_DC17_(_554_v38, !(p3), _554_v38)
		var _583_i6 _dafny.Int
		_ = _583_i6
		_583_i6 = _dafny.Zero
		{
			for !(_554_v38) || ((_582_v60).Dtor_cf25()) {
				{
					if (_583_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_583_i6 = (_583_i6).Plus(_dafny.One)
					var _584_v61 _dafny.Sequence
					_ = _584_v61
					_584_v61 = _dafny.UnicodeSeqOfUtf8Bytes("wfjnh")
					_584_v61 = _584_v61
					_554_v38 = (Companion_Default___.SafeModuloInt(_499_v0, _499_v0)).Cmp(_499_v0) >= 0
					var _585_v62 _dafny.Array
					_ = _585_v62
					var _len0_19 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_19
					var _nw87 _dafny.Array
					_ = _nw87
					if _len0_19.Cmp(_dafny.Zero) == 0 {
						_nw87 = _dafny.NewArray(_len0_19)
					} else {
						var _init19 func(_dafny.Int) _dafny.Int = (func(_586_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_587_i7 _dafny.Int) _dafny.Int {
								return (_587_i7).Times(_586_v0)
							}
						})(_499_v0)
						_ = _init19
						var _element0_19 = _init19(_dafny.Zero)
						_ = _element0_19
						_nw87 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
						(_nw87).ArraySet1(_element0_19, 0)
						var _nativeLen0_19 = (_len0_19).Int()
						_ = _nativeLen0_19
						for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
							(_nw87).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
						}
					}
					_585_v62 = _nw87
					var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_585_v62), 0))
					_ = _index85
					(_585_v62).ArraySet1((_dafny.Zero).Minus(((_dafny.SetOf(_499_v0)).Cardinality()).Plus(_499_v0)), (_index85).Int())
					var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_585_v62), 0))
					_ = _index86
					(_585_v62).ArraySet1((_dafny.Zero).Minus(_499_v0), (_index86).Int())
					var _588_v63 _dafny.Map
					_ = _588_v63
					_588_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf((p1).Select((Companion_Default___.SafeIndex((_585_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_585_v62), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool)))
					var _589_v64 _dafny.Map
					_ = _589_v64
					_589_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_585_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_585_v62), 0))).Int()).(_dafny.Int), _dafny.SetOf(!(p3)))
					var _590_v65 _dafny.Array
					_ = _590_v65
					var _nwElement0_18 _dafny.Set = (p2).Union(p2)
					_ = _nwElement0_18
					var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(21))
					_ = _nw88
					(_nw88).ArraySet1(_nwElement0_18, 0)
					(_nw88).ArraySet1((p2).Difference(p2), 1)
					(_nw88).ArraySet1(p2, 2)
					(_nw88).ArraySet1(p2, 3)
					(_nw88).ArraySet1(p2, 4)
					(_nw88).ArraySet1((p2).Difference(Companion_Default___.Fm37(globalState)), 5)
					(_nw88).ArraySet1(p2, 6)
					(_nw88).ArraySet1((_dafny.SetOf(!(p3))).Difference(p2), 7)
					(_nw88).ArraySet1((p2).Difference(Companion_Default___.Fm37(globalState)), 8)
					(_nw88).ArraySet1(p2, 9)
					(_nw88).ArraySet1(p2, 10)
					(_nw88).ArraySet1((p2).Difference(p2), 11)
					(_nw88).ArraySet1(p2, 12)
					(_nw88).ArraySet1((p2).Difference((func() _dafny.Set {
						if (_588_v63).Contains(_554_v38) {
							return (_588_v63).Get(_554_v38).(_dafny.Set)
						}
						return _dafny.SetOf(true)
					})()), 13)
					(_nw88).ArraySet1(p2, 14)
					(_nw88).ArraySet1(p2, 15)
					(_nw88).ArraySet1(p2, 16)
					(_nw88).ArraySet1((func() _dafny.Set {
						if _554_v38 {
							return p2
						}
						return p2
					})(), 17)
					(_nw88).ArraySet1((p2).Intersection(p2), 18)
					(_nw88).ArraySet1((p2).Union(p2), 19)
					(_nw88).ArraySet1((func() _dafny.Set {
						if (_589_v64).Contains(_499_v0) {
							return (_589_v64).Get(_499_v0).(_dafny.Set)
						}
						return p2
					})(), 20)
					_590_v65 = _nw88
					var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_590_v65), 0))
					_ = _index87
					(_590_v65).ArraySet1(p2, (_index87).Int())
					var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_590_v65), 0))
					_ = _index88
					(_590_v65).ArraySet1(_dafny.SetOf(_554_v38), (_index88).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		(globalState).F15 = _499_v0
		r0 = (func() _dafny.Int {
			if _554_v38 {
				return _dafny.IntOfInt64(786)
			}
			return _dafny.IntOfInt64(111)
		})()
		var _591_v66 D4
		_ = _591_v66
		_591_v66 = Companion_D4_.Create_DC11_(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qpuhuc")))
		var _pat_let_tv7 = _499_v0
		_ = _pat_let_tv7
		r1 = func(_source12 D4) _dafny.Int {
			if _source12.Is_DC12() {
				var _592___mcc_h6 _dafny.Int = _source12.Get_().(D4_DC12).Cf16
				_ = _592___mcc_h6
				var _593_cf16 _dafny.Int = _592___mcc_h6
				_ = _593_cf16
				return _593_cf16
			} else {
				var _594___mcc_h7 _dafny.Sequence = _source12.Get_().(D4_DC11).Cf15
				_ = _594___mcc_h7
				var _595_cf15 _dafny.Sequence = _594___mcc_h7
				_ = _595_cf15
				return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(204), _pat_let_tv7))
			}
		}(_591_v66)
		r2 = (_dafny.Zero).Minus((_499_v0).Times(_dafny.IntOfInt64(798)))
		return r0, r1, r2
	}
}
func (_this *C2) M12(p0 bool, p1 bool, p2 _dafny.Array, p3 bool, globalState *GlobalState) {
	{
		var _596_v0 _dafny.Set
		_ = _596_v0
		_596_v0 = _dafny.SetOf(Companion_Default___.Fm38(p1, globalState))
		var _597_v1 _dafny.CodePoint
		_ = _597_v1
		_597_v1 = _dafny.CodePoint('q')
		var _598_v2 _dafny.Sequence
		_ = _598_v2
		_598_v2 = _dafny.SeqOf(_597_v1)
		var _599_v4 _dafny.Sequence
		_ = _599_v4
		_599_v4 = _dafny.SeqOf(func() _dafny.Set {
			var _coll30 = _dafny.NewBuilder()
			_ = _coll30
			for _iter30 := _dafny.Iterate((_598_v2).Elements()); ; {
				_compr_30, _ok30 := _iter30()
				if !_ok30 {
					break
				}
				var _600_v3 _dafny.CodePoint
				_600_v3 = interface{}(_compr_30).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_598_v2, _600_v3) {
					_coll30.Add(_600_v3)
				}
			}
			return _coll30.ToSet()
		}(), _596_v0)
		var _601_i0 _dafny.Int
		_ = _601_i0
		_601_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_596_v0), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(221))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_662_v0 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_663_i1 _dafny.Int) _dafny.Set {
					return _662_v0
				}
			})(_596_v0))), _599_v4)) {
				{
					if (_601_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_601_i0 = (_601_i0).Plus(_dafny.One)
					var _602_v5 _dafny.Int
					_ = _602_v5
					_602_v5 = _dafny.IntOfInt64(813)
					(globalState).F13 = (func() _dafny.Int {
						if p3 {
							return (_dafny.IntOfInt64(72)).Plus(_602_v5)
						}
						return _602_v5
					})()
					var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))
					_ = _index89
					(p2).ArraySet1(_602_v5, (_index89).Int())
					var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))
					_ = _index90
					(p2).ArraySet1((func() _dafny.Set {
						var _coll31 = _dafny.NewBuilder()
						_ = _coll31
						for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(218), _dafny.IntOfInt64(157))); ; {
							_compr_31, _ok31 := _iter31()
							if !_ok31 {
								break
							}
							var _603_v6 _dafny.Int
							_603_v6 = interface{}(_compr_31).(_dafny.Int)
							if ((_dafny.IntOfInt64(218)).Cmp(_603_v6) <= 0) && ((_603_v6).Cmp(_dafny.IntOfInt64(157)) < 0) {
								_coll31.Add((_603_v6).Plus((_dafny.MultiSetOf(p1, p0)).Cardinality()))
							}
						}
						return _coll31.ToSet()
					}()).Cardinality(), (_index90).Int())
					var _604_v7 D2
					_ = _604_v7
					_604_v7 = Companion_D2_.Create_DC6_(_602_v5)
					var _605_v8 D2
					_ = _605_v8
					_605_v8 = Companion_D2_.Create_DC8_(_604_v7)
					var _source13 D2 = _605_v8
					_ = _source13
					if _source13.Is_DC6() {
						var _606___mcc_h0 _dafny.Int = _source13.Get_().(D2_DC6).Cf7
						_ = _606___mcc_h0
						var _607_cf7 _dafny.Int = _606___mcc_h0
						_ = _607_cf7
						var _608_v9 bool
						_ = _608_v9
						_608_v9 = true
						var _609_v10 _dafny.Map
						_ = _609_v10
						_609_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_602_v5, (_dafny.Zero).Minus(_dafny.IntOfInt64(-213)))
						_608_v9 = (_this).Fm24(_602_v5, p0, _609_v10, _607_cf7, globalState)
						_608_v9 = p3
						var _610_v11 _dafny.Array
						_ = _610_v11
						var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
						_ = _nw89
						_610_v11 = _nw89
						var _611_v12 _dafny.Sequence
						_ = _611_v12
						_611_v12 = _dafny.SeqOf((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))
						var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_610_v11), 0))
						_ = _index91
						(_610_v11).ArraySet1(_611_v12, (_index91).Int())
						var _612_v13 _dafny.Map
						_ = _612_v13
						_612_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), p0)
						var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_610_v11), 0))
						_ = _index92
						(_610_v11).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(428))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg43 _dafny.Int) interface{} {
								return coer43(arg43)
							}
						}((func(_613_cf7 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_614_i2 _dafny.Int) _dafny.Int {
								return _613_cf7
							}
						})(_607_cf7))), (func() _dafny.Sequence {
							if p0 {
								return _611_v12
							}
							return _611_v12
						})()), (Companion_Default___.SafeIndex((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(428))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}((func(_615_cf7 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_616_i2 _dafny.Int) _dafny.Int {
								return _615_cf7
							}
						})(_607_cf7))), (func() _dafny.Sequence {
							if p0 {
								return _611_v12
							}
							return _611_v12
						})())).Cardinality()))).Uint32(), (_612_v13).Cardinality()), (_index92).Int())
						var _617_v14 _dafny.Array
						_ = _617_v14
						var _nw90 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
						_ = _nw90
						_617_v14 = _nw90
						var _618_v15 D0
						_ = _618_v15
						_618_v15 = Companion_D0_.Create_DC1_(_617_v14)
						var _619_v16 _dafny.Set
						_ = _619_v16
						_619_v16 = _dafny.SetOf((_618_v15).Dtor_cf1(), _617_v14)
						var _rhs129 _dafny.Set = (_619_v16).Union(_619_v16)
						_ = _rhs129
						var _rhs130 _dafny.Int = ((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(185))
						_ = _rhs130
						var _rhs131 bool = ((_dafny.Zero).Minus(_607_cf7)).Cmp(Companion_Default___.SafeModuloInt(_607_cf7, _607_cf7)) >= 0
						_ = _rhs131
						var _rhs132 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_598_v2, _dafny.Companion_Sequence_.Concatenate(_598_v2, _dafny.UnicodeSeqOfUtf8Bytes("sqk")))
						_ = _rhs132
						_619_v16 = _rhs129
						_602_v5 = _rhs130
						_608_v9 = _rhs131
						_608_v9 = _rhs132
					} else if _source13.Is_DC7() {
						var _620___mcc_h1 bool = _source13.Get_().(D2_DC7).Cf8
						_ = _620___mcc_h1
						var _621___mcc_h2 _dafny.MultiSet = _source13.Get_().(D2_DC7).Cf9
						_ = _621___mcc_h2
						var _622___mcc_h3 _dafny.Array = _source13.Get_().(D2_DC7).Cf10
						_ = _622___mcc_h3
						var _623_cf10 _dafny.Array = _622___mcc_h3
						_ = _623_cf10
						var _624_cf9 _dafny.MultiSet = _621___mcc_h2
						_ = _624_cf9
						var _625_cf8 bool = _620___mcc_h1
						_ = _625_cf8
						(globalState).F13 = (_dafny.Zero).Minus((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))
						var _626_v17 _dafny.Array
						_ = _626_v17
						var _len0_20 _dafny.Int = _dafny.IntOfInt64(13)
						_ = _len0_20
						var _nw91 _dafny.Array
						_ = _nw91
						if _len0_20.Cmp(_dafny.Zero) == 0 {
							_nw91 = _dafny.NewArray(_len0_20)
						} else {
							var _init20 func(_dafny.Int) bool = (func(_627_p1 bool) func(_dafny.Int) bool {
								return func(_628_i3 _dafny.Int) bool {
									return _627_p1
								}
							})(p1)
							_ = _init20
							var _element0_20 = _init20(_dafny.Zero)
							_ = _element0_20
							_nw91 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
							(_nw91).ArraySet1(_element0_20, 0)
							var _nativeLen0_20 = (_len0_20).Int()
							_ = _nativeLen0_20
							for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
								(_nw91).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
							}
						}
						_626_v17 = _nw91
						var _629_v18 _dafny.MultiSet
						_ = _629_v18
						_629_v18 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(114))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}((func(_630_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_631_i4 _dafny.Int) _dafny.CodePoint {
								return _630_v1
							}
						})(_597_v1))), _598_v2)
						var _632_v19 _dafny.Set
						_ = _632_v19
						_632_v19 = _dafny.SetOf(p3)
						var _633_v20 _dafny.Sequence
						_ = _633_v20
						_633_v20 = _dafny.SeqOf(p1)
						var _634_v21 D1
						_ = _634_v21
						_634_v21 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_633_v20).Cardinality()))
						var _635_v22 _dafny.Map
						_ = _635_v22
						_635_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), _634_v21)
						var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_626_v17), 0))
						_ = _index93
						(_626_v17).ArraySet1(!(!(_629_v18).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("wqrdii"), Companion_Default___.Fm30((_632_v19).Cardinality(), p1, _dafny.SetOf(_602_v5, (_635_v22).Cardinality(), _602_v5), p3, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-842))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg46 _dafny.Int) interface{} {
								return coer46(arg46)
							}
						}((func(_636_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_637_i5 _dafny.Int) _dafny.CodePoint {
								return _636_v1
							}
						})(_597_v1))), _598_v2))), (_index93).Int())
						var _638_v23 _dafny.Map
						_ = _638_v23
						_638_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
						var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(121), _dafny.ArrayLen((_626_v17), 0))
						_ = _index94
						(_626_v17).ArraySet1((func() bool {
							if (_638_v23).Contains(p0) {
								return (_638_v23).Get(p0).(bool)
							}
							return p1
						})(), (_index94).Int())
						_602_v5 = ((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)).Minus(_602_v5)
						_598_v2 = (Companion_D1_.Create_DC3_(_598_v2, true)).Dtor_cf3()
					} else if _source13.Is_DC5() {
						var _639___mcc_h4 _dafny.Array = _source13.Get_().(D2_DC5).Cf6
						_ = _639___mcc_h4
						var _640_cf6 _dafny.Array = _639___mcc_h4
						_ = _640_cf6
						var _641_v24 _dafny.Array
						_ = _641_v24
						var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
						_ = _nw92
						_641_v24 = _nw92
						var _642_v25 _dafny.Array
						_ = _642_v25
						var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
						_ = _nw93
						_642_v25 = _nw93
						var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_641_v24), 0))
						_ = _index95
						(_641_v24).ArraySet1(_642_v25, (_index95).Int())
						var _643_v26 _dafny.Map
						_ = _643_v26
						_643_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _602_v5)
						var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_641_v24), 0))
						_ = _index96
						var _rhs133 _dafny.Int = (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)
						_ = _rhs133
						var _rhs134 _dafny.Int = ((_643_v26).Cardinality()).Plus((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))
						_ = _rhs134
						var _rhs135 _dafny.Array = _642_v25
						_ = _rhs135
						var _rhs136 _dafny.CodePoint = (_598_v2).Select((Companion_Default___.SafeIndex(((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfUint32((_598_v2).Cardinality())), _dafny.IntOfUint32((_598_v2).Cardinality()))).Uint32()).(_dafny.CodePoint)
						_ = _rhs136
						var _rhs137 _dafny.Int = _602_v5
						_ = _rhs137
						var _lhs86 *GlobalState = globalState
						_ = _lhs86
						var _lhs87 *GlobalState = globalState
						_ = _lhs87
						var _lhs88 _dafny.Array = _641_v24
						_ = _lhs88
						var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_641_v24), 0))
						_ = _lhs89
						var _lhs90 *GlobalState = globalState
						_ = _lhs90
						_lhs86.F4 = _rhs133
						_lhs87.F4 = _rhs134
						(_lhs88).ArraySet1(_rhs135, (_lhs89).Int())
						_597_v1 = _rhs136
						_lhs90.F7 = _rhs137
						var _644_v27 _dafny.Sequence
						_ = _644_v27
						_644_v27 = _dafny.SeqOf((_this).Fm2(_602_v5, globalState))
						var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))
						_ = _index97
						(p2).ArraySet1((Companion_Default___.SafeModuloInt((_this).Fm2((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), globalState), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))).Plus((_644_v27).Select((Companion_Default___.SafeIndex(_602_v5, _dafny.IntOfUint32((_644_v27).Cardinality()))).Uint32()).(_dafny.Int)), (_index97).Int())
						var _645_v28 bool
						_ = _645_v28
						_645_v28 = false
						var _646_v29 _dafny.MultiSet
						_ = _646_v29
						_646_v29 = _dafny.MultiSetOf(_dafny.IntOfUint32((_598_v2).Cardinality()), (_dafny.Zero).Minus(_602_v5), _602_v5)
						var _rhs138 _dafny.Int = _602_v5
						_ = _rhs138
						var _rhs139 _dafny.Int = (_dafny.Zero).Minus(_602_v5)
						_ = _rhs139
						var _rhs140 bool = true
						_ = _rhs140
						var _rhs141 _dafny.Int = ((_dafny.IntOfInt64(624)).Times(_602_v5)).Minus((_646_v29).Cardinality())
						_ = _rhs141
						var _rhs142 bool = _645_v28
						_ = _rhs142
						var _lhs91 *GlobalState = globalState
						_ = _lhs91
						var _lhs92 *GlobalState = globalState
						_ = _lhs92
						_lhs91.F15 = _rhs138
						_602_v5 = _rhs139
						_645_v28 = _rhs140
						_lhs92.F7 = _rhs141
						_645_v28 = _rhs142
						var _647_v30 _dafny.Sequence
						_ = _647_v30
						_647_v30 = _dafny.SeqOf(p0)
						var _648_v31 _dafny.Set
						_ = _648_v31
						_648_v31 = _dafny.SetOf(p1, _645_v28)
						var _649_v32 _dafny.Int
						_ = _649_v32
						var _650_v33 _dafny.Int
						_ = _650_v33
						var _651_v34 _dafny.Int
						_ = _651_v34
						var _out6 _dafny.Int
						_ = _out6
						var _out7 _dafny.Int
						_ = _out7
						var _out8 _dafny.Int
						_ = _out8
						_out6, _out7, _out8 = (_this).M11(_dafny.ArrayCastTo((_641_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_641_v24), 0))).Int())), _647_v30, _648_v31, (_602_v5).Cmp(_dafny.IntOfInt64(-823)) <= 0, globalState)
						_649_v32 = _out6
						_650_v33 = _out7
						_651_v34 = _out8
					} else {
						var _652___mcc_h5 D2 = _source13.Get_().(D2_DC8).Cf11
						_ = _652___mcc_h5
						var _653_cf11 D2 = _652___mcc_h5
						_ = _653_cf11
						(globalState).F4 = _602_v5
						var _654_v35 _dafny.Set
						_ = _654_v35
						_654_v35 = _dafny.SetOf((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), _602_v5, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()))
						var _655_v36 _dafny.Sequence
						_ = _655_v36
						_655_v36 = _dafny.SeqOf((_596_v0).Cardinality())
						var _656_v38 _dafny.Sequence
						_ = _656_v38
						_656_v38 = _dafny.SeqOf(_654_v35, _dafny.SetOf((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)), func() _dafny.Set {
							var _coll32 = _dafny.NewBuilder()
							_ = _coll32
							for _iter32 := _dafny.Iterate((_655_v36).Elements()); ; {
								_compr_32, _ok32 := _iter32()
								if !_ok32 {
									break
								}
								var _657_v37 _dafny.Int
								_657_v37 = interface{}(_compr_32).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_655_v36, _657_v37) {
									_coll32.Add((_657_v37).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC20_(_dafny.IntOfInt64(315)), true)).Cardinality()))
								}
							}
							return _coll32.ToSet()
						}(), _dafny.SetOf((_this).Fm2((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), globalState)))
						var _658_v39 _dafny.Sequence
						_ = _658_v39
						_658_v39 = _dafny.SeqOf(_dafny.SetOf((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)), _654_v35, (_656_v38).Select((Companion_Default___.SafeIndex((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_656_v38).Cardinality()))).Uint32()).(_dafny.Set), _654_v35)
						(globalState).F13 = ((_658_v39).Select((Companion_Default___.SafeIndex(((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)).Plus((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_658_v39).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
						var _659_v40 _dafny.Map
						_ = _659_v40
						_659_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), _598_v2)
						var _660_v41 D5
						_ = _660_v41
						_660_v41 = Companion_D5_.Create_DC14_(_597_v1, p3, (_dafny.Zero).Minus(_602_v5), p0)
						_598_v2 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
							if (_659_v40).Contains(_602_v5) {
								return (_659_v40).Get(_602_v5).(_dafny.Sequence)
							}
							return _598_v2
						})(), _dafny.Companion_Sequence_.Update(_598_v2, (Companion_Default___.SafeIndex(_602_v5, _dafny.IntOfUint32((_598_v2).Cardinality()))).Uint32(), (_660_v41).Dtor_cf18()))
						var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((p2), 0))
						_ = _index98
						(p2).ArraySet1(Companion_Default___.SafeModuloInt(_602_v5, _602_v5), (_index98).Int())
					}
					var _661_v42 bool
					_ = _661_v42
					_661_v42 = true
					_661_v42 = p0
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _664_v43 _dafny.Int
		_ = _664_v43
		_664_v43 = _dafny.IntOfInt64(-298)
		var _665_v44 _dafny.Sequence
		_ = _665_v44
		_665_v44 = _dafny.SeqOf((_dafny.Zero).Minus(_664_v43), _dafny.IntOfInt64(580), _664_v43)
		var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((p2), 0))
		_ = _index99
		(p2).ArraySet1(_dafny.IntOfUint32((_665_v44).Cardinality()), (_index99).Int())
		var _666_v45 _dafny.Array
		_ = _666_v45
		var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(4))
		_ = _nw94
		_666_v45 = _nw94
		var _667_v46 _dafny.Map
		_ = _667_v46
		_667_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(632), p1)
		var _668_v47 _dafny.Sequence
		_ = _668_v47
		_668_v47 = _dafny.SeqOf(!((func() bool {
			if (_667_v46).Contains(_dafny.IntOfInt64(264)) {
				return (_667_v46).Get(_dafny.IntOfInt64(264)).(bool)
			}
			return p1
		})()), p3, p3)
		var _669_v48 D7
		_ = _669_v48
		_669_v48 = Companion_D7_.Create_DC22_(p3, _dafny.IntOfUint32((_668_v47).Cardinality()), p1, _dafny.IntOfUint32((_665_v44).Cardinality()))
		var _670_v49 _dafny.Sequence
		_ = _670_v49
		_670_v49 = _dafny.SeqOf(_669_v48, _669_v48)
		var _671_v50 _dafny.MultiSet
		_ = _671_v50
		_671_v50 = _dafny.MultiSetOf(_664_v43, _dafny.IntOfUint32((_670_v49).Cardinality()))
		var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_666_v45), 0))
		_ = _index100
		(_666_v45).ArraySet1(_671_v50, (_index100).Int())
		var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((p2), 0))
		_ = _index101
		var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_666_v45), 0))
		_ = _index102
		var _rhs143 _dafny.Int = _664_v43
		_ = _rhs143
		var _rhs144 _dafny.MultiSet = _671_v50
		_ = _rhs144
		var _lhs93 _dafny.Array = p2
		_ = _lhs93
		var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((p2), 0))
		_ = _lhs94
		var _lhs95 _dafny.Array = _666_v45
		_ = _lhs95
		var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_666_v45), 0))
		_ = _lhs96
		(_lhs93).ArraySet1(_rhs143, (_lhs94).Int())
		(_lhs95).ArraySet1(_rhs144, (_lhs96).Int())
		var _672_v51 _dafny.Array
		_ = _672_v51
		var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw95
		_672_v51 = _nw95
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_672_v51), 0))); ; {
			_guard_loop_0, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _673_i6 _dafny.Int
			_673_i6 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_673_i6).Sign() != -1) && ((_673_i6).Cmp(_dafny.ArrayLen((_672_v51), 0)) < 0)) {
				(_672_v51).ArraySet1((_673_i6).Times(_664_v43), (_673_i6).Int())
			}
		}
		(globalState).F15 = (_this).Fm2(_664_v43, globalState)
		var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((p2), 0))
		_ = _index103
		(p2).ArraySet1((func() _dafny.Int {
			if p1 {
				return _664_v43
			}
			return (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(911), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)
		})(), (_index103).Int())
		var _674_v52 bool
		_ = _674_v52
		_674_v52 = true
		_674_v52 = p0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	F24 D2
	F25 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this.F24 = Companion_D2_.Default()
	_this.F25 = false
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

func (_this *C3) Ctor__(f24 D2, f25 bool) {
	{
		(_this).F24 = f24
		(_this).F25 = f25
	}
}
func (_this *C3) Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return _this.F25
	}
}
func (_this *C3) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if !(_this.F25) {
				return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(328)))).Cardinality()
			}
			return (_dafny.Zero).Minus(_dafny.IntOfInt64(-540))
		})(), _dafny.IntOfInt64(714))
	}
}
func (_this *C3) M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _675_v0 _dafny.Int
		_ = _675_v0
		_675_v0 = _dafny.IntOfInt64(-140)
		var _676_v1 _dafny.Sequence
		_ = _676_v1
		_676_v1 = _dafny.SeqOf(_675_v0, _675_v0)
		var _677_v2 _dafny.Map
		_ = _677_v2
		_677_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_675_v0, (_676_v1).Select((Companion_Default___.SafeIndex(_675_v0, _dafny.IntOfUint32((_676_v1).Cardinality()))).Uint32()).(_dafny.Int))
		var _678_v3 _dafny.Map
		_ = _678_v3
		_678_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F25), _675_v0)
		var _679_v4 _dafny.CodePoint
		_ = _679_v4
		_679_v4 = _dafny.CodePoint('k')
		var _680_v5 _dafny.MultiSet
		_ = _680_v5
		_680_v5 = _dafny.MultiSetOf(_dafny.CodePoint('y'), _679_v4)
		var _681_v6 _dafny.Array
		_ = _681_v6
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_21
		var _nw96 _dafny.Array
		_ = _nw96
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw96 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) D1 = (func(_682_v0 _dafny.Int) func(_dafny.Int) D1 {
				return func(_683_i0 _dafny.Int) D1 {
					return Companion_D1_.Create_DC4_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_682_v0)))
				}
			})(_675_v0)
			_ = _init21
			var _element0_21 = _init21(_dafny.Zero)
			_ = _element0_21
			_nw96 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
			(_nw96).ArraySet1(_element0_21, 0)
			var _nativeLen0_21 = (_len0_21).Int()
			_ = _nativeLen0_21
			for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
				(_nw96).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
			}
		}
		_681_v6 = _nw96
		var _684_v7 _dafny.Map
		_ = _684_v7
		_684_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_681_v6, _this.F25)
		var _685_v8 _dafny.Map
		_ = _685_v8
		_685_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_675_v0, _this.F25)
		var _686_v9 _dafny.Sequence
		_ = _686_v9
		_686_v9 = _dafny.SeqOf(_this.F25, _this.F25, (func() bool {
			if (_685_v8).Contains(_675_v0) {
				return (_685_v8).Get(_675_v0).(bool)
			}
			return _this.F25
		})())
		var _687_v10 _dafny.MultiSet
		_ = _687_v10
		_687_v10 = _dafny.MultiSetOf(true, _this.F25, _this.F25, _this.F25, _this.F25)
		var _688_v11 _dafny.Array
		_ = _688_v11
		var _nwElement0_19 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-941), _675_v0)
		_ = _nwElement0_19
		var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(28))
		_ = _nw97
		(_nw97).ArraySet1(_nwElement0_19, 0)
		(_nw97).ArraySet1(_675_v0, 1)
		(_nw97).ArraySet1((func() _dafny.Int {
			if _this.F25 {
				return _675_v0
			}
			return _dafny.IntOfInt64(422)
		})(), 2)
		(_nw97).ArraySet1(_675_v0, 3)
		(_nw97).ArraySet1(((_677_v2).Cardinality()).Times(_675_v0), 4)
		(_nw97).ArraySet1(_675_v0, 5)
		(_nw97).ArraySet1(_675_v0, 6)
		(_nw97).ArraySet1((func() _dafny.Int {
			if (_678_v3).Contains(!(_this.F25)) {
				return (_678_v3).Get(!(_this.F25)).(_dafny.Int)
			}
			return _675_v0
		})(), 7)
		(_nw97).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((p0).Cardinality()), (func() _dafny.Int {
			if (_680_v5).Contains(_679_v4) {
				return (_680_v5).Multiplicity(_679_v4)
			}
			return _675_v0
		})()), 8)
		(_nw97).ArraySet1(_dafny.IntOfInt64(25), 9)
		(_nw97).ArraySet1(Companion_Default___.SafeDivisionInt(_675_v0, _675_v0), 10)
		(_nw97).ArraySet1(_dafny.IntOfInt64(-127), 11)
		(_nw97).ArraySet1(_675_v0, 12)
		(_nw97).ArraySet1(_675_v0, 13)
		(_nw97).ArraySet1(_dafny.IntOfInt64(19), 14)
		(_nw97).ArraySet1((_684_v7).Cardinality(), 15)
		(_nw97).ArraySet1((_dafny.Zero).Minus(_675_v0), 16)
		(_nw97).ArraySet1(_675_v0, 17)
		(_nw97).ArraySet1(_675_v0, 18)
		(_nw97).ArraySet1(_dafny.IntOfUint32((_686_v9).Cardinality()), 19)
		(_nw97).ArraySet1(_675_v0, 20)
		(_nw97).ArraySet1((_675_v0).Plus(_dafny.IntOfInt64(652)), 21)
		(_nw97).ArraySet1(_675_v0, 22)
		(_nw97).ArraySet1(_dafny.IntOfUint32((_686_v9).Cardinality()), 23)
		(_nw97).ArraySet1((func() _dafny.Int {
			if (_687_v10).Contains(_this.F25) {
				return (_687_v10).Multiplicity(_this.F25)
			}
			return _675_v0
		})(), 24)
		(_nw97).ArraySet1(Companion_Default___.SafeModuloInt(_675_v0, _675_v0), 25)
		(_nw97).ArraySet1(_675_v0, 26)
		(_nw97).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_675_v0), _675_v0), 27)
		_688_v11 = _nw97
		var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_688_v11), 0))
		_ = _index104
		(_688_v11).ArraySet1(_675_v0, (_index104).Int())
		var _689_v12 _dafny.Array
		_ = _689_v12
		var _nwElement0_20 _dafny.Sequence = _dafny.SeqOf(_675_v0, _dafny.IntOfInt64(572), _675_v0)
		_ = _nwElement0_20
		var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(17))
		_ = _nw98
		(_nw98).ArraySet1(_nwElement0_20, 0)
		(_nw98).ArraySet1(_676_v1, 1)
		(_nw98).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(440))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg48 _dafny.Int) interface{} {
				return coer48(arg48)
			}
		}((func(_690_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_691_i1 _dafny.Int) _dafny.Int {
				return _690_v0
			}
		})(_675_v0))), 2)
		(_nw98).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(969)), 3)
		(_nw98).ArraySet1(_dafny.Companion_Sequence_.Update(_676_v1, (Companion_Default___.SafeIndex(_675_v0, _dafny.IntOfUint32((_676_v1).Cardinality()))).Uint32(), _675_v0), 4)
		(_nw98).ArraySet1(_676_v1, 5)
		(_nw98).ArraySet1(_dafny.SeqOf(_675_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25, (_this).Fm2((_678_v3).Cardinality(), globalState))).Cardinality()), 6)
		(_nw98).ArraySet1(_676_v1, 7)
		(_nw98).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nht")).Cardinality())), 8)
		(_nw98).ArraySet1(_dafny.SeqOf(_675_v0, _675_v0), 9)
		(_nw98).ArraySet1(_676_v1, 10)
		(_nw98).ArraySet1((Companion_D6_.Create_DC16_(_676_v1)).Dtor_cf23(), 11)
		(_nw98).ArraySet1(_676_v1, 12)
		(_nw98).ArraySet1(_676_v1, 13)
		(_nw98).ArraySet1(_dafny.Companion_Sequence_.Update(_676_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_676_v1).Cardinality()))).Uint32(), _675_v0), 14)
		(_nw98).ArraySet1(_676_v1, 15)
		(_nw98).ArraySet1(_676_v1, 16)
		_689_v12 = _nw98
		var _692_v13 _dafny.Sequence
		_ = _692_v13
		_692_v13 = _dafny.SeqOf(_689_v12, _689_v12, _689_v12)
		var _693_v14 _dafny.Array
		_ = _693_v14
		var _nwElement0_21 _dafny.Array = (func() _dafny.Array {
			if _this.F25 {
				return _689_v12
			}
			return _689_v12
		})()
		_ = _nwElement0_21
		var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(15))
		_ = _nw99
		(_nw99).ArraySet1(_nwElement0_21, 0)
		(_nw99).ArraySet1((_692_v13).Select((Companion_Default___.SafeIndex(_675_v0, _dafny.IntOfUint32((_692_v13).Cardinality()))).Uint32()).(_dafny.Array), 1)
		(_nw99).ArraySet1((func() _dafny.Array {
			if _this.F25 {
				return _689_v12
			}
			return _689_v12
		})(), 2)
		(_nw99).ArraySet1(_689_v12, 3)
		(_nw99).ArraySet1(_689_v12, 4)
		(_nw99).ArraySet1(_689_v12, 5)
		(_nw99).ArraySet1(_689_v12, 6)
		(_nw99).ArraySet1(_689_v12, 7)
		(_nw99).ArraySet1(_689_v12, 8)
		(_nw99).ArraySet1(_689_v12, 9)
		(_nw99).ArraySet1(_689_v12, 10)
		(_nw99).ArraySet1(_689_v12, 11)
		(_nw99).ArraySet1(_689_v12, 12)
		(_nw99).ArraySet1(_689_v12, 13)
		(_nw99).ArraySet1(_689_v12, 14)
		_693_v14 = _nw99
		var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_693_v14), 0))
		_ = _index105
		(_693_v14).ArraySet1(_689_v12, (_index105).Int())
		var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_688_v11), 0))
		_ = _index106
		var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_693_v14), 0))
		_ = _index107
		var _rhs145 _dafny.Int = _dafny.IntOfInt64(250)
		_ = _rhs145
		var _rhs146 _dafny.Int = _675_v0
		_ = _rhs146
		var _rhs147 _dafny.Array = _689_v12
		_ = _rhs147
		var _rhs148 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((Companion_D1_.Create_DC2_(_675_v0)).Dtor_cf2()), _675_v0)
		_ = _rhs148
		var _lhs97 _dafny.Array = _688_v11
		_ = _lhs97
		var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_688_v11), 0))
		_ = _lhs98
		var _lhs99 *GlobalState = globalState
		_ = _lhs99
		var _lhs100 _dafny.Array = _693_v14
		_ = _lhs100
		var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(229), _dafny.ArrayLen((_693_v14), 0))
		_ = _lhs101
		(_lhs97).ArraySet1(_rhs145, (_lhs98).Int())
		_lhs99.F15 = _rhs146
		(_lhs100).ArraySet1(_rhs147, (_lhs101).Int())
		_675_v0 = _rhs148
		_675_v0 = _675_v0
		var _694_v15 _dafny.Array
		_ = _694_v15
		var _nw100 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw100
		_694_v15 = _nw100
		var _695_v16 _dafny.Sequence
		_ = _695_v16
		_695_v16 = _dafny.SeqOf(_694_v15)
		var _696_v17 D7
		_ = _696_v17
		_696_v17 = Companion_D7_.Create_DC19_(_695_v16)
		var _rhs149 bool = _this.F25
		_ = _rhs149
		var _rhs150 _dafny.CodePoint = _679_v4
		_ = _rhs150
		var _rhs151 _dafny.Sequence = (_696_v17).Dtor_cf28()
		_ = _rhs151
		var _lhs102 *C3 = _this
		_ = _lhs102
		_lhs102.F25 = _rhs149
		_679_v4 = _rhs150
		_695_v16 = _rhs151
		var _hi4 _dafny.Int = (func() _dafny.Int {
			if _this.F25 {
				return _675_v0
			}
			return (_688_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_688_v11), 0))).Int()).(_dafny.Int)
		})()
		_ = _hi4
		for _697_i2 := ((_677_v2).Cardinality()).Minus((_dafny.Zero).Minus(_675_v0)); _697_i2.Cmp(_hi4) < 0; _697_i2 = _697_i2.Plus(_dafny.One) {
			var _698_v18 _dafny.Array
			_ = _698_v18
			var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
			_ = _nw101
			_698_v18 = _nw101
			var _699_v19 _dafny.Sequence
			_ = _699_v19
			_699_v19 = _dafny.SeqOf(_698_v18, _698_v18, _698_v18)
			var _700_v20 _dafny.Array
			_ = _700_v20
			var _nwElement0_22 _dafny.Array = (_699_v19).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_699_v19).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _nwElement0_22
			var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(29))
			_ = _nw102
			(_nw102).ArraySet1(_nwElement0_22, 0)
			(_nw102).ArraySet1(_698_v18, 1)
			(_nw102).ArraySet1(_698_v18, 2)
			(_nw102).ArraySet1(_698_v18, 3)
			(_nw102).ArraySet1(_698_v18, 4)
			(_nw102).ArraySet1(_698_v18, 5)
			(_nw102).ArraySet1(_698_v18, 6)
			(_nw102).ArraySet1(_698_v18, 7)
			(_nw102).ArraySet1(_698_v18, 8)
			(_nw102).ArraySet1((_699_v19).Select((Companion_Default___.SafeIndex((_688_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_688_v11), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_699_v19).Cardinality()))).Uint32()).(_dafny.Array), 9)
			(_nw102).ArraySet1(_698_v18, 10)
			(_nw102).ArraySet1(_698_v18, 11)
			(_nw102).ArraySet1(_698_v18, 12)
			(_nw102).ArraySet1(_698_v18, 13)
			(_nw102).ArraySet1(_698_v18, 14)
			(_nw102).ArraySet1((func() _dafny.Array {
				if true {
					return _698_v18
				}
				return _698_v18
			})(), 15)
			(_nw102).ArraySet1(_698_v18, 16)
			(_nw102).ArraySet1(_698_v18, 17)
			(_nw102).ArraySet1(_698_v18, 18)
			(_nw102).ArraySet1(_698_v18, 19)
			(_nw102).ArraySet1(_698_v18, 20)
			(_nw102).ArraySet1(_698_v18, 21)
			(_nw102).ArraySet1(_698_v18, 22)
			(_nw102).ArraySet1(_698_v18, 23)
			(_nw102).ArraySet1(_698_v18, 24)
			(_nw102).ArraySet1(_698_v18, 25)
			(_nw102).ArraySet1(_698_v18, 26)
			(_nw102).ArraySet1(_698_v18, 27)
			(_nw102).ArraySet1(_698_v18, 28)
			_700_v20 = _nw102
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_700_v20), 0))
			_ = _index108
			(_700_v20).ArraySet1(_698_v18, (_index108).Int())
			var _701_v21 D2
			_ = _701_v21
			_701_v21 = Companion_D2_.Create_DC6_((_688_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_688_v11), 0))).Int()).(_dafny.Int))
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_700_v20), 0))
			_ = _index109
			var _rhs152 _dafny.Array = _698_v18
			_ = _rhs152
			var _rhs153 bool = _this.F25
			_ = _rhs153
			var _rhs154 bool = true
			_ = _rhs154
			var _rhs155 D2 = _701_v21
			_ = _rhs155
			var _lhs103 _dafny.Array = _700_v20
			_ = _lhs103
			var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_700_v20), 0))
			_ = _lhs104
			var _lhs105 *C3 = _this
			_ = _lhs105
			var _lhs106 *C3 = _this
			_ = _lhs106
			(_lhs103).ArraySet1(_rhs152, (_lhs104).Int())
			_lhs105.F25 = _rhs153
			_lhs106.F25 = _rhs154
			_701_v21 = _rhs155
			(globalState).F4 = _675_v0
			(_this).F25 = _this.F25
			(_this).F25 = !((_675_v0).Cmp(_dafny.IntOfUint32((p0).Cardinality())) == 0) || (_this.F25)
		}
		_675_v0 = (_688_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_688_v11), 0))).Int()).(_dafny.Int)
		(_this).F25 = _this.F25
	}
}
func (_this *C3) M10(globalState *GlobalState) (_dafny.Set, _dafny.Array, _dafny.Sequence) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _702_v0 _dafny.Array
		_ = _702_v0
		var _len0_22 _dafny.Int = _dafny.One
		_ = _len0_22
		var _nw103 _dafny.Array
		_ = _nw103
		if _len0_22.Cmp(_dafny.Zero) == 0 {
			_nw103 = _dafny.NewArray(_len0_22)
		} else {
			var _init22 func(_dafny.Int) bool = func(_703_i0 _dafny.Int) bool {
				return true
			}
			_ = _init22
			var _element0_22 = _init22(_dafny.Zero)
			_ = _element0_22
			_nw103 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
			(_nw103).ArraySet1(_element0_22, 0)
			var _nativeLen0_22 = (_len0_22).Int()
			_ = _nativeLen0_22
			for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
				(_nw103).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
			}
		}
		_702_v0 = _nw103
		var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
		_ = _index110
		(_702_v0).ArraySet1(!((func() bool {
			if _this.F25 {
				return _this.F25
			}
			return _this.F25
		})()), (_index110).Int())
		var _704_v1 _dafny.Int
		_ = _704_v1
		_704_v1 = _dafny.IntOfInt64(199)
		var _705_v2 _dafny.Set
		_ = _705_v2
		_705_v2 = _dafny.SetOf(_704_v1)
		var _706_v3 _dafny.Map
		_ = _706_v3
		_706_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_704_v1, _704_v1, (_705_v2).Cardinality(), _704_v1, _704_v1)).Union(_705_v2), _704_v1)
		var _707_v4 _dafny.CodePoint
		_ = _707_v4
		_707_v4 = _dafny.CodePoint('v')
		var _708_v5 _dafny.Sequence
		_ = _708_v5
		_708_v5 = _dafny.SeqOf(_this.F25)
		var _709_v6 D8
		_ = _709_v6
		_709_v6 = Companion_D8_.Create_DC23_(_dafny.SeqOf(_this.F25, _this.F25))
		var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
		_ = _index111
		var _rhs156 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_708_v5, (_709_v6).Dtor_cf39()), _dafny.SeqOf(_this.F25, _this.F25))
		_ = _rhs156
		var _rhs157 _dafny.Int = ((_dafny.Zero).Minus(_704_v1)).Plus(_704_v1)
		_ = _rhs157
		var _rhs158 bool = (Companion_Default___.SafeModuloInt(_704_v1, _704_v1)).Cmp(_dafny.IntOfInt64(-740)) != 0
		_ = _rhs158
		var _rhs159 _dafny.Map = _706_v3
		_ = _rhs159
		var _rhs160 _dafny.CodePoint = _707_v4
		_ = _rhs160
		var _lhs107 _dafny.Array = _702_v0
		_ = _lhs107
		var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
		_ = _lhs108
		var _lhs109 *GlobalState = globalState
		_ = _lhs109
		var _lhs110 *C3 = _this
		_ = _lhs110
		(_lhs107).ArraySet1(_rhs156, (_lhs108).Int())
		_lhs109.F4 = _rhs157
		_lhs110.F25 = _rhs158
		_706_v3 = _rhs159
		_707_v4 = _rhs160
		var _710_v7 _dafny.Array
		_ = _710_v7
		var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
		_ = _nw104
		_710_v7 = _nw104
		var _711_v8 _dafny.Array
		_ = _711_v8
		var _nwElement0_23 _dafny.Array = _710_v7
		_ = _nwElement0_23
		var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(18))
		_ = _nw105
		(_nw105).ArraySet1(_nwElement0_23, 0)
		(_nw105).ArraySet1(_710_v7, 1)
		(_nw105).ArraySet1(_710_v7, 2)
		(_nw105).ArraySet1(_710_v7, 3)
		(_nw105).ArraySet1(_710_v7, 4)
		(_nw105).ArraySet1(_710_v7, 5)
		(_nw105).ArraySet1(_710_v7, 6)
		(_nw105).ArraySet1(_710_v7, 7)
		(_nw105).ArraySet1(_710_v7, 8)
		(_nw105).ArraySet1(_710_v7, 9)
		(_nw105).ArraySet1(_710_v7, 10)
		(_nw105).ArraySet1(_710_v7, 11)
		(_nw105).ArraySet1(_710_v7, 12)
		(_nw105).ArraySet1(_710_v7, 13)
		(_nw105).ArraySet1(_710_v7, 14)
		(_nw105).ArraySet1(_710_v7, 15)
		(_nw105).ArraySet1(_710_v7, 16)
		(_nw105).ArraySet1(_710_v7, 17)
		_711_v8 = _nw105
		_711_v8 = _711_v8
		var _712_v9 _dafny.Map
		_ = _712_v9
		_712_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_702_v0, _702_v0)
		var _713_v10 _dafny.Sequence
		_ = _713_v10
		_713_v10 = _dafny.SeqOf(_704_v1, (_712_v9).Cardinality())
		(_this).F25 = !_dafny.Companion_Sequence_.Contains(_713_v10, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _this.F25 {
				return _713_v10
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(159))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}((func(_714_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_715_i1 _dafny.Int) _dafny.Int {
					return _714_v1
				}
			})(_704_v1)))
		})()).Cardinality()))
		var _716_v11 _dafny.MultiSet
		_ = _716_v11
		_716_v11 = _dafny.MultiSetOf(_this.F25)
		var _717_v12 _dafny.Sequence
		_ = _717_v12
		_717_v12 = _dafny.UnicodeSeqOfUtf8Bytes("bpoxflb")
		var _718_v13 D7
		_ = _718_v13
		_718_v13 = Companion_D7_.Create_DC21_((func() _dafny.Int {
			if (_716_v11).Contains(_this.F25) {
				return (_716_v11).Multiplicity(_this.F25)
			}
			return _dafny.IntOfUint32((_717_v12).Cardinality())
		})(), _704_v1, false, !((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool)) || (_this.F25), false)
		var _source14 D7 = _718_v13
		_ = _source14
		if _source14.Is_DC20() {
			var _719___mcc_h0 _dafny.Int = _source14.Get_().(D7_DC20).Cf29
			_ = _719___mcc_h0
			var _720_cf29 _dafny.Int = _719___mcc_h0
			_ = _720_cf29
			if (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool) {
				var _721_v14 *C1
				_ = _721_v14
				var _nw106 *C1 = New_C1_()
				_ = _nw106
				_nw106.Ctor__()
				_721_v14 = _nw106
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_710_v7), 0))
				_ = _index112
				(_710_v7).ArraySet1(Companion_Default___.SafeModuloInt(_704_v1, _dafny.IntOfUint32((_708_v5).Cardinality())), (_index112).Int())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
				_ = _index113
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_710_v7), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
				_ = _index115
				var _rhs161 bool = _this.F25
				_ = _rhs161
				var _rhs162 bool = (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool)
				_ = _rhs162
				var _rhs163 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(788))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_722_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_723_i2 _dafny.Int) _dafny.CodePoint {
						return _722_v4
					}
				})(_707_v4))), _dafny.UnicodeSeqOfUtf8Bytes("tswns"))).Cardinality())
				_ = _rhs163
				var _rhs164 _dafny.Array = _710_v7
				_ = _rhs164
				var _rhs165 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm39(_704_v1, _707_v4, _713_v10, !(_this.F25), globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_720_cf29, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(318))).Uint32(), func(coer51 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_724_v2 _dafny.Set) func(_dafny.Int) D11 {
					return func(_725_i3 _dafny.Int) D11 {
						return Companion_D11_.Create_DC28_(_724_v2)
					}
				})(_705_v2)))).Cardinality())), _dafny.IntOfUint32((_717_v12).Cardinality())), _713_v10))
				_ = _rhs165
				var _lhs111 _dafny.Array = _702_v0
				_ = _lhs111
				var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
				_ = _lhs112
				var _lhs113 *C3 = _this
				_ = _lhs113
				var _lhs114 _dafny.Array = _710_v7
				_ = _lhs114
				var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_710_v7), 0))
				_ = _lhs115
				var _lhs116 _dafny.Array = _702_v0
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
				_ = _lhs117
				(_lhs111).ArraySet1(_rhs161, (_lhs112).Int())
				_lhs113.F25 = _rhs162
				(_lhs114).ArraySet1(_rhs163, (_lhs115).Int())
				r1 = _rhs164
				(_lhs116).ArraySet1(_rhs165, (_lhs117).Int())
				(_this).F25 = false
				var _726_v15 D1
				_ = _726_v15
				_726_v15 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_713_v10).Cardinality()))
				var _pat_let_tv8 = _704_v1
				_ = _pat_let_tv8
				var _727_v16 _dafny.Array
				_ = _727_v16
				var _nwElement0_24 D1 = _726_v15
				_ = _nwElement0_24
				var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(20))
				_ = _nw107
				(_nw107).ArraySet1(_nwElement0_24, 0)
				(_nw107).ArraySet1(_726_v15, 1)
				(_nw107).ArraySet1(_726_v15, 2)
				(_nw107).ArraySet1(_726_v15, 3)
				(_nw107).ArraySet1(Companion_D1_.Create_DC2_(_704_v1), 4)
				(_nw107).ArraySet1(_726_v15, 5)
				(_nw107).ArraySet1(Companion_D1_.Create_DC2_(_720_cf29), 6)
				(_nw107).ArraySet1(_726_v15, 7)
				(_nw107).ArraySet1(Companion_D1_.Create_DC2_((_713_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_717_v12, (Companion_Default___.SafeIndex((_705_v2).Cardinality(), _dafny.IntOfUint32((_717_v12).Cardinality()))).Uint32(), _707_v4)).Cardinality()), _dafny.IntOfUint32((_713_v10).Cardinality()))).Uint32()).(_dafny.Int)), 8)
				(_nw107).ArraySet1(Companion_D1_.Create_DC2_(_720_cf29), 9)
				(_nw107).ArraySet1(_726_v15, 10)
				(_nw107).ArraySet1(_726_v15, 11)
				(_nw107).ArraySet1(_726_v15, 12)
				(_nw107).ArraySet1(_726_v15, 13)
				(_nw107).ArraySet1(_726_v15, 14)
				(_nw107).ArraySet1(func(_pat_let2_0 D1) D1 {
					return func(_728_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let3_0 _dafny.Int) D1 {
							return func(_729_dt__update_hcf2_h0 _dafny.Int) D1 {
								return Companion_D1_.Create_DC2_(_729_dt__update_hcf2_h0)
							}(_pat_let3_0)
						}(_pat_let_tv8)
					}(_pat_let2_0)
				}(Companion_D1_.Create_DC2_((_710_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_710_v7), 0))).Int()).(_dafny.Int))), 15)
				(_nw107).ArraySet1(_726_v15, 16)
				(_nw107).ArraySet1(_726_v15, 17)
				(_nw107).ArraySet1(_726_v15, 18)
				(_nw107).ArraySet1(Companion_Default___.Fm40((_710_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_710_v7), 0))).Int()).(_dafny.Int), (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool), _this.F25, _709_v6, globalState), 19)
				_727_v16 = _nw107
				_727_v16 = _727_v16
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_710_v7), 0))
				_ = _index116
				(_710_v7).ArraySet1(_dafny.IntOfInt64(603), (_index116).Int())
			} else {
				var _730_v17 _dafny.Map
				_ = _730_v17
				_730_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_705_v2, _this.F25)
				_713_v10 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (func() bool {
						if (_730_v17).Contains(_705_v2) {
							return (_730_v17).Get(_705_v2).(bool)
						}
						return Companion_Default___.Fm26(globalState)
					})() {
						return _713_v10
					}
					return _713_v10
				})(), _dafny.SeqOf(_704_v1))
				var _731_v18 _dafny.Map
				_ = _731_v18
				_731_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v1, (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool))
				var _732_v19 _dafny.Map
				_ = _732_v19
				_732_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_717_v12, _720_cf29)
				_731_v18 = (_731_v18).Update(Companion_Default___.SafeModuloInt((_732_v19).Cardinality(), _704_v1), true)
				_708_v5 = Companion_Default___.Fm27(globalState)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_710_v7), 0))
				_ = _index117
				(_710_v7).ArraySet1(_704_v1, (_index117).Int())
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_710_v7), 0))
				_ = _index118
				(_710_v7).ArraySet1(_720_cf29, (_index118).Int())
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_710_v7), 0))
				_ = _index119
				(_710_v7).ArraySet1(Companion_Default___.SafeModuloInt(_720_cf29, _704_v1), (_index119).Int())
			}
			var _733_v20 _dafny.Set
			_ = _733_v20
			_733_v20 = _dafny.SetOf(Companion_Default___.Fm41(_707_v4, globalState))
			var _734_v21 _dafny.Map
			_ = _734_v21
			_734_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_733_v20, _707_v4)
			_734_v21 = (_734_v21).Update(_733_v20, _707_v4)
			var _735_v22 _dafny.Array
			_ = _735_v22
			var _nw108 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(12))
			_ = _nw108
			_735_v22 = _nw108
			var _736_v23 D3
			_ = _736_v23
			_736_v23 = Companion_D3_.Create_DC10_(_dafny.IntOfInt64(808), _720_cf29)
			var _737_v24 _dafny.Map
			_ = _737_v24
			_737_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_735_v22, _736_v23)
			(_this).F25 = ((_737_v24).Cardinality()).Cmp(_704_v1) == 0
			(globalState).F4 = ((_716_v11).Cardinality()).Minus(_704_v1)
		} else if _source14.Is_DC21() {
			var _738___mcc_h1 _dafny.Int = _source14.Get_().(D7_DC21).Cf30
			_ = _738___mcc_h1
			var _739___mcc_h2 _dafny.Int = _source14.Get_().(D7_DC21).Cf31
			_ = _739___mcc_h2
			var _740___mcc_h3 bool = _source14.Get_().(D7_DC21).Cf32
			_ = _740___mcc_h3
			var _741___mcc_h4 bool = _source14.Get_().(D7_DC21).Cf33
			_ = _741___mcc_h4
			var _742___mcc_h5 bool = _source14.Get_().(D7_DC21).Cf34
			_ = _742___mcc_h5
			var _743_cf34 bool = _742___mcc_h5
			_ = _743_cf34
			var _744_cf33 bool = _741___mcc_h4
			_ = _744_cf33
			var _745_cf32 bool = _740___mcc_h3
			_ = _745_cf32
			var _746_cf31 _dafny.Int = _739___mcc_h2
			_ = _746_cf31
			var _747_cf30 _dafny.Int = _738___mcc_h1
			_ = _747_cf30
			var _pat_let_tv9 = _705_v2
			_ = _pat_let_tv9
			var _source15 D11 = func(_pat_let4_0 D11) D11 {
				return func(_748_dt__update__tmp_h1 D11) D11 {
					return func(_pat_let5_0 _dafny.Set) D11 {
						return func(_749_dt__update_hcf45_h0 _dafny.Set) D11 {
							return Companion_D11_.Create_DC28_(_749_dt__update_hcf45_h0)
						}(_pat_let5_0)
					}(_pat_let_tv9)
				}(_pat_let4_0)
			}(Companion_D11_.Create_DC28_(_705_v2))
			_ = _source15
			if _source15.Is_DC29() {
				(globalState).F4 = (_747_cf30).Times(_747_cf30)
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
				_ = _index120
				(_702_v0).ArraySet1((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool), (_index120).Int())
				var _750_v25 *C2
				_ = _750_v25
				var _nw109 *C2 = New_C2_()
				_ = _nw109
				_nw109.Ctor__()
				_750_v25 = _nw109
				var _751_v26 _dafny.Map
				_ = _751_v26
				_751_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v1, _745_cf32)
				var _752_v27 D7
				_ = _752_v27
				_752_v27 = Companion_D7_.Create_DC20_(_dafny.IntOfInt64(316))
				var _rhs166 _dafny.Int = _746_cf31
				_ = _rhs166
				var _rhs167 bool = !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(true, !(_this.F25)), _dafny.Companion_Sequence_.Update(_708_v5, (Companion_Default___.SafeIndex((_752_v27).Dtor_cf29(), _dafny.IntOfUint32((_708_v5).Cardinality()))).Uint32(), (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool)))
				_ = _rhs167
				var _rhs168 _dafny.Map = _751_v26
				_ = _rhs168
				var _rhs169 bool = !(_this.F25) || (_743_cf34)
				_ = _rhs169
				var _rhs170 bool = _745_cf32
				_ = _rhs170
				var _lhs118 *GlobalState = globalState
				_ = _lhs118
				_lhs118.F4 = _rhs166
				_745_cf32 = _rhs167
				_751_v26 = _rhs168
				_745_cf32 = _rhs169
				_745_cf32 = _rhs170
			} else if _source15.Is_DC30() {
				var _753___mcc_h11 bool = _source15.Get_().(D11_DC30).Cf46
				_ = _753___mcc_h11
				var _754___mcc_h12 _dafny.Int = _source15.Get_().(D11_DC30).Cf47
				_ = _754___mcc_h12
				var _755_cf47 _dafny.Int = _754___mcc_h12
				_ = _755_cf47
				var _756_cf46 bool = _753___mcc_h11
				_ = _756_cf46
				var _757_v28 _dafny.Map
				_ = _757_v28
				_757_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-980), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_743_cf34, _704_v1)).Cardinality())
				var _758_v29 _dafny.Map
				_ = _758_v29
				_758_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_cf32, (func() _dafny.Int {
					if (_757_v28).Contains(_746_cf31) {
						return (_757_v28).Get(_746_cf31).(_dafny.Int)
					}
					return _704_v1
				})())).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(59))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_759_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_760_i4 _dafny.Int) _dafny.CodePoint {
						return _759_v4
					}
				})(_707_v4))))
				var _761_v30 D1
				_ = _761_v30
				_761_v30 = Companion_D1_.Create_DC2_(_747_cf30)
				var _762_v31 _dafny.Set
				_ = _762_v31
				_762_v31 = _dafny.SetOf(_761_v30)
				var _763_v32 _dafny.Map
				_ = _763_v32
				_763_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_758_v29).Merge(_758_v29), _762_v31)
				var _764_v33 _dafny.Map
				_ = _764_v33
				_764_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_755_cf47, _762_v31)
				var _765_v34 _dafny.MultiSet
				_ = _765_v34
				_765_v34 = _dafny.MultiSetOf(_761_v30)
				_763_v32 = (_763_v32).Update(_758_v29, (func() _dafny.Set {
					if (_764_v33).Contains(_747_cf30) {
						return (_764_v33).Get(_747_cf30).(_dafny.Set)
					}
					return func() _dafny.Set {
						var _coll33 = _dafny.NewBuilder()
						_ = _coll33
						for _iter34 := _dafny.Iterate((_765_v34).Elements()); ; {
							_compr_33, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _766_v35 D1
							_766_v35 = interface{}(_compr_33).(D1)
							if (_765_v34).Contains(_766_v35) {
								_coll33.Add(_766_v35)
							}
						}
						return _coll33.ToSet()
					}()
				})())
				(globalState).F4 = _dafny.IntOfInt64(897)
				var _767_v36 *C2
				_ = _767_v36
				var _nw110 *C2 = New_C2_()
				_ = _nw110
				_nw110.Ctor__()
				_767_v36 = _nw110
				var _768_v37 _dafny.Map
				_ = _768_v37
				_768_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool), true)
				var _769_v38 _dafny.Map
				_ = _769_v38
				_769_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_768_v37, _702_v0)
				var _770_v39 _dafny.Array
				_ = _770_v39
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_23
				var _nw111 _dafny.Array
				_ = _nw111
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw111 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Sequence = (func(_771_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_772_i5 _dafny.Int) _dafny.Sequence {
							return _771_v5
						}
					})(_708_v5)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw111 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw111).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw111).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_770_v39 = _nw111
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_770_v39), 0))
				_ = _index121
				(_770_v39).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F25, true), _708_v5), (_index121).Int())
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_770_v39), 0))
				_ = _index122
				var _rhs171 _dafny.Map = _769_v38
				_ = _rhs171
				var _rhs172 _dafny.Sequence = Companion_Default___.Fm27(globalState)
				_ = _rhs172
				var _lhs119 _dafny.Array = _770_v39
				_ = _lhs119
				var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_770_v39), 0))
				_ = _lhs120
				_769_v38 = _rhs171
				(_lhs119).ArraySet1(_rhs172, (_lhs120).Int())
			} else if _source15.Is_DC28() {
				var _773___mcc_h13 _dafny.Set = _source15.Get_().(D11_DC28).Cf45
				_ = _773___mcc_h13
				var _774_cf45 _dafny.Set = _773___mcc_h13
				_ = _774_cf45
				(globalState).F15 = (_dafny.Zero).Minus(_704_v1)
				(globalState).F13 = _704_v1
				_745_cf32 = _this.F25
				(globalState).F7 = _746_cf31
			} else {
				var _775___mcc_h14 D11 = _source15.Get_().(D11_DC31).Cf48
				_ = _775___mcc_h14
				var _776_cf48 D11 = _775___mcc_h14
				_ = _776_cf48
				var _777_v40 _dafny.Map
				_ = _777_v40
				_777_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_744_cf33, _744_cf33))
				var _778_v41 _dafny.Map
				_ = _778_v41
				_778_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm1(_777_v40, _dafny.IntOfUint32((_717_v12).Cardinality()), globalState), _743_cf34)
				var _779_v42 _dafny.Map
				_ = _779_v42
				_779_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_744_cf33, _778_v41)
				_779_v42 = (_779_v42).Update(!((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool)), _778_v41)
				(_this).F25 = (_708_v5).Select((Companion_Default___.SafeIndex(_746_cf31, _dafny.IntOfUint32((_708_v5).Cardinality()))).Uint32()).(bool)
				var _780_v43 _dafny.MultiSet
				_ = _780_v43
				_780_v43 = _dafny.MultiSetOf(_dafny.IntOfUint32((_708_v5).Cardinality()))
				var _781_v44 D1
				_ = _781_v44
				_781_v44 = Companion_D1_.Create_DC4_(_780_v43)
				_781_v44 = Companion_Default___.Fm42(_747_cf30, (func() _dafny.Set {
					var _coll34 = _dafny.NewBuilder()
					_ = _coll34
					for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(376), _dafny.IntOfInt64(392))); ; {
						_compr_34, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _782_v45 _dafny.Int
						_782_v45 = interface{}(_compr_34).(_dafny.Int)
						if ((_dafny.IntOfInt64(376)).Cmp(_782_v45) <= 0) && ((_782_v45).Cmp(_dafny.IntOfInt64(392)) < 0) {
							_coll34.Add((_782_v45).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_747_cf30, _747_cf30)).Cardinality()))
						}
					}
					return _coll34.ToSet()
				}()).Cardinality(), globalState)
				var _783_v46 D0
				_ = _783_v46
				_783_v46 = Companion_D0_.Create_DC0_(_702_v0)
				var _784_v47 _dafny.Sequence
				_ = _784_v47
				_784_v47 = _dafny.SeqOf(_783_v46)
				(_this).F25 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_784_v47, _784_v47), _784_v47)
			}
			_704_v1 = _dafny.IntOfInt64(698)
			var _785_v48 _dafny.MultiSet
			_ = _785_v48
			_785_v48 = _dafny.MultiSetOf(_705_v2, _705_v2, _705_v2)
			var _source16 D5 = Companion_D5_.Create_DC13_((_785_v48).Union(_dafny.MultiSetOf(_705_v2, _705_v2, _705_v2)))
			_ = _source16
			if _source16.Is_DC14() {
				var _786___mcc_h15 _dafny.CodePoint = _source16.Get_().(D5_DC14).Cf18
				_ = _786___mcc_h15
				var _787___mcc_h16 bool = _source16.Get_().(D5_DC14).Cf19
				_ = _787___mcc_h16
				var _788___mcc_h17 _dafny.Int = _source16.Get_().(D5_DC14).Cf20
				_ = _788___mcc_h17
				var _789___mcc_h18 bool = _source16.Get_().(D5_DC14).Cf21
				_ = _789___mcc_h18
				var _790_cf21 bool = _789___mcc_h18
				_ = _790_cf21
				var _791_cf20 _dafny.Int = _788___mcc_h17
				_ = _791_cf20
				var _792_cf19 bool = _787___mcc_h16
				_ = _792_cf19
				var _793_cf18 _dafny.CodePoint = _786___mcc_h15
				_ = _793_cf18
				_790_cf21 = _745_cf32
				var _794_v49 *C0
				_ = _794_v49
				var _nw112 *C0 = New_C0_()
				_ = _nw112
				_nw112.Ctor__()
				_794_v49 = _nw112
				var _795_v50 *C2
				_ = _795_v50
				var _nw113 *C2 = New_C2_()
				_ = _nw113
				_nw113.Ctor__()
				_795_v50 = _nw113
				(globalState).F4 = (func() _dafny.Int {
					if (_716_v11).Contains(_792_cf19) {
						return (_716_v11).Multiplicity(_792_cf19)
					}
					return (func() _dafny.Int {
						if _792_cf19 {
							return (_705_v2).Cardinality()
						}
						return _746_cf31
					})()
				})()
			} else if _source16.Is_DC13() {
				var _796___mcc_h19 _dafny.MultiSet = _source16.Get_().(D5_DC13).Cf17
				_ = _796___mcc_h19
				var _797_cf17 _dafny.MultiSet = _796___mcc_h19
				_ = _797_cf17
				var _798_v51 _dafny.MultiSet
				_ = _798_v51
				_798_v51 = _dafny.MultiSetOf(_746_cf31)
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_710_v7), 0))
				_ = _index123
				(_710_v7).ArraySet1(((_798_v51).Cardinality()).Plus(_dafny.IntOfUint32((_717_v12).Cardinality())), (_index123).Int())
				var _799_v52 D5
				_ = _799_v52
				_799_v52 = Companion_D5_.Create_DC13_(Companion_Default___.Fm43((_dafny.Zero).Minus(_746_cf31), globalState))
				var _800_v53 _dafny.Map
				_ = _800_v53
				_800_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool), _717_v12)
				var _801_v54 _dafny.Sequence
				_ = _801_v54
				_801_v54 = _dafny.SeqOf(_717_v12, _717_v12, _717_v12, _717_v12, _717_v12)
				var _802_v55 _dafny.Sequence
				_ = _802_v55
				_802_v55 = _dafny.SeqOf(_dafny.SeqOf(_717_v12, _717_v12, (func() _dafny.Sequence {
					if (_800_v53).Contains(!((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool))) {
						return (_800_v53).Get(!((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool))).(_dafny.Sequence)
					}
					return _717_v12
				})(), _717_v12), _dafny.Companion_Sequence_.Update(_801_v54, (Companion_Default___.SafeIndex(_704_v1, _dafny.IntOfUint32((_801_v54).Cardinality()))).Uint32(), _717_v12), _801_v54)
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_710_v7), 0))
				_ = _index124
				var _rhs173 _dafny.Int = (_747_cf30).Times(_704_v1)
				_ = _rhs173
				var _rhs174 _dafny.Int = _dafny.IntOfUint32(((_802_v55).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).Fm2(_746_cf31, globalState)), _dafny.IntOfUint32((_802_v55).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
				_ = _rhs174
				var _rhs175 D5 = _799_v52
				_ = _rhs175
				var _rhs176 bool = true
				_ = _rhs176
				var _rhs177 _dafny.MultiSet = _716_v11
				_ = _rhs177
				var _lhs121 _dafny.Array = _710_v7
				_ = _lhs121
				var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_710_v7), 0))
				_ = _lhs122
				(_lhs121).ArraySet1(_rhs173, (_lhs122).Int())
				_704_v1 = _rhs174
				_799_v52 = _rhs175
				_744_cf33 = _rhs176
				_716_v11 = _rhs177
				var _803_v56 _dafny.Array
				_ = _803_v56
				var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(6))
				_ = _nw114
				_803_v56 = _nw114
				var _804_v57 _dafny.Set
				_ = _804_v57
				_804_v57 = _dafny.SetOf(_this.F25, _745_cf32, (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool))
				var _805_v58 D12
				_ = _805_v58
				_805_v58 = Companion_D12_.Create_DC32_(_804_v57)
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_803_v56), 0))
				_ = _index125
				(_803_v56).ArraySet1((_dafny.MultiSetOf(_804_v57, (_805_v58).Dtor_cf49(), _804_v57, _804_v57, _804_v57)).Update(_804_v57, Companion_Default___.Abs(_746_cf31)), (_index125).Int())
				var _806_v59 _dafny.MultiSet
				_ = _806_v59
				_806_v59 = _dafny.MultiSetOf(_804_v57)
				var _807_v60 _dafny.Map
				_ = _807_v60
				_807_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_cf32, _747_cf30)
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_803_v56), 0))
				_ = _index126
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
				_ = _index127
				var _rhs178 _dafny.MultiSet = (_806_v59).Union((_806_v59).Intersection(_806_v59))
				_ = _rhs178
				var _rhs179 bool = Companion_Default___.Fm26(globalState)
				_ = _rhs179
				var _rhs180 _dafny.Int = (_713_v10).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_807_v60).Contains(!(!(_743_cf34))) {
						return (_807_v60).Get(!(!(_743_cf34))).(_dafny.Int)
					}
					return _747_cf30
				})(), _dafny.IntOfUint32((_713_v10).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs180
				var _lhs123 _dafny.Array = _803_v56
				_ = _lhs123
				var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_803_v56), 0))
				_ = _lhs124
				var _lhs125 _dafny.Array = _702_v0
				_ = _lhs125
				var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
				_ = _lhs126
				var _lhs127 *GlobalState = globalState
				_ = _lhs127
				(_lhs123).ArraySet1(_rhs178, (_lhs124).Int())
				(_lhs125).ArraySet1(_rhs179, (_lhs126).Int())
				_lhs127.F7 = _rhs180
				(globalState).F15 = _746_cf31
				var _808_v61 _dafny.Map
				_ = _808_v61
				_808_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v1, _744_cf33)
				var _809_v62 _dafny.Map
				_ = _809_v62
				_809_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm35(globalState), _746_cf31)
				var _810_v63 _dafny.Array
				_ = _810_v63
				var _nwElement0_25 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_713_v10).Cardinality()), _704_v1)
				_ = _nwElement0_25
				var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(10))
				_ = _nw115
				(_nw115).ArraySet1(_nwElement0_25, 0)
				(_nw115).ArraySet1(Companion_Default___.SafeModuloInt((_808_v61).Cardinality(), (_809_v62).Cardinality()), 1)
				(_nw115).ArraySet1(_747_cf30, 2)
				(_nw115).ArraySet1((_dafny.Zero).Minus((_710_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_710_v7), 0))).Int()).(_dafny.Int)), 3)
				(_nw115).ArraySet1(Companion_Default___.SafeDivisionInt((_808_v61).Cardinality(), _747_cf30), 4)
				(_nw115).ArraySet1(_747_cf30, 5)
				(_nw115).ArraySet1((_710_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_710_v7), 0))).Int()).(_dafny.Int), 6)
				(_nw115).ArraySet1((_710_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_710_v7), 0))).Int()).(_dafny.Int), 7)
				(_nw115).ArraySet1(_746_cf31, 8)
				(_nw115).ArraySet1((_this).Fm2((_710_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_710_v7), 0))).Int()).(_dafny.Int), globalState), 9)
				_810_v63 = _nw115
				r1 = _810_v63
			} else {
				var _811___mcc_h20 D5 = _source16.Get_().(D5_DC15).Cf22
				_ = _811___mcc_h20
				var _812_cf22 D5 = _811___mcc_h20
				_ = _812_cf22
				var _813_v64 _dafny.Map
				_ = _813_v64
				_813_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_717_v12).Cardinality()))
				var _814_v65 _dafny.Map
				_ = _814_v65
				_814_v65 = _813_v64
				var _815_v66 _dafny.MultiSet
				_ = _815_v66
				_815_v66 = _dafny.MultiSetOf(_813_v64, (_814_v65), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_708_v5).Select((Companion_Default___.SafeIndex(_747_cf30, _dafny.IntOfUint32((_708_v5).Cardinality()))).Uint32()).(bool), _704_v1)).Merge(Companion_Default___.Fm44(_747_cf30, true, globalState)), (_813_v64).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm26(globalState), _747_cf30)))
				(globalState).F4 = (func() _dafny.Int {
					if (_815_v66).Contains(_813_v64) {
						return (_815_v66).Multiplicity(_813_v64)
					}
					return Companion_Default___.SafeModuloInt(_704_v1, _704_v1)
				})()
				(globalState).F7 = _704_v1
				_710_v7 = (Companion_D2_.Create_DC7_((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool), _716_v11, _710_v7)).Dtor_cf10()
				_743_cf34 = _dafny.Companion_Sequence_.Contains(_713_v10, _746_cf31)
			}
			_716_v11 = (_716_v11).Update(true, Companion_Default___.Abs(_dafny.IntOfInt64(643)))
		} else if _source14.Is_DC22() {
			var _816___mcc_h6 bool = _source14.Get_().(D7_DC22).Cf35
			_ = _816___mcc_h6
			var _817___mcc_h7 _dafny.Int = _source14.Get_().(D7_DC22).Cf36
			_ = _817___mcc_h7
			var _818___mcc_h8 bool = _source14.Get_().(D7_DC22).Cf37
			_ = _818___mcc_h8
			var _819___mcc_h9 _dafny.Int = _source14.Get_().(D7_DC22).Cf38
			_ = _819___mcc_h9
			var _820_cf38 _dafny.Int = _819___mcc_h9
			_ = _820_cf38
			var _821_cf37 bool = _818___mcc_h8
			_ = _821_cf37
			var _822_cf36 _dafny.Int = _817___mcc_h7
			_ = _822_cf36
			var _823_cf35 bool = _816___mcc_h6
			_ = _823_cf35
			(_this).F25 = (_716_v11).IsSubsetOf(_716_v11)
			var _824_v67 _dafny.Set
			_ = _824_v67
			_824_v67 = _dafny.SetOf((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool))
			var _825_v68 _dafny.Map
			_ = _825_v68
			_825_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_824_v67).Cardinality()), _822_cf36)
			var _826_v69 D3
			_ = _826_v69
			_826_v69 = Companion_D3_.Create_DC10_(_704_v1, (_825_v68).Cardinality())
			var _827_v70 _dafny.Map
			_ = _827_v70
			_827_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_717_v12, _dafny.UnicodeSeqOfUtf8Bytes("de"))
			var _pat_let_tv10 = _704_v1
			_ = _pat_let_tv10
			var _828_v71 _dafny.Sequence
			_ = _828_v71
			_828_v71 = _dafny.SeqOf(_826_v69, func(_pat_let6_0 D3) D3 {
				return func(_829_dt__update__tmp_h2 D3) D3 {
					return func(_pat_let7_0 _dafny.Int) D3 {
						return func(_830_dt__update_hcf13_h0 _dafny.Int) D3 {
							return Companion_D3_.Create_DC10_(_830_dt__update_hcf13_h0, (_829_dt__update__tmp_h2).Dtor_cf14())
						}(_pat_let7_0)
					}(_pat_let_tv10)
				}(_pat_let6_0)
			}(_826_v69), _826_v69, _826_v69, Companion_D3_.Create_DC10_((_827_v70).Cardinality(), _dafny.IntOfInt64(-373)))
			var _831_v72 _dafny.Map
			_ = _831_v72
			_831_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool), _828_v71)
			var _832_v73 _dafny.Map
			_ = _832_v73
			_832_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_823_cf35, _this.F25)
			_831_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_832_v73).Contains((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool)) {
					return (_832_v73).Get((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool)).(bool)
				}
				return _823_cf35
			})(), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_826_v69, _826_v69), (Companion_Default___.SafeIndex(_822_cf36, _dafny.IntOfUint32((_dafny.SeqOf(_826_v69, _826_v69)).Cardinality()))).Uint32(), _826_v69))
			_821_cf37 = (_705_v2).IsSubsetOf(_705_v2)
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _index128
			(_702_v0).ArraySet1(Companion_Default___.Fm26(globalState), (_index128).Int())
		} else {
			var _833___mcc_h10 _dafny.Sequence = _source14.Get_().(D7_DC19).Cf28
			_ = _833___mcc_h10
			var _834_cf28 _dafny.Sequence = _833___mcc_h10
			_ = _834_cf28
			var _835_v74 _dafny.Map
			_ = _835_v74
			_835_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(655), _dafny.CodePoint('g'))
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _index129
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _index130
			var _rhs181 bool = _this.F25
			_ = _rhs181
			var _rhs182 _dafny.Map = _835_v74
			_ = _rhs182
			var _rhs183 bool = (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool)
			_ = _rhs183
			var _rhs184 _dafny.Int = (_704_v1).Minus((_dafny.Zero).Minus(_704_v1))
			_ = _rhs184
			var _rhs185 _dafny.Array = _710_v7
			_ = _rhs185
			var _lhs128 _dafny.Array = _702_v0
			_ = _lhs128
			var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _lhs129
			var _lhs130 _dafny.Array = _702_v0
			_ = _lhs130
			var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _lhs131
			(_lhs128).ArraySet1(_rhs181, (_lhs129).Int())
			_835_v74 = _rhs182
			(_lhs130).ArraySet1(_rhs183, (_lhs131).Int())
			_704_v1 = _rhs184
			_710_v7 = _rhs185
			_708_v5 = _708_v5
			(_this).F25 = (func() bool {
				if true {
					return _this.F25
				}
				return (_this.F25) == (_this.F25)
			})()
			var _836_v75 _dafny.Array
			_ = _836_v75
			var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
			_ = _nw116
			_836_v75 = _nw116
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_836_v75), 0))
			_ = _index131
			(_836_v75).ArraySet1((_835_v74).Merge(_835_v74), (_index131).Int())
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_836_v75), 0))
			_ = _index132
			(_836_v75).ArraySet1(_835_v74, (_index132).Int())
		}
		var _837_v76 _dafny.Set
		_ = _837_v76
		_837_v76 = _dafny.SetOf(_710_v7)
		var _838_v77 _dafny.Map
		_ = _838_v77
		_838_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_717_v12, _704_v1)
		var _hi5 _dafny.Int = (func() _dafny.Int {
			if (_838_v77).Contains(_717_v12) {
				return (_838_v77).Get(_717_v12).(_dafny.Int)
			}
			return _704_v1
		})()
		_ = _hi5
		for _839_i6 := (_837_v76).Cardinality(); _839_i6.Cmp(_hi5) < 0; _839_i6 = _839_i6.Plus(_dafny.One) {
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _index133
			(_702_v0).ArraySet1((_718_v13).Dtor_cf34(), (_index133).Int())
			(_this).F25 = ((_839_i6).Plus(_dafny.IntOfInt64(211))).Cmp(_839_i6) != 0
			var _840_v78 *C2
			_ = _840_v78
			var _nw117 *C2 = New_C2_()
			_ = _nw117
			_nw117.Ctor__()
			_840_v78 = _nw117
			var _841_v79 _dafny.Map
			_ = _841_v79
			_841_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25, _840_v78)
			var _842_v80 _dafny.Map
			_ = _842_v80
			_842_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25, _this.F25)
			var _843_v81 _dafny.Map
			_ = _843_v81
			_843_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25, _842_v80)
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _index134
			var _rhs186 bool = (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool)
			_ = _rhs186
			var _rhs187 bool = true
			_ = _rhs187
			var _rhs188 *C2 = (func() *C2 {
				if (_841_v79).Contains((_840_v78).Fm1(_843_v81, _839_i6, globalState)) {
					return (_841_v79).Get((_840_v78).Fm1(_843_v81, _839_i6, globalState)).(*C2)
				}
				return _840_v78
			})()
			_ = _rhs188
			var _lhs132 _dafny.Array = _702_v0
			_ = _lhs132
			var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _lhs133
			var _lhs134 *C3 = _this
			_ = _lhs134
			(_lhs132).ArraySet1(_rhs186, (_lhs133).Int())
			_lhs134.F25 = _rhs187
			_840_v78 = _rhs188
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))
			_ = _index135
			(_702_v0).ArraySet1((_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool), (_index135).Int())
		}
		var _844_v82 *C1
		_ = _844_v82
		var _nw118 *C1 = New_C1_()
		_ = _nw118
		_nw118.Ctor__()
		_844_v82 = _nw118
		r0 = _dafny.SetOf(_704_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pvydf")).Cardinality()))
		var _845_v83 _dafny.MultiSet
		_ = _845_v83
		_845_v83 = _dafny.MultiSetOf(_704_v1)
		r1 = (func() _dafny.Array {
			if ((Companion_Default___.Fm42(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-269))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}(func(_846_i7 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(23)
			}))).Cardinality()), _704_v1, globalState)).Dtor_cf5()).IsDisjointFrom(_845_v83) {
				return _710_v7
			}
			return _710_v7
		})()
		var _847_v84 _dafny.Map
		_ = _847_v84
		_847_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v1, (_702_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_702_v0), 0))).Int()).(bool))
		var _848_v86 _dafny.Sequence
		_ = _848_v86
		_848_v86 = _dafny.SeqOf(_847_v84, func() _dafny.Map {
			var _coll35 = _dafny.NewMapBuilder()
			_ = _coll35
			for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(11), _dafny.IntOfInt64(121))); ; {
				_compr_35, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _849_v85 _dafny.Int
				_849_v85 = interface{}(_compr_35).(_dafny.Int)
				if ((_dafny.IntOfInt64(11)).Cmp(_849_v85) <= 0) && ((_849_v85).Cmp(_dafny.IntOfInt64(121)) < 0) {
					_coll35.Add((_849_v85).Minus(_704_v1), !(true))
				}
			}
			return _coll35.ToMap()
		}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_v1, _this.F25), _847_v84, _847_v84)
		var _850_v87 D14
		_ = _850_v87
		_850_v87 = Companion_D14_.Create_DC35_(_848_v86)
		r2 = (_850_v87).Dtor_cf52()
		return r0, r1, r2
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	dummy byte
}

func New_C4_() *C4 {
	_this := C4{}

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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__() {
	{
	}
}
func (_this *C4) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return ((Companion_D5_.Create_DC13_(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(-472), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('h'))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))))).Dtor_cf17()).Intersection((_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(-140)), func() _dafny.Set {
			var _coll36 = _dafny.NewBuilder()
			_ = _coll36
			for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(470), _dafny.IntOfInt64(446))); ; {
				_compr_36, _ok37 := _iter37()
				if !_ok37 {
					break
				}
				var _851_v0 _dafny.Int
				_851_v0 = interface{}(_compr_36).(_dafny.Int)
				if ((_dafny.IntOfInt64(470)).Cmp(_851_v0) <= 0) && ((_851_v0).Cmp(_dafny.IntOfInt64(446)) < 0) {
					_coll36.Add((_851_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(false, true, false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}(func(_852_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('s')
					}))).Cardinality()))).Cardinality())))
				}
			}
			return _coll36.ToSet()
		}())).Union(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()), _dafny.IntOfInt64(710)))))
	}
}
func (_this *C4) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.SetOf(_dafny.IntOfInt64(604), (_dafny.Zero).Minus(_dafny.IntOfInt64(-912)), _dafny.IntOfInt64(783))).IsProperSubsetOf(_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(241), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), false)).Cardinality(), _dafny.IntOfInt64(-235))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ecngi")).Cardinality())))
	}
}
func (_this *C4) Fm23(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if !(true) {
				return _dafny.IntOfInt64(909)
			}
			return _dafny.IntOfInt64(731)
		})(), (func() _dafny.Int {
			if true {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xueolns")).Cardinality())
			}
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("apewrwfc")).Cardinality())
		})())
	}
}
func (_this *C4) M6(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 D2, globalState *GlobalState) {
	{
		var _853_v0 _dafny.Array
		_ = _853_v0
		var _nwElement0_26 bool = p2
		_ = _nwElement0_26
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(5))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_26, 0)
		(_nw119).ArraySet1(p2, 1)
		(_nw119).ArraySet1(p2, 2)
		(_nw119).ArraySet1(p2, 3)
		(_nw119).ArraySet1(p2, 4)
		_853_v0 = _nw119
		var _854_v1 D0
		_ = _854_v1
		_854_v1 = Companion_D0_.Create_DC1_(_853_v0)
		var _source17 D0 = _854_v1
		_ = _source17
		if _source17.Is_DC1() {
			var _855___mcc_h0 _dafny.Array = _source17.Get_().(D0_DC1).Cf1
			_ = _855___mcc_h0
			var _856_cf1 _dafny.Array = _855___mcc_h0
			_ = _856_cf1
			_856_cf1 = _856_cf1
			var _857_v2 _dafny.Set
			_ = _857_v2
			_857_v2 = _dafny.SetOf(p0)
			var _858_v3 _dafny.Sequence
			_ = _858_v3
			_858_v3 = _dafny.SeqOf(_857_v2)
			var _859_v4 _dafny.Map
			_ = _859_v4
			_859_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _858_v3)
			var _860_v5 _dafny.Sequence
			_ = _860_v5
			_860_v5 = _dafny.SeqOf(p2)
			var _861_v6 _dafny.Sequence
			_ = _861_v6
			_861_v6 = _dafny.UnicodeSeqOfUtf8Bytes("wnhu")
			_859_v4 = (_859_v4).Update((_860_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_861_v6).Cardinality()), _dafny.IntOfUint32((_860_v5).Cardinality()))).Uint32()).(bool), (func() _dafny.Sequence {
				if p2 {
					return _858_v3
				}
				return _858_v3
			})())
			(globalState).F15 = ((p0).Plus(_dafny.IntOfUint32((_861_v6).Cardinality()))).Minus(_dafny.IntOfUint32((_861_v6).Cardinality()))
			var _862_v7 bool
			_ = _862_v7
			_862_v7 = true
			_862_v7 = p2
		} else {
			var _863___mcc_h1 _dafny.Array = _source17.Get_().(D0_DC0).Cf0
			_ = _863___mcc_h1
			var _864_cf0 _dafny.Array = _863___mcc_h1
			_ = _864_cf0
			var _865_v8 _dafny.Array
			_ = _865_v8
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_24
			var _nw120 _dafny.Array
			_ = _nw120
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw120 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Set = (func(_866_p2 bool) func(_dafny.Int) _dafny.Set {
					return func(_867_i0 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_866_p2)).Union(_dafny.SetOf(_866_p2))
					}
				})(p2)
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw120 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw120).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw120).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_865_v8 = _nw120
			var _868_v9 _dafny.Set
			_ = _868_v9
			_868_v9 = _dafny.SetOf(p2, p2)
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_865_v8), 0))
			_ = _index136
			(_865_v8).ArraySet1(_868_v9, (_index136).Int())
			var _869_v10 _dafny.Map
			_ = _869_v10
			_869_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_865_v8), 0))
			_ = _index137
			(_865_v8).ArraySet1(_dafny.SetOf((func() bool {
				if (_869_v10).Contains(p2) {
					return (_869_v10).Get(p2).(bool)
				}
				return p2
			})()), (_index137).Int())
			var _870_v11 _dafny.Map
			_ = _870_v11
			_870_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
			var _871_v12 bool
			_ = _871_v12
			_871_v12 = true
			var _872_v13 _dafny.Sequence
			_ = _872_v13
			_872_v13 = _dafny.UnicodeSeqOfUtf8Bytes("wihxevenh")
			var _873_v14 D1
			_ = _873_v14
			_873_v14 = Companion_D1_.Create_DC3_(_872_v13, p2)
			var _874_v15 T0
			_ = _874_v15
			var _nw121 *C2 = New_C2_()
			_ = _nw121
			_nw121.Ctor__()
			_874_v15 = _nw121
			var _875_v16 _dafny.Set
			_ = _875_v16
			_875_v16 = _dafny.SetOf(_874_v15)
			var _rhs189 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_873_v14).Dtor_cf4())
			_ = _rhs189
			var _rhs190 bool = false
			_ = _rhs190
			var _rhs191 _dafny.Int = (func() _dafny.Int {
				if !((_this).Fm8((_875_v16).Cardinality(), false, p0, globalState)) || (_871_v12) {
					return (_869_v10).Cardinality()
				}
				return p0
			})()
			_ = _rhs191
			var _lhs135 *GlobalState = globalState
			_ = _lhs135
			_870_v11 = _rhs189
			_871_v12 = _rhs190
			_lhs135.F4 = _rhs191
			var _876_v17 _dafny.CodePoint
			_ = _876_v17
			_876_v17 = _dafny.CodePoint('g')
			var _877_v18 _dafny.Set
			_ = _877_v18
			_877_v18 = _dafny.SetOf(_dafny.CodePoint('s'), _876_v17)
			var _878_v19 _dafny.Map
			_ = _878_v19
			_878_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_877_v18).Cardinality())
			_878_v19 = _878_v19
			_869_v10 = ((_869_v10).Merge(_869_v10)).Merge(_869_v10)
		}
		var _879_v20 _dafny.Map
		_ = _879_v20
		_879_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		_879_v20 = _879_v20
		var _880_v21 _dafny.Set
		_ = _880_v21
		_880_v21 = _dafny.SetOf(Companion_Default___.Fm27(globalState))
		if (_880_v21).IsSubsetOf(_880_v21) {
			var _881_v22 _dafny.Array
			_ = _881_v22
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(19))
			_ = _nw122
			_881_v22 = _nw122
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_881_v22), 0))
			_ = _index138
			(_881_v22).ArraySet1CodePoint(Companion_Default___.Fm38(p2, globalState), (_index138).Int())
			var _882_v23 _dafny.Sequence
			_ = _882_v23
			_882_v23 = _dafny.UnicodeSeqOfUtf8Bytes("dnltn")
			var _883_v24 _dafny.Sequence
			_ = _883_v24
			_883_v24 = _dafny.SeqOf(p2)
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_853_v0), 0))
			_ = _index139
			(_853_v0).ArraySet1((_883_v24).Select((Companion_Default___.SafeIndex((p1).Cardinality(), _dafny.IntOfUint32((_883_v24).Cardinality()))).Uint32()).(bool), (_index139).Int())
			var _884_v25 bool
			_ = _884_v25
			_884_v25 = true
			var _885_v26 _dafny.CodePoint
			_ = _885_v26
			_885_v26 = _dafny.CodePoint('f')
			var _886_v27 _dafny.Sequence
			_ = _886_v27
			_886_v27 = _dafny.SeqOf(p0)
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_881_v22), 0))
			_ = _index140
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_853_v0), 0))
			_ = _index141
			var _rhs192 _dafny.CodePoint = _885_v26
			_ = _rhs192
			var _rhs193 _dafny.Sequence = _882_v23
			_ = _rhs193
			var _rhs194 bool = !(!((_this).Fm8((_886_v27).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_886_v27).Cardinality()))).Uint32()).(_dafny.Int), (func() bool {
				if p2 {
					return _884_v25
				}
				return _884_v25
			})(), p0, globalState)))
			_ = _rhs194
			var _rhs195 bool = _884_v25
			_ = _rhs195
			var _lhs136 _dafny.Array = _881_v22
			_ = _lhs136
			var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_881_v22), 0))
			_ = _lhs137
			var _lhs138 _dafny.Array = _853_v0
			_ = _lhs138
			var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_853_v0), 0))
			_ = _lhs139
			(_lhs136).ArraySet1CodePoint(_rhs192, (_lhs137).Int())
			_882_v23 = _rhs193
			(_lhs138).ArraySet1(_rhs194, (_lhs139).Int())
			_884_v25 = _rhs195
			var _887_v28 _dafny.Array
			_ = _887_v28
			var _nw123 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(5))
			_ = _nw123
			_887_v28 = _nw123
			var _888_v29 D1
			_ = _888_v29
			_888_v29 = Companion_D1_.Create_DC2_(p0)
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_887_v28), 0))
			_ = _index142
			(_887_v28).ArraySet1(_888_v29, (_index142).Int())
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_887_v28), 0))
			_ = _index143
			(_887_v28).ArraySet1(_888_v29, (_index143).Int())
			if p2 {
				_882_v23 = _882_v23
				var _889_v30 _dafny.Set
				_ = _889_v30
				_889_v30 = _dafny.SetOf((Companion_D11_.Create_DC30_(p2, p0)).Dtor_cf47())
				var _890_v31 _dafny.Sequence
				_ = _890_v31
				_890_v31 = _dafny.SeqOf(_dafny.SetOf(p0, (_dafny.MultiSetOf(_dafny.IntOfInt64(494), p0)).Cardinality(), p0), _889_v30, _889_v30, _889_v30)
				var _891_v32 _dafny.Set
				_ = _891_v32
				_891_v32 = _dafny.SetOf(Companion_Default___.Fm26(globalState), !(p2))
				var _rhs196 bool = false
				_ = _rhs196
				var _rhs197 _dafny.Sequence = Companion_Default___.Fm30(p0, p2, (_890_v31).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_890_v31).Cardinality()))).Uint32()).(_dafny.Set), !(!(!((_891_v32).IsProperSubsetOf(_891_v32)))), globalState)
				_ = _rhs197
				_884_v25 = _rhs196
				_882_v23 = _rhs197
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_853_v0), 0))
				_ = _index144
				(_853_v0).ArraySet1(!(p2), (_index144).Int())
				var _892_v33 _dafny.Sequence
				_ = _892_v33
				_892_v33 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("rfny"))
				var _893_v34 D4
				_ = _893_v34
				_893_v34 = Companion_D4_.Create_DC11_(_892_v33)
				var _pat_let_tv11 = _892_v33
				_ = _pat_let_tv11
				var _894_v35 _dafny.Map
				_ = _894_v35
				_894_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_853_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_853_v0), 0))).Int()).(bool), func(_pat_let8_0 D4) D4 {
					return func(_895_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let9_0 _dafny.Sequence) D4 {
							return func(_896_dt__update_hcf15_h0 _dafny.Sequence) D4 {
								return Companion_D4_.Create_DC11_(_896_dt__update_hcf15_h0)
							}(_pat_let9_0)
						}(_pat_let_tv11)
					}(_pat_let8_0)
				}(_893_v34))
				var _897_v36 T1
				_ = _897_v36
				var _nw124 *C2 = New_C2_()
				_ = _nw124
				_nw124.Ctor__()
				_897_v36 = _nw124
				var _898_v37 _dafny.Sequence
				_ = _898_v37
				_898_v37 = _dafny.SeqOf(_897_v36, _897_v36)
				(globalState).F13 = (_this).Fm23(_894_v35, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_892_v33).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_892_v33).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_898_v37).Cardinality())), p0, globalState)
				_884_v25 = (((p1).Cardinality()).Minus(p0)).Cmp(p0) > 0
			} else {
				var _899_v38 _dafny.Sequence
				_ = _899_v38
				_899_v38 = _dafny.SeqOf(_882_v23)
				var _900_v39 _dafny.Map
				_ = _900_v39
				_900_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, Companion_D4_.Create_DC11_(_899_v38))
				var _901_v40 _dafny.Set
				_ = _901_v40
				_901_v40 = _dafny.SetOf(p0, (_this).Fm23(_900_v39, _dafny.IntOfUint32((_882_v23).Cardinality()), _dafny.IntOfInt64(-193), globalState), _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), p0, p0)
				_899_v38 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(950))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_902_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_903_i1 _dafny.Int) _dafny.Sequence {
						return _902_v23
					}
				})(_882_v23))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(950))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_904_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_905_i1 _dafny.Int) _dafny.Sequence {
						return _904_v23
					}
				})(_882_v23)))).Cardinality()))).Uint32(), Companion_Default___.Fm30(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(278))).Cardinality()), !(p2), _901_v40, _884_v25, globalState)), _899_v38)
				(globalState).F4 = _dafny.IntOfInt64(310)
				var _906_v41 *C2
				_ = _906_v41
				var _nw125 *C2 = New_C2_()
				_ = _nw125
				_nw125.Ctor__()
				_906_v41 = _nw125
				_882_v23 = _dafny.UnicodeSeqOfUtf8Bytes("bsu")
				var _907_v42 _dafny.Array
				_ = _907_v42
				var _nwElement0_27 bool = (_853_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_853_v0), 0))).Int()).(bool)
				_ = _nwElement0_27
				var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(5))
				_ = _nw126
				(_nw126).ArraySet1(_nwElement0_27, 0)
				(_nw126).ArraySet1(p2, 1)
				(_nw126).ArraySet1(_884_v25, 2)
				(_nw126).ArraySet1(p2, 3)
				(_nw126).ArraySet1(!(!(p2) || (_884_v25)), 4)
				_907_v42 = _nw126
				var _rhs198 _dafny.Array = (func() _dafny.Array {
					if _884_v25 {
						return _907_v42
					}
					return _907_v42
				})()
				_ = _rhs198
				var _rhs199 _dafny.Int = _dafny.IntOfUint32((_883_v24).Cardinality())
				_ = _rhs199
				var _lhs140 *GlobalState = globalState
				_ = _lhs140
				_907_v42 = _rhs198
				_lhs140.F15 = _rhs199
			}
			(globalState).F13 = p0
			var _908_v43 _dafny.Array
			_ = _908_v43
			var _nw127 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(27))
			_ = _nw127
			_908_v43 = _nw127
			var _909_v44 D4
			_ = _909_v44
			_909_v44 = Companion_D4_.Create_DC12_(p0)
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_908_v43), 0))
			_ = _index145
			(_908_v43).ArraySet1(_909_v44, (_index145).Int())
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_908_v43), 0))
			_ = _index146
			(_908_v43).ArraySet1(_909_v44, (_index146).Int())
		} else {
			var _910_v45 _dafny.CodePoint
			_ = _910_v45
			_910_v45 = _dafny.CodePoint('y')
			var _911_v46 _dafny.Map
			_ = _911_v46
			_911_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v45, false)
			var _912_v47 _dafny.Sequence
			_ = _912_v47
			_912_v47 = _dafny.UnicodeSeqOfUtf8Bytes("qbe")
			_911_v46 = (_911_v46).Update(_910_v45, (_dafny.IntOfUint32((_912_v47).Cardinality())).Cmp(p0) >= 0)
			(globalState).F13 = (_dafny.IntOfInt64(-853)).Times(p0)
			var _913_v48 _dafny.Map
			_ = _913_v48
			_913_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(p0)).Cmp((_dafny.Zero).Minus(p0)) > 0, p2)
			_913_v48 = (_913_v48).Update(p2, p2)
			var _914_v49 _dafny.Sequence
			_ = _914_v49
			_914_v49 = _dafny.SeqOf((func() bool {
				if p2 {
					return true
				}
				return p2
			})())
			var _915_v50 _dafny.Set
			_ = _915_v50
			_915_v50 = _dafny.SetOf(p0)
			var _916_v51 _dafny.Array
			_ = _916_v51
			var _nwElement0_28 _dafny.Int = _dafny.IntOfInt64(784)
			_ = _nwElement0_28
			var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(3))
			_ = _nw128
			(_nw128).ArraySet1(_nwElement0_28, 0)
			(_nw128).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, (_915_v50).Cardinality())), 1)
			(_nw128).ArraySet1(p0, 2)
			_916_v51 = _nw128
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_916_v51), 0))
			_ = _index147
			(_916_v51).ArraySet1((_dafny.IntOfInt64(-129)).Plus(p0), (_index147).Int())
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_916_v51), 0))
			_ = _index148
			var _rhs200 _dafny.Sequence = _914_v49
			_ = _rhs200
			var _rhs201 _dafny.Int = _dafny.IntOfInt64(886)
			_ = _rhs201
			var _lhs141 _dafny.Array = _916_v51
			_ = _lhs141
			var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_916_v51), 0))
			_ = _lhs142
			_914_v49 = _rhs200
			(_lhs141).ArraySet1(_rhs201, (_lhs142).Int())
			var _917_v52 bool
			_ = _917_v52
			_917_v52 = false
			_917_v52 = !(((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Map {
				var _coll37 = _dafny.NewMapBuilder()
				_ = _coll37
				for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(745), _dafny.IntOfInt64(-795))); ; {
					_compr_37, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _918_v53 _dafny.Int
					_918_v53 = interface{}(_compr_37).(_dafny.Int)
					if ((_dafny.IntOfInt64(745)).Cmp(_918_v53) <= 0) && ((_918_v53).Cmp(_dafny.IntOfInt64(-795)) < 0) {
						_coll37.Add(Companion_Default___.SafeModuloInt(_918_v53, (_916_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_916_v51), 0))).Int()).(_dafny.Int)), _912_v47)
					}
				}
				return _coll37.ToMap()
			}()).Cardinality()))).Cmp((_916_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_916_v51), 0))).Int()).(_dafny.Int)) >= 0)
		}
		var _919_v54 _dafny.Map
		_ = _919_v54
		_919_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(-152))
		var _920_v55 D4
		_ = _920_v55
		_920_v55 = Companion_D4_.Create_DC12_((_919_v54).Cardinality())
		var _921_v56 _dafny.Sequence
		_ = _921_v56
		_921_v56 = _dafny.UnicodeSeqOfUtf8Bytes("jpfda")
		var _922_v57 _dafny.Set
		_ = _922_v57
		_922_v57 = _dafny.SetOf(_dafny.IntOfUint32((_921_v56).Cardinality()))
		var _923_v58 _dafny.MultiSet
		_ = _923_v58
		_923_v58 = _dafny.MultiSetOf(_922_v57)
		var _924_v59 _dafny.Set
		_ = _924_v59
		_924_v59 = _dafny.SetOf(p2, p2)
		var _925_v60 D7
		_ = _925_v60
		_925_v60 = Companion_D7_.Create_DC22_(true, p0, p2, p0)
		var _926_v61 _dafny.MultiSet
		_ = _926_v61
		_926_v61 = _dafny.MultiSetOf(_dafny.SeqOf(p0))
		var _927_v62 _dafny.Array
		_ = _927_v62
		var _nwElement0_29 _dafny.Int = p0
		_ = _nwElement0_29
		var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(23))
		_ = _nw129
		(_nw129).ArraySet1(_nwElement0_29, 0)
		(_nw129).ArraySet1(_dafny.IntOfInt64(805), 1)
		(_nw129).ArraySet1(p0, 2)
		(_nw129).ArraySet1((_dafny.Zero).Minus((_920_v55).Dtor_cf16()), 3)
		(_nw129).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Cardinality(), p2)).Cardinality(), 4)
		(_nw129).ArraySet1(p0, 5)
		(_nw129).ArraySet1(p0, 6)
		(_nw129).ArraySet1((_923_v58).Cardinality(), 7)
		(_nw129).ArraySet1(p0, 8)
		(_nw129).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(658))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}((func(_928_p0 _dafny.Int, _929_v20 _dafny.Map, _930_p2 bool) func(_dafny.Int) _dafny.Sequence {
			return func(_931_i3 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf((func() bool {
					if (_929_v20).Contains(_928_p0) {
						return (_929_v20).Get(_928_p0).(bool)
					}
					return _930_p2
				})(), (Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-685))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}(func(_932_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})), !(_930_p2))).Dtor_cf4(), _930_p2)
			}
		})(p0, _879_v20, p2))), (Companion_Default___.SafeIndex((_924_v59).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(658))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}((func(_933_p0 _dafny.Int, _934_v20 _dafny.Map, _935_p2 bool) func(_dafny.Int) _dafny.Sequence {
			return func(_936_i3 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf((func() bool {
					if (_934_v20).Contains(_933_p0) {
						return (_934_v20).Get(_933_p0).(bool)
					}
					return _935_p2
				})(), (Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-685))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}(func(_937_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})), !(_935_p2))).Dtor_cf4(), _935_p2)
			}
		})(p0, _879_v20, p2)))).Cardinality()))).Uint32(), _dafny.SeqOf((_925_v60).Dtor_cf37(), p2))).Cardinality()), 9)
		(_nw129).ArraySet1(p0, 10)
		(_nw129).ArraySet1(p0, 11)
		(_nw129).ArraySet1(p0, 12)
		(_nw129).ArraySet1(p0, 13)
		(_nw129).ArraySet1((_dafny.MultiSetOf(p2)).Cardinality(), 14)
		(_nw129).ArraySet1(p0, 15)
		(_nw129).ArraySet1(p0, 16)
		(_nw129).ArraySet1(p0, 17)
		(_nw129).ArraySet1(p0, 18)
		(_nw129).ArraySet1((_926_v61).Cardinality(), 19)
		(_nw129).ArraySet1(p0, 20)
		(_nw129).ArraySet1(p0, 21)
		(_nw129).ArraySet1(p0, 22)
		_927_v62 = _nw129
		var _938_v63 _dafny.Sequence
		_ = _938_v63
		_938_v63 = _dafny.SeqOf(_927_v62, _927_v62)
		var _939_i2 _dafny.Int
		_ = _939_i2
		_939_i2 = _dafny.Zero
		{
			for !(_dafny.Companion_Sequence_.Equal(_938_v63, _938_v63)) {
				{
					if (_939_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_939_i2 = (_939_i2).Plus(_dafny.One)
					var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))
					_ = _index149
					(_927_v62).ArraySet1(p0, (_index149).Int())
					var _940_v64 D14
					_ = _940_v64
					_940_v64 = Companion_D14_.Create_DC35_(Companion_Default___.Fm45(globalState))
					var _941_v65 _dafny.Sequence
					_ = _941_v65
					_941_v65 = _dafny.SeqOf(_940_v64, _940_v64)
					var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))
					_ = _index150
					(_927_v62).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if false {
							return _941_v65
						}
						return _941_v65
					})()).Cardinality()), (_index150).Int())
					if p2 {
						(globalState).F4 = p0
						var _942_v66 bool
						_ = _942_v66
						_942_v66 = true
						_942_v66 = _942_v66
						var _943_v67 _dafny.Set
						_ = _943_v67
						_943_v67 = _dafny.SetOf(_879_v20, _879_v20)
						var _944_v68 _dafny.Sequence
						_ = _944_v68
						_944_v68 = _dafny.SeqOf(p2)
						var _rhs202 bool = !(p2)
						_ = _rhs202
						var _rhs203 bool = _942_v66
						_ = _rhs203
						var _rhs204 bool = !(!(((_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int)).Cmp((_943_v67).Cardinality()) <= 0) || (_dafny.Companion_Sequence_.Contains(_944_v68, _942_v66)))
						_ = _rhs204
						_942_v66 = _rhs202
						_942_v66 = _rhs203
						_942_v66 = _rhs204
						var _945_v69 _dafny.Sequence
						_ = _945_v69
						_945_v69 = _dafny.SeqOf((_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(431), (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int))
						(globalState).F15 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_945_v69, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_945_v69).Cardinality()))).Uint32(), (_919_v54).Cardinality()), _945_v69), _945_v69)).Cardinality())
						var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_853_v0), 0))
						_ = _index151
						(_853_v0).ArraySet1(!(p2), (_index151).Int())
						var _946_v70 _dafny.Array
						_ = _946_v70
						var _nw130 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(11))
						_ = _nw130
						_946_v70 = _nw130
						var _947_v71 D6
						_ = _947_v71
						_947_v71 = Companion_D6_.Create_DC16_(_dafny.SeqOf(_dafny.IntOfUint32((_921_v56).Cardinality()), p0))
						var _948_v72 D6
						_ = _948_v72
						_948_v72 = Companion_D6_.Create_DC18_(_947_v71)
						var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_946_v70), 0))
						_ = _index152
						(_946_v70).ArraySet1(_948_v72, (_index152).Int())
						var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_853_v0), 0))
						_ = _index153
						var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_946_v70), 0))
						_ = _index154
						var _rhs205 bool = (_944_v68).Select((Companion_Default___.SafeIndex((_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_944_v68).Cardinality()))).Uint32()).(bool)
						_ = _rhs205
						var _rhs206 D6 = _948_v72
						_ = _rhs206
						var _rhs207 bool = p2
						_ = _rhs207
						var _lhs143 _dafny.Array = _853_v0
						_ = _lhs143
						var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_853_v0), 0))
						_ = _lhs144
						var _lhs145 _dafny.Array = _946_v70
						_ = _lhs145
						var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_946_v70), 0))
						_ = _lhs146
						(_lhs143).ArraySet1(_rhs205, (_lhs144).Int())
						(_lhs145).ArraySet1(_rhs206, (_lhs146).Int())
						_942_v66 = _rhs207
					} else {
						var _949_v73 bool
						_ = _949_v73
						_949_v73 = true
						_949_v73 = p2
						var _950_v74 _dafny.Array
						_ = _950_v74
						var _len0_25 _dafny.Int = _dafny.IntOfInt64(23)
						_ = _len0_25
						var _nw131 _dafny.Array
						_ = _nw131
						if _len0_25.Cmp(_dafny.Zero) == 0 {
							_nw131 = _dafny.NewArray(_len0_25)
						} else {
							var _init25 func(_dafny.Int) _dafny.Int = (func(_951_v62 _dafny.Array) func(_dafny.Int) _dafny.Int {
								return func(_952_i5 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeDivisionInt(_952_i5, (_951_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_951_v62), 0))).Int()).(_dafny.Int))
								}
							})(_927_v62)
							_ = _init25
							var _element0_25 = _init25(_dafny.Zero)
							_ = _element0_25
							_nw131 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
							(_nw131).ArraySet1(_element0_25, 0)
							var _nativeLen0_25 = (_len0_25).Int()
							_ = _nativeLen0_25
							for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
								(_nw131).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
							}
						}
						_950_v74 = _nw131
						var _953_v75 D2
						_ = _953_v75
						_953_v75 = Companion_D2_.Create_DC5_(_950_v74)
						var _954_v76 D2
						_ = _954_v76
						_954_v76 = Companion_D2_.Create_DC8_(_953_v75)
						var _955_v77 *C3
						_ = _955_v77
						var _nw132 *C3 = New_C3_()
						_ = _nw132
						_nw132.Ctor__(_954_v76, (p0).Cmp((_924_v59).Cardinality()) < 0)
						_955_v77 = _nw132
						var _956_v78 _dafny.Array
						_ = _956_v78
						var _len0_26 _dafny.Int = _dafny.IntOfInt64(23)
						_ = _len0_26
						var _nw133 _dafny.Array
						_ = _nw133
						if _len0_26.Cmp(_dafny.Zero) == 0 {
							_nw133 = _dafny.NewArray(_len0_26)
						} else {
							var _init26 func(_dafny.Int) _dafny.Sequence = (func(_957_v56 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_958_i6 _dafny.Int) _dafny.Sequence {
									return _957_v56
								}
							})(_921_v56)
							_ = _init26
							var _element0_26 = _init26(_dafny.Zero)
							_ = _element0_26
							_nw133 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
							(_nw133).ArraySet1(_element0_26, 0)
							var _nativeLen0_26 = (_len0_26).Int()
							_ = _nativeLen0_26
							for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
								(_nw133).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
							}
						}
						_956_v78 = _nw133
						var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_956_v78), 0))
						_ = _index155
						(_956_v78).ArraySet1(_921_v56, (_index155).Int())
						var _959_v79 _dafny.MultiSet
						_ = _959_v79
						_959_v79 = _dafny.MultiSetOf((_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(729))
						var _960_v80 D8
						_ = _960_v80
						_960_v80 = Companion_D8_.Create_DC24_(p0, !(p2) || ((func() bool {
							if (_879_v20).Contains((_959_v79).Cardinality()) {
								return (_879_v20).Get((_959_v79).Cardinality()).(bool)
							}
							return _955_v77.F25
						})()))
						var _961_v81 _dafny.Map
						_ = _961_v81
						_961_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if _955_v77.F25 {
								return false
							}
							return true
						})(), _853_v0)
						var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_956_v78), 0))
						_ = _index156
						var _rhs208 _dafny.Sequence = _921_v56
						_ = _rhs208
						var _rhs209 D8 = _960_v80
						_ = _rhs209
						var _rhs210 _dafny.Map = (_961_v81).Merge((_961_v81).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _853_v0)))
						_ = _rhs210
						var _lhs147 _dafny.Array = _956_v78
						_ = _lhs147
						var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_956_v78), 0))
						_ = _lhs148
						(_lhs147).ArraySet1(_rhs208, (_lhs148).Int())
						_960_v80 = _rhs209
						_961_v81 = _rhs210
						var _962_v82 _dafny.Map
						_ = _962_v82
						_962_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _950_v74)
						_962_v82 = (_962_v82).Update(Companion_Default___.Fm34(p0, p0, _949_v73, globalState), _950_v74)
						(globalState).F13 = (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int)
					}
					var _963_v83 _dafny.Sequence
					_ = _963_v83
					_963_v83 = _dafny.SeqOf(Companion_Default___.Fm30((_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), p2, _922_v57, true, globalState))
					var _964_v84 D4
					_ = _964_v84
					_964_v84 = Companion_D4_.Create_DC11_(_963_v83)
					var _965_v85 _dafny.Map
					_ = _965_v85
					_965_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _964_v84)
					var _966_v86 _dafny.Sequence
					_ = _966_v86
					_966_v86 = _dafny.SeqOf((_this).Fm23(_965_v85, (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), p0, globalState))
					_966_v86 = _dafny.Companion_Sequence_.Update(_966_v86, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm30(p0, p2, _922_v57, p2, globalState)).Cardinality()), _dafny.IntOfUint32((_966_v86).Cardinality()))).Uint32(), (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int))
					var _967_v87 D11
					_ = _967_v87
					_967_v87 = Companion_D11_.Create_DC29_()
					var _source18 D11 = _967_v87
					_ = _source18
					if _source18.Is_DC29() {
						var _968_v88 bool
						_ = _968_v88
						_968_v88 = false
						_968_v88 = true
						(globalState).F7 = ((_dafny.Zero).Minus(p0)).Plus((_dafny.Zero).Minus(p0))
						_968_v88 = true
						var _969_v89 _dafny.Array
						_ = _969_v89
						var _nw134 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
						_ = _nw134
						_969_v89 = _nw134
						var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
						_ = _nw135
						_969_v89 = _nw135
					} else if _source18.Is_DC30() {
						var _970___mcc_h2 bool = _source18.Get_().(D11_DC30).Cf46
						_ = _970___mcc_h2
						var _971___mcc_h3 _dafny.Int = _source18.Get_().(D11_DC30).Cf47
						_ = _971___mcc_h3
						var _972_cf47 _dafny.Int = _971___mcc_h3
						_ = _972_cf47
						var _973_cf46 bool = _970___mcc_h2
						_ = _973_cf46
						var _974_v90 _dafny.CodePoint
						_ = _974_v90
						_974_v90 = _dafny.CodePoint('b')
						var _975_v91 _dafny.Array
						_ = _975_v91
						var _nwElement0_30 _dafny.Int = (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int)
						_ = _nwElement0_30
						var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(2))
						_ = _nw136
						(_nw136).ArraySet1(_nwElement0_30, 0)
						(_nw136).ArraySet1((Companion_Default___.Fm41(_974_v90, globalState)).Cardinality(), 1)
						_975_v91 = _nw136
						var _976_v92 D2
						_ = _976_v92
						_976_v92 = Companion_D2_.Create_DC7_(_973_cf46, p1, _975_v91)
						var _977_v93 D2
						_ = _977_v93
						_977_v93 = Companion_D2_.Create_DC8_(_976_v92)
						var _pat_let_tv12 = _976_v92
						_ = _pat_let_tv12
						var _978_v94 *C3
						_ = _978_v94
						var _nw137 *C3 = New_C3_()
						_ = _nw137
						_nw137.Ctor__(func(_pat_let10_0 D2) D2 {
							return func(_979_dt__update__tmp_h1 D2) D2 {
								return func(_pat_let11_0 D2) D2 {
									return func(_980_dt__update_hcf11_h0 D2) D2 {
										return Companion_D2_.Create_DC8_(_980_dt__update_hcf11_h0)
									}(_pat_let11_0)
								}(_pat_let_tv12)
							}(_pat_let10_0)
						}(_977_v93), _973_cf46)
						_978_v94 = _nw137
						_973_cf46 = p2
						_974_v90 = _974_v90
						var _981_v95 _dafny.Array
						_ = _981_v95
						var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
						_ = _nw138
						_981_v95 = _nw138
						var _982_v96 D7
						_ = _982_v96
						_982_v96 = Companion_D7_.Create_DC21_((_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(508), _973_cf46, (_this).Fm8(p0, _978_v94.F25, (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), globalState), _978_v94.F25)
						var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_981_v95), 0))
						_ = _index157
						(_981_v95).ArraySet1((func() _dafny.Sequence {
							if (_982_v96).Dtor_cf32() {
								return _dafny.Companion_Sequence_.Update(_966_v86, (Companion_Default___.SafeIndex((_this).Fm23(_965_v85, (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), _972_cf47, globalState), _dafny.IntOfUint32((_966_v86).Cardinality()))).Uint32(), _972_cf47)
							}
							return _966_v86
						})(), (_index157).Int())
						var _983_v97 _dafny.Array
						_ = _983_v97
						var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
						_ = _nw139
						_983_v97 = _nw139
						var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_983_v97), 0))
						_ = _index158
						(_983_v97).ArraySet1(_853_v0, (_index158).Int())
						var _984_v98 _dafny.Array
						_ = _984_v98
						var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
						_ = _nw140
						_984_v98 = _nw140
						var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_984_v98), 0))
						_ = _index159
						(_984_v98).ArraySet1(_921_v56, (_index159).Int())
						var _985_v99 _dafny.Map
						_ = _985_v99
						_985_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_853_v0, _921_v56)
						var _986_v100 _dafny.Sequence
						_ = _986_v100
						_986_v100 = _dafny.SeqOf(_966_v86, _966_v86)
						var _987_v101 _dafny.Sequence
						_ = _987_v101
						_987_v101 = _dafny.SeqOf(_853_v0)
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_981_v95), 0))
						_ = _index160
						var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_983_v97), 0))
						_ = _index161
						var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_984_v98), 0))
						_ = _index162
						var _rhs211 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
							if _978_v94.F25 {
								return _966_v86
							}
							return _966_v86
						})(), (Companion_Default___.SafeIndex(_972_cf47, _dafny.IntOfUint32(((func() _dafny.Sequence {
							if _978_v94.F25 {
								return _966_v86
							}
							return _966_v86
						})()).Cardinality()))).Uint32(), _972_cf47), (Companion_Default___.SafeIndex((p1).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
							if _978_v94.F25 {
								return _966_v86
							}
							return _966_v86
						})(), (Companion_Default___.SafeIndex(_972_cf47, _dafny.IntOfUint32(((func() _dafny.Sequence {
							if _978_v94.F25 {
								return _966_v86
							}
							return _966_v86
						})()).Cardinality()))).Uint32(), _972_cf47)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_986_v100).Cardinality())), _966_v86)
						_ = _rhs211
						var _rhs212 _dafny.Array = (func() _dafny.Array {
							if true {
								return (func() _dafny.Array {
									if p2 {
										return _853_v0
									}
									return _853_v0
								})()
							}
							return (func() _dafny.Array {
								if _973_cf46 {
									return _853_v0
								}
								return (_987_v101).Select((Companion_Default___.SafeIndex((_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_987_v101).Cardinality()))).Uint32()).(_dafny.Array)
							})()
						})()
						_ = _rhs212
						var _rhs213 bool = !(_973_cf46) || (_973_cf46)
						_ = _rhs213
						var _rhs214 _dafny.Sequence = (func() _dafny.Sequence {
							if _973_cf46 {
								return _dafny.Companion_Sequence_.Update(_921_v56, (Companion_Default___.SafeIndex(_972_cf47, _dafny.IntOfUint32((_921_v56).Cardinality()))).Uint32(), Companion_Default___.Fm38(true, globalState))
							}
							return _921_v56
						})()
						_ = _rhs214
						var _rhs215 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_853_v0, _921_v56)).Merge(_985_v99)
						_ = _rhs215
						var _lhs149 _dafny.Array = _981_v95
						_ = _lhs149
						var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_981_v95), 0))
						_ = _lhs150
						var _lhs151 _dafny.Array = _983_v97
						_ = _lhs151
						var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_983_v97), 0))
						_ = _lhs152
						var _lhs153 _dafny.Array = _984_v98
						_ = _lhs153
						var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_984_v98), 0))
						_ = _lhs154
						(_lhs149).ArraySet1(_rhs211, (_lhs150).Int())
						(_lhs151).ArraySet1(_rhs212, (_lhs152).Int())
						_973_cf46 = _rhs213
						(_lhs153).ArraySet1(_rhs214, (_lhs154).Int())
						_985_v99 = _rhs215
					} else if _source18.Is_DC28() {
						var _988___mcc_h4 _dafny.Set = _source18.Get_().(D11_DC28).Cf45
						_ = _988___mcc_h4
						var _989_cf45 _dafny.Set = _988___mcc_h4
						_ = _989_cf45
						var _990_v102 _dafny.Array
						_ = _990_v102
						var _nw141 _dafny.Array = _dafny.NewArrayWithValue(Companion_D11_.Default(), _dafny.IntOfInt64(20))
						_ = _nw141
						_990_v102 = _nw141
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_990_v102), 0))
						_ = _index163
						(_990_v102).ArraySet1(_967_v87, (_index163).Int())
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_990_v102), 0))
						_ = _index164
						(_990_v102).ArraySet1(_967_v87, (_index164).Int())
						var _991_v103 bool
						_ = _991_v103
						_991_v103 = false
						var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))
						_ = _index165
						var _rhs216 _dafny.Set = (_924_v59).Difference((_924_v59).Union(_924_v59))
						_ = _rhs216
						var _rhs217 bool = _991_v103
						_ = _rhs217
						var _rhs218 _dafny.Int = (p0).Plus((func() _dafny.Int {
							if p2 {
								return p0
							}
							return p0
						})())
						_ = _rhs218
						var _rhs219 _dafny.Int = _dafny.IntOfInt64(153)
						_ = _rhs219
						var _lhs155 *GlobalState = globalState
						_ = _lhs155
						var _lhs156 _dafny.Array = _927_v62
						_ = _lhs156
						var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))
						_ = _lhs157
						_924_v59 = _rhs216
						_991_v103 = _rhs217
						_lhs155.F15 = _rhs218
						(_lhs156).ArraySet1(_rhs219, (_lhs157).Int())
						_991_v103 = p2
						var _992_v104 *C1
						_ = _992_v104
						var _nw142 *C1 = New_C1_()
						_ = _nw142
						_nw142.Ctor__()
						_992_v104 = _nw142
					} else {
						var _993___mcc_h5 D11 = _source18.Get_().(D11_DC31).Cf48
						_ = _993___mcc_h5
						var _994_cf48 D11 = _993___mcc_h5
						_ = _994_cf48
						_919_v54 = _919_v54
						(globalState).F4 = p0
						var _995_v105 bool
						_ = _995_v105
						_995_v105 = false
						var _996_v106 _dafny.Array
						_ = _996_v106
						var _len0_27 _dafny.Int = _dafny.IntOfInt64(28)
						_ = _len0_27
						var _nw143 _dafny.Array
						_ = _nw143
						if _len0_27.Cmp(_dafny.Zero) == 0 {
							_nw143 = _dafny.NewArray(_len0_27)
						} else {
							var _init27 func(_dafny.Int) _dafny.Int = func(_997_i7 _dafny.Int) _dafny.Int {
								return (_997_i7).Times(_dafny.IntOfInt64(264))
							}
							_ = _init27
							var _element0_27 = _init27(_dafny.Zero)
							_ = _element0_27
							_nw143 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
							(_nw143).ArraySet1(_element0_27, 0)
							var _nativeLen0_27 = (_len0_27).Int()
							_ = _nativeLen0_27
							for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
								(_nw143).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
							}
						}
						_996_v106 = _nw143
						var _998_v107 D2
						_ = _998_v107
						_998_v107 = Companion_D2_.Create_DC5_(_996_v106)
						var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))
						_ = _index166
						var _rhs220 bool = Companion_Default___.Fm26(globalState)
						_ = _rhs220
						var _rhs221 _dafny.Int = p0
						_ = _rhs221
						var _rhs222 _dafny.Int = (_927_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))).Int()).(_dafny.Int)
						_ = _rhs222
						var _rhs223 D2 = p3
						_ = _rhs223
						var _rhs224 _dafny.Int = p0
						_ = _rhs224
						var _lhs158 *GlobalState = globalState
						_ = _lhs158
						var _lhs159 *GlobalState = globalState
						_ = _lhs159
						var _lhs160 _dafny.Array = _927_v62
						_ = _lhs160
						var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_927_v62), 0))
						_ = _lhs161
						_995_v105 = _rhs220
						_lhs158.F13 = _rhs221
						_lhs159.F13 = _rhs222
						_998_v107 = _rhs223
						(_lhs160).ArraySet1(_rhs224, (_lhs161).Int())
						var _999_v108 _dafny.CodePoint
						_ = _999_v108
						_999_v108 = _dafny.CodePoint('p')
						var _1000_v109 _dafny.Map
						_ = _1000_v109
						_1000_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_v108, _853_v0)
						_1000_v109 = (_1000_v109).Update(_999_v108, _853_v0)
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _1001_v110 bool
		_ = _1001_v110
		_1001_v110 = false
		_1001_v110 = _1001_v110
		var _hi6 _dafny.Int = p0
		_ = _hi6
		for _1002_i8 := (func() _dafny.Int {
			if false {
				return p0
			}
			return p0
		})(); _1002_i8.Cmp(_hi6) < 0; _1002_i8 = _1002_i8.Plus(_dafny.One) {
			var _1003_v111 _dafny.Array
			_ = _1003_v111
			var _nw144 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(19))
			_ = _nw144
			_1003_v111 = _nw144
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_853_v0), 0))
			_ = _index167
			(_853_v0).ArraySet1(p2, (_index167).Int())
			var _1004_v112 D2
			_ = _1004_v112
			_1004_v112 = Companion_D2_.Create_DC7_(_1001_v110, p1, _927_v62)
			var _1005_v113 _dafny.Set
			_ = _1005_v113
			_1005_v113 = _dafny.SetOf(_921_v56, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(204))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}(func(_1006_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})), _921_v56, _921_v56, _921_v56)
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_853_v0), 0))
			_ = _index168
			var _rhs225 bool = (func() bool {
				if _dafny.Companion_Sequence_.Equal(_921_v56, _921_v56) {
					return p2
				}
				return (_1004_v112).Dtor_cf8()
			})()
			_ = _rhs225
			var _rhs226 _dafny.Array = _1003_v111
			_ = _rhs226
			var _rhs227 bool = (_1005_v113).IsSubsetOf(_1005_v113)
			_ = _rhs227
			var _rhs228 bool = _1001_v110
			_ = _rhs228
			var _lhs162 _dafny.Array = _853_v0
			_ = _lhs162
			var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_853_v0), 0))
			_ = _lhs163
			_1001_v110 = _rhs225
			_1003_v111 = _rhs226
			(_lhs162).ArraySet1(_rhs227, (_lhs163).Int())
			_1001_v110 = _rhs228
			var _1007_v114 _dafny.Sequence
			_ = _1007_v114
			_1007_v114 = _dafny.SeqOf(_921_v56, _921_v56)
			var _1008_v115 D4
			_ = _1008_v115
			_1008_v115 = Companion_D4_.Create_DC11_(_1007_v114)
			var _1009_v116 _dafny.Map
			_ = _1009_v116
			_1009_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1001_v110, _1008_v115)
			var _1010_v117 _dafny.MultiSet
			_ = _1010_v117
			_1010_v117 = _dafny.MultiSetOf((_this).Fm23(_1009_v116, _1002_i8, _dafny.IntOfInt64(330), globalState), (_dafny.Zero).Minus(p0))
			_1001_v110 = (_1010_v117).IsSubsetOf(_1010_v117)
			var _1011_v118 _dafny.Sequence
			_ = _1011_v118
			_1011_v118 = _dafny.SeqOf(_1001_v110, Companion_Default___.Fm26(globalState), _1001_v110)
			_1011_v118 = _1011_v118
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_853_v0), 0))
			_ = _index169
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_853_v0), 0))
			_ = _index170
			var _rhs229 bool = _1001_v110
			_ = _rhs229
			var _rhs230 bool = (_1011_v118).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1011_v118).Cardinality()))).Uint32()).(bool)
			_ = _rhs230
			var _lhs164 _dafny.Array = _853_v0
			_ = _lhs164
			var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_853_v0), 0))
			_ = _lhs165
			var _lhs166 _dafny.Array = _853_v0
			_ = _lhs166
			var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_853_v0), 0))
			_ = _lhs167
			(_lhs164).ArraySet1(_rhs229, (_lhs165).Int())
			(_lhs166).ArraySet1(_rhs230, (_lhs167).Int())
		}
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	dummy byte
}

func New_C5_() *C5 {
	_this := C5{}

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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__() {
	{
	}
}
func (_this *C5) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(180)), _dafny.SetOf(_dafny.IntOfInt64(-596))))).Intersection(_dafny.MultiSetOf(func() _dafny.Set {
			var _coll38 = _dafny.NewBuilder()
			_ = _coll38
			for _iter39 := _dafny.Iterate((func() _dafny.Set {
				var _coll39 = _dafny.NewBuilder()
				_ = _coll39
				for _iter40 := _dafny.Iterate((func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(741), _dafny.IntOfInt64(602))); ; {
						_compr_40, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _1012_v0 _dafny.Int
						_1012_v0 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(741)).Cmp(_1012_v0) <= 0) && ((_1012_v0).Cmp(_dafny.IntOfInt64(602)) < 0) {
							_coll40.Add((_1012_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-193))).Cardinality())), _dafny.IntOfInt64(587))
						}
					}
					return _coll40.ToMap()
				}()).Keys().Elements()); ; {
					_compr_39, _ok40 := _iter40()
					if !_ok40 {
						break
					}
					var _1013_v1 _dafny.Int
					_1013_v1 = interface{}(_compr_39).(_dafny.Int)
					if (func() _dafny.Map {
						var _coll41 = _dafny.NewMapBuilder()
						_ = _coll41
						for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(741), _dafny.IntOfInt64(602))); ; {
							_compr_41, _ok42 := _iter42()
							if !_ok42 {
								break
							}
							var _1012_v0 _dafny.Int
							_1012_v0 = interface{}(_compr_41).(_dafny.Int)
							if ((_dafny.IntOfInt64(741)).Cmp(_1012_v0) <= 0) && ((_1012_v0).Cmp(_dafny.IntOfInt64(602)) < 0) {
								_coll41.Add((_1012_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-193))).Cardinality())), _dafny.IntOfInt64(587))
							}
						}
						return _coll41.ToMap()
					}()).Contains(_1013_v1) {
						_coll39.Add(Companion_Default___.SafeDivisionInt(_1013_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("we")).Cardinality())))
					}
				}
				return _coll39.ToSet()
			}()).Elements()); ; {
				_compr_38, _ok39 := _iter39()
				if !_ok39 {
					break
				}
				var _1014_v2 _dafny.Int
				_1014_v2 = interface{}(_compr_38).(_dafny.Int)
				if (func() _dafny.Set {
					var _coll42 = _dafny.NewBuilder()
					_ = _coll42
					for _iter43 := _dafny.Iterate((func() _dafny.Map {
						var _coll43 = _dafny.NewMapBuilder()
						_ = _coll43
						for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(741), _dafny.IntOfInt64(602))); ; {
							_compr_43, _ok44 := _iter44()
							if !_ok44 {
								break
							}
							var _1012_v0 _dafny.Int
							_1012_v0 = interface{}(_compr_43).(_dafny.Int)
							if ((_dafny.IntOfInt64(741)).Cmp(_1012_v0) <= 0) && ((_1012_v0).Cmp(_dafny.IntOfInt64(602)) < 0) {
								_coll43.Add((_1012_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-193))).Cardinality())), _dafny.IntOfInt64(587))
							}
						}
						return _coll43.ToMap()
					}()).Keys().Elements()); ; {
						_compr_42, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _1015_v1 _dafny.Int
						_1015_v1 = interface{}(_compr_42).(_dafny.Int)
						if (func() _dafny.Map {
							var _coll44 = _dafny.NewMapBuilder()
							_ = _coll44
							for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(741), _dafny.IntOfInt64(602))); ; {
								_compr_44, _ok45 := _iter45()
								if !_ok45 {
									break
								}
								var _1012_v0 _dafny.Int
								_1012_v0 = interface{}(_compr_44).(_dafny.Int)
								if ((_dafny.IntOfInt64(741)).Cmp(_1012_v0) <= 0) && ((_1012_v0).Cmp(_dafny.IntOfInt64(602)) < 0) {
									_coll44.Add((_1012_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-193))).Cardinality())), _dafny.IntOfInt64(587))
								}
							}
							return _coll44.ToMap()
						}()).Contains(_1015_v1) {
							_coll42.Add(Companion_Default___.SafeDivisionInt(_1015_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("we")).Cardinality())))
						}
					}
					return _coll42.ToSet()
				}()).Contains(_1014_v2) {
					_coll38.Add(Companion_Default___.SafeDivisionInt(_1014_v2, _dafny.IntOfInt64(769)))
				}
			}
			return _coll38.ToSet()
		}(), func() _dafny.Set {
			var _coll45 = _dafny.NewBuilder()
			_ = _coll45
			for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(374), _dafny.IntOfInt64(894))); ; {
				_compr_45, _ok46 := _iter46()
				if !_ok46 {
					break
				}
				var _1016_v3 _dafny.Int
				_1016_v3 = interface{}(_compr_45).(_dafny.Int)
				if ((_dafny.IntOfInt64(374)).Cmp(_1016_v3) <= 0) && ((_1016_v3).Cmp(_dafny.IntOfInt64(894)) < 0) {
					_coll45.Add((_1016_v3).Times(_dafny.IntOfInt64(245)))
				}
			}
			return _coll45.ToSet()
		}(), func() _dafny.Set {
			var _coll46 = _dafny.NewBuilder()
			_ = _coll46
			for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(666), _dafny.IntOfInt64(713))); ; {
				_compr_46, _ok47 := _iter47()
				if !_ok47 {
					break
				}
				var _1017_v4 _dafny.Int
				_1017_v4 = interface{}(_compr_46).(_dafny.Int)
				if ((_dafny.IntOfInt64(666)).Cmp(_1017_v4) <= 0) && ((_1017_v4).Cmp(_dafny.IntOfInt64(713)) < 0) {
					_coll46.Add((_1017_v4).Times(_dafny.IntOfInt64(144)))
				}
			}
			return _coll46.ToSet()
		}()))).Difference((_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(412), _dafny.IntOfInt64(183), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll47 = _dafny.NewMapBuilder()
			_ = _coll47
			for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(337), _dafny.IntOfInt64(-99))); ; {
				_compr_47, _ok48 := _iter48()
				if !_ok48 {
					break
				}
				var _1018_v5 _dafny.Int
				_1018_v5 = interface{}(_compr_47).(_dafny.Int)
				if ((_dafny.IntOfInt64(337)).Cmp(_1018_v5) <= 0) && ((_1018_v5).Cmp(_dafny.IntOfInt64(-99)) < 0) {
					_coll47.Add((_1018_v5).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(500))).Cardinality()), _dafny.IntOfInt64(853))
				}
			}
			return _coll47.ToMap()
		}()).Cardinality(), _dafny.IntOfInt64(-161))).Cardinality()), func() _dafny.Set {
			var _coll48 = _dafny.NewBuilder()
			_ = _coll48
			for _iter49 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(215), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()))).Elements()); ; {
				_compr_48, _ok49 := _iter49()
				if !_ok49 {
					break
				}
				var _1019_v6 _dafny.Int
				_1019_v6 = interface{}(_compr_48).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(215), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality()))).Contains(_1019_v6) {
					_coll48.Add(Companion_Default___.SafeDivisionInt(_1019_v6, _dafny.IntOfInt64(772)))
				}
			}
			return _coll48.ToSet()
		}(), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jghlcq")).Cardinality()), (_dafny.SetOf(false)).Cardinality()))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('a')))).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(138))))))
	}
}
func (_this *C5) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_dafny.IntOfInt64(99)).Times(_dafny.IntOfInt64(952))).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Cardinality())) > 0
	}
}
func (_this *C5) Fm20(globalState *GlobalState) bool {
	{
		return !((func() bool {
			if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(true, true), _dafny.SeqOf(false)) {
				return (!(!(false))) && (false)
			}
			return (_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-758)))).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(785), _dafny.IntOfInt64(488), _dafny.IntOfInt64(825)))
		})())
	}
}
func (_this *C5) M6(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 D2, globalState *GlobalState) {
	{
		(globalState).F15 = p0
		if p2 {
			var _1020_v0 bool
			_ = _1020_v0
			_1020_v0 = false
			var _1021_v1 _dafny.CodePoint
			_ = _1021_v1
			_1021_v1 = _dafny.CodePoint('g')
			var _1022_v2 _dafny.Set
			_ = _1022_v2
			_1022_v2 = _dafny.SetOf(_1020_v0)
			var _1023_v3 _dafny.Array
			_ = _1023_v3
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_28
			var _nw145 _dafny.Array
			_ = _nw145
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw145 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Int = (func(_1024_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1025_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1025_i0, _1024_p0)
					}
				})(p0)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw145 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw145).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw145).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1023_v3 = _nw145
			var _1026_v4 _dafny.Array
			_ = _1026_v4
			var _nwElement0_31 bool = true
			_ = _nwElement0_31
			var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(23))
			_ = _nw146
			(_nw146).ArraySet1(_nwElement0_31, 0)
			(_nw146).ArraySet1((true) == (_1020_v0), 1)
			(_nw146).ArraySet1(_1020_v0, 2)
			(_nw146).ArraySet1((Companion_Default___.Fm21(_dafny.CodePoint('m'), _1021_v1, globalState)).Dtor_cf4(), 3)
			(_nw146).ArraySet1(p2, 4)
			(_nw146).ArraySet1(p2, 5)
			(_nw146).ArraySet1((func() bool {
				if true {
					return _1020_v0
				}
				return true
			})(), 6)
			(_nw146).ArraySet1(_1020_v0, 7)
			(_nw146).ArraySet1((_this).Fm20(globalState), 8)
			(_nw146).ArraySet1((p1).IsDisjointFrom(p1), 9)
			(_nw146).ArraySet1(p2, 10)
			(_nw146).ArraySet1(p2, 11)
			(_nw146).ArraySet1(p2, 12)
			(_nw146).ArraySet1((_1022_v2).IsSubsetOf(_1022_v2), 13)
			(_nw146).ArraySet1(p2, 14)
			(_nw146).ArraySet1(_1020_v0, 15)
			(_nw146).ArraySet1((_dafny.MultiSetOf(_1023_v3, _1023_v3)).IsProperSubsetOf((_dafny.MultiSetOf(_1023_v3)).Update(_1023_v3, Companion_Default___.Abs((_dafny.Zero).Minus(p0)))), 16)
			(_nw146).ArraySet1(p2, 17)
			(_nw146).ArraySet1((_this).Fm20(globalState), 18)
			(_nw146).ArraySet1(false, 19)
			(_nw146).ArraySet1(!(p2), 20)
			(_nw146).ArraySet1(p2, 21)
			(_nw146).ArraySet1((_this).Fm8(p0, p2, p0, globalState), 22)
			_1026_v4 = _nw146
			var _1027_v5 _dafny.Set
			_ = _1027_v5
			_1027_v5 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("g"))
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))
			_ = _index171
			(_1026_v4).ArraySet1((_1027_v5).IsSubsetOf(_1027_v5), (_index171).Int())
			var _1028_v6 _dafny.Sequence
			_ = _1028_v6
			_1028_v6 = _dafny.UnicodeSeqOfUtf8Bytes("slclxhejc")
			var _1029_v7 _dafny.Sequence
			_ = _1029_v7
			_1029_v7 = _dafny.SeqOf(_1028_v6, _dafny.Companion_Sequence_.Update(_1028_v6, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1028_v6).Cardinality()))).Uint32(), _dafny.CodePoint('x')))
			var _1030_v8 D4
			_ = _1030_v8
			_1030_v8 = Companion_D4_.Create_DC11_(_1029_v7)
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))
			_ = _index172
			var _rhs231 bool = p2
			_ = _rhs231
			var _rhs232 bool = p2
			_ = _rhs232
			var _rhs233 _dafny.Sequence = (_1030_v8).Dtor_cf15()
			_ = _rhs233
			var _lhs168 _dafny.Array = _1026_v4
			_ = _lhs168
			var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))
			_ = _lhs169
			_1020_v0 = _rhs231
			(_lhs168).ArraySet1(_rhs232, (_lhs169).Int())
			_1029_v7 = _rhs233
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))
			_ = _index173
			(_1023_v3).ArraySet1(p0, (_index173).Int())
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))
			_ = _index174
			(_1023_v3).ArraySet1((func() _dafny.Int {
				if (_1026_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))).Int()).(bool) {
					return p0
				}
				return p0
			})(), (_index174).Int())
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))
			_ = _index175
			var _rhs234 bool = ((_1023_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))).Int()).(_dafny.Int)).Cmp(Companion_Default___.SafeDivisionInt((_1023_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))).Int()).(_dafny.Int), (_1023_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))).Int()).(_dafny.Int))) < 0
			_ = _rhs234
			var _rhs235 _dafny.CodePoint = _1021_v1
			_ = _rhs235
			var _lhs170 _dafny.Array = _1026_v4
			_ = _lhs170
			var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))
			_ = _lhs171
			(_lhs170).ArraySet1(_rhs234, (_lhs171).Int())
			_1021_v1 = _rhs235
			var _1031_v9 _dafny.Map
			_ = _1031_v9
			_1031_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1020_v0)
			var _1032_v10 _dafny.Map
			_ = _1032_v10
			_1032_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1026_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))).Int()).(bool), _1023_v3)
			var _1033_v11 _dafny.Map
			_ = _1033_v11
			_1033_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1032_v10).Cardinality()), _1031_v9)
			var _1034_v12 _dafny.Array
			_ = _1034_v12
			var _nwElement0_32 _dafny.Map = _1031_v9
			_ = _nwElement0_32
			var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(17))
			_ = _nw147
			(_nw147).ArraySet1(_nwElement0_32, 0)
			(_nw147).ArraySet1((_1031_v9).Merge(_1031_v9), 1)
			(_nw147).ArraySet1((_1031_v9).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)), 2)
			(_nw147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1026_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))).Int()).(bool), _1020_v0), 3)
			(_nw147).ArraySet1(_1031_v9, 4)
			(_nw147).ArraySet1((_1031_v9).Merge(_1031_v9), 5)
			(_nw147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1026_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))).Int()).(bool), p2), 6)
			(_nw147).ArraySet1(_1031_v9, 7)
			(_nw147).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1026_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))).Int()).(bool))).Update(p2, (_1026_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))).Int()).(bool)), 8)
			(_nw147).ArraySet1((_1031_v9).Merge(_1031_v9), 9)
			(_nw147).ArraySet1(_1031_v9, 10)
			(_nw147).ArraySet1((_1031_v9).Merge(_1031_v9), 11)
			(_nw147).ArraySet1(_1031_v9, 12)
			(_nw147).ArraySet1((func() _dafny.Map {
				if _1020_v0 {
					return Companion_Default___.Fm22(_1029_v7, (_1023_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))).Int()).(_dafny.Int), globalState)
				}
				return (func() _dafny.Map {
					if (_1033_v11).Contains(p0) {
						return (_1033_v11).Get(p0).(_dafny.Map)
					}
					return _1031_v9
				})()
			})(), 13)
			(_nw147).ArraySet1(_1031_v9, 14)
			(_nw147).ArraySet1(_1031_v9, 15)
			(_nw147).ArraySet1(_1031_v9, 16)
			_1034_v12 = _nw147
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_1034_v12), 0))
			_ = _index176
			(_1034_v12).ArraySet1((func() _dafny.Map {
				if false {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1020_v0, (_1026_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))).Int()).(bool))
				}
				return Companion_Default___.Fm22(_1029_v7, p0, globalState)
			})(), (_index176).Int())
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_1034_v12), 0))
			_ = _index177
			var _rhs236 bool = !(p2)
			_ = _rhs236
			var _rhs237 _dafny.Map = (_1031_v9).Merge(_1031_v9)
			_ = _rhs237
			var _rhs238 bool = (_1026_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_1026_v4), 0))).Int()).(bool)
			_ = _rhs238
			var _lhs172 _dafny.Array = _1034_v12
			_ = _lhs172
			var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_1034_v12), 0))
			_ = _lhs173
			_1020_v0 = _rhs236
			(_lhs172).ArraySet1(_rhs237, (_lhs173).Int())
			_1020_v0 = _rhs238
			var _1035_v13 D2
			_ = _1035_v13
			_1035_v13 = Companion_D2_.Create_DC6_(((_1023_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))).Int()).(_dafny.Int)).Plus((_1023_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))).Int()).(_dafny.Int)))
			var _pat_let_tv13 = p0
			_ = _pat_let_tv13
			_1035_v13 = func(_pat_let12_0 D2) D2 {
				return func(_1036_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let13_0 _dafny.Int) D2 {
						return func(_1037_dt__update_hcf7_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC6_(_1037_dt__update_hcf7_h0)
						}(_pat_let13_0)
					}(_pat_let_tv13)
				}(_pat_let12_0)
			}(Companion_D2_.Create_DC6_((_1023_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(876), _dafny.ArrayLen((_1023_v3), 0))).Int()).(_dafny.Int)))
		} else {
			var _1038_v14 _dafny.Array
			_ = _1038_v14
			var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw148
			_1038_v14 = _nw148
			var _1039_v15 D2
			_ = _1039_v15
			_1039_v15 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC5_(_1038_v14))
			var _1040_v16 D2
			_ = _1040_v16
			_1040_v16 = Companion_D2_.Create_DC8_(_1039_v15)
			var _pat_let_tv14 = _1039_v15
			_ = _pat_let_tv14
			var _1041_v17 *C3
			_ = _1041_v17
			var _nw149 *C3 = New_C3_()
			_ = _nw149
			_nw149.Ctor__(func(_pat_let14_0 D2) D2 {
				return func(_1042_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let15_0 D2) D2 {
						return func(_1043_dt__update_hcf11_h0 D2) D2 {
							return Companion_D2_.Create_DC8_(_1043_dt__update_hcf11_h0)
						}(_pat_let15_0)
					}(_pat_let_tv14)
				}(_pat_let14_0)
			}(_1040_v16), p2)
			_1041_v17 = _nw149
			var _1044_v18 _dafny.Sequence
			_ = _1044_v18
			_1044_v18 = _dafny.UnicodeSeqOfUtf8Bytes("g")
			var _1045_v19 _dafny.Set
			_ = _1045_v19
			_1045_v19 = _dafny.SetOf(p0)
			var _1046_v20 _dafny.MultiSet
			_ = _1046_v20
			_1046_v20 = _dafny.MultiSetOf((_1045_v19).Cardinality(), p0)
			var _1047_v21 _dafny.Sequence
			_ = _1047_v21
			_1047_v21 = _dafny.SeqOf(p0)
			var _1048_v22 _dafny.Array
			_ = _1048_v22
			var _nwElement0_33 _dafny.MultiSet = _1046_v20
			_ = _nwElement0_33
			var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(8))
			_ = _nw150
			(_nw150).ArraySet1(_nwElement0_33, 0)
			(_nw150).ArraySet1(_1046_v20, 1)
			(_nw150).ArraySet1((_1046_v20).Difference(_dafny.MultiSetFromSeq(_1047_v21)), 2)
			(_nw150).ArraySet1(_1046_v20, 3)
			(_nw150).ArraySet1(_1046_v20, 4)
			(_nw150).ArraySet1((_1046_v20).Union(_1046_v20), 5)
			(_nw150).ArraySet1((_1046_v20).Intersection((_1046_v20).Update(p0, Companion_Default___.Abs(p0))), 6)
			(_nw150).ArraySet1(_1046_v20, 7)
			_1048_v22 = _nw150
			var _1049_v23 _dafny.Sequence
			_ = _1049_v23
			_1049_v23 = _dafny.SeqOf(_1044_v18, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(790))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}(func(_1050_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})))
			var _1051_v24 _dafny.Map
			_ = _1051_v24
			_1051_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1041_v17.F25)
			var _1052_v25 _dafny.Map
			_ = _1052_v25
			_1052_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1051_v24).Cardinality(), _1041_v17.F25)
			var _1053_v26 D11
			_ = _1053_v26
			_1053_v26 = Companion_D11_.Create_DC30_(_1041_v17.F25, (_1051_v24).Cardinality())
			var _1054_v27 _dafny.Map
			_ = _1054_v27
			_1054_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1053_v26).Dtor_cf46(), p0)
			var _1055_v28 _dafny.Sequence
			_ = _1055_v28
			_1055_v28 = _dafny.SeqOf(_1041_v17.F25, p2)
			var _1056_v29 _dafny.CodePoint
			_ = _1056_v29
			_1056_v29 = _dafny.CodePoint('o')
			var _rhs239 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1044_v18, (_1049_v23).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1049_v23).Cardinality()))).Uint32()).(_dafny.Sequence))
			_ = _rhs239
			var _rhs240 _dafny.Array = _1048_v22
			_ = _rhs240
			var _rhs241 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm39((((_1054_v27).Update(Companion_Default___.Fm26(globalState), _dafny.IntOfInt64(-592))).Cardinality()).Minus(_dafny.IntOfUint32((_1055_v28).Cardinality())), _1056_v29, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(524))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}((func(_1057_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1058_i2 _dafny.Int) _dafny.Int {
					return _1057_p0
				}
			})(p0))), _1047_v21), p2, globalState)).Cardinality())
			_ = _rhs241
			var _rhs242 _dafny.Int = p0
			_ = _rhs242
			var _lhs174 *GlobalState = globalState
			_ = _lhs174
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			_1044_v18 = _rhs239
			_1048_v22 = _rhs240
			_lhs174.F4 = _rhs241
			_lhs175.F13 = _rhs242
			var _1059_v30 _dafny.Map
			_ = _1059_v30
			_1059_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _1060_v31 _dafny.MultiSet
			_ = _1060_v31
			_1060_v31 = _dafny.MultiSetOf(_1059_v30)
			_1038_v14 = (func() _dafny.Array {
				if (_1060_v31).IsProperSubsetOf(_1060_v31) {
					return _1038_v14
				}
				return _1038_v14
			})()
			if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(374))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg64 _dafny.Int) interface{} {
					return coer64(arg64)
				}
			}((func(_1061_v29 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1062_i3 _dafny.Int) _dafny.CodePoint {
					return _1061_v29
				}
			})(_1056_v29))), _1044_v18), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(111))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}(func(_1063_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}))) {
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_1038_v14), 0))
				_ = _index178
				(_1038_v14).ArraySet1(((_1041_v17).Fm2((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1041_v17.F25, _dafny.UnicodeSeqOfUtf8Bytes("xs"))).Cardinality(), globalState)).Plus(_dafny.IntOfUint32((_1055_v28).Cardinality())), (_index178).Int())
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_1038_v14), 0))
				_ = _index179
				(_1038_v14).ArraySet1(p0, (_index179).Int())
				(_1041_v17).F25 = (_1041_v17.F25) == (_dafny.Companion_Sequence_.Equal(_1055_v28, _dafny.SeqOf(_1041_v17.F25)))
				(_1041_v17).F25 = ((_1038_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_1038_v14), 0))).Int()).(_dafny.Int)).Cmp((_1038_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_1038_v14), 0))).Int()).(_dafny.Int)) < 0
				(_1041_v17).F25 = true
				_1059_v30 = (_1059_v30).Update(false, (_1041_v17.F25) || (p2))
			} else {
				_1056_v29 = Companion_Default___.Fm38((_1041_v17.F25) == (p2), globalState)
				_1056_v29 = _1056_v29
				var _1064_v32 _dafny.Set
				_ = _1064_v32
				_1064_v32 = _dafny.SetOf((_1055_v28).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1055_v28).Cardinality()))).Uint32()).(bool))
				(globalState).F4 = (_dafny.Zero).Minus(((_1059_v30).Cardinality()).Plus((_dafny.IntOfUint32((_1044_v18).Cardinality())).Plus((_dafny.MultiSetOf((_1064_v32).Cardinality())).Cardinality())))
				(globalState).F13 = p0
				(_1041_v17).F25 = (func() bool {
					if p2 {
						return p2
					}
					return _1041_v17.F25
				})()
			}
			(_1041_v17).F25 = p2
		}
		var _1065_v33 _dafny.Sequence
		_ = _1065_v33
		_1065_v33 = _dafny.UnicodeSeqOfUtf8Bytes("phf")
		var _1066_v34 _dafny.Map
		_ = _1066_v34
		_1066_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1065_v33, p2)
		var _1067_i5 _dafny.Int
		_ = _1067_i5
		_1067_i5 = _dafny.Zero
		{
			for !((func() bool {
				if (_1066_v34).Contains(_1065_v33) {
					return (_1066_v34).Get(_1065_v33).(bool)
				}
				return p2
			})()) {
				{
					if (_1067_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1067_i5 = (_1067_i5).Plus(_dafny.One)
					var _1068_v35 _dafny.Map
					_ = _1068_v35
					_1068_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(175), p0)
					_1068_v35 = (_1068_v35).Update(Companion_Default___.SafeModuloInt(p0, p0), p0)
					_1068_v35 = _1068_v35
					var _1069_v36 _dafny.CodePoint
					_ = _1069_v36
					_1069_v36 = _dafny.CodePoint('r')
					_1069_v36 = _1069_v36
					var _1070_v37 _dafny.Set
					_ = _1070_v37
					_1070_v37 = _dafny.SetOf(p2)
					var _1071_v38 _dafny.Array
					_ = _1071_v38
					var _nwElement0_34 _dafny.Set = _1070_v37
					_ = _nwElement0_34
					var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(17))
					_ = _nw151
					(_nw151).ArraySet1(_nwElement0_34, 0)
					(_nw151).ArraySet1(_1070_v37, 1)
					(_nw151).ArraySet1(_1070_v37, 2)
					(_nw151).ArraySet1(Companion_Default___.Fm37(globalState), 3)
					(_nw151).ArraySet1(_1070_v37, 4)
					(_nw151).ArraySet1((_1070_v37).Union(Companion_Default___.Fm37(globalState)), 5)
					(_nw151).ArraySet1((func() _dafny.Set {
						if p2 {
							return _1070_v37
						}
						return _1070_v37
					})(), 6)
					(_nw151).ArraySet1(_1070_v37, 7)
					(_nw151).ArraySet1(_1070_v37, 8)
					(_nw151).ArraySet1(_dafny.SetOf(p2), 9)
					(_nw151).ArraySet1(_dafny.SetOf(p2), 10)
					(_nw151).ArraySet1(_1070_v37, 11)
					(_nw151).ArraySet1(_1070_v37, 12)
					(_nw151).ArraySet1(_dafny.SetOf(!(p2)), 13)
					(_nw151).ArraySet1(_1070_v37, 14)
					(_nw151).ArraySet1(_1070_v37, 15)
					(_nw151).ArraySet1(_dafny.SetOf(!(p2), p2, false, p2), 16)
					_1071_v38 = _nw151
					var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_1071_v38), 0))
					_ = _index180
					(_1071_v38).ArraySet1(_1070_v37, (_index180).Int())
					var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_1071_v38), 0))
					_ = _index181
					(_1071_v38).ArraySet1(_1070_v37, (_index181).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1072_v39 _dafny.Array
		_ = _1072_v39
		var _nw152 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
		_ = _nw152
		_1072_v39 = _nw152
		var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_1072_v39), 0))
		_ = _index182
		(_1072_v39).ArraySet1(true, (_index182).Int())
		var _1073_v40 bool
		_ = _1073_v40
		_1073_v40 = false
		var _1074_v41 _dafny.Map
		_ = _1074_v41
		_1074_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
		var _1075_v42 _dafny.Map
		_ = _1075_v42
		_1075_v42 = _1074_v41
		var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_1072_v39), 0))
		_ = _index183
		var _rhs243 _dafny.Int = (p0).Plus(p0)
		_ = _rhs243
		var _rhs244 bool = _1073_v40
		_ = _rhs244
		var _rhs245 bool = !(p2)
		_ = _rhs245
		var _rhs246 _dafny.Map = _1075_v42
		_ = _rhs246
		var _lhs176 *GlobalState = globalState
		_ = _lhs176
		var _lhs177 _dafny.Array = _1072_v39
		_ = _lhs177
		var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_1072_v39), 0))
		_ = _lhs178
		_lhs176.F13 = _rhs243
		(_lhs177).ArraySet1(_rhs244, (_lhs178).Int())
		_1073_v40 = _rhs245
		_1075_v42 = _rhs246
		var _1076_v43 _dafny.Array
		_ = _1076_v43
		var _nw153 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw153
		_1076_v43 = _nw153
		_1076_v43 = _1076_v43
		var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_1072_v39), 0))
		_ = _index184
		(_1072_v39).ArraySet1(p2, (_index184).Int())
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	dummy byte
}

func New_C6_() *C6 {
	_this := C6{}

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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__() {
	{
	}
}
func (_this *C6) Fm15(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Set {
	{
		return ((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(353))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg66 _dafny.Int) interface{} {
				return coer66(arg66)
			}
		}(func(_1077_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _dafny.UnicodeSeqOfUtf8Bytes("vmtigl"))).Intersection(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ggpkkfhx"), _dafny.UnicodeSeqOfUtf8Bytes("mo"), _dafny.UnicodeSeqOfUtf8Bytes("jhtp"), _dafny.UnicodeSeqOfUtf8Bytes("rhv")))).Union(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-201))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg67 _dafny.Int) interface{} {
				return coer67(arg67)
			}
		}(func(_1078_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))))
	}
}
func (_this *C6) Fm16(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C6) M9(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) (_dafny.Int, _dafny.MultiSet, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r1
		var r2 bool = false
		_ = r2
		var _1079_v0 _dafny.MultiSet
		_ = _1079_v0
		_1079_v0 = _dafny.MultiSetOf(p1, (_this).Fm16(p0, globalState))
		var _1080_v1 _dafny.Array
		_ = _1080_v1
		var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw154
		_1080_v1 = _nw154
		var _1081_v2 D2
		_ = _1081_v2
		_1081_v2 = Companion_D2_.Create_DC7_(true, _1079_v0, _1080_v1)
		var _1082_v3 _dafny.MultiSet
		_ = _1082_v3
		_1082_v3 = _dafny.MultiSetOf(_1081_v2, _1081_v2, _1081_v2, _1081_v2, _1081_v2)
		var _1083_v4 _dafny.Sequence
		_ = _1083_v4
		_1083_v4 = _dafny.SeqOf(_1082_v3)
		var _1084_i0 _dafny.Int
		_ = _1084_i0
		_1084_i0 = _dafny.Zero
		{
			for ((_1083_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1083_v4).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsSubsetOf(_1082_v3) {
				{
					if (_1084_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1084_i0 = (_1084_i0).Plus(_dafny.One)
					var _1085_v5 _dafny.CodePoint
					_ = _1085_v5
					_1085_v5 = _dafny.CodePoint('b')
					_1085_v5 = _dafny.CodePoint('x')
					(globalState).F4 = p0
					r2 = (p0).Cmp(p0) != 0
					var _1086_v6 _dafny.Sequence
					_ = _1086_v6
					_1086_v6 = _dafny.SeqOf(Companion_Default___.Fm17(p0, globalState), p0)
					var _1087_v7 _dafny.Map
					_ = _1087_v7
					_1087_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
					var _1088_v8 _dafny.Set
					_ = _1088_v8
					_1088_v8 = _dafny.SetOf((_1086_v6).Select((Companion_Default___.SafeIndex((_1087_v7).Cardinality(), _dafny.IntOfUint32((_1086_v6).Cardinality()))).Uint32()).(_dafny.Int))
					r2 = (_1088_v8).IsDisjointFrom(_1088_v8)
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1089_v9 _dafny.Array
		_ = _1089_v9
		var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
		_ = _nw155
		_1089_v9 = _nw155
		var _1090_v10 _dafny.Map
		_ = _1090_v10
		_1090_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(p1))
		var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1089_v9), 0))
		_ = _index185
		(_1089_v9).ArraySet1((_1090_v10).Merge(func() _dafny.Map {
			var _coll49 = _dafny.NewMapBuilder()
			_ = _coll49
			for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(693), _dafny.IntOfInt64(711))); ; {
				_compr_49, _ok50 := _iter50()
				if !_ok50 {
					break
				}
				var _1091_v11 _dafny.Int
				_1091_v11 = interface{}(_compr_49).(_dafny.Int)
				if ((_dafny.IntOfInt64(693)).Cmp(_1091_v11) <= 0) && ((_1091_v11).Cmp(_dafny.IntOfInt64(711)) < 0) {
					_coll49.Add(Companion_Default___.SafeModuloInt(_1091_v11, (_dafny.Zero).Minus(p0)), p1)
				}
			}
			return _coll49.ToMap()
		}()), (_index185).Int())
		var _1092_v12 _dafny.CodePoint
		_ = _1092_v12
		_1092_v12 = _dafny.CodePoint('h')
		var _1093_v13 _dafny.MultiSet
		_ = _1093_v13
		_1093_v13 = _dafny.MultiSetOf(_1092_v12, _1092_v12)
		var _1094_v14 _dafny.Map
		_ = _1094_v14
		_1094_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1093_v13)
		var _1095_v15 _dafny.Sequence
		_ = _1095_v15
		_1095_v15 = _dafny.UnicodeSeqOfUtf8Bytes("yvqumpwp")
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1089_v9), 0))
		_ = _index186
		(_1089_v9).ArraySet1((Companion_Default___.Fm18(p0, _dafny.IntOfInt64(643), (func() _dafny.MultiSet {
			if (_1094_v14).Contains(_dafny.IntOfUint32((_1095_v15).Cardinality())) {
				return (_1094_v14).Get(_dafny.IntOfUint32((_1095_v15).Cardinality())).(_dafny.MultiSet)
			}
			return _1093_v13
		})(), globalState)).Update(p0, (p2) == (p1)), (_index186).Int())
		var _1096_v16 _dafny.Array
		_ = _1096_v16
		var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
		_ = _nw156
		_1096_v16 = _nw156
		var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))
		_ = _index187
		(_1096_v16).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm19(p1, p1, globalState), _1095_v15), (_index187).Int())
		var _1097_v17 _dafny.Array
		_ = _1097_v17
		var _nwElement0_35 _dafny.Array = _1080_v1
		_ = _nwElement0_35
		var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(14))
		_ = _nw157
		(_nw157).ArraySet1(_nwElement0_35, 0)
		(_nw157).ArraySet1(_1080_v1, 1)
		(_nw157).ArraySet1(_1080_v1, 2)
		(_nw157).ArraySet1(_1080_v1, 3)
		(_nw157).ArraySet1(_1080_v1, 4)
		(_nw157).ArraySet1(_1080_v1, 5)
		(_nw157).ArraySet1(_1080_v1, 6)
		(_nw157).ArraySet1(_1080_v1, 7)
		(_nw157).ArraySet1(_1080_v1, 8)
		(_nw157).ArraySet1(_1080_v1, 9)
		(_nw157).ArraySet1(_1080_v1, 10)
		(_nw157).ArraySet1(_1080_v1, 11)
		(_nw157).ArraySet1(_1080_v1, 12)
		(_nw157).ArraySet1(_1080_v1, 13)
		_1097_v17 = _nw157
		var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1097_v17), 0))
		_ = _index188
		(_1097_v17).ArraySet1(_1080_v1, (_index188).Int())
		var _1098_v18 _dafny.Map
		_ = _1098_v18
		_1098_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC2_(p0), _1095_v15)
		var _1099_v19 D1
		_ = _1099_v19
		_1099_v19 = Companion_D1_.Create_DC2_(p0)
		var _1100_v20 _dafny.Sequence
		_ = _1100_v20
		_1100_v20 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-938))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg68 _dafny.Int) interface{} {
				return coer68(arg68)
			}
		}((func(_1101_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1102_i1 _dafny.Int) _dafny.CodePoint {
				return _1101_v12
			}
		})(_1092_v12))))
		var _1103_v21 _dafny.Map
		_ = _1103_v21
		_1103_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
		var _1104_v22 _dafny.Map
		_ = _1104_v22
		_1104_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2) && (p1), _1080_v1)
		var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))
		_ = _index189
		var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1097_v17), 0))
		_ = _index190
		var _rhs247 _dafny.Int = (p0).Plus(p0)
		_ = _rhs247
		var _rhs248 _dafny.Sequence = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_1098_v18).Contains(_1099_v19) {
				return (_1098_v18).Get(_1099_v19).(_dafny.Sequence)
			}
			return _1095_v15
		})(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1098_v18).Contains(_1099_v19) {
				return (_1098_v18).Get(_1099_v19).(_dafny.Sequence)
			}
			return _1095_v15
		})()).Cardinality()))).Uint32(), _1092_v12)
		_ = _rhs248
		var _rhs249 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1095_v15, _1095_v15), (_1100_v20).Select((Companion_Default___.SafeIndex((_1103_v21).Cardinality(), _dafny.IntOfUint32((_1100_v20).Cardinality()))).Uint32()).(_dafny.Sequence))
		_ = _rhs249
		var _rhs250 _dafny.Array = (func() _dafny.Array {
			if (_1104_v22).Contains((p0).Cmp(p0) > 0) {
				return (_1104_v22).Get((p0).Cmp(p0) > 0).(_dafny.Array)
			}
			return _1080_v1
		})()
		_ = _rhs250
		var _lhs179 *GlobalState = globalState
		_ = _lhs179
		var _lhs180 _dafny.Array = _1096_v16
		_ = _lhs180
		var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))
		_ = _lhs181
		var _lhs182 _dafny.Array = _1097_v17
		_ = _lhs182
		var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1097_v17), 0))
		_ = _lhs183
		_lhs179.F15 = _rhs247
		(_lhs180).ArraySet1(_rhs248, (_lhs181).Int())
		_1095_v15 = _rhs249
		(_lhs182).ArraySet1(_rhs250, (_lhs183).Int())
		if false {
			var _1105_v23 _dafny.Array
			_ = _1105_v23
			var _nw158 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw158
			_1105_v23 = _nw158
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
			_ = _index191
			(_1105_v23).ArraySet1(_dafny.Companion_Sequence_.Contains(_1095_v15, _1092_v12), (_index191).Int())
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
			_ = _index192
			(_1105_v23).ArraySet1(p1, (_index192).Int())
			var _1106_v24 D0
			_ = _1106_v24
			_1106_v24 = Companion_D0_.Create_DC1_(_1105_v23)
			_1106_v24 = _1106_v24
			var _1107_v25 _dafny.Sequence
			_ = _1107_v25
			_1107_v25 = _dafny.SeqOf((_1105_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))).Int()).(bool), false)
			var _1108_v26 _dafny.Map
			_ = _1108_v26
			_1108_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1107_v25, (func() bool {
				if (_1105_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))).Int()).(bool) {
					return false
				}
				return p1
			})())
			var _1109_v27 _dafny.Sequence
			_ = _1109_v27
			_1109_v27 = _dafny.SeqOf(_dafny.SetOf(_1092_v12))
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
			_ = _index193
			(_1105_v23).ArraySet1((func() bool {
				if (_1108_v26).Contains(_dafny.SeqOf(p1)) {
					return (_1108_v26).Get(_dafny.SeqOf(p1)).(bool)
				}
				return (func() bool {
					if ((_1089_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1089_v9), 0))).Int()).(_dafny.Map)).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_1109_v27).Cardinality()))) {
						return ((_1089_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1089_v9), 0))).Int()).(_dafny.Map)).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_1109_v27).Cardinality()))).(bool)
					}
					return p1
				})()
			})(), (_index193).Int())
			var _1110_v28 D0
			_ = _1110_v28
			_1110_v28 = Companion_D0_.Create_DC0_(_1105_v23)
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
			_ = _index194
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
			_ = _index195
			var _rhs251 _dafny.Int = (p0).Times(p0)
			_ = _rhs251
			var _rhs252 bool = (func() bool {
				if (_1090_v10).Contains(p0) {
					return (_1090_v10).Get(p0).(bool)
				}
				return p1
			})()
			_ = _rhs252
			var _rhs253 D0 = _1110_v28
			_ = _rhs253
			var _rhs254 bool = !(p2)
			_ = _rhs254
			var _rhs255 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1107_v25, _1107_v25)
			_ = _rhs255
			var _lhs184 *GlobalState = globalState
			_ = _lhs184
			var _lhs185 _dafny.Array = _1105_v23
			_ = _lhs185
			var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
			_ = _lhs186
			var _lhs187 _dafny.Array = _1105_v23
			_ = _lhs187
			var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
			_ = _lhs188
			_lhs184.F15 = _rhs251
			(_lhs185).ArraySet1(_rhs252, (_lhs186).Int())
			_1110_v28 = _rhs253
			(_lhs187).ArraySet1(_rhs254, (_lhs188).Int())
			_1107_v25 = _rhs255
			var _1111_v29 D1
			_ = _1111_v29
			_1111_v29 = Companion_D1_.Create_DC3_(_1095_v15, p1)
			if !((_1111_v29).Dtor_cf4()) || ((_1105_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))).Int()).(bool)) {
				r2 = (_1105_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))).Int()).(bool)
				var _1112_v30 _dafny.Sequence
				_ = _1112_v30
				_1112_v30 = _dafny.SeqOf(_1105_v23)
				var _1113_v31 _dafny.Set
				_ = _1113_v31
				_1113_v31 = _dafny.SetOf(_dafny.IntOfUint32((_1112_v30).Cardinality()))
				var _1114_v32 _dafny.Map
				_ = _1114_v32
				_1114_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1113_v31, p1)
				var _1115_v33 _dafny.Map
				_ = _1115_v33
				_1115_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
				r2 = (func() bool {
					if (_1114_v32).Contains((_dafny.SetOf(p0, (func() _dafny.Int {
						if (_1079_v0).Contains(p1) {
							return (_1079_v0).Multiplicity(p1)
						}
						return p0
					})())).Difference(_dafny.SetOf((_1115_v33).Cardinality(), _dafny.IntOfInt64(721)))) {
						return (_1114_v32).Get((_dafny.SetOf(p0, (func() _dafny.Int {
							if (_1079_v0).Contains(p1) {
								return (_1079_v0).Multiplicity(p1)
							}
							return p0
						})())).Difference(_dafny.SetOf((_1115_v33).Cardinality(), _dafny.IntOfInt64(721)))).(bool)
					}
					return !((_1105_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))).Int()).(bool))
				})()
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1080_v1), 0))
				_ = _index196
				(_1080_v1).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), (_index196).Int())
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_1080_v1), 0))
				_ = _index197
				(_1080_v1).ArraySet1(p0, (_index197).Int())
				_1095_v15 = _dafny.Companion_Sequence_.Concatenate((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence), _1095_v15)
				_1100_v20 = _1100_v20
			} else {
				var _1116_v34 D3
				_ = _1116_v34
				_1116_v34 = Companion_D3_.Create_DC9_(_1092_v12)
				_1092_v12 = (_1116_v34).Dtor_cf12()
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
				_ = _index198
				var _rhs256 _dafny.Int = p0
				_ = _rhs256
				var _rhs257 _dafny.Int = p0
				_ = _rhs257
				var _rhs258 bool = p1
				_ = _rhs258
				var _lhs189 *GlobalState = globalState
				_ = _lhs189
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				var _lhs191 _dafny.Array = _1105_v23
				_ = _lhs191
				var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
				_ = _lhs192
				_lhs189.F4 = _rhs256
				_lhs190.F4 = _rhs257
				(_lhs191).ArraySet1(_rhs258, (_lhs192).Int())
				var _1117_v35 *C0
				_ = _1117_v35
				var _nw159 *C0 = New_C0_()
				_ = _nw159
				_nw159.Ctor__()
				_1117_v35 = _nw159
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))
				_ = _index199
				(_1105_v23).ArraySet1((Companion_Default___.SafeModuloInt(p0, p0)).Cmp(p0) == 0, (_index199).Int())
				var _1118_v36 D15
				_ = _1118_v36
				_1118_v36 = Companion_D15_.Create_DC38_(_1090_v10)
				var _1119_v37 _dafny.MultiSet
				_ = _1119_v37
				_1119_v37 = _dafny.MultiSetOf((_1118_v36).Dtor_cf55(), (func() _dafny.Map {
					if p1 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), p2)
					}
					return Companion_Default___.Fm35(globalState)
				})(), (_1118_v36).Dtor_cf55())
				var _rhs259 _dafny.MultiSet = _1119_v37
				_ = _rhs259
				var _rhs260 bool = (_1105_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_1105_v23), 0))).Int()).(bool)
				_ = _rhs260
				_1119_v37 = _rhs259
				r2 = _rhs260
			}
		} else {
			if _dafny.Companion_Sequence_.IsPrefixOf((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Update((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1092_v12)) {
				(globalState).F15 = p0
				var _1120_v38 *C4
				_ = _1120_v38
				var _nw160 *C4 = New_C4_()
				_ = _nw160
				_nw160.Ctor__()
				_1120_v38 = _nw160
				var _1121_v39 _dafny.Sequence
				_ = _1121_v39
				_1121_v39 = _dafny.SeqOf(p2, (_1120_v38).Fm8((_dafny.MultiSetOf(p0)).Cardinality(), p1, p0, globalState), p2, p2)
				_1079_v0 = (Companion_Default___.Fm34(_dafny.IntOfInt64(603), _dafny.IntOfUint32((_1121_v39).Cardinality()), p2, globalState)).Union(_1079_v0)
				r2 = p2
				var _1122_v40 _dafny.Sequence
				_ = _1122_v40
				_1122_v40 = _dafny.SeqOf(_1089_v9)
				var _1123_v41 _dafny.Map
				_ = _1123_v41
				_1123_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1089_v9)
				var _1124_v42 _dafny.Set
				_ = _1124_v42
				_1124_v42 = _dafny.SetOf(false)
				var _1125_v43 _dafny.Array
				_ = _1125_v43
				var _nwElement0_36 _dafny.Array = _1089_v9
				_ = _nwElement0_36
				var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(28))
				_ = _nw161
				(_nw161).ArraySet1(_nwElement0_36, 0)
				(_nw161).ArraySet1(_1089_v9, 1)
				(_nw161).ArraySet1(_1089_v9, 2)
				(_nw161).ArraySet1(_1089_v9, 3)
				(_nw161).ArraySet1(_1089_v9, 4)
				(_nw161).ArraySet1(_1089_v9, 5)
				(_nw161).ArraySet1(_1089_v9, 6)
				(_nw161).ArraySet1(_1089_v9, 7)
				(_nw161).ArraySet1(_1089_v9, 8)
				(_nw161).ArraySet1(_1089_v9, 9)
				(_nw161).ArraySet1(_1089_v9, 10)
				(_nw161).ArraySet1((_1122_v40).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1122_v40).Cardinality()))).Uint32()).(_dafny.Array), 11)
				(_nw161).ArraySet1(_1089_v9, 12)
				(_nw161).ArraySet1(_1089_v9, 13)
				(_nw161).ArraySet1(_1089_v9, 14)
				(_nw161).ArraySet1(_1089_v9, 15)
				(_nw161).ArraySet1(_1089_v9, 16)
				(_nw161).ArraySet1(_1089_v9, 17)
				(_nw161).ArraySet1((func() _dafny.Array {
					if p2 {
						return _1089_v9
					}
					return (func() _dafny.Array {
						if (_1123_v41).Contains((_1124_v42).Cardinality()) {
							return (_1123_v41).Get((_1124_v42).Cardinality()).(_dafny.Array)
						}
						return (_1122_v40).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1122_v40).Cardinality()))).Uint32()).(_dafny.Array)
					})()
				})(), 18)
				(_nw161).ArraySet1(_1089_v9, 19)
				(_nw161).ArraySet1(_1089_v9, 20)
				(_nw161).ArraySet1((func() _dafny.Array {
					if false {
						return _1089_v9
					}
					return _1089_v9
				})(), 21)
				(_nw161).ArraySet1(_1089_v9, 22)
				(_nw161).ArraySet1(_1089_v9, 23)
				(_nw161).ArraySet1(_1089_v9, 24)
				(_nw161).ArraySet1(_1089_v9, 25)
				(_nw161).ArraySet1(_1089_v9, 26)
				(_nw161).ArraySet1(_1089_v9, 27)
				_1125_v43 = _nw161
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1125_v43), 0))
				_ = _index200
				(_1125_v43).ArraySet1(_1089_v9, (_index200).Int())
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_1125_v43), 0))
				_ = _index201
				(_1125_v43).ArraySet1(_1089_v9, (_index201).Int())
			} else {
				(globalState).F4 = Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(887))
				r2 = !(((p0).Times(_dafny.IntOfInt64(-924))).Cmp(Companion_Default___.SafeDivisionInt(p0, p0)) != 0)
				var _1126_v44 _dafny.Sequence
				_ = _1126_v44
				_1126_v44 = _dafny.SeqOf((Companion_Default___.Fm33(p1, p1, p0, globalState)).Cardinality(), p0)
				r2 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1126_v44, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-375))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg69 _dafny.Int) interface{} {
						return coer69(arg69)
					}
				}(func(_1127_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(535)
				}))), _1126_v44)
				var _1128_v45 _dafny.Sequence
				_ = _1128_v45
				_1128_v45 = _dafny.SeqOf(!(!(p2)) || (p2))
				var _1129_v46 _dafny.Sequence
				_ = _1129_v46
				_1129_v46 = _dafny.SeqOf(_1103_v21)
				var _1130_v47 _dafny.Sequence
				_ = _1130_v47
				_1130_v47 = _dafny.SeqOf(_1129_v46)
				var _1131_v48 _dafny.MultiSet
				_ = _1131_v48
				_1131_v48 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq((_1130_v47).Select((Companion_Default___.SafeIndex(((_1089_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1089_v9), 0))).Int()).(_dafny.Map)).Cardinality(), _dafny.IntOfUint32((_1130_v47).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality())
				var _rhs261 _dafny.Int = p0
				_ = _rhs261
				var _rhs262 _dafny.Sequence = _dafny.SeqOf(p2, p1, !((_1131_v48).IsSubsetOf(_1131_v48)))
				_ = _rhs262
				var _lhs193 *GlobalState = globalState
				_ = _lhs193
				_lhs193.F15 = _rhs261
				_1128_v45 = _rhs262
				r2 = false
			}
			var _1132_v49 _dafny.MultiSet
			_ = _1132_v49
			_1132_v49 = _dafny.MultiSetOf(p0, p0)
			var _1133_v50 _dafny.Map
			_ = _1133_v50
			_1133_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint), _1132_v49)
			_1133_v50 = (_1133_v50).Update(_1092_v12, (_1132_v49).Union(_1132_v49))
			if Companion_Default___.Fm26(globalState) {
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1080_v1), 0))
				_ = _index202
				(_1080_v1).ArraySet1(p0, (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1080_v1), 0))
				_ = _index203
				(_1080_v1).ArraySet1((Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1080_v1, Companion_Default___.Fm26(globalState))).Cardinality(), p0)).Minus(Companion_Default___.Fm17(p0, globalState)), (_index203).Int())
				r2 = p2
				r2 = (p1) == ((func() bool {
					if p1 {
						return p1
					}
					return p1
				})())
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1080_v1), 0))
				_ = _index204
				(_1080_v1).ArraySet1((p0).Times((_1080_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1080_v1), 0))).Int()).(_dafny.Int)), (_index204).Int())
				r2 = true
			} else {
				var _1134_v51 *C2
				_ = _1134_v51
				var _nw162 *C2 = New_C2_()
				_ = _nw162
				_nw162.Ctor__()
				_1134_v51 = _nw162
				var _1135_v52 _dafny.Array
				_ = _1135_v52
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_29
				var _nw163 _dafny.Array
				_ = _nw163
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw163 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) bool = (func(_1136_p1 bool) func(_dafny.Int) bool {
						return func(_1137_i3 _dafny.Int) bool {
							return _1136_p1
						}
					})(p1)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw163 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw163).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw163).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1135_v52 = _nw163
				var _1138_v53 _dafny.Map
				_ = _1138_v53
				_1138_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_dafny.SeqOf(_1135_v52)).Cardinality()))
				(globalState).F13 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1138_v53).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-592))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}((func(_1139_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1140_i4 _dafny.Int) _dafny.CodePoint {
						return (Companion_D3_.Create_DC9_(_1139_v12)).Dtor_cf12()
					}
				})(_1092_v12)))).Cardinality())))
				r2 = p2
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_1135_v52), 0))
				_ = _index205
				(_1135_v52).ArraySet1(p1, (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_1135_v52), 0))
				_ = _index206
				(_1135_v52).ArraySet1(!(p2) || (p2), (_index206).Int())
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_1097_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1097_v17), 0))).Int()))
				_ = _arr0
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_dafny.ArrayCastTo((_1097_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1097_v17), 0))).Int()))), 0))
				_ = _index207
				(_arr0).ArraySet1(_dafny.IntOfUint32((_1095_v15).Cardinality()), (_index207).Int())
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_1097_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1097_v17), 0))).Int()))
				_ = _arr1
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(200), _dafny.ArrayLen((_dafny.ArrayCastTo((_1097_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1097_v17), 0))).Int()))), 0))
				_ = _index208
				(_arr1).ArraySet1(p0, (_index208).Int())
			}
			if true {
				r2 = (func() bool {
					if (_1103_v21).Contains(!(p2)) {
						return (_1103_v21).Get(!(p2)).(bool)
					}
					return (_1079_v0).IsProperSubsetOf(_1079_v0)
				})()
				var _1141_v54 _dafny.Map
				_ = _1141_v54
				_1141_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1142_v57 _dafny.Array
				_ = _1142_v57
				var _nwElement0_37 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(387), p0)
				_ = _nwElement0_37
				var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(14))
				_ = _nw164
				(_nw164).ArraySet1(_nwElement0_37, 0)
				(_nw164).ArraySet1(_1141_v54, 1)
				(_nw164).ArraySet1(_1141_v54, 2)
				(_nw164).ArraySet1(_1141_v54, 3)
				(_nw164).ArraySet1((_1141_v54).Update((_dafny.Zero).Minus(p0), p0), 4)
				(_nw164).ArraySet1(_1141_v54, 5)
				(_nw164).ArraySet1(_1141_v54, 6)
				(_nw164).ArraySet1(_1141_v54, 7)
				(_nw164).ArraySet1(((_1141_v54).Update(p0, _dafny.IntOfInt64(733))).Merge(_1141_v54), 8)
				(_nw164).ArraySet1((func() _dafny.Map {
					if p2 {
						return _1141_v54
					}
					return _1141_v54
				})(), 9)
				(_nw164).ArraySet1(func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(337), _dafny.IntOfInt64(833))); ; {
						_compr_50, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1143_v55 _dafny.Int
						_1143_v55 = interface{}(_compr_50).(_dafny.Int)
						if ((_dafny.IntOfInt64(337)).Cmp(_1143_v55) <= 0) && ((_1143_v55).Cmp(_dafny.IntOfInt64(833)) < 0) {
							_coll50.Add((_1143_v55).Minus(p0), p0)
						}
					}
					return _coll50.ToMap()
				}(), 10)
				(_nw164).ArraySet1(func() _dafny.Map {
					var _coll51 = _dafny.NewMapBuilder()
					_ = _coll51
					for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(132), _dafny.IntOfInt64(774))); ; {
						_compr_51, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _1144_v56 _dafny.Int
						_1144_v56 = interface{}(_compr_51).(_dafny.Int)
						if ((_dafny.IntOfInt64(132)).Cmp(_1144_v56) <= 0) && ((_1144_v56).Cmp(_dafny.IntOfInt64(774)) < 0) {
							_coll51.Add(Companion_Default___.SafeDivisionInt(_1144_v56, p0), p0)
						}
					}
					return _coll51.ToMap()
				}(), 11)
				(_nw164).ArraySet1(_1141_v54, 12)
				(_nw164).ArraySet1(_1141_v54, 13)
				_1142_v57 = _nw164
				_1142_v57 = _1142_v57
				var _1145_v58 _dafny.Array
				_ = _1145_v58
				var _nw165 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw165
				_1145_v58 = _nw165
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_1145_v58), 0))
				_ = _index209
				(_1145_v58).ArraySet1((_this).Fm16(p0, globalState), (_index209).Int())
				var _1146_v59 _dafny.Sequence
				_ = _1146_v59
				_1146_v59 = _dafny.SeqOf(false, p1)
				var _1147_v60 _dafny.Sequence
				_ = _1147_v60
				_1147_v60 = _dafny.SeqOf(_1132_v49, _dafny.MultiSetOf(p0, _dafny.IntOfInt64(209)))
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_1145_v58), 0))
				_ = _index210
				(_1145_v58).ArraySet1((_dafny.MultiSetFromSeq(_1147_v60)).Contains(((_1132_v49).Update(_dafny.IntOfUint32((_1146_v59).Cardinality()), Companion_Default___.Abs(p0))).Intersection(_dafny.MultiSetOf(p0))), (_index210).Int())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_1145_v58), 0))
				_ = _index211
				(_1145_v58).ArraySet1(p2, (_index211).Int())
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_1145_v58), 0))
				_ = _index212
				(_1145_v58).ArraySet1(p1, (_index212).Int())
			} else {
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))
				_ = _index213
				(_1096_v16).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-551))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1148_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1149_i5 _dafny.Int) _dafny.CodePoint {
						return _1148_v12
					}
				})(_1092_v12))), (_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence)), (_index213).Int())
				var _1150_v61 D4
				_ = _1150_v61
				_1150_v61 = Companion_D4_.Create_DC12_(p0)
				var _1151_v62 _dafny.MultiSet
				_ = _1151_v62
				_1151_v62 = _dafny.MultiSetOf((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence))
				var _1152_v63 _dafny.Map
				_ = _1152_v63
				_1152_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_v61, (_1151_v62).Cardinality())
				_1152_v63 = (_1152_v63).Update(Companion_D4_.Create_DC12_(p0), p0)
				r2 = p2
				(globalState).F15 = (_dafny.Zero).Minus(p0)
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1080_v1), 0))
				_ = _index214
				(_1080_v1).ArraySet1(p0, (_index214).Int())
				var _1153_v64 _dafny.Array
				_ = _1153_v64
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_30
				var _nw166 _dafny.Array
				_ = _nw166
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw166 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) bool = (func(_1154_p1 bool) func(_dafny.Int) bool {
						return func(_1155_i6 _dafny.Int) bool {
							return (func() bool {
								if _1154_p1 {
									return _1154_p1
								}
								return _1154_p1
							})()
						}
					})(p1)
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
				_1153_v64 = _nw166
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1153_v64), 0))
				_ = _index215
				(_1153_v64).ArraySet1((p2) == (p1), (_index215).Int())
				var _1156_v65 _dafny.Map
				_ = _1156_v65
				_1156_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
				var _1157_v66 _dafny.Sequence
				_ = _1157_v66
				_1157_v66 = _dafny.SeqOf(_dafny.IntOfInt64(885), p0)
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1080_v1), 0))
				_ = _index216
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1153_v64), 0))
				_ = _index217
				var _rhs263 _dafny.Int = ((_1156_v65).Merge(_1156_v65)).Cardinality()
				_ = _rhs263
				var _rhs264 _dafny.MultiSet = (_dafny.MultiSetOf(!(p1))).Update(Companion_Default___.Fm26(globalState), Companion_Default___.Abs((p0).Times((_1157_v66).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1157_v66).Cardinality()))).Uint32()).(_dafny.Int))))
				_ = _rhs264
				var _rhs265 bool = (((_1079_v0).Update(p1, Companion_Default___.Abs(p0))).Intersection(_1079_v0)).IsSubsetOf(_dafny.MultiSetOf(p1))
				_ = _rhs265
				var _rhs266 _dafny.Int = p0
				_ = _rhs266
				var _lhs194 _dafny.Array = _1080_v1
				_ = _lhs194
				var _lhs195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(87), _dafny.ArrayLen((_1080_v1), 0))
				_ = _lhs195
				var _lhs196 _dafny.Array = _1153_v64
				_ = _lhs196
				var _lhs197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_1153_v64), 0))
				_ = _lhs197
				var _lhs198 *GlobalState = globalState
				_ = _lhs198
				(_lhs194).ArraySet1(_rhs263, (_lhs195).Int())
				_1079_v0 = _rhs264
				(_lhs196).ArraySet1(_rhs265, (_lhs197).Int())
				_lhs198.F4 = _rhs266
			}
			(globalState).F15 = p0
		}
		var _1158_v67 D2
		_ = _1158_v67
		_1158_v67 = Companion_D2_.Create_DC7_(p1, _1079_v0, _dafny.ArrayCastTo((_1097_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1097_v17), 0))).Int())))
		var _1159_v68 D2
		_ = _1159_v68
		_1159_v68 = Companion_D2_.Create_DC8_(_1158_v67)
		var _1160_v69 *C3
		_ = _1160_v69
		var _nw167 *C3 = New_C3_()
		_ = _nw167
		_nw167.Ctor__(_1159_v68, false)
		_1160_v69 = _nw167
		var _hi7 _dafny.Int = p0
		_ = _hi7
		for _1161_i7 := p0; _1161_i7.Cmp(_hi7) < 0; _1161_i7 = _1161_i7.Plus(_dafny.One) {
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_1080_v1), 0))
			_ = _index218
			(_1080_v1).ArraySet1(_dafny.IntOfInt64(-834), (_index218).Int())
			var _1162_v70 _dafny.Set
			_ = _1162_v70
			_1162_v70 = _dafny.SetOf(_1161_i7)
			var _1163_v71 _dafny.Sequence
			_ = _1163_v71
			_1163_v71 = _dafny.SeqOf(_1162_v70)
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_1080_v1), 0))
			_ = _index219
			var _rhs267 _dafny.Int = p0
			_ = _rhs267
			var _rhs268 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1163_v71).Cardinality()))
			_ = _rhs268
			var _lhs199 _dafny.Array = _1080_v1
			_ = _lhs199
			var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_1080_v1), 0))
			_ = _lhs200
			var _lhs201 *GlobalState = globalState
			_ = _lhs201
			(_lhs199).ArraySet1(_rhs267, (_lhs200).Int())
			_lhs201.F15 = _rhs268
			r2 = _1160_v69.F25
			r2 = !(_1160_v69.F25)
			var _1164_v72 _dafny.MultiSet
			_ = _1164_v72
			_1164_v72 = _dafny.MultiSetOf(p0, (_1080_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_1080_v1), 0))).Int()).(_dafny.Int))
			(globalState).F4 = (func() _dafny.Int {
				if (_1164_v72).Contains(p0) {
					return (_1164_v72).Multiplicity(p0)
				}
				return p0
			})()
		}
		r0 = (_1160_v69).Fm2((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, ((_1089_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1089_v9), 0))).Int()).(_dafny.Map)).Cardinality())), globalState)
		var _1165_v73 _dafny.Map
		_ = _1165_v73
		_1165_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1166_v74 _dafny.Sequence
		_ = _1166_v74
		_1166_v74 = _dafny.SeqOf(_dafny.MultiSetOf((func() _dafny.Int {
			if (_1165_v73).Contains(p0) {
				return (_1165_v73).Get(p0).(_dafny.Int)
			}
			return p0
		})()))
		var _1167_v75 _dafny.Set
		_ = _1167_v75
		_1167_v75 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32(((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence)).Cardinality())))
		var _1168_v76 _dafny.MultiSet
		_ = _1168_v76
		_1168_v76 = _dafny.MultiSetOf((_1167_v75).Cardinality(), _dafny.IntOfUint32(((_1096_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1096_v16), 0))).Int()).(_dafny.Sequence)).Cardinality()))
		r1 = ((_1166_v74).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.IntOfUint32((_1166_v74).Cardinality()))).Uint32()).(_dafny.MultiSet)).Difference((_1168_v76).Update(p0, Companion_Default___.Abs(p0)))
		var _1169_v77 D15
		_ = _1169_v77
		_1169_v77 = Companion_D15_.Create_DC38_((_1089_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_1089_v9), 0))).Int()).(_dafny.Map))
		var _1170_v78 D15
		_ = _1170_v78
		_1170_v78 = Companion_D15_.Create_DC40_(_1169_v77)
		var _1171_v79 D15
		_ = _1171_v79
		_1171_v79 = Companion_D15_.Create_DC40_(_1169_v77)
		var _1172_v80 D15
		_ = _1172_v80
		_1172_v80 = Companion_D15_.Create_DC40_(_1169_v77)
		var _pat_let_tv15 = _1160_v69
		_ = _pat_let_tv15
		var _pat_let_tv16 = _1160_v69
		_ = _pat_let_tv16
		var _pat_let_tv17 = _1160_v69
		_ = _pat_let_tv17
		r2 = func(_source19 D15) bool {
			if _source19.Is_DC39() {
				var _1173___mcc_h0 _dafny.Map = _source19.Get_().(D15_DC39).Cf56
				_ = _1173___mcc_h0
				var _1174___mcc_h1 bool = _source19.Get_().(D15_DC39).Cf57
				_ = _1174___mcc_h1
				var _1175___mcc_h2 _dafny.Int = _source19.Get_().(D15_DC39).Cf58
				_ = _1175___mcc_h2
				var _1176_cf58 _dafny.Int = _1175___mcc_h2
				_ = _1176_cf58
				var _1177_cf57 bool = _1174___mcc_h1
				_ = _1177_cf57
				var _1178_cf56 _dafny.Map = _1173___mcc_h0
				_ = _1178_cf56
				return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_1177_cf57), _dafny.SeqOf(_pat_let_tv15.F25))
			} else if _source19.Is_DC38() {
				var _1179___mcc_h3 _dafny.Map = _source19.Get_().(D15_DC38).Cf55
				_ = _1179___mcc_h3
				var _1180_cf55 _dafny.Map = _1179___mcc_h3
				_ = _1180_cf55
				return _pat_let_tv16.F25
			} else {
				var _1181___mcc_h4 D15 = _source19.Get_().(D15_DC40).Cf59
				_ = _1181___mcc_h4
				var _1182_cf59 D15 = _1181___mcc_h4
				_ = _1182_cf59
				return _pat_let_tv17.F25
			}
		}(_1172_v80)
		return r0, r1, r2
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	dummy byte
}

func New_C7_() *C7 {
	_this := C7{}

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

func (_this *C7) Ctor__() {
	{
	}
}
func (_this *C7) Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("bppvmww"), _dafny.CodePoint('w'))
	}
}
func (_this *C7) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(293))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg72 _dafny.Int) interface{} {
				return coer72(arg72)
			}
		}(func(_1183_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32(((func() _dafny.Sequence {
				if true {
					return _dafny.SeqOf((_dafny.SetOf(true, !(false), true)).Cardinality())
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(888))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}(func(_1184_i1 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("po")).Cardinality()))
				}))
			})()).Cardinality())
		})))).Cardinality()
	}
}
func (_this *C7) Fm12(globalState *GlobalState) bool {
	{
		return (true) == (!((Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("dtaill"), true)).Dtor_cf4()))
	}
}
func (_this *C7) Fm13(p0 D2, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jkonhsjjq"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(542))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg74 _dafny.Int) interface{} {
				return coer74(arg74)
			}
		}(func(_1185_i0 _dafny.Int) _dafny.CodePoint {
			return (Companion_D3_.Create_DC9_(_dafny.CodePoint('l'))).Dtor_cf12()
		})))
	}
}
func (_this *C7) M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _1186_v0 bool
		_ = _1186_v0
		_1186_v0 = true
		var _1187_v1 D1
		_ = _1187_v1
		_1187_v1 = Companion_D1_.Create_DC3_(p0, _1186_v0)
		_1186_v0 = (func() bool {
			if (_1187_v1).Dtor_cf4() {
				return _1186_v0
			}
			return _1186_v0
		})()
		var _1188_v2 _dafny.Int
		_ = _1188_v2
		_1188_v2 = _dafny.IntOfInt64(792)
		var _1189_v3 _dafny.Map
		_ = _1189_v3
		_1189_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1188_v2, false)
		var _1190_v4 _dafny.Array
		_ = _1190_v4
		var _nwElement0_38 _dafny.Int = (_dafny.Zero).Minus(_1188_v2)
		_ = _nwElement0_38
		var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(5))
		_ = _nw168
		(_nw168).ArraySet1(_nwElement0_38, 0)
		(_nw168).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_1189_v3).Cardinality())), 1)
		(_nw168).ArraySet1(_1188_v2, 2)
		(_nw168).ArraySet1(_1188_v2, 3)
		(_nw168).ArraySet1(_1188_v2, 4)
		_1190_v4 = _nw168
		var _1191_v5 _dafny.Map
		_ = _1191_v5
		_1191_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1190_v4, _dafny.IntOfUint32((_dafny.SeqOf(_1188_v2)).Cardinality()))
		_1191_v5 = _1191_v5
		var _1192_v6 _dafny.Map
		_ = _1192_v6
		_1192_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1186_v0, _1188_v2)
		_1192_v6 = _1192_v6
		var _1193_v7 _dafny.MultiSet
		_ = _1193_v7
		_1193_v7 = _dafny.MultiSetOf(_1186_v0)
		var _1194_v8 _dafny.Sequence
		_ = _1194_v8
		_1194_v8 = _dafny.SeqOf(_dafny.IntOfInt64(78))
		var _1195_i0 _dafny.Int
		_ = _1195_i0
		_1195_i0 = _dafny.Zero
		{
			for !(!(((func() _dafny.Int {
				if (_1193_v7).Contains(!(_1186_v0)) {
					return (_1193_v7).Multiplicity(!(_1186_v0))
				}
				return (_1194_v8).Select((Companion_Default___.SafeIndex(_1188_v2, _dafny.IntOfUint32((_1194_v8).Cardinality()))).Uint32()).(_dafny.Int)
			})()).Cmp((func() _dafny.Int {
				if _1186_v0 {
					return _1188_v2
				}
				return _1188_v2
			})()) >= 0)) {
				{
					if (_1195_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1195_i0 = (_1195_i0).Plus(_dafny.One)
					var _1196_v9 _dafny.Sequence
					_ = _1196_v9
					_1196_v9 = _dafny.UnicodeSeqOfUtf8Bytes("vkph")
					var _1197_v10 _dafny.CodePoint
					_ = _1197_v10
					_1197_v10 = _dafny.CodePoint('a')
					var _1198_v11 _dafny.Sequence
					_ = _1198_v11
					_1198_v11 = _dafny.SeqOf(_1193_v7, _1193_v7, _1193_v7)
					var _rhs269 _dafny.MultiSet = (_1198_v11).Select((Companion_Default___.SafeIndex(_1188_v2, _dafny.IntOfUint32((_1198_v11).Cardinality()))).Uint32()).(_dafny.MultiSet)
					_ = _rhs269
					var _rhs270 _dafny.Sequence = (func() _dafny.Sequence {
						if _1186_v0 {
							return _1196_v9
						}
						return _1196_v9
					})()
					_ = _rhs270
					var _rhs271 _dafny.CodePoint = _1197_v10
					_ = _rhs271
					var _rhs272 _dafny.Int = _dafny.IntOfUint32((_1196_v9).Cardinality())
					_ = _rhs272
					_1193_v7 = _rhs269
					_1196_v9 = _rhs270
					_1197_v10 = _rhs271
					_1188_v2 = _rhs272
					var _1199_v12 _dafny.Map
					_ = _1199_v12
					_1199_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1186_v0, _dafny.SeqOf(_1188_v2, _1188_v2))
					var _1200_v13 _dafny.Map
					_ = _1200_v13
					_1200_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
						if (_1199_v12).Contains(_1186_v0) {
							return (_1199_v12).Get(_1186_v0).(_dafny.Sequence)
						}
						return _1194_v8
					})(), _1194_v8), _1186_v0)
					var _1201_v14 _dafny.MultiSet
					_ = _1201_v14
					_1201_v14 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-69))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg75 _dafny.Int) interface{} {
							return coer75(arg75)
						}
					}((func(_1202_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1203_i1 _dafny.Int) _dafny.CodePoint {
							return _1202_v10
						}
					})(_1197_v10)))).Cardinality()), _1188_v2)
					if (func() bool {
						if (_1200_v13).Contains(_1194_v8) {
							return (_1200_v13).Get(_1194_v8).(bool)
						}
						return (_1201_v14).IsSubsetOf(_1201_v14)
					})() {
						var _rhs273 _dafny.Int = Companion_Default___.SafeModuloInt((_1188_v2).Times(_1188_v2), _1188_v2)
						_ = _rhs273
						var _rhs274 _dafny.Int = (_1194_v8).Select((Companion_Default___.SafeIndex(_1188_v2, _dafny.IntOfUint32((_1194_v8).Cardinality()))).Uint32()).(_dafny.Int)
						_ = _rhs274
						var _lhs202 *GlobalState = globalState
						_ = _lhs202
						_lhs202.F7 = _rhs273
						_1188_v2 = _rhs274
						_1190_v4 = (func() _dafny.Array {
							if _1186_v0 {
								return _1190_v4
							}
							return _1190_v4
						})()
						(globalState).F4 = (_dafny.Zero).Minus(_1188_v2)
						(globalState).F15 = (_dafny.Zero).Minus(_1188_v2)
						(globalState).F4 = (_this).Fm2(_1188_v2, globalState)
					} else {
						var _1204_v15 _dafny.Map
						_ = _1204_v15
						_1204_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1186_v0, _1186_v0)
						_1186_v0 = (func() bool {
							if (_1204_v15).Contains(false) {
								return (_1204_v15).Get(false).(bool)
							}
							return _1186_v0
						})()
						var _1205_v16 _dafny.Sequence
						_ = _1205_v16
						_1205_v16 = _dafny.SeqOf(_1186_v0, _1186_v0)
						var _1206_v17 _dafny.Map
						_ = _1206_v17
						_1206_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1186_v0, _1204_v15)
						var _1207_v18 _dafny.Array
						_ = _1207_v18
						var _nwElement0_39 bool = _1186_v0
						_ = _nwElement0_39
						var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(27))
						_ = _nw169
						(_nw169).ArraySet1(_nwElement0_39, 0)
						(_nw169).ArraySet1(_1186_v0, 1)
						(_nw169).ArraySet1(false, 2)
						(_nw169).ArraySet1(_1186_v0, 3)
						(_nw169).ArraySet1(_1186_v0, 4)
						(_nw169).ArraySet1(_1186_v0, 5)
						(_nw169).ArraySet1(!(false), 6)
						(_nw169).ArraySet1(_1186_v0, 7)
						(_nw169).ArraySet1(_1186_v0, 8)
						(_nw169).ArraySet1(_1186_v0, 9)
						(_nw169).ArraySet1(!(_1186_v0), 10)
						(_nw169).ArraySet1(_1186_v0, 11)
						(_nw169).ArraySet1((_1205_v16).Select((Companion_Default___.SafeIndex(_1188_v2, _dafny.IntOfUint32((_1205_v16).Cardinality()))).Uint32()).(bool), 12)
						(_nw169).ArraySet1(_1186_v0, 13)
						(_nw169).ArraySet1(true, 14)
						(_nw169).ArraySet1(_1186_v0, 15)
						(_nw169).ArraySet1(_1186_v0, 16)
						(_nw169).ArraySet1(_1186_v0, 17)
						(_nw169).ArraySet1((_this).Fm1(_1206_v17, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1188_v2, _1188_v2)).Cardinality(), globalState), 18)
						(_nw169).ArraySet1(_1186_v0, 19)
						(_nw169).ArraySet1(_1186_v0, 20)
						(_nw169).ArraySet1((_1205_v16).Select((Companion_Default___.SafeIndex(_1188_v2, _dafny.IntOfUint32((_1205_v16).Cardinality()))).Uint32()).(bool), 21)
						(_nw169).ArraySet1(_1186_v0, 22)
						(_nw169).ArraySet1(_1186_v0, 23)
						(_nw169).ArraySet1(_1186_v0, 24)
						(_nw169).ArraySet1(true, 25)
						(_nw169).ArraySet1((func() bool {
							if (_1189_v3).Contains((_dafny.SetOf(_1188_v2, _1188_v2, _1188_v2)).Cardinality()) {
								return (_1189_v3).Get((_dafny.SetOf(_1188_v2, _1188_v2, _1188_v2)).Cardinality()).(bool)
							}
							return _1186_v0
						})(), 26)
						_1207_v18 = _nw169
						var _1208_v19 _dafny.Map
						_ = _1208_v19
						_1208_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1207_v18, (func() bool {
							if _1186_v0 {
								return !(!(false))
							}
							return _1186_v0
						})())
						_1186_v0 = (func() bool {
							if (_1208_v19).Contains(_1207_v18) {
								return (_1208_v19).Get(_1207_v18).(bool)
							}
							return false
						})()
						var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))
						_ = _index220
						(_1207_v18).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1188_v2)).Contains((_this).Fm1(_1206_v17, _1188_v2, globalState)), (_index220).Int())
						var _1209_v20 _dafny.Set
						_ = _1209_v20
						_1209_v20 = _dafny.SetOf(_1186_v0)
						var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))
						_ = _index221
						(_1207_v18).ArraySet1((_1209_v20).IsDisjointFrom(_1209_v20), (_index221).Int())
						var _1210_v21 _dafny.Sequence
						_ = _1210_v21
						_1210_v21 = _dafny.SeqOf(_dafny.SeqOf(_1188_v2), _1194_v8)
						var _rhs275 bool = _1186_v0
						_ = _rhs275
						var _rhs276 _dafny.Sequence = _1210_v21
						_ = _rhs276
						var _rhs277 _dafny.Int = _1188_v2
						_ = _rhs277
						var _rhs278 _dafny.Map = (_1206_v17).Merge((_1206_v17).Merge(_1206_v17))
						_ = _rhs278
						_1186_v0 = _rhs275
						_1210_v21 = _rhs276
						_1188_v2 = _rhs277
						_1206_v17 = _rhs278
						var _1211_v22 D3
						_ = _1211_v22
						_1211_v22 = Companion_D3_.Create_DC9_(_1197_v10)
						var _1212_v23 _dafny.Sequence
						_ = _1212_v23
						_1212_v23 = _dafny.SeqOf(_1207_v18)
						var _1213_v24 _dafny.Array
						_ = _1213_v24
						var _nwElement0_40 _dafny.Array = _1207_v18
						_ = _nwElement0_40
						var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(7))
						_ = _nw170
						(_nw170).ArraySet1(_nwElement0_40, 0)
						(_nw170).ArraySet1(_1207_v18, 1)
						(_nw170).ArraySet1(_1207_v18, 2)
						(_nw170).ArraySet1((_1212_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.IntOfUint32((_1212_v23).Cardinality()))).Uint32()).(_dafny.Array), 3)
						(_nw170).ArraySet1(_1207_v18, 4)
						(_nw170).ArraySet1(_1207_v18, 5)
						(_nw170).ArraySet1(_1207_v18, 6)
						_1213_v24 = _nw170
						var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_1213_v24), 0))
						_ = _index222
						(_1213_v24).ArraySet1(_1207_v18, (_index222).Int())
						var _1214_v25 _dafny.Map
						_ = _1214_v25
						_1214_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_1186_v0, (_1207_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))).Int()).(bool), (_1207_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))).Int()).(bool))).Cardinality()), _1190_v4)
						var _1215_v26 _dafny.Set
						_ = _1215_v26
						_1215_v26 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ejjgk"))
						var _1216_v28 _dafny.Map
						_ = _1216_v28
						_1216_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
							var _coll52 = _dafny.NewMapBuilder()
							_ = _coll52
							for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(803), _dafny.IntOfInt64(-816))); ; {
								_compr_52, _ok53 := _iter53()
								if !_ok53 {
									break
								}
								var _1217_v27 _dafny.Int
								_1217_v27 = interface{}(_compr_52).(_dafny.Int)
								if ((_dafny.IntOfInt64(803)).Cmp(_1217_v27) <= 0) && ((_1217_v27).Cmp(_dafny.IntOfInt64(-816)) < 0) {
									_coll52.Add(Companion_Default___.SafeModuloInt(_1217_v27, _1188_v2), _1186_v0)
								}
							}
							return _coll52.ToMap()
						}(), !(!((_1207_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))).Int()).(bool))))
						var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))
						_ = _index223
						var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_1213_v24), 0))
						_ = _index224
						var _rhs279 bool = (((_dafny.MultiSetOf(_1186_v0)).Cardinality()).Minus((_1214_v25).Cardinality())).Cmp((_dafny.IntOfInt64(301)).Minus(_1188_v2)) == 0
						_ = _rhs279
						var _rhs280 D3 = Companion_Default___.Fm14((_1201_v14).Cardinality(), _1215_v26, _1188_v2, (_1216_v28).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1189_v3, (_1207_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))).Int()).(bool))).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1205_v16).Cardinality()), (_1207_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))).Int()).(bool))).Update((_dafny.SetOf(_dafny.IntOfInt64(946), _1188_v2)).Cardinality(), _1186_v0), (_1207_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))).Int()).(bool))), globalState)
						_ = _rhs280
						var _rhs281 _dafny.Array = _1207_v18
						_ = _rhs281
						var _lhs203 _dafny.Array = _1207_v18
						_ = _lhs203
						var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1207_v18), 0))
						_ = _lhs204
						var _lhs205 _dafny.Array = _1213_v24
						_ = _lhs205
						var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_1213_v24), 0))
						_ = _lhs206
						(_lhs203).ArraySet1(_rhs279, (_lhs204).Int())
						_1211_v22 = _rhs280
						(_lhs205).ArraySet1(_rhs281, (_lhs206).Int())
					}
					var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1190_v4), 0))
					_ = _index225
					(_1190_v4).ArraySet1(((_1189_v3).Update(_1188_v2, !(_1186_v0))).Cardinality(), (_index225).Int())
					var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1190_v4), 0))
					_ = _index226
					(_1190_v4).ArraySet1((func() _dafny.Int {
						if _1186_v0 {
							return (_dafny.Zero).Minus(_1188_v2)
						}
						return _1188_v2
					})(), (_index226).Int())
					var _1218_v29 D1
					_ = _1218_v29
					_1218_v29 = Companion_D1_.Create_DC2_(_1188_v2)
					if ((_1190_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1190_v4), 0))).Int()).(_dafny.Int)).Cmp((_1218_v29).Dtor_cf2()) > 0 {
						(globalState).F13 = (_1190_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1190_v4), 0))).Int()).(_dafny.Int)
						var _1219_v30 *C0
						_ = _1219_v30
						var _nw171 *C0 = New_C0_()
						_ = _nw171
						_nw171.Ctor__()
						_1219_v30 = _nw171
						var _1220_v31 *C2
						_ = _1220_v31
						var _nw172 *C2 = New_C2_()
						_ = _nw172
						_nw172.Ctor__()
						_1220_v31 = _nw172
						var _1221_v32 _dafny.Sequence
						_ = _1221_v32
						_1221_v32 = _dafny.SeqOf(_1186_v0)
						var _1222_v33 D11
						_ = _1222_v33
						_1222_v33 = Companion_D11_.Create_DC30_(_1186_v0, (_1190_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1190_v4), 0))).Int()).(_dafny.Int))
						var _1223_v34 _dafny.Map
						_ = _1223_v34
						_1223_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1221_v32, _1222_v33)
						var _1224_v35 _dafny.Map
						_ = _1224_v35
						_1224_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, _1186_v0), Companion_D11_.Create_DC30_(_1186_v0, _dafny.IntOfInt64(364)))).Merge(_1223_v34), _1186_v0)
						_1224_v35 = (_1224_v35).Update((_1223_v34).Merge(_1223_v34), _1186_v0)
						(globalState).F13 = (_1190_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1190_v4), 0))).Int()).(_dafny.Int)
					} else {
						var _1225_v36 D4
						_ = _1225_v36
						_1225_v36 = Companion_D4_.Create_DC12_((_1190_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1190_v4), 0))).Int()).(_dafny.Int))
						var _pat_let_tv18 = _1190_v4
						_ = _pat_let_tv18
						var _pat_let_tv19 = _1190_v4
						_ = _pat_let_tv19
						var _1226_v37 _dafny.Map
						_ = _1226_v37
						_1226_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let16_0 D4) D4 {
							return func(_1227_dt__update__tmp_h0 D4) D4 {
								return func(_pat_let17_0 _dafny.Int) D4 {
									return func(_1228_dt__update_hcf16_h0 _dafny.Int) D4 {
										return Companion_D4_.Create_DC12_(_1228_dt__update_hcf16_h0)
									}(_pat_let17_0)
								}((_pat_let_tv19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_pat_let_tv18), 0))).Int()).(_dafny.Int))
							}(_pat_let16_0)
						}(_1225_v36), (_1190_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1190_v4), 0))).Int()).(_dafny.Int))
						var _rhs282 _dafny.Map = (_1226_v37).Merge(_1226_v37)
						_ = _rhs282
						var _rhs283 _dafny.Int = _1188_v2
						_ = _rhs283
						var _rhs284 _dafny.Int = _dafny.IntOfInt64(660)
						_ = _rhs284
						var _lhs207 *GlobalState = globalState
						_ = _lhs207
						var _lhs208 *GlobalState = globalState
						_ = _lhs208
						_1226_v37 = _rhs282
						_lhs207.F13 = _rhs283
						_lhs208.F4 = _rhs284
						var _1229_v38 _dafny.Set
						_ = _1229_v38
						_1229_v38 = _dafny.SetOf(!(_1186_v0), !(false))
						var _1230_v39 D12
						_ = _1230_v39
						_1230_v39 = Companion_D12_.Create_DC32_((func() _dafny.Set {
							if true {
								return _1229_v38
							}
							return _1229_v38
						})())
						_1230_v39 = Companion_D12_.Create_DC32_(_1229_v38)
						(globalState).F7 = _dafny.IntOfInt64(641)
						var _1231_v40 _dafny.Array
						_ = _1231_v40
						var _len0_31 _dafny.Int = _dafny.IntOfInt64(15)
						_ = _len0_31
						var _nw173 _dafny.Array
						_ = _nw173
						if _len0_31.Cmp(_dafny.Zero) == 0 {
							_nw173 = _dafny.NewArray(_len0_31)
						} else {
							var _init31 func(_dafny.Int) _dafny.Sequence = (func(_1232_v0 bool) func(_dafny.Int) _dafny.Sequence {
								return func(_1233_i2 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqOf(_1232_v0, _1232_v0)
								}
							})(_1186_v0)
							_ = _init31
							var _element0_31 = _init31(_dafny.Zero)
							_ = _element0_31
							_nw173 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
							(_nw173).ArraySet1(_element0_31, 0)
							var _nativeLen0_31 = (_len0_31).Int()
							_ = _nativeLen0_31
							for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
								(_nw173).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
							}
						}
						_1231_v40 = _nw173
						var _1234_v41 _dafny.Array
						_ = _1234_v41
						_1234_v41 = _1231_v40
						var _1235_v42 _dafny.Map
						_ = _1235_v42
						_1235_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1201_v14).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_1196_v9).Cardinality()))), _1234_v41)
						_1235_v42 = (_1235_v42).Update(_1201_v14, _1234_v41)
						var _1236_v43 _dafny.Array
						_ = _1236_v43
						var _nw174 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
						_ = _nw174
						_1236_v43 = _nw174
						var _1237_v44 _dafny.Sequence
						_ = _1237_v44
						_1237_v44 = _dafny.SeqOf(_1236_v43, _1236_v43, _1236_v43, _1236_v43)
						var _1238_v45 D7
						_ = _1238_v45
						_1238_v45 = Companion_D7_.Create_DC19_(_1237_v44)
						_1238_v45 = _1238_v45
					}
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1239_v46 _dafny.CodePoint
		_ = _1239_v46
		_1239_v46 = _dafny.CodePoint('k')
		var _1240_v47 _dafny.Array
		_ = _1240_v47
		var _nwElement0_41 bool = true
		_ = _nwElement0_41
		var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(24))
		_ = _nw175
		(_nw175).ArraySet1(_nwElement0_41, 0)
		(_nw175).ArraySet1((_1186_v0) && (true), 1)
		(_nw175).ArraySet1(_1186_v0, 2)
		(_nw175).ArraySet1(_1186_v0, 3)
		(_nw175).ArraySet1((_1188_v2).Cmp(_1188_v2) != 0, 4)
		(_nw175).ArraySet1(_1186_v0, 5)
		(_nw175).ArraySet1(_1186_v0, 6)
		(_nw175).ArraySet1(!(false), 7)
		(_nw175).ArraySet1(_1186_v0, 8)
		(_nw175).ArraySet1(_1186_v0, 9)
		(_nw175).ArraySet1(!_dafny.Companion_Sequence_.Contains(p0, _1239_v46), 10)
		(_nw175).ArraySet1(_1186_v0, 11)
		(_nw175).ArraySet1(_1186_v0, 12)
		(_nw175).ArraySet1((_1186_v0) == (_1186_v0), 13)
		(_nw175).ArraySet1(_1186_v0, 14)
		(_nw175).ArraySet1(_1186_v0, 15)
		(_nw175).ArraySet1(_1186_v0, 16)
		(_nw175).ArraySet1(_1186_v0, 17)
		(_nw175).ArraySet1((func() bool {
			if (_1189_v3).Contains(_dafny.IntOfInt64(-201)) {
				return (_1189_v3).Get(_dafny.IntOfInt64(-201)).(bool)
			}
			return !(_1186_v0)
		})(), 18)
		(_nw175).ArraySet1((_dafny.IntOfInt64(55)).Cmp(_1188_v2) == 0, 19)
		(_nw175).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1194_v8, _1194_v8), 20)
		(_nw175).ArraySet1(Companion_Default___.Fm26(globalState), 21)
		(_nw175).ArraySet1((_this).Fm12(globalState), 22)
		(_nw175).ArraySet1(_1186_v0, 23)
		_1240_v47 = _nw175
		var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1240_v47), 0))
		_ = _index227
		(_1240_v47).ArraySet1(!(!(false)), (_index227).Int())
		var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1240_v47), 0))
		_ = _index228
		(_1240_v47).ArraySet1(_1186_v0, (_index228).Int())
		var _1241_v48 _dafny.Sequence
		_ = _1241_v48
		_1241_v48 = _dafny.SeqOf(_1186_v0)
		var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_1240_v47), 0))
		_ = _index229
		(_1240_v47).ArraySet1((_1188_v2).Cmp(_dafny.IntOfUint32((_1241_v48).Cardinality())) == 0, (_index229).Int())
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C8{}
var _ T1 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) Ctor__() {
	{
	}
}
func (_this *C8) Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(326))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg76 _dafny.Int) interface{} {
				return coer76(arg76)
			}
		}(func(_1242_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
		})), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(_dafny.IntOfInt64(404))).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hwtfiavx")).Cardinality()), (_dafny.MultiSetOf(true)).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(766), _dafny.IntOfInt64(187), _dafny.IntOfInt64(520)), _dafny.SeqOf(_dafny.IntOfInt64(474))))
	}
}
func (_this *C8) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		if !((true) || (true)) {
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("bmw"), !(false))).Dtor_cf3(), _dafny.UnicodeSeqOfUtf8Bytes("yvx"))).Cardinality()))
		} else {
			return _dafny.IntOfInt64(231)
		}
	}
}
func (_this *C8) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf(func() _dafny.Set {
			var _coll53 = _dafny.NewBuilder()
			_ = _coll53
			for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-228), _dafny.IntOfInt64(300))); ; {
				_compr_53, _ok54 := _iter54()
				if !_ok54 {
					break
				}
				var _1243_v0 _dafny.Int
				_1243_v0 = interface{}(_compr_53).(_dafny.Int)
				if ((_dafny.IntOfInt64(-228)).Cmp(_1243_v0) <= 0) && ((_1243_v0).Cmp(_dafny.IntOfInt64(300)) < 0) {
					_coll53.Add((_1243_v0).Minus((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(255))).Cardinality(), _dafny.IntOfInt64(129))).Cardinality()))
				}
			}
			return _coll53.ToSet()
		}())).Union((_dafny.MultiSetOf(func() _dafny.Set {
			var _coll54 = _dafny.NewBuilder()
			_ = _coll54
			for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-367), _dafny.IntOfInt64(626))); ; {
				_compr_54, _ok55 := _iter55()
				if !_ok55 {
					break
				}
				var _1244_v1 _dafny.Int
				_1244_v1 = interface{}(_compr_54).(_dafny.Int)
				if ((_dafny.IntOfInt64(-367)).Cmp(_1244_v1) <= 0) && ((_1244_v1).Cmp(_dafny.IntOfInt64(626)) < 0) {
					_coll54.Add((_1244_v1).Times(_dafny.IntOfInt64(122)))
				}
			}
			return _coll54.ToSet()
		}(), func() _dafny.Set {
			var _coll55 = _dafny.NewBuilder()
			_ = _coll55
			for _iter56 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), (_dafny.SetOf((_dafny.SetOf(_dafny.CodePoint('w'))).Cardinality())).Cardinality())).Keys().Elements()); ; {
				_compr_55, _ok56 := _iter56()
				if !_ok56 {
					break
				}
				var _1245_v2 _dafny.Int
				_1245_v2 = interface{}(_compr_55).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), (_dafny.SetOf((_dafny.SetOf(_dafny.CodePoint('w'))).Cardinality())).Cardinality())).Contains(_1245_v2) {
					_coll55.Add(Companion_Default___.SafeModuloInt(_1245_v2, _dafny.IntOfInt64(477)))
				}
			}
			return _coll55.ToSet()
		}())).Difference(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality())), _dafny.SetOf((_dafny.MultiSetOf(_dafny.SeqOf(!(false), true))).Cardinality(), _dafny.IntOfInt64(622)), _dafny.SetOf(_dafny.IntOfInt64(18)), func() _dafny.Set {
			var _coll56 = _dafny.NewBuilder()
			_ = _coll56
			for _iter57 := _dafny.Iterate((func() _dafny.Map {
				var _coll57 = _dafny.NewMapBuilder()
				_ = _coll57
				for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(723), _dafny.IntOfInt64(-208))); ; {
					_compr_57, _ok58 := _iter58()
					if !_ok58 {
						break
					}
					var _1246_v3 _dafny.Int
					_1246_v3 = interface{}(_compr_57).(_dafny.Int)
					if ((_dafny.IntOfInt64(723)).Cmp(_1246_v3) <= 0) && ((_1246_v3).Cmp(_dafny.IntOfInt64(-208)) < 0) {
						_coll57.Add((_1246_v3).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(273))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg77 _dafny.Int) interface{} {
								return coer77(arg77)
							}
						}(func(_1248_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('a')
						}))).Cardinality())), (func() _dafny.Map {
							var _coll58 = _dafny.NewMapBuilder()
							_ = _coll58
							for _iter59 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xypuautvo"))).Elements()); ; {
								_compr_58, _ok59 := _iter59()
								if !_ok59 {
									break
								}
								var _1247_v4 _dafny.Sequence
								_1247_v4 = interface{}(_compr_58).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xypuautvo")), _1247_v4) {
									_coll58.Add(_1247_v4, false)
								}
							}
							return _coll58.ToMap()
						}()).Cardinality())
					}
				}
				return _coll57.ToMap()
			}()).Keys().Elements()); ; {
				_compr_56, _ok57 := _iter57()
				if !_ok57 {
					break
				}
				var _1249_v5 _dafny.Int
				_1249_v5 = interface{}(_compr_56).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(723), _dafny.IntOfInt64(-208))); ; {
						_compr_59, _ok60 := _iter60()
						if !_ok60 {
							break
						}
						var _1246_v3 _dafny.Int
						_1246_v3 = interface{}(_compr_59).(_dafny.Int)
						if ((_dafny.IntOfInt64(723)).Cmp(_1246_v3) <= 0) && ((_1246_v3).Cmp(_dafny.IntOfInt64(-208)) < 0) {
							_coll59.Add((_1246_v3).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(273))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg78 _dafny.Int) interface{} {
									return coer78(arg78)
								}
							}(func(_1248_i0 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('a')
							}))).Cardinality())), (func() _dafny.Map {
								var _coll60 = _dafny.NewMapBuilder()
								_ = _coll60
								for _iter61 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xypuautvo"))).Elements()); ; {
									_compr_60, _ok61 := _iter61()
									if !_ok61 {
										break
									}
									var _1247_v4 _dafny.Sequence
									_1247_v4 = interface{}(_compr_60).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xypuautvo")), _1247_v4) {
										_coll60.Add(_1247_v4, false)
									}
								}
								return _coll60.ToMap()
							}()).Cardinality())
						}
					}
					return _coll59.ToMap()
				}()).Contains(_1249_v5) {
					_coll56.Add((_1249_v5).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ubinwfp")).Cardinality())))
				}
			}
			return _coll56.ToSet()
		}(), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yhq")).Cardinality())))))
	}
}
func (_this *C8) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		if (_dafny.MultiSetOf(_dafny.SeqOf(true), _dafny.SeqOf(false))).IsSubsetOf(_dafny.MultiSetOf(_dafny.SeqOf(true))) {
			return true
		} else {
			return (_dafny.IntOfInt64(909)).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ydocfi")).Cardinality())) > 0
		}
	}
}
func (_this *C8) M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _1250_v0 bool
		_ = _1250_v0
		_1250_v0 = true
		_1250_v0 = _1250_v0
		var _1251_v1 _dafny.Array
		_ = _1251_v1
		var _len0_32 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_32
		var _nw176 _dafny.Array
		_ = _nw176
		if _len0_32.Cmp(_dafny.Zero) == 0 {
			_nw176 = _dafny.NewArray(_len0_32)
		} else {
			var _init32 func(_dafny.Int) bool = (func(_1252_v0 bool) func(_dafny.Int) bool {
				return func(_1253_i1 _dafny.Int) bool {
					return _1252_v0
				}
			})(_1250_v0)
			_ = _init32
			var _element0_32 = _init32(_dafny.Zero)
			_ = _element0_32
			_nw176 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
			(_nw176).ArraySet1(_element0_32, 0)
			var _nativeLen0_32 = (_len0_32).Int()
			_ = _nativeLen0_32
			for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
				(_nw176).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
			}
		}
		_1251_v1 = _nw176
		for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1251_v1), 0))); ; {
			_guard_loop_1, _ok62 := _iter62()
			if !_ok62 {
				break
			}
			var _1254_i0 _dafny.Int
			_1254_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_1254_i0).Sign() != -1) && ((_1254_i0).Cmp(_dafny.ArrayLen((_1251_v1), 0)) < 0)) {
				(_1251_v1).ArraySet1(!(_1250_v0) || (!(!(_1250_v0))), (_1254_i0).Int())
			}
		}
		var _1255_v2 _dafny.Int
		_ = _1255_v2
		_1255_v2 = _dafny.IntOfInt64(551)
		var _1256_v4 _dafny.Map
		_ = _1256_v4
		_1256_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1255_v2, !(_1250_v0))
		var _hi8 _dafny.Int = (func() _dafny.Int {
			if _1250_v0 {
				return _1255_v2
			}
			return (func() _dafny.Map {
				var _coll61 = _dafny.NewMapBuilder()
				_ = _coll61
				for _iter63 := _dafny.Iterate((_1256_v4).Keys().Elements()); ; {
					_compr_61, _ok63 := _iter63()
					if !_ok63 {
						break
					}
					var _1257_v3 _dafny.Int
					_1257_v3 = interface{}(_compr_61).(_dafny.Int)
					if (_1256_v4).Contains(_1257_v3) {
						_coll61.Add(Companion_Default___.SafeModuloInt(_1257_v3, _1255_v2), _1255_v2)
					}
				}
				return _coll61.ToMap()
			}()).Cardinality()
		})()
		_ = _hi8
		for _1258_i2 := (Companion_Default___.Fm11(_1250_v0, globalState)).Cardinality(); _1258_i2.Cmp(_hi8) < 0; _1258_i2 = _1258_i2.Plus(_dafny.One) {
			var _1259_v5 _dafny.Int
			_ = _1259_v5
			var _1260_v6 bool
			_ = _1260_v6
			var _out9 _dafny.Int
			_ = _out9
			var _out10 bool
			_ = _out10
			_out9, _out10 = (_this).M8(globalState)
			_1259_v5 = _out9
			_1260_v6 = _out10
			var _1261_v7 _dafny.Sequence
			_ = _1261_v7
			_1261_v7 = _dafny.SeqOf(_1250_v0, _1250_v0)
			_1261_v7 = _dafny.Companion_Sequence_.Concatenate(_1261_v7, _1261_v7)
			_1250_v0 = (_1250_v0) && (_1250_v0)
			(globalState).F4 = _1259_v5
		}
		var _1262_v8 *C5
		_ = _1262_v8
		var _nw177 *C5 = New_C5_()
		_ = _nw177
		_nw177.Ctor__()
		_1262_v8 = _nw177
		_1251_v1 = _1251_v1
		var _1263_v9 _dafny.Map
		_ = _1263_v9
		_1263_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1251_v1, _1250_v0)
		var _1264_i3 _dafny.Int
		_ = _1264_i3
		_1264_i3 = _dafny.Zero
		{
			for (func() bool {
				if (_1263_v9).Contains(_1251_v1) {
					return (_1263_v9).Get(_1251_v1).(bool)
				}
				return _1250_v0
			})() {
				{
					if (_1264_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1264_i3 = (_1264_i3).Plus(_dafny.One)
					var _1265_v10 _dafny.Array
					_ = _1265_v10
					var _nw178 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
					_ = _nw178
					_1265_v10 = _nw178
					var _1266_v11 _dafny.MultiSet
					_ = _1266_v11
					_1266_v11 = _dafny.MultiSetOf(_1265_v10)
					var _1267_v12 _dafny.Map
					_ = _1267_v12
					_1267_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1266_v11).IsSubsetOf((_1266_v11).Update(_1265_v10, Companion_Default___.Abs(_1255_v2))), (_1250_v0) == (_1250_v0))
					_1267_v12 = (_1267_v12).Update(_1250_v0, !(_1250_v0))
					var _1268_v13 D7
					_ = _1268_v13
					_1268_v13 = Companion_D7_.Create_DC21_(_dafny.IntOfInt64(637), _dafny.IntOfInt64(-530), _1250_v0, _1250_v0, _1250_v0)
					(globalState).F4 = (_1268_v13).Dtor_cf31()
					var _1269_v14 _dafny.Array
					_ = _1269_v14
					var _len0_33 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_33
					var _nw179 _dafny.Array
					_ = _nw179
					if _len0_33.Cmp(_dafny.Zero) == 0 {
						_nw179 = _dafny.NewArray(_len0_33)
					} else {
						var _init33 func(_dafny.Int) _dafny.MultiSet = (func(_1270_v0 bool) func(_dafny.Int) _dafny.MultiSet {
							return func(_1271_i4 _dafny.Int) _dafny.MultiSet {
								return _dafny.MultiSetOf(_1270_v0)
							}
						})(_1250_v0)
						_ = _init33
						var _element0_33 = _init33(_dafny.Zero)
						_ = _element0_33
						_nw179 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
						(_nw179).ArraySet1(_element0_33, 0)
						var _nativeLen0_33 = (_len0_33).Int()
						_ = _nativeLen0_33
						for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
							(_nw179).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
						}
					}
					_1269_v14 = _nw179
					var _1272_v15 _dafny.Map
					_ = _1272_v15
					_1272_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1269_v14, _1250_v0)
					if (func() bool {
						if (_1272_v15).Contains(_1269_v14) {
							return (_1272_v15).Get(_1269_v14).(bool)
						}
						return (func() bool {
							if _1250_v0 {
								return _1250_v0
							}
							return !(_1250_v0)
						})()
					})() {
						var _1273_v16 D6
						_ = _1273_v16
						_1273_v16 = Companion_D6_.Create_DC16_(_dafny.SeqOf(_1255_v2, _1255_v2, _1255_v2))
						_1273_v16 = _1273_v16
						var _1274_v17 _dafny.MultiSet
						_ = _1274_v17
						_1274_v17 = _dafny.MultiSetOf((_dafny.Zero).Minus(_1255_v2), _1255_v2)
						var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1251_v1), 0))
						_ = _index230
						(_1251_v1).ArraySet1(((func() _dafny.Int {
							if (_1274_v17).Contains(_1255_v2) {
								return (_1274_v17).Multiplicity(_1255_v2)
							}
							return _1255_v2
						})()).Cmp(_1255_v2) > 0, (_index230).Int())
						var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1251_v1), 0))
						_ = _index231
						(_1251_v1).ArraySet1((func() bool {
							if (_1256_v4).Contains(Companion_Default___.SafeModuloInt(_1255_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gmlde")).Cardinality()))) {
								return (_1256_v4).Get(Companion_Default___.SafeModuloInt(_1255_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gmlde")).Cardinality()))).(bool)
							}
							return _1250_v0
						})(), (_index231).Int())
						_1250_v0 = (_1251_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1251_v1), 0))).Int()).(bool)
						var _1275_v18 *C1
						_ = _1275_v18
						var _nw180 *C1 = New_C1_()
						_ = _nw180
						_nw180.Ctor__()
						_1275_v18 = _nw180
						var _1276_v19 D15
						_ = _1276_v19
						_1276_v19 = Companion_D15_.Create_DC38_(_1256_v4)
						var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_1265_v10), 0))
						_ = _index232
						(_1265_v10).ArraySet1((_dafny.Zero).Minus(((_1276_v19).Dtor_cf55()).Cardinality()), (_index232).Int())
						var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1251_v1), 0))
						_ = _index233
						var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_1265_v10), 0))
						_ = _index234
						var _rhs285 bool = (_1251_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1251_v1), 0))).Int()).(bool)
						_ = _rhs285
						var _rhs286 _dafny.Int = _1255_v2
						_ = _rhs286
						var _rhs287 _dafny.Int = Companion_Default___.SafeDivisionInt(_1255_v2, _1255_v2)
						_ = _rhs287
						var _rhs288 bool = (_1251_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1251_v1), 0))).Int()).(bool)
						_ = _rhs288
						var _lhs209 _dafny.Array = _1251_v1
						_ = _lhs209
						var _lhs210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_1251_v1), 0))
						_ = _lhs210
						var _lhs211 *GlobalState = globalState
						_ = _lhs211
						var _lhs212 _dafny.Array = _1265_v10
						_ = _lhs212
						var _lhs213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_1265_v10), 0))
						_ = _lhs213
						(_lhs209).ArraySet1(_rhs285, (_lhs210).Int())
						_lhs211.F7 = _rhs286
						(_lhs212).ArraySet1(_rhs287, (_lhs213).Int())
						_1250_v0 = _rhs288
					} else {
						_1250_v0 = _1250_v0
						var _1277_v20 *C1
						_ = _1277_v20
						var _nw181 *C1 = New_C1_()
						_ = _nw181
						_nw181.Ctor__()
						_1277_v20 = _nw181
						var _1278_v21 _dafny.Sequence
						_ = _1278_v21
						_1278_v21 = _dafny.SeqOf(_1277_v20, _1277_v20, _1277_v20, _1277_v20, _1277_v20)
						_1277_v20 = (_1278_v21).Select((Companion_Default___.SafeIndex(_1255_v2, _dafny.IntOfUint32((_1278_v21).Cardinality()))).Uint32()).(*C1)
						_1255_v2 = _dafny.IntOfInt64(328)
						var _1279_v22 _dafny.Sequence
						_ = _1279_v22
						_1279_v22 = _dafny.UnicodeSeqOfUtf8Bytes("lqpqi")
						var _1280_v23 _dafny.Sequence
						_ = _1280_v23
						_1280_v23 = _dafny.SeqOf((_1255_v2).Plus(_1255_v2))
						var _1281_v24 _dafny.Array
						_ = _1281_v24
						var _nw182 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
						_ = _nw182
						_1281_v24 = _nw182
						var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1281_v24), 0))
						_ = _index235
						(_1281_v24).ArraySet1(_1251_v1, (_index235).Int())
						var _1282_v25 _dafny.Sequence
						_ = _1282_v25
						_1282_v25 = _dafny.SeqOf(_1250_v0)
						var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1281_v24), 0))
						_ = _index236
						var _rhs289 _dafny.Int = (_1277_v20).Fm25(_1250_v0, _1250_v0, globalState)
						_ = _rhs289
						var _rhs290 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(947))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg79 _dafny.Int) interface{} {
								return coer79(arg79)
							}
						}(func(_1283_i5 _dafny.Int) _dafny.CodePoint {
							return (Companion_D3_.Create_DC9_(_dafny.CodePoint('d'))).Dtor_cf12()
						}))
						_ = _rhs290
						var _rhs291 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1280_v23, _1280_v23)
						_ = _rhs291
						var _rhs292 _dafny.Int = Companion_Default___.SafeDivisionInt(_1255_v2, Companion_Default___.SafeModuloInt(_1255_v2, _dafny.IntOfUint32((_1282_v25).Cardinality())))
						_ = _rhs292
						var _rhs293 _dafny.Array = _1251_v1
						_ = _rhs293
						var _lhs214 *GlobalState = globalState
						_ = _lhs214
						var _lhs215 _dafny.Array = _1281_v24
						_ = _lhs215
						var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1281_v24), 0))
						_ = _lhs216
						_lhs214.F4 = _rhs289
						_1279_v22 = _rhs290
						_1280_v23 = _rhs291
						_1255_v2 = _rhs292
						(_lhs215).ArraySet1(_rhs293, (_lhs216).Int())
						var _1284_v26 _dafny.Map
						_ = _1284_v26
						_1284_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1250_v0, _1255_v2)
						_1250_v0 = (func() bool {
							if !(_1284_v26).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1250_v0, _1255_v2)) {
								return false
							}
							return (_1255_v2).Cmp((_dafny.Zero).Minus(_1255_v2)) != 0
						})()
					}
					var _1285_v27 _dafny.MultiSet
					_ = _1285_v27
					_1285_v27 = _dafny.MultiSetOf(_1250_v0)
					var _1286_v28 D2
					_ = _1286_v28
					_1286_v28 = Companion_D2_.Create_DC7_(_1250_v0, _1285_v27, _1265_v10)
					var _1287_v29 D2
					_ = _1287_v29
					_1287_v29 = Companion_D2_.Create_DC8_(_1286_v28)
					var _1288_v30 *C3
					_ = _1288_v30
					var _nw183 *C3 = New_C3_()
					_ = _nw183
					_nw183.Ctor__(_1287_v29, _1250_v0)
					_1288_v30 = _nw183
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
	}
}
func (_this *C8) M6(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 D2, globalState *GlobalState) {
	{
		var _1289_v0 _dafny.Map
		_ = _1289_v0
		_1289_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).Fm2(p0, globalState))
		var _1290_v1 _dafny.Sequence
		_ = _1290_v1
		_1290_v1 = _dafny.UnicodeSeqOfUtf8Bytes("bqid")
		var _1291_v2 D4
		_ = _1291_v2
		_1291_v2 = Companion_D4_.Create_DC11_(_dafny.SeqOf(_1290_v1, _1290_v1, _1290_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(321))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg80 _dafny.Int) interface{} {
				return coer80(arg80)
			}
		}(func(_1292_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})), _1290_v1))
		var _1293_v3 _dafny.Map
		_ = _1293_v3
		_1293_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v0, (func() D4 {
			if p2 {
				return _1291_v2
			}
			return _1291_v2
		})())
		var _1294_v4 _dafny.Sequence
		_ = _1294_v4
		_1294_v4 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(144))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}(func(_1295_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-780))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg82 _dafny.Int) interface{} {
				return coer82(arg82)
			}
		}(func(_1296_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})), _1290_v1, _dafny.UnicodeSeqOfUtf8Bytes("qfjnmala"))
		var _1297_v5 _dafny.Map
		_ = _1297_v5
		_1297_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), p2)
		var _1298_v6 _dafny.Sequence
		_ = _1298_v6
		_1298_v6 = _dafny.SeqOf((_1294_v4).Select((Companion_Default___.SafeIndex((_1297_v5).Cardinality(), _dafny.IntOfUint32((_1294_v4).Cardinality()))).Uint32()).(_dafny.Sequence), _1290_v1)
		_1293_v3 = (_1293_v3).Update(_1289_v0, Companion_D4_.Create_DC11_(_1294_v4))
		var _1299_i3 _dafny.Int
		_ = _1299_i3
		_1299_i3 = _dafny.Zero
		{
			for p2 {
				{
					if (_1299_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1299_i3 = (_1299_i3).Plus(_dafny.One)
					var _1300_v7 _dafny.Array
					_ = _1300_v7
					var _nw184 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
					_ = _nw184
					_1300_v7 = _nw184
					var _1301_v8 _dafny.Sequence
					_ = _1301_v8
					_1301_v8 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2))
					var _1302_v9 _dafny.Sequence
					_ = _1302_v9
					_1302_v9 = _dafny.SeqOf((Companion_D16_.Create_DC41_(_1301_v8)).Dtor_cf60())
					var _1303_v10 _dafny.Map
					_ = _1303_v10
					_1303_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1300_v7, _dafny.Companion_Sequence_.Update((_1302_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1302_v9).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_1302_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1302_v9).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)))
					var _1304_v11 _dafny.Map
					_ = _1304_v11
					_1304_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), p2)
					_1303_v10 = (_1303_v10).Update(_1300_v7, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2), _1304_v11, _1304_v11))
					var _1305_v12 bool
					_ = _1305_v12
					_1305_v12 = true
					_1305_v12 = (_1305_v12) == (_1305_v12)
					var _1306_v13 _dafny.Sequence
					_ = _1306_v13
					_1306_v13 = _dafny.SeqOf(p2)
					var _1307_v14 *C0
					_ = _1307_v14
					var _nw185 *C0 = New_C0_()
					_ = _nw185
					_nw185.Ctor__()
					_1307_v14 = _nw185
					var _1308_v15 _dafny.Set
					_ = _1308_v15
					_1308_v15 = _dafny.SetOf(_1307_v14)
					var _1309_v16 _dafny.Map
					_ = _1309_v16
					_1309_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1290_v1)
					var _1310_v17 _dafny.Array
					_ = _1310_v17
					var _nwElement0_42 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1298_v6).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_1298_v6).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality(), _dafny.IntOfUint32((_1306_v13).Cardinality()))
					_ = _nwElement0_42
					var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(24))
					_ = _nw186
					(_nw186).ArraySet1(_nwElement0_42, 0)
					(_nw186).ArraySet1(_dafny.IntOfInt64(961), 1)
					(_nw186).ArraySet1(p0, 2)
					(_nw186).ArraySet1((_dafny.IntOfInt64(-500)).Times(p0), 3)
					(_nw186).ArraySet1((_this).Fm2(_dafny.IntOfInt64(-428), globalState), 4)
					(_nw186).ArraySet1(p0, 5)
					(_nw186).ArraySet1((_this).Fm2(p0, globalState), 6)
					(_nw186).ArraySet1((p0).Minus(p0), 7)
					(_nw186).ArraySet1((_1308_v15).Cardinality(), 8)
					(_nw186).ArraySet1(p0, 9)
					(_nw186).ArraySet1(_dafny.IntOfInt64(-819), 10)
					(_nw186).ArraySet1(p0, 11)
					(_nw186).ArraySet1((p1).Cardinality(), 12)
					(_nw186).ArraySet1(p0, 13)
					(_nw186).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if _1305_v12 {
							return _1290_v1
						}
						return _1290_v1
					})()).Cardinality()), 14)
					(_nw186).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_1309_v16).Contains(_1305_v12) {
							return (_1309_v16).Get(_1305_v12).(_dafny.Sequence)
						}
						return _1290_v1
					})()).Cardinality()), 15)
					(_nw186).ArraySet1((_dafny.Zero).Minus(((p1).Intersection(p1)).Cardinality()), 16)
					(_nw186).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1306_v13, Companion_Default___.Fm27(globalState))).Cardinality()), 17)
					(_nw186).ArraySet1(p0, 18)
					(_nw186).ArraySet1((p0).Times(p0), 19)
					(_nw186).ArraySet1(p0, 20)
					(_nw186).ArraySet1((_dafny.Zero).Minus(p0), 21)
					(_nw186).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), 22)
					(_nw186).ArraySet1(p0, 23)
					_1310_v17 = _nw186
					var _1311_v18 D8
					_ = _1311_v18
					_1311_v18 = Companion_D8_.Create_DC25_(p0)
					var _pat_let_tv20 = p0
					_ = _pat_let_tv20
					var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1310_v17), 0))
					_ = _index237
					(_1310_v17).ArraySet1((func(_pat_let18_0 D8) D8 {
						return func(_1312_dt__update__tmp_h0 D8) D8 {
							return func(_pat_let19_0 _dafny.Int) D8 {
								return func(_1313_dt__update_hcf42_h0 _dafny.Int) D8 {
									return Companion_D8_.Create_DC25_(_1313_dt__update_hcf42_h0)
								}(_pat_let19_0)
							}(_pat_let_tv20)
						}(_pat_let18_0)
					}(_1311_v18)).Dtor_cf42(), (_index237).Int())
					var _1314_v19 *C5
					_ = _1314_v19
					var _nw187 *C5 = New_C5_()
					_ = _nw187
					_nw187.Ctor__()
					_1314_v19 = _nw187
					var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1310_v17), 0))
					_ = _index238
					var _rhs294 _dafny.Int = _dafny.IntOfInt64(215)
					_ = _rhs294
					var _rhs295 *C5 = _1314_v19
					_ = _rhs295
					var _lhs217 _dafny.Array = _1310_v17
					_ = _lhs217
					var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1310_v17), 0))
					_ = _lhs218
					(_lhs217).ArraySet1(_rhs294, (_lhs218).Int())
					_1314_v19 = _rhs295
					(globalState).F15 = (_1310_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1310_v17), 0))).Int()).(_dafny.Int)
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _hi9 _dafny.Int = p0
		_ = _hi9
		for _1315_i4 := p0; _1315_i4.Cmp(_hi9) < 0; _1315_i4 = _1315_i4.Plus(_dafny.One) {
			var _1316_v20 bool
			_ = _1316_v20
			_1316_v20 = false
			var _1317_v21 _dafny.Array
			_ = _1317_v21
			var _nwElement0_43 bool = p2
			_ = _nwElement0_43
			var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(15))
			_ = _nw188
			(_nw188).ArraySet1(_nwElement0_43, 0)
			(_nw188).ArraySet1(p2, 1)
			(_nw188).ArraySet1(p2, 2)
			(_nw188).ArraySet1(_1316_v20, 3)
			(_nw188).ArraySet1(_1316_v20, 4)
			(_nw188).ArraySet1(!(_1316_v20), 5)
			(_nw188).ArraySet1(Companion_Default___.Fm26(globalState), 6)
			(_nw188).ArraySet1(_1316_v20, 7)
			(_nw188).ArraySet1(_1316_v20, 8)
			(_nw188).ArraySet1(_1316_v20, 9)
			(_nw188).ArraySet1(_1316_v20, 10)
			(_nw188).ArraySet1(_1316_v20, 11)
			(_nw188).ArraySet1(!(_1316_v20), 12)
			(_nw188).ArraySet1(_1316_v20, 13)
			(_nw188).ArraySet1(false, 14)
			_1317_v21 = _nw188
			var _1318_v22 _dafny.Map
			_ = _1318_v22
			_1318_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1316_v20, _1317_v21)
			var _1319_v23 _dafny.Set
			_ = _1319_v23
			_1319_v23 = _dafny.SetOf(p2)
			var _1320_v24 D12
			_ = _1320_v24
			_1320_v24 = Companion_D12_.Create_DC32_(_1319_v23)
			var _1321_v25 _dafny.Map
			_ = _1321_v25
			_1321_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_1318_v22).Contains(_1316_v20) {
					return (_1318_v22).Get(_1316_v20).(_dafny.Array)
				}
				return _1317_v21
			})(), _1320_v24)
			_1316_v20 = !(_1321_v25).Equals(_1321_v25)
			var _1322_v26 D2
			_ = _1322_v26
			_1322_v26 = Companion_D2_.Create_DC6_(p0)
			var _1323_v27 *C3
			_ = _1323_v27
			var _nw189 *C3 = New_C3_()
			_ = _nw189
			_nw189.Ctor__(Companion_D2_.Create_DC8_(_1322_v26), p2)
			_1323_v27 = _nw189
			var _1324_v28 _dafny.Map
			_ = _1324_v28
			_1324_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1290_v1).Cardinality()), p0)
			var _1325_v29 _dafny.Map
			_ = _1325_v29
			_1325_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1324_v28).Cardinality(), _1315_i4)
			var _1326_v30 _dafny.Sequence
			_ = _1326_v30
			_1326_v30 = _dafny.SeqOf((_dafny.Zero).Minus(_1315_i4), Companion_Default___.Fm17((_1324_v28).Cardinality(), globalState))
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1317_v21), 0))
			_ = _index239
			(_1317_v21).ArraySet1((_this).Fm8(_dafny.IntOfUint32((_1326_v30).Cardinality()), !(_1323_v27.F25), _1315_i4, globalState), (_index239).Int())
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1317_v21), 0))
			_ = _index240
			(_1317_v21).ArraySet1(!(false), (_index240).Int())
			var _1327_v31 _dafny.Array
			_ = _1327_v31
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_34
			var _nw190 _dafny.Array
			_ = _nw190
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw190 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.Map = (func(_1328_v27 *C3) func(_dafny.Int) _dafny.Map {
					return func(_1329_i5 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1328_v27.F25), _dafny.IntOfInt64(851))
					}
				})(_1323_v27)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw190 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw190).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw190).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1327_v31 = _nw190
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1327_v31), 0))
			_ = _index241
			(_1327_v31).ArraySet1((_1289_v0), (_index241).Int())
			var _1330_v32 _dafny.CodePoint
			_ = _1330_v32
			_1330_v32 = _dafny.CodePoint('x')
			var _1331_v33 _dafny.Map
			_ = _1331_v33
			_1331_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1330_v32)
			var _1332_v34 _dafny.Array
			_ = _1332_v34
			var _nwElement0_44 _dafny.CodePoint = _dafny.CodePoint('p')
			_ = _nwElement0_44
			var _nw191 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(14))
			_ = _nw191
			(_nw191).ArraySet1CodePoint(_nwElement0_44, 0)
			(_nw191).ArraySet1CodePoint(_1330_v32, 1)
			(_nw191).ArraySet1CodePoint(_1330_v32, 2)
			(_nw191).ArraySet1CodePoint(_dafny.CodePoint('m'), 3)
			(_nw191).ArraySet1CodePoint(_1330_v32, 4)
			(_nw191).ArraySet1CodePoint(_1330_v32, 5)
			(_nw191).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_1331_v33).Contains(_1315_i4) {
					return (_1331_v33).Get(_1315_i4).(_dafny.CodePoint)
				}
				return _1330_v32
			})(), 6)
			(_nw191).ArraySet1CodePoint(_dafny.CodePoint('x'), 7)
			(_nw191).ArraySet1CodePoint(_dafny.CodePoint('h'), 8)
			(_nw191).ArraySet1CodePoint(_1330_v32, 9)
			(_nw191).ArraySet1CodePoint(_1330_v32, 10)
			(_nw191).ArraySet1CodePoint((func() _dafny.CodePoint {
				if !(p2) {
					return _1330_v32
				}
				return _1330_v32
			})(), 11)
			(_nw191).ArraySet1CodePoint(_1330_v32, 12)
			(_nw191).ArraySet1CodePoint(_1330_v32, 13)
			_1332_v34 = _nw191
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1332_v34), 0))
			_ = _index242
			(_1332_v34).ArraySet1CodePoint(_1330_v32, (_index242).Int())
			var _1333_v35 *C5
			_ = _1333_v35
			var _nw192 *C5 = New_C5_()
			_ = _nw192
			_nw192.Ctor__()
			_1333_v35 = _nw192
			var _1334_v36 _dafny.Map
			_ = _1334_v36
			_1334_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1315_i4)
			var _1335_v37 _dafny.Map
			_ = _1335_v37
			_1335_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1334_v36).Cardinality())
			var _1336_v38 _dafny.Sequence
			_ = _1336_v38
			_1336_v38 = _dafny.SeqOf(p1, p1, (p1).Update(p2, Companion_Default___.Abs(p0)))
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1327_v31), 0))
			_ = _index243
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1332_v34), 0))
			_ = _index244
			var _rhs296 bool = _dafny.Companion_Sequence_.Contains(_1290_v1, _1330_v32)
			_ = _rhs296
			var _rhs297 _dafny.Map = (_1334_v36).Update(!((_1317_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1317_v21), 0))).Int()).(bool)) || (true), p0)
			_ = _rhs297
			var _rhs298 _dafny.CodePoint = (_1290_v1).Select((Companion_Default___.SafeIndex(((_1336_v38).Select((Companion_Default___.SafeIndex((_1319_v23).Cardinality(), _dafny.IntOfUint32((_1336_v38).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), _dafny.IntOfUint32((_1290_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs298
			var _rhs299 bool = (_1317_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1317_v21), 0))).Int()).(bool)
			_ = _rhs299
			var _rhs300 *C5 = _1333_v35
			_ = _rhs300
			var _lhs219 _dafny.Array = _1327_v31
			_ = _lhs219
			var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1327_v31), 0))
			_ = _lhs220
			var _lhs221 _dafny.Array = _1332_v34
			_ = _lhs221
			var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1332_v34), 0))
			_ = _lhs222
			_1316_v20 = _rhs296
			(_lhs219).ArraySet1(_rhs297, (_lhs220).Int())
			(_lhs221).ArraySet1CodePoint(_rhs298, (_lhs222).Int())
			_1316_v20 = _rhs299
			_1333_v35 = _rhs300
		}
		var _1337_v39 _dafny.Sequence
		_ = _1337_v39
		_1337_v39 = _dafny.SeqOf(p2)
		var _1338_v40 _dafny.Map
		_ = _1338_v40
		_1338_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1337_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-696), _dafny.IntOfUint32((_1337_v39).Cardinality()))).Uint32()).(bool))
		var _1339_v41 _dafny.MultiSet
		_ = _1339_v41
		_1339_v41 = _dafny.MultiSetOf(Companion_Default___.Fm38(false, globalState))
		var _1340_v42 _dafny.Set
		_ = _1340_v42
		_1340_v42 = _dafny.SetOf((_1339_v41).Cardinality())
		_1338_v40 = (_1338_v40).Update(p2, (_1340_v42).IsDisjointFrom(_1340_v42))
		var _1341_v43 _dafny.Int
		_ = _1341_v43
		var _1342_v44 bool
		_ = _1342_v44
		var _out11 _dafny.Int
		_ = _out11
		var _out12 bool
		_ = _out12
		_out11, _out12 = (_this).M8(globalState)
		_1341_v43 = _out11
		_1342_v44 = _out12
		var _1343_v45 _dafny.CodePoint
		_ = _1343_v45
		_1343_v45 = _dafny.CodePoint('k')
		var _1344_v46 D15
		_ = _1344_v46
		_1344_v46 = Companion_D15_.Create_DC38_(_1297_v5)
		var _pat_let_tv21 = p0
		_ = _pat_let_tv21
		var _pat_let_tv22 = p0
		_ = _pat_let_tv22
		var _rhs301 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1290_v1, _dafny.Companion_Sequence_.Update(_1290_v1, (Companion_Default___.SafeIndex((_dafny.MultiSetOf(_1342_v44, !(_1342_v44))).Cardinality(), _dafny.IntOfUint32((_1290_v1).Cardinality()))).Uint32(), _1343_v45))
		_ = _rhs301
		var _rhs302 _dafny.Int = (_dafny.Zero).Minus(p0)
		_ = _rhs302
		var _rhs303 bool = !(!(_1342_v44)) || ((p0).Cmp(p0) == 0)
		_ = _rhs303
		var _rhs304 _dafny.Int = func(_source20 D15) _dafny.Int {
			if _source20.Is_DC39() {
				var _1345___mcc_h0 _dafny.Map = _source20.Get_().(D15_DC39).Cf56
				_ = _1345___mcc_h0
				var _1346___mcc_h1 bool = _source20.Get_().(D15_DC39).Cf57
				_ = _1346___mcc_h1
				var _1347___mcc_h2 _dafny.Int = _source20.Get_().(D15_DC39).Cf58
				_ = _1347___mcc_h2
				var _1348_cf58 _dafny.Int = _1347___mcc_h2
				_ = _1348_cf58
				var _1349_cf57 bool = _1346___mcc_h1
				_ = _1349_cf57
				var _1350_cf56 _dafny.Map = _1345___mcc_h0
				_ = _1350_cf56
				return _pat_let_tv21
			} else if _source20.Is_DC38() {
				var _1351___mcc_h3 _dafny.Map = _source20.Get_().(D15_DC38).Cf55
				_ = _1351___mcc_h3
				var _1352_cf55 _dafny.Map = _1351___mcc_h3
				_ = _1352_cf55
				return _pat_let_tv22
			} else {
				var _1353___mcc_h4 D15 = _source20.Get_().(D15_DC40).Cf59
				_ = _1353___mcc_h4
				var _1354_cf59 D15 = _1353___mcc_h4
				_ = _1354_cf59
				return (_dafny.IntOfInt64(626)).Times(_dafny.IntOfInt64(818))
			}
		}(_1344_v46)
		_ = _rhs304
		var _rhs305 _dafny.Int = p0
		_ = _rhs305
		var _lhs223 *GlobalState = globalState
		_ = _lhs223
		var _lhs224 *GlobalState = globalState
		_ = _lhs224
		_1290_v1 = _rhs301
		_lhs223.F15 = _rhs302
		_1342_v44 = _rhs303
		_lhs224.F4 = _rhs304
		_1341_v43 = _rhs305
	}
}
func (_this *C8) M8(globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _1355_v0 _dafny.Int
		_ = _1355_v0
		_1355_v0 = _dafny.IntOfInt64(919)
		(globalState).F15 = _1355_v0
		var _1356_v1 bool
		_ = _1356_v1
		_1356_v1 = false
		var _1357_v2 _dafny.Map
		_ = _1357_v2
		_1357_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1356_v1, _1355_v0)
		var _1358_v3 _dafny.Array
		_ = _1358_v3
		var _nwElement0_45 _dafny.Int = _dafny.IntOfInt64(-660)
		_ = _nwElement0_45
		var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(11))
		_ = _nw193
		(_nw193).ArraySet1(_nwElement0_45, 0)
		(_nw193).ArraySet1(((_dafny.Zero).Minus(_1355_v0)).Plus(_1355_v0), 1)
		(_nw193).ArraySet1(_1355_v0, 2)
		(_nw193).ArraySet1(_1355_v0, 3)
		(_nw193).ArraySet1(_dafny.IntOfInt64(541), 4)
		(_nw193).ArraySet1(_1355_v0, 5)
		(_nw193).ArraySet1(_1355_v0, 6)
		(_nw193).ArraySet1((_dafny.SetOf(_1357_v2)).Cardinality(), 7)
		(_nw193).ArraySet1(_1355_v0, 8)
		(_nw193).ArraySet1(_1355_v0, 9)
		(_nw193).ArraySet1((_1355_v0).Plus(_1355_v0), 10)
		_1358_v3 = _nw193
		var _1359_v4 _dafny.Map
		_ = _1359_v4
		_1359_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1356_v1, _1356_v1)
		_1358_v3 = (func() _dafny.Array {
			if !(_1359_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1356_v1)) {
				return _1358_v3
			}
			return _1358_v3
		})()
		var _1360_v5 D6
		_ = _1360_v5
		_1360_v5 = Companion_D6_.Create_DC17_(_1356_v1, _1356_v1, _1356_v1)
		var _pat_let_tv23 = _1355_v0
		_ = _pat_let_tv23
		var _pat_let_tv24 = _1356_v1
		_ = _pat_let_tv24
		var _pat_let_tv25 = _1356_v1
		_ = _pat_let_tv25
		var _pat_let_tv26 = _1355_v0
		_ = _pat_let_tv26
		(globalState).F13 = func(_source21 D6) _dafny.Int {
			if _source21.Is_DC17() {
				var _1361___mcc_h0 bool = _source21.Get_().(D6_DC17).Cf24
				_ = _1361___mcc_h0
				var _1362___mcc_h1 bool = _source21.Get_().(D6_DC17).Cf25
				_ = _1362___mcc_h1
				var _1363___mcc_h2 bool = _source21.Get_().(D6_DC17).Cf26
				_ = _1363___mcc_h2
				var _1364_cf26 bool = _1363___mcc_h2
				_ = _1364_cf26
				var _1365_cf25 bool = _1362___mcc_h1
				_ = _1365_cf25
				var _1366_cf24 bool = _1361___mcc_h0
				_ = _1366_cf24
				return _pat_let_tv23
			} else if _source21.Is_DC16() {
				var _1367___mcc_h3 _dafny.Sequence = _source21.Get_().(D6_DC16).Cf23
				_ = _1367___mcc_h3
				var _1368_cf23 _dafny.Sequence = _1367___mcc_h3
				_ = _1368_cf23
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_pat_let_tv24), _dafny.SeqOf(_pat_let_tv25))).Cardinality())
			} else {
				var _1369___mcc_h4 D6 = _source21.Get_().(D6_DC18).Cf27
				_ = _1369___mcc_h4
				var _1370_cf27 D6 = _1369___mcc_h4
				_ = _1370_cf27
				return _pat_let_tv26
			}
		}(_1360_v5)
		var _rhs306 bool = _1356_v1
		_ = _rhs306
		var _rhs307 _dafny.Int = Companion_Default___.Fm17(_1355_v0, globalState)
		_ = _rhs307
		var _lhs225 *GlobalState = globalState
		_ = _lhs225
		r1 = _rhs306
		_lhs225.F7 = _rhs307
		var _1371_v6 _dafny.Array
		_ = _1371_v6
		var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(13))
		_ = _nw194
		_1371_v6 = _nw194
		var _1372_v7 _dafny.CodePoint
		_ = _1372_v7
		_1372_v7 = _dafny.CodePoint('k')
		var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_1371_v6), 0))
		_ = _index245
		(_1371_v6).ArraySet1CodePoint(_1372_v7, (_index245).Int())
		var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_1371_v6), 0))
		_ = _index246
		(_1371_v6).ArraySet1CodePoint(_1372_v7, (_index246).Int())
		var _1373_v8 _dafny.MultiSet
		_ = _1373_v8
		_1373_v8 = _dafny.MultiSetOf(_1356_v1)
		var _1374_v9 _dafny.Array
		_ = _1374_v9
		var _nwElement0_46 _dafny.MultiSet = _1373_v8
		_ = _nwElement0_46
		var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(8))
		_ = _nw195
		(_nw195).ArraySet1(_nwElement0_46, 0)
		(_nw195).ArraySet1((_1373_v8).Intersection(_dafny.MultiSetOf(_1356_v1)), 1)
		(_nw195).ArraySet1(_1373_v8, 2)
		(_nw195).ArraySet1((_dafny.MultiSetOf(_1356_v1)).Union(_1373_v8), 3)
		(_nw195).ArraySet1(_dafny.MultiSetOf(_1356_v1), 4)
		(_nw195).ArraySet1((_dafny.MultiSetOf(_1356_v1)).Union(_1373_v8), 5)
		(_nw195).ArraySet1(_1373_v8, 6)
		(_nw195).ArraySet1((_1373_v8).Difference(_1373_v8), 7)
		_1374_v9 = _nw195
		var _1375_v10 _dafny.Sequence
		_ = _1375_v10
		_1375_v10 = _dafny.SeqOf(_1356_v1)
		var _1376_v11 _dafny.Sequence
		_ = _1376_v11
		_1376_v11 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_1375_v10), _dafny.MultiSetFromSeq(_1375_v10))
		var _1377_v12 _dafny.Set
		_ = _1377_v12
		_1377_v12 = _dafny.SetOf(_1356_v1, _1356_v1)
		var _1378_v13 _dafny.Sequence
		_ = _1378_v13
		_1378_v13 = _dafny.UnicodeSeqOfUtf8Bytes("apmhaajn")
		var _1379_v14 _dafny.Map
		_ = _1379_v14
		_1379_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_1357_v2).Contains(_1356_v1) {
				return (_1357_v2).Get(_1356_v1).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_1378_v13).Cardinality())
		})(), _1373_v8)
		var _nwElement0_47 _dafny.MultiSet = _dafny.MultiSetFromSeq(_1375_v10)
		_ = _nwElement0_47
		var _nw196 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(12))
		_ = _nw196
		(_nw196).ArraySet1(_nwElement0_47, 0)
		(_nw196).ArraySet1(_1373_v8, 1)
		(_nw196).ArraySet1(_1373_v8, 2)
		(_nw196).ArraySet1(_1373_v8, 3)
		(_nw196).ArraySet1(((_1376_v11).Select((Companion_Default___.SafeIndex(_1355_v0, _dafny.IntOfUint32((_1376_v11).Cardinality()))).Uint32()).(_dafny.MultiSet)).Difference(_1373_v8), 4)
		(_nw196).ArraySet1(_1373_v8, 5)
		(_nw196).ArraySet1((_1373_v8).Update(_1356_v1, Companion_Default___.Abs((_1377_v12).Cardinality())), 6)
		(_nw196).ArraySet1(_dafny.MultiSetFromSeq(_1375_v10), 7)
		(_nw196).ArraySet1(_1373_v8, 8)
		(_nw196).ArraySet1(_1373_v8, 9)
		(_nw196).ArraySet1(_dafny.MultiSetFromSeq(Companion_Default___.Fm27(globalState)), 10)
		(_nw196).ArraySet1((_1373_v8).Difference((func() _dafny.MultiSet {
			if (_1379_v14).Contains(_1355_v0) {
				return (_1379_v14).Get(_1355_v0).(_dafny.MultiSet)
			}
			return _1373_v8
		})()), 11)
		_1374_v9 = _nw196
		r0 = _1355_v0
		var _1380_v15 D3
		_ = _1380_v15
		_1380_v15 = Companion_D3_.Create_DC10_(_1355_v0, _1355_v0)
		var _pat_let_tv27 = _1356_v1
		_ = _pat_let_tv27
		var _pat_let_tv28 = _1356_v1
		_ = _pat_let_tv28
		var _pat_let_tv29 = _1356_v1
		_ = _pat_let_tv29
		r1 = func(_source22 D3) bool {
			if _source22.Is_DC10() {
				var _1381___mcc_h5 _dafny.Int = _source22.Get_().(D3_DC10).Cf13
				_ = _1381___mcc_h5
				var _1382___mcc_h6 _dafny.Int = _source22.Get_().(D3_DC10).Cf14
				_ = _1382___mcc_h6
				var _1383_cf14 _dafny.Int = _1382___mcc_h6
				_ = _1383_cf14
				var _1384_cf13 _dafny.Int = _1381___mcc_h5
				_ = _1384_cf13
				return !(_pat_let_tv27) || (_pat_let_tv28)
			} else {
				var _1385___mcc_h7 _dafny.CodePoint = _source22.Get_().(D3_DC9).Cf12
				_ = _1385___mcc_h7
				var _1386_cf12 _dafny.CodePoint = _1385___mcc_h7
				_ = _1386_cf12
				return _pat_let_tv29
			}
		}(_1380_v15)
		return r0, r1
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	dummy byte
}

func New_C9_() *C9 {
	_this := C9{}

	return &_this
}

type CompanionStruct_C9_ struct {
}

var Companion_C9_ = CompanionStruct_C9_{}

func (_this *C9) Equals(other *C9) bool {
	return _this == other
}

func (_this *C9) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C9)
	return ok && _this.Equals(other)
}

func (*C9) String() string {
	return "_module.C9"
}

func Type_C9_() _dafny.TypeDescriptor {
	return type_C9_{}
}

type type_C9_ struct {
}

func (_this type_C9_) Default() interface{} {
	return (*C9)(nil)
}

func (_this type_C9_) String() string {
	return "main.C9"
}
func (_this *C9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C9{}
var _ T1 = &C9{}
var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) Ctor__() {
	{
	}
}
func (_this *C9) Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((func() D1 {
			if !(!(true)) {
				return Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-16))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}(func(_1387_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				})), false)
			}
			return Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("r"), true)
		})()).Dtor_cf4()
	}
}
func (_this *C9) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), false)).Cardinality(), _dafny.IntOfInt64(330))
	}
}
func (_this *C9) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(-569)), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-381))).Cardinality()), true)).Cardinality(), (_dafny.SetOf(true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(6))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg84 _dafny.Int) interface{} {
				return coer84(arg84)
			}
		}(func(_1388_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		}))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(850))).Cardinality(), _dafny.IntOfInt64(49), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(24))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg85 _dafny.Int) interface{} {
				return coer85(arg85)
			}
		}(func(_1389_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality())))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("otolbkm")).Cardinality()))))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf((_dafny.MultiSetOf(true)).Cardinality()))))).Difference((_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(833)), _dafny.SetOf(_dafny.IntOfInt64(705)), _dafny.SetOf(_dafny.IntOfInt64(909)))).Difference(_dafny.MultiSetOf(func() _dafny.Set {
			var _coll62 = _dafny.NewBuilder()
			_ = _coll62
			for _iter64 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(829))).Elements()); ; {
				_compr_62, _ok64 := _iter64()
				if !_ok64 {
					break
				}
				var _1390_v0 _dafny.Int
				_1390_v0 = interface{}(_compr_62).(_dafny.Int)
				if (_dafny.SetOf(_dafny.IntOfInt64(829))).Contains(_1390_v0) {
					_coll62.Add(Companion_Default___.SafeDivisionInt(_1390_v0, (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(98))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg86 _dafny.Int) interface{} {
							return coer86(arg86)
						}
					}(func(_1391_i2 _dafny.Int) _dafny.Int {
						return (_dafny.MultiSetOf(true)).Cardinality()
					}))).Cardinality()), (_dafny.SetOf(!(false))).Cardinality())).Cardinality()))
				}
			}
			return _coll62.ToSet()
		}())))
	}
}
func (_this *C9) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(390))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(20)))).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qqyrsao")).Cardinality()), _dafny.IntOfInt64(481)))
	}
}
func (_this *C9) Fm9(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		var _source23 D2 = Companion_D2_.Create_DC6_(_dafny.IntOfInt64(-528))
		_ = _source23
		if _source23.Is_DC6() {
			var _1392___mcc_h0 _dafny.Int = _source23.Get_().(D2_DC6).Cf7
			_ = _1392___mcc_h0
			var _1393_cf7 _dafny.Int = _1392___mcc_h0
			_ = _1393_cf7
			return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_1393_cf7)).Cardinality()), _1393_cf7, _1393_cf7, _dafny.IntOfInt64(338)))).IsSubsetOf(_dafny.MultiSetOf((func() _dafny.Map {
				var _coll63 = _dafny.NewMapBuilder()
				_ = _coll63
				for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(484), _dafny.IntOfInt64(492))); ; {
					_compr_63, _ok65 := _iter65()
					if !_ok65 {
						break
					}
					var _1394_v0 _dafny.Int
					_1394_v0 = interface{}(_compr_63).(_dafny.Int)
					if ((_dafny.IntOfInt64(484)).Cmp(_1394_v0) <= 0) && ((_1394_v0).Cmp(_dafny.IntOfInt64(492)) < 0) {
						_coll63.Add(Companion_Default___.SafeDivisionInt(_1394_v0, _1393_cf7), !(false))
					}
				}
				return _coll63.ToMap()
			}()).Cardinality()))
		} else if _source23.Is_DC7() {
			var _1395___mcc_h1 bool = _source23.Get_().(D2_DC7).Cf8
			_ = _1395___mcc_h1
			var _1396___mcc_h2 _dafny.MultiSet = _source23.Get_().(D2_DC7).Cf9
			_ = _1396___mcc_h2
			var _1397___mcc_h3 _dafny.Array = _source23.Get_().(D2_DC7).Cf10
			_ = _1397___mcc_h3
			var _1398_cf10 _dafny.Array = _1397___mcc_h3
			_ = _1398_cf10
			var _1399_cf9 _dafny.MultiSet = _1396___mcc_h2
			_ = _1399_cf9
			var _1400_cf8 bool = _1395___mcc_h1
			_ = _1400_cf8
			return _1400_cf8
		} else if _source23.Is_DC5() {
			var _1401___mcc_h4 _dafny.Array = _source23.Get_().(D2_DC5).Cf6
			_ = _1401___mcc_h4
			var _1402_cf6 _dafny.Array = _1401___mcc_h4
			_ = _1402_cf6
			return ((_dafny.Zero).Minus(_dafny.IntOfInt64(-84))).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()) >= 0
		} else {
			var _1403___mcc_h5 D2 = _source23.Get_().(D2_DC8).Cf11
			_ = _1403___mcc_h5
			var _1404_cf11 D2 = _1403___mcc_h5
			_ = _1404_cf11
			return !(!(!(!(true))))
		}
	}
}
func (_this *C9) Fm10(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C9) M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _1405_v0 _dafny.Int
		_ = _1405_v0
		_1405_v0 = _dafny.IntOfInt64(-428)
		var _1406_v1 _dafny.Map
		_ = _1406_v1
		_1406_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v0, _1405_v0)
		var _1407_v3 _dafny.Map
		_ = _1407_v3
		_1407_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(78))
		var _1408_v4 _dafny.Sequence
		_ = _1408_v4
		_1408_v4 = _dafny.SeqOf((_1407_v3).Cardinality(), _1405_v0)
		var _1409_v10 bool
		_ = _1409_v10
		_1409_v10 = false
		var _1410_v11 _dafny.Array
		_ = _1410_v11
		var _nwElement0_48 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(_1405_v0, globalState), _1405_v0)
		_ = _nwElement0_48
		var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(23))
		_ = _nw197
		(_nw197).ArraySet1(_nwElement0_48, 0)
		(_nw197).ArraySet1(_1406_v1, 1)
		(_nw197).ArraySet1(func() _dafny.Map {
			var _coll64 = _dafny.NewMapBuilder()
			_ = _coll64
			for _iter66 := _dafny.Iterate((_1408_v4).Elements()); ; {
				_compr_64, _ok66 := _iter66()
				if !_ok66 {
					break
				}
				var _1411_v2 _dafny.Int
				_1411_v2 = interface{}(_compr_64).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_1408_v4, _1411_v2) {
					_coll64.Add((_1411_v2).Plus(_1405_v0), _1405_v0)
				}
			}
			return _coll64.ToMap()
		}(), 2)
		(_nw197).ArraySet1(func() _dafny.Map {
			var _coll65 = _dafny.NewMapBuilder()
			_ = _coll65
			for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(42), _dafny.IntOfInt64(698))); ; {
				_compr_65, _ok67 := _iter67()
				if !_ok67 {
					break
				}
				var _1412_v5 _dafny.Int
				_1412_v5 = interface{}(_compr_65).(_dafny.Int)
				if ((_dafny.IntOfInt64(42)).Cmp(_1412_v5) <= 0) && ((_1412_v5).Cmp(_dafny.IntOfInt64(698)) < 0) {
					_coll65.Add(Companion_Default___.SafeDivisionInt(_1412_v5, _1405_v0), _1405_v0)
				}
			}
			return _coll65.ToMap()
		}(), 3)
		(_nw197).ArraySet1(_1406_v1, 4)
		(_nw197).ArraySet1(_1406_v1, 5)
		(_nw197).ArraySet1(_1406_v1, 6)
		(_nw197).ArraySet1((_1406_v1).Update(_1405_v0, _1405_v0), 7)
		(_nw197).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1408_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_1408_v4).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(-607)), 8)
		(_nw197).ArraySet1(func() _dafny.Map {
			var _coll66 = _dafny.NewMapBuilder()
			_ = _coll66
			for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(326), _dafny.IntOfInt64(-776))); ; {
				_compr_66, _ok68 := _iter68()
				if !_ok68 {
					break
				}
				var _1413_v6 _dafny.Int
				_1413_v6 = interface{}(_compr_66).(_dafny.Int)
				if ((_dafny.IntOfInt64(326)).Cmp(_1413_v6) <= 0) && ((_1413_v6).Cmp(_dafny.IntOfInt64(-776)) < 0) {
					_coll66.Add(Companion_Default___.SafeDivisionInt(_1413_v6, _1405_v0), (_1407_v3).Cardinality())
				}
			}
			return _coll66.ToMap()
		}(), 9)
		(_nw197).ArraySet1((_1406_v1).Merge(func() _dafny.Map {
			var _coll67 = _dafny.NewMapBuilder()
			_ = _coll67
			for _iter69 := _dafny.Iterate((_dafny.SeqOf(_1405_v0, _1405_v0)).Elements()); ; {
				_compr_67, _ok69 := _iter69()
				if !_ok69 {
					break
				}
				var _1414_v7 _dafny.Int
				_1414_v7 = interface{}(_compr_67).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1405_v0, _1405_v0), _1414_v7) {
					_coll67.Add(Companion_Default___.SafeDivisionInt(_1414_v7, _1405_v0), _1405_v0)
				}
			}
			return _coll67.ToMap()
		}()), 10)
		(_nw197).ArraySet1((_1406_v1).Merge(_1406_v1), 11)
		(_nw197).ArraySet1((_1406_v1).Merge(_1406_v1), 12)
		(_nw197).ArraySet1(func() _dafny.Map {
			var _coll68 = _dafny.NewMapBuilder()
			_ = _coll68
			for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-666), _dafny.IntOfInt64(400))); ; {
				_compr_68, _ok70 := _iter70()
				if !_ok70 {
					break
				}
				var _1415_v8 _dafny.Int
				_1415_v8 = interface{}(_compr_68).(_dafny.Int)
				if ((_dafny.IntOfInt64(-666)).Cmp(_1415_v8) <= 0) && ((_1415_v8).Cmp(_dafny.IntOfInt64(400)) < 0) {
					_coll68.Add(Companion_Default___.SafeModuloInt(_1415_v8, _dafny.IntOfInt64(365)), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(true, true)).Cardinality())))
				}
			}
			return _coll68.ToMap()
		}(), 13)
		(_nw197).ArraySet1((_1406_v1).Merge(_1406_v1), 14)
		(_nw197).ArraySet1(_1406_v1, 15)
		(_nw197).ArraySet1((func() _dafny.Map {
			var _coll69 = _dafny.NewMapBuilder()
			_ = _coll69
			for _iter71 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v0, _1409_v10)).Keys().Elements()); ; {
				_compr_69, _ok71 := _iter71()
				if !_ok71 {
					break
				}
				var _1416_v9 _dafny.Int
				_1416_v9 = interface{}(_compr_69).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v0, _1409_v10)).Contains(_1416_v9) {
					_coll69.Add((_1416_v9).Times(_dafny.IntOfInt64(215)), _1405_v0)
				}
			}
			return _coll69.ToMap()
		}()).Update(_1405_v0, _1405_v0), 16)
		(_nw197).ArraySet1((_1406_v1).Merge(_1406_v1), 17)
		(_nw197).ArraySet1(_1406_v1, 18)
		(_nw197).ArraySet1((func() _dafny.Map {
			if _1409_v10 {
				return _1406_v1
			}
			return _1406_v1
		})(), 19)
		(_nw197).ArraySet1(_1406_v1, 20)
		(_nw197).ArraySet1(_1406_v1, 21)
		(_nw197).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-308), _1405_v0), 22)
		_1410_v11 = _nw197
		var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1410_v11), 0))
		_ = _index247
		(_1410_v11).ArraySet1((func() _dafny.Map {
			var _coll70 = _dafny.NewMapBuilder()
			_ = _coll70
			for _iter72 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(537), _dafny.IntOfInt64(936))); ; {
				_compr_70, _ok72 := _iter72()
				if !_ok72 {
					break
				}
				var _1417_v12 _dafny.Int
				_1417_v12 = interface{}(_compr_70).(_dafny.Int)
				if ((_dafny.IntOfInt64(537)).Cmp(_1417_v12) <= 0) && ((_1417_v12).Cmp(_dafny.IntOfInt64(936)) < 0) {
					_coll70.Add((_1417_v12).Times(_dafny.IntOfUint32((_1408_v4).Cardinality())), _1405_v0)
				}
			}
			return _coll70.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v0, _dafny.IntOfUint32((p0).Cardinality()))), (_index247).Int())
		var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1410_v11), 0))
		_ = _index248
		(_1410_v11).ArraySet1((_1406_v1).Merge((func() _dafny.Map {
			if !(_1409_v10) {
				return _1406_v1
			}
			return _1406_v1
		})()), (_index248).Int())
		_1406_v1 = (_1406_v1).Update((_1408_v4).Select((Companion_Default___.SafeIndex(_1405_v0, _dafny.IntOfUint32((_1408_v4).Cardinality()))).Uint32()).(_dafny.Int), _1405_v0)
		var _1418_v13 _dafny.Set
		_ = _1418_v13
		_1418_v13 = _dafny.SetOf(_1405_v0, _1405_v0)
		var _hi10 _dafny.Int = (_1418_v13).Cardinality()
		_ = _hi10
		for _1419_i0 := _1405_v0; _1419_i0.Cmp(_hi10) < 0; _1419_i0 = _1419_i0.Plus(_dafny.One) {
			var _1420_v14 _dafny.Map
			_ = _1420_v14
			_1420_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), _1409_v10)
			if (func() bool {
				if (_1420_v14).Contains(_1405_v0) {
					return (_1420_v14).Get(_1405_v0).(bool)
				}
				return _1409_v10
			})() {
				_1409_v10 = _1409_v10
				var _1421_v15 _dafny.MultiSet
				_ = _1421_v15
				_1421_v15 = _dafny.MultiSetOf(_1409_v10)
				var _1422_v16 T1
				_ = _1422_v16
				var _nw198 *C8 = New_C8_()
				_ = _nw198
				_nw198.Ctor__()
				_1422_v16 = _nw198
				var _1423_v17 _dafny.Map
				_ = _1423_v17
				_1423_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1422_v16, _1409_v10)
				var _1424_v18 _dafny.Map
				_ = _1424_v18
				_1424_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1409_v10, _1423_v17)
				var _1425_v19 _dafny.Set
				_ = _1425_v19
				_1425_v19 = _dafny.SetOf(_1409_v10)
				var _1426_v20 _dafny.Array
				_ = _1426_v20
				var _nwElement0_49 _dafny.Int = (_1420_v14).Cardinality()
				_ = _nwElement0_49
				var _nw199 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(10))
				_ = _nw199
				(_nw199).ArraySet1(_nwElement0_49, 0)
				(_nw199).ArraySet1(_1419_i0, 1)
				(_nw199).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1409_v10, !(_1409_v10))).Update(_1409_v10, _1409_v10)).Cardinality(), 2)
				(_nw199).ArraySet1(_1405_v0, 3)
				(_nw199).ArraySet1((_1424_v18).Cardinality(), 4)
				(_nw199).ArraySet1(_dafny.IntOfInt64(829), 5)
				(_nw199).ArraySet1((Companion_Default___.Fm32(_1419_i0, _1419_i0, _1425_v19, globalState)).Cardinality(), 6)
				(_nw199).ArraySet1(_1405_v0, 7)
				(_nw199).ArraySet1(_1405_v0, 8)
				(_nw199).ArraySet1(_1419_i0, 9)
				_1426_v20 = _nw199
				var _1427_v21 D2
				_ = _1427_v21
				_1427_v21 = Companion_D2_.Create_DC7_(false, _1421_v15, _1426_v20)
				var _1428_v22 _dafny.Sequence
				_ = _1428_v22
				_1428_v22 = _dafny.SeqOf(_1409_v10, _1409_v10, _1409_v10, true, _1409_v10)
				var _1429_v23 _dafny.Array
				_ = _1429_v23
				var _nwElement0_50 D2 = _1427_v21
				_ = _nwElement0_50
				var _nw200 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(22))
				_ = _nw200
				(_nw200).ArraySet1(_nwElement0_50, 0)
				(_nw200).ArraySet1(_1427_v21, 1)
				(_nw200).ArraySet1(_1427_v21, 2)
				(_nw200).ArraySet1(_1427_v21, 3)
				(_nw200).ArraySet1(_1427_v21, 4)
				(_nw200).ArraySet1(_1427_v21, 5)
				(_nw200).ArraySet1((func() D2 {
					if _1409_v10 {
						return _1427_v21
					}
					return _1427_v21
				})(), 6)
				(_nw200).ArraySet1(_1427_v21, 7)
				(_nw200).ArraySet1(_1427_v21, 8)
				(_nw200).ArraySet1(_1427_v21, 9)
				(_nw200).ArraySet1(_1427_v21, 10)
				(_nw200).ArraySet1(Companion_D2_.Create_DC7_(_1409_v10, _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1428_v22, (Companion_Default___.SafeIndex(_1405_v0, _dafny.IntOfUint32((_1428_v22).Cardinality()))).Uint32(), _1409_v10)), _1426_v20), 11)
				(_nw200).ArraySet1(_1427_v21, 12)
				(_nw200).ArraySet1(_1427_v21, 13)
				(_nw200).ArraySet1(_1427_v21, 14)
				(_nw200).ArraySet1(Companion_D2_.Create_DC7_(_1409_v10, _1421_v15, _1426_v20), 15)
				(_nw200).ArraySet1(_1427_v21, 16)
				(_nw200).ArraySet1(_1427_v21, 17)
				(_nw200).ArraySet1(Companion_D2_.Create_DC7_(true, _dafny.MultiSetOf(_1409_v10), _1426_v20), 18)
				(_nw200).ArraySet1(Companion_D2_.Create_DC7_(_1409_v10, _1421_v15, _1426_v20), 19)
				(_nw200).ArraySet1(_1427_v21, 20)
				(_nw200).ArraySet1(Companion_D2_.Create_DC7_(_1409_v10, Companion_Default___.Fm34(_dafny.IntOfUint32((Companion_Default___.Fm19(_1409_v10, !(_1409_v10), globalState)).Cardinality()), _1405_v0, (func() bool {
					if (_1420_v14).Contains(_1419_i0) {
						return (_1420_v14).Get(_1419_i0).(bool)
					}
					return _1409_v10
				})(), globalState), _1426_v20), 21)
				_1429_v23 = _nw200
				var _nw201 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(25))
				_ = _nw201
				_1429_v23 = _nw201
				_1409_v10 = (_1425_v19).IsSubsetOf(_1425_v19)
				_1409_v10 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1408_v4, _1408_v4), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(305))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_1430_v10 bool, _1431_v3 _dafny.Map, _1432_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1433_i1 _dafny.Int) _dafny.Int {
						return (func() _dafny.Int {
							if (_1431_v3).Contains(_1430_v10) {
								return (_1431_v3).Get(_1430_v10).(_dafny.Int)
							}
							return _1432_i0
						})()
					}
				})(_1409_v10, _1407_v3, _1419_i0))))
				var _1434_v24 _dafny.Map
				_ = _1434_v24
				_1434_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), _1409_v10)
				var _1435_v25 _dafny.CodePoint
				_ = _1435_v25
				_1435_v25 = _dafny.CodePoint('f')
				var _1436_v26 _dafny.Array
				_ = _1436_v26
				var _nwElement0_51 bool = _1409_v10
				_ = _nwElement0_51
				var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(17))
				_ = _nw202
				(_nw202).ArraySet1(_nwElement0_51, 0)
				(_nw202).ArraySet1(_1409_v10, 1)
				(_nw202).ArraySet1(_1409_v10, 2)
				(_nw202).ArraySet1(false, 3)
				(_nw202).ArraySet1(_1409_v10, 4)
				(_nw202).ArraySet1(_1409_v10, 5)
				(_nw202).ArraySet1(_1409_v10, 6)
				(_nw202).ArraySet1(_1409_v10, 7)
				(_nw202).ArraySet1(_1409_v10, 8)
				(_nw202).ArraySet1(_1409_v10, 9)
				(_nw202).ArraySet1(_1409_v10, 10)
				(_nw202).ArraySet1(_1409_v10, 11)
				(_nw202).ArraySet1(_1409_v10, 12)
				(_nw202).ArraySet1(!(_1409_v10), 13)
				(_nw202).ArraySet1(_1409_v10, 14)
				(_nw202).ArraySet1((func() bool {
					if (_1434_v24).Contains(_1435_v25) {
						return (_1434_v24).Get(_1435_v25).(bool)
					}
					return false
				})(), 15)
				(_nw202).ArraySet1(_1409_v10, 16)
				_1436_v26 = _nw202
				var _1437_v27 _dafny.Map
				_ = _1437_v27
				_1437_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1409_v10, _1436_v26)
				_1437_v27 = (_1437_v27).Update(_1409_v10, (func() _dafny.Array {
					if true {
						return _1436_v26
					}
					return _1436_v26
				})())
			} else {
				var _1438_v28 *C2
				_ = _1438_v28
				var _nw203 *C2 = New_C2_()
				_ = _nw203
				_nw203.Ctor__()
				_1438_v28 = _nw203
				_1438_v28 = _1438_v28
				(globalState).F4 = _1405_v0
				var _1439_v29 _dafny.Map
				_ = _1439_v29
				_1439_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1409_v10, p0)
				_1439_v29 = _1439_v29
				var _1440_v30 *C5
				_ = _1440_v30
				var _nw204 *C5 = New_C5_()
				_ = _nw204
				_nw204.Ctor__()
				_1440_v30 = _nw204
				_1409_v10 = _1409_v10
			}
			var _1441_v31 _dafny.CodePoint
			_ = _1441_v31
			_1441_v31 = _dafny.CodePoint('h')
			var _1442_v32 _dafny.Array
			_ = _1442_v32
			var _nw205 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw205
			_1442_v32 = _nw205
			var _1443_v33 _dafny.Sequence
			_ = _1443_v33
			_1443_v33 = _dafny.SeqOf(_1409_v10, !(false) || (false), _1409_v10)
			var _1444_v34 D5
			_ = _1444_v34
			_1444_v34 = Companion_D5_.Create_DC14_(_1441_v31, _1409_v10, _1405_v0, _1409_v10)
			var _rhs308 bool = (_1409_v10) == (_1409_v10)
			_ = _rhs308
			var _rhs309 bool = !((_1443_v33).Select((Companion_Default___.SafeIndex(_1419_i0, _dafny.IntOfUint32((_1443_v33).Cardinality()))).Uint32()).(bool))
			_ = _rhs309
			var _rhs310 _dafny.CodePoint = (_1444_v34).Dtor_cf18()
			_ = _rhs310
			var _rhs311 _dafny.Array = _1442_v32
			_ = _rhs311
			var _rhs312 _dafny.Int = (_this).Fm2(Companion_Default___.SafeDivisionInt(_1405_v0, _1419_i0), globalState)
			_ = _rhs312
			var _lhs226 *GlobalState = globalState
			_ = _lhs226
			_1409_v10 = _rhs308
			_1409_v10 = _rhs309
			_1441_v31 = _rhs310
			_1442_v32 = _rhs311
			_lhs226.F13 = _rhs312
			var _1445_v35 _dafny.Map
			_ = _1445_v35
			_1445_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_1441_v31), _1442_v32)
			var _1446_v36 _dafny.MultiSet
			_ = _1446_v36
			_1446_v36 = _dafny.MultiSetOf(_1441_v31, _1441_v31, _1441_v31, _1441_v31)
			_1445_v35 = (_1445_v35).Update(_1446_v36, _1442_v32)
			var _1447_v37 _dafny.MultiSet
			_ = _1447_v37
			_1447_v37 = _dafny.MultiSetOf(_1405_v0, (_1407_v3).Cardinality(), _1419_i0)
			var _1448_v38 *C6
			_ = _1448_v38
			var _nw206 *C6 = New_C6_()
			_ = _nw206
			_nw206.Ctor__()
			_1448_v38 = _nw206
			var _1449_v39 _dafny.MultiSet
			_ = _1449_v39
			_1449_v39 = _dafny.MultiSetOf(_1448_v38, _1448_v38, _1448_v38)
			var _1450_v40 _dafny.Array
			_ = _1450_v40
			var _nwElement0_52 _dafny.Int = _dafny.IntOfUint32((p0).Cardinality())
			_ = _nwElement0_52
			var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(17))
			_ = _nw207
			(_nw207).ArraySet1(_nwElement0_52, 0)
			(_nw207).ArraySet1(_dafny.IntOfInt64(264), 1)
			(_nw207).ArraySet1(_1419_i0, 2)
			(_nw207).ArraySet1((_1447_v37).Cardinality(), 3)
			(_nw207).ArraySet1((_1419_i0).Minus((_dafny.Zero).Minus(_1405_v0)), 4)
			(_nw207).ArraySet1(_dafny.IntOfInt64(496), 5)
			(_nw207).ArraySet1(_1419_i0, 6)
			(_nw207).ArraySet1((_this).Fm2(_1405_v0, globalState), 7)
			(_nw207).ArraySet1(_1419_i0, 8)
			(_nw207).ArraySet1(_1405_v0, 9)
			(_nw207).ArraySet1(_1405_v0, 10)
			(_nw207).ArraySet1(_1405_v0, 11)
			(_nw207).ArraySet1(_1405_v0, 12)
			(_nw207).ArraySet1((_dafny.IntOfInt64(210)).Plus((_1447_v37).Cardinality()), 13)
			(_nw207).ArraySet1(_1405_v0, 14)
			(_nw207).ArraySet1((func() _dafny.Int {
				if _1409_v10 {
					return (_1449_v39).Cardinality()
				}
				return _1405_v0
			})(), 15)
			(_nw207).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-570), _1405_v0), 16)
			_1450_v40 = _nw207
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1450_v40), 0))
			_ = _index249
			(_1450_v40).ArraySet1(_1405_v0, (_index249).Int())
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1450_v40), 0))
			_ = _index250
			(_1450_v40).ArraySet1(_1405_v0, (_index250).Int())
		}
		var _1451_v41 _dafny.Array
		_ = _1451_v41
		var _nw208 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
		_ = _nw208
		_1451_v41 = _nw208
		var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))
		_ = _index251
		(_1451_v41).ArraySet1(false, (_index251).Int())
		var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))
		_ = _index252
		(_1451_v41).ArraySet1(_1409_v10, (_index252).Int())
		var _1452_v42 D7
		_ = _1452_v42
		_1452_v42 = Companion_D7_.Create_DC21_(_1405_v0, _1405_v0, _1409_v10, (_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool), _1409_v10)
		if !(((_1452_v42).Dtor_cf32()) == ((_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool))) {
			var _1453_v43 _dafny.Array
			_ = _1453_v43
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_35
			var _nw209 _dafny.Array
			_ = _nw209
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw209 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.MultiSet = (func(_1454_v41 _dafny.Array) func(_dafny.Int) _dafny.MultiSet {
					return func(_1455_i2 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(!(!(true)), (_1454_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1454_v41), 0))).Int()).(bool))
					}
				})(_1451_v41)
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw209 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw209).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw209).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1453_v43 = _nw209
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1453_v43), 0))
			_ = _index253
			(_1453_v43).ArraySet1(_dafny.MultiSetOf(true), (_index253).Int())
			var _1456_v44 _dafny.Set
			_ = _1456_v44
			_1456_v44 = _dafny.SetOf((_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool))
			var _1457_v45 _dafny.Map
			_ = _1457_v45
			_1457_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1456_v44, (_this).Fm9(_1405_v0, globalState))
			var _1458_v46 _dafny.MultiSet
			_ = _1458_v46
			_1458_v46 = _dafny.MultiSetOf((func() bool {
				if (_1457_v45).Contains(_1456_v44) {
					return (_1457_v45).Get(_1456_v44).(bool)
				}
				return (_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool)
			})())
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1453_v43), 0))
			_ = _index254
			(_1453_v43).ArraySet1(_1458_v46, (_index254).Int())
			var _1459_v47 _dafny.Array
			_ = _1459_v47
			var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw210
			_1459_v47 = _nw210
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1459_v47), 0))
			_ = _index255
			(_1459_v47).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-664))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg88 _dafny.Int) interface{} {
					return coer88(arg88)
				}
			}(func(_1460_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			}))).Cardinality()), (_index255).Int())
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1459_v47), 0))
			_ = _index256
			(_1459_v47).ArraySet1(_1405_v0, (_index256).Int())
			_1459_v47 = _1459_v47
			var _1461_v48 _dafny.Sequence
			_ = _1461_v48
			_1461_v48 = _dafny.SeqOf(_1409_v10)
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))
			_ = _index257
			var _rhs313 bool = true
			_ = _rhs313
			var _rhs314 _dafny.Sequence = _1461_v48
			_ = _rhs314
			var _rhs315 _dafny.Set = ((_1456_v44).Difference(_1456_v44)).Difference(_1456_v44)
			_ = _rhs315
			var _rhs316 _dafny.Int = (_1459_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1459_v47), 0))).Int()).(_dafny.Int)
			_ = _rhs316
			var _lhs227 _dafny.Array = _1451_v41
			_ = _lhs227
			var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))
			_ = _lhs228
			var _lhs229 *GlobalState = globalState
			_ = _lhs229
			(_lhs227).ArraySet1(_rhs313, (_lhs228).Int())
			_1461_v48 = _rhs314
			_1456_v44 = _rhs315
			_lhs229.F13 = _rhs316
			var _1462_v49 _dafny.CodePoint
			_ = _1462_v49
			_1462_v49 = _dafny.CodePoint('h')
			var _1463_v50 _dafny.Sequence
			_ = _1463_v50
			_1463_v50 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_1405_v0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _1462_v49), p0), Companion_Default___.Fm19(_1409_v10, _1409_v10, globalState), _dafny.Companion_Sequence_.Concatenate(p0, p0), _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_1459_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_1459_v47), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _1462_v49))
			_1463_v50 = _1463_v50
		} else {
			(globalState).F7 = _1405_v0
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))
			_ = _index258
			(_1451_v41).ArraySet1((func() bool {
				if _1409_v10 {
					return (_1405_v0).Cmp(_1405_v0) > 0
				}
				return _1409_v10
			})(), (_index258).Int())
			var _1464_v51 *C8
			_ = _1464_v51
			var _nw211 *C8 = New_C8_()
			_ = _nw211
			_nw211.Ctor__()
			_1464_v51 = _nw211
			var _1465_v52 _dafny.Array
			_ = _1465_v52
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_36
			var _nw212 _dafny.Array
			_ = _nw212
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw212 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) _dafny.Int = (func(_1466_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1467_i4 _dafny.Int) _dafny.Int {
						return (_1467_i4).Minus(_1466_v0)
					}
				})(_1405_v0)
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw212 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw212).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw212).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1465_v52 = _nw212
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1465_v52), 0))
			_ = _index259
			(_1465_v52).ArraySet1(_1405_v0, (_index259).Int())
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1465_v52), 0))
			_ = _index260
			(_1465_v52).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.Fm17((_this).Fm2(_1405_v0, globalState), globalState)).Times(Companion_Default___.SafeModuloInt(_1405_v0, _1405_v0)))), (_index260).Int())
			if ((_1465_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1465_v52), 0))).Int()).(_dafny.Int)).Cmp(Companion_Default___.SafeModuloInt((_1465_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1465_v52), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p0).Cardinality()))) < 0 {
				var _1468_v53 D6
				_ = _1468_v53
				_1468_v53 = Companion_D6_.Create_DC17_((_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool), (_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool), (_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool))
				_1409_v10 = (_this).Fm8((_1465_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1465_v52), 0))).Int()).(_dafny.Int), (func() bool {
					if _1409_v10 {
						return (_1468_v53).Dtor_cf26()
					}
					return _1409_v10
				})(), (_dafny.Zero).Minus(_1405_v0), globalState)
				var _1469_v54 *C2
				_ = _1469_v54
				var _nw213 *C2 = New_C2_()
				_ = _nw213
				_nw213.Ctor__()
				_1469_v54 = _nw213
				var _1470_v55 _dafny.MultiSet
				_ = _1470_v55
				_1470_v55 = _dafny.MultiSetOf((_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool))
				var _1471_v56 _dafny.Map
				_ = _1471_v56
				_1471_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1465_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1465_v52), 0))).Int()).(_dafny.Int), _1409_v10)
				var _1472_v57 D15
				_ = _1472_v57
				_1472_v57 = Companion_D15_.Create_DC38_(_1471_v56)
				var _1473_v58 _dafny.Map
				_ = _1473_v58
				_1473_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1470_v55).Difference(_1470_v55), _1472_v57)
				var _1474_v59 _dafny.CodePoint
				_ = _1474_v59
				_1474_v59 = _dafny.CodePoint('u')
				var _1475_v60 _dafny.Map
				_ = _1475_v60
				_1475_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1409_v10, _1474_v59)
				_1473_v58 = (_1473_v58).Update(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).Fm10(_1475_v60, _1405_v0, (_1410_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1410_v11), 0))).Int()).(_dafny.Map), globalState), _1409_v10, (_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool), (_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool))), _1472_v57)
				(globalState).F4 = _1405_v0
				(globalState).F4 = _1405_v0
			} else {
				_1408_v4 = _1408_v4
				var _1476_v61 *C4
				_ = _1476_v61
				var _nw214 *C4 = New_C4_()
				_ = _nw214
				_nw214.Ctor__()
				_1476_v61 = _nw214
				(globalState).F4 = (_1465_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1465_v52), 0))).Int()).(_dafny.Int)
				_1407_v3 = (_1407_v3).Update(_1409_v10, (_1465_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_1465_v52), 0))).Int()).(_dafny.Int))
				var _1477_v62 _dafny.MultiSet
				_ = _1477_v62
				_1477_v62 = _dafny.MultiSetOf(!((_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool)))
				var _1478_v63 D2
				_ = _1478_v63
				_1478_v63 = Companion_D2_.Create_DC5_(_1465_v52)
				(_1464_v51).M6(_dafny.IntOfInt64(-905), (_1477_v62).Union(_1477_v62), true, _1478_v63, globalState)
			}
		}
		var _1479_v64 _dafny.Array
		_ = _1479_v64
		var _nw215 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw215
		_1479_v64 = _nw215
		var _1480_v65 D2
		_ = _1480_v65
		_1480_v65 = Companion_D2_.Create_DC7_((_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool), _dafny.MultiSetOf((_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool)), _1479_v64)
		var _1481_v66 D2
		_ = _1481_v66
		_1481_v66 = Companion_D2_.Create_DC8_(_1480_v65)
		var _1482_v67 _dafny.Map
		_ = _1482_v67
		_1482_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v0, _1409_v10)
		var _1483_v68 *C3
		_ = _1483_v68
		var _nw216 *C3 = New_C3_()
		_ = _nw216
		_nw216.Ctor__(_1481_v66, (func() bool {
			if (_1482_v67).Contains(_1405_v0) {
				return (_1482_v67).Get(_1405_v0).(bool)
			}
			return (_1451_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1451_v41), 0))).Int()).(bool)
		})())
		_1483_v68 = _nw216
	}
}
func (_this *C9) M6(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 D2, globalState *GlobalState) {
	{
		var _1484_v0 _dafny.Set
		_ = _1484_v0
		_1484_v0 = _dafny.SetOf(p0)
		var _1485_v1 _dafny.Sequence
		_ = _1485_v1
		_1485_v1 = _dafny.SeqOf(_1484_v0)
		var _1486_v2 _dafny.Sequence
		_ = _1486_v2
		_1486_v2 = _dafny.UnicodeSeqOfUtf8Bytes("fwnxp")
		var _1487_v3 D5
		_ = _1487_v3
		_1487_v3 = Companion_D5_.Create_DC13_(_dafny.MultiSetOf((_1485_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1486_v2).Cardinality()), _dafny.IntOfUint32((_1485_v1).Cardinality()))).Uint32()).(_dafny.Set)))
		var _source24 D5 = _1487_v3
		_ = _source24
		if _source24.Is_DC14() {
			var _1488___mcc_h0 _dafny.CodePoint = _source24.Get_().(D5_DC14).Cf18
			_ = _1488___mcc_h0
			var _1489___mcc_h1 bool = _source24.Get_().(D5_DC14).Cf19
			_ = _1489___mcc_h1
			var _1490___mcc_h2 _dafny.Int = _source24.Get_().(D5_DC14).Cf20
			_ = _1490___mcc_h2
			var _1491___mcc_h3 bool = _source24.Get_().(D5_DC14).Cf21
			_ = _1491___mcc_h3
			var _1492_cf21 bool = _1491___mcc_h3
			_ = _1492_cf21
			var _1493_cf20 _dafny.Int = _1490___mcc_h2
			_ = _1493_cf20
			var _1494_cf19 bool = _1489___mcc_h1
			_ = _1494_cf19
			var _1495_cf18 _dafny.CodePoint = _1488___mcc_h0
			_ = _1495_cf18
			var _1496_v4 _dafny.MultiSet
			_ = _1496_v4
			_1496_v4 = _dafny.MultiSetOf(_dafny.IntOfInt64(435))
			var _1497_v5 _dafny.Map
			_ = _1497_v5
			_1497_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1496_v4)
			var _1498_v6 _dafny.Sequence
			_ = _1498_v6
			_1498_v6 = _dafny.SeqOf(Companion_D1_.Create_DC4_((func() _dafny.MultiSet {
				if (_1497_v5).Contains(p0) {
					return (_1497_v5).Get(p0).(_dafny.MultiSet)
				}
				return _1496_v4
			})()))
			_1498_v6 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-20))).Uint32(), func(coer89 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
				return func(arg89 _dafny.Int) interface{} {
					return coer89(arg89)
				}
			}((func(_1499_p0 _dafny.Int, _1500_cf20 _dafny.Int, _1501_cf19 bool, _1502_cf21 bool) func(_dafny.Int) D1 {
				return func(_1503_i0 _dafny.Int) D1 {
					return Companion_D1_.Create_DC4_(_dafny.MultiSetOf((_dafny.SetOf(_1499_p0, _1500_cf20, _1500_cf20, _dafny.IntOfInt64(-24), (_dafny.SetOf(_1501_cf19, _1502_cf21, !(_1502_cf21))).Cardinality())).Cardinality(), _1499_p0, _1499_p0))
				}
			})(p0, _1493_cf20, _1494_cf19, _1492_cf21)))
			var _1504_v7 D11
			_ = _1504_v7
			_1504_v7 = Companion_D11_.Create_DC30_(p2, Companion_Default___.Fm17((_dafny.Zero).Minus(p0), globalState))
			(globalState).F7 = (_1504_v7).Dtor_cf47()
			var _1505_v8 _dafny.Array
			_ = _1505_v8
			var _nw217 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw217
			_1505_v8 = _nw217
			var _1506_v9 _dafny.Sequence
			_ = _1506_v9
			_1506_v9 = _dafny.SeqOf(_1505_v8, _1505_v8, _1505_v8)
			var _1507_v10 _dafny.Map
			_ = _1507_v10
			_1507_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1505_v8)
			_1506_v9 = _dafny.SeqOf(_1505_v8, _1505_v8, _1505_v8, (func() _dafny.Array {
				if (_1507_v10).Contains(_1494_cf19) {
					return (_1507_v10).Get(_1494_cf19).(_dafny.Array)
				}
				return _1505_v8
			})(), _1505_v8)
			if !(!(_1494_cf19) || ((p1).IsProperSubsetOf(p1))) {
				_1505_v8 = _1505_v8
				var _1508_v12 D11
				_ = _1508_v12
				_1508_v12 = Companion_D11_.Create_DC29_()
				var _1509_v13 _dafny.Set
				_ = _1509_v13
				_1509_v13 = _dafny.SetOf(_1508_v12, _1508_v12)
				var _1510_v14 _dafny.Map
				_ = _1510_v14
				_1510_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1494_cf19)
				var _1511_v15 D12
				_ = _1511_v15
				_1511_v15 = Companion_D12_.Create_DC32_(_dafny.SetOf(p2))
				var _1512_v16 _dafny.Sequence
				_ = _1512_v16
				_1512_v16 = _dafny.SeqOf(_1511_v15, _1511_v15)
				var _1513_v17 _dafny.Map
				_ = _1513_v17
				_1513_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v13, (_1510_v14).Update(_dafny.IntOfUint32((_1512_v16).Cardinality()), _1492_cf21))
				(globalState).F13 = (func() _dafny.Map {
					var _coll71 = _dafny.NewMapBuilder()
					_ = _coll71
					for _iter73 := _dafny.Iterate((_1513_v17).Keys().Elements()); ; {
						_compr_71, _ok73 := _iter73()
						if !_ok73 {
							break
						}
						var _1514_v11 _dafny.Set
						_1514_v11 = interface{}(_compr_71).(_dafny.Set)
						if (_1513_v17).Contains(_1514_v11) {
							_coll71.Add(_1514_v11, false)
						}
					}
					return _coll71.ToMap()
				}()).Cardinality()
				var _1515_v18 _dafny.Map
				_ = _1515_v18
				_1515_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1494_cf19)
				var _1516_v19 _dafny.Map
				_ = _1516_v19
				_1516_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1515_v18)
				var _1517_v20 _dafny.Sequence
				_ = _1517_v20
				_1517_v20 = _dafny.SeqOf(p2, _1492_cf21, _1492_cf21, (func() bool {
					if _1492_cf21 {
						return _1492_cf21
					}
					return (_this).Fm1(_1516_v19, _dafny.IntOfUint32((_1486_v2).Cardinality()), globalState)
				})())
				var _1518_v21 D16
				_ = _1518_v21
				_1518_v21 = Companion_D16_.Create_DC42_(_1494_cf19, _1492_cf21, _1494_cf19)
				var _1519_v22 _dafny.Array
				_ = _1519_v22
				var _nwElement0_53 D16 = _1518_v21
				_ = _nwElement0_53
				var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(12))
				_ = _nw218
				(_nw218).ArraySet1(_nwElement0_53, 0)
				(_nw218).ArraySet1(_1518_v21, 1)
				(_nw218).ArraySet1(_1518_v21, 2)
				(_nw218).ArraySet1(_1518_v21, 3)
				(_nw218).ArraySet1(Companion_D16_.Create_DC42_(p2, (_1518_v21).Dtor_cf61(), _1494_cf19), 4)
				(_nw218).ArraySet1(_1518_v21, 5)
				(_nw218).ArraySet1(Companion_Default___.Fm46(p2, globalState), 6)
				(_nw218).ArraySet1(_1518_v21, 7)
				(_nw218).ArraySet1(_1518_v21, 8)
				(_nw218).ArraySet1(_1518_v21, 9)
				(_nw218).ArraySet1(_1518_v21, 10)
				(_nw218).ArraySet1(_1518_v21, 11)
				_1519_v22 = _nw218
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_1519_v22), 0))
				_ = _index261
				(_1519_v22).ArraySet1(_1518_v21, (_index261).Int())
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_1519_v22), 0))
				_ = _index262
				var _rhs317 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1517_v20, _dafny.Companion_Sequence_.Update(_1517_v20, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1517_v20).Cardinality()))).Uint32(), !(_1492_cf21)))
				_ = _rhs317
				var _rhs318 _dafny.Int = p0
				_ = _rhs318
				var _rhs319 D16 = _1518_v21
				_ = _rhs319
				var _lhs230 *GlobalState = globalState
				_ = _lhs230
				var _lhs231 _dafny.Array = _1519_v22
				_ = _lhs231
				var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_1519_v22), 0))
				_ = _lhs232
				_1517_v20 = _rhs317
				_lhs230.F13 = _rhs318
				(_lhs231).ArraySet1(_rhs319, (_lhs232).Int())
				var _1520_v23 _dafny.Int
				_ = _1520_v23
				var _1521_v24 _dafny.MultiSet
				_ = _1521_v24
				var _out13 _dafny.Int
				_ = _out13
				var _out14 _dafny.MultiSet
				_ = _out14
				_out13, _out14 = (_this).M7(Companion_Default___.SafeModuloInt(_1493_cf20, _dafny.IntOfInt64(821)), _1494_cf19, globalState)
				_1520_v23 = _out13
				_1521_v24 = _out14
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1505_v8), 0))
				_ = _index263
				(_1505_v8).ArraySet1((_this).Fm8(_1493_cf20, _1492_cf21, p0, globalState), (_index263).Int())
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1505_v8), 0))
				_ = _index264
				(_1505_v8).ArraySet1(!(false), (_index264).Int())
			} else {
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_1505_v8), 0))
				_ = _index265
				(_1505_v8).ArraySet1(_1492_cf21, (_index265).Int())
				var _1522_v25 _dafny.Sequence
				_ = _1522_v25
				_1522_v25 = _dafny.SeqOf((_dafny.Zero).Minus(_1493_cf20))
				var _1523_v26 _dafny.MultiSet
				_ = _1523_v26
				_1523_v26 = _dafny.MultiSetOf(_1486_v2, _dafny.Companion_Sequence_.Update(_1486_v2, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1522_v25).Cardinality()), _dafny.IntOfUint32((_1486_v2).Cardinality()))).Uint32(), _1495_cf18), _1486_v2)
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_1505_v8), 0))
				_ = _index266
				(_1505_v8).ArraySet1((Companion_Default___.Fm47(p2, (_dafny.Zero).Minus(_1493_cf20), globalState)).IsProperSubsetOf(_1523_v26), (_index266).Int())
				_1486_v2 = _dafny.Companion_Sequence_.Concatenate(_1486_v2, _dafny.Companion_Sequence_.Concatenate(_1486_v2, _1486_v2))
				var _1524_v27 *C1
				_ = _1524_v27
				var _nw219 *C1 = New_C1_()
				_ = _nw219
				_nw219.Ctor__()
				_1524_v27 = _nw219
				var _1525_v28 _dafny.Sequence
				_ = _1525_v28
				_1525_v28 = _dafny.SeqOf(_1524_v27, _1524_v27)
				var _1526_v29 _dafny.Map
				_ = _1526_v29
				_1526_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(133), _1525_v28)
				var _1527_v30 _dafny.Map
				_ = _1527_v30
				_1527_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(771), _1526_v29)
				var _1528_v31 _dafny.Map
				_ = _1528_v31
				_1528_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1494_cf19), _1493_cf20)
				var _1529_v32 _dafny.Array
				_ = _1529_v32
				var _nwElement0_54 _dafny.Map = (_1526_v29).Merge((_1526_v29).Update(_1493_cf20, _1525_v28))
				_ = _nwElement0_54
				var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(23))
				_ = _nw220
				(_nw220).ArraySet1(_nwElement0_54, 0)
				(_nw220).ArraySet1(_1526_v29, 1)
				(_nw220).ArraySet1(_1526_v29, 2)
				(_nw220).ArraySet1(_1526_v29, 3)
				(_nw220).ArraySet1((_1526_v29).Update(_1493_cf20, _1525_v28), 4)
				(_nw220).ArraySet1(_1526_v29, 5)
				(_nw220).ArraySet1((func() _dafny.Map {
					if _1492_cf21 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1486_v2).Cardinality()), _1525_v28)
					}
					return _1526_v29
				})(), 6)
				(_nw220).ArraySet1(_1526_v29, 7)
				(_nw220).ArraySet1(_1526_v29, 8)
				(_nw220).ArraySet1((_1526_v29).Merge(_1526_v29), 9)
				(_nw220).ArraySet1((_1526_v29).Merge((func() _dafny.Map {
					if (_1527_v30).Contains((_1522_v25).Select((Companion_Default___.SafeIndex(_1493_cf20, _dafny.IntOfUint32((_1522_v25).Cardinality()))).Uint32()).(_dafny.Int)) {
						return (_1527_v30).Get((_1522_v25).Select((Companion_Default___.SafeIndex(_1493_cf20, _dafny.IntOfUint32((_1522_v25).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1493_cf20, _dafny.SeqOf(_1524_v27))
				})()), 10)
				(_nw220).ArraySet1(_1526_v29, 11)
				(_nw220).ArraySet1((_1526_v29).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1525_v28)), 12)
				(_nw220).ArraySet1((_1526_v29).Merge(_1526_v29), 13)
				(_nw220).ArraySet1(_1526_v29, 14)
				(_nw220).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1525_v28)).Update((_dafny.Zero).Minus(_1493_cf20), _1525_v28), 15)
				(_nw220).ArraySet1(_1526_v29, 16)
				(_nw220).ArraySet1(_1526_v29, 17)
				(_nw220).ArraySet1(_1526_v29, 18)
				(_nw220).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1525_v28)).Update(((_1528_v31).Update(_1494_cf19, p0)).Cardinality(), _1525_v28), 19)
				(_nw220).ArraySet1(_1526_v29, 20)
				(_nw220).ArraySet1((_1526_v29).Merge(_1526_v29), 21)
				(_nw220).ArraySet1((_1526_v29).Merge(_1526_v29), 22)
				_1529_v32 = _nw220
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_1529_v32), 0))
				_ = _index267
				(_1529_v32).ArraySet1(_1526_v29, (_index267).Int())
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_1529_v32), 0))
				_ = _index268
				(_1529_v32).ArraySet1(_1526_v29, (_index268).Int())
				(globalState).F4 = p0
				var _1530_v33 _dafny.Map
				_ = _1530_v33
				_1530_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1492_cf21, _1494_cf19)
				var _1531_v34 _dafny.Map
				_ = _1531_v34
				_1531_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1530_v33).Cardinality(), _1493_cf20)
				_1531_v34 = (_1531_v34).Update((_1493_cf20).Plus((_dafny.Zero).Minus(p0)), (_dafny.IntOfInt64(241)).Minus(_1493_cf20))
			}
		} else if _source24.Is_DC13() {
			var _1532___mcc_h4 _dafny.MultiSet = _source24.Get_().(D5_DC13).Cf17
			_ = _1532___mcc_h4
			var _1533_cf17 _dafny.MultiSet = _1532___mcc_h4
			_ = _1533_cf17
			var _1534_v35 bool
			_ = _1534_v35
			_1534_v35 = false
			_1534_v35 = p2
			var _1535_v36 _dafny.Sequence
			_ = _1535_v36
			_1535_v36 = _dafny.SeqOf(_1486_v2, (func() _dafny.Sequence {
				if p2 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(279))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg90 _dafny.Int) interface{} {
							return coer90(arg90)
						}
					}(func(_1536_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('t')
					}))
				}
				return _1486_v2
			})(), _dafny.Companion_Sequence_.Concatenate(_1486_v2, _1486_v2), _1486_v2, _1486_v2)
			var _1537_v37 T0
			_ = _1537_v37
			var _nw221 *C2 = New_C2_()
			_ = _nw221
			_nw221.Ctor__()
			_1537_v37 = _nw221
			var _1538_v38 _dafny.Set
			_ = _1538_v38
			_1538_v38 = _dafny.SetOf(_1537_v37, _1537_v37, _1537_v37, _1537_v37, _1537_v37)
			var _1539_v39 _dafny.Map
			_ = _1539_v39
			_1539_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1538_v38).Cardinality(), p0)
			var _1540_v40 _dafny.Sequence
			_ = _1540_v40
			_1540_v40 = _dafny.SeqOf(p0, p0, (_1539_v39).Cardinality(), _dafny.IntOfInt64(269))
			var _1541_v41 _dafny.Sequence
			_ = _1541_v41
			_1541_v41 = _dafny.SeqOf((_1484_v0).Cardinality(), (_dafny.Zero).Minus((_1540_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("is")).Cardinality()), _dafny.IntOfUint32((_1540_v40).Cardinality()))).Uint32()).(_dafny.Int)), (_dafny.Zero).Minus(p0))
			var _1542_v42 _dafny.Sequence
			_ = _1542_v42
			_1542_v42 = _dafny.SeqOf(_1540_v40)
			_1486_v2 = (_1535_v36).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1542_v42).Cardinality()), _dafny.IntOfUint32((_1535_v36).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _1543_v43 _dafny.Array
			_ = _1543_v43
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_37
			var _nw222 _dafny.Array
			_ = _nw222
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw222 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) bool = (func(_1544_v35 bool) func(_dafny.Int) bool {
					return func(_1545_i2 _dafny.Int) bool {
						return _1544_v35
					}
				})(_1534_v35)
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw222 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw222).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw222).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1543_v43 = _nw222
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1543_v43), 0))
			_ = _index269
			(_1543_v43).ArraySet1(p2, (_index269).Int())
			var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1543_v43), 0))
			_ = _index270
			(_1543_v43).ArraySet1(!(p2), (_index270).Int())
			var _1546_v44 _dafny.Array
			_ = _1546_v44
			var _nw223 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw223
			_1546_v44 = _nw223
			var _1547_v45 _dafny.Array
			_ = _1547_v45
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_38
			var _nw224 _dafny.Array
			_ = _nw224
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw224 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) _dafny.Int = func(_1548_i3 _dafny.Int) _dafny.Int {
					return (_1548_i3).Plus(_dafny.IntOfInt64(-411))
				}
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw224 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw224).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw224).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1547_v45 = _nw224
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_1546_v44), 0))
			_ = _index271
			(_1546_v44).ArraySet1(_1547_v45, (_index271).Int())
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_1546_v44), 0))
			_ = _index272
			var _rhs320 _dafny.Array = _1547_v45
			_ = _rhs320
			var _rhs321 bool = (_1543_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1543_v43), 0))).Int()).(bool)
			_ = _rhs321
			var _rhs322 _dafny.Map = _1539_v39
			_ = _rhs322
			var _rhs323 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1486_v2).Cardinality()), p0)
			_ = _rhs323
			var _lhs233 _dafny.Array = _1546_v44
			_ = _lhs233
			var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_1546_v44), 0))
			_ = _lhs234
			var _lhs235 *GlobalState = globalState
			_ = _lhs235
			(_lhs233).ArraySet1(_rhs320, (_lhs234).Int())
			_1534_v35 = _rhs321
			_1539_v39 = _rhs322
			_lhs235.F7 = _rhs323
		} else {
			var _1549___mcc_h5 D5 = _source24.Get_().(D5_DC15).Cf22
			_ = _1549___mcc_h5
			var _1550_cf22 D5 = _1549___mcc_h5
			_ = _1550_cf22
			if (p0).Cmp((func() _dafny.Int {
				if (p1).Contains(true) {
					return (p1).Multiplicity(true)
				}
				return p0
			})()) > 0 {
				var _1551_v46 _dafny.MultiSet
				_ = _1551_v46
				_1551_v46 = _dafny.MultiSetOf(_dafny.IntOfInt64(793), p0, p0)
				var _1552_v47 _dafny.Sequence
				_ = _1552_v47
				_1552_v47 = _dafny.SeqOf((_1551_v46).Contains(p0), p2, (_this).Fm9(p0, globalState))
				_1552_v47 = _1552_v47
				(globalState).F13 = p0
				var _1553_v48 *C6
				_ = _1553_v48
				var _nw225 *C6 = New_C6_()
				_ = _nw225
				_nw225.Ctor__()
				_1553_v48 = _nw225
				var _1554_v49 _dafny.Sequence
				_ = _1554_v49
				_1554_v49 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), p0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dkgfi"), _1486_v2)).Cardinality()), p0)
				_1554_v49 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(237))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg91 _dafny.Int) interface{} {
						return coer91(arg91)
					}
				}((func(_1555_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1556_i4 _dafny.Int) _dafny.Int {
						return _1555_p0
					}
				})(p0)))
				var _1557_v50 _dafny.Map
				_ = _1557_v50
				_1557_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), p2)
				var _1558_v51 _dafny.Map
				_ = _1558_v51
				_1558_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(898), (func() bool {
					if (_1557_v50).Contains((func() _dafny.Int {
						if (_1551_v46).Contains(p0) {
							return (_1551_v46).Multiplicity(p0)
						}
						return p0
					})()) {
						return (_1557_v50).Get((func() _dafny.Int {
							if (_1551_v46).Contains(p0) {
								return (_1551_v46).Multiplicity(p0)
							}
							return p0
						})()).(bool)
					}
					return (_1552_v47).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1552_v47).Cardinality()))).Uint32()).(bool)
				})())
				var _1559_v52 _dafny.Sequence
				_ = _1559_v52
				_1559_v52 = _dafny.SeqOf(_1558_v51)
				var _1560_v53 D14
				_ = _1560_v53
				_1560_v53 = Companion_D14_.Create_DC35_(_1559_v52)
				var _1561_v54 _dafny.Sequence
				_ = _1561_v54
				_1561_v54 = _dafny.SeqOf(_1560_v53)
				var _1562_v55 D1
				_ = _1562_v55
				_1562_v55 = Companion_D1_.Create_DC3_(_1486_v2, p2)
				_1561_v54 = _dafny.Companion_Sequence_.Concatenate(_1561_v54, _dafny.SeqOf(Companion_D14_.Create_DC35_(_1559_v52), Companion_Default___.Fm48(!(p2), p2, (_1562_v55).Dtor_cf4(), globalState)))
			} else {
				var _1563_v56 bool
				_ = _1563_v56
				_1563_v56 = false
				var _1564_v57 _dafny.Map
				_ = _1564_v57
				_1564_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1484_v0)
				_1563_v56 = (_1484_v0).Equals((func() _dafny.Set {
					if (_1564_v57).Contains(p0) {
						return (_1564_v57).Get(p0).(_dafny.Set)
					}
					return _1484_v0
				})())
				_1563_v56 = false
				(globalState).F7 = ((_1484_v0).Cardinality()).Times(p0)
				var _1565_v58 *C5
				_ = _1565_v58
				var _nw226 *C5 = New_C5_()
				_ = _nw226
				_nw226.Ctor__()
				_1565_v58 = _nw226
				var _1566_v59 _dafny.Sequence
				_ = _1566_v59
				_1566_v59 = _dafny.SeqOf(_1563_v56, p2)
				var _1567_v60 _dafny.Sequence
				_ = _1567_v60
				_1567_v60 = _dafny.SeqOf(p2, (_1566_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1566_v59).Cardinality()), _dafny.IntOfUint32((_1566_v59).Cardinality()))).Uint32()).(bool))
				var _1568_v61 _dafny.Set
				_ = _1568_v61
				_1568_v61 = _dafny.SetOf(p2)
				var _1569_v62 D2
				_ = _1569_v62
				_1569_v62 = Companion_D2_.Create_DC6_(Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus(p0)))
				var _rhs324 _dafny.Sequence = _1486_v2
				_ = _rhs324
				var _rhs325 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1566_v59, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1563_v56, p2), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.IntOfUint32((_dafny.SeqOf(_1563_v56, p2)).Cardinality()))).Uint32(), p2))
				_ = _rhs325
				var _rhs326 _dafny.Set = _1568_v61
				_ = _rhs326
				var _rhs327 D2 = _1569_v62
				_ = _rhs327
				_1486_v2 = _rhs324
				_1566_v59 = _rhs325
				_1568_v61 = _rhs326
				_1569_v62 = _rhs327
			}
			var _1570_v63 _dafny.Map
			_ = _1570_v63
			_1570_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(246))
			var _1571_v64 _dafny.Map
			_ = _1571_v64
			_1571_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1570_v63, p2)
			var _1572_v65 D17
			_ = _1572_v65
			_1572_v65 = Companion_D17_.Create_DC45_(_1571_v64)
			_1571_v64 = (func() _dafny.Map {
				if !(p2) {
					return _1571_v64
				}
				return (_1571_v64).Merge((_1572_v65).Dtor_cf67())
			})()
			var _1573_v66 bool
			_ = _1573_v66
			_1573_v66 = true
			var _1574_v67 _dafny.CodePoint
			_ = _1574_v67
			_1574_v67 = _dafny.CodePoint('j')
			var _1575_v68 _dafny.Map
			_ = _1575_v68
			_1575_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1574_v67, _1573_v66)
			var _1576_v69 _dafny.Array
			_ = _1576_v69
			var _nwElement0_55 bool = p2
			_ = _nwElement0_55
			var _nw227 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(8))
			_ = _nw227
			(_nw227).ArraySet1(_nwElement0_55, 0)
			(_nw227).ArraySet1(false, 1)
			(_nw227).ArraySet1((_this).Fm9(p0, globalState), 2)
			(_nw227).ArraySet1(true, 3)
			(_nw227).ArraySet1(_1573_v66, 4)
			(_nw227).ArraySet1((func() bool {
				if (_1575_v68).Contains(_1574_v67) {
					return (_1575_v68).Get(_1574_v67).(bool)
				}
				return p2
			})(), 5)
			(_nw227).ArraySet1(p2, 6)
			(_nw227).ArraySet1(p2, 7)
			_1576_v69 = _nw227
			var _1577_v70 _dafny.Set
			_ = _1577_v70
			_1577_v70 = _dafny.SetOf(_1576_v69, _1576_v69, _1576_v69)
			var _1578_v71 _dafny.Sequence
			_ = _1578_v71
			_1578_v71 = _dafny.SeqOf(_1573_v66, p2, !(_1573_v66))
			var _rhs328 bool = p2
			_ = _rhs328
			var _rhs329 bool = (_1578_v71).Select((Companion_Default___.SafeIndex((p0).Minus((_1570_v63).Cardinality()), _dafny.IntOfUint32((_1578_v71).Cardinality()))).Uint32()).(bool)
			_ = _rhs329
			var _rhs330 _dafny.Set = _1577_v70
			_ = _rhs330
			_1573_v66 = _rhs328
			_1573_v66 = _rhs329
			_1577_v70 = _rhs330
			var _1579_v72 D14
			_ = _1579_v72
			_1579_v72 = Companion_D14_.Create_DC37_(Companion_D14_.Create_DC36_(_1578_v71))
			var _1580_v73 D14
			_ = _1580_v73
			_1580_v73 = Companion_D14_.Create_DC36_(_1578_v71)
			var _1581_v74 _dafny.Map
			_ = _1581_v74
			_1581_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1573_v66)
			var _pat_let_tv30 = _1580_v73
			_ = _pat_let_tv30
			var _pat_let_tv31 = _1580_v73
			_ = _pat_let_tv31
			var _1582_v75 _dafny.Array
			_ = _1582_v75
			var _nwElement0_56 D14 = _1579_v72
			_ = _nwElement0_56
			var _nw228 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(25))
			_ = _nw228
			(_nw228).ArraySet1(_nwElement0_56, 0)
			(_nw228).ArraySet1(_1579_v72, 1)
			(_nw228).ArraySet1(_1579_v72, 2)
			(_nw228).ArraySet1(func(_pat_let20_0 D14) D14 {
				return func(_1583_dt__update__tmp_h0 D14) D14 {
					return func(_pat_let21_0 D14) D14 {
						return func(_1584_dt__update_hcf54_h0 D14) D14 {
							return Companion_D14_.Create_DC37_(_1584_dt__update_hcf54_h0)
						}(_pat_let21_0)
					}(_pat_let_tv30)
				}(_pat_let20_0)
			}(Companion_D14_.Create_DC37_(_1580_v73)), 3)
			(_nw228).ArraySet1(_1579_v72, 4)
			(_nw228).ArraySet1(func(_pat_let22_0 D14) D14 {
				return func(_1585_dt__update__tmp_h1 D14) D14 {
					return func(_pat_let23_0 D14) D14 {
						return func(_1586_dt__update_hcf54_h1 D14) D14 {
							return Companion_D14_.Create_DC37_(_1586_dt__update_hcf54_h1)
						}(_pat_let23_0)
					}(_pat_let_tv31)
				}(_pat_let22_0)
			}(_1579_v72), 5)
			(_nw228).ArraySet1(_1579_v72, 6)
			(_nw228).ArraySet1(_1579_v72, 7)
			(_nw228).ArraySet1(_1579_v72, 8)
			(_nw228).ArraySet1(_1579_v72, 9)
			(_nw228).ArraySet1(_1579_v72, 10)
			(_nw228).ArraySet1(Companion_D14_.Create_DC37_(_1580_v73), 11)
			(_nw228).ArraySet1(_1579_v72, 12)
			(_nw228).ArraySet1(Companion_D14_.Create_DC37_(Companion_Default___.Fm48(_1573_v66, p2, (func() bool {
				if (_1581_v74).Contains(p0) {
					return (_1581_v74).Get(p0).(bool)
				}
				return p2
			})(), globalState)), 13)
			(_nw228).ArraySet1(_1579_v72, 14)
			(_nw228).ArraySet1(_1579_v72, 15)
			(_nw228).ArraySet1(_1579_v72, 16)
			(_nw228).ArraySet1(_1579_v72, 17)
			(_nw228).ArraySet1(_1579_v72, 18)
			(_nw228).ArraySet1(_1579_v72, 19)
			(_nw228).ArraySet1((func() D14 {
				if p2 {
					return _1579_v72
				}
				return _1579_v72
			})(), 20)
			(_nw228).ArraySet1(Companion_D14_.Create_DC37_(_1580_v73), 21)
			(_nw228).ArraySet1(_1579_v72, 22)
			(_nw228).ArraySet1(_1579_v72, 23)
			(_nw228).ArraySet1(Companion_D14_.Create_DC37_(Companion_D14_.Create_DC35_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(969))).Uint32(), func(coer92 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg92 _dafny.Int) interface{} {
					return coer92(arg92)
				}
			}((func(_1587_v74 _dafny.Map, _1588_p0 _dafny.Int, _1589_p2 bool) func(_dafny.Int) _dafny.Map {
				return func(_1590_i5 _dafny.Int) _dafny.Map {
					return (_1587_v74).Update(_1588_p0, _1589_p2)
				}
			})(_1581_v74, p0, p2))))), 24)
			_1582_v75 = _nw228
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_1582_v75), 0))
			_ = _index273
			(_1582_v75).ArraySet1(_1579_v72, (_index273).Int())
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_1582_v75), 0))
			_ = _index274
			(_1582_v75).ArraySet1(_1579_v72, (_index274).Int())
		}
		if !(((p0).Cmp(p0) <= 0) && (!((p0).Cmp(_dafny.IntOfInt64(807)) >= 0))) {
			_1486_v2 = _dafny.Companion_Sequence_.Concatenate(_1486_v2, _dafny.UnicodeSeqOfUtf8Bytes("eiseggug"))
			var _1591_v76 *C0
			_ = _1591_v76
			var _nw229 *C0 = New_C0_()
			_ = _nw229
			_nw229.Ctor__()
			_1591_v76 = _nw229
			var _1592_v77 _dafny.Array
			_ = _1592_v77
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_39
			var _nw230 _dafny.Array
			_ = _nw230
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw230 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.CodePoint = func(_1593_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw230 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw230).ArraySet1CodePoint(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw230).ArraySet1CodePoint(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1592_v77 = _nw230
			_1592_v77 = _1592_v77
			if !(((_1484_v0).Cardinality()).Cmp(p0) == 0) {
				(globalState).F13 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p0), p0)
				var _1594_v78 _dafny.Sequence
				_ = _1594_v78
				_1594_v78 = _dafny.SeqOf(p2, p2, p2, p2, p2)
				var _1595_v79 _dafny.Sequence
				_ = _1595_v79
				_1595_v79 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_1594_v78).Cardinality()))
				var _1596_v80 _dafny.Sequence
				_ = _1596_v80
				_1596_v80 = _dafny.SeqOf(p0, p0, _dafny.IntOfUint32((_1595_v79).Cardinality()))
				var _1597_v81 _dafny.Sequence
				_ = _1597_v81
				_1597_v81 = _dafny.SeqOf(_dafny.SeqOf(p0), _1596_v80)
				var _1598_v82 _dafny.Map
				_ = _1598_v82
				_1598_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1486_v2)
				var _1599_v83 _dafny.CodePoint
				_ = _1599_v83
				_1599_v83 = _dafny.CodePoint('p')
				var _1600_v84 _dafny.Set
				_ = _1600_v84
				_1600_v84 = _dafny.SetOf(p2)
				var _1601_v85 _dafny.Sequence
				_ = _1601_v85
				_1601_v85 = _dafny.SeqOf(_1600_v84)
				var _1602_v86 _dafny.Map
				_ = _1602_v86
				_1602_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1601_v85, _1486_v2)
				var _1603_v87 _dafny.Map
				_ = _1603_v87
				_1603_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _1604_v88 T0
				_ = _1604_v88
				var _nw231 *C2 = New_C2_()
				_ = _nw231
				_nw231.Ctor__()
				_1604_v88 = _nw231
				var _1605_v89 _dafny.Sequence
				_ = _1605_v89
				_1605_v89 = _dafny.SeqOf(Companion_Default___.Fm49(p0, (_dafny.SetOf(_1604_v88, _1604_v88)).Cardinality(), globalState))
				var _1606_v90 _dafny.Map
				_ = _1606_v90
				_1606_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (p1).Cardinality())).Update(p2, _dafny.IntOfInt64(-649))).Cardinality(), _dafny.SetOf((_1603_v87).Cardinality(), _dafny.IntOfUint32((_1605_v89).Cardinality())))
				var _1607_v91 _dafny.Map
				_ = _1607_v91
				_1607_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1597_v81).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1597_v81).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _1608_v92 _dafny.Array
				_ = _1608_v92
				var _nwElement0_57 _dafny.Sequence = (_1591_v76).Fm28(_1597_v81, globalState)
				_ = _nwElement0_57
				var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(27))
				_ = _nw232
				(_nw232).ArraySet1(_nwElement0_57, 0)
				(_nw232).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1486_v2, _1486_v2), 1)
				(_nw232).ArraySet1((func() _dafny.Sequence {
					if p2 {
						return _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
							if (_1598_v82).Contains(p2) {
								return (_1598_v82).Get(p2).(_dafny.Sequence)
							}
							return _1486_v2
						})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_1598_v82).Contains(p2) {
								return (_1598_v82).Get(p2).(_dafny.Sequence)
							}
							return _1486_v2
						})()).Cardinality()))).Uint32(), _1599_v83)
					}
					return _1486_v2
				})(), 2)
				(_nw232).ArraySet1(_1486_v2, 3)
				(_nw232).ArraySet1(_1486_v2, 4)
				(_nw232).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1486_v2, _1486_v2), 5)
				(_nw232).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(145))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg93 _dafny.Int) interface{} {
						return coer93(arg93)
					}
				}((func(_1609_v83 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1610_i7 _dafny.Int) _dafny.CodePoint {
						return _1609_v83
					}
				})(_1599_v83))), 6)
				(_nw232).ArraySet1((func() _dafny.Sequence {
					if (_1602_v86).Contains(_dafny.SeqOf(_dafny.SetOf(p2, p2, p2), _dafny.SetOf(p2, false), _1600_v84, _1600_v84)) {
						return (_1602_v86).Get(_dafny.SeqOf(_dafny.SetOf(p2, p2, p2), _dafny.SetOf(p2, false), _1600_v84, _1600_v84)).(_dafny.Sequence)
					}
					return _1486_v2
				})(), 7)
				(_nw232).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kuauuuwp"), 8)
				(_nw232).ArraySet1(Companion_Default___.Fm30(_dafny.IntOfInt64(535), p2, _dafny.SetOf((_1606_v90).Cardinality(), _dafny.IntOfInt64(769)), p2, globalState), 9)
				(_nw232).ArraySet1((func() _dafny.Sequence {
					if p2 {
						return _1486_v2
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("eiqjnpdrf")
				})(), 10)
				(_nw232).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1486_v2, _1486_v2), 11)
				(_nw232).ArraySet1(_1486_v2, 12)
				(_nw232).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1486_v2, (_1591_v76).Fm28(_1597_v81, globalState)), 13)
				(_nw232).ArraySet1(_1486_v2, 14)
				(_nw232).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1486_v2, _1486_v2), 15)
				(_nw232).ArraySet1(_1486_v2, 16)
				(_nw232).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1486_v2, _1486_v2), 17)
				(_nw232).ArraySet1((_1591_v76).Fm28(_dafny.SeqOf((func() _dafny.Sequence {
					if (_1607_v91).Contains((_dafny.MultiSetOf(p2)).Cardinality()) {
						return (_1607_v91).Get((_dafny.MultiSetOf(p2)).Cardinality()).(_dafny.Sequence)
					}
					return _1595_v79
				})(), _1596_v80), globalState), 18)
				(_nw232).ArraySet1(_1486_v2, 19)
				(_nw232).ArraySet1((func() _dafny.Sequence {
					if (_1598_v82).Contains(p2) {
						return (_1598_v82).Get(p2).(_dafny.Sequence)
					}
					return _1486_v2
				})(), 20)
				(_nw232).ArraySet1(_1486_v2, 21)
				(_nw232).ArraySet1(_1486_v2, 22)
				(_nw232).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("sxbmntpf"), 23)
				(_nw232).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ggfo"), _1486_v2), 24)
				(_nw232).ArraySet1(_1486_v2, 25)
				(_nw232).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kkqchmef"), 26)
				_1608_v92 = _nw232
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1608_v92), 0))
				_ = _index275
				(_1608_v92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1486_v2, _1486_v2), (_index275).Int())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_1608_v92), 0))
				_ = _index276
				(_1608_v92).ArraySet1(_1486_v2, (_index276).Int())
				(globalState).F15 = p0
				var _1611_v93 _dafny.Array
				_ = _1611_v93
				var _nw233 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw233
				_1611_v93 = _nw233
				var _1612_v94 D2
				_ = _1612_v94
				_1612_v94 = Companion_D2_.Create_DC5_(_1611_v93)
				_1612_v94 = Companion_D2_.Create_DC5_(_1611_v93)
				(globalState).F4 = _dafny.IntOfInt64(493)
			} else {
				var _1613_v95 bool
				_ = _1613_v95
				_1613_v95 = true
				var _1614_v96 _dafny.Array
				_ = _1614_v96
				var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw234
				_1614_v96 = _nw234
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1614_v96), 0))
				_ = _index277
				(_1614_v96).ArraySet1(p0, (_index277).Int())
				var _1615_v97 _dafny.Sequence
				_ = _1615_v97
				_1615_v97 = _dafny.SeqOf(p2)
				var _1616_v98 _dafny.Map
				_ = _1616_v98
				_1616_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1613_v95)
				var _1617_v99 _dafny.Sequence
				_ = _1617_v99
				_1617_v99 = _dafny.SeqOf(_dafny.IntOfInt64(15), _dafny.IntOfInt64(-836))
				var _1618_v100 D15
				_ = _1618_v100
				_1618_v100 = Companion_D15_.Create_DC39_((_1616_v98).Update(_1613_v95, _1613_v95), _1613_v95, _dafny.IntOfUint32((_1617_v99).Cardinality()))
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1614_v96), 0))
				_ = _index278
				var _rhs331 bool = (p1).Contains(p2)
				_ = _rhs331
				var _rhs332 _dafny.Int = (p0).Plus(((_1616_v98).Cardinality()).Minus(p0))
				_ = _rhs332
				var _rhs333 _dafny.Sequence = _1615_v97
				_ = _rhs333
				var _rhs334 bool = (_1618_v100).Dtor_cf57()
				_ = _rhs334
				var _rhs335 bool = _1613_v95
				_ = _rhs335
				var _lhs236 _dafny.Array = _1614_v96
				_ = _lhs236
				var _lhs237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1614_v96), 0))
				_ = _lhs237
				_1613_v95 = _rhs331
				(_lhs236).ArraySet1(_rhs332, (_lhs237).Int())
				_1615_v97 = _rhs333
				_1613_v95 = _rhs334
				_1613_v95 = _rhs335
				var _1619_v101 D14
				_ = _1619_v101
				_1619_v101 = Companion_D14_.Create_DC35_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(7))).Uint32(), func(coer94 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg94 _dafny.Int) interface{} {
						return coer94(arg94)
					}
				}((func(_1620_v96 _dafny.Array, _1621_p2 bool) func(_dafny.Int) _dafny.Map {
					return func(_1622_i8 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1620_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1620_v96), 0))).Int()).(_dafny.Int), _1621_p2)
					}
				})(_1614_v96, p2))))
				var _1623_v102 D14
				_ = _1623_v102
				_1623_v102 = Companion_D14_.Create_DC35_((_1619_v101).Dtor_cf52())
				var _1624_v103 D14
				_ = _1624_v103
				_1624_v103 = Companion_D14_.Create_DC37_(_1623_v102)
				_1624_v103 = _1624_v103
				_1613_v95 = _1613_v95
				var _1625_v104 _dafny.Array
				_ = _1625_v104
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_40
				var _nw235 _dafny.Array
				_ = _nw235
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw235 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) bool = (func(_1626_p1 _dafny.MultiSet) func(_dafny.Int) bool {
						return func(_1627_i9 _dafny.Int) bool {
							return (_1626_p1).IsProperSubsetOf(_1626_p1)
						}
					})(p1)
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw235 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw235).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw235).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1625_v104 = _nw235
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_1625_v104), 0))
				_ = _index279
				(_1625_v104).ArraySet1(p2, (_index279).Int())
				var _1628_v105 _dafny.MultiSet
				_ = _1628_v105
				_1628_v105 = _dafny.MultiSetOf((_1614_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1614_v96), 0))).Int()).(_dafny.Int), (_1617_v99).Select((Companion_Default___.SafeIndex((_1614_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1614_v96), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1617_v99).Cardinality()))).Uint32()).(_dafny.Int))
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_1625_v104), 0))
				_ = _index280
				(_1625_v104).ArraySet1((_1628_v105).IsSubsetOf((_dafny.MultiSetOf((_1614_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1614_v96), 0))).Int()).(_dafny.Int), p0)).Union(Companion_Default___.Fm31((_1614_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1614_v96), 0))).Int()).(_dafny.Int), p0, globalState))), (_index280).Int())
				var _1629_v106 _dafny.Map
				_ = _1629_v106
				_1629_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1617_v99, !((_1625_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_1625_v104), 0))).Int()).(bool)))
				_1629_v106 = _1629_v106
			}
			(globalState).F13 = ((_this).Fm2(p0, globalState)).Minus(p0)
		} else {
			var _1630_v107 _dafny.Set
			_ = _1630_v107
			_1630_v107 = _dafny.SetOf(p2, p2)
			var _1631_v108 _dafny.Map
			_ = _1631_v108
			_1631_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
			var _1632_v109 _dafny.Map
			_ = _1632_v109
			_1632_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _1633_v110 _dafny.Sequence
			_ = _1633_v110
			_1633_v110 = _dafny.SeqOf(!((_this).Fm8((_1630_v107).Cardinality(), (func() bool {
				if (_1631_v108).Contains(true) {
					return (_1631_v108).Get(true).(bool)
				}
				return !(!(p2))
			})(), (func() _dafny.Int {
				if (_1632_v109).Contains(p2) {
					return (_1632_v109).Get(p2).(_dafny.Int)
				}
				return (_dafny.Zero).Minus((_this).Fm2(p0, globalState))
			})(), globalState)), p2)
			var _1634_v111 _dafny.Array
			_ = _1634_v111
			var _nwElement0_58 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), p0))
			_ = _nwElement0_58
			var _nw236 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(24))
			_ = _nw236
			(_nw236).ArraySet1(_nwElement0_58, 0)
			(_nw236).ArraySet1((p0).Minus(_dafny.IntOfUint32((_1633_v110).Cardinality())), 1)
			(_nw236).ArraySet1(p0, 2)
			(_nw236).ArraySet1((_dafny.Zero).Minus(p0), 3)
			(_nw236).ArraySet1(p0, 4)
			(_nw236).ArraySet1(p0, 5)
			(_nw236).ArraySet1(p0, 6)
			(_nw236).ArraySet1(_dafny.IntOfInt64(75), 7)
			(_nw236).ArraySet1(p0, 8)
			(_nw236).ArraySet1(p0, 9)
			(_nw236).ArraySet1(p0, 10)
			(_nw236).ArraySet1(_dafny.IntOfInt64(917), 11)
			(_nw236).ArraySet1(p0, 12)
			(_nw236).ArraySet1(Companion_Default___.SafeDivisionInt((_1630_v107).Cardinality(), p0), 13)
			(_nw236).ArraySet1(_dafny.IntOfInt64(265), 14)
			(_nw236).ArraySet1(p0, 15)
			(_nw236).ArraySet1(p0, 16)
			(_nw236).ArraySet1(p0, 17)
			(_nw236).ArraySet1(p0, 18)
			(_nw236).ArraySet1(p0, 19)
			(_nw236).ArraySet1(p0, 20)
			(_nw236).ArraySet1((p1).Cardinality(), 21)
			(_nw236).ArraySet1((_this).Fm2(Companion_Default___.Fm17(_dafny.IntOfUint32((_1486_v2).Cardinality()), globalState), globalState), 22)
			(_nw236).ArraySet1(p0, 23)
			_1634_v111 = _nw236
			_1634_v111 = _1634_v111
			(globalState).F4 = p0
			var _1635_v112 _dafny.Array
			_ = _1635_v112
			var _len0_41 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_41
			var _nw237 _dafny.Array
			_ = _nw237
			if _len0_41.Cmp(_dafny.Zero) == 0 {
				_nw237 = _dafny.NewArray(_len0_41)
			} else {
				var _init41 func(_dafny.Int) D12 = (func(_1636_v110 _dafny.Sequence, _1637_p0 _dafny.Int) func(_dafny.Int) D12 {
					return func(_1638_i10 _dafny.Int) D12 {
						return Companion_D12_.Create_DC32_(_dafny.SetOf((_1636_v110).Select((Companion_Default___.SafeIndex(_1637_p0, _dafny.IntOfUint32((_1636_v110).Cardinality()))).Uint32()).(bool)))
					}
				})(_1633_v110, p0)
				_ = _init41
				var _element0_41 = _init41(_dafny.Zero)
				_ = _element0_41
				_nw237 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
				(_nw237).ArraySet1(_element0_41, 0)
				var _nativeLen0_41 = (_len0_41).Int()
				_ = _nativeLen0_41
				for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
					(_nw237).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
				}
			}
			_1635_v112 = _nw237
			var _1639_v113 D12
			_ = _1639_v113
			_1639_v113 = Companion_D12_.Create_DC32_(_dafny.SetOf(p2))
			var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_1635_v112), 0))
			_ = _index281
			(_1635_v112).ArraySet1(_1639_v113, (_index281).Int())
			var _1640_v114 _dafny.Array
			_ = _1640_v114
			var _nw238 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(15))
			_ = _nw238
			_1640_v114 = _nw238
			var _1641_v115 D4
			_ = _1641_v115
			_1641_v115 = Companion_D4_.Create_DC11_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(608))).Uint32(), func(coer95 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg95 _dafny.Int) interface{} {
					return coer95(arg95)
				}
			}((func(_1642_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1643_i11 _dafny.Int) _dafny.Sequence {
					return _1642_v2
				}
			})(_1486_v2))))
			var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1640_v114), 0))
			_ = _index282
			(_1640_v114).ArraySet1(_1641_v115, (_index282).Int())
			var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_1635_v112), 0))
			_ = _index283
			var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1640_v114), 0))
			_ = _index284
			var _rhs336 _dafny.Int = (func() _dafny.Int {
				if p2 {
					return p0
				}
				return p0
			})()
			_ = _rhs336
			var _rhs337 _dafny.Set = _1630_v107
			_ = _rhs337
			var _rhs338 D12 = Companion_D12_.Create_DC32_(_1630_v107)
			_ = _rhs338
			var _rhs339 _dafny.Int = p0
			_ = _rhs339
			var _rhs340 D4 = _1641_v115
			_ = _rhs340
			var _lhs238 *GlobalState = globalState
			_ = _lhs238
			var _lhs239 _dafny.Array = _1635_v112
			_ = _lhs239
			var _lhs240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_1635_v112), 0))
			_ = _lhs240
			var _lhs241 *GlobalState = globalState
			_ = _lhs241
			var _lhs242 _dafny.Array = _1640_v114
			_ = _lhs242
			var _lhs243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1640_v114), 0))
			_ = _lhs243
			_lhs238.F4 = _rhs336
			_1630_v107 = _rhs337
			(_lhs239).ArraySet1(_rhs338, (_lhs240).Int())
			_lhs241.F4 = _rhs339
			(_lhs242).ArraySet1(_rhs340, (_lhs243).Int())
			var _1644_v116 _dafny.Sequence
			_ = _1644_v116
			_1644_v116 = _dafny.SeqOf(p0, p0)
			var _1645_v117 _dafny.MultiSet
			_ = _1645_v117
			_1645_v117 = _dafny.MultiSetOf(p0, _dafny.IntOfUint32((_1644_v116).Cardinality()), p0, _dafny.IntOfUint32((_1486_v2).Cardinality()))
			var _1646_v119 _dafny.MultiSet
			_ = _1646_v119
			_1646_v119 = _dafny.MultiSetOf((p0).Times(p0), _dafny.IntOfInt64(-763), (_dafny.Zero).Minus((func() _dafny.Int {
				if (_1645_v117).Contains(p0) {
					return (_1645_v117).Multiplicity(p0)
				}
				return (_dafny.Zero).Minus(p0)
			})()), (func() _dafny.Int {
				if false {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), (func() _dafny.Map {
						var _coll72 = _dafny.NewMapBuilder()
						_ = _coll72
						for _iter74 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(436), _dafny.IntOfInt64(908))); ; {
							_compr_72, _ok74 := _iter74()
							if !_ok74 {
								break
							}
							var _1647_v118 _dafny.Int
							_1647_v118 = interface{}(_compr_72).(_dafny.Int)
							if ((_dafny.IntOfInt64(436)).Cmp(_1647_v118) <= 0) && ((_1647_v118).Cmp(_dafny.IntOfInt64(908)) < 0) {
								_coll72.Add(Companion_Default___.SafeDivisionInt(_1647_v118, p0), p2)
							}
						}
						return _coll72.ToMap()
					}()).Cardinality())).Cardinality()
				}
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(815))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}(func(_1648_i12 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}))).Cardinality())
			})(), p0)
			_1645_v117 = _1645_v117
			if p2 {
				var _1649_v120 _dafny.MultiSet
				_ = _1649_v120
				_1649_v120 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(279))).Uint32(), func(coer97 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}((func(_1650_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1651_i13 _dafny.Int) _dafny.Int {
						return _1650_p0
					}
				})(p0))))
				var _1652_v121 D1
				_ = _1652_v121
				_1652_v121 = Companion_D1_.Create_DC3_(_1486_v2, !(!(true)))
				var _1653_v122 _dafny.Map
				_ = _1653_v122
				_1653_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((((_1649_v120).Update(_1644_v116, Companion_Default___.Abs((_1630_v107).Cardinality()))).Update(_1644_v116, Companion_Default___.Abs(_dafny.IntOfInt64(-51)))).Cardinality()), (_1652_v121).Dtor_cf3())
				_1653_v122 = (func() _dafny.Map {
					if (_1630_v107).IsSubsetOf(_dafny.SetOf(p2)) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1486_v2)
					}
					return (_1653_v122).Merge(_1653_v122)
				})()
				var _1654_v123 bool
				_ = _1654_v123
				_1654_v123 = true
				_1654_v123 = (Companion_Default___.SafeDivisionInt(p0, p0)).Cmp(p0) > 0
				var _1655_v124 _dafny.Array
				_ = _1655_v124
				var _nw239 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw239
				_1655_v124 = _nw239
				var _1656_v125 _dafny.Array
				_ = _1656_v125
				var _nw240 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw240
				_1656_v125 = _nw240
				var _1657_v126 _dafny.Map
				_ = _1657_v126
				_1657_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1656_v125)
				var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1655_v124), 0))
				_ = _index285
				(_1655_v124).ArraySet1((func() _dafny.Array {
					if (_1657_v126).Contains(true) {
						return (_1657_v126).Get(true).(_dafny.Array)
					}
					return _1656_v125
				})(), (_index285).Int())
				var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1655_v124), 0))
				_ = _index286
				(_1655_v124).ArraySet1(_1656_v125, (_index286).Int())
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1655_v124), 0))
				_ = _index287
				(_1655_v124).ArraySet1(_1656_v125, (_index287).Int())
				var _1658_v127 D8
				_ = _1658_v127
				_1658_v127 = Companion_D8_.Create_DC25_((p0).Minus(p0))
				_1658_v127 = _1658_v127
			} else {
				var _1659_v128 _dafny.CodePoint
				_ = _1659_v128
				_1659_v128 = _dafny.CodePoint('x')
				var _1660_v129 _dafny.Array
				_ = _1660_v129
				var _nwElement0_59 _dafny.Sequence = _1486_v2
				_ = _nwElement0_59
				var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(17))
				_ = _nw241
				(_nw241).ArraySet1(_nwElement0_59, 0)
				(_nw241).ArraySet1(_1486_v2, 1)
				(_nw241).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-789))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}(func(_1661_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				})), 2)
				(_nw241).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xxum"), 3)
				(_nw241).ArraySet1(_1486_v2, 4)
				(_nw241).ArraySet1(Companion_Default___.Fm19(p2, !(p2), globalState), 5)
				(_nw241).ArraySet1(_1486_v2, 6)
				(_nw241).ArraySet1(_1486_v2, 7)
				(_nw241).ArraySet1(_1486_v2, 8)
				(_nw241).ArraySet1(_1486_v2, 9)
				(_nw241).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hhvdynhug"), 10)
				(_nw241).ArraySet1(_dafny.Companion_Sequence_.Update(_1486_v2, (Companion_Default___.SafeIndex(Companion_Default___.Fm17((_dafny.Zero).Minus(p0), globalState), _dafny.IntOfUint32((_1486_v2).Cardinality()))).Uint32(), _1659_v128), 11)
				(_nw241).ArraySet1(_1486_v2, 12)
				(_nw241).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg99 _dafny.Int) interface{} {
						return coer99(arg99)
					}
				}(func(_1662_i15 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), 13)
				(_nw241).ArraySet1(_1486_v2, 14)
				(_nw241).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("m"), 15)
				(_nw241).ArraySet1(_1486_v2, 16)
				_1660_v129 = _nw241
				var _1663_v130 _dafny.Map
				_ = _1663_v130
				_1663_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1660_v129, (p0).Cmp(p0) == 0)
				_1663_v130 = (_1663_v130).Update(_1660_v129, p2)
				var _1664_v131 bool
				_ = _1664_v131
				_1664_v131 = true
				var _rhs341 _dafny.Int = p0
				_ = _rhs341
				var _rhs342 bool = (_1664_v131) && ((p1).IsProperSubsetOf(_dafny.MultiSetOf(_1664_v131, p2)))
				_ = _rhs342
				var _rhs343 bool = p2
				_ = _rhs343
				var _lhs244 *GlobalState = globalState
				_ = _lhs244
				_lhs244.F4 = _rhs341
				_1664_v131 = _rhs342
				_1664_v131 = _rhs343
				var _1665_v132 _dafny.Array
				_ = _1665_v132
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_42
				var _nw242 _dafny.Array
				_ = _nw242
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw242 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) bool = (func(_1666_v131 bool) func(_dafny.Int) bool {
						return func(_1667_i16 _dafny.Int) bool {
							return _1666_v131
						}
					})(_1664_v131)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw242 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw242).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw242).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1665_v132 = _nw242
				var _1668_v133 _dafny.Map
				_ = _1668_v133
				_1668_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1664_v131), _1631_v108)
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1665_v132), 0))
				_ = _index288
				(_1665_v132).ArraySet1((_this).Fm1(_1668_v133, _dafny.IntOfUint32((_1633_v110).Cardinality()), globalState), (_index288).Int())
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1665_v132), 0))
				_ = _index289
				(_1665_v132).ArraySet1(false, (_index289).Int())
				var _1669_v134 *C7
				_ = _1669_v134
				var _nw243 *C7 = New_C7_()
				_ = _nw243
				_nw243.Ctor__()
				_1669_v134 = _nw243
				var _1670_v135 _dafny.Map
				_ = _1670_v135
				_1670_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1665_v132, _1665_v132)
				var _1671_v136 _dafny.Map
				_ = _1671_v136
				_1671_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1665_v132, _1670_v135)
				_1670_v135 = (func() _dafny.Map {
					if (_1671_v136).Contains(_1665_v132) {
						return (_1671_v136).Get(_1665_v132).(_dafny.Map)
					}
					return _1670_v135
				})()
			}
		}
		var _1672_v137 _dafny.Array
		_ = _1672_v137
		var _nw244 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw244
		_1672_v137 = _nw244
		var _1673_v138 _dafny.Map
		_ = _1673_v138
		_1673_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v137, p2)
		_1673_v138 = (_1673_v138).Update(_1672_v137, p2)
		var _hi11 _dafny.Int = p0
		_ = _hi11
		for _1674_i17 := Companion_Default___.SafeDivisionInt(p0, p0); _1674_i17.Cmp(_hi11) < 0; _1674_i17 = _1674_i17.Plus(_dafny.One) {
			_1672_v137 = _1672_v137
			var _1675_v139 _dafny.Array
			_ = _1675_v139
			var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw245
			_1675_v139 = _nw245
			var _1676_v140 _dafny.Array
			_ = _1676_v140
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_43
			var _nw246 _dafny.Array
			_ = _nw246
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw246 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.Int = (func(_1677_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1678_i18 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1678_i18, _1677_p0)
					}
				})(p0)
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw246 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw246).ArraySet1(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw246).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1676_v140 = _nw246
			var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1675_v139), 0))
			_ = _index290
			(_1675_v139).ArraySet1(_1676_v140, (_index290).Int())
			var _1679_v141 _dafny.MultiSet
			_ = _1679_v141
			_1679_v141 = _dafny.MultiSetOf((func() _dafny.Int {
				if p2 {
					return _1674_i17
				}
				return _dafny.IntOfInt64(-605)
			})(), (_1674_i17).Times(_1674_i17))
			var _1680_v142 _dafny.Sequence
			_ = _1680_v142
			_1680_v142 = _dafny.SeqOf(_1679_v141, _1679_v141)
			var _1681_v143 T0
			_ = _1681_v143
			var _nw247 *C7 = New_C7_()
			_ = _nw247
			_nw247.Ctor__()
			_1681_v143 = _nw247
			var _1682_v144 _dafny.Sequence
			_ = _1682_v144
			_1682_v144 = _dafny.SeqOf(((_1680_v142).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1680_v142).Cardinality()))).Uint32()).(_dafny.MultiSet)).Update(_1674_i17, Companion_Default___.Abs((Companion_D17_.Create_DC47_(_1681_v143, true, _1674_i17, _1674_i17, Companion_Default___.Fm17(_1674_i17, globalState))).Dtor_cf72())))
			var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1675_v139), 0))
			_ = _index291
			var _rhs344 _dafny.Array = _1676_v140
			_ = _rhs344
			var _rhs345 _dafny.Int = (p1).Cardinality()
			_ = _rhs345
			var _rhs346 _dafny.MultiSet = (_1682_v144).Select((Companion_Default___.SafeIndex(_1674_i17, _dafny.IntOfUint32((_1682_v144).Cardinality()))).Uint32()).(_dafny.MultiSet)
			_ = _rhs346
			var _lhs245 _dafny.Array = _1675_v139
			_ = _lhs245
			var _lhs246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1675_v139), 0))
			_ = _lhs246
			var _lhs247 *GlobalState = globalState
			_ = _lhs247
			(_lhs245).ArraySet1(_rhs344, (_lhs246).Int())
			_lhs247.F4 = _rhs345
			_1679_v141 = _rhs346
			if p2 {
				var _1683_v145 bool
				_ = _1683_v145
				_1683_v145 = true
				var _1684_v146 _dafny.Map
				_ = _1684_v146
				_1684_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1681_v143).Fm2(p0, globalState))
				var _1685_v147 _dafny.Set
				_ = _1685_v147
				_1685_v147 = _dafny.SetOf(_1684_v146, _1684_v146)
				_1683_v145 = (_1685_v147).IsSubsetOf(_1685_v147)
				(globalState).F15 = ((_dafny.MultiSetOf(_dafny.IntOfInt64(-81))).Cardinality()).Minus(Companion_Default___.SafeModuloInt((_1679_v141).Cardinality(), _1674_i17))
				_1683_v145 = _1683_v145
				var _1686_v148 _dafny.MultiSet
				_ = _1686_v148
				_1686_v148 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1687_i17 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1688_i19 _dafny.Int) _dafny.Int {
						return _1687_i17
					}
				})(_1674_i17))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(285))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}((func(_1689_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1690_i20 _dafny.Int) _dafny.Int {
						return _1689_p0
					}
				})(p0))))
				var _1691_v149 _dafny.Sequence
				_ = _1691_v149
				_1691_v149 = _dafny.SeqOf(p0)
				var _1692_v150 _dafny.CodePoint
				_ = _1692_v150
				_1692_v150 = _dafny.CodePoint('y')
				var _1693_v151 _dafny.Set
				_ = _1693_v151
				_1693_v151 = _dafny.SetOf((_this).Fm9(_1674_i17, globalState))
				var _rhs347 bool = p2
				_ = _rhs347
				var _rhs348 _dafny.MultiSet = _dafny.MultiSetOf(_1691_v149, Companion_Default___.Fm39(_1674_i17, _1692_v150, _1691_v149, _1683_v145, globalState))
				_ = _rhs348
				var _rhs349 bool = (_dafny.IntOfInt64(751)).Cmp((_1693_v151).Cardinality()) != 0
				_ = _rhs349
				_1683_v145 = _rhs347
				_1686_v148 = _rhs348
				_1683_v145 = _rhs349
				var _1694_v152 _dafny.Map
				_ = _1694_v152
				_1694_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1683_v145, p0)
				var _1695_v153 _dafny.Array
				_ = _1695_v153
				var _nwElement0_60 _dafny.Map = _1694_v152
				_ = _nwElement0_60
				var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(13))
				_ = _nw248
				(_nw248).ArraySet1(_nwElement0_60, 0)
				(_nw248).ArraySet1(_1694_v152, 1)
				(_nw248).ArraySet1(_1694_v152, 2)
				(_nw248).ArraySet1((_1694_v152).Update(_1683_v145, _1674_i17), 3)
				(_nw248).ArraySet1(_1694_v152, 4)
				(_nw248).ArraySet1(_1694_v152, 5)
				(_nw248).ArraySet1(_1694_v152, 6)
				(_nw248).ArraySet1(Companion_Default___.Fm44(_1674_i17, _1683_v145, globalState), 7)
				(_nw248).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1683_v145, p0), 8)
				(_nw248).ArraySet1(_1694_v152, 9)
				(_nw248).ArraySet1(_1694_v152, 10)
				(_nw248).ArraySet1(_1694_v152, 11)
				(_nw248).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0), 12)
				_1695_v153 = _nw248
				var _1696_v154 _dafny.Map
				_ = _1696_v154
				_1696_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1695_v153, p2)
				_1683_v145 = (func() bool {
					if (_1696_v154).Contains(_1695_v153) {
						return (_1696_v154).Get(_1695_v153).(bool)
					}
					return (func() bool {
						if p2 {
							return false
						}
						return false
					})()
				})()
			} else {
				var _1697_v155 D1
				_ = _1697_v155
				_1697_v155 = Companion_D1_.Create_DC4_(_1679_v141)
				_1679_v141 = (_1679_v141).Intersection((_1697_v155).Dtor_cf5())
				var _1698_v156 *C8
				_ = _1698_v156
				var _nw249 *C8 = New_C8_()
				_ = _nw249
				_nw249.Ctor__()
				_1698_v156 = _nw249
				var _1699_v157 _dafny.Set
				_ = _1699_v157
				_1699_v157 = _dafny.SetOf(p2, true, p2)
				var _1700_v158 _dafny.Map
				_ = _1700_v158
				_1700_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1699_v157)
				var _1701_v159 _dafny.MultiSet
				_ = _1701_v159
				_1701_v159 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg102 _dafny.Int) interface{} {
						return coer102(arg102)
					}
				}(func(_1702_i21 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})))
				_1700_v158 = (_1700_v158).Update(!(_1701_v159).Equals(_1701_v159), _1699_v157)
				var _1703_v160 D1
				_ = _1703_v160
				_1703_v160 = Companion_D1_.Create_DC2_(p0)
				(globalState).F13 = (Companion_Default___.SafeModuloInt(p0, p0)).Minus((_1703_v160).Dtor_cf2())
				_1699_v157 = _dafny.SetOf(p2, p2, p2, p2, p2)
			}
			var _1704_v161 _dafny.CodePoint
			_ = _1704_v161
			_1704_v161 = _dafny.CodePoint('u')
			var _1705_v162 _dafny.Map
			_ = _1705_v162
			_1705_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v137, _1704_v161)
			var _1706_v163 _dafny.Map
			_ = _1706_v163
			_1706_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Minus(p0), (func() _dafny.CodePoint {
				if (_1705_v162).Contains(_1672_v137) {
					return (_1705_v162).Get(_1672_v137).(_dafny.CodePoint)
				}
				return _1704_v161
			})())
			_1706_v163 = (_1706_v163).Update(_dafny.IntOfInt64(115), _1704_v161)
		}
		var _1707_v164 _dafny.Array
		_ = _1707_v164
		var _len0_44 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_44
		var _nw250 _dafny.Array
		_ = _nw250
		if _len0_44.Cmp(_dafny.Zero) == 0 {
			_nw250 = _dafny.NewArray(_len0_44)
		} else {
			var _init44 func(_dafny.Int) _dafny.Int = func(_1708_i22 _dafny.Int) _dafny.Int {
				return (_1708_i22).Plus(_dafny.IntOfInt64(871))
			}
			_ = _init44
			var _element0_44 = _init44(_dafny.Zero)
			_ = _element0_44
			_nw250 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
			(_nw250).ArraySet1(_element0_44, 0)
			var _nativeLen0_44 = (_len0_44).Int()
			_ = _nativeLen0_44
			for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
				(_nw250).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
			}
		}
		_1707_v164 = _nw250
		_1707_v164 = _1707_v164
		var _hi12 _dafny.Int = p0
		_ = _hi12
		for _1709_i23 := p0; _1709_i23.Cmp(_hi12) < 0; _1709_i23 = _1709_i23.Plus(_dafny.One) {
			var _1710_v165 _dafny.Array
			_ = _1710_v165
			var _nw251 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
			_ = _nw251
			_1710_v165 = _nw251
			var _nw252 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
			_ = _nw252
			_1710_v165 = _nw252
			var _1711_v166 T0
			_ = _1711_v166
			var _nw253 *C7 = New_C7_()
			_ = _nw253
			_nw253.Ctor__()
			_1711_v166 = _nw253
			_1711_v166 = _1711_v166
			var _1712_v167 bool
			_ = _1712_v167
			_1712_v167 = false
			_1712_v167 = _1712_v167
			_1712_v167 = _1712_v167
		}
	}
}
func (_this *C9) M7(p0 _dafny.Int, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.MultiSet) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r1
		var _1713_v0 _dafny.Array
		_ = _1713_v0
		var _nw254 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw254
		_1713_v0 = _nw254
		var _1714_v1 _dafny.MultiSet
		_ = _1714_v1
		_1714_v1 = _dafny.MultiSetOf(_1713_v0, _1713_v0)
		r0 = (func() _dafny.Int {
			if (_1714_v1).Contains(_1713_v0) {
				return (_1714_v1).Multiplicity(_1713_v0)
			}
			return p0
		})()
		var _1715_i0 _dafny.Int
		_ = _1715_i0
		_1715_i0 = _dafny.Zero
		{
			for p1 {
				{
					if (_1715_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1715_i0 = (_1715_i0).Plus(_dafny.One)
					if true {
						var _1716_v2 _dafny.Array
						_ = _1716_v2
						var _len0_45 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_45
						var _nw255 _dafny.Array
						_ = _nw255
						if _len0_45.Cmp(_dafny.Zero) == 0 {
							_nw255 = _dafny.NewArray(_len0_45)
						} else {
							var _init45 func(_dafny.Int) D3 = func(_1717_i1 _dafny.Int) D3 {
								return Companion_D3_.Create_DC10_(_dafny.IntOfInt64(65), _dafny.IntOfInt64(626))
							}
							_ = _init45
							var _element0_45 = _init45(_dafny.Zero)
							_ = _element0_45
							_nw255 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
							(_nw255).ArraySet1(_element0_45, 0)
							var _nativeLen0_45 = (_len0_45).Int()
							_ = _nativeLen0_45
							for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
								(_nw255).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
							}
						}
						_1716_v2 = _nw255
						var _1718_v3 _dafny.MultiSet
						_ = _1718_v3
						_1718_v3 = _dafny.MultiSetOf(p1)
						var _1719_v4 _dafny.Sequence
						_ = _1719_v4
						_1719_v4 = _dafny.SeqOf(p1, p1)
						var _1720_v5 _dafny.Map
						_ = _1720_v5
						_1720_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1719_v4).Cardinality()), p0)
						var _1721_v6 _dafny.Sequence
						_ = _1721_v6
						_1721_v6 = _dafny.SeqOf(_1720_v5)
						var _1722_v7 _dafny.Sequence
						_ = _1722_v7
						_1722_v7 = _dafny.SeqOf(((_1718_v3).Update(p1, Companion_Default___.Abs(p0))).Cardinality(), (_this).Fm2(p0, globalState), _dafny.IntOfUint32((_1721_v6).Cardinality()))
						var _1723_v8 D3
						_ = _1723_v8
						_1723_v8 = Companion_D3_.Create_DC10_(_dafny.IntOfUint32((_1722_v7).Cardinality()), p0)
						var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_1716_v2), 0))
						_ = _index292
						(_1716_v2).ArraySet1(_1723_v8, (_index292).Int())
						var _1724_v9 _dafny.Set
						_ = _1724_v9
						_1724_v9 = _dafny.SetOf(!(p1))
						var _1725_v10 _dafny.Set
						_ = _1725_v10
						_1725_v10 = _dafny.SetOf(_1724_v9, _1724_v9)
						var _pat_let_tv32 = _1720_v5
						_ = _pat_let_tv32
						var _pat_let_tv33 = _1725_v10
						_ = _pat_let_tv33
						var _pat_let_tv34 = _1720_v5
						_ = _pat_let_tv34
						var _pat_let_tv35 = _1725_v10
						_ = _pat_let_tv35
						var _pat_let_tv36 = p0
						_ = _pat_let_tv36
						var _pat_let_tv37 = p0
						_ = _pat_let_tv37
						var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_1716_v2), 0))
						_ = _index293
						(_1716_v2).ArraySet1(func(_pat_let24_0 D3) D3 {
							return func(_1726_dt__update__tmp_h0 D3) D3 {
								return func(_pat_let25_0 _dafny.Int) D3 {
									return func(_1727_dt__update_hcf14_h0 _dafny.Int) D3 {
										return Companion_D3_.Create_DC10_((_1726_dt__update__tmp_h0).Dtor_cf13(), _1727_dt__update_hcf14_h0)
									}(_pat_let25_0)
								}(Companion_Default___.SafeModuloInt((func() _dafny.Int {
									if (_pat_let_tv32).Contains((_pat_let_tv33).Cardinality()) {
										return (_pat_let_tv34).Get((_pat_let_tv35).Cardinality()).(_dafny.Int)
									}
									return _pat_let_tv36
								})(), _pat_let_tv37))
							}(_pat_let24_0)
						}(_1723_v8), (_index293).Int())
						var _1728_v11 *C5
						_ = _1728_v11
						var _nw256 *C5 = New_C5_()
						_ = _nw256
						_nw256.Ctor__()
						_1728_v11 = _nw256
						(globalState).F15 = Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus((func() _dafny.Int {
							if p1 {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1728_v11)).Cardinality()
							}
							return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tmqug")).Cardinality())
						})()))
						var _1729_v12 D18
						_ = _1729_v12
						_1729_v12 = Companion_D18_.Create_DC48_(_1720_v5)
						var _1730_v13 _dafny.Map
						_ = _1730_v13
						_1730_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
						var _1731_v14 _dafny.Map
						_ = _1731_v14
						_1731_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1730_v13)
						var _pat_let_tv38 = _1720_v5
						_ = _pat_let_tv38
						var _1732_v15 _dafny.Map
						_ = _1732_v15
						_1732_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p0, p0, (_dafny.Zero).Minus(((func(_pat_let26_0 D18) D18 {
							return func(_1733_dt__update__tmp_h1 D18) D18 {
								return func(_pat_let27_0 _dafny.Map) D18 {
									return func(_1734_dt__update_hcf74_h0 _dafny.Map) D18 {
										return Companion_D18_.Create_DC48_(_1734_dt__update_hcf74_h0)
									}(_pat_let27_0)
								}(_pat_let_tv38)
							}(_pat_let26_0)
						}(_1729_v12)).Dtor_cf74()).Cardinality()))).Cardinality(), (_this).Fm1(_1731_v14, p0, globalState))
						var _1735_v17 _dafny.Set
						_ = _1735_v17
						_1735_v17 = _dafny.SetOf((_dafny.Zero).Minus(p0), p0)
						var _1736_v18 _dafny.Sequence
						_ = _1736_v18
						_1736_v18 = _dafny.SeqOf(func() _dafny.Set {
							var _coll73 = _dafny.NewBuilder()
							_ = _coll73
							for _iter75 := _dafny.Iterate((_1732_v15).Keys().Elements()); ; {
								_compr_73, _ok75 := _iter75()
								if !_ok75 {
									break
								}
								var _1737_v16 _dafny.Int
								_1737_v16 = interface{}(_compr_73).(_dafny.Int)
								if (_1732_v15).Contains(_1737_v16) {
									_coll73.Add((_1737_v16).Minus(_dafny.IntOfInt64(-68)))
								}
							}
							return _coll73.ToSet()
						}(), _1735_v17)
						var _1738_v19 _dafny.Map
						_ = _1738_v19
						_1738_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1735_v17, _dafny.SetOf(p0, _dafny.IntOfUint32((_1722_v7).Cardinality())))
						var _1739_v21 _dafny.Sequence
						_ = _1739_v21
						_1739_v21 = _dafny.SeqOf((_1736_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.IntOfUint32((_1736_v18).Cardinality()))).Uint32()).(_dafny.Set), (func() _dafny.Set {
							if (_1738_v19).Contains(func() _dafny.Set {
								var _coll74 = _dafny.NewBuilder()
								_ = _coll74
								for _iter76 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(594), _dafny.IntOfInt64(392))); ; {
									_compr_74, _ok76 := _iter76()
									if !_ok76 {
										break
									}
									var _1740_v20 _dafny.Int
									_1740_v20 = interface{}(_compr_74).(_dafny.Int)
									if ((_dafny.IntOfInt64(594)).Cmp(_1740_v20) <= 0) && ((_1740_v20).Cmp(_dafny.IntOfInt64(392)) < 0) {
										_coll74.Add((_1740_v20).Plus(p0))
									}
								}
								return _coll74.ToSet()
							}()) {
								return (_1738_v19).Get(func() _dafny.Set {
									var _coll75 = _dafny.NewBuilder()
									_ = _coll75
									for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(594), _dafny.IntOfInt64(392))); ; {
										_compr_75, _ok77 := _iter77()
										if !_ok77 {
											break
										}
										var _1741_v20 _dafny.Int
										_1741_v20 = interface{}(_compr_75).(_dafny.Int)
										if ((_dafny.IntOfInt64(594)).Cmp(_1741_v20) <= 0) && ((_1741_v20).Cmp(_dafny.IntOfInt64(392)) < 0) {
											_coll75.Add((_1741_v20).Plus(p0))
										}
									}
									return _coll75.ToSet()
								}()).(_dafny.Set)
							}
							return _1735_v17
						})(), _dafny.SetOf((_dafny.Zero).Minus(p0), p0, p0), _1735_v17, _1735_v17)
						(globalState).F13 = _dafny.IntOfUint32((_1739_v21).Cardinality())
						var _1742_v22 _dafny.Array
						_ = _1742_v22
						var _nw257 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
						_ = _nw257
						_1742_v22 = _nw257
						var _1743_v23 _dafny.Map
						_ = _1743_v23
						_1743_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1742_v22, p0)
						var _1744_v24 _dafny.Sequence
						_ = _1744_v24
						_1744_v24 = _dafny.UnicodeSeqOfUtf8Bytes("kxs")
						var _1745_v25 _dafny.Map
						_ = _1745_v25
						_1745_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1744_v24, _1743_v23)
						var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1713_v0), 0))
						_ = _index294
						(_1713_v0).ArraySet1(!(_1743_v23).Equals((func() _dafny.Map {
							if (_1745_v25).Contains(_dafny.UnicodeSeqOfUtf8Bytes("idknhjqof")) {
								return (_1745_v25).Get(_dafny.UnicodeSeqOfUtf8Bytes("idknhjqof")).(_dafny.Map)
							}
							return _1743_v23
						})()), (_index294).Int())
						var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1713_v0), 0))
						_ = _index295
						(_1713_v0).ArraySet1((_1728_v11).Fm20(globalState), (_index295).Int())
						var _1746_v26 D2
						_ = _1746_v26
						_1746_v26 = Companion_D2_.Create_DC6_(p0)
						var _1747_v27 _dafny.MultiSet
						_ = _1747_v27
						_1747_v27 = _dafny.MultiSetOf(_1746_v26, _1746_v26, _1746_v26)
						var _1748_v28 _dafny.MultiSet
						_ = _1748_v28
						_1748_v28 = _dafny.MultiSetOf((_1747_v27).Update(_1746_v26, Companion_Default___.Abs(p0)), _1747_v27, _1747_v27, _1747_v27, _1747_v27)
						var _1749_v29 _dafny.CodePoint
						_ = _1749_v29
						_1749_v29 = _dafny.CodePoint('f')
						var _1750_v30 _dafny.Map
						_ = _1750_v30
						_1750_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1713_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1713_v0), 0))).Int()).(bool), _1749_v29)
						var _1751_v31 _dafny.Sequence
						_ = _1751_v31
						_1751_v31 = _dafny.SeqOf(_1747_v27, (_1747_v27).Update(_1746_v26, Companion_Default___.Abs(p0)))
						var _1752_v32 _dafny.Map
						_ = _1752_v32
						_1752_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.MultiSetFromSeq(_1751_v31))
						var _1753_v33 _dafny.Array
						_ = _1753_v33
						var _nwElement0_61 _dafny.MultiSet = _1748_v28
						_ = _nwElement0_61
						var _nw258 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(9))
						_ = _nw258
						(_nw258).ArraySet1(_nwElement0_61, 0)
						(_nw258).ArraySet1((func() _dafny.MultiSet {
							if p1 {
								return _1748_v28
							}
							return _1748_v28
						})(), 1)
						(_nw258).ArraySet1((_1748_v28).Intersection(_1748_v28), 2)
						(_nw258).ArraySet1(Companion_Default___.Fm50((_1713_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1713_v0), 0))).Int()).(bool), (_this).Fm10(_1750_v30, _dafny.IntOfInt64(790), _1720_v5, globalState), p0, (func() bool {
							if (_1732_v15).Contains(p0) {
								return (_1732_v15).Get(p0).(bool)
							}
							return p1
						})(), globalState), 3)
						(_nw258).ArraySet1((func() _dafny.MultiSet {
							if (_1752_v32).Contains(p0) {
								return (_1752_v32).Get(p0).(_dafny.MultiSet)
							}
							return _1748_v28
						})(), 4)
						(_nw258).ArraySet1(_1748_v28, 5)
						(_nw258).ArraySet1(_1748_v28, 6)
						(_nw258).ArraySet1(_1748_v28, 7)
						(_nw258).ArraySet1((Companion_Default___.Fm50((_1713_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_1713_v0), 0))).Int()).(bool), (_this).Fm9(p0, globalState), p0, p1, globalState)).Union(_1748_v28), 8)
						_1753_v33 = _nw258
						var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_1753_v33), 0))
						_ = _index296
						(_1753_v33).ArraySet1(_1748_v28, (_index296).Int())
						var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_1753_v33), 0))
						_ = _index297
						(_1753_v33).ArraySet1(_1748_v28, (_index297).Int())
					} else {
						var _1754_v34 bool
						_ = _1754_v34
						_1754_v34 = true
						var _1755_v35 _dafny.MultiSet
						_ = _1755_v35
						_1755_v35 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(842), (_dafny.Zero).Minus(p0))
						var _1756_v36 _dafny.Sequence
						_ = _1756_v36
						_1756_v36 = _dafny.UnicodeSeqOfUtf8Bytes("l")
						var _rhs350 _dafny.Int = (_dafny.Zero).Minus(p0)
						_ = _rhs350
						var _rhs351 bool = _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
							if p1 {
								return _1756_v36
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(506))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg103 _dafny.Int) interface{} {
									return coer103(arg103)
								}
							}(func(_1757_i2 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('b')
							}))
						})(), _1756_v36)
						_ = _rhs351
						var _rhs352 _dafny.MultiSet = ((_1755_v35).Intersection(_1755_v35)).Union(_1755_v35)
						_ = _rhs352
						var _lhs248 *GlobalState = globalState
						_ = _lhs248
						_lhs248.F13 = _rhs350
						_1754_v34 = _rhs351
						_1755_v35 = _rhs352
						var _1758_v37 _dafny.Set
						_ = _1758_v37
						_1758_v37 = _dafny.SetOf(p0, p0, p0, p0, p0)
						var _1759_v38 _dafny.MultiSet
						_ = _1759_v38
						_1759_v38 = _dafny.MultiSetOf(_1758_v37)
						var _1760_v39 D5
						_ = _1760_v39
						_1760_v39 = Companion_D5_.Create_DC13_(_1759_v38)
						var _1761_v40 _dafny.MultiSet
						_ = _1761_v40
						_1761_v40 = _dafny.MultiSetOf(_1760_v39)
						(globalState).F4 = (_1761_v40).Cardinality()
						var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1713_v0), 0))
						_ = _index298
						(_1713_v0).ArraySet1(p1, (_index298).Int())
						var _1762_v41 _dafny.Map
						_ = _1762_v41
						_1762_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), p1)
						var _1763_v42 _dafny.Map
						_ = _1763_v42
						_1763_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1754_v34, _1762_v41)
						var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1713_v0), 0))
						_ = _index299
						var _rhs353 _dafny.Array = _1713_v0
						_ = _rhs353
						var _rhs354 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((p0).Times((_dafny.Zero).Minus(p0)), Companion_Default___.Fm17(p0, globalState)))
						_ = _rhs354
						var _rhs355 bool = (_this).Fm1(_1763_v42, p0, globalState)
						_ = _rhs355
						var _lhs249 *GlobalState = globalState
						_ = _lhs249
						var _lhs250 _dafny.Array = _1713_v0
						_ = _lhs250
						var _lhs251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1713_v0), 0))
						_ = _lhs251
						_1713_v0 = _rhs353
						_lhs249.F13 = _rhs354
						(_lhs250).ArraySet1(_rhs355, (_lhs251).Int())
						_1754_v34 = (_1758_v37).IsDisjointFrom((func() _dafny.Set {
							if false {
								return _1758_v37
							}
							return _dafny.SetOf(p0)
						})())
						var _1764_v43 _dafny.Map
						_ = _1764_v43
						_1764_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(670), _1756_v36)
						var _1765_v44 _dafny.Map
						_ = _1765_v44
						_1765_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1764_v43).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1756_v36)))
						_1765_v44 = (_1765_v44).Update((_1713_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1713_v0), 0))).Int()).(bool), func() _dafny.Map {
							var _coll76 = _dafny.NewMapBuilder()
							_ = _coll76
							for _iter78 := _dafny.Iterate((_1755_v35).Elements()); ; {
								_compr_76, _ok78 := _iter78()
								if !_ok78 {
									break
								}
								var _1766_v45 _dafny.Int
								_1766_v45 = interface{}(_compr_76).(_dafny.Int)
								if (_1755_v35).Contains(_1766_v45) {
									_coll76.Add((_1766_v45).Minus(p0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(124))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg104 _dafny.Int) interface{} {
											return coer104(arg104)
										}
									}(func(_1767_i3 _dafny.Int) _dafny.CodePoint {
										return _dafny.CodePoint('u')
									})))
								}
							}
							return _coll76.ToMap()
						}())
					}
					(globalState).F15 = Companion_Default___.SafeModuloInt(p0, p0)
					(globalState).F15 = Companion_Default___.Fm17(p0, globalState)
					var _1768_v46 bool
					_ = _1768_v46
					_1768_v46 = true
					var _1769_v47 _dafny.Sequence
					_ = _1769_v47
					_1769_v47 = _dafny.UnicodeSeqOfUtf8Bytes("yfrlx")
					var _1770_v48 _dafny.MultiSet
					_ = _1770_v48
					_1770_v48 = _dafny.MultiSetOf(p1)
					var _rhs356 bool = _1768_v46
					_ = _rhs356
					var _rhs357 bool = p1
					_ = _rhs357
					var _rhs358 _dafny.Sequence = _1769_v47
					_ = _rhs358
					var _rhs359 _dafny.MultiSet = _1770_v48
					_ = _rhs359
					_1768_v46 = _rhs356
					_1768_v46 = _rhs357
					_1769_v47 = _rhs358
					_1770_v48 = _rhs359
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		r0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(p0))
		var _1771_v49 _dafny.Set
		_ = _1771_v49
		_1771_v49 = _dafny.SetOf(p1, p1)
		var _1772_v50 _dafny.Set
		_ = _1772_v50
		_1772_v50 = _dafny.SetOf((_1771_v49).Cardinality())
		var _1773_v51 D11
		_ = _1773_v51
		_1773_v51 = Companion_D11_.Create_DC28_(_1772_v50)
		var _1774_v52 _dafny.Sequence
		_ = _1774_v52
		_1774_v52 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("holgkicqv")).Cardinality()), p0)
		var _1775_v53 _dafny.Sequence
		_ = _1775_v53
		_1775_v53 = _dafny.SeqOf(Companion_Default___.Fm51(_1774_v52, p0, globalState))
		var _1776_v54 D11
		_ = _1776_v54
		_1776_v54 = Companion_D11_.Create_DC31_(Companion_D11_.Create_DC29_())
		var _1777_v55 D19
		_ = _1777_v55
		_1777_v55 = Companion_D19_.Create_DC51_(_1775_v53)
		var _1778_v56 _dafny.Sequence
		_ = _1778_v56
		_1778_v56 = _dafny.SeqOf(_1775_v53, _1775_v53)
		var _pat_let_tv39 = _1773_v51
		_ = _pat_let_tv39
		var _1779_v57 _dafny.Array
		_ = _1779_v57
		var _nwElement0_62 _dafny.Sequence = _dafny.SeqOf(Companion_D11_.Create_DC31_(_1773_v51))
		_ = _nwElement0_62
		var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(21))
		_ = _nw259
		(_nw259).ArraySet1(_nwElement0_62, 0)
		(_nw259).ArraySet1(_1775_v53, 1)
		(_nw259).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D11_.Create_DC31_(Companion_D11_.Create_DC31_(_1773_v51)), _1776_v54, _1776_v54), _1775_v53), 2)
		(_nw259).ArraySet1(_1775_v53, 3)
		(_nw259).ArraySet1(_dafny.SeqOf(_1776_v54), 4)
		(_nw259).ArraySet1(_dafny.SeqOf(func(_pat_let28_0 D11) D11 {
			return func(_1780_dt__update__tmp_h2 D11) D11 {
				return func(_pat_let29_0 D11) D11 {
					return func(_1781_dt__update_hcf48_h0 D11) D11 {
						return Companion_D11_.Create_DC31_(_1781_dt__update_hcf48_h0)
					}(_pat_let29_0)
				}(_pat_let_tv39)
			}(_pat_let28_0)
		}(_1776_v54)), 5)
		(_nw259).ArraySet1(_1775_v53, 6)
		(_nw259).ArraySet1((func() _dafny.Sequence {
			if p1 {
				return _1775_v53
			}
			return (_1777_v55).Dtor_cf80()
		})(), 7)
		(_nw259).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1776_v54), _1775_v53), 8)
		(_nw259).ArraySet1((_1778_v56).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1778_v56).Cardinality()))).Uint32()).(_dafny.Sequence), 9)
		(_nw259).ArraySet1(_1775_v53, 10)
		(_nw259).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm52(globalState), _1775_v53), 11)
		(_nw259).ArraySet1(_1775_v53, 12)
		(_nw259).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer105 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
			return func(arg105 _dafny.Int) interface{} {
				return coer105(arg105)
			}
		}((func(_1782_v54 D11) func(_dafny.Int) D11 {
			return func(_1783_i4 _dafny.Int) D11 {
				return _1782_v54
			}
		})(_1776_v54))), 13)
		(_nw259).ArraySet1(_1775_v53, 14)
		(_nw259).ArraySet1(_1775_v53, 15)
		(_nw259).ArraySet1(_1775_v53, 16)
		(_nw259).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-243))).Uint32(), func(coer106 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
			return func(arg106 _dafny.Int) interface{} {
				return coer106(arg106)
			}
		}((func(_1784_v54 D11) func(_dafny.Int) D11 {
			return func(_1785_i5 _dafny.Int) D11 {
				return _1784_v54
			}
		})(_1776_v54))), 17)
		(_nw259).ArraySet1(_1775_v53, 18)
		(_nw259).ArraySet1(_1775_v53, 19)
		(_nw259).ArraySet1(_dafny.Companion_Sequence_.Update(_1775_v53, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1775_v53).Cardinality()))).Uint32(), Companion_Default___.Fm51(_1774_v52, p0, globalState)), 20)
		_1779_v57 = _nw259
		var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_1779_v57), 0))
		_ = _index300
		(_1779_v57).ArraySet1(_dafny.SeqOf(Companion_D11_.Create_DC31_(_1773_v51), _1776_v54), (_index300).Int())
		var _1786_v58 bool
		_ = _1786_v58
		_1786_v58 = true
		var _1787_v59 D0
		_ = _1787_v59
		_1787_v59 = Companion_D0_.Create_DC0_(_1713_v0)
		var _1788_v60 _dafny.Map
		_ = _1788_v60
		_1788_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
		var _1789_v61 _dafny.Sequence
		_ = _1789_v61
		_1789_v61 = _dafny.SeqOf(_1786_v58, (func() bool {
			if (_1788_v60).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1774_v52, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.IntOfUint32((_1774_v52).Cardinality()))).Uint32(), p0)).Cardinality())) {
				return (_1788_v60).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1774_v52, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.IntOfUint32((_1774_v52).Cardinality()))).Uint32(), p0)).Cardinality())).(bool)
			}
			return _1786_v58
		})())
		var _1790_v62 _dafny.Sequence
		_ = _1790_v62
		_1790_v62 = _dafny.UnicodeSeqOfUtf8Bytes("ksdbdq")
		var _1791_v63 _dafny.Sequence
		_ = _1791_v63
		_1791_v63 = _dafny.SeqOf(_1790_v62, _dafny.UnicodeSeqOfUtf8Bytes("mhqnf"), Companion_Default___.Fm30(p0, _1786_v58, _1772_v50, true, globalState))
		var _1792_v64 D7
		_ = _1792_v64
		_1792_v64 = Companion_D7_.Create_DC21_(p0, p0, false, _1786_v58, p1)
		var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_1779_v57), 0))
		_ = _index301
		var _rhs360 _dafny.Array = ((func() D0 {
			if _1786_v58 {
				return _1787_v59
			}
			return _1787_v59
		})()).Dtor_cf0()
		_ = _rhs360
		var _rhs361 _dafny.Sequence = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm52(globalState), (Companion_Default___.SafeIndex((_dafny.IntOfUint32((_1789_v61).Cardinality())).Plus(_dafny.IntOfInt64(-182)), _dafny.IntOfUint32((Companion_Default___.Fm52(globalState)).Cardinality()))).Uint32(), _1776_v54)
		_ = _rhs361
		var _rhs362 _dafny.Int = _dafny.IntOfUint32(((_1791_v63).Select((Companion_Default___.SafeIndex((p0).Plus(p0), _dafny.IntOfUint32((_1791_v63).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
		_ = _rhs362
		var _rhs363 bool = (((_dafny.Zero).Minus((_1792_v64).Dtor_cf30())).Plus(_dafny.IntOfInt64(-837))).Cmp((p0).Minus(p0)) > 0
		_ = _rhs363
		var _lhs252 _dafny.Array = _1779_v57
		_ = _lhs252
		var _lhs253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_1779_v57), 0))
		_ = _lhs253
		var _lhs254 *GlobalState = globalState
		_ = _lhs254
		_1713_v0 = _rhs360
		(_lhs252).ArraySet1(_rhs361, (_lhs253).Int())
		_lhs254.F4 = _rhs362
		_1786_v58 = _rhs363
		var _1793_i6 _dafny.Int
		_ = _1793_i6
		_1793_i6 = _dafny.Zero
		{
			for p1 {
				{
					if (_1793_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1793_i6 = (_1793_i6).Plus(_dafny.One)
					var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1713_v0), 0))
					_ = _index302
					(_1713_v0).ArraySet1(false, (_index302).Int())
					var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1713_v0), 0))
					_ = _index303
					(_1713_v0).ArraySet1(!((p0).Cmp(p0) <= 0), (_index303).Int())
					var _1794_v65 _dafny.CodePoint
					_ = _1794_v65
					_1794_v65 = _dafny.CodePoint('l')
					var _1795_v66 _dafny.Map
					_ = _1795_v66
					_1795_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
					var _1796_v67 _dafny.Map
					_ = _1796_v67
					_1796_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1786_v58, _1794_v65)
					var _1797_v68 _dafny.Map
					_ = _1797_v68
					_1797_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1790_v62).Cardinality()))
					var _1798_v69 _dafny.Map
					_ = _1798_v69
					_1798_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if (_1795_v66).Contains(p1) {
							return (_1795_v66).Get(p1).(bool)
						}
						return p1
					})(), (_this).Fm10(_1796_v67, p0, _1797_v68, globalState))
					var _1799_v70 _dafny.Map
					_ = _1799_v70
					_1799_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1798_v69)
					var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1713_v0), 0))
					_ = _index304
					var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1713_v0), 0))
					_ = _index305
					var _rhs364 bool = (p0).Cmp(p0) == 0
					_ = _rhs364
					var _rhs365 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1786_v58, p1), _1789_v61), _1789_v61)
					_ = _rhs365
					var _rhs366 bool = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("aqqe"), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1790_v62, _dafny.UnicodeSeqOfUtf8Bytes("behwcy")), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1790_v62, _dafny.UnicodeSeqOfUtf8Bytes("behwcy"))).Cardinality()))).Uint32(), _1794_v65))
					_ = _rhs366
					var _rhs367 bool = (_this).Fm1(_1799_v70, p0, globalState)
					_ = _rhs367
					var _lhs255 _dafny.Array = _1713_v0
					_ = _lhs255
					var _lhs256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1713_v0), 0))
					_ = _lhs256
					var _lhs257 _dafny.Array = _1713_v0
					_ = _lhs257
					var _lhs258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1713_v0), 0))
					_ = _lhs258
					_1786_v58 = _rhs364
					_1789_v61 = _rhs365
					(_lhs255).ArraySet1(_rhs366, (_lhs256).Int())
					(_lhs257).ArraySet1(_rhs367, (_lhs258).Int())
					(globalState).F4 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(488)))
					var _1800_v71 D4
					_ = _1800_v71
					_1800_v71 = Companion_D4_.Create_DC12_((func() _dafny.Int {
						if p1 {
							return (_dafny.Zero).Minus(_dafny.IntOfUint32((_1790_v62).Cardinality()))
						}
						return p0
					})())
					var _source25 D4 = _1800_v71
					_ = _source25
					if _source25.Is_DC12() {
						var _1801___mcc_h0 _dafny.Int = _source25.Get_().(D4_DC12).Cf16
						_ = _1801___mcc_h0
						var _1802_cf16 _dafny.Int = _1801___mcc_h0
						_ = _1802_cf16
						var _1803_v72 *C2
						_ = _1803_v72
						var _nw260 *C2 = New_C2_()
						_ = _nw260
						_nw260.Ctor__()
						_1803_v72 = _nw260
						var _1804_v73 _dafny.Array
						_ = _1804_v73
						var _len0_46 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_46
						var _nw261 _dafny.Array
						_ = _nw261
						if _len0_46.Cmp(_dafny.Zero) == 0 {
							_nw261 = _dafny.NewArray(_len0_46)
						} else {
							var _init46 func(_dafny.Int) _dafny.Int = (func(_1805_cf16 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1806_i7 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeModuloInt(_1806_i7, _1805_cf16)
								}
							})(_1802_cf16)
							_ = _init46
							var _element0_46 = _init46(_dafny.Zero)
							_ = _element0_46
							_nw261 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
							(_nw261).ArraySet1(_element0_46, 0)
							var _nativeLen0_46 = (_len0_46).Int()
							_ = _nativeLen0_46
							for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
								(_nw261).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
							}
						}
						_1804_v73 = _nw261
						var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1804_v73), 0))
						_ = _index306
						(_1804_v73).ArraySet1(p0, (_index306).Int())
						var _1807_v74 _dafny.Array
						_ = _1807_v74
						var _nw262 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
						_ = _nw262
						_1807_v74 = _nw262
						var _1808_v75 _dafny.Map
						_ = _1808_v75
						_1808_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1807_v74, (func() bool {
							if p1 {
								return _1786_v58
							}
							return (_this).Fm8(_dafny.IntOfInt64(634), p1, p0, globalState)
						})())
						var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1804_v73), 0))
						_ = _index307
						(_1804_v73).ArraySet1((_1808_v75).Cardinality(), (_index307).Int())
						var _1809_v76 _dafny.Array
						_ = _1809_v76
						var _nw263 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
						_ = _nw263
						_1809_v76 = _nw263
						var _1810_v77 _dafny.Map
						_ = _1810_v77
						_1810_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1786_v58, _1809_v76)
						var _1811_v78 _dafny.MultiSet
						_ = _1811_v78
						_1811_v78 = _dafny.MultiSetOf(_1802_cf16)
						_1810_v77 = (_1810_v77).Update((_1811_v78).IsProperSubsetOf((_1811_v78).Update(p0, Companion_Default___.Abs(_1802_cf16))), _1809_v76)
						var _1812_v80 D15
						_ = _1812_v80
						_1812_v80 = Companion_D15_.Create_DC38_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1802_cf16, _1786_v58))
						var _1813_v81 _dafny.MultiSet
						_ = _1813_v81
						_1813_v81 = _dafny.MultiSetOf(Companion_D15_.Create_DC38_(_1788_v60), Companion_D15_.Create_DC38_(func() _dafny.Map {
							var _coll77 = _dafny.NewMapBuilder()
							_ = _coll77
							for _iter79 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-584))).Uint32(), func(coer107 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg107 _dafny.Int) interface{} {
									return coer107(arg107)
								}
							}((func(_1814_v73 _dafny.Array) func(_dafny.Int) _dafny.Int {
								return func(_1815_i8 _dafny.Int) _dafny.Int {
									return (_1814_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1814_v73), 0))).Int()).(_dafny.Int)
								}
							})(_1804_v73)))).Elements()); ; {
								_compr_77, _ok79 := _iter79()
								if !_ok79 {
									break
								}
								var _1816_v79 _dafny.Int
								_1816_v79 = interface{}(_compr_77).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-584))).Uint32(), func(coer108 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg108 _dafny.Int) interface{} {
										return coer108(arg108)
									}
								}((func(_1817_v73 _dafny.Array) func(_dafny.Int) _dafny.Int {
									return func(_1815_i8 _dafny.Int) _dafny.Int {
										return (_1817_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1817_v73), 0))).Int()).(_dafny.Int)
									}
								})(_1804_v73))), _1816_v79) {
									_coll77.Add((_1816_v79).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)).Cardinality()), p1)
								}
							}
							return _coll77.ToMap()
						}()), _1812_v80)
						var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(409), _dafny.ArrayLen((_1804_v73), 0))
						_ = _index308
						(_1804_v73).ArraySet1((func() _dafny.Int {
							if (_1813_v81).Contains(_1812_v80) {
								return (_1813_v81).Multiplicity(_1812_v80)
							}
							return ((_dafny.SetOf((_1713_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1713_v0), 0))).Int()).(bool), (_1713_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1713_v0), 0))).Int()).(bool))).Union(_1771_v49)).Cardinality()
						})(), (_index308).Int())
					} else {
						var _1818___mcc_h1 _dafny.Sequence = _source25.Get_().(D4_DC11).Cf15
						_ = _1818___mcc_h1
						var _1819_cf15 _dafny.Sequence = _1818___mcc_h1
						_ = _1819_cf15
						_1786_v58 = !(!(true))
						var _1820_v82 D3
						_ = _1820_v82
						_1820_v82 = Companion_D3_.Create_DC10_(p0, p0)
						_1820_v82 = Companion_Default___.Fm53((_this).Fm2(p0, globalState), globalState)
						_1786_v58 = _1786_v58
						var _1821_v83 _dafny.Map
						_ = _1821_v83
						_1821_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm9(p0, globalState), (p0).Plus(p0))
						var _1822_v84 _dafny.Array
						_ = _1822_v84
						var _nwElement0_63 _dafny.Int = (_dafny.Zero).Minus(p0)
						_ = _nwElement0_63
						var _nw264 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(21))
						_ = _nw264
						(_nw264).ArraySet1(_nwElement0_63, 0)
						(_nw264).ArraySet1(p0, 1)
						(_nw264).ArraySet1((_this).Fm2((_1788_v60).Cardinality(), globalState), 2)
						(_nw264).ArraySet1(p0, 3)
						(_nw264).ArraySet1(p0, 4)
						(_nw264).ArraySet1((_dafny.MultiSetFromSeq(_1774_v52)).Cardinality(), 5)
						(_nw264).ArraySet1(p0, 6)
						(_nw264).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wcuryoktk")).Cardinality()), 7)
						(_nw264).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1790_v62).Cardinality()), _1786_v58)).Cardinality(), 8)
						(_nw264).ArraySet1(p0, 9)
						(_nw264).ArraySet1(p0, 10)
						(_nw264).ArraySet1(p0, 11)
						(_nw264).ArraySet1(p0, 12)
						(_nw264).ArraySet1(p0, 13)
						(_nw264).ArraySet1(p0, 14)
						(_nw264).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1790_v62, (Companion_Default___.SafeIndex((_1798_v69).Cardinality(), _dafny.IntOfUint32((_1790_v62).Cardinality()))).Uint32(), _1794_v65)).Cardinality()), 15)
						(_nw264).ArraySet1(_dafny.IntOfUint32((_1790_v62).Cardinality()), 16)
						(_nw264).ArraySet1(p0, 17)
						(_nw264).ArraySet1(p0, 18)
						(_nw264).ArraySet1(p0, 19)
						(_nw264).ArraySet1((_1774_v52).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.IntOfUint32((_1774_v52).Cardinality()))).Uint32()).(_dafny.Int), 20)
						_1822_v84 = _nw264
						var _1823_v85 _dafny.Map
						_ = _1823_v85
						_1823_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1822_v84, _1774_v52)
						var _1824_v86 _dafny.Sequence
						_ = _1824_v86
						_1824_v86 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(277))).Uint32(), func(coer109 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg109 _dafny.Int) interface{} {
								return coer109(arg109)
							}
						}((func(_1825_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1826_i9 _dafny.Int) _dafny.Int {
								return _1825_p0
							}
						})(p0))))
						var _1827_v87 _dafny.Map
						_ = _1827_v87
						_1827_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1774_v52).Cardinality()), _1789_v61)
						var _1828_v88 D1
						_ = _1828_v88
						_1828_v88 = Companion_D1_.Create_DC3_(_1790_v62, p1)
						_1821_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
							if (_1823_v85).Contains(_1822_v84) {
								return (_1823_v85).Get(_1822_v84).(_dafny.Sequence)
							}
							return (_1824_v86).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
								if (_1827_v87).Contains(p0) {
									return (_1827_v87).Get(p0).(_dafny.Sequence)
								}
								return _1789_v61
							})()).Cardinality())), _dafny.IntOfUint32((_1824_v86).Cardinality()))).Uint32()).(_dafny.Sequence)
						})(), _1774_v52), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1828_v88).Dtor_cf3(), _1790_v62)).Cardinality()))
					}
					(globalState).F7 = ((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(766)))).Times(p0)
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _1829_v89 *C0
		_ = _1829_v89
		var _nw265 *C0 = New_C0_()
		_ = _nw265
		_nw265.Ctor__()
		_1829_v89 = _nw265
		r0 = (_this).Fm2((_dafny.Zero).Minus(p0), globalState)
		var _1830_v90 _dafny.Set
		_ = _1830_v90
		_1830_v90 = _dafny.SetOf(Companion_D0_.Create_DC0_(_1713_v0))
		var _1831_v91 _dafny.MultiSet
		_ = _1831_v91
		_1831_v91 = _dafny.MultiSetOf(_1830_v90)
		r1 = ((_1831_v91).Union(_1831_v91)).Union(_1831_v91)
		return r0, r1
	}
}

// End of class C9

// Definition of class C10
type C10 struct {
	dummy byte
}

func New_C10_() *C10 {
	_this := C10{}

	return &_this
}

type CompanionStruct_C10_ struct {
}

var Companion_C10_ = CompanionStruct_C10_{}

func (_this *C10) Equals(other *C10) bool {
	return _this == other
}

func (_this *C10) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C10)
	return ok && _this.Equals(other)
}

func (*C10) String() string {
	return "_module.C10"
}

func Type_C10_() _dafny.TypeDescriptor {
	return type_C10_{}
}

type type_C10_ struct {
}

func (_this type_C10_) Default() interface{} {
	return (*C10)(nil)
}

func (_this type_C10_) String() string {
	return "main.C10"
}
func (_this *C10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C10{}
var _ _dafny.TraitOffspring = &C10{}

func (_this *C10) Ctor__() {
	{
	}
}
func (_this *C10) Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C10) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(_dafny.IntOfInt64(-601))
	}
}
func (_this *C10) Fm6(p0 bool, p1 bool, p2 _dafny.Map, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(func(_source26 D1) _dafny.Int {
			if _source26.Is_DC3() {
				var _1832___mcc_h0 _dafny.Sequence = _source26.Get_().(D1_DC3).Cf3
				_ = _1832___mcc_h0
				var _1833___mcc_h1 bool = _source26.Get_().(D1_DC3).Cf4
				_ = _1833___mcc_h1
				var _1834_cf4 bool = _1833___mcc_h1
				_ = _1834_cf4
				var _1835_cf3 _dafny.Sequence = _1832___mcc_h0
				_ = _1835_cf3
				return (Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1834_cf4, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_1834_cf4, _1834_cf4)).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(621), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1834_cf4, _1834_cf4)).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfInt64(60))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg110 _dafny.Int) interface{} {
						return coer110(arg110)
					}
				}(func(_1836_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(107)
				}))).Cardinality()), _1834_cf4)).Cardinality(), _dafny.IntOfInt64(262))).Cardinality()))).Cardinality())))).Cardinality())).Dtor_cf2()
			} else if _source26.Is_DC4() {
				var _1837___mcc_h2 _dafny.MultiSet = _source26.Get_().(D1_DC4).Cf5
				_ = _1837___mcc_h2
				var _1838_cf5 _dafny.MultiSet = _1837___mcc_h2
				_ = _1838_cf5
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, !(true), true, !(!(false))), _dafny.SeqOf(true, !(false)))).Cardinality()))
			} else {
				var _1839___mcc_h3 _dafny.Int = _source26.Get_().(D1_DC2).Cf2
				_ = _1839___mcc_h3
				var _1840_cf2 _dafny.Int = _1839___mcc_h3
				_ = _1840_cf2
				return _1840_cf2
			}
		}(Companion_D1_.Create_DC2_(_dafny.IntOfInt64(647))))
	}
}
func (_this *C10) M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _1841_v0 _dafny.Array
		_ = _1841_v0
		var _nw266 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw266
		_1841_v0 = _nw266
		var _1842_v1 _dafny.Sequence
		_ = _1842_v1
		_1842_v1 = _dafny.SeqOf(_1841_v0, _1841_v0)
		var _1843_v2 _dafny.Int
		_ = _1843_v2
		_1843_v2 = _dafny.IntOfInt64(216)
		_1841_v0 = (_1842_v1).Select((Companion_Default___.SafeIndex(_1843_v2, _dafny.IntOfUint32((_1842_v1).Cardinality()))).Uint32()).(_dafny.Array)
		(globalState).F4 = (_1843_v2).Minus(_1843_v2)
		var _1844_v3 bool
		_ = _1844_v3
		_1844_v3 = true
		var _1845_v4 _dafny.Map
		_ = _1845_v4
		_1845_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1844_v3, _1844_v3)
		_1844_v3 = (_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1844_v3, _1845_v4), (_1845_v4).Cardinality(), globalState)
		var _1846_v5 _dafny.MultiSet
		_ = _1846_v5
		_1846_v5 = _dafny.MultiSetOf(_1844_v3)
		var _1847_v6 D2
		_ = _1847_v6
		_1847_v6 = Companion_D2_.Create_DC7_(true, _1846_v5, _1841_v0)
		var _pat_let_tv40 = _1841_v0
		_ = _pat_let_tv40
		var _1848_v7 _dafny.Array
		_ = _1848_v7
		var _nwElement0_64 _dafny.Array = _1841_v0
		_ = _nwElement0_64
		var _nw267 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(9))
		_ = _nw267
		(_nw267).ArraySet1(_nwElement0_64, 0)
		(_nw267).ArraySet1(_1841_v0, 1)
		(_nw267).ArraySet1(_1841_v0, 2)
		(_nw267).ArraySet1(_1841_v0, 3)
		(_nw267).ArraySet1(_1841_v0, 4)
		(_nw267).ArraySet1(_1841_v0, 5)
		(_nw267).ArraySet1(_1841_v0, 6)
		(_nw267).ArraySet1((func(_pat_let30_0 D2) D2 {
			return func(_1849_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let31_0 _dafny.Array) D2 {
					return func(_1850_dt__update_hcf10_h0 _dafny.Array) D2 {
						return func(_pat_let32_0 bool) D2 {
							return func(_1851_dt__update_hcf8_h0 bool) D2 {
								return Companion_D2_.Create_DC7_(_1851_dt__update_hcf8_h0, (_1849_dt__update__tmp_h0).Dtor_cf9(), _1850_dt__update_hcf10_h0)
							}(_pat_let32_0)
						}(false)
					}(_pat_let31_0)
				}(_pat_let_tv40)
			}(_pat_let30_0)
		}(_1847_v6)).Dtor_cf10(), 7)
		(_nw267).ArraySet1(_1841_v0, 8)
		_1848_v7 = _nw267
		var _nw268 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
		_ = _nw268
		_1848_v7 = _nw268
		var _hi13 _dafny.Int = Companion_Default___.SafeDivisionInt((_1846_v5).Cardinality(), (_dafny.Zero).Minus(_1843_v2))
		_ = _hi13
		for _1852_i0 := _1843_v2; _1852_i0.Cmp(_hi13) < 0; _1852_i0 = _1852_i0.Plus(_dafny.One) {
			var _1853_v8 _dafny.Array
			_ = _1853_v8
			var _nw269 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw269
			_1853_v8 = _nw269
			_1853_v8 = _1853_v8
			var _1854_v9 _dafny.CodePoint
			_ = _1854_v9
			_1854_v9 = _dafny.CodePoint('j')
			_1854_v9 = _1854_v9
			var _1855_v10 _dafny.Map
			_ = _1855_v10
			_1855_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1844_v3, _1845_v4)
			var _1856_v11 _dafny.Map
			_ = _1856_v11
			_1856_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1843_v2, (_dafny.Zero).Minus((_1846_v5).Cardinality()))
			var _1857_v12 _dafny.Map
			_ = _1857_v12
			_1857_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).Fm6(false, (_this).Fm1(_1855_v10, _1843_v2, globalState), _1856_v11, _dafny.CodePoint('w'), globalState)).Minus(_dafny.IntOfInt64(446)), (_1846_v5).IsProperSubsetOf(_dafny.MultiSetOf(_1844_v3)))
			_1857_v12 = (_1857_v12).Update((_dafny.Zero).Minus(_1852_i0), _1844_v3)
			var _1858_v13 *C4
			_ = _1858_v13
			var _nw270 *C4 = New_C4_()
			_ = _nw270
			_nw270.Ctor__()
			_1858_v13 = _nw270
		}
		var _1859_v14 _dafny.CodePoint
		_ = _1859_v14
		_1859_v14 = _dafny.CodePoint('u')
		_1859_v14 = _1859_v14
	}
}
func (_this *C10) M5(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) (_dafny.Array, bool, _dafny.CodePoint) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r2
		var _1860_v0 _dafny.Set
		_ = _1860_v0
		_1860_v0 = _dafny.SetOf(p0, p0, p0)
		var _hi14 _dafny.Int = p0
		_ = _hi14
		for _1861_i0 := (_1860_v0).Cardinality(); _1861_i0.Cmp(_hi14) < 0; _1861_i0 = _1861_i0.Plus(_dafny.One) {
			r1 = !(true) || (p2)
			r1 = (p0).Cmp(_1861_i0) > 0
			var _1862_v1 _dafny.MultiSet
			_ = _1862_v1
			_1862_v1 = _dafny.MultiSetOf(_1861_i0, _1861_i0)
			var _1863_v2 _dafny.Sequence
			_ = _1863_v2
			_1863_v2 = _dafny.SeqOf(p0)
			r1 = (_1862_v1).IsSubsetOf((_1862_v1).Union(_dafny.MultiSetFromSeq(_1863_v2)))
			var _1864_v3 _dafny.Array
			_ = _1864_v3
			var _len0_47 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_47
			var _nw271 _dafny.Array
			_ = _nw271
			if _len0_47.Cmp(_dafny.Zero) == 0 {
				_nw271 = _dafny.NewArray(_len0_47)
			} else {
				var _init47 func(_dafny.Int) _dafny.Int = (func(_1865_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1866_i1 _dafny.Int) _dafny.Int {
						return (_1866_i1).Plus(_1865_i0)
					}
				})(_1861_i0)
				_ = _init47
				var _element0_47 = _init47(_dafny.Zero)
				_ = _element0_47
				_nw271 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
				(_nw271).ArraySet1(_element0_47, 0)
				var _nativeLen0_47 = (_len0_47).Int()
				_ = _nativeLen0_47
				for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
					(_nw271).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
				}
			}
			_1864_v3 = _nw271
			var _1867_v4 _dafny.MultiSet
			_ = _1867_v4
			_1867_v4 = _dafny.MultiSetOf(false)
			var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1864_v3), 0))
			_ = _index309
			(_1864_v3).ArraySet1((func() _dafny.Int {
				if (_1862_v1).Contains((_1867_v4).Cardinality()) {
					return (_1862_v1).Multiplicity((_1867_v4).Cardinality())
				}
				return p0
			})(), (_index309).Int())
			var _1868_v5 _dafny.Map
			_ = _1868_v5
			_1868_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1864_v3), 0))
			_ = _index310
			(_1864_v3).ArraySet1((_dafny.Zero).Minus(((_1861_i0).Plus(p0)).Times((func() _dafny.Int {
				if (_1868_v5).Contains(false) {
					return (_1868_v5).Get(false).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_dafny.IntOfInt64(-617))
			})())), (_index310).Int())
		}
		var _1869_v6 _dafny.Map
		_ = _1869_v6
		_1869_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
		var _1870_v7 _dafny.Map
		_ = _1870_v7
		_1870_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1869_v6)
		var _1871_i2 _dafny.Int
		_ = _1871_i2
		_1871_i2 = _dafny.Zero
		{
			for (_this).Fm1(_1870_v7, p0, globalState) {
				{
					if (_1871_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_1871_i2 = (_1871_i2).Plus(_dafny.One)
					var _1872_v8 _dafny.Sequence
					_ = _1872_v8
					_1872_v8 = _dafny.SeqOf(p1)
					r1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1872_v8, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1872_v8).Cardinality()))).Uint32(), p2), _1872_v8), _1872_v8)
					var _1873_v9 _dafny.CodePoint
					_ = _1873_v9
					_1873_v9 = _dafny.CodePoint('k')
					r2 = _1873_v9
					r1 = (func() bool {
						if false {
							return p2
						}
						return p1
					})()
					(globalState).F4 = p0
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		var _1874_v10 _dafny.Map
		_ = _1874_v10
		_1874_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1875_v11 _dafny.Sequence
		_ = _1875_v11
		_1875_v11 = _dafny.SeqOf(!(!(p2)), p1)
		var _1876_v12 D12
		_ = _1876_v12
		_1876_v12 = Companion_D12_.Create_DC33_(p1)
		var _1877_v13 _dafny.CodePoint
		_ = _1877_v13
		_1877_v13 = _dafny.CodePoint('k')
		var _1878_v14 D11
		_ = _1878_v14
		_1878_v14 = Companion_D11_.Create_DC30_(p1, p0)
		var _1879_v15 _dafny.Sequence
		_ = _1879_v15
		_1879_v15 = _dafny.SeqOf(_dafny.IntOfInt64(620), p0)
		var _1880_v16 _dafny.MultiSet
		_ = _1880_v16
		_1880_v16 = _dafny.MultiSetOf(!(p2), p1)
		var _1881_v17 *C5
		_ = _1881_v17
		var _nw272 *C5 = New_C5_()
		_ = _nw272
		_nw272.Ctor__()
		_1881_v17 = _nw272
		var _1882_v18 _dafny.Array
		_ = _1882_v18
		var _nwElement0_65 bool = p1
		_ = _nwElement0_65
		var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(27))
		_ = _nw273
		(_nw273).ArraySet1(_nwElement0_65, 0)
		(_nw273).ArraySet1(p2, 1)
		(_nw273).ArraySet1((_this).Fm1(_1870_v7, (_1874_v10).Cardinality(), globalState), 2)
		(_nw273).ArraySet1(p1, 3)
		(_nw273).ArraySet1(true, 4)
		(_nw273).ArraySet1(!(p1), 5)
		(_nw273).ArraySet1((p0).Cmp(p0) < 0, 6)
		(_nw273).ArraySet1(p1, 7)
		(_nw273).ArraySet1((_this).Fm1(_1870_v7, p0, globalState), 8)
		(_nw273).ArraySet1(p1, 9)
		(_nw273).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1875_v11, _1875_v11), 10)
		(_nw273).ArraySet1(!(p2), 11)
		(_nw273).ArraySet1((p0).Cmp(p0) <= 0, 12)
		(_nw273).ArraySet1((func() bool {
			if p2 {
				return p1
			}
			return (_this).Fm1(_1870_v7, p0, globalState)
		})(), 13)
		(_nw273).ArraySet1((_1876_v12).Dtor_cf50(), 14)
		(_nw273).ArraySet1(((_this).Fm6(p2, p1, _1874_v10, _1877_v13, globalState)).Cmp((_this).Fm6(p2, p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1878_v14).Dtor_cf47(), p0), _1877_v13, globalState)) == 0, 15)
		(_nw273).ArraySet1(p2, 16)
		(_nw273).ArraySet1(p2, 17)
		(_nw273).ArraySet1(p1, 18)
		(_nw273).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1879_v15, _1879_v15), 19)
		(_nw273).ArraySet1(p1, 20)
		(_nw273).ArraySet1(p1, 21)
		(_nw273).ArraySet1((_1880_v16).IsProperSubsetOf(_1880_v16), 22)
		(_nw273).ArraySet1(p2, 23)
		(_nw273).ArraySet1(false, 24)
		(_nw273).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1881_v17)).Contains(!(p1)), 25)
		(_nw273).ArraySet1(p1, 26)
		_1882_v18 = _nw273
		_1882_v18 = _1882_v18
		var _1883_v19 _dafny.Sequence
		_ = _1883_v19
		_1883_v19 = _dafny.UnicodeSeqOfUtf8Bytes("lpee")
		_1883_v19 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm19(p2, p1, globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm19(p2, p1, globalState)).Cardinality()))).Uint32(), _1877_v13)
		if false {
			var _1884_v20 _dafny.Array
			_ = _1884_v20
			var _nw274 _dafny.Array = _dafny.NewArrayWithValue(Companion_D19_.Default(), _dafny.IntOfInt64(23))
			_ = _nw274
			_1884_v20 = _nw274
			_1884_v20 = _1884_v20
			(globalState).F15 = p0
			_1883_v19 = Companion_Default___.Fm30((p0).Minus(p0), (p0).Cmp(p0) != 0, _1860_v0, (func() bool {
				if true {
					return p1
				}
				return p1
			})(), globalState)
			if p2 {
				r1 = p1
				var _1885_v21 _dafny.Array
				_ = _1885_v21
				var _nw275 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
				_ = _nw275
				_1885_v21 = _nw275
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1885_v21), 0))
				_ = _index311
				(_1885_v21).ArraySet1((_dafny.Zero).Minus(p0), (_index311).Int())
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1885_v21), 0))
				_ = _index312
				(_1885_v21).ArraySet1(p0, (_index312).Int())
				var _1886_v22 _dafny.Map
				_ = _1886_v22
				_1886_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if p2 {
						return _1882_v18
					}
					return _1882_v18
				})(), Companion_Default___.Fm17(p0, globalState))
				(globalState).F15 = (func() _dafny.Int {
					if (_1886_v22).Contains(_1882_v18) {
						return (_1886_v22).Get(_1882_v18).(_dafny.Int)
					}
					return (_dafny.Zero).Minus((_1885_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1885_v21), 0))).Int()).(_dafny.Int))
				})()
				var _1887_v23 _dafny.Array
				_ = _1887_v23
				var _nwElement0_66 _dafny.Array = _1885_v21
				_ = _nwElement0_66
				var _nw276 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(17))
				_ = _nw276
				(_nw276).ArraySet1(_nwElement0_66, 0)
				(_nw276).ArraySet1(_1885_v21, 1)
				(_nw276).ArraySet1(_1885_v21, 2)
				(_nw276).ArraySet1(_1885_v21, 3)
				(_nw276).ArraySet1(_1885_v21, 4)
				(_nw276).ArraySet1(_1885_v21, 5)
				(_nw276).ArraySet1(_1885_v21, 6)
				(_nw276).ArraySet1(_1885_v21, 7)
				(_nw276).ArraySet1(_1885_v21, 8)
				(_nw276).ArraySet1(_1885_v21, 9)
				(_nw276).ArraySet1(_1885_v21, 10)
				(_nw276).ArraySet1(_1885_v21, 11)
				(_nw276).ArraySet1(_1885_v21, 12)
				(_nw276).ArraySet1(_1885_v21, 13)
				(_nw276).ArraySet1(_1885_v21, 14)
				(_nw276).ArraySet1(_1885_v21, 15)
				(_nw276).ArraySet1(_1885_v21, 16)
				_1887_v23 = _nw276
				var _1888_v24 D20
				_ = _1888_v24
				_1888_v24 = Companion_D20_.Create_DC54_(_1887_v23)
				var _1889_v25 _dafny.Array
				_ = _1889_v25
				var _nwElement0_67 _dafny.Array = _1887_v23
				_ = _nwElement0_67
				var _nw277 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(4))
				_ = _nw277
				(_nw277).ArraySet1(_nwElement0_67, 0)
				(_nw277).ArraySet1(_1887_v23, 1)
				(_nw277).ArraySet1(_1887_v23, 2)
				(_nw277).ArraySet1((_1888_v24).Dtor_cf84(), 3)
				_1889_v25 = _nw277
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_1889_v25), 0))
				_ = _index313
				(_1889_v25).ArraySet1(_1887_v23, (_index313).Int())
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_1889_v25), 0))
				_ = _index314
				(_1889_v25).ArraySet1(_1887_v23, (_index314).Int())
				var _1890_v26 _dafny.Set
				_ = _1890_v26
				_1890_v26 = _dafny.SetOf(p1)
				var _rhs368 _dafny.Array = _1885_v21
				_ = _rhs368
				var _rhs369 _dafny.Int = (_1885_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1885_v21), 0))).Int()).(_dafny.Int)
				_ = _rhs369
				var _rhs370 bool = (_1890_v26).IsProperSubsetOf(_1890_v26)
				_ = _rhs370
				var _rhs371 bool = _dafny.Companion_Sequence_.Equal(_1883_v19, Companion_Default___.Fm19(false, p1, globalState))
				_ = _rhs371
				var _lhs259 *GlobalState = globalState
				_ = _lhs259
				_1885_v21 = _rhs368
				_lhs259.F13 = _rhs369
				r1 = _rhs370
				r1 = _rhs371
			} else {
				var _1891_v27 _dafny.Array
				_ = _1891_v27
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_48
				var _nw278 _dafny.Array
				_ = _nw278
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw278 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) _dafny.Map = (func(_1892_p1 bool) func(_dafny.Int) _dafny.Map {
						return func(_1893_i3 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1892_p1, _dafny.IntOfInt64(789))
						}
					})(p1)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw278 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw278).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw278).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1891_v27 = _nw278
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1891_v27), 0))
				_ = _index315
				(_1891_v27).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_1879_v15).Cardinality()))), (_index315).Int())
				var _1894_v28 _dafny.Map
				_ = _1894_v28
				_1894_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1891_v27), 0))
				_ = _index316
				var _rhs372 bool = p1
				_ = _rhs372
				var _rhs373 _dafny.Map = ((func() _dafny.Map {
					if p1 {
						return _1894_v28
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
				})()).Merge(_1894_v28)
				_ = _rhs373
				var _rhs374 _dafny.Int = p0
				_ = _rhs374
				var _rhs375 _dafny.CodePoint = _1877_v13
				_ = _rhs375
				var _lhs260 _dafny.Array = _1891_v27
				_ = _lhs260
				var _lhs261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(593), _dafny.ArrayLen((_1891_v27), 0))
				_ = _lhs261
				var _lhs262 *GlobalState = globalState
				_ = _lhs262
				r1 = _rhs372
				(_lhs260).ArraySet1(_rhs373, (_lhs261).Int())
				_lhs262.F7 = _rhs374
				_1877_v13 = _rhs375
				var _1895_v29 _dafny.Array
				_ = _1895_v29
				var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw279
				_1895_v29 = _nw279
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_1895_v29), 0))
				_ = _index317
				(_1895_v29).ArraySet1(p0, (_index317).Int())
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_1895_v29), 0))
				_ = _index318
				(_1895_v29).ArraySet1((_1879_v15).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("booqtq")).Cardinality()), _dafny.IntOfUint32((_1879_v15).Cardinality()))).Uint32()).(_dafny.Int), (_index318).Int())
				var _1896_v30 *C8
				_ = _1896_v30
				var _nw280 *C8 = New_C8_()
				_ = _nw280
				_nw280.Ctor__()
				_1896_v30 = _nw280
				r1 = !((p2) || (p2))
				var _1897_v31 _dafny.Sequence
				_ = _1897_v31
				_1897_v31 = _dafny.SeqOf(_1882_v18)
				var _1898_v32 _dafny.Array
				_ = _1898_v32
				var _nwElement0_68 _dafny.Array = _1882_v18
				_ = _nwElement0_68
				var _nw281 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(23))
				_ = _nw281
				(_nw281).ArraySet1(_nwElement0_68, 0)
				(_nw281).ArraySet1(_1882_v18, 1)
				(_nw281).ArraySet1(_1882_v18, 2)
				(_nw281).ArraySet1(_1882_v18, 3)
				(_nw281).ArraySet1(_1882_v18, 4)
				(_nw281).ArraySet1((func() _dafny.Array {
					if p2 {
						return (_1897_v31).Select((Companion_Default___.SafeIndex((_1874_v10).Cardinality(), _dafny.IntOfUint32((_1897_v31).Cardinality()))).Uint32()).(_dafny.Array)
					}
					return _1882_v18
				})(), 5)
				(_nw281).ArraySet1(_1882_v18, 6)
				(_nw281).ArraySet1(_1882_v18, 7)
				(_nw281).ArraySet1(_1882_v18, 8)
				(_nw281).ArraySet1(_1882_v18, 9)
				(_nw281).ArraySet1(_1882_v18, 10)
				(_nw281).ArraySet1(_1882_v18, 11)
				(_nw281).ArraySet1(_1882_v18, 12)
				(_nw281).ArraySet1(_1882_v18, 13)
				(_nw281).ArraySet1(_1882_v18, 14)
				(_nw281).ArraySet1(_1882_v18, 15)
				(_nw281).ArraySet1(_1882_v18, 16)
				(_nw281).ArraySet1(_1882_v18, 17)
				(_nw281).ArraySet1(_1882_v18, 18)
				(_nw281).ArraySet1(_1882_v18, 19)
				(_nw281).ArraySet1((func() _dafny.Array {
					if p2 {
						return _1882_v18
					}
					return _1882_v18
				})(), 20)
				(_nw281).ArraySet1(_1882_v18, 21)
				(_nw281).ArraySet1(_1882_v18, 22)
				_1898_v32 = _nw281
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_1898_v32), 0))
				_ = _index319
				(_1898_v32).ArraySet1(_1882_v18, (_index319).Int())
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_1898_v32), 0))
				_ = _index320
				(_1898_v32).ArraySet1(_1882_v18, (_index320).Int())
			}
			if p1 {
				r1 = p2
				var _1899_v33 _dafny.Array
				_ = _1899_v33
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_49
				var _nw282 _dafny.Array
				_ = _nw282
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw282 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) _dafny.Set = (func(_1900_v0 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_1901_i4 _dafny.Int) _dafny.Set {
							return _1900_v0
						}
					})(_1860_v0)
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw282 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw282).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw282).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1899_v33 = _nw282
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_1899_v33), 0))
				_ = _index321
				(_1899_v33).ArraySet1(_dafny.SetOf(p0), (_index321).Int())
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_1899_v33), 0))
				_ = _index322
				(_1899_v33).ArraySet1(_1860_v0, (_index322).Int())
				var _1902_v34 _dafny.MultiSet
				_ = _1902_v34
				_1902_v34 = _dafny.MultiSetOf(p0)
				r1 = ((_1902_v34).Difference(_1902_v34)).IsSubsetOf(_1902_v34)
				var _1903_v35 _dafny.Map
				_ = _1903_v35
				_1903_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _1904_v36 _dafny.Map
				_ = _1904_v36
				_1904_v36 = _1903_v35
				var _1905_v37 _dafny.Map
				_ = _1905_v37
				_1905_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1903_v35, _dafny.Companion_Sequence_.Concatenate(_1883_v19, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(976))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}(func(_1906_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))))
				_1905_v37 = (_1905_v37).Update(_1904_v36, _1883_v19)
				var _1907_v38 _dafny.Set
				_ = _1907_v38
				_1907_v38 = _dafny.SetOf(p1)
				var _1908_v39 _dafny.Map
				_ = _1908_v39
				_1908_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Times(p0), _1907_v38)
				_1908_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1907_v38)
			} else {
				var _1909_v40 *C2
				_ = _1909_v40
				var _nw283 *C2 = New_C2_()
				_ = _nw283
				_nw283.Ctor__()
				_1909_v40 = _nw283
				var _1910_v41 _dafny.Sequence
				_ = _1910_v41
				_1910_v41 = _dafny.SeqOf(_1909_v40)
				var _1911_v42 _dafny.Map
				_ = _1911_v42
				_1911_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
				_1909_v40 = (_1910_v41).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1911_v42).Contains(!(p1)) {
						return (_1911_v42).Get(!(p1)).(_dafny.Int)
					}
					return p0
				})(), _dafny.IntOfUint32((_1910_v41).Cardinality()))).Uint32()).(*C2)
				var _1912_v43 _dafny.Array
				_ = _1912_v43
				var _nw284 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
				_ = _nw284
				_1912_v43 = _nw284
				var _1913_v44 D21
				_ = _1913_v44
				_1913_v44 = Companion_D21_.Create_DC56_(_1912_v43)
				var _1914_v45 _dafny.Array
				_ = _1914_v45
				var _nwElement0_69 _dafny.Array = _1912_v43
				_ = _nwElement0_69
				var _nw285 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(23))
				_ = _nw285
				(_nw285).ArraySet1(_nwElement0_69, 0)
				(_nw285).ArraySet1(_1912_v43, 1)
				(_nw285).ArraySet1(_1912_v43, 2)
				(_nw285).ArraySet1(_1912_v43, 3)
				(_nw285).ArraySet1(_1912_v43, 4)
				(_nw285).ArraySet1(_1912_v43, 5)
				(_nw285).ArraySet1(_1912_v43, 6)
				(_nw285).ArraySet1((func() _dafny.Array {
					if p2 {
						return _1912_v43
					}
					return _1912_v43
				})(), 7)
				(_nw285).ArraySet1(_1912_v43, 8)
				(_nw285).ArraySet1(_1912_v43, 9)
				(_nw285).ArraySet1(_1912_v43, 10)
				(_nw285).ArraySet1(_1912_v43, 11)
				(_nw285).ArraySet1(_1912_v43, 12)
				(_nw285).ArraySet1(_1912_v43, 13)
				(_nw285).ArraySet1(_1912_v43, 14)
				(_nw285).ArraySet1(_1912_v43, 15)
				(_nw285).ArraySet1(_1912_v43, 16)
				(_nw285).ArraySet1(_1912_v43, 17)
				(_nw285).ArraySet1((_1913_v44).Dtor_cf89(), 18)
				(_nw285).ArraySet1(_1912_v43, 19)
				(_nw285).ArraySet1(_1912_v43, 20)
				(_nw285).ArraySet1(_1912_v43, 21)
				(_nw285).ArraySet1(_1912_v43, 22)
				_1914_v45 = _nw285
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1914_v45), 0))
				_ = _index323
				(_1914_v45).ArraySet1(_1912_v43, (_index323).Int())
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1914_v45), 0))
				_ = _index324
				(_1914_v45).ArraySet1(_1912_v43, (_index324).Int())
				var _1915_v46 D5
				_ = _1915_v46
				_1915_v46 = Companion_D5_.Create_DC14_(_1877_v13, !(p2), p0, p1)
				var _1916_v47 D5
				_ = _1916_v47
				_1916_v47 = Companion_D5_.Create_DC15_(_1915_v46)
				_1916_v47 = _1916_v47
				var _1917_v48 *C0
				_ = _1917_v48
				var _nw286 *C0 = New_C0_()
				_ = _nw286
				_nw286.Ctor__()
				_1917_v48 = _nw286
				(globalState).F7 = Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_1879_v15).Cardinality()))
			}
		} else {
			var _1918_v49 _dafny.Array
			_ = _1918_v49
			var _nw287 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
			_ = _nw287
			_1918_v49 = _nw287
			var _1919_v50 _dafny.MultiSet
			_ = _1919_v50
			_1919_v50 = _dafny.MultiSetOf(_1918_v49)
			var _1920_v51 _dafny.Map
			_ = _1920_v51
			_1920_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1860_v0, (_1919_v50).Cardinality())
			(globalState).F7 = (func() _dafny.Int {
				if (_1920_v51).Contains((_1860_v0).Union(_1860_v0)) {
					return (_1920_v51).Get((_1860_v0).Union(_1860_v0)).(_dafny.Int)
				}
				return _dafny.IntOfInt64(781)
			})()
			r1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1879_v15, _1879_v15), _1879_v15)
			var _1921_v52 T0
			_ = _1921_v52
			var _nw288 *C7 = New_C7_()
			_ = _nw288
			_nw288.Ctor__()
			_1921_v52 = _nw288
			var _1922_v53 D17
			_ = _1922_v53
			_1922_v53 = Companion_D17_.Create_DC47_(_1921_v52, p1, p0, p0, p0)
			var _1923_v54 _dafny.Map
			_ = _1923_v54
			_1923_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1922_v53, !_dafny.Companion_Sequence_.Contains(_1883_v19, _dafny.CodePoint('h')))
			var _rhs376 _dafny.Array = _1882_v18
			_ = _rhs376
			var _rhs377 _dafny.Map = _1923_v54
			_ = _rhs377
			_1882_v18 = _rhs376
			_1923_v54 = _rhs377
			var _rhs378 _dafny.Int = (p0).Times(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(913))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg112 _dafny.Int) interface{} {
					return coer112(arg112)
				}
			}((func(_1924_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1925_i6 _dafny.Int) _dafny.CodePoint {
					return _1924_v13
				}
			})(_1877_v13))))).Cardinality()))).Times((_1879_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1879_v15).Cardinality()))).Uint32()).(_dafny.Int)))
			_ = _rhs378
			var _rhs379 _dafny.Int = p0
			_ = _rhs379
			var _rhs380 bool = p1
			_ = _rhs380
			var _lhs263 *GlobalState = globalState
			_ = _lhs263
			var _lhs264 *GlobalState = globalState
			_ = _lhs264
			_lhs263.F7 = _rhs378
			_lhs264.F15 = _rhs379
			r1 = _rhs380
			var _1926_v55 *C4
			_ = _1926_v55
			var _nw289 *C4 = New_C4_()
			_ = _nw289
			_nw289.Ctor__()
			_1926_v55 = _nw289
		}
		if p2 {
			(globalState).F7 = (func() _dafny.Int {
				if !(true) {
					return p0
				}
				return p0
			})()
			if p2 {
				var _1927_v56 T0
				_ = _1927_v56
				var _nw290 *C8 = New_C8_()
				_ = _nw290
				_nw290.Ctor__()
				_1927_v56 = _nw290
				var _1928_v57 _dafny.Map
				_ = _1928_v57
				_1928_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1927_v56)
				_1928_v57 = (_1928_v57).Update(p2, (func() T0 {
					if p2 {
						return _1927_v56
					}
					return _1927_v56
				})())
				var _len0_50 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_50
				var _nw291 _dafny.Array
				_ = _nw291
				if _len0_50.Cmp(_dafny.Zero) == 0 {
					_nw291 = _dafny.NewArray(_len0_50)
				} else {
					var _init50 func(_dafny.Int) bool = (func(_1929_p2 bool) func(_dafny.Int) bool {
						return func(_1930_i7 _dafny.Int) bool {
							return (_1929_p2) == (_1929_p2)
						}
					})(p2)
					_ = _init50
					var _element0_50 = _init50(_dafny.Zero)
					_ = _element0_50
					_nw291 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
					(_nw291).ArraySet1(_element0_50, 0)
					var _nativeLen0_50 = (_len0_50).Int()
					_ = _nativeLen0_50
					for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
						(_nw291).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
					}
				}
				_1882_v18 = _nw291
				r1 = p1
				var _1931_v58 _dafny.Sequence
				_ = _1931_v58
				_1931_v58 = _dafny.SeqOf(_1869_v6)
				var _1932_v59 D16
				_ = _1932_v59
				_1932_v59 = Companion_D16_.Create_DC41_(_1931_v58)
				var _1933_v60 D16
				_ = _1933_v60
				_1933_v60 = Companion_D16_.Create_DC44_(_1932_v59)
				_1933_v60 = _1933_v60
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1882_v18), 0))
				_ = _index325
				(_1882_v18).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1875_v11, _1875_v11), (_index325).Int())
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1882_v18), 0))
				_ = _index326
				(_1882_v18).ArraySet1(p1, (_index326).Int())
			} else {
				var _1934_v61 _dafny.Array
				_ = _1934_v61
				var _nw292 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw292
				_1934_v61 = _nw292
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1934_v61), 0))
				_ = _index327
				(_1934_v61).ArraySet1((_dafny.IntOfInt64(986)).Minus(p0), (_index327).Int())
				var _1935_v62 D3
				_ = _1935_v62
				_1935_v62 = Companion_D3_.Create_DC10_(p0, p0)
				var _1936_v63 D2
				_ = _1936_v63
				_1936_v63 = Companion_D2_.Create_DC7_(p2, _1880_v16, _1934_v61)
				var _1937_v64 _dafny.Sequence
				_ = _1937_v64
				_1937_v64 = _dafny.SeqOf(_1936_v63, _1936_v63, _1936_v63)
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1934_v61), 0))
				_ = _index328
				var _rhs381 _dafny.Sequence = _1875_v11
				_ = _rhs381
				var _rhs382 bool = p1
				_ = _rhs382
				var _rhs383 _dafny.Int = (_1935_v62).Dtor_cf14()
				_ = _rhs383
				var _rhs384 bool = _dafny.Companion_Sequence_.Equal(_1937_v64, _dafny.Companion_Sequence_.Concatenate(_1937_v64, _1937_v64))
				_ = _rhs384
				var _rhs385 _dafny.Int = ((p0).Times(p0)).Minus(_dafny.IntOfInt64(-95))
				_ = _rhs385
				var _lhs265 *GlobalState = globalState
				_ = _lhs265
				var _lhs266 _dafny.Array = _1934_v61
				_ = _lhs266
				var _lhs267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1934_v61), 0))
				_ = _lhs267
				_1875_v11 = _rhs381
				r1 = _rhs382
				_lhs265.F13 = _rhs383
				r1 = _rhs384
				(_lhs266).ArraySet1(_rhs385, (_lhs267).Int())
				r1 = p2
				_1934_v61 = _1934_v61
				var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1934_v61), 0))
				_ = _index329
				(_1934_v61).ArraySet1(p0, (_index329).Int())
				var _1938_v65 D2
				_ = _1938_v65
				_1938_v65 = Companion_D2_.Create_DC5_(_1934_v61)
				(_1881_v17).M6((func() _dafny.Int {
					if p2 {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1875_v11, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(128), p1)).Cardinality(), _dafny.IntOfUint32((_1875_v11).Cardinality()))).Uint32(), p1)).Cardinality())
					}
					return (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm19(p2, (Companion_D19_.Create_DC52_(p0, p1, p2)).Dtor_cf82(), globalState)).Cardinality()))))
				})(), (_1880_v16).Union(_1880_v16), (func() bool {
					if p2 {
						return p2
					}
					return p1
				})(), _1938_v65, globalState)
			}
			var _1939_v66 *C0
			_ = _1939_v66
			var _nw293 *C0 = New_C0_()
			_ = _nw293
			_nw293.Ctor__()
			_1939_v66 = _nw293
			var _1940_v67 _dafny.Set
			_ = _1940_v67
			_1940_v67 = _dafny.SetOf(_1860_v0)
			var _1941_v68 _dafny.Array
			_ = _1941_v68
			var _nw294 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw294
			_1941_v68 = _nw294
			var _1942_v69 _dafny.Map
			_ = _1942_v69
			_1942_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1941_v68)
			var _1943_v70 _dafny.Sequence
			_ = _1943_v70
			_1943_v70 = _dafny.SeqOf(_1882_v18, _1882_v18, _1882_v18)
			var _1944_v71 _dafny.MultiSet
			_ = _1944_v71
			_1944_v71 = _dafny.MultiSetOf(p0)
			var _1945_v72 D19
			_ = _1945_v72
			_1945_v72 = Companion_D19_.Create_DC52_(_dafny.IntOfInt64(297), false, p2)
			var _1946_v73 _dafny.Array
			_ = _1946_v73
			var _nwElement0_70 _dafny.Int = p0
			_ = _nwElement0_70
			var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(27))
			_ = _nw295
			(_nw295).ArraySet1(_nwElement0_70, 0)
			(_nw295).ArraySet1(p0, 1)
			(_nw295).ArraySet1((_dafny.SetOf(p0, (_1942_v69).Cardinality(), p0)).Cardinality(), 2)
			(_nw295).ArraySet1(p0, 3)
			(_nw295).ArraySet1(_dafny.IntOfUint32((_1943_v70).Cardinality()), 4)
			(_nw295).ArraySet1(p0, 5)
			(_nw295).ArraySet1(p0, 6)
			(_nw295).ArraySet1(_dafny.IntOfUint32((_1883_v19).Cardinality()), 7)
			(_nw295).ArraySet1((_1944_v71).Cardinality(), 8)
			(_nw295).ArraySet1(p0, 9)
			(_nw295).ArraySet1(p0, 10)
			(_nw295).ArraySet1((func() _dafny.Int {
				if (_1874_v10).Contains(_dafny.IntOfUint32((_1883_v19).Cardinality())) {
					return (_1874_v10).Get(_dafny.IntOfUint32((_1883_v19).Cardinality())).(_dafny.Int)
				}
				return p0
			})(), 11)
			(_nw295).ArraySet1(p0, 12)
			(_nw295).ArraySet1(_dafny.IntOfInt64(827), 13)
			(_nw295).ArraySet1(p0, 14)
			(_nw295).ArraySet1(p0, 15)
			(_nw295).ArraySet1(p0, 16)
			(_nw295).ArraySet1((_1945_v72).Dtor_cf81(), 17)
			(_nw295).ArraySet1((_dafny.Zero).Minus(p0), 18)
			(_nw295).ArraySet1(_dafny.IntOfInt64(607), 19)
			(_nw295).ArraySet1(p0, 20)
			(_nw295).ArraySet1((Companion_Default___.Fm31(p0, p0, globalState)).Cardinality(), 21)
			(_nw295).ArraySet1(_dafny.IntOfUint32((_1883_v19).Cardinality()), 22)
			(_nw295).ArraySet1(p0, 23)
			(_nw295).ArraySet1(p0, 24)
			(_nw295).ArraySet1(p0, 25)
			(_nw295).ArraySet1(p0, 26)
			_1946_v73 = _nw295
			var _1947_v74 _dafny.MultiSet
			_ = _1947_v74
			_1947_v74 = _dafny.MultiSetOf(_1946_v73)
			var _1948_v75 _dafny.Array
			_ = _1948_v75
			var _nwElement0_71 _dafny.MultiSet = _1947_v74
			_ = _nwElement0_71
			var _nw296 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(4))
			_ = _nw296
			(_nw296).ArraySet1(_nwElement0_71, 0)
			(_nw296).ArraySet1(((_1947_v74).Update(_1946_v73, Companion_Default___.Abs(p0))).Difference(_1947_v74), 1)
			(_nw296).ArraySet1(_1947_v74, 2)
			(_nw296).ArraySet1(_1947_v74, 3)
			_1948_v75 = _nw296
			var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1948_v75), 0))
			_ = _index330
			(_1948_v75).ArraySet1(_1947_v74, (_index330).Int())
			var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))
			_ = _index331
			(_1946_v73).ArraySet1(((_dafny.Zero).Minus(p0)).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_1875_v11).Cardinality()))), (_index331).Int())
			var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1948_v75), 0))
			_ = _index332
			var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))
			_ = _index333
			var _rhs386 _dafny.Set = ((_1940_v67).Union(_1940_v67)).Difference((_1940_v67).Difference(_dafny.SetOf(_1860_v0, _1860_v0)))
			_ = _rhs386
			var _rhs387 _dafny.Int = (p0).Plus(_dafny.IntOfUint32((_1875_v11).Cardinality()))
			_ = _rhs387
			var _rhs388 _dafny.MultiSet = (func() _dafny.MultiSet {
				if !(!(p2)) {
					return (_1947_v74).Difference(_1947_v74)
				}
				return _1947_v74
			})()
			_ = _rhs388
			var _rhs389 bool = (p0).Cmp(p0) != 0
			_ = _rhs389
			var _rhs390 _dafny.Int = Companion_Default___.SafeDivisionInt((_1879_v15).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1879_v15).Cardinality()))).Uint32()).(_dafny.Int), p0)
			_ = _rhs390
			var _lhs268 *GlobalState = globalState
			_ = _lhs268
			var _lhs269 _dafny.Array = _1948_v75
			_ = _lhs269
			var _lhs270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1948_v75), 0))
			_ = _lhs270
			var _lhs271 _dafny.Array = _1946_v73
			_ = _lhs271
			var _lhs272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))
			_ = _lhs272
			_1940_v67 = _rhs386
			_lhs268.F13 = _rhs387
			(_lhs269).ArraySet1(_rhs388, (_lhs270).Int())
			r1 = _rhs389
			(_lhs271).ArraySet1(_rhs390, (_lhs272).Int())
			if !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1875_v11, _dafny.SeqOf(p2, p2)), _dafny.Companion_Sequence_.Concatenate(_1875_v11, _1875_v11))) {
				var _1949_v76 *C5
				_ = _1949_v76
				var _nw297 *C5 = New_C5_()
				_ = _nw297
				_nw297.Ctor__()
				_1949_v76 = _nw297
				(_1949_v76).M6((_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int), Companion_Default___.Fm34((_dafny.Zero).Minus(_dafny.IntOfInt64(-881)), (_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int), !(p2), globalState), !(true), Companion_D2_.Create_DC5_(_1941_v68), globalState)
				(globalState).F4 = ((_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus((_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int)))
				var _1950_v77 _dafny.Array
				_ = _1950_v77
				var _nw298 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw298
				_1950_v77 = _nw298
				var _rhs391 _dafny.Array = (func() _dafny.Array {
					if p2 {
						return _1882_v18
					}
					return _1882_v18
				})()
				_ = _rhs391
				var _rhs392 _dafny.Array = _1950_v77
				_ = _rhs392
				var _rhs393 _dafny.Int = (_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int)
				_ = _rhs393
				var _lhs273 *GlobalState = globalState
				_ = _lhs273
				_1882_v18 = _rhs391
				_1950_v77 = _rhs392
				_lhs273.F13 = _rhs393
				var _1951_v78 _dafny.Sequence
				_ = _1951_v78
				_1951_v78 = _dafny.SeqOf(_1941_v68, _1946_v73)
				var _1952_v79 _dafny.Map
				_ = _1952_v79
				_1952_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1951_v78).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1951_v78).Cardinality()))).Uint32()).(_dafny.Array), (_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int))
				_1952_v79 = _1952_v79
			} else {
				var _1953_v80 _dafny.Array
				_ = _1953_v80
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_51
				var _nw299 _dafny.Array
				_ = _nw299
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw299 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) _dafny.Map = (func(_1954_p1 bool, _1955_p2 bool, _1956_v73 _dafny.Array, _1957_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_1958_i8 _dafny.Int) _dafny.Map {
							return (func() _dafny.Map {
								if _1954_p1 {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1955_p2, (_1956_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1956_v73), 0))).Int()).(_dafny.Int))
								}
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1954_p1, _1957_p0)
							})()
						}
					})(p1, p2, _1946_v73, p0)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw299 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw299).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw299).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_1953_v80 = _nw299
				var _1959_v81 _dafny.Map
				_ = _1959_v81
				_1959_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_1953_v80), 0))
				_ = _index334
				(_1953_v80).ArraySet1((_1959_v81).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1945_v72).Dtor_cf83(), p0)), (_index334).Int())
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_1953_v80), 0))
				_ = _index335
				(_1953_v80).ArraySet1((_1959_v81).Merge(_1959_v81), (_index335).Int())
				var _1960_v82 _dafny.Set
				_ = _1960_v82
				_1960_v82 = _dafny.SetOf(p2)
				var _1961_v83 _dafny.Sequence
				_ = _1961_v83
				_1961_v83 = _dafny.SeqOf(_1960_v82, _1960_v82, _1960_v82)
				var _1962_v84 D5
				_ = _1962_v84
				_1962_v84 = Companion_D5_.Create_DC14_(_1877_v13, false, p0, p2)
				var _rhs394 bool = false
				_ = _rhs394
				var _rhs395 bool = (Companion_D18_.Create_DC49_(!(p1), p2, (_1962_v84).Dtor_cf18(), p1)).Dtor_cf78()
				_ = _rhs395
				var _rhs396 _dafny.Sequence = _1961_v83
				_ = _rhs396
				var _rhs397 _dafny.Sequence = _1879_v15
				_ = _rhs397
				r1 = _rhs394
				r1 = _rhs395
				_1961_v83 = _rhs396
				_1879_v15 = _rhs397
				var _1963_v85 D19
				_ = _1963_v85
				_1963_v85 = Companion_D19_.Create_DC53_()
				_1963_v85 = Companion_Default___.Fm54((_1939_v66).Fm29(_1875_v11, _1883_v19, p2, p0, globalState), _1869_v6, _1944_v71, globalState)
				var _1964_v86 *C0
				_ = _1964_v86
				var _nw300 *C0 = New_C0_()
				_ = _nw300
				_nw300.Ctor__()
				_1964_v86 = _nw300
				var _1965_v88 _dafny.Map
				_ = _1965_v88
				_1965_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1877_v13, (_this).Fm2((_1880_v16).Cardinality(), globalState))
				var _1966_v89 _dafny.Sequence
				_ = _1966_v89
				_1966_v89 = _dafny.SeqOf(func() _dafny.Map {
					var _coll78 = _dafny.NewMapBuilder()
					_ = _coll78
					for _iter80 := _dafny.Iterate((_1965_v88).Keys().Elements()); ; {
						_compr_78, _ok80 := _iter80()
						if !_ok80 {
							break
						}
						var _1967_v87 _dafny.CodePoint
						_1967_v87 = interface{}(_compr_78).(_dafny.CodePoint)
						if (_1965_v88).Contains(_1967_v87) {
							_coll78.Add(_1967_v87, p2)
						}
					}
					return _coll78.ToMap()
				}())
				var _1968_v90 T0
				_ = _1968_v90
				var _nw301 *C2 = New_C2_()
				_ = _nw301
				_nw301.Ctor__()
				_1968_v90 = _nw301
				var _1969_v91 D17
				_ = _1969_v91
				_1969_v91 = Companion_D17_.Create_DC47_(_1968_v90, p2, (_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1883_v19).Cardinality()), (_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int))
				var _1970_v92 _dafny.Sequence
				_ = _1970_v92
				_1970_v92 = _dafny.SeqOf((_1953_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_1953_v80), 0))).Int()).(_dafny.Map))
				var _1971_v93 _dafny.Map
				_ = _1971_v93
				_1971_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1883_v19).Select((Companion_Default___.SafeIndex(((_1970_v92).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1970_v92).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.IntOfUint32((_1883_v19).Cardinality()))).Uint32()).(_dafny.CodePoint), (_1881_v17).Fm8(p0, false, (_1968_v90).Fm2(p0, globalState), globalState))
				var _1972_v94 D14
				_ = _1972_v94
				_1972_v94 = Companion_D14_.Create_DC36_(_1875_v11)
				var _1973_v95 _dafny.Map
				_ = _1973_v95
				_1973_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(p1))
				var _1974_v96 _dafny.Array
				_ = _1974_v96
				var _nwElement0_72 _dafny.Sequence = _1966_v89
				_ = _nwElement0_72
				var _nw302 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(12))
				_ = _nw302
				(_nw302).ArraySet1(_nwElement0_72, 0)
				(_nw302).ArraySet1(_1966_v89, 1)
				(_nw302).ArraySet1(_1966_v89, 2)
				(_nw302).ArraySet1(_1966_v89, 3)
				(_nw302).ArraySet1(_dafny.Companion_Sequence_.Update(_1966_v89, (Companion_Default___.SafeIndex((_1969_v91).Dtor_cf73(), _dafny.IntOfUint32((_1966_v89).Cardinality()))).Uint32(), _1971_v93), 4)
				(_nw302).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(881))).Uint32(), func(coer113 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}((func(_1975_v93 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1976_i9 _dafny.Int) _dafny.Map {
						return _1975_v93
					}
				})(_1971_v93))), 5)
				(_nw302).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(858))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg114 _dafny.Int) interface{} {
						return coer114(arg114)
					}
				}((func(_1977_v93 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1978_i10 _dafny.Int) _dafny.Map {
						return _1977_v93
					}
				})(_1971_v93))), 6)
				(_nw302).ArraySet1(_1966_v89, 7)
				(_nw302).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1966_v89, _dafny.SeqOf(_1971_v93, _1971_v93, _1971_v93, _1971_v93)), 8)
				(_nw302).ArraySet1(Companion_Default___.Fm55(p1, p0, p2, _1972_v94, globalState), 9)
				(_nw302).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1966_v89, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-814))).Uint32(), func(coer115 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg115 _dafny.Int) interface{} {
						return coer115(arg115)
					}
				}((func(_1979_v93 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1980_i11 _dafny.Int) _dafny.Map {
						return _1979_v93
					}
				})(_1971_v93)))), 10)
				(_nw302).ArraySet1(Companion_Default___.Fm55(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1973_v95).Contains((_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int)) {
						return (_1973_v95).Get((_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
					}
					return _1875_v11
				})()).Cardinality()), p1, _1972_v94, globalState), 11)
				_1974_v96 = _nw302
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_1974_v96), 0))
				_ = _index336
				(_1974_v96).ArraySet1(_1966_v89, (_index336).Int())
				var _1981_v97 _dafny.Map
				_ = _1981_v97
				_1981_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int), _1877_v13)
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_1974_v96), 0))
				_ = _index337
				var _rhs398 _dafny.Int = (p0).Plus((_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int))
				_ = _rhs398
				var _rhs399 bool = p1
				_ = _rhs399
				var _rhs400 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1966_v89, _1966_v89), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1966_v89, (Companion_Default___.SafeIndex((_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1966_v89).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1877_v13, false)), _dafny.Companion_Sequence_.Update(_1966_v89, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1959_v81).Contains(p2) {
						return (_1959_v81).Get(p2).(_dafny.Int)
					}
					return (_1946_v73).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_1946_v73), 0))).Int()).(_dafny.Int)
				})(), _dafny.IntOfUint32((_1966_v89).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
					if (_1981_v97).Contains((_dafny.Zero).Minus((_1880_v16).Cardinality())) {
						return (_1981_v97).Get((_dafny.Zero).Minus((_1880_v16).Cardinality())).(_dafny.CodePoint)
					}
					return _1877_v13
				})(), p1))))
				_ = _rhs400
				var _lhs274 *GlobalState = globalState
				_ = _lhs274
				var _lhs275 _dafny.Array = _1974_v96
				_ = _lhs275
				var _lhs276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_1974_v96), 0))
				_ = _lhs276
				_lhs274.F15 = _rhs398
				r1 = _rhs399
				(_lhs275).ArraySet1(_rhs400, (_lhs276).Int())
			}
		} else {
			var _1982_v98 D1
			_ = _1982_v98
			_1982_v98 = Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(627))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg116 _dafny.Int) interface{} {
					return coer116(arg116)
				}
			}((func(_1983_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1984_i12 _dafny.Int) _dafny.CodePoint {
					return _1983_v13
				}
			})(_1877_v13))), p1)
			var _pat_let_tv41 = _1883_v19
			_ = _pat_let_tv41
			var _source27 D1 = func(_pat_let33_0 D1) D1 {
				return func(_1985_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let34_0 _dafny.Sequence) D1 {
						return func(_1986_dt__update_hcf3_h0 _dafny.Sequence) D1 {
							return Companion_D1_.Create_DC3_(_1986_dt__update_hcf3_h0, (_1985_dt__update__tmp_h1).Dtor_cf4())
						}(_pat_let34_0)
					}(_pat_let_tv41)
				}(_pat_let33_0)
			}(_1982_v98)
			_ = _source27
			if _source27.Is_DC3() {
				var _1987___mcc_h0 _dafny.Sequence = _source27.Get_().(D1_DC3).Cf3
				_ = _1987___mcc_h0
				var _1988___mcc_h1 bool = _source27.Get_().(D1_DC3).Cf4
				_ = _1988___mcc_h1
				var _1989_cf4 bool = _1988___mcc_h1
				_ = _1989_cf4
				var _1990_cf3 _dafny.Sequence = _1987___mcc_h0
				_ = _1990_cf3
				var _1991_v99 D3
				_ = _1991_v99
				_1991_v99 = Companion_D3_.Create_DC10_(p0, (_1880_v16).Cardinality())
				(globalState).F4 = (_1991_v99).Dtor_cf13()
				var _1992_v100 _dafny.Array
				_ = _1992_v100
				var _nw303 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
				_ = _nw303
				_1992_v100 = _nw303
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1992_v100), 0))
				_ = _index338
				(_1992_v100).ArraySet1(_1990_cf3, (_index338).Int())
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1992_v100), 0))
				_ = _index339
				var _rhs401 bool = !((p1) || (p1))
				_ = _rhs401
				var _rhs402 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1990_cf3, _1990_cf3)
				_ = _rhs402
				var _lhs277 _dafny.Array = _1992_v100
				_ = _lhs277
				var _lhs278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1992_v100), 0))
				_ = _lhs278
				_1989_cf4 = _rhs401
				(_lhs277).ArraySet1(_rhs402, (_lhs278).Int())
				(globalState).F7 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_1874_v10).Contains(_dafny.IntOfInt64(982)) {
						return (_1874_v10).Get(_dafny.IntOfInt64(982)).(_dafny.Int)
					}
					return p0
				})(), (p0).Minus(p0))
				r1 = ((Companion_D16_.Create_DC42_(_1989_cf4, false, true)).Dtor_cf61()) == (p2)
			} else if _source27.Is_DC4() {
				var _1993___mcc_h2 _dafny.MultiSet = _source27.Get_().(D1_DC4).Cf5
				_ = _1993___mcc_h2
				var _1994_cf5 _dafny.MultiSet = _1993___mcc_h2
				_ = _1994_cf5
				var _1995_v101 _dafny.Map
				_ = _1995_v101
				_1995_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p2, p1, p1)).Cardinality(), (Companion_D8_.Create_DC23_(_1875_v11)).Dtor_cf39())
				r1 = (_dafny.IntOfInt64(120)).Cmp(Companion_Default___.Fm17((_1995_v101).Cardinality(), globalState)) == 0
				var _1996_v102 _dafny.Array
				_ = _1996_v102
				var _nw304 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw304
				_1996_v102 = _nw304
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_1996_v102), 0))
				_ = _index340
				(_1996_v102).ArraySet1(_dafny.IntOfUint32((_1883_v19).Cardinality()), (_index340).Int())
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_1996_v102), 0))
				_ = _index341
				(_1996_v102).ArraySet1(p0, (_index341).Int())
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_1882_v18), 0))
				_ = _index342
				(_1882_v18).ArraySet1(p2, (_index342).Int())
				var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_1882_v18), 0))
				_ = _index343
				(_1882_v18).ArraySet1((_1875_v11).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p0, (_1996_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_1996_v102), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_1875_v11).Cardinality()))).Uint32()).(bool), (_index343).Int())
				var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_1882_v18), 0))
				_ = _index344
				(_1882_v18).ArraySet1(p2, (_index344).Int())
			} else {
				var _1997___mcc_h3 _dafny.Int = _source27.Get_().(D1_DC2).Cf2
				_ = _1997___mcc_h3
				var _1998_cf2 _dafny.Int = _1997___mcc_h3
				_ = _1998_cf2
				var _1999_v103 *C2
				_ = _1999_v103
				var _nw305 *C2 = New_C2_()
				_ = _nw305
				_nw305.Ctor__()
				_1999_v103 = _nw305
				r1 = p2
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1882_v18), 0))
				_ = _index345
				(_1882_v18).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1883_v19, _1883_v19), (_index345).Int())
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1882_v18), 0))
				_ = _index346
				var _rhs403 bool = !(false)
				_ = _rhs403
				var _rhs404 _dafny.Int = p0
				_ = _rhs404
				var _rhs405 bool = p1
				_ = _rhs405
				var _rhs406 bool = (_1998_cf2).Cmp(_dafny.IntOfUint32((_1879_v15).Cardinality())) != 0
				_ = _rhs406
				var _lhs279 _dafny.Array = _1882_v18
				_ = _lhs279
				var _lhs280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1882_v18), 0))
				_ = _lhs280
				r1 = _rhs403
				_1998_cf2 = _rhs404
				(_lhs279).ArraySet1(_rhs405, (_lhs280).Int())
				r1 = _rhs406
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_1882_v18), 0))
				_ = _index347
				(_1882_v18).ArraySet1(false, (_index347).Int())
			}
			r1 = p1
			var _2000_v104 _dafny.Sequence
			_ = _2000_v104
			_2000_v104 = _dafny.SeqOf(_1880_v16)
			var _2001_v105 _dafny.Array
			_ = _2001_v105
			var _nwElement0_73 _dafny.MultiSet = _1880_v16
			_ = _nwElement0_73
			var _nw306 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(18))
			_ = _nw306
			(_nw306).ArraySet1(_nwElement0_73, 0)
			(_nw306).ArraySet1((_1880_v16).Union(_1880_v16), 1)
			(_nw306).ArraySet1(Companion_Default___.Fm34(p0, (_1880_v16).Cardinality(), p2, globalState), 2)
			(_nw306).ArraySet1((_1880_v16).Union(_1880_v16), 3)
			(_nw306).ArraySet1(_1880_v16, 4)
			(_nw306).ArraySet1(_1880_v16, 5)
			(_nw306).ArraySet1(_1880_v16, 6)
			(_nw306).ArraySet1(_1880_v16, 7)
			(_nw306).ArraySet1(_1880_v16, 8)
			(_nw306).ArraySet1(_dafny.MultiSetOf(!(p2), !(p1)), 9)
			(_nw306).ArraySet1((_1880_v16).Intersection(_1880_v16), 10)
			(_nw306).ArraySet1((_1880_v16).Intersection(_1880_v16), 11)
			(_nw306).ArraySet1(_1880_v16, 12)
			(_nw306).ArraySet1((_1880_v16).Intersection(_1880_v16), 13)
			(_nw306).ArraySet1(_1880_v16, 14)
			(_nw306).ArraySet1((_2000_v104).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2000_v104).Cardinality()))).Uint32()).(_dafny.MultiSet), 15)
			(_nw306).ArraySet1(_1880_v16, 16)
			(_nw306).ArraySet1(_1880_v16, 17)
			_2001_v105 = _nw306
			var _rhs407 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1883_v19, _dafny.Companion_Sequence_.Concatenate(_1883_v19, _1883_v19))
			_ = _rhs407
			var _rhs408 _dafny.Array = _2001_v105
			_ = _rhs408
			var _rhs409 _dafny.Int = p0
			_ = _rhs409
			var _lhs281 *GlobalState = globalState
			_ = _lhs281
			_1883_v19 = _rhs407
			r0 = _rhs408
			_lhs281.F13 = _rhs409
			var _2002_v106 *C4
			_ = _2002_v106
			var _nw307 *C4 = New_C4_()
			_ = _nw307
			_nw307.Ctor__()
			_2002_v106 = _nw307
			var _2003_v107 _dafny.Map
			_ = _2003_v107
			_2003_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2002_v106), p2)
			_2003_v107 = _2003_v107
			var _2004_v108 _dafny.Array
			_ = _2004_v108
			var _nw308 _dafny.Array = _dafny.NewArrayWithValue(Companion_D21_.Default(), _dafny.IntOfInt64(13))
			_ = _nw308
			_2004_v108 = _nw308
			var _2005_v109 _dafny.Array
			_ = _2005_v109
			var _nw309 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw309
			_2005_v109 = _nw309
			var _2006_v110 D21
			_ = _2006_v110
			_2006_v110 = Companion_D21_.Create_DC56_(_2005_v109)
			var _2007_v111 D21
			_ = _2007_v111
			_2007_v111 = Companion_D21_.Create_DC56_((_2006_v110).Dtor_cf89())
			var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_2004_v108), 0))
			_ = _index348
			(_2004_v108).ArraySet1(_2007_v111, (_index348).Int())
			var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_2004_v108), 0))
			_ = _index349
			(_2004_v108).ArraySet1(_2006_v110, (_index349).Int())
		}
		var _2008_v112 _dafny.Array
		_ = _2008_v112
		var _nw310 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(4))
		_ = _nw310
		_2008_v112 = _nw310
		r0 = _2008_v112
		r1 = p1
		r2 = _1877_v13
		return r0, r1, r2
	}
}

// End of class C10

// Definition of class C11
type C11 struct {
	dummy byte
}

func New_C11_() *C11 {
	_this := C11{}

	return &_this
}

type CompanionStruct_C11_ struct {
}

var Companion_C11_ = CompanionStruct_C11_{}

func (_this *C11) Equals(other *C11) bool {
	return _this == other
}

func (_this *C11) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C11)
	return ok && _this.Equals(other)
}

func (*C11) String() string {
	return "_module.C11"
}

func Type_C11_() _dafny.TypeDescriptor {
	return type_C11_{}
}

type type_C11_ struct {
}

func (_this type_C11_) Default() interface{} {
	return (*C11)(nil)
}

func (_this type_C11_) String() string {
	return "main.C11"
}
func (_this *C11) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C11{}
var _ _dafny.TraitOffspring = &C11{}

func (_this *C11) Ctor__() {
	{
	}
}
func (_this *C11) Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C11) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((func(_source28 D1) _dafny.Sequence {
			if _source28.Is_DC3() {
				var _2009___mcc_h0 _dafny.Sequence = _source28.Get_().(D1_DC3).Cf3
				_ = _2009___mcc_h0
				var _2010___mcc_h1 bool = _source28.Get_().(D1_DC3).Cf4
				_ = _2010___mcc_h1
				var _2011_cf4 bool = _2010___mcc_h1
				_ = _2011_cf4
				var _2012_cf3 _dafny.Sequence = _2009___mcc_h0
				_ = _2012_cf3
				return _2012_cf3
			} else if _source28.Is_DC4() {
				var _2013___mcc_h2 _dafny.MultiSet = _source28.Get_().(D1_DC4).Cf5
				_ = _2013___mcc_h2
				var _2014_cf5 _dafny.MultiSet = _2013___mcc_h2
				_ = _2014_cf5
				return _dafny.UnicodeSeqOfUtf8Bytes("rbfn")
			} else {
				var _2015___mcc_h3 _dafny.Int = _source28.Get_().(D1_DC2).Cf2
				_ = _2015___mcc_h3
				var _2016_cf2 _dafny.Int = _2015___mcc_h3
				_ = _2016_cf2
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("aredcecm"), _dafny.UnicodeSeqOfUtf8Bytes("lopn"))
			}
		}((func() D1 {
			if false {
				return Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("gjiblb"), false)
			}
			return Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("ahql"), true)
		})())).Cardinality())
	}
}
func (_this *C11) M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _2017_v0 _dafny.Sequence
		_ = _2017_v0
		_2017_v0 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		_2017_v0 = _2017_v0
		var _2018_i0 _dafny.Int
		_ = _2018_i0
		_2018_i0 = _dafny.Zero
		{
			for false {
				{
					if (_2018_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_2018_i0 = (_2018_i0).Plus(_dafny.One)
					var _2019_v1 bool
					_ = _2019_v1
					_2019_v1 = false
					_2019_v1 = _2019_v1
					if true {
						_2019_v1 = true
						var _2020_v2 _dafny.CodePoint
						_ = _2020_v2
						_2020_v2 = _dafny.CodePoint('g')
						var _2021_v3 _dafny.Int
						_ = _2021_v3
						_2021_v3 = _dafny.IntOfInt64(-854)
						var _rhs410 _dafny.Int = (func() _dafny.Int {
							if _2019_v1 {
								return _2021_v3
							}
							return _2021_v3
						})()
						_ = _rhs410
						var _rhs411 _dafny.CodePoint = _2020_v2
						_ = _rhs411
						var _lhs282 *GlobalState = globalState
						_ = _lhs282
						_lhs282.F7 = _rhs410
						_2020_v2 = _rhs411
						var _2022_v4 _dafny.Map
						_ = _2022_v4
						_2022_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2021_v3, (_dafny.SetOf(_dafny.IntOfInt64(-798))).Cardinality())
						var _2023_v5 _dafny.Sequence
						_ = _2023_v5
						_2023_v5 = _dafny.SeqOf((_2022_v4).Cardinality(), _2021_v3)
						var _2024_v6 _dafny.Set
						_ = _2024_v6
						_2024_v6 = _dafny.SetOf(_2021_v3, _dafny.IntOfInt64(-516), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-427))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg117 _dafny.Int) interface{} {
								return coer117(arg117)
							}
						}((func(_2025_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2026_i1 _dafny.Int) _dafny.CodePoint {
								return _2025_v2
							}
						})(_2020_v2)))).Cardinality()))
						var _2027_v7 _dafny.Array
						_ = _2027_v7
						var _nwElement0_74 bool = _2019_v1
						_ = _nwElement0_74
						var _nw311 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(5))
						_ = _nw311
						(_nw311).ArraySet1(_nwElement0_74, 0)
						(_nw311).ArraySet1(_2019_v1, 1)
						(_nw311).ArraySet1(_2019_v1, 2)
						(_nw311).ArraySet1(_2019_v1, 3)
						(_nw311).ArraySet1(_2019_v1, 4)
						_2027_v7 = _nw311
						var _2028_v8 _dafny.Map
						_ = _2028_v8
						_2028_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2027_v7, _dafny.IntOfInt64(950))
						var _2029_v9 _dafny.Sequence
						_ = _2029_v9
						_2029_v9 = _dafny.SeqOf(_2019_v1)
						var _2030_v10 _dafny.Array
						_ = _2030_v10
						var _nwElement0_75 _dafny.Int = _2021_v3
						_ = _nwElement0_75
						var _nw312 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(25))
						_ = _nw312
						(_nw312).ArraySet1(_nwElement0_75, 0)
						(_nw312).ArraySet1(_2021_v3, 1)
						(_nw312).ArraySet1(_dafny.IntOfInt64(591), 2)
						(_nw312).ArraySet1(_2021_v3, 3)
						(_nw312).ArraySet1(_dafny.IntOfUint32((_2023_v5).Cardinality()), 4)
						(_nw312).ArraySet1((_2024_v6).Cardinality(), 5)
						(_nw312).ArraySet1((_dafny.Zero).Minus(_2021_v3), 6)
						(_nw312).ArraySet1((_2028_v8).Cardinality(), 7)
						(_nw312).ArraySet1(_2021_v3, 8)
						(_nw312).ArraySet1(_dafny.IntOfInt64(321), 9)
						(_nw312).ArraySet1(_2021_v3, 10)
						(_nw312).ArraySet1(_2021_v3, 11)
						(_nw312).ArraySet1(_2021_v3, 12)
						(_nw312).ArraySet1(_2021_v3, 13)
						(_nw312).ArraySet1(_dafny.IntOfUint32((_2029_v9).Cardinality()), 14)
						(_nw312).ArraySet1(_2021_v3, 15)
						(_nw312).ArraySet1(_2021_v3, 16)
						(_nw312).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 17)
						(_nw312).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_2023_v5).Cardinality())))), 18)
						(_nw312).ArraySet1(_2021_v3, 19)
						(_nw312).ArraySet1(_2021_v3, 20)
						(_nw312).ArraySet1(_2021_v3, 21)
						(_nw312).ArraySet1(_2021_v3, 22)
						(_nw312).ArraySet1((_dafny.Zero).Minus(_2021_v3), 23)
						(_nw312).ArraySet1(_2021_v3, 24)
						_2030_v10 = _nw312
						var _2031_v11 _dafny.Sequence
						_ = _2031_v11
						_2031_v11 = _dafny.SeqOf(_2030_v10)
						var _rhs412 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("l"), _2017_v0), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm5(globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2023_v5).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm5(globalState)).Cardinality()))).Uint32(), _2020_v2))
						_ = _rhs412
						var _rhs413 _dafny.Int = _2021_v3
						_ = _rhs413
						var _rhs414 _dafny.Sequence = _2031_v11
						_ = _rhs414
						var _lhs283 *GlobalState = globalState
						_ = _lhs283
						_2019_v1 = _rhs412
						_lhs283.F15 = _rhs413
						_2031_v11 = _rhs414
						_2017_v0 = p0
						(globalState).F13 = _dafny.IntOfInt64(-597)
					} else {
						var _2032_v12 _dafny.Array
						_ = _2032_v12
						var _nw313 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(21))
						_ = _nw313
						_2032_v12 = _nw313
						var _2033_v13 D1
						_ = _2033_v13
						_2033_v13 = Companion_D1_.Create_DC3_(p0, true)
						var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_2032_v12), 0))
						_ = _index350
						(_2032_v12).ArraySet1(_2033_v13, (_index350).Int())
						var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_2032_v12), 0))
						_ = _index351
						(_2032_v12).ArraySet1(_2033_v13, (_index351).Int())
						_2019_v1 = _2019_v1
						var _2034_v14 _dafny.Int
						_ = _2034_v14
						_2034_v14 = _dafny.IntOfInt64(586)
						var _2035_v15 _dafny.Map
						_ = _2035_v15
						_2035_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2034_v14, _2034_v14)
						(globalState).F13 = (_this).Fm2((_2034_v14).Plus((_2035_v15).Cardinality()), globalState)
						var _2036_v16 _dafny.CodePoint
						_ = _2036_v16
						_2036_v16 = _dafny.CodePoint('w')
						var _2037_v17 _dafny.Map
						_ = _2037_v17
						_2037_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2036_v16, p0)
						_2037_v17 = ((_2037_v17).Merge(_2037_v17)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2036_v16, _2017_v0))
						(globalState).F4 = _2034_v14
					}
					var _2038_v18 _dafny.Array
					_ = _2038_v18
					var _nw314 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(8))
					_ = _nw314
					_2038_v18 = _nw314
					var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_2038_v18), 0))
					_ = _index352
					(_2038_v18).ArraySet1CodePoint(_dafny.CodePoint('l'), (_index352).Int())
					var _2039_v19 _dafny.CodePoint
					_ = _2039_v19
					_2039_v19 = _dafny.CodePoint('m')
					var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_2038_v18), 0))
					_ = _index353
					(_2038_v18).ArraySet1CodePoint(_2039_v19, (_index353).Int())
					var _2040_v20 _dafny.Int
					_ = _2040_v20
					_2040_v20 = _dafny.IntOfInt64(534)
					var _2041_v21 _dafny.Sequence
					_ = _2041_v21
					_2041_v21 = _dafny.SeqOf(_dafny.IntOfInt64(-358), _2040_v20)
					var _2042_v22 _dafny.MultiSet
					_ = _2042_v22
					_2042_v22 = _dafny.MultiSetOf(_2040_v20, _2040_v20, (_dafny.Zero).Minus(_2040_v20), _dafny.IntOfUint32((_2041_v21).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_2040_v20, _2040_v20)).Cardinality()))
					var _2043_v23 _dafny.MultiSet
					_ = _2043_v23
					_2043_v23 = _dafny.MultiSetOf(_2019_v1, _2019_v1)
					var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_2038_v18), 0))
					_ = _index354
					var _rhs415 _dafny.Int = (_dafny.Zero).Minus(((_2042_v22).Update(_2040_v20, Companion_Default___.Abs((_2043_v23).Cardinality()))).Cardinality())
					_ = _rhs415
					var _rhs416 _dafny.CodePoint = _2039_v19
					_ = _rhs416
					var _rhs417 _dafny.Int = (_2040_v20).Plus((_dafny.Zero).Minus(_2040_v20))
					_ = _rhs417
					var _rhs418 _dafny.Int = (_this).Fm2((_2040_v20).Times(_2040_v20), globalState)
					_ = _rhs418
					var _rhs419 bool = (_2040_v20).Cmp(_dafny.IntOfUint32((_2041_v21).Cardinality())) <= 0
					_ = _rhs419
					var _lhs284 *GlobalState = globalState
					_ = _lhs284
					var _lhs285 _dafny.Array = _2038_v18
					_ = _lhs285
					var _lhs286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_2038_v18), 0))
					_ = _lhs286
					var _lhs287 *GlobalState = globalState
					_ = _lhs287
					var _lhs288 *GlobalState = globalState
					_ = _lhs288
					_lhs284.F13 = _rhs415
					(_lhs285).ArraySet1CodePoint(_rhs416, (_lhs286).Int())
					_lhs287.F7 = _rhs417
					_lhs288.F13 = _rhs418
					_2019_v1 = _rhs419
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		var _2044_v24 bool
		_ = _2044_v24
		_2044_v24 = true
		var _2045_v25 _dafny.Int
		_ = _2045_v25
		_2045_v25 = _dafny.IntOfInt64(261)
		var _2046_v26 _dafny.MultiSet
		_ = _2046_v26
		_2046_v26 = _dafny.MultiSetOf(_2045_v25)
		var _2047_v27 D1
		_ = _2047_v27
		_2047_v27 = Companion_D1_.Create_DC4_(_2046_v26)
		var _source29 D1 = (func() D1 {
			if _2044_v24 {
				return _2047_v27
			}
			return _2047_v27
		})()
		_ = _source29
		if _source29.Is_DC3() {
			var _2048___mcc_h0 _dafny.Sequence = _source29.Get_().(D1_DC3).Cf3
			_ = _2048___mcc_h0
			var _2049___mcc_h1 bool = _source29.Get_().(D1_DC3).Cf4
			_ = _2049___mcc_h1
			var _2050_cf4 bool = _2049___mcc_h1
			_ = _2050_cf4
			var _2051_cf3 _dafny.Sequence = _2048___mcc_h0
			_ = _2051_cf3
			var _2052_v28 D1
			_ = _2052_v28
			_2052_v28 = Companion_D1_.Create_DC2_((_2046_v26).Cardinality())
			(globalState).F7 = ((func() D1 {
				if _2050_cf4 {
					return Companion_D1_.Create_DC2_(_dafny.IntOfInt64(495))
				}
				return _2052_v28
			})()).Dtor_cf2()
			var _rhs420 _dafny.Int = _2045_v25
			_ = _rhs420
			var _rhs421 bool = _2050_cf4
			_ = _rhs421
			var _lhs289 *GlobalState = globalState
			_ = _lhs289
			_lhs289.F4 = _rhs420
			_2044_v24 = _rhs421
			_2050_cf4 = _2050_cf4
			var _2053_v29 T0
			_ = _2053_v29
			var _nw315 *C10 = New_C10_()
			_ = _nw315
			_nw315.Ctor__()
			_2053_v29 = _nw315
			_2053_v29 = _2053_v29
		} else if _source29.Is_DC4() {
			var _2054___mcc_h2 _dafny.MultiSet = _source29.Get_().(D1_DC4).Cf5
			_ = _2054___mcc_h2
			var _2055_cf5 _dafny.MultiSet = _2054___mcc_h2
			_ = _2055_cf5
			var _2056_v30 _dafny.Array
			_ = _2056_v30
			var _nwElement0_76 bool = _2044_v24
			_ = _nwElement0_76
			var _nw316 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(7))
			_ = _nw316
			(_nw316).ArraySet1(_nwElement0_76, 0)
			(_nw316).ArraySet1(false, 1)
			(_nw316).ArraySet1(_2044_v24, 2)
			(_nw316).ArraySet1(_2044_v24, 3)
			(_nw316).ArraySet1(_2044_v24, 4)
			(_nw316).ArraySet1(_2044_v24, 5)
			(_nw316).ArraySet1(_2044_v24, 6)
			_2056_v30 = _nw316
			var _2057_v31 _dafny.Sequence
			_ = _2057_v31
			_2057_v31 = _dafny.SeqOf(_2056_v30, _2056_v30)
			var _2058_v32 D7
			_ = _2058_v32
			_2058_v32 = Companion_D7_.Create_DC19_(_2057_v31)
			var _2059_v33 _dafny.Map
			_ = _2059_v33
			_2059_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), _2058_v32)
			var _2060_v34 _dafny.Set
			_ = _2060_v34
			_2060_v34 = _dafny.SetOf(_2045_v25)
			var _2061_v35 _dafny.Array
			_ = _2061_v35
			var _nwElement0_77 _dafny.Int = _dafny.IntOfInt64(223)
			_ = _nwElement0_77
			var _nw317 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(19))
			_ = _nw317
			(_nw317).ArraySet1(_nwElement0_77, 0)
			(_nw317).ArraySet1(_2045_v25, 1)
			(_nw317).ArraySet1(_2045_v25, 2)
			(_nw317).ArraySet1(_2045_v25, 3)
			(_nw317).ArraySet1((_this).Fm2(_dafny.IntOfInt64(704), globalState), 4)
			(_nw317).ArraySet1((_2059_v33).Cardinality(), 5)
			(_nw317).ArraySet1(_2045_v25, 6)
			(_nw317).ArraySet1(_2045_v25, 7)
			(_nw317).ArraySet1(_2045_v25, 8)
			(_nw317).ArraySet1(_dafny.IntOfInt64(447), 9)
			(_nw317).ArraySet1(_2045_v25, 10)
			(_nw317).ArraySet1((_2060_v34).Cardinality(), 11)
			(_nw317).ArraySet1(_2045_v25, 12)
			(_nw317).ArraySet1(_2045_v25, 13)
			(_nw317).ArraySet1(_2045_v25, 14)
			(_nw317).ArraySet1(_2045_v25, 15)
			(_nw317).ArraySet1(_2045_v25, 16)
			(_nw317).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hvirhay")).Cardinality()), 17)
			(_nw317).ArraySet1(_2045_v25, 18)
			_2061_v35 = _nw317
			var _2062_v36 _dafny.Map
			_ = _2062_v36
			_2062_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.IntOfInt64(918)).Times(_2045_v25)), _2061_v35)
			_2062_v36 = (_2062_v36).Update((_dafny.Zero).Minus(_2045_v25), _2061_v35)
			var _2063_v37 *C9
			_ = _2063_v37
			var _nw318 *C9 = New_C9_()
			_ = _nw318
			_nw318.Ctor__()
			_2063_v37 = _nw318
			var _2064_v38 _dafny.CodePoint
			_ = _2064_v38
			_2064_v38 = _dafny.CodePoint('x')
			var _2065_v39 _dafny.Set
			_ = _2065_v39
			_2065_v39 = _dafny.SetOf(_2064_v38, _2064_v38, _2064_v38, Companion_Default___.Fm38(_2044_v24, globalState), _dafny.CodePoint('a'))
			var _2066_v41 _dafny.Sequence
			_ = _2066_v41
			_2066_v41 = _dafny.SeqOf((_2065_v39).Union(func() _dafny.Set {
				var _coll79 = _dafny.NewBuilder()
				_ = _coll79
				for _iter81 := _dafny.Iterate((_2065_v39).Elements()); ; {
					_compr_79, _ok81 := _iter81()
					if !_ok81 {
						break
					}
					var _2067_v40 _dafny.CodePoint
					_2067_v40 = interface{}(_compr_79).(_dafny.CodePoint)
					if (_2065_v39).Contains(_2067_v40) {
						_coll79.Add(_2067_v40)
					}
				}
				return _coll79.ToSet()
			}()))
			_2065_v39 = (_2066_v41).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_2045_v25, _2045_v25), _dafny.IntOfUint32((_2066_v41).Cardinality()))).Uint32()).(_dafny.Set)
			_2045_v25 = _2045_v25
		} else {
			var _2068___mcc_h3 _dafny.Int = _source29.Get_().(D1_DC2).Cf2
			_ = _2068___mcc_h3
			var _2069_cf2 _dafny.Int = _2068___mcc_h3
			_ = _2069_cf2
			var _2070_v42 _dafny.Array
			_ = _2070_v42
			var _nw319 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
			_ = _nw319
			_2070_v42 = _nw319
			var _2071_v43 _dafny.Array
			_ = _2071_v43
			var _nw320 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw320
			_2071_v43 = _nw320
			var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))
			_ = _index355
			(_2070_v42).ArraySet1(_2071_v43, (_index355).Int())
			var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))
			_ = _index356
			var _rhs422 _dafny.Int = _2069_cf2
			_ = _rhs422
			var _rhs423 _dafny.Array = _2071_v43
			_ = _rhs423
			var _rhs424 bool = true
			_ = _rhs424
			var _lhs290 _dafny.Array = _2070_v42
			_ = _lhs290
			var _lhs291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))
			_ = _lhs291
			_2045_v25 = _rhs422
			(_lhs290).ArraySet1(_rhs423, (_lhs291).Int())
			_2044_v24 = _rhs424
			if _2044_v24 {
				var _2072_v44 _dafny.Map
				_ = _2072_v44
				_2072_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_2044_v24) || (_2044_v24), Companion_Default___.Fm5(globalState))
				var _2073_v45 D1
				_ = _2073_v45
				_2073_v45 = Companion_D1_.Create_DC3_(_2017_v0, _2044_v24)
				_2072_v44 = (_2072_v44).Update(_2044_v24, _dafny.Companion_Sequence_.Concatenate((_2073_v45).Dtor_cf3(), _2017_v0))
				var _2074_v46 _dafny.Sequence
				_ = _2074_v46
				_2074_v46 = _dafny.SeqOf(_2044_v24)
				var _2075_v47 _dafny.Map
				_ = _2075_v47
				_2075_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2074_v46, _2044_v24)
				_2075_v47 = (_2075_v47).Merge(_2075_v47)
				var _2076_v48 _dafny.Sequence
				_ = _2076_v48
				_2076_v48 = _dafny.SeqOf(_2046_v26, _2046_v26)
				var _2077_v49 _dafny.Sequence
				_ = _2077_v49
				_2077_v49 = _dafny.SeqOf((_2076_v48).Select((Companion_Default___.SafeIndex(_2045_v25, _dafny.IntOfUint32((_2076_v48).Cardinality()))).Uint32()).(_dafny.MultiSet))
				(globalState).F4 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2076_v48, _2076_v48), _2076_v48)).Cardinality())
				var _2078_v50 *C10
				_ = _2078_v50
				var _nw321 *C10 = New_C10_()
				_ = _nw321
				_nw321.Ctor__()
				_2078_v50 = _nw321
				_2078_v50 = _2078_v50
				(globalState).F15 = _2069_cf2
			} else {
				var _2079_v51 _dafny.Array
				_ = _2079_v51
				var _len0_52 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_52
				var _nw322 _dafny.Array
				_ = _nw322
				if _len0_52.Cmp(_dafny.Zero) == 0 {
					_nw322 = _dafny.NewArray(_len0_52)
				} else {
					var _init52 func(_dafny.Int) _dafny.Set = (func(_2080_v24 bool) func(_dafny.Int) _dafny.Set {
						return func(_2081_i2 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_2080_v24)
						}
					})(_2044_v24)
					_ = _init52
					var _element0_52 = _init52(_dafny.Zero)
					_ = _element0_52
					_nw322 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
					(_nw322).ArraySet1(_element0_52, 0)
					var _nativeLen0_52 = (_len0_52).Int()
					_ = _nativeLen0_52
					for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
						(_nw322).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
					}
				}
				_2079_v51 = _nw322
				_2079_v51 = _2079_v51
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_2070_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))).Int()))
				_ = _arr2
				var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_dafny.ArrayCastTo((_2070_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))).Int()))), 0))
				_ = _index357
				(_arr2).ArraySet1(false, (_index357).Int())
				var _2082_v52 _dafny.Array
				_ = _2082_v52
				var _nw323 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw323
				_2082_v52 = _nw323
				var _2083_v53 _dafny.Set
				_ = _2083_v53
				_2083_v53 = _dafny.SetOf(_2082_v52, _2082_v52)
				var _2084_v54 _dafny.Sequence
				_ = _2084_v54
				_2084_v54 = _dafny.SeqOf((_2083_v53).Cardinality())
				var _arr3 _dafny.Array = _dafny.ArrayCastTo((_2070_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))).Int()))
				_ = _arr3
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_dafny.ArrayCastTo((_2070_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))).Int()))), 0))
				_ = _index358
				(_arr3).ArraySet1(((_2069_cf2).Times(_2069_cf2)).Cmp((_2084_v54).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.IntOfUint32((_2084_v54).Cardinality()))).Uint32()).(_dafny.Int)) > 0, (_index358).Int())
				var _2085_v55 _dafny.CodePoint
				_ = _2085_v55
				_2085_v55 = _dafny.CodePoint('i')
				var _2086_v56 _dafny.MultiSet
				_ = _2086_v56
				_2086_v56 = _dafny.MultiSetOf(_2085_v55)
				var _arr4 _dafny.Array = _dafny.ArrayCastTo((_2070_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))).Int()))
				_ = _arr4
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_dafny.ArrayCastTo((_2070_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))).Int()))), 0))
				_ = _index359
				(_arr4).ArraySet1((_2086_v56).IsDisjointFrom(_2086_v56), (_index359).Int())
				(globalState).F13 = Companion_Default___.Fm17((_dafny.Zero).Minus(_2045_v25), globalState)
				var _arr5 _dafny.Array = _dafny.ArrayCastTo((_2070_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))).Int()))
				_ = _arr5
				var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_dafny.ArrayCastTo((_2070_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2070_v42), 0))).Int()))), 0))
				_ = _index360
				(_arr5).ArraySet1(_2044_v24, (_index360).Int())
			}
			var _2087_v57 _dafny.Map
			_ = _2087_v57
			_2087_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2044_v24, _2044_v24)
			var _2088_v58 _dafny.Map
			_ = _2088_v58
			_2088_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_2087_v57).Contains(!(_2044_v24)) {
					return (_2087_v57).Get(!(_2044_v24)).(bool)
				}
				return _2044_v24
			})(), _2087_v57)
			var _2089_v59 _dafny.Sequence
			_ = _2089_v59
			_2089_v59 = _dafny.SeqOf((_this).Fm1(_2088_v58, _2069_cf2, globalState))
			var _2090_v60 _dafny.Map
			_ = _2090_v60
			_2090_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2089_v59, _2069_cf2)
			var _2091_v61 _dafny.CodePoint
			_ = _2091_v61
			_2091_v61 = _dafny.CodePoint('o')
			var _2092_v62 _dafny.Map
			_ = _2092_v62
			_2092_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2091_v61, _2046_v26)
			var _2093_v64 _dafny.Sequence
			_ = _2093_v64
			_2093_v64 = _dafny.SeqOf(_2092_v62, func() _dafny.Map {
				var _coll80 = _dafny.NewMapBuilder()
				_ = _coll80
				for _iter82 := _dafny.Iterate((_2017_v0).Elements()); ; {
					_compr_80, _ok82 := _iter82()
					if !_ok82 {
						break
					}
					var _2094_v63 _dafny.CodePoint
					_2094_v63 = interface{}(_compr_80).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_2017_v0, _2094_v63) {
						_coll80.Add(_2094_v63, _2046_v26)
					}
				}
				return _coll80.ToMap()
			}(), _2092_v62)
			var _2095_v65 _dafny.Sequence
			_ = _2095_v65
			_2095_v65 = _dafny.SeqOf(_2069_cf2, ((_2090_v60).Cardinality()).Minus(Companion_Default___.Fm17(_2045_v25, globalState)), ((_2093_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_2044_v24, _2044_v24)).Cardinality()), _dafny.IntOfUint32((_2093_v64).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
			_2095_v65 = _2095_v65
			_2044_v24 = ((_dafny.Zero).Minus(_2069_cf2)).Cmp(_dafny.IntOfInt64(678)) >= 0
		}
		var _2096_v66 _dafny.Array
		_ = _2096_v66
		var _nw324 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw324
		_2096_v66 = _nw324
		for _iter83 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2096_v66), 0))); ; {
			_guard_loop_2, _ok83 := _iter83()
			if !_ok83 {
				break
			}
			var _2097_i3 _dafny.Int
			_2097_i3 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_2097_i3).Sign() != -1) && ((_2097_i3).Cmp(_dafny.ArrayLen((_2096_v66), 0)) < 0)) {
				(_2096_v66).ArraySet1((_2097_i3).Minus((_2045_v25).Plus(_2045_v25)), (_2097_i3).Int())
			}
		}
		var _2098_v67 _dafny.Array
		_ = _2098_v67
		var _nw325 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
		_ = _nw325
		_2098_v67 = _nw325
		var _2099_v68 D0
		_ = _2099_v68
		_2099_v68 = Companion_D0_.Create_DC0_(_2098_v67)
		var _pat_let_tv42 = _2098_v67
		_ = _pat_let_tv42
		var _pat_let_tv43 = _2098_v67
		_ = _pat_let_tv43
		var _2100_v69 _dafny.Array
		_ = _2100_v69
		var _nwElement0_78 D0 = func(_pat_let35_0 D0) D0 {
			return func(_2101_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let36_0 _dafny.Array) D0 {
					return func(_2102_dt__update_hcf0_h0 _dafny.Array) D0 {
						return Companion_D0_.Create_DC0_(_2102_dt__update_hcf0_h0)
					}(_pat_let36_0)
				}(_pat_let_tv42)
			}(_pat_let35_0)
		}(_2099_v68)
		_ = _nwElement0_78
		var _nw326 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(5))
		_ = _nw326
		(_nw326).ArraySet1(_nwElement0_78, 0)
		(_nw326).ArraySet1(func(_pat_let37_0 D0) D0 {
			return func(_2103_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let38_0 _dafny.Array) D0 {
					return func(_2104_dt__update_hcf0_h1 _dafny.Array) D0 {
						return Companion_D0_.Create_DC0_(_2104_dt__update_hcf0_h1)
					}(_pat_let38_0)
				}(_pat_let_tv43)
			}(_pat_let37_0)
		}(_2099_v68), 1)
		(_nw326).ArraySet1(_2099_v68, 2)
		(_nw326).ArraySet1(_2099_v68, 3)
		(_nw326).ArraySet1(_2099_v68, 4)
		_2100_v69 = _nw326
		var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_2100_v69), 0))
		_ = _index361
		(_2100_v69).ArraySet1(Companion_D0_.Create_DC0_(_2098_v67), (_index361).Int())
		var _2105_v70 _dafny.Array
		_ = _2105_v70
		var _nw327 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
		_ = _nw327
		_2105_v70 = _nw327
		var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_2105_v70), 0))
		_ = _index362
		(_2105_v70).ArraySet1(_2096_v66, (_index362).Int())
		var _2106_v71 D6
		_ = _2106_v71
		_2106_v71 = Companion_D6_.Create_DC17_(_2044_v24, !(_2044_v24), _2044_v24)
		var _2107_v72 _dafny.Sequence
		_ = _2107_v72
		_2107_v72 = _dafny.SeqOf(_2106_v71, _2106_v71)
		var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_2096_v66), 0))
		_ = _index363
		(_2096_v66).ArraySet1(_dafny.IntOfUint32((_2107_v72).Cardinality()), (_index363).Int())
		var _2108_v73 _dafny.Map
		_ = _2108_v73
		_2108_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if _2044_v24 {
				return _2096_v66
			}
			return _2096_v66
		})(), _2045_v25)
		var _2109_v74 _dafny.MultiSet
		_ = _2109_v74
		_2109_v74 = _dafny.MultiSetOf(_2044_v24, _2044_v24)
		var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_2100_v69), 0))
		_ = _index364
		var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_2105_v70), 0))
		_ = _index365
		var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_2096_v66), 0))
		_ = _index366
		var _rhs425 D0 = _2099_v68
		_ = _rhs425
		var _rhs426 _dafny.Array = _2096_v66
		_ = _rhs426
		var _rhs427 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if _2044_v24 {
				return _2045_v25
			}
			return (_dafny.Zero).Minus(_2045_v25)
		})(), Companion_Default___.SafeModuloInt((_2109_v74).Cardinality(), _2045_v25))
		_ = _rhs427
		var _rhs428 _dafny.Map = _2108_v73
		_ = _rhs428
		var _lhs292 _dafny.Array = _2100_v69
		_ = _lhs292
		var _lhs293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_2100_v69), 0))
		_ = _lhs293
		var _lhs294 _dafny.Array = _2105_v70
		_ = _lhs294
		var _lhs295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_2105_v70), 0))
		_ = _lhs295
		var _lhs296 _dafny.Array = _2096_v66
		_ = _lhs296
		var _lhs297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_2096_v66), 0))
		_ = _lhs297
		(_lhs292).ArraySet1(_rhs425, (_lhs293).Int())
		(_lhs294).ArraySet1(_rhs426, (_lhs295).Int())
		(_lhs296).ArraySet1(_rhs427, (_lhs297).Int())
		_2108_v73 = _rhs428
		var _2110_v75 _dafny.Int
		_ = _2110_v75
		var _2111_v76 _dafny.Map
		_ = _2111_v76
		var _2112_v77 _dafny.Int
		_ = _2112_v77
		var _2113_v78 _dafny.Array
		_ = _2113_v78
		var _out15 _dafny.Int
		_ = _out15
		var _out16 _dafny.Map
		_ = _out16
		var _out17 _dafny.Int
		_ = _out17
		var _out18 _dafny.Array
		_ = _out18
		_out15, _out16, _out17, _out18 = (_this).M3(globalState)
		_2110_v75 = _out15
		_2111_v76 = _out16
		_2112_v77 = _out17
		_2113_v78 = _out18
	}
}
func (_this *C11) M3(globalState *GlobalState) (_dafny.Int, _dafny.Map, _dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r3
		var _2114_v0 *C9
		_ = _2114_v0
		var _nw328 *C9 = New_C9_()
		_ = _nw328
		_nw328.Ctor__()
		_2114_v0 = _nw328
		var _2115_v1 _dafny.Int
		_ = _2115_v1
		_2115_v1 = _dafny.IntOfInt64(196)
		var _2116_v2 D4
		_ = _2116_v2
		_2116_v2 = Companion_D4_.Create_DC12_((_dafny.Zero).Minus(_2115_v1))
		var _2117_v3 _dafny.Map
		_ = _2117_v3
		_2117_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2115_v1, (Companion_Default___.Fm56((_dafny.Zero).Minus(_2115_v1), (_2116_v2).Dtor_cf16(), globalState)).Dtor_cf35())
		_2117_v3 = (_2117_v3).Update(_2115_v1, true)
		var _2118_v4 bool
		_ = _2118_v4
		_2118_v4 = false
		var _2119_v5 D20
		_ = _2119_v5
		_2119_v5 = Companion_D20_.Create_DC55_((_dafny.Zero).Minus(_2115_v1), _2115_v1, _2118_v4, !(_2118_v4))
		_2118_v4 = (_2119_v5).Dtor_cf87()
		var _2120_v6 _dafny.Sequence
		_ = _2120_v6
		_2120_v6 = _dafny.UnicodeSeqOfUtf8Bytes("mxqw")
		var _2121_v7 _dafny.Map
		_ = _2121_v7
		_2121_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2118_v4, _2115_v1)
		var _2122_v8 _dafny.Sequence
		_ = _2122_v8
		_2122_v8 = _dafny.SeqOf(_2118_v4, (_2114_v0).Fm8((_2121_v7).Cardinality(), _2118_v4, _dafny.IntOfInt64(887), globalState))
		var _2123_v9 _dafny.Set
		_ = _2123_v9
		_2123_v9 = _dafny.SetOf(_2115_v1)
		var _2124_v10 D1
		_ = _2124_v10
		_2124_v10 = Companion_D1_.Create_DC3_(Companion_Default___.Fm30(_dafny.IntOfInt64(557), !((_2122_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm5(globalState)).Cardinality()), _dafny.IntOfUint32((_2122_v8).Cardinality()))).Uint32()).(bool)), _2123_v9, false, globalState), false)
		_2120_v6 = (_2124_v10).Dtor_cf3()
		var _2125_v11 _dafny.Map
		_ = _2125_v11
		_2125_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2115_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-67))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg118 _dafny.Int) interface{} {
				return coer118(arg118)
			}
		}(func(_2126_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _2120_v6)).Cardinality()))
		_2125_v11 = (_2125_v11).Update(_2115_v1, _2115_v1)
		var _2127_v12 _dafny.Array
		_ = _2127_v12
		var _nw329 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw329
		_2127_v12 = _nw329
		var _2128_v13 _dafny.MultiSet
		_ = _2128_v13
		_2128_v13 = _dafny.MultiSetOf(_2118_v4, _2118_v4)
		var _2129_v14 _dafny.Map
		_ = _2129_v14
		_2129_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2128_v13).Cardinality(), _2127_v12)
		var _2130_v15 _dafny.Array
		_ = _2130_v15
		var _nwElement0_79 _dafny.Array = _2127_v12
		_ = _nwElement0_79
		var _nw330 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(12))
		_ = _nw330
		(_nw330).ArraySet1(_nwElement0_79, 0)
		(_nw330).ArraySet1(_2127_v12, 1)
		(_nw330).ArraySet1(_2127_v12, 2)
		(_nw330).ArraySet1(_2127_v12, 3)
		(_nw330).ArraySet1(_2127_v12, 4)
		(_nw330).ArraySet1(_2127_v12, 5)
		(_nw330).ArraySet1(_2127_v12, 6)
		(_nw330).ArraySet1(_2127_v12, 7)
		(_nw330).ArraySet1(_2127_v12, 8)
		(_nw330).ArraySet1((func() _dafny.Array {
			if (_2129_v14).Contains(_dafny.IntOfInt64(456)) {
				return (_2129_v14).Get(_dafny.IntOfInt64(456)).(_dafny.Array)
			}
			return _2127_v12
		})(), 9)
		(_nw330).ArraySet1(_2127_v12, 10)
		(_nw330).ArraySet1(_2127_v12, 11)
		_2130_v15 = _nw330
		var _2131_v16 _dafny.Sequence
		_ = _2131_v16
		_2131_v16 = _dafny.SeqOf(_2130_v15, _2130_v15, _2130_v15, _2130_v15)
		var _2132_v17 _dafny.Array
		_ = _2132_v17
		var _nwElement0_80 _dafny.Array = _2130_v15
		_ = _nwElement0_80
		var _nw331 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(25))
		_ = _nw331
		(_nw331).ArraySet1(_nwElement0_80, 0)
		(_nw331).ArraySet1(_2130_v15, 1)
		(_nw331).ArraySet1(_2130_v15, 2)
		(_nw331).ArraySet1(_2130_v15, 3)
		(_nw331).ArraySet1(_2130_v15, 4)
		(_nw331).ArraySet1(_2130_v15, 5)
		(_nw331).ArraySet1(_2130_v15, 6)
		(_nw331).ArraySet1(_2130_v15, 7)
		(_nw331).ArraySet1(_2130_v15, 8)
		(_nw331).ArraySet1(_2130_v15, 9)
		(_nw331).ArraySet1(_2130_v15, 10)
		(_nw331).ArraySet1(_2130_v15, 11)
		(_nw331).ArraySet1(_2130_v15, 12)
		(_nw331).ArraySet1(_2130_v15, 13)
		(_nw331).ArraySet1(_2130_v15, 14)
		(_nw331).ArraySet1(_2130_v15, 15)
		(_nw331).ArraySet1(_2130_v15, 16)
		(_nw331).ArraySet1(_2130_v15, 17)
		(_nw331).ArraySet1(_2130_v15, 18)
		(_nw331).ArraySet1(_2130_v15, 19)
		(_nw331).ArraySet1(_2130_v15, 20)
		(_nw331).ArraySet1(_2130_v15, 21)
		(_nw331).ArraySet1(_2130_v15, 22)
		(_nw331).ArraySet1(_2130_v15, 23)
		(_nw331).ArraySet1((_2131_v16).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_2115_v1), _dafny.IntOfUint32((_2131_v16).Cardinality()))).Uint32()).(_dafny.Array), 24)
		_2132_v17 = _nw331
		var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_2132_v17), 0))
		_ = _index367
		(_2132_v17).ArraySet1(_2130_v15, (_index367).Int())
		var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_2132_v17), 0))
		_ = _index368
		(_2132_v17).ArraySet1(_2130_v15, (_index368).Int())
		var _2133_v18 _dafny.Array
		_ = _2133_v18
		var _len0_53 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_53
		var _nw332 _dafny.Array
		_ = _nw332
		if _len0_53.Cmp(_dafny.Zero) == 0 {
			_nw332 = _dafny.NewArray(_len0_53)
		} else {
			var _init53 func(_dafny.Int) _dafny.CodePoint = func(_2134_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}
			_ = _init53
			var _element0_53 = _init53(_dafny.Zero)
			_ = _element0_53
			_nw332 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
			(_nw332).ArraySet1CodePoint(_element0_53, 0)
			var _nativeLen0_53 = (_len0_53).Int()
			_ = _nativeLen0_53
			for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
				(_nw332).ArraySet1CodePoint(_init53(_dafny.IntOf(_i0_53)), _i0_53)
			}
		}
		_2133_v18 = _nw332
		var _2135_v19 _dafny.Map
		_ = _2135_v19
		_2135_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm30(_2115_v1, _2118_v4, _2123_v9, !(_2118_v4), globalState), _2133_v18)
		r0 = ((_2135_v19).Merge((_2135_v19).Merge(_2135_v19))).Cardinality()
		var _2136_v20 _dafny.Map
		_ = _2136_v20
		_2136_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(_2115_v1, _2115_v1)), _2118_v4)
		var _2137_v21 _dafny.Sequence
		_ = _2137_v21
		_2137_v21 = _dafny.SeqOf(_2136_v20, _2136_v20)
		r1 = (_2137_v21).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(166)).Minus(_2115_v1), _dafny.IntOfUint32((_2137_v21).Cardinality()))).Uint32()).(_dafny.Map)
		r2 = _2115_v1
		var _2138_v22 _dafny.Array
		_ = _2138_v22
		var _len0_54 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_54
		var _nw333 _dafny.Array
		_ = _nw333
		if _len0_54.Cmp(_dafny.Zero) == 0 {
			_nw333 = _dafny.NewArray(_len0_54)
		} else {
			var _init54 func(_dafny.Int) _dafny.Map = (func(_2139_v7 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_2140_i2 _dafny.Int) _dafny.Map {
					return _2139_v7
				}
			})(_2121_v7)
			_ = _init54
			var _element0_54 = _init54(_dafny.Zero)
			_ = _element0_54
			_nw333 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
			(_nw333).ArraySet1(_element0_54, 0)
			var _nativeLen0_54 = (_len0_54).Int()
			_ = _nativeLen0_54
			for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
				(_nw333).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
			}
		}
		_2138_v22 = _nw333
		r3 = _2138_v22
		return r0, r1, r2, r3
	}
}
func (_this *C11) M4(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var _2141_i0 _dafny.Int
		_ = _2141_i0
		_2141_i0 = _dafny.Zero
		{
			for p3 {
				{
					if (_2141_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_2141_i0 = (_2141_i0).Plus(_dafny.One)
					var _2142_v0 bool
					_ = _2142_v0
					_2142_v0 = true
					var _2143_v1 _dafny.CodePoint
					_ = _2143_v1
					_2143_v1 = _dafny.CodePoint('w')
					var _2144_v2 D5
					_ = _2144_v2
					_2144_v2 = Companion_D5_.Create_DC14_(_2143_v1, p1, p0, p3)
					_2142_v0 = (p1) && ((_2144_v2).Dtor_cf21())
					var _2145_v3 _dafny.Int
					_ = _2145_v3
					var _2146_v4 _dafny.Map
					_ = _2146_v4
					var _2147_v5 _dafny.Int
					_ = _2147_v5
					var _2148_v6 _dafny.Array
					_ = _2148_v6
					var _out19 _dafny.Int
					_ = _out19
					var _out20 _dafny.Map
					_ = _out20
					var _out21 _dafny.Int
					_ = _out21
					var _out22 _dafny.Array
					_ = _out22
					_out19, _out20, _out21, _out22 = (_this).M3(globalState)
					_2145_v3 = _out19
					_2146_v4 = _out20
					_2147_v5 = _out21
					_2148_v6 = _out22
					if true {
						(globalState).F13 = _2147_v5
						(globalState).F4 = (_2147_v5).Minus((_this).Fm2(_2147_v5, globalState))
						var _2149_v7 _dafny.Sequence
						_ = _2149_v7
						_2149_v7 = _dafny.UnicodeSeqOfUtf8Bytes("sxs")
						var _2150_v8 _dafny.Map
						_ = _2150_v8
						_2150_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2147_v5, _dafny.IntOfInt64(327))
						var _rhs429 _dafny.Int = Companion_Default___.Fm17(p0, globalState)
						_ = _rhs429
						var _rhs430 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("wse")
						_ = _rhs430
						var _rhs431 _dafny.Int = (_2145_v3).Plus((_dafny.IntOfInt64(723)).Times(Companion_Default___.Fm17(_2145_v3, globalState)))
						_ = _rhs431
						var _rhs432 _dafny.Int = ((_dafny.Zero).Minus(p0)).Times(p0)
						_ = _rhs432
						var _rhs433 _dafny.Map = _2150_v8
						_ = _rhs433
						var _lhs298 *GlobalState = globalState
						_ = _lhs298
						var _lhs299 *GlobalState = globalState
						_ = _lhs299
						var _lhs300 *GlobalState = globalState
						_ = _lhs300
						_lhs298.F4 = _rhs429
						_2149_v7 = _rhs430
						_lhs299.F13 = _rhs431
						_lhs300.F13 = _rhs432
						_2150_v8 = _rhs433
						var _2151_v9 *C0
						_ = _2151_v9
						var _nw334 *C0 = New_C0_()
						_ = _nw334
						_nw334.Ctor__()
						_2151_v9 = _nw334
						var _2152_v10 _dafny.Sequence
						_ = _2152_v10
						_2152_v10 = _dafny.SeqOf(p1, _2142_v0, p1)
						var _2153_v11 _dafny.Sequence
						_ = _2153_v11
						_2153_v11 = _dafny.SeqOf(_dafny.IntOfUint32((_2152_v10).Cardinality()))
						var _2154_v12 _dafny.Sequence
						_ = _2154_v12
						_2154_v12 = _dafny.SeqOf(_2153_v11)
						var _2155_v13 _dafny.Sequence
						_ = _2155_v13
						_2155_v13 = _dafny.SeqOf(_2149_v7, _2149_v7)
						var _2156_v14 _dafny.Array
						_ = _2156_v14
						var _nwElement0_81 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("gif")
						_ = _nwElement0_81
						var _nw335 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(21))
						_ = _nw335
						(_nw335).ArraySet1(_nwElement0_81, 0)
						(_nw335).ArraySet1(_2149_v7, 1)
						(_nw335).ArraySet1(_2149_v7, 2)
						(_nw335).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm5(globalState), (Companion_Default___.SafeIndex(_2147_v5, _dafny.IntOfUint32((Companion_Default___.Fm5(globalState)).Cardinality()))).Uint32(), _2143_v1), 3)
						(_nw335).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-717))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg119 _dafny.Int) interface{} {
								return coer119(arg119)
							}
						}((func(_2157_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2158_i1 _dafny.Int) _dafny.CodePoint {
								return _2157_v1
							}
						})(_2143_v1))), 4)
						(_nw335).ArraySet1(_2149_v7, 5)
						(_nw335).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(343))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg120 _dafny.Int) interface{} {
								return coer120(arg120)
							}
						}((func(_2159_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2160_i2 _dafny.Int) _dafny.CodePoint {
								return _2159_v1
							}
						})(_2143_v1))), 6)
						(_nw335).ArraySet1((_2151_v9).Fm28(_2154_v12, globalState), 7)
						(_nw335).ArraySet1(_2149_v7, 8)
						(_nw335).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2149_v7, (Companion_Default___.SafeIndex(_2145_v3, _dafny.IntOfUint32((_2149_v7).Cardinality()))).Uint32(), _dafny.CodePoint('n')), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2149_v7, (Companion_Default___.SafeIndex(_2145_v3, _dafny.IntOfUint32((_2149_v7).Cardinality()))).Uint32(), _dafny.CodePoint('n'))).Cardinality()))).Uint32(), _dafny.CodePoint('x')), 9)
						(_nw335).ArraySet1(_2149_v7, 10)
						(_nw335).ArraySet1(_2149_v7, 11)
						(_nw335).ArraySet1(_2149_v7, 12)
						(_nw335).ArraySet1(_2149_v7, 13)
						(_nw335).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jc"), 14)
						(_nw335).ArraySet1(_2149_v7, 15)
						(_nw335).ArraySet1(_2149_v7, 16)
						(_nw335).ArraySet1(_2149_v7, 17)
						(_nw335).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("befkyx"), 18)
						(_nw335).ArraySet1((_2155_v13).Select((Companion_Default___.SafeIndex(_2147_v5, _dafny.IntOfUint32((_2155_v13).Cardinality()))).Uint32()).(_dafny.Sequence), 19)
						(_nw335).ArraySet1(_2149_v7, 20)
						_2156_v14 = _nw335
						var _2161_v15 _dafny.Array
						_ = _2161_v15
						var _nw336 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
						_ = _nw336
						_2161_v15 = _nw336
						var _2162_v16 _dafny.Sequence
						_ = _2162_v16
						_2162_v16 = _dafny.SeqOf(_2161_v15, _2161_v15)
						var _2163_v17 _dafny.Map
						_ = _2163_v17
						_2163_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2156_v14, _2162_v16)
						_2163_v17 = (_2163_v17).Update(_2156_v14, _dafny.Companion_Sequence_.Concatenate(_2162_v16, _2162_v16))
					} else {
						var _2164_v18 _dafny.Map
						_ = _2164_v18
						_2164_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_2147_v5)), _2147_v5)
						(globalState).F4 = (func() _dafny.Int {
							if (_2164_v18).Contains((_2147_v5).Plus(_2145_v3)) {
								return (_2164_v18).Get((_2147_v5).Plus(_2145_v3)).(_dafny.Int)
							}
							return p0
						})()
						var _2165_v19 _dafny.Set
						_ = _2165_v19
						_2165_v19 = _dafny.SetOf(!(p1))
						var _2166_v20 _dafny.Array
						_ = _2166_v20
						var _nwElement0_82 bool = _2142_v0
						_ = _nwElement0_82
						var _nw337 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(8))
						_ = _nw337
						(_nw337).ArraySet1(_nwElement0_82, 0)
						(_nw337).ArraySet1(_2142_v0, 1)
						(_nw337).ArraySet1(p1, 2)
						(_nw337).ArraySet1((_2165_v19).IsDisjointFrom(_2165_v19), 3)
						(_nw337).ArraySet1(_2142_v0, 4)
						(_nw337).ArraySet1(p3, 5)
						(_nw337).ArraySet1(p3, 6)
						(_nw337).ArraySet1(p1, 7)
						_2166_v20 = _nw337
						var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_2166_v20), 0))
						_ = _index369
						(_2166_v20).ArraySet1(_2142_v0, (_index369).Int())
						var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_2166_v20), 0))
						_ = _index370
						(_2166_v20).ArraySet1(p1, (_index370).Int())
						var _2167_v21 _dafny.Map
						_ = _2167_v21
						_2167_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2142_v0, !(p3))
						_2142_v0 = !(!((_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2167_v21), p0, globalState)))
						_2145_v3 = p0
						var _2168_v23 _dafny.Sequence
						_ = _2168_v23
						_2168_v23 = _dafny.SeqOf(_2147_v5)
						_2164_v18 = ((_2164_v18).Merge(func() _dafny.Map {
							var _coll81 = _dafny.NewMapBuilder()
							_ = _coll81
							for _iter84 := _dafny.Iterate((_2168_v23).Elements()); ; {
								_compr_81, _ok84 := _iter84()
								if !_ok84 {
									break
								}
								var _2169_v22 _dafny.Int
								_2169_v22 = interface{}(_compr_81).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_2168_v23, _2169_v22) {
									_coll81.Add((_2169_v22).Plus(_2147_v5), _2145_v3)
								}
							}
							return _coll81.ToMap()
						}())).Merge(_2164_v18)
					}
					var _2170_v24 _dafny.Array
					_ = _2170_v24
					var _nw338 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
					_ = _nw338
					_2170_v24 = _nw338
					var _2171_v25 _dafny.Sequence
					_ = _2171_v25
					_2171_v25 = _dafny.UnicodeSeqOfUtf8Bytes("ihvfprud")
					var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_2170_v24), 0))
					_ = _index371
					(_2170_v24).ArraySet1(_2171_v25, (_index371).Int())
					var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_2170_v24), 0))
					_ = _index372
					(_2170_v24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2171_v25, _2171_v25), (_index372).Int())
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		var _2172_v26 _dafny.Map
		_ = _2172_v26
		_2172_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
		var _2173_v27 _dafny.Map
		_ = _2173_v27
		_2173_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2172_v26)
		var _2174_v28 _dafny.Sequence
		_ = _2174_v28
		_2174_v28 = _dafny.SeqOf((_dafny.Zero).Minus(p0))
		var _2175_v29 _dafny.Map
		_ = _2175_v29
		_2175_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).Fm1(_2173_v27, _dafny.IntOfUint32((_2174_v28).Cardinality()), globalState))
		var _2176_v30 _dafny.Map
		_ = _2176_v30
		_2176_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2172_v26)
		var _2177_v31 _dafny.Sequence
		_ = _2177_v31
		_2177_v31 = _dafny.SeqOf(p1)
		var _2178_v32 _dafny.MultiSet
		_ = _2178_v32
		_2178_v32 = _dafny.MultiSetOf(_2177_v31)
		var _2179_v33 _dafny.Set
		_ = _2179_v33
		_2179_v33 = _dafny.SetOf((_this).Fm1(_2173_v27, (_this).Fm2((_2178_v32).Cardinality(), globalState), globalState))
		var _2180_v34 _dafny.CodePoint
		_ = _2180_v34
		_2180_v34 = _dafny.CodePoint('l')
		var _2181_v35 _dafny.Map
		_ = _2181_v35
		_2181_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2180_v34, _dafny.SetOf(p1))
		var _pat_let_tv44 = _2181_v35
		_ = _pat_let_tv44
		var _pat_let_tv45 = _2180_v34
		_ = _pat_let_tv45
		var _pat_let_tv46 = _2181_v35
		_ = _pat_let_tv46
		var _pat_let_tv47 = _2180_v34
		_ = _pat_let_tv47
		var _pat_let_tv48 = _2179_v33
		_ = _pat_let_tv48
		var _source30 D12 = func(_pat_let39_0 D12) D12 {
			return func(_2182_dt__update__tmp_h0 D12) D12 {
				return func(_pat_let40_0 _dafny.Set) D12 {
					return func(_2183_dt__update_hcf49_h0 _dafny.Set) D12 {
						return Companion_D12_.Create_DC32_(_2183_dt__update_hcf49_h0)
					}(_pat_let40_0)
				}((func() _dafny.Set {
					if (_pat_let_tv44).Contains(_pat_let_tv45) {
						return (_pat_let_tv46).Get(_pat_let_tv47).(_dafny.Set)
					}
					return _pat_let_tv48
				})())
			}(_pat_let39_0)
		}(Companion_D12_.Create_DC32_(_2179_v33))
		_ = _source30
		if _source30.Is_DC33() {
			var _2184___mcc_h0 bool = _source30.Get_().(D12_DC33).Cf50
			_ = _2184___mcc_h0
			var _2185_cf50 bool = _2184___mcc_h0
			_ = _2185_cf50
			(globalState).F13 = _dafny.IntOfInt64(461)
			var _2186_v36 D3
			_ = _2186_v36
			_2186_v36 = Companion_D3_.Create_DC10_(_dafny.IntOfInt64(468), p0)
			var _source31 D3 = _2186_v36
			_ = _source31
			if _source31.Is_DC10() {
				var _2187___mcc_h2 _dafny.Int = _source31.Get_().(D3_DC10).Cf13
				_ = _2187___mcc_h2
				var _2188___mcc_h3 _dafny.Int = _source31.Get_().(D3_DC10).Cf14
				_ = _2188___mcc_h3
				var _2189_cf14 _dafny.Int = _2188___mcc_h3
				_ = _2189_cf14
				var _2190_cf13 _dafny.Int = _2187___mcc_h2
				_ = _2190_cf13
				(globalState).F4 = p0
				var _2191_v37 _dafny.Sequence
				_ = _2191_v37
				_2191_v37 = _dafny.UnicodeSeqOfUtf8Bytes("xmyrhykkq")
				var _2192_v38 _dafny.Map
				_ = _2192_v38
				_2192_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2189_cf14, false)
				var _2193_v39 _dafny.Array
				_ = _2193_v39
				var _len0_55 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_55
				var _nw339 _dafny.Array
				_ = _nw339
				if _len0_55.Cmp(_dafny.Zero) == 0 {
					_nw339 = _dafny.NewArray(_len0_55)
				} else {
					var _init55 func(_dafny.Int) _dafny.Int = (func(_2194_cf14 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2195_i3 _dafny.Int) _dafny.Int {
							return (_2195_i3).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(78))).Uint32(), func(coer121 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg121 _dafny.Int) interface{} {
									return coer121(arg121)
								}
							}((func(_2196_cf14 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_2197_i4 _dafny.Int) _dafny.Int {
									return _2196_cf14
								}
							})(_2194_cf14)))).Cardinality()))
						}
					})(_2189_cf14)
					_ = _init55
					var _element0_55 = _init55(_dafny.Zero)
					_ = _element0_55
					_nw339 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
					(_nw339).ArraySet1(_element0_55, 0)
					var _nativeLen0_55 = (_len0_55).Int()
					_ = _nativeLen0_55
					for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
						(_nw339).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
					}
				}
				_2193_v39 = _nw339
				var _2198_v40 D2
				_ = _2198_v40
				_2198_v40 = Companion_D2_.Create_DC5_(_2193_v39)
				var _2199_v41 _dafny.Set
				_ = _2199_v41
				_2199_v41 = _dafny.SetOf(_2198_v40)
				var _rhs434 bool = Companion_Default___.Fm26(globalState)
				_ = _rhs434
				var _rhs435 bool = !((func() bool {
					if p1 {
						return _2185_cf50
					}
					return !(p3)
				})()) || ((func() bool {
					if (_2192_v38).Contains((_2179_v33).Cardinality()) {
						return (_2192_v38).Get((_2179_v33).Cardinality()).(bool)
					}
					return p3
				})())
				_ = _rhs435
				var _rhs436 bool = !_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm19((_2177_v31).Select((Companion_Default___.SafeIndex((_2199_v41).Cardinality(), _dafny.IntOfUint32((_2177_v31).Cardinality()))).Uint32()).(bool), _2185_cf50, globalState), (func() _dafny.CodePoint {
					if (_2177_v31).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.IntOfUint32((_2177_v31).Cardinality()))).Uint32()).(bool) {
						return _2180_v34
					}
					return _2180_v34
				})())
				_ = _rhs436
				var _rhs437 _dafny.Map = _2172_v26
				_ = _rhs437
				var _rhs438 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2191_v37, _2191_v37)
				_ = _rhs438
				_2185_cf50 = _rhs434
				_2185_cf50 = _rhs435
				_2185_cf50 = _rhs436
				_2172_v26 = _rhs437
				_2191_v37 = _rhs438
				var _2200_v42 _dafny.Array
				_ = _2200_v42
				var _nw340 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw340
				_2200_v42 = _nw340
				var _rhs439 _dafny.Array = _2200_v42
				_ = _rhs439
				var _rhs440 bool = (func() bool {
					if p3 {
						return !(p1) || (p3)
					}
					return _2185_cf50
				})()
				_ = _rhs440
				var _rhs441 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-409))).Uint32(), func(coer122 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg122 _dafny.Int) interface{} {
						return coer122(arg122)
					}
				}((func(_2201_cf13 _dafny.Int, _2202_v38 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_2203_i5 _dafny.Int) _dafny.Int {
						return (_2201_cf13).Minus((_2202_v38).Cardinality())
					}
				})(_2190_cf13, _2192_v38)))
				_ = _rhs441
				var _rhs442 _dafny.CodePoint = _2180_v34
				_ = _rhs442
				var _rhs443 bool = p3
				_ = _rhs443
				_2200_v42 = _rhs439
				_2185_cf50 = _rhs440
				_2174_v28 = _rhs441
				_2180_v34 = _rhs442
				_2185_cf50 = _rhs443
				_2191_v37 = _2191_v37
			} else {
				var _2204___mcc_h4 _dafny.CodePoint = _source31.Get_().(D3_DC9).Cf12
				_ = _2204___mcc_h4
				var _2205_cf12 _dafny.CodePoint = _2204___mcc_h4
				_ = _2205_cf12
				_2185_cf50 = Companion_Default___.Fm26(globalState)
				var _2206_v43 *C1
				_ = _2206_v43
				var _nw341 *C1 = New_C1_()
				_ = _nw341
				_nw341.Ctor__()
				_2206_v43 = _nw341
				var _2207_v44 _dafny.Sequence
				_ = _2207_v44
				_2207_v44 = _dafny.UnicodeSeqOfUtf8Bytes("qu")
				_2207_v44 = _2207_v44
				_2185_cf50 = (_2179_v33).IsProperSubsetOf(_dafny.SetOf(false))
			}
			var _2208_v45 _dafny.Map
			_ = _2208_v45
			_2208_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(p0, globalState), p0)
			_2208_v45 = (_2208_v45).Update(Companion_Default___.SafeModuloInt((_2174_v28).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2174_v28).Cardinality()))).Uint32()).(_dafny.Int), p0), (_dafny.Zero).Minus((_this).Fm2(p0, globalState)))
			if false {
				_2185_cf50 = !(true)
				var _2209_v46 _dafny.Sequence
				_ = _2209_v46
				_2209_v46 = _dafny.UnicodeSeqOfUtf8Bytes("auf")
				_2209_v46 = _2209_v46
				(globalState).F13 = p0
				var _2210_v47 *C10
				_ = _2210_v47
				var _nw342 *C10 = New_C10_()
				_ = _nw342
				_nw342.Ctor__()
				_2210_v47 = _nw342
				var _2211_v48 _dafny.MultiSet
				_ = _2211_v48
				_2211_v48 = _dafny.MultiSetOf(_dafny.IntOfInt64(388))
				(globalState).F7 = Companion_Default___.SafeModuloInt(p0, (p0).Plus((func() _dafny.Int {
					if (_2211_v48).Contains(p0) {
						return (_2211_v48).Multiplicity(p0)
					}
					return p0
				})()))
			} else {
				var _2212_v49 _dafny.MultiSet
				_ = _2212_v49
				_2212_v49 = _dafny.MultiSetOf(!(p3))
				var _2213_v50 _dafny.MultiSet
				_ = _2213_v50
				_2213_v50 = _dafny.MultiSetOf((_2212_v49).Update(_2185_cf50, Companion_Default___.Abs(p0)), _dafny.MultiSetOf(p3, p3, p3, _2185_cf50))
				var _2214_v51 _dafny.Array
				_ = _2214_v51
				var _nwElement0_83 bool = _2185_cf50
				_ = _nwElement0_83
				var _nw343 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(4))
				_ = _nw343
				(_nw343).ArraySet1(_nwElement0_83, 0)
				(_nw343).ArraySet1((_2213_v50).IsDisjointFrom(_2213_v50), 1)
				(_nw343).ArraySet1(p1, 2)
				(_nw343).ArraySet1(p1, 3)
				_2214_v51 = _nw343
				var _2215_v52 D6
				_ = _2215_v52
				_2215_v52 = Companion_D6_.Create_DC16_(_dafny.SeqOf(p0))
				var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))
				_ = _index373
				(_2214_v51).ArraySet1(_dafny.Companion_Sequence_.Equal((_2215_v52).Dtor_cf23(), _2174_v28), (_index373).Int())
				var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))
				_ = _index374
				(_2214_v51).ArraySet1(p3, (_index374).Int())
				var _2216_v53 _dafny.Set
				_ = _2216_v53
				_2216_v53 = _dafny.SetOf((_dafny.IntOfUint32((_2174_v28).Cardinality())).Plus(p0), _dafny.IntOfInt64(655))
				_2216_v53 = ((_2216_v53).Union(_2216_v53)).Intersection(_2216_v53)
				var _2217_v54 D7
				_ = _2217_v54
				_2217_v54 = Companion_D7_.Create_DC22_(p3, p0, (_2214_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))).Int()).(bool), p0)
				var _2218_v55 _dafny.MultiSet
				_ = _2218_v55
				_2218_v55 = _dafny.MultiSetOf(_2217_v54, Companion_D7_.Create_DC22_(p3, p0, p1, p0), _2217_v54, Companion_D7_.Create_DC22_(p3, p0, true, p0), _2217_v54)
				var _2219_v56 D11
				_ = _2219_v56
				_2219_v56 = Companion_D11_.Create_DC30_(_2185_cf50, p0)
				var _2220_v57 _dafny.Set
				_ = _2220_v57
				_2220_v57 = _dafny.SetOf(Companion_Default___.Fm38(_2185_cf50, globalState))
				var _2221_v58 _dafny.Map
				_ = _2221_v58
				_2221_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2180_v34, _2180_v34)
				var _2222_v60 D5
				_ = _2222_v60
				_2222_v60 = Companion_D5_.Create_DC13_(Companion_Default___.Fm43(p0, globalState))
				var _2223_v61 _dafny.Map
				_ = _2223_v61
				_2223_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2222_v60)
				var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))
				_ = _index375
				var _rhs444 bool = ((_2220_v57).Difference(func() _dafny.Set {
					var _coll82 = _dafny.NewBuilder()
					_ = _coll82
					for _iter85 := _dafny.Iterate((_2221_v58).Keys().Elements()); ; {
						_compr_82, _ok85 := _iter85()
						if !_ok85 {
							break
						}
						var _2224_v59 _dafny.CodePoint
						_2224_v59 = interface{}(_compr_82).(_dafny.CodePoint)
						if (_2221_v58).Contains(_2224_v59) {
							_coll82.Add(_2224_v59)
						}
					}
					return _coll82.ToSet()
				}())).IsProperSubsetOf(_dafny.SetOf(_2180_v34))
				_ = _rhs444
				var _rhs445 _dafny.MultiSet = Companion_Default___.Fm57(p0, _2223_v61, globalState)
				_ = _rhs445
				var _rhs446 D11 = _2219_v56
				_ = _rhs446
				var _lhs301 _dafny.Array = _2214_v51
				_ = _lhs301
				var _lhs302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))
				_ = _lhs302
				(_lhs301).ArraySet1(_rhs444, (_lhs302).Int())
				_2218_v55 = _rhs445
				_2219_v56 = _rhs446
				var _2225_v62 _dafny.Array
				_ = _2225_v62
				var _nw344 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
				_ = _nw344
				_2225_v62 = _nw344
				var _2226_v63 _dafny.Map
				_ = _2226_v63
				_2226_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2180_v34, p0)
				var _2227_v64 _dafny.Sequence
				_ = _2227_v64
				_2227_v64 = _dafny.UnicodeSeqOfUtf8Bytes("fhuvugk")
				var _2228_v65 *C6
				_ = _2228_v65
				var _nw345 *C6 = New_C6_()
				_ = _nw345
				_nw345.Ctor__()
				_2228_v65 = _nw345
				var _2229_v66 _dafny.Map
				_ = _2229_v66
				_2229_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(788), _2228_v65)
				var _2230_v67 _dafny.Array
				_ = _2230_v67
				var _nwElement0_84 _dafny.Int = (func() _dafny.Int {
					if (_2208_v45).Contains((_2226_v63).Cardinality()) {
						return (_2208_v45).Get((_2226_v63).Cardinality()).(_dafny.Int)
					}
					return p0
				})()
				_ = _nwElement0_84
				var _nw346 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(24))
				_ = _nw346
				(_nw346).ArraySet1(_nwElement0_84, 0)
				(_nw346).ArraySet1(p0, 1)
				(_nw346).ArraySet1(p0, 2)
				(_nw346).ArraySet1(p0, 3)
				(_nw346).ArraySet1(p0, 4)
				(_nw346).ArraySet1(_dafny.IntOfUint32((_2227_v64).Cardinality()), 5)
				(_nw346).ArraySet1(p0, 6)
				(_nw346).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("le")).Cardinality()), 7)
				(_nw346).ArraySet1(p0, 8)
				(_nw346).ArraySet1(p0, 9)
				(_nw346).ArraySet1(p0, 10)
				(_nw346).ArraySet1((_2229_v66).Cardinality(), 11)
				(_nw346).ArraySet1(_dafny.IntOfUint32((_2227_v64).Cardinality()), 12)
				(_nw346).ArraySet1(p0, 13)
				(_nw346).ArraySet1(_dafny.IntOfInt64(720), 14)
				(_nw346).ArraySet1(p0, 15)
				(_nw346).ArraySet1(p0, 16)
				(_nw346).ArraySet1(p0, 17)
				(_nw346).ArraySet1(p0, 18)
				(_nw346).ArraySet1(_dafny.IntOfInt64(-383), 19)
				(_nw346).ArraySet1(p0, 20)
				(_nw346).ArraySet1(p0, 21)
				(_nw346).ArraySet1(p0, 22)
				(_nw346).ArraySet1(p0, 23)
				_2230_v67 = _nw346
				var _2231_v68 _dafny.Map
				_ = _2231_v68
				_2231_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC5_(_2230_v67), _2185_cf50)
				var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_2225_v62), 0))
				_ = _index376
				(_2225_v62).ArraySet1(_2231_v68, (_index376).Int())
				var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_2225_v62), 0))
				_ = _index377
				(_2225_v62).ArraySet1(_2231_v68, (_index377).Int())
				var _2232_v69 _dafny.Array
				_ = _2232_v69
				var _nwElement0_85 _dafny.Sequence = _2177_v31
				_ = _nwElement0_85
				var _nw347 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(25))
				_ = _nw347
				(_nw347).ArraySet1(_nwElement0_85, 0)
				(_nw347).ArraySet1(_2177_v31, 1)
				(_nw347).ArraySet1(_2177_v31, 2)
				(_nw347).ArraySet1(_2177_v31, 3)
				(_nw347).ArraySet1(_2177_v31, 4)
				(_nw347).ArraySet1(_2177_v31, 5)
				(_nw347).ArraySet1(_2177_v31, 6)
				(_nw347).ArraySet1(_dafny.SeqOf((_2214_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))).Int()).(bool)), 7)
				(_nw347).ArraySet1(_2177_v31, 8)
				(_nw347).ArraySet1(_2177_v31, 9)
				(_nw347).ArraySet1(_2177_v31, 10)
				(_nw347).ArraySet1(_dafny.Companion_Sequence_.Update(_2177_v31, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2177_v31).Cardinality()))).Uint32(), (_2214_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))).Int()).(bool)), 11)
				(_nw347).ArraySet1(_dafny.SeqOf(p3), 12)
				(_nw347).ArraySet1(Companion_Default___.Fm27(globalState), 13)
				(_nw347).ArraySet1(_2177_v31, 14)
				(_nw347).ArraySet1(_2177_v31, 15)
				(_nw347).ArraySet1(_2177_v31, 16)
				(_nw347).ArraySet1(_2177_v31, 17)
				(_nw347).ArraySet1(_2177_v31, 18)
				(_nw347).ArraySet1(_2177_v31, 19)
				(_nw347).ArraySet1(_2177_v31, 20)
				(_nw347).ArraySet1(_2177_v31, 21)
				(_nw347).ArraySet1(_dafny.SeqOf(_2185_cf50, true), 22)
				(_nw347).ArraySet1(_2177_v31, 23)
				(_nw347).ArraySet1(_2177_v31, 24)
				_2232_v69 = _nw347
				var _2233_v70 _dafny.Array
				_ = _2233_v70
				_2233_v70 = _2232_v69
				var _2234_v71 _dafny.Sequence
				_ = _2234_v71
				_2234_v71 = _dafny.SeqOf(_2230_v67, _2230_v67, _2230_v67, _2230_v67, _2230_v67)
				var _2235_v72 _dafny.Sequence
				_ = _2235_v72
				_2235_v72 = _dafny.SeqOf((_2234_v71).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-155), _dafny.IntOfUint32((_2234_v71).Cardinality()))).Uint32()).(_dafny.Array))
				var _2236_v73 D2
				_ = _2236_v73
				_2236_v73 = Companion_D2_.Create_DC5_(_2230_v67)
				var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))
				_ = _index378
				var _rhs447 _dafny.Array = _2233_v70
				_ = _rhs447
				var _rhs448 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2230_v67, (_2236_v73).Dtor_cf6()), _2234_v71)
				_ = _rhs448
				var _rhs449 bool = (_dafny.IntOfInt64(868)).Cmp(p0) < 0
				_ = _rhs449
				var _rhs450 bool = (_2228_v65).Fm16(p0, globalState)
				_ = _rhs450
				var _rhs451 bool = p1
				_ = _rhs451
				var _lhs303 _dafny.Array = _2214_v51
				_ = _lhs303
				var _lhs304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_2214_v51), 0))
				_ = _lhs304
				_2233_v70 = _rhs447
				_2235_v72 = _rhs448
				(_lhs303).ArraySet1(_rhs449, (_lhs304).Int())
				_2185_cf50 = _rhs450
				_2185_cf50 = _rhs451
			}
		} else {
			var _2237___mcc_h1 _dafny.Set = _source30.Get_().(D12_DC32).Cf49
			_ = _2237___mcc_h1
			var _2238_cf49 _dafny.Set = _2237___mcc_h1
			_ = _2238_cf49
			var _2239_v74 _dafny.Array
			_ = _2239_v74
			var _nwElement0_86 _dafny.CodePoint = _2180_v34
			_ = _nwElement0_86
			var _nw348 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(11))
			_ = _nw348
			(_nw348).ArraySet1CodePoint(_nwElement0_86, 0)
			(_nw348).ArraySet1CodePoint(Companion_Default___.Fm38(p1, globalState), 1)
			(_nw348).ArraySet1CodePoint(_2180_v34, 2)
			(_nw348).ArraySet1CodePoint(_2180_v34, 3)
			(_nw348).ArraySet1CodePoint(_dafny.CodePoint('i'), 4)
			(_nw348).ArraySet1CodePoint(_2180_v34, 5)
			(_nw348).ArraySet1CodePoint(_2180_v34, 6)
			(_nw348).ArraySet1CodePoint(_2180_v34, 7)
			(_nw348).ArraySet1CodePoint(_dafny.CodePoint('p'), 8)
			(_nw348).ArraySet1CodePoint(_2180_v34, 9)
			(_nw348).ArraySet1CodePoint(_2180_v34, 10)
			_2239_v74 = _nw348
			_2239_v74 = _2239_v74
			var _2240_v75 _dafny.Map
			_ = _2240_v75
			_2240_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _2241_v76 T0
			_ = _2241_v76
			var _nw349 *C9 = New_C9_()
			_ = _nw349
			_nw349.Ctor__()
			_2241_v76 = _nw349
			var _2242_v77 D17
			_ = _2242_v77
			_2242_v77 = Companion_D17_.Create_DC47_(_2241_v76, p3, p0, p0, p0)
			var _2243_v78 D18
			_ = _2243_v78
			_2243_v78 = Companion_D18_.Create_DC49_(p1, true, Companion_Default___.Fm38(p3, globalState), (_2242_v77).Dtor_cf70())
			var _2244_v79 _dafny.Array
			_ = _2244_v79
			var _nwElement0_87 bool = p3
			_ = _nwElement0_87
			var _nw350 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(18))
			_ = _nw350
			(_nw350).ArraySet1(_nwElement0_87, 0)
			(_nw350).ArraySet1(!(p3), 1)
			(_nw350).ArraySet1(false, 2)
			(_nw350).ArraySet1(p1, 3)
			(_nw350).ArraySet1(p3, 4)
			(_nw350).ArraySet1(p1, 5)
			(_nw350).ArraySet1(p1, 6)
			(_nw350).ArraySet1(p1, 7)
			(_nw350).ArraySet1(p3, 8)
			(_nw350).ArraySet1((func() bool {
				if (_2240_v75).Contains(_dafny.IntOfInt64(627)) {
					return (_2240_v75).Get(_dafny.IntOfInt64(627)).(bool)
				}
				return !(false)
			})(), 9)
			(_nw350).ArraySet1(p1, 10)
			(_nw350).ArraySet1(!((_2243_v78).Dtor_cf78()), 11)
			(_nw350).ArraySet1(false, 12)
			(_nw350).ArraySet1(p1, 13)
			(_nw350).ArraySet1(p3, 14)
			(_nw350).ArraySet1(p3, 15)
			(_nw350).ArraySet1(p1, 16)
			(_nw350).ArraySet1(p3, 17)
			_2244_v79 = _nw350
			var _2245_v80 _dafny.Map
			_ = _2245_v80
			_2245_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(595), _2244_v79)
			_2245_v80 = (_2245_v80).Update(p0, _2244_v79)
			var _2246_v81 *C8
			_ = _2246_v81
			var _nw351 *C8 = New_C8_()
			_ = _nw351
			_nw351.Ctor__()
			_2246_v81 = _nw351
			var _2247_v82 _dafny.Array
			_ = _2247_v82
			var _nwElement0_88 _dafny.Array = _2244_v79
			_ = _nwElement0_88
			var _nw352 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(15))
			_ = _nw352
			(_nw352).ArraySet1(_nwElement0_88, 0)
			(_nw352).ArraySet1(_2244_v79, 1)
			(_nw352).ArraySet1(_2244_v79, 2)
			(_nw352).ArraySet1(_2244_v79, 3)
			(_nw352).ArraySet1(_2244_v79, 4)
			(_nw352).ArraySet1(_2244_v79, 5)
			(_nw352).ArraySet1(_2244_v79, 6)
			(_nw352).ArraySet1(_2244_v79, 7)
			(_nw352).ArraySet1(_2244_v79, 8)
			(_nw352).ArraySet1(_2244_v79, 9)
			(_nw352).ArraySet1(_2244_v79, 10)
			(_nw352).ArraySet1(_2244_v79, 11)
			(_nw352).ArraySet1(_2244_v79, 12)
			(_nw352).ArraySet1(_2244_v79, 13)
			(_nw352).ArraySet1(_2244_v79, 14)
			_2247_v82 = _nw352
			_2247_v82 = _2247_v82
		}
		var _2248_v83 *C10
		_ = _2248_v83
		var _nw353 *C10 = New_C10_()
		_ = _nw353
		_nw353.Ctor__()
		_2248_v83 = _nw353
		var _2249_v84 bool
		_ = _2249_v84
		_2249_v84 = false
		_2249_v84 = _2249_v84
		var _2250_v85 D5
		_ = _2250_v85
		_2250_v85 = Companion_D5_.Create_DC14_(_2180_v34, p1, p0, p1)
		var _2251_v86 _dafny.Sequence
		_ = _2251_v86
		_2251_v86 = _dafny.UnicodeSeqOfUtf8Bytes("juct")
		var _2252_v87 _dafny.Array
		_ = _2252_v87
		var _nwElement0_89 _dafny.Sequence = _dafny.SeqOf((_2250_v85).Dtor_cf18(), (_2251_v86).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2251_v86).Cardinality()))).Uint32()).(_dafny.CodePoint))
		_ = _nwElement0_89
		var _nw354 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.IntOfInt64(5))
		_ = _nw354
		(_nw354).ArraySet1(_nwElement0_89, 0)
		(_nw354).ArraySet1(_2251_v86, 1)
		(_nw354).ArraySet1(Companion_Default___.Fm5(globalState), 2)
		(_nw354).ArraySet1(_2251_v86, 3)
		(_nw354).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_2180_v34, (_2251_v86).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2251_v86).Cardinality()))).Uint32()).(_dafny.CodePoint)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2177_v31).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_2180_v34, (_2251_v86).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2251_v86).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()))).Uint32(), _2180_v34), 4)
		_2252_v87 = _nw354
		var _2253_v88 D21
		_ = _2253_v88
		_2253_v88 = Companion_D21_.Create_DC56_(_2252_v87)
		var _source32 D21 = _2253_v88
		_ = _source32
		if _source32.Is_DC57() {
			var _2254___mcc_h5 _dafny.Int = _source32.Get_().(D21_DC57).Cf90
			_ = _2254___mcc_h5
			var _2255___mcc_h6 _dafny.Int = _source32.Get_().(D21_DC57).Cf91
			_ = _2255___mcc_h6
			var _2256___mcc_h7 D5 = _source32.Get_().(D21_DC57).Cf92
			_ = _2256___mcc_h7
			var _2257_cf92 D5 = _2256___mcc_h7
			_ = _2257_cf92
			var _2258_cf91 _dafny.Int = _2255___mcc_h6
			_ = _2258_cf91
			var _2259_cf90 _dafny.Int = _2254___mcc_h5
			_ = _2259_cf90
			var _2260_v89 _dafny.Map
			_ = _2260_v89
			_2260_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2249_v84, _dafny.IntOfInt64(892))
			var _2261_v90 _dafny.Map
			_ = _2261_v90
			_2261_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_2260_v89).Contains(p3) {
					return (_2260_v89).Get(p3).(_dafny.Int)
				}
				return _2259_cf90
			})(), p1)
			_2261_v90 = (Companion_Default___.Fm35(globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2259_cf90, p1))
			_2180_v34 = _dafny.CodePoint('r')
			if _2249_v84 {
				var _2262_v91 T1
				_ = _2262_v91
				var _nw355 *C8 = New_C8_()
				_ = _nw355
				_nw355.Ctor__()
				_2262_v91 = _nw355
				var _2263_v92 _dafny.MultiSet
				_ = _2263_v92
				_2263_v92 = _dafny.MultiSetOf(p3, p1)
				var _rhs452 T1 = _2262_v91
				_ = _rhs452
				var _rhs453 bool = p1
				_ = _rhs453
				var _rhs454 _dafny.MultiSet = _2263_v92
				_ = _rhs454
				_2262_v91 = _rhs452
				_2249_v84 = _rhs453
				_2263_v92 = _rhs454
				_2249_v84 = _2249_v84
				var _2264_v93 _dafny.MultiSet
				_ = _2264_v93
				_2264_v93 = _dafny.MultiSetOf((_2248_v83).Fm2((_2261_v90).Cardinality(), globalState))
				var _2265_v94 _dafny.Set
				_ = _2265_v94
				_2265_v94 = _dafny.SetOf((_2264_v93).Cardinality())
				var _2266_v95 _dafny.Array
				_ = _2266_v95
				var _nwElement0_90 bool = _2249_v84
				_ = _nwElement0_90
				var _nw356 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_90, nil, _dafny.IntOfInt64(2))
				_ = _nw356
				(_nw356).ArraySet1(_nwElement0_90, 0)
				(_nw356).ArraySet1((_2248_v83).Fm1(_2176_v30, (_2265_v94).Cardinality(), globalState), 1)
				_2266_v95 = _nw356
				var _2267_v96 _dafny.Sequence
				_ = _2267_v96
				_2267_v96 = _dafny.SeqOf(_2266_v95)
				var _2268_v97 _dafny.Map
				_ = _2268_v97
				_2268_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2266_v95, _2259_cf90)
				var _rhs455 _dafny.Int = _2259_cf90
				_ = _rhs455
				var _rhs456 _dafny.Sequence = _2267_v96
				_ = _rhs456
				var _rhs457 bool = !(_2268_v97).Contains(_2266_v95)
				_ = _rhs457
				var _rhs458 _dafny.Int = ((_dafny.Zero).Minus(p0)).Plus(_2258_cf91)
				_ = _rhs458
				var _lhs305 *GlobalState = globalState
				_ = _lhs305
				var _lhs306 *GlobalState = globalState
				_ = _lhs306
				_lhs305.F7 = _rhs455
				_2267_v96 = _rhs456
				_2249_v84 = _rhs457
				_lhs306.F15 = _rhs458
				var _2269_v98 *C9
				_ = _2269_v98
				var _nw357 *C9 = New_C9_()
				_ = _nw357
				_nw357.Ctor__()
				_2269_v98 = _nw357
				var _2270_v99 *C6
				_ = _2270_v99
				var _nw358 *C6 = New_C6_()
				_ = _nw358
				_nw358.Ctor__()
				_2270_v99 = _nw358
				var _2271_v100 _dafny.Map
				_ = _2271_v100
				_2271_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if p1 {
						return _2259_cf90
					}
					return _dafny.IntOfUint32((_2251_v86).Cardinality())
				})(), _2270_v99)
				(globalState).F13 = (_2271_v100).Cardinality()
			} else {
				var _2272_v101 _dafny.MultiSet
				_ = _2272_v101
				_2272_v101 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2179_v33).Cardinality(), false)).Cardinality())
				_2174_v28 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2259_cf90, (_2272_v101).Cardinality(), p0), _2174_v28)
				var _2273_v102 D6
				_ = _2273_v102
				_2273_v102 = Companion_D6_.Create_DC16_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(494))).Uint32(), func(coer123 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg123 _dafny.Int) interface{} {
						return coer123(arg123)
					}
				}((func(_2274_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2275_i6 _dafny.Int) _dafny.Int {
						return _2274_p0
					}
				})(p0))))
				var _2276_v103 D6
				_ = _2276_v103
				_2276_v103 = Companion_D6_.Create_DC18_(_2273_v102)
				var _2277_v104 D6
				_ = _2277_v104
				_2277_v104 = Companion_D6_.Create_DC18_(_2276_v103)
				var _2278_v105 D6
				_ = _2278_v105
				_2278_v105 = Companion_D6_.Create_DC18_(_2276_v103)
				var _2279_v106 D6
				_ = _2279_v106
				_2279_v106 = Companion_D6_.Create_DC18_(_2273_v102)
				var _2280_v107 D6
				_ = _2280_v107
				_2280_v107 = Companion_D6_.Create_DC18_(_2276_v103)
				_2280_v107 = Companion_D6_.Create_DC18_(Companion_D6_.Create_DC18_(_2276_v103))
				var _2281_v108 _dafny.Map
				_ = _2281_v108
				_2281_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2249_v84, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_2259_cf90), p1))
				var _2282_v109 _dafny.Map
				_ = _2282_v109
				_2282_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2281_v108).Cardinality())
				_2282_v109 = (_2282_v109).Update(p0, Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(255))).Uint32(), func(coer124 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg124 _dafny.Int) interface{} {
						return coer124(arg124)
					}
				}(func(_2283_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))).Cardinality())))
				_2249_v84 = (_2249_v84) && (p3)
				_2179_v33 = (_2179_v33).Union(_2179_v33)
			}
			if !(p3) || (!(_2249_v84)) {
				var _2284_v111 _dafny.Set
				_ = _2284_v111
				_2284_v111 = _dafny.SetOf(_2258_cf91)
				var _2285_v112 _dafny.MultiSet
				_ = _2285_v112
				_2285_v112 = _dafny.MultiSetOf(func() _dafny.Map {
					var _coll83 = _dafny.NewMapBuilder()
					_ = _coll83
					for _iter86 := _dafny.Iterate((_2284_v111).Elements()); ; {
						_compr_83, _ok86 := _iter86()
						if !_ok86 {
							break
						}
						var _2286_v110 _dafny.Int
						_2286_v110 = interface{}(_compr_83).(_dafny.Int)
						if (_2284_v111).Contains(_2286_v110) {
							_coll83.Add((_2286_v110).Times(_dafny.IntOfUint32((_2251_v86).Cardinality())), Companion_D21_.Create_DC57_(_2259_cf90, _2258_cf91, _2257_cf92))
						}
					}
					return _coll83.ToMap()
				}())
				var _2287_v113 _dafny.MultiSet
				_ = _2287_v113
				_2287_v113 = _dafny.MultiSetOf(p0)
				var _2288_v114 D21
				_ = _2288_v114
				_2288_v114 = Companion_D21_.Create_DC57_((_2287_v113).Cardinality(), _2258_cf91, Companion_D5_.Create_DC14_(_2180_v34, p1, p0, p1))
				var _2289_v115 _dafny.Map
				_ = _2289_v115
				_2289_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2258_cf91, _2288_v114)
				var _2290_v116 _dafny.Sequence
				_ = _2290_v116
				_2290_v116 = _dafny.SeqOf(_2289_v115)
				var _2291_v117 _dafny.Sequence
				_ = _2291_v117
				_2291_v117 = _dafny.SeqOf(_2285_v112, _dafny.MultiSetFromSeq(_2290_v116))
				var _2292_v118 _dafny.Sequence
				_ = _2292_v118
				_2292_v118 = _dafny.SeqOf(_2285_v112, (_2291_v117).Select((Companion_Default___.SafeIndex(_2258_cf91, _dafny.IntOfUint32((_2291_v117).Cardinality()))).Uint32()).(_dafny.MultiSet))
				var _2293_v119 _dafny.Sequence
				_ = _2293_v119
				_2293_v119 = _dafny.SeqOf((_2291_v117).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2291_v117).Cardinality()))).Uint32()).(_dafny.MultiSet))
				var _2294_v120 _dafny.Set
				_ = _2294_v120
				_2294_v120 = _dafny.SetOf(_2180_v34, _2180_v34, _2180_v34)
				var _2295_v121 D22
				_ = _2295_v121
				_2295_v121 = Companion_D22_.Create_DC59_(_2285_v112)
				var _pat_let_tv49 = _2285_v112
				_ = _pat_let_tv49
				var _2296_v122 _dafny.Array
				_ = _2296_v122
				var _nwElement0_91 _dafny.MultiSet = _2285_v112
				_ = _nwElement0_91
				var _nw359 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_91, nil, _dafny.IntOfInt64(29))
				_ = _nw359
				(_nw359).ArraySet1(_nwElement0_91, 0)
				(_nw359).ArraySet1(_2285_v112, 1)
				(_nw359).ArraySet1(_2285_v112, 2)
				(_nw359).ArraySet1(_2285_v112, 3)
				(_nw359).ArraySet1(_2285_v112, 4)
				(_nw359).ArraySet1(_2285_v112, 5)
				(_nw359).ArraySet1(_2285_v112, 6)
				(_nw359).ArraySet1((_2285_v112).Difference((_2292_v118).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2292_v118).Cardinality()))).Uint32()).(_dafny.MultiSet)), 7)
				(_nw359).ArraySet1(_2285_v112, 8)
				(_nw359).ArraySet1(_2285_v112, 9)
				(_nw359).ArraySet1((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_2290_v116, (Companion_Default___.SafeIndex((_2248_v83).Fm6(p1, p3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-245), p0), Companion_Default___.Fm38(p3, globalState), globalState), _dafny.IntOfUint32((_2290_v116).Cardinality()))).Uint32(), _2289_v115))).Union(_2285_v112), 10)
				(_nw359).ArraySet1((_2291_v117).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.IntOfUint32((_2291_v117).Cardinality()))).Uint32()).(_dafny.MultiSet), 11)
				(_nw359).ArraySet1(_2285_v112, 12)
				(_nw359).ArraySet1((_2285_v112).Intersection(_2285_v112), 13)
				(_nw359).ArraySet1(_2285_v112, 14)
				(_nw359).ArraySet1(_2285_v112, 15)
				(_nw359).ArraySet1(_2285_v112, 16)
				(_nw359).ArraySet1(_2285_v112, 17)
				(_nw359).ArraySet1(_2285_v112, 18)
				(_nw359).ArraySet1(_2285_v112, 19)
				(_nw359).ArraySet1((_2293_v119).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.IntOfUint32((_2293_v119).Cardinality()))).Uint32()).(_dafny.MultiSet), 20)
				(_nw359).ArraySet1(_2285_v112, 21)
				(_nw359).ArraySet1(_dafny.MultiSetOf((_2290_v116).Select((Companion_Default___.SafeIndex((_2294_v120).Cardinality(), _dafny.IntOfUint32((_2290_v116).Cardinality()))).Uint32()).(_dafny.Map), Companion_Default___.Fm58(_2249_v84, _2180_v34, globalState)), 22)
				(_nw359).ArraySet1(_2285_v112, 23)
				(_nw359).ArraySet1((func(_pat_let41_0 D22) D22 {
					return func(_2297_dt__update__tmp_h1 D22) D22 {
						return func(_pat_let42_0 _dafny.MultiSet) D22 {
							return func(_2298_dt__update_hcf94_h0 _dafny.MultiSet) D22 {
								return Companion_D22_.Create_DC59_(_2298_dt__update_hcf94_h0)
							}(_pat_let42_0)
						}(_pat_let_tv49)
					}(_pat_let41_0)
				}(_2295_v121)).Dtor_cf94(), 24)
				(_nw359).ArraySet1(_2285_v112, 25)
				(_nw359).ArraySet1(_2285_v112, 26)
				(_nw359).ArraySet1(_2285_v112, 27)
				(_nw359).ArraySet1((_2285_v112).Intersection(_2285_v112), 28)
				_2296_v122 = _nw359
				var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_2296_v122), 0))
				_ = _index379
				(_2296_v122).ArraySet1(_2285_v112, (_index379).Int())
				var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_2296_v122), 0))
				_ = _index380
				(_2296_v122).ArraySet1(_2285_v112, (_index380).Int())
				_2249_v84 = _2249_v84
				_2249_v84 = _2249_v84
				var _2299_v123 _dafny.Array
				_ = _2299_v123
				var _len0_56 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_56
				var _nw360 _dafny.Array
				_ = _nw360
				if _len0_56.Cmp(_dafny.Zero) == 0 {
					_nw360 = _dafny.NewArray(_len0_56)
				} else {
					var _init56 func(_dafny.Int) _dafny.Int = (func(_2300_cf91 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2301_i8 _dafny.Int) _dafny.Int {
							return (_2301_i8).Plus((_dafny.Zero).Minus(_2300_cf91))
						}
					})(_2258_cf91)
					_ = _init56
					var _element0_56 = _init56(_dafny.Zero)
					_ = _element0_56
					_nw360 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
					(_nw360).ArraySet1(_element0_56, 0)
					var _nativeLen0_56 = (_len0_56).Int()
					_ = _nativeLen0_56
					for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
						(_nw360).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
					}
				}
				_2299_v123 = _nw360
				var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2299_v123), 0))
				_ = _index381
				(_2299_v123).ArraySet1(p0, (_index381).Int())
				var _2302_v124 _dafny.MultiSet
				_ = _2302_v124
				_2302_v124 = _dafny.MultiSetOf(p1, !(_2249_v84), false)
				var _2303_v125 _dafny.Map
				_ = _2303_v125
				_2303_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2251_v86, _2259_cf90)
				var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2299_v123), 0))
				_ = _index382
				var _rhs459 _dafny.Int = Companion_Default___.SafeDivisionInt(_2258_cf91, (_2287_v113).Cardinality())
				_ = _rhs459
				var _rhs460 _dafny.Int = (_2302_v124).Cardinality()
				_ = _rhs460
				var _rhs461 _dafny.Int = (func() _dafny.Int {
					if (_2303_v125).Contains(_2251_v86) {
						return (_2303_v125).Get(_2251_v86).(_dafny.Int)
					}
					return Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32((_2251_v86).Cardinality()))
				})()
				_ = _rhs461
				var _lhs307 *GlobalState = globalState
				_ = _lhs307
				var _lhs308 _dafny.Array = _2299_v123
				_ = _lhs308
				var _lhs309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2299_v123), 0))
				_ = _lhs309
				var _lhs310 *GlobalState = globalState
				_ = _lhs310
				_lhs307.F7 = _rhs459
				(_lhs308).ArraySet1(_rhs460, (_lhs309).Int())
				_lhs310.F13 = _rhs461
				var _2304_v126 _dafny.Sequence
				_ = _2304_v126
				_2304_v126 = _dafny.SeqOf(_2284_v111, _2284_v111, (_2284_v111).Difference(_dafny.SetOf(_dafny.IntOfInt64(-54), _dafny.IntOfInt64(672), _2259_cf90, _2258_cf91)))
				_2304_v126 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2304_v126, _2304_v126), _2304_v126)
			} else {
				var _2305_v127 _dafny.Array
				_ = _2305_v127
				var _len0_57 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_57
				var _nw361 _dafny.Array
				_ = _nw361
				if _len0_57.Cmp(_dafny.Zero) == 0 {
					_nw361 = _dafny.NewArray(_len0_57)
				} else {
					var _init57 func(_dafny.Int) _dafny.CodePoint = (func(_2306_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2307_i9 _dafny.Int) _dafny.CodePoint {
							return _2306_v34
						}
					})(_2180_v34)
					_ = _init57
					var _element0_57 = _init57(_dafny.Zero)
					_ = _element0_57
					_nw361 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
					(_nw361).ArraySet1CodePoint(_element0_57, 0)
					var _nativeLen0_57 = (_len0_57).Int()
					_ = _nativeLen0_57
					for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
						(_nw361).ArraySet1CodePoint(_init57(_dafny.IntOf(_i0_57)), _i0_57)
					}
				}
				_2305_v127 = _nw361
				var _2308_v128 _dafny.Map
				_ = _2308_v128
				_2308_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2305_v127)
				_2308_v128 = (_2308_v128).Update(p1, _2305_v127)
				var _2309_v129 _dafny.Map
				_ = _2309_v129
				_2309_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_2259_cf90, _2259_cf90), _dafny.Companion_Sequence_.Concatenate(_2174_v28, _2174_v28))
				_2309_v129 = (_2309_v129).Update((_2259_cf90).Times(p0), _2174_v28)
				_2249_v84 = p1
				var _2310_v130 *C6
				_ = _2310_v130
				var _nw362 *C6 = New_C6_()
				_ = _nw362
				_nw362.Ctor__()
				_2310_v130 = _nw362
				_2174_v28 = _dafny.Companion_Sequence_.Concatenate(_2174_v28, _2174_v28)
			}
		} else if _source32.Is_DC56() {
			var _2311___mcc_h8 _dafny.Array = _source32.Get_().(D21_DC56).Cf89
			_ = _2311___mcc_h8
			var _2312_cf89 _dafny.Array = _2311___mcc_h8
			_ = _2312_cf89
			(globalState).F7 = p0
			_2249_v84 = p3
			var _2313_v131 *C9
			_ = _2313_v131
			var _nw363 *C9 = New_C9_()
			_ = _nw363
			_nw363.Ctor__()
			_2313_v131 = _nw363
			var _2314_v132 _dafny.Array
			_ = _2314_v132
			var _len0_58 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_58
			var _nw364 _dafny.Array
			_ = _nw364
			if _len0_58.Cmp(_dafny.Zero) == 0 {
				_nw364 = _dafny.NewArray(_len0_58)
			} else {
				var _init58 func(_dafny.Int) bool = (func(_2315_p1 bool) func(_dafny.Int) bool {
					return func(_2316_i10 _dafny.Int) bool {
						return _2315_p1
					}
				})(p1)
				_ = _init58
				var _element0_58 = _init58(_dafny.Zero)
				_ = _element0_58
				_nw364 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
				(_nw364).ArraySet1(_element0_58, 0)
				var _nativeLen0_58 = (_len0_58).Int()
				_ = _nativeLen0_58
				for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
					(_nw364).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
				}
			}
			_2314_v132 = _nw364
			var _2317_v133 D0
			_ = _2317_v133
			_2317_v133 = Companion_D0_.Create_DC0_(_2314_v132)
			var _2318_v134 _dafny.Map
			_ = _2318_v134
			_2318_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), _2314_v132)
			var _2319_v135 _dafny.Map
			_ = _2319_v135
			_2319_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2249_v84, _2314_v132)
			var _2320_v136 _dafny.Sequence
			_ = _2320_v136
			_2320_v136 = _dafny.SeqOf(_2314_v132)
			var _2321_v137 _dafny.Array
			_ = _2321_v137
			var _nwElement0_92 _dafny.Array = _2314_v132
			_ = _nwElement0_92
			var _nw365 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_92, nil, _dafny.IntOfInt64(27))
			_ = _nw365
			(_nw365).ArraySet1(_nwElement0_92, 0)
			(_nw365).ArraySet1((_2317_v133).Dtor_cf0(), 1)
			(_nw365).ArraySet1(_2314_v132, 2)
			(_nw365).ArraySet1(_2314_v132, 3)
			(_nw365).ArraySet1(_2314_v132, 4)
			(_nw365).ArraySet1(_2314_v132, 5)
			(_nw365).ArraySet1(_2314_v132, 6)
			(_nw365).ArraySet1((func() _dafny.Array {
				if (_2318_v134).Contains(p0) {
					return (_2318_v134).Get(p0).(_dafny.Array)
				}
				return _2314_v132
			})(), 7)
			(_nw365).ArraySet1((func() _dafny.Array {
				if (_2319_v135).Contains(p3) {
					return (_2319_v135).Get(p3).(_dafny.Array)
				}
				return _2314_v132
			})(), 8)
			(_nw365).ArraySet1(_2314_v132, 9)
			(_nw365).ArraySet1((func() _dafny.Array {
				if p1 {
					return _2314_v132
				}
				return _2314_v132
			})(), 10)
			(_nw365).ArraySet1(_2314_v132, 11)
			(_nw365).ArraySet1(_2314_v132, 12)
			(_nw365).ArraySet1(_2314_v132, 13)
			(_nw365).ArraySet1(_2314_v132, 14)
			(_nw365).ArraySet1(_2314_v132, 15)
			(_nw365).ArraySet1((_2320_v136).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2320_v136).Cardinality()))).Uint32()).(_dafny.Array), 16)
			(_nw365).ArraySet1(_2314_v132, 17)
			(_nw365).ArraySet1(_2314_v132, 18)
			(_nw365).ArraySet1(_2314_v132, 19)
			(_nw365).ArraySet1(_2314_v132, 20)
			(_nw365).ArraySet1(_2314_v132, 21)
			(_nw365).ArraySet1((func() _dafny.Array {
				if (_2319_v135).Contains(true) {
					return (_2319_v135).Get(true).(_dafny.Array)
				}
				return (_2320_v136).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(397))).Uint32(), func(coer125 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg125 _dafny.Int) interface{} {
						return coer125(arg125)
					}
				}(func(_2322_i11 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-388)
				}))).Cardinality()), _dafny.IntOfUint32((_2320_v136).Cardinality()))).Uint32()).(_dafny.Array)
			})(), 22)
			(_nw365).ArraySet1(_2314_v132, 23)
			(_nw365).ArraySet1(_2314_v132, 24)
			(_nw365).ArraySet1(_2314_v132, 25)
			(_nw365).ArraySet1(_2314_v132, 26)
			_2321_v137 = _nw365
			var _2323_v138 D0
			_ = _2323_v138
			_2323_v138 = Companion_D0_.Create_DC1_(_2314_v132)
			var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.ArrayLen((_2321_v137), 0))
			_ = _index383
			(_2321_v137).ArraySet1((_2323_v138).Dtor_cf1(), (_index383).Int())
			var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.ArrayLen((_2321_v137), 0))
			_ = _index384
			var _len0_59 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_59
			var _nw366 _dafny.Array
			_ = _nw366
			if _len0_59.Cmp(_dafny.Zero) == 0 {
				_nw366 = _dafny.NewArray(_len0_59)
			} else {
				var _init59 func(_dafny.Int) bool = (func(_2324_p1 bool) func(_dafny.Int) bool {
					return func(_2325_i12 _dafny.Int) bool {
						return _2324_p1
					}
				})(p1)
				_ = _init59
				var _element0_59 = _init59(_dafny.Zero)
				_ = _element0_59
				_nw366 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
				(_nw366).ArraySet1(_element0_59, 0)
				var _nativeLen0_59 = (_len0_59).Int()
				_ = _nativeLen0_59
				for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
					(_nw366).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
				}
			}
			(_2321_v137).ArraySet1(_nw366, (_index384).Int())
		} else {
			var _2326___mcc_h9 D21 = _source32.Get_().(D21_DC58).Cf93
			_ = _2326___mcc_h9
			var _2327_cf93 D21 = _2326___mcc_h9
			_ = _2327_cf93
			(globalState).F4 = p0
			var _2328_v139 _dafny.Array
			_ = _2328_v139
			var _len0_60 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_60
			var _nw367 _dafny.Array
			_ = _nw367
			if _len0_60.Cmp(_dafny.Zero) == 0 {
				_nw367 = _dafny.NewArray(_len0_60)
			} else {
				var _init60 func(_dafny.Int) _dafny.Int = (func(_2329_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2330_i13 _dafny.Int) _dafny.Int {
						return (_2330_i13).Times(_2329_p0)
					}
				})(p0)
				_ = _init60
				var _element0_60 = _init60(_dafny.Zero)
				_ = _element0_60
				_nw367 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
				(_nw367).ArraySet1(_element0_60, 0)
				var _nativeLen0_60 = (_len0_60).Int()
				_ = _nativeLen0_60
				for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
					(_nw367).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
				}
			}
			_2328_v139 = _nw367
			var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_2328_v139), 0))
			_ = _index385
			(_2328_v139).ArraySet1(p0, (_index385).Int())
			var _2331_v140 _dafny.Map
			_ = _2331_v140
			_2331_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2328_v139, (_dafny.Zero).Minus(p0))
			var _index386 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_2328_v139), 0))
			_ = _index386
			var _rhs462 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (((_2331_v140).Update(_2328_v139, p0)).Update(_2328_v139, _dafny.IntOfUint32((_2177_v31).Cardinality()))).Cardinality())
			_ = _rhs462
			var _rhs463 _dafny.Int = (p0).Minus(p0)
			_ = _rhs463
			var _lhs311 _dafny.Array = _2328_v139
			_ = _lhs311
			var _lhs312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_2328_v139), 0))
			_ = _lhs312
			var _lhs313 *GlobalState = globalState
			_ = _lhs313
			(_lhs311).ArraySet1(_rhs462, (_lhs312).Int())
			_lhs313.F7 = _rhs463
			_2180_v34 = _dafny.CodePoint('b')
			var _2332_v141 _dafny.Set
			_ = _2332_v141
			_2332_v141 = _dafny.SetOf(_2180_v34)
			var _2333_v142 D22
			_ = _2333_v142
			_2333_v142 = Companion_D22_.Create_DC61_(p0, _dafny.IntOfInt64(128), (_this).Fm2((_2328_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_2328_v139), 0))).Int()).(_dafny.Int), globalState), _dafny.Companion_Sequence_.Update(_2251_v86, (Companion_Default___.SafeIndex((_2332_v141).Cardinality(), _dafny.IntOfUint32((_2251_v86).Cardinality()))).Uint32(), _2180_v34))
			var _2334_v143 _dafny.MultiSet
			_ = _2334_v143
			_2334_v143 = _dafny.MultiSetOf(_2333_v142)
			(globalState).F13 = (_dafny.Zero).Minus(((_2328_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_2328_v139), 0))).Int()).(_dafny.Int)).Minus((_2334_v143).Cardinality()))
		}
		(globalState).F7 = p0
		var _2335_v144 _dafny.MultiSet
		_ = _2335_v144
		_2335_v144 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0))
		r0 = _2335_v144
		return r0
	}
}

// End of class C11

// Definition of class C12
type C12 struct {
	_f22 _dafny.Int
	_f23 _dafny.Sequence
}

func New_C12_() *C12 {
	_this := C12{}

	_this._f22 = _dafny.Zero
	_this._f23 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C12_ struct {
}

var Companion_C12_ = CompanionStruct_C12_{}

func (_this *C12) Equals(other *C12) bool {
	return _this == other
}

func (_this *C12) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C12)
	return ok && _this.Equals(other)
}

func (*C12) String() string {
	return "_module.C12"
}

func Type_C12_() _dafny.TypeDescriptor {
	return type_C12_{}
}

type type_C12_ struct {
}

func (_this type_C12_) Default() interface{} {
	return (*C12)(nil)
}

func (_this type_C12_) String() string {
	return "main.C12"
}
func (_this *C12) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C12{}
var _ _dafny.TraitOffspring = &C12{}

func (_this *C12) Ctor__(f22 _dafny.Int, f23 _dafny.Sequence) {
	{
		(_this)._f22 = f22
		(_this)._f23 = f23
	}
}
func (_this *C12) Fm1(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_this).F22()).Cmp((_this).F22()) >= 0
	}
}
func (_this *C12) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F22()).Times((_dafny.MultiSetOf(true)).Cardinality())
	}
}
func (_this *C12) Fm3(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf((_this).F22())).Cardinality()), Companion_Default___.SafeModuloInt((_this).F22(), _dafny.IntOfInt64(561)))
	}
}
func (_this *C12) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, p3 D1, globalState *GlobalState) _dafny.Set {
	{
		return ((_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true))).Cardinality())).Union(_dafny.SetOf((_this).F22(), (_dafny.MultiSetOf(true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F22())).Cardinality()))).Intersection((_dafny.SetOf((_dafny.MultiSetOf((_this).F22(), (_this).F22())).Cardinality(), (_this).F22())).Intersection(func() _dafny.Set {
			var _coll84 = _dafny.NewBuilder()
			_ = _coll84
			for _iter87 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), false)).Keys().Elements()); ; {
				_compr_84, _ok87 := _iter87()
				if !_ok87 {
					break
				}
				var _2336_v0 _dafny.Int
				_2336_v0 = interface{}(_compr_84).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), false)).Contains(_2336_v0) {
					_coll84.Add((_2336_v0).Times(_dafny.IntOfInt64(-991)))
				}
			}
			return _coll84.ToSet()
		}()))
	}
}
func (_this *C12) M1(p0 _dafny.Sequence, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _2337_v0 bool
		_ = _2337_v0
		_2337_v0 = true
		(globalState).F4 = ((func() _dafny.Int {
			if _2337_v0 {
				return (_this).F22()
			}
			return (_this).F22()
		})()).Times((_this).F22())
		var _2338_i0 _dafny.Int
		_ = _2338_i0
		_2338_i0 = _dafny.Zero
		{
			for _2337_v0 {
				{
					if (_2338_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L15
					}
					_2338_i0 = (_2338_i0).Plus(_dafny.One)
					var _2339_v1 _dafny.CodePoint
					_ = _2339_v1
					_2339_v1 = _dafny.CodePoint('r')
					_2339_v1 = (p0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F22()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint)
					var _2340_v2 _dafny.Array
					_ = _2340_v2
					var _nw368 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
					_ = _nw368
					_2340_v2 = _nw368
					_2340_v2 = _2340_v2
					var _2341_v3 _dafny.Set
					_ = _2341_v3
					_2341_v3 = _dafny.SetOf((_this).F22(), (_this).Fm2((_this).F22(), globalState))
					var _2342_v4 _dafny.Map
					_ = _2342_v4
					_2342_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2341_v3).IsProperSubsetOf(_2341_v3), _2337_v0)
					if (func() bool {
						if (_2342_v4).Contains(_2337_v0) {
							return (_2342_v4).Get(_2337_v0).(bool)
						}
						return false
					})() {
						var _2343_v5 _dafny.Map
						_ = _2343_v5
						_2343_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2337_v0, _2342_v4)
						var _2344_v6 _dafny.Sequence
						_ = _2344_v6
						_2344_v6 = _dafny.SeqOf(_2337_v0, _2337_v0, _2337_v0)
						var _2345_v7 _dafny.Array
						_ = _2345_v7
						var _nwElement0_93 bool = (_this).Fm1(_2343_v5, (_this).F22(), globalState)
						_ = _nwElement0_93
						var _nw369 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_93, nil, _dafny.IntOfInt64(6))
						_ = _nw369
						(_nw369).ArraySet1(_nwElement0_93, 0)
						(_nw369).ArraySet1((_2344_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.IntOfUint32((_2344_v6).Cardinality()))).Uint32()).(bool), 1)
						(_nw369).ArraySet1((_dafny.IntOfInt64(873)).Cmp((_this).F22()) <= 0, 2)
						(_nw369).ArraySet1(true, 3)
						(_nw369).ArraySet1(_dafny.Companion_Sequence_.Equal(p0, p0), 4)
						(_nw369).ArraySet1(_2337_v0, 5)
						_2345_v7 = _nw369
						var _index387 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_2345_v7), 0))
						_ = _index387
						(_2345_v7).ArraySet1(_2337_v0, (_index387).Int())
						var _index388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_2345_v7), 0))
						_ = _index388
						(_2345_v7).ArraySet1(_2337_v0, (_index388).Int())
						(globalState).F7 = (_this).F22()
						(globalState).F13 = (_this).Fm2((_this).F22(), globalState)
						var _index389 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_2345_v7), 0))
						_ = _index389
						(_2345_v7).ArraySet1(!_dafny.Companion_Sequence_.Contains((_this).F23(), _2339_v1), (_index389).Int())
						var _nw370 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
						_ = _nw370
						_2345_v7 = _nw370
					} else {
						var _2346_v8 _dafny.Sequence
						_ = _2346_v8
						_2346_v8 = _dafny.UnicodeSeqOfUtf8Bytes("hlrwshec")
						_2346_v8 = (func() _dafny.Sequence {
							if ((_this).F22()).Cmp(_dafny.IntOfInt64(-642)) < 0 {
								return (_this).F23()
							}
							return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("umwm"), _dafny.UnicodeSeqOfUtf8Bytes("npllbkx")), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F22()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("umwm"), _dafny.UnicodeSeqOfUtf8Bytes("npllbkx"))).Cardinality()))).Uint32(), _2339_v1)
						})()
						var _2347_v9 _dafny.Array
						_ = _2347_v9
						var _nw371 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
						_ = _nw371
						_2347_v9 = _nw371
						var _2348_v10 _dafny.Array
						_ = _2348_v10
						var _len0_61 _dafny.Int = _dafny.IntOfInt64(20)
						_ = _len0_61
						var _nw372 _dafny.Array
						_ = _nw372
						if _len0_61.Cmp(_dafny.Zero) == 0 {
							_nw372 = _dafny.NewArray(_len0_61)
						} else {
							var _init61 func(_dafny.Int) _dafny.Map = (func(_2349_v0 bool) func(_dafny.Int) _dafny.Map {
								return func(_2350_i1 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2349_v0, _dafny.IntOfInt64(-494))
								}
							})(_2337_v0)
							_ = _init61
							var _element0_61 = _init61(_dafny.Zero)
							_ = _element0_61
							_nw372 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
							(_nw372).ArraySet1(_element0_61, 0)
							var _nativeLen0_61 = (_len0_61).Int()
							_ = _nativeLen0_61
							for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
								(_nw372).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
							}
						}
						_2348_v10 = _nw372
						var _index390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_2347_v9), 0))
						_ = _index390
						(_2347_v9).ArraySet1(_2348_v10, (_index390).Int())
						var _2351_v11 _dafny.Array
						_ = _2351_v11
						var _len0_62 _dafny.Int = _dafny.IntOfInt64(6)
						_ = _len0_62
						var _nw373 _dafny.Array
						_ = _nw373
						if _len0_62.Cmp(_dafny.Zero) == 0 {
							_nw373 = _dafny.NewArray(_len0_62)
						} else {
							var _init62 func(_dafny.Int) bool = (func(_2352_v0 bool) func(_dafny.Int) bool {
								return func(_2353_i2 _dafny.Int) bool {
									return _2352_v0
								}
							})(_2337_v0)
							_ = _init62
							var _element0_62 = _init62(_dafny.Zero)
							_ = _element0_62
							_nw373 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
							(_nw373).ArraySet1(_element0_62, 0)
							var _nativeLen0_62 = (_len0_62).Int()
							_ = _nativeLen0_62
							for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
								(_nw373).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
							}
						}
						_2351_v11 = _nw373
						var _2354_v12 _dafny.Sequence
						_ = _2354_v12
						_2354_v12 = _dafny.SeqOf(_2351_v11)
						var _index391 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_2347_v9), 0))
						_ = _index391
						var _rhs464 _dafny.Array = _2348_v10
						_ = _rhs464
						var _rhs465 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("oexfuyyh")
						_ = _rhs465
						var _rhs466 _dafny.Int = _dafny.IntOfUint32((_2354_v12).Cardinality())
						_ = _rhs466
						var _lhs314 _dafny.Array = _2347_v9
						_ = _lhs314
						var _lhs315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_2347_v9), 0))
						_ = _lhs315
						var _lhs316 *GlobalState = globalState
						_ = _lhs316
						(_lhs314).ArraySet1(_rhs464, (_lhs315).Int())
						_2346_v8 = _rhs465
						_lhs316.F7 = _rhs466
						_2346_v8 = p0
						var _2355_v13 *C6
						_ = _2355_v13
						var _nw374 *C6 = New_C6_()
						_ = _nw374
						_nw374.Ctor__()
						_2355_v13 = _nw374
						(globalState).F7 = ((_this).F22()).Times((_this).F22())
					}
					_2337_v0 = !(_2337_v0)
					goto C15
				}
			C15:
			}
			goto L15
		}
	L15:
		var _2356_v14 _dafny.Map
		_ = _2356_v14
		_2356_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _2337_v0)
		var _2357_v15 D14
		_ = _2357_v15
		_2357_v15 = Companion_D14_.Create_DC37_(Companion_D14_.Create_DC35_(_dafny.SeqOf(_2356_v14)))
		var _2358_v16 _dafny.Sequence
		_ = _2358_v16
		_2358_v16 = _dafny.SeqOf(_2337_v0, _2337_v0, _2337_v0)
		_2357_v15 = Companion_D14_.Create_DC37_(Companion_D14_.Create_DC36_(_2358_v16))
		var _hi15 _dafny.Int = (_this).F22()
		_ = _hi15
		for _2359_i3 := (_this).F22(); _2359_i3.Cmp(_hi15) < 0; _2359_i3 = _2359_i3.Plus(_dafny.One) {
			if ((_this).F22()).Cmp(_2359_i3) > 0 {
				var _2360_v17 _dafny.Array
				_ = _2360_v17
				var _len0_63 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_63
				var _nw375 _dafny.Array
				_ = _nw375
				if _len0_63.Cmp(_dafny.Zero) == 0 {
					_nw375 = _dafny.NewArray(_len0_63)
				} else {
					var _init63 func(_dafny.Int) _dafny.Int = (func(_2361_i3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2362_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_2362_i4, _2361_i3)
						}
					})(_2359_i3)
					_ = _init63
					var _element0_63 = _init63(_dafny.Zero)
					_ = _element0_63
					_nw375 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
					(_nw375).ArraySet1(_element0_63, 0)
					var _nativeLen0_63 = (_len0_63).Int()
					_ = _nativeLen0_63
					for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
						(_nw375).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
					}
				}
				_2360_v17 = _nw375
				_2360_v17 = _2360_v17
				_2337_v0 = (_2337_v0) == (_2337_v0)
				var _2363_v18 _dafny.Array
				_ = _2363_v18
				var _nw376 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
				_ = _nw376
				_2363_v18 = _nw376
				var _2364_v19 _dafny.CodePoint
				_ = _2364_v19
				_2364_v19 = _dafny.CodePoint('a')
				var _index392 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_2363_v18), 0))
				_ = _index392
				(_2363_v18).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("spf"), (Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("spf")).Cardinality()))).Uint32(), _2364_v19), (_index392).Int())
				var _2365_v20 _dafny.Sequence
				_ = _2365_v20
				_2365_v20 = _dafny.UnicodeSeqOfUtf8Bytes("s")
				var _index393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_2363_v18), 0))
				_ = _index393
				var _rhs467 _dafny.Array = _2360_v17
				_ = _rhs467
				var _rhs468 _dafny.Sequence = (_this).F23()
				_ = _rhs468
				var _rhs469 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p0, p0), (Companion_Default___.SafeIndex(((_this).F22()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(847))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg126 _dafny.Int) interface{} {
						return coer126(arg126)
					}
				}((func(_2366_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2367_i5 _dafny.Int) _dafny.CodePoint {
						return _2366_v19
					}
				})(_2364_v19)))).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()))).Uint32(), _2364_v19), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_2359_i3, _dafny.IntOfInt64(534)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p0, p0), (Companion_Default___.SafeIndex(((_this).F22()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(847))).Uint32(), func(coer127 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg127 _dafny.Int) interface{} {
						return coer127(arg127)
					}
				}((func(_2368_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2369_i5 _dafny.Int) _dafny.CodePoint {
						return _2368_v19
					}
				})(_2364_v19)))).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()))).Uint32(), _2364_v19)).Cardinality()))).Uint32(), _2364_v19), (Companion_Default___.SafeIndex((func() _dafny.Int {
					if _2337_v0 {
						return (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_2358_v16).Cardinality())))
					}
					return _2359_i3
				})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p0, p0), (Companion_Default___.SafeIndex(((_this).F22()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(847))).Uint32(), func(coer128 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg128 _dafny.Int) interface{} {
						return coer128(arg128)
					}
				}((func(_2370_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2371_i5 _dafny.Int) _dafny.CodePoint {
						return _2370_v19
					}
				})(_2364_v19)))).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()))).Uint32(), _2364_v19), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_2359_i3, _dafny.IntOfInt64(534)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(p0, p0), (Companion_Default___.SafeIndex(((_this).F22()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(847))).Uint32(), func(coer129 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg129 _dafny.Int) interface{} {
						return coer129(arg129)
					}
				}((func(_2372_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2373_i5 _dafny.Int) _dafny.CodePoint {
						return _2372_v19
					}
				})(_2364_v19)))).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()))).Uint32(), _2364_v19)).Cardinality()))).Uint32(), _2364_v19)).Cardinality()))).Uint32(), _2364_v19)
				_ = _rhs469
				var _lhs317 _dafny.Array = _2363_v18
				_ = _lhs317
				var _lhs318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_2363_v18), 0))
				_ = _lhs318
				_2360_v17 = _rhs467
				(_lhs317).ArraySet1(_rhs468, (_lhs318).Int())
				_2365_v20 = _rhs469
				var _2374_v21 D8
				_ = _2374_v21
				_2374_v21 = Companion_D8_.Create_DC25_((_this).F22())
				_2374_v21 = func(_pat_let43_0 D8) D8 {
					return func(_2377_dt__update__tmp_h1 D8) D8 {
						return func(_pat_let46_0 _dafny.Int) D8 {
							return func(_2378_dt__update_hcf42_h1 _dafny.Int) D8 {
								return Companion_D8_.Create_DC25_(_2378_dt__update_hcf42_h1)
							}(_pat_let46_0)
						}((_2359_i3).Minus((_this).F22()))
					}(_pat_let43_0)
				}(func(_pat_let44_0 D8) D8 {
					return func(_2375_dt__update__tmp_h0 D8) D8 {
						return func(_pat_let45_0 _dafny.Int) D8 {
							return func(_2376_dt__update_hcf42_h0 _dafny.Int) D8 {
								return Companion_D8_.Create_DC25_(_2376_dt__update_hcf42_h0)
							}(_pat_let45_0)
						}(_2359_i3)
					}(_pat_let44_0)
				}(_2374_v21))
				_2364_v19 = _2364_v19
			} else {
				var _2379_v22 _dafny.MultiSet
				_ = _2379_v22
				_2379_v22 = _dafny.MultiSetOf(_2337_v0)
				(globalState).F13 = (_2379_v22).Cardinality()
				var _2380_v23 _dafny.Array
				_ = _2380_v23
				var _nw377 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw377
				_2380_v23 = _nw377
				var _index394 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2380_v23), 0))
				_ = _index394
				(_2380_v23).ArraySet1(_2359_i3, (_index394).Int())
				var _2381_v24 _dafny.Map
				_ = _2381_v24
				_2381_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2337_v0, _2337_v0)
				var _2382_v25 D15
				_ = _2382_v25
				_2382_v25 = Companion_D15_.Create_DC39_(_2381_v24, _2337_v0, (_2379_v22).Cardinality())
				var _pat_let_tv50 = _2337_v0
				_ = _pat_let_tv50
				var _index395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2380_v23), 0))
				_ = _index395
				(_2380_v23).ArraySet1((func() _dafny.Int {
					if _2337_v0 {
						return (func(_pat_let47_0 D15) D15 {
							return func(_2383_dt__update__tmp_h2 D15) D15 {
								return func(_pat_let48_0 bool) D15 {
									return func(_2384_dt__update_hcf57_h0 bool) D15 {
										return Companion_D15_.Create_DC39_((_2383_dt__update__tmp_h2).Dtor_cf56(), _2384_dt__update_hcf57_h0, (_2383_dt__update__tmp_h2).Dtor_cf58())
									}(_pat_let48_0)
								}(!(_pat_let_tv50))
							}(_pat_let47_0)
						}(_2382_v25)).Dtor_cf58()
					}
					return (_dafny.Zero).Minus((_dafny.Zero).Minus(((_this).F22()).Times(_2359_i3)))
				})(), (_index395).Int())
				var _2385_v26 *C2
				_ = _2385_v26
				var _nw378 *C2 = New_C2_()
				_ = _nw378
				_nw378.Ctor__()
				_2385_v26 = _nw378
				_2337_v0 = !(_2337_v0) || ((_2385_v26).Fm8(_dafny.IntOfInt64(24), _2337_v0, _2359_i3, globalState))
				var _2386_v27 _dafny.Array
				_ = _2386_v27
				var _nw379 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw379
				_2386_v27 = _nw379
				var _2387_v28 _dafny.Array
				_ = _2387_v28
				var _nwElement0_94 bool = _2337_v0
				_ = _nwElement0_94
				var _nw380 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_94, nil, _dafny.IntOfInt64(18))
				_ = _nw380
				(_nw380).ArraySet1(_nwElement0_94, 0)
				(_nw380).ArraySet1(_2337_v0, 1)
				(_nw380).ArraySet1(_2337_v0, 2)
				(_nw380).ArraySet1(_2337_v0, 3)
				(_nw380).ArraySet1(!(_2337_v0), 4)
				(_nw380).ArraySet1(_2337_v0, 5)
				(_nw380).ArraySet1(_2337_v0, 6)
				(_nw380).ArraySet1(_2337_v0, 7)
				(_nw380).ArraySet1(_2337_v0, 8)
				(_nw380).ArraySet1(_2337_v0, 9)
				(_nw380).ArraySet1(_2337_v0, 10)
				(_nw380).ArraySet1(Companion_Default___.Fm26(globalState), 11)
				(_nw380).ArraySet1(_2337_v0, 12)
				(_nw380).ArraySet1(_2337_v0, 13)
				(_nw380).ArraySet1((_2385_v26).Fm8((_2380_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2380_v23), 0))).Int()).(_dafny.Int), _2337_v0, Companion_Default___.Fm17(_dafny.IntOfInt64(57), globalState), globalState), 14)
				(_nw380).ArraySet1(_2337_v0, 15)
				(_nw380).ArraySet1(_2337_v0, 16)
				(_nw380).ArraySet1((func() bool {
					if (_2356_v14).Contains(_2359_i3) {
						return (_2356_v14).Get(_2359_i3).(bool)
					}
					return _2337_v0
				})(), 17)
				_2387_v28 = _nw380
				var _index396 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_2386_v27), 0))
				_ = _index396
				(_2386_v27).ArraySet1(_2387_v28, (_index396).Int())
				var _2388_v29 _dafny.CodePoint
				_ = _2388_v29
				_2388_v29 = _dafny.CodePoint('a')
				var _2389_v30 _dafny.MultiSet
				_ = _2389_v30
				_2389_v30 = _dafny.MultiSetOf(_2388_v29, _2388_v29)
				var _2390_v31 D7
				_ = _2390_v31
				_2390_v31 = Companion_D7_.Create_DC22_(_2337_v0, (_2380_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2380_v23), 0))).Int()).(_dafny.Int), _2337_v0, (_2380_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2380_v23), 0))).Int()).(_dafny.Int))
				var _index397 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_2386_v27), 0))
				_ = _index397
				var _nwElement0_95 bool = (_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2381_v24), (_dafny.Zero).Minus(_2359_i3), globalState)
				_ = _nwElement0_95
				var _nw381 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_95, nil, _dafny.IntOfInt64(22))
				_ = _nw381
				(_nw381).ArraySet1(_nwElement0_95, 0)
				(_nw381).ArraySet1((_2337_v0) || (_2337_v0), 1)
				(_nw381).ArraySet1(_2337_v0, 2)
				(_nw381).ArraySet1(_2337_v0, 3)
				(_nw381).ArraySet1((func() bool {
					if _2337_v0 {
						return _2337_v0
					}
					return _2337_v0
				})(), 4)
				(_nw381).ArraySet1(_2337_v0, 5)
				(_nw381).ArraySet1((_2389_v30).IsSubsetOf(_2389_v30), 6)
				(_nw381).ArraySet1(true, 7)
				(_nw381).ArraySet1(!(true) || (_2337_v0), 8)
				(_nw381).ArraySet1((func() bool {
					if _2337_v0 {
						return _2337_v0
					}
					return _2337_v0
				})(), 9)
				(_nw381).ArraySet1(_2337_v0, 10)
				(_nw381).ArraySet1(_2337_v0, 11)
				(_nw381).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqOf((_2380_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_2380_v23), 0))).Int()).(_dafny.Int)))).Contains((_dafny.Zero).Minus(_2359_i3)), 12)
				(_nw381).ArraySet1(_2337_v0, 13)
				(_nw381).ArraySet1(!(_2337_v0), 14)
				(_nw381).ArraySet1((func() bool {
					if (_2356_v14).Contains((Companion_Default___.Fm35(globalState)).Cardinality()) {
						return (_2356_v14).Get((Companion_Default___.Fm35(globalState)).Cardinality()).(bool)
					}
					return (_2390_v31).Dtor_cf37()
				})(), 15)
				(_nw381).ArraySet1(false, 16)
				(_nw381).ArraySet1(_2337_v0, 17)
				(_nw381).ArraySet1(_2337_v0, 18)
				(_nw381).ArraySet1(true, 19)
				(_nw381).ArraySet1(false, 20)
				(_nw381).ArraySet1(!(_2337_v0), 21)
				(_2386_v27).ArraySet1(_nw381, (_index397).Int())
			}
			var _2391_v32 _dafny.MultiSet
			_ = _2391_v32
			_2391_v32 = _dafny.MultiSetOf(false, _2337_v0)
			var _2392_v33 _dafny.Array
			_ = _2392_v33
			var _nwElement0_96 bool = _2337_v0
			_ = _nwElement0_96
			var _nw382 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_96, nil, _dafny.IntOfInt64(2))
			_ = _nw382
			(_nw382).ArraySet1(_nwElement0_96, 0)
			(_nw382).ArraySet1((_dafny.MultiSetOf(_2337_v0)).IsProperSubsetOf(_2391_v32), 1)
			_2392_v33 = _nw382
			var _index398 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))
			_ = _index398
			(_2392_v33).ArraySet1(!(_2337_v0) || (_2337_v0), (_index398).Int())
			var _index399 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))
			_ = _index399
			(_2392_v33).ArraySet1(_2337_v0, (_index399).Int())
			var _2393_v34 _dafny.Array
			_ = _2393_v34
			var _nw383 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw383
			_2393_v34 = _nw383
			var _2394_v35 _dafny.Array
			_ = _2394_v35
			var _nwElement0_97 _dafny.Int = _dafny.IntOfInt64(684)
			_ = _nwElement0_97
			var _nw384 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_97, nil, _dafny.IntOfInt64(19))
			_ = _nw384
			(_nw384).ArraySet1(_nwElement0_97, 0)
			(_nw384).ArraySet1((_this).F22(), 1)
			(_nw384).ArraySet1((_this).F22(), 2)
			(_nw384).ArraySet1((_this).F22(), 3)
			(_nw384).ArraySet1(_dafny.IntOfInt64(871), 4)
			(_nw384).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm27(globalState)).Cardinality()), 5)
			(_nw384).ArraySet1(_2359_i3, 6)
			(_nw384).ArraySet1((_this).F22(), 7)
			(_nw384).ArraySet1(_2359_i3, 8)
			(_nw384).ArraySet1((_this).F22(), 9)
			(_nw384).ArraySet1((_this).F22(), 10)
			(_nw384).ArraySet1(_dafny.IntOfUint32((_2358_v16).Cardinality()), 11)
			(_nw384).ArraySet1((_this).F22(), 12)
			(_nw384).ArraySet1(_2359_i3, 13)
			(_nw384).ArraySet1(_2359_i3, 14)
			(_nw384).ArraySet1((_2391_v32).Cardinality(), 15)
			(_nw384).ArraySet1((_this).F22(), 16)
			(_nw384).ArraySet1((_this).F22(), 17)
			(_nw384).ArraySet1((_this).F22(), 18)
			_2394_v35 = _nw384
			var _2395_v36 D2
			_ = _2395_v36
			_2395_v36 = Companion_D2_.Create_DC5_(_2394_v35)
			var _2396_v38 _dafny.Map
			_ = _2396_v38
			_2396_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), Companion_Default___.Fm30(Companion_Default___.Fm17(_2359_i3, globalState), (_2392_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))).Int()).(bool), func() _dafny.Set {
				var _coll85 = _dafny.NewBuilder()
				_ = _coll85
				for _iter88 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(303), _dafny.IntOfInt64(-84))); ; {
					_compr_85, _ok88 := _iter88()
					if !_ok88 {
						break
					}
					var _2397_v37 _dafny.Int
					_2397_v37 = interface{}(_compr_85).(_dafny.Int)
					if ((_dafny.IntOfInt64(303)).Cmp(_2397_v37) <= 0) && ((_2397_v37).Cmp(_dafny.IntOfInt64(-84)) < 0) {
						_coll85.Add((_2397_v37).Times(_2359_i3))
					}
				}
				return _coll85.ToSet()
			}(), _2337_v0, globalState))
			var _2398_v39 _dafny.Sequence
			_ = _2398_v39
			_2398_v39 = _dafny.SeqOf(_2396_v38, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _dafny.UnicodeSeqOfUtf8Bytes("gpjnbt")), _2396_v38)
			var _2399_v40 _dafny.Map
			_ = _2399_v40
			_2399_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_2398_v39).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F22()), _dafny.IntOfUint32((_2398_v39).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _2394_v35)
			var _2400_v41 _dafny.Array
			_ = _2400_v41
			var _nwElement0_98 _dafny.Array = (_2395_v36).Dtor_cf6()
			_ = _nwElement0_98
			var _nw385 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_98, nil, _dafny.IntOfInt64(12))
			_ = _nw385
			(_nw385).ArraySet1(_nwElement0_98, 0)
			(_nw385).ArraySet1(_2394_v35, 1)
			(_nw385).ArraySet1(_2394_v35, 2)
			(_nw385).ArraySet1(_2394_v35, 3)
			(_nw385).ArraySet1(_2394_v35, 4)
			(_nw385).ArraySet1(_2394_v35, 5)
			(_nw385).ArraySet1(_2394_v35, 6)
			(_nw385).ArraySet1(_2394_v35, 7)
			(_nw385).ArraySet1(_2394_v35, 8)
			(_nw385).ArraySet1(_2394_v35, 9)
			(_nw385).ArraySet1(_2394_v35, 10)
			(_nw385).ArraySet1((func() _dafny.Array {
				if (_2399_v40).Contains(_2359_i3) {
					return (_2399_v40).Get(_2359_i3).(_dafny.Array)
				}
				return _2394_v35
			})(), 11)
			_2400_v41 = _nw385
			var _index400 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_2393_v34), 0))
			_ = _index400
			(_2393_v34).ArraySet1(_2400_v41, (_index400).Int())
			var _index401 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_2393_v34), 0))
			_ = _index401
			(_2393_v34).ArraySet1(_2400_v41, (_index401).Int())
			if _dafny.Companion_Sequence_.IsPrefixOf(p0, p0) {
				(globalState).F15 = (_this).F22()
				var _2401_v42 _dafny.Sequence
				_ = _2401_v42
				_2401_v42 = _dafny.UnicodeSeqOfUtf8Bytes("ks")
				_2401_v42 = (func() _dafny.Sequence {
					if (_2392_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))).Int()).(bool) {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-693))).Uint32(), func(coer130 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg130 _dafny.Int) interface{} {
								return coer130(arg130)
							}
						}(func(_2402_i6 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('k')
						}))
					}
					return _dafny.Companion_Sequence_.Concatenate(_2401_v42, p0)
				})()
				var _2403_v43 _dafny.Array
				_ = _2403_v43
				var _nw386 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw386
				_2403_v43 = _nw386
				var _index402 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_2403_v43), 0))
				_ = _index402
				(_2403_v43).ArraySet1(_2392_v33, (_index402).Int())
				var _2404_v44 *C2
				_ = _2404_v44
				var _nw387 *C2 = New_C2_()
				_ = _nw387
				_nw387.Ctor__()
				_2404_v44 = _nw387
				var _2405_v45 _dafny.Map
				_ = _2405_v45
				_2405_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2404_v44, _2392_v33)
				var _index403 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_2403_v43), 0))
				_ = _index403
				(_2403_v43).ArraySet1((func() _dafny.Array {
					if (_2405_v45).Contains(_2404_v44) {
						return (_2405_v45).Get(_2404_v44).(_dafny.Array)
					}
					return _2392_v33
				})(), (_index403).Int())
				var _2406_v46 _dafny.Set
				_ = _2406_v46
				_2406_v46 = _dafny.SetOf((_2392_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))).Int()).(bool))
				_2356_v14 = (_2356_v14).Update(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F22()), _2359_i3), (_2406_v46).IsDisjointFrom(_2406_v46))
				_2406_v46 = (_2406_v46).Union(_2406_v46)
			} else {
				var _2407_v47 T0
				_ = _2407_v47
				var _nw388 *C11 = New_C11_()
				_ = _nw388
				_nw388.Ctor__()
				_2407_v47 = _nw388
				var _2408_v48 _dafny.Map
				_ = _2408_v48
				_2408_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2407_v47, (_2392_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))).Int()).(bool))
				var _2409_v49 _dafny.Sequence
				_ = _2409_v49
				_2409_v49 = _dafny.SeqOf(_2408_v48, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2407_v47, (_2392_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))).Int()).(bool)))
				_2337_v0 = !(!_dafny.Companion_Sequence_.Equal(_2409_v49, _2409_v49))
				var _index404 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))
				_ = _index404
				(_2392_v33).ArraySet1(!(_2337_v0), (_index404).Int())
				var _2410_v50 *C2
				_ = _2410_v50
				var _nw389 *C2 = New_C2_()
				_ = _nw389
				_nw389.Ctor__()
				_2410_v50 = _nw389
				var _2411_v51 _dafny.Sequence
				_ = _2411_v51
				_2411_v51 = _dafny.UnicodeSeqOfUtf8Bytes("jhounbr")
				var _2412_v52 D0
				_ = _2412_v52
				_2412_v52 = Companion_D0_.Create_DC1_(_2392_v33)
				var _index405 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))
				_ = _index405
				var _rhs470 _dafny.Int = (_dafny.Zero).Minus(_2359_i3)
				_ = _rhs470
				var _rhs471 bool = _2337_v0
				_ = _rhs471
				var _rhs472 _dafny.Sequence = (func() _dafny.Sequence {
					if (_2396_v38).Contains(_dafny.IntOfUint32(((_this).F23()).Cardinality())) {
						return (_2396_v38).Get(_dafny.IntOfUint32(((_this).F23()).Cardinality())).(_dafny.Sequence)
					}
					return (_this).F23()
				})()
				_ = _rhs472
				var _rhs473 D0 = _2412_v52
				_ = _rhs473
				var _lhs319 *GlobalState = globalState
				_ = _lhs319
				var _lhs320 _dafny.Array = _2392_v33
				_ = _lhs320
				var _lhs321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_2392_v33), 0))
				_ = _lhs321
				_lhs319.F15 = _rhs470
				(_lhs320).ArraySet1(_rhs471, (_lhs321).Int())
				_2411_v51 = _rhs472
				_2412_v52 = _rhs473
				var _2413_v53 _dafny.Sequence
				_ = _2413_v53
				_2413_v53 = _dafny.SeqOf(Companion_Default___.SafeModuloInt((_this).F22(), _2359_i3))
				var _2414_v54 _dafny.Sequence
				_ = _2414_v54
				_2414_v54 = _dafny.SeqOf(_2411_v51)
				var _rhs474 _dafny.Int = (_2356_v14).Cardinality()
				_ = _rhs474
				var _rhs475 _dafny.Sequence = _2413_v53
				_ = _rhs475
				var _rhs476 _dafny.Sequence = p0
				_ = _rhs476
				var _rhs477 _dafny.Sequence = _2414_v54
				_ = _rhs477
				var _lhs322 *GlobalState = globalState
				_ = _lhs322
				_lhs322.F4 = _rhs474
				_2413_v53 = _rhs475
				_2411_v51 = _rhs476
				_2414_v54 = _rhs477
			}
		}
		var _2415_v55 *C0
		_ = _2415_v55
		var _nw390 *C0 = New_C0_()
		_ = _nw390
		_nw390.Ctor__()
		_2415_v55 = _nw390
		var _2416_v56 _dafny.Array
		_ = _2416_v56
		var _len0_64 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_64
		var _nw391 _dafny.Array
		_ = _nw391
		if _len0_64.Cmp(_dafny.Zero) == 0 {
			_nw391 = _dafny.NewArray(_len0_64)
		} else {
			var _init64 func(_dafny.Int) _dafny.Int = func(_2417_i7 _dafny.Int) _dafny.Int {
				return (_2417_i7).Minus((_this).F22())
			}
			_ = _init64
			var _element0_64 = _init64(_dafny.Zero)
			_ = _element0_64
			_nw391 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
			(_nw391).ArraySet1(_element0_64, 0)
			var _nativeLen0_64 = (_len0_64).Int()
			_ = _nativeLen0_64
			for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
				(_nw391).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
			}
		}
		_2416_v56 = _nw391
		var _index406 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_2416_v56), 0))
		_ = _index406
		(_2416_v56).ArraySet1((_this).F22(), (_index406).Int())
		var _2418_v57 T1
		_ = _2418_v57
		var _nw392 *C2 = New_C2_()
		_ = _nw392
		_nw392.Ctor__()
		_2418_v57 = _nw392
		var _2419_v58 _dafny.CodePoint
		_ = _2419_v58
		_2419_v58 = _dafny.CodePoint('n')
		var _2420_v59 _dafny.Set
		_ = _2420_v59
		_2420_v59 = _dafny.SetOf(_2337_v0)
		var _2421_v60 D5
		_ = _2421_v60
		_2421_v60 = Companion_D5_.Create_DC14_(_2419_v58, _2337_v0, (_2420_v59).Cardinality(), _2337_v0)
		var _2422_v61 _dafny.Map
		_ = _2422_v61
		_2422_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2421_v60, (_this).F22())
		var _2423_v62 _dafny.Map
		_ = _2423_v62
		_2423_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), (_2422_v61).Cardinality())
		var _index407 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_2416_v56), 0))
		_ = _index407
		(_2416_v56).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_2418_v57, _2418_v57)).Cardinality())).Minus((func() _dafny.Int {
			if (_2423_v62).Contains(_2419_v58) {
				return (_2423_v62).Get(_2419_v58).(_dafny.Int)
			}
			return (_this).F22()
		})()), (_index407).Int())
	}
}
func (_this *C12) M2(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) {
	{
		var _2424_v0 _dafny.CodePoint
		_ = _2424_v0
		_2424_v0 = _dafny.CodePoint('k')
		var _2425_v1 _dafny.Set
		_ = _2425_v1
		_2425_v1 = _dafny.SetOf((_this).F23(), _dafny.Companion_Sequence_.Update((_this).F23(), (Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32(((_this).F23()).Cardinality()))).Uint32(), _2424_v0))
		if ((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("jieavydw"), Companion_Default___.Fm19(p1, p1, globalState))).Union(_2425_v1)).Contains((_this).F23()) {
			var _2426_v2 _dafny.Sequence
			_ = _2426_v2
			_2426_v2 = _dafny.SeqOf(p1, !(!(!(p1))))
			var _2427_v3 _dafny.Sequence
			_ = _2427_v3
			_2427_v3 = _dafny.SeqOf((_2426_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2426_v2).Cardinality()), _dafny.IntOfUint32((_2426_v2).Cardinality()))).Uint32()).(bool))
			var _2428_v4 D18
			_ = _2428_v4
			_2428_v4 = Companion_D18_.Create_DC49_(p1, p1, _2424_v0, (_2427_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2427_v3).Cardinality()))).Uint32()).(bool))
			var _source33 D18 = _2428_v4
			_ = _source33
			if _source33.Is_DC49() {
				var _2429___mcc_h0 bool = _source33.Get_().(D18_DC49).Cf75
				_ = _2429___mcc_h0
				var _2430___mcc_h1 bool = _source33.Get_().(D18_DC49).Cf76
				_ = _2430___mcc_h1
				var _2431___mcc_h2 _dafny.CodePoint = _source33.Get_().(D18_DC49).Cf77
				_ = _2431___mcc_h2
				var _2432___mcc_h3 bool = _source33.Get_().(D18_DC49).Cf78
				_ = _2432___mcc_h3
				var _2433_cf78 bool = _2432___mcc_h3
				_ = _2433_cf78
				var _2434_cf77 _dafny.CodePoint = _2431___mcc_h2
				_ = _2434_cf77
				var _2435_cf76 bool = _2430___mcc_h1
				_ = _2435_cf76
				var _2436_cf75 bool = _2429___mcc_h0
				_ = _2436_cf75
				var _2437_v5 _dafny.Array
				_ = _2437_v5
				var _nw393 _dafny.Array = _dafny.NewArrayWithValue(Companion_D5_.Default(), _dafny.IntOfInt64(21))
				_ = _nw393
				_2437_v5 = _nw393
				var _2438_v6 _dafny.Map
				_ = _2438_v6
				_2438_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2433_cf78)
				var _2439_v7 _dafny.Set
				_ = _2439_v7
				_2439_v7 = _dafny.SetOf((_2438_v6).Cardinality())
				var _2440_v9 _dafny.MultiSet
				_ = _2440_v9
				_2440_v9 = _dafny.MultiSetOf(_2439_v7, func() _dafny.Set {
					var _coll86 = _dafny.NewBuilder()
					_ = _coll86
					for _iter89 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(713), _dafny.IntOfInt64(165))); ; {
						_compr_86, _ok89 := _iter89()
						if !_ok89 {
							break
						}
						var _2441_v8 _dafny.Int
						_2441_v8 = interface{}(_compr_86).(_dafny.Int)
						if ((_dafny.IntOfInt64(713)).Cmp(_2441_v8) <= 0) && ((_2441_v8).Cmp(_dafny.IntOfInt64(165)) < 0) {
							_coll86.Add((_2441_v8).Times(p0))
						}
					}
					return _coll86.ToSet()
				}())
				var _2442_v10 D5
				_ = _2442_v10
				_2442_v10 = Companion_D5_.Create_DC13_(_2440_v9)
				var _index408 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_2437_v5), 0))
				_ = _index408
				(_2437_v5).ArraySet1(_2442_v10, (_index408).Int())
				var _pat_let_tv51 = _2439_v7
				_ = _pat_let_tv51
				var _pat_let_tv52 = _2439_v7
				_ = _pat_let_tv52
				var _pat_let_tv53 = _2440_v9
				_ = _pat_let_tv53
				var _index409 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_2437_v5), 0))
				_ = _index409
				(_2437_v5).ArraySet1(func(_pat_let49_0 D5) D5 {
					return func(_2443_dt__update__tmp_h0 D5) D5 {
						return func(_pat_let50_0 _dafny.MultiSet) D5 {
							return func(_2444_dt__update_hcf17_h0 _dafny.MultiSet) D5 {
								return Companion_D5_.Create_DC13_(_2444_dt__update_hcf17_h0)
							}(_pat_let50_0)
						}((_dafny.MultiSetOf(_pat_let_tv51, _pat_let_tv52)).Difference(_pat_let_tv53))
					}(_pat_let49_0)
				}(_2442_v10), (_index409).Int())
				var _2445_v11 _dafny.MultiSet
				_ = _2445_v11
				_2445_v11 = _dafny.MultiSetOf(false)
				var _2446_v12 _dafny.Array
				_ = _2446_v12
				var _nw394 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
				_ = _nw394
				_2446_v12 = _nw394
				var _2447_v13 D2
				_ = _2447_v13
				_2447_v13 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC7_((func() bool {
					if (_2438_v6).Contains((_this).F22()) {
						return (_2438_v6).Get((_this).F22()).(bool)
					}
					return _2436_cf75
				})(), _2445_v11, _2446_v12))
				var _2448_v14 *C3
				_ = _2448_v14
				var _nw395 *C3 = New_C3_()
				_ = _nw395
				_nw395.Ctor__(_2447_v13, false)
				_2448_v14 = _nw395
				var _rhs478 _dafny.Int = (p2).Minus((_this).F22())
				_ = _rhs478
				var _rhs479 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2426_v2, _dafny.SeqOf(_2448_v14.F25))).Cardinality())).Cmp(p0) < 0
				_ = _rhs479
				var _rhs480 *C3 = _2448_v14
				_ = _rhs480
				var _lhs323 *GlobalState = globalState
				_ = _lhs323
				_lhs323.F4 = _rhs478
				_2433_cf78 = _rhs479
				_2448_v14 = _rhs480
				_2433_cf78 = !(_2436_cf75)
				var _2449_v15 _dafny.MultiSet
				_ = _2449_v15
				_2449_v15 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq(_2427_v3)).Cardinality(), (_this).F22())
				var _2450_v16 _dafny.Sequence
				_ = _2450_v16
				_2450_v16 = _dafny.SeqOf((_2449_v15).Cardinality(), _dafny.IntOfUint32(((_this).F23()).Cardinality()))
				(globalState).F15 = _dafny.IntOfUint32((_2450_v16).Cardinality())
			} else if _source33.Is_DC48() {
				var _2451___mcc_h4 _dafny.Map = _source33.Get_().(D18_DC48).Cf74
				_ = _2451___mcc_h4
				var _2452_cf74 _dafny.Map = _2451___mcc_h4
				_ = _2452_cf74
				var _2453_v17 _dafny.Array
				_ = _2453_v17
				var _nw396 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw396
				_2453_v17 = _nw396
				var _2454_v18 _dafny.Array
				_ = _2454_v18
				var _len0_65 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_65
				var _nw397 _dafny.Array
				_ = _nw397
				if _len0_65.Cmp(_dafny.Zero) == 0 {
					_nw397 = _dafny.NewArray(_len0_65)
				} else {
					var _init65 func(_dafny.Int) _dafny.CodePoint = (func(_2455_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2456_i0 _dafny.Int) _dafny.CodePoint {
							return _2455_v0
						}
					})(_2424_v0)
					_ = _init65
					var _element0_65 = _init65(_dafny.Zero)
					_ = _element0_65
					_nw397 = _dafny.NewArrayFromExample(_element0_65, nil, _len0_65)
					(_nw397).ArraySet1CodePoint(_element0_65, 0)
					var _nativeLen0_65 = (_len0_65).Int()
					_ = _nativeLen0_65
					for _i0_65 := 1; _i0_65 < _nativeLen0_65; _i0_65++ {
						(_nw397).ArraySet1CodePoint(_init65(_dafny.IntOf(_i0_65)), _i0_65)
					}
				}
				_2454_v18 = _nw397
				var _2457_v19 _dafny.Sequence
				_ = _2457_v19
				_2457_v19 = _dafny.SeqOf(_2454_v18)
				var _index410 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_2453_v17), 0))
				_ = _index410
				(_2453_v17).ArraySet1(_2457_v19, (_index410).Int())
				var _2458_v20 _dafny.Sequence
				_ = _2458_v20
				_2458_v20 = _dafny.SeqOf(_2457_v19)
				var _2459_v21 _dafny.Sequence
				_ = _2459_v21
				_2459_v21 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2454_v18, _2454_v18, _2454_v18), _2457_v19), _dafny.Companion_Sequence_.Update((_2458_v20).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2458_v20).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F22()), _dafny.IntOfUint32(((_2458_v20).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2458_v20).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _2454_v18), _2457_v19)
				var _index411 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_2453_v17), 0))
				_ = _index411
				(_2453_v17).ArraySet1((_2459_v21).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32(((_this).F23()).Cardinality())).Minus((_dafny.Zero).Minus((_this).F22())), _dafny.IntOfUint32((_2459_v21).Cardinality()))).Uint32()).(_dafny.Sequence), (_index411).Int())
				var _2460_v22 _dafny.MultiSet
				_ = _2460_v22
				_2460_v22 = _dafny.MultiSetOf(p1, p1)
				var _2461_v23 D1
				_ = _2461_v23
				_2461_v23 = Companion_D1_.Create_DC2_((func() _dafny.Int {
					if (_2460_v22).Contains(p1) {
						return (_2460_v22).Multiplicity(p1)
					}
					return _dafny.IntOfInt64(143)
				})())
				var _2462_v24 D20
				_ = _2462_v24
				_2462_v24 = Companion_D20_.Create_DC55_(p2, (_this).F22(), p1, p1)
				var _2463_v25 D8
				_ = _2463_v25
				_2463_v25 = Companion_D8_.Create_DC24_(p0, false)
				var _2464_v26 _dafny.Map
				_ = _2464_v26
				_2464_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2463_v25))
				var _2465_v27 _dafny.Map
				_ = _2465_v27
				_2465_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2463_v25)
				var _2466_v28 _dafny.Set
				_ = _2466_v28
				_2466_v28 = _dafny.SetOf(!(!(p1)), !(p1))
				var _2467_v29 _dafny.Array
				_ = _2467_v29
				var _nwElement0_99 _dafny.Int = (_this).F22()
				_ = _nwElement0_99
				var _nw398 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_99, nil, _dafny.IntOfInt64(11))
				_ = _nw398
				(_nw398).ArraySet1(_nwElement0_99, 0)
				(_nw398).ArraySet1(p0, 1)
				(_nw398).ArraySet1(_dafny.IntOfInt64(678), 2)
				(_nw398).ArraySet1(p2, 3)
				(_nw398).ArraySet1((_this).F22(), 4)
				(_nw398).ArraySet1((_2462_v24).Dtor_cf85(), 5)
				(_nw398).ArraySet1(((func() _dafny.Map {
					if (_2464_v26).Contains(p2) {
						return (_2464_v26).Get(p2).(_dafny.Map)
					}
					return _2465_v27
				})()).Cardinality(), 6)
				(_nw398).ArraySet1((_2460_v22).Cardinality(), 7)
				(_nw398).ArraySet1((_2466_v28).Cardinality(), 8)
				(_nw398).ArraySet1(p2, 9)
				(_nw398).ArraySet1(_dafny.IntOfInt64(-907), 10)
				_2467_v29 = _nw398
				var _2468_v30 _dafny.Map
				_ = _2468_v30
				_2468_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2467_v29, _dafny.MultiSetOf(p1))
				var _2469_v31 _dafny.Array
				_ = _2469_v31
				var _nwElement0_100 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_2426_v2, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm19(p1, p1, globalState)).Cardinality()), _dafny.IntOfUint32((_2426_v2).Cardinality()))).Uint32(), p1))).Union(_2460_v22)
				_ = _nwElement0_100
				var _nw399 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_100, nil, _dafny.IntOfInt64(21))
				_ = _nw399
				(_nw399).ArraySet1(_nwElement0_100, 0)
				(_nw399).ArraySet1(_dafny.MultiSetOf(!(p1), p1), 1)
				(_nw399).ArraySet1(_dafny.MultiSetOf(p1), 2)
				(_nw399).ArraySet1((_2460_v22).Update(false, Companion_Default___.Abs(p0)), 3)
				(_nw399).ArraySet1((_2460_v22).Union(Companion_Default___.Fm34((_this).F22(), p0, p1, globalState)), 4)
				(_nw399).ArraySet1(_dafny.MultiSetFromSeq(_2427_v3), 5)
				(_nw399).ArraySet1(_2460_v22, 6)
				(_nw399).ArraySet1((_2460_v22).Update(false, Companion_Default___.Abs(p2)), 7)
				(_nw399).ArraySet1(_2460_v22, 8)
				(_nw399).ArraySet1(_2460_v22, 9)
				(_nw399).ArraySet1(_dafny.MultiSetOf(!(p1)), 10)
				(_nw399).ArraySet1((_2460_v22).Union(_2460_v22), 11)
				(_nw399).ArraySet1(_2460_v22, 12)
				(_nw399).ArraySet1(_2460_v22, 13)
				(_nw399).ArraySet1((_2460_v22).Union(_2460_v22), 14)
				(_nw399).ArraySet1(_2460_v22, 15)
				(_nw399).ArraySet1((func() _dafny.MultiSet {
					if p1 {
						return _2460_v22
					}
					return _2460_v22
				})(), 16)
				(_nw399).ArraySet1(_2460_v22, 17)
				(_nw399).ArraySet1(_dafny.MultiSetOf(p1), 18)
				(_nw399).ArraySet1((func() _dafny.MultiSet {
					if p1 {
						return _2460_v22
					}
					return _2460_v22
				})(), 19)
				(_nw399).ArraySet1((func() _dafny.MultiSet {
					if (_2468_v30).Contains(_2467_v29) {
						return (_2468_v30).Get(_2467_v29).(_dafny.MultiSet)
					}
					return _2460_v22
				})(), 20)
				_2469_v31 = _nw399
				var _2470_v32 _dafny.Map
				_ = _2470_v32
				_2470_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_2427_v3).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2427_v3).Cardinality()))).Uint32()).(bool))
				var _index412 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_2454_v18), 0))
				_ = _index412
				(_2454_v18).ArraySet1CodePoint(((_this).F23()).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32(((_this).F23()).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index412).Int())
				var _2471_v33 _dafny.Sequence
				_ = _2471_v33
				_2471_v33 = _dafny.SeqOf(_2469_v31, _2469_v31)
				var _index413 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_2454_v18), 0))
				_ = _index413
				var _rhs481 D1 = _2461_v23
				_ = _rhs481
				var _rhs482 _dafny.Array = (_2471_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.IntOfUint32((_2471_v33).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs482
				var _rhs483 _dafny.Map = _2470_v32
				_ = _rhs483
				var _rhs484 _dafny.CodePoint = _2424_v0
				_ = _rhs484
				var _rhs485 _dafny.Int = (_dafny.Zero).Minus(p2)
				_ = _rhs485
				var _lhs324 _dafny.Array = _2454_v18
				_ = _lhs324
				var _lhs325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_2454_v18), 0))
				_ = _lhs325
				var _lhs326 *GlobalState = globalState
				_ = _lhs326
				_2461_v23 = _rhs481
				_2469_v31 = _rhs482
				_2470_v32 = _rhs483
				(_lhs324).ArraySet1CodePoint(_rhs484, (_lhs325).Int())
				_lhs326.F13 = _rhs485
				(globalState).F15 = (_dafny.Zero).Minus(p0)
				var _2472_v34 _dafny.Sequence
				_ = _2472_v34
				_2472_v34 = _dafny.SeqOf(p0, p2)
				(globalState).F15 = (_2472_v34).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.IntOfUint32((_2472_v34).Cardinality()))).Uint32()).(_dafny.Int)
			} else {
				var _2473___mcc_h5 D18 = _source33.Get_().(D18_DC50).Cf79
				_ = _2473___mcc_h5
				var _2474_cf79 D18 = _2473___mcc_h5
				_ = _2474_cf79
				var _2475_v35 _dafny.Map
				_ = _2475_v35
				_2475_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _2476_v36 _dafny.Map
				_ = _2476_v36
				_2476_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2475_v35)
				var _2477_v37 _dafny.Array
				_ = _2477_v37
				var _nwElement0_101 _dafny.CodePoint = _2424_v0
				_ = _nwElement0_101
				var _nw400 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_101, nil, _dafny.IntOfInt64(20))
				_ = _nw400
				(_nw400).ArraySet1CodePoint(_nwElement0_101, 0)
				(_nw400).ArraySet1CodePoint(_2424_v0, 1)
				(_nw400).ArraySet1CodePoint(_2424_v0, 2)
				(_nw400).ArraySet1CodePoint(_2424_v0, 3)
				(_nw400).ArraySet1CodePoint(_2424_v0, 4)
				(_nw400).ArraySet1CodePoint(_2424_v0, 5)
				(_nw400).ArraySet1CodePoint((func() _dafny.CodePoint {
					if p1 {
						return _2424_v0
					}
					return _2424_v0
				})(), 6)
				(_nw400).ArraySet1CodePoint(_2424_v0, 7)
				(_nw400).ArraySet1CodePoint(Companion_Default___.Fm38(p1, globalState), 8)
				(_nw400).ArraySet1CodePoint(_2424_v0, 9)
				(_nw400).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_this).Fm1(_2476_v36, _dafny.IntOfUint32((_2426_v2).Cardinality()), globalState) {
						return _2424_v0
					}
					return _2424_v0
				})(), 10)
				(_nw400).ArraySet1CodePoint(_2424_v0, 11)
				(_nw400).ArraySet1CodePoint(_2424_v0, 12)
				(_nw400).ArraySet1CodePoint((func() _dafny.CodePoint {
					if p1 {
						return _2424_v0
					}
					return _2424_v0
				})(), 13)
				(_nw400).ArraySet1CodePoint(_2424_v0, 14)
				(_nw400).ArraySet1CodePoint(_2424_v0, 15)
				(_nw400).ArraySet1CodePoint(_2424_v0, 16)
				(_nw400).ArraySet1CodePoint(_2424_v0, 17)
				(_nw400).ArraySet1CodePoint(_2424_v0, 18)
				(_nw400).ArraySet1CodePoint(_2424_v0, 19)
				_2477_v37 = _nw400
				_2477_v37 = _2477_v37
				var _2478_v38 _dafny.Map
				_ = _2478_v38
				_2478_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F22()), p2)
				var _2479_v39 _dafny.Map
				_ = _2479_v39
				_2479_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F23(), Companion_D18_.Create_DC48_(_2478_v38))
				var _2480_v41 _dafny.MultiSet
				_ = _2480_v41
				_2480_v41 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-696))).Uint32(), func(coer131 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg131 _dafny.Int) interface{} {
						return coer131(arg131)
					}
				}((func(_2481_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2482_i1 _dafny.Int) _dafny.CodePoint {
						return _2481_v0
					}
				})(_2424_v0))))
				(globalState).F15 = ((_this).F22()).Times(((_2479_v39).Merge(func() _dafny.Map {
					var _coll87 = _dafny.NewMapBuilder()
					_ = _coll87
					for _iter90 := _dafny.Iterate((_2480_v41).Elements()); ; {
						_compr_87, _ok90 := _iter90()
						if !_ok90 {
							break
						}
						var _2483_v40 _dafny.Sequence
						_2483_v40 = interface{}(_compr_87).(_dafny.Sequence)
						if (_2480_v41).Contains(_2483_v40) {
							_coll87.Add(_2483_v40, Companion_D18_.Create_DC48_(_2478_v38))
						}
					}
					return _coll87.ToMap()
				}())).Cardinality())
				var _2484_v42 bool
				_ = _2484_v42
				_2484_v42 = true
				_2484_v42 = _2484_v42
				var _2485_v43 _dafny.Set
				_ = _2485_v43
				_2485_v43 = _dafny.SetOf(p1)
				var _2486_v44 _dafny.Sequence
				_ = _2486_v44
				_2486_v44 = _dafny.SeqOf((_this).F22(), (_2485_v43).Cardinality(), p2)
				var _2487_v45 D6
				_ = _2487_v45
				_2487_v45 = Companion_D6_.Create_DC17_((_this).Fm1(_2476_v36, _dafny.IntOfUint32((_2486_v44).Cardinality()), globalState), !(_2484_v42), p1)
				_2484_v42 = (_2487_v45).Dtor_cf24()
			}
			var _2488_v46 _dafny.Set
			_ = _2488_v46
			_2488_v46 = _dafny.SetOf(!(p1), p1)
			var _2489_v47 _dafny.Sequence
			_ = _2489_v47
			_2489_v47 = _dafny.SeqOf(p2, (_2488_v46).Cardinality(), (_this).F22())
			var _2490_v48 _dafny.Sequence
			_ = _2490_v48
			_2490_v48 = _dafny.SeqOf(_2489_v47, _2489_v47)
			_2490_v48 = _dafny.Companion_Sequence_.Update(_2490_v48, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(p0))).Cardinality()), _dafny.IntOfUint32((_2490_v48).Cardinality()))).Uint32(), _2489_v47)
			(globalState).F7 = Companion_Default___.Fm17(p0, globalState)
			_2427_v3 = _2427_v3
			var _2491_v49 _dafny.Array
			_ = _2491_v49
			var _len0_66 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_66
			var _nw401 _dafny.Array
			_ = _nw401
			if _len0_66.Cmp(_dafny.Zero) == 0 {
				_nw401 = _dafny.NewArray(_len0_66)
			} else {
				var _init66 func(_dafny.Int) bool = (func(_2492_p1 bool) func(_dafny.Int) bool {
					return func(_2493_i2 _dafny.Int) bool {
						return _2492_p1
					}
				})(p1)
				_ = _init66
				var _element0_66 = _init66(_dafny.Zero)
				_ = _element0_66
				_nw401 = _dafny.NewArrayFromExample(_element0_66, nil, _len0_66)
				(_nw401).ArraySet1(_element0_66, 0)
				var _nativeLen0_66 = (_len0_66).Int()
				_ = _nativeLen0_66
				for _i0_66 := 1; _i0_66 < _nativeLen0_66; _i0_66++ {
					(_nw401).ArraySet1(_init66(_dafny.IntOf(_i0_66)), _i0_66)
				}
			}
			_2491_v49 = _nw401
			var _index414 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_2491_v49), 0))
			_ = _index414
			(_2491_v49).ArraySet1(p1, (_index414).Int())
			var _index415 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_2491_v49), 0))
			_ = _index415
			(_2491_v49).ArraySet1(p1, (_index415).Int())
		} else {
			var _2494_v50 _dafny.Array
			_ = _2494_v50
			var _len0_67 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_67
			var _nw402 _dafny.Array
			_ = _nw402
			if _len0_67.Cmp(_dafny.Zero) == 0 {
				_nw402 = _dafny.NewArray(_len0_67)
			} else {
				var _init67 func(_dafny.Int) bool = (func(_2495_p1 bool) func(_dafny.Int) bool {
					return func(_2496_i3 _dafny.Int) bool {
						return _2495_p1
					}
				})(p1)
				_ = _init67
				var _element0_67 = _init67(_dafny.Zero)
				_ = _element0_67
				_nw402 = _dafny.NewArrayFromExample(_element0_67, nil, _len0_67)
				(_nw402).ArraySet1(_element0_67, 0)
				var _nativeLen0_67 = (_len0_67).Int()
				_ = _nativeLen0_67
				for _i0_67 := 1; _i0_67 < _nativeLen0_67; _i0_67++ {
					(_nw402).ArraySet1(_init67(_dafny.IntOf(_i0_67)), _i0_67)
				}
			}
			_2494_v50 = _nw402
			_2494_v50 = _2494_v50
			var _2497_v51 bool
			_ = _2497_v51
			_2497_v51 = true
			var _2498_v52 T0
			_ = _2498_v52
			var _nw403 *C11 = New_C11_()
			_ = _nw403
			_nw403.Ctor__()
			_2498_v52 = _nw403
			var _2499_v53 _dafny.Set
			_ = _2499_v53
			_2499_v53 = _dafny.SetOf(_2498_v52, _2498_v52)
			var _2500_v54 D7
			_ = _2500_v54
			_2500_v54 = Companion_D7_.Create_DC22_(p1, p0, p1, (_2499_v53).Cardinality())
			var _2501_v55 _dafny.Map
			_ = _2501_v55
			_2501_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2497_v51)
			var _2502_v56 D7
			_ = _2502_v56
			_2502_v56 = Companion_D7_.Create_DC22_(!((_2500_v54).Dtor_cf35()), (_dafny.Zero).Minus(p0), (func() bool {
				if (_2501_v55).Contains(_2497_v51) {
					return (_2501_v55).Get(_2497_v51).(bool)
				}
				return true
			})(), p0)
			_2497_v51 = (_2502_v56).Dtor_cf37()
			var _2503_v57 _dafny.MultiSet
			_ = _2503_v57
			_2503_v57 = _dafny.MultiSetOf(_2497_v51)
			_2497_v51 = !(!(_2503_v57).Contains(p1))
			var _2504_v58 *C8
			_ = _2504_v58
			var _nw404 *C8 = New_C8_()
			_ = _nw404
			_nw404.Ctor__()
			_2504_v58 = _nw404
			_2494_v50 = _2494_v50
		}
		var _2505_v59 *C4
		_ = _2505_v59
		var _nw405 *C4 = New_C4_()
		_ = _nw405
		_nw405.Ctor__()
		_2505_v59 = _nw405
		var _2506_v61 _dafny.MultiSet
		_ = _2506_v61
		_2506_v61 = _dafny.MultiSetOf((_this).F22(), p0, (_this).F22())
		var _2507_v62 _dafny.Sequence
		_ = _2507_v62
		_2507_v62 = _dafny.SeqOf(func() _dafny.Map {
			var _coll88 = _dafny.NewMapBuilder()
			_ = _coll88
			for _iter91 := _dafny.Iterate((_2506_v61).Elements()); ; {
				_compr_88, _ok91 := _iter91()
				if !_ok91 {
					break
				}
				var _2508_v60 _dafny.Int
				_2508_v60 = interface{}(_compr_88).(_dafny.Int)
				if (_2506_v61).Contains(_2508_v60) {
					_coll88.Add(Companion_Default___.SafeDivisionInt(_2508_v60, _dafny.IntOfInt64(13)), p1)
				}
			}
			return _coll88.ToMap()
		}())
		var _2509_v63 D14
		_ = _2509_v63
		_2509_v63 = Companion_D14_.Create_DC35_(_2507_v62)
		var _2510_i4 _dafny.Int
		_ = _2510_i4
		_2510_i4 = _dafny.Zero
		{
			var _pat_let_tv54 = p1
			_ = _pat_let_tv54
			var _pat_let_tv55 = p1
			_ = _pat_let_tv55
			var _pat_let_tv56 = p1
			_ = _pat_let_tv56
			for func(_source35 D14) bool {
				if _source35.Is_DC36() {
					var _2583___mcc_h18 _dafny.Sequence = _source35.Get_().(D14_DC36).Cf53
					_ = _2583___mcc_h18
					var _2584_cf53 _dafny.Sequence = _2583___mcc_h18
					_ = _2584_cf53
					return _pat_let_tv54
				} else if _source35.Is_DC35() {
					var _2585___mcc_h19 _dafny.Sequence = _source35.Get_().(D14_DC35).Cf52
					_ = _2585___mcc_h19
					var _2586_cf52 _dafny.Sequence = _2585___mcc_h19
					_ = _2586_cf52
					return _pat_let_tv55
				} else {
					var _2587___mcc_h20 D14 = _source35.Get_().(D14_DC37).Cf54
					_ = _2587___mcc_h20
					var _2588_cf54 D14 = _2587___mcc_h20
					_ = _2588_cf54
					return !(_pat_let_tv56)
				}
			}(_2509_v63) {
				{
					if (_2510_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L16
					}
					_2510_i4 = (_2510_i4).Plus(_dafny.One)
					var _2511_v64 *C8
					_ = _2511_v64
					var _nw406 *C8 = New_C8_()
					_ = _nw406
					_nw406.Ctor__()
					_2511_v64 = _nw406
					var _2512_v65 D22
					_ = _2512_v65
					_2512_v65 = Companion_D22_.Create_DC60_((_this).F22(), (_this).F22())
					var _source34 D22 = _2512_v65
					_ = _source34
					if _source34.Is_DC60() {
						var _2513___mcc_h6 _dafny.Int = _source34.Get_().(D22_DC60).Cf95
						_ = _2513___mcc_h6
						var _2514___mcc_h7 _dafny.Int = _source34.Get_().(D22_DC60).Cf96
						_ = _2514___mcc_h7
						var _2515_cf96 _dafny.Int = _2514___mcc_h7
						_ = _2515_cf96
						var _2516_cf95 _dafny.Int = _2513___mcc_h6
						_ = _2516_cf95
						var _2517_v66 *C9
						_ = _2517_v66
						var _nw407 *C9 = New_C9_()
						_ = _nw407
						_nw407.Ctor__()
						_2517_v66 = _nw407
						var _2518_v67 bool
						_ = _2518_v67
						_2518_v67 = true
						_2518_v67 = true
						var _2519_v68 _dafny.Array
						_ = _2519_v68
						var _len0_68 _dafny.Int = _dafny.One
						_ = _len0_68
						var _nw408 _dafny.Array
						_ = _nw408
						if _len0_68.Cmp(_dafny.Zero) == 0 {
							_nw408 = _dafny.NewArray(_len0_68)
						} else {
							var _init68 func(_dafny.Int) bool = (func(_2520_p1 bool) func(_dafny.Int) bool {
								return func(_2521_i5 _dafny.Int) bool {
									return _2520_p1
								}
							})(p1)
							_ = _init68
							var _element0_68 = _init68(_dafny.Zero)
							_ = _element0_68
							_nw408 = _dafny.NewArrayFromExample(_element0_68, nil, _len0_68)
							(_nw408).ArraySet1(_element0_68, 0)
							var _nativeLen0_68 = (_len0_68).Int()
							_ = _nativeLen0_68
							for _i0_68 := 1; _i0_68 < _nativeLen0_68; _i0_68++ {
								(_nw408).ArraySet1(_init68(_dafny.IntOf(_i0_68)), _i0_68)
							}
						}
						_2519_v68 = _nw408
						_2519_v68 = _2519_v68
						(globalState).F7 = _2516_cf95
					} else if _source34.Is_DC61() {
						var _2522___mcc_h8 _dafny.Int = _source34.Get_().(D22_DC61).Cf97
						_ = _2522___mcc_h8
						var _2523___mcc_h9 _dafny.Int = _source34.Get_().(D22_DC61).Cf98
						_ = _2523___mcc_h9
						var _2524___mcc_h10 _dafny.Int = _source34.Get_().(D22_DC61).Cf99
						_ = _2524___mcc_h10
						var _2525___mcc_h11 _dafny.Sequence = _source34.Get_().(D22_DC61).Cf100
						_ = _2525___mcc_h11
						var _2526_cf100 _dafny.Sequence = _2525___mcc_h11
						_ = _2526_cf100
						var _2527_cf99 _dafny.Int = _2524___mcc_h10
						_ = _2527_cf99
						var _2528_cf98 _dafny.Int = _2523___mcc_h9
						_ = _2528_cf98
						var _2529_cf97 _dafny.Int = _2522___mcc_h8
						_ = _2529_cf97
						var _2530_v69 bool
						_ = _2530_v69
						_2530_v69 = true
						_2530_v69 = p1
						var _2531_v70 _dafny.Array
						_ = _2531_v70
						var _len0_69 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_69
						var _nw409 _dafny.Array
						_ = _nw409
						if _len0_69.Cmp(_dafny.Zero) == 0 {
							_nw409 = _dafny.NewArray(_len0_69)
						} else {
							var _init69 func(_dafny.Int) bool = (func(_2532_v0 _dafny.CodePoint) func(_dafny.Int) bool {
								return func(_2533_i6 _dafny.Int) bool {
									return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-459))).Uint32(), func(coer132 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg132 _dafny.Int) interface{} {
											return coer132(arg132)
										}
									}((func(_2534_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
										return func(_2535_i7 _dafny.Int) _dafny.CodePoint {
											return _2534_v0
										}
									})(_2532_v0))), _dafny.UnicodeSeqOfUtf8Bytes("claiu"))
								}
							})(_2424_v0)
							_ = _init69
							var _element0_69 = _init69(_dafny.Zero)
							_ = _element0_69
							_nw409 = _dafny.NewArrayFromExample(_element0_69, nil, _len0_69)
							(_nw409).ArraySet1(_element0_69, 0)
							var _nativeLen0_69 = (_len0_69).Int()
							_ = _nativeLen0_69
							for _i0_69 := 1; _i0_69 < _nativeLen0_69; _i0_69++ {
								(_nw409).ArraySet1(_init69(_dafny.IntOf(_i0_69)), _i0_69)
							}
						}
						_2531_v70 = _nw409
						_2531_v70 = _2531_v70
						var _2536_v71 _dafny.MultiSet
						_ = _2536_v71
						_2536_v71 = _dafny.MultiSetOf(p1)
						var _2537_v72 _dafny.Map
						_ = _2537_v72
						_2537_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2527_cf99, p2)
						var _2538_v73 _dafny.Sequence
						_ = _2538_v73
						_2538_v73 = _dafny.SeqOf(_2530_v69)
						var _2539_v74 _dafny.Set
						_ = _2539_v74
						_2539_v74 = _dafny.SetOf((_this).F22())
						var _2540_v75 D22
						_ = _2540_v75
						_2540_v75 = Companion_D22_.Create_DC62_(_2526_cf100, p2, (_2539_v74).Cardinality(), _2529_cf97, p1)
						var _2541_v76 _dafny.Array
						_ = _2541_v76
						var _nwElement0_102 _dafny.Int = (func() _dafny.Int {
							if (_2537_v72).Contains(p2) {
								return (_2537_v72).Get(p2).(_dafny.Int)
							}
							return _2527_cf99
						})()
						_ = _nwElement0_102
						var _nw410 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_102, nil, _dafny.IntOfInt64(23))
						_ = _nw410
						(_nw410).ArraySet1(_nwElement0_102, 0)
						(_nw410).ArraySet1(p2, 1)
						(_nw410).ArraySet1(_dafny.IntOfUint32((_2538_v73).Cardinality()), 2)
						(_nw410).ArraySet1(_dafny.IntOfInt64(343), 3)
						(_nw410).ArraySet1((_this).F22(), 4)
						(_nw410).ArraySet1(_2528_cf98, 5)
						(_nw410).ArraySet1(p0, 6)
						(_nw410).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2540_v75)).Cardinality(), 7)
						(_nw410).ArraySet1((_dafny.Zero).Minus(p2), 8)
						(_nw410).ArraySet1((_dafny.Zero).Minus(_2529_cf97), 9)
						(_nw410).ArraySet1(_2528_cf98, 10)
						(_nw410).ArraySet1(p2, 11)
						(_nw410).ArraySet1(_dafny.IntOfInt64(-744), 12)
						(_nw410).ArraySet1(_2529_cf97, 13)
						(_nw410).ArraySet1(p0, 14)
						(_nw410).ArraySet1((_this).F22(), 15)
						(_nw410).ArraySet1((_this).F22(), 16)
						(_nw410).ArraySet1(_dafny.IntOfInt64(661), 17)
						(_nw410).ArraySet1(p0, 18)
						(_nw410).ArraySet1((_this).Fm3(p1, globalState), 19)
						(_nw410).ArraySet1((_this).F22(), 20)
						(_nw410).ArraySet1(_2528_cf98, 21)
						(_nw410).ArraySet1(_dafny.IntOfUint32((_2526_cf100).Cardinality()), 22)
						_2541_v76 = _nw410
						var _2542_v77 D2
						_ = _2542_v77
						_2542_v77 = Companion_D2_.Create_DC5_(_2541_v76)
						(_2511_v64).M6(_dafny.IntOfUint32(((_this).F23()).Cardinality()), _2536_v71, (func() bool {
							if _2530_v69 {
								return p1
							}
							return p1
						})(), _2542_v77, globalState)
						_2424_v0 = _2424_v0
					} else if _source34.Is_DC62() {
						var _2543___mcc_h12 _dafny.Sequence = _source34.Get_().(D22_DC62).Cf101
						_ = _2543___mcc_h12
						var _2544___mcc_h13 _dafny.Int = _source34.Get_().(D22_DC62).Cf102
						_ = _2544___mcc_h13
						var _2545___mcc_h14 _dafny.Int = _source34.Get_().(D22_DC62).Cf103
						_ = _2545___mcc_h14
						var _2546___mcc_h15 _dafny.Int = _source34.Get_().(D22_DC62).Cf104
						_ = _2546___mcc_h15
						var _2547___mcc_h16 bool = _source34.Get_().(D22_DC62).Cf105
						_ = _2547___mcc_h16
						var _2548_cf105 bool = _2547___mcc_h16
						_ = _2548_cf105
						var _2549_cf104 _dafny.Int = _2546___mcc_h15
						_ = _2549_cf104
						var _2550_cf103 _dafny.Int = _2545___mcc_h14
						_ = _2550_cf103
						var _2551_cf102 _dafny.Int = _2544___mcc_h13
						_ = _2551_cf102
						var _2552_cf101 _dafny.Sequence = _2543___mcc_h12
						_ = _2552_cf101
						var _2553_v78 _dafny.Map
						_ = _2553_v78
						_2553_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if p1 {
								return p1
							}
							return p1
						})(), (_this).F22())
						_2553_v78 = _2553_v78
						var _2554_v79 _dafny.Array
						_ = _2554_v79
						var _len0_70 _dafny.Int = _dafny.IntOfInt64(21)
						_ = _len0_70
						var _nw411 _dafny.Array
						_ = _nw411
						if _len0_70.Cmp(_dafny.Zero) == 0 {
							_nw411 = _dafny.NewArray(_len0_70)
						} else {
							var _init70 func(_dafny.Int) _dafny.Int = (func(_2555_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_2556_i8 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeModuloInt(_2556_i8, _2555_p0)
								}
							})(p0)
							_ = _init70
							var _element0_70 = _init70(_dafny.Zero)
							_ = _element0_70
							_nw411 = _dafny.NewArrayFromExample(_element0_70, nil, _len0_70)
							(_nw411).ArraySet1(_element0_70, 0)
							var _nativeLen0_70 = (_len0_70).Int()
							_ = _nativeLen0_70
							for _i0_70 := 1; _i0_70 < _nativeLen0_70; _i0_70++ {
								(_nw411).ArraySet1(_init70(_dafny.IntOf(_i0_70)), _i0_70)
							}
						}
						_2554_v79 = _nw411
						var _index416 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_2554_v79), 0))
						_ = _index416
						(_2554_v79).ArraySet1(_2551_cf102, (_index416).Int())
						var _index417 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_2554_v79), 0))
						_ = _index417
						(_2554_v79).ArraySet1(_dafny.IntOfInt64(205), (_index417).Int())
						var _2557_v80 D2
						_ = _2557_v80
						_2557_v80 = Companion_D2_.Create_DC5_(_2554_v79)
						var _2558_v81 D2
						_ = _2558_v81
						_2558_v81 = Companion_D2_.Create_DC8_(_2557_v80)
						_2558_v81 = _2558_v81
						var _2559_v82 D22
						_ = _2559_v82
						_2559_v82 = Companion_D22_.Create_DC61_(_2551_cf102, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nwpuonys")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("goqrmbm")).Cardinality()), _2552_cf101)
						var _2560_v83 _dafny.Map
						_ = _2560_v83
						_2560_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
						var _2561_v84 _dafny.Map
						_ = _2561_v84
						_2561_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2560_v83, _2552_cf101)
						var _2562_v85 _dafny.Map
						_ = _2562_v85
						_2562_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2554_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_2554_v79), 0))).Int()).(_dafny.Int), _2552_cf101)
						var _2563_v86 _dafny.Array
						_ = _2563_v86
						var _nwElement0_103 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2552_cf101, (_this).F23())
						_ = _nwElement0_103
						var _nw412 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_103, nil, _dafny.IntOfInt64(18))
						_ = _nw412
						(_nw412).ArraySet1(_nwElement0_103, 0)
						(_nw412).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F23(), _2552_cf101), 1)
						(_nw412).ArraySet1((_this).F23(), 2)
						(_nw412).ArraySet1((_2559_v82).Dtor_cf100(), 3)
						(_nw412).ArraySet1((_this).F23(), 4)
						(_nw412).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_2559_v82).Dtor_cf100(), _2552_cf101), 5)
						(_nw412).ArraySet1(_2552_cf101, 6)
						(_nw412).ArraySet1(_dafny.Companion_Sequence_.Update(_2552_cf101, (Companion_Default___.SafeIndex((_2554_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_2554_v79), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2552_cf101).Cardinality()))).Uint32(), _dafny.CodePoint('o')), 7)
						(_nw412).ArraySet1(_dafny.Companion_Sequence_.Update(_2552_cf101, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-44), _dafny.IntOfUint32((_2552_cf101).Cardinality()))).Uint32(), _2424_v0), 8)
						(_nw412).ArraySet1(_2552_cf101, 9)
						(_nw412).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
							if (_2561_v84).Contains(_2560_v83) {
								return (_2561_v84).Get(_2560_v83).(_dafny.Sequence)
							}
							return (_this).F23()
						})(), (_this).F23()), 10)
						(_nw412).ArraySet1((_this).F23(), 11)
						(_nw412).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("s"), 12)
						(_nw412).ArraySet1((_this).F23(), 13)
						(_nw412).ArraySet1(_2552_cf101, 14)
						(_nw412).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jfd"), 15)
						(_nw412).ArraySet1((func() _dafny.Sequence {
							if (_2562_v85).Contains((_2511_v64).Fm2((_dafny.Zero).Minus(_2550_cf103), globalState)) {
								return (_2562_v85).Get((_2511_v64).Fm2((_dafny.Zero).Minus(_2550_cf103), globalState)).(_dafny.Sequence)
							}
							return (_this).F23()
						})(), 16)
						(_nw412).ArraySet1((_this).F23(), 17)
						_2563_v86 = _nw412
						var _index418 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_2563_v86), 0))
						_ = _index418
						(_2563_v86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bnimhcmi"), (_this).F23()), (_index418).Int())
						var _2564_v87 _dafny.Sequence
						_ = _2564_v87
						_2564_v87 = _dafny.SeqOf(_2548_cf105)
						var _2565_v88 _dafny.Array
						_ = _2565_v88
						var _len0_71 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_71
						var _nw413 _dafny.Array
						_ = _nw413
						if _len0_71.Cmp(_dafny.Zero) == 0 {
							_nw413 = _dafny.NewArray(_len0_71)
						} else {
							var _init71 func(_dafny.Int) bool = func(_2566_i9 _dafny.Int) bool {
								return false
							}
							_ = _init71
							var _element0_71 = _init71(_dafny.Zero)
							_ = _element0_71
							_nw413 = _dafny.NewArrayFromExample(_element0_71, nil, _len0_71)
							(_nw413).ArraySet1(_element0_71, 0)
							var _nativeLen0_71 = (_len0_71).Int()
							_ = _nativeLen0_71
							for _i0_71 := 1; _i0_71 < _nativeLen0_71; _i0_71++ {
								(_nw413).ArraySet1(_init71(_dafny.IntOf(_i0_71)), _i0_71)
							}
						}
						_2565_v88 = _nw413
						var _2567_v89 D0
						_ = _2567_v89
						_2567_v89 = Companion_D0_.Create_DC1_(_2565_v88)
						var _2568_v90 _dafny.Map
						_ = _2568_v90
						_2568_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2567_v89, _2564_v87)
						var _2569_v91 T1
						_ = _2569_v91
						var _nw414 *C5 = New_C5_()
						_ = _nw414
						_nw414.Ctor__()
						_2569_v91 = _nw414
						var _2570_v92 _dafny.Map
						_ = _2570_v92
						_2570_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2569_v91, _2564_v87)
						var _2571_v93 _dafny.Array
						_ = _2571_v93
						var _nwElement0_104 _dafny.Sequence = _2564_v87
						_ = _nwElement0_104
						var _nw415 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_104, nil, _dafny.IntOfInt64(26))
						_ = _nw415
						(_nw415).ArraySet1(_nwElement0_104, 0)
						(_nw415).ArraySet1(_2564_v87, 1)
						(_nw415).ArraySet1(_dafny.SeqOf(_2548_cf105, _2548_cf105), 2)
						(_nw415).ArraySet1(_dafny.SeqOf(_2548_cf105), 3)
						(_nw415).ArraySet1(_2564_v87, 4)
						(_nw415).ArraySet1(_dafny.SeqOf((_2505_v59).Fm8(_2549_cf104, _2548_cf105, _dafny.IntOfInt64(302), globalState)), 5)
						(_nw415).ArraySet1((Companion_D14_.Create_DC36_(_dafny.SeqOf(p1, _2548_cf105))).Dtor_cf53(), 6)
						(_nw415).ArraySet1(_2564_v87, 7)
						(_nw415).ArraySet1(_dafny.SeqOf(p1), 8)
						(_nw415).ArraySet1(_2564_v87, 9)
						(_nw415).ArraySet1(_2564_v87, 10)
						(_nw415).ArraySet1(_2564_v87, 11)
						(_nw415).ArraySet1(_2564_v87, 12)
						(_nw415).ArraySet1((func() _dafny.Sequence {
							if (_2568_v90).Contains(_2567_v89) {
								return (_2568_v90).Get(_2567_v89).(_dafny.Sequence)
							}
							return _2564_v87
						})(), 13)
						(_nw415).ArraySet1(_2564_v87, 14)
						(_nw415).ArraySet1(_dafny.Companion_Sequence_.Update(_2564_v87, (Companion_Default___.SafeIndex(_2549_cf104, _dafny.IntOfUint32((_2564_v87).Cardinality()))).Uint32(), _2548_cf105), 15)
						(_nw415).ArraySet1(_2564_v87, 16)
						(_nw415).ArraySet1(_2564_v87, 17)
						(_nw415).ArraySet1(_2564_v87, 18)
						(_nw415).ArraySet1(_2564_v87, 19)
						(_nw415).ArraySet1(_2564_v87, 20)
						(_nw415).ArraySet1(Companion_Default___.Fm27(globalState), 21)
						(_nw415).ArraySet1(_2564_v87, 22)
						(_nw415).ArraySet1(_2564_v87, 23)
						(_nw415).ArraySet1(_2564_v87, 24)
						(_nw415).ArraySet1((func() _dafny.Sequence {
							if (_2570_v92).Contains(_2569_v91) {
								return (_2570_v92).Get(_2569_v91).(_dafny.Sequence)
							}
							return _2564_v87
						})(), 25)
						_2571_v93 = _nw415
						var _2572_v94 _dafny.Sequence
						_ = _2572_v94
						_2572_v94 = _dafny.SeqOf(_2571_v93, _2571_v93, _2571_v93)
						var _2573_v95 _dafny.Array
						_ = _2573_v95
						var _nwElement0_105 _dafny.Array = _2571_v93
						_ = _nwElement0_105
						var _nw416 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_105, nil, _dafny.IntOfInt64(17))
						_ = _nw416
						(_nw416).ArraySet1(_nwElement0_105, 0)
						(_nw416).ArraySet1(_2571_v93, 1)
						(_nw416).ArraySet1(_2571_v93, 2)
						(_nw416).ArraySet1(_2571_v93, 3)
						(_nw416).ArraySet1(_2571_v93, 4)
						(_nw416).ArraySet1(_2571_v93, 5)
						(_nw416).ArraySet1(_2571_v93, 6)
						(_nw416).ArraySet1(_2571_v93, 7)
						(_nw416).ArraySet1(_2571_v93, 8)
						(_nw416).ArraySet1(_2571_v93, 9)
						(_nw416).ArraySet1(_2571_v93, 10)
						(_nw416).ArraySet1(_2571_v93, 11)
						(_nw416).ArraySet1((_2572_v94).Select((Companion_Default___.SafeIndex(_2550_cf103, _dafny.IntOfUint32((_2572_v94).Cardinality()))).Uint32()).(_dafny.Array), 12)
						(_nw416).ArraySet1(_2571_v93, 13)
						(_nw416).ArraySet1(_2571_v93, 14)
						(_nw416).ArraySet1(_2571_v93, 15)
						(_nw416).ArraySet1(_2571_v93, 16)
						_2573_v95 = _nw416
						var _index419 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_2573_v95), 0))
						_ = _index419
						(_2573_v95).ArraySet1(_2571_v93, (_index419).Int())
						var _index420 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_2563_v86), 0))
						_ = _index420
						var _index421 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_2573_v95), 0))
						_ = _index421
						var _rhs486 _dafny.Sequence = (_this).F23()
						_ = _rhs486
						var _rhs487 _dafny.Array = (func() _dafny.Array {
							if p1 {
								return _2571_v93
							}
							return _2571_v93
						})()
						_ = _rhs487
						var _rhs488 _dafny.Int = (_dafny.Zero).Minus((_2549_cf104).Minus((_2554_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(350), _dafny.ArrayLen((_2554_v79), 0))).Int()).(_dafny.Int)))
						_ = _rhs488
						var _rhs489 _dafny.CodePoint = _2424_v0
						_ = _rhs489
						var _lhs327 _dafny.Array = _2563_v86
						_ = _lhs327
						var _lhs328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_2563_v86), 0))
						_ = _lhs328
						var _lhs329 _dafny.Array = _2573_v95
						_ = _lhs329
						var _lhs330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_2573_v95), 0))
						_ = _lhs330
						var _lhs331 *GlobalState = globalState
						_ = _lhs331
						(_lhs327).ArraySet1(_rhs486, (_lhs328).Int())
						(_lhs329).ArraySet1(_rhs487, (_lhs330).Int())
						_lhs331.F15 = _rhs488
						_2424_v0 = _rhs489
					} else {
						var _2574___mcc_h17 _dafny.MultiSet = _source34.Get_().(D22_DC59).Cf94
						_ = _2574___mcc_h17
						var _2575_cf94 _dafny.MultiSet = _2574___mcc_h17
						_ = _2575_cf94
						var _2576_v96 D8
						_ = _2576_v96
						_2576_v96 = Companion_D8_.Create_DC25_(((_dafny.Zero).Minus(p2)).Times((_this).F22()))
						_2576_v96 = _2576_v96
						(globalState).F13 = (func() _dafny.Int {
							if p1 {
								return p2
							}
							return p0
						})()
						var _2577_v97 bool
						_ = _2577_v97
						_2577_v97 = false
						_2577_v97 = _2577_v97
						var _2578_v98 _dafny.Sequence
						_ = _2578_v98
						_2578_v98 = _dafny.UnicodeSeqOfUtf8Bytes("vdybm")
						var _rhs490 _dafny.Sequence = (_this).F23()
						_ = _rhs490
						var _rhs491 _dafny.MultiSet = ((_2506_v61).Intersection(Companion_Default___.Fm31((_this).F22(), (_this).F22(), globalState))).Difference((_dafny.MultiSetOf((_this).F22(), p0, _dafny.IntOfInt64(-8))).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_2578_v98).Cardinality()))))
						_ = _rhs491
						_2578_v98 = _rhs490
						_2506_v61 = _rhs491
					}
					var _2579_v99 bool
					_ = _2579_v99
					_2579_v99 = true
					_2579_v99 = _2579_v99
					var _2580_v100 _dafny.Sequence
					_ = _2580_v100
					_2580_v100 = _dafny.UnicodeSeqOfUtf8Bytes("kkiv")
					_2580_v100 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-119))).Uint32(), func(coer133 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg133 _dafny.Int) interface{} {
							return coer133(arg133)
						}
					}((func(_2581_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2582_i10 _dafny.Int) _dafny.CodePoint {
							return _2581_v0
						}
					})(_2424_v0)))
					goto C16
				}
			C16:
			}
			goto L16
		}
	L16:
		var _2589_v101 _dafny.Array
		_ = _2589_v101
		var _nw417 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw417
		_2589_v101 = _nw417
		var _2590_v102 _dafny.Sequence
		_ = _2590_v102
		_2590_v102 = _dafny.UnicodeSeqOfUtf8Bytes("fwc")
		var _rhs492 _dafny.Array = _2589_v101
		_ = _rhs492
		var _rhs493 _dafny.Sequence = (_this).F23()
		_ = _rhs493
		_2589_v101 = _rhs492
		_2590_v102 = _rhs493
		var _2591_v103 D6
		_ = _2591_v103
		_2591_v103 = Companion_D6_.Create_DC17_(p1, true, p1)
		var _2592_i11 _dafny.Int
		_ = _2592_i11
		_2592_i11 = _dafny.Zero
		{
			for ((func() D6 {
				if p1 {
					return _2591_v103
				}
				return _2591_v103
			})()).Dtor_cf26() {
				{
					if (_2592_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L17
					}
					_2592_i11 = (_2592_i11).Plus(_dafny.One)
					var _2593_v104 _dafny.MultiSet
					_ = _2593_v104
					_2593_v104 = _dafny.MultiSetOf(p1, !(p1), (func() bool {
						if p1 {
							return p1
						}
						return p1
					})(), false, p1)
					var _2594_v105 _dafny.Set
					_ = _2594_v105
					_2594_v105 = _dafny.SetOf(!(!(p1)))
					var _rhs494 _dafny.MultiSet = (_dafny.MultiSetOf(p1)).Intersection(_dafny.MultiSetOf(p1, p1))
					_ = _rhs494
					var _rhs495 _dafny.Array = _2589_v101
					_ = _rhs495
					var _rhs496 _dafny.Set = (_2594_v105).Union(_dafny.SetOf(p1))
					_ = _rhs496
					_2593_v104 = _rhs494
					_2589_v101 = _rhs495
					_2594_v105 = _rhs496
					if false {
						(globalState).F13 = p0
						var _2595_v106 _dafny.Array
						_ = _2595_v106
						var _len0_72 _dafny.Int = _dafny.IntOfInt64(11)
						_ = _len0_72
						var _nw418 _dafny.Array
						_ = _nw418
						if _len0_72.Cmp(_dafny.Zero) == 0 {
							_nw418 = _dafny.NewArray(_len0_72)
						} else {
							var _init72 func(_dafny.Int) bool = (func(_2596_p1 bool) func(_dafny.Int) bool {
								return func(_2597_i12 _dafny.Int) bool {
									return _2596_p1
								}
							})(p1)
							_ = _init72
							var _element0_72 = _init72(_dafny.Zero)
							_ = _element0_72
							_nw418 = _dafny.NewArrayFromExample(_element0_72, nil, _len0_72)
							(_nw418).ArraySet1(_element0_72, 0)
							var _nativeLen0_72 = (_len0_72).Int()
							_ = _nativeLen0_72
							for _i0_72 := 1; _i0_72 < _nativeLen0_72; _i0_72++ {
								(_nw418).ArraySet1(_init72(_dafny.IntOf(_i0_72)), _i0_72)
							}
						}
						_2595_v106 = _nw418
						var _2598_v107 _dafny.Map
						_ = _2598_v107
						_2598_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2594_v105, p1)
						var _index422 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_2595_v106), 0))
						_ = _index422
						(_2595_v106).ArraySet1(((_2505_v59).Fm8((_2598_v107).Cardinality(), p1, (_dafny.Zero).Minus(p2), globalState)) || (p1), (_index422).Int())
						var _2599_v108 bool
						_ = _2599_v108
						_2599_v108 = true
						var _2600_v109 _dafny.Sequence
						_ = _2600_v109
						_2600_v109 = _dafny.SeqOf(p2, p0)
						var _2601_v110 D1
						_ = _2601_v110
						_2601_v110 = Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("ohan"), false)
						var _index423 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_2595_v106), 0))
						_ = _index423
						var _rhs497 _dafny.Int = (Companion_Default___.SafeDivisionInt((_this).F22(), _dafny.IntOfUint32((_2590_v102).Cardinality()))).Times(_dafny.IntOfUint32((_2600_v109).Cardinality()))
						_ = _rhs497
						var _rhs498 bool = true
						_ = _rhs498
						var _rhs499 _dafny.Sequence = Companion_Default___.Fm30(p2, (p2).Cmp((_this).F22()) >= 0, (_this).Fm4(_2590_v102, p2, _2599_v108, _2601_v110, globalState), _2599_v108, globalState)
						_ = _rhs499
						var _rhs500 bool = (p0).Cmp(p0) >= 0
						_ = _rhs500
						var _rhs501 bool = p1
						_ = _rhs501
						var _lhs332 *GlobalState = globalState
						_ = _lhs332
						var _lhs333 _dafny.Array = _2595_v106
						_ = _lhs333
						var _lhs334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_2595_v106), 0))
						_ = _lhs334
						_lhs332.F13 = _rhs497
						(_lhs333).ArraySet1(_rhs498, (_lhs334).Int())
						_2590_v102 = _rhs499
						_2599_v108 = _rhs500
						_2599_v108 = _rhs501
						var _2602_v111 D14
						_ = _2602_v111
						_2602_v111 = Companion_D14_.Create_DC35_(_2507_v62)
						var _2603_v112 D14
						_ = _2603_v112
						_2603_v112 = Companion_D14_.Create_DC37_(_2602_v111)
						var _2604_v113 D14
						_ = _2604_v113
						_2604_v113 = Companion_D14_.Create_DC37_(_2602_v111)
						var _2605_v114 D14
						_ = _2605_v114
						_2605_v114 = Companion_D14_.Create_DC37_(_2603_v112)
						var _2606_v115 _dafny.Map
						_ = _2606_v115
						_2606_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2605_v114, p2)
						var _2607_v116 T0
						_ = _2607_v116
						var _nw419 *C9 = New_C9_()
						_ = _nw419
						_nw419.Ctor__()
						_2607_v116 = _nw419
						var _2608_v117 D17
						_ = _2608_v117
						_2608_v117 = Companion_D17_.Create_DC47_(_2607_v116, _2599_v108, (_2506_v61).Cardinality(), (_2607_v116).Fm2(p2, globalState), (_this).F22())
						var _2609_v118 _dafny.Map
						_ = _2609_v118
						_2609_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2599_v108, _dafny.IntOfInt64(708))
						(globalState).F15 = (Companion_Default___.SafeDivisionInt((func() _dafny.Int {
							if (_2606_v115).Contains(_2605_v114) {
								return (_2606_v115).Get(_2605_v114).(_dafny.Int)
							}
							return p0
						})(), (_this).F22())).Minus(((_2608_v117).Dtor_cf72()).Minus((func() _dafny.Int {
							if (_2609_v118).Contains(_2599_v108) {
								return (_2609_v118).Get(_2599_v108).(_dafny.Int)
							}
							return p2
						})()))
						var _index424 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_2589_v101), 0))
						_ = _index424
						(_2589_v101).ArraySet1(p2, (_index424).Int())
						var _index425 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_2589_v101), 0))
						_ = _index425
						(_2589_v101).ArraySet1(p0, (_index425).Int())
						var _2610_v119 *C9
						_ = _2610_v119
						var _nw420 *C9 = New_C9_()
						_ = _nw420
						_nw420.Ctor__()
						_2610_v119 = _nw420
						var _2611_v120 _dafny.MultiSet
						_ = _2611_v120
						_2611_v120 = _dafny.MultiSetOf(_2610_v119, _2610_v119, _2610_v119, _2610_v119, _2610_v119)
						_2611_v120 = _dafny.MultiSetOf(_2610_v119, _2610_v119)
					} else {
						var _2612_v121 _dafny.Sequence
						_ = _2612_v121
						_2612_v121 = _dafny.SeqOf(p1)
						_2612_v121 = (func() _dafny.Sequence {
							if p1 {
								return _2612_v121
							}
							return _2612_v121
						})()
						var _2613_v122 bool
						_ = _2613_v122
						_2613_v122 = false
						var _rhs502 bool = !(_2613_v122) || (_2613_v122)
						_ = _rhs502
						var _rhs503 _dafny.MultiSet = _2593_v104
						_ = _rhs503
						_2613_v122 = _rhs502
						_2593_v104 = _rhs503
						var _2614_v123 _dafny.Sequence
						_ = _2614_v123
						_2614_v123 = _dafny.SeqOf((_this).F23())
						_2590_v102 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_2614_v123).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_2614_v123).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((_2614_v123).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_2614_v123).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _2424_v0), (Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_2614_v123).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_2614_v123).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((_2614_v123).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_2614_v123).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _2424_v0)).Cardinality()))).Uint32(), _2424_v0), _2590_v102), _dafny.UnicodeSeqOfUtf8Bytes("rxdtrqap"))
						_2589_v101 = _2589_v101
						var _index426 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_2589_v101), 0))
						_ = _index426
						(_2589_v101).ArraySet1((_dafny.Zero).Minus(p0), (_index426).Int())
						var _index427 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_2589_v101), 0))
						_ = _index427
						(_2589_v101).ArraySet1(p0, (_index427).Int())
					}
					(globalState).F13 = _dafny.IntOfUint32((Companion_Default___.Fm59(_2424_v0, p1, globalState)).Cardinality())
					var _2615_v124 *C1
					_ = _2615_v124
					var _nw421 *C1 = New_C1_()
					_ = _nw421
					_nw421.Ctor__()
					_2615_v124 = _nw421
					goto C17
				}
			C17:
			}
			goto L17
		}
	L17:
		var _2616_v125 *C1
		_ = _2616_v125
		var _nw422 *C1 = New_C1_()
		_ = _nw422
		_nw422.Ctor__()
		_2616_v125 = _nw422
	}
}
func (_this *C12) F22() _dafny.Int {
	{
		return _this._f22
	}
}
func (_this *C12) F23() _dafny.Sequence {
	{
		return _this._f23
	}
}

// End of class C12

// Definition of class C13
type C13 struct {
	F21  bool
	_f20 _dafny.Int
}

func New_C13_() *C13 {
	_this := C13{}

	_this.F21 = false
	_this._f20 = _dafny.Zero
	return &_this
}

type CompanionStruct_C13_ struct {
}

var Companion_C13_ = CompanionStruct_C13_{}

func (_this *C13) Equals(other *C13) bool {
	return _this == other
}

func (_this *C13) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C13)
	return ok && _this.Equals(other)
}

func (*C13) String() string {
	return "_module.C13"
}

func Type_C13_() _dafny.TypeDescriptor {
	return type_C13_{}
}

type type_C13_ struct {
}

func (_this type_C13_) Default() interface{} {
	return (*C13)(nil)
}

func (_this type_C13_) String() string {
	return "main.C13"
}
func (_this *C13) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C13{}

func (_this *C13) Ctor__(f20 _dafny.Int, f21 bool) {
	{
		(_this)._f20 = f20
		(_this).F21 = f21
	}
}
func (_this *C13) Fm0(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())).Merge((func() _dafny.Map {
			var _coll89 = _dafny.NewMapBuilder()
			_ = _coll89
			for _iter92 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(180), _dafny.IntOfInt64(472))); ; {
				_compr_89, _ok92 := _iter92()
				if !_ok92 {
					break
				}
				var _2617_v0 _dafny.Int
				_2617_v0 = interface{}(_compr_89).(_dafny.Int)
				if ((_dafny.IntOfInt64(180)).Cmp(_2617_v0) <= 0) && ((_2617_v0).Cmp(_dafny.IntOfInt64(472)) < 0) {
					_coll89.Add(Companion_Default___.SafeModuloInt(_2617_v0, (_this).F20()), (_this).F20())
				}
			}
			return _coll89.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F21, _this.F21)).Cardinality())))
	}
}
func (_this *C13) M0(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		if p2 {
			var _2618_v0 _dafny.Array
			_ = _2618_v0
			var _nw423 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw423
			_2618_v0 = _nw423
			var _2619_v1 D0
			_ = _2619_v1
			_2619_v1 = Companion_D0_.Create_DC0_(_2618_v0)
			var _2620_v2 _dafny.Sequence
			_ = _2620_v2
			_2620_v2 = _dafny.SeqOf(_2618_v0, _2618_v0, _2618_v0, _2618_v0)
			var _pat_let_tv57 = _2618_v0
			_ = _pat_let_tv57
			var _2621_v3 _dafny.Sequence
			_ = _2621_v3
			_2621_v3 = _dafny.SeqOf(_2618_v0, (func(_pat_let51_0 D0) D0 {
				return func(_2622_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let52_0 _dafny.Array) D0 {
						return func(_2623_dt__update_hcf0_h0 _dafny.Array) D0 {
							return Companion_D0_.Create_DC0_(_2623_dt__update_hcf0_h0)
						}(_pat_let52_0)
					}(_pat_let_tv57)
				}(_pat_let51_0)
			}(_2619_v1)).Dtor_cf0(), (_2620_v2).Select((Companion_Default___.SafeIndex((Companion_D1_.Create_DC2_((_dafny.MultiSetOf((_this).F20())).Cardinality())).Dtor_cf2(), _dafny.IntOfUint32((_2620_v2).Cardinality()))).Uint32()).(_dafny.Array), _2618_v0)
			_2621_v3 = _dafny.Companion_Sequence_.Concatenate(_2620_v2, _2621_v3)
			var _2624_v4 _dafny.CodePoint
			_ = _2624_v4
			_2624_v4 = _dafny.CodePoint('c')
			var _2625_v5 _dafny.Map
			_ = _2625_v5
			_2625_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2624_v4, _2624_v4)
			_2625_v5 = (_2625_v5).Update(_2624_v4, _2624_v4)
			var _2626_v6 *C7
			_ = _2626_v6
			var _nw424 *C7 = New_C7_()
			_ = _nw424
			_nw424.Ctor__()
			_2626_v6 = _nw424
			(_this).F21 = _this.F21
			var _2627_v7 _dafny.Sequence
			_ = _2627_v7
			_2627_v7 = _dafny.UnicodeSeqOfUtf8Bytes("dj")
			var _2628_v8 *C12
			_ = _2628_v8
			var _nw425 *C12 = New_C12_()
			_ = _nw425
			_nw425.Ctor__(_dafny.IntOfInt64(-287), _2627_v7)
			_2628_v8 = _nw425
		} else {
			(globalState).F13 = (_dafny.Zero).Minus((_this).F20())
			var _2629_v9 *C2
			_ = _2629_v9
			var _nw426 *C2 = New_C2_()
			_ = _nw426
			_nw426.Ctor__()
			_2629_v9 = _nw426
			var _2630_v10 _dafny.MultiSet
			_ = _2630_v10
			_2630_v10 = _dafny.MultiSetOf(false, false, p1, p2)
			var _2631_v11 _dafny.Set
			_ = _2631_v11
			_2631_v11 = _dafny.SetOf(p2, p2)
			var _2632_v12 _dafny.Sequence
			_ = _2632_v12
			_2632_v12 = _dafny.SeqOf(_this.F21, p1)
			_2630_v10 = Companion_Default___.Fm34(((_2631_v11).Cardinality()).Times((_this).F20()), (_dafny.MultiSetFromSeq(_2632_v12)).Cardinality(), (_this.F21) || ((_2632_v12).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_2632_v12).Cardinality()))).Uint32()).(bool)), globalState)
			var _2633_v13 _dafny.Array
			_ = _2633_v13
			var _len0_73 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_73
			var _nw427 _dafny.Array
			_ = _nw427
			if _len0_73.Cmp(_dafny.Zero) == 0 {
				_nw427 = _dafny.NewArray(_len0_73)
			} else {
				var _init73 func(_dafny.Int) _dafny.Int = func(_2634_i0 _dafny.Int) _dafny.Int {
					return (_2634_i0).Plus(_dafny.IntOfInt64(439))
				}
				_ = _init73
				var _element0_73 = _init73(_dafny.Zero)
				_ = _element0_73
				_nw427 = _dafny.NewArrayFromExample(_element0_73, nil, _len0_73)
				(_nw427).ArraySet1(_element0_73, 0)
				var _nativeLen0_73 = (_len0_73).Int()
				_ = _nativeLen0_73
				for _i0_73 := 1; _i0_73 < _nativeLen0_73; _i0_73++ {
					(_nw427).ArraySet1(_init73(_dafny.IntOf(_i0_73)), _i0_73)
				}
			}
			_2633_v13 = _nw427
			var _2635_v14 _dafny.Set
			_ = _2635_v14
			_2635_v14 = _dafny.SetOf(_2633_v13, _2633_v13, _2633_v13, _2633_v13)
			(_this).F21 = ((_2635_v14).Intersection(_2635_v14)).IsProperSubsetOf((_dafny.SetOf(_2633_v13, _2633_v13)).Intersection(_2635_v14))
			var _2636_v15 _dafny.Map
			_ = _2636_v15
			_2636_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(978)).Cmp(_dafny.IntOfUint32((_2632_v12).Cardinality())) == 0, (_2631_v11).Intersection(_2631_v11))
			_2636_v15 = _2636_v15
		}
		var _2637_v16 _dafny.Sequence
		_ = _2637_v16
		_2637_v16 = _dafny.SeqOf(p1)
		var _2638_v17 _dafny.Map
		_ = _2638_v17
		_2638_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F21, _dafny.SeqOf(false, p1, p1, p0, p0))
		var _2639_v18 _dafny.Array
		_ = _2639_v18
		var _nwElement0_106 _dafny.Sequence = _2637_v16
		_ = _nwElement0_106
		var _nw428 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_106, nil, _dafny.IntOfInt64(6))
		_ = _nw428
		(_nw428).ArraySet1(_nwElement0_106, 0)
		(_nw428).ArraySet1(_2637_v16, 1)
		(_nw428).ArraySet1((func() _dafny.Sequence {
			if (_2638_v17).Contains(_this.F21) {
				return (_2638_v17).Get(_this.F21).(_dafny.Sequence)
			}
			return _2637_v16
		})(), 2)
		(_nw428).ArraySet1(Companion_Default___.Fm27(globalState), 3)
		(_nw428).ArraySet1((Companion_D8_.Create_DC23_(_2637_v16)).Dtor_cf39(), 4)
		(_nw428).ArraySet1(_2637_v16, 5)
		_2639_v18 = _nw428
		var _2640_v19 _dafny.MultiSet
		_ = _2640_v19
		_2640_v19 = _dafny.MultiSetOf((_this).F20(), (_this).F20(), (_this).F20())
		var _2641_v20 _dafny.MultiSet
		_ = _2641_v20
		_2641_v20 = _dafny.MultiSetOf(((_2640_v19).Update((_this).F20(), Companion_Default___.Abs((_this).F20()))).Cardinality())
		var _2642_v21 _dafny.Sequence
		_ = _2642_v21
		_2642_v21 = _dafny.SeqOf(_2640_v19, _2641_v20)
		var _2643_v22 _dafny.Sequence
		_ = _2643_v22
		_2643_v22 = _dafny.UnicodeSeqOfUtf8Bytes("onoohtgc")
		var _2644_v23 _dafny.Map
		_ = _2644_v23
		_2644_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), false)
		var _2645_v24 _dafny.Array
		_ = _2645_v24
		var _nwElement0_107 bool = p2
		_ = _nwElement0_107
		var _nw429 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_107, nil, _dafny.IntOfInt64(22))
		_ = _nw429
		(_nw429).ArraySet1(_nwElement0_107, 0)
		(_nw429).ArraySet1(p1, 1)
		(_nw429).ArraySet1(Companion_Default___.Fm26(globalState), 2)
		(_nw429).ArraySet1(_dafny.Companion_Sequence_.Equal(_2642_v21, _2642_v21), 3)
		(_nw429).ArraySet1(p2, 4)
		(_nw429).ArraySet1(p1, 5)
		(_nw429).ArraySet1(_this.F21, 6)
		(_nw429).ArraySet1(!(!(p2) || (p0)), 7)
		(_nw429).ArraySet1(p2, 8)
		(_nw429).ArraySet1(p0, 9)
		(_nw429).ArraySet1(p0, 10)
		(_nw429).ArraySet1((_2637_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2643_v22).Cardinality()), _dafny.IntOfUint32((_2637_v16).Cardinality()))).Uint32()).(bool), 11)
		(_nw429).ArraySet1(Companion_Default___.Fm26(globalState), 12)
		(_nw429).ArraySet1((func() bool {
			if !(p2) {
				return p2
			}
			return (func() bool {
				if (_2644_v23).Contains((_this).F20()) {
					return (_2644_v23).Get((_this).F20()).(bool)
				}
				return p1
			})()
		})(), 13)
		(_nw429).ArraySet1(p2, 14)
		(_nw429).ArraySet1(p0, 15)
		(_nw429).ArraySet1(_this.F21, 16)
		(_nw429).ArraySet1(true, 17)
		(_nw429).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_2643_v22, _2643_v22), 18)
		(_nw429).ArraySet1(p2, 19)
		(_nw429).ArraySet1(p2, 20)
		(_nw429).ArraySet1((p2) && (_this.F21), 21)
		_2645_v24 = _nw429
		var _index428 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_2645_v24), 0))
		_ = _index428
		(_2645_v24).ArraySet1((func() bool {
			if _this.F21 {
				return p0
			}
			return _this.F21
		})(), (_index428).Int())
		var _2646_v25 _dafny.CodePoint
		_ = _2646_v25
		_2646_v25 = _dafny.CodePoint('l')
		var _pat_let_tv58 = _2643_v22
		_ = _pat_let_tv58
		var _pat_let_tv59 = _2643_v22
		_ = _pat_let_tv59
		var _index429 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_2645_v24), 0))
		_ = _index429
		var _rhs504 _dafny.Int = func(_source36 D7) _dafny.Int {
			if _source36.Is_DC20() {
				var _2647___mcc_h0 _dafny.Int = _source36.Get_().(D7_DC20).Cf29
				_ = _2647___mcc_h0
				var _2648_cf29 _dafny.Int = _2647___mcc_h0
				_ = _2648_cf29
				return _dafny.IntOfInt64(881)
			} else if _source36.Is_DC21() {
				var _2649___mcc_h1 _dafny.Int = _source36.Get_().(D7_DC21).Cf30
				_ = _2649___mcc_h1
				var _2650___mcc_h2 _dafny.Int = _source36.Get_().(D7_DC21).Cf31
				_ = _2650___mcc_h2
				var _2651___mcc_h3 bool = _source36.Get_().(D7_DC21).Cf32
				_ = _2651___mcc_h3
				var _2652___mcc_h4 bool = _source36.Get_().(D7_DC21).Cf33
				_ = _2652___mcc_h4
				var _2653___mcc_h5 bool = _source36.Get_().(D7_DC21).Cf34
				_ = _2653___mcc_h5
				var _2654_cf34 bool = _2653___mcc_h5
				_ = _2654_cf34
				var _2655_cf33 bool = _2652___mcc_h4
				_ = _2655_cf33
				var _2656_cf32 bool = _2651___mcc_h3
				_ = _2656_cf32
				var _2657_cf31 _dafny.Int = _2650___mcc_h2
				_ = _2657_cf31
				var _2658_cf30 _dafny.Int = _2649___mcc_h1
				_ = _2658_cf30
				return _dafny.IntOfInt64(-379)
			} else if _source36.Is_DC22() {
				var _2659___mcc_h6 bool = _source36.Get_().(D7_DC22).Cf35
				_ = _2659___mcc_h6
				var _2660___mcc_h7 _dafny.Int = _source36.Get_().(D7_DC22).Cf36
				_ = _2660___mcc_h7
				var _2661___mcc_h8 bool = _source36.Get_().(D7_DC22).Cf37
				_ = _2661___mcc_h8
				var _2662___mcc_h9 _dafny.Int = _source36.Get_().(D7_DC22).Cf38
				_ = _2662___mcc_h9
				var _2663_cf38 _dafny.Int = _2662___mcc_h9
				_ = _2663_cf38
				var _2664_cf37 bool = _2661___mcc_h8
				_ = _2664_cf37
				var _2665_cf36 _dafny.Int = _2660___mcc_h7
				_ = _2665_cf36
				var _2666_cf35 bool = _2659___mcc_h6
				_ = _2666_cf35
				return (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(170))).Uint32(), func(coer134 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg134 _dafny.Int) interface{} {
						return coer134(arg134)
					}
				}(func(_2667_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				})), _dafny.UnicodeSeqOfUtf8Bytes("ug"), _pat_let_tv58, (Companion_D22_.Create_DC62_(_pat_let_tv59, (_this).F20(), _2663_cf38, _2663_cf38, false)).Dtor_cf101())).Cardinality()
			} else {
				var _2668___mcc_h10 _dafny.Sequence = _source36.Get_().(D7_DC19).Cf28
				_ = _2668___mcc_h10
				var _2669_cf28 _dafny.Sequence = _2668___mcc_h10
				_ = _2669_cf28
				return (_this).F20()
			}
		}(Companion_D7_.Create_DC21_((_this).F20(), (_this).F20(), (_2637_v16).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_2637_v16).Cardinality()))).Uint32()).(bool), _this.F21, false))
		_ = _rhs504
		var _rhs505 _dafny.Array = _2639_v18
		_ = _rhs505
		var _rhs506 bool = !(_dafny.Companion_Sequence_.Contains(_2643_v22, _2646_v25))
		_ = _rhs506
		var _rhs507 _dafny.Int = (_this).F20()
		_ = _rhs507
		var _lhs335 *GlobalState = globalState
		_ = _lhs335
		var _lhs336 _dafny.Array = _2645_v24
		_ = _lhs336
		var _lhs337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_2645_v24), 0))
		_ = _lhs337
		var _lhs338 *GlobalState = globalState
		_ = _lhs338
		_lhs335.F15 = _rhs504
		_2639_v18 = _rhs505
		(_lhs336).ArraySet1(_rhs506, (_lhs337).Int())
		_lhs338.F7 = _rhs507
		var _2670_v26 _dafny.MultiSet
		_ = _2670_v26
		_2670_v26 = _dafny.MultiSetOf(_this.F21)
		var _2671_v27 _dafny.Sequence
		_ = _2671_v27
		_2671_v27 = _dafny.SeqOf((_2670_v26).Cardinality())
		var _2672_v28 _dafny.Map
		_ = _2672_v28
		_2672_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F21)
		var _2673_v29 _dafny.Map
		_ = _2673_v29
		_2673_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2671_v27).Cardinality()), (_2672_v28).Cardinality())
		var _2674_v30 D18
		_ = _2674_v30
		_2674_v30 = Companion_D18_.Create_DC48_(_2673_v29)
		var _2675_v31 _dafny.Sequence
		_ = _2675_v31
		_2675_v31 = _dafny.SeqOf(_2674_v30, _2674_v30)
		var _2676_v32 _dafny.Array
		_ = _2676_v32
		var _nw430 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
		_ = _nw430
		_2676_v32 = _nw430
		var _2677_v33 D2
		_ = _2677_v33
		_2677_v33 = Companion_D2_.Create_DC7_(p2, _dafny.MultiSetOf(true, !(p0)), _2676_v32)
		var _2678_v34 D2
		_ = _2678_v34
		_2678_v34 = Companion_D2_.Create_DC8_(_2677_v33)
		var _2679_v35 *C3
		_ = _2679_v35
		var _nw431 *C3 = New_C3_()
		_ = _nw431
		_nw431.Ctor__(_2678_v34, false)
		_2679_v35 = _nw431
		var _2680_v36 _dafny.Set
		_ = _2680_v36
		_2680_v36 = _dafny.SetOf(_2679_v35)
		var _2681_v37 _dafny.Set
		_ = _2681_v37
		_2681_v37 = _dafny.SetOf(_2679_v35.F25)
		var _2682_v38 _dafny.Map
		_ = _2682_v38
		_2682_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _2637_v16)
		var _2683_v39 _dafny.Map
		_ = _2683_v39
		_2683_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2646_v25, _2682_v38)
		var _2684_v40 _dafny.Array
		_ = _2684_v40
		var _nwElement0_108 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_2643_v22).Cardinality()))).Cardinality())
		_ = _nwElement0_108
		var _nw432 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_108, nil, _dafny.IntOfInt64(23))
		_ = _nw432
		(_nw432).ArraySet1(_nwElement0_108, 0)
		(_nw432).ArraySet1((func() _dafny.Int {
			if (_2670_v26).Contains(p1) {
				return (_2670_v26).Multiplicity(p1)
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _2675_v31)).Cardinality()
		})(), 1)
		(_nw432).ArraySet1((_this).F20(), 2)
		(_nw432).ArraySet1((_2680_v36).Cardinality(), 3)
		(_nw432).ArraySet1((_this).F20(), 4)
		(_nw432).ArraySet1((_this).F20(), 5)
		(_nw432).ArraySet1((_this).F20(), 6)
		(_nw432).ArraySet1((_2641_v20).Cardinality(), 7)
		(_nw432).ArraySet1((_this).F20(), 8)
		(_nw432).ArraySet1(_dafny.IntOfUint32((_2671_v27).Cardinality()), 9)
		(_nw432).ArraySet1((_this).F20(), 10)
		(_nw432).ArraySet1((_2681_v37).Cardinality(), 11)
		(_nw432).ArraySet1((_this).F20(), 12)
		(_nw432).ArraySet1((_this).F20(), 13)
		(_nw432).ArraySet1((_this).F20(), 14)
		(_nw432).ArraySet1((_this).F20(), 15)
		(_nw432).ArraySet1((_this).F20(), 16)
		(_nw432).ArraySet1((_2672_v28).Cardinality(), 17)
		(_nw432).ArraySet1((_this).F20(), 18)
		(_nw432).ArraySet1(((func() _dafny.Map {
			if (_2683_v39).Contains(Companion_Default___.Fm38(_2679_v35.F25, globalState)) {
				return (_2683_v39).Get(Companion_Default___.Fm38(_2679_v35.F25, globalState)).(_dafny.Map)
			}
			return _2682_v38
		})()).Cardinality(), 19)
		(_nw432).ArraySet1(_dafny.IntOfInt64(587), 20)
		(_nw432).ArraySet1((_this).F20(), 21)
		(_nw432).ArraySet1(Companion_Default___.Fm17((_this).F20(), globalState), 22)
		_2684_v40 = _nw432
		var _2685_v41 D2
		_ = _2685_v41
		_2685_v41 = Companion_D2_.Create_DC5_(_2676_v32)
		_2685_v41 = _2685_v41
		var _2686_v42 *C10
		_ = _2686_v42
		var _nw433 *C10 = New_C10_()
		_ = _nw433
		_nw433.Ctor__()
		_2686_v42 = _nw433
		var _2687_v43 _dafny.Set
		_ = _2687_v43
		_2687_v43 = _dafny.SetOf(_dafny.IntOfInt64(190), (_this).F20())
		var _2688_v44 _dafny.Array
		_ = _2688_v44
		var _nwElement0_109 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("uxr")
		_ = _nwElement0_109
		var _nw434 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_109, nil, _dafny.IntOfInt64(25))
		_ = _nw434
		(_nw434).ArraySet1(_nwElement0_109, 0)
		(_nw434).ArraySet1(_2643_v22, 1)
		(_nw434).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hlbyiwl"), 2)
		(_nw434).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2643_v22, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer135 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg135 _dafny.Int) interface{} {
				return coer135(arg135)
			}
		}((func(_2689_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2690_i2 _dafny.Int) _dafny.CodePoint {
				return _2689_v25
			}
		})(_2646_v25)))), 3)
		(_nw434).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(488))).Uint32(), func(coer136 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg136 _dafny.Int) interface{} {
				return coer136(arg136)
			}
		}((func(_2691_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2692_i3 _dafny.Int) _dafny.CodePoint {
				return _2691_v25
			}
		})(_2646_v25))), 4)
		(_nw434).ArraySet1(_2643_v22, 5)
		(_nw434).ArraySet1(_2643_v22, 6)
		(_nw434).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2643_v22, _2643_v22), 7)
		(_nw434).ArraySet1(_dafny.Companion_Sequence_.Update(_2643_v22, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_2643_v22).Cardinality()))).Uint32(), _2646_v25), 8)
		(_nw434).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2643_v22, _2643_v22), 9)
		(_nw434).ArraySet1(_2643_v22, 10)
		(_nw434).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(858))).Uint32(), func(coer137 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg137 _dafny.Int) interface{} {
				return coer137(arg137)
			}
		}(func(_2693_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		})), 11)
		(_nw434).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(757))).Uint32(), func(coer138 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg138 _dafny.Int) interface{} {
				return coer138(arg138)
			}
		}((func(_2694_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2695_i5 _dafny.Int) _dafny.CodePoint {
				return _2694_v25
			}
		})(_2646_v25))), 12)
		(_nw434).ArraySet1(_2643_v22, 13)
		(_nw434).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-511))).Uint32(), func(coer139 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg139 _dafny.Int) interface{} {
				return coer139(arg139)
			}
		}((func(_2696_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2697_i6 _dafny.Int) _dafny.CodePoint {
				return _2696_v25
			}
		})(_2646_v25))), _2643_v22), 14)
		(_nw434).ArraySet1(_2643_v22, 15)
		(_nw434).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(30))).Uint32(), func(coer140 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg140 _dafny.Int) interface{} {
				return coer140(arg140)
			}
		}((func(_2698_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2699_i7 _dafny.Int) _dafny.CodePoint {
				return _2698_v25
			}
		})(_2646_v25))), _2643_v22), 16)
		(_nw434).ArraySet1(Companion_Default___.Fm30((_this).F20(), false, _2687_v43, p0, globalState), 17)
		(_nw434).ArraySet1((func() _dafny.Sequence {
			if (_2645_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_2645_v24), 0))).Int()).(bool) {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(992))).Uint32(), func(coer141 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg141 _dafny.Int) interface{} {
						return coer141(arg141)
					}
				}((func(_2700_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2701_i8 _dafny.Int) _dafny.CodePoint {
						return _2700_v25
					}
				})(_2646_v25)))
			}
			return _2643_v22
		})(), 18)
		(_nw434).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ihj"), 19)
		(_nw434).ArraySet1(_2643_v22, 20)
		(_nw434).ArraySet1(_2643_v22, 21)
		(_nw434).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xfq"), _2643_v22), 22)
		(_nw434).ArraySet1(_2643_v22, 23)
		(_nw434).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer142 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg142 _dafny.Int) interface{} {
				return coer142(arg142)
			}
		}((func(_2702_v25 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2703_i9 _dafny.Int) _dafny.CodePoint {
				return _2702_v25
			}
		})(_2646_v25))), _2643_v22), 24)
		_2688_v44 = _nw434
		var _index430 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_2688_v44), 0))
		_ = _index430
		(_2688_v44).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("idcvyt"), _dafny.Companion_Sequence_.Update(_2643_v22, (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_2643_v22).Cardinality()))).Uint32(), Companion_Default___.Fm38(_this.F21, globalState))), (_index430).Int())
		var _index431 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_2688_v44), 0))
		_ = _index431
		(_2688_v44).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2643_v22, _2643_v22), (_index431).Int())
		var _2704_v45 _dafny.Array
		_ = _2704_v45
		var _len0_74 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_74
		var _nw435 _dafny.Array
		_ = _nw435
		if _len0_74.Cmp(_dafny.Zero) == 0 {
			_nw435 = _dafny.NewArray(_len0_74)
		} else {
			var _init74 func(_dafny.Int) _dafny.Map = (func(_2705_v23 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_2706_i11 _dafny.Int) _dafny.Map {
					return _2705_v23
				}
			})(_2644_v23)
			_ = _init74
			var _element0_74 = _init74(_dafny.Zero)
			_ = _element0_74
			_nw435 = _dafny.NewArrayFromExample(_element0_74, nil, _len0_74)
			(_nw435).ArraySet1(_element0_74, 0)
			var _nativeLen0_74 = (_len0_74).Int()
			_ = _nativeLen0_74
			for _i0_74 := 1; _i0_74 < _nativeLen0_74; _i0_74++ {
				(_nw435).ArraySet1(_init74(_dafny.IntOf(_i0_74)), _i0_74)
			}
		}
		_2704_v45 = _nw435
		for _iter93 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2704_v45), 0))); ; {
			_guard_loop_3, _ok93 := _iter93()
			if !_ok93 {
				break
			}
			var _2707_i10 _dafny.Int
			_2707_i10 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_2707_i10).Sign() != -1) && ((_2707_i10).Cmp(_dafny.ArrayLen((_2704_v45), 0)) < 0)) {
				(_2704_v45).ArraySet1(_2644_v23, (_2707_i10).Int())
			}
		}
		r0 = _2646_v25
		return r0
	}
}
func (_this *C13) F20() _dafny.Int {
	{
		return _this._f20
	}
}

// End of class C13
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
