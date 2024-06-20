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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	return (!(true)) || ((!(true)) || (false))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(299), _dafny.IntOfInt64(553)), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(162))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	}))).Cardinality())), func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(955), _dafny.IntOfInt64(500))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(955)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(500)) < 0) {
				_coll0.Add((_1_v0).Plus((_dafny.SetOf(true, !(true))).Cardinality()))
			}
		}
		return _coll0.ToSet()
	}(), func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v1 _dafny.Int
			_2_v1 = interface{}(_compr_1).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()), _2_v1) {
				_coll1.Add((_2_v1).Minus(_dafny.IntOfInt64(601)))
			}
		}
		return _coll1.ToSet()
	}())).Cardinality()), _dafny.IntOfInt64(528)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(429))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("smswnlmsn")).Cardinality())))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	if false {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true))
	} else {
		return _dafny.SeqOf(true)
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(!(true))).Intersection(_dafny.MultiSetOf(false))).Intersection((_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(!(true), false)))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(713), true))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v0 _dafny.Map
			_4_v0 = interface{}(_compr_2).(_dafny.Map)
			if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(713), true))).Contains(_4_v0) {
				_coll2.Add(_4_v0)
			}
		}
		return _coll2.ToSet()
	}()).Intersection(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), false)))).Difference(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(413), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality())), false), func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-871), _dafny.IntOfInt64(-831))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _5_v1 _dafny.Int
			_5_v1 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(-871)).Cmp(_5_v1) <= 0) && ((_5_v1).Cmp(_dafny.IntOfInt64(-831)) < 0) {
				_coll3.Add((_5_v1).Plus(_dafny.IntOfInt64(-737)), false)
			}
		}
		return _coll3.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('n')
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(255), ((func() _dafny.Set {
		if !(!(true)) {
			return _dafny.SetOf(Companion_D12_.Create_DC26_(_dafny.SetOf(_dafny.IntOfInt64(387))), Companion_D12_.Create_DC26_(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(569), true)).Cardinality())), Companion_D12_.Create_DC26_(_dafny.SetOf(_dafny.IntOfInt64(427), _dafny.IntOfInt64(687), _dafny.IntOfInt64(659), _dafny.IntOfInt64(249))), Companion_D12_.Create_DC26_(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf(true))).Cardinality()), false)).Cardinality())))
		}
		return _dafny.SetOf(Companion_D12_.Create_DC26_(func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(71), _dafny.IntOfInt64(108))); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _6_v0 _dafny.Int
				_6_v0 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(71)).Cmp(_6_v0) <= 0) && ((_6_v0).Cmp(_dafny.IntOfInt64(108)) < 0) {
					_coll4.Add((_6_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(!(true)))).Cardinality()))
				}
			}
			return _coll4.ToSet()
		}()), Companion_D12_.Create_DC26_(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cao")).Cardinality()), _dafny.IntOfInt64(353))), Companion_D12_.Create_DC26_(func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(388), _dafny.IntOfInt64(551))); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _7_v1 _dafny.Int
				_7_v1 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(388)).Cmp(_7_v1) <= 0) && ((_7_v1).Cmp(_dafny.IntOfInt64(551)) < 0) {
					_coll5.Add(Companion_Default___.SafeDivisionInt(_7_v1, (func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('e'))).Elements()); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _8_v2 _dafny.CodePoint
							_8_v2 = interface{}(_compr_6).(_dafny.CodePoint)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('e')), _8_v2) {
								_coll6.Add(_8_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('h'))).Cardinality())
							}
						}
						return _coll6.ToMap()
					}()).Cardinality()))
				}
			}
			return _coll5.ToSet()
		}()), Companion_D12_.Create_DC26_(_dafny.SetOf(_dafny.IntOfInt64(129))))
	})()).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true)).Intersection(_dafny.SetOf(true))).Union((_dafny.SetOf(false, true, !(true))).Difference(_dafny.SetOf(!(!(true)))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC6_((_dafny.MultiSetOf(!(true))).Cardinality(), _dafny.IntOfInt64(-906), false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC6_(_dafny.IntOfInt64(799), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-306))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_9_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wmk")).Cardinality())
	}))).Cardinality()), true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qi")).Cardinality()), (func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("tmyf"), !(false))).Keys().Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _10_v0 _dafny.Sequence
			_10_v0 = interface{}(_compr_7).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("tmyf"), !(false))).Contains(_10_v0) {
				_coll7.Add(_10_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ecjrmps")).Cardinality())))
			}
		}
		return _coll7.ToMap()
	}()).Cardinality(), true)))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, globalState *GlobalState) D10 {
	var _source0 D6 = Companion_D6_.Create_DC12_(true, _dafny.IntOfInt64(932))
	_ = _source0
	if _source0.Is_DC12() {
		var _11___mcc_h0 bool = _source0.Get_().(D6_DC12).Cf17
		_ = _11___mcc_h0
		var _12___mcc_h1 _dafny.Int = _source0.Get_().(D6_DC12).Cf18
		_ = _12___mcc_h1
		var _13_cf18 _dafny.Int = _12___mcc_h1
		_ = _13_cf18
		var _14_cf17 bool = _11___mcc_h0
		_ = _14_cf17
		return Companion_D10_.Create_DC18_(_dafny.SetOf(_14_cf17))
	} else {
		var _15___mcc_h2 T0 = _source0.Get_().(D6_DC11).Cf16
		_ = _15___mcc_h2
		var _16_cf16 T0 = _15___mcc_h2
		_ = _16_cf16
		return Companion_D10_.Create_DC18_(_dafny.SetOf(!(true)))
	}
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(10))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-385))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_18_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-51))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_19_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Map, globalState *GlobalState) D10 {
	return Companion_D10_.Create_DC20_((_dafny.SetOf(_dafny.IntOfInt64(266))).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfInt64(246))), (_dafny.IntOfInt64(-107)).Plus(_dafny.IntOfInt64(449)))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	var _source1 D2 = (func() D2 {
		if true {
			return Companion_D2_.Create_DC2_(false)
		}
		return Companion_D2_.Create_DC2_(true)
	})()
	_ = _source1
	if _source1.Is_DC3() {
		var _20___mcc_h0 _dafny.Int = _source1.Get_().(D2_DC3).Cf3
		_ = _20___mcc_h0
		var _21___mcc_h1 _dafny.Int = _source1.Get_().(D2_DC3).Cf4
		_ = _21___mcc_h1
		var _22_cf4 _dafny.Int = _21___mcc_h1
		_ = _22_cf4
		var _23_cf3 _dafny.Int = _20___mcc_h0
		_ = _23_cf3
		return _23_cf3
	} else {
		var _24___mcc_h2 bool = _source1.Get_().(D2_DC2).Cf2
		_ = _24___mcc_h2
		var _25_cf2 bool = _24___mcc_h2
		_ = _25_cf2
		return (_dafny.Zero).Minus((Companion_D11_.Create_DC24_(_25_cf2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_cf2, _dafny.IntOfInt64(205))).Cardinality())).Dtor_cf37())
	}
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(!(true)))).Merge((func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D11_.Create_DC23_(_dafny.IntOfInt64(355), !(true))).Dtor_cf35(), !(false))
	})())
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, globalState *GlobalState) D10 {
	var _source2 D11 = (func() D11 {
		if !(true) {
			return Companion_D11_.Create_DC24_(true, _dafny.IntOfInt64(961))
		}
		return Companion_D11_.Create_DC24_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.CodePoint('a'), _dafny.CodePoint('l'), _dafny.CodePoint('a'))).Cardinality(), _dafny.IntOfInt64(634))).Cardinality())
	})()
	_ = _source2
	if _source2.Is_DC23() {
		var _26___mcc_h0 _dafny.Int = _source2.Get_().(D11_DC23).Cf34
		_ = _26___mcc_h0
		var _27___mcc_h1 bool = _source2.Get_().(D11_DC23).Cf35
		_ = _27___mcc_h1
		var _28_cf35 bool = _27___mcc_h1
		_ = _28_cf35
		var _29_cf34 _dafny.Int = _26___mcc_h0
		_ = _29_cf34
		return Companion_D10_.Create_DC19_(_dafny.CodePoint('r'))
	} else if _source2.Is_DC24() {
		var _30___mcc_h2 bool = _source2.Get_().(D11_DC24).Cf36
		_ = _30___mcc_h2
		var _31___mcc_h3 _dafny.Int = _source2.Get_().(D11_DC24).Cf37
		_ = _31___mcc_h3
		var _32_cf37 _dafny.Int = _31___mcc_h3
		_ = _32_cf37
		var _33_cf36 bool = _30___mcc_h2
		_ = _33_cf36
		return Companion_D10_.Create_DC19_(_dafny.CodePoint('x'))
	} else if _source2.Is_DC22() {
		var _34___mcc_h4 _dafny.Array = _source2.Get_().(D11_DC22).Cf33
		_ = _34___mcc_h4
		var _35_cf33 _dafny.Array = _34___mcc_h4
		_ = _35_cf33
		return Companion_D10_.Create_DC19_(_dafny.CodePoint('k'))
	} else {
		var _36___mcc_h5 D11 = _source2.Get_().(D11_DC25).Cf38
		_ = _36___mcc_h5
		var _37_cf38 D11 = _36___mcc_h5
		_ = _37_cf38
		return Companion_D10_.Create_DC19_(_dafny.CodePoint('x'))
	}
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source3 D9 = (func() D9 {
		if true {
			return Companion_D9_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC6_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(46), _dafny.IntOfInt64(327))).Cardinality(), _dafny.IntOfInt64(723), !(false))))
		}
		return Companion_D9_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nmqjwa")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_38_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))).Cardinality()), false)))
	})()
	_ = _source3
	if _source3.Is_DC17() {
		var _39___mcc_h0 _dafny.Int = _source3.Get_().(D9_DC17).Cf25
		_ = _39___mcc_h0
		var _40_cf25 _dafny.Int = _39___mcc_h0
		_ = _40_cf25
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D7_.Create_DC13_(func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_dafny.SeqOf(_40_cf25)).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _41_v0 _dafny.Int
				_41_v0 = interface{}(_compr_8).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_40_cf25), _41_v0) {
					_coll8.Add((_41_v0).Minus(_40_cf25), false)
				}
			}
			return _coll8.ToMap()
		}())), _dafny.SeqOf(Companion_D7_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_cf25, !(false)))))
	} else {
		var _42___mcc_h1 _dafny.Map = _source3.Get_().(D9_DC16).Cf24
		_ = _42___mcc_h1
		var _43_cf24 _dafny.Map = _42___mcc_h1
		_ = _43_cf24
		return _dafny.SeqOf(Companion_D7_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-943))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_44_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		}))).Cardinality()), false)), Companion_D7_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(853), true)))
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC23_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, !(false), true), _dafny.IntOfInt64(132))).Cardinality()), ((_dafny.SetOf(_dafny.IntOfInt64(249))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality())) != 0)
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC6_((_dafny.IntOfInt64(8)).Plus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.IntOfUint32(((func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('k'), _dafny.CodePoint('x'))
		}
		return _dafny.SeqOf(_dafny.CodePoint('l'))
	})()).Cardinality()), false)
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) D9 {
	if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(_dafny.IntOfInt64(290), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(977))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_45_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))).Cardinality()), _dafny.IntOfInt64(-498), (_dafny.MultiSetOf(_dafny.IntOfInt64(752))).Cardinality())).Cardinality())).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(618))) {
		return Companion_D9_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC6_((_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfInt64(630), true)))
	} else {
		return Companion_D9_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D3_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cnempwcsy")).Cardinality()), (_dafny.SetOf(true)).Cardinality(), false)))
	}
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), (_dafny.MultiSetOf(_dafny.IntOfInt64(633))).Cardinality())).Keys().Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _46_v1 _dafny.CodePoint
				_46_v1 = interface{}(_compr_10).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), (_dafny.MultiSetOf(_dafny.IntOfInt64(633))).Cardinality())).Contains(_46_v1) {
					_coll10.Add(_46_v1)
				}
			}
			return _coll10.ToSet()
		}()).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _47_v0 _dafny.CodePoint
			_47_v0 = interface{}(_compr_9).(_dafny.CodePoint)
			if (func() _dafny.Set {
				var _coll11 = _dafny.NewBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), (_dafny.MultiSetOf(_dafny.IntOfInt64(633))).Cardinality())).Keys().Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _48_v1 _dafny.CodePoint
					_48_v1 = interface{}(_compr_11).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), (_dafny.MultiSetOf(_dafny.IntOfInt64(633))).Cardinality())).Contains(_48_v1) {
						_coll11.Add(_48_v1)
					}
				}
				return _coll11.ToSet()
			}()).Contains(_47_v0) {
				_coll9.Add(_47_v0, false)
			}
		}
		return _coll9.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), false)))
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(639), false)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(680))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_49_i0 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(!(true))).Cardinality()
	})))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf((_dafny.SetOf(false)).Cardinality()))).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false, false, false, false), _dafny.IntOfInt64(-332)))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC17_((_dafny.SetOf(false, true)).Cardinality()), false)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC17_(_dafny.IntOfInt64(288)), !(true))).Merge(func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.SetOf(Companion_D9_.Create_DC17_(_dafny.IntOfInt64(859)), Companion_D9_.Create_DC17_(_dafny.IntOfInt64(413)), Companion_D9_.Create_DC17_((func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(152))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_50_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("ujowcoxtg")
			}))).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _51_v1 _dafny.Sequence
				_51_v1 = interface{}(_compr_13).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(152))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}(func(_50_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("ujowcoxtg")
				})), _51_v1) {
					_coll13.Add(_51_v1)
				}
			}
			return _coll13.ToSet()
		}()).Cardinality()))).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _52_v0 D9
			_52_v0 = interface{}(_compr_12).(D9)
			if (_dafny.SetOf(Companion_D9_.Create_DC17_(_dafny.IntOfInt64(859)), Companion_D9_.Create_DC17_(_dafny.IntOfInt64(413)), Companion_D9_.Create_DC17_((func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(152))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}(func(_50_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("ujowcoxtg")
				}))).Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _53_v1 _dafny.Sequence
					_53_v1 = interface{}(_compr_14).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(152))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg13 _dafny.Int) interface{} {
							return coer13(arg13)
						}
					}(func(_50_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("ujowcoxtg")
					})), _53_v1) {
						_coll14.Add(_53_v1)
					}
				}
				return _coll14.ToSet()
			}()).Cardinality()))).Contains(_52_v0) {
				_coll12.Add(_52_v0, !(true))
			}
		}
		return _coll12.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if false {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_54_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		}))
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("qt")
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll15 = _dafny.NewMapBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate(((func() _dafny.Set {
			var _coll16 = _dafny.NewBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_55_i0 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.UnicodeSeqOfUtf8Bytes("ro"))
			}))).Elements()); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _56_v1 _dafny.Map
				_56_v1 = interface{}(_compr_16).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}(func(_55_i0 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.UnicodeSeqOfUtf8Bytes("ro"))
				})), _56_v1) {
					_coll16.Add(_56_v1)
				}
			}
			return _coll16.ToSet()
		}()).Intersection(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("qduxisjho"))))).Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _57_v0 _dafny.Map
			_57_v0 = interface{}(_compr_15).(_dafny.Map)
			if ((func() _dafny.Set {
				var _coll17 = _dafny.NewBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}(func(_55_i0 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.UnicodeSeqOfUtf8Bytes("ro"))
				}))).Elements()); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _58_v1 _dafny.Map
					_58_v1 = interface{}(_compr_17).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(590))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}(func(_55_i0 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.UnicodeSeqOfUtf8Bytes("ro"))
					})), _58_v1) {
						_coll17.Add(_58_v1)
					}
				}
				return _coll17.ToSet()
			}()).Intersection(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("qduxisjho"))))).Contains(_57_v0) {
				_coll15.Add(_57_v0, _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("fyriywjbp"), _dafny.UnicodeSeqOfUtf8Bytes("mniamr")))
			}
		}
		return _coll15.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-432))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(152))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus((_dafny.MultiSetOf(true, !(false))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll18 = _dafny.NewMapBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(672), _dafny.IntOfInt64(958))); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _59_v0 _dafny.Int
			_59_v0 = interface{}(_compr_18).(_dafny.Int)
			if ((_dafny.IntOfInt64(672)).Cmp(_59_v0) <= 0) && ((_59_v0).Cmp(_dafny.IntOfInt64(958)) < 0) {
				_coll18.Add(Companion_Default___.SafeModuloInt(_59_v0, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true)), _dafny.SeqOf(true, false))))
			}
		}
		return _coll18.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("gvtucu"), (func() _dafny.Map {
		var _coll19 = _dafny.NewMapBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(370), _dafny.IntOfInt64(-582))); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _60_v0 _dafny.Int
			_60_v0 = interface{}(_compr_19).(_dafny.Int)
			if ((_dafny.IntOfInt64(370)).Cmp(_60_v0) <= 0) && ((_60_v0).Cmp(_dafny.IntOfInt64(-582)) < 0) {
				_coll19.Add(Companion_Default___.SafeDivisionInt(_60_v0, _dafny.IntOfInt64(-876)), false)
			}
		}
		return _coll19.ToMap()
	}()).Cardinality())).Merge((Companion_D16_.Create_DC33_(func() _dafny.Map {
		var _coll20 = _dafny.NewMapBuilder()
		_ = _coll20
		for _iter20 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("xtjcm"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(607))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_61_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(662))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}(func(_62_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		})))).Elements()); ; {
			_compr_20, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _63_v1 _dafny.Sequence
			_63_v1 = interface{}(_compr_20).(_dafny.Sequence)
			if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("xtjcm"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(607))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_61_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(662))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}(func(_62_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})))).Contains(_63_v1) {
				_coll20.Add(_63_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(656))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}(func(_64_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(473)
				}))).Cardinality()), true)).Cardinality())
			}
		}
		return _coll20.ToMap()
	}())).Dtor_cf49())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("o"), _dafny.IntOfInt64(612)))
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(-661)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qhld")).Cardinality()), (_dafny.Zero).Minus((_dafny.SetOf(false, true)).Cardinality())))).Intersection(_dafny.SetOf(_dafny.SeqOf((func() _dafny.Set {
		var _coll21 = _dafny.NewBuilder()
		_ = _coll21
		for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(918), _dafny.IntOfInt64(553))); ; {
			_compr_21, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _65_v0 _dafny.Int
			_65_v0 = interface{}(_compr_21).(_dafny.Int)
			if ((_dafny.IntOfInt64(918)).Cmp(_65_v0) <= 0) && ((_65_v0).Cmp(_dafny.IntOfInt64(553)) < 0) {
				_coll21.Add((_65_v0).Plus(_dafny.IntOfInt64(-644)))
			}
		}
		return _coll21.ToSet()
	}()).Cardinality())))).Union((func() _dafny.Set {
		var _coll22 = _dafny.NewBuilder()
		_ = _coll22
		for _iter22 := _dafny.Iterate((func() _dafny.Set {
			var _coll23 = _dafny.NewBuilder()
			_ = _coll23
			for _iter23 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(601)))).Elements()); ; {
				_compr_23, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _66_v1 _dafny.Sequence
				_66_v1 = interface{}(_compr_23).(_dafny.Sequence)
				if (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(601)))).Contains(_66_v1) {
					_coll23.Add(_66_v1)
				}
			}
			return _coll23.ToSet()
		}()).Elements()); ; {
			_compr_22, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _67_v2 _dafny.Sequence
			_67_v2 = interface{}(_compr_22).(_dafny.Sequence)
			if (func() _dafny.Set {
				var _coll24 = _dafny.NewBuilder()
				_ = _coll24
				for _iter24 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(601)))).Elements()); ; {
					_compr_24, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _68_v1 _dafny.Sequence
					_68_v1 = interface{}(_compr_24).(_dafny.Sequence)
					if (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(601)))).Contains(_68_v1) {
						_coll24.Add(_68_v1)
					}
				}
				return _coll24.ToSet()
			}()).Contains(_67_v2) {
				_coll22.Add(_67_v2)
			}
		}
		return _coll22.ToSet()
	}()).Difference(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(894)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(-476), _dafny.IntOfInt64(198)), _dafny.SeqOf(_dafny.IntOfInt64(439)))))
}
func (_static *CompanionStruct_Default___) M6(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	var r0 _dafny.CodePoint = _dafny.CodePoint('D')
	_ = r0
	var _hi0 _dafny.Int = (_dafny.Zero).Minus(p0)
	_ = _hi0
	for _69_i0 := p0; _69_i0.Cmp(_hi0) < 0; _69_i0 = _69_i0.Plus(_dafny.One) {
		if p1 {
			var _70_v0 _dafny.Array
			_ = _70_v0
			var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw0
			_70_v0 = _nw0
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_70_v0), 0))
			_ = _index0
			(_70_v0).ArraySet1(p2, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_70_v0), 0))
			_ = _index1
			(_70_v0).ArraySet1(p3, (_index1).Int())
			var _71_v1 _dafny.CodePoint
			_ = _71_v1
			_71_v1 = _dafny.CodePoint('f')
			var _72_v2 *C3
			_ = _72_v2
			var _nw1 *C3 = New_C3_()
			_ = _nw1
			_nw1.Ctor__(Companion_Default___.Fm0(_69_i0, (_70_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_70_v0), 0))).Int()).(bool), p3, p0, globalState), _71_v1, p0)
			_72_v2 = _nw1
			_72_v2 = _72_v2
			var _73_v3 _dafny.MultiSet
			_ = _73_v3
			_73_v3 = _dafny.MultiSetOf(Companion_Default___.Fm0(p0, (_70_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_70_v0), 0))).Int()).(bool), _72_v2.F31, _dafny.IntOfInt64(-620), globalState), p3, Companion_Default___.Fm0(p0, (_70_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_70_v0), 0))).Int()).(bool), _72_v2.F31, _69_i0, globalState), p1)
			var _74_v4 _dafny.Map
			_ = _74_v4
			_74_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v3, _69_i0)
			var _75_v5 _dafny.Sequence
			_ = _75_v5
			_75_v5 = _dafny.UnicodeSeqOfUtf8Bytes("vah")
			var _76_v6 *C4
			_ = _76_v6
			var _nw2 *C4 = New_C4_()
			_ = _nw2
			_nw2.Ctor__(_74_v4, _dafny.Companion_Sequence_.IsProperPrefixOf(_75_v5, _75_v5), _71_v1, p0)
			_76_v6 = _nw2
			var _77_v7 T1
			_ = _77_v7
			var _nw3 *C3 = New_C3_()
			_ = _nw3
			_nw3.Ctor__(p2, _71_v1, p0)
			_77_v7 = _nw3
			var _78_v8 _dafny.Sequence
			_ = _78_v8
			_78_v8 = _dafny.SeqOf(_77_v7, _77_v7)
			var _79_v9 *C5
			_ = _79_v9
			var _nw4 *C5 = New_C5_()
			_ = _nw4
			_nw4.Ctor__((_76_v6).F30(), _71_v1, (_69_i0).Times(_dafny.IntOfUint32((_78_v8).Cardinality())))
			_79_v9 = _nw4
			_79_v9 = _79_v9
		} else {
			r0 = Companion_Default___.Fm10((_69_i0).Cmp(_69_i0) >= 0, p2, p1, p3, globalState)
			var _80_v10 _dafny.Sequence
			_ = _80_v10
			_80_v10 = _dafny.UnicodeSeqOfUtf8Bytes("rthyuxu")
			var _81_v11 _dafny.MultiSet
			_ = _81_v11
			_81_v11 = _dafny.MultiSetOf(Companion_Default___.Fm5(_dafny.IntOfUint32((_80_v10).Cardinality()), _69_i0, globalState))
			var _82_v12 _dafny.Array
			_ = _82_v12
			var _nwElement0_0 bool = p2
			_ = _nwElement0_0
			var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(8))
			_ = _nw5
			(_nw5).ArraySet1(_nwElement0_0, 0)
			(_nw5).ArraySet1((func() bool {
				if p2 {
					return p3
				}
				return p3
			})(), 1)
			(_nw5).ArraySet1((_81_v11).IsSubsetOf(_81_v11), 2)
			(_nw5).ArraySet1(p2, 3)
			(_nw5).ArraySet1(true, 4)
			(_nw5).ArraySet1(p1, 5)
			(_nw5).ArraySet1((p0).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_69_i0))) != 0, 6)
			(_nw5).ArraySet1((p3) || (p3), 7)
			_82_v12 = _nw5
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_82_v12), 0))
			_ = _index2
			(_82_v12).ArraySet1(p2, (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_82_v12), 0))
			_ = _index3
			(_82_v12).ArraySet1(p3, (_index3).Int())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_82_v12), 0))
			_ = _index4
			(_82_v12).ArraySet1(p2, (_index4).Int())
			var _83_v13 _dafny.Map
			_ = _83_v13
			_83_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-334), p1)
			_83_v13 = (_83_v13).Update(_69_i0, (_82_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_82_v12), 0))).Int()).(bool))
			var _84_v14 _dafny.Array
			_ = _84_v14
			var _len0_0 _dafny.Int = _dafny.One
			_ = _len0_0
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Sequence = (func(_85_v12 _dafny.Array, _86_p1 bool, _87_p2 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_88_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(!((_85_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_85_v12), 0))).Int()).(bool)), _86_p1, _87_p2, !(_86_p1))
					}
				})(_82_v12, p1, p2)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw6 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw6).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw6).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_84_v14 = _nw6
			var _89_v15 _dafny.Sequence
			_ = _89_v15
			_89_v15 = _dafny.SeqOf(p3)
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_84_v14), 0))
			_ = _index5
			(_84_v14).ArraySet1(_89_v15, (_index5).Int())
			var _90_v16 _dafny.Set
			_ = _90_v16
			_90_v16 = _dafny.SetOf(_83_v13)
			var _91_v17 D3
			_ = _91_v17
			_91_v17 = Companion_D3_.Create_DC4_(_90_v16)
			var _92_v18 *C0
			_ = _92_v18
			var _nw7 *C0 = New_C0_()
			_ = _nw7
			_nw7.Ctor__()
			_92_v18 = _nw7
			var _93_v19 _dafny.CodePoint
			_ = _93_v19
			_93_v19 = _dafny.CodePoint('w')
			var _94_v21 _dafny.Set
			_ = _94_v21
			_94_v21 = _dafny.SetOf(_93_v19)
			var _95_v22 D10
			_ = _95_v22
			_95_v22 = Companion_D10_.Create_DC21_(_dafny.IntOfInt64(787), (_69_i0).Cmp(_dafny.IntOfInt64(296)) == 0, !(func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter25 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v19, _dafny.IntOfInt64(736))).Keys().Elements()); ; {
					_compr_25, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _96_v20 _dafny.CodePoint
					_96_v20 = interface{}(_compr_25).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v19, _dafny.IntOfInt64(736))).Contains(_96_v20) {
						_coll25.Add(_96_v20)
					}
				}
				return _coll25.ToSet()
			}()).Equals(_94_v21))
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_84_v14), 0))
			_ = _index6
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_82_v12), 0))
			_ = _index7
			var _rhs0 _dafny.Sequence = _89_v15
			_ = _rhs0
			var _rhs1 bool = ((p0).Cmp(_69_i0) < 0) || ((p0).Cmp(_dafny.IntOfUint32((_80_v10).Cardinality())) >= 0)
			_ = _rhs1
			var _rhs2 D3 = _91_v17
			_ = _rhs2
			var _rhs3 *C0 = _92_v18
			_ = _rhs3
			var _rhs4 D10 = _95_v22
			_ = _rhs4
			var _lhs0 _dafny.Array = _84_v14
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_84_v14), 0))
			_ = _lhs1
			var _lhs2 _dafny.Array = _82_v12
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_82_v12), 0))
			_ = _lhs3
			(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
			(_lhs2).ArraySet1(_rhs1, (_lhs3).Int())
			_91_v17 = _rhs2
			_92_v18 = _rhs3
			_95_v22 = _rhs4
		}
		var _97_v23 _dafny.Sequence
		_ = _97_v23
		_97_v23 = _dafny.SeqOf(!(p1), p3)
		var _98_v24 _dafny.Array
		_ = _98_v24
		var _nwElement0_1 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_97_v23, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_97_v23).Cardinality()))).Uint32(), p1), _97_v23)
		_ = _nwElement0_1
		var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(17))
		_ = _nw8
		(_nw8).ArraySet1(_nwElement0_1, 0)
		(_nw8).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_97_v23, _97_v23), 1)
		(_nw8).ArraySet1(_97_v23, 2)
		(_nw8).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_97_v23, _97_v23), 3)
		(_nw8).ArraySet1((func() _dafny.Sequence {
			if p2 {
				return _97_v23
			}
			return _97_v23
		})(), 4)
		(_nw8).ArraySet1(_97_v23, 5)
		(_nw8).ArraySet1(_97_v23, 6)
		(_nw8).ArraySet1(_97_v23, 7)
		(_nw8).ArraySet1(_97_v23, 8)
		(_nw8).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_97_v23, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_97_v23).Cardinality()))).Uint32(), !(p1)), _97_v23), 9)
		(_nw8).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3), _97_v23), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3), _97_v23)).Cardinality()))).Uint32(), p1), 10)
		(_nw8).ArraySet1(_dafny.Companion_Sequence_.Update(_97_v23, (Companion_Default___.SafeIndex(_69_i0, _dafny.IntOfUint32((_97_v23).Cardinality()))).Uint32(), p1), 11)
		(_nw8).ArraySet1(_97_v23, 12)
		(_nw8).ArraySet1(_97_v23, 13)
		(_nw8).ArraySet1(_97_v23, 14)
		(_nw8).ArraySet1(_97_v23, 15)
		(_nw8).ArraySet1(_97_v23, 16)
		_98_v24 = _nw8
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_98_v24), 0))
		_ = _index8
		(_98_v24).ArraySet1(_97_v23, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_98_v24), 0))
		_ = _index9
		(_98_v24).ArraySet1(Companion_Default___.Fm7(globalState), (_index9).Int())
		var _99_v25 D11
		_ = _99_v25
		_99_v25 = Companion_D11_.Create_DC23_(p0, p1)
		var _100_v26 D11
		_ = _100_v26
		_100_v26 = Companion_D11_.Create_DC25_(_99_v25)
		var _101_v27 _dafny.Map
		_ = _101_v27
		_101_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v26, p1)
		var _rhs5 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm17((_101_v27).Equals(_101_v27), !(p1) || (p1), p3, globalState))
		_ = _rhs5
		var _rhs6 bool = p3
		_ = _rhs6
		var _lhs4 *GlobalState = globalState
		_ = _lhs4
		var _lhs5 *GlobalState = globalState
		_ = _lhs5
		_lhs4.F2 = _rhs5
		_lhs5.F18 = _rhs6
		var _102_v28 _dafny.Sequence
		_ = _102_v28
		_102_v28 = _dafny.UnicodeSeqOfUtf8Bytes("rkvodkl")
		if !_dafny.Companion_Sequence_.Equal(_102_v28, _102_v28) {
			var _103_v29 *C2
			_ = _103_v29
			var _nw9 *C2 = New_C2_()
			_ = _nw9
			_nw9.Ctor__((_69_i0).Plus(_69_i0), !(p2) || (p2))
			_103_v29 = _nw9
			(globalState).F18 = true
			var _104_v31 _dafny.CodePoint
			_ = _104_v31
			_104_v31 = _dafny.CodePoint('l')
			var _105_v32 *C1
			_ = _105_v32
			var _nw10 *C1 = New_C1_()
			_ = _nw10
			_nw10.Ctor__(_104_v31, p0)
			_105_v32 = _nw10
			var _106_v33 _dafny.Map
			_ = _106_v33
			_106_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(330), _105_v32)
			var _107_v34 _dafny.Map
			_ = _107_v34
			_107_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_103_v29).F28())
			var _108_v35 *C6
			_ = _108_v35
			var _nw11 *C6 = New_C6_()
			_ = _nw11
			_nw11.Ctor__(((func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(606), _dafny.IntOfInt64(984))); ; {
					_compr_26, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _109_v30 _dafny.Int
					_109_v30 = interface{}(_compr_26).(_dafny.Int)
					if ((_dafny.IntOfInt64(606)).Cmp(_109_v30) <= 0) && ((_109_v30).Cmp(_dafny.IntOfInt64(984)) < 0) {
						_coll26.Add((_109_v30).Times((_103_v29).F27()), !(false))
					}
				}
				return _coll26.ToMap()
			}()).Update(p0, (_97_v23).Select((Companion_Default___.SafeIndex((_106_v33).Cardinality(), _dafny.IntOfUint32((_97_v23).Cardinality()))).Uint32()).(bool))).Merge(_107_v34))
			_108_v35 = _nw11
			(globalState).F3 = (_dafny.UnicodeSeqOfUtf8Bytes("rohb"))
			(globalState).F16 = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt((_103_v29).F27(), _69_i0)).Minus((func() _dafny.Int {
				if p3 {
					return p0
				}
				return (_dafny.Zero).Minus(p0)
			})()))
		} else {
			var _110_v36 _dafny.MultiSet
			_ = _110_v36
			_110_v36 = _dafny.MultiSetOf(p3)
			var _111_v37 _dafny.Map
			_ = _111_v37
			_111_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _110_v36)
			(globalState).F18 = ((func() _dafny.MultiSet {
				if (_111_v37).Contains(_dafny.IntOfInt64(-178)) {
					return (_111_v37).Get(_dafny.IntOfInt64(-178)).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(false)
			})()).IsDisjointFrom(_110_v36)
			var _112_v38 _dafny.Sequence
			_ = _112_v38
			_112_v38 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate((_98_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_98_v24), 0))).Int()).(_dafny.Sequence), _97_v23), _dafny.SeqOf(!(p2), p1, p3, p3, p2), _dafny.SeqOf(Companion_Default___.Fm0(_dafny.IntOfUint32((_102_v28).Cardinality()), p1, p3, _69_i0, globalState)), _97_v23, _97_v23)
			_112_v38 = _112_v38
			var _113_v39 _dafny.Set
			_ = _113_v39
			_113_v39 = _dafny.SetOf(_102_v28)
			var _114_v40 _dafny.Map
			_ = _114_v40
			_114_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_113_v39).Cardinality(), _69_i0)
			var _115_v41 _dafny.Sequence
			_ = _115_v41
			_115_v41 = _dafny.SeqOf(_114_v40)
			(globalState).F18 = Companion_Default___.Fm0(_69_i0, true, p1, (p0).Times(((_115_v41).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_115_v41).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()), globalState)
			(globalState).F10 = (p0).Plus(_69_i0)
			(globalState).F0 = (_69_i0).Minus((_dafny.Zero).Minus(p0))
		}
	}
	var _116_v42 _dafny.Sequence
	_ = _116_v42
	_116_v42 = _dafny.SeqOf(p1, p3)
	var _117_v43 *C5
	_ = _117_v43
	var _nw12 *C5 = New_C5_()
	_ = _nw12
	_nw12.Ctor__(p3, Companion_Default___.Fm10(true, (_116_v42).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_116_v42).Cardinality()))).Uint32()).(bool), p1, false, globalState), Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(12)))
	_117_v43 = _nw12
	if p2 {
		(globalState).F10 = (_dafny.Zero).Minus(p0)
		var _118_v44 _dafny.CodePoint
		_ = _118_v44
		_118_v44 = _dafny.CodePoint('r')
		var _119_v45 *C3
		_ = _119_v45
		var _nw13 *C3 = New_C3_()
		_ = _nw13
		_nw13.Ctor__(true, _118_v44, (p0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qicyr")).Cardinality())))
		_119_v45 = _nw13
		var _120_v46 _dafny.MultiSet
		_ = _120_v46
		_120_v46 = _dafny.MultiSetOf(p1)
		var _121_v47 _dafny.Map
		_ = _121_v47
		_121_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v46, p0)
		var _122_v48 T0
		_ = _122_v48
		var _nw14 *C4 = New_C4_()
		_ = _nw14
		_nw14.Ctor__((_121_v47).Merge(_121_v47), (_117_v43).F26(), _118_v44, p0)
		_122_v48 = _nw14
		_122_v48 = _122_v48
		var _123_v49 _dafny.Set
		_ = _123_v49
		_123_v49 = _dafny.SetOf((_117_v43).F26())
		var _124_v50 _dafny.Sequence
		_ = _124_v50
		_124_v50 = _dafny.SeqOf(_dafny.IntOfInt64(754), _122_v48.F25())
		var _125_v51 *C4
		_ = _125_v51
		var _nw15 *C4 = New_C4_()
		_ = _nw15
		_nw15.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v46, _122_v48.F25()), (p0).Cmp((_123_v49).Cardinality()) < 0, Companion_Default___.Fm10(!(p3), !(p1), (_119_v45).Fm18(globalState), (_117_v43).F26(), globalState), _dafny.IntOfUint32((_124_v50).Cardinality()))
		_125_v51 = _nw15
		var _126_v52 _dafny.Array
		_ = _126_v52
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_1
		var _nw16 _dafny.Array
		_ = _nw16
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw16 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_127_v51 *C4) func(_dafny.Int) bool {
				return func(_128_i2 _dafny.Int) bool {
					return (_127_v51).F30()
				}
			})(_125_v51)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw16 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw16).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw16).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_126_v52 = _nw16
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_126_v52), 0))
		_ = _index10
		(_126_v52).ArraySet1((Companion_Default___.Fm17(p1, p3, !((_117_v43).F26()), globalState)).Cmp(Companion_Default___.Fm17((_117_v43).F26(), p3, _119_v45.F31, globalState)) >= 0, (_index10).Int())
		var _129_v53 _dafny.Array
		_ = _129_v53
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_2
		var _nw17 _dafny.Array
		_ = _nw17
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw17 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = func(_130_i3 _dafny.Int) _dafny.Int {
				return (_130_i3).Plus(_dafny.IntOfInt64(928))
			}
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw17 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw17).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw17).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_129_v53 = _nw17
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_129_v53), 0))
		_ = _index11
		(_129_v53).ArraySet1(p0, (_index11).Int())
		var _131_v54 _dafny.Sequence
		_ = _131_v54
		_131_v54 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_124_v50, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_124_v50).Cardinality()))).Uint32(), _dafny.IntOfInt64(86)), _124_v50, _dafny.SeqOf(p0), _124_v50, _124_v50)
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_126_v52), 0))
		_ = _index12
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_129_v53), 0))
		_ = _index13
		var _rhs7 _dafny.Int = _122_v48.F25()
		_ = _rhs7
		var _rhs8 *C4 = _125_v51
		_ = _rhs8
		var _rhs9 bool = (_dafny.SetOf(p1)).Contains(p1)
		_ = _rhs9
		var _rhs10 _dafny.Int = _dafny.IntOfInt64(-14)
		_ = _rhs10
		var _rhs11 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_122_v48.F25(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_131_v54, _131_v54)).Cardinality())))
		_ = _rhs11
		var _lhs6 *GlobalState = globalState
		_ = _lhs6
		var _lhs7 _dafny.Array = _126_v52
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_126_v52), 0))
		_ = _lhs8
		var _lhs9 *GlobalState = globalState
		_ = _lhs9
		var _lhs10 _dafny.Array = _129_v53
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_129_v53), 0))
		_ = _lhs11
		_lhs6.F10 = _rhs7
		_125_v51 = _rhs8
		(_lhs7).ArraySet1(_rhs9, (_lhs8).Int())
		_lhs9.F11 = _rhs10
		(_lhs10).ArraySet1(_rhs11, (_lhs11).Int())
		var _132_v55 _dafny.Map
		_ = _132_v55
		_132_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_129_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_129_v53), 0))).Int()).(_dafny.Int))
		var _133_v56 _dafny.Array
		_ = _133_v56
		var _nwElement0_2 _dafny.Map = (func() _dafny.Map {
			if (_125_v51).F30() {
				return _132_v55
			}
			return _132_v55
		})()
		_ = _nwElement0_2
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_2, 0)
		(_nw18).ArraySet1(_132_v55, 1)
		_133_v56 = _nw18
		var _134_v57 _dafny.Map
		_ = _134_v57
		_134_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_116_v42).Cardinality())), !((_117_v43).F26()))
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_133_v56), 0))
		_ = _index14
		(_133_v56).ArraySet1(Companion_Default___.Fm31((_134_v57).Cardinality(), (_117_v43).F26(), Companion_Default___.Fm0(_dafny.IntOfInt64(967), p3, p3, _dafny.IntOfUint32((_116_v42).Cardinality()), globalState), _dafny.IntOfInt64(262), globalState), (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_133_v56), 0))
		_ = _index15
		(_133_v56).ArraySet1(_132_v55, (_index15).Int())
	} else {
		var _135_v58 _dafny.Array
		_ = _135_v58
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_3
		var _nw19 _dafny.Array
		_ = _nw19
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw19 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Int = (func(_136_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_137_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_137_i4, _136_p0)
				}
			})(p0)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw19 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw19).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw19).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_135_v58 = _nw19
		var _138_v59 _dafny.Sequence
		_ = _138_v59
		_138_v59 = _dafny.UnicodeSeqOfUtf8Bytes("qdj")
		var _139_v60 _dafny.Sequence
		_ = _139_v60
		_139_v60 = _dafny.SeqOf(_138_v59)
		var _140_v61 _dafny.CodePoint
		_ = _140_v61
		_140_v61 = _dafny.CodePoint('q')
		var _141_v62 _dafny.Map
		_ = _141_v62
		_141_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _140_v61)
		var _142_v63 _dafny.Array
		_ = _142_v63
		var _nwElement0_3 bool = p2
		_ = _nwElement0_3
		var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(13))
		_ = _nw20
		(_nw20).ArraySet1(_nwElement0_3, 0)
		(_nw20).ArraySet1(true, 1)
		(_nw20).ArraySet1(p2, 2)
		(_nw20).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_139_v60, _139_v60), 3)
		(_nw20).ArraySet1(p3, 4)
		(_nw20).ArraySet1(true, 5)
		(_nw20).ArraySet1(p2, 6)
		(_nw20).ArraySet1((_117_v43).F26(), 7)
		(_nw20).ArraySet1(!_dafny.Companion_Sequence_.Contains(_138_v59, _140_v61), 8)
		(_nw20).ArraySet1(p3, 9)
		(_nw20).ArraySet1(!_dafny.Companion_Sequence_.Contains(_116_v42, p2), 10)
		(_nw20).ArraySet1(!(_141_v62).Equals(_141_v62), 11)
		(_nw20).ArraySet1((p0).Cmp(_dafny.IntOfInt64(-372)) <= 0, 12)
		_142_v63 = _nw20
		var _rhs12 bool = !((_117_v43).F26()) || (p3)
		_ = _rhs12
		var _rhs13 _dafny.Array = _135_v58
		_ = _rhs13
		var _rhs14 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(34))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_143_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))).Cardinality())
		_ = _rhs14
		var _rhs15 _dafny.Int = (p0).Plus(p0)
		_ = _rhs15
		var _rhs16 _dafny.Array = _142_v63
		_ = _rhs16
		var _lhs12 *GlobalState = globalState
		_ = _lhs12
		var _lhs13 *GlobalState = globalState
		_ = _lhs13
		var _lhs14 *GlobalState = globalState
		_ = _lhs14
		_lhs12.F18 = _rhs12
		_135_v58 = _rhs13
		_lhs13.F16 = _rhs14
		_lhs14.F11 = _rhs15
		_142_v63 = _rhs16
		var _144_v64 D13
		_ = _144_v64
		_144_v64 = Companion_D13_.Create_DC29_((_117_v43).F26(), (_117_v43).F26(), _142_v63)
		var _145_v65 _dafny.Map
		_ = _145_v65
		_145_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D13_.Create_DC29_(p2, Companion_Default___.Fm0(p0, p3, p2, p0, globalState), _142_v63), _135_v58)
		var _rhs17 bool = p1
		_ = _rhs17
		var _rhs18 bool = (_145_v65).Contains(_144_v64)
		_ = _rhs18
		var _rhs19 _dafny.Int = p0
		_ = _rhs19
		var _rhs20 _dafny.Int = (func() _dafny.Int {
			if !(p1) {
				return p0
			}
			return Companion_Default___.Fm5(p0, p0, globalState)
		})()
		_ = _rhs20
		var _lhs15 *GlobalState = globalState
		_ = _lhs15
		var _lhs16 *GlobalState = globalState
		_ = _lhs16
		var _lhs17 *GlobalState = globalState
		_ = _lhs17
		var _lhs18 *GlobalState = globalState
		_ = _lhs18
		_lhs15.F18 = _rhs17
		_lhs16.F18 = _rhs18
		_lhs17.F16 = _rhs19
		_lhs18.F22 = _rhs20
		var _146_v66 *C2
		_ = _146_v66
		var _nw21 *C2 = New_C2_()
		_ = _nw21
		_nw21.Ctor__(p0, (_117_v43).F26())
		_146_v66 = _nw21
		var _147_v67 _dafny.Set
		_ = _147_v67
		_147_v67 = _dafny.SetOf(_146_v66, _146_v66, _146_v66, _146_v66, _146_v66)
		var _148_v68 _dafny.Sequence
		_ = _148_v68
		_148_v68 = _dafny.SeqOf(_147_v67, _147_v67, _147_v67, _147_v67, _147_v67)
		_148_v68 = _dafny.Companion_Sequence_.Concatenate(_148_v68, _148_v68)
		var _149_v69 _dafny.Map
		_ = _149_v69
		_149_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_146_v66).F28(), (_117_v43).F26())
		_149_v69 = (_149_v69).Update(p1, p1)
		var _150_v70 *C3
		_ = _150_v70
		var _nw22 *C3 = New_C3_()
		_ = _nw22
		_nw22.Ctor__(p1, _140_v61, p0)
		_150_v70 = _nw22
		var _151_v71 _dafny.Sequence
		_ = _151_v71
		_151_v71 = _dafny.SeqOf(_150_v70, _150_v70, _150_v70, _150_v70)
		var _152_v72 _dafny.Map
		_ = _152_v72
		_152_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_117_v43).F26(), _151_v71)
		var _153_v73 _dafny.Sequence
		_ = _153_v73
		_153_v73 = _dafny.SeqOf(_151_v71, _151_v71, _151_v71, _151_v71)
		var _154_v74 _dafny.Map
		_ = _154_v74
		_154_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_146_v66).F27(), _151_v71)
		var _155_v75 T0
		_ = _155_v75
		var _nw23 *C1 = New_C1_()
		_ = _nw23
		_nw23.Ctor__(_dafny.CodePoint('g'), p0)
		_155_v75 = _nw23
		var _156_v76 _dafny.Map
		_ = _156_v76
		_156_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v75, _151_v71)
		var _157_v77 _dafny.Array
		_ = _157_v77
		var _nwElement0_4 _dafny.Sequence = (func() _dafny.Sequence {
			if p2 {
				return _151_v71
			}
			return _151_v71
		})()
		_ = _nwElement0_4
		var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(21))
		_ = _nw24
		(_nw24).ArraySet1(_nwElement0_4, 0)
		(_nw24).ArraySet1(_151_v71, 1)
		(_nw24).ArraySet1(_151_v71, 2)
		(_nw24).ArraySet1(_151_v71, 3)
		(_nw24).ArraySet1(_151_v71, 4)
		(_nw24).ArraySet1(_dafny.SeqOf(_150_v70), 5)
		(_nw24).ArraySet1(_151_v71, 6)
		(_nw24).ArraySet1(_151_v71, 7)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_151_v71, _dafny.Companion_Sequence_.Update(_151_v71, (Companion_Default___.SafeIndex((_146_v66).F27(), _dafny.IntOfUint32((_151_v71).Cardinality()))).Uint32(), _150_v70)), 8)
		(_nw24).ArraySet1(_151_v71, 9)
		(_nw24).ArraySet1(_151_v71, 10)
		(_nw24).ArraySet1(_dafny.SeqOf(_150_v70, _150_v70), 11)
		(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_152_v72).Contains(!(p2)) {
				return (_152_v72).Get(!(p2)).(_dafny.Sequence)
			}
			return _151_v71
		})(), _dafny.Companion_Sequence_.Update((_153_v73).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_153_v73).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_153_v73).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_153_v73).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _150_v70)), 12)
		(_nw24).ArraySet1(_151_v71, 13)
		(_nw24).ArraySet1((func() _dafny.Sequence {
			if (_154_v74).Contains((_146_v66).F27()) {
				return (_154_v74).Get((_146_v66).F27()).(_dafny.Sequence)
			}
			return _151_v71
		})(), 14)
		(_nw24).ArraySet1(_151_v71, 15)
		(_nw24).ArraySet1(_dafny.SeqOf(_150_v70), 16)
		(_nw24).ArraySet1(_151_v71, 17)
		(_nw24).ArraySet1(_151_v71, 18)
		(_nw24).ArraySet1((func() _dafny.Sequence {
			if (_156_v76).Contains(_155_v75) {
				return (_156_v76).Get(_155_v75).(_dafny.Sequence)
			}
			return _151_v71
		})(), 19)
		(_nw24).ArraySet1(_151_v71, 20)
		_157_v77 = _nw24
		_157_v77 = _157_v77
	}
	var _158_v78 _dafny.Array
	_ = _158_v78
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(22)
	_ = _len0_4
	var _nw25 _dafny.Array
	_ = _nw25
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw25 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) bool = (func(_159_v43 *C5) func(_dafny.Int) bool {
			return func(_160_i7 _dafny.Int) bool {
				return (_159_v43).F26()
			}
		})(_117_v43)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw25 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw25).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw25).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_158_v78 = _nw25
	var _source4 D13 = Companion_D13_.Create_DC29_(p2, !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(995))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}((func(_161_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_162_i6 _dafny.Int) _dafny.Int {
			return _161_p0
		}
	})(p0))), _dafny.SeqOf(_dafny.IntOfInt64(7), p0)), _158_v78)
	_ = _source4
	if _source4.Is_DC29() {
		var _163___mcc_h0 bool = _source4.Get_().(D13_DC29).Cf43
		_ = _163___mcc_h0
		var _164___mcc_h1 bool = _source4.Get_().(D13_DC29).Cf44
		_ = _164___mcc_h1
		var _165___mcc_h2 _dafny.Array = _source4.Get_().(D13_DC29).Cf45
		_ = _165___mcc_h2
		var _166_cf45 _dafny.Array = _165___mcc_h2
		_ = _166_cf45
		var _167_cf44 bool = _164___mcc_h1
		_ = _167_cf44
		var _168_cf43 bool = _163___mcc_h0
		_ = _168_cf43
		var _169_v79 _dafny.Map
		_ = _169_v79
		_169_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
		var _170_v80 _dafny.Map
		_ = _170_v80
		_170_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _169_v79)
		var _171_v81 _dafny.Sequence
		_ = _171_v81
		_171_v81 = _dafny.UnicodeSeqOfUtf8Bytes("mw")
		(globalState).F16 = ((_169_v79).Merge((func() _dafny.Map {
			if (_170_v80).Contains(_dafny.IntOfUint32((_171_v81).Cardinality())) {
				return (_170_v80).Get(_dafny.IntOfUint32((_171_v81).Cardinality())).(_dafny.Map)
			}
			return _169_v79
		})())).Cardinality()
		var _172_v82 *C0
		_ = _172_v82
		var _nw26 *C0 = New_C0_()
		_ = _nw26
		_nw26.Ctor__()
		_172_v82 = _nw26
		var _173_v83 _dafny.Array
		_ = _173_v83
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(3))
		_ = _nw27
		_173_v83 = _nw27
		var _174_v84 D10
		_ = _174_v84
		_174_v84 = Companion_D10_.Create_DC20_(_167_cf44, Companion_Default___.Fm5(p0, p0, globalState))
		var _175_v85 _dafny.Set
		_ = _175_v85
		_175_v85 = _dafny.SetOf(_174_v84)
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_173_v83), 0))
		_ = _index16
		(_173_v83).ArraySet1(_175_v85, (_index16).Int())
		var _176_v87 _dafny.Sequence
		_ = _176_v87
		_176_v87 = _dafny.SeqOf(_174_v84, _174_v84)
		var _177_v88 _dafny.Map
		_ = _177_v88
		_177_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Set {
			var _coll27 = _dafny.NewBuilder()
			_ = _coll27
			for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(961), _dafny.IntOfInt64(914))); ; {
				_compr_27, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _178_v86 _dafny.Int
				_178_v86 = interface{}(_compr_27).(_dafny.Int)
				if ((_dafny.IntOfInt64(961)).Cmp(_178_v86) <= 0) && ((_178_v86).Cmp(_dafny.IntOfInt64(914)) < 0) {
					_coll27.Add((_178_v86).Plus(_dafny.IntOfUint32((_171_v81).Cardinality())))
				}
			}
			return _coll27.ToSet()
		}()).Cardinality()), _176_v87)
		var _179_v89 _dafny.Sequence
		_ = _179_v89
		_179_v89 = _dafny.SeqOf(p0)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_173_v83), 0))
		_ = _index17
		(_173_v83).ArraySet1((_175_v85).Union(func() _dafny.Set {
			var _coll28 = _dafny.NewBuilder()
			_ = _coll28
			for _iter28 := _dafny.Iterate(((func() _dafny.Sequence {
				if (_177_v88).Contains(_dafny.IntOfUint32((_179_v89).Cardinality())) {
					return (_177_v88).Get(_dafny.IntOfUint32((_179_v89).Cardinality())).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer26 func(_dafny.Int) D10) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}((func(_180_cf44 bool, _181_v43 *C5, _182_p2 bool, _183_p0 _dafny.Int) func(_dafny.Int) D10 {
					return func(_184_i8 _dafny.Int) D10 {
						return Companion_D10_.Create_DC20_((Companion_D6_.Create_DC12_(_180_cf44, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_181_v43).F26(), _182_p2)).Cardinality())).Dtor_cf17(), _183_p0)
					}
				})(_167_cf44, _117_v43, p2, p0)))
			})()).Elements()); ; {
				_compr_28, _ok28 := _iter28()
				if !_ok28 {
					break
				}
				var _185_v90 D10
				_185_v90 = interface{}(_compr_28).(D10)
				if _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
					if (_177_v88).Contains(_dafny.IntOfUint32((_179_v89).Cardinality())) {
						return (_177_v88).Get(_dafny.IntOfUint32((_179_v89).Cardinality())).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer27 func(_dafny.Int) D10) func(_dafny.Int) interface{} {
						return func(arg27 _dafny.Int) interface{} {
							return coer27(arg27)
						}
					}((func(_186_cf44 bool, _187_v43 *C5, _188_p2 bool, _189_p0 _dafny.Int) func(_dafny.Int) D10 {
						return func(_184_i8 _dafny.Int) D10 {
							return Companion_D10_.Create_DC20_((Companion_D6_.Create_DC12_(_186_cf44, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_187_v43).F26(), _188_p2)).Cardinality())).Dtor_cf17(), _189_p0)
						}
					})(_167_cf44, _117_v43, p2, p0)))
				})(), _185_v90) {
					_coll28.Add(_185_v90)
				}
			}
			return _coll28.ToSet()
		}()), (_index17).Int())
		var _190_v91 _dafny.Array
		_ = _190_v91
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw28
		_190_v91 = _nw28
		var _191_v92 _dafny.Map
		_ = _191_v92
		_191_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _117_v43)
		var _192_v93 D2
		_ = _192_v93
		_192_v93 = Companion_D2_.Create_DC3_(_dafny.IntOfUint32((_171_v81).Cardinality()), (_191_v92).Cardinality())
		var _193_v94 D7
		_ = _193_v94
		_193_v94 = Companion_D7_.Create_DC14_(p2, p0, (_192_v93).Dtor_cf3())
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_190_v91), 0))
		_ = _index18
		(_190_v91).ArraySet1(Companion_Default___.SafeDivisionInt((_193_v94).Dtor_cf22(), _dafny.IntOfUint32((Companion_Default___.Fm26(globalState)).Cardinality())), (_index18).Int())
		var _194_v95 *C3
		_ = _194_v95
		var _nw29 *C3 = New_C3_()
		_ = _nw29
		_nw29.Ctor__(_167_cf44, (_171_v81).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_171_v81).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.IntOfInt64(547))
		_194_v95 = _nw29
		var _195_v96 _dafny.Map
		_ = _195_v96
		_195_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v95, _171_v81)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_190_v91), 0))
		_ = _index19
		(_190_v91).ArraySet1(Companion_Default___.Fm17((_195_v96).Contains(_194_v95), p2, (_117_v43).F26(), globalState), (_index19).Int())
	} else if _source4.Is_DC28() {
		var _196___mcc_h3 _dafny.Set = _source4.Get_().(D13_DC28).Cf42
		_ = _196___mcc_h3
		var _197_cf42 _dafny.Set = _196___mcc_h3
		_ = _197_cf42
		(globalState).F10 = (p0).Times((_dafny.IntOfInt64(667)).Minus(_dafny.IntOfInt64(450)))
		var _198_v97 _dafny.MultiSet
		_ = _198_v97
		_198_v97 = _dafny.MultiSetOf(p2)
		var _199_v98 _dafny.Map
		_ = _199_v98
		_199_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_198_v97, p0)
		var _200_v99 _dafny.CodePoint
		_ = _200_v99
		_200_v99 = _dafny.CodePoint('v')
		var _201_v100 _dafny.Map
		_ = _201_v100
		_201_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
		var _202_v101 T1
		_ = _202_v101
		var _nw30 *C4 = New_C4_()
		_ = _nw30
		_nw30.Ctor__(_199_v98, p3, _200_v99, p0)
		_202_v101 = _nw30
		var _203_v102 _dafny.Map
		_ = _203_v102
		_203_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_202_v101, p0)
		var _204_v103 _dafny.MultiSet
		_ = _204_v103
		_204_v103 = _dafny.MultiSetOf((_203_v102).Cardinality(), _202_v101.F25())
		var _205_v104 _dafny.Sequence
		_ = _205_v104
		_205_v104 = _dafny.SeqOf(p0, (_201_v100).Cardinality(), (_204_v103).Cardinality())
		var _206_v105 T0
		_ = _206_v105
		var _nw31 *C4 = New_C4_()
		_ = _nw31
		_nw31.Ctor__(_199_v98, p3, _200_v99, _dafny.IntOfUint32((_205_v104).Cardinality()))
		_206_v105 = _nw31
		var _207_v106 _dafny.MultiSet
		_ = _207_v106
		_207_v106 = _dafny.MultiSetOf((func() T0 {
			if p2 {
				return _206_v105
			}
			return _206_v105
		})(), _206_v105)
		var _rhs21 _dafny.MultiSet = (_207_v106).Difference(_207_v106)
		_ = _rhs21
		var _rhs22 bool = ((p0).Plus((_dafny.Zero).Minus(_206_v105.F25()))).Cmp(p0) != 0
		_ = _rhs22
		var _lhs19 *GlobalState = globalState
		_ = _lhs19
		_207_v106 = _rhs21
		_lhs19.F18 = _rhs22
		if !((!(p3)) && ((func() bool {
			if p3 {
				return (_117_v43).F26()
			}
			return p2
		})())) {
			var _208_v107 *C3
			_ = _208_v107
			var _nw32 *C3 = New_C3_()
			_ = _nw32
			_nw32.Ctor__((_117_v43).F26(), (_206_v105).F24(), _206_v105.F25())
			_208_v107 = _nw32
			var _209_v108 *C3
			_ = _209_v108
			_209_v108 = _208_v107
			var _210_v109 _dafny.Map
			_ = _210_v109
			_210_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_209_v108), _dafny.SeqOf(_dafny.IntOfInt64(563)))
			var _211_v110 _dafny.MultiSet
			_ = _211_v110
			_211_v110 = _dafny.MultiSetOf(_205_v104, _205_v104, _dafny.SeqOf(_206_v105.F25()))
			var _212_v111 _dafny.Array
			_ = _212_v111
			var _nwElement0_5 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_205_v104, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-929), _dafny.IntOfUint32((_205_v104).Cardinality()))).Uint32(), (_205_v104).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm17((_117_v43).F26(), p1, false, globalState), _dafny.IntOfUint32((_205_v104).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _nwElement0_5
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(20))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_5, 0)
			(_nw33).ArraySet1(_205_v104, 1)
			(_nw33).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(817))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_213_v101 T1) func(_dafny.Int) _dafny.Int {
				return func(_214_i9 _dafny.Int) _dafny.Int {
					return _213_v101.F25()
				}
			})(_202_v101))), 2)
			(_nw33).ArraySet1((func() _dafny.Sequence {
				if (_210_v109).Contains(_208_v107) {
					return (_210_v109).Get(_208_v107).(_dafny.Sequence)
				}
				return _205_v104
			})(), 3)
			(_nw33).ArraySet1(_205_v104, 4)
			(_nw33).ArraySet1(_205_v104, 5)
			(_nw33).ArraySet1(_205_v104, 6)
			(_nw33).ArraySet1(_205_v104, 7)
			(_nw33).ArraySet1(_205_v104, 8)
			(_nw33).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-330))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_215_v97 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
				return func(_216_i10 _dafny.Int) _dafny.Int {
					return (_215_v97).Cardinality()
				}
			})(_198_v97))), 9)
			(_nw33).ArraySet1(_205_v104, 10)
			(_nw33).ArraySet1(_205_v104, 11)
			(_nw33).ArraySet1(_205_v104, 12)
			(_nw33).ArraySet1(_dafny.SeqOf(_206_v105.F25(), _dafny.IntOfUint32((_116_v42).Cardinality()), _202_v101.F25()), 13)
			(_nw33).ArraySet1(_205_v104, 14)
			(_nw33).ArraySet1(_205_v104, 15)
			(_nw33).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-866))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_217_v104 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_218_i11 _dafny.Int) _dafny.Int {
					return (_217_v104).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.IntOfUint32((_217_v104).Cardinality()))).Uint32()).(_dafny.Int)
				}
			})(_205_v104))), 16)
			(_nw33).ArraySet1(_205_v104, 17)
			(_nw33).ArraySet1(_205_v104, 18)
			(_nw33).ArraySet1(_dafny.Companion_Sequence_.Update(_205_v104, (Companion_Default___.SafeIndex(_206_v105.F25(), _dafny.IntOfUint32((_205_v104).Cardinality()))).Uint32(), (_211_v110).Cardinality()), 19)
			_212_v111 = _nw33
			var _219_v112 _dafny.Map
			_ = _219_v112
			_219_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_117_v43).F26(), _212_v111)
			_219_v112 = (_219_v112).Update((_117_v43).F26(), _212_v111)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_212_v111), 0))
			_ = _index20
			(_212_v111).ArraySet1(_205_v104, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_212_v111), 0))
			_ = _index21
			(_212_v111).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_206_v105.F25(), _202_v101.F25()), _dafny.Companion_Sequence_.Concatenate(_205_v104, _205_v104)), (_index21).Int())
			var _220_v113 _dafny.Map
			_ = _220_v113
			_220_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_116_v42).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_116_v42).Cardinality()))).Uint32()).(bool))
			var _221_v114 D11
			_ = _221_v114
			_221_v114 = Companion_D11_.Create_DC24_((_117_v43).F26(), (Companion_Default___.Fm34(p2, (_220_v113).Cardinality(), (_117_v43).F26(), globalState)).Cardinality())
			var _222_v115 *C1
			_ = _222_v115
			var _nw34 *C1 = New_C1_()
			_ = _nw34
			_nw34.Ctor__((_202_v101).F24(), (_202_v101.F25()).Minus((_221_v114).Dtor_cf37()))
			_222_v115 = _nw34
			var _223_v116 D3
			_ = _223_v116
			_223_v116 = Companion_D3_.Create_DC5_(_116_v42, _202_v101.F25(), _202_v101.F25())
			var _224_v117 D10
			_ = _224_v117
			_224_v117 = Companion_D10_.Create_DC20_((_117_v43).F26(), (_223_v116).Dtor_cf7())
			(globalState).F18 = (_224_v117).Dtor_cf28()
			_200_v99 = _dafny.CodePoint('a')
		} else {
			var _225_v118 _dafny.Map
			_ = _225_v118
			_225_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_117_v43).F26(), p0)
			(globalState).F11 = (func() _dafny.Int {
				if (_225_v118).Contains(p3) {
					return (_225_v118).Get(p3).(_dafny.Int)
				}
				return _202_v101.F25()
			})()
			var _226_v119 D3
			_ = _226_v119
			_226_v119 = Companion_D3_.Create_DC5_(_116_v42, _206_v105.F25(), _206_v105.F25())
			_226_v119 = _226_v119
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_158_v78), 0))
			_ = _index22
			(_158_v78).ArraySet1((_198_v97).IsDisjointFrom(_198_v97), (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_158_v78), 0))
			_ = _index23
			(_158_v78).ArraySet1((_117_v43).F26(), (_index23).Int())
			var _227_v120 _dafny.Array
			_ = _227_v120
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw35
			_227_v120 = _nw35
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_227_v120), 0))
			_ = _index24
			(_227_v120).ArraySet1(_202_v101.F25(), (_index24).Int())
			var _228_v121 D10
			_ = _228_v121
			_228_v121 = Companion_D10_.Create_DC21_(_202_v101.F25(), !((_158_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_158_v78), 0))).Int()).(bool)), (_117_v43).F26())
			var _229_v122 D11
			_ = _229_v122
			_229_v122 = Companion_D11_.Create_DC25_(Companion_D11_.Create_DC24_(p3, Companion_Default___.Fm17((Companion_D11_.Create_DC24_(p3, _206_v105.F25())).Dtor_cf36(), (_228_v121).Dtor_cf32(), p2, globalState)))
			var _230_v123 _dafny.MultiSet
			_ = _230_v123
			_230_v123 = _dafny.MultiSetOf(_229_v122, _229_v122, _229_v122)
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_227_v120), 0))
			_ = _index25
			var _rhs23 _dafny.Int = _dafny.IntOfInt64(795)
			_ = _rhs23
			var _rhs24 _dafny.Int = (func() _dafny.Int {
				if (_225_v118).Contains((_117_v43).F26()) {
					return (_225_v118).Get((_117_v43).F26()).(_dafny.Int)
				}
				return (_230_v123).Cardinality()
			})()
			_ = _rhs24
			var _rhs25 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(-545)).Plus(_206_v105.F25()), _202_v101.F25())
			_ = _rhs25
			var _lhs20 _dafny.Array = _227_v120
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_227_v120), 0))
			_ = _lhs21
			var _lhs22 *GlobalState = globalState
			_ = _lhs22
			var _lhs23 *GlobalState = globalState
			_ = _lhs23
			(_lhs20).ArraySet1(_rhs23, (_lhs21).Int())
			_lhs22.F22 = _rhs24
			_lhs23.F10 = _rhs25
			var _231_v124 _dafny.Array
			_ = _231_v124
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_5
			var _nw36 _dafny.Array
			_ = _nw36
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw36 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_232_v101 T1, _233_v120 _dafny.Array) func(_dafny.Int) bool {
					return func(_234_i12 _dafny.Int) bool {
						return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf((_233_v120).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_233_v120), 0))).Int()).(_dafny.Int), _232_v101.F25()), _dafny.SetOf(_232_v101.F25())))).IsSubsetOf(_dafny.MultiSetOf(_dafny.SetOf(_232_v101.F25())))
					}
				})(_202_v101, _227_v120)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw36 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw36).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw36).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_231_v124 = _nw36
			_231_v124 = (func() _dafny.Array {
				if p1 {
					return _231_v124
				}
				return _231_v124
			})()
		}
		var _235_v125 _dafny.Array
		_ = _235_v125
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_6
		var _nw37 _dafny.Array
		_ = _nw37
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw37 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Sequence = (func(_236_v104 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_237_i13 _dafny.Int) _dafny.Sequence {
					return _236_v104
				}
			})(_205_v104)
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
		_235_v125 = _nw37
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_235_v125), 0))
		_ = _index26
		(_235_v125).ArraySet1(_205_v104, (_index26).Int())
		var _238_v126 D10
		_ = _238_v126
		_238_v126 = Companion_D10_.Create_DC19_((_206_v105).F24())
		var _239_v127 _dafny.Map
		_ = _239_v127
		_239_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v78, (_117_v43).F26())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_235_v125), 0))
		_ = _index27
		var _rhs26 _dafny.Int = (_206_v105.F25()).Times(_206_v105.F25())
		_ = _rhs26
		var _rhs27 _dafny.Int = (_dafny.Zero).Minus(p0)
		_ = _rhs27
		var _rhs28 _dafny.Sequence = _dafny.SeqOf((_dafny.IntOfInt64(703)).Plus(_dafny.IntOfInt64(-504)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pvsknr")).Cardinality()))
		_ = _rhs28
		var _rhs29 _dafny.Int = ((_239_v127).Cardinality()).Plus(_206_v105.F25())
		_ = _rhs29
		var _rhs30 D10 = _238_v126
		_ = _rhs30
		var _lhs24 *GlobalState = globalState
		_ = _lhs24
		var _lhs25 T1 = _202_v101
		_ = _lhs25
		var _lhs26 _dafny.Array = _235_v125
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_235_v125), 0))
		_ = _lhs27
		var _lhs28 *GlobalState = globalState
		_ = _lhs28
		_lhs24.F0 = _rhs26
		_lhs25.F25_set_(_rhs27)
		(_lhs26).ArraySet1(_rhs28, (_lhs27).Int())
		_lhs28.F2 = _rhs29
		_238_v126 = _rhs30
	} else {
		var _240___mcc_h4 D13 = _source4.Get_().(D13_DC30).Cf46
		_ = _240___mcc_h4
		var _241_cf46 D13 = _240___mcc_h4
		_ = _241_cf46
		var _242_v128 _dafny.Set
		_ = _242_v128
		_242_v128 = _dafny.SetOf(p0)
		var _243_v129 _dafny.Map
		_ = _243_v129
		_243_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_242_v128).Cardinality())
		var _244_v130 D11
		_ = _244_v130
		_244_v130 = Companion_D11_.Create_DC24_(p2, _dafny.IntOfUint32((_116_v42).Cardinality()))
		var _245_v131 _dafny.Map
		_ = _245_v131
		_245_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_243_v129).Contains((_244_v130).Dtor_cf36()) {
				return (_243_v129).Get((_244_v130).Dtor_cf36()).(_dafny.Int)
			}
			return p0
		})(), p0)
		(globalState).F10 = (func() _dafny.Int {
			if (_245_v131).Contains(p0) {
				return (_245_v131).Get(p0).(_dafny.Int)
			}
			return p0
		})()
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_158_v78), 0))
		_ = _index28
		(_158_v78).ArraySet1(p2, (_index28).Int())
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_158_v78), 0))
		_ = _index29
		var _rhs31 bool = (_117_v43).F26()
		_ = _rhs31
		var _rhs32 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(p0))
		_ = _rhs32
		var _lhs29 _dafny.Array = _158_v78
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_158_v78), 0))
		_ = _lhs30
		var _lhs31 *GlobalState = globalState
		_ = _lhs31
		(_lhs29).ArraySet1(_rhs31, (_lhs30).Int())
		_lhs31.F0 = _rhs32
		var _246_v132 _dafny.Map
		_ = _246_v132
		_246_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
		var _247_v133 _dafny.Set
		_ = _247_v133
		_247_v133 = _dafny.SetOf(_246_v132)
		var _248_v134 D3
		_ = _248_v134
		_248_v134 = Companion_D3_.Create_DC4_(_247_v133)
		var _249_v135 _dafny.Sequence
		_ = _249_v135
		_249_v135 = _116_v42
		var _250_v136 _dafny.Map
		_ = _250_v136
		_250_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_249_v135).Cardinality())), Companion_D3_.Create_DC4_(_247_v133))
		var _251_v137 _dafny.Array
		_ = _251_v137
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_7
		var _nw38 _dafny.Array
		_ = _nw38
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw38 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.CodePoint = func(_252_i14 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			}
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw38 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw38).ArraySet1CodePoint(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw38).ArraySet1CodePoint(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_251_v137 = _nw38
		var _253_v138 _dafny.Sequence
		_ = _253_v138
		_253_v138 = _dafny.SeqOf(_251_v137, _251_v137)
		var _254_v139 bool
		_ = _254_v139
		var _255_v140 _dafny.Sequence
		_ = _255_v140
		var _256_v141 bool
		_ = _256_v141
		var _out0 bool
		_ = _out0
		var _out1 _dafny.Sequence
		_ = _out1
		var _out2 bool
		_ = _out2
		_out0, _out1, _out2 = (_117_v43).M1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _248_v134)).Merge(_250_v136), ((_116_v42).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_116_v42).Cardinality()))).Uint32()).(bool)) == (p3), _253_v138, globalState)
		_254_v139 = _out0
		_255_v140 = _out1
		_256_v141 = _out2
		var _257_v142 *C0
		_ = _257_v142
		var _nw39 *C0 = New_C0_()
		_ = _nw39
		_nw39.Ctor__()
		_257_v142 = _nw39
	}
	var _258_v143 _dafny.Map
	_ = _258_v143
	_258_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
	var _259_v144 _dafny.Map
	_ = _259_v144
	_259_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_258_v143, _158_v78)
	_259_v144 = (_259_v144).Update(func() _dafny.Map {
		var _coll29 = _dafny.NewMapBuilder()
		_ = _coll29
		for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(691), _dafny.IntOfInt64(699))); ; {
			_compr_29, _ok29 := _iter29()
			if !_ok29 {
				break
			}
			var _260_v145 _dafny.Int
			_260_v145 = interface{}(_compr_29).(_dafny.Int)
			if ((_dafny.IntOfInt64(691)).Cmp(_260_v145) <= 0) && ((_260_v145).Cmp(_dafny.IntOfInt64(699)) < 0) {
				_coll29.Add((_260_v145).Times(_dafny.IntOfInt64(919)), p3)
			}
		}
		return _coll29.ToMap()
	}(), _158_v78)
	var _261_v146 _dafny.Array
	_ = _261_v146
	var _len0_8 _dafny.Int = _dafny.IntOfInt64(10)
	_ = _len0_8
	var _nw40 _dafny.Array
	_ = _nw40
	if _len0_8.Cmp(_dafny.Zero) == 0 {
		_nw40 = _dafny.NewArray(_len0_8)
	} else {
		var _init8 func(_dafny.Int) _dafny.Int = (func(_262_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_263_i15 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_263_i15, _262_p0)
			}
		})(p0)
		_ = _init8
		var _element0_8 = _init8(_dafny.Zero)
		_ = _element0_8
		_nw40 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
		(_nw40).ArraySet1(_element0_8, 0)
		var _nativeLen0_8 = (_len0_8).Int()
		_ = _nativeLen0_8
		for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
			(_nw40).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
		}
	}
	_261_v146 = _nw40
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_261_v146), 0))
	_ = _index30
	(_261_v146).ArraySet1(p0, (_index30).Int())
	var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_261_v146), 0))
	_ = _index31
	(_261_v146).ArraySet1((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_116_v42, _116_v42), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_116_v42, _116_v42)).Cardinality()))).Uint32(), p3))).Cardinality(), (_index31).Int())
	var _264_v147 D10
	_ = _264_v147
	_264_v147 = Companion_D10_.Create_DC21_(_dafny.IntOfInt64(124), p2, p1)
	r0 = func(_source5 D10) _dafny.CodePoint {
		if _source5.Is_DC19() {
			var _265___mcc_h5 _dafny.CodePoint = _source5.Get_().(D10_DC19).Cf27
			_ = _265___mcc_h5
			var _266_cf27 _dafny.CodePoint = _265___mcc_h5
			_ = _266_cf27
			return _266_cf27
		} else if _source5.Is_DC20() {
			var _267___mcc_h6 bool = _source5.Get_().(D10_DC20).Cf28
			_ = _267___mcc_h6
			var _268___mcc_h7 _dafny.Int = _source5.Get_().(D10_DC20).Cf29
			_ = _268___mcc_h7
			var _269_cf29 _dafny.Int = _268___mcc_h7
			_ = _269_cf29
			var _270_cf28 bool = _267___mcc_h6
			_ = _270_cf28
			return _dafny.CodePoint('l')
		} else if _source5.Is_DC21() {
			var _271___mcc_h8 _dafny.Int = _source5.Get_().(D10_DC21).Cf30
			_ = _271___mcc_h8
			var _272___mcc_h9 bool = _source5.Get_().(D10_DC21).Cf31
			_ = _272___mcc_h9
			var _273___mcc_h10 bool = _source5.Get_().(D10_DC21).Cf32
			_ = _273___mcc_h10
			var _274_cf32 bool = _273___mcc_h10
			_ = _274_cf32
			var _275_cf31 bool = _272___mcc_h9
			_ = _275_cf31
			var _276_cf30 _dafny.Int = _271___mcc_h8
			_ = _276_cf30
			return _dafny.CodePoint('n')
		} else {
			var _277___mcc_h11 _dafny.Set = _source5.Get_().(D10_DC18).Cf26
			_ = _277___mcc_h11
			var _278_cf26 _dafny.Set = _277___mcc_h11
			_ = _278_cf26
			return _dafny.CodePoint('o')
		}
	}(_264_v147)
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _279_v0 _dafny.Sequence
	_ = _279_v0
	_279_v0 = _dafny.SeqOf(_dafny.CodePoint('w'))
	var _280_v1 bool
	_ = _280_v1
	_280_v1 = true
	var _281_v2 _dafny.Int
	_ = _281_v2
	_281_v2 = _dafny.IntOfInt64(742)
	var _282_v3 _dafny.Map
	_ = _282_v3
	_282_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v1, _281_v2)
	var _283_v4 _dafny.Array
	_ = _283_v4
	var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
	_ = _nw41
	_283_v4 = _nw41
	var _284_v5 _dafny.Array
	_ = _284_v5
	var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
	_ = _nw42
	_284_v5 = _nw42
	var _285_v6 _dafny.Map
	_ = _285_v6
	_285_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v4, _284_v5)
	var _286_v7 _dafny.MultiSet
	_ = _286_v7
	_286_v7 = _dafny.MultiSetOf(_280_v1)
	var _287_v8 _dafny.Sequence
	_ = _287_v8
	_287_v8 = _dafny.SeqOf((_286_v7).Cardinality())
	var _288_v9 _dafny.Set
	_ = _288_v9
	_288_v9 = _dafny.SetOf(_287_v8)
	var _289_globalState *GlobalState
	_ = _289_globalState
	var _nw43 *GlobalState = New_GlobalState_()
	_ = _nw43
	_nw43.Ctor__(_dafny.IntOfInt64(-928), _dafny.IntOfInt64(141), _dafny.IntOfInt64(124), _dafny.Companion_Sequence_.Concatenate(_279_v0, _279_v0), _dafny.IntOfInt64(227), false, _dafny.CodePoint('d'), _282_v3, _285_v6, _dafny.IntOfInt64(-50), _dafny.IntOfInt64(-492), _dafny.IntOfInt64(-627), false, (_282_v3).Merge(_282_v3), _288_v9, _dafny.IntOfInt64(968), _dafny.IntOfInt64(343), false, false, _287_v8, _dafny.IntOfInt64(425), true, _dafny.IntOfInt64(-118))
	_289_globalState = _nw43
	var _290_v10 _dafny.MultiSet
	_ = _290_v10
	_290_v10 = _dafny.MultiSetOf(_281_v2, _281_v2)
	var _291_v11 _dafny.Sequence
	_ = _291_v11
	_291_v11 = _dafny.SeqOf(_280_v1, Companion_Default___.Fm0((_290_v10).Cardinality(), _280_v1, _280_v1, (_282_v3).Cardinality(), _289_globalState))
	var _292_v12 _dafny.Sequence
	_ = _292_v12
	_292_v12 = _291_v11
	var _293_v13 _dafny.Sequence
	_ = _293_v13
	_293_v13 = _dafny.SeqOf(_284_v5, _284_v5, _284_v5, _284_v5)
	var _294_v14 _dafny.Sequence
	_ = _294_v14
	_294_v14 = _279_v0
	var _295_v15 _dafny.Sequence
	_ = _295_v15
	_295_v15 = _dafny.SeqOf((_294_v14), _279_v0, _279_v0, _dafny.UnicodeSeqOfUtf8Bytes("fln"), _279_v0)
	var _296_v16 _dafny.CodePoint
	_ = _296_v16
	_296_v16 = _dafny.CodePoint('j')
	var _rhs33 _dafny.Sequence = _279_v0
	_ = _rhs33
	var _rhs34 _dafny.Int = _281_v2
	_ = _rhs34
	var _rhs35 _dafny.Int = _281_v2
	_ = _rhs35
	var _rhs36 _dafny.Int = (_dafny.Zero).Minus((_281_v2).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_292_v12), _291_v11)).Cardinality())))
	_ = _rhs36
	var _rhs37 _dafny.Array = (_293_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_295_v15).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_295_v15).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_295_v15).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_295_v15).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32(((_295_v15).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_295_v15).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _296_v16)).Cardinality()), _dafny.IntOfUint32((_293_v13).Cardinality()))).Uint32()).(_dafny.Array)
	_ = _rhs37
	var _lhs32 *GlobalState = _289_globalState
	_ = _lhs32
	var _lhs33 *GlobalState = _289_globalState
	_ = _lhs33
	var _lhs34 *GlobalState = _289_globalState
	_ = _lhs34
	_279_v0 = _rhs33
	_lhs32.F11 = _rhs34
	_lhs33.F10 = _rhs35
	_lhs34.F16 = _rhs36
	_284_v5 = _rhs37
	(_289_globalState).F2 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_287_v8, _287_v8), _287_v8)).Cardinality()))
	var _297_i0 _dafny.Int
	_ = _297_i0
	_297_i0 = _dafny.Zero
	{
		for true {
			{
				if (_297_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_297_i0 = (_297_i0).Plus(_dafny.One)
				var _298_v17 *C4
				_ = _298_v17
				var _nw44 *C4 = New_C4_()
				_ = _nw44
				_nw44.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v7, _dafny.IntOfUint32((_291_v11).Cardinality())), !(_280_v1), _296_v16, _281_v2)
				_298_v17 = _nw44
				_296_v16 = (func() _dafny.CodePoint {
					if !(!(_280_v1)) || (_280_v1) {
						return _296_v16
					}
					return _296_v16
				})()
				var _source6 _dafny.Sequence = _294_v14
				_ = _source6
				var _299___mcc_h0 _dafny.Sequence = _source6
				_ = _299___mcc_h0
				var _300_cf1 _dafny.Sequence = _299___mcc_h0
				_ = _300_cf1
				var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_283_v4), 0))
				_ = _index32
				(_283_v4).ArraySet1(false, (_index32).Int())
				var _301_v18 _dafny.Map
				_ = _301_v18
				_301_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v2, (func() _dafny.Int {
					if (_282_v3).Contains((_291_v11).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_291_v11).Cardinality()))).Uint32()).(bool)) {
						return (_282_v3).Get((_291_v11).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_291_v11).Cardinality()))).Uint32()).(bool)).(_dafny.Int)
					}
					return _281_v2
				})())
				var _302_v19 D3
				_ = _302_v19
				_302_v19 = Companion_D3_.Create_DC6_((func() _dafny.Int {
					if (_301_v18).Contains(_281_v2) {
						return (_301_v18).Get(_281_v2).(_dafny.Int)
					}
					return _281_v2
				})(), _281_v2, (_298_v17).F30())
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_283_v4), 0))
				_ = _index33
				(_283_v4).ArraySet1((_302_v19).Dtor_cf11(), (_index33).Int())
				(_289_globalState).F0 = (_287_v8).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_287_v8).Cardinality()))).Uint32()).(_dafny.Int)
				(_289_globalState).F15 = _281_v2
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_283_v4), 0))
				_ = _index34
				(_283_v4).ArraySet1(false, (_index34).Int())
				var _303_v20 D5
				_ = _303_v20
				_303_v20 = Companion_D5_.Create_DC10_(false, true)
				var _304_v21 _dafny.Map
				_ = _304_v21
				_304_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_298_v17).F30())), _303_v20)
				(_298_v17).M5((_304_v21).Merge(_304_v21), true, _281_v2, _289_globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _305_v22 _dafny.Set
	_ = _305_v22
	_305_v22 = _dafny.SetOf(_281_v2, _dafny.IntOfUint32((_291_v11).Cardinality()))
	var _306_v23 _dafny.Map
	_ = _306_v23
	_306_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_287_v8).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_287_v8).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.Fm0(_281_v2, _280_v1, false, (func() _dafny.Int {
		if (_282_v3).Contains(_280_v1) {
			return (_282_v3).Get(_280_v1).(_dafny.Int)
		}
		return _281_v2
	})(), _289_globalState))
	(_289_globalState).F2 = (func() _dafny.Int {
		if false {
			return (_305_v22).Cardinality()
		}
		return (_306_v23).Cardinality()
	})()
	var _307_v24 _dafny.CodePoint
	_ = _307_v24
	var _out3 _dafny.CodePoint
	_ = _out3
	_out3 = Companion_Default___.M6(_281_v2, _280_v1, _280_v1, _280_v1, _289_globalState)
	_307_v24 = _out3
	var _308_v25 D11
	_ = _308_v25
	_308_v25 = Companion_D11_.Create_DC23_(_281_v2, _280_v1)
	var _309_v26 D11
	_ = _309_v26
	_309_v26 = Companion_D11_.Create_DC25_(_308_v25)
	var _310_v27 _dafny.Array
	_ = _310_v27
	var _nwElement0_6 D11 = _309_v26
	_ = _nwElement0_6
	var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(4))
	_ = _nw45
	(_nw45).ArraySet1(_nwElement0_6, 0)
	(_nw45).ArraySet1(_309_v26, 1)
	(_nw45).ArraySet1(_309_v26, 2)
	(_nw45).ArraySet1(Companion_D11_.Create_DC25_(_308_v25), 3)
	_310_v27 = _nw45
	var _311_v28 _dafny.Set
	_ = _311_v28
	_311_v28 = _dafny.SetOf(_310_v27, _310_v27)
	var _312_v29 _dafny.Map
	_ = _312_v29
	_312_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v1, _311_v28)
	_312_v29 = (_312_v29).Update(_280_v1, _311_v28)
	if _280_v1 {
		var _313_v30 T1
		_ = _313_v30
		var _nw46 *C3 = New_C3_()
		_ = _nw46
		_nw46.Ctor__(true, _296_v16, (_287_v8).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_287_v8).Cardinality()))).Uint32()).(_dafny.Int))
		_313_v30 = _nw46
		var _314_v31 _dafny.MultiSet
		_ = _314_v31
		_314_v31 = _dafny.MultiSetOf(_313_v30)
		var _315_v32 _dafny.CodePoint
		_ = _315_v32
		var _out4 _dafny.CodePoint
		_ = _out4
		_out4 = Companion_Default___.M6(_281_v2, (_281_v2).Cmp(Companion_Default___.Fm5(_281_v2, (_314_v31).Cardinality(), _289_globalState)) == 0, _dafny.Companion_Sequence_.Equal(_279_v0, _279_v0), (_281_v2).Cmp(_281_v2) < 0, _289_globalState)
		_315_v32 = _out4
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xvkixbphw"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(618))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}((func(_316_v24 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_317_i1 _dafny.Int) _dafny.CodePoint {
				return _316_v24
			}
		})(_307_v24)))), _dafny.Companion_Sequence_.Concatenate(_279_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(912))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}((func(_318_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_319_i2 _dafny.Int) _dafny.CodePoint {
				return _318_v16
			}
		})(_296_v16))))) {
			_280_v1 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_287_v8, _287_v8), ((_dafny.SetOf(_313_v30)).Cardinality()).Times(_313_v30.F25()))
			(_289_globalState).F18 = _280_v1
			var _320_v33 _dafny.CodePoint
			_ = _320_v33
			var _out5 _dafny.CodePoint
			_ = _out5
			_out5 = Companion_Default___.M6((_313_v30.F25()).Times(_281_v2), _280_v1, (_286_v7).IsDisjointFrom(_dafny.MultiSetOf(_280_v1, _280_v1)), (_281_v2).Cmp(_313_v30.F25()) >= 0, _289_globalState)
			_320_v33 = _out5
			var _321_v34 _dafny.Map
			_ = _321_v34
			_321_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v7, _313_v30.F25())
			var _322_v36 _dafny.Set
			_ = _322_v36
			_322_v36 = _dafny.SetOf(Companion_Default___.Fm8(_313_v30.F25(), true, _dafny.SeqOf(_280_v1), _289_globalState))
			var _323_v37 *C4
			_ = _323_v37
			var _nw47 *C4 = New_C4_()
			_ = _nw47
			_nw47.Ctor__((_321_v34).Merge(func() _dafny.Map {
				var _coll30 = _dafny.NewMapBuilder()
				_ = _coll30
				for _iter30 := _dafny.Iterate((_322_v36).Elements()); ; {
					_compr_30, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _324_v35 _dafny.MultiSet
					_324_v35 = interface{}(_compr_30).(_dafny.MultiSet)
					if (_322_v36).Contains(_324_v35) {
						_coll30.Add(_324_v35, _313_v30.F25())
					}
				}
				return _coll30.ToMap()
			}()), (func() bool {
				if true {
					return _280_v1
				}
				return _280_v1
			})(), _dafny.CodePoint('b'), _313_v30.F25())
			_323_v37 = _nw47
			(_289_globalState).F3 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(941))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}((func(_325_v30 T1) func(_dafny.Int) _dafny.CodePoint {
				return func(_326_i3 _dafny.Int) _dafny.CodePoint {
					return (_325_v30).F24()
				}
			})(_313_v30)))
		} else {
			var _327_v38 _dafny.Map
			_ = _327_v38
			_327_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_290_v10).IsDisjointFrom(_290_v10), _283_v4)
			var _328_v39 _dafny.Map
			_ = _328_v39
			_328_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_313_v30).F24(), _280_v1)
			var _329_v40 D13
			_ = _329_v40
			_329_v40 = Companion_D13_.Create_DC29_(_280_v1, _280_v1, _283_v4)
			_327_v38 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v1, _283_v4)).Update((func() bool {
				if (_328_v39).Contains((_313_v30).F24()) {
					return (_328_v39).Get((_313_v30).F24()).(bool)
				}
				return _280_v1
			})(), (_329_v40).Dtor_cf45())).Merge(_327_v38)
			(_289_globalState).F18 = _280_v1
			var _330_v41 *C6
			_ = _330_v41
			var _nw48 *C6 = New_C6_()
			_ = _nw48
			_nw48.Ctor__(_306_v23)
			_330_v41 = _nw48
			var _331_v42 _dafny.Map
			_ = _331_v42
			_331_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_280_v1, _dafny.CodePoint('u'))
			_331_v42 = (_331_v42).Update(_280_v1, (_313_v30).F24())
			var _332_v43 _dafny.Array
			_ = _332_v43
			var _nw49 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
			_ = _nw49
			_332_v43 = _nw49
			_332_v43 = (func() _dafny.Array {
				if false {
					return _332_v43
				}
				return _332_v43
			})()
		}
		(_289_globalState).F18 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_296_v16, _280_v1)).Contains(_307_v24)
		(_289_globalState).F16 = _281_v2
		if ((func() bool {
			if _280_v1 {
				return _280_v1
			}
			return _280_v1
		})()) == (false) {
			var _333_v44 _dafny.Set
			_ = _333_v44
			_333_v44 = _dafny.SetOf((_313_v30.F25()).Cmp(_281_v2) != 0)
			var _334_v45 _dafny.Array
			_ = _334_v45
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(13))
			_ = _nw50
			_334_v45 = _nw50
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_334_v45), 0))
			_ = _index35
			(_334_v45).ArraySet1CodePoint(_315_v32, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_334_v45), 0))
			_ = _index36
			var _rhs38 _dafny.Int = (func() _dafny.Int {
				if (_291_v11).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_291_v11).Cardinality()))).Uint32()).(bool) {
					return (_281_v2).Plus(_281_v2)
				}
				return _313_v30.F25()
			})()
			_ = _rhs38
			var _rhs39 _dafny.Set = ((_dafny.SetOf(_280_v1)).Difference(_333_v44)).Difference(_333_v44)
			_ = _rhs39
			var _rhs40 bool = _280_v1
			_ = _rhs40
			var _rhs41 _dafny.CodePoint = _dafny.CodePoint('y')
			_ = _rhs41
			var _rhs42 _dafny.Array = (func() _dafny.Array {
				if _280_v1 {
					return _284_v5
				}
				return _284_v5
			})()
			_ = _rhs42
			var _lhs35 *GlobalState = _289_globalState
			_ = _lhs35
			var _lhs36 *GlobalState = _289_globalState
			_ = _lhs36
			var _lhs37 _dafny.Array = _334_v45
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_334_v45), 0))
			_ = _lhs38
			_lhs35.F16 = _rhs38
			_333_v44 = _rhs39
			_lhs36.F18 = _rhs40
			(_lhs37).ArraySet1CodePoint(_rhs41, (_lhs38).Int())
			_284_v5 = _rhs42
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_284_v5), 0))
			_ = _index37
			(_284_v5).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_287_v8, _287_v8)).Cardinality()), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_284_v5), 0))
			_ = _index38
			(_284_v5).ArraySet1(_313_v30.F25(), (_index38).Int())
			var _335_v46 _dafny.MultiSet
			_ = _335_v46
			_335_v46 = _dafny.MultiSetOf(_279_v0)
			(_289_globalState).F16 = (_dafny.Zero).Minus((_335_v46).Cardinality())
			_306_v23 = (_306_v23).Update(_313_v30.F25(), _280_v1)
			(_289_globalState).F18 = _280_v1
		} else {
			var _336_v47 _dafny.Map
			_ = _336_v47
			_336_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v7, _313_v30.F25())
			var _337_v48 D3
			_ = _337_v48
			_337_v48 = Companion_D3_.Create_DC6_(_281_v2, _313_v30.F25(), _280_v1)
			var _338_v49 D5
			_ = _338_v49
			_338_v49 = Companion_D5_.Create_DC10_(_280_v1, (_337_v48).Dtor_cf11())
			var _339_v50 T0
			_ = _339_v50
			var _nw51 *C4 = New_C4_()
			_ = _nw51
			_nw51.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf((func() bool {
				if (_306_v23).Contains(_281_v2) {
					return (_306_v23).Get(_281_v2).(bool)
				}
				return _280_v1
			})(), _280_v1)), _281_v2)).Merge(_336_v47), _280_v1, Companion_Default___.Fm10(_280_v1, _280_v1, !((_338_v49).Dtor_cf15()), _280_v1, _289_globalState), _281_v2)
			_339_v50 = _nw51
			var _340_v51 _dafny.Set
			_ = _340_v51
			_340_v51 = _dafny.SetOf(true)
			var _341_v52 D6
			_ = _341_v52
			_341_v52 = Companion_D6_.Create_DC11_(_339_v50)
			_339_v50 = (func() T0 {
				if Companion_Default___.Fm0((_340_v51).Cardinality(), false, _280_v1, _281_v2, _289_globalState) {
					return _339_v50
				}
				return (_341_v52).Dtor_cf16()
			})()
			(_289_globalState).F18 = _280_v1
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_284_v5), 0))
			_ = _index39
			(_284_v5).ArraySet1((_281_v2).Times(_339_v50.F25()), (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_284_v5), 0))
			_ = _index40
			(_284_v5).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_291_v11).Cardinality())), (_index40).Int())
			(_289_globalState).F18 = _280_v1
			var _342_v53 *C0
			_ = _342_v53
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__()
			_342_v53 = _nw52
		}
	} else {
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_284_v5), 0))
		_ = _index41
		(_284_v5).ArraySet1(_281_v2, (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.ArrayLen((_284_v5), 0))
		_ = _index42
		(_284_v5).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(161), _dafny.IntOfUint32((_dafny.SeqOf(_281_v2)).Cardinality())), _281_v2), (_index42).Int())
		(_289_globalState).F11 = _dafny.IntOfUint32((_279_v0).Cardinality())
		var _343_v54 *C2
		_ = _343_v54
		var _nw53 *C2 = New_C2_()
		_ = _nw53
		_nw53.Ctor__(_dafny.IntOfUint32((Companion_Default___.Fm7(_289_globalState)).Cardinality()), _280_v1)
		_343_v54 = _nw53
		var _344_v55 _dafny.Sequence
		_ = _344_v55
		_344_v55 = _dafny.SeqOf((func() *C2 {
			if _280_v1 {
				return _343_v54
			}
			return _343_v54
		})())
		_344_v55 = _dafny.Companion_Sequence_.Concatenate(_344_v55, _dafny.Companion_Sequence_.Concatenate(_344_v55, _344_v55))
		_291_v11 = _291_v11
		_343_v54 = _343_v54
	}
	if _280_v1 {
		var _345_v56 _dafny.CodePoint
		_ = _345_v56
		var _out6 _dafny.CodePoint
		_ = _out6
		_out6 = Companion_Default___.M6(_dafny.IntOfInt64(138), !((_dafny.IntOfInt64(-166)).Cmp(_281_v2) > 0), (_280_v1) == (_280_v1), _280_v1, _289_globalState)
		_345_v56 = _out6
		var _346_v57 *C0
		_ = _346_v57
		var _nw54 *C0 = New_C0_()
		_ = _nw54
		_nw54.Ctor__()
		_346_v57 = _nw54
		var _347_v58 _dafny.Map
		_ = _347_v58
		_347_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v57, _dafny.UnicodeSeqOfUtf8Bytes("kmijxxnsm"))
		_347_v58 = (_347_v58).Update(_346_v57, _dafny.UnicodeSeqOfUtf8Bytes("ogexx"))
		var _348_v59 *C5
		_ = _348_v59
		var _nw55 *C5 = New_C5_()
		_ = _nw55
		_nw55.Ctor__(_280_v1, _307_v24, _281_v2)
		_348_v59 = _nw55
		_348_v59 = _348_v59
		_284_v5 = _284_v5
		var _349_v60 _dafny.Set
		_ = _349_v60
		_349_v60 = _dafny.SetOf(_284_v5)
		var _350_v61 D13
		_ = _350_v61
		_350_v61 = Companion_D13_.Create_DC28_(_349_v60)
		_350_v61 = _350_v61
	} else {
		var _351_v62 _dafny.Map
		_ = _351_v62
		_351_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v2, _286_v7)
		var _352_v63 _dafny.Map
		_ = _352_v63
		_352_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_351_v62, _dafny.IntOfUint32((_279_v0).Cardinality()))
		_352_v63 = (_352_v63).Update(Companion_Default___.Fm32(_280_v1, _280_v1, _289_globalState), (_dafny.Zero).Minus((func() _dafny.Int {
			if (_286_v7).Contains(_280_v1) {
				return (_286_v7).Multiplicity(_280_v1)
			}
			return _281_v2
		})()))
		_281_v2 = (_dafny.Zero).Minus(_281_v2)
		(_289_globalState).F19 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(2))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}((func(_353_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_354_i4 _dafny.Int) _dafny.Int {
				return (_353_v2).Times(_353_v2)
			}
		})(_281_v2)))
		_280_v1 = !(_280_v1)
		_280_v1 = (_dafny.IntOfUint32((Companion_Default___.Fm15((_282_v3).Cardinality(), _281_v2, _281_v2, _289_globalState)).Cardinality())).Cmp(_281_v2) < 0
	}
	var _355_v64 D3
	_ = _355_v64
	_355_v64 = Companion_D3_.Create_DC6_(_281_v2, (_dafny.Zero).Minus(_281_v2), _280_v1)
	var _356_v65 _dafny.Map
	_ = _356_v65
	_356_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v2, _281_v2)
	var _357_v66 _dafny.Map
	_ = _357_v66
	_357_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v2, (_356_v65).Cardinality())
	var _358_v67 _dafny.CodePoint
	_ = _358_v67
	var _out7 _dafny.CodePoint
	_ = _out7
	_out7 = Companion_Default___.M6(_281_v2, Companion_Default___.Fm0(_281_v2, Companion_Default___.Fm0((_355_v64).Dtor_cf10(), _280_v1, _280_v1, _281_v2, _289_globalState), _280_v1, _dafny.IntOfUint32((_279_v0).Cardinality()), _289_globalState), _280_v1, ((func() _dafny.Int {
		if (_357_v66).Contains(Companion_Default___.Fm5(_281_v2, Companion_Default___.Fm17(_280_v1, _280_v1, _280_v1, _289_globalState), _289_globalState)) {
			return (_357_v66).Get(Companion_Default___.Fm5(_281_v2, Companion_Default___.Fm17(_280_v1, _280_v1, _280_v1, _289_globalState), _289_globalState)).(_dafny.Int)
		}
		return _dafny.IntOfUint32((_279_v0).Cardinality())
	})()).Cmp(_281_v2) <= 0, _289_globalState)
	_358_v67 = _out7
	var _359_i5 _dafny.Int
	_ = _359_i5
	_359_i5 = _dafny.Zero
	{
		for _280_v1 {
			{
				if (_359_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_359_i5 = (_359_i5).Plus(_dafny.One)
				if (false) || (_280_v1) {
					(_289_globalState).F18 = _280_v1
					var _360_v68 _dafny.Map
					_ = _360_v68
					_360_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v7, _281_v2)
					var _361_v69 *C4
					_ = _361_v69
					var _nw56 *C4 = New_C4_()
					_ = _nw56
					_nw56.Ctor__(_360_v68, !((_280_v1) == (_280_v1)), _dafny.CodePoint('f'), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_281_v2, _281_v2)))
					_361_v69 = _nw56
					var _362_v70 D12
					_ = _362_v70
					_362_v70 = Companion_D12_.Create_DC27_(_281_v2, _280_v1)
					(_289_globalState).F15 = (_281_v2).Minus((_362_v70).Dtor_cf40())
					var _363_v71 D5
					_ = _363_v71
					_363_v71 = Companion_D5_.Create_DC10_((_361_v69).F30(), _280_v1)
					var _pat_let_tv0 = _280_v1
					_ = _pat_let_tv0
					var _364_v72 _dafny.Map
					_ = _364_v72
					_364_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_361_v69).F30(), func(_pat_let0_0 D5) D5 {
						return func(_365_dt__update__tmp_h0 D5) D5 {
							return func(_pat_let1_0 bool) D5 {
								return func(_366_dt__update_hcf14_h0 bool) D5 {
									return Companion_D5_.Create_DC10_(_366_dt__update_hcf14_h0, (_365_dt__update__tmp_h0).Dtor_cf15())
								}(_pat_let1_0)
							}(_pat_let_tv0)
						}(_pat_let0_0)
					}(_363_v71))
					(_361_v69).M5(_364_v72, _280_v1, Companion_Default___.SafeModuloInt(_281_v2, _281_v2), _289_globalState)
					_293_v13 = _293_v13
				} else {
					var _367_v73 *C1
					_ = _367_v73
					var _nw57 *C1 = New_C1_()
					_ = _nw57
					_nw57.Ctor__(_296_v16, (_281_v2).Minus(_dafny.IntOfUint32((_279_v0).Cardinality())))
					_367_v73 = _nw57
					(_289_globalState).F10 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_279_v0, _dafny.UnicodeSeqOfUtf8Bytes("untmsi"))).Cardinality())
					var _368_v74 _dafny.Map
					_ = _368_v74
					_368_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_307_v24, _280_v1)
					var _369_v75 _dafny.Array
					_ = _369_v75
					var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
					_ = _nw58
					_369_v75 = _nw58
					var _370_v76 _dafny.Map
					_ = _370_v76
					_370_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if _280_v1 {
							return (_368_v74).Cardinality()
						}
						return _281_v2
					})(), _369_v75)
					_370_v76 = (_370_v76).Update((_367_v73).Fm6(_289_globalState), _369_v75)
					var _371_v77 T1
					_ = _371_v77
					var _nw59 *C3 = New_C3_()
					_ = _nw59
					_nw59.Ctor__(true, _307_v24, _dafny.IntOfUint32((_287_v8).Cardinality()))
					_371_v77 = _nw59
					var _372_v78 _dafny.Map
					_ = _372_v78
					_372_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v77, _280_v1)
					_280_v1 = (func() bool {
						if _280_v1 {
							return (_372_v78).Contains(_371_v77)
						}
						return ((func() bool {
							if (_306_v23).Contains(_281_v2) {
								return (_306_v23).Get(_281_v2).(bool)
							}
							return _280_v1
						})()) == (Companion_Default___.Fm0(_dafny.IntOfInt64(763), true, _280_v1, _371_v77.F25(), _289_globalState))
					})()
					var _373_v79 _dafny.Map
					_ = _373_v79
					_373_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v7, _371_v77.F25())
					var _374_v81 _dafny.Set
					_ = _374_v81
					_374_v81 = _dafny.SetOf(_dafny.MultiSetOf(!(_280_v1), _280_v1))
					var _375_v82 *C4
					_ = _375_v82
					var _nw60 *C4 = New_C4_()
					_ = _nw60
					_nw60.Ctor__((_373_v79).Merge(func() _dafny.Map {
						var _coll31 = _dafny.NewMapBuilder()
						_ = _coll31
						for _iter31 := _dafny.Iterate((_374_v81).Elements()); ; {
							_compr_31, _ok31 := _iter31()
							if !_ok31 {
								break
							}
							var _376_v80 _dafny.MultiSet
							_376_v80 = interface{}(_compr_31).(_dafny.MultiSet)
							if (_374_v81).Contains(_376_v80) {
								_coll31.Add(_376_v80, _371_v77.F25())
							}
						}
						return _coll31.ToMap()
					}()), _280_v1, _307_v24, _281_v2)
					_375_v82 = _nw60
				}
				var _377_v83 _dafny.Set
				_ = _377_v83
				_377_v83 = _dafny.SetOf(_284_v5, _284_v5, _284_v5)
				var _378_v84 _dafny.Map
				_ = _378_v84
				_378_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_291_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_279_v0).Cardinality()), _dafny.IntOfUint32((_291_v11).Cardinality()))).Uint32()).(bool)) || (_280_v1), _377_v83)
				_378_v84 = (_378_v84).Update(_280_v1, (func() _dafny.Set {
					if _280_v1 {
						return _377_v83
					}
					return _377_v83
				})())
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_283_v4), 0))
				_ = _index43
				(_283_v4).ArraySet1(_280_v1, (_index43).Int())
				var _379_v85 _dafny.Array
				_ = _379_v85
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_9
				var _nw61 _dafny.Array
				_ = _nw61
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw61 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Sequence = (func(_380_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_381_i6 _dafny.Int) _dafny.Sequence {
							return _380_v0
						}
					})(_279_v0)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw61 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw61).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw61).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_379_v85 = _nw61
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_379_v85), 0))
				_ = _index44
				(_379_v85).ArraySet1(_279_v0, (_index44).Int())
				var _382_v86 _dafny.MultiSet
				_ = _382_v86
				_382_v86 = _dafny.MultiSetOf(_358_v67)
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_283_v4), 0))
				_ = _index45
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_379_v85), 0))
				_ = _index46
				var _rhs43 _dafny.Int = Companion_Default___.Fm5(_281_v2, ((_382_v86).Difference(_dafny.MultiSetOf(_dafny.CodePoint('f')))).Cardinality(), _289_globalState)
				_ = _rhs43
				var _rhs44 bool = _280_v1
				_ = _rhs44
				var _rhs45 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_279_v0, _dafny.UnicodeSeqOfUtf8Bytes("wdfjarp"))
				_ = _rhs45
				var _lhs39 *GlobalState = _289_globalState
				_ = _lhs39
				var _lhs40 _dafny.Array = _283_v4
				_ = _lhs40
				var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_283_v4), 0))
				_ = _lhs41
				var _lhs42 _dafny.Array = _379_v85
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_379_v85), 0))
				_ = _lhs43
				_lhs39.F2 = _rhs43
				(_lhs40).ArraySet1(_rhs44, (_lhs41).Int())
				(_lhs42).ArraySet1(_rhs45, (_lhs43).Int())
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_284_v5), 0))
				_ = _index47
				(_284_v5).ArraySet1((_287_v8).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_287_v8).Cardinality()))).Uint32()).(_dafny.Int), (_index47).Int())
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_284_v5), 0))
				_ = _index48
				(_284_v5).ArraySet1((func() _dafny.Int {
					if (_290_v10).Contains(_281_v2) {
						return (_290_v10).Multiplicity(_281_v2)
					}
					return _281_v2
				})(), (_index48).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_281_v2 = _281_v2
	(_289_globalState).F15 = (func() _dafny.Int {
		if _280_v1 {
			return _281_v2
		}
		return _281_v2
	})()
	var _383_v87 _dafny.Map
	_ = _383_v87
	_383_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _284_v5)
	var _384_v88 _dafny.Array
	_ = _384_v88
	var _nwElement0_7 _dafny.Array = _284_v5
	_ = _nwElement0_7
	var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(11))
	_ = _nw62
	(_nw62).ArraySet1(_nwElement0_7, 0)
	(_nw62).ArraySet1(_284_v5, 1)
	(_nw62).ArraySet1(_284_v5, 2)
	(_nw62).ArraySet1(_284_v5, 3)
	(_nw62).ArraySet1((func() _dafny.Array {
		if (_383_v87).Contains(_280_v1) {
			return (_383_v87).Get(_280_v1).(_dafny.Array)
		}
		return _284_v5
	})(), 4)
	(_nw62).ArraySet1(_284_v5, 5)
	(_nw62).ArraySet1(_284_v5, 6)
	(_nw62).ArraySet1(_284_v5, 7)
	(_nw62).ArraySet1(_284_v5, 8)
	(_nw62).ArraySet1(_284_v5, 9)
	(_nw62).ArraySet1(_284_v5, 10)
	_384_v88 = _nw62
	_384_v88 = _384_v88
	var _385_v89 _dafny.Array
	_ = _385_v89
	var _len0_10 _dafny.Int = _dafny.IntOfInt64(14)
	_ = _len0_10
	var _nw63 _dafny.Array
	_ = _nw63
	if _len0_10.Cmp(_dafny.Zero) == 0 {
		_nw63 = _dafny.NewArray(_len0_10)
	} else {
		var _init10 func(_dafny.Int) _dafny.Sequence = func(_386_i8 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(903))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_387_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			}))
		}
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
	_385_v89 = _nw63
	for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_385_v89), 0))); ; {
		_guard_loop_0, _ok32 := _iter32()
		if !_ok32 {
			break
		}
		var _388_i7 _dafny.Int
		_388_i7 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_388_i7).Sign() != -1) && ((_388_i7).Cmp(_dafny.ArrayLen((_385_v89), 0)) < 0)) {
			(_385_v89).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sn"), _279_v0), _dafny.UnicodeSeqOfUtf8Bytes("infvwkvh")), (_388_i7).Int())
		}
	}
	var _389_v90 _dafny.Map
	_ = _389_v90
	_389_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_291_v11).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm17(!(_280_v1), _280_v1, _280_v1, _289_globalState), _dafny.IntOfUint32((_291_v11).Cardinality()))).Uint32()).(bool), _280_v1)
	if !((func() bool {
		if (_389_v90).Contains(_280_v1) {
			return (_389_v90).Get(_280_v1).(bool)
		}
		return (_281_v2).Cmp((_dafny.Zero).Minus(_281_v2)) > 0
	})()) {
		var _390_v91 *C6
		_ = _390_v91
		var _nw64 *C6 = New_C6_()
		_ = _nw64
		_nw64.Ctor__((func() _dafny.Map {
			if _280_v1 {
				return (_306_v23).Update(_281_v2, false)
			}
			return _306_v23
		})())
		_390_v91 = _nw64
		var _391_v92 *C3
		_ = _391_v92
		var _nw65 *C3 = New_C3_()
		_ = _nw65
		_nw65.Ctor__(_280_v1, _358_v67, _dafny.IntOfInt64(224))
		_391_v92 = _nw65
		var _392_v93 *C1
		_ = _392_v93
		var _nw66 *C1 = New_C1_()
		_ = _nw66
		_nw66.Ctor__(_358_v67, (_287_v8).Select((Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_287_v8).Cardinality()))).Uint32()).(_dafny.Int))
		_392_v93 = _nw66
		_392_v93 = (func() *C1 {
			if !(_391_v92.F31) {
				return _392_v93
			}
			return _392_v93
		})()
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_284_v5), 0))
		_ = _index49
		(_284_v5).ArraySet1(_281_v2, (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_284_v5), 0))
		_ = _index50
		(_284_v5).ArraySet1((func() _dafny.Int {
			if (_356_v65).Contains(_281_v2) {
				return (_356_v65).Get(_281_v2).(_dafny.Int)
			}
			return Companion_Default___.SafeDivisionInt(_281_v2, _dafny.IntOfUint32((_291_v11).Cardinality()))
		})(), (_index50).Int())
		(_289_globalState).F15 = (((_284_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_284_v5), 0))).Int()).(_dafny.Int)).Minus(_281_v2)).Minus(_281_v2)
	} else {
		var _393_v94 D7
		_ = _393_v94
		_393_v94 = Companion_D7_.Create_DC13_(_306_v23)
		var _394_v95 *C1
		_ = _394_v95
		var _nw67 *C1 = New_C1_()
		_ = _nw67
		_nw67.Ctor__(_307_v24, (((_393_v94).Dtor_cf19()).Cardinality()).Minus(_281_v2))
		_394_v95 = _nw67
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_284_v5), 0))
		_ = _index51
		(_284_v5).ArraySet1(_281_v2, (_index51).Int())
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_385_v89), 0))
		_ = _index52
		(_385_v89).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yj"), _279_v0), (_index52).Int())
		var _395_v96 D10
		_ = _395_v96
		_395_v96 = Companion_D10_.Create_DC21_((func() _dafny.Int {
			if (_282_v3).Contains(_280_v1) {
				return (_282_v3).Get(_280_v1).(_dafny.Int)
			}
			return _281_v2
		})(), _280_v1, !(false))
		var _396_v97 _dafny.MultiSet
		_ = _396_v97
		_396_v97 = _dafny.MultiSetOf(_395_v96, _395_v96, _395_v96, _395_v96, Companion_D10_.Create_DC21_(_281_v2, !(Companion_Default___.Fm0((_dafny.MultiSetFromSeq(_291_v11)).Cardinality(), _280_v1, _280_v1, _dafny.IntOfUint32((_279_v0).Cardinality()), _289_globalState)), _280_v1))
		var _pat_let_tv1 = _280_v1
		_ = _pat_let_tv1
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_284_v5), 0))
		_ = _index53
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_385_v89), 0))
		_ = _index54
		var _rhs46 _dafny.Int = _281_v2
		_ = _rhs46
		var _rhs47 _dafny.Sequence = _279_v0
		_ = _rhs47
		var _rhs48 _dafny.Set = (_305_v22).Union(_305_v22)
		_ = _rhs48
		var _rhs49 D10 = func(_pat_let2_0 D10) D10 {
			return func(_397_dt__update__tmp_h1 D10) D10 {
				return func(_pat_let3_0 bool) D10 {
					return func(_398_dt__update_hcf32_h0 bool) D10 {
						return Companion_D10_.Create_DC21_((_397_dt__update__tmp_h1).Dtor_cf30(), (_397_dt__update__tmp_h1).Dtor_cf31(), _398_dt__update_hcf32_h0)
					}(_pat_let3_0)
				}(_pat_let_tv1)
			}(_pat_let2_0)
		}(_395_v96)
		_ = _rhs49
		var _rhs50 _dafny.MultiSet = ((func() _dafny.MultiSet {
			if _280_v1 {
				return _396_v97
			}
			return _396_v97
		})()).Intersection((_dafny.MultiSetOf(_395_v96)).Union(_396_v97))
		_ = _rhs50
		var _lhs44 _dafny.Array = _284_v5
		_ = _lhs44
		var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_284_v5), 0))
		_ = _lhs45
		var _lhs46 _dafny.Array = _385_v89
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_385_v89), 0))
		_ = _lhs47
		(_lhs44).ArraySet1(_rhs46, (_lhs45).Int())
		(_lhs46).ArraySet1(_rhs47, (_lhs47).Int())
		_305_v22 = _rhs48
		_395_v96 = _rhs49
		_396_v97 = _rhs50
		var _399_v98 _dafny.Map
		_ = _399_v98
		_399_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v7, _dafny.IntOfInt64(683))
		var _400_v99 T1
		_ = _400_v99
		var _nw68 *C4 = New_C4_()
		_ = _nw68
		_nw68.Ctor__(_399_v98, _280_v1, _296_v16, _dafny.IntOfInt64(350))
		_400_v99 = _nw68
		var _401_v100 _dafny.Map
		_ = _401_v100
		_401_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(482), _400_v99)
		_401_v100 = (_401_v100).Update((func() _dafny.Int {
			if _280_v1 {
				return (_282_v3).Cardinality()
			}
			return _281_v2
		})(), _400_v99)
		var _402_v101 _dafny.Map
		_ = _402_v101
		_402_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("b"), (Companion_Default___.SafeIndex(_281_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality()))).Uint32(), _296_v16), _281_v2)
		var _403_v102 _dafny.Map
		_ = _403_v102
		_403_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_394_v95).Fm6(_289_globalState), _402_v101)
		_403_v102 = (_403_v102).Update(_400_v99.F25(), Companion_Default___.Fm33(_289_globalState))
		_279_v0 = (_385_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_385_v89), 0))).Int()).(_dafny.Sequence)
	}
	for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_284_v5), 0))); ; {
		_guard_loop_1, _ok33 := _iter33()
		if !_ok33 {
			break
		}
		var _404_i10 _dafny.Int
		_404_i10 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_404_i10).Sign() != -1) && ((_404_i10).Cmp(_dafny.ArrayLen((_284_v5), 0)) < 0)) {
			(_284_v5).ArraySet1(Companion_Default___.SafeModuloInt(_404_i10, _281_v2), (_404_i10).Int())
		}
	}
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_279_v0, _dafny.SeqOf(_dafny.CodePoint('w'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_280_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_281_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_282_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(742))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_283_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_283_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v5).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_285_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v7).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_287_v8, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_288_v9).Equals(_dafny.SetOf(_dafny.SeqOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_289_globalState.F3, _dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('w'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_globalState).F7()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(742))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_globalState).F8()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_globalState.F11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState.F13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(742))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_289_globalState).F14()).Equals(_dafny.SetOf(_dafny.SeqOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_globalState.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_289_globalState.F19, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_globalState.F22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v10).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(742), _dafny.IntOfInt64(742))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_291_v11, _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_292_v12), _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_293_v13).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_294_v14), _dafny.SeqOf(_dafny.CodePoint('w'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_295_v15, _dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('w')), _dafny.SeqOf(_dafny.CodePoint('w')), _dafny.SeqOf(_dafny.CodePoint('w')), _dafny.UnicodeSeqOfUtf8Bytes("fln"), _dafny.SeqOf(_dafny.CodePoint('w')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_296_v16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_297_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_305_v22).Equals(_dafny.SetOf(_dafny.IntOfInt64(742), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v23).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_307_v24)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v25).Dtor_cf34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v25).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_309_v26).Dtor_cf38()).Dtor_cf34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_309_v26).Dtor_cf38()).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_310_v27).ArrayGet1((_dafny.Zero).Int()).(D11)).Dtor_cf38()).Dtor_cf34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_310_v27).ArrayGet1((_dafny.Zero).Int()).(D11)).Dtor_cf38()).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_310_v27).ArrayGet1((_dafny.One).Int()).(D11)).Dtor_cf38()).Dtor_cf34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_310_v27).ArrayGet1((_dafny.One).Int()).(D11)).Dtor_cf38()).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_310_v27).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D11)).Dtor_cf38()).Dtor_cf34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_310_v27).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D11)).Dtor_cf38()).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_310_v27).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D11)).Dtor_cf38()).Dtor_cf34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_310_v27).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D11)).Dtor_cf38()).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_311_v28).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_312_v29).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_355_v64).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_355_v64).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_355_v64).Dtor_cf11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_356_v65).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(742), _dafny.IntOfInt64(742))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_357_v66).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(742), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_358_v67)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_359_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_383_v87).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_384_v88).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_385_v89).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('w'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v89).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v90).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
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

type D1_DC1 struct {
	Cf1 _dafny.Sequence
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Sequence) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D1) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC1:
		{
			return "D1.DC1" + "(" + data.Cf1.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
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

type D2_DC3 struct {
	Cf3 _dafny.Int
	Cf4 _dafny.Int
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf3 _dafny.Int, Cf4 _dafny.Int) D2 {
	return D2{D2_DC3{Cf3, Cf4}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

type D2_DC2 struct {
	Cf2 bool
}

func (D2_DC2) isD2() {}

func (CompanionStruct_D2_) Create_DC2_(Cf2 bool) D2 {
	return D2{D2_DC2{Cf2}}
}

func (_this D2) Is_DC2() bool {
	_, ok := _this.Get_().(D2_DC2)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC3_(_dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D2_DC3).Cf3
}

func (_this D2) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D2_DC3).Cf4
}

func (_this D2) Dtor_cf2() bool {
	return _this.Get_().(D2_DC2).Cf2
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D2_DC2:
		{
			return "D2.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D2_DC2:
		{
			data2, ok := other.Get_().(D2_DC2)
			return ok && data1.Cf2 == data2.Cf2
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

type D3_DC5 struct {
	Cf6 _dafny.Sequence
	Cf7 _dafny.Int
	Cf8 _dafny.Int
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf6 _dafny.Sequence, Cf7 _dafny.Int, Cf8 _dafny.Int) D3 {
	return D3{D3_DC5{Cf6, Cf7, Cf8}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

type D3_DC6 struct {
	Cf9  _dafny.Int
	Cf10 _dafny.Int
	Cf11 bool
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf9 _dafny.Int, Cf10 _dafny.Int, Cf11 bool) D3 {
	return D3{D3_DC6{Cf9, Cf10, Cf11}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC4 struct {
	Cf5 _dafny.Set
}

func (D3_DC4) isD3() {}

func (CompanionStruct_D3_) Create_DC4_(Cf5 _dafny.Set) D3 {
	return D3{D3_DC4{Cf5}}
}

func (_this D3) Is_DC4() bool {
	_, ok := _this.Get_().(D3_DC4)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC5_(_dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this D3) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D3_DC5).Cf6
}

func (_this D3) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D3_DC5).Cf7
}

func (_this D3) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D3_DC5).Cf8
}

func (_this D3) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf9
}

func (_this D3) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf10
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC6).Cf11
}

func (_this D3) Dtor_cf5() _dafny.Set {
	return _this.Get_().(D3_DC4).Cf5
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC4:
		{
			return "D3.DC4" + "(" + _dafny.String(data.Cf5) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11 == data2.Cf11
		}
	case D3_DC4:
		{
			data2, ok := other.Get_().(D3_DC4)
			return ok && data1.Cf5.Equals(data2.Cf5)
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

type D4_DC8 struct {
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_() D4 {
	return D4{D4_DC8{}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC7 struct {
	Cf12 _dafny.Array
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf12 _dafny.Array) D4 {
	return D4{D4_DC7{Cf12}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC8_()
}

func (_this D4) Dtor_cf12() _dafny.Array {
	return _this.Get_().(D4_DC7).Cf12
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8"
		}
	case D4_DC7:
		{
			return "D4.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC8:
		{
			_, ok := other.Get_().(D4_DC8)
			return ok
		}
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf12 == data2.Cf12
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

type D5_DC10 struct {
	Cf14 bool
	Cf15 bool
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf14 bool, Cf15 bool) D5 {
	return D5{D5_DC10{Cf14, Cf15}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

type D5_DC9 struct {
	Cf13 _dafny.Array
}

func (D5_DC9) isD5() {}

func (CompanionStruct_D5_) Create_DC9_(Cf13 _dafny.Array) D5 {
	return D5{D5_DC9{Cf13}}
}

func (_this D5) Is_DC9() bool {
	_, ok := _this.Get_().(D5_DC9)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC10_(false, false)
}

func (_this D5) Dtor_cf14() bool {
	return _this.Get_().(D5_DC10).Cf14
}

func (_this D5) Dtor_cf15() bool {
	return _this.Get_().(D5_DC10).Cf15
}

func (_this D5) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D5_DC9).Cf13
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D5_DC9:
		{
			return "D5.DC9" + "(" + _dafny.String(data.Cf13) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15
		}
	case D5_DC9:
		{
			data2, ok := other.Get_().(D5_DC9)
			return ok && data1.Cf13 == data2.Cf13
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

type D6_DC12 struct {
	Cf17 bool
	Cf18 _dafny.Int
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf17 bool, Cf18 _dafny.Int) D6 {
	return D6{D6_DC12{Cf17, Cf18}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC11 struct {
	Cf16 T0
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf16 T0) D6 {
	return D6{D6_DC11{Cf16}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC12_(false, _dafny.Zero)
}

func (_this D6) Dtor_cf17() bool {
	return _this.Get_().(D6_DC12).Cf17
}

func (_this D6) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D6_DC12).Cf18
}

func (_this D6) Dtor_cf16() T0 {
	return _this.Get_().(D6_DC11).Cf16
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D6_DC11:
		{
			return "D6.DC11" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0
		}
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
			return ok && _dafny.AreEqual(data1.Cf16, data2.Cf16)
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

type D7_DC14 struct {
	Cf20 bool
	Cf21 _dafny.Int
	Cf22 _dafny.Int
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf20 bool, Cf21 _dafny.Int, Cf22 _dafny.Int) D7 {
	return D7{D7_DC14{Cf20, Cf21, Cf22}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

type D7_DC13 struct {
	Cf19 _dafny.Map
}

func (D7_DC13) isD7() {}

func (CompanionStruct_D7_) Create_DC13_(Cf19 _dafny.Map) D7 {
	return D7{D7_DC13{Cf19}}
}

func (_this D7) Is_DC13() bool {
	_, ok := _this.Get_().(D7_DC13)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC14_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D7) Dtor_cf20() bool {
	return _this.Get_().(D7_DC14).Cf20
}

func (_this D7) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D7_DC14).Cf21
}

func (_this D7) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D7_DC14).Cf22
}

func (_this D7) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D7_DC13).Cf19
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D7_DC13:
		{
			return "D7.DC13" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D7_DC13:
		{
			data2, ok := other.Get_().(D7_DC13)
			return ok && data1.Cf19.Equals(data2.Cf19)
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

type D8_DC15 struct {
	Cf23 _dafny.Sequence
}

func (D8_DC15) isD8() {}

func (CompanionStruct_D8_) Create_DC15_(Cf23 _dafny.Sequence) D8 {
	return D8{D8_DC15{Cf23}}
}

func (_this D8) Is_DC15() bool {
	_, ok := _this.Get_().(D8_DC15)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D8_DC15).Cf23
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC15:
		{
			return "D8.DC15" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC15:
		{
			data2, ok := other.Get_().(D8_DC15)
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D9_DC17 struct {
	Cf25 _dafny.Int
}

func (D9_DC17) isD9() {}

func (CompanionStruct_D9_) Create_DC17_(Cf25 _dafny.Int) D9 {
	return D9{D9_DC17{Cf25}}
}

func (_this D9) Is_DC17() bool {
	_, ok := _this.Get_().(D9_DC17)
	return ok
}

type D9_DC16 struct {
	Cf24 _dafny.Map
}

func (D9_DC16) isD9() {}

func (CompanionStruct_D9_) Create_DC16_(Cf24 _dafny.Map) D9 {
	return D9{D9_DC16{Cf24}}
}

func (_this D9) Is_DC16() bool {
	_, ok := _this.Get_().(D9_DC16)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC17_(_dafny.Zero)
}

func (_this D9) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D9_DC17).Cf25
}

func (_this D9) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D9_DC16).Cf24
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC17:
		{
			return "D9.DC17" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D9_DC16:
		{
			return "D9.DC16" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC17:
		{
			data2, ok := other.Get_().(D9_DC17)
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D9_DC16:
		{
			data2, ok := other.Get_().(D9_DC16)
			return ok && data1.Cf24.Equals(data2.Cf24)
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

type D10_DC19 struct {
	Cf27 _dafny.CodePoint
}

func (D10_DC19) isD10() {}

func (CompanionStruct_D10_) Create_DC19_(Cf27 _dafny.CodePoint) D10 {
	return D10{D10_DC19{Cf27}}
}

func (_this D10) Is_DC19() bool {
	_, ok := _this.Get_().(D10_DC19)
	return ok
}

type D10_DC20 struct {
	Cf28 bool
	Cf29 _dafny.Int
}

func (D10_DC20) isD10() {}

func (CompanionStruct_D10_) Create_DC20_(Cf28 bool, Cf29 _dafny.Int) D10 {
	return D10{D10_DC20{Cf28, Cf29}}
}

func (_this D10) Is_DC20() bool {
	_, ok := _this.Get_().(D10_DC20)
	return ok
}

type D10_DC21 struct {
	Cf30 _dafny.Int
	Cf31 bool
	Cf32 bool
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf30 _dafny.Int, Cf31 bool, Cf32 bool) D10 {
	return D10{D10_DC21{Cf30, Cf31, Cf32}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

type D10_DC18 struct {
	Cf26 _dafny.Set
}

func (D10_DC18) isD10() {}

func (CompanionStruct_D10_) Create_DC18_(Cf26 _dafny.Set) D10 {
	return D10{D10_DC18{Cf26}}
}

func (_this D10) Is_DC18() bool {
	_, ok := _this.Get_().(D10_DC18)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC19_(_dafny.CodePoint('D'))
}

func (_this D10) Dtor_cf27() _dafny.CodePoint {
	return _this.Get_().(D10_DC19).Cf27
}

func (_this D10) Dtor_cf28() bool {
	return _this.Get_().(D10_DC20).Cf28
}

func (_this D10) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D10_DC20).Cf29
}

func (_this D10) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D10_DC21).Cf30
}

func (_this D10) Dtor_cf31() bool {
	return _this.Get_().(D10_DC21).Cf31
}

func (_this D10) Dtor_cf32() bool {
	return _this.Get_().(D10_DC21).Cf32
}

func (_this D10) Dtor_cf26() _dafny.Set {
	return _this.Get_().(D10_DC18).Cf26
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC19:
		{
			return "D10.DC19" + "(" + _dafny.String(data.Cf27) + ")"
		}
	case D10_DC20:
		{
			return "D10.DC20" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D10_DC18:
		{
			return "D10.DC18" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC19:
		{
			data2, ok := other.Get_().(D10_DC19)
			return ok && data1.Cf27 == data2.Cf27
		}
	case D10_DC20:
		{
			data2, ok := other.Get_().(D10_DC20)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32
		}
	case D10_DC18:
		{
			data2, ok := other.Get_().(D10_DC18)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D11_DC23 struct {
	Cf34 _dafny.Int
	Cf35 bool
}

func (D11_DC23) isD11() {}

func (CompanionStruct_D11_) Create_DC23_(Cf34 _dafny.Int, Cf35 bool) D11 {
	return D11{D11_DC23{Cf34, Cf35}}
}

func (_this D11) Is_DC23() bool {
	_, ok := _this.Get_().(D11_DC23)
	return ok
}

type D11_DC24 struct {
	Cf36 bool
	Cf37 _dafny.Int
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf36 bool, Cf37 _dafny.Int) D11 {
	return D11{D11_DC24{Cf36, Cf37}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

type D11_DC22 struct {
	Cf33 _dafny.Array
}

func (D11_DC22) isD11() {}

func (CompanionStruct_D11_) Create_DC22_(Cf33 _dafny.Array) D11 {
	return D11{D11_DC22{Cf33}}
}

func (_this D11) Is_DC22() bool {
	_, ok := _this.Get_().(D11_DC22)
	return ok
}

type D11_DC25 struct {
	Cf38 D11
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf38 D11) D11 {
	return D11{D11_DC25{Cf38}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC23_(_dafny.Zero, false)
}

func (_this D11) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D11_DC23).Cf34
}

func (_this D11) Dtor_cf35() bool {
	return _this.Get_().(D11_DC23).Cf35
}

func (_this D11) Dtor_cf36() bool {
	return _this.Get_().(D11_DC24).Cf36
}

func (_this D11) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D11_DC24).Cf37
}

func (_this D11) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D11_DC22).Cf33
}

func (_this D11) Dtor_cf38() D11 {
	return _this.Get_().(D11_DC25).Cf38
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC23:
		{
			return "D11.DC23" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D11_DC22:
		{
			return "D11.DC22" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC23:
		{
			data2, ok := other.Get_().(D11_DC23)
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35
		}
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D11_DC22:
		{
			data2, ok := other.Get_().(D11_DC22)
			return ok && data1.Cf33 == data2.Cf33
		}
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf38.Equals(data2.Cf38)
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
	Cf40 _dafny.Int
	Cf41 bool
}

func (D12_DC27) isD12() {}

func (CompanionStruct_D12_) Create_DC27_(Cf40 _dafny.Int, Cf41 bool) D12 {
	return D12{D12_DC27{Cf40, Cf41}}
}

func (_this D12) Is_DC27() bool {
	_, ok := _this.Get_().(D12_DC27)
	return ok
}

type D12_DC26 struct {
	Cf39 _dafny.Set
}

func (D12_DC26) isD12() {}

func (CompanionStruct_D12_) Create_DC26_(Cf39 _dafny.Set) D12 {
	return D12{D12_DC26{Cf39}}
}

func (_this D12) Is_DC26() bool {
	_, ok := _this.Get_().(D12_DC26)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC27_(_dafny.Zero, false)
}

func (_this D12) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D12_DC27).Cf40
}

func (_this D12) Dtor_cf41() bool {
	return _this.Get_().(D12_DC27).Cf41
}

func (_this D12) Dtor_cf39() _dafny.Set {
	return _this.Get_().(D12_DC26).Cf39
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC27:
		{
			return "D12.DC27" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D12_DC26:
		{
			return "D12.DC26" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41 == data2.Cf41
		}
	case D12_DC26:
		{
			data2, ok := other.Get_().(D12_DC26)
			return ok && data1.Cf39.Equals(data2.Cf39)
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

type D13_DC29 struct {
	Cf43 bool
	Cf44 bool
	Cf45 _dafny.Array
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf43 bool, Cf44 bool, Cf45 _dafny.Array) D13 {
	return D13{D13_DC29{Cf43, Cf44, Cf45}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

type D13_DC28 struct {
	Cf42 _dafny.Set
}

func (D13_DC28) isD13() {}

func (CompanionStruct_D13_) Create_DC28_(Cf42 _dafny.Set) D13 {
	return D13{D13_DC28{Cf42}}
}

func (_this D13) Is_DC28() bool {
	_, ok := _this.Get_().(D13_DC28)
	return ok
}

type D13_DC30 struct {
	Cf46 D13
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf46 D13) D13 {
	return D13{D13_DC30{Cf46}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC29_(false, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D13) Dtor_cf43() bool {
	return _this.Get_().(D13_DC29).Cf43
}

func (_this D13) Dtor_cf44() bool {
	return _this.Get_().(D13_DC29).Cf44
}

func (_this D13) Dtor_cf45() _dafny.Array {
	return _this.Get_().(D13_DC29).Cf45
}

func (_this D13) Dtor_cf42() _dafny.Set {
	return _this.Get_().(D13_DC28).Cf42
}

func (_this D13) Dtor_cf46() D13 {
	return _this.Get_().(D13_DC30).Cf46
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ")"
		}
	case D13_DC28:
		{
			return "D13.DC28" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC29:
		{
			data2, ok := other.Get_().(D13_DC29)
			return ok && data1.Cf43 == data2.Cf43 && data1.Cf44 == data2.Cf44 && data1.Cf45 == data2.Cf45
		}
	case D13_DC28:
		{
			data2, ok := other.Get_().(D13_DC28)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D13_DC30:
		{
			data2, ok := other.Get_().(D13_DC30)
			return ok && data1.Cf46.Equals(data2.Cf46)
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

type D14_DC31 struct {
	Cf47 _dafny.Sequence
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf47 _dafny.Sequence) D14 {
	return D14{D14_DC31{Cf47}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D14) Dtor_cf47() _dafny.Sequence {
	return _this.Get_().(D14_DC31).Cf47
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf47) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC31:
		{
			data2, ok := other.Get_().(D14_DC31)
			return ok && data1.Cf47.Equals(data2.Cf47)
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
	Cf48 *C3
}

func (D15_DC32) isD15() {}

func (CompanionStruct_D15_) Create_DC32_(Cf48 *C3) D15 {
	return D15{D15_DC32{Cf48}}
}

func (_this D15) Is_DC32() bool {
	_, ok := _this.Get_().(D15_DC32)
	return ok
}

func (CompanionStruct_D15_) Default() *C3 {
	return (*C3)(nil)
}

func (_this D15) Dtor_cf48() *C3 {
	return _this.Get_().(D15_DC32).Cf48
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC32:
		{
			return "D15.DC32" + "(" + _dafny.String(data.Cf48) + ")"
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
			return ok && data1.Cf48 == data2.Cf48
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

type D16_DC34 struct {
	Cf50 bool
}

func (D16_DC34) isD16() {}

func (CompanionStruct_D16_) Create_DC34_(Cf50 bool) D16 {
	return D16{D16_DC34{Cf50}}
}

func (_this D16) Is_DC34() bool {
	_, ok := _this.Get_().(D16_DC34)
	return ok
}

type D16_DC33 struct {
	Cf49 _dafny.Map
}

func (D16_DC33) isD16() {}

func (CompanionStruct_D16_) Create_DC33_(Cf49 _dafny.Map) D16 {
	return D16{D16_DC33{Cf49}}
}

func (_this D16) Is_DC33() bool {
	_, ok := _this.Get_().(D16_DC33)
	return ok
}

type D16_DC35 struct {
	Cf51 D16
}

func (D16_DC35) isD16() {}

func (CompanionStruct_D16_) Create_DC35_(Cf51 D16) D16 {
	return D16{D16_DC35{Cf51}}
}

func (_this D16) Is_DC35() bool {
	_, ok := _this.Get_().(D16_DC35)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC34_(false)
}

func (_this D16) Dtor_cf50() bool {
	return _this.Get_().(D16_DC34).Cf50
}

func (_this D16) Dtor_cf49() _dafny.Map {
	return _this.Get_().(D16_DC33).Cf49
}

func (_this D16) Dtor_cf51() D16 {
	return _this.Get_().(D16_DC35).Cf51
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC34:
		{
			return "D16.DC34" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D16_DC33:
		{
			return "D16.DC33" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D16_DC35:
		{
			return "D16.DC35" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC34:
		{
			data2, ok := other.Get_().(D16_DC34)
			return ok && data1.Cf50 == data2.Cf50
		}
	case D16_DC33:
		{
			data2, ok := other.Get_().(D16_DC33)
			return ok && data1.Cf49.Equals(data2.Cf49)
		}
	case D16_DC35:
		{
			data2, ok := other.Get_().(D16_DC35)
			return ok && data1.Cf51.Equals(data2.Cf51)
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
	F25() _dafny.Int
	F25_set_(value _dafny.Int)
	Fm3(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence
	F24() _dafny.CodePoint
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
	F25() _dafny.Int
	F25_set_(value _dafny.Int)
	Fm3(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence
	F24() _dafny.CodePoint
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
	F0   _dafny.Int
	F2   _dafny.Int
	F3   _dafny.Sequence
	F10  _dafny.Int
	F11  _dafny.Int
	F13  _dafny.Map
	F15  _dafny.Int
	F16  _dafny.Int
	F18  bool
	F19  _dafny.Sequence
	F22  _dafny.Int
	_f1  _dafny.Int
	_f4  _dafny.Int
	_f5  bool
	_f6  _dafny.CodePoint
	_f7  _dafny.Map
	_f8  _dafny.Map
	_f9  _dafny.Int
	_f12 bool
	_f14 _dafny.Set
	_f17 bool
	_f20 _dafny.Int
	_f21 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F2 = _dafny.Zero
	_this.F3 = _dafny.EmptySeq
	_this.F10 = _dafny.Zero
	_this.F11 = _dafny.Zero
	_this.F13 = _dafny.EmptyMap
	_this.F15 = _dafny.Zero
	_this.F16 = _dafny.Zero
	_this.F18 = false
	_this.F19 = _dafny.EmptySeq
	_this.F22 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f4 = _dafny.Zero
	_this._f5 = false
	_this._f6 = _dafny.CodePoint('D')
	_this._f7 = _dafny.EmptyMap
	_this._f8 = _dafny.EmptyMap
	_this._f9 = _dafny.Zero
	_this._f12 = false
	_this._f14 = _dafny.EmptySet
	_this._f17 = false
	_this._f20 = _dafny.Zero
	_this._f21 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Sequence, f4 _dafny.Int, f5 bool, f6 _dafny.CodePoint, f7 _dafny.Map, f8 _dafny.Map, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Int, f12 bool, f13 _dafny.Map, f14 _dafny.Set, f15 _dafny.Int, f16 _dafny.Int, f17 bool, f18 bool, f19 _dafny.Sequence, f20 _dafny.Int, f21 bool, f22 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this).F16 = f16
		(_this)._f17 = f17
		(_this).F18 = f18
		(_this).F19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this).F22 = f22
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() bool {
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
func (_this *GlobalState) F8() _dafny.Map {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Set {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F20() _dafny.Int {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() bool {
	{
		return _this._f21
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
	_f25 _dafny.Int
	_f24 _dafny.CodePoint
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f25 = _dafny.Zero
	_this._f24 = _dafny.CodePoint('D')
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

func (_this *C1) F25() _dafny.Int {
	return _this._f25
}
func (_this *C1) F25_set_(value _dafny.Int) {
	_this._f25 = value
}
func (_this *C1) F24() _dafny.CodePoint {
	return _this._f24
}
func (_this *C1) Ctor__(f24 _dafny.CodePoint, f25 _dafny.Int) {
	{
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C1) Fm3(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ywk"), _dafny.UnicodeSeqOfUtf8Bytes("nyrhpcet"))
	}
}
func (_this *C1) Fm6(globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_this.F25(), _this.F25()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, true), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-385))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg36 _dafny.Int) interface{} {
				return coer36(arg36)
			}
		}(func(_405_i0 _dafny.Int) _dafny.Int {
			return _this.F25()
		}))).Cardinality()))).Cardinality()))).Cardinality())
	}
}
func (_this *C1) M4(p0 bool, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _406_v0 _dafny.Map
		_ = _406_v0
		_406_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _407_v1 _dafny.Array
		_ = _407_v1
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_11
		var _nw69 _dafny.Array
		_ = _nw69
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw69 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) bool = (func(_408_p0 bool) func(_dafny.Int) bool {
				return func(_409_i0 _dafny.Int) bool {
					return _408_p0
				}
			})(p0)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw69 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw69).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw69).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_407_v1 = _nw69
		var _410_v2 _dafny.Set
		_ = _410_v2
		_410_v2 = _dafny.SetOf(_407_v1)
		_406_v0 = (_406_v0).Update((_410_v2).IsSubsetOf(_410_v2), p0)
		(_this).F25_set_(p1)
		var _hi1 _dafny.Int = (_this).Fm6(globalState)
		_ = _hi1
		for _411_i1 := p1; _411_i1.Cmp(_hi1) < 0; _411_i1 = _411_i1.Plus(_dafny.One) {
			var _412_v3 _dafny.Array
			_ = _412_v3
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
			_ = _nw70
			_412_v3 = _nw70
			_412_v3 = _412_v3
			var _413_v4 D5
			_ = _413_v4
			_413_v4 = Companion_D5_.Create_DC10_(p0, !(p0))
			var _source7 D5 = _413_v4
			_ = _source7
			if _source7.Is_DC10() {
				var _414___mcc_h0 bool = _source7.Get_().(D5_DC10).Cf14
				_ = _414___mcc_h0
				var _415___mcc_h1 bool = _source7.Get_().(D5_DC10).Cf15
				_ = _415___mcc_h1
				var _416_cf15 bool = _415___mcc_h1
				_ = _416_cf15
				var _417_cf14 bool = _414___mcc_h0
				_ = _417_cf14
				var _418_v5 _dafny.Array
				_ = _418_v5
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw71
				_418_v5 = _nw71
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_418_v5), 0))
				_ = _index55
				(_418_v5).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqOf(_416_cf15))).Cardinality(), (_index55).Int())
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_407_v1), 0))
				_ = _index56
				(_407_v1).ArraySet1(_417_cf14, (_index56).Int())
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_418_v5), 0))
				_ = _index57
				(_418_v5).ArraySet1(_411_i1, (_index57).Int())
				var _419_v6 _dafny.Map
				_ = _419_v6
				_419_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), p0)
				var _420_v7 _dafny.Sequence
				_ = _420_v7
				_420_v7 = _dafny.UnicodeSeqOfUtf8Bytes("tpjon")
				var _421_v8 _dafny.Sequence
				_ = _421_v8
				_421_v8 = _dafny.SeqOf(_416_cf15, p0)
				var _422_v10 _dafny.Sequence
				_ = _422_v10
				_422_v10 = _dafny.SeqOf(_this.F25())
				var _423_v11 _dafny.Set
				_ = _423_v11
				_423_v11 = _dafny.SetOf(p0)
				var _424_v12 D7
				_ = _424_v12
				_424_v12 = Companion_D7_.Create_DC13_(_419_v6)
				var _425_v14 _dafny.Array
				_ = _425_v14
				var _nwElement0_8 _dafny.Map = _419_v6
				_ = _nwElement0_8
				var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(28))
				_ = _nw72
				(_nw72).ArraySet1(_nwElement0_8, 0)
				(_nw72).ArraySet1(_419_v6, 1)
				(_nw72).ArraySet1(_419_v6, 2)
				(_nw72).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_420_v7).Cardinality()), (_421_v8).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_421_v8).Cardinality()))).Uint32()).(bool)), 3)
				(_nw72).ArraySet1(_419_v6, 4)
				(_nw72).ArraySet1(_419_v6, 5)
				(_nw72).ArraySet1(_419_v6, 6)
				(_nw72).ArraySet1(_419_v6, 7)
				(_nw72).ArraySet1(_419_v6, 8)
				(_nw72).ArraySet1(_419_v6, 9)
				(_nw72).ArraySet1(_419_v6, 10)
				(_nw72).ArraySet1(_419_v6, 11)
				(_nw72).ArraySet1(_419_v6, 12)
				(_nw72).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_420_v7).Cardinality()), _417_cf14), 13)
				(_nw72).ArraySet1(_419_v6, 14)
				(_nw72).ArraySet1(_419_v6, 15)
				(_nw72).ArraySet1((_419_v6).Update((_this).Fm6(globalState), _416_cf15), 16)
				(_nw72).ArraySet1(func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter34 := _dafny.Iterate((_422_v10).Elements()); ; {
						_compr_32, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _426_v9 _dafny.Int
						_426_v9 = interface{}(_compr_32).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_422_v10, _426_v9) {
							_coll32.Add(Companion_Default___.SafeModuloInt(_426_v9, _dafny.IntOfInt64(560)), _417_cf14)
						}
					}
					return _coll32.ToMap()
				}(), 17)
				(_nw72).ArraySet1((_419_v6).Update((_423_v11).Cardinality(), _416_cf15), 18)
				(_nw72).ArraySet1(_419_v6, 19)
				(_nw72).ArraySet1(_419_v6, 20)
				(_nw72).ArraySet1(_419_v6, 21)
				(_nw72).ArraySet1(_419_v6, 22)
				(_nw72).ArraySet1(_419_v6, 23)
				(_nw72).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _417_cf14), 24)
				(_nw72).ArraySet1((_424_v12).Dtor_cf19(), 25)
				(_nw72).ArraySet1(_419_v6, 26)
				(_nw72).ArraySet1(func() _dafny.Map {
					var _coll33 = _dafny.NewMapBuilder()
					_ = _coll33
					for _iter35 := _dafny.Iterate((_422_v10).Elements()); ; {
						_compr_33, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _427_v13 _dafny.Int
						_427_v13 = interface{}(_compr_33).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_422_v10, _427_v13) {
							_coll33.Add((_427_v13).Minus(_411_i1), _417_cf14)
						}
					}
					return _coll33.ToMap()
				}(), 27)
				_425_v14 = _nw72
				var _428_v15 D4
				_ = _428_v15
				_428_v15 = Companion_D4_.Create_DC7_(_425_v14)
				var _429_v16 _dafny.Map
				_ = _429_v16
				_429_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(_411_i1))
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_418_v5), 0))
				_ = _index58
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_407_v1), 0))
				_ = _index59
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_418_v5), 0))
				_ = _index60
				var _rhs51 _dafny.Int = (_411_i1).Minus((_411_i1).Minus(_411_i1))
				_ = _rhs51
				var _rhs52 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("duj")).Cardinality())
				_ = _rhs52
				var _rhs53 bool = _417_cf14
				_ = _rhs53
				var _rhs54 _dafny.Int = (_dafny.Zero).Minus((_411_i1).Times(Companion_Default___.SafeModuloInt((_429_v16).Cardinality(), _this.F25())))
				_ = _rhs54
				var _rhs55 D4 = _428_v15
				_ = _rhs55
				var _lhs48 _dafny.Array = _418_v5
				_ = _lhs48
				var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_418_v5), 0))
				_ = _lhs49
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				var _lhs51 _dafny.Array = _407_v1
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_407_v1), 0))
				_ = _lhs52
				var _lhs53 _dafny.Array = _418_v5
				_ = _lhs53
				var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_418_v5), 0))
				_ = _lhs54
				(_lhs48).ArraySet1(_rhs51, (_lhs49).Int())
				_lhs50.F22 = _rhs52
				(_lhs51).ArraySet1(_rhs53, (_lhs52).Int())
				(_lhs53).ArraySet1(_rhs54, (_lhs54).Int())
				_428_v15 = _rhs55
				var _430_v17 _dafny.Array
				_ = _430_v17
				var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw73
				_430_v17 = _nw73
				_430_v17 = _430_v17
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_418_v5), 0))
				_ = _index61
				(_418_v5).ArraySet1((Companion_Default___.SafeDivisionInt((_418_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_418_v5), 0))).Int()).(_dafny.Int), p1)).Times(_411_i1), (_index61).Int())
				_418_v5 = (func() _dafny.Array {
					if (func() bool {
						if _416_cf15 {
							return (_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool)
						}
						return true
					})() {
						return _418_v5
					}
					return _418_v5
				})()
			} else {
				var _431___mcc_h2 _dafny.Array = _source7.Get_().(D5_DC9).Cf13
				_ = _431___mcc_h2
				var _432_cf13 _dafny.Array = _431___mcc_h2
				_ = _432_cf13
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_432_cf13), 0))
				_ = _index62
				(_432_cf13).ArraySet1(_411_i1, (_index62).Int())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_432_cf13), 0))
				_ = _index63
				(_432_cf13).ArraySet1((_dafny.Zero).Minus((_411_i1).Minus(_411_i1)), (_index63).Int())
				(globalState).F2 = (_432_cf13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_432_cf13), 0))).Int()).(_dafny.Int)
				var _433_v18 _dafny.Map
				_ = _433_v18
				_433_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _432_cf13)
				var _434_v19 _dafny.MultiSet
				_ = _434_v19
				_434_v19 = _dafny.MultiSetOf(p0)
				var _435_v20 _dafny.Map
				_ = _435_v20
				_435_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_i1, true)
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_432_cf13), 0))
				_ = _index64
				var _rhs56 _dafny.Int = (func() _dafny.Int {
					if (_434_v19).Contains(!(true) || (Companion_Default___.Fm0(_411_i1, p0, p0, _this.F25(), globalState))) {
						return (_434_v19).Multiplicity(!(true) || (Companion_Default___.Fm0(_411_i1, p0, p0, _this.F25(), globalState)))
					}
					return (_435_v20).Cardinality()
				})()
				_ = _rhs56
				var _rhs57 _dafny.Map = (_433_v18).Merge(_433_v18)
				_ = _rhs57
				var _rhs58 _dafny.Int = _411_i1
				_ = _rhs58
				var _rhs59 bool = p0
				_ = _rhs59
				var _rhs60 _dafny.Array = _432_cf13
				_ = _rhs60
				var _lhs55 *GlobalState = globalState
				_ = _lhs55
				var _lhs56 _dafny.Array = _432_cf13
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_432_cf13), 0))
				_ = _lhs57
				var _lhs58 *GlobalState = globalState
				_ = _lhs58
				_lhs55.F15 = _rhs56
				_433_v18 = _rhs57
				(_lhs56).ArraySet1(_rhs58, (_lhs57).Int())
				_lhs58.F18 = _rhs59
				_432_cf13 = _rhs60
				var _436_v21 *C0
				_ = _436_v21
				var _nw74 *C0 = New_C0_()
				_ = _nw74
				_nw74.Ctor__()
				_436_v21 = _nw74
			}
			var _437_v22 _dafny.Array
			_ = _437_v22
			var _nwElement0_9 _dafny.Array = (func() _dafny.Array {
				if p0 {
					return _407_v1
				}
				return _407_v1
			})()
			_ = _nwElement0_9
			var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(21))
			_ = _nw75
			(_nw75).ArraySet1(_nwElement0_9, 0)
			(_nw75).ArraySet1(_407_v1, 1)
			(_nw75).ArraySet1((func() _dafny.Array {
				if p0 {
					return _407_v1
				}
				return _407_v1
			})(), 2)
			(_nw75).ArraySet1(_407_v1, 3)
			(_nw75).ArraySet1(_407_v1, 4)
			(_nw75).ArraySet1((func() _dafny.Array {
				if p0 {
					return _407_v1
				}
				return _407_v1
			})(), 5)
			(_nw75).ArraySet1(_407_v1, 6)
			(_nw75).ArraySet1(_407_v1, 7)
			(_nw75).ArraySet1(_407_v1, 8)
			(_nw75).ArraySet1(_407_v1, 9)
			(_nw75).ArraySet1(_407_v1, 10)
			(_nw75).ArraySet1(_407_v1, 11)
			(_nw75).ArraySet1(_407_v1, 12)
			(_nw75).ArraySet1(_407_v1, 13)
			(_nw75).ArraySet1(_407_v1, 14)
			(_nw75).ArraySet1(_407_v1, 15)
			(_nw75).ArraySet1(_407_v1, 16)
			(_nw75).ArraySet1(_407_v1, 17)
			(_nw75).ArraySet1(_407_v1, 18)
			(_nw75).ArraySet1(_407_v1, 19)
			(_nw75).ArraySet1(_407_v1, 20)
			_437_v22 = _nw75
			var _438_v23 _dafny.Sequence
			_ = _438_v23
			_438_v23 = _dafny.SeqOf(Companion_D5_.Create_DC10_(p0, false))
			var _439_v24 _dafny.Sequence
			_ = _439_v24
			_439_v24 = _dafny.SeqOf(!(p0))
			var _440_v25 _dafny.Sequence
			_ = _440_v25
			_440_v25 = _dafny.UnicodeSeqOfUtf8Bytes("nnhp")
			var _441_v26 _dafny.Sequence
			_ = _441_v26
			_441_v26 = _dafny.SeqOf(_dafny.IntOfUint32((_440_v25).Cardinality()))
			var _442_v27 _dafny.Array
			_ = _442_v27
			var _nwElement0_10 _dafny.Int = (_411_i1).Minus(_this.F25())
			_ = _nwElement0_10
			var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(14))
			_ = _nw76
			(_nw76).ArraySet1(_nwElement0_10, 0)
			(_nw76).ArraySet1(_411_i1, 1)
			(_nw76).ArraySet1(_this.F25(), 2)
			(_nw76).ArraySet1((func() _dafny.Int {
				if p0 {
					return _dafny.IntOfUint32((_438_v23).Cardinality())
				}
				return _this.F25()
			})(), 3)
			(_nw76).ArraySet1(_411_i1, 4)
			(_nw76).ArraySet1(_this.F25(), 5)
			(_nw76).ArraySet1(_dafny.IntOfUint32((_439_v24).Cardinality()), 6)
			(_nw76).ArraySet1(_411_i1, 7)
			(_nw76).ArraySet1(p1, 8)
			(_nw76).ArraySet1(_411_i1, 9)
			(_nw76).ArraySet1((_411_i1).Minus(_411_i1), 10)
			(_nw76).ArraySet1(p1, 11)
			(_nw76).ArraySet1(_this.F25(), 12)
			(_nw76).ArraySet1((_441_v26).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_441_v26).Cardinality()))).Uint32()).(_dafny.Int), 13)
			_442_v27 = _nw76
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_442_v27), 0))
			_ = _index65
			(_442_v27).ArraySet1(_this.F25(), (_index65).Int())
			var _443_v28 _dafny.Set
			_ = _443_v28
			_443_v28 = _dafny.SetOf(p0, false, p0, p0)
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_442_v27), 0))
			_ = _index66
			var _rhs61 bool = ((_443_v28).Intersection(_443_v28)).IsProperSubsetOf((_dafny.SetOf(p0)).Difference(_443_v28))
			_ = _rhs61
			var _rhs62 _dafny.Array = _437_v22
			_ = _rhs62
			var _rhs63 _dafny.Int = ((_this.F25()).Minus((_dafny.Zero).Minus(p1))).Times(_dafny.IntOfInt64(-958))
			_ = _rhs63
			var _rhs64 D5 = _413_v4
			_ = _rhs64
			var _lhs59 *GlobalState = globalState
			_ = _lhs59
			var _lhs60 _dafny.Array = _442_v27
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_442_v27), 0))
			_ = _lhs61
			_lhs59.F18 = _rhs61
			_437_v22 = _rhs62
			(_lhs60).ArraySet1(_rhs63, (_lhs61).Int())
			_413_v4 = _rhs64
			var _444_v29 D6
			_ = _444_v29
			_444_v29 = Companion_D6_.Create_DC12_(p0, _this.F25())
			(globalState).F18 = (_444_v29).Dtor_cf17()
		}
		var _445_v30 _dafny.Set
		_ = _445_v30
		_445_v30 = _dafny.SetOf(!(p0))
		var _446_v31 _dafny.Sequence
		_ = _446_v31
		_446_v31 = _dafny.SeqOf((_445_v30).Cardinality(), p1, p1)
		r3 = (func() _dafny.Int {
			if p0 {
				return _dafny.IntOfInt64(-490)
			}
			return (_446_v31).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_446_v31).Cardinality()))).Uint32()).(_dafny.Int)
		})()
		(globalState).F18 = p0
		var _447_v32 _dafny.Map
		_ = _447_v32
		_447_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(295), _this.F25())
		if (_this.F25()).Cmp((func() _dafny.Int {
			if (_447_v32).Contains(p1) {
				return (_447_v32).Get(p1).(_dafny.Int)
			}
			return _this.F25()
		})()) >= 0 {
			if !((func() bool {
				if p0 {
					return p0
				}
				return p0
			})()) {
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_407_v1), 0))
				_ = _index67
				(_407_v1).ArraySet1((func() bool {
					if p0 {
						return p0
					}
					return p0
				})(), (_index67).Int())
				var _448_v33 _dafny.Map
				_ = _448_v33
				_448_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), p0)
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_407_v1), 0))
				_ = _index68
				(_407_v1).ArraySet1(((func() _dafny.Map {
					if p0 {
						return (_448_v33).Update(_this.F25(), false)
					}
					return _448_v33
				})()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(176), true)), (_index68).Int())
				var _449_v34 _dafny.Array
				_ = _449_v34
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_12
				var _nw77 _dafny.Array
				_ = _nw77
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw77 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) bool = (func(_450_v1 _dafny.Array) func(_dafny.Int) bool {
						return func(_451_i2 _dafny.Int) bool {
							return (_450_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_450_v1), 0))).Int()).(bool)
						}
					})(_407_v1)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw77 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw77).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw77).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_449_v34 = _nw77
				_449_v34 = _449_v34
				r0 = _dafny.IntOfInt64(-777)
				var _452_v35 _dafny.Sequence
				_ = _452_v35
				_452_v35 = _dafny.SeqOf(!((_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool)))
				_452_v35 = Companion_Default___.Fm7(globalState)
				(globalState).F18 = Companion_Default___.Fm0((p1).Times(_this.F25()), (_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool), p0, (_this).Fm6(globalState), globalState)
			} else {
				var _453_v36 _dafny.MultiSet
				_ = _453_v36
				_453_v36 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(309))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}(func(_454_i3 _dafny.Int) _dafny.CodePoint {
					return (_this).F24()
				}))).Cardinality()))
				var _455_v37 D7
				_ = _455_v37
				_455_v37 = Companion_D7_.Create_DC14_(p0, _dafny.IntOfInt64(882), _this.F25())
				(globalState).F18 = ((_445_v30).Cardinality()).Cmp(((_453_v36).Cardinality()).Minus((_455_v37).Dtor_cf21())) > 0
				(globalState).F18 = p0
				var _456_v38 _dafny.Sequence
				_ = _456_v38
				_456_v38 = _dafny.SeqOf(p0)
				var _457_v39 _dafny.Sequence
				_ = _457_v39
				_457_v39 = _dafny.UnicodeSeqOfUtf8Bytes("ybuioxvd")
				var _458_v40 _dafny.MultiSet
				_ = _458_v40
				_458_v40 = _dafny.MultiSetOf(p0, p0, p0, p0)
				var _459_v41 _dafny.Array
				_ = _459_v41
				var _nwElement0_11 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqOf(_this.F25())).Cardinality())).Plus(_this.F25())
				_ = _nwElement0_11
				var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(12))
				_ = _nw78
				(_nw78).ArraySet1(_nwElement0_11, 0)
				(_nw78).ArraySet1((_this.F25()).Minus((_dafny.Zero).Minus(p1)), 1)
				(_nw78).ArraySet1((_this).Fm6(globalState), 2)
				(_nw78).ArraySet1(p1, 3)
				(_nw78).ArraySet1(_this.F25(), 4)
				(_nw78).ArraySet1(((Companion_Default___.Fm8(_dafny.IntOfUint32((_457_v39).Cardinality()), p0, _456_v38, globalState)).Difference(_458_v40)).Cardinality(), 5)
				(_nw78).ArraySet1(p1, 6)
				(_nw78).ArraySet1((_this.F25()).Plus(p1), 7)
				(_nw78).ArraySet1((_dafny.IntOfInt64(687)).Times(_this.F25()), 8)
				(_nw78).ArraySet1((_this).Fm6(globalState), 9)
				(_nw78).ArraySet1(p1, 10)
				(_nw78).ArraySet1((_dafny.Zero).Minus(p1), 11)
				_459_v41 = _nw78
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_459_v41), 0))
				_ = _index69
				(_459_v41).ArraySet1((_this).Fm6(globalState), (_index69).Int())
				var _460_v42 D6
				_ = _460_v42
				_460_v42 = Companion_D6_.Create_DC12_(p0, p1)
				var _461_v43 D3
				_ = _461_v43
				_461_v43 = Companion_D3_.Create_DC4_(Companion_Default___.Fm9(true, _dafny.Companion_Sequence_.Update(_456_v38, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_456_v38).Cardinality()))).Uint32(), (_460_v42).Dtor_cf17()), globalState))
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_459_v41), 0))
				_ = _index70
				var _rhs65 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0), (Companion_Default___.SafeIndex(_this.F25(), _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))).Uint32(), Companion_Default___.Fm0(p1, false, p0, _this.F25(), globalState))
				_ = _rhs65
				var _rhs66 _dafny.Sequence = (_this).Fm3((_this).F24(), globalState)
				_ = _rhs66
				var _rhs67 _dafny.Int = (_this).Fm6(globalState)
				_ = _rhs67
				var _rhs68 D3 = _461_v43
				_ = _rhs68
				var _lhs62 *GlobalState = globalState
				_ = _lhs62
				var _lhs63 _dafny.Array = _459_v41
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_459_v41), 0))
				_ = _lhs64
				_456_v38 = _rhs65
				_lhs62.F3 = _rhs66
				(_lhs63).ArraySet1(_rhs67, (_lhs64).Int())
				_461_v43 = _rhs68
				var _462_v44 D5
				_ = _462_v44
				_462_v44 = Companion_D5_.Create_DC9_(_459_v41)
				var _pat_let_tv2 = _459_v41
				_ = _pat_let_tv2
				var _463_v45 _dafny.Array
				_ = _463_v45
				var _nwElement0_12 D5 = Companion_D5_.Create_DC9_(_459_v41)
				_ = _nwElement0_12
				var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(16))
				_ = _nw79
				(_nw79).ArraySet1(_nwElement0_12, 0)
				(_nw79).ArraySet1(_462_v44, 1)
				(_nw79).ArraySet1(_462_v44, 2)
				(_nw79).ArraySet1(Companion_D5_.Create_DC9_(_459_v41), 3)
				(_nw79).ArraySet1(Companion_D5_.Create_DC9_(_459_v41), 4)
				(_nw79).ArraySet1(_462_v44, 5)
				(_nw79).ArraySet1(_462_v44, 6)
				(_nw79).ArraySet1(_462_v44, 7)
				(_nw79).ArraySet1(Companion_D5_.Create_DC9_(_459_v41), 8)
				(_nw79).ArraySet1(func(_pat_let4_0 D5) D5 {
					return func(_464_dt__update__tmp_h0 D5) D5 {
						return func(_pat_let5_0 _dafny.Array) D5 {
							return func(_465_dt__update_hcf13_h0 _dafny.Array) D5 {
								return Companion_D5_.Create_DC9_(_465_dt__update_hcf13_h0)
							}(_pat_let5_0)
						}(_pat_let_tv2)
					}(_pat_let4_0)
				}(Companion_D5_.Create_DC9_(_459_v41)), 9)
				(_nw79).ArraySet1(_462_v44, 10)
				(_nw79).ArraySet1(_462_v44, 11)
				(_nw79).ArraySet1(_462_v44, 12)
				(_nw79).ArraySet1(_462_v44, 13)
				(_nw79).ArraySet1(_462_v44, 14)
				(_nw79).ArraySet1(Companion_D5_.Create_DC9_(_459_v41), 15)
				_463_v45 = _nw79
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_463_v45), 0))
				_ = _index71
				(_463_v45).ArraySet1(_462_v44, (_index71).Int())
				var _pat_let_tv3 = _459_v41
				_ = _pat_let_tv3
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_463_v45), 0))
				_ = _index72
				(_463_v45).ArraySet1(func(_pat_let6_0 D5) D5 {
					return func(_466_dt__update__tmp_h1 D5) D5 {
						return func(_pat_let7_0 _dafny.Array) D5 {
							return func(_467_dt__update_hcf13_h1 _dafny.Array) D5 {
								return Companion_D5_.Create_DC9_(_467_dt__update_hcf13_h1)
							}(_pat_let7_0)
						}(_pat_let_tv3)
					}(_pat_let6_0)
				}(_462_v44), (_index72).Int())
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_407_v1), 0))
				_ = _index73
				(_407_v1).ArraySet1(p0, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_407_v1), 0))
				_ = _index74
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_459_v41), 0))
				_ = _index75
				var _rhs69 bool = p0
				_ = _rhs69
				var _rhs70 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_this.F25(), (_459_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_459_v41), 0))).Int()).(_dafny.Int)))
				_ = _rhs70
				var _rhs71 _dafny.Int = p1
				_ = _rhs71
				var _rhs72 _dafny.Int = p1
				_ = _rhs72
				var _rhs73 bool = p0
				_ = _rhs73
				var _lhs65 _dafny.Array = _407_v1
				_ = _lhs65
				var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_407_v1), 0))
				_ = _lhs66
				var _lhs67 _dafny.Array = _459_v41
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_459_v41), 0))
				_ = _lhs68
				var _lhs69 *C1 = _this
				_ = _lhs69
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				(_lhs65).ArraySet1(_rhs69, (_lhs66).Int())
				(_lhs67).ArraySet1(_rhs70, (_lhs68).Int())
				_lhs69.F25_set_(_rhs71)
				r0 = _rhs72
				_lhs70.F18 = _rhs73
			}
			var _468_v46 _dafny.Array
			_ = _468_v46
			var _nwElement0_13 _dafny.CodePoint = _dafny.CodePoint('u')
			_ = _nwElement0_13
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(10))
			_ = _nw80
			(_nw80).ArraySet1CodePoint(_nwElement0_13, 0)
			(_nw80).ArraySet1CodePoint(_dafny.CodePoint('e'), 1)
			(_nw80).ArraySet1CodePoint((_this).F24(), 2)
			(_nw80).ArraySet1CodePoint((_this).F24(), 3)
			(_nw80).ArraySet1CodePoint(Companion_Default___.Fm10(p0, p0, p0, p0, globalState), 4)
			(_nw80).ArraySet1CodePoint((_this).F24(), 5)
			(_nw80).ArraySet1CodePoint((_this).F24(), 6)
			(_nw80).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return (_this).F24()
				}
				return (_this).F24()
			})(), 7)
			(_nw80).ArraySet1CodePoint((_this).F24(), 8)
			(_nw80).ArraySet1CodePoint((_this).F24(), 9)
			_468_v46 = _nw80
			_468_v46 = _468_v46
			(globalState).F18 = !(true)
			var _469_v47 D7
			_ = _469_v47
			_469_v47 = Companion_D7_.Create_DC14_(p0, p1, p1)
			(globalState).F18 = (_469_v47).Dtor_cf20()
			var _470_v48 D2
			_ = _470_v48
			_470_v48 = Companion_D2_.Create_DC2_(p0)
			_470_v48 = _470_v48
		} else {
			var _471_v49 *C0
			_ = _471_v49
			var _nw81 *C0 = New_C0_()
			_ = _nw81
			_nw81.Ctor__()
			_471_v49 = _nw81
			(globalState).F18 = p0
			var _472_v50 _dafny.Sequence
			_ = _472_v50
			_472_v50 = _dafny.SeqOf(p0)
			var _473_v51 _dafny.Map
			_ = _473_v51
			_473_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), _472_v50)
			var _474_v52 _dafny.Sequence
			_ = _474_v52
			_474_v52 = _dafny.UnicodeSeqOfUtf8Bytes("fwhru")
			var _rhs74 bool = !((_dafny.MultiSetOf(p1, p1)).IsDisjointFrom(_dafny.MultiSetFromSeq(_446_v31)))
			_ = _rhs74
			var _rhs75 _dafny.Int = _dafny.IntOfUint32((_474_v52).Cardinality())
			_ = _rhs75
			var _rhs76 _dafny.Map = _473_v51
			_ = _rhs76
			var _lhs71 *GlobalState = globalState
			_ = _lhs71
			var _lhs72 *GlobalState = globalState
			_ = _lhs72
			_lhs71.F18 = _rhs74
			_lhs72.F11 = _rhs75
			_473_v51 = _rhs76
			var _475_v53 _dafny.MultiSet
			_ = _475_v53
			_475_v53 = _dafny.MultiSetOf(!(p0))
			var _rhs77 _dafny.Int = (_this.F25()).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_this.F25(), _dafny.IntOfInt64(-356))).Cardinality()))).Cardinality()))
			_ = _rhs77
			var _rhs78 _dafny.Int = p1
			_ = _rhs78
			var _rhs79 bool = (Companion_Default___.Fm8(_this.F25(), p0, _472_v50, globalState)).IsSubsetOf((_475_v53).Union(_475_v53))
			_ = _rhs79
			var _lhs73 *GlobalState = globalState
			_ = _lhs73
			var _lhs74 *GlobalState = globalState
			_ = _lhs74
			var _lhs75 *GlobalState = globalState
			_ = _lhs75
			_lhs73.F22 = _rhs77
			_lhs74.F16 = _rhs78
			_lhs75.F18 = _rhs79
			if p0 {
				var _476_v54 _dafny.Array
				_ = _476_v54
				var _nwElement0_14 _dafny.Map = _447_v32
				_ = _nwElement0_14
				var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(22))
				_ = _nw82
				(_nw82).ArraySet1(_nwElement0_14, 0)
				(_nw82).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), p1)).Merge((_447_v32).Update((_this).Fm6(globalState), p1)), 1)
				(_nw82).ArraySet1((_447_v32).Merge(_447_v32), 2)
				(_nw82).ArraySet1(_447_v32, 3)
				(_nw82).ArraySet1(_447_v32, 4)
				(_nw82).ArraySet1(_447_v32, 5)
				(_nw82).ArraySet1(_447_v32, 6)
				(_nw82).ArraySet1(_447_v32, 7)
				(_nw82).ArraySet1((_447_v32).Merge(_447_v32), 8)
				(_nw82).ArraySet1(_447_v32, 9)
				(_nw82).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Merge(_447_v32), 10)
				(_nw82).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F25()), 11)
				(_nw82).ArraySet1(_447_v32, 12)
				(_nw82).ArraySet1(_447_v32, 13)
				(_nw82).ArraySet1((_447_v32).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm6(globalState), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), _407_v1)).Cardinality())), 14)
				(_nw82).ArraySet1((Companion_Default___.Fm11(_dafny.CodePoint('c'), p0, globalState)).Merge(_447_v32), 15)
				(_nw82).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), _this.F25()), 16)
				(_nw82).ArraySet1(_447_v32, 17)
				(_nw82).ArraySet1((func() _dafny.Map {
					if true {
						return _447_v32
					}
					return (_447_v32).Update(p1, (_dafny.SetOf(p0, p0, p0)).Cardinality())
				})(), 18)
				(_nw82).ArraySet1(_447_v32, 19)
				(_nw82).ArraySet1(_447_v32, 20)
				(_nw82).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), 21)
				_476_v54 = _nw82
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_476_v54), 0))
				_ = _index76
				(_476_v54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), (_index76).Int())
				var _477_v55 _dafny.Map
				_ = _477_v55
				_477_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), (_this).F24())
				var _478_v56 _dafny.Map
				_ = _478_v56
				_478_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), true))
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_476_v54), 0))
				_ = _index77
				var _rhs80 _dafny.Map = Companion_Default___.Fm11((func() _dafny.CodePoint {
					if (_477_v55).Contains(p1) {
						return (_477_v55).Get(p1).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('t')
				})(), p0, globalState)
				_ = _rhs80
				var _rhs81 _dafny.Int = (func() _dafny.Int {
					if p0 {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_478_v56).Cardinality(), (_this).Fm6(globalState))).Cardinality()
					}
					return (_dafny.IntOfInt64(248)).Minus(_dafny.IntOfInt64(281))
				})()
				_ = _rhs81
				var _rhs82 bool = p0
				_ = _rhs82
				var _rhs83 *C0 = _471_v49
				_ = _rhs83
				var _lhs76 _dafny.Array = _476_v54
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_476_v54), 0))
				_ = _lhs77
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				(_lhs76).ArraySet1(_rhs80, (_lhs77).Int())
				_lhs78.F22 = _rhs81
				_lhs79.F18 = _rhs82
				_471_v49 = _rhs83
				(globalState).F18 = p0
				(globalState).F18 = p0
				var _479_v57 _dafny.Map
				_ = _479_v57
				_479_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Minus(p1), _406_v0)
				_479_v57 = _479_v57
				(globalState).F18 = p0
			} else {
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))
				_ = _index78
				(_407_v1).ArraySet1((_this.F25()).Cmp(p1) <= 0, (_index78).Int())
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))
				_ = _index79
				(_407_v1).ArraySet1((_472_v50).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((Companion_D3_.Create_DC5_(_472_v50, _dafny.IntOfInt64(-73), _dafny.IntOfInt64(-879))).Dtor_cf6()).Cardinality()), _dafny.IntOfUint32((_472_v50).Cardinality()))).Uint32()).(bool), (_index79).Int())
				(globalState).F0 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if p0 {
						return _this.F25()
					}
					return (_this).Fm6(globalState)
				})(), Companion_Default___.SafeDivisionInt(p1, p1))
				(globalState).F11 = _this.F25()
				var _480_v58 _dafny.Sequence
				_ = _480_v58
				_480_v58 = _dafny.SeqOf(_445_v30)
				var _481_v59 _dafny.Array
				_ = _481_v59
				var _nwElement0_15 _dafny.Set = _445_v30
				_ = _nwElement0_15
				var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(20))
				_ = _nw83
				(_nw83).ArraySet1(_nwElement0_15, 0)
				(_nw83).ArraySet1((_445_v30).Difference(_445_v30), 1)
				(_nw83).ArraySet1((_445_v30).Difference(_dafny.SetOf((_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool), p0)), 2)
				(_nw83).ArraySet1(_dafny.SetOf(!((_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool))), 3)
				(_nw83).ArraySet1(_dafny.SetOf((_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool)), 4)
				(_nw83).ArraySet1(_445_v30, 5)
				(_nw83).ArraySet1((_480_v58).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_dafny.IntOfUint32((_446_v31).Cardinality()), p1)).Cardinality(), _dafny.IntOfUint32((_480_v58).Cardinality()))).Uint32()).(_dafny.Set), 6)
				(_nw83).ArraySet1(_445_v30, 7)
				(_nw83).ArraySet1(_445_v30, 8)
				(_nw83).ArraySet1(_445_v30, 9)
				(_nw83).ArraySet1(_dafny.SetOf((_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool), (_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool)), 10)
				(_nw83).ArraySet1((_dafny.SetOf((_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool), p0)).Intersection(_445_v30), 11)
				(_nw83).ArraySet1(_445_v30, 12)
				(_nw83).ArraySet1(_445_v30, 13)
				(_nw83).ArraySet1((_445_v30).Union(_445_v30), 14)
				(_nw83).ArraySet1(_445_v30, 15)
				(_nw83).ArraySet1(_445_v30, 16)
				(_nw83).ArraySet1(_445_v30, 17)
				(_nw83).ArraySet1(_445_v30, 18)
				(_nw83).ArraySet1(_445_v30, 19)
				_481_v59 = _nw83
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_481_v59), 0))
				_ = _index80
				(_481_v59).ArraySet1((_445_v30).Difference(_445_v30), (_index80).Int())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_481_v59), 0))
				_ = _index81
				var _rhs84 _dafny.Set = (Companion_Default___.Fm12((_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool), p0, true, globalState)).Union((_445_v30).Intersection(_445_v30))
				_ = _rhs84
				var _rhs85 bool = (_407_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))).Int()).(bool)
				_ = _rhs85
				var _lhs80 _dafny.Array = _481_v59
				_ = _lhs80
				var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_481_v59), 0))
				_ = _lhs81
				var _lhs82 *GlobalState = globalState
				_ = _lhs82
				(_lhs80).ArraySet1(_rhs84, (_lhs81).Int())
				_lhs82.F18 = _rhs85
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_407_v1), 0))
				_ = _index82
				(_407_v1).ArraySet1(Companion_Default___.Fm0(Companion_Default___.SafeDivisionInt(_this.F25(), _dafny.IntOfUint32((_474_v52).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F25())).Contains(_dafny.IntOfInt64(658)), !((_445_v30).IsSubsetOf((_481_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(772), _dafny.ArrayLen((_481_v59), 0))).Int()).(_dafny.Set))), (_dafny.Zero).Minus(p1), globalState), (_index82).Int())
			}
		}
		r0 = Companion_Default___.SafeModuloInt(_this.F25(), p1)
		r1 = ((func() _dafny.Int {
			if true {
				return (_this).Fm6(globalState)
			}
			return p1
		})()).Minus((func() _dafny.Int {
			if Companion_Default___.Fm0(p1, p0, p0, _this.F25(), globalState) {
				return p1
			}
			return _this.F25()
		})())
		r2 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p1, _this.F25())))
		r3 = _this.F25()
		return r0, r1, r2, r3
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f27 _dafny.Int
	_f28 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f27 = _dafny.Zero
	_this._f28 = false
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f27 _dafny.Int, f28 bool) {
	{
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *C2) Fm4(p0 _dafny.Set, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf((_this).F28(), (_this).F28(), (_this).F28(), (_this).F28(), (_this).F28())).Intersection((_dafny.MultiSetOf((_this).F28(), (_this).F28(), true, (_this).F28())).Union(_dafny.MultiSetOf((_this).F28())))
	}
}
func (_this *C2) M3(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _482_v0 _dafny.CodePoint
		_ = _482_v0
		_482_v0 = _dafny.CodePoint('m')
		var _483_v1 _dafny.Sequence
		_ = _483_v1
		_483_v1 = _dafny.UnicodeSeqOfUtf8Bytes("gxn")
		var _484_v2 _dafny.Sequence
		_ = _484_v2
		_484_v2 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(269))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}((func(_485_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_486_i0 _dafny.Int) _dafny.CodePoint {
				return _485_v0
			}
		})(_482_v0)))
		var _487_v3 _dafny.Array
		_ = _487_v3
		var _nwElement0_16 _dafny.Sequence = _483_v1
		_ = _nwElement0_16
		var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(12))
		_ = _nw84
		(_nw84).ArraySet1(_nwElement0_16, 0)
		(_nw84).ArraySet1(_484_v2, 1)
		(_nw84).ArraySet1(_484_v2, 2)
		(_nw84).ArraySet1(_483_v1, 3)
		(_nw84).ArraySet1(_484_v2, 4)
		(_nw84).ArraySet1(_484_v2, 5)
		(_nw84).ArraySet1(_484_v2, 6)
		(_nw84).ArraySet1(_484_v2, 7)
		(_nw84).ArraySet1(_483_v1, 8)
		(_nw84).ArraySet1(_484_v2, 9)
		(_nw84).ArraySet1(_484_v2, 10)
		(_nw84).ArraySet1(_484_v2, 11)
		_487_v3 = _nw84
		var _488_v4 _dafny.Map
		_ = _488_v4
		_488_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F27()).Plus((_this).F27()), _487_v3)
		var _489_v5 _dafny.Set
		_ = _489_v5
		_489_v5 = _dafny.SetOf((_this).F27(), (_this).F27(), (_this).F27(), _dafny.IntOfInt64(-942))
		var _rhs86 bool = (_this).F28()
		_ = _rhs86
		var _rhs87 _dafny.CodePoint = _dafny.CodePoint('i')
		_ = _rhs87
		var _rhs88 _dafny.Sequence = _483_v1
		_ = _rhs88
		var _rhs89 _dafny.Array = (func() _dafny.Array {
			if (_488_v4).Contains((_489_v5).Cardinality()) {
				return (_488_v4).Get((_489_v5).Cardinality()).(_dafny.Array)
			}
			return _487_v3
		})()
		_ = _rhs89
		var _lhs83 *GlobalState = globalState
		_ = _lhs83
		var _lhs84 *GlobalState = globalState
		_ = _lhs84
		_lhs83.F18 = _rhs86
		_482_v0 = _rhs87
		_lhs84.F3 = _rhs88
		_487_v3 = _rhs89
		var _490_v6 _dafny.MultiSet
		_ = _490_v6
		_490_v6 = _dafny.MultiSetOf(p0, (_this).F28())
		var _491_v7 D3
		_ = _491_v7
		_491_v7 = Companion_D3_.Create_DC5_(_dafny.SeqOf((_this).F28()), (_490_v6).Cardinality(), (_this).F27())
		var _492_v8 _dafny.Sequence
		_ = _492_v8
		_492_v8 = (_491_v7).Dtor_cf6()
		var _493_v9 _dafny.Sequence
		_ = _493_v9
		_493_v9 = _dafny.SeqOf((_this).F28())
		var _source8 _dafny.Sequence = _493_v9
		_ = _source8
		var _494___mcc_h0 _dafny.Sequence = _source8
		_ = _494___mcc_h0
		var _495_cf0 _dafny.Sequence = _494___mcc_h0
		_ = _495_cf0
		if _dafny.Companion_Sequence_.IsPrefixOf(_483_v1, _dafny.UnicodeSeqOfUtf8Bytes("q")) {
			var _496_v10 _dafny.Array
			_ = _496_v10
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
			_ = _nw85
			_496_v10 = _nw85
			_496_v10 = _496_v10
			var _497_v11 _dafny.Map
			_ = _497_v11
			_497_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_483_v1, (_491_v7).Dtor_cf7())
			_497_v11 = (_497_v11).Update(_dafny.Companion_Sequence_.Concatenate(_483_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(59))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg39 _dafny.Int) interface{} {
					return coer39(arg39)
				}
			}(func(_498_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			}))), (_dafny.Zero).Minus(((_this).F27()).Times((_this).F27())))
			(globalState).F3 = _483_v1
			var _499_v12 _dafny.Array
			_ = _499_v12
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
			_ = _nw86
			_499_v12 = _nw86
			var _500_v13 D4
			_ = _500_v13
			_500_v13 = Companion_D4_.Create_DC7_(_499_v12)
			_499_v12 = (_500_v13).Dtor_cf12()
			var _501_v14 _dafny.Set
			_ = _501_v14
			_501_v14 = _dafny.SetOf((_this).F28(), p0, (!(p1)) && (p1), ((_this).F28()) == (p0), (_this).F28())
			(globalState).F22 = (_501_v14).Cardinality()
		} else {
			(globalState).F0 = (_this).F27()
			var _502_v15 *C0
			_ = _502_v15
			var _nw87 *C0 = New_C0_()
			_ = _nw87
			_nw87.Ctor__()
			_502_v15 = _nw87
			_490_v6 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_493_v9, (Companion_D3_.Create_DC5_(_495_cf0, (_this).F27(), (_this).F27())).Dtor_cf6()))).Intersection(_490_v6)
			var _503_v16 _dafny.Array
			_ = _503_v16
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw88
			_503_v16 = _nw88
			var _504_v17 D5
			_ = _504_v17
			_504_v17 = Companion_D5_.Create_DC9_(_503_v16)
			var _505_v18 _dafny.Sequence
			_ = _505_v18
			_505_v18 = _dafny.SeqOf((_504_v17).Dtor_cf13())
			(globalState).F11 = _dafny.IntOfUint32((_505_v18).Cardinality())
			var _506_v19 D5
			_ = _506_v19
			_506_v19 = Companion_D5_.Create_DC10_(p1, (_495_cf0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_495_cf0).Cardinality()))).Uint32()).(bool))
			var _507_v20 _dafny.Array
			_ = _507_v20
			var _nwElement0_17 bool = p1
			_ = _nwElement0_17
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(19))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_17, 0)
			(_nw89).ArraySet1(!(((_this).F28()) || (p0)), 1)
			(_nw89).ArraySet1(!(p0), 2)
			(_nw89).ArraySet1((func() bool {
				if (_495_cf0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_495_cf0).Cardinality()))).Uint32()).(bool) {
					return (_this).F28()
				}
				return !((_this).F28())
			})(), 3)
			(_nw89).ArraySet1(!((_489_v5).IsDisjointFrom(_dafny.SetOf((_this).F27()))), 4)
			(_nw89).ArraySet1((_dafny.MultiSetOf(!((_this).F28()), (_495_cf0).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_495_cf0).Cardinality()))).Uint32()).(bool))).Equals(_490_v6), 5)
			(_nw89).ArraySet1((_this).F28(), 6)
			(_nw89).ArraySet1(!(p1) || (p1), 7)
			(_nw89).ArraySet1((_this).F28(), 8)
			(_nw89).ArraySet1((_506_v19).Dtor_cf14(), 9)
			(_nw89).ArraySet1((_this).F28(), 10)
			(_nw89).ArraySet1(p1, 11)
			(_nw89).ArraySet1((_this).F28(), 12)
			(_nw89).ArraySet1(Companion_Default___.Fm0((_this).F27(), p1, (_this).F28(), (_this).F27(), globalState), 13)
			(_nw89).ArraySet1(p1, 14)
			(_nw89).ArraySet1(!(p0), 15)
			(_nw89).ArraySet1(((_this).F27()).Cmp((_this).F27()) > 0, 16)
			(_nw89).ArraySet1(p1, 17)
			(_nw89).ArraySet1((_this).F28(), 18)
			_507_v20 = _nw89
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_507_v20), 0))
			_ = _index83
			(_507_v20).ArraySet1(((_this).F27()).Cmp((_this).F27()) > 0, (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_507_v20), 0))
			_ = _index84
			(_507_v20).ArraySet1((func() bool {
				if (_this).F28() {
					return p0
				}
				return p0
			})(), (_index84).Int())
		}
		(globalState).F19 = _dafny.SeqOf((_this).F27(), (_this).F27())
		var _508_v21 _dafny.Array
		_ = _508_v21
		var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
		_ = _nw90
		_508_v21 = _nw90
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_13
		var _nw91 _dafny.Array
		_ = _nw91
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw91 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.CodePoint = (func(_509_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_510_i2 _dafny.Int) _dafny.CodePoint {
					return _509_v0
				}
			})(_482_v0)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw91 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw91).ArraySet1CodePoint(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw91).ArraySet1CodePoint(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_508_v21 = _nw91
		var _511_v22 D3
		_ = _511_v22
		_511_v22 = Companion_D3_.Create_DC6_(_dafny.IntOfInt64(617), (_dafny.Zero).Minus((_this).F27()), p0)
		var _512_v23 _dafny.Map
		_ = _512_v23
		_512_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), p1)
		var _513_v24 _dafny.Map
		_ = _513_v24
		_513_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(381))
		var _514_v25 _dafny.MultiSet
		_ = _514_v25
		_514_v25 = _dafny.MultiSetOf((_511_v22).Dtor_cf10(), (_this).F27(), (_512_v23).Cardinality(), ((_489_v5).Union(_489_v5)).Cardinality(), (((_513_v24).Update(p1, (_this).F27())).Merge((_513_v24).Update(p0, (_this).F27()))).Cardinality())
		var _515_v26 _dafny.Sequence
		_ = _515_v26
		_515_v26 = _dafny.SeqOf(_508_v21)
		var _516_v27 _dafny.Sequence
		_ = _516_v27
		_516_v27 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus((_this).F27())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_this).F27()), _513_v24)
		var _rhs90 _dafny.MultiSet = ((_514_v25).Union(_514_v25)).Union((_514_v25).Intersection(_dafny.MultiSetOf((func() _dafny.Int {
			if (_514_v25).Contains((_this).F27()) {
				return (_514_v25).Multiplicity((_this).F27())
			}
			return (_this).F27()
		})())))
		_ = _rhs90
		var _rhs91 _dafny.Sequence = _dafny.SeqOf(_508_v21)
		_ = _rhs91
		var _rhs92 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("o"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_516_v27, _dafny.SeqOf(_513_v24, _513_v24))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("o")).Cardinality()))).Uint32(), _482_v0)
		_ = _rhs92
		var _rhs93 bool = !(p1)
		_ = _rhs93
		var _lhs85 *GlobalState = globalState
		_ = _lhs85
		var _lhs86 *GlobalState = globalState
		_ = _lhs86
		_514_v25 = _rhs90
		_515_v26 = _rhs91
		_lhs85.F3 = _rhs92
		_lhs86.F18 = _rhs93
		var _517_v28 _dafny.Map
		_ = _517_v28
		_517_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _518_v29 _dafny.Set
		_ = _518_v29
		_518_v29 = _dafny.SetOf(_517_v28)
		if !(!((_518_v29).IsProperSubsetOf(_518_v29))) {
			var _519_v30 _dafny.MultiSet
			_ = _519_v30
			_519_v30 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F27()), _dafny.IntOfInt64(-840))
			var _520_v31 _dafny.Sequence
			_ = _520_v31
			_520_v31 = _dafny.SeqOf(_489_v5)
			var _521_v32 _dafny.Array
			_ = _521_v32
			var _nwElement0_18 _dafny.Int = (_519_v30).Cardinality()
			_ = _nwElement0_18
			var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(24))
			_ = _nw92
			(_nw92).ArraySet1(_nwElement0_18, 0)
			(_nw92).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.IntOfInt64(872))).Cardinality(), 1)
			(_nw92).ArraySet1((_this).F27(), 2)
			(_nw92).ArraySet1(_dafny.IntOfUint32((_483_v1).Cardinality()), 3)
			(_nw92).ArraySet1((_this).F27(), 4)
			(_nw92).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()), 5)
			(_nw92).ArraySet1((_this).F27(), 6)
			(_nw92).ArraySet1((_dafny.Zero).Minus((_this).F27()), 7)
			(_nw92).ArraySet1((_this).F27(), 8)
			(_nw92).ArraySet1((_this).F27(), 9)
			(_nw92).ArraySet1((_this).F27(), 10)
			(_nw92).ArraySet1((_this).F27(), 11)
			(_nw92).ArraySet1((_this).F27(), 12)
			(_nw92).ArraySet1((_this).F27(), 13)
			(_nw92).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xrr")).Cardinality()), 14)
			(_nw92).ArraySet1((_this).F27(), 15)
			(_nw92).ArraySet1((_this).F27(), 16)
			(_nw92).ArraySet1((_this).F27(), 17)
			(_nw92).ArraySet1((_this).F27(), 18)
			(_nw92).ArraySet1((_489_v5).Cardinality(), 19)
			(_nw92).ArraySet1(_dafny.IntOfUint32((_520_v31).Cardinality()), 20)
			(_nw92).ArraySet1((_this).F27(), 21)
			(_nw92).ArraySet1((_this).F27(), 22)
			(_nw92).ArraySet1((_this).F27(), 23)
			_521_v32 = _nw92
			var _522_v33 D5
			_ = _522_v33
			_522_v33 = Companion_D5_.Create_DC9_(_521_v32)
			var _source9 D5 = _522_v33
			_ = _source9
			if _source9.Is_DC10() {
				var _523___mcc_h1 bool = _source9.Get_().(D5_DC10).Cf14
				_ = _523___mcc_h1
				var _524___mcc_h2 bool = _source9.Get_().(D5_DC10).Cf15
				_ = _524___mcc_h2
				var _525_cf15 bool = _524___mcc_h2
				_ = _525_cf15
				var _526_cf14 bool = _523___mcc_h1
				_ = _526_cf14
				var _527_v34 _dafny.Set
				_ = _527_v34
				_527_v34 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ofy"), _483_v1)
				var _528_v35 D3
				_ = _528_v35
				_528_v35 = Companion_D3_.Create_DC6_((func() _dafny.Int {
					if p0 {
						return (_this).F27()
					}
					return (_this).F27()
				})(), ((_527_v34).Union(_527_v34)).Cardinality(), true)
				var _529_v36 _dafny.Array
				_ = _529_v36
				var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
				_ = _nw93
				_529_v36 = _nw93
				var _530_v37 _dafny.Array
				_ = _530_v37
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_14
				var _nw94 _dafny.Array
				_ = _nw94
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw94 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) _dafny.MultiSet = (func(_531_v30 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_532_i3 _dafny.Int) _dafny.MultiSet {
							return _531_v30
						}
					})(_519_v30)
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
				_530_v37 = _nw94
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_529_v36), 0))
				_ = _index85
				(_529_v36).ArraySet1(_530_v37, (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_529_v36), 0))
				_ = _index86
				var _rhs94 bool = !(p1)
				_ = _rhs94
				var _rhs95 D3 = _528_v35
				_ = _rhs95
				var _rhs96 _dafny.Array = _530_v37
				_ = _rhs96
				var _lhs87 *GlobalState = globalState
				_ = _lhs87
				var _lhs88 _dafny.Array = _529_v36
				_ = _lhs88
				var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_529_v36), 0))
				_ = _lhs89
				_lhs87.F18 = _rhs94
				_528_v35 = _rhs95
				(_lhs88).ArraySet1(_rhs96, (_lhs89).Int())
				var _533_v38 _dafny.Map
				_ = _533_v38
				_533_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.IntOfInt64(200))
				(globalState).F18 = ((_519_v30).Intersection(_dafny.MultiSetOf((_this).F27(), (_this).F27(), (_this).F27(), (_this).F27(), Companion_Default___.Fm5((_this).F27(), (_this).F27(), globalState)))).Equals(_dafny.MultiSetOf((_this).F27(), (_533_v38).Cardinality()))
				var _534_v39 _dafny.Map
				_ = _534_v39
				_534_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p1)
				var _535_v40 _dafny.Set
				_ = _535_v40
				_535_v40 = _dafny.SetOf(_534_v39, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p0), _534_v39, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_533_v38).Contains((_this).F27()) {
						return (_533_v38).Get((_this).F27()).(_dafny.Int)
					}
					return (_this).F27()
				})(), (_this).F28()))
				var _536_v41 D3
				_ = _536_v41
				_536_v41 = Companion_D3_.Create_DC4_(_535_v40)
				var _537_v42 _dafny.Map
				_ = _537_v42
				_537_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _536_v41)
				_537_v42 = (_537_v42).Update((_this).F27(), Companion_D3_.Create_DC4_(_535_v40))
				var _538_v43 _dafny.MultiSet
				_ = _538_v43
				_538_v43 = _dafny.MultiSetOf(_521_v32, _521_v32, _521_v32)
				var _rhs97 bool = _526_cf14
				_ = _rhs97
				var _rhs98 bool = !(p1)
				_ = _rhs98
				var _rhs99 bool = p0
				_ = _rhs99
				var _rhs100 _dafny.MultiSet = _dafny.MultiSetOf((_this).F27(), (_this).F27(), ((_538_v43).Intersection(_538_v43)).Cardinality())
				_ = _rhs100
				var _lhs90 *GlobalState = globalState
				_ = _lhs90
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				_lhs90.F18 = _rhs97
				_526_cf14 = _rhs98
				_lhs91.F18 = _rhs99
				_519_v30 = _rhs100
			} else {
				var _539___mcc_h3 _dafny.Array = _source9.Get_().(D5_DC9).Cf13
				_ = _539___mcc_h3
				var _540_cf13 _dafny.Array = _539___mcc_h3
				_ = _540_cf13
				_484_v2 = _484_v2
				var _541_v44 _dafny.Sequence
				_ = _541_v44
				_541_v44 = _dafny.SeqOf(_dafny.IntOfUint32((_493_v9).Cardinality()))
				var _542_v45 _dafny.Map
				_ = _542_v45
				_542_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v44, (_dafny.IntOfUint32((_483_v1).Cardinality())).Cmp((_this).F27()) >= 0)
				_542_v45 = (_542_v45).Update(_dafny.Companion_Sequence_.Concatenate(_541_v44, _541_v44), true)
				(globalState).F22 = _dafny.IntOfInt64(409)
				(globalState).F0 = (_this).F27()
			}
			(globalState).F15 = (_dafny.IntOfUint32((_483_v1).Cardinality())).Plus(((_this).F27()).Times((_this).F27()))
			var _543_v46 _dafny.Sequence
			_ = _543_v46
			_543_v46 = _dafny.SeqOf((_this).F27(), (_this).F27())
			var _544_v47 D2
			_ = _544_v47
			_544_v47 = Companion_D2_.Create_DC2_((func() bool {
				if p1 {
					return p0
				}
				return Companion_Default___.Fm0(_dafny.IntOfUint32((_543_v46).Cardinality()), p0, p0, (_this).F27(), globalState)
			})())
			var _source10 D2 = _544_v47
			_ = _source10
			if _source10.Is_DC3() {
				var _545___mcc_h4 _dafny.Int = _source10.Get_().(D2_DC3).Cf3
				_ = _545___mcc_h4
				var _546___mcc_h5 _dafny.Int = _source10.Get_().(D2_DC3).Cf4
				_ = _546___mcc_h5
				var _547_cf4 _dafny.Int = _546___mcc_h5
				_ = _547_cf4
				var _548_cf3 _dafny.Int = _545___mcc_h4
				_ = _548_cf3
				var _549_v48 _dafny.Map
				_ = _549_v48
				_549_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _521_v32)
				_549_v48 = (_549_v48).Update(p1, _521_v32)
				var _550_v49 *C0
				_ = _550_v49
				var _nw95 *C0 = New_C0_()
				_ = _nw95
				_nw95.Ctor__()
				_550_v49 = _nw95
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_521_v32), 0))
				_ = _index87
				(_521_v32).ArraySet1((_548_cf3).Plus((_dafny.SetOf(false)).Cardinality()), (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_521_v32), 0))
				_ = _index88
				(_521_v32).ArraySet1(_547_cf4, (_index88).Int())
				var _551_v50 _dafny.Sequence
				_ = _551_v50
				_551_v50 = _dafny.SeqOf(_483_v1, _dafny.UnicodeSeqOfUtf8Bytes("sjtnwy"), _483_v1)
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_521_v32), 0))
				_ = _index89
				var _rhs101 _dafny.Int = (_this).F27()
				_ = _rhs101
				var _rhs102 _dafny.Int = (_521_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_521_v32), 0))).Int()).(_dafny.Int)
				_ = _rhs102
				var _rhs103 _dafny.Int = (_this).F27()
				_ = _rhs103
				var _rhs104 _dafny.CodePoint = _482_v0
				_ = _rhs104
				var _rhs105 _dafny.Sequence = (_551_v50).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_543_v46).Cardinality()), _dafny.IntOfInt64(682)), _dafny.IntOfUint32((_551_v50).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs105
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				var _lhs93 _dafny.Array = _521_v32
				_ = _lhs93
				var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.ArrayLen((_521_v32), 0))
				_ = _lhs94
				var _lhs95 *GlobalState = globalState
				_ = _lhs95
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				_lhs92.F16 = _rhs101
				(_lhs93).ArraySet1(_rhs102, (_lhs94).Int())
				_lhs95.F22 = _rhs103
				_482_v0 = _rhs104
				_lhs96.F3 = _rhs105
			} else {
				var _552___mcc_h6 bool = _source10.Get_().(D2_DC2).Cf2
				_ = _552___mcc_h6
				var _553_cf2 bool = _552___mcc_h6
				_ = _553_cf2
				(globalState).F16 = (_this).F27()
				var _554_v51 _dafny.Map
				_ = _554_v51
				_554_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_483_v1, p1)
				(globalState).F18 = (func() bool {
					if (_554_v51).Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mffkpiqe")).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1)).Cardinality()))).Uint32(), _482_v0)) {
						return (_554_v51).Get(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mffkpiqe")).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1)).Cardinality()))).Uint32(), _482_v0)).(bool)
					}
					return (_this).F28()
				})()
				var _555_v52 _dafny.Array
				_ = _555_v52
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_15
				var _nw96 _dafny.Array
				_ = _nw96
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw96 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) bool = func(_556_i4 _dafny.Int) bool {
						return ((_this).F27()).Cmp(_dafny.IntOfInt64(73)) != 0
					}
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw96 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw96).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw96).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_555_v52 = _nw96
				_555_v52 = _555_v52
				var _557_v53 T0
				_ = _557_v53
				var _nw97 *C1 = New_C1_()
				_ = _nw97
				_nw97.Ctor__(_dafny.CodePoint('a'), (_this).F27())
				_557_v53 = _nw97
				var _558_v54 D6
				_ = _558_v54
				_558_v54 = Companion_D6_.Create_DC11_(_557_v53)
				var _559_v55 _dafny.Sequence
				_ = _559_v55
				_559_v55 = _dafny.SeqOf(_557_v53)
				var _560_v56 _dafny.Sequence
				_ = _560_v56
				_560_v56 = _dafny.SeqOf(_dafny.SeqOf((_558_v54).Dtor_cf16(), _557_v53), _559_v55, _559_v55, _559_v55, _559_v55)
				var _561_v57 _dafny.MultiSet
				_ = _561_v57
				_561_v57 = _dafny.MultiSetOf(_559_v55, _559_v55, _559_v55)
				var _562_v58 _dafny.Array
				_ = _562_v58
				var _nwElement0_19 _dafny.MultiSet = _dafny.MultiSetFromSeq(_560_v56)
				_ = _nwElement0_19
				var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(10))
				_ = _nw98
				(_nw98).ArraySet1(_nwElement0_19, 0)
				(_nw98).ArraySet1(_561_v57, 1)
				(_nw98).ArraySet1(_561_v57, 2)
				(_nw98).ArraySet1(_561_v57, 3)
				(_nw98).ArraySet1(_561_v57, 4)
				(_nw98).ArraySet1(_dafny.MultiSetFromSeq(_560_v56), 5)
				(_nw98).ArraySet1(_561_v57, 6)
				(_nw98).ArraySet1((func() _dafny.MultiSet {
					if p1 {
						return _561_v57
					}
					return _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_559_v55, (Companion_Default___.SafeIndex((_518_v29).Cardinality(), _dafny.IntOfUint32((_559_v55).Cardinality()))).Uint32(), _557_v53), _559_v55, _dafny.SeqOf(_557_v53, _557_v53), _559_v55, _559_v55)
				})(), 7)
				(_nw98).ArraySet1((_dafny.MultiSetFromSeq(_560_v56)).Update(_559_v55, Companion_Default___.Abs(_557_v53.F25())), 8)
				(_nw98).ArraySet1(_561_v57, 9)
				_562_v58 = _nw98
				var _563_v59 _dafny.Sequence
				_ = _563_v59
				_563_v59 = _559_v55
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_562_v58), 0))
				_ = _index90
				(_562_v58).ArraySet1(_dafny.MultiSetOf(_dafny.SeqOf(_557_v53), (_563_v59), _559_v55), (_index90).Int())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_562_v58), 0))
				_ = _index91
				var _rhs106 _dafny.MultiSet = _561_v57
				_ = _rhs106
				var _rhs107 _dafny.MultiSet = _490_v6
				_ = _rhs107
				var _lhs97 _dafny.Array = _562_v58
				_ = _lhs97
				var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_562_v58), 0))
				_ = _lhs98
				(_lhs97).ArraySet1(_rhs106, (_lhs98).Int())
				_490_v6 = _rhs107
			}
			var _564_v60 T0
			_ = _564_v60
			var _nw99 *C1 = New_C1_()
			_ = _nw99
			_nw99.Ctor__(_482_v0, (_this).F27())
			_564_v60 = _nw99
			var _565_v61 _dafny.Map
			_ = _565_v61
			_565_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_564_v60, p0)
			var _566_v62 D6
			_ = _566_v62
			_566_v62 = Companion_D6_.Create_DC12_((func() bool {
				if (_565_v61).Contains(_564_v60) {
					return (_565_v61).Get(_564_v60).(bool)
				}
				return !((_this).F28())
			})(), Companion_Default___.Fm5(_dafny.IntOfInt64(650), (_dafny.Zero).Minus(Companion_Default___.Fm5(_564_v60.F25(), _564_v60.F25(), globalState)), globalState))
			(globalState).F18 = (_566_v62).Dtor_cf17()
			(globalState).F18 = ((_dafny.Zero).Minus(((_this).F27()).Times(_564_v60.F25()))).Cmp(((_this).F27()).Minus((_this).F27())) <= 0
		} else {
			(globalState).F18 = !((func() bool {
				if (_this).F28() {
					return !(_517_v28).Equals(_517_v28)
				}
				return (p1) || (!((_this).F28()))
			})())
			(globalState).F18 = !(p1)
			if !(false) || ((_this).F28()) {
				var _567_v63 *C0
				_ = _567_v63
				var _nw100 *C0 = New_C0_()
				_ = _nw100
				_nw100.Ctor__()
				_567_v63 = _nw100
				var _568_v64 D4
				_ = _568_v64
				_568_v64 = Companion_D4_.Create_DC8_()
				var _569_v65 _dafny.Map
				_ = _569_v65
				_569_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), Companion_Default___.Fm0(_dafny.IntOfInt64(-11), p0, p1, (_this).F27(), globalState))
				var _570_v66 _dafny.MultiSet
				_ = _570_v66
				_570_v66 = _dafny.MultiSetOf((_569_v65).Cardinality())
				var _571_v67 _dafny.Map
				_ = _571_v67
				_571_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_570_v66).Union(_570_v66), p1)
				var _572_v68 _dafny.Sequence
				_ = _572_v68
				_572_v68 = _dafny.SeqOf((_this).F27())
				var _rhs108 D4 = Companion_D4_.Create_DC8_()
				_ = _rhs108
				var _rhs109 _dafny.Map = _571_v67
				_ = _rhs109
				var _rhs110 _dafny.Int = (((_this).F27()).Minus(_dafny.IntOfInt64(337))).Plus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_572_v68).Cardinality()), (_this).F27()))
				_ = _rhs110
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				_568_v64 = _rhs108
				_571_v67 = _rhs109
				_lhs99.F15 = _rhs110
				(globalState).F16 = (_this).F27()
				var _573_v69 _dafny.Array
				_ = _573_v69
				var _nwElement0_20 _dafny.CodePoint = _482_v0
				_ = _nwElement0_20
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(5))
				_ = _nw101
				(_nw101).ArraySet1CodePoint(_nwElement0_20, 0)
				(_nw101).ArraySet1CodePoint(Companion_Default___.Fm10((_this).F28(), p0, p0, p1, globalState), 1)
				(_nw101).ArraySet1CodePoint(_482_v0, 2)
				(_nw101).ArraySet1CodePoint(_482_v0, 3)
				(_nw101).ArraySet1CodePoint(_482_v0, 4)
				_573_v69 = _nw101
				var _574_v70 _dafny.Map
				_ = _574_v70
				_574_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _573_v69)
				_574_v70 = (_574_v70).Update((_this).F27(), _573_v69)
				(globalState).F18 = ((_572_v68).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_572_v68).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_this).F27()) >= 0
			} else {
				var _575_v71 T0
				_ = _575_v71
				var _nw102 *C1 = New_C1_()
				_ = _nw102
				_nw102.Ctor__(_482_v0, (_490_v6).Cardinality())
				_575_v71 = _nw102
				var _576_v72 _dafny.Map
				_ = _576_v72
				_576_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5((_this).F27(), (_this).F27(), globalState), _575_v71)
				var _577_v73 D7
				_ = _577_v73
				_577_v73 = Companion_D7_.Create_DC14_(p0, (_576_v72).Cardinality(), _575_v71.F25())
				var _578_v74 _dafny.Map
				_ = _578_v74
				_578_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_D3_.Create_DC6_((_this).F27(), (_this).F27(), (_577_v73).Dtor_cf20()))
				var _579_v75 D9
				_ = _579_v75
				_579_v75 = Companion_D9_.Create_DC16_(Companion_Default___.Fm13(_483_v1, globalState))
				_578_v74 = (_579_v75).Dtor_cf24()
				(globalState).F2 = _dafny.IntOfUint32((_483_v1).Cardinality())
				var _580_v76 _dafny.Array
				_ = _580_v76
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
				_ = _nw103
				_580_v76 = _nw103
				var _581_v77 _dafny.Array
				_ = _581_v77
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_16
				var _nw104 _dafny.Array
				_ = _nw104
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw104 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) bool = func(_582_i5 _dafny.Int) bool {
						return false
					}
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw104 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw104).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw104).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_581_v77 = _nw104
				var _583_v78 _dafny.Sequence
				_ = _583_v78
				_583_v78 = _dafny.SeqOf(_581_v77)
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_580_v76), 0))
				_ = _index92
				(_580_v76).ArraySet1((func() _dafny.Sequence {
					if p1 {
						return _583_v78
					}
					return _583_v78
				})(), (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_580_v76), 0))
				_ = _index93
				var _rhs111 bool = !(p1)
				_ = _rhs111
				var _rhs112 bool = !(p0)
				_ = _rhs112
				var _rhs113 _dafny.Int = (_this).F27()
				_ = _rhs113
				var _rhs114 _dafny.Sequence = _583_v78
				_ = _rhs114
				var _lhs100 *GlobalState = globalState
				_ = _lhs100
				var _lhs101 *GlobalState = globalState
				_ = _lhs101
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				var _lhs103 _dafny.Array = _580_v76
				_ = _lhs103
				var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_580_v76), 0))
				_ = _lhs104
				_lhs100.F18 = _rhs111
				_lhs101.F18 = _rhs112
				_lhs102.F22 = _rhs113
				(_lhs103).ArraySet1(_rhs114, (_lhs104).Int())
				var _584_v79 _dafny.Map
				_ = _584_v79
				_584_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F28())
				var _585_v80 _dafny.Map
				_ = _585_v80
				_585_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_584_v79).Update(_575_v71.F25(), p1))
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_581_v77), 0))
				_ = _index94
				(_581_v77).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _584_v79)).Equals(_585_v80), (_index94).Int())
				var _586_v81 D3
				_ = _586_v81
				_586_v81 = Companion_D3_.Create_DC6_((_584_v79).Cardinality(), (_this).F27(), p0)
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_581_v77), 0))
				_ = _index95
				var _rhs115 _dafny.Int = (_dafny.IntOfInt64(-589)).Times((_dafny.IntOfInt64(615)).Minus((_this).F27()))
				_ = _rhs115
				var _rhs116 _dafny.Int = (_586_v81).Dtor_cf9()
				_ = _rhs116
				var _rhs117 bool = p0
				_ = _rhs117
				var _rhs118 _dafny.Int = (_575_v71.F25()).Times((_this).F27())
				_ = _rhs118
				var _lhs105 T0 = _575_v71
				_ = _lhs105
				var _lhs106 *GlobalState = globalState
				_ = _lhs106
				var _lhs107 _dafny.Array = _581_v77
				_ = _lhs107
				var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_581_v77), 0))
				_ = _lhs108
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				_lhs105.F25_set_(_rhs115)
				_lhs106.F0 = _rhs116
				(_lhs107).ArraySet1(_rhs117, (_lhs108).Int())
				_lhs109.F22 = _rhs118
				var _587_v82 D7
				_ = _587_v82
				_587_v82 = Companion_D7_.Create_DC13_(_584_v79)
				var _588_v83 _dafny.MultiSet
				_ = _588_v83
				_588_v83 = _dafny.MultiSetOf(_587_v82, _587_v82)
				var _589_v84 *C1
				_ = _589_v84
				var _nw105 *C1 = New_C1_()
				_ = _nw105
				_nw105.Ctor__(_482_v0, (func() _dafny.Int {
					if (_588_v83).Contains(_587_v82) {
						return (_588_v83).Multiplicity(_587_v82)
					}
					return _dafny.IntOfUint32((_483_v1).Cardinality())
				})())
				_589_v84 = _nw105
			}
			var _590_v85 _dafny.Set
			_ = _590_v85
			_590_v85 = _dafny.SetOf(p0, p0)
			var _rhs119 bool = (_this).F28()
			_ = _rhs119
			var _rhs120 bool = p1
			_ = _rhs120
			var _rhs121 _dafny.Int = (_dafny.Zero).Minus((_this).F27())
			_ = _rhs121
			var _rhs122 bool = Companion_Default___.Fm0((_dafny.Zero).Minus(((_this).F27()).Plus((_this).F27())), p1, (_590_v85).IsDisjointFrom((Companion_Default___.Fm14((_this).F28(), globalState)).Dtor_cf26()), (_this).F27(), globalState)
			_ = _rhs122
			var _lhs110 *GlobalState = globalState
			_ = _lhs110
			var _lhs111 *GlobalState = globalState
			_ = _lhs111
			var _lhs112 *GlobalState = globalState
			_ = _lhs112
			var _lhs113 *GlobalState = globalState
			_ = _lhs113
			_lhs110.F18 = _rhs119
			_lhs111.F18 = _rhs120
			_lhs112.F22 = _rhs121
			_lhs113.F18 = _rhs122
			var _591_v86 _dafny.MultiSet
			_ = _591_v86
			_591_v86 = _dafny.MultiSetOf(_482_v0)
			if (Companion_Default___.Fm5(_dafny.IntOfInt64(-135), (_591_v86).Cardinality(), globalState)).Cmp((_this).F27()) >= 0 {
				(globalState).F18 = !(p0)
				(globalState).F18 = p1
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_487_v3), 0))
				_ = _index96
				(_487_v3).ArraySet1(_484_v2, (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_487_v3), 0))
				_ = _index97
				(_487_v3).ArraySet1(_484_v2, (_index97).Int())
				var _592_v87 D5
				_ = _592_v87
				_592_v87 = Companion_D5_.Create_DC10_(p1, p1)
				var _593_v88 _dafny.Map
				_ = _593_v88
				_593_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _483_v1)
				var _594_v89 _dafny.Array
				_ = _594_v89
				var _nwElement0_21 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1)).Cardinality()))).Uint32(), _482_v0)
				_ = _nwElement0_21
				var _nw106 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(9))
				_ = _nw106
				(_nw106).ArraySet1(_nwElement0_21, 0)
				(_nw106).ArraySet1(_483_v1, 1)
				(_nw106).ArraySet1(_483_v1, 2)
				(_nw106).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("htgefcx"), 3)
				(_nw106).ArraySet1(((_487_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_487_v3), 0))).Int()).(_dafny.Sequence)), 4)
				(_nw106).ArraySet1(_483_v1, 5)
				(_nw106).ArraySet1(_483_v1, 6)
				(_nw106).ArraySet1((func() _dafny.Sequence {
					if (_593_v88).Contains((_this).F27()) {
						return (_593_v88).Get((_this).F27()).(_dafny.Sequence)
					}
					return Companion_Default___.Fm15((_this).F27(), (_this).F27(), (_dafny.Zero).Minus((_this).F27()), globalState)
				})(), 7)
				(_nw106).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1), 8)
				_594_v89 = _nw106
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_594_v89), 0))
				_ = _index98
				(_594_v89).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1), (_index98).Int())
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_594_v89), 0))
				_ = _index99
				var _rhs123 D5 = func(_pat_let8_0 D5) D5 {
					return func(_595_dt__update__tmp_h2 D5) D5 {
						return func(_pat_let9_0 bool) D5 {
							return func(_596_dt__update_hcf15_h0 bool) D5 {
								return Companion_D5_.Create_DC10_((_595_dt__update__tmp_h2).Dtor_cf14(), _596_dt__update_hcf15_h0)
							}(_pat_let9_0)
						}((_this).F28())
					}(_pat_let8_0)
				}(_592_v87)
				_ = _rhs123
				var _rhs124 _dafny.Int = (_this).F27()
				_ = _rhs124
				var _rhs125 bool = p0
				_ = _rhs125
				var _rhs126 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_483_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-798))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}(func(_597_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})))
				_ = _rhs126
				var _lhs114 *GlobalState = globalState
				_ = _lhs114
				var _lhs115 *GlobalState = globalState
				_ = _lhs115
				var _lhs116 _dafny.Array = _594_v89
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_594_v89), 0))
				_ = _lhs117
				_592_v87 = _rhs123
				_lhs114.F10 = _rhs124
				_lhs115.F18 = _rhs125
				(_lhs116).ArraySet1(_rhs126, (_lhs117).Int())
				var _598_v90 _dafny.Map
				_ = _598_v90
				_598_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0) || (p1), _dafny.Companion_Sequence_.Concatenate(_493_v9, _493_v9))
				_598_v90 = (_598_v90).Update(false, _493_v9)
			} else {
				var _599_v91 _dafny.Array
				_ = _599_v91
				var _nwElement0_22 _dafny.Sequence = _dafny.SeqOf(p1, p0, p1)
				_ = _nwElement0_22
				var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(27))
				_ = _nw107
				(_nw107).ArraySet1(_nwElement0_22, 0)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_493_v9, _dafny.SeqOf(p1, p1)), 1)
				(_nw107).ArraySet1(_493_v9, 2)
				(_nw107).ArraySet1(_493_v9, 3)
				(_nw107).ArraySet1(_493_v9, 4)
				(_nw107).ArraySet1(_493_v9, 5)
				(_nw107).ArraySet1(_493_v9, 6)
				(_nw107).ArraySet1(_493_v9, 7)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_493_v9, Companion_Default___.Fm7(globalState)), 8)
				(_nw107).ArraySet1(_dafny.SeqOf(p0, (_this).F28(), p0, p0), 9)
				(_nw107).ArraySet1(_493_v9, 10)
				(_nw107).ArraySet1((func() _dafny.Sequence {
					if p1 {
						return _493_v9
					}
					return _493_v9
				})(), 11)
				(_nw107).ArraySet1(_493_v9, 12)
				(_nw107).ArraySet1(_493_v9, 13)
				(_nw107).ArraySet1(_493_v9, 14)
				(_nw107).ArraySet1(_493_v9, 15)
				(_nw107).ArraySet1(_493_v9, 16)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1), _493_v9), 17)
				(_nw107).ArraySet1(_493_v9, 18)
				(_nw107).ArraySet1(_493_v9, 19)
				(_nw107).ArraySet1(_493_v9, 20)
				(_nw107).ArraySet1(_493_v9, 21)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_493_v9, _493_v9), 22)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_493_v9, _493_v9), 23)
				(_nw107).ArraySet1(_493_v9, 24)
				(_nw107).ArraySet1(_493_v9, 25)
				(_nw107).ArraySet1(_dafny.SeqOf(p0), 26)
				_599_v91 = _nw107
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_599_v91), 0))
				_ = _index100
				(_599_v91).ArraySet1(_493_v9, (_index100).Int())
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_599_v91), 0))
				_ = _index101
				(_599_v91).ArraySet1(_493_v9, (_index101).Int())
				(globalState).F18 = p1
				var _600_v92 _dafny.Map
				_ = _600_v92
				_600_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _dafny.Companion_Sequence_.Update(_483_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-629), _dafny.IntOfUint32((_483_v1).Cardinality()))).Uint32(), _482_v0))
				_600_v92 = (_600_v92).Update(p1, _dafny.Companion_Sequence_.Concatenate(_483_v1, _483_v1))
				(globalState).F0 = (func() _dafny.Int {
					if !((_this).F28()) {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("html"), _dafny.UnicodeSeqOfUtf8Bytes("i"))).Cardinality())
					}
					return (_dafny.IntOfUint32((_dafny.SeqOf(p0, (_this).F28())).Cardinality())).Times(_dafny.IntOfInt64(86))
				})()
				(globalState).F18 = !((_this).F28())
			}
		}
		var _601_v93 _dafny.Sequence
		_ = _601_v93
		_601_v93 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("fyp"), _483_v1)
		var _602_v94 D5
		_ = _602_v94
		_602_v94 = Companion_D5_.Create_DC10_(p1, false)
		if !(!_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_483_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}((func(_603_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_604_i7 _dafny.Int) _dafny.CodePoint {
				return _603_v0
			}
		})(_482_v0)))), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_483_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg42 _dafny.Int) interface{} {
				return coer42(arg42)
			}
		}((func(_605_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_606_i7 _dafny.Int) _dafny.CodePoint {
				return _605_v0
			}
		})(_482_v0))))).Cardinality()))).Uint32(), _482_v0), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_601_v93).Select((Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_601_v93).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_483_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg43 _dafny.Int) interface{} {
				return coer43(arg43)
			}
		}((func(_607_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_608_i7 _dafny.Int) _dafny.CodePoint {
				return _607_v0
			}
		})(_482_v0)))), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_483_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg44 _dafny.Int) interface{} {
				return coer44(arg44)
			}
		}((func(_609_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_610_i7 _dafny.Int) _dafny.CodePoint {
				return _609_v0
			}
		})(_482_v0))))).Cardinality()))).Uint32(), _482_v0)).Cardinality()))).Uint32(), _482_v0), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_602_v94).Dtor_cf15() {
				return (_483_v1)
			}
			return _483_v1
		})(), (Companion_Default___.SafeIndex((_this).F27(), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_602_v94).Dtor_cf15() {
				return (_483_v1)
			}
			return _483_v1
		})()).Cardinality()))).Uint32(), _482_v0))) {
			(globalState).F15 = (_this).F27()
			(globalState).F18 = false
			var _611_v95 _dafny.Sequence
			_ = _611_v95
			_611_v95 = _dafny.SeqOf((_this).F27())
			var _612_v96 _dafny.Array
			_ = _612_v96
			var _nwElement0_23 bool = p1
			_ = _nwElement0_23
			var _nw108 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(15))
			_ = _nw108
			(_nw108).ArraySet1(_nwElement0_23, 0)
			(_nw108).ArraySet1((_this).F28(), 1)
			(_nw108).ArraySet1(Companion_Default___.Fm0((_this).F27(), p1, p1, (_this).F27(), globalState), 2)
			(_nw108).ArraySet1(p1, 3)
			(_nw108).ArraySet1(p0, 4)
			(_nw108).ArraySet1(!((_this).F28()), 5)
			(_nw108).ArraySet1(p1, 6)
			(_nw108).ArraySet1((_this).F28(), 7)
			(_nw108).ArraySet1(p1, 8)
			(_nw108).ArraySet1(p0, 9)
			(_nw108).ArraySet1((_this).F28(), 10)
			(_nw108).ArraySet1(!(p1), 11)
			(_nw108).ArraySet1((_this).F28(), 12)
			(_nw108).ArraySet1(p1, 13)
			(_nw108).ArraySet1((_602_v94).Dtor_cf15(), 14)
			_612_v96 = _nw108
			var _613_v97 _dafny.Map
			_ = _613_v97
			_613_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_612_v96, p1)
			var _614_v98 _dafny.Map
			_ = _614_v98
			_614_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
				if (_490_v6).Contains((_this).F28()) {
					return (_490_v6).Multiplicity((_this).F28())
				}
				return (_this).F27()
			})()).Minus(_dafny.IntOfUint32((_611_v95).Cardinality())), (_613_v97).Contains(_612_v96))
			var _615_v99 _dafny.Array
			_ = _615_v99
			var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw109
			_615_v99 = _nw109
			var _616_v100 _dafny.Map
			_ = _616_v100
			_616_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_615_v99, Companion_Default___.Fm0((_this).F27(), !(false), (_this).F28(), (_this).F27(), globalState))
			_614_v98 = (_614_v98).Update((_this).F27(), ((_dafny.Zero).Minus((_this).F27())).Cmp((_616_v100).Cardinality()) < 0)
			var _rhs127 _dafny.CodePoint = _dafny.CodePoint('b')
			_ = _rhs127
			var _rhs128 _dafny.Sequence = _483_v1
			_ = _rhs128
			var _lhs118 *GlobalState = globalState
			_ = _lhs118
			_482_v0 = _rhs127
			_lhs118.F3 = _rhs128
			if p1 {
				(globalState).F11 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((Companion_Default___.Fm15(_dafny.IntOfUint32((_483_v1).Cardinality()), (_this).F27(), Companion_Default___.Fm5(_dafny.IntOfUint32((_611_v95).Cardinality()), (_this).F27(), globalState), globalState)).Cardinality()), (func() _dafny.Int {
					if (_490_v6).Contains(p0) {
						return (_490_v6).Multiplicity(p0)
					}
					return (_this).F27()
				})()))
				var _617_v101 _dafny.Array
				_ = _617_v101
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
				_ = _nw110
				_617_v101 = _nw110
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_617_v101), 0))
				_ = _index102
				(_617_v101).ArraySet1(_483_v1, (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_617_v101), 0))
				_ = _index103
				(_617_v101).ArraySet1(_483_v1, (_index103).Int())
				(globalState).F10 = (_dafny.Zero).Minus((_this).F27())
				(globalState).F18 = false
				var _618_v102 _dafny.MultiSet
				_ = _618_v102
				_618_v102 = _dafny.MultiSetOf(_492_v8, _493_v9, _492_v8)
				var _619_v103 _dafny.MultiSet
				_ = _619_v103
				_619_v103 = _dafny.MultiSetOf((_this).F27(), (_618_v102).Cardinality())
				_619_v103 = _619_v103
			} else {
				(globalState).F18 = (p1) == (p1)
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_615_v99), 0))
				_ = _index104
				(_615_v99).ArraySet1(_dafny.IntOfInt64(213), (_index104).Int())
				var _620_v104 _dafny.MultiSet
				_ = _620_v104
				_620_v104 = _dafny.MultiSetOf((_this).F27())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_615_v99), 0))
				_ = _index105
				(_615_v99).ArraySet1((func() _dafny.Int {
					if (_this).F28() {
						return (_this).F27()
					}
					return (func() _dafny.Int {
						if p1 {
							return (_620_v104).Cardinality()
						}
						return (_this).F27()
					})()
				})(), (_index105).Int())
				(globalState).F2 = Companion_Default___.SafeDivisionInt(((_615_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_615_v99), 0))).Int()).(_dafny.Int)).Minus((_dafny.Zero).Minus((_615_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_615_v99), 0))).Int()).(_dafny.Int))), Companion_Default___.SafeModuloInt((_615_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_615_v99), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(588)))
				var _621_v105 D3
				_ = _621_v105
				_621_v105 = Companion_D3_.Create_DC6_(_dafny.IntOfUint32((_483_v1).Cardinality()), (_615_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_615_v99), 0))).Int()).(_dafny.Int), p0)
				var _622_v106 D9
				_ = _622_v106
				_622_v106 = Companion_D9_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _621_v105))
				_622_v106 = _622_v106
				var _623_v107 D10
				_ = _623_v107
				_623_v107 = Companion_D10_.Create_DC20_(p1, Companion_Default___.Fm5((_this).F27(), (_615_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_615_v99), 0))).Int()).(_dafny.Int), globalState))
				_623_v107 = _623_v107
			}
		} else {
			_517_v28 = (_517_v28).Merge((_517_v28).Merge(_517_v28))
			var _624_v108 D2
			_ = _624_v108
			_624_v108 = Companion_D2_.Create_DC3_((_this).F27(), (_this).F27())
			var _625_v109 _dafny.Map
			_ = _625_v109
			_625_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_624_v108, (_dafny.Zero).Minus((_this).F27()))
			_625_v109 = (_625_v109).Update(_624_v108, _dafny.IntOfInt64(446))
			var _626_v110 _dafny.Array
			_ = _626_v110
			var _nw111 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw111
			_626_v110 = _nw111
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_626_v110), 0))
			_ = _index106
			(_626_v110).ArraySet1(p1, (_index106).Int())
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_626_v110), 0))
			_ = _index107
			(_626_v110).ArraySet1(p0, (_index107).Int())
			var _627_v111 _dafny.Sequence
			_ = _627_v111
			_627_v111 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!((_this).F28()))).Cardinality()))
			(globalState).F0 = (_627_v111).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_493_v9).Cardinality()), _dafny.IntOfUint32((_627_v111).Cardinality()))).Uint32()).(_dafny.Int)
			var _628_v112 _dafny.Map
			_ = _628_v112
			_628_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F27())
			(globalState).F15 = ((_628_v112).Cardinality()).Times((_this).F27())
		}
		var _629_v113 *C1
		_ = _629_v113
		var _nw112 *C1 = New_C1_()
		_ = _nw112
		_nw112.Ctor__(_482_v0, ((_this).F27()).Minus((_this).F27()))
		_629_v113 = _nw112
		var _630_i8 _dafny.Int
		_ = _630_i8
		_630_i8 = _dafny.Zero
		{
			for ((_490_v6).Difference(((_dafny.MultiSetOf(p1)).Update(p1, Companion_Default___.Abs((_this).F27()))).Update(p0, Companion_Default___.Abs(_dafny.IntOfInt64(103))))).IsDisjointFrom(_490_v6) {
				{
					if (_630_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_630_i8 = (_630_i8).Plus(_dafny.One)
					(globalState).F18 = p1
					(globalState).F15 = _dafny.IntOfUint32((_493_v9).Cardinality())
					(globalState).F18 = ((_490_v6).Difference(_490_v6)).IsDisjointFrom(_490_v6)
					(globalState).F10 = (_this).F27()
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		r0 = func(_source11 D2) _dafny.Int {
			if _source11.Is_DC3() {
				var _631___mcc_h7 _dafny.Int = _source11.Get_().(D2_DC3).Cf3
				_ = _631___mcc_h7
				var _632___mcc_h8 _dafny.Int = _source11.Get_().(D2_DC3).Cf4
				_ = _632___mcc_h8
				var _633_cf4 _dafny.Int = _632___mcc_h8
				_ = _633_cf4
				var _634_cf3 _dafny.Int = _631___mcc_h7
				_ = _634_cf3
				return (_this).F27()
			} else {
				var _635___mcc_h9 bool = _source11.Get_().(D2_DC2).Cf2
				_ = _635___mcc_h9
				var _636_cf2 bool = _635___mcc_h9
				_ = _636_cf2
				return (_dafny.Zero).Minus((_this).F27())
			}
		}(Companion_D2_.Create_DC2_((_this).F28()))
		var _637_v114 _dafny.MultiSet
		_ = _637_v114
		_637_v114 = _dafny.MultiSetOf(_482_v0)
		r1 = (_dafny.Zero).Minus((func() _dafny.Int {
			if (_637_v114).Contains(_482_v0) {
				return (_637_v114).Multiplicity(_482_v0)
			}
			return (_this).F27()
		})())
		return r0, r1
	}
}
func (_this *C2) F27() _dafny.Int {
	{
		return _this._f27
	}
}
func (_this *C2) F28() bool {
	{
		return _this._f28
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f25 _dafny.Int
	_f24 _dafny.CodePoint
	F31  bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f25 = _dafny.Zero
	_this._f24 = _dafny.CodePoint('D')
	_this.F31 = false
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

func (_this *C3) F25() _dafny.Int {
	return _this._f25
}
func (_this *C3) F25_set_(value _dafny.Int) {
	_this._f25 = value
}
func (_this *C3) F24() _dafny.CodePoint {
	return _this._f24
}
func (_this *C3) Ctor__(f31 bool, f24 _dafny.CodePoint, f25 _dafny.Int) {
	{
		(_this).F31 = f31
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C3) Fm3(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pcxk"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(893))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg45 _dafny.Int) interface{} {
				return coer45(arg45)
			}
		}(func(_638_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})))
	}
}
func (_this *C3) Fm18(globalState *GlobalState) bool {
	{
		return !(_this.F31)
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f25 _dafny.Int
	_f24 _dafny.CodePoint
	F29  _dafny.Map
	_f30 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f25 = _dafny.Zero
	_this._f24 = _dafny.CodePoint('D')
	_this.F29 = _dafny.EmptyMap
	_this._f30 = false
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

func (_this *C4) F25() _dafny.Int {
	return _this._f25
}
func (_this *C4) F25_set_(value _dafny.Int) {
	_this._f25 = value
}
func (_this *C4) F24() _dafny.CodePoint {
	return _this._f24
}
func (_this *C4) Ctor__(f29 _dafny.Map, f30 bool, f24 _dafny.CodePoint, f25 _dafny.Int) {
	{
		(_this).F29 = f29
		(_this)._f30 = f30
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C4) Fm3(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("if"), _dafny.UnicodeSeqOfUtf8Bytes("kaxkef")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("e"), _dafny.UnicodeSeqOfUtf8Bytes("oqanv")))
	}
}
func (_this *C4) M5(p0 _dafny.Map, p1 bool, p2 _dafny.Int, globalState *GlobalState) {
	{
		(globalState).F11 = p2
		var _639_i0 _dafny.Int
		_ = _639_i0
		_639_i0 = _dafny.Zero
		{
			for p1 {
				{
					if (_639_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_639_i0 = (_639_i0).Plus(_dafny.One)
					var _640_v0 _dafny.Map
					_ = _640_v0
					_640_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC2_(p1), p2)
					var _641_v1 _dafny.Map
					_ = _641_v1
					_641_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_dafny.Zero).Minus((Companion_Default___.Fm16(_640_v0, globalState)).Dtor_cf29()))
					var _642_v2 _dafny.Sequence
					_ = _642_v2
					_642_v2 = _dafny.UnicodeSeqOfUtf8Bytes("aqdbq")
					var _643_v3 _dafny.Map
					_ = _643_v3
					_643_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_642_v2, Companion_Default___.Fm0(p2, (_this).F30(), p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}(func(_644_i1 _dafny.Int) _dafny.Int {
						return _this.F25()
					}))).Cardinality()), globalState))
					_641_v1 = (_641_v1).Update(((_643_v3).Merge(_643_v3)).Cardinality(), p2)
					var _645_v4 _dafny.Array
					_ = _645_v4
					var _nw113 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
					_ = _nw113
					_645_v4 = _nw113
					var _646_v5 _dafny.Map
					_ = _646_v5
					_646_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _645_v4)
					_645_v4 = (func() _dafny.Array {
						if (_646_v5).Contains((_this).F30()) {
							return (_646_v5).Get((_this).F30()).(_dafny.Array)
						}
						return _645_v4
					})()
					_641_v1 = (_641_v1).Update(_this.F25(), Companion_Default___.SafeDivisionInt(_this.F25(), (_dafny.Zero).Minus(p2)))
					var _647_v6 _dafny.Sequence
					_ = _647_v6
					_647_v6 = _dafny.SeqOf((_this).F30())
					var _648_v7 _dafny.Array
					_ = _648_v7
					var _len0_17 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_17
					var _nw114 _dafny.Array
					_ = _nw114
					if _len0_17.Cmp(_dafny.Zero) == 0 {
						_nw114 = _dafny.NewArray(_len0_17)
					} else {
						var _init17 func(_dafny.Int) _dafny.Map = (func(_649_v1 _dafny.Map) func(_dafny.Int) _dafny.Map {
							return func(_650_i2 _dafny.Int) _dafny.Map {
								return _649_v1
							}
						})(_641_v1)
						_ = _init17
						var _element0_17 = _init17(_dafny.Zero)
						_ = _element0_17
						_nw114 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
						(_nw114).ArraySet1(_element0_17, 0)
						var _nativeLen0_17 = (_len0_17).Int()
						_ = _nativeLen0_17
						for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
							(_nw114).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
						}
					}
					_648_v7 = _nw114
					var _651_v8 _dafny.MultiSet
					_ = _651_v8
					_651_v8 = _dafny.MultiSetOf(_648_v7)
					var _652_v9 _dafny.Set
					_ = _652_v9
					_652_v9 = _dafny.SetOf(!((_this).F30()), (_this).F30())
					var _653_v10 _dafny.MultiSet
					_ = _653_v10
					_653_v10 = _dafny.MultiSetOf(_this.F25())
					var _654_v11 _dafny.Set
					_ = _654_v11
					_654_v11 = _dafny.SetOf(_dafny.IntOfInt64(489), _dafny.IntOfUint32((_642_v2).Cardinality()))
					var _655_v12 _dafny.Map
					_ = _655_v12
					_655_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
					var _656_v13 _dafny.Sequence
					_ = _656_v13
					_656_v13 = _dafny.SeqOf(_655_v12)
					var _657_v14 _dafny.Map
					_ = _657_v14
					_657_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_656_v13).Select((Companion_Default___.SafeIndex(_this.F25(), _dafny.IntOfUint32((_656_v13).Cardinality()))).Uint32()).(_dafny.Map), (_this).F24())
					var _658_v15 _dafny.Sequence
					_ = _658_v15
					_658_v15 = _dafny.SeqOf(p2)
					var _659_v16 _dafny.Map
					_ = _659_v16
					_659_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), (_this).F24())
					var _660_v17 _dafny.Array
					_ = _660_v17
					var _nwElement0_24 _dafny.Int = Companion_Default___.SafeDivisionInt(_this.F25(), _this.F25())
					_ = _nwElement0_24
					var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(27))
					_ = _nw115
					(_nw115).ArraySet1(_nwElement0_24, 0)
					(_nw115).ArraySet1(p2, 1)
					(_nw115).ArraySet1(_dafny.IntOfUint32((_642_v2).Cardinality()), 2)
					(_nw115).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_647_v6).Cardinality()), _this.F25()), 3)
					(_nw115).ArraySet1(_this.F25(), 4)
					(_nw115).ArraySet1(_dafny.IntOfUint32((_642_v2).Cardinality()), 5)
					(_nw115).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
						if (_651_v8).Contains(_648_v7) {
							return (_651_v8).Multiplicity(_648_v7)
						}
						return (_652_v9).Cardinality()
					})(), _dafny.IntOfInt64(108)), 6)
					(_nw115).ArraySet1(_this.F25(), 7)
					(_nw115).ArraySet1((p2).Times(p2), 8)
					(_nw115).ArraySet1((func() _dafny.Int {
						if (_653_v10).Contains(p2) {
							return (_653_v10).Multiplicity(p2)
						}
						return p2
					})(), 9)
					(_nw115).ArraySet1(Companion_Default___.SafeModuloInt((_654_v11).Cardinality(), p2), 10)
					(_nw115).ArraySet1(_dafny.IntOfInt64(-441), 11)
					(_nw115).ArraySet1((_657_v14).Cardinality(), 12)
					(_nw115).ArraySet1(p2, 13)
					(_nw115).ArraySet1(_dafny.IntOfInt64(141), 14)
					(_nw115).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}(func(_661_i3 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(499)
					}))).Cardinality()))), 15)
					(_nw115).ArraySet1(_this.F25(), 16)
					(_nw115).ArraySet1(p2, 17)
					(_nw115).ArraySet1(_this.F25(), 18)
					(_nw115).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg48 _dafny.Int) interface{} {
							return coer48(arg48)
						}
					}(func(_662_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					})), (Companion_Default___.SafeIndex(_this.F25(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg49 _dafny.Int) interface{} {
							return coer49(arg49)
						}
					}(func(_663_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					}))).Cardinality()))).Uint32(), (_this).F24())).Cardinality()), _dafny.IntOfInt64(257))), 19)
					(_nw115).ArraySet1(p2, 20)
					(_nw115).ArraySet1((func() _dafny.Int {
						if (_this).F30() {
							return p2
						}
						return Companion_Default___.Fm17((_this).F30(), true, p1, globalState)
					})(), 21)
					(_nw115).ArraySet1(_dafny.IntOfUint32((_658_v15).Cardinality()), 22)
					(_nw115).ArraySet1((_dafny.Zero).Minus(p2), 23)
					(_nw115).ArraySet1(((_659_v16).Cardinality()).Minus(_dafny.IntOfInt64(624)), 24)
					(_nw115).ArraySet1((_dafny.IntOfInt64(286)).Plus(Companion_Default___.Fm17(p1, (_this).F30(), p1, globalState)), 25)
					(_nw115).ArraySet1(_this.F25(), 26)
					_660_v17 = _nw115
					var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_660_v17), 0))
					_ = _index108
					(_660_v17).ArraySet1(p2, (_index108).Int())
					var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_660_v17), 0))
					_ = _index109
					(_660_v17).ArraySet1(_this.F25(), (_index109).Int())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _664_v18 _dafny.Map
		_ = _664_v18
		_664_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), p1)
		var _665_v19 T1
		_ = _665_v19
		var _nw116 *C3 = New_C3_()
		_ = _nw116
		_nw116.Ctor__(p1, (_this).F24(), (_664_v18).Cardinality())
		_665_v19 = _nw116
		var _666_v20 _dafny.Sequence
		_ = _666_v20
		_666_v20 = _dafny.SeqOf(_665_v19)
		(globalState).F11 = (_dafny.Zero).Minus((_this.F25()).Times(_dafny.IntOfUint32((_dafny.SeqOf(_666_v20, _666_v20, _666_v20, _666_v20)).Cardinality())))
		var _667_v21 _dafny.Map
		_ = _667_v21
		_667_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(682), p1)
		var _668_v22 _dafny.Array
		_ = _668_v22
		var _nwElement0_25 bool = ((func() bool {
			if (_667_v21).Contains(_this.F25()) {
				return (_667_v21).Get(_this.F25()).(bool)
			}
			return p1
		})()) && (p1)
		_ = _nwElement0_25
		var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(3))
		_ = _nw117
		(_nw117).ArraySet1(_nwElement0_25, 0)
		(_nw117).ArraySet1((func() bool {
			if (_664_v18).Contains((_this).F30()) {
				return (_664_v18).Get((_this).F30()).(bool)
			}
			return (_this).F30()
		})(), 1)
		(_nw117).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("uxddo"), _dafny.UnicodeSeqOfUtf8Bytes("teel")), 2)
		_668_v22 = _nw117
		var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_668_v22), 0))
		_ = _index110
		(_668_v22).ArraySet1(p1, (_index110).Int())
		var _669_v23 _dafny.Map
		_ = _669_v23
		_669_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_668_v22), 0))
		_ = _index111
		(_668_v22).ArraySet1(Companion_Default___.Fm0((_669_v23).Cardinality(), p1, p1, _665_v19.F25(), globalState), (_index111).Int())
		_667_v21 = (_667_v21).Update(_665_v19.F25(), (_this.F25()).Cmp((_dafny.Zero).Minus(_665_v19.F25())) >= 0)
		var _670_v24 *C1
		_ = _670_v24
		var _nw118 *C1 = New_C1_()
		_ = _nw118
		_nw118.Ctor__((_665_v19).F24(), Companion_Default___.SafeDivisionInt(_665_v19.F25(), (_664_v18).Cardinality()))
		_670_v24 = _nw118
	}
}
func (_this *C4) F30() bool {
	{
		return _this._f30
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f25 _dafny.Int
	_f24 _dafny.CodePoint
	_f26 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f25 = _dafny.Zero
	_this._f24 = _dafny.CodePoint('D')
	_this._f26 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F25() _dafny.Int {
	return _this._f25
}
func (_this *C5) F25_set_(value _dafny.Int) {
	_this._f25 = value
}
func (_this *C5) F24() _dafny.CodePoint {
	return _this._f24
}
func (_this *C5) Ctor__(f26 bool, f24 _dafny.CodePoint, f25 _dafny.Int) {
	{
		(_this)._f26 = f26
		(_this)._f24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C5) Fm3(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("enhmpxpn"), (func() _dafny.Sequence {
			if (_this).F26() {
				return _dafny.UnicodeSeqOfUtf8Bytes("bob")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("ihqo")
		})())
	}
}
func (_this *C5) M1(p0 _dafny.Map, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 bool = false
		_ = r2
		var _671_v0 *C2
		_ = _671_v0
		var _nw119 *C2 = New_C2_()
		_ = _nw119
		_nw119.Ctor__(_this.F25(), p1)
		_671_v0 = _nw119
		var _672_i0 _dafny.Int
		_ = _672_i0
		_672_i0 = _dafny.Zero
		{
			var _pat_let_tv4 = _671_v0
			_ = _pat_let_tv4
			var _pat_let_tv5 = _671_v0
			_ = _pat_let_tv5
			var _pat_let_tv6 = p1
			_ = _pat_let_tv6
			var _pat_let_tv7 = _671_v0
			_ = _pat_let_tv7
			var _pat_let_tv8 = _671_v0
			_ = _pat_let_tv8
			for !(func(_source12 D10) bool {
				if _source12.Is_DC19() {
					var _702___mcc_h0 _dafny.CodePoint = _source12.Get_().(D10_DC19).Cf27
					_ = _702___mcc_h0
					var _703_cf27 _dafny.CodePoint = _702___mcc_h0
					_ = _703_cf27
					return !(!(_dafny.SetOf((_pat_let_tv4).F28())).Equals(_dafny.SetOf((_pat_let_tv5).F28(), _pat_let_tv6)))
				} else if _source12.Is_DC20() {
					var _704___mcc_h1 bool = _source12.Get_().(D10_DC20).Cf28
					_ = _704___mcc_h1
					var _705___mcc_h2 _dafny.Int = _source12.Get_().(D10_DC20).Cf29
					_ = _705___mcc_h2
					var _706_cf29 _dafny.Int = _705___mcc_h2
					_ = _706_cf29
					var _707_cf28 bool = _704___mcc_h1
					_ = _707_cf28
					return !(true)
				} else if _source12.Is_DC21() {
					var _708___mcc_h3 _dafny.Int = _source12.Get_().(D10_DC21).Cf30
					_ = _708___mcc_h3
					var _709___mcc_h4 bool = _source12.Get_().(D10_DC21).Cf31
					_ = _709___mcc_h4
					var _710___mcc_h5 bool = _source12.Get_().(D10_DC21).Cf32
					_ = _710___mcc_h5
					var _711_cf32 bool = _710___mcc_h5
					_ = _711_cf32
					var _712_cf31 bool = _709___mcc_h4
					_ = _712_cf31
					var _713_cf30 _dafny.Int = _708___mcc_h3
					_ = _713_cf30
					return (_pat_let_tv7).F28()
				} else {
					var _714___mcc_h6 _dafny.Set = _source12.Get_().(D10_DC18).Cf26
					_ = _714___mcc_h6
					var _715_cf26 _dafny.Set = _714___mcc_h6
					_ = _715_cf26
					return (_pat_let_tv8).F28()
				}
			}(Companion_Default___.Fm14((_this).F26(), globalState))) {
				{
					if (_672_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_672_i0 = (_672_i0).Plus(_dafny.One)
					var _673_v1 _dafny.Array
					_ = _673_v1
					var _len0_18 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_18
					var _nw120 _dafny.Array
					_ = _nw120
					if _len0_18.Cmp(_dafny.Zero) == 0 {
						_nw120 = _dafny.NewArray(_len0_18)
					} else {
						var _init18 func(_dafny.Int) bool = func(_674_i1 _dafny.Int) bool {
							return (_this).F26()
						}
						_ = _init18
						var _element0_18 = _init18(_dafny.Zero)
						_ = _element0_18
						_nw120 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
						(_nw120).ArraySet1(_element0_18, 0)
						var _nativeLen0_18 = (_len0_18).Int()
						_ = _nativeLen0_18
						for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
							(_nw120).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
						}
					}
					_673_v1 = _nw120
					var _675_v2 _dafny.Sequence
					_ = _675_v2
					_675_v2 = _dafny.SeqOf(_673_v1)
					var _676_v3 T1
					_ = _676_v3
					var _nw121 *C3 = New_C3_()
					_ = _nw121
					_nw121.Ctor__((_671_v0).F28(), (_this).F24(), ((_671_v0).F27()).Minus((_dafny.Zero).Minus(_this.F25())))
					_676_v3 = _nw121
					var _rhs129 _dafny.Int = _676_v3.F25()
					_ = _rhs129
					var _rhs130 bool = !((_671_v0).F28())
					_ = _rhs130
					var _rhs131 bool = p1
					_ = _rhs131
					var _rhs132 _dafny.Sequence = _675_v2
					_ = _rhs132
					var _rhs133 T1 = _676_v3
					_ = _rhs133
					var _lhs119 *GlobalState = globalState
					_ = _lhs119
					var _lhs120 *GlobalState = globalState
					_ = _lhs120
					_lhs119.F0 = _rhs129
					_lhs120.F18 = _rhs130
					r0 = _rhs131
					_675_v2 = _rhs132
					_676_v3 = _rhs133
					r0 = false
					var _677_v4 _dafny.Sequence
					_ = _677_v4
					_677_v4 = _dafny.UnicodeSeqOfUtf8Bytes("vfafw")
					(globalState).F3 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_677_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-183))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg50 _dafny.Int) interface{} {
							return coer50(arg50)
						}
					}((func(_678_v3 T1) func(_dafny.Int) _dafny.CodePoint {
						return func(_679_i2 _dafny.Int) _dafny.CodePoint {
							return (_678_v3).F24()
						}
					})(_676_v3)))), (func() _dafny.Sequence {
						if (_this).F26() {
							return _677_v4
						}
						return _677_v4
					})())
					if (_671_v0).F28() {
						r2 = (_671_v0).F28()
						(globalState).F18 = !((((_671_v0).F27()).Minus((_dafny.Zero).Minus(_676_v3.F25()))).Cmp((_671_v0).F27()) > 0)
						(globalState).F0 = Companion_Default___.SafeModuloInt((_671_v0).F27(), (_671_v0).F27())
						(globalState).F22 = (_676_v3.F25()).Plus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(783), (_671_v0).F27())))
						var _680_v5 _dafny.Map
						_ = _680_v5
						_680_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_671_v0).F28(), p1)
						var _681_v6 _dafny.Array
						_ = _681_v6
						var _nw122 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
						_ = _nw122
						_681_v6 = _nw122
						var _682_v7 _dafny.Map
						_ = _682_v7
						_682_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_680_v5, _681_v6)
						var _683_v8 _dafny.Sequence
						_ = _683_v8
						_683_v8 = _dafny.SeqOf(_681_v6)
						var _684_v9 *C0
						_ = _684_v9
						var _nw123 *C0 = New_C0_()
						_ = _nw123
						_nw123.Ctor__()
						_684_v9 = _nw123
						var _685_v10 _dafny.Sequence
						_ = _685_v10
						_685_v10 = _dafny.SeqOf(_684_v9)
						var _686_v11 _dafny.Map
						_ = _686_v11
						_686_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_671_v0).F27(), _dafny.IntOfUint32((_685_v10).Cardinality()))
						var _687_v12 D11
						_ = _687_v12
						_687_v12 = Companion_D11_.Create_DC22_(_681_v6)
						var _688_v13 _dafny.Array
						_ = _688_v13
						var _nwElement0_26 _dafny.Map = (_682_v7).Merge(_682_v7)
						_ = _nwElement0_26
						var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(25))
						_ = _nw124
						(_nw124).ArraySet1(_nwElement0_26, 0)
						(_nw124).ArraySet1((_682_v7).Merge(_682_v7), 1)
						(_nw124).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_680_v5, _681_v6)).Update(_680_v5, _681_v6), 2)
						(_nw124).ArraySet1((func() _dafny.Map {
							if false {
								return _682_v7
							}
							return _682_v7
						})(), 3)
						(_nw124).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_680_v5, (_683_v8).Select((Companion_Default___.SafeIndex((_686_v11).Cardinality(), _dafny.IntOfUint32((_683_v8).Cardinality()))).Uint32()).(_dafny.Array)), 4)
						(_nw124).ArraySet1(_682_v7, 5)
						(_nw124).ArraySet1(_682_v7, 6)
						(_nw124).ArraySet1((_682_v7).Merge(_682_v7), 7)
						(_nw124).ArraySet1((_682_v7).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_680_v5, _681_v6)), 8)
						(_nw124).ArraySet1(_682_v7, 9)
						(_nw124).ArraySet1(_682_v7, 10)
						(_nw124).ArraySet1(_682_v7, 11)
						(_nw124).ArraySet1(_682_v7, 12)
						(_nw124).ArraySet1(_682_v7, 13)
						(_nw124).ArraySet1(_682_v7, 14)
						(_nw124).ArraySet1(_682_v7, 15)
						(_nw124).ArraySet1((_682_v7).Merge(_682_v7), 16)
						(_nw124).ArraySet1(_682_v7, 17)
						(_nw124).ArraySet1(_682_v7, 18)
						(_nw124).ArraySet1((_682_v7).Update(_680_v5, (_687_v12).Dtor_cf33()), 19)
						(_nw124).ArraySet1((_682_v7).Merge(_682_v7), 20)
						(_nw124).ArraySet1((_682_v7).Merge((_682_v7).Update(_680_v5, _681_v6)), 21)
						(_nw124).ArraySet1(_682_v7, 22)
						(_nw124).ArraySet1(_682_v7, 23)
						(_nw124).ArraySet1(_682_v7, 24)
						_688_v13 = _nw124
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_688_v13), 0))
						_ = _index112
						(_688_v13).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm19((_671_v0).F28(), globalState), _681_v6)).Merge(_682_v7), (_index112).Int())
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_688_v13), 0))
						_ = _index113
						(_688_v13).ArraySet1((_682_v7).Merge(_682_v7), (_index113).Int())
					} else {
						var _689_v14 D4
						_ = _689_v14
						_689_v14 = Companion_D4_.Create_DC8_()
						var _690_v15 _dafny.MultiSet
						_ = _690_v15
						_690_v15 = _dafny.MultiSetOf(Companion_D4_.Create_DC8_(), _689_v14)
						var _691_v16 _dafny.Sequence
						_ = _691_v16
						_691_v16 = _dafny.SeqOf(_689_v14, _689_v14)
						var _rhs134 bool = p1
						_ = _rhs134
						var _rhs135 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_691_v16)).Union(_dafny.MultiSetOf(_689_v14, _689_v14, _689_v14, _689_v14, _689_v14))
						_ = _rhs135
						var _rhs136 bool = (_this).F26()
						_ = _rhs136
						var _rhs137 bool = p1
						_ = _rhs137
						var _lhs121 *GlobalState = globalState
						_ = _lhs121
						var _lhs122 *GlobalState = globalState
						_ = _lhs122
						_lhs121.F18 = _rhs134
						_690_v15 = _rhs135
						r2 = _rhs136
						_lhs122.F18 = _rhs137
						var _692_v17 _dafny.Array
						_ = _692_v17
						var _len0_19 _dafny.Int = _dafny.IntOfInt64(27)
						_ = _len0_19
						var _nw125 _dafny.Array
						_ = _nw125
						if _len0_19.Cmp(_dafny.Zero) == 0 {
							_nw125 = _dafny.NewArray(_len0_19)
						} else {
							var _init19 func(_dafny.Int) _dafny.Int = (func(_693_v0 *C2) func(_dafny.Int) _dafny.Int {
								return func(_694_i3 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeDivisionInt(_694_i3, (_dafny.Zero).Minus((_693_v0).F27()))
								}
							})(_671_v0)
							_ = _init19
							var _element0_19 = _init19(_dafny.Zero)
							_ = _element0_19
							_nw125 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
							(_nw125).ArraySet1(_element0_19, 0)
							var _nativeLen0_19 = (_len0_19).Int()
							_ = _nativeLen0_19
							for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
								(_nw125).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
							}
						}
						_692_v17 = _nw125
						var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
						_ = _nw126
						_692_v17 = _nw126
						var _695_v18 D10
						_ = _695_v18
						_695_v18 = Companion_D10_.Create_DC19_((_676_v3).F24())
						var _696_v19 D10
						_ = _696_v19
						_696_v19 = Companion_D10_.Create_DC21_(_676_v3.F25(), (_this).F26(), (_671_v0).F28())
						_695_v18 = (func() D10 {
							if !((_671_v0).F28()) {
								return _695_v18
							}
							return (func() D10 {
								if (_696_v19).Dtor_cf32() {
									return func(_pat_let10_0 D10) D10 {
										return func(_697_dt__update__tmp_h0 D10) D10 {
											return func(_pat_let11_0 _dafny.CodePoint) D10 {
												return func(_698_dt__update_hcf27_h0 _dafny.CodePoint) D10 {
													return Companion_D10_.Create_DC19_(_698_dt__update_hcf27_h0)
												}(_pat_let11_0)
											}(_dafny.CodePoint('g'))
										}(_pat_let10_0)
									}(_695_v18)
								}
								return Companion_Default___.Fm20(_676_v3.F25(), true, globalState)
							})()
						})()
						var _699_v20 _dafny.Int
						_ = _699_v20
						var _700_v21 _dafny.Int
						_ = _700_v21
						var _out8 _dafny.Int
						_ = _out8
						var _out9 _dafny.Int
						_ = _out9
						_out8, _out9 = (_671_v0).M3((_this).F26(), (_this).F26(), globalState)
						_699_v20 = _out8
						_700_v21 = _out9
						var _701_v22 *C0
						_ = _701_v22
						var _nw127 *C0 = New_C0_()
						_ = _nw127
						_nw127.Ctor__()
						_701_v22 = _nw127
						_701_v22 = _701_v22
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _716_v23 _dafny.Sequence
		_ = _716_v23
		_716_v23 = _dafny.SeqOf(p1, (_this).F26(), p1, p1, p1)
		var _717_v24 _dafny.Map
		_ = _717_v24
		_717_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), p1)
		var _718_v25 _dafny.Sequence
		_ = _718_v25
		_718_v25 = _dafny.UnicodeSeqOfUtf8Bytes("pdxyd")
		var _719_v26 _dafny.Array
		_ = _719_v26
		var _nwElement0_27 bool = !(p1)
		_ = _nwElement0_27
		var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(20))
		_ = _nw128
		(_nw128).ArraySet1(_nwElement0_27, 0)
		(_nw128).ArraySet1(!((func() bool {
			if p1 {
				return (_671_v0).F28()
			}
			return p1
		})()), 1)
		(_nw128).ArraySet1((_this).F26(), 2)
		(_nw128).ArraySet1(((_dafny.Zero).Minus(_this.F25())).Cmp(_this.F25()) >= 0, 3)
		(_nw128).ArraySet1(!(!((_this.F25()).Cmp(_this.F25()) < 0)), 4)
		(_nw128).ArraySet1((_this).F26(), 5)
		(_nw128).ArraySet1((_this).F26(), 6)
		(_nw128).ArraySet1((_this).F26(), 7)
		(_nw128).ArraySet1((_this).F26(), 8)
		(_nw128).ArraySet1(p1, 9)
		(_nw128).ArraySet1(p1, 10)
		(_nw128).ArraySet1(((_671_v0).F28()) == (Companion_Default___.Fm0((_671_v0).F27(), p1, (_671_v0).F28(), _this.F25(), globalState)), 11)
		(_nw128).ArraySet1(((_this).F26()) || (p1), 12)
		(_nw128).ArraySet1((_716_v23).Select((Companion_Default___.SafeIndex((_671_v0).F27(), _dafny.IntOfUint32((_716_v23).Cardinality()))).Uint32()).(bool), 13)
		(_nw128).ArraySet1(!((_this).F26()), 14)
		(_nw128).ArraySet1((_671_v0).F28(), 15)
		(_nw128).ArraySet1((func() bool {
			if (_717_v24).Contains(_this.F25()) {
				return (_717_v24).Get(_this.F25()).(bool)
			}
			return (_this).F26()
		})(), 16)
		(_nw128).ArraySet1(p1, 17)
		(_nw128).ArraySet1((_this.F25()).Cmp(_dafny.IntOfUint32((_718_v25).Cardinality())) >= 0, 18)
		(_nw128).ArraySet1((_671_v0).F28(), 19)
		_719_v26 = _nw128
		for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_719_v26), 0))); ; {
			_guard_loop_2, _ok36 := _iter36()
			if !_ok36 {
				break
			}
			var _720_i4 _dafny.Int
			_720_i4 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_720_i4).Sign() != -1) && ((_720_i4).Cmp(_dafny.ArrayLen((_719_v26), 0)) < 0)) {
				(_719_v26).ArraySet1(((func() _dafny.MultiSet {
					if (_671_v0).F28() {
						return _dafny.MultiSetOf((_671_v0).F28())
					}
					return _dafny.MultiSetOf((_671_v0).F28(), !(p1), (_671_v0).F28())
				})()).IsProperSubsetOf((_dafny.MultiSetOf((_671_v0).F28())).Intersection(_dafny.MultiSetFromSeq(_716_v23))), (_720_i4).Int())
			}
		}
		for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_719_v26), 0))); ; {
			_guard_loop_3, _ok37 := _iter37()
			if !_ok37 {
				break
			}
			var _721_i5 _dafny.Int
			_721_i5 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_721_i5).Sign() != -1) && ((_721_i5).Cmp(_dafny.ArrayLen((_719_v26), 0)) < 0)) {
				(_719_v26).ArraySet1((_this).F26(), (_721_i5).Int())
			}
		}
		var _722_v27 _dafny.Map
		_ = _722_v27
		_722_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F25()).Times((_671_v0).F27()), _671_v0)
		_671_v0 = (func() *C2 {
			if (_722_v27).Contains((_671_v0).F27()) {
				return (_722_v27).Get((_671_v0).F27()).(*C2)
			}
			return _671_v0
		})()
		if p1 {
			var _723_v28 _dafny.CodePoint
			_ = _723_v28
			_723_v28 = _dafny.CodePoint('d')
			_723_v28 = (_this).F24()
			var _724_v29 D7
			_ = _724_v29
			_724_v29 = Companion_D7_.Create_DC14_((_this).F26(), (_dafny.Zero).Minus(_this.F25()), (_671_v0).F27())
			var _725_v30 _dafny.Map
			_ = _725_v30
			_725_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_716_v23), (_724_v29).Dtor_cf22())
			var _726_v31 _dafny.Map
			_ = _726_v31
			_726_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_717_v24).Contains((_671_v0).F27()) {
					return (_717_v24).Get((_671_v0).F27()).(bool)
				}
				return true
			})(), !((_671_v0).F28()))
			var _727_v32 *C4
			_ = _727_v32
			var _nw129 *C4 = New_C4_()
			_ = _nw129
			_nw129.Ctor__(_725_v30, (func() bool {
				if (_726_v31).Contains((_this).F26()) {
					return (_726_v31).Get((_this).F26()).(bool)
				}
				return p1
			})(), _dafny.CodePoint('f'), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(656), _this.F25()))
			_727_v32 = _nw129
			_727_v32 = _727_v32
			if !(false) {
				(globalState).F11 = _this.F25()
				var _728_v33 D7
				_ = _728_v33
				_728_v33 = Companion_D7_.Create_DC13_(_717_v24)
				var _729_v34 _dafny.MultiSet
				_ = _729_v34
				_729_v34 = _dafny.MultiSetOf((func() D7 {
					if (_this).F26() {
						return _728_v33
					}
					return _728_v33
				})())
				var _730_v35 _dafny.Array
				_ = _730_v35
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_20
				var _nw130 _dafny.Array
				_ = _nw130
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw130 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Int = func(_731_i6 _dafny.Int) _dafny.Int {
						return (_731_i6).Minus(_dafny.IntOfInt64(-813))
					}
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw130 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw130).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw130).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_730_v35 = _nw130
				var _732_v36 D6
				_ = _732_v36
				_732_v36 = Companion_D6_.Create_DC12_(p1, (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("pshrqd"), (Companion_Default___.SafeIndex((_671_v0).F27(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pshrqd")).Cardinality()))).Uint32(), _723_v28)).Cardinality())).Plus((_671_v0).F27()))
				var _rhs138 _dafny.MultiSet = _dafny.MultiSetFromSeq(Companion_Default___.Fm21(Companion_Default___.SafeModuloInt((_671_v0).F27(), (_671_v0).F27()), Companion_Default___.SafeModuloInt((_671_v0).F27(), (_671_v0).F27()), globalState))
				_ = _rhs138
				var _rhs139 _dafny.Array = _730_v35
				_ = _rhs139
				var _rhs140 D6 = _732_v36
				_ = _rhs140
				_729_v34 = _rhs138
				_730_v35 = _rhs139
				_732_v36 = _rhs140
				(globalState).F3 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_718_v25, _718_v25), _718_v25)
				var _733_v37 D10
				_ = _733_v37
				_733_v37 = Companion_D10_.Create_DC21_(((_726_v31).Update(p1, (_671_v0).F28())).Cardinality(), (_671_v0).F28(), (_671_v0).F28())
				var _734_v38 _dafny.Set
				_ = _734_v38
				_734_v38 = _dafny.SetOf(_733_v37)
				_734_v38 = _734_v38
				(globalState).F18 = _dafny.Companion_Sequence_.IsPrefixOf(_716_v23, Companion_Default___.Fm7(globalState))
			} else {
				var _735_v39 _dafny.Array
				_ = _735_v39
				var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw131
				_735_v39 = _nw131
				var _736_v40 _dafny.MultiSet
				_ = _736_v40
				_736_v40 = _dafny.MultiSetOf((_671_v0).F28())
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_735_v39), 0))
				_ = _index114
				(_735_v39).ArraySet1(Companion_Default___.SafeModuloInt((_736_v40).Cardinality(), _dafny.IntOfInt64(515)), (_index114).Int())
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_735_v39), 0))
				_ = _index115
				(_735_v39).ArraySet1(_dafny.IntOfInt64(460), (_index115).Int())
				var _737_v41 T0
				_ = _737_v41
				var _nw132 *C1 = New_C1_()
				_ = _nw132
				_nw132.Ctor__((_this).F24(), (_dafny.IntOfUint32((_716_v23).Cardinality())).Times((_671_v0).F27()))
				_737_v41 = _nw132
				var _738_v42 _dafny.Map
				_ = _738_v42
				_738_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_671_v0).F27(), (_this).F24())
				var _739_v43 _dafny.Set
				_ = _739_v43
				_739_v43 = _dafny.SetOf(_dafny.IntOfUint32((_718_v25).Cardinality()), (_671_v0).F27())
				var _nw133 *C4 = New_C4_()
				_ = _nw133
				_nw133.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf((_671_v0).F28(), (_671_v0).F28())), _737_v41.F25()), (_671_v0).F28(), (func() _dafny.CodePoint {
					if (_738_v42).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kanvy")).Cardinality())) {
						return (_738_v42).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kanvy")).Cardinality())).(_dafny.CodePoint)
					}
					return _723_v28
				})(), ((_739_v43).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(408))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_740_v28 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_741_i7 _dafny.Int) _dafny.CodePoint {
						return _740_v28
					}
				})(_723_v28)))).Cardinality())))
				_737_v41 = _nw133
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_735_v39), 0))
				_ = _index116
				(_735_v39).ArraySet1((_671_v0).F27(), (_index116).Int())
				var _742_v44 _dafny.Array
				_ = _742_v44
				var _len0_21 _dafny.Int = _dafny.One
				_ = _len0_21
				var _nw134 _dafny.Array
				_ = _nw134
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw134 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) D3 = (func(_743_v41 T0, _744_v39 _dafny.Array, _745_v0 *C2) func(_dafny.Int) D3 {
						return func(_746_i8 _dafny.Int) D3 {
							return Companion_D3_.Create_DC6_(_743_v41.F25(), (_744_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_744_v39), 0))).Int()).(_dafny.Int), (_745_v0).F28())
						}
					})(_737_v41, _735_v39, _671_v0)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw134 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw134).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw134).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_742_v44 = _nw134
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_742_v44), 0))
				_ = _index117
				(_742_v44).ArraySet1(Companion_D3_.Create_DC6_(Companion_Default___.Fm17((_671_v0).F28(), (_727_v32).F30(), p1, globalState), _this.F25(), Companion_Default___.Fm0(_this.F25(), !((_727_v32).F30()), false, _this.F25(), globalState)), (_index117).Int())
				var _747_v45 _dafny.Map
				_ = _747_v45
				_747_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_671_v0).F28(), (_735_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_735_v39), 0))).Int()).(_dafny.Int))
				var _748_v46 D3
				_ = _748_v46
				_748_v46 = Companion_D3_.Create_DC6_((_671_v0).F27(), _737_v41.F25(), p1)
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_742_v44), 0))
				_ = _index118
				var _rhs141 _dafny.Map = _747_v45
				_ = _rhs141
				var _rhs142 D3 = _748_v46
				_ = _rhs142
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				var _lhs124 _dafny.Array = _742_v44
				_ = _lhs124
				var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_742_v44), 0))
				_ = _lhs125
				_lhs123.F13 = _rhs141
				(_lhs124).ArraySet1(_rhs142, (_lhs125).Int())
				var _749_v47 _dafny.Array
				_ = _749_v47
				var _nw135 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
				_ = _nw135
				_749_v47 = _nw135
				var _750_v48 D11
				_ = _750_v48
				_750_v48 = Companion_D11_.Create_DC22_(_749_v47)
				var _751_v49 _dafny.Sequence
				_ = _751_v49
				_751_v49 = _dafny.SeqOf(_750_v48, _750_v48)
				var _752_v50 _dafny.Sequence
				_ = _752_v50
				_752_v50 = _dafny.SeqOf((_dafny.Zero).Minus((_735_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_735_v39), 0))).Int()).(_dafny.Int)))
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_735_v39), 0))
				_ = _index119
				var _rhs143 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_751_v49, _dafny.Companion_Sequence_.Update(_751_v49, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_752_v50).Cardinality()), _dafny.IntOfUint32((_751_v49).Cardinality()))).Uint32(), _750_v48))
				_ = _rhs143
				var _rhs144 _dafny.Int = ((_671_v0).F27()).Times(_737_v41.F25())
				_ = _rhs144
				var _rhs145 _dafny.Array = _719_v26
				_ = _rhs145
				var _lhs126 _dafny.Array = _735_v39
				_ = _lhs126
				var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_735_v39), 0))
				_ = _lhs127
				_751_v49 = _rhs143
				(_lhs126).ArraySet1(_rhs144, (_lhs127).Int())
				_719_v26 = _rhs145
			}
			var _753_v51 *C1
			_ = _753_v51
			var _nw136 *C1 = New_C1_()
			_ = _nw136
			_nw136.Ctor__(_dafny.CodePoint('v'), _dafny.IntOfUint32((_718_v25).Cardinality()))
			_753_v51 = _nw136
			var _754_v53 _dafny.MultiSet
			_ = _754_v53
			_754_v53 = _dafny.MultiSetOf((_671_v0).F28(), (_671_v0).F28(), !((_671_v0).F28()), !(true), (func() bool {
				if (_726_v31).Contains((_671_v0).F28()) {
					return (_726_v31).Get((_671_v0).F28()).(bool)
				}
				return !((_this).F26())
			})())
			var _rhs146 _dafny.Int = (_dafny.Zero).Minus((_671_v0).F27())
			_ = _rhs146
			var _rhs147 bool = !(p1)
			_ = _rhs147
			var _rhs148 _dafny.Array = _719_v26
			_ = _rhs148
			var _rhs149 bool = ((_671_v0).Fm4(func() _dafny.Set {
				var _coll34 = _dafny.NewBuilder()
				_ = _coll34
				for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(93), _dafny.IntOfInt64(576))); ; {
					_compr_34, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _755_v52 _dafny.Int
					_755_v52 = interface{}(_compr_34).(_dafny.Int)
					if ((_dafny.IntOfInt64(93)).Cmp(_755_v52) <= 0) && ((_755_v52).Cmp(_dafny.IntOfInt64(576)) < 0) {
						_coll34.Add(Companion_Default___.SafeModuloInt(_755_v52, (_671_v0).F27()))
					}
				}
				return _coll34.ToSet()
			}(), _dafny.IntOfUint32((_718_v25).Cardinality()), (_671_v0).F28(), globalState)).IsDisjointFrom(_754_v53)
			_ = _rhs149
			var _lhs128 *GlobalState = globalState
			_ = _lhs128
			var _lhs129 *GlobalState = globalState
			_ = _lhs129
			_lhs128.F22 = _rhs146
			r0 = _rhs147
			_719_v26 = _rhs148
			_lhs129.F18 = _rhs149
		} else {
			r1 = _716_v23
			(globalState).F2 = (((_671_v0).F27()).Minus(Companion_Default___.Fm17(!(true), (_this).F26(), p1, globalState))).Plus((_671_v0).F27())
			var _756_v54 _dafny.Map
			_ = _756_v54
			_756_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_671_v0).F28(), _dafny.IntOfUint32((_718_v25).Cardinality()))
			var _757_v55 _dafny.Map
			_ = _757_v55
			_757_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm0(_dafny.IntOfInt64(-691), (_716_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-920), _dafny.IntOfUint32((_716_v23).Cardinality()))).Uint32()).(bool), (_this).F26(), (func() _dafny.Int {
				if (_756_v54).Contains(true) {
					return (_756_v54).Get(true).(_dafny.Int)
				}
				return (_dafny.Zero).Minus((_671_v0).F27())
			})(), globalState)), (_671_v0).F27())
			(globalState).F13 = _756_v54
			if p1 {
				var _758_v56 _dafny.Map
				_ = _758_v56
				_758_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_671_v0).F27())
				var _759_v57 _dafny.MultiSet
				_ = _759_v57
				_759_v57 = _dafny.MultiSetOf(_716_v23, _716_v23)
				var _760_v58 _dafny.MultiSet
				_ = _760_v58
				_760_v58 = _dafny.MultiSetOf((_671_v0).F28(), (_671_v0).F28(), p1, Companion_Default___.Fm0((_671_v0).F27(), (_671_v0).F28(), (_671_v0).F28(), (_671_v0).F27(), globalState))
				var _761_v59 _dafny.Map
				_ = _761_v59
				_761_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_671_v0).F27(), (_671_v0).F27())
				var _762_v60 _dafny.Set
				_ = _762_v60
				_762_v60 = _dafny.SetOf((_671_v0).F27())
				var _763_v61 _dafny.Sequence
				_ = _763_v61
				_763_v61 = _dafny.SeqOf((func() _dafny.Int {
					if (_761_v59).Contains((_762_v60).Cardinality()) {
						return (_761_v59).Get((_762_v60).Cardinality()).(_dafny.Int)
					}
					return _dafny.IntOfInt64(125)
				})(), _dafny.IntOfUint32((_dafny.SeqOf(true, (_this).F26(), (_671_v0).F28(), p1, (_671_v0).F28())).Cardinality()), (_671_v0).F27(), (_671_v0).F27(), (_761_v59).Cardinality())
				var _764_v62 _dafny.Array
				_ = _764_v62
				var _nwElement0_28 _dafny.Int = Companion_Default___.Fm5((_671_v0).F27(), (_671_v0).F27(), globalState)
				_ = _nwElement0_28
				var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(26))
				_ = _nw137
				(_nw137).ArraySet1(_nwElement0_28, 0)
				(_nw137).ArraySet1((_671_v0).F27(), 1)
				(_nw137).ArraySet1(_this.F25(), 2)
				(_nw137).ArraySet1((func() _dafny.Int {
					if (_759_v57).Contains(_dafny.SeqOf((_this).F26())) {
						return (_759_v57).Multiplicity(_dafny.SeqOf((_this).F26()))
					}
					return (_671_v0).F27()
				})(), 3)
				(_nw137).ArraySet1((_671_v0).F27(), 4)
				(_nw137).ArraySet1((_671_v0).F27(), 5)
				(_nw137).ArraySet1((_671_v0).F27(), 6)
				(_nw137).ArraySet1((_671_v0).F27(), 7)
				(_nw137).ArraySet1(_this.F25(), 8)
				(_nw137).ArraySet1((_671_v0).F27(), 9)
				(_nw137).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 10)
				(_nw137).ArraySet1((_671_v0).F27(), 11)
				(_nw137).ArraySet1((_671_v0).F27(), 12)
				(_nw137).ArraySet1((_671_v0).F27(), 13)
				(_nw137).ArraySet1((_671_v0).F27(), 14)
				(_nw137).ArraySet1((_760_v58).Cardinality(), 15)
				(_nw137).ArraySet1((_671_v0).F27(), 16)
				(_nw137).ArraySet1((_671_v0).F27(), 17)
				(_nw137).ArraySet1((_671_v0).F27(), 18)
				(_nw137).ArraySet1(_dafny.IntOfUint32((_716_v23).Cardinality()), 19)
				(_nw137).ArraySet1(_this.F25(), 20)
				(_nw137).ArraySet1(_this.F25(), 21)
				(_nw137).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_763_v61).Cardinality())), 22)
				(_nw137).ArraySet1(_this.F25(), 23)
				(_nw137).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uxb")).Cardinality()), 24)
				(_nw137).ArraySet1((_671_v0).F27(), 25)
				_764_v62 = _nw137
				var _765_v63 _dafny.Array
				_ = _765_v63
				var _nwElement0_29 _dafny.Array = _764_v62
				_ = _nwElement0_29
				var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(13))
				_ = _nw138
				(_nw138).ArraySet1(_nwElement0_29, 0)
				(_nw138).ArraySet1(_764_v62, 1)
				(_nw138).ArraySet1(_764_v62, 2)
				(_nw138).ArraySet1(_764_v62, 3)
				(_nw138).ArraySet1(_764_v62, 4)
				(_nw138).ArraySet1(_764_v62, 5)
				(_nw138).ArraySet1(_764_v62, 6)
				(_nw138).ArraySet1(_764_v62, 7)
				(_nw138).ArraySet1(_764_v62, 8)
				(_nw138).ArraySet1(_764_v62, 9)
				(_nw138).ArraySet1(_764_v62, 10)
				(_nw138).ArraySet1(_764_v62, 11)
				(_nw138).ArraySet1(_764_v62, 12)
				_765_v63 = _nw138
				var _766_v64 _dafny.Sequence
				_ = _766_v64
				_766_v64 = _dafny.UnicodeSeqOfUtf8Bytes("bedhejjw")
				var _767_v65 _dafny.Map
				_ = _767_v65
				_767_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_766_v64), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_671_v0).F27())).Merge(_758_v56))
				var _rhs150 *C2 = _671_v0
				_ = _rhs150
				var _rhs151 _dafny.Map = (func() _dafny.Map {
					if (_767_v65).Contains(_dafny.Companion_Sequence_.Concatenate(_718_v25, _718_v25)) {
						return (_767_v65).Get(_dafny.Companion_Sequence_.Concatenate(_718_v25, _718_v25)).(_dafny.Map)
					}
					return _758_v56
				})()
				_ = _rhs151
				var _rhs152 _dafny.Int = (_671_v0).F27()
				_ = _rhs152
				var _rhs153 _dafny.Array = _765_v63
				_ = _rhs153
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				_671_v0 = _rhs150
				_758_v56 = _rhs151
				_lhs130.F10 = _rhs152
				_765_v63 = _rhs153
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.ArrayLen((_764_v62), 0))
				_ = _index120
				(_764_v62).ArraySet1((_671_v0).F27(), (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.ArrayLen((_764_v62), 0))
				_ = _index121
				(_764_v62).ArraySet1((_this.F25()).Times(_this.F25()), (_index121).Int())
				var _768_v66 _dafny.Map
				_ = _768_v66
				var _769_v67 bool
				_ = _769_v67
				var _out10 _dafny.Map
				_ = _out10
				var _out11 bool
				_ = _out11
				_out10, _out11 = (_this).M2(globalState)
				_768_v66 = _out10
				_769_v67 = _out11
				var _770_v68 *C3
				_ = _770_v68
				var _nw139 *C3 = New_C3_()
				_ = _nw139
				_nw139.Ctor__(_769_v67, (_this).F24(), _this.F25())
				_770_v68 = _nw139
				_671_v0 = _671_v0
			} else {
				var _771_v69 _dafny.Map
				_ = _771_v69
				_771_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), Companion_Default___.Fm10((_this).F26(), (_671_v0).F28(), false, (_671_v0).F28(), globalState))
				var _772_v70 _dafny.Sequence
				_ = _772_v70
				_772_v70 = _dafny.SeqOf(_771_v69, _771_v69, _771_v69)
				_772_v70 = _772_v70
				var _773_v71 *C1
				_ = _773_v71
				var _nw140 *C1 = New_C1_()
				_ = _nw140
				_nw140.Ctor__((_this).F24(), (_671_v0).F27())
				_773_v71 = _nw140
				var _774_v72 _dafny.Map
				_ = _774_v72
				_774_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _773_v71)
				_773_v71 = (func() *C1 {
					if (_774_v72).Contains((_this).F26()) {
						return (_774_v72).Get((_this).F26()).(*C1)
					}
					return _773_v71
				})()
				var _775_v73 _dafny.Map
				_ = _775_v73
				_775_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _716_v23)
				_775_v73 = (_775_v73).Update(false, _716_v23)
				var _776_v74 _dafny.Set
				_ = _776_v74
				_776_v74 = _dafny.SetOf((_this).F26())
				r2 = (_776_v74).IsSubsetOf(_776_v74)
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_719_v26), 0))
				_ = _index122
				(_719_v26).ArraySet1((_this).F26(), (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_719_v26), 0))
				_ = _index123
				(_719_v26).ArraySet1((_671_v0).F28(), (_index123).Int())
			}
			var _777_v75 _dafny.Array
			_ = _777_v75
			var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw141
			_777_v75 = _nw141
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_777_v75), 0))
			_ = _index124
			(_777_v75).ArraySet1(_719_v26, (_index124).Int())
			var _778_v76 _dafny.Array
			_ = _778_v76
			var _nw142 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw142
			_778_v76 = _nw142
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_778_v76), 0))
			_ = _index125
			(_778_v76).ArraySet1(_dafny.IntOfInt64(-380), (_index125).Int())
			var _779_v77 _dafny.Sequence
			_ = _779_v77
			_779_v77 = _dafny.SeqOf(_719_v26, _719_v26)
			var _780_v78 D3
			_ = _780_v78
			_780_v78 = Companion_D3_.Create_DC5_(_716_v23, _this.F25(), (_671_v0).F27())
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_777_v75), 0))
			_ = _index126
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_778_v76), 0))
			_ = _index127
			var _rhs154 _dafny.Array = (_779_v77).Select((Companion_Default___.SafeIndex((_671_v0).F27(), _dafny.IntOfUint32((_779_v77).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs154
			var _rhs155 _dafny.Int = _this.F25()
			_ = _rhs155
			var _rhs156 _dafny.Int = Companion_Default___.SafeDivisionInt((_this.F25()).Times((_671_v0).F27()), (_671_v0).F27())
			_ = _rhs156
			var _rhs157 bool = (_716_v23).Select((Companion_Default___.SafeIndex((_780_v78).Dtor_cf8(), _dafny.IntOfUint32((_716_v23).Cardinality()))).Uint32()).(bool)
			_ = _rhs157
			var _lhs131 _dafny.Array = _777_v75
			_ = _lhs131
			var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_777_v75), 0))
			_ = _lhs132
			var _lhs133 _dafny.Array = _778_v76
			_ = _lhs133
			var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_778_v76), 0))
			_ = _lhs134
			var _lhs135 *GlobalState = globalState
			_ = _lhs135
			(_lhs131).ArraySet1(_rhs154, (_lhs132).Int())
			(_lhs133).ArraySet1(_rhs155, (_lhs134).Int())
			_lhs135.F15 = _rhs156
			r2 = _rhs157
		}
		r0 = (_this).F26()
		r1 = _dafny.Companion_Sequence_.Update(_716_v23, (Companion_Default___.SafeIndex((_this.F25()).Minus(_dafny.IntOfInt64(758)), _dafny.IntOfUint32((_716_v23).Cardinality()))).Uint32(), (_671_v0).F28())
		r2 = (_671_v0).F28()
		return r0, r1, r2
	}
}
func (_this *C5) M2(globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _781_v0 D7
		_ = _781_v0
		_781_v0 = Companion_D7_.Create_DC14_((_this).F26(), _dafny.IntOfUint32(((_this).Fm3((_this).F24(), globalState)).Cardinality()), _this.F25())
		if (func() bool {
			if (_this).F26() {
				return (_781_v0).Dtor_cf20()
			}
			return !(((_dafny.Zero).Minus(_this.F25())).Cmp(_dafny.IntOfInt64(961)) > 0)
		})() {
			var _782_v2 _dafny.Map
			_ = _782_v2
			_782_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), (_this).F26())
			var _783_v3 _dafny.Set
			_ = _783_v3
			_783_v3 = _dafny.SetOf(_this.F25(), (func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter39 := _dafny.Iterate((_782_v2).Keys().Elements()); ; {
					_compr_35, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _784_v1 _dafny.CodePoint
					_784_v1 = interface{}(_compr_35).(_dafny.CodePoint)
					if (_782_v2).Contains(_784_v1) {
						_coll35.Add(_784_v1, (_dafny.Zero).Minus(_this.F25()))
					}
				}
				return _coll35.ToMap()
			}()).Cardinality())
			var _785_v4 D11
			_ = _785_v4
			_785_v4 = Companion_D11_.Create_DC23_((_783_v3).Cardinality(), (_this).F26())
			var _786_v5 *C3
			_ = _786_v5
			var _nw143 *C3 = New_C3_()
			_ = _nw143
			_nw143.Ctor__(!((_785_v4).Dtor_cf35()), (_this).F24(), (func() _dafny.Int {
				if (_this).F26() {
					return _this.F25()
				}
				return _this.F25()
			})())
			_786_v5 = _nw143
			var _787_v6 _dafny.CodePoint
			_ = _787_v6
			_787_v6 = _dafny.CodePoint('f')
			var _rhs158 _dafny.Int = (Companion_Default___.Fm22(_dafny.IntOfInt64(949), !(!((_this).F26())), _786_v5.F31, globalState)).Dtor_cf34()
			_ = _rhs158
			var _rhs159 _dafny.CodePoint = _787_v6
			_ = _rhs159
			var _rhs160 bool = _786_v5.F31
			_ = _rhs160
			var _rhs161 bool = (_this).F26()
			_ = _rhs161
			var _lhs136 *GlobalState = globalState
			_ = _lhs136
			var _lhs137 *GlobalState = globalState
			_ = _lhs137
			var _lhs138 *GlobalState = globalState
			_ = _lhs138
			_lhs136.F0 = _rhs158
			_787_v6 = _rhs159
			_lhs137.F18 = _rhs160
			_lhs138.F18 = _rhs161
			var _788_v7 _dafny.MultiSet
			_ = _788_v7
			_788_v7 = _dafny.MultiSetOf(_this.F25(), _this.F25())
			var _789_v8 _dafny.Sequence
			_ = _789_v8
			_789_v8 = _dafny.SeqOf(_this.F25())
			(_this).F25_set_(Companion_Default___.Fm17((_this).F26(), (_this).F26(), (_dafny.MultiSetFromSeq(_789_v8)).IsProperSubsetOf(_788_v7), globalState))
			var _790_v9 _dafny.Map
			_ = _790_v9
			_790_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_789_v8, _786_v5.F31)
			if (_790_v9).Equals(_790_v9) {
				var _791_v10 _dafny.Sequence
				_ = _791_v10
				_791_v10 = _dafny.SeqOf(_786_v5.F31)
				_791_v10 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_791_v10, _791_v10), _791_v10)
				(globalState).F18 = _786_v5.F31
				var _792_v11 _dafny.Map
				_ = _792_v11
				_792_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_791_v10, (Companion_Default___.SafeIndex(_this.F25(), _dafny.IntOfUint32((_791_v10).Cardinality()))).Uint32(), _786_v5.F31), _791_v10))
				_792_v11 = (_792_v11).Update((_this).F24(), _791_v10)
				var _793_v12 _dafny.Array
				_ = _793_v12
				var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
				_ = _nw144
				_793_v12 = _nw144
				_793_v12 = _793_v12
				(_786_v5).F31 = _786_v5.F31
			} else {
				var _794_v13 _dafny.Map
				_ = _794_v13
				_794_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm17(_786_v5.F31, _786_v5.F31, _786_v5.F31, globalState), false)
				_794_v13 = _794_v13
				_787_v6 = (_this).F24()
				(globalState).F10 = (_this.F25()).Minus(_this.F25())
				var _795_v14 _dafny.Array
				_ = _795_v14
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_22
				var _nw145 _dafny.Array
				_ = _nw145
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw145 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) bool = (func(_796_v5 *C3) func(_dafny.Int) bool {
						return func(_797_i0 _dafny.Int) bool {
							return _796_v5.F31
						}
					})(_786_v5)
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
				_795_v14 = _nw145
				var _798_v15 _dafny.Sequence
				_ = _798_v15
				_798_v15 = _dafny.SeqOf(_788_v7)
				var _799_v16 _dafny.Map
				_ = _799_v16
				_799_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_798_v15, _798_v15), _795_v14)
				_795_v14 = (func() _dafny.Array {
					if (_799_v16).Contains((_this).F26()) {
						return (_799_v16).Get((_this).F26()).(_dafny.Array)
					}
					return _795_v14
				})()
				var _800_v17 _dafny.Map
				_ = _800_v17
				_800_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_786_v5.F31), (_783_v3).Cardinality())
				var _801_v18 _dafny.Map
				_ = _801_v18
				_801_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10((_this).F26(), _786_v5.F31, true, (_this).F26(), globalState), (_this).F24())
				var _802_v19 T0
				_ = _802_v19
				var _nw146 *C4 = New_C4_()
				_ = _nw146
				_nw146.Ctor__(_800_v17, _786_v5.F31, (func() _dafny.CodePoint {
					if (_801_v18).Contains((_this).F24()) {
						return (_801_v18).Get((_this).F24()).(_dafny.CodePoint)
					}
					return _787_v6
				})(), _this.F25())
				_802_v19 = _nw146
				var _803_v20 _dafny.Map
				_ = _803_v20
				_803_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v19, (_this).F24())
				_803_v20 = (_803_v20).Update(_802_v19, (_802_v19).F24())
			}
			var _804_v21 _dafny.Array
			_ = _804_v21
			var _nwElement0_30 bool = _786_v5.F31
			_ = _nwElement0_30
			var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.One)
			_ = _nw147
			(_nw147).ArraySet1(_nwElement0_30, 0)
			_804_v21 = _nw147
			var _805_v22 _dafny.Array
			_ = _805_v22
			var _nw148 _dafny.Array = _dafny.NewArrayWithValue(Companion_D11_.Default(), _dafny.IntOfInt64(21))
			_ = _nw148
			_805_v22 = _nw148
			var _806_v23 _dafny.Array
			_ = _806_v23
			var _nw149 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
			_ = _nw149
			_806_v23 = _nw149
			var _807_v24 D11
			_ = _807_v24
			_807_v24 = Companion_D11_.Create_DC22_(_806_v23)
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_805_v22), 0))
			_ = _index128
			(_805_v22).ArraySet1(_807_v24, (_index128).Int())
			var _808_v25 _dafny.Array
			_ = _808_v25
			var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
			_ = _nw150
			_808_v25 = _nw150
			var _809_v26 _dafny.Array
			_ = _809_v26
			var _nw151 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
			_ = _nw151
			_809_v26 = _nw151
			var _810_v27 _dafny.Map
			_ = _810_v27
			_810_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _809_v26)
			var _811_v28 _dafny.Sequence
			_ = _811_v28
			_811_v28 = _dafny.SeqOf(_804_v21)
			var _812_v29 _dafny.Map
			_ = _812_v29
			_812_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_810_v27).Merge(_810_v27), (_811_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.IntOfUint32((_811_v28).Cardinality()))).Uint32()).(_dafny.Array))
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_805_v22), 0))
			_ = _index129
			var _rhs162 _dafny.Array = (func() _dafny.Array {
				if (_812_v29).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _809_v26)) {
					return (_812_v29).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _809_v26)).(_dafny.Array)
				}
				return _804_v21
			})()
			_ = _rhs162
			var _rhs163 D11 = _807_v24
			_ = _rhs163
			var _rhs164 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_this.F25()), _this.F25())
			_ = _rhs164
			var _rhs165 _dafny.Array = _808_v25
			_ = _rhs165
			var _lhs139 _dafny.Array = _805_v22
			_ = _lhs139
			var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_805_v22), 0))
			_ = _lhs140
			var _lhs141 *GlobalState = globalState
			_ = _lhs141
			_804_v21 = _rhs162
			(_lhs139).ArraySet1(_rhs163, (_lhs140).Int())
			_lhs141.F16 = _rhs164
			_808_v25 = _rhs165
		} else {
			var _813_v30 _dafny.Array
			_ = _813_v30
			var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
			_ = _nw152
			_813_v30 = _nw152
			var _814_v31 _dafny.Sequence
			_ = _814_v31
			_814_v31 = _dafny.SeqOf(Companion_D10_.Create_DC21_(_this.F25(), (_this).F26(), (_this).F26()))
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_813_v30), 0))
			_ = _index130
			(_813_v30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_814_v31, _814_v31), (_index130).Int())
			var _815_v32 D10
			_ = _815_v32
			_815_v32 = Companion_D10_.Create_DC21_(_dafny.IntOfInt64(745), (_this).F26(), (_this).F26())
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_813_v30), 0))
			_ = _index131
			(_813_v30).ArraySet1(_dafny.SeqOf(_815_v32, _815_v32, _815_v32, _815_v32, _815_v32), (_index131).Int())
			var _816_v33 _dafny.Sequence
			_ = _816_v33
			_816_v33 = _dafny.UnicodeSeqOfUtf8Bytes("qxbymslm")
			var _817_v34 D11
			_ = _817_v34
			_817_v34 = Companion_D11_.Create_DC24_((_this).F26(), _dafny.IntOfInt64(802))
			var _818_v35 D4
			_ = _818_v35
			_818_v35 = Companion_D4_.Create_DC8_()
			var _819_v36 _dafny.Map
			_ = _819_v36
			_819_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_817_v34).Dtor_cf36(), _818_v35)
			var _820_v37 _dafny.Map
			_ = _820_v37
			_820_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_816_v33, _819_v36)
			_820_v37 = (_820_v37).Update(_816_v33, _819_v36)
			var _821_v38 _dafny.Map
			_ = _821_v38
			_821_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), (_this).F26())
			var _822_v39 _dafny.Map
			_ = _822_v39
			_822_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F26())
			var _823_v40 _dafny.Sequence
			_ = _823_v40
			_823_v40 = _dafny.SeqOf((_821_v38).Cardinality(), _this.F25(), _this.F25(), (_822_v39).Cardinality())
			var _824_v41 _dafny.Sequence
			_ = _824_v41
			_824_v41 = _dafny.SeqOf((_dafny.Zero).Minus(_this.F25()), _dafny.IntOfUint32((_823_v40).Cardinality()), _dafny.IntOfInt64(-217))
			var _825_v42 _dafny.Set
			_ = _825_v42
			_825_v42 = _dafny.SetOf((_824_v41).Select((Companion_Default___.SafeIndex(_this.F25(), _dafny.IntOfUint32((_824_v41).Cardinality()))).Uint32()).(_dafny.Int))
			var _826_v43 _dafny.Map
			_ = _826_v43
			_826_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_825_v42).Cardinality())
			(globalState).F22 = (func() _dafny.Int {
				if Companion_Default___.Fm0(_this.F25(), (_this).F26(), (_this).F26(), _this.F25(), globalState) {
					return (func() _dafny.Int {
						if (_826_v43).Contains((_this).F26()) {
							return (_826_v43).Get((_this).F26()).(_dafny.Int)
						}
						return _this.F25()
					})()
				}
				return _this.F25()
			})()
			var _827_v44 *C1
			_ = _827_v44
			var _nw153 *C1 = New_C1_()
			_ = _nw153
			_nw153.Ctor__((_this).F24(), Companion_Default___.SafeDivisionInt(_this.F25(), _this.F25()))
			_827_v44 = _nw153
			var _828_v45 _dafny.Array
			_ = _828_v45
			var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw154
			_828_v45 = _nw154
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_828_v45), 0))
			_ = _index132
			(_828_v45).ArraySet1(Companion_Default___.Fm5(_dafny.IntOfUint32((_816_v33).Cardinality()), _this.F25(), globalState), (_index132).Int())
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_828_v45), 0))
			_ = _index133
			(_828_v45).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_this.F25()).Plus(Companion_Default___.SafeDivisionInt(_this.F25(), _this.F25())))), (_index133).Int())
		}
		var _829_v46 D10
		_ = _829_v46
		_829_v46 = Companion_D10_.Create_DC21_(_this.F25(), !((_this).F26()), (_this).F26())
		var _830_v47 _dafny.Array
		_ = _830_v47
		var _len0_23 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_23
		var _nw155 _dafny.Array
		_ = _nw155
		if _len0_23.Cmp(_dafny.Zero) == 0 {
			_nw155 = _dafny.NewArray(_len0_23)
		} else {
			var _init23 func(_dafny.Int) _dafny.Map = func(_831_i1 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), !((_this).F26()))
			}
			_ = _init23
			var _element0_23 = _init23(_dafny.Zero)
			_ = _element0_23
			_nw155 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
			(_nw155).ArraySet1(_element0_23, 0)
			var _nativeLen0_23 = (_len0_23).Int()
			_ = _nativeLen0_23
			for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
				(_nw155).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
			}
		}
		_830_v47 = _nw155
		var _source13 D4 = Companion_D4_.Create_DC7_((func() _dafny.Array {
			if (_829_v46).Dtor_cf32() {
				return _830_v47
			}
			return _830_v47
		})())
		_ = _source13
		if _source13.Is_DC8() {
			var _832_v48 _dafny.Array
			_ = _832_v48
			var _nwElement0_31 _dafny.Int = Companion_Default___.SafeDivisionInt(_this.F25(), _this.F25())
			_ = _nwElement0_31
			var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.One)
			_ = _nw156
			(_nw156).ArraySet1(_nwElement0_31, 0)
			_832_v48 = _nw156
			var _833_v49 _dafny.Sequence
			_ = _833_v49
			_833_v49 = _dafny.UnicodeSeqOfUtf8Bytes("v")
			var _834_v50 _dafny.MultiSet
			_ = _834_v50
			_834_v50 = _dafny.MultiSetOf(_this.F25(), _dafny.IntOfUint32((_833_v49).Cardinality()))
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))
			_ = _index134
			(_832_v48).ArraySet1(((_834_v50).Cardinality()).Minus(_this.F25()), (_index134).Int())
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))
			_ = _index135
			(_832_v48).ArraySet1(_this.F25(), (_index135).Int())
			var _835_v51 D10
			_ = _835_v51
			_835_v51 = Companion_D10_.Create_DC19_((_this).F24())
			var _836_v52 _dafny.Array
			_ = _836_v52
			var _nwElement0_32 _dafny.CodePoint = (_this).F24()
			_ = _nwElement0_32
			var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
			_ = _nw157
			(_nw157).ArraySet1CodePoint(_nwElement0_32, 0)
			(_nw157).ArraySet1CodePoint((_this).F24(), 1)
			(_nw157).ArraySet1CodePoint((_this).F24(), 2)
			(_nw157).ArraySet1CodePoint((_this).F24(), 3)
			(_nw157).ArraySet1CodePoint((_this).F24(), 4)
			(_nw157).ArraySet1CodePoint(_dafny.CodePoint('q'), 5)
			(_nw157).ArraySet1CodePoint((_this).F24(), 6)
			(_nw157).ArraySet1CodePoint((_this).F24(), 7)
			(_nw157).ArraySet1CodePoint((_this).F24(), 8)
			(_nw157).ArraySet1CodePoint(Companion_Default___.Fm10((_this).F26(), (_this).F26(), (_this).F26(), (_this).F26(), globalState), 9)
			(_nw157).ArraySet1CodePoint((_this).F24(), 10)
			(_nw157).ArraySet1CodePoint((_this).F24(), 11)
			(_nw157).ArraySet1CodePoint((_this).F24(), 12)
			(_nw157).ArraySet1CodePoint((_this).F24(), 13)
			(_nw157).ArraySet1CodePoint(Companion_Default___.Fm10((_this).F26(), (_this).F26(), (_this).F26(), (_this).F26(), globalState), 14)
			(_nw157).ArraySet1CodePoint((_this).F24(), 15)
			(_nw157).ArraySet1CodePoint((_this).F24(), 16)
			(_nw157).ArraySet1CodePoint(_dafny.CodePoint('u'), 17)
			(_nw157).ArraySet1CodePoint((_this).F24(), 18)
			(_nw157).ArraySet1CodePoint((_this).F24(), 19)
			(_nw157).ArraySet1CodePoint((_this).F24(), 20)
			(_nw157).ArraySet1CodePoint((_835_v51).Dtor_cf27(), 21)
			(_nw157).ArraySet1CodePoint((_this).F24(), 22)
			(_nw157).ArraySet1CodePoint((_this).F24(), 23)
			(_nw157).ArraySet1CodePoint((_this).F24(), 24)
			_836_v52 = _nw157
			var _837_v53 _dafny.Sequence
			_ = _837_v53
			_837_v53 = _dafny.SeqOf(_833_v49, _833_v49, _833_v49, _dafny.UnicodeSeqOfUtf8Bytes("kucbeyarw"), (_this).Fm3((_this).F24(), globalState))
			var _838_v54 _dafny.MultiSet
			_ = _838_v54
			_838_v54 = _dafny.MultiSetOf(_837_v53, _837_v53, _837_v53, _837_v53, _837_v53)
			var _rhs166 _dafny.Array = _836_v52
			_ = _rhs166
			var _rhs167 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_838_v54).Contains(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(920))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}(func(_839_i2 _dafny.Int) _dafny.CodePoint {
					return (_this).F24()
				})))) {
					return (_838_v54).Multiplicity(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(920))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}(func(_840_i2 _dafny.Int) _dafny.CodePoint {
						return (_this).F24()
					}))))
				}
				return _this.F25()
			})())
			_ = _rhs167
			var _rhs168 bool = (_this).F26()
			_ = _rhs168
			var _lhs142 *GlobalState = globalState
			_ = _lhs142
			_836_v52 = _rhs166
			_lhs142.F16 = _rhs167
			r1 = _rhs168
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))
			_ = _index136
			(_832_v48).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(267), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(845))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}(func(_841_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), _833_v49)).Cardinality())), (_index136).Int())
			var _842_v55 D3
			_ = _842_v55
			_842_v55 = Companion_D3_.Create_DC6_((_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int), (_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int), true)
			if (_842_v55).Dtor_cf11() {
				var _843_v56 _dafny.Map
				_ = _843_v56
				_843_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.IntOfUint32((_833_v49).Cardinality()), (_this).F26(), (_this).F26(), _this.F25(), globalState), _832_v48)
				_832_v48 = (func() _dafny.Array {
					if (_843_v56).Contains((_this).F26()) {
						return (_843_v56).Get((_this).F26()).(_dafny.Array)
					}
					return _832_v48
				})()
				var _844_v57 _dafny.Map
				_ = _844_v57
				_844_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), !(Companion_Default___.Fm0(_this.F25(), (_this).F26(), (_this).F26(), _this.F25(), globalState)))
				_844_v57 = (_844_v57).Update(true, (_this).F26())
				var _845_v58 _dafny.Map
				_ = _845_v58
				_845_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _842_v55)
				var _846_v59 D9
				_ = _846_v59
				_846_v59 = Companion_D9_.Create_DC16_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _842_v55)).Merge((_845_v58).Update(!((_this).F26()), Companion_Default___.Fm23(globalState))))
				_846_v59 = Companion_Default___.Fm24((_this).F26(), (_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int), _833_v49, globalState)
				var _847_v60 _dafny.Map
				_ = _847_v60
				_847_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F26())
				var _848_v61 _dafny.Map
				_ = _848_v61
				_848_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_847_v60).Cardinality(), _this.F25())
				var _849_v64 _dafny.Set
				_ = _849_v64
				_849_v64 = _dafny.SetOf((_this).F24(), (_this).F24(), (_this).F24(), (_this).F24())
				var _850_v65 _dafny.Sequence
				_ = _850_v65
				_850_v65 = _dafny.SeqOf(_849_v64, _849_v64)
				var _851_v66 _dafny.Map
				_ = _851_v66
				_851_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC19_(_dafny.CodePoint('c')), false)
				var _852_v67 _dafny.Map
				_ = _852_v67
				_852_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F26()), _851_v66)
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))
				_ = _index137
				var _rhs169 _dafny.Map = func() _dafny.Map {
					var _coll36 = _dafny.NewMapBuilder()
					_ = _coll36
					for _iter40 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), _this.F25())).Keys().Elements()); ; {
						_compr_36, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _853_v62 _dafny.Int
						_853_v62 = interface{}(_compr_36).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F25(), _this.F25())).Contains(_853_v62) {
							_coll36.Add((_853_v62).Plus(_this.F25()), ((func() _dafny.Map {
								var _coll37 = _dafny.NewMapBuilder()
								_ = _coll37
								for _iter41 := _dafny.Iterate((_834_v50).Elements()); ; {
									_compr_37, _ok41 := _iter41()
									if !_ok41 {
										break
									}
									var _854_v63 _dafny.Int
									_854_v63 = interface{}(_compr_37).(_dafny.Int)
									if (_834_v50).Contains(_854_v63) {
										_coll37.Add((_854_v63).Plus(_this.F25()), (_this).F26())
									}
								}
								return _coll37.ToMap()
							}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int))).Cardinality(), (_this).F26()))).Cardinality())
						}
					}
					return _coll36.ToMap()
				}()
				_ = _rhs169
				var _rhs170 bool = _dafny.Companion_Sequence_.Equal(_850_v65, _dafny.Companion_Sequence_.Concatenate(_850_v65, _850_v65))
				_ = _rhs170
				var _rhs171 _dafny.Int = Companion_Default___.SafeModuloInt((_852_v67).Cardinality(), (_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int))
				_ = _rhs171
				var _lhs143 *GlobalState = globalState
				_ = _lhs143
				var _lhs144 _dafny.Array = _832_v48
				_ = _lhs144
				var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))
				_ = _lhs145
				_848_v61 = _rhs169
				_lhs143.F18 = _rhs170
				(_lhs144).ArraySet1(_rhs171, (_lhs145).Int())
				var _855_v68 _dafny.Sequence
				_ = _855_v68
				_855_v68 = _dafny.SeqOf((_dafny.Zero).Minus(_this.F25()), _dafny.IntOfUint32((_833_v49).Cardinality()))
				var _856_v69 _dafny.Set
				_ = _856_v69
				_856_v69 = _dafny.SetOf((_this).F26(), (_this).F26())
				var _857_v70 _dafny.MultiSet
				_ = _857_v70
				_857_v70 = _dafny.MultiSetOf((_this).F26())
				var _858_v71 _dafny.Map
				_ = _858_v71
				_858_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_this.F25(), (Companion_Default___.Fm25((_856_v69).Cardinality(), (_this).F26(), (_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int), globalState)).Cardinality()), _857_v70)
				var _859_v72 _dafny.Sequence
				_ = _859_v72
				_859_v72 = _dafny.SeqOf((_this).F26())
				var _860_v73 _dafny.Map
				_ = _860_v73
				_860_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int), (_this).F26())
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))
				_ = _index138
				var _rhs172 _dafny.Int = _dafny.IntOfInt64(33)
				_ = _rhs172
				var _rhs173 _dafny.Int = (func() _dafny.Int {
					if (_848_v61).Contains((_dafny.Zero).Minus(((_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_855_v68).Cardinality())))) {
						return (_848_v61).Get((_dafny.Zero).Minus(((_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int)).Times(_dafny.IntOfUint32((_855_v68).Cardinality())))).(_dafny.Int)
					}
					return (((func() _dafny.MultiSet {
						if (_858_v71).Contains(Companion_Default___.Fm26(globalState)) {
							return (_858_v71).Get(Companion_Default___.Fm26(globalState)).(_dafny.MultiSet)
						}
						return _857_v70
					})()).Intersection(_dafny.MultiSetFromSeq(_859_v72))).Cardinality()
				})()
				_ = _rhs173
				var _rhs174 bool = (func() bool {
					if (_860_v73).Contains((func() _dafny.Int {
						if (_857_v70).Contains((_this).F26()) {
							return (_857_v70).Multiplicity((_this).F26())
						}
						return (_dafny.Zero).Minus((_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int))
					})()) {
						return (_860_v73).Get((func() _dafny.Int {
							if (_857_v70).Contains((_this).F26()) {
								return (_857_v70).Multiplicity((_this).F26())
							}
							return (_dafny.Zero).Minus((_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int))
						})()).(bool)
					}
					return !((func() bool {
						if (_844_v57).Contains((_this).F26()) {
							return (_844_v57).Get((_this).F26()).(bool)
						}
						return (_this).F26()
					})())
				})()
				_ = _rhs174
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				var _lhs147 _dafny.Array = _832_v48
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))
				_ = _lhs148
				var _lhs149 *GlobalState = globalState
				_ = _lhs149
				_lhs146.F2 = _rhs172
				(_lhs147).ArraySet1(_rhs173, (_lhs148).Int())
				_lhs149.F18 = _rhs174
			} else {
				var _861_v74 _dafny.Array
				_ = _861_v74
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_24
				var _nw158 _dafny.Array
				_ = _nw158
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw158 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Sequence = func(_862_i4 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_this.F25(), _this.F25())
					}
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw158 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw158).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw158).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_861_v74 = _nw158
				var _863_v75 _dafny.Map
				_ = _863_v75
				_863_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_861_v74, true)
				(globalState).F18 = (func() bool {
					if (_863_v75).Contains(_861_v74) {
						return (_863_v75).Get(_861_v74).(bool)
					}
					return (_this).F26()
				})()
				var _rhs175 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(913))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_864_v49 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_865_i5 _dafny.Int) _dafny.Sequence {
						return _864_v49
					}
				})(_833_v49)))
				_ = _rhs175
				var _rhs176 bool = (_this).F26()
				_ = _rhs176
				_837_v53 = _rhs175
				r1 = _rhs176
				(globalState).F2 = _this.F25()
				var _866_v76 *C0
				_ = _866_v76
				var _nw159 *C0 = New_C0_()
				_ = _nw159
				_nw159.Ctor__()
				_866_v76 = _nw159
				_866_v76 = _866_v76
				var _867_v77 _dafny.Sequence
				_ = _867_v77
				_867_v77 = _dafny.SeqOf(!((_this).F26()))
				var _868_v78 _dafny.Array
				_ = _868_v78
				var _nwElement0_33 _dafny.Sequence = _867_v77
				_ = _nwElement0_33
				var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(11))
				_ = _nw160
				(_nw160).ArraySet1(_nwElement0_33, 0)
				(_nw160).ArraySet1(Companion_Default___.Fm7(globalState), 1)
				(_nw160).ArraySet1(_867_v77, 2)
				(_nw160).ArraySet1(_867_v77, 3)
				(_nw160).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_867_v77, (Companion_Default___.SafeIndex((_832_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(676), _dafny.ArrayLen((_832_v48), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_867_v77).Cardinality()))).Uint32(), (_this).F26()), _867_v77), 4)
				(_nw160).ArraySet1(_dafny.SeqOf(!((_this).F26()), (_this).F26(), (_this).F26(), (_this).F26(), (_this).F26()), 5)
				(_nw160).ArraySet1(_867_v77, 6)
				(_nw160).ArraySet1(_dafny.SeqOf(false, (_this).F26()), 7)
				(_nw160).ArraySet1(_867_v77, 8)
				(_nw160).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F26(), (_this).F26()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-115), _dafny.IntOfUint32((_dafny.SeqOf((_this).F26(), (_this).F26())).Cardinality()))).Uint32(), (_this).F26()), _867_v77), 9)
				(_nw160).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_867_v77, _867_v77), 10)
				_868_v78 = _nw160
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_868_v78), 0))
				_ = _index139
				(_868_v78).ArraySet1(_867_v77, (_index139).Int())
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_868_v78), 0))
				_ = _index140
				(_868_v78).ArraySet1(_867_v77, (_index140).Int())
			}
		} else {
			var _869___mcc_h0 _dafny.Array = _source13.Get_().(D4_DC7).Cf12
			_ = _869___mcc_h0
			var _870_cf12 _dafny.Array = _869___mcc_h0
			_ = _870_cf12
			(globalState).F19 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(323))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}(func(_871_i6 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uuvoe"), _dafny.UnicodeSeqOfUtf8Bytes("knnyuorh"))).Cardinality())
			}))
			var _872_v79 _dafny.Array
			_ = _872_v79
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_25
			var _nw161 _dafny.Array
			_ = _nw161
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw161 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Int = func(_873_i7 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_873_i7, (_dafny.SetOf((_this).F26(), (_this).F26())).Cardinality())
				}
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw161 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw161).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw161).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_872_v79 = _nw161
			_872_v79 = _872_v79
			r1 = (_this).F26()
			if (_this).F26() {
				var _874_v80 _dafny.Sequence
				_ = _874_v80
				_874_v80 = _dafny.UnicodeSeqOfUtf8Bytes("ksmojox")
				(globalState).F10 = _dafny.IntOfUint32((_874_v80).Cardinality())
				var _875_v81 _dafny.Sequence
				_ = _875_v81
				_875_v81 = _dafny.SeqOf((_this).F26(), !((_this).F26()))
				var _876_v82 _dafny.Array
				_ = _876_v82
				var _nwElement0_34 bool = (_dafny.IntOfInt64(502)).Cmp(_this.F25()) > 0
				_ = _nwElement0_34
				var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(10))
				_ = _nw162
				(_nw162).ArraySet1(_nwElement0_34, 0)
				(_nw162).ArraySet1(true, 1)
				(_nw162).ArraySet1((_this).F26(), 2)
				(_nw162).ArraySet1((_this).F26(), 3)
				(_nw162).ArraySet1((_this).F26(), 4)
				(_nw162).ArraySet1((_this).F26(), 5)
				(_nw162).ArraySet1(true, 6)
				(_nw162).ArraySet1((func() bool {
					if (_this).F26() {
						return (_this).F26()
					}
					return (_this).F26()
				})(), 7)
				(_nw162).ArraySet1((_this).F26(), 8)
				(_nw162).ArraySet1(_dafny.Companion_Sequence_.Equal(_875_v81, _875_v81), 9)
				_876_v82 = _nw162
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_876_v82), 0))
				_ = _index141
				(_876_v82).ArraySet1((_this).F26(), (_index141).Int())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_876_v82), 0))
				_ = _index142
				(_876_v82).ArraySet1((_this).F26(), (_index142).Int())
				var _877_v83 _dafny.Array
				_ = _877_v83
				var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
				_ = _nw163
				_877_v83 = _nw163
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_877_v83), 0))
				_ = _index143
				(_877_v83).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_874_v80, _874_v80), (_index143).Int())
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_877_v83), 0))
				_ = _index144
				var _rhs177 _dafny.Array = _876_v82
				_ = _rhs177
				var _rhs178 _dafny.Sequence = _874_v80
				_ = _rhs178
				var _lhs150 _dafny.Array = _877_v83
				_ = _lhs150
				var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_877_v83), 0))
				_ = _lhs151
				_876_v82 = _rhs177
				(_lhs150).ArraySet1(_rhs178, (_lhs151).Int())
				var _878_v84 *C4
				_ = _878_v84
				var _nw164 *C4 = New_C4_()
				_ = _nw164
				_nw164.Ctor__(Companion_Default___.Fm27(_this.F25(), globalState), Companion_Default___.Fm0(_this.F25(), (_this).F26(), false, _this.F25(), globalState), (_this).F24(), _this.F25())
				_878_v84 = _nw164
				var _879_v85 *C3
				_ = _879_v85
				var _nw165 *C3 = New_C3_()
				_ = _nw165
				_nw165.Ctor__((_this).F26(), Companion_Default___.Fm10((_878_v84).F30(), (_878_v84).F30(), (_this).F26(), (_876_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_876_v82), 0))).Int()).(bool), globalState), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_874_v80).Cardinality()), _this.F25()))
				_879_v85 = _nw165
			} else {
				var _880_v86 _dafny.Sequence
				_ = _880_v86
				_880_v86 = _dafny.UnicodeSeqOfUtf8Bytes("xukawq")
				var _881_v87 _dafny.MultiSet
				_ = _881_v87
				_881_v87 = _dafny.MultiSetOf(_880_v86, _dafny.UnicodeSeqOfUtf8Bytes("ylh"), _880_v86, _880_v86, _880_v86)
				var _882_v88 _dafny.Sequence
				_ = _882_v88
				_882_v88 = _dafny.UnicodeSeqOfUtf8Bytes("qxgte")
				var _883_v89 _dafny.Sequence
				_ = _883_v89
				_883_v89 = _dafny.SeqOf((_this).F26(), (_this).F26())
				var _884_v90 *C3
				_ = _884_v90
				var _nw166 *C3 = New_C3_()
				_ = _nw166
				_nw166.Ctor__(!((_881_v87).IsDisjointFrom(_dafny.MultiSetOf((_882_v88)))), (_this).F24(), _dafny.IntOfUint32((_883_v89).Cardinality()))
				_884_v90 = _nw166
				var _885_v91 _dafny.Array
				_ = _885_v91
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_26
				var _nw167 _dafny.Array
				_ = _nw167
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw167 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Sequence = (func(_886_v86 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_887_i8 _dafny.Int) _dafny.Sequence {
							return _886_v86
						}
					})(_880_v86)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw167 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw167).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw167).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_885_v91 = _nw167
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_885_v91), 0))
				_ = _index145
				(_885_v91).ArraySet1(_880_v86, (_index145).Int())
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_885_v91), 0))
				_ = _index146
				(_885_v91).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_880_v86, _dafny.Companion_Sequence_.Concatenate(_880_v86, _dafny.UnicodeSeqOfUtf8Bytes("ixxmbm"))), (_index146).Int())
				(globalState).F22 = _dafny.IntOfUint32((_880_v86).Cardinality())
				(globalState).F3 = _880_v86
				(globalState).F18 = (_this).F26()
			}
		}
		var _888_v92 _dafny.Array
		_ = _888_v92
		var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
		_ = _nw168
		_888_v92 = _nw168
		for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_888_v92), 0))); ; {
			_guard_loop_4, _ok42 := _iter42()
			if !_ok42 {
				break
			}
			var _889_i9 _dafny.Int
			_889_i9 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_889_i9).Sign() != -1) && ((_889_i9).Cmp(_dafny.ArrayLen((_888_v92), 0)) < 0)) {
				(_888_v92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gvibnejd"), _dafny.UnicodeSeqOfUtf8Bytes("polnxwjt")), (_889_i9).Int())
			}
		}
		var _890_v94 _dafny.Array
		_ = _890_v94
		var _len0_27 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_27
		var _nw169 _dafny.Array
		_ = _nw169
		if _len0_27.Cmp(_dafny.Zero) == 0 {
			_nw169 = _dafny.NewArray(_len0_27)
		} else {
			var _init27 func(_dafny.Int) _dafny.Map = func(_891_i10 _dafny.Int) _dafny.Map {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F26())).Merge(func() _dafny.Map {
					var _coll38 = _dafny.NewMapBuilder()
					_ = _coll38
					for _iter43 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F26())).Keys().Elements()); ; {
						_compr_38, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _892_v93 _dafny.CodePoint
						_892_v93 = interface{}(_compr_38).(_dafny.CodePoint)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F24(), (_this).F26())).Contains(_892_v93) {
							_coll38.Add(_892_v93, (_this).F26())
						}
					}
					return _coll38.ToMap()
				}())
			}
			_ = _init27
			var _element0_27 = _init27(_dafny.Zero)
			_ = _element0_27
			_nw169 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
			(_nw169).ArraySet1(_element0_27, 0)
			var _nativeLen0_27 = (_len0_27).Int()
			_ = _nativeLen0_27
			for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
				(_nw169).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
			}
		}
		_890_v94 = _nw169
		_890_v94 = _890_v94
		var _893_v95 _dafny.Array
		_ = _893_v95
		var _nw170 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
		_ = _nw170
		_893_v95 = _nw170
		var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_893_v95), 0))
		_ = _index147
		(_893_v95).ArraySet1((_this).F26(), (_index147).Int())
		var _894_v96 _dafny.Set
		_ = _894_v96
		_894_v96 = _dafny.SetOf((_this).F26(), (_this).F26())
		var _895_v97 _dafny.Map
		_ = _895_v97
		_895_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_894_v96, _this.F25())
		var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_893_v95), 0))
		_ = _index148
		(_893_v95).ArraySet1((func() bool {
			if (_this).F26() {
				return Companion_Default___.Fm0((_895_v97).Cardinality(), (_this).F26(), false, _this.F25(), globalState)
			}
			return (_this).F26()
		})(), (_index148).Int())
		var _896_v98 _dafny.Map
		_ = _896_v98
		_896_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F25()).Times(_this.F25()), _dafny.IntOfInt64(198))
		_896_v98 = (_896_v98).Update(_this.F25(), Companion_Default___.SafeModuloInt(_this.F25(), _this.F25()))
		var _897_v99 _dafny.Map
		_ = _897_v99
		_897_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_893_v95, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))
		r0 = (_897_v99).Merge((_897_v99).Merge(_897_v99))
		r1 = false
		return r0, r1
	}
}
func (_this *C5) F26() bool {
	{
		return _this._f26
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	F23 _dafny.Map
}

func New_C6_() *C6 {
	_this := C6{}

	_this.F23 = _dafny.EmptyMap
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

func (_this *C6) Ctor__(f23 _dafny.Map) {
	{
		(_this).F23 = f23
	}
}
func (_this *C6) Fm1(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-509))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("clwiitdbt")).Cardinality())))).Equals((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-196))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}(func(_898_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		}))).Cardinality()))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(-610))))
	}
}
func (_this *C6) Fm2(globalState *GlobalState) bool {
	{
		return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(339))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg58 _dafny.Int) interface{} {
				return coer58(arg58)
			}
		}(func(_899_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})), _dafny.UnicodeSeqOfUtf8Bytes("rideqr"))).Cardinality())).Cmp((_dafny.IntOfInt64(-714)).Plus((func() _dafny.Map {
			var _coll39 = _dafny.NewMapBuilder()
			_ = _coll39
			for _iter44 := _dafny.Iterate((_dafny.MultiSetOf(_this.F23)).Elements()); ; {
				_compr_39, _ok44 := _iter44()
				if !_ok44 {
					break
				}
				var _900_v0 _dafny.Map
				_900_v0 = interface{}(_compr_39).(_dafny.Map)
				if (_dafny.MultiSetOf(_this.F23)).Contains(_900_v0) {
					_coll39.Add(_900_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yf")).Cardinality()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t'))).Cardinality()))).Cardinality())))
				}
			}
			return _coll39.ToMap()
		}()).Cardinality())) <= 0
	}
}
func (_this *C6) M0(p0 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _901_v0 bool
		_ = _901_v0
		_901_v0 = true
		if _901_v0 {
			(globalState).F2 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(p0, p0), _dafny.IntOfInt64(203))
			var _902_v1 _dafny.Sequence
			_ = _902_v1
			_902_v1 = _dafny.SeqOf(_901_v0)
			if (_902_v1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_902_v1).Cardinality()))).Uint32()).(bool) {
				(globalState).F19 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(806))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}((func(_903_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_904_i0 _dafny.Int) _dafny.Int {
						return (_903_p0).Times(_903_p0)
					}
				})(p0)))
				(globalState).F10 = (p0).Times((p0).Minus(p0))
				var _905_v2 _dafny.Map
				_ = _905_v2
				_905_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_this.F23).Contains(p0) {
						return (_this.F23).Get(p0).(bool)
					}
					return _901_v0
				})(), _901_v0)
				var _906_v3 D2
				_ = _906_v3
				_906_v3 = Companion_D2_.Create_DC2_(_901_v0)
				var _907_v4 _dafny.Sequence
				_ = _907_v4
				_907_v4 = _dafny.UnicodeSeqOfUtf8Bytes("yxdeimc")
				var _908_v5 _dafny.CodePoint
				_ = _908_v5
				_908_v5 = _dafny.CodePoint('n')
				var _909_v6 _dafny.Array
				_ = _909_v6
				var _nwElement0_35 bool = true
				_ = _nwElement0_35
				var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(15))
				_ = _nw171
				(_nw171).ArraySet1(_nwElement0_35, 0)
				(_nw171).ArraySet1(_901_v0, 1)
				(_nw171).ArraySet1((func() bool {
					if (_905_v2).Contains(_901_v0) {
						return (_905_v2).Get(_901_v0).(bool)
					}
					return _901_v0
				})(), 2)
				(_nw171).ArraySet1(_901_v0, 3)
				(_nw171).ArraySet1(_901_v0, 4)
				(_nw171).ArraySet1(_901_v0, 5)
				(_nw171).ArraySet1(_901_v0, 6)
				(_nw171).ArraySet1((_906_v3).Dtor_cf2(), 7)
				(_nw171).ArraySet1(!(!(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("lkxsn"), _907_v4))), 8)
				(_nw171).ArraySet1(true, 9)
				(_nw171).ArraySet1((func() bool {
					if _901_v0 {
						return _901_v0
					}
					return true
				})(), 10)
				(_nw171).ArraySet1(!_dafny.Companion_Sequence_.Contains(_907_v4, _908_v5), 11)
				(_nw171).ArraySet1(_901_v0, 12)
				(_nw171).ArraySet1(_901_v0, 13)
				(_nw171).ArraySet1(_901_v0, 14)
				_909_v6 = _nw171
				var _910_v7 _dafny.Sequence
				_ = _910_v7
				_910_v7 = _dafny.SeqOf(p0, p0)
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_909_v6), 0))
				_ = _index149
				(_909_v6).ArraySet1(!_dafny.Companion_Sequence_.Equal(_910_v7, _910_v7), (_index149).Int())
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_909_v6), 0))
				_ = _index150
				(_909_v6).ArraySet1((p0).Cmp(p0) <= 0, (_index150).Int())
				var _911_v9 _dafny.Set
				_ = _911_v9
				_911_v9 = _dafny.SetOf((func() _dafny.Map {
					if _901_v0 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_909_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_909_v6), 0))).Int()).(bool))
					}
					return _this.F23
				})(), _this.F23, _this.F23, func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(807), _dafny.IntOfInt64(609))); ; {
						_compr_40, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _912_v8 _dafny.Int
						_912_v8 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(807)).Cmp(_912_v8) <= 0) && ((_912_v8).Cmp(_dafny.IntOfInt64(609)) < 0) {
							_coll40.Add(Companion_Default___.SafeModuloInt(_912_v8, p0), (_909_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_909_v6), 0))).Int()).(bool))
						}
					}
					return _coll40.ToMap()
				}(), _this.F23)
				var _913_v10 _dafny.Array
				_ = _913_v10
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw172
				_913_v10 = _nw172
				var _914_v11 _dafny.Array
				_ = _914_v11
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_28
				var _nw173 _dafny.Array
				_ = _nw173
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw173 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.CodePoint = (func(_915_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_916_i1 _dafny.Int) _dafny.CodePoint {
							return _915_v5
						}
					})(_908_v5)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw173 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw173).ArraySet1CodePoint(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw173).ArraySet1CodePoint(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_914_v11 = _nw173
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_913_v10), 0))
				_ = _index151
				(_913_v10).ArraySet1(_914_v11, (_index151).Int())
				var _917_v12 _dafny.Sequence
				_ = _917_v12
				_917_v12 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}(func(_918_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))
				var _919_v13 _dafny.MultiSet
				_ = _919_v13
				_919_v13 = _dafny.MultiSetOf(_917_v12, _917_v12)
				var _920_v14 _dafny.Sequence
				_ = _920_v14
				_920_v14 = _dafny.SeqOf(_dafny.SeqOf((_this).Fm1(_919_v13, globalState), _901_v0, false))
				var _921_v15 _dafny.Map
				_ = _921_v15
				_921_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_920_v14).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_920_v14).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_913_v10), 0))
				_ = _index152
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_909_v6), 0))
				_ = _index153
				var _rhs179 _dafny.Set = (_911_v9).Intersection((Companion_D3_.Create_DC4_(_911_v9)).Dtor_cf5())
				_ = _rhs179
				var _rhs180 _dafny.Array = _914_v11
				_ = _rhs180
				var _rhs181 bool = !_dafny.Companion_Sequence_.Equal(_910_v7, _910_v7)
				_ = _rhs181
				var _rhs182 bool = _901_v0
				_ = _rhs182
				var _rhs183 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_902_v1, (func() _dafny.Sequence {
					if (_921_v15).Contains(p0) {
						return (_921_v15).Get(p0).(_dafny.Sequence)
					}
					return _dafny.SeqOf((_909_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_909_v6), 0))).Int()).(bool), (_909_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_909_v6), 0))).Int()).(bool))
				})())
				_ = _rhs183
				var _lhs152 _dafny.Array = _913_v10
				_ = _lhs152
				var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.ArrayLen((_913_v10), 0))
				_ = _lhs153
				var _lhs154 _dafny.Array = _909_v6
				_ = _lhs154
				var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_909_v6), 0))
				_ = _lhs155
				_911_v9 = _rhs179
				(_lhs152).ArraySet1(_rhs180, (_lhs153).Int())
				(_lhs154).ArraySet1(_rhs181, (_lhs155).Int())
				_901_v0 = _rhs182
				_902_v1 = _rhs183
				var _rhs184 _dafny.Int = (p0).Plus(p0)
				_ = _rhs184
				var _rhs185 _dafny.Sequence = _910_v7
				_ = _rhs185
				var _lhs156 *GlobalState = globalState
				_ = _lhs156
				_lhs156.F2 = _rhs184
				_910_v7 = _rhs185
			} else {
				var _922_v16 D11
				_ = _922_v16
				_922_v16 = Companion_D11_.Create_DC24_(_901_v0, p0)
				var _923_v17 _dafny.Map
				_ = _923_v17
				_923_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_922_v16, _901_v0)
				var _924_v18 _dafny.CodePoint
				_ = _924_v18
				_924_v18 = _dafny.CodePoint('p')
				var _925_v19 _dafny.Sequence
				_ = _925_v19
				_925_v19 = _dafny.UnicodeSeqOfUtf8Bytes("csjg")
				var _926_v20 *C4
				_ = _926_v20
				var _nw174 *C4 = New_C4_()
				_ = _nw174
				_nw174.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), _dafny.IntOfInt64(707)), (func() bool {
					if (_923_v17).Contains(_922_v16) {
						return (_923_v17).Get(_922_v16).(bool)
					}
					return _901_v0
				})(), _924_v18, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_925_v19).Cardinality()), p0))
				_926_v20 = _nw174
				var _927_v21 _dafny.Map
				_ = _927_v21
				_927_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), _dafny.CodePoint('k'))
				_927_v21 = (_927_v21).Update((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p0, p0)), _dafny.CodePoint('e'))
				var _928_v22 _dafny.Set
				_ = _928_v22
				_928_v22 = _dafny.SetOf(p0)
				var _929_v23 D12
				_ = _929_v23
				_929_v23 = Companion_D12_.Create_DC26_(_928_v22)
				r2 = Companion_Default___.Fm15(p0, (_dafny.Zero).Minus(p0), ((_929_v23).Dtor_cf39()).Cardinality(), globalState)
				var _930_v24 _dafny.Array
				_ = _930_v24
				var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw175
				_930_v24 = _nw175
				_930_v24 = _930_v24
				var _931_v25 D7
				_ = _931_v25
				_931_v25 = Companion_D7_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(703), _901_v0))
				var _932_v26 _dafny.Map
				_ = _932_v26
				_932_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v25, _902_v1)
				var _933_v27 _dafny.Map
				_ = _933_v27
				_933_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_926_v20).F30(), (_932_v26).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v25, _902_v1)))
				_933_v27 = (_933_v27).Update((_926_v20).F30(), (_932_v26).Update(_931_v25, _902_v1))
			}
			var _934_v28 _dafny.Array
			_ = _934_v28
			var _len0_29 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_29
			var _nw176 _dafny.Array
			_ = _nw176
			if _len0_29.Cmp(_dafny.Zero) == 0 {
				_nw176 = _dafny.NewArray(_len0_29)
			} else {
				var _init29 func(_dafny.Int) _dafny.Sequence = (func(_935_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_936_i3 _dafny.Int) _dafny.Sequence {
						return _935_v1
					}
				})(_902_v1)
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
			_934_v28 = _nw176
			var _937_v29 _dafny.Sequence
			_ = _937_v29
			_937_v29 = _dafny.UnicodeSeqOfUtf8Bytes("telw")
			var _938_v30 _dafny.Sequence
			_ = _938_v30
			_938_v30 = _dafny.SeqOf(_902_v1)
			var _nwElement0_36 _dafny.Sequence = _902_v1
			_ = _nwElement0_36
			var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(22))
			_ = _nw177
			(_nw177).ArraySet1(_nwElement0_36, 0)
			(_nw177).ArraySet1(_902_v1, 1)
			(_nw177).ArraySet1(_dafny.SeqOf(_901_v0, _901_v0, _901_v0), 2)
			(_nw177).ArraySet1(_902_v1, 3)
			(_nw177).ArraySet1(_902_v1, 4)
			(_nw177).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_902_v1, _dafny.Companion_Sequence_.Update(_902_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_937_v29).Cardinality()), _dafny.IntOfUint32((_902_v1).Cardinality()))).Uint32(), _901_v0)), 5)
			(_nw177).ArraySet1(_902_v1, 6)
			(_nw177).ArraySet1(_902_v1, 7)
			(_nw177).ArraySet1(_902_v1, 8)
			(_nw177).ArraySet1(_dafny.SeqOf(true, _901_v0, _901_v0, !((func() bool {
				if (_this.F23).Contains(_dafny.IntOfUint32((_902_v1).Cardinality())) {
					return (_this.F23).Get(_dafny.IntOfUint32((_902_v1).Cardinality())).(bool)
				}
				return _901_v0
			})())), 9)
			(_nw177).ArraySet1(_902_v1, 10)
			(_nw177).ArraySet1(_902_v1, 11)
			(_nw177).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_902_v1, _902_v1), 12)
			(_nw177).ArraySet1(_902_v1, 13)
			(_nw177).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_938_v30).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_938_v30).Cardinality()))).Uint32()).(_dafny.Sequence), _902_v1), 14)
			(_nw177).ArraySet1(_902_v1, 15)
			(_nw177).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_902_v1, _902_v1), 16)
			(_nw177).ArraySet1(_dafny.SeqOf(_901_v0), 17)
			(_nw177).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_902_v1, _902_v1), 18)
			(_nw177).ArraySet1(Companion_Default___.Fm7(globalState), 19)
			(_nw177).ArraySet1(_902_v1, 20)
			(_nw177).ArraySet1(_902_v1, 21)
			_934_v28 = _nw177
			var _939_v31 _dafny.Sequence
			_ = _939_v31
			_939_v31 = _dafny.SeqOf((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_dafny.SeqOf(_901_v0, false)).Cardinality()))
			var _rhs186 _dafny.Int = Companion_Default___.SafeDivisionInt((_939_v31).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_939_v31).Cardinality()))).Uint32()).(_dafny.Int), p0)
			_ = _rhs186
			var _rhs187 bool = _901_v0
			_ = _rhs187
			var _rhs188 bool = _901_v0
			_ = _rhs188
			var _lhs157 *GlobalState = globalState
			_ = _lhs157
			var _lhs158 *GlobalState = globalState
			_ = _lhs158
			_lhs157.F10 = _rhs186
			_lhs158.F18 = _rhs187
			_901_v0 = _rhs188
			r1 = p0
		} else {
			var _940_v32 _dafny.Sequence
			_ = _940_v32
			_940_v32 = _dafny.UnicodeSeqOfUtf8Bytes("kysw")
			var _941_v33 D9
			_ = _941_v33
			_941_v33 = Companion_D9_.Create_DC17_(_dafny.IntOfUint32((_940_v32).Cardinality()))
			var _942_v34 _dafny.Sequence
			_ = _942_v34
			_942_v34 = _dafny.SeqOf(p0)
			var _943_v35 _dafny.Map
			_ = _943_v35
			_943_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_941_v33, _dafny.Companion_Sequence_.Contains(_942_v34, p0))
			_943_v35 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC17_(p0), _901_v0)).Merge(_943_v35)).Merge(Companion_Default___.Fm28(p0, !(_901_v0), _901_v0, globalState))
			_901_v0 = false
			var _944_v36 _dafny.MultiSet
			_ = _944_v36
			_944_v36 = _dafny.MultiSetOf(_901_v0, _901_v0, true, _901_v0)
			var _945_v37 _dafny.Map
			_ = _945_v37
			_945_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_944_v36, p0)
			var _946_v38 _dafny.CodePoint
			_ = _946_v38
			_946_v38 = _dafny.CodePoint('u')
			var _947_v39 T0
			_ = _947_v39
			var _nw178 *C4 = New_C4_()
			_ = _nw178
			_nw178.Ctor__(_945_v37, _901_v0, _946_v38, p0)
			_947_v39 = _nw178
			var _948_v40 _dafny.Sequence
			_ = _948_v40
			_948_v40 = _dafny.SeqOf(_947_v39, _947_v39, _947_v39, _947_v39)
			var _949_v41 _dafny.Sequence
			_ = _949_v41
			_949_v41 = _948_v40
			var _950_v42 _dafny.Set
			_ = _950_v42
			_950_v42 = _dafny.SetOf(_949_v41)
			(globalState).F18 = !(!((_950_v42).IsDisjointFrom(_950_v42)))
			var _951_v44 _dafny.Sequence
			_ = _951_v44
			_951_v44 = _dafny.SeqOf(_944_v36)
			var _952_v45 *C4
			_ = _952_v45
			var _nw179 *C4 = New_C4_()
			_ = _nw179
			_nw179.Ctor__((func() _dafny.Map {
				var _coll41 = _dafny.NewMapBuilder()
				_ = _coll41
				for _iter46 := _dafny.Iterate((_951_v44).Elements()); ; {
					_compr_41, _ok46 := _iter46()
					if !_ok46 {
						break
					}
					var _953_v43 _dafny.MultiSet
					_953_v43 = interface{}(_compr_41).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_951_v44, _953_v43) {
						_coll41.Add(_953_v43, _dafny.IntOfInt64(943))
					}
				}
				return _coll41.ToMap()
			}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_944_v36, p0)), _901_v0, Companion_Default___.Fm10(!(_901_v0), _901_v0, _901_v0, _901_v0, globalState), p0)
			_952_v45 = _nw179
			var _954_v46 *C0
			_ = _954_v46
			var _nw180 *C0 = New_C0_()
			_ = _nw180
			_nw180.Ctor__()
			_954_v46 = _nw180
			var _955_v47 _dafny.Map
			_ = _955_v47
			_955_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_952_v45).F30(), _954_v46)
			var _rhs189 _dafny.Int = p0
			_ = _rhs189
			var _rhs190 _dafny.CodePoint = (_947_v39).F24()
			_ = _rhs190
			var _rhs191 *C4 = _952_v45
			_ = _rhs191
			var _rhs192 *C0 = (func() *C0 {
				if (_955_v47).Contains(_901_v0) {
					return (_955_v47).Get(_901_v0).(*C0)
				}
				return _954_v46
			})()
			_ = _rhs192
			var _lhs159 *GlobalState = globalState
			_ = _lhs159
			_lhs159.F11 = _rhs189
			_946_v38 = _rhs190
			_952_v45 = _rhs191
			_954_v46 = _rhs192
			(globalState).F10 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_940_v32, _dafny.Companion_Sequence_.Concatenate(_940_v32, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}((func(_956_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_957_i4 _dafny.Int) _dafny.CodePoint {
					return _956_v38
				}
			})(_946_v38)))))).Cardinality())
		}
		if _901_v0 {
			var _958_v48 _dafny.Sequence
			_ = _958_v48
			_958_v48 = _dafny.UnicodeSeqOfUtf8Bytes("sepaclhxe")
			var _959_v49 _dafny.CodePoint
			_ = _959_v49
			_959_v49 = _dafny.CodePoint('g')
			var _960_v50 _dafny.Map
			_ = _960_v50
			_960_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _959_v49)
			var _961_v51 *C5
			_ = _961_v51
			var _nw181 *C5 = New_C5_()
			_ = _nw181
			_nw181.Ctor__(Companion_Default___.Fm0(_dafny.IntOfUint32((_958_v48).Cardinality()), !(_901_v0), _901_v0, (_dafny.Zero).Minus(p0), globalState), (func() _dafny.CodePoint {
				if (_960_v50).Contains(p0) {
					return (_960_v50).Get(p0).(_dafny.CodePoint)
				}
				return _959_v49
			})(), p0)
			_961_v51 = _nw181
			var _962_v52 _dafny.Set
			_ = _962_v52
			_962_v52 = _dafny.SetOf(((_dafny.Zero).Minus(p0)).Cmp(p0) > 0)
			_962_v52 = _962_v52
			var _963_v53 _dafny.Array
			_ = _963_v53
			var _nw182 _dafny.Array = _dafny.NewArrayWithValue(Companion_D9_.Default(), _dafny.IntOfInt64(18))
			_ = _nw182
			_963_v53 = _nw182
			var _964_v54 _dafny.Map
			_ = _964_v54
			_964_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_961_v51).F26(), Companion_D3_.Create_DC6_(p0, p0, false))
			var _965_v55 D9
			_ = _965_v55
			_965_v55 = Companion_D9_.Create_DC16_(_964_v54)
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_963_v53), 0))
			_ = _index154
			(_963_v53).ArraySet1(_965_v55, (_index154).Int())
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_963_v53), 0))
			_ = _index155
			(_963_v53).ArraySet1(Companion_D9_.Create_DC16_((Companion_D9_.Create_DC16_(_964_v54)).Dtor_cf24()), (_index155).Int())
			(globalState).F18 = !(!(_901_v0))
			var _966_v56 _dafny.Array
			_ = _966_v56
			var _nwElement0_37 _dafny.CodePoint = _959_v49
			_ = _nwElement0_37
			var _nw183 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(2))
			_ = _nw183
			(_nw183).ArraySet1CodePoint(_nwElement0_37, 0)
			(_nw183).ArraySet1CodePoint(_959_v49, 1)
			_966_v56 = _nw183
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_966_v56), 0))
			_ = _index156
			(_966_v56).ArraySet1CodePoint(_959_v49, (_index156).Int())
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_966_v56), 0))
			_ = _index157
			(_966_v56).ArraySet1CodePoint(_959_v49, (_index157).Int())
		} else {
			var _967_v57 _dafny.MultiSet
			_ = _967_v57
			_967_v57 = _dafny.MultiSetOf(p0, p0)
			var _968_v58 D7
			_ = _968_v58
			_968_v58 = Companion_D7_.Create_DC14_(_901_v0, p0, p0)
			var _source14 D7 = (func() D7 {
				if !((_967_v57).Equals(_967_v57)) {
					return _968_v58
				}
				return _968_v58
			})()
			_ = _source14
			if _source14.Is_DC14() {
				var _969___mcc_h0 bool = _source14.Get_().(D7_DC14).Cf20
				_ = _969___mcc_h0
				var _970___mcc_h1 _dafny.Int = _source14.Get_().(D7_DC14).Cf21
				_ = _970___mcc_h1
				var _971___mcc_h2 _dafny.Int = _source14.Get_().(D7_DC14).Cf22
				_ = _971___mcc_h2
				var _972_cf22 _dafny.Int = _971___mcc_h2
				_ = _972_cf22
				var _973_cf21 _dafny.Int = _970___mcc_h1
				_ = _973_cf21
				var _974_cf20 bool = _969___mcc_h0
				_ = _974_cf20
				var _975_v59 _dafny.CodePoint
				_ = _975_v59
				_975_v59 = _dafny.CodePoint('w')
				_975_v59 = (func() _dafny.CodePoint {
					if (_974_cf20) || (_901_v0) {
						return _975_v59
					}
					return _975_v59
				})()
				_975_v59 = (func() _dafny.CodePoint {
					if false {
						return _975_v59
					}
					return _975_v59
				})()
				(globalState).F18 = !(false) || ((_974_cf20) && (_901_v0))
				var _976_v60 _dafny.Array
				_ = _976_v60
				var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw184
				_976_v60 = _nw184
				var _977_v61 _dafny.MultiSet
				_ = _977_v61
				_977_v61 = _dafny.MultiSetOf(_974_cf20)
				var _978_v62 _dafny.Map
				_ = _978_v62
				_978_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_977_v61, _972_cf22)
				var _979_v63 *C4
				_ = _979_v63
				var _nw185 *C4 = New_C4_()
				_ = _nw185
				_nw185.Ctor__(_978_v62, _901_v0, _975_v59, _972_cf22)
				_979_v63 = _nw185
				var _980_v64 _dafny.Map
				_ = _980_v64
				_980_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_979_v63, (_979_v63).F30())
				var _981_v65 _dafny.Map
				_ = _981_v65
				_981_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(596), _980_v64)
				var _982_v66 _dafny.Sequence
				_ = _982_v66
				_982_v66 = _dafny.UnicodeSeqOfUtf8Bytes("x")
				var _983_v67 _dafny.Set
				_ = _983_v67
				_983_v67 = _dafny.SetOf(_dafny.IntOfInt64(440))
				var _984_v68 _dafny.Sequence
				_ = _984_v68
				_984_v68 = _dafny.SeqOf((_979_v63).F30(), true, (_979_v63).F30(), _901_v0, _974_cf20)
				var _985_v69 _dafny.Array
				_ = _985_v69
				var _nwElement0_38 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(false, _974_cf20, _974_cf20)).Cardinality())
				_ = _nwElement0_38
				var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(8))
				_ = _nw186
				(_nw186).ArraySet1(_nwElement0_38, 0)
				(_nw186).ArraySet1(p0, 1)
				(_nw186).ArraySet1((_dafny.MultiSetOf(_974_cf20)).Cardinality(), 2)
				(_nw186).ArraySet1((_981_v65).Cardinality(), 3)
				(_nw186).ArraySet1(_dafny.IntOfUint32((_982_v66).Cardinality()), 4)
				(_nw186).ArraySet1((_983_v67).Cardinality(), 5)
				(_nw186).ArraySet1(Companion_Default___.Fm17(_901_v0, (_979_v63).F30(), Companion_Default___.Fm0(_dafny.IntOfUint32((_984_v68).Cardinality()), _974_cf20, _974_cf20, p0, globalState), globalState), 6)
				(_nw186).ArraySet1(_dafny.IntOfUint32((_984_v68).Cardinality()), 7)
				_985_v69 = _nw186
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_976_v60), 0))
				_ = _index158
				(_976_v60).ArraySet1(_985_v69, (_index158).Int())
				var _986_v70 D5
				_ = _986_v70
				_986_v70 = Companion_D5_.Create_DC9_(_985_v69)
				var _pat_let_tv9 = _985_v69
				_ = _pat_let_tv9
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_976_v60), 0))
				_ = _index159
				(_976_v60).ArraySet1((func(_pat_let12_0 D5) D5 {
					return func(_987_dt__update__tmp_h0 D5) D5 {
						return func(_pat_let13_0 _dafny.Array) D5 {
							return func(_988_dt__update_hcf13_h0 _dafny.Array) D5 {
								return Companion_D5_.Create_DC9_(_988_dt__update_hcf13_h0)
							}(_pat_let13_0)
						}(_pat_let_tv9)
					}(_pat_let12_0)
				}(_986_v70)).Dtor_cf13(), (_index159).Int())
			} else {
				var _989___mcc_h3 _dafny.Map = _source14.Get_().(D7_DC13).Cf19
				_ = _989___mcc_h3
				var _990_cf19 _dafny.Map = _989___mcc_h3
				_ = _990_cf19
				var _991_v71 _dafny.Array
				_ = _991_v71
				var _nwElement0_39 bool = _901_v0
				_ = _nwElement0_39
				var _nw187 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(4))
				_ = _nw187
				(_nw187).ArraySet1(_nwElement0_39, 0)
				(_nw187).ArraySet1(_901_v0, 1)
				(_nw187).ArraySet1(_901_v0, 2)
				(_nw187).ArraySet1(_901_v0, 3)
				_991_v71 = _nw187
				var _992_v72 _dafny.Map
				_ = _992_v72
				_992_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v0, p0)
				var _993_v73 _dafny.Map
				_ = _993_v73
				_993_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _994_v74 _dafny.Map
				_ = _994_v74
				_994_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v71, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_992_v72).Cardinality(), (_993_v73).Cardinality())).Merge(_993_v73))
				var _995_v75 _dafny.Array
				_ = _995_v75
				var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw188
				_995_v75 = _nw188
				var _996_v76 _dafny.Array
				_ = _996_v76
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_30
				var _nw189 _dafny.Array
				_ = _nw189
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw189 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.CodePoint = func(_997_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('q')
					}
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw189 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw189).ArraySet1CodePoint(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw189).ArraySet1CodePoint(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_996_v76 = _nw189
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_995_v75), 0))
				_ = _index160
				(_995_v75).ArraySet1(_996_v76, (_index160).Int())
				var _998_v77 _dafny.Sequence
				_ = _998_v77
				_998_v77 = _dafny.SeqOf(_994_v74, _994_v74, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v71, (_993_v73).Update(p0, (Companion_D11_.Create_DC24_(Companion_Default___.Fm0(p0, _901_v0, _901_v0, p0, globalState), p0)).Dtor_cf37())), _994_v74)
				var _999_v78 _dafny.Sequence
				_ = _999_v78
				_999_v78 = _dafny.SeqOf((_994_v74).Merge(_994_v74), (_994_v74).Merge(_994_v74), (_998_v77).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_998_v77).Cardinality()))).Uint32()).(_dafny.Map), _994_v74)
				var _1000_v79 _dafny.Array
				_ = _1000_v79
				var _nwElement0_40 _dafny.Int = p0
				_ = _nwElement0_40
				var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(12))
				_ = _nw190
				(_nw190).ArraySet1(_nwElement0_40, 0)
				(_nw190).ArraySet1(p0, 1)
				(_nw190).ArraySet1(p0, 2)
				(_nw190).ArraySet1(p0, 3)
				(_nw190).ArraySet1(p0, 4)
				(_nw190).ArraySet1(_dafny.IntOfInt64(838), 5)
				(_nw190).ArraySet1(p0, 6)
				(_nw190).ArraySet1(p0, 7)
				(_nw190).ArraySet1(p0, 8)
				(_nw190).ArraySet1(p0, 9)
				(_nw190).ArraySet1(p0, 10)
				(_nw190).ArraySet1(p0, 11)
				_1000_v79 = _nw190
				var _1001_v80 D13
				_ = _1001_v80
				_1001_v80 = Companion_D13_.Create_DC28_(_dafny.SetOf(_1000_v79))
				var _1002_v81 D5
				_ = _1002_v81
				_1002_v81 = Companion_D5_.Create_DC9_(_1000_v79)
				var _1003_v82 _dafny.Set
				_ = _1003_v82
				_1003_v82 = _dafny.SetOf(_1000_v79, (_1002_v81).Dtor_cf13(), _1000_v79, _1000_v79)
				var _1004_v83 _dafny.MultiSet
				_ = _1004_v83
				_1004_v83 = _dafny.MultiSetOf(_901_v0)
				var _1005_v84 _dafny.Sequence
				_ = _1005_v84
				_1005_v84 = _dafny.UnicodeSeqOfUtf8Bytes("sc")
				var _1006_v85 D12
				_ = _1006_v85
				_1006_v85 = Companion_D12_.Create_DC27_((func() _dafny.Int {
					if (_967_v57).Contains(p0) {
						return (_967_v57).Multiplicity(p0)
					}
					return (func() _dafny.Int {
						if (_1004_v83).Contains(_901_v0) {
							return (_1004_v83).Multiplicity(_901_v0)
						}
						return _dafny.IntOfUint32((_1005_v84).Cardinality())
					})()
				})(), _901_v0)
				var _1007_v86 _dafny.Map
				_ = _1007_v86
				_1007_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _901_v0)
				var _1008_v87 _dafny.Map
				_ = _1008_v87
				_1008_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1007_v86)
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_995_v75), 0))
				_ = _index161
				var _rhs193 _dafny.Map = (_999_v78).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_999_v78).Cardinality()))).Uint32()).(_dafny.Map)
				_ = _rhs193
				var _rhs194 _dafny.Array = _996_v76
				_ = _rhs194
				var _rhs195 _dafny.Int = p0
				_ = _rhs195
				var _rhs196 bool = ((_1001_v80).Dtor_cf42()).IsDisjointFrom(_1003_v82)
				_ = _rhs196
				var _rhs197 bool = (Companion_D7_.Create_DC14_((_1006_v85).Dtor_cf41(), p0, ((func() _dafny.Map {
					if (_1008_v87).Contains(p0) {
						return (_1008_v87).Get(p0).(_dafny.Map)
					}
					return _1007_v86
				})()).Cardinality())).Dtor_cf20()
				_ = _rhs197
				var _lhs160 _dafny.Array = _995_v75
				_ = _lhs160
				var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_995_v75), 0))
				_ = _lhs161
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				var _lhs163 *GlobalState = globalState
				_ = _lhs163
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				_994_v74 = _rhs193
				(_lhs160).ArraySet1(_rhs194, (_lhs161).Int())
				_lhs162.F16 = _rhs195
				_lhs163.F18 = _rhs196
				_lhs164.F18 = _rhs197
				var _1009_v88 _dafny.CodePoint
				_ = _1009_v88
				_1009_v88 = _dafny.CodePoint('k')
				_1009_v88 = Companion_Default___.Fm10(_901_v0, (false) || (_901_v0), _901_v0, _901_v0, globalState)
				var _1010_v89 _dafny.MultiSet
				_ = _1010_v89
				_1010_v89 = _dafny.MultiSetOf(Companion_Default___.Fm29(p0, globalState))
				_993_v73 = Companion_Default___.Fm11(_1009_v88, (_this).Fm1(_1010_v89, globalState), globalState)
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_991_v71), 0))
				_ = _index162
				(_991_v71).ArraySet1(!(_901_v0), (_index162).Int())
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_991_v71), 0))
				_ = _index163
				(_991_v71).ArraySet1(!(_901_v0), (_index163).Int())
			}
			var _1011_v90 _dafny.Sequence
			_ = _1011_v90
			_1011_v90 = _dafny.UnicodeSeqOfUtf8Bytes("sc")
			var _1012_v91 _dafny.Sequence
			_ = _1012_v91
			_1012_v91 = _dafny.SeqOf(_dafny.IntOfUint32((_1011_v90).Cardinality()))
			(globalState).F11 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1012_v91).Cardinality()), p0)
			var _1013_v92 D11
			_ = _1013_v92
			_1013_v92 = Companion_D11_.Create_DC24_(_901_v0, p0)
			var _1014_v93 _dafny.Sequence
			_ = _1014_v93
			_1014_v93 = _dafny.SeqOf((_1013_v92).Dtor_cf36())
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_1014_v93, _1014_v93) {
				(globalState).F11 = p0
				var _1015_v94 *C2
				_ = _1015_v94
				var _nw191 *C2 = New_C2_()
				_ = _nw191
				_nw191.Ctor__(p0, _901_v0)
				_1015_v94 = _nw191
				var _1016_v95 _dafny.Map
				_ = _1016_v95
				_1016_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v0, _1015_v94)
				var _1017_v96 _dafny.Map
				_ = _1017_v96
				_1017_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v0, _901_v0)
				var _1018_v97 _dafny.Set
				_ = _1018_v97
				_1018_v97 = _dafny.SetOf(_1011_v90)
				var _1019_v98 _dafny.MultiSet
				_ = _1019_v98
				_1019_v98 = _dafny.MultiSetOf(_901_v0)
				var _1020_v99 _dafny.Sequence
				_ = _1020_v99
				_1020_v99 = _dafny.UnicodeSeqOfUtf8Bytes("tkymklej")
				var _1021_v100 _dafny.MultiSet
				_ = _1021_v100
				_1021_v100 = _dafny.MultiSetOf(_1020_v99, _1011_v90, _1011_v90, _1020_v99, _1020_v99)
				var _1022_v101 _dafny.Array
				_ = _1022_v101
				var _nwElement0_41 bool = ((_1016_v95).Cardinality()).Cmp((_1015_v94).F27()) != 0
				_ = _nwElement0_41
				var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(20))
				_ = _nw192
				(_nw192).ArraySet1(_nwElement0_41, 0)
				(_nw192).ArraySet1((_1015_v94).F28(), 1)
				(_nw192).ArraySet1((func() bool {
					if (_1017_v96).Contains(true) {
						return (_1017_v96).Get(true).(bool)
					}
					return (_1015_v94).F28()
				})(), 2)
				(_nw192).ArraySet1((_1015_v94).F28(), 3)
				(_nw192).ArraySet1((_901_v0) || ((_1015_v94).F28()), 4)
				(_nw192).ArraySet1((_1015_v94).F28(), 5)
				(_nw192).ArraySet1(_901_v0, 6)
				(_nw192).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1014_v93, (_1015_v94).F28()), 7)
				(_nw192).ArraySet1((_1015_v94).F28(), 8)
				(_nw192).ArraySet1((_1015_v94).F28(), 9)
				(_nw192).ArraySet1(_901_v0, 10)
				(_nw192).ArraySet1((_1018_v97).IsDisjointFrom(_1018_v97), 11)
				(_nw192).ArraySet1((_1015_v94).F28(), 12)
				(_nw192).ArraySet1(Companion_Default___.Fm0(p0, (_1015_v94).F28(), _901_v0, p0, globalState), 13)
				(_nw192).ArraySet1((_1019_v98).IsDisjointFrom(_1019_v98), 14)
				(_nw192).ArraySet1(_901_v0, 15)
				(_nw192).ArraySet1(!((_this).Fm1(_1021_v100, globalState)), 16)
				(_nw192).ArraySet1(_901_v0, 17)
				(_nw192).ArraySet1((p0).Cmp(_dafny.IntOfInt64(348)) < 0, 18)
				(_nw192).ArraySet1(!(((_1015_v94).F28()) && (false)), 19)
				_1022_v101 = _nw192
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_1022_v101), 0))
				_ = _index164
				(_1022_v101).ArraySet1(_901_v0, (_index164).Int())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_1022_v101), 0))
				_ = _index165
				(_1022_v101).ArraySet1(!(_901_v0) || (false), (_index165).Int())
				var _1023_v102 D9
				_ = _1023_v102
				_1023_v102 = Companion_D9_.Create_DC17_(_dafny.IntOfInt64(338))
				var _1024_v103 _dafny.Array
				_ = _1024_v103
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_31
				var _nw193 _dafny.Array
				_ = _nw193
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw193 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Sequence = (func(_1025_v94 *C2, _1026_p0 _dafny.Int, _1027_v91 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1028_i6 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1025_v94).F27(), _1026_p0), _1027_v91)
						}
					})(_1015_v94, p0, _1012_v91)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw193 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw193).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw193).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1024_v103 = _nw193
				var _1029_v104 _dafny.Sequence
				_ = _1029_v104
				_1029_v104 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(820))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_1030_v94 *C2) func(_dafny.Int) _dafny.Int {
					return func(_1031_i7 _dafny.Int) _dafny.Int {
						return (_1030_v94).F27()
					}
				})(_1015_v94)))
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_1024_v103), 0))
				_ = _index166
				(_1024_v103).ArraySet1((_1029_v104), (_index166).Int())
				var _1032_v105 D13
				_ = _1032_v105
				_1032_v105 = Companion_D13_.Create_DC29_((_1022_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_1022_v101), 0))).Int()).(bool), _901_v0, _1022_v101)
				var _1033_v106 _dafny.Map
				_ = _1033_v106
				_1033_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1032_v105).Dtor_cf45(), (_1015_v94).F28())
				var _1034_v107 _dafny.Sequence
				_ = _1034_v107
				_1034_v107 = _dafny.SeqOf(_1019_v98)
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_1024_v103), 0))
				_ = _index167
				var _rhs198 D9 = (func() D9 {
					if (func() bool {
						if (_1033_v106).Contains(_1022_v101) {
							return (_1033_v106).Get(_1022_v101).(bool)
						}
						return (_1015_v94).F28()
					})() {
						return _1023_v102
					}
					return Companion_D9_.Create_DC17_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1022_v101).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_1022_v101), 0))).Int()).(bool), (_1015_v94).F28())).Cardinality())
				})()
				_ = _rhs198
				var _rhs199 _dafny.Sequence = _dafny.SeqOf(p0)
				_ = _rhs199
				var _rhs200 _dafny.Int = ((func() _dafny.Int {
					if (_1015_v94).F28() {
						return _dafny.IntOfUint32((_1034_v107).Cardinality())
					}
					return (_1015_v94).F27()
				})()).Minus(p0)
				_ = _rhs200
				var _rhs201 _dafny.Array = _1022_v101
				_ = _rhs201
				var _lhs165 _dafny.Array = _1024_v103
				_ = _lhs165
				var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_1024_v103), 0))
				_ = _lhs166
				var _lhs167 *GlobalState = globalState
				_ = _lhs167
				_1023_v102 = _rhs198
				(_lhs165).ArraySet1(_rhs199, (_lhs166).Int())
				_lhs167.F16 = _rhs200
				_1022_v101 = _rhs201
				(globalState).F18 = (p0).Cmp(p0) > 0
				(globalState).F0 = (_dafny.Zero).Minus((_1015_v94).F27())
			} else {
				var _1035_v108 _dafny.Array
				_ = _1035_v108
				var _nwElement0_42 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1012_v91, _1012_v91)
				_ = _nwElement0_42
				var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(21))
				_ = _nw194
				(_nw194).ArraySet1(_nwElement0_42, 0)
				(_nw194).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1012_v91, _1012_v91), 1)
				(_nw194).ArraySet1(_dafny.SeqOf((_dafny.Zero).Minus(p0)), 2)
				(_nw194).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1012_v91, _1012_v91), 3)
				(_nw194).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0), (Companion_Default___.SafeIndex(Companion_Default___.Fm5((_967_v57).Cardinality(), p0, globalState), _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1012_v91).Cardinality())), 4)
				(_nw194).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1012_v91, _1012_v91), 5)
				(_nw194).ArraySet1(_1012_v91, 6)
				(_nw194).ArraySet1(_1012_v91, 7)
				(_nw194).ArraySet1(_1012_v91, 8)
				(_nw194).ArraySet1(_1012_v91, 9)
				(_nw194).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1012_v91, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(229))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_1036_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1037_i8 _dafny.Int) _dafny.Int {
						return _1036_p0
					}
				})(p0)))), 10)
				(_nw194).ArraySet1(_1012_v91, 11)
				(_nw194).ArraySet1(_1012_v91, 12)
				(_nw194).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1012_v91, _1012_v91), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1012_v91, _1012_v91)).Cardinality()))).Uint32(), p0), 13)
				(_nw194).ArraySet1(_1012_v91, 14)
				(_nw194).ArraySet1(_dafny.SeqOf(p0), 15)
				(_nw194).ArraySet1(_1012_v91, 16)
				(_nw194).ArraySet1(_1012_v91, 17)
				(_nw194).ArraySet1(_1012_v91, 18)
				(_nw194).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(370))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_1038_v91 _dafny.Sequence, _1039_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1040_i9 _dafny.Int) _dafny.Int {
						return (_1038_v91).Select((Companion_Default___.SafeIndex(_1039_p0, _dafny.IntOfUint32((_1038_v91).Cardinality()))).Uint32()).(_dafny.Int)
					}
				})(_1012_v91, p0))), 19)
				(_nw194).ArraySet1(_1012_v91, 20)
				_1035_v108 = _nw194
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1035_v108), 0))
				_ = _index168
				(_1035_v108).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(483))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_1041_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1042_i10 _dafny.Int) _dafny.Int {
						return _1041_p0
					}
				})(p0))), (_index168).Int())
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_1035_v108), 0))
				_ = _index169
				(_1035_v108).ArraySet1(_1012_v91, (_index169).Int())
				_901_v0 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("rbg"), (func() _dafny.Sequence {
					if _901_v0 {
						return _1011_v90
					}
					return _1011_v90
				})())
				var _1043_v109 *C5
				_ = _1043_v109
				var _nw195 *C5 = New_C5_()
				_ = _nw195
				_nw195.Ctor__(_901_v0, (_1011_v90).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1011_v90).Cardinality()))).Uint32()).(_dafny.CodePoint), p0)
				_1043_v109 = _nw195
				var _1044_v110 *C0
				_ = _1044_v110
				var _nw196 *C0 = New_C0_()
				_ = _nw196
				_nw196.Ctor__()
				_1044_v110 = _nw196
				var _1045_v111 _dafny.Map
				_ = _1045_v111
				_1045_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ocqonr"), (_1043_v109).F26())
				_1045_v111 = (_1045_v111).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(928))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}(func(_1046_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})), !(false)))
			}
			var _1047_v112 _dafny.Sequence
			_ = _1047_v112
			_1047_v112 = _1011_v90
			var _1048_v113 _dafny.Map
			_ = _1048_v113
			_1048_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v0, _1047_v112)
			var _1049_v114 _dafny.Map
			_ = _1049_v114
			_1049_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1048_v113, _901_v0)
			(globalState).F0 = ((_1049_v114).Merge((_1049_v114).Merge(Companion_Default___.Fm30(_dafny.IntOfInt64(746), p0, _901_v0, _901_v0, globalState)))).Cardinality()
			(globalState).F2 = p0
		}
		var _1050_v115 _dafny.CodePoint
		_ = _1050_v115
		_1050_v115 = _dafny.CodePoint('p')
		var _1051_v116 D10
		_ = _1051_v116
		_1051_v116 = Companion_D10_.Create_DC19_(_1050_v115)
		var _1052_v117 _dafny.Sequence
		_ = _1052_v117
		_1052_v117 = _dafny.SeqOf(_1051_v116)
		var _1053_v118 _dafny.Sequence
		_ = _1053_v118
		_1053_v118 = _dafny.UnicodeSeqOfUtf8Bytes("okgjhwo")
		var _1054_v119 _dafny.Map
		_ = _1054_v119
		_1054_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1050_v115, _1053_v118)
		var _1055_v120 _dafny.Map
		_ = _1055_v120
		_1055_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_1052_v117, _1052_v117), (_1054_v119).Merge(_1054_v119))
		_1055_v120 = (_1055_v120).Update(_901_v0, _1054_v119)
		var _1056_v121 _dafny.Set
		_ = _1056_v121
		_1056_v121 = _dafny.SetOf(Companion_Default___.Fm17(!(false), true, false, globalState))
		var _1057_v122 _dafny.Map
		_ = _1057_v122
		_1057_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v0, _901_v0)
		var _1058_v123 _dafny.MultiSet
		_ = _1058_v123
		_1058_v123 = _dafny.MultiSetOf(Companion_Default___.Fm5((_1057_v122).Cardinality(), p0, globalState))
		_1056_v121 = func() _dafny.Set {
			var _coll42 = _dafny.NewBuilder()
			_ = _coll42
			for _iter47 := _dafny.Iterate((_1058_v123).Elements()); ; {
				_compr_42, _ok47 := _iter47()
				if !_ok47 {
					break
				}
				var _1059_v124 _dafny.Int
				_1059_v124 = interface{}(_compr_42).(_dafny.Int)
				if (_1058_v123).Contains(_1059_v124) {
					_coll42.Add(Companion_Default___.SafeDivisionInt(_1059_v124, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("e")).Cardinality())).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC5_(_dafny.SeqOf(false), _dafny.IntOfInt64(492), (func() _dafny.Set {
						var _coll43 = _dafny.NewBuilder()
						_ = _coll43
						for _iter48 := _dafny.Iterate((_dafny.SetOf((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-941)), _dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()), _dafny.IntOfInt64(774), _dafny.IntOfInt64(187))).Cardinality())).Elements()); ; {
							_compr_43, _ok48 := _iter48()
							if !_ok48 {
								break
							}
							var _1060_v125 _dafny.Int
							_1060_v125 = interface{}(_compr_43).(_dafny.Int)
							if (_dafny.SetOf((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-941)), _dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()), _dafny.IntOfInt64(774), _dafny.IntOfInt64(187))).Cardinality())).Contains(_1060_v125) {
								_coll43.Add((_1060_v125).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xn")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(780))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg67 _dafny.Int) interface{} {
										return coer67(arg67)
									}
								}(func(_1061_i12 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('h')
								}))).Cardinality()))).Cardinality())))
							}
						}
						return _coll43.ToSet()
					}()).Cardinality()), !(!(false)))).Cardinality())))
				}
			}
			return _coll42.ToSet()
		}()
		var _1062_v126 _dafny.Sequence
		_ = _1062_v126
		_1062_v126 = _dafny.SeqOf(_901_v0, !(_901_v0) || (_901_v0), _901_v0, _901_v0)
		if (_1062_v126).Select((Companion_Default___.SafeIndex((p0).Plus(p0), _dafny.IntOfUint32((_1062_v126).Cardinality()))).Uint32()).(bool) {
			(globalState).F0 = _dafny.IntOfInt64(330)
			var _1063_v127 *C0
			_ = _1063_v127
			var _nw197 *C0 = New_C0_()
			_ = _nw197
			_nw197.Ctor__()
			_1063_v127 = _nw197
			r2 = _dafny.Companion_Sequence_.Concatenate(_1053_v118, _1053_v118)
			_1058_v123 = _1058_v123
			var _1064_v128 _dafny.Map
			_ = _1064_v128
			_1064_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v0, p0)
			(globalState).F15 = ((func() _dafny.Int {
				if (_1064_v128).Contains(_901_v0) {
					return (_1064_v128).Get(_901_v0).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1053_v118, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1053_v118).Cardinality()))).Uint32(), _1050_v115)).Cardinality())
			})()).Plus(p0)
		} else {
			var _1065_v129 *C3
			_ = _1065_v129
			var _nw198 *C3 = New_C3_()
			_ = _nw198
			_nw198.Ctor__(true, _dafny.CodePoint('r'), p0)
			_1065_v129 = _nw198
			var _1066_v130 _dafny.Map
			_ = _1066_v130
			_1066_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1065_v129, (_1065_v129).Fm18(globalState))
			(globalState).F18 = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1065_v129, true)).Merge(_1066_v130)).Cardinality()).Cmp(p0) != 0
			var _1067_v131 _dafny.Sequence
			_ = _1067_v131
			_1067_v131 = _dafny.SeqOf(p0, p0)
			var _1068_v132 *C1
			_ = _1068_v132
			var _nw199 *C1 = New_C1_()
			_ = _nw199
			_nw199.Ctor__(_1050_v115, _dafny.IntOfUint32((_1067_v131).Cardinality()))
			_1068_v132 = _nw199
			var _1069_v133 _dafny.Map
			_ = _1069_v133
			_1069_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1068_v132)
			var _1070_v134 _dafny.Set
			_ = _1070_v134
			_1070_v134 = _dafny.SetOf((func() *C1 {
				if (_1069_v133).Contains(_1065_v129.F31) {
					return (_1069_v133).Get(_1065_v129.F31).(*C1)
				}
				return _1068_v132
			})(), _1068_v132)
			var _1071_v135 D7
			_ = _1071_v135
			_1071_v135 = Companion_D7_.Create_DC13_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1070_v134).Cardinality(), _901_v0)).Update((_dafny.Zero).Minus(p0), _1065_v129.F31))
			var _pat_let_tv10 = _1054_v119
			_ = _pat_let_tv10
			var _pat_let_tv11 = _901_v0
			_ = _pat_let_tv11
			var _rhs202 D7 = func(_pat_let14_0 D7) D7 {
				return func(_1072_dt__update__tmp_h1 D7) D7 {
					return func(_pat_let15_0 _dafny.Map) D7 {
						return func(_1073_dt__update_hcf19_h0 _dafny.Map) D7 {
							return Companion_D7_.Create_DC13_(_1073_dt__update_hcf19_h0)
						}(_pat_let15_0)
					}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv10).Cardinality(), _pat_let_tv11))
				}(_pat_let14_0)
			}(_1071_v135)
			_ = _rhs202
			var _rhs203 bool = _1065_v129.F31
			_ = _rhs203
			var _lhs168 *C3 = _1065_v129
			_ = _lhs168
			_1071_v135 = _rhs202
			_lhs168.F31 = _rhs203
			_901_v0 = !(_1065_v129.F31) || ((func() bool {
				if _1065_v129.F31 {
					return _1065_v129.F31
				}
				return _1065_v129.F31
			})())
			(globalState).F18 = _1065_v129.F31
			var _1074_v136 *C1
			_ = _1074_v136
			var _nw200 *C1 = New_C1_()
			_ = _nw200
			_nw200.Ctor__(_1050_v115, p0)
			_1074_v136 = _nw200
		}
		var _1075_v137 D9
		_ = _1075_v137
		_1075_v137 = Companion_D9_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v0, Companion_D3_.Create_DC6_(p0, p0, !(_901_v0))))
		var _pat_let_tv12 = _1056_v121
		_ = _pat_let_tv12
		var _pat_let_tv13 = _901_v0
		_ = _pat_let_tv13
		if func(_source15 D9) bool {
			if _source15.Is_DC17() {
				var _1076___mcc_h5 _dafny.Int = _source15.Get_().(D9_DC17).Cf25
				_ = _1076___mcc_h5
				var _1077_cf25 _dafny.Int = _1076___mcc_h5
				_ = _1077_cf25
				return ((_pat_let_tv12).Cardinality()).Cmp(_dafny.IntOfInt64(619)) > 0
			} else {
				var _1078___mcc_h6 _dafny.Map = _source15.Get_().(D9_DC16).Cf24
				_ = _1078___mcc_h6
				var _1079_cf24 _dafny.Map = _1078___mcc_h6
				_ = _1079_cf24
				return _pat_let_tv13
			}
		}(_1075_v137) {
			var _1080_v138 _dafny.Sequence
			_ = _1080_v138
			_1080_v138 = _1053_v118
			var _source16 _dafny.Sequence = _1080_v138
			_ = _source16
			var _1081___mcc_h4 _dafny.Sequence = _source16
			_ = _1081___mcc_h4
			var _1082_cf1 _dafny.Sequence = _1081___mcc_h4
			_ = _1082_cf1
			var _1083_v139 _dafny.Array
			_ = _1083_v139
			var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
			_ = _nw201
			_1083_v139 = _nw201
			var _1084_v140 _dafny.Sequence
			_ = _1084_v140
			_1084_v140 = _dafny.SeqOf(Companion_Default___.Fm5(p0, p0, globalState))
			var _1085_v141 _dafny.Map
			_ = _1085_v141
			_1085_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1083_v139, (func() _dafny.Int {
				if true {
					return _dafny.IntOfUint32((_1084_v140).Cardinality())
				}
				return p0
			})())
			r1 = (_1085_v141).Cardinality()
			_1058_v123 = (_1058_v123).Difference(_1058_v123)
			(_this).F23 = (_this.F23).Update(_dafny.IntOfUint32((_1082_cf1).Cardinality()), _901_v0)
			r2 = _1053_v118
			r1 = (p0).Plus(Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-498))))
			(globalState).F18 = !(!(true))
			var _1086_v142 _dafny.Array
			_ = _1086_v142
			var _nwElement0_43 bool = _901_v0
			_ = _nwElement0_43
			var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(4))
			_ = _nw202
			(_nw202).ArraySet1(_nwElement0_43, 0)
			(_nw202).ArraySet1(true, 1)
			(_nw202).ArraySet1(_901_v0, 2)
			(_nw202).ArraySet1(_901_v0, 3)
			_1086_v142 = _nw202
			var _1087_v143 _dafny.Sequence
			_ = _1087_v143
			_1087_v143 = _dafny.SeqOf(_1086_v142)
			var _1088_v144 _dafny.Map
			_ = _1088_v144
			_1088_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1087_v143).Cardinality()), _1050_v115)
			_1088_v144 = (_1088_v144).Update(_dafny.IntOfUint32((_1053_v118).Cardinality()), Companion_Default___.Fm10(_901_v0, _901_v0, false, _901_v0, globalState))
			var _1089_v145 _dafny.Array
			_ = _1089_v145
			var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw203
			_1089_v145 = _nw203
			var _1090_v146 _dafny.Map
			_ = _1090_v146
			_1090_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_901_v0, _1089_v145)
			if !((_1090_v146).Merge(_1090_v146)).Equals((_1090_v146).Merge((_1090_v146).Update(false, _1089_v145))) {
				var _1091_v147 _dafny.MultiSet
				_ = _1091_v147
				_1091_v147 = _dafny.MultiSetOf(_901_v0, _901_v0)
				_1091_v147 = (_1091_v147).Difference(_1091_v147)
				(globalState).F18 = _901_v0
				var _1092_v148 _dafny.Array
				_ = _1092_v148
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_32
				var _nw204 _dafny.Array
				_ = _nw204
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw204 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.Sequence = (func(_1093_v118 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1094_i13 _dafny.Int) _dafny.Sequence {
							return _1093_v118
						}
					})(_1053_v118)
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw204 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw204).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw204).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1092_v148 = _nw204
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_1092_v148), 0))
				_ = _index170
				(_1092_v148).ArraySet1(_1053_v118, (_index170).Int())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(110), _dafny.ArrayLen((_1092_v148), 0))
				_ = _index171
				(_1092_v148).ArraySet1(_1053_v118, (_index171).Int())
				var _1095_v149 _dafny.Array
				_ = _1095_v149
				var _nw205 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(8))
				_ = _nw205
				_1095_v149 = _nw205
				var _1096_v150 _dafny.Sequence
				_ = _1096_v150
				_1096_v150 = _dafny.SeqOf(_1056_v121, _dafny.SetOf((_dafny.Zero).Minus(p0)))
				var _1097_v151 _dafny.Sequence
				_ = _1097_v151
				_1097_v151 = _dafny.SeqOf(p0, p0)
				var _1098_v152 _dafny.Sequence
				_ = _1098_v152
				_1098_v152 = _dafny.SeqOf(_1097_v151)
				var _1099_v153 _dafny.Sequence
				_ = _1099_v153
				_1099_v153 = _dafny.SeqOf((_1098_v152).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1057_v122).Cardinality()), _dafny.IntOfUint32((_1098_v152).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf(p0))
				var _1100_v154 _dafny.Sequence
				_ = _1100_v154
				_1100_v154 = _dafny.SeqOf(_1056_v121, _1056_v121, (_1096_v150).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.IntOfUint32((_1096_v150).Cardinality()))).Uint32()).(_dafny.Set), _dafny.SetOf(_dafny.IntOfUint32((_1098_v152).Cardinality()), p0), _dafny.SetOf(p0, p0))
				var _1101_v155 _dafny.Map
				_ = _1101_v155
				_1101_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_901_v0)).Cardinality(), p0)
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1095_v149), 0))
				_ = _index172
				(_1095_v149).ArraySet1((_1096_v150).Select((Companion_Default___.SafeIndex((_1101_v155).Cardinality(), _dafny.IntOfUint32((_1096_v150).Cardinality()))).Uint32()).(_dafny.Set), (_index172).Int())
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1086_v142), 0))
				_ = _index173
				(_1086_v142).ArraySet1(_901_v0, (_index173).Int())
				var _1102_v156 _dafny.Sequence
				_ = _1102_v156
				_1102_v156 = _dafny.SeqOf(_this.F23, _this.F23, _this.F23)
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1095_v149), 0))
				_ = _index174
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1086_v142), 0))
				_ = _index175
				var _rhs204 _dafny.Set = (_1056_v121).Intersection((_1056_v121).Union(_dafny.SetOf(p0, p0)))
				_ = _rhs204
				var _rhs205 _dafny.Int = _dafny.IntOfInt64(374)
				_ = _rhs205
				var _rhs206 bool = _dafny.Companion_Sequence_.Contains(_1102_v156, (_this.F23).Merge(_this.F23))
				_ = _rhs206
				var _lhs169 _dafny.Array = _1095_v149
				_ = _lhs169
				var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1095_v149), 0))
				_ = _lhs170
				var _lhs171 *GlobalState = globalState
				_ = _lhs171
				var _lhs172 _dafny.Array = _1086_v142
				_ = _lhs172
				var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1086_v142), 0))
				_ = _lhs173
				(_lhs169).ArraySet1(_rhs204, (_lhs170).Int())
				_lhs171.F10 = _rhs205
				(_lhs172).ArraySet1(_rhs206, (_lhs173).Int())
				(globalState).F10 = ((func() _dafny.MultiSet {
					if (_1086_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_1086_v142), 0))).Int()).(bool) {
						return _1091_v147
					}
					return _1091_v147
				})()).Cardinality()
			} else {
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))
				_ = _index176
				(_1086_v142).ArraySet1(_901_v0, (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))
				_ = _index177
				(_1086_v142).ArraySet1(!(_901_v0), (_index177).Int())
				var _1103_v157 _dafny.Map
				_ = _1103_v157
				_1103_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), p0)
				var _1104_v158 _dafny.Sequence
				_ = _1104_v158
				_1104_v158 = _dafny.SeqOf(_1103_v157)
				var _1105_v159 _dafny.Sequence
				_ = _1105_v159
				_1105_v159 = _dafny.SeqOf(_dafny.IntOfInt64(794), ((_1103_v157).Merge((_1104_v158).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1104_v158).Cardinality()))).Uint32()).(_dafny.Map))).Cardinality(), p0, p0)
				(globalState).F15 = (_1105_v159).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1105_v159).Cardinality()))).Uint32()).(_dafny.Int)
				var _1106_v160 *C0
				_ = _1106_v160
				var _nw206 *C0 = New_C0_()
				_ = _nw206
				_nw206.Ctor__()
				_1106_v160 = _nw206
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))
				_ = _index178
				var _rhs207 bool = _901_v0
				_ = _rhs207
				var _rhs208 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1105_v159, _1105_v159)
				_ = _rhs208
				var _rhs209 *C0 = _1106_v160
				_ = _rhs209
				var _rhs210 bool = (_1086_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))).Int()).(bool)
				_ = _rhs210
				var _rhs211 bool = (_1058_v123).IsDisjointFrom((_1058_v123).Update(p0, Companion_Default___.Abs((_dafny.Zero).Minus((Companion_Default___.Fm31(_dafny.IntOfUint32((_1105_v159).Cardinality()), (_1086_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))).Int()).(bool), _901_v0, _dafny.IntOfUint32((_1053_v118).Cardinality()), globalState)).Cardinality()))))
				_ = _rhs211
				var _lhs174 *GlobalState = globalState
				_ = _lhs174
				var _lhs175 _dafny.Array = _1086_v142
				_ = _lhs175
				var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))
				_ = _lhs176
				var _lhs177 *GlobalState = globalState
				_ = _lhs177
				_901_v0 = _rhs207
				_lhs174.F19 = _rhs208
				_1106_v160 = _rhs209
				(_lhs175).ArraySet1(_rhs210, (_lhs176).Int())
				_lhs177.F18 = _rhs211
				var _1107_v161 _dafny.Array
				_ = _1107_v161
				var _nwElement0_44 bool = (_1086_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))).Int()).(bool)
				_ = _nwElement0_44
				var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(6))
				_ = _nw207
				(_nw207).ArraySet1(_nwElement0_44, 0)
				(_nw207).ArraySet1(_901_v0, 1)
				(_nw207).ArraySet1((_1086_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))).Int()).(bool), 2)
				(_nw207).ArraySet1(_901_v0, 3)
				(_nw207).ArraySet1(false, 4)
				(_nw207).ArraySet1(false, 5)
				_1107_v161 = _nw207
				var _1108_v162 _dafny.Map
				_ = _1108_v162
				_1108_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1107_v161, (_1058_v123).Cardinality())
				_1108_v162 = (_1108_v162).Update(_1107_v161, (p0).Minus((_dafny.Zero).Minus(p0)))
				var _1109_v163 _dafny.Array
				_ = _1109_v163
				var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
				_ = _nw208
				_1109_v163 = _nw208
				var _1110_v164 _dafny.Map
				_ = _1110_v164
				_1110_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("oo"), (_1086_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))).Int()).(bool))
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_1109_v163), 0))
				_ = _index179
				(_1109_v163).ArraySet1(_1110_v164, (_index179).Int())
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_1109_v163), 0))
				_ = _index180
				(_1109_v163).ArraySet1(((_1110_v164).Update(_1053_v118, (_1086_v142).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_1086_v142), 0))).Int()).(bool))).Merge(_1110_v164), (_index180).Int())
			}
		} else {
			var _1111_v165 _dafny.Array
			_ = _1111_v165
			var _nw209 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw209
			_1111_v165 = _nw209
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1111_v165), 0))
			_ = _index181
			(_1111_v165).ArraySet1(_901_v0, (_index181).Int())
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1111_v165), 0))
			_ = _index182
			(_1111_v165).ArraySet1(!(_901_v0), (_index182).Int())
			(globalState).F16 = Companion_Default___.SafeModuloInt((_1057_v122).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_1111_v165).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1111_v165), 0))).Int()).(bool))).Cardinality()))
			(globalState).F18 = !(Companion_Default___.Fm0(p0, (p0).Cmp(p0) < 0, false, p0, globalState))
			_1053_v118 = (func() _dafny.Sequence {
				if (_1111_v165).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1111_v165), 0))).Int()).(bool) {
					return _1053_v118
				}
				return _dafny.Companion_Sequence_.Concatenate(_1053_v118, _1053_v118)
			})()
			var _1112_v166 *C3
			_ = _1112_v166
			var _nw210 *C3 = New_C3_()
			_ = _nw210
			_nw210.Ctor__(_901_v0, _1050_v115, p0)
			_1112_v166 = _nw210
			_1112_v166 = (func() *C3 {
				if (_1111_v165).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_1111_v165), 0))).Int()).(bool) {
					return _1112_v166
				}
				return _1112_v166
			})()
		}
		r0 = p0
		r1 = Companion_Default___.Fm17(_901_v0, _901_v0, (_1062_v126).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1053_v118, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.IntOfUint32((_1053_v118).Cardinality()))).Uint32(), Companion_Default___.Fm10(_901_v0, _901_v0, _901_v0, _901_v0, globalState))).Cardinality()), _dafny.IntOfUint32((_1062_v126).Cardinality()))).Uint32()).(bool), globalState)
		r2 = _1053_v118
		return r0, r1, r2
	}
}

// End of class C6
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
